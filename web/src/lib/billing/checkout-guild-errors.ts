// web/src/lib/billing/checkout-guild-errors.ts
// The busy labels and error lines of the checkout guild step (design/specs/2026-10-27-subscribe-and-join.md
// §9.1). Error lines are written from the API's error codes, never its sentence.

/** One claim attempt per account in this many days (the API's claim cooldown). */
export const CLAIM_COOLDOWN_DAYS = 30;

export const checkoutGuildBusy = {
  checkout: 'Redirecting to secure checkout…',
  claiming: 'Claiming…',
  starting: 'Starting…',
  confirming: 'Confirming…',
} as const;

export const checkoutGuildClaimErrors = {
  alreadyClaimed: (guild: string): string => `Someone claimed ${guild} a moment ago.`,
  pending: (guild: string): string => `An officer started a claim on ${guild} a moment ago.`,
  notEligible: (guild: string): string => `Your character in ${guild} is not an officer or the guild master.`,
  noBattlenet: 'Claiming as guild master needs your Battle.net account linked.',
  rateLimited: `You made a claim in the last ${CLAIM_COOLDOWN_DAYS} days.`,
  holdsAnother: "You already hold another guild's claim.",
  sameAccount: 'You started this claim, so a second officer confirms it.',
  generic: 'That did not work; try again.',
} as const;

export const checkoutGuildCheckoutErrors = {
  notOpen: 'Purchases are not open yet. Check back soon — nothing was charged.',
  alreadyOnPlan: (guild: string): string => `${guild} is already on a plan.`,
  generic: 'That did not work; try again.',
  tryAgain: 'Try again',
} as const;

export const checkoutGuildPortalError = 'Billing did not open. Try again.';
