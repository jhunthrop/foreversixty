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
  pointerLink: 'Setup',
} as const;
