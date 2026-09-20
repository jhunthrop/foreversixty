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
import type { RankAnswer, StageRequests } from './bulk-types';
import type { ShardProgress } from './estimate';
import { unwrapOrThrow, type CountAnswer, type PlanAnswer, type RequestValidation } from './engine';
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
  | { kind: 'needsMore'; token: number; result: string; request: string }
  | { kind: 'validate'; token: number; request: string }
  | { kind: 'count'; token: number; request: string }
  | { kind: 'plan'; token: number; request: string }
  | { kind: 'rank'; token: number; request: string; stage: string; results: string }
  | { kind: 'weights'; token: number; callbackId: string; request: string }
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
  /** Whether a target-error run has another step to do. The engine decides, not the page. */
  needsMore(result: string, request: string): Promise<boolean>;
  /** `api.SimRequest.Validate` over an edited request. */
  validate(request: string): Promise<RequestValidation>;
  /** How many combinations a bulk request expands to; a cap breach is an answer, not a throw. */
  count(request: string): Promise<CountAnswer>;
  /** The bulk planner's first stage, on worker 0. A cap breach is an answer, not a throw. */
  plan(request: string): Promise<PlanAnswer>;
  /** Scores a finished bulk stage, on worker 0. `next` for another stage, `result` for the last. */
  rank(request: string, stage: string, results: string): Promise<RankAnswer>;
  /** A weights run, on worker 0, progress through the same shard callback a run uses. */
  weights(
    request: string,
    callbackId: string,
    onProgress: (progress: ShardProgress) => void,
  ): Promise<string>;
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
      const shardResults = shards.map((request, index) =>
        send<string>(
          index % size,
          (token) => ({ kind: 'run', token, callbackId: `${callbackId}-${index}`, request }),
          (progress) => onProgress({ ...progress, shard: index }),
        ),
      );
      // One shard failing (an engine error, not a user abort) leaves the rest of this run's
      // shards with nothing telling them to stop, since Promise.all settles on the first
      // rejection but every other worker just keeps computing and posting progress into a
      // closure the caller has already treated as dead. Aborting the run's own id catches
      // every sibling shard through the same <callbackId>-<index> prefix rule the pool gives
      // them, so a failure winds the whole run down instead of leaving orphans running.
      return Promise.all(shardResults).catch((error: unknown) => {
        abort(callbackId);
        throw error;
      });
    },
    combine(results) {
      return send<string>(0, (token) => ({ kind: 'combine', token, results: [...results] }));
    },
    async needsMore(result, request) {
      // The fake lane's answer has not necessarily been through unwrapOrThrow already (see
      // fake-worker.ts): a malformed envelope must fail here no matter which lane produced it.
      const answer = unwrapOrThrow(
        await send<string>(0, (token) => ({ kind: 'needsMore', token, result, request })),
      );
      return (JSON.parse(answer) as { needs_more?: boolean }).needs_more === true;
    },
    async validate(request) {
      const answer = unwrapOrThrow(await send<string>(0, (token) => ({ kind: 'validate', token, request })));
      return JSON.parse(answer) as RequestValidation;
    },
    async count(request) {
      const answer = await send<string>(0, (token) => ({ kind: 'count', token, request }));
      const parsed = JSON.parse(answer) as {
        error?: string;
        cap?: number;
        combinations?: number;
      };
      // `cap_exceeded` is the one error envelope on this lane that is a real answer; any
      // other error from the engine is a genuine failure and is thrown as one.
      if (parsed.error === 'cap_exceeded') {
        return { ok: false, cap: parsed.cap ?? 0, combinations: parsed.combinations ?? 0 };
      }
      if (parsed.error !== undefined) throw new Error(parsed.error);
      return { ok: true, combinations: parsed.combinations ?? 0 };
    },
    async plan(request) {
      const answer = await send<string>(0, (token) => ({ kind: 'plan', token, request }));
      const parsed = JSON.parse(answer) as {
        error?: string;
        cap?: number;
        combinations?: number;
      };
      // `cap_exceeded` is the one error envelope simPlan answers with rather than throws;
      // every other failure is a genuine `{"error": "..."}` (see engine.ts's simPlan doc).
      if (parsed.error === 'cap_exceeded') {
        return { ok: false, cap: parsed.cap ?? 0, combinations: parsed.combinations ?? 0 };
      }
      if (parsed.error !== undefined) throw new Error(parsed.error);
      return { ok: true, stage: parsed as unknown as StageRequests };
    },
    async rank(request, stage, results) {
      const answer = unwrapOrThrow(
        await send<string>(0, (token) => ({ kind: 'rank', token, request, stage, results })),
      );
      return JSON.parse(answer) as RankAnswer;
    },
    weights(request, callbackId, onProgress) {
      return send<string>(
        0,
        (token) => ({ kind: 'weights', token, callbackId, request }),
        (progress) => onProgress({ ...progress, shard: 0 }),
      );
    },
    abort,
    terminate,
  };
}
