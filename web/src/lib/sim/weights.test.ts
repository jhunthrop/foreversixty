// web/src/lib/sim/weights.test.ts
import { describe, expect, it } from 'vitest';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import specsJson from '../../fixtures/sim/specs.json';
import {
  CASTER_REFERENCE,
  DEFAULT_REFERENCE,
  defaultStatsFor,
  fallbackReferenceFor,
  formatWeightError,
  hasWeightStats,
  isDpsSpec,
  isSignificant,
  pawnString,
  pickableStatsFor,
  referenceFor,
  statLabel,
  weightScale,
  weightStatsFor,
  weightsEngineIterations,
  weightsIterationsFor,
  WEIGHTS_BROWSER_DEFAULT_ITERATIONS,
  WEIGHTS_ITERATIONS_FACTOR,
  WEIGHT_STATS,
} from './weights';
import { PRECISION_ITERATIONS } from './precision';
import { specRow } from './spec-label';
import { SPECS } from './specs';
import { WEIGHT_ERROR_BELOW_THRESHOLD } from './copy';
import type { StatWeight, WeightsResult } from './bulk-types';
import type { SpecFidelity } from './types';

const result = weightsResultJson as unknown as WeightsResult;
const specs = specsJson as unknown as SpecFidelity[];

describe('referenceFor', () => {
  it('takes the spec’s own reference stat from GET /v1/specs', () => {
    expect(referenceFor('warrior-fury', specs)).toBe('attack_power');
    expect(referenceFor('mage-frost', specs)).toBe('spell_power');
  });

  it('falls back to the pinned default for the kind of spec, not to one number', () => {
    // Contract 10.8: attack_power for melee and hunters, spell_power for casters.
    expect(referenceFor('warrior-fury', [])).toBe(DEFAULT_REFERENCE);
    expect(referenceFor('mage-fire', [])).toBe(CASTER_REFERENCE);
    expect(referenceFor('druid-balance', specs)).toBe(CASTER_REFERENCE);
    expect(fallbackReferenceFor('hunter-marksmanship')).toBe(DEFAULT_REFERENCE);
    expect(fallbackReferenceFor('nonesuch-spec')).toBe(DEFAULT_REFERENCE);
  });
});

describe('defaultStatsFor', () => {
  it('always includes the reference stat, exactly once, first', () => {
    const stats = defaultStatsFor('warrior-fury', 'attack_power');
    expect(stats[0]).toBe('attack_power');
    expect(stats.filter((stat) => stat === 'attack_power')).toHaveLength(1);
    expect(new Set(stats).size).toBe(stats.length);
  });

  it('is scoped to weightStats when the spec sent one, the reference stat still first', () => {
    const stats = defaultStatsFor('warrior-fury', 'attack_power', ['attack_power', 'strength', 'agility']);
    expect(stats).toEqual(['attack_power', 'strength', 'agility']);
  });
});

describe('statLabel', () => {
  it('names every stat the picker offers', () => {
    for (const stat of WEIGHT_STATS) expect(statLabel(stat.id)).toBe(stat.label);
  });

  // 2026-09-21 result-page review round 3: this used to return the bare id unchanged for
  // anything BY_ID had no row for, which is exactly how `frost_power` leaked onto the page
  // as itself -- PAWN_KEYS had no entry for any school-specific power stat. Humanised, never
  // the raw id, is the same last-resort stats.ts's own statLabel already uses.
  it('humanises an id neither BY_ID nor simCopy.statLabel carries, rather than returning it bare', () => {
    expect(statLabel('nonesuch')).toBe('Nonesuch');
    expect(statLabel('some_unlisted_stat')).toBe('Some unlisted stat');
  });
});

describe('weightScale', () => {
  it('is the largest weight plus its error, so no bar overflows its track', () => {
    // The widest row computed from the fixture itself, not a copied literal: a regenerated
    // fixture's own numbers change every time the engine reruns it, and a hand-copied
    // number here would silently drift from what the fixture actually carries.
    const widest = Math.max(...result.weights.map((row) => row.weight + row.error));
    expect(weightScale(result.weights)).toBeCloseTo(widest, 6);
    expect(weightScale([])).toBe(1);
  });
});

describe('pawnString', () => {
  it('is a Pawn v1 line with the class, the spec and two decimals a piece', () => {
    // The fixture's own "agility" row is insignificant -- left out here the same way a
    // stat Pawn has no key for is, below.
    expect(pawnString('warrior-fury', result.weights)).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury, Strength=1.50, AttackPower=1.00, ' +
        'CritRating=7.03, HitRating=6.87, HasteRating=5.85 )',
    );
  });

  it('leaves out a stat Pawn has no key for rather than inventing one', () => {
    expect(pawnString('warrior-fury', [{ stat: 'nonesuch', weight: 3, error: 0 }])).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury )',
    );
  });

  it('D45: leaves out a weight the engine flagged insignificant', () => {
    const weights: StatWeight[] = [
      { stat: 'attack_power', weight: 1, error: 0 },
      { stat: 'crit', weight: 14.37, error: 22.55, insignificant: true },
    ];
    expect(pawnString('warrior-fury', weights)).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury, AttackPower=1.00 )',
    );
  });

  it('D46: never disagrees with isSignificant about which rows are included -- the regression guard', () => {
    // The table greys a row exactly when `isSignificant` is false (StatWeights.svelte,
    // SavedWeights.svelte); pawnString must include exactly the rows that predicate calls
    // significant, and no others, for every row that has a Pawn key at all. Asserted
    // generically -- over the whole pinned vocabulary, not one hand-picked stat -- so this
    // guard cannot be satisfied by coincidence the way a single fixed-string test could be.
    const weights: StatWeight[] = WEIGHT_STATS.map((stat, index) => ({
      stat: stat.id,
      weight: index + 1,
      error: 0.1,
      insignificant: index % 2 === 0,
    }));
    const pawn = pawnString('warrior-fury', weights);
    for (const row of weights) {
      const key = WEIGHT_STATS.find((stat) => stat.id === row.stat)?.pawn ?? '';
      if (key === '') continue; // Pawn has no key for this stat regardless of significance.
      const included = pawn.includes(`${key}=`);
      expect(included, `${row.stat} (insignificant: ${row.insignificant})`).toBe(isSignificant(row));
    }
  });
});

describe('isSignificant', () => {
  it('treats absent or false insignificant as significant', () => {
    expect(isSignificant({ insignificant: undefined })).toBe(true);
    expect(isSignificant({ insignificant: false })).toBe(true);
    expect(isSignificant({})).toBe(true);
  });

  it('treats insignificant: true as not significant', () => {
    expect(isSignificant({ insignificant: true })).toBe(false);
  });
});

describe('pickableStatsFor (sub-item 4: only the spec’s own stats)', () => {
  it('offers the full pinned vocabulary when the engine sent no weight_stats', () => {
    expect(pickableStatsFor(undefined)).toBe(WEIGHT_STATS);
    expect(pickableStatsFor([])).toBe(WEIGHT_STATS);
  });

  it('offers only the spec’s own stats, in the engine’s order, when it sent one', () => {
    const pickable = pickableStatsFor(['attack_power', 'strength']);
    expect(pickable.map((stat) => stat.id)).toEqual(['attack_power', 'strength']);
    // Retail-only entries (D45) are never offered once the engine names its own list --
    // by construction (they are simply not in the list), not by a deny-list here.
    for (const wrong of ['expertise', 'spell_haste', 'armor_penetration', 'mp5', 'feral_attack_power']) {
      expect(pickable.map((stat) => stat.id)).not.toContain(wrong);
    }
  });
});

// 2026-09-21 result-page review round 3: a Frost Mage's "Stats to weigh" checklist and
// results table showed a row literally labelled `frost_power` -- PAWN_KEYS (weights.ts) had
// no entry for any school-specific power stat, so pickableStatsFor's own fallback handed
// back the raw wire id as its label. This walks every spec specs.ts carries (the client's
// own mirror of sim/specs/specs.go) and asserts every id its own weight_stats names comes
// back with a real, human label -- never the snake_case id itself.
describe('every spec’s pickable stats have a real label, never the raw id', () => {
  const SNAKE_CASE = /^[a-z]+(_[a-z]+)*$/;

  for (const spec of SPECS) {
    it(`${spec.spec}: every weight_stats id resolves to a label, not its own key`, () => {
      const pickable = pickableStatsFor(spec.weight_stats);
      for (const stat of pickable) {
        expect(stat.label, `${spec.spec}'s ${stat.id} has no real label`).not.toBe(stat.id);
        // A label that still happens to be snake_case (rather than failing to differ from
        // the id at all) is the same leak in a different shape -- catches a stat whose id
        // and intended label would coincidentally differ only by casing.
        expect(
          SNAKE_CASE.test(stat.label),
          `${spec.spec}'s ${stat.id} label "${stat.label}" still reads snake_case`,
        ).toBe(false);
      }
    });
  }
});

describe('hasWeightStats (final whole-branch review, Finding 3: the note must not disagree with the picker)', () => {
  it('is false for undefined and for an empty list -- the exact two shapes pickableStatsFor falls back on', () => {
    expect(hasWeightStats(undefined)).toBe(false);
    expect(hasWeightStats([])).toBe(false);
  });

  it('is true for any non-empty list', () => {
    expect(hasWeightStats(['attack_power'])).toBe(true);
  });
});

describe('weightStatsFor', () => {
  // Final whole-branch review, Finding 2: `GET /v1/specs` carries no `weight_stats` column
  // today (api/internal/sims/specs.go's own SpecFidelity has no such field) -- only the test
  // fixture used to hand-carry it, which hid the fact that the picker's spec-scoping shipped
  // inert. The real source of the list is the client's own generated `specs.ts`
  // (data/curated/specs.json), the same table `specRow`/`isDpsSpec` already read.
  it('falls back to the spec’s own curated weight_stats (specs.ts) when the API row carries no column -- the shape GET /v1/specs sends today', () => {
    expect(weightStatsFor('warrior-fury', specs)).toEqual(specRow('warrior-fury')?.weight_stats);
    expect(weightStatsFor('mage-frost', specs)).toEqual(specRow('mage-frost')?.weight_stats);
    // Neither is empty -- otherwise this test would not distinguish the curated fallback
    // from the "nobody has heard of this spec" case below.
    expect(specRow('warrior-fury')?.weight_stats.length).toBeGreaterThan(0);
    expect(specRow('mage-frost')?.weight_stats.length).toBeGreaterThan(0);
  });

  it('prefers an API-supplied weight_stats over the curated list, if the API ever sends one', () => {
    const apiRow: SpecFidelity = {
      ...specs[0],
      spec: 'warrior-fury',
      weight_stats: ['attack_power', 'strength'],
    };
    expect(weightStatsFor('warrior-fury', [apiRow])).toEqual(['attack_power', 'strength']);
  });

  it('is undefined for a spec neither the API nor the curated list has a row for', () => {
    expect(weightStatsFor('nonesuch-spec', specs)).toBeUndefined();
    expect(specRow('nonesuch-spec')).toBeNull();
  });
});

describe('formatWeightError (final whole-branch review, Finding 4: a non-zero error must never print as zero)', () => {
  it('keeps two decimals for a value that already reads clearly at that precision', () => {
    expect(formatWeightError(0.06)).toBe('0.06');
    expect(formatWeightError(22.55)).toBe('22.55');
  });

  it('reads a genuine zero as "0.00", the same width as every other row', () => {
    expect(formatWeightError(0)).toBe('0.00');
  });

  it('never rounds a non-zero error away to "0.00"', () => {
    expect(formatWeightError(0.001)).toBe(WEIGHT_ERROR_BELOW_THRESHOLD);
    expect(formatWeightError(0.001)).not.toBe('0.00');
  });
});

describe('isDpsSpec (sub-item 5: the honest refusal)', () => {
  it('is true for a dps spec', () => {
    expect(isDpsSpec('warrior-fury')).toBe(true);
  });

  it('is false for a healer or tank spec, and for an unrecognised one', () => {
    expect(isDpsSpec('druid-restoration')).toBe(false);
    expect(isDpsSpec('nonesuch-spec')).toBe(false);
  });
});

describe('the stat vocabulary (contract 10.8, pinned)', () => {
  it('carries one hit and one crit, and splits only haste', () => {
    const ids = WEIGHT_STATS.map((stat) => stat.id);
    expect(ids).toContain('hit');
    expect(ids).toContain('crit');
    expect(ids).toContain('melee_haste');
    expect(ids).toContain('spell_haste');
    for (const wrong of ['melee_hit', 'spell_hit', 'melee_crit', 'spell_crit', 'haste']) {
      expect(ids).not.toContain(wrong);
    }
  });

  it('spells MP5 as mp5', () => {
    const ids = WEIGHT_STATS.map((stat) => stat.id);
    expect(ids).toContain('mp5');
    expect(ids).not.toContain('m_p5');
  });

  it('offers only ids the pinned proto.Stat list carries', () => {
    const pinned = new Set([
      'strength',
      'agility',
      'stamina',
      'intellect',
      'spirit',
      'spell_power',
      'arcane_power',
      'fire_power',
      'frost_power',
      'holy_power',
      'nature_power',
      'shadow_power',
      'mp5',
      'hit',
      'crit',
      'spell_haste',
      'spell_penetration',
      'attack_power',
      'melee_haste',
      'armor_penetration',
      'expertise',
      'mana',
      'energy',
      'rage',
      'armor',
      'ranged_attack_power',
      'defense',
      'block',
      'block_value',
      'dodge',
      'parry',
      'health',
      'arcane_resistance',
      'fire_resistance',
      'frost_resistance',
      'nature_resistance',
      'shadow_resistance',
      'bonus_armor',
      'healing_power',
      'spell_damage',
      'feral_attack_power',
    ]);
    for (const stat of WEIGHT_STATS) expect(pinned.has(stat.id), stat.id).toBe(true);
  });
});

describe('weightsEngineIterations (Task 8, sub-item 2): the real engine cost, not the wire’s nominal count', () => {
  it('matches contract 10.9’s own formula: a baseline pass plus a low and a high pass per stat, each at Iterations * factor / 2', () => {
    expect(WEIGHTS_ITERATIONS_FACTOR).toBe(8);
    // The exact figure this task's own brief named for an eight-stat spec at "normal":
    // (3000 * 8 / 2) * (1 + 2*8) = 12,000 * 17 = 204,000 -- 68x the wire's own "3,000".
    expect(weightsEngineIterations(PRECISION_ITERATIONS.normal, 8)).toBe(204_000);
    expect(weightsEngineIterations(PRECISION_ITERATIONS.fast, 8)).toBe(34_000);
  });

  it('grows with both the iteration base and the stat count', () => {
    expect(weightsEngineIterations(500, 1)).toBeLessThan(weightsEngineIterations(500, 8));
    expect(weightsEngineIterations(500, 8)).toBeLessThan(weightsEngineIterations(3000, 8));
  });
});

describe('weightsIterationsFor (Task 8, sub-item 2): the browser lane is guarded, the server lane is not', () => {
  it('sends a smaller default on the browser lane than the plain-run PRECISION_ITERATIONS.fast', () => {
    expect(WEIGHTS_BROWSER_DEFAULT_ITERATIONS).toBeLessThan(PRECISION_ITERATIONS.fast);
    expect(weightsIterationsFor('fast', 'browser')).toBe(WEIGHTS_BROWSER_DEFAULT_ITERATIONS);
  });

  it('leaves the server lane at PRECISION_ITERATIONS, unchanged, at every precision', () => {
    for (const id of ['fast', 'normal', 'high'] as const) {
      expect(weightsIterationsFor(id, 'server')).toBe(PRECISION_ITERATIONS[id]);
    }
  });

  it('leaves normal and high unchanged on the browser lane too -- only the default is guarded, the rest is disclosed', () => {
    expect(weightsIterationsFor('normal', 'browser')).toBe(PRECISION_ITERATIONS.normal);
    expect(weightsIterationsFor('high', 'browser')).toBe(PRECISION_ITERATIONS.high);
  });

  // Measured, not guessed (task-8-report.md): a real wasm weights run for an eight-stat
  // spec at PRECISION_ITERATIONS.fast (34,000 real engine iterations) took 76.7 s in an
  // actual browser (Apple M4 Pro, one wasm worker) -- about 443 real iterations per second.
  // The guarded browser default must stay well inside a minute even at the worst case a
  // fallback spec can reach: the full pinned vocabulary, not just a narrowed eight stats.
  it('the guarded browser default finishes with real headroom under a minute, even at the full stat vocabulary', () => {
    const measuredIterationsPerSecond = 443;
    const worstCaseStatCount = WEIGHT_STATS.length;
    const cost = weightsEngineIterations(WEIGHTS_BROWSER_DEFAULT_ITERATIONS, worstCaseStatCount);
    expect(cost / measuredIterationsPerSecond).toBeLessThan(30);
  });
});
