// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { coachModelReview, coachTeaser } from '../../test-support/coach-fixtures';
import { CoachApiError, fetchCoachHistory, fetchCoachRecent, fetchFightCoach, fetchNightLine } from './api';

const BASE = 'https://api.test';
let requested: string[] = [];

function answer(status: number, data: unknown, code?: string): void {
  vi.stubGlobal(
    'fetch',
    vi.fn((request: Request) => {
      const url = new URL(request.url);
      requested.push(url.pathname + url.search);
      const body =
        status < 400
          ? { ok: true, data, error: null }
          : { ok: false, data: null, error: { code, message: 'x' } };
      return Promise.resolve(
        new Response(JSON.stringify(body), { status, headers: { 'content-type': 'application/json' } }),
      );
    }),
  );
}

beforeEach(() => {
  requested = [];
});
afterEach(() => vi.unstubAllGlobals());

describe('fetchFightCoach', () => {
  it('sends the player key and returns null on 404', async () => {
    answer(404, null, 'not_found');
    expect(await fetchFightCoach('r 1', 3, 'us-x-Thrall', BASE)).toBeNull();
    expect(requested).toEqual(['/v1/reports/r%201/fights/3/coach?player=us-x-Thrall']);
  });

  it('the row strip form adds line=1', async () => {
    answer(200, coachTeaser());
    await fetchFightCoach('r1', 3, 'k', BASE, { line: true });
    expect(requested[0]).toBe('/v1/reports/r1/fights/3/coach?player=k&line=1');
  });

  it('a teaser parses as a teaser and drops unknown fields', async () => {
    answer(200, { ...coachTeaser(), leak: 'x' });
    const got = await fetchFightCoach('r1', 3, 'k', BASE);
    expect(got?.access).toBe('teaser');
    expect(got).toEqual(coachTeaser());
  });

  it('a full review round-trips, and an officer answer keeps its access', async () => {
    answer(200, coachModelReview({ access: 'officer' }));
    expect(await fetchFightCoach('r1', 3, 'k', BASE)).toEqual(coachModelReview({ access: 'officer' }));
  });

  it('an older shape with missing fields never throws', async () => {
    answer(200, {
      access: 'full',
      status: 'strange',
      review: { fixFirst: 4 },
      anchors: { a1: { kind: 'nope' } },
    });
    expect(await fetchFightCoach('r1', 3, 'k', BASE)).toMatchObject({
      status: 'ready',
      source: 'rules',
      review: { fixFirst: null, thenFix: [] },
      anchors: {},
    });
  });

  it('a failure throws CoachApiError with status and code', async () => {
    answer(500, null, 'internal');
    await expect(fetchFightCoach('r1', 3, 'k', BASE)).rejects.toBeInstanceOf(CoachApiError);
    await expect(fetchFightCoach('r1', 3, 'k', BASE)).rejects.toMatchObject({
      status: 500,
      code: 'internal',
    });
  });
});

describe('the other routes', () => {
  it('recent normalises the last raid and history items', async () => {
    answer(200, {
      last_raid: { zone: 'Molten Core', first: { stem: 's', on: 5, of: 8 } },
      pending: true,
      reviews: [{ boss: 'Ragnaros', kill: false }, 7],
    });
    const got = await fetchCoachRecent(BASE);
    expect(got.pending).toBe(true);
    expect(got.last_raid?.first).toEqual({ stem: 's', on: 5, of: 8 });
    expect(got.reviews).toHaveLength(1);
    expect(got.reviews[0]).toMatchObject({ boss: 'Ragnaros', kill: false });
  });

  it('history encodes the path and passes the cursor', async () => {
    answer(200, { reviews: [], next_cursor: 'c2' });
    const page = await fetchCoachHistory('us', 'forever', 'Thrall Jr', 'c1', BASE);
    expect(requested[0]).toBe('/v1/characters/us/forever/Thrall%20Jr/coach/history?cursor=c1');
    expect(page.next_cursor).toBe('c2');
  });

  it('the night line is null when the API sends none', async () => {
    answer(200, {});
    expect(await fetchNightLine('r1', 'k', BASE)).toBeNull();
    answer(200, { stem: 'Bloodthirst casts lost', on: 2, of: 4, finding_count: 3, fight_index: 1 });
    expect(await fetchNightLine('r1', 'k', BASE)).toMatchObject({ on: 2, of: 4 });
  });
});
