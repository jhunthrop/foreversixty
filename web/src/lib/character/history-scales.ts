// web/src/lib/character/history-scales.ts
// The chart's pure arithmetic (experience spec 2026-10-24-character-history 3.5): the plot
// rectangle, the y domains and their labels, the proportional-to-time x axis and its tick
// rule. No data types here; history-chart.ts feeds it numbers.
import { SCORE_MAX, TIER_FROM } from '../score/score';

export interface ChartSize {
  width: number;
  height: number;
  /** Below 640 px: narrower y gutter, three x labels. */
  phone: boolean;
}

export interface PlotRect {
  left: number;
  top: number;
  width: number;
  height: number;
}

export interface Domain {
  min: number;
  max: number;
}

export type TickAnchor = 'start' | 'middle' | 'end';

export interface AxisTick {
  /** Milliseconds since the epoch. */
  time: number;
  /** Pixels from the plot's left edge. */
  x: number;
  anchor: TickAnchor;
}

const MARGIN_LEFT = 40;
const MARGIN_LEFT_PHONE = 36;
const MARGIN_RIGHT = 14;
const MARGIN_TOP = 12;
const MARGIN_BOTTOM = 26;

const DOMAIN_STEP = 100;
const SCORE_PAD = 60;
export const MIN_SCORE_SPAN = 400;
const LABEL_EVERY_TIER_ABOVE_SPAN = 1200;
const FINE_LABEL_SPAN = 600;
const COARSE_STEP = 200;

/** The colour ladder's bounds: a parse chart's floor is the rung below its lowest parse. */
const PARSE_FLOORS: readonly { below: number; floor: number }[] = [
  { below: 30, floor: 0 },
  { below: 55, floor: 25 },
  { below: 80, floor: 50 },
];
const PARSE_TOP_FLOOR = 75;
export const PARSE_CEILING = 100;
const PARSE_RUNGS = [25, 50, 75, 100];

export const MIN_LABEL_GAP = 96;
const PHONE_LABELS = 3;
const DESKTOP_LABELS = 6;

export function plotRect(size: ChartSize): PlotRect {
  const left = size.phone ? MARGIN_LEFT_PHONE : MARGIN_LEFT;
  return {
    left,
    top: MARGIN_TOP,
    width: Math.max(0, size.width - left - MARGIN_RIGHT),
    height: Math.max(0, size.height - MARGIN_TOP - MARGIN_BOTTOM),
  };
}

function requireValues(values: readonly number[]): void {
  if (values.length === 0) throw new Error('a chart domain needs at least one value');
}

/** floor((min - 60) / 100) * 100 to ceil((max + 60) / 100) * 100, at least 400 wide, inside 0 to 2100. */
export function scoreDomain(values: readonly number[]): Domain {
  requireValues(values);
  let min = Math.floor((Math.min(...values) - SCORE_PAD) / DOMAIN_STEP) * DOMAIN_STEP;
  let max = Math.ceil((Math.max(...values) + SCORE_PAD) / DOMAIN_STEP) * DOMAIN_STEP;
  if (max - min < MIN_SCORE_SPAN) {
    min = Math.floor(((min + max) / 2 - MIN_SCORE_SPAN / 2) / DOMAIN_STEP) * DOMAIN_STEP;
    max = min + MIN_SCORE_SPAN;
  }
  // Clamping shifts the window rather than shrinking it, so the 400 minimum survives.
  if (min < 0) return { min: 0, max: max - min };
  if (max > SCORE_MAX) return { min: min - (max - SCORE_MAX), max: SCORE_MAX };
  return { min, max };
}

/** The ladder's own rungs (0, 25, 50, 75) to 100, never the exact data, so it does not rescale as points arrive. */
export function parseDomain(values: readonly number[]): Domain {
  requireValues(values);
  const lowest = Math.min(...values);
  const rung = PARSE_FLOORS.find(({ below }) => lowest < below);
  return { min: rung?.floor ?? PARSE_TOP_FLOOR, max: PARSE_CEILING };
}

function multiplesIn(domain: Domain, step: number): number[] {
  const first = Math.ceil(domain.min / step) * step;
  const out: number[] = [];
  for (let value = first; value <= domain.max; value += step) out.push(value);
  return out;
}

/** Tier boundaries inside a wide domain, else every 200 (every 100 up to a 600 span). */
export function scoreTicks(domain: Domain): number[] {
  const span = domain.max - domain.min;
  if (span > LABEL_EVERY_TIER_ABOVE_SPAN) {
    return Object.values(TIER_FROM).filter((from) => from >= domain.min && from <= domain.max);
  }
  return multiplesIn(domain, span <= FINE_LABEL_SPAN ? DOMAIN_STEP : COARSE_STEP);
}

/** 25, 50, 75 and 100 inside the domain, and 0 when the domain starts there. */
export function parseTicks(domain: Domain): number[] {
  const rungs = PARSE_RUNGS.filter((value) => value >= domain.min && value <= domain.max);
  return domain.min === 0 ? [0, ...rungs] : rungs;
}

/** Maps a domain value onto [from, to]; a zero-width domain maps to `from`. */
export function linearScale(domain: Domain, from: number, to: number): (value: number) => number {
  const span = domain.max - domain.min;
  return (value) => (span === 0 ? from : from + ((value - domain.min) / span) * (to - from));
}

/** y grows downward: the domain's max sits at the plot's top. */
export function yScale(domain: Domain, plot: PlotRect): (value: number) => number {
  return linearScale(domain, plot.top + plot.height, plot.top);
}

/**
 * The x labels: three on a phone (first, middle, last), up to six otherwise, evenly spaced in
 * time and dropped one at a time until they are at least 96 px apart. The first is start-aligned
 * to the plot's left edge, the last end-aligned to its right edge, the rest centred.
 */
export function axisTicks(first: number, last: number, plotWidth: number, phone: boolean): AxisTick[] {
  if (last <= first) return [{ time: first, x: 0, anchor: 'start' }];
  let count = phone ? PHONE_LABELS : DESKTOP_LABELS;
  while (count > 2 && plotWidth / (count - 1) < MIN_LABEL_GAP) count -= 1;
  return Array.from({ length: count }, (_, index) => {
    const share = index / (count - 1);
    const anchor: TickAnchor = index === 0 ? 'start' : index === count - 1 ? 'end' : 'middle';
    return { time: first + (last - first) * share, x: plotWidth * share, anchor };
  });
}
