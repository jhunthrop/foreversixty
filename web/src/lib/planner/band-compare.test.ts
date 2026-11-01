// web/src/lib/planner/band-compare.test.ts
import { describe, expect, it } from 'vitest';
import talents from '../../fixtures/planner/talents/warrior.json';
import contractWarrior from '../../../../docs/contracts/fixtures/app/cdn.talents.warrior.json';
import {
  VISITOR_DEFAULT_BAND,
  bandDiffStatus,
  bandForPanel,
  bandTreeRanksFrom,
  diffAgainstBand,
  loadFromBand,
  loadFromBandAtLevel,
} from './band-compare';
import { bandCompareCopy } from './copy';
import { indexTalents } from './rules';
import type { TalentFile } from './types';

const index = indexTalents(talents as TalentFile);

// The case list the apps run too (TalentRulesTests.swift, TalentRulesTest.kt): the real warrior
// tree and the Fury BiS bands of the active build.
const warrior = indexTalents(contractWarrior as TalentFile);
/** The Fury BiS file's 50 to 59 Horde guide build (41 points). */
const BAND_50 = '34200000000000000-05153105022011501-000000000000000000';
/** The Fury BiS file's level-60 raid guide build (51 points). */
const BAND_60 = '34320003002000000-05153105022011501-200000000000000000';
/** A level-58 build (49 points) holding every point of BAND_50: the level-60 path less Improved Bloodrage. */
const LEVEL_58 = '34320003002-05153105022011501-0';
/** BAND_60 without Fury's 5-point second talent: 46 points, 5 of BAND_50's missing, 10 stray. */
const STRAY_AND_SHORT = '34320003002000000-00153105022011501-200000000000000000';

function ranksOf(talentIndex: typeof warrior, build: string): Map<number, number> {
  const trees = bandTreeRanksFrom(build);
  return new Map(
    talentIndex.trees.flatMap((tree, t) =>
      tree.talents.map((talent, i) => [talent.id, trees[t]?.[i] ?? 0] as [number, number]),
    ),
  );
}

describe('bandTreeRanksFrom', () => {
  it('splits a band talents string into one base-36 digit array per tree', () => {
    expect(bandTreeRanksFrom('503200000-0-0')).toEqual([[5, 0, 3, 2, 0, 0, 0, 0, 0], [0], [0]]);
  });

  it('reads base 36, so a rank above nine is a letter', () => {
    expect(bandTreeRanksFrom('c-0-0')).toEqual([[12], [0], [0]]);
  });
});

describe('diffAgainstBand', () => {
  it('reports no diff and no ids against an empty build and an all-zero band', () => {
    const result = diffAgainstBand(index, new Map(), '0-0-0');
    expect(result.diffCount).toBe(0);
    expect(result.diffTalentIds.size).toBe(0);
  });

  it('a build short of the band, but inside it, counts the band points still to take', () => {
    // Arms talent 1001 (tier 0, column 0) at rank 3 in the band, 0 in the build.
    const result = diffAgainstBand(index, new Map(), '3-0-0');
    expect(result.diffCount).toBe(0);
    expect(result.missingCount).toBe(3);
    expect([...result.diffTalentIds]).toEqual([1001]);
    expect(result.bandRankById.get(1001)).toBe(3);
  });

  it('an empty band: every build point is beyond its total, so none of them differ', () => {
    const result = diffAgainstBand(index, new Map([[1001, 3]]), '0-0-0');
    expect(result.diffCount).toBe(0);
    expect(result.missingCount).toBe(0);
    expect(result.extraCount).toBe(3);
    expect([...result.diffTalentIds]).toEqual([1001]);
  });

  it('never reports more differing points than the build holds, however different it is', () => {
    const buildPoints = 3;
    const result = diffAgainstBand(
      index,
      new Map([
        [1001, 0],
        [1002, 3],
      ]),
      '3-0-0',
    );
    expect(result.diffCount).toBeLessThanOrEqual(buildPoints);
    expect(result.diffCount).toBe(3);
    expect(result.missingCount).toBe(3);
  });

  it('matches (zero diff, no ids) once the build equals the band', () => {
    const result = diffAgainstBand(index, new Map([[1001, 3]]), '3-0-0');
    expect(result.diffCount).toBe(0);
    expect(result.missingCount).toBe(0);
    expect(result.diffTalentIds.size).toBe(0);
  });

  it('clamps a band digit above a talent’s own max rank, never reporting more than it can hold', () => {
    // 1001's max_rank is 3; a band digit of 9 clamps to 3, same as orderFromRanks would.
    const result = diffAgainstBand(index, new Map(), '9-0-0');
    expect(result.bandRankById.get(1001)).toBe(3);
    expect(result.missingCount).toBe(3);
  });
});

describe('diffAgainstBand: points beyond the band (shared case list)', () => {
  it('the same build matches with nothing beyond', () => {
    const result = diffAgainstBand(warrior, ranksOf(warrior, BAND_50), BAND_50);
    expect(result).toMatchObject({ diffCount: 0, missingCount: 0, extraCount: 0 });
    expect(bandDiffStatus(result)).toEqual({ kind: 'matches', beyond: 0 });
  });

  it('a build inside the band and short of it reads as points still to take', () => {
    const result = diffAgainstBand(warrior, ranksOf(warrior, BAND_50), BAND_60);
    expect(result).toMatchObject({ diffCount: 0, missingCount: 10, extraCount: 0 });
    expect(bandDiffStatus(result)).toEqual({ kind: 'toTake', count: 10 });
  });

  it('level 58: 49 points holding the whole 41-point band match, 8 beyond the guide', () => {
    const result = diffAgainstBand(warrior, ranksOf(warrior, LEVEL_58), BAND_50);
    expect(result).toMatchObject({ diffCount: 0, missingCount: 0, extraCount: 8 });
    expect(bandDiffStatus(result)).toEqual({ kind: 'matches', beyond: 8 });
    expect(bandCompareCopy.beyondGuide(8)).toBe('+8 beyond the guide');
  });

  it('stray points differ only up to the band total: 10 stray − (46 − 41) = 5', () => {
    const result = diffAgainstBand(warrior, ranksOf(warrior, STRAY_AND_SHORT), BAND_50);
    expect(result).toMatchObject({ diffCount: 5, missingCount: 5, extraCount: 5 });
    expect(bandDiffStatus(result)).toEqual({ kind: 'differs', count: 5, beyond: 5 });
  });
});

describe('loadFromBand', () => {
  it('reconstructs a legal order reaching the band’s own ranks, with nothing dropped', () => {
    const { order, dropped } = loadFromBand(index, '3-0-0');
    expect(order).toEqual([1001, 1001, 1001]);
    expect(dropped).toEqual([]);
  });

  it('drops a rank no legal order can reach (a tier/prerequisite the band string never grants)', () => {
    // Talent 1004 (Tactical Mastery, tier 1) needs 1002 at rank >= 2 and 5 points already
    // in the tree to clear its own tier gate; the band string asks for 1004 alone, so the
    // reconstruction cannot legally place it.
    const { order, dropped } = loadFromBand(index, '0005000-0-0');
    expect(dropped).toEqual([1004]);
    expect(order.includes(1004)).toBe(false);
  });
});

describe('loadFromBandAtLevel', () => {
  // The band's whole build is three points (level 12), so each case sits at a known cut.
  it('loads the band’s whole build with no character selected', () => {
    const result = loadFromBandAtLevel(index, '3-0-0', null);
    expect(result.order).toEqual([1001, 1001, 1001]);
    expect(result.fullPoints).toBe(3);
  });

  it('cuts the point order at a character inside the band (points = level - 9)', () => {
    const result = loadFromBandAtLevel(index, '3-0-0', 11);
    expect(result.order).toEqual([1001, 1001]);
    expect(result.fullPoints).toBe(3);
  });

  it('loads the whole build for a character past the band', () => {
    const result = loadFromBandAtLevel(index, '3-0-0', 29);
    expect(result.order).toHaveLength(3);
  });

  it('cuts to the character’s points when the character is below the band', () => {
    expect(loadFromBandAtLevel(index, '3-0-0', 10).order).toEqual([1001]);
  });

  it('does not load an empty build for a character with no point yet', () => {
    expect(loadFromBandAtLevel(index, '3-0-0', 8).order).toHaveLength(3);
  });
});

describe('bandForPanel (W.8)', () => {
  it('opens a visitor with no character and nothing spent on the level-60 build', () => {
    expect(VISITOR_DEFAULT_BAND).toBe(60);
    expect(bandForPanel(0, null, 20)).toBe(60);
  });

  it('follows the build once a point is spent, or a character is selected', () => {
    expect(bandForPanel(1, null, 20)).toBe(20);
    expect(bandForPanel(0, 34, 30)).toBe(30);
  });
});

describe('the band card’s heading', () => {
  it('titles a single level "Level 60 build" and a span "Levels 20 to 29 build"', () => {
    expect(bandCompareCopy.heading('60')).toBe('Level 60 build');
    expect(bandCompareCopy.heading('20 to 29')).toBe('Levels 20 to 29 build');
    expect(bandCompareCopy.loadButton('60')).toBe('Load the level 60 build');
  });
});
