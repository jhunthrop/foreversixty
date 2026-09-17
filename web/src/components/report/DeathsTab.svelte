<!-- web/src/components/report/DeathsTab.svelte -->
<!-- The deaths recap. One card per death, in the order they happened: the killing blow,
     the last ten damage events with the health they left, the auras that were up and the
     ones that had just fallen off, a button that sets the window to the twenty seconds
     before it, and the link into the planner.
     Cards rather than a table at every width: a death is read top to bottom, and the
     last-ten list is the whole point of the view. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import {
    classColorVar,
    formatAmount,
    formatDuration,
    formatDurationPrecise,
  } from '../../lib/report/format';
  import { believableHealth, healthPct } from '../../lib/report/death-health';
  import { plannerLinkFor } from '../../lib/report/planner-link';
  import CopyCsv from './CopyCsv.svelte';
  import type { CastRow, CombatantRow, DamageRef, Death, HealRef, PullMark } from '../../lib/report/types';
  import { deathWindow, type TimeWindow } from '../../lib/report/window';

  let {
    deaths,
    casts = [],
    open = [],
    onPatch = () => {},
    onSelectPlayer = undefined,
    durationMs,
    pulls = [],
    combatants,
    classOf,
    dataBuild,
    treeSizesFor,
    onWindow,
  }: {
    /** The night's pulls, so a death over the night can say when in its pull it came. */
    pulls?: PullMark[];
    deaths: Death[];
    /** The fight's casts, to tell a player still dead from one raised: a cast means alive. */
    casts?: CastRow[];
    open?: string[];
    onPatch?: (patch: { openDeaths?: string[] }) => void;
    /** Narrows the page to one player, from their name on the card. */
    onSelectPlayer?: (guid: string) => void;
    durationMs: number;
    combatants: CombatantRow[];
    classOf: Map<string, string>;
    dataBuild: string;
    /** Talents per tree for a class, or [] when the planner has no data for it. */
    treeSizesFor: (className: string | undefined) => number[];
    onWindow: (window: TimeWindow) => void;
  } = $props();

  const ordered = $derived([...deaths].sort((a, b) => a.at_ms - b.at_ms));
  /** One line per death: when, who, what killed them and from whom, and the pull over a night. */
  function csvLines(): string[][] {
    return [
      ['At', 'Player', 'Killing blow', 'From', 'Amount', 'Overkill', 'Pull'],
      ...ordered.map((death) => [
        // On the pull's clock, the same as the cards: over the night a death is in a pull.
        formatDurationPrecise(inPull(death)),
        // The full name, realm included, the way every other CSV on the report writes it.
        death.name,
        death.killing_blow === undefined ? '' : death.killing_blow.spell_name || 'Melee',
        death.killing_blow === undefined ? '' : splitUnitName(death.killing_blow.source_name).name,
        death.killing_blow === undefined ? '' : String(death.killing_blow.amount),
        death.killing_blow?.overkill === undefined ? '' : String(death.killing_blow.overkill),
        death.label ?? '',
      ]),
    ];
  }
  const FOLD_ABOVE = 6;
  /** Cards the reader has opened by hand; every card is open while there are few. */
  /** The cards opened by hand, from the url, so a pasted link opens the same ones. */
  const opened = $derived(new Set(open));
  // A link that arrives with a card already open (Mechanics' Deaths links) lands at the
  // top of a long page; bring the first open card into view once, on arrival.
  let scrolledTo = '';
  $effect(() => {
    const first = open[0];
    if (first === undefined || first === scrolledTo) return;
    scrolledTo = first;
    document.getElementById(`death-${first}`)?.scrollIntoView({ block: 'start' });
  });
  const foldAll = $derived(ordered.length > FOLD_ABOVE);
  const isOpen = (death: Death): boolean => !foldAll || opened.has(`${death.guid}-${death.at_ms}`);
  function toggle(death: Death): void {
    const key = `${death.guid}-${death.at_ms}`;
    onPatch({ openDeaths: opened.has(key) ? open.filter((entry) => entry !== key) : [...open, key] });
  }

  type LastEvent =
    { kind: 'damage'; at_ms: number; hit: DamageRef } | { kind: 'heal'; at_ms: number; heal: HealRef };

  /**
   * The damage and the heals in one list, in order: whether anyone was healing the
   * player while the damage came in is the healer's whole question, and two separate
   * lists make the reader interleave them by hand. Summaries from engines before 0.2.0
   * carry no heals, and then this is the damage alone.
   */
  /** The row that is the killing blow: same instant, spell and amount. */
  function isKillingBlow(death: Death, hit: DamageRef): boolean {
    const blow = death.killing_blow;
    return (
      blow !== undefined &&
      blow.at_ms === hit.at_ms &&
      blow.spell_id === hit.spell_id &&
      blow.amount === hit.amount
    );
  }

  /**
   * True when the last recorded hit demonstrably left them alive: its health-after is
   * known and above zero. Then the log never showed what killed them.
   */
  /**
   * The recorded blow was not what ended them: it left health behind and carried no
   * overkill. A blow with overkill was lethal whatever health the client's rounding
   * left on the line (a Gloom Squall at 14,141 with 14,140 over reads "1 of 32,540").
   */
  function lethalHitMissing(death: Death, card: { events: LastEvent[]; pcts: (number | null)[] }): boolean {
    const blow = death.killing_blow;
    if (blow === undefined || (blow.overkill ?? 0) > 0) return false;
    // The card's believed reading decides, not the raw one: a blow whose reading the card
    // could not believe is read as the hit that ended them, never as "no lethal hit".
    const left = blowPct(death, card);
    return left !== null && left > 0;
  }

  function lastEvents(death: Death): LastEvent[] {
    // Nothing after the death itself: a hit the log wrote after UNIT_DIED is noise.
    const cutoff = death.killing_blow?.at_ms ?? death.at_ms;
    const events: LastEvent[] = [
      ...death.last
        .filter((hit) => hit.at_ms <= cutoff)
        .map((hit): LastEvent => ({ kind: 'damage', at_ms: hit.at_ms, hit })),
      ...(death.heals ?? [])
        .filter((heal) => heal.at_ms <= cutoff)
        .map((heal): LastEvent => ({ kind: 'heal', at_ms: heal.at_ms, heal })),
    ];
    return events.sort((a, b) => a.at_ms - b.at_ms);
  }

  function healingBefore(death: Death): number {
    return (death.heals ?? []).reduce((sum, heal) => sum + heal.amount - (heal.overheal ?? 0), 0);
  }

  /** The combat log's null GUID: falling, drowning, lava, the fight's own environment. */
  const NULL_GUID = /^0+$/;

  function sourceName(guid: string, name: string): string {
    if (NULL_GUID.test(guid) || name === '' || NULL_GUID.test(name)) return 'Environment';
    return splitUnitName(name).name;
  }

  /** Seconds before the death a hit landed, as "-3.2s"; "0.0s" for the killing blow. */
  function beforeDeath(death: Death, atMs: number): string {
    return `-${((death.at_ms - atMs) / 1000).toFixed(1)}s`;
  }

  /**
   * What a hit took off the health bar. The engine's amount is what landed: a swing is
   * read from its landed line and an absorb miss lands for nothing, with the absorb kept
   * beside it, so nothing is subtracted here -- subtracting the absorb again read a
   * 15,308 hit as 14,021.
   */
  function landed(hit: DamageRef): number {
    return Math.max(0, hit.amount);
  }

  /**
   * One card's rows with the health reading each of them can be believed on. The client
   * writes a full bar on a swing a shield ate whole and keeps writing it for a swing or
   * two after, so a reading that rises on a damage row with no heal before it is the
   * client's and not the player's: death-health.ts drops it, and the bar, the run-up
   * sentence and the "no lethal hit" sentence all read the same list.
   */
  function cardHealth(death: Death): { events: LastEvent[]; pcts: (number | null)[] } {
    const events = lastEvents(death);
    return {
      events,
      pcts: believableHealth(
        events.map((event) =>
          event.kind === 'heal'
            ? { kind: 'heal' as const, hp_after: event.heal.hp_after, max_hp: event.heal.max_hp }
            : { kind: 'damage' as const, hp_after: event.hit.hp_after, max_hp: event.hit.max_hp },
        ),
      ),
    };
  }
  /** The believable reading on the killing blow's own row, when the card has one. */
  function blowPct(death: Death, card: { events: LastEvent[]; pcts: (number | null)[] }): number | null {
    const at = card.events.findIndex((event) => event.kind === 'damage' && isKillingBlow(death, event.hit));
    return at === -1
      ? death.killing_blow === undefined
        ? null
        : healthPct(death.killing_blow)
      : card.pcts[at];
  }

  /**
   * "Burst" or "bleed": the one word a healer wants first. The last hits are the engine's
   * final ten damage events; if the player went from above 80% to dead inside three
   * seconds, no heal was going to land in time, and that is a different conversation
   * from a health bar that drained over fifteen seconds while nobody was healing it.
   */
  function shape(death: Death, pcts: (number | null)[]): string | null {
    const hits = death.last;
    if (hits.length < 2) return null;
    const first = hits[0];
    // The health the run-up started from is the highest the recap saw, not the first
    // row's: heals between hits can lift it, and "from 4%" when they were at 47% two
    // hits later is the wrong story.
    // The heals count too: a Shatter that lifted them to 50% is part of the span's story.
    // Only the readings the card believes, though: a client's post-absorb full bar named
    // as the peak told a healer the player was fine at 73% when they were under 30%.
    const firstPct = pcts.reduce<number | null>(
      (highest, pct) => (pct === null ? highest : Math.max(highest ?? 0, pct)),
      null,
    );
    if (firstPct === null) return null;
    const spanMs = death.at_ms - first.at_ms;
    const total = hits.reduce((sum, hit) => sum + landed(hit), 0);
    const spanText = formatDuration(spanMs);
    const healed = healingBefore(death);
    const healedText =
      death.heals === undefined
        ? ''
        : healed > 0
          ? ` ${formatAmount(healed)} of healing landed in the same span.`
          : ' No healing landed on them in that span.';
    const from = `the highest they stood at in that span was ${Math.round(firstPct)}%`;
    if (spanMs <= 3000 && firstPct >= 60)
      return `Burst: they took ${formatAmount(total)} in ${spanText}; ${from}.${healedText}`;
    return `They took ${formatAmount(total)} over ${spanText}; ${from}.${healedText}`;
  }
  function linkFor(death: Death): ReturnType<typeof plannerLinkFor> {
    const combatant = combatants.find((row) => row.guid === death.guid);
    if (combatant === undefined) return null;
    const className = death.class ?? classOf.get(death.guid);
    return plannerLinkFor({ dataBuild, className, treeSizes: treeSizesFor(className), combatant });
  }

  /** When in its own pull a death came: the night's clock, less the pull's start. */
  function inPull(death: Death): number {
    const pull = pulls.find((mark) => mark.start_ms <= death.at_ms && death.at_ms < mark.end_ms);
    return death.at_ms - (pull?.start_ms ?? 0);
  }

  /** The players who had died earlier in the same pull and cast nothing since: still dead at this one. */
  function deadAt(death: Death): Death[] {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Set<string>();
    return deaths
      .filter(
        (other) => other.guid !== death.guid && other.at_ms < death.at_ms && other.label === death.label,
      )
      .filter((other) => {
        const raised = casts
          .filter((row) => row.guid === other.guid)
          .some((row) => row.sequence.some((at) => at > other.at_ms && at <= death.at_ms));
        if (raised || seen.has(other.guid)) return false;
        seen.add(other.guid);
        return true;
      });
  }
</script>

{#if ordered.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">Nobody died in this window.</p>
{:else}
  <!-- Past a handful, the cards start folded to their first line: forty-four open cards
       over a whole night were a forty-thousand-pixel page. -->
  <CopyCsv lines={csvLines} />
  <ul class="flex flex-col gap-4" data-testid="deaths-tab">
    <!-- A death's last hits can repeat a spell in one millisecond (a DoT tick and its
         crit, a cleave), so the row key carries its index too; a bare timestamp-and-spell
         key threw on the duplicate and blanked the whole tab. -->
    {#each ordered as death (`${death.guid}-${death.at_ms}`)}
      {@const link = linkFor(death)}
      {@const alreadyDead = deadAt(death)}
      {@const card = cardHealth(death)}
      <li
        class="border-line rounded-panel bg-raised flex flex-col gap-3 border p-3"
        data-testid={`death-${death.guid}`}
        id={`death-${death.guid}-${death.at_ms}`}
      >
        <div class="flex flex-wrap items-baseline gap-x-3 gap-y-1">
          {#if foldAll}
            <button
              type="button"
              class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-0"
              aria-expanded={isOpen(death)}
              aria-label={`${isOpen(death) ? 'Fold' : 'Open'} ${splitUnitName(death.name).name}’s death at ${formatDuration(death.at_ms)}`}
              title={`${isOpen(death) ? 'Fold' : 'Open'} this death card`}
              data-testid="death-toggle"
              onclick={() => toggle(death)}>{isOpen(death) ? 'Fold' : 'Open'}</button
            >
          {/if}
          <span
            class="text-[15px] font-semibold"
            style={`color: ${classColorVar(death.class ?? classOf.get(death.guid))}`}
          >
            {#if onSelectPlayer}
              <button
                type="button"
                class="inline-flex min-h-11 items-center underline-offset-2 hover:underline md:min-h-0"
                style={`color: ${classColorVar(death.class ?? classOf.get(death.guid))}`}
                title="Show only this player"
                onclick={() => onSelectPlayer(death.guid)}>{splitUnitName(death.name).name}</button
              >
            {:else}
              {splitUnitName(death.name).name}
            {/if}
          </span>
          {#if death.label}
            <span class="text-muted text-[13px]" data-testid="death-label"
              >{death.label} · <span class="tabular font-mono">{formatDuration(inPull(death))}</span></span
            >
          {:else}
            <span class="text-muted tabular font-mono text-[13px]">{formatDuration(death.at_ms)}</span>
          {/if}
          {#if alreadyDead.length > 0}
            <span class="text-muted text-[13px]" data-testid="death-already-dead"
              >already dead: {alreadyDead
                .map((entry) => `${splitUnitName(entry.name).name} (since ${formatDuration(inPull(entry))})`)
                .join(', ')}</span
            >
          {/if}
          {#if death.killing_blow}
            <span class="text-[13px]">
              <span
                title={lethalHitMissing(death, card)
                  ? 'No damage line ended them: the death came from something the log does not write as damage (a fall, an instant kill, a timer running out), so this is the last hit before it'
                  : 'The hit that ended them'}
                >{lethalHitMissing(death, card) ? 'last hit by' : 'killed by'}</span
              >
              {sourceName(death.killing_blow.source_guid, death.killing_blow.source_name)} ·
              {NULL_GUID.test(death.killing_blow.source_guid) && death.killing_blow.spell_name === ''
                ? 'a fall, a hazard or an untracked source'
                : death.killing_blow.spell_name === ''
                  ? 'Melee'
                  : death.killing_blow.spell_name} ·
              <span class="tabular font-mono">{formatAmount(death.killing_blow.amount)}</span>
              {#if death.killing_blow.overkill}
                <span
                  class="text-muted"
                  title="Overkill: the part of that hit beyond the health they had left. The hit is counted whole here."
                  >(<span class="tabular font-mono">{formatAmount(death.killing_blow.overkill)}</span> overkill)</span
                >
              {/if}
            </span>
          {/if}
          {#if death.last.some((hit) => hit.max_hp)}
            <span
              class="text-muted text-[13px]"
              data-testid="death-max-hp"
              title="Their full health bar, for reading the hits against"
            >
              max health
              <span class="tabular font-mono"
                >{formatAmount(Math.max(...death.last.map((hit) => hit.max_hp ?? 0)))}</span
              >
            </span>
          {/if}
          {#if death.release_ms}
            <span class="text-muted text-[13px]">
              released after
              <span class="tabular font-mono">{formatDuration(death.release_ms - death.at_ms)}</span>
            </span>
          {/if}
        </div>

        {#if isOpen(death)}
          {#if shape(death, card.pcts)}
            <p class="text-[13px]" data-testid="death-shape">{shape(death, card.pcts)}</p>
          {/if}
          {#if death.killing_blow && lethalHitMissing(death, card)}
            <p class="text-muted text-[13px]" data-testid="death-unlogged">
              The log shows no lethal hit: the last recorded hit left them at
              {#if blowPct(death, card) !== null}
                <span class="tabular font-mono">{Math.round(blowPct(death, card) ?? 0)}%</span>,
              {:else}
                unknown health,
              {/if}
              and they died
              <span class="tabular font-mono">{beforeDeath(death, death.killing_blow.at_ms).slice(1)}</span> later.
            </p>
          {/if}

          <div class="overflow-x-auto">
            <table class="w-full table-fixed text-[13px] md:table-auto">
              <caption class="label text-muted text-left">
                {death.heals === undefined ? 'Last hits' : 'Last hits and heals'}
              </caption>
              <thead>
                <tr class="text-muted label">
                  <th class="w-14 py-1 pr-2 text-left font-bold md:w-auto" title="Seconds before the death"
                    >Before</th
                  >
                  <th class="py-1 pr-3 text-left font-bold">Ability</th>
                  <th class="hidden py-1 pr-3 text-left font-bold md:table-cell">From</th>
                  <th class="w-20 py-1 pr-3 text-right font-bold md:w-auto">Amount</th>
                  <th
                    class="w-24 py-1 text-left font-bold md:w-auto"
                    title="Health left after the hit; blank when the log's reading cannot be believed (it rose above the reading before it with no heal between)"
                    >Health after</th
                  >
                </tr>
              </thead>
              <tbody>
                {#each card.events as event, i (`${event.at_ms}-${i}`)}
                  {#if event.kind === 'heal'}
                    {@const heal = event.heal}
                    {@const healed = card.pcts[i]}
                    <tr class="border-line-soft border-b" data-testid="death-heal">
                      <td
                        class="text-muted tabular py-1 pr-3 font-mono"
                        title={formatDurationPrecise(heal.at_ms)}>{beforeDeath(death, heal.at_ms)}</td
                      >
                      <td class="text-kill py-1 pr-3 break-words md:truncate" title={heal.spell_name}
                        >{heal.spell_name}<span class="text-muted block text-[11px] md:hidden"
                          >from {splitUnitName(heal.source_name).name}</span
                        ></td
                      >
                      <td class="text-muted hidden truncate py-1 pr-3 md:table-cell"
                        >{splitUnitName(heal.source_name).name}</td
                      >
                      <td class="text-kill tabular py-1 pr-3 text-right font-mono"
                        >+{formatAmount(heal.amount - (heal.overheal ?? 0))}{#if heal.overheal}
                          <span class="text-muted block text-[11px] md:inline" title="Overhealing"
                            >({formatAmount(heal.overheal)} over)</span
                          >{/if}</td
                      >
                      <td class="w-[30%] py-1">
                        {#if healed !== null}
                          <span class="flex items-center gap-2">
                            <span
                              class="bg-line-soft block h-[8px] flex-1"
                              title={`${heal.hp_after ?? 0} of ${heal.max_hp}`}
                            >
                              <span
                                class="block h-full {healed <= 35
                                  ? 'bg-death'
                                  : healed <= 65
                                    ? 'bg-ember'
                                    : 'bg-kill'}"
                                style={`width: ${healed}%`}
                              ></span>
                            </span>
                            <span class="text-muted tabular w-[36px] text-right font-mono text-[12px]"
                              >{Math.round(healed)}%</span
                            >
                          </span>
                        {/if}
                      </td>
                    </tr>
                  {:else}
                    {@const hit = event.hit}
                    {@const lethal = isKillingBlow(death, hit)}
                    <!-- The last recorded hit reads as the kill only when it was one: a hit that
                         left health behind keeps its own figure, as the sentence above says. -->
                    {@const pct = lethal && !lethalHitMissing(death, card) ? 0 : card.pcts[i]}
                    <tr class="border-line-soft border-b">
                      <td
                        class="text-muted tabular py-1 pr-3 font-mono"
                        title={formatDurationPrecise(hit.at_ms)}>{beforeDeath(death, hit.at_ms)}</td
                      >
                      <td
                        class="py-1 pr-3 break-words md:truncate"
                        title={hit.spell_name === '' ? 'Melee' : hit.spell_name}
                        >{hit.spell_name === '' ? 'Melee' : hit.spell_name}<span
                          class="text-muted block text-[11px] md:hidden"
                          >from {sourceName(hit.source_guid, hit.source_name)}</span
                        ></td
                      >
                      <td
                        class="text-muted hidden truncate py-1 pr-3 md:table-cell"
                        title={sourceName(hit.source_guid, hit.source_name)}
                        >{sourceName(hit.source_guid, hit.source_name)}</td
                      >
                      <td
                        class="tabular py-1 pr-3 text-right font-mono"
                        class:text-muted={landed(hit) === 0}
                        title={landed(hit) === 0 ? 'Absorbed in full: nothing landed' : undefined}
                        >{formatAmount(landed(hit))}{#if (hit.absorbed ?? 0) > 0}
                          <span class="text-muted block text-[11px] md:inline" title="Soaked by a shield"
                            >({formatAmount(hit.absorbed ?? 0)} absorbed)</span
                          >{/if}{#if lethal && death.killing_blow?.overkill}
                          <span class="text-muted block text-[11px] md:inline" title="Overkill"
                            >({formatAmount(death.killing_blow.overkill)} overkill)</span
                          >{/if}</td
                      >
                      <td class="w-[30%] py-1">
                        {#if pct !== null}
                          <span class="flex items-center gap-2">
                            <span
                              class="bg-line-soft block h-[8px] flex-1"
                              title={`${hit.hp_after ?? 0} of ${hit.max_hp}`}
                            >
                              <span
                                class="block h-full {pct <= 35
                                  ? 'bg-death'
                                  : pct <= 65
                                    ? 'bg-ember'
                                    : 'bg-kill'}"
                                style={`width: ${pct}%`}
                              ></span>
                            </span>
                            <span class="text-muted tabular w-[36px] text-right font-mono text-[12px]"
                              >{Math.round(pct)}%</span
                            >
                          </span>
                        {/if}
                      </td>
                    </tr>
                  {/if}
                {/each}
              </tbody>
            </table>
          </div>

          {#if death.auras_held.length > 0 || death.auras_lost.some((aura) => aura.at_ms < death.at_ms - 50)}
            <p class="text-[13px]">
              {#if death.auras_held.length > 0}
                <span class="label text-muted" title="Buffs and debuffs on them at the moment they died"
                  >Up</span
                >
                {death.auras_held.map((aura) => aura.name || `Spell #${aura.spell_id}`).join(', ')}
              {/if}
              {#if death.auras_lost.some((aura) => aura.at_ms < death.at_ms - 50)}
                <span class="label text-muted ml-3" title="Auras that dropped in the seconds before the death"
                  >Just lost</span
                >
                <!-- Only what dropped before the death: the death itself strips every aura,
                     and listing those as "just lost" said nothing. -->
                {#each death.auras_lost.filter((aura) => aura.at_ms < death.at_ms - 50) as aura, i (`${aura.spell_id}-${i}`)}
                  {#if i > 0},{/if}
                  {aura.name}
                  <span
                    class="text-muted tabular font-mono text-[12px]"
                    title={formatDurationPrecise(aura.at_ms)}>{beforeDeath(death, aura.at_ms)}</span
                  >
                {/each}
              {/if}
            </p>
          {/if}

          <div class="flex flex-wrap gap-2">
            <button
              type="button"
              class="border-line-warm rounded-control text-text inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
              data-testid="death-window"
              onclick={() => onWindow(deathWindow(death.at_ms, durationMs))}
            >
              The 20s before this
            </button>
            {#if link}
              <a
                class="border-line-warm-strong rounded-control text-strong inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
                href={link.href}
                data-testid="death-build-link"
              >
                {link.label}
              </a>
            {/if}
          </div>
        {/if}
      </li>
    {/each}
  </ul>
{/if}
