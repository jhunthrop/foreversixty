// web/src/lib/sim/spread.ts
// The DPS spread (spec 2026-11-04-sim-results 3.6): the engine's own 1 DPS bins grouped into at
// most 40 bars, and the 16th and 84th percentiles read off the real counts. Nothing here draws a
// shape the run did not produce.
import type { DpsHistogram } from './types';

export const MAX_BARS = 40;
const LOW_PERCENTILE = 0.16;
const HIGH_PERCENTILE = 0.84;

export interface SpreadBar {
  /** The bar's first DPS value. */
  low: number;
  /** The bar's last DPS value. */
  high: number;
  count: number;
  /** Inside the 16th to 84th percentile. */
  central: boolean;
}

export interface Spread {
  bars: SpreadBar[];
  peak: number;
  total: number;
  /** The 16th and 84th percentile DPS. */
  low: number;
  high: number;
}

/** The DPS below which `share` of the runs fall, from the raw bins. */
export function percentileOf(histogram: DpsHistogram, share: number): number {
  const total = histogram.counts.reduce((sum, count) => sum + count, 0);
  const wanted = total * share;
  let seen = 0;
  for (const [index, count] of histogram.counts.entries()) {
    seen += count;
    if (seen >= wanted) return histogram.start + index * histogram.bin_dps;
  }
  return histogram.start + (histogram.counts.length - 1) * histogram.bin_dps;
}

/** The bins grouped into `ceil(bins / 40)` DPS buckets. Null when there is nothing to draw. */
export function spreadOf(histogram: DpsHistogram | undefined): Spread | null {
  if (histogram === undefined || histogram.counts.length === 0) return null;
  const total = histogram.counts.reduce((sum, count) => sum + count, 0);
  if (total === 0) return null;
  const step = Math.ceil(histogram.counts.length / MAX_BARS);
  const low = percentileOf(histogram, LOW_PERCENTILE);
  const high = percentileOf(histogram, HIGH_PERCENTILE);
  const bars: SpreadBar[] = [];
  for (let index = 0; index < histogram.counts.length; index += step) {
    const barLow = histogram.start + index * histogram.bin_dps;
    const barHigh = barLow + step * histogram.bin_dps - 1;
    const count = histogram.counts.slice(index, index + step).reduce((sum, value) => sum + value, 0);
    bars.push({ low: barLow, high: barHigh, count, central: barHigh >= low && barLow <= high });
  }
  return { bars, peak: Math.max(...bars.map((bar) => bar.count)), total, low, high };
}
