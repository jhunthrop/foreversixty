// web/src/lib/guild/billing-state.ts
// What the guild settings #billing section shows, from the guild billing read: which plan line,
// which line 2, whether the chooser, the size line and the payer buttons draw. Pure, so every
// state is pinned without a page.
import { guildFitsRaidTier } from '../billing/checkout-params';
import { GUILD_TIER_MEMBER_LIMIT, PLANS, type Interval } from '../billing/plans';
import type { GuildBillingSettings, GuildPlanView, GuildTier } from '../billing/contract';
import type { GuildPlanLineState, Line2State } from '../billing/guild-billing-copy';

/** A plan an operator granted: the scoped officer tools, or a full plan with no Stripe billing contact. */
export interface GrantState {
  kind: 'grant';
  /** True for the officer-tools scope, which still offers the chooser. */
  scoped: boolean;
}

export type BillingViewState = GuildPlanLineState | { kind: 'none' } | GrantState;

const ENDED_STATUSES: readonly string[] = ['canceled', 'incomplete_expired'];

/** The tier that fits a verified count: Guild up to the limit, Raid guild above it. */
export function fitsTier(count: number): GuildTier {
  return guildFitsRaidTier(count) ? 'raid_guild' : 'guild';
}

function hasEnded(plan: GuildPlanView, now: Date): boolean {
  if (ENDED_STATUSES.includes(plan.status)) return true;
  if (!plan.cancel_at_period_end || plan.current_period_end === null) return false;
  return Date.parse(plan.current_period_end) <= now.getTime();
}

function ownPrice(plan: GuildPlanView): { amountCents: number; interval: Interval } | null {
  return plan.amount_cents === null || plan.interval === null
    ? null
    : { amountCents: plan.amount_cents, interval: plan.interval };
}

function stripeState(view: GuildBillingSettings, plan: GuildPlanView, now: Date): GuildPlanLineState | null {
  const periodEnd = plan.current_period_end ?? '';
  if (hasEnded(plan, now)) return { kind: 'lapsed', ended: periodEnd };
  if (plan.cancel_at_period_end) {
    return view.transfer?.state === 'accepted'
      ? { kind: 'ending-transfer', tier: plan.tier, ends: periodEnd, toName: view.transfer.to_name }
      : { kind: 'ending', tier: plan.tier, ends: periodEnd };
  }
  const price = ownPrice(plan);
  if (price === null) return null;
  if (plan.status === 'trialing' && plan.trial_end !== null) {
    return { kind: 'trial', tier: plan.tier, firstCharge: plan.trial_end, ...price };
  }
  return { kind: 'active', tier: plan.tier, renews: periodEnd, ...price };
}

export function billingState(view: GuildBillingSettings | null, now: Date): BillingViewState {
  const plan = view?.billing ?? null;
  if (view === null || plan === null) return { kind: 'none' };
  if (plan.scope === 'officer_tools') return { kind: 'grant', scoped: true };
  if (plan.billed_by === '') return { kind: 'grant', scoped: false };
  return stripeState(view, plan, now) ?? { kind: 'grant', scoped: false };
}

/** Line 2 under the plan line; a grant states its own explanation and has none here. */
export function lineTwoState(state: BillingViewState, plan: GuildPlanView | null): Line2State | null {
  switch (state.kind) {
    case 'trial':
      return { kind: 'trial', cancelBefore: state.firstCharge };
    case 'active':
      return plan?.you_are_billing_contact === true
        ? { kind: 'billed-to-you' }
        : { kind: 'billed-by', name: plan?.billed_by ?? '' };
    case 'ending':
    case 'ending-transfer':
      return { kind: 'ending' };
    case 'lapsed':
      return { kind: 'read-only' };
    default:
      return null;
  }
}

/** Whether the reader is a payer: the claim holder, who gets the buttons. */
export function canBill(view: GuildBillingSettings): boolean {
  return view.can.includes('billing');
}

/** The chooser draws for no plan, a lapsed plan and the scoped officer-tools grant. */
export function showsChooser(state: BillingViewState): boolean {
  return state.kind === 'none' || state.kind === 'lapsed' || (state.kind === 'grant' && state.scoped);
}

/** Guild tier with more verified members than it covers (the leader counted). */
export function showSizeLine(view: GuildBillingSettings): boolean {
  const plan = view.billing;
  return (
    plan !== null &&
    plan.scope === 'plan' &&
    plan.tier === 'guild' &&
    !ENDED_STATUSES.includes(plan.status) &&
    view.verified_count > GUILD_TIER_MEMBER_LIMIT
  );
}

/** Where the `Move to Raid guild` button draws: on the plan line, or on the size line when that shows. */
export function movesFromPlanLine(view: GuildBillingSettings, state: BillingViewState): boolean {
  return (
    canBill(view) &&
    view.billing?.tier === 'guild' &&
    (state.kind === 'active' || state.kind === 'trial') &&
    !showSizeLine(view)
  );
}

/** The plan's seat count (Guild 1, Raid guild 3); without a plan, the rows the read sent. */
export function planSeatCount(view: GuildBillingSettings): number {
  return view.billing === null ? view.seats.length : PLANS[view.billing.tier].seats;
}

/** An offer to take over billing waits for this reader: the receiver's block is drawn. */
export function offeredToReader(view: GuildBillingSettings): boolean {
  return view.transfer?.to_you === true && view.transfer.state === 'offered';
}

/** The read-only line under Records: every non-payer, except while an offer waits for them. */
export function showsReadOnlyLine(view: GuildBillingSettings, payer: boolean): boolean {
  return !payer && !offeredToReader(view);
}
