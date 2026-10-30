// web/src/lib/current-character-layout.ts
// The one fixed height every current-character chip and its reserved slot must share, so
// a page that reserves space for the chip before hydration (a tool page's own skeleton and
// min-h budget, bulk-skeleton.ts) and CurrentCharacterChip.svelte itself never disagree.
// Two 44px rows on phone (label, then a non-wrapping action row), collapsing to one 44px
// row on desktop -- CurrentCharacterChip.svelte's own header comment. Pulled out of that
// component (fix round 1) so a mounting page can reserve the identical height without
// importing a Svelte component into a plain build-time string.
export const CHIP_HEIGHT = 'h-[88px] md:h-11';

// The current-character strip's own fixed height on the sim pages (leftovers lane A): an
// account-style card row -- the crest or tile and name, then the four doors and Forget. One
// 56 px row inside the card's 2 px of border on desktop (measured 56-58); on a phone the
// identity takes a line and the doors wrap to two (measured 148-150). Fixed so the hydrated
// strip and the static slot that reserves it can never disagree.
export const STRIP_HEIGHT = 'h-[150px] md:h-[58px]';

// The flex gap SimView.svelte's and ToolsView.svelte's own root `<div>` both use between
// the chip and whatever follows it. Fix round 1, Important #2: a static shell's
// pre-hydration markup that stacks the chip slot directly against the next element with no
// gap (a plain, non-flex wrapper) renders everything below the chip 22px/32px higher than
// the hydrated root does -- pulled out here so a shell's own wrapper can reuse the
// identical value instead of a literal that can silently drift from the real root's class.
export const VIEW_GAP = 'gap-[22px] md:gap-8';
