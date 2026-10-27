// web/src/lib/guild/settings-view.ts
// The guild settings page in a player's words (lifecycle review round 4, N14 / UX 7 and 8): who
// holds the claim, what the plan is and when it renews, and who is billed. Pure, so the states
// are pinned without a page.
import type { GuildBillingView } from '../account/api';
import { utcDayMonthYear } from '../dates';
import type { ClaimStateView, MemberRef } from './api';

export const settingsWords = {
  officerRank: 'Officers are guild ranks at or above',
  officerRankHelp:
    'Rank number, 0 is the guild master. This rank and every rank above it can approve members and award loot.',
  releaseOnClaimPage: 'Release on the claim page',
  visibilityHelp: 'Public, Unlisted (link only) or Guild only. Applies to new reports.',
  guildPlan: 'Guild plan',
  officerTools: 'Officer tools',
  noBillingContact: 'No billing contact yet',
  grantedByUs: 'Granted by Forever Sixty.',
  billingUnavailable: 'Billing is not available right now. Try again in a few minutes.',
  tryAgain: 'Try again',
  guildSettingsLink: 'Guild settings',
} as const;

/** "Claimed by hunthrop#1894" for a claimed guild; null otherwise (a contested claim shows the
 *  frozen notice instead, never both). */
export function claimedByLine(claim: ClaimStateView, claimedBy: MemberRef | null): string | null {
  if (claim.state !== 'claimed') return null;
  const name = claimedBy?.battletag ?? claim.claimed_by_name;
  return name === undefined || name === '' ? 'Claimed' : `Claimed by ${name}`;
}

function planEndDate(billing: Pick<GuildBillingView, 'current_period_end'>): string {
  return billing.current_period_end === null ? '' : utcDayMonthYear(billing.current_period_end);
}

/** What a scoped grant covers and what follows it, in the grant's own words. */
export function officerToolsExplanation(billing: Pick<GuildBillingView, 'current_period_end'>): string {
  const date = planEndDate(billing);
  const until = date === '' ? '' : `, until ${date}`;
  return `Granted by Forever Sixty for the first raid nights: Award, Undo and Nudge only${until}. After that they need the guild plan.`;
}

/** Manage billing needs a Stripe customer: the billing contact, and only a plan with one. */
export function canManageBilling(
  billing: Pick<GuildBillingView, 'you_are_billing_contact' | 'billed_by'>,
): boolean {
  return billing.you_are_billing_contact && billing.billed_by !== '';
}
