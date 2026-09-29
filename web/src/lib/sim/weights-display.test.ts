// web/src/lib/sim/weights-display.test.ts
import { describe, expect, it } from 'vitest';
import { buildWeightDisplayRows, refAbbrev, statLabelForSpec } from './weights-display';
import type { StatWeight } from './types';

// The owner's own repro (2026-09-28): a level-20 hunter-marksmanship weights run --
// attack_power dropped from the spec's weight_stats by this same lane (data/curated/
// specs.json), so it never reaches this module at all. `insignificant` here is exactly
// what the JSON this lane's Go change writes carries (sim/cmd/leveling-bis/report.go's
// isWeightSignificant, a 25%-of-value error bar): this module never recomputes that ratio
// itself -- the wire's own flag, read through weights.ts's `isSignificant`, is the single
// source of truth, the same way every other significance-aware component on the site
// already reads it.
const hunterWeights: StatWeight[] = [
  { stat: 'ranged_attack_power', weight: 1, error: 0 },
  { stat: 'agility', weight: 2.05, error: 0.1 },
  { stat: 'crit', weight: 8.41, error: 0.3 },
  { stat: 'hit', weight: 5.23, error: 0.2 },
  { stat: 'melee_haste', weight: 14.87, error: 6.0, insignificant: true },
];

describe('statLabelForSpec', () => {
  it('renders melee_haste as "Attack speed" for a hunter, since the engine has no separate ranged-haste stat and MeleeHaste is what speeds a hunter’s shots', () => {
    expect(statLabelForSpec('melee_haste', 'hunter-marksmanship')).toBe('Attack speed');
  });

  it('keeps the plain "Melee haste" label for a class the swing-speed override does not name', () => {
    expect(statLabelForSpec('melee_haste', 'warrior-fury')).toBe('Melee haste');
  });

  it('falls through to the ordinary vocabulary for every other id, hunter or not', () => {
    expect(statLabelForSpec('agility', 'hunter-marksmanship')).toBe('Agility');
    expect(statLabelForSpec('crit', 'hunter-marksmanship')).toBe('Crit');
  });
});

describe('buildWeightDisplayRows', () => {
  const rows = buildWeightDisplayRows(hunterWeights, 'ranged_attack_power', 'hunter-marksmanship');

  it('never renders a raw snake_case id', () => {
    for (const row of rows) {
      expect(row.label).not.toMatch(/_/);
      expect(row.sentence).not.toMatch(/[a-z]_[a-z]/);
    }
  });

  it('puts the reference stat first regardless of its own weight', () => {
    expect(rows[0].stat).toBe('ranged_attack_power');
    expect(rows[0].isReference).toBe(true);
  });

  it('sorts every other row by weight, highest first', () => {
    const nonReference = rows.filter((row) => !row.isReference).map((row) => row.stat);
    expect(nonReference).toEqual(['melee_haste', 'crit', 'hit', 'agility']);
  });

  it('labels melee_haste "Attack speed" on this hunter row, not "Melee haste"', () => {
    const meleeHaste = rows.find((row) => row.stat === 'melee_haste');
    expect(meleeHaste?.label).toBe('Attack speed');
  });

  it('renders a significant weight as "1 <stat> = <weight> <reference>"', () => {
    const agility = rows.find((row) => row.stat === 'agility');
    expect(agility?.significant).toBe(true);
    expect(agility?.sentence).toBe('1 Agility = 2.05 Ranged attack power');
  });

  it('flags a weight whose error is not under 25% of its own value as not significant, and says so rather than printing a bare number', () => {
    const attackSpeed = rows.find((row) => row.stat === 'melee_haste');
    expect(attackSpeed?.significant).toBe(false);
    expect(attackSpeed?.sentence).toContain('not significant');
    expect(attackSpeed?.sentence).toContain('±');
  });

  it('never drops an insignificant row (tenet 4: nothing hidden when it is the point)', () => {
    expect(rows).toHaveLength(hunterWeights.length);
  });

  it('gives the reference row its own sentence rather than "1 Ranged attack power = 1.00 Ranged attack power"', () => {
    expect(rows[0].sentence).not.toMatch(/^1 /);
    expect(rows[0].sentence.toLowerCase()).toContain('reference');
  });

  it('an unknown spec falls back to the plain vocabulary rather than throwing', () => {
    const fallback = buildWeightDisplayRows(hunterWeights, 'ranged_attack_power', 'nobody-heard-of-this');
    const meleeHaste = fallback.find((row) => row.stat === 'melee_haste');
    expect(meleeHaste?.label).toBe('Melee haste');
  });
});

describe('refAbbrev', () => {
  it('takes one initial per word 3+ letters long for a multi-word label', () => {
    expect(refAbbrev('Attack power')).toBe('AP');
    expect(refAbbrev('Ranged attack power')).toBe('RAP');
    expect(refAbbrev('Spell power')).toBe('SP');
  });

  it('keeps a one-word label’s own first two letters, capitalised, rather than a single initial', () => {
    expect(refAbbrev('Strength')).toBe('ST');
    expect(refAbbrev('Agility')).toBe('AG');
  });
});
