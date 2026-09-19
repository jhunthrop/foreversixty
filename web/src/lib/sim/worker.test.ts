// web/src/lib/sim/worker.test.ts
// Fix round 1: two protocol bugs the review found, both about what an abort actually reaches.
//
// Neither test touches a real Worker -- createPool's `spawn` option is exactly what lets a
// fake worker (backed by the fake engine, exercising the real message protocol) stand in for
// the Worker boundary, so these run in the default (node) environment.
import { describe, expect, it, vi } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { createFakeWorker } from '../../test-support/fake-worker';
import { DEFAULT_ENCOUNTER, type SimRequest } from './types';
import { ENGINE_VERSION } from './version';
import { createPool } from './worker';

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
  iterations: 40,
  random_seed: 7,
};

const json = (value: SimRequest): string => JSON.stringify(value);

describe('createPool abort', () => {
  it('a numeric callback id does not abort a different run that shares it as a string prefix', async () => {
    vi.useFakeTimers();
    try {
      // One worker, so both runs' single shard land on the same fake worker and its one
      // tokenOf map -- exactly the collision the review found: "sim-1" is a string-prefix of
      // "sim-10-0" and, unanchored, aborting run "sim-1" also caught run "sim-10"'s shard.
      const engine = createFakeEngine({ tickMs: 10, ticks: 4 });
      const pool = createPool({ hardwareConcurrency: 1, spawn: () => createFakeWorker(engine) });

      const run1 = pool.run([json(request)], 'sim-1', () => {});
      const run10 = pool.run([json(request)], 'sim-10', () => {});
      const settled1 = expect(run1).rejects.toThrow(/aborted/);

      await vi.advanceTimersByTimeAsync(15);
      pool.abort('sim-1');
      await vi.advanceTimersByTimeAsync(100);

      await settled1;
      await expect(run10).resolves.toHaveLength(1);
    } finally {
      vi.useRealTimers();
    }
  });

  it('aborts the sibling shards of a run when one shard fails, instead of leaving them running', async () => {
    vi.useFakeTimers();
    try {
      const failing = createFakeEngine({ tickMs: 10, failWith: 'shard 0 blew up' });
      const healthy = createFakeEngine({ tickMs: 10, ticks: 8 });
      const healthyAbort = vi.spyOn(healthy, 'simAbort');
      const pool = createPool({
        hardwareConcurrency: 2,
        spawn: (index) => createFakeWorker(index === 0 ? failing : healthy),
      });

      const shard1Iterations: number[] = [];
      const run = pool.run([json(request), json({ ...request, iterations: 800 })], 'sim-fail', (progress) => {
        if (progress.shard === 1) shard1Iterations.push(progress.iterationsDone);
      });
      const settled = expect(run).rejects.toThrow('shard 0 blew up');

      await vi.advanceTimersByTimeAsync(50);
      await settled;

      // The failing shard's own id is "sim-fail-0"; the sibling worker.ts must reach is
      // "sim-fail-1", found only by aborting the run's own callback id after the rejection.
      expect(healthyAbort).toHaveBeenCalledWith('sim-fail-1');

      // Advance well past the 80ms the healthy shard needs to finish its 8 ticks
      // unaborted; if it were still running, it would report iterationsDone: 800.
      await vi.advanceTimersByTimeAsync(200);
      expect(shard1Iterations).not.toContain(800);
    } finally {
      vi.useRealTimers();
    }
  });
});
