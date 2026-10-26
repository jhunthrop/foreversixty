// web/src/lib/guild/night.ts
// The pure half of the raid-night surfaces (design/specs/2026-10-27-guild-attendance.md 3.3, 3.5,
// 4): the five attendance groups in order, the `Your night` sentence and facts, the copied
// Discord text, and the readiness freshness count. Strings come from attendance-copy.ts.
import {
  BIGGEST_CHANGES_ROWS,
  DAY_MS,
  DISCORD_REPORT_MAX_LINES,
  REPORT_ROWS_BEFORE_SHOW_ALL,
  SECONDS_PER_MINUTE,
  attendanceCopy,
  cameAsOutcome,
  cameRatio,
  deaths,
  lateOutcome,
  durationLabel,
  kills,
  killAfter,
  killTimeLabel,
  metricLabel,
  moveDelta,
  wipes,
} from './attendance-copy';
import type {
  RaidAttendanceRow,
  RaidKill,
  RaidMoveRow,
  RaidMoves,
  RaidNight,
  RaidReport,
} from './attendance-contract';
import type { GuildRaidRole } from './api';
import { discordText, yourNightText } from './night-copy';
import { SEP, raidCopy } from './raid-copy';
import { rosterOrder } from './raid-roster';
import { roleForSpec } from './raid-roles';
import { DEFAULT_LOCALE, dateLabel } from './raid-time';

export type AttendanceGroupKey = 'noShow' | 'late' | 'joined' | 'present' | 'benchOut';

export interface AttendanceGroup {
  key: AttendanceGroupKey;
  label: string;
  rows: RaidAttendanceRow[];
  /** Present and Bench-and-out take the compact three-column rows (spec 3.5). */
  compact: boolean;
}

const GROUP_ORDER: readonly AttendanceGroupKey[] = ['noShow', 'late', 'joined', 'present', 'benchOut'];
const COMPACT_GROUPS: ReadonlySet<AttendanceGroupKey> = new Set(['present', 'benchOut']);

/** Signed In or Late: the characters the night's `signed` count and `Late` group are about. */
function attending(row: RaidAttendanceRow): boolean {
  return row.signed === 'in' || row.signed === 'late';
}

/** The one group a row sits in. `Late` holds only characters signed In or Late; a Bench or
 *  unsigned character seen late is `Joined`, never `Late` (spec 3.5). */
export function attendanceGroupOf(row: RaidAttendanceRow): AttendanceGroupKey {
  if (row.outcome === 'no_show') return 'noShow';
  if (row.outcome === 'absent') return 'benchOut';
  if (!attending(row)) return 'joined';
  return row.outcome === 'late' ? 'late' : 'present';
}

/** The role the sign-up gave, else the spec's; `null` when neither knows (such a row sorts last). */
function roleOf(row: RaidAttendanceRow): GuildRaidRole | null {
  return row.role ?? roleForSpec(row.spec);
}

function inRosterOrder(rows: readonly RaidAttendanceRow[]): RaidAttendanceRow[] {
  return rows
    .map((row) => ({ row, key: { role: roleOf(row), class: row.class, name: row.name } }))
    .sort((a, b) => rosterOrder(a.key, b.key))
    .map(({ row }) => row);
}

/** The attendance rows in the spec's five groups, in order, each in role, class, name order;
 *  an empty group is left out. */
export function groupAttendance(rows: readonly RaidAttendanceRow[]): AttendanceGroup[] {
  const labels = attendanceCopy.attendance.groups;
  return GROUP_ORDER.map((key) => ({
    key,
    label: labels[key],
    rows: inRosterOrder(rows.filter((row) => attendanceGroupOf(row) === key)),
    compact: COMPACT_GROUPS.has(key),
  })).filter((group) => group.rows.length > 0);
}

/** The Excused checkbox shows on `Not in the log` and `Late` rows only (spec 3.5). */
export function canExcuse(row: RaidAttendanceRow): boolean {
  return row.outcome === 'no_show' || row.outcome === 'late';
}

/** The first `REPORT_ROWS_BEFORE_SHOW_ALL` rows, or all of them once expanded. */
export function visibleRows<T>(rows: readonly T[], expanded: boolean): readonly T[] {
  return expanded ? rows : rows.slice(0, REPORT_ROWS_BEFORE_SHOW_ALL);
}

/** The viewer's own rows, the header character first (spec 3.3); the rest keep their order. */
export function yourNightRows(
  rows: readonly RaidAttendanceRow[],
  headerCharacterKey: string | null,
): RaidAttendanceRow[] {
  const own = rows.filter((row) => row.own);
  return [
    ...own.filter((r) => r.character_key === headerCharacterKey),
    ...own.filter((r) => r.character_key !== headerCharacterKey),
  ];
}

function signedWord(row: RaidAttendanceRow): string {
  return row.signed === null ? '' : raidCopy.segments[row.signed];
}

function roleLabelOf(row: RaidAttendanceRow): string | null {
  const role = roleOf(row);
  return role === null ? null : raidCopy.roleLabels[role];
}

/** Minutes after the first pull a late row was first seen; `null` when the log gave no offset. */
function lateMinutes(row: RaidAttendanceRow): number | null {
  return row.first_seen_offset_s === null ? null : Math.round(row.first_seen_offset_s / SECONDS_PER_MINUTE);
}

/** The attendance row's outcome cell (spec 3.5): `Not in the log`, `Came as Grimtotem`,
 *  `Late, first seen 22 min after the first pull`, `Present`; nothing for a Bench or Out character
 *  who stayed away. */
export function outcomeText(row: RaidAttendanceRow): string {
  if (row.came_as !== null) return cameAsOutcome(row.came_as);
  if (row.outcome === 'no_show') return attendanceCopy.attendance.notInLog;
  if (row.outcome === 'absent') return '';
  const minutes = lateMinutes(row);
  if (row.outcome === 'late' && minutes !== null) return lateOutcome(minutes);
  return attendanceCopy.attendance.present;
}

/** The outcome sentence of `Your night` (spec 6.1). */
export function yourNightSentence(row: RaidAttendanceRow): string {
  const state = signedWord(row);
  const { pulls_present: n, pulls_total: total } = row;
  if (row.signed === null) return yourNightText.unsigned(n, total);
  if (row.outcome === 'no_show') return yourNightText.notInLog(state, roleLabelOf(row));
  if (row.outcome === 'absent') return yourNightText.stayedOut(state);
  if (!attending(row)) return yourNightText.joined(state, n, total);
  if (row.came_as !== null) return yourNightText.cameAs(state, row.name, row.came_as, n, total);
  const minutes = lateMinutes(row);
  if (row.outcome === 'late' && minutes !== null) {
    return yourNightText.late(state, roleLabelOf(row), minutes, n, total);
  }
  return yourNightText.present(state, roleLabelOf(row), n, total);
}

/** The muted facts under the sentence; `rising` colours `up 53` in `--kill`. */
export type YourNightFact =
  | { kind: 'deaths'; text: string }
  | { kind: 'move'; lead: string; change: string; tail: string; rising: boolean }
  | { kind: 'excused'; text: string };

function moveFact(move: RaidMoveRow, previousDate: string): YourNightFact {
  const { rising, size } = moveDelta(move.before, move.after);
  return {
    kind: 'move',
    lead: `${metricLabel(move.metric)} ${move.after} on ${move.encounter}, `,
    change: `${rising ? 'up' : 'down'} ${size}`,
    tail: ` since ${previousDate}`,
    rising,
  };
}

function ownMoves(moves: RaidMoves, characterKey: string): RaidMoveRow[] {
  return moves.rows.filter((move) => move.character_key === characterKey && move.after !== move.before);
}

/** `2 deaths`, the moves against the previous raid, `Excused`: each only when it applies. */
export function yourNightFacts(
  row: RaidAttendanceRow,
  report: RaidReport | null,
  tz: string,
  locale: string = DEFAULT_LOCALE,
): YourNightFact[] {
  const facts: YourNightFact[] = [];
  const died = report?.deaths.find((d) => d.character_key === row.character_key);
  const deathCount = died === undefined ? 0 : died.on_wipes + died.on_kills;
  if (deathCount > 0) facts.push({ kind: 'deaths', text: deaths(deathCount) });
  if (report?.moves.previous != null) {
    const previousDate = dateLabel(new Date(report.moves.previous.starts_at), tz, locale);
    for (const move of ownMoves(report.moves, row.character_key)) facts.push(moveFact(move, previousDate));
  }
  if (row.excused) facts.push({ kind: 'excused', text: attendanceCopy.yourNight.excused });
  return facts;
}

export interface DiscordReportInput {
  title: string;
  startsAt: string;
  guildName: string;
  tz: string;
  night: RaidNight;
  report: RaidReport;
  /** The raid page's absolute URL with the report anchor. */
  reportUrl: string;
  locale?: string;
}

function summaryLine(night: RaidNight): string {
  const figures = [
    night.kills > 0 ? kills(night.kills) : null,
    night.wipes > 0 ? wipes(night.wipes) : null,
    night.duration_ms > 0 ? durationLabel(night.duration_ms) : null,
  ].filter((part): part is string => part !== null);
  const head = figures.length === 0 ? '' : `${figures.join(', ')}. `;
  return `${head}${cameRatio(night)}.`;
}

type KilledEntry = RaidKill & { kill_ms: number };

function isKilled(entry: RaidKill): entry is KilledEntry {
  return entry.killed && entry.kill_ms !== null;
}

function killsLine(report: RaidReport): string | null {
  const killed = report.kill_order.filter(isKilled);
  if (killed.length === 0) return null;
  const parts = killed.map((k) => `${k.name} ${killTimeLabel(k.kill_ms)} (${killAfter(k.wipes_before)})`);
  return `${discordText.kills} ${parts.join(', ')}`;
}

const MOST_DEATHS_NAMED = 3;

function deathsLine(report: RaidReport): string | null {
  const ranked = report.deaths
    .map((d) => ({ name: d.name, count: d.on_wipes + d.on_kills }))
    .filter((d) => d.count > 0)
    .sort((a, b) => b.count - a.count)
    .slice(0, MOST_DEATHS_NAMED);
  return ranked.length === 0
    ? null
    : `${discordText.mostDeaths} ${ranked.map((d) => `${d.name} ${d.count}`).join(', ')}`;
}

function lootLines(report: RaidReport): string[] {
  return report.loot
    .filter((group) => group.awards.length > 0)
    .map(
      (group) =>
        `${discordText.loot} ${group.boss}: ${group.awards.map((a) => `${a.item_name} to ${a.name}`).join('; ')}`,
    );
}

/** Rises only: a fall is never posted (spec 4). */
function changesLine(report: RaidReport): string | null {
  const rises = report.moves.rows
    .filter((m) => m.after > m.before)
    .sort((a, b) => b.after - b.before - (a.after - a.before))
    .slice(0, BIGGEST_CHANGES_ROWS);
  if (rises.length === 0) return null;
  const parts = rises.map((m) => `${m.name} ${metricLabel(m.metric)} ${moveDelta(m.before, m.after).text}`);
  return `${discordText.changes} ${parts.join(', ')}`;
}

function present(lines: readonly (string | null)[]): string[] {
  return lines.filter((line): line is string => line !== null);
}

/** The loot lines that fit, the last of them `+N more` when the cap leaves some out. */
function fitLoot(lines: readonly string[], room: number): string[] {
  if (lines.length <= room) return [...lines];
  const shown = Math.max(0, room - 1);
  return room === 0 ? [] : [...lines.slice(0, shown), discordText.more(lines.length - shown)];
}

/** Plain text in guild time, at most `DISCORD_REPORT_MAX_LINES` lines, no absentee names (R-A3).
 *  Loot lines give way first when a night has more bosses than the cap leaves room for. */
export function discordReportText(input: DiscordReportInput): string {
  const { report } = input;
  const locale = input.locale ?? DEFAULT_LOCALE;
  const head = [
    [input.title, dateLabel(new Date(input.startsAt), input.tz, locale), input.guildName].join(SEP),
    summaryLine(input.night),
  ];
  const beforeLoot = present([...head, killsLine(report), deathsLine(report)]);
  const afterLoot = present([changesLine(report), `${discordText.report} ${input.reportUrl}`]);
  const lootRoom = Math.max(0, DISCORD_REPORT_MAX_LINES - beforeLoot.length - afterLoot.length);
  return [...beforeLoot, ...fitLoot(lootLines(report), lootRoom), ...afterLoot].join('\n');
}

/** Whole days a sync was before the raid (`Synced 9 days before the raid`), never negative. */
export function daysBeforeRaid(loggedAt: string, raidStartsAt: string): number {
  const gap = new Date(raidStartsAt).getTime() - new Date(loggedAt).getTime();
  return Math.max(0, Math.floor(gap / DAY_MS));
}
