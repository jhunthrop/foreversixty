// web/src/lib/billing/plans.ts
// The one place a Forever Sixty price literal is written. It mirrors the API catalog
// (api plans table); every plan-facing copy string and checkout link builds its price text
// from here.
export type PlanKey = 'premium' | 'guild' | 'raid_guild';
export type Interval = 'monthly' | 'yearly';

/** Display order: yearly first. */
export const INTERVALS: readonly Interval[] = ['yearly', 'monthly'];
/** Display order of the plan cards. */
export const PLAN_ORDER: readonly PlanKey[] = ['premium', 'guild', 'raid_guild'];

export interface Plan {
  key: PlanKey;
  name: 'Premium' | 'Guild' | 'Raid guild';
  monthly: number;
  yearly: number;
  founderMonthly: number;
  founderYearly: number;
  /** True for the guild plans: one subscription covers a guild, not just the buyer. */
  guild: boolean;
  /** Premium seats the subscriber can give to guild members. */
  seats: 0 | 1 | 3;
}

export const PLANS: Record<PlanKey, Plan> = {
  premium: {
    key: 'premium',
    name: 'Premium',
    monthly: 6,
    yearly: 50,
    founderMonthly: 5,
    founderYearly: 40,
    guild: false,
    seats: 0,
  },
  guild: {
    key: 'guild',
    name: 'Guild',
    monthly: 15,
    yearly: 150,
    founderMonthly: 12,
    founderYearly: 120,
    guild: true,
    seats: 1,
  },
  raid_guild: {
    key: 'raid_guild',
    name: 'Raid guild',
    monthly: 25,
    yearly: 250,
    founderMonthly: 20,
    founderYearly: 200,
    guild: true,
    seats: 3,
  },
};

/** The founder window closes at 6 January 23:59:59 America/Los_Angeles; it has no opening instant. */
export const FOUNDER_WINDOW_CLOSES = '2027-01-07T08:00:00Z';
/** The last founder day, one date style for every surface: `6 Jan 2027`. */
export const FOUNDER_LAST_DAY = new Date(Date.parse(FOUNDER_WINDOW_CLOSES) - 1).toLocaleDateString('en-GB', {
  day: 'numeric',
  month: 'short',
  year: 'numeric',
  timeZone: 'America/Los_Angeles',
});
export const GUILD_TIER_MEMBER_LIMIT = 25;
export const GUILD_TRIAL_DAYS = 7;
/** Days before a guild trial ends that the reminder goes out. */
export const TRIAL_REMINDER_DAYS = 3;
/** Raid sizes a Raid guild price is divided across, largest first. */
export const RAID_SIZES = [40, 20, 10] as const;
export type RaidSize = (typeof RAID_SIZES)[number];
export const SEAT_MOVE_LOCK_DAYS = 30;

const MONTHS_PER_YEAR = 12;
const CENTS_PER_DOLLAR = 100;
const FOUNDER_WINDOW_CLOSES_MS = Date.parse(FOUNDER_WINDOW_CLOSES);

export function founderOpen(now: Date): boolean {
  return now.getTime() < FOUNDER_WINDOW_CLOSES_MS;
}

/** The price a plan entry charges for an interval; pure, so callers can pass any table's entry. */
export function entryPrice(entry: Plan, interval: Interval, founder: boolean): number {
  if (interval === 'monthly') return founder ? entry.founderMonthly : entry.monthly;
  return founder ? entry.founderYearly : entry.yearly;
}

export function priceFor(plan: PlanKey, interval: Interval, founder: boolean): number {
  return entryPrice(PLANS[plan], interval, founder);
}

/** `$3.33`, `$10`, `$12.50`: cents only when the amount is not whole. */
export function formatAmount(usd: number): string {
  const cents = Math.round(usd * CENTS_PER_DOLLAR);
  const dollars = cents / CENTS_PER_DOLLAR;
  return cents % CENTS_PER_DOLLAR === 0 ? `$${dollars}` : `$${dollars.toFixed(2)}`;
}

/** The words after a price figure: ` a year`, ` a month` (leading space included). */
export function intervalUnit(interval: Interval): string {
  return ` a ${interval === 'monthly' ? 'month' : 'year'}`;
}

/** Running-text price: `$40 a year`, `$5 a month`. */
export function formatPrice(usd: number, interval: Interval): string {
  return `${formatAmount(usd)}${intervalUnit(interval)}`;
}

/** The list price shown beside a founder price: `$50 after 6 Jan 2027`, `$6 a month after 6 Jan 2027`. */
export function listPriceAfterFounder(plan: PlanKey, interval: Interval): string {
  const amount = formatAmount(priceFor(plan, interval, false));
  return `${amount}${interval === 'monthly' ? ' a month' : ''} after ${FOUNDER_LAST_DAY}`;
}

export function perMonthLine(yearlyUsd: number): string {
  return `${formatAmount(yearlyUsd / MONTHS_PER_YEAR)} a month, billed yearly`;
}

export function formatCents(cents: number, interval: Interval): string {
  return formatPrice(cents / CENTS_PER_DOLLAR, interval);
}
