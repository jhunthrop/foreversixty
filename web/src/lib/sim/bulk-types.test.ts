// web/src/lib/sim/bulk-types.test.ts
import { describe, expect, it } from 'vitest';
import {
  BROWSER_CAP,
  LOW_CORE_CAP,
  SERVER_CAP,
  STAGES_BY_PRECISION,
  browserCap,
  finalIterations,
  isCapExceeded,
  requestKind,
  BULK_PRECISIONS,
  type BulkMode,
  type BulkRequest,
  type BulkResult,
  type Candidate,
  type Combination,
  type Combo,
  type GearSet,
  type Precision,
  type RankAnswer,
  type StageRequests,
  type Stage,
  type StatWeight,
  type Substitution,
  type TalentLoadout,
  type WeightsRequest,
  type WeightsResult,
  type WeightsSpec,
  type BulkSpec,
  type SimKind,
} from './bulk-types';
import type { SimRequest, SimResult } from './types';

const base: SimRequest = {
  engine_version: 'testver',
  spec: 'warrior-fury',
  source: { kind: 'addon', ref: '', captured_at: '2026-09-19T00:00:00Z' },
  character: {
    name: 'Fury',
    race: 'orc',
    class: 'warrior',
    level: 60,
    talents: '0-5530515-',
    gear: [{ slot: 'head', item_id: 12640 }],
    buffs: [],
    consumes: [],
  },
  encounter: { duration_sec: 180, variation: 0.2, targets: 1, execute_ratio: 0.25, profile: '' },
  iterations: 3000,
  random_seed: 0,
};

describe('browserCap', () => {
  it('is 400 on a desktop and 200 on four cores or fewer', () => {
    expect(browserCap(8)).toBe(BROWSER_CAP);
    expect(browserCap(6)).toBe(BROWSER_CAP);
    expect(browserCap(4)).toBe(LOW_CORE_CAP);
    expect(browserCap(2)).toBe(LOW_CORE_CAP);
  });

  it('assumes the low-core cap when the browser will not say', () => {
    expect(browserCap(undefined)).toBe(LOW_CORE_CAP);
    expect(browserCap(Number.NaN)).toBe(LOW_CORE_CAP);
    expect(browserCap(0)).toBe(LOW_CORE_CAP);
  });
});

describe('requestKind', () => {
  it('derives run, gear, talents, drops and weights the way SimRequest.Kind() does', () => {
    expect(requestKind(base)).toBe('run');
    for (const mode of ['gear', 'talents', 'drops'] as const) {
      const request: BulkRequest = {
        ...base,
        bulk: { mode, candidates: [], precision: 'normal', cap: BROWSER_CAP },
      };
      expect(requestKind(request)).toBe(mode);
    }
    const weights: WeightsRequest = { ...base, weights: { stats: ['crit'], reference: 'crit' } };
    expect(requestKind(weights)).toBe('weights');
  });

  it('prefers bulk over weights, because a bulk request is never a weights run', () => {
    const both = {
      ...base,
      bulk: { mode: 'gear' as const, candidates: [], precision: 'fast' as const, cap: 400 },
      weights: { stats: ['crit'], reference: 'crit' },
    };
    expect(requestKind(both)).toBe('gear');
  });
});

describe('isCapExceeded', () => {
  it('recognises the wasm refusal and nothing else', () => {
    expect(isCapExceeded({ error: 'cap_exceeded', cap: 400, combinations: 812 })).toBe(true);
    expect(isCapExceeded({ error: 'request: unknown buff' })).toBe(false);
    expect(isCapExceeded({ stage: 1, iterations: 100, requests: [], combos: [] })).toBe(false);
    expect(isCapExceeded(null)).toBe(false);
  });
});

describe('BULK_PRECISIONS and finalIterations', () => {
  it('is the contract’s three, in ladder order', () => {
    expect([...BULK_PRECISIONS]).toEqual(['fast', 'normal', 'high']);
  });

  it('gives each precision its final-stage iteration count (contract 10.1 A3)', () => {
    expect(finalIterations('fast')).toBe(3000);
    expect(finalIterations('normal')).toBe(3000);
    expect(finalIterations('high')).toBe(10_000);
  });
});

describe('the lane caps', () => {
  it('are 400 in the browser and 5,000 on the server (contract 10.1 A2)', () => {
    expect(BROWSER_CAP).toBe(400);
    expect(SERVER_CAP).toBe(5000);
  });
});

describe('STAGES_BY_PRECISION', () => {
  it('is the progress line’s stage count per precision: fast 3, normal 2, high 2', () => {
    expect(STAGES_BY_PRECISION).toEqual({ fast: 3, normal: 2, high: 2 });
  });
});

// --- re-export surface ---
//
// Everything below is a compile-time check, in the same spirit as kind.test.ts's fixture
// cast: `astro check` fails loudly if part A renames or reshapes any of these, rather than
// a later task discovering it as a silent runtime mismatch. Nothing here needs a runtime
// assertion because an interface has no runtime representation; the values still exercised
// (mode, precision, kind, slot, origin) are read back to confirm the field names line up
// with what `requestKind`/`isCapExceeded` above actually consume.
describe('the re-exported part-A shapes', () => {
  it('accepts a Candidate, TalentLoadout and GearSet shaped exactly as sim/api/envelope.go emits them', () => {
    const candidate: Candidate = {
      slot: '',
      item_id: 19019,
      origin: 'drop:ragnaros',
      source_name: 'Ragnaros',
    };
    const loadout: TalentLoadout = { name: 'Arms-Fury', talents: '0-5530515-' };
    const set: GearSet = { name: 'Dreadnaught', gear: [{ slot: 'head', item_id: 16955 }] };
    expect(candidate.origin).toBe('drop:ragnaros');
    expect(loadout.talents).toBe('0-5530515-');
    expect(set.gear[0]?.slot).toBe('head');
  });

  it('accepts a BulkSpec and WeightsSpec built the way BulkRequest/WeightsRequest carry them', () => {
    const bulk: BulkSpec = { mode: 'gear', candidates: [], precision: 'high', cap: BROWSER_CAP };
    const weights: WeightsSpec = { stats: ['crit', 'agility'], reference: 'crit' };
    expect(bulk.mode).toBe('gear');
    expect(weights.reference).toBe('crit');
  });

  it('accepts a Substitution, Combo, Stage and StatWeight as sim/bulk emits them', () => {
    const substitution: Substitution = { kind: 'item', slot: 'head', item_id: 16955, name: 'Helm of Wrath' };
    const combo: Combo = {
      substitutions: [substitution],
      dps: { mean: 1200, stddev: 10, error: 1, min: 1150, max: 1250 },
      delta: { mean: 40, stddev: 2, error: 0.5, min: 30, max: 50 },
      group: 0,
    };
    const stage: Stage = { iterations: 3000, combos: 12 };
    const weight: StatWeight = { stat: 'crit', weight: 1, error: 0.01 };
    expect(combo.group).toBe(0);
    expect(stage.combos).toBe(12);
    expect(weight.stat).toBe('crit');
  });

  it('accepts every SimKind value, the same vocabulary requestKind answers with', () => {
    const kinds: SimKind[] = ['run', 'gear', 'talents', 'drops', 'weights'];
    expect(kinds).toHaveLength(5);
  });
});

describe('the new bulk-types declarations', () => {
  it('types BulkResult and WeightsResult as SimResult plus their bulk-only fields', () => {
    const result: BulkResult = {
      engine_version: 'testver',
      request: base,
      lane: 'browser',
      dps: { mean: 1200, stddev: 10, error: 1, min: 1150, max: 1250 },
      iterations_run: 3000,
      duration_ms: 500,
      summary: {} as BulkResult['summary'],
      combos: [],
      equipped: { mean: 1160, stddev: 10, error: 1, min: 1100, max: 1200 },
      stages: [{ iterations: 3000, combos: 4 }],
    };
    const weightsResult: WeightsResult = {
      engine_version: 'testver',
      request: base,
      lane: 'browser',
      dps: { mean: 1200, stddev: 10, error: 1, min: 1150, max: 1250 },
      iterations_run: 3000,
      duration_ms: 500,
      summary: {} as WeightsResult['summary'],
      weights: [{ stat: 'crit', weight: 1, error: 0.01 }],
    };
    expect(result.stages).toHaveLength(1);
    expect(weightsResult.weights).toHaveLength(1);
  });

  it('types Combination, StageRequests and RankAnswer as simPlan/simRank exchange them', () => {
    const combination: Combination = { request: base, substitutions: [] };
    const stageRequests: StageRequests = {
      stage: 1,
      iterations: 1000,
      requests: [base],
      combos: [combination],
    };
    const nextAnswer: RankAnswer = { next: stageRequests };
    const finishedResult: SimResult = {
      engine_version: 'testver',
      request: base,
      lane: 'browser',
      dps: { mean: 1200, stddev: 10, error: 1, min: 1150, max: 1250 },
      iterations_run: 3000,
      duration_ms: 500,
      summary: {} as SimResult['summary'],
    };
    const doneAnswer: RankAnswer = { result: finishedResult };
    expect(nextAnswer.next?.combos).toHaveLength(1);
    expect(doneAnswer.result?.lane).toBe('browser');
  });

  it('accepts every BulkMode value and derives Precision from BULK_PRECISIONS', () => {
    const modes: BulkMode[] = ['gear', 'talents', 'drops'];
    const precisions: Precision[] = [...BULK_PRECISIONS];
    expect(modes).toHaveLength(3);
    expect(precisions).toEqual(['fast', 'normal', 'high']);
  });
});
