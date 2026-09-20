// web/src/lib/sim/weights.test.ts
import { describe, expect, it } from 'vitest';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import specsJson from '../../fixtures/sim/specs.json';
import {
  CASTER_REFERENCE,
  DEFAULT_REFERENCE,
  defaultStatsFor,
  fallbackReferenceFor,
  pawnString,
  referenceFor,
  statLabel,
  weightScale,
  WEIGHT_STATS,
} from './weights';
import type { WeightsResult } from './bulk-types';
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
});

describe('statLabel', () => {
  it('names every stat the picker offers', () => {
    for (const stat of WEIGHT_STATS) expect(statLabel(stat.id)).toBe(stat.label);
    expect(statLabel('nonesuch')).toBe('nonesuch');
  });
});

describe('weightScale', () => {
  it('is the largest weight plus its error, so no bar overflows its track', () => {
    expect(weightScale(result.weights)).toBeCloseTo(27.3 + 1.2, 6);
    expect(weightScale([])).toBe(1);
  });
});

describe('pawnString', () => {
  it('is a Pawn v1 line with the class, the spec and two decimals a piece', () => {
    expect(pawnString('warrior-fury', result.weights)).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury, AttackPower=1.00, Strength=2.14, ' +
        'CritRating=21.70, HitRating=27.30, Agility=1.32, HasteRating=18.40 )',
    );
  });

  it('leaves out a stat Pawn has no key for rather than inventing one', () => {
    expect(pawnString('warrior-fury', [{ stat: 'nonesuch', weight: 3, error: 0 }])).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury )',
    );
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
