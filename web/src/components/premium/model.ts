// web/src/components/premium/model.ts
// What a priced plan card needs from the plan table and the page copy, in one place, so the
// card, its price block and its buttons cannot disagree. No price literal: amounts come from
// lib/billing/plans.
import { plansPageCopy } from '../../lib/billing/plans-page-copy';
import {
  GUILD_TRIAL_DAYS,
  formatAmount,
  formatPrice,
  listPriceAfterFounder,
  priceFor,
  type Interval,
  type PlanKey,
} from '../../lib/billing/plans';

export const FOUNDER_STATES = [true, false] as const;

interface PlanBlock {
  name: string;
  audience: string;
  features: readonly string[];
}

const BLOCKS: Record<PlanKey, PlanBlock> = {
  premium: plansPageCopy.premium,
  guild: plansPageCopy.guild,
  raid_guild: plansPageCopy.raidGuild,
};

export function planBlock(plan: PlanKey): PlanBlock {
  return BLOCKS[plan];
}

/** `data-fd` value for a founder state: the attribute the page's CSS switches on. */
export function founderKey(founder: boolean): 'open' | 'closed' {
  return founder ? 'open' : 'closed';
}

export function checkoutHref(plan: PlanKey, interval: Interval): string {
  return `/premium/checkout?plan=${plan}&interval=${interval}`;
}

export function buttonLabel(plan: PlanKey, interval: Interval, founder: boolean): string {
  const price = formatPrice(priceFor(plan, interval, founder), interval);
  const { subscribe, subscribeGuild } = plansPageCopy.buttons;
  return plan === 'premium' ? subscribe(price) : subscribeGuild(price);
}

export interface PriceCell {
  amount: string;
  listLine: string | null;
  suffix: string;
  secondLine: string;
  founderLine: string | null;
}

export function priceCell(plan: PlanKey, interval: Interval, founder: boolean): PriceCell {
  const { price } = plansPageCopy;
  const yearly = priceFor(plan, 'yearly', founder);
  return {
    amount: formatAmount(priceFor(plan, interval, founder)),
    listLine: founder ? listPriceAfterFounder(plan, interval) : null,
    suffix: price.suffix[interval],
    secondLine:
      interval === 'yearly' ? price.perMonth(yearly) : price.orYearly(formatPrice(yearly, 'yearly')),
    founderLine: founder ? price.founder : null,
  };
}

/** The trial line split so its lead can be set strong without a second copy of the words. */
export function trialParts(): { lead: string; rest: string } {
  const lead = `First ${GUILD_TRIAL_DAYS} days free`;
  return { lead, rest: plansPageCopy.trial.slice(lead.length) };
}
