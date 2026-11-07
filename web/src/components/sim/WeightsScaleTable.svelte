<!-- web/src/components/sim/WeightsScaleTable.svelte -->
<!-- A weights run read the way the published weights are (owner, 2026-09-30): the
     SimulationCraft scale-factor table the BiS page's StatWeightsPanel.astro draws -- Stat,
     bar, Scale (top significant per-point stat = 1.00), Per point (DPS), Error -- with haste
     per 1% in the caption, never a row. The live tool (StatWeights.svelte) and a saved run
     (SavedWeights.svelte) both mount this, over lib/bis/scale-view.ts's one view-model.
     Drawn as an open ui/StatTable: the rating stats' "14 Crit rating = 1% Crit" is the row's
     visible sub line rather than a hover-only title; a stat the engine could not separate
     from noise reads muted, with no bar. -->
<script lang="ts">
  import { bisCopy } from '../../lib/bis/copy';
  import type { ScaleRow, ScaleView } from '../../lib/bis/scale-view';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import Meter from '../ui/Meter.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';

  let { view }: { view: ScaleView } = $props();

  const caption = $derived(
    view.hasteCaption === undefined ? view.noteLine : `${view.noteLine} ${view.hasteCaption}`,
  );

  const COLUMNS: readonly StatColumn[] = [
    { id: 'scale', label: 'Scale' },
    { id: 'per-point', label: 'Per point' },
    { id: 'error', label: 'Error' },
  ];

  function figuresOf(row: ScaleRow): StatFigure[] {
    if (!row.significant) {
      return [
        { value: bisCopy.weightsNotSignificant, muted: true },
        { value: '', muted: true },
        { value: '', muted: true },
      ];
    }
    return [
      { value: row.scaleFactor.toFixed(2) },
      {
        value: row.dpsPerPoint === undefined ? '' : `${row.dpsPerPoint.toFixed(3)} ${row.perPointUnit}`,
        muted: true,
      },
      { value: `±${row.scaleError.toFixed(2)}`, muted: true },
    ];
  }
</script>

<div class="flex flex-col gap-2" data-testid="sim-weights-table">
  <p class="text-muted text-[13px]" data-testid="sim-weights-scale-caption">{caption}</p>
  <StatTable columns={COLUMNS} nameHead="Stat" open label="Stat weights">
    {#each view.rows as row (row.stat)}
      <StatRow figures={figuresOf(row)} testid={`sim-weight-${row.stat}`}>
        {#snippet name()}
          <span class={row.significant ? '' : 'text-muted'}>{row.label}</span>
        {/snippet}
        {#snippet sub()}{row.ratingFactorTitle ?? ''}{/snippet}
        {#snippet meter()}
          {#if row.significant}<Meter value={row.barPercent / 100} />{/if}
        {/snippet}
      </StatRow>
    {/each}
  </StatTable>
</div>
