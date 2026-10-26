// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { forgetPrivate, invalidate } from '../data/query';
import { NO_LOG_YET } from '../../test-support/raid-fixtures';
import { GuildApiError, type GuildRaidView } from './api';
import {
  createRaid,
  fetchGuildRaid,
  putSignup,
  setRaidStatus,
  updateRaid,
  type RaidInput,
} from './raids-api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

function errorEnvelope(status: number, code: string): Response {
  return new Response(
    JSON.stringify({ ok: false, data: null, error: { code, message: 'No' }, request_id: 'r' }),
    {
      status,
      headers: { 'content-type': 'application/json' },
    },
  );
}

function raid(id: number): GuildRaidView {
  return {
    id,
    zone: 'onyxia',
    zone_name: "Onyxia's Lair",
    title: "Onyxia's Lair",
    starts_at: '2026-12-11T02:00:00Z',
    duration_minutes: 180,
    status: 'planned',
    size: 40,
    guild_timezone: 'America/Chicago',
    notes_html: '',
    counts: { in: 0, late: 0, bench: 0, out: 0, tanks: 0, healers: 0, dps: 0 },
    signups: [],
    not_signed: [],
    viewer: { characters: [] },
    discord: { available: true, connected: false, post_default: false },
    ...NO_LOG_YET,
  };
}

const input: RaidInput = {
  zone: 'onyxia',
  title: '',
  starts_at: '2026-12-11T02:00:00Z',
  duration_minutes: 180,
  notes: '',
  timezone: 'America/Chicago',
  post_to_discord: false,
};

function sentRequest(upstream: ReturnType<typeof vi.fn<GlobalFetch>>, call = 0): Request {
  return upstream.mock.calls[call][0] as Request;
}

afterEach(() => {
  forgetPrivate();
  invalidate('');
  vi.unstubAllGlobals();
});

describe('fetchGuildRaid', () => {
  it('reads one raid under the guild path', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(raid(12)));
    vi.stubGlobal('fetch', upstream);
    const result = await fetchGuildRaid(7, 12, API);
    expect(result.id).toBe(12);
    expect(sentRequest(upstream).url).toBe(`${API}/v1/guilds/7/raids/12`);
  });

  it('is cleared by a write to the same guild', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(raid(12)));
    vi.stubGlobal('fetch', upstream);
    await fetchGuildRaid(7, 12, API);
    await setRaidStatus(7, 12, 'lock', API);
    upstream.mockClear();
    await fetchGuildRaid(7, 12, API);
    expect(upstream).toHaveBeenCalledTimes(1);
  });

  it('rethrows a refusal with its code', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => errorEnvelope(404, 'not_found')),
    );
    const error = await fetchGuildRaid(7, 99, API).catch((caught: unknown) => caught);
    expect(error).toBeInstanceOf(GuildApiError);
    expect((error as GuildApiError).code).toBe('not_found');
  });
});

describe('officer writes', () => {
  it('creates with POST and answers every raid made', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ raids: [raid(1), raid(2)] }, 201));
    vi.stubGlobal('fetch', upstream);
    const made = await createRaid(7, { ...input, repeat_weeks: 3 }, API);
    expect(made.map((r) => r.id)).toEqual([1, 2]);
    const request = sentRequest(upstream);
    expect(request.method).toBe('POST');
    expect(request.url).toBe(`${API}/v1/guilds/7/raids`);
    expect(await request.json()).toMatchObject({ zone: 'onyxia', repeat_weeks: 3 });
  });

  it('updates with PATCH', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(raid(12)));
    vi.stubGlobal('fetch', upstream);
    await updateRaid(7, 12, input, API);
    expect(sentRequest(upstream).method).toBe('PATCH');
    expect(sentRequest(upstream).url).toBe(`${API}/v1/guilds/7/raids/12`);
  });

  it.each(['lock', 'unlock', 'cancel'] as const)('%s posts to its own path', async (action) => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(raid(12)));
    vi.stubGlobal('fetch', upstream);
    await setRaidStatus(7, 12, action, API);
    expect(sentRequest(upstream).method).toBe('POST');
    expect(sentRequest(upstream).url).toBe(`${API}/v1/guilds/7/raids/12/${action}`);
  });

  it('rethrows guild_plan_required and raid_cap with their codes', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => errorEnvelope(403, 'guild_plan_required')),
    );
    const refused = await createRaid(7, input, API).catch((caught: unknown) => caught);
    expect((refused as GuildApiError).code).toBe('guild_plan_required');

    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => errorEnvelope(422, 'raid_cap')),
    );
    const capped = await createRaid(7, input, API).catch((caught: unknown) => caught);
    expect((capped as GuildApiError).code).toBe('raid_cap');
  });
});

describe('putSignup', () => {
  it('puts the state, role and note, and answers the raid and who moved', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ raid: raid(12), moved: 'Lightbrand' }));
    vi.stubGlobal('fetch', upstream);
    const result = await putSignup(
      7,
      12,
      'us/pvp/hunthrop',
      { state: 'in', role: 'dps', note: 'late 5' },
      API,
    );
    expect(result.moved).toBe('Lightbrand');
    expect(result.raid.id).toBe(12);
    const request = sentRequest(upstream);
    expect(request.method).toBe('PUT');
    expect(request.url).toBe(`${API}/v1/guilds/7/raids/12/signups/us/pvp/hunthrop`);
    expect(await request.json()).toEqual({ state: 'in', role: 'dps', note: 'late 5' });
  });

  it('rethrows the lock and cancel races with their codes', async () => {
    for (const code of ['raid_locked', 'raid_cancelled']) {
      vi.stubGlobal(
        'fetch',
        vi.fn<GlobalFetch>(async () => errorEnvelope(409, code)),
      );
      const error = await putSignup(7, 12, 'k', { state: 'in', role: 'dps', note: '' }, API).catch(
        (caught: unknown) => caught,
      );
      expect(error).toBeInstanceOf(GuildApiError);
      expect((error as GuildApiError).code).toBe(code);
      expect((error as GuildApiError).status).toBe(409);
    }
  });
});
