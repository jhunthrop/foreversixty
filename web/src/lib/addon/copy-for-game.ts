// web/src/lib/addon/copy-for-game.ts
// What "Copy for the game" does when pressed, kept out of CopyForGame.svelte so it runs
// under vitest without a DOM (this project renders components with Svelte's server entry
// only): copy the code, say so for 3 s, then go back; or, when the clipboard refuses, say
// that and let the component show the code to select by hand.
import { addonCopy } from './copy';

/** How long the button reads `addonCopy.copiedForGame` after a copy. */
export const COPIED_FOR_GAME_MS = 3000;

export type CopyForGameState = 'idle' | 'copied' | 'failed';

/** The button's own label. A failure is said beside the button, not on it. */
export function copyForGameLabel(state: CopyForGameState): string {
  return state === 'copied' ? addonCopy.copiedForGame : addonCopy.copyForGame;
}

export interface CopyForGame {
  copy(code: string): Promise<void>;
  /** Clears a pending revert, so an unmounted button is never written to. */
  dispose(): void;
}

export function createCopyForGame(onchange: (state: CopyForGameState) => void): CopyForGame {
  let timer: ReturnType<typeof setTimeout> | null = null;

  const clear = (): void => {
    if (timer !== null) clearTimeout(timer);
    timer = null;
  };

  return {
    async copy(code: string): Promise<void> {
      if (code === '') return;
      clear();
      try {
        // Read at click time, never at module load: `navigator` does not exist during SSR,
        // and `navigator.clipboard` is undefined on an insecure page -- both land in the
        // catch below as an ordinary failure.
        await navigator.clipboard.writeText(code);
      } catch {
        onchange('failed');
        return;
      }
      onchange('copied');
      timer = setTimeout(() => {
        timer = null;
        onchange('idle');
      }, COPIED_FOR_GAME_MS);
    },
    dispose: clear,
  };
}
