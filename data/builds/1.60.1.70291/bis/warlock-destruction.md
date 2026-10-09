# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.7. Weights run: 1.7s. Verify run: 1.0s. 151 eligible items had no known source.

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

Set DPS (verified): 86.0. Weights run: 1.7s. Verify run: 1.0s. 266 eligible items had no known source.

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

Set DPS (verified): 140.8. Weights run: 1.4s. Verify run: 1.1s. 346 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.597 ± 0.020, crit=0.116 ± 0.004 per rating point (14 rating = 1%, 1.629 per %), hit=0.289 ± 0.019 per rating point (10 rating = 1%, 2.894 per %), spell_haste=0.661 ± 0.160, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.283 ± 0.000, fire_power=0.717 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.58 DPS) | yes | Living Cowl (5608, -2.13 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -2.19 DPS) [dungeon]; Augural Shroud (2620, -2.43 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (3.13 DPS) | yes | Scorn's Icy Choker (23169, -1.26 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.54 DPS) [quest]; Darkspear Warding Pendant (272074, -2.02 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (3.92 DPS) | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest]; Green Silken Shoulders (7057, -1.17 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.4 spell_power points (3.82 DPS) | yes | Guardian Cloak (5965, -1.43 DPS) [crafted]; Icy Cloak (4327, -1.96 DPS) [crafted]; Long Silken Cloak (4326, -2.02 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.6 spell_power points (6.80 DPS) | yes | Dreamweave Vest (10021, -0.59 DPS) [crafted]; Robe of Power (7054, -1.17 DPS) [crafted]; Elemental Raiment (9434, -1.22 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.3 spell_power points (2.48 DPS) | yes | Arcane Runed Bracers (4744, -0.08 DPS) [quest]; Spidertank Oilrag (9448, -0.08 DPS) [dungeon]; Condor Bracers (15864, -0.62 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (5.42 DPS) | yes | Black Mageweave Gloves (10003, -1.43 DPS) [crafted]; Fiery Handwraps (254025, -1.64 DPS) [crafted]; Red Mageweave Gloves (10018, -2.36 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.4 spell_power points (4.36 DPS) | yes | Fiery Cord (254041, -0.42 DPS) [crafted]; Star Belt (4329, -0.90 DPS) [crafted]; Deathmage Sash (10771, -1.16 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (5.63 DPS) | yes | Flame Leggings (253991, -1.85 DPS) [crafted]; Abomination Skin Leggings (23173, -1.96 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.22 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.38 DPS) | yes | Fiery Slippers (254005, -3.16 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.41 DPS) [crafted]; Acidic Walkers (9454, -3.78 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (3.45 DPS) | yes | Ring of Forlorn Spirits (2043, -1.33 DPS) [quest]; Reedknot Ring (9622, -1.59 DPS) [quest]; Minor Channeling Ring (1449, -1.81 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.39 DPS) | yes | Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.53 DPS) [quest]; Minor Channeling Ring (1449, -0.75 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.32 DPS) [dungeon]; Searing Golden Blade (12260, -4.42 DPS) [crafted] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.72 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Noh'Orahil (15107, -0.26 DPS) [quest]; Orb of Lorica (11262, -0.54 DPS) [quest] |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (140.8 DPS) | yes | Twisted Nether Wand (249144, -0.08 DPS) [crafted]; Earthen Rod (9381, -0.08 DPS) [dungeon]; Jaina's Firestarter (13064, -1.83 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Umbral Wand

No-known-source sample (15 of 346, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 245.5. Weights run: 1.4s. Verify run: 1.4s. 436 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.218 ± 0.022, crit=0.168 ± 0.005 per rating point (14 rating = 1%, 2.352 per %), hit=0.436 ± 0.025 per rating point (10 rating = 1%, 4.364 per %), spell_haste=not significant (-0.048 ± 0.402), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.279 ± 0.000, fire_power=0.721 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.70 DPS) | yes | Red Mageweave Headband (10033, -1.04 DPS) [crafted]; Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.71 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.7 spell_power points (2.49 DPS) | yes | Scorn's Icy Choker (23169, -0.12 DPS) [dungeon]; Mindburst Medallion (11196, -0.41 DPS) [quest]; Horizon Choker (13085, -1.62 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (245.5 DPS) | yes | Kentic Amice (11624, -0.03 DPS) [dungeon]; Netherflame Shoulders (254053, -0.83 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.32 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.37 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.24 DPS) [dungeon]; Runecloth Cloak (13860, -1.30 DPS) [crafted]; Cindercloth Cloak (14044, -1.92 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.4 spell_power points (6.66 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.67 DPS) [world_drop]; Dreamweave Vest (10021, -0.97 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.9 spell_power points (3.11 DPS) | yes | Phoenix Bindings (210781, -0.44 DPS) [crafted]; Arcane Runed Bracers (4744, -0.54 DPS) [quest]; Spidertank Oilrag (9448, -0.54 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (5.38 DPS) | yes | Fiery Gloves (254099, -0.85 DPS) [crafted]; Black Mageweave Gloves (10003, -1.10 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.12 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.8 spell_power points (5.35 DPS) | yes | Fiery Waistcord (254085, -0.89 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.11 DPS) [rep]; Satyrmane Sash (17755, -2.93 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.2 spell_power points (7.18 DPS) | yes | Knight's Dreadweave Leggings (220888, -2.34 DPS) [vendor]; Red Mageweave Pants (10009, -2.44 DPS) [crafted]; Wizardweave Leggings (14132, -4.37 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.85 DPS) | yes | Fiery Sandals (254111, -1.96 DPS, sim-verified) [crafted]; Cindercloth Boots (10044, -2.53 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.76 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.71 DPS) | yes | Philanthropist's Ring (281635, -0.55 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.43 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.42 DPS) | yes | Philanthropist's Ring (281635, -0.26 DPS) [quest]; Cyclopean Band (11824, -0.42 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.14 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.71 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.18 DPS) [world]; Inventor's Focal Sword (17719, -0.75 DPS) [dungeon]; Blade of Eternal Darkness (17780, -21.90 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (245.5 DPS) | yes | Wand of Allistarj (13065, -3.01 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.39 DPS) [crafted]; Pyric Caduceus (11748, -12.08 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 474.4. Weights run: 1.6s. Verify run: 3.1s. 1064 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.378 ± 0.029, crit=0.209 ± 0.006 per rating point (14 rating = 1%, 2.924 per %), hit=0.586 ± 0.033 per rating point (10 rating = 1%, 5.864 per %), spell_haste=-2.575 ± 0.475, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.378 ± 0.001, fire_power=0.622 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.0 spell_power points (12.99 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.90 DPS) [pvp]; Deathmist Mask (226909, -2.47 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.65 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.82 DPS) [quest]; Diana's Pearl Necklace (22403, -1.62 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.6 spell_power points (14.00 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.64 DPS) [pvp]; Burial Shawl (18681, -3.75 DPS) [dungeon]; Argent Shoulders (19059, -4.17 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.9 spell_power points (9.79 DPS) | yes | Crystalline Threaded Cape (20697, -1.33 DPS) [world]; Amplifying Cloak (18350, -2.71 DPS) [dungeon]; Hide of the Wild (18510, -2.79 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.4 spell_power points (19.43 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.27 DPS) [pvp]; Robe of Everlasting Night (18385, -3.43 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -6.62 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.0 spell_power points (9.84 DPS) | yes | Sublime Wristguards (18497, -3.63 DPS) [dungeon]; Runecloth Cuffs (254123, -4.03 DPS) [crafted]; Deathmist Bracers (226907, -5.30 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.70 DPS) [crafted]; Sandworm Skin Gloves (20716, -4.68 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.2 spell_power points (15.41 DPS) | yes | Belt of the Archmage (18405, -4.49 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.75 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -6.84 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 41.6 spell_power points (16.37 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.81 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -3.43 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.5 spell_power points (9.65 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Dragonrider Boots (18102, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.06 DPS) [dungeon]; Maiden's Circle (13001, -6.53 DPS) [world_drop]; Naglering (11669, -16.70 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.69 DPS) [dungeon]; Maiden's Circle (13001, -2.17 DPS) [world_drop]; Naglering (11669, -9.56 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.36 DPS) [quest]; Weakness Analyzer (272438, -2.75 DPS) [vendor]; Draconic Infused Emblem (22268, -15.89 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (474.4 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.72 DPS) [quest]; Teebu's Blazing Longsword (1728, -34.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+34.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.62 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.54 DPS) [world]; Torch of Light (279246, -34.69 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1064, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 886.8. Weights run: 1.7s. Verify run: 3.3s. 1064 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.135 ± 0.034, crit=0.283 ± 0.008 per rating point (14 rating = 1%, 3.962 per %), hit=0.875 ± 0.050 per rating point (10 rating = 1%, 8.752 per %), spell_haste=not significant (-0.677 ± 0.659), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.295 ± 0.001, fire_power=0.705 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.47 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (12.46 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS, sim-verified) [quest]; Diana's Pearl Necklace (22403, -1.79 DPS) [dungeon]; Amulet of the Dawn (22657, -2.97 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 33.0 spell_power points (18.68 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -3.22 DPS) [pvp]; Argent Shoulders (19059, -4.63 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -5.82 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.8 spell_power points (14.63 DPS) | yes | Crystalline Threaded Cape (20697, -3.00 DPS) [world]; Amplifying Cloak (18350, -4.43 DPS) [dungeon]; Hide of the Wild (18510, -5.94 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.2 spell_power points (26.73 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.78 DPS) [pvp]; Robe of Everlasting Night (18385, -10.45 DPS) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -11.05 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.1 spell_power points (13.07 DPS) | yes | Sublime Wristguards (18497, -5.51 DPS) [dungeon]; Runecloth Cuffs (254123, -6.08 DPS) [crafted]; Netherflame Cuffs (254065, -7.34 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (15.67 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.80 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 36.7 spell_power points (20.79 DPS) | yes | Ban'thok Sash (11662, -8.20 DPS) [dungeon]; Belt of the Archmage (18405, -9.73 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -9.84 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 41.3 spell_power points (23.37 DPS) | yes | Marshal's Dreadweave Leggings (231587, -0.97 DPS) [pvp]; Skyshroud Leggings (13170, -3.51 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -6.52 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.59 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.13 DPS) [crafted]; Omnicast Boots (11822, -1.35 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -7.48 DPS) [dungeon]; Maiden's Circle (13001, -10.77 DPS) [world_drop]; Naglering (11669, -29.96 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, +0.00 DPS) [dungeon]; Songstone of Ironforge (12543, -2.57 DPS) [quest]; Maiden's Circle (13001, -2.57 DPS) [world_drop] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -3.40 DPS) [quest]; Weakness Analyzer (272438, -3.96 DPS) [vendor]; Draconic Infused Emblem (22268, -22.84 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (886.8 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.22 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -49.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+56.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.91 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.23 DPS) [world]; Torch of Light (279246, -56.05 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Spire of Hakkar; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1064, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.4. Weights run: 1.7s. Verify run: 1.0s. 140 eligible items had no known source.

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

Set DPS (verified): 84.4. Weights run: 1.7s. Verify run: 1.0s. 249 eligible items had no known source.

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

Set DPS (verified): 138.0. Weights run: 1.4s. Verify run: 1.2s. 324 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.597 ± 0.020, crit=0.116 ± 0.004 per rating point (14 rating = 1%, 1.629 per %), hit=0.289 ± 0.019 per rating point (10 rating = 1%, 2.894 per %), spell_haste=0.661 ± 0.160, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.283 ± 0.000, fire_power=0.717 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.58 DPS) | yes | Living Cowl (5608, -2.13 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -2.19 DPS) [dungeon]; Augural Shroud (2620, -2.38 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (3.13 DPS) | yes | Scorn's Icy Choker (23169, -1.18 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.54 DPS) [quest]; Darkspear Warding Pendant (272074, -2.02 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (3.92 DPS) | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest]; Green Silken Shoulders (7057, -1.09 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (138.0 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Mantle of Lady Falther'ess (23178, +0.00 DPS) [dungeon]; Icy Cloak (4327, -0.53 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.6 spell_power points (6.80 DPS) | yes | Dreamweave Vest (10021, -0.59 DPS) [crafted]; Robe of Power (7054, -1.17 DPS) [crafted]; Elemental Raiment (9434, -1.22 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.3 spell_power points (2.48 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.14 DPS) [quest]; Spidertank Oilrag (9448, -0.65 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (5.42 DPS) | yes | Black Mageweave Gloves (10003, -1.43 DPS) [crafted]; Fiery Handwraps (254025, -1.64 DPS) [crafted]; Red Mageweave Gloves (10018, -2.58 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.4 spell_power points (4.36 DPS) | yes | Fiery Cord (254041, -0.42 DPS) [crafted]; Star Belt (4329, -0.90 DPS) [crafted]; Deathmage Sash (10771, -1.06 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (5.63 DPS) | yes | Flame Leggings (253991, -1.85 DPS) [crafted]; Abomination Skin Leggings (23173, -1.96 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.99 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.38 DPS) | yes | Gilded Slippers (254001, -3.41 DPS) [crafted]; Fiery Slippers (254005, -3.60 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.78 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (3.45 DPS) | yes | Reedknot Ring (9622, -1.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.86 DPS) [vendor]; Black Widow Band (6199, -2.34 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.39 DPS) | yes | Reedknot Ring (9622, -0.63 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.80 DPS) [vendor]; Black Widow Band (6199, -1.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.32 DPS) [dungeon]; Searing Golden Blade (12260, -4.42 DPS) [crafted] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.72 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Noh'Orahil (15107, -0.26 DPS) [quest]; Orb of Mystic Insight (249394, -0.91 DPS) [crafted] |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Twisted Nether Wand (249144, -0.08 DPS) [crafted]; Earthen Rod (9381, -0.08 DPS) [dungeon]; Jaina's Firestarter (13064, -1.55 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Umbral Wand

No-known-source sample (15 of 324, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 243.7. Weights run: 1.4s. Verify run: 1.3s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.218 ± 0.022, crit=0.168 ± 0.005 per rating point (14 rating = 1%, 2.352 per %), hit=0.436 ± 0.025 per rating point (10 rating = 1%, 4.364 per %), spell_haste=not significant (-0.048 ± 0.402), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.279 ± 0.000, fire_power=0.721 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.70 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Red Mageweave Headband (10033, -1.23 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.71 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.7 spell_power points (2.49 DPS) | yes | Mindburst Medallion (11196, -0.41 DPS) [quest]; Horizon Choker (13085, -1.62 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.72 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.9 spell_power points (4.83 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Netherflame Shoulders (254053, -0.83 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.32 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.37 DPS) | yes | Cindercloth Cloak (14044, -1.20 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.24 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.49 DPS, sim-verified) [quest] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.4 spell_power points (6.66 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.67 DPS) [world_drop]; Dreamweave Vest (10021, -0.97 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.9 spell_power points (3.11 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Phoenix Bindings (210781, -0.88 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (5.38 DPS) | yes | Fiery Gloves (254099, -1.04 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.10 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.12 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.8 spell_power points (5.35 DPS) | yes | Fiery Waistcord (254085, -0.89 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.11 DPS) [rep]; Satyrmane Sash (17755, -3.06 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.2 spell_power points (7.18 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -2.34 DPS) [vendor]; Red Mageweave Pants (10009, -2.44 DPS) [crafted]; Wizardweave Leggings (14132, -4.33 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.85 DPS) | yes | Cindercloth Boots (10044, -2.53 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.76 DPS) [vendor]; Fiery Sandals (254111, -2.98 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.71 DPS) | yes | Philanthropist's Ring (281635, -0.55 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon]; Runed Ring (862, -1.71 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.42 DPS) | yes | Philanthropist's Ring (281635, -0.26 DPS) [quest]; Cyclopean Band (11824, -0.42 DPS) [dungeon]; Runed Ring (862, -1.43 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.71 DPS) [world_drop]; Rune of the Guard Captain (19120, -2.55 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.25 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.18 DPS) [world]; Inventor's Focal Sword (17719, -0.75 DPS) [dungeon]; Blade of Eternal Darkness (17780, -22.20 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (243.7 DPS) | yes | Wand of Allistarj (13065, -3.01 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.39 DPS) [crafted]; Pyric Caduceus (11748, -11.06 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 462.1. Weights run: 1.6s. Verify run: 2.9s. 1052 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.378 ± 0.029, crit=0.209 ± 0.006 per rating point (14 rating = 1%, 2.924 per %), hit=0.586 ± 0.033 per rating point (10 rating = 1%, 5.864 per %), spell_haste=-2.575 ± 0.475, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.378 ± 0.001, fire_power=0.622 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.87 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.65 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.82 DPS) [quest]; Diana's Pearl Necklace (22403, -1.62 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.6 spell_power points (14.00 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.64 DPS) [pvp]; Burial Shawl (18681, -2.96 DPS, sim-verified) [dungeon]; Argent Shoulders (19059, -4.17 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.9 spell_power points (9.79 DPS) | yes | Crystalline Threaded Cape (20697, -1.33 DPS) [world]; Amplifying Cloak (18350, -2.71 DPS) [dungeon]; Hide of the Wild (18510, -2.79 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.4 spell_power points (19.43 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.27 DPS) [pvp]; Robe of Everlasting Night (18385, -4.64 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -6.62 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.0 spell_power points (9.84 DPS) | yes | Sublime Wristguards (18497, -3.63 DPS) [dungeon]; Runecloth Cuffs (254123, -4.03 DPS) [crafted]; Deathmist Bracers (226907, -5.30 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.9 spell_power points (11.36 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.94 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.2 spell_power points (15.41 DPS) | yes | Belt of the Archmage (18405, -4.62 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.75 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -6.84 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 41.6 spell_power points (16.37 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.81 DPS) [dungeon]; Outrider's Silk Leggings (22747, -2.53 DPS) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.5 spell_power points (9.65 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.19 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.21 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.06 DPS) [dungeon]; Maiden's Circle (13001, -6.53 DPS) [world_drop]; Naglering (11669, -15.25 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.69 DPS) [dungeon]; Maiden's Circle (13001, -2.17 DPS) [world_drop]; Naglering (11669, -9.16 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.36 DPS) [quest]; Weakness Analyzer (272438, -2.75 DPS) [vendor]; Draconic Infused Emblem (22268, -15.63 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (462.1 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.72 DPS) [quest]; Teebu's Blazing Longsword (1728, -32.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+35.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.62 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.54 DPS) [world]; Torch of Light (279246, -35.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1052, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 879.7. Weights run: 1.7s. Verify run: 3.3s. 1052 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.135 ± 0.034, crit=0.283 ± 0.008 per rating point (14 rating = 1%, 3.962 per %), hit=0.875 ± 0.050 per rating point (10 rating = 1%, 8.752 per %), spell_haste=not significant (-0.677 ± 0.659), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.295 ± 0.001, fire_power=0.705 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.47 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (12.46 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS, sim-verified) [quest]; Diana's Pearl Necklace (22403, -1.79 DPS) [dungeon]; Amulet of the Dawn (22657, -2.97 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 33.0 spell_power points (18.68 DPS) | yes | Warlord's Dreadweave Mantle (231592, -3.22 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.82 DPS) [crafted]; Argent Shoulders (19059, -6.83 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.8 spell_power points (14.63 DPS) | yes | Crystalline Threaded Cape (20697, -3.00 DPS) [world]; Amplifying Cloak (18350, -4.43 DPS) [dungeon]; Hide of the Wild (18510, -5.94 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.2 spell_power points (26.73 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.78 DPS) [pvp]; Robe of Everlasting Night (18385, -10.45 DPS) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -11.05 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.1 spell_power points (13.07 DPS) | yes | Sublime Wristguards (18497, -5.51 DPS) [dungeon]; Runecloth Cuffs (254123, -6.08 DPS) [crafted]; Netherflame Cuffs (254065, -7.34 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -0.49 DPS) [dungeon]; Inferno Gloves (18408, -1.80 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 36.7 spell_power points (20.79 DPS) | yes | Belt of the Archmage (18405, -3.12 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -8.20 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -9.84 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 41.3 spell_power points (23.37 DPS) | yes | General's Dreadweave Pants (231588, -0.97 DPS) [pvp]; Skyshroud Leggings (13170, -3.51 DPS) [dungeon]; Outrider's Silk Leggings (22747, -6.06 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.59 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.13 DPS) [crafted]; Omnicast Boots (11822, -1.35 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -7.48 DPS) [dungeon]; Maiden's Circle (13001, -10.77 DPS) [world_drop]; Naglering (11669, -22.40 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, +0.00 DPS) [dungeon]; Eye of Orgrimmar (12545, -2.57 DPS) [quest]; Maiden's Circle (13001, -2.57 DPS) [world_drop] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -3.40 DPS) [quest]; Weakness Analyzer (272438, -3.96 DPS) [vendor]; Draconic Infused Emblem (22268, -22.72 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (879.7 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.22 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -47.88 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+57.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.91 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.23 DPS) [world]; Torch of Light (279246, -57.72 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Spire of Hakkar; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1052, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

