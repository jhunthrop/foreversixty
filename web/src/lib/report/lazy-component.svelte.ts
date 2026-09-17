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

export interface LazyComponent<Props extends Record<string, unknown>> {
  /** The resolved component, or null before `load()` has settled. */
  readonly current: Component<Props> | null;
  /** Starts the import if it has not started yet; safe to call on every render. */
  load: () => void;
}

export function createLazyComponent<Props extends Record<string, unknown>>(
  loader: () => Promise<{ default: Component<Props> }>,
): LazyComponent<Props> {
  let current = $state<Component<Props> | null>(null);
  let loading = false;
  return {
    get current() {
      return current;
    },
    load(): void {
      if (current !== null || loading) return;
      loading = true;
      void loader().then((module) => {
        current = module.default;
      });
    },
  };
}
