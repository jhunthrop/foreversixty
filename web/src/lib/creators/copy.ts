// web/src/lib/creators/copy.ts
// Every word on the creator program's four pages (design/specs/2026-11-08-creator-program.md
// section 3): the owner's /admin/creators and /admin/creators/payouts, and the creator's
// /creators/onboarding and /creators/stats. The mocks read this module (creators_facts.mjs).

export const adminCreatorsCopy = {
  pageTitle: 'Creators',
  job: 'Who brought players, what they earned this month, and whether Stripe can pay them.',
  pageDescription: 'The owner’s creator program: players, earnings and payouts.',
  signedOut: 'Sign in with the owner’s account to see the creators.',
  notOwner: 'This page is for the site’s owner.',
  loadFailed: 'The creators didn’t load.',
  stripeOff: 'Stripe isn’t configured, so codes and payouts are off.',
  payoutsOff: 'Payouts are switched off on the server. Statements can be read but not paid.',
  tableTitle: 'Your creators',
  tableMeta: (count: number): string => `${count} ${count === 1 ? 'creator' : 'creators'}`,
  empty: 'No creators yet. Add the first one.',
  sub: (slug: string, code: string, since: string): string => `/go/${slug} · code ${code} · since ${since}`,
  columns: {
    clicks: 'Clicks',
    signUps: 'Signed up',
    returning: 'Returning',
    revenue: 'Premium net',
    owed: 'Owed so far',
  },
  allTime: 'all time',
  /** Under Owed so far: the open month it covers and the cap. */
  owedNote: (month: string, cap: string, capped: boolean): string =>
    capped ? `${month} · ${cap} cap reached` : `${month} · of ${cap} cap`,
  edit: 'Edit',
  editLabel: (name: string): string => `Edit ${name}`,
  statementsLink: 'Monthly statements',
  addTitle: 'Add a creator',
  editTitle: (name: string): string => `Edit ${name}`,
  formJob: 'The link and the code are made once and can’t change after.',
  fields: {
    name: { label: 'Name', hint: 'What the creator goes by on stream.', placeholder: 'e.g. Grimbrew' },
    channel: { label: 'Channel', hint: 'Their Twitch or YouTube address.', placeholder: 'https://twitch.tv/…' },
    slug: { label: 'Link', hint: 'Becomes foreversixty.gg/go/your-slug.', placeholder: 'e.g. grimbrew' },
    code: { label: 'Promo code', hint: 'Capitals and digits, 4 to 20. Players get the creator discount on a first payment.', placeholder: 'e.g. GRIMBREW' },
    startsOn: { label: 'Starts on', hint: 'Players who signed up before this day never count.', placeholder: '2026-11-08' },
    cap: { label: 'Monthly cap', hint: 'Dollars a month at most, like 500.', placeholder: '500' },
  },
  add: 'Add creator',
  save: 'Save',
  saving: 'Saving',
  cancelEdit: 'Cancel',
  added: (name: string): string => `${name} is in. Copy their invite from Edit.`,
  saved: 'Saved.',
  saveFailed: 'That wasn’t saved; try again.',
  invite: 'Copy onboarding invite',
  inviteCopied: 'Invite copied. It works once, for 14 days.',
  inviteFailed: 'The invite wasn’t made; try again.',
  refresh: 'Refresh Stripe status',
  refreshed: 'Stripe status read just now.',
  refreshFailed: 'Stripe didn’t answer; try again.',
  fixed: (slug: string, code: string): string => `/go/${slug} · code ${code}`,
} as const;

export const stripeStatusCopy = {
  ready: 'Payouts ready',
  onboarding: 'Onboarding',
  needs_info: 'Needs info',
  none: 'Not started',
} as const;

export const payoutsCopy = {
  pageTitle: 'Statements',
  job: 'Each creator’s month. Nothing is paid until you approve it.',
  pageDescription: 'The owner’s monthly creator statements and payouts.',
  payoutsOff: 'Payouts are switched off on the server. Statements can be read but not paid.',
  monthLabel: 'Month',
  openMonth: (label: string): string => `${label} (open)`,
  loadFailed: 'The statements didn’t load.',
  tableTitle: (count: number): string => `${count} ${count === 1 ? 'creator' : 'creators'}`,
  tableMeta: (total: string): string => `${total} to pay`,
  empty: 'No creators yet.',
  back: 'Creators',
  columns: {
    bounties: 'Returning',
    share: 'Share',
    refunds: 'Refunds',
    cap: 'Over cap',
    total: 'Total',
  },
  bounties: (count: number, each: string): string => `${count} × ${each}`,
  status: {
    open: 'Month still open',
    ready: 'Ready to approve',
    nothing: 'Nothing owed',
    blocked: 'Stripe can’t pay yet: onboarding not finished',
    paid: (date: string, transfer: string): string => `Paid ${date} · transfer ${transfer}`,
    settled: (date: string): string => `Settled ${date} · nothing to transfer`,
  },
  notReady: 'Stripe not ready',
  approve: (amount: string): string => `Approve ${amount}`,
  approving: 'Paying',
  confirmTitle: (name: string, amount: string, month: string): string => `Pay ${name} ${amount} for ${month}?`,
  confirmBody: (amount: string): string =>
    `Stripe transfers ${amount} to their account now. This can’t be undone here.`,
  confirmPay: (amount: string): string => `Pay ${amount}`,
  cancel: 'Cancel',
  approveFailed: (reason: string): string => `Stripe didn’t take the transfer: ${reason}. Nothing was paid; try again.`,
  lines: 'Lines',
  linesLabel: (name: string): string => `Show the lines for ${name}`,
  linesTitle: (name: string): string => `Lines · ${name}`,
  linesEmpty: 'No lines this month.',
  line: {
    bounty: (player: number, date: string): string => `Player #${player} · returning · ${date}`,
    share: (player: number, date: string): string => `Player #${player} · paid ${date}`,
    refund: (player: number, date: string): string => `Player #${player} · refund ${date}`,
  },
} as const;

export const onboardingCopy = {
  pageTitle: 'Creator payouts',
  pageDescription: 'Set up payouts for the Forever Sixty creator program.',
  inviteTitle: 'Your invite',
  stripeTitle: 'Stripe',
  inviteSignedOut: 'Sign in with Battle.net to accept your invite to the creator program.',
  inviteJoin: (name: string): string => `You’re joining the creator program as ${name}.`,
  accept: 'Accept and set up payouts',
  accepting: 'Accepting',
  stripeLead: 'Stripe handles your payout details. We never see your bank account.',
  continue: 'Continue to Stripe',
  continuing: 'Opening Stripe',
  doneReady: 'You’re set up. Payouts arrive after each month is approved.',
  doneNotReady: 'Stripe still needs a few details.',
  checking: 'Checking with Stripe',
  statsLink: 'Your stats',
  notCreator: 'This page is for creators in the program. Ask us for an invite.',
  badInvite: 'That invite has expired or was already used. Ask us for a new one.',
  failed: 'Stripe didn’t answer; try again.',
} as const;

export const creatorStatsCopy = {
  pageDescription: 'Your players and payouts in the Forever Sixty creator program.',
  meta: (slug: string, code: string): string => `foreversixty.gg/go/${slug} · code ${code}`,
  loadFailed: 'Your stats didn’t load.',
  signedOut: 'Sign in with Battle.net to see your creator stats.',
  summaryLabel: 'Your figures',
  figures: {
    month: 'Earned so far',
    players: 'Players',
    returning: 'Returning',
    paying: 'Paying',
  },
  monthNote: (month: string, cap: string): string => `${month}, estimated, ${cap} cap max`,
  allTime: 'all time',
  payingNote: 'paying now',
  playersTitle: 'Your players',
  playersMeta: 'by week · no names, by design',
  playersEmpty: 'No players yet. Share your link or your code.',
  /** One week of sign-ups. A week still running shows only its sign-ups (spec section 1, Privacy). */
  week: (week: string, signedUp: number, returning: number | null, paying: number | null): string =>
    returning === null || paying === null
      ? `Week of ${week}: ${signedUp} signed up so far`
      : `Week of ${week}: ${signedUp} signed up · ${returning} returning · ${paying} paying`,
  weeksCapped: 'Showing the last 26 weeks.',
  payoutsTitle: 'Payouts',
  payoutsEmpty: 'No payouts yet. Each month is paid after it closes and is approved.',
  setup: 'Set up payouts',
} as const;
