# Player verdict: mobile Home (iOS and Android), 2026-10-31

Question this screen answers: "What should I replace on my character, and how much does it help?" The answer (11 upgrades, +76.2 DPS) is the first thing under the name. Good.

## Verdict: HOLD

## Blockers
1. The upgrades list looks unsorted. Rows read +11.1, (no figure), +23.1, +11.0, +4.9. The biggest gain, Hyper Deluxe Sniper Rifle at +23.1, is third. Frozen Heart of the Mountain sits second with a blank gain column. A player reads this as broken sorting or a missing number. Either sort by gain with the unmodelled rows last, or show a visible placeholder in the gain column. The site's order must be explained if it is deliberate.
2. Sources are raw strings or bare names, not Wowhead-style lines. Amulet of the Darkmoon shows "Neck · 1200 Tickets - Amulet of the Darkmoon": the item name repeats itself and wraps to two lines. Beastmaster's Cap shows only "Saving the Best for Last", with no "Quest:", no NPC, no zone. Legionite Glaive shows only "Blacksmithing", with no recipe source and no skill level. Raider Gloves show "Creeg Bothunk" with no zone and no drop chance. Add "Quest · NPC · zone", "Crafted · 300 Blacksmithing · recipe from X" and "{n}% from {boss}".
3. The weapon tooltip is not what the client prints. It shows "Main hand ... Two-Handed Weapon". The client says "Two-Hand" on the left and the weapon type (Polearm, Sword) on the right. The weapon type is missing entirely.
4. Off hand shows "Empty" in Your gear for a character wielding a two-hander. It reads as a missing upgrade. Hide the row or mark it "Two-hander equipped".
5. Faction and phase picks to verify (accuracy reviewer):
   - Sentinel's Chain Leggings is picked for an Orc hunter. I believe it is the Alliance Silverwing item and the Horde one is Outrider's. If so it is a wrong pick.
   - Signet Ring of the Bronze Dragonflight is a ring pick for "60 · raid". It belongs to the Ahn'Qiraj quest chain, which is not available at launch.
   - Dawnstalker Leggings is shown at ilvl 98 and Sadist's Collar at ilvl 83. Confirm both against Era data; they look like later-tier items.
   - Raider Gloves, Dawn Armor and Hyper Deluxe Sniper Rifle Mk XVII: I do not recognise them. Confirm they are real Forever items.

## Shoulds
- Effect lines have no "Equip:" or "Use:" prefix. "+54 Attack power" is lowercase here while the set bonus says "Attack Power". Stats are not in the client's order or wording. The Binds, armor and standing gaps are known and owned by another lane, but the tooltip is not a tooltip until they land.
- The set block lists bonuses but not the six piece names with owned ones highlighted, as the client does. There is no "(0/6)" count.
- The tooltip repeats the item name twice: once in the sheet header and again inside the box. Drop the header duplicate or the box title.
- Runners-up show only "Same DPS as the pick". Show source and rank for High Warlord's Pig Sticker and Poker (rank 18, not obtainable at launch) and the Warlord's helms, so the player knows what is reachable.
- iOS Home has about 380px of dead black above the hero with only a tiny "Velaren". There is no class crest. "UPGRADES FOR YOUR LEVEL RANGE" is odd for a level-60 character. "Best in slot for 60, together" is unclear.
- "+76.2 DPS" is a sum that leaves out six unmodelled gains. Say "at least +76.2 DPS", or say it is the sum of the modelled rows.
- "Not in our item data yet" is honest, but the icon is a blank dark tile, the same as "Empty". Give it a question-mark icon and the item ID, and make it tappable so the player can see what we think they wear.
- Gear rows put 3 lines of wrapped name beside two icons ("Dawnstalker Legguards ilvl 98 → Sentinel's Chain Leggings"). Names wrap hard on narrow phones. The gear list has no per-slot DPS gain, so it cannot be used to prioritise.
- The Frozen Heart of the Mountain row says "Gain not modelled yet" and the Use effect is not described. Say it is a Forever craft.
- The quality colours and the icon borders are right: purple epic and blue rare. The "DPS" unit is always shown, and "+11.1 DPS" reads well.
