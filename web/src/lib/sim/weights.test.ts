// web/src/lib/sim/weights.test.ts
import { describe, expect, it } from 'vitest';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import specsJson from '../../fixtures/sim/specs.json';
import {
  CASTER_REFERENCE,
  DEFAULT_REFERENCE,
  defaultStatsFor,
  fallbackReferenceFor,
  isDpsSpec,
  isSignificant,
  pawnString,
  pickableStatsFor,
  referenceFor,
  statLabel,
  weightScale,
  weightStatsFor,
  WEIGHT_STATS,
} from './weights';
import type { StatWeight, WeightsResult } from './bulk-types';
import type { SpecFidelity } from './types';

const result = weightsResultJson as unknown as WeightsResult;
const specs = specsJson as unknown as SpecFidelity[];

describe('referenceFor', () => {
  it('takes the spec’s own reference stat from GET /v1/specs', () => {
    expect(referenceFor('warrior-fury', specs)).toBe('attack_power');
    expect(referenceFor('mage-frost', specs)).toBe('spell_power');
  });

  it('falls back to the pinned default for the kind of spec, not to one number', () => {
    // Contract 10.8: attack_power for melee and hunters, spell_power for casters.
    expect(referenceFor('warrior-fury', [])).toBe(DEFAULT_REFERENCE);
    expect(referenceFor('mage-fire', [])).toBe(CASTER_REFERENCE);
    expect(referenceFor('druid-balance', specs)).toBe(CASTER_REFERENCE);
    expect(fallbackReferenceFor('hunter-marksmanship')).toBe(DEFAULT_REFERENCE);
    expect(fallbackReferenceFor('nonesuch-spec')).toBe(DEFAULT_REFERENCE);
  });
});

describe('defaultStatsFor', () => {
  it('always includes the reference stat, exactly once, first', () => {
    const stats = defaultStatsFor('warrior-fury', 'attack_power');
    expect(stats[0]).toBe('attack_power');
    expect(stats.filter((stat) => stat === 'attack_power')).toHaveLength(1);
    expect(new Set(stats).size).toBe(stats.length);
  });

  it('is scoped to weightStats when the spec sent one, the reference stat still first', () => {
    const stats = defaultStatsFor('warrior-fury', 'attack_power', ['attack_power', 'strength', 'agility']);
    expect(stats).toEqual(['attack_power', 'strength', 'agility']);
  });
});

describe('statLabel', () => {
  it('names every stat the picker offers', () => {
    for (const stat of WEIGHT_STATS) expect(statLabel(stat.id)).toBe(stat.label);
    expect(statLabel('nonesuch')).toBe('nonesuch');
  });
});

describe('weightScale', () => {
  it('is the largest weight plus its error, so no bar overflows its track', () => {
    // The widest row computed from the fixture itself, not a copied literal: a regenerated
    // fixture's own numbers change every time the engine reruns it, and a hand-copied
    // number here would silently drift from what the fixture actually carries.
    const widest = Math.max(...result.weights.map((row) => row.weight + row.error));
    expect(weightScale(result.weights)).toBeCloseTo(widest, 6);
    expect(weightScale([])).toBe(1);
  });
});

describe('pawnString', () => {
  it('is a Pawn v1 line with the class, the spec and two decimals a piece', () => {
    // The fixture's own "agility" row is insignificant -- left out here the same way a
    // stat Pawn has no key for is, below.
    expect(pawnString('warrior-fury', result.weights)).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury, Strength=1.50, AttackPower=1.00, ' +
        'CritRating=7.03, HitRating=6.87, HasteRating=5.85 )',
    );
  });

  it('leaves out a stat Pawn has no key for rather than inventing one', () => {
    expect(pawnString('warrior-fury', [{ stat: 'nonesuch', weight: 3, error: 0 }])).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury )',
    );
  });

  it('D45: leaves out a weight the engine flagged insignificant', () => {
    const weights: StatWeight[] = [
      { stat: 'attack_power', weight: 1, error: 0 },
      { stat: 'crit', weight: 14.37, error: 22.55, insignificant: true },
    ];
    expect(pawnString('warrior-fury', weights)).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury, AttackPower=1.00 )',
    );
  });

  it('D46: never disagrees with isSignificant about which rows are included -- the regression guard', () => {
    // The table greys a row exactly when `isSignificant` is false (StatWeights.svelte,
    // SavedWeights.svelte); pawnString must include exactly the rows that predicate calls
    // significant, and no others, for every row that has a Pawn key at all. Asserted
    // generically -- over the whole pinned vocabulary, not one hand-picked stat -- so this
    // guard cannot be satisfied by coincidence the way a single fixed-string test could be.
    const weights: StatWeight[] = WEIGHT_STATS.map((stat, index) => ({
      stat: stat.id,
      weight: index + 1,
      error: 0.1,
      insignificant: index % 2 === 0,
    }));
    const pawn = pawnString('warrior-fury', weights);
    for (const row of weights) {
      const key = WEIGHT_STATS.find((stat) => stat.id === row.stat)?.pawn ?? '';
      if (key === '') continue; // Pawn has no key for this stat regardless of significance.
      const included = pawn.includes(`${key}=`);
      expect(included, `${row.stat} (insignificant: ${row.insignificant})`).toBe(isSignificant(row));
    }
  });
});

describe('isSignificant', () => {
  it('treats absent or false insignificant as significant', () => {
    expect(isSignificant({ insignificant: undefined })).toBe(true);
    expect(isSignificant({ insignificant: false })).toBe(true);
    expect(isSignificant({})).toBe(true);
  });

  it('treats insignificant: true as not significant', () => {
    expect(isSignificant({ insignificant: true })).toBe(false);
  });
});

describe('pickableStatsFor (sub-item 4: only the spec’s own stats)', () => {
  it('offers the full pinned vocabulary when the engine sent no weight_stats', () => {
    expect(pickableStatsFor(undefined)).toBe(WEIGHT_STATS);
    expect(pickableStatsFor([])).toBe(WEIGHT_STATS);
  });

  it('offers only the spec’s own stats, in the engine’s order, when it sent one', () => {
    const pickable = pickableStatsFor(['attack_power', 'strength']);
    expect(pickable.map((stat) => stat.id)).toEqual(['attack_power', 'strength']);
    // Retail-only entries (D45) are never offered once the engine names its own list --
    // by construction (they are simply not in the list), not by a deny-list here.
    for (const wrong of ['expertise', 'spell_haste', 'armor_penetration', 'mp5', 'feral_attack_power']) {
      expect(pickable.map((stat) => stat.id)).not.toContain(wrong);
    }
  });
});

describe('weightStatsFor', () => {
  it('reads the spec’s own weight_stats column from GET /v1/specs', () => {
    expect(weightStatsFor('warrior-fury', specs)).toEqual([
      'attack_power',
      'strength',
      'agility',
      'crit',
      'hit',
      'melee_haste',
      'expertise',
      'armor_penetration',
    ]);
  });

  it('is undefined for a spec with no row, or a row predating the column', () => {
    expect(weightStatsFor('mage-frost', specs)).toBeUndefined();
    expect(weightStatsFor('nonesuch-spec', specs)).toBeUndefined();
  });
});

describe('isDpsSpec (sub-item 5: the honest refusal)', () => {
  it('is true for a dps spec', () => {
    expect(isDpsSpec('warrior-fury')).toBe(true);
  });

  it('is false for a healer or tank spec, and for an unrecognised one', () => {
    expect(isDpsSpec('druid-restoration')).toBe(false);
    expect(isDpsSpec('nonesuch-spec')).toBe(false);
  });
});

describe('the stat vocabulary (contract 10.8, pinned)', () => {
  it('carries one hit and one crit, and splits only haste', () => {
    const ids = WEIGHT_STATS.map((stat) => stat.id);
    expect(ids).toContain('hit');
    expect(ids).toContain('crit');
    expect(ids).toContain('melee_haste');
    expect(ids).toContain('spell_haste');
    for (const wrong of ['melee_hit', 'spell_hit', 'melee_crit', 'spell_crit', 'haste']) {
      expect(ids).not.toContain(wrong);
    }
  });

  it('spells MP5 as mp5', () => {
    const ids = WEIGHT_STATS.map((stat) => stat.id);
    expect(ids).toContain('mp5');
    expect(ids).not.toContain('m_p5');
  });

  it('offers only ids the pinned proto.Stat list carries', () => {
    const pinned = new Set([
      'strength',
      'agility',
      'stamina',
      'intellect',
      'spirit',
      'spell_power',
      'arcane_power',
      'fire_power',
      'frost_power',
      'holy_power',
      'nature_power',
      'shadow_power',
      'mp5',
      'hit',
      'crit',
      'spell_haste',
      'spell_penetration',
      'attack_power',
      'melee_haste',
      'armor_penetration',
      'expertise',
      'mana',
      'energy',
      'rage',
      'armor',
      'ranged_attack_power',
      'defense',
      'block',
      'block_value',
      'dodge',
      'parry',
      'health',
      'arcane_resistance',
      'fire_resistance',
      'frost_resistance',
      'nature_resistance',
      'shadow_resistance',
      'bonus_armor',
      'healing_power',
      'spell_damage',
      'feral_attack_power',
    ]);
    for (const stat of WEIGHT_STATS) expect(pinned.has(stat.id), stat.id).toBe(true);
  });
});
