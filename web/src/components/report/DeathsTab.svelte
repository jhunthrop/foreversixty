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
  import { plannerLinkFor } from '../../lib/report/planner-link';
  import type { CombatantRow, DamageRef, Death, HealRef } from '../../lib/report/types';
  import { deathWindow, type TimeWindow } from '../../lib/report/window';

  let {
    deaths,
    durationMs,
    combatants,
    classOf,
    dataBuild,
    treeSizesFor,
    onWindow,
  }: {
    deaths: Death[];
    durationMs: number;
    combatants: CombatantRow[];
    classOf: Map<string, string>;
    dataBuild: string;
    /** Talents per tree for a class, or [] when the planner has no data for it. */
    treeSizesFor: (className: string | undefined) => number[];
    onWindow: (window: TimeWindow) => void;
  } = $props();

  const ordered = $derived([...deaths].sort((a, b) => a.at_ms - b.at_ms));

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
  function lethalHitMissing(death: Death): boolean {
    const blow = death.killing_blow;
    return blow !== undefined && blow.max_hp !== undefined && blow.max_hp > 0 && (blow.hp_after ?? 0) > 0;
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

  function healthPct(hit: Death['last'][number]): number | null {
    if (!hit.max_hp) return null;
    return Math.max(0, Math.min(100, ((hit.hp_after ?? 0) / hit.max_hp) * 100));
  }

  /**
   * "Burst" or "bleed": the one word a healer wants first. The last hits are the engine's
   * final ten damage events; if the player went from above 80% to dead inside three
   * seconds, no heal was going to land in time, and that is a different conversation
   * from a health bar that drained over fifteen seconds while nobody was healing it.
   */
  function shape(death: Death): string | null {
    const hits = death.last;
    if (hits.length < 2) return null;
    const first = hits[0];
    // The health the run-up started from is the highest the recap saw, not the first
    // row's: heals between hits can lift it, and "from 4%" when they were at 47% two
    // hits later is the wrong story.
    const firstPct = hits.reduce<number | null>((highest, hit) => {
      const pct = healthPct(hit);
      return pct === null ? highest : Math.max(highest ?? 0, pct);
    }, null);
    if (firstPct === null) return null;
    const spanMs = death.at_ms - first.at_ms;
    const total = hits.reduce((sum, hit) => sum + hit.amount, 0);
    const spanText = formatDuration(spanMs);
    const healed = healingBefore(death);
    const healedText =
      death.heals === undefined
        ? ''
        : healed > 0
          ? ` ${formatAmount(healed)} of healing landed in the same span.`
          : ' No healing landed on them in that span.';
    if (spanMs <= 3000 && firstPct >= 60)
      return `Burst: ${formatAmount(total)} in ${spanText}, from ${Math.round(firstPct)}% health.${healedText}`;
    return `${formatAmount(total)} over ${spanText}, from ${Math.round(firstPct)}% health.${healedText}`;
  }
  function linkFor(death: Death): ReturnType<typeof plannerLinkFor> {
    const combatant = combatants.find((row) => row.guid === death.guid);
    if (combatant === undefined) return null;
    const className = death.class ?? classOf.get(death.guid);
    return plannerLinkFor({ dataBuild, className, treeSizes: treeSizesFor(className), combatant });
  }
</script>

{#if ordered.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">Nobody died in this window.</p>
{:else}
  <ul class="flex flex-col gap-4" data-testid="deaths-tab">
    <!-- A death's last hits can repeat a spell in one millisecond (a DoT tick and its
         crit, a cleave), so the row key carries its index too; a bare timestamp-and-spell
         key threw on the duplicate and blanked the whole tab. -->
    {#each ordered as death (`${death.guid}-${death.at_ms}`)}
      {@const link = linkFor(death)}
      <li
        class="border-line rounded-panel bg-raised flex flex-col gap-3 border p-3"
        data-testid={`death-${death.guid}`}
      >
        <div class="flex flex-wrap items-baseline gap-x-3 gap-y-1">
          <span
            class="text-[15px] font-semibold"
            style={`color: ${classColorVar(death.class ?? classOf.get(death.guid))}`}
          >
            {splitUnitName(death.name).name}
          </span>
          {#if death.label}
            <span class="text-muted text-[13px]" data-testid="death-label">{death.label}</span>
          {:else}
            <span class="text-muted tabular font-mono text-[13px]">{formatDuration(death.at_ms)}</span>
          {/if}
          {#if death.killing_blow}
            <span class="text-[13px]">
              killed by {sourceName(death.killing_blow.source_guid, death.killing_blow.source_name)} ·
              {death.killing_blow.spell_name === '' ? 'Melee' : death.killing_blow.spell_name} ·
              <span class="tabular font-mono">{formatAmount(death.killing_blow.amount)}</span>
              {#if death.killing_blow.overkill}
                <span class="text-muted">
                  (<span class="tabular font-mono">{formatAmount(death.killing_blow.overkill)}</span> overkill)
                </span>
              {/if}
            </span>
          {/if}
          {#if death.release_ms}
            <span class="text-muted text-[13px]">
              released after
              <span class="tabular font-mono">{formatDuration(death.release_ms - death.at_ms)}</span>
            </span>
          {/if}
        </div>

        {#if shape(death)}
          <p class="text-[13px]" data-testid="death-shape">{shape(death)}</p>
        {/if}
        {#if death.killing_blow && lethalHitMissing(death)}
          <p class="text-muted text-[13px]" data-testid="death-unlogged">
            The log shows no lethal hit: the last recorded hit left them at
            {#if healthPct(death.killing_blow) !== null}
              <span class="tabular font-mono">{Math.round(healthPct(death.killing_blow) ?? 0)}%</span>,
            {:else}
              unknown health,
            {/if}
            and they died
            <span class="tabular font-mono">{beforeDeath(death, death.killing_blow.at_ms).slice(1)}</span> later.
          </p>
        {/if}

        <div class="overflow-x-auto">
          <table class="w-full text-[13px]">
            <caption class="label text-muted text-left">
              {death.heals === undefined ? 'Last hits' : 'Last hits and heals'}
            </caption>
            <thead>
              <tr class="text-muted label">
                <th class="py-1 pr-3 text-left font-bold" title="Seconds before the death">Before</th>
                <th class="py-1 pr-3 text-left font-bold">Ability</th>
                <th class="py-1 pr-3 text-left font-bold">From</th>
                <th class="py-1 pr-3 text-right font-bold">Amount</th>
                <th class="py-1 text-left font-bold" title="Health left after the hit">Health after</th>
              </tr>
            </thead>
            <tbody>
              {#each lastEvents(death) as event, i (`${event.at_ms}-${i}`)}
                {#if event.kind === 'heal'}
                  {@const heal = event.heal}
                  <tr class="border-line-soft border-b" data-testid="death-heal">
                    <td
                      class="text-muted tabular py-1 pr-3 font-mono"
                      title={formatDurationPrecise(heal.at_ms)}>{beforeDeath(death, heal.at_ms)}</td
                    >
                    <td class="text-kill py-1 pr-3">{heal.spell_name}</td>
                    <td class="text-muted truncate py-1 pr-3">{splitUnitName(heal.source_name).name}</td>
                    <td class="text-kill tabular py-1 pr-3 text-right font-mono"
                      >+{formatAmount(heal.amount - (heal.overheal ?? 0))}{#if heal.overheal}
                        <span class="text-muted text-[11px]" title="Overhealing">
                          ({formatAmount(heal.overheal)} over)</span
                        >{/if}</td
                    >
                    <td class="py-1"></td>
                  </tr>
                {:else}
                  {@const hit = event.hit}
                  {@const lethal = isKillingBlow(death, hit)}
                  {@const pct = lethal ? 0 : healthPct(hit)}
                  <tr class="border-line-soft border-b">
                    <td
                      class="text-muted tabular py-1 pr-3 font-mono"
                      title={formatDurationPrecise(hit.at_ms)}>{beforeDeath(death, hit.at_ms)}</td
                    >
                    <td class="py-1 pr-3">{hit.spell_name === '' ? 'Melee' : hit.spell_name}</td>
                    <td class="text-muted truncate py-1 pr-3"
                      >{sourceName(hit.source_guid, hit.source_name)}</td
                    >
                    <td class="tabular py-1 pr-3 text-right font-mono">{formatAmount(hit.amount)}</td>
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

        {#if death.auras_held.length > 0 || death.auras_lost.length > 0}
          <p class="text-[13px]">
            {#if death.auras_held.length > 0}
              <span class="label text-muted">Up</span>
              {death.auras_held.map((aura) => aura.name).join(', ')}
            {/if}
            {#if death.auras_lost.length > 0}
              <span class="label text-muted ml-3">Just lost</span>
              {death.auras_lost.map((aura) => aura.name).join(', ')}
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
      </li>
    {/each}
  </ul>
{/if}
