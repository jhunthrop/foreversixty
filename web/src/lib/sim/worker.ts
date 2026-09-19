// web/src/lib/sim/worker.ts
// The browser lane's pool: min(8, hardwareConcurrency) workers, one engine instance each,
// every message a JSON string.
//
// simSplit and simCombine are wasm exports and the main thread holds no instance, so both
// are routed to worker 0 -- loading a second copy of the module on the main thread would
// cost the sim page's LCP budget for nothing. run.ts still decides what is combined and
// when, which is the ownership the contract's Web section specifies.
//
// The pool tears itself down on pagehide, so a navigation mid-run leaves no workers behind.
import type { ShardProgress } from './estimate';
import type { SimProgressUpdate } from './types';

export const MAX_WORKERS = 8;
const DEFAULT_CORES = 4;

export function poolSize(hardwareConcurrency: number | undefined): number {
  const cores =
    typeof hardwareConcurrency === 'number' && Number.isFinite(hardwareConcurrency) && hardwareConcurrency > 0
      ? Math.floor(hardwareConcurrency)
      : DEFAULT_CORES;
  return Math.max(1, Math.min(MAX_WORKERS, cores));
}

export type ToWorker =
  | { kind: 'split'; token: number; request: string; shards: number }
  | { kind: 'run'; token: number; callbackId: string; request: string }
  | { kind: 'combine'; token: number; results: string[] }
  | { kind: 'abort'; callbackId: string };

export type FromWorker =
  | { kind: 'progress'; token: number; update: SimProgressUpdate }
  | { kind: 'one'; token: number; result: string }
  | { kind: 'many'; token: number; results: string[] }
  | { kind: 'failed'; token: number; message: string };

export interface PoolWorker {
  postMessage(message: ToWorker): void;
  addEventListener(type: 'message', listener: (event: { data: FromWorker }) => void): void;
  terminate(): void;
}

export interface PoolOptions {
  hardwareConcurrency?: number;
  spawn?: (index: number) => PoolWorker;
}

export interface SimPool {
  readonly size: number;
  split(request: string, shards: number): Promise<string[]>;
  run(
    shards: readonly string[],
    callbackId: string,
    onProgress: (progress: ShardProgress) => void,
  ): Promise<string[]>;
  combine(results: readonly string[]): Promise<string>;
  abort(callbackId: string): void;
  terminate(): void;
}

interface Pending {
  resolve: (value: never) => void;
  reject: (error: Error) => void;
  onProgress?: (progress: ShardProgress) => void;
  shard: number;
}

function defaultSpawn(): PoolWorker {
  return new Worker(new URL('./sim.worker.ts', import.meta.url), {
    type: 'module',
    name: 'forever-sim',
  }) as unknown as PoolWorker;
}

export function createPool(options: PoolOptions = {}): SimPool {
  const size = poolSize(options.hardwareConcurrency ?? globalThis.navigator?.hardwareConcurrency);
  const spawn = options.spawn ?? defaultSpawn;

  const workers: PoolWorker[] = [];
  const pending = new Map<number, Pending>();
  let nextToken = 1;
  let stopped = false;

  function workerAt(index: number): PoolWorker {
    let worker = workers[index];
    if (worker === undefined) {
      worker = spawn(index);
      worker.addEventListener('message', (event) => {
        const message = event.data;
        const entry = pending.get(message.token);
        if (entry === undefined) return;
        if (message.kind === 'progress') {
          entry.onProgress?.({
            shard: entry.shard,
            iterationsDone: message.update.iterations_run,
            estimate: message.update.dps,
          });
          return;
        }
        pending.delete(message.token);
        if (message.kind === 'failed') entry.reject(new Error(message.message));
        else if (message.kind === 'one') (entry.resolve as (v: string) => void)(message.result);
        else (entry.resolve as (v: string[]) => void)(message.results);
      });
      workers[index] = worker;
    }
    return worker;
  }

  function send<T>(
    index: number,
    build: (token: number) => ToWorker,
    onProgress?: (progress: ShardProgress) => void,
  ): Promise<T> {
    if (stopped) return Promise.reject(new Error('sim pool: terminated'));
    const token = nextToken;
    nextToken += 1;
    return new Promise<T>((resolve, reject) => {
      pending.set(token, {
        resolve: resolve as unknown as (value: never) => void,
        reject,
        onProgress,
        shard: index,
      });
      workerAt(index).postMessage(build(token));
    });
  }

  function abort(callbackId: string): void {
    for (const worker of workers) worker?.postMessage({ kind: 'abort', callbackId });
  }

  function terminate(): void {
    stopped = true;
    for (const entry of pending.values()) entry.reject(new Error('sim pool: terminated'));
    pending.clear();
    for (const worker of workers) worker.terminate();
    workers.length = 0;
    globalThis.removeEventListener?.('pagehide', onPageHide);
  }

  function onPageHide(): void {
    terminate();
  }
  globalThis.addEventListener?.('pagehide', onPageHide);

  return {
    get size() {
      return size;
    },
    split(request, shards) {
      return send<string[]>(0, (token) => ({ kind: 'split', token, request, shards }));
    },
    run(shards, callbackId, onProgress) {
      return Promise.all(
        shards.map((request, index) =>
          send<string>(
            index % size,
            (token) => ({ kind: 'run', token, callbackId: `${callbackId}-${index}`, request }),
            (progress) => onProgress({ ...progress, shard: index }),
          ),
        ),
      );
    },
    combine(results) {
      return send<string>(0, (token) => ({ kind: 'combine', token, results: [...results] }));
    },
    abort,
    terminate,
  };
}
