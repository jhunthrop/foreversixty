// web/src/lib/character-search.ts
// The home page's character search (Task L8.2 ruling 7): what a player types -- `Name`, or
// `Name-Realm` the way the game and the addon write a unit -- turned into the profile's
// address through characters.ts's own characterHref, so the search and every other link to
// a character build the same path.
import {
  characterHref,
  parseCharacterPath,
  rulesetFromRealm,
  splitUnitName,
  type Region,
  type Ruleset,
} from './characters';

/** The region and ruleset a search starts on, the same pair the site's placeholder character
 *  and the addon fixtures use (lib/character-header/placeholder.ts). */
export const DEFAULT_REGION: Region = 'us';
export const DEFAULT_RULESET: Ruleset = 'normal';

/**
 * The profile a search names, or null when nothing a character could be called was typed.
 * A `-Realm` suffix picks the ruleset when the realm's name says one (characters.ts's
 * rulesetFromRealm: `-Hardcore`, `-PvP`, `Classic Beta PvE 2`); a hyphen that names no
 * ruleset is part of the name (`Jean-Luc`), since the game allows hyphenated names.
 */
export function characterSearchHref(
  query: string,
  region: Region = DEFAULT_REGION,
  ruleset: Ruleset = DEFAULT_RULESET,
): string | null {
  const typed = query.trim();
  if (typed === '') return null;
  const { name, segment } = splitUnitName(typed);
  const fromRealm = segment === '' ? '' : rulesetFromRealm(segment);
  const href =
    fromRealm === '' ? characterHref(region, ruleset, typed) : characterHref(region, fromRealm, name);
  return parseCharacterPath(href) === null ? null : href;
}

/** The one search row a lookup should open, from GET /v1/characters/search's rows. */
export interface LookupRow {
  key: string;
  region: string;
  ruleset: string;
  name: string;
}

/**
 * The character a lookup meant (first-real-user feedback, 2026-11-04: a phone user typed their
 * name, kept the form's Normal default and got "not found" for a character the site had on
 * PvP). Among the search rows whose name is what was typed, or starts with it as a whole word
 * ("Frickin" for "Frickin Laserbeams"), the one on the chosen region and ruleset wins, then the
 * chosen region, then the only one there is. Undefined when nothing or several match.
 */
export function pickLookupMatch(
  rows: readonly LookupRow[],
  typed: string,
  region: string,
  ruleset: string,
): LookupRow | undefined {
  const want = typed.trim().toLowerCase();
  if (want === '') return undefined;
  const named = rows.filter((row) => {
    const name = row.name.toLowerCase();
    return name === want || name.startsWith(`${want} `);
  });
  const exact = named.filter((row) => row.name.toLowerCase() === want);
  const pool = exact.length > 0 ? exact : named;
  const here = pool.filter((row) => row.region === region && row.ruleset === ruleset);
  if (here.length === 1) return here[0];
  const sameRegion = pool.filter((row) => row.region === region);
  if (sameRegion.length === 1) return sameRegion[0];
  return pool.length === 1 ? pool[0] : undefined;
}
