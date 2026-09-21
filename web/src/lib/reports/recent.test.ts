// web/src/lib/reports/recent.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { RECENT_REPORTS_FAILED, RecentReportsError, fetchRecentReports } from './recent';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

afterEach(() => vi.unstubAllGlobals());

describe('fetchRecentReports', () => {
  it('reads the first page with no cursor in the query', async () => {
    const page = {
      rows: [
        {
          id: 'abc123',
          title: 'Progress night',
          created_at: '2026-12-09T22:10:00Z',
          fight_count: 8,
          kill_count: 3,
        },
      ],
    };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(page));
    vi.stubGlobal('fetch', upstream);

    const result = await fetchRecentReports(undefined, API);

    expect(result.rows[0].title).toBe('Progress night');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/reports/recent`);
  });

  it('passes a given cursor back as the query parameter, url-encoded', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ rows: [] }));
    vi.stubGlobal('fetch', upstream);

    await fetchRecentReports('a/b+c', API);

    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/reports/recent?cursor=a%2Fb%2Bc`);
  });

  it('sends no credentials: the feed answers the same for every visitor', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ rows: [] }));
    vi.stubGlobal('fetch', upstream);

    await fetchRecentReports(undefined, API);

    expect((upstream.mock.calls[0][0] as Request).credentials).toBe('omit');
  });

  it('turns a failed response into a RecentReportsError', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 500)),
    );
    await expect(fetchRecentReports(undefined, API)).rejects.toBeInstanceOf(RecentReportsError);
  });

  it('turns a network failure into a RecentReportsError with the fallback message', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(fetchRecentReports(undefined, API)).rejects.toThrow(RECENT_REPORTS_FAILED);
  });

  it('treats an envelope with no data as an empty page rather than throwing', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(null, 200));
    vi.stubGlobal('fetch', upstream);

    const result = await fetchRecentReports(undefined, API);
    expect(result.rows).toEqual([]);
  });
});
