// web/src/lib/billing/premium-owned.ts
// The checkout page's "You already have Premium." panel (design/specs/2026-10-27-subscribe-and-join.md
// 4.2 and 9.2): which words and which button the account's one Premium source gets. Pure: the
// caller passes the clock and zone. The API's refusal sentence is never shown.
import type { Me } from '../account/api';
import { accountPlanLine, premiumSource } from '../account/plan-line';
import { BillingApiError } from './api';
import { premiumSourceCopy } from './premium-source-copy';

export const ACCOUNT_PLAN_HREF = '/account#plan';
const CONFLICT_STATUS = 409;
const CONFLICT_CODE = 'conflict';

export type OwnedAction = { kind: 'portal'; label: string } | { kind: 'link'; label: string; href: string };

export interface PremiumOwnedView {
  title: string;
  line: string;
  line2: string | null;
  founder: boolean;
  /** The own subscription's last payment failed; the alert takes line 2's place. */
  pastDue: boolean;
  action: OwnedAction;
}

const SEE_PLAN: OwnedAction = {
  kind: 'link',
  label: premiumSourceCopy.buttons.seePlan,
  href: ACCOUNT_PLAN_HREF,
};

export function premiumOwnedView(me: Me, now: Date, zone?: string): PremiumOwnedView {
  const plan = accountPlanLine(me, now, zone);
  const manage = plan.billingButton === 'manage';
  // A refusal for a source this read cannot see has no plan words: the generic panel.
  const known = premiumSource(me, now) !== 'none';
  return {
    title: premiumSourceCopy.panelTitle,
    line: known ? plan.line : premiumSourceCopy.fallback,
    line2: known ? plan.line2 : null,
    founder: plan.founder,
    pastDue: manage && me.entitlements?.billing?.status === 'past_due',
    action: manage ? { kind: 'portal', label: premiumSourceCopy.buttons.manage } : SEE_PLAN,
  };
}

/** The personal-Premium checkout's one 409: the account already has Premium by some source. */
export function isPremiumAlreadyOwned(error: unknown): boolean {
  return (
    error instanceof BillingApiError &&
    error.status === CONFLICT_STATUS &&
    (error.fields?.source !== undefined || error.code === CONFLICT_CODE)
  );
}
