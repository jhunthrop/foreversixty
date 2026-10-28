// web/src/lib/billing/premium-owned-flow.test.ts
import { describe, expect, it, vi } from 'vitest';
import type { BillingView, Me } from '../account/api';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { BillingApiError } from './api';
import { runPremiumCheckout, type PremiumCheckoutDeps } from './premium-owned-flow';

const NOW = new Date('2026-11-12T12:00:00Z');
const REQUEST = { plan: 'premium', interval: 'yearly' } as const;
const GRANT: BillingView = {
  plan: 'premium',
  status: 'active',
  current_period_end: null,
  cancel_at_period_end: false,
  founder: false,
  interval: null,
  amount_cents: null,
  source: 'grant',
};

function meWith(billing: BillingView | null): Me {
  return {
    user: { id: 7, battletag: 'F#1', email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ billing }),
  };
}

function deps(over: Partial<PremiumCheckoutDeps>): PremiumCheckoutDeps {
  return {
    fetchMe: async () => null,
    forgetSession: vi.fn(),
    startCheckout: vi.fn(),
    ...over,
  };
}

describe('runPremiumCheckout', () => {
  it('an account with Premium is owned and checkout is never called', async () => {
    const start = vi.fn();
    const me = meWith(GRANT);
    const outcome = await runPremiumCheckout(deps({ startCheckout: start }), me, REQUEST, NOW);
    expect(outcome).toEqual({ kind: 'owned', me });
    expect(start).not.toHaveBeenCalled();
  });

  it('an account without Premium redirects to checkout', async () => {
    const start = vi.fn().mockResolvedValue({ checkout_url: 'https://pay.example/x' });
    const outcome = await runPremiumCheckout(deps({ startCheckout: start }), meWith(null), REQUEST, NOW);
    expect(outcome).toEqual({ kind: 'redirect', url: 'https://pay.example/x' });
  });

  it('a 409 with fields.source forgets the session and returns the re-read account', async () => {
    const forget = vi.fn();
    const fresh = meWith(GRANT);
    const start = vi
      .fn()
      .mockRejectedValue(
        new BillingApiError('this account already has Premium', 409, 'conflict', { source: 'grant' }),
      );
    const outcome = await runPremiumCheckout(
      deps({ startCheckout: start, forgetSession: forget, fetchMe: async () => fresh }),
      meWith(null),
      REQUEST,
      NOW,
    );
    expect(forget).toHaveBeenCalled();
    expect(outcome).toEqual({ kind: 'owned', me: fresh });
  });

  it('a 409 with no source still resolves to owned (the generic panel)', async () => {
    const free = meWith(null);
    const start = vi.fn().mockRejectedValue(new BillingApiError('x', 409, 'conflict'));
    const outcome = await runPremiumCheckout(
      deps({ startCheckout: start, fetchMe: async () => free }),
      free,
      REQUEST,
      NOW,
    );
    expect(outcome).toEqual({ kind: 'owned', me: free });
  });

  it('a signed-out re-read goes to login', async () => {
    const start = vi.fn().mockRejectedValue(new BillingApiError('x', 409, 'conflict'));
    const outcome = await runPremiumCheckout(deps({ startCheckout: start }), meWith(null), REQUEST, NOW);
    expect(outcome).toEqual({ kind: 'login' });
  });

  it('other failures rethrow for the launcher to word', async () => {
    const failure = new BillingApiError('down', 503);
    const start = vi.fn().mockRejectedValue(failure);
    await expect(runPremiumCheckout(deps({ startCheckout: start }), meWith(null), REQUEST, NOW)).rejects.toBe(
      failure,
    );
  });
});
