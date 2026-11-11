// web/src/lib/billing/guild-billing-copy.ts
// The guild settings #billing strings of design/specs/2026-10-27-launch-pricing.md section 6.2:
// Plan line, Line 2, Chooser, Size line, Upgrade confirm, Records and Gate. Prices and limits come
// from ./plans, never written here. Dates arrive as stored: a date-only value reads as written, an instant on the viewer calendar.
import type { GuildBillingView } from '../account/api';
import { zonedDayMonth, zonedDayMonthYear } from '../dates';
import { settingsWords } from '../guild/settings-view';
import type { GuildTier } from './contract';
import { counted } from '../plural';
import { priceFor } from './plan-costs';
import {
  GUILD_TIER_MEMBER_LIMIT,
  GUILD_TRIAL_DAYS,
  PLANS,
  formatUsd,
  formatCents,
  formatPrice,
  type Interval,
} from './plans';
import { plansPageCopy } from './plans-page-copy';
import { withRetry } from '../ui/copy';

const CENTS_PER_DOLLAR = 100;
const EXTRA_RAID_SEATS = PLANS.raid_guild.seats - PLANS.guild.seats;

export type GuildPlanLineState =
  | { kind: 'trial'; tier: GuildTier; firstCharge: string; amountCents: number; interval: Interval }
  | { kind: 'active'; tier: GuildTier; renews: string; amountCents: number; interval: Interval }
  | { kind: 'ending'; tier: GuildTier; ends: string }
  | { kind: 'ending-transfer'; tier: GuildTier; ends: string; toName: string }
  | { kind: 'lapsed'; ended: string };

export type Line2State =
  | { kind: 'trial'; cancelBefore: string }
  | { kind: 'billed-to-you' }
  | { kind: 'billed-by'; name: string }
  | { kind: 'ending' }
  | { kind: 'read-only' };

/** The facts a plan line reads: shared by the account's `me` view and the guild billing read. */
type PlanLineFacts = Pick<
  GuildBillingView,
  'scope' | 'billed_by' | 'current_period_end' | 'cancel_at_period_end' | 'you_are_billing_contact'
>;

const tierName = (tier: GuildTier): string => PLANS[tier].name;

/** The plan line under the card's title, which names the tier: the line never repeats it (guild
 *  loop round 3, web4, after `Raid guild` over `Guild plan`). */
function planLine(state: GuildPlanLineState, zone?: string): string {
  switch (state.kind) {
    case 'trial':
      return `Trial · first charge ${zonedDayMonthYear(state.firstCharge, zone)}: ${formatUsd(state.amountCents / CENTS_PER_DOLLAR)}`;
    case 'active':
      return `Renews ${zonedDayMonthYear(state.renews, zone)} · ${formatCents(state.amountCents, state.interval)}`;
    case 'ending':
      return `Ends ${zonedDayMonthYear(state.ends, zone)}`;
    case 'ending-transfer':
      return `Ends ${zonedDayMonthYear(state.ends, zone)} · then billed by ${state.toName}`;
    case 'lapsed':
      return `Guild plan ended ${zonedDayMonthYear(state.ended, zone)}.`;
  }
}

function line2(state: Line2State, zone?: string): string {
  switch (state.kind) {
    case 'trial':
      return `Cancel before ${zonedDayMonth(state.cancelBefore, zone)} and nothing is charged. Billed to you.`;
    case 'billed-to-you':
      return 'Billed to you.';
    case 'billed-by':
      return `Billed by ${state.name}`;
    case 'ending':
      return 'Officer tools turn read-only then. Records stay, and CSV export stays.';
    case 'read-only':
      return 'Officer tools are read-only. Raids, attendance, loot and reports stay readable, and the export below always works.';
  }
}

const UPGRADE_LEAD_IN = 'Move to Raid guild? Today you pay ';

/** The prorated sentence in three pieces, so the amount can be a placeholder while the preview loads. */
function upgradeParts(
  proratedCents: number,
  interval: Interval,
  from: string,
  nextCents: number,
  zone?: string,
): { before: string; amount: string; after: string } {
  const period = interval === 'yearly' ? 'year' : 'month';
  return {
    before: UPGRADE_LEAD_IN,
    amount: formatUsd(proratedCents / CENTS_PER_DOLLAR),
    after: ` for the rest of this ${period}. From ${zonedDayMonthYear(from, zone)} it is ${formatCents(nextCents, interval)}.`,
  };
}

/** The prorated sentence: the rest of "this year" or "this month", then the next price. */
function upgradeSentence(
  proratedCents: number,
  interval: Interval,
  from: string,
  nextCents: number,
  zone?: string,
): string {
  const { before, amount, after } = upgradeParts(proratedCents, interval, from, nextCents, zone);
  return `${before}${amount}${after}`;
}

const tierMemberLine = (tier: GuildTier): string => {
  const seats = PLANS[tier].seats;
  const seatWord = seats === 1 ? 'seat' : 'seats';
  return tier === 'guild'
    ? `Up to ${GUILD_TIER_MEMBER_LIMIT} verified members · ${seats} Premium ${seatWord}`
    : `Any size · loot council and readiness · ${seats} Premium ${seatWord}`;
};

export const guildBillingCopy = {
  loadFailed: withRetry('Billing did not load.'),
  planLine,
  founderPill: 'Founder',
  line2,
  chooser: {
    /** The plan radio group's accessible name. */
    planGroup: 'Plan',
    /** The one link to the plans page's guild region (spec 2026-10-29-premium-entry section 7). */
    compare: 'Compare Guild and Raid guild',
    tierName,
    tierLine: tierMemberLine,
    verifiedToday: (count: number): string => `${counted(count, 'verified member')} today.`,
    startTrial: `Start the ${GUILD_TRIAL_DAYS}-day trial`,
    subscribe: (tier: GuildTier, interval: Interval, founder: boolean): string => {
      return `Subscribe · ${formatPrice(priceFor(tier, interval, founder), interval)}`;
    },
    subscribeAgain: 'Subscribe again',
    trialLine: plansPageCopy.trial,
    founderLine: 'Founder price until 6 January 2027, kept while the plan stays active.',
    opening: 'Opening checkout…',
    failed: withRetry('Checkout did not open.'),
  },
  size: {
    line: (count: number): string =>
      `${count} verified members. Guild covers up to ${GUILD_TIER_MEMBER_LIMIT}; nothing switches off. Raid guild covers any size.`,
    move: 'Move to Raid guild',
  },
  upgrade: {
    sentence: upgradeSentence,
    parts: upgradeParts,
    leadIn: UPGRADE_LEAD_IN,
    // Controller's 2026-10-27 persona-round wording (spec 6.2, Upgrade confirm).
    unconfirmed: 'Moved. Stripe is still confirming the charge.',
    unconfirmedStale: 'Stripe is still confirming the move. Reload in a minute to see the plan.',
    previewExpired: "Today's amount was updated. Check it, then move.",
    commit: 'Move to Raid guild',
    committing: 'Moving…',
    keep: 'Keep Guild',
    done: `Moved. Loot council, readiness and ${EXTRA_RAID_SEATS} more seats are on now.`,
    failed: withRetry('That did not go through. Nothing was charged.'),
  },
  records: {
    heading: 'Records',
    body: "Your guild's raids, attendance and loot are kept for good.",
    attendance: 'Download attendance (CSV)',
    loot: 'Download loot (CSV)',
    preparing: 'Preparing…',
    failed: withRetry('The file did not download.'),
    fileName: (guildSlug: string, kind: 'attendance' | 'loot'): string => `${guildSlug}-${kind}.csv`,
  },
  /** An officer who may not bill (spec 6.2, Read-only officer; controller's 2026-10-27 persona round). */
  readOnly: 'The subscriber or the guild leader manages billing and seats.',
  gate: {
    text: 'Loot council and the readiness board are part of Raid guild. ',
    action: 'Move to Raid guild',
  },
} as const;

/**
 * The plan row under the card's title (which names the tier, so the row never repeats it; guild
 * loop round 3, web4): "Renews 31 Dec 2026", "Ends 31 Dec 2026", for a plan an operator granted
 * (no Stripe billing contact) "Until 31 Dec 2026", and "Active" with no date. The scoped grant
 * keeps its own name: "Officer tools · until 31 Dec 2026". The date is the one that was stored: a
 * Stripe period end is an instant read on the viewer calendar; an operator grant (no billing
 * contact, or the scoped officer tools) is a calendar day the API stores at midnight UTC, so it
 * reads in UTC and is the 31st for a viewer in any zone.
 */
export function guildPlanLine(billing: PlanLineFacts, zone?: string): string {
  const isGrant = billing.scope === 'officer_tools' || billing.billed_by === '';
  const date =
    billing.current_period_end === null
      ? ''
      : zonedDayMonthYear(billing.current_period_end, isGrant ? 'UTC' : zone);
  if (billing.scope === 'officer_tools') {
    return date === '' ? settingsWords.officerTools : `${settingsWords.officerTools} · until ${date}`;
  }
  if (date === '') return settingsWords.planActive;
  const verb = billing.billed_by === '' ? 'Until' : billing.cancel_at_period_end ? 'Ends' : 'Renews';
  return `${verb} ${date}`;
}

/** Who pays: "Billed to you.", "Billed by hunthrop#1894", or "Granted by Forever Sixty." */
export function billedLine(billing: Pick<PlanLineFacts, 'billed_by' | 'you_are_billing_contact'>): string {
  if (billing.you_are_billing_contact) return line2({ kind: 'billed-to-you' });
  return billing.billed_by === ''
    ? settingsWords.grantedByUs
    : line2({ kind: 'billed-by', name: billing.billed_by });
}
