// web/src/lib/sim/bulk-run.test.ts
import { describe, expect, it, vi } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { createFakeWorker } from '../../test-support/fake-worker';
import bulkResultJson from '../../fixtures/sim/bulk-result.json';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import {
  BulkCapError,
  BulkRunError,
  BulkValidationError,
  chunk,
  countCombinations,
  runBulk,
  runWeightsRun,
  stageProgressLine,
  type BulkProgress,
} from './bulk-run';
import type { BulkRequest, BulkResult, WeightsRequest, WeightsResult } from './bulk-types';
import type { SimRequest, SimResult } from './types';
import { createPool, type SimPool } from './worker';

const baseRequest = (bulkResultJson as unknown as BulkResult).request as BulkRequest;
const weightsRequest = (weightsResultJson as unknown as WeightsResult).request as WeightsRequest;

function gearRequest(overrides: Partial<BulkRequest['bulk']> = {}): BulkRequest {
  return {
    ...baseRequest,
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
      ...overrides,
    },
  };
}

function pool(options: { tickMs?: number; ticks?: number; size?: number } = {}): SimPool {
  const engine = createFakeEngine({ tickMs: options.tickMs ?? 0, ticks: options.ticks ?? 1 });
  return createPool({
    hardwareConcurrency: options.size ?? 4,
    spawn: () => createFakeWorker(engine),
  });
}

describe('chunk', () => {
  it('splits into runs of at most size, keeping order', () => {
    expect(chunk([1, 2, 3, 4, 5], 2)).toEqual([[1, 2], [3, 4], [5]]);
    expect(chunk([], 3)).toEqual([]);
    expect(chunk([1], 0)).toEqual([[1]]);
  });
});

describe('stageProgressLine', () => {
  it('reads the way the design writes it', () => {
    const progress: BulkProgress = { stage: 2, stages: 3, combosDone: 31, combosTotal: 96 };
    expect(stageProgressLine(progress)).toBe('stage 2 of 3 · 31 of 96 combinations');
  });
});

describe('countCombinations', () => {
  it('answers the live count', async () => {
    const p = pool();
    await expect(countCombinations(p, gearRequest())).resolves.toBe(3);
    p.terminate();
  });

  it('throws the cap refusal with both numbers rather than trimming', async () => {
    const p = pool();
    const error = await countCombinations(p, gearRequest({ cap: 2 })).catch((e: unknown) => e);
    expect(error).toBeInstanceOf(BulkCapError);
    expect((error as BulkCapError).cap).toBe(2);
    expect((error as BulkCapError).combinations).toBe(3);
    p.terminate();
  });

  it('validates before counting: a malformed request surfaces its own errors, not a count (engine-lane rule 2)', async () => {
    const p = pool();
    const countSpy = vi.spyOn(p, 'count');
    const invalid: BulkRequest = { ...gearRequest(), iterations: 0 };
    const error = await countCombinations(p, invalid).catch((e: unknown) => e);
    expect(error).toBeInstanceOf(BulkValidationError);
    expect((error as BulkValidationError).errors).toEqual([
      { field: 'iterations', message: 'iterations must be a positive number' },
    ]);
    expect((error as BulkValidationError).detail).toBe('iterations: iterations must be a positive number');
    // The one thing this rule exists to prevent: a bad request must never reach `simCount`
    // and come back as a cap or combination answer that means nothing.
    expect(countSpy).not.toHaveBeenCalled();
    p.terminate();
  });
});

describe('runBulk', () => {
  it('runs every stage and returns a ranked result', async () => {
    const p = pool();
    const seen: BulkProgress[] = [];
    const handle = runBulk(p, gearRequest(), (progress) => seen.push({ ...progress }));
    const result = (await handle.result) as BulkResult;

    expect(result.combos.length).toBeGreaterThan(0);
    expect(result.equipped.mean).toBeGreaterThan(0);
    expect(result.stages).toHaveLength(3);
    expect(result.lane).toBe('browser');
    expect(result.aborted).toBeUndefined();
    expect(seen.map((progress) => progress.stage)).toEqual(expect.arrayContaining([1, 2, 3]));
    expect(seen.every((progress) => progress.stages === 3)).toBe(true);
    p.terminate();
  });

  it('refuses past the cap before it runs anything', async () => {
    const p = pool();
    const handle = runBulk(p, gearRequest({ cap: 1 }), () => {});
    await expect(handle.result).rejects.toBeInstanceOf(BulkCapError);
    p.terminate();
  });

  it('a stop returns the combinations that finished, marked partial', async () => {
    // Cancels the instant the progress callback reports the first combo done, rather than
    // guessing a wall-clock offset: a one-worker pool runs the equipped baseline and each
    // combo one request at a time, so `combosDone >= 1` fires exactly between the first
    // combo finishing and the second one's `pool.run` call. `handle.cancel()` then runs as
    // a microtask queued ahead of the second combo's own first tick (a real timer, however
    // small `tickMs` is), so it always lands before that combo's engine call checks whether
    // it was aborted -- deterministic regardless of runner speed or scheduling jitter, with
    // no dependency on two independently-scheduled `setTimeout` chains landing in order.
    const p = pool({ tickMs: 0, ticks: 1, size: 1 });
    let resolveOneComboDone: () => void = () => {};
    const oneComboDone = new Promise<void>((resolve) => {
      resolveOneComboDone = resolve;
    });
    const handle = runBulk(p, gearRequest(), (progress) => {
      if (progress.combosDone >= 1) resolveOneComboDone();
    });
    await oneComboDone;
    handle.cancel();
    const result = (await handle.result) as BulkResult;
    expect(result.aborted).toBe(true);
    expect(result.combos.length).toBeGreaterThanOrEqual(1);
    p.terminate();
  });

  it('a stop before the equipped run finishes has nothing to return', async () => {
    const p = pool({ tickMs: 200, ticks: 2, size: 1 });
    const handle = runBulk(p, gearRequest(), () => {});
    handle.cancel();
    const error = await handle.result.catch((e: unknown) => e);
    expect(error).toBeInstanceOf(BulkRunError);
    expect((error as BulkRunError).cancelled).toBe(true);
    p.terminate();
  });

  it('keeps the engine’s own words on a failure', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1, failWith: 'request: unknown buff: "x"' });
    const p = createPool({ hardwareConcurrency: 2, spawn: () => createFakeWorker(engine) });
    const handle = runBulk(p, gearRequest(), () => {});
    const error = await handle.result.catch((e: unknown) => e);
    expect(error).toBeInstanceOf(BulkRunError);
    expect((error as BulkRunError).cancelled).toBe(false);
    expect((error as BulkRunError).detail).toContain('unknown buff');
    p.terminate();
  });

  it('stamps the browser lane and the wall clock', async () => {
    const p = pool();
    const clock = vi.fn().mockReturnValueOnce(1000).mockReturnValue(4500);
    const handle = runBulk(p, gearRequest(), () => {}, clock);
    const result = await handle.result;
    expect(result.duration_ms).toBe(3500);
    p.terminate();
  });

  it('stamps the wall clock on the abort-with-partial exit path too (engine-lane rule 4)', async () => {
    // Browser results otherwise report `duration_ms: 0` until something stamps them; the
    // success path is covered above, this pins the OTHER exit path `finish()` shares with
    // it -- a stop that still has a partial result to hand back.
    const p = pool({ tickMs: 0, ticks: 1, size: 1 });
    const clock = vi.fn().mockReturnValueOnce(2000).mockReturnValue(2750);
    let resolveOneComboDone: () => void = () => {};
    const oneComboDone = new Promise<void>((resolve) => {
      resolveOneComboDone = resolve;
    });
    const handle = runBulk(
      p,
      gearRequest(),
      (progress) => {
        if (progress.combosDone >= 1) resolveOneComboDone();
      },
      clock,
    );
    await oneComboDone;
    handle.cancel();
    const result = (await handle.result) as BulkResult;
    expect(result.aborted).toBe(true);
    expect(result.duration_ms).toBe(750);
    p.terminate();
  });
});

describe('order preservation (engine-lane rule 1)', () => {
  it('keeps results in request order even when shards resolve out of order', async () => {
    // A stub SimPool, not the fake engine: every combo in a real stage shares the same
    // `random_seed`/iterations and so produces byte-identical DPS (Task 3's report), which
    // makes it impossible to tell combos apart by their results at all. Here each request
    // carries its own `random_seed` as a fingerprint, and `run()` deliberately resolves the
    // LAST shard of each chunk first -- proving that a reordering bug in `runBulk`'s own
    // chunking/collection (not `pool.run`'s `Promise.all`, which is a language guarantee)
    // would be caught: `rank` is handed the wrong-order results if one exists anywhere on
    // this path.
    const fixtureSummary = (bulkResultJson as unknown as BulkResult).summary;
    const { bulk, ...bare } = baseRequest;
    void bulk;
    const withSeed = (seed: number): SimRequest => ({ ...bare, random_seed: seed });
    const resultFor = (request: SimRequest): SimResult => ({
      engine_version: request.engine_version,
      request,
      lane: 'browser',
      dps: { mean: request.random_seed, stddev: 0, error: 0, min: 0, max: 0 },
      iterations_run: request.iterations,
      duration_ms: 0,
      summary: fixtureSummary,
    });

    const rankCalls: string[] = [];
    const notUsed = (): never => {
      throw new Error('not used by this test');
    };
    const stubPool: SimPool = {
      size: 2,
      split: notUsed,
      combine: notUsed,
      needsMore: notUsed,
      validate: notUsed,
      count: notUsed,
      weights: notUsed,
      abort: () => {},
      terminate: () => {},
      async plan() {
        return {
          ok: true,
          stage: {
            stage: 1,
            iterations: 100,
            requests: [0, 1, 2, 3].map(withSeed),
            combos: [1, 2, 3].map((seed) => ({ request: withSeed(seed), substitutions: [] })),
            ran: [],
          },
        };
      },
      run(shards) {
        const promises = shards.map(
          (json, index) =>
            new Promise<string>((resolve) => {
              const request = JSON.parse(json) as SimRequest;
              const reverseDelay = (shards.length - 1 - index) * 5;
              setTimeout(() => resolve(JSON.stringify(resultFor(request))), reverseDelay);
            }),
        );
        return Promise.all(promises);
      },
      async rank(_request, _stage, resultsJSON) {
        rankCalls.push(resultsJSON);
        const result: BulkResult = {
          ...resultFor(bare),
          combos: [],
          equipped: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
          stages: [],
        };
        return { result };
      },
    };

    const handle = runBulk(stubPool, gearRequest(), () => {});
    await handle.result;

    expect(rankCalls).toHaveLength(1);
    const seenSeeds = (JSON.parse(rankCalls[0]) as SimResult[]).map((r) => r.dps.mean);
    expect(seenSeeds).toEqual([0, 1, 2, 3]);
  });
});

describe('runWeightsRun', () => {
  it('returns the weights and ticks progress', async () => {
    const p = pool({ tickMs: 0, ticks: 3 });
    const ticks: number[] = [];
    const handle = runWeightsRun(p, weightsRequest, (done) => ticks.push(done));
    const result = (await handle.result) as WeightsResult;
    expect(result.weights.length).toBe(6);
    expect(ticks.length).toBeGreaterThan(0);
    p.terminate();
  });

  it('stamps the wall clock (engine-lane rule 4)', async () => {
    const p = pool({ tickMs: 0, ticks: 3 });
    const clock = vi.fn().mockReturnValueOnce(500).mockReturnValue(900);
    const handle = runWeightsRun(p, weightsRequest, () => {}, clock);
    const result = await handle.result;
    expect(result.duration_ms).toBe(400);
    p.terminate();
  });
});
