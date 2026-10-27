// web/src/lib/guild/billing-errors.ts
// What each refusal of the guild billing writes means for the section: pure, so every API answer
// (409 upgrade_failed, preview_expired, not_upgradable, transfer_in_progress; 502 upgrade_unconfirmed;
// 503 billing_unavailable) is pinned without a page.
import { BillingApiError } from '../billing/api';
import { seatsCopy } from '../billing/seats-copy';

export type UpgradeOutcome =
  /** The preview aged out: price it again and say so. */
  | 'reopen'
  /** Nothing was charged: say so and keep the confirm. */
  | 'declined'
  /** The plan changed under the reader (or is already Raid guild): show the plan as it is, no line. */
  | 'reload-quiet'
  /** The change may have applied: show the plan as it is and say Stripe is confirming. */
  | 'unconfirmed'
  /** A transfer Checkout is open (409 transfer_in_progress): nothing changed; say a change is underway. */
  | 'in-progress'
  /** Any other refusal (a 4xx that is not a clean decline, or a throw that is not an API answer). */
  | 'generic';

const code = (thrown: unknown): string | null => (thrown instanceof BillingApiError ? thrown.code : null);

export function upgradeOutcome(thrown: unknown): UpgradeOutcome {
  if (!(thrown instanceof BillingApiError)) return 'generic';
  if (thrown.code === 'preview_expired') return 'reopen';
  if (thrown.code === 'not_upgradable') return 'reload-quiet';
  if (thrown.code === 'transfer_in_progress') return 'in-progress';
  if (thrown.code === 'upgrade_unconfirmed' || thrown.status >= 500 || thrown.status === 0) {
    return 'unconfirmed';
  }
  return thrown.status === 409 && thrown.code === 'upgrade_failed' ? 'declined' : 'generic';
}

/** 409 transfer_in_progress: name the waiting offer when there is one, else a plan change. */
export function inProgressLine(openOfferTo: string | null): string {
  return openOfferTo === null || openOfferTo === ''
    ? seatsCopy.whoPays.inProgress
    : seatsCopy.whoPays.inProgressOffer(openOfferTo);
}

/** The line under a failed transfer send or withdraw; `openOfferTo` names an offer still waiting. */
export function transferErrorLine(thrown: unknown, openOfferTo: string | null = null): string {
  switch (code(thrown)) {
    case 'transfer_in_progress':
      return inProgressLine(openOfferTo);
    case 'billing_unavailable':
      return seatsCopy.whoPays.billingUnavailable;
    default:
      return seatsCopy.whoPays.failed;
  }
}
