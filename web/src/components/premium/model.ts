// web/src/components/premium/model.ts
// What a priced plan card needs from the plan table and the page copy, in one place, so the
// card, its price block and its buttons cannot disagree. No price literal: amounts come from
// lib/billing/plans.
import { listPrice, priceFor } from '../../lib/billing/plan-costs';
import {
  perMemberLine,
  perRaiderLine,
  plansPageCopy,
  type PremiumFeature,
} from '../../lib/billing/plans-page-copy';
import {
  GUILD_TRIAL_DAYS,
  formatAmount,
  formatPrice,
  type Interval,
  type PlanKey,
} from '../../lib/billing/plans';

export const FOUNDER_STATES = [true, false] as const;

/** Direct children of a plan card, each one row of its desktop subgrid (`grid-row: span 6` in
 *  PlanCard.astro). Neighbours sharing that row set carry `lg:row-span-6` in premium.astro. */
export const PLAN_CARD_ROWS = 6;

interface PlanBlock {
  name: string;
  audience: string;
  features: readonly string[] | readonly PremiumFeature[];
}

/** A card's feature lines with their anchors; a plain string line has none. */
export function featureLines(features: readonly (string | PremiumFeature)[]): readonly PremiumFeature[] {
  return features.map((line) => (typeof line === 'string' ? { text: line, anchor: null } : line));
}

/** The line that carries an anchor's `id`: the first line of that anchor, so one id per anchor. */
export function isFirstOfAnchor(lines: readonly PremiumFeature[], index: number): boolean {
  const anchor = lines[index]?.anchor ?? null;
  return anchor !== null && lines.findIndex((line) => line.anchor === anchor) === index;
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
  suffix: string;
  secondLine: string;
  founderLine: string | null;
}

export function priceCell(plan: PlanKey, interval: Interval, founder: boolean): PriceCell {
  const { price } = plansPageCopy;
  const yearly = priceFor(plan, 'yearly', founder);
  return {
    amount: formatAmount(priceFor(plan, interval, founder)),
    suffix: price.suffix[interval],
    secondLine:
      interval === 'yearly' ? price.perMonth(yearly) : price.orYearly(formatPrice(yearly, 'yearly')),
    founderLine: founder ? price.founderUntil(listPrice(plan, interval)) : null,
  };
}

/** The per-member (Guild) or per-raider (Raid guild) cost line; Premium has none. */
export function costLine(plan: PlanKey, interval: Interval, founder: boolean): string | null {
  if (plan === 'guild') return perMemberLine(interval, founder);
  if (plan === 'raid_guild') return perRaiderLine(interval, founder);
  return null;
}

/** The trial line split so its lead can be set strong without a second copy of the words. */
export function trialParts(): { lead: string; rest: string } {
  const lead = `First ${GUILD_TRIAL_DAYS} days free`;
  return { lead, rest: plansPageCopy.trial.slice(lead.length) };
}
