// web/src/lib/billing/subscribe-step.ts
// The pure rules of the checkout guild step: card order, the lead sentence, the plan-switch
// note, the verify sub-line, the hand-off link and the copy field's label. No DOM, no fetch.
import { checkoutGuildCopy as copy } from './checkout-guild-copy';
import { guildFitsRaidTier } from './checkout-params';
import { parseInterval, parsePlan } from './plan-params';
import type { PlanKey } from './plans';

/** The API's `subscribe.step` words (spec §8), in resolution order. */
export type SubscribeStep =
  | 'ready'
  | 'claim'
  | 'confirm_claim'
  | 'start_claim'
  | 'claim_link_bnet'
  | 'claimed_ask'
  | 'ask_leader'
  | 'claim_pending_yours'
  | 'claim_pending_ask'
  | 'checkout_open'
  | 'frozen'
  | 'claim_blocked'
  | 'on_plan';

/** The fields of a `/v1/me` guild entry these rules read; `MeGuild` satisfies it. */
export interface StepGuild {
  id: number;
  region: string;
  ruleset: string;
  name: string;
  subscribe?: { step: SubscribeStep | null } | null;
  claim?: { state: 'unclaimed' | 'pending' | 'claimed' | 'contested'; holder_name: string | null } | null;
  verify?: { has_approver: boolean } | null;
}

/** Spec §5.1's card order. */
export const STEP_ORDER: readonly SubscribeStep[] = [
  'ready',
  'claim',
  'confirm_claim',
  'start_claim',
  'claim_link_bnet',
  'claimed_ask',
  'ask_leader',
  'claim_pending_yours',
  'claim_pending_ask',
  'checkout_open',
  'frozen',
  'claim_blocked',
  'on_plan',
];

const SITE_ORIGIN = 'https://foreversixty.gg';
const CLAIM_LINK_STEPS: readonly SubscribeStep[] = ['ask_leader', 'claim_pending_yours', 'claim_pending_ask'];
const NO_VERIFY_LINE_STEPS: readonly SubscribeStep[] = ['claim', 'start_claim', 'confirm_claim'];
const SHORT_VERIFY_STEPS: readonly SubscribeStep[] = ['ask_leader', 'claim_pending_ask'];

function stepOf(guild: StepGuild): SubscribeStep | null {
  return guild.subscribe?.step ?? null;
}

function rankOf(guild: StepGuild): number {
  const step = stepOf(guild);
  return step === null ? STEP_ORDER.length : STEP_ORDER.indexOf(step);
}

/** Cards in step order, ties in the order given; `firstId`, when in the list, leads. */
export function orderGuildCards<T extends StepGuild>(guilds: readonly T[], firstId?: number): T[] {
  const byStep = guilds
    .map((guild, index) => ({ guild, index }))
    .sort((a, b) => rankOf(a.guild) - rankOf(b.guild) || a.index - b.index)
    .map(({ guild }) => guild);
  const first = byStep.find((guild) => guild.id === firstId);
  return first === undefined ? byStep : [first, ...byStep.filter((guild) => guild !== first)];
}

function isUnclaimed(guild: StepGuild): boolean {
  if (stepOf(guild) === 'on_plan') return false;
  const state = guild.claim?.state;
  return state === undefined || state === 'unclaimed' || state === 'pending';
}

/** Spec §9.1's lead rows: any `ready` card, else by how many guilds still have no claim holder. */
export function leadSentence(guilds: readonly StepGuild[]): string {
  if (guilds.some((guild) => stepOf(guild) === 'ready')) return copy.lead.ready;
  const unclaimed = guilds.filter(isUnclaimed).length;
  if (unclaimed === 0) return copy.lead.allClaimed;
  return unclaimed === guilds.length ? copy.lead.allUnclaimed(unclaimed) : copy.lead.someClaimed;
}

export interface SwitchNote {
  text: string;
  link: string;
  to: PlanKey;
}

/** The note above a `ready` button when the chosen guild plan does not fit the verified count. */
export function switchNote(plan: PlanKey, verifiedCount: number): SwitchNote | null {
  const needsRaidGuild = guildFitsRaidTier(verifiedCount);
  if (plan === 'guild' && needsRaidGuild) {
    return {
      text: copy.switchPlan.overLimit(verifiedCount),
      link: copy.switchPlan.toRaidGuild,
      to: 'raid_guild',
    };
  }
  if (plan === 'raid_guild' && !needsRaidGuild) {
    return { text: copy.switchPlan.withinLimit(verifiedCount), link: copy.switchPlan.toGuild, to: 'guild' };
  }
  return null;
}

/** Spec §5.4: what a viewer with no verified character reads under the step. */
export function verifySubLine(guild: StepGuild, viewerVerified: boolean): string | null {
  const step = stepOf(guild);
  if (viewerVerified) return null;
  if (step !== null && NO_VERIFY_LINE_STEPS.includes(step)) return null;
  if (guild.verify?.has_approver === true) return copy.verify.approver;
  if (step !== null && SHORT_VERIFY_STEPS.includes(step)) return copy.verify.short;
  return copy.verify.noApprover(guild.name);
}

/** Spec §6.1: the link a viewer sends to whoever can act. Carries no account id, token or name. */
export function handoffUrl(
  guild: Pick<StepGuild, 'region' | 'ruleset'> & { slug: string },
  step: SubscribeStep,
  plan: string,
  interval: string,
  origin: string = SITE_ORIGIN,
): string {
  const isClaimLink = CLAIM_LINK_STEPS.includes(step);
  if (!isClaimLink && step !== 'claimed_ask') throw new Error(`step ${step} has no hand-off link`);
  const query = new URLSearchParams();
  const knownPlan = parsePlan(plan);
  const knownInterval = parseInterval(interval);
  if (knownPlan !== undefined && knownPlan !== 'premium') query.set('plan', knownPlan);
  if (knownInterval !== undefined) query.set('interval', knownInterval);
  const queryText = query.toString();
  const search = queryText === '' ? '' : `?${queryText}`;
  const base = `${origin}/guild/${guild.region}/${guild.ruleset}/${encodeURIComponent(guild.slug)}`;
  return isClaimLink ? `${base}/claim${search}` : `${base}/settings${search}#billing`;
}

/** The screen-reader label of the hand-off link's field, by who it is for. */
export function copyFieldLabel(guild: StepGuild, step: SubscribeStep, viewerIsOfficer: boolean): string {
  if (step === 'claim_pending_yours' || step === 'claim_pending_ask')
    return copy.copyField.labelSecondOfficer;
  const holder = guild.claim?.holder_name;
  if (step === 'claimed_ask' && viewerIsOfficer && holder != null) return copy.copyField.labelPerson(holder);
  return copy.copyField.labelGuildMaster(guild.name);
}
