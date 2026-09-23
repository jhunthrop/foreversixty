// web/src/lib/sim/sim-items.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { invalidate } from '../data/query';
import { isKnownItem, knownItemIds, loadSimItems } from './sim-items';

function stubFetch(map: Record<string, unknown>, status = 200): void {
  vi.stubGlobal(
    'fetch',
    vi.fn(async (input: string) => {
      const body = map[String(input)];
      if (body === undefined) return new Response('not found', { status: 404 });
      return new Response(JSON.stringify(body), {
        status,
        headers: { 'content-type': 'application/json' },
      });
    }),
  );
}

afterEach(() => {
  // loadSimItems now caches through query.ts's shared, module-level store (scope:
  // 'public', via planner/load.ts's fetchJson), so a later test that reuses the same
  // build's url -- every test in the describe block below reads '/data/b1/simitems.json'
  // -- would otherwise see a still-fresh entry and never call `fetch` at all.
  invalidate('');
  vi.unstubAllGlobals();
});

describe('loadSimItems', () => {
  it('returns the parsed file when the build ships one', async () => {
    stubFetch({ '/data/b1/simitems.json': { build: 'b1', items: [1, 2, 3] } });
    await expect(loadSimItems('b1')).resolves.toEqual({ build: 'b1', items: [1, 2, 3] });
  });

  it('resolves to null when the build ships no simitems.json', async () => {
    stubFetch({});
    await expect(loadSimItems('b1')).resolves.toBeNull();
  });

  it('rejects when the server errors, rather than reporting no file', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => new Response('boom', { status: 500 })),
    );
    await expect(loadSimItems('b1')).rejects.toMatchObject({ status: 500 });
  });
});

describe('knownItemIds', () => {
  it('is null for a build with no simitems.json -- nothing to filter against', () => {
    expect(knownItemIds(null)).toBeNull();
  });

  it('is the file’s ids as a Set, even when the file lists none', () => {
    expect(knownItemIds({ build: 'b1', items: [12640, 16966] })).toEqual(new Set([12640, 16966]));
    expect(knownItemIds({ build: 'b1', items: [] })).toEqual(new Set());
  });
});

describe('isKnownItem', () => {
  it('accepts every id when there is nothing to filter against', () => {
    expect(isKnownItem(16963, null)).toBe(true);
    expect(isKnownItem(21550, null)).toBe(true);
  });

  it('accepts an id the set carries and refuses one it does not', () => {
    const known = new Set([12640, 16963]);
    expect(isKnownItem(12640, known)).toBe(true);
    expect(isKnownItem(21550, known)).toBe(false);
  });

  it('refuses every id against an empty set -- a build whose engine database is genuinely empty', () => {
    expect(isKnownItem(12640, new Set())).toBe(false);
  });
});
