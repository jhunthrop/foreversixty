<!-- web/src/components/chart/ChartLegend.svelte -->
<!-- The chip row under a chart's head (controls and charts spec §B5.3, §B6.4): the medians first
     with a dashed swatch, then the drawn players by the window's figure. A click toggles a line;
     hover or focus highlights it and shows `Only`; Alt-click is Only. An always-drawn chip has the
     pin and cannot be turned off. On a phone the row is one line, `+{n}` counts the chips that did
     not fit, and while a finger scrubs the plot each chip reads the value at the finger. -->
<script lang="ts">
  import { valueAt } from '../../lib/chart/buckets';
  import { chartCopy } from '../../lib/chart/copy';
  import type { ChartSeries } from '../../lib/chart/types';
  import { valueLabel } from '../../lib/chart/value-axis';
  import Glyph from '../ui/Glyph.svelte';

  let {
    series,
    hidden,
    highlightId,
    notDrawn,
    coarse = false,
    scrubMs = null,
    loading = false,
    formatValue = valueLabel,
    unit = '',
    onToggle,
    onOnly,
    onHighlight,
    onMore,
  }: {
    series: readonly ChartSeries[];
    hidden: ReadonlySet<string>;
    highlightId: string | null;
    /** The players the chart does not draw: `+{n} more`. */
    notDrawn: number;
    coarse?: boolean;
    /** A touch scrub's second: every chip reads its value there. */
    scrubMs?: number | null;
    loading?: boolean;
    formatValue?: (value: number) => string;
    /** What every chip's value is (`DPS`): the head's unit, over the head's span. */
    unit?: string;
    onToggle: (id: string) => void;
    onOnly: (id: string) => void;
    onHighlight: (id: string | null) => void;
    onMore: () => void;
  } = $props();

  let row = $state<HTMLDivElement | null>(null);
  let rowWidth = $state(0);
  /** On a phone, how many chips fit on the one row. */
  let fitting = $state(Number.POSITIVE_INFINITY);

  $effect(() => {
    const element = row;
    void [series, rowWidth, coarse];
    if (element === null || !coarse) {
      fitting = Number.POSITIVE_INFINITY;
      return;
    }
    const chips = [...element.querySelectorAll<HTMLElement>('[data-chip]')];
    // Widths, not positions: a chip already moved out of the row has no useful offsetLeft.
    const room = element.clientWidth - 24 - 44;
    let used = 0;
    let count = 0;
    for (const chip of chips) {
      used += chip.offsetWidth + (count === 0 ? 0 : 6);
      if (used > room) break;
      count += 1;
    }
    fitting = Math.max(1, count);
  });

  const valueOf = (line: ChartSeries): string => {
    if (scrubMs !== null) {
      const value = valueAt(line, scrubMs);
      return value === null || value === undefined ? '—' : formatValue(value);
    }
    return formatValue(line.figure);
  };
  /** Medians, then the pinned players (tanks, the Source player, you), then the rest: at every
   *  width the always-drawn lines are named before `+n` (fix round 2). On a phone the first chip
   *  alone carries the unit, so the row has room for the pinned names. */
  const medians = $derived(series.filter((line) => line.kind === 'median'));
  const pinnedLines = $derived(series.filter((line) => line.kind !== 'median' && line.pinned));
  const rest = $derived(series.filter((line) => line.kind !== 'median' && !line.pinned));
  /** A phone's one row names the pinned players first: they are the lines it promises to draw. */
  const ordered = $derived(
    coarse ? [...pinnedLines, ...medians, ...rest] : [...medians, ...pinnedLines, ...rest],
  );
  const overflow = $derived(coarse ? Math.max(0, series.length - fitting) : 0);
  const moreCount = $derived(coarse ? overflow + notDrawn : notDrawn);
</script>

<div
  bind:this={row}
  bind:clientWidth={rowWidth}
  class="legend flex gap-x-1.5 gap-y-1"
  class:flex-wrap={!coarse}
  class:coarse
  data-testid="chart-legend"
>
  {#if loading}
    {#each [96, 110, 88, 102, 92] as chipWidth, index (index)}
      <span class="skeleton-block h-7 shrink-0 rounded-[4px]" style:width="{chipWidth}px" aria-hidden="true"
      ></span>
    {/each}
  {:else}
    {#each ordered as line, index (line.id)}
      {@const off = hidden.has(line.id)}
      <span
        class="chip-wrap relative inline-flex shrink-0"
        class:overflowed={coarse && index >= fitting}
        data-chip
        role="presentation"
        onpointerenter={() => !coarse && onHighlight(line.id)}
        onpointerleave={() => !coarse && onHighlight(null)}
      >
        <button
          type="button"
          class="chip inline-flex items-center gap-1.5"
          class:off
          class:hl={highlightId === line.id}
          aria-pressed={!off}
          aria-label={off ? chartCopy.show(line.label) : chartCopy.hide(line.label)}
          title={line.pinned ? chartCopy.alwaysDrawn : undefined}
          data-testid="legend-chip"
          data-id={line.id}
          onclick={(event) => {
            if (event.altKey) onOnly(line.id);
            else if (!line.pinned) onToggle(line.id);
          }}
          onfocus={() => onHighlight(line.id)}
          onblur={() => onHighlight(null)}
        >
          <span
            class="swatch"
            class:dashed={line.kind === 'median'}
            style:--swatch={line.colour}
            aria-hidden="true"
          ></span>
          {#if line.pinned}<Glyph name="pin" size={12} class="text-muted" />{/if}
          <span class="name" style:color={line.kind === 'median' ? 'var(--color-strong)' : line.colour}
            >{line.label}</span
          >
          {#if line.you}<span class="you">{chartCopy.youChip}</span>{/if}
          <span class="value numeric" data-testid="legend-value"
            >{valueOf(line)}{#if unit !== '' && (!coarse || index === 0)}<span class="unit">{unit}</span
              >{/if}</span
          >
        </button>
        {#if !coarse && highlightId === line.id}
          <button
            type="button"
            class="only"
            aria-label={chartCopy.showOnly(line.label)}
            data-testid="legend-only"
            onclick={() => onOnly(line.id)}
            onfocus={() => onHighlight(line.id)}>{chartCopy.only}</button
          >
        {/if}
      </span>
    {/each}
    {#if moreCount > 0}
      <button type="button" class="more" class:ml-auto={coarse} data-testid="legend-more" onclick={onMore}
        >{coarse ? chartCopy.morePhone(moreCount) : chartCopy.more(moreCount)}</button
      >
    {/if}
  {/if}
</div>

<style>
  .legend {
    position: relative;
    min-height: 42px;
    padding: 8px 12px 6px;
    border-top: 1px solid var(--color-line-soft);
    align-items: center;
    overflow: hidden;
  }
  .legend.coarse {
    min-height: 37px;
    padding-top: 4px;
    padding-bottom: 4px;
    flex-wrap: nowrap;
  }
  /* A chip that did not fit the phone's one row: measured, never seen. */
  .chip-wrap.overflowed {
    position: absolute;
    visibility: hidden;
  }
  .chip {
    height: 28px;
    padding: 0 8px;
    border-radius: var(--radius-pill);
    background: transparent;
    font-size: 13px;
    cursor: pointer;
  }
  /* A coarse pointer's chip is a 44 px target that still lays out in the 28 px row. */
  .coarse .chip,
  .coarse .more {
    height: 44px;
    margin: -8px 0;
    padding: 0 6px;
  }
  .chip:hover,
  .chip.hl {
    background: rgba(255, 255, 255, 0.04);
  }
  .chip:focus-visible,
  .only:focus-visible,
  .more:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .chip-wrap:has(.only:focus-visible) .chip {
    outline: none;
  }
  .chip.off {
    opacity: 0.4;
  }
  .swatch {
    width: 12px;
    height: 3px;
    border-radius: 1px;
    background: var(--swatch);
  }
  .swatch.dashed {
    background: repeating-linear-gradient(90deg, var(--swatch) 0 4px, transparent 4px 7px);
  }
  .chip.off .swatch {
    background: transparent;
    box-shadow: inset 0 0 0 1px var(--swatch);
    height: 5px;
  }
  .name {
    font-weight: 600;
    white-space: nowrap;
  }
  .you {
    font-size: 10px;
    font-weight: 700;
    letter-spacing: 0.08em;
    color: var(--color-gold);
  }
  .value {
    min-width: 44px;
    text-align: left;
    font-weight: 500;
    color: var(--color-muted);
  }
  .coarse .value {
    min-width: 36px;
  }
  .unit {
    margin-left: 3px;
    font-size: 11px;
  }
  /* The fine pointer's chip reserves the slot its Only button appears in, so the button never
     covers the value and no chip moves on hover. */
  .legend:not(.coarse) .chip-wrap {
    padding-right: 38px;
  }
  .only {
    position: absolute;
    right: 2px;
    top: 50%;
    transform: translateY(-50%);
    padding: 2px 4px;
    border-radius: 4px;
    background: var(--color-raised);
    color: var(--color-gold);
    font-size: 12px;
    font-weight: 600;
  }
  .more {
    height: 28px;
    padding: 0 6px;
    color: var(--color-gold);
    font-size: 13px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    white-space: nowrap;
  }
</style>
