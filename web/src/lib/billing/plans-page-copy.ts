// web/src/lib/billing/plans-page-copy.ts
// Every visible string on /premium except the checkout chooser (billing/copy.ts), verbatim from
// design/specs/2026-10-27-launch-pricing.md section 6.1 as amended by B5 (the sample coach review's
// frame is design/specs/2026-10-28-polish.md section 3.6). No exclamation marks, no urgency. Price
// text is built from ./plans, never written here.
import { perMemberAmount, perRaiderAmounts } from './plan-costs';
import { coachCopy } from '../coach/copy';
import {
  GUILD_TIER_MEMBER_LIMIT,
  FOUNDER_LAST_DAY,
  FOUNDER_PRICE_UNTIL,
  GUILD_TRIAL_DAYS,
  PLANS,
  RAID_SIZES,
  REFUND_WINDOW_DAYS,
  RENEWAL_REMINDER_DAYS,
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

/** The `/premium` anchors a contextual link lands on (spec 2026-10-29-premium-entry section 8). */
export type PremiumAnchor = 'coach' | 'server-sims' | 'history' | 'notifications';

export interface PremiumFeature {
  readonly text: string;
  /** The first line of an anchor carries its id; every line of it carries `data-anchor`. */
  readonly anchor: PremiumAnchor | null;
}

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
    open: `Founder prices until ${FOUNDER_PRICE_UNTIL}. Buy by then and the price stays for as long as you stay subscribed.`,
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
    features: [
      'Addon export and import',
      'BiS lists for every spec and level band',
      'Talent planner',
      'Browser sims',
      'Logs, live logging and parses',
      'Rankings and the tier list',
      'Class guides',
    ],
  },
  premium: {
    name: 'Premium',
    audience: 'For one player',
    support: 'Free stays complete. Premium pays for the servers and keeps the site going.',
    // Each line's `/premium` anchor (spec 2026-10-29-premium-entry section 8): a contextual link
    // lands on its line, and every line of one anchor is highlighted together.
    features: [
      {
        text: 'Coach: after every logged raid, what to fix first. It names the casts you missed, the buffs you let drop and the consumables you skipped.',
        anchor: 'coach',
      },
      {
        text: `Try up to ${count(CAPS.server)} gear, talent or drop combinations in one sim and find your best DPS. Your browser stops at ${count(CAPS.browser)}.`,
        anchor: 'server-sims',
      },
      {
        text: 'Close the tab. Server sims keep running, and the site tells you when they finish.',
        anchor: 'server-sims',
      },
      {
        text: `Runs go up to ${count(LANE_ITERATION_CEILING.server)} iterations instead of ${count(LANE_ITERATION_CEILING.browser)}, so close upgrades come out clear.`,
        anchor: 'server-sims',
      },
      {
        text: 'See if you are improving: your Forever Score and every boss parse on a chart over time, with the fights behind each point.',
        anchor: 'history',
      },
      {
        text: 'A message in your own Discord channel when a server sim finishes or a log is parsed.',
        anchor: 'notifications',
      },
      {
        text: "A Supporter mark beside your name on your characters and in your guild's roster.",
        anchor: null,
      },
    ] as readonly PremiumFeature[],
    scope:
      'Sims cover damage specs; healing and tanking specs are not simulated yet. The coach reviews every spec; damage specs get the most findings.',
  },
  guild: {
    name: 'Guild',
    audience: `For a guild of up to ${GUILD_TIER_MEMBER_LIMIT} verified members`,
    features: [
      'Post raids on the site and in Discord; members sign up with one tap.',
      'See who came: attendance and a report for every raid night from its log.',
      'The Forever Sixty bot in your Discord: raid posts with buttons, plus /bis and /char lookups.',
      `${PLANS.guild.seats} Premium seat to give a member, such as your raid leader.`,
      `Up to ${GUILD_TIER_MEMBER_LIMIT} verified members, so a guild that runs ${RAID_SIZES[2]}- or ${RAID_SIZES[1]}-person raids fits under Guild. Nothing switches off if you grow past it.`,
    ],
  },
  raidGuild: {
    name: 'Raid guild',
    audience: 'For a raiding guild of any size',
    features: [
      'Loot council: each drop ranks the raiders it helps most, with what they wear now, their attendance and their loot this tier. Officers award it from the same screen.',
      'Readiness board: before the pull, one list of who is missing enchants, consumables, talent points or upgrades, worst first.',
      `${PLANS.raid_guild.seats} Premium seats to give members.`,
      'Any number of members.',
      'Everything in Guild: raids, sign-ups, attendance and the Discord bot.',
    ],
  },
  price: {
    suffix: { yearly: ' /year', monthly: ' /month' } as Record<Interval, string>,
    perMonth: (yearlyUsd: number): string => perMonthLine(yearlyUsd),
    orYearly: (yearlyPrice: string): string => `Or ${yearlyPrice}.`,
    founderUntil: (listPrice: string): string =>
      `Founder price until ${FOUNDER_LAST_DAY}, then ${listPrice}.`,
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
  // The sample coach review in column 3 (polish spec 2026-10-28 section 3.6). The review's own
  // words come from the coach fixture; only the frame is written here.
  coachSample: {
    tag: 'Sample',
    label: 'What the coach says after a raid',
    context: (spec: string, cls: string, zone: string, size: number, kill: boolean): string =>
      `${spec} ${cls} · ${zone} · ${size}-player ${kill ? 'kill' : 'wipe'}`,
    fixFirst: coachCopy.labelFixFirst,
    then: coachCopy.labelThen,
    doingWell: coachCopy.labelWell,
    caption: 'A sample review. Yours is written from your own log after every raid you upload.',
    ariaLabel: 'A sample coach review',
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
      sentence: `We email ${RENEWAL_REMINDER_DAYS} days before every yearly renewal. Cancel in one click; the plan runs to the end of the period you paid for.`,
    },
    {
      term: 'A new guild leader',
      sentence:
        'The subscriber hands the plan to another officer from guild settings. The founder price goes with it.',
    },
    {
      term: 'Refunds',
      sentence: `Ask within ${REFUND_WINDOW_DAYS} days of a first charge for a refund. After that there are no part-year refunds.`,
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
