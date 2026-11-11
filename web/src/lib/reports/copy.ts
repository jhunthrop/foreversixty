// web/src/lib/reports/copy.ts
// Every visible string RecentReports.svelte and the /logs page use, in the site's own
// honest-copy voice: never promises what is not built, no exclamation marks.
import { withRetry } from '../ui/copy';
export const recentReportsCopy = {
  heading: 'Recent public reports',
  /** The placeholder's screen-reader name while the feed loads: never the word "Loading" in page text. */
  placeholderLabel: 'Recent reports',
  failed: 'Recent reports did not load.',
  older: 'Older reports',
} as const;

/** The page head's own copy (logs spec 2026-11-07 section 4; directions b1: the title alone). */
export const logsBandCopy = {
  h1: 'Logs',
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
  companionFirst: 'Press Sign in with Battle.net in the companion. A code is the fallback.',
  success: (deviceName: string): string =>
    `${deviceName} paired. It starts uploading as soon as you are logging.`,
  action: 'Get a pairing code',
  failed: withRetry('The pairing code did not load.'),
  codeLine: (minutes: number): string =>
    `Type this into the companion within ${minutes} minutes. It pairs one device and cannot be reused.`,
  /** /setup's row 3 while a code is out and unused. */
  waiting: 'Waiting for the companion to use the code',
} as const;
