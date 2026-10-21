// web/src/lib/reports/copy.ts
// Every visible string RecentReports.svelte and the /logs page use, in the site's own
// honest-copy voice: never promises what is not built, no exclamation marks.
export const recentReportsCopy = {
  heading: 'Recent public reports',
  /** The placeholder's screen-reader name while the feed loads: never the word "Loading" in page text. */
  placeholderLabel: 'Recent reports',
  failed: 'Recent reports did not load.',
  empty: 'No public reports yet. The first raid logs land in December; dungeon logs are welcome now.',
  older: 'Older reports',
} as const;

/** The page band's own copy (logs spec 2026-10-11 section 2.2): eyebrow, h1 and the one job line. */
export const logsBandCopy = {
  eyebrow: 'Combat logs',
  h1: 'Logs',
  jobLine: 'Upload a combat log or log live, then read every pull.',
} as const;

/** /logs' companion column: the pairing island, and a pointer to /setup for the download and
 *  /combatlog steps instead of repeating them. The pairing labels name what the code is for. */
export const logsCompanionCopy = {
  pointer: 'Log live while you raid.',
  /** The signed-out card's one button, to the companion's own step on /setup. */
  pointerLink: 'Get the companion',
  pointerHref: '/setup#companion',
  pairLabel: 'Pair it with your account',
  pairAnotherLabel: 'Pair another device',
} as const;

/** `Your reports`' guild tab strip (§4.C.1), one new string beside `myReportsCopy`. */
export const logsCopy = {
  guildEmpty: (guildName: string): string => `No ${guildName} reports yet.`,
} as const;

/** The companion panel's pairing lines (§4.D.2): the success line replaces the code block in
 *  place once a new device appears; signed out, one plain sentence and no control. */
export const pairingCopy = {
  signedOut: 'Live logging needs a signed-in account.',
  /** One line above the code button: the companion's own Battle.net sign-in is the way in. */
  companionFirst:
    'Using the companion? Press Sign in with Battle.net in it. The code below is only for the fallback.',
  success: (deviceName: string): string =>
    `${deviceName} paired. It starts uploading as soon as you are logging.`,
} as const;
