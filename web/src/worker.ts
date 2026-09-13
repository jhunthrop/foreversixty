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

export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const url = new URL(request.url);
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

    return new Response(upstream.body, {
      status: upstream.status,
      statusText: upstream.statusText,
      headers: upstream.headers,
    });
  },
};
