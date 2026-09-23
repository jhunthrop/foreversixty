// web/src/lib/account/session-cache.ts
// The last /v1/me answer, kept in localStorage so the next page load renders who is signed
// in and their characters at hydration instead of after a network round trip: the island
// shows the snapshot, then revalidates in the background and announces a change.
//
// Trust rules: the snapshot is only ever read while the readable half of the session's
// cookie pair (fs_csrf) is present, so a signed-out browser never shows a stale account;
// it expires after SNAPSHOT_TTL_MS; it is cleared on sign-out and whenever the API answers
// that there is no session. It holds what /v1/me holds (tag, characters, guilds,
// entitlements): personal, not secret, and never a token.
import type { Me } from './api';

export const SNAPSHOT_KEY = 'fs.me';
export const SNAPSHOT_TTL_MS = 10 * 60 * 1000;
/** Dispatched on `window` when a background revalidation changed the snapshot. */
export const ME_UPDATED = 'fs:me-updated';

interface Snapshot {
  me: Me;
  savedAt: number;
}

function storageOf(storage?: Storage): Storage | null {
  if (storage !== undefined) return storage;
  try {
    return globalThis.localStorage ?? null;
  } catch {
    return null;
  }
}

/** True when the session's readable cookie is present: the only time a snapshot is trusted. */
export function sessionHinted(
  cookie: string = typeof document === 'undefined' ? '' : document.cookie,
): boolean {
  return /(?:^|;\s*)fs_csrf=[^;]+/.test(cookie);
}

export function readSnapshot(now: number = Date.now(), storage?: Storage, cookie?: string): Me | null {
  if (!sessionHinted(cookie)) return null;
  const target = storageOf(storage);
  if (target === null) return null;
  try {
    const raw = target.getItem(SNAPSHOT_KEY);
    if (raw === null) return null;
    const parsed = JSON.parse(raw) as Partial<Snapshot>;
    if (typeof parsed.savedAt !== 'number' || parsed.me === undefined) return null;
    if (now - parsed.savedAt > SNAPSHOT_TTL_MS) return null;
    return parsed.me;
  } catch {
    return null;
  }
}

export function writeSnapshot(me: Me, now: number = Date.now(), storage?: Storage): void {
  const target = storageOf(storage);
  if (target === null) return;
  try {
    target.setItem(SNAPSHOT_KEY, JSON.stringify({ me, savedAt: now } satisfies Snapshot));
  } catch {
    // Quota, private mode, or a disabled store: the snapshot is a convenience only.
  }
}

export function clearSnapshot(storage?: Storage): void {
  const target = storageOf(storage);
  if (target === null) return;
  try {
    target.removeItem(SNAPSHOT_KEY);
  } catch {
    // Same as above.
  }
}

/** Whether two answers differ, for deciding whether to announce a revalidation. */
export function sameMe(a: Me | null, b: Me | null): boolean {
  return JSON.stringify(a) === JSON.stringify(b);
}
