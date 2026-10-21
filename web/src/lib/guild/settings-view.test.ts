import { describe, expect, it } from 'vitest';
import type { GuildBillingView } from '../account/api';
import { billedLine, canManageBilling, claimedByLine, guildPlanLine, settingsWords } from './settings-view';

function billing(overrides: Partial<GuildBillingView> = {}): GuildBillingView {
  return {
    status: 'active',
    current_period_end: '2026-12-31T00:00:00Z',
    cancel_at_period_end: false,
    billed_by: 'hunthrop#1894',
    you_are_billing_contact: false,
    ...overrides,
  };
}

describe('claimedByLine', () => {
  it('names the claimant, from the settings answer first', () => {
    expect(claimedByLine({ state: 'claimed', frozen: false }, { battletag: 'hunthrop#1894' })).toBe(
      'Claimed by hunthrop#1894',
    );
    expect(claimedByLine({ state: 'claimed', frozen: false, claimed_by_name: 'hunthrop#1894' }, null)).toBe(
      'Claimed by hunthrop#1894',
    );
  });

  it('says nothing of a guild nobody holds', () => {
    expect(claimedByLine({ state: 'unclaimed', frozen: false }, null)).toBeNull();
    expect(claimedByLine({ state: 'pending', frozen: false }, null)).toBeNull();
    expect(claimedByLine({ state: 'contested', frozen: true }, { battletag: 'hunthrop#1894' })).toBeNull();
  });
});

describe('guildPlanLine', () => {
  it('prints a Stripe plan as renewing on the stored date, in UTC', () => {
    expect(guildPlanLine(billing())).toBe('Guild plan · renews 31 Dec 2026');
  });

  it('prints a cancelled plan as ending', () => {
    expect(guildPlanLine(billing({ cancel_at_period_end: true }))).toBe('Guild plan · ends 31 Dec 2026');
  });

  it('prints an operator grant as lasting until the granted date', () => {
    expect(guildPlanLine(billing({ billed_by: '' }))).toBe('Guild plan · until 31 Dec 2026');
  });

  it('names the plan alone when no end date is stored', () => {
    expect(guildPlanLine(billing({ current_period_end: null }))).toBe('Guild plan');
  });
});

describe('billedLine and the Manage billing rule', () => {
  it('says who is billed', () => {
    expect(billedLine(billing())).toBe('Billed by hunthrop#1894');
    expect(billedLine(billing({ you_are_billing_contact: true }))).toBe('Billed to you.');
    expect(billedLine(billing({ billed_by: '' }))).toBe(settingsWords.noBillingContact);
  });

  it('offers Manage billing only to the billing contact of a plan that has a Stripe customer', () => {
    expect(canManageBilling(billing({ you_are_billing_contact: true }))).toBe(true);
    expect(canManageBilling(billing())).toBe(false);
    expect(canManageBilling(billing({ billed_by: '', you_are_billing_contact: true }))).toBe(false);
  });
});
