// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { ARRIVAL_CAPTURE_SCRIPT } from './arrival';
import { ARRIVAL_KEY, ARRIVAL_NEXT_KEY, ARRIVAL_TTL_MS, BEACON_SRC, BEACON_TOKEN } from './config';
import { loadBeacon, promoteArrival, sendArrival } from './later';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

/** Runs the inline capture the way the page does. */
function capture(): void {
  new Function(ARRIVAL_CAPTURE_SCRIPT)();
}

function signIn(): void {
  document.cookie = 'fs_csrf=tok%2Ben';
}

function stored(key = ARRIVAL_KEY): { slug: string; at: number } | null {
  const raw = localStorage.getItem(key);
  return raw === null ? null : (JSON.parse(raw) as { slug: string; at: number });
}

function keep(slug: string, at = Date.now()): void {
  localStorage.setItem(ARRIVAL_KEY, JSON.stringify({ slug, at }));
}

beforeEach(() => {
  localStorage.clear();
  document.cookie = 'fs_csrf=; expires=Thu, 01 Jan 1970 00:00:00 GMT';
  document.querySelectorAll('script').forEach((s) => s.remove());
  history.replaceState(null, '', '/');
});

afterEach(() => vi.unstubAllGlobals());

describe('the inline capture', () => {
  it('keeps a valid slug and takes it out of the address bar', () => {
    history.replaceState(null, '', '/sim?spec=fury&golink=reddit-sim#top');
    capture();
    expect(stored(ARRIVAL_NEXT_KEY)?.slug).toBe('reddit-sim');
    expect(location.pathname + location.search + location.hash).toBe('/sim?spec=fury#top');
  });

  it('drops a malformed slug but still cleans the address bar', () => {
    history.replaceState(null, '', '/?golink=%3Cscript%3E');
    capture();
    expect(stored()).toBeNull();
    expect(location.search).toBe('');
  });

  it('never touches the simulator’s own ?ref= (a character or a fight)', () => {
    history.replaceState(null, '', '/sim?source=armory&ref=us%2Fnormal%2Fobnoxious&mode=compare');
    capture();
    expect(stored()).toBeNull();
    expect(location.search).toBe('?source=armory&ref=us%2Fnormal%2Fobnoxious&mode=compare');
  });

  it('leaves a page with no golink alone', () => {
    history.replaceState(null, '', '/guides?class=rogue');
    capture();
    expect(stored()).toBeNull();
    expect(location.search).toBe('?class=rogue');
  });

  it('stays small: every page’s HTML carries it', () => {
    expect(ARRIVAL_CAPTURE_SCRIPT.length).toBeLessThan(400);
    expect(ARRIVAL_CAPTURE_SCRIPT).not.toContain('cloudflareinsights');
  });
});

describe('the arrival, after paint', () => {
  it('sends the kept slug with the CSRF token, then forgets it', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => new Response('{}', { status: 200 }));
    vi.stubGlobal('fetch', upstream);
    keep('reddit');
    signIn();
    await sendArrival(API);
    expect(upstream).toHaveBeenCalledTimes(1);
    const [url, init] = upstream.mock.calls[0];
    expect(url).toBe(`${API}/v1/me/arrival`);
    expect(init?.method).toBe('POST');
    expect(init?.credentials).toBe('include');
    expect(new Headers(init?.headers).get('x-csrf-token')).toBe('tok+en');
    expect(JSON.parse(String(init?.body))).toEqual({ slug: 'reddit', at: expect.any(Number) as number });
    expect(stored()).toBeNull();
  });

  it('keeps the slug for later while signed out', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    keep('discord');
    await sendArrival(API);
    expect(upstream).not.toHaveBeenCalled();
    expect(stored()?.slug).toBe('discord');
  });

  it('keeps it when the session turns out to be gone, so the sign-in still sends it', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => new Response('{}', { status: 401 })),
    );
    keep('wago');
    signIn();
    await sendArrival(API);
    expect(stored()?.slug).toBe('wago');
  });

  it('drops an arrival older than the API would take, unsent', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    keep('wago', Date.now() - ARRIVAL_TTL_MS - 1);
    signIn();
    await sendArrival(API);
    expect(upstream).not.toHaveBeenCalled();
    expect(stored()).toBeNull();
  });

  it('drops a malformed one, unsent', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    localStorage.setItem(ARRIVAL_KEY, '{"slug":"Not A Slug","at":1}');
    signIn();
    await sendArrival(API);
    expect(upstream).not.toHaveBeenCalled();
    expect(stored()).toBeNull();
  });
});

describe('first touch', () => {
  it('keeps the first link a browser carried, not a later one', () => {
    const t0 = Date.now() - 5 * 24 * 60 * 60 * 1000;
    keep('grimbrew', t0);
    history.replaceState(null, '', '/?golink=lunawick');
    capture();
    expect(promoteArrival(localStorage, Date.now())).toEqual({ slug: 'grimbrew', at: t0 });
    expect(stored(ARRIVAL_NEXT_KEY)).toBeNull();
  });

  it('takes the newest click once the first has gone stale', () => {
    keep('grimbrew', Date.now() - ARRIVAL_TTL_MS - 1);
    history.replaceState(null, '', '/?golink=lunawick');
    capture();
    expect(promoteArrival(localStorage, Date.now())?.slug).toBe('lunawick');
    expect(stored()?.slug).toBe('lunawick');
  });

  it('keeps a click for 30 days, the creator link window', () => {
    expect(ARRIVAL_TTL_MS).toBe(30 * 24 * 60 * 60 * 1000);
    keep('grimbrew', Date.now() - ARRIVAL_TTL_MS + 60_000);
    expect(promoteArrival(localStorage, Date.now())?.slug).toBe('grimbrew');
  });

  it('sends the first click with its time', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => new Response('{}', { status: 200 }));
    vi.stubGlobal('fetch', upstream);
    const t0 = Date.now() - 3 * 24 * 60 * 60 * 1000;
    keep('grimbrew', t0);
    history.replaceState(null, '', '/?golink=lunawick');
    capture();
    signIn();
    await sendArrival(API);
    expect(JSON.parse(String(upstream.mock.calls[0][1]?.body))).toEqual({ slug: 'grimbrew', at: t0 });
  });
});

describe('the visit beacon', () => {
  const beacons = () => [...document.head.querySelectorAll<HTMLScriptElement>(`script[src="${BEACON_SRC}"]`)];

  it('loads once on the production host, cookieless, SPA tracking off', () => {
    loadBeacon(location.hostname);
    loadBeacon(location.hostname);
    expect(beacons()).toHaveLength(1);
    expect(JSON.parse(beacons()[0].getAttribute('data-cf-beacon') ?? '')).toEqual({
      token: BEACON_TOKEN,
      spa: false,
    });
  });

  it('never loads anywhere else', () => {
    loadBeacon();
    expect(beacons()).toHaveLength(0);
  });

  it('never loads on the apps’ sign-in hand-off page', () => {
    history.replaceState(null, '', '/app/auth');
    loadBeacon(location.hostname);
    expect(beacons()).toHaveLength(0);
  });
});
