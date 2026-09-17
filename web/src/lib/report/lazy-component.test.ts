// web/src/lib/report/lazy-component.test.ts
import { describe, expect, it, vi } from 'vitest';
import { createLazyComponent } from './lazy-component.svelte';

// A stand-in for a Svelte component module: createLazyComponent only ever reads `.default`
// off what the loader resolves, so a plain marker object is enough.
const fakeComponent = {} as never;

function flush(): Promise<void> {
  // Two microtask hops: one for the loader's own promise, one for the `.then`/`.catch`
  // handler chained onto it inside `load()`.
  return Promise.resolve().then(() => Promise.resolve());
}

describe('createLazyComponent', () => {
  it('loads once and caches: a second load() call does not call the loader again', async () => {
    const loader = vi.fn().mockResolvedValue({ default: fakeComponent });
    const lazy = createLazyComponent(loader);

    expect(lazy.current).toBeNull();
    lazy.load();
    expect(loader).toHaveBeenCalledTimes(1);
    lazy.load();
    expect(loader).toHaveBeenCalledTimes(1);
    await flush();

    expect(lazy.current).toBe(fakeComponent);
    expect(lazy.loading).toBe(false);
    expect(lazy.error).toBe('');

    lazy.load();
    expect(loader).toHaveBeenCalledTimes(1);
  });

  it('a rejected loader resets loading and records the error', async () => {
    const loader = vi.fn().mockRejectedValue(new Error('offline'));
    const lazy = createLazyComponent(loader);

    lazy.load();
    expect(lazy.loading).toBe(true);
    await flush();

    expect(lazy.current).toBeNull();
    expect(lazy.loading).toBe(false);
    expect(lazy.error).toBe('This view did not load (offline).');
  });

  it('a later load() call retries rather than short-circuiting, and can succeed', async () => {
    const loader = vi.fn().mockRejectedValueOnce(new Error('offline')).mockResolvedValueOnce({
      default: fakeComponent,
    });
    const lazy = createLazyComponent(loader);

    lazy.load();
    await flush();
    expect(lazy.error).not.toBe('');
    expect(lazy.current).toBeNull();

    lazy.load();
    expect(loader).toHaveBeenCalledTimes(2);
    expect(lazy.loading).toBe(true);
    // The retry clears the previous failure's message immediately, not only once it settles.
    expect(lazy.error).toBe('');
    await flush();

    expect(lazy.current).toBe(fakeComponent);
    expect(lazy.loading).toBe(false);
    expect(lazy.error).toBe('');
  });

  it('a rejection with a non-Error value still resets loading and records a message', async () => {
    const loader = vi.fn().mockRejectedValue('boom');
    const lazy = createLazyComponent(loader);

    lazy.load();
    await flush();

    expect(lazy.loading).toBe(false);
    expect(lazy.error).toBe('This view did not load.');
  });
});
