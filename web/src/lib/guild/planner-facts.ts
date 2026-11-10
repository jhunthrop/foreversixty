// web/src/lib/guild/planner-facts.ts
// The figures each sign-up carries on the raid planner (guild dense spec, h-guild-tabs): item level
// against the guild's median, the raid's readiness checks it fails (officers, from the readiness
// read for this raid), attendance over the last 30 nights and the raider's rating. They come from
// reads the page already has (the guild's roster on /home, readiness?raid=), keyed by character.
import type { GuildReadinessRow, GuildRosterRow } from './api';

export interface PlannerFacts {
  itemLevel: number | null;
  /** Against the guild's median item level; null without both. */
  itemLevelDelta: number | null;
  /** Readiness checks failed for this raid; null when the viewer has no readiness read. */
  short: number | null;
  /** Nights present of the last 30 days' nights, as a percentage; null with none. */
  attendancePct: number | null;
  rating: number | null;
}

/** The median of the guild's known item levels, rounded; null with none. */
export function medianItemLevel(roster: readonly GuildRosterRow[]): number | null {
  const levels = roster
    .map((row) => row.item_level)
    .filter((v): v is number => typeof v === 'number')
    .toSorted((a, b) => a - b);
  return levels.length === 0 ? null : levels[Math.floor(levels.length / 2)];
}

export function plannerFacts(
  roster: readonly GuildRosterRow[],
  readiness: readonly GuildReadinessRow[] | null,
): ReadonlyMap<string, PlannerFacts> {
  const median = medianItemLevel(roster);
  const failing = new Map((readiness ?? []).map((row) => [row.character_key, row.failing]));
  return new Map(
    roster.map((row) => {
      const ilvl = typeof row.item_level === 'number' ? row.item_level : null;
      const att = row.attendance_30 ?? row.attendance;
      return [
        row.character_key,
        {
          itemLevel: ilvl,
          itemLevelDelta: ilvl === null || median === null ? null : ilvl - median,
          short: readiness === null ? null : (failing.get(row.character_key) ?? null),
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

/** The figures on one line where the columns fold (below 1024): `64 ilvl · 1 short · rating 77 · 80%`. */
export function phoneFacts(facts: PlannerFacts): string {
  return [
    facts.itemLevel === null ? '' : `${facts.itemLevel} ilvl`,
    facts.short === null || facts.short === 0 ? '' : `${facts.short} short`,
    facts.rating === null ? '' : `rating ${facts.rating}`,
    facts.attendancePct === null ? '' : `${facts.attendancePct}%`,
  ]
    .filter((part) => part !== '')
    .join(' · ');
}
