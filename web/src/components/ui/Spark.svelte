<!-- web/src/components/ui/Spark.svelte -->
<!-- A one-line series in a cell or a Figure (logs spec 2026-11-07 section 3A, decision P1): a 24 px
     line over the fight (a resource bar, a rating trend), or with `marks` a tick strip (a cast
     sequence). Decorative: the figures beside it carry the numbers. -->
<script lang="ts">
  import { sparkPoints, sparkTickX } from '../../lib/ui/spark';

  let {
    values = [],
    marks,
    top,
    color = 'var(--color-gold)',
  }: {
    values?: readonly number[];
    /** Tick positions 0 to 1; given, the spark is a tick strip. */
    marks?: readonly number[];
    /** The series' ceiling (100 for a rating); its own largest value by default. */
    top?: number;
    color?: string;
  } = $props();
</script>

<svg class="block h-6 w-full" viewBox="0 0 1000 24" preserveAspectRatio="none" aria-hidden="true">
  {#if marks !== undefined}
    <rect x="0" y="11" width="1000" height="2" fill="var(--color-line-soft)" />
    {#each marks as mark, index (index)}<rect
        x={sparkTickX(mark)}
        y="4"
        width="4"
        height="16"
        rx="1"
        fill={color}
      />{/each}
  {:else}
    <polyline
      fill="none"
      stroke={color}
      stroke-width="2"
      vector-effect="non-scaling-stroke"
      points={sparkPoints(values, top)}
    />
  {/if}
</svg>
