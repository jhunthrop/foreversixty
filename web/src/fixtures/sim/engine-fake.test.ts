import { describe, expect, it, vi } from 'vitest';
import type { SimProgressUpdate, SimRequest, SimResult } from '../../lib/sim/types';
import { DEFAULT_ENCOUNTER } from '../../lib/sim/types';
import { ENGINE_VERSION } from '../../lib/sim/version';
import {
  isCapExceeded,
  type BulkRequest,
  type BulkResult,
  type CapExceeded,
  type Precision,
  type RankAnswer,
  type StageRequests,
  type WeightsRequest,
  type WeightsResult,
} from '../../lib/sim/bulk-types';
import { createFakeEngine } from './engine-fake';
import fixtureResultJson from './result.json';
import fixtureBulkResultJson from './bulk-result.json';
import fixtureWeightsResultJson from './weights-result.json';

const fixtureResult = fixtureResultJson as unknown as SimResult;
const fixtureRequest = fixtureResult.request;
const fixtureBulkResult = fixtureBulkResultJson as unknown as BulkResult & { request: BulkRequest };
const fixtureWeightsResult = fixtureWeightsResultJson as unknown as WeightsResult & {
  request: WeightsRequest;
};

const request: SimRequest = {
  engine_version: ENGINE_VERSION,
  spec: 'warrior-fury',
  source: { kind: 'addon', ref: '', captured_at: '2026-09-14T10:00:00Z' },
  character: {
    name: 'Thrallgar',
    race: 'orc',
    class: 'warrior',
    level: 60,
    talents: '-5530515-',
    gear: [{ slot: 'head', item_id: 12640 }],
    buffs: [],
    consumes: [],
  },
  encounter: DEFAULT_ENCOUNTER,
  iterations: 3000,
  random_seed: 7,
};

const json = (value: SimRequest): string => JSON.stringify(value);
const resultOf = (resultJSON: string): SimResult => JSON.parse(resultJSON) as SimResult;
// simSplit returns ONE JSON string encoding a SimRequest[] (main.go), never a JS array.
const splitOf = (splitJSON: string): SimRequest[] => JSON.parse(splitJSON) as SimRequest[];

describe('simSplit', () => {
  it('returns one JSON string, not a JS array', () => {
    const engine = createFakeEngine({ tickMs: 0 });
    const splitJSON = engine.simSplit(json(request), 8);
    expect(typeof splitJSON).toBe('string');
    expect(Array.isArray(splitJSON)).toBe(false);
  });

  it('splits iterations so the shards sum to the original', () => {
    const engine = createFakeEngine({ tickMs: 0 });
    const shards = splitOf(engine.simSplit(json(request), 8));
    expect(shards).toHaveLength(8);
    expect(shards.reduce((sum, shard) => sum + shard.iterations, 0)).toBe(3000);
  });

  it('never emits a shard with no iterations', () => {
    const engine = createFakeEngine({ tickMs: 0 });
    expect(splitOf(engine.simSplit(json({ ...request, iterations: 3 }), 8))).toHaveLength(3);
  });

  it('returns request objects, not JSON strings, so the pool re-stringifies each shard', () => {
    const engine = createFakeEngine({ tickMs: 0 });
    const [first] = splitOf(engine.simSplit(json(request), 4));
    expect(typeof first).toBe('object');
    expect(first.character.talents).toBe('-5530515-');
  });
});

describe('simRun', () => {
  it('reports progress before it resolves, and the progress sums to the shard', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 4 });
    const seen: number[] = [];
    engine.onProgress((callbackId, progressJSON) => {
      expect(callbackId).toBe('run-1');
      seen.push((JSON.parse(progressJSON) as SimProgressUpdate).iterations_run);
    });
    const result = resultOf(await engine.simRun(json({ ...request, iterations: 400 }), 'run-1'));
    expect(seen).toEqual([100, 200, 300, 400]);
    expect(result.iterations_run).toBe(400);
  });

  it('returns a finished SimResult, summary included, because the adapter runs inside', async () => {
    const engine = createFakeEngine({ tickMs: 0 });
    const result = resultOf(await engine.simRun(json(request), 'run-2'));
    expect(result.summary.engine_version).toBe(`sim:${ENGINE_VERSION}`);
    // The engine names abilities by action key, not display name (sim/adapter.ActionName);
    // resolving them is Task 23's job and no test here may assume a display name.
    expect(result.summary.damage_done[0].abilities[0].name).toMatch(/^(spell|item|other):/);
    expect(result.lane).toBe('browser');
  });

  it('gives the same numbers for the same seed, and different ones for a different seed', async () => {
    const engine = createFakeEngine({ tickMs: 0 });
    const a = resultOf(await engine.simRun(json(request), 'a'));
    const b = resultOf(await engine.simRun(json(request), 'b'));
    const c = resultOf(await engine.simRun(json({ ...request, random_seed: 8 }), 'c'));
    expect(a.dps.mean).toBe(b.dps.mean);
    expect(c.dps.mean).not.toBe(a.dps.mean);
  });

  it('rejects when the run is aborted mid-flight, and simAbort answers {"aborted": true}', async () => {
    vi.useFakeTimers();
    const engine = createFakeEngine({ tickMs: 10, ticks: 4 });
    const pending = engine.simRun(json(request), 'run-abort');
    const settled = expect(pending).rejects.toThrow('aborted');
    await vi.advanceTimersByTimeAsync(15);
    expect(JSON.parse(engine.simAbort('run-abort'))).toEqual({ aborted: true });
    await vi.advanceTimersByTimeAsync(50);
    await settled;
    vi.useRealTimers();
  });

  it('answers {"aborted": false} for a run that was never started, distinguishing it from a real stop', () => {
    const engine = createFakeEngine({ tickMs: 0 });
    expect(JSON.parse(engine.simAbort('nobody'))).toEqual({ aborted: false });
  });
});

describe('simCombine', () => {
  it('takes one JSON string, not a JS array', async () => {
    const engine = createFakeEngine({ tickMs: 0 });
    const shards = splitOf(engine.simSplit(json(request), 4)).map((shard) => json(shard));
    const results = await Promise.all(shards.map((shard, i) => engine.simRun(shard, `s${i}`)));
    const combinedJSON = engine.simCombine(JSON.stringify(results.map((r) => resultOf(r))));
    expect(typeof combinedJSON).toBe('string');
  });

  it('pools the shard distributions into one finished result', async () => {
    const engine = createFakeEngine({ tickMs: 0 });
    const shards = splitOf(engine.simSplit(json(request), 4)).map((shard) => json(shard));
    const results = await Promise.all(shards.map((shard, i) => engine.simRun(shard, `s${i}`)));
    const combined = resultOf(engine.simCombine(JSON.stringify(results.map((r) => resultOf(r)))));
    expect(combined.iterations_run).toBe(3000);
    expect(combined.dps.mean).toBeGreaterThan(0);
    expect(combined.dps.error).toBeCloseTo(combined.dps.stddev / Math.sqrt(3000), 9);
    expect(combined.summary.damage_done[0].guid).toBe('sim-player');
  });
});

describe('simNeedsMore', () => {
  const request = (over: Partial<SimRequest> = {}): string =>
    JSON.stringify({ ...fixtureRequest, iterations: 30_000, target_error: 0.005, ...over });
  const result = (mean: number, error: number, iterationsRun: number): string =>
    JSON.stringify({
      ...fixtureResult,
      dps: { mean, stddev: 100, error, min: 0, max: 0 },
      iterations_run: iterationsRun,
    });

  it('asks for more while the relative error is over the target', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(1000, 20, 2000), request()))).toEqual({
      needs_more: true,
    });
  });

  it('stops once the relative error is inside the target', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(1000, 4, 2000), request()))).toEqual({
      needs_more: false,
    });
  });

  it('stops at the ceiling however wide the band still is', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(1000, 99, 30_000), request()))).toEqual({
      needs_more: false,
    });
  });

  it('never asks for more on a fixed-count run', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(1000, 99, 500), request({ target_error: 0 })))).toEqual({
      needs_more: false,
    });
  });

  it('stops rather than dividing by a mean of zero', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(0, 0, 1000), request()))).toEqual({
      needs_more: false,
    });
  });
});

describe('simValidate', () => {
  it('passes the fixture request', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simValidate(JSON.stringify(fixtureRequest)))).toEqual({
      ok: true,
      errors: [],
    });
  });

  it('names every field it refuses, so the drawer can put each one beside its line', () => {
    const engine = createFakeEngine();
    const broken = { ...fixtureRequest, spec: '', iterations: 0 };
    const answer = JSON.parse(engine.simValidate(JSON.stringify(broken))) as {
      ok: boolean;
      errors: { field: string; message: string }[];
    };
    expect(answer.ok).toBe(false);
    expect(answer.errors.map((row) => row.field).sort()).toEqual(['iterations', 'spec']);
    expect(answer.errors.every((row) => row.message.length > 0)).toBe(true);
  });

  it('answers the error envelope for a string that is not JSON at all', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simValidate('{nope'))).toHaveProperty('error');
  });
});

describe('simCount', () => {
  const bulk = (candidates: number, cap = 400): string =>
    JSON.stringify({
      ...fixtureRequest,
      bulk: {
        mode: 'gear',
        candidates: Array.from({ length: candidates }, (_, i) => ({
          slot: 'head',
          item_id: 16963 + i,
          origin: 'bag',
        })),
        precision: 'normal',
        cap,
      },
    });

  it('counts the combinations a bulk request expands to', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simCount(bulk(3)))).toEqual({ combinations: 3 });
  });

  it('answers cap_exceeded with both numbers rather than throwing them away', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simCount(bulk(5, 4)))).toEqual({
      error: 'cap_exceeded',
      cap: 4,
      combinations: 5,
    });
  });

  it('counts a request with no bulk block as no combinations at all', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simCount(JSON.stringify(fixtureRequest)))).toEqual({ combinations: 0 });
  });

  // Fix round 1 (controller ruling): simCount must agree with simPlan, since Task 15's live
  // cap-notice UI and its client-side server-cap gate both read simCount. simCount now
  // reuses simPlan's own `expand()` instead of a separately-derived sum, so the two cannot
  // drift again -- this invariant, not a hardcoded number, is what protects that.
  it('agrees with simPlan on a multi-slot gear request', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const request = {
      ...fixtureRequest,
      bulk: {
        mode: 'gear',
        candidates: [
          { slot: 'head', item_id: 16963, origin: 'bag' },
          { slot: 'shoulder', item_id: 16966, origin: 'bank' },
        ],
        talents: [],
        sets: [],
        locked: [],
        precision: 'fast',
        cap: 400,
      },
    };
    const requestJSON = JSON.stringify(request);
    const counted = JSON.parse(engine.simCount(requestJSON)) as { combinations: number };
    const planned = JSON.parse(engine.simPlan(requestJSON)) as StageRequests;
    expect(counted.combinations).toBe(planned.combos.length);
    expect(counted.combinations).toBe(3);
  });

  it('agrees with simPlan on a request carrying consumable alternatives (contract 10.1 A5)', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const request = {
      ...fixtureRequest,
      bulk: {
        mode: 'gear',
        candidates: [
          { slot: 'head', item_id: 16963, origin: 'bag' },
          { slot: 'shoulder', item_id: 16966, origin: 'bank' },
        ],
        talents: [],
        sets: [],
        locked: [],
        consumables: [['flask_of_supreme_power'], ['elixir_of_the_mongoose']],
        precision: 'fast',
        cap: 20,
      },
    };
    const requestJSON = JSON.stringify(request);
    const counted = JSON.parse(engine.simCount(requestJSON)) as { combinations: number };
    const planned = JSON.parse(engine.simPlan(requestJSON)) as StageRequests;
    expect(counted.combinations).toBe(planned.combos.length);
    expect(counted.combinations).toBe(8);
  });

  it('agrees with simPlan on the shared fixture request (sets and consumables, Task 2 fix)', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const requestJSON = JSON.stringify(fixtureBulkResult.request);
    const counted = JSON.parse(engine.simCount(requestJSON)) as { combinations: number };
    const planned = JSON.parse(engine.simPlan(requestJSON)) as StageRequests;
    expect(counted.combinations).toBe(planned.combos.length);
  });
});

describe('the fake engine’s bulk exports', () => {
  const gearRequest = (cap = 400, precision: Precision = 'fast'): string =>
    JSON.stringify({
      ...fixtureBulkResult.request,
      bulk: {
        mode: 'gear',
        candidates: [
          { slot: 'head', item_id: 16963, origin: 'bag' },
          { slot: 'shoulder', item_id: 16966, origin: 'bank' },
        ],
        talents: [],
        sets: [],
        locked: [],
        precision,
        cap,
      },
    });

  it('plans the equipped set first and every combination after it', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const stage = JSON.parse(engine.simPlan(gearRequest())) as StageRequests;
    expect(stage.stage).toBe(1);
    expect(stage.iterations).toBe(100);
    // two singles plus the pair, plus the equipped set at index 0
    expect(stage.requests).toHaveLength(4);
    expect(stage.combos).toHaveLength(3);
    expect(stage.requests[0].iterations).toBe(100);
  });

  it('starts a normal-precision plan at 1,000 iterations', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const stage = JSON.parse(engine.simPlan(gearRequest(400, 'normal'))) as StageRequests;
    expect(stage.iterations).toBe(1000);
  });

  it('refuses a plan past the cap with the count, and never trims', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const refusal: unknown = JSON.parse(engine.simPlan(gearRequest(2)));
    expect(isCapExceeded(refusal)).toBe(true);
    expect((refusal as CapExceeded).combinations).toBe(3);
    expect((refusal as CapExceeded).cap).toBe(2);
  });

  it('multiplies the gear product by the consumable alternatives (contract 10.1 A5)', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const request = JSON.parse(gearRequest(20)) as BulkRequest;
    request.bulk.consumables = [['flask_of_supreme_power'], ['elixir_of_the_mongoose']];
    const stage = JSON.parse(engine.simPlan(JSON.stringify(request))) as StageRequests;
    // (1 head + 1) x (1 shoulder + 1) combinations = 4, including gear left untouched; each
    // one gets tried against both consumable lists (A5: "each inner list replaces Consumes
    // for that combination"), so nothing is left "untouched" once a consumables list is
    // present -- 4 x 2 = 8.
    expect(stage.combos).toHaveLength(8);
  });

  // The brief's draft of this test also asserted an invalid `bulk.precision` makes
  // simValidate answer `{ok: false}` with a `bulk.precision` field error. Omitted: the
  // shipped simValidate (part A, left untouched per controller ruling) validates only
  // top-level SimRequest fields and does not inspect `bulk.*` at all, so that assertion
  // would fail against current behaviour -- see the task-3 report for the full note.
  it('answers simValidate and simNeedsMore in the shapes contract 10.2 names', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    expect(JSON.parse(engine.simValidate(gearRequest()))).toEqual({ ok: true, errors: [] });
    expect(JSON.parse(engine.simNeedsMore(JSON.stringify(fixtureBulkResult), gearRequest()))).toEqual({
      needs_more: false,
    });
  });

  it('ranks a stage into the next one and finally into a result', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const request = gearRequest();
    let stage = JSON.parse(engine.simPlan(request)) as StageRequests;
    let final: SimResult | null = null;
    for (let guard = 0; guard < 5 && final === null; guard += 1) {
      const results = await Promise.all(
        stage.requests.map((entry, index) =>
          engine.simRun(JSON.stringify(entry), `plan-${stage.stage}-${index}`),
        ),
      );
      const answer = JSON.parse(
        engine.simRank(request, JSON.stringify(stage), `[${results.join(',')}]`),
      ) as RankAnswer;
      if (answer.result !== undefined) final = answer.result;
      else stage = answer.next!;
    }
    const bulk = final as (SimResult & { combos: NonNullable<SimResult['combos']> }) | null;
    expect(bulk).not.toBeNull();
    expect(bulk!.stages!.map((entry) => entry.iterations)).toEqual([100, 1000, 3000]);
    expect(bulk!.equipped!.mean).toBeGreaterThan(0);
    const means = bulk!.combos.map((combo) => combo.dps.mean);
    expect([...means].sort((a, b) => b - a)).toEqual(means);
    expect(bulk!.combos[0].group).toBe(0);
    // Contract 10.1 A6: an item substitution comes back named, so the page never re-joins.
    expect(bulk!.combos[0].substitutions[0].name).toBeTruthy();
  });

  it('carries the ladder history on the stage object, not in the engine (contract 10.1 A10)', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const stage = JSON.parse(engine.simPlan(gearRequest())) as StageRequests;
    expect(stage.ran).toEqual([]);
  });

  it('answers a weights request with the reference stat at exactly 1', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const json = await engine.simWeights(JSON.stringify(fixtureWeightsResult.request), 'w-1');
    const result = JSON.parse(json) as WeightsResult;
    expect(result.weights.find((row) => row.stat === 'attack_power')?.weight).toBe(1);
    expect(result.weights).toHaveLength(6);
  });

  it('gives the same numbers for the same request, twice', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const one = await engine.simWeights(JSON.stringify(fixtureWeightsResult.request), 'w-1');
    const two = await engine.simWeights(JSON.stringify(fixtureWeightsResult.request), 'w-2');
    expect(JSON.parse(one).weights).toEqual(JSON.parse(two).weights);
  });
});
