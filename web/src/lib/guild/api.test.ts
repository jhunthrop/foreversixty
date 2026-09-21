// web/src/lib/guild/api.test.ts
// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  GUILD_API_FAILED,
  GuildApiError,
  acceptInvite,
  approveCharacter,
  claimGuild,
  confirmClaim,
  fetchGuildHome,
  fetchGuildSettings,
  leaveGuild,
  releaseClaim,
  removeCharacter,
  rotateInvite,
  updateConsent,
  updateGuildSettings,
} from './api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

afterEach(() => vi.unstubAllGlobals());

describe('fetchGuildHome', () => {
  it('reads the guild home by numeric id', async () => {
    const home = {
      guild: { id: 42, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' },
      viewer: { rank: 'member', verified: true, can_manage: false },
      reports: { rows: [] },
      roster: [],
    };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(home));
    vi.stubGlobal('fetch', upstream);

    const result = await fetchGuildHome(42, API);

    expect(result.guild.name).toBe('The Last Watch');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/guilds/42/home`);
  });

  it('turns a failure into a GuildApiError', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 403)),
    );
    await expect(fetchGuildHome(42, API)).rejects.toBeInstanceOf(GuildApiError);
  });
});

describe('claim flow', () => {
  it('claims, confirms and releases with POST', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ status: 'confirmed' }));
    vi.stubGlobal('fetch', upstream);

    await claimGuild(42, API);
    expect((upstream.mock.calls[0][0] as Request).method).toBe('POST');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/guilds/42/claim`);

    await confirmClaim(42, API);
    expect((upstream.mock.calls[1][0] as Request).url).toBe(`${API}/v1/guilds/42/claim/confirm`);

    await releaseClaim(42, API);
    expect((upstream.mock.calls[2][0] as Request).url).toBe(`${API}/v1/guilds/42/claim/release`);
  });
});

describe('settings', () => {
  it('PATCHes only the fields given', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 2,
        claimed_by: null,
        claim_pending: false,
        invite: { rotated_at: null },
      }),
    );
    vi.stubGlobal('fetch', upstream);

    await updateGuildSettings(42, { officer_max_rank_index: 2 }, API);

    const request = upstream.mock.calls[0][0] as Request;
    expect(request.method).toBe('PATCH');
    expect(await request.clone().json()).toEqual({ officer_max_rank_index: 2 });
  });

  it('reads settings and rotates the invite', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { battletag: 'Fixture#1234' },
        claim_pending: false,
        invite: { rotated_at: '2026-09-01T00:00:00Z' },
      }),
    );
    vi.stubGlobal('fetch', upstream);
    const settings = await fetchGuildSettings(42, API);
    expect(settings.claimed_by?.battletag).toBe('Fixture#1234');

    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () =>
        envelope({
          token: 'abc',
          url: 'https://foreversixty.gg/guild/invite/abc',
          rotated_at: '2026-09-21T00:00:00Z',
        }),
      ),
    );
    const rotated = await rotateInvite(42, API);
    expect(rotated.token).toBe('abc');
  });
});

describe('roster management', () => {
  it('percent-encodes a character_key containing slashes for approve and remove', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        character_key: 'us/hardcore/thrallgar',
        region: 'us',
        ruleset: 'hardcore',
        name: 'Thrallgar',
        rank: 'member',
        verified: true,
        logged_recently: false,
        consent: 'gear',
      }),
    );
    vi.stubGlobal('fetch', upstream);

    await approveCharacter(42, 'us/hardcore/thrallgar', API);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/guilds/42/characters/us%2Fhardcore%2Fthrallgar/approve`,
    );
    expect((upstream.mock.calls[0][0] as Request).method).toBe('POST');

    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ status: 'removed' })),
    );
    await removeCharacter(42, 'us/hardcore/thrallgar', API);
  });
});

describe('membership', () => {
  it('PATCHes consent and DELETEs to leave', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ consent: 'gear_bags' }));
    vi.stubGlobal('fetch', upstream);
    const updated = await updateConsent(42, 'gear_bags', API);
    expect(updated.consent).toBe('gear_bags');
    const request = upstream.mock.calls[0][0] as Request;
    expect(request.method).toBe('PATCH');
    expect(await request.clone().json()).toEqual({ consent: 'gear_bags' });

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
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/guilds/invite/abc-123/accept`);
  });

  it('surfaces GUILD_API_FAILED when the network fails outright', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(fetchGuildHome(42, API)).rejects.toThrow(GUILD_API_FAILED);
  });
});
