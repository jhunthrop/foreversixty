// web/src/lib/guild/planner-facts.ts
// The figures each sign-up carries on the raid planner (guild dense spec, h-guild-tabs): the last
// five raid nights (officers, from attendance/recent), item level against the guild's median, the
// raid's readiness checks it fails and what they are (officers from readiness?raid=, a member's
// own row from their checklist), the raider's rating and attendance over the last 30 days. Every
// figure comes from a read the page makes anyway, keyed by character (attendance by account).
import type { GuildReadinessRow, GuildRosterRow } from './api';
import type { AttendanceCell, GuildChecklist, GuildRecentAttendance } from './overview-api';
import { counted } from '../plural';
import { readinessShortfalls } from './readiness-view';

export interface PlannerFacts {
  /** The last raid nights, oldest first; null without the officers' attendance read. */
  cells: readonly AttendanceCell[] | null;
  itemLevel: number | null;
  /** Against the guild's median item level; null without both. */
  itemLevelDelta: number | null;
  /** Readiness checks failed for this raid; null when the viewer has no read of this raider's. */
  short: number | null;
  /** What those checks are, in the readiness board's words. */
  why: readonly string[];
  /** Nights present of the last 30 days' nights, as a percentage; null with none. */
  attendancePct: number | null;
  rating: number | null;
}

/** The reads the planner joins: readiness (officers), attendance (officers), the viewer's checklist. */
export interface PlannerReads {
  readiness: readonly GuildReadinessRow[] | null;
  attendance: GuildRecentAttendance | null;
  checklist: GuildChecklist | null;
}

/** The median of the guild's known item levels; null with none. */
export function medianItemLevel(roster: readonly GuildRosterRow[]): number | null {
  const levels = roster
    .map((row) => row.item_level)
    .filter((v): v is number => typeof v === 'number')
    .toSorted((a, b) => a - b);
  return levels.length === 0 ? null : levels[Math.floor(levels.length / 2)];
}

/** A member's own checklist as the readiness board words its failures. */
export function checklistShortfalls(checklist: GuildChecklist): string[] {
  return checklist.items
    .filter((item) => item.state === 'fail')
    .flatMap((item): string[] => {
      switch (item.key) {
        case 'gear_gap':
          return [counted(item.upgrades ?? 0, 'upgrade')];
        case 'enchants':
          return [`${counted(item.missing_slots?.length ?? 0, 'enchant')} missing`];
        case 'flask_food':
          return (item.missing ?? []).map((m) => (m === 'food' ? 'no food' : 'no flask or elixir'));
        case 'attunement':
          return ['not attuned'];
        default:
          return [];
      }
    });
}

export function plannerFacts(
  roster: readonly GuildRosterRow[],
  reads: PlannerReads,
): ReadonlyMap<string, PlannerFacts> {
  const median = medianItemLevel(roster);
  const readiness = new Map((reads.readiness ?? []).map((row) => [row.character_key, row]));
  const nights = new Map((reads.attendance?.rows ?? []).map((row) => [row.account_key, row.cells]));
  const own = reads.checklist;
  return new Map(
    roster.map((row) => {
      const ilvl = typeof row.item_level === 'number' ? row.item_level : null;
      const att = row.attendance_30 ?? row.attendance;
      const ready = readiness.get(row.character_key);
      const why =
        ready !== undefined
          ? readinessShortfalls(ready, { attunement: true }, false)
          : own?.character.character_key === row.character_key
            ? checklistShortfalls(own)
            : null;
      return [
        row.character_key,
        {
          cells: row.account_key === undefined ? null : (nights.get(row.account_key) ?? null),
          itemLevel: ilvl,
          itemLevelDelta: ilvl === null || median === null ? null : ilvl - median,
          short: why === null ? null : why.length,
          why: why ?? [],
          attendancePct:
            att === undefined || att.nights === 0 ? null : Math.round((100 * att.present) / att.nights),
          rating: row.rating?.overall ?? null,
        },
      ];
    }),
  );
}

/** `+2`, `−1`, `+0`. */
export function signed(delta: number): string {
  return delta < 0 ? `−${Math.abs(delta)}` : `+${delta}`;
}

/** The figures on one line where the columns fold (below 1024): `64 ilvl · rating 77 · 80%`. */
export function phoneFacts(facts: PlannerFacts): string {
  return [
    facts.itemLevel === null ? '' : `${facts.itemLevel} ilvl`,
    facts.rating === null ? '' : `rating ${facts.rating}`,
    facts.attendancePct === null ? '' : `${facts.attendancePct}%`,
  ]
    .filter((part) => part !== '')
    .join(' · ');
}
