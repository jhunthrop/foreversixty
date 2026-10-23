// web/src/lib/guides/leveling-talents.ts
// A leveling band's talents as the list a player works through, in the order to take them: the
// talents of the band before keep their places, and what the band adds comes after them, the
// build's main tree first and each tree from its first row down. A talent whose rank rose since
// the band before is marked, so the strip also shows what each band adds. Pure over the decoded
// trees `decodeBandTalents` produces.
import type { DecodedTreeView } from './build-tree';

export interface LevelingTalentChip {
  name: string;
  icon: string;
  rank: number;
  maxRank: number;
  /** True when this band spends more points here than the band before it. */
  isNew: boolean;
}

/** One chip list per band, bands in level order; `undefined` (a band whose talents did not
 *  decode) gives an empty list and does not break the order for the next one. */
export function levelingChipRows(
  bands: readonly (readonly DecodedTreeView[] | undefined)[],
): LevelingTalentChip[][] {
  const order: number[] = [];
  const ranks = new Map<number, number>();
  return bands.map((trees, bandIndex) => {
    if (trees === undefined) return [];
    const lit = new Map<number, { talent: DecodedTreeView['talents'][number]; main: number }>();
    const byPoints = trees
      .map((tree, index) => ({ tree, index }))
      .sort((a, b) => b.tree.points - a.tree.points || a.index - b.index);
    byPoints.forEach(({ tree }, main) => {
      for (const talent of tree.talents) if (talent.rank > 0) lit.set(talent.id, { talent, main });
    });
    const added = [...lit.values()]
      .filter(({ talent }) => !order.includes(talent.id))
      .sort((a, b) => a.main - b.main || a.talent.tier - b.talent.tier || a.talent.column - b.talent.column);
    for (const { talent } of added) order.push(talent.id);
    const chips = order
      .filter((id) => lit.has(id))
      .map((id): LevelingTalentChip => {
        const { talent } = lit.get(id)!;
        return {
          name: talent.name,
          icon: talent.icon,
          rank: talent.rank,
          maxRank: talent.max_rank,
          isNew: bandIndex > 0 && talent.rank > (ranks.get(id) ?? 0),
        };
      });
    for (const { talent } of lit.values()) ranks.set(talent.id, talent.rank);
    return chips;
  });
}
