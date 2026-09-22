// web/src/lib/account/layout.ts
// Reserved-height constants for /account's loading skeletons and its signed-out/signed-in
// first-screen match (brief 2026-09-22 §B2/§B6/§B8). Re-measured against a real
// `FOREVER_DATA=fixture` build at 360px after the page-header/two-column/rail-panel
// redesign (account-visual lane, 2026-09-22), with route-stubbed /v1/me responses so the
// measurement isn't confounded by network conditions. The title itself (`<h1>`) now
// renders identically in every status -- loading, failed, signed-out, ready -- so it is
// common to all four and excluded from the figures below; what each constant reserves is
// everything from just under the title to the bottom of its section. As with the previous
// lane's measurement, `npm run lhci`'s own account.html run cannot confirm this directly:
// its sandbox has no CORS allowance for localhost to call the real API, so it always lands
// on the `failed`/LoadError branch instead -- see SIGNED_OUT_MIN_H's own comment.

/** The identity line (BattleTag/email, sign-in method, Sign out), the compact
 *  current-character chip, and the Characters panel -- everything the ready view shows
 *  before "Your reports". measured against meBnetFixture at 360px, 2026-09-22 */
export const CHARACTERS_SKELETON_MIN_H = 'min-h-[692px]';

/** "Your reports" and the rail's Devices panel -- the two sections that follow Characters
 *  in reading order at 360px (the two-column grid only applies at `lg`; below that every
 *  panel stacks in one column). measured against meBnetFixture at 360px, 2026-09-22 */
export const IDENTITY_SKELETON_MIN_H = 'min-h-[332px]';

/** The rail's You (pseudonym toggle) and Guilds-and-plan panels.
 *  measured against meBnetFixture at 360px, 2026-09-22 */
export const MORE_SKELETON_MIN_H = 'min-h-[308px]';

/** The signed-out prompt's floor. `/account`'s loading branch always renders all three
 *  skeletons above (Account.svelte does not know `signedIn` until `/v1/me` answers), so a
 *  genuinely signed-out visitor only avoids a layout shift if this block reserves the SAME
 *  total height the loading skeletons already reserved. Same fixed-height-regardless-of-
 *  branch pattern CurrentCharacterChip.svelte's CHIP_HEIGHT and the pairing-mode block
 *  above already use in this file, for the same reason.
 *  1396px = CHARACTERS_SKELETON_MIN_H (692) + IDENTITY_SKELETON_MIN_H (332) +
 *  MORE_SKELETON_MIN_H (308) + two gap-8 gaps (2 x 32) between them, matching the loading
 *  branch's `flex flex-col gap-8` wrapper exactly. Unverifiable against `npm run lhci`'s
 *  own account.html run for the CORS reason above; the previous lane confirmed via direct
 *  DOM/CLS instrumentation that the exact figure does not change lhci's measured CLS
 *  (the sandbox never reaches this branch at all), so 1396px is kept because it is the
 *  mathematically correct value for a real signed-out visitor. */
export const SIGNED_OUT_MIN_H = 'min-h-[1396px]';
