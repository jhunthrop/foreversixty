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
