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
  connectDenied:
    'Stripe payouts key missing or lacks Connect permissions. Payouts can’t be made until it’s set.',
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
    channel: {
      label: 'Channel',
      hint: 'Their Twitch or YouTube address.',
      placeholder: 'https://twitch.tv/…',
    },
    slug: { label: 'Link', hint: 'Becomes foreversixty.gg/go/your-slug.', placeholder: 'e.g. grimbrew' },
    code: {
      label: 'Creator code',
      hint: 'Capitals and digits, 4 to 20. It gets viewers a free month of Premium, or a free month for their guild.',
      placeholder: 'e.g. GRIMBREW',
    },
    startsOn: {
      label: 'Starts on',
      hint: 'Players who signed up before this day never count.',
      placeholder: '2026-11-08',
    },
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
  confirmTitle: (name: string, amount: string, month: string): string =>
    `Pay ${name} ${amount} for ${month}?`,
  confirmBody: (amount: string): string =>
    `Stripe transfers ${amount} to their account now. This can’t be undone here.`,
  confirmPay: (amount: string): string => `Pay ${amount}`,
  cancel: 'Cancel',
  approveFailed: (reason: string): string =>
    `Stripe didn’t take the transfer: ${reason}. Nothing was paid; try again.`,
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
  monthNote: (month: string, cap: string): string =>
    `${month}, an estimate until the month closes · ${cap} monthly cap`,
  allTime: 'all time',
  payingNote: 'paying now',
  playersTitle: 'Your players',
  playersMeta: 'by week · the latest weeks fill in over 14 days · no names',
  playersEmpty: 'No players yet. Share your link or your code.',
  /** One week of sign-ups. A week still running shows only its sign-ups (spec section 1, Privacy). */
  week: (week: string, signedUp: number, returning: number | null, paying: number | null): string =>
    returning === null || paying === null
      ? `Week of ${week}: ${signedUp} signed up so far`
      : `Week of ${week}: ${signedUp} signed up · ${returning} returning · ${paying} paying`,
  weeksCapped: 'Showing the last 26 weeks.',
  payoutsTitle: 'Payouts',
  /** Opens the creator's Stripe Express dashboard (payouts, bank details, the 1099) by a fresh login link. */
  dashboard: 'Payouts and tax forms',
  dashboardOpening: 'Opening Stripe',
  dashboardFailed: 'Stripe didn’t open; try again.',
  /** The profile and account menu item for a signed-in creator. */
  menuItem: 'Creator stats',
  payoutsEmpty: 'No payouts yet. Each month is paid after it closes and is approved.',
  setup: 'Set up payouts',
} as const;

/** One section of the public /creators page: a heading and its paragraphs. */
export interface ProgramSection {
  id: string;
  title: string;
  paragraphs: readonly string[];
}

/** The public /creators page (spec section 3e). Plain words: what a creator gets and what they owe us. */
export const programCopy = {
  pageTitle: 'Creators',
  pageDescription:
    'The Forever Sixty creator program: $2 for each player who comes back, half of their Premium for a year, paid monthly through Stripe.',
  lead: 'If you stream or make videos about World of Warcraft: Forever, you can earn from the players you send us.',
  apply: 'Apply',
  footerLabel: 'Creators',
  sections: [
    {
      id: 'what-you-get',
      title: 'What you get',
      paragraphs: [
        '$2 for each player you send who comes back. A player comes back when they sign in with Battle.net, import at least one Forever character, and use the site again on a later day within 14 days of signing up.',
        'Half of the Premium they pay, after Stripe’s fees and any refunds, for 12 months from their first payment. Guild plans count when the officer who pays came through you.',
        'Both add up: $2 once when a player comes back, then half of each of their payments for a year.',
        'Premium on your own account and the Raid guild plan for your guild while you are in the program.',
      ],
    },
    {
      id: 'how-players-count',
      title: 'How players count',
      paragraphs: [
        'You get a link of your own and a code. Your code gets viewers a free month of Premium, or a free month for their guild; they enter it when they subscribe, and a card is still needed. A player is yours if their first visit came through your link within 30 days before they signed up, or if they use your code.',
        'Each account counts once, for one creator. Accounts made before you joined, and your own account, never count.',
      ],
    },
    {
      id: 'payouts',
      title: 'Payouts',
      paragraphs: [
        'We add up each month after it ends and, once we approve it, pay it through Stripe to a payout account you set up with them. We never see your bank details. Each creator has a monthly cap, $500 unless we agree another.',
        'You can see your players by week, and what you have earned, on your own stats page. A player counts as coming back up to 14 days after signing up, so the latest weeks fill in late. We never show you who they are.',
        'Stripe asks for your tax details when you set up payouts. If you leave the program, your free Premium ends, and the months you already earned are still paid.',
      ],
    },
    {
      id: 'disclosure',
      title: 'Saying it is sponsored',
      paragraphs: [
        'You earn money when people use your link or code, so you have to say so wherever you share it. Say “affiliate link” or “sponsored” on stream or in the video, and put the same words next to your link or code. The FTC and Twitch both require it.',
      ],
    },
  ] satisfies readonly ProgramSection[],
  applyTitle: 'Apply',
  applyLead: 'Tell us about your channel. We read every application.',
} as const;

export const PLATFORMS = ['twitch', 'youtube'] as const;
export type Platform = (typeof PLATFORMS)[number];

/** /creators/apply (spec section 3f). */
export const applyCopy = {
  pageTitle: 'Apply to the creator program',
  pageDescription: 'Apply to the Forever Sixty creator program.',
  formTitle: 'Your channel',
  signedOut: 'Sign in with Battle.net to apply.',
  fields: {
    channel: {
      label: 'Channel',
      hint: 'Your Twitch or YouTube address.',
      placeholder: 'https://twitch.tv/…',
    },
    platform: { label: 'Main platform' },
    audience: {
      label: 'Viewers or subscribers',
      hint: 'Roughly. Leave it empty if you would rather not say.',
      placeholder: 'e.g. 1200',
    },
    about: { label: 'What do you make about Forever?' },
    email: { label: 'Contact email', hint: 'Optional. Only we see it.', placeholder: 'you@example.com' },
  },
  platforms: { twitch: 'Twitch', youtube: 'YouTube' } satisfies Record<Platform, string>,
  consent: 'We review every application; we’ll reply on your Forever Sixty account.',
  submit: 'Send application',
  sending: 'Sending',
  sendFailed: 'That wasn’t sent; try again.',
  counter: (used: number, max: number): string => `${used} of ${max} characters`,
  disclose: 'I’ll say it’s an affiliate link or sponsored wherever I share my link or code.',
  discloseNeeded: 'Tick this to apply: the FTC and Twitch require it.',
  statusTitle: 'Your application',
  pending: (date: string): string => `Sent ${date}. We’ll reply here.`,
  pendingTag: 'Pending',
  approved: 'You’re in. Set up payouts so we can pay you.',
  approvedTag: 'Approved',
  setup: 'Set up payouts',
  declined: 'Not this time. Thank you for applying.',
  declinedTag: 'Not this time',
  alreadyCreator: 'You’re already in the creator program.',
  statsLink: 'Your stats',
} as const;

/** The owner's Requests section on /admin/creators (spec section 3a). */
export const requestsCopy = {
  title: 'Requests',
  meta: (count: number): string => `${count} waiting`,
  empty: 'No requests waiting.',
  sub: (platform: string, size: string, date: string): string => `${platform} · ${size} · ${date}`,
  sizeUnknown: 'size not given',
  size: (n: string): string => `about ${n}`,
  approve: 'Approve',
  decline: 'Decline',
  declined: 'Declined.',
  approveTitle: (channel: string): string => `Approve ${channel}`,
  approveJob: 'Their account is already bound; no invite is needed.',
  /** The owner's account-menu badge while requests wait. */
  badge: (count: number): string => `${count} creator ${count === 1 ? 'request' : 'requests'}`,
} as const;
