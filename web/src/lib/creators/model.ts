// web/src/lib/creators/model.ts
// What one player is worth to a creator in their first year, from the program's terms
// (PROGRAM_TERMS) and the plan table (lib/billing/plans): the bounty once, then the share of every
// payment for the share's months. Before Stripe's fees, which the site does not know.
import {
  FOUNDER_LAST_DAY,
  PLANS,
  entryPrice,
  formatAmount,
  formatPrice,
  type Interval,
  type PlanKey,
} from '../billing/plans';
import { PROGRAM_TERMS } from './copy';

const MONTHS_PER_YEAR = 12;
const PERCENT = 100;

/** The plans a worked example shows: a player's Premium both ways, and a guild's officer on each guild plan. */
export const EXAMPLE_PLANS: readonly { plan: PlanKey; interval: Interval }[] = [
  { plan: 'premium', interval: 'yearly' },
  { plan: 'premium', interval: 'monthly' },
  { plan: 'guild', interval: 'yearly' },
  { plan: 'raid_guild', interval: 'yearly' },
];

/** Dollars a creator earns in a player's first year on `plan`, paid per `interval`. */
export function firstYearEarnings(plan: PlanKey, interval: Interval, founder: boolean): number {
  const price = entryPrice(PLANS[plan], interval, founder);
  const paymentsInShare =
    interval === 'monthly' ? PROGRAM_TERMS.shareMonths : PROGRAM_TERMS.shareMonths / MONTHS_PER_YEAR;
  return PROGRAM_TERMS.bountyUsd + (price * paymentsInShare * PROGRAM_TERMS.sharePercent) / PERCENT;
}

export interface EarningsRow {
  readonly id: string;
  /** `Premium, $40 a year`. */
  readonly label: string;
  /** `$22`. */
  readonly value: string;
  /** While founder prices last, the same player at list price: `$27 after 6 Jan 2027`. */
  readonly sub?: string;
}

/** The worked examples at today's prices (founder prices while the window is open). */
export function earningsRows(founder: boolean): readonly EarningsRow[] {
  return EXAMPLE_PLANS.map(({ plan, interval }) => ({
    id: `${plan}-${interval}`,
    label: `${PLANS[plan].name}, ${formatPrice(entryPrice(PLANS[plan], interval, founder), interval)}`,
    value: formatAmount(firstYearEarnings(plan, interval, founder)),
    ...(founder
      ? { sub: `${formatAmount(firstYearEarnings(plan, interval, false))} after ${FOUNDER_LAST_DAY}` }
      : {}),
  }));
}
