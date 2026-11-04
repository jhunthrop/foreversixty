# Round 3 racial sims (client 1.60.1.70755)

sim/cmd/spec-breakdown on each guide's raidBuild (or build) via -talents, raid preset, band 60, published gear, 10,000 iterations, seeds 1 and 50001 averaged, error the larger standard error. Every playable race of the class per data/builds/1.60.1.70755/combos.json. Skyborne needs the paid Heroic Pack: never the recommended pick. Run one at a time, niced, 2026-11-04.

```
druid-balance        alliance night-elf 569.9  high-order-skyborne 569.9  (±0.31)
druid-balance        horde    windshaper-skyborne 581.3  tauren 575.0  (±0.30)
druid-feral          alliance night-elf 675.5  high-order-skyborne 675.4  (±0.24)
druid-feral          horde    windshaper-skyborne 679.5  tauren 671.0  (±0.24)
mage-arcane          alliance gnome 614.0  human 611.8  high-order-skyborne 607.5  (±0.35)
mage-arcane          horde    undead 623.3  orc 619.3  troll 616.0  (±0.36)
mage-fire            alliance human 667.2  gnome 666.8  high-order-skyborne 655.9  (±0.44)
mage-fire            horde    undead 676.0  orc 673.3  troll 667.3  (±0.44)
mage-frost           alliance high-order-skyborne 630.9  gnome 627.3  human 625.9  (±0.51)
mage-frost           horde    undead 633.4  orc 632.4  troll 632.2  (±0.51)
priest-shadow        alliance human 690.7  night-elf 690.4  dwarf 689.6  gnome 687.8  (±0.27)
priest-shadow        horde    undead 695.0  troll 687.8  (±0.27)
rogue-assassination  alliance night-elf 746.1  dwarf 745.4  human 744.9  high-order-skyborne 744.7  gnome 732.3  (±0.35)
rogue-assassination  horde    undead 743.3  orc 739.4  windshaper-skyborne 739.2  troll 738.4  (±0.36)
rogue-combat         alliance gnome 759.3  high-order-skyborne 756.2  night-elf 755.9  human 755.1  dwarf 753.4  (±0.32)
rogue-combat         horde    undead 759.3  orc 756.6  windshaper-skyborne 756.0  troll 756.0  (±0.32)
rogue-subtlety       alliance night-elf 640.0  human 637.1  high-order-skyborne 636.9  gnome 635.2  dwarf 633.8  (±0.25)
rogue-subtlety       horde    undead 641.7  orc 638.1  windshaper-skyborne 636.9  troll 636.0  (±0.24)
shaman-elemental     alliance dwarf 623.2  (±0.30)
shaman-elemental     horde    windshaper-skyborne 633.4  orc 630.5  troll 624.8  tauren 624.6  (±0.30)
shaman-enhancement   alliance dwarf 795.6  (±0.51)
shaman-enhancement   horde    windshaper-skyborne 796.0  orc 795.6  troll 792.8  tauren 787.4  (±0.51)
```

Raw (seed 1 | seed 50001):
```
mage-arcane alliance human 612.04 0.35 | 611.61 0.35
mage-arcane horde orc 619.20 0.36 | 619.43 0.35
mage-arcane horde undead 622.92 0.35 | 623.60 0.35
mage-arcane alliance gnome 614.17 0.35 | 613.82 0.35
mage-arcane horde troll 616.40 0.35 | 615.51 0.35
mage-arcane alliance high-order-skyborne 607.53 0.34 | 607.45 0.34
mage-fire alliance human 666.84 0.44 | 667.66 0.44
mage-fire horde orc 673.02 0.44 | 673.54 0.43
mage-fire horde undead 675.49 0.43 | 676.42 0.44
mage-fire alliance gnome 666.42 0.44 | 667.12 0.44
mage-fire horde troll 666.95 0.43 | 667.66 0.43
mage-fire alliance high-order-skyborne 655.71 0.44 | 656.00 0.44
mage-frost alliance human 626.55 0.51 | 625.21 0.51
mage-frost horde orc 632.63 0.51 | 632.18 0.51
mage-frost horde undead 633.93 0.50 | 632.88 0.50
mage-frost alliance gnome 627.98 0.51 | 626.61 0.51
mage-frost horde troll 632.34 0.50 | 631.96 0.51
mage-frost alliance high-order-skyborne 631.44 0.50 | 630.28 0.50
rogue-assassination alliance human 744.59 0.35 | 745.28 0.35
rogue-assassination horde orc 739.30 0.35 | 739.47 0.36
rogue-assassination alliance dwarf 745.77 0.34 | 745.02 0.34
rogue-assassination alliance night-elf 745.74 0.35 | 746.41 0.35
rogue-assassination horde undead 743.02 0.35 | 743.62 0.35
rogue-assassination alliance gnome 732.44 0.34 | 732.16 0.34
rogue-assassination horde troll 738.68 0.35 | 738.04 0.35
rogue-assassination alliance high-order-skyborne 745.23 0.35 | 744.09 0.35
rogue-assassination horde windshaper-skyborne 739.32 0.36 | 739.17 0.35
rogue-combat alliance human 755.14 0.31 | 755.13 0.32
rogue-combat horde orc 756.22 0.32 | 756.93 0.31
rogue-combat alliance dwarf 753.27 0.31 | 753.60 0.32
rogue-combat alliance night-elf 755.98 0.31 | 755.85 0.32
rogue-combat horde undead 759.53 0.32 | 759.05 0.32
rogue-combat alliance gnome 759.38 0.32 | 759.19 0.32
rogue-combat horde troll 756.39 0.31 | 755.68 0.32
rogue-combat alliance high-order-skyborne 756.36 0.32 | 756.07 0.31
rogue-combat horde windshaper-skyborne 756.12 0.31 | 755.96 0.32
rogue-subtlety alliance human 637.36 0.24 | 636.75 0.24
rogue-subtlety horde orc 638.01 0.24 | 638.26 0.24
rogue-subtlety alliance dwarf 633.80 0.24 | 633.80 0.24
rogue-subtlety alliance night-elf 640.33 0.24 | 639.71 0.24
rogue-subtlety horde undead 641.47 0.24 | 641.91 0.24
rogue-subtlety alliance gnome 635.25 0.25 | 635.11 0.25
rogue-subtlety horde troll 635.80 0.24 | 636.15 0.24
rogue-subtlety alliance high-order-skyborne 636.55 0.24 | 637.32 0.24
rogue-subtlety horde windshaper-skyborne 636.60 0.24 | 637.29 0.24
shaman-elemental horde orc 630.37 0.30 | 630.61 0.30
shaman-elemental alliance dwarf 622.83 0.30 | 623.67 0.29
shaman-elemental horde tauren 624.43 0.30 | 624.70 0.30
shaman-elemental horde troll 624.75 0.30 | 624.88 0.30
shaman-elemental horde windshaper-skyborne 633.43 0.30 | 633.44 0.30
shaman-enhancement horde orc 794.52 0.50 | 796.66 0.51
shaman-enhancement alliance dwarf 794.61 0.51 | 796.50 0.50
shaman-enhancement horde tauren 786.41 0.50 | 788.45 0.50
shaman-enhancement horde troll 793.03 0.51 | 792.67 0.51
shaman-enhancement horde windshaper-skyborne 795.84 0.51 | 796.13 0.51
priest-shadow alliance human 690.30 0.27 | 691.10 0.27
priest-shadow alliance dwarf 689.45 0.26 | 689.78 0.27
priest-shadow alliance night-elf 690.14 0.27 | 690.68 0.27
priest-shadow horde undead 695.19 0.27 | 694.86 0.27
priest-shadow alliance gnome 687.46 0.27 | 688.13 0.27
priest-shadow horde troll 687.33 0.26 | 688.21 0.27
druid-balance alliance night-elf 569.50 0.31 | 570.32 0.31
druid-balance horde tauren 574.94 0.30 | 575.14 0.30
druid-balance alliance high-order-skyborne 569.79 0.31 | 569.93 0.31
druid-balance horde windshaper-skyborne 581.15 0.30 | 581.48 0.30
druid-feral alliance night-elf 675.50 0.24 | 675.41 0.24
druid-feral horde tauren 671.13 0.24 | 670.82 0.23
druid-feral alliance high-order-skyborne 675.28 0.24 | 675.44 0.24
druid-feral horde windshaper-skyborne 679.56 0.24 | 679.44 0.24
```

## Mage, bare 60 (guide raid builds, seed 1 | seed 50001)
```
mage-arcane bare gnome 351.20 0.22 | 351.23 0.22
mage-arcane bare human 350.11 0.22 | 350.72 0.22
mage-arcane bare undead 358.11 0.24 | 358.45 0.24
mage-arcane bare orc 353.71 0.24 | 353.57 0.24
mage-arcane bare troll 352.23 0.24 | 351.52 0.24
mage-fire bare gnome 342.54 0.33 | 343.04 0.32
mage-fire bare human 347.33 0.33 | 348.14 0.32
mage-fire bare undead 339.63 0.32 | 339.47 0.32
mage-fire bare orc 335.15 0.32 | 335.71 0.32
mage-fire bare troll 331.83 0.32 | 332.41 0.32
mage-frost bare gnome 365.32 0.29 | 365.39 0.29
mage-frost bare human 363.87 0.29 | 364.00 0.29
mage-frost bare undead 369.98 0.30 | 370.11 0.30
mage-frost bare orc 366.39 0.30 | 366.50 0.30
mage-frost bare troll 367.23 0.30 | 367.14 0.30
```
Decision: Arcane [gnome, undead]; Fire [human, undead] (raid tie, Human +1.4% bare, beyond error); Frost [gnome, undead] (Undead leads raid and bare; Skyborne is paid).

## Druid Restoration build (healer fight, 300 s, 10,000 iterations, seed 1)
Published 05312001/0/55530051531132: effective HPS 1014.5. Improved Moonfire point moved to Improved Regrowth (05302001/0/555300515311321): 1019.2 (+0.5%). NOT adopted: under the 1% bar, and it breaks the Restoration ladder-first invariant (see round3/druid.md).

## Priest healers (healer fight, effective HPS, seed 1 / seed 50001)
```
priest-discipline alliance human effHPS 597.9 598.0
priest-discipline alliance dwarf effHPS 581.8 581.9
priest-discipline alliance night-elf effHPS 585.1 585.0
priest-discipline alliance gnome effHPS 600.5 600.8
priest-discipline horde undead effHPS 585.2 585.2
priest-discipline horde troll effHPS 580.9 581.2
priest-holy alliance human effHPS 827.4 827.5
priest-holy alliance dwarf effHPS 805.8 805.5
priest-holy alliance night-elf effHPS 808.5 808.1
priest-holy alliance gnome effHPS 821.5 821.7
priest-holy horde undead effHPS 837.0 837.4
priest-holy horde troll effHPS 804.6 804.3
```
Decision: Discipline [gnome, undead] (Gnome 600.6 vs Human 598.0); Holy [human, undead] (Human 827.5 vs Gnome 821.6; Undead 837.2). Shadow [human, undead] from the DPS table (Human 690.7, Gnome last at 687.8; Undead 695.0 vs Troll 687.8).

## Protection warrior (tank fight, [damage taken/s, TMI, threat/s], seed 1 / seed 50001)
```
dwarf [262.4 25.8 1053.1 ] [262.3 25.8 1053.4 ]
human [261.6 25.9 1048.4 ] [261.9 25.9 1047.7 ]
night-elf [256.5 25.6 1049.0 ] [256.1 25.5 1048.3 ]
gnome [260.6 25.9 1050.0 ] [260.7 25.9 1047.7 ]
tauren [262.7 24.4 1042.7 ] [263.0 24.4 1041.3 ]
orc [262.3 25.9 1050.4 ] [262.2 25.8 1048.9 ]
undead [262.3 25.9 1063.0 ] [262.3 25.9 1062.3 ]
troll [261.4 25.8 1046.1 ] [261.4 25.8 1044.5 ]
```
Decision: Protection [night-elf, tauren]: Night Elf takes about 2% less damage than any Alliance race (and lower TMI than Dwarf); Tauren has the lowest TMI of all. Caveat: the engine never casts Stoneform (no rotation line), so Dwarf's emergency button is unmeasured; the prose names it as Dwarf's case.
