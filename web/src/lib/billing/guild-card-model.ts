// web/src/lib/billing/guild-card-model.ts
// What a checkout guild card draws, as pure rules: the viewer's characters in the guild, the
// facts line, and the step sentence. Words come from checkout-guild-copy; no DOM, no fetch.
import type { MeCharacter } from '../account/me-types';
import { mainCharacter } from '../account/main-character';
import type { GuildCoveredView, GuildSubscribeView } from './contract';
import { checkoutGuildCopy as copy, type GuildTier } from './checkout-guild-copy';
import { claimedAskLine, type StepGuild } from './subscribe-step';

/** The most characters one card names before `+{n} more`. */
export const MAX_YOU_ENTRIES = 2;

export interface YouEntry {
  key: string;
  name: string;
  characterClass: string;
  rankWord: string;
  verified: boolean;
}

export interface YouRow {
  entries: YouEntry[];
  /** `+{n} more`, or null when everyone fits. */
  more: string | null;
}

const RANK_WORDS: Readonly<Record<string, string>> = {
  leader: copy.rank.leader,
  officer: copy.rank.officer,
};

function rankWord(rank: string | undefined): string {
  return (rank !== undefined && RANK_WORDS[rank]) || copy.rank.member;
}

/** The viewer's characters in this guild, the main first, at most two shown. */
export function youRow(
  characters: readonly MeCharacter[],
  guildId: number,
  mainKey: string | undefined,
): YouRow {
  const inGuild = characters.filter((character) => character.guild?.id === guildId);
  const mainKeyInGuild = mainCharacter(inGuild, mainKey)?.key;
  const ordered = [...inGuild].sort(
    (a, b) => Number(b.key === mainKeyInGuild) - Number(a.key === mainKeyInGuild),
  );
  const entries = ordered.slice(0, MAX_YOU_ENTRIES).map((character) => ({
    key: character.key,
    name: character.name,
    characterClass: character.class ?? '',
    rankWord: rankWord(character.guild?.rank),
    verified: character.guild?.verified === true,
  }));
  const hidden = ordered.length - entries.length;
  return { entries, more: hidden > 0 ? copy.more(hidden) : null };
}

function planWord(covered: GuildCoveredView | null | undefined): string {
  if (covered == null) return copy.facts.noPlan;
  switch (covered.state) {
    case 'trial':
      return copy.facts.trial(covered.tier);
    case 'ending':
      return covered.ends_at === null
        ? copy.facts.plan(covered.tier)
        : copy.facts.ends(covered.tier, covered.ends_at);
    case 'lapsed':
      return covered.ends_at === null ? copy.facts.noPlan : copy.facts.ended(covered.ends_at);
    default:
      return copy.facts.plan(covered.tier);
  }
}

/** One fact: the figure (drawn in mono) and the words after it; a fact with no number has no figure. */
export interface Fact {
  figure: string | null;
  rest: string;
}

function countFact(count: number, text: string): Fact {
  const figure = String(count);
  return { figure, rest: text.slice(figure.length) };
}

export interface FactsGuild {
  verified_count: number;
  roster_count?: number | null;
  covered?: GuildCoveredView | null;
}

/** The facts, in order: verified count, roster count (when known), the plan word. */
export function guildFacts(guild: FactsGuild): Fact[] {
  const verified =
    guild.verified_count > 0
      ? countFact(guild.verified_count, copy.facts.verified(guild.verified_count))
      : { figure: null, rest: copy.facts.noneVerified };
  const roster =
    guild.roster_count == null
      ? []
      : [countFact(guild.roster_count, copy.facts.onRoster(guild.roster_count))];
  return [verified, ...roster, { figure: null, rest: planWord(guild.covered) }];
}

/** The id of the first guild whose step differs between two reads, or null; a new guild counts. */
export function changedGuildId(
  before: { guilds: readonly StepGuildRef[] } | null,
  after: { guilds: readonly StepGuildRef[] },
): number | null {
  if (before === null) return null;
  const changed = after.guilds.find((guild) => {
    const previous = before.guilds.find((other) => other.id === guild.id);
    return previous !== undefined && previous.subscribe?.step !== guild.subscribe?.step;
  });
  return changed?.id ?? null;
}

interface StepGuildRef {
  id: number;
  subscribe?: { step: string | null } | null;
}

export interface SentenceGuild extends StepGuild {
  subscribe?: GuildSubscribeView | null;
  covered?: GuildCoveredView | null;
}

function onPlanSentence(guild: SentenceGuild, chosen: GuildTier): string {
  const covered = guild.covered;
  const tier = covered?.tier ?? chosen;
  if (covered?.you_pay === true) return copy.onPlan.youPay(guild.name, tier);
  if (covered != null && covered.billed_by !== '')
    return copy.onPlan.billedBy(guild.name, tier, covered.billed_by);
  return copy.onPlan.unnamed(guild.name, tier);
}

/** The read-only seat line under an `on_plan` card the viewer does not pay for, or null. */
export function seatLine(covered: GuildCoveredView | null | undefined): string | null {
  if (covered == null || covered.you_pay || covered.seats_total === 0) return null;
  if (covered.you_hold_seat) return copy.onPlan.seatHeld;
  if (covered.seats_given >= covered.seats_total) return copy.onPlan.seatsAllGiven(covered.seats_total);
  return copy.onPlan.seatsGiven(covered.seats_given, covered.seats_total, covered.billed_by || null);
}

function blockedSentence(guild: SentenceGuild): string | null {
  const subscribe = guild.subscribe;
  if (subscribe?.reason === 'holds_other' && subscribe.other_guild_name !== null)
    return copy.claimBlocked.holdsOther(subscribe.other_guild_name, guild.name);
  if (subscribe?.reason === 'rate_limited' && subscribe.available_at !== null)
    return copy.claimBlocked.rateLimited(guild.name, subscribe.available_at);
  return null;
}

function readySentence(guild: SentenceGuild): string {
  const covered = guild.covered;
  if (covered?.state === 'lapsed' && covered.ends_at !== null)
    return copy.ready.lapsed(guild.name, covered.ends_at);
  return copy.ready.holdsClaim(guild.name);
}

/** Spec §9.1's one sentence for the card's step; null when the step has none (or lacks its data). */
export function stepSentence(
  guild: SentenceGuild,
  chosen: GuildTier,
  viewerIsOfficer: boolean,
): string | null {
  const name = guild.name;
  switch (guild.subscribe?.step) {
    case 'ready':
      return readySentence(guild);
    case 'claim':
      return copy.claim.sentence(name);
    case 'claim_link_bnet':
      return copy.claimLinkBattlenet.sentence(name);
    case 'start_claim':
      return copy.startClaim.sentence(name);
    case 'confirm_claim':
      return copy.confirmClaim.sentence(guild.claim?.holder_name ?? copy.rank.officer, name);
    case 'claimed_ask':
      return claimedAskLine(guild, viewerIsOfficer);
    case 'ask_leader':
      return copy.askLeader(name);
    case 'claim_pending_yours':
      return copy.claimPendingYours(name);
    case 'claim_pending_ask':
      return copy.claimPendingAsk(name);
    case 'checkout_open':
      return copy.checkoutOpen(name);
    case 'frozen':
      return copy.frozen(name);
    case 'claim_blocked':
      return blockedSentence(guild);
    case 'on_plan':
      return onPlanSentence(guild, chosen);
    default:
      return null;
  }
}
