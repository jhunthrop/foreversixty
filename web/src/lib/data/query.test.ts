// web/src/lib/data/query.test.ts
// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { forgetPrivate, invalidate, query, sessionHinted, setQueryData, subscribe } from './query';

function memoryStorage(): Storage {
  const map = new Map<string, string>();
  return {
    get length() {
      return map.size;
    },
    clear: () => map.clear(),
    getItem: (k) => map.get(k) ?? null,
    key: (i) => [...map.keys()][i] ?? null,
    removeItem: (k) => void map.delete(k),
    setItem: (k, v) => void map.set(k, v),
  };
}

function hashFor(key: string): string {
  let h = 0x811c9dc5;
  for (let i = 0; i < key.length; i++) {
    h ^= key.charCodeAt(i);
    h = Math.imul(h, 0x01000193);
  }
  return (h >>> 0).toString(16);
}

beforeEach(() => {
  vi.stubGlobal('localStorage', memoryStorage());
  document.cookie = 'fs_csrf=; Max-Age=0; path=/';
});
afterEach(() => {
  vi.unstubAllGlobals();
  document.cookie = 'fs_csrf=; Max-Age=0; path=/';
});

describe('query()', () => {
  it('dedupes concurrent calls for the same key into one in-flight load', async () => {
    const load = vi.fn().mockResolvedValue({ n: 1 });
    const [a, b] = await Promise.all([
      query('k1', load, { scope: 'public', ttlMs: 60_000 }),
      query('k1', load, { scope: 'public', ttlMs: 60_000 }),
    ]);
    expect(load).toHaveBeenCalledTimes(1);
    expect(a).toEqual({ n: 1 });
    expect(b).toEqual({ n: 1 });
  });

  it('serves a fresh cached answer instantly and revalidates in the background', async () => {
    let call = 0;
    const load = vi.fn().mockImplementation(() => Promise.resolve({ n: ++call }));
    const first = await query('k2', load, { scope: 'public', ttlMs: 60_000 });
    expect(first).toEqual({ n: 1 });

    const second = await query('k2', load, { scope: 'public', ttlMs: 60_000 });
    expect(second).toEqual({ n: 1 }); // instant, from cache, before the background load below settles
    await Promise.resolve().then(() => Promise.resolve());
    expect(load).toHaveBeenCalledTimes(2); // the background revalidation ran
  });

  it('does not show a stored entry older than its ttlMs; treats it as a first load', async () => {
    let now = 0;
    const realNow = Date.now;
    vi.spyOn(Date, 'now').mockImplementation(() => now);
    const load = vi.fn().mockResolvedValue({ n: 1 });
    await query('k3', load, { scope: 'public', ttlMs: 1000 });
    now = 2000;
    const load2 = vi.fn().mockResolvedValue({ n: 2 });
    const result = await query('k3', load2, { scope: 'public', ttlMs: 1000 });
    expect(result).toEqual({ n: 2 });
    expect(load2).toHaveBeenCalledTimes(1);
    Date.now = realNow;
  });

  it('a failed background revalidation keeps the stale answer and does not throw', async () => {
    const load = vi.fn().mockResolvedValueOnce({ n: 1 }).mockRejectedValueOnce(new Error('offline'));
    await query('k4', load, { scope: 'public', ttlMs: 60_000 });
    const second = await query('k4', load, { scope: 'public', ttlMs: 60_000 });
    expect(second).toEqual({ n: 1 });
    await Promise.resolve().then(() => Promise.resolve());
    // no unhandled rejection, no throw: the test completing is the assertion
  });

  it('a failed first load rejects', async () => {
    const load = vi.fn().mockRejectedValue(new Error('offline'));
    await expect(query('k5', load, { scope: 'public', ttlMs: 60_000 })).rejects.toThrow('offline');
  });

  it('persists a public entry and hydrates a fresh module instance from it', async () => {
    const load = vi.fn().mockResolvedValue({ n: 7 });
    await query('k6', load, { scope: 'public', ttlMs: 60_000 });
    vi.resetModules();
    const fresh = await import('./query');
    const load2 = vi.fn().mockResolvedValue({ n: 999 });
    const result = await fresh.query('k6', load2, { scope: 'public', ttlMs: 60_000 });
    expect(result).toEqual({ n: 7 }); // hydrated from storage, no reload needed (still fresh)
    expect(load2).not.toHaveBeenCalled();
  });

  it('only reads a persisted private entry while the session cookie is present', async () => {
    document.cookie = 'fs_csrf=tok; path=/';
    const load = vi.fn().mockResolvedValue({ n: 3 });
    await query('k7', load, { scope: 'private', ttlMs: 60_000 });
    document.cookie = 'fs_csrf=; Max-Age=0; path=/';
    vi.resetModules();
    const fresh = await import('./query');
    const load2 = vi.fn().mockResolvedValue({ n: 4 });
    const result = await fresh.query('k7', load2, { scope: 'private', ttlMs: 60_000 });
    expect(result).toEqual({ n: 4 }); // cookie gone: the snapshot is not trusted, reloaded
    expect(load2).toHaveBeenCalledTimes(1);
  });

  it('drops a persisted entry on a version mismatch', async () => {
    const load = vi.fn().mockResolvedValue({ n: 1 });
    await query('k8', load, { scope: 'public', ttlMs: 60_000, version: 1 });
    vi.resetModules();
    const fresh = await import('./query');
    const load2 = vi.fn().mockResolvedValue({ n: 2 });
    const result = await fresh.query('k8', load2, { scope: 'public', ttlMs: 60_000, version: 2 });
    expect(result).toEqual({ n: 2 });
    expect(load2).toHaveBeenCalledTimes(1);
  });

  it('does not persist an entry over 256 KB', async () => {
    const big = { s: 'x'.repeat(300 * 1024) };
    await query('k9', () => Promise.resolve(big), { scope: 'public', ttlMs: 60_000 });
    expect(localStorage.getItem(`fs.q.${hashFor('k9')}`)).toBeNull();
  });

  it('evicts the oldest persisted entry once more than 64 exist', async () => {
    let now = 0;
    vi.spyOn(Date, 'now').mockImplementation(() => now);
    for (let i = 0; i < 65; i++) {
      now = i;
      await query(`evict-${i}`, () => Promise.resolve({ i }), { scope: 'public', ttlMs: 60_000 });
    }
    const keys = Object.keys(localStorage).filter((k) => k.startsWith('fs.q.'));
    expect(keys.length).toBeLessThanOrEqual(64);
    expect(localStorage.getItem(`fs.q.${hashFor('evict-0')}`)).toBeNull(); // oldest evicted
    vi.restoreAllMocks();
  });

  it('subscribe() replays the current state synchronously, then follows updates', async () => {
    const states: string[] = [];
    const unsubscribe = subscribe('k10', (s) => states.push(s.status));
    expect(states).toEqual(['idle']);
    await query('k10', () => Promise.resolve({ n: 1 }), { scope: 'public', ttlMs: 60_000 });
    expect(states).toEqual(['idle', 'loading', 'ready']);
    unsubscribe();
  });

  it('setQueryData writes a value directly, visible to the next query() call', async () => {
    setQueryData('k11', { n: 42 });
    const load = vi.fn().mockResolvedValue({ n: 0 });
    const result = await query('k11', load, { scope: 'public', ttlMs: 60_000 });
    expect(result).toEqual({ n: 42 });
    expect(load).not.toHaveBeenCalled();
  });

  it('invalidate(prefix) forces the next query() call to reload', async () => {
    const load = vi.fn().mockResolvedValue({ n: 1 });
    await query('list/a', load, { scope: 'public', ttlMs: 60_000 });
    invalidate('list/');
    const load2 = vi.fn().mockResolvedValue({ n: 2 });
    const result = await query('list/a', load2, { scope: 'public', ttlMs: 60_000 });
    expect(result).toEqual({ n: 2 });
    expect(load2).toHaveBeenCalledTimes(1);
  });

  it('forgetPrivate clears private entries but leaves public ones', async () => {
    document.cookie = 'fs_csrf=tok; path=/';
    await query('priv', () => Promise.resolve({ n: 1 }), { scope: 'private', ttlMs: 60_000 });
    await query('pub', () => Promise.resolve({ n: 2 }), { scope: 'public', ttlMs: 60_000 });
    forgetPrivate();
    const loadPriv = vi.fn().mockResolvedValue({ n: 3 });
    const loadPub = vi.fn().mockResolvedValue({ n: 4 });
    await query('priv', loadPriv, { scope: 'private', ttlMs: 60_000 });
    await query('pub', loadPub, { scope: 'public', ttlMs: 60_000 });
    expect(loadPriv).toHaveBeenCalledTimes(1);
    expect(loadPub).not.toHaveBeenCalled();
  });

  it('a 401 on a private load forgets every private entry', async () => {
    document.cookie = 'fs_csrf=tok; path=/';
    await query('priv-a', () => Promise.resolve({ n: 1 }), { scope: 'private', ttlMs: 60_000 });
    class Unauthorized extends Error {
      status = 401;
    }
    const load = vi.fn().mockRejectedValue(new Unauthorized('no'));
    await expect(query('priv-b', load, { scope: 'private', ttlMs: 0 })).rejects.toThrow();
    const loadAgain = vi.fn().mockResolvedValue({ n: 9 });
    await query('priv-a', loadAgain, { scope: 'private', ttlMs: 60_000 });
    expect(loadAgain).toHaveBeenCalledTimes(1); // forgotten by the 401, so this is a real load
  });
});

describe('sessionHinted', () => {
  it('is true only when fs_csrf is present', () => {
    expect(sessionHinted('fs_csrf=abc; other=1')).toBe(true);
    expect(sessionHinted('other=1')).toBe(false);
  });
});
