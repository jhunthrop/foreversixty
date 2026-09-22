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
