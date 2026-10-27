import { describe, expect, it } from 'vitest';
import {
  FOUNDER_RAID_GUILD_READ,
  GUILD_TIER_28_READ,
  LAPSED_READ,
  TRANSFER_RECEIVED,
  TRIAL_READ,
  emptySeat,
  guildBillingRead,
  planView,
} from '../../test-support/billing-fixtures';
import {
  billingState,
  canBill,
  fitsTier,
  lineTwoState,
  movesFromPlanLine,
  offeredToReader,
  planSeatCount,
  showSizeLine,
  showsChooser,
  showsReadOnlyLine,
} from './billing-state';

const NOW = new Date('2026-11-15T12:00:00Z');

describe('billingState', () => {
  it('is none before a plan, and for a missing read', () => {
    expect(billingState(guildBillingRead(), NOW)).toEqual({ kind: 'none' });
    expect(billingState(null, NOW)).toEqual({ kind: 'none' });
  });

  it('reads a trial with its first charge', () => {
    expect(billingState(TRIAL_READ, NOW)).toEqual({
      kind: 'trial',
      tier: 'guild',
      firstCharge: '2026-11-04T00:00:00Z',
      amountCents: 1500,
      interval: 'monthly',
    });
  });

  it('reads an active plan with its renewal and price', () => {
    expect(billingState(FOUNDER_RAID_GUILD_READ, NOW)).toMatchObject({
      kind: 'active',
      tier: 'raid_guild',
      amountCents: 20000,
      interval: 'yearly',
    });
  });

  it('reads cancel at period end as ending, and as ending-transfer once a transfer is accepted', () => {
    const ending = guildBillingRead({ billing: planView({ cancel_at_period_end: true }) });
    expect(billingState(ending, NOW)).toMatchObject({ kind: 'ending', ends: '2026-12-01T00:00:00Z' });
    const handed = {
      ...ending,
      transfer: {
        to_name: 'Mirelle',
        to_you: false,
        from_name: 'x',
        at: '',
        state: 'accepted' as const,
        starts_at: null,
        amount_cents: null,
        interval: null,
        founder: false,
      },
    };
    expect(billingState(handed, NOW)).toMatchObject({ kind: 'ending-transfer', toName: 'Mirelle' });
  });

  it('reads a canceled plan, and an ending plan past its end, as lapsed', () => {
    expect(billingState(LAPSED_READ, NOW)).toEqual({ kind: 'lapsed', ended: '2026-10-01T00:00:00Z' });
    const past = guildBillingRead({
      billing: planView({ cancel_at_period_end: true, current_period_end: '2026-11-01T00:00:00Z' }),
    });
    expect(billingState(past, NOW).kind).toBe('lapsed');
  });

  it('reads an operator grant, scoped or full', () => {
    const scoped = guildBillingRead({ billing: planView({ scope: 'officer_tools', billed_by: '' }) });
    expect(billingState(scoped, NOW)).toEqual({ kind: 'grant', scoped: true });
    const full = guildBillingRead({ billing: planView({ billed_by: '', you_are_billing_contact: false }) });
    expect(billingState(full, NOW)).toEqual({ kind: 'grant', scoped: false });
  });
});

describe('the derived rules', () => {
  it('fitsTier switches above 25', () => {
    expect(fitsTier(25)).toBe('guild');
    expect(fitsTier(26)).toBe('raid_guild');
  });

  it('showSizeLine needs Guild tier over 25 members', () => {
    expect(showSizeLine(GUILD_TIER_28_READ)).toBe(true);
    expect(showSizeLine({ ...GUILD_TIER_28_READ, verified_count: 25 })).toBe(false);
    expect(showSizeLine(FOUNDER_RAID_GUILD_READ)).toBe(false);
    expect(showSizeLine(guildBillingRead({ verified_count: 40 }))).toBe(false);
  });

  it('the chooser shows for none, lapsed and the scoped grant only', () => {
    expect(showsChooser({ kind: 'none' })).toBe(true);
    expect(showsChooser({ kind: 'lapsed', ended: '' })).toBe(true);
    expect(showsChooser({ kind: 'grant', scoped: true })).toBe(true);
    expect(showsChooser({ kind: 'grant', scoped: false })).toBe(false);
    expect(showsChooser(billingState(TRIAL_READ, NOW))).toBe(false);
  });

  it('Move to Raid guild draws on the plan line only when the size line does not', () => {
    expect(movesFromPlanLine(TRIAL_READ, billingState(TRIAL_READ, NOW))).toBe(true);
    expect(movesFromPlanLine(GUILD_TIER_28_READ, billingState(GUILD_TIER_28_READ, NOW))).toBe(false);
    expect(movesFromPlanLine(FOUNDER_RAID_GUILD_READ, billingState(FOUNDER_RAID_GUILD_READ, NOW))).toBe(
      false,
    );
  });

  it('canBill reads the billing word', () => {
    expect(canBill(TRIAL_READ)).toBe(true);
    expect(canBill(LAPSED_READ)).toBe(false);
  });

  it('line 2 follows the plan line', () => {
    const active = billingState(GUILD_TIER_28_READ, NOW);
    expect(lineTwoState(active, planView())).toEqual({ kind: 'billed-to-you' });
    expect(lineTwoState(active, planView({ you_are_billing_contact: false }))).toEqual({
      kind: 'billed-by',
      name: 'Thornwick',
    });
    expect(lineTwoState({ kind: 'lapsed', ended: '' }, null)).toEqual({ kind: 'read-only' });
    expect(lineTwoState({ kind: 'none' }, null)).toBeNull();
  });
});

describe('the read-only line and the seat count', () => {
  const plan = planView({ you_are_billing_contact: false });

  it('a non-payer reads the line unless an offer to take over billing waits for them', () => {
    expect(showsReadOnlyLine(guildBillingRead({ billing: plan }), false)).toBe(true);
    const offered = guildBillingRead({ billing: plan, transfer: TRANSFER_RECEIVED });
    expect(offeredToReader(offered)).toBe(true);
    expect(showsReadOnlyLine(offered, false)).toBe(false);
    expect(showsReadOnlyLine(guildBillingRead({ billing: plan }), true)).toBe(false);
  });

  it('the seat count is the plan tier own, not the rows sent', () => {
    const raid = guildBillingRead({ billing: planView({ tier: 'raid_guild' }), seats: [emptySeat(1)] });
    expect(planSeatCount(raid)).toBe(3);
    expect(planSeatCount(guildBillingRead({ billing: planView() }))).toBe(1);
    expect(planSeatCount(guildBillingRead({ seats: [emptySeat(1), emptySeat(2)] }))).toBe(2);
  });
});
