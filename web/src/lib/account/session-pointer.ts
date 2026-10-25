// web/src/lib/account/session-pointer.ts
// The pointer at the start of a signed-in session. When `/v1/me` first answers for a browser
// session and the stored pointer is absent, a nameless pasted export, or an account character this
// account does not hold (a previous player's, left in the browser), the account's main becomes the
// current character -- whichever way the player signed in (the Battle.net return, the email link
// and the dev sign-in all end the same). After that the pointer changes only when the player
// chooses; signing out clears it.
import {
  announcePageLoadedCharacter,
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
  if (pointer.source === 'armory') {
    const held = me.characters.find((character) => character.key === pointer.ref);
    if (held === undefined) return pointerForCharacter(main);
    // The label carries the spec the pointer was written with; the rule may read the record
    // differently now (new talents, a better record), so the label follows it.
    const label = pointerForCharacter(held).label;
    return label === pointer.label ? null : { ...pointer, label };
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
  // The heal is the site settling its own default, not a player choosing: the displays re-read the
  // pointer, the followers that load a character in place leave the page alone.
  announcePageLoadedCharacter();
}

/** Sign-out: the pointer belongs to the account that is leaving. */
export function clearPointerOnSignOut(storage?: Storage): void {
  clearCurrent(storage);
  markHealed(false);
}
