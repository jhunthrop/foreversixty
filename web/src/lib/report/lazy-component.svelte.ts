// web/src/lib/report/lazy-component.svelte.ts
// A tiny cache around a dynamically imported Svelte component, for the report island's
// non-landing modes and views (Compare, Mechanics, Rankings, Timelines, Events, Queries):
// each one ships as its own chunk instead of the entry bundle everyone pays for on load,
// and is fetched the first time its mode or view is actually selected. `import()` on an
// already-loaded module resolves instantly from the browser's own module cache, but doing
// that inside a template on every reactive pass would still round-trip a microtask and
// unmount/remount the component; caching the resolved component here keeps one stable
// instance for the life of the page, the same identity an eagerly imported component would
// have had across the mode or view being switched away and back.
import type { Component } from 'svelte';

/** The load/error/retry surface, without the resolved component: what a fallback renders. */
export interface LazyLoadState {
  /** Set when the last `load()` attempt failed; cleared the moment a retry starts. */
  readonly error: string;
  /** True while an import is in flight. */
  readonly loading: boolean;
  /** Starts the import if one is not already in flight; safe to call on every render.
   *  After a failure, `current` is still null and `loading` is still false, so the next
   *  call retries rather than short-circuiting. */
  load: () => void;
}

export interface LazyComponent<Props extends Record<string, unknown>> extends LazyLoadState {
  /** The resolved component, or null before `load()` has settled. */
  readonly current: Component<Props> | null;
}

export function createLazyComponent<Props extends Record<string, unknown>>(
  loader: () => Promise<{ default: Component<Props> }>,
): LazyComponent<Props> {
  let current = $state<Component<Props> | null>(null);
  let loading = $state(false);
  let error = $state('');
  return {
    get current() {
      return current;
    },
    get error() {
      return error;
    },
    get loading() {
      return loading;
    },
    load(): void {
      if (current !== null || loading) return;
      loading = true;
      error = '';
      void loader()
        .then((module) => {
          current = module.default;
        })
        .catch((thrown: unknown) => {
          // Offline after the entry loaded, or a chunk evicted from the browser's cache: the
          // mode or view must not stay a permanent blank panel. The caller renders `error`
          // with a retry that calls `load()` again -- `loading` is reset below regardless of
          // outcome, so that retry is a normal, un-short-circuited attempt.
          error = `This view did not load${thrown instanceof Error ? ` (${thrown.message})` : ''}.`;
        })
        .finally(() => {
          loading = false;
        });
    },
  };
}
