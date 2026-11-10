# UX review, guides family, round 4 (2026-11-10, ux-designer)

Captures: round4-captures/guides{,-you,-class,-spec}-{1440,390}.png. Baseline: round3/ux-guides.md. handoffs.md items (TalentCell/SpellIcon glyph: f-content; Toc "Jump to": f-selected) are noted, not failed.

## /guides: SHIP
- Compact band, nine plates, every spec in view at 1440 and 390. Blank icons are fixture artefacts.
- Polish (open): Mage and Warlock plates pad to Druid's height, ~50 px dead plate (ClassSpecPicker.astro, align-items: start).

## /guides signed in: SHIP
- Felsnap answer block sits above the plates at both widths.
- Handed off, unchanged: "Your guide:" prefix (f-content); "Your spec" pill wraps under Beast Mastery at 390 (GuideSpecRow.astro, f-selected).

## /guides/hunter: SHIP
- Three spec plates, What changed, Races (bordered, faction-tinted), Sources; no em dashes.
- Polish (1440): left column is empty from y 700 to 1160 beside the Races/Sources rail ([class]/index.astro).
- Polish (390): "new" tag floats far right of wrapped "High Order Skyborne" and "Windshaper Skyborne" in the 2-col race grid.

## /guides/hunter/survival: SHIP (handed-off items open)
- 1440: build plate (trees, Rotation, Stat priority, Load/Sim) fills the first screen under the band; lit talent nodes render as bordered icons.
- 390: build plate and Rotation (dark bordered placeholders, no glyph) are above the fold.
- Handed off, still visible at 390: browser broken-image glyph on lit talent nodes (TalentCell via SpellIcon, f-content); sticky bar reads "Survival Hunter · On this page" + "Change" (Toc.svelte, f-selected).
- Handed off, unchanged: lede "Survival is a melee spec in Forever", "DPS · Updated", dead plate beside prose, 17-slot BiS repeat, icon-less extras rows, pasted racials, duplicate "Enchant Weapon - Agility".

## Checklist
- [x] Compact bands, no big header, four pages
- [x] Race, rotation and cooldown icon fallback
- [x] No em dashes seen
- [x] Answer visible without scrolling at 390 and 1440 on all four pages
- [ ] No browser glyph in talent trees at 390 (handed off)
- [ ] "Jump to" at 390 (handed off)
- [ ] Dead plate, extras icons, BiS strip (handed off)
