# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 1531001000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 39.2. Weights run: 1.1s. Verify run: 0.7s. 145 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.294 ± 0.008, crit=0.058 ± 0.002 per rating point (14 rating = 1%, 0.807 per %), hit=0.201 ± 0.007 per rating point (10 rating = 1%, 2.015 per %), spell_haste=0.568 ± 0.100, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.236 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.81 DPS) | yes | Shadow Goggles (4373, -1.62 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.6 spell_power points (1.04 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.49 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.70 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.54 DPS) | yes | Black Whelp Cloak (7283, -0.14 DPS) [crafted]; Caretaker's Cape (20428, -0.14 DPS) [rep]; Pearl-clasped Cloak (5542, -0.15 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.88 DPS) | yes | Green Woolen Vest (2582, -0.33 DPS) [crafted]; Bloody Apron (6226, -0.33 DPS) [dungeon]; Gray Woolen Robe (2585, -1.16 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.5 spell_power points (0.20 DPS) | yes | Windsong Bangles (263336, -0.06 DPS) [quest]; Repurposed Hair Band (281256, -0.12 DPS) [quest]; Bright Bracers (3647, -1.06 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.95 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.29 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.60 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.70 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.38 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.80 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (39.2 DPS) | yes | Silk-threaded Trousers (1929, -0.10 DPS) [dungeon]; Rumpled Kilt (274741, -0.37 DPS) [vendor]; Abomination Skin Leggings (23173, -0.52 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (1.11 DPS) | yes | Pristine Boots (253889, -0.58 DPS) [crafted]; Woolen Boots (2583, -0.76 DPS) [crafted]; Red Woolen Boots (4313, -1.02 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.76 DPS) | yes | Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.64 DPS) [world_drop]; The 1 Ring (8350, -0.72 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.68 DPS) | yes | Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.64 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.9 spell_power points (0.40 DPS) | yes | Witching Stave (1484, -0.20 DPS) [dungeon]; Staff of Westfall (2042, -0.20 DPS) [quest]; Lesser Staff of the Spire (1300, -0.62 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 167.6 spell_power points (22.72 DPS) | yes | Skycaller (12984, -0.51 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.19 DPS) [quest]; Sizzle Stick (8071, -4.39 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 145, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 1531005113100000000-00000000000000000-000000000000000000)

Set DPS (verified): 92.8. Weights run: 1.2s. Verify run: 0.8s. 264 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.600 ± 0.019, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.986 per %), hit=0.270 ± 0.028 per rating point (10 rating = 1%, 2.701 per %), spell_haste=-7.141 ± 0.517, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.541 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (92.8 DPS) | yes | Silk Headband (7050, -0.38 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.57 DPS) [crafted]; Enchanter's Cowl (4322, -2.93 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.6 spell_power points (2.02 DPS) | yes | Darkspear Warding Pendant (272075, -1.45 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.56 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.88 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (2.75 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.57 DPS) [quest]; Magician's Mantle (12998, -0.76 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.95 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272078, -0.04 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.8 spell_power points (3.21 DPS) | yes | Beguiler Robes (7728, -0.11 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.31 DPS) [dungeon]; Death Speaker Robes (6682, -0.61 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.72 DPS) | yes | Nightsky Wristbands (6407, -1.03 DPS) [world_drop]; Stonecloth Bindings (14416, -1.15 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.89 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.2 spell_power points (1.37 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Shilly Mitts (9609, -0.04 DPS) [quest]; Truefaith Gloves (7049, -0.08 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.8 spell_power points (2.44 DPS) | yes | Belt of Arugal (6392, -0.38 DPS) [dungeon]; Crimson Silk Belt (7055, -0.50 DPS) [crafted]; Invoker's Cord (215366, -0.53 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.8 spell_power points (2.63 DPS) | yes | Gaze Dreamer Pants (6903, -0.34 DPS) [dungeon]; Pristine Leggings (253987, -0.50 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.80 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.2 spell_power points (2.14 DPS) | yes | Spidersilk Boots (4320, -0.34 DPS) [crafted]; Glinteye Slippers (273024, -0.80 DPS) [dungeon]; Acidic Walkers (9454, -2.31 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.72 DPS) | yes | Minor Channeling Ring (1449, -0.53 DPS) [quest]; Ring of Power Regulation (273030, -0.57 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.57 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.34 DPS) | yes | Minor Channeling Ring (1449, -0.15 DPS) [quest]; Ring of Power Regulation (273030, -0.19 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.19 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (6.87 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.05 DPS) [quest]; Glimmering Staff (249392, -5.61 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 174.6 spell_power points (33.32 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.36 DPS) [crafted]; Gravestone Scepter (7001, -4.32 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 1531005113100001531-00000000000000000-000000000000000000)

Set DPS (verified): 141.0. Weights run: 1.3s. Verify run: 0.9s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.817 ± 0.043, crit=0.160 ± 0.005 per rating point (14 rating = 1%, 2.238 per %), hit=0.553 ± 0.057 per rating point (10 rating = 1%, 5.534 per %), spell_haste=-14.786 ± 1.089, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.570 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 21.9 spell_power points (4.46 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.18 DPS) [crafted]; Augural Shroud (2620, -0.55 DPS) [world]; Whitemane's Chapeau (7720, -1.37 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.5 spell_power points (2.76 DPS) | yes | Scorn's Icy Choker (23169, -0.33 DPS) [dungeon]; Triune Amulet (7722, -1.13 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.59 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.6 spell_power points (3.59 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.50 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.1 spell_power points (2.05 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.22 DPS) [vendor]; Icy Cloak (4327, -0.63 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.9 spell_power points (5.48 DPS) | yes | Dreamweave Vest (10021, -0.32 DPS) [crafted]; Robe of Power (7054, -0.63 DPS) [crafted]; Elemental Raiment (9434, -1.20 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.83 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Windchaser Cuffs (14429, -0.34 DPS) [world_drop]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (4.33 DPS) | yes | Red Mageweave Gloves (10018, -0.43 DPS) [crafted]; Black Mageweave Gloves (10003, -1.28 DPS) [crafted]; Stormcloth Gloves (10011, -1.52 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.3 spell_power points (3.92 DPS) | yes | Highlander's Cloth Girdle (20098, -0.40 DPS) [rep]; Gilded Cord (254037, -0.96 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.8 spell_power points (4.85 DPS) | yes | Crimson Silk Pantaloons (7062, -1.06 DPS) [crafted]; Abomination Skin Leggings (23173, -1.68 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.09 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Gilded Slippers (254001, -2.04 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.54 DPS) [dungeon]; Spidersilk Boots (4320, -2.80 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.87 DPS) | yes | Lorekeeper's Ring (19524, -1.04 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.24 DPS) [quest]; Reedknot Ring (9622, -1.44 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.1 spell_power points (2.05 DPS) | yes | Lorekeeper's Ring (19524, -0.22 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.42 DPS) [quest]; Reedknot Ring (9622, -0.63 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (141.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.33 DPS) [dungeon]; Spine of the Devourer (273040, -2.96 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.5 spell_power points (3.15 DPS) | yes | Orb of Lorica (11262, -0.26 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.30 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.72 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 198.3 spell_power points (40.38 DPS) | yes | Nether Force Wand (11263, -2.55 DPS) [quest]; Icefury Wand (7514, -2.69 DPS) [quest]; Ragefire Wand (7513, -2.74 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 1531005113100001531-03202300000000000-000000000000000000)

Set DPS (verified): 238.9. Weights run: 1.2s. Verify run: 1.2s. 436 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.060 ± 0.065, crit=0.239 ± 0.007 per rating point (14 rating = 1%, 3.346 per %), hit=0.702 ± 0.090 per rating point (10 rating = 1%, 7.021 per %), spell_haste=-21.271 ± 1.665, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.585 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.2 spell_power points (7.90 DPS) | yes | Dreamweave Circlet (10041, -1.69 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.99 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.59 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.5 spell_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.13 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.42 DPS) [dungeon]; Mindburst Medallion (11196, -0.61 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.1 spell_power points (6.31 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -2.20 DPS) [vendor]; Inquisitor's Shawl (19507, -2.22 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.17 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.4 spell_power points (4.00 DPS) | yes | Runecloth Cloak (13860, -0.57 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.09 DPS) [vendor]; Big Voodoo Cloak (8216, -1.14 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.2 spell_power points (7.90 DPS) | yes | Runecloth Tunic (13857, -2.27 DPS) [crafted]; Robe of the Magi (1716, -2.33 DPS) [world_drop]; Runecloth Robe (13858, -3.81 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.5 spell_power points (2.86 DPS) | yes | Nethergeld Cuffs (254061, -0.02 DPS) [crafted]; Forgotten Wraps (9433, -0.36 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.36 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 33.9 spell_power points (6.66 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -2.23 DPS) [vendor]; Dreamweave Gloves (10019, -2.28 DPS) [crafted]; Raider Handwraps (272098, -3.67 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.1 spell_power points (5.14 DPS) | yes | Satyrmane Sash (17755, -0.30 DPS) [dungeon]; Deathmage Sash (10771, -0.64 DPS) [dungeon]; Runecloth Belt (13856, -1.07 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 28.1 spell_power points (5.52 DPS) | yes | Red Mageweave Pants (10009, -0.26 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.18 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.24 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 24.6 spell_power points (4.83 DPS) | yes | Earthen Silk Slippers (254013, -0.11 DPS) [crafted]; Gilded Sandals (254107, -0.79 DPS) [crafted]; Southsea Mojo Boots (20641, -0.96 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.3 spell_power points (3.01 DPS) | yes | Lorekeeper's Ring (19523, -0.65 DPS) [rep]; Snake Hoop (6750, -0.79 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.56 DPS) | yes | Lorekeeper's Ring (19523, -0.20 DPS) [rep]; Snake Hoop (6750, -0.33 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (238.9 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (238.9 DPS) | yes | Mark of the Chosen (17774, -3.86 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (238.9 DPS) | yes | Spellshifter Rod (9527, -0.46 DPS) [quest]; Kindling Stave (11750, -1.58 DPS) [dungeon]; Hanzo Sword (8190, -40.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 267.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.52 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.18 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 1531005113100001531-03202300000000000-053200000000000000)

Set DPS (verified): 443.1. Weights run: 1.3s. Verify run: 3.6s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.093 ± 0.070, crit=0.399 ± 0.011 per rating point (14 rating = 1%, 5.591 per %), hit=0.932 ± 0.107 per rating point (10 rating = 1%, 9.322 per %), spell_haste=-30.393 ± 1.876, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.527 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 53.9 spell_power points (12.76 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.81 DPS) [pvp]; Magister's Crown (16686, -17.05 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (443.1 DPS) | yes | Beads of Ogre Mojo (22149, -0.73 DPS) [quest]; Pebble of Kajaro (19600, -1.42 DPS) [quest]; Jewel of Kajaro (19601, -1.42 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 49.0 spell_power points (11.59 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.80 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.88 DPS) [crafted]; Darkspear Shoulderpads (272103, -4.73 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 34.1 spell_power points (8.06 DPS) | yes | Hide of the Wild (18510, -2.16 DPS) [crafted]; Spritecaster Cape (11623, -3.20 DPS) [dungeon]; Runecloth Cloak (13860, -3.86 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.7 spell_power points (13.89 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.36 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.20 DPS) [pvp]; Sorcerer's Robes (226932, -8.47 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 30.7 spell_power points (7.28 DPS) | yes | Runecloth Cuffs (254123, -2.09 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.36 DPS) [pvp]; Magiskull Cuffs (13107, -2.88 DPS) [world_drop] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (443.1 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Raider Handwraps (272097, -13.45 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.4 spell_power points (13.81 DPS) | yes | Belt of the Archmage (18405, -3.62 DPS) [crafted]; Magister's Belt (16685, -6.72 DPS) [dungeon]; Magician's Cord (272393, -9.28 DPS, sim-verified) [vendor] |
| legs | Sorcerer's Leggings (226933) | Anthion's Parting Words [quest] | sim-verified (443.1 DPS) | yes | Sentinel's Silk Leggings (22752, +0.00 DPS) [rep]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (443.1 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.27 DPS) [dungeon]; Sorcerer's Boots (22064, -12.08 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 47.2 spell_power points (11.16 DPS) | yes | Maiden's Circle (13001, -5.09 DPS) [world_drop]; Channeler's Ring (272406, -5.33 DPS) [vendor]; Deep Rooted Ring (19109, -5.73 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 34.0 spell_power points (8.05 DPS) | yes | Maiden's Circle (13001, -1.98 DPS) [world_drop]; Channeler's Ring (272406, -2.22 DPS) [vendor]; Deep Rooted Ring (19109, -2.63 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (443.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (443.1 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.33 DPS) [crafted]; Uther's Strength (11302, -1.89 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (443.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.61 DPS) [quest]; Teebu's Blazing Longsword (1728, -48.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 330.3 spell_power points (78.18 DPS) | yes | Wand of Biting Cold (19108, -7.25 DPS, sim-verified) [quest]; Stormrager (16997, -14.20 DPS) [quest]; Brilliant Wand (249385, -14.66 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sorcerer's Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 1531005113100101531-03000000000000000-053500030000000000)

Set DPS (verified): 662.3. Weights run: 1.5s. Verify run: 2.7s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.453 ± 0.013, crit=0.465 ± 0.012 per rating point (14 rating = 1%, 6.512 per %), hit=0.996 ± 0.038 per rating point (10 rating = 1%, 9.959 per %), spell_haste=-13.693 ± 0.827, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.475 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 38.8 spell_power points (17.45 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.43 DPS) [pvp]; Red Mageweave Headband (10033, -7.32 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (662.3 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Jewel of Kajaro (19601, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.50 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 40.3 spell_power points (18.11 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.83 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.34 DPS, sim-verified) [crafted]; Lieutenant Commander's Silk Mantle (227102, -6.21 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.6 spell_power points (13.29 DPS) | yes | Hide of the Wild (18510, -4.97 DPS) [crafted]; Spritecaster Cape (11623, -5.78 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -7.00 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 51.9 spell_power points (23.35 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.13 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -7.52 DPS) [pvp]; Acumen Robes (17775, -13.51 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.6 spell_power points (11.52 DPS) | yes | Runecloth Cuffs (254123, -4.54 DPS) [crafted]; Sorcerer's Bindings (226929, -6.33 DPS) [quest]; Bracers of Hope (22667, -6.58 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.3 spell_power points (13.15 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 44.9 spell_power points (20.19 DPS) | yes | Belt of the Archmage (18405, -5.02 DPS) [crafted]; Magician's Cord (272393, -6.76 DPS) [vendor]; Highlander's Cloth Girdle (20047, -9.75 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.6 spell_power points (22.27 DPS) | yes | Marshal's Silk Leggings (231605, -1.79 DPS) [pvp]; Sorcerer's Leggings (226933, -6.00 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -6.44 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 28.2 spell_power points (12.69 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -1.22 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 42.0 spell_power points (18.89 DPS) | yes | Maiden's Circle (13001, -9.38 DPS) [world_drop]; Blessed Band of Light (272407, -9.92 DPS) [vendor]; Deep Rooted Ring (19109, -11.46 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 27.0 spell_power points (12.13 DPS) | yes | Maiden's Circle (13001, -2.61 DPS) [world_drop]; Blessed Band of Light (272407, -3.16 DPS) [vendor]; Deep Rooted Ring (19109, -4.70 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (662.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (662.3 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.26 DPS) [crafted]; Uther's Strength (11302, -3.60 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (662.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.46 DPS) [quest]; Teebu's Blazing Longsword (1728, -42.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 176.8 spell_power points (79.46 DPS) | yes | Brilliant Wand (249385, -14.41 DPS) [crafted]; Wand of Biting Cold (19108, -15.46 DPS) [quest]; Stormrager (16997, -15.75 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (orc, 1531001000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.8. Weights run: 1.1s. Verify run: 0.7s. 134 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.294 ± 0.008, crit=0.058 ± 0.002 per rating point (14 rating = 1%, 0.807 per %), hit=0.201 ± 0.007 per rating point (10 rating = 1%, 2.015 per %), spell_haste=0.568 ± 0.100, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.236 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.81 DPS) | yes | Shadow Goggles (4373, -1.51 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.6 spell_power points (1.04 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.49 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.74 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.54 DPS) | yes | Battle Healer's Cloak (20427, -0.14 DPS) [rep]; Pearl-clasped Cloak (5542, -0.15 DPS) [crafted]; Black Whelp Cloak (7283, -0.57 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.88 DPS) | yes | Green Woolen Vest (2582, -0.33 DPS) [crafted]; Bloody Apron (6226, -0.33 DPS) [dungeon]; Gray Woolen Robe (2585, -1.17 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.41 DPS) | yes | Owlbeard Bracers (16981, -0.19 DPS) [quest]; Mindthrust Bracers (1974, -0.21 DPS) [dungeon]; Featherbead Bracers (15452, -0.21 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.95 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.29 DPS) [crafted]; Apothecary Gloves (10919, -0.41 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.70 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.38 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.97 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.3 spell_power points (1.54 DPS) | yes | Filigreed Pristine Leggings (253937, -0.49 DPS) [crafted]; Silk-threaded Trousers (1929, -0.59 DPS) [dungeon]; Rumpled Kilt (274741, -0.86 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (1.11 DPS) | yes | Pristine Boots (253889, -0.58 DPS) [crafted]; Woolen Boots (2583, -0.76 DPS) [crafted]; Red Woolen Boots (4313, -0.95 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.68 DPS) | yes | Loop of Sacrifice (281673, -0.48 DPS) [quest]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; The 1 Ring (8350, -0.64 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.8 spell_power points (0.24 DPS) | yes | Loop of Sacrifice (281673, -0.04 DPS) [quest]; Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; The 1 Ring (8350, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.9 spell_power points (0.40 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.16 DPS) [world]; Clear Crystal Rod (16894, -0.20 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 167.6 spell_power points (22.72 DPS) | yes | Skycaller (12984, -0.65 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.39 DPS) [quest]; Greater Magic Wand (11288, -4.95 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 134, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (orc, 1531005113100000000-00000000000000000-000000000000000000)

Set DPS (verified): 88.8. Weights run: 1.2s. Verify run: 0.8s. 245 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.600 ± 0.019, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.986 per %), hit=0.270 ± 0.028 per rating point (10 rating = 1%, 2.701 per %), spell_haste=-7.141 ± 0.517, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.541 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (88.8 DPS) | yes | Silk Headband (7050, -0.38 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.57 DPS) [crafted]; Enchanter's Cowl (4322, -2.09 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.6 spell_power points (2.02 DPS) | yes | Darkspear Warding Pendant (272075, -1.45 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.56 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.81 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (2.75 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.57 DPS) [quest]; Magician's Mantle (12998, -0.76 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.04 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.8 spell_power points (3.21 DPS) | yes | Beguiler Robes (7728, -0.11 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.31 DPS) [dungeon]; Death Speaker Robes (6682, -0.61 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.72 DPS) | yes | Nightsky Wristbands (6407, -1.03 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.15 DPS) [quest]; Glowing Magical Bracelets (13106, -1.61 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.0 spell_power points (1.72 DPS) | yes | Serpent Gloves (5970, -0.38 DPS) [dungeon]; Truefaith Gloves (7049, -0.42 DPS) [crafted]; Gnoll Casting Gloves (892, -0.57 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.8 spell_power points (2.44 DPS) | yes | Warsong Sash (16975, -0.34 DPS) [quest]; Belt of Arugal (6392, -0.38 DPS) [dungeon]; Crimson Silk Belt (7055, -0.50 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.8 spell_power points (2.63 DPS) | yes | Gaze Dreamer Pants (6903, -0.34 DPS) [dungeon]; Pristine Leggings (253987, -0.50 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.80 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.2 spell_power points (2.14 DPS) | yes | Spidersilk Boots (4320, -0.34 DPS) [crafted]; Glinteye Slippers (273024, -0.80 DPS) [dungeon]; Acidic Walkers (9454, -2.46 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.72 DPS) | yes | Ring of Power Regulation (273030, -0.57 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.57 DPS) [vendor]; Solid Gold Band (279845, -0.95 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.34 DPS) | yes | Ring of Power Regulation (273030, -0.19 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.19 DPS) [vendor]; Solid Gold Band (279845, -0.57 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (6.87 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -5.61 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -5.73 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 174.6 spell_power points (33.32 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.36 DPS) [crafted]; Unstable Power Core (279847, -3.61 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 245, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (orc, 1531005113100001531-00000000000000000-000000000000000000)

Set DPS (verified): 137.6. Weights run: 1.3s. Verify run: 1.0s. 320 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.817 ± 0.043, crit=0.160 ± 0.005 per rating point (14 rating = 1%, 2.238 per %), hit=0.553 ± 0.057 per rating point (10 rating = 1%, 5.534 per %), spell_haste=-14.786 ± 1.089, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.570 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 21.9 spell_power points (4.46 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.18 DPS) [crafted]; Augural Shroud (2620, -0.55 DPS) [world]; Whitemane's Chapeau (7720, -1.37 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.5 spell_power points (2.76 DPS) | yes | Scorn's Icy Choker (23169, -0.33 DPS) [dungeon]; Triune Amulet (7722, -1.13 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.59 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.6 spell_power points (3.59 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.50 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.1 spell_power points (2.05 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.22 DPS) [vendor]; Icy Cloak (4327, -0.63 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.9 spell_power points (5.48 DPS) | yes | Dreamweave Vest (10021, -0.32 DPS) [crafted]; Robe of Power (7054, -0.63 DPS) [crafted]; Elemental Raiment (9434, -1.20 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (137.6 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Windchaser Cuffs (14429, -0.34 DPS) [world_drop]; Radiant Silver Bracers (4545, -2.22 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (4.33 DPS) | yes | Red Mageweave Gloves (10018, -0.43 DPS) [crafted]; Black Mageweave Gloves (10003, -1.28 DPS) [crafted]; Stormcloth Gloves (10011, -1.52 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.3 spell_power points (3.92 DPS) | yes | Defiler's Cloth Girdle (20166, -0.40 DPS) [rep]; Gilded Cord (254037, -0.96 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.8 spell_power points (4.85 DPS) | yes | Crimson Silk Pantaloons (7062, -1.06 DPS) [crafted]; Abomination Skin Leggings (23173, -1.68 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.09 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Gilded Slippers (254001, -2.30 DPS) [crafted]; Acidic Walkers (9454, -2.54 DPS) [dungeon]; Spidersilk Boots (4320, -2.80 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.87 DPS) | yes | Advisor's Ring (19520, -1.04 DPS) [rep]; Reedknot Ring (9622, -1.44 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.1 spell_power points (2.05 DPS) | yes | Advisor's Ring (19520, -0.22 DPS) [rep]; Reedknot Ring (9622, -0.63 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.33 DPS) [dungeon]; Spine of the Devourer (273040, -2.96 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.5 spell_power points (3.15 DPS) | yes | Orb of the Forgotten Seer (7685, -0.30 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.72 DPS) [crafted]; Omega Orb (7749, -1.31 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 198.3 spell_power points (40.38 DPS) | yes | Nether Force Wand (11263, -2.55 DPS) [quest]; Icefury Wand (7514, -2.69 DPS) [quest]; Ragefire Wand (7513, -2.74 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 320, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 1531005113100001531-03202300000000000-000000000000000000)

Set DPS (verified): 229.1. Weights run: 1.2s. Verify run: 1.1s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.060 ± 0.065, crit=0.239 ± 0.007 per rating point (14 rating = 1%, 3.346 per %), hit=0.702 ± 0.090 per rating point (10 rating = 1%, 7.021 per %), spell_haste=-21.271 ± 1.665, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.585 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.2 spell_power points (7.90 DPS) | yes | Dreamweave Circlet (10041, -1.69 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.99 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.59 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.5 spell_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.13 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.42 DPS) [dungeon]; Mindburst Medallion (11196, -0.61 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.1 spell_power points (6.31 DPS) | yes | Red Mageweave Shoulders (10029, -1.80 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.20 DPS) [vendor]; Inquisitor's Shawl (19507, -2.22 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.5 spell_power points (4.23 DPS) | yes | Spritecaster Cape (11623, -0.23 DPS) [dungeon]; Runecloth Cloak (13860, -0.80 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.32 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.2 spell_power points (7.90 DPS) | yes | Runecloth Tunic (13857, -2.27 DPS) [crafted]; Robe of the Magi (1716, -2.33 DPS) [world_drop]; Runecloth Robe (13858, -2.94 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.5 spell_power points (2.86 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.02 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 33.9 spell_power points (6.66 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -2.23 DPS) [vendor]; Dreamweave Gloves (10019, -2.28 DPS) [crafted]; Raider Handwraps (272098, -3.16 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.1 spell_power points (5.14 DPS) | yes | Satyrmane Sash (17755, -0.30 DPS) [dungeon]; Deathmage Sash (10771, -0.64 DPS) [dungeon]; Runecloth Belt (13856, -1.07 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 28.1 spell_power points (5.52 DPS) | yes | Red Mageweave Pants (10009, -0.26 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.18 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.24 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 24.6 spell_power points (4.83 DPS) | yes | Gilded Sandals (254107, -0.79 DPS) [crafted]; Southsea Mojo Boots (20641, -0.96 DPS) [quest]; Earthen Silk Slippers (254013, -3.58 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.3 spell_power points (3.01 DPS) | yes | Advisor's Ring (19519, -0.65 DPS) [rep]; Snake Hoop (6750, -0.79 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.56 DPS) | yes | Advisor's Ring (19519, -0.20 DPS) [rep]; Snake Hoop (6750, -0.33 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (229.1 DPS) | yes | Rune of the Guard Captain (19120, -0.28 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (229.1 DPS) | yes | Rune of the Guard Captain (19120, -0.21 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (229.1 DPS) | yes | Spellshifter Rod (9527, -0.46 DPS) [quest]; Kindling Stave (11750, -1.58 DPS) [dungeon]; Hanzo Sword (8190, -40.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 267.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.52 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.18 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 1531005113100001531-03202300000000000-053200000000000000)

Set DPS (verified): 431.6. Weights run: 1.3s. Verify run: 3.5s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.093 ± 0.070, crit=0.399 ± 0.011 per rating point (14 rating = 1%, 5.591 per %), hit=0.932 ± 0.107 per rating point (10 rating = 1%, 9.322 per %), spell_haste=-30.393 ± 1.876, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.527 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 53.9 spell_power points (12.76 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.81 DPS) [pvp]; Magister's Crown (16686, -20.08 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (431.6 DPS) | yes | Beads of Ogre Mojo (22149, -0.73 DPS) [quest]; Pebble of Kajaro (19600, -1.42 DPS) [quest]; Jewel of Kajaro (19601, -1.42 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 49.0 spell_power points (11.59 DPS) | yes | Warlord's Silk Amice (231594, -1.80 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.88 DPS) [crafted]; Darkspear Shoulderpads (272103, -5.77 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 34.1 spell_power points (8.06 DPS) | yes | Hide of the Wild (18510, -2.16 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.89 DPS) [quest]; Spritecaster Cape (11623, -3.20 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.7 spell_power points (13.89 DPS) | yes | Warlord's Silk Raiment (231596, -0.36 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.20 DPS) [pvp]; Sorcerer's Robes (226932, -7.71 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 30.7 spell_power points (7.28 DPS) | yes | Runecloth Cuffs (254123, -2.09 DPS) [crafted]; General's Silk Cuffs (16538, -2.36 DPS) [pvp]; Magiskull Cuffs (13107, -2.88 DPS) [world_drop] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (431.6 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Raider Handwraps (272097, -14.03 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.4 spell_power points (13.81 DPS) | yes | Belt of the Archmage (18405, -3.62 DPS) [crafted]; Magister's Belt (16685, -6.72 DPS) [dungeon]; Magician's Cord (272393, -9.20 DPS, sim-verified) [vendor] |
| legs | Sorcerer's Leggings (226933) | Anthion's Parting Words [quest] | sim-verified (431.6 DPS) | yes | Outrider's Silk Leggings (22747, +0.00 DPS) [rep]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (237815, -10.98 DPS, sim-verified) [vendor] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (431.6 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.27 DPS) [dungeon]; Sorcerer's Boots (22064, -17.45 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 47.2 spell_power points (11.16 DPS) | yes | Maiden's Circle (13001, -5.09 DPS) [world_drop]; Channeler's Ring (272406, -5.33 DPS) [vendor]; Deep Rooted Ring (19109, -5.73 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 34.0 spell_power points (8.05 DPS) | yes | Maiden's Circle (13001, -1.98 DPS) [world_drop]; Channeler's Ring (272406, -2.22 DPS) [vendor]; Deep Rooted Ring (19109, -2.63 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (431.6 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (431.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.33 DPS) [crafted]; Rune of the Guard Captain (19120, -1.77 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (431.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.61 DPS) [quest]; Teebu's Blazing Longsword (1728, -53.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 330.3 spell_power points (78.18 DPS) | yes | Wand of Biting Cold (19108, -6.63 DPS, sim-verified) [quest]; Stormrager (16997, -14.20 DPS) [quest]; Brilliant Wand (249385, -14.66 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sorcerer's Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 1531005113100101531-03000000000000000-053500030000000000)

Set DPS (verified): 651.9. Weights run: 1.5s. Verify run: 2.4s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.453 ± 0.013, crit=0.465 ± 0.012 per rating point (14 rating = 1%, 6.512 per %), hit=0.996 ± 0.038 per rating point (10 rating = 1%, 9.959 per %), spell_haste=-13.693 ± 0.827, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.475 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 38.8 spell_power points (17.45 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.43 DPS) [pvp]; Red Mageweave Headband (10033, -7.41 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (651.9 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.50 DPS) [quest]; Jewel of Kajaro (19601, -3.69 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 40.3 spell_power points (18.11 DPS) | yes | Warlord's Silk Amice (231594, -3.83 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.32 DPS, sim-verified) [crafted]; Champion's Silk Mantle (227104, -6.21 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.6 spell_power points (13.29 DPS) | yes | Hide of the Wild (18510, -4.97 DPS) [crafted]; Spritecaster Cape (11623, -5.78 DPS) [dungeon]; Deep Woodlands Cloak (19121, -6.07 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 51.9 spell_power points (23.35 DPS) | yes | Warlord's Silk Raiment (231596, -2.13 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -7.52 DPS) [pvp]; Acumen Robes (17775, -14.12 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.6 spell_power points (11.52 DPS) | yes | Runecloth Cuffs (254123, -4.54 DPS) [crafted]; Sorcerer's Bindings (226929, -6.33 DPS) [quest]; Bracers of Hope (22667, -6.58 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.3 spell_power points (13.15 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 44.9 spell_power points (20.19 DPS) | yes | Belt of the Archmage (18405, -5.04 DPS, sim-verified) [crafted]; Magician's Cord (272393, -6.76 DPS) [vendor]; Defiler's Cloth Girdle (20163, -9.75 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.6 spell_power points (22.27 DPS) | yes | General's Silk Trousers (231595, -1.79 DPS) [pvp]; Sorcerer's Leggings (226933, -6.00 DPS) [quest]; Outrider's Silk Leggings (22747, -10.52 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 28.2 spell_power points (12.69 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.22 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 42.0 spell_power points (18.89 DPS) | yes | Maiden's Circle (13001, -9.38 DPS) [world_drop]; Blessed Band of Light (272407, -9.92 DPS) [vendor]; Deep Rooted Ring (19109, -11.46 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 27.0 spell_power points (12.13 DPS) | yes | Maiden's Circle (13001, -2.61 DPS) [world_drop]; Blessed Band of Light (272407, -3.16 DPS) [vendor]; Deep Rooted Ring (19109, -4.70 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (651.9 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (651.9 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.26 DPS) [crafted]; Rune of the Guard Captain (19120, -3.16 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (651.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.46 DPS) [quest]; Teebu's Blazing Longsword (1728, -42.25 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 176.8 spell_power points (79.46 DPS) | yes | Brilliant Wand (249385, -14.41 DPS) [crafted]; Wand of Biting Cold (19108, -15.46 DPS) [quest]; Stormrager (16997, -15.75 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

