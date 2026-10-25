# Character history: UX review on screenshots

Reviewer: ux-designer. Build: fixture, head cde23dbf plus working tree. Spec: design/specs/2026-10-24-character-history.md (with the two controller amendments; the guild ring is treated as fixed).

## Verdict: SHIP WITH FIXES

The layout, hierarchy, colours, selection model, gap drawing, lists, 2000 px centring and CLS match the spec. Four must-fix items and a short polish list follow. No blocker.

## Deviations

| # | Element | Spec said | Shipped | Severity |
|---|---|---|---|---|
| 1 | Score list, boss name beside `Private` / `Guild` (history-owner-section-1440, -390) | 3.7: the word follows the boss name, 12 px muted | Boss name is hard-clipped with no ellipsis: `Golemagg the Incinerat` + Private at 1440, `Golemagg the Incin` at 390. The name loses letters to a secondary word. | must fix |
| 2 | Loading at the reserve (history-loading-1440) | 4.2 and 3.11: the panel at its reserve with skeleton bars; the reserved box holds the skeleton | The skeleton panel is about 285 px tall; the rest of the 1197 px reserve is an empty black hole of about 500 px between the card and In game. Anchors the page but reads as a broken page for the whole load. | must fix |
| 3 | Free owner sentence at 390 (history-free-owner-focus-390) | 3.9: one sentence, `See plans` an inline link, 44 px min height | `See plans` is an inline-flex box that wraps to its own line; the closing `.` is orphaned on a third line outside the link, and there is no space before the link (`time.See plans`). The focus ring is a 44 px box around the two words. | must fix |
| 4 | Hover preview (history-hover-section-1440, -1024, -2000) | 3.8 and 9.4: readout previews the nearest point, guide goes warm | The capture is identical to the owner state: readout still 1830, no warm guide, no preview. Either the hover did not reach the plot or the preview does not render. Not verifiable. | must fix (recapture, fix if real) |
| 5 | Sub-headings `Score over time`, `Parse by boss`, list labels (all sections) | 3.3: 15 px 600 `--color-strong`, Barlow | Set in the serif small-caps `section-title` face, the same face as `HISTORY`. The mock shows Barlow 15 px 600. Two levels of heading now look alike. | polish |
| 6 | Parse list, one-kill state (history-single-kill-section-1440) | 4.2: note, then the one-row list under `Kills of Gehennas` | Matches, but the list label starts at y 562 while the score list label sits at 592, so the two lists are 30 px out of line and 330 px of empty column sits above. | polish |
| 7 | Fixture density | Mock: 6 bosses, varied parses and DPS | Capture rows are all 80 / 1,402 / 91% with two bosses. Not a build defect, but it hides colour-ladder and truncation behaviour (two-digit differences, long boss names, 5 pills). | polish (richer fixture for the next capture round) |
| 8 | Step buttons | 44 x 44, chevron glyph | 44 x 44, larger chevron than the mock; right disabled glyph reads dim. Acceptable and clearer than the mock. | none |
| 9 | Gap label position | Inside band, centred | Centred vertically in the band; the mock puts it at the top. Both are within 3.6. | none |

## Region by region

- 3.2 section: heading, caption, one panel with no gold highlight, two columns at 1440 and 2000, one column at 1024 and 390, score first. Chart tops align across columns (both at y 255 in the 1440 crop). Pass.
- 3.3 score column: control-row note, `Rings are private or guild reports.` (the capture is pre-amendment for the ring note; confirm it now reads `Rings are private reports.`), readout 28 px mono in the Epic purple, `11 Dec 2026 · 2 fights`, step buttons, `Up 1090 since 12 Oct`, `Gap 1 Dec to 8 Dec: Score not kept` under the chart when narrow (history-gap-note-390). Pass.
- 3.4 parse column: pills with counts, selected pill gold border, readout `88 parse`, `11 Dec 2026 · kill 5 of 5`, `Up 34 since 17 Nov`. Pass. The sideways fade at 390 is not exercised (two pills fit).
- 3.5 chart: tier bands in quality colours, tier names, boundaries, ring points for private and guild days, 3 px dots, 5 px selected dot in its tier colour, y labels only where reached (1000, 1500; 25 to 100), three x labels at 390, six at 1440, first and last edge-aligned. Parse y starts at 25 for a minimum of 54, per the rule. No animation, no legend. Pass.
- 3.6 gaps: band with dashed edges and `No rated fights` inside at 1440 and 390; narrow gap with the line under the chart. Pass.
- 3.7 lists: eight rows then `Show all 9 fights`; marked rows with gold left edge and ground (two rows on 11 Dec); parse list leads with the date; Spec column omitted; at 390 Date and Spec go. `Private` and `Guild` words present. See item 1.
- 3.8 selection: stepped state moves readout, guide, selected ring point and marked row (history-stepped). Disabled state of the later button at the newest point visible. Hover: see item 4.
- 4 states: loading (item 2), error (`Your history did not load.` and `Try again`, collapses the panel to one row, accepted), focus on plot (2 px gold ring around the whole plot box), focus on step button, focus on pill, single-kill, free owner, visitor (card only), anonymised (404). Pass except as noted.
- 5 hierarchy: the readout is the largest type in the section and the only coloured text outside the parse numbers; the card keeps the gold top highlight. Pass.
- 6 copy: every string seen matches section 6 verbatim, apart from the pre-amendment ring note.
- 7 widths: at 2000 the panel is centred and capped (about 1248 px, charts 588 px); at 1024 one column; at 390 the readout and chart fit between the heading and the first screen end (heading top to axis about 465 px, under the 500 px budget). Pass.
- 9 performance: hinted CLS 0 at 390 and 0.000128 at 1440, under the 0.02 budget. Reserve fallbacks 2120 / 2240 / 1197 px accepted as the measured values; they replace the spec's estimates.

## Side-by-side with the mock at 1440

Same-scale compositions: the build matches the mock's skeleton almost exactly. Panel at 1344 wide, two columns with a 32 px gutter, sub-headings, control row, readout with two steppers on the right, chart with bands, `Up N since` line, list label, list. Chart heights agree (about 255 px each) and the tops align. Three differences a player would notice. First, the mock's sub-headings are plain Barlow; the build's are small-caps serif, so the panel reads heavier at the top (item 5). Second, the mock has six pills and varied rows with colour ladder spread; the build's fixture is two pills and identical rows, so the build looks emptier in the parse column (item 7). Third, the mock's step chevrons are small; the build's are larger and read better. Verdict on composition: faithful, with the gap in the loading state (item 2) the only large visual difference.

## Section 9 checklist

1. Whole section, 390 / 1440 / 2000: done.
2. Gap with reason inside and narrow gap, 1440 and 390: done.
3. Rings and control-row note: done (note text is pre-amendment; re-check after the rebuild).
4. Selection: middle point selected done; hover preview NOT verified (item 4); step buttons rest, focus-visible and later-disabled done; step-button hover, active and earlier-disabled (first point) NOT captured; plot focus-visible done.
5. Picker: default, focus-visible, other boss, one kill done; picker hover, active and the 390 sideways fade NOT captured.
6. Lists: collapsed, expanded (show-all), marked row, Private and Guild rows done.
7. States: loading and error done (loading fails item 2); empty (`No history yet`), one score point and no kills NOT captured.
8. Free owner and visitor at 1440 and 390 done; `See plans` focus-visible done (fails item 3).
9. Account page History row, all three variants: NOT captured.
10. CLS numbers: done.
11. Side-by-side with 1024 and 2000: done here.

Missing captures that gate the next round, in order: hover preview (item 4), the empty / one-point / no-kills states, the account row (three variants). The picker hover and the step-button hover and first-point disabled are not blockers; verify in the same recapture.

## FIX list

1. Boss name never clips for the Private or Guild word: ellipsis the name only after it has taken the width it needs, give the word `flex-none`, or place the word on the date line at 390.
2. Loading skeleton fills the reserved box (grow the chart skeleton or give the panel `min-height` equal to the reserve), so no 500 px void.
3. Free-owner sentence: keep `See plans` inline (not inline-flex), a plain space before it, the full stop inside the sentence flow, and the 44 px target by padding.
4. Recapture the hover preview with the pointer over the plot; if the preview and warm guide still do not show, fix them.
5. Polish: sub-headings to Barlow 15 px 600 `--color-strong`; align the one-kill list label with the score list label; richer fixture (long boss name, five pills, varied parses) for the next capture round.

## Re-check, 2026-10-25, head cca3bc36

Verdict: **SHIP**. No blocker and no must-fix remain. Two polish notes, neither gating.

| FIX item | Result | Evidence |
|---|---|---|
| 1 Boss name beside `Private` / `Guild` | Fixed | At 1440 the name ellipsises (`Golemagg the Inciner…`) beside a non-shrinking `Private`; the `Guild` row fits whole. At 390 the word drops to a second line under the full name (history-picker-fade-section-390, history-no-kills-section-390). |
| 2 Loading skeleton fills the reserve | Fixed | history-loading-1440: skeleton panel runs from the caption to about 1376 px, with sub-heading, control row, readout, chart and eight list rows per column; no void before In game. history-skeleton-hinted resolves in place (hinted CLS 0 at 390, 0.00276 at 1440, under 0.02). |
| 3 Free-owner sentence at 390 | Fixed, one polish | The sentence reads `History charts draw your score and parses over time.` with `See plans` on its own underlined line; the orphan full stop is gone. Polish: the 44 px focus box overlaps the bottom of the sentence line above it; add 4 px of space above the link. |
| 4 Hover preview | Fixed | history-hover-1440 shows `1240 Rare`, `17 Nov 2026 · 1 fight` and the warm guide at 17 Nov while the selection stays on 1830; the list marks do not move. |
| 5 Sub-headings | Fixed | `Score over time`, `Parse by boss` and the list labels are Barlow 15 px 600, clearly below the serif section heading. |
| 6 One-kill label alignment | Fixed | `Kills of Gehennas` now sits at the same y (592) as `Fights behind the score`. |

Previously unverified states:

- Empty (history-empty-section-1440): one panel, `No history yet. Each rated fight you upload adds a point.` 13 px muted, no chart. Pass.
- One point (history-one-point-section-1440): readout 1830 Epic `11 Dec 2026 · 1 fight`, both step buttons disabled, `First score recorded 11 Dec. A line needs a second point.`, the one-row list. Pass. Polish: the score list label sits 30 px above the parse list label because the score column has no chart; acceptable.
- No kills (history-no-kills-section-390): `No kills on record yet. Parse charts start with your first kill.` under `Parse by boss`, no picker; the score column is unaffected. Pass.
- Picker with eight bosses (history-picker-fade-section-390 and history-picker-hover-section-1440): pills never wrap, scroll sideways, fade at the right edge, the selected pill is first and gold-bordered, one chart and list per boss, the parse y axis starts at 50 for a minimum of 57 as the ladder rule says. Pass. The pill hover border is too faint to tell from rest in the 1440 crop; not a gate.
- Step buttons: hover, disabled and focus-visible captures exist and match 4.1 (the later button dims at the newest point). The first-point disabled state appears on the one-point capture. Pass.
- Ring note: `Rings are private reports.` Pass.
- Account row: `Open on your main character` (link, wraps to two lines in the 1440 plan column, acceptable), and `Sparkline only` for free. The no-main `On` variant was captured and reads as the spec describes. Pass.

### Section 9 checklist

1. Whole section 390 / 1440 / 2000: ticked.
2. Gap inside the band and narrow gap: ticked.
3. Rings and the control-row note: ticked.
4. Selection: middle selected, hover preview, step rest / hover / focus / disabled, plot focus: ticked.
5. Picker: default, hover, focus, one-kill, another boss, 390 fade: ticked (active is a 1-frame state, unverified, not a gate).
6. Lists: collapsed, expanded, marked, Private, Guild: ticked.
7. States: loading, error, empty, one point, no kills: ticked.
8. Free owner, visitor, `See plans` focus: ticked.
9. Account row, three variants: ticked.
10. CLS: ticked (0 and 0.00276).
11. Side-by-side with 1024 and 2000: ticked.

Remaining polish only: 4 px above `See plans` at 390; a visible pill hover border. Both can ride the next pass.
