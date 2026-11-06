// web/src/lib/tiers/tier-me.ts
// The signed-in half of the tier list, loaded only after the page has settled and only when a
// session is hinted: which of the tier list's specs the visitor's selected character is.
import { fetchMeOnce, type Me } from '../account/api';
import { selectedCharacter } from '../account/hero-character';
import { readCurrent, type CurrentCharacter } from '../current-character';
import { specViewKeyForCharacter } from './tier-callout';

/** What the answer slot shows: the spec's place, a character whose spec is not known yet, a
 *  failed read, or nothing (signed out, or no character). */
export type TierMe =
  { kind: 'spec'; key: string } | { kind: 'unknown'; name: string } | { kind: 'error' } | { kind: 'none' };

/** The answer for an account read: pure, so the tests feed it fixtures. */
export function tierMeFrom(me: Me | null, current: CurrentCharacter | null): TierMe {
  if (me === null) return { kind: 'none' };
  const character = selectedCharacter(current, me);
  if (character === null) return { kind: 'none' };
  const key = specViewKeyForCharacter(character);
  return key === undefined ? { kind: 'unknown', name: character.name } : { kind: 'spec', key };
}

export async function findTierMe(): Promise<TierMe> {
  try {
    return tierMeFrom(await fetchMeOnce(), readCurrent());
  } catch {
    return { kind: 'error' };
  }
}
