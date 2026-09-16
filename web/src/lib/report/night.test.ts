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

  it('leaves the night\u2019s mechanics undefined when no pull carried a mechanics block', () => {
    // A report parsed before engine 0.3.0 has no block at all, which is a different
    // answer from "these bosses have no table" and must not be flattened into it.
    expect(nightSummary(fights, summaries).mechanics).toBeUndefined();
  });

  it('folds mechanics per boss across pulls, keeping two bosses\u2019 rows for one spell apart', () => {
    // Two bosses whose tables both list spell 331415. Merged by spell id alone, Kryxis\u2019s
    // hit would be added to Kaal\u2019s row and the night would report one mechanic on three
    // pulls of a boss that was pulled twice.
    const byEncounter = [
      { ...fights[0], encounter_id: 2363 },
      fights[1],
      { ...fights[2], encounter_id: 2363 },
      { ...fights[3], encounter_id: 2360 },
    ];
    const block = (damage: number, firstMs: number, lastMs: number, casts: number) => ({
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
              first_ms: firstMs,
              last_ms: lastMs,
              killed: false,
            },
          ],
        },
        { spell_id: 5, name: 'Corrupted Blood', kind: 'interrupt' as const, casts, stopped: 1 },
      ],
    });
    const withMechanics = new Map(summaries);
    withMechanics.set(1, { ...summaries.get(1)!, mechanics: block(100, 4000, 6000, 3) });
    withMechanics.set(3, { ...summaries.get(3)!, mechanics: block(250, 1000, 9000, 2) });
    withMechanics.set(4, { ...summary(4, 30_000, [roster('A', 100)]), mechanics: block(900, 500, 500, 1) });
    const night = nightSummary(byEncounter, withMechanics);

    expect(night.mechanics?.table_found).toBe(true);
    expect(night.mechanics?.rows).toHaveLength(4);

    const kaal = night.mechanics?.rows.find((row) => row.spell_id === 331415 && row.encounter_id === 2363);
    expect(kaal).toMatchObject({ encounter: 'Kaal', pulls_hit: 2 });
    // Pull 1 runs 0\u201310_000 ms and pull 3 10_000\u201330_000, so the earliest hit is pull 1\u2019s
    // 4000 and the latest is pull 3\u2019s 9000 shifted by pull 3\u2019s 10_000 ms offset.
    expect(kaal?.players?.[0]).toMatchObject({
      guid: 'Player-1',
      hits: 2,
      damage: 350,
      first_ms: 4000,
      last_ms: 19_000,
      pulls: 2,
    });
    expect(
      night.mechanics?.rows.find((row) => row.spell_id === 5 && row.encounter_id === 2363),
    ).toMatchObject({ casts: 5, stopped: 2, pulls_hit: 2 });

    const kryxis = night.mechanics?.rows.find((row) => row.spell_id === 331415 && row.encounter_id === 2360);
    expect(kryxis).toMatchObject({ encounter: 'Kryxis', pulls_hit: 1 });
    expect(kryxis?.players?.[0]).toMatchObject({ hits: 1, damage: 900, pulls: 1 });

    // The denominator behind "hit someone on 2 of 2 pulls": the pulls of each boss the
    // fold actually loaded, in the order the night met them.
    expect(night.mechanics?.bosses).toEqual([
      { encounter_id: 2363, name: 'Kaal', pulls: 2 },
      { encounter_id: 2360, name: 'Kryxis', pulls: 1 },
    ]);
  });

  it('folds per-target threat by enemy name and shifts taunts onto the night’s clock', () => {
    const kaalFights = [fight(1, 'Kaal', false, 10_000), fight(2, 'Kaal', true, 10_000)];
    const kaalSummaries = new Map<number, Summary>([
      [1, summary(1, 10_000, [roster('Player-1', 1000)])],
      [2, summary(2, 10_000, [roster('Player-1', 500)])],
    ]);
    const pairs = (threat: number) => [
      { guid: 'Player-1', name: 'Tank', target_guid: 'Creature-a', target_name: 'Kaal', threat },
    ];
    const withThreat = new Map(kaalSummaries);
    withThreat.set(1, {
      ...kaalSummaries.get(1)!,
      threat_by_target: pairs(100),
      taunts: [
        {
          at_ms: 1000,
          source_guid: 'Player-1',
          source_name: 'Tank',
          target_guid: 'Creature-a',
          target_name: 'Kaal',
          spell_id: 355,
          spell_name: 'Taunt',
        },
      ],
    });
    // A second pull of the same boss is a new GUID for the add; folding by name keeps it
    // as one row instead of two.
    withThreat.set(2, {
      ...kaalSummaries.get(2)!,
      threat_by_target: [{ ...pairs(50)[0], target_guid: 'Creature-b' }],
    });
    const night = nightSummary(kaalFights, withThreat);
    expect(night.threat_by_target).toEqual([
      { guid: 'Player-1', name: 'Tank', target_guid: 'Kaal', target_name: 'Kaal', threat: 150 },
    ]);
    expect(night.taunts?.[0].label).toMatch(/pull 1/);
    expect(night.taunts?.[0].at_ms).toBe(1000);
    // Pull 2's summary carries the key with nothing in it, so the night can say "no
    // taunts" rather than "parsed before taunts were kept".
    expect(nightSummary(kaalFights, kaalSummaries).taunts).toBeUndefined();
  });

  it('drops the per-second series from the night’s threat pairs: a night has no one clock', () => {
    const fights = [fight(3, 'Kaal', false, 10_000), fight(4, 'Kaal', true, 10_000)];
    const summaries = new Map([
      [3, summary(3, 10_000, [roster('A', 10)])],
      [4, summary(4, 10_000, [roster('A', 10)])],
    ]);
    const pair = (threat: number, series: number[]) => [
      {
        guid: 'Player-1',
        name: 'Tank',
        target_guid: 'Creature-a',
        target_name: 'Kaal',
        threat,
        series,
      },
    ];
    summaries.set(3, { ...summaries.get(3)!, threat_by_target: pair(100, [60, 40]) });
    summaries.set(4, { ...summaries.get(4)!, threat_by_target: pair(50, [50]) });
    const night = nightSummary(fights, summaries);
    expect(night.threat_by_target?.[0].threat).toBe(150);
    expect(night.threat_by_target?.[0].series).toBeUndefined();
  });

  it('offsets hits by their pull, skips a pull whose table was not found, and folds a no-player row by its counts', () => {
    const withMechanics = new Map(summaries);
    // Pull 1 (10_000 ms, offset 0): table not found, but the fixture still carries a
    // stray row -- it must be ignored entirely, not merged.
    withMechanics.set(1, {
      ...summaries.get(1)!,
      mechanics: {
        table_found: false,
        rows: [
          { spell_id: 331415, name: 'Wicked Gash', kind: 'avoidable', players: [] },
          { spell_id: 5, name: 'Corrupted Blood', kind: 'interrupt', casts: 99, stopped: 99 },
        ],
      },
    });
    // Pull 3 (20_000 ms, offset 10_000): a real table with an avoidable hit near the
    // start of the offset window and a no-player interrupt row.
    withMechanics.set(3, {
      ...summaries.get(3)!,
      mechanics: {
        table_found: true,
        rows: [
          {
            spell_id: 331415,
            name: 'Wicked Gash',
            kind: 'avoidable',
            players: [
              {
                guid: 'Player-1',
                name: 'Hobolol',
                hits: 1,
                damage: 250,
                first_ms: 500,
                last_ms: 1500,
                killed: false,
              },
            ],
          },
          { spell_id: 5, name: 'Corrupted Blood', kind: 'interrupt', casts: 3, stopped: 1 },
        ],
      },
    });
    const night = nightSummary(fights, withMechanics);

    expect(night.mechanics?.table_found).toBe(true);
    expect(night.mechanics?.rows).toHaveLength(2);
    // No encounter id on these fights, so the boss name keys the rows instead.
    expect(night.mechanics?.rows.every((entry) => entry.encounter === 'Kaal')).toBe(true);

    const avoidable = night.mechanics?.rows.find((entry) => entry.spell_id === 331415);
    // Pull 1's row was dropped with its pull (table_found: false), so this hit is only
    // pull 3's, shifted onto the night's clock by pull 3's 10_000 ms offset.
    expect(avoidable?.pulls_hit).toBe(1);
    expect(avoidable?.players?.[0]).toMatchObject({
      guid: 'Player-1',
      hits: 1,
      damage: 250,
      first_ms: 10_500,
      last_ms: 11_500,
      pulls: 1,
    });

    const interrupt = night.mechanics?.rows.find((entry) => entry.spell_id === 5);
    // Pull 1's 99/99 stray counts must not appear: only pull 3's 3/1 do.
    expect(interrupt).toMatchObject({ casts: 3, stopped: 1, pulls_hit: 1, players: undefined });
  });
});
