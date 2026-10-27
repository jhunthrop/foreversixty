// web/src/lib/billing/guild-billing-copy.ts
// The guild settings #billing strings of design/specs/2026-10-27-launch-pricing.md section 6.2:
// Plan line, Line 2, Chooser, Size line, Upgrade confirm, Records and Gate. Prices and limits come
// from ./plans, never written here. Every date arrives as stored and is read in UTC.
import type { GuildBillingView } from '../account/api';
import { utcDayMonth, utcDayMonthYear } from '../dates';
import { settingsWords } from '../guild/settings-view';
import type { GuildTier } from './contract';
import {
  GUILD_TIER_MEMBER_LIMIT,
  GUILD_TRIAL_DAYS,
  PLANS,
  formatAmount,
  formatCents,
  formatPrice,
  priceFor,
  type Interval,
} from './plans';
import { plansPageCopy } from './plans-page-copy';

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

const tierName = (tier: GuildTier): string => PLANS[tier].name;

function planLine(state: GuildPlanLineState): string {
  switch (state.kind) {
    case 'trial':
      return `${tierName(state.tier)} · trial · first charge ${utcDayMonthYear(state.firstCharge)}: ${formatAmount(state.amountCents / CENTS_PER_DOLLAR)}`;
    case 'active':
      return `${tierName(state.tier)} · renews ${utcDayMonthYear(state.renews)} · ${formatCents(state.amountCents, state.interval)}`;
    case 'ending':
      return `${tierName(state.tier)} · ends ${utcDayMonthYear(state.ends)}`;
    case 'ending-transfer':
      return `${tierName(state.tier)} · ends ${utcDayMonthYear(state.ends)} · then billed by ${state.toName}`;
    case 'lapsed':
      return `Guild plan ended ${utcDayMonthYear(state.ended)}.`;
  }
}

function line2(state: Line2State): string {
  switch (state.kind) {
    case 'trial':
      return `Cancel before ${utcDayMonth(state.cancelBefore)} and nothing is charged. Billed to you.`;
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

/** The prorated sentence: the rest of "this year" or "this month", then the next price. */
function upgradeSentence(proratedCents: number, interval: Interval, from: string, nextCents: number): string {
  const period = interval === 'yearly' ? 'year' : 'month';
  return `Move to Raid guild? Today you pay ${formatAmount(proratedCents / CENTS_PER_DOLLAR)} for the rest of this ${period}. From ${utcDayMonthYear(from)} it is ${formatCents(nextCents, interval)}.`;
}

const tierMemberLine = (tier: GuildTier): string => {
  const seats = PLANS[tier].seats;
  const seatWord = seats === 1 ? 'seat' : 'seats';
  return tier === 'guild'
    ? `Up to ${GUILD_TIER_MEMBER_LIMIT} verified members · ${seats} Premium ${seatWord}`
    : `Any size · loot council and readiness · ${seats} Premium ${seatWord}`;
};

export const guildBillingCopy = {
  planLine,
  founderPill: 'Founder',
  line2,
  chooser: {
    tierName,
    tierLine: tierMemberLine,
    verifiedToday: (count: number): string => `${count} verified members today.`,
    startTrial: `Start the ${GUILD_TRIAL_DAYS}-day trial`,
    subscribe: (tier: GuildTier, interval: Interval, founder: boolean): string => {
      return `Subscribe · ${formatPrice(priceFor(tier, interval, founder), interval)}`;
    },
    subscribeAgain: 'Subscribe again',
    trialLine: plansPageCopy.trial,
    founderLine: 'Founder price until 6 January, kept while the plan stays active.',
    opening: 'Opening checkout…',
    failed: 'Checkout did not open. Try again.',
  },
  size: {
    line: (count: number): string =>
      `${count} verified members. Guild covers up to ${GUILD_TIER_MEMBER_LIMIT}; nothing switches off. Raid guild covers any size.`,
    move: 'Move to Raid guild',
  },
  upgrade: {
    sentence: upgradeSentence,
    commit: 'Move to Raid guild',
    committing: 'Moving…',
    keep: 'Keep Guild',
    done: `Moved. Loot council, readiness and ${EXTRA_RAID_SEATS} more seats are on now.`,
    failed: 'That did not go through. Nothing was charged. Try again.',
  },
  records: {
    heading: 'Records',
    body: "Your guild's raids, attendance and loot are kept for good.",
    attendance: 'Download attendance (CSV)',
    loot: 'Download loot (CSV)',
    preparing: 'Preparing…',
    failed: 'The file did not download. Try again.',
    fileName: (guildSlug: string, kind: 'attendance' | 'loot'): string => `${guildSlug}-${kind}.csv`,
  },
  gate: {
    text: 'Loot council and the readiness board are part of Raid guild. ',
    action: 'Move to Raid guild',
  },
} as const;

/**
 * "Guild plan · renews 31 Dec 2026", "Guild plan · ends 31 Dec 2026", or for a plan an operator
 * granted (no Stripe billing contact) "Guild plan · until 31 Dec 2026". The date is the one that
 * was stored, read in UTC: a grant until 2026-12-31 is the 31st for a viewer in any zone.
 */
export function guildPlanLine(billing: GuildBillingView): string {
  const date = billing.current_period_end === null ? '' : utcDayMonthYear(billing.current_period_end);
  if (billing.scope === 'officer_tools') {
    return date === '' ? settingsWords.officerTools : `${settingsWords.officerTools} · until ${date}`;
  }
  if (date === '') return settingsWords.guildPlan;
  const verb = billing.billed_by === '' ? 'until' : billing.cancel_at_period_end ? 'ends' : 'renews';
  return `${settingsWords.guildPlan} · ${verb} ${date}`;
}

/** Who pays: "Billed to you.", "Billed by hunthrop#1894", or "Granted by Forever Sixty." */
export function billedLine(billing: GuildBillingView): string {
  if (billing.you_are_billing_contact) return line2({ kind: 'billed-to-you' });
  return billing.billed_by === ''
    ? settingsWords.grantedByUs
    : line2({ kind: 'billed-by', name: billing.billed_by });
}
