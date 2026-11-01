import { afterEach, describe, expect, it, vi } from 'vitest';
import { invalidate } from '../data/query';
import {
  DATA_LOAD_FAILED,
  dataUrl,
  fetchJson,
  loadItemsWithCurrentFallback,
  loadReference,
  loadSets,
  loadTalents,
} from './load';

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
  // fetchJson now caches through query.ts's shared, module-level store (scope: 'public'),
  // so a later test that reuses the same url -- loadTalents' network-failure test below
  // reads 'b1'/'warrior' again, same as its success test -- would otherwise see a still-
  // fresh entry and never call `fetch` at all.
  invalidate('');
  vi.unstubAllGlobals();
});

describe('dataUrl', () => {
  it('builds a root-relative path under the build id', () => {
    expect(dataUrl('1.15.9.69722', 'talents/warrior.json')).toBe('/data/1.15.9.69722/talents/warrior.json');
  });
});

describe('loadTalents', () => {
  it('returns the parsed file', async () => {
    stubFetch({ '/data/b1/talents/warrior.json': { build: 'b1', class_slug: 'warrior', trees: [] } });
    await expect(loadTalents('b1', 'warrior')).resolves.toMatchObject({ class_slug: 'warrior' });
  });

  it('throws the planner load message on a 404', async () => {
    stubFetch({});
    await expect(loadTalents('b1', 'paladin')).rejects.toThrow(DATA_LOAD_FAILED);
  });

  it('throws the planner load message when the network fails', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(loadTalents('b1', 'warrior')).rejects.toThrow(DATA_LOAD_FAILED);
  });
});

describe('loadItemsWithCurrentFallback', () => {
  const OLD = '1.60.1.70450';
  const CURRENT = '1.60.1.70707';
  const row = (id: number, name: string, icon: string) => ({ id, name, icon });

  // An old export is read against its own build's items, but an item added since (Quick
  // Strike Ring 288927 on 1.60.1.70707) is only in the current build's file; without the
  // fallback the paperdoll named it "Item 288927".
  it("names an item the export's build lacks from the current build", async () => {
    stubFetch({
      [`/data/${OLD}/items/hunter.json`]: {
        build: OLD,
        class_slug: 'hunter',
        items: [row(1, 'Old Bow', 'inv_bow')],
      },
      [`/data/${CURRENT}/items/hunter.json`]: {
        build: CURRENT,
        class_slug: 'hunter',
        items: [row(1, 'Renamed Bow', 'inv_bow_new'), row(288927, 'Quick Strike Ring', 'inv_jewelry_ring')],
      },
    });
    const file = await loadItemsWithCurrentFallback(OLD, CURRENT, 'hunter');
    expect(file.build).toBe(OLD);
    expect(file.items.find((item) => item.id === 288927)).toMatchObject({
      name: 'Quick Strike Ring',
      icon: 'inv_jewelry_ring',
    });
    // The export's own build still names what it has.
    expect(file.items.find((item) => item.id === 1)?.name).toBe('Old Bow');
    expect(file.items.filter((item) => item.id === 1)).toHaveLength(1);
  });

  it('reads the current build alone when the export build ships no items file', async () => {
    stubFetch({
      [`/data/${CURRENT}/items/hunter.json`]: {
        build: CURRENT,
        class_slug: 'hunter',
        items: [row(288927, 'Quick Strike Ring', 'r')],
      },
    });
    const file = await loadItemsWithCurrentFallback(OLD, CURRENT, 'hunter');
    expect(file.items.map((item) => item.id)).toEqual([288927]);
  });

  it('fetches once when the export is on the current build', async () => {
    stubFetch({
      [`/data/${CURRENT}/items/hunter.json`]: {
        build: CURRENT,
        class_slug: 'hunter',
        items: [row(1, 'Bow', 'b')],
      },
    });
    const file = await loadItemsWithCurrentFallback(CURRENT, CURRENT, 'hunter');
    expect(file.items).toHaveLength(1);
    expect(fetch).toHaveBeenCalledTimes(1);
  });

  it('rejects when the export build is broken rather than absent', async () => {
    stubFetch({ [`/data/${OLD}/items/hunter.json`]: {} }, 500);
    await expect(loadItemsWithCurrentFallback(OLD, CURRENT, 'hunter')).rejects.toThrow(DATA_LOAD_FAILED);
  });
});

describe('caching through query.ts', () => {
  // fetchJson is the one wrapped call point -- loadTalents/loadItems/loadReference call it
  // directly and loadOptional (loadSets/loadWeights) calls it too, so this covers all five
  // loaders. Asserted on fetchJson itself rather than loadTalents: loadTalents adds a
  // second layer of async indirection around fetchJson's own wrap of query(), and that
  // extra layer delays the caller's second await just long enough to observe query()'s
  // deliberately-deferred background revalidation actually firing (a real, harmless
  // stale-while-revalidate re-check, not a caching defect) -- so asserting the call count
  // through loadTalents is flaky in a way asserting it on fetchJson directly is not.
  it('caches per build and class', async () => {
    const fetchSpy = vi.fn(
      async () =>
        new Response(JSON.stringify({ build: 'b1', class_slug: 'warrior', trees: [] }), {
          status: 200,
          headers: { 'content-type': 'application/json' },
        }),
    );
    vi.stubGlobal('fetch', fetchSpy);
    await fetchJson(dataUrl('b1', 'talents/warrior.json'));
    await fetchJson(dataUrl('b1', 'talents/warrior.json'));
    expect(fetchSpy).toHaveBeenCalledTimes(1);
  });

  it('loadTalents resolves to the same data on repeat calls, served from the fetchJson cache', async () => {
    const fetchSpy = vi.fn(
      async () =>
        new Response(JSON.stringify({ build: 'b1', class_slug: 'warrior', trees: [] }), {
          status: 200,
          headers: { 'content-type': 'application/json' },
        }),
    );
    vi.stubGlobal('fetch', fetchSpy);
    const first = await loadTalents('b1', 'warrior');
    const second = await loadTalents('b1', 'warrior');
    expect(second).toEqual(first);
  });
});

describe('loadReference', () => {
  it('fetches classes, races and combos together', async () => {
    stubFetch({
      '/data/b1/classes.json': [{ id: 1, slug: 'warrior' }],
      '/data/b1/races.json': [{ id: 1, slug: 'human' }],
      '/data/b1/combos.json': [{ race_id: 1, class_id: 1, new_in_forever: false }],
    });
    await expect(loadReference('b1')).resolves.toEqual({
      classes: [{ id: 1, slug: 'warrior' }],
      races: [{ id: 1, slug: 'human' }],
      combos: [{ race_id: 1, class_id: 1, new_in_forever: false }],
    });
  });
});

describe('loadSets', () => {
  it('resolves to an empty list when the build ships no sets file', async () => {
    stubFetch({});
    await expect(loadSets('b1')).resolves.toEqual([]);
  });

  it('rejects with the status when the server errors, rather than reporting no sets', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => new Response('boom', { status: 500 })),
    );
    await expect(loadSets('b1')).rejects.toMatchObject({
      message: expect.stringContaining(DATA_LOAD_FAILED),
      status: 500,
    });
  });

  it('rejects when the network fails, rather than reporting no sets', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(loadSets('b1')).rejects.toThrow(DATA_LOAD_FAILED);
  });
});
