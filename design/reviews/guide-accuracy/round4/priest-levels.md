# Round 4: Priest "Levels 1 to 19"

Added a `### Levels 1 to 19` subsection as the first thing under `## Leveling` in `web/src/content/guides/priest/discipline.md`, `holy.md` and `shadow.md` (index.md untouched). Three bullets each, about 140 to 150 words.

## What each spec got

- **Shadow:** Smite carries damage to 10 (SW:P 4, Smite, wand); PW:S 6, Mind Blast 10, Inner Fire 12, Psychic Scream 14. Talents 10 to 18: Shadow Focus 5 (5% Shadow hit), Blackout 4 (stun chance). Gear and wand bullet. Merged with the existing build paragraph: its "Shadow Focus from 10 to 14, Blackout to 18" opening now reads "From 19 the build takes Spirit Tap to 21, ..." so the talent path is not said twice.
- **Discipline:** PW:S 6, Renew 8, Mind Blast 10, Inner Fire 12, SW:P 4 opener, Dispel Magic 18. Talents: Twin Disciplines 10 to 14 (instant spells 5%: SW:P, PW:S, Renew), Mental Strength 15 to 19 (10% Intellect). Gear and wand bullet.
- **Holy:** Smite is damage until Holy Fire at 20 (SW:P 4, Smite, wand); PW:S 6, Renew 8, Inner Fire 12. Talents: Twilight Focus 10 to 12 (pushback, 70% at 3), Improved Renew 13 to 15 (15%), Holy Specialization from 16. Existing text keeps Heal at 16 and downranking, not repeated.
- **Shared gear bullet:** cloth, quest rewards over vendor, replace the starting mace with the first quest staff or caster mace; wands: Lesser Magic Wand (Enchanting, required level 5), Greater Magic Wand (required level 13), quest wands Elven Wand (Bashal'Aran, Darkshore, Alliance) and Cinder Wand (Arugal's Folly, Silverpine Forest, Horde).

## Data checks (1.60.1.70755)

- `trainables/priest.json`: Smite, Lesser Heal, PW:F 1; SW:P 4; PW:S 6; Renew, Fade 8; Mind Blast 10; Inner Fire 12; Psychic Scream 14; Heal 16; Dispel Magic 18; Holy Fire, Devouring Plague, Fear Ward, Flash Heal 20. Racial priest spells at 10/20 match Warcraft Tavern (Gnome: Confounding Flash 10, Contingency Plan 20). Shadowfiend shows at level 1 in the trainables but is the SoD rune that index.md says is not in Forever, so it is not mentioned.
- `addon-builds.json` leveling orders decoded against `talents/priest.json` (tiers/columns 0-based in the file, 1-based in the order): Disc TD 10-14, MS 15-19; Holy TF 10-12, IR 13-15, HS 16-20; Shadow SF 10-14, Blackout 15-18, Spirit Tap 19-21. Descriptions from the rank text.
- `items.json`: Lesser Magic Wand 11287 req 5, Greater Magic Wand 11288 req 13. `loot.json` quests: Elven Wand 5604 from Bashal'Aran (alliance, quest 13), Cinder Wand 5242 from Arugal's Folly (horde, quest 15).
- Not named: the level 4/5 "Garments of ..." robe quests are in classic-db and loot.json but no Forever source confirms them, so they stay out.

## Sources

- Warcraft Tavern, Priest Leveling Guide (Forever): https://www.warcrafttavern.com/forever/guides/priest-leveling/ (trainer levels, racial spells at 10 and 20, Lesser/Greater Magic Wand at 5/13)
- Icy Veins, Shadow Priest Leveling (Forever): https://www.icy-veins.com/wow-forever/shadow-priest-leveling (403, not read)
- warcraft.wiki.gg, Garments of the Light (Elwynn Forest): https://warcraft.wiki.gg/wiki/Garments_of_the_Light_(Elwynn_Forest) (Classic only, hence omitted)
