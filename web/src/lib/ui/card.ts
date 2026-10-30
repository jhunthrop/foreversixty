// web/src/lib/ui/card.ts
// The design system's card surface as class strings, for boxes that are not yet a `Card`
// component (site alignment spec 2026-10-29 S1b, polish spec 4.2). The utilities live in
// global.css: the warm gradient, the gold-tinted border, the 12 px panel radius and the deep shadow.
/** A primary card: the surface with the inset gold top edge (the /account card). */
export const CARD = 'card-surface';
/** A secondary card (prose asides, tables inside prose, tool-canvas frames): no top edge. */
export const CARD_QUIET = 'card-surface-quiet';
