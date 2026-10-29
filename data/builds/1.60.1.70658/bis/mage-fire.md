# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 0000000000000000000-23510000000000000-000000000000000000)

Set DPS (verified): 28.2. Weights run: 1.1s. Verify run: 0.8s. 145 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.413 ± 0.011, crit=0.127 ± 0.004 per rating point (14 rating = 1%, 1.778 per %), hit=0.389 ± 0.013 per rating point (10 rating = 1%, 3.895 per %), spell_haste=not significant (0.525 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.71 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Shadow Goggles (4373, -0.58 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.7 spell_power points (0.56 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.13 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.31 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Black Whelp Cloak (7283, -0.06 DPS) [crafted]; Caretaker's Cape (20428, -0.06 DPS) [rep]; Pearl-clasped Cloak (5542, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 11.1 spell_power points (0.72 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.32 DPS) [crafted]; Green Woolen Robe (6243, -0.44 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.1 spell_power points (0.13 DPS) | yes | Windsong Bangles (263336, -0.07 DPS) [quest]; Repurposed Hair Band (281256, -0.08 DPS) [quest]; Bright Bracers (3647, -0.56 DPS, sim-verified) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.58 DPS) | yes | Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world]; Flame Gloves (253917, -0.22 DPS, sim-verified) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.7 spell_power points (0.56 DPS) | yes | Pristine Sash (253925, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.35 DPS) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.5 spell_power points (0.87 DPS) | yes | Phoenix Pants (4317, -0.12 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted]; Abomination Skin Leggings (23173, -0.37 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 spell_power points (0.56 DPS) | yes | Pristine Boots (253889, -0.29 DPS) [crafted]; Red Woolen Boots (4313, -0.30 DPS) [crafted]; Flame Boots (253893, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.38 DPS) | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.30 DPS) [world_drop]; The 1 Ring (8350, -0.35 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.32 DPS) | yes | Volcanic Rock Ring (12053, -0.24 DPS) [world_drop]; The 1 Ring (8350, -0.30 DPS) [world]; Lavishly Jeweled Ring (1156, -0.55 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 4.1 spell_power points (0.27 DPS) | yes | Witching Stave (1484, -0.13 DPS) [dungeon]; Staff of Westfall (2042, -0.13 DPS) [quest]; Lesser Staff of the Spire (1300, -0.21 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 347.7 spell_power points (22.50 DPS) | yes | Skycaller (12984, -0.95 DPS, sim-verified) [world_drop]; Deepblaze (279896, -3.97 DPS) [quest]; Sizzle Stick (8071, -4.53 DPS) [quest] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 145, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 0000000000000000000-23552100120000000-000000000000000000)

Set DPS (verified): 50.1. Weights run: 1.3s. Verify run: 0.8s. 264 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.700 ± 0.013, crit=0.172 ± 0.007 per rating point (14 rating = 1%, 2.406 per %), hit=0.360 ± 0.017 per rating point (10 rating = 1%, 3.602 per %), spell_haste=0.949 ± 0.187, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.32 DPS) | yes | Enchanter's Cowl (4322, -0.09 DPS) [crafted]; Holy Shroud (2721, -0.28 DPS) [world_drop]; Flame Circlet (253953, -0.28 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 spell_power points (1.06 DPS) | yes | Darkspear Warding Pendant (272075, -0.73 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.85 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (1.44 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 5.6 spell_power points (0.53 DPS) | yes | Hillman's Cloak (3719, -0.06 DPS) [crafted]; Resilient Cape (14400, -0.13 DPS) [world_drop]; Vine Pruner's Cloak (279835, -0.61 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.1 spell_power points (1.71 DPS) | yes | Flame Gown (253965, -0.02 DPS) [crafted]; Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.23 DPS) | yes | Spidertank Oilrag (9448, -0.38 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.70 DPS) [world_drop]; Nightsky Wristbands (6407, -0.83 DPS) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.1 spell_power points (0.86 DPS) | yes | Town Clerk's Mittens (270029, -0.11 DPS) [quest]; Truefaith Gloves (7049, -0.19 DPS) [crafted]; Phoenix Gloves (4331, -0.50 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.1 spell_power points (1.24 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Crimson Silk Belt (7055, -0.21 DPS) [crafted]; Invoker's Cord (215366, -0.25 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.9 spell_power points (1.78 DPS) | yes | Abomination Skin Leggings (23173, -0.41 DPS) [dungeon]; Smoldering Pants (3073, -0.56 DPS) [world]; Filigreed Flame Leggings (253941, -0.68 DPS, sim-verified) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.9 spell_power points (1.69 DPS) | yes | Gilded Slippers (254001, -0.57 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.90 DPS) | yes | Minor Channeling Ring (1449, -0.29 DPS) [quest]; Ring of Power Regulation (273030, -0.33 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.66 DPS) | yes | Minor Channeling Ring (1449, -0.06 DPS) [quest]; Ring of Power Regulation (273030, -0.09 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.39 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.51 DPS) [quest]; Glimmering Staff (249392, -2.67 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 351.3 spell_power points (33.12 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.65 DPS) [crafted]; Scorching Wand (5213, -3.71 DPS) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Flame Gloves; waist: Highlander's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 0000000000000000000-23552100130125010-000000000000000000)

Set DPS (verified): 82.2. Weights run: 1.3s. Verify run: 1.0s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.916 ± 0.031, crit=0.259 ± 0.020 per rating point (14 rating = 1%, 3.630 per %), hit=0.489 ± 0.040 per rating point (10 rating = 1%, 4.886 per %), spell_haste=2.140 ± 0.521, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.6 spell_power points (2.72 DPS) | yes | Augural Shroud (2620, -0.39 DPS) [world]; Whitemane's Chapeau (7720, -0.85 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.78 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.3 spell_power points (1.65 DPS) | yes | Triune Amulet (7722, -0.73 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.91 DPS) [vendor]; Scorn's Icy Choker (23169, -1.05 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.9 spell_power points (2.18 DPS) | yes | Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest]; Green Silken Shoulders (7057, -1.00 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.6 spell_power points (1.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.06 DPS) [vendor]; Blackforge Cape (6424, -0.37 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.5 spell_power points (3.17 DPS) | yes | Cindercloth Robe (10042, -0.06 DPS) [crafted]; Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.29 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.50 DPS) | yes | Arcane Runed Bracers (4744, -0.46 DPS) [quest]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Windchaser Cuffs (14429, -0.55 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.7 spell_power points (2.50 DPS) | yes | Fiery Handwraps (254025, -0.14 DPS) [crafted]; Red Mageweave Gloves (10018, -0.17 DPS) [crafted]; Crimson Silk Gloves (7064, -0.25 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 21.3 spell_power points (2.46 DPS) | yes | Deathmage Sash (10771, -0.07 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.42 DPS) [rep]; Gilded Cord (254037, -0.69 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.0 spell_power points (2.88 DPS) | yes | Flame Leggings (253991, -0.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.59 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.98 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Fiery Slippers (254005, -0.53 DPS) [crafted]; Gilded Slippers (254001, -1.22 DPS) [crafted]; Acidic Walkers (9454, -1.35 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.68 DPS) | yes | Lorekeeper's Ring (19524, -0.64 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.76 DPS) [quest]; Reedknot Ring (9622, -0.87 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.6 spell_power points (1.22 DPS) | yes | Lorekeeper's Ring (19524, -0.18 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.30 DPS) [quest]; Reedknot Ring (9622, -0.41 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.21 DPS) [dungeon]; Searing Golden Blade (12260, -0.68 DPS) [crafted] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (82.2 DPS) | yes | Orb of the Forgotten Seer (7685, -0.13 DPS) [dungeon]; Rod of Molten Fire (2565, -0.25 DPS) [world_drop]; Celestial Orb (7515, -1.11 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 347.1 spell_power points (40.01 DPS) | yes | Ragefire Wand (7513, -1.52 DPS) [quest]; Nether Force Wand (11263, -2.37 DPS) [quest]; Icefury Wand (7514, -2.51 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 2050000000000000000-23552100130125031-000000000000000000)

Set DPS (verified): 149.0. Weights run: 1.3s. Verify run: 1.3s. 436 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.886 ± 0.034, crit=0.303 ± 0.027 per rating point (14 rating = 1%, 4.237 per %), hit=0.558 ± 0.056 per rating point (10 rating = 1%, 5.585 per %), spell_haste=not significant (0.439 ± 0.642), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.7 spell_power points (4.90 DPS) | yes | Dreamweave Circlet (10041, -0.92 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.05 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.30 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.1 spell_power points (1.88 DPS) | yes | Horizon Choker (13085, -0.22 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.24 DPS) [dungeon]; Mindburst Medallion (11196, -0.37 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.0 spell_power points (3.86 DPS) | yes | Netherflame Shoulders (254053, -0.65 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.17 DPS) [vendor] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.1 spell_power points (2.68 DPS) | yes | Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted]; Cloak of Fire (14134, -0.81 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.7 spell_power points (4.90 DPS) | yes | Cindercloth Robe (10042, -1.30 DPS) [crafted]; Runecloth Tunic (13857, -1.33 DPS) [crafted]; Robe of the Magi (1716, -2.38 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.2 spell_power points (2.56 DPS) | yes | Nethergeld Cuffs (254061, -0.80 DPS) [crafted]; Phoenix Bindings (210781, -0.83 DPS) [crafted]; Bloodband Bracers (11469, -0.83 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.0 spell_power points (4.00 DPS) | yes | Fiery Gloves (254099, -0.28 DPS) [crafted]; Raider Handwraps (272098, -0.47 DPS) [vendor]; Dreamweave Gloves (10019, -1.13 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.0 spell_power points (3.60 DPS) | yes | Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Dawnspire Cord (12466, -0.55 DPS) [dungeon]; Fiery Cord (254041, -0.79 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 26.9 spell_power points (3.58 DPS) | yes | Red Mageweave Pants (10009, -0.30 DPS) [crafted]; Flame Leggings (253991, -0.89 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.98 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.0 spell_power points (3.73 DPS) | yes | Earthen Silk Slippers (254013, -0.53 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -0.86 DPS) [vendor]; Cindercloth Boots (10044, -0.93 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.92 DPS) | yes | Snake Hoop (6750, -0.53 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.86 DPS) [quest]; Band of the Unicorn (7553, -2.88 DPS, sim-verified) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-verified (149.0 DPS) | yes | Snake Hoop (6750, -0.21 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.53 DPS) [quest]; Band of the Unicorn (7553, -1.55 DPS, sim-verified) [world_drop] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.18 DPS) [quest]; Kindling Stave (11750, -0.68 DPS) [dungeon]; Hanzo Sword (8190, -24.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 393.6 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.83 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.68 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Sorcerer's Gauntlets; waist: Fiery Waistcord; legs: Knight's Dreadweave Leggings; feet: Fiery Sandals; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 2050015100000000000-23552100130125031-003000000000000000)

Set DPS (verified): 488.2. Weights run: 1.5s. Verify run: 2.6s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.845 ± 0.052, crit=0.465 ± 0.044 per rating point (14 rating = 1%, 6.503 per %), hit=1.218 ± 0.089 per rating point (10 rating = 1%, 12.181 per %), spell_haste=-11.057 ± 1.138, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 48.6 spell_power points (15.15 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.84 DPS) [pvp]; Magister's Crown (16686, -3.83 DPS) [dungeon] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (488.2 DPS) | yes | Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Amulet of the Dawn (22657, +0.00 DPS) [quest]; Chains of the Lich (23125, +0.00 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.2 spell_power points (14.39 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.65 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.64 DPS) [crafted]; Darkspear Shoulderpads (272103, -11.52 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 34.9 spell_power points (10.89 DPS) | yes | Hide of the Wild (18510, -3.89 DPS) [crafted]; Mageflame Cloak (13007, -4.34 DPS) [world_drop]; Cindercloth Cloak (14044, -4.73 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.6 spell_power points (17.65 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.86 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.60 DPS) [pvp]; Robes of Fiery Devastation (279266, -11.06 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.8 spell_power points (8.96 DPS) | yes | Runecloth Cuffs (254123, -2.90 DPS) [crafted]; Netherflame Cuffs (254065, -3.07 DPS) [crafted]; Marshal's Silk Bracers (16438, -3.96 DPS) [pvp] |
| hands | Inferno Gloves (18408) | Tailoring [crafted] | 40.6 spell_power points (12.65 DPS) | yes | Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gloves (16440, -1.08 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -1.08 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 55.8 spell_power points (17.38 DPS) | yes | Magician's Cord (272393, -5.14 DPS) [vendor]; Fiery Waistcord (254085, -9.09 DPS) [crafted]; Belt of the Archmage (18405, -11.49 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 53.5 spell_power points (16.65 DPS) | yes | Marshal's Silk Leggings (231605, -0.02 DPS) [pvp]; Sorcerer's Leggings (226933, -2.73 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -3.61 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.5 spell_power points (10.75 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.36 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 47.8 spell_power points (14.89 DPS) | yes | Channeler's Ring (272406, -7.41 DPS) [vendor]; Maiden's Circle (13001, -7.44 DPS) [world_drop]; Blessed Band of Light (272407, -7.98 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 31.3 spell_power points (9.75 DPS) | yes | Channeler's Ring (272406, -2.27 DPS) [vendor]; Maiden's Circle (13001, -2.30 DPS) [world_drop]; Blessed Band of Light (272407, -2.84 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (488.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (488.2 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -2.49 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -10.37 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (488.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.33 DPS) [quest]; Teebu's Blazing Longsword (1728, -67.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 252.4 spell_power points (78.63 DPS) | yes | Brilliant Wand (249385, -14.40 DPS) [crafted]; Stormrager (16997, -14.62 DPS) [quest]; Wand of Biting Cold (19108, -14.63 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Jewel of Kajaro; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Inferno Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 2050015100000000000-23452100130135031-003000000000000000)

Set DPS (verified): 872.3. Weights run: 1.7s. Verify run: 2.8s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.438 ± 0.029, crit=0.763 ± 0.040 per rating point (14 rating = 1%, 10.676 per %), hit=1.133 ± 0.072 per rating point (10 rating = 1%, 11.331 per %), spell_haste=not significant (-0.035 ± 0.561), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.6 spell_power points (24.49 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.76 DPS) [pvp]; Knight-Lieutenant's Dreadweave Hat (220889, -7.23 DPS, sim-verified) [vendor] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Jewel of Kajaro (19601, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.75 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.3 spell_power points (25.42 DPS) | yes | Mantle of the Timbermaw (19050, -7.18 DPS, sim-verified) [crafted]; Field Marshal's Silk Spaulders (231602, -7.28 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -7.90 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.8 spell_power points (17.71 DPS) | yes | Mageflame Cloak (13007, -5.84 DPS, sim-verified) [world_drop]; Hide of the Wild (18510, -7.15 DPS) [crafted]; Spritecaster Cape (11623, -8.16 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.9 spell_power points (32.13 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.76 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.66 DPS) [pvp]; Robes of Fiery Devastation (279266, -14.15 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.5 spell_power points (14.65 DPS) | yes | Netherflame Cuffs (254065, -5.42 DPS) [crafted]; Runecloth Cuffs (254123, -5.82 DPS) [crafted]; Phoenix Bindings (210781, -7.18 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+14.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Silk Gloves (16440, -1.42 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -1.42 DPS) [vendor]; Inferno Gloves (18408, -14.19 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 46.0 spell_power points (26.41 DPS) | yes | Belt of the Archmage (18405, -4.76 DPS) [crafted]; Magician's Cord (272393, -9.45 DPS) [vendor]; Highlander's Cloth Girdle (20047, -10.72 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 57.7 spell_power points (33.17 DPS) | yes | Marshal's Silk Leggings (231605, -4.76 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -10.69 DPS) [pvp]; Sorcerer's Leggings (226933, -11.00 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 28.0 spell_power points (16.09 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -1.51 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 43.3 spell_power points (24.86 DPS) | yes | Don Julio's Band (19325, -12.22 DPS) [rep]; Blessed Band of Light (272407, -12.61 DPS) [vendor]; Maiden's Circle (13001, -12.76 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 26.8 spell_power points (15.41 DPS) | yes | Don Julio's Band (19325, -2.77 DPS) [rep]; Blessed Band of Light (272407, -3.15 DPS) [vendor]; Maiden's Circle (13001, -3.30 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+10.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -4.60 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -10.38 DPS, sim-verified) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.68 DPS) [quest]; Teebu's Blazing Longsword (1728, -58.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 139.6 spell_power points (80.21 DPS) | yes | Brilliant Wand (249385, -13.69 DPS) [crafted]; Wand of Biting Cold (19108, -16.21 DPS) [quest]; Stormrager (16997, -16.26 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (orc, 0000000000000000000-23510000000000000-000000000000000000)

Set DPS (verified): 26.3. Weights run: 1.1s. Verify run: 0.8s. 134 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.413 ± 0.011, crit=0.127 ± 0.004 per rating point (14 rating = 1%, 1.778 per %), hit=0.389 ± 0.013 per rating point (10 rating = 1%, 3.895 per %), spell_haste=not significant (0.525 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.71 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Shadow Goggles (4373, -0.58 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.7 spell_power points (0.56 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.31 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.37 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Black Whelp Cloak (7283, -0.06 DPS) [crafted]; Battle Healer's Cloak (20427, -0.06 DPS) [rep]; Pearl-clasped Cloak (5542, -0.36 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 11.1 spell_power points (0.72 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.32 DPS) [crafted]; Green Woolen Robe (6243, -0.44 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.19 DPS) | yes | Mindthrust Bracers (1974, -0.06 DPS) [dungeon]; Featherbead Bracers (15452, -0.06 DPS) [quest]; Owlbeard Bracers (16981, -0.08 DPS) [quest] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.58 DPS) | yes | Flame Gloves (253917, -0.05 DPS) [crafted]; Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.7 spell_power points (0.56 DPS) | yes | Pristine Sash (253925, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.35 DPS) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.5 spell_power points (0.87 DPS) | yes | Phoenix Pants (4317, -0.12 DPS) [crafted]; Abomination Skin Leggings (23173, -0.23 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 spell_power points (0.56 DPS) | yes | Flame Boots (253893, -0.24 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.29 DPS) [crafted]; Red Woolen Boots (4313, -0.30 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.32 DPS) | yes | Loop of Sacrifice (281673, -0.19 DPS) [quest]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop]; The 1 Ring (8350, -0.30 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.5 spell_power points (0.16 DPS) | yes | Volcanic Rock Ring (12053, -0.08 DPS) [world_drop]; The 1 Ring (8350, -0.13 DPS) [world]; Loop of Sacrifice (281673, -0.16 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 4.1 spell_power points (0.27 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.11 DPS) [world]; Clear Crystal Rod (16894, -0.13 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 347.7 spell_power points (22.50 DPS) | yes | Skycaller (12984, -0.76 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.53 DPS) [quest]; Greater Magic Wand (11288, -4.87 DPS) [crafted] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Tabitha's Cuffs; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 134, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (orc, 0000000000000000000-23552100120000000-000000000000000000)

Set DPS (verified): 48.0. Weights run: 1.3s. Verify run: 0.8s. 245 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.700 ± 0.013, crit=0.172 ± 0.007 per rating point (14 rating = 1%, 2.406 per %), hit=0.360 ± 0.017 per rating point (10 rating = 1%, 3.602 per %), spell_haste=0.949 ± 0.187, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.32 DPS) | yes | Enchanter's Cowl (4322, -0.09 DPS) [crafted]; Holy Shroud (2721, -0.28 DPS) [world_drop]; Flame Circlet (253953, -0.28 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 spell_power points (1.06 DPS) | yes | Darkspear Warding Pendant (272075, -0.73 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.82 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (1.44 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 5.6 spell_power points (0.53 DPS) | yes | Windsong Drape (15468, -0.06 DPS) [quest]; Resilient Cape (14400, -0.13 DPS) [world_drop]; Hillman's Cloak (3719, -0.55 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.1 spell_power points (1.71 DPS) | yes | Flame Gown (253965, -0.02 DPS) [crafted]; Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.23 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -0.70 DPS) [world_drop]; Nightsky Wristbands (6407, -0.83 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.5 spell_power points (0.90 DPS) | yes | Flame Gloves (253917, -0.04 DPS) [crafted]; Phoenix Gloves (4331, -0.05 DPS) [crafted]; Truefaith Gloves (7049, -0.23 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.1 spell_power points (1.24 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Warsong Sash (16975, -0.20 DPS) [quest]; Crimson Silk Belt (7055, -0.21 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.9 spell_power points (1.78 DPS) | yes | Filigreed Flame Leggings (253941, -0.35 DPS) [crafted]; Abomination Skin Leggings (23173, -0.41 DPS) [dungeon]; Smoldering Pants (3073, -0.56 DPS) [world] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.9 spell_power points (1.69 DPS) | yes | Gilded Slippers (254001, -0.57 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.90 DPS) | yes | Ring of Power Regulation (273030, -0.33 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.66 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Ring of Power Regulation (273030, -0.09 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.39 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -2.67 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -2.73 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 351.3 spell_power points (33.12 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Unstable Power Core (279847, -3.41 DPS) [quest]; Greater Mystic Wand (217287, -3.65 DPS) [crafted] |

**New at 30:** head: Filigreed Flame Circlet; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 245, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (orc, 0000000000000000000-23552100130125010-000000000000000000)

Set DPS (verified): 77.5. Weights run: 1.3s. Verify run: 1.0s. 320 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.916 ± 0.031, crit=0.259 ± 0.020 per rating point (14 rating = 1%, 3.630 per %), hit=0.489 ± 0.040 per rating point (10 rating = 1%, 4.886 per %), spell_haste=2.140 ± 0.521, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.6 spell_power points (2.72 DPS) | yes | Augural Shroud (2620, -0.39 DPS) [world]; Whitemane's Chapeau (7720, -0.85 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.36 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.3 spell_power points (1.65 DPS) | yes | Triune Amulet (7722, -0.73 DPS) [dungeon]; Scorn's Icy Choker (23169, -0.76 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.91 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.9 spell_power points (2.18 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.6 spell_power points (1.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.06 DPS) [vendor]; Blackforge Cape (6424, -0.37 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.5 spell_power points (3.17 DPS) | yes | Cindercloth Robe (10042, -0.06 DPS) [crafted]; Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.29 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.50 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.19 DPS) [quest]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.7 spell_power points (2.50 DPS) | yes | Fiery Handwraps (254025, -0.14 DPS) [crafted]; Red Mageweave Gloves (10018, -0.17 DPS) [crafted]; Crimson Silk Gloves (7064, -0.25 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 21.3 spell_power points (2.46 DPS) | yes | Deathmage Sash (10771, -0.07 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.42 DPS) [rep]; Gilded Cord (254037, -0.69 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.0 spell_power points (2.88 DPS) | yes | Flame Leggings (253991, -0.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.59 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.98 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Fiery Slippers (254005, -0.86 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.22 DPS) [crafted]; Acidic Walkers (9454, -1.35 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.68 DPS) | yes | Advisor's Ring (19520, -0.64 DPS) [rep]; Reedknot Ring (9622, -0.87 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.6 spell_power points (1.22 DPS) | yes | Advisor's Ring (19520, -0.18 DPS) [rep]; Reedknot Ring (9622, -0.41 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (77.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.21 DPS) [dungeon]; Searing Golden Blade (12260, -0.68 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.7 spell_power points (1.82 DPS) | yes | Orb of the Forgotten Seer (7685, -0.20 DPS) [dungeon]; Rod of Molten Fire (2565, -0.32 DPS) [world_drop]; Orb of Mystic Insight (249394, -0.37 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 347.1 spell_power points (40.01 DPS) | yes | Ragefire Wand (7513, -1.52 DPS) [quest]; Nether Force Wand (11263, -2.37 DPS) [quest]; Icefury Wand (7514, -2.51 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 320, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 2050000000000000000-23552100130125031-000000000000000000)

Set DPS (verified): 142.3. Weights run: 1.3s. Verify run: 1.2s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.886 ± 0.034, crit=0.303 ± 0.027 per rating point (14 rating = 1%, 4.237 per %), hit=0.558 ± 0.056 per rating point (10 rating = 1%, 5.585 per %), spell_haste=not significant (0.439 ± 0.642), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.7 spell_power points (4.90 DPS) | yes | Dreamweave Circlet (10041, -0.92 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.05 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.30 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.1 spell_power points (1.88 DPS) | yes | Horizon Choker (13085, -0.22 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.24 DPS) [dungeon]; Mindburst Medallion (11196, -0.37 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.0 spell_power points (3.86 DPS) | yes | Netherflame Shoulders (254053, -0.65 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.17 DPS) [vendor] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.1 spell_power points (2.68 DPS) | yes | Deep Woodlands Cloak (19121, -0.02 DPS) [quest]; Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.7 spell_power points (4.90 DPS) | yes | Cindercloth Robe (10042, -1.30 DPS) [crafted]; Runecloth Tunic (13857, -1.33 DPS) [crafted]; Robe of the Magi (1716, -2.28 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.2 spell_power points (2.56 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.80 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.0 spell_power points (4.00 DPS) | yes | Raider Handwraps (272098, -0.47 DPS) [vendor]; Dreamweave Gloves (10019, -1.13 DPS) [crafted]; Fiery Gloves (254099, -1.54 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.0 spell_power points (3.60 DPS) | yes | Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Dawnspire Cord (12466, -0.55 DPS) [dungeon]; Fiery Cord (254041, -0.79 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 26.9 spell_power points (3.58 DPS) | yes | Red Mageweave Pants (10009, -0.30 DPS) [crafted]; Flame Leggings (253991, -0.89 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.98 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.0 spell_power points (3.73 DPS) | yes | Earthen Silk Slippers (254013, -0.53 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -0.86 DPS) [vendor]; Cindercloth Boots (10044, -0.93 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.92 DPS) | yes | Advisor's Ring (19519, -0.32 DPS) [rep]; Snake Hoop (6750, -0.53 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.73 DPS) | yes | Advisor's Ring (19519, -0.13 DPS) [rep]; Snake Hoop (6750, -0.34 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (142.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (142.3 DPS) | yes | Frozen Heart of the Mountain (249469, -0.13 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (142.3 DPS) | yes | Spellshifter Rod (9527, -0.18 DPS) [quest]; Kindling Stave (11750, -0.68 DPS) [dungeon]; Hanzo Sword (8190, -24.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 393.6 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.83 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.68 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Sorcerer's Gauntlets; waist: Fiery Waistcord; legs: Stone Guard's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Rune of the Guard Captain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 2050015100000000000-23552100130125031-003000000000000000)

Set DPS (verified): 485.8. Weights run: 1.5s. Verify run: 3.8s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.845 ± 0.052, crit=0.465 ± 0.044 per rating point (14 rating = 1%, 6.503 per %), hit=1.218 ± 0.089 per rating point (10 rating = 1%, 12.181 per %), spell_haste=-11.057 ± 1.138, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 48.6 spell_power points (15.15 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.84 DPS) [pvp]; Magister's Crown (16686, -19.99 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (485.8 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.89 DPS) [quest]; Chains of the Lich (23125, -1.24 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.2 spell_power points (14.39 DPS) | yes | Warlord's Silk Amice (231594, -2.65 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.64 DPS) [crafted]; Darkspear Shoulderpads (272103, -9.53 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 34.9 spell_power points (10.89 DPS) | yes | Hide of the Wild (18510, -3.89 DPS) [crafted]; Mageflame Cloak (13007, -4.34 DPS) [world_drop]; Cindercloth Cloak (14044, -4.73 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.6 spell_power points (17.65 DPS) | yes | Warlord's Silk Raiment (231596, -0.86 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.60 DPS) [pvp]; Robes of Fiery Devastation (279266, -9.06 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.8 spell_power points (8.96 DPS) | yes | Runecloth Cuffs (254123, -2.90 DPS) [crafted]; Netherflame Cuffs (254065, -3.07 DPS) [crafted]; General's Silk Cuffs (16538, -3.96 DPS) [pvp] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (485.8 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Inferno Gloves (18408, -17.16 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 55.8 spell_power points (17.38 DPS) | yes | Magician's Cord (272393, -5.14 DPS) [vendor]; Belt of the Archmage (18405, -6.93 DPS, sim-verified) [crafted]; Fiery Waistcord (254085, -9.09 DPS) [crafted] |
| legs | Sorcerer's Leggings (226933) | Anthion's Parting Words [quest] | sim-verified (485.8 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.20 DPS) [rep]; Sentinel's Silk Leggings (237815, -12.19 DPS, sim-verified) [vendor] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (485.8 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.36 DPS) [dungeon]; Sorcerer's Boots (22064, -7.58 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 47.8 spell_power points (14.89 DPS) | yes | Channeler's Ring (272406, -7.41 DPS) [vendor]; Maiden's Circle (13001, -7.44 DPS) [world_drop]; Blessed Band of Light (272407, -7.98 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 31.3 spell_power points (9.75 DPS) | yes | Channeler's Ring (272406, -2.27 DPS) [vendor]; Maiden's Circle (13001, -2.30 DPS) [world_drop]; Blessed Band of Light (272407, -2.84 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (485.8 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (485.8 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.95 DPS) [crafted]; Rune of the Guard Captain (19120, -1.71 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (485.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.33 DPS) [quest]; Teebu's Blazing Longsword (1728, -67.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 252.4 spell_power points (78.63 DPS) | yes | Brilliant Wand (249385, -14.40 DPS) [crafted]; Stormrager (16997, -14.62 DPS) [quest]; Wand of Biting Cold (19108, -14.63 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sorcerer's Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 2050015100000000000-23452100130135031-003000000000000000)

Set DPS (verified): 846.5. Weights run: 1.7s. Verify run: 2.6s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.438 ± 0.029, crit=0.763 ± 0.040 per rating point (14 rating = 1%, 10.676 per %), hit=1.133 ± 0.072 per rating point (10 rating = 1%, 11.331 per %), spell_haste=not significant (-0.035 ± 0.561), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.6 spell_power points (24.49 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.76 DPS) [pvp]; Red Mageweave Headband (10033, -16.70 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (846.5 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.75 DPS) [quest]; Jewel of Kajaro (19601, -8.19 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.3 spell_power points (25.42 DPS) | yes | Mantle of the Timbermaw (19050, -7.26 DPS, sim-verified) [crafted]; Warlord's Silk Amice (231594, -7.28 DPS) [pvp]; Champion's Silk Mantle (227104, -7.90 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.8 spell_power points (17.71 DPS) | yes | Hide of the Wild (18510, -7.15 DPS) [crafted]; Spritecaster Cape (11623, -8.16 DPS) [dungeon]; Mageflame Cloak (13007, -22.80 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.9 spell_power points (32.13 DPS) | yes | Warlord's Silk Raiment (231596, -2.76 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -9.66 DPS) [pvp]; Robes of Fiery Devastation (279266, -17.78 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.5 spell_power points (14.65 DPS) | yes | Netherflame Cuffs (254065, -5.42 DPS) [crafted]; Runecloth Cuffs (254123, -5.82 DPS) [crafted]; Phoenix Bindings (210781, -7.18 DPS) [crafted] |
| hands | Inferno Gloves (18408) | Tailoring [crafted] | 36.9 spell_power points (21.22 DPS) | yes | Gloves of Spell Mastery (14146, -1.27 DPS) [crafted]; General's Silk Handguards (16540, -2.69 DPS) [vendor]; General's Silk Gauntlets (231599, -2.69 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 46.0 spell_power points (26.41 DPS) | yes | Magician's Cord (272393, -9.45 DPS) [vendor]; Defiler's Cloth Girdle (20163, -10.72 DPS) [rep]; Belt of the Archmage (18405, -12.02 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 57.7 spell_power points (33.17 DPS) | yes | General's Silk Trousers (231595, -4.76 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -10.69 DPS) [pvp]; Sorcerer's Leggings (226933, -19.17 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 28.0 spell_power points (16.09 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.51 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 43.3 spell_power points (24.86 DPS) | yes | Don Julio's Band (19325, -12.22 DPS) [rep]; Blessed Band of Light (272407, -12.61 DPS) [vendor]; Maiden's Circle (13001, -12.76 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 26.8 spell_power points (15.41 DPS) | yes | Don Julio's Band (19325, -2.77 DPS) [rep]; Blessed Band of Light (272407, -3.15 DPS) [vendor]; Maiden's Circle (13001, -3.30 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (846.5 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (846.5 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Rune of the Guard Captain (19120, -1.30 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (846.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.68 DPS) [quest]; Teebu's Blazing Longsword (1728, -56.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 139.6 spell_power points (80.21 DPS) | yes | Brilliant Wand (249385, -13.69 DPS) [crafted]; Wand of Biting Cold (19108, -16.21 DPS) [quest]; Stormrager (16997, -16.26 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Inferno Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

