# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 28.6. Weights run: 1.3s. Verify run: 1.2s. 196 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.521 ± 0.013, crit=0.084 ± 0.004 per rating point (14 rating = 1%, 1.182 per %), hit=0.333 ± 0.014 per rating point (10 rating = 1%, 3.327 per %), spell_haste=not significant (0.387 ± 0.198), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.50 DPS) | yes | Shadow Goggles (4373, -1.12 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.7 spell_power points (0.80 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.22 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.47 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Caretaker's Cape (20428, -0.08 DPS) [rep]; Pearl-clasped Cloak (5542, -0.18 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.6 spell_power points (0.63 DPS) | yes | Green Woolen Robe (6243, -0.25 DPS) [crafted]; Green Woolen Vest (2582, -0.30 DPS) [crafted]; Gray Woolen Robe (2585, -0.51 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.6 spell_power points (0.22 DPS) | yes | Bright Bracers (3647, -0.04 DPS) [world_drop]; Mystic's Bracelets (14366, -0.13 DPS) [world_drop]; Repurposed Hair Band (281256, -0.13 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.58 DPS) | yes | Pristine Gloves (253913, -0.12 DPS) [crafted]; Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.32 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.1 spell_power points (0.50 DPS) | yes | Keller's Girdle (2911, -0.16 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.36 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.2 spell_power points (1.09 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.51 DPS) [dungeon]; Rumpled Kilt (274741, -0.67 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.1 spell_power points (0.75 DPS) | yes | Pristine Boots (253889, -0.10 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.42 DPS) [crafted]; Woolen Boots (2583, -0.50 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.0 spell_power points (0.50 DPS) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.46 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.41 DPS) | yes | Volcanic Rock Ring (12053, -0.28 DPS) [world_drop]; The 1 Ring (8350, -0.37 DPS) [world]; Lavishly Jeweled Ring (1156, -0.43 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 5.2 spell_power points (0.43 DPS) | yes | Lesser Staff of the Spire (1300, -0.17 DPS) [world]; Witching Stave (1484, -0.22 DPS) [dungeon]; Staff of Westfall (2042, -0.22 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 273.1 spell_power points (22.56 DPS) | yes | Skycaller (12984, -1.00 DPS) [world_drop]; Deepblaze (279896, -4.03 DPS) [quest]; Sizzle Stick (8071, -4.49 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash

### Band 30 (gnome, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 59.7. Weights run: 1.3s. Verify run: 1.3s. 317 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.554 ± 0.010, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.375 per %), hit=0.220 ± 0.011 per rating point (10 rating = 1%, 2.200 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Enchanter's Cowl (4322, -1.33 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.17 DPS) | yes | Darkspear Warding Pendant (272075, -0.91 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.92 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.92 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.59 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Magician's Mantle (12998, -0.45 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Hillman's Cloak (3719, -0.02 DPS) [crafted]; Darkspear Raider's Cloak (272078, -0.09 DPS) [vendor]; Vine Pruner's Cloak (279835, -0.72 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.2 spell_power points (1.84 DPS) | yes | Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Beguiler Robes (7728, -0.31 DPS) [dungeon]; Death Speaker Robes (6682, -0.35 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.02 DPS) | yes | Nightsky Wristbands (6407, -0.64 DPS) [world_drop]; Stonecloth Bindings (14416, -0.71 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.95 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.1 spell_power points (1.15 DPS) | yes | Serpent Gloves (5970, -0.35 DPS) [dungeon]; Truefaith Gloves (7049, -0.39 DPS) [crafted]; Shilly Mitts (9609, -0.59 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (1.44 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Crimson Silk Belt (7055, -0.32 DPS) [crafted]; Invoker's Cord (215366, -0.33 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.4 spell_power points (1.52 DPS) | yes | Pristine Leggings (253987, -0.29 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.30 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.47 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.23 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Nimbus Boots (6998, -0.55 DPS) [quest]; Acidic Walkers (9454, -1.13 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.79 DPS) | yes | Snake Hoop (6750, -0.35 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Minor Channeling Ring (1449, -1.31 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.24 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon]; Minor Channeling Ring (1449, -1.07 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | 30.0 spell_power points (3.40 DPS) | yes | Twisted Chanter's Staff (890, -2.78 DPS) [world_drop]; Lesser Staff of the Spire (1300, -3.03 DPS) [world]; Glimmering Staff (249392, -4.96 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 291.8 spell_power points (33.11 DPS) | yes | Greater Mystic Wand (217287, -3.54 DPS) [crafted]; Gravestone Scepter (7001, -4.11 DPS) [quest]; Scorching Wand (5213, -4.26 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff; ranged: Starfaller

No-known-source sample (15 of 317, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff

### Band 40 (gnome, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 127.3. Weights run: 1.3s. Verify run: 1.4s. 400 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.088 ± 0.017, crit=0.014 ± 0.001 per rating point (14 rating = 1%, 0.189 per %), hit=0.283 ± 0.021 per rating point (10 rating = 1%, 2.829 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 21.9 spell_power points (4.16 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.17 DPS) [crafted]; Thinking Cap (2624, -0.64 DPS) [world]; Enchanter's Cowl (4322, -0.95 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (2.99 DPS) | yes | Scorn's Icy Choker (23169, -1.52 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.54 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.54 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.1 spell_power points (4.02 DPS) | yes | Bloodmage Mantle (7684, -0.45 DPS) [dungeon]; Death Speaker Mantle (6685, -0.60 DPS) [dungeon]; Green Silken Shoulders (7057, -1.43 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 12.0 spell_power points (2.28 DPS) | yes | Long Silken Cloak (4326, -0.10 DPS) [crafted]; Guardian Cloak (5965, -0.10 DPS) [crafted]; Blackforge Cape (6424, -0.62 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.5 spell_power points (5.43 DPS) | yes | Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.28 DPS) [crafted]; Green Silk Armor (7065, -1.02 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 9.8 spell_power points (1.86 DPS) | yes | Arcane Runed Bracers (4744, -0.15 DPS) [quest]; Spidertank Oilrag (9448, -0.15 DPS) [dungeon]; Mistscape Bracers (4045, -0.21 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.4 spell_power points (4.25 DPS) | yes | Red Mageweave Gloves (10018, -0.72 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -1.01 DPS) [crafted]; Town Clerk's Mittens (270029, -1.21 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.3 spell_power points (4.44 DPS) | yes | Highlander's Cloth Girdle (20098, -0.91 DPS, sim-verified) [rep]; Gilded Cord (254037, -1.26 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.1 spell_power points (5.15 DPS) | yes | Crimson Silk Pantaloons (7062, -1.46 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.78 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.16 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.56 DPS) | yes | Gilded Slippers (254001, -1.32 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.96 DPS) [dungeon]; Spidersilk Boots (4320, -2.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.94 DPS) | yes | Ring of Forlorn Spirits (2043, -1.42 DPS) [quest]; Ogremind Ring (1993, -1.49 DPS) [world_drop]; Voodoo Band (1996, -1.49 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.71 DPS) | yes | Ogremind Ring (1993, -0.26 DPS) [world_drop]; Voodoo Band (1996, -0.26 DPS) [world]; Ring of Forlorn Spirits (2043, -1.82 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (127.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.41 DPS) [dungeon]; Hand of Righteousness (7721, -4.13 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 16.9 spell_power points (3.21 DPS) | yes | Thrash's Trash (276204, -0.55 DPS) [vendor]; Orb of Mystic Insight (249394, -0.64 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.24 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 213.6 spell_power points (40.62 DPS) | yes | Umbral Wand (5216, -4.95 DPS) [dungeon]; Twisted Nether Wand (249144, -5.48 DPS) [crafted]; Burning Sliver (5249, -6.79 DPS) [quest] |

**New at 40:** head: Augural Shroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers; 3225 Bloodstained Knife

### Band 50 (gnome, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 212.3. Weights run: 1.3s. Verify run: 3.1s. 494 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.454 ± 0.027, crit=0.026 ± 0.002 per rating point (14 rating = 1%, 0.368 per %), hit=0.537 ± 0.039 per rating point (10 rating = 1%, 5.372 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.1 spell_power points (8.48 DPS) | yes | Dreamweave Circlet (10041, -1.61 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.87 DPS) [vendor]; Bad Mojo Mask (9470, -3.10 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.4 spell_power points (3.59 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Scorn's Icy Choker (23169, -0.82 DPS) [dungeon]; Mindburst Medallion (11196, -0.99 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.2 spell_power points (6.91 DPS) | yes | Red Mageweave Shoulders (10029, -2.25 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -2.34 DPS) [dungeon]; Green Silken Shoulders (7057, -2.68 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.7 spell_power points (4.01 DPS) | yes | Darkspear Raider's Cloak (272076, -0.42 DPS) [vendor]; Big Voodoo Cloak (8216, -0.82 DPS) [crafted]; Runecloth Cloak (13860, -1.03 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.1 spell_power points (8.48 DPS) | yes | Runecloth Tunic (13857, -2.66 DPS) [crafted]; Runecloth Robe (13858, -2.82 DPS, sim-verified) [crafted]; Robe of Power (7054, -2.93 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 18.1 spell_power points (3.19 DPS) | yes | Forgotten Wraps (9433, -0.11 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.11 DPS) [quest]; Nethergeld Cuffs (254061, -0.16 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.0 spell_power points (6.88 DPS) | yes | Virtuous Hands (226958, -0.19 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.28 DPS) [vendor]; Red Mageweave Gloves (10018, -2.37 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 33.6 spell_power points (5.93 DPS) | yes | Satyrmane Sash (17755, -0.90 DPS) [dungeon]; Runecloth Belt (13856, -1.44 DPS) [crafted]; Deathmage Sash (10771, -1.83 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 31.4 spell_power points (5.55 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.29 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.40 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.80 DPS) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthen Silk Slippers (254013, -0.01 DPS) [crafted]; Southsea Mojo Boots (20641, -0.02 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -2.56 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.3 spell_power points (3.05 DPS) | yes | Ogremind Ring (1993, -1.25 DPS) [world_drop]; Voodoo Band (1996, -1.25 DPS) [world]; Band of the Unicorn (7553, -4.15 DPS, sim-verified) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Ogremind Ring (1993, -0.32 DPS) [world_drop]; Voodoo Band (1996, -0.32 DPS) [world]; Band of the Unicorn (7553, -2.11 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -3.39 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.76 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -3.33 DPS) [dungeon]; Hammer of the Northern Wind (810, -25.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.46 DPS) [world_drop]; Pyric Caduceus (11748, -3.23 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.72 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Dawnspire Cord; feet: Gilded Sandals; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Noxious Shooter

No-known-source sample (15 of 494, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers; 3225 Bloodstained Knife

### Band 60 (gnome, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 415.9. Weights run: 1.8s. Verify run: 8.2s. 1005 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.292 ± 0.025, crit=0.026 ± 0.002 per rating point (14 rating = 1%, 0.362 per %), hit=0.585 ± 0.040 per rating point (10 rating = 1%, 5.850 per %), spell_haste=not significant (-0.002 ± 0.013), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | sim-verified (415.9 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Magister's Crown (16686, -13.13 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 31.8 spell_power points (8.85 DPS) | yes | Beads of Ogre Mojo (22149, -0.68 DPS, sim-verified) [quest]; Archlight Talisman (15856, -1.91 DPS) [quest]; Chains of the Lich (23125, -2.72 DPS) [dungeon] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | sim-verified (415.9 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -6.89 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.2 spell_power points (8.96 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -2.90 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -3.20 DPS) [vendor] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 49.2 spell_power points (13.70 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -12.53 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 32.3 spell_power points (9.00 DPS) | yes | Marshal's Satin Bracers (17606, -2.17 DPS) [pvp]; Runecloth Cuffs (254123, -2.34 DPS) [crafted]; Virtuous Wraps (226953, -2.82 DPS) [vendor] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | sim-verified (415.9 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Raider Handwraps (272097, -8.78 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 59.3 spell_power points (16.49 DPS) | yes | Belt of the Archmage (18405, -3.51 DPS, sim-verified) [crafted]; Magister's Belt (16685, -7.00 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -7.89 DPS) [rep] |
| legs | Virtuous Leggings (226956) | Mokvar [vendor] | sim-verified (415.9 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (22752, -6.78 DPS, sim-verified) [rep] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 42.7 spell_power points (11.87 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Omnicast Boots (11822, -8.51 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.5 spell_power points (12.65 DPS) | yes | Deep Rooted Ring (19109, -5.72 DPS) [world]; Channeler's Ring (272406, -5.99 DPS) [vendor]; Band of Piety (22681, -7.11 DPS) [quest] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 36.2 spell_power points (10.08 DPS) | yes | Deep Rooted Ring (19109, -3.40 DPS, sim-verified) [world]; Channeler's Ring (272406, -3.42 DPS) [vendor]; Band of Piety (22681, -4.53 DPS) [quest] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (415.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (415.9 DPS) | yes | Weakness Analyzer (272438, -2.04 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.23 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.43 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (415.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -2.28 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (415.9 DPS) | yes | Stormrager (16997, -0.11 DPS) [quest]; Wand of Biting Cold (19108, -0.60 DPS) [quest]; Torch of Light (279246, -8.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Virtuous Hands; waist: Knowledge of the Timbermaw; legs: Virtuous Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 1005, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers; 3225 Bloodstained Knife

### Band 60, raid preset (gnome, 325001031300000000-00000000000000000-523120501201300251)

Set DPS (verified): 573.1. Weights run: 1.6s. Verify run: 5.2s. 1005 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.014 ± 0.002, crit=0.062 ± 0.006 per rating point (14 rating = 1%, 0.864 per %), hit=0.459 ± 0.050 per rating point (10 rating = 1%, 4.590 per %), spell_haste=-7.784 ± 0.506, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 27.1 spell_power points (10.05 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -0.05 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.15 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.52 DPS) [quest]; Beads of Ogre Mojo (22149, -3.27 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.1 spell_power points (10.40 DPS) | yes | Field Marshal's Satin Epaulets (231621, -1.03 DPS) [vendor]; Argent Shoulders (19059, -1.17 DPS, sim-verified) [crafted]; Virtuous Epaulets (226955, -1.60 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.7 spell_power points (7.67 DPS) | yes | Hide of the Wild (18510, -2.43 DPS) [crafted]; Spritecaster Cape (11623, -2.45 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -2.48 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-verified (573.1 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Virtuous Gown (226960, -8.49 DPS, sim-verified) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.1 spell_power points (8.19 DPS) | yes | Runecloth Cuffs (254123, -4.06 DPS) [crafted]; Arcane Runed Bracers (4744, -4.86 DPS) [quest]; Spidertank Oilrag (9448, -4.86 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.15 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Knight-Lieutenant's Satin Grips (227139, -0.34 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 29.9 spell_power points (11.08 DPS) | yes | Stormpike Cloth Girdle (19094, -4.36 DPS) [rep]; Belt of the Archmage (18405, -5.27 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -5.54 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 33.9 spell_power points (12.55 DPS) | yes | Marshal's Satin Pants (17603, -0.58 DPS) [vendor]; Marshal's Satin Leggings (231619, -0.58 DPS) [vendor]; Ghostweave Pants (14144, -2.92 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.89 DPS) | yes | Marshal's Satin Sandals (17607, -0.27 DPS) [vendor]; Marshal's Satin Treads (231620, -0.27 DPS) [vendor]; Virtuous Slippers (226959, -0.66 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 32.7 spell_power points (12.12 DPS) | yes | Blessed Band of Light (272407, -6.72 DPS) [vendor]; Lorekeeper's Ring (19522, -6.93 DPS) [rep]; Band of the Unicorn (7553, -7.30 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.2 spell_power points (8.21 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Lorekeeper's Ring (19522, -3.02 DPS) [rep]; Band of the Unicorn (7553, -3.39 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+17.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.21 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.96 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.66 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.68 DPS) [dungeon]; The Lobotomizer (19324, -30.60 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 213.2 spell_power points (78.98 DPS) | yes | Wand of Biting Cold (19108, +0.00 DPS, sim-verified) [quest]; Brilliant Wand (249385, -16.03 DPS) [crafted]; Stormrager (16997, -16.27 DPS) [quest] |

**New at 60:** head: Virtuous Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Truefaith Vestments; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Earthen Silk Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1005, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers; 3225 Bloodstained Knife

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 33.5. Weights run: 1.3s. Verify run: 1.2s. 186 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.521 ± 0.013, crit=0.084 ± 0.004 per rating point (14 rating = 1%, 1.182 per %), hit=0.333 ± 0.014 per rating point (10 rating = 1%, 3.327 per %), spell_haste=not significant (0.387 ± 0.198), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.50 DPS) | yes | Shadow Goggles (4373, -1.04 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.7 spell_power points (0.80 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.47 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.04 DPS) [crafted]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Battle Healer's Cloak (20427, -0.08 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.6 spell_power points (0.63 DPS) | yes | Green Woolen Robe (6243, -0.25 DPS) [crafted]; Green Woolen Vest (2582, -0.30 DPS) [crafted]; Gray Woolen Robe (2585, -0.70 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.25 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest]; Bright Bracers (3647, -0.08 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.58 DPS) | yes | Gnoll Casting Gloves (892, -0.08 DPS) [world]; Pristine Gloves (253913, -0.12 DPS) [crafted]; Apothecary Gloves (10919, -0.25 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.1 spell_power points (0.50 DPS) | yes | Keller's Girdle (2911, -0.16 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.49 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.2 spell_power points (1.09 DPS) | yes | Filigreed Pristine Leggings (253937, -0.33 DPS) [crafted]; Silk-threaded Trousers (1929, -0.51 DPS) [dungeon]; Rumpled Kilt (274741, -0.67 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.1 spell_power points (0.75 DPS) | yes | Pristine Boots (253889, -0.16 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.42 DPS) [crafted]; Woolen Boots (2583, -0.50 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.41 DPS) | yes | Loop of Sacrifice (281673, -0.20 DPS) [quest]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop]; The 1 Ring (8350, -0.37 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.1 spell_power points (0.26 DPS) | yes | Loop of Sacrifice (281673, -0.04 DPS) [quest]; Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; The 1 Ring (8350, -0.22 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 5.2 spell_power points (0.43 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Lesser Staff of the Spire (1300, -0.17 DPS) [world]; Clear Crystal Rod (16894, -0.22 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 273.1 spell_power points (22.56 DPS) | yes | Skycaller (12984, -2.11 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.49 DPS) [quest]; Greater Magic Wand (11288, -4.89 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff

### Band 30 (undead, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 53.3. Weights run: 1.3s. Verify run: 1.3s. 304 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.554 ± 0.010, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.375 per %), hit=0.220 ± 0.011 per rating point (10 rating = 1%, 2.200 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Enchanter's Cowl (4322, -1.10 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.17 DPS) | yes | Crystal Starfire Medallion (5003, -0.92 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.92 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.01 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.59 DPS) | yes | Death Speaker Mantle (6685, -0.24 DPS, sim-verified) [dungeon]; Mantle of Woe (7750, -0.27 DPS) [quest]; Fairywing Mantle (9536, -0.34 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.57 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.06 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.2 spell_power points (1.84 DPS) | yes | Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Beguiler Robes (7728, -0.31 DPS) [dungeon]; Death Speaker Robes (6682, -0.35 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.02 DPS) | yes | Nightsky Wristbands (6407, -0.64 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.68 DPS) [quest]; Glowing Magical Bracelets (13106, -0.89 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.8 spell_power points (1.00 DPS) | yes | Truefaith Gloves (7049, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.31 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.7 spell_power points (1.44 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Crimson Silk Belt (7055, -0.32 DPS) [crafted]; Warsong Sash (16975, -0.32 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.4 spell_power points (1.52 DPS) | yes | Pristine Leggings (253987, -0.29 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.47 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.47 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.23 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Boots of the Enchanter (4325, -0.67 DPS) [crafted]; Acidic Walkers (9454, -1.23 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.79 DPS) | yes | Snake Hoop (6750, -0.35 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.68 DPS) | yes | Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Snake Hoop (6750, -0.84 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 6.1 spell_power points (0.69 DPS) | yes | Twisted Chanter's Staff (890, -0.06 DPS) [world_drop]; Polished Walking Staff (16889, -0.25 DPS) [quest]; Gnarled Necromancer's Staff (251534, -0.39 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Unstable Power Core (279847, -1.25 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.54 DPS) [crafted]; Gravestone Scepter (7001, -4.11 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Glimmering Staff; ranged: Starfaller

No-known-source sample (15 of 304, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers; 3225 Bloodstained Knife

### Band 40 (undead, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 119.3. Weights run: 1.3s. Verify run: 1.5s. 384 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.088 ± 0.017, crit=0.014 ± 0.001 per rating point (14 rating = 1%, 0.189 per %), hit=0.283 ± 0.021 per rating point (10 rating = 1%, 2.829 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 21.9 spell_power points (4.16 DPS) | yes | Thinking Cap (2624, -0.64 DPS) [world]; Enchanter's Cowl (4322, -0.95 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.30 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (2.99 DPS) | yes | Necklace of Calisea (1714, -1.54 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.54 DPS) [vendor]; Scorn's Icy Choker (23169, -1.83 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.1 spell_power points (4.02 DPS) | yes | Bloodmage Mantle (7684, -0.45 DPS) [dungeon]; Mantle of Woe (7750, -0.59 DPS) [quest]; Green Silken Shoulders (7057, -1.86 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 12.0 spell_power points (2.28 DPS) | yes | Guardian Cloak (5965, -0.10 DPS) [crafted]; Blackforge Cape (6424, -0.62 DPS) [dungeon]; Long Silken Cloak (4326, -1.19 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.5 spell_power points (5.43 DPS) | yes | Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.28 DPS) [crafted]; Zealot's Robe (17043, -0.50 DPS) [quest] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.7 spell_power points (2.42 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.70 DPS) [dungeon]; Windchaser Cuffs (14429, -0.74 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.4 spell_power points (4.25 DPS) | yes | Red Mageweave Gloves (10018, -0.68 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -1.01 DPS) [crafted]; Gilded Handwraps (254021, -1.28 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.3 spell_power points (4.44 DPS) | yes | Gilded Cord (254037, -1.26 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.72 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.1 spell_power points (5.15 DPS) | yes | Crimson Silk Pantaloons (7062, -1.24 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.78 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.16 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.56 DPS) | yes | Gilded Slippers (254001, -0.88 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.96 DPS) [dungeon]; Spidersilk Boots (4320, -2.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.94 DPS) | yes | Ogremind Ring (1993, -1.49 DPS) [world_drop]; Voodoo Band (1996, -1.49 DPS) [world]; Mindbender Loop (5009, -1.49 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.71 DPS) | yes | Ogremind Ring (1993, -0.26 DPS) [world_drop]; Mindbender Loop (5009, -0.26 DPS) [world_drop]; Voodoo Band (1996, -2.29 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (119.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.41 DPS) [dungeon]; Skullbreaker (17039, -0.59 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.66 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Mystic Insight (249394, -0.09 DPS) [crafted]; Tome of the Darkspear Prophecy (272089, -0.87 DPS) [vendor] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 213.6 spell_power points (40.62 DPS) | yes | Umbral Wand (5216, -4.95 DPS) [dungeon]; Twisted Nether Wand (249144, -5.48 DPS) [crafted]; Dancing Flame (6806, -5.78 DPS) [quest] |

**New at 40:** head: Augural Shroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 384, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers

### Band 50 (undead, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 199.3. Weights run: 1.3s. Verify run: 2.6s. 475 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.454 ± 0.027, crit=0.026 ± 0.002 per rating point (14 rating = 1%, 0.368 per %), hit=0.537 ± 0.039 per rating point (10 rating = 1%, 5.372 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.1 spell_power points (8.48 DPS) | yes | Dreamweave Circlet (10041, -1.91 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.87 DPS) [vendor]; Bad Mojo Mask (9470, -3.10 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.4 spell_power points (3.59 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Scorn's Icy Choker (23169, -0.82 DPS) [dungeon]; Mindburst Medallion (11196, -0.99 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.2 spell_power points (6.91 DPS) | yes | Inquisitor's Shawl (19507, -2.34 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.43 DPS, sim-verified) [crafted]; Green Silken Shoulders (7057, -2.68 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 25.1 spell_power points (4.43 DPS) | yes | Spritecaster Cape (11623, -0.42 DPS) [dungeon]; Runecloth Cloak (13860, -0.79 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.83 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.1 spell_power points (8.48 DPS) | yes | Runecloth Tunic (13857, -2.66 DPS) [crafted]; Robe of Power (7054, -2.93 DPS) [crafted]; Runecloth Robe (13858, -2.94 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 18.1 spell_power points (3.19 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Shizzle's Nozzle Wiper (11917, -0.11 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.0 spell_power points (6.88 DPS) | yes | Virtuous Hands (226958, -0.19 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.28 DPS) [vendor]; Red Mageweave Gloves (10018, -2.37 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 33.6 spell_power points (5.93 DPS) | yes | Satyrmane Sash (17755, -0.90 DPS) [dungeon]; Runecloth Belt (13856, -1.44 DPS) [crafted]; Deathmage Sash (10771, -2.31 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 31.4 spell_power points (5.55 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.29 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.40 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.80 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.5 spell_power points (4.67 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Earthen Silk Slippers (254013, -0.43 DPS) [crafted]; Southsea Mojo Boots (20641, -0.43 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.3 spell_power points (3.05 DPS) | yes | Ogremind Ring (1993, -1.25 DPS) [world_drop]; Voodoo Band (1996, -1.25 DPS) [world]; Band of the Unicorn (7553, -3.83 DPS, sim-verified) [world_drop] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ogremind Ring (1993, -0.32 DPS) [world_drop]; Voodoo Band (1996, -0.32 DPS) [world]; Band of the Unicorn (7553, -2.03 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.19 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.97 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.76 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -3.33 DPS) [dungeon]; Hammer of the Northern Wind (810, -25.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.46 DPS) [world_drop]; Pyric Caduceus (11748, -3.21 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.72 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Dawnspire Cord; feet: First Sergeant's Dreadweave Boots; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Noxious Shooter

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers

### Band 60 (undead, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 427.6. Weights run: 1.8s. Verify run: 7.9s. 998 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.292 ± 0.025, crit=0.026 ± 0.002 per rating point (14 rating = 1%, 0.362 per %), hit=0.585 ± 0.040 per rating point (10 rating = 1%, 5.850 per %), spell_haste=not significant (-0.002 ± 0.013), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | sim-verified (427.6 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -13.37 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 31.8 spell_power points (8.85 DPS) | yes | Beads of Ogre Mojo (22149, -0.81 DPS, sim-verified) [quest]; Archlight Talisman (15856, -1.91 DPS) [quest]; Chains of the Lich (23125, -2.72 DPS) [dungeon] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | sim-verified (427.6 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -7.81 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.2 spell_power points (8.96 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Deep Woodlands Cloak (19121, -2.38 DPS) [quest]; Spritecaster Cape (11623, -2.90 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 49.2 spell_power points (13.70 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -12.71 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 32.3 spell_power points (9.00 DPS) | yes | General's Satin Bracers (17619, -2.17 DPS) [pvp]; Runecloth Cuffs (254123, -2.34 DPS) [crafted]; Virtuous Wraps (226953, -2.82 DPS) [vendor] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | sim-verified (427.6 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Raider Handwraps (272097, -8.68 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 59.3 spell_power points (16.49 DPS) | yes | Belt of the Archmage (18405, -4.03 DPS, sim-verified) [crafted]; Magister's Belt (16685, -7.00 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -7.89 DPS) [rep] |
| legs | Virtuous Leggings (226956) | Mokvar [vendor] | sim-verified (427.6 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Outrider's Silk Leggings (22747, -6.56 DPS, sim-verified) [rep] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 42.7 spell_power points (11.87 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Omnicast Boots (11822, -8.67 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.5 spell_power points (12.65 DPS) | yes | Deep Rooted Ring (19109, -5.72 DPS) [world]; Channeler's Ring (272406, -5.99 DPS) [vendor]; Band of Piety (22681, -7.11 DPS) [quest] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 36.2 spell_power points (10.08 DPS) | yes | Deep Rooted Ring (19109, -3.37 DPS, sim-verified) [world]; Channeler's Ring (272406, -3.42 DPS) [vendor]; Band of Piety (22681, -4.53 DPS) [quest] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (427.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (427.6 DPS) | yes | Weakness Analyzer (272438, -2.05 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.23 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.43 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (427.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -2.28 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (427.6 DPS) | yes | Stormrager (16997, -0.11 DPS) [quest]; Wand of Biting Cold (19108, -0.60 DPS) [quest]; Torch of Light (279246, -10.38 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Virtuous Hands; waist: Knowledge of the Timbermaw; legs: Virtuous Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 998, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers

### Band 60, raid preset (undead, 325001031300000000-00000000000000000-523120501201300251)

Set DPS (verified): 583.4. Weights run: 1.6s. Verify run: 5.1s. 998 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.014 ± 0.002, crit=0.062 ± 0.006 per rating point (14 rating = 1%, 0.864 per %), hit=0.459 ± 0.050 per rating point (10 rating = 1%, 4.590 per %), spell_haste=-7.784 ± 0.506, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 27.1 spell_power points (10.05 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -0.05 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.15 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.52 DPS) [quest]; Beads of Ogre Mojo (22149, -3.27 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.1 spell_power points (10.40 DPS) | yes | Warlord's Satin Epaulets (231611, -1.03 DPS) [vendor]; Argent Shoulders (19059, -1.14 DPS) [crafted]; Virtuous Epaulets (226955, -1.60 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.7 spell_power points (7.67 DPS) | yes | Hide of the Wild (18510, -2.43 DPS) [crafted]; Spritecaster Cape (11623, -2.45 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -2.48 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-verified (583.4 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Robe of the Magi (1716, -1.08 DPS) [world_drop]; Virtuous Gown (226960, -8.35 DPS, sim-verified) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.1 spell_power points (8.19 DPS) | yes | Runecloth Cuffs (254123, -4.06 DPS) [crafted]; Spidertank Oilrag (9448, -4.86 DPS) [dungeon]; Virtuous Wraps (226953, -5.17 DPS) [vendor] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.15 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; Virtuous Hands (226958, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 29.9 spell_power points (11.08 DPS) | yes | Frostwolf Cloth Belt (19090, -4.36 DPS) [rep]; Belt of the Archmage (18405, -5.20 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -5.54 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 33.9 spell_power points (12.55 DPS) | yes | General's Satin Leggings (231614, -0.58 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.61 DPS, sim-verified) [rep]; Ghostweave Pants (14144, -2.92 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.89 DPS) | yes | General's Satin Boots (17618, -0.27 DPS) [vendor]; General's Satin Treads (231610, -0.27 DPS) [vendor]; Virtuous Slippers (226959, -0.66 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 32.7 spell_power points (12.12 DPS) | yes | Blessed Band of Light (272407, -6.72 DPS) [vendor]; Advisor's Ring (19518, -6.93 DPS) [rep]; Band of the Unicorn (7553, -7.30 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.2 spell_power points (8.21 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Advisor's Ring (19518, -3.02 DPS) [rep]; Band of the Unicorn (7553, -3.39 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+17.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.21 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.96 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.66 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.68 DPS) [dungeon]; The Lobotomizer (19324, -30.70 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 213.2 spell_power points (78.98 DPS) | yes | Wand of Biting Cold (19108, +0.00 DPS, sim-verified) [quest]; Brilliant Wand (249385, -16.03 DPS) [crafted]; Stormrager (16997, -16.27 DPS) [quest] |

**New at 60:** head: Virtuous Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Truefaith Vestments; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Earthen Silk Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 998, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3223 Frostmane Scepter; 3224 Silver-lined Bracers

