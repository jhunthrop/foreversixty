# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 0000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 32.7. Weights run: 1.0s. Verify run: 0.7s. 145 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.377 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.006 per %), hit=0.223 ± 0.007 per rating point (10 rating = 1%, 2.229 per %), spell_haste=not significant (0.176 ± 0.108), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.65 DPS) | yes | Shadow Goggles (4373, -1.18 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (0.91 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.67 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.44 DPS) | yes | Pearl-clasped Cloak (5542, -0.09 DPS) [crafted]; Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Caretaker's Cape (20428, -0.11 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.75 DPS) | yes | Green Woolen Robe (6243, -0.30 DPS) [crafted]; Green Woolen Vest (2582, -0.31 DPS) [crafted]; Gray Woolen Robe (2585, -0.95 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.9 spell_power points (0.21 DPS) | yes | Windsong Bangles (263336, -0.10 DPS) [quest]; Repurposed Hair Band (281256, -0.12 DPS) [quest]; Bright Bracers (3647, -0.63 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.76 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.23 DPS, sim-verified) [world]; Heavy Woolen Gloves (4310, -0.46 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.60 DPS) | yes | Novice Ardent's Sash (253887, -0.26 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.79 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.0 spell_power points (1.31 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.55 DPS) [dungeon]; Rumpled Kilt (274741, -0.76 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.93 DPS) | yes | Pristine Boots (253889, -0.38 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.49 DPS) [crafted]; Woolen Boots (2583, -0.63 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.63 DPS) | yes | Lavishly Jeweled Ring (1156, -0.38 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop]; The 1 Ring (8350, -0.59 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.54 DPS) | yes | Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.49 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.50 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.8 spell_power points (0.41 DPS) | yes | Witching Stave (1484, -0.21 DPS) [dungeon]; Staff of Westfall (2042, -0.21 DPS) [quest]; Lesser Staff of the Spire (1300, -0.27 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 207.8 spell_power points (22.64 DPS) | yes | Skycaller (12984, -1.07 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.11 DPS) [quest]; Sizzle Stick (8071, -4.44 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 145, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 0000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 63.6. Weights run: 1.0s. Verify run: 0.7s. 264 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.386 ± 0.009, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.589 per %), hit=0.230 ± 0.010 per rating point (10 rating = 1%, 2.300 per %), spell_haste=not significant (0.509 ± 0.155), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.68 DPS) | yes | Silk Headband (7050, -0.31 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.46 DPS) [crafted]; Enchanter's Cowl (4322, -1.87 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 spell_power points (1.42 DPS) | yes | Darkspear Warding Pendant (272075, -1.13 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.19 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.35 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (1.91 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Invoker's Mantle (215365, -0.54 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.76 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Prelacy Cape (7004, -0.15 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.0 spell_power points (2.14 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Tree Bark Jacket (1486, -0.16 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.21 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.37 DPS) | yes | Nightsky Wristbands (6407, -1.02 DPS) [world_drop]; Stonecloth Bindings (14416, -1.08 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.59 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.07 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Truefaith Gloves (7049, -0.13 DPS) [crafted]; Gnoll Casting Gloves (892, -0.15 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.2 spell_power points (1.86 DPS) | yes | Belt of Arugal (6392, -0.31 DPS) [dungeon]; Invoker's Cord (215366, -0.49 DPS) [crafted]; Crimson Silk Belt (7055, -0.53 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.1 spell_power points (1.85 DPS) | yes | Gaze Dreamer Pants (6903, -0.01 DPS) [dungeon]; Pristine Leggings (253987, -0.36 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.58 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (1.48 DPS) | yes | Acidic Walkers (9454, -0.25 DPS) [dungeon]; Glinteye Slippers (273024, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -2.07 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.9 spell_power points (1.21 DPS) | yes | Ring of Power Regulation (273030, -0.29 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.29 DPS) [vendor]; Minor Channeling Ring (1449, -0.33 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.07 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.15 DPS) [vendor]; Minor Channeling Ring (1449, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (5.50 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -2.44 DPS) [quest]; Glimmering Staff (249392, -4.85 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 216.8 spell_power points (33.10 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.33 DPS) [crafted]; Gravestone Scepter (7001, -4.10 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 0000000000000000000-00000000000000000-253510130100300250)

Set DPS (verified): 90.8. Weights run: 1.1s. Verify run: 0.9s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.601 ± 0.019, crit=0.167 ± 0.007 per rating point (14 rating = 1%, 2.341 per %), hit=0.404 ± 0.022 per rating point (10 rating = 1%, 4.044 per %), spell_haste=1.980 ± 0.315, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.18 DPS) [world]; Whitemane's Chapeau (7720, -0.78 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.96 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (1.78 DPS) | yes | Triune Amulet (7722, -0.57 DPS) [dungeon]; Scorn's Icy Choker (23169, -0.96 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.15 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (2.23 DPS) | yes | Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -0.90 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.0 spell_power points (1.36 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.30 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.36 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.6 spell_power points (3.86 DPS) | yes | Dreamweave Vest (10021, -0.33 DPS) [crafted]; Robe of Power (7054, -0.66 DPS) [crafted]; Elemental Raiment (9434, -0.69 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Condor Bracers (15864, -0.30 DPS) [quest]; Windchaser Cuffs (14429, -0.54 DPS) [world_drop]; Arcane Runed Bracers (4744, -0.93 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (3.08 DPS) | yes | Red Mageweave Gloves (10018, -0.51 DPS) [crafted]; Black Mageweave Gloves (10003, -0.81 DPS) [crafted]; Gilded Handwraps (254021, -1.24 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.45 DPS) [crafted]; Gilded Cord (254037, -0.48 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.00 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (3.20 DPS) | yes | Crimson Silk Pantaloons (7062, -1.01 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.12 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.39 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.62 DPS) | yes | Gilded Slippers (254001, -1.93 DPS) [crafted]; Acidic Walkers (9454, -2.14 DPS) [dungeon]; Spidersilk Boots (4320, -2.20 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (1.96 DPS) | yes | Lorekeeper's Ring (19524, -0.60 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.75 DPS) [quest]; Reedknot Ring (9622, -0.91 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.36 DPS) | yes | Lorekeeper's Ring (19524, +0.00 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.15 DPS) [quest]; Reedknot Ring (9622, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.18 DPS) [dungeon]; Spine of the Devourer (273040, -1.93 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (2.23 DPS) | yes | Orb of the Forgotten Seer (7685, -0.12 DPS) [dungeon]; Orb of Lorica (11262, -0.42 DPS) [quest]; Orb of Mystic Insight (249394, -0.63 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 264.8 spell_power points (39.92 DPS) | yes | Nether Force Wand (11263, -2.32 DPS) [quest]; Icefury Wand (7514, -2.46 DPS) [quest]; Ragefire Wand (7513, -2.51 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 2030005000000000000-00000000000000000-253510130100300250)

Set DPS (verified): 225.0. Weights run: 1.0s. Verify run: 1.1s. 436 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.818 ± 0.034, crit=0.268 ± 0.010 per rating point (14 rating = 1%, 3.755 per %), hit=0.751 ± 0.045 per rating point (10 rating = 1%, 7.513 per %), spell_haste=7.477 ± 0.550, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.4 spell_power points (7.73 DPS) | yes | Dreamweave Circlet (10041, -1.35 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.70 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.83 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.5 spell_power points (2.96 DPS) | yes | Scorn's Icy Choker (23169, -0.36 DPS) [dungeon]; Horizon Choker (13085, -0.46 DPS) [world_drop]; Mindburst Medallion (11196, -0.58 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.7 spell_power points (6.06 DPS) | yes | Red Mageweave Shoulders (10029, -1.85 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.88 DPS) [vendor]; Inquisitor's Shawl (19507, -2.21 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.9 spell_power points (4.13 DPS) | yes | Runecloth Cloak (13860, -0.74 DPS) [crafted]; Big Voodoo Cloak (8216, -1.43 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.63 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.4 spell_power points (7.73 DPS) | yes | Robe of the Magi (1716, -1.85 DPS) [world_drop]; Runecloth Tunic (13857, -2.05 DPS) [crafted]; Dreamweave Vest (10021, -2.19 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.7 spell_power points (2.78 DPS) | yes | Forgotten Wraps (9433, -0.64 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.64 DPS) [quest]; Bloodband Bracers (11469, -3.31 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.0 spell_power points (6.77 DPS) | yes | Raider Handwraps (272098, -1.31 DPS) [vendor]; Dreamweave Gloves (10019, -2.12 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.32 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.2 spell_power points (4.85 DPS) | yes | Dawnspire Cord (12466, -0.14 DPS) [dungeon]; Deathmage Sash (10771, -0.64 DPS) [dungeon]; Runecloth Belt (13856, -0.95 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 25.6 spell_power points (5.59 DPS) | yes | Red Mageweave Pants (10009, -0.38 DPS) [crafted]; Wizardweave Leggings (14132, -1.44 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.52 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.25 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.25 DPS) [vendor]; Gilded Sandals (254107, -1.23 DPS) [crafted]; Southsea Mojo Boots (20641, -1.53 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (3.08 DPS) | yes | Lorekeeper's Ring (19523, -0.46 DPS) [rep]; Snake Hoop (6750, -0.87 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.84 DPS) | yes | Lorekeeper's Ring (19523, -0.22 DPS) [rep]; Snake Hoop (6750, -0.64 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (225.0 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (225.0 DPS) | yes | Mark of the Chosen (17774, -3.82 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (225.0 DPS) | yes | Spellshifter Rod (9527, -0.20 DPS) [quest]; Kindling Stave (11750, -0.85 DPS) [dungeon]; Hanzo Sword (8190, -38.24 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 240.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.41 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.00 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Knight's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 2030005000000000000-11302300000000000-253510130100300250)

Set DPS (verified): 342.9. Weights run: 1.1s. Verify run: 2.2s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.668 ± 0.036, crit=0.367 ± 0.012 per rating point (14 rating = 1%, 5.140 per %), hit=1.015 ± 0.053 per rating point (10 rating = 1%, 10.150 per %), spell_haste=7.435 ± 0.592, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.8 spell_power points (9.82 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.07 DPS) [pvp]; Red Mageweave Headband (10033, -5.17 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (342.9 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.39 DPS) [quest]; Chains of the Lich (23125, -0.39 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 42.2 spell_power points (9.66 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.64 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.78 DPS) [vendor]; Mantle of the Timbermaw (19050, -3.64 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.5 spell_power points (7.22 DPS) | yes | Hide of the Wild (18510, -2.48 DPS) [crafted]; Spritecaster Cape (11623, -3.09 DPS) [dungeon]; Runecloth Cloak (13860, -3.93 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 53.2 spell_power points (12.19 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.84 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.59 DPS) [pvp]; Sorcerer's Robes (226932, -7.13 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 27.3 spell_power points (6.27 DPS) | yes | Runecloth Cuffs (254123, -2.22 DPS) [crafted]; Sorcerer's Bindings (226929, -3.13 DPS) [quest]; Bracers of Hope (22667, -3.21 DPS) [quest] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 31.5 spell_power points (7.22 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 49.8 spell_power points (11.43 DPS) | yes | Magician's Cord (272393, -3.40 DPS) [vendor]; Stormpike Cloth Girdle (19094, -5.77 DPS) [rep]; Belt of the Archmage (18405, -6.70 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.0 spell_power points (11.22 DPS) | yes | Marshal's Silk Leggings (231605, -0.11 DPS) [pvp]; Sorcerer's Leggings (226933, -2.09 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -2.63 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 31.7 spell_power points (7.26 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.84 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 44.2 spell_power points (10.12 DPS) | yes | Maiden's Circle (13001, -4.93 DPS) [world_drop]; Blessed Band of Light (272407, -5.50 DPS) [vendor]; Channeler's Ring (272406, -5.65 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 29.3 spell_power points (6.73 DPS) | yes | Maiden's Circle (13001, -1.53 DPS) [world_drop]; Blessed Band of Light (272407, -2.11 DPS) [vendor]; Channeler's Ring (272406, -2.26 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (342.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (342.9 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.12 DPS) [crafted]; Uther's Strength (11302, -1.83 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (342.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.00 DPS) [quest]; Teebu's Blazing Longsword (1728, -42.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 340.8 spell_power points (78.14 DPS) | yes | Wand of Biting Cold (19108, -3.38 DPS, sim-verified) [quest]; Stormrager (16997, -14.68 DPS) [quest]; Brilliant Wand (249385, -15.42 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 2030005000000000000-13102000000000000-253510130101301251)

Set DPS (verified): 524.2. Weights run: 1.2s. Verify run: 2.3s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.158 ± 0.012, crit=0.418 ± 0.013 per rating point (14 rating = 1%, 5.857 per %), hit=0.986 ± 0.039 per rating point (10 rating = 1%, 9.857 per %), spell_haste=6.933 ± 0.244, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 30.8 spell_power points (11.72 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.42 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.43 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (524.2 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.88 DPS) [quest]; Jewel of Kajaro (19601, -4.72 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.2 spell_power points (13.40 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.99 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.93 DPS) [crafted]; Argent Shoulders (19059, -5.71 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.1 spell_power points (10.32 DPS) | yes | Hide of the Wild (18510, -4.39 DPS) [crafted]; Spritecaster Cape (11623, -4.63 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -4.99 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 47.8 spell_power points (18.17 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.36 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.93 DPS) [pvp]; Robe of the Magi (1716, -11.87 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.3 spell_power points (8.85 DPS) | yes | Runecloth Cuffs (254123, -4.07 DPS) [crafted]; Arcane Runed Bracers (4744, -5.43 DPS) [quest]; Spidertank Oilrag (9448, -5.43 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.8 spell_power points (10.57 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.3 spell_power points (14.59 DPS) | yes | Belt of the Archmage (18405, -4.67 DPS, sim-verified) [crafted]; Magician's Cord (272393, -5.91 DPS) [vendor]; Highlander's Cloth Girdle (20047, -6.67 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.3 spell_power points (17.24 DPS) | yes | Marshal's Silk Leggings (231605, -2.39 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -5.99 DPS) [pvp]; Sorcerer's Leggings (226933, -7.60 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.13 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.18 DPS) [quest]; Sorcerer's Sandals (226931, -0.18 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 39.3 spell_power points (14.95 DPS) | yes | Blessed Band of Light (272407, -7.39 DPS) [vendor]; Maiden's Circle (13001, -7.68 DPS) [world_drop]; Don Julio's Band (19325, -8.97 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 23.7 spell_power points (9.03 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Maiden's Circle (13001, -1.76 DPS) [world_drop]; Don Julio's Band (19325, -3.05 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (524.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (524.2 DPS) | yes | Frozen Heart of the Mountain (249469, -1.95 DPS) [crafted]; Weakness Analyzer (272438, -2.52 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.04 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (524.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.04 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -34.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 207.7 spell_power points (79.04 DPS) | yes | Wand of Biting Cold (19108, -4.86 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.62 DPS) [crafted]; Stormrager (16997, -16.05 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 0000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 30.2. Weights run: 1.0s. Verify run: 0.7s. 134 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.377 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.006 per %), hit=0.223 ± 0.007 per rating point (10 rating = 1%, 2.229 per %), spell_haste=not significant (0.176 ± 0.108), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.65 DPS) | yes | Shadow Goggles (4373, -1.12 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (0.91 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.61 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.44 DPS) | yes | Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Battle Healer's Cloak (20427, -0.11 DPS) [rep]; Pearl-clasped Cloak (5542, -0.23 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.75 DPS) | yes | Green Woolen Robe (6243, -0.30 DPS) [crafted]; Green Woolen Vest (2582, -0.31 DPS) [crafted]; Gray Woolen Robe (2585, -0.97 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.33 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Featherbead Bracers (15452, -0.12 DPS) [quest]; Owlbeard Bracers (16981, -0.14 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.76 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.21 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.33 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.60 DPS) | yes | Novice Ardent's Sash (253887, -0.26 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.73 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.0 spell_power points (1.31 DPS) | yes | Filigreed Pristine Leggings (253937, -0.41 DPS) [crafted]; Silk-threaded Trousers (1929, -0.55 DPS) [dungeon]; Rumpled Kilt (274741, -0.76 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.93 DPS) | yes | Pristine Boots (253889, -0.37 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.49 DPS) [crafted]; Woolen Boots (2583, -0.63 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.54 DPS) | yes | Loop of Sacrifice (281673, -0.34 DPS) [quest]; Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; The 1 Ring (8350, -0.50 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.3 spell_power points (0.25 DPS) | yes | Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; The 1 Ring (8350, -0.21 DPS) [world]; Loop of Sacrifice (281673, -0.26 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.8 spell_power points (0.41 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.16 DPS) [world]; Clear Crystal Rod (16894, -0.21 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 207.8 spell_power points (22.64 DPS) | yes | Skycaller (12984, -1.14 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.44 DPS) [quest]; Greater Magic Wand (11288, -4.92 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 134, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (troll, 0000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 60.6. Weights run: 1.0s. Verify run: 0.7s. 245 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.386 ± 0.009, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.589 per %), hit=0.230 ± 0.010 per rating point (10 rating = 1%, 2.300 per %), spell_haste=not significant (0.509 ± 0.155), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.68 DPS) | yes | Silk Headband (7050, -0.31 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.46 DPS) [crafted]; Enchanter's Cowl (4322, -1.60 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 spell_power points (1.42 DPS) | yes | Darkspear Warding Pendant (272075, -1.13 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.19 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.26 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (1.91 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Invoker's Mantle (215365, -0.54 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.76 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.0 spell_power points (2.14 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Tree Bark Jacket (1486, -0.16 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.21 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.37 DPS) | yes | Tabitha's Cuffs (251486, -0.92 DPS) [quest]; Nightsky Wristbands (6407, -1.02 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.14 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.9 spell_power points (1.21 DPS) | yes | Truefaith Gloves (7049, -0.27 DPS) [crafted]; Gnoll Casting Gloves (892, -0.29 DPS) [world]; Serpent Gloves (5970, -0.39 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.2 spell_power points (1.86 DPS) | yes | Belt of Arugal (6392, -0.31 DPS) [dungeon]; Warsong Sash (16975, -0.47 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.49 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.1 spell_power points (1.85 DPS) | yes | Gaze Dreamer Pants (6903, -0.01 DPS) [dungeon]; Pristine Leggings (253987, -0.36 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.58 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (1.48 DPS) | yes | Acidic Walkers (9454, -0.25 DPS) [dungeon]; Glinteye Slippers (273024, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.93 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.9 spell_power points (1.21 DPS) | yes | Ring of Power Regulation (273030, -0.29 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.29 DPS) [vendor]; Solid Gold Band (279845, -0.60 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.07 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.15 DPS) [vendor]; Solid Gold Band (279845, -0.46 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (5.50 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -4.85 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -4.91 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 216.8 spell_power points (33.10 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.33 DPS) [crafted]; Unstable Power Core (279847, -3.39 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 245, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (troll, 0000000000000000000-00000000000000000-253510130100300250)

Set DPS (verified): 86.4. Weights run: 1.1s. Verify run: 0.9s. 320 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.601 ± 0.019, crit=0.167 ± 0.007 per rating point (14 rating = 1%, 2.341 per %), hit=0.404 ± 0.022 per rating point (10 rating = 1%, 4.044 per %), spell_haste=1.980 ± 0.315, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (86.4 DPS) | yes | Augural Shroud (2620, -0.18 DPS) [world]; Whitemane's Chapeau (7720, -0.78 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.87 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (1.78 DPS) | yes | Triune Amulet (7722, -0.57 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.15 DPS) [vendor]; Scorn's Icy Choker (23169, -1.18 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (2.23 DPS) | yes | Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -1.12 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.0 spell_power points (1.36 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.30 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.36 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.6 spell_power points (3.86 DPS) | yes | Dreamweave Vest (10021, -0.33 DPS) [crafted]; Robe of Power (7054, -0.66 DPS) [crafted]; Elemental Raiment (9434, -0.69 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.36 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.03 DPS) [quest]; Condor Bracers (15864, -0.30 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (3.08 DPS) | yes | Red Mageweave Gloves (10018, -0.51 DPS) [crafted]; Black Mageweave Gloves (10003, -0.81 DPS) [crafted]; Gilded Handwraps (254021, -1.24 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.4 spell_power points (2.47 DPS) | yes | Deathmage Sash (10771, +0.00 DPS, sim-verified) [dungeon]; Star Belt (4329, -0.51 DPS) [crafted]; Gilded Cord (254037, -0.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (3.20 DPS) | yes | Crimson Silk Pantaloons (7062, -0.81 DPS) [crafted]; Abomination Skin Leggings (23173, -1.12 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.39 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.62 DPS) | yes | Gilded Slippers (254001, -0.66 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.14 DPS) [dungeon]; Spidersilk Boots (4320, -2.20 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (1.96 DPS) | yes | Advisor's Ring (19520, -0.60 DPS) [rep]; Reedknot Ring (9622, -0.91 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.36 DPS) | yes | Advisor's Ring (19520, -0.00 DPS) [rep]; Reedknot Ring (9622, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.18 DPS) [dungeon]; Spine of the Devourer (273040, -1.93 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (2.23 DPS) | yes | Orb of the Forgotten Seer (7685, -0.12 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.63 DPS) [crafted]; Omega Orb (7749, -0.87 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 264.8 spell_power points (39.92 DPS) | yes | Nether Force Wand (11263, -1.03 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.46 DPS) [quest]; Ragefire Wand (7513, -2.51 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 320, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (troll, 2030005000000000000-00000000000000000-253510130100300250)

Set DPS (verified): 218.9. Weights run: 1.0s. Verify run: 1.0s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.818 ± 0.034, crit=0.268 ± 0.010 per rating point (14 rating = 1%, 3.755 per %), hit=0.751 ± 0.045 per rating point (10 rating = 1%, 7.513 per %), spell_haste=7.477 ± 0.550, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.4 spell_power points (7.73 DPS) | yes | Dreamweave Circlet (10041, -1.35 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.70 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.83 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.5 spell_power points (2.96 DPS) | yes | Horizon Choker (13085, -0.46 DPS) [world_drop]; Mindburst Medallion (11196, -0.58 DPS) [quest]; Scorn's Icy Choker (23169, -2.72 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.7 spell_power points (6.06 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.88 DPS) [vendor]; Inquisitor's Shawl (19507, -2.21 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.71 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.4 spell_power points (4.23 DPS) | yes | Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.83 DPS) [crafted]; Big Voodoo Cloak (8216, -1.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.4 spell_power points (7.73 DPS) | yes | Robe of the Magi (1716, -1.85 DPS) [world_drop]; Runecloth Tunic (13857, -2.05 DPS) [crafted]; Dreamweave Vest (10021, -2.19 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.7 spell_power points (2.78 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -2.58 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.0 spell_power points (6.77 DPS) | yes | Raider Handwraps (272098, -1.31 DPS) [vendor]; Dreamweave Gloves (10019, -2.12 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.32 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.2 spell_power points (4.85 DPS) | yes | Dawnspire Cord (12466, -0.14 DPS) [dungeon]; Deathmage Sash (10771, -0.64 DPS) [dungeon]; Runecloth Belt (13856, -0.95 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 25.6 spell_power points (5.59 DPS) | yes | Red Mageweave Pants (10009, -0.38 DPS) [crafted]; Wizardweave Leggings (14132, -1.44 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.52 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.25 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.25 DPS) [vendor]; Gilded Sandals (254107, -1.23 DPS) [crafted]; Southsea Mojo Boots (20641, -1.53 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (3.08 DPS) | yes | Advisor's Ring (19519, -0.46 DPS) [rep]; Snake Hoop (6750, -0.87 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.84 DPS) | yes | Advisor's Ring (19519, -0.22 DPS) [rep]; Snake Hoop (6750, -0.64 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (218.9 DPS) | yes | Rune of the Guard Captain (19120, -0.33 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (218.9 DPS) | yes | Rune of the Guard Captain (19120, -0.16 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (218.9 DPS) | yes | Spellshifter Rod (9527, -0.20 DPS) [quest]; Kindling Stave (11750, -0.85 DPS) [dungeon]; Hanzo Sword (8190, -36.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 240.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.41 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.00 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Stone Guard's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (troll, 2030005000000000000-11302300000000000-253510130100300250)

Set DPS (verified): 339.3. Weights run: 1.1s. Verify run: 3.0s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.668 ± 0.036, crit=0.367 ± 0.012 per rating point (14 rating = 1%, 5.140 per %), hit=1.015 ± 0.053 per rating point (10 rating = 1%, 10.150 per %), spell_haste=7.435 ± 0.592, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.8 spell_power points (9.82 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.07 DPS) [pvp]; Red Mageweave Headband (10033, -9.62 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (339.3 DPS) | yes | Orb of the Darkmoon (19426, -0.39 DPS) [quest]; Chains of the Lich (23125, -0.39 DPS) [dungeon]; Jewel of Kajaro (19601, -3.30 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 42.2 spell_power points (9.66 DPS) | yes | Warlord's Silk Amice (231594, -1.64 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.78 DPS) [vendor]; Mantle of the Timbermaw (19050, -5.25 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.5 spell_power points (7.22 DPS) | yes | Hide of the Wild (18510, -2.48 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.09 DPS) [quest]; Spritecaster Cape (11623, -3.09 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 53.2 spell_power points (12.19 DPS) | yes | Warlord's Silk Raiment (231596, -0.84 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.59 DPS) [pvp]; Sorcerer's Robes (226932, -7.21 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.3 spell_power points (6.27 DPS) | yes | Runecloth Cuffs (254123, -2.22 DPS) [crafted]; Sorcerer's Bindings (226929, -3.13 DPS) [quest]; Bracers of Hope (22667, -3.21 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (339.3 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, -5.07 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 49.8 spell_power points (11.43 DPS) | yes | Magician's Cord (272393, -3.40 DPS) [vendor]; Frostwolf Cloth Belt (19090, -5.77 DPS) [rep]; Belt of the Archmage (18405, -9.24 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.0 spell_power points (11.22 DPS) | yes | General's Silk Trousers (231595, -0.11 DPS) [pvp]; Sorcerer's Leggings (226933, -2.09 DPS) [quest]; Outrider's Silk Leggings (22747, -5.31 DPS, sim-verified) [rep] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (339.3 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.84 DPS) [dungeon]; Sorcerer's Boots (22064, -5.17 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 44.2 spell_power points (10.12 DPS) | yes | Maiden's Circle (13001, -4.93 DPS) [world_drop]; Blessed Band of Light (272407, -5.50 DPS) [vendor]; Channeler's Ring (272406, -5.65 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 29.3 spell_power points (6.73 DPS) | yes | Maiden's Circle (13001, -1.53 DPS) [world_drop]; Blessed Band of Light (272407, -2.11 DPS) [vendor]; Channeler's Ring (272406, -2.26 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (339.3 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (339.3 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.12 DPS) [crafted]; Rune of the Guard Captain (19120, -1.58 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (339.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.00 DPS) [quest]; Teebu's Blazing Longsword (1728, -44.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 340.8 spell_power points (78.14 DPS) | yes | Wand of Biting Cold (19108, -7.83 DPS, sim-verified) [quest]; Stormrager (16997, -14.68 DPS) [quest]; Brilliant Wand (249385, -15.42 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (troll, 2030005000000000000-13102000000000000-253510130101301251)

Set DPS (verified): 522.6. Weights run: 1.2s. Verify run: 2.2s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.158 ± 0.012, crit=0.418 ± 0.013 per rating point (14 rating = 1%, 5.857 per %), hit=0.986 ± 0.039 per rating point (10 rating = 1%, 9.857 per %), spell_haste=6.933 ± 0.244, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 30.8 spell_power points (11.72 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.42 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.81 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (522.6 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.88 DPS) [quest]; Jewel of Kajaro (19601, -4.74 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.2 spell_power points (13.40 DPS) | yes | Warlord's Silk Amice (231594, -2.99 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.93 DPS) [crafted]; Argent Shoulders (19059, -6.52 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.1 spell_power points (10.32 DPS) | yes | Hide of the Wild (18510, -4.39 DPS) [crafted]; Spritecaster Cape (11623, -4.63 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -4.99 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 47.8 spell_power points (18.17 DPS) | yes | Warlord's Silk Raiment (231596, -2.36 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.93 DPS) [pvp]; Robe of the Magi (1716, -11.90 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.3 spell_power points (8.85 DPS) | yes | Runecloth Cuffs (254123, -4.07 DPS) [crafted]; Spidertank Oilrag (9448, -5.43 DPS) [dungeon]; Sorcerer's Bindings (226929, -5.59 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.8 spell_power points (10.57 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.3 spell_power points (14.59 DPS) | yes | Belt of the Archmage (18405, -4.66 DPS, sim-verified) [crafted]; Magician's Cord (272393, -5.91 DPS) [vendor]; Defiler's Cloth Girdle (20163, -6.67 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.3 spell_power points (17.24 DPS) | yes | General's Silk Trousers (231595, -2.39 DPS) [pvp]; Outrider's Silk Leggings (22747, -5.44 DPS) [rep]; Sorcerer's Leggings (226933, -8.12 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.13 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.18 DPS) [quest]; Sorcerer's Sandals (226931, -0.18 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 39.3 spell_power points (14.95 DPS) | yes | Blessed Band of Light (272407, -7.39 DPS) [vendor]; Maiden's Circle (13001, -7.68 DPS) [world_drop]; Don Julio's Band (19325, -8.97 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 23.7 spell_power points (9.03 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Maiden's Circle (13001, -1.76 DPS) [world_drop]; Don Julio's Band (19325, -3.05 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (522.6 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (522.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.95 DPS) [crafted]; Rune of the Guard Captain (19120, -2.70 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (522.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.04 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -34.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 207.7 spell_power points (79.04 DPS) | yes | Wand of Biting Cold (19108, -5.27 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.62 DPS) [crafted]; Stormrager (16997, -16.05 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

