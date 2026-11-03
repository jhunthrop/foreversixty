// web/src/lib/golinks/redirect.ts
// foreversixty.gg/go/<slug>: the tracked short links. The Worker answers with a 302 to the
// link's target, a path on this site with the slug handed along as ?golink=, and counts the
// click after the answer has gone (the execution context's waitUntil), so a click never
// waits on the database. The count carries the slug and "bot or not" and nothing else: no
// IP, no user agent, no cookie reaches the API from here.
import { isBotAgent } from './bots';
import { GO_PREFIX, isSlug, LINK_PARAM } from './config';

/** The Workers runtime's execution context, declared inline like worker.ts's other bindings. */
export interface ExecutionContext {
  waitUntil(promise: Promise<unknown>): void;
}

export interface GoLinkEnv {
  API_BASE_URL: string;
}

/** The API caches a target for a minute; the isolate keeps it as long, a miss included. */
const TARGET_TTL_MS = 60_000;
/** An outage is not stable: retried after a few seconds, while every click still lands home. */
const UNAVAILABLE_TTL_MS = 5_000;

const targets = new Map<string, { target: string | null; expires: number }>();

/** Tests only: a fresh isolate has resolved nothing. */
export function forgetGoLinks(): void {
  targets.clear();
}

/**
 * The landing URL for a target, or null when the target is not a path on this site. The API
 * and the database refuse such a target already; this is the last check before a Location
 * header, so an open redirect needs all three to fail.
 */
export function landingUrl(target: string, slug: string, origin: string): string | null {
  if (!target.startsWith('/') || target.startsWith('//') || target.includes('\\')) return null;
  const url = new URL(target, origin);
  if (url.origin !== origin) return null;
  url.searchParams.set(LINK_PARAM, slug);
  return url.toString();
}

function remember(slug: string, target: string | null, ttl: number): string | null {
  targets.set(slug, { target, expires: Date.now() + ttl });
  return target;
}

async function resolveTarget(slug: string, env: GoLinkEnv): Promise<string | null> {
  const cached = targets.get(slug);
  if (cached !== undefined && cached.expires > Date.now()) return cached.target;
  let response: Response;
  try {
    response = await fetch(new Request(`${env.API_BASE_URL}/v1/golinks/${slug}`));
  } catch {
    return remember(slug, null, UNAVAILABLE_TTL_MS);
  }
  if (response.status === 404) return remember(slug, null, TARGET_TTL_MS);
  if (!response.ok) return remember(slug, null, UNAVAILABLE_TTL_MS);
  try {
    const envelope = (await response.json()) as { data?: { target?: unknown } | null };
    const target = envelope.data?.target;
    return remember(slug, typeof target === 'string' ? target : null, TARGET_TTL_MS);
  } catch {
    return remember(slug, null, UNAVAILABLE_TTL_MS);
  }
}

function countClick(slug: string, bot: boolean, env: GoLinkEnv): Promise<unknown> {
  return fetch(
    new Request(`${env.API_BASE_URL}/v1/golinks/${slug}/clicks`, {
      method: 'POST',
      headers: { 'content-type': 'application/json' },
      body: JSON.stringify({ bot }),
    }),
  ).catch(() => undefined);
}

function redirect(location: string): Response {
  return new Response(null, {
    status: 302,
    // Never cached anywhere, so every click reaches the Worker and is counted.
    headers: { location, 'cache-control': 'no-store', 'x-robots-tag': 'noindex' },
  });
}

/** The slug a /go/ path names, or null for any other shape. */
function slugOf(pathname: string): string | null {
  const rest = pathname.slice(GO_PREFIX.length).replace(/\/$/, '');
  return isSlug(rest) ? rest : null;
}

export async function serveGoLink(
  request: Request,
  env: GoLinkEnv,
  ctx?: ExecutionContext,
): Promise<Response> {
  if (request.method !== 'GET' && request.method !== 'HEAD') {
    return new Response('Method not allowed', { status: 405, headers: { allow: 'GET, HEAD' } });
  }
  const url = new URL(request.url);
  const home = new URL('/', url.origin).toString();
  const slug = slugOf(url.pathname);
  if (slug === null) return redirect(home);

  const target = await resolveTarget(slug, env);
  const landing = target === null ? null : landingUrl(target, slug, url.origin);
  if (landing === null) return redirect(home);

  // A HEAD is a link checker or a preview, never a person following the link.
  const bot = request.method === 'HEAD' || isBotAgent(request.headers.get('user-agent') ?? '');
  ctx?.waitUntil(countClick(slug, bot, env));
  return redirect(landing);
}
