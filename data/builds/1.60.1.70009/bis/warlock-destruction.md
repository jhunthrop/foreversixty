# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.1. Weights run: 2.3s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.045 ± 0.003, crit=0.025 ± 0.001 per rating point (14 rating = 1%, 0.356 per %), hit=0.106 ± 0.003 per rating point (10 rating = 1%, 1.063 per %), spell_haste=0.510 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.761 ± 0.002, fire_power=0.239 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.25 DPS) | yes | Shadow Goggles (4373, -2.71 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.4 spell_power points (1.12 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.22 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.83 DPS) | yes | Feyscale Cloak (6632, -0.21 DPS) [dungeon]; Caretaker's Cape (20428, -0.21 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.2 spell_power points (1.08 DPS) | yes | Green Woolen Vest (2582, -0.25 DPS) [crafted]; Bloody Apron (6226, -0.25 DPS) [dungeon]; Gray Woolen Robe (2585, -1.38 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.21 DPS) | yes | Mindthrust Bracers (1974, -0.16 DPS) [dungeon]; Bright Bracers (3647, -0.17 DPS) [world_drop]; Repurposed Hair Band (281256, -0.19 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.45 DPS) | yes | Gnoll Casting Gloves (892, -0.28 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.60 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.02 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.2 spell_power points (0.87 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Keller's Girdle (2911, -0.79 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.95 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.4 spell_power points (1.94 DPS) | yes | Filigreed Pristine Leggings (253937, -0.64 DPS) [crafted]; Silk-threaded Trousers (1929, -0.69 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.2 spell_power points (1.49 DPS) | yes | Red Woolen Boots (4313, -0.66 DPS) [crafted]; Feather Padded Treads (285345, -0.67 DPS, sim-verified) [world]; Pristine Boots (253889, -0.84 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.1 spell_power points (1.06 DPS) | yes | Sludge-Stained Band (286535, -0.43 DPS) [world]; Lavishly Jeweled Ring (1156, -1.00 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.03 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.04 DPS) | yes | Sludge-Stained Band (286535, -0.57 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.98 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.01 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.4 spell_power points (0.09 DPS) | yes | Channeler's Staff (4437, -0.02 DPS) [world]; Lesser Staff of the Spire (1300, -0.04 DPS) [world]; Staff of Westfall (2042, -0.05 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 110.4 spell_power points (22.93 DPS) | yes | Skycaller (12984, -1.54 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.64 DPS) [dungeon]; Sizzle Stick (8071, -4.24 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053225101000000)

Set DPS (verified): 71.1. Weights run: 2.4s. Verify run: 1.3s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.290 ± 0.008, crit=0.065 ± 0.002 per rating point (14 rating = 1%, 0.911 per %), hit=0.156 ± 0.007 per rating point (10 rating = 1%, 1.564 per %), spell_haste=0.651 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.771 ± 0.002, fire_power=0.229 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.62 DPS) | yes | Enchanter's Cowl (4322, -0.50 DPS) [crafted]; Embalmed Shroud (7691, -0.72 DPS) [dungeon]; Silk Headband (7050, -1.20 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (2.08 DPS) | yes | Crystal Starfire Medallion (5003, -1.81 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.81 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.18 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.6 spell_power points (2.77 DPS) | yes | Fairywing Mantle (9536, -0.72 DPS) [quest]; Invoker's Mantle (215365, -0.75 DPS) [crafted]; Death Speaker Mantle (6685, -1.26 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.19 DPS) | yes | Heavy Woolen Cloak (4311, -0.24 DPS) [crafted]; Prelacy Cape (7004, -0.24 DPS) [quest]; Repairman's Cape (9605, -0.78 DPS, sim-verified) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.10 DPS) | yes | Death Speaker Robes (6682, -0.67 DPS) [dungeon]; Pristine Gown (253961, -0.95 DPS) [crafted]; Green Silk Armor (7065, -1.12 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.15 DPS) | yes | Nightsky Wristbands (6407, -1.73 DPS) [world_drop]; Stonecloth Bindings (14416, -1.80 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.48 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.2 spell_power points (1.72 DPS) | yes | Shilly Mitts (9609, +0.00 DPS, sim-verified) [quest]; Serpent Gloves (5970, -0.05 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.28 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 spell_power points (2.83 DPS) | yes | Belt of Arugal (6392, -0.75 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.82 DPS) [crafted]; Crimson Silk Belt (7055, -0.92 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.86 DPS) | yes | Pristine Leggings (253987, -0.71 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.02 DPS) [crafted]; Abomination Skin Leggings (23173, -1.19 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.0 spell_power points (2.15 DPS) | yes | Acidic Walkers (9454, -0.41 DPS) [dungeon]; Nimbus Boots (6998, -0.72 DPS) [quest]; Spidersilk Boots (4320, -2.45 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.67 DPS) | yes | Minor Channeling Ring (1449, -0.34 DPS) [quest]; Electrocutioner Lagnut (9447, -0.95 DPS) [dungeon]; Sludge-Stained Band (286535, -0.95 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.43 DPS) | yes | Electrocutioner Lagnut (9447, -0.72 DPS) [dungeon]; Sludge-Stained Band (286535, -0.72 DPS) [world]; Minor Channeling Ring (1449, -2.29 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.15 DPS) | yes | Twisted Chanter's Staff (890, -1.45 DPS) [world_drop]; Channeler's Staff (4437, -1.59 DPS) [world]; Glimmering Staff (249392, -1.92 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.7 spell_power points (2.08 DPS) | yes | Dwarven Tome (279898, -1.02 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.13 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -1.13 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 142.3 spell_power points (33.93 DPS) | yes | Starfaller (13063, -1.03 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.73 DPS) [crafted]; Gravestone Scepter (7001, -4.93 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053225103101321)

Set DPS (verified): 134.0. Weights run: 2.0s. Verify run: 1.2s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.656 ± 0.020, crit=0.115 ± 0.004 per rating point (14 rating = 1%, 1.607 per %), hit=0.289 ± 0.020 per rating point (10 rating = 1%, 2.885 per %), spell_haste=not significant (0.404 ± 0.234), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.293 ± 0.000, fire_power=0.707 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.19 DPS) | yes | Augural Shroud (2620, -1.78 DPS, sim-verified) [world]; Living Cowl (5608, -1.98 DPS) [world]; Enchanter's Cowl (4322, -2.09 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.9 spell_power points (2.71 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.90 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.57 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.57 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.5 spell_power points (3.84 DPS) | yes | Bloodmage Mantle (7684, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest]; Green Silken Shoulders (7057, -0.72 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.9 spell_power points (3.69 DPS) | yes | Guardian Cloak (5965, -1.39 DPS) [crafted]; Long Silken Cloak (4326, -1.81 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -1.90 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.9 spell_power points (6.42 DPS) | yes | Dreamweave Vest (10021, -0.50 DPS) [crafted]; Robe of Power (7054, -1.00 DPS) [crafted]; Elemental Raiment (9434, -1.22 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.23 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.49 DPS) [quest]; Windchaser Cuffs (14429, -0.77 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (5.10 DPS) | yes | Black Mageweave Gloves (10003, -1.39 DPS) [crafted]; Red Mageweave Gloves (10018, -1.68 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.99 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.8 spell_power points (4.17 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.89 DPS) [crafted]; Star Belt (4329, -0.95 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.9 spell_power points (5.41 DPS) | yes | Crimson Silk Pantaloons (7062, -1.52 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.89 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.94 DPS) | yes | Gilded Slippers (254001, -2.66 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.40 DPS) [dungeon]; Spidersilk Boots (4320, -3.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (3.45 DPS) | yes | Ring of Forlorn Spirits (2043, -1.47 DPS) [quest]; Reedknot Ring (9622, -1.72 DPS) [quest]; Minor Channeling Ring (1449, -1.89 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.23 DPS) | yes | Reedknot Ring (9622, -0.49 DPS) [quest]; Minor Channeling Ring (1449, -0.66 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.90 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Dar'Orahil (15106, -2.45 DPS) [quest]; Windweaver Staff (7757, -2.51 DPS) [dungeon]; Gut Ripper (2164, -6.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (134.0 DPS) | yes | Earthen Rod (9381, -0.08 DPS) [dungeon]; Twisted Nether Wand (249144, -0.19 DPS) [crafted]; Jaina's Firestarter (13064, -1.95 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Umbral Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 212.1. Weights run: 1.9s. Verify run: 1.5s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.360 ± 0.023, crit=0.149 ± 0.005 per rating point (14 rating = 1%, 2.087 per %), hit=0.407 ± 0.024 per rating point (10 rating = 1%, 4.070 per %), spell_haste=not significant (0.443 ± 0.310), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.293 ± 0.000, fire_power=0.707 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.33 DPS) | yes | Dreamweave Circlet (10041, -0.65 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.63 DPS) [crafted]; Red Mageweave Headband (10033, -2.27 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.2 spell_power points (2.49 DPS) | yes | Mindburst Medallion (11196, -0.27 DPS) [quest]; Horizon Choker (13085, -1.12 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.51 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 19.5 spell_power points (5.29 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.67 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.69 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.2 spell_power points (4.39 DPS) | yes | Runecloth Cloak (13860, -1.16 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.18 DPS, sim-verified) [dungeon]; Nightfall Drape (12465, -1.94 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 26.2 spell_power points (7.12 DPS) | yes | Robe of the Magi (1716, -0.55 DPS) [world_drop]; Dreamweave Vest (10021, -1.35 DPS) [crafted]; Elemental Raiment (9434, -1.41 DPS) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 9.5 spell_power points (2.59 DPS) | yes | Spidertank Oilrag (9448, -0.14 DPS) [dungeon]; Bloodband Bracers (11469, -0.35 DPS) [quest]; Arcane Runed Bracers (4744, -3.53 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.4 spell_power points (5.28 DPS) | yes | Runecloth Gloves (13863, -1.14 DPS) [crafted]; Black Mageweave Gloves (10003, -1.21 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.07 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 20.0 spell_power points (5.44 DPS) | yes | Highlander's Cloth Girdle (20098, -1.25 DPS) [rep]; Ghostweave Cord (254073, -1.64 DPS) [crafted]; Satyrmane Sash (17755, -3.04 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.6 spell_power points (7.22 DPS) | yes | Knight's Dreadweave Leggings (220888, -2.22 DPS) [vendor]; Red Mageweave Pants (10009, -2.25 DPS) [crafted]; Wizardweave Leggings (14132, -4.75 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.52 DPS) | yes | Gilded Sandals (254107, -2.65 DPS) [crafted]; Black Mageweave Boots (10026, -2.85 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.79 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.53 DPS) | yes | Lorekeeper's Ring (19523, -0.27 DPS) [rep]; Cyclopean Band (11824, -0.40 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.2 spell_power points (3.30 DPS) | yes | Lorekeeper's Ring (19523, -0.04 DPS) [rep]; Cyclopean Band (11824, -0.17 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.88 DPS, sim-verified) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.60 DPS) [world_drop]; Soul Harvester (20536, -2.76 DPS) [quest]; Blade of Eternal Darkness (17780, -5.02 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (212.1 DPS) | yes | Wand of Allistarj (13065, -2.94 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.44 DPS) [crafted]; Pyric Caduceus (11748, -5.03 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 422.8. Weights run: 2.0s. Verify run: 3.3s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.493 ± 0.033, crit=0.219 ± 0.006 per rating point (14 rating = 1%, 3.068 per %), hit=0.627 ± 0.034 per rating point (10 rating = 1%, 6.265 per %), spell_haste=not significant (-1.774 ± 0.541), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.386 ± 0.001, fire_power=0.614 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.37 DPS) [pvp]; Deathmist Mask (226909, -4.57 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.12 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.22 DPS) [quest]; Diana's Pearl Necklace (22403, -1.03 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.5 spell_power points (13.83 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.51 DPS) [pvp]; Burial Shawl (18681, -1.85 DPS, sim-verified) [dungeon]; Mantle of the Timbermaw (19050, -4.06 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.2 spell_power points (9.67 DPS) | yes | Crystalline Threaded Cape (20697, -1.56 DPS) [world]; Hide of the Wild (18510, -2.69 DPS) [crafted]; Amplifying Cloak (18350, -3.03 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.4 spell_power points (18.62 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -2.44 DPS) [pvp]; Robe of Everlasting Night (18385, -2.92 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -5.75 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.9 spell_power points (9.58 DPS) | yes | Sublime Wristguards (18497, -3.33 DPS) [dungeon]; Runecloth Cuffs (254123, -3.70 DPS) [crafted]; Deathmist Bracers (226907, -4.81 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.5 spell_power points (10.88 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -1.40 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 42.1 spell_power points (15.55 DPS) | yes | Belt of the Archmage (18405, -4.11 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.80 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -7.08 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 43.1 spell_power points (15.90 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -3.19 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.9 spell_power points (9.57 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Dragonrider Boots (18102, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.91 DPS) [dungeon]; Maiden's Circle (13001, -2.20 DPS) [world_drop]; Naglering (11669, -8.65 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.81 DPS) [dungeon]; Maiden's Circle (13001, -1.11 DPS) [world_drop]; Naglering (11669, -7.66 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+15.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.21 DPS) [quest]; Weakness Analyzer (272438, -2.58 DPS) [vendor]; Serenity Field (272439, -5.54 DPS) [vendor] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -1.67 DPS, sim-verified) [quest] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.02 DPS) [world]; Teebu's Blazing Longsword (1728, -18.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+23.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.49 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.10 DPS) [world]; Torch of Light (279246, -23.01 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 836.6. Weights run: 6.7s. Verify run: 3.6s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.076 ± 0.025, crit=0.284 ± 0.008 per rating point (14 rating = 1%, 3.978 per %), hit=0.842 ± 0.053 per rating point (10 rating = 1%, 8.424 per %), spell_haste=not significant (-0.602 ± 0.664), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.301 ± 0.001, fire_power=0.699 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.6 spell_power points (16.28 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -1.92 DPS) [crafted]; Deathmist Mask (226909, -3.81 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (11.70 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.11 DPS) [dungeon]; Amulet of the Dawn (22657, -3.20 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 32.1 spell_power points (17.09 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -3.10 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.40 DPS) [crafted]; Argent Shoulders (19059, -8.25 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.0 spell_power points (13.32 DPS) | yes | Crystalline Threaded Cape (20697, -3.23 DPS, sim-verified) [world]; Amplifying Cloak (18350, -3.74 DPS) [dungeon]; Hide of the Wild (18510, -5.46 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (24.83 DPS) | yes | Robe of Everlasting Night (18385, -5.77 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -6.84 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.72 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.6 spell_power points (12.03 DPS) | yes | Sublime Wristguards (18497, -5.24 DPS) [dungeon]; Runecloth Cuffs (254123, -5.77 DPS) [crafted]; Arcane Runed Bracers (4744, -7.24 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (14.56 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -2.64 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 35.1 spell_power points (18.67 DPS) | yes | Ban'thok Sash (11662, -7.36 DPS) [dungeon]; Belt of the Archmage (18405, -8.41 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -8.69 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.7 spell_power points (21.66 DPS) | yes | Marshal's Dreadweave Leggings (231587, -1.21 DPS) [pvp]; Skyshroud Leggings (13170, -3.25 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -6.24 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.77 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.06 DPS) [crafted]; Omnicast Boots (11822, -1.64 DPS) [dungeon] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.54 DPS) [quest]; Maiden's Circle (13001, -3.13 DPS) [world_drop]; Naglering (11669, -20.74 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.69 DPS) [quest]; Maiden's Circle (13001, -2.29 DPS) [world_drop]; Naglering (11669, -13.11 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+21.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -3.19 DPS) [quest]; Weakness Analyzer (272438, -3.72 DPS) [vendor]; Serenity Field (272439, -7.98 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -2.39 DPS) [world]; Teebu's Blazing Longsword (1728, -29.54 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (836.6 DPS) | yes | Bonecreeper Stylus (13938, -1.06 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.35 DPS) [world]; Torch of Light (279246, -33.95 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 38.3. Weights run: 2.3s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.045 ± 0.003, crit=0.025 ± 0.001 per rating point (14 rating = 1%, 0.356 per %), hit=0.106 ± 0.003 per rating point (10 rating = 1%, 1.063 per %), spell_haste=0.510 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.761 ± 0.002, fire_power=0.239 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.25 DPS) | yes | Shadow Goggles (4373, -2.64 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.4 spell_power points (1.12 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.05 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.83 DPS) | yes | Feyscale Cloak (6632, -0.21 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.21 DPS) [rep]; Black Whelp Cloak (7283, -0.45 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.2 spell_power points (1.08 DPS) | yes | Green Woolen Vest (2582, -0.25 DPS) [crafted]; Bloody Apron (6226, -0.25 DPS) [dungeon]; Gray Woolen Robe (2585, -1.42 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.1 spell_power points (0.23 DPS) | yes | Windsong Bangles (263336, -0.02 DPS) [quest]; Tabitha's Cuffs (251486, -0.17 DPS) [quest]; Featherbead Bracers (15452, -0.18 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.45 DPS) | yes | Gnoll Casting Gloves (892, -0.28 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.60 DPS) [crafted]; Apothecary Gloves (10919, -0.62 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.2 spell_power points (0.87 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Keller's Girdle (2911, -0.79 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.88 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.4 spell_power points (1.94 DPS) | yes | Silk-threaded Trousers (1929, -0.55 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.64 DPS) [crafted]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.2 spell_power points (1.49 DPS) | yes | Feather Padded Treads (285345, -0.53 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.66 DPS) [crafted]; Pristine Boots (253889, -0.84 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.04 DPS) | yes | Lavishly Jeweled Ring (1156, -0.98 DPS) [dungeon]; Loop of Sacrifice (281673, -0.99 DPS) [quest]; Volcanic Rock Ring (12053, -1.01 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.62 DPS) | yes | Loop of Sacrifice (281673, -0.58 DPS) [quest]; Volcanic Rock Ring (12053, -0.60 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.85 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.4 spell_power points (0.09 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.02 DPS) [world]; Lesser Staff of the Spire (1300, -0.04 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 110.4 spell_power points (22.93 DPS) | yes | Skycaller (12984, -1.46 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.64 DPS) [dungeon]; Sizzle Stick (8071, -4.24 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-0000000000000000000-2053225101000000)

Set DPS (verified): 70.3. Weights run: 2.4s. Verify run: 1.3s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.290 ± 0.008, crit=0.065 ± 0.002 per rating point (14 rating = 1%, 0.911 per %), hit=0.156 ± 0.007 per rating point (10 rating = 1%, 1.564 per %), spell_haste=0.651 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.771 ± 0.002, fire_power=0.229 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.62 DPS) | yes | Enchanter's Cowl (4322, -0.50 DPS) [crafted]; Silk Headband (7050, -0.51 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.72 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (2.08 DPS) | yes | Crystal Starfire Medallion (5003, -1.81 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.81 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.95 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.6 spell_power points (2.77 DPS) | yes | Fairywing Mantle (9536, -0.72 DPS) [quest]; Invoker's Mantle (215365, -0.75 DPS) [crafted]; Death Speaker Mantle (6685, -0.91 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.19 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.24 DPS) [crafted]; Battle Healer's Cloak (19529, -0.24 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.10 DPS) | yes | Death Speaker Robes (6682, -0.67 DPS) [dungeon]; Green Silk Armor (7065, -0.70 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.95 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.15 DPS) | yes | Nightsky Wristbands (6407, -1.73 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.73 DPS) [quest]; Glowing Magical Bracelets (13106, -2.51 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.5 spell_power points (1.78 DPS) | yes | Serpent Gloves (5970, -0.11 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.35 DPS) [world]; Truefaith Gloves (7049, -0.38 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 spell_power points (2.83 DPS) | yes | Warsong Sash (16975, -0.21 DPS) [quest]; Belt of Arugal (6392, -0.48 DPS) [dungeon]; Invoker's Cord (215366, -0.82 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.86 DPS) | yes | Pristine Leggings (253987, -0.71 DPS) [crafted]; Abomination Skin Leggings (23173, -0.96 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -1.02 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.0 spell_power points (2.15 DPS) | yes | Acidic Walkers (9454, -0.41 DPS) [dungeon]; Boots of the Enchanter (4325, -0.96 DPS) [crafted]; Spidersilk Boots (4320, -2.15 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.67 DPS) | yes | Electrocutioner Lagnut (9447, -0.95 DPS) [dungeon]; Sludge-Stained Band (286535, -0.95 DPS) [world]; Black Widow Band (6199, -1.18 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.43 DPS) | yes | Electrocutioner Lagnut (9447, -0.72 DPS) [dungeon]; Black Widow Band (6199, -0.95 DPS) [world]; Sludge-Stained Band (286535, -2.19 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.15 DPS) | yes | Twisted Chanter's Staff (890, -1.45 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.45 DPS) [quest]; Glimmering Staff (249392, -2.03 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.7 spell_power points (2.08 DPS) | yes | Orb of Souls (249395, -1.13 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.33 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.37 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 142.3 spell_power points (33.93 DPS) | yes | Starfaller (13063, -1.07 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.73 DPS) [crafted]; Gravestone Scepter (7001, -4.93 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2053225103101321)

Set DPS (verified): 130.2. Weights run: 2.0s. Verify run: 1.2s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.656 ± 0.020, crit=0.115 ± 0.004 per rating point (14 rating = 1%, 1.607 per %), hit=0.289 ± 0.020 per rating point (10 rating = 1%, 2.885 per %), spell_haste=not significant (0.404 ± 0.234), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.293 ± 0.000, fire_power=0.707 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.19 DPS) | yes | Living Cowl (5608, -1.98 DPS) [world]; Enchanter's Cowl (4322, -2.09 DPS) [crafted]; Augural Shroud (2620, -2.31 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.9 spell_power points (2.71 DPS) | yes | Triune Amulet (7722, -1.57 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.57 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.58 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.5 spell_power points (3.84 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.9 spell_power points (3.69 DPS) | yes | Guardian Cloak (5965, -1.39 DPS) [crafted]; Long Silken Cloak (4326, -1.78 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -1.90 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.9 spell_power points (6.42 DPS) | yes | Dreamweave Vest (10021, -0.65 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.00 DPS) [crafted]; Elemental Raiment (9434, -1.22 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.3 spell_power points (2.29 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.06 DPS) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (5.10 DPS) | yes | Black Mageweave Gloves (10003, -1.39 DPS) [crafted]; Gilded Handwraps (254021, -1.99 DPS) [crafted]; Red Mageweave Gloves (10018, -2.23 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.8 spell_power points (4.17 DPS) | yes | Defiler's Cloth Girdle (20166, -0.05 DPS) [rep]; Gilded Cord (254037, -0.89 DPS) [crafted]; Star Belt (4329, -0.95 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.9 spell_power points (5.41 DPS) | yes | Crimson Silk Pantaloons (7062, -1.21 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.89 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.94 DPS) | yes | Gilded Slippers (254001, -2.61 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.40 DPS) [dungeon]; Spidersilk Boots (4320, -3.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (3.45 DPS) | yes | Reedknot Ring (9622, -1.72 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.96 DPS) [vendor]; Black Widow Band (6199, -2.31 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.23 DPS) | yes | Sea Giant's Toe Ring (274746, -0.74 DPS) [vendor]; Reedknot Ring (9622, -0.87 DPS, sim-verified) [quest]; Black Widow Band (6199, -1.09 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (130.2 DPS) | yes | Staff of Dar'Orahil (15106, -2.45 DPS) [quest]; Windweaver Staff (7757, -2.51 DPS) [dungeon]; Gut Ripper (2164, -6.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 163.1 spell_power points (40.35 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.76 DPS) [dungeon]; Twisted Nether Wand (249144, -4.87 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 210.4. Weights run: 1.9s. Verify run: 1.4s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.360 ± 0.023, crit=0.149 ± 0.005 per rating point (14 rating = 1%, 2.087 per %), hit=0.407 ± 0.024 per rating point (10 rating = 1%, 4.070 per %), spell_haste=not significant (0.443 ± 0.310), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.293 ± 0.000, fire_power=0.707 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.33 DPS) | yes | Red Mageweave Headband (10033, -0.22 DPS) [crafted]; Dreamweave Circlet (10041, -0.65 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.63 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.2 spell_power points (2.49 DPS) | yes | Mindburst Medallion (11196, -0.27 DPS) [quest]; Horizon Choker (13085, -1.12 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.51 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 19.5 spell_power points (5.29 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.67 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.69 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.2 spell_power points (4.39 DPS) | yes | Deep Woodlands Cloak (19121, -0.25 DPS) [quest]; Mantle of Lady Falther'ess (23178, -1.06 DPS) [dungeon]; Runecloth Cloak (13860, -1.16 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 26.2 spell_power points (7.12 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.35 DPS) [crafted]; Elemental Raiment (9434, -1.41 DPS) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 9.5 spell_power points (2.59 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -1.41 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.4 spell_power points (5.28 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -0.87 DPS) [vendor]; Runecloth Gloves (13863, -1.14 DPS) [crafted]; Black Mageweave Gloves (10003, -1.21 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 20.0 spell_power points (5.44 DPS) | yes | Defiler's Cloth Girdle (20166, -1.25 DPS) [rep]; Ghostweave Cord (254073, -1.64 DPS) [crafted]; Satyrmane Sash (17755, -2.93 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.6 spell_power points (7.22 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -2.22 DPS) [vendor]; Red Mageweave Pants (10009, -2.25 DPS) [crafted]; Wizardweave Leggings (14132, -3.59 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.52 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.19 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -2.65 DPS) [crafted]; Black Mageweave Boots (10026, -2.85 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.53 DPS) | yes | Advisor's Ring (19519, -0.27 DPS) [rep]; Cyclopean Band (11824, -0.40 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.2 spell_power points (3.30 DPS) | yes | Advisor's Ring (19519, -0.04 DPS) [rep]; Cyclopean Band (11824, -0.17 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.22 DPS) [quest] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.84 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -2.49 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.60 DPS) [world_drop]; Soul Harvester (20536, -2.76 DPS) [quest]; Blade of Eternal Darkness (17780, -3.19 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (210.4 DPS) | yes | Wand of Allistarj (13065, -2.94 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.44 DPS) [crafted]; Pyric Caduceus (11748, -7.74 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 419.7. Weights run: 2.0s. Verify run: 3.3s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.493 ± 0.033, crit=0.219 ± 0.006 per rating point (14 rating = 1%, 3.068 per %), hit=0.627 ± 0.034 per rating point (10 rating = 1%, 6.265 per %), spell_haste=not significant (-1.774 ± 0.541), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.386 ± 0.001, fire_power=0.614 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.37 DPS) [pvp]; Deathmist Mask (226909, -5.28 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.12 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.22 DPS) [quest]; Diana's Pearl Necklace (22403, -1.03 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.5 spell_power points (13.83 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.51 DPS) [pvp]; Burial Shawl (18681, -2.44 DPS, sim-verified) [dungeon]; Mantle of the Timbermaw (19050, -4.06 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.2 spell_power points (9.67 DPS) | yes | Crystalline Threaded Cape (20697, -1.56 DPS) [world]; Hide of the Wild (18510, -2.69 DPS) [crafted]; Amplifying Cloak (18350, -3.03 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.4 spell_power points (18.62 DPS) | yes | Warlord's Dreadweave Robe (231591, -2.44 DPS) [pvp]; Robe of Everlasting Night (18385, -3.86 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -5.75 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.9 spell_power points (9.58 DPS) | yes | Sublime Wristguards (18497, -3.33 DPS) [dungeon]; Runecloth Cuffs (254123, -3.70 DPS) [crafted]; Deathmist Bracers (226907, -4.81 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.5 spell_power points (10.88 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -1.40 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 42.1 spell_power points (15.55 DPS) | yes | Belt of the Archmage (18405, -2.36 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.80 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -7.08 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 43.1 spell_power points (15.90 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.10 DPS) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.9 spell_power points (9.57 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.01 DPS) [dungeon]; Blood Guard's Dreadweave Walkers (227098, -0.56 DPS) [pvp] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.91 DPS) [dungeon]; Maiden's Circle (13001, -2.20 DPS) [world_drop]; Naglering (11669, -5.34 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.81 DPS) [dungeon]; Maiden's Circle (13001, -1.11 DPS) [world_drop]; Naglering (11669, -5.36 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+14.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.21 DPS) [quest]; Weakness Analyzer (272438, -2.58 DPS) [vendor]; Serenity Field (272439, -5.54 DPS) [vendor] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.02 DPS) [world]; Teebu's Blazing Longsword (1728, -17.24 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+25.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.49 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.10 DPS) [world]; Torch of Light (279246, -25.31 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 833.4. Weights run: 6.7s. Verify run: 3.6s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.076 ± 0.025, crit=0.284 ± 0.008 per rating point (14 rating = 1%, 3.978 per %), hit=0.842 ± 0.053 per rating point (10 rating = 1%, 8.424 per %), spell_haste=not significant (-0.602 ± 0.664), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.301 ± 0.001, fire_power=0.699 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.6 spell_power points (16.28 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -1.92 DPS) [crafted]; Deathmist Mask (226909, -3.53 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (11.70 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.11 DPS) [dungeon]; Amulet of the Dawn (22657, -3.20 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 32.1 spell_power points (17.09 DPS) | yes | Warlord's Dreadweave Mantle (231592, -3.10 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.40 DPS) [crafted]; Argent Shoulders (19059, -6.40 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.0 spell_power points (13.32 DPS) | yes | Amplifying Cloak (18350, -3.74 DPS) [dungeon]; Crystalline Threaded Cape (20697, -5.30 DPS, sim-verified) [world]; Hide of the Wild (18510, -5.46 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (24.83 DPS) | yes | Robe of Everlasting Night (18385, -4.57 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -6.84 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.72 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.6 spell_power points (12.03 DPS) | yes | Sublime Wristguards (18497, -5.24 DPS) [dungeon]; Runecloth Cuffs (254123, -5.77 DPS) [crafted]; Spidertank Oilrag (9448, -7.24 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (14.56 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -2.64 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 35.1 spell_power points (18.67 DPS) | yes | Ban'thok Sash (11662, -7.36 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -8.69 DPS) [rep]; Belt of the Archmage (18405, -9.68 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.7 spell_power points (21.66 DPS) | yes | General's Dreadweave Pants (231588, -1.21 DPS) [pvp]; Skyshroud Leggings (13170, -3.25 DPS) [dungeon]; Outrider's Silk Leggings (22747, -5.99 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.77 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.06 DPS) [crafted]; Omnicast Boots (11822, -1.64 DPS) [dungeon] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.54 DPS) [quest]; Maiden's Circle (13001, -3.13 DPS) [world_drop]; Naglering (11669, -18.64 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.69 DPS) [quest]; Maiden's Circle (13001, -2.29 DPS) [world_drop]; Naglering (11669, -12.56 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+21.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -3.19 DPS) [quest]; Weakness Analyzer (272438, -3.72 DPS) [vendor]; Serenity Field (272439, -7.98 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -2.39 DPS) [world]; Teebu's Blazing Longsword (1728, -29.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (833.4 DPS) | yes | Bonecreeper Stylus (13938, -1.06 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.35 DPS) [world]; Torch of Light (279246, -39.11 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

