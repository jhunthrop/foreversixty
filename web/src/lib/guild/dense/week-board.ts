// web/src/lib/guild/dense/week-board.ts
// The officer's week on the Overview (design/specs/2026-11-09-guild-dense.md section 5.1): the next
// raid's seats against the comp target with the raiders who could fill each gap, who needs a word
// before it, last night's risers and fallers by rating, the roles still short if everyone answers,
// and who is slipping. Pure: built from the reads the Overview already makes.
import type { GuildRaidGroup, GuildRaidView, GuildReadinessRow, GuildRosterRow, GuildRosterRating } from '../api';
import type { RaidReportParse } from '../attendance-contract';
import type { AttendanceCell, GuildRecentAttendance } from '../overview-api';
import { RAID_GROUP_ORDER } from './roster-board';

/** How many filler chips a gap row shows before "+N more". */
export const FILLERS_SHOWN = 3;
/** Missed nights among the last SLIP_WINDOW that make a raider slipping (attendance_recent.go). */
export const SLIP_MISSES = 2;
export const SLIP_WINDOW = 4;
/** Benched this many nights or more puts a raider on the Slipping list too. */
export const BENCH_NIGHTS = 2;
/** How many raiders each side of last night's rating changes lists. */
export const MOVERS_EACH = 3;
/** A rating fall of this many points is worth a word. */
export const WORD_RATING_FALL = 10;
/** Days without an export that are worth a word. */
export const WORD_STALE_DAYS = 7;
const DAY_MS = 86_400_000;

const TARGET_KEY = { tank: 'tanks', healer: 'healers', melee: 'melee', ranged: 'ranged' } as const;

export interface Filler {
  key: string;
  name: string;
  cls: string;
  /** Not answered: Ping. Benched: move to In. */
  kind: 'ping' | 'bench';
}

export interface CompGapRow {
  group: GuildRaidGroup;
  need: number;
  in: number;
  late: number;
  /** Seats open counting Late as coming; 0 when full. */
  gap: number;
  fillers: Filler[];
  /** Seats still open if every filler says yes. */
  stillShort: number;
}

function groupCount(view: GuildRaidView, group: GuildRaidGroup, state: 'in' | 'late'): number {
  return view.signups.filter((s) => s.state === state && (s.group ?? 'ranged') === group).length;
}

/** The raid's seats per role against its target, with the fillers of each open seat. */
export function compGaps(view: GuildRaidView): CompGapRow[] {
  const target = view.comp_target;
  if (target === null || target === undefined) return [];
  return RAID_GROUP_ORDER.map((group) => {
    const need = target[TARGET_KEY[group]];
    const inCount = groupCount(view, group, 'in');
    const late = groupCount(view, group, 'late');
    const gap = Math.max(0, need - inCount - late);
    const fillers: Filler[] = [
      ...view.not_signed
        .filter((n) => (n.group ?? 'ranged') === group)
        .map((n) => ({ key: n.character_key, name: n.name, cls: n.class, kind: 'ping' as const })),
      ...view.signups
        .filter((s) => s.state === 'bench' && (s.group ?? 'ranged') === group)
        .map((s) => ({ key: s.character_key, name: s.name, cls: s.class, kind: 'bench' as const })),
    ];
    return { group, need, in: inCount, late, gap, fillers: gap > 0 ? fillers : [], stillShort: Math.max(0, gap - fillers.length) };
  });
}

export function seatsOpen(rows: readonly CompGapRow[]): number {
  return rows.reduce((sum, row) => sum + row.gap, 0);
}

export interface Mover {
  key: string;
  name: string;
  cls: string;
  spec: string;
  rating: number;
  change: number;
  value: number;
  metric: string;
  weakest: { part: keyof Omit<GuildRosterRating, 'overall'>; value: number } | null;
}

const PARTS = ['output', 'survival', 'mechanics', 'utility', 'preparation', 'activity'] as const;

export function weakestPart(rating: GuildRosterRating | null | undefined): Mover['weakest'] {
  if (rating === null || rating === undefined) return null;
  const part = PARTS.reduce((low, p) => (rating[p] < rating[low] ? p : low), PARTS[0]);
  return { part, value: rating[part] };
}

/** One rated parse per raider on a night: their best rating among the night's kills. */
function bestByRaider(parses: readonly RaidReportParse[]): Map<string, RaidReportParse> {
  const out = new Map<string, RaidReportParse>();
  for (const p of parses) {
    if (p.rating === null) continue;
    const had = out.get(p.character_key);
    if (had === undefined || (had.rating ?? 0) < p.rating) out.set(p.character_key, p);
  }
  return out;
}

/** Last night's raiders by rating change against the night before: risers then fallers. */
export function ratingMovers(
  last: readonly RaidReportParse[],
  before: readonly RaidReportParse[],
  roster: readonly GuildRosterRow[],
): { up: Mover[]; down: Mover[] } {
  const prev = bestByRaider(before);
  const byKey = new Map(roster.map((r) => [r.character_key, r] as const));
  const moved: Mover[] = [];
  for (const p of bestByRaider(last).values()) {
    const was = prev.get(p.character_key);
    if (was === undefined || was.rating === null || p.rating === null) continue;
    const row = byKey.get(p.character_key);
    moved.push({
      key: p.character_key,
      name: p.name,
      cls: p.class,
      spec: row?.spec_name ?? row?.spec ?? '',
      rating: Math.round(p.rating),
      change: Math.round(p.rating - was.rating),
      value: p.value,
      metric: p.metric,
      weakest: weakestPart(row?.rating),
    });
  }
  const sorted = [...moved].sort((a, b) => b.change - a.change || a.name.localeCompare(b.name));
  return {
    up: sorted.filter((m) => m.change > 0).slice(0, MOVERS_EACH),
    down: sorted
      .filter((m) => m.change < 0)
      .reverse()
      .slice(0, MOVERS_EACH),
  };
}

export interface SlipRow {
  name: string;
  cls: string;
  cells: readonly AttendanceCell[];
  present: number;
  nights: number;
}

/** Raiders who missed SLIP_MISSES of the last SLIP_WINDOW nights, or sat on the bench often; most misses first. */
export function slipping(grid: GuildRecentAttendance | null): SlipRow[] {
  if (grid === null) return [];
  const misses = (cells: readonly AttendanceCell[]) => cells.slice(-SLIP_WINDOW).filter((c) => c === 'absent').length;
  return grid.rows
    .filter((r) => misses(r.cells) >= SLIP_MISSES || r.cells.filter((c) => c === 'bench').length >= BENCH_NIGHTS)
    .map((r) => ({
      name: r.name,
      cls: r.class,
      cells: r.cells,
      present: r.cells.filter((c) => c === 'present').length,
      nights: r.cells.filter((c) => c !== 'none').length,
    }))
    .sort((a, b) => misses(b.cells) - misses(a.cells) || a.name.localeCompare(b.name));
}

export type WordAction = 'profile' | 'ping' | 'nudge' | 'report';

export interface Word {
  key: string;
  name: string;
  cls: string;
  reasons: string[];
  action: WordAction;
}

export interface WordInputs {
  view: GuildRaidView | null;
  roster: readonly GuildRosterRow[];
  grid: GuildRecentAttendance | null;
  readiness: readonly GuildReadinessRow[];
  movers: { up: Mover[]; down: Mover[] };
  now: Date;
}

const SHORT_WORDS: ReadonlyArray<[(r: GuildReadinessRow) => boolean, (r: GuildReadinessRow) => string]> = [
  [(r) => r.attuned === false, () => 'not attuned'],
  [(r) => r.enchants.missing_slots.length > 0, (r) => `${r.enchants.missing_slots.length} enchants missing`],
  [(r) => r.flask_food != null && (!r.flask_food.flask_or_elixir || !r.flask_food.food), () => 'no flask or food'],
  [(r) => r.talent_points_read !== false && r.talent_points_unspent > 0, (r) => `${r.talent_points_unspent} talent points unspent`],
];

export function shortReasons(row: GuildReadinessRow): string[] {
  return SHORT_WORDS.filter(([fails]) => fails(row)).map(([, word]) => word(row));
}

function daysSince(iso: string | undefined, now: Date): number | null {
  return iso === undefined || iso === '' ? null : Math.floor((now.getTime() - new Date(iso).getTime()) / DAY_MS);
}

/** Who the officer talks to before the raid, with the reason in figures and one action each. */
export function needsAWord(inputs: WordInputs): Word[] {
  const words = new Map<string, Word>();
  const add = (key: string, name: string, cls: string, reason: string, action: WordAction) => {
    const had = words.get(key);
    if (had === undefined) words.set(key, { key, name, cls, reasons: [reason], action });
    else had.reasons.push(reason);
  };
  const missesByName = new Map(
    (inputs.grid?.rows ?? []).map((r) => [r.name, r.cells.slice(-SLIP_WINDOW).filter((c) => c === 'absent').length]),
  );
  const rosterByKey = new Map(inputs.roster.map((r) => [r.character_key, r] as const));
  for (const n of inputs.view?.not_signed_raiders ?? []) {
    const misses = missesByName.get(n.name) ?? 0;
    if (misses < SLIP_MISSES) continue;
    const row = rosterByKey.get(n.character_key);
    const stale = daysSince(row?.logged_at, inputs.now);
    add(n.character_key, n.name, n.class, 'Not answered', n.discord_linked ? 'ping' : 'profile');
    add(n.character_key, n.name, n.class, `missed ${misses} of the last ${SLIP_WINDOW}`, 'ping');
    if (stale !== null && stale >= WORD_STALE_DAYS) add(n.character_key, n.name, n.class, `last log ${stale} d ago`, 'ping');
    if (!n.discord_linked) add(n.character_key, n.name, n.class, 'no Discord', 'profile');
  }
  for (const r of inputs.readiness) {
    const why = shortReasons(r);
    if (why.length > 0 && !words.has(r.character_key)) for (const w of why) add(r.character_key, r.name, r.class, w, 'nudge');
  }
  for (const m of inputs.movers.down) {
    if (-m.change < WORD_RATING_FALL || words.has(m.key)) continue;
    const part = m.weakest === null ? '' : ` · ${m.weakest.part} ${m.weakest.value}`;
    add(m.key, m.name, m.cls, `Rating ${m.rating - m.change} to ${m.rating}${part}`, 'report');
  }
  return [...words.values()];
}

/** A role still short if every filler answers yes. */
export interface RecruitGap {
  group: GuildRaidGroup;
  short: number;
}

export function recruitGaps(rows: readonly CompGapRow[]): RecruitGap[] {
  return rows.filter((r) => r.stillShort > 0).map((r) => ({ group: r.group, short: r.stillShort }));
}
