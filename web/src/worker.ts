// The only server code in web/. wrangler.jsonc sets run_worker_first to ["/b/*"], so this
// handler runs for shared build pages and their preview cards and nothing else; every other
// path is served straight from the static assets with no Worker execution at all.
//
// Shared builds are rendered by the Go API, but their links have to stay on foreversixty.gg:
// that is the domain people paste into Discord, and it is what the canonical tag and the
// Open Graph tags claim. So this proxies rather than redirects.
//
// The Workers runtime types are declared inline instead of pulling in
// @cloudflare/workers-types, which would have to be added to the Astro tsconfig's `types`
// and would then apply to every file in the project.
export interface Env {
  /** Origin of the Go API, from `vars` in wrangler.jsonc. */
  API_BASE_URL: string;
  /** The static assets binding. */
  ASSETS: { fetch(request: Request): Promise<Response> };
  /**
   * The foreversixty-logs bucket. Optional because `astro preview` and the Playwright
   * suite serve the static assets with no Worker and no bindings, and the checked-in
   * report fixture under public/logs-data/ has to keep working there.
   */
  LOGS?: R2Bucket;
}

/**
 * The R2 surface this Worker uses, declared inline for the same reason the ASSETS binding
 * is: @cloudflare/workers-types would have to go into the Astro tsconfig's `types` and
 * would then apply to every file in the project. `httpMetadata` is read field by field
 * rather than through the runtime's `writeHttpMetadata`, so the vitest fake stays a plain
 * object rather than a reimplementation of a Workers method.
 */
export interface R2HttpMetadata {
  contentType?: string;
  cacheControl?: string;
  contentEncoding?: string;
}

export interface R2ObjectBody {
  body: ReadableStream | null;
  httpMetadata?: R2HttpMetadata;
  httpEtag: string;
  size: number;
}

export interface R2Bucket {
  get(key: string): Promise<R2ObjectBody | null>;
}

const BUILD_PATH = /^\/b\/[^/]+(\/card\.png)?$/;

/** Only these reach the API; cookies never do, so the build page stays cacheable. */
const FORWARDED_HEADERS = ['accept', 'accept-language'];

function buildUpstreamHeaders(request: Request): Headers {
  const headers = new Headers();
  for (const name of FORWARDED_HEADERS) {
    const value = request.headers.get(name);
    if (value !== null) headers.set(name, value);
  }

  // Cloudflare sets CF-Connecting-IP at its edge and strips any copy a client sent, so it is
  // the one trustworthy source for the caller's IP. A client's own X-Forwarded-For is never
  // read, so it can never be smuggled through to the API.
  const clientIp = request.headers.get('cf-connecting-ip');
  if (clientIp !== null) headers.set('x-forwarded-for', clientIp);

  return headers;
}

async function unavailable(request: Request, env: Env): Promise<Response> {
  const url = new URL(request.url);
  const fallback = await env.ASSETS.fetch(new Request(new URL('/b-unavailable', url).toString()));
  const headers = new Headers(fallback.headers);
  headers.set('cache-control', 'no-store');
  return new Response(fallback.body, { status: 503, headers });
}

/**
 * Report files. `/logs-data/reports/<id>/<file>` maps one-to-one onto the engine's
 * store.Keys layout inside the bucket, so the path is the key with the prefix removed.
 *
 * The id pattern is the contract's: 12 characters of lowercase base32 (a-z plus 2-7). The
 * rest of the path is restricted to the four files the engine writes, which is also how
 * raw/ chunks are kept out -- they are `private` in the bucket and are downloaded from the
 * API by their owner, never from here.
 */
const LOGS_DATA_PATH =
  /^\/logs-data\/(reports\/([a-z2-7]{12})\/(?:report\.json|fights\/\d{1,6}\/(?:summary\.json|live\.json|events\.parquet)))$/;

/**
 * Everything under this prefix is either a report file matching LOGS_DATA_PATH or nothing
 * this route serves -- raw/ chunks (private, downloaded from the API by their owner) and
 * any other shape included. Falling through to ASSETS for those would be wrong: they are
 * not static site pages, so a 404 is refused directly rather than handed to the asset
 * pipeline as if it were an ordinary path.
 */
const LOGS_DATA_PREFIX = '/logs-data/';

/** Only these two are served without a session. Everything else goes through the API. */
const PUBLIC_VISIBILITIES = new Set(['public', 'unlisted']);

/**
 * Visibility, cached for sixty seconds per the contract. A module-level Map rather than
 * `caches.default`: it is per-isolate, which is exactly the granularity the contract asks
 * for, it survives across requests the way isolates do, and it is testable in plain vitest
 * with fake timers. A report that flips to private is therefore public for at most a
 * minute longer -- the contract's own bound.
 */
const VISIBILITY_TTL_MS = 60_000;
const visibilityCache = new Map<string, { visibility: string; expires: number }>();

type VisibilityLookup = { kind: 'ok'; visibility: string } | { kind: 'missing' } | { kind: 'unavailable' };

async function reportVisibility(id: string, env: Env): Promise<VisibilityLookup> {
  const cached = visibilityCache.get(id);
  if (cached !== undefined && cached.expires > Date.now()) {
    return { kind: 'ok', visibility: cached.visibility };
  }

  let response: Response;
  try {
    response = await fetch(`${env.API_BASE_URL}/v1/reports/${id}/visibility`);
  } catch {
    return { kind: 'unavailable' };
  }
  if (response.status === 404) return { kind: 'missing' };
  if (!response.ok) return { kind: 'unavailable' };

  let envelope: { ok: boolean; data: { visibility?: string } | null };
  try {
    envelope = (await response.json()) as { ok: boolean; data: { visibility?: string } | null };
  } catch {
    return { kind: 'unavailable' };
  }
  const visibility = envelope.data?.visibility;
  if (typeof visibility !== 'string' || visibility === '') return { kind: 'unavailable' };

  visibilityCache.set(id, { visibility, expires: Date.now() + VISIBILITY_TTL_MS });
  return { kind: 'ok', visibility };
}

function refuse(status: number, message: string, extra: Record<string, string> = {}): Response {
  return new Response(message, {
    status,
    headers: { 'content-type': 'text/plain; charset=utf-8', 'cache-control': 'no-store', ...extra },
  });
}

async function serveReportFile(request: Request, env: Env, key: string, id: string): Promise<Response> {
  if (request.method !== 'GET' && request.method !== 'HEAD') {
    return refuse(405, 'Method not allowed', { allow: 'GET, HEAD' });
  }

  const lookup = await reportVisibility(id, env);
  if (lookup.kind === 'missing') return refuse(404, 'No report with that id');
  if (lookup.kind === 'unavailable') return refuse(503, 'Report visibility cannot be checked right now');
  if (!PUBLIC_VISIBILITIES.has(lookup.visibility)) return refuse(403, 'That report is not public');

  // No bucket bound (preview, Playwright) or no such object: the static assets carry the
  // checked-in fixture at the same path, and otherwise answer with the 404 page.
  const object = await env.LOGS?.get(key);
  if (object === null || object === undefined) return env.ASSETS.fetch(request);

  const headers = new Headers();
  if (object.httpMetadata?.contentType) headers.set('content-type', object.httpMetadata.contentType);
  if (object.httpMetadata?.cacheControl) headers.set('cache-control', object.httpMetadata.cacheControl);
  if (object.httpMetadata?.contentEncoding) {
    headers.set('content-encoding', object.httpMetadata.contentEncoding);
  }
  headers.set('etag', object.httpEtag);
  return new Response(request.method === 'HEAD' ? null : object.body, { status: 200, headers });
}

export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const url = new URL(request.url);
    const logsData = LOGS_DATA_PATH.exec(url.pathname);
    if (logsData !== null) return serveReportFile(request, env, logsData[1], logsData[2]);
    if (url.pathname.startsWith(LOGS_DATA_PREFIX)) return refuse(404, 'Not a report file');

    if (!BUILD_PATH.test(url.pathname)) return env.ASSETS.fetch(request);

    let upstream: Response;
    try {
      upstream = await fetch(
        new Request(`${env.API_BASE_URL}${url.pathname}${url.search}`, {
          method: 'GET',
          headers: buildUpstreamHeaders(request),
          redirect: 'manual',
        }),
      );
    } catch {
      return unavailable(request, env);
    }

    // A 404 is the API saying the build does not exist, and it renders its own page for
    // that. Only a server-side failure means the API itself is unusable.
    if (upstream.status >= 500) return unavailable(request, env);

    // The rest of upstream's headers pass through unchanged -- ETag, Vary, Last-Modified,
    // Content-Encoding and the like all need to keep working as the API-rendered page grows,
    // so this is deliberately not an allow-list. Set-Cookie is the one exception: this
    // response can carry a cacheable Cache-Control and sit behind Cloudflare's shared edge
    // cache, so any Set-Cookie the API ever emitted here would be a cache-poisoning vector.
    // The Fetch spec special-cases Set-Cookie as the one response header that is never
    // combined: `Headers` keeps repeated Set-Cookie entries distinct internally (checked here
    // with Node's native Headers/Response -- multiple Set-Cookie values survive a `new
    // Headers(response.headers)` copy as separate entries, confirmed via `getSetCookie()`)
    // rather than folding them into one comma-joined value, and a single `delete` call
    // removes every one of them, so this is not a partial fix when the API sends more than
    // one Set-Cookie.
    const headers = new Headers(upstream.headers);
    headers.delete('set-cookie');

    return new Response(upstream.body, {
      status: upstream.status,
      statusText: upstream.statusText,
      headers,
    });
  },
};
