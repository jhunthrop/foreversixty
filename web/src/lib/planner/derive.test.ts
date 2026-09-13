// web/src/lib/planner/derive.test.ts
import { describe, expect, it } from 'vitest';
import fixtureItems from '../../fixtures/planner/items/warrior.json';
import fixtureSets from '../../fixtures/planner/sets.json';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import {
  activeSetBonuses,
  equippedItems,
  levelForIndex,
  levelReached,
  pointsPerTree,
  ranksByTalent,
  splitLabel,
  statTotals,
} from './derive';
import { indexItems, indexTalents } from './rules';
import type { ItemFile, ItemSet, TalentFile } from './types';

const index = indexTalents(fixtureTalents as TalentFile);
const items = indexItems(fixtureItems as ItemFile);
const sets = fixtureSets as ItemSet[];
const repeat = (id: number, n: number): number[] => Array.from({ length: n }, () => id);

describe('levels', () => {
  it('spends the first point at level 10 and counts up by one', () => {
    expect(levelForIndex(0)).toBe(10);
    expect(levelForIndex(41)).toBe(51);
  });

  it('reports level 9 before any point and the last point level after', () => {
    expect(levelReached([])).toBe(9);
    expect(levelReached([1001])).toBe(10);
    expect(levelReached(repeat(1001, 3))).toBe(12);
  });
});

describe('ranksByTalent', () => {
  it('counts the points in each talent', () => {
    expect([...ranksByTalent([1001, 1002, 1001]).entries()].sort()).toEqual([
      [1001, 2],
      [1002, 1],
    ]);
    expect(ranksByTalent([]).size).toBe(0);
  });
});

describe('the split', () => {
  it('reports one number per tree in left-to-right order', () => {
    const order = repeat(1001, 3).concat(repeat(2001, 2));
    expect(pointsPerTree(index, order)).toEqual([3, 2]);
    expect(splitLabel(index, order)).toBe('3/2');
  });

  it('is all zeroes for an empty build', () => {
    expect(pointsPerTree(index, [])).toEqual([0, 0]);
    expect(splitLabel(index, [])).toBe('0/0');
  });
});

describe('gear totals', () => {
  it('resolves the equipped items in slot order and skips empty slots', () => {
    expect(equippedItems(items, { head: 12640, main_hand: 12784 }).map((i) => i.id)).toEqual([12640, 12784]);
    expect(equippedItems(items, {})).toEqual([]);
    expect(equippedItems(items, { head: 999999 })).toEqual([]);
  });

  it('sums stats and folds armor into the armor key', () => {
    const equipped = equippedItems(items, { head: 12640, shoulder: 16966 });
    expect(statTotals(equipped)).toEqual({
      strength: 38,
      crit: 2,
      hit: 2,
      stamina: 18,
      armor: 1065,
    });
  });

  it('returns an empty total for no gear', () => {
    expect(statTotals([])).toEqual({});
  });
});

describe('set bonuses', () => {
  it('activates a bonus once enough pieces are equipped', () => {
    const two = equippedItems(items, { head: 16963, shoulder: 16966 });
    expect(activeSetBonuses(two, sets)).toEqual([{ set: sets[0], pieces: 2, active: sets[0].bonuses }]);
  });

  it('lists the set with no active bonus when only one piece is worn', () => {
    const one = equippedItems(items, { head: 16963 });
    expect(activeSetBonuses(one, sets)).toEqual([{ set: sets[0], pieces: 1, active: [] }]);
  });

  it('lists nothing when no set piece is worn', () => {
    expect(activeSetBonuses(equippedItems(items, { head: 12640 }), sets)).toEqual([]);
  });
});
