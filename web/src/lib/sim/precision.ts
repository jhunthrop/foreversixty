// web/src/lib/sim/precision.ts
// Design 4.2 and contract 1.2. Three fixed counts and one stopping rule.
//
// "Until ±0.5%" is Raidbots' Smart Sim, made visible: the run continues in thousand-
// iteration steps until the DPS error is inside half a per cent, or until the lane's
// ceiling is reached, and the results line says which of the two stopped it. The decision
// itself is never taken here -- `SimPool.needsMore` asks the engine after every step (see
// run.ts). This module only carries the numbers.
//
// Contract A3 requires a target-error run's `Iterations` to be a positive multiple of
// `StepIterations` at or under the lane's ceiling; both ceilings below satisfy that, and
// precision.test.ts asserts it so a future ceiling cannot quietly break `Validate`.
import type { Estimate, SimRequest } from './types';
import { STEP_ITERATIONS_DEFAULT } from './types';

export type Lane = 'browser' | 'server';

export type PrecisionId = 'fast' | 'normal' | 'high' | 'target-error';

export const PRECISIONS: readonly PrecisionId[] = ['fast', 'normal', 'high', 'target-error'];

/** `BulkSpec.Precision` takes only the three fixed counts (contract 1.3). */
export const BULK_PRECISIONS: readonly Exclude<PrecisionId, 'target-error'>[] = ['fast', 'normal', 'high'];

/**
 * A plain run's own iteration count for each fixed precision. **Plain-run only** -- a bulk
 * request's final-stage count is a different figure (contract A3: a bulk `fast` run's final
 * stage runs 3,000, not 500). `BULK_FINAL_ITERATIONS` below is that map; part B's pages read
 * that one, not this one, for `BulkSpec.iterations`.
 */
export const PRECISION_ITERATIONS: Record<Exclude<PrecisionId, 'target-error'>, number> = {
  fast: 500,
  normal: 3000,
  high: 10_000,
};

/**
 * A bulk request's final-stage iteration count, per contract A3: "a bulk request's
 * `Iterations` is its precision's final-stage count (3,000; 10,000 for `high`)" -- so a
 * bulk `fast` run's final stage runs at the same count as `normal`. Nothing on `/sim`
 * builds a bulk request; this lives here, beside `PRECISION_ITERATIONS`, so part B's bulk
 * pages read the one number the contract actually specifies rather than reimplementing it.
 */
export const BULK_FINAL_ITERATIONS: Record<Exclude<PrecisionId, 'target-error'>, number> = {
  fast: 3000,
  normal: 3000,
  high: 10_000,
};

/** Half a per cent, which is what the control is named after. */
export const TARGET_ERROR = 0.005;

export const STEP_ITERATIONS = STEP_ITERATIONS_DEFAULT;

export const LANE_ITERATION_CEILING: Record<Lane, number> = { browser: 30_000, server: 100_000 };

/**
 * Contract A2: how many combinations a bulk request may expand to on each lane. The
 * server figure is 5,000, not the design's 20,000 -- at the measured native rate a
 * 20,000-combination fast run cannot finish inside the job's fifteen minutes.
 *
 * Nothing on `/sim` uses this; it lives beside the iteration ceilings because they are
 * the same kind of fact, and part B's cap notice would otherwise keep a second copy.
 */
export const CAPS: Record<Lane, number> = { browser: 400, server: 5000 };

export interface PrecisionPlan {
  /** The count for a fixed run, the ceiling for a target-error run. */
  iterations: number;
  /** 0 for a fixed run. */
  targetError: number;
  /** 0 for a fixed run. */
  step: number;
}

export function precisionPlan(id: PrecisionId, lane: Lane): PrecisionPlan {
  if (id === 'target-error') {
    return { iterations: LANE_ITERATION_CEILING[lane], targetError: TARGET_ERROR, step: STEP_ITERATIONS };
  }
  return { iterations: PRECISION_ITERATIONS[id], targetError: 0, step: 0 };
}

/**
 * A stored request read back as the control that produced it, for a saved sim and for a
 * request pasted into the drawer. A count no precision names -- a hand-edited request --
 * reads as `normal`: the control has to show something, and the request itself, not the
 * control, is what will be re-run.
 */
export function precisionOf(request: Pick<SimRequest, 'iterations' | 'target_error'>): PrecisionId {
  if ((request.target_error ?? 0) > 0) return 'target-error';
  const found = (Object.keys(PRECISION_ITERATIONS) as Exclude<PrecisionId, 'target-error'>[]).find(
    (id) => PRECISION_ITERATIONS[id] === request.iterations,
  );
  return found ?? 'normal';
}

/** The band as a fraction of the figure. Zero at a mean of zero rather than infinity. */
export function relativeError(estimate: Estimate): number {
  return estimate.mean === 0 ? 0 : estimate.error / estimate.mean;
}
