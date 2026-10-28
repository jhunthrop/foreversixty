// web/src/lib/ui/copy-feedback.ts
// The copy-to-clipboard feedback the shared link field shows: "Copied" for two seconds
// after a successful write, a failure line after a rejected one. Pure and injectable so
// the timing is testable without a DOM.

export const COPIED_VISIBLE_MS = 2000;

export type CopyFeedbackState = 'idle' | 'copied' | 'failed';

export interface CopyFeedback {
  copy: (text: string) => Promise<void>;
  dispose: () => void;
}

export function createCopyFeedback(
  writeText: (text: string) => Promise<void>,
  onState: (state: CopyFeedbackState) => void,
  visibleMs: number = COPIED_VISIBLE_MS,
): CopyFeedback {
  let timer: ReturnType<typeof setTimeout> | undefined;
  return {
    async copy(text) {
      try {
        await writeText(text);
        onState('copied');
        clearTimeout(timer);
        timer = setTimeout(() => onState('idle'), visibleMs);
      } catch {
        clearTimeout(timer);
        onState('failed');
      }
    },
    dispose() {
      clearTimeout(timer);
    },
  };
}
