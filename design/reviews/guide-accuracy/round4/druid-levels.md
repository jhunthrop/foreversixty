# Druid: Levels 1 to 19 (round 4)

## What was added

Each spec guide (balance, feral, feral-bear, restoration) now opens `## Leveling` with `### Levels 1 to 19` (three or four bullets: spells and class quest, first talents, gear). No heading was added after it.

- **Balance:** Wrath 1 (ranks 6/14), Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18. Talents: Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19. Merged: "take Improved Wrath first" removed from the existing paragraph.
- **Feral:** caster play until 20 (Moonfire 4, Wrath), Bear Form + Maul + Demoralizing Roar 10, Bash 14, Swipe 16. Talents: Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19. Merged: the existing "Before that, play as a caster ... Bear Form from 10" sentence moved into the new block; the existing text now starts "Cat Form arrives at 20. Solo mobs face you...".
- **Feral-bear:** Bear Form + Maul + Growl + Demoralizing Roar 10, Enrage 12, Bash 14, Swipe 16, Maul rank 2 18. Talents: Ferocity 10-14, Furor 15-19 (100% chance of 10 Rage on Bear Form shift at 5/5). Merged: existing "Bear Form comes at 10 and Dire Bear Form at 40" became "Dire Bear Form comes at 40."
- **Restoration:** Healing Touch 1, Rejuvenation 4, Regrowth 12, Revive 12 (out-of-combat res), Cure Poison 14. Talents: Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19.
- **Shared:** Bear Form quest at 10: Heeding the Call (Darnassus / Thunder Bluff), The Great Ursera Spirit on Zephras Isle for Skyborne. Aquatic Form quest at 16 rewards Aquarius Belt (16608) or Idol of Shifting Tides (263411). Druids cannot use wands (no wand skill in trainables; Warcraft Tavern's Forever weapon list has none), so the caster gear line points to a quest staff; ferals take the highest-DPS staff or two-handed mace because forms hit with weapon DPS (index.md).

## Data checks (1.60.1.70755)

- `trainables/druid.json`: every spell level above matches its skill_line_ability rank. Note: spell 5487 (Bear Form) is grouped under the name "Dire Bear Form" with 9634 at 40; the guides call the level 10 form Bear Form.
- `addon-builds.json` `druid-<spec>.leveling.order[0..9]`, mapped 1-based to `talents/druid.json` (tier/column 0-based): balance [1,1,2]x5, [1,1,3]x2, [1,2,1]x3; feral [2,1,2]x5, [2,2,1]x2, [2,2,2], [2,2,3]x2; feral-bear [2,1,2]x5, [2,1,3]x5; restoration [3,1,3]x5, [3,2,2]x3, [3,1,2]x2. Rank descriptions quoted from the client text (Improved Wrath 0.5 sec / 50% mana, Genesis 4%, Moonglow 25%, Ferocity 5, Feral Swiftness 4% dodge, Feral Instinct 10% Swipe, Brutal Impact +1 sec / -30 sec, Furor 100% for 10 Rage, Heart of the Wild 10% Int / 20% bear Stamina, Reflection 50%, Nature's Focus 28%).
- `loot.json` quests: "Aquatic Form" (quest 31 Horde, 5061 Alliance, level 16, druid only) rewards 16608 and 263411. The Bear Form chain has no item reward, so it is confirmed from the sources below only.
- Left out on purpose: no wand named (druids have none); Thorns dropped for length; Polearms (20) are past the range.

## Sources

- Method, WoW Forever Druid Leveling Guide & Talents: https://www.method.gg/wow-forever/wow-forever-druid-leveling-guide-and-talents (class quest table: Heeding the Call, The Great Ursera Spirit, Aquatic Form at 16)
- Warcraft Tavern, Druid Leveling Guide (Forever): https://www.warcrafttavern.com/forever/guides/druid-leveling/ (ability levels 1 to 20, class quests, weapon list without wands, Idol of Shifting Tides)
- Icy Veins, Feral Druid leveling (Forever): https://www.icy-veins.com/wow-forever/feral-druid-leveling (search snippet: Bear Form quest chain at 10 from any druid trainer)
