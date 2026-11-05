// web/src/lib/sim/spread.test.ts
// Spec 2026-11-04-sim-results 3.6 and section 10, check 6.
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import { MAX_BARS, percentileOf, spreadOf } from './spread';
import type { DpsHistogram } from './types';

const histogram = (full as unknown as { dps_histogram: DpsHistogram }).dps_histogram;

describe('spreadOf', () => {
  it('groups the 385 bins into 10 DPS bars, at most 40, and keeps every run', () => {
    const spread = spreadOf(histogram);
    expect(spread?.bars.length).toBeLessThanOrEqual(MAX_BARS);
    expect(spread?.bars[0]).toMatchObject({ low: 731, high: 740 });
    expect(spread?.total).toBe(10_000);
    expect(spread?.bars.reduce((sum, bar) => sum + bar.count, 0)).toBe(10_000);
  });

  it('reads the two-thirds band off the real bins', () => {
    expect(percentileOf(histogram, 0.16)).toBe(856);
    expect(percentileOf(histogram, 0.84)).toBe(952);
    expect(spreadOf(histogram)).toMatchObject({ low: 856, high: 952 });
  });

  it('marks only the bars that touch the band as central', () => {
    const bars = spreadOf(histogram)?.bars ?? [];
    expect(bars.find((bar) => bar.low === 731)?.central).toBe(false);
    expect(bars.find((bar) => bar.low <= 904 && bar.high >= 904)?.central).toBe(true);
  });

  it('draws nothing without bins: the older result gets the range strip', () => {
    expect(spreadOf(undefined)).toBeNull();
    expect(spreadOf({ start: 1, bin_dps: 1, counts: [] })).toBeNull();
    expect(spreadOf({ start: 1, bin_dps: 1, counts: [0, 0] })).toBeNull();
  });
});
