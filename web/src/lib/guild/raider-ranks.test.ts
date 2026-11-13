// web/src/lib/guild/raider-ranks.test.ts
import { describe, expect, it } from 'vitest';
import { raiderRankRows, raiderRanksCopy, raiderRanksSummary, sameRanks, toggleRank } from './raider-ranks';

const LINES = [
  { rank_index: 3, members: 12, at_sixty: 10, names: ['Bramblek', 'Ostavine', 'Pellwyn'] },
  { rank_index: 0, members: 1, at_sixty: 1, names: ['Harrowgate'] },
  { rank_index: 1, members: 2, at_sixty: 2, names: ['Kestrelle', 'Dunmarrow'] },
  { rank_index: 4, members: 5, at_sixty: 1, names: [] },
];

describe('raiderRankRows', () => {
  it('lists every roster rank highest first, with its role and holders', () => {
    const rows = raiderRankRows(LINES, 1, [3]);
    expect(rows.map((row) => row.rankIndex)).toEqual([0, 1, 3, 4]);
    expect(rows.map((row) => row.role)).toEqual(['Guild master', 'Officer', '', '']);
    expect(rows[2]).toEqual({
      rankIndex: 3,
      title: 'Rank 3',
      role: '',
      sub: '10 of 12 at 60 · Bramblek, Ostavine, Pellwyn\u00a0+9',
      checked: true,
    });
    expect(rows[3].sub).toBe('1 of 5 at 60');
    expect(rows[0].sub).toBe('1 at 60 · Harrowgate');
    expect(rows[1].sub).toBe('all 2 at 60 · Kestrelle, Dunmarrow');
    expect(rows.filter((row) => row.checked).map((row) => row.rankIndex)).toEqual([3]);
  });

  it('keeps a saved rank nobody holds now, so a save never drops it unseen', () => {
    const rows = raiderRankRows(LINES, 1, [3, 6]);
    expect(rows.at(-1)).toMatchObject({ rankIndex: 6, sub: raiderRanksCopy.nobody, checked: true });
  });

  it('is empty with no synced ranks and nothing saved', () => {
    expect(raiderRankRows([], 1, [])).toEqual([]);
  });
});

describe('toggleRank and sameRanks', () => {
  it('ticks and clears without touching the input', () => {
    const chosen = [4, 2];
    expect(toggleRank(chosen, 3, true)).toEqual([2, 3, 4]);
    expect(toggleRank(chosen, 4, false)).toEqual([2]);
    expect(toggleRank(chosen, 2, true)).toEqual([2, 4]);
    expect(chosen).toEqual([4, 2]);
  });

  it('compares as sets of ranks', () => {
    expect(sameRanks([3, 2], [2, 3])).toBe(true);
    expect(sameRanks([2], [2, 3])).toBe(false);
    expect(sameRanks([], [])).toBe(true);
  });
});

describe('raiderRanksSummary', () => {
  it('says the addon counts every level 60 with no box ticked', () => {
    expect(raiderRanksSummary(LINES, [])).toBe(raiderRanksCopy.unset);
    expect(raiderRanksSummary(null, [])).toBe(raiderRanksCopy.unset);
  });

  it('counts the ticked ranks, the characters they cover and the 60s among them', () => {
    expect(raiderRanksSummary(LINES, [3, 4])).toBe('17 characters, 11 at 60');
    expect(raiderRanksSummary(LINES, [0])).toBe('1 character, 1 at 60');
  });

  it('leaves out what it cannot know', () => {
    expect(raiderRanksSummary(null, [2, 3])).toBe('2 ranks');
    const older = LINES.map(({ at_sixty: _, ...line }) => line);
    expect(raiderRanksSummary(older, [3])).toBe('12 characters');
  });
});
