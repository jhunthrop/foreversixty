// web/src/lib/creators/copy.ts
// The public /creators page: the creator program's terms (owner-approved, design/specs/
// 2026-11-08-creator-program.md section 3e on the creator-program branch) and its apply path.
// The terms are numbers in PROGRAM_TERMS, written once; every sentence and the page's figures
// build from them. The program's server side (attribution, statements, Connect payouts, the
// apply form) is not on main yet, so Apply is an email to the support address, the one channel
// that exists today. The mocks read this module (design/mocks/creators_facts.mjs).

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

/** Where an application goes until the apply form ships with the program's server side. */
export const CREATORS_APPLY_EMAIL = 'help@foreversixty.gg';
const APPLY_SUBJECT = 'Creator program';

/** The mailto an Apply button opens: the support address with the subject filled in. */
export const CREATORS_APPLY_HREF = `mailto:${CREATORS_APPLY_EMAIL}?subject=${encodeURIComponent(APPLY_SUBJECT)}`;

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
    unit: 'a player',
    line: `then ${T.sharePercent}% of their Premium for ${T.shareMonths} months`,
  },
  /** The worked examples in the first section's lane (lib/creators/model.ts). */
  examplesLabel: 'One player, first year',
  examplesNote: 'Before Stripe fees.',
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
        { id: 'once', label: 'Each account', value: 'once' },
      ],
    },
    {
      section: 'payouts',
      label: 'Each month',
      figures: [
        { id: 'close', label: 'Closes', value: 'the 1st' },
        { id: 'paid', label: 'Paid', value: 'first week' },
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
        `${usd(T.bountyUsd)} for each player you send who comes back. A player comes back when they sign in with Battle.net, import at least one Forever character, and use the site again on a later day within ${T.returnDays} days of signing up.`,
        `Half of the Premium they pay, after Stripe’s fees and any refunds, for ${T.shareMonths} months from their first payment. Guild plans count when the officer who pays came through you.`,
        `Both add up: ${usd(T.bountyUsd)} once when a player comes back, then half of each of their payments for a year.`,
        'You also get Premium on your own account and the Raid guild plan for your guild while you are in the program.',
      ],
    },
    {
      id: 'how-players-count',
      title: 'How players count',
      paragraphs: [
        `You get a link of your own and a code. Your code gets viewers a free month of Premium, or a free month for their guild; they enter it when they subscribe, and a card is still needed. A player is yours if their first visit came through your link within ${T.attributionDays} days before they signed up, or if they use your code.`,
        'Each account counts once, for one creator. Accounts made before you joined, and your own account, never count.',
      ],
    },
    {
      id: 'payouts',
      title: 'Payouts',
      paragraphs: [
        `Each month closes on the 1st. We approve it during the first week and pay it through Stripe to a payout account you set up with them. We never see your bank details. A month under ${usd(T.minimumPayoutUsd)} carries over to the next one. Each creator has a monthly cap, ${usd(T.monthlyCapUsd)} unless we agree another.`,
        `A player counts as coming back up to ${T.returnDays} days after signing up, so the latest weeks fill in late. We never tell you who they are.`,
        'Stripe asks for your tax details when you set up payouts. If you leave the program, your free Premium ends, and the months you already earned are still paid.',
      ],
    },
    {
      id: 'disclosure',
      title: 'Saying it is sponsored',
      paragraphs: [
        'You earn money when people use your link or code, so you have to say so wherever you share it. Say “affiliate link” or “sponsored” on stream or in the video, and put the same words next to your link or code. The FTC and Twitch both require it, and so do these terms: the first time it is missing we warn you, and a second time ends your place in the program.',
      ],
    },
  ] satisfies readonly ProgramSection[],
  applySection: {
    id: 'apply',
    title: 'Apply',
    paragraphs: [
      `Email ${CREATORS_APPLY_EMAIL} with your Twitch or YouTube channel, roughly how many watch, and what you make about Forever. We read every application and answer by email.`,
    ],
  } satisfies ProgramSection,
  applyAction: 'Email your application',
} as const;
