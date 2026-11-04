# Mage: Levels 1 to 19 (round 4)

## What was added

Each spec guide (arcane, fire, frost) now opens `## Leveling` with `### Levels 1 to 19` (three bullets: spells, first talents, wand and gear). The existing leveling text now sits under a new `### Levels 20 to 60` heading, so it no longer renders as part of the 1 to 19 block.

- **Arcane:** Frostbolt + Conjure Water 4, Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14. Talents: Wand Specialization 10, Improved Channeling 11-14, Arcane Resilience 15, Arcane Concentration 16-19. The old first sentence "Put every point in Arcane from 10" became "Keep every point in Arcane."
- **Fire:** Fireball 1 (ranks 6/12/18), Frostbolt 4, Fire Blast 6, Frost Nova 10, Flamestrike 16. Talents: Incineration 10-14, Accelerant 15, Ignite 16-19.
- **Frost:** Frostbolt 4 (ranks 8/14/20), Conjure Water 4, Fire Blast 6, Frost Nova 10. Talents: Frost Warding 10-11, Improved Frostbolt 12-16, Precision 17-19. Merged: these talents were removed from the existing paragraph, which now starts at Ice Shards in the early 20s.
- **Shared gear bullet:** Lesser Magic Wand (11287, req 5) and Greater Magic Wand (11288, req 13) from Enchanting. Quest wands: Riptear's Spare Arm (Rear Guard Patrol, Horde, q13) and Elven Wand (Bashal'Aran, Alliance, q13). The level 10 class quest pays a staff, orb or wand (Mirror Lake: Ley Staff/Ley Orb; Mage-tastic Gizmonitor and The Balnir Farmstead: Arcane Staff/Arcane Orb; Boughs in the Wind: Windbough Wand). Cloth quest rewards over vendor gear.

## Data checks (1.60.1.70755)

- `trainables/mage.json`: every level named above matches its skill_line_ability rank.
- `addon-builds.json` `mage-<spec>.leveling.order[0..9]`, mapped 1-based to `talents/mage.json` (tier/column 0-based): the orders and rank descriptions match the text (Wand Spec 13%, Incineration 10% at 5/5, Accelerant 0.1 sec, Frost Warding 30%, Improved Frostbolt 0.5 sec, Precision 6%, Arcane Resilience 25% of Int).
- `items.json` / `loot.json` quests: confirmed the wands and class quest rewards listed above.
- Left out on purpose: Sleep (level 8) and the four Volleys (level 8) show up in the data only as `class_spell`, not SkillLineAbility, so it is not confirmed a trainer teaches them. No Horde robe quest is named because Spellfire Robe is not in loot.json.

## Sources

- Warcraft Tavern, Mage Leveling Guide (Forever): https://www.warcrafttavern.com/forever/guides/mage-leveling/ (class-quest names at 10 and 15; spell levels cross-checked against the data)
- Mobalytics, Mage leveling guide (Forever): https://mobalytics.gg/wow-forever/classes/mage-leveling-guide
- Icy Veins, Frost Mage leveling (Forever): https://www.icy-veins.com/wow-forever/frost-mage-leveling (403 on fetch; search snippet only)
