// web/src/lib/billing/premium-page.ts
// What /premium shows a signed-in account (premium spec 2026-11-08 sections 4.1, 4.3, 4.5; gap G1,
// G2): the head answer (the plan it is on now), at most one state line (a failed payment, a seat
// that ended), and the guild column a billing officer's guild is on. Pure: the island reads /v1/me
// and draws these; signed out none of it runs.
import type { Me } from '../account/api';
import { accountPlanLine, premiumSource } from '../account/plan-line';
import { checkoutGuildCopy, type GuildTier } from './checkout-guild-copy';

export interface PremiumAnswerView {
  name: string;
  /** `Renews 4 Nov 2027 · $40 a year`, `A seat from Olympus XXVII`; empty when there is nothing. */
  sub: string;
  founder: boolean;
}

/** The head answer: the plan's name, its date and price or source, and the Founder mark. */
export function premiumAnswerView(me: Me, now: Date, zone?: string): PremiumAnswerView {
  const { head, founder } = accountPlanLine(me, now, zone);
  const when = head.when === null ? '' : `${head.when.label} ${head.when.date}`;
  return { name: head.name, sub: [when, head.extra ?? ''].filter((x) => x !== '').join(' · '), founder };
}

export type PremiumStateLine =
  { kind: 'pastDue' } | { kind: 'ended'; text: string; action: string; href: string } | null;

/** The one state line under the bar: a failed payment first, else a seat that ended. */
export function premiumStateLine(me: Me, now: Date, zone?: string): PremiumStateLine {
  if (me.entitlements?.billing?.status === 'past_due') return { kind: 'pastDue' };
  const { ended } = accountPlanLine(me, now, zone);
  return ended === null ? null : { kind: 'ended', ...ended };
}

/** Whose plan column the page marks: `free` while the account has no Premium, else `owned`. */
export function premiumMine(me: Me, now: Date): 'free' | 'owned' {
  return premiumSource(me, now) === 'none' ? 'free' : 'owned';
}

export interface GuildOnPlanView {
  guildId: number;
  guildName: string;
  line: string;
  /** The account pays: Manage billing opens the guild's portal. */
  manage: boolean;
}

/** A guild of this account on `tier` (scope plan) that the account pays for or bills as an officer,
 *  with the sentence its column reads; a member's guild is not marked (spec 4.5: an officer). */
export function guildOnPlan(me: Me, tier: GuildTier): GuildOnPlanView | null {
  const guild = me.guilds.find(
    (g) =>
      g.plan?.tier === tier &&
      g.plan.scope === 'plan' &&
      (g.plan.you_are_billing_contact || (g.can ?? []).includes('billing')),
  );
  const plan = guild?.plan;
  if (guild === undefined || plan == null) return null;
  const { onPlan } = checkoutGuildCopy;
  const line = plan.you_are_billing_contact
    ? onPlan.youPay(guild.name, tier)
    : plan.billed_by !== ''
      ? onPlan.billedBy(guild.name, tier, plan.billed_by)
      : onPlan.unnamed(guild.name, tier);
  return { guildId: guild.id, guildName: guild.name, line, manage: plan.you_are_billing_contact };
}
