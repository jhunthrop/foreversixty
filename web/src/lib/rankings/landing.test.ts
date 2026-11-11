// web/src/lib/rankings/landing.test.ts
import { describe, expect, it } from 'vitest';
import type { RankedEncounter, RankingRow } from './api';
import { leaderEncounter, specLeaders, specRole, topParses } from './landing';

const row = (rank: number, name: string, cls: string, spec: string, value: number): RankingRow => ({
  rank,
  player: { key: `us/normal/${name.toLowerCase()}`, name, class: cls, spec },
  value,
  size: 40,
  fought_at: '2026-12-12T21:00:00Z',
  duration_ms: 90_000,
  talent_split: '31/20/0',
  trinkets: [],
  buff_count: 0,
  report_id: `r${rank}`,
  fight_index: 1,
  state: 'ok',
  execution_score: null,
});

const encounter = (name: string, size: number | null, kills: number): RankedEncounter => ({
  id: kills,
  name,
  slug: name.toLowerCase(),
  zone: null,
  size,
  kills,
});

describe('leaderEncounter', () => {
  it('is the raid boss with the most kills, else the first encounter, else none', () => {
    expect(leaderEncounter([encounter('Saltspine', 5, 9), encounter('Onyxia', 40, 4)])?.name).toBe('Onyxia');
    expect(leaderEncounter([encounter('Saltspine', 5, 9)])?.name).toBe('Saltspine');
    expect(leaderEncounter([])).toBeNull();
  });
});

describe('specRole', () => {
  it('reads the role from the class and spec, DPS when unknown', () => {
    expect(specRole('Priest', 'Holy')).toBe('healer');
    expect(specRole('Warrior', 'Protection')).toBe('tank');
    expect(specRole('Warrior', 'Fury')).toBe('dps');
    expect(specRole('Warrior', 'Gladiator')).toBe('dps');
  });
});

describe('topParses', () => {
  it('keeps the first three rows with the value and its percentile', () => {
    const board = {
      total: 212,
      rows: [1, 2, 3, 4].map((n) => row(n, `P${n}`, 'Warrior', 'Fury', 900 - n)),
    };
    const top = topParses(board);
    expect(top.map((r) => r.rank)).toEqual(['1', '2', '3']);
    expect(top[0]?.figures.map((f) => [f.value, f.unit ?? ''])).toEqual([
      ['899', ''],
      ['100', ''],
    ]);
    expect(top[0]?.figures[0]?.href).toBe('/reports/r1?fight=1');
  });
});

describe('specLeaders', () => {
  const href = (cls: string, spec: string) => `/rankings/onyxia?class=${cls}&spec=${spec}`;

  it('groups damage by DPS and healing by HPS, each spec once, best first, spec leading the row', () => {
    const dps = {
      total: 5,
      rows: [
        row(1, 'Fury1', 'Warrior', 'Fury', 900),
        row(2, 'Fury2', 'Warrior', 'Fury', 880),
        row(3, 'Mage1', 'Mage', 'Frost', 870),
        row(4, 'HolyDps', 'Priest', 'Holy', 100),
      ],
    };
    const hps = {
      total: 2,
      rows: [row(1, 'Holy1', 'Priest', 'Holy', 700), row(2, 'Fury9', 'Warrior', 'Fury', 5)],
    };
    const { damage, healing } = specLeaders(dps, hps, href);
    expect(damage.map((r) => [r.rank, r.name, r.sub])).toEqual([
      ['1', 'Fury Warrior', 'Fury1'],
      ['2', 'Frost Mage', 'Mage1'],
    ]);
    expect(damage[1]?.href).toBe('/rankings/onyxia?class=mage&spec=Frost');
    expect(healing.map((r) => [r.name, r.figures[0]?.value])).toEqual([['Holy Priest', '700']]);
  });

  it('has no healing group when the HPS board did not load', () => {
    const board = { total: 1, rows: [row(1, 'Fury1', 'Warrior', 'Fury', 900)] };
    expect(specLeaders(board, null, href).healing).toEqual([]);
  });
});

describe('specLeaders with a log that named no spec', () => {
  it('leaves the row out instead of throwing (the API omits `spec` then)', () => {
    const bare = {
      ...row(1, 'Psycho', 'Shaman', '', 113),
      player: { key: 'us/pvp/psycho', name: 'Psycho', class: 'Shaman' },
    };
    const board = { total: 2, rows: [bare as RankingRow, row(2, 'Thoralf', 'Paladin', 'Retribution', 102)] };
    expect(specLeaders(board, null, () => '#').damage.map((leader) => leader.name)).toEqual([
      'Retribution Paladin',
    ]);
  });
});
