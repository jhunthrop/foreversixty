# info-density-reviewer, round 2 (Me with no site activity)

Verdict: PASS

Both items I asked for are done. The upgrade rows now carry slot, source and drop chance, and the hit-cap argument holds: `data.hitToCap` describes the list's set, not what the character wears, so leaving it out is correct. Professions is a real client read and fills the level 60 left column honestly. I checked `me-empty-l25.png` and `me-empty-l60.png` and `readProfessions` in MeTab.lua.

**Must-fix**
None.

**Shoulds, ranked**
1. **Enchant rows truncate and repeat themselves.** In `me-empty-l25.png` the rows read `No enchant: Enchant Chest - Minor Stats · Che...` and `No enchant: Enchant Bracer - Minor Agility · ...`. The slot is already in the enchant's name, and it is cut off after it anyway. Drop the "No enchant:" prefix, since the warning colour carries it. Drop the trailing slot, or put it first (`Chest  Minor Stats`) so a cut loses the effect, not the slot. This frees about 80 px per row. At 390 wide, the whole row needs this fix.
2. **Level 60 left column is still thin.** `me-empty-l60.png` has two rows and about 400 px of empty ground, against a full right column. Cheap real fillers:
   - The XP-style row is gone at 60. Add `Gear` / `9 of 16 best in slot` as a left row.
   - Add `Upgrades` / `+28.1 DPS available`, from `GearTab.upgrades(doll).total` (my earlier should 2).
   This also lets the right column shed its heading note. If the owner prefers one Gear home, move the whole Gear block left at 60 only, because Right now is full there.
3. **"+N more" is still missing.** `GearTab.upgrades(doll).count` minus the three shown, as a muted note on the last upgrade row, links the player to the Gear tab for the rest.
4. **Professions guard.** `readProfessions` checks `GetNumSkillLines` but then calls `_G.GetSkillLineInfo` unguarded (MeTab.lua:670). If a client has one and not the other, it throws inside `readOwn`. Guard it with `type(_G.GetSkillLineInfo) == "function"`. `readOwn` runs the gear read under `pcall` but this one is bare, so the whole tab would fail.
5. **Professions row is plain text.** A veteran would expect the skill cap too (`Skinning 112 of 150`). Only add it if the skill line's max rank is returned: `GetSkillLineInfo` gives it as the fifth return. This is cheap and real.

**Needs data**
- Gold, played time and reputation are still not read by the addon.
- Score, logs and sim stay behind the companion call to action.

Files: /Users/jh/code/forever-me-tab/addon/ForeverSixty/surfaces/MeTab.lua, /Users/jh/code/forever-me-tab/design/reviews/2026-11-14-me-empty-and-sources/me-empty-l25.png, /Users/jh/code/forever-me-tab/design/reviews/2026-11-14-me-empty-and-sources/me-empty-l60.png.
