# Hunter: Levels 1 to 19 (round 4)

Each spec guide now opens `## Leveling` with `### Levels 1 to 19` (four bullets), and the old Leveling prose sits under `### Levels 20 to 60`.

## Added per spec

- **Beast Mastery**: first trainables (Serpent Sting and Aspect of the Monkey 4, Arcane Shot and Hunter's Mark 6, Concussive Shot 8, Aspect of the Hawk 10), the level 10 taming quests and Mend Pet 12, the talent start, and ranged quest weapon plus leather to 40. The existing Deadly Aspects / Focused Fire paragraph and "tame at 10" line moved into the new subsection, so nothing is said twice.
- **Marksmanship**: the same trainables plus Multi-Shot 18, the taming quests, Lethal Attacks 10 to 14, Improved Stings 15, Efficiency from 16, and the same gear line.
- **Survival (melee)**: Raptor Strike 1 (ranks 8 and 16), Monkey and Serpent Sting 4, Hunter's Mark 6, Wing Clip 12, Mongoose Bite 16 (after a dodge), the taming quests, Improved Tracking 10 to 14 with the matching Track spells (Beasts 1, Humanoids 10, Undead 18), Savage Strikes 15 and 16, Survivalist from 17, and a slow two-handed quest weapon plus leather to 40.

## Data checks (1.60.1.70755)

- Levels come from `trainables/hunter.json` rank levels. I left out the level 1 `Aspect of the Viper` (415423) and `Heart of the Lion` (409580) entries: they are SoD rune spell ids, not trainer spells. `Wyvern Strike` and `Widow Bite` are `class_spell` entries (a talent and a pet ability), so they are not level 1 to 19 trainables either.
- Talent orders come from `addon-builds.json` `hunter-<spec>.leveling.order[0..9]`, mapped against `talents/hunter.json`:
  - BM: Deadly Aspects x5 (10 to 14), Focused Fire x2 (15, 16), Endurance Training from 17.
  - MM: Lethal Attacks x5, Improved Stings x1 (15), Efficiency from 16.
  - SV: Improved Tracking x5, Savage Strikes x2 (15, 16), Survivalist from 17.
- Per-rank effects come from the talent rank descriptions: Deadly Aspects 2% per rank, Focused Fire 1% per rank, Lethal Attacks 1%, Improved Stings rank 1 is 6% Serpent Sting, Efficiency 3%, Improved Tracking 1%, Savage Strikes 2%, Survivalist 2%.
- Spell text comes from client `Spell.csv` (1.60.1.70532, the newest build with raw Spell.csv): Mongoose Bite "Can only be performed after you dodge"; Raptor Strike "melee weapon damage plus X". The second one is why the guide says to use a slow weapon.
- Armor (leather to 40, mail at 40) follows the owner's Forever class rules, and the existing guide text already says "Mail at 40". I named no wand because hunters do not use one.

## Sources

- Warcraft Tavern, Hunter Pets Guide (Forever): https://www.warcrafttavern.com/forever/guides/hunter-pets/ (taming chain from the trainer at 10 teaches Tame Beast; feeding and happiness)
- Wowhead Classic, Hunter class quests: https://cn.classic.wowhead.com/guides/hunter-class-quests-classic-wow (the level 10 chain is per race)
