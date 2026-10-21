// web/src/lib/guild/settings-view.ts
// The guild settings page in a player's words (lifecycle review round 4, N14 / UX 7 and 8): who
// holds the claim, what the plan is and when it renews, and who is billed. Pure, so the states
// are pinned without a page.
import type { GuildBillingView } from '../account/api';
import { utcDayMonthYear } from '../dates';
import type { ClaimStateView, MemberRef } from './api';

export const settingsWords = {
  officerRank: 'Officers are guild ranks at or above:',
  officerRankHelp:
    'Rank number, 0 is the guild master. This rank and every rank above it can approve members and award loot.',
  releaseOnClaimPage: 'Release on the claim page',
  visibilityHelp: 'Public, Unlisted (link only) or Guild only. Applies to new reports.',
  guildPlan: 'Guild plan',
  noBillingContact: 'No billing contact yet',
} as const;

/** "Claimed by hunthrop#1894" for a claimed guild; null otherwise (a contested claim shows the
 *  frozen notice instead, never both). */
export function claimedByLine(claim: ClaimStateView, claimedBy: MemberRef | null): string | null {
  if (claim.state !== 'claimed') return null;
  const name = claimedBy?.battletag ?? claim.claimed_by_name;
  return name === undefined || name === '' ? 'Claimed' : `Claimed by ${name}`;
}

/**
 * "Guild plan · renews 31 Dec 2026", "Guild plan · ends 31 Dec 2026", or for a plan an operator
 * granted (no Stripe billing contact) "Guild plan · until 31 Dec 2026". The date is the one that
 * was stored, read in UTC: a grant until 2026-12-31 is the 31st for a viewer in any zone.
 */
export function guildPlanLine(billing: GuildBillingView): string {
  const date = billing.current_period_end === null ? '' : utcDayMonthYear(billing.current_period_end);
  if (date === '') return settingsWords.guildPlan;
  const verb = billing.billed_by === '' ? 'until' : billing.cancel_at_period_end ? 'ends' : 'renews';
  return `${settingsWords.guildPlan} · ${verb} ${date}`;
}

/** Who pays: "Billed to you.", "Billed by hunthrop#1894", or "No billing contact yet". */
export function billedLine(billing: GuildBillingView): string {
  if (billing.you_are_billing_contact) return 'Billed to you.';
  return billing.billed_by === '' ? settingsWords.noBillingContact : `Billed by ${billing.billed_by}`;
}

/** Manage billing needs a Stripe customer: the billing contact, and only a plan with one. */
export function canManageBilling(billing: GuildBillingView): boolean {
  return billing.you_are_billing_contact && billing.billed_by !== '';
}
