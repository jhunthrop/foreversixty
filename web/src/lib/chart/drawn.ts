// web/src/lib/chart/drawn.ts
// Who a chart draws and what each line's numbers are (controls and charts spec §B5.3, D5). Pure:
// the report's adapters (rate.ts, threat.ts) call these and the unit tests pin them.
import { BUCKET_MS, sliceSeries, type TimeWindow } from '../report/window';

/** The default number of players a chart draws, before the always-drawn ones. */
export const DEFAULT_TOP_N = 8;

/** The rate charts' smoothing: a trailing average over this many seconds. */
export const AVERAGE_SECONDS = 5;

/**
 * A trailing average: each second is the mean of itself and the seconds before it, up to
 * `seconds` of them. A run of zeros brings it to exactly zero and keeps it there, so a phase
 * where the boss takes no damage draws flat at zero, never a gap and never a slope across it.
 */
export function trailingAverage(values: readonly number[], seconds = AVERAGE_SECONDS): number[] {
  let running = 0;
  return values.map((value, index) => {
    running += value;
    if (index >= seconds) running -= values[index - seconds];
    return running / Math.min(seconds, index + 1);
  });
}

/** A running total: the value at a second is the standing then. */
export function cumulative(values: readonly number[]): number[] {
  let running = 0;
  return values.map((value) => (running += value));
}

/** The window's total over its seconds: DPS, HPS or DTPS over the window. */
export function windowRate(values: readonly number[], window: TimeWindow): number {
  const seconds = (window.endMs - window.startMs) / 1000;
  if (seconds <= 0) return 0;
  return sliceSeries([...values], window).reduce((sum, value) => sum + value, 0) / seconds;
}

/** A cumulative line's value at the window's last second (its last value when the line is shorter). */
export function standingAt(running: readonly number[], window: TimeWindow): number {
  if (running.length === 0) return 0;
  const index = Math.min(running.length, Math.ceil(window.endMs / BUCKET_MS)) - 1;
  return index < 0 ? 0 : running[index];
}

/** Every figure rounded first, then added: the head's raid figure is the sum a raider makes. */
export function roundedSum(figures: readonly number[]): number {
  return figures.reduce((sum, figure) => sum + Math.round(figure), 0);
}

export function median(values: readonly number[]): number {
  if (values.length === 0) return 0;
  const sorted = [...values].sort((a, b) => a - b);
  const middle = Math.floor(sorted.length / 2);
  return sorted.length % 2 === 1 ? sorted[middle] : (sorted[middle - 1] + sorted[middle]) / 2;
}

/** The per-second median across several lines. */
export function medianSeries(lines: readonly (readonly number[])[]): number[] {
  const length = lines.reduce((longest, line) => Math.max(longest, line.length), 0);
  return Array.from({ length }, (_, index) => median(lines.map((line) => line[index] ?? 0)));
}

export interface DrawCandidate {
  id: string;
  figure: number;
}

export interface DrawSelection {
  candidates: readonly DrawCandidate[];
  /** Always drawn whenever they are in the fight. */
  pinned: ReadonlySet<string>;
  /** The Players control's choice; null keeps the default top N. */
  chosen: ReadonlySet<string> | null;
  topN?: number;
}

/** The ids drawn, highest figure first: the top N (or the chosen set) plus the pinned players. */
export function selectDrawn({ candidates, pinned, chosen, topN = DEFAULT_TOP_N }: DrawSelection): string[] {
  const ranked = [...candidates].sort((a, b) => b.figure - a.figure);
  const base = chosen === null ? ranked.slice(0, topN).map((entry) => entry.id) : [...chosen];
  const wanted = new Set([...base, ...pinned]);
  return ranked.filter((entry) => wanted.has(entry.id)).map((entry) => entry.id);
}

export interface DeadSpan {
  startMs: number;
  endMs: number;
}

/** The line stops at a death (the death's own second is its last point) and resumes after a raise. */
export function stopAtDeath(values: readonly number[], spans: readonly DeadSpan[]): (number | null)[] {
  if (spans.length === 0) return [...values];
  return values.map((value, index) => {
    const at = index * BUCKET_MS;
    return spans.some((span) => at > span.startMs && at <= span.endMs) ? null : value;
  });
}
