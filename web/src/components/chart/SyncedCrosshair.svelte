<!-- web/src/components/chart/SyncedCrosshair.svelte -->
<!-- The shared crosshair on a lane or SVG view that is not a uPlot chart (controls and charts
     spec §B6.3, §B10 rows 4 to 6): a 1 px dashed line at the page cursor's second, inside a
     positioned parent that spans the window. Only the chart under the pointer shows a tooltip;
     this view draws the line. -->
<script lang="ts">
  import { cursorMsFor } from '../../lib/chart/cursor.svelte';
  import type { TimeWindow } from '../../lib/report/window';

  let { syncKey, window: shown }: { syncKey: string; window: TimeWindow } = $props();

  const ms = $derived(cursorMsFor(syncKey));
  const left = $derived(
    ms === null || ms < shown.startMs || ms > shown.endMs
      ? null
      : ((ms - shown.startMs) / Math.max(shown.endMs - shown.startMs, 1)) * 100,
  );
</script>

{#if left !== null}
  <span class="crosshair" style:left="{left}%" data-testid="synced-crosshair" aria-hidden="true"></span>
{/if}

<style>
  .crosshair {
    position: absolute;
    top: 0;
    bottom: 0;
    width: 0;
    border-left: 1px dashed color-mix(in srgb, var(--color-text) 70%, transparent);
    pointer-events: none;
    z-index: 1;
  }
</style>
