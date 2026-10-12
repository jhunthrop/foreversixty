# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 28.2. Weights run: 1.3s. Verify run: 1.3s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.340 ± 0.016, crit=0.133 ± 0.006 per rating point (14 rating = 1%, 1.855 per %), hit=0.451 ± 0.020 per rating point (10 rating = 1%, 4.506 per %), spell_haste=not significant (0.497 ± 0.185), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.71 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Shadow Goggles (4373, -0.60 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.52 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.11 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Feyscale Cloak (6632, -0.06 DPS) [dungeon]; Black Whelp Cloak (7283, -0.06 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 10.7 spell_power points (0.69 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.32 DPS) [crafted]; Green Woolen Robe (6243, -0.43 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.7 spell_power points (0.11 DPS) | yes | Windsong Bangles (263336, -0.05 DPS) [quest]; Repurposed Hair Band (281256, -0.07 DPS) [quest]; Bright Bracers (3647, -0.56 DPS, sim-verified) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.58 DPS) | yes | Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world]; Flame Gloves (253917, -0.22 DPS, sim-verified) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.4 spell_power points (0.54 DPS) | yes | Pristine Sash (253925, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Novice Ardent's Sash (253887, -0.34 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.0 spell_power points (0.84 DPS) | yes | Phoenix Pants (4317, -0.11 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted]; Abomination Skin Leggings (23173, -0.37 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.54 DPS) | yes | Feather Padded Treads (285345, -0.22 DPS) [world]; Pristine Boots (253889, -0.28 DPS) [crafted]; Flame Boots (253893, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.37 DPS) | yes | Sludge-Stained Band (286535, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.30 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.32 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop]; Sludge-Stained Band (286535, -0.50 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.4 spell_power points (0.22 DPS) | yes | Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world]; Staff of Westfall (2042, -0.11 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 349.0 spell_power points (22.50 DPS) | yes | Skycaller (12984, -0.95 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.21 DPS) [dungeon]; Deepblaze (279896, -3.97 DPS) [quest] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 53.0. Weights run: 1.4s. Verify run: 1.6s. 260 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.690 ± 0.018, crit=0.164 ± 0.010 per rating point (14 rating = 1%, 2.301 per %), hit=0.366 ± 0.024 per rating point (10 rating = 1%, 3.660 per %), spell_haste=not significant (0.657 ± 0.273), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.32 DPS) | yes | Enchanter's Cowl (4322, -0.10 DPS) [crafted]; Holy Shroud (2721, -0.28 DPS) [world_drop]; Flame Circlet (253953, -0.28 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.1 spell_power points (1.05 DPS) | yes | Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.91 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.2 spell_power points (1.43 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Cloak of Rot (4462, -0.02 DPS) [world]; Darkspear Raider's Cloak (272078, -0.02 DPS) [vendor]; Vine Pruner's Cloak (279835, -0.53 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.0 spell_power points (1.69 DPS) | yes | Flame Gown (253965, -0.01 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon]; Death Speaker Robes (6682, -0.32 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.22 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -0.70 DPS) [world_drop]; Nightsky Wristbands (6407, -0.83 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.6 spell_power points (1.09 DPS) | yes | Phoenix Gloves (4331, -0.24 DPS) [crafted]; Truefaith Gloves (7049, -0.43 DPS) [crafted]; Flame Gloves (253917, -0.62 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.1 spell_power points (1.23 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Crimson Silk Belt (7055, -0.21 DPS) [crafted]; Invoker's Cord (215366, -0.25 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.8 spell_power points (1.77 DPS) | yes | Abomination Skin Leggings (23173, -0.41 DPS) [dungeon]; Filigreed Flame Leggings (253941, -0.55 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.55 DPS) [world] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.8 spell_power points (1.68 DPS) | yes | Gilded Slippers (254001, -0.56 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.20 DPS) [world]; Snake Hoop (6750, -0.20 DPS) [quest]; Minor Channeling Ring (1449, -1.28 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.11 DPS) [world]; Snake Hoop (6750, -0.11 DPS) [quest]; Minor Channeling Ring (1449, -0.65 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Hardened Root Staff (1317, -1.37 DPS, sim-verified) [quest]; Glimmering Staff (249392, -1.77 DPS) [crafted]; Twisted Chanter's Staff (890, -1.84 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 356.3 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.37 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Scorching Wand (5213, -4.08 DPS) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 000000000000000000-23552100130103041-0000000000000000000)

Set DPS (verified): 81.1. Weights run: 1.5s. Verify run: 1.7s. 334 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.794 ± 0.037, crit=0.206 ± 0.021 per rating point (14 rating = 1%, 2.884 per %), hit=0.417 ± 0.050 per rating point (10 rating = 1%, 4.168 per %), spell_haste=2.677 ± 0.663, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.58 DPS) | yes | Augural Shroud (2620, -0.25 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -0.82 DPS) [dungeon]; Filigreed Flame Circlet (253979, -0.86 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (1.64 DPS) | yes | Scorn's Icy Choker (23169, -0.79 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.96 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.3 spell_power points (2.13 DPS) | yes | Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -0.74 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.0 spell_power points (1.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.15 DPS) [vendor]; Icy Cloak (4327, -0.36 DPS) [crafted] |
| chest | Cindercloth Robe (10042) | Tailoring [crafted] | 27.0 spell_power points (3.31 DPS) | yes | Robe of the Magi (1716, -0.03 DPS) [world_drop]; Dreamweave Vest (10021, -0.23 DPS) [crafted]; Robe of Power (7054, -0.43 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.59 DPS) | yes | Arcane Runed Bracers (4744, -0.49 DPS) [quest]; Spidertank Oilrag (9448, -0.49 DPS) [dungeon]; Windchaser Cuffs (14429, -0.72 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (2.60 DPS) | yes | Fiery Handwraps (254025, -0.20 DPS) [crafted]; Red Mageweave Gloves (10018, -0.27 DPS) [crafted]; Crimson Silk Gloves (7064, -0.30 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.4 spell_power points (2.50 DPS) | yes | Highlander's Cloth Girdle (20098, -0.39 DPS) [rep]; Gilded Cord (254037, -0.74 DPS) [crafted]; Deathmage Sash (10771, -0.78 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (2.89 DPS) | yes | Flame Leggings (253991, -0.49 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.64 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.95 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.94 DPS) | yes | Fiery Slippers (254005, -1.08 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.40 DPS) [crafted]; Acidic Walkers (9454, -1.55 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.71 DPS) | yes | Ring of Forlorn Spirits (2043, -0.73 DPS) [quest]; Reedknot Ring (9622, -0.86 DPS) [quest]; Minor Channeling Ring (1449, -0.91 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.10 DPS) | yes | Reedknot Ring (9622, -0.25 DPS) [quest]; Minor Channeling Ring (1449, -0.30 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.53 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (81.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Searing Golden Blade (12260, -0.71 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.4 spell_power points (1.89 DPS) | yes | Orb of the Forgotten Seer (7685, -0.17 DPS) [dungeon]; Thrash's Trash (276204, -0.17 DPS) [vendor]; Orb of Lorica (11262, -0.18 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 325.8 spell_power points (39.96 DPS) | yes | Ragefire Wand (7513, -1.43 DPS) [quest]; Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Cindercloth Robe; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 334, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 151.1. Weights run: 1.5s. Verify run: 2.2s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.944 ± 0.056, crit=0.317 ± 0.032 per rating point (14 rating = 1%, 4.437 per %), hit=0.682 ± 0.072 per rating point (10 rating = 1%, 6.817 per %), spell_haste=not significant (-3.785 ± 1.019), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.9 spell_power points (5.50 DPS) | yes | Dreamweave Circlet (10041, -1.08 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.18 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.58 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.6 spell_power points (2.11 DPS) | yes | Horizon Choker (13085, -0.19 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.27 DPS) [dungeon]; Mindburst Medallion (11196, -0.42 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 30.0 spell_power points (4.36 DPS) | yes | Netherflame Shoulders (254053, -0.79 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.28 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.32 DPS) [vendor] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.6 spell_power points (2.99 DPS) | yes | Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Runecloth Cloak (13860, -0.58 DPS) [crafted]; Cloak of Fire (14134, -0.95 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.9 spell_power points (5.50 DPS) | yes | Robe of the Magi (1716, -1.48 DPS) [world_drop]; Runecloth Tunic (13857, -1.53 DPS) [crafted]; Runecloth Robe (13858, -1.57 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.6 spell_power points (2.85 DPS) | yes | Nethergeld Cuffs (254061, -0.87 DPS) [crafted]; Bloodband Bracers (11469, -0.89 DPS) [quest]; Phoenix Bindings (210781, -0.96 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 32.0 spell_power points (4.65 DPS) | yes | Fiery Gloves (254099, -0.52 DPS) [crafted]; Raider Handwraps (272098, -0.62 DPS) [vendor]; Dreamweave Gloves (10019, -1.49 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.5 spell_power points (3.99 DPS) | yes | Dawnspire Cord (12466, -0.52 DPS) [dungeon]; Satyrmane Sash (17755, -0.59 DPS) [dungeon]; Fiery Cord (254041, -0.86 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 27.8 spell_power points (4.03 DPS) | yes | Red Mageweave Pants (10009, -0.35 DPS) [crafted]; Flame Leggings (253991, -1.04 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.09 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.5 spell_power points (4.14 DPS) | yes | Earthen Silk Slippers (254013, -0.65 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -0.75 DPS) [vendor]; Cindercloth Boots (10044, -1.09 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (2.14 DPS) | yes | Lorekeeper's Ring (19523, -0.40 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.98 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.89 DPS) | yes | Lorekeeper's Ring (19523, -0.15 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.73 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (151.1 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (151.1 DPS) | yes | Mark of the Chosen (17774, -2.46 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (151.1 DPS) | yes | Kindling Stave (11750, -0.64 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.78 DPS) [dungeon]; Hanzo Sword (8190, -18.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 361.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.59 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Sorcerer's Gauntlets; waist: Fiery Waistcord; legs: Knight's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 433.6. Weights run: 1.2s. Verify run: 5.7s. 876 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.011, intellect=0.772 ± 0.069, crit=0.669 ± 0.052 per rating point (14 rating = 1%, 9.363 per %), hit=1.020 ± 0.110 per rating point (10 rating = 1%, 10.196 per %), spell_haste=-11.105 ± 1.531, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 49.7 spell_power points (14.07 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.53 DPS) [pvp]; Red Mageweave Headband (10033, -26.83 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (433.6 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.79 DPS) [quest]; Chains of the Lich (23125, -0.86 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 47.9 spell_power points (13.59 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.22 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -4.28 DPS) [pvp]; Mantle of the Timbermaw (19050, -7.23 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.4 spell_power points (9.17 DPS) | yes | Hide of the Wild (18510, -3.02 DPS) [crafted]; Mageflame Cloak (13007, -3.22 DPS) [world_drop]; Cindercloth Cloak (14044, -3.74 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.6 spell_power points (16.61 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.89 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.29 DPS) [pvp]; Robes of Fiery Devastation (279266, -7.37 DPS, sim-verified) [crafted] |
| wrist | Sorcerer's Bindings (226929) | An Earnest Proposition [quest] | sim-verified (433.6 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (433.6 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Inferno Gloves (18408, -12.97 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.2 spell_power points (14.79 DPS) | yes | Magician's Cord (272393, -4.15 DPS) [vendor]; Highlander's Cloth Girdle (20047, -6.85 DPS) [rep]; Belt of the Archmage (18405, -10.86 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 58.4 spell_power points (16.56 DPS) | yes | Marshal's Silk Leggings (231605, -1.03 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -4.24 DPS) [pvp]; Sorcerer's Leggings (226933, -4.79 DPS) [quest] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (433.6 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.16 DPS) [dungeon]; Sorcerer's Boots (22064, -12.11 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.1 spell_power points (12.79 DPS) | yes | Channeler's Ring (272406, -6.84 DPS) [vendor]; Blessed Band of Light (272407, -7.07 DPS) [vendor]; Deep Rooted Ring (19109, -7.20 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 30.5 spell_power points (8.64 DPS) | yes | Blessed Band of Light (272407, -2.92 DPS) [vendor]; Deep Rooted Ring (19109, -3.05 DPS) [world]; Channeler's Ring (272406, -6.05 DPS, sim-verified) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (433.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (433.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.37 DPS) [crafted]; Uther's Strength (11302, -2.27 DPS) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (433.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.71 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -38.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 276.9 spell_power points (78.46 DPS) | yes | Wand of Biting Cold (19108, -9.32 DPS, sim-verified) [quest]; Stormrager (16997, -14.68 DPS) [quest]; Brilliant Wand (249385, -14.80 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Sorcerer's Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 876, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 205015100000000000-23252100130133051-0050000000000000000)

Set DPS (verified): 797.0. Weights run: 1.4s. Verify run: 4.1s. 876 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.367 ± 0.039, crit=0.859 ± 0.062 per rating point (14 rating = 1%, 12.019 per %), hit=1.130 ± 0.091 per rating point (10 rating = 1%, 11.305 per %), spell_haste=not significant (-0.073 ± 0.679), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.2 spell_power points (21.63 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.32 DPS) [pvp]; Knight-Lieutenant's Dreadweave Hat (220889, -9.99 DPS, sim-verified) [vendor] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (797.0 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Jewel of Kajaro (19601, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.14 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.5 spell_power points (22.83 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -6.90 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -7.19 DPS) [pvp]; Mantle of the Timbermaw (19050, -7.46 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.2 spell_power points (15.50 DPS) | yes | Hide of the Wild (18510, -6.44 DPS) [crafted]; Spritecaster Cape (11623, -7.20 DPS) [dungeon]; Mageflame Cloak (13007, -10.57 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.4 spell_power points (28.92 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.65 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -8.80 DPS) [pvp]; Robes of Fiery Devastation (279266, -19.68 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.9 spell_power points (12.78 DPS) | yes | Netherflame Cuffs (254065, -4.80 DPS) [crafted]; Runecloth Cuffs (254123, -5.26 DPS) [crafted]; Phoenix Bindings (210781, -6.12 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 36.7 spell_power points (18.82 DPS) | yes | Marshal's Silk Gloves (16440, -2.72 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.72 DPS) [vendor]; Inferno Gloves (18408, -9.64 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 44.4 spell_power points (22.75 DPS) | yes | Belt of the Archmage (18405, -3.33 DPS) [crafted]; Highlander's Cloth Girdle (20047, -8.29 DPS) [rep]; Magician's Cord (272393, -8.49 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 59.7 spell_power points (30.61 DPS) | yes | Marshal's Silk Leggings (231605, -5.30 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -10.48 DPS) [pvp]; Sorcerer's Leggings (226933, -26.14 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 26.9 spell_power points (13.78 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -1.13 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 42.6 spell_power points (21.84 DPS) | yes | Don Julio's Band (19325, -9.89 DPS) [rep]; Blessed Band of Light (272407, -10.92 DPS) [vendor]; Channeler's Ring (272406, -13.41 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 26.0 spell_power points (13.35 DPS) | yes | Don Julio's Band (19325, -1.39 DPS) [rep]; Blessed Band of Light (272407, -2.43 DPS) [vendor]; Channeler's Ring (272406, -4.92 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (797.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (797.0 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.96 DPS) [crafted]; Uther's Strength (11302, -4.10 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (797.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -1.00 DPS) [quest]; Teebu's Blazing Longsword (1728, -55.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.8 spell_power points (79.84 DPS) | yes | Brilliant Wand (249385, -14.32 DPS) [crafted]; Wand of Biting Cold (19108, -15.84 DPS) [quest]; Stormrager (16997, -16.20 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 876, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 26.3. Weights run: 1.3s. Verify run: 1.4s. 136 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.340 ± 0.016, crit=0.133 ± 0.006 per rating point (14 rating = 1%, 1.855 per %), hit=0.451 ± 0.020 per rating point (10 rating = 1%, 4.506 per %), spell_haste=not significant (0.497 ± 0.185), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.71 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Shadow Goggles (4373, -0.60 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.52 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.37 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Pearl-clasped Cloak (5542, -0.06 DPS) [crafted]; Feyscale Cloak (6632, -0.06 DPS) [dungeon]; Black Whelp Cloak (7283, -0.06 DPS) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 10.7 spell_power points (0.69 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.32 DPS) [crafted]; Green Woolen Robe (6243, -0.43 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.19 DPS) | yes | Mindthrust Bracers (1974, -0.08 DPS) [dungeon]; Featherbead Bracers (15452, -0.08 DPS) [quest]; Owlbeard Bracers (16981, -0.09 DPS) [quest] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.58 DPS) | yes | Flame Gloves (253917, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.4 spell_power points (0.54 DPS) | yes | Pristine Sash (253925, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Novice Ardent's Sash (253887, -0.34 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.0 spell_power points (0.84 DPS) | yes | Abomination Skin Leggings (23173, -0.09 DPS) [dungeon]; Phoenix Pants (4317, -0.11 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.54 DPS) | yes | Flame Boots (253893, -0.14 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.22 DPS) [world]; Pristine Boots (253889, -0.28 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.32 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Loop of Sacrifice (281673, -0.21 DPS) [quest]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.19 DPS) | yes | Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Volcanic Rock Ring (12053, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.4 spell_power points (0.22 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 349.0 spell_power points (22.50 DPS) | yes | Skycaller (12984, -0.65 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.21 DPS) [dungeon]; Sizzle Stick (8071, -4.53 DPS) [quest] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Tabitha's Cuffs; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 136, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (orc, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 50.1. Weights run: 1.4s. Verify run: 1.7s. 242 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.690 ± 0.018, crit=0.164 ± 0.010 per rating point (14 rating = 1%, 2.301 per %), hit=0.366 ± 0.024 per rating point (10 rating = 1%, 3.660 per %), spell_haste=not significant (0.657 ± 0.273), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.18 DPS) [world_drop]; Flame Circlet (253953, -0.18 DPS) [crafted]; Filigreed Flame Circlet (253979, -0.65 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.1 spell_power points (1.05 DPS) | yes | Darkspear Warding Pendant (272075, -0.67 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.2 spell_power points (1.43 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Mantle of Woe (7750, -0.18 DPS) [quest]; Fairywing Mantle (9536, -0.28 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Hillman's Cloak (3719, -0.05 DPS) [crafted]; Windsong Drape (15468, -0.05 DPS) [quest]; Cloak of Rot (4462, -0.66 DPS, sim-verified) [world] |
| chest | Flame Gown (253965) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.15 DPS) [dungeon]; Death Speaker Robes (6682, -0.30 DPS) [dungeon]; Green Silk Armor (7065, -0.52 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.22 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -0.70 DPS) [world_drop]; Nightsky Wristbands (6407, -0.83 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.5 spell_power points (0.89 DPS) | yes | Flame Gloves (253917, -0.04 DPS) [crafted]; Phoenix Gloves (4331, -0.04 DPS) [crafted]; Truefaith Gloves (7049, -0.22 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warsong Sash (16975, -0.01 DPS) [quest]; Crimson Silk Belt (7055, -0.02 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.70 DPS, sim-verified) [rep] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.8 spell_power points (1.77 DPS) | yes | Filigreed Flame Leggings (253941, -0.38 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.41 DPS) [dungeon]; Smoldering Pants (3073, -0.55 DPS) [world] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.8 spell_power points (1.68 DPS) | yes | Gilded Slippers (254001, -0.56 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.20 DPS) [world]; Snake Hoop (6750, -0.20 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.56 DPS) | yes | Snake Hoop (6750, -0.11 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; Black Widow Band (6199, -0.47 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 26.5 spell_power points (2.49 DPS) | yes | Twisted Chanter's Staff (890, -1.84 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.84 DPS) [quest]; Glimmering Staff (249392, -5.62 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Unstable Power Core (279847) | Source of Power [quest] | 363.0 spell_power points (34.12 DPS) | yes | Necrotic Wand (7708, -0.63 DPS) [dungeon]; Starfaller (13063, -1.00 DPS) [world_drop]; Greater Mystic Wand (217287, -4.65 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Flame Gown; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Unstable Power Core

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (orc, 000000000000000000-23552100130103041-0000000000000000000)

Set DPS (verified): 78.3. Weights run: 1.5s. Verify run: 1.8s. 312 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.794 ± 0.037, crit=0.206 ± 0.021 per rating point (14 rating = 1%, 2.884 per %), hit=0.417 ± 0.050 per rating point (10 rating = 1%, 4.168 per %), spell_haste=2.677 ± 0.663, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.58 DPS) | yes | Augural Shroud (2620, -0.25 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -0.82 DPS) [dungeon]; Filigreed Flame Circlet (253979, -0.86 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (1.64 DPS) | yes | Scorn's Icy Choker (23169, -0.19 DPS) [dungeon]; Necklace of Calisea (1714, -0.96 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.3 spell_power points (2.13 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.0 spell_power points (1.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.15 DPS) [vendor]; Icy Cloak (4327, -0.36 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (78.3 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.40 DPS) [crafted]; Cindercloth Robe (10042, -0.85 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.59 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.32 DPS) [quest]; Spidertank Oilrag (9448, -0.49 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (2.60 DPS) | yes | Fiery Handwraps (254025, -0.20 DPS) [crafted]; Red Mageweave Gloves (10018, -0.27 DPS) [crafted]; Crimson Silk Gloves (7064, -0.30 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.4 spell_power points (2.50 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.39 DPS) [rep]; Gilded Cord (254037, -0.74 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (2.89 DPS) | yes | Flame Leggings (253991, -0.49 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.64 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.95 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.94 DPS) | yes | Fiery Slippers (254005, -0.67 DPS) [crafted]; Gilded Slippers (254001, -1.40 DPS) [crafted]; Acidic Walkers (9454, -1.55 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.71 DPS) | yes | Reedknot Ring (9622, -0.86 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.98 DPS) [vendor]; Black Widow Band (6199, -1.03 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.10 DPS) | yes | Sea Giant's Toe Ring (274746, -0.37 DPS) [vendor]; Black Widow Band (6199, -0.42 DPS) [world]; Reedknot Ring (9622, -0.99 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Searing Golden Blade (12260, -0.71 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.4 spell_power points (1.89 DPS) | yes | Orb of the Forgotten Seer (7685, -0.17 DPS) [dungeon]; Thrash's Trash (276204, -0.17 DPS) [vendor]; Rod of Molten Fire (2565, -0.29 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 325.8 spell_power points (39.96 DPS) | yes | Ragefire Wand (7513, -1.43 DPS) [quest]; Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 312, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 145.2. Weights run: 1.5s. Verify run: 1.8s. 398 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.944 ± 0.056, crit=0.317 ± 0.032 per rating point (14 rating = 1%, 4.437 per %), hit=0.682 ± 0.072 per rating point (10 rating = 1%, 6.817 per %), spell_haste=not significant (-3.785 ± 1.019), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.9 spell_power points (5.50 DPS) | yes | Dreamweave Circlet (10041, -1.08 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.18 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.58 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.6 spell_power points (2.11 DPS) | yes | Horizon Choker (13085, -0.19 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.27 DPS) [dungeon]; Mindburst Medallion (11196, -0.42 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 30.0 spell_power points (4.36 DPS) | yes | Netherflame Shoulders (254053, -0.79 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.28 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.32 DPS) [vendor] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.6 spell_power points (2.99 DPS) | yes | Deep Woodlands Cloak (19121, -0.01 DPS) [quest]; Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Runecloth Cloak (13860, -0.58 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.9 spell_power points (5.50 DPS) | yes | Robe of the Magi (1716, -1.48 DPS) [world_drop]; Runecloth Tunic (13857, -1.53 DPS) [crafted]; Runecloth Robe (13858, -1.57 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.6 spell_power points (2.85 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.87 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 32.0 spell_power points (4.65 DPS) | yes | Fiery Gloves (254099, -0.52 DPS) [crafted]; Raider Handwraps (272098, -0.62 DPS) [vendor]; Dreamweave Gloves (10019, -1.49 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.5 spell_power points (3.99 DPS) | yes | Dawnspire Cord (12466, -0.52 DPS) [dungeon]; Satyrmane Sash (17755, -0.59 DPS) [dungeon]; Fiery Cord (254041, -0.86 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 27.8 spell_power points (4.03 DPS) | yes | Red Mageweave Pants (10009, -0.35 DPS) [crafted]; Flame Leggings (253991, -1.04 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.09 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.5 spell_power points (4.14 DPS) | yes | Earthen Silk Slippers (254013, -0.65 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -0.75 DPS) [vendor]; Cindercloth Boots (10044, -1.09 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (2.14 DPS) | yes | Advisor's Ring (19519, -0.40 DPS) [rep]; Reedknot Ring (9622, -1.12 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.89 DPS) | yes | Advisor's Ring (19519, -0.15 DPS) [rep]; Reedknot Ring (9622, -0.87 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (145.2 DPS) | yes | Uther's Strength (11302, -0.02 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (145.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (145.2 DPS) | yes | Kindling Stave (11750, -0.64 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.78 DPS) [dungeon]; Hanzo Sword (8190, -18.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 361.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.59 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Sorcerer's Gauntlets; waist: Fiery Waistcord; legs: Stone Guard's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 398, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 421.6. Weights run: 1.2s. Verify run: 5.7s. 866 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.011, intellect=0.772 ± 0.069, crit=0.669 ± 0.052 per rating point (14 rating = 1%, 9.363 per %), hit=1.020 ± 0.110 per rating point (10 rating = 1%, 10.196 per %), spell_haste=-11.105 ± 1.531, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 49.7 spell_power points (14.07 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.53 DPS) [pvp]; Red Mageweave Headband (10033, -20.00 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (421.6 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.79 DPS) [quest]; Chains of the Lich (23125, -0.86 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 47.9 spell_power points (13.59 DPS) | yes | Warlord's Silk Amice (231594, -3.22 DPS) [pvp]; Champion's Silk Mantle (227104, -4.28 DPS) [pvp]; Mantle of the Timbermaw (19050, -8.05 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.4 spell_power points (9.17 DPS) | yes | Hide of the Wild (18510, -3.02 DPS) [crafted]; Mageflame Cloak (13007, -3.22 DPS) [world_drop]; Cindercloth Cloak (14044, -3.74 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.6 spell_power points (16.61 DPS) | yes | Warlord's Silk Raiment (231596, -0.89 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.29 DPS) [pvp]; Robes of Fiery Devastation (279266, -12.01 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.2 spell_power points (7.98 DPS) | yes | Runecloth Cuffs (254123, -2.68 DPS) [crafted]; Netherflame Cuffs (254065, -2.77 DPS) [crafted]; Sorcerer's Bindings (226929, -3.81 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (421.6 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Inferno Gloves (18408, -8.84 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.2 spell_power points (14.79 DPS) | yes | Magician's Cord (272393, -4.15 DPS) [vendor]; Belt of the Archmage (18405, -6.28 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -6.85 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 58.4 spell_power points (16.56 DPS) | yes | General's Silk Trousers (231595, -1.03 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -4.24 DPS) [pvp]; Outrider's Silk Leggings (22747, -14.44 DPS, sim-verified) [rep] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (421.6 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.16 DPS) [dungeon]; Sorcerer's Boots (22064, -7.12 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.1 spell_power points (12.79 DPS) | yes | Channeler's Ring (272406, -6.84 DPS) [vendor]; Blessed Band of Light (272407, -7.07 DPS) [vendor]; Deep Rooted Ring (19109, -7.20 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 30.5 spell_power points (8.64 DPS) | yes | Channeler's Ring (272406, -2.69 DPS) [vendor]; Blessed Band of Light (272407, -2.92 DPS) [vendor]; Deep Rooted Ring (19109, -3.05 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (421.6 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (421.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.37 DPS) [crafted]; Rune of the Guard Captain (19120, -1.94 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (421.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.71 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -42.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 276.9 spell_power points (78.46 DPS) | yes | Wand of Biting Cold (19108, -10.03 DPS, sim-verified) [quest]; Stormrager (16997, -14.68 DPS) [quest]; Brilliant Wand (249385, -14.80 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 866, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 205015100000000000-23252100130133051-0050000000000000000)

Set DPS (verified): 786.3. Weights run: 1.4s. Verify run: 4.3s. 866 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.367 ± 0.039, crit=0.859 ± 0.062 per rating point (14 rating = 1%, 12.019 per %), hit=1.130 ± 0.091 per rating point (10 rating = 1%, 11.305 per %), spell_haste=not significant (-0.073 ± 0.679), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.2 spell_power points (21.63 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.32 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -9.52 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.14 DPS) [quest]; Jewel of Kajaro (19601, -6.01 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.5 spell_power points (22.83 DPS) | yes | Champion's Silk Mantle (227104, -6.90 DPS) [pvp]; Warlord's Silk Amice (231594, -7.19 DPS) [pvp]; Mantle of the Timbermaw (19050, -8.78 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.2 spell_power points (15.50 DPS) | yes | Hide of the Wild (18510, -6.44 DPS) [crafted]; Spritecaster Cape (11623, -7.20 DPS) [dungeon]; Mageflame Cloak (13007, -9.95 DPS, sim-verified) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.4 spell_power points (28.92 DPS) | yes | Warlord's Silk Raiment (231596, -2.65 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -8.80 DPS) [pvp]; Robes of Fiery Devastation (279266, -11.97 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.9 spell_power points (12.78 DPS) | yes | Netherflame Cuffs (254065, -4.80 DPS) [crafted]; Runecloth Cuffs (254123, -5.26 DPS) [crafted]; Phoenix Bindings (210781, -6.12 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 36.7 spell_power points (18.82 DPS) | yes | Inferno Gloves (18408, -0.21 DPS) [crafted]; General's Silk Handguards (16540, -2.72 DPS) [vendor]; General's Silk Gauntlets (231599, -2.72 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 44.4 spell_power points (22.75 DPS) | yes | Belt of the Archmage (18405, -3.33 DPS) [crafted]; Defiler's Cloth Girdle (20163, -8.29 DPS) [rep]; Magician's Cord (272393, -8.49 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 59.7 spell_power points (30.61 DPS) | yes | General's Silk Trousers (231595, -5.30 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -10.48 DPS) [pvp]; Sorcerer's Leggings (226933, -13.12 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 26.9 spell_power points (13.78 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.13 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 42.6 spell_power points (21.84 DPS) | yes | Elemental Focus Band (20682, -8.49 DPS) [world]; Blessed Band of Light (272407, -10.92 DPS) [vendor]; Channeler's Ring (272406, -13.41 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (786.3 DPS) | yes | Blessed Band of Light (272407, -1.03 DPS) [vendor]; Channeler's Ring (272406, -3.52 DPS) [vendor]; Elemental Focus Band (20682, -13.67 DPS, sim-verified) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+19.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.96 DPS) [crafted]; Rune of the Guard Captain (19120, -3.12 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -1.00 DPS) [quest]; Teebu's Blazing Longsword (1728, -52.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.8 spell_power points (79.84 DPS) | yes | Brilliant Wand (249385, -14.32 DPS) [crafted]; Wand of Biting Cold (19108, -15.84 DPS) [quest]; Stormrager (16997, -16.20 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 866, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

