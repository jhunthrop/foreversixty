// web/src/lib/guides/leveling-talents.ts
// A leveling band's talents as the list a player works through: each talent's icon, name and
// rank, in the order to take them (the build's main tree first, each tree from its first row
// down). A talent whose rank rose since the band before is marked, so the strip also shows
// what each band adds. Pure over the decoded trees `decodeBandTalents` produces.
import type { DecodedTreeView } from './build-tree';

export interface LevelingTalentChip {
  name: string;
  icon: string;
  rank: number;
  maxRank: number;
  /** True when this band spends more points here than the band before it. */
  isNew: boolean;
}

export function levelingTalentChips(
  trees: readonly DecodedTreeView[],
  previous?: readonly DecodedTreeView[],
): LevelingTalentChip[] {
  const before = new Map<number, number>();
  for (const tree of previous ?? []) for (const talent of tree.talents) before.set(talent.id, talent.rank);
  const ordered = trees
    .map((tree, index) => ({ tree, index }))
    .sort((a, b) => b.tree.points - a.tree.points || a.index - b.index);
  return ordered.flatMap(({ tree }) =>
    [...tree.talents]
      .filter((talent) => talent.rank > 0)
      .sort((a, b) => a.tier - b.tier || a.column - b.column)
      .map((talent) => ({
        name: talent.name,
        icon: talent.icon,
        rank: talent.rank,
        maxRank: talent.max_rank,
        isNew: previous !== undefined && talent.rank > (before.get(talent.id) ?? 0),
      })),
  );
}
