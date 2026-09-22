// web/src/lib/billing/plans.ts
// The one place a Forever Sixty price literal is written. Every plan-facing copy string
// (billing/copy.ts) and every checkout link builds its price text from here, per the
// coordinator's "single data file, no price literal repeated in copy" instruction.
export type PlanKey = 'premium' | 'guild';
export type Interval = 'monthly' | 'yearly';

export interface Plan {
  key: PlanKey;
  name: string;
  monthly: number;
  yearly: number;
  /** True for the guild plan: one subscription covers every member, not just the buyer. */
  coversGuild: boolean;
}

export const PLANS: Record<PlanKey, Plan> = {
  premium: { key: 'premium', name: 'Premium', monthly: 4, yearly: 40, coversGuild: false },
  guild: { key: 'guild', name: 'Guild', monthly: 15, yearly: 150, coversGuild: true },
};

export function formatPrice(amountUsd: number, interval: Interval): string {
  return `$${amountUsd}/${interval === 'monthly' ? 'month' : 'year'}`;
}
