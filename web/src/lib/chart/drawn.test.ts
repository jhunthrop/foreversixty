// web/src/lib/chart/drawn.test.ts
import { describe, expect, it } from 'vitest';
import {
  cumulative,
  DEFAULT_TOP_N,
  median,
  medianSeries,
  roundedSum,
  selectDrawn,
  standingAt,
  stopAtDeath,
  trailingAverage,
  windowRate,
} from './drawn';

describe('trailingAverage', () => {
  it('averages the last five seconds, over the seconds there are at the start', () => {
    expect(trailingAverage([10, 20, 30, 40, 50, 60])).toEqual([10, 15, 20, 25, 30, 40]);
  });

  it('falls to zero and stays flat through a phase where the boss takes no damage', () => {
    // Ragnaros submerged: five seconds of nothing brings the average to exactly zero, and it
    // stays at zero (never null, never interpolated) until damage comes back.
    const series = [100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 0, 0, 0, 50];
    const averaged = trailingAverage(series);
    expect(averaged.slice(10, 13)).toEqual([0, 0, 0]);
    expect(averaged.every((value) => Number.isFinite(value))).toBe(true);
    expect(averaged[13]).toBe(10);
  });
});

describe('cumulative', () => {
  it('is a running total, so the value at a second is the standing then', () => {
    expect(cumulative([1, 2, 0, 4])).toEqual([1, 3, 3, 7]);
  });
});

describe('windowRate and standingAt', () => {
  it('is the window total over the window seconds', () => {
    expect(windowRate([10, 20, 30, 40], { startMs: 1000, endMs: 3000 })).toBe(25);
    expect(windowRate([10, 20], { startMs: 0, endMs: 0 })).toBe(0);
  });

  it('reads the cumulative value at the window end', () => {
    expect(standingAt([1, 3, 6, 10], { startMs: 0, endMs: 3000 })).toBe(6);
    expect(standingAt([1, 3, 6, 10], { startMs: 0, endMs: 9000 })).toBe(10);
  });
});

describe('roundedSum', () => {
  it('rounds each player first, so a raider adding the legend gets the head', () => {
    expect(roundedSum([100.4, 100.4, 100.4])).toBe(300);
    expect(Math.round(100.4 * 3)).toBe(301);
  });
});

describe('median and medianSeries', () => {
  it('takes the middle value, or the mean of the middle two', () => {
    expect(median([3, 1, 2])).toBe(2);
    expect(median([4, 1, 3, 2])).toBe(2.5);
    expect(median([])).toBe(0);
  });

  it('is a per-second median across the players', () => {
    expect(
      medianSeries([
        [1, 10, 5],
        [2, 20, 5],
        [9, 30, 5],
      ]),
    ).toEqual([2, 20, 5]);
  });
});

describe('selectDrawn', () => {
  const candidates = Array.from({ length: 40 }, (_, index) => ({ id: `p${index}`, figure: 1000 - index }));

  it('draws the top 8 of N by the window figure, highest first', () => {
    const picked = selectDrawn({ candidates, pinned: new Set(), chosen: null });
    expect(DEFAULT_TOP_N).toBe(8);
    expect(picked).toEqual(['p0', 'p1', 'p2', 'p3', 'p4', 'p5', 'p6', 'p7']);
  });

  it('always draws the pinned players (the Source player, you) on top of the 8', () => {
    const picked = selectDrawn({ candidates, pinned: new Set(['p33', 'p20']), chosen: null });
    expect(picked).toHaveLength(10);
    expect(picked.slice(-2)).toEqual(['p20', 'p33']);
  });

  it('does not double count a pinned player already in the 8', () => {
    expect(selectDrawn({ candidates, pinned: new Set(['p1']), chosen: null })).toHaveLength(8);
  });

  it('draws exactly the chosen players, plus the pinned ones, once the Players control changes', () => {
    const picked = selectDrawn({ candidates, pinned: new Set(['p39']), chosen: new Set(['p12', 'p3']) });
    expect(picked).toEqual(['p3', 'p12', 'p39']);
  });

  it('splits the default into the top 4 and the bottom 4, so the low performers are named', () => {
    const picked = selectDrawn({ candidates, pinned: new Set(), chosen: null, preset: 'split' });
    expect(picked).toEqual(['p0', 'p1', 'p2', 'p3', 'p36', 'p37', 'p38', 'p39']);
  });

  it('takes the bottom from the metric’s own pool (damage dealers on damage done)', () => {
    const pool = new Set(candidates.slice(0, 30).map((entry) => entry.id));
    const split = selectDrawn({
      candidates,
      pinned: new Set(),
      chosen: null,
      preset: 'split',
      bottomPool: pool,
    });
    expect(split.slice(-4)).toEqual(['p26', 'p27', 'p28', 'p29']);
    const bottom = selectDrawn({
      candidates,
      pinned: new Set(),
      chosen: null,
      preset: 'bottom',
      bottomPool: pool,
    });
    expect(bottom).toEqual(['p22', 'p23', 'p24', 'p25', 'p26', 'p27', 'p28', 'p29']);
  });

  it('never draws a player twice when the roster is small', () => {
    const few = candidates.slice(0, 5);
    expect(selectDrawn({ candidates: few, pinned: new Set(), chosen: null, preset: 'split' })).toEqual([
      'p0',
      'p1',
      'p2',
      'p3',
      'p4',
    ]);
  });

  it('ignores a pinned player who is not in the fight', () => {
    expect(selectDrawn({ candidates, pinned: new Set(['ghost']), chosen: null })).not.toContain('ghost');
  });
});

describe('stopAtDeath', () => {
  it('ends the line at the death and resumes it after a raise', () => {
    const values = [5, 5, 5, 5, 5, 5];
    expect(stopAtDeath(values, [{ startMs: 2000, endMs: 4000 }])).toEqual([5, 5, 5, null, null, 5]);
  });

  it('leaves a living player untouched', () => {
    expect(stopAtDeath([1, 2], [])).toEqual([1, 2]);
  });
});
