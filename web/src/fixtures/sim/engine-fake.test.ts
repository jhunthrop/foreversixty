import { describe, expect, it, vi } from 'vitest';
import type { SimProgressUpdate, SimRequest, SimResult } from '../../lib/sim/types';
import { DEFAULT_ENCOUNTER } from '../../lib/sim/types';
import { ENGINE_VERSION } from '../../lib/sim/version';
import { createFakeEngine } from './engine-fake';

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
