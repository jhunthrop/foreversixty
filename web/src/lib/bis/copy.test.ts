// web/src/lib/bis/copy.test.ts
import { describe, expect, it } from 'vitest';
import { bisCopy, pvpRankTitle } from './copy';

describe('pvpRankTitle', () => {
  it('reads Vanilla’s own rank ladder, offset by the client’s +4 RequiredPVPRank', () => {
    expect(pvpRankTitle('alliance', 5)).toBe('Private');
    expect(pvpRankTitle('alliance', 11)).toBe('Knight-Lieutenant');
    expect(pvpRankTitle('alliance', 18)).toBe('Grand Marshal');
    expect(pvpRankTitle('horde', 5)).toBe('Scout');
    expect(pvpRankTitle('horde', 11)).toBe('Blood Guard');
    expect(pvpRankTitle('horde', 18)).toBe('High Warlord');
  });

  it('is undefined for a rank outside the 14-rank ladder', () => {
    expect(pvpRankTitle('alliance', 4)).toBeUndefined();
    expect(pvpRankTitle('alliance', 19)).toBeUndefined();
  });
});

describe('pvpSourceLabel', () => {
  it('names the rank, its title and the faction -- never the bare "Rank N" bucket name', () => {
    expect(bisCopy.pvpSourceLabel(11, 'alliance')).toBe('PvP rank 11 · Knight-Lieutenant · Alliance');
    expect(bisCopy.pvpSourceLabel(11, 'horde')).toBe('PvP rank 11 · Blood Guard · Horde');
  });

  it('falls back to rank and faction alone when the rank has no known title', () => {
    expect(bisCopy.pvpSourceLabel(4, 'alliance')).toBe('PvP rank 4 · Alliance');
  });
});

describe('unit-aware wording', () => {
  it('defaults every figure to DPS and swaps only the unit for a healer', () => {
    expect(bisCopy.alternativeGapLabel(-0.8, 'HPS')).toBe('−0.8 HPS');
    expect(bisCopy.alternativeGapLabel(0, 'HPS')).toBe('same HPS');
    expect(bisCopy.indexSpecDps60(412.34)).toBe('Level 60: 412.3 DPS');
    expect(bisCopy.indexSpecDps60(412.34, 'HPS')).toBe('Level 60: 412.3 HPS');
    expect(bisCopy.runnersUpHeaderLabel()).toBe('Runners-up · DPS vs the pick');
    expect(bisCopy.runnersUpHeaderLabel('HPS')).toBe('Runners-up · HPS vs the pick');
    expect(bisCopy.youGainOverWorn(2.5, 'Old Ring', 'HPS')).toBe('+2.5 HPS over your Old Ring');
    expect(bisCopy.setDpsDelta(3, 'HPS')).toBe('+3.0 HPS since the last band');
    expect(bisCopy.lowValueTitle('HPS')).toContain('adds HPS');
    expect(bisCopy.headerSummary(60, undefined)).toContain('raises your damage most at 60');
    expect(bisCopy.headerSummary(50, 59, 'HPS')).toContain('raises your healing most from 50 to 59');
    expect(bisCopy.noSourcedItemFirst('trinket', '20 to 29', 30)).toContain('raises your damage. The first');
    expect(bisCopy.noSourcedItemFirst('trinket', '20 to 29', undefined, 'HPS')).toBe(
      'No trinket you can get at 20 to 29 raises your healing.',
    );
  });
});

describe('alternativeGapLabel', () => {
  it('reads "same DPS" for an exact tie', () => {
    expect(bisCopy.alternativeGapLabel(0)).toBe('same DPS');
  });

  it('reads "same DPS" for anything under the 0.05 DPS tie threshold (fix round 1: the ranker’s real dps_delta is a float, essentially never an exact 0)', () => {
    expect(bisCopy.alternativeGapLabel(0.03)).toBe('same DPS');
    expect(bisCopy.alternativeGapLabel(-0.049)).toBe('same DPS');
  });

  it('reads a signed, rounded number at or above the threshold', () => {
    expect(bisCopy.alternativeGapLabel(-0.8)).toBe('−0.8 DPS');
    expect(bisCopy.alternativeGapLabel(0.05)).toBe('+0.1 DPS');
    expect(bisCopy.alternativeGapLabel(1.2)).toBe('+1.2 DPS');
  });
});

describe('alternativeMetaLabel', () => {
  it('names only the item level when the requirement is at or under the band', () => {
    expect(bisCopy.alternativeMetaLabel(24, 18, 20)).toBe('ilvl 24');
    expect(bisCopy.alternativeMetaLabel(24, 20, 20)).toBe('ilvl 24');
  });

  it('adds "needs <level>" only when the requirement is above the band', () => {
    expect(bisCopy.alternativeMetaLabel(24, 21, 20)).toBe('ilvl 24 · needs 21');
  });
});
