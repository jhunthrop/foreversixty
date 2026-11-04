// web/src/lib/guides/leveling-talents.ts
// A leveling band's talents as the list a player works through, in the order the points are
// spent: one order for every band (the ladder's own point order, `guideLadderOrder` in
// leveling-builds.ts), each talent placed by its first point, so a talent
// never moves between bands. A talent whose rank rose since the band before is marked, so the
// strip also shows what each band adds. Pure over the decoded trees `decodeBandTalents` produces.
import type { DecodedTreeView } from './build-tree';

export interface LevelingTalentChip {
  id: number;
  name: string;
  icon: string;
  rank: number;
  maxRank: number;
  /** True when this band spends more points here than the band before it. */
  isNew: boolean;
}

/** One chip list per band, bands in level order; `undefined` (a band whose talents did not
 *  decode) gives an empty list. `spendOrder` is one talent id per point; a lit talent it does
 *  not name follows the ones it does, by tree position, row and column. */
export function levelingChipRows(
  bands: readonly (readonly DecodedTreeView[] | undefined)[],
  spendOrder: readonly number[],
): LevelingTalentChip[][] {
  const firstPoint = new Map<number, number>();
  spendOrder.forEach((id, point) => {
    if (!firstPoint.has(id)) firstPoint.set(id, point);
  });
  const ranks = new Map<number, number>();
  return bands.map((trees, bandIndex) => {
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
        isNew: bandIndex > 0 && talent.rank > (ranks.get(talent.id) ?? 0),
      }));
    for (const talent of lit) ranks.set(talent.id, talent.rank);
    return chips;
  });
}
