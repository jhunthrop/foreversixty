// web/src/lib/guild/raid-report-styles.ts
// Recipes the Report, Your night and Attendance share (attendance spec 3.3 to 3.5, 4): the own-row
// wash, the link row's hover and focus ring, the 8 px bar track, and the `Show all` text button.
import { RAID_TEXT_BUTTON } from './raid-buttons';

/** The viewer's own row (G1's roster wash). */
export const OWN_ROW_WASH = 'rgba(229,185,85,.06)';

/** A row that links: `--card-top` on hover, a 2 px gold ring inside it on focus. */
export const ROW_LINK =
  'hover:bg-card-top active:bg-[#171d2b] focus-visible:outline-gold focus-visible:outline-2 focus-visible:-outline-offset-2';

/** `Show all {n} raiders`: the text button, `--gold-hover` on hover. */
export const SHOW_ALL_BUTTON = `${RAID_TEXT_BUTTON} hover:text-gold-hover self-start`;

/** The 8 px bar's track and its muted fill (never `--wipe`, ruling R-A3). */
export const BAR_TRACK = 'bg-line-soft relative inline-block h-2 overflow-hidden rounded-full';
export const BAR_FILL = 'bg-muted absolute inset-y-0 left-0';

/** The row divider every block uses. */
export const ROW_DIVIDER = 'border-line-soft border-b last:border-b-0';
