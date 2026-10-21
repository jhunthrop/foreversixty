// web/src/lib/account/adopt-held.ts
// A page that loads an export naming one of the account's own characters makes that character
// current -- the account holds it, so it is the player's own and the header should say so. Any
// other named export is somebody else's (a guild member's gear, another player's character
// page) and a load never changes the pointer for it: the pointer changes only when the player
// chooses (current-character.ts).
import { announcePageLoadedCharacter, readCurrent, writeCurrent } from '../current-character';
import { fetchMeOnce, type MeCharacter } from './api';
import { sameCharacter } from './hero-character';
import { pointerForCharacter } from './main-character';

/** The account character `name`/`classSlug` names (whether or not the pointer needed to move), or null. */
export async function adoptHeldCharacter(
  name: string,
  classSlug: string,
  storage?: Storage,
): Promise<MeCharacter | null> {
  if (name === '') return null;
  const me = await fetchMeOnce().catch(() => null);
  const held = me?.characters.find((row) => sameCharacter(row, name, classSlug));
  if (held === undefined) return null;
  const current = readCurrent(storage);
  if (current?.source === 'armory' && current.ref === held.key) return held;
  writeCurrent(pointerForCharacter(held), storage);
  announcePageLoadedCharacter();
  return held;
}
