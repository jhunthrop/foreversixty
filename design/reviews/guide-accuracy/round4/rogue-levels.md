# Rogue: Levels 1 to 19 (round 4)

Added `### Levels 1 to 19` as the first block under `## Leveling` in assassination.md, combat.md and subtlety.md (index.md untouched).

## Per spec

- **Assassination**: training path (Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Garrote 14, Ambush 18), the Gouge, step behind, Backstab trick, Poised at 18, talents Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19; leather, quest rewards, replace Worn Dagger, throwing knives.
- **Combat**: training path (Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush 18), Poised; talents Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds from 19; quest sword/axe/mace over the Worn Dagger. Merged: removed "Efficient Strikes at 13" from the existing paragraph (now in the new block); it keeps "Aggression by 30".
- **Subtlety**: training path (Backstab 4, Gouge 6, Slice and Dice and Sprint 10, Garrote 14, Ambush 18), keep a dagger, Poised; talents Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19. Merged: existing paragraph no longer restates Master of Deception / Opportunity, now starts at Premeditation 20.

## Data checks (1.60.1.70755)

- trainables/rogue.json: Sinister Strike, Eviscerate, Stealth 1; Backstab 4; Gouge 6; Evasion 8; Sap, Slice and Dice, Sprint 10; Kick 12; Garrote, Expose Armor 14; Feint 16; Ambush, Poised 18; Instant/Crippling Poison, Rupture 20. Cheap Shot is 26 (so it is not named as a sub-20 opener). Dual Wield and Pick Lock are not in the trainables file, so no level is named for them.
- spelltips/rogue.json: Poised (1328825) "While at maximum Energy, and out of combat, reduce the Energy cost of your next ability by 10." Ambush rank 1 costs 60 (50 with Poised). Backstab and Ambush require a main-hand dagger.
- addon-builds.json leveling.order, decoded against talents/rogue.json:
  - rogue-assassination: [1,1,1]x3 Improved Gouge (+1.5 s at 3/3), [1,1,2]x2 Remorseless Attacks (40% at 2/2), [1,1,3]x5 Malice (5%), then Ruthlessness at 20.
  - rogue-combat: [2,1,1]x3 Coup de Grace (Eviscerate +20% at 3/3, Combat tree in Forever), [2,1,2]x1 Efficient Strikes (3 Energy at 1/2), [2,1,3]x5 Lightning Reflexes (5% dodge), [2,2,1] Puncturing Wounds from 19.
  - rogue-subtlety: [3,1,2]x3 Master of Deception (3 levels), [3,1,3]x5 Opportunity (20%), [3,2,1]x1 Dirty Deeds (10 Energy, Garrote from the front, at 1/2), [3,2,2]x1 Camouflage at 19, Premeditation at 20.
- items.json: Worn Dagger and Small Throwing Knife exist.

## Sources

- Warcraft Tavern, Rogue Leveling Guide (Forever): https://www.warcrafttavern.com/forever/guides/rogue-leveling/ (training levels match the data; it lists Dual Wield at 10 and the level 20 poison quests "Mathias and the Defias" / "The Shattered Salute", left out as unverifiable in 70755 and above 19)
- wow.gg, WoW Forever Combat Rogue (beta): https://wow.gg/guides/wow-forever-combat-rogue-beta (Stealth into Ambush opener at low level)
- Icy Veins and Mobalytics Forever rogue leveling pages returned 403.
