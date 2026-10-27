// web/src/lib/guild/upgrade-facts.ts
// What the Move to Raid guild sentence reads (launch pricing spec 3.2.4). Before the preview lands
// the plan frames it: its interval, its renewal date and the Raid guild price at the same founder
// flag (the lock carries, product spec 8.2), so only the prorated amount waits on a Skeleton. Once
// the preview lands its amounts win; the date stays the plan line's own so the two never disagree.
import type { GuildPlanView, UpgradePreview } from '../billing/contract';
import { priceFor, type Interval } from '../billing/plans';

const CENTS_PER_DOLLAR = 100;

export interface UpgradeFacts {
  /** Null until the preview lands. */
  proratedCents: number | null;
  interval: Interval;
  renewsAt: string;
  nextCents: number;
}

type PlanFacts = Pick<GuildPlanView, 'interval' | 'founder' | 'current_period_end'>;

/** Null when neither the plan nor the preview can frame the sentence. */
export function upgradeFacts(plan: PlanFacts | null, preview: UpgradePreview | null): UpgradeFacts | null {
  const planDate = plan?.current_period_end ?? null;
  if (preview !== null) {
    return {
      proratedCents: preview.amount_due_cents,
      interval: preview.interval,
      renewsAt: planDate ?? preview.renews_at,
      nextCents: preview.next_amount_cents,
    };
  }
  if (plan === null || plan.interval === null || planDate === null) return null;
  return {
    proratedCents: null,
    interval: plan.interval,
    renewsAt: planDate,
    nextCents: priceFor('raid_guild', plan.interval, plan.founder) * CENTS_PER_DOLLAR,
  };
}
