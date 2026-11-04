// web/src/lib/guides/leveling-talents.ts
// A leveling range's talents as the list a player works through, in the order the points are
// spent (the ladder's own point order, `guideLadderOrder` in leveling-builds.ts): the talents a
// range carries from earlier ranges by their first point, then the ones it adds by the level it
// adds them at, so a row reads left to right in taking order. A range lists what a player holds by its LAST level, and each talent names the
// levels its points come at inside the range (the ladder spends one point a level from 10),
// so a reader anywhere in "Level 30 to 39" sees every point up to their own level. A talent
// whose rank rose since the range before is marked. Pure over the decoded trees
// `decodeBandTalents` produces.
import { FIRST_POINT_LEVEL } from '../planner/types';
import type { DecodedTreeView } from './build-tree';

export interface LevelingTalentChip {
  id: number;
  name: string;
  icon: string;
  rank: number;
  maxRank: number;
  /** True when this range spends more points here than the range before it. */
  isNew: boolean;
  /** The levels, ascending, whose point goes here, from the range's `from` to its `to`. */
  levels: number[];
}

/** One range of the strip: the trees held at its last level (`undefined` when they did not
 *  decode) and the levels whose points its list names. */
export interface LevelingChipRange {
  trees: readonly DecodedTreeView[] | undefined;
  from: number;
  to: number;
}

/** One chip list per range, ranges in level order; a range whose talents did not decode gives
 *  an empty list. `spendOrder` is one talent id per point, the point at index i taken at level
 *  10 + i; a lit talent it does not name follows the ones it does, by tree position, row and
 *  column. */
export function levelingChipRows(
  ranges: readonly LevelingChipRange[],
  spendOrder: readonly number[],
): LevelingTalentChip[][] {
  const firstPoint = new Map<number, number>();
  spendOrder.forEach((id, point) => {
    if (!firstPoint.has(id)) firstPoint.set(id, point);
  });
  const ranks = new Map<number, number>();
  return ranges.map(({ trees, from, to }, rangeIndex) => {
    if (trees === undefined) return [];
    const lit = [...trees]
      .sort((a, b) => a.position - b.position)
      .flatMap((tree) =>
        [...tree.talents]
          .filter((talent) => talent.rank > 0)
          .sort((a, b) => a.tier - b.tier || a.column - b.column),
      );
    const unnamed = spendOrder.length;
    const place = (id: number, fallback: number): number => firstPoint.get(id) ?? unnamed + fallback;
    const chips = lit
      .map((talent, fallback) => ({ talent, at: place(talent.id, fallback) }))
      .sort((a, b) => a.at - b.at)
      .map(({ talent }): LevelingTalentChip => ({
        id: talent.id,
        name: talent.name,
        icon: talent.icon,
        rank: talent.rank,
        maxRank: talent.max_rank,
        isNew: rangeIndex > 0 && talent.rank > (ranks.get(talent.id) ?? 0),
        levels: levelsOf(spendOrder, talent.id, from, to),
      }));
    for (const talent of lit) ranks.set(talent.id, talent.rank);
    return inTakingOrder(chips, firstPoint, from);
  });
}

/** A row in the order a player takes its points: what it carries from points before `from`
 *  first (already in first-point order), then what it adds, by the first level it adds them at,
 *  then any talent the order places later or not at all. */
function inTakingOrder(
  chips: readonly LevelingTalentChip[],
  firstPoint: ReadonlyMap<number, number>,
  from: number,
): LevelingTalentChip[] {
  const added = chips.filter((chip) => chip.levels.length > 0);
  const carried = chips.filter(
    (chip) => chip.levels.length === 0 && FIRST_POINT_LEVEL + (firstPoint.get(chip.id) ?? Infinity) < from,
  );
  const rest = chips.filter((chip) => !added.includes(chip) && !carried.includes(chip));
  return [...carried, ...[...added].sort((a, b) => a.levels[0]! - b.levels[0]!), ...rest];
}

function levelsOf(spendOrder: readonly number[], id: number, from: number, to: number): number[] {
  return spendOrder.flatMap((spent, point) => {
    const level = FIRST_POINT_LEVEL + point;
    return spent === id && level >= from && level <= to ? [level] : [];
  });
}

/** Ascending levels folded into [first, last] runs: 30, 31, 32, 35 gives [30, 32], [35, 35]. */
export function levelRuns(levels: readonly number[]): [number, number][] {
  return levels.reduce<[number, number][]>((runs, level) => {
    const last = runs.at(-1);
    return last !== undefined && level === last[1] + 1
      ? [...runs.slice(0, -1), [last[0], level]]
      : [...runs, [level, level]];
  }, []);
}

/** Ascending levels as the runs a reader scans: 30, 31, 32, 35 gives "30-32, 35". */
export function levelRunsText(levels: readonly number[]): string {
  return levelRuns(levels)
    .map(([first, last]) => (first === last ? `${first}` : `${first}-${last}`))
    .join(', ');
}
