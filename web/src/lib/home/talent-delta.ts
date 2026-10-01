// web/src/lib/home/talent-delta.ts
// Home rebuild spec §3.B.2's Talents card: how far a character's actual talent split has
// drifted from a leveling-BiS band's own recommended build.

/**
 * `trees` (`MeCharacter.build.talents.trees`: three per-tree rank strings in tree order,
 * one digit per talent cell, e.g. `"503200000"`) against `bandTalents` (a BiS band's own
 * `talents` field: the identical per-tree digit convention, the three trees joined with
 * `"-"`, e.g. `"0000000000000000-35300000000000000-000000000000000000"`).
 *
 * The result is the sum, over every talent cell in every tree, of the absolute difference
 * between the character's rank there and the band's: `delta = Σ |worn[i] - band[i]|`. That
 * sum is exactly "points spent where the band spends none, plus points the band spends that
 * the character does not" -- for one cell, `|w - b| = max(w - b, 0) + max(b - w, 0)`, and
 * the first term is points the character put somewhere the band's own build does not, the
 * second is points the band's build asks for that the character has not matched there;
 * summing that identity across every cell gives the total above. A character and a band
 * with the same total point count but a different allocation across cells still shows a
 * nonzero delta: spending the same number of points differently from the band is exactly
 * what this is meant to catch, not merely "fewer points than the band." Missing digits (a
 * shorter string than its counterpart) read as `0`, the same "nothing spent there" default
 * a talent string's own trailing zeros already carry.
 */
export function talentDeltaFor(trees: readonly string[], bandTalents: string): number {
  const bandTrees = bandTalents.split('-');
  const treeCount = Math.max(trees.length, bandTrees.length);
  let delta = 0;
  for (let treeIndex = 0; treeIndex < treeCount; treeIndex += 1) {
    const worn = trees[treeIndex] ?? '';
    const band = bandTrees[treeIndex] ?? '';
    const cellCount = Math.max(worn.length, band.length);
    for (let cell = 0; cell < cellCount; cell += 1) {
      const wornRank = Number.parseInt(worn[cell] ?? '0', 10) || 0;
      const bandRank = Number.parseInt(band[cell] ?? '0', 10) || 0;
      delta += Math.abs(wornRank - bandRank);
    }
  }
  return delta;
}
