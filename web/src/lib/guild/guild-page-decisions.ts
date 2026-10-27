// web/src/lib/guild/guild-page-decisions.ts
// The decisions Guild.svelte takes around its tab reads, kept pure so they are table-tested: what the
// Readiness tab shows after a (re-)read, and what a band press does to the Raids tab.
import type { GuildReadinessRow } from './api';
import type { AttendanceSource, ReadinessRaid } from './attendance-contract';
import { readinessRows, type ReadinessBoard } from './readiness-raid';

/** The attendance source a surface renders until the roster says otherwise. */
export const DEFAULT_ATTENDANCE_SOURCE: AttendanceSource = 'reports';

export const attendanceSourceOf = (home: { attendance_source?: AttendanceSource } | null): AttendanceSource =>
  home?.attendance_source ?? DEFAULT_ATTENDANCE_SOURCE;

export interface HomeNextRaid {
  id: number;
  guild_timezone: string | null;
}

export interface ReadinessTabView {
  rows: GuildReadinessRow[];
  /** The filtered board's raid; null shows no toggle (no raid, or its read failed). */
  raid: ReadinessRaid | null;
  /** The raid page the `Showing` line links; null with no raid. */
  raidId: number | null;
  /** Taken from home only when home's next raid is the readiness raid; otherwise unknown (viewer zone). */
  guildTimezone: string | null;
}

/** What the Readiness tab draws. The toggle is the page's own state, so it survives a re-read; a board
 *  without a filtered read (the filter dropped or no raid planned) shows the guild's rows whatever it says. */
export function readinessTabView(
  board: ReadinessBoard | null,
  nextRaidOnly: boolean,
  homeNextRaid: HomeNextRaid | null,
): ReadinessTabView {
  if (board === null) return { rows: [], raid: null, raidId: null, guildTimezone: null };
  const raid = board.raid?.raid ?? null;
  return {
    rows: readinessRows(board, nextRaidOnly),
    raid,
    raidId: raid?.id ?? null,
    guildTimezone: raid !== null && homeNextRaid?.id === raid.id ? homeNextRaid.guild_timezone : null,
  };
}

export type TabFetchStatus = 'idle' | 'loading' | 'ready' | 'missing';

export type RaidsChangeAction = 'reread' | 'reset';

/** A band press moved the Raids tab's data: an open tab (or one still loading) reads again in place;
 *  anything else is marked unread so it reads when next shown. */
export function raidsChangeAction(
  hasGuild: boolean,
  raidsStatus: TabFetchStatus,
  tabOpen: boolean,
): RaidsChangeAction {
  return hasGuild && (raidsStatus === 'loading' || (tabOpen && raidsStatus === 'ready')) ? 'reread' : 'reset';
}
