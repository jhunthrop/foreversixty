<!-- web/src/components/sim/CompareView.svelte -->
<!-- The sim beside the fight. The words come first and the tables second, because most
     people will read "Bloodthirst cast 33 times, the sim expects 41" and stop -- that is
     the answer, and the tables are the working.
     Nothing here is coloured. A red-and-green diff would be a fourth colour vocabulary on
     a page that already has class colours, rarity colours and percentile tokens, and the
     design system does not allow colour to carry meaning on its own anyway.
     The ability table is the view's one Plate (a StatTable, sim spec 2026-11-06 §4.7); the
     aura uptimes an open StatTable under it. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { pct as pctOf, type Comparison } from '../../lib/sim/compare';
  import type { StatColumn } from '../../lib/ui/stat-table';
  import Plate from '../ui/Plate.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';

  let {
    comparison,
    simDuration,
    actualDuration,
  }: { comparison: Comparison; simDuration: number; actualDuration: number } = $props();

  // The rounding rule itself (whole percent, own duration) lives once, in compare.ts, so
  // this table's numbers and compareSummaries' own uptime lines can never round differently.
  const pct = (ms: number, of: number): string => `${pctOf(ms, of)}%`;
  const amount = (value: number): string => Math.round(value).toLocaleString('en-US');

  const ABILITY_COLUMNS: readonly StatColumn[] = [
    { id: 'actual-casts', label: simCopy.compareActualCasts },
    { id: 'sim-casts', label: simCopy.compareSimCasts },
    { id: 'actual-damage', label: simCopy.compareActualDamage },
    { id: 'sim-damage', label: simCopy.compareSimDamage },
  ];
  const AURA_COLUMNS: readonly StatColumn[] = [
    { id: 'actual', label: simCopy.compareActual },
    { id: 'simulated', label: simCopy.compareSimulated },
  ];
</script>

<div class="flex flex-col gap-4" data-testid="sim-compare">
  <p class="text-strong px-[18px] text-[15px] md:px-0" role="status" data-testid="compare-headline">
    {comparison.headline}
  </p>

  {#if comparison.lines.length > 0}
    <ul class="flex flex-col gap-1 px-[18px] md:px-0" data-testid="compare-lines">
      <!-- Keyed by position, never by text: a line is prose, not an identity, and two rows
           can word the same gap identically. A repeated {#each} key is a Svelte 5 runtime
           error, and the list is rebuilt whole on every new result anyway. -->
      {#each comparison.lines as line, index (index)}
        <li class="text-[14px]">{line}</li>
      {/each}
    </ul>
  {/if}

  <div class="px-[18px] md:px-0">
    <Plate testid="compare-abilities">
      <StatTable columns={ABILITY_COLUMNS} nameHead={simCopy.compareAbility} bar={false}>
        <!-- Keyed and tested by id, never by name: a resolved name can hold a space ("Heroic
             Strike"), a colon ("spell:25286") or a slash, none of which belong in a test id or
             in a Svelte key. -->
        {#each comparison.abilities as row (row.spellId)}
          <StatRow
            bar={false}
            figures={[
              { value: String(row.actualCasts) },
              { value: String(row.simCasts), muted: true },
              { value: amount(row.actualDamage) },
              { value: amount(row.simDamage), muted: true },
            ]}
            phone={`${row.simCasts} ${simCopy.compareSimCasts.toLowerCase()}`}
            testid={`compare-ability-${row.spellId}`}
          >
            {#snippet name()}<span class="block truncate">{row.name}</span>{/snippet}
            {#snippet sub()}<span class="ui-st-phone"
                >{amount(row.actualDamage)}
                {simCopy.compareActual} · {amount(row.simDamage)}
                {simCopy.compareSimulated}</span
              >{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    </Plate>
  </div>

  <div class="px-[18px] md:px-0">
    <!-- Keyed and tested by id, never by name, for the same reason the ability table is:
         two real buffs at different spell ids (a rank mismatch between the sim's loadout
         and the logged fight's) can resolve to the same display name, and a repeated
         {#each} key is a Svelte 5 runtime error. -->
    <StatTable columns={AURA_COLUMNS} nameHead={simCopy.compareBuff} bar={false} open testid="compare-auras">
      {#each comparison.auras as row (row.spellId)}
        <StatRow
          bar={false}
          figures={[
            { value: pct(row.actualUptimeMs, actualDuration) },
            { value: pct(row.simUptimeMs, simDuration), muted: true },
          ]}
          phone={`${pct(row.simUptimeMs, simDuration)} ${simCopy.compareSimulated.toLowerCase()}`}
          testid={`compare-aura-${row.spellId}`}
        >
          {#snippet name()}<span class="block truncate">{row.name}</span>{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </div>

  <p class="text-muted px-[18px] text-[13px] leading-5 md:px-0" data-testid="compare-footnote">
    {simCopy.compareFootnote}
  </p>
</div>
