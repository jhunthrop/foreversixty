// @vitest-environment jsdom
import { describe, expect, it, vi } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { simCopy } from './copy';
import { SimRunError, buildSimRequest, runSim, type RunInput, type RunUpdate } from './run';
import { DEFAULT_ENCOUNTER, ITERATIONS } from './types';
import { ENGINE_VERSION } from './version';
import { createBrokenWorker, createFakeWorker } from '../../test-support/fake-worker';
import { createPool, type PoolWorker, type SimPool } from './worker';

const input: RunInput = {
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
  iterations: ITERATIONS.normal,
  randomSeed: 5,
};

/** The shared in-process worker (Task 3), wrapped so each test names its own tick rate. */
function fakeWorker(tickMs = 0): PoolWorker {
  return createFakeWorker(createFakeEngine({ tickMs, ticks: 3 }));
}

describe('buildSimRequest', () => {
  it('is the contract envelope, with the character inline and no protobuf anywhere', () => {
    const request = buildSimRequest(input);
    expect(request).toEqual({
      engine_version: ENGINE_VERSION,
      spec: 'warrior-fury',
      source: { kind: 'addon', ref: '', captured_at: '2026-09-14T10:00:00Z' },
      character: input.character,
      encounter: DEFAULT_ENCOUNTER,
      iterations: 3000,
      random_seed: 5,
    });
    expect('raw' in request).toBe(false);
  });

  it('round-trips through JSON unchanged, which is all the wasm is handed', () => {
    expect(JSON.parse(JSON.stringify(buildSimRequest(input)))).toEqual(buildSimRequest(input));
  });
});

describe('runSim', () => {
  it('opens at zero, refines as shards land, and finishes on the engine’s own numbers', async () => {
    const pool = createPool({ hardwareConcurrency: 4, spawn: () => fakeWorker() });
    const updates: RunUpdate[] = [];
    const result = await runSim(pool, input, (update) => updates.push(update)).result;

    expect(updates[0]).toEqual({
      estimate: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      iterationsDone: 0,
      iterationsTotal: 3000,
      relativeError: 0,
    });
    const progressed = updates.map((u) => u.iterationsDone);
    expect(progressed).toEqual([...progressed].sort((a, b) => a - b));
    expect(updates.at(-1)?.iterationsDone).toBe(3000);
    expect(result.iterations_run).toBe(3000);
    expect(result.lane).toBe('browser');
    expect(result.dps.error).toBeCloseTo(result.dps.stddev / Math.sqrt(3000), 9);
    pool.terminate();
  });

  it('gets a finished summary from the engine, not one it built itself', async () => {
    const pool = createPool({ hardwareConcurrency: 4, spawn: () => fakeWorker() });
    const result = await runSim(pool, input, () => {}).result;
    expect(result.summary.engine_version).toBe(`sim:${ENGINE_VERSION}`);
    // The engine names nothing: these are its own action keys, straight out of the golden
    // the fixture was derived from (Task 2 Step 5). Resolving them is Task 23's job, and
    // no test in group A may assume a display name the adapter never writes.
    expect(result.summary.damage_done[0].abilities[0].name).toBe('spell:25286');
    expect(result.summary.auras[0].name).toBe('spell:9910');
    // And the row identity is the client spell id for those two, because both are plain
    // untagged spells; anything tagged sits at or above 2,000,000.
    expect(result.summary.damage_done[0].abilities[0].spell_id).toBe(25286);
    pool.terminate();
  });

  it('pools the shard progress into the numbers the engine itself combines to', async () => {
    const pool = createPool({ hardwareConcurrency: 4, spawn: () => fakeWorker() });
    const atFinish: RunUpdate[] = [];
    const result = await runSim(pool, input, (update) => {
      if (update.iterationsDone === 3000) atFinish.push(update);
    }).result;
    const last = atFinish.find((update) => update.iterationsDone === 3000);
    expect(last?.estimate.mean).toBeCloseTo(result.dps.mean, 6);
    expect(last?.estimate.stddev).toBeCloseTo(result.dps.stddev, 6);
    pool.terminate();
  });

  it('is within the contract’s error for 3,000 iterations', async () => {
    const pool = createPool({ hardwareConcurrency: 4, spawn: () => fakeWorker() });
    const result = await runSim(pool, input, () => {}).result;
    expect((1.96 * result.dps.error) / result.dps.mean).toBeLessThan(0.0075);
    pool.terminate();
  });

  it('reports a cancel as cancelled, not as a failure', async () => {
    vi.useFakeTimers();
    const pool = createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker(20) });
    const handle = runSim(pool, input, () => {});
    const settled = handle.result.catch((error: unknown) => error);
    await vi.advanceTimersByTimeAsync(25);
    handle.cancel();
    await vi.advanceTimersByTimeAsync(200);
    const error = await settled;
    expect(error).toBeInstanceOf(SimRunError);
    expect((error as SimRunError).cancelled).toBe(true);
    expect((error as SimRunError).message).toBe(simCopy.stopped);
    pool.terminate();
    vi.useRealTimers();
  });

  it('reports an engine failure as a failure, with copy a player can read', async () => {
    const pool = createPool({ hardwareConcurrency: 1, spawn: () => createBrokenWorker() });
    const error = await runSim(pool, input, () => {}).result.catch((e: unknown) => e);
    expect((error as SimRunError).cancelled).toBe(false);
    expect((error as SimRunError).message).toBe(simCopy.failed);
    pool.terminate();
  });

  it('keeps the engine’s own words beside the generic sentence', async () => {
    // sim/request answers an unknown buff with `request: unknown buff: "battle-shout"`,
    // which names the thing to fix; simCopy.failed alone does not. The UI shows both
    // (Task 13). This drives the real engine surface through the real pool -- the fake's
    // `failWith` option (Task 3) makes simRun reject with exactly that message -- rather
    // than a hand-built worker, so the message survives every hop the real one does:
    // simRun, sim.worker.ts's catch, the pool's `failed` message, and runSim's detailOf.
    const message = 'request: unknown buff: "battle-shout"';
    const pool = createPool({
      hardwareConcurrency: 1,
      spawn: () => createFakeWorker(createFakeEngine({ tickMs: 0, failWith: message })),
    });
    const error = await runSim(pool, input, () => {}).result.catch((e: unknown) => e);
    expect((error as SimRunError).message).toBe(simCopy.failed);
    expect((error as SimRunError).detail).toBe(message);
    pool.terminate();
  });

  it('never asks for more shards than there are iterations', async () => {
    const pool = createPool({ hardwareConcurrency: 8, spawn: () => fakeWorker() });
    const result = await runSim(pool, { ...input, iterations: 3 }, () => {}).result;
    expect(result.iterations_run).toBe(3);
    pool.terminate();
  });
});

describe('a target-error run', () => {
  /**
   * A fake engine whose relative error shrinks with every step, so the loop terminates on
   * the engine's own answer rather than on a count this test chose. `tickMs: 0` keeps it
   * instant; `ticks: 1` keeps the progress traffic down.
   */
  function stepPool(): SimPool {
    return createPool({
      hardwareConcurrency: 2,
      spawn: () => createFakeWorker(createFakeEngine({ tickMs: 0, ticks: 1 })),
    });
  }

  it('runs in thousand-iteration steps and stops when the engine says the band is inside the target', async () => {
    const pool = stepPool();
    const updates: RunUpdate[] = [];
    const handle = runSim(
      pool,
      { ...input, iterations: 30_000, targetError: 0.005, stepIterations: 1000 },
      (update) => updates.push(update),
    );
    const result = await handle.result;

    // Every step is a multiple of the step size, nothing overshoots the ceiling, and the
    // pooled figure is the whole run's, not the last step's.
    expect(result.iterations_run % 1000).toBe(0);
    expect(result.iterations_run).toBeGreaterThanOrEqual(1000);
    expect(result.iterations_run).toBeLessThanOrEqual(30_000);
    expect(result.dps.error / result.dps.mean).toBeLessThanOrEqual(0.005);
    expect(result.request.target_error).toBe(0.005);
    // The ceiling is what the request carries, so a saved sim says what bounded it.
    expect(result.request.iterations).toBe(30_000);
    // The progress line has an error to show from the first step onwards.
    expect(updates.at(-1)?.relativeError).toBeLessThanOrEqual(0.005);
    expect(updates.at(-1)?.iterationsDone).toBe(result.iterations_run);
    pool.terminate();
  });

  it('stops at the ceiling when the band never closes', async () => {
    const pool = stepPool();
    const handle = runSim(
      pool,
      { ...input, iterations: 3000, targetError: 0.0000001, stepIterations: 1000 },
      () => {},
    );
    const result = await handle.result;
    expect(result.iterations_run).toBe(3000);
    pool.terminate();
  });

  it('is today’s single pass when no target error is asked for', async () => {
    const pool = stepPool();
    const handle = runSim(pool, { ...input, iterations: 500 }, () => {});
    const result = await handle.result;
    expect(result.iterations_run).toBe(500);
    expect(result.request.target_error).toBeUndefined();
    pool.terminate();
  });

  it('a stop between steps ends the run rather than starting another one', async () => {
    const pool = stepPool();
    const handle = runSim(
      pool,
      { ...input, iterations: 30_000, targetError: 0.0000001, stepIterations: 1000 },
      () => {},
    );
    handle.cancel();
    await expect(handle.result).rejects.toThrow(SimRunError);
    pool.terminate();
  });
});
