// web/src/lib/report/compare.test.ts
import { describe, expect, it } from 'vitest';
import {
  abilityDiff,
  metricTable,
  phaseWindow,
  playerAbilityDiff,
  sharedPhases,
  sideScaled,
} from './compare';
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

  it('names the id when two rows share a name, and only then', () => {
    // A Mistweaver's two Essence Fonts: one name, two spell ids, nothing else to tell
    // them apart -- the Healing tab shows the id here and Compare now does too.
    const fonts = summary(
      [],
      [
        actor('P2', [
          ability(191840, 'Essence Font', 12_995),
          ability(344006, 'Essence Font', 2853),
          ability(115175, 'Soothing Mist', 900),
        ]),
      ],
    );
    expect(abilityDiff(fonts, summary([], []), 'P2', 'hps').map((row) => [row.name, row.id])).toEqual([
      ['Essence Font', 191840],
      ['Essence Font', 344006],
      ['Soothing Mist', undefined],
    ]);
  });

  it('leaves a pet’s copy of one spell without an id: the pet’s name already tells them apart', () => {
    const owner = summary([actor('P1', [ability(1, 'Melee', 100), ability(1, 'Melee', 60, 'Ashfang')])]);
    const rows = abilityDiff(owner, summary([actor('P1', [])]), 'P1', 'dps');
    expect(rows.map((row) => row.id)).toEqual([undefined, undefined]);
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

describe('aligning two pulls by phase', () => {
  const phased = (phases: { name: string; start_ms: number; end_ms: number }[]): Summary => ({
    ...summary([]),
    phases,
  });
  const left = phased([
    { name: 'Phase 1', start_ms: 0, end_ms: 20_000 },
    { name: 'Phase 2', start_ms: 20_000, end_ms: 90_000 },
  ]);
  const right = phased([
    { name: 'Phase 1', start_ms: 0, end_ms: 35_000 },
    { name: 'Phase 2', start_ms: 35_000, end_ms: 60_000 },
    { name: 'Phase 3', start_ms: 60_000, end_ms: 120_000 },
  ]);

  it('offers only the phases both pulls reached, in the first pull’s order', () => {
    expect(sharedPhases(left, right)).toEqual(['Phase 1', 'Phase 2']);
  });

  it('gives each side its own span for the named phase', () => {
    expect(phaseWindow(left, 'Phase 2')).toEqual({ startMs: 20_000, endMs: 90_000 });
    expect(phaseWindow(right, 'Phase 2')).toEqual({ startMs: 35_000, endMs: 60_000 });
  });

  it('has nothing to offer when a side has no phases at all', () => {
    expect(sharedPhases(left, phased([]))).toEqual([]);
    expect(sharedPhases(left, null)).toEqual([]);
    expect(phaseWindow(phased([]), 'Phase 2')).toBeNull();
    expect(phaseWindow(null, 'Phase 2')).toBeNull();
  });
});

describe('sideScaled', () => {
  const loaded = summary([]);
  const aWindow = { startMs: 0, endMs: 20_000 };

  it('is false with no window: the side reads its whole fight, measured rather than prorated', () => {
    expect(sideScaled(loaded, null)).toBe(false);
  });

  it('is false with no summary: a side not yet loaded has nothing to prorate', () => {
    expect(sideScaled(null, aWindow)).toBe(false);
  });

  it('is true only for a loaded side narrowed by a window, phase-picked or otherwise', () => {
    expect(sideScaled(loaded, aWindow)).toBe(true);
  });
});
