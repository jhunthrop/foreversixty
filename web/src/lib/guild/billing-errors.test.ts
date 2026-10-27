import { describe, expect, it } from 'vitest';
import { BillingApiError } from '../billing/api';
import { guildBillingCopy } from '../billing/guild-billing-copy';
import { seatsCopy } from '../billing/seats-copy';
import { transferErrorLine, upgradeOutcome } from './billing-errors';

const refusal = (status: number, code: string | null): BillingApiError =>
  new BillingApiError('x', status, code);

describe('upgradeOutcome', () => {
  it('a declined card keeps the confirm and says nothing was charged', () => {
    expect(upgradeOutcome(refusal(409, 'upgrade_failed'))).toBe('declined');
  });

  it('an aged-out preview is priced again', () => {
    expect(upgradeOutcome(refusal(409, 'preview_expired'))).toBe('reopen');
  });

  it('not_upgradable reloads quietly', () => {
    expect(upgradeOutcome(refusal(409, 'not_upgradable'))).toBe('reload-quiet');
  });

  it('upgrade_unconfirmed, any 5xx and a network failure may have applied', () => {
    expect(upgradeOutcome(refusal(502, 'upgrade_unconfirmed'))).toBe('unconfirmed');
    expect(upgradeOutcome(refusal(503, 'billing_unavailable'))).toBe('unconfirmed');
    expect(upgradeOutcome(refusal(500, null))).toBe('unconfirmed');
    expect(upgradeOutcome(refusal(0, null))).toBe('unconfirmed');
  });

  it('an open transfer checkout says a change is already in progress', () => {
    expect(upgradeOutcome(refusal(409, 'transfer_in_progress'))).toBe('in-progress');
  });

  it('any other 4xx, and a throw that is not an API answer, is the generic failure', () => {
    expect(upgradeOutcome(refusal(403, 'forbidden'))).toBe('generic');
    expect(upgradeOutcome(new TypeError('bug'))).toBe('generic');
  });
});

describe('transferErrorLine', () => {
  it('maps transfer_in_progress and billing_unavailable to their own lines', () => {
    expect(transferErrorLine(refusal(409, 'transfer_in_progress'))).toBe(seatsCopy.whoPays.inProgress);
    expect(transferErrorLine(refusal(503, 'billing_unavailable'))).toBe(seatsCopy.whoPays.billingUnavailable);
  });

  it('anything else keeps the spec line', () => {
    expect(transferErrorLine(refusal(500, null))).toBe('The offer did not send. Try again.');
  });
});

describe('the strings added beyond the design spec', () => {
  it('are stated', () => {
    expect(guildBillingCopy.upgrade.unconfirmed).toBe(
      'Stripe is still confirming the change. The plan below is up to date.',
    );
    expect(seatsCopy.whoPays.inProgress).toContain('already in progress');
    expect(seatsCopy.whoPays.billingUnavailable).toBe(
      'Billing is not available right now. Try again in a few minutes.',
    );
  });
});
