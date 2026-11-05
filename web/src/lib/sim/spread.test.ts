// web/src/lib/sim/spread.test.ts
// Spec 2026-11-04-sim-results 3.6 and section 10, check 6.
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import { MAX_BARS, percentileOf, spreadOf } from './spread';
import type { DpsHistogram } from './types';

const histogram = (full as unknown as { dps_histogram: DpsHistogram }).dps_histogram;

describe('spreadOf', () => {
  it('draws the 40 bins of 10 DPS as they are, and keeps every run', () => {
    const spread = spreadOf(histogram);
    expect(spread?.bars).toHaveLength(MAX_BARS);
    expect(spread?.bars[0]).toMatchObject({ low: 730, high: 739 });
    expect(spread?.total).toBe(10_000);
    expect(spread?.bars.reduce((sum, bar) => sum + bar.count, 0)).toBe(10_000);
  });

  it('groups past 40 bins: 385 bins of 1 DPS become 10 DPS bars', () => {
    const spread = spreadOf({ start: 731, bin_dps: 1, counts: Array.from({ length: 385 }, () => 1) });
    expect(spread?.bars.length).toBeLessThanOrEqual(MAX_BARS);
    expect(spread?.bars[0]).toMatchObject({ low: 731, high: 740, count: 10 });
  });

  it('reads the two-thirds band off the real bins, by hand from the cumulative counts', () => {
    const { start, bin_dps: bin, counts } = histogram;
    const at = (share: number): number => {
      let seen = 0;
      for (const [i, count] of counts.entries()) {
        if (seen + count >= 10_000 * share) return start + (i + (10_000 * share - seen) / count) * bin;
        seen += count;
      }
      throw new Error('past the end');
    };
    expect(percentileOf(histogram, 0.16)).toBeCloseTo(at(0.16), 6);
    expect(percentileOf(histogram, 0.84)).toBeCloseTo(at(0.84), 6);
    // Within one spread of the mean, as the run's own 49.5 says.
    expect(Math.abs(904.37 - 49.46 - percentileOf(histogram, 0.16))).toBeLessThan(10);
    expect(Math.abs(904.37 + 49.46 - percentileOf(histogram, 0.84))).toBeLessThan(10);
  });

  it('marks only the bars that touch the band as central', () => {
    const bars = spreadOf(histogram)?.bars ?? [];
    expect(bars.find((bar) => bar.low === 730)?.central).toBe(false);
    expect(bars.find((bar) => bar.low <= 904 && bar.high >= 904)?.central).toBe(true);
  });

  it('draws nothing without bins: the older result gets the range strip', () => {
    expect(spreadOf(undefined)).toBeNull();
    expect(spreadOf({ start: 1, bin_dps: 1, counts: [] })).toBeNull();
    expect(spreadOf({ start: 1, bin_dps: 1, counts: [0, 0] })).toBeNull();
  });
});
