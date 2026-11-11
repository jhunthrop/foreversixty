// web/src/lib/ui/copy-state.svelte.ts
// The reactive face of lib/ui/copy-feedback.ts for a component whose copy control is not the
// shared CopyButton (a row of Nudge buttons, a link inside a sentence): `state` is idle, copied
// or failed, and `key` names which of several things was copied, so a list lights one row.
// The owner disposes it on destroy: `$effect(() => () => copy.dispose())`.
import { uiCopy } from './copy';
import { createCopyFeedback, writeClipboard, type CopyFeedback, type CopyFeedbackState } from './copy-feedback';

export class CopyState {
  state = $state<CopyFeedbackState>('idle');
  key = $state<string | null>(null);
  readonly #feedback: CopyFeedback;

  constructor(writeText: (text: string) => Promise<void> = writeClipboard) {
    this.#feedback = createCopyFeedback((next) => {
      this.state = next;
      if (next === 'idle') this.key = null;
    }, writeText);
  }

  /** Writes `text`; `key` marks which row or control asked. */
  copy(text: string, key: string | null = null): Promise<void> {
    this.key = key;
    return this.#feedback.copy(text);
  }

  /** True while `key` (or the single control, with no key) shows "Copied". */
  copied(key: string | null = null): boolean {
    return this.state === 'copied' && this.key === key;
  }

  /** True when the last try for `key` was refused. */
  failed(key: string | null = null): boolean {
    return this.state === 'failed' && this.key === key;
  }

  /** The control's words: `idle`, then "Copied", or "Copy failed" after a refusal. */
  label(idle: string, key: string | null = null): string {
    if (this.copied(key)) return uiCopy.copied;
    if (this.failed(key)) return uiCopy.copyFailed;
    return idle;
  }

  dispose(): void {
    this.#feedback.dispose();
  }
}
