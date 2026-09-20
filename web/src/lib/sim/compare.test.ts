// web/src/lib/sim/compare.test.ts
import { describe, expect, it } from 'vitest';
import type { ActionNames } from './action-names';
import type { Summary } from '../report/types';
import { simCopy } from './copy';
import { compareSummaries, MAX_ABILITY_LINES, MAX_AURA_LINES } from './compare';

const names: ActionNames = {
  spell: { '25286': 'Heroic Strike', '23894': 'Bloodthirst', '12974': 'Flurry', '12292': 'Death Wish' },
  item: {},
};

const empty = {
  damage_taken: [],
  healing: [],
  healing_taken: [],
  deaths: [],
  interrupts: [],
  dispels: [],
  resources: [],
  threat: [],
  combatants: [],
  roster: [],
};

function ability(spellId: number, name: string, total: number) {
  return { spell_id: spellId, name, total, effective: total, hits: 1, crits: 0, ticks: 0, min: 1, max: 1 };
}

/**
 * The sim: 270,000 over 180 s is 1,500 DPS. Its melee is two tagged rows at derived ids,
 * exactly as sim/adapter emits them; everything else is a plain spell at its client id.
 */
function simSummary(): Summary {
  return {
    ...empty,
    engine_version: 'sim:test',
    fight_index: 1,
    duration_ms: 180_000,
    damage_done: [
      {
        guid: 'sim-player',
        name: 'Fury',
        class: 'warrior',
        total: 270_000,
        effective: 270_000,
        active_ms: 180_000,
        targets: [],
        series: [],
        abilities: [
          ability(25286, 'spell:25286', 120_000),
          ability(10022000007, 'other:attack/1', 60_000),
          ability(20022000007, 'other:attack/2', 30_000),
          ability(23894, 'spell:23894', 60_000),
        ],
      },
    ],
    casts: [castRow('sim-player', 25286, 'spell:25286', 41), castRow('sim-player', 23894, 'spell:23894', 22)],
    auras: [
      auraRow('sim-player', 12974, 'spell:12974', 140_400),
      auraRow('sim-player', 12292, 'spell:12292', 108_000),
    ],
  };
}

/** The same fight, played: 216,000 over 180 s is 1,200 DPS, which is 80% of the sim. */
function actualSummary(): Summary {
  return {
    ...empty,
    engine_version: '0.5.3',
    fight_index: 1,
    duration_ms: 180_000,
    damage_done: [
      {
        guid: 'log-player',
        name: 'Thrallgar',
        class: 'Warrior',
        total: 216_000,
        effective: 216_000,
        active_ms: 180_000,
        targets: [],
        series: [],
        abilities: [
          ability(25286, 'Heroic Strike', 96_000),
          // A combat log records one melee row and gives it spell id 0.
          ability(0, 'Melee', 72_000),
          ability(23894, 'Bloodthirst', 48_000),
        ],
      },
    ],
    casts: [
      castRow('log-player', 25286, 'Heroic Strike', 33),
      castRow('log-player', 23894, 'Bloodthirst', 19),
    ],
    auras: [
      auraRow('log-player', 12974, 'Flurry', 109_800),
      auraRow('log-player', 12292, 'Death Wish', 72_000),
    ],
  };
}

function castRow(guid: string, spellId: number, spellName: string, succeeded: number) {
  return {
    guid,
    name: guid,
    spell_id: spellId,
    spell_name: spellName,
    started: succeeded,
    succeeded,
    failed: 0,
    cast_time_ms: 0,
    sequence: [],
  };
}

function auraRow(guid: string, spellId: number, name: string, uptimeMs: number) {
  return {
    target_guid: guid,
    target_name: guid,
    spell_id: spellId,
    name,
    type: 'BUFF',
    applications: 1,
    max_stacks: 1,
    uptime_ms: uptimeMs,
    segments: [],
    appliers: [guid],
  };
}

describe('compareSummaries', () => {
  it('leads with the two DPS figures and what the fraction means', () => {
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', names);
    expect(comparison.headline).toBe(
      'This fight did 1,200 DPS; the sim expects 1,500. That is 80% of what this gear can do.',
    );
    expect(comparison.executionScore).toBeCloseTo(0.8, 6);
  });

  it('joins a plain spell by its client id, whatever the two sides call it', () => {
    // The sim row is named `spell:25286`, the fight row `Heroic Strike`, and both carry
    // spell id 25286 -- so they are one row, and the name shown is the resolved one.
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', names);
    const row = comparison.abilities.find((entry) => entry.name === 'Heroic Strike');
    expect(row).toEqual({
      // The row's identity is the client spell id, which is also its test id in CompareView.
      spellId: 25286,
      name: 'Heroic Strike',
      simCasts: 41,
      actualCasts: 33,
      simDamage: 120_000,
      actualDamage: 96_000,
    });
  });

  it('joins the sim’s tagged melee rows onto the log’s single one, by name', () => {
    // Every derived id is at or above SYNTHETIC_ID_BASE and means nothing to a combat log,
    // and the log's own melee row carries spell id 0 -- so neither side can be joined by
    // number here. The two tagged sim rows fold onto one, and their damage sums to 90,000.
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', names);
    const row = comparison.abilities.find((entry) => entry.name === 'Melee');
    expect(row?.simDamage).toBe(90_000);
    expect(row?.actualDamage).toBe(72_000);
    // A name-keyed row still gets an integer identity: the sim side's first row id.
    expect(row?.spellId).toBe(10022000007);
  });

  it('joins a tagged row of a real spell onto its plain row, by the client id they share', () => {
    // The real engine reports Heroic Strike twice -- "spell:25286" and the tagged
    // "spell:25286/1" at a derived id -- and the log reports it once. One row, damage and
    // casts summed, one line about it; a second row under the same name printed the same
    // sentence twice, which is a repeated {#each} key on the page.
    const sim = simSummary();
    sim.damage_done[0].abilities.push(ability(10002025286, 'spell:25286/1', 30_000));
    sim.damage_done[0].total += 30_000;
    sim.casts.push(castRow('sim-player', 10002025286, 'spell:25286/1', 3));
    const comparison = compareSummaries(sim, actualSummary(), 'Thrallgar', names);
    const rows = comparison.abilities.filter((row) => row.name === 'Heroic Strike');
    expect(rows).toHaveLength(1);
    expect(rows[0]?.spellId).toBe(25286);
    expect(rows[0]?.simDamage).toBe(150_000);
    expect(rows[0]?.simCasts).toBe(44);
    expect(new Set(comparison.lines).size).toBe(comparison.lines.length);
  });

  it('never joins two different actions by number just because both are id zero', () => {
    const sim = simSummary();
    const actual = actualSummary();
    // A second log row at spell id 0 -- a environmental hit, say -- must stay its own row.
    actual.damage_done[0].abilities.push(ability(0, 'Falling', 1));
    actual.damage_done[0].total += 1;
    const comparison = compareSummaries(sim, actual, 'Thrallgar', names);
    expect(comparison.abilities.map((row) => row.name)).toContain('Falling');
    expect(comparison.abilities.find((row) => row.name === 'Falling')?.simDamage).toBe(0);
    // Every row's identity is distinct, which is what CompareView's {#each} key needs.
    const ids = comparison.abilities.map((row) => row.spellId);
    expect(new Set(ids).size).toBe(ids.length);
  });

  it('pairs every ability the two runs have between them, biggest first', () => {
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', names);
    expect(comparison.abilities.map((row) => row.name)).toEqual(['Heroic Strike', 'Melee', 'Bloodthirst']);
  });

  it('explains each cast-count gap in words, biggest first', () => {
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', names);
    expect(comparison.lines[0]).toBe('Heroic Strike cast 33 times, the sim expects 41.');
    expect(comparison.lines).toContain('Bloodthirst cast 19 times, the sim expects 22.');
  });

  it('explains each uptime gap the way the design writes it', () => {
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', names);
    expect(comparison.lines).toContain('Flurry uptime 61% against 78%.');
    expect(comparison.lines).toContain('Death Wish uptime 40% against 60%.');
  });

  it('falls back to a humanised label rather than the raw key when the build has no name table', () => {
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', null);
    expect(comparison.abilities.map((row) => row.name)).toContain('Spell 25286');
  });

  it('clamps the score to the range the contract stores', () => {
    const actual = actualSummary();
    actual.damage_done[0].total *= 9;
    expect(compareSummaries(simSummary(), actual, 'Thrallgar', names).executionScore).toBe(2);
  });

  it('says so when the fight has no damage for this character', () => {
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Someone Else', names);
    expect(comparison.headline).toBe(simCopy.compareNoPlayer);
    expect(comparison.executionScore).toBeNull();
    expect(comparison.lines).toEqual([]);
  });

  it('lists no more than five ability lines and three aura lines', () => {
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', names);
    expect(comparison.lines.length).toBeLessThanOrEqual(MAX_ABILITY_LINES + MAX_AURA_LINES);
  });

  it('gives every aura row a stable id, not just a name', () => {
    const comparison = compareSummaries(simSummary(), actualSummary(), 'Thrallgar', names);
    const flurry = comparison.auras.find((row) => row.name === 'Flurry');
    expect(flurry?.spellId).toBe(12974);
  });

  it('keeps two same-named auras at different ids apart, each with its own identity', () => {
    // Two real ranks of the same buff can resolve to the same display name -- one on the
    // sim side, a different one logged -- and both are below SYNTHETIC_ID_BASE, so tier 1
    // keys each by its own id and neither folds onto the other. CompareView.svelte's
    // `{#each}` keys this table on `spellId`, and a Svelte 5 duplicate key throws at
    // runtime, so this is the regression the aura table's own id field exists to prevent.
    const sim: Summary = {
      ...empty,
      engine_version: 'sim:test',
      fight_index: 1,
      duration_ms: 100_000,
      damage_done: [
        {
          guid: 'sim-player',
          name: 'Fury',
          class: 'warrior',
          total: 100_000,
          effective: 100_000,
          active_ms: 100_000,
          targets: [],
          series: [],
          abilities: [ability(1, 'Slam', 100_000)],
        },
      ],
      casts: [],
      auras: [
        auraRow('sim-player', 100, 'Renewed Vigor', 50_000),
        auraRow('sim-player', 200, 'Renewed Vigor', 30_000),
      ],
    };
    const actual: Summary = {
      ...empty,
      engine_version: '0.5.3',
      fight_index: 1,
      duration_ms: 100_000,
      damage_done: [
        {
          guid: 'log-player',
          name: 'Thrallgar',
          class: 'Warrior',
          total: 90_000,
          effective: 90_000,
          active_ms: 100_000,
          targets: [],
          series: [],
          abilities: [ability(1, 'Slam', 90_000)],
        },
      ],
      casts: [],
      auras: [auraRow('log-player', 100, 'Renewed Vigor', 40_000)],
    };

    const comparison = compareSummaries(sim, actual, 'Thrallgar', names);
    const renewedVigor = comparison.auras.filter((row) => row.name === 'Renewed Vigor');
    expect(renewedVigor).toHaveLength(2);
    expect(renewedVigor.map((row) => row.spellId).sort((a, b) => a - b)).toEqual([100, 200]);
    // Every row's identity is distinct, which is what CompareView's {#each} key needs.
    const ids = comparison.auras.map((row) => row.spellId);
    expect(new Set(ids).size).toBe(ids.length);
  });

  it('never mutates the summaries it compares', () => {
    const sim = simSummary();
    const actual = actualSummary();
    const simSnapshot = structuredClone(sim);
    const actualSnapshot = structuredClone(actual);

    compareSummaries(sim, actual, 'Thrallgar', names);

    expect(sim).toEqual(simSnapshot);
    expect(actual).toEqual(actualSnapshot);
  });
});
