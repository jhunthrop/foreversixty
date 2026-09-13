// web/src/lib/planner/share.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import { cardUrlFor, RATE_LIMIT_MESSAGE, SAVE_FAILED_MESSAGE, saveBuild, shareUrlFor } from './share';
import type { BuildDraft } from './types';

const draft: BuildDraft = {
  class_id: 1,
  race_id: 1,
  tree_version: '1.15.9.69722',
  point_order: [1001],
  gear: {},
};

type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function stubResponse(status: number, body: unknown) {
  const spy = vi.fn<GlobalFetch>(
    async () =>
      new Response(JSON.stringify(body), {
        status,
        headers: { 'content-type': 'application/json' },
      }),
  );
  vi.stubGlobal('fetch', spy);
  return spy;
}

afterEach(() => vi.unstubAllGlobals());

describe('share urls', () => {
  it('point at the site, not the API', () => {
    expect(shareUrlFor('k7x2qm4a')).toBe('https://foreversixty.gg/b/k7x2qm4a');
    expect(cardUrlFor('k7x2qm4a')).toBe('https://foreversixty.gg/b/k7x2qm4a/card.png');
  });
});

describe('saveBuild', () => {
  it('posts the draft as the contract body and returns the saved build', async () => {
    const spy = stubResponse(201, {
      ok: true,
      data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
      error: null,
      request_id: 'r1',
    });
    await expect(saveBuild(draft, 'https://api.test')).resolves.toEqual({
      ok: true,
      build: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
    });
    const [url, init] = spy.mock.calls[0];
    expect(String(url)).toBe('https://api.test/v1/builds');
    expect(init?.method).toBe('POST');
    expect(JSON.parse(String(init?.body))).toEqual(draft);
  });

  it('treats a 200 the same as a 201, because the id already existed', async () => {
    stubResponse(200, {
      ok: true,
      data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
      error: null,
      request_id: 'r1',
    });
    await expect(saveBuild(draft, 'https://api.test')).resolves.toMatchObject({ ok: true });
  });

  it('surfaces the API message and its per-field messages on a 400', async () => {
    stubResponse(400, {
      ok: false,
      data: null,
      error: {
        message: 'build is not valid',
        fields: { 'point_order[7]': 'Tier 2 of Holy needs 10 points in Holy first' },
      },
      request_id: 'r1',
    });
    await expect(saveBuild(draft, 'https://api.test')).resolves.toEqual({
      ok: false,
      message: 'build is not valid',
      fields: { 'point_order[7]': 'Tier 2 of Holy needs 10 points in Holy first' },
    });
  });

  it('uses the rate-limit wording on a 429', async () => {
    stubResponse(429, { ok: false, data: null, error: { message: 'slow down' }, request_id: 'r1' });
    await expect(saveBuild(draft, 'https://api.test')).resolves.toEqual({
      ok: false,
      message: RATE_LIMIT_MESSAGE,
      fields: {},
    });
  });

  it('falls back to the generic message when the API is unreachable', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(saveBuild(draft, 'https://api.test')).resolves.toEqual({
      ok: false,
      message: SAVE_FAILED_MESSAGE,
      fields: {},
    });
  });

  it('falls back to the generic message when the body is not the envelope', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn(async () => new Response('<html>502</html>', { status: 502 })),
    );
    await expect(saveBuild(draft, 'https://api.test')).resolves.toEqual({
      ok: false,
      message: SAVE_FAILED_MESSAGE,
      fields: {},
    });
  });
});
