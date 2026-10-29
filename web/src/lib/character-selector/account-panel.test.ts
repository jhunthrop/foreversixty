import { describe, expect, it } from 'vitest';
import { FOUNDER_PREMIUM_FACTS, ME_SEAT, mockEntitlements } from '../../test-support/billing-fixtures';
import type { BillingView, Me, MeCharacter } from '../account/api';
import { accountPanelOf } from './account-panel';

const NOW = Date.parse('2026-11-12T12:00:00Z');
const OWN: BillingView = {
  plan: 'premium',
  status: 'active',
  current_period_end: '2026-12-20T00:00:00Z',
  cancel_at_period_end: false,
  founder: false,
  interval: 'monthly',
  amount_cents: 600,
};
const SHAMAN: MeCharacter = {
  key: 'us/normal/lightningbolltz',
  region: 'us',
  ruleset: 'normal',
  name: 'Lightningbolltz',
  class: 'Shaman',
};

function meWith(user: Partial<Me['user']>, billing: BillingView | null = null, extra: Partial<Me> = {}): Me {
  return {
    user: { id: 1, battletag: 'Hunthrop#1894', email: null, role: 'user', anonymize: false, ...user },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ billing }),
    ...extra,
  };
}

describe('accountPanelOf', () => {
  it('heads a Battle.net account with its BattleTag', () => {
    const panel = accountPanelOf(meWith({}), null, NOW);
    expect(panel.head).toEqual({ name: 'Hunthrop#1894', via: 'bnet', premium: false });
  });

  it('heads an email account with the whole address', () => {
    const panel = accountPanelOf(meWith({ battletag: null, email: 'a@b.co' }), null, NOW);
    expect(panel.head).toMatchObject({ name: 'a@b.co', via: 'email' });
  });

  it('tags Premium for a live plan or a held seat, not for none', () => {
    expect(accountPanelOf(meWith({}, OWN), null, NOW).head.premium).toBe(true);
    expect(accountPanelOf(meWith({}, null, { seat: ME_SEAT }), null, NOW).head.premium).toBe(true);
    expect(accountPanelOf(meWith({}), null, NOW).head.premium).toBe(false);
  });

  it('a free account: Plan reads Free and goes to /premium', () => {
    expect(accountPanelOf(meWith({}), null, NOW).plan).toEqual({
      sub: 'Free',
      sub2: null,
      href: '/premium',
      founder: false,
    });
  });

  it('an own plan reads its price over its renewal date and goes to /account#plan', () => {
    const plan = accountPanelOf(meWith({}, OWN), null, NOW).plan;
    expect(plan.sub).toBe('Premium · $6/mo');
    expect(plan.sub2).toMatch(/^Renews \d+ Dec 2026$/);
    expect(plan.href).toBe('/account#plan');
  });

  it('a yearly plan that stops renewing says Ends', () => {
    const plan = accountPanelOf(
      meWith({}, { ...OWN, interval: 'yearly', amount_cents: 5000, cancel_at_period_end: true }),
      null,
      NOW,
    ).plan;
    expect(plan.sub).toBe('Premium · $50/yr');
    expect(plan.sub2).toMatch(/^Ends \d+ Dec 2026$/);
  });

  it('a seat keeps the account page line verbatim with no second line', () => {
    const plan = accountPanelOf(meWith({}, null, { seat: ME_SEAT }), null, NOW).plan;
    expect(plan.sub).toMatch(/^Premium · a seat from /);
    expect(plan.sub2).toBeNull();
  });

  it('a founder price carries the Founder tag', () => {
    const founder = { ...OWN, ...FOUNDER_PREMIUM_FACTS };
    expect(accountPanelOf(meWith({}, founder), null, NOW).plan.founder).toBe(true);
  });

  it('Your profile exists only for a character saved to the account', () => {
    expect(accountPanelOf(meWith({}), SHAMAN, NOW).profile).toEqual({
      href: '/character/us/normal/lightningbolltz',
    });
    expect(accountPanelOf(meWith({}), null, NOW).profile).toBeNull();
  });
});
