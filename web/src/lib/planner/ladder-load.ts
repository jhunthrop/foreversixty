// web/src/lib/planner/ladder-load.ts
// The spec ladder orders (/data/leveling-plans.json) and the prefix check that spends a band in
// one, loaded on demand: only a band load needs them, so they stay out of planner-island.js.
import { bandTreeRanksFrom } from './band-compare';
import { validateOrder, type TalentIndex } from './rules';
import { loadLevelingPlans } from '../leveling/plans-file';

/** The ladder order for `spec`'s band `bandTalents`, or undefined to reconstruct one. */
export type LadderOrderFor = (index: TalentIndex, bandTalents: string, spec: string) => number[] | undefined;

/** The first points of `ladder` that spend exactly `treeRanks`, when that prefix is a legal
 *  order (`validateOrder`); undefined when the band is not a point of this ladder. */
export function ladderPrefixFor(
  index: TalentIndex,
  treeRanks: readonly (readonly number[])[],
  ladder: readonly number[],
): number[] | undefined {
  const wanted = new Map<number, number>();
  index.trees.forEach((tree, ti) =>
    tree.talents.forEach((talent, cell) => {
      const rank = treeRanks[ti]?.[cell] ?? 0;
      if (rank > 0) wanted.set(talent.id, rank);
    }),
  );
  const points = [...wanted.values()].reduce((sum, rank) => sum + rank, 0);
  if (points === 0 || points > ladder.length) return undefined;
  const prefix = ladder.slice(0, points);
  const got = new Map<number, number>();
  for (const id of prefix) got.set(id, (got.get(id) ?? 0) + 1);
  const same = got.size === wanted.size && [...wanted].every(([id, rank]) => got.get(id) === rank);
  return same && validateOrder(index, prefix).length === 0 ? prefix : undefined;
}

/** Reads the plans once and answers band loads from them; never fails (no ladder: reconstruct). */
export async function loadLadderOrderFor(): Promise<LadderOrderFor> {
  const plans = await loadLevelingPlans();
  return (index, bandTalents, spec) => {
    const ladder = plans[spec]?.ladder;
    return ladder === undefined ? undefined : ladderPrefixFor(index, bandTreeRanksFrom(bandTalents), ladder);
  };
}
