// web/src/lib/rankings/layout.ts
// The height Rankings.svelte's main board reserves while its rows are in flight. A module
// rather than an inline literal for the same reason guild/layout.ts is one: every sizing
// constant this lane added is readable in one place next to the reasoning behind it.
//
// Eight rows at h-11 with the Skeleton's own gap-3 between them come to ~436px, so this
// reserves the board the skeleton itself draws. The encounter picker's skeleton keeps no
// minHeight: it sits inside a panel the page has already sized, and it is three short
// lines rather than a board.
export const RANKINGS_LOADING_MIN_H = 'min-h-[440px]';

/**
 * The character board's column template, shared by the header line and every row so the
 * two can never drift apart: rank, name, guild, then seven columns at md and up (size, the
 * metric, executed, date, length, build, report); on a
 * phone three columns (rank, name, the build and report links) with the figures folded
 * into the row's own second line.
 */
export const RANKING_ROW_GRID =
  'grid-cols-[40px_minmax(0,1fr)_auto] md:grid-cols-[40px_minmax(140px,1.4fr)_minmax(120px,1fr)_64px_88px_72px_96px_64px_72px_64px]';
