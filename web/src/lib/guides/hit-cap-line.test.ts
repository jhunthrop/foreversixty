// web/src/lib/guides/hit-cap-line.test.ts
import { describe, expect, it } from 'vitest';
import type { DecodedTalentView, DecodedTreeView } from './build-tree';
import { guideExpertiseCapLine, guideHitCapLine, schoolHitTalent } from './hit-cap-line';

function tree(...talents: [string, number][]): DecodedTreeView {
  const views = talents.map(
    ([name, rank], index) =>
      ({
        id: index,
        name,
        icon: 'x',
        max_rank: 3,
        tier: index,
        column: 0,
        prereq_talent_id: null,
        prereq_rank: null,
        ranks: [],
        spell_id: index,
        rank,
      }) satisfies DecodedTalentView,
  );
  return { name: 'T', position: 0, talents: views, points: views.reduce((s, t) => s + t.rank, 0) };
}

describe('guideHitCapLine', () => {
  it('splits a caster’s cap into talent and what is still short: "Hit cap: 16% · 3% from Precision, 13% still to find"', () => {
    const line = guideHitCapLine(
      { kind: 'spell', baseline: 0, spell: 16, school: { names: ['Frost'], hit: 3, to_cap: 13 } },
      'Precision',
    );
    expect(line?.text).toBe('Hit cap: 16% · 3% from Precision, 13% still to find');
  });

  it('says "talents" when the build names no school-hit talent', () => {
    const line = guideHitCapLine({
      kind: 'spell',
      baseline: 0,
      spell: 16,
      school: { names: ['Frost'], hit: 3, to_cap: 13 },
    });
    expect(line?.text).toBe('Hit cap: 16% · 3% from talents, 13% still to find');
  });

  it('leaves the split off when no talent adds hit', () => {
    expect(guideHitCapLine({ kind: 'spell', baseline: 0, spell: 16 })?.text).toBe('Hit cap: 16%');
  });

  it('words a melee cap, with the white-swing distance for a dual wielder', () => {
    expect(guideHitCapLine({ baseline: 0, specials: 9, white: 28 })?.text).toBe(
      'Hit cap: 9% (special attacks), 28% (white swings)',
    );
    expect(guideHitCapLine({ baseline: 0, specials: 6.5 })?.text).toBe('Hit cap: 6.5%');
  });

  it('names the cap, not the distance to it, when the character starts with some hit', () => {
    expect(guideHitCapLine({ kind: 'spell', baseline: 6, spell: 10 })?.text).toBe(
      'Hit cap: 16% · 6% hit from BiS gear and buffs, 10% still to find',
    );
    expect(guideHitCapLine({ baseline: 3, specials: 6 })?.text).toBe(
      'Hit cap: 9% · 3% hit from BiS gear, talents and buffs, 6% still to find',
    );
    expect(guideHitCapLine({ baseline: 3, specials: 6, white: 25 })?.text).toBe(
      'Hit cap: 9% (special attacks), 28% (white swings) · 3% hit from BiS gear, talents and buffs, 6% still to find',
    );
    expect(
      guideHitCapLine(
        { kind: 'spell', baseline: 2, spell: 14, school: { names: ['Frost'], hit: 5, to_cap: 11 } },
        'Precision',
      )?.text,
    ).toBe('Hit cap: 16% · 2% hit from BiS gear and buffs, 3% from Precision, 11% still to find');
  });

  it('reads "reached" for a cap the character is already past', () => {
    expect(guideHitCapLine({ baseline: 10, specials: 0 })?.text).toBe(
      'Hit cap: reached · 10% hit from BiS gear, talents and buffs',
    );
  });

  it('is undefined when the band publishes no cap', () => {
    expect(guideHitCapLine(null)).toBeUndefined();
    expect(guideHitCapLine(undefined)).toBeUndefined();
  });
});

describe('schoolHitTalent', () => {
  it('finds the school-hit talent the build spends points in', () => {
    expect(schoolHitTalent([tree(['Frostbolt', 5], ['Precision', 3])])).toBe('Precision');
  });

  it('ignores one with no points and returns undefined when there is none', () => {
    expect(schoolHitTalent([tree(['Precision', 0])])).toBeUndefined();
    expect(schoolHitTalent([])).toBeUndefined();
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
      'Expertise cap: reached (dodge), 14% (parry) · 8% expertise from BiS gear and talents',
    );
    expect(guideExpertiseCapLine({ baseline: 2, dodge: 4.5 })?.text).toBe(
      'Expertise cap: 6.5% (dodge) · 2% expertise from BiS gear and talents',
    );
  });
});
