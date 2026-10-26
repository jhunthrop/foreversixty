// web/src/lib/guild/raid-layout.ts
// The raid surfaces' loading reserves (raids spec 4 "Page states", "Raids tab loading"), each equal
// to the ready height measured on the fixture build (design/reviews/2026-10-26-guild-raids, the
// pre-review fix round), so the swap from reserve to content moves nothing. Tiny on purpose:
// GuildShell's entry chunk imports it.

/** The raid band, ready: back link, eyebrow, title, when, count and tabs, for a member. */
export const RAID_BAND_MIN = 'min-h-[347px] md:min-h-[311px]';

/** The sign-up strip, ready: panel padding and sub-head, then each row (183 / 101 px for the
 *  first, one px less for each later row, which loses its top divider). */
export const RAID_STRIP_RESERVE =
  'h-[calc(62px+var(--strip-rows)*182px)] md:h-[calc(70px+var(--strip-rows)*100px)]';

/** The Raids tab's planned block, ready: the tab's header row and the `PLANNED n` sub-head (62 px),
 *  then each planned row (97 phone, 71 desktop). */
export const PLANNED_RESERVE =
  'h-[calc(62px+var(--planned-rows)*97px)] md:h-[calc(62px+var(--planned-rows)*71px)]';

/** A guild view is at least one screen tall from its first paint, so the footer starts below the
 *  fold and never jumps when the view under it grows (phone content starts at 72 px, desktop at 109). */
export const GUILD_VIEW_MIN = 'min-h-[calc(100svh-72px)] md:min-h-[calc(100svh-109px)]';
