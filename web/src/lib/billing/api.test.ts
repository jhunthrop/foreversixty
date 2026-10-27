// web/src/lib/billing/api.test.ts
// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { BillingApiError, openPortal, startCheckout } from './api';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

function envelope(data: unknown, status = 200): Response {
  return new Response(JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

afterEach(() => vi.unstubAllGlobals());

describe('startCheckout', () => {
  it('posts plan, interval and nothing price-shaped', async () => {
    let body: unknown;
    const fetchMock = vi.fn<GlobalFetch>(async (input) => {
      body = JSON.parse((input as Request).body ? await (input as Request).text() : '{}');
      return envelope({ checkout_url: 'https://checkout.stripe.com/x' });
    });
    vi.stubGlobal('fetch', fetchMock);
    const result = await startCheckout({ plan: 'premium', interval: 'monthly' }, API);
    expect(result.checkout_url).toBe('https://checkout.stripe.com/x');
    expect(body).toEqual({ plan: 'premium', interval: 'monthly' });
  });

  it('sends comeback_id when a comeback is carried, and omits it otherwise', async () => {
    const bodies: unknown[] = [];
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async (input) => {
        bodies.push(JSON.parse(await (input as Request).text()));
        return envelope({ checkout_url: 'https://checkout.stripe.com/x' });
      }),
    );
    await startCheckout({ plan: 'premium', interval: 'monthly', comebackId: 9 }, API);
    await startCheckout({ plan: 'premium', interval: 'monthly' }, API);
    expect(bodies).toEqual([
      { plan: 'premium', interval: 'monthly', comeback_id: 9 },
      { plan: 'premium', interval: 'monthly' },
    ]);
  });

  it('carries the envelope code on the error', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(
        async () =>
          new Response(
            JSON.stringify({
              ok: false,
              data: null,
              error: { code: 'seat_locked', message: 'Locked.' },
              request_id: 'r',
            }),
            { status: 409, headers: { 'content-type': 'application/json' } },
          ),
      ),
    );
    await expect(startCheckout({ plan: 'premium', interval: 'monthly' }, API)).rejects.toMatchObject({
      status: 409,
      code: 'seat_locked',
    });
  });

  it('throws BillingApiError(503) when billing is not configured yet', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope({ message: 'billing_unavailable' }, 503)),
    );
    await expect(startCheckout({ plan: 'premium', interval: 'monthly' }, API)).rejects.toMatchObject({
      status: 503,
    });
  });

  it('throws BillingApiError(409) when a guild already has the plan', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => envelope(null, 409)),
    );
    await expect(
      startCheckout({ plan: 'guild', interval: 'monthly', guildId: 42 }, API),
    ).rejects.toBeInstanceOf(BillingApiError);
  });
});

describe('openPortal', () => {
  it('posts an empty body for personal billing and returns the portal url', async () => {
    let body: unknown;
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async (input) => {
        body = JSON.parse((input as Request).body ? await (input as Request).text() : '{}');
        return envelope({ portal_url: 'https://billing.stripe.com/p' });
      }),
    );
    const result = await openPortal(undefined, API);
    expect(result.portal_url).toBe('https://billing.stripe.com/p');
    expect(body).toEqual({});
  });
});
