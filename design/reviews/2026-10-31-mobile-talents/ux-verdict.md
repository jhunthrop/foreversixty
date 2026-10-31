# F6 Talents, built review (ux-designer), 2026-10-31

Reviewed: the captures in this folder against `design/specs/2026-10-31-mobile-talents.md` and the f6 boards. Rendered output only.

## Blockers

1. **Home row and F6 disagree.** Home's Talents row reads `1 of your 51 points differ from the 50 to 59 guide build` (home-talents-row, both platforms). F6 for the same character reads `8 of your 49 points differ ...` (app-talents-differs). Level 58 allows 49 points, so 51 is wrong on its face. Spec section 3 and section 13: one `TalentCompare` feeds both. The count the player taps through must be the count they land on.
2. **Home Talents row collapses at AX5 on iOS.** `app-talents-differs-large-ios` shows the row text one to four characters per line ("Tal / ent / s", "diffe / r"), with the DIFFERS tag eating the row width. The tag must stack above the text from AX1 (spec 10). Unreadable, so a blocker even though the row is Home's.
3. **The AX5 F6 tree capture on iOS is not F6.** `app-talents-differs-large-ios` and `-large-tree-ios` are both the Home screen. The iOS large-text F6 status, stacked tabs and tree are unproven in the running app. Only the golden `talents-differs-large-ios` covers them, and the golden is correct (stacked segments, stacked tab rows with gold rail, dots, tag above sentence).

## Must fix

4. **Captures missing from section 14:** `talents-differs-360`, `-se`, `talents-differs-restoration-tab`, `talents-rank-large` and `-large-end` (only the right-column large golden exists, Android; none for iOS), and no app capture of the iOS or Android stacked tabs at large text in the running app (Android large-tree shows the scrolled tree only). The 360 and SE captures are the fit proof for pitch 93 and the first-screen answer.
5. **Home row copy** ends `guide build` with no full stop on both platforms. Spec revision 1 item 10: `... guide build.` Either the row is a deliberate exception (then amend the spec) or fix it.

## Polish

6. iOS rank sheet at `.medium`: about 45 percent of the sheet is empty ground below the panel (app-talents-rank-ios, goldens). Known deviation. Accepted for now, but a content-height detent is the spec and the empty slab reads as unfinished. Keep `.large` available.
7. Android rank sheet: about 100 px between the drag handle and the header row; the iOS gap is about 45. Tighten to the spec's 20 side / 16 top rhythm.
8. iOS scrolled: faint 1 px fragments of scrolled content (two short dashes at y 265, app-talents-differs-scrolled-ios and rank-ios) show between the solid bar and the pinned segments. The pinned section ground should extend to the bar.
9. `talents.loading` renders `Loading talents...` with three dots; the spec value is `Loading talents…` (one ellipsis character).
10. Talents row glyph differs: iOS sparkles, Android star. Fine if the shared icon map says so; confirm.
11. The error golden leaves the whole screen below the footer empty. Acceptable; no layout shift.

## Region by region

| Region | Spec | Shipped | Verdict |
|---|---|---|---|
| Bar | iOS inline centred title, solid ground; Android TopAppBar | Matches both | OK |
| Status block | Tag + sentence, one line, wraps together; Matches green, Differs gold | Matches; Differs wraps to two lines at default; tag stacks above at AX5 | OK |
| Unspent / missing / beyond lines | meta muted, in order | Present, in order, beyond line renders on the 49-point case | OK |
| Segments | 44 iOS rounded rect, 48 Android pill with check | Matches | OK |
| Tabs | Equal width, gold underline, mono points, 8 pt dot on trees with dashes | Matches; Arms dot in app capture, Elemental and Restoration in the vector; none on Matches or Guide | OK |
| Pinned compare bar | Pinned when scrolled, not at AX1+ | Pinned on both; not pinned at large on Android | OK (item 8) |
| Tree cells | 48, green maxed/filled, open full face, locked 40 percent, no gold | Matches; locked cells read dim, open cells full | OK |
| Badges | Mono n/max, hang 6 | Matches, readable at 360-equivalent Android | OK |
| Dashes and chips | wipe over, gold under, signed chip top left | Matches: +2 red, -2 gold, chips legible, ring does not clip | OK |
| Prerequisite lines | Under cells, arrowheads, gold when met | Matches; grey vs gold distinction visible | OK |
| Legend, footer, link | Legend when dashed, tapHint otherwise; link to /planner with external glyph | Matches | OK |
| Capstone / scroll | Tier 7 half under the bar in Differs, accepted | iOS half hidden under the floating tab bar; Android visible. Accepted by spec item 8 | OK |
| Rank sheet content | Icon 56, mixed-case name, Rank n/max, panel, Next rank, wipe requirement, dashed guide line | Matches on both; requirement line correct (`Requires 30 points in Elemental Talents`, `Requires 5 points in Enrage`); guide box dashed gold; maxed has no Next rank | OK |
| Tapped cell visibility | Ringed, 28 above sheet top | iOS and Android both ring the cell and keep it above the sheet | OK |
| Loading | Skeletons at real heights, 14 cells, no spinner | Matches (item 9) | OK |
| No export / no spec / matches / error / offline | Per section 6 | Copy verbatim; offline strip on Android correct | OK |
| Large text sheet | Single full-height detent, wraps, scrolls | Android golden wraps to 7 lines each, reaches the guide box | OK |
| Home row chevron | Talents only | Present, Sims untouched | OK |

## Known deviations judged

- iOS sheet detents `[.medium, .large]`: accepted, polish item 6.
- Home chevron on Talents only: accepted, correct affordance.
- No golden for stale export or beyond: accepted; beyond appears in the app captures with the right sentence.

## Checklist

- [ ] Reached only from Home's Talents row; stack pops on 401 (not shown; entry is shown, 401 not captured)
- [x] Guide band by the planner's rule (50 to 59 band chosen for level 58, 60 for the vector)
- [ ] Count equals `diffCount`, `Matches` iff zero dashes, totals summed: true on F6, false against Home (blocker 1)
- [x] 4 columns, 48 icons, 7 tiers reserved, nothing clips (360 and SE captures missing, so only partly proven)
- [x] Cells, badges, lines with arrowheads, chips on Your build only, no gold on Your build, dot on tabs
- [x] Sheet content, guide line only on difference, tapped cell stays visible
- [ ] Large text: F6 holds on Android and in the iOS golden; Home row breaks at AX5 (blocker 2); iOS running-app capture missing (blocker 3)
- [x] Every section 6 state has a render (stale and beyond-band excepted)
- [x] No gear, DPS, weights or sim figure on F6; footer links to /planner
- [ ] Section 14 captures attached in full (item 4)

VERDICT: SHIP WITH FIXES
