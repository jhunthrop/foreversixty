<!-- web/src/components/sim/DpsDistribution.svelte -->
<!-- The run's DPS spread (spec 2026-11-04-sim-results 3.6): the engine's real bins as at most 40
     bars, the mean line, the two-thirds caption and six tiles. A result saved before the engine
     kept the bins draws a range strip instead, and says so: a drawn shape the run did not produce
     would break the accuracy tenet. No chart library: the bars are spans, the tiles are the data.
     Tiles sit under the plot, or beside it where the card is full width (1024 to 1279). -->
<script lang="ts">
  import { grouped, oneDecimal, spreadCopy } from '../../lib/sim/result-copy';
  import { spreadOf, type SpreadBar } from '../../lib/sim/spread';
  import type { DpsHistogram, Estimate } from '../../lib/sim/types';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import StatTile from '../ui/StatTile.svelte';

  let {
    estimate,
    iterationsRun,
    histogram,
  }: { estimate: Estimate; iterationsRun: number; histogram?: DpsHistogram } = $props();

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
    { id: 'error', label: spreadCopy.tiles.error, value: `± ${oneDecimal(estimate.error)}` },
    { id: 'spread', label: spreadCopy.tiles.spread, value: `± ${oneDecimal(estimate.stddev)}` },
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

<div class="grid grid-cols-1 gap-4 p-5 lg:grid-cols-[7fr_5fr] xl:grid-cols-1" data-testid="sim-distribution">
  <div class="flex min-w-0 flex-col gap-3">
    {#if spread !== null}
      <div class="pt-6">
        <!-- One image with a readout: hover or arrow keys step the bars (spec 3.6). -->
        <!-- svelte-ignore a11y_no_noninteractive_tabindex, a11y_no_noninteractive_element_interactions -->
        <div
          bind:this={plot}
          class="focus-visible:outline-gold border-line relative flex h-[168px] items-end gap-[2px] border-b focus-visible:outline-2 focus-visible:outline-offset-2"
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
              class={`min-w-0 flex-1 rounded-t-[2px] ${index === active ? 'bg-gold-light' : bar.central ? 'bg-[rgba(229,185,85,.7)]' : 'bg-line'}`}
              style:height={`${Math.max(1, (bar.count / spread.peak) * 100)}%`}
              data-testid="sim-spread-bar"
              onpointerenter={() => (active = index)}
              aria-hidden="true"
            ></span>
          {/each}
          <span
            class="bg-gold-light pointer-events-none absolute -top-1.5 bottom-0 w-[2px] -translate-x-1/2"
            style:left={`${meanAt}%`}
            data-testid="sim-distribution-mean"
          ></span>
          <span
            class="text-gold-light numeric pointer-events-none absolute -top-6 -translate-x-1/2 text-[13px] font-semibold"
            style:left={`${meanAt}%`}>{oneDecimal(estimate.mean)}</span
          >
        </div>
        <div class="text-muted numeric mt-2 flex justify-between text-[13px]">
          <span>{grouped(estimate.min)}</span>
          <span>{spreadCopy.axis}</span>
          <span>{grouped(estimate.max)}</span>
        </div>
      </div>
      <p class="text-muted text-[13px]" data-testid="sim-spread-caption">
        {spreadCopy.caption(grouped(spread.low), grouped(spread.high))}
      </p>
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
            class="absolute inset-y-0 bg-[rgba(229,185,85,.35)]"
            style:left={`${Math.max(0, pct(estimate.mean - estimate.stddev))}%`}
            style:right={`${Math.max(0, 100 - pct(estimate.mean + estimate.stddev))}%`}
          ></span>
          <span
            class="bg-gold-light absolute -inset-y-1 w-[2px] -translate-x-1/2"
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
  <div class="grid grid-cols-2 gap-2 self-start sm:grid-cols-3" data-testid="sim-spread-tiles">
    {#each tiles as tile (tile.id)}
      <StatTile label={tile.label} value={tile.value} testid={`sim-spread-tile-${tile.id}`} />
    {/each}
  </div>
</div>
