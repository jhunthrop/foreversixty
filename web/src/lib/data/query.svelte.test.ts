// @vitest-environment jsdom
// web/src/lib/data/query.svelte.test.ts
// A plain `.test.ts` file is never run through the Svelte preprocessor (only `.svelte` and
// `.svelte.ts` files are, per @astrojs/svelte's file matching), so `$effect.root(...)`
// compiler syntax cannot appear here directly -- it would be a syntax error. `withEffectRoot`
// (exported from the compiled `./query.svelte` module) is the plain-function door into a
// root: it is Svelte's `$effect.root` compiled inside a file the preprocessor does touch.
// The jsdom pragma is required, not cosmetic: without a browser-like environment, vite-plugin-
// svelte compiles query.svelte.ts against Svelte's *server* runtime, where $effect is a no-op
// (see the design notes' "$effect never runs during svelte/server's render()") -- so $effect.root
// would never invoke its callback and every handle below would stay unset.
import { describe, expect, it, vi } from 'vitest';
import { createQueryState, withEffectRoot } from './query.svelte';

function flush(): Promise<void> {
  return Promise.resolve().then(() => Promise.resolve());
}

describe('createQueryState', () => {
  it('starts loading, then reflects the loaded data', async () => {
    let handle: ReturnType<typeof createQueryState<{ n: number }>>;
    const cleanup = withEffectRoot(() => {
      handle = createQueryState('qs1', () => Promise.resolve({ n: 5 }), {
        scope: 'public',
        ttlMs: 60_000,
      });
    });
    expect(handle!.status).toBe('loading');
    await flush();
    expect(handle!.status).toBe('ready');
    expect(handle!.data).toEqual({ n: 5 });
    cleanup();
  });

  it('records a failure', async () => {
    let handle: ReturnType<typeof createQueryState<{ n: number }>>;
    const cleanup = withEffectRoot(() => {
      handle = createQueryState('qs2', () => Promise.reject(new Error('offline')), {
        scope: 'public',
        ttlMs: 60_000,
      });
    });
    await flush();
    expect(handle!.status).toBe('failed');
    expect(handle!.error).toBe('offline');
    cleanup();
  });

  it('refresh() re-fires the same load in place', async () => {
    const load = vi.fn().mockResolvedValue({ n: 1 });
    let handle: ReturnType<typeof createQueryState<{ n: number }>>;
    const cleanup = withEffectRoot(() => {
      handle = createQueryState('qs3', load, { scope: 'public', ttlMs: 60_000 });
    });
    await flush();
    expect(load).toHaveBeenCalledTimes(1);
    handle!.refresh();
    await flush();
    expect(load).toHaveBeenCalledTimes(2);
    cleanup();
  });

  it('unsubscribes on cleanup: a later notification does not touch a destroyed handle', async () => {
    let handle: ReturnType<typeof createQueryState<{ n: number }>>;
    let resolveLoad: (v: { n: number }) => void = () => {};
    const cleanup = withEffectRoot(() => {
      handle = createQueryState(
        'qs4',
        () =>
          new Promise((resolve) => {
            resolveLoad = resolve;
          }),
        { scope: 'public', ttlMs: 60_000 },
      );
    });
    cleanup();
    resolveLoad({ n: 1 });
    await flush();
    expect(handle!.status).toBe('loading'); // frozen at teardown, never touched again
  });
});
