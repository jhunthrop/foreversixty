// web/src/lib/billing/plan-costs.ts
// Per-raider and per-member cost amounts. Kept apart from plans.ts so a test that mocks
// ./plans moves every amount computed here.
import { GUILD_TIER_MEMBER_LIMIT, RAID_SIZES, formatAmount, priceFor, type Interval } from './plans';

/** The Raid guild price divided by each raid size, in `RAID_SIZES` order. */
export function perRaiderAmounts(interval: Interval, founder: boolean): readonly string[] {
  const price = priceFor('raid_guild', interval, founder);
  return RAID_SIZES.map((size) => formatAmount(price / size));
}

/** The Guild price divided by the member limit. */
export function perMemberAmount(interval: Interval, founder: boolean): string {
  return formatAmount(priceFor('guild', interval, founder) / GUILD_TIER_MEMBER_LIMIT);
}
