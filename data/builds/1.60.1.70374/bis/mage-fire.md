# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 28.2. Weights run: 0.8s. Verify run: 0.6s. 147 eligible items had no known source.

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

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 50.1. Weights run: 0.9s. Verify run: 0.6s. 261 eligible items had no known source.

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

No-known-source sample (15 of 261, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 000000000000000000-23552100130103041-0000000000000000000)

Set DPS (verified): 82.2. Weights run: 1.0s. Verify run: 0.7s. 339 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.774 ± 0.026, crit=0.219 ± 0.015 per rating point (14 rating = 1%, 3.059 per %), hit=0.376 ± 0.034 per rating point (10 rating = 1%, 3.765 per %), spell_haste=2.538 ± 0.457, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 21.2 spell_power points (2.61 DPS) | yes | Augural Shroud (2620, -0.30 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -0.86 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.33 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.2 spell_power points (1.63 DPS) | yes | Triune Amulet (7722, -0.64 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor]; Scorn's Icy Choker (23169, -0.99 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.1 spell_power points (2.11 DPS) | yes | Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -0.94 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.9 spell_power points (1.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.17 DPS) [vendor]; Icy Cloak (4327, -0.35 DPS) [crafted] |
| chest | Cindercloth Robe (10042) | Tailoring [crafted] | 27.0 spell_power points (3.33 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.25 DPS) [crafted]; Robe of Power (7054, -0.46 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.60 DPS) | yes | Arcane Runed Bracers (4744, -0.49 DPS) [quest]; Spidertank Oilrag (9448, -0.49 DPS) [dungeon]; Condor Bracers (15864, -0.74 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.60 DPS) | yes | Fiery Handwraps (254025, -0.21 DPS) [crafted]; Red Mageweave Gloves (10018, -0.29 DPS) [crafted]; Crimson Silk Gloves (7064, -0.30 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (2.49 DPS) | yes | Deathmage Sash (10771, -0.20 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.38 DPS) [rep]; Gilded Cord (254037, -0.74 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.3 spell_power points (2.87 DPS) | yes | Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Flame Leggings (253991, -0.83 DPS, sim-verified) [crafted]; Filigreed Flame Leggings (253941, -0.94 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.96 DPS) | yes | Fiery Slippers (254005, -0.69 DPS) [crafted]; Gilded Slippers (254001, -1.43 DPS) [crafted]; Acidic Walkers (9454, -1.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (1.71 DPS) | yes | Lorekeeper's Ring (19524, -0.60 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.72 DPS) [quest]; Reedknot Ring (9622, -0.85 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.9 spell_power points (1.22 DPS) | yes | Lorekeeper's Ring (19524, -0.11 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.23 DPS) [quest]; Reedknot Ring (9622, -0.35 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (82.2 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Searing Golden Blade (12260, -0.71 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.3 spell_power points (1.89 DPS) | yes | Orb of the Forgotten Seer (7685, -0.16 DPS) [dungeon]; Thrash's Trash (276204, -0.16 DPS) [vendor]; Orb of Lorica (11262, -0.20 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 323.6 spell_power points (39.95 DPS) | yes | Ragefire Wand (7513, -1.42 DPS) [quest]; Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Cindercloth Robe; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 339, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 151.1. Weights run: 0.9s. Verify run: 1.0s. 432 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.939 ± 0.038, crit=0.343 ± 0.024 per rating point (14 rating = 1%, 4.796 per %), hit=0.622 ± 0.048 per rating point (10 rating = 1%, 6.225 per %), spell_haste=-3.557 ± 0.738, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.8 spell_power points (5.51 DPS) | yes | Dreamweave Circlet (10041, -1.08 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.13 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.57 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.5 spell_power points (2.12 DPS) | yes | Horizon Choker (13085, -0.20 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.27 DPS) [dungeon]; Mindburst Medallion (11196, -0.42 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.9 spell_power points (4.36 DPS) | yes | Netherflame Shoulders (254053, -0.79 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.26 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.29 DPS) [crafted] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.5 spell_power points (2.99 DPS) | yes | Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Runecloth Cloak (13860, -0.58 DPS) [crafted]; Cloak of Fire (14134, -0.95 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.8 spell_power points (5.51 DPS) | yes | Robe of the Magi (1716, -1.48 DPS) [world_drop]; Runecloth Tunic (13857, -1.52 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.55 DPS) [vendor] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.6 spell_power points (2.85 DPS) | yes | Nethergeld Cuffs (254061, -0.87 DPS) [crafted]; Bloodband Bracers (11469, -0.89 DPS) [quest]; Phoenix Bindings (210781, -0.96 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.4 spell_power points (4.57 DPS) | yes | Fiery Gloves (254099, -0.43 DPS) [crafted]; Raider Handwraps (272098, -0.54 DPS) [vendor]; Dreamweave Gloves (10019, -1.40 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.5 spell_power points (4.00 DPS) | yes | Dawnspire Cord (12466, -0.53 DPS) [dungeon]; Satyrmane Sash (17755, -0.59 DPS) [dungeon]; Fiery Cord (254041, -0.87 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 28.1 spell_power points (4.09 DPS) | yes | Red Mageweave Pants (10009, -0.41 DPS) [crafted]; Flame Leggings (253991, -1.09 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.15 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.5 spell_power points (4.15 DPS) | yes | Earthen Silk Slippers (254013, -0.65 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -0.84 DPS) [vendor]; Cindercloth Boots (10044, -1.09 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (2.14 DPS) | yes | Lorekeeper's Ring (19523, -0.39 DPS) [rep]; Snake Hoop (6750, -0.58 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.90 DPS) | yes | Lorekeeper's Ring (19523, -0.15 DPS) [rep]; Snake Hoop (6750, -0.34 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (151.1 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (151.1 DPS) | yes | Mark of the Chosen (17774, -2.46 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (151.1 DPS) | yes | Kindling Stave (11750, -0.63 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.78 DPS) [dungeon]; Hanzo Sword (8190, -18.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 360.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.58 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Sorcerer's Gauntlets; waist: Fiery Waistcord; legs: Knight's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 432, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 433.6. Weights run: 1.0s. Verify run: 2.9s. 886 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.762 ± 0.050, crit=0.690 ± 0.040 per rating point (14 rating = 1%, 9.664 per %), hit=0.992 ± 0.076 per rating point (10 rating = 1%, 9.923 per %), spell_haste=-10.683 ± 1.080, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 49.7 spell_power points (13.98 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.50 DPS) [pvp]; Red Mageweave Headband (10033, -26.83 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (433.6 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.78 DPS) [quest]; Chains of the Lich (23125, -0.82 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.1 spell_power points (13.53 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.28 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -4.23 DPS) [pvp]; Mantle of the Timbermaw (19050, -7.23 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.0 spell_power points (9.01 DPS) | yes | Hide of the Wild (18510, -2.92 DPS) [crafted]; Mageflame Cloak (13007, -3.10 DPS) [world_drop]; Cindercloth Cloak (14044, -3.63 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.8 spell_power points (16.54 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.90 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.27 DPS) [vendor]; Robes of Fiery Devastation (279266, -7.37 DPS, sim-verified) [crafted] |
| wrist | Sorcerer's Bindings (226929) | An Earnest Proposition [quest] | sim-verified (433.6 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (433.6 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Inferno Gloves (18408, -12.97 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 51.7 spell_power points (14.54 DPS) | yes | Magician's Cord (272393, -4.05 DPS) [vendor]; Highlander's Cloth Girdle (20047, -6.60 DPS) [rep]; Belt of the Archmage (18405, -10.86 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 58.9 spell_power points (16.58 DPS) | yes | Marshal's Silk Leggings (231605, -1.14 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -4.31 DPS) [vendor]; Sorcerer's Leggings (226933, -5.02 DPS) [quest] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (433.6 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.14 DPS) [dungeon]; Sorcerer's Boots (22064, -12.11 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 44.8 spell_power points (12.60 DPS) | yes | Maiden's Circle (13001, -6.03 DPS) [world_drop]; Channeler's Ring (272406, -6.80 DPS) [vendor]; Blessed Band of Light (272407, -6.99 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 30.4 spell_power points (8.55 DPS) | yes | Maiden's Circle (13001, -1.98 DPS) [world_drop]; Channeler's Ring (272406, -2.75 DPS) [vendor]; Blessed Band of Light (272407, -2.94 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (433.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (433.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.43 DPS) [crafted]; Uther's Strength (11302, -2.25 DPS) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (433.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.67 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -38.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 278.9 spell_power points (78.45 DPS) | yes | Wand of Biting Cold (19108, -9.32 DPS, sim-verified) [quest]; Stormrager (16997, -14.69 DPS) [quest]; Brilliant Wand (249385, -14.84 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Sorcerer's Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 886, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 205015100000000000-23252100130133051-0050000000000000000)

Set DPS (verified): 797.0. Weights run: 1.2s. Verify run: 2.0s. 886 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.353 ± 0.026, crit=0.844 ± 0.043 per rating point (14 rating = 1%, 11.813 per %), hit=1.023 ± 0.062 per rating point (10 rating = 1%, 10.227 per %), spell_haste=not significant (0.338 ± 0.450), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 41.6 spell_power points (21.37 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.27 DPS) [pvp]; Knight-Lieutenant's Dreadweave Hat (220889, -7.61 DPS, sim-verified) [vendor] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.24 DPS) [quest]; Jewel of Kajaro (19601, -7.25 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.1 spell_power points (22.63 DPS) | yes | Mantle of the Timbermaw (19050, -5.91 DPS, sim-verified) [crafted]; Lieutenant Commander's Silk Mantle (227102, -6.88 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -7.09 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.1 spell_power points (14.91 DPS) | yes | Hide of the Wild (18510, -5.91 DPS) [crafted]; Spritecaster Cape (11623, -6.64 DPS) [dungeon]; Mageflame Cloak (13007, -12.58 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.1 spell_power points (28.76 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.69 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -8.84 DPS) [vendor]; Robes of Fiery Devastation (279266, -19.83 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.8 spell_power points (12.74 DPS) | yes | Netherflame Cuffs (254065, -4.80 DPS) [crafted]; Runecloth Cuffs (254123, -5.28 DPS) [crafted]; Phoenix Bindings (210781, -6.07 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (797.0 DPS) | yes | Marshal's Silk Gloves (16440, -2.52 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.52 DPS) [vendor]; Inferno Gloves (18408, -9.64 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.0 spell_power points (22.06 DPS) | yes | Belt of the Archmage (18405, -7.23 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -7.73 DPS) [rep]; Magician's Cord (272393, -7.96 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 59.2 spell_power points (30.35 DPS) | yes | Marshal's Silk Leggings (231605, -5.27 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -10.44 DPS) [vendor]; Sorcerer's Leggings (226933, -27.30 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 26.7 spell_power points (13.67 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -1.09 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 41.4 spell_power points (21.24 DPS) | yes | Don Julio's Band (19325, -9.94 DPS) [rep]; Maiden's Circle (13001, -10.74 DPS) [world_drop]; Blessed Band of Light (272407, -10.87 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 25.9 spell_power points (13.28 DPS) | yes | Don Julio's Band (19325, -1.97 DPS) [rep]; Maiden's Circle (13001, -2.78 DPS) [world_drop]; Blessed Band of Light (272407, -2.90 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+18.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.46 DPS) [crafted]; Uther's Strength (11302, -4.10 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -1.08 DPS) [quest]; Teebu's Blazing Longsword (1728, -51.81 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.6 spell_power points (79.84 DPS) | yes | Brilliant Wand (249385, -14.37 DPS) [crafted]; Wand of Biting Cold (19108, -15.84 DPS) [quest]; Stormrager (16997, -16.24 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 886, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 26.3. Weights run: 0.8s. Verify run: 0.6s. 136 eligible items had no known source.

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

No-known-source sample (15 of 136, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (orc, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 48.0. Weights run: 0.9s. Verify run: 0.6s. 242 eligible items had no known source.

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

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (orc, 000000000000000000-23552100130103041-0000000000000000000)

Set DPS (verified): 79.6. Weights run: 1.0s. Verify run: 0.7s. 316 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.774 ± 0.026, crit=0.219 ± 0.015 per rating point (14 rating = 1%, 3.059 per %), hit=0.376 ± 0.034 per rating point (10 rating = 1%, 3.765 per %), spell_haste=2.538 ± 0.457, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 21.2 spell_power points (2.61 DPS) | yes | Augural Shroud (2620, -0.30 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -0.86 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.17 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.2 spell_power points (1.63 DPS) | yes | Triune Amulet (7722, -0.64 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor]; Scorn's Icy Choker (23169, -1.16 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.1 spell_power points (2.11 DPS) | yes | Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -1.11 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.9 spell_power points (1.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.17 DPS) [vendor]; Icy Cloak (4327, -0.35 DPS) [crafted] |
| chest | Cindercloth Robe (10042) | Tailoring [crafted] | 27.0 spell_power points (3.33 DPS) | yes | Robe of the Magi (1716, -0.04 DPS) [world_drop]; Dreamweave Vest (10021, -0.25 DPS) [crafted]; Robe of Power (7054, -0.46 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.60 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.49 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.96 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.60 DPS) | yes | Fiery Handwraps (254025, -0.21 DPS) [crafted]; Red Mageweave Gloves (10018, -0.29 DPS) [crafted]; Crimson Silk Gloves (7064, -0.30 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (2.49 DPS) | yes | Defiler's Cloth Girdle (20166, -0.38 DPS) [rep]; Gilded Cord (254037, -0.74 DPS) [crafted]; Deathmage Sash (10771, -1.13 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.3 spell_power points (2.87 DPS) | yes | Flame Leggings (253991, -0.48 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.94 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.96 DPS) | yes | Gilded Slippers (254001, -1.43 DPS) [crafted]; Fiery Slippers (254005, -1.44 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (1.71 DPS) | yes | Advisor's Ring (19520, -0.60 DPS) [rep]; Reedknot Ring (9622, -0.85 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.9 spell_power points (1.22 DPS) | yes | Advisor's Ring (19520, -0.11 DPS) [rep]; Reedknot Ring (9622, -0.35 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (79.6 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Searing Golden Blade (12260, -0.71 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.3 spell_power points (1.89 DPS) | yes | Orb of the Forgotten Seer (7685, -0.16 DPS) [dungeon]; Thrash's Trash (276204, -0.16 DPS) [vendor]; Rod of Molten Fire (2565, -0.29 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 323.6 spell_power points (39.95 DPS) | yes | Ragefire Wand (7513, -1.13 DPS, sim-verified) [quest]; Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Cindercloth Robe; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 316, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 145.2. Weights run: 0.9s. Verify run: 0.9s. 405 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.939 ± 0.038, crit=0.343 ± 0.024 per rating point (14 rating = 1%, 4.796 per %), hit=0.622 ± 0.048 per rating point (10 rating = 1%, 6.225 per %), spell_haste=-3.557 ± 0.738, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.8 spell_power points (5.51 DPS) | yes | Dreamweave Circlet (10041, -1.08 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.13 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.57 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.5 spell_power points (2.12 DPS) | yes | Horizon Choker (13085, -0.20 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.27 DPS) [dungeon]; Mindburst Medallion (11196, -0.42 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.9 spell_power points (4.36 DPS) | yes | Netherflame Shoulders (254053, -0.79 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.26 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.29 DPS) [crafted] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.5 spell_power points (2.99 DPS) | yes | Deep Woodlands Cloak (19121, -0.01 DPS) [quest]; Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Runecloth Cloak (13860, -0.58 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.8 spell_power points (5.51 DPS) | yes | Robe of the Magi (1716, -1.48 DPS) [world_drop]; Runecloth Tunic (13857, -1.52 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.55 DPS) [vendor] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.6 spell_power points (2.85 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.87 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.4 spell_power points (4.57 DPS) | yes | Fiery Gloves (254099, -0.43 DPS) [crafted]; Raider Handwraps (272098, -0.54 DPS) [vendor]; Dreamweave Gloves (10019, -1.40 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.5 spell_power points (4.00 DPS) | yes | Dawnspire Cord (12466, -0.53 DPS) [dungeon]; Satyrmane Sash (17755, -0.59 DPS) [dungeon]; Fiery Cord (254041, -0.87 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 28.1 spell_power points (4.09 DPS) | yes | Red Mageweave Pants (10009, -0.41 DPS) [crafted]; Flame Leggings (253991, -1.09 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.15 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.5 spell_power points (4.15 DPS) | yes | Earthen Silk Slippers (254013, -0.65 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -0.84 DPS) [vendor]; Cindercloth Boots (10044, -1.09 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (2.14 DPS) | yes | Advisor's Ring (19519, -0.39 DPS) [rep]; Snake Hoop (6750, -0.58 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.90 DPS) | yes | Advisor's Ring (19519, -0.15 DPS) [rep]; Snake Hoop (6750, -0.34 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (145.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (145.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (145.2 DPS) | yes | Kindling Stave (11750, -0.63 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.78 DPS) [dungeon]; Hanzo Sword (8190, -18.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 360.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.58 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Sorcerer's Gauntlets; waist: Fiery Waistcord; legs: Stone Guard's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 405, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 421.6. Weights run: 1.0s. Verify run: 2.7s. 875 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.762 ± 0.050, crit=0.690 ± 0.040 per rating point (14 rating = 1%, 9.664 per %), hit=0.992 ± 0.076 per rating point (10 rating = 1%, 9.923 per %), spell_haste=-10.683 ± 1.080, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 49.7 spell_power points (13.98 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.50 DPS) [pvp]; Red Mageweave Headband (10033, -20.00 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (421.6 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.78 DPS) [quest]; Chains of the Lich (23125, -0.82 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.1 spell_power points (13.53 DPS) | yes | Warlord's Silk Amice (231594, -3.28 DPS) [pvp]; Champion's Silk Mantle (227104, -4.23 DPS) [pvp]; Mantle of the Timbermaw (19050, -8.05 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.0 spell_power points (9.01 DPS) | yes | Hide of the Wild (18510, -2.92 DPS) [crafted]; Mageflame Cloak (13007, -3.10 DPS) [world_drop]; Cindercloth Cloak (14044, -3.63 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.8 spell_power points (16.54 DPS) | yes | Warlord's Silk Raiment (231596, -0.90 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.27 DPS) [pvp]; Robes of Fiery Devastation (279266, -12.01 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.1 spell_power points (7.90 DPS) | yes | Runecloth Cuffs (254123, -2.67 DPS) [crafted]; Netherflame Cuffs (254065, -2.75 DPS) [crafted]; Sorcerer's Bindings (226929, -3.79 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (421.6 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Inferno Gloves (18408, -8.84 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 51.7 spell_power points (14.54 DPS) | yes | Magician's Cord (272393, -4.05 DPS) [vendor]; Belt of the Archmage (18405, -6.28 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -6.60 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 58.9 spell_power points (16.58 DPS) | yes | General's Silk Trousers (231595, -1.14 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -4.31 DPS) [pvp]; Outrider's Silk Leggings (22747, -14.44 DPS, sim-verified) [rep] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (421.6 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.14 DPS) [dungeon]; Sorcerer's Boots (22064, -7.12 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 44.8 spell_power points (12.60 DPS) | yes | Maiden's Circle (13001, -6.03 DPS) [world_drop]; Channeler's Ring (272406, -6.80 DPS) [vendor]; Blessed Band of Light (272407, -6.99 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 30.4 spell_power points (8.55 DPS) | yes | Maiden's Circle (13001, -1.98 DPS) [world_drop]; Channeler's Ring (272406, -2.75 DPS) [vendor]; Blessed Band of Light (272407, -2.94 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (421.6 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (421.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.43 DPS) [crafted]; Rune of the Guard Captain (19120, -1.98 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (421.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.67 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -42.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 278.9 spell_power points (78.45 DPS) | yes | Wand of Biting Cold (19108, -10.03 DPS, sim-verified) [quest]; Stormrager (16997, -14.69 DPS) [quest]; Brilliant Wand (249385, -14.84 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 875, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 205015100000000000-23252100130133051-0050000000000000000)

Set DPS (verified): 769.4. Weights run: 1.2s. Verify run: 1.9s. 875 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.353 ± 0.026, crit=0.844 ± 0.043 per rating point (14 rating = 1%, 11.813 per %), hit=1.023 ± 0.062 per rating point (10 rating = 1%, 10.227 per %), spell_haste=not significant (0.338 ± 0.450), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 41.6 spell_power points (21.37 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.27 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -20.01 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (769.4 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.24 DPS) [quest]; Jewel of Kajaro (19601, -6.13 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.1 spell_power points (22.63 DPS) | yes | Mantle of the Timbermaw (19050, -6.77 DPS, sim-verified) [crafted]; Champion's Silk Mantle (227104, -6.88 DPS) [pvp]; Warlord's Silk Amice (231594, -7.09 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.1 spell_power points (14.91 DPS) | yes | Hide of the Wild (18510, -5.91 DPS) [crafted]; Spritecaster Cape (11623, -6.64 DPS) [dungeon]; Mageflame Cloak (13007, -14.15 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.1 spell_power points (28.76 DPS) | yes | Warlord's Silk Raiment (231596, -2.69 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -8.84 DPS) [pvp]; Robes of Fiery Devastation (279266, -19.75 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.8 spell_power points (12.74 DPS) | yes | Netherflame Cuffs (254065, -4.80 DPS) [crafted]; Runecloth Cuffs (254123, -5.28 DPS) [crafted]; Phoenix Bindings (210781, -6.07 DPS) [crafted] |
| hands | Inferno Gloves (18408) | Tailoring [crafted] | 36.2 spell_power points (18.56 DPS) | yes | Gloves of Spell Mastery (14146, -0.01 DPS) [crafted]; General's Silk Handguards (16540, -2.53 DPS) [vendor]; General's Silk Gauntlets (231599, -2.53 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.0 spell_power points (22.06 DPS) | yes | Defiler's Cloth Girdle (20163, -7.73 DPS) [rep]; Belt of the Archmage (18405, -7.88 DPS, sim-verified) [crafted]; Magician's Cord (272393, -7.96 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 59.2 spell_power points (30.35 DPS) | yes | General's Silk Trousers (231595, -5.27 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -10.44 DPS) [pvp]; Sorcerer's Leggings (226933, -22.07 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 26.7 spell_power points (13.67 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.09 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 41.4 spell_power points (21.24 DPS) | yes | Don Julio's Band (19325, -9.94 DPS) [rep]; Maiden's Circle (13001, -10.74 DPS) [world_drop]; Blessed Band of Light (272407, -10.87 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 25.9 spell_power points (13.28 DPS) | yes | Don Julio's Band (19325, -1.97 DPS) [rep]; Maiden's Circle (13001, -2.78 DPS) [world_drop]; Blessed Band of Light (272407, -2.90 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (769.4 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (769.4 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.46 DPS) [crafted]; Rune of the Guard Captain (19120, -3.51 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (769.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -1.08 DPS) [quest]; Teebu's Blazing Longsword (1728, -54.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.6 spell_power points (79.84 DPS) | yes | Brilliant Wand (249385, -14.37 DPS) [crafted]; Wand of Biting Cold (19108, -15.84 DPS) [quest]; Stormrager (16997, -16.24 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Inferno Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 875, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

