// web/src/lib/account/main-character.ts
// The one character the hub points at on arrival (spec 2026-09-22 §3.1): the simmable
// character (one the site holds a build for) with the newest build.captured_at, ties
// broken by level; with nothing simmable, the highest level; with no characters at all,
// null. A pure function of the list `/v1/me` already returned, so the hub's own effect
// just calls it and writes the result -- no fetch of its own.
import type { CurrentCharacter } from '../current-character';
import type { MeCharacter } from './api';

function levelOf(character: MeCharacter): number {
  return character.level ?? 0;
}

function betterSimmable(a: MeCharacter, b: MeCharacter): MeCharacter {
  const atA = a.build?.captured_at ?? '';
  const atB = b.build?.captured_at ?? '';
  if (atA !== atB) return atA > atB ? a : b;
  return levelOf(b) > levelOf(a) ? b : a;
}

function betterByLevel(a: MeCharacter, b: MeCharacter): MeCharacter {
  return levelOf(b) > levelOf(a) ? b : a;
}

export function mainCharacter(characters: readonly MeCharacter[]): MeCharacter | null {
  if (characters.length === 0) return null;
  const simmable = characters.filter((character) => character.build !== undefined);
  if (simmable.length > 0) return simmable.reduce(betterSimmable);
  return [...characters].reduce(betterByLevel);
}

/** The current-character pointer for the hub's own arrival write (spec §3.1: "sets the
 *  current-character pointer to it"). Always an `'armory'`-kind pointer, the same kind
 *  every other stored-character load writes (`sources.ts`'s `fromStoredCharacter`) --
 *  `/sim` and the character page already know how to restore one; `/planner` does not
 *  (documented in `current-character-planner.ts`), which is an existing, intentional
 *  limit this pointer does not change. */
export function pointerForCharacter(character: MeCharacter): CurrentCharacter {
  const classSlug = character.class?.toLowerCase() ?? '';
  const label = character.class === undefined ? character.name : `${character.name} · ${character.class}`;
  return {
    source: 'armory',
    ref: character.key,
    label,
    classSlug,
    savedAt: new Date().toISOString(),
  };
}
