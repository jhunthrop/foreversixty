// web/src/lib/account/guild-row-line.ts
// What an account page guild row says beyond its name and plan line (design/specs/
// 2026-10-27-subscribe-and-join.md 6.5): whether the viewer is verified there, and, for a guild
// with no plan, the way to its checkout step. Pure; the component draws it.
import { checkoutHref } from '../../components/premium/model';
import { guildFitsRaidTier } from '../billing/checkout-params';
import type { GuildSubscribeView } from '../billing/contract';
import type { PlanKey } from '../billing/plans';

export const accountGuildRowCopy = {
  notVerified: 'Not verified',
  noPlan: (guild: string): string => `${guild} has no plan.`,
  nextStep: 'See the next step',
} as const;

/** The fields of a `MeGuild` the row reads. */
export interface GuildRowFacts {
  id: number;
  name: string;
  verified: boolean;
  verified_count: number;
  covered?: unknown;
  subscribe?: Pick<GuildSubscribeView, 'step'> | null;
}

export interface GuildRowNextStep {
  text: string;
  action: string;
  href: string;
}

/** `raid_guild` above the Guild plan's member limit, else `guild` (spec 6.5). */
export function fittingGuildPlan(verifiedCount: number): Extract<PlanKey, 'guild' | 'raid_guild'> {
  return guildFitsRaidTier(verifiedCount) ? 'raid_guild' : 'guild';
}

export function notVerifiedMark(guild: GuildRowFacts): string | null {
  return guild.verified ? null : accountGuildRowCopy.notVerified;
}

/** The no-plan line and its link; null on a plan or from an API that sends no step (today's row). */
export function guildRowNextStep(guild: GuildRowFacts): GuildRowNextStep | null {
  const step = guild.subscribe?.step;
  if (step === undefined || step === 'on_plan' || (guild.covered ?? null) !== null) return null;
  return {
    text: accountGuildRowCopy.noPlan(guild.name),
    action: accountGuildRowCopy.nextStep,
    href: `${checkoutHref(fittingGuildPlan(guild.verified_count), 'yearly')}#guild-${guild.id}`,
  };
}
