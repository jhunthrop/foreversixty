import { describe, expect, it } from 'vitest';
import {
  EMPTY_ESTIMATE,
  combineEstimate,
  confidenceBand,
  formatMargin,
  type ShardProgress,
} from './estimate';

/** Turns raw samples into the shape a shard reports, so the test can check against truth. */
function shardsFromSamples(groups: number[][]): ShardProgress[] {
  return groups.map((values, shard) => {
    const n = values.length;
    const mean = values.reduce((a, b) => a + b, 0) / n;
    const variance = values.reduce((a, b) => a + (b - mean) ** 2, 0) / n;
    const stddev = Math.sqrt(variance);
    return {
      shard,
      iterationsDone: n,
      estimate: {
        mean,
        stddev,
        error: stddev / Math.sqrt(n),
        min: Math.min(...values),
        max: Math.max(...values),
      },
    };
  });
}

describe('combineEstimate', () => {
  it('matches the statistics of the pooled sample, not an average of averages', () => {
    const all = [900, 1100, 1000, 1200, 800, 1050, 950, 1300, 700, 1000];
    const combined = combineEstimate(shardsFromSamples([all.slice(0, 4), all.slice(4, 7), all.slice(7)]));
    const mean = all.reduce((a, b) => a + b, 0) / all.length;
    const variance = all.reduce((a, b) => a + (b - mean) ** 2, 0) / all.length;
    expect(combined.iterationsDone).toBe(10);
    expect(combined.estimate.mean).toBeCloseTo(mean, 9);
    expect(combined.estimate.stddev).toBeCloseTo(Math.sqrt(variance), 9);
    expect(combined.estimate.error).toBeCloseTo(Math.sqrt(variance / 10), 9);
    expect(combined.estimate.min).toBe(700);
    expect(combined.estimate.max).toBe(1300);
  });

  it('ignores shards that have not reported an iteration yet', () => {
    const [only] = shardsFromSamples([[1000, 1000, 1000]]);
    const idle: ShardProgress = { shard: 1, iterationsDone: 0, estimate: EMPTY_ESTIMATE };
    const combined = combineEstimate([only, idle]);
    expect(combined.iterationsDone).toBe(3);
    expect(combined.estimate.mean).toBe(1000);
    expect(combined.estimate.min).toBe(1000);
  });

  it('is all zeros before any shard reports, so the page can render it', () => {
    expect(combineEstimate([])).toEqual({ estimate: EMPTY_ESTIMATE, iterationsDone: 0 });
  });
});

describe('confidenceBand', () => {
  it('is the 95% band, which is what the page shows beside the figure', () => {
    expect(confidenceBand({ mean: 1400, stddev: 140, error: 2.5, min: 0, max: 0 })).toBeCloseTo(4.9, 9);
  });
});

describe('formatMargin', () => {
  // The controller's ruling: below 10, one decimal; 10 or more, an integer; a non-zero
  // margin that would round to 0.0 at one decimal reads "< 0.1" rather than "0" -- the
  // tank-sim/healer-sim defect this fixes. A genuinely zero margin still reads "0".
  it.each([
    [0, '0'],
    [0.04, '< 0.1'],
    [0.4, '0.4'],
    [4.24, '4.2'],
    [9.95, '9.9'],
    [10, '10'],
    [147.6, '148'],
  ])('formats %p as %p', (value, expected) => {
    expect(formatMargin(value)).toBe(expected);
  });
});
