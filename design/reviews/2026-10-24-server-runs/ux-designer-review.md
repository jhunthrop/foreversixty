# Server runs: after-build review (ux-designer), 2026-10-24

Reviewed on the 48 captures plus the mock render. Read at displayed scale; pixel values below are measured from the captures (1440 and 390 native, others scaled).

## Verdict: SHIP WITH FIXES (no blockers)

## Side-by-side with the mock at 1440
- Grid matches: crest 32, title left, state block right-aligned two lines, `Run again` on the far right, five rows separated by 1 px soft lines, panel width equals the siblings (96 to 1344 at 1440; centred 376 to 1624 at 2000).
- Same order as the mock: heading, list, run block. First row state at y 372 (limit 900); at 390 the first state line is at about y 405 (limit 560).
- Unread row: gold dot, tinted ground. In the mock the tint reads as warm gold; shipped it reads as a neutral brown-grey. Acceptable, still the only tinted row.
- Done chevron: mock 16 px; shipped about 6 px and nearly invisible. Polish (fix below).
- Band line at 1440: 44 px row, sentence 13 px, mono DPS, gold `Open result →`, muted underlined `and 2 more`, dismiss at the right. Matches mock. Dismiss glyph is about 8 px; spec 16 px.
- Mock rows carry titles ("Molten Core drops") and per-spec crests. The fixture has no titles, so the title-over-meta grammar is not proven by these captures.

## Deviations
| # | Element | Spec | Shipped | Severity |
|---|---|---|---|---|
| 1 | Done-row `›` (3.3, 4) | 16 px, goes gold on hover | about 6 px glyph, 1440 and 390; hover capture shows no visible gold step | must fix |
| 2 | Band dismiss `×` (3.4) | 16 px glyph, 44 x 44 target | glyph about 8 px; target not measurable from capture; focus ring at 390 looks about 36 px square | must fix (verify 44) |
| 3 | Loading skeleton (4) | three rows: title 180x15, meta 120x12, state 140x14, in the reserved box | at 390: three blocks of bars, no crest circles, no row separators, no state bar on the right | polish |
| 4 | Error state 390 (4) | `LoadError` inside the reserved box | message and `Try again` sit bare, with no panel; the reserve collapses to about 50 px, so the run block jumps up 30 px when the read fails | polish |
| 5 | Empty state 1440 (3.2, 6) | one `EmptyState` panel at the reserve height | text-only line in a bare panel, 64 px, correct copy. Passes the 64 px rule, but it is a plain text card (tenet) | polish |
| 6 | Phone state line (3.3) | state block under the title: line 1 state, line 2 age | state and age share one line when they fit (done, running without estimate, queued); running with estimate and failed wrap age to line 2 | accepted, see ruling |
| 7 | Phone done chevron (3.3) | muted `›` in the row | top-right, aligned to the title line, not centred on the row | polish |
| 8 | Running dot | gold dot | captured mid-pulse, reads olive. Acceptable for a 1.6 s opacity animation; a reduced-motion capture was not supplied | polish |
| 9 | Known: list above `SimRunBlock` | as spec 3.1 / 11.1 | as spec | none |
| 10 | Known: running and failed rows not links | as spec 8.5 | as spec | none |
| 11 | Known: line only in the band layout | spec 3.4 says on every page with the band | not in the stacked card layout. Ruled by controller; the BiS, Rankings and Logs captures at 1440 show it | accepted |
| 12 | Known: px reserve | amended 3.7 | CLS 0.0006 / 0.0004 (list), 0.0059 / 0.0002 (home) against a 0.02 limit | pass |

### Rulings on the controller's known items
- Phone state and age on one line: accepted. Rows measure 99.5 / 118 / 161 px, all above 64. Condition: the wrap must stay to the two cases above, and the mono age keeps its right-hand place when the line fits. Not a defect.
- List above the run block: accepted; the run block starts at y 707 at 1440 and stays in the first screen at 1440. At 390 with five rows it starts at y 950 (below the fold); with the empty state it is inside the screen. Acceptable because the list is the answer.

## Region checks
- 3.2 heading: `Your server runs`, 15 px section title, nothing else on the row. Pass. Overflow footer line not captured (no 6-row capture): unverified.
- 3.3 copy: `Running · takes about 3 minutes`, `Started 2 minutes ago`, `Finished 14 minutes ago`, `Failed: stopped before completion`, `Ended 2 hours ago`. All verbatim. Overdue (`taking longer than usual`) and `Starting` not captured: unverified, covered only by unit tests.
- 3.4 band line: sentence, `and 2 more`, `Open result →` desktop; `→` on phone with `and 2 more` on its own line. Pass. Dismissed returns to the ready height (dismissed captures match the no-line band). 2000 centred in the 1344 band. Pass.
- 4 states: Run again rest, hover, focus, pending, started and could-not-start are all captured at the specified positions. The error line sits under the state block at 1440 and under the age at 390. Pass. The 402 / `See plans` variant not captured.
- 6 account copy: `On · 4 runs this month` fits one line at 390 (value column about 170 px); underline intact. The fixture shows `Not on Premium` above `On`, a fixture artefact only.

## Section 9 checklist
1. /sim payer, five states, 390 / 1440 / 2000: ticked.
2. Empty, 1440 / 2000 / 390, run block directly under: ticked.
3. Loading and error, 1440 and 390: ticked (see 3 and 4 above).
4. Run again rest, hover (1440), focus, pending, started, could-not-start: ticked.
5. Done row hover and focus at 1440: ticked (hover chevron step not visible).
6. Home line one / three / dismissed at 390, 1440, 2000; BiS, Rankings, Logs 1440: ticked.
7. Band focus on link and dismiss at 390: ticked.
8. Account row four variants plus off, 390 and 1440: ticked.
9. CLS numbers printed: ticked, as supplied.
10. Side-by-side with the mock plus 1024 and 2000: not ticked. No 1024 capture; side-by-side composed here in prose only.

## FIX list
1. Done-row chevron to 16 px, muted; gold on hover and focus; centred vertically on the row at phone width.
2. Band dismiss glyph to 16 px; confirm the hit box is 44 x 44 (focus ring currently reads smaller) and re-capture.
3. Loading skeleton: add the crest circle, row separators and the right-hand state bar; use the reserve height.
4. Error state: draw it in a panel at the reserve height so the run block does not jump on failure.
5. Supply the missing captures: 1024 for the list and the band line, a 6-row overflow footer, an overdue row, a `Starting` row, and a reduced-motion dot.

## Re-check, 2026-10-24 (head b6690e3d)

Fix items:
1. Done-row chevron: pass. 16 px stroke SVG, muted, vertically centred at 1440 and 1024 (rows at y 511 / 576). Titled rows and overflow captures agree.
2. Band dismiss: pass. The `×` now reads at 16 px (home-line-one-1440, x 970 y 208), same tint row, no clipping.
3. Loading skeleton: pass. Three rows, each with a crest square, title and meta bars, a right-hand state bar and 1 px separators, inside the panel (390).
4. Error state: pass. `Your server runs did not load.` and `Try again` sit inside a bordered panel; run block starts at y 453 against 333 + 100 = 433 panel edge plus the 24 px gap, no jump beyond the reported 354x136 / 1248x100 box.

Previously unverified:
- 1024: pass. Desktop grid, no overflow, `Run again` and chevrons aligned, list matches the container width. The header's `Signe...` truncation is pre-existing and not this lane's.
- Overflow footer: pass. Five rows then `Older finished runs are in Your sims below.` at 12 px muted inside the panel. Note: the fixture headline reads `1200 DPS` with no thousands separator. If that is the fallback path (`{dps} DPS`), it must use `formatAmount`; if it is the stubbed API headline, ignore. Check once, polish.
- Overdue: pass. `Running · taking longer than usual` / `Started 6 minutes ago`.
- Starting: pass. Static gold dot, `Starting` / `Started just now`.
- Reduced motion: pass. Dots render solid gold (compare the olive mid-pulse dots in the standard capture); no animation state visible.
- 402: pass. `Server sims are part of Premium.` plus underlined `See plans` under the state block, `Run again` returns to rest.
- Titled rows: pass. Title 15 px 600 over `Frost Mage · Sim` meta, unread dot before the title, matching the mock's row grammar.

Checklist item 10, ticked. Prose side-by-side at 1440 against server-runs-states.png: same page order (static heading, list, run block); list panel 96 to 1344, five rows of 65 px, crest 32 px at x 127, title block left, state block right-aligned with the chevron at x 1321, `Run again` 112 px wide at the far right of the failed row; unread row is the only tinted row and carries the gold dot; band line is 44 px with `Open result →`, `and 2 more` and a 16 px dismiss. Differences from the mock: the unread tint is a neutral warm grey rather than the mock's gold wash (accepted, still unique on the page); the mock's long titles ("Molten Core drops") are shorter in the fixture (no truncation case captured, covered by unit tests).

### Final verdict: SHIP
Remaining items, none blocking: confirm the `1200 DPS` fixture headline is not the unformatted fallback path.
