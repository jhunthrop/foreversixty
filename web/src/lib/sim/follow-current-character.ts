// web/src/lib/sim/follow-current-character.ts
// A /sim* page that follows the header character selector: when the pointer changes to a
// character the page did not load itself, load that character the way a fresh visit would
// (`runBootstrapRestore` with an empty URL), in place. Shared by SimView.svelte and
// ToolsView.svelte so the two cannot drift.
import { onCurrentCharacterChange, readCurrent, type CurrentCharacter } from '../current-character';
import { runBootstrapRestore, type CharacterLoaders } from './character-bootstrap';

export interface FollowOptions {
  loaders: CharacterLoaders;
  /** Read after each load settles: whether the page now holds a character. */
  characterLoaded: () => boolean;
  /** True on a page that shows a fixed request (a saved sim or share link): it never follows. */
  pinned?: () => boolean;
  /** True while the page holds work a session-default heal must not replace (a run in progress, a
   *  `?code=` build on screen). A player's own selection still follows. */
  holdsWork?: () => boolean;
  /** Called with true while a load runs and false once the queue is empty. */
  onBusy?: (busy: boolean) => void;
}

const NO_URL = { code: '', source: '', ref: '' } as const;

/** Phases (the simulator's and the bulk tools') in which a run, or the engine for one, is under way. */
const BUSY_PHASES: ReadonlySet<string> = new Set(['loading-engine', 'counting', 'running']);

/** What a store reports for `simHoldsWork`; both the simulator's and the bulk tools' stores fit. */
export interface SimWorkState {
  phase: string;
  serverRunning: boolean;
}

/**
 * Whether the page holds work a session-default heal must not replace: a run (or its engine
 * load) in progress, a server run, or a `?code=` build on screen.
 */
export function simHoldsWork(state: SimWorkState, bootstrapCode: string): boolean {
  return BUSY_PHASES.has(state.phase) || state.serverRunning || bootstrapCode !== '';
}

/**
 * Subscribes; returns the unsubscribe. Loads run one after another in the order the choices
 * were made, so the last choice is always the one left on screen.
 */
export function followCurrentCharacter(options: FollowOptions): () => void {
  let queue: Promise<void> = Promise.resolve();
  let pending = 0;
  const load = async (pointer: CurrentCharacter): Promise<void> => {
    await runBootstrapRestore(options.loaders, NO_URL, pointer, options.characterLoaded, {
      storeHandlesUrl: false,
    });
  };
  const stop = onCurrentCharacterChange(({ fromPageLoad, fromSessionHeal }) => {
    const pointer = readCurrent();
    if (fromPageLoad || pointer === null || options.pinned?.() === true) return;
    if (fromSessionHeal && options.holdsWork?.() === true) return;
    pending += 1;
    options.onBusy?.(true);
    queue = queue
      .then(() => load(pointer))
      .catch(() => undefined)
      .finally(() => {
        pending -= 1;
        if (pending === 0) options.onBusy?.(false);
      });
  });
  return stop;
}
