// web/src/lib/ui/talent-tree.ts
// The one talent tree's data and pure rules (talent-tree lane 2026-11-10; planner spec 2026-11-08
// section 3A, 4.4). Every tree on the site, the planner's editable trees and every read-only
// copy (guides, sim, /sim/talents), is ui/TalentTrees over these views: the caller resolves its
// talents to cells (rank, state, tooltip text) and the component draws and routes gestures, never
// reading a planner store. The grid helpers are the keyboard model: trees are sparse, so focus
// moves over the occupied cells.
import { talentTipKey, talentTipModel, type GameTipModel } from '../tooltips/game-tip';

/** A cell's colour (decision P2, the client's TalentFrame): can spend, partial, maxed, locked. */
export type TalentCellState = 'avail' | 'part' | 'max' | 'lock';

/**
 * What a tree is for. `edit`: the planner (points, rules, the per-tree reset, a foot). `view`: a
 * read-only build at a readable size (the sim, /sim/talents). `compact`: a read-only build sized
 * to its box, three trees side by side (a guide's At a glance and Leveling lane). The look, the
 * tooltip and the keyboard are the same in all three; only `edit` adds and removes points.
 */
export type TalentTreesMode = 'edit' | 'view' | 'compact';

/** The cell size scale (tokens in TalentTrees.svelte): `lg` 44 / 48 from 768 / 56 from 1760;
 *  `md` 44 below 768, 36 from 768; `sm` fits three trees to the box, 18 to 32. */
export type TalentCellSize = 'sm' | 'md' | 'lg';

export const MODE_CELL_SIZE: Readonly<Record<TalentTreesMode, TalentCellSize>> = {
  edit: 'lg',
  view: 'md',
  compact: 'sm',
};

export interface TalentCellView {
  id: number;
  tier: number;
  column: number;
  name: string;
  /** The client icon's stem: the tooltip's icon. */
  icon: string;
  iconSrc: string;
  rank: number;
  max: number;
  state: TalentCellState;
  /** Every rank's text, rank 1 first. Absent on a page whose tooltip host serves the text (a
   *  guide's bundle): the cell then names its tooltip by key only. */
  ranks?: readonly string[];
  /** The cell's rank differs from the loaded or compared band: the gold dot. */
  differs?: boolean;
  /** Spent past the character's level (ruling P4): the dashed, dimmed rank pill. */
  planned?: boolean;
  /** Appended to the accessible name when planned, e.g. "planned past level 34". */
  plannedLabel?: string;
  /** The tooltip's requirement line in --wipe (a refusal), '' for none. */
  req?: string;
  /** The tooltip's gold band line, '' for none. */
  band?: string;
  /** Drawn locked (no points left at the character's level, decision P3) but a press still plans
   *  a point ahead (P4): not aria-disabled. */
  plannable?: boolean;
}

export interface TalentConnectorView {
  d: string;
  met: boolean;
  /** The prerequisite's talent id and the rank the dependent needs of it: a read-only tree that
   *  changes rank (the guide's level slider) recomputes `met` from these. */
  from: number;
  needs: number;
}

export interface TalentTreeView {
  id: number;
  name: string;
  /** The spec icon in the tree head; '' when the build ships none. */
  iconSrc: string;
  points: number;
  /** The client's tree art behind the grid. */
  artSrc?: string;
  columns: number;
  tiers: number;
  viewBox: string;
  connectors: readonly TalentConnectorView[];
  cells: readonly TalentCellView[];
}

/** The editable trees' words (the planner's copy). */
export interface TalentTreesCopy {
  /** The per-tree reset's visible word: `Reset`. */
  resetTree: string;
  /** The per-tree reset's accessible name: `Reset Beast Mastery`. */
  resetTreeLabel: (treeName: string) => string;
}

/** The delay before an editable talent's tooltip opens on hover or focus (spec 5.4). */
export const TIP_DELAY_MS = 150;
/** A primary-button press this long removes a point (spec 4.4). */
export const LONG_PRESS_MS = 500;

/** A cell's colour: maxed gold, partial green, a point you can spend green, else locked grey. */
export function cellState(rank: number, max: number, canSpend: boolean): TalentCellState {
  if (rank >= max) return 'max';
  if (rank > 0) return 'part';
  return canSpend ? 'avail' : 'lock';
}

/** A read-only cell's colour: the client's frame for a build you are reading (nothing to spend). */
export function viewState(rank: number, max: number): TalentCellState {
  return cellState(rank, max, false);
}

/** The cell's accessible name: `Pathfinding, rank 0 of 2`, plus the planned mark when planned. */
export function cellLabel(cell: TalentCellView): string {
  const base = `${cell.name}, rank ${cell.rank} of ${cell.max}`;
  return cell.planned && cell.plannedLabel ? `${base}, ${cell.plannedLabel}` : base;
}

/** The key a page's tooltip host opens for a cell (`talent:<id>:<rank>`). */
export function cellTipKey(cell: TalentCellView): string {
  return talentTipKey(cell.id, cell.rank);
}

/** The cell's tooltip, the site's one game tooltip (GameTooltip): name and `Rank n/m`, the
 *  requirement when the point is refused, this rank's text, the next rank's, the band line. Null
 *  when the cell carries no text (its page's host serves it by key). */
export function cellTipModel(cell: TalentCellView): GameTipModel | null {
  if (cell.ranks === undefined) return null;
  const model = talentTipModel(
    { name: cell.name, icon: cell.icon, maxRank: cell.max, ranks: [...cell.ranks] },
    cell.rank,
  );
  return {
    ...model,
    ...(cell.req ? { requirement: cell.req } : {}),
    ...(cell.band ? { note: cell.band } : {}),
  };
}

/** Occupied cells, top to bottom then left to right: the keyboard's reading order. */
export function sortedCells(cells: readonly TalentCellView[]): TalentCellView[] {
  return [...cells].sort((a, b) => a.tier - b.tier || a.column - b.column);
}

/** The tree as tiers × columns, `undefined` where no talent sits (an empty gridcell). */
export function gridRows(tree: TalentTreeView): (TalentCellView | undefined)[][] {
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
export function moveCell(
  cells: readonly TalentCellView[],
  currentId: number,
  dx: number,
  dy: number,
): number {
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
export function openTreeIndex(trees: readonly TalentTreeView[]): number {
  return trees.reduce((best, tree, index) => (tree.points > trees[best].points ? index : best), 0);
}

/** The cell that holds a tree's one tab stop: the last focused when it still exists, else the first. */
export function tabStopId(tree: TalentTreeView, lastFocused: number | undefined): number | undefined {
  const order = sortedCells(tree.cells);
  return order.find((cell) => cell.id === lastFocused)?.id ?? order[0]?.id;
}

/** True when a loaded icon is the data pipeline's 1×1 placeholder: show the dark socket instead. */
export function isPlaceholderIcon(image: EventTarget): boolean {
  return image instanceof HTMLImageElement && image.naturalWidth <= 1;
}

/**
 * Read-only trees at other ranks over the same layout (the guide's level slider): each cell's
 * rank from `rankOf` (capped at its max), its state, each tree's points and each arrow's `met`.
 */
export function withRanks(
  trees: readonly TalentTreeView[],
  rankOf: (cell: TalentCellView, treeIndex: number, position: number) => number,
): TalentTreeView[] {
  return trees.map((tree, treeIndex) => {
    const positions = new Map(sortedCells(tree.cells).map((cell, position) => [cell.id, position]));
    const cells = tree.cells.map((cell) => {
      const rank = Math.min(cell.max, Math.max(0, rankOf(cell, treeIndex, positions.get(cell.id) ?? 0)));
      return { ...cell, rank, state: viewState(rank, cell.max) };
    });
    const ranks = new Map(cells.map((cell) => [cell.id, cell.rank]));
    return {
      ...tree,
      cells,
      points: cells.reduce((sum, cell) => sum + cell.rank, 0),
      connectors: tree.connectors.map((c) => ({ ...c, met: (ranks.get(c.from) ?? 0) >= c.needs })),
    };
  });
}

/** The trees at a ladder step: its talent string (`a-b-c`, one digit per talent in each tree's
 *  tier-then-column order, `ladderTalentString`) over the trees' layout. */
export function treesAtTalents(trees: readonly TalentTreeView[], talents: string): TalentTreeView[] {
  const digits = talents.split('-');
  return withRanks(trees, (_cell, treeIndex, position) => Number(digits[treeIndex]?.[position] ?? 0) || 0);
}
