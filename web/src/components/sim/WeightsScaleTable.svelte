<!-- web/src/components/sim/WeightsScaleTable.svelte -->
<!-- A weights run read the way the published weights are (owner, 2026-09-30): the
     SimulationCraft scale-factor table the BiS page's StatWeightsPanel.astro draws -- Stat,
     bar, Scale (top significant per-point stat = 1.00), Per point (DPS), Error -- with haste
     per 1% in the caption, never a row. The live tool (StatWeights.svelte) and a saved run
     (SavedWeights.svelte) both mount this, over lib/bis/scale-view.ts's one view-model. -->
<script lang="ts">
  import { bisCopy } from '../../lib/bis/copy';
  import type { ScaleView } from '../../lib/bis/scale-view';

  let { view }: { view: ScaleView } = $props();

  const caption = $derived(
    view.hasteCaption === undefined ? view.noteLine : `${view.noteLine} ${view.hasteCaption}`,
  );
</script>

<div class="flex flex-col gap-2" data-testid="sim-weights-table">
  <p class="text-muted text-[12px]" data-testid="sim-weights-scale-caption">{caption}</p>
  <div class="weights-table" role="table">
    <div class="weights-row border-line-soft border-b pb-1.5" role="row">
      <span class="quiet-label" role="columnheader">Stat</span>
      <span aria-hidden="true"></span>
      <span class="quiet-label text-right" role="columnheader">Scale</span>
      <span class="quiet-label text-right" role="columnheader">Per point</span>
      <span class="quiet-label text-right" role="columnheader">Error</span>
    </div>
    {#each view.rows as row (row.stat)}
      <div
        class="weights-row border-line-soft min-h-11 border-b py-2 {row.significant ? '' : 'opacity-50'}"
        role="row"
        title={row.significant ? row.ratingFactorTitle : `${row.label} not significant`}
        data-testid={`sim-weight-${row.stat}`}
      >
        <span class="text-strong text-[13px] leading-tight" role="cell">{row.label}</span>
        <span class="bg-line-soft relative block h-2 overflow-hidden rounded-full" aria-hidden="true">
          {#if row.significant}
            <span class="bg-gold block h-full rounded-full" style={`width:${row.barPercent}%`}></span>
          {/if}
        </span>
        {#if row.significant}
          <span class="numeric text-strong text-right text-[13px]" role="cell"
            >{row.scaleFactor.toFixed(2)}</span
          >
          <span class="numeric text-muted text-right text-[11px]" role="cell">
            {row.dpsPerPoint !== undefined ? `${row.dpsPerPoint.toFixed(3)} ${row.perPointUnit}` : ''}
          </span>
          <span class="numeric text-muted text-right text-[11px]" role="cell"
            >±{row.scaleError.toFixed(2)}</span
          >
        {:else}
          <span class="text-muted col-[3/-1] text-right text-[12px] whitespace-nowrap" role="cell">
            {bisCopy.weightsNotSignificant}
          </span>
        {/if}
      </div>
    {/each}
  </div>
</div>

<style>
  .weights-table {
    display: flex;
    flex-direction: column;
  }
  .weights-row {
    display: grid;
    grid-template-columns: minmax(104px, 1.2fr) minmax(40px, 3fr) 52px 80px 52px;
    gap: 8px;
    align-items: center;
  }
  @media (max-width: 767px) {
    .weights-row {
      grid-template-columns: 90px minmax(24px, 1fr) 44px 72px 44px;
      gap: 6px;
    }
  }
</style>
