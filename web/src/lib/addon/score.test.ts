import { describe, expect, it } from 'vitest';
import { scoreItem, specKeyFor, type SpecWeights } from './score';
import type { Item } from '../planner/types';

const item = (stats: Record<string, number>, armor = 0): Item =>
  ({
    id: 1,
    name: 'x',
    icon: 'i',
    slot: 'head',
    quality: 4,
    required_level: 0,
    item_level: 60,
    armor,
    stats,
    set_id: null,
    unique: false,
  }) as unknown as Item;

const weights: SpecWeights = { spell_power: 1, intellect: 0.45, armor: 0.12, healing_power: 1 };

describe('scoreItem', () => {
  it('sums weight times stat', () => {
    expect(scoreItem(item({ spell_power: 23, intellect: 10 }), weights)).toBeCloseTo(27.5);
  });

  it('counts armour, which is a column and not a stat on the record', () => {
    expect(scoreItem(item({}, 400), weights)).toBeCloseTo(48);
  });

  it('maps the planner stat names onto the contract vocabulary', () => {
    expect(scoreItem(item({ healing: 40 }), weights)).toBeCloseTo(40);
  });

  it('scores an unweighted stat as zero rather than refusing the item', () => {
    expect(scoreItem(item({ dodge: 9 }), weights)).toBe(0);
  });
});

describe('specKeyFor', () => {
  it('takes the tree with the most points', () => {
    expect(specKeyFor('paladin', [2, 31, 0])).toBe('paladin-protection');
  });

  it('resolves a tie to the first tree', () => {
    expect(specKeyFor('paladin', [10, 10, 0])).toBe('paladin-holy');
  });

  it('returns the first tree for an unspent build', () => {
    expect(specKeyFor('paladin', [0, 0, 0])).toBe('paladin-holy');
  });
});
