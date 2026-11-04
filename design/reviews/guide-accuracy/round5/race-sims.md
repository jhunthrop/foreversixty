# Round 5 racial sims (client 1.60.1.70755)

## Shaman Restoration (healer fight, effective HPS, seed 1 / seed 50001, guide build, raid band 60)
```
shaman-restoration dwarf 719.8 719.8
shaman-restoration tauren 718.8 719.0
shaman-restoration troll 721.2 720.9
shaman-restoration orc 723.0 723.1
```
Client evidence: Blood Fury 20572 reads "Increases Attack Power and Spell Power by 10% for 15 sec." Its spell power effect is aura 317 with misc value 0, no school mask and no damage-or-healing selector, so it raises healing too (engine sim/core/racial_blood_fury.go at d1d9f2775 reads the same row). Orc leads every Horde race on both seeds, by about 2 HPS over Troll and 4 over Tauren, outside the seed spread of 0.3.

Decision (integrator): Restoration shaman recommendedRaces becomes [dwarf, orc]. The race prose names Blood Fury as the Horde healing pick (10% spell power, which covers your heals, for 15 seconds every 2 minutes), keeps Tauren's Endurance as the safer pick when you are the one taking damage, and Troll's Berserking after them. Dwarf stays the Alliance pick.

## Paladin Retribution, Human against Dwarf on the re-ranked gear (review pass)
```
human 1 780.79 0.52
human 50001 780.26 0.52
dwarf 1 783.39 0.52
dwarf 50001 782.97 0.52
lvlbuild-raid human 1 612.03 0.33
lvlbuild-bare human 1 283.14 0.15
lvlbuild-raid human 50001 611.88 0.33
lvlbuild-bare human 50001 282.99 0.16
lvlbuild-raid dwarf 1 612.56 0.33
lvlbuild-bare dwarf 1 280.02 0.16
lvlbuild-raid dwarf 50001 612.15 0.32
lvlbuild-bare dwarf 50001 279.92 0.16
```
Raid build (the first four lines): Dwarf 783.2 against Human 780.5. Leveling build in a raid: even. Solo (bare): Human 283.1 against Dwarf 280.0. Human stays the pick, and the Dwarf line ('even in a raid') holds.
