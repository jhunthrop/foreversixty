# info-density-reviewer, round 3 (Me with no site activity)

Verdict: PASS

I checked `me-empty-l25.png` and `me-empty-l60.png` in the round 3 render. The three fixes landed and the page is dense with real local data.

**Must-fix**
None.

**Confirmed fixed**
- **Enchant rows:** they now read `Chest: Minor Stats` and `Wrist: Minor Agility`, with the slot first and nothing cut off.
- **Level 60:** Gear joins the shorter column under Where you stand, so the left column is filled and the right column holds only Right now and Your sim. There is no dead block now.
- **Upgrade rows:** each carries slot, source and drop chance underneath.
- **`GetSkillLineInfo`:** guarded.

**Shoulds, ranked**
1. **The Impaling Harpoon source line still cuts at level 25.** It reads `The Deadmines · Captain Greenskin · 30% dr...`, and the drop chance is the part that is lost. Drop the dungeon name when the boss is named, or shrink to `Captain Greenskin · 30% drop`. The drop chance is the number a raider wants, so cut the area, not the chance.
2. **Level 60 right column has about 300 px of empty ground** under Your sim. Moving Gear left fixed the left column but left the right one short. This is honest, and I do not want invented content. If the owner wants it filled, the real options are `+N more` upgrades (`GearTab.upgrades(doll).count`) or `+28.1 DPS available` (`.total`). Both are optional.
3. **No "+N more" or total** anywhere on the Gear block, as before. The `8 of 14 best in slot` note shows slots, not upgrade potential.
4. **Professions** could show the cap as `Skinning 112 of 150`, if the fifth return of `GetSkillLineInfo` gives it. Optional.

**Needs data**
- Gold, played time and reputation are not read by the addon.
- Score, logs and sim stay behind the companion call to action.

Files: /Users/jh/code/forever-me-tab/design/reviews/2026-11-14-me-empty-and-sources/me-empty-l25.png, /Users/jh/code/forever-me-tab/design/reviews/2026-11-14-me-empty-and-sources/me-empty-l60.png, /Users/jh/code/forever-me-tab/addon/ForeverSixty/surfaces/MeTab.lua.
