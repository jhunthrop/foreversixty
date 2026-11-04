# Round 4 racial sims (client 1.60.1.70755)

Paladin Protection on the tank fight ([damage taken/s, TMI, threat/s], seed 1 / seed 50001) and Holy on the healer fight (effective HPS), guide builds, raid band 60, 10,000 iterations.
```
paladin-protection dwarf [362.7 41.1 1113.8 ] [361.9 41.1 1112.5 ]
paladin-protection human [361.4 41.2 1110.3 ] [360.5 41.2 1109.4 ]
paladin-protection undead [367.5 41.0 1125.5 ] [367.0 40.9 1125.2 ]
paladin-holy dwarf [660.9] [661.0]
paladin-holy human [676.5] [677.0]
paladin-holy undead [662.0] [662.4]
```
Decision: Protection keeps Dwarf for the Alliance (a tie with Human on damage taken and TMI; Stoneform is not cast by the engine, so Dwarf's emergency button is unmeasured) and Undead stays the Horde race (lowest TMI and the most threat). Holy keeps Human (676.8 against Undead 662.2 and Dwarf 661.0).

## Warlock Horde (guide raid builds, raid band 60, seed 1 | seed 50001)
```
warlock-affliction orc 582.10 0.29 | 582.89 0.29
warlock-affliction undead 581.84 0.29 | 582.27 0.29
warlock-affliction troll 575.79 0.29 | 575.81 0.29
warlock-demonology orc 644.64 0.37 | 645.22 0.38
warlock-demonology undead 643.19 0.37 | 642.40 0.37
warlock-demonology troll 639.39 0.37 | 640.23 0.38
warlock-destruction orc 703.26 0.35 | 703.63 0.35
warlock-destruction undead 701.77 0.34 | 701.36 0.34
warlock-destruction troll 695.88 0.34 | 695.56 0.34
```
Decision: Orc stays the Horde pick for all three (Affliction Orc 582.5 against Undead 582.1, a tie; Demonology 644.9 against 642.8 and Destruction 703.4 against 701.6, Orc beyond error; Troll last).

## Druid bear (tank fight, [damage taken/s, TMI, threat/s], seed 1 / seed 50001)
```
druid-feral-bear night-elf [465.8 37.3 964.2 ] [465.3 37.3 965.2 ]
druid-feral-bear tauren [473.5 35.1 955.7 ] [472.9 35.1 956.3 ]
```
Decision: druids are faction-locked (Night Elf on the Alliance, Tauren on the Horde; Skyborne is the paid option), so neither is 'the tank pick'. Night Elf takes about 1.6% less damage and holds more threat; Tauren's health is steadier (TMI 35.1 against 37.3).
