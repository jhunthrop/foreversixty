// @vitest-environment jsdom
// web/src/lib/account/api.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { mockEntitlements, premiumEntitlements } from '../../test-support/billing-fixtures';
import { invalidate, subscribe } from '../data/query';
import {
  ACCOUNT_FAILED,
  AccountError,
  type BillingView,
  battlenetStartUrl,
  companionsOf,
  connectDeviceApprove,
  connectDeviceDeny,
  connectDeviceLookup,
  csrfToken,
  effectiveServerSims,
  fetchMe,
  fetchMeOnce,
  forgetSession,
  listCompanions,
  listDevices,
  meKey,
  pairDevice,
  postMyExports,
  queueBuildToAddon,
  REPORTS_PER_PAGE,
  requestEmailLink,
  requestEnvelope,
  revokeDevice,
  setAnonymize,
  setMainCharacter,
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
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
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
  // setAnonymize/setMainCharacter now write into query.ts's shared, module-level cache
  // (setQueryData), so every test in this file -- not just the ones under the `fetchMeOnce`
  // describe below -- must clear it, or a later test's `fetchMeOnce(API)` sees a stale
  // "fresh" entry an earlier test left behind and never calls `fetch` at all. `forgetSession`
  // only forgets `scope: 'private'` entries; `invalidate('')` (every key) also resets one
  // `setQueryData` can create with no prior `query()` call for that key, which defaults to
  // `scope: 'public'` and so survives `forgetSession`.
  forgetSession();
  invalidate('');
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

  it('a signed-in phone app is a device but not a companion', async () => {
    const devices = [
      {
        id: 'pc',
        name: 'Raid PC',
        platform: 'windows',
        kind: 'companion' as const,
        created_at: 't',
        last_seen_at: null,
      },
      {
        id: 'ph',
        name: 'iPhone',
        platform: 'ios',
        kind: 'app' as const,
        created_at: 't',
        last_seen_at: null,
      },
      // An API from before devices carried a kind lists companions only.
      { id: 'old', name: 'Old PC', platform: 'linux', created_at: 't', last_seen_at: null },
    ];
    expect(companionsOf(devices).map((d) => d.id)).toEqual(['pc', 'old']);
    vi.stubGlobal('fetch', vi.fn().mockResolvedValue(envelope(devices)));
    await expect(listDevices(API)).resolves.toHaveLength(3);
    await expect(listCompanions(API)).resolves.toEqual([devices[0], devices[2]]);
  });

  it('listDevices shares a cached answer across two calls without a second request', async () => {
    const fetchSpy = vi
      .fn()
      .mockResolvedValue(
        envelope([{ id: 'd1', name: 'Phone', platform: 'ios', created_at: 't', last_seen_at: null }]),
      );
    vi.stubGlobal('fetch', fetchSpy);
    await listDevices(API);
    await listDevices(API);
    expect(fetchSpy).toHaveBeenCalledTimes(1);
  });

  it('revokeDevice invalidates the devices list', async () => {
    const fetchSpy = vi
      .fn()
      .mockResolvedValueOnce(
        envelope([{ id: 'd1', name: 'Phone', platform: 'ios', created_at: 't', last_seen_at: null }]),
      )
      .mockResolvedValueOnce(envelope(null)) // the DELETE
      .mockResolvedValueOnce(envelope([]));
    vi.stubGlobal('fetch', fetchSpy);
    await listDevices(API);
    await revokeDevice('d1', API);
    await listDevices(API);
    expect(fetchSpy).toHaveBeenCalledTimes(3); // list, delete, re-list (invalidated, not cached)
  });

  it('signs out and sets the anonymize flag', async () => {
    // A dedicated apiBase, not the shared `API` constant: `setAnonymize` now writes its
    // response into query.ts's cache via `setQueryData`, and this test's fixture response
    // (`{ ok: true }`, not a real `Me`) would otherwise seed a bogus `${API}/v1/me` entry
    // that outlives this test and confuses the `fetchMeOnce` tests below, which share `API`.
    const apiBase = `${API}/signout-fixture`;
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ ok: true }));
    vi.stubGlobal('fetch', upstream);

    await signOut(apiBase);
    await setAnonymize(true, apiBase);

    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${apiBase}/v1/sessions`);
    expect((upstream.mock.calls[0][0] as Request).method).toBe('DELETE');
    const patch = upstream.mock.calls[1][0] as Request;
    expect(patch.method).toBe('PATCH');
    expect(patch.url).toBe(`${apiBase}/v1/me`);
    expect(await patch.json()).toEqual({ anonymize: true });
  });

  it('a sign-out the server refuses keeps the session: the cached /v1/me is never dropped', async () => {
    const apiBase = `${API}/signout-refused`;
    const me = { user: { id: 1 }, characters: [], guilds: [], entitlements: mockEntitlements() };
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>().mockResolvedValueOnce(envelope(me)).mockResolvedValueOnce(envelope(null, 500)),
    );
    await fetchMeOnce(apiBase);
    const statuses: string[] = [];
    const unsubscribe = subscribe(meKey(apiBase), (state) => statuses.push(state.status));

    await expect(signOut(apiBase)).rejects.toThrow();
    unsubscribe();

    expect(statuses).toEqual(['ready']);
  });

  it.each([
    ['setAnonymize', (base: string) => setAnonymize(true, base)],
    ['setMainCharacter', (base: string) => setMainCharacter('us/hardcore/elyra', base)],
  ])('%s caches the PATCH answer with its price facts normalized', async (name, patch) => {
    const apiBase = `${API}/patch-normalize-${name}`;
    const me = {
      user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
      characters: [],
      guilds: [],
      entitlements: mockEntitlements({
        billing: {
          plan: 'premium',
          status: 'active',
          current_period_end: null,
          cancel_at_period_end: false,
          founder: false,
          interval: '',
          amount_cents: '',
        } as unknown as BillingView,
      }),
    };
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(me)),
    );
    await patch(apiBase);
    const cached = await fetchMeOnce(apiBase);
    expect(cached?.entitlements?.billing).toMatchObject({ interval: null, amount_cents: null });
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

  it('posts a signed-in paste to /v1/me/exports and returns the written characters', async () => {
    document.cookie = 'fs_csrf=abc; path=/';
    const written = [
      { key: 'us/hardcore/elyra-duskvale', region: 'us', ruleset: 'hardcore', name: 'Elyra Duskvale' },
    ];
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ characters: written }));
    vi.stubGlobal('fetch', upstream);

    const input = [{ name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', export: 'FS1:1:priest:...' }];
    await expect(postMyExports(input, API)).resolves.toEqual(written);

    const request = upstream.mock.calls[0][0] as Request;
    expect(request.method).toBe('POST');
    expect(request.url).toBe(`${API}/v1/me/exports`);
    expect(request.headers.get('x-csrf-token')).toBe('abc');
    expect(await request.json()).toEqual({ exports: input });
  });

  it('surfaces the API error message when a paste save fails', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(
        async () =>
          new Response(
            JSON.stringify({
              ok: false,
              data: null,
              error: { message: 'That export is too long.' },
              request_id: 'r',
            }),
            { status: 400, headers: { 'content-type': 'application/json' } },
          ),
      ),
    );
    await expect(
      postMyExports([{ name: 'X', region: 'us', ruleset: 'pvp', export: 'FS1:...' }], API),
    ).rejects.toThrow('That export is too long.');
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

  it('shares a cached answer across two calls without a second request', async () => {
    const fetchSpy = vi
      .fn()
      .mockResolvedValue(envelope({ rows: [], total: 0, page: 1, per_page: REPORTS_PER_PAGE }));
    vi.stubGlobal('fetch', fetchSpy);
    const { listMyReports } = await import('./api');
    await listMyReports(1, API);
    await listMyReports(1, API);
    expect(fetchSpy).toHaveBeenCalledTimes(1);
  });
});

describe('fetchMeOnce', () => {
  it('asks the API once however many islands want the session', async () => {
    const fetchMock = vi.fn<GlobalFetch>(async () => envelope(ME));
    vi.stubGlobal('fetch', fetchMock);
    const [first, second] = await Promise.all([fetchMeOnce(API), fetchMeOnce(API)]);
    expect(first).toEqual(second);
    expect(fetchMock).toHaveBeenCalledTimes(1);
  });

  it('does not remember a failure, so the next caller asks again', async () => {
    const fetchMock = vi
      .fn<GlobalFetch>()
      .mockRejectedValueOnce(new TypeError('offline'))
      .mockResolvedValueOnce(envelope(ME));
    vi.stubGlobal('fetch', fetchMock);
    await expect(fetchMeOnce(API)).rejects.toBeDefined();
    await expect(fetchMeOnce(API)).resolves.toMatchObject({ user: { id: 7 } });
    expect(fetchMock).toHaveBeenCalledTimes(2);
  });

  it('shares one in-flight request across concurrent callers on a page', async () => {
    const fetchSpy = vi.fn().mockResolvedValue(envelope(ME));
    vi.stubGlobal('fetch', fetchSpy);
    const [a, b] = await Promise.all([fetchMeOnce(API), fetchMeOnce(API)]);
    expect(a).toMatchObject(ME);
    expect(a?.guilds[0]?.subscribe).toBeNull();
    expect(b).toMatchObject(ME);
    expect(fetchSpy).toHaveBeenCalledTimes(1);
  });

  it('forgetSession clears the shared /v1/me cache: the next fetchMeOnce call is a real request', async () => {
    const fetchSpy = vi.fn().mockResolvedValue(envelope(ME));
    vi.stubGlobal('fetch', fetchSpy);
    await fetchMeOnce(API);
    forgetSession();
    await fetchMeOnce(API);
    expect(fetchSpy).toHaveBeenCalledTimes(2);
  });
});

describe('log_coach on a parsed /v1/me', () => {
  it.each([
    ['Premium', premiumEntitlements(), true],
    ['free', mockEntitlements(), false],
  ])('a %s account reads log_coach %s', async (_name, entitlements, want) => {
    const me = { ...ME, entitlements };
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(me)),
    );
    const parsed = await fetchMe(`${API}/log-coach-${want}`);
    expect(parsed?.entitlements?.log_coach).toBe(want);
  });
});

describe('effectiveServerSims', () => {
  it('is false for a signed-out visitor', () => {
    expect(effectiveServerSims(null)).toBe(false);
  });

  it('prefers entitlements.server_sims when present', () => {
    const me = {
      user: { id: 1, battletag: 'A', email: null, role: 'user', anonymize: false, premium: false },
      characters: [],
      guilds: [],
      entitlements: {
        server_sims: true,
        retention: false,
        multi_compare: false,
        history: false,
        notifications: false,
        officer_views: false,
        roster_check: false,
        log_coach: false,
        supporter_mark: false,
        billing: null,
      },
    } as unknown as Parameters<typeof effectiveServerSims>[0];
    expect(effectiveServerSims(me)).toBe(true);
  });

  it('falls back to user.premium when entitlements is absent (pre-API-lane fixtures)', () => {
    const me = {
      user: { id: 1, battletag: 'A', email: null, role: 'user', anonymize: false, premium: true },
      characters: [],
      guilds: [],
    } as unknown as Parameters<typeof effectiveServerSims>[0];
    expect(effectiveServerSims(me)).toBe(true);
  });
});

describe('ETag revalidation', () => {
  it('sends If-None-Match on a repeat GET after seeing an ETag, and resolves the 304 to the prior data', async () => {
    const first = new Response(JSON.stringify({ ok: true, data: { n: 1 }, error: null }), {
      status: 200,
      headers: { 'content-type': 'application/json', etag: 'W/"abc"' },
    });
    const second = new Response(null, { status: 304, headers: { etag: 'W/"abc"' } });
    const fetchSpy = vi.fn().mockResolvedValueOnce(first).mockResolvedValueOnce(second);
    vi.stubGlobal('fetch', fetchSpy);

    const a = await requestEnvelope<{ n: number }>('/v1/etag-demo', API);
    expect(a.data).toEqual({ n: 1 });

    const b = await requestEnvelope<{ n: number }>('/v1/etag-demo', API);
    expect(b.data).toEqual({ n: 1 }); // the 304's cached body
    expect(b.status).toBe(304);

    const sentHeaders = fetchSpy.mock.calls[1][0].headers as Headers;
    expect(sentHeaders.get('if-none-match')).toBe('W/"abc"');
  });

  it('never sends If-None-Match for a non-GET request', async () => {
    const fetchSpy = vi.fn().mockResolvedValue(envelope({ ok: true }));
    vi.stubGlobal('fetch', fetchSpy);
    await requestEnvelope('/v1/etag-demo-2', API, { method: 'POST', body: {} });
    const sentHeaders = fetchSpy.mock.calls[0][0].headers as Headers;
    expect(sentHeaders.has('if-none-match')).toBe(false);
  });
});

describe('companion connect calls', () => {
  it('looks a code up with the session cookie', async () => {
    const fetchMock = vi.fn<GlobalFetch>(async () =>
      envelope({ name: 'JHUNT-PC', platform: 'Windows', expires_in: 500 }),
    );
    vi.stubGlobal('fetch', fetchMock);
    const found = await connectDeviceLookup('KQ4-7TM', API);
    expect(found.name).toBe('JHUNT-PC');
    const request = fetchMock.mock.calls[0][0] as Request;
    expect(request.url).toBe(`${API}/v1/devices/connect/lookup?code=KQ4-7TM`);
    expect(request.credentials).toBe('include');
  });

  it('posts approve and deny with the code and the CSRF header', async () => {
    document.cookie = 'fs_csrf=tok; path=/';
    const fetchMock = vi.fn<GlobalFetch>(async (input) =>
      (input as Request).url.endsWith('/deny')
        ? new Response(null, { status: 204 })
        : envelope({ name: 'JHUNT-PC', platform: 'Windows' }),
    );
    vi.stubGlobal('fetch', fetchMock);
    expect((await connectDeviceApprove('KQ4-7TM', API)).name).toBe('JHUNT-PC');
    await connectDeviceDeny('KQ4-7TM', API);
    const [approve, deny] = fetchMock.mock.calls.map((c) => c[0] as Request);
    expect(approve.method).toBe('POST');
    expect(approve.url).toBe(`${API}/v1/devices/connect/approve`);
    expect(approve.headers.get('x-csrf-token')).toBe('tok');
    expect(await approve.clone().json()).toEqual({ code: 'KQ4-7TM' });
    expect(deny.url).toBe(`${API}/v1/devices/connect/deny`);
  });

  it('surfaces 404 and 409 as AccountError statuses', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 409)),
    );
    await expect(connectDeviceApprove('KQ4-7TM', API)).rejects.toMatchObject({ status: 409 });
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 404)),
    );
    await expect(connectDeviceLookup('KQ4-7TM', API)).rejects.toMatchObject({ status: 404 });
  });

  it('reports a network error as status 0', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => Promise.reject(new TypeError('offline'))),
    );
    await expect(connectDeviceApprove('KQ4-7TM', API)).rejects.toMatchObject({ status: 0 });
  });

  it('sends next with the email link request when asked', async () => {
    const fetchMock = vi.fn<GlobalFetch>(async () => envelope({ sent: true }, 202));
    vi.stubGlobal('fetch', fetchMock);
    await requestEmailLink('a@b.co', API, '/connect?code=KQ4-7TM');
    expect(await (fetchMock.mock.calls[0][0] as Request).clone().json()).toEqual({
      email: 'a@b.co',
      next: '/connect?code=KQ4-7TM',
    });
  });
});

describe('queueBuildToAddon', () => {
  it('posts the build id and character key with the CSRF header, and resolves on 201', async () => {
    document.cookie = 'fs_csrf=tok';
    const upstream = vi.fn<GlobalFetch>().mockResolvedValue(envelope({ build_id: 'k7x2qm4a' }, 201));
    vi.stubGlobal('fetch', upstream);

    await expect(queueBuildToAddon('k7x2qm4a', 'us/hardcore/obnoxious-yell', API)).resolves.toBeUndefined();

    const request = upstream.mock.calls[0][0] as Request;
    expect(request.method).toBe('POST');
    expect(request.url).toBe(`${API}/v1/addon/inbox`);
    expect(request.headers.get('x-csrf-token')).toBe('tok');
    expect(request.credentials).toBe('include');
    expect(await request.json()).toEqual({
      build_id: 'k7x2qm4a',
      character_key: 'us/hardcore/obnoxious-yell',
    });
  });

  it('omits the character key when none is given', async () => {
    const upstream = vi.fn<GlobalFetch>().mockResolvedValue(envelope({ build_id: 'k7x2qm4a' }, 201));
    vi.stubGlobal('fetch', upstream);
    await queueBuildToAddon('k7x2qm4a', undefined, API);
    expect(await (upstream.mock.calls[0][0] as Request).json()).toEqual({ build_id: 'k7x2qm4a' });
  });

  it('throws an AccountError carrying 401 for a signed-out visitor', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>().mockResolvedValue(envelope(null, 401)));
    await expect(queueBuildToAddon('k7x2qm4a', 'us/hardcore/x', API)).rejects.toMatchObject({
      name: 'AccountError',
      status: 401,
    });
  });

  it('throws an AccountError carrying 400 for an id the API refuses', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>().mockResolvedValue(envelope(null, 400)));
    const failure = await queueBuildToAddon('nope', undefined, API).catch((e: unknown) => e);
    expect(failure).toBeInstanceOf(AccountError);
    expect((failure as AccountError).status).toBe(400);
  });
});
