<!-- web/src/components/character/ChartPlot.svelte -->
<!-- The drawing and the pointer and keyboard of both History charts (experience spec
     2026-10-24-character-history 3.5 and 3.8). Everything it draws comes from the geometry
     history-chart.ts built; it measures its own box and reports the size so the parent can
     build that geometry. No animation: the resting state is the only state. -->
<script lang="ts">
  import { drawnDotIndexes, type ChartGeometry } from '../../lib/character/history-chart';
  import type { ChartSize } from '../../lib/character/history-scales';
  import { nearestIndex, type Step } from '../../lib/character/history-select';
  import { tierLabel, tierTextColor } from '../../lib/score/score';

  /** The viewport width below which the chart uses its phone aspect, margins and labels. */
  const PHONE_BELOW_PX = 640;
  const KEY_STEPS: Record<string, Step> = {
    ArrowLeft: 'earlier',
    ArrowRight: 'later',
    Home: 'first',
    End: 'newest',
  };

  let {
    geometry,
    size,
    label,
    selected,
    hover,
    selectedFill,
    rings = new Set<number>(),
    onmeasure,
    onselect,
    onhover,
    onstep,
  }: {
    /** Null before the first measure, or for a series too short to draw. */
    geometry: ChartGeometry | null;
    /** The size the geometry was built for; null until measured. */
    size: ChartSize | null;
    label: string;
    selected: number;
    hover: number | null;
    /** The selected point's fill: its tier colour (score) or parse colour (parse). */
    selectedFill: string | null;
    /** Indexes drawn as rings (private or guild points). */
    rings?: ReadonlySet<number>;
    onmeasure: (size: ChartSize) => void;
    onselect: (index: number) => void;
    onhover: (index: number | null) => void;
    onstep: (move: Step) => void;
  } = $props();

  let box: HTMLDivElement | undefined = $state();

  $effect(() => {
    const element = box;
    if (element === undefined) return;
    const observer = new ResizeObserver(([entry]) => {
      const { width, height } = entry.contentRect;
      if (width > 0 && height > 0) {
        onmeasure({ width, height, phone: window.innerWidth < PHONE_BELOW_PX });
      }
    });
    observer.observe(element);
    return () => observer.disconnect();
  });

  const shown = $derived(hover ?? selected);
  const drawn = $derived(geometry === null ? [] : drawnDotIndexes(geometry.dots.length, selected));
  const guideX = $derived(geometry?.dots[shown]?.x ?? null);

  function nearestTo(event: PointerEvent): number {
    const left = box?.getBoundingClientRect().left ?? 0;
    return nearestIndex(
      (geometry?.dots ?? []).map((dot) => dot.x),
      event.clientX - left,
    );
  }

  function onpointermove(event: PointerEvent): void {
    if (event.pointerType === 'mouse' || event.buttons > 0) onhover(nearestTo(event));
  }

  function onpointerup(event: PointerEvent): void {
    onselect(nearestTo(event));
    if (event.pointerType !== 'mouse') onhover(null);
  }

  function onkeydown(event: KeyboardEvent): void {
    const move = KEY_STEPS[event.key];
    if (move !== undefined) {
      event.preventDefault();
      onstep(move);
    } else if (event.key === ' ') {
      event.preventDefault();
    }
  }
</script>

<div
  bind:this={box}
  class="plot-box rounded-control relative aspect-[3/2] w-full max-w-full sm:aspect-[16/5] xl:aspect-[2/1]"
  style="touch-action: pan-y"
  role="group"
  tabindex="0"
  aria-label={label}
  data-testid="history-plot"
  {onpointermove}
  {onpointerup}
  onpointerleave={() => onhover(null)}
  onpointercancel={() => onhover(null)}
  {onkeydown}
>
  {#if geometry !== null && size !== null}
    {@const plot = geometry.plot}
    <svg
      class="absolute inset-0 block"
      width={size.width}
      height={size.height}
      viewBox={`0 0 ${size.width} ${size.height}`}
      aria-hidden="true"
      data-testid="history-svg"
    >
      {#each geometry.bands as band (band.tier)}
        {#if band.color !== null}
          <rect
            x={plot.left}
            y={band.top}
            width={plot.width}
            height={band.bottom - band.top}
            fill={band.color}
            fill-opacity={band.opacity}
          />
        {/if}
        {#if band.top > plot.top + 0.5}
          <line
            x1={plot.left}
            x2={plot.left + plot.width}
            y1={band.top}
            y2={band.top}
            stroke="var(--color-line)"
          />
        {/if}
        {#if band.labelled}
          <text
            x={plot.left + 8}
            y={band.bottom - 6}
            font-size="11"
            font-weight="600"
            fill={tierTextColor(band.tier)}
            fill-opacity="0.85"
            data-testid="history-band-label">{tierLabel(band.tier)}</text
          >
        {/if}
      {/each}

      {#each geometry.yTicks as tick (tick.value)}
        {#if !tick.boundary}
          <line
            x1={plot.left}
            x2={plot.left + plot.width}
            y1={tick.y}
            y2={tick.y}
            stroke="var(--color-line-soft)"
          />
        {/if}
        <text
          x={plot.left - 8}
          y={tick.y}
          dy="0.32em"
          text-anchor="end"
          font-size="11"
          class="numeric"
          fill="var(--color-muted)">{tick.value}</text
        >
      {/each}

      {#each geometry.gaps as gap (gap.from + gap.to)}
        <rect
          x={gap.left}
          y={plot.top}
          width={gap.right - gap.left}
          height={plot.height}
          fill="var(--color-muted)"
          fill-opacity="0.08"
          data-testid="history-gap"
        />
        {#each [gap.left, gap.right] as edge (edge)}
          <line
            x1={edge}
            x2={edge}
            y1={plot.top}
            y2={plot.top + plot.height}
            stroke="var(--color-muted)"
            stroke-opacity="0.4"
            stroke-dasharray="3 3"
          />
        {/each}
        {#each gap.labelLines as line (line.y)}
          <text
            x={(gap.left + gap.right) / 2}
            y={line.y}
            dy="0.32em"
            text-anchor="middle"
            font-size="11"
            fill="var(--color-muted)"
            data-testid="history-gap-label">{line.text}</text
          >
        {/each}
      {/each}

      {#if guideX !== null}
        <line
          x1={guideX}
          x2={guideX}
          y1={plot.top}
          y2={plot.top + plot.height}
          stroke={hover === null ? 'var(--color-line)' : 'var(--color-line-warm-strong)'}
          data-testid="history-guide"
        />
      {/if}

      {#each geometry.runs as run (run)}
        <path
          d={run}
          fill="none"
          stroke="var(--color-gold)"
          stroke-width="2"
          stroke-linejoin="round"
          stroke-linecap="round"
        />
      {/each}

      {#each drawn as index (index)}
        {@const dot = geometry.dots[index]}
        {@const ring = rings.has(index)}
        {#if index === selected}
          <circle
            cx={dot.x}
            cy={dot.y}
            r="5"
            fill={ring || selectedFill === null ? 'var(--color-raised)' : selectedFill}
            stroke="var(--color-strong)"
            stroke-width="2"
            data-testid="history-selected-dot"
          />
        {:else if ring}
          <circle
            cx={dot.x}
            cy={dot.y}
            r="3.5"
            fill="var(--color-raised)"
            stroke="var(--color-gold)"
            stroke-width="1.5"
            data-testid="history-ring"
          />
        {:else}
          <circle cx={dot.x} cy={dot.y} r="3" fill="var(--color-gold)" />
        {/if}
      {/each}

      {#each geometry.xTicks as tick (tick.x)}
        <text
          x={tick.x}
          y={plot.top + plot.height + 18}
          text-anchor={tick.anchor}
          font-size="11"
          class="numeric"
          fill="var(--color-muted)"
          data-testid="history-x-label">{tick.label}</text
        >
      {/each}
    </svg>
  {/if}
</div>
