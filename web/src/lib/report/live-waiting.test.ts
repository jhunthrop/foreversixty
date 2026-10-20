// @vitest-environment jsdom
// web/src/lib/report/live-waiting.test.ts
// The meta decides whether a report exists. A live report whose first fight has not been
// uploaded has a meta and no report.json, and that is "waiting", not "No report with that id".
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { resetQueryCache } from '../data/query';
import fixtureMeta from '../../fixtures/report/meta.json';
import {
  REPORT_DATA_MISSING,
  REPORT_NOT_FOUND,
  ReportLoadError,
  fetchReportFileFor,
  fetchReportMeta,
  fetchReportMetaAfterUpload,
} from './load';

const API = 'https://api.foreversixty.test';
const BASE = 'https://data.foreversixty.test/reports/xeduvzlm6cwn';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

const json = (body: unknown, status = 200): Response =>
  new Response(JSON.stringify(body), { status, headers: { 'content-type': 'application/json' } });
const envelope = (data: unknown, status = 200): Response =>
  json({ ok: status < 400, data, error: null, request_id: 'r' }, status);

beforeEach(() => {
  resetQueryCache();
});
afterEach(() => {
  resetQueryCache();
  vi.unstubAllGlobals();
});

describe('a report with a meta and no report.json', () => {
  it.each(['live', 'processing'] as const)(
    'is waiting, not missing, while the meta says %s',
    async (status) => {
      vi.stubGlobal(
        'fetch',
        vi.fn<GlobalFetch>(async () => json(null, 404)),
      );
      expect(await fetchReportFileFor(status, BASE)).toBeNull();
    },
  );

  it('is a data error, not a missing report, once the meta says complete', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => json(null, 404)),
    );
    const failure = await fetchReportFileFor('complete', BASE).catch((thrown: unknown) => thrown);
    expect(failure).toBeInstanceOf(ReportLoadError);
    expect((failure as ReportLoadError).message).toBe(REPORT_DATA_MISSING);
    expect((failure as ReportLoadError).message).not.toBe(REPORT_NOT_FOUND);
  });

  it('still fails for anything that is not a 404', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => json(null, 500)),
    );
    await expect(fetchReportFileFor('live', BASE)).rejects.toMatchObject({ status: 500 });
  });

  it('switches to the file as soon as a poll finds it', async () => {
    let calls = 0;
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        calls += 1;
        return calls < 3
          ? json(null, 404)
          : json({ engine_version: '0.1.0', fights: [{ index: 1 }], units: null });
      }),
    );
    expect(await fetchReportFileFor('live', BASE, 'fresh')).toBeNull();
    expect(await fetchReportFileFor('live', BASE, 'fresh')).toBeNull();
    const file = await fetchReportFileFor('live', BASE, 'fresh');
    expect(file?.fights).toHaveLength(1);
    expect(file?.units).toEqual([]);
  });
});

describe('the meta itself', () => {
  it('reads a live meta with no fights as a report that exists', async () => {
    const live = { ...fixtureMeta, status: 'live', fights: [], players: [], title: '', zone: '' };
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(live)),
    );
    const meta = await fetchReportMeta('fixture2abcd', API);
    expect(meta.status).toBe('live');
    expect(meta.fights).toEqual([]);
  });

  it('is missing only when the meta itself answers 404, after the upload grace', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 404)),
    );
    await expect(fetchReportMeta('fixture2abcd', API)).rejects.toMatchObject({ message: REPORT_NOT_FOUND });
    await expect(
      fetchReportMetaAfterUpload('fixture2abcd', {
        apiBase: API,
        uploadedAt: 0,
        now: () => 60_000,
        sleep: async () => {},
      }),
    ).rejects.toMatchObject({ message: REPORT_NOT_FOUND });
  });
});
