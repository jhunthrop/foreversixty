// web/src/lib/guild/roster-sort.ts
// Guild control-centre spec §4.B: the Roster tab's filter bar, five-column sort header, and
// "own row pinned first" rule. Pure over `GuildRosterRow` (api.ts) so GuildRosterTable.svelte
// never duplicates a comparator inline.
import type { GuildRosterRow } from './api';

export type RosterSortKey = 'rank' | 'ilvl' | 'rating' | 'attendance' | 'seen';
export type RosterRoleFilter = 'all' | 'tank' | 'healer' | 'dps';

export interface RosterFilters {
  role: RosterRoleFilter;
  classFilter: string;
  verifiedOnly: boolean;
  belowFloorOnly: boolean;
}

export const DEFAULT_ROSTER_FILTERS: RosterFilters = {
  role: 'all',
  classFilter: 'all',
  verifiedOnly: false,
  belowFloorOnly: false,
};

/** Spec §4.A's "below the rating floor" count: percentile 25 within the roster's own
 *  rated population. `null` when nobody on the roster has a rating yet (never a fabricated
 *  floor of 0). */
export function ratingFloor(rows: readonly GuildRosterRow[]): number | null {
  const values = rows
    .map((row) => row.rating?.overall)
    .filter((value): value is number => value !== undefined && value !== null)
    .sort((a, b) => a - b);
  if (values.length === 0) return null;
  const index = Math.min(Math.floor(values.length * 0.25), values.length - 1);
  return values[index];
}

/** A row with no rating is never counted as "below" -- that is a missing fact, not a
 *  failing one (tenet 1). */
export function isBelowRatingFloor(row: GuildRosterRow, floor: number | null): boolean {
  if (floor === null) return false;
  const overall = row.rating?.overall;
  return overall !== undefined && overall !== null && overall < floor;
}

/** Filters combine as AND (spec §4.B). */
export function applyRosterFilters(
  rows: readonly GuildRosterRow[],
  filters: RosterFilters,
  floor: number | null,
): GuildRosterRow[] {
  return rows.filter((row) => {
    if (filters.role !== 'all' && row.role !== filters.role) return false;
    if (filters.classFilter !== 'all' && row.class !== filters.classFilter) return false;
    if (filters.verifiedOnly && !row.verified) return false;
    if (filters.belowFloorOnly && !isBelowRatingFloor(row, floor)) return false;
    return true;
  });
}

function rankWeight(rank: GuildRosterRow['rank']): number {
  return rank === 'leader' ? 2 : rank === 'officer' ? 1 : 0;
}

function attendanceRatio(row: GuildRosterRow): number {
  const attendance = row.attendance;
  if (attendance === undefined || attendance.nights === 0) return -1;
  return attendance.present / attendance.nights;
}

function loggedAtMs(row: GuildRosterRow): number {
  if (row.logged_at === undefined) return row.logged_recently ? 1 : 0;
  const parsed = Date.parse(row.logged_at);
  return Number.isNaN(parsed) ? 0 : parsed;
}

/** One comparator per sort-header column (spec §4.B): Rank, Item level, Rating,
 *  Attendance, Last seen -- descending in every case (rank-highest, highest ilvl, etc.
 *  first), never mutates `rows`. */
export function sortRoster(rows: readonly GuildRosterRow[], key: RosterSortKey): GuildRosterRow[] {
  const sorted = [...rows];
  sorted.sort((a, b) => {
    switch (key) {
      case 'rank':
        return rankWeight(b.rank) - rankWeight(a.rank);
      case 'ilvl':
        return (b.item_level ?? -1) - (a.item_level ?? -1);
      case 'rating':
        return (b.rating?.overall ?? -1) - (a.rating?.overall ?? -1);
      case 'attendance':
        return attendanceRatio(b) - attendanceRatio(a);
      case 'seen':
        return loggedAtMs(b) - loggedAtMs(a);
      default:
        return 0;
    }
  });
  return sorted;
}

/** Spec §4.B: "Member's own row is pinned first" -- applied after sorting, so the rest of
 *  the table stays in the chosen sort order. */
export function pinOwnRowFirst(
  rows: readonly GuildRosterRow[],
  myCharacterKey: string | null,
): GuildRosterRow[] {
  if (myCharacterKey === null) return [...rows];
  return [...rows].sort(
    (a, b) => Number(b.character_key === myCharacterKey) - Number(a.character_key === myCharacterKey),
  );
}

export interface OrderRosterOptions {
  filters: RosterFilters;
  floor: number | null;
  sortKey: RosterSortKey;
  myCharacterKey: string | null;
}

/**
 * The Roster tab's full row order (spec §4.B): unverified rows always lead (unchanged v1
 * rule, `orderRoster`'s own reasoning, never filtered out -- an officer must always see
 * who is waiting), then the verified rows filtered, sorted and pinned.
 */
export function orderRosterForTab(
  rows: readonly GuildRosterRow[],
  options: OrderRosterOptions,
): GuildRosterRow[] {
  const unverified = rows.filter((row) => !row.verified);
  const verified = applyRosterFilters(
    rows.filter((row) => row.verified),
    options.filters,
    options.floor,
  );
  const ordered = pinOwnRowFirst(sortRoster(verified, options.sortKey), options.myCharacterKey);
  return [...unverified, ...ordered];
}
