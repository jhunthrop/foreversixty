// web/src/lib/report/compare.test.ts
import { describe, expect, it } from 'vitest';
import { abilityDiff, metricTable, playerAbilityDiff } from './compare';
import type { Ability, Actor, Summary } from './types';

function ability(spell_id: number, name: string, effective: number, via?: string): Ability {
  return { spell_id, name, via, total: effective, effective, hits: 1, crits: 0, ticks: 0, min: 0, max: 0 };
}

function actor(guid: string, abilities: Ability[]): Actor {
  return {
    guid,
    name: guid,
    total: abilities.reduce((sum, a) => sum + a.total, 0),
    effective: abilities.reduce((sum, a) => sum + a.effective, 0),
    active_ms: 1000,
    abilities,
    targets: [],
    series: [],
  };
}

function summary(damage: Actor[], healing: Actor[] = []): Summary {
  return {
    engine_version: 'test',
    fight_index: 1,
    duration_ms: 10_000,
    damage_done: damage,
    damage_taken: [],
    healing,
    healing_taken: [],
    deaths: [],
    auras: [],
    casts: [],
    interrupts: [],
    dispels: [],
    resources: [],
    threat: [],
    combatants: [],
    roster: [],
  };
}

describe('metricTable', () => {
  it('sends each metric to the table its ability split lives in, and threat to none', () => {
    expect(metricTable('damage_done')).toBe('damage_done');
    expect(metricTable('dps')).toBe('damage_done');
    expect(metricTable('healing_done')).toBe('healing');
    expect(metricTable('hps')).toBe('healing');
    expect(metricTable('damage_taken')).toBe('damage_taken');
    expect(metricTable('dtps')).toBe('damage_taken');
    expect(metricTable('threat')).toBeNull();
    expect(metricTable('tps')).toBeNull();
  });
});

describe('abilityDiff', () => {
  const left = summary([actor('P1', [ability(1, 'Slam', 400), ability(2, 'Cleave', 100)])]);
  const right = summary([actor('P1', [ability(1, 'Slam', 250), ability(3, 'Execute', 300)])]);

  it('lists both pulls’ abilities with a signed difference, biggest first', () => {
    const rows = abilityDiff(left, right, 'P1', 'dps');
    expect(rows.map((row) => [row.name, row.a, row.b, row.delta])).toEqual([
      ['Execute', null, 300, -300],
      ['Slam', 400, 250, 150],
      ['Cleave', 100, null, 100],
    ]);
  });

  it('keeps a pet’s ability apart from its owner’s own of the same spell', () => {
    const owner = summary([actor('P1', [ability(1, 'Melee', 100), ability(1, 'Melee', 60, 'Ashfang')])]);
    const rows = abilityDiff(owner, summary([actor('P1', [])]), 'P1', 'dps');
    expect(rows.map((row) => [row.name, row.via, row.a])).toEqual([
      ['Melee', undefined, 100],
      ['Melee', 'Ashfang', 60],
    ]);
  });

  it('has nothing to split for threat, which has no per-ability table', () => {
    expect(abilityDiff(left, right, 'P1', 'threat')).toEqual([]);
  });

  it('is empty rather than throwing when a side is missing or the player is not in it', () => {
    expect(abilityDiff(left, null, 'P1', 'dps')).toEqual([
      { key: '1|', name: 'Slam', via: undefined, a: 400, b: null, delta: 400 },
      { key: '2|', name: 'Cleave', via: undefined, a: 100, b: null, delta: 100 },
    ]);
    expect(abilityDiff(left, right, 'nobody', 'dps')).toEqual([]);
  });

  it('reads healing off the healing table, not the damage one', () => {
    const healers = summary([], [actor('P2', [ability(9, 'Heal', 900)])]);
    expect(abilityDiff(healers, summary([], []), 'P2', 'hps').map((row) => row.name)).toEqual(['Heal']);
    expect(abilityDiff(healers, summary([], []), 'P2', 'dps')).toEqual([]);
  });
});

describe('playerAbilityDiff', () => {
  it('puts two players of one fight side by side', () => {
    const one = summary([
      actor('P1', [ability(1, 'Slam', 400)]),
      actor('P2', [ability(5, 'Frostbolt', 310)]),
    ]);
    expect(
      playerAbilityDiff(one, 'P1', 'P2', 'dps').map((row) => [row.name, row.a, row.b, row.delta]),
    ).toEqual([
      ['Slam', 400, null, 400],
      ['Frostbolt', null, 310, -310],
    ]);
  });
});
