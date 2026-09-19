// web/src/lib/sim/sim.worker.ts
// One engine instance, loaded once, behind the pool's message protocol.
//
// Progress is the only part that is not a plain request and reply: Go reports it by calling
// a global the host installs, keyed by the callback id the run was started with, so this
// keeps a callbackId-to-token map and translates each callback into a message for the pool.
/// <reference lib="webworker" />
import { loadEngine, type EngineModule } from './engine';
import type { SimProgressUpdate } from './types';
import type { FromWorker, ToWorker } from './worker';

const scope = self as unknown as DedicatedWorkerGlobalScope;

let engine: EngineModule | null = null;
let loading: Promise<EngineModule> | null = null;
const tokenOf = new Map<string, number>();

function engineOnce(): Promise<EngineModule> {
  if (engine !== null) return Promise.resolve(engine);
  loading ??= loadEngine().then((loaded) => {
    loaded.onProgress((callbackId, progressJSON) => {
      const token = tokenOf.get(callbackId);
      if (token === undefined) return;
      reply({ kind: 'progress', token, update: JSON.parse(progressJSON) as SimProgressUpdate });
    });
    engine = loaded;
    return loaded;
  });
  return loading;
}

function reply(message: FromWorker): void {
  scope.postMessage(message);
}

async function handle(message: ToWorker): Promise<void> {
  if (message.kind === 'abort') {
    for (const callbackId of tokenOf.keys()) {
      if (callbackId.startsWith(message.callbackId)) engine?.simAbort(callbackId);
    }
    return;
  }
  const token = message.token;
  try {
    const loaded = await engineOnce();
    if (message.kind === 'split') {
      reply({ kind: 'many', token, results: loaded.simSplit(message.request, message.shards) });
      return;
    }
    if (message.kind === 'combine') {
      reply({ kind: 'one', token, result: loaded.simCombine(message.results) });
      return;
    }
    tokenOf.set(message.callbackId, token);
    try {
      reply({ kind: 'one', token, result: await loaded.simRun(message.request, message.callbackId) });
    } finally {
      tokenOf.delete(message.callbackId);
    }
  } catch (error) {
    reply({ kind: 'failed', token, message: error instanceof Error ? error.message : String(error) });
  }
}

scope.addEventListener('message', (event: MessageEvent<ToWorker>) => {
  void handle(event.data);
});
