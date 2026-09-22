<!-- web/src/components/report/TimeChart.svelte -->
<!-- The per-second chart, drawn on a canvas with no chart library: one path, one fill, a
     grid, death ticks and a brush rectangle is a hundred lines of drawing code and a
     library is forty kilobytes the island's budget does not have.
     Brushing is the primary interaction, so it has three ways in: drag on the canvas, the
     two range inputs beneath it (which are how a keyboard and a screen reader work the
     chart), and the presets. All three call the same onWindow. -->
<script lang="ts">
  import { formatAmount, formatDuration } from '../../lib/report/format';
  import { BUCKET_MS, isFullWindow, type TimeWindow } from '../../lib/report/window';

  let {
    series,
    extra = [],
    marks = [],
    phases = [],
    perSecond = true,
    durationMs,
    window: current,
    deaths,
    label,
    onWindow,
  }: {
    series: number[];
    /** More lines drawn behind the main one, each in its token's colour. */
    extra?: { label: string; series: number[]; token: string }[];
    /** Ticks on the time axis, named on hover: the taunts under a threat chart. */
    marks?: { atMs: number; label: string }[];
    /** The fight's phases, drawn as bands with their names at their left edge. */
    phases?: { name: string; start_ms: number; end_ms: number }[];
    /** False when the lines are a running total rather than a rate: the caption and readout drop "per second". */
    perSecond?: boolean;
    durationMs: number;
    window: TimeWindow;
    deaths: { at_ms: number; name: string }[];
    label: string;
    onWindow: (window: TimeWindow | null) => void;
  } = $props();

  const HEIGHT = 96;

  let canvas = $state<HTMLCanvasElement | null>(null);
  let width = $state(720);
  let dragFrom = $state<number | null>(null);
  let dragTo = $state<number | null>(null);
  /** The second under the pointer, or null when it is off the canvas. */
  let hoverMs = $state<number | null>(null);
  const seriesLength = $derived(
    extra.reduce((longest, line) => Math.max(longest, line.series.length), series.length),
  );
  const hoverIndex = $derived(
    hoverMs === null ? null : Math.min(seriesLength - 1, Math.floor(hoverMs / BUCKET_MS)),
  );
  const hoverValue = $derived(hoverIndex === null ? null : (series[hoverIndex] ?? 0));
  const hoverExtra = $derived(
    hoverIndex === null
      ? []
      : extra.map((line) => ({ label: line.label, value: line.series[hoverIndex] ?? 0 })),
  );

  const peak = $derived(
    [series, ...extra.map((line) => line.series)]
      .flat()
      .reduce((highest, value) => Math.max(highest, value), 0),
  );

  /** The one place "/s" is spelled: the readout and the scale gutter both read off it. */
  const unit = $derived(perSecond ? '/s' : '');

  /** A fight time as a percentage of the chart's width, for the labels drawn over the canvas. */
  const pct = (ms: number): number => (durationMs === 0 ? 0 : (ms / durationMs) * 100);
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
    const death = styles.getPropertyValue('--color-death').trim();

    context.strokeStyle = line;
    context.lineWidth = 1;
    for (let at = 0; at <= durationMs; at += 10_000) {
      const x = Math.round(xOf(at)) + 0.5;
      context.beginPath();
      context.moveTo(x, 0);
      context.lineTo(x, HEIGHT);
      context.stroke();
    }

    phases.forEach((phase, index) => {
      if (index % 2 === 1) {
        context.fillStyle = `${gold}0d`;
        context.fillRect(xOf(phase.start_ms), 0, xOf(phase.end_ms) - xOf(phase.start_ms), HEIGHT);
      }
      if (phase.start_ms <= 0) return;
      const x = Math.round(xOf(phase.start_ms)) + 0.5;
      context.strokeStyle = gold;
      context.lineWidth = 1;
      context.setLineDash([2, 2]);
      context.beginPath();
      context.moveTo(x, 0);
      context.lineTo(x, HEIGHT);
      context.stroke();
      context.setLineDash([]);
    });

    for (const line of extra) {
      if (peak <= 0 || line.series.length === 0) continue;
      const lineStep = width / line.series.length;
      const name = line.token.replace(/^var\(/, '').replace(/\)$/, '');
      const colour = styles.getPropertyValue(name).trim() || line.token;
      context.beginPath();
      line.series.forEach((value, index) => {
        const y = HEIGHT - (value / peak) * (HEIGHT - 8);
        if (index === 0) context.moveTo(0, y);
        context.lineTo(index * lineStep, y);
        context.lineTo((index + 1) * lineStep, y);
      });
      context.strokeStyle = colour;
      context.lineWidth = 1.25;
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

    context.strokeStyle = death;
    context.fillStyle = death;
    context.lineWidth = 1.5;
    context.setLineDash([4, 3]);
    for (const death of deaths) {
      const x = Math.round(xOf(death.at_ms)) + 0.5;
      context.beginPath();
      context.moveTo(x, 6);
      context.lineTo(x, HEIGHT);
      context.stroke();
      // A cap at the top: the mark is a death, not a spike of the line under it.
      context.beginPath();
      context.moveTo(x - 4, 0);
      context.lineTo(x + 4, 0);
      context.lineTo(x, 6);
      context.closePath();
      context.fill();
    }
    context.setLineDash([]);

    if (marks.length > 0) {
      context.strokeStyle = gold;
      context.fillStyle = gold;
      context.lineWidth = 2;
      for (const mark of marks) {
        const x = Math.round(xOf(mark.atMs)) + 0.5;
        context.beginPath();
        context.moveTo(x, HEIGHT);
        context.lineTo(x, HEIGHT - 14);
        context.stroke();
        context.beginPath();
        context.arc(x, HEIGHT - 16, 3, 0, Math.PI * 2);
        context.fill();
      }
    }

    if (hoverMs !== null) {
      const x = Math.round(xOf(hoverMs)) + 0.5;
      context.strokeStyle = styles.getPropertyValue('--color-text').trim();
      context.lineWidth = 1;
      context.setLineDash([3, 3]);
      context.beginPath();
      context.moveTo(x, 0);
      context.lineTo(x, HEIGHT);
      context.stroke();
      context.setLineDash([]);
    }

    const from = dragFrom ?? current.startMs;
    const to = dragTo ?? current.endMs;
    if (!(from <= 0 && to >= durationMs)) {
      context.fillStyle = `${gold}22`;
      context.fillRect(xOf(Math.min(from, to)), 0, Math.abs(xOf(to) - xOf(from)), HEIGHT);
    }
  }

  // Drawing is deferred to the next animation frame and coalesced: the mount effect and
  // the ResizeObserver's first measurement both fire before anything is painted, so drawing
  // synchronously paints the canvas twice (once at the 720px default) inside the island's
  // mount task, which is the task Lighthouse's blocking-time budget measures. One frame
  // later is invisible to the viewer and moves the work into its own short task.
  // Until the ResizeObserver has measured the canvas, `width` is a guess, and a frame drawn
  // from it would be painted squashed into the real width; the observer's first callback
  // draws instead, synchronously, which still lands before the page's first paint.
  let measured = false;
  let frame = 0;
  function scheduleDraw(): void {
    if (frame !== 0) return;
    frame = requestAnimationFrame(() => {
      frame = 0;
      if (measured) draw();
    });
  }

  $effect(() => {
    // Re-reads series, window, deaths and width, so any of them redraws the canvas.
    void [series, extra, marks, phases, current, deaths, width, peak, hoverMs];
    scheduleDraw();
  });

  $effect(() => {
    return () => {
      if (frame !== 0) cancelAnimationFrame(frame);
    };
  });

  $effect(() => {
    const element = canvas;
    if (element === null) return;
    const observer = new ResizeObserver(([entry]) => {
      width = Math.max(240, Math.round(entry.contentRect.width));
      measured = true;
      draw();
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
    hoverMs = msOf(localX(event));
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
    <span class="label text-muted"
      >{#if series.length > 0}<span class="whitespace-nowrap"
          ><span class="bg-gold mr-1 inline-block h-[2px] w-[14px] align-middle" aria-hidden="true"
          ></span>{label}{perSecond ? ' per second' : ''}</span
        >{:else}<span class="whitespace-nowrap">{label}</span>{/if}{#each extra as line (line.label)}
        <span class="ml-3 tracking-normal whitespace-nowrap normal-case" data-testid="chart-line-label"
          ><span
            class="mr-1 inline-block h-[2px] w-[14px] align-middle"
            style={`background: ${line.token}`}
            aria-hidden="true"
          ></span>{line.label}</span
        >{/each}{#if deaths.length > 0}
        <span class="ml-3 tracking-normal whitespace-nowrap normal-case"
          ><span class="bg-death mr-1 inline-block h-[10px] w-[2px] align-middle" aria-hidden="true"
          ></span>{deaths.length === 1 ? 'a death' : `${deaths.length} deaths`}</span
        >{/if}</span
    >
    <span class="tabular text-muted font-mono text-[12px]" data-testid="window-label" aria-live="polite">
      {#if hoverMs !== null && hoverValue !== null}
        <span class="text-text mr-3" data-testid="chart-readout"
          >{formatDuration(hoverMs)}{#if series.length > 0}
            · {label.toLowerCase()}
            {formatAmount(hoverValue)}{unit}{/if}{#each hoverExtra as line (line.label)}
            · {line.label.toLowerCase()} {formatAmount(line.value)}{unit}{/each}</span
        >
      {/if}
      {isFullWindow(current, durationMs)
        ? `Whole fight · ${formatDuration(durationMs)}`
        : `${formatDuration(current.startMs)} to ${formatDuration(current.endMs)}`}
    </span>
  </figcaption>

  <!-- The scale sits over the canvas's left edge: the peak, half of it and zero, so the
       line says how much and not only when. The canvas keeps its full width for the brush. -->
  <div class="relative pl-12">
    <canvas
      bind:this={canvas}
      class="w-full touch-none"
      style={`height: ${HEIGHT}px`}
      onpointerdown={onDown}
      onpointermove={onMove}
      onpointerup={onUp}
      onpointercancel={onUp}
      onpointerleave={() => (hoverMs = null)}
      data-testid="time-chart-canvas"
      aria-hidden="true"
    ></canvas>
    {#if peak > 0}
      <div
        class="text-muted pointer-events-none absolute inset-y-0 left-0 flex w-11 flex-col items-end justify-between py-0.5 pr-1 text-right font-mono text-[10px] leading-none"
        data-testid="chart-scale"
        aria-hidden="true"
      >
        <span>{formatAmount(peak)}{unit}</span>
        <span>{formatAmount(peak / 2)}{unit}</span>
        <span>0</span>
      </div>
    {/if}
    {#if marks.length > 0}
      <div class="pointer-events-none absolute inset-y-0 right-0 left-12" aria-hidden="true">
        {#each marks as mark, position (`${mark.atMs}-${position}`)}
          <span
            class="bg-gold pointer-events-auto absolute bottom-0 block h-[18px] w-[3px]"
            style={`left: ${durationMs === 0 ? 0 : (mark.atMs / durationMs) * 100}%`}
            title={`${mark.label} · ${formatDuration(mark.atMs)}`}
            data-testid="chart-mark"
          ></span>
        {/each}
      </div>
    {/if}
    {#if phases.length > 0}
      <div class="pointer-events-none absolute inset-y-0 right-0 left-12" aria-hidden="true">
        {#each phases as phase, i (`${i}-${phase.name}`)}
          {@const from = pct(phase.start_ms)}
          {@const width = pct(phase.end_ms) - from}
          <!-- The name is clipped to its own band: a long curated name on a short phase
               ends in an ellipsis rather than running into the next band's name. -->
          <span
            class="text-muted absolute top-0 overflow-hidden font-mono text-[10px] leading-none text-ellipsis whitespace-nowrap"
            style={`left: calc(${from}% + 2px); max-width: calc(${width}% - 4px)`}
            data-testid="phase-band">{phase.name}</span
          >
        {/each}
      </div>
    {/if}
  </div>

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
