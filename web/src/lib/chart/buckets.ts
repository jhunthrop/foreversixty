// web/src/lib/chart/buckets.ts
// How finely a rate chart draws a fight. The summary's series are one value per second (the logs
// engine's Options.Bucket, logs/engine/summary/summary.go); the chart averages them into buckets
// whose width follows the fight's length: a 10 s pull draws every second, a long fight keeps the
// 5 s average, and no series draws more than MAX_POINTS points.
import { BUCKET_MS } from '../report/window';

/** The widths a point may take, in seconds. */
export const BUCKET_STEPS_S = [1, 2, 5, 10, 15, 30, 60] as const;
/** The widest a point is while the fight has room: the 5 s average long fights have always drawn. */
export const DEFAULT_BUCKET_S = 5;
/** The fewest points a fight draws before its bucket narrows below the default. */
export const MIN_POINTS = 30;
/** The most points a series draws: a very long fight widens its bucket past the default. */
export const MAX_POINTS = 300;
/** The finest label step: on a one-second chart, a label every other point. */
export const MIN_LABEL_STEP_S = 2;

export interface ChartBucket {
  /** Seconds per point. */
  seconds: number;
  /** The time axis' finest label step, in seconds (time-axis.ts): never finer than the points. */
  labelFloorS: number;
}

/**
 * The widest step up to the default that still gives MIN_POINTS points (one second when none
 * does), then widened until the fight fits in MAX_POINTS.
 */
export function chartBucket(durationMs: number): ChartBucket {
  const spanS = Number.isFinite(durationMs) && durationMs > 0 ? Math.ceil(durationMs / BUCKET_MS) : 0;
  const points = (step: number): number => Math.ceil(spanS / step);
  const fine = [...BUCKET_STEPS_S]
    .reverse()
    .find((step) => step <= DEFAULT_BUCKET_S && points(step) >= MIN_POINTS);
  const start = fine ?? BUCKET_STEPS_S[0];
  const seconds =
    BUCKET_STEPS_S.find((step) => step >= start && points(step) <= MAX_POINTS) ?? BUCKET_STEPS_S.at(-1)!;
  return { seconds, labelFloorS: Math.max(MIN_LABEL_STEP_S, seconds) };
}

/**
 * Each bucket's mean per second over the fight's `spanS` seconds; the last bucket is averaged over
 * the seconds it has. A series that stops at its last busy second reads zero past it.
 */
export function bucketMeans(values: readonly number[], seconds: number, spanS: number): number[] {
  return Array.from({ length: Math.ceil(spanS / seconds) }, (_, index) => {
    const from = index * seconds;
    const to = Math.min(from + seconds, spanS);
    let sum = 0;
    for (let second = from; second < to; second += 1) sum += values[second] ?? 0;
    return sum / (to - from);
  });
}

/** A line's value at a moment: the bucket the moment falls in. */
export function valueAt<T>(line: { values: readonly T[]; stepMs?: number }, ms: number): T | undefined {
  return line.values[Math.floor(ms / (line.stepMs ?? BUCKET_MS))];
}
