# Round 11 UX review (1440 and 390)
Handed-off items (handoffs.md) excluded. Report-deaths 390 keeps Sim this fight and Copy link in the 44 px "···" Report actions menu by design. Coach has no count because no loaded payload carries one.

## guides-spec: SHIP
- 390 fell from about 11,280 px to 8,089 px. Best in slot ("Show 17 items"), Enchants and consumables ("Show 15 enchants and consumables"), Races ("Show 8 races"), Professions and Leveling are folded with counts or labels. Answer (Build plate, Load this build, Sim this build) is above the fold at both widths.
- 1440: Overview two columns, Races rail beside prose, no empty plate.
- Polish: 390 is 89 px over the 8,000 target; fold Rotation cooldown prose or trim the Overview "What the raid adds" block.
- Polish, 1440: Best in slot paperdoll plate has about 120 px dead band either side; Leveling "Talents by level" plate floats right of the prose with dead space under it.

## character-owner-60-premium: SHIP
- Round 10 must-fix landed enough: History now reads "7 fights" with a chevron; Coach is a chevron row. Coach count accepted as data-blocked.
- Pass: Upgrade banner, Standing, Activity carry the next step; Lucifron delta and sparkline sit on a second line at 390.
- Polish, 1440 and 390: History and Coach are still content-width links, not full-width divided rows; give them the row treatment so the tap target is 44 px.
- Polish, 390: Upgrade banner has no chevron and no "Upgrade:" label (1440 has the label); make both match and make the whole banner a link.
- Polish, 1440: Lucifron "+82 vs 27 Oct" and the sparkline still float mid-row.

## report-deaths: SHIP
- Pass at 390: "···" actions menu beside "5 deaths" holds Sim this fight and Copy link (by design); first death, last hits and heals, Killing blow tag, class-coloured names above the fold. 1440 shows Sim this fight (primary) and Copy link in This pull.
- Polish, 390: tab strip still clips "Damage ta" with no fade or chevron cue.
- Polish, 390: chart legend is six stacked lines (about 180 px) before the chart; wrap two per row.
- Polish, 1440: rail holds only This pull; about 880 px empty beside the chart.

## sim-talents: SHIP WITH FIXES
- Must fix, 390: each tree renders at the left edge about 125 px wide with about 165 px dead to the right, three trees stacked per build; page is 3,390 px. Show the three trees across the row or centre them.
- Must fix, 1440: the Your current build row has an empty right column where Guide build shows Gains and Drops; show "No changes" or the delta against the guide build.
- Must fix: the two compare checkboxes have no label for what they select and Run is disabled with no reason; add "Pick builds to compare" beside Run at both widths.
- Pass: compact title band, tab strip, Gains green and Drops red with signed numbers, no em dashes.
- Polish: "No saved builds" and a bare yellow "Add a build"; give it the dashed-plate empty state.

## sim-specs: SHIP
- 1440: ranked rows with spec icon, class-coloured name, rank, bar and DPS; reads as the Details! rank list. Healers and tanks wrap onto two rows.
- 390: bar column dropped, names wrap on five specs, number column holds.
- Polish, 390: the bar is the scan cue and it is gone; keep a 3 px bar under each name.
- Polish: healers and tanks are a bare list with no number and no link; link each to its guide.

## Checklist
- [x] guides-spec 390 folds with counts
- [ ] guides-spec 390 under 8,000 px (8,089)
- [x] report-deaths 390 actions present via menu
- [x] character History/Coach chevron rows (not full-width)
- [ ] sim-talents 390 tree dead space
- [ ] sim-talents Run reason and checkbox label
- [x] no em dashes on any of the five pages
