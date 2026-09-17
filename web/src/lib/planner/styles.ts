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

/**
 * Same recipe as SECONDARY_BUTTON, but 44px tall on every breakpoint instead of shrinking to
 * 36px on desktop. For controls outside the planner -- the homepage subscribe button -- where
 * the 44px hit target is a fixed requirement, not just a phone-only minimum.
 */
export const SECONDARY_BUTTON_FIXED =
  'rounded-control inline-flex h-11 items-center border text-[12px] font-bold tracking-[0.06em] uppercase';

/**
 * The four states a talent cell is read at a glance by, from design/DESIGN-SYSTEM.md
 * and the game's own frame: gold says "you can spend here", green says "you have",
 * and a dimmed grey says "not yet".
 */
type CellState = 'locked' | 'available' | 'filled' | 'maxed';

export function cellState(rank: number, maxRank: number, available: boolean): CellState {
  if (rank >= maxRank) return 'maxed';
  if (rank > 0) return 'filled';
  return available ? 'available' : 'locked';
}

/** The cell's own border. `locked` dims the icon with it, which is one state, not two. */
export const CELL_BORDER: Record<CellState, string> = {
  maxed: 'border-kill',
  filled: 'border-kill',
  available: 'border-gold',
  locked: 'border-line-soft opacity-50',
};

/**
 * The rank pill in the cell's corner. Filled and maxed share the green border and
 * differ by the number they show, which is what tells them apart in game too.
 */
export const CELL_PILL: Record<CellState, string> = {
  maxed: 'border-kill text-kill',
  filled: 'border-kill text-text',
  available: 'border-gold text-gold',
  locked: 'border-line text-muted',
};

/**
 * A prerequisite link's stroke, keyed by whether it's met: gold once the prerequisite
 * holds the rank its dependent needs, the same grey `border-line` uses otherwise. A
 * `stroke-` utility rather than `border-`, so it lives beside CELL_BORDER/CELL_PILL
 * instead of inside either of them.
 */
export const CONNECTOR_STROKE: Record<'met' | 'unmet', string> = {
  met: 'stroke-gold',
  unmet: 'stroke-line',
};
