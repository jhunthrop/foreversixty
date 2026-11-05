// web/src/lib/guild/join-plan.ts
// Spec 2026-10-27 §6.3: the invite's joined state names the guild's plan, from `/v1/me` `covered`.
import type { GuildCoveredView } from '../billing/contract';
import { PLANS } from '../billing/plans';
import { guildJoinCopy } from './copy';
import { isOfficerRank } from './roster-model';

/** A guild has a plan while something covers it; a lapsed plan is no plan. */
export function guildHasPlan(covered: GuildCoveredView | null | undefined): covered is GuildCoveredView {
  return covered !== null && covered !== undefined && covered.state !== 'lapsed';
}

/** `{Guild} is on the {Tier} plan.` or `{Guild} is on the free plan.`; a lapsed plan is no plan. */
export function joinedPlanLine(guildName: string, covered: GuildCoveredView | null | undefined): string {
  if (!guildHasPlan(covered)) return guildJoinCopy.noPlanYet(guildName);
  return guildJoinCopy.onPlan(guildName, PLANS[covered.tier].name);
}

/**
 * The joined state's plan sentence for this viewer: an officer or the guild master reads it; a
 * member, who has nothing to do about the plan, gets none (guild loop round 3).
 */
export function joinedPlanLineFor(
  guildName: string,
  entry: { rank?: string; covered?: GuildCoveredView | null },
): string | null {
  return isOfficerRank(entry.rank ?? '') ? joinedPlanLine(guildName, entry.covered) : null;
}
