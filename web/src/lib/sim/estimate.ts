// web/src/lib/sim/estimate.ts
// Pooling shard distributions into one, for the DPS figure that updates while the sim runs.
//
// An average of the shards' averages is only right when the shards are equal-sized, and a
// mean of their standard deviations is never right. Both are pooled properly here from the
// sum and the sum of squares each shard implies (n * (s^2 + m^2)), which is exactly what
// the wasm's own simCombine does at the end -- run.test.ts asserts the two agree, which is
// what lets the page show this number a wasm round-trip earlier.
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
