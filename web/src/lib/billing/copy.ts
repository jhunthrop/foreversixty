// web/src/lib/billing/copy.ts
// Every visible string /premium, /premium/checkout and the account page's billing block
// use, in the site's own honest voice: no exclamation marks, no urgency, no invented
// discounts. Prices are never written here as literals -- they come from ./plans.
import { PLANS, formatPrice, type Interval, type PlanKey } from './plans';

export const premiumCopy = {
  freeForeverHeading: 'What is free, forever',
  freeForeverBody: [
    'Planner, browser simulator, logs, live logging, rankings, the addon — always free, no ads.',
    'Nothing that comes from Battle.net is ever behind a paywall — your characters and your sign-in stay free no matter what.',
  ],
  freeForeverNote: 'A player who never pays never hits a wall that makes the site feel broken.',
  plansHeading: 'The two plans',
  premiumFeatures: [
    'Run sims on our servers (5,000 combinations, any precision, no tab left open).',
    'Two years of log retention instead of ninety days.',
    'Compare more than two builds at once.',
    'Character history charts over time.',
    'Discord notifications.',
    'A supporter mark on your profile.',
  ],
  guildFeatures: [
    'Everything in Premium, for everyone in the guild.',
    "Officer tools once they ship: the loot council helper, the raid readiness board, and the rest — a claimed guild's officers see what's live today on the guild's settings page.",
  ],
  noAds:
    'No ads. Not for anyone, paying or not. The site is paid for by Premium and the guild plan, not by ads, and that is true for everyone.',
  intervalLabel: { monthly: 'Monthly', yearly: 'Yearly' } as Record<Interval, string>,
  buyLabel: (plan: PlanKey, interval: Interval): string =>
    `Subscribe — ${formatPrice(PLANS[plan][interval], interval)}`,
  pricedName: (plan: PlanKey, interval: Interval): string =>
    `${PLANS[plan].name} — ${formatPrice(PLANS[plan][interval], interval)}`,
  faq: [
    {
      question: 'What happens if I cancel?',
      answer:
        'You keep everything through the end of the period you already paid for. After that, the account works exactly like a free one — nothing is deleted.',
    },
    {
      question: 'What happens to my long-retention logs if I stop?',
      answer:
        'They stay at the two-year retention for thirty days after your subscription ends, in case you resubscribe. After that, retention falls back to ninety days like every free account, and anything older than that is not kept.',
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
        'No. Not for anyone, paying or not. The site is paid for by Premium and the guild plan, not by ads, and that is not a perk you are buying — it is true for everyone.',
    },
    {
      question: 'Who am I paying?',
      answer:
        'COMMISH LLC, the company behind Forever Sixty. Stripe handles the payment; we never see or store your card.',
    },
  ],
};

export const checkoutCopy = {
  redirecting: 'Redirecting to secure checkout…',
  signInRequired: 'Sign in to continue.',
  notOpenYet: 'Purchases are not open yet. Check back soon — nothing was charged.',
  guildAlreadyOnPlan: "This guild is already on the plan. Manage its billing from the guild's settings page.",
  genericFailure: 'That did not work; try again.',
  processing: 'Confirming your payment…',
  confirmed: "You're all set.",
  confirmationSlow: 'This can take a moment; refresh if it does not update.',
  retry: 'Try again',
};

export const billingBlockCopy = {
  notSubscribed: 'Not on Premium.',
  seePlans: 'See plans',
  manageBilling: 'Manage billing',
  renews: 'Renews',
  ends: 'Ends',
  pastDueBanner: 'Your last payment failed. Update your card to keep access.',
  guildBilledBy: (name: string): string => `Billed by ${name}`,
};
