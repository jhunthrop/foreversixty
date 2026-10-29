// web/src/lib/billing/plan-costs.ts
// Every price lookup and the per-raider and per-member amounts built on it. Kept apart from plans.ts so a test that mocks
// ./plans moves every amount computed here: every price is read from the imported PLANS.
import {
  FOUNDER_LAST_DAY,
  GUILD_TIER_MEMBER_LIMIT,
  PLANS,
  RAID_SIZES,
  entryPrice,
  formatAmount,
  type Interval,
  type PlanKey,
} from './plans';

export function priceFor(plan: PlanKey, interval: Interval, founder: boolean): number {
  return entryPrice(PLANS[plan], interval, founder);
}

/** The list price with its interval: `$50`, `$6 a month`. */
export function listPrice(plan: PlanKey, interval: Interval): string {
  return `${formatAmount(priceFor(plan, interval, false))}${interval === 'monthly' ? ' a month' : ''}`;
}

/** The list price shown beside a founder price: `$50 after 6 Jan 2027`, `$6 a month after 6 Jan 2027`. */
export function listPriceAfterFounder(plan: PlanKey, interval: Interval): string {
  return `${listPrice(plan, interval)} after ${FOUNDER_LAST_DAY}`;
}

/** The Raid guild price divided by each raid size, in `RAID_SIZES` order. */
export function perRaiderAmounts(interval: Interval, founder: boolean): readonly string[] {
  const price = priceFor('raid_guild', interval, founder);
  return RAID_SIZES.map((size) => formatAmount(price / size));
}

/** The Guild price divided by the member limit. */
export function perMemberAmount(interval: Interval, founder: boolean): string {
  return formatAmount(priceFor('guild', interval, founder) / GUILD_TIER_MEMBER_LIMIT);
}
