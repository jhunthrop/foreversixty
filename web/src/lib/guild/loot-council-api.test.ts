// @vitest-environment jsdom
// The loot council's transport (api/openapi.yaml loot/nights, loot/drops/{item}/candidates,
// loot/votes, loot/interest, loot/awards, export/loot.csv|json).
import { afterEach, describe, expect, it, vi } from 'vitest';
import { forgetPrivate, invalidate } from '../data/query';
import { GuildApiError } from './api';
import {
  awardDrop,
  fetchDropCandidates,
  fetchLootInterest,
  fetchLootNights,
  lootHistoryHref,
  putLootInterest,
  putLootVote,
} from './loot-council-api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  const error = status < 400 ? null : { code: 'raid_guild_required', message: 'No' };
  return new Response(JSON.stringify({ ok: status < 400, data, error, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

function stub(response: () => Response): ReturnType<typeof vi.fn<GlobalFetch>> {
  const upstream = vi.fn<GlobalFetch>(async () => response());
  vi.stubGlobal('fetch', upstream);
  return upstream;
}

const sent = (upstream: ReturnType<typeof vi.fn<GlobalFetch>>): Request =>
  upstream.mock.calls[0]![0] as Request;
const target = { night: 'raid:12', item_id: 17068, character_key: 'us/hardcore/kraggor' };

afterEach(() => {
  forgetPrivate();
  invalidate('');
  vi.unstubAllGlobals();
});

describe('loot council reads', () => {
  it('reads the nights, a drop of a night and the viewer flags', async () => {
    let upstream = stub(() => envelope({ officers: 4, nights: [] }));
    expect(await fetchLootNights(7, API)).toEqual({ officers: 4, nights: [] });
    expect(sent(upstream).url).toBe(`${API}/v1/guilds/7/loot/nights`);

    upstream = stub(() => envelope({ candidates: [] }));
    await fetchDropCandidates(7, 17068, 'raid:12', API);
    expect(sent(upstream).url).toBe(`${API}/v1/guilds/7/loot/drops/17068/candidates?night=raid%3A12`);

    upstream = stub(() => envelope({ night: 'live', flags: [] }));
    await fetchLootInterest(7, 'live', API);
    expect(sent(upstream).url).toBe(`${API}/v1/guilds/7/loot/interest?night=live`);
  });
});

describe('loot council writes', () => {
  it('puts a vote and answers the tally', async () => {
    const votes = { yes: 1, no: 0, voters: 4, mine: 'yes' };
    const upstream = stub(() => envelope({ ...target, votes }));
    expect(await putLootVote(7, target, 'yes', API)).toEqual(votes);
    expect(sent(upstream).method).toBe('PUT');
    expect(await sent(upstream).clone().json()).toEqual({ ...target, vote: 'yes' });
  });

  it('puts a flag, null removing it', async () => {
    const upstream = stub(() => envelope({ ...target, ms_os: null }));
    expect(await putLootInterest(7, target, null, API)).toBeNull();
    expect(sent(upstream).url).toBe(`${API}/v1/guilds/7/loot/interest`);
    expect(await sent(upstream).clone().json()).toEqual({ ...target, ms_os: null });
  });

  it('posts an award with its night, MS/OS and reason', async () => {
    const upstream = stub(() => envelope({ award_id: 9, status: 'awarded' }));
    await awardDrop(7, { ...target, encounter_id: 1084, ms_os: 'os', reason: 'Main spec passed' }, API);
    expect(sent(upstream).method).toBe('POST');
    expect(await sent(upstream).clone().json()).toEqual({
      ...target,
      encounter_id: 1084,
      ms_os: 'os',
      reason: 'Main spec passed',
    });
  });

  it('rethrows a plan refusal with its code', async () => {
    stub(() => envelope(null, 403));
    const error = await putLootVote(7, target, 'no', API).catch((caught: unknown) => caught);
    expect(error).toBeInstanceOf(GuildApiError);
    expect((error as GuildApiError).code).toBe('raid_guild_required');
  });
});

describe('lootHistoryHref', () => {
  it('names the format and only the dates given', () => {
    expect(lootHistoryHref(7, 'csv', { from: '', to: '' }, API)).toBe(`${API}/v1/guilds/7/export/loot.csv`);
    expect(lootHistoryHref(7, 'json', { from: '2026-12-01', to: '' }, API)).toBe(
      `${API}/v1/guilds/7/export/loot.json?from=2026-12-01`,
    );
    expect(lootHistoryHref(7, 'csv', { from: '2026-12-01', to: '2026-12-31' }, API)).toBe(
      `${API}/v1/guilds/7/export/loot.csv?from=2026-12-01&to=2026-12-31`,
    );
  });
});
