// web/src/lib/ui/talent-editor.ts
// The TalentEditor primitive's data and pure rules (planner spec 2026-11-08 section 3A, 4.4): the
// trees a caller hands ui/TalentEditor already resolved to cells (rank, state, tooltip lines), so
// the primitive draws and routes gestures and never reads a planner store. The grid helpers here
// are the editor's keyboard model: trees are sparse, so focus moves over the occupied cells.

/** A cell's colour (decision P2, the client's TalentFrame): can spend, partial, maxed, locked. */
export type TalentCellState = 'avail' | 'part' | 'max' | 'lock';

export interface EditorCell {
  id: number;
  tier: number;
  column: number;
  name: string;
  iconSrc: string;
  rank: number;
  max: number;
  state: TalentCellState;
  /** The cell's rank differs from the loaded or compared band: the gold dot. */
  differs?: boolean;
  /** Spent past the character's level (ruling P4): the dashed, dimmed rank pill. */
  planned?: boolean;
  /** Appended to the accessible name when planned, e.g. "planned past level 34". */
  plannedLabel?: string;
  /** This rank's text ('' at rank 0). */
  now: string;
  /** The next rank's text ('' at max). */
  next: string;
  /** The tooltip's requirement line in --wipe (a refusal), '' for none. */
  req?: string;
  /** The tooltip's gold band line, '' for none. */
  band?: string;
  /** Drawn locked (no points left at the character's level, decision P3) but a press still plans
   *  a point ahead (P4): not aria-disabled. */
  plannable?: boolean;
}

export interface EditorConnector {
  d: string;
  met: boolean;
}

export interface EditorTree {
  id: number;
  name: string;
  iconSrc: string;
  points: number;
  artSrc?: string;
  columns: number;
  tiers: number;
  viewBox: string;
  connectors: readonly EditorConnector[];
  cells: readonly EditorCell[];
}

export interface TalentEditorCopy {
  /** The tooltip's rank line: `Rank 2 of 5`. */
  rankOf: (rank: number, max: number) => string;
  /** The label over the next rank's text: `Next rank`. */
  nextRank: string;
  /** The per-tree reset's visible word: `Reset`. */
  resetTree: string;
  /** The per-tree reset's accessible name: `Reset Beast Mastery`. */
  resetTreeLabel: (treeName: string) => string;
}

/** The delay before a talent's tooltip opens on hover or focus (spec 5.4). */
export const TIP_DELAY_MS = 150;
/** A primary-button press this long removes a point (spec 4.4). */
export const LONG_PRESS_MS = 500;

/** The cell's accessible name: `Pathfinding, rank 0 of 2`, plus the planned mark when planned. */
export function cellLabel(cell: EditorCell): string {
  const base = `${cell.name}, rank ${cell.rank} of ${cell.max}`;
  return cell.planned && cell.plannedLabel ? `${base}, ${cell.plannedLabel}` : base;
}

/** Occupied cells, top to bottom then left to right: the keyboard's reading order. */
export function sortedCells(cells: readonly EditorCell[]): EditorCell[] {
  return [...cells].sort((a, b) => a.tier - b.tier || a.column - b.column);
}

/** The tree as tiers × columns, `undefined` where no talent sits (an empty gridcell). */
export function gridRows(tree: EditorTree): (EditorCell | undefined)[][] {
  const byPlace = new Map(tree.cells.map((cell) => [`${cell.tier}:${cell.column}`, cell]));
  return Array.from({ length: tree.tiers }, (_, tier) =>
    Array.from({ length: tree.columns }, (_, column) => byPlace.get(`${tier}:${column}`)),
  );
}

/**
 * The cell id arrow keys move to from `currentId`. Horizontal moves walk the reading order (the
 * end of a tier continues into the next); vertical moves jump to the nearest column in the next
 * occupied tier. Both stop at the edges. An unknown id returns the first cell's.
 */
export function moveCell(cells: readonly EditorCell[], currentId: number, dx: number, dy: number): number {
  const order = sortedCells(cells);
  const current = order.find((cell) => cell.id === currentId);
  if (!current) return order[0]?.id ?? currentId;
  if (dx !== 0) {
    const at = order.indexOf(current);
    return order[Math.min(Math.max(at + dx, 0), order.length - 1)].id;
  }
  if (dy === 0) return currentId;
  const tiers = [...new Set(order.map((cell) => cell.tier))];
  const candidates =
    dy > 0 ? tiers.filter((t) => t > current.tier) : tiers.filter((t) => t < current.tier).reverse();
  const tier = candidates[0];
  if (tier === undefined) return currentId;
  const row = order.filter((cell) => cell.tier === tier);
  return row.reduce((best, cell) =>
    Math.abs(cell.column - current.column) < Math.abs(best.column - current.column) ? cell : best,
  ).id;
}

/** What a key does on a focused cell: move, remove a point, or nothing (Enter is the button's click). */
export type GridKeyAction = { kind: 'move'; dx: number; dy: number } | { kind: 'remove' } | null;

const ARROWS: Readonly<Record<string, readonly [number, number]>> = {
  ArrowRight: [1, 0],
  ArrowLeft: [-1, 0],
  ArrowDown: [0, 1],
  ArrowUp: [0, -1],
};

export function gridKeyAction(key: string): GridKeyAction {
  const delta = ARROWS[key];
  if (delta) return { kind: 'move', dx: delta[0], dy: delta[1] };
  return key === 'Backspace' || key === 'Delete' ? { kind: 'remove' } : null;
}

/** The tree a phone opens on: the one with the most points, the first when none or tied. */
export function openTreeIndex(trees: readonly EditorTree[]): number {
  return trees.reduce((best, tree, index) => (tree.points > trees[best].points ? index : best), 0);
}

/** The cell that holds a tree's one tab stop: the last focused when it still exists, else the first. */
export function tabStopId(tree: EditorTree, lastFocused: number | undefined): number | undefined {
  const order = sortedCells(tree.cells);
  return order.find((cell) => cell.id === lastFocused)?.id ?? order[0]?.id;
}

/** True when a loaded icon is the data pipeline's 1×1 placeholder: show the dark socket instead. */
export function isPlaceholderIcon(image: EventTarget): boolean {
  return image instanceof HTMLImageElement && image.naturalWidth <= 1;
}
