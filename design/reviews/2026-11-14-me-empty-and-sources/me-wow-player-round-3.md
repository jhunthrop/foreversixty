# wow-player, round 3 (Me with no site activity)

Verdict: FIX

The enchant rows, the "Spend 1 point" line, the source lines and the level 5 prompt are all fixed. One real defect is left, plus one small should.

**Must-fix**
1. **The level-25 no-companion capture still truncates the Harpoon source.** The line reads "The Deadmines · Captain Greenskin · 30% dr…". The level-26 capture fits it whole because Gear moved to the left column. The right column in me-empty-l25 is too narrow for the same string. The player sees "30% dr…" on the one drop they care about. Either use the same wide layout, drop the "30%" suffix into the right-aligned DPS column, or shorten to "Captain Greenskin · 30%". The "drop" is implied.

**Shoulds**
2. **"Level 10" under "First talent" and "Level 30" under "New gear list" read fine now.** One more tweak would help: "New gear list: Level 30" is still vague. Say "Gear list 30" or name the band, since the Gear tab labels it that way.
3. **"8 of 14 best in slot" and "Rested 40%" stay as you decided.** Same note as before: the in-game word for "on your list" is not "best in slot". It is your call.
4. **"Spend 1 point" at level 26 now says what to do but not where.** "Builds" on the right is the link. Fine as it is.

**Data to check (accuracy reviewer)**
- **Gloves of Holy Might, "World drop", +11.2 DPS.** Still the largest gain on any capture. Confirm it is a hunter-usable glove and a real world drop.
- **Forest Stalker's Bracers, "Silverwing Sentinels", epic.** Confirm the rarity and the reputation standing required.
- **Windreaver Greaves, "Scholomance · Kirtonos the Herald".** This matches the Classic source. Verify.
- **Impaling Harpoon, Captain Greenskin at "30% drop".** Verify the chance against the loot table.
- **Back: Lesser Agility at 60.** Real and sensible, though a 60 would normally have the Greater Agility cloak enchant (if the client has it), so confirm the plan's level gate.

**Good**
- The level 5 prompt is a deliberate choice and reads right.
- The talent counts agree.
- The site line is one clear call.
- The gear rows now carry slot, source and chance.
- "Not simmed yet" with "Run a sim" is honest.
- Moving Gear to the left column fixes the squeeze at 26 and 60.
