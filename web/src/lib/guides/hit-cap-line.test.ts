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
  it('splits a caster’s cap into talent and gear: "Hit cap: 16% · 3% from Elemental Precision, 13% from gear"', () => {
    const line = guideHitCapLine(
      { kind: 'spell', baseline: 0, spell: 16, school: { names: ['Frost'], hit: 3, to_cap: 13 } },
      'Elemental Precision',
    );
    expect(line?.text).toBe('Hit cap: 16% · 3% from Elemental Precision, 13% from gear');
  });

  it('says "talents" when the build names no school-hit talent', () => {
    const line = guideHitCapLine({
      kind: 'spell',
      baseline: 0,
      spell: 16,
      school: { names: ['Frost'], hit: 3, to_cap: 13 },
    });
    expect(line?.text).toBe('Hit cap: 16% · 3% from talents, 13% from gear');
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

  it('is undefined when the band publishes no cap', () => {
    expect(guideHitCapLine(null)).toBeUndefined();
    expect(guideHitCapLine(undefined)).toBeUndefined();
  });
});

describe('schoolHitTalent', () => {
  it('finds the school-hit talent the build spends points in', () => {
    expect(schoolHitTalent([tree(['Frostbolt', 5], ['Elemental Precision', 3])])).toBe('Elemental Precision');
  });

  it('ignores one with no points and returns undefined when there is none', () => {
    expect(schoolHitTalent([tree(['Elemental Precision', 0])])).toBeUndefined();
    expect(schoolHitTalent([])).toBeUndefined();
  });
});

describe('guideExpertiseCapLine', () => {
  it('words the dodge cap, with parry for a tank', () => {
    expect(guideExpertiseCapLine({ baseline: 0, dodge: 6.5 })?.text).toBe(
      'Expertise cap: 6.5% (dodge)',
    );
    expect(guideExpertiseCapLine({ baseline: 0, dodge: 5.5, parry: 14 })?.text).toBe(
      'Expertise cap: 5.5% (dodge), 14% (parry)',
    );
    expect(guideExpertiseCapLine(null)).toBeUndefined();
  });
});
