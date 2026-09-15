// web/src/lib/rankings/api.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  RANKINGS_FAILED,
  RankingsError,
  encounterSlug,
  fetchCharacter,
  fetchGuild,
  fetchRankings,
} from './api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

const PAGE = {
  rows: [
    {
      rank: 1,
      player: { key: 'us/hardcore/elyra-duskvale', name: 'Elyra Duskvale', class: 'Priest', spec: 'Shadow' },
      guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
      value: 1840.2,
      size: 40,
      fought_at: '2026-12-09T22:10:00Z',
      duration_ms: 240_000,
      talent_split: '31/20/0',
      build_id: 'k7x2qm4a',
      trinkets: [19339, 18820],
      buff_count: 11,
      report_id: 'fixture2abcd',
      fight_index: 3,
      state: 'ok',
    },
  ],
  total: 812,
  page: 1,
  per_page: 100,
  updated_at: '2026-12-09T22:15:00Z',
};

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

afterEach(() => vi.unstubAllGlobals());

describe('fetchRankings', () => {
  it('sends only the filters that are set, in the contract’s parameter names', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(PAGE));
    vi.stubGlobal('fetch', upstream);

    const page = await fetchRankings(
      { encounter: 'warden-kelthas', metric: 'dps', ruleset: 'hardcore', spec: '', page: 2 },
      API,
    );

    expect(page.rows[0].player.name).toBe('Elyra Duskvale');
    expect(page.rows[0].guild?.ruleset).toBe('hardcore');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/rankings?encounter=warden-kelthas&metric=dps&ruleset=hardcore&page=2`,
    );
  });

  it('turns a failure into a RankingsError a page can print', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(fetchRankings({ encounter: 'x' }, API)).rejects.toBeInstanceOf(RankingsError);
    await expect(fetchRankings({ encounter: 'x' }, API)).rejects.toThrow(RANKINGS_FAILED);
  });
});

describe('character and guild', () => {
  it('reads a character by region, ruleset and slug', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore' },
        best: [],
        history: [],
        builds_seen: [],
      }),
    );
    vi.stubGlobal('fetch', upstream);

    await fetchCharacter({ region: 'us', ruleset: 'hardcore', slug: 'elyra-duskvale' }, API);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/characters/us/hardcore/elyra-duskvale`,
    );
  });

  it('reads a guild the same way', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({
        guild: { name: 'The Last Watch', region: 'eu', ruleset: 'normal' },
        progression: [],
        roster_best: [],
        reports: [],
      }),
    );
    vi.stubGlobal('fetch', upstream);

    await fetchGuild({ region: 'eu', ruleset: 'normal', slug: 'the-last-watch' }, API);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/guilds/eu/normal/the-last-watch`);
  });
});

describe('encounterSlug', () => {
  it('lowercases and hyphenates a plain multi-word name', () => {
    expect(encounterSlug('Warden Kelthas')).toBe('warden-kelthas');
  });

  it('collapses a run of punctuation to one hyphen rather than one per character', () => {
    expect(encounterSlug("Kel'Thas Sunstrider")).toBe('kel-thas-sunstrider');
    expect(encounterSlug('A -- B')).toBe('a-b');
  });

  it('trims a leading or trailing hyphen rather than keeping it', () => {
    expect(encounterSlug('  Leading and Trailing  ')).toBe('leading-and-trailing');
  });

  it('leaves a single word alone but for the case', () => {
    expect(encounterSlug('Trash')).toBe('trash');
  });

  it('is total: an empty name slugs to an empty string, not a throw', () => {
    expect(encounterSlug('')).toBe('');
  });
});
