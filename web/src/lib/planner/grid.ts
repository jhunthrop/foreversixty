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

/**
 * The cell arrow keys should move to. Horizontal movement walks the flat list, so the end of
 * a tier continues into the next one; vertical movement jumps to the nearest column in the
 * next occupied tier. Both stop at the edges rather than wrapping.
 */
export function moveFocus(cells: GridCell[], current: GridCell, dx: number, dy: number): GridCell {
  if (cells.length === 0) return current;

  if (dx !== 0) {
    const at = cells.indexOf(current);
    const next = Math.min(Math.max(at + dx, 0), cells.length - 1);
    return cells[next];
  }

  if (dy !== 0) {
    const tiers = [...new Set(cells.map((cell) => cell.tier))].sort((a, b) => a - b);
    const candidates =
      dy > 0
        ? tiers.filter((tier) => tier > current.tier)
        : tiers.filter((tier) => tier < current.tier).reverse();
    for (const tier of candidates) {
      const row = cells.filter((cell) => cell.tier === tier);
      if (row.length === 0) continue;
      return row.reduce((best, cell) =>
        Math.abs(cell.column - current.column) < Math.abs(best.column - current.column) ? cell : best,
      );
    }
  }

  return current;
}
