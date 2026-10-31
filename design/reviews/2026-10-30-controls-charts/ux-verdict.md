# UX verdict: controls and charts spec (second designer)

Reviewed: design/specs/2026-10-30-controls-charts.md and the seven boards in design/mocks/2026-10-30-controls-charts/renders/. Read-only; no code touched.

## Verdict: HOLD

The direction is right (own listbox in the top layer, uPlot, one chart per tab, drag/brush/URL as one window). Two items make the spec misreadable or fail its own first rule. Fix them and it ships to build.

## Blockers

1. **The phone fold promise (B1.4) is not met by the board, and the spec's number leaves out the page above the chart.**
   - Spec: at 360 x 640 the threat head, Enemy field, legend, plot and brush fit in "about 520 px".
   - Board (controls-charts-phone.png, 390): the card alone runs from y 285 to y 865, about 580 px, and the Source field plus the tab strip sit above it (about 130 px) under the sticky header. The brush is below the fold at 640 px tall, and the answer to "what happened" (the plot) starts about 150 px down.
   - Required: the spec states one layout that fits. Suggested: Enemy field inline in the head row at 44 px beside a shortened title; legend capped to one row of chips plus `+{n} more`; plot 160 px; the scrub hint removed from the first view. Name the total height and measure it with the Source row and tabs included. Update the board to match.
   - Also undefined: what `+2 more` does when tapped. State it (expands the legend, or opens a list of off series).

2. **The searchable list contradicts itself on where focus lives (A5 vs A6).**
   - A5: "focus stays on the trigger, aria-activedescendant names the row."
   - A6: in a searchable list "the filter input is the role=combobox with aria-activedescendant".
   - Both cannot hold. The trigger would be a combobox with `aria-expanded` while focus sits in a second combobox. An engineer will pick one and screen readers will get either a dead trigger or two comboboxes.
   - Required: for searchable lists, say focus moves to the filter input on open; the trigger keeps `aria-expanded` and `aria-controls` but is not the active-descendant owner; the list is `aria-labelledby` the field label; focus returns to the trigger on close. For plain lists, focus stays on the trigger. Add the axe check for both variants in section D.

## Shoulds

Dropdown

1. **Hover sets the active row (A4.3) and Tab commits the active row (A6).** A mouse parked over the list changes the value when the user tabs away. Track pointer-active and keyboard-active separately. Tab commits only a keyboard-moved row; a parked pointer does not.
2. **Where the popover mounts is not specified.** Say the list is a child of the field's wrapper, not `document.body`, so it stays interactive inside any modal dialog or focus trap (RaidForm, sim sheets, GuildSettings). Add: Escape closes the list and stops propagation so it does not close the surrounding sheet.
3. **Hidden native select and validation.** A hidden `<select required>` cannot show the browser's validation bubble. State that the trigger takes the error state and focus when the form is invalid.
4. **Touch on a fine-pointer device** (a touchscreen laptop with a mouse) gets 36 px rows. Give rows min-height 44 px when the last input was touch (`pointerType`), or state that 36 px is accepted.
5. **Dense field text is 13 px; its list rows are 14 px.** The text size jumps on open. Use 13 px rows in the `dense` list, or justify it.
6. **Closed field vs /account.** The list surface matches `.csel-panel` (card surface, 12 px radius, highlight, gold left edge). The closed field does not: the account trigger is a 999 px pill with a gold glow when open (`--glow-gold`) and a `--color-strong` chevron; the spec's field is an 8 px rectangle with a plain gold border and a gold chevron. Either state that fields are intentionally rectangular, or add the open glow so the two read as one family. As drawn, the closed field is the old native-looking box; the premium read comes only from the open list.
7. **Open list alignment.** In the chart-states board the Source list opens `bottom-start` and overshoots the card's right edge by about 23 px. Specify `bottom-end` when the field sits in the right slot of a head or bar.
8. **Option accessible name.** Descriptions inside options will be read as part of the name. Put the description in `aria-describedby` on the option.
9. **Reduced-motion and `popover="manual"`:** state the fallback for browsers without the Popover API (Safari before 17): the list stays `position: fixed` with the same layer rule, or the native select stays. One sentence.

Charts

10. **Empty state contradicts the spec (B7).** Spec: axes, gridlines and the brush stay. Board (chart-states-1440.png, Empty): no brush and no legend row. Head meta for the empty window is also undefined (the board shows `Brokk · 3:20 to 4:10`; B5.2 says `{name} · {n} DPS · 5 s average`). Define the empty meta and redraw the board with the brush.
11. **Loading frame height does not match ready.** Loading card is about 406 px tall; the ready Damage done card is about 424 px. That is the layout shift B11.2 forbids. Match the reserved heights and add the Log scale switch and the brush strip to the loading frame (the spec lists the brush box; the board omits it).
12. **Tooltip placement covers the live data right of the crosshair** (threat boards: it hides the lines for about 280 px). Acceptable on hover, but say it flips left when the cursor is in the right 55% of the plot, not only when it would pass the frame edge, so the lines you are reading toward stay visible.
13. **Phase label collides with a marker line** (damage-timeline board: `PHASE 1` is crossed by the phase-start marker line). Give the label a `--surface-card` backing or offset it 4 px.
14. **Touch scrub layout shift.** The time pill leads the legend row while scrubbing; the legend wraps to three rows on the board. Say the pill is absolutely positioned or reserved in the row so no chip moves.
15. **2000 px.** The chart stays at 1248 px with about 376 px of empty gutter each side. That follows the centred-column rule and is correct, but the plot is 260 px tall by about 1180 px wide (4.5:1). Consider 300 px from 1440 up, since the owner works at 2000.
16. **Source field on the Threat tab.** The mode bar's Source and the head's Enemy are two fields beside each other on 1440. Say what Source does on Threat (filter the lines, or nothing). If nothing, hide it there.
17. **Legend `Only` and chip focus order.** Add focus-visible styling for the `Only` button inside the chip (the spec names the control but not its state).

Consistency

18. **Tenets and site alignment.** Compact head, quiet card, `quiet-label` and the Card slots are used correctly (tenets 9, 10, 11, 13 satisfied on paper). The spec adds shared components (`TimeSeriesChart`, `ChartLegend`, `ChartTooltip`, `OverviewBrush`, popover core). Tenet 10 says a new shared component is its own design review; this spec is that review, so say so in the header.
19. **Tenet 12 (compare to the reference).** The Warcraft Logs and Details! citations are present; add one line each on why ours is as good or better (class-coloured isolate with `Only`, synced crosshair across tabs).

## What is right (keep)

- Top layer plus an overlay counter makes the owner's bug class impossible, not just fixed.
- The native select stays the source of truth and the picker on phones; no custom touch list to maintain.
- Brush, drag, presets and URL are one `setWindow`.
- Every state of the closed field and every copy string is specified; the call-site table (A10) is complete and testable by grep.
- uPlot over ECharts is correctly argued with the report island budget.
