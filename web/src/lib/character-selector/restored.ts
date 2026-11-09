// web/src/lib/character-selector/restored.ts
// The header selector's "Restored your last character" line and its Forget (2026-10-31, which
// retired the simulator's spine bar that used to carry both). When a tab's session begins,
// Base.astro's pre-paint script records which character the stored pointer named at that
// moment (sessionStorage[SESSION_START_POINTER_KEY], before any island can write a pointer):
// an empty string when none was stored. The current character is restored while the pointer
// still names that same character. Compared by identity rather than by the pointer's savedAt,
// because the simulator rewrites the pointer each time it loads the restored character.
import { CURRENT_CHARACTER_CHANGED, clearCurrent, type CurrentCharacter } from '../current-character';

/** The sessionStorage key of the pointer identity this tab session started with. */
export const SESSION_START_POINTER_KEY = 'fs.sessionStartPointer';

/** sessionStorage: '1' when this tab's page load began a fresh browser session (consumed by
 *  account/session-pointer.ts). */
export const SESSION_FRESH_KEY = 'fs.sessionFresh';

/** localStorage: the last page load of this browser, in epoch ms. */
export const LAST_ACTIVE_KEY = 'fs.lastActiveAt';

/** A browser session ends after this long with no page of the site opened. */
export const SESSION_IDLE_MS = 30 * 60_000;

/**
 * Whether a page load begins a fresh browser session: nothing recorded, or the last page load
 * was SESSION_IDLE_MS or more ago. Base.astro's pre-paint script runs the same rule (an inline
 * script cannot import). A tab opened mid-session is not fresh, so a choice made in another tab
 * of the same visit stands.
 */
export function isFreshBrowserSession(lastActiveMs: number | null, nowMs: number): boolean {
  return lastActiveMs === null || !(lastActiveMs > 0) || nowMs - lastActiveMs >= SESSION_IDLE_MS;
}

/** `source:ref`, the identity Base.astro's pre-paint script writes in the same shape. */
export function pointerIdentity(pointer: Pick<CurrentCharacter, 'source' | 'ref'>): string {
  return `${pointer.source}:${pointer.ref}`;
}

function sessionStorageOf(storage: Storage | undefined): Storage | null {
  if (storage !== undefined) return storage;
  try {
    return globalThis.sessionStorage ?? null;
  } catch {
    return null;
  }
}

/** The identity the session started with: '' for no pointer, null when nothing was recorded. */
export function readSessionStartPointer(storage?: Storage): string | null {
  const target = sessionStorageOf(storage);
  if (target === null) return null;
  try {
    return target.getItem(SESSION_START_POINTER_KEY);
  } catch {
    return null;
  }
}

/**
 * Whether the selector offers "Restored your last character" and Forget: the pointer still names
 * the character this tab session started with. Never for the account's main (`mainKey`):
 * forgetting it would leave the main in its place, so the line would offer an action that
 * changes nothing.
 */
export function isRestoredPointer(
  pointer: CurrentCharacter | null,
  sessionStartIdentity: string | null,
  mainKey: string | null,
): boolean {
  if (pointer === null || sessionStartIdentity === null || sessionStartIdentity === '') return false;
  if (pointer.source === 'armory' && pointer.ref === mainKey) return false;
  return pointerIdentity(pointer) === sessionStartIdentity;
}

export interface ForgetDeps {
  storage?: Storage;
  dispatch: (event: Event) => void;
}

/** Forget: clears the pointer (the account's main, if any, takes its place) and tells every island. */
export function forgetPointer(deps: ForgetDeps): void {
  clearCurrent(deps.storage);
  deps.dispatch(new Event(CURRENT_CHARACTER_CHANGED));
}
