// web/src/lib/guild/report-copy.svelte.ts
// `Copy for Discord` (attendance spec 4): the clipboard write, `Copied` for 4 s, and on failure the
// text kept so the page can show it selected in a read-only textarea. The clipboard and the timer
// come in, so the tests drive them.
import { COPIED_VISIBLE_MS } from './attendance-copy';

export type ReportCopyState = 'idle' | 'copied' | 'failed';

export class ReportCopy {
  state = $state<ReportCopyState>('idle');
  /** The text the failure fallback shows. */
  text = $state('');

  private readonly writeText: (text: string) => Promise<void>;
  private timer: ReturnType<typeof setTimeout> | undefined;

  constructor(writeText: (text: string) => Promise<void>) {
    this.writeText = writeText;
  }

  copy = async (text: string): Promise<void> => {
    clearTimeout(this.timer);
    this.text = text;
    try {
      await this.writeText(text);
    } catch {
      this.state = 'failed';
      return;
    }
    this.state = 'copied';
    this.timer = setTimeout(() => {
      this.state = 'idle';
    }, COPIED_VISIBLE_MS);
  };

  stop(): void {
    clearTimeout(this.timer);
  }
}
