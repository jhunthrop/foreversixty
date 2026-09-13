import { afterEach, describe, expect, it, vi } from 'vitest';
import { DATA_LOAD_FAILED, dataUrl, loadReference, loadSets, loadTalents } from './load';

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

afterEach(() => vi.unstubAllGlobals());

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
});
