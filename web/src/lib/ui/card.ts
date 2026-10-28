// web/src/lib/ui/card.ts
// The design system's card recipe (--raised, 1 px --line, 6 px radius, 1 px gold top highlight),
// as one class string: there is no global `.card` class, so every card-shaped box shares this.
/** The surface alone: the /premium plan and example cards add their own top highlight, or none. */
export const CARD_SURFACE = 'bg-raised border-line rounded-panel border';
export const CARD = `${CARD_SURFACE} shadow-[inset_0_1px_0_rgba(229,185,85,.35)]`;
