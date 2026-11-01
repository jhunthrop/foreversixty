# F12 search and the item page: ux-designer build verdict

Reviewer: ux-designer, 2026-11-01. Reviewed what renders: `build/ios/*.png` (iPhone, default and AX5) and `build/android/*.png` (default and font scale 2.0) against spec §5.4, §3.2, §6, §7, §9, v1 shared rules and the boards `f12-search-ios.png` and `f12-search-and.png`. No diff was reviewed.

# HOLD

The screens match the boards closely: Binds line, Cloth right-aligned, pick card, class-coloured spec rows, `Pick` and `Runner-up` tags, `Drop chance not known`, grouped results in the right order, the offline strip and the no-results copy are all correct. Three rendering defects and two evidence gaps block the merge.

## Scope calls

| Item | Call |
|---|---|
| Watching group and the star left out (F11 not built) | Acceptable. A group with nothing in it is not drawn, per spec. Characters rows show no star at all, which reads as finished; fine until F11. Re-capture F12.1 and F12.2 with Watching when F11 lands. |
| Characters, guides, Rankings, planner open the website | Acceptable for this lane, with one fix: see SHOULD 5 (no signal that the tap leaves the app). |
| Binds line and Cloth from test data | Acceptable. The capture shows the ship-blocker line the design verdict asked for (`Binds when picked up`, `f12-4-pick-ios.png`, `f12-4-pick-and.png`). Old data must show no line and no gap; the lane's data test owns that. |
| `Binds when picked up` instead of the mock's `equipped` | Accepted. classic-db is the source. Update the board copy when the boards are next touched. |

## MUST (blocks merge)

**M1. iOS, item page, unknown id: a loading skeleton is frozen above the error.** `f12-4-unknown-ios.png` shows the empty icon tile and a grey name bar, then `This item is not in our data.` in muted grey. It reads as a page still loading. Android (`f12-unknown-and.png`) is right: back arrow and the sentence, no skeleton. Spec §5.4: unknown id prints the sentence. Fix: on the unknown state draw no tile and no bar; match Android, with the sentence at the same weight and colour on both platforms.

**M2. iOS, every chip row: stray vertical bars at both ends of each chip.** `f12-1-empty-ios.png`, `f12-2-all-ios.png`, `f12-3-items-ios.png`, `f12-4-pick-ios.png`, `f12-4-loading-ios.png`, `f12-6-offline-ios.png` show a thin vertical line outside each capsule's left and right edge (clearest beside `All`, `Item`, `Best in slot for`). The boards draw clean capsules; Android (`f12-1-empty-and.png`) has none. Likely a rectangular border or focus overlay under a capsule clip. Fix: one capsule shape for fill, stroke and hit area; no rectangle stroke. Re-capture all six.

**M3. iOS, item page, sources failed: the `Best in slot for` heading hangs over nothing.** `f12-4-uses-failed-ios.png`: `Sources did not load.` and `Try again` sit under `Where it comes from`, then `Best in slot for` is a bare label with an empty page below. Both sections come from the same item-uses shard, so the player cannot tell the second is also missing. Spec §7: every state has words. Fix: when the shard fails, draw one error block that covers both sections (one label pair removed or one message per section), with one `Try again`. Do not draw a heading with no body. Add the Android capture of the same state (missing).

**M4. Evidence: the `not on your list` capture does not prove the state.** `f12-4-not-on-list-ios.png` and `f12-4-not-listed-and.png` are the pick page with the card removed. `Elemental Shaman  Pick` is still in `Best in slot for`, which says the item is on the Elemental list. A reader sees a contradiction (a pick that is not yours). Spec acceptance: an item not on your list. Fix: capture a real item that Elemental Shaman lists nowhere. If the card shows only for the active band and the row shows the spec's best placement, add the band to the pick card meta (see S3) so the two agree, and say so in the capture note.

**M5. Evidence: acceptance captures missing.** Spec §5.4 acceptance asks for the following.
- Android F12.2 with all four groups: `f12-2-all-and.png` and `f12-2-all-large-and.png` stop before `Characters`. Capture the Characters group, the `You` tag, and the row shape on Android.
- F12.5 with the `Best in slot for` chip selected on iOS: `f12-5-bis-for-ios.png` is a full-height scroll with `Item` still selected. Android's (`f12-5-bis-for-and.png`) is the right shape.
- Index failed on iOS (only `f12-index-failed-and.png` exists).
- Android loading, runner-up and not-listed exist; Android `f12-4-uses-failed` does not (M3).

## SHOULD

**S1. iOS, search field has no field.** `f12-1-empty-ios.png`: the magnifier, placeholder and `Cancel` float on the page; the board draws a glass field, 44 pt, with a surface and radius. Android (`f12-1-empty-and.png`) draws the pill. Fix: give the iOS field its 44 pt rounded surface (`--surface-card` over the glass material) so the primary control reads as a field. If the simulator drops the material, the fallback fill must still show.

**S2. iOS, item page, source rows use a different glyph from the board and Android.** `f12-4-pick-ios.png` draws a half-filled shield; the board and `f12-4-pick-and.png` draw crossed swords. Fix: one boss glyph on both platforms, the board's.

**S3. Both, pick card meta is thinner on a runner-up.** `f12-4-runner-up-ios.png`, `f12-4-runner-up-and.png`: `Elemental · 60 · Horde`; the pick card reads `Elemental · 60 · Phase 1 raid · Horde`. Spec: `{Spec} · {band} · {preset} · {Faction}`. If the preset is empty, say so in the capture note; otherwise print it. Also the runner-up page still lists `Elemental Shaman  Pick` below, with no sentence that the pick is in another band; one clause in the card meta solves it.

**S4. Both, Recent guide row colour differs.** `f12-1-empty-ios.png` draws `Elemental Shaman` in class blue; `f12-1-empty-and.png` draws it white. Spec row shape: guide name in the class colour. The board draws white in Recent and blue in results. Pick one rule (class colour everywhere is the spec) and apply it on both platforms and on the board.

**S5. Both, rows that leave the app look like in-app rows.** Characters, guides, `Rankings` and `Elemental Shaman guide` open the website but show the same chevron (`f12-1-empty-ios.png`, `f12-2-all-ios.png`). Fix: a trailing external-link glyph on those rows until F9, F11 and F13 land, and remove it per row as each lands. One constant in code, not a per-row edit.

**S6. Both, loading skeleton does not match the loaded page; layout shifts.** `f12-4-loading-ios.png`: tooltip block about 220 pt, sources 120 pt, best in slot 120 pt, no pick-card slot. Loaded (`f12-4-pick-ios.png`): tooltip about 260 pt, sources 143 pt, best in slot about 315 pt, plus a 66 pt card. The page jumps by well over 200 pt on arrival. Android (`f12-4-loading-and.png`) drops the meta line that iOS keeps, shows no chip strip, and its tooltip block (about 400 dp) is taller than the loaded one. Spec §5.4 loading: tile real size, name from the index, rest skeleton; §8 no layout shift. Fix: skeleton heights equal the loaded heights (tooltip 260, sources two rows, six spec rows); both platforms show the same head (index meta kept) and the chip strip; the pick-card slot is reserved only when the character has a list, otherwise none.

**S7. Android, largest text: the tag squeezes the name.** `f12-4-pick-large-and.png`: `Balance Druid` and `Elemental Shaman` wrap to two lines and `60 · Phase 1 raid` breaks as `60 · Phase 1` / `raid` because `PICK` stays on the right. iOS AX5 (`f12-4-pick-ax5-ios.png`) puts the tag under the name, which is the spec §9 pattern (value under the name). Fix: below the font-scale breakpoint, drop the tag and chevron under the meta on Android, as iOS does. Source rows at 2.0 have the same squeeze (`Blackrock Depths ·` / `Dungeon · Drop chance` / `not known`); stack the same way.

**S8. iOS, scope chips do not fit at 402 pt.** `f12-2-all-ios.png`: `Talents` is cut by the right edge, and the board shows all five chips whole. A scroll is allowed, but the cut chip should not be a surprise at the default size. Fix: 15 pt label with 12 pt side padding so five fit at 402 pt; below 380 pt let them scroll. Same on Android, where `Talents` is cut at 412 dp (`f12-1-empty-and.png`).

**S9. Android, offline strip is square and flat.** `f12-6-offline-and.png`: the strip has square corners and a single bottom rule; iOS (`f12-6-offline-ios.png`) and the board draw a rounded 12 pt card. Fix: the same rounded strip component.

**S10. iOS, item page back button has no label.** `f12-4-pick-ios.png`: a bare chevron; the board draws `Search` in the back pill. Fix: back title `Search` when pushed from search, `Back` otherwise, 44 pt target.

**S11. Test data: `You` row is not the current character.** `f12-2-all-ios.png` tags `Elegy · Priest · Level 58` as `You`, while `f12-1-empty-ios.png` says `Lightningbolltz`. Fix the fixture so one capture set tells one story.

## NICE

- iOS AX5 (`f12-2-all-ax5-ios.png`): item and talent rows leave a column about half the screen wide, so names break to three lines. Stack the icon over the text at AX5 or let the meta go under the icon; Android 2.0 is calmer (`f12-2-all-large-and.png`).
- Android `f12-5-bis-for-and.png`: the sticky chips cut the scrolled row with no fade; add a 12 dp fade.
- The expanded spec list (`f12-5-bis-for-ios.png`) has no way back; add `Show fewer`.
- The explanatory lines `A drop chance prints when the game's data has one.` and `Each row opens that list...` are on the board, so they stay; once the third list ships, review whether they still earn their place (tenet 14).
- Android loading shows the title in the bar and the name again in the head; fine per the board, noted for the next bar pass.

## Checklist (spec §5.4)

- [x] Binds line shows a client word (`f12-4-pick-ios.png`).
- [x] `Cloth` on the right of `Back`.
- [x] Name in quality colour; meta `{Quality} · {Slot} · item level`.
- [x] Pick card only when on the list; `Yours` tag.
- [x] `Drop chance not known`, never `0%`.
- [x] Six rows then `5 more specs`; Pick before Runner-up; class then spec order (`f12-5-bis-for-ios.png`).
- [x] All order Guides, Talents, Items, Characters; `Show all {n}` rows; `1 rank`.
- [x] `Nothing matches` copy and the offline strip, verbatim.
- [x] Android search pill, 56 dp.
- [x] AX5 and 2.0 captures on F12.2 and F12.4 exist; icons and quality colours do not collapse.
- [ ] iOS unknown id clean (M1).
- [ ] iOS chips clean (M2).
- [ ] Sources failed covers both sections (M3).
- [ ] Not-on-list proven on a real item (M4).
- [ ] All acceptance captures present (M5).
- [ ] No layout shift on load (S6).

Re-run me after M1 to M5 with new captures for: `f12-4-unknown-ios`, one chip row on iOS, `f12-4-uses-failed` on both, a true not-on-list item, Android Characters group, iOS `Best in slot for` selected, iOS index-failed.
