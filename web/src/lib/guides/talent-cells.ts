// web/src/lib/guides/talent-cells.ts
// A decoded build as the TalentTrees primitive draws it: each talent's cell, its rank, and the
// game tooltip key at that rank (rank 1 for an unlearned talent) when the guide's tooltip
// bundle carries the talent.
import type { TalentTree } from '../../components/ui/TalentTrees.svelte';
import { talentTipKey, type GuideTipBundle } from '../tooltips/game-tip';
import type { DecodedTreeView } from './build-tree';

export function talentTreesOf(
  trees: readonly DecodedTreeView[],
  bundle?: Pick<GuideTipBundle, 'talents'>,
): TalentTree[] {
  return trees.map((tree) => ({
    name: tree.name,
    points: tree.points,
    talents: tree.talents.map((talent) => ({
      tier: talent.tier,
      column: talent.column,
      name: talent.name,
      icon: talent.icon,
      rank: talent.rank,
      maxRank: talent.max_rank,
      ...(bundle?.talents[String(talent.id)] === undefined
        ? {}
        : { tipKey: talentTipKey(talent.id, Math.max(1, talent.rank)) }),
    })),
  }));
}

/** The trees at a ladder step: the step's talent string ("a-b-c", one digit per talent in each
 *  tree's tier-then-column order, `ladderTalentString`) over the build's own layout. */
export function treesAtStep(trees: readonly TalentTree[], talents: string): TalentTree[] {
  const digits = talents.split('-');
  return trees.map((tree, index) => {
    const ranks = digits[index] ?? '';
    const cells = tree.talents.map((cell, i) => ({ ...cell, rank: Number(ranks[i] ?? 0) || 0 }));
    return { ...tree, points: cells.reduce((sum, cell) => sum + cell.rank, 0), talents: cells };
  });
}
