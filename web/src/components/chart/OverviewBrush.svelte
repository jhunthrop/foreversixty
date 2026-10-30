<!-- web/src/components/chart/OverviewBrush.svelte -->
<!-- The whole-fight strip under every report chart (controls and charts spec §B5.9): the raid's
     damage per second as a muted area, phase starts, deaths, the window and its two handles.
     The handles ARE the two native range inputs the old TimeChart drew (the keyboard and
     screen-reader path to the window), layered over the strip and styled. A handle's drag shows
     the new edge at once and calls onWindow on release; dragging the window body pans it.
     Keys: arrows ±1 s, Shift+arrows ±10 s, PageUp and PageDown ±30 s, Home and End. The inputs
     take any value (step="any"): a whole-second step that does not divide the fight's length made
     the browser snap the end handle down a second short of the end. Drags snap in lib/chart/brush. -->
<script lang="ts">
  import { brushSnap } from '../../lib/chart/brush';
  import { clockLabel, endClockLabel } from '../../lib/chart/time-axis';
  import { BUCKET_MS, type TimeWindow } from '../../lib/report/window';

  let {
    domainMs,
    window: current,
    shape,
    phaseStarts = [],
    deaths = [],
    onWindow,
    coarse = false,
  }: {
    domainMs: number;
    window: TimeWindow;
    /** The raid's damage per second over the whole fight, one value per second. */
    shape: readonly number[];
    phaseStarts?: readonly number[];
    deaths?: readonly number[];
    onWindow: (window: TimeWindow | null) => void;
    /** A coarse pointer: 12 px handles with 44 px hit areas and a 32 px strip. */
    coarse?: boolean;
  } = $props();

  /** The edges while a handle or the body is being dragged; null when nothing is. */
  let draft = $state<TimeWindow | null>(null);
  const shown = $derived(draft ?? current);
  const zoomed = $derived(shown.startMs > 0 || shown.endMs < domainMs);
  /** An edge on the clock: the fight's end reads as the fight list reads its length. */
  const edgeLabel = (ms: number): string => (ms >= domainMs ? endClockLabel(domainMs) : clockLabel(ms));
  const pct = (ms: number): number =>
    domainMs <= 0 ? 0 : (Math.max(0, Math.min(domainMs, ms)) / domainMs) * 100;

  /** The strip's area: a path in a 1000 x 100 box, stretched to the strip. */
  const area = $derived.by(() => {
    if (shape.length === 0) return '';
    const peak = shape.reduce((highest, value) => Math.max(highest, value), 0);
    if (peak <= 0) return '';
    const step = 1000 / shape.length;
    const points = shape.map(
      (value, index) => `${(index * step + step / 2).toFixed(1)},${(100 - (value / peak) * 88).toFixed(1)}`,
    );
    return `M0,100 L${points.join(' L')} L1000,100 Z`;
  });

  function commit(next: TimeWindow): void {
    draft = null;
    const startMs = Math.max(0, Math.min(next.startMs, domainMs - BUCKET_MS));
    const endMs = Math.max(startMs + BUCKET_MS, Math.min(next.endMs, domainMs));
    onWindow(startMs <= 0 && endMs >= domainMs ? null : { startMs, endMs });
  }

  function moved(edge: 'start' | 'end', ms: number): TimeWindow {
    const base = draft ?? current;
    return edge === 'start'
      ? { startMs: Math.min(ms, base.endMs - BUCKET_MS), endMs: base.endMs }
      : { startMs: base.startMs, endMs: Math.max(ms, base.startMs + BUCKET_MS) };
  }

  function onKey(edge: 'start' | 'end', event: KeyboardEvent): void {
    const value = edge === 'start' ? shown.startMs : shown.endMs;
    const steps: Record<string, number> = {
      ArrowRight: 1,
      ArrowUp: 1,
      ArrowLeft: -1,
      ArrowDown: -1,
      PageUp: 30,
      PageDown: -30,
    };
    let next: number | null = null;
    if (event.key in steps)
      next =
        value + steps[event.key] * (event.shiftKey && event.key.startsWith('Arrow') ? 10 : 1) * BUCKET_MS;
    else if (event.key === 'Home') next = 0;
    else if (event.key === 'End') next = domainMs;
    if (next === null) return;
    event.preventDefault();
    commit(moved(edge, Math.max(0, Math.min(domainMs, next))));
  }

  /** Panning the window body: the press's offset is kept, the window's length is too. */
  let pan: { fromX: number; width: number; window: TimeWindow } | null = null;
  function panDown(event: PointerEvent): void {
    if (!zoomed) return;
    const strip = (event.currentTarget as HTMLElement).parentElement!.getBoundingClientRect();
    (event.currentTarget as HTMLElement).setPointerCapture(event.pointerId);
    pan = { fromX: event.clientX, width: strip.width, window: current };
  }
  function panMove(event: PointerEvent): void {
    if (pan === null) return;
    const length = pan.window.endMs - pan.window.startMs;
    const delta =
      Math.round((((event.clientX - pan.fromX) / Math.max(pan.width, 1)) * domainMs) / BUCKET_MS) * BUCKET_MS;
    const startMs = Math.max(0, Math.min(domainMs - length, pan.window.startMs + delta));
    draft = { startMs, endMs: startMs + length };
  }
  function panUp(): void {
    if (pan === null) return;
    pan = null;
    if (draft !== null) commit(draft);
  }
</script>

<div class="overview pt-1.5" data-testid="overview-brush" data-coarse={coarse ? 'true' : undefined}>
  <div class="strip relative overflow-visible rounded-[6px] border" style:height={coarse ? '32px' : '36px'}>
    <svg
      class="absolute inset-0 h-full w-full"
      viewBox="0 0 1000 100"
      preserveAspectRatio="none"
      aria-hidden="true"
    >
      {#if area !== ''}<path d={area} class="shape" />{/if}
    </svg>
    {#each phaseStarts as at (at)}
      <span
        class="phase absolute inset-y-0 w-0 border-l border-dashed"
        style:left="{pct(at)}%"
        aria-hidden="true"
      ></span>
    {/each}
    {#each deaths as at, index (`${at}-${index}`)}
      <span class="death absolute inset-y-0 w-0" style:left="{pct(at)}%" aria-hidden="true"></span>
    {/each}
    <span class="dim absolute inset-y-0 left-0" style:width="{pct(shown.startMs)}%" aria-hidden="true"></span>
    <span class="dim absolute inset-y-0 right-0" style:width="{100 - pct(shown.endMs)}%" aria-hidden="true"
    ></span>
    {#if zoomed}
      <span
        class="window absolute inset-y-0 cursor-grab"
        style:left="{pct(shown.startMs)}%"
        style:width="{pct(shown.endMs) - pct(shown.startMs)}%"
        onpointerdown={panDown}
        onpointermove={panMove}
        onpointerup={panUp}
        onpointercancel={panUp}
        data-testid="brush-window"
        aria-hidden="true"
      ></span>
    {/if}
    <input
      class="handle"
      type="range"
      min="0"
      max={domainMs}
      step="any"
      value={shown.startMs}
      aria-label="Start"
      aria-valuetext={edgeLabel(shown.startMs)}
      data-testid="window-start"
      oninput={(event) => (draft = moved('start', brushSnap(Number(event.currentTarget.value), domainMs)))}
      onchange={(event) => commit(moved('start', brushSnap(Number(event.currentTarget.value), domainMs)))}
      onkeydown={(event) => onKey('start', event)}
    />
    <input
      class="handle"
      type="range"
      min="0"
      max={domainMs}
      step="any"
      value={shown.endMs}
      aria-label="End"
      aria-valuetext={edgeLabel(shown.endMs)}
      data-testid="window-end"
      oninput={(event) => (draft = moved('end', brushSnap(Number(event.currentTarget.value), domainMs)))}
      onchange={(event) => commit(moved('end', brushSnap(Number(event.currentTarget.value), domainMs)))}
      onkeydown={(event) => onKey('end', event)}
    />
  </div>
  <div class="labels numeric relative h-4 text-[11px] leading-4" aria-hidden="true">
    <span class="text-muted absolute left-0">0:00</span>
    {#if zoomed}
      <span
        class="text-gold absolute -translate-x-1/2 font-bold"
        style:left="{pct(shown.startMs)}%"
        data-testid="brush-start-label">{clockLabel(shown.startMs)}</span
      >
      <span
        class="text-gold absolute -translate-x-1/2 font-bold"
        style:left="{pct(shown.endMs)}%"
        data-testid="brush-end-label">{edgeLabel(shown.endMs)}</span
      >
    {/if}
    <span class="text-muted absolute right-0" data-testid="brush-domain-label">{endClockLabel(domainMs)}</span
    >
  </div>
</div>

<style>
  .strip {
    background: rgba(255, 255, 255, 0.025);
    border-color: var(--color-line-soft);
  }
  .shape {
    fill: color-mix(in srgb, var(--color-muted) 22%, transparent);
    stroke: color-mix(in srgb, var(--color-muted) 50%, transparent);
    stroke-width: 1;
    vector-effect: non-scaling-stroke;
  }
  .phase {
    border-color: color-mix(in srgb, var(--color-gold) 50%, transparent);
  }
  .death {
    border-left: 1.5px solid var(--color-death);
  }
  .dim {
    background: rgba(7, 9, 13, 0.6);
    pointer-events: none;
  }
  .window {
    background: rgba(229, 185, 85, 0.1);
    border: 1px solid color-mix(in srgb, var(--color-gold) 80%, transparent);
  }
  /* The two range inputs lie over the whole strip; only their thumbs take the pointer, so the
     body under them stays draggable and a press on bare strip does nothing. */
  .handle {
    position: absolute;
    inset: 0;
    width: 100%;
    height: 100%;
    margin: 0;
    background: transparent;
    appearance: none;
    pointer-events: none;
  }
  .handle:focus-visible {
    outline: none;
  }
  .handle::-webkit-slider-runnable-track {
    background: transparent;
    height: 100%;
  }
  .handle::-moz-range-track {
    background: transparent;
  }
  .handle::-webkit-slider-thumb {
    appearance: none;
    pointer-events: auto;
    cursor: ew-resize;
    width: 24px;
    height: 36px;
    margin-top: 0;
    border: none;
    background:
      linear-gradient(#0b0e14, #0b0e14) center 9px / 2px 1px no-repeat,
      linear-gradient(#0b0e14, #0b0e14) center 26px / 2px 1px no-repeat,
      linear-gradient(var(--color-gold), var(--color-gold)) center / 8px 24px no-repeat;
    border-radius: 3px;
  }
  .handle::-moz-range-thumb {
    pointer-events: auto;
    cursor: ew-resize;
    width: 8px;
    height: 24px;
    border: none;
    border-radius: 3px;
    background: var(--color-gold);
  }
  .handle:focus-visible::-webkit-slider-thumb {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  /* On a coarse pointer each input's own box is the 44 px hit target, past the 32 px strip. */
  .overview[data-coarse='true'] .handle {
    top: -7px;
    height: calc(100% + 14px);
  }
  .overview[data-coarse='true'] .handle::-webkit-slider-thumb {
    width: 44px;
    height: 44px;
    background: linear-gradient(var(--color-gold), var(--color-gold)) center / 12px 24px no-repeat;
  }
</style>
