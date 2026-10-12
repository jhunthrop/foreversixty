// web/src/lib/score/sparkline.test.ts
// Task L8.4 ruling 3: the Forever Score card's sparkline as a pure model -- score_history in,
// the line, the newest point and its tier colour, the two date labels and the table rows out.
import { describe, expect, it } from 'vitest';
import labels from '../../../../docs/contracts/fixtures/score-labels.json';
import type { ScorePoint } from './score';
import { SPARKLINE_HEIGHT, SPARKLINE_PAD, SPARKLINE_WIDTH, scoreSparkline } from './sparkline';

const point = (date: string, score: number, tier = 'rare'): ScorePoint => ({ date, score, tier });

/** The (x, y) pairs a path's M/L commands draw through. */
function coords(path: string): [number, number][] {
  return [...path.matchAll(/[ML]\s*(-?[\d.]+)[ ,](-?[\d.]+)/g)].map((m) => [Number(m[1]), Number(m[2])]);
}

describe('scoreSparkline', () => {
  it('is null for no history, for one point, and for a missing list', () => {
    expect(scoreSparkline([])).toBeNull();
    expect(scoreSparkline([point('2026-11-04', 1830)])).toBeNull();
    expect(scoreSparkline(undefined)).toBeNull();
    expect(scoreSparkline(null)).toBeNull();
  });

  it('draws one vertex per point, left to right, oldest first, inside the box', () => {
    const line = scoreSparkline([
      point('2026-11-01', 900, 'uncommon'),
      point('2026-11-02', 1200),
      point('2026-11-04', 1100),
    ]);
    expect(line).not.toBeNull();
    const xy = coords(line!.path);
    expect(xy).toHaveLength(3);
    expect(xy[0][0]).toBeLessThan(xy[1][0]);
    expect(xy[1][0]).toBeLessThan(xy[2][0]);
    for (const [x, y] of xy) {
      expect(x).toBeGreaterThanOrEqual(0);
      expect(x).toBeLessThanOrEqual(SPARKLINE_WIDTH);
      expect(y).toBeGreaterThanOrEqual(0);
      expect(y).toBeLessThanOrEqual(SPARKLINE_HEIGHT);
    }
    // Higher score, higher on the page (smaller y): 1200 tops, 900 bottoms.
    expect(xy[1][1]).toBeLessThan(xy[2][1]);
    expect(xy[2][1]).toBeLessThan(xy[0][1]);
    expect(line!.path.startsWith('M')).toBe(true);
  });

  it('marks the newest point in its tier’s colour', () => {
    const line = scoreSparkline([point('2026-11-01', 900, 'uncommon'), point('2026-11-04', 1830, 'epic')])!;
    const xy = coords(line.path);
    expect(line.last.x).toBeCloseTo(xy[1][0]);
    expect(line.last.y).toBeCloseTo(xy[1][1]);
    expect(line.last.color).toBe(labels.tiers.epic.color);
    expect(line.last.score).toBe(1830);
  });

  it('leaves the mark uncoloured (null) for a tier the site does not know', () => {
    const line = scoreSparkline([point('2026-11-01', 900), point('2026-11-04', 1830, 'mythic')])!;
    expect(line.last.color).toBeNull();
  });

  it('labels the first and last dates the way the addon dates a score ("4 Nov")', () => {
    const line = scoreSparkline([
      point('2026-10-12', 900),
      point('2026-10-20', 1000),
      point('2026-11-04', 1100),
    ])!;
    expect(line.firstLabel).toBe('12 Oct');
    expect(line.lastLabel).toBe('4 Nov');
  });

  it('spans at least 200 points vertically, centred on the data, so a tiny move stays small', () => {
    const line = scoreSparkline([point('2026-11-01', 1000), point('2026-11-02', 1010)])!;
    const [[, older], [, newer]] = coords(line.path);
    const inner = SPARKLINE_HEIGHT - 2 * SPARKLINE_PAD;
    // 10 points of a 200-point span, centred on 1005: both sit near the middle.
    expect(older - newer).toBeCloseTo((10 / 200) * inner, 0);
    expect((older + newer) / 2).toBeCloseTo(SPARKLINE_HEIGHT / 2, 0);
  });

  it('uses the data’s own range once it is wider than 200 points', () => {
    const line = scoreSparkline([point('2026-11-01', 900), point('2026-11-02', 1500)])!;
    const [[, low], [, high]] = coords(line.path);
    expect(low).toBeCloseTo(SPARKLINE_HEIGHT - SPARKLINE_PAD);
    expect(high).toBeCloseTo(SPARKLINE_PAD);
  });

  it('gives a flat history a flat line through the middle', () => {
    const line = scoreSparkline([point('2026-11-01', 1000), point('2026-11-02', 1000)])!;
    const ys = coords(line.path).map(([, y]) => y);
    expect(ys[0]).toBeCloseTo(SPARKLINE_HEIGHT / 2);
    expect(ys[1]).toBeCloseTo(SPARKLINE_HEIGHT / 2);
  });

  it('keeps the newest 30 points and orders them by date whatever order they came in', () => {
    const many = Array.from({ length: 32 }, (_, i) => {
      const day = new Date(Date.UTC(2026, 9, 1 + i)).toISOString().slice(0, 10);
      return point(day, 500 + i * 10);
    }).reverse();
    const line = scoreSparkline(many)!;
    expect(coords(line.path)).toHaveLength(30);
    expect(line.rows).toHaveLength(30);
    expect(line.rows[0]).toEqual({ date: '3 Oct 2026', score: 520 });
    expect(line.rows[29]).toEqual({ date: '1 Nov 2026', score: 810 });
    expect(line.firstLabel).toBe('3 Oct');
    expect(line.lastLabel).toBe('1 Nov');
  });

  it('gives the table equivalent every point with its full date and score, oldest first', () => {
    const line = scoreSparkline([point('2026-11-01', 900), point('2026-11-04', 1830, 'epic')])!;
    expect(line.rows).toEqual([
      { date: '1 Nov 2026', score: 900 },
      { date: '4 Nov 2026', score: 1830 },
    ]);
  });

  it('drops points whose date or score cannot be read, then needs two of what is left', () => {
    const bad = { date: 'yesterday', score: 1000, tier: 'rare' };
    const nan = { date: '2026-11-03', score: Number.NaN, tier: 'rare' };
    expect(scoreSparkline([bad, point('2026-11-04', 1100)])).toBeNull();
    expect(
      coords(scoreSparkline([point('2026-11-01', 900), nan, point('2026-11-04', 1100)])!.path),
    ).toHaveLength(2);
  });
});
