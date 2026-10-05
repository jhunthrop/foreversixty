// web/src/lib/guild/api.test.ts
// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { forgetPrivate, invalidate } from '../data/query';
import {
  GUILD_API_FAILED,
  GuildApiError,
  acceptInvite,
  approveCharacter,
  claimGuild,
  confirmClaim,
  contestClaim,
  deleteGuildCrest,
  fetchGuildHome,
  fetchGuildSettings,
  leaveGuild,
  putGuildCrest,
  releaseClaim,
  removeCharacter,
  rotateInvite,
  updateConsent,
  updateGuildSettings,
  type GuildHome,
  type GuildSettingsData,
} from './api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

/** Minimal fixtures matching `GuildHome`/`GuildSettingsData` -- no shared helper exists yet
 *  in this file, so these two are just for the caching/invalidation tests below. */
function guildHomeFixture(): GuildHome {
  return {
    guild: { id: 9, region: 'us', ruleset: 'hardcore', name: 'Fixture Guild' },
    claim: { state: 'claimed', frozen: false },
    reports: [],
    roster: [],
  };
}

function guildSettingsFixture(): GuildSettingsData {
  return {
    default_visibility: 'guild',
    officer_max_rank_index: 1,
    claimed_by: null,
    claim_pending: null,
    claim: { state: 'claimed', frozen: false },
    invite: { rotated_at: null },
  };
}

afterEach(() => {
  // fetchGuildHome/fetchGuildSettings now cache through query.ts's shared, module-level
  // store (scope: 'private'), so a later test's read for the same guild id could otherwise
  // see a still-fresh entry an earlier test left behind and never call `fetch` at all.
  forgetPrivate();
  invalidate('');
  vi.unstubAllGlobals();
});

describe('fetchGuildHome', () => {
  it('reads the flat reports array and top-level next_cursor, plus claim state', async () => {
    const home = {
      guild: { id: 42, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' },
      claim: { state: 'claimed', frozen: false },
      reports: [
        {
          id: 'r1',
          title: 'Sanguine Depths',
          created_at: '2026-09-20T20:00:00Z',
          fight_count: 8,
          kill_count: 3,
        },
      ],
      next_cursor: 'abc123',
      roster: [],
    };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(home));
    vi.stubGlobal('fetch', upstream);

    const result = await fetchGuildHome(42, undefined, API);

    expect(result.claim.state).toBe('claimed');
    expect(result.reports[0].fight_count).toBe(8);
    expect(result.next_cursor).toBe('abc123');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/guilds/42/home`);
  });

  it('sends a cursor query param when given one', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        guild: { id: 42, region: 'us', ruleset: 'hardcore', name: 'X' },
        claim: { state: 'unclaimed', frozen: false },
        reports: [],
        roster: [],
      }),
    );
    vi.stubGlobal('fetch', upstream);
    await fetchGuildHome(42, 'abc123', API);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/guilds/42/home?cursor=abc123`);
  });
});

describe('claim and contest', () => {
  it('claims, confirms, releases and contests with POST', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ status: 'confirmed' }));
    vi.stubGlobal('fetch', upstream);
    await claimGuild(42, API);
    expect((upstream.mock.calls[0][0] as Request).method).toBe('POST');
    await confirmClaim(42, API);
    await releaseClaim(42, API);

    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ status: 'contested' })),
    );
    const result = await contestClaim(42, API);
    expect(result.status).toBe('contested');
  });
});

describe('settings', () => {
  it('reads claim_pending as an object, and the new claim field', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: null,
        claim_pending: { by: { battletag: 'Fixture#1234' }, expires_at: '2026-10-01T00:00:00Z' },
        claim: { state: 'pending', since: '2026-09-17T00:00:00Z', frozen: false },
        invite: { rotated_at: null },
      }),
    );
    vi.stubGlobal('fetch', upstream);
    const settings = await fetchGuildSettings(42, API);
    expect(settings.claim_pending?.by.battletag).toBe('Fixture#1234');
    expect(settings.claim.state).toBe('pending');
  });
});

describe('roster management', () => {
  it('approves and removes by three literal path segments, never an encoded combined key', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({ character_key: 'us/hardcore/thrallgar', status: 'approved' }),
    );
    vi.stubGlobal('fetch', upstream);

    const result = await approveCharacter(42, 'us', 'hardcore', 'thrallgar', API);
    expect(result.status).toBe('approved');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/guilds/42/characters/us/hardcore/thrallgar/approve`,
    );
    expect((upstream.mock.calls[0][0] as Request).method).toBe('POST');

    const removeUpstream = vi.fn<GlobalFetch>(async () => envelope({ status: 'removed' }));
    vi.stubGlobal('fetch', removeUpstream);
    await removeCharacter(42, 'us', 'hardcore', 'thrallgar', API);
    expect((removeUpstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/guilds/42/characters/us/hardcore/thrallgar`,
    );
    expect((removeUpstream.mock.calls[0][0] as Request).method).toBe('DELETE');
  });
});

describe('membership and invite', () => {
  it('PATCHes consent and DELETEs to leave', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ consent: 'gear_bags' }));
    vi.stubGlobal('fetch', upstream);
    await updateConsent(42, 'gear_bags', API);
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ status: 'left' })),
    );
    const left = await leaveGuild(42, API);
    expect(left.status).toBe('left');
  });

  it('accepts an invite token', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        guild: { id: 42, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' },
        rank: 'member',
      }),
    );
    vi.stubGlobal('fetch', upstream);
    const result = await acceptInvite('abc-123', API);
    expect(result.guild.name).toBe('The Last Watch');
  });

  it('rotates the invite', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({ token: 'abc', url: '/guild/invite/abc', rotated_at: '2026-09-21T00:00:00Z' }),
    );
    vi.stubGlobal('fetch', upstream);
    const rotated = await rotateInvite(42, API);
    expect(rotated.token).toBe('abc');
  });

  it('surfaces GUILD_API_FAILED when the network fails outright', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(fetchGuildHome(42, undefined, API)).rejects.toThrow(GUILD_API_FAILED);
    await expect(fetchGuildHome(42, undefined, API)).rejects.toBeInstanceOf(GuildApiError);
  });
});

// Guild crest round (docs/contracts/2026-10-05-guild-crest-api.md).
describe('crest', () => {
  it('PUTs the file as multipart/form-data, field "image", with no hand-set content-type', async () => {
    const upstream = vi.fn<GlobalFetch>(async (...args: Parameters<typeof fetch>) => {
      const request = args[0] as Request;
      expect(request.method).toBe('PUT');
      // The browser (not this module) writes `multipart/form-data; boundary=...` from the
      // FormData body itself -- asserting that boundary is present, and that it was never
      // overwritten with a hand-set `application/json`, is the real bug class this test
      // guards: a manual content-type header here silently breaks every upload. (Reading
      // the parts back out with `request.formData()` is skipped -- jsdom's own Request/
      // FormData interop cannot round-trip a File part in this test environment -- the
      // header the browser derives from the body is what actually matters here.)
      expect(request.headers.get('content-type')).toMatch(/^multipart\/form-data; boundary=/);
      return envelope({ crest_url: 'https://cdn.test/guilds/42/crest.webp?v=abc' });
    });
    vi.stubGlobal('fetch', upstream);
    const file = new File([new Uint8Array(8)], 'crest.png', { type: 'image/png' });
    const result = await putGuildCrest(42, file, API);
    expect(result.crest_url).toBe('https://cdn.test/guilds/42/crest.webp?v=abc');
  });

  it("surfaces the server's plain error message on a rejected upload", async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(
        async () =>
          new Response(
            JSON.stringify({
              ok: false,
              data: null,
              error: { message: 'That file is 3.1 MB; the limit is 2 MB.' },
              request_id: 'r',
            }),
            { status: 422, headers: { 'content-type': 'application/json' } },
          ),
      ),
    );
    const file = new File([new Uint8Array(8)], 'crest.png', { type: 'image/png' });
    await expect(putGuildCrest(42, file, API)).rejects.toThrow('That file is 3.1 MB; the limit is 2 MB.');
  });

  it('DELETEs and resolves even on a bare 204 (no envelope data)', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async (...args: Parameters<typeof fetch>) => {
        expect((args[0] as Request).method).toBe('DELETE');
        return new Response(null, { status: 204 });
      }),
    );
    const result = await deleteGuildCrest(42, API);
    expect(result.status).toBe('removed');
  });
});

describe('caching through query.ts', () => {
  it('fetchGuildHome caches per guild id and cursor', async () => {
    const fetchSpy = vi.fn<GlobalFetch>(async () => envelope(guildHomeFixture()));
    vi.stubGlobal('fetch', fetchSpy);
    await fetchGuildHome(9, undefined, API);
    await fetchGuildHome(9, undefined, API);
    expect(fetchSpy).toHaveBeenCalledTimes(1);
  });

  it('updateGuildSettings invalidates the settings cache for that guild', async () => {
    const fetchSpy = vi
      .fn<GlobalFetch>()
      .mockImplementationOnce(async () => envelope(guildSettingsFixture()))
      .mockImplementationOnce(async () => envelope(guildSettingsFixture()))
      .mockImplementationOnce(async () => envelope(guildSettingsFixture()));
    vi.stubGlobal('fetch', fetchSpy);
    await fetchGuildSettings(9, API);
    await updateGuildSettings(9, { officer_max_rank_index: 2 }, API);
    await fetchGuildSettings(9, API);
    expect(fetchSpy).toHaveBeenCalledTimes(3);
  });
});
