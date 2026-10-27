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
export const GUILD_TIER_MEMBER_LIMIT = 25;
export const GUILD_TRIAL_DAYS = 7;

const MONTHS_PER_YEAR = 12;
const CENTS_PER_DOLLAR = 100;
const FOUNDER_WINDOW_CLOSES_MS = Date.parse(FOUNDER_WINDOW_CLOSES);

export function founderOpen(now: Date): boolean {
  return now.getTime() < FOUNDER_WINDOW_CLOSES_MS;
}

export function priceFor(plan: PlanKey, interval: Interval, founder: boolean): number {
  const entry = PLANS[plan];
  if (interval === 'monthly') return founder ? entry.founderMonthly : entry.monthly;
  return founder ? entry.founderYearly : entry.yearly;
}

/** `$3.33`, `$10`, `$12.50`: cents only when the amount is not whole. */
export function formatAmount(usd: number): string {
  const cents = Math.round(usd * CENTS_PER_DOLLAR);
  const dollars = cents / CENTS_PER_DOLLAR;
  return cents % CENTS_PER_DOLLAR === 0 ? `$${dollars}` : `$${dollars.toFixed(2)}`;
}

/** Running-text price: `$40 a year`, `$5 a month`. */
export function formatPrice(usd: number, interval: Interval): string {
  return `${formatAmount(usd)} a ${interval === 'monthly' ? 'month' : 'year'}`;
}

export function perMonthLine(yearlyUsd: number): string {
  return `${formatAmount(yearlyUsd / MONTHS_PER_YEAR)} a month, billed yearly`;
}

export function formatCents(cents: number, interval: Interval): string {
  return formatPrice(cents / CENTS_PER_DOLLAR, interval);
}
