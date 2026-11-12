// web/src/lib/rankings/rank-row-view.test.ts
import { describe, expect, it } from 'vitest';
import type { RankingRow } from './api';
import { killLength, percentileLabel, rankRowView } from './rank-row-view';

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
const BOARD = { total: 212, top: 888, unit: 'DPS' };

describe('rankRowView', () => {
  it('reads a kill as the RankRow draws it', () => {
    const view = rankRowView(ROW, BOARD);
    expect(view.rank).toBe('1');
    expect(view.rankColor).toBe('var(--color-parse-gold)');
    expect(view.sub).toBe('Olympus XXVII · 17/34/0');
    expect(view.fraction).toBe(1);
    // The Date rides under the figure below 768 (phone only), beside the percentile.
    expect(view.figures.map((figure) => figure.value)).toEqual(['888', '100', view.gap.text, '90%', '2:41']);
    expect(view.figures[2]).toMatchObject({ label: 'Date', bare: true, phoneOnly: true });
    expect(view.figures[0]?.href).toBe('/reports/rep1?fight=3');
    expect(view.links).toEqual([
      {
        label: 'Report',
        href: '/reports/rep1?fight=3',
        title: 'Open the kill in its report',
        testid: 'ranking-report',
      },
    ]);
    expect(view.note).toBeUndefined();
  });

  it('colours ranks 2 and 3 of 212 pink and 4 orange (B4), names a moderated kill', () => {
    expect(rankRowView({ ...ROW, rank: 2 }, BOARD).rankColor).toBe('var(--color-parse-pink)');
    expect(rankRowView({ ...ROW, rank: 4 }, BOARD).rankColor).toBe('var(--color-parse-orange)');
    expect(rankRowView({ ...ROW, state: 'removed' }, BOARD).note?.text).toBe('Removed by a moderator');
  });

  it('names the spec under the guild on a board of several classes (the crest says the class)', () => {
    const view = rankRowView(ROW, { ...BOARD, multiClass: true });
    expect(view.sub).toBe('Olympus XXVII · Fury');
    expect(rankRowView({ ...ROW, guild: undefined }, { ...BOARD, multiClass: true }).sub).toBe(
      'Fury Warrior',
    );
    expect(rankRowView({ ...ROW, guild: undefined }, BOARD).sub).toBe('17/34/0');
  });

  it('floors the percentile the colour reads: 100, then 99 for ranks 2 and 3 of 212', () => {
    expect(percentileLabel(1, 212)).toBe('100');
    expect(percentileLabel(2, 212)).toBe('99');
    expect(percentileLabel(3, 212)).toBe('99');
    expect(percentileLabel(4, 212)).toBe('98');
    expect(percentileLabel(212, 212)).toBe('0');
  });

  it('prints a length as m:ss', () => {
    expect(killLength(161_000)).toBe('2:41');
    expect(killLength(9_400)).toBe('0:09');
  });
});

describe('a row whose log named no spec', () => {
  it('names the class alone, never "undefined"', () => {
    const row = {
      rank: 1,
      player: { key: 'us/pvp/psycho-giga', name: 'Psycho Giga', class: 'Shaman' },
      value: 113,
      size: 5,
      fought_at: '2026-11-09T00:20:16Z',
      duration_ms: 47224,
      talent_split: '0/0/0',
      trinkets: [],
      buff_count: 0,
      report_id: 'gxrgvz4yzsi2',
      fight_index: 40,
      state: 'ok',
      execution_score: null,
    } as unknown as RankingRow;
    const view = rankRowView(row, { total: 2, top: 113, unit: 'DPS', multiClass: true });
    expect(view.sub).toBe('Shaman');
  });
});
