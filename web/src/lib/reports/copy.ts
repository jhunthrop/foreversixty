// web/src/lib/reports/copy.ts
// Every visible string RecentReports.svelte and /logs' framing line use, in the site's own
// honest-copy voice: never promises what is not built, no exclamation marks.
export const recentReportsCopy = {
  heading: 'Recent public reports',
  failed: 'Recent reports did not load.',
  empty: 'No public reports yet. The first raid logs land in December; dungeon logs are welcome now.',
  older: 'Older reports',
} as const;

/** /logs' one line for a visitor who is not raiding yet, spec section 4's exact wording. */
export const logsFraming =
  'Logs are for group content at any level: a dungeon run logs the same way a raid does.';

/** /logs' companion column, spec section 3.4: it keeps the pairing island but points to
 *  /setup for the download and /combatlog steps instead of repeating them. */
export const logsCompanionCopy = {
  pointer: 'Downloads and the in-game /combatlog step are on the setup page.',
  pointerLink: 'Get set up',
} as const;

/** The header band's own copy (logs landing spec 2026-10-04 §4.A/§6). The two anchors in
 *  `description` stay `.text-nav underline`, the live page's existing link style. */
export const logsHeroCopy = {
  eyebrow: 'Logs',
  /** The sentence's lead clause; the two trailing anchors (`Log live with the companion`,
   *  `upload a log file`) are markup in LogsHero.svelte, not plain copy, the same way the
   *  live page's own intro paragraph always kept its two links out of a copy constant. */
  description: 'Every pull, ranked: damage, healing, deaths, buffs, casts and threat.',
  openButton: 'Open this report',
  /** The sample hero's own line in the per-player hook's place (§4.A.1): no fetch, states
   *  the same capability honestly for the highest-traffic, anonymous case. */
  sampleHookLine: "Open it to see every player's gear next to the planner and the simulator.",
  failed: 'Logs did not load.',
} as const;

/** `Your reports`' guild tab strip (§4.C.1), one new string beside `myReportsCopy`. */
export const logsCopy = {
  guildEmpty: (guildName: string): string => `No ${guildName} reports yet.`,
} as const;

/** The companion panel's pairing success line (§4.D.2), replacing the code block in place
 *  once a new device appears. */
export const pairingCopy = {
  success: (deviceName: string): string =>
    `${deviceName} paired. It starts uploading as soon as you are logging.`,
} as const;
