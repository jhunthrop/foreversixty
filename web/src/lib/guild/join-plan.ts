// web/src/lib/guild/join-plan.ts
// Spec 2026-10-27 §6.3: the invite's joined state names the guild's plan, from `/v1/me` `covered`.
import type { GuildCoveredView } from '../billing/contract';
import { PLANS } from '../billing/plans';
import { guildJoinCopy } from './copy';

/** `{Guild} is on the {Tier} plan.` or `{Guild} has no plan yet.`; a lapsed plan is no plan. */
export function joinedPlanLine(guildName: string, covered: GuildCoveredView | null | undefined): string {
  if (covered === null || covered === undefined || covered.state === 'lapsed') {
    return guildJoinCopy.noPlanYet(guildName);
  }
  return guildJoinCopy.onPlan(guildName, PLANS[covered.tier].name);
}
