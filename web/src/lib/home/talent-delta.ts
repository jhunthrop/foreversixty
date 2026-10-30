// web/src/lib/home/talent-delta.ts
// Home rebuild spec §3.B.2's Talents card: how far a character's actual talent split has
// drifted from a leveling-BiS band's own recommended build.

/** One talent cell's share of the count: the band's points there the character has not matched. */
export function unmatchedBandPoints(wornRank: number, bandRank: number): number {
  return Math.max(bandRank - wornRank, 0);
}

/**
 * `trees` (`MeCharacter.build.talents.trees`: three per-tree rank strings in tree order,
 * one digit per talent cell, e.g. `"503200000"`) against `bandTalents` (a BiS band's own
 * `talents` field: the identical per-tree digit convention, the three trees joined with
 * `"-"`, e.g. `"0000000000000000-35300000000000000-000000000000000000"`).
 *
 * The result is how many of the band build's points the character has not matched:
 * `Σ max(band[i] - worn[i], 0)` over every talent cell. It is the numerator of "N of the
 * band's M points differ", so it can never exceed M (each cell contributes at most the
 * band's own rank there). Summing `|worn - band|` instead would also count every point the
 * character spent where the band spends none, which doubles a fully different build to 2M
 * ("14 of the 11 points differ"). A character who spent the same number of points as the
 * band elsewhere is still told how many of the band's points they are missing, which is the
 * figure the card's denominator names. Missing digits (a shorter string than its
 * counterpart) read as `0`, the same "nothing spent there" default a talent string's own
 * trailing zeros already carry.
 */
export function talentDeltaFor(trees: readonly string[], bandTalents: string): number {
  const bandTrees = bandTalents.split('-');
  const treeCount = Math.max(trees.length, bandTrees.length);
  let unmatched = 0;
  for (let treeIndex = 0; treeIndex < treeCount; treeIndex += 1) {
    const worn = trees[treeIndex] ?? '';
    const band = bandTrees[treeIndex] ?? '';
    const cellCount = Math.max(worn.length, band.length);
    for (let cell = 0; cell < cellCount; cell += 1) {
      const wornRank = Number.parseInt(worn[cell] ?? '0', 10) || 0;
      const bandRank = Number.parseInt(band[cell] ?? '0', 10) || 0;
      unmatched += unmatchedBandPoints(wornRank, bandRank);
    }
  }
  return unmatched;
}
