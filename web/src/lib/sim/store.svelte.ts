// web/src/lib/sim/store.svelte.ts
// The simulator island's single source of truth, in the same shape as the planner's
// store.svelte.ts: runes compile here, no DOM dependency, so it is unit-tested like any
// other module and every component below it is a pure render of what it exposes.
//
// Two decisions this file makes and the components must not second-guess:
//
//   * A failed load keeps the character already on screen. Pasting a bad build id after
//     loading a good addon export must not empty the page; the message says what went wrong
//     and the strip stays.
//   * The pool is created on the first run, never on mount. The wasm is 4 MB and the sim
//     page has the site's 1.6 s mobile LCP budget, so nothing touches the engine until the
//     player asks for a number. `init.pool` exists for tests and for the planner, which
//     brings its own.
import { loadItems, loadReference, loadTalents } from '../planner/load';
import type { ClassRow, Item, RaceRow, TalentFile } from '../planner/types';
import { indexTalents } from '../planner/rules';
import { dispatchServerSim, fetchSim, fetchSimProgress, saveSim, SimApiError } from './api';
import { characterFromFs1, needsRace, toCharacterSpec, type SimCharacter } from './character';
import { loadActionNames, type ActionNames } from './action-names';
import { simCopy } from './copy';
import { EMPTY_ESTIMATE } from './estimate';
import { precisionPlan, relativeError, type Lane, type PrecisionId } from './precision';
import { buildSimRequest, runSim, SimRunError, type RunHandle, type RunInput } from './run';
import { defaultSettings, type SimSettings } from './settings';
import {
  fromAddonExport,
  fromLoggedFight,
  fromPlannerBuild,
  fromStoredCharacter,
  type LoadContext,
  type SourceResult,
} from './sources';
import type { CharacterPath } from '../characters';
import type { Estimate, SimProgress, SimResult, SourceKind } from './types';
import { createPool, type SimPool } from './worker';

export type SimPhase = 'idle' | 'loading-character' | 'loading-engine' | 'running' | 'done' | 'error';

/**
 * The server lane's poll interval (`runOnServer` below). `SimStoreInit.serverPollMs` is
 * the seam a test overrides instead of mocking timers -- the same reason `run.ts` takes a
 * `now` function rather than calling `Date.now()` itself.
 */
const DEFAULT_SERVER_POLL_MS = 2000;

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/**
 * An unsaved planner build's own FS1 code, into a `'manual'`-sourced character. This is
 * `sources.ts`'s `fromAddonExport` in every step but the source it stamps: an addon export
 * and a "Sim this build" link decode through the identical `FS1:…` grammar and the
 * identical `characterFromFs1`, and only differ in where the character came from, which the
 * strip's pill has to say honestly (`sourcePill` reads `'addon'` as "Addon export, …" and
 * anything else, `'manual'` included, as "Entered by hand").
 */
async function fromPlannerCode(code: string, ctx: LoadContext): Promise<SourceResult> {
  const classSlug = code.trim().split(':')[2] ?? '';
  let talents: TalentFile;
  let classes: ClassRow[];
  let races: RaceRow[];
  try {
    const [loadedTalents, reference] = await Promise.all([
      loadTalents(ctx.treeVersion, classSlug),
      loadReference(ctx.treeVersion),
    ]);
    talents = loadedTalents;
    classes = reference.classes;
    races = reference.races;
  } catch {
    return { ok: false, message: simCopy.characterFailed };
  }
  return characterFromFs1(code, talents, classes, races, {
    kind: 'manual',
    ref: '',
    captured_at: new Date().toISOString(),
  });
}

// A plain Map, not a SvelteMap: `items` below is replaced wholesale when the class changes
// and only ever read by key, so per-key tracking would be machinery for mutations that
// never happen.
function toItemMap(items: readonly Item[]): Map<number, Item> {
  return new Map(items.map((item) => [item.id, item]));
}

/** The `source`/`ref` half of the URL bootstrap: the three sources a plain string ref
 *  identifies. `'armory'` and `'manual'` have no ref-shaped loader (armory needs a whole
 *  `CharacterPath`; manual is `code`'s job above), so a link naming either bootstraps
 *  nothing rather than guessing at one. */
function bootstrapSource(
  source: SourceKind | '',
  ref: string,
  ctx: LoadContext,
): Promise<SourceResult> | null {
  switch (source) {
    case 'addon':
      return fromAddonExport(ref, ctx);
    case 'build':
      return fromPlannerBuild(ref, ctx);
    case 'fight':
      return fromLoggedFight(ref, ctx);
    default:
      return null;
  }
}

export interface SimStoreInit {
  treeVersion: string;
  apiBase?: string;
  /** Injected by tests and by the planner, which keeps one pool for its live estimates. */
  pool?: SimPool;
  /**
   * The URL's own bootstrap, read once by the page that creates this store (SimView, from
   * `window.location.search`) and passed in rather than read from `window` here -- this
   * module stays DOM-free and unit-testable, per the header note above.
   *
   * `code` is an unsaved planner build's own FS1 code -- Planner.svelte's "Sim this build"
   * link before the build is saved (`/sim?code=…`, `encodeFS1` on the planner's live
   * state). It decodes through the same `characterFromFs1` every source funnels through,
   * adopted as a `'manual'` source: the same kind `characterFromPlanner` gives the
   * planner's own live state (character.ts), because that is exactly what this is -- a
   * build the player has not saved, not an addon export, a saved build or a logged fight.
   *
   * `source`/`ref` is the saved-build link (`/sim?source=build&ref=<id>`) and the same
   * vocabulary a logged fight or an addon push already uses. `code` wins when both are
   * present, though no link the site writes ever carries both.
   */
  code?: string;
  source?: SourceKind | '';
  ref?: string;
  /** Overrides `runOnServer`'s poll interval. Tests pass a short one; production takes the default. */
  serverPollMs?: number;
}

export function createSimStore(init: SimStoreInit) {
  const ctx = { treeVersion: init.treeVersion, apiBase: init.apiBase };

  let phase = $state<SimPhase>('idle');
  let character = $state<SimCharacter | null>(null);
  let settings = $state<SimSettings>(defaultSettings());
  let precisionId = $state<PrecisionId>('normal');
  /** `error / mean` of the figure on screen, for the progress line and the details card. */
  let relative = $state(0);
  let estimate = $state<Estimate>(EMPTY_ESTIMATE);
  let iterationsDone = $state(0);
  let iterationsTotal = $state(0);
  let result = $state<SimResult | null>(null);
  let message = $state<string | null>(null);
  // The engine's own message for the last failure. sim/request names the buff or
  // consumable id it refused, and that text is shown verbatim rather than paraphrased.
  let detail = $state('');
  // Read from `user.premium` on GET /v1/me by SimView; false until it says otherwise, so a
  // signed-out or non-premium visitor never sees a control they cannot use.
  let premium = $state(false);
  /**
   * The build's race list, for the strip's one-time picker. A logged fight records no race
   * (Task 7), so a character loaded from one arrives with `race_slug === PENDING_RACE` and
   * the player answers once before the run control enables. Loaded alongside the item file
   * in `adopt()`, from the same `loadReference` the planner conversion already needs.
   */
  let races = $state<RaceRow[]>([]);
  let talents = $state<TalentFile | null>(null);
  let actionNames = $state<ActionNames | null>(null);
  let loadedNamesFor = '';
  let items = $state<Map<number, Item>>(toItemMap([]));

  let pool: SimPool | null = init.pool ?? null;
  let handle: RunHandle | null = null;
  // Set by stop(), read at the end of run(). A worker can finish a shard between the abort
  // message being sent and the engine noticing it -- the pool's own protocol does not
  // guarantee the cancellation wins the race -- so run() checks this rather than trusting
  // that a resolved handle.result means the player's Stop click was too late to matter.
  let stopRequested = false;

  // True for the duration of a runOnServer() call, set synchronously before its first
  // `await` -- both the reentrancy guard (a second call while one is in flight is a no-op)
  // and what RunControl disables the primary button on, the server lane's own equivalent
  // of run()'s `phase = 'loading-engine'` guard.
  let serverRunning = $state(false);
  // Bumped by dispose() and by adopt() (a new character invalidates whatever server-lane
  // poll was still running for the old one). runOnServer() captures its own generation and
  // checks it against this after every `await`; once they disagree it returns without
  // touching any more state, which is what keeps a disposed store's poll loop from writing
  // to fields nothing renders any more.
  let serverRunGeneration = 0;

  function poolOnce(): SimPool {
    pool ??= createPool({});
    return pool;
  }

  /**
   * Puts the figure back to the last completed result after a cancelled run, rather than
   * leaving it at whatever the aborted run's own progress happened to report last -- which
   * can be `EMPTY_ESTIMATE`, since `run.ts`'s `execute()` reports that once, synchronously,
   * before the pool's first real tick. Without this a fast Run-then-Stop blanks a good
   * number the player never asked to discard; with it, "the number on screen is the one
   * from before this run" (this file's own comment where `run()` calls it) is actually true.
   */
  function restorePreviousResult(): void {
    if (result !== null) {
      estimate = result.dps;
      iterationsDone = result.iterations_run;
      iterationsTotal = result.request.iterations;
      relative = relativeError(result.dps);
    } else {
      estimate = EMPTY_ESTIMATE;
      iterationsDone = 0;
      iterationsTotal = 0;
      relative = 0;
    }
  }

  /** Every source funnels through here, so the failure rule lives in one place. */
  async function adopt(load: Promise<SourceResult>): Promise<void> {
    // A new character invalidates any server-lane poll still in flight for the old one --
    // see runOnServer()'s own generation check.
    serverRunGeneration += 1;
    serverRunning = false;
    phase = 'loading-character';
    message = null;
    const outcome = await load;
    if (!outcome.ok) {
      // The character already on screen stays: a bad paste is not a reason to empty a page.
      // `idle` either way -- with a character the page is back where it was, without one it
      // is back on the empty state -- and the message is what tells the two apart.
      message = outcome.message;
      phase = 'idle';
      return;
    }
    character = outcome.character;
    result = null;
    estimate = EMPTY_ESTIMATE;
    iterationsDone = 0;
    iterationsTotal = 0;
    relative = 0;
    phase = 'idle';
    try {
      const file = await loadItems(outcome.character.tree_version, outcome.character.class_slug);
      items = toItemMap(file.items);
    } catch {
      // The strip renders slot names and "Empty" without the item file; it is a nicety, and
      // a failed fetch here must not stop a player from running a sim.
      items = toItemMap([]);
    }
    try {
      // The talent file the character's point order was built against. `run()` needs a
      // TalentIndex to turn `point_order` into the engine's talents string, and this is the
      // one place that index is built, from the same file every source already validated
      // the order with -- a run can never disagree with the strip about what the tree says.
      talents = await loadTalents(outcome.character.tree_version, outcome.character.class_slug);
    } catch {
      // No file means no index means run() refuses with simCopy.failed rather than send the
      // engine a guess. The strip and the gear grid still render.
      talents = null;
    }
    try {
      // Only for the strip's race picker, and only when the character needs one -- which is
      // the logged-fight source alone. Every other source already knows the race.
      if (needsRace(outcome.character) && races.length === 0) {
        races = (await loadReference(outcome.character.tree_version)).races;
      }
    } catch {
      // No list means no picker; the strip says so rather than rendering an empty select.
      races = [];
    }
    await ensureActionNames(outcome.character.class_slug);
  }

  /**
   * The build's name table for a class, once per class. Called from `adopt()`, never from a
   * `$effect`: the fetch must not re-run because something unrelated in the store changed.
   *
   * A build with no table renders engine action keys, which is legible and honest, and is
   * not worth an error banner on a page whose numbers are all correct.
   */
  async function ensureActionNames(classSlug: string): Promise<void> {
    if (classSlug === '' || loadedNamesFor === classSlug) return;
    loadedNamesFor = classSlug;
    try {
      actionNames = await loadActionNames(init.treeVersion, classSlug);
    } catch {
      actionNames = null;
    }
  }

  // The URL's own bootstrap, kicked off once here rather than by the component: `code` wins
  // when present, otherwise `source`/`ref` dispatches to the same loaders `loadAddon`,
  // `loadBuild` and `loadFight` expose below. Neither present resolves immediately, so
  // `ready` is always safe to await. A refusal (a class mismatch, an unreachable talent, an
  // unknown race) runs through `adopt()` exactly as a pasted source does: `message` carries
  // the reason and a race-pending character still arrives, needing the strip's picker.
  const ready: Promise<void> =
    init.code !== undefined && init.code !== ''
      ? adopt(fromPlannerCode(init.code, ctx))
      : (() => {
          const load =
            init.source !== undefined && init.ref !== undefined && init.ref !== ''
              ? bootstrapSource(init.source, init.ref, ctx)
              : null;
          return load === null ? Promise.resolve() : adopt(load);
        })();

  return {
    /** Resolves once the URL's own bootstrap character, if any, has been adopted. */
    ready,
    get phase() {
      return phase;
    },
    get character() {
      return character;
    },
    get settings() {
      return settings;
    },
    get precisionId() {
      return precisionId;
    },
    /** `error / mean` of the figure on screen, for the progress line and the details card. */
    get relativeError() {
      return relative;
    },
    /**
     * Which lane the figure on screen ran on: `run()` is always `'browser'`,
     * `runOnServer()` is always `'server'` -- there is no third. Reflects whichever ran
     * most recently rather than a player-facing choice; `run()` and `runOnServer()` are
     * two separate buttons, not one control with a lane setting.
     */
    get lane(): Lane {
      return serverRunning ? 'server' : 'browser';
    },
    get estimate() {
      return estimate;
    },
    get iterationsDone() {
      return iterationsDone;
    },
    get iterationsTotal() {
      return iterationsTotal;
    },
    get result() {
      return result;
    },
    get message() {
      return message;
    },
    /** The engine's own message for the last failure, or empty. Shown verbatim, never paraphrased. */
    get detail() {
      return detail;
    },
    get actionNames() {
      return actionNames;
    },
    get items() {
      return items;
    },
    get premium() {
      return premium;
    },
    /** True for the duration of a `runOnServer()` call. RunControl disables the primary
     *  button and the server-lane button on it, the same way it does for `loadingEngine`. */
    get serverRunning() {
      return serverRunning;
    },
    /** The build's races, for the strip's picker. Empty until a character has loaded. */
    get races() {
      return races;
    },
    /** True while the loaded character still needs a race before it can be simmed. */
    get needsRace() {
      return character !== null && needsRace(character);
    },

    setPremium(value: boolean): void {
      premium = value;
    },
    /**
     * The player answering the strip's race question. It is a whole-character replacement
     * rather than a mutation, the way every other state change in this file is, and it
     * clears the result: a different race is a different sim.
     */
    setRace(slug: string): void {
      if (character === null) return;
      character = { ...character, race_slug: slug };
      result = null;
      estimate = EMPTY_ESTIMATE;
      iterationsDone = 0;
      iterationsTotal = 0;
      relative = 0;
      message = null;
      phase = 'idle';
    },
    /** One line, so compare mode reports its failures through the same alert every source uses. */
    setMessage(text: string): void {
      message = text;
    },
    setSettings(next: SimSettings): void {
      settings = next;
    },
    setPrecisionId(value: PrecisionId): void {
      precisionId = value;
    },

    loadAddon: (code: string) => adopt(fromAddonExport(code, ctx)),
    loadBuild: (id: string) => adopt(fromPlannerBuild(id, ctx)),
    loadFight: (ref: string) => adopt(fromLoggedFight(ref, ctx)),
    loadStored: (path: CharacterPath) => adopt(fromStoredCharacter(path, ctx)),

    /** Adopts a result the page was handed rather than ran: a saved sim, or a server run. */
    adoptResult(next: SimResult): void {
      result = next;
      estimate = next.dps;
      iterationsDone = next.iterations_run;
      iterationsTotal = next.request.iterations;
      relative = relativeError(next.dps);
      phase = 'done';
    },

    async run(): Promise<void> {
      if (character === null) {
        message = simCopy.noCharacter;
        return;
      }
      message = null;
      detail = '';
      stopRequested = false;
      phase = 'loading-engine';
      const plan = precisionPlan(precisionId, 'browser');
      iterationsTotal = plan.iterations;
      iterationsDone = 0;
      relative = 0;

      // The talent index comes from the file the character was loaded with, so the engine's
      // talents string is built from the same data the strip is rendering.
      const index = talents === null ? null : indexTalents(talents);
      if (index === null) {
        message = simCopy.failed;
        phase = 'error';
        return;
      }

      const input: RunInput = {
        spec: character.spec,
        source: character.source,
        character: toCharacterSpec(character, index, settings.buffs, settings.consumables),
        encounter: settings.encounter,
        iterations: plan.iterations,
        targetError: plan.targetError,
        stepIterations: plan.step,
      };

      phase = 'running';
      handle = runSim(poolOnce(), input, (update) => {
        // A shard can still report progress after stop() fires and before the engine has
        // noticed the abort message; the figure on screen must not keep moving once the
        // player has asked it to stop.
        if (stopRequested) return;
        estimate = update.estimate;
        iterationsDone = update.iterationsDone;
        iterationsTotal = update.iterationsTotal;
        relative = update.relativeError;
      });

      try {
        const finished = await handle.result;
        // handle.result can resolve with a real result even after stop(): the abort message
        // and the engine's own last tick can cross in flight, and a shard mid-tick when the
        // message arrives finishes it rather than discarding the work. The player's Stop
        // still wins -- the number on screen is the one from before this run, not a result
        // they asked to discard.
        if (stopRequested) {
          message = simCopy.stopped;
          restorePreviousResult();
          phase = result !== null ? 'done' : 'idle';
          return;
        }
        result = finished;
        phase = 'done';
      } catch (error) {
        const failure = error instanceof SimRunError ? error : null;
        message = failure?.cancelled === true ? simCopy.stopped : (failure?.message ?? simCopy.failed);
        // The engine's own words, kept beside ours: sim/request names the buff or
        // consumable id it could not map, and that is the only thing that says what to
        // change. RunControl renders it under the message, verbatim.
        detail = failure?.detail ?? '';
        if (failure?.cancelled === true) restorePreviousResult();
        phase = failure?.cancelled === true && result !== null ? 'done' : 'error';
      } finally {
        handle = null;
      }
    },

    /**
     * Runs the same request on the server lane, then polls it to a finish.
     *
     * `buildSimRequest` is the one place a `SimRequest` is assembled (`run()` above uses
     * it too), so the browser and server lanes can never disagree about what they simmed.
     * A 402 comes back from `dispatchServerSim` as a `SimApiError` already carrying
     * `simCopy.premiumRequired` (api.ts's `asSimError`); it is shown as `message` and
     * nothing else here has been touched yet, so the browser lane's own estimate and
     * result stay on screen exactly as they were -- there is nothing to roll back.
     *
     * Once dispatched, `fetchSimProgress` is polled every `serverPollMs` (2s in
     * production), updating `estimate.mean` and `iterationsDone` the way the browser
     * pool's own progress callback does, so RunControl renders both lanes identically.
     * `fetchSim` then fetches the finished result -- progress alone carries no summary.
     *
     * Two guards, both against the same class of bug -- state written by a run nothing
     * wants any more:
     *   - `serverRunning` is set synchronously, before the first `await`, so a second call
     *     that lands while one is already in flight is a no-op. Without this a fast double
     *     click dispatches (and pays for) the same premium run twice.
     *   - `generation` is this call's own snapshot of `serverRunGeneration`. `dispose()`
     *     and a new `adopt()` both bump the counter, and every state write below is guarded
     *     by `stillCurrent()`, which compares the two. A poll that outlives the component
     *     (a navigation away from /sim) or the character it was run for (a new source
     *     pasted mid-poll) then stops touching `phase`/`estimate`/`result` on its next
     *     check, rather than looping forever against a store nothing renders any more.
     */
    async runOnServer(): Promise<void> {
      if (serverRunning) return;
      if (character === null) {
        message = simCopy.noCharacter;
        return;
      }
      const index = talents === null ? null : indexTalents(talents);
      if (index === null) {
        message = simCopy.failed;
        return;
      }

      serverRunning = true;
      const generation = ++serverRunGeneration;
      const stillCurrent = (): boolean => generation === serverRunGeneration;

      try {
        message = null;
        detail = '';

        const plan = precisionPlan(precisionId, 'server');
        const request = buildSimRequest({
          spec: character.spec,
          source: character.source,
          character: toCharacterSpec(character, index, settings.buffs, settings.consumables),
          encounter: settings.encounter,
          iterations: plan.iterations,
          targetError: plan.targetError,
        });

        let simId: string;
        try {
          simId = await dispatchServerSim(request, init.apiBase);
        } catch (error) {
          if (stillCurrent()) message = error instanceof SimApiError ? error.message : simCopy.failed;
          return;
        }
        if (!stillCurrent()) return;

        iterationsTotal = plan.iterations;
        iterationsDone = 0;
        relative = 0;

        const pollMs = init.serverPollMs ?? DEFAULT_SERVER_POLL_MS;
        for (;;) {
          await delay(pollMs);
          if (!stillCurrent()) return;

          let progress: SimProgress;
          try {
            progress = await fetchSimProgress(simId, init.apiBase);
          } catch (error) {
            if (stillCurrent()) {
              message = error instanceof SimApiError ? error.message : simCopy.failed;
              phase = result !== null ? 'done' : 'error';
            }
            return;
          }
          if (!stillCurrent()) return;

          iterationsDone = progress.iterations_done;
          if (progress.dps !== undefined) estimate = { ...estimate, mean: progress.dps };

          if (progress.state === 'error') {
            message = simCopy.failed;
            phase = result !== null ? 'done' : 'error';
            return;
          }
          if (progress.state === 'done') {
            try {
              const finished = await fetchSim(simId, init.apiBase);
              if (stillCurrent()) {
                result = finished;
                estimate = finished.dps;
                iterationsDone = finished.iterations_run;
                iterationsTotal = finished.request.iterations;
                relative = relativeError(finished.dps);
                phase = 'done';
              }
            } catch (error) {
              if (stillCurrent()) {
                message = error instanceof SimApiError ? error.message : simCopy.failed;
                phase = result !== null ? 'done' : 'error';
              }
            }
            return;
          }
        }
      } finally {
        // Only clears the flag this call itself set: dispose()/adopt() already cleared it
        // (and moved the generation past this call's own) when they are the reason this is
        // running, and a newer runOnServer() call may have set it again by now.
        if (stillCurrent()) serverRunning = false;
      }
    },

    stop(): void {
      stopRequested = true;
      handle?.cancel();
    },

    /**
     * Saves the last finished result, optionally under a title -- the contract's `sims.
     * title` column, pre-filled by the caller with `settingsLabel(store.settings)` and
     * editable before the press. Null on failure, without touching `message`: a save
     * failure is the save form's own concern (`simCopy.saveFailed` beside its button, per
     * the design), not the run control's -- setting the shared field here would raise a
     * second, unrelated alert next to a run that did not fail.
     */
    async save(title?: string): Promise<string | null> {
      if (result === null) return null;
      try {
        return await saveSim(result, init.apiBase, title ?? '');
      } catch {
        return null;
      }
    },

    dispose(): void {
      handle?.cancel();
      serverRunGeneration += 1;
      serverRunning = false;
      pool?.terminate();
      pool = null;
    },
  };
}
