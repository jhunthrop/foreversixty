# Warrior: Levels 1 to 19 (round 4)

Added a `### Levels 1 to 19` subsection as the first block under `## Leveling` in arms.md, fury.md and protection.md (index.md untouched).

## Per spec

- **Arms**: Charge/Rend 4, Thunder Clap 6, Hamstring 8, Overpower 12, Tactical Mastery 14 (keeps 10 Rage); level 10 race warrior quest (Defensive Stance, Sunder Armor, Taunt, one-hander follow-up); talents 10 Improved Heroic Strike, 11-13 Deflection, 14-16 Improved Rend, 17-18 Improved Charge; quest/dungeon two-hander over the starter, mail until plate at 40.
- **Fury**: same abilities and quest; talents 10-14 Cruelty 5/5, 15 Lingering Rage 1, 16-20 Unbridled Wrath 5/5; Dual Wield at 20, two-hander until then; mail to 40.
- **Protection**: Thunder Clap 6, quest at 10, Shield Bash 12, Revenge and Demoralizing Shout 14, Shield Block 16 (blocks 2 attacks); talents 10-11 Improved Bloodrage, 12-16 Shield Specialization, 17+ Anticipation; one-hander and shield over the starter; mail to 40. Merged: removed the old "Defensive Stance trains at 10 and Shield Block at 16" and "take Improved Bloodrage, Shield Specialization and Anticipation first" sentences from the existing paragraph, which the new subsection now covers.

## Data checks (1.60.1.70755)

- trainables/warrior.json: Charge 4, Rend 4, Thunder Clap 6, Hamstring 8, Bloodrage/Defensive Stance/Sunder Armor/Taunt 10, Overpower/Shield Bash 12, Revenge/Demoralizing Shout/Tactical Mastery (1310185) 14, Shield Block 16, Victory Rush/Slam/Cleave/Retaliation 20, Berserker Stance 30.
- spelltips/warrior.json: Tactical Mastery "retain up to 10 Rage"; Shield Block "75% block for 7 sec, only 2 attacks"; Revenge follows block/dodge/parry; Rend is a bleed.
- addon-builds.json leveling.order mapped against talents/warrior.json (1-based tab/tier/column): arms [1,1,1] IHS, [1,1,2]x3 Deflection, [1,1,3]x3 Improved Rend, [1,2,1]x2 Improved Charge, then Improved Tactical Mastery; fury [2,1,3]x5 Cruelty, [2,2,2] Lingering Rage, [2,2,3]x5 Unbridled Wrath (5/5 = 60%), [2,3,1] Furious Precision at 21; protection [3,1,1]x2 Improved Bloodrage (50%), [3,1,2]x5 Shield Specialization (5 Rage on block at 5/5), [3,2,1] Anticipation from 17.
- Dual Wield is not in trainables (skill, not a spell); level 20 taken from Warcraft Tavern Forever and consistent with Fury's Furious Precision (off-hand hit) at 21.
- Armor: mail to 40, plate at 40 (owner class rule). No wand (not a caster).

## Sources

- Warcraft Tavern, Forever Warrior Leveling Guide: https://www.warcrafttavern.com/forever/guides/warrior-leveling/ (level 10 class quests teach Defensive Stance, Taunt, Sunder Armor; Dual Wield at 20)
- Method, WoW Forever Warrior Leveling Guide and Talents: https://www.method.gg/wow-forever/wow-forever-warrior-leveling-guide-and-talents (Defensive Stance quest reward one-hander follow-up; weapon dependence)
- Icy Veins, Forever Arms Warrior guide: https://www.icy-veins.com/wow-forever/arms-warrior-melee-dps-pve-guide (Tactical Mastery at 14, 10 Rage)
