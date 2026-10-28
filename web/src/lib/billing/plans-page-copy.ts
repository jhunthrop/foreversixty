// web/src/lib/billing/plans-page-copy.ts
// Every visible string on /premium except the checkout chooser (billing/copy.ts), verbatim from
// design/specs/2026-10-27-launch-pricing.md section 6.1. No exclamation marks, no urgency. Price
// text is built from ./plans, never written here.
import { perMemberAmount, perRaiderAmounts } from './plan-costs';
import {
  exampleCaption,
  serverRunExample,
  type CaptionTemplates,
  type ServerRunExample,
} from './premium-examples';
import {
  GUILD_TIER_MEMBER_LIMIT,
  GUILD_TRIAL_DAYS,
  RAID_SIZES,
  TRIAL_REMINDER_DAYS,
  SEAT_MOVE_LOCK_DAYS,
  perMonthLine,
  type Interval,
} from './plans';
import { CAPS, LANE_ITERATION_CEILING } from '../sim/precision';

export interface PlanRule {
  term: string;
  sentence: string;
}

export interface FaqEntry {
  question: string;
  answer: string;
}

const count = (n: number): string => n.toLocaleString('en-US');

const INTERVAL_WORD: Record<Interval, string> = { yearly: 'a year', monthly: 'a month' };

function perMember(interval: Interval, founder: boolean): string {
  return `${perMemberAmount(interval, founder)} per member ${INTERVAL_WORD[interval]} at ${GUILD_TIER_MEMBER_LIMIT} members.`;
}

function perRaider(interval: Interval, founder: boolean): string {
  const [a40, a20, a10] = perRaiderAmounts(interval, founder);
  const [s40, s20, s10] = RAID_SIZES;
  return `${a40} per raider ${INTERVAL_WORD[interval]} in a ${s40}-person raid, ${a20} in a ${s20}, ${a10} in a ${s10}.`;
}

const TRIAL_CANCEL = `Cancel before day ${GUILD_TRIAL_DAYS + 1} and nothing is charged`;

const serverRunCaptionTemplates: CaptionTemplates = {
  preset: (f) =>
    `A level-${f.level} ${f.specLabel} in the site's preset gear: ${f.combinations} combinations on our servers, ${f.ranAt}.`,
  ownGear: (f) =>
    `${f.characterName}, ${f.specLabel}: ${f.combinations} combinations on our servers, ${f.ranAt}.`,
};

export const plansPageCopy = {
  title: 'Plans',
  description:
    'What is free on Forever Sixty for good, what Premium, Guild and Raid guild add, and what each costs.',
  band: {
    eyebrow: 'PLANS',
    heading: 'What is free, and what the plans add',
    lede: 'Everything a player needs is free, and the free list never shrinks. Premium and the guild plans pay for the site. There are no ads for anyone.',
    guildLead: 'Running a guild?',
    guildLink: 'Guild and Raid guild are further down',
  },
  guildRegion: {
    heading: 'For guild leaders',
    lede: 'One officer pays, and the whole guild gets raids, sign-ups and attendance. Members never pay to sign up.',
    members:
      'Plans count verified members: players with a Forever Sixty account whose character the guild has confirmed. People in your Discord without an account do not count.',
  },
  founder: {
    label: 'Founder',
    open: 'Founder prices until 6 January 2027. Buy by then and the price stays for as long as you stay subscribed.',
  },
  interval: {
    groupLabel: 'Billing interval',
    labels: { yearly: 'Yearly', monthly: 'Monthly' } as Record<Interval, string>,
  },
  free: {
    name: 'Free',
    audience: 'For every player',
    price: '$0',
    note: 'No card, no trial, no end date.',
    signIn: 'Sign in with Battle.net from the top of any page.',
    features: [
      'Addon export and import',
      'BiS lists for every spec and level band',
      'Talent planner',
      'Browser sims',
      'Logs, live logging and parses',
      'Rankings and the tier list',
      'Class guides',
    ],
    promise: 'This list never shrinks.',
  },
  premium: {
    name: 'Premium',
    audience: 'For one player',
    support: 'Free stays complete. Premium pays for the servers and keeps the site going.',
    features: [
      `Try up to ${count(CAPS.server)} gear, talent or drop combinations in one sim and find your best DPS. Your browser stops at ${count(CAPS.browser)}.`,
      'Close the tab. Server sims keep running, and the site tells you when they finish.',
      `Runs go up to ${count(LANE_ITERATION_CEILING.server)} iterations instead of ${count(LANE_ITERATION_CEILING.browser)}, so close upgrades come out clear.`,
      'See if you are improving: your Forever Score and every boss parse on a chart over time, with the fights behind each point.',
      'A message in your own Discord channel when a server sim finishes or a log is parsed.',
      "A Supporter mark beside your name on your characters and in your guild's roster.",
    ] as readonly string[],
    scope:
      'Sims cover damage specs; healing and tanking specs are not simulated yet. History, Discord messages and the Supporter mark work for every spec.',
    // Held until the log coach ships (B5): never rendered, never in `features`.
    featuresHeld: {
      coach:
        'Coach: after every logged raid, what to fix first — the casts you missed, the buffs you let drop, the consumables you skipped — learned from the engine and from the top logs on the site.',
    },
  },
  guild: {
    name: 'Guild',
    audience: `For a guild of up to ${GUILD_TIER_MEMBER_LIMIT} verified members`,
    features: [
      'Post raids on the site and in Discord; members sign up with one tap.',
      'See who came: attendance and a report for every raid night from its log.',
      'The Forever Sixty bot in your Discord: raid posts with buttons, plus /bis and /char lookups.',
      '1 Premium seat to give a member, such as your raid leader.',
      `Up to ${GUILD_TIER_MEMBER_LIMIT} verified members, so a guild that runs 10- or 20-person raids fits under Guild. Nothing switches off if you grow past it.`,
    ],
  },
  raidGuild: {
    name: 'Raid guild',
    audience: 'For a raiding guild of any size',
    features: [
      'Loot council: each drop ranks the raiders it helps most, with what they wear now, their attendance and their loot this tier. Officers award it from the same screen.',
      'Readiness board: before the pull, one list of who is missing enchants, consumables, talent points or upgrades, worst first.',
      '3 Premium seats to give members.',
      'Any number of members.',
      'Everything in Guild: raids, sign-ups, attendance and the Discord bot.',
    ],
  },
  price: {
    suffix: { yearly: ' /year', monthly: ' /month' } as Record<Interval, string>,
    perMonth: (yearlyUsd: number): string => perMonthLine(yearlyUsd),
    orYearly: (yearlyPrice: string): string => `Or ${yearlyPrice}.`,
    founder: 'Founder price, kept while you stay subscribed.',
    perMember,
    perRaider,
  },
  buttons: {
    subscribe: (price: string): string => `Subscribe · ${price}`,
    subscribeGuild: (price: string): string => `Subscribe a guild · ${price}`,
  },
  trial: `First ${GUILD_TRIAL_DAYS} days free, once per guild. ${TRIAL_CANCEL}; guild settings shows the date of the first charge.`,
  // Defined but appended to nothing until the trial email ships (B10 ruling 1); that lane appends it
  // to `trial` and the Trial rule and flips the guard test.
  trialEmail: `We email you ${TRIAL_REMINDER_DAYS} days before the trial ends.`,
  examples: {
    serverRun: {
      label: 'EXAMPLE · A REAL SERVER RUN',
      caption: serverRunCaptionTemplates,
      link: 'Open this result',
    },
  },
  historyPromise:
    'If a guild plan ends, its raids, attendance, loot and reports stay readable for good, and officers can export attendance and loot as CSV at any time.',
  howItWorks: 'How the guild plans work',
  rules: [
    {
      term: 'Verified members',
      sentence:
        "A verified member is a player with a Forever Sixty account whose character the guild has confirmed: an officer approved it, they joined by the guild's invite link, or the character is in the guild's logs. Each player counts once, however many characters they have in the guild. Discord members without an account do not count.",
    },
    {
      term: 'Size',
      sentence: `Guild counts verified members each day, the leader included. Above ${GUILD_TIER_MEMBER_LIMIT} the officer is told before renewal and offered Raid guild. Nothing switches off mid-plan.`,
    },
    {
      term: 'Paying together',
      sentence: "One person's card pays for the plan. Members cannot split the bill on the site.",
    },
    {
      term: 'Seats',
      sentence:
        'The subscriber or the guild leader gives each seat to any verified member. A seat holder has Premium while the plan lasts.',
    },
    {
      term: 'Moving a seat',
      sentence: `A seat can move once every ${SEAT_MOVE_LOCK_DAYS} days. The old holder keeps Premium until it moves. If a holder leaves the guild, the seat comes back at once.`,
    },
    {
      term: 'Already on Premium',
      sentence:
        'A member who pays for Premium accepts a seat first. Their own Premium stops renewing and the seat starts when it ends, so nobody pays twice. If the seat is taken away later, they get a link to come back at the price they had.',
    },
    {
      term: 'Trial',
      sentence: `A guild's first ${GUILD_TRIAL_DAYS} days are free, once per guild. The card is added at the start, and guild settings shows the date of the first charge from day one. ${TRIAL_CANCEL}.`,
    },
    {
      term: 'Renewals',
      sentence:
        'We email 7 days before every yearly renewal. Cancel in one click; the plan runs to the end of the period you paid for.',
    },
    {
      term: 'A new guild leader',
      sentence:
        'The subscriber hands the plan to another officer from guild settings. The founder price goes with it.',
    },
    {
      term: 'Refunds',
      sentence:
        'Ask within 14 days of a first charge for a refund. After that there are no part-year refunds.',
    },
  ] as PlanRule[],
  foot: 'Prices in US dollars. Cancel any time; you keep the plan to the end of the period you paid for.',
  faq: [
    {
      question: 'What happens if I cancel?',
      answer:
        "You keep everything through the end of the period you already paid for. After that, the account works like a free one, and nothing is deleted. A guild's records stay readable for good.",
    },
    {
      question: 'Can I get a refund?',
      answer: "Email us and we'll sort it out — see the refund policy for the details.",
    },
    {
      question: 'Is anything from Battle.net ever paid?',
      answer:
        "No. Your characters, your sign-in, and everything the game itself tells us about you stays free, always — Blizzard's own rules for using their data require this, and we would want it that way regardless.",
    },
    {
      question: 'Are there ads?',
      answer:
        'No. Not for anyone, paying or not. The site is paid for by Premium and the guild plans, not by ads, and that is not a perk you are buying — it is true for everyone.',
    },
    {
      question: 'Who am I paying?',
      answer:
        'COMMISH LLC, the company behind Forever Sixty. Stripe handles the payment; we never see or store your card.',
    },
  ] as FaqEntry[],
};

export const perMemberLine = plansPageCopy.price.perMember;
export const perRaiderLine = plansPageCopy.price.perRaider;

/** The server-run example's caption, also the image alt. */
export function serverRunCaption(ex: ServerRunExample = serverRunExample): string {
  return exampleCaption(ex, plansPageCopy.examples.serverRun.caption);
}
