// web/src/lib/account/session-pointer.ts
// The pointer at the start of a signed-in session. When `/v1/me` first answers for a browser
// session and the stored pointer is absent, a nameless pasted export, or an account character this
// account does not hold (a previous player's, left in the browser), the account's main becomes the
// current character -- whichever way the player signed in (the Battle.net return, the email link
// and the dev sign-in all end the same). After that the pointer changes only when the player
// chooses; signing out clears it.
//
// In a fresh browser session (restored.ts: the first page after a long idle gap), a pointer
// restored from the earlier visit gives way to the character played most recently
// (last-played.ts): the player logged another character in since. A pointer changed in this
// session is the player's own choice and always stands.
import {
  announceSessionHeal,
  clearCurrent,
  readCurrent,
  writeCurrent,
  type CurrentCharacter,
} from '../current-character';
import { pointerIdentity, readSessionStartPointer, SESSION_FRESH_KEY } from '../character-selector/restored';
import type { Me } from './api';
import { heroCharacter, pointerNameOf } from './hero-character';
import { lastPlayedCharacter } from './last-played';
import { mainCharacter, pointerForCharacter } from './main-character';

/** sessionStorage: '1' once the rule has run in this tab's session (Base.astro clears it when a fresh one begins). */
export const POINTER_HEALED_KEY = 'fs.pointerHealed';

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

/**
 * The pointer a fresh session writes over a restored one, or null when it stands: the character
 * played last when it is another one (a tie keeps the restored character), else the heal rule.
 */
export function pointerForFreshSession(restored: CurrentCharacter, me: Me): CurrentCharacter | null {
  const named = heroCharacter(restored, me.characters);
  const latest = lastPlayedCharacter(me.characters, [named?.key, me.main_character_key]);
  if (latest === null || latest.key === named?.key) return pointerToHeal(restored, me);
  return pointerForCharacter(latest);
}

/** Reads and clears the fresh-session mark Base.astro's pre-paint script set for this tab. */
function takeFreshSession(): boolean {
  try {
    const fresh = globalThis.sessionStorage?.getItem(SESSION_FRESH_KEY) === '1';
    globalThis.sessionStorage?.removeItem(SESSION_FRESH_KEY);
    return fresh;
  } catch {
    return false;
  }
}

function healedThisSession(): boolean {
  try {
    return globalThis.sessionStorage?.getItem(POINTER_HEALED_KEY) === '1';
  } catch {
    return false;
  }
}

function markHealed(on: boolean): void {
  try {
    if (on) globalThis.sessionStorage?.setItem(POINTER_HEALED_KEY, '1');
    else globalThis.sessionStorage?.removeItem(POINTER_HEALED_KEY);
  } catch {
    // A private window: the rule simply runs again on the next page.
  }
}

/** Runs the rule once per browser session, on the first signed-in `/v1/me` answer. */
export function settlePointerForSession(me: Me | null, storage?: Storage): void {
  if (me === null || healedThisSession()) return;
  markHealed(true);
  const pointer = readCurrent(storage);
  const restored = takeFreshSession() && pointer !== null && pointerIdentity(pointer) === readSessionStartPointer();
  const next = restored ? pointerForFreshSession(pointer, me) : pointerToHeal(pointer, me);
  if (next === null) return;
  writeCurrent(next, storage);
  // Tagged: the pages move to the main as for a choice, but not over the visitor's own work.
  announceSessionHeal();
}

/** Sign-out: the pointer belongs to the account that is leaving. */
export function clearPointerOnSignOut(storage?: Storage): void {
  clearCurrent(storage);
  markHealed(false);
}
