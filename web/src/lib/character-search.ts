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
