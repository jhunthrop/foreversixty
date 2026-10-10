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
  /** Called with the `restored` flag after each load. */
  onSettled?: (restored: boolean) => void;
  /** Called with true while a load runs and false once the queue is empty. */
  onBusy?: (busy: boolean) => void;
}

const NO_URL = { code: '', source: '', ref: '' } as const;

/**
 * Subscribes; returns the unsubscribe. Loads run one after another in the order the choices
 * were made, so the last choice is always the one left on screen.
 */
export function followCurrentCharacter(options: FollowOptions): () => void {
  let queue: Promise<void> = Promise.resolve();
  let pending = 0;
  const load = async (pointer: CurrentCharacter): Promise<void> => {
    const restored = await runBootstrapRestore(options.loaders, NO_URL, pointer, options.characterLoaded, {
      storeHandlesUrl: false,
    });
    options.onSettled?.(restored);
  };
  const stop = onCurrentCharacterChange(({ fromPageLoad }) => {
    const pointer = readCurrent();
    if (fromPageLoad || pointer === null || options.pinned?.() === true) return;
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
