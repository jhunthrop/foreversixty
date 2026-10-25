// web/src/lib/sim/server-runs-poll.ts
// The rules under the "Your server runs" list that are not rendering: which rows are still in
// flight, which flipped to done while the member watched, and the two localStorage hints
// (experience spec 2026-10-24 sections 3.5 and 3.7) the pre-paint script in Base.astro reads
// back. Pure functions plus guarded storage access; the reactive wrapper is
// server-runs-controller.svelte.ts.
import { MAX_VISIBLE_RUNS } from './server-runs';
import type { SimListRow } from './types';

/** How often the list re-reads while a run is queued or running and the tab is visible. */
export const POLL_INTERVAL_MS = 10_000;
/** After this page starts a server run, how long before the list first looks for the new row. */
export const RUN_START_REFRESH_MS = 2_000;
/** Last list read's row count (an integer string); sizes the reserved box on the next visit. */
export const COUNT_STORAGE_KEY = 'fs.serverRuns.count';
/** When a run was last started or seen in flight (epoch ms); the band's reserve hint. */
export const LIVE_STORAGE_KEY = 'fs.serverRuns.live';
export const LIVE_TTL_MS = 2 * 60 * 60 * 1000;

/** What the `Run again` button on a failed row is showing. */
export type RerunStatus = 'idle' | 'pending' | 'started' | 'failed' | 'lapsed';

export function isInFlight(row: SimListRow): boolean {
  return row.state === 'queued' || row.state === 'running';
}

export function hasInFlight(rows: readonly SimListRow[] | null): boolean {
  return rows !== null && rows.some(isInFlight);
}

/** Rows that were queued or running in `before` and are done and unopened in `after`. */
export function newlyFinished(
  before: readonly SimListRow[] | null,
  after: readonly SimListRow[],
): SimListRow[] {
  if (before === null) return [];
  const inFlightIds = new Set(before.filter(isInFlight).map((row) => row.sim_id));
  return after.filter((row) => row.state === 'done' && row.seen === false && inFlightIds.has(row.sim_id));
}

/** `rows` with the given ids marked seen; the input is not changed. */
export function withSeen(rows: readonly SimListRow[], ids: readonly string[]): SimListRow[] {
  return rows.map((row) => (ids.includes(row.sim_id) ? { ...row, seen: true } : row));
}

/** Rows the reserved box holds: the last count, at least one, at most the visible five. */
export function reservedRowCount(stored: number): number {
  return Math.max(1, Math.min(Number.isFinite(stored) ? Math.floor(stored) : 1, MAX_VISIBLE_RUNS));
}

function storage(): Storage | null {
  try {
    return globalThis.localStorage ?? null;
  } catch {
    return null;
  }
}

export function writeRowCount(count: number): void {
  try {
    storage()?.setItem(COUNT_STORAGE_KEY, String(count));
  } catch {
    // Private mode or quota: the hint is a convenience.
  }
}

export function readRowCount(): number {
  try {
    return reservedRowCount(Number(storage()?.getItem(COUNT_STORAGE_KEY) ?? 1));
  } catch {
    return 1;
  }
}

export function markLive(now: number = Date.now()): void {
  try {
    storage()?.setItem(LIVE_STORAGE_KEY, String(now));
  } catch {
    // Same as writeRowCount.
  }
}

export function isLive(now: number = Date.now()): boolean {
  try {
    const stamp = Number(storage()?.getItem(LIVE_STORAGE_KEY));
    return Number.isFinite(stamp) && stamp > 0 && now - stamp < LIVE_TTL_MS;
  } catch {
    return false;
  }
}
