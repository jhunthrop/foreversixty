// web/src/lib/account/seat-offer-flow.ts
// The seat offer's state machine, apart from the markup: which call a button makes, what
// the outcome is, and the status line each outcome reads.
import { accountPlanCopy } from '../billing/account-plan-copy';
import type { MeSeatOffer } from '../billing/contract';

export type OfferOutcome = 'idle' | 'accepting' | 'declining' | 'accepted' | 'declined' | 'failed';
export type OfferCall = (guildId: number) => Promise<void>;

/** Runs one answer; `failed` when the API refused or the network did. */
export async function answerOffer(
  call: OfferCall,
  guildId: number,
  done: 'accepted' | 'declined',
): Promise<'accepted' | 'declined' | 'failed'> {
  try {
    await call(guildId);
    return done;
  } catch {
    return 'failed';
  }
}

/** The line that replaces the offer once it is answered; empty while it is still open. */
export function offerStatusLine(outcome: OfferOutcome, offer: MeSeatOffer, zone?: string): string {
  if (outcome === 'declined') return accountPlanCopy.offer.declined;
  if (outcome !== 'accepted' || offer.own_ends_at === null) return '';
  return accountPlanCopy.offer.accepted(offer.own_ends_at, zone);
}
