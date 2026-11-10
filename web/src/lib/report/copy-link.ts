// web/src/lib/report/copy-link.ts
// Copy link, the report's one share action (This pull on a desktop, the band's Report actions
// menu on a phone): the page's own url, which carries the fight, tab, window and source, so a
// pasted link opens exactly this view. One implementation for both places.
import { reportCopy } from './copy';

/** How long the button reads Copied (or Copy failed) before it returns to Copy link. */
export const COPIED_FOR_MS = 2000;

/** Writes `url` to the clipboard; resolves to the button's word for the outcome. */
export async function copyReportLink(
  url: string,
  clipboard: Pick<Clipboard, 'writeText'> | null = globalThis.navigator?.clipboard ?? null,
): Promise<string> {
  try {
    if (clipboard === null) throw new Error('no clipboard');
    await clipboard.writeText(url);
    return reportCopy.copied;
  } catch {
    return reportCopy.copyFailed;
  }
}
