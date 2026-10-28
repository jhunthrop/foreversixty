// web/src/lib/account/guild-row-line.ts
// What an account page guild row says beyond its name and plan line (design/specs/
// 2026-10-27-subscribe-and-join.md 6.5): whether the viewer is verified there, and, for a guild
// with no plan, the way to its checkout step. Pure; the component draws it.
import { checkoutHref } from '../../components/premium/model';
import { accountPlanCopy } from '../billing/account-plan-copy';
import { fittingGuildPlan } from '../billing/checkout-params';
import type { GuildCoveredView, GuildSubscribeView } from '../billing/contract';

const copy = accountPlanCopy.guildRow;

/** The fields of a `MeGuild` the row reads. */
export interface GuildRowFacts {
  id: number;
  name: string;
  verified: boolean;
  verified_count: number;
  covered?: GuildCoveredView | null;
  subscribe?: Pick<GuildSubscribeView, 'step'> | null;
}

export interface GuildRowNextStep {
  text: string;
  action: string;
  label: string;
  href: string;
}

export function notVerifiedMark(guild: GuildRowFacts): string | null {
  return guild.verified ? null : copy.notVerified;
}

/** The no-plan line and its link; null on a plan or from an API that sends no step (today's row). */
export function guildRowNextStep(guild: GuildRowFacts): GuildRowNextStep | null {
  const step = guild.subscribe?.step;
  if (step === undefined || step === 'on_plan' || (guild.covered ?? null) !== null) return null;
  return {
    text: copy.noPlan(guild.name),
    action: copy.nextStep,
    label: copy.nextStepLabel(copy.nextStep, guild.name),
    href: `${checkoutHref(fittingGuildPlan(guild.verified_count), 'yearly')}#guild-${guild.id}`,
  };
}
