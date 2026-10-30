# Build review: report charts (UX), 2026-10-30

Spec: design/specs/2026-10-30-controls-charts.md Part B. Reviewed from the captures in this folder against the boards. The Players control is exempt as briefed.

## Verdict: HOLD

## Blockers

1. Damage chart head on a phone (build-charts-damage-390.png): the Window select is squeezed to about 58 px and shows a clipped `\` instead of `Whole fight`. The head wraps Players, Window and Log scale badly. Spec B5.2 and the 44 px dense field rule: every head field must show its value in full at 390. Give the Window field its own row (as the Enemy field has on the threat chart) or a min width that fits `Whole fight`.
2. Threat chart head names the enemy two ways (build-charts-threat-1440.png and -390.png): the title reads `Estimated threat on Ragnaros` (desktop) while the Enemy field says `Every enemy`. On a phone the title drops the name and the field says `Every enemy`, so the one thing the spec says the field must do (name the enemy) fails. The title, the tooltip phase line and the field must agree. Make the field show the enemy the chart is drawn for, or make the title `on every enemy`.

## Shoulds

1. Players field reads `Top 8 of 3` on the 3-player timelines and resources captures. Copy must not name a limit above the count. Use `All 3` when N is 8 or fewer.
2. Legend `Only` on hover (build-charts-damage-2000.png, Ashka chip) overlaps the value and truncates it to `42`, and shifts the next chip. Reserve its width or place it after the value without clipping.
3. Threat table lists all 40 rows (build-charts-threat-1440.png), so the page runs about 3,500 px. The board shows six rows and `34 more rows`. Restore the cap with the show-more row so the tab keeps its answer first.
4. Tooltip event order differs from the board: build lists death, taunt, target change; board lists death, target change, taunt. Pick one rule (time, then kind) and state it in B5.6.
5. Phone threat capture: the death, taunt and target-change markers stack at one x and the leader line runs through the icon group, so the three read as a smear. Apply the 22 px step to all three or fold them into a counted cluster.
6. Axis end label `1:00` overhangs the plot's right edge at 1440 on the timelines and resources captures (and the lane axis below it sits 13 px right of the chart axis). Right-align the last tick inside the frame and share the plot's left and right insets with the lanes.
7. Resources capture: the rows under the chart have no column header and the first row's top rule is cut; the `Damage done` chart above a Resources table repeats the page's one-chart rule only on phones. Confirm this is the intended desktop state.
8. Timelines lanes still have the one-line readout and no tooltip (lane report, Open). Acceptable for this pass; schedule it.
9. Not captured: loading, empty, error, wipe fixture (D15), the 360 legend sheet (D17), log scale (D13), reduced motion (D18). Attach them before final sign-off.

## What passes

- Phone fold: the threat chart ends near 580 px at 390 x 640 with Source and Table pinned at the top, one legend row with `+38`, the hint in the event line, the caveat under the chart. Matches the 581 px stack.
- Tooltip: sorted, class-coloured, events first with the killing blow and overkill, dead last, share column, flips left in the right 55 percent, sits above the plot.
- Legend: pinned chips with the pin glyph, `YOU`, raid median first, `+n MORE`.
- Markers and phase bands: death, cast, taunt and target-change icons in the lane; `RAGNAROS` and `SUBMERGE` labels backed; submerge flat at zero.
- Brush: window handles, whole-fight strip with death and phase ticks, `Whole fight` ghost button when zoomed.
- 2000 px: chart stays in the 1248 px column, centred.
- Consistency with /account and the alignment pass: quiet card frame, 12 px radius, section-title head, same field recipe, gold focus ring on the open Players filter.

## Re-check (fix round 2)

Verdict: SHIP.

- Blocker 1 resolved. build-charts-damage-390.png: Players, Window and Log scale each sit on a full-width row, and Window reads `Whole fight` in full.
- Blocker 2 resolved. build-charts-threat-1440.png and -390.png: the title reads `Estimated threat, every enemy` and the Enemy field reads `Every enemy`. The tooltip head agrees but truncates at 1440 (`estimated threat, every e...`). Polish.
- Nothing new from the fix round. The 390 fold holds, the phone head drops its meta line, and the legend lists pinned players first.
- Still open from the shoulds, unchanged: the threat table lists all 40 rows, and the state captures D13, D15, D17 and D18 are missing. Neither blocks.
