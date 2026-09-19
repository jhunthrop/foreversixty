// web/src/test-support/fake-worker.ts
// One in-process worker speaking exactly the protocol web/src/lib/sim/sim.worker.ts speaks,
// over any EngineModule. Tasks 5, 11 and 20 all drive the real pool through it, so the pool
// and the protocol are tested for real and only the Worker boundary is stood in for.
//
// It is deliberately not a mock: every message is handled the way sim.worker.ts handles it,
// including the callbackId-to-token map that turns an engine progress callback into a
// protocol message, and the prefix rule that makes one abort stop every shard of a run.
import type { EngineModule } from '../lib/sim/engine';
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
            if (id.startsWith(prefix)) engine.simAbort(id);
          }
          return;
        }
        const token = message.token;
        try {
          if (message.kind === 'split') {
            emit({ kind: 'many', token, results: engine.simSplit(message.request, message.shards) });
          } else if (message.kind === 'combine') {
            emit({ kind: 'one', token, result: engine.simCombine(message.results) });
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
