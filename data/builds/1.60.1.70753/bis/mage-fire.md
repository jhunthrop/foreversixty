# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-51500000000000000-000000000000000000)

Set DPS (verified): 32.7. Weights run: 1.1s. Verify run: 0.8s. 185 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.459 ± 0.012, crit=0.112 ± 0.005 per rating point (14 rating = 1%, 1.565 per %), hit=0.240 ± 0.012 per rating point (10 rating = 1%, 2.405 per %), spell_haste=not significant (0.405 ± 0.166), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.067 ± 0.045), mp5=0.724 ± 0.143

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.12 DPS) | yes | Pristine Circlet (253949, -0.81 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.84 DPS) [crafted]; Flying Tiger Goggles (4368, -1.09 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.1 spell_power points (0.01 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.1 spell_power points (0.93 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.52 DPS) [crafted]; Slime-encrusted Pads (6461, -0.71 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.72 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.3 spell_power points (0.43 DPS) | yes | Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Sanguine Cape (14376, -0.25 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.29 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 11.5 spell_power points (1.17 DPS) | yes | Gray Woolen Robe (2585, -0.53 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.65 DPS, sim-verified) [crafted]; Green Woolen Robe (6243, -0.72 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 4.5 spell_power points (0.45 DPS) | yes | Repurposed Hair Band (281256, -0.35 DPS) [quest]; Windsong Bangles (263336, -0.35 DPS) [quest]; Bright Bracers (3647, -0.39 DPS, sim-verified) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.91 DPS) | yes | Flame Gloves (253917, -0.06 DPS) [crafted]; Serpent Gloves (5970, -0.20 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.30 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.8 spell_power points (0.90 DPS) | yes | Novice Arcanist's Sash (253885, -0.35 DPS) [crafted]; Pristine Sash (253925, -0.49 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.52 DPS) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 14.0 spell_power points (1.43 DPS) | yes | Filigreed Pristine Leggings (253937, -0.51 DPS) [crafted]; Phoenix Pants (4317, -0.61 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.71 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.8 spell_power points (0.90 DPS) | yes | Flame Boots (253893, -0.15 DPS) [crafted]; Pristine Boots (253889, -0.45 DPS) [crafted]; Red Woolen Boots (4313, -0.46 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.9 spell_power points (0.60 DPS) | yes | Volcanic Rock Ring (12053, -0.46 DPS) [world_drop]; Furlfeather Ring (286745, -0.50 DPS) [world]; The 1 Ring (8350, -0.55 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.8 spell_power points (0.28 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; Furlfeather Ring (286745, -0.18 DPS) [world]; The 1 Ring (8350, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 4.6 spell_power points (0.47 DPS) | yes | Staff of Westfall (2042, -0.19 DPS) [quest]; Witching Stave (1484, -0.23 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.26 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 219.5 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.75 DPS) [world_drop]; Deepblaze (279896, -3.78 DPS) [quest]; Sizzle Stick (8071, -4.46 DPS) [quest] |

**New at 20:** head: Flame Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 185, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 000000000000000000-51521001230100000-000000000000000000)

Set DPS (verified): 68.8. Weights run: 1.2s. Verify run: 0.9s. 323 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.747 ± 0.018, crit=0.175 ± 0.011 per rating point (14 rating = 1%, 2.451 per %), hit=0.290 ± 0.017 per rating point (10 rating = 1%, 2.896 per %), spell_haste=not significant (0.092 ± 0.106), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.275 ± 0.064, mp5=not significant (0.406 ± 0.123)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (2.00 DPS) | yes | Enchanter's Cowl (4322, -0.08 DPS) [crafted]; Holy Shroud (2721, -0.19 DPS) [world_drop]; Flame Circlet (253953, -0.43 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.5 spell_power points (0.64 DPS) | yes | Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.41 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.7 spell_power points (2.25 DPS) | yes | Magician's Mantle (12998, -0.57 DPS) [world_drop]; Invoker's Mantle (215365, -0.71 DPS) [crafted]; Death Speaker Mantle (6685, -1.16 DPS, sim-verified) [dungeon] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.6 spell_power points (0.80 DPS) | yes | Glowing Thresher Cape (6901, -0.06 DPS) [dungeon]; Heavy Woolen Cloak (4311, -0.07 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.09 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 20.2 spell_power points (2.89 DPS) | yes | Green Silk Armor (7065, -0.21 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.46 DPS) [dungeon]; Flame Gown (253965, -1.37 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.86 DPS) | yes | Spidertank Oilrag (9448, -0.40 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.01 DPS) [world_drop]; Nightsky Wristbands (6407, -1.10 DPS) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.2 spell_power points (1.32 DPS) | yes | Phoenix Gloves (4331, -0.03 DPS) [crafted]; Town Clerk's Mittens (270029, -0.14 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.27 DPS) [dungeon] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 12.6 spell_power points (1.81 DPS) | yes | Invoker's Cord (215366, -0.27 DPS) [crafted]; Flame Sash (253929, -0.38 DPS) [crafted]; Crimson Silk Belt (7055, -2.62 DPS, sim-verified) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.6 spell_power points (2.95 DPS) | yes | Filigreed Flame Leggings (253941, -0.68 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -0.88 DPS) [dungeon]; Pristine Leggings (253987, -1.00 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.3 spell_power points (2.77 DPS) | yes | Acidic Walkers (9454, -1.04 DPS) [dungeon]; Gilded Slippers (254001, -1.16 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.34 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (1.39 DPS) | yes | Minor Channeling Ring (1449, -0.46 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.69 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 7.2 spell_power points (1.03 DPS) | yes | Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Runemender's Seal (273031, -0.33 DPS) [world]; Minor Channeling Ring (1449, -0.67 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 8.2 spell_power points (1.18 DPS) | yes | Cryptbone Staff (2013, -0.38 DPS) [world]; Staff of the Friar (3415, -0.38 DPS) [dungeon]; Twisted Chanter's Staff (890, -6.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 232.4 spell_power points (33.28 DPS) | yes | Wand of Mana Concentration (273043, -1.56 DPS) [world]; Gyromatic Icemaker (9489, -2.93 DPS) [dungeon]; Greater Mystic Wand (217287, -3.57 DPS) [crafted] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Beguiler Robes; wrist: Phoenix Bindings; hands: Flame Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Glimmering Staff; ranged: Starfaller

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 000000000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 147.9. Weights run: 1.4s. Verify run: 1.0s. 421 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.762 ± 0.025, crit=0.248 ± 0.018 per rating point (14 rating = 1%, 3.475 per %), hit=0.366 ± 0.031 per rating point (10 rating = 1%, 3.659 per %), spell_haste=not significant (0.229 ± 0.260), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.225 ± 0.166), mp5=not significant (0.468 ± 0.237)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.0 spell_power points (4.57 DPS) | yes | Augural Shroud (2620, -0.87 DPS) [world]; Whitemane's Chapeau (7720, -1.13 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.62 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.1 spell_power points (2.61 DPS) | yes | Triune Amulet (7722, -0.70 DPS) [dungeon]; Necklace of Calisea (1714, -1.23 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.70 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.9 spell_power points (3.36 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.21 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.8 spell_power points (1.95 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.29 DPS) [crafted]; Blackforge Cape (6424, -0.56 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 28.0 spell_power points (5.57 DPS) | yes | Cindercloth Robe (10042, -0.20 DPS) [crafted]; Robe of Power (7054, -0.61 DPS) [crafted]; Robe of the Magi (1716, -4.30 DPS, sim-verified) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (2.59 DPS) | yes | Spidertank Oilrag (9448, -0.52 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.80 DPS) [quest]; Condor Bracers (15864, -1.19 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.6 spell_power points (4.50 DPS) | yes | Red Mageweave Gloves (10018, -0.80 DPS) [crafted]; Crimson Silk Gloves (7064, -0.81 DPS) [crafted]; Fiery Handwraps (254025, -2.09 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 21.4 spell_power points (4.27 DPS) | yes | Gilded Cord (254037, -1.19 DPS) [crafted]; Scorching Sash (4117, -1.28 DPS) [quest]; Deathmage Sash (10771, -1.76 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.1 spell_power points (4.61 DPS) | yes | Flame Leggings (253991, -0.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.04 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.33 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.6 spell_power points (5.09 DPS) | yes | Fiery Slippers (254005, -1.26 DPS) [crafted]; Gilded Slippers (254001, -2.46 DPS) [crafted]; Acidic Walkers (9454, -2.70 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (2.75 DPS) | yes | Ring of Forlorn Spirits (2043, -1.16 DPS) [quest]; Ring of Power Regulation (273030, -1.27 DPS) [dungeon]; Reedknot Ring (9622, -1.36 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.8 spell_power points (1.95 DPS) | yes | Ring of Power Regulation (273030, -0.48 DPS) [dungeon]; Reedknot Ring (9622, -0.56 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.43 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (147.9 DPS) | yes | Illusionary Rod (7713, -2.66 DPS) [dungeon]; Glimmering Staff (249392, -2.90 DPS) [crafted]; Searing Golden Blade (12260, -14.82 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 203.1 spell_power points (40.42 DPS) | yes | Ragefire Wand (7513, -1.03 DPS) [quest]; Nether Force Wand (11263, -2.63 DPS) [quest]; Icefury Wand (7514, -2.78 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 252010000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 190.7. Weights run: 1.3s. Verify run: 1.2s. 537 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.708 ± 0.037, crit=0.357 ± 0.028 per rating point (14 rating = 1%, 4.991 per %), hit=0.536 ± 0.046 per rating point (10 rating = 1%, 5.360 per %), spell_haste=not significant (0.071 ± 0.368), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.991 ± 0.331), mp5=not significant (1.052 ± 0.342)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 40.0 spell_power points (7.79 DPS) | yes | Gemburst Circlet (10751, -1.19 DPS) [quest]; Red Mageweave Headband (10033, -1.33 DPS) [crafted]; Papal Fez (9431, -2.15 DPS) [dungeon] |
| neck | Mindburst Medallion (11196) | Tremors of the Earth [quest] | 15.2 spell_power points (2.96 DPS) | yes | Triune Amulet (7722, -0.05 DPS) [dungeon]; Horizon Choker (13085, -0.26 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -0.49 DPS) [dungeon] |
| shoulder | Netherflame Shoulders (254053) | Tailoring [crafted] | 28.6 spell_power points (5.57 DPS) | yes | Rotgrip Mantle (17732, -0.56 DPS) [dungeon]; Flameseer Mantle (11310, -0.92 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.56 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.2 spell_power points (4.71 DPS) | yes | Runecloth Cloak (13860, -1.86 DPS) [crafted]; Cindercloth Cloak (14044, -1.93 DPS, sim-verified) [crafted]; Cloak of Fire (14134, -1.99 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 38.2 spell_power points (7.45 DPS) | yes | Ghostweave Vest (14141, -0.20 DPS) [crafted]; Cindercloth Vest (14042, -0.65 DPS) [crafted]; Acumen Robes (17775, -0.99 DPS) [quest] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.9 spell_power points (4.46 DPS) | yes | Nethergeld Cuffs (254061, -1.17 DPS) [crafted]; Bloodband Bracers (11469, -1.67 DPS) [quest]; Phoenix Bindings (210781, -1.93 DPS) [crafted] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 34.0 spell_power points (6.63 DPS) | yes | Dreamweave Gloves (10019, -1.22 DPS) [crafted]; Runecloth Gloves (13863, -1.31 DPS) [crafted]; Cindercloth Gloves (14043, -3.82 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 25.6 spell_power points (4.99 DPS) | yes | Fiery Waistcord (254085, -0.05 DPS) [crafted]; Ghostweave Cord (254073, -0.14 DPS) [crafted]; Scorching Sash (4117, -0.72 DPS) [quest] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 34.8 spell_power points (6.78 DPS) | yes | Cindercloth Leggings (12256, -0.18 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.24 DPS) [dungeon]; Dalewind Trousers (13008, -0.71 DPS) [world_drop] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.9 spell_power points (6.03 DPS) | yes | Fiery Sandals (254111, -0.89 DPS) [crafted]; Vinerot Sandals (17748, -1.69 DPS) [dungeon]; Fiery Slippers (254005, -1.76 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 16.1 spell_power points (3.14 DPS) | yes | Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Choking Band (11868, -0.64 DPS) [quest]; Coldwater Ring (4550, -0.99 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (2.64 DPS) | yes | Band of the Unicorn (7553, -0.11 DPS) [world_drop]; Choking Band (11868, -0.13 DPS) [quest]; Coldwater Ring (4550, -0.49 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (190.7 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (190.7 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (190.7 DPS) | yes | Spellshifter Rod (9527, -0.62 DPS) [quest]; Inventor's Focal Sword (17719, -1.66 DPS) [dungeon]; Spire of Hakkar (10844, -2.98 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 285.5 spell_power points (55.61 DPS) | yes | Noxious Shooter (17745, -4.01 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -7.19 DPS) [world_drop]; Lesser Eternal Wand (249232, -8.30 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Mindburst Medallion; shoulder: Netherflame Shoulders; back: Spritecaster Cape; wrist: Netherflame Cuffs; hands: Fiery Gloves; legs: Senior Designer's Pantaloons; finger1: Eye of Adaegus; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 537, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 252015000000200000-51521001230133031-003000000000000000)

Set DPS (verified): 352.3. Weights run: 1.5s. Verify run: 3.4s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.624 ± 0.047, crit=0.540 ± 0.040 per rating point (14 rating = 1%, 7.566 per %), hit=0.894 ± 0.074 per rating point (10 rating = 1%, 8.941 per %), spell_haste=-6.621 ± 0.737, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.487 ± 0.608), mp5=not significant (0.452 ± 0.619)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (352.3 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Magister's Crown (16686, -8.88 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.2 spell_power points (4.48 DPS) | yes | Mindburst Medallion (11196, -1.49 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.53 DPS) [dungeon]; Triune Amulet (7722, -1.68 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 32.7 spell_power points (8.03 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.32 DPS) [pvp]; Shroud of the Nathrezim (18720, -1.04 DPS) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (5.16 DPS) | yes | Spritecaster Cape (11623, -0.08 DPS) [dungeon]; Cindercloth Cloak (14044, -0.74 DPS) [crafted]; Cloak of Fire (14134, -1.72 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.1 spell_power points (13.52 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.23 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.30 DPS) [pvp]; Robes of Fiery Devastation (279266, -5.44 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (352.3 DPS) | yes | Netherflame Cuffs (254065, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Magister's Bindings (16683, -6.42 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (352.3 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Magister's Gloves (16684, -4.24 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (352.3 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Fiery Waistcord (254085, +0.00 DPS) [crafted]; Magister's Belt (16685, -8.33 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (352.3 DPS) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (352.3 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, +0.00 DPS) [crafted]; Omnicast Boots (11822, -6.92 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.3 spell_power points (6.21 DPS) | yes | Band of the Unicorn (7553, -3.02 DPS) [world_drop]; Band of the Hierophant (13096, -3.36 DPS) [world_drop]; Eye of Adaegus (5266, -3.86 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.22 DPS) | yes | Band of the Unicorn (7553, -0.03 DPS) [world_drop]; Band of the Hierophant (13096, -0.37 DPS) [world_drop]; Eye of Adaegus (5266, -0.87 DPS) [world_drop] |
| trinket1 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (352.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Thunderbrew's Boot Flask (744, -0.24 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (352.3 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (352.3 DPS) | yes | Grand Marshal's Stave (234571, -1.41 DPS) [pvp]; Spire of Hakkar (10844, -11.41 DPS) [world]; Kindling Stave (11750, -18.10 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 323.9 spell_power points (79.56 DPS) | yes | Stormrager (16997, -8.33 DPS, sim-verified) [quest]; Brilliant Wand (249385, -16.69 DPS) [crafted]; Torch of Austen (13004, -22.77 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; trinket1: Evonice's Landin' Pilla; trinket2: Ankh of Life; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 253005000000200000-51521001230133031-003000000000000000)

Set DPS (verified): 690.9. Weights run: 1.6s. Verify run: 2.3s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.551 ± 0.039, crit=0.826 ± 0.046 per rating point (14 rating = 1%, 11.568 per %), hit=1.113 ± 0.075 per rating point (10 rating = 1%, 11.130 per %), spell_haste=-3.538 ± 0.639, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.002, spirit=not significant (0.784 ± 0.350), mp5=not significant (1.003 ± 0.408)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreadmist Mask (16698) | Scholomance: Darkmaster Gandling [dungeon] | 36.1 spell_power points (19.25 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -0.09 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.5 spell_power points (9.34 DPS) | yes | Mindburst Medallion (11196, -2.28 DPS) [quest]; Lady Alizabeth's Pendant (13002, -3.06 DPS) [world_drop]; Triune Amulet (7722, -7.87 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.7 spell_power points (19.06 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -4.85 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.0 spell_power points (11.74 DPS) | yes | Mageflame Cloak (13007, -0.54 DPS) [world_drop]; Cindercloth Cloak (14044, -2.46 DPS) [crafted]; Cloak of Fire (14134, -4.27 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.2 spell_power points (31.03 DPS) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.57 DPS) [pvp]; Robes of Fiery Devastation (279266, -17.19 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 20.8 spell_power points (11.08 DPS) | yes | Marshal's Silk Bracers (16438, -1.31 DPS) [pvp]; Runecloth Cuffs (254123, -2.28 DPS) [crafted]; Arena Wristguards (18709, -2.40 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.6 spell_power points (20.08 DPS) | yes | Marshal's Silk Gloves (16440, -0.06 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -0.06 DPS) [vendor]; Fiery Gloves (254099, -14.00 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 25.8 spell_power points (13.75 DPS) | yes | Marshal's Silk Sash (16439, -0.71 DPS) [pvp]; Fiery Waistcord (254085, -0.97 DPS) [crafted]; Dreadmist Belt (16702, -1.07 DPS) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 41.8 spell_power points (22.27 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -0.04 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 29.5 spell_power points (15.73 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.65 DPS) [crafted]; Wolfrunner Shoes (13101, -0.96 DPS) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 26.6 spell_power points (14.17 DPS) | yes | Band of the Unicorn (7553, -7.23 DPS) [world_drop]; Philanthropist's Ring (281635, -7.36 DPS) [quest]; Eye of Adaegus (5266, -7.38 DPS) [world_drop] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 14.1 spell_power points (7.54 DPS) | yes | Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Philanthropist's Ring (281635, -0.74 DPS) [quest]; Eye of Adaegus (5266, -0.76 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (690.9 DPS) | yes | Ankh of Life (1713, -0.32 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.16 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (690.9 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (690.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -20.08 DPS, sim-verified) [dungeon]; Spire of Hakkar (10844, -21.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 158.6 spell_power points (84.57 DPS) | yes | Brilliant Wand (249385, -18.13 DPS) [crafted]; Stormrager (16997, -20.40 DPS) [quest]; Pyric Caduceus (11748, -23.88 DPS) [dungeon] |

**New at 60:** head: Dreadmist Mask; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; hands: Gloves of Spell Mastery; waist: Wisdom of the Timbermaw; legs: Magister's Leggings; finger1: Maiden's Circle; finger2: Band of the Hierophant; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 000000000000000000-51500000000000000-000000000000000000)

Set DPS (verified): 34.5. Weights run: 1.1s. Verify run: 0.9s. 179 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.459 ± 0.012, crit=0.112 ± 0.005 per rating point (14 rating = 1%, 1.565 per %), hit=0.240 ± 0.012 per rating point (10 rating = 1%, 2.405 per %), spell_haste=not significant (0.405 ± 0.166), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.067 ± 0.045), mp5=0.724 ± 0.143

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (1.12 DPS) | yes | Pristine Circlet (253949, -0.78 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.84 DPS) [crafted]; Flying Tiger Goggles (4368, -1.09 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.2 spell_power points (0.02 DPS) | yes | Scholarly Pendant (277203, -0.01 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.1 spell_power points (0.93 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.52 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.64 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.71 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.3 spell_power points (0.43 DPS) | yes | Pearl-clasped Cloak (5542, -0.09 DPS) [crafted]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Mystic Shawl (3449, -0.23 DPS) [quest] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 11.5 spell_power points (1.17 DPS) | yes | Gray Woolen Robe (2585, -0.53 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.62 DPS, sim-verified) [crafted]; Green Woolen Robe (6243, -0.72 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 4.5 spell_power points (0.45 DPS) | yes | Owlbeard Bracers (16981, -0.26 DPS) [quest]; Bright Bracers (3647, -0.27 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.73 DPS, sim-verified) [quest] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.91 DPS) | yes | Flame Gloves (253917, -0.06 DPS) [crafted]; Serpent Gloves (5970, -0.20 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.30 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.8 spell_power points (0.90 DPS) | yes | Novice Arcanist's Sash (253885, -0.35 DPS) [crafted]; Pristine Sash (253925, -0.47 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.52 DPS) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 14.0 spell_power points (1.43 DPS) | yes | Filigreed Pristine Leggings (253937, -0.51 DPS) [crafted]; Phoenix Pants (4317, -0.66 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.71 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.8 spell_power points (0.90 DPS) | yes | Pristine Boots (253889, -0.45 DPS) [crafted]; Red Woolen Boots (4313, -0.46 DPS) [crafted]; Flame Boots (253893, -0.66 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.8 spell_power points (0.28 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; Furlfeather Ring (286745, -0.18 DPS) [world]; The 1 Ring (8350, -0.23 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.3 spell_power points (0.23 DPS) | yes | Furlfeather Ring (286745, -0.13 DPS) [world]; The 1 Ring (8350, -0.18 DPS) [world]; Volcanic Rock Ring (12053, -0.40 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-verified (34.5 DPS) | yes | Lesser Staff of the Spire (1300, -0.05 DPS) [world]; Staff of Orgrimmar (15444, -0.05 DPS) [quest]; Twisted Chanter's Staff (890, -3.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 219.5 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.31 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.46 DPS) [quest]; Greater Magic Wand (11288, -4.60 DPS) [crafted] |

**New at 20:** head: Flame Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Crescent Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 000000000000000000-51521001230100000-000000000000000000)

Set DPS (verified): 67.1. Weights run: 1.2s. Verify run: 1.0s. 313 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.747 ± 0.018, crit=0.175 ± 0.011 per rating point (14 rating = 1%, 2.451 per %), hit=0.290 ± 0.017 per rating point (10 rating = 1%, 2.896 per %), spell_haste=not significant (0.092 ± 0.106), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=0.275 ± 0.064, mp5=not significant (0.406 ± 0.123)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (2.00 DPS) | yes | Enchanter's Cowl (4322, -0.08 DPS) [crafted]; Holy Shroud (2721, -0.19 DPS) [world_drop]; Flame Circlet (253953, -0.43 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.5 spell_power points (0.64 DPS) | yes | Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.41 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.7 spell_power points (2.25 DPS) | yes | Death Speaker Mantle (6685, -0.22 DPS) [dungeon]; Magician's Mantle (12998, -0.57 DPS) [world_drop]; Invoker's Mantle (215365, -0.71 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | 6.1 spell_power points (0.87 DPS) | yes | Glowing Thresher Cape (6901, -0.13 DPS) [dungeon]; Heavy Woolen Cloak (4311, -0.14 DPS) [crafted]; Hillman's Cloak (3719, -0.16 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 20.2 spell_power points (2.89 DPS) | yes | Flame Gown (253965, -0.12 DPS) [crafted]; Green Silk Armor (7065, -0.21 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.46 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Magical Bracelets (13106, -0.61 DPS) [world_drop]; Nightsky Wristbands (6407, -0.70 DPS) [world_drop]; Phoenix Bindings (210781, -0.75 DPS, sim-verified) [crafted] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.7 spell_power points (1.39 DPS) | yes | Phoenix Gloves (4331, -0.11 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.34 DPS) [dungeon]; Flame Gloves (253917, -0.82 DPS, sim-verified) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 12.6 spell_power points (1.81 DPS) | yes | Crimson Silk Belt (7055, -0.20 DPS) [crafted]; Warsong Sash (16975, -0.23 DPS) [quest]; Invoker's Cord (215366, -0.27 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 20.6 spell_power points (2.95 DPS) | yes | Gaze Dreamer Pants (6903, -0.88 DPS) [dungeon]; Pristine Leggings (253987, -1.00 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.18 DPS, sim-verified) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 19.3 spell_power points (2.77 DPS) | yes | Acidic Walkers (9454, -1.04 DPS) [dungeon]; Gilded Slippers (254001, -1.12 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.34 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (1.39 DPS) | yes | Solid Gold Band (279845, -0.62 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.69 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 7.2 spell_power points (1.03 DPS) | yes | Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Runemender's Seal (273031, -0.33 DPS) [world]; Solid Gold Band (279845, -1.31 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 8.2 spell_power points (1.18 DPS) | yes | Cryptbone Staff (2013, -0.38 DPS) [world]; Staff of the Friar (3415, -0.38 DPS) [dungeon]; Twisted Chanter's Staff (890, -6.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Starfaller (13063, -0.89 DPS, sim-verified) [world_drop]; Gyromatic Icemaker (9489, -1.36 DPS) [dungeon]; Greater Mystic Wand (217287, -2.00 DPS) [crafted] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Glimmering Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (orc, 000000000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 144.1. Weights run: 1.4s. Verify run: 1.0s. 411 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.762 ± 0.025, crit=0.248 ± 0.018 per rating point (14 rating = 1%, 3.475 per %), hit=0.366 ± 0.031 per rating point (10 rating = 1%, 3.659 per %), spell_haste=not significant (0.229 ± 0.260), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.225 ± 0.166), mp5=not significant (0.468 ± 0.237)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 23.0 spell_power points (4.57 DPS) | yes | Augural Shroud (2620, -0.87 DPS) [world]; Whitemane's Chapeau (7720, -1.13 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -5.62 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.1 spell_power points (2.61 DPS) | yes | Triune Amulet (7722, -0.70 DPS) [dungeon]; Necklace of Calisea (1714, -1.23 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.70 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.9 spell_power points (3.36 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.21 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.8 spell_power points (1.95 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.29 DPS) [crafted]; Blackforge Cape (6424, -0.56 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 28.0 spell_power points (5.57 DPS) | yes | Cindercloth Robe (10042, -0.20 DPS) [crafted]; Robe of Power (7054, -0.61 DPS) [crafted]; Robe of the Magi (1716, -1.98 DPS, sim-verified) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (2.59 DPS) | yes | Spidertank Oilrag (9448, -0.52 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.58 DPS) [quest]; Condor Bracers (15864, -1.19 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.6 spell_power points (4.50 DPS) | yes | Red Mageweave Gloves (10018, -0.80 DPS) [crafted]; Crimson Silk Gloves (7064, -0.81 DPS) [crafted]; Fiery Handwraps (254025, -2.51 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 21.4 spell_power points (4.27 DPS) | yes | Gilded Cord (254037, -1.19 DPS) [crafted]; Scorching Sash (4117, -1.28 DPS) [quest]; Deathmage Sash (10771, -3.58 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.1 spell_power points (4.61 DPS) | yes | Flame Leggings (253991, -0.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.04 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.33 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.6 spell_power points (5.09 DPS) | yes | Gilded Slippers (254001, -2.46 DPS) [crafted]; Acidic Walkers (9454, -2.70 DPS) [dungeon]; Fiery Slippers (254005, -4.34 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (2.75 DPS) | yes | Ring of Power Regulation (273030, -1.27 DPS) [dungeon]; Reedknot Ring (9622, -1.36 DPS) [quest]; Voodoo Band (1996, -1.55 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.8 spell_power points (1.95 DPS) | yes | Ring of Power Regulation (273030, -0.48 DPS) [dungeon]; Reedknot Ring (9622, -0.56 DPS) [quest]; Voodoo Band (1996, -0.76 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (144.1 DPS) | yes | Illusionary Rod (7713, -2.66 DPS) [dungeon]; Glimmering Staff (249392, -2.90 DPS) [crafted]; Searing Golden Blade (12260, -16.84 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 203.1 spell_power points (40.42 DPS) | yes | Ragefire Wand (7513, -1.03 DPS) [quest]; Nether Force Wand (11263, -2.63 DPS) [quest]; Icefury Wand (7514, -2.78 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 252010000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 184.5. Weights run: 1.3s. Verify run: 1.2s. 527 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.708 ± 0.037, crit=0.357 ± 0.028 per rating point (14 rating = 1%, 4.991 per %), hit=0.536 ± 0.046 per rating point (10 rating = 1%, 5.360 per %), spell_haste=not significant (0.071 ± 0.368), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.991 ± 0.331), mp5=not significant (1.052 ± 0.342)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 40.0 spell_power points (7.79 DPS) | yes | Red Mageweave Headband (10033, -1.33 DPS) [crafted]; Gemburst Circlet (10751, -1.91 DPS, sim-verified) [quest]; Papal Fez (9431, -2.15 DPS) [dungeon] |
| neck | Mindburst Medallion (11196) | Broken Alliances [quest] | 15.2 spell_power points (2.96 DPS) | yes | Triune Amulet (7722, -0.05 DPS) [dungeon]; Horizon Choker (13085, -0.26 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -0.49 DPS) [dungeon] |
| shoulder | Netherflame Shoulders (254053) | Tailoring [crafted] | 28.6 spell_power points (5.57 DPS) | yes | Rotgrip Mantle (17732, -0.56 DPS) [dungeon]; Flameseer Mantle (11310, -0.92 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.56 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.2 spell_power points (4.71 DPS) | yes | Cindercloth Cloak (14044, -1.08 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.13 DPS) [quest]; Runecloth Cloak (13860, -1.86 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 38.2 spell_power points (7.45 DPS) | yes | Cindercloth Vest (14042, -0.65 DPS) [crafted]; Acumen Robes (17775, -0.99 DPS) [quest]; Ghostweave Vest (14141, -2.61 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 22.9 spell_power points (4.46 DPS) | yes | Nethergeld Cuffs (254061, -1.17 DPS) [crafted]; Bloodband Bracers (11469, -1.67 DPS) [quest]; Phoenix Bindings (210781, -1.93 DPS) [crafted] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 34.0 spell_power points (6.63 DPS) | yes | Dreamweave Gloves (10019, -1.22 DPS) [crafted]; Runecloth Gloves (13863, -1.31 DPS) [crafted]; Cindercloth Gloves (14043, -4.75 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 25.6 spell_power points (4.99 DPS) | yes | Fiery Waistcord (254085, -0.05 DPS) [crafted]; Ghostweave Cord (254073, -0.14 DPS) [crafted]; Scorching Sash (4117, -0.72 DPS) [quest] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 34.8 spell_power points (6.78 DPS) | yes | Cindercloth Leggings (12256, -0.18 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.24 DPS) [dungeon]; Dalewind Trousers (13008, -0.71 DPS) [world_drop] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.9 spell_power points (6.03 DPS) | yes | Fiery Sandals (254111, -0.89 DPS) [crafted]; Vinerot Sandals (17748, -1.69 DPS) [dungeon]; Fiery Slippers (254005, -1.76 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 16.1 spell_power points (3.14 DPS) | yes | Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Coldwater Ring (4550, -0.99 DPS) [quest]; Welken Ring (5011, -1.10 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (2.64 DPS) | yes | Band of the Unicorn (7553, -0.11 DPS) [world_drop]; Coldwater Ring (4550, -0.49 DPS) [quest]; Welken Ring (5011, -0.60 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (184.5 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (184.5 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (184.5 DPS) | yes | Spellshifter Rod (9527, -0.62 DPS) [quest]; Inventor's Focal Sword (17719, -1.66 DPS) [dungeon]; Spire of Hakkar (10844, -3.64 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 285.5 spell_power points (55.61 DPS) | yes | Noxious Shooter (17745, -3.37 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -7.19 DPS) [world_drop]; Lesser Eternal Wand (249232, -8.30 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Mindburst Medallion; shoulder: Netherflame Shoulders; back: Spritecaster Cape; wrist: Netherflame Cuffs; hands: Fiery Gloves; legs: Senior Designer's Pantaloons; finger1: Eye of Adaegus; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 527, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 252015000000200000-51521001230133031-003000000000000000)

Set DPS (verified): 343.1. Weights run: 1.5s. Verify run: 3.4s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.624 ± 0.047, crit=0.540 ± 0.040 per rating point (14 rating = 1%, 7.566 per %), hit=0.894 ± 0.074 per rating point (10 rating = 1%, 8.941 per %), spell_haste=-6.621 ± 0.737, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003, spirit=not significant (0.487 ± 0.608), mp5=not significant (0.452 ± 0.619)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (343.1 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Magister's Crown (16686, -7.55 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.2 spell_power points (4.48 DPS) | yes | Mindburst Medallion (11196, -1.49 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.53 DPS) [dungeon]; Triune Amulet (7722, -1.68 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 32.7 spell_power points (8.03 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -0.32 DPS) [pvp]; Shroud of the Nathrezim (18720, -1.04 DPS) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (5.16 DPS) | yes | Spritecaster Cape (11623, -0.08 DPS) [dungeon]; Cindercloth Cloak (14044, -0.74 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.83 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.1 spell_power points (13.52 DPS) | yes | Warlord's Silk Raiment (231596, -0.23 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.30 DPS) [pvp]; Robes of Fiery Devastation (279266, -8.87 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (343.1 DPS) | yes | Netherflame Cuffs (254065, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Magister's Bindings (16683, -6.07 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (343.1 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Magister's Gloves (16684, -3.60 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (343.1 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Fiery Waistcord (254085, +0.00 DPS) [crafted]; Magister's Belt (16685, -9.20 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (343.1 DPS) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (343.1 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, +0.00 DPS) [crafted]; Omnicast Boots (11822, -6.84 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.3 spell_power points (6.21 DPS) | yes | Band of the Unicorn (7553, -3.02 DPS) [world_drop]; Band of the Hierophant (13096, -3.36 DPS) [world_drop]; Eye of Adaegus (5266, -3.86 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.22 DPS) | yes | Band of the Unicorn (7553, -0.03 DPS) [world_drop]; Band of the Hierophant (13096, -0.37 DPS) [world_drop]; Eye of Adaegus (5266, -0.87 DPS) [world_drop] |
| trinket1 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (343.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (343.1 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -3.85 DPS, sim-verified) [world_drop] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (343.1 DPS) | yes | High Warlord's War Staff (234549, -1.41 DPS) [pvp]; Spire of Hakkar (10844, -11.41 DPS) [world]; Kindling Stave (11750, -17.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 323.9 spell_power points (79.56 DPS) | yes | Stormrager (16997, -8.31 DPS, sim-verified) [quest]; Brilliant Wand (249385, -16.69 DPS) [crafted]; Torch of Austen (13004, -22.77 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; trinket1: Evonice's Landin' Pilla; trinket2: Ankh of Life; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 253005000000200000-51521001230133031-003000000000000000)

Set DPS (verified): 699.4. Weights run: 1.6s. Verify run: 3.6s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.551 ± 0.039, crit=0.826 ± 0.046 per rating point (14 rating = 1%, 11.568 per %), hit=1.113 ± 0.075 per rating point (10 rating = 1%, 11.130 per %), spell_haste=-3.538 ± 0.639, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.002, spirit=not significant (0.784 ± 0.350), mp5=not significant (1.003 ± 0.408)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (699.4 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Dreadmist Mask (16698, -10.18 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.5 spell_power points (9.34 DPS) | yes | Mindburst Medallion (11196, -2.28 DPS) [quest]; Lady Alizabeth's Pendant (13002, -3.06 DPS) [world_drop]; Triune Amulet (7722, -6.72 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.7 spell_power points (19.06 DPS) | yes | Champion's Silk Mantle (227104, +0.00 DPS) [pvp]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.09 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.0 spell_power points (11.74 DPS) | yes | Mageflame Cloak (13007, -0.54 DPS) [world_drop]; Cindercloth Cloak (14044, -2.46 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.69 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.2 spell_power points (31.03 DPS) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.57 DPS) [pvp]; Robes of Fiery Devastation (279266, -16.94 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 20.8 spell_power points (11.08 DPS) | yes | General's Silk Cuffs (16538, -1.31 DPS) [pvp]; Runecloth Cuffs (254123, -2.28 DPS) [crafted]; Arena Wristguards (18709, -2.40 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.6 spell_power points (20.08 DPS) | yes | General's Silk Handguards (16540, -0.06 DPS) [vendor]; General's Silk Gauntlets (231599, -0.06 DPS) [vendor]; Fiery Gloves (254099, -7.99 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (699.4 DPS) | yes | General's Silk Sash (16537, +0.00 DPS) [pvp]; Fiery Waistcord (254085, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -8.87 DPS, sim-verified) [crafted] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 41.8 spell_power points (22.27 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -10.58 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 29.5 spell_power points (15.73 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.65 DPS) [crafted]; Wolfrunner Shoes (13101, -0.96 DPS) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 26.6 spell_power points (14.17 DPS) | yes | Philanthropist's Ring (281635, -7.36 DPS) [quest]; Eye of Adaegus (5266, -7.38 DPS) [world_drop]; Band of the Hierophant (13096, -9.70 DPS, sim-verified) [world_drop] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | sim-verified (699.4 DPS) | yes | Philanthropist's Ring (281635, -0.13 DPS) [quest]; Eye of Adaegus (5266, -0.15 DPS) [world_drop]; Band of the Hierophant (13096, -4.37 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (699.4 DPS) | yes | Ankh of Life (1713, -0.32 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.16 DPS) [quest]; Shard of the Splithooves (10659, -2.13 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (699.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (699.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -21.22 DPS) [world]; Kindling Stave (11750, -22.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 158.6 spell_power points (84.57 DPS) | yes | Brilliant Wand (249385, -18.13 DPS) [crafted]; Stormrager (16997, -20.40 DPS) [quest]; Pyric Caduceus (11748, -23.88 DPS) [dungeon] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

