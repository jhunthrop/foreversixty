// web/src/lib/planner/styles.ts
// Class strings used by more than one planner component, so a design-system control is spelled
// out in exactly one place.

/**
 * The secondary button (docs/design-system.md: 15 px 600, sentence case, the control radius),
 * 36px tall -- 44px on phone, where it has to clear the hit-target minimum. Callers add the
 * border and text colour they want, and their own horizontal padding. The one action a page
 * exists for takes PRIMARY_BUTTON instead.
 *
 * Logs build review, fix round 1 (2026-10-04): a Secondary button carried no hover state at
 * all (caught on `LogsHero.svelte`'s "Open this report", which is this token with no hover
 * override of its own) -- design/DESIGN-SYSTEM.md's own "120ms hover transitions on cards and
 * links" budget applies here too, so it is added to the shared token rather than to one
 * caller: a hover-variant class always outranks a caller's own plain `border-*`/`text-*`
 * class in CSS, regardless of which comes first in the concatenated string, so every existing
 * caller's own colours are untouched outside `:hover`.
 */
const SECONDARY_BUTTON_HOVER =
  'transition-colors duration-[120ms] hover:border-gold-hover hover:text-gold-hover';

/**
 * Launch pricing spec 4, Disabled: a disabled or aria-disabled _FIXED button reads as off at
 * 50 % with a not-allowed cursor (global.css `.control-fixed`). A busy one keeps its own look.
 */
const FIXED_BUTTON_DISABLED = 'control-fixed';

export const SECONDARY_BUTTON = `rounded-field inline-flex h-11 items-center border button-label md:h-9 ${SECONDARY_BUTTON_HOVER}`;

/**
 * Same recipe as SECONDARY_BUTTON, but 44px tall on every breakpoint instead of shrinking to
 * 36px on desktop. For controls outside the planner -- the homepage subscribe button -- where
 * the 44px hit target is a fixed requirement, not just a phone-only minimum.
 */
export const SECONDARY_BUTTON_FIXED = `rounded-field inline-flex h-11 items-center border button-label ${SECONDARY_BUTTON_HOVER} ${FIXED_BUTTON_DISABLED}`;

/**
 * The primary button: the same recipe in the round-5 teal, for the single action a page is for
 * (Sign in with Battle.net on /login and /setup, Run sim on /sim, Sim this build in the
 * planner). One per view, never two side by side: two filled buttons are two secondaries.
 * Callers add their own horizontal padding.
 */
export const PRIMARY_BUTTON =
  'rounded-field border-teal-edge bg-teal text-strong hover:bg-teal-edge inline-flex h-11 items-center border button-label md:h-9';

/** PRIMARY_BUTTON at 44px on every breakpoint, as SECONDARY_BUTTON_FIXED is to SECONDARY_BUTTON. */
export const PRIMARY_BUTTON_FIXED = `rounded-field border-teal-edge bg-teal text-strong hover:bg-teal-edge inline-flex h-11 items-center border button-label ${FIXED_BUTTON_DISABLED}`;
