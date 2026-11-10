// web/src/lib/guides/glance.ts
// The guide's lane rules (design/specs/2026-11-08-guides.md, rulings 0A): the leveller's next
// talent point, the cheap ranks, the level 60 rows and the Racials lane in faction groups. The At a
// glance plate is the build alone (h-audit 2026-11-09: the Rotation and Stat priority sections own
// those lists). Pure over the page's own models, so the page and the mock draw one truth.
import { readBuildJson } from './build-files';
import type { GuideRaceRow } from './race-table';
import type { RotationItem } from './rotation-lists';
import type { LadderStep } from './leveling-builds';

/** A clause segment that names a level (`at 26 to 29`, `from 35`, `to 49`) is levelling advice. */
const LEVEL_WORDS = /\b(?:at|to|from)\s+\d{2}\b/;

/**
 * The clause as it reads at level 60: its first segment, and every later comma segment that
 * names no level. "On every swing it is ready, with an older rank on your bar at 26 to 29, ..."
 * reads "On every swing it is ready"; rank swaps while levelling stay in the Leveling prose.
 */
export function level60Clause(clause: string): string {
  const [first = '', ...rest] = clause.split(', ');
  return [first, ...rest.filter((segment) => !LEVEL_WORDS.test(segment))].join(', ');
}

/** The next level after `level` whose point the ladder places, or undefined at the end. */
export function nextTalentPoint(
  steps: readonly LadderStep[],
  level: number,
): { level: number; name: string } | undefined {
  const next = steps.find((step) => step.level > level && step.name !== null);
  return next === undefined || next.name === null ? undefined : { level: next.level, name: next.name };
}

/** The Racials lane: Alliance then Horde, the recommended races first in each, otherwise in
 *  the build's order. */
export function racesByFaction(
  rows: readonly GuideRaceRow[],
): { faction: 'alliance' | 'horde'; rows: GuideRaceRow[] }[] {
  return (['alliance', 'horde'] as const)
    .map((faction) => {
      const ofFaction = rows.filter((row) => row.faction === faction);
      return {
        faction,
        rows: [...ofFaction.filter((row) => row.recommended), ...ofFaction.filter((row) => !row.recommended)],
      };
    })
    .filter((group) => group.rows.length > 0);
}

interface SpellRankEntry {
  id: number;
  level: number;
  rank: number;
}

/** The highest rank of `name` a class learns by level 60 (`spellranks.json`); undefined for a
 *  spell with no ranks. */
export function topRankAt60(build: string, classSlug: string, name: string): number | undefined {
  const file = readBuildJson<{ classes: Record<string, Record<string, SpellRankEntry[]>> }>(
    build,
    'spellranks.json',
  );
  const ranks = (file?.classes[classSlug]?.[name] ?? []).filter(
    (entry) => entry.level <= 60 && entry.rank > 0,
  );
  return ranks.length === 0 ? undefined : Math.max(...ranks.map((entry) => entry.rank));
}

/** A line the guide casts below the top rank it could, and whose rotation prose names that
 *  rank ("Renew always up at rank 7"): a deliberate cheap rank, said so. A lower rank the prose
 *  never chose (a top rank that only a raid book teaches) is not called cheap. */
export function isCheapRank(item: RotationItem, topRank: number | undefined, prose: string): boolean {
  if (item.rank === undefined || topRank === undefined || item.rank >= topRank) return false;
  const name = item.name.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  return new RegExp(`\\b${name}\\b[^.]*?\\brank ${item.rank}\\b`, 'i').test(prose);
}

/** The rotation lane's lists at their level 60 clauses (level60Clause): the rank swaps while
 *  levelling ("with an older rank on your bar at 26 to 29, ...") stay in the Leveling prose, so
 *  every priority row reads as one line (P8, 2026-11-09). */
export function atLevel60<
  T extends { opener: RotationItem[]; priority: RotationItem[]; cooldowns: RotationItem[] },
>(lists: T): T {
  const at60 = (items: RotationItem[]): RotationItem[] =>
    items.map((item) => ({ ...item, clause: level60Clause(item.clause) }));
  return {
    ...lists,
    opener: at60(lists.opener),
    priority: at60(lists.priority),
    cooldowns: at60(lists.cooldowns),
  };
}

/** Marks every row the guide casts at a cheap rank on purpose (`cheapRank`). */
export function markCheapRanks<
  T extends { opener: RotationItem[]; priority: RotationItem[]; cooldowns: RotationItem[] },
>(lists: T, topRank: (name: string) => number | undefined, prose: string): T {
  const mark = (items: RotationItem[]): RotationItem[] =>
    items.map((item) => (isCheapRank(item, topRank(item.name), prose) ? { ...item, cheapRank: true } : item));
  return {
    ...lists,
    opener: mark(lists.opener),
    priority: mark(lists.priority),
    cooldowns: mark(lists.cooldowns),
  };
}
