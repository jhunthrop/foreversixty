// The Forever Score panel's parts (directions mock b2-character): one scale, the quality ladder.
import { describe, expect, it } from 'vitest';
import { fromLowest, partQuality, scorePartRows, scoreTrend } from './score-parts';

describe('scoreTrend', () => {
  const history = [
    { date: '2026-10-20', score: 1520, tier: 'epic' },
    { date: '2026-10-27', score: 1690, tier: 'epic' },
    { date: '2026-11-03', score: 1830, tier: 'epic' },
  ] as const;

  it('is the series oldest first, the change since its first point and that date', () => {
    expect(scoreTrend([...history].reverse())).toEqual({
      values: [0, 170, 310],
      change: 310,
      since: '2026-10-20',
    });
  });

  it('is null with fewer than two points: one week is not a trend', () => {
    expect(scoreTrend(history.slice(0, 1))).toBeNull();
    expect(scoreTrend(undefined)).toBeNull();
  });
});

describe('fromLowest', () => {
  it('draws the movement: every point over the lowest', () => {
    expect(fromLowest([63, 60, 65])).toEqual([3, 0, 5]);
  });
});

describe('partQuality', () => {
  it('reads how full a part is on the item quality ladder', () => {
    expect([0.95, 0.9, 0.8, 0.75, 0.6, 0.5, 0.49, 0].map(partQuality)).toEqual([
      'epic',
      'epic',
      'rare',
      'rare',
      'uncommon',
      'uncommon',
      'common',
      'common',
    ]);
  });
});

describe('scorePartRows', () => {
  const rows = scorePartRows({ gear: 410, experience: 520, parses: 470, attendance: 330, discipline: 100 });

  it('keeps the contract order, each part labelled with its points', () => {
    expect(rows.map((row) => [row.key, row.label, row.points])).toEqual([
      ['gear', 'Gear', 410],
      ['experience', 'Experience', 520],
      ['parses', 'Parses', 470],
      ['attendance', 'Attendance', 330],
      ['discipline', 'Preparation', 100],
    ]);
  });

  it('draws every track on one scale: as long as its maximum against the largest', () => {
    expect(rows.map((row) => Number(row.track.toFixed(3)))).toEqual([0.714, 1, 0.857, 0.286, 0.143]);
  });

  it('fills to the share of its own maximum, capped at full, coloured by the ladder', () => {
    expect(rows.map((row) => Number(row.fill.toFixed(3)))).toEqual([0.82, 0.743, 0.783, 1, 1]);
    expect(rows.map((row) => row.color)).toEqual([
      'var(--color-rarity-rare)',
      'var(--color-rarity-uncommon)',
      'var(--color-rarity-rare)',
      'var(--color-rarity-epic)',
      'var(--color-rarity-epic)',
    ]);
  });

  it('reads a missing part as zero', () => {
    expect(scorePartRows(undefined).every((row) => row.points === 0 && row.fill === 0)).toBe(true);
  });
});
