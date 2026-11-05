// web/src/lib/sim/buff-effect.test.ts
// The "No effect on your damage" rule (spec 2026-11-04-sim-results 3.12, build ruling R2), on the
// published entries' shape and the Fury weights.
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import { damageSchools, hasNoDamageEffect, type BuffEffectRow } from './buff-effect';
import type { SimResult } from './types';

const FURY = { attack_power: 1, strength: 1.9, agility: 0.7, hit: 32, crit: 18 };
const MELEE = ['physical'];
const row = (fields: Partial<BuffEffectRow>): BuffEffectRow => ({ target: false, ...fields });
const none = (
  fields: Partial<BuffEffectRow>,
  schools = MELEE,
  weights: Record<string, number> = FURY,
): boolean => hasNoDamageEffect(row(fields), weights, schools);

describe('hasNoDamageEffect', () => {
  it('files what a Fury warrior gains nothing from', () => {
    expect(none({ stats: { intellect: 31 } })).toBe(true);
    expect(none({ stats: { mp5: 25 } })).toBe(true);
    expect(none({ stats: { ranged_attack_power: 110 } })).toBe(true);
    expect(none({ stats: { spell_crit: 3 } })).toBe(true);
    expect(none({})).toBe(true);
  });

  it('keeps what changes a swing', () => {
    expect(none({ stats: { attack_power: 112 } })).toBe(false);
    expect(none({ stats: { target_armor: -505 } })).toBe(false);
    expect(none({ melee_effect: true })).toBe(false);
  });

  it('reads a school debuff against the schools the run dealt damage in', () => {
    expect(none({ school_damage: ['fire', 'frost'] })).toBe(true);
    expect(none({ school_damage: ['fire', 'frost'] }, ['fire'])).toBe(false);
  });

  it('armor means nothing to a caster, spell power everything', () => {
    expect(none({ stats: { target_armor: -505 } }, ['fire'], { spell_power: 1 })).toBe(true);
    expect(none({ stats: { spell_crit: 3 } }, ['fire'], { spell_power: 1 })).toBe(false);
  });

  it('claims nothing from a file with no effect fields, or an id it does not carry', () => {
    expect(hasNoDamageEffect({ stats: { intellect: 31 } }, FURY, MELEE)).toBe(false);
    expect(hasNoDamageEffect(undefined, FURY, MELEE)).toBe(false);
  });
});

describe('damageSchools', () => {
  it('reads the Fury fixture as physical only', () => {
    expect(damageSchools((full as unknown as SimResult).summary.damage_done)).toEqual(['physical']);
  });
});
