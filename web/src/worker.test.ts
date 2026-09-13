import { afterEach, describe, expect, it, vi } from 'vitest';
import worker from './worker';

const API_BASE_URL = 'https://api.foreversixty.test';

// Typing the mocks with the real signatures is what makes `mock.calls[0][0]` type-check
// under `astro check`, which reads every file in the project including this one.
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;
type AssetFetch = (request: Request) => Promise<Response>;

function envWith(assetBody = 'asset') {
  return {
    API_BASE_URL,
    ASSETS: { fetch: vi.fn<AssetFetch>(async () => new Response(assetBody, { status: 200 })) },
  };
}

afterEach(() => vi.unstubAllGlobals());

describe('worker fetch handler', () => {
  it('proxies /b/:id to the API, keeping the path, query and cache headers', async () => {
    const upstream = vi.fn<GlobalFetch>(
      async () =>
        new Response('<html>build</html>', {
          status: 200,
          headers: { 'content-type': 'text/html', 'cache-control': 'public, max-age=60' },
        }),
    );
    vi.stubGlobal('fetch', upstream);
    const env = envWith();

    const response = await worker.fetch(
      new Request('https://foreversixty.gg/b/k7x2qm4a?from=discord', {
        headers: {
          accept: 'text/html',
          'accept-language': 'en-GB',
          cookie: 'session=secret',
          'cf-connecting-ip': '203.0.113.5',
        },
      }),
      env,
    );

    expect(response.status).toBe(200);
    expect(response.headers.get('cache-control')).toBe('public, max-age=60');
    expect(env.ASSETS.fetch).not.toHaveBeenCalled();

    const proxied = upstream.mock.calls[0][0] as Request;
    expect(proxied.url).toBe(`${API_BASE_URL}/b/k7x2qm4a?from=discord`);
    expect(proxied.headers.get('accept')).toBe('text/html');
    expect(proxied.headers.get('accept-language')).toBe('en-GB');
    expect(proxied.headers.get('cookie')).toBeNull();
    expect(proxied.headers.get('x-forwarded-for')).toBe('203.0.113.5');
  });

  it('sets X-Forwarded-For from CF-Connecting-IP and never forwards a client-supplied value', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => new Response('ok', { status: 200 }));
    vi.stubGlobal('fetch', upstream);
    const env = envWith();

    await worker.fetch(
      new Request('https://foreversixty.gg/b/k7x2qm4a', {
        headers: { 'cf-connecting-ip': '203.0.113.5', 'x-forwarded-for': '10.0.0.1, evil' },
      }),
      env,
    );

    const proxied = upstream.mock.calls[0][0] as Request;
    expect(proxied.headers.get('x-forwarded-for')).toBe('203.0.113.5');
  });

  it('proxies the card image too', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => new Response('png', { status: 200 }));
    vi.stubGlobal('fetch', upstream);
    await worker.fetch(new Request('https://foreversixty.gg/b/k7x2qm4a/card.png'), envWith());
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API_BASE_URL}/b/k7x2qm4a/card.png`);
  });

  it('serves every other path from the static assets and never calls the API', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    const env = envWith('about page');
    const response = await worker.fetch(new Request('https://foreversixty.gg/about'), env);
    expect(await response.text()).toBe('about page');
    expect(env.ASSETS.fetch).toHaveBeenCalledTimes(1);
    expect(upstream).not.toHaveBeenCalled();
  });

  it('does not treat /builds as a build page', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    const env = envWith();
    await worker.fetch(new Request('https://foreversixty.gg/builds'), env);
    expect(env.ASSETS.fetch).toHaveBeenCalledTimes(1);
    expect(upstream).not.toHaveBeenCalled();
  });

  it('falls back to the static page with 503 and no caching when the API is down', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('connect ECONNREFUSED');
      }),
    );
    const env = envWith('planner is over here');
    const response = await worker.fetch(new Request('https://foreversixty.gg/b/k7x2qm4a'), env);
    expect(response.status).toBe(503);
    expect(response.headers.get('cache-control')).toBe('no-store');
    expect(await response.text()).toBe('planner is over here');
    expect((env.ASSETS.fetch.mock.calls[0][0] as Request).url).toBe('https://foreversixty.gg/b-unavailable');
  });

  it('passes an API 404 straight through rather than masking it', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => new Response('no such build', { status: 404 })),
    );
    const env = envWith();
    const response = await worker.fetch(new Request('https://foreversixty.gg/b/nope'), env);
    expect(response.status).toBe(404);
    expect(env.ASSETS.fetch).not.toHaveBeenCalled();
  });

  it('falls back when the API answers 5xx', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => new Response('boom', { status: 502 })),
    );
    const env = envWith('planner is over here');
    const response = await worker.fetch(new Request('https://foreversixty.gg/b/k7x2qm4a'), env);
    expect(response.status).toBe(503);
  });
});
