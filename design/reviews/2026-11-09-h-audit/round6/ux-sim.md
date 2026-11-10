# ux-designer review, sim pages, h-audit round 6 (2026-11-10)

Reviewed round6-captures/{sim-landing-in, sim-result, sim-saved, sim-drops}-{1440,390}.png against round5/ux-sim.md, BRIEF, handoffs. Handed-off and fixture items not failed.

| Page | R5 | R6 |
|---|---|---|
| sim-landing-in | SHIP | SHIP |
| sim-result | NOT SHIPPABLE | SHIP WITH FIXES |
| sim-saved | SHIP | SHIP |
| sim-drops | SHIP WITH FIXES | SHIP WITH FIXES |

## sim-landing-in: SHIP
1440: Rotation ends y about 1635, Gear/Talents/Your sims stack ends about 1605; columns balanced. 390 is 3,789 px, Run sim in the first screen, Talents folded to one row (Arms 17 / Fury 32 / Protection 2).
- polish (390): about 50 px of empty band between the character line and the tabs.
- polish: "Item level 65" on gear rows (handed to h-doll).

## sim-result: SHIP WITH FIXES
Fixed from R5: 390 is now 390 px wide (no horizontal overflow), timeline fits, Rage, Buffs and debuffs, Rotation and Talents are folded. 7,555 px (was 9,206 at 871 wide). 904.4 DPS, Save this sim and Find upgrades are in the first screen. 1440 answer in the band.
- must fix (390): the settings form (Fight, Length, Targets, Buffs, Precision, More settings) is fully open under the tabs, about 600 px before Damage by ability. The landing page shows the one-line "Patchwerk 3:00 · Raid / Change" bar. Use the same bar here with "Run again". SimView / the run-bar wrapper.
- polish (390): This run (about 900 px) is still open. Handed off to sim-phone.spec (openSimSection), so not failed; fold it when the spec is updated.
- polish (1440): Rage and Buffs and debuffs sit in the 460 px left column, leaving about 1,100 px of bare background to their right once Rotation ends. Let them span the full width or move Rotation below them.
- polish: Copy for the game / Copy JSON / Download JSON uneven rows at 390 (one full, two half).

## sim-saved: SHIP
390 is 6,731 px, fits 390, answer (Fury Warrior, 904.4 DPS, Open in the sim, Copy link) in the first screen; Rage, Buffs and Rotation folded.
- polish (390): This run still open, about 900 px; fold it.
- polish (1440): same dead column right of Rage and Buffs (about 1,300 px) as on sim-result.

## sim-drops: SHIP WITH FIXES
1,210 px at 1440, 1,848 at 390. Raids as one prose line is accepted for the fixture (no raid opens before 9 December); released raids render as rows.
- must fix (1440+390): the Crafted footnote prints duplicates: "No date yet: Blacksmithing, Blacksmithing, Blacksmithing, Blacksmithing, Engineering, Leatherworking, Leatherworking, Tailoring". Dedupe to profession names, or drop the line for professions already shown as rows. At 390 it is also a 6-line grey prose block (Droptimizer.svelte / SourcePicker.svelte).
- polish: the Raids line should scan as chips ("Onyxia's Lair 9 Dec", "No date yet") rather than a sentence.
- polish: Crafted is the only group open with nothing picked.

## Checklist
- [x] No empty icon squares, no thin pages, one crest language, no em dashes.
- [x] Landing-in 1440 column balance, 390 length.
- [x] Answer visible at 390 on all four.
- [x] Result 390 fits the viewport.
- [x] Result 390 Rage, Buffs, Rotation folded.
- [ ] Result 390 settings folded to the one-line bar (must fix).
- [ ] Drops duplicate-name footnote (must fix).
- [~] Drops Raids group: accepted for the fixture.
