// web/src/lib/billing/subscribe-step.ts
// The pure rules of the checkout guild step: card order, the lead sentence, the plan-switch
// note, the verify sub-line, the hand-off link and the copy field's label. No DOM, no fetch.
import { checkoutGuildCopy as copy } from './checkout-guild-copy';
import type { ClaimState, SubscribeStep } from './contract';
import { guildFitsRaidTier } from './checkout-params';
import { parseInterval, parsePlan } from './plan-params';
import type { PlanKey } from './plans';

/** The fields of a `/v1/me` guild entry these rules read; `MeGuild` satisfies it. */
export interface StepGuild {
  id: number;
  region: string;
  ruleset: string;
  name: string;
  /** The guild page's slug, as the API sends it; never derived from `name`. */
  slug: string;
  subscribe?: { step: SubscribeStep | null } | null;
  claim?: {
    state: ClaimState;
    yours: boolean;
    holder_name: string | null;
  } | null;
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

/** The steps whose card holds a copy link (spec §5.3). */
export type HandoffStep = 'ask_leader' | 'claimed_ask' | 'claim_pending_yours' | 'claim_pending_ask';

const SITE_ORIGIN = 'https://foreversixty.gg';
/** Claiming verifies the claimant, and a claim_blocked viewer is the guild master: no verify line. */
const NO_VERIFY_LINE_STEPS: readonly SubscribeStep[] = ['claim', 'start_claim', 'confirm_claim', 'claim_blocked'];
/** Steps where the viewer can act on the claim themselves: the lead never says the next step is someone else's. */
const OWN_CLAIM_STEPS: readonly SubscribeStep[] = [
  'claim',
  'start_claim',
  'confirm_claim',
  'claim_link_bnet',
  'claim_blocked',
];
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

function actsOnClaim(guild: StepGuild): boolean {
  const step = stepOf(guild);
  return step !== null && OWN_CLAIM_STEPS.includes(step);
}

function isUnclaimed(guild: StepGuild): boolean {
  if (stepOf(guild) === 'on_plan') return false;
  const state = guild.claim?.state;
  return state === undefined || state === 'unclaimed' || state === 'pending';
}

/**
 * Spec §9.1's lead rows ('' for no guilds: the empty states draw instead): any `ready` card; else,
 * when a card is a step the viewer takes on the claim themselves, the some-claimed form (gate fix
 * G6); else by how many guilds still have no claim holder.
 */
export function leadSentence(guilds: readonly StepGuild[]): string {
  if (guilds.length === 0) return '';
  if (guilds.some((guild) => stepOf(guild) === 'ready')) return copy.lead.ready;
  if (guilds.some(actsOnClaim)) return copy.lead.someClaimed;
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
  if (guild.claim?.state === 'claimed') return copy.verify.claimedNoApprover;
  return copy.verify.noApprover(guild.name);
}

/** Spec §6.1: the link a viewer sends to whoever can act. Carries no account id, token or name. */
export function handoffUrl(
  guild: Pick<StepGuild, 'region' | 'ruleset' | 'slug'>,
  step: HandoffStep,
  plan: string,
  interval: string,
  origin: string = SITE_ORIGIN,
): string {
  const isClaimLink = step !== 'claimed_ask';
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
export function copyFieldLabel(guild: StepGuild, step: HandoffStep, viewerIsOfficer: boolean): string {
  if (step === 'claim_pending_yours' || step === 'claim_pending_ask')
    return copy.copyField.labelSecondOfficer;
  const holder = guild.claim?.holder_name;
  if (step === 'claimed_ask' && viewerIsOfficer && holder != null) return copy.copyField.labelPerson(holder);
  return copy.copyField.labelGuildMaster(guild.name);
}

const HANDOFF_STEPS: readonly SubscribeStep[] = [
  'ask_leader',
  'claimed_ask',
  'claim_pending_yours',
  'claim_pending_ask',
];

/** The claim holder's own `claimed_ask` card has nobody to send a link to (controller ruling 2026-10-27). */
function holdsClaimedAsk(guild: StepGuild): boolean {
  return stepOf(guild) === 'claimed_ask' && guild.claim?.yours === true;
}

/** The card's hand-off step, or null when its card shows no copy field. Callers build the link from this. */
export function handoffStepOf(guild: StepGuild): HandoffStep | null {
  const step = stepOf(guild);
  if (step === null || !HANDOFF_STEPS.includes(step) || holdsClaimedAsk(guild)) return null;
  return step as HandoffStep;
}

/** Spec §5.3: the card holds a copy field, except for the claim holder's own `claimed_ask`. */
export function showsCopyField(guild: StepGuild): boolean {
  return handoffStepOf(guild) !== null;
}

/** Spec §5.4 / Decision 8: the moderator fallback shows on `ask_leader` and `claimed_ask`, not the holder's own. */
export function showsModeratorFallback(guild: StepGuild): boolean {
  const step = stepOf(guild);
  return step === 'ask_leader' || (step === 'claimed_ask' && !holdsClaimedAsk(guild));
}

/** Spec §5.4: the other-character line shows on `ask_leader` only. */
export function showsOtherCharacter(step: SubscribeStep): boolean {
  return step === 'ask_leader';
}

/** The `claimed_ask` sentence: the claim holder's own card first, then officer, then member. */
export function claimedAskLine(guild: StepGuild, viewerIsOfficer: boolean): string {
  if (guild.claim?.yours === true) return copy.claimedAsk.holdsClaim(guild.name);
  const holder = guild.claim?.holder_name;
  if (viewerIsOfficer && holder != null) return copy.claimedAsk.officer(holder, guild.name);
  return copy.claimedAsk.member(guild.name);
}
