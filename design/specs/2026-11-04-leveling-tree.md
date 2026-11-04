# Leveling tree on spec guides (2026-11-04)

Owner request: show the leveling build in a normal talent tree like the other builds, and start at level 10 so the player sees the whole journey. Build window: about three hours, landed by 20:30 UTC. Every decision below is final.

## 1. The question and the answer above the fold

Question: "What do I have in my talents at level N, and what do I take next?"

Answer visible without scrolling once the player reaches the block (the block is a section, not the page top; #leveling still lands on Levels 1 to 19):

- 360 px phone: the caption ("Level 60: Name, rank n"), the level control row (44 px), and the Load in planner link are all visible together, then the first tree starts directly under them. The tree fills below; the player scrolls the three stacked trees, but the control stays above them and the caption always names the newest point.
- 1280 px desktop: caption, control, link and all three trees (three columns) in one view of 900 px height.

The control sits ABOVE the trees on every width (caption first, then control, then trees, then the link). The link sits under the control row, not under the trees, so a phone player never scrolls past three trees to find it.

## 2. Reference

- In-game talent frame (the 3-tree panel at the character window): borrowed as the tree itself, through the existing read-only TreeGrid. Nothing re-drawn.
- Wowhead talent calculator "level N" slider and Classic Era talent planners that step through one point at a time: borrowed as one slider with a point-by-point step and a caption that names the point just taken.
- The guide's own "Level 60 build" and "Raid build" trees under Talents and builds: borrowed whole (same columns, same tooltips, same skeleton and error states).

Not borrowed: OrderStrip, BandCompare, any planner chrome. The planner bundle is at budget; this block imports TreeGrid only, which GuideBuildTree already imports.

## 3. Layout, region by region

Placement. In `web/src/pages/guides/[class]/[spec].astro` the Leveling section becomes, in this order:

1. `## Leveling` heading (unchanged), 2. "Levels 1 to 19" prose (unchanged), 3. NEW page-owned `### Talents by level` heading with id `talents-by-level` and the block below, 4. the markdown `### Levels 20 to 60` heading and its prose (unchanged, no markdown file edited).

Reason for the move: the markdown heading says Levels 20 to 60 and its prose is about 20 to 60 in every guide. The tree starts at 10, so it gets its own heading rather than a mislabelled one. The old strip's `splitAtLevelingStrip` call keeps working: render `levelingEarlyHtml`, then the new block, then `levelingStripHeadingHtml`, then `levelingLateHtml`.

Intro line under the h3 (13 px, `--color-nav`, margin 0 0 12px): "One talent point a level, from 10 to 60. Move the level to see the build at that point." Static HTML, server rendered, no island needed.

The block is one island, `GuideLevelingTree.svelte` (client:visible), inside a container `data-testid="guide-leveling-tree"`. Regions top to bottom:

A. Caption (primary of the block).
- Line 1: 16 px / 22 px, 600, `--color-strong`: "Level 23: Improved Gouge, rank 2". The level number is `numeric`. The talent name is in `--color-strong`; no quality colour (talents have none).
- Line 2: 13 px / 20 px, `--color-nav`: "14 points spent, 37 to go" (spent = level minus 9; to go = 60 minus level; 0 at 60 reads "14 points spent" style: "51 points spent, all placed").
- Reserved height: 42 px from 640 px, 64 px below 640 px (line 1 may wrap to two lines on a phone). Always reserved, so a long name never moves the control.
- `aria-live="polite"`, `aria-atomic="true"`, `data-testid="guide-leveling-caption"`. This is the only live region.

B. Level control (a row, 44 px tall, gap 8 px, margin 8px 0 12px):
- Step back button, 44 x 44, glyph `Glyph` chevron-left, `aria-label="Previous level"`.
- Slider: native `<input type="range" min="10" max="60" step="1">`, flex 1, `aria-label="Level"`, `aria-valuetext="Level 23"`. Track 6 px tall, `--color-line`, filled part `--color-gold`; thumb 24 px circle `--color-gold` inside a 44 px tall hit area (the input is 44 px tall, thumb centred). No existing range control in components/ui; OverviewBrush.svelte is a time brush and is not reused. Style this one inline in the island with `<style>`, using only the tokens above.
- Step forward button, 44 x 44, chevron-right, `aria-label="Next level"`.
- Level readout right of the buttons is not needed: the caption owns the number. Under the track, aria-hidden 11 px `--color-nav` ticks are not drawn. Keep it plain.
- Buttons use the secondary ground and border from the Switch/Select family: faint white ground, 1px `--color-line`, 8px radius, 120 ms hover to a 30% gold border.

C. Load in planner link: the existing `ActionLink.svelte` ("Load in planner", 14 px 600 gold, trailing chevron, 44 px tall), directly under the control row, left aligned. `href` = `/planner?spec=<fullSpecId>&talents=<encoded string for the chosen level>`. `data-testid="guide-leveling-load"`. Its `aria-label` follows the level: "Load the level 23 build in the planner".

D. Trees: new shared `GuideTreeColumns.svelte` (extracted from GuideBuildTree's ready branch: the `grid grid-cols-1 gap-4 md:grid-cols-3` wrapper, one `h3.section-title` per tree, `TreeGrid readOnly`). Three columns from 768 px (md), stacked below. Tooltips identical to the other builds because it is the same TreeGrid. No numbering on cells (TalentCell is in the planner bundle and is not touched). The tree on screen is the whole story: lit cells are the points spent; the caption names the newest one.

Trees that hold zero points at the chosen level stay drawn, dimmed by TreeGrid's own empty state (as the other builds' off-spec trees are). Tree headings show the tree name only, as today.

E. Hierarchy. Primary: the caption. Secondary: the slider. Tertiary: the link (gold text, not a filled button; the only filled gold button on a guide page is the rail's). One emphasis level: gold is used only for the slider fill/thumb and the link.

## 4. Behaviour and shared code

New and changed files:

1. `web/src/lib/guides/leveling-builds.ts`: add `guideLadderSteps(build, classSlug, spec, ownTreeIndex, code)` returning `{ steps: {id:number; name:string; rank:number}[]; strings: string[] }`, 51 entries each, index i = level i+10. It calls `climbLadder` once; `rank` is the count of that talent id in order[0..i]; `strings[i]` is the talent string after i+1 points (`row.join('')` per tree, joined by `-`, snapshotted inside the climb, same format as `ladderTalentString`). `guideLadderOrder` and `guideLevelingBuilds` stay if anything else imports them; delete them only if the grep is empty. A test holds `strings[level-10]` equal to `ladderTalentString(..., level)` for every level 10 to 60 for Combat rogue, Shadow priest and BM hunter, and `strings[50]` equal to the guide's level 60 digits.
2. `web/src/lib/guides/tree-store.ts`: extract `loadGuideTreeStore(code, order?)` from GuideBuildTree's `load()` (decode, `loadTalents`, `indexTalents`, `orderFromRanks`, `createPlannerStore({readOnly:true})`, `setTalents`). Returns `{ store, talentFile }`. GuideBuildTree calls it with no `order`.
3. `web/src/components/guides/GuideTreeColumns.svelte`: the columns markup above. GuideBuildTree uses it too.
4. `web/src/components/guides/GuideLevelingTree.svelte`: props `code` (the guide's level 60 FS1), `spec` (fullSpecId), `steps`, `strings`. State: `level` (default 60), `status`. On ready it holds the talent file and the full ladder order (`steps.map(s => s.id)`). On a level change it builds a fresh read-only store with `order.slice(0, level - 9)` (the ladder obeys the planner's tier and prerequisite rules, so every prefix is legal) and passes it to `GuideTreeColumns`. The talent file is fetched once.
5. `[spec].astro`: replace `LevelingBandStrip` with `<GuideLevelingTree client:visible ... />`. Delete `LevelingBandStrip.astro`, `leveling-talents.ts` (and its test), the copy keys `levelingLegend`, `levelingNewTitle`, `levelRangePoints` if unused, and update any e2e test that reads `guide-leveling-strip` or `guide-leveling-row-*` to the new test ids. Do not remove `plannerHref` building from the page if the rail still uses it.

Server render (Astro renders Svelte on the server before the island loads): the caption at level 60 from `steps[50]`, the full control row at level 60 (slider value 60, buttons enabled-looking but see below), the Load in planner link with the level 60 string, and the tree area as `Skeleton` at the ready height. After load, the same elements stay, the tree area swaps skeleton for one `.reveal`, and nothing else moves.

Default level: 60. Reason: it is the finished build the Talents and builds section already shows, the server HTML and the hydrated HTML are then identical (no mismatch), and it works for a visitor with no character. No signed-in or URL-driven default in this pass.

## 5. States of every interactive element

Slider:
- Default: track `--color-line`, fill gold to the thumb, thumb gold 24 px.
- Hover (fine pointer only, `@media (hover: hover) and (pointer: fine)`): thumb scales to 28 px, 120 ms.
- Focus-visible: 2px gold ring at 2px offset around the thumb (`outline` on `::-webkit-slider-thumb` and `::-moz-range-thumb` via box-shadow `0 0 0 2px var(--color-bg), 0 0 0 4px var(--color-gold)`).
- Active (dragging): thumb 28 px, fill follows.
- Loading (before the talent file arrives): the slider is `disabled`, 45% opacity, not-allowed cursor, out of the tab order; the caption and link are already correct for level 60 and stay. Changing nothing visually but the opacity when it enables.
- Disabled: as loading. Empty: does not occur (51 steps always exist; if `steps` is empty the page renders no block at all, same rule as today's `levelingRows.length > 0`).
- Error: the tree area shows `LoadError` (message "This build did not load", the existing `guidesCopy.treeLoadFailed`, Retry button). The caption, slider and link are hidden in this state only if the failure is a decode failure; for a fetch failure they stay, the slider stays disabled, the link stays live (it needs no fetch). Container height: the error box has the same min-height as the skeleton, so nothing moves.

Step buttons: default secondary look; hover warm border; focus-visible 2px gold ring; active 98% scale off, ground at 8% white; disabled (Previous at 10, Next at 60, or while loading) 45% with `aria-disabled` not removed from DOM, `disabled` attribute set, out of the tab order.

Load in planner: ActionLink states as in the design system (hover lightens gold, chevron 2px right; focus-visible 2px gold ring; active `--gold-hover`).

Keyboard: Tab order is Previous, slider, Next, Load in planner, then into the trees (TreeGrid's roving tabindex, unchanged). On the slider: Left/Down -1, Right/Up +1 (native), PageUp/PageDown +/-10 (set `onkeydown` to do this explicitly, since browsers differ), Home 10, End 60. Cells in the tree remain focusable, show tooltips on focus as in the other builds; Enter on a cell does nothing (read-only).

Touch parity: both buttons are 44 x 44 and repeat on hold is NOT implemented (tap per level; the slider covers fast moves). Dragging the thumb updates the tree live on `input`, throttled with `requestAnimationFrame`. Cell tooltips open on tap exactly as in the other builds.

Live region: only the caption. Updates on every `input` event would flood a screen reader; update the caption text immediately but set `aria-live="off"` while the pointer is dragging (pointerdown to pointerup) and back to `polite` on release, and on each key or button press it is polite. The slider's own `aria-valuetext` is "Level 23".

## 6. Copy, verbatim

- h3: "Talents by level"
- Intro: "One talent point a level, from 10 to 60. Move the level to see the build at that point."
- Caption line 1: "Level 23: Improved Gouge, rank 2" (rank is the rank the talent holds after that point)
- Caption line 2: "14 points spent, 37 to go"; at level 60: "51 points spent, all placed"; at level 10: "1 point spent, 50 to go"
- Slider label: "Level"; valuetext: "Level 23"
- Buttons: "Previous level", "Next level"
- Link: "Load in planner"; aria-label "Load the level 23 build in the planner"
- Tree error: "This build did not load" (existing); retry label from LoadError, unchanged
- Skeleton: no text; `aria-busy="true"` on the tree area while loading

Singular rule: "1 point spent", otherwise "N points spent". No em dashes anywhere.

## 7. Phone and desktop

- 390 px: caption (64 px reserved), control row 44 px with buttons 44 x 44 and the slider taking the remaining 390 - 36 (page gutter 18 each side) - 88 (buttons) - 16 (gaps) = 250 px, Load in planner link, then three stacked trees at full width. The control never collapses or hides; no horizontal scroll at 360.
- 1440 px: container width is the article column (Content layout, with the sticky rail beside it at lg); caption 42 px, the control row max-width 560 px left aligned (the slider does not stretch across a 900 px column), three tree columns from md.
- 2000 px: same as 1440, container centred by the layout, nothing stretches.
- Never collapses: caption, slider, buttons, link.

## 8. Performance and polish

- No layout shift: the server renders the caption, control row and link at final size; the tree area is a `Skeleton` with the same `min-h` as the ready grid (measure the ready three-column height at 1440 and the stacked height at 390 on Combat rogue, take the largest of the three specs, and set `min-h` per breakpoint accordingly). CLS must read 0.00 in the capture. Level changes do not change tree heights (same trees, same cells).
- No missing-icon flash: TreeGrid and TalentCell already load icons for the whole tree; because the store is rebuilt per level, the tree DOM must NOT be remounted. Key the `{#each}` on `tree.id` as GuideBuildTree does, and pass the new store as a prop so cells update in place.
- Page weight: props add about 6 KB to the HTML (51 steps and 51 strings). Acceptable; do not send per-level objects of full trees.
- Lighthouse: `web/lighthouserc.json` covers `/guides/warrior/fury.html`; the page already ships the guide tree island, and this adds one small Svelte component and no new library. Budgets must stay green; run the lhci shard that includes that URL.
- Reduced motion: no animation except the existing `.reveal`, which reduced-motion already switches off.
- Remove the old strip's global style block with the component.

## 9. Acceptance screenshots (implementer attaches)

Specs: Combat rogue, Shadow priest, BM hunter. Viewports: 1440 and 390 for each (six captures minimum), plus one 2000 capture of Combat rogue.

States per spec/viewport pair:
1. Default (level 60), fully loaded, block top to link to end of the third tree.
2. Level 10: caption "Level 10: <talent>, rank 1", all trees showing a single lit cell.
3. A mid level (30) after a keyboard step with focus ring visible on the slider.

Extra captures (Combat rogue 390 only): loading (skeleton, slider disabled), error (LoadError), a talent tooltip open on a lit cell at level 30, Load in planner `href` shown in a devtools crop or asserted in the test.

Also attached: a CLS reading of 0.00 for `/guides/rogue/combat` and a unit-test result for `guideLadderSteps` against `ladderTalentString` for all 51 levels on the three specs.

## 10. Checklist for the review

- [ ] The old chip strip and its legend are gone; no `guide-leveling-strip` anywhere.
- [ ] `#leveling` lands on Levels 1 to 19; the h3 "Talents by level" follows that prose; the markdown "Levels 20 to 60" heading and prose follow the block.
- [ ] Caption names level, talent and rank; line 2 counts are right at 10, 30, 60.
- [ ] Slider, both buttons and keys (arrows, PageUp/PageDown, Home, End) change the tree and caption together.
- [ ] Load in planner `href` matches `ladderTalentString` for the chosen level.
- [ ] Tooltips on the leveling tree match the ones on the Level 60 build tree.
- [ ] 44 px hit targets at 390; no horizontal scroll at 360.
- [ ] CLS 0.00; no icon flash on level change; lhci green.
