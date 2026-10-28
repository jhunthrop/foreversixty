// web/src/lib/billing/checkout-params.ts
// The one place /premium/checkout reads its query string. The page is static, so only the
// hydrated launcher ever sees the real string; every value is checked against a closed set
// or shape, and anything else makes the whole request invalid (never echoed, never guessed).
import type { MeGuild } from '../account/api';
import { parseInterval, parsePlan } from './plan-params';
import { GUILD_TIER_MEMBER_LIMIT, type Interval, type PlanKey } from './plans';

export interface CheckoutParams {
  plan: PlanKey;
  interval: Interval;
  guildId?: number;
  /** Premium only: the ended-seat resubscribe link's id. */
  comebackId?: number;
  status: 'success' | null;
}

const POSITIVE_INTEGER = /^[1-9][0-9]*$/;

export function isGuildPlan(plan: PlanKey): boolean {
  return plan !== 'premium';
}

/** Spec §5 / §5.7: the guild step applies, known from the URL alone, so its loading state
 *  draws before /v1/me answers. A `guild_id` request (spec §5.6) shows it too: the step starts a
 *  `ready` guild's checkout itself and floats any other guild's card first. */
export function showsGuildStep(params: CheckoutParams): boolean {
  return isGuildPlan(params.plan) && params.status !== 'success';
}

/** undefined when absent, null when present but not a positive integer. */
function positiveInteger(raw: string | null): number | undefined | null {
  if (raw === null) return undefined;
  return POSITIVE_INTEGER.test(raw) ? Number.parseInt(raw, 10) : null;
}

export function parseCheckoutParams(search: string): CheckoutParams | null {
  const params = new URLSearchParams(search);
  const plan = parsePlan(params.get('plan'));
  const interval = parseInterval(params.get('interval'));
  const guildId = positiveInteger(params.get('guild_id'));
  const comebackId = positiveInteger(params.get('comeback'));
  if (plan === undefined || interval === undefined) return null;
  if (guildId === null || comebackId === null) return null;
  if (!isGuildPlan(plan) && guildId !== undefined) return null;
  if (isGuildPlan(plan) && comebackId !== undefined) return null;
  const status = params.get('status') === 'success' ? 'success' : null;
  if (status === 'success' && isGuildPlan(plan) && guildId === undefined) return null;
  return { plan, interval, guildId, comebackId, status };
}

/** Guilds this account may subscribe: it can bill, and the guild has no plan yet. */
export function billableGuilds(guilds: readonly MeGuild[]): MeGuild[] {
  return guilds.filter((guild) => (guild.can ?? []).includes('billing') && guild.plan == null);
}

export function guildFitsRaidTier(verifiedCount: number): boolean {
  return verifiedCount > GUILD_TIER_MEMBER_LIMIT;
}

/** `raid_guild` above the Guild plan's member limit, else `guild` (spec 6.5). */
export function fittingGuildPlan(verifiedCount: number): Extract<PlanKey, 'guild' | 'raid_guild'> {
  return guildFitsRaidTier(verifiedCount) ? 'raid_guild' : 'guild';
}
