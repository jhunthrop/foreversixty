<!-- web/src/components/sim/DpsDistribution.svelte -->
<!-- The run's DPS spread (spec 2026-11-04-sim-results 3.6; sim spec 2026-11-06 §4.5): the engine's
     real bins as at most 40 bars, the mean line (its figure is the Mean tile, as the mock's `Histogram`), the two-thirds caption, then six small Figures in
     three columns. A result saved before the engine kept the bins draws a range strip instead, and
     says so: a drawn shape the run did not produce would break the accuracy tenet. No chart
     library: the bars are spans. It draws no frame: the caller's open Section holds it. Error is
     the page's one ± (ruling S2: the 95% band, unlabelled here); Std dev carries no ±. Under the tiles, when the run fought the
     spec's published fight (lib/sim/bis-benchmark.ts), the best-in-slot set's DPS and the gap to it
     with both runs' 95% band, linking to that set on Best in slot. -->
<script lang="ts">
  import { bisPageHref } from '../../lib/bis/hover';
  import type { BenchmarkMatch } from '../../lib/sim/bis-benchmark';
  import { confidenceBand } from '../../lib/sim/estimate';
  import { grouped, oneDecimal, spreadCopy } from '../../lib/sim/result-copy';
  import { spreadOf, type SpreadBar } from '../../lib/sim/spread';
  import type { DpsHistogram, Estimate } from '../../lib/sim/types';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import Figure from '../ui/Figure.svelte';
  import Row from '../ui/Row.svelte';

  let {
    estimate,
    iterationsRun,
    histogram,
    spec = '',
    benchmark = null,
  }: {
    estimate: Estimate;
    iterationsRun: number;
    histogram?: DpsHistogram;
    /** The run's spec key, for the set's Best in slot link. */
    spec?: string;
    benchmark?: BenchmarkMatch | null;
  } = $props();

  const gapLine = $derived(
    benchmark === null
      ? ''
      : spreadCopy.benchmarkGap(oneDecimal(Math.abs(benchmark.gap / benchmark.dps) * 100), benchmark.gap < 0),
  );

  /** The readout's width and the side of the plot past which it flips left. */
  const TOOLTIP_WIDTH = 240;
  const TOOLTIP_FLIP_SHARE = 0.55;

  const spread = $derived(spreadOf(histogram));
  const range = $derived(estimate.max - estimate.min);
  /** Where a DPS figure sits across the plot, 0..100. */
  const pct = (value: number): number => (range <= 0 ? 50 : ((value - estimate.min) / range) * 100);
  /** The histogram's own axis runs over its bins, which the range rounds inside. */
  const barsLow = $derived(spread?.bars[0]?.low ?? estimate.min);
  const barsHigh = $derived(spread?.bars.at(-1)?.high ?? estimate.max);
  const meanAt = $derived(((estimate.mean - barsLow) / Math.max(1, barsHigh - barsLow)) * 100);

  const tiles = $derived([
    { id: 'mean', label: spreadCopy.tiles.mean, value: oneDecimal(estimate.mean) },
    { id: 'error', label: spreadCopy.tiles.error, value: `± ${oneDecimal(confidenceBand(estimate))}` },
    { id: 'stddev', label: spreadCopy.tiles.stddev, value: oneDecimal(estimate.stddev) },
    { id: 'lowest', label: spreadCopy.tiles.lowest, value: grouped(estimate.min) },
    { id: 'highest', label: spreadCopy.tiles.highest, value: grouped(estimate.max) },
    { id: 'runs', label: spreadCopy.tiles.runs, value: grouped(iterationsRun) },
  ]);

  // The readout's panel is the report's own, fetched the first time a bar is pointed at.
  const tooltipLazy = createLazyComponent(() => import('../chart/ChartTooltip.svelte'));

  let plot = $state<HTMLDivElement | undefined>(undefined);
  let active = $state<number | null>(null);
  let focused = $state(false);
  const activeBar = $derived<SpreadBar | null>(active === null ? null : (spread?.bars[active] ?? null));

  function moveTo(index: number): void {
    const last = (spread?.bars.length ?? 1) - 1;
    active = Math.max(0, Math.min(last, index));
  }

  function onKeydown(event: KeyboardEvent): void {
    const last = (spread?.bars.length ?? 1) - 1;
    const from = active ?? -1;
    const next: Record<string, number> = { ArrowRight: from + 1, ArrowLeft: from - 1, Home: 0, End: last };
    const target = next[event.key];
    if (target === undefined) return;
    event.preventDefault();
    moveTo(target);
  }

  const tooltipBox = $derived.by(() => {
    if (plot === undefined || active === null || spread === null) return { x: 0, y: 0, flip: false };
    const rect = plot.getBoundingClientRect();
    const share = (active + 0.5) / spread.bars.length;
    return { x: rect.left + rect.width * share, y: rect.top, flip: share > TOOLTIP_FLIP_SHARE };
  });
</script>

<div class="flex flex-col gap-3" data-testid="sim-distribution">
  <div class="flex min-w-0 flex-col gap-3">
    {#if spread !== null}
      <div>
        <!-- One image with a readout: hover or arrow keys step the bars (spec 3.6). -->
        <!-- svelte-ignore a11y_no_noninteractive_tabindex, a11y_no_noninteractive_element_interactions -->
        <div
          bind:this={plot}
          class="focus-visible:outline-gold relative flex h-[168px] items-end gap-[2px] focus-visible:outline-2 focus-visible:outline-offset-2"
          role="img"
          tabindex="0"
          aria-label={spreadCopy.image(
            grouped(spread.total),
            grouped(estimate.min),
            grouped(estimate.max),
            grouped(estimate.mean),
          )}
          data-testid="sim-spread-plot"
          onfocus={() => {
            focused = true;
            tooltipLazy.load();
          }}
          onpointerenter={tooltipLazy.load}
          onblur={() => {
            focused = false;
            active = null;
          }}
          onkeydown={onKeydown}
          onpointerleave={() => {
            if (!focused) active = null;
          }}
        >
          {#each spread.bars as bar, index (bar.low)}
            <span
              class={`min-w-0 flex-1 rounded-t-[2px] ${index === active ? 'bg-accent' : bar.central ? 'bg-accent/70' : 'bg-strong/20'}`}
              style:height={`${Math.max(1, (bar.count / spread.peak) * 100)}%`}
              data-testid="sim-spread-bar"
              onpointerenter={() => (active = index)}
              aria-hidden="true"
            ></span>
          {/each}
          <span
            class="bg-gold pointer-events-none absolute -top-1.5 bottom-0 w-[2px] -translate-x-1/2"
            style:left={`${meanAt}%`}
            data-testid="sim-distribution-mean"
          ></span>
        </div>
        <div class="text-muted numeric mt-2 flex justify-between text-[13px]">
          <span>{grouped(estimate.min)}</span>
          <span>{spreadCopy.axis}</span>
          <span>{grouped(estimate.max)}</span>
        </div>
      </div>
      {#if tooltipLazy.current}
        <tooltipLazy.current
          open={activeBar !== null}
          x={tooltipBox.x}
          y={tooltipBox.y}
          flip={tooltipBox.flip}
          title={activeBar === null
            ? ''
            : spreadCopy.bar(grouped(activeBar.low), grouped(activeBar.high), grouped(activeBar.count))}
          note=""
          events={[]}
          rows={[]}
          more={0}
          width={TOOLTIP_WIDTH}
        />
      {/if}
    {:else if range > 0}
      <!-- A result saved before the engine kept the bins: the range it did keep. -->
      <div class="flex flex-col gap-2 pt-2" data-testid="sim-spread-range">
        <div class="bg-line-soft relative h-2 rounded-[4px]">
          <span
            class="bg-accent/35 absolute inset-y-0"
            style:left={`${Math.max(0, pct(estimate.mean - estimate.stddev))}%`}
            style:right={`${Math.max(0, 100 - pct(estimate.mean + estimate.stddev))}%`}
          ></span>
          <span
            class="bg-gold absolute -inset-y-1 w-[2px] -translate-x-1/2"
            style:left={`${pct(estimate.mean)}%`}
            data-testid="sim-distribution-mean"
          ></span>
        </div>
        <div class="text-muted numeric flex justify-between text-[13px]">
          <span>{grouped(estimate.min)}</span>
          <span>{spreadCopy.rangeMean(oneDecimal(estimate.mean))}</span>
          <span>{grouped(estimate.max)}</span>
        </div>
      </div>
      <p class="text-muted text-[13px]" data-testid="sim-spread-caption">{spreadCopy.rangeCaption}</p>
    {/if}
  </div>
  <div class="grid grid-cols-3 gap-x-4 gap-y-[14px]" data-testid="sim-spread-tiles">
    {#each tiles as tile (tile.id)}
      <Figure label={tile.label} value={tile.value} size="sm" testid={`sim-spread-tile-${tile.id}`} />
    {/each}
  </div>
  {#if benchmark !== null}
    <Row
      title={spreadCopy.benchmark(benchmark.band, benchmark.preset)}
      href={bisPageHref(spec, benchmark.faction, benchmark.band)}
      open
      testid="sim-spread-benchmark"
    >
      <!-- "1,004.0 · 9.9% under": the set's DPS and the run against it, once (h-audit 2026-11-09). -->
      {#snippet trail()}<span class="text-[15px] tabular-nums"
          ><b class="text-strong">{oneDecimal(benchmark.dps)}</b>
          <span class="text-muted">· {gapLine}</span></span
        >{/snippet}
    </Row>
  {/if}
</div>
