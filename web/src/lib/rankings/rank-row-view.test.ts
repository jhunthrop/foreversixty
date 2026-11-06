// web/src/lib/rankings/rank-row-view.test.ts
import { describe, expect, it } from 'vitest';
import type { RankingRow } from './api';
import { killLength, rankRowView } from './rank-row-view';

const ROW: RankingRow = {
  rank: 1,
  player: { key: 'us/normal/thorgrim-ashveil', name: 'Thorgrim Ashveil', class: 'Warrior', spec: 'Fury' },
  guild: { name: 'Olympus XXVII', ruleset: 'normal', region: 'us' },
  value: 888,
  size: 40,
  fought_at: '2026-12-12T21:00:00Z',
  duration_ms: 161_000,
  talent_split: '17/34/0',
  build_id: 'abc',
  trinkets: [],
  buff_count: 0,
  report_id: 'rep1',
  fight_index: 3,
  state: 'ok',
  execution_score: 0.9,
};
const BOARD = { total: 212, top: 888, unit: 'DPS', now: new Date('2026-12-13T12:00:00Z') };

describe('rankRowView', () => {
  it('reads a kill as the RankRow draws it', () => {
    const view = rankRowView(ROW, BOARD);
    expect(view.rank).toBe('1');
    expect(view.rankColor).toBe('var(--color-parse-gold)');
    expect(view.sub).toBe('Olympus XXVII · Fury Warrior · 40-man');
    expect(view.fraction).toBe(1);
    expect(view.figures.map((figure) => figure.value)).toEqual(['888', '–', '90%', '2:41']);
    expect(view.links.map((link) => link.href)).toEqual(['/reports/rep1?fight=3', '/b/abc']);
    expect(view.note).toBeUndefined();
  });

  it('colours ranks 2 and 3 of 212 pink and 4 orange (B4), names a moderated kill', () => {
    expect(rankRowView({ ...ROW, rank: 2 }, BOARD).rankColor).toBe('var(--color-parse-pink)');
    expect(rankRowView({ ...ROW, rank: 4 }, BOARD).rankColor).toBe('var(--color-parse-orange)');
    expect(rankRowView({ ...ROW, state: 'removed' }, BOARD).note?.text).toBe('Removed by a moderator');
  });

  it('prints an item level once the row carries one, and no Build link without a build', () => {
    const view = rankRowView({ ...ROW, item_level: 71, build_id: undefined }, BOARD);
    expect(view.figures[1]?.value).toBe('71');
    expect(view.links).toHaveLength(1);
  });

  it('prints a length as m:ss', () => {
    expect(killLength(161_000)).toBe('2:41');
    expect(killLength(9_400)).toBe('0:09');
  });
});
