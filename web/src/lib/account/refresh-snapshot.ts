// web/src/lib/account/refresh-snapshot.ts
// "Refresh from Battle.net" leaves the page for Battle.net's sign-in and comes back to
// /account?refreshed=1 (design/specs/2026-11-04-account-characters.md section 6). What the
// account held before the redirect is kept in sessionStorage, so the page can say on return
// what the refresh did: updated, no changes, or that Battle.net did not answer (the import
// time did not move).
import type { Me } from './me-types';

const KEY = 'fs.refresh.snapshot';

export interface RefreshSnapshot {
  /** `me.bnet_imported_at` before the refresh; null when the account had never imported. */
  at: string | null;
  /** Each character's level by key (0 when unknown). */
  chars: Record<string, number>;
}

export type RefreshResult =
  | { kind: 'changed'; count: number; added: number }
  | { kind: 'unchanged'; count: number }
  | { kind: 'error' };

export function snapshotOf(me: Pick<Me, 'characters' | 'bnet_imported_at'>): RefreshSnapshot {
  return {
    at: me.bnet_imported_at ?? null,
    chars: Object.fromEntries(me.characters.map((c) => [c.key, c.level ?? 0])),
  };
}

/** What the refresh did, from the snapshot taken before it and the account after it. With no
 *  snapshot (a new tab, storage blocked) a return from the refresh reads as updated. */
export function compareSnapshot(
  before: RefreshSnapshot | null,
  after: Pick<Me, 'characters' | 'bnet_imported_at'>,
): RefreshResult {
  const count = after.characters.length;
  if (before === null) return { kind: 'changed', count, added: 0 };
  const advanced =
    after.bnet_imported_at !== undefined &&
    (before.at === null || Date.parse(after.bnet_imported_at) > Date.parse(before.at));
  if (!advanced) return { kind: 'error' };
  const now = snapshotOf(after).chars;
  const added = Object.keys(now).filter((key) => !(key in before.chars)).length;
  const removed = Object.keys(before.chars).some((key) => !(key in now));
  const levelled = Object.keys(now).some((key) => key in before.chars && before.chars[key] !== now[key]);
  return added > 0 || removed || levelled ? { kind: 'changed', count, added } : { kind: 'unchanged', count };
}

/** The storage the snapshot lives in; any of its calls may throw (private mode, blocked). */
export interface SnapshotStore {
  getItem(key: string): string | null;
  setItem(key: string, value: string): void;
  removeItem(key: string): void;
}

export function saveSnapshot(snapshot: RefreshSnapshot, store: SnapshotStore): void {
  try {
    store.setItem(KEY, JSON.stringify(snapshot));
  } catch {
    // Storage blocked: the return reads as updated (compareSnapshot with no snapshot).
  }
}

/** The saved snapshot, removed as it is read; null when there is none or it does not parse. */
export function takeSnapshot(store: SnapshotStore): RefreshSnapshot | null {
  try {
    const raw = store.getItem(KEY);
    store.removeItem(KEY);
    if (raw === null) return null;
    const parsed = JSON.parse(raw) as Partial<RefreshSnapshot>;
    if (typeof parsed.chars !== 'object' || parsed.chars === null) return null;
    return { at: typeof parsed.at === 'string' ? parsed.at : null, chars: parsed.chars };
  } catch {
    return null;
  }
}
