// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { ARRIVAL_SCRIPT, arrivalScriptData } from './arrival';
import { ARRIVAL_KEY, ARRIVAL_TTL_MS, BEACON_SRC, BEACON_TOKEN } from './config';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

/** Runs the inline script the way the browser does: as the current script, with its data. */
function run(overrides: Record<string, string> = {}): void {
  const el = document.createElement('script');
  for (const [name, value] of Object.entries({ ...arrivalScriptData(API), ...overrides })) {
    el.setAttribute(name, value);
  }
  document.body.append(el);
  Object.defineProperty(document, 'currentScript', { value: el, configurable: true });
  new Function(ARRIVAL_SCRIPT)();
  Object.defineProperty(document, 'currentScript', { value: null, configurable: true });
}

function signIn(): void {
  document.cookie = 'fs_csrf=tok%2Ben';
}

function stored(): { slug: string; at: number } | null {
  const raw = localStorage.getItem(ARRIVAL_KEY);
  return raw === null ? null : (JSON.parse(raw) as { slug: string; at: number });
}

beforeEach(() => {
  vi.useFakeTimers();
  localStorage.clear();
  document.cookie = 'fs_csrf=; expires=Thu, 01 Jan 1970 00:00:00 GMT';
  document.head.querySelectorAll('script').forEach((s) => s.remove());
  history.replaceState(null, '', '/');
});

afterEach(() => {
  vi.useRealTimers();
  vi.unstubAllGlobals();
});

describe('the landing', () => {
  it('keeps a valid slug and takes it out of the address bar', () => {
    history.replaceState(null, '', '/sim?spec=fury&golink=reddit-sim#top');
    run();
    expect(stored()?.slug).toBe('reddit-sim');
    expect(location.pathname + location.search + location.hash).toBe('/sim?spec=fury#top');
  });

  it('drops a malformed slug but still cleans the address bar', () => {
    history.replaceState(null, '', '/?golink=%3Cscript%3E');
    run();
    expect(stored()).toBeNull();
    expect(location.search).toBe('');
  });

  it('never touches the simulator’s own ?ref= (a character or a fight)', () => {
    history.replaceState(null, '', '/sim?source=armory&ref=us%2Fnormal%2Fobnoxious&mode=compare');
    run();
    expect(stored()).toBeNull();
    expect(location.search).toBe('?source=armory&ref=us%2Fnormal%2Fobnoxious&mode=compare');
  });

  it('leaves a page with no golink alone', () => {
    history.replaceState(null, '', '/guides?class=rogue');
    run();
    expect(stored()).toBeNull();
    expect(location.search).toBe('?class=rogue');
  });
});

describe('the first signed-in page', () => {
  it('sends the slug after load with the CSRF token, then forgets it', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => new Response('{}', { status: 200 }));
    vi.stubGlobal('fetch', upstream);
    history.replaceState(null, '', '/?golink=reddit');
    signIn();
    run();
    expect(upstream).not.toHaveBeenCalled();
    await vi.runAllTimersAsync();
    expect(upstream).toHaveBeenCalledTimes(1);
    const [url, init] = upstream.mock.calls[0];
    expect(url).toBe(`${API}/v1/me/arrival`);
    expect(init?.method).toBe('POST');
    expect(init?.credentials).toBe('include');
    expect(new Headers(init?.headers).get('x-csrf-token')).toBe('tok+en');
    expect(JSON.parse(String(init?.body))).toEqual({ slug: 'reddit' });
    expect(stored()).toBeNull();
  });

  it('keeps the slug for later while signed out', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    history.replaceState(null, '', '/?golink=discord');
    run();
    await vi.runAllTimersAsync();
    expect(upstream).not.toHaveBeenCalled();
    expect(stored()?.slug).toBe('discord');
  });

  it('keeps it when the session turns out to be gone, so the sign-in still sends it', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => new Response('{}', { status: 401 })),
    );
    localStorage.setItem(ARRIVAL_KEY, JSON.stringify({ slug: 'wago', at: Date.now() }));
    signIn();
    run();
    await vi.runAllTimersAsync();
    expect(stored()?.slug).toBe('wago');
  });

  it('drops an arrival older than the API would take, unsent', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    localStorage.setItem(ARRIVAL_KEY, JSON.stringify({ slug: 'wago', at: Date.now() - ARRIVAL_TTL_MS - 1 }));
    signIn();
    run();
    await vi.runAllTimersAsync();
    expect(upstream).not.toHaveBeenCalled();
    expect(stored()).toBeNull();
  });
});

describe('the visit beacon', () => {
  const beacons = () => [...document.head.querySelectorAll<HTMLScriptElement>(`script[src="${BEACON_SRC}"]`)];

  it('loads after load on the production host, cookieless, SPA tracking off', async () => {
    run({ 'data-beacon-host': location.hostname });
    expect(beacons()).toHaveLength(0);
    await vi.runAllTimersAsync();
    expect(beacons()).toHaveLength(1);
    expect(JSON.parse(beacons()[0].getAttribute('data-cf-beacon') ?? '')).toEqual({
      token: BEACON_TOKEN,
      spa: false,
    });
  });

  it('never loads anywhere else', async () => {
    run();
    await vi.runAllTimersAsync();
    expect(beacons()).toHaveLength(0);
  });

  it('never loads on the apps’ sign-in hand-off page', async () => {
    history.replaceState(null, '', '/app/auth');
    run({ 'data-beacon-host': location.hostname });
    await vi.runAllTimersAsync();
    expect(beacons()).toHaveLength(0);
  });
});
