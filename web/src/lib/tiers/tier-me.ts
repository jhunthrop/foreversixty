// web/src/lib/tiers/tier-me.ts
// The signed-in half of the tier list, loaded only after the page has settled and only when a
// session is hinted: which of the tier list's specs the visitor's selected character is.
import { fetchMeOnce } from '../account/api';
import { selectedCharacter } from '../account/hero-character';
import { readCurrent } from '../current-character';
import { specViewKeyForCharacter } from './tier-callout';

/** The `<class>/<spec>` key of the selected character, or `undefined` when signed out, when the
 *  read fails, or while its class or spec is not known yet. */
export async function findCharacterKey(): Promise<string | undefined> {
  const me = await fetchMeOnce().catch(() => null);
  if (me === null) return undefined;
  const character = selectedCharacter(readCurrent(), me);
  return character === null ? undefined : specViewKeyForCharacter(character);
}
