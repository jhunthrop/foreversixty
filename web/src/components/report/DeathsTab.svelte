<!-- web/src/components/report/DeathsTab.svelte -->
<!-- The Deaths tab's Plate (logs spec 2026-11-07 6.2, the SHOULDs as ruled): one block per death
     in the scope, in the order they happened. A Row: the time, the crest, the name, "Killed by
     {source} · {ability} · {amount} ({overkill} overkill) · {max} health", and The 20 s before it ›
     (the window). Then one StatTable, Last hits and heals, in time order: a hit reads its ability,
     "{time} · from {source}", −amount and the health after as numbers ("74 / 114") with a green
     Meter, "dead" on the killing blow; a heal received is a green row with +amount. The foot: the
     run-up in one sentence, what they held, and "No healing landed in the last 10 seconds." when
     none did. Past six deaths each block folds to its Row. -->
<script lang="ts">
  import EmptyLine from './EmptyLine.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { unitLabel } from '../../lib/report/names.svelte';
  import {
    classColorVar,
    formatAmount,
    formatDuration,
    formatDurationPrecise,
  } from '../../lib/report/format';
  import { believableHealth, healthPct } from '../../lib/report/death-health';
  import { plannerLinkFor } from '../../lib/report/planner-link';
  import type { StatFigure } from '../../lib/ui/stat-table';
  import Meter from '../ui/Meter.svelte';
  import Plate from '../ui/Plate.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import ClassIcon from './ClassIcon.svelte';
  import SpellIcon from './SpellIcon.svelte';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import type { CastRow, CombatantRow, DamageRef, Death, HealRef, PullMark } from '../../lib/report/types';
  import { deathWindow, type TimeWindow } from '../../lib/report/window';

  let {
    deaths,
    casts = [],
    open = [],
    onPatch = () => {},
    durationMs,
    pulls = [],
    combatants,
    classOf,
    dataBuild,
    treeSizesFor,
    onWindow,
    spellIcon,
  }: {
    /** The night's pulls, so a death over the night can say when in its pull it came. */
    pulls?: PullMark[];
    deaths: Death[];
    /** The fight's casts, to tell a player still dead from one raised: a cast means alive. */
    casts?: CastRow[];
    open?: string[];
    onPatch?: (patch: { openDeaths?: string[] }) => void;
    durationMs: number;
    combatants: CombatantRow[];
    classOf: Map<string, string>;
    dataBuild: string;
    /** Talents per tree for a class, or [] when the planner has no data for it. */
    treeSizesFor: (className: string | undefined) => number[];
    onWindow: (window: TimeWindow) => void;
    spellIcon: SpellIconOf;
  } = $props();

  const ordered = $derived([...deaths].sort((a, b) => a.at_ms - b.at_ms));
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
  const COLUMNS = [
    { id: 'amount', label: 'Amount' },
    { id: 'health', label: 'Health' },
  ] as const;

  /** "74 / 114" after a hit or a heal; '' when the log carries no health reading. */
  function healthText(after: number | undefined, max: number | undefined): string {
    return after === undefined || max === undefined || max === 0
      ? ''
      : `${formatAmount(after)} / ${formatAmount(max)}`;
  }
  function hitFigures(death: Death, hit: DamageRef, lethal: boolean): StatFigure[] {
    const health =
      lethal && !lethalHitMissing(death, cardHealth(death)) ? 'dead' : healthText(hit.hp_after, hit.max_hp);
    return [
      { value: `−${formatAmount(landed(hit))}`, muted: landed(hit) === 0 },
      { value: health || '–', muted: health === '' },
    ];
  }
  function healFigures(heal: HealRef): StatFigure[] {
    const health = healthText(heal.hp_after, heal.max_hp);
    return [
      { value: `+${formatAmount(heal.amount - (heal.overheal ?? 0))}` },
      { value: health || '–', muted: health === '' },
    ];
  }
  function maxHealth(death: Death): number {
    return Math.max(0, ...death.last.map((hit) => hit.max_hp ?? 0), death.killing_blow?.max_hp ?? 0);
  }
  /** The block's Row sub line: what killed them, how hard, over how much, and how full their bar was. */
  function killedBy(death: Death, card: { events: LastEvent[]; pcts: (number | null)[] }): string {
    const blow = death.killing_blow;
    if (blow === undefined) return '';
    const spell =
      NULL_GUID.test(blow.source_guid) && blow.spell_name === ''
        ? 'a fall, a hazard or an untracked source'
        : blow.spell_name === ''
          ? 'Melee'
          : blow.spell_name;
    const overkill = blow.overkill ? ` (${formatAmount(blow.overkill)} overkill)` : '';
    const max = maxHealth(death);
    return [
      `${lethalHitMissing(death, card) ? 'Last hit by' : 'Killed by'} ${sourceName(blow.source_guid, blow.source_name)}`,
      spell,
      `${formatAmount(blow.amount)}${overkill}`,
      max > 0 ? `${formatAmount(max)} health` : '',
    ]
      .filter((part) => part !== '')
      .join(' · ');
  }
  /** The block's foot: the run-up, what they held, and whether healing landed in the last 10 s. */
  function foot(death: Death, card: { events: LastEvent[]; pcts: (number | null)[] }): string {
    const held = death.auras_held.map((aura) => aura.name || `Spell #${aura.spell_id}`);
    const healed = (death.heals ?? []).some((heal) => heal.at_ms >= death.at_ms - 10_000);
    const already = deadAt(death);
    return [
      shape(death, card.pcts) ?? '',
      held.length > 0 ? `Holding: ${held.join(', ')}.` : '',
      death.heals !== undefined && !healed ? 'No healing landed in the last 10 seconds.' : '',
      already.length > 0
        ? `Already dead: ${already.map((entry) => `${splitUnitName(entry.name).name} (since ${formatDuration(inPull(entry))})`).join(', ')}.`
        : '',
    ]
      .filter((part) => part !== '')
      .join(' ');
  }
</script>

<Section title="Deaths" meta={`${ordered.length} in this fight`} testid="deaths-tab">
  <Plate label="Deaths">
    {#if ordered.length === 0}
      <EmptyLine text="Nobody died." testid="table-empty" />
    {/if}
    {#each ordered as death (`${death.guid}-${death.at_ms}`)}
      {@const link = linkFor(death)}
      {@const card = cardHealth(death)}
      {@const colour = classColorVar(death.class ?? classOf.get(death.guid))}
      {@const footLine = foot(death, card)}
      <div class="death-block" data-testid={`death-${death.guid}`} id={`death-${death.guid}-${death.at_ms}`}>
        <Row title={unitLabel(death.name).text} sub={killedBy(death, card)}>
          {#snippet lead()}<span class="text-muted text-[13px] tabular-nums" data-testid="death-label"
              >{#if death.label}{death.label} ·
              {/if}{formatDuration(inPull(death))}</span
            ><ClassIcon className={death.class ?? classOf.get(death.guid)} size={32} />{/snippet}
          {#snippet titleContent()}<span style:color={colour}>{unitLabel(death.name).text}</span>{/snippet}
          {#snippet trail()}
            {#if foldAll}
              <button
                type="button"
                class="text-nav text-[14px] font-semibold"
                aria-expanded={isOpen(death)}
                data-testid="death-toggle"
                onclick={() => toggle(death)}>{isOpen(death) ? 'Fold' : 'Open'}</button
              >
            {/if}
            <button
              type="button"
              class="text-gold text-[14px] font-semibold whitespace-nowrap"
              data-testid="death-window"
              onclick={() => onWindow(deathWindow(death.at_ms, durationMs))}>The 20 s before it ›</button
            >
          {/snippet}
        </Row>
        {#if isOpen(death)}
          <StatTable
            columns={COLUMNS}
            nameHead={death.heals === undefined ? 'Last hits' : 'Last hits and heals'}
            barHead="Health after"
            label={`${unitLabel(death.name).name}: last hits and heals`}
          >
            {#each card.events as event, i (`${event.at_ms}-${i}`)}
              {#if event.kind === 'heal'}
                {@const heal = event.heal}
                <StatRow
                  figures={healFigures(heal)}
                  phone={`+${formatAmount(heal.amount)}`}
                  testid="death-heal"
                >
                  {#snippet icon()}<SpellIcon url={spellIcon(heal.spell_id)} />{/snippet}
                  {#snippet name()}<b class="text-kill">{heal.spell_name}</b>{/snippet}
                  {#snippet sub()}{formatDurationPrecise(heal.at_ms)} · from {splitUnitName(heal.source_name)
                      .name}{/snippet}
                  {#snippet meter()}{#if card.pcts[i] !== null}<Meter
                        value={(card.pcts[i] ?? 0) / 100}
                        tone="kill"
                      />{/if}{/snippet}
                </StatRow>
              {:else}
                {@const hit = event.hit}
                {@const lethal = isKillingBlow(death, hit)}
                {@const pct = lethal && !lethalHitMissing(death, card) ? 0 : card.pcts[i]}
                <StatRow figures={hitFigures(death, hit, lethal)} phone={`−${formatAmount(landed(hit))}`}>
                  {#snippet icon()}<SpellIcon url={spellIcon(hit.spell_id)} />{/snippet}
                  {#snippet name()}<b class="text-strong"
                      >{hit.spell_name === '' ? 'Melee' : hit.spell_name}</b
                    >{/snippet}
                  {#snippet sub()}{formatDurationPrecise(hit.at_ms)} · from {sourceName(
                      hit.source_guid,
                      hit.source_name,
                    )}{/snippet}
                  {#snippet meter()}{#if pct !== null}<Meter value={pct / 100} tone="kill" />{/if}{/snippet}
                </StatRow>
              {/if}
            {/each}
          </StatTable>
          {#if footLine !== '' || link}
            <p class="text-small text-muted px-4 pt-2 pb-3" data-testid="death-shape">
              {footLine}{#if link}
                <a class="text-gold ml-1" href={link.href} data-testid="death-build-link">{link.label} ›</a
                >{/if}
            </p>
          {/if}
        {/if}
      </div>
    {/each}
  </Plate>
</Section>
