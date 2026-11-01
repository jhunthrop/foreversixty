# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-51500000000000000-000000000000000000)

Set DPS (verified): 27.2. Weights run: 0.7s. Verify run: 0.5s. 160 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.354 ± 0.014, crit=0.146 ± 0.005 per rating point (14 rating = 1%, 2.047 per %), hit=0.328 ± 0.017 per rating point (10 rating = 1%, 3.280 per %), spell_haste=not significant (-0.176 ± 0.231), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.77 DPS) | yes | Pristine Circlet (253949, -0.35 DPS) [crafted]; Shadow Goggles (4373, -0.65 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.2 spell_power points (0.58 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.50 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.28 DPS) | yes | Pearl-clasped Cloak (5542, -0.07 DPS) [crafted]; Black Whelp Cloak (7283, -0.07 DPS) [crafted]; Sanguine Cape (14376, -0.18 DPS) [world_drop] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 10.8 spell_power points (0.76 DPS) | yes | Filigreed Pristine Gown (253901, -0.28 DPS) [crafted]; Gray Woolen Robe (2585, -0.35 DPS) [crafted]; Green Woolen Robe (6243, -0.47 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.8 spell_power points (0.12 DPS) | yes | Windsong Bangles (263336, -0.05 DPS) [quest]; Repurposed Hair Band (281256, -0.07 DPS) [quest]; Bright Bracers (3647, -0.50 DPS, sim-verified) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.63 DPS) | yes | Flame Gloves (253917, -0.07 DPS) [crafted]; Serpent Gloves (5970, -0.14 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.21 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.4 spell_power points (0.59 DPS) | yes | Pristine Sash (253925, -0.21 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.24 DPS) [crafted]; Novice Ardent's Sash (253887, -0.38 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.1 spell_power points (0.92 DPS) | yes | Filigreed Pristine Leggings (253937, -0.35 DPS) [crafted]; Silk-threaded Trousers (1929, -0.43 DPS) [dungeon]; Phoenix Pants (4317, -0.52 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.59 DPS) | yes | Flame Boots (253893, -0.10 DPS) [crafted]; Pristine Boots (253889, -0.31 DPS) [crafted]; Red Woolen Boots (4313, -0.31 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.40 DPS) | yes | Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; The 1 Ring (8350, -0.38 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.1 spell_power points (0.15 DPS) | yes | The 1 Ring (8350, -0.12 DPS) [world]; Volcanic Rock Ring (12053, -0.34 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.5 spell_power points (0.25 DPS) | yes | Witching Stave (1484, -0.12 DPS) [dungeon]; Staff of Westfall (2042, -0.12 DPS) [quest]; Lesser Staff of the Spire (1300, -0.30 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 320.2 spell_power points (22.52 DPS) | yes | Skycaller (12984, -0.94 DPS, sim-verified) [world_drop]; Deepblaze (279896, -3.99 DPS) [quest]; Sizzle Stick (8071, -4.52 DPS) [quest] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 160, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 000000000000000000-51521001230100000-000000000000000000)

Set DPS (verified): 52.5. Weights run: 0.8s. Verify run: 0.5s. 298 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.746 ± 0.014, crit=0.172 ± 0.010 per rating point (14 rating = 1%, 2.412 per %), hit=0.335 ± 0.019 per rating point (10 rating = 1%, 3.346 per %), spell_haste=not significant (-0.166 ± 0.134), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.44 DPS) | yes | Enchanter's Cowl (4322, -0.06 DPS) [crafted]; Holy Shroud (2721, -0.31 DPS) [world_drop]; Flame Circlet (253953, -0.31 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.5 spell_power points (0.46 DPS) | yes | Crystal Starfire Medallion (5003, -0.15 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.15 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.7 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.31 DPS) [quest]; Magician's Mantle (12998, -0.41 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.51 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Resilient Cape (14400, -0.05 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.08 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.7 spell_power points (1.92 DPS) | yes | Flame Gown (253965, -0.05 DPS) [crafted]; Beguiler Robes (7728, -0.08 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.33 DPS) | yes | Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.72 DPS) [world_drop]; Nightsky Wristbands (6407, -0.87 DPS) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.2 spell_power points (0.95 DPS) | yes | Phoenix Gloves (4331, -0.02 DPS) [crafted]; Town Clerk's Mittens (270029, -0.10 DPS) [quest]; Truefaith Gloves (7049, -0.21 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.2 spell_power points (1.15 DPS) | yes | Crimson Silk Belt (7055, -0.00 DPS) [crafted]; Invoker's Cord (215366, -0.05 DPS) [crafted]; Flame Sash (253929, -0.13 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 19.2 spell_power points (1.97 DPS) | yes | Filigreed Flame Leggings (253941, -0.52 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.64 DPS) [world]; Phoenix Pants (4317, -0.64 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 18.2 spell_power points (1.87 DPS) | yes | Gilded Slippers (254001, -0.62 DPS) [crafted]; Acidic Walkers (9454, -0.74 DPS) [dungeon]; Spidersilk Boots (4320, -0.85 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (1.00 DPS) | yes | Ring of Power Regulation (273030, -0.38 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon]; Azora's Will (4999, -0.62 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (0.67 DPS) | yes | Ring of Power Regulation (273030, -0.05 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Azora's Will (4999, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.69 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.64 DPS) [quest]; Glimmering Staff (249392, -2.85 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 323.2 spell_power points (33.17 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.65 DPS) [crafted]; Scorching Wand (5213, -3.70 DPS) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Flame Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Minor Channeling Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 000000000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 90.1. Weights run: 0.9s. Verify run: 0.7s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.776 ± 0.024, crit=0.213 ± 0.019 per rating point (14 rating = 1%, 2.982 per %), hit=0.387 ± 0.041 per rating point (10 rating = 1%, 3.866 per %), spell_haste=not significant (0.360 ± 0.302), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 21.2 spell_power points (2.87 DPS) | yes | Augural Shroud (2620, -0.33 DPS) [world]; Whitemane's Chapeau (7720, -0.87 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.93 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.2 spell_power points (1.79 DPS) | yes | Necklace of Calisea (1714, -1.05 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.16 DPS) [world]; Triune Amulet (7722, -1.95 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.1 spell_power points (2.31 DPS) | yes | Bloodmage Mantle (7684, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest]; Green Silken Shoulders (7057, -1.22 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.9 spell_power points (1.34 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.39 DPS) [crafted]; Blackforge Cape (6424, -0.50 DPS) [dungeon] |
| chest | Cindercloth Robe (10042) | Tailoring [crafted] | 27.0 spell_power points (3.66 DPS) | yes | Robe of the Magi (1716, -0.05 DPS) [world_drop]; Dreamweave Vest (10021, -0.27 DPS) [crafted]; Robe of Power (7054, -0.50 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.76 DPS) | yes | Arcane Runed Bracers (4744, -0.54 DPS) [quest]; Spidertank Oilrag (9448, -0.54 DPS) [dungeon]; Condor Bracers (15864, -0.81 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.86 DPS) | yes | Fiery Handwraps (254025, -0.23 DPS) [crafted]; Red Mageweave Gloves (10018, -0.32 DPS) [crafted]; Crimson Silk Gloves (7064, -0.33 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (2.74 DPS) | yes | Deathmage Sash (10771, -0.21 DPS) [dungeon]; Gilded Cord (254037, -0.81 DPS) [crafted]; Scorching Sash (4117, -0.98 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.3 spell_power points (3.16 DPS) | yes | Crimson Silk Pantaloons (7062, -0.71 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.04 DPS) [crafted]; Flame Leggings (253991, -1.31 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.25 DPS) | yes | Fiery Slippers (254005, -1.00 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.57 DPS) [crafted]; Acidic Walkers (9454, -1.73 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (1.88 DPS) | yes | Ring of Forlorn Spirits (2043, -0.80 DPS) [quest]; Reedknot Ring (9622, -0.93 DPS) [quest]; Minor Channeling Ring (1449, -0.99 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.9 spell_power points (1.34 DPS) | yes | Ring of Forlorn Spirits (2043, -0.25 DPS) [quest]; Reedknot Ring (9622, -0.39 DPS) [quest]; Minor Channeling Ring (1449, -0.45 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (90.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.21 DPS) [dungeon]; Searing Golden Blade (12260, -0.78 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.3 spell_power points (2.08 DPS) | yes | Orb of the Forgotten Seer (7685, -0.18 DPS) [dungeon]; Orb of Lorica (11262, -0.21 DPS) [quest]; Rod of Molten Fire (2565, -0.32 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 295.4 spell_power points (40.01 DPS) | yes | Ragefire Wand (7513, -1.34 DPS) [quest]; Nether Force Wand (11263, -2.37 DPS) [quest]; Icefury Wand (7514, -2.51 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Cindercloth Robe; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 252010000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 148.1. Weights run: 0.9s. Verify run: 0.8s. 512 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.900 ± 0.033, crit=0.361 ± 0.029 per rating point (14 rating = 1%, 5.055 per %), hit=0.585 ± 0.053 per rating point (10 rating = 1%, 5.854 per %), spell_haste=not significant (1.066 ± 0.517), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.0 spell_power points (5.07 DPS) | yes | Dreamweave Circlet (10041, -0.96 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -0.98 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.37 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.2 spell_power points (1.94 DPS) | yes | Horizon Choker (13085, -0.22 DPS) [world_drop]; Mindburst Medallion (11196, -0.38 DPS) [quest]; Triune Amulet (7722, -0.85 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.2 spell_power points (4.00 DPS) | yes | Netherflame Shoulders (254053, -0.68 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.10 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.19 DPS) [crafted] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.2 spell_power points (2.77 DPS) | yes | Spritecaster Cape (11623, -0.11 DPS) [dungeon]; Runecloth Cloak (13860, -0.55 DPS) [crafted]; Cloak of Fire (14134, -0.85 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.0 spell_power points (5.07 DPS) | yes | Cindercloth Robe (10042, -1.37 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.38 DPS) [vendor]; Robe of the Magi (1716, -2.07 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.3 spell_power points (2.64 DPS) | yes | Nethergeld Cuffs (254061, -0.82 DPS) [crafted]; Bloodband Bracers (11469, -0.85 DPS) [quest]; Phoenix Bindings (210781, -0.86 DPS) [crafted] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 28.0 spell_power points (3.83 DPS) | yes | Dreamweave Gloves (10019, -0.88 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.94 DPS) [vendor]; Fiery Handwraps (254025, -1.05 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.1 spell_power points (3.71 DPS) | yes | Dawnspire Cord (12466, -0.55 DPS) [dungeon]; Satyrmane Sash (17755, -0.56 DPS) [dungeon]; Fiery Cord (254041, -0.81 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 27.9 spell_power points (3.81 DPS) | yes | Red Mageweave Pants (10009, -0.42 DPS) [crafted]; Flame Leggings (253991, -1.03 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.12 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.1 spell_power points (3.85 DPS) | yes | Earthen Silk Slippers (254013, -0.56 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -0.84 DPS) [vendor]; Cindercloth Boots (10044, -0.97 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.99 DPS) | yes | Snake Hoop (6750, -0.55 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.89 DPS) [quest]; Reedknot Ring (9622, -1.03 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.78 DPS) | yes | Snake Hoop (6750, -0.34 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.68 DPS) [quest]; Reedknot Ring (9622, -0.82 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (148.1 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (148.1 DPS) | yes | Mark of the Chosen (17774, -2.29 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (148.1 DPS) | yes | Spellshifter Rod (9527, -0.19 DPS) [quest]; Kindling Stave (11750, -0.73 DPS) [dungeon]; Hanzo Sword (8190, -24.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 383.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.82 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.65 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Knight's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 252015000000200000-51521001230133031-003000000000000000)

Set DPS (verified): 424.1. Weights run: 1.0s. Verify run: 2.6s. 1131 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.779 ± 0.044, crit=0.521 ± 0.039 per rating point (14 rating = 1%, 7.300 per %), hit=1.106 ± 0.076 per rating point (10 rating = 1%, 11.059 per %), spell_haste=4.256 ± 0.644, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (424.1 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -9.43 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 19.8 spell_power points (6.54 DPS) | yes | Glowing Eye of Mordresh (10769, -2.17 DPS) [dungeon]; Horizon Choker (13085, -2.93 DPS) [world_drop]; Mindburst Medallion (11196, -3.01 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.4 spell_power points (11.37 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.18 DPS) [pvp]; Rotgrip Mantle (17732, -9.01 DPS, sim-verified) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (6.94 DPS) | yes | Cindercloth Cloak (14044, -0.58 DPS) [crafted]; Spritecaster Cape (11623, -0.77 DPS) [dungeon]; Runecloth Cloak (13860, -1.90 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.7 spell_power points (18.71 DPS) | yes | Field Marshal's Silk Vestments (231603, -1.02 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.99 DPS) [pvp]; Robes of Fiery Devastation (279266, -8.82 DPS, sim-verified) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 18.8 spell_power points (6.21 DPS) | yes | Netherflame Cuffs (254065, -0.11 DPS) [crafted]; Marshal's Silk Bracers (16438, -1.32 DPS) [pvp]; Dreadmist Bracers (16703, -1.72 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.4 spell_power points (10.37 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Fiery Gloves (254099, -9.39 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (424.1 DPS) | yes | Wisdom of the Timbermaw (19047, -0.48 DPS) [crafted]; Satyrmane Sash (17755, -0.52 DPS) [dungeon]; Fiery Waistcord (254085, -10.03 DPS, sim-verified) [crafted] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (424.1 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -15.43 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.4 spell_power points (9.69 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.18 DPS) [pvp]; Fiery Sandals (254111, -0.77 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.5 spell_power points (7.75 DPS) | yes | Band of the Unicorn (7553, -3.45 DPS) [world_drop]; Snake Hoop (6750, -4.48 DPS) [quest]; Glowing Crystal Ring (18402, -4.66 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (4.59 DPS) | yes | Band of the Unicorn (7553, -0.30 DPS) [world_drop]; Snake Hoop (6750, -1.32 DPS) [quest]; Glowing Crystal Ring (18402, -1.50 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (424.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (424.1 DPS) | yes | Mark of the Chosen (17774, -10.40 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (424.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.22 DPS) [quest]; Teebu's Blazing Longsword (1728, -65.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 238.4 spell_power points (78.74 DPS) | yes | Brilliant Wand (249385, -14.39 DPS) [crafted]; Stormrager (16997, -14.76 DPS) [quest]; Torch of Austen (13004, -21.95 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1131, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 251015000000000000-54521001230133031-003000000000000000)

Set DPS (verified): 750.3. Weights run: 1.1s. Verify run: 1.6s. 1131 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.457 ± 0.028, crit=0.885 ± 0.043 per rating point (14 rating = 1%, 12.385 per %), hit=1.147 ± 0.077 per rating point (10 rating = 1%, 11.469 per %), spell_haste=3.512 ± 0.408, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 28.1 spell_power points (16.26 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, +0.00 DPS) [crafted]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.6 spell_power points (9.57 DPS) | yes | Glowing Eye of Mordresh (10769, -3.42 DPS) [dungeon]; Mindburst Medallion (11196, -4.52 DPS) [quest]; Triune Amulet (7722, -4.95 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.3 spell_power points (20.41 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -1.68 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -2.01 DPS) [pvp]; Argent Shoulders (19059, -11.19 DPS, sim-verified) [crafted] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (12.13 DPS) | yes | Spritecaster Cape (11623, -2.46 DPS) [dungeon]; Cindercloth Cloak (14044, -2.51 DPS) [crafted]; Cloak of Fire (14134, -4.04 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 57.9 spell_power points (33.44 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.72 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.66 DPS) [pvp]; Robes of Fiery Devastation (279266, -20.65 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 16.2 spell_power points (9.36 DPS) | yes | Runecloth Cuffs (254123, -0.36 DPS) [crafted]; Phoenix Bindings (210781, -1.85 DPS) [crafted]; Arena Wristguards (18709, -2.20 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 38.3 spell_power points (22.15 DPS) | yes | Marshal's Silk Gloves (16440, -3.38 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -3.38 DPS) [vendor]; Fiery Gloves (254099, -15.08 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 23.1 spell_power points (13.35 DPS) | yes | Satyrmane Sash (17755, -2.62 DPS) [dungeon]; Fiery Cord (254041, -3.15 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.30 DPS) [crafted] |
| legs | Cindercloth Pants (14045) | Tailoring [crafted] | 31.5 spell_power points (18.19 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -0.93 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.5 spell_power points (14.72 DPS) | yes | Knight-Lieutenant's Silk Walkers (227112, +0.00 DPS) [pvp]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Fiery Sandals (254111, -0.79 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.2 spell_power points (12.25 DPS) | yes | Philanthropist's Ring (281635, -5.15 DPS) [quest]; Snake Hoop (6750, -7.46 DPS) [quest]; Ring of Forlorn Spirits (2043, -7.63 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.51 DPS) | yes | Philanthropist's Ring (281635, -0.41 DPS) [quest]; Snake Hoop (6750, -2.72 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.89 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (750.3 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (750.3 DPS) | yes | Mark of the Chosen (17774, -7.88 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (750.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.56 DPS) [quest]; Teebu's Blazing Longsword (1728, -57.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 138.9 spell_power points (80.23 DPS) | yes | Brilliant Wand (249385, -13.60 DPS) [crafted]; Stormrager (16997, -16.22 DPS) [quest]; Greater Eternal Wand (249237, -21.70 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; hands: Gloves of Spell Mastery; legs: Cindercloth Pants; feet: Omnicast Boots; finger1: Maiden's Circle; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1131, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 000000000000000000-51500000000000000-000000000000000000)

Set DPS (verified): 25.2. Weights run: 0.7s. Verify run: 0.6s. 154 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.354 ± 0.014, crit=0.146 ± 0.005 per rating point (14 rating = 1%, 2.047 per %), hit=0.328 ± 0.017 per rating point (10 rating = 1%, 3.280 per %), spell_haste=not significant (-0.176 ± 0.231), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.77 DPS) | yes | Pristine Circlet (253949, -0.35 DPS) [crafted]; Shadow Goggles (4373, -0.65 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.2 spell_power points (0.58 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.30 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.28 DPS) | yes | Pearl-clasped Cloak (5542, -0.07 DPS) [crafted]; Black Whelp Cloak (7283, -0.07 DPS) [crafted]; Mystic Shawl (3449, -0.14 DPS) [quest] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 10.8 spell_power points (0.76 DPS) | yes | Filigreed Pristine Gown (253901, -0.28 DPS) [crafted]; Gray Woolen Robe (2585, -0.35 DPS) [crafted]; Green Woolen Robe (6243, -0.47 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.21 DPS) | yes | Mindthrust Bracers (1974, -0.09 DPS) [dungeon]; Featherbead Bracers (15452, -0.09 DPS) [quest]; Owlbeard Bracers (16981, -0.09 DPS) [quest] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.63 DPS) | yes | Flame Gloves (253917, -0.07 DPS) [crafted]; Serpent Gloves (5970, -0.14 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.21 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.4 spell_power points (0.59 DPS) | yes | Pristine Sash (253925, -0.21 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.24 DPS) [crafted]; Novice Ardent's Sash (253887, -0.38 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.1 spell_power points (0.92 DPS) | yes | Phoenix Pants (4317, -0.12 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.35 DPS) [crafted]; Silk-threaded Trousers (1929, -0.43 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.59 DPS) | yes | Flame Boots (253893, -0.10 DPS) [crafted]; Pristine Boots (253889, -0.31 DPS) [crafted]; Red Woolen Boots (4313, -0.31 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.1 spell_power points (0.15 DPS) | yes | Volcanic Rock Ring (12053, -0.07 DPS) [world_drop]; The 1 Ring (8350, -0.12 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 1.8 spell_power points (0.12 DPS) | yes | Volcanic Rock Ring (12053, -0.05 DPS) [world_drop]; The 1 Ring (8350, -0.10 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.5 spell_power points (0.25 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.10 DPS) [world]; Clear Crystal Rod (16894, -0.12 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 320.2 spell_power points (22.52 DPS) | yes | Skycaller (12984, -0.86 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.52 DPS) [quest]; Greater Magic Wand (11288, -4.88 DPS) [crafted] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Tabitha's Cuffs; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 154, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 000000000000000000-51521001230100000-000000000000000000)

Set DPS (verified): 51.3. Weights run: 0.8s. Verify run: 0.6s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.746 ± 0.014, crit=0.172 ± 0.010 per rating point (14 rating = 1%, 2.412 per %), hit=0.335 ± 0.019 per rating point (10 rating = 1%, 3.346 per %), spell_haste=not significant (-0.166 ± 0.134), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.44 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Holy Shroud (2721, -0.31 DPS) [world_drop]; Flame Circlet (253953, -0.31 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.5 spell_power points (0.46 DPS) | yes | Crystal Starfire Medallion (5003, -0.15 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.15 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.7 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.31 DPS) [quest]; Magician's Mantle (12998, -0.41 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.51 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Resilient Cape (14400, -0.05 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.08 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.7 spell_power points (1.92 DPS) | yes | Flame Gown (253965, -0.05 DPS) [crafted]; Beguiler Robes (7728, -0.08 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (51.3 DPS) | yes | Glowing Magical Bracelets (13106, -0.31 DPS) [world_drop]; Nightsky Wristbands (6407, -0.46 DPS) [world_drop]; Phoenix Bindings (210781, -0.53 DPS, sim-verified) [crafted] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.7 spell_power points (1.00 DPS) | yes | Flame Gloves (253917, -0.05 DPS) [crafted]; Phoenix Gloves (4331, -0.08 DPS) [crafted]; Truefaith Gloves (7049, -0.26 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.2 spell_power points (1.15 DPS) | yes | Crimson Silk Belt (7055, -0.00 DPS) [crafted]; Warsong Sash (16975, -0.02 DPS) [quest]; Invoker's Cord (215366, -0.05 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 19.2 spell_power points (1.97 DPS) | yes | Filigreed Flame Leggings (253941, -0.38 DPS) [crafted]; Smoldering Pants (3073, -0.64 DPS) [world]; Phoenix Pants (4317, -0.64 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 18.2 spell_power points (1.87 DPS) | yes | Gilded Slippers (254001, -0.62 DPS) [crafted]; Acidic Walkers (9454, -0.74 DPS) [dungeon]; Spidersilk Boots (4320, -0.85 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon]; Solid Gold Band (279845, -0.59 DPS) [quest]; Azora's Will (4999, -0.62 DPS) [world_drop] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.62 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Solid Gold Band (279845, -0.21 DPS) [quest]; Azora's Will (4999, -0.23 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.69 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -2.85 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -2.93 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 323.2 spell_power points (33.17 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Unstable Power Core (279847, -3.46 DPS) [quest]; Greater Mystic Wand (217287, -3.65 DPS) [crafted] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (orc, 000000000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 85.5. Weights run: 0.9s. Verify run: 0.7s. 386 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.776 ± 0.024, crit=0.213 ± 0.019 per rating point (14 rating = 1%, 2.982 per %), hit=0.387 ± 0.041 per rating point (10 rating = 1%, 3.866 per %), spell_haste=not significant (0.360 ± 0.302), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 21.2 spell_power points (2.87 DPS) | yes | Augural Shroud (2620, -0.33 DPS) [world]; Whitemane's Chapeau (7720, -0.87 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.22 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.2 spell_power points (1.79 DPS) | yes | Necklace of Calisea (1714, -1.05 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.16 DPS) [world]; Triune Amulet (7722, -1.41 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.1 spell_power points (2.31 DPS) | yes | Bloodmage Mantle (7684, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest]; Green Silken Shoulders (7057, -1.23 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.9 spell_power points (1.34 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.39 DPS) [crafted]; Blackforge Cape (6424, -0.50 DPS) [dungeon] |
| chest | Cindercloth Robe (10042) | Tailoring [crafted] | 27.0 spell_power points (3.66 DPS) | yes | Robe of the Magi (1716, -0.05 DPS) [world_drop]; Dreamweave Vest (10021, -0.27 DPS) [crafted]; Robe of Power (7054, -0.50 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.76 DPS) | yes | Radiant Silver Bracers (4545, -0.38 DPS) [quest]; Spidertank Oilrag (9448, -0.54 DPS) [dungeon]; Condor Bracers (15864, -0.81 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.86 DPS) | yes | Fiery Handwraps (254025, -0.23 DPS) [crafted]; Red Mageweave Gloves (10018, -0.32 DPS) [crafted]; Crimson Silk Gloves (7064, -0.33 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (2.74 DPS) | yes | Deathmage Sash (10771, -0.21 DPS) [dungeon]; Gilded Cord (254037, -0.81 DPS) [crafted]; Scorching Sash (4117, -0.98 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.3 spell_power points (3.16 DPS) | yes | Flame Leggings (253991, -0.53 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.71 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.04 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.25 DPS) | yes | Fiery Slippers (254005, -0.75 DPS) [crafted]; Gilded Slippers (254001, -1.57 DPS) [crafted]; Acidic Walkers (9454, -1.73 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (1.88 DPS) | yes | Reedknot Ring (9622, -0.93 DPS) [quest]; Ring of Power Regulation (273030, -1.07 DPS) [dungeon]; Voodoo Band (1996, -1.14 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.9 spell_power points (1.34 DPS) | yes | Reedknot Ring (9622, -0.39 DPS) [quest]; Ring of Power Regulation (273030, -0.53 DPS) [dungeon]; Voodoo Band (1996, -0.60 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (85.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.21 DPS) [dungeon]; Searing Golden Blade (12260, -0.78 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.3 spell_power points (2.08 DPS) | yes | Orb of the Forgotten Seer (7685, -0.18 DPS) [dungeon]; Rod of Molten Fire (2565, -0.32 DPS) [world_drop]; Orb of Mystic Insight (249394, -0.50 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 295.4 spell_power points (40.01 DPS) | yes | Ragefire Wand (7513, -1.34 DPS) [quest]; Nether Force Wand (11263, -2.37 DPS) [quest]; Icefury Wand (7514, -2.51 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Cindercloth Robe; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 386, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 252010000000000000-51521001230133031-000000000000000000)

Set DPS (verified): 141.9. Weights run: 0.9s. Verify run: 0.8s. 502 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.900 ± 0.033, crit=0.361 ± 0.029 per rating point (14 rating = 1%, 5.055 per %), hit=0.585 ± 0.053 per rating point (10 rating = 1%, 5.854 per %), spell_haste=not significant (1.066 ± 0.517), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.0 spell_power points (5.07 DPS) | yes | Dreamweave Circlet (10041, -0.96 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -0.98 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.37 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.2 spell_power points (1.94 DPS) | yes | Horizon Choker (13085, -0.22 DPS) [world_drop]; Mindburst Medallion (11196, -0.38 DPS) [quest]; Triune Amulet (7722, -0.85 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.2 spell_power points (4.00 DPS) | yes | Netherflame Shoulders (254053, -0.68 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.10 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.19 DPS) [crafted] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.2 spell_power points (2.77 DPS) | yes | Deep Woodlands Cloak (19121, -0.01 DPS) [quest]; Spritecaster Cape (11623, -0.11 DPS) [dungeon]; Runecloth Cloak (13860, -0.55 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.0 spell_power points (5.07 DPS) | yes | Cindercloth Robe (10042, -1.37 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.38 DPS) [vendor]; Robe of the Magi (1716, -2.09 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.3 spell_power points (2.64 DPS) | yes | Nethergeld Cuffs (254061, -0.82 DPS) [crafted]; Bloodband Bracers (11469, -0.85 DPS) [quest]; Phoenix Bindings (210781, -0.86 DPS) [crafted] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 28.0 spell_power points (3.83 DPS) | yes | Dreamweave Gloves (10019, -0.88 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -0.94 DPS) [vendor]; Fiery Handwraps (254025, -1.05 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 27.1 spell_power points (3.71 DPS) | yes | Dawnspire Cord (12466, -0.55 DPS) [dungeon]; Satyrmane Sash (17755, -0.56 DPS) [dungeon]; Fiery Cord (254041, -0.81 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 27.9 spell_power points (3.81 DPS) | yes | Red Mageweave Pants (10009, -0.42 DPS) [crafted]; Flame Leggings (253991, -1.03 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.12 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.1 spell_power points (3.85 DPS) | yes | Earthen Silk Slippers (254013, -0.56 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -0.84 DPS) [vendor]; Cindercloth Boots (10044, -0.97 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.99 DPS) | yes | Snake Hoop (6750, -0.55 DPS) [quest]; Reedknot Ring (9622, -1.03 DPS) [quest]; Voodoo Band (1996, -1.12 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.78 DPS) | yes | Snake Hoop (6750, -0.34 DPS) [quest]; Reedknot Ring (9622, -0.82 DPS) [quest]; Voodoo Band (1996, -0.92 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (141.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (141.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Spellshifter Rod (9527, -0.19 DPS) [quest]; Kindling Stave (11750, -0.73 DPS) [dungeon]; Hanzo Sword (8190, -23.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 383.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.82 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.65 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Stone Guard's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 502, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 252015000000200000-51521001230133031-003000000000000000)

Set DPS (verified): 407.1. Weights run: 1.0s. Verify run: 2.8s. 1123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.779 ± 0.044, crit=0.521 ± 0.039 per rating point (14 rating = 1%, 7.300 per %), hit=1.106 ± 0.076 per rating point (10 rating = 1%, 11.059 per %), spell_haste=4.256 ± 0.644, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (407.1 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -9.06 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 19.8 spell_power points (6.54 DPS) | yes | Glowing Eye of Mordresh (10769, -2.17 DPS) [dungeon]; Horizon Choker (13085, -2.93 DPS) [world_drop]; Mindburst Medallion (11196, -3.01 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.4 spell_power points (11.37 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.18 DPS) [pvp]; Rotgrip Mantle (17732, -9.12 DPS, sim-verified) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (6.94 DPS) | yes | Cindercloth Cloak (14044, -0.58 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.66 DPS) [quest]; Spritecaster Cape (11623, -0.77 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.7 spell_power points (18.71 DPS) | yes | Warlord's Silk Raiment (231596, -1.02 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.99 DPS) [pvp]; Robes of Fiery Devastation (279266, -8.17 DPS, sim-verified) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 18.8 spell_power points (6.21 DPS) | yes | Netherflame Cuffs (254065, -0.11 DPS) [crafted]; General's Silk Cuffs (16538, -1.32 DPS) [pvp]; Dreadmist Bracers (16703, -1.72 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.4 spell_power points (10.37 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Fiery Gloves (254099, -10.22 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (407.1 DPS) | yes | Fiery Waistcord (254085, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -0.48 DPS) [crafted]; Satyrmane Sash (17755, -0.52 DPS) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (407.1 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -12.99 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.4 spell_power points (9.69 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.18 DPS) [pvp]; Fiery Sandals (254111, -0.77 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.5 spell_power points (7.75 DPS) | yes | Band of the Unicorn (7553, -3.45 DPS) [world_drop]; Snake Hoop (6750, -4.48 DPS) [quest]; Glowing Crystal Ring (18402, -4.66 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (4.59 DPS) | yes | Band of the Unicorn (7553, -0.30 DPS) [world_drop]; Snake Hoop (6750, -1.32 DPS) [quest]; Glowing Crystal Ring (18402, -1.50 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (407.1 DPS) | yes | Uther's Strength (11302, -1.31 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (407.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.22 DPS) [quest]; Teebu's Blazing Longsword (1728, -58.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 238.4 spell_power points (78.74 DPS) | yes | Brilliant Wand (249385, -14.39 DPS) [crafted]; Stormrager (16997, -14.76 DPS) [quest]; Torch of Austen (13004, -21.95 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1123, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 251015000000000000-54521001230133031-003000000000000000)

Set DPS (verified): 732.7. Weights run: 1.1s. Verify run: 1.6s. 1123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.457 ± 0.028, crit=0.885 ± 0.043 per rating point (14 rating = 1%, 12.385 per %), hit=1.147 ± 0.077 per rating point (10 rating = 1%, 11.469 per %), spell_haste=3.512 ± 0.408, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 28.1 spell_power points (16.26 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, +0.00 DPS) [crafted]; Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.6 spell_power points (9.57 DPS) | yes | Glowing Eye of Mordresh (10769, -3.42 DPS) [dungeon]; Mindburst Medallion (11196, -4.52 DPS) [quest]; Triune Amulet (7722, -4.95 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.3 spell_power points (20.41 DPS) | yes | Champion's Silk Mantle (227104, -1.68 DPS) [pvp]; Warlord's Silk Amice (231594, -2.01 DPS) [pvp]; Argent Shoulders (19059, -14.65 DPS, sim-verified) [crafted] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (12.13 DPS) | yes | Spritecaster Cape (11623, -2.46 DPS) [dungeon]; Cindercloth Cloak (14044, -2.51 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.82 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 57.9 spell_power points (33.44 DPS) | yes | Warlord's Silk Raiment (231596, -2.72 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -9.66 DPS) [pvp]; Robes of Fiery Devastation (279266, -20.02 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 16.2 spell_power points (9.36 DPS) | yes | Runecloth Cuffs (254123, -0.36 DPS) [crafted]; Phoenix Bindings (210781, -1.85 DPS) [crafted]; Arena Wristguards (18709, -2.20 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 38.3 spell_power points (22.15 DPS) | yes | General's Silk Handguards (16540, -3.38 DPS) [vendor]; General's Silk Gauntlets (231599, -3.38 DPS) [vendor]; Fiery Gloves (254099, -22.06 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 23.1 spell_power points (13.35 DPS) | yes | Satyrmane Sash (17755, -2.62 DPS) [dungeon]; Fiery Cord (254041, -3.15 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.30 DPS) [crafted] |
| legs | Cindercloth Pants (14045) | Tailoring [crafted] | 31.5 spell_power points (18.19 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Stone Guard's Dreadweave Leggings (220906, -0.93 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.5 spell_power points (14.72 DPS) | yes | Blood Guard's Silk Walkers (227110, +0.00 DPS) [pvp]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Fiery Sandals (254111, -0.79 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.2 spell_power points (12.25 DPS) | yes | Philanthropist's Ring (281635, -5.15 DPS) [quest]; Snake Hoop (6750, -7.46 DPS) [quest]; Reedknot Ring (9622, -8.20 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.51 DPS) | yes | Philanthropist's Ring (281635, -0.41 DPS) [quest]; Snake Hoop (6750, -2.72 DPS) [quest]; Reedknot Ring (9622, -3.47 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (732.7 DPS) | yes | Uther's Strength (11302, -2.50 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (732.7 DPS) | yes | Uther's Strength (11302, -1.17 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (732.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.56 DPS) [quest]; Teebu's Blazing Longsword (1728, -58.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 138.9 spell_power points (80.23 DPS) | yes | Brilliant Wand (249385, -13.60 DPS) [crafted]; Stormrager (16997, -16.22 DPS) [quest]; Greater Eternal Wand (249237, -21.70 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; hands: Gloves of Spell Mastery; legs: Cindercloth Pants; feet: Omnicast Boots; finger1: Maiden's Circle; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1123, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

