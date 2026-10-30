<!-- web/src/components/chart/TimeSeriesPlot.svelte -->
<!-- The shared report chart's working part (controls and charts spec §B4 to §B7): the lazy chunk
     TimeSeriesChart loads, with uPlot inside it. It owns the uPlot instance, the marker lane, the crosshair and its time pill, the drag
     selection, the tooltip, the keyboard and the touch scrub, the overview brush and the
     loading, empty and error states, all inside one reserved height. The adapters (report/
     TimeChart, report/ThreatChart) give it series, events and phases; the legend above it is
     ChartLegend. Every colour is read from the tokens at draw time. -->
<script lang="ts">
  import uPlot from 'uplot';
  import { untrack } from 'svelte';
  import { chartCopy } from '../../lib/chart/copy';
  import { chartCursor, clearChartCursor, setChartCursor } from '../../lib/chart/cursor.svelte';
  import { clusterRows, eventsNear, laneItems, type LaneItem } from '../../lib/chart/events';
  import { resolveColour, withAlpha } from '../../lib/chart/palette';
  import { clockLabel, timeTicks } from '../../lib/chart/time-axis';
  import { tooltipLines } from '../../lib/chart/tooltip';
  import type { ChartEvent, ChartPhase, ChartSeries } from '../../lib/chart/types';
  import {
    CHART_AXIS_PX,
    CHART_LANE_PX,
    CHART_EVENT_LINE_PX,
    gutterFor,
    plotHeightFor,
    type PlotHeights,
  } from '../../lib/chart/layout';
  import { LOG_FLOOR, logTicks, niceTop, valueLabel, valueTicks } from '../../lib/chart/value-axis';
  import { subscribeOverlays } from '../../lib/ui/overlay';
  import type { TimeWindow } from '../../lib/report/window';
  import Glyph from '../ui/Glyph.svelte';
  import ChartTooltip, { type TooltipRow } from './ChartTooltip.svelte';
  import OverviewBrush from './OverviewBrush.svelte';

  let {
    series,
    events = [],
    phases = [],
    domainMs,
    window: current,
    onWindow,
    plotHeights,
    scale = 'linear',
    syncKey,
    mode = 'rate',
    label,
    emptyCopy,
    hidden = new Set<string>(),
    highlightId = null,
    brushShape = [],
    loading = false,
    compact = false,
    tooltipWidth = 268,
    formatValue = valueLabel,
    shareOf,
    onScrub,
    onOpenDeath,
    testid = 'chart-frame',
  }: {
    series: readonly ChartSeries[];
    events?: readonly ChartEvent[];
    phases?: readonly ChartPhase[];
    /** The whole fight, in ms. */
    domainMs: number;
    window: TimeWindow;
    onWindow: (window: TimeWindow | null) => void;
    plotHeights: PlotHeights;
    scale?: 'linear' | 'log';
    /** Charts and lane views on one key share the crosshair (the fight). */
    syncKey: string;
    mode?: 'rate' | 'cumulative';
    /** The chart's title, for its accessible name. */
    label: string;
    emptyCopy: string;
    hidden?: ReadonlySet<string>;
    highlightId?: string | null;
    /** The raid's damage per second over the whole fight, for the brush's shape. */
    brushShape?: readonly number[];
    /** The data is still in flight. */
    loading?: boolean;
    /** The sample report's 96 px strip: no lane, no brush, no event line. */
    compact?: boolean;
    tooltipWidth?: number;
    formatValue?: (value: number) => string;
    /** The threat chart's share column: a series' share of everyone at that second. */
    shareOf?: (value: number, ms: number) => string;
    /** A touch scrub's second, so the legend reads out; null when it ends. */
    onScrub?: (ms: number | null) => void;
    /** A death marker (or a collapsed cluster) pressed: the report opens it on the Deaths tab in place. */
    onOpenDeath?: (deaths: readonly ChartEvent[]) => void;
    testid?: string;
  } = $props();

  const owner = `chart-${Math.random().toString(36).slice(2)}`;
  const AXIS_PX = CHART_AXIS_PX;
  const LANE_PX = CHART_LANE_PX;
  const EVENT_LINE_PX = CHART_EVENT_LINE_PX;
  const SCRUB_HINT_KEY = 'fs.chart.scrubHint';

  // The viewport decides the plot height, the y gutter and the pointer kind.
  let viewport = $state(typeof window === 'undefined' ? 1440 : window.innerWidth);
  let coarse = $state(false);
  let reducedMotion = $state(false);
  $effect(() => {
    const coarseQuery = window.matchMedia('(pointer: coarse), (hover: none)');
    const motionQuery = window.matchMedia('(prefers-reduced-motion: reduce)');
    const read = (): void => {
      viewport = window.innerWidth;
      coarse = coarseQuery.matches;
      reducedMotion = motionQuery.matches;
    };
    read();
    window.addEventListener('resize', read);
    coarseQuery.addEventListener('change', read);
    motionQuery.addEventListener('change', read);
    return () => {
      window.removeEventListener('resize', read);
      coarseQuery.removeEventListener('change', read);
      motionQuery.removeEventListener('change', read);
    };
  });
  const plotHeight = $derived(plotHeightFor(viewport, plotHeights));
  const gutter = $derived(gutterFor(viewport));
  const laneEvents = $derived(events.filter((event) => event.kind !== 'phase'));
  const showLane = $derived(!compact && laneEvents.length > 0);
  const showEventLine = $derived(!compact && coarse);
  const zoomed = $derived(current.startMs > 0 || current.endMs < domainMs);

  /** The visible series have nothing inside the window: the empty state. */
  const empty = $derived.by(() => {
    const from = Math.floor(current.startMs / 1000);
    const to = Math.ceil(current.endMs / 1000);
    return !series.some(
      (line) =>
        !hidden.has(line.id) &&
        line.kind !== 'median' &&
        line.values.slice(from, to + 1).some((value) => (value ?? 0) > 0),
    );
  });

  const status = $derived<'loading' | 'ready'>(loading ? 'loading' : 'ready');
  /** Menus, lists and sheets open over the page (spec §C): while any is, the chart stands down. */
  let openOverlays = $state(0);
  $effect(() => subscribeOverlays((count) => (openOverlays = count)));
  /** An icon the build does not ship falls back to the marker's glyph instead of a broken image. */
  let brokenIcons = $state<ReadonlySet<string>>(new Set());
  function markBroken(url: string | null | undefined): void {
    if (url) brokenIcons = new Set([...brokenIcons, url]);
  }
  let host = $state<HTMLDivElement | null>(null);
  let width = $state(0);
  let chart = $state.raw<uPlot | null>(null);
  /** Bumped whenever the plot's scale or size moves, so the DOM overlays re-place. */
  let layout = $state(0);
  /** Re-places the overlays; called from uPlot's hooks, which can fire inside an effect. */
  function bump(): void {
    untrack(() => (layout += 1));
  }
  const measured = $derived(width > 0);

  // ---- colours, read from the tokens at draw time --------------------------------------
  function colour(token: string, alpha = 1): string {
    const resolved = host === null ? token : resolveColour(host, token);
    return alpha === 1 ? resolved : withAlpha(resolved, alpha);
  }

  // ---- uPlot options ------------------------------------------------------------------
  const font = '11px Barlow, system-ui, sans-serif';

  /**
   * Every line over the whole fight, one value per second. A summary's series stops at its last
   * busy second: past it a rate is zero and a running total holds its last value.
   */
  const full = $derived.by((): ChartSeries[] => {
    const length = Math.max(1, Math.ceil(domainMs / 1000) + 1);
    return series.map((line) => {
      if (line.values.length >= length) return line;
      const tail = mode === 'cumulative' ? (line.values.at(-1) ?? 0) : 0;
      const pad = line.diedAtMs !== undefined && line.values.at(-1) === null ? null : tail;
      return {
        ...line,
        values: [...line.values, ...Array.from({ length: length - line.values.length }, () => pad)],
      };
    });
  });

  function data(): uPlot.AlignedData {
    const xs = Array.from({ length: full[0]?.values.length ?? 1 }, (_, index) => index);
    const lines = full.map((line) =>
      line.values.map((value) =>
        value === null ? null : scale === 'log' ? Math.max(LOG_FLOOR, value) : value,
      ),
    );
    return [xs, ...lines] as uPlot.AlignedData;
  }

  function options(): uPlot.Options {
    const fill = (line: ChartSeries): ((u: uPlot) => string) | undefined =>
      line.kind === 'single' ? () => colour(line.colour, 0.2) : undefined;
    return {
      width: Math.max(width, 200),
      height: plotHeight + AXIS_PX,
      padding: [8, 12, 0, 0],
      legend: { show: false },
      cursor: {
        y: false,
        drag: { x: false, y: false },
        focus: { prox: -1 },
        points: {
          size: 7,
          width: 1.5,
          fill: (u, index) => colour(series[index - 1]?.colour ?? 'var(--color-text)'),
          stroke: () => colour('var(--color-bg)'),
        },
      },
      focus: { alpha: 0.25 },
      scales: {
        x: { time: false, auto: false, range: () => [current.startMs / 1000, current.endMs / 1000] },
        y:
          scale === 'log'
            ? {
                distr: 3,
                log: 10,
                range: (_u, _min, max) => [LOG_FLOOR, 10 ** Math.ceil(Math.log10(Math.max(max, 10)))],
              }
            : { range: (_u, _min, max) => [0, niceTop(max)] },
      },
      axes: [
        {
          size: AXIS_PX,
          font,
          stroke: () => colour('var(--color-muted)'),
          grid: { show: false },
          ticks: { show: true, size: 4, width: 1, stroke: () => colour('var(--color-line)') },
          border: { show: true, width: 1, stroke: () => colour('var(--color-line)') },
          splits: (u) => timeTicks(u.scales.x.min ?? 0, u.scales.x.max ?? 0, u.bbox.width / devicePixelRatio),
          values: (_u, splits) => splits.map((at) => clockLabel(at * 1000)),
          gap: 4,
        },
        {
          size: gutter,
          font,
          stroke: () => colour('var(--color-muted)'),
          gap: 8,
          ticks: { show: false },
          grid: { show: true, width: 1, stroke: () => colour('var(--color-line-soft)') },
          splits: (u) => (scale === 'log' ? logTicks(u.scales.y.max ?? 10) : valueTicks(u.scales.y.max ?? 1)),
          values: (_u, splits) => splits.map((value) => valueLabel(value)),
        },
      ],
      series: [
        {},
        ...series.map((line) => ({
          label: line.label,
          stroke: () => colour(line.colour, line.kind === 'median' ? 0.7 : 1),
          width: 1.5,
          dash: line.dash === undefined ? undefined : [...line.dash],
          fill: fill(line),
          points: { show: false },
          spanGaps: false,
          show: !hidden.has(line.id),
        })),
      ],
      hooks: {
        drawClear: [drawBands],
        draw: [drawMarks],
        setCursor: [onCursor],
        setScale: [bump],
        setSize: [bump],
      },
    };
  }

  /** Every second phase is a faint band (spec §B5.4), under the lines. */
  function drawBands(u: uPlot): void {
    const { ctx } = u;
    const { top, height } = u.bbox;
    ctx.save();
    ctx.fillStyle = 'rgba(229,185,85,0.05)';
    phases.forEach((phase, index) => {
      if (index % 2 === 0) return;
      const from = u.valToPos(phase.startMs / 1000, 'x', true);
      const to = u.valToPos(phase.endMs / 1000, 'x', true);
      ctx.fillRect(from, top, to - from, height);
    });
    ctx.restore();
  }

  /** Phase starts, the plot lines of every marker, and a dead player's hollow end dot. */
  function drawMarks(u: uPlot): void {
    const { ctx } = u;
    const { top, height, left, width: plotWidth } = u.bbox;
    const ratio = devicePixelRatio;
    const line = (ms: number, stroke: string, dash: number[], lineWidth: number): void => {
      const x = Math.round(u.valToPos(ms / 1000, 'x', true)) + 0.5;
      if (x < left || x > left + plotWidth) return;
      ctx.strokeStyle = stroke;
      ctx.lineWidth = lineWidth * ratio;
      ctx.setLineDash(dash.map((part) => part * ratio));
      ctx.beginPath();
      ctx.moveTo(x, top);
      ctx.lineTo(x, top + height);
      ctx.stroke();
    };
    ctx.save();
    for (const phase of phases)
      if (phase.startMs > 0) line(phase.startMs, colour('var(--color-gold)', 0.7), [2, 3], 1);
    for (const event of events) {
      if (event.kind === 'death') line(event.atMs, colour('var(--color-death)'), [4, 3], 1.5);
      else if (event.kind === 'target-change') line(event.atMs, colour('var(--color-text)', 0.35), [1, 3], 1);
      else if (event.kind === 'taunt' || event.kind === 'cast')
        line(event.atMs, colour('var(--color-gold)', 0.35), [], 1);
    }
    ctx.setLineDash([]);
    series.forEach((entry, index) => {
      if (entry.diedAtMs === undefined || hidden.has(entry.id)) return;
      const at = Math.floor(entry.diedAtMs / 1000);
      const value = u.data[index + 1][at];
      if (value === null || value === undefined) return;
      const x = u.valToPos(at, 'x', true);
      const y = u.valToPos(value, 'y', true);
      if (x < left || x > left + plotWidth) return;
      ctx.strokeStyle = colour(entry.colour);
      ctx.lineWidth = 1.5 * ratio;
      ctx.fillStyle = colour('var(--color-bg)');
      ctx.beginPath();
      ctx.arc(x, y, 3 * ratio, 0, Math.PI * 2);
      ctx.fill();
      ctx.stroke();
    });
    ctx.restore();
  }

  // ---- the cursor ------------------------------------------------------------------------
  /** The second the crosshair is on, when this chart's own pointer, key or finger put it there. */
  let ownMs = $state<number | null>(null);
  let applyingExternal = false;

  function onCursor(u: uPlot): void {
    if (applyingExternal) return;
    const left = u.cursor.left ?? -1;
    if (left < 0 || openOverlays > 0) {
      if (!scrubbing && !keyboardActive) {
        ownMs = null;
        clearChartCursor(owner);
      }
      return;
    }
    const ms = Math.round(u.posToVal(left, 'x')) * 1000;
    ownMs = Math.max(0, Math.min(domainMs, ms));
    setChartCursor({ syncKey, ms: ownMs, owner });
  }

  /** Another view on the page moved the shared cursor: draw the crosshair there, no tooltip. */
  $effect(() => {
    const shared = chartCursor();
    const u = chart;
    if (u === null) return;
    untrack(() => {
      if (shared === null || shared.syncKey !== syncKey || shared.owner === owner || openOverlays > 0) {
        if (shared?.owner !== owner && ownMs === null) {
          applyingExternal = true;
          u.setCursor({ left: -10, top: -10 }, false);
          applyingExternal = false;
        }
        return;
      }
      applyingExternal = true;
      u.setCursor({ left: u.valToPos(shared.ms / 1000, 'x'), top: 10 }, false);
      applyingExternal = false;
    });
  });

  /** A menu opening over the plot stands the crosshair and tooltip down (spec §C). */
  $effect(() => {
    if (openOverlays === 0) return;
    untrack(() => {
      ownMs = null;
      clearChartCursor(owner);
      if (chart !== null) {
        applyingExternal = true;
        chart.setCursor({ left: -10, top: -10 }, false);
        applyingExternal = false;
      }
    });
  });

  function placeCursor(ms: number): void {
    const u = chart;
    if (u === null) return;
    const clamped = Math.max(current.startMs, Math.min(current.endMs, ms));
    ownMs = clamped;
    setChartCursor({ syncKey, ms: clamped, owner });
    applyingExternal = true;
    u.setCursor({ left: u.valToPos(clamped / 1000, 'x'), top: 10 }, false);
    applyingExternal = false;
  }

  // ---- building and updating the uPlot ---------------------------------------------------
  $effect(() => {
    const element = host;
    // Rebuilt when the lines, the scale, the height, the gutter or the phases change.
    void [series, scale, plotHeight, gutter, phases, events, domainMs];
    if (element === null || !measured) return;
    const built = untrack(
      () => new uPlot(options(), data(), element.querySelector<HTMLElement>('[data-uplot]')!),
    );
    chart = built;
    bump();
    return () => {
      built.destroy();
      if (chart === built) chart = null;
    };
  });

  $effect(() => {
    const u = chart;
    if (u === null || width === 0) return;
    if (Math.abs(u.width - width) > 0.5) untrack(() => u.setSize({ width, height: plotHeight + AXIS_PX }));
  });

  $effect(() => {
    const { startMs, endMs } = current;
    const u = chart;
    if (u === null) return;
    untrack(() => {
      if (u.scales.x.min !== startMs / 1000 || u.scales.x.max !== endMs / 1000)
        u.setScale('x', { min: startMs / 1000, max: endMs / 1000 });
    });
  });

  $effect(() => {
    const u = chart;
    const off = hidden;
    const focus = highlightId;
    if (u === null) return;
    untrack(() => {
      series.forEach((line, index) => {
        const show = !off.has(line.id);
        if (u.series[index + 1].show !== show) u.setSeries(index + 1, { show }, false);
        u.series[index + 1].width = focus === line.id ? 2.5 : 1.5;
      });
      const at = series.findIndex((line) => line.id === focus);
      u.setSeries(at < 0 ? null : at + 1, { focus: at >= 0 }, false);
      u.redraw(false, false);
    });
  });

  // ---- positions for the DOM overlays ------------------------------------------------------
  const plotLeft = $derived.by(() => {
    void layout;
    return chart === null ? gutter : chart.bbox.left / devicePixelRatio;
  });
  const plotWidth = $derived.by(() => {
    void layout;
    return chart === null ? Math.max(0, width - gutter - 12) : chart.bbox.width / devicePixelRatio;
  });
  function xOf(ms: number): number {
    void layout;
    const span = current.endMs - current.startMs;
    return plotLeft + (span <= 0 ? 0 : ((ms - current.startMs) / span) * plotWidth);
  }
  const visibleEvents = $derived(
    laneEvents.filter((event) => event.atMs >= current.startMs && event.atMs <= current.endMs),
  );
  const lane = $derived.by((): LaneItem[] => {
    void layout;
    return laneItems(visibleEvents, xOf);
  });
  const phaseLabels = $derived(
    phases
      .filter((phase) => phase.endMs > current.startMs && phase.startMs < current.endMs)
      .map((phase) => ({
        phase,
        left: xOf(Math.max(phase.startMs, current.startMs)) + 6,
      })),
  );

  // ---- the drag selection (fine pointer) ---------------------------------------------------
  let dragFrom = $state<number | null>(null);
  let dragTo = $state<number | null>(null);
  function msAt(clientX: number): number {
    const box = host?.getBoundingClientRect();
    if (box === undefined) return 0;
    const x = clientX - box.left - plotLeft;
    const span = current.endMs - current.startMs;
    const ms = current.startMs + (x / Math.max(plotWidth, 1)) * span;
    return Math.round(Math.max(current.startMs, Math.min(current.endMs, ms)) / 1000) * 1000;
  }

  // ---- touch: hold 300 ms to scrub ---------------------------------------------------------
  let scrubbing = $state(false);
  let scrubbedOnce = $state(false);
  let holdTimer: ReturnType<typeof setTimeout> | null = null;
  let pressAt: { x: number; y: number } | null = null;
  $effect(() => {
    try {
      scrubbedOnce = globalThis.localStorage?.getItem(SCRUB_HINT_KEY) === '1';
    } catch {
      /* storage can be unavailable; the hint simply shows */
    }
  });

  function cancelHold(): void {
    if (holdTimer !== null) clearTimeout(holdTimer);
    holdTimer = null;
  }

  function onPointerDown(event: PointerEvent): void {
    if (openOverlays > 0) return;
    if (event.pointerType === 'touch') {
      pressAt = { x: event.clientX, y: event.clientY };
      cancelHold();
      const clientX = event.clientX;
      const target = event.currentTarget as HTMLElement;
      const pointerId = event.pointerId;
      holdTimer = setTimeout(() => {
        scrubbing = true;
        try {
          target.setPointerCapture(pointerId);
        } catch {
          /* the pointer may have gone */
        }
        placeCursor(msAt(clientX));
        onScrub?.(ownMs);
        if (!scrubbedOnce) {
          scrubbedOnce = true;
          try {
            globalThis.localStorage?.setItem(SCRUB_HINT_KEY, '1');
          } catch {
            /* ignore */
          }
        }
      }, 300);
      return;
    }
    if (event.button !== 0 || compact) return;
    dragFrom = msAt(event.clientX);
    dragTo = dragFrom;
  }

  function onPointerMove(event: PointerEvent): void {
    if (event.pointerType === 'touch') {
      if (scrubbing) {
        placeCursor(msAt(event.clientX));
        onScrub?.(ownMs);
      } else if (pressAt !== null && Math.hypot(event.clientX - pressAt.x, event.clientY - pressAt.y) > 8) {
        cancelHold();
      }
      return;
    }
    if (dragFrom !== null) dragTo = msAt(event.clientX);
  }

  function onPointerUp(event: PointerEvent): void {
    if (event.pointerType === 'touch') {
      cancelHold();
      pressAt = null;
      scrubbing = false;
      return;
    }
    if (dragFrom === null || dragTo === null) return;
    const startMs = Math.min(dragFrom, dragTo);
    const endMs = Math.max(dragFrom, dragTo);
    dragFrom = null;
    dragTo = null;
    if (endMs - startMs >= 1000) onWindow({ startMs, endMs });
  }

  /** While scrubbing, a finger's move is the scrub, not a scroll. */
  $effect(() => {
    const element = host;
    if (element === null) return;
    const block = (event: TouchEvent): void => {
      if (scrubbing) event.preventDefault();
    };
    element.addEventListener('touchmove', block, { passive: false });
    return () => element.removeEventListener('touchmove', block);
  });

  /** A tap outside the plot clears the last reading. */
  $effect(() => {
    if (!coarse) return;
    const clear = (event: PointerEvent): void => {
      if (host !== null && !host.contains(event.target as Node) && ownMs !== null && !scrubbing) {
        ownMs = null;
        clearChartCursor(owner);
        onScrub?.(null);
        chart?.setCursor({ left: -10, top: -10 }, false);
      }
    };
    document.addEventListener('pointerdown', clear);
    return () => document.removeEventListener('pointerdown', clear);
  });

  // ---- keyboard ----------------------------------------------------------------------------
  let keyboardActive = false;
  let liveText = $state('');
  let liveTimer: ReturnType<typeof setTimeout> | null = null;
  function onKey(event: KeyboardEvent): void {
    if (event.key === 'Escape' && zoomed) {
      event.preventDefault();
      onWindow(null);
      return;
    }
    if (event.key !== 'ArrowLeft' && event.key !== 'ArrowRight') return;
    event.preventDefault();
    keyboardActive = true;
    const step = (event.shiftKey ? 10 : 1) * 1000 * (event.key === 'ArrowLeft' ? -1 : 1);
    placeCursor((ownMs ?? current.startMs) + step);
    if (liveTimer !== null) clearTimeout(liveTimer);
    liveTimer = setTimeout(() => {
      const at = ownMs;
      if (at === null) return;
      liveText = chartCopy.live(
        at,
        rowsAt(at)
          .lines.slice(0, 3)
          .map((row) => ({ name: row.label, value: row.valueText })),
      );
    }, 300);
  }
  function onBlur(): void {
    keyboardActive = false;
    ownMs = null;
    clearChartCursor(owner);
  }

  // ---- the tooltip ---------------------------------------------------------------------------
  function rowsAt(ms: number): ReturnType<typeof tooltipLines> {
    return tooltipLines(full, ms, { hidden, highlightId, format: formatValue });
  }
  const phaseAt = (ms: number): string | undefined =>
    phases.find((phase) => ms >= phase.startMs && ms < phase.endMs)?.name;
  let hoveredItem = $state<LaneItem | null>(null);
  const tooltipMs = $derived(hoveredItem?.startMs ?? ownMs);
  const tooltipOpen = $derived(
    tooltipMs !== null &&
      !coarse &&
      !scrubbing &&
      dragFrom === null &&
      openOverlays === 0 &&
      status === 'ready',
  );
  const tooltip = $derived.by(() => {
    if (tooltipMs === null) return null;
    const cluster = hoveredItem?.kind === 'cluster' && hoveredItem.tone === 'death' ? hoveredItem : null;
    if (cluster !== null) {
      const rows = clusterRows(cluster.events);
      return {
        title: `${clockLabel(cluster.startMs)} to ${clockLabel(cluster.endMs)}`,
        note: chartCopy.deathCluster(cluster.events.length),
        events: rows.shown.map((event) => ({ ...event, text: event.clusterText ?? event.text })),
        rows: [] as TooltipRow[],
        more: 0,
        clusterMore: rows.more,
      };
    }
    const { lines, more } = rowsAt(tooltipMs);
    return {
      title: clockLabel(tooltipMs),
      note: chartCopy.tooltipRight(phaseAt(tooltipMs), mode === 'cumulative'),
      events: eventsNear(events, tooltipMs).map((event) =>
        event.iconUrl && brokenIcons.has(event.iconUrl) ? { ...event, iconUrl: null } : event,
      ),
      rows: lines.map((line) => ({
        id: line.id,
        label: line.label,
        colour: line.colour,
        value: line.valueText,
        share:
          shareOf === undefined || line.dead || line.value === null
            ? undefined
            : shareOf(line.value, tooltipMs),
        dead: line.dead,
        highlighted: line.highlighted,
      })),
      more,
      clusterMore: 0,
    };
  });
  const box = $derived.by(() => {
    void layout;
    const rect = host?.getBoundingClientRect();
    return rect ?? { left: 0, top: 0 };
  });
  const crossX = $derived(tooltipMs === null ? 0 : xOf(tooltipMs));
  const tooltipFlip = $derived(crossX - plotLeft > plotWidth * 0.45);

  // ---- the event line (coarse pointer) -------------------------------------------------------
  const lineEvents = $derived(
    ownMs === null ? [] : eventsNear(events, ownMs).filter((event) => event.kind !== 'phase'),
  );

  const dragLabel = $derived(
    dragFrom === null || dragTo === null || dragFrom === dragTo
      ? null
      : { from: Math.min(dragFrom, dragTo), to: Math.max(dragFrom, dragTo) },
  );
</script>

<div
  class="tsc flex flex-col"
  data-testid={testid}
  data-state={status}
  data-empty={empty && status === 'ready' ? 'true' : undefined}
  data-compact={compact ? 'true' : undefined}
  data-scale={scale}
>
  {#if showLane}
    <div class="lane relative" class:coarse style:height="{LANE_PX}px" data-testid="marker-lane">
      {#if status === 'ready'}
        {#each lane as item (item.key)}
          {#if item.leader}
            <span
              class="leader absolute"
              style:left="{item.timeX}px"
              style:width="{item.x - item.timeX}px"
              aria-hidden="true"
            ></span>
          {/if}
          {#if item.tone === 'death'}
            {@const first = item.events[0]}
            <a
              class="marker death absolute"
              class:cluster={item.kind === 'cluster'}
              style:left="{item.x}px"
              href={first.href}
              aria-label={item.kind === 'cluster'
                ? `${chartCopy.deathCluster(item.events.length)}. ${chartCopy.deathAria(first.text)}`
                : chartCopy.deathAria(first.text)}
              data-testid="death-marker"
              data-at={item.startMs}
              onpointerenter={() => (hoveredItem = item)}
              onpointerleave={() => (hoveredItem = null)}
              onfocus={() => (hoveredItem = item)}
              onblur={() => (hoveredItem = null)}
              onclick={(event) => {
                if (onOpenDeath === undefined || event.metaKey || event.ctrlKey || event.shiftKey) return;
                event.preventDefault();
                onOpenDeath(item.events);
              }}
            >
              <span class="face">
                {#if item.kind === 'cluster'}
                  <span class="count numeric">{item.events.length}</span>
                {:else}
                  <Glyph name="skull" size={14} />
                {/if}
              </span>
            </a>
          {:else}
            <span
              class="marker gold absolute"
              class:cluster={item.kind === 'cluster'}
              class:target={item.events[0].kind === 'target-change'}
              style:left="{item.x}px"
              data-testid="lane-marker"
              data-kind={item.events[0].kind}
              title={item.events.map((event) => event.text).join('\n')}
              role="img"
              aria-label={item.events.map((event) => event.text).join('. ')}
              onpointerenter={() => (hoveredItem = item)}
              onpointerleave={() => (hoveredItem = null)}
            >
              <span class="face">
                {#if item.kind === 'cluster'}
                  <span class="count numeric">{item.events.length}</span>
                {:else if item.events[0].iconUrl && !brokenIcons.has(item.events[0].iconUrl)}
                  <img
                    src={item.events[0].iconUrl}
                    alt=""
                    width="18"
                    height="18"
                    onerror={() => markBroken(item.events[0].iconUrl)}
                  />
                {:else}
                  <Glyph name={item.events[0].kind === 'target-change' ? 'target' : 'bolt'} size={14} />
                {/if}
              </span>
            </span>
          {/if}
        {/each}
      {/if}
    </div>
  {/if}

  <!-- svelte-ignore a11y_no_noninteractive_tabindex -->
  <div
    bind:this={host}
    bind:clientWidth={width}
    class="plot relative"
    class:coarse
    style:height="{plotHeight + AXIS_PX}px"
    tabindex={status === 'ready' ? 0 : -1}
    role="group"
    aria-roledescription="chart"
    aria-label={status === 'loading'
      ? chartCopy.loading
      : chartCopy.plotAria(label, current.startMs, current.endMs)}
    aria-busy={status === 'loading' ? 'true' : undefined}
    data-testid="chart-plot"
    onkeydown={onKey}
    onblur={onBlur}
    onpointerdown={onPointerDown}
    onpointermove={onPointerMove}
    onpointerup={onPointerUp}
    onpointercancel={onPointerUp}
    ondblclick={() => zoomed && onWindow(null)}
  >
    <div data-uplot class="h-full w-full" class:invisible={status !== 'ready'}></div>
    {#if status === 'loading'}
      <span
        class="skeleton-block absolute"
        style:left="{gutter}px"
        style:right="12px"
        style:top="8px"
        style:bottom="{AXIS_PX}px"
        aria-hidden="true"
      ></span>
    {:else}
      {#each phaseLabels as { phase, left } (phase.name + phase.startMs)}
        <span class="phase-label absolute" style:left="{left}px" data-testid="phase-band">{phase.name}</span>
      {/each}
      {#if empty}
        <p
          class="text-muted pointer-events-none absolute inset-x-0 text-center text-[14px]"
          style:top="{plotHeight / 2 - 10}px"
          data-testid="chart-empty"
        >
          {emptyCopy}
        </p>
      {/if}
      {#if dragLabel !== null}
        <span
          class="selection absolute"
          style:left="{xOf(dragLabel.from)}px"
          style:width="{xOf(dragLabel.to) - xOf(dragLabel.from)}px"
          style:top="8px"
          style:bottom="{AXIS_PX}px"
          aria-hidden="true"
        ></span>
        <span
          class="drag-label numeric absolute -translate-x-1/2"
          style:left="{(xOf(dragLabel.from) + xOf(dragLabel.to)) / 2}px"
          data-testid="drag-label">{chartCopy.dragLabel(dragLabel.from, dragLabel.to)}</span
        >
      {/if}
      {#if ownMs !== null || chartCursor()?.syncKey === syncKey}
        {@const pillMs = ownMs ?? chartCursor()?.ms ?? 0}
        {#if pillMs >= current.startMs && pillMs <= current.endMs}
          <span
            class="time-pill numeric absolute -translate-x-1/2"
            style:left="{xOf(pillMs)}px"
            style:top="{plotHeight + 10}px"
            data-testid="time-pill">{clockLabel(pillMs)}</span
          >
        {/if}
      {/if}
    {/if}
  </div>

  {#if showEventLine}
    <div
      class="event-line flex items-center gap-1.5 text-[12px]"
      style:height="{EVENT_LINE_PX}px"
      data-testid="event-line"
    >
      {#if lineEvents.length > 0}
        {@const first = lineEvents[0]}
        {#if first.kind === 'death' && first.href}
          <a
            class="death-link flex min-w-0 flex-1 items-center gap-1.5"
            href={first.href}
            data-testid="event-line-death"
            onclick={(event) => {
              if (onOpenDeath === undefined) return;
              event.preventDefault();
              onOpenDeath([first]);
            }}
          >
            <Glyph name="skull" size={14} class="text-death" />
            <span class="truncate">{first.shortText ?? first.text}</span>
            {#if lineEvents.length > 1}<span class="text-muted">+{lineEvents.length - 1}</span>{/if}
            <Glyph name="chevron" size={14} />
          </a>
        {:else}
          <span class="truncate">{first.shortText ?? first.text}</span>
          {#if lineEvents.length > 1}<span class="text-muted">+{lineEvents.length - 1}</span>{/if}
        {/if}
      {:else if !scrubbedOnce}
        <span class="text-muted" data-testid="scrub-hint">{chartCopy.touchHint}</span>
      {/if}
    </div>
  {/if}

  {#if !compact}
    <div style:padding-left="{plotLeft}px" style:padding-right="12px">
      {#if status === 'loading'}
        <div class="pt-1.5">
          <span class="skeleton-block block h-[36px] rounded-[6px]" aria-hidden="true"></span><span
            class="block h-4"
          ></span>
        </div>
      {:else}
        <OverviewBrush
          {domainMs}
          window={current}
          shape={brushShape}
          phaseStarts={phases.filter((phase) => phase.startMs > 0).map((phase) => phase.startMs)}
          deaths={events.filter((event) => event.kind === 'death').map((event) => event.atMs)}
          {onWindow}
          {coarse}
        />
      {/if}
    </div>
    <div class="pb-2"></div>
  {/if}

  <p class="sr-only" aria-live="polite" data-testid="chart-live">{liveText}</p>

  {#if tooltip !== null}
    <ChartTooltip
      open={tooltipOpen}
      x={box.left + crossX}
      y={box.top}
      flip={tooltipFlip}
      title={tooltip.title}
      note={tooltip.note}
      events={tooltip.events}
      rows={tooltip.rows}
      more={tooltip.more}
      width={tooltipWidth}
      threat={shareOf !== undefined}
    />
  {/if}
  <span class="hidden" data-reduced-motion={reducedMotion ? 'true' : undefined}></span>
</div>

<style>
  .lane {
    margin-top: 4px;
  }
  /* A marker is its hit area (24 x 26, 44 x 44 on a coarse pointer); its face is the 20 px tile. */
  .marker {
    top: 0;
    width: 24px;
    height: 26px;
    margin-left: -12px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .lane.coarse .marker {
    top: -9px;
    width: 44px;
    height: 44px;
    margin-left: -22px;
  }
  .face {
    width: 20px;
    height: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .marker.death .face {
    border-radius: 999px;
    background: var(--color-death-soft);
    border: 1px solid color-mix(in srgb, var(--color-death) 60%, transparent);
    color: var(--color-death);
  }
  .marker.death.cluster .face {
    border-color: var(--color-death);
  }
  .marker.death:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 0;
    border-radius: 6px;
  }
  .marker.gold .face {
    border-radius: 4px;
    background: var(--color-raised);
    border: 1px solid color-mix(in srgb, var(--color-gold) 70%, transparent);
    color: var(--color-gold);
  }
  .marker.gold.target .face {
    border-color: color-mix(in srgb, var(--color-text) 50%, transparent);
    color: var(--color-text);
  }
  .marker.gold.cluster .face {
    border-radius: 999px;
    border-color: var(--color-gold);
  }
  .marker img {
    width: 18px;
    height: 18px;
    border-radius: 3px;
  }
  .count {
    font-size: 11px;
    font-weight: 700;
    color: #ffb4b0;
  }
  .marker.gold .count {
    color: var(--color-gold);
  }
  .leader {
    top: 12px;
    height: 1px;
    background: color-mix(in srgb, var(--color-gold) 60%, transparent);
  }
  .plot {
    touch-action: pan-y;
    outline: none;
    overflow-x: clip;
  }
  .plot:focus-visible {
    box-shadow: inset 0 0 0 2px var(--color-gold);
    border-radius: 6px;
  }
  .phase-label {
    top: 12px;
    height: 16px;
    padding: 0 5px;
    border-radius: 3px;
    background: rgba(13, 15, 21, 0.9);
    color: var(--color-muted);
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.12em;
    line-height: 16px;
    text-transform: uppercase;
    pointer-events: none;
    white-space: nowrap;
  }
  .selection {
    background: rgba(229, 185, 85, 0.13);
    border-left: 1px solid var(--color-gold);
    border-right: 1px solid var(--color-gold);
    pointer-events: none;
  }
  .drag-label {
    top: 34px;
    padding: 3px 8px;
    border-radius: 6px;
    background: #0d0f15;
    border: 1px solid color-mix(in srgb, var(--color-gold) 60%, transparent);
    color: var(--color-strong);
    font-size: 12px;
    font-weight: 600;
    pointer-events: none;
    white-space: nowrap;
  }
  .time-pill {
    width: 40px;
    height: 17px;
    border-radius: 4px;
    background: var(--color-text);
    color: var(--color-bg);
    font-size: 11px;
    font-weight: 700;
    line-height: 17px;
    text-align: center;
    pointer-events: none;
  }
  .event-line {
    padding: 0 12px;
    color: var(--color-text);
  }
  .death-link {
    color: #ffb4b0;
  }
  /* uPlot's own layout rules (its stylesheet, trimmed to what this chart uses). */
  .tsc :global(.uplot) {
    font-family: inherit;
    line-height: 1;
    width: 100%;
  }
  .tsc :global(.uplot *) {
    box-sizing: border-box;
  }
  .tsc :global(.u-wrap) {
    position: relative;
    user-select: none;
  }
  .tsc :global(.u-over),
  .tsc :global(.u-under) {
    position: absolute;
  }
  .tsc :global(.u-under) {
    overflow: hidden;
  }
  .tsc :global(.uplot canvas) {
    display: block;
    position: relative;
    width: 100%;
    height: 100%;
  }
  .tsc :global(.u-axis) {
    position: absolute;
  }
  .tsc :global(.u-select) {
    display: none;
  }
  .tsc :global(.u-cursor-x) {
    position: absolute;
    left: 0;
    top: 0;
    height: 100%;
    pointer-events: none;
    will-change: transform;
    border-right: 1px dashed color-mix(in srgb, var(--color-text) 70%, transparent);
  }
  .tsc :global(.u-cursor-pt) {
    position: absolute;
    top: 0;
    left: 0;
    border-radius: 50%;
    border: 0 solid;
    pointer-events: none;
    will-change: transform;
    background-clip: padding-box !important;
  }
  .tsc :global(.u-cursor-x.u-off),
  .tsc :global(.u-cursor-pt.u-off),
  .tsc :global(.u-axis.u-off) {
    display: none;
  }
</style>
