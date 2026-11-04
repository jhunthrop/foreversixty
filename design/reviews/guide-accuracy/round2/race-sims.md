# Round 2 racial sims (client 1.60.1.70755, engine d1d9f2775)

sim/cmd/spec-breakdown, band 60, published gear and talents kept (`-race` swaps only the race and its rank choices), 10,000 iterations, seeds 1 and 50001 averaged; error is the larger of the two runs' standard errors. Run one at a time, niced, 2026-11-04. A race whose racial needs a weapon type (Dwarf guns, Troll bows, Orc axes) is measured on the published weapons, so its number is what that gear gives it.

| Spec | Preset | Faction | Race | DPS | Error |
|---|---|---|---|---|---|
| hunter-beast-mastery | bare | alliance | dwarf | 329.69 | 0.12 |
| hunter-beast-mastery | bare | alliance | high-order-skyborne | 330.87 | 0.12 |
| hunter-beast-mastery | bare | alliance | human | 332.47 | 0.12 |
| hunter-beast-mastery | bare | alliance | night-elf | 333.37 | 0.12 |
| hunter-beast-mastery | bare | horde | orc | 332.99 | 0.12 |
| hunter-beast-mastery | bare | horde | tauren | 325.84 | 0.12 |
| hunter-beast-mastery | bare | horde | troll | 329.51 | 0.12 |
| hunter-beast-mastery | bare | horde | windshaper-skyborne | 330.31 | 0.12 |
| hunter-beast-mastery | raid | alliance | dwarf | 658.98 | 0.22 |
| hunter-beast-mastery | raid | alliance | high-order-skyborne | 660.96 | 0.22 |
| hunter-beast-mastery | raid | alliance | human | 660.75 | 0.23 |
| hunter-beast-mastery | raid | alliance | night-elf | 664.14 | 0.23 |
| hunter-beast-mastery | raid | horde | orc | 668.54 | 0.23 |
| hunter-beast-mastery | raid | horde | tauren | 655.73 | 0.22 |
| hunter-beast-mastery | raid | horde | troll | 663.39 | 0.22 |
| hunter-beast-mastery | raid | horde | windshaper-skyborne | 663.65 | 0.22 |
| hunter-marksmanship | bare | alliance | dwarf | 330.62 | 0.13 |
| hunter-marksmanship | bare | alliance | high-order-skyborne | 332.26 | 0.14 |
| hunter-marksmanship | bare | alliance | human | 332.76 | 0.14 |
| hunter-marksmanship | bare | alliance | night-elf | 334.31 | 0.14 |
| hunter-marksmanship | bare | horde | orc | 334.44 | 0.13 |
| hunter-marksmanship | bare | horde | tauren | 328.27 | 0.13 |
| hunter-marksmanship | bare | horde | troll | 332.62 | 0.13 |
| hunter-marksmanship | bare | horde | windshaper-skyborne | 333.24 | 0.13 |
| hunter-marksmanship | raid | alliance | dwarf | 703.68 | 0.27 |
| hunter-marksmanship | raid | alliance | high-order-skyborne | 706.56 | 0.28 |
| hunter-marksmanship | raid | alliance | human | 706.01 | 0.28 |
| hunter-marksmanship | raid | alliance | night-elf | 709.29 | 0.28 |
| hunter-marksmanship | raid | horde | orc | 711.01 | 0.28 |
| hunter-marksmanship | raid | horde | tauren | 699.40 | 0.27 |
| hunter-marksmanship | raid | horde | troll | 707.88 | 0.28 |
| hunter-marksmanship | raid | horde | windshaper-skyborne | 708.97 | 0.28 |
| hunter-survival | bare | alliance | dwarf | 300.18 | 0.15 |
| hunter-survival | bare | alliance | high-order-skyborne | 300.84 | 0.15 |
| hunter-survival | bare | alliance | human | 301.12 | 0.15 |
| hunter-survival | bare | alliance | night-elf | 302.95 | 0.16 |
| hunter-survival | bare | horde | orc | 296.69 | 0.15 |
| hunter-survival | bare | horde | tauren | 293.75 | 0.15 |
| hunter-survival | bare | horde | troll | 296.84 | 0.15 |
| hunter-survival | bare | horde | windshaper-skyborne | 297.10 | 0.15 |
| hunter-survival | raid | alliance | dwarf | 747.66 | 0.38 |
| hunter-survival | raid | alliance | high-order-skyborne | 749.91 | 0.38 |
| hunter-survival | raid | alliance | human | 751.75 | 0.39 |
| hunter-survival | raid | alliance | night-elf | 755.13 | 0.39 |
| hunter-survival | raid | horde | orc | 753.53 | 0.38 |
| hunter-survival | raid | horde | tauren | 744.08 | 0.38 |
| hunter-survival | raid | horde | troll | 751.37 | 0.38 |
| hunter-survival | raid | horde | windshaper-skyborne | 750.62 | 0.38 |
| paladin-retribution | bare | alliance | dwarf | 261.03 | 0.15 |
| paladin-retribution | bare | alliance | human | 264.17 | 0.15 |
| paladin-retribution | raid | alliance | dwarf | 761.75 | 0.50 |
| paladin-retribution | raid | alliance | human | 761.98 | 0.51 |
| warlock-affliction | bare | alliance | gnome | 305.86 | 0.10 |
| warlock-affliction | bare | alliance | human | 305.63 | 0.10 |
| warlock-affliction | raid | alliance | gnome | 570.04 | 0.29 |
| warlock-affliction | raid | alliance | human | 573.58 | 0.29 |

## Warrior (raid preset, same method)

| Spec | Faction | Race | DPS | Error |
|---|---|---|---|---|
| warrior-arms | alliance | high-order-skyborne | 754.14 | 0.45 |
| warrior-arms | alliance | night-elf | 750.54 | 0.46 |
| warrior-arms | alliance | gnome | 749.44 | 0.47 |
| warrior-arms | alliance | human | 748.81 | 0.46 |
| warrior-arms | alliance | dwarf | 745.18 | 0.46 |
| warrior-arms | horde | windshaper-skyborne | 764.63 | 0.46 |
| warrior-arms | horde | undead | 762.61 | 0.46 |
| warrior-arms | horde | orc | 761.33 | 0.47 |
| warrior-arms | horde | troll | 757.86 | 0.46 |
| warrior-arms | horde | tauren | 753.85 | 0.46 |

warrior-fury could not run: its published BiS entry (data/builds/1.60.1.70755/bis/warrior-fury.json) carries the talent string 34320003002000000-051531025020101031-200000000000000000, which is illegal on 70755 (Bloodthirst needs 30 Fury points above its row; the string has 29). For the BiS refresh, not this round.

## Rerun on the GUIDE raid builds (these decide)

The tables above ran each spec on its published BiS entry's talents, which for several specs differ from the guide's raid build. These runs pass the guide's own raidBuild (or build, where a spec has one) with -talents; same presets, iterations and seeds. Values are seed 1 | seed 50001.

```
warlock-affliction alliance human 577.52 0.29 | 577.95 0.29
warlock-affliction alliance gnome 574.02 0.29 | 574.30 0.30
warlock-demonology alliance human 637.10 0.38 | 637.47 0.38
warlock-demonology alliance gnome 635.29 0.38 | 635.27 0.38
warlock-destruction alliance human 684.26 0.34 | 683.88 0.33
warlock-destruction alliance gnome 680.92 0.33 | 680.74 0.33
paladin-retribution alliance human 761.98 0.51 | 761.98 0.51
paladin-retribution alliance dwarf 761.57 0.50 | 761.93 0.50
warrior-arms alliance high-order-skyborne 755.31 0.45 | 755.11 0.45
warrior-arms alliance night-elf 752.13 0.46 | 751.77 0.45
warrior-arms alliance gnome 750.50 0.47 | 750.87 0.47
warrior-arms alliance human 750.70 0.46 | 750.66 0.45
warrior-arms alliance dwarf 745.76 0.45 | 745.86 0.45
warrior-arms horde windshaper-skyborne 766.54 0.45 | 765.19 0.45
warrior-arms horde undead 765.10 0.46 | 764.42 0.46
warrior-arms horde orc 763.29 0.47 | 763.18 0.46
warrior-arms horde troll 759.14 0.45 | 758.71 0.46
warrior-fury alliance high-order-skyborne 832.78 0.45 | 832.53 0.44
warrior-fury alliance night-elf 830.70 0.44 | 831.87 0.44
warrior-fury alliance gnome 829.35 0.45 | 829.16 0.45
warrior-fury alliance human 830.07 0.44 | 829.12 0.44
warrior-fury alliance dwarf 826.77 0.44 | 827.19 0.44
warrior-fury horde windshaper-skyborne 833.81 0.42 | 834.87 0.42
warrior-fury horde undead 835.61 0.41 | 834.96 0.41
warrior-fury horde orc 833.38 0.42 | 832.53 0.42
warrior-fury horde troll 831.02 0.42 | 831.58 0.42

```

Hunters on the guide builds (seed 1 | seed 50001):
```
beast-mastery night-elf 674.81 | 674.91   dwarf 669.52 | 669.09   human 671.25 | 671.58   orc 679.00 | 678.83   troll 673.65 | 673.30
marksmanship  night-elf 745.93 | 745.94   dwarf 739.95 | 739.95   human 742.45 | 742.48   orc 743.60 | 743.44   troll 745.51 | 745.66
survival      night-elf 820.44 | 819.41   dwarf 812.62 | 812.33   human 816.95 | 815.93   orc 818.80 | 818.27   troll 818.59 | 818.99
```
Warlock builds (guide raid strings converted with the class tree order): Destruction published 680.83, Molten Skin and Destructive Reach into Aftermath 5 + Unholy Power 4 688.71, into Pandemic 3 + Malevolence 3 692.68 (errors 0.33-0.35); Affliction published 574.16, Nightfall 2 for Grim Reach 2 578.20 (error 0.29). Survival Rapid Killing 2 for Lethal Attacks 2 is not a legal build (Careful Aim needs 5 Marksmanship points above its row).
