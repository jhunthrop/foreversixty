# ux-designer review, sim pages, h-audit round 3 (2026-11-10)

Reviewed: round3-captures/{sim-landing-out, sim-landing-in, sim-result, sim-saved, sim-gear, sim-talents, sim-drops, sim-weights, sim-specs}-{1440,390}.png
Against: round2/ux-sim.md, BRIEF, handoffs. Handed-off (timeline Off/On label, SimGearSection stat strip) and fixture artefacts (candidates limited to Head/Main hand, talent "current build 0/20/0", empty saved builds, "0 valid combinations") not failed.

## Verdicts
| Page | R2 | R3 |
|---|---|---|
| sim-landing-out | SHIP | SHIP |
| sim-landing-in | SWF | SHIP |
| sim-result | SWF | SHIP WITH FIXES |
| sim-saved | SWF | SHIP WITH FIXES |
| sim-gear | SWF | SHIP |
| sim-talents | SWF | SHIP |
| sim-drops | NOT SHIPPABLE | SHIP WITH FIXES |
| sim-weights | SHIP | SHIP |
| sim-specs | SWF | SHIP |

## sim-landing-out: SHIP
- polish (390): Hunter and Warlock rows still wrap one spec to a second line (SimLanding.svelte / SpecGrid.svelte).

## sim-landing-in: SHIP
Fixed: Rotation step 1 now carries the icon box and matches steps 2-12 at both widths. 1440 columns balance.
- polish (390): 4,297 px; the Talents grid (about 700 px) sits before Rotation. Collapse Talents behind Disclosure on phone (BiS/planner owns talents).

## sim-result: SHIP WITH FIXES
Answer (904.4 DPS, error, vs BiS, Save this sim, Find upgrades) in the band at both widths.
- must fix (390): page is 11,769 px, every section still open (was 12,189). Round 2 asked for Disclosure on Rage, Buffs and debuffs, This run, Rotation; nothing collapsed. Keep Damage by ability, DPS spread, Gear open. SimView.svelte / LazySection.svelte.
- polish (1440+390): Rage "Refunded on a miss, dodge or parry" now has a bare glyph, not the framed icon box of its siblings. SimResourceSection.svelte.
- polish: Copy for the game / Copy JSON / Download JSON still two uneven rows.
- polish (390): "Up all fight" chips run 7 loose rows.

## sim-saved: SHIP WITH FIXES
Share band correct (Open in the sim, Copy link, Saved 4 Nov).
- must fix (390): 11,381 px, same collapse as sim-result. Same file.
- check: title "Fury Warrior" only; show the character name if the sim carries one.
- polish: Rage glyph, as sim-result.

## sim-gear: SHIP
Fixed: Add a build and Add set are quiet yellow text; Run is the only filled action; Copy and modify is a text link.
- polish (390): Candidate rows put the item level and "Copy and modify" on a second line with a 40 px dead gap between slots.
- polish (1440+390): Find an item list clips the last row mid-text ("Lightbreaker"); end on a whole row or add a fade.
- polish: Whole sets and Talent builds can sit in Disclosure.
- check: full-gear capture still owed (16 slots).

## sim-talents: SHIP
Fixed: the guide build lists its changes (+5 Unbridled Wrath ... -1 Piercing Howl) beside the grid at 1440; Add a build is quiet text; 390 stacks cleanly.
- polish: the diff is one ungrouped list of 24 lines; group by tree (Arms, Fury, Protection).

## sim-drops: SHIP WITH FIXES
Fixed: 1,223 px at 1440 and 1,716 px at 390 (was 5,570); Dungeons, Reputation and PvP collapse with "0 of n ticked"; Crafted rows carry framed profession icons; zero-count rows gone; one Show unreleased content.
- must fix (1440+390): the Raids/Dungeons/Crafted/Reputation/PvP chips are ticked while the page says "0 valid combinations" and every group reads "0 of n ticked". Either tick the groups to match or untick the chips. Droptimizer.svelte / SourcePicker.svelte.
- polish: Raids prose ("Onyxia's Lair: First raids, 9 December 2026" plus a grey run-on of eight unopened raids) is 12 px grey; show one line plus chips.
- polish: only Crafted is open, with no reason; open the first group with ticked items, or all closed.

## sim-weights: SHIP
- polish: "3,120 iterations" engine talk above the strip; drop or "Runs in your browser."
- polish: the Haste note is a grey 12 px line in the plate header; make it a table foot row.

## sim-specs: SHIP
Fixed: spec crest per row at 28 px, rank and DPS bar per row, "What it does" is now a quiet chevron, rows 44 px at both widths, healers and tanks on a plate with crests.
- polish (1440): about 700 px of empty row between the spec name and the rank/bar; start the bar nearer the name.
- polish (390): "Marksmanship Hunter" and "Enhancement Shaman" touch the rank; truncate or drop to 13 px.
- polish: list is sorted by class; sort by DPS rank.

## Checklist
- [x] No empty icon squares: landing-in step 1, drops, specs.
- [ ] Rage "Refunded" uses a bare glyph, not the framed box (polish).
- [x] No thin pages.
- [x] One crest language: spec crests on specs.
- [x] Answer without scrolling at 390 on every page.
- [ ] Phone collapse of secondary sections: result and saved not done (must fix).
- [x] One primary per region: top gear, talents.
- [ ] No engine meta talk: weights "3,120 iterations" (polish).
- [x] No em dashes.
- [x] Number takes its unit.
- [x] Class and quality colours on plates.
- [ ] Drops: ticked chips contradict "0 of n ticked" (must fix).
