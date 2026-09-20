// web/src/lib/sim/bulk-store-request.ts
// The bulk store's request-building surface, extracted from bulk-store.svelte.ts (Task 15,
// following the seam Task 10's fix round 1 already flagged: "the natural next split is the
// request-building cluster... into a bulk-store-request.ts... following store-request.ts's
// own dependency-injection pattern"). Task 15 mounts part A's RequestDrawer here, and that
// component's real props (`disabled`, `onvalidate`, `onapply`, `onrun`, `onshare`, not just
// `onapply`) need four store methods with the same shape as store-request.ts's own
// `buildRequest`/`validateRequest`/`applyRequest`/`runRequest` -- adding them in
// bulk-store.svelte.ts directly would have pushed it well past the 800-line cap.
//
// `$state` cannot cross a module boundary, so this module holds no state of its own: every
// read is a getter (or a plain field, for the two values that never change after
// construction) and every write is a setter, all passed in by `createBulkStore` once, at
// construction -- the same seam store-request.ts uses for `store.svelte.ts`.
import { indexTalents } from '../planner/rules';
import type { TalentFile } from '../planner/types';
import {
  BulkCapError,
  BulkRunError,
  BulkValidationError,
  countCombinations,
  runBulk,
  runWeightsRun,
  type BulkProgress,
  type BulkRunHandle,
} from './bulk-run';
import type { BulkPhase, SimTool } from './bulk-store.svelte';
import {
  SERVER_CAP,
  finalIterations,
  isBulkPrecision,
  type BulkMode,
  type BulkRequest,
  type GearSet,
  type Precision,
  type TalentLoadout,
  type WeightsRequest,
} from './bulk-types';
import { buildBulkSpec, validateBulk, type CandidateRow } from './candidates';
import { toCharacterSpec, type SimCharacter } from './character';
import { bulkCopy, simCopy, weightsUnsupportedSpec } from './copy';
import type { RequestValidation } from './engine';
import { PRECISION_ITERATIONS } from './precision';
import { specLabel } from './spec-label';
import type { SimSettings } from './settings';
import type { CharacterSpec, SimResult, SpecFidelity } from './types';
import { ENGINE_VERSION } from './version';
import { defaultStatsFor, isDpsSpec, referenceFor, weightStatsFor } from './weights';
import type { SimPool } from './worker';

/** Everything `envelope`/`currentSpec`/the four request methods read or write. */
export interface BulkRequestDeps {
  tool: SimTool;
  /** null only for the `weights` tool, which sends no `bulk` block at all. */
  mode: BulkMode | null;
  getCharacter(): SimCharacter | null;
  getTalentFile(): TalentFile | null;
  getSettings(): SimSettings;
  getRows(): CandidateRow[];
  getLocked(): string[];
  getLoadouts(): TalentLoadout[];
  getNamedSets(): GearSet[];
  getPrecision(): Precision;
  getCap(): number;
  getConsumableIds(): string[];
  getStats(): string[];
  /** The one stat `WeightsSpec.Reference` must name -- the store's own, never re-derived. */
  getReferenceStat(): string;
  setPrecision(value: Precision): void;
  setCap(value: number): void;
  setLocked(value: string[]): void;
  setLoadouts(value: TalentLoadout[]): void;
  setNamedSets(value: GearSet[]): void;
  setStats(value: string[]): void;
  scheduleCount(): void;
  poolOnce(): SimPool;
}

/** The character as the engine wants it, or null while a character or its talents are missing. */
function characterSpecOrNull(deps: BulkRequestDeps): CharacterSpec | null {
  const character = deps.getCharacter();
  const talentFile = deps.getTalentFile();
  if (character === null || talentFile === null) return null;
  const settings = deps.getSettings();
  return toCharacterSpec(character, indexTalents(talentFile), settings.buffs, settings.consumables);
}

/**
 * The fields a bulk and a weights request share, or null while a character or its talent
 * file is missing. One function so a run's envelope and a count's can never drift apart --
 * only `iterations` and the trailing block differ.
 */
export function envelope(deps: BulkRequestDeps, iterations: number): Omit<BulkRequest, 'bulk'> | null {
  const character = deps.getCharacter();
  const spec = characterSpecOrNull(deps);
  if (spec === null || character === null) return null;
  return {
    engine_version: ENGINE_VERSION,
    spec: character.spec,
    source: character.source,
    character: spec,
    encounter: deps.getSettings().encounter,
    iterations,
    random_seed: 0,
  };
}

/** The candidate list and its mode, from the store's own ticked rows and lists. */
export function currentSpec(deps: BulkRequestDeps): BulkRequest['bulk'] | null {
  if (deps.mode === null) return null;
  return buildBulkSpec({
    mode: deps.mode,
    rows: deps.getRows(),
    locked: deps.getLocked(),
    loadouts: deps.getLoadouts(),
    sets: deps.getNamedSets(),
    precision: deps.getPrecision(),
    cap: deps.getCap(),
    // "Try each of these" is one alternative list per ticked consumable, not every subset
    // of them (contract 10.1 A5).
    consumables: deps.getConsumableIds().map((id) => [id]),
  });
}

/**
 * Everything `recount` needs beyond `BulkRequestDeps`: the phase gate (read, for the
 * finally-block's "was I the one who started counting" check, and written) and the four
 * fields a count can change -- `combinations`, both cap notices, and `message`/`detail` for
 * a malformed request.
 */
export interface RecountDeps extends BulkRequestDeps {
  getPhase(): BulkPhase;
  setPhase(value: BulkPhase): void;
  setMessage(value: string | null): void;
  setDetail(value: string): void;
  setCombinations(value: number | null): void;
  setCapNotice(value: { cap: number; combinations: number } | null): void;
  setServerCapNotice(value: { cap: number; combinations: number } | null): void;
}

/**
 * The live combination count, debounced by the caller (`bulk-store.svelte.ts`'s own
 * `scheduleCount`, still there -- this function only counts, once). Moved out of
 * `bulk-store.svelte.ts` in fix round 2 (that file was at the 800-line cap, its third split
 * point exhausted): a behaviour-preserving move, not a rewrite -- every branch below is
 * byte-for-byte the same decision its prior home made, with closure variables replaced by
 * `deps` getters/setters, the identical seam `buildRequest`/`envelope` already use.
 */
export async function recount(deps: RecountDeps): Promise<void> {
  // One place owns clearing, at the top, before any exit -- `runBulkAndSettle`'s own rule.
  // Otherwise `BulkValidationError`'s message/detail outlive the next recount once the
  // player fixes what was wrong (fix round 3).
  deps.setMessage(null);
  deps.setDetail('');
  if (deps.mode === null || deps.getCharacter() === null) return;
  const bulk = currentSpec(deps);
  if (bulk === null || validateBulk(bulk) !== null) {
    deps.setCombinations(null);
    deps.setCapNotice(null);
    // The invalid-spec exit used to leave a stale server-cap notice on screen after the
    // player unticked everything past 5,000 (fix round 1, Important 3) -- every exit now
    // clears all three of the same fields.
    deps.setServerCapNotice(null);
    return;
  }
  const base = envelope(deps, finalIterations(deps.getPrecision()));
  if (base === null) return;
  const request: BulkRequest = { ...base, bulk };
  deps.setPhase('counting');
  try {
    const combinations = await countCombinations(deps.poolOnce(), request);
    deps.setCombinations(combinations);
    deps.setCapNotice(null);
    // Derived directly from the count on the success path too, not only through a
    // `BulkCapError` (fix round 1, Minor): `setCap()` is public, so a browser cap raised
    // past 5,000 must not let a 6,000-combination count succeed without this notice.
    deps.setServerCapNotice(combinations > SERVER_CAP ? { cap: SERVER_CAP, combinations } : null);
  } catch (error) {
    if (error instanceof BulkCapError) {
      deps.setCombinations(error.combinations);
      deps.setCapNotice({ cap: error.cap, combinations: error.combinations });
      // Past the premium lane's own 5,000 too (contract 10.1 A2), so the page does not
      // offer a server run the API would refuse at submit with `cap_exceeded`.
      deps.setServerCapNotice(
        error.combinations > SERVER_CAP ? { cap: SERVER_CAP, combinations: error.combinations } : null,
      );
    } else if (error instanceof BulkValidationError) {
      // Engine-lane rule 2: `countCombinations` now validates before it counts, so a
      // malformed request lands here instead of a meaningless cap or count answer. Unlike
      // the generic branch below, this is always actionable by the player (something on
      // the request itself is wrong), so it surfaces through `message`/`detail` the same
      // way a run failure does -- `sim-message` is gated on `message !== null`, so leaving
      // it unset (as the generic branch does) would make `detail` invisible.
      deps.setCombinations(null);
      deps.setCapNotice(null);
      deps.setServerCapNotice(null);
      deps.setMessage(error.message);
      deps.setDetail(error.detail);
    } else {
      // Not a cap or validation refusal: a genuine engine error while merely counting -- an
      // unknown candidate item id (`bulk: the build has no such item: <id>`) is the case
      // this was written for. The count blanks rather than showing a stale number, and
      // unlike the old behaviour, `message` is set too: leaving it unset made `detail`
      // invisible (`sim-message` only renders when `message !== null`, the same gate the
      // `BulkValidationError` branch above already respects), which is exactly how this
      // used to read "Counting combinations…" forever with no explanation on screen.
      deps.setCombinations(null);
      deps.setCapNotice(null);
      deps.setServerCapNotice(null);
      deps.setMessage(bulkCopy.countFailed);
      deps.setDetail(error instanceof Error ? error.message : '');
    }
  } finally {
    if (deps.getPhase() === 'counting') deps.setPhase('idle');
  }
}

export type RequestOutcome = { request: BulkRequest | WeightsRequest } | { error: string };

/** `store.svelte.ts`'s own `buildRequest`, for a bulk or weights request. Refuses with a
 *  reason rather than throwing: `run()` shows the reason as `message`, never a stack trace. */
export function buildRequest(deps: BulkRequestDeps): RequestOutcome {
  const character = deps.getCharacter();
  if (character === null) return { error: bulkCopy.needCharacter };
  if (deps.tool === 'weights') {
    // The healer-sim defect (BLOCKER 2): a spec the engine has no dps model for used to
    // reach the pool and fail silently for 64 seconds. Refused here, before any engine
    // call, with a sentence a player can act on -- never the engine's own words, which
    // name internal spec ids (final whole-branch review: a raw `combine: part 0 failed:
    // request: …` string shown verbatim).
    if (!isDpsSpec(character.spec)) {
      return { error: weightsUnsupportedSpec(specLabel(character.spec)) };
    }
    // Contract 10.8: `WeightsSpec.Reference` is required. An empty `stats` list has
    // nothing to send as one.
    const stats = deps.getStats();
    if (stats.length === 0) return { error: bulkCopy.weightsNeedStats };
    // D45: a weights run used to ignore `precision` entirely and always send `normal`'s
    // 3,000 -- the fixed count `PRECISION_ITERATIONS` already carries for a flat (non-
    // staged) run, the same map `/sim`'s own plain run reads for `fast`/`normal`/`high`.
    const base = envelope(deps, PRECISION_ITERATIONS[deps.getPrecision()]);
    if (base === null) return { error: simCopy.failed };
    // The store's own `referenceStat`, not a second read of `stats[0]`: the page renders
    // "Reference: …" from the former, and two derivations of the one value is how the line
    // and the request come to disagree (final whole-branch review, Minor 8).
    return { request: { ...base, weights: { stats: [...stats], reference: deps.getReferenceStat() } } };
  }
  // Contract 10.1 A3: a bulk request's iterations ARE its precision's final stage.
  const base = envelope(deps, finalIterations(deps.getPrecision()));
  if (base === null) return { error: simCopy.failed };
  const bulk = currentSpec(deps);
  if (bulk === null) return { error: simCopy.failed };
  const refusal = validateBulk(bulk);
  if (refusal !== null) return { error: refusal };
  return { request: { ...base, bulk } };
}

/** Like `buildRequest` but silent, for part A's Advanced drawer (design 8): the drawer
 *  renders what would be sent, it never runs on its own, so an invalid or incomplete state
 *  is simply nothing to preview rather than a message to show. */
export function previewRequest(deps: BulkRequestDeps): BulkRequest | WeightsRequest | null {
  if (deps.getCharacter() === null) return null;
  if (deps.tool === 'weights') {
    const base = envelope(deps, PRECISION_ITERATIONS[deps.getPrecision()]);
    const stats = deps.getStats();
    return base === null
      ? null
      : { ...base, weights: { stats: [...stats], reference: deps.getReferenceStat() } };
  }
  const base = envelope(deps, finalIterations(deps.getPrecision()));
  if (base === null) return null;
  const bulk = currentSpec(deps);
  return bulk === null ? null : { ...base, bulk };
}

/**
 * A request edited in the drawer, adopted whole. Only the two blocks this store owns are
 * read back -- the encounter and the buffs belong to `settings`, which part A's own panel
 * owns, and writing them from here would fight it. The candidate list is not read back
 * either: a `Candidate[]` is derived from ticked rows, not a shape this store re-derives
 * rows from.
 */
export function applyRequestFields(deps: BulkRequestDeps, next: unknown): void {
  const parsed = next as Partial<BulkRequest & WeightsRequest>;
  if (parsed.bulk !== undefined) {
    // Narrowed, never cast: a drawer-edited `precision` is whatever the player typed, and
    // a cast would have put "quick" into the store and through to the wire. An unreadable
    // value leaves the store's own precision alone (final whole-branch review, Minor 8).
    if (isBulkPrecision(parsed.bulk.precision)) deps.setPrecision(parsed.bulk.precision);
    deps.setCap(parsed.bulk.cap);
    deps.setLocked([...(parsed.bulk.locked ?? [])]);
    deps.setLoadouts([...(parsed.bulk.talents ?? [])]);
    deps.setNamedSets([...(parsed.bulk.sets ?? [])]);
  }
  if (parsed.weights !== undefined) deps.setStats([...parsed.weights.stats]);
  deps.scheduleCount();
}

/** `api.SimRequest.Validate`, inside the wasm -- the drawer's own inline JSON check
 *  (design 8), the identical call `store.svelte.ts`'s own `validateRequest` makes. */
export function validateRequestJson(deps: BulkRequestDeps, json: string): Promise<RequestValidation> {
  return deps.poolOnce().validate(json);
}

/**
 * The weights picker's seed, run whenever a character or the spec list newly become
 * available. Returns `currentStats` unchanged once it is non-empty, so a player's own
 * edits are never overwritten -- a pure function rather than a store method, since seeding
 * is "what should `stats` be", not an action with side effects of its own.
 */
export function seededStats(
  tool: SimTool,
  character: SimCharacter | null,
  specRows: SpecFidelity[],
  currentStats: string[],
): string[] {
  if (tool !== 'weights' || currentStats.length > 0 || character === null) return currentStats;
  return defaultStatsFor(
    character.spec,
    referenceFor(character.spec, specRows),
    weightStatsFor(character.spec, specRows),
  );
}

/** Everything `runBulkAndSettle` needs from the store's own `$state`. */
export interface BulkRunDeps {
  tool: SimTool;
  poolOnce(): SimPool;
  getStopRequested(): boolean;
  setStopRequested(value: boolean): void;
  setPhase(value: BulkPhase): void;
  setProgress(value: BulkProgress | null): void;
  setMessage(value: string | null): void;
  setDetail(value: string): void;
  getResult(): SimResult | null;
  setResult(value: SimResult | null): void;
  setHandle(value: BulkRunHandle | null): void;
  setCapNotice(value: { cap: number; combinations: number } | null): void;
  setCombinations(value: number | null): void;
  bumpServerGeneration(): void;
}

/**
 * One `BulkRequest`/`WeightsRequest`, run to a finish or a stop -- the cancel/restore/phase
 * decision `run()` (the store's own ticked-candidates request) and `runRequest()` (the
 * drawer's edited-and-validated one, design 8's escape hatch) share, so there is exactly
 * one rule for what a finished, a cancelled and a genuinely failed run each mean.
 */
export async function runBulkAndSettle(
  deps: BulkRunDeps,
  request: BulkRequest | WeightsRequest,
): Promise<void> {
  deps.bumpServerGeneration();
  deps.setMessage(null);
  deps.setDetail('');
  deps.setStopRequested(false);
  deps.setPhase('running');
  deps.setProgress(null);

  const handle =
    deps.tool === 'weights'
      ? runWeightsRun(deps.poolOnce(), request as WeightsRequest, () => {})
      : runBulk(deps.poolOnce(), request as BulkRequest, (next) => {
          if (deps.getStopRequested()) return;
          deps.setProgress(next);
        });
  deps.setHandle(handle);

  try {
    const finished = await handle.result;
    deps.setResult(finished);
    if (finished.aborted === true) deps.setMessage(bulkCopy.partial);
    deps.setPhase('done');
  } catch (error) {
    if (error instanceof BulkCapError) {
      deps.setCapNotice({ cap: error.cap, combinations: error.combinations });
      deps.setCombinations(error.combinations);
      deps.setMessage(error.message);
      deps.setPhase('idle');
      return;
    }
    const failure = error instanceof BulkRunError ? error : null;
    deps.setMessage(
      failure?.cancelled === true ? simCopy.stopped : (failure?.message ?? bulkCopy.bulkFailed),
    );
    deps.setDetail(failure?.detail ?? '');
    deps.setPhase(failure?.cancelled === true && deps.getResult() !== null ? 'done' : 'error');
  } finally {
    deps.setHandle(null);
    deps.setProgress(null);
  }
}
