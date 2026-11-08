// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { AccountError } from '../account/api';
import { APP_TESTERS_PATH, ADMIN_APP_TESTERS_PATH, listAppTesters, signUpTester } from './api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

const answer = (status: number, body: unknown): Response =>
  new Response(JSON.stringify(body), { status, headers: { 'content-type': 'application/json' } });

afterEach(() => vi.unstubAllGlobals());

describe('the tester signup', () => {
  it('posts the address and phone with no cookie, and reads created', async () => {
    const upstream = vi.fn<GlobalFetch>(async () =>
      answer(201, { ok: true, data: { status: 'created' }, error: null }),
    );
    vi.stubGlobal('fetch', upstream);
    await expect(signUpTester({ email: 'you@example.com', platform: 'android' }, API)).resolves.toBe(
      'created',
    );
    const request = upstream.mock.calls[0][0] as Request;
    expect(request.url).toBe(`${API}${APP_TESTERS_PATH}`);
    expect(request.method).toBe('POST');
    expect(request.credentials).toBe('omit');
    expect(await request.json()).toEqual({ email: 'you@example.com', platform: 'android' });
  });

  it('reads an address already on the list', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => answer(200, { ok: true, data: { status: 'exists' }, error: null })),
    );
    await expect(signUpTester({ email: 'you@example.com', platform: 'ios' }, API)).resolves.toBe('exists');
  });

  it('throws the API refusal with its fields', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () =>
        answer(422, {
          ok: false,
          data: null,
          error: { code: 'invalid', message: 'no', fields: { email: 'bad' } },
        }),
      ),
    );
    const error = await signUpTester({ email: 'x', platform: 'ios' }, API).catch((e: unknown) => e);
    expect(error).toBeInstanceOf(AccountError);
    expect((error as AccountError).status).toBe(422);
    expect((error as AccountError).fields).toEqual({ email: 'bad' });
  });
});

describe("the owner's list", () => {
  it('reads every signup with the session', async () => {
    const testers = [{ email: 'a@example.com', platform: 'ios', created_at: '2026-11-08T12:00:00Z' }];
    const upstream = vi.fn<GlobalFetch>(async () =>
      answer(200, { ok: true, data: { testers }, error: null }),
    );
    vi.stubGlobal('fetch', upstream);
    await expect(listAppTesters(API)).resolves.toEqual(testers);
    const request = upstream.mock.calls[0][0] as Request;
    expect(request.url).toBe(`${API}${ADMIN_APP_TESTERS_PATH}`);
    expect(request.credentials).toBe('include');
  });
});
