# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 2.4s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.101, intellect=0.342 ± 0.017, crit=0.050 ± 0.002 per rating point (14 rating = 1%, 0.695 per %), hit=0.189 ± 0.001 per rating point (10 rating = 1%, 1.895 per %), spell_haste=not significant (0.091 ± 0.109), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.101, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.69 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.79 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.39 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.37 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.7 spell_power points (0.66 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.27 DPS) [crafted]; Gray Woolen Robe (2585, -1.09 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.7 spell_power points (0.17 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.07 DPS) [quest]; Repurposed Hair Band (281256, -0.10 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.42 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.53 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (36.0 DPS) | yes | Silk-threaded Trousers (1929, -0.10 DPS) [dungeon]; Rumpled Kilt (274741, -0.30 DPS) [vendor]; Abomination Skin Leggings (23173, -0.75 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.43 DPS) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Feather Padded Treads (285345, -0.50 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.56 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.46 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.49 DPS) | yes | Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Sludge-Stained Band (286535, -0.39 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.4 spell_power points (0.34 DPS) | yes | Channeler's Staff (4437, -0.07 DPS) [world]; Lesser Staff of the Spire (1300, -0.13 DPS) [world]; Staff of Westfall (2042, -0.17 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 230.1 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.29 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 55.7. Weights run: 2.4s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.094, intellect=0.161 ± 0.007, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.486 per %), hit=0.141 ± 0.001 per rating point (10 rating = 1%, 1.415 per %), spell_haste=0.953 ± 0.107, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.094, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.77 DPS) | yes | Embalmed Shroud (7691, -0.48 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.48 DPS) [crafted]; Silk Headband (7050, -0.54 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -1.18 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.18 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.51 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (1.68 DPS) | yes | Death Speaker Mantle (6685, -0.43 DPS) [dungeon]; Fairywing Mantle (9536, -0.48 DPS) [quest]; Invoker's Mantle (215365, -0.57 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.80 DPS) | yes | Prelacy Cape (7004, -0.16 DPS) [quest]; Caretaker's Cape (19533, -0.16 DPS) [rep]; Heavy Woolen Cloak (4311, -0.30 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.09 DPS) | yes | Green Silk Armor (7065, -0.31 DPS) [crafted]; Death Speaker Robes (6682, -0.68 DPS) [dungeon]; Pristine Gown (253961, -0.78 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.45 DPS) | yes | Windsong Bangles (263336, -1.29 DPS) [quest]; Nightsky Wristbands (6407, -1.29 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.51 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.13 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.16 DPS) [world]; Town Clerk's Mittens (270029, -0.20 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (1.85 DPS) | yes | Belt of Arugal (6392, -0.54 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.56 DPS) [dungeon]; Invoker's Cord (215366, -0.59 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.93 DPS) | yes | Abomination Skin Leggings (23173, -0.28 DPS) [dungeon]; Pristine Leggings (253987, -0.62 DPS) [crafted]; Silk-threaded Trousers (1929, -0.80 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (1.31 DPS) | yes | Acidic Walkers (9454, -0.30 DPS) [dungeon]; Nimbus Boots (6998, -0.34 DPS) [quest]; Spidersilk Boots (4320, -1.58 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.13 DPS) | yes | Minor Channeling Ring (1449, -0.27 DPS) [quest]; Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon]; Sludge-Stained Band (286535, -0.64 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.96 DPS) | yes | Electrocutioner Lagnut (9447, -0.48 DPS) [dungeon]; Sludge-Stained Band (286535, -0.48 DPS) [world]; Minor Channeling Ring (1449, -1.58 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.45 DPS) | yes | Twisted Chanter's Staff (890, -1.19 DPS) [world_drop]; Channeler's Staff (4437, -1.24 DPS) [world]; Glimmering Staff (249392, -1.42 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (1.28 DPS) | yes | Eye of Paleth (2943, -0.64 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.64 DPS) [world]; Dwarven Tome (279898, -1.00 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 209.5 spell_power points (33.69 DPS) | yes | Starfaller (13063, -0.73 DPS) [world_drop]; Greater Mystic Wand (217287, -3.89 DPS) [crafted]; Gravestone Scepter (7001, -4.69 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 111.8. Weights run: 1.8s. Verify run: 1.1s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.246, intellect=0.770 ± 0.016, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.004 per %), hit=0.313 ± 0.002 per rating point (10 rating = 1%, 3.132 per %), spell_haste=1.746 ± 0.268, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.246, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.09 DPS) | yes | Corpseshroud (10574, -0.64 DPS) [dungeon]; Enchanter's Cowl (4322, -0.73 DPS) [crafted]; Augural Shroud (2620, -2.22 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.6 spell_power points (1.16 DPS) | yes | Triune Amulet (7722, -0.62 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.62 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.66 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.0 spell_power points (1.70 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.9 spell_power points (1.59 DPS) | yes | Guardian Cloak (5965, -0.61 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.74 DPS) [vendor]; Long Silken Cloak (4326, -1.95 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.6 spell_power points (2.66 DPS) | yes | Robe of Power (7054, -0.34 DPS) [crafted]; Elemental Raiment (9434, -0.56 DPS) [world_drop]; Dreamweave Vest (10021, -0.97 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (0.90 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.20 DPS) [quest]; Windchaser Cuffs (14429, -0.21 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.10 DPS) | yes | Black Mageweave Gloves (10003, -0.61 DPS) [crafted]; Gilded Handwraps (254021, -0.77 DPS) [crafted]; Red Mageweave Gloves (10018, -1.47 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (111.8 DPS) | yes | Gilded Cord (254037, -0.29 DPS) [crafted]; Star Belt (4329, -0.41 DPS) [crafted]; Deathmage Sash (10771, -1.40 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.2 spell_power points (2.32 DPS) | yes | Abomination Skin Leggings (23173, -0.81 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.01 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.37 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.39 DPS) | yes | Acidic Walkers (9454, -1.28 DPS) [dungeon]; Spidersilk Boots (4320, -1.39 DPS) [crafted]; Gilded Slippers (254001, -2.25 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.46 DPS) | yes | Ring of Forlorn Spirits (2043, -0.66 DPS) [quest]; Reedknot Ring (9622, -0.76 DPS) [quest]; Minor Channeling Ring (1449, -0.81 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.90 DPS) | yes | Ring of Forlorn Spirits (2043, -0.10 DPS) [quest]; Reedknot Ring (9622, -0.20 DPS) [quest]; Minor Channeling Ring (1449, -0.25 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Dar'Orahil (15106, -0.84 DPS) [quest]; Windweaver Staff (7757, -0.84 DPS) [dungeon]; Gut Ripper (2164, -6.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 399.4 spell_power points (39.84 DPS) | yes | Umbral Wand (5216, -4.17 DPS) [dungeon]; Earthen Rod (9381, -4.25 DPS) [dungeon]; Twisted Nether Wand (249144, -5.24 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 153.1. Weights run: 1.8s. Verify run: 1.2s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.159, intellect=0.281 ± 0.010, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.602 per %), hit=0.161 ± 0.002 per rating point (10 rating = 1%, 1.608 per %), spell_haste=not significant (0.396 ± 0.180), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.943 ± 0.159, fire_power=0.059 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.98 DPS) | yes | Dreamweave Circlet (10041, -0.83 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.55 DPS) [crafted]; Red Mageweave Headband (10033, -1.82 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (2.25 DPS) | yes | Mindburst Medallion (11196, -0.26 DPS) [quest]; Horizon Choker (13085, -1.23 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.52 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.1 spell_power points (4.67 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.43 DPS) [crafted]; Bloodmage Mantle (7684, -1.69 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.7 spell_power points (4.06 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.08 DPS) [dungeon]; Runecloth Cloak (13860, -1.15 DPS) [crafted]; Nightfall Drape (12465, -1.73 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.6 spell_power points (6.37 DPS) | yes | Robe of the Magi (1716, -0.24 DPS) [world_drop]; Elemental Raiment (9434, -0.93 DPS) [world_drop]; Dreamweave Vest (10021, -1.06 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.33 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.01 DPS) [crafted]; Bloodband Bracers (11469, -0.38 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (4.95 DPS) | yes | Black Mageweave Gloves (10003, -1.07 DPS) [crafted]; Runecloth Gloves (13863, -1.19 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.73 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.8 spell_power points (4.35 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.44 DPS) [rep]; Ghostweave Cord (254073, -0.73 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.8 spell_power points (6.68 DPS) | yes | Red Mageweave Pants (10009, -2.18 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.54 DPS) [vendor]; Wizardweave Leggings (14132, -3.76 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.21 DPS) | yes | Gilded Sandals (254107, -1.36 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -2.85 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.07 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.36 DPS) | yes | Philanthropist's Ring (281635, -0.34 DPS) [quest]; Cyclopean Band (11824, -0.53 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.29 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.10 DPS) | yes | Cyclopean Band (11824, -0.27 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.03 DPS) [quest]; Philanthropist's Ring (281635, -1.04 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.92 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.74 DPS) [dungeon]; Scorn's Focal Dagger (23168, -2.85 DPS) [dungeon]; Blade of Eternal Darkness (17780, -5.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (153.1 DPS) | yes | Pyric Caduceus (11748, -2.27 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.87 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.47 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 283.3. Weights run: 5.1s. Verify run: 1.1s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.382, intellect=0.122 ± 0.024, crit=0.120 ± 0.005 per rating point (14 rating = 1%, 1.680 per %), hit=0.378 ± 0.004 per rating point (10 rating = 1%, 3.778 per %), spell_haste=not significant (-0.175 ± 0.442), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.919 ± 0.382), fire_power=0.081 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.0 spell_power points (5.92 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.17 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.27 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (4.20 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.04 DPS) [quest]; Kezan's Taint (19604, -1.34 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.5 spell_power points (5.83 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.66 DPS) [pvp]; Burial Shawl (18681, -1.64 DPS) [dungeon]; Argent Shoulders (19059, -3.16 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.8 spell_power points (3.97 DPS) | yes | Crystalline Threaded Cape (20697, -0.05 DPS) [world]; Amplifying Cloak (18350, -0.53 DPS) [dungeon]; Hide of the Wild (18510, -1.06 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (9.00 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -2.33 DPS) [pvp]; Robe of Everlasting Night (18385, -2.42 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -3.76 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.0 spell_power points (4.39 DPS) | yes | Sublime Wristguards (18497, -1.86 DPS) [dungeon]; Runecloth Cuffs (254123, -2.06 DPS) [crafted]; Arcane Runed Bracers (4744, -2.67 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (5.28 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Hands of Power (13253, -0.17 DPS) [dungeon]; Earth Warder's Gloves (21318, -1.07 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 31.5 spell_power points (6.01 DPS) | yes | Stormpike Cloth Girdle (19094, -2.34 DPS) [rep]; Highlander's Cloth Girdle (20047, -2.88 DPS) [rep]; Belt of the Archmage (18405, -4.10 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 36.6 spell_power points (6.99 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -1.34 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.59 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.38 DPS) [crafted]; Omnicast Boots (11822, -0.49 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (283.3 DPS) | yes | Songstone of Ironforge (12543, -0.86 DPS) [quest]; Maiden's Circle (13001, -0.86 DPS) [world_drop]; Naglering (11669, -7.05 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (283.3 DPS) | yes | Songstone of Ironforge (12543, -0.57 DPS) [quest]; Maiden's Circle (13001, -0.57 DPS) [world_drop]; Naglering (11669, -6.95 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (283.3 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (283.3 DPS) | yes | Weakness Analyzer (272438, -1.34 DPS) [vendor]; Serenity Field (272439, -2.87 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.90 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (283.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.67 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -14.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 407.6 spell_power points (77.91 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.95 DPS) [dungeon]; Bonecreeper Stylus (13938, -13.08 DPS) [dungeon]; Wand of Biting Cold (19108, -13.91 DPS) [quest] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 34.9. Weights run: 2.4s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.101, intellect=0.342 ± 0.017, crit=0.050 ± 0.002 per rating point (14 rating = 1%, 0.695 per %), hit=0.189 ± 0.001 per rating point (10 rating = 1%, 1.895 per %), spell_haste=not significant (0.091 ± 0.109), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.101, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.72 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.79 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.39 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.43 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.7 spell_power points (0.66 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.27 DPS) [crafted]; Gray Woolen Robe (2585, -1.36 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest]; Owlbeard Bracers (16981, -0.04 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Pristine Gloves (253913, -0.19 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.29 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.53 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (34.9 DPS) | yes | Silk-threaded Trousers (1929, -0.10 DPS) [dungeon]; Rumpled Kilt (274741, -0.30 DPS) [vendor]; Abomination Skin Leggings (23173, -0.67 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.43 DPS) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Feather Padded Treads (285345, -0.67 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.49 DPS) | yes | Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Loop of Sacrifice (281673, -0.32 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.4 spell_power points (0.34 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.07 DPS) [world]; Lesser Staff of the Spire (1300, -0.13 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 230.1 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Sizzle Stick (8071, -4.46 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 54.1. Weights run: 2.4s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.094, intellect=0.161 ± 0.007, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.486 per %), hit=0.141 ± 0.001 per rating point (10 rating = 1%, 1.415 per %), spell_haste=0.953 ± 0.107, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.094, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.77 DPS) | yes | Silk Headband (7050, -0.42 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.48 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -1.18 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.18 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.38 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (1.68 DPS) | yes | Invoker's Mantle (215365, -0.42 DPS) [crafted]; Death Speaker Mantle (6685, -0.43 DPS) [dungeon]; Chestnut Mantle (17695, -1.28 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.80 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.16 DPS) [crafted]; Battle Healer's Cloak (19529, -0.16 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.09 DPS) | yes | Green Silk Armor (7065, -0.31 DPS) [crafted]; Death Speaker Robes (6682, -0.68 DPS) [dungeon]; High Robe of the Adjudicator (3461, -0.75 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.45 DPS) | yes | Glowing Magical Bracelets (13106, -1.24 DPS) [world_drop]; Windsong Bangles (263336, -1.29 DPS) [quest]; Owlbeard Bracers (16981, -1.48 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.13 DPS) | yes | Jutebraid Gloves (10654, -0.03 DPS) [quest]; Gnoll Casting Gloves (892, -0.16 DPS) [world]; Truefaith Gloves (7049, -0.24 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (1.85 DPS) | yes | Warsong Sash (16975, -0.28 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.32 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.56 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.93 DPS) | yes | Abomination Skin Leggings (23173, -0.28 DPS) [dungeon]; Pristine Leggings (253987, -0.62 DPS) [crafted]; Silk-threaded Trousers (1929, -0.80 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (1.31 DPS) | yes | Acidic Walkers (9454, -0.30 DPS) [dungeon]; Boots of the Enchanter (4325, -0.50 DPS) [crafted]; Spidersilk Boots (4320, -1.75 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.13 DPS) | yes | Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon]; Sludge-Stained Band (286535, -0.64 DPS) [world]; Sacred Band (6669, -0.80 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.96 DPS) | yes | Electrocutioner Lagnut (9447, -0.48 DPS) [dungeon]; Sacred Band (6669, -0.64 DPS) [quest]; Sludge-Stained Band (286535, -1.65 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.45 DPS) | yes | Twisted Chanter's Staff (890, -1.19 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.19 DPS) [quest]; Glimmering Staff (249392, -1.25 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (1.28 DPS) | yes | Orb of Souls (249395, -0.64 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.86 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.68 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 209.5 spell_power points (33.69 DPS) | yes | Starfaller (13063, -0.59 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.89 DPS) [crafted]; Gravestone Scepter (7001, -4.69 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 108.8. Weights run: 1.8s. Verify run: 1.1s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.246, intellect=0.770 ± 0.016, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.004 per %), hit=0.313 ± 0.002 per rating point (10 rating = 1%, 3.132 per %), spell_haste=1.746 ± 0.268, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.246, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.09 DPS) | yes | Corpseshroud (10574, -0.64 DPS) [dungeon]; Enchanter's Cowl (4322, -0.73 DPS) [crafted]; Augural Shroud (2620, -1.49 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.6 spell_power points (1.16 DPS) | yes | Triune Amulet (7722, -0.62 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.62 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.76 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.0 spell_power points (1.70 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.9 spell_power points (1.59 DPS) | yes | Guardian Cloak (5965, -0.61 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.74 DPS) [vendor]; Long Silken Cloak (4326, -1.60 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.6 spell_power points (2.66 DPS) | yes | Robe of Power (7054, -0.34 DPS) [crafted]; Elemental Raiment (9434, -0.56 DPS) [world_drop]; Dreamweave Vest (10021, -0.93 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (108.8 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.20 DPS) [quest]; Radiant Silver Bracers (4545, -1.26 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.10 DPS) | yes | Black Mageweave Gloves (10003, -0.61 DPS) [crafted]; Gilded Handwraps (254021, -0.77 DPS) [crafted]; Red Mageweave Gloves (10018, -1.40 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.5 spell_power points (1.85 DPS) | yes | Defiler's Cloth Girdle (20166, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.44 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.2 spell_power points (2.32 DPS) | yes | Abomination Skin Leggings (23173, -0.81 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.01 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.62 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.39 DPS) | yes | Acidic Walkers (9454, -1.28 DPS) [dungeon]; Spidersilk Boots (4320, -1.39 DPS) [crafted]; Gilded Slippers (254001, -2.00 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.46 DPS) | yes | Reedknot Ring (9622, -0.76 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.86 DPS) [vendor]; Black Widow Band (6199, -0.92 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.90 DPS) | yes | Sea Giant's Toe Ring (274746, -0.30 DPS) [vendor]; Black Widow Band (6199, -0.36 DPS) [world]; Reedknot Ring (9622, -0.55 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Dar'Orahil (15106, -0.84 DPS) [quest]; Windweaver Staff (7757, -0.84 DPS) [dungeon]; Gut Ripper (2164, -6.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 399.4 spell_power points (39.84 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.25 DPS) [dungeon]; Twisted Nether Wand (249144, -5.24 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 151.7. Weights run: 1.8s. Verify run: 1.2s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.159, intellect=0.281 ± 0.010, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.602 per %), hit=0.161 ± 0.002 per rating point (10 rating = 1%, 1.608 per %), spell_haste=not significant (0.396 ± 0.180), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.943 ± 0.159, fire_power=0.059 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.98 DPS) | yes | Dreamweave Circlet (10041, -0.83 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.55 DPS) [crafted]; Red Mageweave Headband (10033, -1.61 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (2.25 DPS) | yes | Mindburst Medallion (11196, -0.26 DPS) [quest]; Horizon Choker (13085, -1.23 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.52 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.1 spell_power points (4.67 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.43 DPS) [crafted]; Bloodmage Mantle (7684, -1.69 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.7 spell_power points (4.06 DPS) | yes | Deep Woodlands Cloak (19121, -0.70 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.08 DPS) [dungeon]; Runecloth Cloak (13860, -1.15 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.6 spell_power points (6.37 DPS) | yes | Robe of the Magi (1716, -0.24 DPS) [world_drop]; Elemental Raiment (9434, -0.93 DPS) [world_drop]; Dreamweave Vest (10021, -1.06 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.33 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (4.95 DPS) | yes | Black Mageweave Gloves (10003, -1.07 DPS) [crafted]; Runecloth Gloves (13863, -1.19 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.30 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Cloth Girdle (20166, -0.03 DPS) [rep]; Ghostweave Cord (254073, -0.32 DPS) [crafted]; Satyrmane Sash (17755, -1.95 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.8 spell_power points (6.68 DPS) | yes | Red Mageweave Pants (10009, -2.18 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.54 DPS) [vendor]; Wizardweave Leggings (14132, -3.21 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.21 DPS) | yes | Gilded Sandals (254107, -0.52 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -2.85 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.07 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.36 DPS) | yes | Philanthropist's Ring (281635, -0.34 DPS) [quest]; Cyclopean Band (11824, -0.53 DPS) [dungeon]; Runed Ring (862, -1.55 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.10 DPS) | yes | Cyclopean Band (11824, -0.27 DPS) [dungeon]; Philanthropist's Ring (281635, -0.62 DPS, sim-verified) [quest]; Runed Ring (862, -1.29 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.69 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -2.81 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.08 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.74 DPS) [dungeon]; Scorn's Focal Dagger (23168, -2.85 DPS) [dungeon]; Blade of Eternal Darkness (17780, -4.30 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Pyric Caduceus (11748, -2.17 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.87 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.47 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 277.5. Weights run: 5.1s. Verify run: 1.2s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.382, intellect=0.122 ± 0.024, crit=0.120 ± 0.005 per rating point (14 rating = 1%, 1.680 per %), hit=0.378 ± 0.004 per rating point (10 rating = 1%, 3.778 per %), spell_haste=not significant (-0.175 ± 0.442), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.919 ± 0.382), fire_power=0.081 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.0 spell_power points (5.92 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.17 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.81 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (4.20 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.04 DPS) [quest]; Kezan's Taint (19604, -1.34 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.5 spell_power points (5.83 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.66 DPS) [pvp]; Burial Shawl (18681, -1.64 DPS) [dungeon]; Argent Shoulders (19059, -3.12 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.8 spell_power points (3.97 DPS) | yes | Amplifying Cloak (18350, -0.53 DPS) [dungeon]; Hide of the Wild (18510, -1.06 DPS) [crafted]; Crystalline Threaded Cape (20697, -2.00 DPS, sim-verified) [world] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (9.00 DPS) | yes | Warlord's Dreadweave Robe (231591, -2.33 DPS) [pvp]; Robe of Everlasting Night (18385, -3.10 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -3.76 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.0 spell_power points (4.39 DPS) | yes | Sublime Wristguards (18497, -1.86 DPS) [dungeon]; Runecloth Cuffs (254123, -2.06 DPS) [crafted]; Spidertank Oilrag (9448, -2.67 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (5.28 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -0.17 DPS) [dungeon]; Earth Warder's Gloves (21318, -1.07 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 31.5 spell_power points (6.01 DPS) | yes | Frostwolf Cloth Belt (19090, -2.34 DPS) [rep]; Defiler's Cloth Girdle (20163, -2.88 DPS) [rep]; Belt of the Archmage (18405, -5.30 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 36.6 spell_power points (6.99 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.20 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.59 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.38 DPS) [crafted]; Omnicast Boots (11822, -0.49 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (277.5 DPS) | yes | Eye of Orgrimmar (12545, -0.86 DPS) [quest]; Maiden's Circle (13001, -0.86 DPS) [world_drop]; Naglering (11669, -7.96 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (277.5 DPS) | yes | Eye of Orgrimmar (12545, -0.57 DPS) [quest]; Maiden's Circle (13001, -0.57 DPS) [world_drop]; Naglering (11669, -8.90 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (277.5 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (277.5 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (277.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.67 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -15.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 407.6 spell_power points (77.91 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -2.64 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -13.08 DPS) [dungeon]; Wand of Biting Cold (19108, -13.91 DPS) [quest] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

