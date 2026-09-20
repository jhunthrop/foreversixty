// web/src/test-support/fake-worker.ts
// One in-process worker speaking exactly the protocol web/src/lib/sim/sim.worker.ts speaks,
// over any EngineModule. Tasks 5, 11 and 20 all drive the real pool through it, so the pool
// and the protocol are tested for real and only the Worker boundary is stood in for.
//
// It is deliberately not a mock: every message is handled the way sim.worker.ts handles it,
// including the callbackId-to-token map that turns an engine progress callback into a
// protocol message, and the prefix rule that makes one abort stop every shard of a run.
import { unwrapOrThrow, type EngineModule } from '../lib/sim/engine';
import { combineInputFromShards, shardsFromSplit } from '../lib/sim/engine-protocol';
import type { SimProgressUpdate } from '../lib/sim/types';
import type { FromWorker, PoolWorker, ToWorker } from '../lib/sim/worker';

export function createFakeWorker(engine: EngineModule): PoolWorker {
  const listeners: ((event: { data: FromWorker }) => void)[] = [];
  const emit = (data: FromWorker): void => listeners.forEach((listen) => listen({ data }));
  const tokenOf = new Map<string, number>();

  engine.onProgress((callbackId, progressJSON) => {
    const token = tokenOf.get(callbackId);
    if (token === undefined) return;
    emit({ kind: 'progress', token, update: JSON.parse(progressJSON) as SimProgressUpdate });
  });

  return {
    addEventListener: (_type, listener) => {
      listeners.push(listener);
    },
    terminate: () => {},
    postMessage: (message: ToWorker) => {
      void (async () => {
        if (message.kind === 'abort') {
          const prefix = `${message.callbackId}-`;
          for (const id of tokenOf.keys()) {
            // Mirrors sim.worker.ts: a run's shards live under `${callbackId}-${index}`
            // (the prefix match); a weights run is registered under its own callbackId
            // verbatim, so it needs the exact-match branch too.
            if (id === message.callbackId || id.startsWith(prefix)) engine.simAbort(id);
          }
          return;
        }
        const token = message.token;
        try {
          if (message.kind === 'split') {
            emit({
              kind: 'many',
              token,
              results: shardsFromSplit(engine.simSplit(message.request, message.shards)),
            });
          } else if (message.kind === 'combine') {
            emit({ kind: 'one', token, result: engine.simCombine(combineInputFromShards(message.results)) });
          } else if (message.kind === 'needsMore') {
            emit({
              kind: 'one',
              token,
              result: unwrapOrThrow(engine.simNeedsMore(message.result, message.request)),
            });
          } else if (message.kind === 'validate') {
            emit({ kind: 'one', token, result: unwrapOrThrow(engine.simValidate(message.request)) });
          } else if (message.kind === 'count') {
            emit({ kind: 'one', token, result: engine.simCount(message.request) });
          } else if (message.kind === 'plan') {
            // Not unwrapped: `cap_exceeded` is a legal answer simPlan gives, not a failure
            // (mirrors `count` right above, and engine.ts's own simPlan doc).
            emit({ kind: 'one', token, result: engine.simPlan(message.request) });
          } else if (message.kind === 'rank') {
            emit({
              kind: 'one',
              token,
              result: unwrapOrThrow(engine.simRank(message.request, message.stage, message.results)),
            });
          } else if (message.kind === 'weights') {
            tokenOf.set(message.callbackId, token);
            try {
              emit({
                kind: 'one',
                token,
                result: await engine.simWeights(message.request, message.callbackId),
              });
            } finally {
              tokenOf.delete(message.callbackId);
            }
          } else {
            tokenOf.set(message.callbackId, token);
            try {
              emit({
                kind: 'one',
                token,
                result: await engine.simRun(message.request, message.callbackId),
              });
            } finally {
              tokenOf.delete(message.callbackId);
            }
          }
        } catch (error) {
          emit({ kind: 'failed', token, message: (error as Error).message });
        }
      })();
    },
  };
}

/**
 * The failure-path counterpart to createFakeWorker: a worker that answers every run with
 * `{kind: 'failed'}`, for the "an engine failure surfaces as a failure" tests -- previously
 * a byte-identical 10-line fixture hand-rolled at `run.test.ts:124` and
 * `live-dps.test.ts:106` (L8, final whole-branch review).
 */
export function createBrokenWorker(message = 'wasm trap'): PoolWorker {
  let emit: (data: FromWorker) => void = () => {};
  return {
    addEventListener: (_type, listener) => {
      emit = (data) => listener({ data });
    },
    terminate: () => {},
    postMessage: (toWorker: ToWorker) => {
      if (toWorker.kind === 'abort') return;
      emit({ kind: 'failed', token: toWorker.token, message });
    },
  };
}
