// web/src/lib/rating/report-ratings.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { createReportRatingsFetch } from './report-ratings.svelte';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

// A macrotask tick rather than a fixed number of microtask hops: unlike
// lazy-component.test.ts's mocked loader (which resolves in one hop),
// createReportRatingsFetch's load() runs through fetchReportRatings -> account/api.ts's
// requestEnvelope -> a real Response.json() -- a chain several microtasks deep in Node's
// fetch implementation. A `setTimeout` callback only runs once every microtask queued
// ahead of it (however many hops that took) has drained, so this flushes the whole chain
// regardless of its depth instead of guessing a hop count.
function flush(): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, 0));
}

afterEach(() => vi.unstubAllGlobals());

describe('createReportRatingsFetch', () => {
  it('loads, then reflects ready with the resolved data', async () => {
    const data = { fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] };
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(data)),
    );
    const hook = createReportRatingsFetch(API);

    expect(hook.status).toBe('idle');
    hook.load('fixture2abcd', 3);
    expect(hook.status).toBe('loading');
    await flush();

    expect(hook.status).toBe('ready');
    expect(hook.data).toEqual(data);
    expect(hook.error).toBe('');
  });

  it('a failure sets status failed and a message, and clears any stale data', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 500)),
    );
    const hook = createReportRatingsFetch(API);

    hook.load('fixture2abcd', 3);
    await flush();

    expect(hook.status).toBe('failed');
    expect(hook.data).toBeNull();
    expect(hook.error).not.toBe('');
  });

  it('switching fight while a request is in flight drops the stale response', async () => {
    let resolveFirst: ((r: Response) => void) | undefined;
    const upstream = vi.fn<GlobalFetch>(
      () =>
        new Promise<Response>((resolve) => {
          resolveFirst = resolve;
        }),
    );
    vi.stubGlobal('fetch', upstream);
    const hook = createReportRatingsFetch(API);

    hook.load('fixture2abcd', 1);
    const secondData = {
      fight_index: 2,
      kill: true,
      kill_time_band: 'typical',
      model_version: 'v1',
      players: [],
    };
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(secondData)),
    );
    hook.load('fixture2abcd', 2);
    await flush();

    expect(hook.data).toEqual(secondData);

    // The first request finally resolves; it must not clobber the second's already-landed data.
    resolveFirst?.(
      envelope({ fight_index: 1, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] }),
    );
    await flush();
    expect(hook.data).toEqual(secondData);
  });

  it('a repeat load() for the same fight while already ready is a no-op (no second fetch)', async () => {
    const data = { fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(data));
    vi.stubGlobal('fetch', upstream);
    const hook = createReportRatingsFetch(API);

    hook.load('fixture2abcd', 3);
    await flush();
    hook.load('fixture2abcd', 3);

    expect(upstream).toHaveBeenCalledTimes(1);
  });
});
