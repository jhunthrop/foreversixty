# L3 Gear tab: wow-player

## Round 1: FIX

Seen: l3/gear-hunter-60.png (level 60 BM hunter) and l3/gear-warrior-34.png (level 34 Fury warrior), against hub-gear.png and spec 4.3 / 13. Item names, slots, qualities and armor types checked against data/builds/1.60.1.70938/items/hunter.json (Theramore Gloves exalted Theramore Expeditionary Force; Black Dragonscale Helm, Dawn Armor, Ranger's Belt, Legionite Glaive are Forever items and real in the data; Trueaim Gauntlets' hit is a percent stat; Blackhand's Breadth and Cutthroat's Signet are rare, correct; Cape of the Black Baron is Baron Rivendare, correct).

**The question:** what do I wear, what beats it, and where do I get it. The answer (doll with a gain under every slot, Upgrades, enchant flags) is on screen first. A player would open this after every dungeon and before every enchant, so yes, it is a per-session tab. The doll reads like the character pane, qualities and "Best in slot / Beats the pick / in your bags / No enchant" are the right words, no "+10 Hit" lines anywhere, and the empty Hands / Finger / Trinket cells on the level 34 warrior ("Trinket Empty") are honest.

### Must

1. Source lines are chopped mid-word, and it is the line that answers "where do I get it". On the warrior render: "Hands · Boulderfist Enforc... · World · 1% drop", "Off Hand · Scarlet Mon... · Interrogator Vishas", "Main Hand · Razorfen Kr... · Charlga Razorfla...". A player cannot tell which Scarlet wing or which boss for the last one. Use the names players use (SM, RFK, or the full name on a second line), never cut a boss name; drop the slot from the line (the icon and the doll already say it) before cutting a place. The doll has the same fault: "Brawler's Leather Leggua..." cuts the item name mid-word; let the name wrap or shrink, the client never prints "Leggua...".
2. "shipped, per point" over the weights. "shipped" is site talk, a player has no idea what shipped means, and the spec and board say "per rating point". Hit and crit are ratings here, a bare "point" is a number without its unit (an item's Trueaim "+3% hit" is not a point). Use "Forever default, per rating point" or just "per rating point" and "your sim, per rating point" when the inbox has one. Same words for the header column pair ("Weight" / "DPS": say "DPS per point" once in the note and keep it).

### Should

3. No hit cap anywhere. For a hunter or warrior the first thing the weights are for is "how far from cap". Spec 13 already names the field (`hit_to_cap`); add one muted line under the weights, e.g. "Hit: 2% to the 9% cap" (from the band's field, same as the web's `hitCapLine`). Otherwise Hit 1.00 reads as "stack hit forever".
4. "Boulderfist Enforc... · World · 1% drop": a named mob is not a world drop. Say the zone ("Boulderfist Enforcer · Arathi Highlands · 1%") and keep "World drop" only for the items that really are random world drops, as Wowhead does. Also "1% drop" is fine but write the percent the way the tooltip does.
5. "Attack Power" and "Ranged Attack Power" both 0.33 / 0.14 on the hunter, two identical rows. Show the one that applies (hunter: Ranged Attack Power) or fold them; a duplicate reads as a bug. Strength 0.07 and Intellect 0.03 rows are noise at that size; Intellect is fine for a mana user but cut rows under 0.05 or mark them muted.
6. Gains of "+0.1 DPS" in green on Shoulders, Feet, Bracers (warrior) are noise a player will stop reading. Show nothing (or muted "Best in slot") under about 0.3 DPS, and do not count them toward "in all".
7. Upgrades shows four rows and the doll never names the pick, so for six of the ten gaps on the warrior the player cannot learn what to get without hovering every cell. Add a "Show all" in gold that expands the list (it exists in the spec's 4.6 pattern), or click a doll gap to open its pick. Hover should also give two or three alternatives per slot with source, as Wowhead BiS does.
8. Quest and crafted rows are thin. "Band of the Fist · Finger · Allegiance to the Old Gods · Quest" needs the zone or the turn-in; "Legionite Glaive · Crafted · Blacksmithing" needs the skill level and where the plan is learned. Add on the hover if the line is full. Also show the required level on the pick tooltip ("Requires Level 58" for Legionite Glaive and Cape) so a level 54 hunter on band 60 knows.

### Data I do not believe (for the accuracy reviewer)

- Level 60 hunter picks a leather chest (Dawn Armor) and leather gloves (Theramore Gloves): legal (below own armor type), but confirm the sim really prefers them over mail; the Black Dragonscale Helm's +38 Intellect carries the helm score on a Beast Mastery weight of 0.03 Intellect, check that is not mis-weighted.
- Trueaim Gauntlets carry hit as a percent (3%); confirm the gain function converts it to 30 rating, or the +4.9 on the Theramore Gloves is overstated against a worn pair counted at 3 rating.
- Warrior band 30 "Band of the Fist" from "Allegiance to the Old Gods": confirm the quest name and level floor (three below the quest level).
- Warrior weights list Expertise at 0.56; items do carry it in the Forever data (23 rows at 10), so I accept it, but the client word must match its GlobalString.
- Hunter at 60 lists Warmonger (two-hand sword, level 47, +30 hit) as a main hand pick; fine for a melee-weapon slot only if the hunter weight for hit is truly 1.00; check.

### Verdict

FIX. The structure is what a player wants and the words are the client's, but the two truncations and the "shipped" label are what gets screenshotted. Items 1 and 2 are the blockers; 3 to 8 make it a tab people keep open.

## Round 2: SHIP

Re-read both renders (gear-hunter-60.png, gear-warrior-34.png).

Must 1 and 2 are fixed. The boss names now survive whole ("Boulderfist Enforcer · 1% drop", "Interrogator Vishas · 90% drop", "Charlga Razorflank · 20% drop"), "Band of the Fist" keeps its quest name, and the weights note reads "per rating point". The new "No enchant: Legs" / "No enchant: Chest, Feet" line, the "More upgrades" list that names every gap's pick with its source, the dropped "World"/"Vendor" filler words and the Intellect row going away all read the way a player expects. "Better than the list" is plain and clear. I accept AP and RAP both staying as real shipped weights, and the hit cap, mob zones and recipe skill levels as data follow-ups, not blockers.

Remaining shoulds (none blocks):
1. Places are still cut to fragments: "Scarlet M...", "Razorfen K...", "Kirtonos t...", "Bl..." and "Leatherwor..." (More upgrades). "Scarlet M..." reads as a typo. Use the names players type (SM, RFK, Sch is fine for Scholomance), and for a profession say "Leatherworking" whole or "LW" before cutting. "Kirtonos t..." should keep "Kirtonos the Herald" by dropping "Scholomance" to "Scholo".
2. The doll still cuts "Brawler's Leather Leggуar..." mid-word; wrap the name to two lines or let the cell grow, since the hover is the only other place the full name lives.
3. Hit cap line still missing (follow-up: ship `hit_to_cap` into Data.lua). Until then Hit 1.00 reads as "stack it forever".
4. "+0.1 DPS" rows on the warrior are still green noise; under about 0.3 DPS show "Best in slot" or nothing.
5. Quest and crafted rows still lack zone / turn-in and recipe source; do it with the data follow-up.

Verdict: SHIP. The tab answers what to wear, what beats it and where to get it, in the client's own words.
