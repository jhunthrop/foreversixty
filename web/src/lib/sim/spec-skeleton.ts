// web/src/lib/sim/spec-skeleton.ts
// The specs grid's reserved shape, shared by the static shell (sim/specs.astro, which
// renders it before the island ever mounts) and SpecGrid.svelte (which renders the same
// markup for its own loading and error states) so all three moments -- first paint,
// loading, and the CORS-blocked error lhci's static origin always measures -- reserve the
// same height. One source for the card count and classes, so a spec added to or removed
// from the dps list can never leave the shell's own reservation stale (round 5,
// final whole-branch review: sim/specs.astro's shell reserved no height for the grid at
// all -- #sim measured 42px before hydration and 3684px after, a footer shift the
// layout-shifts audit could not attribute to a font because the real cause was never a
// font; it was this).
import { dpsSpecs } from './spec-label';

/** The specs grid's own responsive column classes. */
export const SPEC_GRID_CLASSES = 'grid grid-cols-1 gap-3 px-[18px] md:grid-cols-2 md:px-0 lg:grid-cols-3';

/** One placeholder card's classes -- the height every real SpecCard reserves before content. */
export const SPEC_CARD_SKELETON_CLASSES = 'bg-card-top border-line rounded-panel min-h-[168px] border';

/** One slot index per dps spec, so the skeleton always has the same card count the real grid will. */
export function specSkeletonSlots(): number[] {
  return Array.from({ length: dpsSpecs().length }, (_, index) => index);
}
