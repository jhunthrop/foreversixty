// web/src/lib/sim/weights.test.ts
import { describe, expect, it } from 'vitest';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import specsJson from '../../fixtures/sim/specs.json';
import {
  CASTER_REFERENCE,
  DEFAULT_REFERENCE,
  RANGED_REFERENCE,
  defaultStatsFor,
  fallbackReferenceFor,
  formatWeightError,
  hasWeightStats,
  isDpsSpec,
  isSignificant,
  pickableStatsFor,
  pickableStatsForSpec,
  toolScaleView,
  referenceFor,
  statLabel,
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
    // Contract 10.8 (corrected 2026-09-28): attack_power for melee, ranged_attack_power
    // for hunters, spell_power for casters.
    expect(referenceFor('warrior-fury', [])).toBe(DEFAULT_REFERENCE);
    expect(referenceFor('mage-fire', [])).toBe(CASTER_REFERENCE);
    expect(referenceFor('druid-balance', specs)).toBe(CASTER_REFERENCE);
    expect(fallbackReferenceFor('hunter-marksmanship')).toBe(RANGED_REFERENCE);
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

// Round 11 (owner ruling 2026-09-30, memory feedback-stat-weights-simc): the tool shows weights
// the way the published ones are -- SimulationCraft scale factors, the top significant
// per-point stat = 1.00, sorted, with DPS per point and the error; haste per 1% in a caption,
// never a row; no Pawn string. It used to show reference = 1.00 with a Pawn export.
describe('toolScaleView', () => {
  const weights: StatWeight[] = [
    { stat: 'attack_power', weight: 1, error: 0.01, dps_per_point: 0.06 },
    { stat: 'strength', weight: 2, error: 0.04, dps_per_point: 0.12 },
    { stat: 'crit', weight: 0.5, error: 0.02, dps_per_point: 0.03 },
    { stat: 'agility', weight: 1.5, error: 3, insignificant: true, dps_per_point: 0.09 },
    { stat: 'melee_haste', weight: 6, error: 0.3, dps_per_point: 0.36 },
  ];
  const view = toolScaleView(weights, 'warrior-fury', 'attack_power');

  // The shared weight reading (lib/bis/weight-reading.ts): measurable rows first, descending,
  // then the insignificant ones.
  it('normalizes to the top significant per-point stat and sorts descending', () => {
    expect(view.rows.map((row) => row.stat)).toEqual(['strength', 'attack_power', 'crit', 'agility']);
    expect(view.rows[0].scaleFactor).toBe(1);
    expect(view.rows.find((row) => row.stat === 'attack_power')?.scaleFactor).toBeCloseTo(0.5, 6);
    expect(view.rows.find((row) => row.stat === 'crit')?.scaleError).toBeCloseTo(0.01, 6);
    expect(view.noteLine).toContain('normalized to Strength = 1.00');
  });

  it("carries each stat's DPS per point and greys an insignificant row", () => {
    expect(view.rows.find((row) => row.stat === 'strength')?.dpsPerPoint).toBeCloseTo(0.12, 6);
    expect(view.rows.find((row) => row.stat === 'agility')?.significant).toBe(false);
  });

  it('puts haste per 1% in the caption, never in the table', () => {
    expect(view.rows.some((row) => row.stat === 'melee_haste')).toBe(false);
    expect(view.hasteCaption).toBe('Haste: 3.00 per 1%');
  });

  it('shows no DPS per point from an engine that sent none, never a made-up one', () => {
    const old = toolScaleView(
      [{ stat: 'attack_power', weight: 1, error: 0 }],
      'warrior-fury',
      'attack_power',
    );
    expect(old.rows[0].dpsPerPoint).toBeUndefined();
  });

  it('the fixture run reads the same way', () => {
    const reference = result.request.weights?.reference;
    expect(reference).toBeDefined();
    const fixture = toolScaleView(result.weights, 'warrior-fury', reference ?? '');
    expect(Math.max(...fixture.rows.map((row) => row.scaleFactor))).toBe(1);
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

describe('pickableStatsForSpec', () => {
  it('offers a Fury warrior only its own stats, never a spell school or MP5 (seo-fix 2026-11-10)', () => {
    const ids = pickableStatsForSpec('warrior-fury').map((stat) => stat.id);
    expect(ids).toEqual(specRow('warrior-fury')?.weight_stats);
    expect(ids).not.toContain('mp5');
    expect(ids.some((id) => id.endsWith('_power') && id !== 'attack_power')).toBe(false);
  });

  it('falls back to the full vocabulary only for a spec nobody has heard of', () => {
    expect(pickableStatsForSpec('nobody-knows')).toEqual(pickableStatsFor(undefined));
  });
});
