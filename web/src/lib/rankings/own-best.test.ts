// web/src/lib/rankings/own-best.test.ts
import { describe, expect, it } from 'vitest';
import type { RankingRow } from './api';
import { ownBestFacts } from './own-best';

const ROW: RankingRow = {
  rank: 14,
  player: { key: 'us/normal/obnoxious-yell', name: 'Obnoxious Yell', class: 'Warrior', spec: 'Fury' },
  value: 612,
  size: 40,
  fought_at: '2026-12-10T22:10:00Z',
  duration_ms: 161_000,
  talent_split: '31/20/0',
  trinkets: [],
  buff_count: 0,
  report_id: 'r',
  fight_index: 1,
  state: 'ok',
  execution_score: 0.94,
};

describe('ownBestFacts', () => {
  it('reads the placed kill as DPS, Executed, Length and the day', () => {
    const facts = ownBestFacts({ character_rank: { rank: 14, row: ROW } }, { unit: 'DPS' });
    expect(facts.map((fact) => [fact.label, fact.value])).toEqual([
      ['DPS', '612'],
      ['Executed', '94%'],
      ['Length', '2:41'],
      ['Killed', expect.stringMatching(/10 Dec/)],
    ]);
  });

  it('has nothing when the board did not place the viewer', () => {
    expect(ownBestFacts({ character_rank: null }, { unit: 'DPS' })).toEqual([]);
    expect(ownBestFacts({}, { unit: 'DPS' })).toEqual([]);
  });
});
