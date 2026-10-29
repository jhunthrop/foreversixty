// web/src/components/account/AccountPlanPanel.test.ts
// The Plan panel's headline by source (design/specs/2026-10-27-subscribe-and-join.md 4.3).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { BillingView, Me } from '../../lib/account/api';
import { accountPlanLine } from '../../lib/account/plan-line';
import { ME_SEAT, mockEntitlements } from '../../test-support/billing-fixtures';
import AccountPremiumBlock from './AccountPremiumBlock.svelte';

const NOW = new Date('2026-11-12T12:00:00Z');
const ZONE = 'UTC';

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
const GRANT_OPEN: BillingView = {
  ...STRIPE,
  source: 'grant',
  current_period_end: null,
  amount_cents: null,
  interval: null,
};
const HELD_SEAT = { ...ME_SEAT, starts_at: null };

function meWith(extra: Partial<Me>, billing: BillingView | null = null, portal = false): Me {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ billing, personal_portal: portal }),
    ...extra,
  };
}

function panel(me: Me): string {
  const plan = accountPlanLine(me, NOW, ZONE);
  const pastDue = me.entitlements?.billing?.status === 'past_due' && plan.billingButton === 'manage';
  return render(AccountPremiumBlock, {
    props: { plan, pastDue, onBilling: () => undefined, seePlans: plan.line === 'Not on Premium.' },
  }).body;
}

const BUTTON = 'data-testid="account-billing-button"';

describe('Account plan headline', () => {
  it('shows a grant with no end date and no billing button (the owner account)', () => {
    const html = panel(meWith({}, GRANT_OPEN));
    expect(html).toContain('Premium · granted, no end date');
    expect(html).toContain('Given by Forever Sixty. Nothing is billed to you.');
    expect(html).not.toContain(BUTTON);
    expect(html).not.toMatch(/Manage billing|Billing history/);
  });

  it('shows Billing history for a seat with a personal portal', () => {
    const html = panel(meWith({ seat: HELD_SEAT }, null, true));
    expect(html).toContain(BUTTON);
    expect(html).toContain('Billing history');
    expect(html).not.toContain('Manage billing');
  });

  it('shows no button for a seat without a portal', () => {
    expect(panel(meWith({ seat: HELD_SEAT }))).not.toContain(BUTTON);
  });

  it('shows Manage billing and the founder pill for an own plan', () => {
    const html = panel(meWith({}, { ...STRIPE, founder: true }));
    expect(html).toContain('Manage billing');
    expect(html).toContain('account-founder');
  });

  it('shows the alert in line 2 place for a past-due own plan', () => {
    const html = panel(meWith({}, { ...STRIPE, status: 'past_due' }));
    expect(html).toContain('role="alert"');
    expect(html).toContain('Your last payment failed.');
    expect(html).toContain('Manage billing');
    expect(html.indexOf('account-billing-row')).toBeLessThan(html.indexOf('account-past-due'));
    expect(html.indexOf('account-past-due')).toBeLessThan(html.indexOf('account-premium-block'));
  });

  it('puts the billing button after the feature rows', () => {
    const html = panel(meWith({}, STRIPE));
    expect(html.indexOf('account-premium-block')).toBeLessThan(html.indexOf(BUTTON));
  });

  it('shows no alert for an active plan', () => {
    expect(panel(meWith({}, STRIPE))).not.toContain('role="alert"');
  });

  it('shows See plans, linked to /premium, on the free plan', () => {
    const html = panel(meWith({}));
    expect(html).toContain('Not on Premium.');
    expect(html).toContain('href="/premium#premium"');
    expect(html).toContain('See plans');
  });
});
