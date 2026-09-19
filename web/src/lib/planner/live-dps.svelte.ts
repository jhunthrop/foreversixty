// web/src/lib/planner/live-dps.svelte.ts
// A 500-iteration sim behind every planner edit, and the three rules that make it usable.
//
//   1. The pool is created on the first request, never in this factory. /planner.html has a
//      0.90 performance budget and a 100 ms total-blocking-time budget in lighthouserc.json;
//      a 4 MB wasm module on page load ends both. Nothing on the planner's critical path may
//      reach engine.ts, which is why sim.worker.ts -- a separate Vite chunk -- is the only
//      file that imports it.
//   2. A burst of edits is one run. Clicking through five talents must not queue five runs.
//   3. The last figure stays on screen, dimmed, while the next one is computed. A number
//      that blanks and reappears on every click is worse than a stale number that dims.
import type { TalentIndex } from './rules';
import { EMPTY_ESTIMATE } from '../sim/estimate';
import { runSim, SimRunError, type RunHandle } from '../sim/run';
import { defaultSettings } from '../sim/settings';
import { toCharacterSpec, type SimCharacter } from '../sim/character';
import { simCopy } from '../sim/copy';
import { ITERATIONS, type Estimate } from '../sim/types';
import { createPool, type SimPool } from '../sim/worker';

export const LIVE_DEBOUNCE_MS = 350;

export type LiveState = 'off' | 'pending' | 'running' | 'ready' | 'error';

export function createLiveDps(init: { pool?: SimPool; debounceMs?: number } = {}) {
  const debounceMs = init.debounceMs ?? LIVE_DEBOUNCE_MS;

  let state = $state<LiveState>('off');
  let estimate = $state<Estimate>(EMPTY_ESTIMATE);
  let iterationsRun = $state(0);
  let message = $state<string | null>(null);

  let pool: SimPool | null = init.pool ?? null;
  let handle: RunHandle | null = null;
  let timer: ReturnType<typeof setTimeout> | null = null;

  function poolOnce(): SimPool {
    pool ??= createPool({});
    return pool;
  }

  async function start(character: SimCharacter, index: TalentIndex): Promise<void> {
    state = 'running';
    const settings = defaultSettings();
    handle = runSim(
      poolOnce(),
      {
        spec: character.spec,
        source: character.source,
        character: toCharacterSpec(character, index, settings.buffs, settings.consumables),
        encounter: settings.encounter,
        iterations: ITERATIONS.live,
      },
      () => {
        // The planner shows one figure at the end, not a ticking one: a number that moves
        // while a player is reading their talent tree is noise, and 500 iterations is fast
        // enough that there is nothing to fill.
      },
    );

    try {
      const result = await handle.result;
      estimate = result.dps;
      iterationsRun = result.iterations_run;
      message = null;
      state = 'ready';
    } catch (error) {
      // A cancel is this module's own doing -- a newer edit arrived -- and is not a failure
      // the player should be told about; the newer run will set the state.
      if (error instanceof SimRunError && error.cancelled) return;
      message = simCopy.liveDpsFailed;
      state = 'error';
    } finally {
      handle = null;
    }
  }

  return {
    get state() {
      return state;
    },
    get estimate() {
      return estimate;
    },
    get iterationsRun() {
      return iterationsRun;
    },
    get message() {
      return message;
    },

    request(character: SimCharacter | null, index: TalentIndex | null): void {
      if (timer !== null) clearTimeout(timer);
      handle?.cancel();
      if (character === null || index === null) {
        state = 'off';
        return;
      }
      state = 'pending';
      timer = setTimeout(() => {
        timer = null;
        void start(character, index);
      }, debounceMs);
    },

    dispose(): void {
      if (timer !== null) clearTimeout(timer);
      timer = null;
      handle?.cancel();
      pool?.terminate();
      pool = null;
    },
  };
}

export type LiveDps = ReturnType<typeof createLiveDps>;
