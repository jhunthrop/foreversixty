// @vitest-environment jsdom
// web/src/lib/report/processing.test.ts
// A report being parsed is never remembered; a finished one is. And the 404 right after an
// upload is "not visible yet" for a while, then "No report with that id".
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { resetQueryCache, storageKey } from '../data/query';
import fixtureMeta from '../../fixtures/report/meta.json';
import { justUploadedAt, markJustUploaded } from '../reports/just-uploaded';
import { REPORT_NOT_FOUND, fetchReportMeta, fetchReportMetaAfterUpload } from './load';

const API = 'https://api.foreversixty.test';
const KEY = `${API}/v1/reports/fixture2abcd`;
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

const meta = (status: string) => ({ ...fixtureMeta, status });

beforeEach(() => {
  resetQueryCache();
});
afterEach(() => {
  resetQueryCache();
  vi.unstubAllGlobals();
});

describe('the report cache rule', () => {
  it('never persists or reuses a processing answer', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(meta('processing')));
    vi.stubGlobal('fetch', upstream);

    expect((await fetchReportMeta('fixture2abcd', API)).status).toBe('processing');
    expect(localStorage.getItem(storageKey(KEY))).toBeNull();

    await fetchReportMeta('fixture2abcd', API);
    expect(upstream).toHaveBeenCalledTimes(2); // the second read went to the network
  });

  it('never persists a live answer either', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(meta('live'))),
    );
    await fetchReportMeta('fixture2abcd', API);
    expect(localStorage.getItem(storageKey(KEY))).toBeNull();
  });

  it('persists a complete answer and serves it without asking again', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(meta('complete')));
    vi.stubGlobal('fetch', upstream);

    await fetchReportMeta('fixture2abcd', API);
    expect(localStorage.getItem(storageKey(KEY))).not.toBeNull();
    await fetchReportMeta('fixture2abcd', API);
    expect(upstream).toHaveBeenCalledTimes(1);
  });

  it('persists a failed answer', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(meta('failed'))),
    );
    await fetchReportMeta('fixture2abcd', API);
    expect(localStorage.getItem(storageKey(KEY))).not.toBeNull();
  });

  it('drops a processing snapshot an older build left in storage, and reads the network', async () => {
    localStorage.setItem(
      storageKey(KEY),
      JSON.stringify({ key: KEY, scope: 'public', v: 0, savedAt: Date.now(), data: meta('processing') }),
    );
    const upstream = vi.fn<GlobalFetch>(async () => envelope(meta('complete')));
    vi.stubGlobal('fetch', upstream);

    expect((await fetchReportMeta('fixture2abcd', API)).status).toBe('complete');
    expect(upstream).toHaveBeenCalledTimes(1);
  });

  it('moves from processing to complete and then caches', async () => {
    let status = 'processing';
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(meta(status))),
    );
    expect((await fetchReportMeta('fixture2abcd', API)).status).toBe('processing');
    status = 'complete';
    expect((await fetchReportMeta('fixture2abcd', API)).status).toBe('complete');
    expect(localStorage.getItem(storageKey(KEY))).not.toBeNull();
  });
});

describe('the 404 grace after an upload', () => {
  const instant = { sleep: async () => {}, apiBase: API };

  it('retries a 404 until the report appears', async () => {
    let calls = 0;
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => (++calls < 3 ? envelope(null, 404) : envelope(meta('processing')))),
    );
    const waiting = vi.fn();
    const got = await fetchReportMetaAfterUpload('fixture2abcd', {
      ...instant,
      uploadedAt: 1000,
      now: () => 2000,
      onWaiting: waiting,
    });
    expect(got.status).toBe('processing');
    expect(calls).toBe(3);
    expect(waiting).toHaveBeenCalledTimes(2);
  });

  it('retries every two seconds by default', async () => {
    let calls = 0;
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => (++calls < 2 ? envelope(null, 404) : envelope(meta('processing')))),
    );
    const sleep = vi.fn(async () => {});
    await fetchReportMetaAfterUpload('fixture2abcd', { apiBase: API, uploadedAt: 0, now: () => 0, sleep });
    expect(sleep).toHaveBeenCalledWith(2000);
  });

  it('gives up with "No report with that id" once thirty seconds have passed', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 404)),
    );
    let clock = 0;
    const sleep = async (ms: number) => {
      clock += ms;
    };
    await expect(
      fetchReportMetaAfterUpload('fixture2abcd', { apiBase: API, uploadedAt: 0, now: () => clock, sleep }),
    ).rejects.toThrow(REPORT_NOT_FOUND);
    expect(clock).toBe(30_000);
  });

  it('does not retry a 404 when the grace window is already over', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(null, 404));
    vi.stubGlobal('fetch', upstream);
    await expect(
      fetchReportMetaAfterUpload('fixture2abcd', { ...instant, uploadedAt: 0, now: () => 31_000 }),
    ).rejects.toThrow(REPORT_NOT_FOUND);
    expect(upstream).toHaveBeenCalledTimes(1);
  });

  it('does not retry a refusal', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope(null, 403));
    vi.stubGlobal('fetch', upstream);
    await expect(
      fetchReportMetaAfterUpload('fixture2abcd', { ...instant, uploadedAt: 0, now: () => 0 }),
    ).rejects.toThrow();
    expect(upstream).toHaveBeenCalledTimes(1);
  });
});

describe('the just-uploaded mark', () => {
  it('is per report id, and absent for one never uploaded here', () => {
    expect(justUploadedAt('aaaaaaaaaaaa')).toBeNull();
    markJustUploaded('aaaaaaaaaaaa', 1234);
    expect(justUploadedAt('aaaaaaaaaaaa')).toBe(1234);
    expect(justUploadedAt('bbbbbbbbbbbb')).toBeNull();
  });
});
