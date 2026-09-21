// web/src/lib/rankings/api.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  RANKINGS_FAILED,
  RankingsError,
  encounterSlug,
  fetchCharacter,
  fetchCharacterRating,
  fetchEncounters,
  fetchGuild,
  fetchRankings,
  fetchReportRatings,
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
        guild: { id: 501, name: 'The Last Watch', region: 'eu', ruleset: 'normal' },
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

describe('fetchEncounters', () => {
  it('reads the encounter list the /rankings picker offers', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({ rows: [{ id: 9001, name: 'Warden Kelthas', slug: 'warden-kelthas' }] }),
    );
    vi.stubGlobal('fetch', upstream);

    const { rows } = await fetchEncounters(API);

    expect(rows).toEqual([{ id: 9001, name: 'Warden Kelthas', slug: 'warden-kelthas' }]);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/encounters`);
  });

  it('turns a failure into a RankingsError, like every other read here', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(fetchEncounters(API)).rejects.toBeInstanceOf(RankingsError);
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

const REPORT_RATINGS = {
  fight_index: 3,
  kill: true,
  kill_time_band: 'typical',
  model_version: 'rating-2026-09-21',
  players: [
    {
      player_key: 'us/normal/simfury',
      player_name: 'Simfury',
      class: 'Warrior',
      spec: 'Fury',
      role: 'dps',
      overall: 70,
      overall_uncapped: 70,
      overall_capped: false,
      basis: 'percentile',
      components: [],
    },
  ],
};

describe('fetchReportRatings', () => {
  it('reads the per-fight report card from the ratings endpoint', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(REPORT_RATINGS));
    vi.stubGlobal('fetch', upstream);

    const ratings = await fetchReportRatings('fixture2abcd', 3, API);

    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/reports/fixture2abcd/fights/3/ratings`,
    );
    expect((upstream.mock.calls[0][0] as Request).credentials).toBe('omit');
    expect(ratings.players[0].player_name).toBe('Simfury');
  });

  it('an empty roster (Ruling 1’s launch state) resolves normally, not as an error', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ ...REPORT_RATINGS, players: [] })),
    );
    const ratings = await fetchReportRatings('fixture2abcd', 3, API);
    expect(ratings.players).toEqual([]);
  });

  it('a transport failure raises RankingsError', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 500)),
    );
    await expect(fetchReportRatings('fixture2abcd', 3, API)).rejects.toBeInstanceOf(RankingsError);
  });
});

describe('fetchCharacterRating', () => {
  it('reads the aggregate from the character rating endpoint', async () => {
    const data = {
      player_key: 'us/hardcore/elyra-duskvale',
      sample_size: 12,
      trend: [{ fought_at: '2026-12-09T22:10:00Z', overall: 62, report_id: 'fixture2abcd', fight_index: 2 }],
      best_component: 'preparation',
      worst_component: 'activity',
      latest: null,
    };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(data));
    vi.stubGlobal('fetch', upstream);

    const rating = await fetchCharacterRating(
      { region: 'us', ruleset: 'hardcore', slug: 'elyra-duskvale' },
      API,
    );

    expect((upstream.mock.calls[0][0] as Request).url).toBe(
      `${API}/v1/characters/us/hardcore/elyra-duskvale/rating`,
    );
    expect((upstream.mock.calls[0][0] as Request).credentials).toBe('omit');
    expect(rating.sample_size).toBe(12);
  });

  it('an anonymized character 404s, and the fetcher raises RankingsError (Ruling 2)', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 404)),
    );
    await expect(
      fetchCharacterRating({ region: 'us', ruleset: 'hardcore', slug: 'hidden' }, API),
    ).rejects.toBeInstanceOf(RankingsError);
  });
});
