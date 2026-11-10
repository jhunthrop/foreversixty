# AI-slop review, sim pages, round 3 (2026-11-09)

Method: all nine pages at 1440 read in full. At 390 read landing-in, gear, talents, drops, weights, specs in full; result and saved 390 not reopened (12k px, same components as 1440; saved 1440 read). Handed-off items and fixture artefacts ignored. Weights haste caption kept per owner rule.

## First impression
Reads as a player's tool. No glow, gradient, glass, em dash or marketing verb. Round 2 fixes landed: band sub-line gone, "No sim yet" gone, Rage panel unit fixed, fidelity caption and uncast rotation steps gone, "Entered by hand" and "of 51" gone, saved-builds dev-speak gone, Talents now shows a per-talent diff, Crafted rows carry icons, Healers and tanks folded to one line, "What it does" disclosure gone.

## Verdicts
- sim-landing-out: PASS
- sim-landing-in: PASS WITH FIXES
- sim-result: PASS WITH FIXES
- sim-saved: PASS WITH FIXES
- sim-gear: PASS WITH FIXES
- sim-talents: PASS
- sim-drops: PASS WITH FIXES
- sim-weights: PASS WITH FIXES
- sim-specs: PASS WITH FIXES

## Remaining findings (not in handoffs.md)
sim-weights, 1440+390 (web/src/lib/sim/copy.ts ~1085, Weights page component)
- 7/22 "3,120 iterations" caption under the band. Move to the Run title or drop.
- 7/22 grey line "Per point of stat, normalized to Hit = 1.00, with DPS per point and the sim error." explains the column headers. Keep only the Haste sentence (owner rule).
- 22 "Level 60 raid set" and "Open Top Gear" are two competing links; keep one in the band.
- Stats have no icons; "Attack power" disabled ticked box is unexplained.
sim-specs, 1440+390
- 22 "Raid BiS DPS at 60" beside the Damage specs title while every row is a bare number. Cut it.
- Row carets with no label read as decoration; make them a real toggle affordance or remove.
sim-gear, 1440+390
- 22 "Equipped" under every candidate name and "Lock to equipped" twice. Mark the equipped row once.
- 22 labels that repeat the placeholder: "Name"/"Name", "Paste an export string" twice (Whole sets), "Name" (Find an item). Keep one.
- Find an item list still clips mid-row ("Lightbreaker Cuirass") at both widths.
- 390: ilvl "61" drops onto its own line left of "Copy and modify", reads detached.
sim-drops, 1440+390 (lib/sim/loot.ts)
- 7 grey "All professions" line under Crafted. Cut.
- Dungeons, Reputation, PvP are captured collapsed, so round-2 PvP rank rows and Bosses lines are unverified; recapture expanded.
sim-landing-in, 390
- Both full talent trees still sit above Rotation; page is 4.3k px. Collapse trees to the 17/32/2 split on phone (round-2 finding, not done).
- 22 "Item level 65" repeated on eight gear rows; "ilvl 65" or bare number.
sim-result / sim-saved, 1440 (SimRotationSection.svelte)
- 7 grey paragraph under 8 of 10 rotation steps. Concrete, but heavy; keep text for the non-obvious steps only.
- 22 "Level 60 raid BiS 1,004.0 DPS / 100 DPS under · 9.9% · ± 2.3" says the same thing three ways; keep "1,004.0 · 9.9% under".

## Authentic, keep
904.4 DPS hierarchy; ability table with miss/dodge/glance; execute-phase timeline; rage ledger with lost at cap; buff uptimes; Your sims; Try each with real consumables; talent diff list with +5/-5 per talent; weights table with bars and rating conversions; class-coloured crests; rank + DPS bars on specs; Crafted rows with professions, counts and ilvl.

## Handed off (noted, not failed)
Timeline Off/On toggle unlabeled (f-content). Sim gear stat strip and BiS DPS target (h-doll).
