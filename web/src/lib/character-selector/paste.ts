// web/src/lib/character-selector/paste.ts
// "Paste an export" inside the selector (spec 3.F): the same decoder the planner and /addon
// paste box use, then the pointer it earns. An export that names a character already on the
// account becomes that account character ('armory', by key); anything else stays a pasted
// export ('code'). Loaded on demand: the decoder is not on the page-load path of every page.
import type { MeCharacter } from '../account/api';
import { characterSlug } from '../characters';
import { pointerForCharacter } from '../account/main-character';
import type { CurrentCharacter } from '../current-character';
import { decodeFS1 } from '../planner/fs1';

export type PasteResult =
  { ok: true; pointer: CurrentCharacter; account: MeCharacter | null } | { ok: false };

function capitalise(slug: string): string {
  return slug.charAt(0).toUpperCase() + slug.slice(1);
}

/** The account character an export or a pasted pointer names: same name (a key slug counts) and class, case-insensitive. */
export function sameCharacter(character: MeCharacter, name: string, classSlug: string): boolean {
  return (
    characterSlug(character.name) === characterSlug(name) && (character.class ?? '').toLowerCase() === classSlug
  );
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
