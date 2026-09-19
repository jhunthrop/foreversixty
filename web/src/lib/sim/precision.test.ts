import { describe, expect, it } from 'vitest';
import {
  BULK_FINAL_ITERATIONS,
  BULK_PRECISIONS,
  CAPS,
  LANE_ITERATION_CEILING,
  PRECISIONS,
  PRECISION_ITERATIONS,
  STEP_ITERATIONS,
  TARGET_ERROR,
  precisionOf,
  precisionPlan,
  relativeError,
} from './precision';
import type { SimRequest } from './types';

describe('the precision vocabulary', () => {
  it('is the design’s three counts plus the target-error run', () => {
    expect(PRECISIONS).toEqual(['fast', 'normal', 'high', 'target-error']);
    expect(PRECISION_ITERATIONS).toEqual({ fast: 500, normal: 3000, high: 10_000 });
  });

  it('offers only the three fixed counts to a bulk request, which the contract’s BulkSpec takes', () => {
    expect(BULK_PRECISIONS).toEqual(['fast', 'normal', 'high']);
  });

  it('carries contract A2’s combination caps, the server one at five thousand', () => {
    expect(CAPS).toEqual({ browser: 400, server: 5000 });
  });

  it('carries the contract’s half a per cent, thousand-iteration step and lane ceilings', () => {
    expect(TARGET_ERROR).toBe(0.005);
    expect(STEP_ITERATIONS).toBe(1000);
    expect(LANE_ITERATION_CEILING).toEqual({ browser: 30_000, server: 100_000 });
  });

  it('keeps every ceiling a positive multiple of the step, which is contract A3’s rule', () => {
    for (const ceiling of Object.values(LANE_ITERATION_CEILING)) {
      expect(ceiling).toBeGreaterThan(0);
      expect(ceiling % STEP_ITERATIONS).toBe(0);
    }
  });

  // Task 6 ruling: contract A3's bulk final-stage count is 3,000 for fast (same as
  // normal) and 10,000 for high -- a different map from the plain-run PRECISION_ITERATIONS
  // above, which part B reads for its own bulk pages.
  it('carries contract A3’s bulk final-stage counts, a different map from the plain-run one', () => {
    expect(BULK_FINAL_ITERATIONS).toEqual({ fast: 3000, normal: 3000, high: 10_000 });
  });
});

describe('precisionPlan', () => {
  it('turns a fixed precision into a count with no target and no step', () => {
    expect(precisionPlan('normal', 'browser')).toEqual({ iterations: 3000, targetError: 0, step: 0 });
    expect(precisionPlan('high', 'server')).toEqual({ iterations: 10_000, targetError: 0, step: 0 });
  });

  it('turns the target-error run into the lane’s ceiling, the target and the step', () => {
    expect(precisionPlan('target-error', 'browser')).toEqual({
      iterations: 30_000,
      targetError: 0.005,
      step: 1000,
    });
    expect(precisionPlan('target-error', 'server')).toEqual({
      iterations: 100_000,
      targetError: 0.005,
      step: 1000,
    });
  });
});

describe('precisionOf', () => {
  const request = (over: Partial<SimRequest>): SimRequest => ({ ...over }) as SimRequest;

  it('reads a stored request back as the precision that produced it', () => {
    expect(precisionOf(request({ iterations: 500 }))).toBe('fast');
    expect(precisionOf(request({ iterations: 3000 }))).toBe('normal');
    expect(precisionOf(request({ iterations: 10_000 }))).toBe('high');
    expect(precisionOf(request({ iterations: 30_000, target_error: 0.005 }))).toBe('target-error');
  });

  it('falls back to normal for a count no precision names', () => {
    expect(precisionOf(request({ iterations: 1234 }))).toBe('normal');
  });
});

describe('relativeError', () => {
  it('is the error over the mean, and zero rather than infinity at a mean of zero', () => {
    expect(relativeError({ mean: 1000, stddev: 0, error: 5, min: 0, max: 0 })).toBeCloseTo(0.005);
    expect(relativeError({ mean: 0, stddev: 0, error: 5, min: 0, max: 0 })).toBe(0);
  });
});
