# ux-designer gate: lane G2 (home foot text, /rankings snapshot, PageHeader lead order)

Judged on the shots in shots/ plus RankingsLanding.svelte, Rankings.svelte, HomeAbout.svelte. Reference for the boards: the Warcraft Logs rankings table and the Details! window (name in class colour, bar, number). Reference for the home foot: the Wowhead class hub footer (text plus link chips).

## Must fix

1. **"2 ranked" contradicts the three rows above it** (rankings landing, all three encounters, 1440 and 390). Each card lists three characters and says "2 ranked". Either the count is something else (guilds? kills?) and needs a label, or it is wrong. The Saltspine board says "10 ranked" for 10 rows, so the landing count must mean the same thing: total ranked characters on that encounter, never fewer than the rows shown. Verify the number against the data, then fix the label or the figure.
2. **Zero-DPS rows are ranked** (Saltspine board, rows 6 to 10: Arkadia Starwoven, Glaucus Strom, Spork Xo, Turbo Turkey, Woodz Holditdowntw, all "0" DPS with an empty bar and a percentile of 50 to 10). A crawler and a player both read "0 DPS, 50th percentile" as a bug. Drop rows with 0 DPS from the board and the "ranked" count, or show them with a real value. Owner of the fix is the snapshot builder (filter at build), not the view.
3. **"Leaders by spec" renders as an empty panel on the landing** (1440 and 390, JS on and off). A heading plus "No ranked kill on this board yet." is thin text on the page we want ranked, and it reads as broken. Until live rows carry a spec, do not render the panel. When it returns, it goes in the empty right third at 1440 (the encounters card is 880 px wide and leaves a 370 px void), not under the card.

## Should fix

4. **Executed column is "—" on every row** (Saltspine, 1440). An all-empty column is noise. Hide the column when every value is empty.
5. **"Updated 11 Nov" is missing from the no-JS rankings header** (rankings-1440-nojs: "Tier list ›" only; with JS it reads "Tier list ›  Updated 11 Nov"). The stamp is the freshness signal a crawler should read. Render it at build time. This also removes a hydration-time text change in the meta line.
6. **Home "What's here" last line is body-weight.** "Data from client build 1.60.1.70790 ... ranked 10 Nov." is provenance. Set it at 13 px, text-muted, so it is visibly secondary to the three paragraphs above it.
7. **Home foot columns end 120 px apart at 1440** (What's here ends y 1397, Leveling BiS by spec ends y 1517). Either let What's here stretch to the row height or set the grid to align-items: start and accept it; a half-filled panel next to a full one is the current look and it reads unfinished.
8. **Leveling BiS chips are about 26 px tall on 390** with a ~1 px row gap. Make each chip a 44 px tap target (padding, not margin) on coarse pointers. Keep the class colour on the name; the Tanks / Healers / Damage labels are correct and stay.
9. **Phone board rows: percentile and date sit on one line** ("100  9 Nov", rankings-saltspine-390). They read as one number. Put the date in the sub-line under the class ("Shaman · 9 Nov") and leave the percentile alone under DPS. The landing mobile cards also drop the Percentile column header and the legend; add the "Percentile" legend row under the Encounters heading as the Saltspine board has.
10. **PageHeader lead is 4 lines on a 390 px tool page** (sim-weights-390: lead ends y ~210, tabs at ~250, the weights table beyond 1400). The lead-under-title order is right and matches planner-1440 (title, lead, meta "Human", 12 px gaps, no shift). On sim and planner, clamp the lead to 3 lines at 390 so the tabs and the first panel start higher.

## Checked and fine

- Home: the class picker stays first (390: y 120 to 740); the text panels are at the foot, as briefed. No hero, no pitch. Links in the prose are the site yellow with underline-free hover as elsewhere. Server text is identical with JS off apart from the feed skeleton, which holds the same height (no shift).
- Rankings landing 1440 and 390: names in class colour, crest, bar, DPS, percentile in the percentile colours, same as the board. Cards at 390 have no horizontal scroll.
- Saltspine board 1440 and 390: no-JS and JS shots match.
- PageHeader site-wide: planner and sim-weights both show the lead directly under the title and the meta line under it.

VERDICT: SHIP WITH FIXES (must fix 1, 2, 3 before merge; 4 to 10 are shoulds)

## Round 2

Judged on the round-2 shots (home 1440/390, rankings 1440 with and without JS, Saltspine 1440 with and without JS).

### Must fix

1. **"2 kills" still sits above three characters** (rankings landing, all three encounters, 1440, JS on and off). The label changed, the contradiction did not. Three rows are three ranked kills at minimum, and a reader counts rows. Fix the figure (count of ranked rows on that encounter, not fewer than shown, from the same trimmed data as the rows) or show only as many rows as the count. Verify against the snapshot.
2. **Saltspine no-JS says "10 ranked" over 5 rows** (rankings-saltspine-1440-nojs: five rows, Buzz Killa last at 27 DPS; heading "10 ranked"). The trim removed the rows but not the count. Compute the count from the trimmed rows.
3. **No-JS and JS Saltspine boards differ.** The JS shot still shows rows 6 to 10 at 0 DPS with percentiles 50 to 10; the no-JS shot ends at row 5. The table grows by 5 rows on hydration (layout shift) and a player still sees the 0 DPS rows I flagged in round 1. Deferring the live-board fix to the API is acceptable for the API. It is not acceptable that the view and the snapshot disagree: filter 0-DPS rows in the board view too (one line, same rule as the snapshot) so both render 5 rows.

### Fixed since round 1 (verified)

- Leaders by spec panel is gone from the landing (1440).
- Home foot: two panels end together at y ~1441 (1440); build line is muted 13 px and reads as provenance on 1440 and 390.
- BiS by spec chips on 390 have visible row spacing and wrap cleanly.
- Landing label reads "N kills" with the chevron (the label is right, the number is item 1).

### Accepted as not done

- Executed column all dashes, the empty right rail, phone percentile/date line: pre-existing board layout, out of lane. Log them as follow-ups; the all-dash column is the first one to take (hide when every value is empty).
- "Updated" stamp: the no-JS rankings header is still missing it (JS shot shows "Updated 11 Nov"), so a text change lands in the meta line on hydration. Your locale reason holds for a formatted date; render the ISO date server-side in a fixed format ("11 Nov", en-GB day-month) and let the client keep the same format. Should fix, not blocking.
- Lead clamp on phone: accepted as a PageHeader-wide change; log it.

VERDICT: SHIP WITH FIXES (must fix 1, 2, 3: kill count equals or exceeds the rows shown, ranked count from trimmed rows, board view filters 0-DPS so no-JS and JS match)

## Round 3

Judged on rankings-1440, rankings-saltspine-1440 and rankings-saltspine-1440-nojs.

- Must fix 2 and 3 closed. Saltspine reads "5 ranked" over 5 rows with and without JS, the rows match one to one (Psycho Giga 114 down to Buzz Killa 27), no 0-DPS rows, no height change on hydration. Percentiles re-ranked on the five rows (100, 80, 60, 40, 20) and the landing cards agree with the board for the same characters.
- Must fix 1 closed on the explanation. "2 kills" counts kills, a 5-player kill puts up to five characters on the board, and the label says kills, not characters. The Saltspine board (5 ranked from 2 kills) is consistent with that. No change needed.
- Still open, accepted as follow-ups (not blocking): the no-JS header lacks "Updated 11 Nov" (text appears on hydration in the meta line), and the Executed column is all dashes on every row.

VERDICT: SHIP (follow-ups: server-render the Updated stamp in a fixed format; hide Executed when every value is empty)
