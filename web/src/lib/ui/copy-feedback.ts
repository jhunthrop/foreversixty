// web/src/lib/ui/copy-feedback.ts
// The one clipboard write and its feedback (component audit 2026-11-10 cluster 13, owner ruling
// 2026-11-10: "Copied" clears after about two seconds, everywhere). Every copy control goes through
// here: `writeClipboard` is the site's only navigator.clipboard call (state-guard.test.ts), and
// `createCopyFeedback` reports copied, then idle after COPIED_VISIBLE_MS, or failed (until the next
// try) when the clipboard refuses. Pure and injectable so the timing is testable without a DOM.

export const COPIED_VISIBLE_MS = 2000;

export type CopyFeedbackState = 'idle' | 'copied' | 'failed';

export interface CopyFeedback {
  copy: (text: string) => Promise<void>;
  dispose: () => void;
}

/** The clipboard write. Rejects where the browser has no clipboard or refuses it. */
export function writeClipboard(text: string): Promise<void> {
  if (typeof navigator === 'undefined' || navigator.clipboard === undefined) {
    return Promise.reject(new Error('clipboard unavailable'));
  }
  return navigator.clipboard.writeText(text);
}

export function createCopyFeedback(
  onState: (state: CopyFeedbackState) => void,
  writeText: (text: string) => Promise<void> = writeClipboard,
): CopyFeedback {
  let timer: ReturnType<typeof setTimeout> | undefined;
  return {
    async copy(text) {
      clearTimeout(timer);
      try {
        await writeText(text);
        onState('copied');
        timer = setTimeout(() => onState('idle'), COPIED_VISIBLE_MS);
      } catch {
        // A refusal stands until the next try: the caller's failure line says what to do instead.
        onState('failed');
      }
    },
    dispose() {
      clearTimeout(timer);
    },
  };
}
