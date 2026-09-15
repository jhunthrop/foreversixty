// @vitest-environment jsdom
// web/src/lib/account/api.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  ACCOUNT_FAILED,
  AccountError,
  battlenetStartUrl,
  csrfToken,
  fetchMe,
  listDevices,
  pairDevice,
  requestEmailLink,
  revokeDevice,
  setAnonymize,
  signOut,
} from './api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
  characters: [
    {
      key: 'us/hardcore/elyra-duskvale',
      region: 'us',
      ruleset: 'hardcore',
      name: 'Elyra Duskvale',
      class: 'Priest',
    },
  ],
  guilds: [{ id: 3, region: 'us', ruleset: 'hardcore', name: 'The Last Watch', rank: 'officer' }],
};

afterEach(() => {
  vi.unstubAllGlobals();
  document.cookie = 'fs_csrf=; Max-Age=0; path=/';
});

describe('the account API', () => {
  it('reads the CSRF cookie the API sets', () => {
    document.cookie = 'other=1; path=/';
    document.cookie = 'fs_csrf=tok%3Aen; path=/';
    expect(csrfToken()).toBe('tok:en');
  });

  it('returns null rather than throwing when nobody is signed in', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 401)),
    );
    await expect(fetchMe(API)).resolves.toBeNull();
  });

  it('reads the signed-in user, their characters and their guilds', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(ME));
    vi.stubGlobal('fetch', upstream);

    const me = await fetchMe(API);

    expect(me?.user.battletag).toBe('Fixture#1234');
    expect(me?.characters[0].name).toBe('Elyra Duskvale');
    expect(me?.guilds[0].rank).toBe('officer');
    const request = upstream.mock.calls[0][0] as Request;
    expect(request.url).toBe(`${API}/v1/me`);
    expect(request.credentials).toBe('include');
  });

  it('builds the Battle.net start url with the page to come back to', () => {
    expect(battlenetStartUrl('/logs', API)).toBe(`${API}/v1/auth/battlenet/start?next=%2Flogs`);
  });

  it('posts an email link request with the CSRF header', async () => {
    document.cookie = 'fs_csrf=abc; path=/';
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ sent: true }));
    vi.stubGlobal('fetch', upstream);

    await expect(requestEmailLink('raider@example.com', API)).resolves.toBeUndefined();

    const request = upstream.mock.calls[0][0] as Request;
    expect(request.method).toBe('POST');
    expect(request.headers.get('x-csrf-token')).toBe('abc');
    expect(await request.json()).toEqual({ email: 'raider@example.com' });
  });

  it('surfaces the rate limit the contract sets on email links', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(
        async () =>
          new Response(
            JSON.stringify({
              ok: false,
              data: null,
              error: { message: 'Too many sign-in links; try again in an hour.' },
              request_id: 'r',
            }),
            { status: 429, headers: { 'content-type': 'application/json' } },
          ),
      ),
    );
    await expect(requestEmailLink('raider@example.com', API)).rejects.toThrow(
      'Too many sign-in links; try again in an hour.',
    );
  });

  it('lists, pairs and revokes devices', async () => {
    const devices = [
      {
        id: 'dev1',
        name: 'Raid PC',
        platform: 'windows',
        created_at: '2026-11-05T10:00:00Z',
        last_seen_at: null,
      },
    ];
    const upstream = vi.fn<GlobalFetch>(async (input) => {
      const request = input as Request;
      if (request.method === 'POST') return envelope({ code: '4821-9930', expires_in: 600 });
      if (request.method === 'DELETE') return envelope({ revoked: true });
      return envelope(devices);
    });
    vi.stubGlobal('fetch', upstream);

    await expect(listDevices(API)).resolves.toEqual(devices);
    await expect(pairDevice(API)).resolves.toEqual({ code: '4821-9930', expires_in: 600 });
    await expect(revokeDevice('dev1', API)).resolves.toBeUndefined();
    expect((upstream.mock.calls[2][0] as Request).url).toBe(`${API}/v1/devices/dev1`);
  });

  it('signs out and sets the anonymize flag', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ ok: true }));
    vi.stubGlobal('fetch', upstream);

    await signOut(API);
    await setAnonymize(true, API);

    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/sessions`);
    expect((upstream.mock.calls[0][0] as Request).method).toBe('DELETE');
    const patch = upstream.mock.calls[1][0] as Request;
    expect(patch.method).toBe('PATCH');
    expect(patch.url).toBe(`${API}/v1/me`);
    expect(await patch.json()).toEqual({ anonymize: true });
  });

  it('turns a network failure into an AccountError with a message a page can print', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(listDevices(API)).rejects.toBeInstanceOf(AccountError);
    await expect(listDevices(API)).rejects.toThrow(ACCOUNT_FAILED);
  });
});

describe('listMyReports', () => {
  it('asks for the signed-in user’s reports, one page at a time', async () => {
    const page = {
      rows: [
        {
          id: 'fixture2abcd',
          title: 'Sanguine Depths, fixture night',
          zone: 'Sanguine Depths',
          status: 'complete',
          visibility: 'public',
          created_at: '2026-09-26T20:09:00Z',
          fight_count: 3,
          kill_count: 1,
        },
      ],
      total: 1,
      page: 1,
      per_page: 100,
    };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(page));
    vi.stubGlobal('fetch', upstream);

    const { listMyReports } = await import('./api');
    await expect(listMyReports(2, API)).resolves.toEqual(page);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/reports?mine=1&page=2`);
  });

  it('treats a signed-out visitor as an empty page rather than an error', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 401)),
    );
    const { listMyReports } = await import('./api');
    await expect(listMyReports(1, API)).resolves.toEqual({ rows: [], total: 0, page: 1, per_page: 100 });
  });
});
