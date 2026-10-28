// web/src/lib/billing/premium-owned.test.ts
import { describe, expect, it } from 'vitest';
import type { BillingView, Me } from '../account/api';
import { ME_SEAT, mockEntitlements } from '../../test-support/billing-fixtures';
import { isPremiumAlreadyOwned, premiumOwnedView } from './premium-owned';
import { BillingApiError } from './api';

const NOW = new Date('2026-11-12T12:00:00Z');
const STRIPE: BillingView = {
  plan: 'premium',
  status: 'active',
  current_period_end: '2026-12-22T00:00:00Z',
  cancel_at_period_end: false,
  founder: false,
  interval: 'monthly',
  amount_cents: 500,
  source: 'stripe',
};
const GRANT: BillingView = {
  ...STRIPE,
  source: 'grant',
  current_period_end: null,
  amount_cents: null,
  interval: null,
};

function meWith(billing: BillingView | null, extra: Partial<Me> = {}): Me {
  return {
    user: { id: 7, battletag: 'F#1', email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ billing }),
    ...extra,
  };
}

describe('premiumOwnedView', () => {
  it('own: plan line, Manage billing opens the portal', () => {
    const view = premiumOwnedView(meWith(STRIPE), NOW, 'UTC');
    expect(view.title).toBe('You already have Premium.');
    expect(view.line).toContain('22');
    expect(view.action).toEqual({ kind: 'portal', label: 'Manage billing' });
    expect(view.pastDue).toBe(false);
  });

  it('own, past due: the alert flag is set', () => {
    expect(premiumOwnedView(meWith({ ...STRIPE, status: 'past_due' }), NOW).pastDue).toBe(true);
  });

  it('grant: both lines, a See your plan link', () => {
    const view = premiumOwnedView(meWith(GRANT), NOW);
    expect(view.line).toBe('Premium · granted, no end date');
    expect(view.line2).toBe('Given by Forever Sixty. Nothing is billed to you.');
    expect(view.action).toEqual({ kind: 'link', label: 'See your plan', href: '/account#plan' });
    expect(view.pastDue).toBe(false);
  });

  it('grant with an end date', () => {
    const view = premiumOwnedView(
      meWith({ ...GRANT, current_period_end: '2027-01-05T00:00:00Z' }),
      NOW,
      'UTC',
    );
    expect(view.line).toContain('granted until');
    expect(view.line2).toContain('ends on its own');
  });

  it('seat: the seat line and a See your plan link', () => {
    const view = premiumOwnedView(meWith(null, { seat: { ...ME_SEAT, starts_at: null } }), NOW);
    expect(view.line).toContain(ME_SEAT.guild_name);
    expect(view.action.kind).toBe('link');
  });

  it('an account with no visible Premium (source unknown) gets the generic panel', () => {
    const view = premiumOwnedView(meWith(null), NOW);
    expect(view.line).toBe('Premium');
    expect(view.line2).toBeNull();
    expect(view.action).toEqual({ kind: 'link', label: 'See your plan', href: '/account#plan' });
  });
});

describe('isPremiumAlreadyOwned', () => {
  it('is a 409 carrying a source or the conflict code', () => {
    expect(isPremiumAlreadyOwned(new BillingApiError('x', 409, null, { source: 'grant' }))).toBe(true);
    expect(isPremiumAlreadyOwned(new BillingApiError('x', 409, 'conflict'))).toBe(true);
    expect(isPremiumAlreadyOwned(new BillingApiError('x', 409, 'no_transfer_offer'))).toBe(false);
    expect(isPremiumAlreadyOwned(new BillingApiError('x', 503))).toBe(false);
    expect(isPremiumAlreadyOwned(new Error('x'))).toBe(false);
  });
});
