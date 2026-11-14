# info-density-reviewer, round 1 (Me with no site activity)

Verdict: FAIL

The page answers its question, but the Gear block drops three facts the addon already holds, and one row is blank on the captures that need it. The build and Right now regions are correct. All three must-fix items are cheap, because `GearTab.read` and `GearTab.doll` have the data in hand when `readGear` runs.

**Must-fix**

1. **Upgrade rows have no slot or source.** `GearTab.fullWhere(cell)` already returns "Feet · Scholomance · Kirtonos the Herald · 12% drop" for every upgrade, with the drop chance and standing built in. `readGear` copies only `item_id` and `gain` into `upgrades` (MeTab.lua:647-649). A raider wants "where do I get Ranger Bow" next to "+3.9 DPS". Fix by adding `where = GearTab.fullWhere(cell)` to each upgrade, then either a muted second line on the row or the existing hover. The spec's own rule says every tooltip fact must also be in the row text, so the line is better. At 20 px rows, three muted second lines add about 36 px. The right column has the room: it is empty below the enchant row in `me-empty-l25.png`.
2. **The "Next" row is missing at level 60, and the level 60 right column is the full one.** At level 60 the left column holds two rows and about 250 px of empty ground (`me-empty-l60.png`), while Gear sits in the right column. This is the weakest panel in the set. The spec drops rows 3 and 4 at 60 when no point waits. That is honest but thin. The local data to fill it:
   - `GearTab.hitCap(data, spec, band, faction, condition)` gives hit to cap, per spec and band, with the lines already built (`data.hitToCap`). Add a "Hit" row to Where you stand or Gear (for example "Hit 7.2 of 8% to cap"), one row from the existing `lines`. It is the first number a level 60 raider checks.
   - The `Progress` points-to-plan figure is shown, but "points spent versus allowed" is a standing fact the left column could also carry at 60, next to Build.
   Do not add more than the hit row. This is a must, because "level 60 fresh" is the most common state of a returning raider.
3. **The left and right columns are unbalanced at level 5.** In `me-empty-l5.png` Right now holds a single "Pick your spec" row beside three left rows, with 200 px of dead ground below it. It is honest. What it lacks is any drill-down: the row's "Builds" link is the only action. The spec accepts this as "thinnest by design", so I treat it as a should (see should 5), not a fail.

**Shoulds, ranked by value**

1. **"Best in slot" tally has no list.** The note "9 of 16 best in slot" counts slots, but the 7 slots not at best are only shown through the top 3 upgrades. `GearTab.upgrades(doll).count` and `.rest` exist. Add a muted "+4 more" on the right of the last upgrade row, linking to the Gear tab (`.rest`, `.count`). Right now the player cannot tell whether there are 1 or 6 more upgrades.
2. **Total upgrade gain.** `GearTab.upgrades(doll).total` is the unrounded sum of every gain. Show "+24.6 DPS available" as the Gear heading's right note, beside or in place of the best-in-slot tally. It is context for the three rows: "these three are 60% of the total".
3. **Enchant row is missing slot context.** "2 enchants missing Chest, Wrist" is right. At 60, `unenchantedSlots` could also name the item worn in that slot, but that is optional. Skip.
4. **Experience row has no figure.** `62%` only; the client gives `UnitXP`/`UnitXPMax` (already read as `xp` and `xpMax`). Show "5,120 of 8,300"? The spec bans thousands separators, so "5120 / 8300" is a plain figure and costs nothing. Optional; the percentage already answers the question.
5. **Level 5 Right now.** Nothing local to add. Leave the empty ground rather than invent data.
6. **Next band row names no item.** "Level 30, a new gear list" could read "Level 30, a new gear list (+N upgrades)" only if the shipped data can count them; it cannot without a doll at that band, so leave it.

**Needs data (do not build)**
- Professions: no `GetSkillLineInfo` read yet.
- Gold, played time, reputation: not read by the addon.
- Anything about score, logs or sim: out of reach until the player links the site or companion, which the one call to action already covers.

**Checks that pass**
- The site link appears once per capture.
- No Right now row is repeated in the left column.
- No invented numbers.
- Level 5 hides Gear correctly.
- Colour use matches the spec: gold for the link, green for gains, warning for enchants.

Files: /Users/jh/code/forever-me-tab/addon/ForeverSixty/surfaces/MeTab.lua (readGear, line 636), /Users/jh/code/forever-me-tab/addon/ForeverSixty/surfaces/GearTab.lua (fullWhere, hitCap, upgrades).
