// web/src/lib/guild/dense/roster-board.ts
// The dense Roster (design/specs/2026-11-09-guild-dense.md section 5.3): one row per raider joining
// what the guild reads already return: the home's roster row (rating, item level, best parse,
// attendance, gap, last log), the attendance grid's last nights, the loot council's awards this
// tier and the readiness board's checks. Grouped by the raid group the API sends, each group head
// with the comp target's seats. Pure: the component draws what this returns.
import type { GuildRaidGroup, GuildReadinessRow, GuildRosterRow } from '../api';
import { missingCount, readinessShortfalls } from '../readiness-view';
import type { LootDecision, LootNights } from '../loot-council-api';
import type { AttendanceCell, GuildCompTarget, GuildRecentAttendance } from '../overview-api';

export const RAID_GROUP_ORDER: readonly GuildRaidGroup[] = ['tank', 'healer', 'melee', 'ranged'];

export const GROUP_WORD: Readonly<Record<GuildRaidGroup, string>> = {
  tank: 'Tanks',
  healer: 'Healers',
  melee: 'Melee',
  ranged: 'Ranged',
};

const TARGET_KEY: Readonly<Record<GuildRaidGroup, keyof Omit<GuildCompTarget, 'source'>>> = {
  tank: 'tanks',
  healer: 'healers',
  melee: 'melee',
  ranged: 'ranged',
};

const DAY_MS = 86_400_000;
/** From this many days without an export the Seen cell shows; from STALE_DAYS it is red. */
export const SEEN_FROM_DAYS = 2;
export const STALE_DAYS = 7;

export interface BoardRow {
  key: string;
  name: string;
  cls: string;
  spec: string;
  group: GuildRaidGroup;
  rank: GuildRosterRow['rank'];
  rating: number | null;
  itemLevel: number | null;
  /** Item level minus the guild's median; null without either. */
  itemLevelDelta: number | null;
  /** The BiS-file gain of the raider's upgrades (readiness, else the addon's gap); null for a tank. */
  gap: number | null;
  gapUnit: 'dps' | 'hps' | 'tps' | null;
  bestPercentile: number | null;
  /** The last nights, oldest first; empty without the grid. */
  cells: readonly AttendanceCell[];
  /** Attendance over the last 30 nights as a whole percent; null without it. */
  attendancePct: number | null;
  items: number;
  /** Readiness checks failed for the next raid; null when the board has no row for the raider. */
  short: number | null;
  /** Whole days since the last export; null without one. */
  seenDays: number | null;
  /** The picker's Discord name for a character picked on Discord, not yet confirmed (officers only). */
  pickedBy: string | null;
}

export interface BoardGroup {
  group: GuildRaidGroup;
  need: number | null;
  medianItemLevel: number | null;
  rows: BoardRow[];
}

export interface BoardInputs {
  roster: readonly GuildRosterRow[];
  attendance: GuildRecentAttendance | null;
  loot: LootNights | null;
  readiness: readonly GuildReadinessRow[] | null;
  medianItemLevel: number | null;
  target: GuildCompTarget | null;
  now: Date;
}

/** Every award this tier, once per award (a night's decisions; a drop's award repeats one). */
export function tierAwards(loot: LootNights | null): LootDecision[] {
  if (loot === null) return [];
  const seen = new Map<number, LootDecision>();
  for (const night of loot.nights) for (const d of night.decisions) seen.set(d.award_id, d);
  return [...seen.values()].sort((a, b) => a.at.localeCompare(b.at));
}

export function median(values: readonly number[]): number | null {
  if (values.length === 0) return null;
  const sorted = [...values].sort((a, b) => a - b);
  const mid = Math.floor(sorted.length / 2);
  return sorted.length % 2 === 1 ? sorted[mid] : (sorted[mid - 1] + sorted[mid]) / 2;
}

function groupOf(row: GuildRosterRow): GuildRaidGroup {
  if (row.group !== undefined) return row.group;
  return row.role === 'tank' || row.role === 'healer' ? row.role : 'ranged';
}

function percentOf(att: GuildRosterRow['attendance_30'] | GuildRosterRow['attendance']): number | null {
  if (att === undefined || att === null || att.nights === 0) return null;
  return Math.round((100 * att.present) / att.nights);
}

function daysSince(iso: string | undefined, now: Date): number | null {
  if (iso === undefined || iso === '') return null;
  return Math.max(0, Math.floor((now.getTime() - new Date(iso).getTime()) / DAY_MS));
}

function boardRow(
  row: GuildRosterRow,
  inputs: BoardInputs,
  cellsByAccount: ReadonlyMap<string, readonly AttendanceCell[]>,
  itemsByKey: ReadonlyMap<string, number>,
  readyByKey: ReadonlyMap<string, GuildReadinessRow>,
): BoardRow {
  const ready = readyByKey.get(row.character_key);
  const gap = ready?.gear_gap?.gain_dps ?? row.addon?.gear_gap ?? null;
  const unit = ready?.gear_gap?.unit ?? null;
  const ilvl = row.item_level ?? null;
  return {
    key: row.character_key,
    name: row.name,
    cls: row.class ?? '',
    spec: row.spec_name ?? row.spec ?? '',
    group: groupOf(row),
    rank: row.rank,
    rating: row.rating?.overall ?? null,
    itemLevel: ilvl,
    itemLevelDelta: ilvl === null || inputs.medianItemLevel === null ? null : ilvl - inputs.medianItemLevel,
    gap: groupOf(row) === 'tank' ? null : gap,
    gapUnit: unit,
    bestPercentile: row.best_parse?.percentile ?? null,
    cells: cellsByAccount.get(row.account_key ?? '') ?? [],
    attendancePct: percentOf(row.attendance_30 ?? row.attendance),
    items: itemsByKey.get(row.character_key) ?? 0,
    short:
      ready === undefined
        ? null
        : missingCount(ready, readinessShortfalls(ready, { attunement: true }, false)),
    seenDays: daysSince(row.logged_at, inputs.now),
    pickedBy: row.discord_picked_by ?? null,
  };
}

/** The roster's raiders, grouped by raid group, each group best rating first. */
export function rosterBoard(inputs: BoardInputs): BoardGroup[] {
  const cellsByAccount = new Map(
    (inputs.attendance?.rows ?? []).map((r) => [r.account_key, r.cells] as const),
  );
  const itemsByKey = new Map<string, number>();
  for (const award of tierAwards(inputs.loot))
    itemsByKey.set(award.character_key, (itemsByKey.get(award.character_key) ?? 0) + 1);
  const readyByKey = new Map((inputs.readiness ?? []).map((r) => [r.character_key, r] as const));
  const rows = inputs.roster.map((row) => boardRow(row, inputs, cellsByAccount, itemsByKey, readyByKey));
  return RAID_GROUP_ORDER.map((group) => {
    const inGroup = rows.filter((r) => r.group === group);
    return {
      group,
      need: inputs.target === null ? null : inputs.target[TARGET_KEY[group]],
      medianItemLevel: median(inGroup.flatMap((r) => (r.itemLevel === null ? [] : [r.itemLevel]))),
      rows: sortBoard(inGroup, 'rating'),
    };
  }).filter((g) => g.rows.length > 0);
}

export type BoardSortKey = 'rating' | 'ilvl' | 'gap' | 'best' | 'nights' | 'loot' | 'ready' | 'seen';

const SORT_VALUE: Readonly<Record<BoardSortKey, (r: BoardRow) => number | null>> = {
  rating: (r) => r.rating,
  ilvl: (r) => r.itemLevel,
  gap: (r) => r.gap,
  best: (r) => r.bestPercentile,
  nights: (r) => r.attendancePct,
  loot: (r) => r.items,
  ready: (r) => r.short,
  seen: (r) => r.seenDays,
};

/** Highest first, a missing figure last, ties by name. */
export function sortBoard(rows: readonly BoardRow[], key: BoardSortKey): BoardRow[] {
  const value = SORT_VALUE[key];
  return [...rows].sort((a, b) => {
    const va = value(a);
    const vb = value(b);
    if (va === vb) return a.name.localeCompare(b.name);
    if (va === null) return 1;
    if (vb === null) return -1;
    return vb - va;
  });
}

/** The roster's raiders of one group, or all, whose name starts a word with `text`. */
export function filterBoard(
  groups: readonly BoardGroup[],
  only: GuildRaidGroup | 'all',
  text: string,
): BoardGroup[] {
  const needle = text.trim().toLowerCase();
  return groups
    .filter((g) => only === 'all' || g.group === only)
    .map((g) => ({
      ...g,
      rows: g.rows.filter((r) => needle === '' || r.name.toLowerCase().includes(needle)),
    }))
    .filter((g) => g.rows.length > 0);
}

/** The Seen cell: blank under SEEN_FROM_DAYS, else `3 d`. */
export function seenText(days: number | null): string {
  if (days === null) return '–';
  return days < 1 ? 'today' : `${days} d`;
}

/** `Tanks 3 · need 3 · median item level 63`'s parts after the title. */
export function groupLine(group: BoardGroup): string {
  return [
    group.need === null ? '' : `need ${group.need}`,
    group.medianItemLevel === null ? '' : `median item level ${Math.floor(group.medianItemLevel)}`,
  ]
    .filter((p) => p !== '')
    .join(' · ');
}
