<!-- web/src/components/sim/SavedCombos.svelte -->
<!-- A saved Top Gear, talent compare or Droptimizer result at /sim/<id>: read-only, so the
     ranked table and the per-slot summary without the "open in planner" and "copy to
     addon" buttons -- those act on the reader's own character, and a saved page is not
     theirs. "Run this yourself" beside the header (SavedSim.svelte's own button) is the
     way back to a page that is. Task 20.

     The item file is fetched here for icons and rarity colours only -- every name on the
     page comes off the result (contract 10.1 A6), so a saved sim whose class file 404s
     still reads correctly, with slot names and ids in place of icons.

     Sim spec 2026-11-06 §4.5 grammar: the ranked table is the page's one Plate (a StatTable, the
     rank in the lead column, the gain the primary figure), the per-slot summary an open Section
     of Rows under it. -->
<script lang="ts">
  import { loadItems } from '../../lib/planner/load';
  import type { Item } from '../../lib/planner/types';
  import { slotLabelForClass } from '../../lib/bis/slot-display-labels';
  import type { BulkResult } from '../../lib/sim/bulk-types';
  import { closestOverlappingPair, comboRows, deltaLabel, rowName, slotSummary } from '../../lib/sim/combos';
  import { bulkCopy } from '../../lib/sim/copy';
  import { confidenceBand, formatMargin } from '../../lib/sim/estimate';
  import { poolQualityCopy } from '../../lib/sim/pool-quality-copy';
  import type { StatColumn } from '../../lib/ui/stat-table';
  import Plate from '../ui/Plate.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import SubstitutionChips from './tools/SubstitutionChips.svelte';

  import { grouped } from '../../lib/format/number';
  let { result, treeVersion }: { result: BulkResult; treeVersion: string } = $props();

  let items = $state<Map<number, Item>>(new Map());

  // One load, for this component's one (unchanging) result -- not an `$effect`, the same
  // reason SavedSim.svelte's own item-file load is a plain call rather than one.
  void (async () => {
    try {
      const file = await loadItems(treeVersion, result.request.character.class);
      items = new Map(file.items.map((item) => [item.id, item]));
    } catch {
      // The chips and the slot summary render off the result's own names regardless
      // (contract 10.1 A6); a failed fetch only costs the icon and the rarity colour.
      items = new Map();
    }
  })();

  const rows = $derived(comboRows(result));
  const summary = $derived(slotSummary(result));
  /** newcomer round 4 / dps D25-pattern: the same overbroad warning this saved read-only
   *  view shared with the live ComboResults.svelte (combos.ts's `closestOverlappingPair`). */
  const overlap = $derived(closestOverlappingPair(rows));
  const isTalentsMode = $derived(result.request.bulk?.mode === 'talents');
  const equippedFigure = $derived(grouped(result.equipped.mean));
  const equippedBand = $derived(formatMargin(confidenceBand(result.equipped)));
  const classSlug = $derived(result.request.character.class);

  /** The gain leads (the primary figure, kept on a phone); the DPS and the share follow. */
  const COLUMNS: readonly StatColumn[] = [
    { id: 'gain', label: bulkCopy.resultsDelta },
    { id: 'dps', label: bulkCopy.resultsDps },
    { id: 'percent', label: bulkCopy.resultsPercent },
  ];
</script>

{#snippet withinErrorLine()}{bulkCopy.withinError}{/snippet}

<div class="flex flex-col gap-[var(--section-gap)]" data-testid="sim-combos">
  <div class="flex flex-col gap-3">
    <p class="text-strong text-[14px] leading-5 tabular-nums" data-testid="sim-equipped-line">
      {bulkCopy.resultsEquipped}: {equippedFigure} ± {equippedBand}
      <span class="text-muted">{bulkCopy.ranAtStages(result.stages)}</span>
    </p>
    {#if isTalentsMode}
      <p class="text-muted text-[13px] leading-5" data-testid="sim-talents-gear-locked">
        {poolQualityCopy.talentsGearLockedNote}
      </p>
    {/if}
    {#if rows.length === 0}
      <p class="text-muted text-[14px] leading-5">{bulkCopy.noGain}</p>
    {:else}
      <Plate label={bulkCopy.resultsChange}>
        <StatTable
          columns={COLUMNS}
          nameHead={bulkCopy.resultsChange}
          bar={false}
          label={bulkCopy.resultsChange}
        >
          {#each rows as row (row.combo.substitutions
            .map((sub) => `${sub.kind}:${sub.slot ?? ''}:${sub.item_id ?? sub.name ?? ''}`)
            .join('|'))}
            <StatRow
              bar={false}
              figures={[
                { value: deltaLabel(row.combo.delta) },
                { value: grouped(row.combo.dps.mean) },
                { value: `${row.percent.toFixed(1)}%`, muted: true },
              ]}
              phone={grouped(row.combo.dps.mean)}
              sub={row.withinError ? withinErrorLine : undefined}
              testid="sim-combo-row"
            >
              {#snippet icon()}<span class="text-muted text-[13px] tabular-nums">{row.rank}</span>{/snippet}
              {#snippet name()}<SubstitutionChips
                  substitutions={row.combo.substitutions}
                  {items}
                  {treeVersion}
                  {classSlug}
                />{/snippet}
            </StatRow>
          {/each}
        </StatTable>
      </Plate>
      {#if overlap !== null}
        <p class="text-muted text-[13px] leading-5" data-testid="sim-within-error-note">
          {poolQualityCopy.withinErrorNoteNaming(rowName(overlap.a), rowName(overlap.b))}
        </p>
      {/if}
    {/if}
  </div>

  {#if summary.length > 0}
    <Section title={bulkCopy.slotSummary} meta={bulkCopy.slotSummaryNote} panel testid="sim-slot-summary">
      {#each summary as row (`${row.slot}-${row.item_id}`)}
        <Row title={row.name} sub={slotLabelForClass(row.slot, classSlug)} density="slim" open>
          {#snippet trail()}<span class="text-gold text-[14px] font-semibold tabular-nums"
              >{row.gain === null ? '–' : `+${grouped(row.gain)}`}</span
            >{/snippet}
        </Row>
      {/each}
    </Section>
  {/if}
</div>
