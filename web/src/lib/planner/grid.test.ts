// web/src/lib/planner/grid.test.ts
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { gridCells, gridSize, moveFocus } from './grid';
import type { TalentFile } from './types';

const arms = (fixtureTalents as TalentFile).trees[0];
const cells = gridCells(arms);
const cellAt = (tier: number, column: number) => cells.find((c) => c.tier === tier && c.column === column)!;

describe('gridCells', () => {
  it('sorts by tier then column', () => {
    expect(cells.map((c) => [c.tier, c.column])).toEqual([
      [0, 0],
      [0, 1],
      [1, 0],
      [1, 1],
      [2, 0],
      [2, 1],
      [3, 0],
    ]);
  });
});

describe('gridSize', () => {
  it('reports the number of tiers and columns the grid needs', () => {
    expect(gridSize(arms)).toEqual({ tiers: 4, columns: 2 });
  });
});

describe('moveFocus', () => {
  it('walks the flat list left and right, crossing tier boundaries', () => {
    expect(moveFocus(cells, cellAt(0, 0), 1, 0)).toBe(cellAt(0, 1));
    expect(moveFocus(cells, cellAt(0, 1), 1, 0)).toBe(cellAt(1, 0));
    expect(moveFocus(cells, cellAt(1, 0), -1, 0)).toBe(cellAt(0, 1));
  });

  it('stops at both ends rather than wrapping around', () => {
    expect(moveFocus(cells, cellAt(0, 0), -1, 0)).toBe(cellAt(0, 0));
    expect(moveFocus(cells, cellAt(3, 0), 1, 0)).toBe(cellAt(3, 0));
  });

  it('moves down and up to the nearest column in the next occupied tier', () => {
    expect(moveFocus(cells, cellAt(0, 1), 0, 1)).toBe(cellAt(1, 1));
    expect(moveFocus(cells, cellAt(2, 1), 0, 1)).toBe(cellAt(3, 0));
    expect(moveFocus(cells, cellAt(3, 0), 0, -1)).toBe(cellAt(2, 0));
  });

  it('stays put when there is no tier in that direction', () => {
    expect(moveFocus(cells, cellAt(0, 0), 0, -1)).toBe(cellAt(0, 0));
    expect(moveFocus(cells, cellAt(3, 0), 0, 1)).toBe(cellAt(3, 0));
  });
});
