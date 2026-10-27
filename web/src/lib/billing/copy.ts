// web/src/lib/billing/copy.ts
// Every visible string /premium, /premium/checkout and the account page's billing block
// use, in the site's own honest voice: no exclamation marks, no urgency, no invented
// discounts. Prices are never written here as literals -- they come from ./plans.

export const checkoutCopy = {
  redirecting: 'Redirecting to secure checkout…',
  signInRequired: 'Sign in to continue.',
  canceled: 'Checkout was canceled. Nothing was charged.',
  notOpenYet: 'Purchases are not open yet. Check back soon — nothing was charged.',
  guildAlreadyOnPlan: "This guild is already on the plan. Manage its billing from the guild's settings page.",
  genericFailure: 'That did not work; try again.',
  processing: 'Confirming your payment…',
  confirmed: "You're all set.",
  confirmationSlow: 'This can take a moment; refresh if it does not update.',
  retry: 'Try again',
  whichGuild: 'Which guild?',
  verifiedMembers: (n: number): string => `${n} verified members`,
  fitsRaidGuild: (n: number): string => `Raid guild fits ${n} members`,
  noBillableGuild: 'You can subscribe a guild you lead or have claimed.',
  findYourGuild: 'Find your guild',
  guildsDidNotLoad: 'Your guilds did not load.',
  tryAgain: 'Try again',
};

export const billingBlockCopy = {
  notSubscribed: 'Not on Premium.',
  seePlans: 'See plans',
  manageBilling: 'Manage billing',
  renews: 'Renews',
  ends: 'Ends',
  pastDueBanner: 'Your last payment failed. Update your card to keep access.',
  guildBilledBy: (name: string): string => `Billed by ${name}`,
  guildNotSubscribed: 'This guild is not on the guild plan yet.',
  subscribeTheGuild: 'Subscribe the guild',
};

/** "On · 4 runs this month"; plain "On" when an older /v1/me sent no count. */
function serverSimsOn(runsThisMonth?: number): string {
  if (runsThisMonth === undefined) return 'On';
  if (runsThisMonth === 0) return 'On · no runs yet this month';
  const noun = runsThisMonth === 1 ? 'run' : 'runs';
  return `On · ${runsThisMonth.toLocaleString('en-US')} ${noun} this month`;
}

export const premiumBlockCopy = {
  rows: {
    serverSims: {
      label: 'Server sims',
      on: serverSimsOn,
      off: 'Browser lane only',
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
  supporterChip: 'Supporter',
  supporterTooltip: 'Supports Forever Sixty (Premium or guild plan)',
} as const;
