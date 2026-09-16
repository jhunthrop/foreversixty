import { describe, expect, it } from 'vitest';
import { aggregateNight, nightSummary } from './night';
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

describe('nightSummary auras', () => {
  const track = (target_guid: string, target_name: string, spell_id: number) => ({
    target_guid,
    target_name,
    spell_id,
    name: 'Hanging Chains',
    type: 'BUFF' as const,
    school: 1,
    applications: 1,
    max_stacks: 1,
    uptime_ms: 1000,
    segments: [{ start_ms: 0, end_ms: 1000 }],
    appliers: [],
  });

  it('names a unit the log left "Unknown" on one pull from the pull that knew it, as one row', () => {
    const fights = [fight(1, 'Kryxis', false, 10000), fight(2, 'Kryxis', true, 10000)];
    const unnamed = { ...summary(1, 10000, []), auras: [track('Creature-1', 'Unknown', 5)] };
    const named = { ...summary(2, 10000, []), auras: [track('Creature-1', 'Hanging Chain', 5)] };
    const night = nightSummary(
      fights,
      new Map([
        [1, unnamed],
        [2, named],
      ]) as never,
    );
    expect(night.auras.map((row) => [row.target_guid, row.target_name, row.uptime_ms])).toEqual([
      ['Creature-1', 'Hanging Chain', 2000],
    ]);
    expect(night.auras[0]?.time_ms).toBe(20000);
  });
});

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

  it('folds the pulls into one summary the fight tabs can show, with deaths labelled by pull', () => {
    const withActors = new Map<number, Summary>([
      [
        1,
        {
          ...summary(1, 10_000, [roster('A', 1000, 1)]),
          damage_done: [
            {
              guid: 'A',
              name: 'A',
              total: 1000,
              effective: 1000,
              active_ms: 5000,
              abilities: [
                {
                  spell_id: 1,
                  name: 'Bolt',
                  total: 1000,
                  effective: 1000,
                  hits: 2,
                  crits: 1,
                  ticks: 0,
                  min: 400,
                  max: 600,
                },
              ],
              targets: [{ guid: 'B', name: 'Boss', total: 1000 }],
              series: [500, 500],
            },
          ],
          interrupts: [
            {
              kind: 'interrupt',
              source_guid: 'A',
              source_name: 'A',
              target_guid: 'B',
              target_name: 'Boss',
              spell_id: 5,
              spell_name: 'Kick',
              extra_spell_id: 9,
              extra_spell_name: 'Cast',
              count: 1,
            },
          ],
        },
      ],
      [
        3,
        {
          ...summary(3, 20_000, [roster('A', 4000)]),
          damage_done: [
            {
              guid: 'A',
              name: 'A',
              total: 4000,
              effective: 4000,
              active_ms: 15_000,
              abilities: [
                {
                  spell_id: 1,
                  name: 'Bolt',
                  total: 4000,
                  effective: 4000,
                  hits: 5,
                  crits: 2,
                  ticks: 0,
                  min: 300,
                  max: 900,
                },
              ],
              targets: [{ guid: 'B', name: 'Boss', total: 4000 }],
              series: [2000, 2000],
            },
          ],
          interrupts: [
            {
              kind: 'interrupt',
              source_guid: 'A',
              source_name: 'A',
              target_guid: 'B',
              target_name: 'Boss',
              spell_id: 5,
              spell_name: 'Kick',
              extra_spell_id: 9,
              extra_spell_name: 'Cast',
              count: 2,
            },
          ],
        },
      ],
    ]);
    const night = nightSummary(fights, withActors);
    expect(night.duration_ms).toBe(30_000);
    expect(night.damage_done).toHaveLength(1);
    expect(night.damage_done[0].effective).toBe(5000);
    expect(night.damage_done[0].abilities[0]).toMatchObject({
      total: 5000,
      hits: 7,
      crits: 3,
      min: 300,
      max: 900,
    });
    expect(night.damage_done[0].targets[0].total).toBe(5000);
    expect(night.interrupts[0].count).toBe(3);
    expect(night.deaths.map((death) => death.label)).toEqual(['Kaal · pull 1']);
    expect(night.roster[0].dps).toBeCloseTo(5000 / 30);
  });

  it('folds mechanics across pulls, counting the pulls each one hit anyone on', () => {
    const withMechanics = new Map(summaries);
    const block = (damage: number) => ({
      table_found: true,
      rows: [
        {
          spell_id: 331415,
          name: 'Wicked Gash',
          kind: 'avoidable' as const,
          players: [
            {
              guid: 'Player-1',
              name: 'Hobolol',
              hits: 1,
              damage,
              first_ms: 1000,
              last_ms: 1000,
              killed: false,
            },
          ],
        },
      ],
    });
    withMechanics.set(1, { ...summaries.get(1)!, mechanics: block(100) });
    withMechanics.set(3, { ...summaries.get(3)!, mechanics: block(250) });
    const night = nightSummary(fights, withMechanics);
    const row = night.mechanics?.rows.find((entry) => entry.spell_id === 331415);
    expect(row?.pulls_hit).toBe(2);
    expect(row?.players?.[0]).toMatchObject({ guid: 'Player-1', hits: 2, damage: 350, pulls: 2 });
    expect(night.mechanics?.table_found).toBe(true);
  });
});
