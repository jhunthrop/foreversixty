// The only server code in web/. wrangler.jsonc's run_worker_first lists exactly the
// prefixes this handler owns -- the shared build proxy, report files and the DuckDB
// runtime out of R2, and the four shells -- and every other path is served straight from
// the static assets with no Worker execution at all.
//
// Shared builds are rendered by the Go API, but their links have to stay on foreversixty.gg:
// that is the domain people paste into Discord, and it is what the canonical tag and the
// Open Graph tags claim. So this proxies rather than redirects.
//
// The Workers runtime types are declared inline instead of pulling in
// @cloudflare/workers-types, which would have to be added to the Astro tsconfig's `types`
// and would then apply to every file in the project.
import { parseCharacterPath, parseGuildPath } from './lib/characters';
import { DUCKDB_RUNTIME_PREFIX, duckdbRuntimeKey } from './lib/report/duckdb-runtime';
import {
  characterShellMeta,
  guildShellMeta,
  rankingsShellMeta,
  reportShellMeta,
  type CharacterHead,
  type GuildHead,
  type RankingsHead,
  type ShellMeta,
} from './lib/report/og-meta';
import type { ReportMeta } from './lib/report/types';

export interface Env {
  /** Origin of the Go API, from `vars` in wrangler.jsonc. */
  API_BASE_URL: string;
  /** The static assets binding. */
  ASSETS: { fetch(request: Request): Promise<Response> };
  /**
   * The foreversixty-logs bucket. It holds both a report's files and the two DuckDB
   * engine modules. Optional because `astro preview` and the Playwright suite serve the
   * static assets with no Worker and no bindings, and the checked-in report fixture under
   * public/logs-data/ has to keep working there.
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

/**
 * A "no such report" is as stable an answer as a real one and is cached for as long:
 * without this, every request for an id nothing answers for -- a crawler working through
 * guessed ids, a stale link being retried -- costs an uncached subrequest to the API.
 */
const MISSING_TTL_MS = VISIBILITY_TTL_MS;

/**
 * An outage is not stable, so it is cached only long enough to stop a recovering API being
 * hit once per inbound request while it comes back. Short enough that the site follows it
 * back up within one visitor's reload.
 */
const UNAVAILABLE_TTL_MS = 5_000;

type VisibilityLookup = { kind: 'ok'; visibility: string } | { kind: 'missing' } | { kind: 'unavailable' };

const visibilityCache = new Map<string, { lookup: VisibilityLookup; expires: number }>();

const VISIBILITY_TTL_BY_KIND: Record<VisibilityLookup['kind'], number> = {
  ok: VISIBILITY_TTL_MS,
  missing: MISSING_TTL_MS,
  unavailable: UNAVAILABLE_TTL_MS,
};

function remember(id: string, lookup: VisibilityLookup): VisibilityLookup {
  visibilityCache.set(id, { lookup, expires: Date.now() + VISIBILITY_TTL_BY_KIND[lookup.kind] });
  return lookup;
}

async function reportVisibility(id: string, env: Env): Promise<VisibilityLookup> {
  const cached = visibilityCache.get(id);
  if (cached !== undefined && cached.expires > Date.now()) return cached.lookup;

  let response: Response;
  try {
    response = await fetch(`${env.API_BASE_URL}/v1/reports/${id}/visibility`);
  } catch {
    return remember(id, { kind: 'unavailable' });
  }
  if (response.status === 404) return remember(id, { kind: 'missing' });
  if (!response.ok) return remember(id, { kind: 'unavailable' });

  let envelope: { ok: boolean; data: { visibility?: string } | null };
  try {
    envelope = (await response.json()) as { ok: boolean; data: { visibility?: string } | null };
  } catch {
    return remember(id, { kind: 'unavailable' });
  }
  const visibility = envelope.data?.visibility;
  if (typeof visibility !== 'string' || visibility === '') {
    return remember(id, { kind: 'unavailable' });
  }

  return remember(id, { kind: 'ok', visibility });
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

/**
 * The two DuckDB engine modules, out of the same bucket a report's files come from.
 *
 * They are here rather than in dist/ because Cloudflare caps a static asset at 25 MiB and
 * these are 32.7 and 37.5 MiB (scripts/duckdb-runtime.mjs). Unlike /logs-data/*, there is
 * no visibility check: these are the vendored public bytes of an npm package, identical
 * for every visitor, and the only reason they are served from here at all is the site's
 * rule that no page makes a third-party request.
 *
 * The path carries the package version and duckdbRuntimeKey() matches it exactly, so the
 * response can be cached forever and anything else under the prefix -- an old version, a
 * filename this site does not publish, a traversal attempt -- is a 404 rather than a
 * lookup.
 */
async function serveDuckdbRuntime(request: Request, env: Env, pathname: string): Promise<Response> {
  if (request.method !== 'GET' && request.method !== 'HEAD') {
    return refuse(405, 'Method not allowed', { allow: 'GET, HEAD' });
  }

  const key = duckdbRuntimeKey(pathname);
  if (key === null) return refuse(404, 'Not a DuckDB runtime file');

  // No ASSETS fallback, unlike /logs-data/*: these bytes are never published as static
  // assets, so a miss means the upload in web/README.md's user-owned steps has not been
  // done for this version and there is nothing else to serve.
  const object = await env.LOGS?.get(key);
  if (object === null || object === undefined) return refuse(404, 'That DuckDB runtime file is not uploaded');

  // Set here rather than read from the object: WebAssembly.instantiateStreaming refuses
  // anything but application/wasm, and that would then depend on the content type whoever
  // ran the upload happened to pass.
  const headers = new Headers({
    'content-type': 'application/wasm',
    'cache-control': 'public, max-age=31536000, immutable',
    etag: object.httpEtag,
  });
  return new Response(request.method === 'HEAD' ? null : object.body, { status: 200, headers });
}

/**
 * Shared build pages, rendered by the Go API. Extracted alongside the other named route
 * handlers for consistency with the shape the R2 route above established -- one early-return
 * check per route in `fetch()`, dispatching to its own standalone function -- though its
 * behaviour is unchanged from before this file grew a shell router.
 */
async function serveBuildPage(request: Request, env: Env, url: URL): Promise<Response> {
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
}

/**
 * The four Worker-served shells. Each prefix maps to one built asset; the id, slug or
 * character path in the URL is what the island reads at runtime and what the head values
 * are rewritten from here.
 *
 * The rankings, character and guild shells arrive with Tasks 19 and 20. Until then their
 * asset is missing and ASSETS answers 404, which is exactly what should happen.
 */
const SHELL_ROUTES = [
  { prefix: '/reports/', asset: '/reports.html' },
  { prefix: '/rankings/', asset: '/rankings.html' },
  { prefix: '/character/', asset: '/character.html' },
  { prefix: '/guild/', asset: '/guild.html' },
] as const;

const REPORT_ID = /^\/reports\/([a-z2-7]{12})$/;
const RANKINGS_SLUG = /^\/rankings\/([a-z0-9-]{1,64})$/;

/** Unwraps the Phase 0 envelope, or null for any failure at all: a shell is never worth an error page. */
async function apiData<T>(url: string): Promise<T | null> {
  try {
    // No cookies, no client headers: this response is cached at the edge for everyone, so
    // nothing about the individual visitor may influence it. A fresh Request (rather than a
    // bare url string) carries that empty header set explicitly rather than by omission.
    const response = await fetch(new Request(url));
    if (!response.ok) return null;
    const envelope = (await response.json()) as { ok: boolean; data: T | null };
    return envelope.ok ? envelope.data : null;
  } catch {
    return null;
  }
}

interface ShellHead {
  meta: ShellMeta;
  /** False for unlisted, private and guild reports, which must not be indexed. */
  indexable: boolean;
}

async function shellHead(url: URL, env: Env): Promise<ShellHead | null> {
  const report = REPORT_ID.exec(url.pathname);
  if (report !== null) {
    // The same gate /logs-data/* uses (Task 4): a report's title, zone and kill count are
    // never fetched -- let alone rewritten into a cached, anonymous-readable response --
    // until the cheap visibility endpoint says the report is public or unlisted. This is
    // deliberately a second, independent check from GET /v1/reports/{id}'s own answer: that
    // endpoint is not guaranteed to refuse a private or guild report to an anonymous caller,
    // and this route must not trust it to.
    const lookup = await reportVisibility(report[1], env);
    if (lookup.kind !== 'ok' || !PUBLIC_VISIBILITIES.has(lookup.visibility)) return null;
    const data = await apiData<ReportMeta>(`${env.API_BASE_URL}/v1/reports/${report[1]}`);
    if (data === null) return null;
    return { meta: reportShellMeta(data, env.API_BASE_URL), indexable: data.visibility === 'public' };
  }

  const rankings = RANKINGS_SLUG.exec(url.pathname);
  if (rankings !== null) {
    const data = await apiData<RankingsHead>(
      `${env.API_BASE_URL}/v1/rankings?encounter=${rankings[1]}&page=1`,
    );
    if (data === null) return null;
    return { meta: rankingsShellMeta(rankings[1], data), indexable: true };
  }

  const character = parseCharacterPath(url.pathname);
  if (character !== null) {
    const data = await apiData<CharacterHead>(
      `${env.API_BASE_URL}/v1/characters/${character.region}/${character.ruleset}/${character.slug}`,
    );
    if (data === null) return null;
    return { meta: characterShellMeta(character, data), indexable: true };
  }

  const guild = parseGuildPath(url.pathname);
  if (guild !== null) {
    const data = await apiData<GuildHead>(
      `${env.API_BASE_URL}/v1/guilds/${guild.region}/${guild.ruleset}/${guild.slug}`,
    );
    if (data === null) return null;
    return { meta: guildShellMeta(guild, data), indexable: true };
  }

  return null;
}

/**
 * Whether this path names one report, encounter, character or guild. A path under a shell
 * prefix that does not -- `/reports/fixture2abcd/extra`, a prerendered fixture page's own
 * asset request -- belongs to ASSETS, not to the shell.
 */
function shellPathIsAddressable(url: URL): boolean {
  return (
    REPORT_ID.test(url.pathname) ||
    RANKINGS_SLUG.test(url.pathname) ||
    parseCharacterPath(url.pathname) !== null ||
    parseGuildPath(url.pathname) !== null
  );
}

/**
 * HTMLRewriter is a Workers-runtime global. It is reached through globalThis rather than
 * used bare so the tests can stub it and so a runtime without it -- `vitest`, a future
 * local preview -- serves the shell unrewritten instead of throwing. The island renders
 * the page either way; only the unfurl is lost.
 */
type ElementHandlers = {
  element(element: { setAttribute(n: string, v: string): void; setInnerContent(t: string): void }): void;
};
type Rewriter = {
  on(selector: string, handlers: ElementHandlers): Rewriter;
  transform(response: Response): Response;
};
type RewriterConstructor = new () => Rewriter;

function rewriteHead(response: Response, meta: ShellMeta): Response {
  const Ctor = (globalThis as { HTMLRewriter?: RewriterConstructor }).HTMLRewriter;
  if (Ctor === undefined) return response;
  return new Ctor()
    .on('title', { element: (element) => element.setInnerContent(meta.title) })
    .on('meta[data-og="description"]', {
      element: (element) => element.setAttribute('content', meta.description),
    })
    .on('meta[data-og="og-title"]', { element: (element) => element.setAttribute('content', meta.title) })
    .on('meta[data-og="og-description"]', {
      element: (element) => element.setAttribute('content', meta.description),
    })
    .on('meta[data-og="og-url"]', { element: (element) => element.setAttribute('content', meta.canonical) })
    .on('meta[data-og="og-image"]', { element: (element) => element.setAttribute('content', meta.image) })
    .on('link[data-og="canonical"]', { element: (element) => element.setAttribute('href', meta.canonical) })
    .transform(response);
}

async function serveShell(env: Env, url: URL, asset: string): Promise<Response> {
  const shell = await env.ASSETS.fetch(new Request(new URL(asset, url).toString(), { method: 'GET' }));
  if (!shell.ok) return shell;

  const head = await shellHead(url, env);
  const headers = new Headers(shell.headers);
  // One minute: a report's title and fight count change while a raid night is being
  // logged, and an unfurl a crawler fetched an hour ago should not be the one people see.
  headers.set('cache-control', 'public, max-age=60');
  // Carries no confidential data either way, but a crawler indexing an un-rewritten
  // placeholder for a real id (a private report it was refused, an id nothing answered for)
  // confirms that id exists, so both the no-data and the not-indexable case are noindex.
  if (head === null || !head.indexable) headers.set('x-robots-tag', 'noindex');

  const body = new Response(shell.body, { status: 200, headers });
  return head === null ? body : rewriteHead(body, head.meta);
}

export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const url = new URL(request.url);
    const logsData = LOGS_DATA_PATH.exec(url.pathname);
    if (logsData !== null) return serveReportFile(request, env, logsData[1], logsData[2]);
    if (url.pathname.startsWith(LOGS_DATA_PREFIX)) return refuse(404, 'Not a report file');

    if (url.pathname.startsWith(DUCKDB_RUNTIME_PREFIX)) {
      return serveDuckdbRuntime(request, env, url.pathname);
    }

    const shell = SHELL_ROUTES.find((route) => url.pathname.startsWith(route.prefix));
    // A path under a shell prefix that is not a valid id, slug or character path -- a
    // prerendered fixture page's own asset request, a stray segment -- belongs to ASSETS,
    // which serves it or 404s there.
    if (shell !== undefined && shellPathIsAddressable(url)) {
      return serveShell(env, url, shell.asset);
    }

    if (!BUILD_PATH.test(url.pathname)) return env.ASSETS.fetch(request);

    return serveBuildPage(request, env, url);
  },
};
