// @vitest-environment jsdom
import { describe, expect, it, vi } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import fixtureResultJson from '../../fixtures/sim/result.json';
import { createFakeWorker } from '../../test-support/fake-worker';
import { simCopy } from './copy';
import type { RunHandle } from './run';
import type { SimSettings } from './settings';
import type { SimPhase } from './store.svelte';
import { createRequestMethods, type StoreRequestDeps } from './store-request';
import type { SimRequest, SimResult } from './types';
import { createPool, type PoolWorker, type SimPool } from './worker';

const fixtureResult = fixtureResultJson as unknown as SimResult;

const request: SimRequest = {
  engine_version: 'edc0c8e9a',
  spec: 'warrior-fury',
  source: { kind: 'addon', ref: '', captured_at: '2026-09-19T10:00:00Z' },
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
  encounter: { duration_sec: 180, variation: 0.2, targets: 1, execute_ratio: 0.25, profile: '' },
  iterations: 3000,
  random_seed: 0,
};

/** The shared in-process worker (Task 3), wrapped so each test names its own tick rate. */
function fakeWorker(tickMs = 0): PoolWorker {
  return createFakeWorker(createFakeEngine({ tickMs, ticks: 3 }));
}

interface Harness {
  deps: StoreRequestDeps;
  /** What the store's own `$state` fields would read, if this were the real store. */
  snapshot(): { phase: SimPhase; message: string | null; detail: string; result: SimResult | null };
  handle(): RunHandle | null;
  restoreCalls: { count: number };
}

/**
 * A `StoreRequestDeps` whose state lives in plain closures instead of `$state` -- the same
 * substitution `store.svelte.ts` itself makes for `createRequestMethods`, just without a
 * real store around it, so `runRequest`'s own logic is exercised in isolation. Every getter
 * unrelated to `runRequest` (settings, precision, character, adopt, fromPlannerCode) throws
 * if called, so a future edit that reaches for one by mistake fails this test rather than
 * passing silently.
 */
function harness(pool: SimPool, initialResult: SimResult | null): Harness {
  let phase: SimPhase = 'idle';
  let message: string | null = null;
  let detail = '';
  let result: SimResult | null = initialResult;
  let stopRequested = false;
  let handle: RunHandle | null = null;
  const restoreCalls = { count: 0 };
  const unused = (name: string) => (): never => {
    throw new Error(`runRequest should not touch ${name}`);
  };

  const deps: StoreRequestDeps = {
    getCharacter: unused('getCharacter'),
    getTalents: unused('getTalents'),
    getSettings: unused('getSettings') as () => SimSettings,
    setSettings: () => {
      throw new Error('runRequest should not touch setSettings');
    },
    getPrecisionId: unused('getPrecisionId'),
    setPrecisionId: () => {
      throw new Error('runRequest should not touch setPrecisionId');
    },
    getResult: () => result,
    setResult: (value) => {
      result = value;
    },
    setMessage: (value) => {
      message = value;
    },
    setDetail: (value) => {
      detail = value;
    },
    setPhase: (value) => {
      phase = value;
    },
    setProgress: () => {},
    getStopRequested: () => stopRequested,
    setStopRequested: (value) => {
      stopRequested = value;
    },
    setHandle: (value) => {
      handle = value;
    },
    poolOnce: () => pool,
    adopt: unused('adopt'),
    restorePreviousResult: () => {
      restoreCalls.count += 1;
    },
    fromPlannerCode: unused('fromPlannerCode'),
    treeVersion: '1.15.9.69722',
  };

  return { deps, snapshot: () => ({ phase, message, detail, result }), handle: () => handle, restoreCalls };
}

describe('createRequestMethods -- validateRequest', () => {
  it('delegates to the pool, never judging validity itself', async () => {
    const pool = createPool({ hardwareConcurrency: 1, spawn: () => fakeWorker() });
    const { deps } = harness(pool, null);
    const methods = createRequestMethods(deps);
    const answer = await methods.validateRequest(JSON.stringify(request));
    expect(answer.ok).toBe(true);
    pool.terminate();
  });
});

describe('createRequestMethods -- runRequest', () => {
  it('finishes a fixed-count run with a result and phase done', async () => {
    const pool = createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker() });
    const { deps, snapshot } = harness(pool, null);
    const methods = createRequestMethods(deps);
    await methods.runRequest(request);
    expect(snapshot().phase).toBe('done');
    expect(snapshot().result?.iterations_run).toBe(3000);
    pool.terminate();
  });

  // Task 15 fix round 1: the brief's own runRequest snippet set `phase = 'error'`
  // unconditionally in its catch, which would have shown "error" (and dropped the
  // previous result) on a Stop click mid-run-this-request -- run()'s own catch restores
  // the previous result and lands on 'done' when one exists. These two tests pin that
  // parity so a future "simplification" back to the brief's snippet fails loudly.
  it('a cancelled run restores the previous result and lands on done, not error', async () => {
    vi.useFakeTimers();
    const pool = createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker(20) });
    const { deps, snapshot, handle, restoreCalls } = harness(pool, fixtureResult);
    const methods = createRequestMethods(deps);

    const finished = methods.runRequest(request);
    await vi.advanceTimersByTimeAsync(25);
    handle()?.cancel();
    await vi.advanceTimersByTimeAsync(400);
    await finished;

    expect(restoreCalls.count).toBe(1);
    expect(snapshot().message).toBe(simCopy.stopped);
    expect(snapshot().phase).toBe('done');
    pool.terminate();
    vi.useRealTimers();
  });

  it('a cancelled run with nothing to restore still reports stopped, on phase error', async () => {
    vi.useFakeTimers();
    const pool = createPool({ hardwareConcurrency: 2, spawn: () => fakeWorker(20) });
    const { deps, snapshot, handle, restoreCalls } = harness(pool, null);
    const methods = createRequestMethods(deps);

    const finished = methods.runRequest(request);
    await vi.advanceTimersByTimeAsync(25);
    handle()?.cancel();
    await vi.advanceTimersByTimeAsync(400);
    await finished;

    expect(restoreCalls.count).toBe(1);
    expect(snapshot().message).toBe(simCopy.stopped);
    expect(snapshot().phase).toBe('error');
    pool.terminate();
    vi.useRealTimers();
  });
});
