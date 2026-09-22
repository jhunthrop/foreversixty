// web/src/lib/account/layout.ts
// Reserved-height constants for /account's loading skeletons and its signed-out/signed-in
// first-screen match (spec 2026-09-22 §2.5, §2.6). Values below are measured against a
// real `FOREVER_DATA=fixture` build at 360px (Task 7 of
// docs/superpowers/plans/2026-09-22-states-account.md), with route-stubbed /v1/me
// responses so the measurement isn't confounded by network conditions -- the loading ->
// signed-in-ready transition measures 0.0066 CLS and the loading -> signed-out transition
// measures 0 CLS with these constants (Lighthouse's own CLS budget, web/lighthouserc.json,
// is 0.05 -- both are well inside it). `npm run lhci`'s own account.html run cannot
// confirm this directly: its sandbox has no CORS allowance for localhost to call the real
// API, so it always lands on the unrelated `failed`/LoadError branch instead -- see
// SIGNED_OUT_MIN_H's own comment and Task 7's report for that finding.

/** Characters section: heading + intro line + 3 placeholder rows at 360px.
 *  measured against meBnetFixture at 360px, 2026-09-22 */
export const CHARACTERS_SKELETON_MIN_H = 'min-h-[352px]';

/** Devices + You panel: pairing block, device row, and the identity/sign-out block.
 *  measured against meBnetFixture at 360px, 2026-09-22 */
export const IDENTITY_SKELETON_MIN_H = 'min-h-[428px]';

/** Guilds + Billing + Reports panel: the tallest panel once a guild row is present.
 *  measured against meBnetFixture at 360px, 2026-09-22 */
export const MORE_SKELETON_MIN_H = 'min-h-[144px]';

/** The signed-out prompt's floor. `/account`'s loading branch always renders all three
 *  skeletons above (Account.svelte does not know `signedIn` until `/v1/me` answers), so a
 *  genuinely signed-out visitor only avoids a layout shift if this block reserves the SAME
 *  total height the loading skeletons already reserved -- not just the shorter
 *  "Characters only" first screen a first-pass reading of spec 2026-09-22 §2.5 suggested.
 *  Same fixed-height-regardless-of-branch pattern CurrentCharacterChip.svelte's CHIP_HEIGHT
 *  and the pairing-mode block above already use in this file, for the same reason.
 *  988px = CHARACTERS_SKELETON_MIN_H (352) + IDENTITY_SKELETON_MIN_H (428) +
 *  MORE_SKELETON_MIN_H (144) + two gap-8 gaps (2 x 32) between them, matching the loading
 *  branch's `flex flex-col gap-8` wrapper exactly. This value is unverifiable against
 *  `npm run lhci`'s own account.html run: that sandbox has no CORS allowance for
 *  localhost to call the real API, so `/v1/me` fails there and the page always lands on
 *  the `failed`/`LoadError` branch (44px), never this one -- confirmed by instrumenting a
 *  real headless run directly (see docs/superpowers/plans/2026-09-22-states-account.md
 *  Task 7's report). The 552px-vs-988px choice made no difference to the measured
 *  account.html CLS (0.3703125 either way) for exactly that reason; 988px is kept because
 *  it is the mathematically correct value for a real signed-out visitor, verified by
 *  direct DOM/CLS instrumentation with the sandbox's CORS block worked around. */
export const SIGNED_OUT_MIN_H = 'min-h-[988px]';
