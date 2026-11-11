// web/src/lib/billing/copy.ts
// Every visible string /premium, /premium/checkout and the account page's billing block
// use, in the site's own honest voice: no exclamation marks, no urgency, no invented
// discounts. Prices are never written here as literals -- they come from ./plans.
import { priceFor } from './plan-costs';
import { PLANS, formatPrice, type Interval, type PlanKey } from './plans';
import { sharedPremiumSupporterChip } from '../copy/shared/premium.generated';
import { grouped } from '../format/number';
import { uiCopy } from '../ui/copy';

export const checkoutCopy = {
  redirecting: 'Redirecting to secure checkout…',
  signInRequired: 'Sign in to continue.',
  canceled: 'Checkout was canceled. Nothing was charged.',
  notOpenYet: 'Purchases are not open yet. Nothing was charged. Check back soon.',
  genericFailure: 'That did not work; try again.',
  processing: 'Confirming your payment…',
  confirmed: 'Payment confirmed.',
  confirmationSlow: 'This can take a moment; refresh if it does not update.',
  retry: uiCopy.retry,
  whichGuild: 'Which guild?',
  /** Under Which guild? (player 9): the plan and price picked on /premium, `Raid guild · $200 a year`. */
  chosenPlan: (plan: PlanKey, interval: Interval, founder: boolean): string =>
    `${PLANS[plan].name} · ${formatPrice(priceFor(plan, interval, founder), interval)}`,
  guildPlanLabel: 'Guild',
  comebackGone: 'That resubscribe link is no longer available. Start again from the premium page.',
  verifiedMembers: (n: number): string => `${n} verified members`,
  fitsRaidGuild: (n: number): string => `Raid guild fits ${n} members`,
  noBillableGuild: 'You can subscribe a guild you lead or have claimed.',
  findYourGuild: 'Find your guild',
  guildsDidNotLoad: 'Your guilds did not load.',
  tryAgain: uiCopy.retry,
  /** The checkout Steps (spec 2026-11-08 section 6). */
  steps: { signIn: 'Sign in', pay: 'Pay with Stripe', on: (plan: string): string => `${plan} on` },
  signedInAs: (who: string): string => `Signed in as ${who}`,
  stripeLine: 'Stripe handles the card; we never see it.',
  paid: 'Stripe took the payment.',
  /** A guild checkout in its free trial took no payment (ruling K1). */
  trialSaved: (date: string): string => `Card saved. First charge ${date}.`,
  checkAgain: 'Check again',
  /** Above the next steps once Premium is on (ruling K2): `$40 a year · renews 4 Nov 2027 · receipt from Stripe by email`. */
  renewalLine: (price: string, date: string): string =>
    `${price} · renews ${date} · receipt from Stripe by email`,
  nextSim: 'Run a server sim',
  nextCoach: 'Upload a log for the coach',
  /** The confirmed panel's head once the plan is on (redesign review 2026-11-09). */
  planOn: (plan: string): string => `${plan} is on.`,
  /** The confirmed panel's plan row: what the account page holds. */
  seePlanSub: 'Renewal date, receipts and cancel',
  /** The confirmed panel's next rows, named for assistive tech. */
  nextLabel: 'What to do next',
};

export const billingBlockCopy = {
  notSubscribed: 'Not on Premium.',
  seePlans: 'See plans',
  manageBilling: 'Manage billing',
  renews: 'Renews',
  ends: 'Ends',
  pastDueTag: 'Payment failed',
  pastDueBanner: 'Your last payment failed. Update your card to keep access.',
  guildBilledBy: (name: string): string => `Billed by ${name}`,
  guildNotSubscribed: 'This guild is not on the guild plan yet.',
  subscribeTheGuild: 'Subscribe your guild',
};

/** "On · 4 runs this month"; plain "On" when an older /v1/me sent no count. */
function serverSimsOn(runsThisMonth?: number): string {
  if (runsThisMonth === undefined) return 'On';
  if (runsThisMonth === 0) return 'On · no runs yet this month';
  const noun = runsThisMonth === 1 ? 'run' : 'runs';
  return `On · ${grouped(runsThisMonth)} ${noun} this month`;
}

export const premiumBlockCopy = {
  rows: {
    serverSims: {
      label: 'Server sims',
      on: serverSimsOn,
      // Controller's 2026-10-27 persona round (player 13).
      off: 'In your browser',
    },
    history: {
      label: 'History charts',
      on: 'On',
      onWithMain: 'Open on your main character',
      off: 'Sparkline only',
    },
    /** The value is a function of the block's state: see account/notifications.ts `premiumRowValue`. */
    notifications: { label: 'Discord notifications' },
    supporter: { label: 'Supporter mark', on: 'Shown on your characters', off: 'Not shown' },
  },
  supporterChip: sharedPremiumSupporterChip,
  supporterTooltip: 'Supports Forever Sixty (Premium or guild plan)',
} as const;
