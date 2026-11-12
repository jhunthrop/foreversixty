// web/src/lib/creators/copy.ts
// Every word of the creator program's pages (design/specs/2026-11-08-creator-program.md section 3):
// the public /creators page (its terms are numbers in PROGRAM_TERMS, written once; every sentence
// and figure builds from them), the owner's /admin/creators and its statements, the creator's
// /creators/onboarding and /creators/stats, and /creators/apply. The mocks read this module
// (design/mocks/creators_facts.mjs). Failure lines end in the shared retry word (lib/ui/copy.ts);
// busy words are the shared ones or name the call ("Opening Stripe…").
import { busyCopy, withRetry } from '../ui/copy';

/** The program's numbers, one place (the terms the sentences below are built from). */
export const PROGRAM_TERMS = {
  /** Dollars for each player who comes back. */
  bountyUsd: 2,
  /** Days after signing up a player has to come back on a later day. */
  returnDays: 14,
  /** Percent of a player's Premium (after Stripe's fees and refunds) the creator earns. */
  sharePercent: 50,
  /** Months the share runs from the player's first payment. */
  shareMonths: 12,
  /** Days before signing up that a first visit through the creator's link still counts. */
  attributionDays: 30,
  /** A month under this many dollars carries over. */
  minimumPayoutUsd: 10,
  /** Each creator's monthly cap in dollars, unless agreed otherwise. */
  monthlyCapUsd: 500,
} as const;

/** Where Apply goes: the signed-in form (section 3f). Signed out, that page asks for Battle.net
 *  sign-in and comes back to itself. */
export const CREATORS_APPLY_PATH = '/creators/apply';

/** One section of the public /creators page: a heading and its paragraphs. */
export interface ProgramSection {
  readonly id: string;
  readonly title: string;
  readonly paragraphs: readonly string[];
}

/** One figure in a section's lane: a label and its value with the unit. */
export interface ProgramFigure {
  readonly id: string;
  readonly label: string;
  readonly value: string;
}

/** A section's figures beside its prose: the lane's label and its figures. */
export interface ProgramLane {
  readonly section: string;
  readonly label: string;
  readonly figures: readonly ProgramFigure[];
}

const T = PROGRAM_TERMS;
const usd = (n: number): string => `$${n}`;

/** The public /creators page. Plain words: what a creator gets and what they owe us. */
export const programCopy = {
  pageTitle: 'Creators',
  pageDescription: `The Forever Sixty creator program: ${usd(T.bountyUsd)} for each player who comes back, half of their Premium for a year, paid monthly through Stripe.`,
  lead: 'If you stream or make videos about World of Warcraft: Forever, you can earn from the players you send us.',
  apply: 'Apply',
  contentsLabel: 'On this page',
  /** The head's answer: the bounty as the figure, the share on its one line. */
  answer: {
    value: usd(T.bountyUsd),
    unit: 'a returning player',
    line: `then ${T.sharePercent}% of their Premium for ${T.shareMonths} months`,
  },
  /** The worked examples in the first section's lane (lib/creators/model.ts). */
  examplesLabel: 'One player, first year',
  examplesNote: 'Before Stripe’s fees; you are paid on what is left after them.',
  lanes: [
    {
      section: 'what-you-get',
      label: 'The terms',
      figures: [
        { id: 'bounty', label: 'Each returning player', value: usd(T.bountyUsd) },
        { id: 'share', label: 'Their Premium', value: `${T.sharePercent}% for ${T.shareMonths} months` },
      ],
    },
    {
      section: 'how-players-count',
      label: 'Counting',
      figures: [
        { id: 'window', label: 'Link window', value: `${T.attributionDays} days` },
        { id: 'return', label: 'Coming back', value: `within ${T.returnDays} days` },
        { id: 'once', label: 'Each account', value: 'once, for one creator' },
      ],
    },
    {
      section: 'payouts',
      label: 'Each month',
      figures: [
        { id: 'close', label: 'Closes', value: 'the 1st' },
        { id: 'approved', label: 'Approved', value: 'by the 7th' },
        { id: 'minimum', label: 'Minimum', value: usd(T.minimumPayoutUsd) },
        { id: 'cap', label: 'Cap', value: usd(T.monthlyCapUsd) },
      ],
    },
  ] satisfies readonly ProgramLane[],
  sections: [
    {
      id: 'what-you-get',
      title: 'What you get',
      paragraphs: [
        `${usd(T.bountyUsd)} for each player you send who comes back. A player comes back when they sign in with Battle.net, import at least one Forever character, and sign in again on a later day within ${T.returnDays} days of signing up.`,
        `Half of the Premium they pay, after Stripe’s fees and any refunds, for ${T.shareMonths} months from their first payment. Guild plans count when the officer who pays came through you.`,
        `Both add up: ${usd(T.bountyUsd)} once when a player comes back, then half of each of their payments for a year.`,
        'You also get Premium on your own account and the Raid guild plan for your guild while you are in the program.',
      ],
    },
    {
      id: 'how-players-count',
      title: 'How players count',
      paragraphs: [
        `You get a link of your own and a code. Viewers enter your code at checkout and get a free month of Premium, or a free month for their guild; a card is still needed. A player is yours if their first visit came through your link within ${T.attributionDays} days before they signed up, or if they use your code.`,
        'Each account counts once, for one creator: a player who came through your link and later types another creator’s code stays yours. Accounts made before you joined, and your own account, never count.',
      ],
    },
    {
      id: 'payouts',
      title: 'Payouts',
      paragraphs: [
        `Each month closes on the 1st. We approve it by the 7th, and the moment we do Stripe moves it to a payout account you set up with them. We never see your bank details. A month under ${usd(T.minimumPayoutUsd)} carries over to the next one. Each creator has a monthly cap, ${usd(T.monthlyCapUsd)} unless we agree another; anything over the cap is not carried over.`,
        `A player counts as coming back up to ${T.returnDays} days after signing up, so the latest weeks fill in late. We never tell you who they are.`,
        'Stripe asks for your tax details when you set up payouts. If you leave the program, your free Premium ends, and the months you already earned are still paid.',
      ],
    },
    {
      id: 'disclosure',
      title: 'Saying it is sponsored',
      paragraphs: [
        'You earn money when people use your link or code, so you have to say so wherever you share it. Say “affiliate link” or “sponsored” on stream or in the video, and put the same words next to your link or code. The FTC and Twitch both require it, and so do these terms: the first time it is missing we warn you, and a second time can end your place in the program.',
      ],
    },
  ] satisfies readonly ProgramSection[],
  applySection: {
    id: 'apply',
    title: 'Apply',
    paragraphs: [
      'Tell us your Twitch or YouTube channel, roughly how many watch, and what you make about Forever. We read every application and reply on your Forever Sixty account.',
    ],
  } satisfies ProgramSection,
  applyAction: 'Apply',
} as const;

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
  tableMeta: (count: number): string => `${count} ${count === 1 ? 'creator' : 'creators'} · counts all time`,
  empty: 'No creators yet. Add the first one.',
  sub: (slug: string, code: string, since: string): string => `/go/${slug} · code ${code} · since ${since}`,
  columns: {
    clicks: 'Clicks',
    signUps: 'Signed up',
    returning: 'Returning',
    revenue: 'Premium net',
    owed: 'Owed so far',
  },
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
    name: { label: 'Name', placeholder: 'e.g. Grimbrew' },
    channel: { label: 'Channel', placeholder: 'https://twitch.tv/…' },
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
    cap: { label: 'Monthly cap', placeholder: '500' },
  },
  add: 'Add creator',
  save: 'Save',
  saving: busyCopy.saving,
  cancelEdit: 'Cancel',
  added: (name: string): string => `${name} is in. Copy their invite from Edit.`,
  saved: 'Saved.',
  saveFailed: withRetry('That wasn’t saved'),
  invite: 'Copy onboarding invite',
  inviting: 'Making the invite…',
  /** The button's word for about two seconds after the copy (the shared Copied rule). */
  inviteCopiedShort: 'Invite copied',
  inviteCopied: 'Invite copied. It works once, for 14 days.',
  inviteFailed: withRetry('The invite wasn’t made'),
  refresh: 'Refresh Stripe status',
  refreshing: 'Asking Stripe…',
  refreshed: 'Stripe status read just now.',
  refreshFailed: withRetry('Stripe didn’t answer'),
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
    carry: 'Under the $10 minimum: it carries over to next month',
    blocked: 'Stripe can’t pay yet: onboarding not finished',
    paid: (date: string, transfer: string): string => `Paid ${date} · transfer ${transfer}`,
    settled: (date: string): string => `Settled ${date} · nothing to transfer`,
  },
  notReady: 'Stripe not ready',
  approve: (amount: string): string => `Approve ${amount}`,
  approving: 'Paying…',
  confirmTitle: (name: string, amount: string, month: string): string =>
    `Pay ${name} ${amount} for ${month}?`,
  confirmBody: (amount: string): string =>
    `Stripe transfers ${amount} to their account now. This can’t be undone here.`,
  confirmPay: (amount: string): string => `Pay ${amount}`,
  balanceChecking: 'Checking your Stripe balance…',
  balanceLine: (available: string, amount: string): string =>
    `Available in Stripe: ${available}. This payout: ${amount}.`,
  balanceShort:
    'Not enough in your Stripe balance. Add funds in Stripe (Balances → Add to balance) or wait for subscription payments to clear.',
  balanceFailed: withRetry('Your Stripe balance didn’t load'),
  transferShort:
    'Stripe didn’t have enough in your balance for this payout. Nothing was paid; approve it again once the balance covers it.',
  cancel: 'Cancel',
  approveFailed: (reason: string): string =>
    withRetry(
      reason === ''
        ? 'Stripe didn’t take the transfer. Nothing was paid'
        : `Stripe didn’t take the transfer: ${reason}. Nothing was paid`,
    ),
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
  accepting: 'Accepting…',
  stripeLead: 'Stripe handles your payout details. We never see your bank account.',
  continue: 'Continue to Stripe',
  continuing: 'Opening Stripe…',
  doneReady: 'You’re set up. Payouts arrive after each month is approved.',
  doneNotReady: 'Stripe still needs a few details.',
  checking: 'Checking with Stripe…',
  statsLink: 'Your stats',
  notCreator: 'This page is for creators in the program. Ask us for an invite.',
  badInvite: 'That invite has expired or was already used. Ask us for a new one.',
  failed: withRetry('Stripe didn’t answer'),
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
  /** Under Earned so far: the open month (an estimate until it closes) and its cap. */
  monthNote: (month: string, cap: string): string => `${month} so far · of ${cap} cap`,
  returningNote: (share: string): string => `${share} of players`,
  payingNote: (share: string): string => `${share} of returning`,
  playersTitle: 'Your players',
  playersMeta: 'no names, by design',
  columns: { signedUp: 'Signed up', returning: 'Returning', paying: 'Paying' },
  weekTitle: (week: string): string => `Week of ${week}`,
  filling: 'still filling',
  playersEmpty: 'No players yet. Share your link or your code.',
  weeksCapped: 'Showing the last 26 weeks.',
  payoutsTitle: 'Payouts',
  /** The last closed month and the day it is approved by (the 7th of the open month). */
  closedMonth: (month: string, by: string): string => `${month} closed on the 1st. We approve it by ${by}.`,
  /** The same month once its approval day has passed. */
  closedMonthNow: (month: string): string => `${month} closed on the 1st. We are approving it now.`,
  paidSoFar: (amount: string): string => `${amount} paid so far`,
  /** Opens the creator's Stripe Express dashboard (payouts, bank details, the 1099) by a fresh login link. */
  dashboard: 'Payouts and tax forms',
  dashboardOpening: 'Opening Stripe…',
  dashboardFailed: withRetry('Stripe didn’t open'),
  /** The profile and account menu item for a signed-in creator. */
  menuItem: 'Creator stats',
  setup: 'Set up payouts',
} as const;

/** One section of the public /creators page: a heading and its paragraphs. */

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
  consent: 'We review every application and reply on your Forever Sixty account.',
  submit: 'Send application',
  sending: busyCopy.sending,
  sendFailed: withRetry('That wasn’t sent'),
  counter: (used: number, max: number): string => `${used} of ${max} characters`,
  disclose: 'I’ll say it’s an affiliate link or sponsored wherever I share my link or code.',
  discloseNeeded: 'Tick this to apply: the FTC and Twitch require it.',
  statusTitle: 'Your application',
  pending: (date: string): string => `Sent ${date}. We’ll reply here.`,
  pendingTag: 'Pending',
  approved: 'You’re in. Set up payouts so we can pay you.',
  approvedNext: 'Your link and code are on your stats page.',
  approvedTag: 'Approved',
  setup: 'Set up payouts',
  declined: 'Not this time. Thank you for applying.',
  reapplyOn: (date: string): string => `You can apply again from ${date}.`,
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
