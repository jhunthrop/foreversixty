<!-- web/src/components/sim/TimelineBrush.svelte -->
<!-- The strip under the timeline (spec 2026-11-04-sim-results 3.7): the whole fight's resource
     line, and in Close up the window. Dragging the window pans it, dragging a handle resizes it
     (never under 10 s). Arrows pan 5 s; Shift with an arrow resizes. A press on bare strip moves
     the window there, or opens Close up there from Whole fight. It shares the charts' look, not
     chart/OverviewBrush: that one edits a raid chart's window from 0:00, one second at a time,
     with two range inputs, and this window starts before the pull and moves by a 5 s model. -->
<script lang="ts">
  import { clockLabel } from '../../lib/chart/time-axis';
  import { CHART_INK, inkCss } from '../../lib/chart/theme';
  import {
    BRUSH_STEP_MS,
    panWindow,
    resizeWindow,
    type TimelineModel,
    type ZoomWindow,
  } from '../../lib/sim/timeline';
  import { timelineCopy } from '../../lib/sim/timeline-copy';
  import { brushLine } from '../../lib/sim/timeline-draw';

  const STRIP_HEIGHT = 44;

  let {
    model,
    window: current,
    width,
    onWindow,
    onOpen,
  }: {
    model: TimelineModel;
    /** Null in Whole fight: the window is hidden. */
    window: ZoomWindow | null;
    width: number;
    onWindow: (window: ZoomWindow) => void;
    /** A press on the strip in Whole fight: Close up centred there. */
    onOpen: (atMs: number) => void;
  } = $props();

  const span = $derived(model.endMs - model.startMs);
  const path = $derived(width > 0 ? brushLine(model.resource, model, width, STRIP_HEIGHT) : '');
  const xFor = (atMs: number): number => (span <= 0 ? 0 : ((atMs - model.startMs) / span) * width);
  const msPerPx = $derived(width <= 0 ? 0 : span / width);

  type Grip = 'pan' | 'start' | 'end';
  let drag: { grip: Grip; fromX: number; from: ZoomWindow } | null = null;

  function press(grip: Grip, event: PointerEvent): void {
    if (current === null) return;
    event.stopPropagation();
    (event.currentTarget as HTMLElement).setPointerCapture(event.pointerId);
    drag = { grip, fromX: event.clientX, from: current };
  }

  function move(event: PointerEvent): void {
    if (drag === null) return;
    const delta = (event.clientX - drag.fromX) * msPerPx;
    onWindow(
      drag.grip === 'pan'
        ? panWindow(drag.from, delta, model)
        : resizeWindow(drag.from, drag.grip, delta, model),
    );
  }

  function release(): void {
    drag = null;
  }

  function pressStrip(event: PointerEvent): void {
    const strip = (event.currentTarget as HTMLElement).getBoundingClientRect();
    const atMs = model.startMs + (event.clientX - strip.left) * msPerPx;
    if (current === null) {
      onOpen(atMs);
      return;
    }
    const length = current.endMs - current.startMs;
    onWindow(panWindow(current, atMs - length / 2 - current.startMs, model));
  }

  function key(event: KeyboardEvent): void {
    if (current === null || (event.key !== 'ArrowLeft' && event.key !== 'ArrowRight')) return;
    event.preventDefault();
    const delta = event.key === 'ArrowRight' ? BRUSH_STEP_MS : -BRUSH_STEP_MS;
    onWindow(event.shiftKey ? resizeWindow(current, 'end', delta, model) : panWindow(current, delta, model));
  }
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  class="relative touch-pan-y"
  style:height="{STRIP_HEIGHT}px"
  style:width="{width}px"
  onpointerdown={pressStrip}
  data-testid="timeline-brush"
>
  <svg class="absolute inset-0" {width} height={STRIP_HEIGHT} aria-hidden="true">
    {#if path !== ''}<path d={path} fill="none" stroke={inkCss(CHART_INK.quiet)} stroke-width="1" />{/if}
  </svg>
  {#if current !== null}
    <div
      class="window"
      role="slider"
      tabindex="0"
      aria-label={timelineCopy.brushLabel}
      aria-valuemin={model.startMs}
      aria-valuemax={model.endMs}
      aria-valuenow={current.startMs}
      aria-valuetext={`${clockLabel(current.startMs)} to ${clockLabel(current.endMs)}`}
      style:left="{xFor(current.startMs)}px"
      style:width="{xFor(current.endMs) - xFor(current.startMs)}px"
      onpointerdown={(event) => press('pan', event)}
      onpointermove={move}
      onpointerup={release}
      onpointercancel={release}
      onkeydown={key}
      data-testid="timeline-brush-window"
    >
      {#each ['start', 'end'] as const as edge (edge)}
        <span
          class="grip"
          data-edge={edge}
          onpointerdown={(event) => press(edge, event)}
          onpointermove={move}
          onpointerup={release}
          onpointercancel={release}
          data-testid={`timeline-brush-${edge}`}
          aria-hidden="true"
        ></span>
      {/each}
    </div>
  {/if}
</div>

<style>
  .window {
    position: absolute;
    top: 4px;
    bottom: 4px;
    border: 1px solid var(--color-accent);
    border-radius: 4px;
    background: color-mix(in srgb, var(--color-accent) 10%, transparent);
    cursor: grab;
    touch-action: none;
  }
  .window:active {
    cursor: grabbing;
  }
  .window:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  /* A 6 x 20 handle in a hit area wide enough for a finger: the visible part is the ::after. */
  .grip {
    position: absolute;
    top: -4px;
    bottom: -4px;
    width: 22px;
    cursor: ew-resize;
    touch-action: none;
  }
  .grip[data-edge='start'] {
    left: -11px;
  }
  .grip[data-edge='end'] {
    right: -11px;
  }
  .grip::after {
    content: '';
    position: absolute;
    top: 12px;
    left: 8px;
    width: 6px;
    height: 20px;
    border-radius: 2px;
    background: var(--color-accent);
  }
  @media (hover: hover) and (pointer: fine) {
    .grip:hover::after {
      background: color-mix(in srgb, var(--color-accent) 80%, var(--color-strong));
    }
  }
  @media (pointer: coarse) {
    .grip {
      width: 44px;
    }
    .grip[data-edge='start'] {
      left: -22px;
    }
    .grip[data-edge='end'] {
      right: -22px;
    }
    .grip::after {
      left: 19px;
    }
  }
</style>
