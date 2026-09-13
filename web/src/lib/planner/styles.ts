// web/src/lib/planner/styles.ts
// Class strings used by more than one planner component, so a design-system control is spelled
// out in exactly one place.

/**
 * design/DESIGN-SYSTEM.md "Secondary button": warm border, uppercase 12px 700 at 0.06em tracking,
 * 36px tall -- 44px on phone, where it has to clear the hit-target minimum. There is no primary
 * button. Callers add the border and text colour they want, and their own horizontal padding.
 */
export const SECONDARY_BUTTON =
  'rounded-control inline-flex h-11 items-center border text-[12px] font-bold tracking-[0.06em] uppercase md:h-9';
