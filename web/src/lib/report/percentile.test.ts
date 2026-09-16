// web/src/lib/report/percentile.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { createPercentileLoader, percentileKey } from './percentile';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

const query = {
  encounterId: 9001,
  difficulty: 8,
  spec: 'Protection',
  phase: 'raids-1',
  metric: 'dps',
  value: 110,
};

afterEach(() => vi.unstubAllGlobals());

describe('parse percentiles', () => {
  it('keys a query by everything that changes the answer', () => {
    expect(percentileKey(query)).toBe('9001|8|Protection|raids-1|dps|110');
  });

  it('asks the API once per distinct query and caches the answer', async () => {
    const upstream = vi.fn<GlobalFetch>(
      async () =>
        new Response(JSON.stringify({ ok: true, data: { percentile: 83.4 }, error: null, request_id: 'r' }), {
          status: 200,
          headers: { 'content-type': 'application/json' },
        }),
    );
    vi.stubGlobal('fetch', upstream);
    const loader = createPercentileLoader(API);

    await expect(loader.load([query, query])).resolves.toEqual({
      placements: new Map([[percentileKey(query), { percentile: 83.4, ranked: 0 }]]),
      unavailable: false,
    });
    await loader.load([query]);
    expect(upstream).toHaveBeenCalledTimes(1);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/rankings/percentile?encounter=9001&difficulty=8&spec=Protection&phase=raids-1&metric=dps&value=110`,
    );
  });

  it('runs at most six requests at a time', async () => {
    let running = 0;
    let peak = 0;
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        running += 1;
        peak = Math.max(peak, running);
        await new Promise((resolve) => setTimeout(resolve, 1));
        running -= 1;
        return new Response(
          JSON.stringify({ ok: true, data: { percentile: 1 }, error: null, request_id: 'r' }),
          {
            status: 200,
            headers: { 'content-type': 'application/json' },
          },
        );
      }),
    );
    const loader = createPercentileLoader(API);
    await loader.load(Array.from({ length: 20 }, (_, index) => ({ ...query, value: index })));
    expect(peak).toBeLessThanOrEqual(6);
  });

  it('leaves a row without a percentile rather than failing the table', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(createPercentileLoader(API).load([query])).resolves.toEqual({
      placements: new Map(),
      unavailable: true,
    });
  });

  it('says the rankings were unavailable on a rate limit, and asks again next time', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => new Response('slow down', { status: 429 }));
    vi.stubGlobal('fetch', upstream);
    const loader = createPercentileLoader(API);
    expect((await loader.load([query])).unavailable).toBe(true);
    expect((await loader.load([query])).unavailable).toBe(true);
    expect(upstream).toHaveBeenCalledTimes(2);
  });

  it('remembers a bracket nothing is ranked in, and does not ask again', async () => {
    const upstream = vi.fn<GlobalFetch>(
      async () =>
        new Response(
          JSON.stringify({ ok: false, data: null, error: { code: 'not_found' }, request_id: 'r' }),
          {
            status: 404,
            headers: { 'content-type': 'application/json' },
          },
        ),
    );
    vi.stubGlobal('fetch', upstream);
    const loader = createPercentileLoader(API);
    expect(await loader.load([query])).toEqual({ placements: new Map(), unavailable: false });
    expect(await loader.load([query])).toEqual({ placements: new Map(), unavailable: false });
    expect(upstream).toHaveBeenCalledTimes(1);
  });
});
