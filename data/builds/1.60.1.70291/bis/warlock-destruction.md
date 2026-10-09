# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.7. Weights run: 2.3s. Verify run: 1.3s. 151 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.063 ± 0.003, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.413 per %), hit=0.121 ± 0.004 per rating point (10 rating = 1%, 1.214 per %), spell_haste=0.614 ± 0.045, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.002, fire_power=0.235 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.20 DPS) | yes | Shadow Goggles (4373, -1.14 DPS) [crafted]; Flame Circlet (253953, -2.39 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.6 spell_power points (1.11 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.19 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.31 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Feyscale Cloak (6632, -0.20 DPS) [dungeon]; Caretaker's Cape (20428, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.24 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.3 spell_power points (1.06 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -1.39 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.14 DPS) [dungeon]; Bright Bracers (3647, -0.15 DPS) [world_drop]; Repurposed Hair Band (281256, -0.17 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.56 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.97 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.3 spell_power points (0.85 DPS) | yes | Novice Ardent's Sash (253887, -0.41 DPS) [crafted]; Flame Sash (253929, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.96 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 spell_power points (1.90 DPS) | yes | Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted]; Silk-threaded Trousers (1929, -0.66 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.3 spell_power points (1.45 DPS) | yes | Red Woolen Boots (4313, -0.65 DPS) [crafted]; Feather Padded Treads (285345, -0.66 DPS, sim-verified) [world]; Pristine Boots (253889, -0.81 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.1 spell_power points (1.02 DPS) | yes | Sludge-Stained Band (286535, -0.42 DPS) [world]; Lavishly Jeweled Ring (1156, -0.95 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.99 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.00 DPS) | yes | Sludge-Stained Band (286535, -0.57 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.92 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.96 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.6 spell_power points (0.13 DPS) | yes | Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.05 DPS) [world]; Staff of Westfall (2042, -0.06 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.7 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.58 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.62 DPS) [dungeon]; Sizzle Stick (8071, -4.26 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 151, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053225101000000)

Set DPS (verified): 86.0. Weights run: 2.5s. Verify run: 1.4s. 266 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.337 ± 0.008, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.980 per %), hit=0.166 ± 0.007 per rating point (10 rating = 1%, 1.665 per %), spell_haste=not significant (0.109 ± 0.120), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.767 ± 0.002, fire_power=0.233 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.56 DPS) | yes | Silk Headband (7050, -0.46 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.70 DPS) [crafted]; Enchanter's Cowl (4322, -3.87 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.10 DPS) | yes | Crystal Starfire Medallion (5003, -1.78 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.78 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.96 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (2.79 DPS) | yes | Death Speaker Mantle (6685, -0.65 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.70 DPS) [quest]; Invoker's Mantle (215365, -0.78 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (1.39 DPS) | yes | Hillman's Cloak (3719, -0.23 DPS) [crafted]; Repairman's Cape (9605, -0.38 DPS) [quest]; Heavy Woolen Cloak (4311, -0.46 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.22 DPS) [dungeon]; Beguiler Robes (7728, -0.30 DPS) [dungeon]; Green Silk Armor (7065, -0.85 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.09 DPS) | yes | Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop]; Nightsky Wristbands (6407, -1.62 DPS) [world_drop]; Phoenix Bindings (210781, -1.65 DPS, sim-verified) [crafted] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.7 spell_power points (1.79 DPS) | yes | Shilly Mitts (9609, +0.00 DPS, sim-verified) [quest]; Serpent Gloves (5970, -0.16 DPS) [dungeon]; Truefaith Gloves (7049, -0.39 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.0 spell_power points (2.79 DPS) | yes | Belt of Arugal (6392, -0.46 DPS) [dungeon]; Invoker's Cord (215366, -0.77 DPS) [crafted]; Crimson Silk Belt (7055, -0.85 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.79 DPS) | yes | Pristine Leggings (253987, -0.61 DPS) [crafted]; Abomination Skin Leggings (23173, -0.77 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.92 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (2.17 DPS) | yes | Acidic Walkers (9454, -0.39 DPS) [dungeon]; Nimbus Boots (6998, -0.78 DPS) [quest]; Spidersilk Boots (4320, -2.04 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.63 DPS) | yes | Minor Channeling Ring (1449, -0.31 DPS) [quest]; Electrocutioner Lagnut (9447, -0.93 DPS) [dungeon]; Sludge-Stained Band (286535, -0.93 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.39 DPS) | yes | Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon]; Sludge-Stained Band (286535, -0.70 DPS) [world]; Minor Channeling Ring (1449, -1.71 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Hardened Root Staff (1317, -1.77 DPS, sim-verified) [quest]; Scorn's Focal Dagger (23168, -3.64 DPS) [dungeon]; Staff of Soran'ruk (15109, -3.89 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 146.0 spell_power points (33.91 DPS) | yes | Starfaller (13063, -0.93 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.75 DPS) [crafted]; Scorching Wand (5213, -4.73 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 266, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053225103101321)

Set DPS (verified): 134.1. Weights run: 2.1s. Verify run: 1.4s. 346 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.662 ± 0.021, crit=0.116 ± 0.004 per rating point (14 rating = 1%, 1.617 per %), hit=0.289 ± 0.020 per rating point (10 rating = 1%, 2.886 per %), spell_haste=not significant (0.387 ± 0.235), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.292 ± 0.000, fire_power=0.708 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.22 DPS) | yes | Augural Shroud (2620, -1.91 DPS, sim-verified) [world]; Electromagnetic Gigaflux Reactivator (9492, -1.92 DPS) [dungeon]; Living Cowl (5608, -1.99 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.3 spell_power points (3.06 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.41 DPS) [quest]; Scorn's Icy Choker (23169, -1.81 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.90 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 spell_power points (3.88 DPS) | yes | Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest]; Green Silken Shoulders (7057, -1.73 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.0 spell_power points (3.72 DPS) | yes | Guardian Cloak (5965, -1.40 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.91 DPS) [vendor]; Long Silken Cloak (4326, -2.35 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 spell_power points (6.46 DPS) | yes | Dreamweave Vest (10021, -0.50 DPS) [crafted]; Robe of Power (7054, -1.00 DPS) [crafted]; Elemental Raiment (9434, -1.24 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.2 spell_power points (2.29 DPS) | yes | Arcane Runed Bracers (4744, -0.05 DPS) [quest]; Spidertank Oilrag (9448, -0.05 DPS) [dungeon]; Condor Bracers (15864, -0.55 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (5.13 DPS) | yes | Black Mageweave Gloves (10003, -1.40 DPS) [crafted]; Fiery Handwraps (254025, -1.52 DPS) [crafted]; Red Mageweave Gloves (10018, -1.86 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.9 spell_power points (4.21 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Fiery Cord (254041, -0.43 DPS) [crafted]; Gilded Cord (254037, -0.90 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.9 spell_power points (5.45 DPS) | yes | Crimson Silk Pantaloons (7062, -1.44 DPS, sim-verified) [crafted]; Flame Leggings (253991, -1.84 DPS) [crafted]; Abomination Skin Leggings (23173, -1.90 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.97 DPS) | yes | Gilded Slippers (254001, -3.07 DPS) [crafted]; Acidic Walkers (9454, -3.41 DPS) [dungeon]; Fiery Slippers (254005, -4.00 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (3.31 DPS) | yes | Ring of Forlorn Spirits (2043, -1.32 DPS) [quest]; Reedknot Ring (9622, -1.57 DPS) [quest]; Minor Channeling Ring (1449, -1.74 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.24 DPS) | yes | Reedknot Ring (9622, -0.50 DPS) [quest]; Minor Channeling Ring (1449, -0.67 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.96 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.33 DPS) [dungeon]; Searing Golden Blade (12260, -4.24 DPS) [crafted] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.48 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Noh'Orahil (15107, -0.19 DPS) [quest]; Orb of Lorica (11262, -0.34 DPS) [quest] |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (134.1 DPS) | yes | Earthen Rod (9381, -0.08 DPS) [dungeon]; Twisted Nether Wand (249144, -0.18 DPS) [crafted]; Jaina's Firestarter (13064, -1.87 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Umbral Wand

No-known-source sample (15 of 346, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 233.7. Weights run: 2.0s. Verify run: 1.6s. 436 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.365 ± 0.023, crit=0.151 ± 0.005 per rating point (14 rating = 1%, 2.109 per %), hit=0.406 ± 0.024 per rating point (10 rating = 1%, 4.059 per %), spell_haste=not significant (0.459 ± 0.310), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.293 ± 0.000, fire_power=0.707 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.35 DPS) | yes | Dreamweave Circlet (10041, -0.64 DPS) [crafted]; Red Mageweave Headband (10033, -1.55 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.63 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.9 spell_power points (2.70 DPS) | yes | Mindburst Medallion (11196, -0.47 DPS) [quest]; Scorn's Icy Choker (23169, -0.87 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.31 DPS) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Netherflame Shoulders (254053, -1.04 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.46 DPS) [vendor]; Rotgrip Mantle (17732, -2.34 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.2 spell_power points (4.41 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.96 DPS, sim-verified) [dungeon]; Cindercloth Cloak (14044, -1.11 DPS) [crafted]; Runecloth Cloak (13860, -1.16 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 26.3 spell_power points (7.16 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.37 DPS) [crafted]; Runecloth Tunic (13857, -1.44 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 11.7 spell_power points (3.20 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Phoenix Bindings (210781, -0.70 DPS) [crafted]; Arcane Runed Bracers (4744, -0.75 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.5 spell_power points (5.30 DPS) | yes | Fiery Gloves (254099, -0.65 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.86 DPS) [vendor]; Runecloth Gloves (13863, -1.14 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 20.1 spell_power points (5.47 DPS) | yes | Fiery Waistcord (254085, -0.91 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.26 DPS) [rep]; Satyrmane Sash (17755, -2.01 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.7 spell_power points (7.26 DPS) | yes | Knight's Dreadweave Leggings (220888, -2.22 DPS) [vendor]; Red Mageweave Pants (10009, -2.25 DPS) [crafted]; Wizardweave Leggings (14132, -2.64 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.53 DPS) | yes | Fiery Sandals (254111, -2.25 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.36 DPS) [vendor]; Cindercloth Boots (10044, -2.49 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.54 DPS) | yes | Philanthropist's Ring (281635, -0.32 DPS) [quest]; Cyclopean Band (11824, -0.39 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.36 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.27 DPS) | yes | Philanthropist's Ring (281635, -0.05 DPS) [quest]; Cyclopean Band (11824, -0.12 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.09 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.63 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.54 DPS) [quest]; Spire of Hakkar (10844, -0.65 DPS) [world]; Blade of Eternal Darkness (17780, -21.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.94 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.43 DPS) [crafted]; Pyric Caduceus (11748, -7.47 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 448.8. Weights run: 2.2s. Verify run: 3.5s. 1064 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.493 ± 0.033, crit=0.219 ± 0.006 per rating point (14 rating = 1%, 3.068 per %), hit=0.627 ± 0.034 per rating point (10 rating = 1%, 6.265 per %), spell_haste=not significant (-1.774 ± 0.541), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.386 ± 0.001, fire_power=0.614 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.37 DPS) [pvp]; Deathmist Mask (226909, -4.86 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.12 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.22 DPS) [quest]; Diana's Pearl Necklace (22403, -1.03 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.5 spell_power points (13.83 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.51 DPS) [pvp]; Burial Shawl (18681, -1.81 DPS, sim-verified) [dungeon]; Mantle of the Timbermaw (19050, -4.06 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.2 spell_power points (9.67 DPS) | yes | Crystalline Threaded Cape (20697, -1.56 DPS) [world]; Hide of the Wild (18510, -2.69 DPS) [crafted]; Amplifying Cloak (18350, -3.03 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.4 spell_power points (18.62 DPS) | yes | Robe of Everlasting Night (18385, -2.25 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -2.44 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -5.75 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.9 spell_power points (9.58 DPS) | yes | Sublime Wristguards (18497, -3.33 DPS) [dungeon]; Runecloth Cuffs (254123, -3.70 DPS) [crafted]; Deathmist Bracers (226907, -4.81 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -1.21 DPS) [quest]; Sandworm Skin Gloves (20716, -4.20 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 42.1 spell_power points (15.55 DPS) | yes | Belt of the Archmage (18405, -4.05 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.80 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -7.08 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 43.1 spell_power points (15.90 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -3.19 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.9 spell_power points (9.57 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Dragonrider Boots (18102, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.07 DPS) [dungeon]; Maiden's Circle (13001, -6.37 DPS) [world_drop]; Naglering (11669, -14.56 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.91 DPS) [dungeon]; Maiden's Circle (13001, -2.20 DPS) [world_drop]; Naglering (11669, -7.74 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+15.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.21 DPS) [quest]; Weakness Analyzer (272438, -2.58 DPS) [vendor]; Serenity Field (272439, -5.54 DPS) [vendor] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -1.78 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.21 DPS) [quest]; Teebu's Blazing Longsword (1728, -30.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+25.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.49 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.10 DPS) [world]; Torch of Light (279246, -25.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1064, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 845.4. Weights run: 7.1s. Verify run: 4.0s. 1064 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.093 ± 0.025, crit=0.275 ± 0.009 per rating point (14 rating = 1%, 3.847 per %), hit=0.832 ± 0.052 per rating point (10 rating = 1%, 8.320 per %), spell_haste=not significant (-0.579 ± 0.655), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.305 ± 0.001, fire_power=0.695 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.7 spell_power points (16.04 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -1.14 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -1.95 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (11.48 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.05 DPS) [dungeon]; Amulet of the Dawn (22657, -3.02 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 32.2 spell_power points (16.82 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.95 DPS) [pvp]; Argent Shoulders (19059, -3.71 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -5.31 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.1 spell_power points (13.08 DPS) | yes | Crystalline Threaded Cape (20697, -2.45 DPS) [world]; Amplifying Cloak (18350, -3.69 DPS) [dungeon]; Hide of the Wild (18510, -5.29 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.8 spell_power points (24.43 DPS) | yes | Robe of Everlasting Night (18385, -5.99 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -6.57 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.42 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.7 spell_power points (11.87 DPS) | yes | Sublime Wristguards (18497, -5.12 DPS) [dungeon]; Runecloth Cuffs (254123, -5.64 DPS) [crafted]; Netherflame Cuffs (254065, -6.81 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.5 spell_power points (14.33 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.92 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 35.4 spell_power points (18.45 DPS) | yes | Belt of the Archmage (18405, -5.32 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -7.32 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -8.58 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.6 spell_power points (21.19 DPS) | yes | Marshal's Dreadweave Leggings (231587, -0.97 DPS) [pvp]; Skyshroud Leggings (13170, -3.07 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -5.95 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.52 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.04 DPS) [crafted]; Omnicast Boots (11822, -1.50 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Elemental Focus Band (20682, -7.37 DPS) [world]; Maiden's Circle (13001, -9.65 DPS) [world_drop]; Naglering (11669, -22.28 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Elemental Focus Band (20682, -0.68 DPS) [world]; Maiden's Circle (13001, -2.96 DPS) [world_drop]; Naglering (11669, -16.50 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+21.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -3.13 DPS) [quest]; Weakness Analyzer (272438, -3.65 DPS) [vendor]; Serenity Field (272439, -7.82 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.46 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -45.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (845.4 DPS) | yes | Bonecreeper Stylus (13938, -1.03 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.21 DPS) [world]; Torch of Light (279246, -36.75 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Rune Band of Wizardry; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Spire of Hakkar; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1064, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.4. Weights run: 2.3s. Verify run: 1.2s. 140 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.063 ± 0.003, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.413 per %), hit=0.121 ± 0.004 per rating point (10 rating = 1%, 1.214 per %), spell_haste=0.614 ± 0.045, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.002, fire_power=0.235 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.20 DPS) | yes | Shadow Goggles (4373, -1.14 DPS) [crafted]; Flame Circlet (253953, -2.38 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.6 spell_power points (1.11 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.31 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.31 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Feyscale Cloak (6632, -0.20 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.53 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.3 spell_power points (1.06 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -1.51 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.60 DPS) | yes | Windsong Bangles (263336, -0.40 DPS) [quest]; Featherbead Bracers (15452, -0.54 DPS) [quest]; Owlbeard Bracers (16981, -0.55 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Gnoll Casting Gloves (892, -0.34 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.56 DPS) [crafted]; Apothecary Gloves (10919, -0.60 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.3 spell_power points (0.85 DPS) | yes | Novice Ardent's Sash (253887, -0.41 DPS) [crafted]; Flame Sash (253929, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.11 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 spell_power points (1.90 DPS) | yes | Silk-threaded Trousers (1929, -0.61 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.3 spell_power points (1.45 DPS) | yes | Red Woolen Boots (4313, -0.65 DPS) [crafted]; Feather Padded Treads (285345, -0.68 DPS, sim-verified) [world]; Pristine Boots (253889, -0.81 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.92 DPS) [dungeon]; Loop of Sacrifice (281673, -0.94 DPS) [quest]; Volcanic Rock Ring (12053, -0.96 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.60 DPS) | yes | Loop of Sacrifice (281673, -0.54 DPS) [quest]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.84 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.6 spell_power points (0.13 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.05 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.7 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.72 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.62 DPS) [dungeon]; Sizzle Stick (8071, -4.26 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-0000000000000000000-2053225101000000)

Set DPS (verified): 84.4. Weights run: 2.5s. Verify run: 1.3s. 249 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.337 ± 0.008, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.980 per %), hit=0.166 ± 0.007 per rating point (10 rating = 1%, 1.665 per %), spell_haste=not significant (0.109 ± 0.120), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.767 ± 0.002, fire_power=0.233 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.56 DPS) | yes | Silk Headband (7050, -0.46 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.70 DPS) [crafted]; Enchanter's Cowl (4322, -3.62 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.10 DPS) | yes | Crystal Starfire Medallion (5003, -1.78 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.78 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.09 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (2.79 DPS) | yes | Death Speaker Mantle (6685, -0.58 DPS, sim-verified) [dungeon]; Mantle of Woe (7750, -0.69 DPS) [quest]; Fairywing Mantle (9536, -0.70 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.16 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Battle Healer's Cloak (19529, -0.23 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.4 spell_power points (3.11 DPS) | yes | Tree Bark Jacket (1486, -0.09 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.31 DPS) [dungeon]; Beguiler Robes (7728, -0.39 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.09 DPS) | yes | Tabitha's Cuffs (251486, -1.39 DPS) [quest]; Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop]; Phoenix Bindings (210781, -2.16 DPS, sim-verified) [crafted] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.7 spell_power points (1.78 DPS) | yes | Serpent Gloves (5970, -0.16 DPS) [dungeon]; Truefaith Gloves (7049, -0.39 DPS) [crafted]; Gnoll Casting Gloves (892, -0.39 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (2.79 DPS) | yes | Belt of Arugal (6392, -0.46 DPS) [dungeon]; Warsong Sash (16975, -0.62 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.77 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.79 DPS) | yes | Pristine Leggings (253987, -0.61 DPS) [crafted]; Abomination Skin Leggings (23173, -0.86 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.92 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (2.17 DPS) | yes | Acidic Walkers (9454, -0.39 DPS) [dungeon]; Fiery Slippers (254005, -0.92 DPS) [crafted]; Spidersilk Boots (4320, -2.37 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.63 DPS) | yes | Electrocutioner Lagnut (9447, -0.93 DPS) [dungeon]; Sludge-Stained Band (286535, -0.93 DPS) [world]; Black Widow Band (6199, -1.08 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.39 DPS) | yes | Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon]; Black Widow Band (6199, -0.85 DPS) [world]; Sludge-Stained Band (286535, -2.42 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 24.7 spell_power points (5.73 DPS) | yes | Staff of Soran'ruk (15109, -3.89 DPS) [quest]; Glimmering Staff (249392, -4.87 DPS) [crafted]; Scorn's Focal Dagger (23168, -10.06 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (84.4 DPS) | yes | Starfaller (13063, -0.73 DPS) [world_drop]; Unstable Power Core (279847, -3.10 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.75 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2053225103101321)

Set DPS (verified): 130.6. Weights run: 2.1s. Verify run: 1.3s. 324 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.662 ± 0.021, crit=0.116 ± 0.004 per rating point (14 rating = 1%, 1.617 per %), hit=0.289 ± 0.020 per rating point (10 rating = 1%, 2.886 per %), spell_haste=not significant (0.387 ± 0.235), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.292 ± 0.000, fire_power=0.708 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.22 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -1.92 DPS) [dungeon]; Living Cowl (5608, -1.99 DPS) [world]; Augural Shroud (2620, -2.60 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.3 spell_power points (3.06 DPS) | yes | Scorn's Icy Choker (23169, -0.95 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.41 DPS) [quest]; Darkspear Warding Pendant (272074, -1.90 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 spell_power points (3.88 DPS) | yes | Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest]; Green Silken Shoulders (7057, -0.86 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.0 spell_power points (3.72 DPS) | yes | Guardian Cloak (5965, -1.40 DPS) [crafted]; Long Silken Cloak (4326, -1.81 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -1.91 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 spell_power points (6.46 DPS) | yes | Dreamweave Vest (10021, -0.98 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.00 DPS) [crafted]; Zealot's Robe (17043, -1.08 DPS) [quest] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.3 spell_power points (2.31 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Phoenix Bindings (210781, -0.02 DPS) [crafted]; Spidertank Oilrag (9448, -0.07 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (5.13 DPS) | yes | Black Mageweave Gloves (10003, -1.40 DPS) [crafted]; Fiery Handwraps (254025, -1.52 DPS) [crafted]; Red Mageweave Gloves (10018, -2.71 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.9 spell_power points (4.21 DPS) | yes | Defiler's Cloth Girdle (20166, -0.07 DPS) [rep]; Fiery Cord (254041, -0.43 DPS) [crafted]; Gilded Cord (254037, -0.90 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.9 spell_power points (5.45 DPS) | yes | Crimson Silk Pantaloons (7062, -1.58 DPS, sim-verified) [crafted]; Flame Leggings (253991, -1.84 DPS) [crafted]; Abomination Skin Leggings (23173, -1.90 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.97 DPS) | yes | Gilded Slippers (254001, -3.07 DPS) [crafted]; Acidic Walkers (9454, -3.41 DPS) [dungeon]; Fiery Slippers (254005, -4.00 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (3.31 DPS) | yes | Reedknot Ring (9622, -1.57 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.82 DPS) [vendor]; Black Widow Band (6199, -2.16 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.24 DPS) | yes | Sea Giant's Toe Ring (274746, -0.75 DPS) [vendor]; Reedknot Ring (9622, -1.02 DPS, sim-verified) [quest]; Black Widow Band (6199, -1.09 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (130.6 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.33 DPS) [dungeon]; Searing Golden Blade (12260, -4.24 DPS) [crafted] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.48 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Noh'Orahil (15107, -0.19 DPS) [quest]; Orb of Mystic Insight (249394, -0.75 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 162.4 spell_power points (40.37 DPS) | yes | Umbral Wand (5216, -4.70 DPS) [dungeon]; Earthen Rod (9381, -4.78 DPS) [dungeon]; Twisted Nether Wand (249144, -4.88 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 324, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 231.8. Weights run: 2.0s. Verify run: 1.5s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.365 ± 0.023, crit=0.151 ± 0.005 per rating point (14 rating = 1%, 2.109 per %), hit=0.406 ± 0.024 per rating point (10 rating = 1%, 4.059 per %), spell_haste=not significant (0.459 ± 0.310), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.293 ± 0.000, fire_power=0.707 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.35 DPS) | yes | Red Mageweave Headband (10033, -0.19 DPS) [crafted]; Dreamweave Circlet (10041, -0.64 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.63 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.9 spell_power points (2.70 DPS) | yes | Mindburst Medallion (11196, -0.47 DPS) [quest]; Scorn's Icy Choker (23169, -1.17 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.31 DPS) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Netherflame Shoulders (254053, -1.04 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.46 DPS) [vendor]; Rotgrip Mantle (17732, -2.53 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.2 spell_power points (4.41 DPS) | yes | Deep Woodlands Cloak (19121, -0.25 DPS) [quest]; Mantle of Lady Falther'ess (23178, -1.06 DPS) [dungeon]; Cindercloth Cloak (14044, -1.11 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 26.3 spell_power points (7.16 DPS) | yes | Robe of the Magi (1716, -0.58 DPS) [world_drop]; Dreamweave Vest (10021, -1.37 DPS) [crafted]; Runecloth Tunic (13857, -1.44 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 11.7 spell_power points (3.20 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.5 spell_power points (5.30 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -0.86 DPS) [vendor]; Runecloth Gloves (13863, -1.14 DPS) [crafted]; Fiery Gloves (254099, -1.33 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 20.1 spell_power points (5.47 DPS) | yes | Fiery Waistcord (254085, -0.91 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.26 DPS) [rep]; Satyrmane Sash (17755, -3.07 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.7 spell_power points (7.26 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -2.22 DPS) [vendor]; Red Mageweave Pants (10009, -2.25 DPS) [crafted]; Wizardweave Leggings (14132, -5.24 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.53 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.36 DPS) [vendor]; Cindercloth Boots (10044, -2.49 DPS) [crafted]; Fiery Sandals (254111, -3.14 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.54 DPS) | yes | Philanthropist's Ring (281635, -0.32 DPS) [quest]; Cyclopean Band (11824, -0.39 DPS) [dungeon]; Runed Ring (862, -1.63 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.27 DPS) | yes | Philanthropist's Ring (281635, -0.05 DPS) [quest]; Cyclopean Band (11824, -0.12 DPS) [dungeon]; Runed Ring (862, -1.36 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.63 DPS) [world_drop]; Rune of the Guard Captain (19120, -2.49 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.22 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.54 DPS) [quest]; Spire of Hakkar (10844, -0.65 DPS) [world]; Blade of Eternal Darkness (17780, -21.19 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.94 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.43 DPS) [crafted]; Pyric Caduceus (11748, -7.60 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 437.8. Weights run: 2.2s. Verify run: 3.5s. 1052 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.493 ± 0.033, crit=0.219 ± 0.006 per rating point (14 rating = 1%, 3.068 per %), hit=0.627 ± 0.034 per rating point (10 rating = 1%, 6.265 per %), spell_haste=not significant (-1.774 ± 0.541), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.386 ± 0.001, fire_power=0.614 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.37 DPS) [pvp]; Deathmist Mask (226909, -6.23 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.12 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.22 DPS) [quest]; Diana's Pearl Necklace (22403, -1.03 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.5 spell_power points (13.83 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.51 DPS) [pvp]; Burial Shawl (18681, -2.04 DPS, sim-verified) [dungeon]; Mantle of the Timbermaw (19050, -4.06 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.2 spell_power points (9.67 DPS) | yes | Crystalline Threaded Cape (20697, -1.56 DPS) [world]; Hide of the Wild (18510, -2.69 DPS) [crafted]; Amplifying Cloak (18350, -3.03 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.4 spell_power points (18.62 DPS) | yes | Warlord's Dreadweave Robe (231591, -2.44 DPS) [pvp]; Robe of Everlasting Night (18385, -3.56 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -5.75 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.9 spell_power points (9.58 DPS) | yes | Sublime Wristguards (18497, -3.33 DPS) [dungeon]; Runecloth Cuffs (254123, -3.70 DPS) [crafted]; Deathmist Bracers (226907, -4.81 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.5 spell_power points (10.88 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -1.40 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 42.1 spell_power points (15.55 DPS) | yes | Belt of the Archmage (18405, -3.81 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.80 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -7.08 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 43.1 spell_power points (15.90 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.89 DPS) [dungeon]; Outrider's Silk Leggings (22747, -2.10 DPS) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.9 spell_power points (9.57 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.01 DPS) [dungeon]; Blood Guard's Dreadweave Walkers (227098, -0.56 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.07 DPS) [dungeon]; Maiden's Circle (13001, -6.37 DPS) [world_drop]; Naglering (11669, -12.39 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.91 DPS) [dungeon]; Maiden's Circle (13001, -2.20 DPS) [world_drop]; Naglering (11669, -5.32 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+14.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.21 DPS) [quest]; Weakness Analyzer (272438, -2.58 DPS) [vendor]; Serenity Field (272439, -5.54 DPS) [vendor] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.21 DPS) [quest]; Teebu's Blazing Longsword (1728, -28.66 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+28.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.49 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.10 DPS) [world]; Torch of Light (279246, -28.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1052, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 842.0. Weights run: 7.1s. Verify run: 3.9s. 1052 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.093 ± 0.025, crit=0.275 ± 0.009 per rating point (14 rating = 1%, 3.847 per %), hit=0.832 ± 0.052 per rating point (10 rating = 1%, 8.320 per %), spell_haste=not significant (-0.579 ± 0.655), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.305 ± 0.001, fire_power=0.695 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.7 spell_power points (16.04 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -1.95 DPS) [crafted]; Deathmist Mask (226909, -5.41 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (11.48 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.05 DPS) [dungeon]; Amulet of the Dawn (22657, -3.02 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 32.2 spell_power points (16.82 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.95 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.31 DPS) [crafted]; Argent Shoulders (19059, -11.00 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.1 spell_power points (13.08 DPS) | yes | Amplifying Cloak (18350, -3.69 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.42 DPS, sim-verified) [world]; Hide of the Wild (18510, -5.29 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.8 spell_power points (24.43 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.57 DPS) [pvp]; Robe of Everlasting Night (18385, -6.74 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -10.42 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.7 spell_power points (11.87 DPS) | yes | Sublime Wristguards (18497, -5.12 DPS) [dungeon]; Runecloth Cuffs (254123, -5.64 DPS) [crafted]; Netherflame Cuffs (254065, -6.81 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.5 spell_power points (14.33 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -0.47 DPS) [dungeon]; Inferno Gloves (18408, -1.92 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 35.4 spell_power points (18.45 DPS) | yes | Ban'thok Sash (11662, -7.32 DPS) [dungeon]; Belt of the Archmage (18405, -7.91 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -8.58 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.6 spell_power points (21.19 DPS) | yes | General's Dreadweave Pants (231588, -0.97 DPS) [pvp]; Skyshroud Leggings (13170, -3.07 DPS) [dungeon]; Outrider's Silk Leggings (22747, -5.66 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.52 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.04 DPS) [crafted]; Omnicast Boots (11822, -1.50 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Elemental Focus Band (20682, -7.37 DPS) [world]; Maiden's Circle (13001, -9.65 DPS) [world_drop]; Naglering (11669, -23.86 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Elemental Focus Band (20682, -0.68 DPS) [world]; Maiden's Circle (13001, -2.96 DPS) [world_drop]; Naglering (11669, -16.84 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+21.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -3.13 DPS) [quest]; Weakness Analyzer (272438, -3.65 DPS) [vendor]; Serenity Field (272439, -7.82 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.46 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -48.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (842.0 DPS) | yes | Bonecreeper Stylus (13938, -1.03 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.21 DPS) [world]; Torch of Light (279246, -39.26 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Rune Band of Wizardry; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Spire of Hakkar; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1052, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

