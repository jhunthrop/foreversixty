<!-- web/src/components/chart/TimeSeriesChart.svelte -->
<!-- The shared report chart (controls and charts spec §B4): every report chart is this. It is the
     frame that stays in report-island.js: it reserves the loaded chart's exact height, shows the
     loading skeleton (§B7: skeleton plot and brush, no spinner) and the error state with Try
     again, and hands every prop to TimeSeriesPlot, the lazy chunk that holds uPlot. -->
<script lang="ts">
  import type { ComponentProps } from 'svelte';
  import { chartCopy } from '../../lib/chart/copy';
  import {
    brushHeight,
    CHART_AXIS_PX,
    CHART_EVENT_LINE_PX,
    CHART_LANE_PX,
    gutterFor,
    plotHeightFor,
  } from '../../lib/chart/layout';
  import { loadChartChunk, type PlotComponent } from '../../lib/chart/uplot-loader';
  import LoadError from '../ui/LoadError.svelte';

  type Props = ComponentProps<PlotComponent>;
  let props: Props = $props();

  let Plot = $state<PlotComponent | null>(null);
  let failed = $state(false);
  function load(): void {
    failed = false;
    loadChartChunk().then(
      (component) => (Plot = component),
      () => (failed = true),
    );
  }
  if (typeof window !== 'undefined') load();

  let viewport = $state(typeof window === 'undefined' ? 1440 : window.innerWidth);
  let coarse = $state(false);
  $effect(() => {
    const query = window.matchMedia('(pointer: coarse), (hover: none)');
    const read = (): void => {
      viewport = window.innerWidth;
      coarse = query.matches;
    };
    read();
    window.addEventListener('resize', read);
    query.addEventListener('change', read);
    return () => {
      window.removeEventListener('resize', read);
      query.removeEventListener('change', read);
    };
  });
  const compact = $derived(props.compact === true);
  const lane = $derived(!compact && (props.events ?? []).some((event) => event.kind !== 'phase'));
  const plotHeight = $derived(plotHeightFor(viewport, props.plotHeights) + CHART_AXIS_PX);
  const gutter = $derived(gutterFor(viewport));
</script>

{#if Plot !== null}
  <Plot {...props} />
{:else}
  <div
    class="flex flex-col"
    data-testid={props.testid ?? 'chart-frame'}
    data-state={failed ? 'error' : 'loading'}
    data-compact={compact ? 'true' : undefined}
  >
    {#if lane}<div style:height="{CHART_LANE_PX}px" style:margin-top="4px"></div>{/if}
    <div
      class="relative"
      style:height="{plotHeight}px"
      role="group"
      aria-roledescription="chart"
      aria-label={chartCopy.loading}
      aria-busy={failed ? undefined : 'true'}
      data-testid="chart-plot"
    >
      {#if failed}
        <div class="absolute inset-0 flex items-center justify-center">
          <LoadError message={chartCopy.error} onRetry={load} testid="chart-error" />
        </div>
      {:else}
        <span
          class="skeleton-block absolute"
          style:left="{gutter}px"
          style:right="12px"
          style:top="8px"
          style:bottom="{CHART_AXIS_PX}px"
          aria-hidden="true"
        ></span>
      {/if}
    </div>
    {#if !compact && coarse}<div style:height="{CHART_EVENT_LINE_PX}px"></div>{/if}
    {#if !compact}
      <div style:height="{brushHeight(coarse)}px" style:padding="6px 12px 24px {gutter}px">
        {#if !failed}<span class="skeleton-block block h-full rounded-[6px]" aria-hidden="true"></span>{/if}
      </div>
    {/if}
  </div>
{/if}
