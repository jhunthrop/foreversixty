<!-- web/src/components/ui/PlateColumns.svelte -->
<!-- The one Plate split into labelled columns by hairlines (guides spec 2026-11-08 §3A): the
     PlateColumns side by side from 1024, stacked below. Each child is a PlateColumn. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import Plate from './Plate.svelte';

  let {
    count,
    label,
    testid,
    children,
  }: {
    /** How many columns the plate holds (1 to 3). */
    count: number;
    /** The plate's accessible name. */
    label: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<Plate {label} {testid}>
  <div class="ui-pcols" style:--pcols={count}>{@render children()}</div>
</Plate>

<style>
  .ui-pcols {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
  }
  .ui-pcols > :global(.ui-pcol + .ui-pcol) {
    border-top: 1px solid var(--color-line-soft);
  }
  @media (min-width: 1024px) {
    .ui-pcols {
      grid-template-columns: repeat(var(--pcols), minmax(0, 1fr));
    }
    .ui-pcols > :global(.ui-pcol + .ui-pcol) {
      border-top: 0;
      border-left: 1px solid var(--color-line-soft);
    }
  }
</style>
