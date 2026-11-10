# AI-slop review, sim pages, round 1 (2026-11-09)

Method: all nine pages at 1440 read in full. At 390 only sim-landing-out and sim-gear were opened; the other seven 390 captures were NOT reviewed.
Framed inset panels are sanctioned by docs/design-system.md (Plate frame, two or three per page), so not counted as tell 8. No glows, gradients, glass, em dashes or marketing verbs found on any page.

## First impression
Result and Saved read as made by a player: real ability rows, timeline, rage ledger, buff uptime. The tool pages (Top Gear, Talents, Droptimizer, Specs) read as form dumps in a nice shell, with explainer copy a dev would not write.

## Verdicts
- sim-landing-out: PASS WITH FIXES
- sim-landing-in: PASS WITH FIXES
- sim-result: PASS WITH FIXES
- sim-saved: PASS
- sim-gear: FAIL (tell 13, thin empty page)
- sim-talents: FAIL (tell 13, thin empty page)
- sim-drops: FAIL (thin generic checkbox wall, no icons, no level ranges)
- sim-weights: PASS WITH FIXES
- sim-specs: PASS WITH FIXES

## Findings
Common (all tool pages): band sub-line "The simulator runs damage specs only. Healing and tanking specs are not simulated yet." (tell 7/21; veterans know). Cut from every tab; Spec support already says it.

sim-landing-out
- 7 band sub-line "Damage specs only. Nothing is uploaded until you save." 1440/390, web/src/lib/sim/page-copy.ts:8. Cut; keep the "Get set up" hint under the paste box.
sim-landing-in
- 22/21 band right "No sim yet / Press Run sim. It runs here, in your browser." 1440, web/src/lib/sim/page-copy.ts:65-66. Cut the caption; the Run sim button is the instruction.
- Layout: Gear+Talents column ends y~1330 while Rotation runs to 1890, ~900px of bare ground under Talents (1440). Move Your sims under Talents or drop the 12-step Rotation card pre-run (restates the guide; one job per page), web/src/components/sim/SimRotationSection.svelte.
- 22 gear sub-lines mix "Item level 65" and "Strength +8" per piece; pick one.
sim-result
- 22 Rage panel "Gained 2,127 rage / Spent 2,116 rage / Lost at the cap 57 rage" web/src/components/sim/SimResourceSection.svelte. Panel title carries the unit; show 2,127 / 2,116 / 57 (2.7%).
- 22 "Recklessness counts once a pull here. In a raid it is back every 30 minutes." under the run bar. Keep only if it changes the number (SimFidelityCaption.svelte).
- 22 "Missed in this run: 0% of specials · 13.9% of white swings." "in this run" is implied. Write "Miss 0% specials, 13.9% white".
- 8 Rotation card keeps greyed "Not cast in this run" steps (5, 11); drop uncast steps. SimRotationSection.svelte.
- Credit: 904.4 DPS is the clear top of the hierarchy.
sim-saved
- none blocking. Same Rotation note as result.
sim-gear
- 13 Candidates panel with two rows plus "Nothing to try yet. Tick an item..." above it; left column is empty from y=750 to 2280 beside a tall rail (1440). Empty state pretends to be a workspace. Pre-fill candidates with best-in-slot upgrades or put Find an item in the left column. lib/sim/copy.ts:919.
- 7/21 "Three stages: everything at 100 iterations, the survivors at 1,000, the finalists at 3,000." copy.ts:925 beside the empty line. Move to the Fast select tooltip.
- 7 "Try each of these as a candidate rather than a setting." (copy.ts:879) and "A set replaces every slot at once. Paste a second export string to add one." Cut both.
- 22 "20 talent points. Simmed as a level 60 with these talents." (lib/sim/handoff-copy.ts:26) plus "Entered by hand" (lib/sim/sources.ts:145). One line: "Orc Fury Warrior 60 · ilvl 62.0".
- 22 "Your saved builds could not be read; the rest of the page still works." (copy.ts:906) shown by default in Talent builds, 1440/390. Dev-speak error as content; show "No saved builds". Reword even if fixture-only.
- Find an item list clipped mid-row ("Lightbreaker Cuirass") at both widths; show the scroll or paginate.
sim-talents
- 13 One panel "Talent builds" with two unticked checkboxes and the same error line; page ~1000px tall with ~150px of content. Render the two builds as trees with the point split (17/32/2) and the talent diff so there is something to pick.
- 7/21 "Nothing to try yet" and "Three stages" strip repeats on a page with two checkboxes. Cut.
- 22 same saved-builds error line (copy.ts:906).
sim-drops
- Thin and generic: 7331px at 1440 of bare checkboxes in flat lists; no icons or crests, no level ranges on dungeons (tell 19). web/src/components/sim/tools (Droptimizer), lib/sim/loot.ts.
- 19 "Every boss 320" at the far right of each row unlabelled; "Bosses 33" disclosure and "Trash and chests" repeated under 21 dungeons. Show "320 items" once.
- Reputation (40 rows) and PvP (45 rows "Rank 10 (Alliance)", counts 1-9) are one flat list each. Group reputation by faction, PvP by rank with both factions on one row.
- 22 Raids block: seven grey lines "Not open yet; no date announced." Collapse to one line with Show unreleased.
- "No source ticked yet" sits below the list, off screen; pin it by the Run bar.
sim-weights
- 7/21 "Per point of stat, normalized to Hit = 1.00, with DPS per point and the sim error. Haste: 9.25 per 1%, not in the table because no item at this level range has it". Column heads say it; make Haste a table row.
- 21 "3,120 engine iterations in your browser: a baseline pass plus a low and a high pass for each stat weighed." (copy.ts:1085). Move to the Run button title.
- 22 "Reference: Attack power" under Stats to weigh; the disabled ticked box says it.
- Stats have no icons; right column empty under Stats to weigh (1440).
- Credit: Scale / Per point / Error table in SimC convention is specific.
sim-specs
- 17 all 18 damage rows carry a "Not yet" chip (1440), copy.ts:328, lib/sim/spec-state.ts. Status with no state behind it if live; if fixture-only ignore, else show the real state or no chip.
- Rows are 65px tall with only a name and a disclosure; add the spec crest (one crest language) and tighten to ~40px. components/sim/SpecCard.svelte, SpecGrid.svelte.
- "What it does" repeated 18 times is noise; one link per row for a veteran is too many.

## Authentic, keep
Timeline with Execute phase band, rage ledger with lost-at-cap, buff uptimes with application counts, DPS spread with mean/error/std dev, engine sha in This run, Your sims rows (900 DPS, Fury · 3 Nov), SimC-convention weights, class colours on plates, quality-coloured items with enchant lines, Droptimizer's real zone names.
