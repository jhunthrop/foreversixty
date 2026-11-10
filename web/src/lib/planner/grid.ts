// web/src/lib/planner/grid.ts
// Talent trees are sparse: a tier may have one talent or four, and columns skip. Keyboard
// navigation therefore cannot be index arithmetic on a dense grid, so it lives here as a
// pure function over the occupied cells and is unit-tested without a browser.
import type { Talent, TalentTree } from './types';

export interface GridCell {
  tier: number;
  column: number;
  talent: Talent;
}

/** Occupied cells, sorted top-to-bottom then left-to-right. */
export function gridCells(tree: TalentTree): GridCell[] {
  return tree.talents
    .map((talent) => ({ tier: talent.tier, column: talent.column, talent }))
    .sort((a, b) => a.tier - b.tier || a.column - b.column);
}

/** How many rows and columns the CSS grid needs to hold the tree. */
export function gridSize(tree: TalentTree): { tiers: number; columns: number } {
  const tiers = tree.talents.reduce((max, t) => Math.max(max, t.tier + 1), 0);
  const columns = tree.talents.reduce((max, t) => Math.max(max, t.column + 1), 0);
  return { tiers, columns };
}
