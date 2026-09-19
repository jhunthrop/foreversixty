// web/src/lib/sim/bulk-run.test.ts
import { describe, expect, it, vi } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { createFakeWorker } from '../../test-support/fake-worker';
import bulkResultJson from '../../fixtures/sim/bulk-result.json';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import {
  BulkCapError,
  BulkRunError,
  chunk,
  countCombinations,
  runBulk,
  runWeightsRun,
  stageProgressLine,
  type BulkProgress,
} from './bulk-run';
import type { BulkRequest, BulkResult, WeightsRequest, WeightsResult } from './bulk-types';
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
    // A one-worker pool runs the equipped baseline and each combo one request at a time,
    // 60ms apiece (two 30ms ticks). By 135ms the baseline and the first combo have both
    // finished and the second combo is mid-run (its next tick lands at 150ms), which is
    // where the stop lands -- comfortably clear of either boundary.
    const p = pool({ tickMs: 30, ticks: 2, size: 1 });
    const handle = runBulk(p, gearRequest(), () => {});
    await new Promise((resolve) => setTimeout(resolve, 135));
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
});
