# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 28.2. Weights run: 1.2s. Verify run: 1.0s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.342 ± 0.011, crit=0.138 ± 0.004 per rating point (14 rating = 1%, 1.931 per %), hit=0.427 ± 0.013 per rating point (10 rating = 1%, 4.273 per %), spell_haste=0.549 ± 0.129, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.71 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Shadow Goggles (4373, -0.60 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.52 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.11 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Feyscale Cloak (6632, -0.06 DPS) [dungeon]; Black Whelp Cloak (7283, -0.06 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 10.7 spell_power points (0.69 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.32 DPS) [crafted]; Green Woolen Robe (6243, -0.43 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.7 spell_power points (0.11 DPS) | yes | Windsong Bangles (263336, -0.05 DPS) [quest]; Repurposed Hair Band (281256, -0.07 DPS) [quest]; Bright Bracers (3647, -0.56 DPS, sim-verified) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.58 DPS) | yes | Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world]; Flame Gloves (253917, -0.22 DPS, sim-verified) [crafted] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.4 spell_power points (0.54 DPS) | yes | Pristine Sash (253925, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Novice Ardent's Sash (253887, -0.35 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.1 spell_power points (0.84 DPS) | yes | Phoenix Pants (4317, -0.11 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted]; Abomination Skin Leggings (23173, -0.37 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.54 DPS) | yes | Feather Padded Treads (285345, -0.22 DPS) [world]; Pristine Boots (253889, -0.28 DPS) [crafted]; Flame Boots (253893, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.37 DPS) | yes | Sludge-Stained Band (286535, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.30 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.32 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop]; Sludge-Stained Band (286535, -0.50 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.4 spell_power points (0.22 DPS) | yes | Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world]; Staff of Westfall (2042, -0.11 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 347.7 spell_power points (22.50 DPS) | yes | Skycaller (12984, -0.95 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.21 DPS) [dungeon]; Deepblaze (279896, -3.97 DPS) [quest] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 53.0. Weights run: 1.3s. Verify run: 1.1s. 266 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.705 ± 0.013, crit=0.174 ± 0.008 per rating point (14 rating = 1%, 2.429 per %), hit=0.356 ± 0.017 per rating point (10 rating = 1%, 3.560 per %), spell_haste=0.908 ± 0.190, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.32 DPS) | yes | Enchanter's Cowl (4322, -0.09 DPS) [crafted]; Holy Shroud (2721, -0.28 DPS) [world_drop]; Flame Circlet (253953, -0.28 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 spell_power points (1.06 DPS) | yes | Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.93 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (1.45 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Cloak of Rot (4462, -0.02 DPS) [world]; Darkspear Raider's Cloak (272078, -0.02 DPS) [vendor]; Vine Pruner's Cloak (279835, -0.51 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.2 spell_power points (1.71 DPS) | yes | Flame Gown (253965, -0.02 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon]; Death Speaker Robes (6682, -0.32 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.23 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -0.69 DPS) [world_drop]; Nightsky Wristbands (6407, -0.83 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.8 spell_power points (1.11 DPS) | yes | Phoenix Gloves (4331, -0.26 DPS) [crafted]; Truefaith Gloves (7049, -0.44 DPS) [crafted]; Flame Gloves (253917, -0.63 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.1 spell_power points (1.24 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Crimson Silk Belt (7055, -0.21 DPS) [crafted]; Invoker's Cord (215366, -0.24 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.9 spell_power points (1.78 DPS) | yes | Abomination Skin Leggings (23173, -0.40 DPS) [dungeon]; Filigreed Flame Leggings (253941, -0.52 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.56 DPS) [world] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.9 spell_power points (1.69 DPS) | yes | Gilded Slippers (254001, -0.57 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.19 DPS) [world]; Snake Hoop (6750, -0.19 DPS) [quest]; Minor Channeling Ring (1449, -1.32 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.10 DPS) [world]; Snake Hoop (6750, -0.10 DPS) [quest]; Minor Channeling Ring (1449, -0.59 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Hardened Root Staff (1317, -1.39 DPS, sim-verified) [quest]; Scorn's Focal Dagger (23168, -1.65 DPS) [dungeon]; Glimmering Staff (249392, -1.77 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 355.4 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.38 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Scorching Wand (5213, -4.08 DPS) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Flame Leggings; feet: Fiery Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 266, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-23552100130103041-0000000000000000000)

Set DPS (verified): 82.3. Weights run: 1.5s. Verify run: 1.3s. 346 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.774 ± 0.026, crit=0.219 ± 0.015 per rating point (14 rating = 1%, 3.059 per %), hit=0.376 ± 0.034 per rating point (10 rating = 1%, 3.765 per %), spell_haste=2.538 ± 0.457, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.59 DPS) | yes | Augural Shroud (2620, -0.28 DPS) [world]; Corpseshroud (10574, -0.78 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.84 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.2 spell_power points (1.63 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.67 DPS) [quest]; Scorn's Icy Choker (23169, -0.72 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.1 spell_power points (2.11 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.0 spell_power points (1.97 DPS) | yes | Guardian Cloak (5965, -0.75 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.92 DPS) [vendor]; Long Silken Cloak (4326, -1.21 DPS, sim-verified) [crafted] |
| chest | Cindercloth Robe (10042) | Tailoring [crafted] | 27.0 spell_power points (3.33 DPS) | yes | Robe of the Magi (1716, -0.04 DPS) [world_drop]; Dreamweave Vest (10021, -0.25 DPS) [crafted]; Robe of Power (7054, -0.46 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.60 DPS) | yes | Arcane Runed Bracers (4744, -0.49 DPS) [quest]; Spidertank Oilrag (9448, -0.49 DPS) [dungeon]; Condor Bracers (15864, -0.74 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.60 DPS) | yes | Fiery Handwraps (254025, -0.21 DPS) [crafted]; Red Mageweave Gloves (10018, -0.29 DPS) [crafted]; Crimson Silk Gloves (7064, -0.30 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (2.49 DPS) | yes | Deathmage Sash (10771, -0.20 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.38 DPS) [rep]; Gilded Cord (254037, -0.74 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.3 spell_power points (2.87 DPS) | yes | Flame Leggings (253991, -0.48 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.94 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.96 DPS) | yes | Fiery Slippers (254005, -0.87 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.43 DPS) [crafted]; Acidic Walkers (9454, -1.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (1.71 DPS) | yes | Ring of Forlorn Spirits (2043, -0.72 DPS) [quest]; Reedknot Ring (9622, -0.85 DPS) [quest]; Minor Channeling Ring (1449, -0.90 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.11 DPS) | yes | Ring of Forlorn Spirits (2043, -0.12 DPS) [quest]; Reedknot Ring (9622, -0.25 DPS) [quest]; Minor Channeling Ring (1449, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (82.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Searing Golden Blade (12260, -0.71 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.3 spell_power points (1.89 DPS) | yes | Orb of the Forgotten Seer (7685, -0.16 DPS) [dungeon]; Thrash's Trash (276204, -0.16 DPS) [vendor]; Orb of Lorica (11262, -0.20 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 323.6 spell_power points (39.95 DPS) | yes | Ragefire Wand (7513, -1.42 DPS) [quest]; Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Cindercloth Robe; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 346, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 156.2. Weights run: 1.4s. Verify run: 1.7s. 443 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.939 ± 0.038, crit=0.343 ± 0.024 per rating point (14 rating = 1%, 4.796 per %), hit=0.622 ± 0.048 per rating point (10 rating = 1%, 6.225 per %), spell_haste=-3.557 ± 0.738, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.8 spell_power points (5.51 DPS) | yes | Dreamweave Circlet (10041, -1.08 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.13 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.57 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcane Crystal Pendant (20037, +0.00 DPS) [quest]; Horizon Choker (13085, -0.20 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.27 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.9 spell_power points (4.36 DPS) | yes | Kentic Amice (11624, -0.54 DPS) [dungeon]; Netherflame Shoulders (254053, -0.79 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.26 DPS) [vendor] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.5 spell_power points (2.99 DPS) | yes | Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.45 DPS) [dungeon]; Runecloth Cloak (13860, -0.58 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.8 spell_power points (5.51 DPS) | yes | Robe of the Magi (1716, -1.48 DPS) [world_drop]; Runecloth Tunic (13857, -1.52 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.55 DPS) [vendor] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.6 spell_power points (2.85 DPS) | yes | Nethergeld Cuffs (254061, -0.87 DPS) [crafted]; Bloodband Bracers (11469, -0.89 DPS) [quest]; Aristocratic Cuffs (12546, -2.29 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.4 spell_power points (4.57 DPS) | yes | Raider Handwraps (272098, -0.54 DPS) [vendor]; Dreamweave Gloves (10019, -1.40 DPS) [crafted]; Fiery Gloves (254099, -2.83 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | sim-verified (156.2 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Dawnspire Cord (12466, -0.53 DPS) [dungeon]; Satyrmane Sash (17755, -0.59 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.4 spell_power points (4.72 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Flame Leggings (253991, -1.72 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.56 DPS, sim-verified) [vendor] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.5 spell_power points (4.15 DPS) | yes | Earthen Silk Slippers (254013, -0.65 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -0.84 DPS) [vendor]; Cindercloth Boots (10044, -1.09 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.6 spell_power points (2.27 DPS) | yes | Brainlash (6440, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.52 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (2.14 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.25 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.39 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.53 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.63 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.78 DPS) [dungeon]; Blade of Eternal Darkness (17780, -13.06 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 360.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Sorcerer's Gauntlets; waist: Fiery Waistcord; legs: Spellshock Leggings; feet: Fiery Sandals; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 443, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 442.4. Weights run: 1.5s. Verify run: 5.2s. 1073 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.923 ± 0.054, crit=0.719 ± 0.043 per rating point (14 rating = 1%, 10.065 per %), hit=1.212 ± 0.083 per rating point (10 rating = 1%, 12.117 per %), spell_haste=-10.458 ± 1.052, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 54.1 spell_power points (14.58 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.74 DPS) [pvp]; Magister's Crown (16686, -14.91 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (442.4 DPS) | yes | Amulet of the Dawn (22657, -0.41 DPS) [quest]; Beads of Ogre Mojo (22149, -1.19 DPS) [quest]; Jewel of Kajaro (19601, -9.73 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 50.9 spell_power points (13.71 DPS) | yes | Mantle of the Timbermaw (19050, -3.19 DPS) [crafted]; Field Marshal's Silk Spaulders (231602, -3.25 DPS) [pvp]; Darkspear Shoulderpads (272103, -3.77 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.5 spell_power points (9.56 DPS) | yes | Hide of the Wild (18510, -3.31 DPS) [crafted]; Shroud of Arcane Mastery (22330, -3.56 DPS) [dungeon]; Crystalline Threaded Cape (20697, -7.98 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 61.1 spell_power points (16.47 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.64 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.88 DPS) [pvp]; Robes of Fiery Devastation (279266, -11.23 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.4 spell_power points (7.92 DPS) | yes | Sublime Wristguards (18497, -2.20 DPS) [dungeon]; Runecloth Cuffs (254123, -2.47 DPS) [crafted]; Netherflame Cuffs (254065, -2.67 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (442.4 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Inferno Gloves (18408, -11.53 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 57.4 spell_power points (15.47 DPS) | yes | Magician's Cord (272393, -4.38 DPS) [vendor]; Ban'thok Sash (11662, -6.24 DPS) [dungeon]; Belt of the Archmage (18405, -10.36 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 61.4 spell_power points (16.53 DPS) | yes | Marshal's Silk Leggings (231605, -0.76 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.94 DPS) [pvp]; Sorcerer's Leggings (226933, -6.47 DPS, sim-verified) [quest] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (442.4 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.81 DPS) [dungeon]; Sorcerer's Boots (22064, -6.85 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (442.4 DPS) | yes | Rune Band of Wizardry (22339, -5.47 DPS) [dungeon]; Channeler's Ring (272406, -6.30 DPS) [vendor]; Naglering (11669, -29.06 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (442.4 DPS) | yes | Rune Band of Wizardry (22339, -1.09 DPS) [dungeon]; Channeler's Ring (272406, -1.92 DPS) [vendor]; Naglering (11669, -12.06 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (442.4 DPS) | yes | Weakness Analyzer (272438, -1.89 DPS) [vendor]; Blackhand's Breadth (13965, -2.39 DPS) [quest]; Darkmoon Card: Blue Dragon (19288, -22.30 DPS, sim-verified) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (442.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Darkmoon Card: Blue Dragon (19288, -19.50 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (442.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -1.12 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -40.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 290.9 spell_power points (78.38 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.56 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.79 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.46 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; ranged: Torch of Light

No-known-source sample (15 of 1073, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 205015100000000000-23252100130133051-0050000000000000000)

Set DPS (verified): 824.2. Weights run: 1.6s. Verify run: 4.1s. 1073 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.409 ± 0.029, crit=0.870 ± 0.043 per rating point (14 rating = 1%, 12.183 per %), hit=1.128 ± 0.065 per rating point (10 rating = 1%, 11.277 per %), spell_haste=not significant (-0.196 ± 0.515), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 43.4 spell_power points (22.35 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.47 DPS) [pvp]; Crimson Felt Hat (18727, -7.90 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (824.2 DPS) | yes | Orb of the Darkmoon (19426, -0.80 DPS) [quest]; Chains of the Lich (23125, -0.80 DPS) [dungeon]; Jewel of Kajaro (19601, -13.36 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.3 spell_power points (23.34 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -7.02 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -7.31 DPS) [pvp]; Mantle of the Timbermaw (19050, -8.01 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.5 spell_power points (15.73 DPS) | yes | Mageflame Cloak (13007, -4.92 DPS) [world_drop]; Hide of the Wild (18510, -6.42 DPS) [crafted]; Crystalline Threaded Cape (20697, -7.56 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 57.1 spell_power points (29.40 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.55 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -8.73 DPS) [pvp]; Robes of Fiery Devastation (279266, -22.12 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.3 spell_power points (13.02 DPS) | yes | Sublime Wristguards (18497, -4.73 DPS) [dungeon]; Netherflame Cuffs (254065, -4.85 DPS) [crafted]; Runecloth Cuffs (254123, -5.25 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.5 spell_power points (19.29 DPS) | yes | Marshal's Silk Gloves (16440, -2.86 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.86 DPS) [vendor]; Inferno Gloves (18408, -15.44 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 45.3 spell_power points (23.32 DPS) | yes | Magician's Cord (272393, -8.48 DPS) [vendor]; Highlander's Cloth Girdle (20047, -8.57 DPS) [rep]; Belt of the Archmage (18405, -10.97 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 60.5 spell_power points (31.14 DPS) | yes | Marshal's Silk Leggings (231605, -5.20 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -10.47 DPS) [pvp]; Skyshroud Leggings (13170, -22.56 DPS, sim-verified) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 27.5 spell_power points (14.18 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -1.26 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (824.2 DPS) | yes | Elemental Focus Band (20682, -8.48 DPS) [world]; Don Julio's Band (19325, -10.04 DPS) [rep]; Naglering (11669, -34.33 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (824.2 DPS) | yes | Elemental Focus Band (20682, -0.40 DPS) [world]; Don Julio's Band (19325, -1.97 DPS) [rep]; Naglering (11669, -27.51 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (824.2 DPS) | yes | Eye of the Beast (13968, -2.39 DPS) [quest]; Weakness Analyzer (272438, -3.61 DPS) [vendor]; Serenity Field (272439, -7.73 DPS) [vendor] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (824.2 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Weakness Analyzer (272438, -1.22 DPS) [vendor]; Serenity Field (272439, -5.34 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (824.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.78 DPS) [quest]; Teebu's Blazing Longsword (1728, -59.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.0 spell_power points (79.85 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.33 DPS) [dungeon]; Bonecreeper Stylus (13938, -10.71 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.05 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Rune Band of Wizardry; trinket1: Briarwood Reed; trinket2: Blackhand's Breadth; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1073, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 26.3. Weights run: 1.2s. Verify run: 1.0s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.342 ± 0.011, crit=0.138 ± 0.004 per rating point (14 rating = 1%, 1.931 per %), hit=0.427 ± 0.013 per rating point (10 rating = 1%, 4.273 per %), spell_haste=0.549 ± 0.129, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.71 DPS) | yes | Pristine Circlet (253949, -0.32 DPS) [crafted]; Shadow Goggles (4373, -0.60 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.52 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.37 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Pearl-clasped Cloak (5542, -0.06 DPS) [crafted]; Feyscale Cloak (6632, -0.06 DPS) [dungeon]; Black Whelp Cloak (7283, -0.06 DPS) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 10.7 spell_power points (0.69 DPS) | yes | Filigreed Pristine Gown (253901, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.32 DPS) [crafted]; Green Woolen Robe (6243, -0.43 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.19 DPS) | yes | Mindthrust Bracers (1974, -0.08 DPS) [dungeon]; Featherbead Bracers (15452, -0.08 DPS) [quest]; Owlbeard Bracers (16981, -0.09 DPS) [quest] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.58 DPS) | yes | Flame Gloves (253917, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.4 spell_power points (0.54 DPS) | yes | Pristine Sash (253925, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.22 DPS) [crafted]; Novice Ardent's Sash (253887, -0.35 DPS) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.1 spell_power points (0.84 DPS) | yes | Abomination Skin Leggings (23173, -0.09 DPS) [dungeon]; Phoenix Pants (4317, -0.11 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.54 DPS) | yes | Flame Boots (253893, -0.14 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.22 DPS) [world]; Pristine Boots (253889, -0.28 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.32 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Loop of Sacrifice (281673, -0.21 DPS) [quest]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.19 DPS) | yes | Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Volcanic Rock Ring (12053, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.4 spell_power points (0.22 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 347.7 spell_power points (22.50 DPS) | yes | Skycaller (12984, -0.65 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.21 DPS) [dungeon]; Sizzle Stick (8071, -4.53 DPS) [quest] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Tabitha's Cuffs; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 50.1. Weights run: 1.3s. Verify run: 1.1s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.705 ± 0.013, crit=0.174 ± 0.008 per rating point (14 rating = 1%, 2.429 per %), hit=0.356 ± 0.017 per rating point (10 rating = 1%, 3.560 per %), spell_haste=0.908 ± 0.190, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.19 DPS) [world_drop]; Flame Circlet (253953, -0.19 DPS) [crafted]; Filigreed Flame Circlet (253979, -0.65 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 spell_power points (1.06 DPS) | yes | Darkspear Warding Pendant (272075, -0.69 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (1.45 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Mantle of Woe (7750, -0.18 DPS) [quest]; Fairywing Mantle (9536, -0.28 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Hillman's Cloak (3719, -0.06 DPS) [crafted]; Windsong Drape (15468, -0.06 DPS) [quest]; Cloak of Rot (4462, -0.68 DPS, sim-verified) [world] |
| chest | Flame Gown (253965) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.14 DPS) [dungeon]; Death Speaker Robes (6682, -0.30 DPS) [dungeon]; Green Silk Armor (7065, -0.50 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.23 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Glowing Magical Bracelets (13106, -0.69 DPS) [world_drop]; Nightsky Wristbands (6407, -0.83 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.5 spell_power points (0.90 DPS) | yes | Flame Gloves (253917, -0.04 DPS) [crafted]; Phoenix Gloves (4331, -0.05 DPS) [crafted]; Truefaith Gloves (7049, -0.23 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warsong Sash (16975, -0.01 DPS) [quest]; Crimson Silk Belt (7055, -0.02 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.70 DPS, sim-verified) [rep] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.9 spell_power points (1.78 DPS) | yes | Abomination Skin Leggings (23173, -0.40 DPS) [dungeon]; Filigreed Flame Leggings (253941, -0.41 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.56 DPS) [world] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.9 spell_power points (1.69 DPS) | yes | Gilded Slippers (254001, -0.57 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.19 DPS) [world]; Snake Hoop (6750, -0.19 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.57 DPS) | yes | Snake Hoop (6750, -0.10 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; Black Widow Band (6199, -0.49 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 26.5 spell_power points (2.50 DPS) | yes | Glimmering Staff (249392, -1.77 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -1.83 DPS) [quest]; Scorn's Focal Dagger (23168, -5.78 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Unstable Power Core (279847) | Source of Power [quest] | 362.1 spell_power points (34.12 DPS) | yes | Necrotic Wand (7708, -0.63 DPS) [dungeon]; Starfaller (13063, -0.99 DPS) [world_drop]; Greater Mystic Wand (217287, -4.65 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Flame Gown; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Unstable Power Core

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552100130103041-0000000000000000000)

Set DPS (verified): 79.4. Weights run: 1.5s. Verify run: 1.3s. 323 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.774 ± 0.026, crit=0.219 ± 0.015 per rating point (14 rating = 1%, 3.059 per %), hit=0.376 ± 0.034 per rating point (10 rating = 1%, 3.765 per %), spell_haste=2.538 ± 0.457, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.59 DPS) | yes | Augural Shroud (2620, -0.28 DPS) [world]; Corpseshroud (10574, -0.78 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.84 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.2 spell_power points (1.63 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.67 DPS) [quest]; Scorn's Icy Choker (23169, -0.86 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.1 spell_power points (2.11 DPS) | yes | Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -0.81 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.0 spell_power points (1.97 DPS) | yes | Guardian Cloak (5965, -0.75 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.92 DPS) [vendor]; Long Silken Cloak (4326, -1.93 DPS, sim-verified) [crafted] |
| chest | Cindercloth Robe (10042) | Tailoring [crafted] | 27.0 spell_power points (3.33 DPS) | yes | Robe of the Magi (1716, -0.04 DPS) [world_drop]; Dreamweave Vest (10021, -0.25 DPS) [crafted]; Robe of Power (7054, -0.46 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.60 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.35 DPS) [quest]; Spidertank Oilrag (9448, -0.49 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.60 DPS) | yes | Red Mageweave Gloves (10018, -0.29 DPS) [crafted]; Crimson Silk Gloves (7064, -0.30 DPS) [crafted]; Fiery Handwraps (254025, -0.90 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.2 spell_power points (2.49 DPS) | yes | Deathmage Sash (10771, -0.20 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.38 DPS) [rep]; Gilded Cord (254037, -0.74 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.3 spell_power points (2.87 DPS) | yes | Flame Leggings (253991, -0.48 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.94 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.96 DPS) | yes | Fiery Slippers (254005, -0.69 DPS) [crafted]; Gilded Slippers (254001, -1.43 DPS) [crafted]; Acidic Walkers (9454, -1.58 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (1.71 DPS) | yes | Reedknot Ring (9622, -0.85 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.97 DPS) [vendor]; Black Widow Band (6199, -1.04 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.11 DPS) | yes | Sea Giant's Toe Ring (274746, -0.37 DPS) [vendor]; Black Widow Band (6199, -0.44 DPS) [world]; Reedknot Ring (9622, -1.34 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (79.4 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Searing Golden Blade (12260, -0.71 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.3 spell_power points (1.89 DPS) | yes | Orb of the Forgotten Seer (7685, -0.16 DPS) [dungeon]; Thrash's Trash (276204, -0.16 DPS) [vendor]; Rod of Molten Fire (2565, -0.29 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 323.6 spell_power points (39.95 DPS) | yes | Ragefire Wand (7513, -1.42 DPS) [quest]; Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Cindercloth Robe; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 151.1. Weights run: 1.4s. Verify run: 1.6s. 416 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.939 ± 0.038, crit=0.343 ± 0.024 per rating point (14 rating = 1%, 4.796 per %), hit=0.622 ± 0.048 per rating point (10 rating = 1%, 6.225 per %), spell_haste=-3.557 ± 0.738, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.8 spell_power points (5.51 DPS) | yes | Blood Guard's Dreadweave Hat (220907, -1.13 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.57 DPS) [crafted]; Dreamweave Circlet (10041, -2.26 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.6 spell_power points (3.15 DPS) | yes | Horizon Choker (13085, -1.24 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.31 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.53 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.9 spell_power points (4.36 DPS) | yes | Kentic Amice (11624, -0.54 DPS) [dungeon]; Netherflame Shoulders (254053, -0.79 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.26 DPS) [vendor] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.5 spell_power points (2.99 DPS) | yes | Deep Woodlands Cloak (19121, -0.01 DPS) [quest]; Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.45 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.8 spell_power points (5.51 DPS) | yes | Runecloth Tunic (13857, -1.52 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.55 DPS) [vendor]; Robe of the Magi (1716, -4.41 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.6 spell_power points (2.85 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Aristocratic Cuffs (12546, -0.80 DPS) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.4 spell_power points (4.57 DPS) | yes | Raider Handwraps (272098, -0.54 DPS) [vendor]; Dreamweave Gloves (10019, -1.40 DPS) [crafted]; Fiery Gloves (254099, -3.08 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Dawnspire Cord (12466, -0.53 DPS) [dungeon]; Satyrmane Sash (17755, -0.59 DPS) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | sim-verified (151.1 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Red Mageweave Pants (10009, -0.41 DPS) [crafted]; Flame Leggings (253991, -1.09 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 28.5 spell_power points (4.15 DPS) | yes | Earthen Silk Slippers (254013, -0.65 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -0.84 DPS) [vendor]; Cindercloth Boots (10044, -1.09 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.6 spell_power points (2.27 DPS) | yes | Brainlash (6440, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Advisor's Ring (19519, -0.52 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (2.14 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.25 DPS) [world_drop]; Advisor's Ring (19519, -0.39 DPS) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.24 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.18 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.63 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.78 DPS) [dungeon]; Blade of Eternal Darkness (17780, -13.71 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 360.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Sorcerer's Gauntlets; waist: Fiery Waistcord; legs: Stone Guard's Dreadweave Leggings; feet: Fiery Sandals; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 419.7. Weights run: 1.5s. Verify run: 5.1s. 1061 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.923 ± 0.054, crit=0.719 ± 0.043 per rating point (14 rating = 1%, 10.065 per %), hit=1.212 ± 0.083 per rating point (10 rating = 1%, 12.117 per %), spell_haste=-10.458 ± 1.052, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 54.1 spell_power points (14.58 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.74 DPS) [pvp]; Magister's Crown (16686, -19.68 DPS, sim-verified) [dungeon] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Amulet of the Dawn (22657, +0.00 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 50.9 spell_power points (13.71 DPS) | yes | Mantle of the Timbermaw (19050, -3.19 DPS) [crafted]; Warlord's Silk Amice (231594, -3.25 DPS) [pvp]; Darkspear Shoulderpads (272103, -3.77 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.5 spell_power points (9.56 DPS) | yes | Crystalline Threaded Cape (20697, -3.18 DPS) [world]; Hide of the Wild (18510, -3.31 DPS) [crafted]; Shroud of Arcane Mastery (22330, -3.56 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 61.1 spell_power points (16.47 DPS) | yes | Warlord's Silk Raiment (231596, -0.64 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.88 DPS) [pvp]; Robes of Fiery Devastation (279266, -6.30 DPS, sim-verified) [crafted] |
| wrist | Sorcerer's Bindings (226929) | An Earnest Proposition [quest] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; Inferno Gloves (18408, -14.05 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 57.4 spell_power points (15.47 DPS) | yes | Magician's Cord (272393, -4.38 DPS) [vendor]; Ban'thok Sash (11662, -6.24 DPS) [dungeon]; Belt of the Archmage (18405, -8.94 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 61.4 spell_power points (16.53 DPS) | yes | General's Silk Trousers (231595, -0.76 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -3.94 DPS) [pvp]; Sorcerer's Leggings (226933, -7.41 DPS, sim-verified) [quest] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.81 DPS) [dungeon]; Sorcerer's Boots (22064, -15.45 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -5.47 DPS) [dungeon]; Channeler's Ring (272406, -6.30 DPS) [vendor]; Naglering (11669, -17.33 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.09 DPS) [dungeon]; Channeler's Ring (272406, -1.92 DPS) [vendor]; Naglering (11669, -7.00 DPS, sim-verified) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -0.50 DPS) [quest]; Eye of the Beast (13968, -0.50 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (419.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -1.12 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -35.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.23 DPS) [dungeon]; Sparkling Crystal Wand (20672, -0.90 DPS) [world]; Torch of Light (279246, -6.23 DPS, sim-verified) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Jewel of Kajaro; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Sorcerer's Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Weakness Analyzer; trinket2: Serenity Field; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1061, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (orc, 205015100000000000-23252100130133051-0050000000000000000)

Set DPS (verified): 805.1. Weights run: 1.6s. Verify run: 3.9s. 1061 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.409 ± 0.029, crit=0.870 ± 0.043 per rating point (14 rating = 1%, 12.183 per %), hit=1.128 ± 0.065 per rating point (10 rating = 1%, 11.277 per %), spell_haste=not significant (-0.196 ± 0.515), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 43.4 spell_power points (22.35 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.47 DPS) [pvp]; Crimson Felt Hat (18727, -5.22 DPS) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (805.1 DPS) | yes | Orb of the Darkmoon (19426, -0.80 DPS) [quest]; Chains of the Lich (23125, -0.80 DPS) [dungeon]; Jewel of Kajaro (19601, -12.86 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.3 spell_power points (23.34 DPS) | yes | Mantle of the Timbermaw (19050, -6.70 DPS, sim-verified) [crafted]; Champion's Silk Mantle (227104, -7.02 DPS) [pvp]; Warlord's Silk Amice (231594, -7.31 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.5 spell_power points (15.73 DPS) | yes | Crystalline Threaded Cape (20697, -4.59 DPS) [world]; Mageflame Cloak (13007, -4.92 DPS) [world_drop]; Hide of the Wild (18510, -6.42 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 57.1 spell_power points (29.40 DPS) | yes | Warlord's Silk Raiment (231596, -2.55 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -8.73 DPS) [pvp]; Robes of Fiery Devastation (279266, -17.59 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.3 spell_power points (13.02 DPS) | yes | Sublime Wristguards (18497, -4.73 DPS) [dungeon]; Netherflame Cuffs (254065, -4.85 DPS) [crafted]; Runecloth Cuffs (254123, -5.25 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.5 spell_power points (19.29 DPS) | yes | General's Silk Handguards (16540, -2.86 DPS) [vendor]; General's Silk Gauntlets (231599, -2.86 DPS) [vendor]; Inferno Gloves (18408, -6.91 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 45.3 spell_power points (23.32 DPS) | yes | Belt of the Archmage (18405, -3.37 DPS) [crafted]; Magician's Cord (272393, -8.48 DPS) [vendor]; Defiler's Cloth Girdle (20163, -8.57 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 60.5 spell_power points (31.14 DPS) | yes | General's Silk Trousers (231595, -5.20 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -10.47 DPS) [pvp]; Skyshroud Leggings (13170, -14.73 DPS, sim-verified) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 27.5 spell_power points (14.18 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.26 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (805.1 DPS) | yes | Elemental Focus Band (20682, -8.48 DPS) [world]; Don Julio's Band (19325, -10.04 DPS) [rep]; Naglering (11669, -33.15 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (805.1 DPS) | yes | Elemental Focus Band (20682, -0.40 DPS) [world]; Don Julio's Band (19325, -1.97 DPS) [rep]; Naglering (11669, -25.21 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (805.1 DPS) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Eye of the Beast (13968, -2.39 DPS) [quest]; Weakness Analyzer (272438, -3.61 DPS) [vendor] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (805.1 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -1.22 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (805.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.78 DPS) [quest]; Teebu's Blazing Longsword (1728, -56.38 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.0 spell_power points (79.85 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.33 DPS) [dungeon]; Bonecreeper Stylus (13938, -10.71 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.05 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Rune Band of Wizardry; trinket1: Briarwood Reed; trinket2: Blackhand's Breadth; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1061, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

