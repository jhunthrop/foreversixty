import { describe, expect, it } from 'vitest';
import fixture from '../../fixtures/sim/result.json';
import { percentLabel, runDetails } from './details';
import type { SimResult } from './types';

const base = fixture as unknown as SimResult;

function result(over: Partial<SimResult>): SimResult {
  return { ...base, ...over };
}

describe('runDetails', () => {
  it('reports the 95% band in DPS and the relative standard error as a percent', () => {
    const details = runDetails(result({ dps: { mean: 1000, stddev: 100, error: 4, min: 0, max: 0 } }));
    // 1.96 * 4 = 7.84, rounded.
    expect(details.bandDps).toBe(8);
    expect(details.errorPercent).toBeCloseTo(0.004);
  });

  it('carries the iterations, the wall clock, the engine and the lane straight through', () => {
    const details = runDetails(result({ iterations_run: 7000, duration_ms: 12_345, lane: 'server' }));
    expect(details.iterations).toBe(7000);
    expect(details.processingMs).toBe(12_345);
    expect(details.lane).toBe('server');
    expect(details.engineVersion).toBe(base.engine_version);
  });

  it('says a target-error run stopped at the ceiling, and says nothing of the sort otherwise', () => {
    const capped = result({
      iterations_run: 30_000,
      request: { ...base.request, iterations: 30_000, target_error: 0.005 },
    });
    expect(runDetails(capped).hitCeiling).toBe(true);

    const converged = result({
      iterations_run: 4000,
      request: { ...base.request, iterations: 30_000, target_error: 0.005 },
    });
    expect(runDetails(converged).hitCeiling).toBe(false);

    const fixed = result({ iterations_run: 3000, request: { ...base.request, iterations: 3000 } });
    expect(runDetails(fixed).hitCeiling).toBe(false);
  });

  it('never divides by a mean of zero', () => {
    expect(runDetails(result({ dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 } })).errorPercent).toBe(
      0,
    );
  });
});

describe('percentLabel', () => {
  it('is two decimals and a sign-free per cent', () => {
    expect(percentLabel(0.004)).toBe('0.40%');
    expect(percentLabel(0.005)).toBe('0.50%');
    expect(percentLabel(0)).toBe('0.00%');
  });
});
