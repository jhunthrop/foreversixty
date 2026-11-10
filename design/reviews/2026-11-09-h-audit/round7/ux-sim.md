# ux-designer review, sim pages, h-audit round 7 (final, 2026-11-10)

Reviewed round7-captures/{sim-landing-in, sim-result, sim-saved, sim-drops}-{1440,390}.png against round6/ux-sim.md, BRIEF, handoffs. Handed-off and fixture items not failed.

| Page | R6 | R7 |
|---|---|---|
| sim-landing-in | SHIP | SHIP |
| sim-result | SHIP WITH FIXES | SHIP |
| sim-saved | SHIP | SHIP |
| sim-drops | SHIP WITH FIXES | SHIP |

## sim-landing-in: SHIP
1440 columns balanced (Rotation and Your sims end together), Run sim in the band. 390 is 3,673 px, "Patchwerk 3:00 - Raid / Change / Run sim" bar in the first screen, Talents folded to one row.
- polish: "Item level 65" on gear rows (handed to h-doll).

## sim-result: SHIP
R6 must-fix closed: at 390 the settings form is the one-line "Patchwerk 3:00 - Raid / Change / Run again" bar; 904.4 DPS, Save this sim and Find upgrades sit above it. 390 is 7,099 px (was 7,555). At 1440 Rage and Buffs and debuffs now span the full width; the dead column is gone.
- polish (390): This run still open (handed to sim-phone.spec).
- polish (390): Copy for the game / Copy JSON / Download JSON uneven rows.

## sim-saved: SHIP
Answer (Fury Warrior, 904.4 DPS, Open in the sim, Copy link) in the first screen at both widths; 1440 dead column fixed the same way. 390 is 6,711 px.
- polish (390): This run still open.

## sim-drops: SHIP
R6 must-fix closed: the Crafted footnote now lists each profession once. 1,191 px at 1440, 1,770 at 390.
- polish: the footnote opens "First raids, 9 December 2026:" under Crafted, which reads as a raid line for professions. Drop the prefix or say "Opens 9 December 2026:". Droptimizer.svelte / SourcePicker.svelte.
- polish: Raids is still one prose line (accepted for the fixture).

## Checklist
- [x] No empty icon squares, thin pages, em dashes; one crest language.
- [x] Answer in the first screen at 390 on all four.
- [x] Result 390 settings folded to the one-line bar.
- [x] Result/Saved 1440 no dead column.
- [x] Drops footnote deduped.
