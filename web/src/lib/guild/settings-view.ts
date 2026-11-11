// web/src/lib/guild/settings-view.ts
// The guild settings page in a player's words (lifecycle review round 4, N14 / UX 7 and 8): who
// holds the claim, what the plan is and when it renews, and who is billed. Pure, so the states
// are pinned without a page.
import type { GuildBillingView } from '../account/api';
import { utcDayMonthYear } from '../dates';
import type { SelectOption } from '../ui/listbox-model';
import type { GuildVisibility } from './api';
import { uiCopy } from '../ui/copy';

/** Report visibility in the settings' words, the order the select lists them. */
export const VISIBILITY_OPTIONS: readonly (SelectOption & { value: GuildVisibility })[] = [
  { value: 'public', label: 'Public' },
  { value: 'unlisted', label: 'Unlisted' },
  { value: 'guild', label: 'Guild only' },
];

/** A visibility's word ("Guild only"); the value itself for one the list does not know. */
export function visibilityWord(value: string): string {
  return VISIBILITY_OPTIONS.find((o) => o.value === value)?.label ?? value;
}

export const settingsWords = {
  officerRank: 'Officer ranks',
  releaseOnClaimPage: 'Release on the claim page',
  guildPlan: 'Guild plan',
  /** The plan row with no date to show (the card's title names the tier). */
  planActive: 'Active',
  officerTools: 'Officer tools',
  noBillingContact: 'No billing contact yet',
  grantedByUs: 'Granted by Forever Sixty.',
  billingUnavailable: `Billing is not available right now. ${uiCopy.retry} in a few minutes.`,
  tryAgain: uiCopy.retry,
  guildSettingsLink: 'Guild settings',
} as const;

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
