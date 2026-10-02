// web/src/lib/planner/band-compare.test.ts
import { describe, expect, it } from 'vitest';
import talents from '../../fixtures/planner/talents/warrior.json';
import { bandTreeRanksFrom, diffAgainstBand, loadFromBand } from './band-compare';
import { indexTalents } from './rules';
import type { TalentFile } from './types';

const index = indexTalents(talents as TalentFile);

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

  it('sums the absolute rank difference per talent id, same as talentDeltaFor', () => {
    // Arms talent 1001 (tier 0, column 0) at rank 3 in the band, 0 in the build.
    const result = diffAgainstBand(index, new Map(), '3-0-0');
    expect(result.diffCount).toBe(3);
    expect([...result.diffTalentIds]).toEqual([1001]);
    expect(result.bandRankById.get(1001)).toBe(3);
  });

  it('counts a build that overspends a cell the band does not, same as underspending one it does', () => {
    const result = diffAgainstBand(index, new Map([[1001, 3]]), '0-0-0');
    expect(result.diffCount).toBe(3);
    expect([...result.diffTalentIds]).toEqual([1001]);
  });

  it('matches (zero diff, no ids) once the build equals the band', () => {
    const result = diffAgainstBand(index, new Map([[1001, 3]]), '3-0-0');
    expect(result.diffCount).toBe(0);
    expect(result.diffTalentIds.size).toBe(0);
  });

  it('clamps a band digit above a talent’s own max rank, never reporting more than it can hold', () => {
    // 1001's max_rank is 3; a band digit of 9 clamps to 3, same as orderFromRanks would.
    const result = diffAgainstBand(index, new Map(), '9-0-0');
    expect(result.bandRankById.get(1001)).toBe(3);
    expect(result.diffCount).toBe(3);
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
