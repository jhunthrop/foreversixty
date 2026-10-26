// web/src/lib/guild/readiness-raid.ts
// The Readiness tab's next-raid filter (attendance spec 3.7): the board read twice (the guild's, then
// the next raid's In and Late sign-ups), which one the toggle shows, and what the Last synced cell says.
import { fetchGuildReadiness, type GuildReadinessPage, type GuildReadinessRow } from './api';
import { attendanceCopy, syncedBefore } from './attendance-copy';
import { daysBeforeRaid } from './night';
import type { ReadinessRaid } from './attendance-contract';

export interface ReadinessBoard {
  /** The guild's whole board; its `next_raid` says whether a filter exists. */
  all: GuildReadinessPage;
  /** The next raid's board, or null with no raid planned or when that read failed. */
  raid: GuildReadinessPage | null;
}

type ReadinessFetch = (guildId: number, raidId?: number) => Promise<GuildReadinessPage>;

/** Reads the board, then (when a raid is planned or locked) that raid's own board. A filtered read
 *  that fails leaves the whole board and no toggle: the tab never hides the guild's rows. */
export async function loadReadinessBoard(
  guildId: number,
  fetchPage: ReadinessFetch = fetchGuildReadiness,
): Promise<ReadinessBoard> {
  const all = await fetchPage(guildId);
  const nextRaid = all.next_raid ?? null;
  if (nextRaid === null) return { all, raid: null };
  try {
    return { all, raid: await fetchPage(guildId, nextRaid.id) };
  } catch {
    return { all, raid: null };
  }
}

export const toggleVisible = (board: ReadinessBoard): boolean => board.raid?.raid !== undefined;

/** The rows to draw: the raid's sign-ups while the toggle is on, the guild's board otherwise. */
export function readinessRows(board: ReadinessBoard, nextRaidOnly: boolean): GuildReadinessRow[] {
  return nextRaidOnly && board.raid !== null ? board.raid.rows : board.all.rows;
}

export type SyncCell =
  { kind: 'no_export'; text: string } | { kind: 'stale'; text: string } | { kind: 'day' };

/** The Last synced cell: no export, an export synced before the raid's cutoff, or the day word. */
export function syncCell(row: GuildReadinessRow, raid: ReadinessRaid | null): SyncCell {
  if (row.no_export) return { kind: 'no_export', text: attendanceCopy.readiness.noExport };
  if (raid !== null && row.synced_before_cutoff === true) {
    return { kind: 'stale', text: syncedBefore(daysBeforeRaid(row.logged_at, raid.starts_at)) };
  }
  return { kind: 'day' };
}
