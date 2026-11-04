import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import ratingFixture from '../../../../docs/contracts/fixtures/rating-per-percent.json';
import activeBuild from '../../data/active-build.json';
import { RATING_PER_PERCENT, scoreItem, specKeyFor, type SpecWeights } from './score';
import type { Item } from '../planner/types';

const item = (stats: Record<string, number>, armor = 0, percent_stats: Record<string, number> = {}): Item =>
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
    percent_stats,
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

// Every weight is per 1% (stat-weights.json, a band's weight_per_percent); an item's line is
// rating points (sim audit round 17: the score multiplied rating by a per-1% weight).
describe('scoreItem with rating stats', () => {
  // data/builds/1.60.1.70755/stat-weights.json warrior-fury.
  const fury: SpecWeights = {
    attack_power: 1,
    strength: 1.9,
    agility: 0.7,
    hit: 32,
    crit: 18,
    expertise: 18,
  };
  // data/builds/1.60.1.70755/items/warrior.json, hands.
  const annihilation = item({ strength: 35, stamina: 15, crit: 14, hit: 10 }, 615);
  const fists = item({ stamina: 26, strength: 26, crit: 1, hit: 1 }, 629, { crit: 1, hit: 1 });
  const enforcement = item({ agility: 20, strength: 28, stamina: 6, hit: 10 }, 140);
  const batskin = item({ agility: 10, stamina: 9, hit: 20 }, 113);

  it('reads hit and crit rating as percent before the per-1% weight', () => {
    // 35 * 1.9 + 14 / 14 * 18 + 10 / 10 * 32
    expect(scoreItem(annihilation, fury)).toBeCloseTo(116.5);
    // 10 * 0.7 + 20 / 10 * 32: was 647
    expect(scoreItem(batskin, fury)).toBeCloseTo(71);
  });

  it('keeps an on-equip literal percent as percent', () => {
    // 26 * 1.9 + 1 * 18 + 1 * 32
    expect(scoreItem(fists, fury)).toBeCloseTo(99.4);
  });

  it('orders the Fury hands Annihilation, Fists, Enforcement, with Batskin below them', () => {
    const named = { annihilation, fists, enforcement, batskin };
    const order = Object.entries(named)
      .sort(([, a], [, b]) => scoreItem(b, fury) - scoreItem(a, fury))
      .map(([name]) => name);
    expect(order).toEqual(['annihilation', 'fists', 'enforcement', 'batskin']);
  });

  it('reads expertise rating at 10 per 1%', () => {
    // Waywatcher Gauntlets 240081: 20 expertise = 2%.
    expect(scoreItem(item({ expertise: 20 }), fury)).toBeCloseTo(36);
  });

  it('reads dodge, parry and block at their own factors and defense one for one', () => {
    const tank: SpecWeights = { dodge: 12, parry: 15, block: 5, defense: 2 };
    expect(scoreItem(item({ dodge: 12, parry: 15, block: 5, defense: 3 }), tank)).toBeCloseTo(
      12 + 15 + 5 + 6,
    );
  });
});

describe('RATING_PER_PERCENT', () => {
  it('is the shared contract fixture', () => {
    expect(RATING_PER_PERCENT).toEqual(ratingFixture.rating_per_percent);
  });

  it("matches every level of the active build's combatratings.txt", () => {
    const path = new URL(
      `../../../../data/builds/${activeBuild.build}/gametables/combatratings.txt`,
      import.meta.url,
    );
    const [header, ...rows] = readFileSync(path, 'utf8')
      .split(/\r?\n/)
      .filter((line) => line !== '')
      .map((line) => line.split('\t'));
    const columns: Record<string, string[]> = {
      hit: ['Hit - Melee', 'Hit - Ranged', 'Hit - Spell'],
      crit: ['Crit - Melee', 'Crit - Ranged', 'Crit - Spell'],
      expertise: ['Expertise'],
      dodge: ['Dodge'],
      parry: ['Parry'],
      block: ['Block'],
      defense: ['Defense Skill'],
      haste: ['Haste - Melee', 'Haste - Ranged', 'Haste - Spell'],
    };
    expect(Object.keys(columns).sort()).toEqual(Object.keys(ratingFixture.rating_per_percent).sort());
    expect(rows.length).toBeGreaterThanOrEqual(60);
    for (const [stat, names] of Object.entries(columns)) {
      for (const name of names) {
        const index = header!.indexOf(name);
        expect(index, name).toBeGreaterThan(0);
        for (const row of rows) {
          expect(Number(row[index]), `${name} at level ${row[0]}`).toBe(
            ratingFixture.rating_per_percent[stat as keyof typeof ratingFixture.rating_per_percent],
          );
        }
      }
    }
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
