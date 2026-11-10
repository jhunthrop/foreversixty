# ux-designer review, sim pages, h-audit round 5 (2026-11-10)

Reviewed round5-captures/{sim-landing-in, sim-result, sim-saved, sim-drops}-{1440,390}.png against round4/ux-sim.md, BRIEF, handoffs. Handed-off and fixture items not failed.

| Page | R4 | R5 |
|---|---|---|
| sim-landing-in | SWF | SHIP |
| sim-result | SWF | NOT SHIPPABLE |
| sim-saved | SHIP | SHIP |
| sim-drops | SWF | SHIP WITH FIXES |

## sim-landing-in: SHIP
Fixed: Your sims sits under Talents; left column ends y about 1605, right (Rotation) y about 1670. Columns balanced, no dead background. 390 is 3,825 px (was 4,021).
- polish: "Item level 65" on gear rows (handed to h-doll).

## sim-result: NOT SHIPPABLE
- blocker (390): the capture is 871 px wide, so the page overflows horizontally. The Timeline grid (Whole fight) is not constrained to the 390 viewport and drags the page wide. sim-saved's timeline fits, so the fix is in how SimView/SimTimeline is wrapped here (overflow-x:auto inside its plate, min-width:0 on the grid parent).
- must fix (390), fourth round: still 9,206 px at 871 wide. Rage, Buffs and debuffs, This run, Your sims and Rotation are all open. sim-saved collapses Rage, Buffs and Rotation with the same component. Keep Damage by ability, DPS spread and Gear open; collapse the rest. SimView.svelte, same props sim-saved passes.
- polish: Copy for the game / Copy JSON / Download JSON still two uneven rows; "Up all fight" chips run loose rows.
- 1440: answer in the band, no deviations.

## sim-saved: SHIP
390 is 6,731 px, fits 390 wide, answer (Fury Warrior, 904.4 DPS, Open in the sim, Copy link) in the first screen.
- polish (390): This run (about 900 px) still open; collapse it.

## sim-drops: SHIP WITH FIXES
Length fine (1,167 px at 1440, 1,688 at 390).
- must fix (1440+390), second round: Raids is a ticked chip but renders as one prose line ("Onyxia's Lair: First raids, 9 December 2026 / No date yet: ...") with no sources to pick. Render it as a group like Crafted (rows with icon, items, ilvl; undated raids as disabled rows reading "No date yet"). Droptimizer.svelte / SourcePicker.svelte.
- polish: only Crafted is open with no reason; open the first group that has items or close all.

## Checklist
- [x] No empty icon squares, no thin pages, one crest language, no em dashes.
- [x] Landing-in 1440 column balance.
- [x] Answer visible at 390 on all four.
- [ ] Result 390 fits the viewport (blocker).
- [ ] Result 390 secondary sections collapsed (must fix).
- [ ] Drops Raids group has sources to pick (must fix).
