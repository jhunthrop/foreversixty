<!-- web/src/components/ui/PlateColumns.svelte -->
<!-- The one Plate split into labelled columns by hairlines (guides spec 2026-11-08 §3A): the
     PlateColumns side by side from 1024, stacked below. Each child is a PlateColumn. `frame` is
     the Plate's (a featured plate: the guide's At a glance is `framed`). `firstWide`: from 1024 to
     1279 a three-column plate gives its first column the whole first row and the other two share
     the second (the guide's three talent trees do not fit a third of the plate there). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import Plate from './Plate.svelte';
  import type { PanelFrame } from '../../lib/ui/primitives';

  let {
    count,
    label,
    frame = 'plain',
    firstWide = false,
    testid,
    children,
  }: {
    /** How many columns the plate holds (1 to 3). */
    count: number;
    /** The plate's accessible name. */
    label: string;
    /** The Plate's frame (ui/Plate). */
    frame?: PanelFrame;
    /** From 1024 to 1279, a three-column plate's first column takes the first row. */
    firstWide?: boolean;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<Plate {label} {frame} {testid}>
  <div class="ui-pcols" style:--pcols={count} data-first-wide={firstWide && count === 3 ? '' : undefined}>
    {@render children()}
  </div>
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
  @media (min-width: 1024px) and (max-width: 1279px) {
    .ui-pcols[data-first-wide] {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
    .ui-pcols[data-first-wide] > :global(.ui-pcol:first-child) {
      grid-column: 1 / -1;
    }
    .ui-pcols[data-first-wide] > :global(.ui-pcol:nth-child(2)) {
      border-left: 0;
      border-top: 1px solid var(--color-line-soft);
    }
    .ui-pcols[data-first-wide] > :global(.ui-pcol:nth-child(3)) {
      border-top: 1px solid var(--color-line-soft);
    }
  }
</style>
