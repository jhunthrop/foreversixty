// web/src/lib/sim/estimate.ts
// Pooling shard distributions into one, for the DPS figure that updates while the sim runs.
//
// An average of the shards' averages is only right when the shards are equal-sized, and a
// mean of their standard deviations is never right. Both are pooled properly here from the
// sum and the sum of squares each shard implies (n * (s^2 + m^2)), which is exactly what
// the wasm's own simCombine does at the end -- run.test.ts asserts the two agree, which is
// what lets the page show this number a wasm round-trip earlier.
import { MARGIN_BELOW_THRESHOLD } from './copy';
import type { Estimate } from './types';

export const EMPTY_ESTIMATE: Estimate = { mean: 0, stddev: 0, error: 0, min: 0, max: 0 };

export interface ShardProgress {
  shard: number;
  iterationsDone: number;
  estimate: Estimate;
}

export interface PooledEstimate {
  estimate: Estimate;
  iterationsDone: number;
}

export function combineEstimate(shards: readonly ShardProgress[]): PooledEstimate {
  let n = 0;
  let sum = 0;
  let sumSquares = 0;
  let min = Number.POSITIVE_INFINITY;
  let max = Number.NEGATIVE_INFINITY;

  for (const shard of shards) {
    if (shard.iterationsDone <= 0) continue;
    const { mean, stddev } = shard.estimate;
    n += shard.iterationsDone;
    sum += mean * shard.iterationsDone;
    sumSquares += shard.iterationsDone * (stddev * stddev + mean * mean);
    if (shard.estimate.min < min) min = shard.estimate.min;
    if (shard.estimate.max > max) max = shard.estimate.max;
  }

  if (n === 0) return { estimate: EMPTY_ESTIMATE, iterationsDone: 0 };

  const mean = sum / n;
  const variance = Math.max(0, sumSquares / n - mean * mean);
  const stddev = Math.sqrt(variance);
  return { estimate: { mean, stddev, error: stddev / Math.sqrt(n), min, max }, iterationsDone: n };
}

export function confidenceBand(estimate: Estimate): number {
  return 1.96 * estimate.error;
}

/**
 * Every "±" figure on the page -- the run headline, a saved sim, the planner strip, a
 * bulk tool's equipped-set line, the Droptimizer gain column, the details card and the OG
 * unfurl -- renders its margin through this one function, so none of them can print a
 * different number for the same run (tank-sim review.md D3, healer-sim review.md Minor).
 *
 * A margin below 10 reads with one decimal ("± 4.2"); 10 or more reads as an integer, as
 * it always has. A non-zero margin never rounds away to "0": one that would round to 0.0
 * at one decimal reads "< 0.1" instead -- the controller's ruling, chosen over a sliding
 * decimal count because one rule reads the same everywhere and never implies false
 * precision on a noisy run. A genuinely zero margin (single iteration, no variance) still
 * reads "0".
 */
export function formatMargin(value: number): string {
  if (value === 0) return '0';
  const oneDecimal = value.toFixed(1);
  if (oneDecimal === '0.0') return MARGIN_BELOW_THRESHOLD;
  // Decided off the *rounded* value, not the raw one: 9.96 and 9.99 both round to "10.0"
  // at one decimal, and must join the integer branch the same way an actual 10 does,
  // rather than printing a "10.0" no value ever should.
  return parseFloat(oneDecimal) >= 10 ? Math.round(value).toLocaleString('en-US') : oneDecimal;
}
