// web/src/components/checkout/PremiumOwnedPanel.test.ts
// The "You already have Premium." panel per source (spec 4.2, 9.2, case 14). Static render:
// this config's `render()` runs no effects, so the launcher's flow is tested in
// lib/billing/premium-owned-flow.test.ts and the portal click by the checkout-launcher e2e case Task 25 adds.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { BillingView, Me } from '../../lib/account/api';
import { ME_SEAT, mockEntitlements } from '../../test-support/billing-fixtures';
import PremiumOwnedPanel from './PremiumOwnedPanel.svelte';

const NOW = new Date('2026-11-12T12:00:00Z');
const STRIPE: BillingView = {
  plan: 'premium',
  status: 'active',
  current_period_end: '2026-12-22T00:00:00Z',
  cancel_at_period_end: false,
  founder: false,
  interval: 'yearly',
  amount_cents: 5000,
  source: 'stripe',
};
const GRANT_OPEN: BillingView = {
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

function html(me: Me): string {
  return render(PremiumOwnedPanel, { props: { me, now: NOW, zone: 'UTC', onPortal: () => undefined } }).body;
}

const MANAGE = '>Manage billing<';
const SEE_PLAN = 'href="/account#plan"';

describe('PremiumOwnedPanel', () => {
  it.each([
    ['own renewing', meWith(STRIPE), MANAGE],
    ['own ending', meWith({ ...STRIPE, cancel_at_period_end: true }), MANAGE],
    ['past due', meWith({ ...STRIPE, status: 'past_due' }), MANAGE],
    ['seat', meWith(null, { seat: { ...ME_SEAT, starts_at: null } }), SEE_PLAN],
    ['grant open-ended', meWith(GRANT_OPEN), SEE_PLAN],
    ['grant until', meWith({ ...GRANT_OPEN, current_period_end: '2027-01-05T00:00:00Z' }), SEE_PLAN],
  ])('%s: an h2 title and the right button', (_name, me, button) => {
    const body = html(me);
    expect(body).toMatch(/<h2[^>]*tabindex="-1"[^>]*>\s*You already have Premium\./);
    expect(body).not.toContain('<h1');
    expect(body).toContain(button);
  });

  it('grant: both plan lines, no billing button', () => {
    const body = html(meWith(GRANT_OPEN));
    expect(body).toContain('Premium · granted, no end date');
    expect(body).toContain('Given by Forever Sixty. Nothing is billed to you.');
    expect(body).not.toContain('Manage billing');
  });

  it('grant until a date names it and says it ends on its own', () => {
    const body = html(meWith({ ...GRANT_OPEN, current_period_end: '2027-01-05T00:00:00Z' }));
    expect(body).toContain('granted until');
    expect(body).toContain('ends on its own');
  });

  it('past due: the alert replaces line 2', () => {
    const body = html(meWith({ ...STRIPE, status: 'past_due' }));
    expect(body).toContain('role="alert"');
    expect(body).toContain('Your last payment failed');
  });

  it('an own plan with the founder price shows the pill', () => {
    expect(html(meWith({ ...STRIPE, founder: true }))).toContain('Founder');
    expect(html(meWith(STRIPE))).not.toContain('Founder');
  });

  it('an account whose source is unknown gets the generic panel, never an API sentence', () => {
    const body = html(meWith(null));
    expect(body).toContain('You already have Premium.');
    expect(body).toContain(SEE_PLAN);
    expect(body).not.toContain('Not on Premium');
  });
});
