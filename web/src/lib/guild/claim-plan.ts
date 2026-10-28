// web/src/lib/guild/claim-plan.ts
// The plan a claim link was made for (design/specs/2026-10-27-subscribe-and-join.md §6.2): read
// from the page's own query, drawn as one line, handed on in the member's copy link, and turned
// into the `ready` button's label once the claim is held. A value outside the closed sets is
// dropped, so a tampered `plan=premium` reads as no plan at all.
import type { CharacterPath } from '../characters';
import { checkoutGuildCopy } from '../billing/checkout-guild-copy';
import type { GuildTier } from '../billing/contract';
import { priceFor } from '../billing/plan-costs';
import { parseInterval, parsePlan } from '../billing/plan-params';
import { PLANS, formatPrice, type Interval } from '../billing/plans';
import { handoffUrl } from '../billing/subscribe-step';
import { guildClaimCopy } from './copy';

export interface ClaimPlan {
  plan: GuildTier;
  interval: Interval;
}

const DEFAULT_INTERVAL: Interval = 'yearly';

/** The guild plan and interval a claim page was opened with, or null when `plan` is absent or not a guild plan. */
export function parseClaimPlan(search: string): ClaimPlan | null {
  const params = new URLSearchParams(search);
  const plan = parsePlan(params.get('plan'));
  if (plan !== 'guild' && plan !== 'raid_guild') return null;
  return { plan, interval: parseInterval(params.get('interval')) ?? DEFAULT_INTERVAL };
}

function priceText({ plan, interval }: ClaimPlan, founder: boolean): string {
  return formatPrice(priceFor(plan, interval, founder), interval);
}

/** The line under the heading when the link names a plan. */
export function claimPlanLine(guildName: string, claimPlan: ClaimPlan, founder: boolean): string {
  return guildClaimCopy.planLine(guildName, PLANS[claimPlan.plan].name, priceText(claimPlan, founder));
}

/** The claim page's own link, carrying the plan it was opened with, for a member to send the guild master. */
export function claimLinkFor(path: CharacterPath, claimPlan: ClaimPlan | null, origin: string): string {
  return handoffUrl(path, 'ask_leader', claimPlan?.plan ?? '', claimPlan?.interval ?? '', origin);
}

/** The `ready` button's label: the trial while it is unused, else the price. */
export function subscribeLabel(claimPlan: ClaimPlan, trialUsed: boolean | null, founder: boolean): string {
  return trialUsed === false
    ? checkoutGuildCopy.ready.startTrial
    : checkoutGuildCopy.ready.subscribe(priceText(claimPlan, founder));
}

/** `/setup` with `next` set to a claim page's path and query. */
export function setupHrefFor(setupPath: string, claimPath: string): string {
  return `${setupPath}?next=${encodeURIComponent(claimPath)}`;
}
