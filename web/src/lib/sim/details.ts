// web/src/lib/sim/details.ts
// Design 5.1's details card: "margin of error, iterations, processing time, engine version
// and lane as a details card beside the results, the way Raidbots' sidebar reads".
//
// Two figures, never the same one twice: `bandDps` is the 95% confidence band the big
// figure already shows (1.96 x the standard error), and `errorPercent` is the relative
// standard error, which is the quantity `simNeedsMore` compares against `target_error`.
// The "until +/-0.5%" control is named after the second, so the second is what any percent
// on this lane means.
import { confidenceBand, formatMargin } from './estimate';
import { relativeError } from './precision';
import type { SimResult } from './types';

export interface RunDetails {
  /** The 95% confidence band, in DPS, already formatted through estimate.ts's formatMargin. */
  bandDps: string;
  /** `error / mean`, as a fraction. */
  errorPercent: number;
  iterations: number;
  processingMs: number;
  engineVersion: string;
  lane: 'browser' | 'server';
  /**
   * True when a target-error run reached the lane's ceiling with the band still open --
   * the design's "the results line says which" of the two things stopped it.
   */
  hitCeiling: boolean;
}

export function runDetails(result: SimResult): RunDetails {
  const target = result.request.target_error ?? 0;
  return {
    bandDps: formatMargin(confidenceBand(result.dps)),
    errorPercent: relativeError(result.dps),
    iterations: result.iterations_run,
    processingMs: result.duration_ms,
    engineVersion: result.engine_version,
    lane: result.lane,
    hitCeiling: target > 0 && result.iterations_run >= result.request.iterations,
  };
}

/** A fraction as a per cent with two decimals: 0.004 is "0.40%". */
export function percentLabel(fraction: number): string {
  return `${(fraction * 100).toFixed(2)}%`;
}
