// web/src/lib/account/main-character.ts
// The one character the hub points at on arrival (spec 2026-09-22 §3.1): the simmable
// character (one the site holds a build for) with the newest build.captured_at, ties
// broken by level; with nothing simmable, the highest level; with no characters at all,
// null. A pure function of the list `/v1/me` already returned, so the hub's own effect
// just calls it and writes the result -- no fetch of its own.
import type { CurrentCharacter } from '../current-character';
import { characterSpecName } from '../character-spec';
import { specName } from '../guild/spec-name';
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

/**
 * The account's main: the character the player chose (`main_character_key`) when it is
 * still one of theirs; otherwise the site's own guess until they choose, the simmable
 * character with the freshest build, then the highest level. Everything else is an alt.
 */
export function mainCharacter(characters: readonly MeCharacter[], mainKey?: string): MeCharacter | null {
  if (characters.length === 0) return null;
  const chosen = mainKey === undefined ? undefined : characters.find((c) => c.key === mainKey);
  if (chosen !== undefined) return chosen;
  const simmable = characters.filter((character) => character.build !== undefined);
  if (simmable.length > 0) return simmable.reduce(betterSimmable);
  return [...characters].reduce(betterByLevel);
}

/** The current-character pointer for the hub's own arrival write (spec §3.1: "sets the
 *  current-character pointer to it"). Always an `'armory'`-kind pointer, the same kind
 *  every other stored-character load writes (`sources.ts`'s `fromStoredCharacter`) --
 *  `/sim`, `/planner` and the character page all restore one by key. */
export function pointerForCharacter(character: MeCharacter): CurrentCharacter {
  const classSlug = character.class?.toLowerCase() ?? '';
  // The class as the game names it ("Survival Hunter"), whatever case the API sent.
  const className = character.class === undefined ? undefined : specName(character.class);
  const spec = characterSpecName(character);
  const specClass = [spec, className].filter((part) => part !== undefined).join(' ');
  const label = specClass === '' ? character.name : `${character.name} · ${specClass}`;
  return {
    source: 'armory',
    ref: character.key,
    label,
    classSlug,
    savedAt: new Date().toISOString(),
  };
}
