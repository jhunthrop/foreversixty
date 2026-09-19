// web/src/lib/sim/store-request.ts
// The request drawer's four store methods, extracted from store.svelte.ts (ruling, Task
// 15): `store.svelte.ts` was 712 of the lane's 800-line ceiling, and these four would have
// pushed it past. The public surface does not move -- `store.svelte.ts` still exposes
// `buildRequest`, `validateRequest`, `applyRequest` and `runRequest` as methods on the
// store it returns; this module only holds their bodies.
//
// `$state` cannot cross a module boundary: the reactive fields these methods read and
// write are `let` bindings inside `createSimStore`'s own closure, and Svelte's runes are
// ordinary lexical `let`s under the hood, not a value that can be exported. So the seam is
// dependency injection -- `StoreRequestDeps` is a plain object of getters, setters and the
// store's own private helpers, built once in `store.svelte.ts` and closed over there. This
// module never touches `$state` itself and needs no `.svelte.ts` extension.
import { indexTalents } from '../planner/rules';
import type { TalentFile } from '../planner/types';
import { codeForCharacterSpec, toCharacterSpec, type SimCharacter } from './character';
import { simCopy } from './copy';
import type { RequestValidation } from './engine';
import { EMPTY_ESTIMATE } from './estimate';
import { precisionOf, precisionPlan, STEP_ITERATIONS, type PrecisionId } from './precision';
import { settingsFromRequest } from './request-json';
import { buildSimRequest, runSim, SimRunError, type RunHandle, type RunInput, type RunUpdate } from './run';
import type { SimSettings } from './settings';
import type { SourceResult } from './sources';
import type { SimPhase } from './store.svelte';
import type { SimRequest, SimResult } from './types';
import type { SimPool } from './worker';

/**
 * Everything `runAndSettle` needs to run one `RunInput` to a finish or a stop. Both
 * `run()` (store.svelte.ts) and `runRequest()` below build a deps object satisfying this
 * (structurally -- `StoreRequestDeps` extends it) and hand it the input they each built
 * their own way; from here on there is exactly one rule for what "cancelled" means.
 */
export interface RunSettleDeps {
  getStopRequested(): boolean;
  /** `estimate`, `iterationsDone`, `iterationsTotal` and `relative`, written together --
   *  the four fields a run's own progress callback writes, from `RunUpdate`'s own shape
   *  rather than a fifth copy of their names. */
  setProgress(update: RunUpdate): void;
  setPhase(value: SimPhase): void;
  setMessage(value: string | null): void;
  setDetail(value: string): void;
  getResult(): SimResult | null;
  setResult(value: SimResult | null): void;
  setHandle(value: RunHandle | null): void;
  /** Puts the figure back to the last completed result after a cancelled run. */
  restorePreviousResult(): void;
}

/**
 * Everything the four methods below need from `createSimStore`'s own closure. Every
 * setter here mutates exactly the `$state` field `store.svelte.ts` declares under the
 * matching name -- this module writes state, it just does not own any.
 */
export interface StoreRequestDeps extends RunSettleDeps {
  getCharacter(): SimCharacter | null;
  getTalents(): TalentFile | null;
  getSettings(): SimSettings;
  setSettings(value: SimSettings): void;
  getPrecisionId(): PrecisionId;
  setPrecisionId(value: PrecisionId): void;
  setStopRequested(value: boolean): void;
  /** Creates the pool on first use; every other lane already funnels through this. */
  poolOnce(): SimPool;
  /** The store's own single load-and-adopt path -- the failure rule lives there, once. */
  adopt(load: Promise<SourceResult>): Promise<void>;
  /** An FS1 code decoded into a `'manual'`-sourced character, bound to the store's own
   *  `LoadContext`. */
  fromPlannerCode(code: string): Promise<SourceResult>;
  /** `init.treeVersion`, for the FS1 code `applyRequest` encodes. */
  treeVersion: string;
}

/**
 * One `RunInput`, run to a finish or a stop. This is the cancel/restore/phase decision
 * `run()` and `runRequest()` both need and, until this fix round, both duplicated --
 * this lane's review has now caught that six times across the branch, and two copies of
 * one rule drift the first time someone fixes only one of them. Callers differ only in
 * how they build `RunInput` (`run()` from the page's own settings, `runRequest()` from
 * the edited text); from `phase: 'running'` on, there is exactly one rule, here.
 */
export async function runAndSettle(deps: RunSettleDeps, pool: SimPool, input: RunInput): Promise<void> {
  deps.setPhase('running');
  const handle = runSim(pool, input, (update) => {
    // A shard can still report progress after stop() fires and before the engine has
    // noticed the abort message; the figure on screen must not keep moving once the
    // player has asked it to stop.
    if (deps.getStopRequested()) return;
    deps.setProgress(update);
  });
  deps.setHandle(handle);

  try {
    const finished = await handle.result;
    // handle.result can resolve with a real result even after stop(): the abort message
    // and the engine's own last tick can cross in flight, and a shard mid-tick when the
    // message arrives finishes it rather than discarding the work. The player's Stop
    // still wins -- the number on screen is the one from before this run, not a result
    // they asked to discard.
    if (deps.getStopRequested()) {
      deps.setMessage(simCopy.stopped);
      deps.restorePreviousResult();
      deps.setPhase(deps.getResult() !== null ? 'done' : 'idle');
      return;
    }
    deps.setResult(finished);
    deps.setPhase('done');
  } catch (error) {
    const failure = error instanceof SimRunError ? error : null;
    deps.setMessage(failure?.cancelled === true ? simCopy.stopped : (failure?.message ?? simCopy.failed));
    // The engine's own words, kept beside ours: sim/request names the buff or consumable
    // id it could not map, and that is the only thing that says what to change.
    deps.setDetail(failure?.detail ?? '');
    if (failure?.cancelled === true) deps.restorePreviousResult();
    deps.setPhase(failure?.cancelled === true && deps.getResult() !== null ? 'done' : 'error');
  } finally {
    deps.setHandle(null);
  }
}

export interface RequestMethods {
  /**
   * The request the page would send right now, or null while there is no character or no
   * talent index. One function, so the drawer, and eventually the share link, can never
   * disagree about what "this request" is -- the same reason `buildSimRequest` exists.
   */
  buildRequest(): SimRequest | null;
  /** `api.SimRequest.Validate`, inside the wasm. Never a rule written here. */
  validateRequest(json: string): Promise<RequestValidation>;
  /**
   * A pasted request as page state: settings and precision exactly, and the character
   * through the same FS1 route "Run this yourself" already uses. `codeForCharacterSpec`
   * carries the request's own gear list -- enchants and suffixes included (contract 10.5)
   * -- so this is lossless for everything `CharacterSpec` models, the same as `runRequest`
   * below.
   */
  applyRequest(request: SimRequest): Promise<void>;
  /** The edited request, run exactly as written. The escape hatch of design 8. */
  runRequest(request: SimRequest): Promise<void>;
}

export function createRequestMethods(deps: StoreRequestDeps): RequestMethods {
  return {
    buildRequest(): SimRequest | null {
      const character = deps.getCharacter();
      if (character === null) return null;
      const talentFile = deps.getTalents();
      const index = talentFile === null ? null : indexTalents(talentFile);
      if (index === null) return null;
      const settings = deps.getSettings();
      const plan = precisionPlan(deps.getPrecisionId(), 'browser');
      return buildSimRequest({
        spec: character.spec,
        source: character.source,
        character: toCharacterSpec(
          character,
          index,
          settings.buffs,
          settings.consumables,
          settings.cooldowns,
        ),
        encounter: settings.encounter,
        iterations: plan.iterations,
        targetError: plan.targetError,
        stepIterations: plan.step,
      });
    },

    validateRequest(json: string): Promise<RequestValidation> {
      return deps.poolOnce().validate(json);
    },

    async applyRequest(request: SimRequest): Promise<void> {
      deps.setSettings(settingsFromRequest(request));
      deps.setPrecisionId(precisionOf(request));
      // codeForCharacterSpec: contract 10.5 lets a gear entry carry
      // `item_id[:enchant[:suffix]]`, and `characterFromFs1` reads it straight back onto
      // `gear_slots`. Apply is therefore lossless for everything `CharacterSpec` models,
      // which is what makes a shared request a full reproduction rather than an
      // approximation of one.
      await deps.adopt(deps.fromPlannerCode(codeForCharacterSpec(request.character, deps.treeVersion)));
    },

    async runRequest(request: SimRequest): Promise<void> {
      deps.setMessage(null);
      deps.setDetail('');
      deps.setStopRequested(false);
      deps.setPhase('loading-engine');
      deps.setProgress({
        estimate: EMPTY_ESTIMATE,
        iterationsDone: 0,
        iterationsTotal: request.iterations,
        relativeError: 0,
      });
      await runAndSettle(deps, deps.poolOnce(), {
        spec: request.spec,
        source: request.source,
        character: request.character,
        encounter: request.encounter,
        iterations: request.iterations,
        randomSeed: request.random_seed,
        targetError: request.target_error,
        stepIterations: STEP_ITERATIONS,
      });
    },
  };
}
