import { describe, expect, it } from 'vitest';
import { inSource, scopeSource } from './source';
import type { Summary } from './types';

const players = new Set(['Player-1', 'Player-2']);

function summary(): Summary {
  return {
    engine_version: 'test',
    fight_index: 1,
    duration_ms: 10_000,
    damage_done: [
      {
        guid: 'Player-1',
        name: 'One',
        total: 1,
        effective: 1,
        active_ms: 1,
        abilities: [],
        targets: [],
        series: [1],
      },
      {
        guid: 'Player-2',
        name: 'Two',
        total: 1,
        effective: 1,
        active_ms: 1,
        abilities: [],
        targets: [],
        series: [1],
      },
      {
        guid: 'Creature-9',
        name: 'Boss',
        total: 1,
        effective: 1,
        active_ms: 1,
        abilities: [],
        targets: [],
        series: [1],
      },
    ],
    damage_taken: [],
    healing: [],
    healing_taken: [],
    deaths: [
      { guid: 'Player-1', name: 'One', at_ms: 5000, last: [], auras_held: [], auras_lost: [] },
      { guid: 'Player-2', name: 'Two', at_ms: 6000, last: [], auras_held: [], auras_lost: [] },
    ] as Summary['deaths'],
    auras: [
      {
        target_guid: 'Player-1',
        target_name: 'One',
        spell_id: 1,
        name: 'A',
        type: 'BUFF',
        applications: 1,
        max_stacks: 1,
        uptime_ms: 1,
        segments: [],
        appliers: [],
      },
      {
        target_guid: 'Creature-9',
        target_name: 'Boss',
        spell_id: 2,
        name: 'B',
        type: 'DEBUFF',
        applications: 1,
        max_stacks: 1,
        uptime_ms: 1,
        segments: [],
        appliers: ['Player-2'],
      },
    ] as Summary['auras'],
    casts: [
      {
        guid: 'Player-1',
        name: 'One',
        owner_guid: 'Player-1',
        spell_id: 1,
        spell_name: 'Slam',
        started: 1,
        succeeded: 1,
        failed: 0,
        cast_time_ms: 0,
        sequence: [1000],
      },
      {
        guid: 'Pet-7',
        name: 'Ashfang',
        owner_guid: 'Player-2',
        spell_id: 2,
        spell_name: 'Bite',
        started: 1,
        succeeded: 1,
        failed: 0,
        cast_time_ms: 0,
        sequence: [2000],
      },
      {
        guid: 'Creature-9',
        name: 'Boss',
        owner_guid: 'Creature-9',
        spell_id: 3,
        spell_name: 'Cleave',
        started: 1,
        succeeded: 1,
        failed: 0,
        cast_time_ms: 0,
        sequence: [3000],
      },
    ],
    interrupts: [],
    dispels: [],
    resources: [],
    threat: [],
    threat_by_target: [
      { guid: 'Player-1', name: 'One', target_guid: 'Creature-9', target_name: 'Boss', threat: 10 },
      { guid: 'Player-2', name: 'Two', target_guid: 'Creature-9', target_name: 'Boss', threat: 20 },
    ],
    taunts: [
      {
        at_ms: 1000,
        source_guid: 'Player-1',
        source_name: 'One',
        target_guid: 'Creature-9',
        target_name: 'Boss',
        spell_id: 355,
        spell_name: 'Taunt',
      },
    ],
    combatants: [],
    roster: [
      {
        guid: 'Player-1',
        name: 'One',
        role: 'dps',
        active_ms: 1,
        activity_pct: 1,
        deaths: 1,
        damage_done: 1,
        healing_done: 0,
        damage_taken: 0,
        dps: 1,
        hps: 0,
        dtps: 0,
      },
      {
        guid: 'Player-2',
        name: 'Two',
        role: 'dps',
        active_ms: 1,
        activity_pct: 1,
        deaths: 1,
        damage_done: 1,
        healing_done: 0,
        damage_taken: 0,
        dps: 1,
        hps: 0,
        dtps: 0,
      },
    ],
  };
}

describe('inSource', () => {
  it('reads friendlies, enemies and one unit', () => {
    expect(inSource('Player-1', 'friendlies', players)).toBe(true);
    expect(inSource('Creature-9', 'friendlies', players)).toBe(false);
    expect(inSource('Creature-9', 'enemies', players)).toBe(true);
    expect(inSource('Player-2', 'Player-1', players)).toBe(false);
    expect(inSource('Player-1', 'Player-1', players)).toBe(true);
  });

  it('treats everyone as friendly when the player list is unknown', () => {
    expect(inSource('Player-1', 'friendlies', new Set())).toBe(true);
  });
});

describe('scopeSource', () => {
  it('cuts enemies out of the unit tables under the default scope, and leaves the roster alone', () => {
    const scoped = scopeSource(summary(), 'friendlies', players);
    expect(scoped.roster.map((row) => row.guid)).toEqual(['Player-1', 'Player-2']);
    expect(scoped.auras.map((track) => track.target_guid)).toEqual(['Player-1']);
  });

  it('returns the same summary when the player list is unknown', () => {
    const input = summary();
    expect(scopeSource(input, 'friendlies', new Set())).toBe(input);
  });

  it('narrows every unit-keyed table to the chosen player', () => {
    const scoped = scopeSource(summary(), 'Player-1', players);
    expect(scoped.roster.map((row) => row.guid)).toEqual(['Player-1']);
    expect(scoped.damage_done.map((actor) => actor.guid)).toEqual(['Player-1']);
    expect(scoped.deaths.map((death) => death.guid)).toEqual(['Player-1']);
    expect(scoped.auras.map((track) => track.target_guid)).toEqual(['Player-1']);
    expect(scoped.casts.map((row) => row.spell_name)).toEqual(['Slam']);
  });

  it('keeps only enemies for the enemies scope', () => {
    const scoped = scopeSource(summary(), 'enemies', players);
    expect(scoped.roster).toEqual([]);
    expect(scoped.auras.map((track) => track.target_guid)).toEqual(['Creature-9']);
  });

  it('keeps the debuffs a player applied on enemies under that player’s scope', () => {
    const scoped = scopeSource(summary(), 'Player-2', players);
    expect(scoped.auras.map((track) => track.target_guid)).toEqual(['Creature-9']);
  });

  it('narrows per-target threat by the player and taunts by their source', () => {
    const scoped = scopeSource(summary(), 'Player-1', players);
    expect(scoped.threat_by_target?.map((pair) => pair.guid)).toEqual(['Player-1']);
    expect(scoped.taunts?.map((taunt) => taunt.source_guid)).toEqual(['Player-1']);
  });

  it('keeps every player’s per-target threat under the friendlies scope', () => {
    const scoped = scopeSource(summary(), 'friendlies', players);
    expect(scoped.threat_by_target?.map((pair) => pair.guid)).toEqual(['Player-1', 'Player-2']);
  });
});

describe('casts under the source scope', () => {
  it('keeps a pet’s casts under the player who owns it', () => {
    const scoped = scopeSource(summary(), 'Player-2', players);
    expect(scoped.casts.map((row) => row.spell_name)).toEqual(['Bite']);
  });

  it('shows a pet’s casts under all friendlies, where the pet itself is not a player', () => {
    const scoped = scopeSource(summary(), 'friendlies', players);
    expect(scoped.casts.map((row) => row.spell_name).sort()).toEqual(['Bite', 'Slam']);
  });

  it('leaves a pet out of the enemies, because its owner is a friendly', () => {
    const scoped = scopeSource(summary(), 'enemies', players, new Set([...players, 'Pet-7']));
    expect(scoped.casts.map((row) => row.spell_name)).toEqual(['Cleave']);
  });

  it('falls back to the caster’s own guid on a report parsed before owners were kept', () => {
    const old = summary();
    old.casts = old.casts.map((row) => ({ ...row, owner_guid: undefined }));
    expect(scopeSource(old, 'Player-2', players).casts).toEqual([]);
  });
});
