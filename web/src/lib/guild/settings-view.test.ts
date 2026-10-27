import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import type { GuildBillingView } from '../account/api';
import { canManageBilling, claimedByLine, officerToolsExplanation, settingsWords } from './settings-view';

function billing(overrides: Partial<GuildBillingView> = {}): GuildBillingView {
  return {
    tier: 'guild',
    you_hold_seat: false,
    status: 'active',
    current_period_end: '2026-12-31T00:00:00Z',
    cancel_at_period_end: false,
    scope: 'plan',
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

describe('the Manage billing rule', () => {
  it('offers Manage billing only to the billing contact of a plan that has a Stripe customer', () => {
    expect(canManageBilling(billing({ you_are_billing_contact: true }))).toBe(true);
    expect(canManageBilling(billing())).toBe(false);
    expect(canManageBilling(billing({ billed_by: '', you_are_billing_contact: true }))).toBe(false);
  });
});

describe('billing copy per state', () => {
  it('a plan with no Stripe customer cannot open the portal', () => {
    expect(canManageBilling(billing({ billed_by: '', current_period_end: '2026-12-01T00:00:00Z' }))).toBe(
      false,
    );
  });

  it('the unavailable state says so and offers Try again', () => {
    expect(settingsWords.billingUnavailable).toBe(
      'Billing is not available right now. Try again in a few minutes.',
    );
    expect(settingsWords.tryAgain).toBe('Try again');
  });
});

// Every officer-tools gate string below is read out of the spec's section 7 table.
const GATE_SPEC = new URL('../../../../design/specs/2026-10-25-officer-tools-gate.md', import.meta.url);
const COPY_HEADING = '## 7. Copy, verbatim';
const DATE_PLACEHOLDER = { 'd Mon yyyy': '12 Dec 2026' };
const scoped = (overrides: Partial<GuildBillingView> = {}): GuildBillingView =>
  billing({
    scope: 'officer_tools',
    billed_by: '',
    current_period_end: '2026-12-12T00:00:00Z',
    ...overrides,
  });

describe('the officer tools scoped grant', () => {
  it('reads the explanation from the spec', () => {
    const expected = specString(GATE_SPEC, COPY_HEADING, 'Settings explanation, scoped', 0, DATE_PLACEHOLDER);
    expect(officerToolsExplanation(scoped())).toBe(expected);
  });
});

describe('the scoped explanation without an end date', () => {
  it('drops the until clause instead of rendering until .', () => {
    expect(officerToolsExplanation(scoped({ current_period_end: null }))).toBe(
      'Granted by Forever Sixty for the first raid nights: Award, Undo and Nudge only. After that they need the guild plan.',
    );
  });
});
