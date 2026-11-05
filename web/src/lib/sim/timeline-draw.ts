// web/src/lib/sim/timeline-draw.ts
// The timeline's plot as one SVG string (spec 2026-11-04-sim-results 3.7 and 8): no chart library
// and no per-tick DOM node a framework must track, so a mount stays cheap. Nothing from the
// stored result reaches this string as text: every value is formatted as a number and every
// colour and label is a constant here, which is why `{@html}` over it is safe for a result a
// stranger saved.
import { clockLabel } from '../chart/time-axis';
import {
  clipSegments,
  mergeTicks,
  ROW_HEIGHT,
  stacksAt,
  type BarLane,
  type DrawnTick,
  type PlotRow,
  type ResourceLine,
  type TickLane,
  type TimelineEvent,
  type TimelineModel,
  type ZoomWindow,
} from './timeline';
import { timelineCopy } from './timeline-copy';

/** Axis steps (spec 3.7): 15 s on a wide plot, 30 s under 900 px of plot, 5 s in Close up. */
const CLOSE_UP_STEP_S = 5;
const WHOLE_STEP_S = 15;
const NARROW_WHOLE_STEP_S = 30;
const NARROW_PLOT_PX = 900;

const AXIS_LABEL_Y = 20;
/** A label closer than this to the plot's edge would be cut: the gridline stays, the label goes. */
const AXIS_LABEL_MARGIN_PX = 18;
const GRID_TOP = ROW_HEIGHT.axis - 6;
const RESOURCE_INSET = 8;

const COLOUR = {
  cooldown: 'var(--color-text)',
  crit: 'var(--color-gold-light)',
  missed: 'var(--color-muted)',
  buff: 'var(--color-text)',
  debuff: 'var(--color-gold-deep)',
  grid: 'var(--color-line-soft)',
  pullLine: 'var(--color-line-warm-strong)',
  capLine: 'var(--color-line-warm-strong)',
  gold: 'var(--color-gold)',
  label: 'var(--color-muted)',
  ember: 'var(--color-ember)',
  stack: 'var(--color-strong)',
} as const;
const WASH = 'rgba(214,110,40,.07)';
const LINE_AREA = 'rgba(229,185,85,.12)';
const BUFF_OPACITY = 0.35;
const DEBUFF_OPACITY = 0.78;
const SWING_OPACITY = 0.8;
const GLANCE_OPACITY = 0.6;

/** The smallest bar that still reads, and the width from which it prints its stack count. */
const MIN_BAR_PX = 2;
const STACK_LABEL_MIN_PX = 18;

export interface PlotView {
  window: ZoomWindow;
  width: number;
  closeUp: boolean;
}

const round = (value: number): string => Number(value).toFixed(1);

export function pxPerMs(view: PlotView): number {
  const span = view.window.endMs - view.window.startMs;
  return span <= 0 ? 0 : view.width / span;
}

export function xOf(atMs: number, view: PlotView): number {
  return (atMs - view.window.startMs) * pxPerMs(view);
}

export function timeOf(x: number, view: PlotView): number {
  const scale = pxPerMs(view);
  return scale === 0 ? view.window.startMs : view.window.startMs + x / scale;
}

export function axisStepSeconds(view: Pick<PlotView, 'closeUp' | 'width'>): number {
  if (view.closeUp) return CLOSE_UP_STEP_S;
  return view.width < NARROW_PLOT_PX ? NARROW_WHOLE_STEP_S : WHOLE_STEP_S;
}

/** The labelled times in view, in ms, on whole steps. The pull is drawn on its own. */
export function axisTimes(view: PlotView): number[] {
  const stepMs = axisStepSeconds(view) * 1000;
  const times: number[] = [];
  for (let at = Math.ceil(view.window.startMs / stepMs) * stepMs; at < view.window.endMs; at += stepMs) {
    times.push(at);
  }
  return times;
}

interface TickShape {
  width: number;
  height: number;
  fill: string;
  opacity: number;
  outlined: boolean;
}

const MISSED = new Set(['miss', 'dodge', 'parry']);

/** A tick's look (spec 3.7, Ticks, with the approved mock's swing sizes). */
export function tickShape(lane: TickLane, event: TimelineEvent, merged = false): TickShape {
  const swing = lane.group === 'swings';
  const base = {
    width: swing ? 2 : 3,
    height: swing ? 10 : 14,
    fill: lane.schoolToken === null ? COLOUR.cooldown : `var(${lane.schoolToken})`,
    opacity: swing ? SWING_OPACITY : 1,
    outlined: false,
  };
  const outcome = lane.periodic ? undefined : event.outcome;
  const shape =
    outcome === 'crit'
      ? { ...base, width: base.width + 1, height: base.height + 4, fill: COLOUR.crit }
      : outcome !== undefined && MISSED.has(outcome)
        ? { ...base, fill: COLOUR.missed, outlined: true, opacity: 1 }
        : outcome === 'glance'
          ? { ...base, opacity: GLANCE_OPACITY }
          : base;
  return merged ? { ...shape, width: Math.max(shape.width, 3) } : shape;
}

function tickRect(x: number, laneY: number, laneHeight: number, shape: TickShape): string {
  const left = round(x - shape.width / 2);
  const top = round(laneY + (laneHeight - shape.height) / 2);
  const paint = shape.outlined
    ? `fill="none" stroke="${shape.fill}" stroke-width="1"`
    : `fill="${shape.fill}"${shape.opacity === 1 ? '' : ` opacity="${shape.opacity}"`}`;
  return `<rect x="${left}" y="${top}" width="${shape.width}" height="${shape.height}" rx="1" ${paint}/>`;
}

/** Every tick lane's drawn ticks for this view: the markup draws them, the pointer hit-tests them. */
export function drawnTicks(rows: readonly PlotRow[], view: PlotView): Map<string, DrawnTick[]> {
  const scale = pxPerMs(view);
  const drawn = new Map<string, DrawnTick[]>();
  for (const row of rows) {
    if (row.kind === 'lane' && row.lane.kind === 'ticks') {
      drawn.set(row.lane.id, mergeTicks(row.lane.events, view.window, scale));
    }
  }
  return drawn;
}

function tickLaneMarkup(
  lane: TickLane,
  ticks: readonly DrawnTick[],
  y: number,
  height: number,
  view: PlotView,
): string {
  let marks = '';
  if (view.closeUp) {
    // The queue press: a hollow tick, in Close up only, so Whole fight never reads as twice the
    // presses (spec 3.7).
    for (const tick of mergeTicks(lane.queued, view.window, pxPerMs(view))) {
      marks += tickRect(tick.x, y, height, { ...tickShape(lane, tick.event), outlined: true, opacity: 1 });
    }
  }
  for (const tick of ticks) marks += tickRect(tick.x, y, height, tickShape(lane, tick.event, tick.count > 1));
  return marks;
}

function barLaneMarkup(lane: BarLane, y: number, height: number, view: PlotView): string {
  const isWindow = lane.group === 'window';
  const barHeight = isWindow ? 6 : 10;
  const fill = lane.group === 'target' ? COLOUR.debuff : COLOUR.buff;
  const opacity = lane.group === 'target' ? DEBUFF_OPACITY : BUFF_OPACITY;
  const top = y + (height - barHeight) / 2;
  let marks = '';
  for (const [start, end] of clipSegments(lane.segments, view.window)) {
    const left = xOf(start, view);
    const width = Math.max(MIN_BAR_PX, xOf(end, view) - left);
    marks += `<rect x="${round(left)}" y="${round(top)}" width="${round(width)}" height="${barHeight}" rx="3" fill="${fill}" opacity="${opacity}"/>`;
    const stacks = stacksAt(lane.stacks, start);
    if (stacks > 1 && width >= STACK_LABEL_MIN_PX) {
      marks += `<text x="${round(left + 4)}" y="${round(top + 9)}" font-size="11" font-weight="600" fill="${COLOUR.stack}">${Math.round(stacks)}</text>`;
    }
  }
  return marks;
}

/** The resource as a step line: each value held until the next event, then the area under it. */
export function resourcePaths(
  line: ResourceLine,
  rowY: number,
  rowHeight: number,
  view: PlotView,
): { line: string; area: string; capY: number; floorY: number } {
  const floorY = rowY + rowHeight - RESOURCE_INSET;
  const capY = rowY + RESOURCE_INSET;
  const yOf = (value: number): number =>
    line.cap <= 0 ? floorY : floorY - (value / line.cap) * (floorY - capY);
  const { startMs, endMs } = view.window;
  let path = '';
  let held: number | null = null;
  let firstX: number | null = null;
  for (const [at, value] of line.points) {
    if (at <= startMs) {
      held = value;
      continue;
    }
    if (at > endMs) break;
    const x = round(xOf(at, view));
    if (path === '') {
      firstX = held === null ? xOf(at, view) : 0;
      path = `M${round(firstX)},${round(yOf(held ?? value))}`;
    }
    path += `H${x}V${round(yOf(value))}`;
    held = value;
  }
  if (path === '' && held !== null) {
    firstX = 0;
    path = `M0,${round(yOf(held))}`;
  }
  if (path === '') return { line: '', area: '', capY, floorY };
  // The line runs to the window's edge: the last value holds until the fight ends.
  path += `H${round(view.width)}`;
  return { line: path, area: `${path}V${round(floorY)}H${round(firstX ?? 0)}Z`, capY, floorY };
}

function resourceMarkup(line: ResourceLine, y: number, height: number, view: PlotView): string {
  const paths = resourcePaths(line, y, height, view);
  const cap = `<line x1="0" x2="${round(view.width)}" y1="${round(paths.capY)}" y2="${round(paths.capY)}" stroke="${COLOUR.capLine}" stroke-dasharray="2 4"/>`;
  if (paths.line === '') return cap;
  return (
    cap +
    `<path d="${paths.area}" fill="${LINE_AREA}"/>` +
    `<path d="${paths.line}" fill="none" stroke="${COLOUR.gold}" stroke-width="1.5"/>`
  );
}

function axisMarkup(height: number, view: PlotView): string {
  let marks = '';
  for (const at of axisTimes(view)) {
    if (at === 0) continue;
    const position = xOf(at, view);
    const x = round(position);
    marks += `<line x1="${x}" x2="${x}" y1="${GRID_TOP}" y2="${height}" stroke="${COLOUR.grid}"/>`;
    if (position < AXIS_LABEL_MARGIN_PX || position > view.width - AXIS_LABEL_MARGIN_PX) continue;
    marks += `<text x="${x}" y="${AXIS_LABEL_Y}" text-anchor="middle" font-size="13" font-weight="600" fill="${COLOUR.label}">${clockLabel(at)}</text>`;
  }
  if (view.window.startMs <= 0 && view.window.endMs > 0) {
    const x = xOf(0, view);
    marks += `<line x1="${round(x)}" x2="${round(x)}" y1="${GRID_TOP}" y2="${height}" stroke="${COLOUR.pullLine}" stroke-dasharray="3 3"/>`;
    marks += `<text x="${round(x + 4)}" y="${AXIS_LABEL_Y}" font-size="13" font-weight="600" fill="${COLOUR.gold}">${timelineCopy.axisPull}</text>`;
  }
  return marks;
}

function washMarkup(executeAtMs: number | null, height: number, view: PlotView): string {
  if (executeAtMs === null || executeAtMs >= view.window.endMs) return '';
  const from = Math.max(executeAtMs, view.window.startMs);
  const x = xOf(from, view);
  const wash = `<rect x="${round(x)}" y="${ROW_HEIGHT.axis}" width="${round(view.width - x)}" height="${height - ROW_HEIGHT.axis}" fill="${WASH}"/>`;
  return (
    wash +
    `<text x="${round(x + 6)}" y="${ROW_HEIGHT.axis + 14}" font-size="11" font-weight="700" letter-spacing="1.54" fill="${COLOUR.ember}">${timelineCopy.executeWash}</text>`
  );
}

/** The whole plot, axis to resource line, as SVG children. */
export function plotMarkup(
  model: Pick<TimelineModel, 'executeAtMs'>,
  rows: readonly PlotRow[],
  height: number,
  ticks: ReadonlyMap<string, readonly DrawnTick[]>,
  view: PlotView,
): string {
  let marks = washMarkup(model.executeAtMs, height, view) + axisMarkup(height, view);
  rows.forEach((row, index) => {
    if (row.kind === 'resource') marks += resourceMarkup(row.line, row.y, row.height, view);
    if (row.kind !== 'lane') return;
    const inner =
      row.lane.kind === 'ticks'
        ? tickLaneMarkup(row.lane, ticks.get(row.lane.id) ?? [], row.y, row.height, view)
        : barLaneMarkup(row.lane, row.y, row.height, view);
    // The row's index, never its action key: the markup carries no text from the stored result.
    marks += `<g data-row="${index}">${inner}</g>`;
  });
  return marks;
}

/** The whole fight's resource line for the brush, in a `width` x `height` strip. */
export function brushLine(
  line: ResourceLine | null,
  model: Pick<TimelineModel, 'startMs' | 'endMs'>,
  width: number,
  height: number,
): string {
  if (line === null || line.cap <= 0) return '';
  const span = model.endMs - model.startMs;
  if (span <= 0) return '';
  const inset = 8;
  const usable = height - inset * 2;
  return line.points
    .map(
      ([at, value], index) =>
        `${index === 0 ? 'M' : 'L'}${round(((at - model.startMs) / span) * width)},${round(inset + usable - (value / line.cap) * usable)}`,
    )
    .join('');
}
