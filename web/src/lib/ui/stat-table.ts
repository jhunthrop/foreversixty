// web/src/lib/ui/stat-table.ts
// StatTable's props and its one sort rule (sim spec 2026-11-06 §4.7, ruling S8): the figure
// columns a table draws, the sorted column, and what each head says to assistive tech.

/** One figure column: its head, and whether pressing the head sorts by it. */
export interface StatColumn {
  id: string;
  label: string;
  sortable?: boolean;
  /** A game-tooltip key the head opens (the Crit head's note). */
  tip?: string;
}

export type SortDirection = 'desc' | 'asc';

export interface StatSort {
  column: string;
  direction: SortDirection;
}

/** One figure cell of a row: the primary (first) cell is the sorted, strong one. */
export interface StatFigure {
  value: string;
  /** A dash or a missing figure, drawn muted. */
  muted?: boolean;
  /** A signed difference: a gain in the kill colour, a loss in the wipe colour. */
  tone?: 'kill' | 'wipe' | 'gold';
  /** A colour the figure takes from data (a parse percentile), over the tone. */
  color?: string;
  /** The cell's spoken name when the figure alone says too little (a parse: "98th percentile"). */
  label?: string;
  /** The cell's testid, for a test that reads one figure. */
  testid?: string;
}

const ARIA_SORT: Readonly<Record<SortDirection, 'descending' | 'ascending'>> = {
  desc: 'descending',
  asc: 'ascending',
};

/** What a sortable head's columnheader says: its direction when sorted, else `none`. A head that
 *  does not sort carries no aria-sort at all. */
export function ariaSortOf(
  column: StatColumn,
  sort: StatSort | undefined,
): 'descending' | 'ascending' | 'none' | undefined {
  if (column.sortable !== true) return undefined;
  return sort?.column === column.id ? ARIA_SORT[sort.direction] : 'none';
}

/** The sorted head's mark after its label. */
export function sortMark(direction: SortDirection): string {
  return direction === 'desc' ? '▾' : '▴';
}
