# Round 12 (final) UX review, 1440 and 390
Handed-off items (handoffs.md) excluded.

## sim-landing-out: SHIP
- Answer (Sim your character, input plus Run) is the first plate at both widths. 1440 best-in-slot set is a 3-column crest grid with class-coloured names. Other ways in is a labelled row list.
- Polish, 390: spec links stack one per line (Hunter and Rogue about 130 px each) with the crest floating mid-block; wrap them inline as at 1440. Page is 2,448 px.

## sim-talents: SHIP
- Round 11 must-fixes landed: "Pick builds to compare." sits beside Run at both widths; Your current build now shows Gains and Drops against the guide build at 1440 and 390; trees fill the row at 390 (about 270 px wide, no left-edge strand).
- Polish: "No saved builds" and the bare yellow "Add a build" still lack the dashed-plate empty state. 1440 current-build left column is dead under the checkbox. 390 is 4,642 px (six stacked trees plus drop lists); acceptable on a compare page.

## sim-drops: SHIP
- Reads as a source picker: Show checkboxes, collapsed Raids/Dungeons with counts, Crafted open with profession icons, item count and ilvl. Run disabled with "Tick a source." beside it. Icons render, no empty squares, no em dashes.
- Polish: Raids "0 of 7 open" with nothing to tick on a fresh state; say which raids open when.

## sim-specs: SHIP WITH FIXES
- 1440: rank list with spec icon, class-coloured name, rank, bar, DPS; reads as the Details! list. Pass.
- Must fix, 390: the Healers and tanks block runs two columns inside the plate; names wrap to two lines and "Healer" for Restoration Druid touches the plate edge (about x 369 against a 357 inner edge). Fix: one column at 390 (name left, role right), 44 px rows.
- Polish, 390: DPS bar still dropped; add a 3 px bar under each name. Healers and tanks still unlinked.

## guides-spec: SHIP
- Build plate with Load this build and Sim this build above the fold at both widths; "Jump to" label landed; Best in slot, Enchants and consumables, Races, Professions, Leveling fold with counts.
- 390 is 8,061 px, 61 px over the 8,000 target (was 8,089). Polish: fold the Rotation cooldown prose.
- Polish, 1440: Best in slot paperdoll plate has dead band either side; Leveling Talents by level plate floats beside prose.

## Checklist
- [x] sim-talents Run reason (Pick builds to compare.)
- [x] sim-talents 390 trees fill the row
- [x] sim-talents current build Gains/Drops
- [x] guides-spec "Jump to"
- [ ] guides-spec 390 under 8,000 px (8,061)
- [ ] sim-specs 390 healers/tanks inside plate padding
- [x] no em dashes on the five pages
