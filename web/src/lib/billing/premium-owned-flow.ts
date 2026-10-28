// web/src/lib/billing/premium-owned-flow.ts
// The personal-Premium checkout's decision (design/specs/2026-10-27-subscribe-and-join.md 4.4):
// read the account first, and never start a second checkout for one that already has Premium.
// A refusal that says the account has it is answered by re-reading /v1/me, never by showing
// the API's sentence. Other failures rethrow for the caller's own messages.
import type { Me } from '../account/api';
import { premiumSource } from '../account/plan-line';
import type { startCheckout } from './api';
import { isPremiumAlreadyOwned } from './premium-owned';

export interface PremiumCheckoutDeps {
  fetchMe: () => Promise<Me | null>;
  forgetSession: () => void;
  startCheckout: typeof startCheckout;
}

export type PremiumCheckoutOutcome =
  { kind: 'owned'; me: Me } | { kind: 'redirect'; url: string } | { kind: 'login' };

export async function runPremiumCheckout(
  deps: PremiumCheckoutDeps,
  me: Me,
  request: Parameters<typeof startCheckout>[0],
  now: Date,
): Promise<PremiumCheckoutOutcome> {
  if (premiumSource(me, now) !== 'none') return { kind: 'owned', me };
  try {
    const result = await deps.startCheckout(request);
    return { kind: 'redirect', url: result.checkout_url };
  } catch (error) {
    if (!isPremiumAlreadyOwned(error)) throw error;
    deps.forgetSession();
    const fresh = await deps.fetchMe();
    return fresh === null ? { kind: 'login' } : { kind: 'owned', me: fresh };
  }
}
