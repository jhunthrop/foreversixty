// web/src/lib/guides/hit-cap-line.test.ts
import { describe, expect, it } from 'vitest';
import { guideExpertiseCapLine, guideHitCapLine } from './hit-cap-line';
import type { TalentHit } from './talent-hit';

const PRECISION_MAGE: TalentHit = { name: 'Precision', scope: 'spell', percent: 6 };
const SUPPRESSION: TalentHit = { name: 'Suppression', scope: 'hit', percent: 5 };
const PRECISION_ROGUE: TalentHit = { name: 'Precision', scope: 'hit', percent: 3 };
const SHADOW_FOCUS: TalentHit = { name: 'Shadow Focus', scope: 'school', percent: 5 };
const WEAPON_EXPERTISE: TalentHit = { name: 'Weapon Expertise', scope: 'expertise', percent: 2 };

describe('guideHitCapLine', () => {
  it('names a stat-wide talent apart from the gear: Arcane, all 6% from Precision', () => {
    expect(guideHitCapLine({ kind: 'spell', baseline: 6, spell: 10 }, [PRECISION_MAGE])?.text).toBe(
      'Hit cap: 16% · 6% from Precision, 10% still to find',
    );
  });

  it('splits the gear and buffs from the talent when both count', () => {
    expect(guideHitCapLine({ kind: 'spell', baseline: 7, spell: 9 }, [PRECISION_MAGE])?.text).toBe(
      'Hit cap: 16% · 1% hit from BiS gear and buffs, 6% from Precision, 9% still to find',
    );
    expect(guideHitCapLine({ kind: 'spell', baseline: 6, spell: 10 }, [SUPPRESSION])?.text).toBe(
      'Hit cap: 16% · 1% hit from BiS gear and buffs, 5% from Suppression, 10% still to find',
    );
  });

  it('names a school talent and the spells it covers: Shadow, 5% from Shadow Focus', () => {
    const line = guideHitCapLine(
      { kind: 'spell', baseline: 0, spell: 16, school: { names: ['Shadow'], hit: 5, to_cap: 11 } },
      [SHADOW_FOCUS],
    );
    expect(line?.text).toBe('Hit cap: 16% · 5% from Shadow Focus on Shadow spells, 11% still to find');
  });

  it('says "talents" for a school bonus its talents do not account for', () => {
    const line = guideHitCapLine({
      kind: 'spell',
      baseline: 0,
      spell: 16,
      school: { names: ['Frost'], hit: 3, to_cap: 13 },
    });
    expect(line?.text).toBe('Hit cap: 16% · 3% from talents on Frost spells, 13% still to find');
  });

  it('leaves the split off when nothing counts toward the cap yet', () => {
    expect(guideHitCapLine({ kind: 'spell', baseline: 0, spell: 16 })?.text).toBe('Hit cap: 16%');
    expect(guideHitCapLine({ baseline: 0, specials: 9, white: 28 })?.text).toBe(
      'Hit cap: 9% (special attacks), 28% (white swings)',
    );
    expect(guideHitCapLine({ baseline: 0, specials: 6.5 })?.text).toBe('Hit cap: 6.5%');
  });

  it('names no talent for a melee set whose build takes none: Fury, all gear and buffs', () => {
    expect(guideHitCapLine({ baseline: 8.9, specials: 0.1, white: 19.1 })?.text).toBe(
      'Hit cap: 9% (special attacks), 28% (white swings) · 8.9% hit from BiS gear and buffs, 0.1% still to find',
    );
  });

  it('counts only the talents a melee baseline holds', () => {
    expect(guideHitCapLine({ baseline: 6.3, specials: 2.7, white: 21.7 }, [PRECISION_ROGUE])?.text).toBe(
      'Hit cap: 9% (special attacks), 28% (white swings) · 3.3% hit from BiS gear and buffs, 3% from Precision, 2.7% still to find',
    );
    // A spell-only talent is in no melee figure.
    expect(guideHitCapLine({ baseline: 3, specials: 6 }, [PRECISION_MAGE])?.text).toBe(
      'Hit cap: 9% · 3% hit from BiS gear and buffs, 6% still to find',
    );
  });

  it('reads "reached" for a cap the character is already past', () => {
    expect(guideHitCapLine({ baseline: 10, specials: 0 })?.text).toBe(
      'Hit cap: reached · 10% hit from BiS gear and buffs',
    );
  });

  it('reads every source as one figure when the talents add up to more than the baseline', () => {
    expect(guideHitCapLine({ baseline: 2, specials: 7 }, [SUPPRESSION])?.text).toBe(
      'Hit cap: 9% · 2% hit from BiS gear, buffs and talents, 7% still to find',
    );
  });

  it('is undefined when the band publishes no cap', () => {
    expect(guideHitCapLine(null)).toBeUndefined();
    expect(guideHitCapLine(undefined)).toBeUndefined();
  });
});

describe('guideExpertiseCapLine', () => {
  it('is left out when nothing gives the character Expertise (no Forever item carries it)', () => {
    expect(guideExpertiseCapLine({ baseline: 0, dodge: 6.5 })).toBeUndefined();
    expect(guideExpertiseCapLine({ baseline: 0, dodge: 5.5, parry: 14 })).toBeUndefined();
    expect(guideExpertiseCapLine(null)).toBeUndefined();
  });

  it('names the caps for a tank who starts with expertise, never a 0% cap', () => {
    // The real tank shape: 8% expertise, dodge already gone, parry 6% away (boss parry 14%).
    expect(guideExpertiseCapLine({ baseline: 8, dodge: 0, parry: 6.000000000000002 })?.text).toBe(
      'Expertise cap: reached (dodge), 14% (parry) · 8% expertise from BiS gear',
    );
    expect(guideExpertiseCapLine({ baseline: 2, dodge: 4.5 }, [WEAPON_EXPERTISE])?.text).toBe(
      'Expertise cap: 6.5% (dodge) · 2% from Weapon Expertise',
    );
  });
});
