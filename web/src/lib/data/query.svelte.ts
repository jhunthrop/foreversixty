// web/src/lib/data/query.svelte.ts
// A Svelte 5 rune wrapper around web/src/lib/data/query.ts: every island named in spec
// 2026-09-23 §3.3 uses this in place of its own $state status + $effect fetch.
import { query, subscribe, type QueryOptions, type QueryState } from './query';

export interface QueryStateHandle<T> {
  readonly data: T | null;
  readonly status: QueryState<T>['status'];
  readonly error: string;
  readonly stale: boolean;
  /** Re-fires `load` in place, ignoring freshness -- what a LoadError's retry button calls. */
  refresh(): void;
}

export function createQueryState<T>(
  key: string,
  load: () => Promise<T>,
  options: QueryOptions<T>,
): QueryStateHandle<T> {
  // 'loading', not 'idle': Svelte 5's $effect never runs during svelte/server's render(), so
  // this initial value IS the server-rendered/pre-hydration output -- it must match every
  // existing island's own `status = $state('loading')` convention (the placeholder that
  // reserves layout before the client ever runs). See the design notes above Task 1.
  let state = $state<QueryState<T>>({ data: null, status: 'loading', error: '', stale: false });

  $effect(() => {
    void query(key, load, options).catch(() => {
      // The rejection already reached `state` via the subscription below; nothing else to do.
    });
    const unsubscribe = subscribe<T>(key, (next) => {
      state = next;
    });
    return unsubscribe;
  });

  return {
    get data() {
      return state.data;
    },
    get status() {
      return state.status;
    },
    get error() {
      return state.error;
    },
    get stale() {
      return state.stale;
    },
    refresh(): void {
      void query(key, load, { ...options, ttlMs: 0 }).catch(() => {});
    },
  };
}

/** A plain-function door into `$effect.root`, for a test file the Svelte preprocessor never
 *  touches (see query.svelte.test.ts). Returns the root's cleanup function, unchanged. */
export function withEffectRoot(fn: () => void): () => void {
  return $effect.root(fn);
}
