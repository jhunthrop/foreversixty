# Paladin: Levels 1 to 19 (round 4)

Added a `### Levels 1 to 19` subsection as the first block under `## Leveling` in holy.md, protection.md and retribution.md (index.md untouched). The existing Leveling text now sits under `### Levels 20 to 60`.

## Per spec

- **Retribution**: Seal of Righteousness and Holy Light 1, Judgement and Blessing of Might 4, Holy Strike 6, Hammer of Justice 8, Lay on Hands 10; Judgement keeps the Seal, both on a 10 sec cooldown. Redemption class quest at 12. Talents 10-14 Benediction 5/5 (10% instant mana), 15-16 Improved Judgement 2/2 (8 sec), 17-19 Conviction 3/3 (6% crit). Slow two-hander (Seal of Righteousness scales with weapon speed), quest rewards over vendor, no wands, mail to 40.
- **Holy**: Holy Light 1 (ranks 6, 14), Judgement 4, Holy Strike 6, Purify 8, Lay on Hands 10, Blessing of Wisdom 14; Redemption quest at 12. Talents (if leveling Holy) 10-14 Spiritual Focus 5/5 (70%), 15-17 Healing Light 3/3 (12%), 18-19 Reverence 2/2 (30% regen while casting). Same gear bullet. Existing text (level as Ret, Flash of Light at 20, Oath of Light at 25) not repeated.
- **Protection**: Devotion Aura 1, Judgement 4, Holy Strike 6, Hammer of Justice 8, Righteous Fury and Seal of Fury 10 (shield absorb, judgement taunts 3 sec); Redemption quest at 12. Talents 10 Divine Intellect 1/5 (2% Int), 11-15 Redoubt 5/5 (30% block, 10 sec or 5 blocks), 16-18 Precision 3/3 (3% hit), 19 Sacred Duty 1/2 (30 sec off Divine Protection). One-hander and shield; mail to 40. Merged: the existing "you get Righteous Fury and Seal of Fury at level 10" now reads "Righteous Fury and Seal of Fury are enough to hold..." so the level is said once.

## Data checks (1.60.1.70755)

- trainables/paladin.json: Devotion Aura, Holy Light, Seal of Righteousness 1; Blessing of Might, Judgement 4; Divine Protection, Holy Strike, Seal of the Crusader, Holy Light r2 6; Hammer of Justice, Purify 8; Blessing of Protection, Lay on Hands, Righteous Fury, Seal of Fury 10; Redemption 12; Blessing of Wisdom, Holy Light r3 14; Retribution Aura 16; Blessing of Freedom 18.
- Warcraft Tavern lists Righteous Fury at 16; the data has it at 10 (id 25780), so the guide follows the data.
- spelltips/paladin.json: Judgement "Does not consume the Seal", 10 sec cooldown; Holy Strike 10 sec cooldown; Seal of Righteousness "Slower weapons cause more Holy damage per swing"; Seal of Fury shield absorb 50% of the Holy damage, judgement deals Holy damage and taunts for 3 sec.
- addon-builds.json leveling.order decoded against talents/paladin.json (file tiers/columns 0-based, order 1-based): holy [1,1,2]x5 Spiritual Focus, [1,2,1]x3 Healing Light, [1,2,2]x2 Reverence, [1,2,4] Purifying Power at 20; protection [2,1,2] Divine Intellect, [2,1,3]x5 Redoubt, [2,2,1]x3 Precision, [2,2,4] Sacred Duty at 19, Templar's Bulwark 20; retribution [3,1,3]x5 Benediction, [3,2,1]x2 Improved Judgement, [3,2,2]x3 Conviction, Sanctified Judgement from 20. Rank text from the talent descriptions.
- combos.json: paladins are Human, Dwarf and (new in Forever) Undead; no race-specific quest named.
- Armor: mail to 40, plate at 40 (owner class rule). No wand (paladins cannot use them).

## Sources

- Warcraft Tavern, Forever Paladin Leveling Guide: https://www.warcrafttavern.com/forever/guides/paladin-leveling/ (Redemption class quest at 12, Retribution talents early even for tanks, Seal of Righteousness weapon speed, Seal of Fury as the tank seal)
