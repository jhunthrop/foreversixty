import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { forgetGoLinks, landingUrl, serveGoLink } from './redirect';

const API = 'https://api.foreversixty.test';
const env = { API_BASE_URL: API };

type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

function context() {
  const pending: Promise<unknown>[] = [];
  return { pending, waitUntil: vi.fn((p: Promise<unknown>) => void pending.push(p)) };
}

function api(target: string | null, status = 200) {
  return vi.fn<GlobalFetch>(async (input) => {
    const request = input as Request;
    if (request.method === 'POST') return new Response(null, { status: 204 });
    if (target === null) return envelope(null, status);
    return envelope({ slug: 'reddit', target });
  });
}

const go = (path: string, init: RequestInit = {}) => new Request(`https://foreversixty.gg${path}`, init);
const HUMAN = {
  'user-agent':
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 Version/17.1 Safari/605.1.15',
};

beforeEach(() => forgetGoLinks());
afterEach(() => vi.unstubAllGlobals());

describe('landingUrl', () => {
  it('keeps a target on this site and hands the slug along', () => {
    expect(landingUrl('/', 'reddit', 'https://foreversixty.gg')).toBe(
      'https://foreversixty.gg/?golink=reddit',
    );
    expect(landingUrl('/premium?plan=guild#top', 'x', 'https://foreversixty.gg')).toBe(
      'https://foreversixty.gg/premium?plan=guild&golink=x#top',
    );
  });

  it.each(['//evil.example', 'https://evil.example/', '/\\evil.example', 'evil', ''])(
    'refuses %j',
    (target) => {
      expect(landingUrl(target, 'x', 'https://foreversixty.gg')).toBeNull();
    },
  );
});

describe('serveGoLink', () => {
  it('redirects to the target with the slug, uncached, and counts a person after answering', async () => {
    const upstream = api('/sim');
    vi.stubGlobal('fetch', upstream);
    const ctx = context();
    const response = await serveGoLink(go('/go/reddit-sim', { headers: HUMAN }), env, ctx);
    expect(response.status).toBe(302);
    expect(response.headers.get('location')).toBe('https://foreversixty.gg/sim?golink=reddit-sim');
    expect(response.headers.get('cache-control')).toBe('no-store');
    expect(ctx.waitUntil).toHaveBeenCalledTimes(1);
    await Promise.all(ctx.pending);
    const count = upstream.mock.calls.map((c) => c[0] as Request).find((r) => r.method === 'POST');
    expect(count?.url).toBe(`${API}/v1/golinks/reddit-sim/clicks`);
    expect(await count?.json()).toEqual({ bot: false });
    // Nothing about the visitor goes to the API.
    expect(count?.headers.get('user-agent')).toBeNull();
    expect(count?.headers.get('x-forwarded-for')).toBeNull();
  });

  it('counts a link preview as a bot, and a HEAD as one too', async () => {
    const upstream = api('/');
    vi.stubGlobal('fetch', upstream);
    for (const init of [
      { headers: { 'user-agent': 'Mozilla/5.0 (compatible; Discordbot/2.0)' } },
      { method: 'HEAD', headers: HUMAN },
    ]) {
      const ctx = context();
      const response = await serveGoLink(go('/go/discord', init), env, ctx);
      expect(response.status).toBe(302);
      await Promise.all(ctx.pending);
    }
    const bodies = await Promise.all(
      upstream.mock.calls
        .map((c) => c[0] as Request)
        .filter((r) => r.method === 'POST')
        .map((r) => r.json()),
    );
    expect(bodies).toEqual([{ bot: true }, { bot: true }]);
  });

  it('resolves a slug once a minute, not once a click', async () => {
    const upstream = api('/');
    vi.stubGlobal('fetch', upstream);
    for (let i = 0; i < 3; i++) await serveGoLink(go('/go/wago', { headers: HUMAN }), env, context());
    const reads = upstream.mock.calls.map((c) => c[0] as Request).filter((r) => r.method === 'GET');
    expect(reads).toHaveLength(1);
  });

  it('sends an unknown slug home without counting it', async () => {
    vi.stubGlobal('fetch', api(null, 404));
    const ctx = context();
    const response = await serveGoLink(go('/go/nope', { headers: HUMAN }), env, ctx);
    expect(response.headers.get('location')).toBe('https://foreversixty.gg/');
    expect(ctx.waitUntil).not.toHaveBeenCalled();
  });

  it('sends a malformed slug home without asking the API', async () => {
    const upstream = api('/');
    vi.stubGlobal('fetch', upstream);
    for (const path of ['/go/Bad_Slug', '/go/', '/go/a/b']) {
      const response = await serveGoLink(go(path), env, context());
      expect(response.headers.get('location')).toBe('https://foreversixty.gg/');
    }
    expect(upstream).not.toHaveBeenCalled();
  });

  it('never follows a target off the site, even if the API answered one', async () => {
    vi.stubGlobal('fetch', api('//evil.example/phish'));
    const ctx = context();
    const response = await serveGoLink(go('/go/reddit', { headers: HUMAN }), env, ctx);
    expect(response.headers.get('location')).toBe('https://foreversixty.gg/');
    expect(ctx.waitUntil).not.toHaveBeenCalled();
  });

  it('still sends the visitor home when the API is down', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new Error('down');
      }),
    );
    const response = await serveGoLink(go('/go/reddit', { headers: HUMAN }), env, context());
    expect(response.status).toBe(302);
    expect(response.headers.get('location')).toBe('https://foreversixty.gg/');
  });

  it('refuses anything but a read', async () => {
    vi.stubGlobal('fetch', api('/'));
    const response = await serveGoLink(go('/go/reddit', { method: 'POST' }), env, context());
    expect(response.status).toBe(405);
  });
});
