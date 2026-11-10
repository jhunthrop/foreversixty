# ux-designer review, sim pages, h-audit round 4 (2026-11-10)

Reviewed: round4-captures/{sim-landing-out, sim-landing-in, sim-result, sim-saved, sim-gear, sim-talents, sim-drops, sim-weights, sim-specs}-{1440,390}.png against round3/ux-sim.md, BRIEF, handoffs. Handed-off (timeline Off/On label, SimGearSection stat strip, full-gear capture) and fixture artefacts (0 valid combinations, 0/20/0 current build, empty saved builds) not failed.

## Verdicts
| Page | R3 | R4 |
|---|---|---|
| sim-landing-out | SHIP | SHIP |
| sim-landing-in | SHIP | SHIP WITH FIXES |
| sim-result | SWF | SHIP WITH FIXES |
| sim-saved | SWF | SHIP |
| sim-gear | SHIP | SHIP |
| sim-talents | SHIP | SHIP |
| sim-drops | SWF | SHIP WITH FIXES |
| sim-weights | SHIP | SHIP |
| sim-specs | SHIP | SHIP |

## sim-landing-out: SHIP
- polish (390): Warlock (Destruction) and Hunter (Survival) still wrap to line 2 (SpecGrid.svelte). 1440 three-column grid is clean.

## sim-landing-in: SHIP WITH FIXES
Fixed: Talents collapsed to a one-line bar at 390 (4,021 px, was 4,297); Rotation step 1 matches steps 2-12; Your sims added.
- must fix (1440): left column ends at Talents (y about 1330) while Rotation and Your sims run to y about 2150. About 800 px of bare background under Gear and Talents. Move Your sims under Talents or above Rotation so the columns end within about 150 px. SimLanding.svelte layout.
- polish: "Item level 65" repeated on gear rows (handed off to h-doll).

## sim-result: SHIP WITH FIXES
Answer (904.4 DPS, error, vs BiS, Save this sim, Find upgrades) in the band at both widths. Rage glyph is now framed.
- must fix (390), third round: 10,515 px. Rage, Buffs and debuffs, This run, Your sims and Rotation are all still open; sim-saved collapses them, so the component supports it. Keep Damage by ability, DPS spread and Gear open; collapse the rest on phone. SimView.svelte, same props sim-saved passes.
- polish: Copy for the game / Copy JSON / Download JSON still two uneven rows.
- polish (390): "Up all fight" chips run 7 loose rows.

## sim-saved: SHIP
Fixed: Rage, Buffs and debuffs and Rotation collapse at 390 (6,731 px, was 11,381); Rage glyph framed; share band correct.
- polish (390): This run (about 900 px) is still open; collapse it too.
- check: title "Fury Warrior" only; show the character name if the sim carries one.

## sim-gear: SHIP
Fixed: Find an item ends on a whole row with "8 more" and a count line at both widths; Add a build is quiet text; Run is the only filled action.
- polish (390): candidate rows wrap "Lionheart Helm" to two lines beside the item level and "Copy and modify"; give the name the width.
- polish: Whole sets and Talent builds can sit in Disclosure.

## sim-talents: SHIP
- polish: the guide-build diff is still one ungrouped list of 24 lines; group by tree (Arms, Fury, Protection).

## sim-drops: SHIP WITH FIXES
Fixed: the "0 of n ticked" counts are gone, so the Show chips now read as a visibility filter and no longer contradict the page. 1,203 px at 1440, 1,696 px at 390.
- must fix (1440+390): Raids is a ticked chip but its group is only prose ("Onyxia's Lair: First raids, 9 December 2026" plus a grey 12 px run-on of unopened raids) with no sources to tick. Render Raids as a group like Crafted (rows with icon and ilvl). Droptimizer.svelte / SourcePicker.svelte.
- polish: only Crafted is open with no reason; open the first group with items, or close all.

## sim-weights: SHIP
Fixed: "3,120 iterations" is gone.
- polish: the Haste note is still a grey 12 px line in the plate header; make it a table foot row.

## sim-specs: SHIP
Fixed: long spec names wrap cleanly with room before the rank at 390; the DPS bar sits beside rank and DPS at 1440; healers and tanks fill a wrapped row at 1440.
- polish (1440): about 650 px of empty row between spec name and rank; widen nothing, start the bar after the name.
- polish: sorted by class; sort by DPS rank.

## Checklist
- [x] No empty icon squares.
- [x] No thin pages.
- [x] One crest language.
- [x] Answer without scrolling at 390 on every page.
- [ ] Phone collapse of secondary sections: sim-result still open (must fix); sim-saved done.
- [ ] Landing-in 1440 column balance (must fix).
- [ ] Drops Raids group has nothing to pick (must fix).
- [x] One primary per region.
- [x] No engine meta talk.
- [x] No em dashes.
- [x] Number takes its unit.
- [x] Class and quality colours on plates.
- [x] Rage glyph framed.
