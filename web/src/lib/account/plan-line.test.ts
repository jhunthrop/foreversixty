import { describe, expect, it } from 'vitest';
import {
  FOUNDER_PREMIUM_FACTS,
  ME_ENDED_SEAT,
  ME_SEAT,
  mockEntitlements,
} from '../../test-support/billing-fixtures';
import type { BillingView, GuildBillingView, Me } from './api';
import { accountPlanLine, guildRowPlanLine } from './plan-line';

const ZONE = 'UTC';
const NOW = new Date('2026-11-12T12:00:00Z');

const OWN: BillingView = {
  plan: 'premium',
  status: 'active',
  current_period_end: '2026-12-20T00:00:00Z',
  cancel_at_period_end: false,
  founder: false,
  interval: 'monthly',
  amount_cents: 600,
};

function meWith(extra: Partial<Me>, billing: BillingView | null = null): Me {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ billing }),
    ...extra,
  };
}

describe('accountPlanLine', () => {
  it('own Premium renews with its price and keeps Manage billing', () => {
    const out = accountPlanLine(meWith({}, OWN), NOW, ZONE);
    expect(out).toEqual({
      line: 'Premium · renews 20 Dec 2026 · $6 a month',
      line2: null,
      founder: false,
      manageBilling: true,
      ended: null,
    });
  });

  it('a founder price carries the Founder flag', () => {
    const out = accountPlanLine(meWith({}, { ...OWN, ...FOUNDER_PREMIUM_FACTS }), NOW, ZONE);
    expect(out.founder).toBe(true);
  });

  it('own Premium that stops renewing reads its end', () => {
    const out = accountPlanLine(meWith({}, { ...OWN, cancel_at_period_end: true }), NOW, ZONE);
    expect(out.line).toBe('Premium · ends 20 Dec 2026');
  });

  it('own Premium ending into a seat names the seat and keeps Manage billing', () => {
    const out = accountPlanLine(meWith({ seat: ME_SEAT }, OWN), NOW, ZONE);
    expect(out.line).toBe('Premium · ends 20 Dec 2026 · then a seat from The Last Watch');
    expect(out.manageBilling).toBe(true);
    expect(out.line2).toBeNull();
  });

  it('a seat holder reads who gave it and has no Manage billing', () => {
    const seat = { ...ME_SEAT, starts_at: null };
    const out = accountPlanLine(meWith({ seat }), NOW, ZONE);
    expect(out.line).toBe('Premium · a seat from The Last Watch');
    expect(out.line2).toBe(
      "Given by Fixture#1234. It lasts while the guild's plan does; nothing is billed to you.",
    );
    expect(out.manageBilling).toBe(false);
    expect(out.founder).toBe(false);
  });

  it('a seat whose start has passed reads as a seat even beside a lapsed own plan', () => {
    const out = accountPlanLine(meWith({ seat: ME_SEAT }, OWN), new Date('2026-12-21T00:00:00Z'), ZONE);
    expect(out.line).toBe('Premium · a seat from The Last Watch');
    expect(out.manageBilling).toBe(false);
  });

  it('free reads Not on Premium with no ended seat', () => {
    expect(accountPlanLine(meWith({}), NOW, ZONE)).toEqual({
      line: 'Not on Premium.',
      line2: null,
      founder: false,
      manageBilling: false,
      ended: null,
    });
  });

  it('an ended seat offers the comeback link at the price they paid', () => {
    const out = accountPlanLine(meWith({ ended_seat: ME_ENDED_SEAT }), NOW, ZONE);
    expect(out.line).toBe('Not on Premium.');
    expect(out.ended).toEqual({
      text: 'Your seat from The Last Watch ended 25 Nov. ',
      action: 'Come back at $6 a month',
      href: '/premium/checkout?plan=premium&interval=monthly&comeback=9',
    });
  });

  it('an ended seat with no remembered price says Come back', () => {
    const ended = { ...ME_ENDED_SEAT, amount_cents: null, interval: null };
    expect(accountPlanLine(meWith({ ended_seat: ended }), NOW, ZONE).ended?.action).toBe('Come back');
  });
});

describe('guildRowPlanLine', () => {
  const plan: GuildBillingView = {
    tier: 'guild',
    you_hold_seat: false,
    scope: 'plan',
    status: 'active',
    current_period_end: '2026-12-12T00:00:00Z',
    cancel_at_period_end: false,
    billed_by: 'Fixture#1234',
    you_are_billing_contact: false,
  };

  it('reads Guild plan, then Raid guild plan for the raid tier', () => {
    expect(guildRowPlanLine(plan, ZONE)).toBe('Guild plan · renews 12 Dec 2026');
    expect(guildRowPlanLine({ ...plan, tier: 'raid_guild' }, ZONE)).toBe(
      'Raid guild plan · renews 12 Dec 2026',
    );
  });

  it('adds the seat note when the account holds one', () => {
    expect(guildRowPlanLine({ ...plan, you_hold_seat: true }, ZONE)).toBe(
      'Guild plan · renews 12 Dec 2026 · you hold a seat',
    );
  });
});
