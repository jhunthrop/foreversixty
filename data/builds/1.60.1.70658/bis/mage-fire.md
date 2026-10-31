# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 0000000000000000000-23510000000000000-000000000000000000)

Set DPS (verified): 27.6. Weights run: 1.1s. Verify run: 0.8s. 154 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.413 ± 0.011, crit=0.127 ± 0.004 per rating point (14 rating = 1%, 1.778 per %), hit=0.389 ± 0.013 per rating point (10 rating = 1%, 3.895 per %), spell_haste=not significant (0.525 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.71 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Shadow Goggles (4373, -0.58 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.7 spell_power points (0.56 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.18 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.31 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Pearl-clasped Cloak (5542, -0.05 DPS) [crafted]; Black Whelp Cloak (7283, -0.06 DPS) [crafted]; Sanguine Cape (14376, -0.15 DPS) [world_drop] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 11.1 spell_power points (0.72 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.32 DPS) [crafted]; Green Woolen Robe (6243, -0.44 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.1 spell_power points (0.13 DPS) | yes | Windsong Bangles (263336, -0.07 DPS) [quest]; Repurposed Hair Band (281256, -0.08 DPS) [quest]; Bright Bracers (3647, -0.34 DPS, sim-verified) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.58 DPS) | yes | Flame Gloves (253917, -0.05 DPS) [crafted]; Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.7 spell_power points (0.56 DPS) | yes | Pristine Sash (253925, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.35 DPS) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.5 spell_power points (0.87 DPS) | yes | Phoenix Pants (4317, -0.28 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted]; Silk-threaded Trousers (1929, -0.42 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 spell_power points (0.56 DPS) | yes | Flame Boots (253893, -0.09 DPS) [crafted]; Pristine Boots (253889, -0.29 DPS) [crafted]; Red Woolen Boots (4313, -0.30 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.38 DPS) | yes | Volcanic Rock Ring (12053, -0.30 DPS) [world_drop]; The 1 Ring (8350, -0.35 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.5 spell_power points (0.16 DPS) | yes | Volcanic Rock Ring (12053, -0.08 DPS) [world_drop]; The 1 Ring (8350, -0.13 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 4.1 spell_power points (0.27 DPS) | yes | Lesser Staff of the Spire (1300, -0.11 DPS) [world]; Witching Stave (1484, -0.13 DPS) [dungeon]; Staff of Westfall (2042, -0.13 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 347.7 spell_power points (22.50 DPS) | yes | Skycaller (12984, -0.57 DPS, sim-verified) [world_drop]; Deepblaze (279896, -3.97 DPS) [quest]; Sizzle Stick (8071, -4.53 DPS) [quest] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 154, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 0000000000000000000-23552100120000000-000000000000000000)

Set DPS (verified): 48.8. Weights run: 1.2s. Verify run: 0.8s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.700 ± 0.013, crit=0.172 ± 0.007 per rating point (14 rating = 1%, 2.406 per %), hit=0.360 ± 0.017 per rating point (10 rating = 1%, 3.602 per %), spell_haste=0.949 ± 0.187, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.32 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Holy Shroud (2721, -0.28 DPS) [world_drop]; Flame Circlet (253953, -0.28 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.2 spell_power points (0.40 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (1.44 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.47 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Resilient Cape (14400, -0.08 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.08 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.1 spell_power points (1.71 DPS) | yes | Flame Gown (253965, -0.02 DPS) [crafted]; Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.23 DPS) | yes | Spidertank Oilrag (9448, -0.38 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.70 DPS) [world_drop]; Nightsky Wristbands (6407, -0.83 DPS) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.1 spell_power points (0.86 DPS) | yes | Phoenix Gloves (4331, -0.01 DPS) [crafted]; Town Clerk's Mittens (270029, -0.11 DPS) [quest]; Truefaith Gloves (7049, -0.19 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.1 spell_power points (1.05 DPS) | yes | Crimson Silk Belt (7055, -0.02 DPS) [crafted]; Invoker's Cord (215366, -0.06 DPS) [crafted]; Flame Sash (253929, -0.12 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.9 spell_power points (1.78 DPS) | yes | Filigreed Flame Leggings (253941, -0.54 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.56 DPS) [world]; Phoenix Pants (4317, -0.58 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.9 spell_power points (1.69 DPS) | yes | Gilded Slippers (254001, -0.57 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.90 DPS) | yes | Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon]; Azora's Will (4999, -0.57 DPS) [world_drop]; Minor Channeling Ring (1449, -1.32 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (48.8 DPS) | yes | Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; Azora's Will (4999, -0.24 DPS) [world_drop]; Minor Channeling Ring (1449, -0.52 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.39 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.51 DPS) [quest]; Glimmering Staff (249392, -2.67 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 351.3 spell_power points (33.12 DPS) | yes | Wand of Mana Concentration (273043, -0.40 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.65 DPS) [crafted]; Scorching Wand (5213, -3.71 DPS) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Flame Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 0000000000000000000-23552100130125010-000000000000000000)

Set DPS (verified): 82.2. Weights run: 1.4s. Verify run: 1.0s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.916 ± 0.031, crit=0.259 ± 0.020 per rating point (14 rating = 1%, 3.630 per %), hit=0.489 ± 0.040 per rating point (10 rating = 1%, 4.886 per %), spell_haste=2.140 ± 0.521, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.6 spell_power points (2.72 DPS) | yes | Augural Shroud (2620, -0.39 DPS) [world]; Whitemane's Chapeau (7720, -0.85 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.78 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.3 spell_power points (1.65 DPS) | yes | Necklace of Calisea (1714, -0.91 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.02 DPS) [world]; Triune Amulet (7722, -1.26 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.9 spell_power points (2.18 DPS) | yes | Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest]; Green Silken Shoulders (7057, -1.00 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.6 spell_power points (1.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.37 DPS) [dungeon]; Icy Cloak (4327, -0.41 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.5 spell_power points (3.17 DPS) | yes | Cindercloth Robe (10042, -0.06 DPS) [crafted]; Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.29 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.50 DPS) | yes | Arcane Runed Bracers (4744, -0.46 DPS) [quest]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Windchaser Cuffs (14429, -0.55 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.7 spell_power points (2.50 DPS) | yes | Fiery Handwraps (254025, -0.14 DPS) [crafted]; Red Mageweave Gloves (10018, -0.17 DPS) [crafted]; Crimson Silk Gloves (7064, -0.25 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 21.3 spell_power points (2.46 DPS) | yes | Deathmage Sash (10771, -0.07 DPS) [dungeon]; Gilded Cord (254037, -0.69 DPS) [crafted]; Scorching Sash (4117, -0.96 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.0 spell_power points (2.88 DPS) | yes | Flame Leggings (253991, -0.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.59 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.98 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Fiery Slippers (254005, -0.53 DPS) [crafted]; Gilded Slippers (254001, -1.22 DPS) [crafted]; Acidic Walkers (9454, -1.35 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.68 DPS) | yes | Ring of Forlorn Spirits (2043, -0.76 DPS) [quest]; Reedknot Ring (9622, -0.87 DPS) [quest]; Minor Channeling Ring (1449, -0.89 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.6 spell_power points (1.22 DPS) | yes | Ring of Forlorn Spirits (2043, -0.30 DPS) [quest]; Reedknot Ring (9622, -0.41 DPS) [quest]; Minor Channeling Ring (1449, -0.43 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.21 DPS) [dungeon]; Searing Golden Blade (12260, -0.68 DPS) [crafted] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (82.2 DPS) | yes | Orb of the Forgotten Seer (7685, -0.13 DPS) [dungeon]; Rod of Molten Fire (2565, -0.25 DPS) [world_drop]; Celestial Orb (7515, -1.11 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 347.1 spell_power points (40.01 DPS) | yes | Ragefire Wand (7513, -1.52 DPS) [quest]; Nether Force Wand (11263, -2.37 DPS) [quest]; Icefury Wand (7514, -2.51 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 2050000000000000000-23552100130125031-000000000000000000)

Set DPS (verified): 147.2. Weights run: 1.3s. Verify run: 1.1s. 506 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.886 ± 0.034, crit=0.303 ± 0.027 per rating point (14 rating = 1%, 4.237 per %), hit=0.558 ± 0.056 per rating point (10 rating = 1%, 5.585 per %), spell_haste=not significant (0.439 ± 0.642), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.7 spell_power points (4.90 DPS) | yes | Knight-Lieutenant's Dreadweave Hat (220889, -1.05 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.30 DPS) [crafted]; Dreamweave Circlet (10041, -1.75 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.1 spell_power points (1.88 DPS) | yes | Horizon Choker (13085, -0.22 DPS) [world_drop]; Mindburst Medallion (11196, -0.37 DPS) [quest]; Triune Amulet (7722, -0.81 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.0 spell_power points (3.86 DPS) | yes | Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.17 DPS) [vendor]; Netherflame Shoulders (254053, -1.62 DPS, sim-verified) [crafted] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.1 spell_power points (2.68 DPS) | yes | Runecloth Cloak (13860, -0.53 DPS) [crafted]; Cloak of Fire (14134, -0.81 DPS) [crafted]; Spritecaster Cape (11623, -1.59 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.7 spell_power points (4.90 DPS) | yes | Cindercloth Robe (10042, -1.30 DPS) [crafted]; Runecloth Tunic (13857, -1.33 DPS) [crafted]; Robe of the Magi (1716, -2.66 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.2 spell_power points (2.56 DPS) | yes | Nethergeld Cuffs (254061, -0.80 DPS) [crafted]; Phoenix Bindings (210781, -0.83 DPS) [crafted]; Bloodband Bracers (11469, -0.83 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 27.9 spell_power points (3.72 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -0.92 DPS) [vendor]; Fiery Handwraps (254025, -1.02 DPS) [crafted]; Dreamweave Gloves (10019, -1.68 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.0 spell_power points (3.60 DPS) | yes | Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Dawnspire Cord (12466, -0.55 DPS) [dungeon]; Fiery Cord (254041, -0.79 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 26.9 spell_power points (3.58 DPS) | yes | Red Mageweave Pants (10009, -0.30 DPS) [crafted]; Flame Leggings (253991, -0.89 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.98 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.0 spell_power points (3.73 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.86 DPS) [vendor]; Cindercloth Boots (10044, -0.93 DPS) [crafted]; Earthen Silk Slippers (254013, -1.77 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.92 DPS) | yes | Snake Hoop (6750, -0.53 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.86 DPS) [quest]; Reedknot Ring (9622, -0.99 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.73 DPS) | yes | Snake Hoop (6750, -0.34 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.67 DPS) [quest]; Reedknot Ring (9622, -0.80 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (147.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (147.2 DPS) | yes | Mark of the Chosen (17774, -2.25 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (147.2 DPS) | yes | Spellshifter Rod (9527, -0.18 DPS) [quest]; Kindling Stave (11750, -0.68 DPS) [dungeon]; Hanzo Sword (8190, -25.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 393.6 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.83 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.68 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Knight's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 506, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 2050015100000000000-23552100130125031-003000000000000000)

Set DPS (verified): 414.1. Weights run: 1.5s. Verify run: 3.6s. 1125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.845 ± 0.052, crit=0.465 ± 0.044 per rating point (14 rating = 1%, 6.503 per %), hit=1.218 ± 0.089 per rating point (10 rating = 1%, 12.181 per %), spell_haste=-11.057 ± 1.138, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (414.1 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Magister's Crown (16686, -7.74 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.4 spell_power points (6.37 DPS) | yes | Glowing Eye of Mordresh (10769, -2.08 DPS) [dungeon]; Horizon Choker (13085, -2.69 DPS) [world_drop]; Mindburst Medallion (11196, -2.92 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.5 spell_power points (10.74 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.15 DPS) [pvp]; Rotgrip Mantle (17732, -7.06 DPS, sim-verified) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (6.54 DPS) | yes | Spritecaster Cape (11623, -0.60 DPS) [dungeon]; Runecloth Cloak (13860, -1.63 DPS) [crafted]; Cindercloth Cloak (14044, -8.26 DPS, sim-verified) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.6 spell_power points (17.65 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.86 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.60 DPS) [pvp]; Robes of Fiery Devastation (279266, -9.72 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (414.1 DPS) | yes | Marshal's Silk Bracers (16438, +0.00 DPS) [pvp]; Netherflame Cuffs (254065, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -11.59 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.5 spell_power points (9.49 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Fiery Gloves (254099, -10.15 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (414.1 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Fiery Waistcord (254085, +0.00 DPS) [crafted]; Magister's Belt (16685, -14.34 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (414.1 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -10.00 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 30.1 spell_power points (9.39 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.21 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.73 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.9 spell_power points (7.45 DPS) | yes | Band of the Unicorn (7553, -3.40 DPS) [world_drop]; Snake Hoop (6750, -4.27 DPS) [quest]; Glowing Crystal Ring (18402, -4.29 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (4.43 DPS) | yes | Snake Hoop (6750, -1.25 DPS) [quest]; Glowing Crystal Ring (18402, -1.27 DPS) [quest]; Band of the Unicorn (7553, -8.22 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (414.1 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (414.1 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (414.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.33 DPS) [quest]; Teebu's Blazing Longsword (1728, -63.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (414.1 DPS) | yes | Stormrager (16997, -0.22 DPS) [quest]; Torch of Austen (13004, -7.44 DPS) [world_drop]; Torch of Light (279246, -12.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; ranged: Brilliant Wand

No-known-source sample (15 of 1125, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 2050015100000000000-23452100130135031-003000000000000000)

Set DPS (verified): 743.0. Weights run: 1.6s. Verify run: 2.2s. 1125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.438 ± 0.029, crit=0.763 ± 0.040 per rating point (14 rating = 1%, 10.676 per %), hit=1.133 ± 0.072 per rating point (10 rating = 1%, 11.331 per %), spell_haste=not significant (-0.035 ± 0.561), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 27.8 spell_power points (15.95 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, +0.00 DPS) [crafted]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.4 spell_power points (9.41 DPS) | yes | Glowing Eye of Mordresh (10769, -3.38 DPS) [dungeon]; Mindburst Medallion (11196, -4.45 DPS) [quest]; Triune Amulet (7722, -4.82 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 33.4 spell_power points (19.17 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.03 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.65 DPS) [pvp]; Argent Shoulders (19059, -12.32 DPS, sim-verified) [crafted] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (12.06 DPS) | yes | Spritecaster Cape (11623, -2.51 DPS) [dungeon]; Cindercloth Cloak (14044, -2.58 DPS) [crafted]; Cloak of Fire (14134, -4.02 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.9 spell_power points (32.13 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.76 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.66 DPS) [pvp]; Robes of Fiery Devastation (279266, -17.59 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 16.1 spell_power points (9.23 DPS) | yes | Runecloth Cuffs (254123, -0.39 DPS) [crafted]; Phoenix Bindings (210781, -1.76 DPS) [crafted]; Arena Wristguards (18709, -3.10 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 34.7 spell_power points (19.95 DPS) | yes | Marshal's Silk Gloves (16440, -1.42 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -1.42 DPS) [vendor]; Fiery Gloves (254099, -13.53 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 22.9 spell_power points (13.18 DPS) | yes | Satyrmane Sash (17755, -2.62 DPS) [dungeon]; Fiery Cord (254041, -3.12 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.34 DPS) [crafted] |
| legs | Cindercloth Pants (14045) | Tailoring [crafted] | 31.3 spell_power points (17.96 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -1.91 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.3 spell_power points (14.51 DPS) | yes | Knight-Lieutenant's Silk Walkers (227112, +0.00 DPS) [pvp]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.72 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.1 spell_power points (12.10 DPS) | yes | Philanthropist's Ring (281635, -5.10 DPS) [quest]; Snake Hoop (6750, -7.40 DPS) [quest]; Ring of Forlorn Spirits (2043, -7.51 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.47 DPS) | yes | Philanthropist's Ring (281635, -0.46 DPS) [quest]; Snake Hoop (6750, -2.76 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.87 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (743.0 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (743.0 DPS) | yes | Mark of the Chosen (17774, -7.78 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (743.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.68 DPS) [quest]; Teebu's Blazing Longsword (1728, -53.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 139.6 spell_power points (80.21 DPS) | yes | Brilliant Wand (249385, -13.69 DPS) [crafted]; Stormrager (16997, -16.26 DPS) [quest]; Greater Eternal Wand (249237, -21.71 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; hands: Gloves of Spell Mastery; legs: Cindercloth Pants; feet: Omnicast Boots; finger1: Maiden's Circle; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1125, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 0000000000000000000-23510000000000000-000000000000000000)

Set DPS (verified): 25.8. Weights run: 1.1s. Verify run: 0.8s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.413 ± 0.011, crit=0.127 ± 0.004 per rating point (14 rating = 1%, 1.778 per %), hit=0.389 ± 0.013 per rating point (10 rating = 1%, 3.895 per %), spell_haste=not significant (0.525 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.71 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Shadow Goggles (4373, -0.58 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.7 spell_power points (0.56 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.31 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.37 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Black Whelp Cloak (7283, -0.06 DPS) [crafted]; Mystic Shawl (3449, -0.13 DPS) [quest]; Pearl-clasped Cloak (5542, -0.24 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 11.1 spell_power points (0.72 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.32 DPS) [crafted]; Green Woolen Robe (6243, -0.44 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.19 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS, sim-verified) [quest]; Mindthrust Bracers (1974, -0.06 DPS) [dungeon]; Owlbeard Bracers (16981, -0.08 DPS) [quest] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.58 DPS) | yes | Flame Gloves (253917, -0.05 DPS) [crafted]; Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.7 spell_power points (0.56 DPS) | yes | Pristine Sash (253925, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.35 DPS) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.5 spell_power points (0.87 DPS) | yes | Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted]; Phoenix Pants (4317, -0.34 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.42 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 spell_power points (0.56 DPS) | yes | Flame Boots (253893, -0.25 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.29 DPS) [crafted]; Red Woolen Boots (4313, -0.30 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.5 spell_power points (0.16 DPS) | yes | Volcanic Rock Ring (12053, -0.08 DPS) [world_drop]; The 1 Ring (8350, -0.13 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.1 spell_power points (0.13 DPS) | yes | Volcanic Rock Ring (12053, -0.05 DPS) [world_drop]; The 1 Ring (8350, -0.11 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 4.1 spell_power points (0.27 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.11 DPS) [world]; Clear Crystal Rod (16894, -0.13 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 347.7 spell_power points (22.50 DPS) | yes | Skycaller (12984, -0.62 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.53 DPS) [quest]; Greater Magic Wand (11288, -4.87 DPS) [crafted] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Tabitha's Cuffs; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 0000000000000000000-23552100120000000-000000000000000000)

Set DPS (verified): 47.0. Weights run: 1.2s. Verify run: 0.8s. 282 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.700 ± 0.013, crit=0.172 ± 0.007 per rating point (14 rating = 1%, 2.406 per %), hit=0.360 ± 0.017 per rating point (10 rating = 1%, 3.602 per %), spell_haste=0.949 ± 0.187, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.32 DPS) | yes | Enchanter's Cowl (4322, -0.09 DPS) [crafted]; Holy Shroud (2721, -0.28 DPS) [world_drop]; Flame Circlet (253953, -0.28 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.2 spell_power points (0.40 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (1.44 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.47 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Resilient Cape (14400, -0.08 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.08 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.1 spell_power points (1.71 DPS) | yes | Flame Gown (253965, -0.02 DPS) [crafted]; Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.23 DPS) | yes | Spidertank Oilrag (9448, -0.38 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.70 DPS) [world_drop]; Nightsky Wristbands (6407, -0.83 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.5 spell_power points (0.90 DPS) | yes | Phoenix Gloves (4331, -0.05 DPS) [crafted]; Truefaith Gloves (7049, -0.23 DPS) [crafted]; Flame Gloves (253917, -0.36 DPS, sim-verified) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.1 spell_power points (1.05 DPS) | yes | Crimson Silk Belt (7055, -0.02 DPS) [crafted]; Invoker's Cord (215366, -0.06 DPS) [crafted]; Warsong Sash (16975, -0.54 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.9 spell_power points (1.78 DPS) | yes | Filigreed Flame Leggings (253941, -0.55 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.56 DPS) [world]; Phoenix Pants (4317, -0.58 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.9 spell_power points (1.69 DPS) | yes | Gilded Slippers (254001, -0.57 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.90 DPS) | yes | Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon]; Solid Gold Band (279845, -0.52 DPS) [quest]; Azora's Will (4999, -0.57 DPS) [world_drop] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.57 DPS) | yes | Solid Gold Band (279845, -0.19 DPS) [quest]; Azora's Will (4999, -0.24 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.63 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.39 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -2.67 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -2.73 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 351.3 spell_power points (33.12 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Unstable Power Core (279847, -3.41 DPS) [quest]; Greater Mystic Wand (217287, -3.65 DPS) [crafted] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 282, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (orc, 0000000000000000000-23552100130125010-000000000000000000)

Set DPS (verified): 77.5. Weights run: 1.4s. Verify run: 1.0s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.916 ± 0.031, crit=0.259 ± 0.020 per rating point (14 rating = 1%, 3.630 per %), hit=0.489 ± 0.040 per rating point (10 rating = 1%, 4.886 per %), spell_haste=2.140 ± 0.521, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.6 spell_power points (2.72 DPS) | yes | Augural Shroud (2620, -0.39 DPS) [world]; Whitemane's Chapeau (7720, -0.85 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.36 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.3 spell_power points (1.65 DPS) | yes | Necklace of Calisea (1714, -0.91 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.02 DPS) [world]; Triune Amulet (7722, -1.55 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.9 spell_power points (2.18 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.6 spell_power points (1.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.37 DPS) [dungeon]; Icy Cloak (4327, -0.41 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.5 spell_power points (3.17 DPS) | yes | Cindercloth Robe (10042, -0.06 DPS) [crafted]; Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.29 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.50 DPS) | yes | Radiant Silver Bracers (4545, -0.19 DPS) [quest]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Windchaser Cuffs (14429, -0.55 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.7 spell_power points (2.50 DPS) | yes | Fiery Handwraps (254025, -0.14 DPS) [crafted]; Red Mageweave Gloves (10018, -0.17 DPS) [crafted]; Crimson Silk Gloves (7064, -0.25 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 21.3 spell_power points (2.46 DPS) | yes | Deathmage Sash (10771, -0.07 DPS) [dungeon]; Gilded Cord (254037, -0.69 DPS) [crafted]; Scorching Sash (4117, -0.96 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.0 spell_power points (2.88 DPS) | yes | Flame Leggings (253991, -0.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.59 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.98 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Fiery Slippers (254005, -0.86 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.22 DPS) [crafted]; Acidic Walkers (9454, -1.35 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.68 DPS) | yes | Reedknot Ring (9622, -0.87 DPS) [quest]; Ogremind Ring (1993, -0.94 DPS) [world_drop]; Voodoo Band (1996, -0.94 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.6 spell_power points (1.22 DPS) | yes | Ogremind Ring (1993, -0.48 DPS) [world_drop]; Voodoo Band (1996, -0.48 DPS) [world]; Reedknot Ring (9622, -0.92 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (77.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.21 DPS) [dungeon]; Searing Golden Blade (12260, -0.68 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.7 spell_power points (1.82 DPS) | yes | Orb of the Forgotten Seer (7685, -0.20 DPS) [dungeon]; Rod of Molten Fire (2565, -0.32 DPS) [world_drop]; Orb of Mystic Insight (249394, -0.37 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 347.1 spell_power points (40.01 DPS) | yes | Ragefire Wand (7513, -1.52 DPS) [quest]; Nether Force Wand (11263, -2.37 DPS) [quest]; Icefury Wand (7514, -2.51 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 2050000000000000000-23552100130125031-000000000000000000)

Set DPS (verified): 141.2. Weights run: 1.3s. Verify run: 1.2s. 496 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.886 ± 0.034, crit=0.303 ± 0.027 per rating point (14 rating = 1%, 4.237 per %), hit=0.558 ± 0.056 per rating point (10 rating = 1%, 5.585 per %), spell_haste=not significant (0.439 ± 0.642), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.7 spell_power points (4.90 DPS) | yes | Dreamweave Circlet (10041, -0.92 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.05 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.30 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.1 spell_power points (1.88 DPS) | yes | Mindburst Medallion (11196, -0.37 DPS) [quest]; Triune Amulet (7722, -0.81 DPS) [dungeon]; Horizon Choker (13085, -1.58 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.0 spell_power points (3.86 DPS) | yes | Netherflame Shoulders (254053, -0.65 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.17 DPS) [vendor] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.1 spell_power points (2.68 DPS) | yes | Deep Woodlands Cloak (19121, -0.02 DPS) [quest]; Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.7 spell_power points (4.90 DPS) | yes | Cindercloth Robe (10042, -1.30 DPS) [crafted]; Runecloth Tunic (13857, -1.33 DPS) [crafted]; Robe of the Magi (1716, -2.78 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.2 spell_power points (2.56 DPS) | yes | Nethergeld Cuffs (254061, -0.80 DPS) [crafted]; Phoenix Bindings (210781, -0.83 DPS) [crafted]; Bloodband Bracers (11469, -0.83 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 27.9 spell_power points (3.72 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -0.92 DPS) [vendor]; Fiery Handwraps (254025, -1.02 DPS) [crafted]; Dreamweave Gloves (10019, -2.32 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.0 spell_power points (3.60 DPS) | yes | Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Dawnspire Cord (12466, -0.55 DPS) [dungeon]; Fiery Cord (254041, -0.79 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 26.9 spell_power points (3.58 DPS) | yes | Red Mageweave Pants (10009, -0.30 DPS) [crafted]; Flame Leggings (253991, -0.89 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.98 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.0 spell_power points (3.73 DPS) | yes | Earthen Silk Slippers (254013, -0.53 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -0.86 DPS) [vendor]; Cindercloth Boots (10044, -0.93 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.92 DPS) | yes | Snake Hoop (6750, -0.53 DPS) [quest]; Reedknot Ring (9622, -0.99 DPS) [quest]; Voodoo Band (1996, -1.10 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.73 DPS) | yes | Snake Hoop (6750, -0.34 DPS) [quest]; Reedknot Ring (9622, -0.80 DPS) [quest]; Voodoo Band (1996, -0.91 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (141.2 DPS) | yes | Rune of the Guard Captain (19120, -0.28 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (141.2 DPS) | yes | Rune of the Guard Captain (19120, -0.15 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (141.2 DPS) | yes | Spellshifter Rod (9527, -0.18 DPS) [quest]; Kindling Stave (11750, -0.68 DPS) [dungeon]; Hanzo Sword (8190, -24.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 393.6 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.83 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.68 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Stone Guard's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 496, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 2050015100000000000-23552100130125031-003000000000000000)

Set DPS (verified): 384.0. Weights run: 1.5s. Verify run: 3.5s. 1117 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.845 ± 0.052, crit=0.465 ± 0.044 per rating point (14 rating = 1%, 6.503 per %), hit=1.218 ± 0.089 per rating point (10 rating = 1%, 12.181 per %), spell_haste=-11.057 ± 1.138, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 36.3 spell_power points (11.32 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -8.87 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.4 spell_power points (6.37 DPS) | yes | Glowing Eye of Mordresh (10769, -2.08 DPS) [dungeon]; Horizon Choker (13085, -2.69 DPS) [world_drop]; Mindburst Medallion (11196, -2.92 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.5 spell_power points (10.74 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.15 DPS) [pvp]; Rotgrip Mantle (17732, -7.81 DPS, sim-verified) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (6.54 DPS) | yes | Cindercloth Cloak (14044, -0.39 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.43 DPS) [quest]; Spritecaster Cape (11623, -0.60 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.6 spell_power points (17.65 DPS) | yes | Warlord's Silk Raiment (231596, -0.86 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.60 DPS) [pvp]; Robes of Fiery Devastation (279266, -12.80 DPS, sim-verified) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 19.4 spell_power points (6.06 DPS) | yes | Netherflame Cuffs (254065, -0.17 DPS) [crafted]; General's Silk Cuffs (16538, -1.06 DPS) [pvp]; Magiskull Cuffs (13107, -1.58 DPS) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.5 spell_power points (9.49 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Fiery Gloves (254099, +0.00 DPS) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (384.0 DPS) | yes | Wisdom of the Timbermaw (19047, -0.60 DPS) [crafted]; Satyrmane Sash (17755, -0.71 DPS) [dungeon]; Fiery Waistcord (254085, -8.70 DPS, sim-verified) [crafted] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (384.0 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -12.13 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 30.1 spell_power points (9.39 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.21 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -0.73 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.9 spell_power points (7.45 DPS) | yes | Band of the Unicorn (7553, -3.40 DPS) [world_drop]; Snake Hoop (6750, -4.27 DPS) [quest]; Glowing Crystal Ring (18402, -4.29 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (4.43 DPS) | yes | Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Snake Hoop (6750, -1.25 DPS) [quest]; Glowing Crystal Ring (18402, -1.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (384.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (384.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.33 DPS) [quest]; Teebu's Blazing Longsword (1728, -56.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 252.4 spell_power points (78.63 DPS) | yes | Brilliant Wand (249385, -14.40 DPS) [crafted]; Stormrager (16997, -14.62 DPS) [quest]; Torch of Austen (13004, -21.84 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 2050015100000000000-23452100130135031-003000000000000000)

Set DPS (verified): 724.1. Weights run: 1.6s. Verify run: 2.2s. 1117 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.438 ± 0.029, crit=0.763 ± 0.040 per rating point (14 rating = 1%, 10.676 per %), hit=1.133 ± 0.072 per rating point (10 rating = 1%, 11.331 per %), spell_haste=not significant (-0.035 ± 0.561), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 27.8 spell_power points (15.95 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, +0.00 DPS) [crafted]; Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.4 spell_power points (9.41 DPS) | yes | Glowing Eye of Mordresh (10769, -3.38 DPS) [dungeon]; Mindburst Medallion (11196, -4.45 DPS) [quest]; Triune Amulet (7722, -4.82 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 33.4 spell_power points (19.17 DPS) | yes | Warlord's Silk Amice (231594, -1.03 DPS) [pvp]; Champion's Silk Mantle (227104, -1.65 DPS) [pvp]; Argent Shoulders (19059, -4.81 DPS) [crafted] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (12.06 DPS) | yes | Spritecaster Cape (11623, -2.51 DPS) [dungeon]; Cindercloth Cloak (14044, -2.58 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.90 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.9 spell_power points (32.13 DPS) | yes | Warlord's Silk Raiment (231596, -2.76 DPS) [pvp]; Robes of Fiery Devastation (279266, -9.59 DPS, sim-verified) [crafted]; Legionnaire's Silk Tunic (227106, -9.66 DPS) [pvp] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 16.1 spell_power points (9.23 DPS) | yes | Runecloth Cuffs (254123, -0.39 DPS) [crafted]; Phoenix Bindings (210781, -1.76 DPS) [crafted]; Arena Wristguards (18709, -3.10 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 34.7 spell_power points (19.95 DPS) | yes | General's Silk Handguards (16540, -1.42 DPS) [vendor]; General's Silk Gauntlets (231599, -1.42 DPS) [vendor]; Fiery Gloves (254099, -7.06 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 22.9 spell_power points (13.18 DPS) | yes | Satyrmane Sash (17755, -2.62 DPS) [dungeon]; Fiery Cord (254041, -3.12 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.34 DPS) [crafted] |
| legs | Cindercloth Pants (14045) | Tailoring [crafted] | 31.3 spell_power points (17.96 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Stone Guard's Dreadweave Leggings (220906, -1.91 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.3 spell_power points (14.51 DPS) | yes | Blood Guard's Silk Walkers (227110, +0.00 DPS) [pvp]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.72 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.1 spell_power points (12.10 DPS) | yes | Philanthropist's Ring (281635, -5.10 DPS) [quest]; Snake Hoop (6750, -7.40 DPS) [quest]; Reedknot Ring (9622, -8.08 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.47 DPS) | yes | Philanthropist's Ring (281635, -0.46 DPS) [quest]; Snake Hoop (6750, -2.76 DPS) [quest]; Reedknot Ring (9622, -3.45 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (724.1 DPS) | yes | Uther's Strength (11302, -1.11 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (724.1 DPS) | yes | Uther's Strength (11302, -2.41 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (724.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.68 DPS) [quest]; Teebu's Blazing Longsword (1728, -51.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 139.6 spell_power points (80.21 DPS) | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -16.26 DPS) [quest]; Greater Eternal Wand (249237, -21.71 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; hands: Gloves of Spell Mastery; legs: Cindercloth Pants; feet: Omnicast Boots; finger1: Maiden's Circle; trinket1: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

