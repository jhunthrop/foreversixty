// web/src/lib/current-character-layout.ts
// The one fixed height every current-character chip and its reserved slot must share, so
// a page that reserves space for the chip before hydration (a tool page's own skeleton and
// min-h budget, bulk-skeleton.ts) and CurrentCharacterChip.svelte itself never disagree.
// Two 44px rows on phone (label, then a non-wrapping action row), collapsing to one 44px
// row on desktop -- CurrentCharacterChip.svelte's own header comment. Pulled out of that
// component (fix round 1) so a mounting page can reserve the identical height without
// importing a Svelte component into a plain build-time string.
export const CHIP_HEIGHT = 'h-[88px] md:h-11';

// The flex gap SimView.svelte's and ToolsView.svelte's own root `<div>` both use between
// the chip and whatever follows it. Fix round 1, Important #2: a static shell's
// pre-hydration markup that stacks the chip slot directly against the next element with no
// gap (a plain, non-flex wrapper) renders everything below the chip 22px/32px higher than
// the hydrated root does -- pulled out here so a shell's own wrapper can reuse the
// identical value instead of a literal that can silently drift from the real root's class.
export const VIEW_GAP = 'gap-[22px] md:gap-8';
