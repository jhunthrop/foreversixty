// web/src/lib/character-selector/paste.ts
// "Paste an export" inside the selector (spec 3.F): the same decoder the planner and /addon
// paste box use, then the pointer it earns. An export that names a character already on the
// account becomes that account character ('armory', by key); anything else stays a pasted
// export ('code'). Loaded on demand: the decoder is not on the page-load path of every page.
import { fetchMeOnce, type MeCharacter } from '../account/api';
import { sameCharacter } from '../account/hero-character';
import { pointerForCharacter } from '../account/main-character';
import { announcePageLoadedCharacter, writeCurrent, type CurrentCharacter } from '../current-character';
import { decodeFS1 } from '../planner/fs1';

export { sameCharacter };

export type PasteResult =
  { ok: true; pointer: CurrentCharacter; account: MeCharacter | null } | { ok: false };

function capitalise(slug: string): string {
  return slug.charAt(0).toUpperCase() + slug.slice(1);
}

/** Decodes `input`; a failure carries nothing, the selector shows its one fixed sentence. */
export function resolvePastedExport(
  input: string,
  accountCharacters: readonly MeCharacter[],
  now: Date,
): PasteResult {
  const code = input.trim();
  const decoded = decodeFS1(code);
  if (!decoded.ok) return { ok: false };
  const { classSlug } = decoded.build;
  const name = decoded.build.character?.name;
  const account =
    name === undefined
      ? null
      : (accountCharacters.find((character) => sameCharacter(character, name, classSlug)) ?? null);
  if (account !== null) {
    return { ok: true, account, pointer: { ...pointerForCharacter(account), savedAt: now.toISOString() } };
  }
  const className = capitalise(classSlug);
  return {
    ok: true,
    account: null,
    pointer: {
      source: 'code',
      ref: code,
      label: name === undefined ? className : `${name} · ${className}`,
      classSlug,
      savedAt: now.toISOString(),
    },
  };
}

/**
 * A player pasting an export into a box on a page (the planner's Import, /addon, /setup) is
 * choosing it: an export that names a character becomes the current one (the account's own
 * when it holds one by that name, else a pasted export). One that names nobody is read and
 * changes nothing -- a class word is never made the pointer. Never throws.
 */
export async function recordPastedExport(code: string, storage?: Storage): Promise<void> {
  const decoded = decodeFS1(code.trim());
  if (!decoded.ok || decoded.build.character === undefined) return;
  // The export is the current character at once; the account, once it answers, may know it by key.
  const first = resolvePastedExport(code, [], new Date());
  if (first.ok) {
    writeCurrent(first.pointer, storage);
    announcePageLoadedCharacter();
  }
  const me = await fetchMeOnce().catch(() => null);
  const held = resolvePastedExport(code, me?.characters ?? [], new Date());
  if (!held.ok || held.account === null) return;
  writeCurrent(held.pointer, storage);
  announcePageLoadedCharacter();
}
