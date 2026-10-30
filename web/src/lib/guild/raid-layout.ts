// web/src/lib/guild/raid-layout.ts
// The raid surfaces' loading reserves (raids spec 4 "Page states", "Raids tab loading"), each equal
// to the ready height measured on the fixture build (design/reviews/2026-10-26-guild-raids, the
// pre-review and review rounds), so the swap from reserve to content moves nothing. Tiny on purpose:
// GuildShell's entry chunk imports it.

/** The raid head, ready (no band padding since site alignment D1: 38 px less on a phone, 60 from md): back link, eyebrow, title, when, count and tabs, for a member. */
export const RAID_BAND_MIN = 'min-h-[309px] md:min-h-[251px]';

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

/** A done raid's band lines (attendance spec 4 "Page states"): the h1, the when line, the night
 *  line and the match line in place of the count line, measured ready on the fixture. A done raid
 *  has no officer row; an officer's match line holds the 44 px `Not this log`. */
const RAID_DONE_BAND_MEMBER = 'min-h-[223px] md:min-h-[157px]';
const RAID_DONE_BAND_OFFICER = 'min-h-[267px] md:min-h-[157px]';

/** The loading band's lines reserve for the page the viewer last saw. */
export function bandLinesReserve(done: boolean, officer: boolean): string {
  if (done) return officer ? RAID_DONE_BAND_OFFICER : RAID_DONE_BAND_MEMBER;
  return officer ? `${RAID_BAND_LINES_MIN} ${RAID_BAND_OFFICER}` : RAID_BAND_LINES_MIN;
}

/** `Your night` for a member: the panel and its head, then one row per own row it last drew,
 *  measured on the fixture's one-row panel (228 / 136 px). */
export const YOUR_NIGHT_RESERVE =
  'h-[calc(56px+var(--strip-rows)*172px)] md:h-[calc(68px+var(--strip-rows)*68px)]';

/** The Report, ready (spec 4 "Page states"): one column below 1024, two from it. */
export const REPORT_RESERVE = 'min-h-[900px] lg:min-h-[520px]';

/** Attendance for a cached officer (spec 4 "Page states"). */
export const ATTENDANCE_RESERVE = 'min-h-[420px]';

/** A guild view with no full-bleed band at its top (the landing, claim, join, settings, the address
 *  that does not parse, and their chunk placeholders) keeps 28 px under the header from `md`: `main`
 *  itself has none there, so a band sits flush against the header (UX N1). */
export const GUILD_BANDLESS_TOP = 'md:pt-7';

/** A guild view is at least one screen tall from its first paint, so the footer starts below the
 *  fold and never jumps when the view under it grows (phone content starts at 72 px, desktop at 81). */
export const GUILD_VIEW_MIN = 'min-h-[calc(100svh-72px)] md:min-h-[calc(100svh-81px)]';
