// web/src/lib/planner/connectors.ts
// Where the game draws a line from a prerequisite to the talent that needs it, as
// a path in the tree grid's own pixel space. Pure and unit-tested: TreeGrid puts
// these in an SVG overlay sized to the grid, so nothing here measures the DOM.
//
// Three shapes, which is what the client draws: straight down a column, straight
// across a row, and an elbow -- down the prerequisite's column to the dependent's
// row, then across into it. Build 1.60.1.69893 uses the first for 67 of its 69
// prerequisites and the second for the other two.
//
// Every line starts and ends at a cell's edge rather than its centre, so no line
// crosses an icon.
import type { TalentTree } from './types';

/** The talent button's side in CSS pixels; TalentCell is h-11 w-11. */
export const CELL_PX = 44;
/** The grid's gap; TreeGrid is gap-2. */
export const GAP_PX = 8;
/** Centre-to-centre distance between neighbouring cells. */
export const PITCH_PX = CELL_PX + GAP_PX;
export const HALF_PX = CELL_PX / 2;

export interface Cell {
  tier: number;
  column: number;
}

export interface Connector {
  /** Stable key for the `{#each}` that draws these. */
  id: string;
  /** The prerequisite's talent id. */
  from: number;
  /** The dependent's talent id. */
  to: number;
  /** An SVG path, in the coordinate space `gridViewBox` describes. */
  d: string;
  /** True once the prerequisite holds the rank the dependent needs. */
  met: boolean;
}

export function cellCentre({ tier, column }: Cell): { x: number; y: number } {
  return { x: column * PITCH_PX + HALF_PX, y: tier * PITCH_PX + HALF_PX };
}

export function connectorPath(from: Cell, to: Cell): string {
  const a = cellCentre(from);
  const b = cellCentre(to);
  if (from.column === to.column) {
    return `M ${a.x} ${a.y + HALF_PX} V ${b.y - HALF_PX}`;
  }
  const side = to.column > from.column ? 1 : -1;
  if (from.tier === to.tier) {
    return `M ${a.x + side * HALF_PX} ${a.y} H ${b.x - side * HALF_PX}`;
  }
  return `M ${a.x} ${a.y + HALF_PX} V ${b.y} H ${b.x - side * HALF_PX}`;
}

/** The viewBox for an overlay covering a grid of this many tiers and columns. */
export function gridViewBox(tiers: number, columns: number): string {
  return `0 0 ${columns * PITCH_PX - GAP_PX} ${tiers * PITCH_PX - GAP_PX}`;
}

export function connectorsFor(tree: TalentTree, ranks: Map<number, number>): Connector[] {
  const byId = new Map(tree.talents.map((talent) => [talent.id, talent]));
  const connectors: Connector[] = [];
  for (const talent of tree.talents) {
    if (talent.prereq_talent_id === null) continue;
    const prereq = byId.get(talent.prereq_talent_id);
    if (prereq === undefined) continue;
    connectors.push({
      id: `${prereq.id}-${talent.id}`,
      from: prereq.id,
      to: talent.id,
      d: connectorPath(prereq, talent),
      met: (ranks.get(prereq.id) ?? 0) >= (talent.prereq_rank ?? 0),
    });
  }
  return connectors.sort((a, b) => a.from - b.from || a.to - b.to);
}
