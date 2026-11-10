<!-- web/src/components/ui/PlateColumn.svelte -->
<!-- One column of a PlateColumns plate: a 12 px caps label, an optional Small meta at its right,
     its body, and a foot (the column's buttons or link) at the bottom. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    label,
    meta,
    foot,
    testid,
    children,
  }: {
    label: string;
    meta?: string;
    foot?: Snippet;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div class="ui-pcol" data-testid={testid}>
  <div class="ui-pcol-head">
    <h2 class="text-label text-muted m-0 font-bold">{label}</h2>
    {#if meta}<span class="text-small text-muted ml-auto text-right">{meta}</span>{/if}
  </div>
  {@render children()}
  {#if foot}<div class="ui-pcol-foot">{@render foot()}</div>{/if}
</div>

<style>
  .ui-pcol {
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding: 16px 20px 20px;
    min-width: 0;
  }
  .ui-pcol-head {
    display: flex;
    align-items: baseline;
    gap: 12px;
  }
  /* The plate's tables are compact: 40 px rows (guides spec §4.3). */
  .ui-pcol :global(.ui-st-r) {
    min-height: 40px;
    padding-block: 6px;
  }
  /* A plate column is narrow even on a wide screen: the name keeps its share of the row. */
  .ui-pcol :global(.ui-st) {
    --st-name: minmax(0, 2fr);
    --st-bar: minmax(0, 1fr);
  }
  /* A plate table's lead cell is its rank, not a 28 px icon; slim rows breathe a pixel more. */
  .ui-pcol :global(.ui-st-r .ui-st-icon) {
    height: auto;
  }
  .ui-pcol :global(.ui-row) {
    padding-block: 7px;
  }
  /* Below 1024 a plate's table keeps its bar on the row (a 96 px track), not under the name. */
  @media (max-width: 1023px) {
    .ui-pcol :global(.ui-st .ui-st-c) {
      grid-template-columns: var(--st-icon) minmax(0, 1fr) minmax(0, 96px) auto;
    }
    .ui-pcol :global(.ui-st[data-open='true'] .ui-st-r .ui-st-bar) {
      display: block;
      grid-column: auto;
      padding-top: 0;
    }
    /* The bar stays on the row here, so the figure keeps its own track after it. */
    .ui-pcol :global(.ui-st[data-open='true'] .ui-st-r .ui-st-f[data-primary='true']) {
      grid-row: auto;
      grid-column: auto;
    }
  }
  .ui-pcol-foot {
    margin-top: auto;
    display: flex;
    gap: 12px;
    flex-wrap: wrap;
    align-items: center;
    padding-top: 4px;
  }
</style>
