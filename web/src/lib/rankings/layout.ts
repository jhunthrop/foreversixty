// web/src/lib/rankings/layout.ts
// The sizes Rankings.svelte reserves and shares between its column line and its rows, in one
// place next to the reasoning behind them.
//
// The height the main board reserves while its rows are in flight: eight rows at h-11 with the
// Skeleton's own gap-3 between them come to ~436px, so this reserves the board the skeleton
// itself draws. The encounter picker's skeleton keeps no minHeight: it sits inside a panel the
// page has already sized, and it is three short lines rather than a board.
export const RANKINGS_LOADING_MIN_H = 'min-h-[440px]';

/** From here up a row has every column; below it Size and Length drop and the guild moves under the name. */
export const WIDE_ROW_PX = 1100;
/** From here up a row is a table row; below it, a name over a muted line with the figure and Report. */
export const TABLE_ROW_PX = 600;

/**
 * The character board's column template, shared by the column line and every row so the two can
 * never drift apart (design/specs/2026-10-11-rankings.md section 3, decision 6):
 *   wide    # · Character · Guild · Size · {Metric} · Executed · Length · Date · Build · Report
 *   middle  # · Character (guild under the name) · {Metric} · Executed · Date · Build · Report
 *   phone   # · Character (a muted line of size, length, executed and build) · {Metric} · Report
 * The cells a width drops are `display: none`, so the columns that remain fill the template in
 * order.
 */
export const RANKING_ROW_GRID =
  'grid-cols-[32px_minmax(0,1fr)_auto_52px] min-[600px]:grid-cols-[44px_minmax(150px,1.5fr)_96px_76px_100px_84px_60px] min-[1100px]:grid-cols-[44px_minmax(150px,1.4fr)_minmax(120px,1fr)_56px_96px_76px_76px_100px_84px_60px]';

/** Cells that exist from the wide row up (Guild, Size, Length). */
export const WIDE_ONLY = 'hidden min-[1100px]:block';
/** Cells that exist from the table row up (Executed, Date, Build). */
export const TABLE_UP = 'hidden min-[600px]:block';
