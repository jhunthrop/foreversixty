<!-- web/src/components/report/TimeChart.svelte -->
<!-- The per-second chart, drawn on a canvas with no chart library: one path, one fill, a
     grid, death ticks and a brush rectangle is a hundred lines of drawing code and a
     library is forty kilobytes the island's budget does not have.
     Brushing is the primary interaction, so it has three ways in: drag on the canvas, the
     two range inputs beneath it (which are how a keyboard and a screen reader work the
     chart), and the presets. All three call the same onWindow. -->
<script lang="ts">
  import { formatDuration } from '../../lib/report/format';
  import { BUCKET_MS, isFullWindow, type TimeWindow } from '../../lib/report/window';

  let {
    series,
    durationMs,
    window: current,
    deaths,
    label,
    onWindow,
  }: {
    series: number[];
    durationMs: number;
    window: TimeWindow;
    deaths: { at_ms: number; name: string }[];
    label: string;
    onWindow: (window: TimeWindow | null) => void;
  } = $props();

  const HEIGHT = 120;

  let canvas = $state<HTMLCanvasElement | null>(null);
  let width = $state(720);
  let dragFrom = $state<number | null>(null);
  let dragTo = $state<number | null>(null);

  const peak = $derived(series.reduce((highest, value) => Math.max(highest, value), 0));

  function xOf(ms: number): number {
    return durationMs === 0 ? 0 : (ms / durationMs) * width;
  }

  function msOf(x: number): number {
    return Math.max(0, Math.min(durationMs, (x / Math.max(width, 1)) * durationMs));
  }

  function draw(): void {
    const element = canvas;
    if (element === null) return;
    const ratio = globalThis.devicePixelRatio ?? 1;
    element.width = Math.round(width * ratio);
    element.height = Math.round(HEIGHT * ratio);
    const context = element.getContext('2d');
    if (context === null) return;
    context.scale(ratio, ratio);
    context.clearRect(0, 0, width, HEIGHT);

    // Every colour is read off the element's computed style so the tokens in
    // src/styles/tokens.css stay the single source and no hex appears here.
    const styles = getComputedStyle(element);
    const gold = styles.getPropertyValue('--color-gold').trim();
    const line = styles.getPropertyValue('--color-line-soft').trim();
    const ember = styles.getPropertyValue('--color-ember').trim();

    context.strokeStyle = line;
    context.lineWidth = 1;
    for (let at = 0; at <= durationMs; at += 10_000) {
      const x = Math.round(xOf(at)) + 0.5;
      context.beginPath();
      context.moveTo(x, 0);
      context.lineTo(x, HEIGHT);
      context.stroke();
    }

    if (peak > 0 && series.length > 0) {
      const step = width / series.length;
      context.beginPath();
      context.moveTo(0, HEIGHT);
      series.forEach((value, index) => {
        context.lineTo(index * step, HEIGHT - (value / peak) * (HEIGHT - 8));
        context.lineTo((index + 1) * step, HEIGHT - (value / peak) * (HEIGHT - 8));
      });
      context.lineTo(width, HEIGHT);
      context.closePath();
      context.fillStyle = `${gold}33`;
      context.fill();
      context.strokeStyle = gold;
      context.lineWidth = 1.5;
      context.stroke();
    }

    context.strokeStyle = ember;
    context.lineWidth = 2;
    for (const death of deaths) {
      const x = Math.round(xOf(death.at_ms)) + 0.5;
      context.beginPath();
      context.moveTo(x, 0);
      context.lineTo(x, HEIGHT);
      context.stroke();
    }

    const from = dragFrom ?? current.startMs;
    const to = dragTo ?? current.endMs;
    if (!(from <= 0 && to >= durationMs)) {
      context.fillStyle = `${gold}22`;
      context.fillRect(xOf(Math.min(from, to)), 0, Math.abs(xOf(to) - xOf(from)), HEIGHT);
    }
  }

  $effect(() => {
    // Re-reads series, window, deaths and width, so any of them redraws the canvas.
    void [series, current, deaths, width, peak];
    draw();
  });

  $effect(() => {
    const element = canvas;
    if (element === null) return;
    const observer = new ResizeObserver(([entry]) => {
      width = Math.max(240, Math.round(entry.contentRect.width));
    });
    observer.observe(element);
    return () => observer.disconnect();
  });

  function localX(event: PointerEvent): number {
    const bounds = (event.currentTarget as HTMLCanvasElement).getBoundingClientRect();
    return event.clientX - bounds.left;
  }

  function onDown(event: PointerEvent): void {
    (event.currentTarget as HTMLCanvasElement).setPointerCapture(event.pointerId);
    dragFrom = msOf(localX(event));
    dragTo = dragFrom;
  }

  function onMove(event: PointerEvent): void {
    if (dragFrom === null) return;
    dragTo = msOf(localX(event));
  }

  function onUp(): void {
    if (dragFrom === null || dragTo === null) return;
    const startMs = Math.round(Math.min(dragFrom, dragTo));
    const endMs = Math.round(Math.max(dragFrom, dragTo));
    dragFrom = null;
    dragTo = null;
    // A tap rather than a drag clears the window, which is the fastest way back out.
    onWindow(endMs - startMs < BUCKET_MS ? null : { startMs, endMs });
  }
</script>

<figure
  class="border-line rounded-panel bg-raised m-0 flex flex-col gap-2 border p-3"
  data-testid="time-chart"
>
  <figcaption class="flex flex-wrap items-baseline justify-between gap-2">
    <span class="label text-muted">{label} per second</span>
    <span class="tabular text-muted font-mono text-[12px]" data-testid="window-label" aria-live="polite">
      {isFullWindow(current, durationMs)
        ? `Whole fight · ${formatDuration(durationMs)}`
        : `${formatDuration(current.startMs)} to ${formatDuration(current.endMs)}`}
    </span>
  </figcaption>

  <canvas
    bind:this={canvas}
    class="w-full touch-none"
    style={`height: ${HEIGHT}px`}
    onpointerdown={onDown}
    onpointermove={onMove}
    onpointerup={onUp}
    onpointercancel={onUp}
    data-testid="time-chart-canvas"
    aria-hidden="true"
  ></canvas>

  <!-- The 44px minimum is on each input, not on the label around it: a range input is
       dragged by a press anywhere inside its own box, and the label's height does nothing
       for that. Native is 16px tall. h-11 makes the whole strip the handle, which is why
       the thumb pseudo-element needs no sizing of its own -- and leaving the control's
       native appearance alone is what keeps accent-gold painting the filled track. -->
  <div class="flex flex-col gap-2 md:flex-row md:items-center md:gap-4">
    <label class="label text-muted flex min-h-11 flex-1 items-center gap-2 md:min-h-0" for="window-start">
      Start
      <input
        id="window-start"
        class="accent-gold h-11 flex-1 md:h-9"
        type="range"
        min="0"
        max={durationMs}
        step={BUCKET_MS}
        value={current.startMs}
        aria-valuetext={formatDuration(current.startMs)}
        data-testid="window-start"
        oninput={(event) => {
          const startMs = Number(event.currentTarget.value);
          onWindow({ startMs, endMs: Math.max(startMs + BUCKET_MS, current.endMs) });
        }}
      />
    </label>
    <label class="label text-muted flex min-h-11 flex-1 items-center gap-2 md:min-h-0" for="window-end">
      End
      <input
        id="window-end"
        class="accent-gold h-11 flex-1 md:h-9"
        type="range"
        min="0"
        max={durationMs}
        step={BUCKET_MS}
        value={current.endMs}
        aria-valuetext={formatDuration(current.endMs)}
        data-testid="window-end"
        oninput={(event) => {
          const endMs = Number(event.currentTarget.value);
          onWindow({ startMs: Math.min(current.startMs, endMs - BUCKET_MS), endMs });
        }}
      />
    </label>
    <button
      type="button"
      class="border-line-warm rounded-control text-text inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
      onclick={() => onWindow(null)}
      data-testid="window-reset"
    >
      Whole fight
    </button>
  </div>
</figure>
