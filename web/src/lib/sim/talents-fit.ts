// web/src/lib/sim/talents-fit.ts
// Whether a stored talents string can be drawn on a build's trees (spec 2026-11-04-sim-results
// 4.7). The planner's own gates decide the prerequisites and the tier rule, through the order
// `orderFromRanks` can rebuild: a rank it could not place breaks one of them.
import { orderFromRanks } from '../planner/fs1';
import type { TalentIndex } from '../planner/rules';
import { ranksFromTalentsString } from './character';

const ONE_DIGIT_PER_TALENT = /^\d*$/;

/**
 * True when the string fits. A part may be shorter than its tree, because the engine's string
 * trims each tree's trailing zeros (`talentsString`, character.ts); a part longer than its tree,
 * a digit past a talent's top rank, or a rank whose gates are not met does not fit.
 */
export function talentsFit(talents: string, index: TalentIndex): boolean {
  const parts = talents.split('-');
  if (parts.length !== index.trees.length) return false;
  const shaped = parts.every(
    (part, tree) =>
      ONE_DIGIT_PER_TALENT.test(part) &&
      part.length <= index.trees[tree]!.talents.length &&
      Array.from(part).every((digit, at) => Number(digit) <= index.trees[tree]!.talents[at]!.max_rank),
  );
  return shaped && orderFromRanks(index, ranksFromTalentsString(talents)).dropped.length === 0;
}

/** The tree the phone opens on: the most points, the first on a tie. */
export function fullestTree(talents: string): number {
  const points = ranksFromTalentsString(talents).map((ranks) => ranks.reduce((sum, rank) => sum + rank, 0));
  return points.indexOf(Math.max(...points));
}
