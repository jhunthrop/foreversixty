// web/src/lib/billing/plans-page-copy.ts
// Every visible string on /premium except the checkout chooser (billing/copy.ts), verbatim from
// design/specs/2026-10-27-launch-pricing.md section 6.1. No exclamation marks, no urgency. Price
// text is built from ./plans, never written here.
import {
  GUILD_TIER_MEMBER_LIMIT,
  GUILD_TRIAL_DAYS,
  SEAT_MOVE_LOCK_DAYS,
  perMonthLine,
  type Interval,
} from './plans';

export interface PlanRule {
  term: string;
  sentence: string;
}

export interface FaqEntry {
  question: string;
  answer: string;
}

export const plansPageCopy = {
  title: 'Plans',
  description:
    'What is free on Forever Sixty for good, what Premium, Guild and Raid guild add, and what each costs.',
  band: {
    eyebrow: 'PLANS',
    heading: 'What is free, and what the plans add',
    lede: 'Everything a player needs is free, and the free list never shrinks. Premium and the guild plans pay for the site. There are no ads for anyone.',
  },
  founder: {
    label: 'Founder',
    open: 'Founder prices until 6 January. Buy by then and the price stays for as long as you stay subscribed.',
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
    summary: {
      name: 'Free',
      price: '$0',
      line: 'Everything a player needs. This list never shrinks.',
      srOnly: 'See the free list',
    },
  },
  premium: {
    name: 'Premium',
    audience: 'For one player',
    features: [
      'Server sims: 5,000 combinations, any precision, no tab left open',
      'History charts of your score and parses',
      'Discord notifications when a server sim or a log finishes',
      'A supporter mark on your characters',
    ],
  },
  guild: {
    name: 'Guild',
    audience: `For a guild of up to ${GUILD_TIER_MEMBER_LIMIT}`,
    features: [
      `Up to ${GUILD_TIER_MEMBER_LIMIT} verified members`,
      'Raids and sign-ups, on the site and in Discord',
      'Attendance and the raid-night report',
      "The guild's Discord channel and bot",
      '1 Premium seat to give to a member',
    ],
  },
  raidGuild: {
    name: 'Raid guild',
    audience: 'For a raiding guild of any size',
    features: [
      'Everything in Guild',
      'Any number of members',
      'Loot council',
      'The readiness board, with consumables',
      '3 Premium seats to give to members',
    ],
  },
  price: {
    suffix: { yearly: ' /year', monthly: ' /month' } as Record<Interval, string>,
    perMonth: (yearlyUsd: number): string => perMonthLine(yearlyUsd),
    orYearly: (yearlyPrice: string): string => `Or ${yearlyPrice}.`,
    founder: 'Founder price, kept while you stay subscribed.',
  },
  buttons: {
    subscribe: (price: string): string => `Subscribe · ${price}`,
    subscribeGuild: (price: string): string => `Subscribe a guild · ${price}`,
  },
  trial: `First ${GUILD_TRIAL_DAYS} days free, once per guild. Cancel before day ${GUILD_TRIAL_DAYS + 1} and nothing is charged.`,
  historyPromise:
    'If a guild plan ends, its raids, attendance, loot and reports stay readable for good, and officers can export attendance and loot as CSV at any time.',
  howItWorks: 'How the guild plans work',
  rules: [
    {
      term: 'Size',
      sentence: `Guild counts verified members each day, the leader included. Above ${GUILD_TIER_MEMBER_LIMIT} the officer is told before renewal and offered Raid guild. Nothing switches off mid-plan.`,
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
      sentence: `A guild's first ${GUILD_TRIAL_DAYS} days are free, once per guild. The card is added at the start; cancel before day ${GUILD_TRIAL_DAYS + 1} and nothing is charged.`,
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
