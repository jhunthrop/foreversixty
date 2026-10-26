// web/src/lib/guild/raid-layout.ts
// The raid surfaces' loading reserves (raids spec 4 "Page states", "Raids tab loading"), each equal
// to the ready height measured on the fixture build (design/reviews/2026-10-26-guild-raids, the
// pre-review and review rounds), so the swap from reserve to content moves nothing. Tiny on purpose:
// GuildShell's entry chunk imports it.

/** The raid band, ready: back link, eyebrow, title, when, count and tabs, for a member. */
export const RAID_BAND_MIN = 'min-h-[347px] md:min-h-[311px]';

/** The band's lines 3–5 while the raid loads (h1, the when line and the count line with the gaps
 *  between them): the back link, eyebrow and tabs around them are drawn for real. */
export const RAID_BAND_LINES_MIN = 'min-h-[163px] md:min-h-[105px]';

/** An officer's band also holds the officer row: on phone two wrapped rows of buttons under the
 *  count (12 + 96), from md the row on the eyebrow line, 22 px taller than the eyebrow alone. */
export const RAID_BAND_OFFICER = 'pb-[108px] md:pb-[22px]';

/** The sign-up strip, ready: panel padding and sub-head, then each row (183 / 101 px for the
 *  first, one px less for each later row, which loses its top divider). */
export const RAID_STRIP_RESERVE =
  'h-[calc(62px+var(--strip-rows)*182px)] md:h-[calc(70px+var(--strip-rows)*100px)]';

/** The Raids tab's planned block, ready: the tab's header row and the `PLANNED n` sub-head (62 px),
 *  then each planned row: 71 desktop (the date block sets it); on phone the mean of the fixture's
 *  rows, 114 with a one-line title and 136.5 with a two-line one (121.5). */
export const PLANNED_RESERVE =
  'h-[calc(62px+var(--planned-rows)*121.5px)] md:h-[calc(62px+var(--planned-rows)*71px)]';

/** A guild view is at least one screen tall from its first paint, so the footer starts below the
 *  fold and never jumps when the view under it grows (phone content starts at 72 px, desktop at 109). */
export const GUILD_VIEW_MIN = 'min-h-[calc(100svh-72px)] md:min-h-[calc(100svh-109px)]';
