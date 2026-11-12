// web/src/lib/addon/copy-for-game.ts
// What "Copy for the game" does when pressed, kept out of CopyForGame.svelte so it runs
// under vitest without a DOM: the shared clipboard feedback (lib/ui/copy-feedback.ts, "Copied"
// for two seconds), with the game's own words on the button; when the clipboard refuses, the
// component shows the code to select by hand. An empty code (no build yet) does nothing.
import { createCopyFeedback, type CopyFeedbackState } from '../ui/copy-feedback';
import { addonCopy } from './copy';

export type CopyForGameState = CopyFeedbackState;

/** The button's own label. A failure is said beside the button, not on it. */
export function copyForGameLabel(state: CopyForGameState): string {
  return state === 'copied' ? addonCopy.copiedForGame : addonCopy.copyForGame;
}

export interface CopyForGame {
  copy(code: string): Promise<void>;
  /** Clears a pending revert, so an unmounted button is never written to. */
  dispose(): void;
}

export function createCopyForGame(
  onchange: (state: CopyForGameState) => void,
  writeText?: (text: string) => Promise<void>,
): CopyForGame {
  const feedback = createCopyFeedback(onchange, writeText);
  return {
    copy: (code) => (code === '' ? Promise.resolve() : feedback.copy(code)),
    dispose: feedback.dispose,
  };
}
