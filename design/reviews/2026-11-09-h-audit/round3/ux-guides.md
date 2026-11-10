# UX review, guides family, round 3 (2026-11-10, ux-designer)

Captures: round3-captures/guides{,-you,-class,-spec}-{1440,390}.png. Baseline: round2/ux-guides.md. handoffs.md items noted, not failed.

## /guides: SHIP
- Compact band, nine plates, every spec in view at 1440 and 390. Blank icons are fixture artefacts.
- Polish (open): Mage and Warlock plates pad to Druid's height, ~50 px dead plate (ClassSpecPicker.astro, align-items: start).

## /guides signed in: SHIP
- Felsnap answer block above the plates. "Your guide:" prefix (f-content) and "Your spec" pill under Beast Mastery at 390 (f-selected) are handed off.

## /guides/hunter: SHIP
- Race icons bordered and faction-tinted; no role tag; specs, what changed, races, sources.
- Polish (1440): Races rail stops y 670, Sources runs on beside an empty right column ([class]/index.astro).

## /guides/hunter/survival: SHIP WITH FIXES
Fixed since round 2: Rotation, Opener and Cooldowns rows show the dark bordered placeholder (1440), no browser glyph.
- Must fix (1440 and 390): the talent trees still show the browser broken-image glyph on every lit node (390: 7 of 13 Survival nodes, 5 in Marksmanship) and a grey glyph on every dimmed node. Source: planner/TalentCell.svelte via ui/SpellIcon.svelte; the onerror fallback is not applied in the tree. It is the first thing under the band.
- Noted, handed off (Toc.svelte): sticky bar at 390 still "Survival Hunter · On this page" + "Change".
- Polish: "Enchant Weapon - Agility" appears twice (survival.md).
Handed off, unchanged: dead plate beside prose, 17-slot BiS repeat, icon-less extras rows, pasted racials, lede, "DPS · Updated", lane text links.

## Checklist
- [x] Compact bands, no big header, four pages
- [x] Race icon fallback
- [x] Rotation and cooldown icon fallback
- [x] No em dashes seen
- [ ] No browser glyph in talent trees (spec page)
- [ ] "Jump to" at 390 (handed off)
- [ ] Dead plate, extras icons, BiS strip (handed off)
