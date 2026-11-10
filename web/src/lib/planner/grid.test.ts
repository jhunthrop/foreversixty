// web/src/lib/planner/grid.test.ts
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { gridCells, gridSize } from './grid';
import type { TalentFile } from './types';

const arms = (fixtureTalents as TalentFile).trees[0];
const cells = gridCells(arms);

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
