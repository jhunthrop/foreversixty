// web/src/lib/report/load.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import fixtureMeta from '../../fixtures/report/meta.json';
import fixtureReport from '../../fixtures/report/report.json';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import {
  POLL_INTERVAL_MS, REPORT_FORBIDDEN, REPORT_LOAD_FAILED, REPORT_NOT_FOUND, ReportLoadError,
  createPoller, eventsUrl, fetchAccessUrl, fetchLive, fetchReportFile, fetchReportMeta, fetchSummary,
} from './load';

const API = 'https://api.foreversixty.test';
// Absolute, per the Amendments: data_base_url is always
// `https://foreversixty.gg/logs-data/reports/<id>` (public/unlisted) or a signed absolute
// url from /access (private/guild) -- production never sends a relative one.
const DATA = 'https://logs.foreversixty.test/reports/fixture2abcd';

type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function json(body: unknown, status = 200): Response {
  return new Response(JSON.stringify(body), { status, headers: { 'content-type': 'application/json' } });
}

function envelope(data: unknown, status = 200): Response {
  return json({ ok: status < 400, data, error: null, request_id: 'req-1' }, status);
}

afterEach(() => {
  vi.unstubAllGlobals();
  vi.useRealTimers();
});

describe('fetchReportMeta', () => {
  it('unwraps the envelope and sends the session cookie', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(fixtureMeta));
    vi.stubGlobal('fetch', upstream);

    const meta = await fetchReportMeta('fixture2abcd', API);

    expect(meta.id).toBe('fixture2abcd');
    expect(meta.fights).toHaveLength(4);
    const request = upstream.mock.calls[0][0] as Request;
    expect(request.url).toBe(`${API}/v1/reports/fixture2abcd`);
    expect(request.credentials).toBe('include');
  });

  it('reports a missing report and a refused one differently', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(null, 404)));
    await expect(fetchReportMeta('nope', API)).rejects.toThrow(REPORT_NOT_FOUND);

    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => envelope(null, 403)));
    await expect(fetchReportMeta('secret', API)).rejects.toThrow(REPORT_FORBIDDEN);
  });

  it('turns a network failure into a ReportLoadError rather than a TypeError', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => { throw new TypeError('offline'); }));
    await expect(fetchReportMeta('fixture2abcd', API)).rejects.toBeInstanceOf(ReportLoadError);
    await expect(fetchReportMeta('fixture2abcd', API)).rejects.toThrow(REPORT_LOAD_FAILED);
  });
});

describe('fetchAccessUrl', () => {
  it('asks for the signed base url a private report needs', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      envelope({ data_base_url: 'https://signed.example/reports/fixture2abcd' }),
    );
    vi.stubGlobal('fetch', upstream);

    const url = await fetchAccessUrl('fixture2abcd', API);

    expect(url).toBe('https://signed.example/reports/fixture2abcd');
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${API}/v1/reports/fixture2abcd/access`);
  });
});

describe('the report files', () => {
  it('reads report.json from the data base url and normalises its arrays', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => json({ ...fixtureReport, units: null }));
    vi.stubGlobal('fetch', upstream);

    const file = await fetchReportFile(DATA);

    expect(file.fights.map((f) => f.index)).toEqual([1, 2, 3, 4]);
    expect(file.units).toEqual([]);
    expect((upstream.mock.calls[0][0] as Request).url).toContain(`${DATA}/report.json`);
  });

  it('busts the edge cache on report.json but not on a fight summary', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => json(fixtureReport));
    vi.stubGlobal('fetch', upstream);
    await fetchReportFile(DATA);
    expect((upstream.mock.calls[0][0] as Request).cache).toBe('no-cache');

    const summaries = vi.fn<GlobalFetch>(async () => json(fixtureSummary));
    vi.stubGlobal('fetch', summaries);
    await fetchSummary(DATA, 3);
    expect((summaries.mock.calls[0][0] as Request).cache).toBe('default');
  });

  it('reads one fight summary by index', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => json(fixtureSummary));
    vi.stubGlobal('fetch', upstream);

    const summary = await fetchSummary(DATA, 3);

    expect(summary.fight_index).toBe(3);
    expect(summary.deaths).toHaveLength(1);
    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${DATA}/fights/3/summary.json`);
  });

  it('treats a missing live.json as "the fight is not open" rather than an error', async () => {
    vi.stubGlobal('fetch', vi.fn<GlobalFetch>(async () => new Response('', { status: 404 })));
    await expect(fetchLive(DATA, 3)).resolves.toBeNull();
  });

  it('preserves a signed data_base_url query string byte-for-byte', async () => {
    // A signed url from GET /v1/reports/{id}/access carries a signature in its query
    // string; the request the browser sends must match it exactly, with no re-encoding.
    const signed = 'https://signed.example/reports/fixture2abcd?sig=AbC%2B123&exp=1700000000';
    const upstream = vi.fn<GlobalFetch>(async () => json(fixtureSummary));
    vi.stubGlobal('fetch', upstream);

    await fetchSummary(signed, 3);

    expect((upstream.mock.calls[0][0] as Request).url).toBe(`${signed}/fights/3/summary.json`);
  });

  it('names the events file without fetching it', () => {
    expect(eventsUrl(DATA, 2)).toBe(`${DATA}/fights/2/events.parquet`);
  });
});

describe('createPoller', () => {
  it('runs the tick every five seconds until it is stopped', async () => {
    vi.useFakeTimers();
    const tick = vi.fn(async () => {});
    const poller = createPoller(tick);

    poller.start();
    expect(POLL_INTERVAL_MS).toBe(5000);
    await vi.advanceTimersByTimeAsync(12_000);
    expect(tick).toHaveBeenCalledTimes(2);

    poller.stop();
    await vi.advanceTimersByTimeAsync(20_000);
    expect(tick).toHaveBeenCalledTimes(2);
  });

  it('never overlaps two ticks, and keeps going after one throws', async () => {
    vi.useFakeTimers();
    let running = 0;
    let peak = 0;
    let calls = 0;
    const poller = createPoller(async () => {
      calls += 1;
      running += 1;
      peak = Math.max(peak, running);
      if (calls === 1) {
        running -= 1;
        throw new Error('one bad poll');
      }
      await new Promise((resolve) => setTimeout(resolve, 9000));
      running -= 1;
    });

    poller.start();
    await vi.advanceTimersByTimeAsync(30_000);
    poller.stop();

    expect(peak).toBe(1);
    expect(calls).toBeGreaterThan(1);
  });
});
