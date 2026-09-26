// web/src/lib/planner/styles.test.ts
import { describe, expect, it } from 'vitest';
import { CELL_BORDER, CELL_PILL, cellState, treeRowColumnsClass } from './styles';

describe('cellState', () => {
  it('is locked when no point can go in yet', () => {
    expect(cellState(0, 5, false)).toBe('locked');
  });

  it('is available once a point could go in', () => {
    expect(cellState(0, 5, true)).toBe('available');
  });

  it('is filled while it holds points below its cap', () => {
    expect(cellState(1, 5, true)).toBe('filled');
    expect(cellState(4, 5, true)).toBe('filled');
  });

  it('is maxed at the cap, whether or not anything else could be added', () => {
    expect(cellState(5, 5, false)).toBe('maxed');
    expect(cellState(1, 1, false)).toBe('maxed');
  });

  it('has a border and a pill class for every state', () => {
    for (const state of ['locked', 'available', 'filled', 'maxed'] as const) {
      expect(CELL_BORDER[state]).toBeTruthy();
      expect(CELL_PILL[state]).toBeTruthy();
    }
  });
});

describe('treeRowColumnsClass', () => {
  it('gives each tree its own column for one, two or three trees', () => {
    expect(treeRowColumnsClass(1)).toBe('md:grid-cols-1');
    expect(treeRowColumnsClass(2)).toBe('md:grid-cols-2');
    expect(treeRowColumnsClass(3)).toBe('md:grid-cols-3');
  });

  it('falls back to three columns for a count the table has never seen', () => {
    expect(treeRowColumnsClass(4)).toBe('md:grid-cols-3');
    expect(treeRowColumnsClass(0)).toBe('md:grid-cols-3');
  });
});
