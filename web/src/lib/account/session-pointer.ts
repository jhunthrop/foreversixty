// web/src/lib/account/session-pointer.ts
// The pointer at the start of a signed-in session. When `/v1/me` first answers for a browser
// session and the stored pointer is absent, a nameless pasted export, or an account character this
// account does not hold (a previous player's, left in the browser), the account's main becomes the
// current character -- whichever way the player signed in (the Battle.net return, the email link
// and the dev sign-in all end the same). After that the pointer changes only when the player
// chooses; signing out clears it.
import {
  CURRENT_CHARACTER_CHANGED,
  clearCurrent,
  readCurrent,
  writeCurrent,
  type CurrentCharacter,
} from '../current-character';
import type { Me } from './api';
import { pointerNameOf } from './hero-character';
import { mainCharacter, pointerForCharacter } from './main-character';

const HEALED_KEY = 'fs.pointerHealed';

/** The pointer to write for this answer, or null when the stored one stands. */
export function pointerToHeal(pointer: CurrentCharacter | null, me: Me): CurrentCharacter | null {
  const main = mainCharacter(me.characters, me.main_character_key);
  if (main === null) return null;
  if (pointer === null) return pointerForCharacter(main);
  if ((pointer.source === 'code' || pointer.source === 'addon') && pointerNameOf(pointer) === '') {
    return pointerForCharacter(main);
  }
  if (pointer.source === 'armory' && !me.characters.some((character) => character.key === pointer.ref)) {
    return pointerForCharacter(main);
  }
  return null;
}

function healedThisSession(): boolean {
  try {
    return globalThis.sessionStorage?.getItem(HEALED_KEY) === '1';
  } catch {
    return false;
  }
}

function markHealed(on: boolean): void {
  try {
    if (on) globalThis.sessionStorage?.setItem(HEALED_KEY, '1');
    else globalThis.sessionStorage?.removeItem(HEALED_KEY);
  } catch {
    // A private window: the rule simply runs again on the next page.
  }
}

/** Runs the rule once per browser session, on the first signed-in `/v1/me` answer. */
export function settlePointerForSession(me: Me | null, storage?: Storage): void {
  if (me === null || healedThisSession()) return;
  markHealed(true);
  const next = pointerToHeal(readCurrent(storage), me);
  if (next === null) return;
  writeCurrent(next, storage);
  if (typeof window !== 'undefined') window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
}

/** Sign-out: the pointer belongs to the account that is leaving. */
export function clearPointerOnSignOut(storage?: Storage): void {
  clearCurrent(storage);
  markHealed(false);
}
