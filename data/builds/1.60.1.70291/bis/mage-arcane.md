# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 153002000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 29.9. Weights run: 1.5s. Verify run: 1.3s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.306 ± 0.011, crit=0.079 ± 0.003 per rating point (14 rating = 1%, 1.108 per %), hit=0.300 ± 0.014 per rating point (10 rating = 1%, 3.004 per %), spell_haste=not significant (0.455 ± 0.116), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.465 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -1.09 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.8 spell_power points (0.73 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.36 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.76 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Caretaker's Cape (20428, -0.09 DPS) [rep]; Pearl-clasped Cloak (5542, -0.10 DPS) [crafted]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.62 DPS) | yes | Green Woolen Vest (2582, -0.24 DPS) [crafted]; Bloody Apron (6226, -0.24 DPS) [dungeon]; Gray Woolen Robe (2585, -0.80 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.5 spell_power points (0.14 DPS) | yes | Windsong Bangles (263336, -0.05 DPS) [quest]; Repurposed Hair Band (281256, -0.09 DPS) [quest]; Bright Bracers (3647, -0.80 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.66 DPS) | yes | Gnoll Casting Gloves (892, -0.18 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.20 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.42 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.49 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.57 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.4 spell_power points (1.08 DPS) | yes | Filigreed Pristine Leggings (253937, -0.34 DPS) [crafted]; Silk-threaded Trousers (1929, -0.42 DPS) [dungeon]; Rumpled Kilt (274741, -0.61 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.78 DPS) | yes | Pristine Boots (253889, -0.41 DPS) [crafted]; Woolen Boots (2583, -0.53 DPS) [crafted]; Red Woolen Boots (4313, -0.68 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.53 DPS) | yes | Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; The 1 Ring (8350, -0.50 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.47 DPS) | yes | Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; The 1 Ring (8350, -0.44 DPS) [world]; Lavishly Jeweled Ring (1156, -0.68 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.1 spell_power points (0.29 DPS) | yes | Witching Stave (1484, -0.14 DPS) [dungeon]; Staff of Westfall (2042, -0.14 DPS) [quest]; Lesser Staff of the Spire (1300, -0.53 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 238.7 spell_power points (22.59 DPS) | yes | Skycaller (12984, -0.49 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.06 DPS) [quest]; Sizzle Stick (8071, -4.47 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 153005113100010000-00000000000000000-0000000000000000000)

Set DPS (verified): 117.2. Weights run: 1.4s. Verify run: 1.5s. 260 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.392 ± 0.021, crit=0.067 ± 0.002 per rating point (14 rating = 1%, 0.937 per %), hit=0.251 ± 0.039 per rating point (10 rating = 1%, 2.510 per %), spell_haste=not significant (0.535 ± 0.325), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.902 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.95 DPS) | yes | Silk Headband (7050, -0.54 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.80 DPS) [crafted]; Enchanter's Cowl (4322, -4.98 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (2.50 DPS) | yes | Crystal Starfire Medallion (5003, -2.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.08 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.79 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (3.35 DPS) | yes | Fairywing Mantle (9536, -0.80 DPS) [quest]; Invoker's Mantle (215365, -0.96 DPS) [crafted]; Death Speaker Mantle (6685, -1.15 DPS, sim-verified) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (1.61 DPS) | yes | Hillman's Cloak (3719, -0.27 DPS) [crafted]; Repairman's Cape (9605, -0.38 DPS) [quest]; Heavy Woolen Cloak (4311, -0.54 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.1 spell_power points (3.77 DPS) | yes | Tree Bark Jacket (1486, -0.29 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.37 DPS) [dungeon]; Beguiler Robes (7728, -0.52 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.41 DPS) | yes | Nightsky Wristbands (6407, -1.78 DPS) [world_drop]; Stonecloth Bindings (14416, -1.88 DPS) [world_drop]; Glowing Magical Bracelets (13106, -3.32 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.3 spell_power points (2.23 DPS) | yes | Serpent Gloves (5970, -0.35 DPS) [dungeon]; Shilly Mitts (9609, -0.35 DPS) [quest]; Truefaith Gloves (7049, -0.57 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.2 spell_power points (3.26 DPS) | yes | Belt of Arugal (6392, -0.54 DPS) [dungeon]; Invoker's Cord (215366, -0.86 DPS) [crafted]; Crimson Silk Belt (7055, -0.92 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.1 spell_power points (3.25 DPS) | yes | Gaze Dreamer Pants (6903, -0.04 DPS) [dungeon]; Pristine Leggings (253987, -0.64 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.01 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (2.61 DPS) | yes | Acidic Walkers (9454, -0.43 DPS) [dungeon]; Nimbus Boots (6998, -1.00 DPS) [quest]; Spidersilk Boots (4320, -2.74 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.87 DPS) | yes | Minor Channeling Ring (1449, -0.33 DPS) [quest]; Electrocutioner Lagnut (9447, -1.07 DPS) [dungeon]; Snake Hoop (6750, -1.14 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.61 DPS) | yes | Electrocutioner Lagnut (9447, -0.80 DPS) [dungeon]; Snake Hoop (6750, -0.87 DPS) [quest]; Minor Channeling Ring (1449, -2.55 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | 30.0 spell_power points (8.03 DPS) | yes | Twisted Chanter's Staff (890, -6.98 DPS) [world_drop]; Lesser Staff of the Spire (1300, -7.40 DPS) [world]; Glimmering Staff (249392, -13.60 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Greater Mystic Wand (217287) | Enchanting [crafted] | sim-verified (117.2 DPS) | yes | Gravestone Scepter (7001, -1.34 DPS) [quest]; Scorching Wand (5213, -1.49 DPS) [world_drop]; Starfaller (13063, -1.76 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff; ranged: Greater Mystic Wand

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 153005113100011531-00000000000000000-0000000000000000000)

Set DPS (verified): 181.9. Weights run: 1.4s. Verify run: 1.6s. 334 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.583 ± 0.038, crit=0.166 ± 0.005 per rating point (14 rating = 1%, 2.321 per %), hit=0.367 ± 0.078 per rating point (10 rating = 1%, 3.673 per %), spell_haste=not significant (1.299 ± 0.570), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.919 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.53 DPS) | yes | Augural Shroud (2620, -1.74 DPS, sim-verified) [world]; Living Cowl (5608, -2.49 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -2.59 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (3.63 DPS) | yes | Scorn's Icy Choker (23169, -1.93 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.36 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.36 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.54 DPS) | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.54 DPS) [quest]; Green Silken Shoulders (7057, -1.80 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.9 spell_power points (2.77 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.60 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.78 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (7.93 DPS) | yes | Dreamweave Vest (10021, -0.70 DPS) [crafted]; Robe of Power (7054, -1.40 DPS) [crafted]; Elemental Raiment (9434, -1.40 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.80 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.62 DPS) [quest]; Windchaser Cuffs (14429, -1.17 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (6.33 DPS) | yes | Black Mageweave Gloves (10003, -1.66 DPS) [crafted]; Red Mageweave Gloves (10018, -2.04 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.57 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.3 spell_power points (5.08 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Star Belt (4329, -1.04 DPS) [crafted]; Gilded Cord (254037, -1.14 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.53 DPS) | yes | Abomination Skin Leggings (23173, -2.28 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.28 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.80 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.47 DPS) | yes | Gilded Slippers (254001, -1.97 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.46 DPS) [dungeon]; Spidersilk Boots (4320, -4.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (4.02 DPS) | yes | Ring of Forlorn Spirits (2043, -1.53 DPS) [quest]; Reedknot Ring (9622, -1.84 DPS) [quest]; Minor Channeling Ring (1449, -2.10 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.80 DPS) | yes | Reedknot Ring (9622, -0.62 DPS) [quest]; Minor Channeling Ring (1449, -0.88 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (181.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.36 DPS) [dungeon]; Hardened Root Staff (1317, -5.81 DPS) [quest] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (4.59 DPS) | yes | Orb of the Forgotten Seer (7685, -0.23 DPS) [dungeon]; Thrash's Trash (276204, -0.23 DPS) [vendor]; Orb of Lorica (11262, -0.91 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 130.1 spell_power points (40.47 DPS) | yes | Nether Force Wand (11263, -2.59 DPS) [quest]; Icefury Wand (7514, -2.73 DPS) [quest]; Ragefire Wand (7513, -2.78 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 334, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 153005113100011531-03202300000000000-0000000000000000000)

Set DPS (verified): 294.0. Weights run: 1.4s. Verify run: 2.1s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.688 ± 0.059, crit=0.254 ± 0.009 per rating point (14 rating = 1%, 3.560 per %), hit=0.695 ± 0.123 per rating point (10 rating = 1%, 6.954 per %), spell_haste=not significant (1.201 ± 0.837), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.922 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 32.8 spell_power points (10.39 DPS) | yes | Dreamweave Circlet (10041, -1.55 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.83 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.20 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.5 spell_power points (3.97 DPS) | yes | Scorn's Icy Choker (23169, -0.44 DPS) [dungeon]; Mindburst Medallion (11196, -0.75 DPS) [quest]; Horizon Choker (13085, -0.91 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.4 spell_power points (8.05 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -2.42 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.56 DPS) [crafted]; Black Mageweave Shoulders (10027, -2.92 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.1 spell_power points (5.75 DPS) | yes | Runecloth Cloak (13860, -2.08 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -2.20 DPS) [crafted]; Darkspear Raider's Cloak (272076, -2.70 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 32.8 spell_power points (10.39 DPS) | yes | Robe of the Magi (1716, -2.10 DPS) [world_drop]; Runecloth Tunic (13857, -2.60 DPS) [crafted]; Dreamweave Vest (10021, -2.72 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.8 spell_power points (3.75 DPS) | yes | Arcane Runed Bracers (4744, -0.89 DPS) [quest]; Spidertank Oilrag (9448, -0.89 DPS) [dungeon]; Bloodband Bracers (11469, -3.33 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.6 spell_power points (9.07 DPS) | yes | Raider Handwraps (272098, -2.05 DPS) [vendor]; Dreamweave Gloves (10019, -2.48 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.98 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 20.9 spell_power points (6.62 DPS) | yes | Dawnspire Cord (12466, -0.57 DPS) [dungeon]; Deathmage Sash (10771, -1.13 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.31 DPS) [rep] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 23.8 spell_power points (7.56 DPS) | yes | Red Mageweave Pants (10009, -0.49 DPS) [crafted]; Wizardweave Leggings (14132, -1.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.18 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.62 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.91 DPS) [vendor]; Gilded Sandals (254107, -2.16 DPS) [crafted]; Black Mageweave Boots (10026, -2.60 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (4.26 DPS) | yes | Lorekeeper's Ring (19523, -0.46 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.73 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.13 DPS) | yes | Lorekeeper's Ring (19523, -0.32 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.59 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (294.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (294.0 DPS) | yes | Mark of the Chosen (17774, -2.22 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (294.0 DPS) | yes | Kindling Stave (11750, -0.50 DPS) [dungeon]; Spire of Hakkar (10844, -2.48 DPS) [world]; Hanzo Sword (8190, -33.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 165.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.21 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Knight's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 498.6. Weights run: 1.2s. Verify run: 4.0s. 876 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.841 ± 0.067, crit=0.407 ± 0.013 per rating point (14 rating = 1%, 5.699 per %), hit=0.958 ± 0.163 per rating point (10 rating = 1%, 9.579 per %), spell_haste=not significant (2.418 ± 1.134), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.914 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.7 spell_power points (16.05 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.98 DPS) [pvp]; Magister's Crown (16686, -7.91 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (498.6 DPS) | yes | Beads of Ogre Mojo (22149, -0.96 DPS) [quest]; Chains of the Lich (23125, -1.32 DPS) [dungeon]; Jewel of Kajaro (19601, -2.57 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.3 spell_power points (15.24 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.59 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.93 DPS) [crafted]; Darkspear Shoulderpads (272103, -9.40 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.3 spell_power points (10.87 DPS) | yes | Hide of the Wild (18510, -3.33 DPS) [crafted]; Spritecaster Cape (11623, -4.46 DPS) [dungeon]; Runecloth Cloak (13860, -5.58 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.8 spell_power points (18.77 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.94 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.98 DPS) [pvp]; Sorcerer's Robes (226932, -13.85 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.7 spell_power points (9.66 DPS) | yes | Runecloth Cuffs (254123, -3.13 DPS) [crafted]; Marshal's Silk Bracers (16438, -4.29 DPS) [pvp]; Sorcerer's Bindings (226929, -4.48 DPS) [quest] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 33.4 spell_power points (11.22 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.1 spell_power points (17.86 DPS) | yes | Belt of the Archmage (18405, -4.68 DPS) [crafted]; Stormpike Cloth Girdle (19094, -8.97 DPS) [rep]; Magician's Cord (272393, -10.38 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.8 spell_power points (17.43 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -3.29 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -3.64 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.5 spell_power points (11.59 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.47 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.1 spell_power points (15.19 DPS) | yes | Channeler's Ring (272406, -8.00 DPS) [vendor]; Deep Rooted Ring (19109, -8.32 DPS) [world]; Blessed Band of Light (272407, -8.60 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 31.3 spell_power points (10.51 DPS) | yes | Deep Rooted Ring (19109, -3.65 DPS) [world]; Blessed Band of Light (272407, -3.93 DPS) [vendor]; Channeler's Ring (272406, -4.81 DPS, sim-verified) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (498.6 DPS) | yes | Frozen Heart of the Mountain (249469, -1.81 DPS) [crafted]; Uther's Strength (11302, -2.69 DPS) [world_drop]; Weakness Analyzer (272438, -2.83 DPS, sim-verified) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (498.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (498.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -1.09 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -37.44 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 234.2 spell_power points (78.78 DPS) | yes | Brilliant Wand (249385, -14.19 DPS) [crafted]; Stormrager (16997, -14.67 DPS) [quest]; Wand of Biting Cold (19108, -14.78 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Serenity Field; trinket2: Talisman of Ascendance; ranged: Torch of Light

No-known-source sample (15 of 876, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 153005113100011531-03000000000000000-0545000300000000000)

Set DPS (verified): 737.7. Weights run: 1.2s. Verify run: 4.1s. 876 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.147 ± 0.007, crit=0.461 ± 0.014 per rating point (14 rating = 1%, 6.447 per %), hit=0.990 ± 0.047 per rating point (10 rating = 1%, 9.901 per %), spell_haste=4.536 ± 0.186, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.894 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.1 spell_power points (15.87 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.52 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.03 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.60 DPS) [quest]; Jewel of Kajaro (19601, -5.41 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.6 spell_power points (18.19 DPS) | yes | Field Marshal's Silk Spaulders (231602, -4.31 DPS) [pvp]; Argent Shoulders (19059, -5.43 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.84 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.1 spell_power points (13.81 DPS) | yes | Hide of the Wild (18510, -5.92 DPS) [crafted]; Spritecaster Cape (11623, -6.22 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -6.67 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.2 spell_power points (24.59 DPS) | yes | Field Marshal's Silk Vestments (231603, -3.20 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.32 DPS) [pvp]; Robe of the Magi (1716, -15.89 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.2 spell_power points (11.82 DPS) | yes | Runecloth Cuffs (254123, -5.46 DPS) [crafted]; Arcane Runed Bracers (4744, -7.23 DPS) [quest]; Spidertank Oilrag (9448, -7.23 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (14.15 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.1 spell_power points (19.45 DPS) | yes | Belt of the Archmage (18405, -4.76 DPS) [crafted]; Magician's Cord (272393, -7.96 DPS) [vendor]; Highlander's Cloth Girdle (20047, -8.57 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.4 spell_power points (23.65 DPS) | yes | Marshal's Silk Leggings (231605, -3.56 DPS) [pvp]; Sorcerer's Leggings (226933, -6.66 DPS, sim-verified) [quest]; Knight-Captain's Silk Legguards (227109, -8.38 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.24 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.33 DPS) [quest]; Sorcerer's Sandals (226931, -0.33 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 39.2 spell_power points (20.01 DPS) | yes | Elemental Focus Band (20682, -7.96 DPS) [world]; Don Julio's Band (19325, -11.67 DPS) [rep]; Lorekeeper's Ring (19522, -12.87 DPS) [rep] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (737.7 DPS) | yes | Don Julio's Band (19325, -1.81 DPS) [rep]; Lorekeeper's Ring (19522, -3.01 DPS) [rep]; Elemental Focus Band (20682, -8.34 DPS, sim-verified) [world] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.60 DPS) [crafted]; Weakness Analyzer (272438, -3.34 DPS, sim-verified) [vendor]; Uther's Strength (11302, -4.08 DPS) [world_drop] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.12 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -45.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 156.5 spell_power points (79.82 DPS) | yes | Brilliant Wand (249385, -15.13 DPS) [crafted]; Wand of Biting Cold (19108, -15.82 DPS) [quest]; Stormrager (16997, -16.76 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket1: Serenity Field; trinket2: Talisman of Ascendance; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 876, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (orc, 153002000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 27.9. Weights run: 1.5s. Verify run: 1.4s. 136 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.306 ± 0.011, crit=0.079 ± 0.003 per rating point (14 rating = 1%, 1.108 per %), hit=0.300 ± 0.014 per rating point (10 rating = 1%, 3.004 per %), spell_haste=not significant (0.455 ± 0.116), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.465 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -0.65 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Roadwatcher's Confidence (281265, -0.56 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.8 spell_power points (0.73 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.14 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.36 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Black Whelp Cloak (7283, +0.00 DPS, sim-verified) [crafted]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Pearl-clasped Cloak (5542, -0.10 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.62 DPS) | yes | Gray Woolen Robe (2585, -0.09 DPS) [crafted]; Green Woolen Vest (2582, -0.24 DPS) [crafted]; Bloody Apron (6226, -0.24 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.28 DPS) | yes | Owlbeard Bracers (16981, -0.13 DPS) [quest]; Mindthrust Bracers (1974, -0.14 DPS) [dungeon]; Featherbead Bracers (15452, -0.14 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.66 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.20 DPS) [crafted]; Apothecary Gloves (10919, -0.28 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.49 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.34 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk-threaded Trousers (1929, -0.08 DPS) [dungeon]; Rumpled Kilt (274741, -0.27 DPS) [vendor]; Abomination Skin Leggings (23173, -0.72 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.78 DPS) | yes | Red Woolen Boots (4313, -0.30 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.41 DPS) [crafted]; Woolen Boots (2583, -0.53 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.47 DPS) | yes | Loop of Sacrifice (281673, -0.33 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; The 1 Ring (8350, -0.44 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.8 spell_power points (0.17 DPS) | yes | Loop of Sacrifice (281673, -0.03 DPS) [quest]; Volcanic Rock Ring (12053, -0.09 DPS) [world_drop]; The 1 Ring (8350, -0.14 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.1 spell_power points (0.29 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Clear Crystal Rod (16894, -0.14 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 238.7 spell_power points (22.59 DPS) | yes | Skycaller (12984, -0.23 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.47 DPS) [quest]; Greater Magic Wand (11288, -4.90 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 136, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (orc, 153005113100010000-00000000000000000-0000000000000000000)

Set DPS (verified): 97.6. Weights run: 1.4s. Verify run: 1.3s. 242 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.392 ± 0.021, crit=0.067 ± 0.002 per rating point (14 rating = 1%, 0.937 per %), hit=0.251 ± 0.039 per rating point (10 rating = 1%, 2.510 per %), spell_haste=not significant (0.535 ± 0.325), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.902 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.95 DPS) | yes | Silk Headband (7050, -0.54 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.80 DPS) [crafted]; Enchanter's Cowl (4322, -4.90 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (2.50 DPS) | yes | Crystal Starfire Medallion (5003, -2.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.08 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.35 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (3.35 DPS) | yes | Mantle of Woe (7750, -0.76 DPS) [quest]; Death Speaker Mantle (6685, -0.77 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.80 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.34 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Battle Healer's Cloak (19529, -0.27 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.1 spell_power points (3.77 DPS) | yes | Tree Bark Jacket (1486, -0.29 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.37 DPS) [dungeon]; Beguiler Robes (7728, -0.52 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.41 DPS) | yes | Tabitha's Cuffs (251486, -1.61 DPS) [quest]; Nightsky Wristbands (6407, -1.78 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.78 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.0 spell_power points (2.13 DPS) | yes | Serpent Gloves (5970, -0.26 DPS) [dungeon]; Truefaith Gloves (7049, -0.48 DPS) [crafted]; Gnoll Casting Gloves (892, -0.52 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.2 spell_power points (3.26 DPS) | yes | Warsong Sash (16975, -0.31 DPS) [quest]; Belt of Arugal (6392, -0.54 DPS) [dungeon]; Invoker's Cord (215366, -0.86 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.1 spell_power points (3.25 DPS) | yes | Gaze Dreamer Pants (6903, -0.04 DPS) [dungeon]; Pristine Leggings (253987, -0.64 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.01 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (2.61 DPS) | yes | Acidic Walkers (9454, -0.43 DPS) [dungeon]; Boots of the Enchanter (4325, -1.27 DPS) [crafted]; Spidersilk Boots (4320, -2.69 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.87 DPS) | yes | Electrocutioner Lagnut (9447, -1.07 DPS) [dungeon]; Snake Hoop (6750, -1.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.24 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.61 DPS) | yes | Snake Hoop (6750, -0.87 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.98 DPS) [dungeon]; Electrocutioner Lagnut (9447, -2.20 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.3 spell_power points (1.15 DPS) | yes | Twisted Chanter's Staff (890, -0.10 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.10 DPS) [quest]; Polished Walking Staff (16889, -0.42 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Unstable Power Core (279847) | Source of Power [quest] | 127.4 spell_power points (34.12 DPS) | yes | Starfaller (13063, -0.84 DPS) [world_drop]; Greater Mystic Wand (217287, -3.78 DPS) [crafted]; Gravestone Scepter (7001, -5.12 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Glimmering Staff; ranged: Unstable Power Core

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (orc, 153005113100011531-00000000000000000-0000000000000000000)

Set DPS (verified): 176.3. Weights run: 1.4s. Verify run: 1.6s. 312 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.583 ± 0.038, crit=0.166 ± 0.005 per rating point (14 rating = 1%, 2.321 per %), hit=0.367 ± 0.078 per rating point (10 rating = 1%, 3.673 per %), spell_haste=not significant (1.299 ± 0.570), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.919 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.53 DPS) | yes | Living Cowl (5608, -2.49 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -2.59 DPS) [dungeon]; Augural Shroud (2620, -2.74 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (3.63 DPS) | yes | Scorn's Icy Choker (23169, -1.62 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.36 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.36 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.54 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.54 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.9 spell_power points (2.77 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.60 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.78 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (7.93 DPS) | yes | Dreamweave Vest (10021, -0.70 DPS) [crafted]; Robe of Power (7054, -1.40 DPS) [crafted]; Elemental Raiment (9434, -1.40 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.80 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.62 DPS) [quest]; Radiant Silver Bracers (4545, -1.85 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (6.33 DPS) | yes | Black Mageweave Gloves (10003, -1.66 DPS) [crafted]; Red Mageweave Gloves (10018, -2.44 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.57 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.3 spell_power points (5.08 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Star Belt (4329, -1.04 DPS) [crafted]; Gilded Cord (254037, -1.14 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.53 DPS) | yes | Abomination Skin Leggings (23173, -2.28 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.54 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.80 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.47 DPS) | yes | Gilded Slippers (254001, -2.81 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.46 DPS) [dungeon]; Spidersilk Boots (4320, -4.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (4.02 DPS) | yes | Reedknot Ring (9622, -1.84 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.15 DPS) [vendor]; Voodoo Band (1996, -2.75 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.80 DPS) | yes | Reedknot Ring (9622, -0.62 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.93 DPS) [vendor]; Voodoo Band (1996, -1.53 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (176.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.36 DPS) [dungeon]; Windweaver Staff (7757, -12.42 DPS) [dungeon] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (4.59 DPS) | yes | Orb of the Forgotten Seer (7685, -0.23 DPS) [dungeon]; Thrash's Trash (276204, -0.23 DPS) [vendor]; Orb of Mystic Insight (249394, -1.32 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 130.1 spell_power points (40.47 DPS) | yes | Nether Force Wand (11263, -2.59 DPS) [quest]; Icefury Wand (7514, -2.73 DPS) [quest]; Ragefire Wand (7513, -2.78 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 312, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 153005113100011531-03202300000000000-0000000000000000000)

Set DPS (verified): 284.9. Weights run: 1.4s. Verify run: 1.7s. 398 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.688 ± 0.059, crit=0.254 ± 0.009 per rating point (14 rating = 1%, 3.560 per %), hit=0.695 ± 0.123 per rating point (10 rating = 1%, 6.954 per %), spell_haste=not significant (1.201 ± 0.837), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.922 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 32.8 spell_power points (10.39 DPS) | yes | Dreamweave Circlet (10041, -1.55 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.83 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.20 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.5 spell_power points (3.97 DPS) | yes | Scorn's Icy Choker (23169, -0.44 DPS) [dungeon]; Mindburst Medallion (11196, -0.75 DPS) [quest]; Horizon Choker (13085, -0.91 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.4 spell_power points (8.05 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -2.42 DPS) [vendor]; Black Mageweave Shoulders (10027, -2.92 DPS) [crafted]; Red Mageweave Shoulders (10029, -3.08 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.2 spell_power points (5.77 DPS) | yes | Spritecaster Cape (11623, -0.02 DPS) [dungeon]; Runecloth Cloak (13860, -1.17 DPS) [crafted]; Big Voodoo Cloak (8216, -2.22 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 32.8 spell_power points (10.39 DPS) | yes | Robe of the Magi (1716, -2.10 DPS) [world_drop]; Runecloth Tunic (13857, -2.60 DPS) [crafted]; Dreamweave Vest (10021, -2.72 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.8 spell_power points (3.75 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -3.47 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.6 spell_power points (9.07 DPS) | yes | Raider Handwraps (272098, -2.05 DPS) [vendor]; Dreamweave Gloves (10019, -2.48 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.98 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 20.9 spell_power points (6.62 DPS) | yes | Dawnspire Cord (12466, -0.57 DPS) [dungeon]; Deathmage Sash (10771, -1.13 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.31 DPS) [rep] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 23.8 spell_power points (7.56 DPS) | yes | Red Mageweave Pants (10009, -0.49 DPS) [crafted]; Wizardweave Leggings (14132, -1.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.18 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.62 DPS) | yes | Gilded Sandals (254107, -2.16 DPS) [crafted]; Black Mageweave Boots (10026, -2.60 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.10 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (4.26 DPS) | yes | Advisor's Ring (19519, -0.46 DPS) [rep]; Reedknot Ring (9622, -2.04 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.13 DPS) | yes | Advisor's Ring (19519, -0.32 DPS) [rep]; Reedknot Ring (9622, -1.90 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (284.9 DPS) | yes | Rune of the Guard Captain (19120, -0.44 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (284.9 DPS) | yes | Rune of the Guard Captain (19120, -2.72 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (284.9 DPS) | yes | Kindling Stave (11750, -0.50 DPS) [dungeon]; Spire of Hakkar (10844, -2.48 DPS) [world]; Hanzo Sword (8190, -33.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 165.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.21 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Stone Guard's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 398, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 482.9. Weights run: 1.2s. Verify run: 3.6s. 866 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.841 ± 0.067, crit=0.407 ± 0.013 per rating point (14 rating = 1%, 5.699 per %), hit=0.958 ± 0.163 per rating point (10 rating = 1%, 9.579 per %), spell_haste=not significant (2.418 ± 1.134), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.914 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.7 spell_power points (16.05 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.98 DPS) [pvp]; Magister's Crown (16686, -7.29 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (482.9 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.96 DPS) [quest]; Chains of the Lich (23125, -1.32 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.3 spell_power points (15.24 DPS) | yes | Warlord's Silk Amice (231594, -2.59 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.93 DPS) [crafted]; Darkspear Shoulderpads (272103, -8.26 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.3 spell_power points (10.87 DPS) | yes | Hide of the Wild (18510, -3.33 DPS) [crafted]; Deep Woodlands Cloak (19121, -4.28 DPS) [quest]; Spritecaster Cape (11623, -4.46 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.8 spell_power points (18.77 DPS) | yes | Warlord's Silk Raiment (231596, -0.94 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.98 DPS) [pvp]; Sorcerer's Robes (226932, -12.81 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.7 spell_power points (9.66 DPS) | yes | Runecloth Cuffs (254123, -3.13 DPS) [crafted]; General's Silk Cuffs (16538, -4.29 DPS) [pvp]; Sorcerer's Bindings (226929, -4.48 DPS) [quest] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 33.4 spell_power points (11.22 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.1 spell_power points (17.86 DPS) | yes | Belt of the Archmage (18405, -4.68 DPS) [crafted]; Frostwolf Cloth Belt (19090, -8.97 DPS) [rep]; Magician's Cord (272393, -9.95 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.8 spell_power points (17.43 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -3.29 DPS) [quest]; Outrider's Silk Leggings (22747, -7.52 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.5 spell_power points (11.59 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.47 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.1 spell_power points (15.19 DPS) | yes | Channeler's Ring (272406, -8.00 DPS) [vendor]; Deep Rooted Ring (19109, -8.32 DPS) [world]; Blessed Band of Light (272407, -8.60 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 31.3 spell_power points (10.51 DPS) | yes | Deep Rooted Ring (19109, -3.65 DPS) [world]; Channeler's Ring (272406, -3.87 DPS, sim-verified) [vendor]; Blessed Band of Light (272407, -3.93 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (482.9 DPS) | yes | Frozen Heart of the Mountain (249469, -1.81 DPS) [crafted]; Rune of the Guard Captain (19120, -2.45 DPS) [quest]; Weakness Analyzer (272438, -2.77 DPS, sim-verified) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (482.9 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (482.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -1.09 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -36.81 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 234.2 spell_power points (78.78 DPS) | yes | Brilliant Wand (249385, -14.19 DPS) [crafted]; Stormrager (16997, -14.67 DPS) [quest]; Wand of Biting Cold (19108, -14.78 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Serenity Field; trinket2: Talisman of Ascendance; ranged: Torch of Light

No-known-source sample (15 of 866, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 153005113100011531-03000000000000000-0545000300000000000)

Set DPS (verified): 721.0. Weights run: 1.2s. Verify run: 3.8s. 866 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.147 ± 0.007, crit=0.461 ± 0.014 per rating point (14 rating = 1%, 6.447 per %), hit=0.990 ± 0.047 per rating point (10 rating = 1%, 9.901 per %), spell_haste=4.536 ± 0.186, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.894 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.1 spell_power points (15.87 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.52 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.09 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.60 DPS) [quest]; Jewel of Kajaro (19601, -5.68 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.6 spell_power points (18.19 DPS) | yes | Warlord's Silk Amice (231594, -4.31 DPS) [pvp]; Argent Shoulders (19059, -5.43 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.60 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.1 spell_power points (13.81 DPS) | yes | Hide of the Wild (18510, -5.92 DPS) [crafted]; Spritecaster Cape (11623, -6.22 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -6.67 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.2 spell_power points (24.59 DPS) | yes | Warlord's Silk Raiment (231596, -3.20 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -9.32 DPS) [pvp]; Robe of the Magi (1716, -15.65 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.2 spell_power points (11.82 DPS) | yes | Runecloth Cuffs (254123, -5.46 DPS) [crafted]; Spidertank Oilrag (9448, -7.23 DPS) [dungeon]; Sorcerer's Bindings (226929, -7.50 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (14.15 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.1 spell_power points (19.45 DPS) | yes | Belt of the Archmage (18405, -5.45 DPS, sim-verified) [crafted]; Magician's Cord (272393, -7.96 DPS) [vendor]; Defiler's Cloth Girdle (20163, -8.57 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.4 spell_power points (23.65 DPS) | yes | General's Silk Trousers (231595, -3.56 DPS) [pvp]; Sorcerer's Leggings (226933, -6.88 DPS, sim-verified) [quest]; Outrider's Silk Leggings (22747, -7.94 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.24 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.33 DPS) [quest]; Sorcerer's Sandals (226931, -0.33 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 39.2 spell_power points (20.01 DPS) | yes | Elemental Focus Band (20682, -7.96 DPS) [world]; Don Julio's Band (19325, -11.67 DPS) [rep]; Advisor's Ring (19518, -12.87 DPS) [rep] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (721.0 DPS) | yes | Don Julio's Band (19325, -1.81 DPS) [rep]; Advisor's Ring (19518, -3.01 DPS) [rep]; Elemental Focus Band (20682, -8.33 DPS, sim-verified) [world] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.60 DPS) [crafted]; Weakness Analyzer (272438, -3.29 DPS, sim-verified) [vendor]; Rune of the Guard Captain (19120, -3.61 DPS) [quest] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Darkmoon Card: Blue Dragon (19288, -5.08 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.12 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -45.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 156.5 spell_power points (79.82 DPS) | yes | Brilliant Wand (249385, -15.13 DPS) [crafted]; Wand of Biting Cold (19108, -15.82 DPS) [quest]; Stormrager (16997, -16.76 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket1: Serenity Field; trinket2: Talisman of Ascendance; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 866, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

