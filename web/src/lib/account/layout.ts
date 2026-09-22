// web/src/lib/account/layout.ts
// Reserved-height constants for /account's loading skeletons and its signed-out/signed-in
// first-screen match (spec 2026-09-22 §2.5, §2.6). First-pass values built from the same
// atomic units current-character-layout.ts uses (44px rows, 12px/32px gaps); Task 7 of
// docs/superpowers/plans/2026-09-22-states-account.md measures the built page and corrects
// them so the skeleton's height matches its ready view exactly (Lighthouse CLS budget,
// web/lighthouserc.json, is 0.05).

/** Characters section: heading + intro line + 3 placeholder rows at 360px. */
export const CHARACTERS_SKELETON_MIN_H = 'min-h-[360px]';

/** Devices + You panel: pairing block, device row, and the identity/sign-out block. */
export const IDENTITY_SKELETON_MIN_H = 'min-h-[280px]';

/** Guilds + Billing + Reports panel: the tallest panel once a guild row is present. */
export const MORE_SKELETON_MIN_H = 'min-h-[220px]';

/** The signed-out prompt's floor, matched to the signed-in first screen (title + toast
 *  slot + Characters) so the footer does not move when the session resolves either way. */
export const SIGNED_OUT_MIN_H = 'min-h-[420px]';
