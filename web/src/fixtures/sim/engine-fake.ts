// web/src/fixtures/sim/engine-fake.ts
// The contract's stub for sim.wasm, with the same four-function surface, so swapping the
// real one in is one environment variable and nothing else.
//
// Five behaviours it has to get right, because the UI can see all five:
//   * iterations split and sum exactly, including the remainder;
//   * progress arrives in ticks, so the DPS figure on screen really does move;
//   * the same seed gives the same numbers, so a paired run is paired;
//   * an abort rejects the pending run rather than resolving it;
//   * simCombine pools the partials the way the real one does, and the combined result
//     carries the fixture summary, so every report component mounts.
//
// The numbers are drawn from the fixture's own mean and standard deviation with a seeded
// generator rather than randomly, so a Playwright assertion on the DPS figure is stable.
import type { EngineModule, ProgressHandler } from '../../lib/sim/engine';
import type { SimRequest, SimResult } from '../../lib/sim/types';
import fixtureResultJson from './result.json';

const fixture = fixtureResultJson as unknown as SimResult;

export interface FakeEngineOptions {
  tickMs?: number;
  ticks?: number;
  /**
   * Makes every `simRun` reject with this exact message instead of running. Task 5's
   * "keeps the engine's own words" test uses it to prove the message survives every hop to
   * `SimRunError.detail` -- the real wasm answers an unknown buff id with
   * `request: unknown buff: "battle-shout"`, and that sentence is the only thing that
   * tells a player what to change.
   */
  failWith?: string;
}

function seeded(seed: number): () => number {
  let state = seed >>> 0;
  return () => {
    state = (state + 0x6d2b79f5) >>> 0;
    let t = state;
    t = Math.imul(t ^ (t >>> 15), t | 1);
    t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
    return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
  };
}

function normal(random: () => number, mean: number, stddev: number): number {
  const u = Math.max(random(), Number.EPSILON);
  const v = random();
  return mean + stddev * Math.sqrt(-2 * Math.log(u)) * Math.cos(2 * Math.PI * v);
}

function statsOf(samples: number[]): SimResult['dps'] & { n: number } {
  const n = samples.length;
  const mean = samples.reduce((a, b) => a + b, 0) / n;
  const variance = samples.reduce((a, b) => a + (b - mean) ** 2, 0) / n;
  const stddev = Math.sqrt(variance);
  return {
    n,
    mean,
    stddev,
    error: n > 0 ? stddev / Math.sqrt(n) : 0,
    min: Math.min(...samples),
    max: Math.max(...samples),
  };
}

const delay = (ms: number): Promise<void> =>
  ms <= 0 ? Promise.resolve() : new Promise((resolve) => setTimeout(resolve, ms));

/** The `{"error": "..."}` envelope simNeedsMore, simValidate and simCount all answer with
 * for input that is not JSON at all -- one place for the message-extraction logic. */
const errorEnvelope = (error: unknown): string =>
  JSON.stringify({ error: error instanceof Error ? error.message : String(error) });

export function createFakeEngine(options: FakeEngineOptions = {}): EngineModule {
  const tickMs = options.tickMs ?? 120;
  const ticks = options.ticks ?? 10;
  const failWith = options.failWith ?? '';
  const aborted = new Set<string>();
  // Tracks a run's callback id for exactly as long as simRun is in flight, so simAbort can
  // answer {"aborted": false} for an id nothing registered -- main.go's own distinction
  // between "you stopped it" and "you called this wrong".
  const active = new Set<string>();
  let progress: ProgressHandler = () => {};

  return {
    onProgress(handler) {
      progress = handler;
    },

    simSplit(requestJSON, n) {
      const request = JSON.parse(requestJSON) as SimRequest;
      const used = Math.max(1, Math.min(n, request.iterations));
      const base = Math.floor(request.iterations / used);
      const remainder = request.iterations % used;
      const parts: SimRequest[] = Array.from({ length: used }, (_, i) => ({
        ...request,
        iterations: base + (i < remainder ? 1 : 0),
        random_seed: request.random_seed === 0 ? i + 1 : request.random_seed * 1000 + i,
      }));
      return JSON.stringify(parts);
    },

    async simRun(requestJSON, callbackId) {
      aborted.delete(callbackId);
      active.add(callbackId);
      try {
        if (failWith !== '') {
          await delay(tickMs);
          throw new Error(failWith);
        }
        const request = JSON.parse(requestJSON) as SimRequest;
        const random = seeded(request.random_seed * 7919 + request.iterations);
        const samples: number[] = [];
        const perTick = Math.max(1, Math.ceil(request.iterations / ticks));
        const startedAt = Date.now();

        while (samples.length < request.iterations) {
          await delay(tickMs);
          if (aborted.has(callbackId)) {
            aborted.delete(callbackId);
            throw new Error(`sim run ${callbackId} aborted`);
          }
          const upTo = Math.min(request.iterations, samples.length + perTick);
          while (samples.length < upTo) {
            samples.push(normal(random, fixture.dps.mean, fixture.dps.stddev));
          }
          const { n, ...dps } = statsOf(samples);
          progress(callbackId, JSON.stringify({ iterations_run: n, dps }));
        }

        const { n, ...dps } = statsOf(samples);
        return JSON.stringify({
          engine_version: request.engine_version,
          request,
          lane: 'browser',
          dps,
          iterations_run: n,
          duration_ms: Date.now() - startedAt,
          summary: fixture.summary,
        } satisfies SimResult);
      } finally {
        active.delete(callbackId);
      }
    },

    simCombine(resultsJSON) {
      const parts = JSON.parse(resultsJSON) as SimResult[];
      let n = 0;
      let sum = 0;
      let sumSquares = 0;
      let min = Number.POSITIVE_INFINITY;
      let max = Number.NEGATIVE_INFINITY;
      let durationMs = 0;
      for (const part of parts) {
        n += part.iterations_run;
        sum += part.dps.mean * part.iterations_run;
        sumSquares += part.iterations_run * (part.dps.stddev ** 2 + part.dps.mean ** 2);
        min = Math.min(min, part.dps.min);
        max = Math.max(max, part.dps.max);
        durationMs = Math.max(durationMs, part.duration_ms);
      }
      const mean = n === 0 ? 0 : sum / n;
      const variance = n === 0 ? 0 : Math.max(0, sumSquares / n - mean * mean);
      const stddev = Math.sqrt(variance);
      const first = parts[0];
      return JSON.stringify({
        engine_version: first.engine_version,
        request: { ...first.request, iterations: n },
        lane: 'browser',
        dps: {
          mean,
          stddev,
          error: n > 0 ? stddev / Math.sqrt(n) : 0,
          min: n === 0 ? 0 : min,
          max: n === 0 ? 0 : max,
        },
        iterations_run: n,
        duration_ms: durationMs,
        summary: fixture.summary,
      } satisfies SimResult);
    },

    simAbort(callbackId) {
      const registered = active.has(callbackId);
      if (registered) aborted.add(callbackId);
      return JSON.stringify({ aborted: registered });
    },

    simNeedsMore(resultJSON, requestJSON) {
      try {
        const result = JSON.parse(resultJSON) as SimResult;
        const request = JSON.parse(requestJSON) as SimRequest;
        const target = request.target_error ?? 0;
        // Four ways to be done, and the real engine agrees on all four: this was not a
        // target-error run; nothing has been measured, so a relative error is undefined;
        // the band is inside the target; or the ceiling is reached.
        const needsMore =
          target > 0 &&
          result.dps.mean > 0 &&
          result.dps.error / result.dps.mean > target &&
          result.iterations_run < request.iterations;
        return JSON.stringify({ needs_more: needsMore });
      } catch (error) {
        return errorEnvelope(error);
      }
    },

    simValidate(requestJSON) {
      let request: SimRequest;
      try {
        request = JSON.parse(requestJSON) as SimRequest;
      } catch (error) {
        return errorEnvelope(error);
      }
      // A short stand-in for api.SimRequest.Validate: the checks the drawer's own tests
      // exercise. The real wasm runs the whole thing, and the drawer renders whatever
      // fields come back, so a fake that refuses fewer things cannot make the page wrong.
      const errors: { field: string; message: string }[] = [];
      if (request.engine_version === undefined || request.engine_version === '') {
        errors.push({ field: 'engine_version', message: 'engine_version is required' });
      }
      if (request.spec === undefined || request.spec === '') {
        errors.push({ field: 'spec', message: 'spec is required' });
      }
      if (!(request.iterations > 0)) {
        errors.push({ field: 'iterations', message: 'iterations must be a positive number' });
      }
      const targets = request.encounter?.targets;
      if (targets !== undefined && (targets < 1 || targets > 10)) {
        errors.push({ field: 'encounter.targets', message: 'targets must be between 1 and 10' });
      }
      return JSON.stringify({ ok: errors.length === 0, errors });
    },

    simCount(requestJSON) {
      let request: SimRequest;
      try {
        request = JSON.parse(requestJSON) as SimRequest;
      } catch (error) {
        return errorEnvelope(error);
      }
      const bulk = request.bulk;
      if (bulk === undefined) return JSON.stringify({ combinations: 0 });
      // A stand-in for `sim/bulk`'s own expansion, which knows slot fit, unique-equipped
      // and weapon shapes; the real wasm counts properly. The fake counts what it can see
      // -- one combination per candidate, per talent loadout, per set -- which is enough
      // for the page's cap notice and its e2e to be exercised honestly.
      const combinations =
        (bulk.candidates?.length ?? 0) + (bulk.talents?.length ?? 0) + (bulk.sets?.length ?? 0);
      if (bulk.cap > 0 && combinations > bulk.cap) {
        return JSON.stringify({ error: 'cap_exceeded', cap: bulk.cap, combinations });
      }
      return JSON.stringify({ combinations });
    },
  };
}
