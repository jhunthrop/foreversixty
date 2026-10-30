<!-- web/src/components/chart/ChartTooltip.svelte -->
<!-- The chart tooltip (controls and charts spec §B5.7): the time, the events within ±2 s (a death
     with its killing blow), then the series at that second, sorted, in their colours. It renders
     in the browser's top layer (popover="manual"), so no card, stacking context or short plot
     ever clips it; without the Popover API it is a fixed box at the same place. -->
<script lang="ts" module>
  export interface TooltipRow {
    id: string;
    label: string;
    colour: string;
    value: string;
    share?: string;
    dead: boolean;
    highlighted: boolean;
  }
</script>

<script lang="ts">
  import { clockLabel } from '../../lib/chart/time-axis';
  import { chartCopy } from '../../lib/chart/copy';
  import type { ChartEvent } from '../../lib/chart/types';
  import Glyph from '../ui/Glyph.svelte';

  let {
    open,
    x,
    y,
    flip,
    title,
    note,
    events,
    rows,
    more,
    width,
    threat = false,
  }: {
    open: boolean;
    /** The crosshair's x and the plot's top, in viewport px. */
    x: number;
    y: number;
    /** True in the plot's right 55 %: the box sits left of the crosshair. */
    flip: boolean;
    title: string;
    note: string;
    events: readonly ChartEvent[];
    rows: readonly TooltipRow[];
    more: number;
    width: number;
    threat?: boolean;
  } = $props();

  let box = $state<HTMLDivElement | null>(null);
  let height = $state(0);
  const supported =
    typeof HTMLElement !== 'undefined' && typeof HTMLElement.prototype.showPopover === 'function';

  $effect(() => {
    const element = box;
    if (element === null || !supported) return;
    if (open && !element.matches(':popover-open')) element.showPopover();
    if (!open && element.matches(':popover-open')) element.hidePopover();
  });

  const left = $derived.by(() => {
    const viewport = typeof window === 'undefined' ? 1440 : window.innerWidth;
    const wanted = flip ? x - 16 - width : x + 16;
    return Math.max(8, Math.min(viewport - width - 8, wanted));
  });
  const top = $derived.by(() => {
    const viewport = typeof window === 'undefined' ? 900 : window.innerHeight;
    return Math.max(8, Math.min(viewport - height - 8, y));
  });

  const glyphOf = (event: ChartEvent): 'skull' | 'target' | 'bolt' =>
    event.kind === 'death' ? 'skull' : event.kind === 'target-change' ? 'target' : 'bolt';
</script>

<div
  bind:this={box}
  bind:clientHeight={height}
  popover="manual"
  class="chart-tooltip"
  class:fallback={!supported}
  class:shown={open}
  style:left="{left}px"
  style:top="{top}px"
  style:width="{width}px"
  data-testid="chart-tooltip"
  aria-hidden="true"
>
  <div class="head flex items-baseline justify-between gap-3 px-3 pb-1.5">
    <span class="text-strong numeric text-[14px] font-bold" data-testid="chart-tooltip-time">{title}</span>
    <span class="text-muted truncate text-[12px]">{note}</span>
  </div>
  {#if events.length > 0}
    <ul class="events flex flex-col gap-1 px-3 py-1.5" data-testid="chart-tooltip-events">
      {#each events as event (event.key)}
        <li class="flex items-start gap-1.5 text-[12px] leading-4" class:death={event.kind === 'death'}>
          <span class="text-muted numeric w-[30px] shrink-0">{clockLabel(event.atMs)}</span>
          {#if event.iconUrl}
            <img src={event.iconUrl} alt="" width="16" height="16" class="rounded-[3px]" />
          {:else}
            <span class="icon flex h-4 w-4 shrink-0 items-center justify-center rounded-[3px]"
              ><Glyph name={glyphOf(event)} size={12} /></span
            >
          {/if}
          <span>{event.text}</span>
        </li>
      {/each}
    </ul>
  {/if}
  {#if threat && rows.length > 0}
    <div class="text-muted flex justify-end gap-2 px-3 pt-1 text-[11px]">
      <span>{chartCopy.tooltipThreatColumn}</span><span class="w-[50px] text-right"
        >{chartCopy.tooltipShareColumn}</span
      >
    </div>
  {/if}
  <ul class="flex flex-col" data-testid="chart-tooltip-rows">
    {#each rows as row (row.id)}
      <li
        class="row flex h-[22px] items-center gap-2 px-3"
        class:hl={row.highlighted}
        data-testid="chart-tooltip-row"
      >
        <span
          class="h-3 w-[3px] shrink-0 rounded-[1px]"
          style:background={row.dead ? 'var(--color-muted)' : row.colour}
        ></span>
        <span
          class="min-w-0 flex-1 truncate text-[13px] font-semibold"
          style:color={row.dead ? 'var(--color-muted)' : row.colour}>{row.label}</span
        >
        <span class="numeric text-[13px] font-semibold {row.dead ? 'text-muted' : 'text-strong'}"
          >{row.value}</span
        >
        {#if threat}
          <span class="text-muted numeric w-[50px] text-right text-[12px]">{row.share ?? ''}</span>
        {/if}
      </li>
    {/each}
  </ul>
  {#if more > 0}
    <p class="text-muted px-3 pt-1 text-[12px]">{chartCopy.tooltipMore(more)}</p>
  {/if}
</div>

<style>
  .chart-tooltip {
    position: fixed;
    inset: auto;
    margin: 0;
    padding: 8px 0 6px;
    border: 1px solid var(--color-line-card);
    border-radius: 10px;
    background: #0d0f15;
    color: var(--color-text);
    box-shadow:
      inset 0 1px 0 var(--card-highlight),
      0 16px 38px rgba(0, 0, 0, 0.6);
    pointer-events: none;
    overflow: hidden;
    transition: opacity 80ms;
  }
  .chart-tooltip.fallback {
    display: none;
    z-index: 60;
  }
  .chart-tooltip.fallback.shown {
    display: block;
  }
  .head {
    border-bottom: 1px solid var(--color-line-soft);
  }
  .events {
    border-bottom: 1px solid var(--color-line-soft);
  }
  .events .death {
    color: #ffb4b0;
  }
  .icon {
    background: var(--color-raised);
    color: var(--color-text);
  }
  .death .icon {
    background: var(--color-death-soft);
    color: var(--color-death);
  }
  .row.hl {
    background: rgba(229, 185, 85, 0.08);
  }
  @media (prefers-reduced-motion: reduce) {
    .chart-tooltip {
      transition: none;
    }
  }
</style>
