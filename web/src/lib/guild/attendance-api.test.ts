// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { forgetPrivate, invalidate } from '../data/query';
import { GuildApiError } from './api';
import {
  clearMatch,
  fetchGuildReadiness,
  fetchMatchCandidates,
  postReport,
  setExcused,
  setMatch,
} from './attendance-api';

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
    { status, headers: { 'content-type': 'application/json' } },
  );
}

function stub(response: () => Response): ReturnType<typeof vi.fn<GlobalFetch>> {
  const upstream = vi.fn<GlobalFetch>(async () => response());
  vi.stubGlobal('fetch', upstream);
  return upstream;
}

function sent(upstream: ReturnType<typeof vi.fn<GlobalFetch>>, call = 0): Request {
  return upstream.mock.calls[call]![0] as Request;
}

afterEach(() => {
  forgetPrivate();
  invalidate('');
  vi.unstubAllGlobals();
});

describe('fetchMatchCandidates', () => {
  it('reads the picker list under the raid', async () => {
    const upstream = stub(() => envelope({ zone_known: true, rows: [] }));
    const result = await fetchMatchCandidates(7, 12, API);
    expect(result).toEqual({ zone_known: true, rows: [] });
    expect(sent(upstream).url).toBe(`${API}/v1/guilds/7/raids/12/match-candidates`);
  });
});

describe('match writes', () => {
  it('sets a match with the report id', async () => {
    const upstream = stub(() => envelope({ id: 12 }));
    await setMatch(7, 12, 'abc', API);
    const request = sent(upstream);
    expect(request.method).toBe('PUT');
    expect(request.url).toBe(`${API}/v1/guilds/7/raids/12/match`);
    expect(await request.clone().json()).toEqual({ report_id: 'abc' });
  });

  it('clears a match with a DELETE', async () => {
    const upstream = stub(() => envelope({ id: 12 }));
    await clearMatch(7, 12, API);
    expect(sent(upstream).method).toBe('DELETE');
    expect(sent(upstream).url).toBe(`${API}/v1/guilds/7/raids/12/match`);
  });

  it('rethrows report_matched with its code', async () => {
    stub(() => errorEnvelope(409, 'report_matched'));
    const error = await setMatch(7, 12, 'abc', API).catch((caught: unknown) => caught);
    expect(error).toBeInstanceOf(GuildApiError);
    expect((error as GuildApiError).code).toBe('report_matched');
  });
});

describe('setExcused', () => {
  it('puts the flag at the character key, one segment each', async () => {
    const upstream = stub(() => envelope({ id: 12 }));
    await setExcused(7, 12, 'us/classic/Pyre wisp', true, API);
    const request = sent(upstream);
    expect(request.method).toBe('PUT');
    expect(request.url).toBe(`${API}/v1/guilds/7/raids/12/attendance/us/classic/Pyre%20wisp`);
    expect(await request.clone().json()).toEqual({ excused: true });
  });
});

describe('postReport', () => {
  it('posts and answers posted_at', async () => {
    const upstream = stub(() => envelope({ posted_at: '2026-12-11T05:58:00Z' }));
    const result = await postReport(7, 12, API);
    expect(result.posted_at).toBe('2026-12-11T05:58:00Z');
    expect(sent(upstream).url).toBe(`${API}/v1/guilds/7/raids/12/report/post`);
    expect(sent(upstream).method).toBe('POST');
  });

  it.each([
    [409, 'already_posted'],
    [409, 'conflict'],
    [502, 'post_failed'],
  ])('rethrows %i %s', async (status, code) => {
    stub(() => errorEnvelope(status, code));
    const error = await postReport(7, 12, API).catch((caught: unknown) => caught);
    expect((error as GuildApiError).code).toBe(code);
    expect((error as GuildApiError).status).toBe(status);
  });
});

describe('fetchGuildReadiness', () => {
  const page = { median_item_level: 0, generated_at: '', rows: [] };

  it('asks for the raid filter only when given a raid', async () => {
    const upstream = stub(() => envelope(page));
    await fetchGuildReadiness(7, 12, API);
    await fetchGuildReadiness(7, undefined, API);
    expect(sent(upstream, 0).url).toBe(`${API}/v1/guilds/7/readiness?raid=12`);
    expect(sent(upstream, 1).url).toBe(`${API}/v1/guilds/7/readiness`);
  });

  it('is cleared by a write to the guild', async () => {
    const upstream = stub(() => envelope(page));
    await fetchGuildReadiness(7, 12, API);
    await clearMatch(7, 12, API);
    upstream.mockClear();
    await fetchGuildReadiness(7, 12, API);
    expect(upstream).toHaveBeenCalledTimes(1);
  });

  it('rethrows a 404 for a closed raid', async () => {
    stub(() => errorEnvelope(404, 'not_found'));
    const error = await fetchGuildReadiness(7, 99, API).catch((caught: unknown) => caught);
    expect((error as GuildApiError).code).toBe('not_found');
  });
});
