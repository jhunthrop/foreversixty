import { afterEach, describe, expect, it, vi } from 'vitest';
import worker, { type Env } from './worker';
import { DUCKDB_WASM_VERSION, duckdbRuntimeUrl } from './lib/report/duckdb-runtime';
import { fixtureResult } from './test-support/sim-api';

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

  it('strips every Set-Cookie from the proxied response without touching Cache-Control', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => {
      const headers = new Headers({ 'cache-control': 'public, max-age=60' });
      headers.append('set-cookie', 'a=1; Path=/');
      headers.append('set-cookie', 'b=2; Path=/');
      return new Response('<html>build</html>', { status: 200, headers });
    });
    vi.stubGlobal('fetch', upstream);
    const env = envWith();

    const response = await worker.fetch(new Request('https://foreversixty.gg/b/k7x2qm4a'), env);

    expect(response.headers.get('set-cookie')).toBeNull();
    expect(response.headers.getSetCookie()).toEqual([]);
    expect(response.headers.get('cache-control')).toBe('public, max-age=60');
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

// --- appended to web/src/worker.test.ts ---

interface FakeObject {
  body: string;
  httpMetadata: { contentType?: string; cacheControl?: string; contentEncoding?: string };
  httpEtag: string;
}

function bucketWith(objects: Record<string, FakeObject>) {
  return {
    get: vi.fn(async (key: string) => {
      const found = objects[key];
      if (found === undefined) return null;
      return {
        body: new Response(found.body).body,
        httpMetadata: found.httpMetadata,
        httpEtag: found.httpEtag,
        size: found.body.length,
      };
    }),
  };
}

const SUMMARY_KEY = (id: string) => `reports/${id}/fights/3/summary.json`;

function summaryObject(): FakeObject {
  return {
    body: '{"fight_index":3}',
    httpMetadata: {
      contentType: 'application/json',
      cacheControl: 'public, max-age=31536000, immutable',
    },
    httpEtag: '"abc123"',
  };
}

/** The API's answer to GET /v1/reports/{id}/visibility, in the Phase 0 envelope. */
function visibilityResponse(visibility: string, status = 200): Response {
  return new Response(
    JSON.stringify({ ok: status < 400, data: { visibility }, error: null, request_id: 'req-1' }),
    { status, headers: { 'content-type': 'application/json' } },
  );
}

describe('/logs-data/* served from the R2 bucket', () => {
  it('serves a public report’s summary with the headers R2 stored', async () => {
    const id = 'pubaaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('public')),
    );
    const env = { ...envWith(), LOGS: bucketWith({ [SUMMARY_KEY(id)]: summaryObject() }) };

    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`),
      env,
    );

    expect(response.status).toBe(200);
    expect(await response.text()).toBe('{"fight_index":3}');
    expect(response.headers.get('content-type')).toBe('application/json');
    expect(response.headers.get('cache-control')).toBe('public, max-age=31536000, immutable');
    expect(response.headers.get('etag')).toBe('"abc123"');
    expect(env.LOGS.get).toHaveBeenCalledWith(SUMMARY_KEY(id));
  });

  it('serves an unlisted report too', async () => {
    const id = 'unlaaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('unlisted')),
    );
    const env = { ...envWith(), LOGS: bucketWith({ [SUMMARY_KEY(id)]: summaryObject() }) };
    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`),
      env,
    );
    expect(response.status).toBe(200);
  });

  it('refuses a private report and never touches the bucket', async () => {
    const id = 'privaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('private')),
    );
    const env = { ...envWith(), LOGS: bucketWith({ [SUMMARY_KEY(id)]: summaryObject() }) };

    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`),
      env,
    );

    expect(response.status).toBe(403);
    expect(response.headers.get('cache-control')).toBe('no-store');
    expect(env.LOGS.get).not.toHaveBeenCalled();
  });

  it('refuses a guild report the same way', async () => {
    const id = 'gldaaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('guild')),
    );
    const env = { ...envWith(), LOGS: bucketWith({}) };
    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/report.json`),
      env,
    );
    expect(response.status).toBe(403);
  });

  it('caches the visibility answer for sixty seconds, then asks again', async () => {
    vi.useFakeTimers();
    const id = 'cachaaaaaaaa';
    const upstream = vi.fn<GlobalFetch>(async () => visibilityResponse('public'));
    vi.stubGlobal('fetch', upstream);
    const env = { ...envWith(), LOGS: bucketWith({ [SUMMARY_KEY(id)]: summaryObject() }) };
    const url = `https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`;

    await worker.fetch(new Request(url), env);
    await worker.fetch(new Request(url), env);
    expect(upstream).toHaveBeenCalledTimes(1);

    vi.setSystemTime(Date.now() + 61_000);
    await worker.fetch(new Request(url), env);
    expect(upstream).toHaveBeenCalledTimes(2);
    vi.useRealTimers();
  });

  it('answers 503 when the API cannot say, and shields it from a stampede for five seconds', async () => {
    vi.useFakeTimers();
    const id = 'downaaaaaaaa';
    const upstream = vi.fn<GlobalFetch>(async () => {
      throw new TypeError('offline');
    });
    vi.stubGlobal('fetch', upstream);
    const env = { ...envWith(), LOGS: bucketWith({ [SUMMARY_KEY(id)]: summaryObject() }) };
    const url = `https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`;

    const response = await worker.fetch(new Request(url), env);

    expect(response.status).toBe(503);
    expect(response.headers.get('cache-control')).toBe('no-store');
    expect(env.LOGS.get).not.toHaveBeenCalled();

    // An outage must not cost the recovering API one subrequest per inbound request. The
    // response itself still says no-store: it is the Worker's own lookup that is held, not
    // the visitor's copy of the failure.
    await worker.fetch(new Request(url), env);
    expect(upstream).toHaveBeenCalledTimes(1);

    // And it is only five seconds, so the site follows the API back up straight away.
    vi.setSystemTime(Date.now() + 6_000);
    await worker.fetch(new Request(url), env);
    expect(upstream).toHaveBeenCalledTimes(2);
    vi.useRealTimers();
  });

  it('answers 404 for a report the API does not know, and remembers that for a minute', async () => {
    vi.useFakeTimers();
    const id = 'gonaaaaaaaaa';
    const upstream = vi.fn<GlobalFetch>(async () => visibilityResponse('', 404));
    vi.stubGlobal('fetch', upstream);
    const env = { ...envWith(), LOGS: bucketWith({}) };
    const url = `https://foreversixty.gg/logs-data/reports/${id}/report.json`;

    const response = await worker.fetch(new Request(url), env);
    expect(response.status).toBe(404);

    // A crawler working through guessed ids, or one dead link being retried, costs one
    // subrequest a minute rather than one per request: an id nothing answers for is as
    // stable an answer as a real visibility.
    await worker.fetch(new Request(url), env);
    await worker.fetch(new Request(url), env);
    expect(upstream).toHaveBeenCalledTimes(1);

    vi.setSystemTime(Date.now() + 61_000);
    await worker.fetch(new Request(url), env);
    expect(upstream).toHaveBeenCalledTimes(2);
    vi.useRealTimers();
  });

  it('falls through to the static assets when the object is missing, which is how the fixture serves in preview', async () => {
    const id = 'fixaaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('public')),
    );
    const env = { ...envWith('fixture bytes'), LOGS: bucketWith({}) };

    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/report.json`),
      env,
    );

    expect(await response.text()).toBe('fixture bytes');
    expect(env.ASSETS.fetch).toHaveBeenCalledTimes(1);
  });

  it('falls through to the static assets when there is no bucket bound at all', async () => {
    const id = 'nobaaaaaaaaa';
    const upstream = vi.fn<GlobalFetch>(async () => visibilityResponse('public'));
    vi.stubGlobal('fetch', upstream);
    const env = envWith('fixture bytes');

    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/report.json`),
      env,
    );

    expect(await response.text()).toBe('fixture bytes');
  });

  it('rejects a path that is not a report file, without calling the API', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    const env = { ...envWith(), LOGS: bucketWith({}) };

    // A literal '../../secret' is not tested here: the Request constructor resolves
    // dot-segments before the Worker ever sees the URL, so a plain `..` never reaches this
    // code as anything but an ordinary (non-/logs-data/) path. The percent-encoded form
    // below stays literal through URL parsing, so it's the one that actually exercises the
    // handler's rejection.
    for (const path of [
      '/logs-data/reports/%2e%2e%2f%2e%2e%2fsecret',
      '/logs-data/reports/..%2f..%2fsecret',
      '/logs-data/reports/SHOUTING1234/report.json',
      '/logs-data/reports/short/report.json',
      '/logs-data/uploads/abc/raw.txt.zst',
      '/logs-data/',
    ]) {
      const response = await worker.fetch(new Request(`https://foreversixty.gg${path}`), env);
      expect(response.status, path).toBe(404);
    }
    expect(upstream).not.toHaveBeenCalled();
    expect(env.LOGS.get).not.toHaveBeenCalled();
  });

  it('refuses raw chunks, which are private to their owner and downloaded from the API', async () => {
    const id = 'rawaaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('public')),
    );
    const env = { ...envWith(), LOGS: bucketWith({ [`reports/${id}/raw/0.zst`]: summaryObject() }) };

    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/raw/0.zst`),
      env,
    );

    expect(response.status).toBe(404);
    expect(env.LOGS.get).not.toHaveBeenCalled();
  });

  it('answers 405 for anything but GET and HEAD', async () => {
    const id = 'putaaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('public')),
    );
    const env = { ...envWith(), LOGS: bucketWith({}) };
    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/report.json`, { method: 'PUT' }),
      env,
    );
    expect(response.status).toBe(405);
    expect(response.headers.get('allow')).toBe('GET, HEAD');
  });

  it('answers a matching If-None-Match with a 304 and no body', async () => {
    const id = 'pubaaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('public')),
    );
    const env = { ...envWith(), LOGS: bucketWith({ [SUMMARY_KEY(id)]: summaryObject() }) };
    const response = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`, {
        headers: { 'if-none-match': '"abc123"' },
      }),
      env,
    );
    expect(response.status).toBe(304);
    expect(await response.text()).toBe('');
    expect(response.headers.get('etag')).toBe('"abc123"');
    expect(response.headers.get('cache-control')).toBe('public, max-age=31536000, immutable');

    const weak = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`, {
        headers: { 'if-none-match': 'W/"abc123", "other"' },
      }),
      env,
    );
    expect(weak.status).toBe(304);

    const stale = await worker.fetch(
      new Request(`https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`, {
        headers: { 'if-none-match': '"older"' },
      }),
      env,
    );
    expect(stale.status).toBe(200);
    expect(await stale.text()).toBe('{"fight_index":3}');
  });

  it('serves a repeat request from the edge cache without reading the bucket', async () => {
    const id = 'pubaaaaaaaaa';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => visibilityResponse('public')),
    );
    const store = new Map<string, Response>();
    vi.stubGlobal('caches', {
      default: {
        match: async (url: string) => store.get(url)?.clone(),
        put: async (url: string, response: Response) => {
          store.set(url, response);
        },
      },
    });
    const env = { ...envWith(), LOGS: bucketWith({ [SUMMARY_KEY(id)]: summaryObject() }) };
    const get = env.LOGS.get;
    const url = `https://foreversixty.gg/logs-data/reports/${id}/fights/3/summary.json`;
    const first = await worker.fetch(new Request(url), env);
    expect(first.status).toBe(200);
    expect(await first.text()).toBe('{"fight_index":3}');
    expect(get).toHaveBeenCalledTimes(1);

    const second = await worker.fetch(new Request(url), env);
    expect(second.status).toBe(200);
    expect(await second.text()).toBe('{"fight_index":3}');
    expect(second.headers.get('etag')).toBe('"abc123"');
    expect(get).toHaveBeenCalledTimes(1);

    const revalidated = await worker.fetch(
      new Request(url, { headers: { 'if-none-match': '"abc123"' } }),
      env,
    );
    expect(revalidated.status).toBe(304);
    expect(get).toHaveBeenCalledTimes(1);

    // A different engine version on the query is a different cache entry and reaches the
    // bucket again: that is how a re-parse's new bytes get past a year-long cache.
    const reparsed = await worker.fetch(new Request(`${url}?v=0.3.5`), env);
    expect(reparsed.status).toBe(200);
    expect(get).toHaveBeenCalledTimes(2);
    expect(get).toHaveBeenLastCalledWith(SUMMARY_KEY(id));
  });
});

describe('/duckdb-runtime/* served from the R2 bucket', () => {
  const MODULE_KEY = `runtime/duckdb/${DUCKDB_WASM_VERSION}/duckdb-eh.wasm`;
  const MODULE_URL = `https://foreversixty.gg${duckdbRuntimeUrl('duckdb-eh.wasm')}`;

  function wasmObject(): FakeObject {
    return { body: '\u0000asm', httpMetadata: {}, httpEtag: '"wasm1"' };
  }

  // These two files are 32.7 and 37.5 MiB, over Cloudflare's 25 MiB static-asset cap, so
  // they cannot be in dist/ at all: a deploy that carries them fails for the whole site.
  it('serves the engine module with an immutable year and the WebAssembly type', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => visibilityResponse('public'));
    vi.stubGlobal('fetch', upstream);
    const env = { ...envWith(), LOGS: bucketWith({ [MODULE_KEY]: wasmObject() }) };

    const response = await worker.fetch(new Request(MODULE_URL), env);

    expect(response.status).toBe(200);
    expect(response.headers.get('content-type')).toBe('application/wasm');
    expect(response.headers.get('cache-control')).toBe('public, max-age=31536000, immutable');
    expect(response.headers.get('etag')).toBe('"wasm1"');
    expect(env.LOGS.get).toHaveBeenCalledWith(MODULE_KEY);
    // Public bytes out of an npm package, not report data: no visibility lookup, and so no
    // API subrequest for a file every visitor gets the same copy of.
    expect(upstream).not.toHaveBeenCalled();
  });

  it('answers a HEAD without a body', async () => {
    const env = { ...envWith(), LOGS: bucketWith({ [MODULE_KEY]: wasmObject() }) };
    const response = await worker.fetch(new Request(MODULE_URL, { method: 'HEAD' }), env);
    expect(response.status).toBe(200);
    expect(await response.text()).toBe('');
  });

  it('answers 405 for anything but GET and HEAD', async () => {
    const env = { ...envWith(), LOGS: bucketWith({ [MODULE_KEY]: wasmObject() }) };
    const response = await worker.fetch(new Request(MODULE_URL, { method: 'PUT' }), env);
    expect(response.status).toBe(405);
    expect(response.headers.get('allow')).toBe('GET, HEAD');
  });

  it('404s anything else under the prefix without touching the bucket or the assets', async () => {
    const env = { ...envWith(), LOGS: bucketWith({ [MODULE_KEY]: wasmObject() }) };
    const refused = [
      `https://foreversixty.gg/duckdb-runtime/${DUCKDB_WASM_VERSION}/duckdb-coi.wasm`,
      'https://foreversixty.gg/duckdb-runtime/1.31.0/duckdb-eh.wasm',
      `https://foreversixty.gg/duckdb-runtime/${DUCKDB_WASM_VERSION}/`,
    ];
    for (const url of refused) {
      const response = await worker.fetch(new Request(url), env);
      expect(response.status).toBe(404);
    }
    expect(env.LOGS.get).not.toHaveBeenCalled();
    expect(env.ASSETS.fetch).not.toHaveBeenCalled();
  });

  it('404s rather than falling back to the assets when the upload has not been done', async () => {
    const env = { ...envWith('a 404 page'), LOGS: bucketWith({}) };
    const response = await worker.fetch(new Request(MODULE_URL), env);
    expect(response.status).toBe(404);
    expect(env.ASSETS.fetch).not.toHaveBeenCalled();
  });
});

// --- appended to web/src/worker.test.ts ---
import { FakeHTMLRewriter } from './test-support/html-rewriter';
import fixtureMeta from './fixtures/report/meta.json';

const SHELL_HTML = `<!doctype html><html><head><title>Report · Forever Sixty</title>
<meta name="description" content="placeholder" data-og="description" />
<link rel="canonical" href="https://foreversixty.gg/reports" data-og="canonical" />
<meta property="og:title" content="Report · Forever Sixty" data-og="og-title" />
<meta property="og:description" content="placeholder" data-og="og-description" />
<meta property="og:url" content="https://foreversixty.gg/reports" data-og="og-url" />
<meta property="og:image" content="https://foreversixty.gg/og/reports.png" data-og="og-image" />
</head><body><div id="report" data-report-mount></div></body></html>`;

function shellEnv() {
  return {
    API_BASE_URL,
    ASSETS: {
      fetch: vi.fn<AssetFetch>(async (request: Request) =>
        new URL(request.url).pathname === '/reports.html'
          ? new Response(SHELL_HTML, { status: 200, headers: { 'content-type': 'text/html' } })
          : new Response('not found', { status: 404 }),
      ),
    },
  };
}

describe('shell routes with rewritten unfurl tags', () => {
  it('serves dist/reports.html for any id, rewritten from the API’s JSON', async () => {
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(
        async () =>
          new Response(JSON.stringify({ ok: true, data: fixtureMeta, error: null, request_id: 'r' }), {
            status: 200,
            headers: { 'content-type': 'application/json' },
          }),
      ),
    );
    const env = shellEnv();

    const response = await worker.fetch(new Request('https://foreversixty.gg/reports/fixture2abcd'), env);
    const html = await response.text();

    expect(response.status).toBe(200);
    expect(response.headers.get('content-type')).toContain('text/html');
    expect(response.headers.get('cache-control')).toBe('public, max-age=60');
    expect(html).toContain('<title>Sanguine Depths, fixture night · Forever Sixty</title>');
    expect(html).toContain('content="4 fights in Sanguine Depths, 2 boss kills, logged 2026-09-26."');
    expect(html).toContain(`content="${API_BASE_URL}/reports/fixture2abcd/card.png"`);
    expect(html).toContain('href="https://foreversixty.gg/reports/fixture2abcd"');
    expect(html).toContain('content="https://foreversixty.gg/reports/fixture2abcd"');
    expect((env.ASSETS.fetch.mock.calls[0][0] as Request).url).toBe('https://foreversixty.gg/reports.html');
  });

  it('never sends the visitor’s cookies to the API for a cacheable shell', async () => {
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    const upstream = vi.fn<GlobalFetch>(
      async () =>
        new Response(JSON.stringify({ ok: true, data: fixtureMeta, error: null, request_id: 'r' }), {
          status: 200,
          headers: { 'content-type': 'application/json' },
        }),
    );
    vi.stubGlobal('fetch', upstream);

    await worker.fetch(
      new Request('https://foreversixty.gg/reports/fixture2abcd', {
        headers: { cookie: 'fs_session=secret' },
      }),
      shellEnv(),
    );

    // Two calls happen for a report shell: the visibility check (a bare url string, which
    // carries no headers by construction) and the report data fetch (a Request built fresh
    // by apiData). Every call this route makes is checked, not just the first.
    for (const [arg] of upstream.mock.calls) {
      if (arg instanceof Request) expect(arg.headers.get('cookie')).toBeNull();
    }
    expect(upstream.mock.calls.some(([arg]) => arg instanceof Request)).toBe(true);
  });

  it('gates the report branch on the visibility endpoint, never the report data itself, and never leaks a private report’s title', async () => {
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    const upstream = vi.fn<GlobalFetch>(async (input) => {
      const url = typeof input === 'string' ? input : (input as Request).url;
      if (url.endsWith('/visibility')) {
        return new Response(
          JSON.stringify({ ok: true, data: { visibility: 'private' }, error: null, request_id: 'r' }),
          { status: 200, headers: { 'content-type': 'application/json' } },
        );
      }
      // If the report branch ever reaches this, the gate has failed: fail loudly rather
      // than answering with data the test would otherwise appear to pass against.
      throw new Error(`unexpected fetch to ${url}`);
    });
    vi.stubGlobal('fetch', upstream);

    const response = await worker.fetch(
      new Request('https://foreversixty.gg/reports/fixture2abcd'),
      shellEnv(),
    );
    const html = await response.text();

    expect(response.status).toBe(200);
    expect(html).toContain('<title>Report · Forever Sixty</title>');
    expect(html).not.toContain('Sanguine Depths');
    expect(response.headers.get('x-robots-tag')).toBe('noindex');
    expect(upstream).toHaveBeenCalledTimes(1);
  });

  it('gates a guild report the same way as a private one', async () => {
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    const upstream = vi.fn<GlobalFetch>(async (input) => {
      const url = typeof input === 'string' ? input : (input as Request).url;
      if (url.endsWith('/visibility')) {
        return new Response(
          JSON.stringify({ ok: true, data: { visibility: 'guild' }, error: null, request_id: 'r' }),
          { status: 200, headers: { 'content-type': 'application/json' } },
        );
      }
      throw new Error(`unexpected fetch to ${url}`);
    });
    vi.stubGlobal('fetch', upstream);

    const response = await worker.fetch(
      new Request('https://foreversixty.gg/reports/fixture2abcd'),
      shellEnv(),
    );
    const html = await response.text();

    expect(html).toContain('<title>Report · Forever Sixty</title>');
    expect(response.headers.get('x-robots-tag')).toBe('noindex');
    expect(upstream).toHaveBeenCalledTimes(1);
  });

  it('marks a non-public report noindex', async () => {
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(
        async () =>
          new Response(
            JSON.stringify({
              ok: true,
              data: { ...fixtureMeta, visibility: 'unlisted' },
              error: null,
              request_id: 'r',
            }),
            { status: 200, headers: { 'content-type': 'application/json' } },
          ),
      ),
    );
    const response = await worker.fetch(
      new Request('https://foreversixty.gg/reports/fixture2abcd'),
      shellEnv(),
    );
    expect(response.headers.get('x-robots-tag')).toBe('noindex');
  });

  it('serves the shell unrewritten when the API cannot answer, rather than failing the page', async () => {
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );

    const response = await worker.fetch(
      new Request('https://foreversixty.gg/reports/fixture2abcd'),
      shellEnv(),
    );
    const html = await response.text();

    expect(response.status).toBe(200);
    expect(html).toContain('<title>Report · Forever Sixty</title>');
    expect(html).toContain('data-report-mount');
    // An unrewritten placeholder is still noindex: a crawler indexing it would confirm the
    // id exists even though it carries no report data.
    expect(response.headers.get('x-robots-tag')).toBe('noindex');
  });

  it('passes a shell prefix whose asset does not exist yet straight through', async () => {
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>());
    const env = shellEnv();
    const response = await worker.fetch(new Request('https://foreversixty.gg/rankings/warden-kelthas'), env);
    expect(response.status).toBe(404);
    expect((env.ASSETS.fetch.mock.calls[0][0] as Request).url).toBe('https://foreversixty.gg/rankings.html');
  });

  it('serves the shell unrewritten where HTMLRewriter does not exist', async () => {
    vi.stubGlobal('HTMLRewriter', undefined);
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(
        async () =>
          new Response(JSON.stringify({ ok: true, data: fixtureMeta, error: null, request_id: 'r' }), {
            status: 200,
            headers: { 'content-type': 'application/json' },
          }),
      ),
    );
    const html = await (
      await worker.fetch(new Request('https://foreversixty.gg/reports/fixture2abcd'), shellEnv())
    ).text();
    expect(html).toContain('<title>Report · Forever Sixty</title>');
  });

  it('leaves a path under a shell prefix that is not a valid id to the assets', async () => {
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    const env = shellEnv();
    const response = await worker.fetch(new Request('https://foreversixty.gg/reports/NOT-AN-ID/extra'), env);
    expect(upstream).not.toHaveBeenCalled();
    expect((env.ASSETS.fetch.mock.calls[0][0] as Request).url).toBe(
      'https://foreversixty.gg/reports/NOT-AN-ID/extra',
    );
    expect(response.status).toBe(404);
  });
});

// --- appended for Task 9: the sim shell route ---

/** The Phase 0 envelope, the same shape every route in this repository answers in. */
function json(body: unknown): Response {
  return new Response(JSON.stringify(body), { status: 200, headers: { 'content-type': 'application/json' } });
}

/** An ASSETS binding that answers `matchPath` with `html` and 404s everything else. */
function shellAssets(matchPath: string, html = SHELL_HTML) {
  return {
    fetch: vi.fn<AssetFetch>(async (request: Request) =>
      new URL(request.url).pathname === matchPath
        ? new Response(html, { status: 200, headers: { 'content-type': 'text/html' } })
        : new Response('not found', { status: 404 }),
    ),
  };
}

describe('/sim/<sim_id>', () => {
  it('serves sim.html with the head rewritten from the API', async () => {
    const assets = shellAssets('/sim.html');
    const upstream = vi.fn(async (request: Request) => {
      expect(new URL(request.url).pathname).toBe('/v1/sims/simfixtureab');
      return json({ ok: true, data: fixtureResult, error: null, request_id: 'r' });
    });
    vi.stubGlobal('fetch', upstream);
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);

    const response = await worker.fetch(new Request('https://foreversixty.gg/sim/simfixtureab'), {
      API_BASE_URL: 'https://api.test',
      ASSETS: assets,
    } as unknown as Env);

    const html = await response.text();
    expect(response.status).toBe(200);
    expect(response.headers.get('cache-control')).toBe('public, max-age=60');
    expect(response.headers.get('x-robots-tag')).toBeNull();
    expect(html).toContain('Fury Warrior, 1,039 DPS · Forever Sixty');
    expect(html).toContain('https://foreversixty.gg/sim/simfixtureab');
  });

  it('serves the shell unrewritten and noindex when the API cannot answer', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => new Response('nope', { status: 500 })),
    );
    vi.stubGlobal('HTMLRewriter', FakeHTMLRewriter);
    const response = await worker.fetch(new Request('https://foreversixty.gg/sim/simfixtureab'), {
      API_BASE_URL: 'https://api.test',
      ASSETS: shellAssets('/sim.html'),
    } as unknown as Env);
    expect(response.status).toBe(200);
    expect(response.headers.get('x-robots-tag')).toBe('noindex');
  });

  it('leaves /sim and /sim/specs to the static assets, with no API call at all', async () => {
    const upstream = vi.fn();
    vi.stubGlobal('fetch', upstream);
    const assets = { fetch: vi.fn(async () => new Response('static', { status: 200 })) };
    for (const path of ['/sim', '/sim/specs', '/sim/not-an-id']) {
      await worker.fetch(new Request(`https://foreversixty.gg${path}`), {
        API_BASE_URL: 'https://api.test',
        ASSETS: assets,
      } as unknown as Env);
    }
    expect(upstream).not.toHaveBeenCalled();
    expect(assets.fetch).toHaveBeenCalledTimes(3);
  });
});
