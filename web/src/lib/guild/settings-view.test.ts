import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import type { GuildBillingView } from '../account/api';
import { billedLine, guildPlanLine } from '../billing/guild-billing-copy';
import { guildRosterModelCopy } from './copy';
import {
  billingState,
  canManageBilling,
  claimedByLine,
  officerToolsExplanation,
  notOnPlanLines,
  settingsWords,
} from './settings-view';

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
    expect(billedLine(billing({ billed_by: '' }))).toBe('Granted by Forever Sixty.');
  });

  it('offers Manage billing only to the billing contact of a plan that has a Stripe customer', () => {
    expect(canManageBilling(billing({ you_are_billing_contact: true }))).toBe(true);
    expect(canManageBilling(billing())).toBe(false);
    expect(canManageBilling(billing({ billed_by: '', you_are_billing_contact: true }))).toBe(false);
  });
});

describe('billing copy per state', () => {
  it('not on the plan: the fact, then the real price and what officers get', () => {
    const lines = notOnPlanLines();
    expect(lines.headline).toBe('This guild is not on the guild plan yet.');
    expect(lines.detail).toBe(
      '$15 a month, cancel any time. Officers get the loot council helper and the raid readiness board, with more as they ship.',
    );
    expect(notOnPlanLines(20).detail).toContain('$20 a month');
  });

  it('with no price it says monthly and prints no number', () => {
    for (const none of [0, Number.NaN, -1]) {
      const detail = notOnPlanLines(none).detail;
      expect(detail).toBe(
        'Monthly, cancel any time. Officers get the loot council helper and the raid readiness board, with more as they ship.',
      );
      expect(detail).not.toMatch(/\d/);
    }
  });

  it('an operator grant is told apart from a Stripe plan and from no plan', () => {
    expect(billingState(null, false)).toBe('not-on-plan');
    expect(billingState(billing(), false)).toBe('on-plan');
    expect(billingState(billing({ billed_by: '' }), false)).toBe('granted');
    expect(billingState(billing(), true)).toBe('unavailable');
    const granted = billing({ billed_by: '', current_period_end: '2026-12-01T00:00:00Z' });
    expect(guildPlanLine(granted)).toBe('Guild plan · until 1 Dec 2026');
    expect(billedLine(granted)).toBe('Granted by Forever Sixty.');
    expect(canManageBilling(granted)).toBe(false);
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
  it('is its own billing state, ahead of the billed_by rule', () => {
    expect(billingState(scoped(), false)).toBe('officer-tools');
    expect(billingState(scoped({ billed_by: 'hunthrop#1894' }), false)).toBe('officer-tools');
    expect(billingState(scoped(), true)).toBe('unavailable');
  });

  it('keeps a full plan on the plan states', () => {
    expect(billingState(billing(), false)).toBe('on-plan');
    expect(billingState(billing({ billed_by: '' }), false)).toBe('granted');
  });

  it('reads the plan line from the spec', () => {
    const expected = specString(GATE_SPEC, COPY_HEADING, 'Settings plan line, scoped', 0, DATE_PLACEHOLDER);
    expect(guildPlanLine(scoped())).toBe(expected);
    expect(guildPlanLine(scoped())).toBe('Officer tools · until 12 Dec 2026');
  });

  it('reads the explanation from the spec', () => {
    const expected = specString(GATE_SPEC, COPY_HEADING, 'Settings explanation, scoped', 0, DATE_PLACEHOLDER);
    expect(officerToolsExplanation(scoped())).toBe(expected);
  });

  it('reads the guild-master line from the spec', () => {
    expect(guildRosterModelCopy.guildMasterOnly).toBe(
      specString(GATE_SPEC, COPY_HEADING, 'Settings, not the claim holder (reused)'),
    );
  });
});

describe('the scoped explanation without an end date', () => {
  it('drops the until clause instead of rendering until .', () => {
    expect(officerToolsExplanation(scoped({ current_period_end: null }))).toBe(
      'Granted by Forever Sixty for the first raid nights: Award, Undo and Nudge only. After that they need the guild plan.',
    );
  });
});
