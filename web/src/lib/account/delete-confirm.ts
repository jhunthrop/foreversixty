// web/src/lib/account/delete-confirm.ts
// The account deletion dialog's rules (DELETE /v1/me, operation deleteAccount): what the person
// types to confirm, and the client-side check that enables the button. The server takes the
// account's BattleTag or the word "delete", trimmed and in any case; the check here is the same
// rule, so the button never enables for something the server would refuse.
import { accountPageCopy } from './account-page-copy';

const { confirmWord } = accountPageCopy.deleteDialog;

/** The word the field asks for: the account's BattleTag, or "delete" for an account with none. */
export function confirmWordFor(battletag: string | null | undefined): string {
  const tag = battletag?.trim() ?? '';
  return tag === '' ? confirmWord : tag;
}

/** Whether what was typed confirms: the BattleTag or "delete", trimmed, case-insensitive. */
export function confirmMatches(typed: string, battletag: string | null | undefined): boolean {
  const wanted = typed.trim().toLowerCase();
  if (wanted === '') return false;
  const tag = battletag?.trim().toLowerCase() ?? '';
  return wanted === confirmWord || (tag !== '' && wanted === tag);
}

/** One dialog line, split around the words that link (when it has a link). */
export interface LineParts {
  before: string;
  link?: { label: string; href: string };
  after: string;
}

export function lineParts(line: { text: string; link?: { label: string; href: string } }): LineParts {
  const at = line.link === undefined ? -1 : line.text.indexOf(line.link.label);
  if (line.link === undefined || at < 0) return { before: line.text, after: '' };
  return {
    before: line.text.slice(0, at),
    link: line.link,
    after: line.text.slice(at + line.link.label.length),
  };
}

/** What a deletion attempt came to, for the dialog and the page that owns it. */
export type DeleteOutcome =
  | { kind: 'deleted' }
  /** 401: the session was already gone; the page treats the visitor as signed out. */
  | { kind: 'signed-out' }
  /** 400 confirm_mismatch or invalid: the line under the field. */
  | { kind: 'mismatch' }
  /** Anything else (409 guild_obligation or staff_account, 429, 503 billing_unavailable): the API's
   *  own message, shown verbatim. */
  | { kind: 'refused'; message: string };

/** Runs the deletion and sorts its answer. `fallback` is the message for a refusal with none. */
export async function deleteOutcome(run: () => Promise<unknown>, fallback: string): Promise<DeleteOutcome> {
  try {
    await run();
    return { kind: 'deleted' };
  } catch (thrown) {
    const status = typeof thrown === 'object' && thrown !== null && 'status' in thrown ? thrown.status : 0;
    if (status === 401) return { kind: 'signed-out' };
    if (status === 400) return { kind: 'mismatch' };
    const message = thrown instanceof Error && thrown.message !== '' ? thrown.message : fallback;
    return { kind: 'refused', message };
  }
}
