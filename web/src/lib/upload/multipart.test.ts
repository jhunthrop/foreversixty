// @vitest-environment jsdom
// web/src/lib/upload/multipart.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  MAX_UPLOAD_BYTES,
  PART_SIZE_BYTES,
  UPLOAD_CANCELLED,
  UPLOAD_TOO_LARGE,
  completeUpload,
  createUpload,
  putPart,
  uploadParts,
  type CreatedUpload,
} from './multipart';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

/** The three-part XMLHttpRequest surface putPart uses, and nothing else. */
class FakeXhr {
  static sent: { url: string; size: number }[] = [];
  static failUntil = 0;
  static attempts = 0;
  /** Urls whose signature has expired: R2 answers 403 and no retry of the same url helps. */
  static expired = new Set<string>();
  /** Called as soon as a send is in flight, so a test can cancel mid-part. */
  static onSend: (() => void) | null = null;

  status = 0;
  upload = { onprogress: null as ((event: { loaded: number; total: number }) => void) | null };
  onload: (() => void) | null = null;
  onerror: (() => void) | null = null;
  onabort: (() => void) | null = null;
  private url = '';
  private etag = '"etag-1"';
  private done = false;

  open(_method: string, url: string): void {
    this.url = url;
  }

  getResponseHeader(name: string): string | null {
    return name.toLowerCase() === 'etag' ? this.etag : null;
  }

  send(body: Blob): void {
    FakeXhr.attempts += 1;
    FakeXhr.onSend?.();
    queueMicrotask(() => {
      if (this.done) return;
      this.done = true;
      if (FakeXhr.expired.has(this.url)) {
        this.status = 403;
        this.onload?.();
        return;
      }
      if (FakeXhr.attempts <= FakeXhr.failUntil) {
        this.status = 500;
        this.onload?.();
        return;
      }
      this.upload.onprogress?.({ loaded: body.size / 2, total: body.size });
      this.upload.onprogress?.({ loaded: body.size, total: body.size });
      this.status = 200;
      this.etag = `"etag-${FakeXhr.sent.length + 1}"`;
      FakeXhr.sent.push({ url: this.url, size: body.size });
      this.onload?.();
    });
  }

  abort(): void {
    if (this.done) return;
    this.done = true;
    this.onabort?.();
  }
}

function fakeFile(size: number, name = 'WoWCombatLog.txt'): File {
  // A sparse Blob: the tests only ever read `size` and slice it.
  const blob = new Blob([new Uint8Array(Math.min(size, 1024))]);
  return Object.defineProperty(new File([blob], name, { type: 'text/plain' }), 'size', {
    value: size,
  }) as File;
}

afterEach(() => {
  vi.unstubAllGlobals();
  FakeXhr.sent = [];
  FakeXhr.failUntil = 0;
  FakeXhr.attempts = 0;
  FakeXhr.expired = new Set();
  FakeXhr.onSend = null;
});

describe('createUpload', () => {
  it('asks the API for signed part urls and sends the CSRF header', async () => {
    const created = {
      upload_id: 'up1',
      parts: [{ number: 1, url: 'https://r2.example/part1' }],
      complete_url: 'https://r2.example/complete',
    };
    const upstream = vi.fn<GlobalFetch>(async () => envelope(created));
    vi.stubGlobal('fetch', upstream);
    document.cookie = 'fs_csrf=tok; path=/';

    await expect(createUpload(fakeFile(1024), API)).resolves.toEqual(created);

    const request = upstream.mock.calls[0][0] as Request;
    expect(request.url).toBe(`${API}/v1/uploads`);
    expect(request.headers.get('x-csrf-token')).toBe('tok');
    expect(await request.json()).toEqual({ size_bytes: 1024, filename: 'WoWCombatLog.txt' });
  });

  it('refuses a file over the ceiling before asking the API anything', async () => {
    const upstream = vi.fn<GlobalFetch>();
    vi.stubGlobal('fetch', upstream);
    await expect(createUpload(fakeFile(MAX_UPLOAD_BYTES + 1), API)).rejects.toThrow(UPLOAD_TOO_LARGE);
    expect(upstream).not.toHaveBeenCalled();
  });
});

describe('putPart', () => {
  it('PUTs the bytes and returns the ETag R2 gave back', async () => {
    const seen: number[] = [];
    const etag = await putPart(
      'https://r2.example/part1',
      new Blob(['hello']),
      (loaded) => seen.push(loaded),
      () => new FakeXhr() as unknown as XMLHttpRequest,
    );
    expect(etag).toBe('"etag-1"');
    expect(seen).toEqual([2.5, 5]);
  });
});

describe('uploadParts', () => {
  const created: CreatedUpload = {
    upload_id: 'up1',
    parts: [
      { number: 1, url: 'https://r2.example/p1' },
      { number: 2, url: 'https://r2.example/p2' },
      { number: 3, url: 'https://r2.example/p3' },
    ],
    complete_url: 'https://r2.example/complete',
  };

  it('slices the file at the contract’s 64 MiB part size and reports bytes as they go', async () => {
    expect(PART_SIZE_BYTES).toBe(64 * 1024 * 1024);
    const file = fakeFile(PART_SIZE_BYTES * 2 + 10);
    const seen: number[] = [];

    const { uploadId, etags } = await uploadParts(
      file,
      created,
      (progress) => seen.push(progress.uploadedBytes),
      { makeXhr: () => new FakeXhr() as unknown as XMLHttpRequest, retryDelayMs: 0 },
    );

    expect(uploadId).toBe('up1');
    expect(etags).toEqual([
      { number: 1, etag: '"etag-1"' },
      { number: 2, etag: '"etag-2"' },
      { number: 3, etag: '"etag-3"' },
    ]);
    expect(FakeXhr.sent.map((part) => part.size)).toEqual([PART_SIZE_BYTES, PART_SIZE_BYTES, 10]);
    expect(seen[seen.length - 1]).toBe(file.size);
    expect(seen).toEqual([...seen].sort((a, b) => a - b));
  });

  it('retries a failed part instead of restarting the upload', async () => {
    FakeXhr.failUntil = 2;
    const file = fakeFile(10);

    const { etags } = await uploadParts(file, { ...created, parts: [created.parts[0]] }, () => {}, {
      makeXhr: () => new FakeXhr() as unknown as XMLHttpRequest,
      retryDelayMs: 0,
    });

    expect(etags).toHaveLength(1);
    expect(FakeXhr.attempts).toBe(3);
    expect(FakeXhr.sent).toHaveLength(1);
  });

  it('gives up after three attempts on one part', async () => {
    FakeXhr.failUntil = 99;
    await expect(
      uploadParts(fakeFile(10), { ...created, parts: [created.parts[0]] }, () => {}, {
        makeXhr: () => new FakeXhr() as unknown as XMLHttpRequest,
        retryDelayMs: 0,
      }),
    ).rejects.toThrow('part 1');
    expect(FakeXhr.attempts).toBe(3);
  });

  // The signed part urls last an hour and the ceiling is 4 GiB, so an upload on a home
  // uplink can outlive its own signatures. Before this it threw away every part already
  // sent and said "part 2 failed three times"; the whole night had to be uploaded again
  // from the beginning, by hand.
  it('starts a fresh upload when a signature expires, rather than losing the file', async () => {
    const file = fakeFile(PART_SIZE_BYTES + 10);
    const twoParts = { ...created, parts: created.parts.slice(0, 2) };
    FakeXhr.expired = new Set(['https://r2.example/p2']);
    const restart = vi.fn(async () => ({
      upload_id: 'up2',
      parts: [
        { number: 1, url: 'https://r2.example/fresh1' },
        { number: 2, url: 'https://r2.example/fresh2' },
      ],
      complete_url: 'https://r2.example/complete2',
    }));

    const { uploadId, etags } = await uploadParts(file, twoParts, () => {}, {
      makeXhr: () => new FakeXhr() as unknown as XMLHttpRequest,
      retryDelayMs: 0,
      restart,
    });

    expect(restart).toHaveBeenCalledTimes(1);
    // Every part again, against the fresh upload: POST /v1/uploads mints a new object key
    // and a new R2 multipart upload (api/internal/reports/uploads.go), so the parts already
    // sent belong to an object the new upload could never be completed with. The id that
    // comes back is the one that has to be completed.
    expect(uploadId).toBe('up2');
    expect(etags).toHaveLength(2);
    expect(FakeXhr.sent.map((part) => part.url)).toEqual([
      'https://r2.example/p1',
      'https://r2.example/fresh1',
      'https://r2.example/fresh2',
    ]);
  });

  it('does not retry an expired signature against the same url, and restarts only once', async () => {
    FakeXhr.expired = new Set(['https://r2.example/p1', 'https://r2.example/fresh1']);
    const restart = vi.fn(async () => ({
      upload_id: 'up2',
      parts: [{ number: 1, url: 'https://r2.example/fresh1' }],
      complete_url: 'https://r2.example/complete2',
    }));

    await expect(
      uploadParts(fakeFile(10), { ...created, parts: [created.parts[0]] }, () => {}, {
        makeXhr: () => new FakeXhr() as unknown as XMLHttpRequest,
        retryDelayMs: 0,
        restart,
      }),
    ).rejects.toThrow('403');

    // One attempt each, not three: retrying a url whose signature no longer verifies is
    // three more 403s. And one restart, not a loop.
    expect(FakeXhr.attempts).toBe(2);
    expect(restart).toHaveBeenCalledTimes(1);
  });

  // The signal used to be read between parts only, and nothing ever called xhr.abort(), so
  // Cancel did nothing at all for as long as the 64 MiB part in flight took to finish.
  it('aborts the request in flight when the signal fires, and sends nothing after it', async () => {
    const controller = new AbortController();
    FakeXhr.onSend = () => controller.abort();

    await expect(
      uploadParts(fakeFile(PART_SIZE_BYTES + 10), created, () => {}, {
        makeXhr: () => new FakeXhr() as unknown as XMLHttpRequest,
        retryDelayMs: 0,
        signal: controller.signal,
      }),
    ).rejects.toThrow(UPLOAD_CANCELLED);

    expect(FakeXhr.sent).toEqual([]);
    expect(FakeXhr.attempts).toBe(1);
  });

  it('never restarts an upload that was cancelled', async () => {
    const controller = new AbortController();
    controller.abort();
    const restart = vi.fn();

    await expect(
      uploadParts(fakeFile(10), created, () => {}, {
        makeXhr: () => new FakeXhr() as unknown as XMLHttpRequest,
        signal: controller.signal,
        restart,
      }),
    ).rejects.toThrow(UPLOAD_CANCELLED);
    expect(restart).not.toHaveBeenCalled();
    expect(FakeXhr.attempts).toBe(0);
  });
});

describe('completeUpload', () => {
  it('posts the etags with the title and visibility and returns the report id', async () => {
    const upstream = vi.fn<GlobalFetch>(async () => envelope({ report_id: 'fixture2abcd' }, 202));
    vi.stubGlobal('fetch', upstream);

    const id = await completeUpload(
      'up1',
      [{ number: 1, etag: '"e1"' }],
      { title: 'Tuesday', visibility: 'unlisted' },
      API,
    );

    expect(id).toBe('fixture2abcd');
    const request = upstream.mock.calls[0][0] as Request;
    expect(request.url).toBe(`${API}/v1/uploads/up1/complete`);
    expect(await request.json()).toEqual({
      etags: [{ number: 1, etag: '"e1"' }],
      title: 'Tuesday',
      visibility: 'unlisted',
    });
  });
});
