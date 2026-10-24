# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 0000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 31.7. Weights run: 0.7s. Verify run: 0.5s. 145 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.365 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.008 per %), hit=0.238 ± 0.007 per rating point (10 rating = 1%, 2.381 per %), spell_haste=not significant (0.287 ± 0.111), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.63 DPS) | yes | Shadow Goggles (4373, -1.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.3 spell_power points (0.87 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.45 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.71 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.42 DPS) | yes | Pearl-clasped Cloak (5542, -0.10 DPS) [crafted]; Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Caretaker's Cape (20428, -0.11 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.72 DPS) | yes | Green Woolen Robe (6243, -0.29 DPS) [crafted]; Green Woolen Vest (2582, -0.30 DPS) [crafted]; Gray Woolen Robe (2585, -0.98 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.8 spell_power points (0.19 DPS) | yes | Windsong Bangles (263336, -0.09 DPS) [quest]; Repurposed Hair Band (281256, -0.12 DPS) [quest]; Bright Bracers (3647, -0.55 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.74 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Heavy Woolen Gloves (4310, -0.45 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.57 DPS) | yes | Novice Ardent's Sash (253887, -0.25 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.76 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.9 spell_power points (1.25 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.52 DPS) [dungeon]; Rumpled Kilt (274741, -0.73 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.89 DPS) | yes | Pristine Boots (253889, -0.37 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.47 DPS) [crafted]; Woolen Boots (2583, -0.60 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.60 DPS) | yes | Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.49 DPS) [world_drop]; The 1 Ring (8350, -0.56 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.53 DPS) | yes | Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.6 spell_power points (0.38 DPS) | yes | Witching Stave (1484, -0.19 DPS) [dungeon]; Staff of Westfall (2042, -0.19 DPS) [quest]; Lesser Staff of the Spire (1300, -0.28 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 215.1 spell_power points (22.63 DPS) | yes | Skycaller (12984, -1.01 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.10 DPS) [quest]; Sizzle Stick (8071, -4.45 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 145, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 0000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 61.7. Weights run: 0.7s. Verify run: 0.5s. 264 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.373 ± 0.010, crit=0.112 ± 0.004 per rating point (14 rating = 1%, 1.572 per %), hit=0.255 ± 0.010 per rating point (10 rating = 1%, 2.553 per %), spell_haste=0.694 ± 0.163, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.62 DPS) | yes | Silk Headband (7050, -0.29 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.44 DPS) [crafted]; Enchanter's Cowl (4322, -1.73 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.2 spell_power points (1.36 DPS) | yes | Darkspear Warding Pendant (272075, -1.09 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.14 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.29 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 spell_power points (1.82 DPS) | yes | Death Speaker Mantle (6685, -0.33 DPS) [dungeon]; Fairywing Mantle (9536, -0.44 DPS) [quest]; Invoker's Mantle (215365, -0.51 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.74 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Prelacy Cape (7004, -0.15 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.8 spell_power points (2.04 DPS) | yes | Beguiler Robes (7728, +0.00 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.12 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.20 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.33 DPS) | yes | Nightsky Wristbands (6407, -1.00 DPS) [world_drop]; Stonecloth Bindings (14416, -1.05 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.32 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.03 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Truefaith Gloves (7049, -0.13 DPS) [crafted]; Gnoll Casting Gloves (892, -0.15 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.1 spell_power points (1.79 DPS) | yes | Belt of Arugal (6392, -0.29 DPS) [dungeon]; Invoker's Cord (215366, -0.48 DPS) [crafted]; Crimson Silk Belt (7055, -0.52 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.77 DPS) | yes | Abomination Skin Leggings (23173, -0.00 DPS) [dungeon]; Pristine Leggings (253987, -0.35 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.56 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.6 spell_power points (1.42 DPS) | yes | Acidic Walkers (9454, -0.24 DPS) [dungeon]; Glinteye Slippers (273024, -0.38 DPS) [dungeon]; Spidersilk Boots (4320, -2.12 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.9 spell_power points (1.16 DPS) | yes | Ring of Power Regulation (273030, -0.27 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.27 DPS) [vendor]; Minor Channeling Ring (1449, -0.31 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.03 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.15 DPS) [vendor]; Minor Channeling Ring (1449, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (5.31 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -2.36 DPS) [quest]; Glimmering Staff (249392, -4.70 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 224.3 spell_power points (33.08 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.34 DPS) [crafted]; Gravestone Scepter (7001, -4.08 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 0000000000000000000-00000000000000000-253510130100300250)

Set DPS (verified): 87.8. Weights run: 0.8s. Verify run: 0.7s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.582 ± 0.021, crit=0.168 ± 0.007 per rating point (14 rating = 1%, 2.355 per %), hit=0.420 ± 0.023 per rating point (10 rating = 1%, 4.199 per %), spell_haste=1.719 ± 0.334, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.16 DPS) [world]; Living Cowl (5608, -0.71 DPS) [world]; Spellpower Goggles Xtreme (10502, -1.64 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (1.69 DPS) | yes | Triune Amulet (7722, -0.53 DPS) [dungeon]; Scorn's Icy Choker (23169, -0.90 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.10 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (2.11 DPS) | yes | Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Berylline Pads (4197, -0.25 DPS) [quest]; Green Silken Shoulders (7057, -0.84 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.9 spell_power points (1.29 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.28 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.36 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (3.69 DPS) | yes | Dreamweave Vest (10021, -0.33 DPS) [crafted]; Elemental Raiment (9434, -0.65 DPS) [world_drop]; Robe of Power (7054, -0.65 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.30 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest]; Windchaser Cuffs (14429, -0.54 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (2.94 DPS) | yes | Red Mageweave Gloves (10018, -0.51 DPS) [crafted]; Black Mageweave Gloves (10003, -0.77 DPS) [crafted]; Gilded Handwraps (254021, -1.20 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.40 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.43 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (3.04 DPS) | yes | Crimson Silk Pantaloons (7062, -0.78 DPS) [crafted]; Abomination Skin Leggings (23173, -1.06 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.30 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.48 DPS) | yes | Gilded Slippers (254001, -1.09 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.08 DPS) [dungeon]; Spidersilk Boots (4320, -2.12 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (1.87 DPS) | yes | Snake Hoop (6750, -0.58 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.71 DPS) [quest]; Reedknot Ring (9622, -0.86 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.30 DPS) | yes | Snake Hoop (6750, -0.01 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.14 DPS) [quest]; Reedknot Ring (9622, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.17 DPS) [dungeon]; Spine of the Devourer (273040, -1.83 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.7 spell_power points (2.14 DPS) | yes | Orb of the Forgotten Seer (7685, -0.11 DPS) [dungeon]; Orb of Lorica (11262, -0.42 DPS) [quest]; Orb of Mystic Insight (249394, -0.62 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 275.4 spell_power points (39.89 DPS) | yes | Nether Force Wand (11263, -2.30 DPS) [quest]; Icefury Wand (7514, -2.44 DPS) [quest]; Ragefire Wand (7513, -2.49 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 2030005000000000000-00000000000000000-253510130100300250)

Set DPS (verified): 213.8. Weights run: 0.8s. Verify run: 0.8s. 436 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.822 ± 0.034, crit=0.272 ± 0.010 per rating point (14 rating = 1%, 3.805 per %), hit=0.828 ± 0.045 per rating point (10 rating = 1%, 8.277 per %), spell_haste=6.671 ± 0.606, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.4 spell_power points (7.35 DPS) | yes | Dreamweave Circlet (10041, -1.29 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.61 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.75 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 spell_power points (2.82 DPS) | yes | Scorn's Icy Choker (23169, -0.34 DPS) [dungeon]; Horizon Choker (13085, -0.43 DPS) [world_drop]; Mindburst Medallion (11196, -0.55 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.8 spell_power points (5.77 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -1.78 DPS) [vendor]; Inquisitor's Shawl (19507, -2.10 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.69 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.9 spell_power points (3.93 DPS) | yes | Runecloth Cloak (13860, -0.70 DPS) [crafted]; Big Voodoo Cloak (8216, -1.36 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.54 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.4 spell_power points (7.35 DPS) | yes | Robe of the Magi (1716, -1.76 DPS) [world_drop]; Runecloth Tunic (13857, -1.95 DPS) [crafted]; Dreamweave Vest (10021, -2.08 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.8 spell_power points (2.65 DPS) | yes | Forgotten Wraps (9433, -0.60 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.60 DPS) [quest]; Bloodband Bracers (11469, -2.73 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.8 spell_power points (6.59 DPS) | yes | Raider Handwraps (272098, -1.39 DPS) [vendor]; Dreamweave Gloves (10019, -2.18 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.36 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.2 spell_power points (4.61 DPS) | yes | Dawnspire Cord (12466, -0.13 DPS) [dungeon]; Deathmage Sash (10771, -0.60 DPS) [dungeon]; Runecloth Belt (13856, -0.90 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 25.7 spell_power points (5.33 DPS) | yes | Red Mageweave Pants (10009, -0.37 DPS) [crafted]; Wizardweave Leggings (14132, -1.38 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.45 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.98 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.07 DPS) [vendor]; Gilded Sandals (254107, -1.16 DPS) [crafted]; Southsea Mojo Boots (20641, -1.44 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.93 DPS) | yes | Lorekeeper's Ring (19523, -0.44 DPS) [rep]; Snake Hoop (6750, -0.83 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.70 DPS) | yes | Lorekeeper's Ring (19523, -0.21 DPS) [rep]; Snake Hoop (6750, -0.60 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (213.8 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (213.8 DPS) | yes | Mark of the Chosen (17774, -3.62 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (213.8 DPS) | yes | Spellshifter Rod (9527, -0.19 DPS) [quest]; Kindling Stave (11750, -0.82 DPS) [dungeon]; Hanzo Sword (8190, -35.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 253.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.46 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.09 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Knight's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 2030005000000000000-11302300000000000-253510130100300250)

Set DPS (verified): 324.0. Weights run: 0.8s. Verify run: 2.3s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.595 ± 0.038, crit=0.369 ± 0.012 per rating point (14 rating = 1%, 5.166 per %), hit=1.037 ± 0.051 per rating point (10 rating = 1%, 10.373 per %), spell_haste=7.127 ± 0.624, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 41.0 spell_power points (8.93 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.91 DPS) [pvp]; Red Mageweave Headband (10033, -9.00 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (324.0 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.16 DPS) [quest]; Chains of the Lich (23125, -0.16 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 41.1 spell_power points (8.94 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.56 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.84 DPS) [vendor]; Mantle of the Timbermaw (19050, -5.34 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.1 spell_power points (6.78 DPS) | yes | Hide of the Wild (18510, -2.43 DPS) [crafted]; Spritecaster Cape (11623, -2.95 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -3.73 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 52.3 spell_power points (11.39 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.88 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.49 DPS) [pvp]; Acumen Robes (17775, -8.29 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.8 spell_power points (5.83 DPS) | yes | Runecloth Cuffs (254123, -2.14 DPS) [crafted]; Sorcerer's Bindings (226929, -3.01 DPS) [quest]; Bracers of Hope (22667, -3.09 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (324.0 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, -3.39 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 48.5 spell_power points (10.55 DPS) | yes | Magician's Cord (272393, -3.30 DPS) [vendor]; Stormpike Cloth Girdle (19094, -5.34 DPS) [rep]; Belt of the Archmage (18405, -7.00 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 48.3 spell_power points (10.51 DPS) | yes | Marshal's Silk Leggings (231605, -0.26 DPS) [pvp]; Sorcerer's Leggings (226933, -2.04 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -2.61 DPS) [pvp] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (324.0 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.74 DPS) [dungeon]; Sorcerer's Boots (22064, -3.57 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 43.7 spell_power points (9.52 DPS) | yes | Maiden's Circle (13001, -4.69 DPS) [world_drop]; Blessed Band of Light (272407, -5.08 DPS) [vendor]; Channeler's Ring (272406, -5.45 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 28.5 spell_power points (6.21 DPS) | yes | Blessed Band of Light (272407, -1.78 DPS) [vendor]; Channeler's Ring (272406, -2.14 DPS) [vendor]; Maiden's Circle (13001, -3.60 DPS, sim-verified) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (324.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (324.0 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.02 DPS) [crafted]; Uther's Strength (11302, -1.74 DPS) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (324.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -0.09 DPS) [world_drop]; Teebu's Blazing Longsword (1728, -29.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 358.6 spell_power points (78.07 DPS) | yes | Wand of Biting Cold (19108, -5.88 DPS, sim-verified) [quest]; Stormrager (16997, -14.73 DPS) [quest]; Brilliant Wand (249385, -15.62 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spellshifter Rod; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 2030005000000000000-13102000000000000-253510130101301251)

Set DPS (verified): 505.9. Weights run: 0.8s. Verify run: 1.7s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.178 ± 0.014, crit=0.418 ± 0.013 per rating point (14 rating = 1%, 5.857 per %), hit=1.012 ± 0.040 per rating point (10 rating = 1%, 10.123 per %), spell_haste=6.995 ± 0.276, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.3 spell_power points (11.46 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.46 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.33 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (505.9 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.72 DPS) [quest]; Jewel of Kajaro (19601, -4.25 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.5 spell_power points (13.01 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.88 DPS) [pvp]; Argent Shoulders (19059, -3.85 DPS) [crafted]; Mantle of the Timbermaw (19050, -4.05 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.5 spell_power points (10.08 DPS) | yes | Hide of the Wild (18510, -2.77 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -4.57 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -4.96 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.0 spell_power points (17.57 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.24 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.63 DPS) [pvp]; Robe of the Magi (1716, -11.42 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.4 spell_power points (8.58 DPS) | yes | Runecloth Cuffs (254123, -3.90 DPS) [crafted]; Arcane Runed Bracers (4744, -5.28 DPS) [quest]; Spidertank Oilrag (9448, -5.28 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.9 spell_power points (10.21 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.0 spell_power points (14.29 DPS) | yes | Magician's Cord (272393, -5.77 DPS) [vendor]; Belt of the Archmage (18405, -5.95 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -6.63 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.5 spell_power points (16.66 DPS) | yes | Marshal's Silk Leggings (231605, -2.22 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -5.71 DPS) [pvp]; Sorcerer's Leggings (226933, -6.28 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.79 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.06 DPS) [quest]; Sorcerer's Sandals (226931, -0.06 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 39.7 spell_power points (14.54 DPS) | yes | Blessed Band of Light (272407, -7.18 DPS) [vendor]; Maiden's Circle (13001, -7.50 DPS) [world_drop]; Don Julio's Band (19325, -8.69 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 24.0 spell_power points (8.77 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Maiden's Circle (13001, -1.73 DPS) [world_drop]; Don Julio's Band (19325, -2.92 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (505.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (505.9 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.79 DPS) [crafted]; Uther's Strength (11302, -2.93 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (505.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.05 DPS) [world]; Teebu's Blazing Longsword (1728, -33.31 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 215.7 spell_power points (78.96 DPS) | yes | Wand of Biting Cold (19108, -4.76 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.63 DPS) [crafted]; Stormrager (16997, -15.94 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 0000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 29.4. Weights run: 0.7s. Verify run: 0.5s. 134 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.365 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.008 per %), hit=0.238 ± 0.007 per rating point (10 rating = 1%, 2.381 per %), spell_haste=not significant (0.287 ± 0.111), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.63 DPS) | yes | Shadow Goggles (4373, -1.11 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.3 spell_power points (0.87 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.45 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.60 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.42 DPS) | yes | Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Battle Healer's Cloak (20427, -0.11 DPS) [rep]; Pearl-clasped Cloak (5542, -0.28 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.72 DPS) | yes | Green Woolen Robe (6243, -0.29 DPS) [crafted]; Green Woolen Vest (2582, -0.30 DPS) [crafted]; Gray Woolen Robe (2585, -0.91 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.32 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Featherbead Bracers (15452, -0.12 DPS) [quest]; Owlbeard Bracers (16981, -0.13 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.74 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.21 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.32 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.57 DPS) | yes | Novice Ardent's Sash (253887, -0.25 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.9 spell_power points (1.25 DPS) | yes | Filigreed Pristine Leggings (253937, -0.39 DPS) [crafted]; Silk-threaded Trousers (1929, -0.52 DPS) [dungeon]; Rumpled Kilt (274741, -0.73 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.89 DPS) | yes | Pristine Boots (253889, -0.36 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.47 DPS) [crafted]; Woolen Boots (2583, -0.60 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.53 DPS) | yes | Loop of Sacrifice (281673, -0.33 DPS) [quest]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; The 1 Ring (8350, -0.49 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.2 spell_power points (0.23 DPS) | yes | Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; The 1 Ring (8350, -0.19 DPS) [world]; Loop of Sacrifice (281673, -0.25 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.6 spell_power points (0.38 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.15 DPS) [world]; Clear Crystal Rod (16894, -0.19 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 215.1 spell_power points (22.63 DPS) | yes | Skycaller (12984, -1.16 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.45 DPS) [quest]; Greater Magic Wand (11288, -4.92 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 134, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (troll, 0000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 59.2. Weights run: 0.7s. Verify run: 0.6s. 245 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.373 ± 0.010, crit=0.112 ± 0.004 per rating point (14 rating = 1%, 1.572 per %), hit=0.255 ± 0.010 per rating point (10 rating = 1%, 2.553 per %), spell_haste=0.694 ± 0.163, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.62 DPS) | yes | Silk Headband (7050, -0.29 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.44 DPS) [crafted]; Enchanter's Cowl (4322, -1.58 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.2 spell_power points (1.36 DPS) | yes | Darkspear Warding Pendant (272075, -1.09 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.14 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.20 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 spell_power points (1.82 DPS) | yes | Death Speaker Mantle (6685, -0.33 DPS) [dungeon]; Fairywing Mantle (9536, -0.44 DPS) [quest]; Invoker's Mantle (215365, -0.51 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.74 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.8 spell_power points (2.04 DPS) | yes | Beguiler Robes (7728, -0.05 DPS) [dungeon]; Tree Bark Jacket (1486, -0.12 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.20 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.33 DPS) | yes | Glowing Magical Bracelets (13106, -0.89 DPS) [world_drop]; Nightsky Wristbands (6407, -1.00 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.81 DPS, sim-verified) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.9 spell_power points (1.16 DPS) | yes | Truefaith Gloves (7049, -0.26 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS) [world]; Serpent Gloves (5970, -0.52 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.1 spell_power points (1.79 DPS) | yes | Warsong Sash (16975, -0.16 DPS) [quest]; Belt of Arugal (6392, -0.29 DPS) [dungeon]; Invoker's Cord (215366, -0.48 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.77 DPS) | yes | Abomination Skin Leggings (23173, -0.00 DPS) [dungeon]; Pristine Leggings (253987, -0.35 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.56 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.6 spell_power points (1.42 DPS) | yes | Acidic Walkers (9454, -0.24 DPS) [dungeon]; Glinteye Slippers (273024, -0.38 DPS) [dungeon]; Spidersilk Boots (4320, -1.60 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.9 spell_power points (1.16 DPS) | yes | Ring of Power Regulation (273030, -0.27 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.27 DPS) [vendor]; Solid Gold Band (279845, -0.57 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.03 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.15 DPS) [vendor]; Solid Gold Band (279845, -0.44 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (5.31 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -4.70 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -4.76 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | sim-verified (59.2 DPS) | yes | Starfaller (13063, -0.59 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -1.66 DPS) [crafted]; Unstable Power Core (279847, -1.69 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 245, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (troll, 0000000000000000000-00000000000000000-253510130100300250)

Set DPS (verified): 84.3. Weights run: 0.8s. Verify run: 0.7s. 320 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.582 ± 0.021, crit=0.168 ± 0.007 per rating point (14 rating = 1%, 2.355 per %), hit=0.420 ± 0.023 per rating point (10 rating = 1%, 4.199 per %), spell_haste=1.719 ± 0.334, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (84.3 DPS) | yes | Augural Shroud (2620, -0.16 DPS) [world]; Living Cowl (5608, -0.71 DPS) [world]; Spellpower Goggles Xtreme (10502, -1.52 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (1.69 DPS) | yes | Triune Amulet (7722, -0.53 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.10 DPS) [vendor]; Scorn's Icy Choker (23169, -1.82 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (2.11 DPS) | yes | Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Berylline Pads (4197, -0.25 DPS) [quest]; Green Silken Shoulders (7057, -1.76 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.9 spell_power points (1.29 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.28 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.36 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (3.69 DPS) | yes | Dreamweave Vest (10021, -0.33 DPS) [crafted]; Elemental Raiment (9434, -0.65 DPS) [world_drop]; Robe of Power (7054, -0.65 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.30 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.05 DPS) [quest]; Condor Bracers (15864, -0.29 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (2.94 DPS) | yes | Black Mageweave Gloves (10003, -0.77 DPS) [crafted]; Red Mageweave Gloves (10018, -0.91 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.20 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.3 spell_power points (2.36 DPS) | yes | Deathmage Sash (10771, -0.09 DPS) [dungeon]; Star Belt (4329, -0.48 DPS) [crafted]; Gilded Cord (254037, -0.53 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (3.04 DPS) | yes | Abomination Skin Leggings (23173, -1.06 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.13 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.30 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.48 DPS) | yes | Gilded Slippers (254001, -1.03 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.08 DPS) [dungeon]; Spidersilk Boots (4320, -2.12 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (1.87 DPS) | yes | Snake Hoop (6750, -0.58 DPS) [quest]; Reedknot Ring (9622, -0.86 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.00 DPS) [vendor] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.30 DPS) | yes | Reedknot Ring (9622, -0.29 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.43 DPS) [vendor]; Snake Hoop (6750, -0.68 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.17 DPS) [dungeon]; Spine of the Devourer (273040, -1.83 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.7 spell_power points (2.14 DPS) | yes | Orb of the Forgotten Seer (7685, -0.11 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.62 DPS) [crafted]; Omega Orb (7749, -0.83 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 275.4 spell_power points (39.89 DPS) | yes | Nether Force Wand (11263, -1.15 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.44 DPS) [quest]; Ragefire Wand (7513, -2.49 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 320, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (troll, 2030005000000000000-00000000000000000-253510130100300250)

Set DPS (verified): 205.9. Weights run: 0.8s. Verify run: 0.8s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.822 ± 0.034, crit=0.272 ± 0.010 per rating point (14 rating = 1%, 3.805 per %), hit=0.828 ± 0.045 per rating point (10 rating = 1%, 8.277 per %), spell_haste=6.671 ± 0.606, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.4 spell_power points (7.35 DPS) | yes | Dreamweave Circlet (10041, -1.29 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.61 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.75 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 spell_power points (2.82 DPS) | yes | Scorn's Icy Choker (23169, -0.34 DPS) [dungeon]; Horizon Choker (13085, -0.43 DPS) [world_drop]; Mindburst Medallion (11196, -0.55 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.8 spell_power points (5.77 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.78 DPS) [vendor]; Inquisitor's Shawl (19507, -2.10 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.61 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.4 spell_power points (4.02 DPS) | yes | Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.79 DPS) [crafted]; Big Voodoo Cloak (8216, -1.45 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.4 spell_power points (7.35 DPS) | yes | Robe of the Magi (1716, -1.76 DPS) [world_drop]; Runecloth Tunic (13857, -1.95 DPS) [crafted]; Dreamweave Vest (10021, -2.08 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.8 spell_power points (2.65 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.07 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.8 spell_power points (6.59 DPS) | yes | Raider Handwraps (272098, -1.39 DPS) [vendor]; Dreamweave Gloves (10019, -2.18 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.36 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.2 spell_power points (4.61 DPS) | yes | Dawnspire Cord (12466, -0.13 DPS) [dungeon]; Deathmage Sash (10771, -0.60 DPS) [dungeon]; Runecloth Belt (13856, -0.90 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 25.7 spell_power points (5.33 DPS) | yes | Red Mageweave Pants (10009, -0.37 DPS) [crafted]; Wizardweave Leggings (14132, -1.38 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.45 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.98 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.07 DPS) [vendor]; Gilded Sandals (254107, -1.16 DPS) [crafted]; Southsea Mojo Boots (20641, -1.44 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.93 DPS) | yes | Advisor's Ring (19519, -0.44 DPS) [rep]; Snake Hoop (6750, -0.83 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.70 DPS) | yes | Advisor's Ring (19519, -0.21 DPS) [rep]; Snake Hoop (6750, -0.60 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (205.9 DPS) | yes | Uther's Strength (11302, -0.30 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (205.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (205.9 DPS) | yes | Spellshifter Rod (9527, -0.19 DPS) [quest]; Kindling Stave (11750, -0.82 DPS) [dungeon]; Hanzo Sword (8190, -35.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 253.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.46 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.09 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Stone Guard's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (troll, 2030005000000000000-11302300000000000-253510130100300250)

Set DPS (verified): 313.2. Weights run: 0.8s. Verify run: 2.4s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.595 ± 0.038, crit=0.369 ± 0.012 per rating point (14 rating = 1%, 5.166 per %), hit=1.037 ± 0.051 per rating point (10 rating = 1%, 10.373 per %), spell_haste=7.127 ± 0.624, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 41.0 spell_power points (8.93 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.91 DPS) [pvp]; Red Mageweave Headband (10033, -6.35 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (313.2 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.16 DPS) [quest]; Chains of the Lich (23125, -0.16 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 41.1 spell_power points (8.94 DPS) | yes | Warlord's Silk Amice (231594, -1.56 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.44 DPS) [crafted]; Darkspear Shoulderpads (272103, -2.84 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.1 spell_power points (6.78 DPS) | yes | Hide of the Wild (18510, -2.43 DPS) [crafted]; Spritecaster Cape (11623, -2.95 DPS) [dungeon]; Deep Woodlands Cloak (19121, -3.00 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 52.3 spell_power points (11.39 DPS) | yes | Warlord's Silk Raiment (231596, -0.88 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.49 DPS) [pvp]; Acumen Robes (17775, -7.09 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.8 spell_power points (5.83 DPS) | yes | Runecloth Cuffs (254123, -2.14 DPS) [crafted]; Sorcerer's Bindings (226929, -3.01 DPS) [quest]; Bracers of Hope (22667, -3.09 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (313.2 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, -4.01 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 48.5 spell_power points (10.55 DPS) | yes | Magician's Cord (272393, -3.30 DPS) [vendor]; Belt of the Archmage (18405, -5.31 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -5.34 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 48.3 spell_power points (10.51 DPS) | yes | General's Silk Trousers (231595, -0.26 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.95 DPS) [rep]; Sorcerer's Leggings (226933, -2.04 DPS) [quest] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (313.2 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.74 DPS) [dungeon]; Sorcerer's Boots (22064, -4.00 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 43.7 spell_power points (9.52 DPS) | yes | Maiden's Circle (13001, -4.69 DPS) [world_drop]; Blessed Band of Light (272407, -5.08 DPS) [vendor]; Channeler's Ring (272406, -5.45 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 28.5 spell_power points (6.21 DPS) | yes | Maiden's Circle (13001, -1.39 DPS) [world_drop]; Blessed Band of Light (272407, -1.78 DPS) [vendor]; Channeler's Ring (272406, -2.14 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (313.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (313.2 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.02 DPS) [crafted]; Rune of the Guard Captain (19120, -1.47 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (313.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -0.09 DPS) [world_drop]; Teebu's Blazing Longsword (1728, -27.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 358.6 spell_power points (78.07 DPS) | yes | Wand of Biting Cold (19108, -5.20 DPS, sim-verified) [quest]; Stormrager (16997, -14.73 DPS) [quest]; Brilliant Wand (249385, -15.62 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spellshifter Rod; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (troll, 2030005000000000000-13102000000000000-253510130101301251)

Set DPS (verified): 500.2. Weights run: 0.8s. Verify run: 1.6s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.178 ± 0.014, crit=0.418 ± 0.013 per rating point (14 rating = 1%, 5.857 per %), hit=1.012 ± 0.040 per rating point (10 rating = 1%, 10.123 per %), spell_haste=6.995 ± 0.276, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.3 spell_power points (11.46 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.46 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.97 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (500.2 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.72 DPS) [quest]; Jewel of Kajaro (19601, -4.65 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.5 spell_power points (13.01 DPS) | yes | Warlord's Silk Amice (231594, -2.88 DPS) [pvp]; Argent Shoulders (19059, -3.85 DPS) [crafted]; Mantle of the Timbermaw (19050, -4.51 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.5 spell_power points (10.08 DPS) | yes | Hide of the Wild (18510, -4.31 DPS) [crafted]; Spritecaster Cape (11623, -4.57 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -4.96 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.0 spell_power points (17.57 DPS) | yes | Warlord's Silk Raiment (231596, -2.24 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.63 DPS) [pvp]; Robe of the Magi (1716, -12.40 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.4 spell_power points (8.58 DPS) | yes | Runecloth Cuffs (254123, -3.90 DPS) [crafted]; Spidertank Oilrag (9448, -5.28 DPS) [dungeon]; Sorcerer's Bindings (226929, -5.36 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.9 spell_power points (10.21 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.0 spell_power points (14.29 DPS) | yes | Magician's Cord (272393, -5.77 DPS) [vendor]; Defiler's Cloth Girdle (20163, -6.63 DPS) [rep]; Belt of the Archmage (18405, -6.75 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.5 spell_power points (16.66 DPS) | yes | General's Silk Trousers (231595, -2.22 DPS) [pvp]; Outrider's Silk Leggings (22747, -5.17 DPS) [rep]; Sorcerer's Leggings (226933, -6.87 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.79 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.06 DPS) [quest]; Sorcerer's Sandals (226931, -0.06 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 39.7 spell_power points (14.54 DPS) | yes | Blessed Band of Light (272407, -7.18 DPS) [vendor]; Maiden's Circle (13001, -7.50 DPS) [world_drop]; Don Julio's Band (19325, -8.69 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 24.0 spell_power points (8.77 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Maiden's Circle (13001, -1.73 DPS) [world_drop]; Don Julio's Band (19325, -2.92 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (500.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -5.89 DPS, sim-verified) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (500.2 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Rune of the Guard Captain (19120, -0.74 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (500.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.05 DPS) [world]; Teebu's Blazing Longsword (1728, -33.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 215.7 spell_power points (78.96 DPS) | yes | Wand of Biting Cold (19108, -4.85 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.63 DPS) [crafted]; Stormrager (16997, -15.94 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

