// web/src/lib/billing/checkout-guild-copy.ts
// Every string of the checkout guild step, verbatim from design/specs/2026-10-27-subscribe-and-join.md
// §9.1. A `{…}` slot in the spec is a function argument here. The figure 25 is
// GUILD_TIER_MEMBER_LIMIT and every 7 is GUILD_TRIAL_DAYS; dates arrive as stored values.
import { utcDayMonthYear } from '../dates';
import {
  CLAIM_COOLDOWN_DAYS,
  checkoutGuildBusy,
  checkoutGuildCheckoutErrors,
  checkoutGuildClaimErrors,
  checkoutGuildPortalError,
} from './checkout-guild-errors';
import { GUILD_TIER_MEMBER_LIMIT, GUILD_TRIAL_DAYS, PLANS, type PlanKey } from './plans';

export type GuildTier = Extract<PlanKey, 'guild' | 'raid_guild'>;

const tierWord = (tier: GuildTier): string => PLANS[tier].name;
const date = utcDayMonthYear;
const CLAIM_RULE = "Only a guild's guild master can claim it here and subscribe it.";
const LAST_CLAIM_WINDOW = `You may attempt one claim every ${CLAIM_COOLDOWN_DAYS} days`;

export const checkoutGuildCopy = {
  band: { label: 'CHECKOUT', title: 'Subscribe a guild', invalidTitle: 'Checkout' },
  planSummary: {
    founder: 'Founder',
    trial: `First ${GUILD_TRIAL_DAYS} days free`,
    changePlan: 'Change plan',
  },
  sizeLine: {
    guild: `Up to ${GUILD_TIER_MEMBER_LIMIT} verified members. A 10-person raid team fits, and a 20-person one with a small bench.`,
    raidGuild: `Any size. For guilds over ${GUILD_TIER_MEMBER_LIMIT} verified members.`,
  },
  lead: {
    ready: 'Choose the guild this plan is for.',
    allUnclaimed: (count: number): string => {
      const status =
        count === 1
          ? 'Your guild has not been claimed yet'
          : count === 2
            ? 'Neither of your guilds has been claimed yet'
            : 'None of your guilds has been claimed yet';
      return `${CLAIM_RULE} ${status}, so the next step is theirs.`;
    },
    someClaimed: `${CLAIM_RULE} Here is the next step for each of your guilds.`,
    allClaimed:
      "The guild master who holds a guild's claim subscribes it. Here is where each of your guilds stands.",
  },
  rank: { leader: 'Guild master', officer: 'Officer', member: 'Member' },
  verifyMark: { verified: 'Verified', notVerified: 'Not verified' },
  more: (count: number): string => `+${count} more`,
  facts: {
    verified: (count: number): string => `${count} verified`,
    noneVerified: 'No one verified yet',
    onRoster: (count: number): string => `${count} on the roster`,
    noPlan: 'No plan',
    plan: (tier: GuildTier): string => `${tierWord(tier)} plan`,
    trial: (tier: GuildTier): string => `${tierWord(tier)} plan · trial`,
    ends: (tier: GuildTier, endsAt: string): string => `${tierWord(tier)} plan · ends ${date(endsAt)}`,
    ended: (endedAt: string): string => `Plan ended ${date(endedAt)}`,
  },
  ready: {
    holdsClaim: (guild: string): string => `You hold ${guild}'s claim.`,
    lapsed: (guild: string, endedAt: string): string =>
      `${guild}'s plan ended ${date(endedAt)}. You hold its claim.`,
    startTrial: `Start the ${GUILD_TRIAL_DAYS}-day trial`,
    subscribe: (price: string): string => `Subscribe · ${price}`,
    subscribeAgain: (price: string): string => `Subscribe again · ${price}`,
  },
  switchPlan: {
    overLimit: (verified: number): string =>
      `${verified} verified members, more than Guild's ${GUILD_TIER_MEMBER_LIMIT}.`,
    withinLimit: (verified: number): string =>
      `${verified} verified members. Guild covers up to ${GUILD_TIER_MEMBER_LIMIT} for less.`,
    toRaidGuild: 'Switch to Raid guild',
    toGuild: 'Switch to Guild',
  },
  claim: {
    sentence: (guild: string): string => `You lead ${guild}. Claim it here, then subscribe it.`,
    button: (guild: string): string => `Claim ${guild}`,
  },
  claimLinkBattlenet: {
    sentence: (guild: string): string =>
      `You lead ${guild}. Claiming as guild master needs your Battle.net account linked.`,
    button: 'Link Battle.net',
  },
  startClaim: {
    sentence: (guild: string): string =>
      `Nobody holds ${guild}'s claim yet. As an officer you can start one; a second officer or the guild master confirms it.`,
    button: 'Start a claim',
  },
  confirmClaim: {
    sentence: (name: string, guild: string): string =>
      `${name} has started a claim on ${guild}. As an officer you can confirm it.`,
    button: 'Confirm the claim',
  },
  confirm: {
    claim: (guild: string): string =>
      `Claim ${guild}? ${LAST_CLAIM_WINDOW}, hold one claimed guild at a time, and a claim can be contested.`,
    yesClaim: 'Yes, claim it',
    startClaim: (guild: string): string =>
      `Start a claim on ${guild}? A second officer or the guild master confirms it. ${LAST_CLAIM_WINDOW}.`,
    yesStartClaim: 'Yes, start the claim',
    confirmClaim: (name: string, guild: string): string =>
      `Confirm ${name}'s claim on ${guild}? They will run the guild here and can subscribe it.`,
    yesConfirmClaim: 'Yes, confirm it',
    cancel: 'Cancel',
  },
  done: {
    claim: (guild: string): string => `Claimed. You hold ${guild}.`,
    startClaim: 'Claim started. It waits for a second officer or the guild master.',
    confirmClaim: (name: string, guild: string): string => `Confirmed. ${name} holds ${guild}.`,
  },
  claimPendingYours: (guild: string): string =>
    `Your claim on ${guild} waits for a second officer or the guild master. Send them this link.`,
  claimPendingAsk: (guild: string): string =>
    `An officer has started a claim on ${guild}. Once a second officer or the guild master confirms it, they can subscribe it. Send them this link.`,
  claimedAsk: {
    holdsClaim: (guild: string): string =>
      `You hold ${guild}'s claim. Subscribing it needs a verified officer or guild-master character.`,
    officer: (name: string, guild: string): string =>
      `${name} holds ${guild}'s claim and subscribes it. Send them this link.`,
    member: (guild: string): string =>
      `The guild master who holds ${guild}'s claim subscribes it. Send them this link.`,
  },
  askLeader: (guild: string): string =>
    `Only ${guild}'s guild master can claim it and subscribe it. Send them this link.`,
  otherCharacter: (guild: string): string =>
    `Lead ${guild} on another character? Save that character with the addon, then come back.`,
  getSetUp: 'Get set up',
  moderatorFallback: 'Guild master gone? Ask us on Discord and we will sort the claim by hand.',
  moderatorLinkText: 'Ask us on Discord',
  claimBlocked: {
    holdsOther: (other: string, guild: string): string =>
      `You already hold ${other}'s claim. One account holds one guild's claim, so an officer of ${guild} claims it instead.`,
    rateLimited: (guild: string, availableAt: string): string =>
      `You made a claim in the last ${CLAIM_COOLDOWN_DAYS} days. You can claim ${guild} from ${date(availableAt)}.`,
  },
  frozen: (guild: string): string =>
    `${guild}'s claim is contested. Nobody can subscribe it until a moderator resolves the contest.`,
  checkoutOpen: (guild: string): string =>
    `Someone already has a checkout open for ${guild}. It closes on its own if they do not finish.`,
  checkAgain: 'Check again',
  openGuild: 'Open the guild',
  onPlan: {
    youPay: (guild: string, tier: GuildTier): string =>
      `${guild} is on the ${tierWord(tier)} plan. You pay for it.`,
    billedBy: (guild: string, tier: GuildTier, name: string): string =>
      `${guild} is on the ${tierWord(tier)} plan, billed by ${name}.`,
    unnamed: (guild: string, tier: GuildTier): string => `${guild} is on the ${tierWord(tier)} plan.`,
    manageBilling: 'Manage billing',
    moveToRaidGuild: 'Move to Raid guild',
    seatHeld: 'You hold one of its Premium seats.',
    seatsGiven: (given: number, total: number, payerName: string | null): string =>
      `${given} of ${total} Premium seats are given. Ask ${payerName ?? 'the guild master'} for one.`,
    seatsAllGiven: (total: number): string => `All ${total} of its Premium seats are given.`,
  },
  verify: {
    approver:
      "Not verified yet. An officer approves you from the roster, the guild's invite link verifies you at once, or you are verified when you turn up in the guild's logs.",
    noApprover: (guild: string): string =>
      `Not verified yet. You are verified when you turn up in the guild's logs, or once the guild master claims ${guild} and approves you.`,
    short: "Not verified yet. Turning up in the guild's logs also verifies you.",
  },
  copyField: {
    labelGuildMaster: (guild: string): string => `Link for ${guild}'s guild master`,
    labelPerson: (name: string): string => `Link for ${name}`,
    labelSecondOfficer: 'Link for a second officer',
    copy: 'Copy link',
    copied: 'Copied',
    share: 'Share',
    failed: 'Could not copy. Select the link and copy it.',
  },
  share: {
    title: (guild: string): string => `Subscribe ${guild} on Forever Sixty`,
    claimText: (guild: string, tier: GuildTier): string =>
      `${guild} could use the ${tierWord(tier)} plan on Forever Sixty. Claim the guild here, then subscribe it.`,
    settingsText: (guild: string, tier: GuildTier): string =>
      `${guild} could use the ${tierWord(tier)} plan on Forever Sixty. Subscribe it from the guild's billing.`,
  },
  busy: checkoutGuildBusy,
  claimErrors: checkoutGuildClaimErrors,
  checkoutErrors: checkoutGuildCheckoutErrors,
  portalError: checkoutGuildPortalError,
  loading: 'Loading your guilds',
  loadError: { title: 'Your guilds did not load.', retry: 'Try again' },
  emptyA: {
    title: 'We do not know your guild yet.',
    body: 'Your guilds come from your characters. Save a character with the addon on Get set up, and its guild shows here.',
    button: 'Get set up',
  },
  emptyB: {
    title: 'None of your characters is in a guild we know.',
    body: "If you joined a guild since, save your character with the addon again. An officer's invite link also adds you at once.",
    button: 'Get set up',
  },
  foot: {
    missingGuild: 'A guild missing? Save the character that is in it with the addon.',
    getSetUp: 'Get set up',
    premium: (price: string): string => `Or Premium for yourself, ${price}.`,
    getPremium: 'Get Premium',
  },
  invalidLink: {
    text: 'Something is wrong with that link. Start again from the premium page.',
    button: 'Premium page',
  },
} as const;

const CLAIM_ERROR_LINES: Readonly<Record<string, (guild: string) => string>> = {
  already_claimed: checkoutGuildCopy.claimErrors.alreadyClaimed,
  claim_pending: checkoutGuildCopy.claimErrors.pending,
  not_eligible: checkoutGuildCopy.claimErrors.notEligible,
  rate_limited: () => checkoutGuildCopy.claimErrors.rateLimited,
  holds_other: () => checkoutGuildCopy.claimErrors.holdsAnother,
  same_account: () => checkoutGuildCopy.claimErrors.sameAccount,
};

/** The §9.1 line for a claim refusal key (claim-flow.ts maps the API status to one); never the API sentence. */
export function claimErrorLine(code: string, guild: string): string {
  return CLAIM_ERROR_LINES[code]?.(guild) ?? checkoutGuildCopy.claimErrors.generic;
}
