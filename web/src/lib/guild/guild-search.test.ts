import { afterEach, describe, expect, it, vi } from 'vitest';
import { searchCandidates, searchGuilds } from './guild-search';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

const ROW = { name: 'Lifecycle Test', region: 'us', ruleset: 'pvp', path: '/guild/us/pvp/lifecycle-test' };

afterEach(() => vi.unstubAllGlobals());

describe('searchCandidates', () => {
  it('reads rows, or a bare array, keeping only rows with the three names', () => {
    const expected = [{ name: 'Lifecycle Test', region: 'us', ruleset: 'pvp' }];
    expect(searchCandidates({ rows: [ROW, { name: 'No place' }, null] })).toEqual(expected);
    expect(searchCandidates([ROW])).toEqual(expected);
    expect(searchCandidates(null)).toEqual([]);
    expect(searchCandidates({})).toEqual([]);
  });
});

describe('searchGuilds', () => {
  it('asks the public search with the typed words and region', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ rows: [ROW] }));
    vi.stubGlobal('fetch', upstream);
    const rows = await searchGuilds('lifecyle tset', 'us', API);
    expect(rows.map((row) => row.name)).toEqual(['Lifecycle Test']);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/guilds/search?q=lifecyle+tset&region=us`,
    );
  });

  it('throws on a failed read, so the page can stay silent', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 500)),
    );
    await expect(searchGuilds('lifecycle', 'us', API)).rejects.toThrow();
  });
});
