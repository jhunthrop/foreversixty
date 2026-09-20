// web/src/lib/sim/bulk-types.ts
// TypeScript mirrors of sim/api/envelope.go's bulk and weights additions, contract sections
// 1.3, 1.4, 2 and 4. Every key here is that file's `json:` tag.
//
// Part A (`sim-parity-web-a`) already declared the shapes shared with the plain-run page --
// `Candidate`, `TalentLoadout`, `GearSet`, `BulkSpec`, `WeightsSpec`, `Substitution`,
// `Combo`, `Stage`, `StatWeight`, `SimRequest.bulk`/`.weights`, `SimResult.weights`,
// `SimKind`/`requestKind` -- in `./types` and `./kind`. This module RE-EXPORTS those (so
// every later task in this lane has one import site, as the plan intends) and declares
// only what genuinely has no other home: the request/result envelopes that make `bulk`/
// `weights` required rather than optional, the wasm plan/rank exchange shapes, the cap
// refusal, and the browser cap rule. Nothing here retypes a number `precision.ts` already
// owns (`CAPS`, `BULK_PRECISIONS`, `BULK_FINAL_ITERATIONS`) -- see the two-line comment
// above each derived constant below for where its value actually comes from.
//
// Nothing here computes a statistic. Expansion, staging, ranking and the within-error
// grouping all happen inside sim/bulk in the wasm; TypeScript names the shapes and renders
// them.
import type {
  BulkSpec,
  Combo,
  Estimate,
  SimProgress,
  SimRequest,
  SimResult,
  Stage,
  StatWeight,
  Substitution,
  WeightsSpec,
} from './types';
import { BULK_FINAL_ITERATIONS, BULK_PRECISIONS, CAPS } from './precision';

// --- re-exports: already shipped by part A, one import site for this lane ---
export type {
  BulkSpec,
  Candidate,
  Combo,
  GearSet,
  Stage,
  StatWeight,
  Substitution,
  TalentLoadout,
  WeightsSpec,
} from './types';
export type { SimKind } from './kind';
/**
 * The derived-kind function is part A's `requestKind`; no second name is introduced for it
 * (controller ruling F11) -- re-exported under its own spelling, not an alias, so a later
 * task can never end up choosing between two names for the same function.
 */
export { requestKind } from './kind';

// --- new: the mode vocabulary and the precision ladder ---

export type BulkMode = 'gear' | 'talents' | 'drops';

/**
 * The contract's three, in ladder order: more stages and fewer survivors, left to right.
 * A straight re-export of `precision.ts`'s own `BULK_PRECISIONS` -- one array, not a second
 * one under a different name -- kept under its own name rather than rebranded to
 * `PRECISIONS`: `precision.ts` already exports a *different*, four-value `PRECISIONS`
 * (it also carries `target-error`), and a same-named-but-different export here would force
 * every future file needing both to alias one on import, with no compiler signal until
 * that exact moment.
 */
export { BULK_PRECISIONS } from './precision';
export type Precision = (typeof BULK_PRECISIONS)[number];

/**
 * How many stages a precision runs, for the progress line only. The iteration counts and
 * the survivor cuts are sim/bulk's and are never mirrored here: fast is 100 -> 1000 ->
 * 3000, normal is 1000 -> 3000, high is 1000 -> 10000 (contract 1.3).
 */
export const STAGES_BY_PRECISION: Record<Precision, number> = { fast: 3, normal: 2, high: 2 };

/**
 * `BulkSpec.precision` is a bare `string` on the wire, so anything that reads one back --
 * a request hand-edited in the drawer, a stored request replayed by the premium lane --
 * must narrow it rather than cast it (final whole-branch review, Minor 8). The caller
 * decides the fallback; this only answers whether the value is one of the three.
 */
export function isBulkPrecision(value: string): value is Precision {
  return (BULK_PRECISIONS as readonly string[]).includes(value);
}

/**
 * The iteration count a bulk request carries in `SimRequest.iterations` -- its precision's
 * FINAL stage, not `precision.ts`'s plain-run count (contract 10.1 A3). Wraps
 * `BULK_FINAL_ITERATIONS`, the one place 3,000/10,000 are typed, rather than repeating
 * them here.
 */
export function finalIterations(precision: Precision): number {
  return BULK_FINAL_ITERATIONS[precision];
}

// --- new: the browser cap rule (contract 1.3's Caps, plus design 2.3's low-core halving) ---

/** `precision.ts`'s `CAPS.browser` -- 400 -- under this module's own name. */
export const BROWSER_CAP = CAPS.browser;
/**
 * `precision.ts`'s `CAPS.server` -- 5,000, not contract 1.3's 20,000: contract 10.1 A2
 * lowered it because a 20,000-combination fast run cannot finish inside the job's 15-minute
 * timeout.
 */
export const SERVER_CAP = CAPS.server;
/** Design 2.3: "a phone gets half the cap by hardwareConcurrency". */
export const LOW_CORE_CAP = BROWSER_CAP / 2;
export const LOW_CORE_THRESHOLD = 4;

/**
 * The cap this device gets. An absent or nonsensical `hardwareConcurrency` takes the low
 * cap rather than the high one: a browser that will not say how many cores it has is far
 * more often a phone with few than a workstation with many, and the run button always says
 * what premium would allow anyway.
 */
export function browserCap(hardwareConcurrency: number | undefined): number {
  if (typeof hardwareConcurrency !== 'number' || !Number.isFinite(hardwareConcurrency)) {
    return LOW_CORE_CAP;
  }
  if (hardwareConcurrency <= 0) return LOW_CORE_CAP;
  return hardwareConcurrency <= LOW_CORE_THRESHOLD ? LOW_CORE_CAP : BROWSER_CAP;
}

// --- new: the envelopes that make bulk/weights required, not optional ---

/** A `SimRequest` that is actually a bulk run: `bulk` is required, never optional. */
export interface BulkRequest extends SimRequest {
  bulk: BulkSpec;
}

/** A `SimRequest` that is actually a weights run: `weights` is required, never optional. */
export interface WeightsRequest extends SimRequest {
  weights: WeightsSpec;
}

/** A `SimResult` for a bulk kind: `combos`/`equipped`/`stages` are required, never optional. */
export interface BulkResult extends SimResult {
  combos: Combo[];
  /** The base character at the final stage. */
  equipped: Estimate;
  stages: Stage[];
}

/** A `SimResult` for a weights kind: `weights` is required, never optional. */
export interface WeightsResult extends SimResult {
  weights: StatWeight[];
}

// --- new: the wasm plan/rank exchange (simPlan, simRank) ---

/** One combination `simPlan`/`simRank` asks the pool to run. */
export interface Combination {
  request: SimRequest;
  substitutions: Substitution[];
}

/** What `simPlan` returns, and what `simRank` returns as `next`. */
export interface StageRequests {
  stage: number;
  iterations: number;
  /** requests[0] is always the equipped set. */
  requests: SimRequest[];
  /** Parallel to requests[1:]. */
  combos: Combination[];
  /**
   * The ladder's history so far (contract 10.1 A10): what each finished stage cost. It
   * rides on the stage object across the wasm boundary so `simRank` can fill
   * `SimResult.stages` without the page keeping a tally of its own.
   */
  ran?: Stage[];
}

/** What `simRank` answers: one of the two is set. */
export interface RankAnswer {
  next?: StageRequests;
  result?: SimResult;
}

/**
 * sim/bulk's ErrCapExceeded across the wasm boundary. Plain `{"error": "..."}` carries no
 * numbers and the page has to say what the cap was and by how much the list overran it, so
 * the cap refusal is the one structured error the exports answer with.
 */
export interface CapExceeded {
  error: 'cap_exceeded';
  cap: number;
  combinations: number;
}

export function isCapExceeded(value: unknown): value is CapExceeded {
  if (typeof value !== 'object' || value === null) return false;
  const record = value as Record<string, unknown>;
  return (
    record.error === 'cap_exceeded' &&
    typeof record.cap === 'number' &&
    typeof record.combinations === 'number'
  );
}

// --- new: GET /v1/sims/<id>/progress for a bulk job (contract 2, "Progress (+)") ---

/**
 * `fetchBulkProgress`'s answer. `SimProgress` (`./types`) already carries `stage`,
 * `combos_done` and `combos_total` as optional fields -- Task 1 added them there directly
 * (commit 9d96cac), which is contract A11's "additive progress widening" already done at
 * the source. A plain alias, not a redeclared `extends` block: repeating those three fields
 * on a second interface would be the same optional shape typed twice for no reason, and
 * this lane's own DRY rule is "no code duplication... single source of truth". The alias
 * still gives this lane the contract's own name to import from `./bulk-types`, its one
 * import site, without a second place those three fields could drift out of sync.
 */
export type BulkServerProgress = SimProgress;
