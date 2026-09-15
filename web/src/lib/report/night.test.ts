import { describe, expect, it } from 'vitest';
import { aggregateNight } from './night';
import type { FightEntry, RosterRow, Summary } from './types';

function fight(index: number, name: string, kill: boolean, duration_ms: number): FightEntry {
  return {
    index,
    kind: 'encounter',
    name,
    kill,
    in_progress: false,
    start: '2026-09-26T20:10:00Z',
    end: '2026-09-26T20:11:00Z',
    duration_ms,
    players: [],
    deaths: 0,
    npc_kills: 1,
  };
}

function roster(guid: string, damage: number, deaths = 0): RosterRow {
  return {
    guid,
    name: guid,
    class: 'Shaman',
    spec: 'Elemental',
    role: 'dps',
    active_ms: 5000,
    activity_pct: 50,
    deaths,
    damage_done: damage,
    healing_done: 0,
    damage_taken: 100,
    dps: 0,
    hps: 0,
    dtps: 0,
  };
}

function summary(fight_index: number, duration_ms: number, rows: RosterRow[]): Summary {
  return {
    engine_version: 't',
    fight_index,
    duration_ms,
    damage_done: [],
    damage_taken: [],
    healing: [],
    healing_taken: [],
    deaths: rows.flatMap((row) =>
      Array.from({ length: row.deaths }, () => ({ guid: row.guid })),
    ) as Summary['deaths'],
    auras: [],
    casts: [],
    interrupts: [],
    dispels: [],
    resources: [],
    threat: [],
    combatants: [],
    roster: rows,
  };
}

describe('aggregateNight', () => {
  const fights = [
    fight(1, 'Kaal', false, 10_000),
    { ...fight(2, 'Trash', false, 5000), kind: 'trash' },
    fight(3, 'Kaal', true, 20_000),
    fight(4, 'Kryxis', true, 30_000),
  ];
  const summaries = new Map<number, Summary>([
    [1, summary(1, 10_000, [roster('A', 1000, 1), roster('B', 500)])],
    [3, summary(3, 20_000, [roster('A', 4000), roster('B', 1000)])],
  ]);

  it('folds every loaded boss pull into one row per player, per second over their own time', () => {
    const night = aggregateNight(fights, summaries);
    expect(night.fights).toBe(2);
    expect(night.expected).toBe(3);
    expect(night.time_ms).toBe(30_000);
    expect(night.deaths).toBe(1);
    const a = night.players[0];
    expect(a.guid).toBe('A');
    expect(a.damage_done).toBe(5000);
    expect(a.dps).toBeCloseTo(5000 / 30);
    expect(a.deaths).toBe(1);
    expect(a.fights).toBe(2);
    expect(a.by_fight.map((entry) => entry.boss)).toEqual(['Kaal', 'Kaal']);
  });

  it('counts pulls, kills, wipes and the quickest kill per boss, including pulls not yet loaded', () => {
    const night = aggregateNight(fights, summaries);
    expect(night.bosses.map((boss) => boss.name)).toEqual(['Kaal', 'Kryxis']);
    expect(night.bosses[0]).toMatchObject({
      pulls: 2,
      kills: 1,
      wipes: 1,
      time_ms: 30_000,
      best: { index: 3, duration_ms: 20_000 },
    });
    expect(night.bosses[1]).toMatchObject({
      pulls: 1,
      kills: 1,
      wipes: 0,
      best: { index: 4, duration_ms: 30_000 },
    });
  });
});
