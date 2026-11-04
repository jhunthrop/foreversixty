# Warlock: Levels 1 to 19 (round 4)

Added `### Levels 1 to 19` as the first block under `## Leveling` in affliction.md, demonology.md and destruction.md (index.md untouched). Three bullets each: first trainables, the leveling talent start from 10, gear and wand. Existing Leveling prose follows unchanged.

## Per spec

- **Affliction**: Corruption 4, Life Tap 6, Bane of Agony 8, Voidwalker class quest 10, Drain Life 14. Talents: Improved Life Tap 10 to 11, Suppression 12 to 16, Improved Corruption 17 to 21 (instant at 21); tip to open with the 2 s Corruption until then.
- **Demonology**: Corruption 4, Life Tap 6, Bane of Agony 8, Voidwalker 10, Health Funnel 12. Talents: Improved Health Funnel 10 to 11 (live from 12), Demonic Embrace 12 to 14, Unholy Power 15 to 19.
- **Destruction**: Immolate 1, Corruption 4, Life Tap 6, Bane of Agony 8, Voidwalker 10, Drain Life 14. Talents: Destructive Reach 10 to 11, Bane 12 to 16, Molten Skin 17 to 19.
- **All**: Lesser Magic Wand (11287, req 5) and Greater Magic Wand (11288, req 13), both Enchanting; cloth quest rewards over vendor gear; replace the starting weapon with the first Int/Stam quest staff or dagger.

## Data checks (1.60.1.70755)

- trainables/warlock.json: Immolate 1, Corruption 4 (rank 1 cast 2000 ms), Life Tap 6, Bane of Agony 8, Fear 8, Drain Soul/Create Healthstone/Summon Voidwalker 10, Health Funnel 12, Drain Life 14. Shadow Bolt and Immolate are Destruction skill line (Destructive Reach applies).
- addon-builds.json `warlock-<spec>`.leveling.order indexes 0 to 11 decoded against talents/warlock.json; talent effects quoted from the rank descriptions (ILT 20%, Suppression 5%/20%, Imp Corruption 2 s/10%, IHF 40%/30%/no threat/any health, Demonic Embrace 9% at 3/5, Unholy Power 10%, Destructive Reach 20%, Bane 0.5 s, Molten Skin 6% at 3/5).
- items.json + loot.json: 11287 and 11288 sourced from Enchanting.
- Not named: race specific Voidwalker quest titles (not confirmable in the 70755 data); said "class quest from your warlock trainer".
- Note: Warcraft Tavern's beta build takes Improved Corruption first; the shipped order takes Improved Life Tap and Suppression first, so the text follows the data.

## Sources

- Warcraft Tavern, Warlock Leveling Guide (Forever): https://www.warcrafttavern.com/forever/guides/warlock-leveling/ (Voidwalker by class quest at 10, trainer levels, Enchanting wands)
- Blizzard forums, Warlock 1-20 Feedback (Forever beta): https://us.forums.blizzard.com/en/wow/t/2362831 (levels 6 to 10 rough until the Voidwalker; wand heavy)
- Mobalytics, WoW Forever warlock leveling guide: https://mobalytics.gg/wow-forever/classes/warlock-leveling-guide
- Icy Veins, Affliction Warlock Leveling (Forever): https://www.icy-veins.com/wow-forever/affliction-warlock-leveling
