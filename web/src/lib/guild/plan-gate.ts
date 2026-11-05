// web/src/lib/guild/plan-gate.ts
// The gate sentence's two parts (spec 2026-11-05 §3.2): what the feature needs, then what this
// viewer can do about it. Part 1 depends on the feature, part 2 on the viewer and the guild's
// billing, never on the tab. Pure, so every row of the spec's table is pinned without a page.
import type { GuildCoveredView, GuildSubscribeView } from '../billing/contract';
import type { ClaimHolder, ClaimStateView } from './api';
import { guildGateCopy } from './copy';

export type GateFeature = 'raids' | 'discord' | 'loot' | 'readiness';

/** The Raid guild features; Raids and Discord come with both plans. */
const RAID_GUILD_FEATURES: ReadonlySet<GateFeature> = new Set(['loot', 'readiness']);

const BILLING_HASH = '#billing';

export interface GateFacts {
  /** The claim as home or settings read it; null when neither answered. */
  claim: Pick<ClaimStateView, 'state' | 'yours' | 'claimed_by'> | null;
  /** The viewer's rank here is guild master. */
  guildMaster: boolean;
  /** `/v1/me` `subscribe.trial_used`: known only once the viewer can subscribe. */
  trialUsed: boolean | null;
  /** The guild had a plan and it ended. */
  lapsed: boolean;
  /** The guild is on the Guild plan, so only the Raid guild features are gated. */
  guildTier: boolean;
  settingsHref: string;
  claimHref: string;
}

export type GateAction =
  | { kind: 'link'; label: string; href: string }
  /** `Ask {crest}{name} to turn it on.`; a null holder has no character on the roster. */
  | { kind: 'ask'; holder: ClaimHolder | null }
  | { kind: 'note'; text: string };

export interface GateFactsInput {
  claim: GateFacts['claim'];
  membership: {
    rank?: string;
    covered?: Pick<GuildCoveredView, 'state'> | null;
    subscribe?: Pick<GuildSubscribeView, 'trial_used'> | null;
  } | null;
  /** Both tools flags as the page holds them: on with Raid guild off is the Guild plan. */
  officerTools: boolean;
  raidGuildTools: boolean;
  settingsHref: string;
  claimHref: string;
}

/** The gate's facts from what a guild page already reads: no request of its own. */
export function gateFactsFrom(input: GateFactsInput): GateFacts {
  const { claim, membership, officerTools, raidGuildTools, settingsHref, claimHref } = input;
  return {
    claim,
    guildMaster: membership?.rank === 'leader',
    trialUsed: membership?.subscribe?.trial_used ?? null,
    lapsed: membership?.covered?.state === 'lapsed',
    guildTier: officerTools && !raidGuildTools,
    settingsHref,
    claimHref,
  };
}

/** The holder's billing link: the plan preselected from a Raid guild feature. */
function holderAction(feature: GateFeature, facts: GateFacts): GateAction {
  const raidGuildFeature = RAID_GUILD_FEATURES.has(feature);
  if (raidGuildFeature && facts.guildTier) {
    return {
      kind: 'link',
      label: guildGateCopy.moveToRaidGuild,
      href: `${facts.settingsHref}?upgrade=1${BILLING_HASH}`,
    };
  }
  const query = raidGuildFeature ? '?plan=raid_guild' : '';
  const trialOpen = facts.trialUsed === false && !facts.lapsed;
  return {
    kind: 'link',
    label: trialOpen ? guildGateCopy.startTrial : guildGateCopy.subscribe,
    href: `${facts.settingsHref}${query}${BILLING_HASH}`,
  };
}

/** Part 2 of the sentence, for this viewer. */
export function gateAction(feature: GateFeature, facts: GateFacts): GateAction {
  const { claim } = facts;
  if (claim === null) return { kind: 'ask', holder: null };
  if (claim.state === 'unclaimed' || claim.state === 'pending') {
    return facts.guildMaster
      ? { kind: 'link', label: guildGateCopy.claimFirst, href: facts.claimHref }
      : { kind: 'note', text: guildGateCopy.masterClaimsFirst };
  }
  return claim.yours === true
    ? holderAction(feature, facts)
    : { kind: 'ask', holder: claim.claimed_by ?? null };
}
