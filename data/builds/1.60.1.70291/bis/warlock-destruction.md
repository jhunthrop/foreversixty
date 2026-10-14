# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.7. Weights run: 1.8s. Verify run: 1.5s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.062 ± 0.004, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.432 per %), hit=0.125 ± 0.005 per rating point (10 rating = 1%, 1.252 per %), spell_haste=0.602 ± 0.065, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.764 ± 0.002, fire_power=0.236 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.20 DPS) | yes | Shadow Goggles (4373, -1.13 DPS) [crafted]; Flame Circlet (253953, -2.39 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.6 spell_power points (1.11 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.19 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.31 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Caretaker's Cape (20428, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.24 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.3 spell_power points (1.06 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -1.39 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.14 DPS) [dungeon]; Bright Bracers (3647, -0.15 DPS) [world_drop]; Repurposed Hair Band (281256, -0.17 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.39 DPS) | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.56 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.97 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.2 spell_power points (0.85 DPS) | yes | Novice Ardent's Sash (253887, -0.41 DPS) [crafted]; Flame Sash (253929, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.96 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 spell_power points (1.89 DPS) | yes | Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted]; Silk-threaded Trousers (1929, -0.66 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.2 spell_power points (1.44 DPS) | yes | Pristine Boots (253889, -0.81 DPS) [crafted]; Red Woolen Boots (4313, -0.83 DPS, sim-verified) [crafted]; Woolen Boots (2583, -1.02 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.1 spell_power points (1.02 DPS) | yes | Lavishly Jeweled Ring (1156, -0.95 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.98 DPS) [world_drop]; The 1 Ring (8350, -1.01 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.00 DPS) | yes | Volcanic Rock Ring (12053, -0.96 DPS) [world_drop]; The 1 Ring (8350, -0.98 DPS) [world]; Lavishly Jeweled Ring (1156, -1.28 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.6 spell_power points (0.12 DPS) | yes | Witching Stave (1484, -0.06 DPS) [dungeon]; Staff of Westfall (2042, -0.06 DPS) [quest]; Lesser Staff of the Spire (1300, -0.18 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 115.0 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.58 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.26 DPS) [quest]; Deepblaze (279896, -4.38 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053225101000000)

Set DPS (verified): 83.1. Weights run: 2.0s. Verify run: 1.8s. 260 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.338 ± 0.012, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.004 per %), hit=0.168 ± 0.011 per rating point (10 rating = 1%, 1.683 per %), spell_haste=not significant (0.095 ± 0.167), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.767 ± 0.002, fire_power=0.233 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.55 DPS) | yes | Silk Headband (7050, -0.46 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.70 DPS) [crafted]; Enchanter's Cowl (4322, -4.03 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.09 DPS) | yes | Crystal Starfire Medallion (5003, -1.78 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.78 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.85 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (2.79 DPS) | yes | Fairywing Mantle (9536, -0.70 DPS) [quest]; Invoker's Mantle (215365, -0.78 DPS) [crafted]; Death Speaker Mantle (6685, -1.09 DPS, sim-verified) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (1.39 DPS) | yes | Hillman's Cloak (3719, -0.23 DPS) [crafted]; Repairman's Cape (9605, -0.38 DPS) [quest]; Heavy Woolen Cloak (4311, -0.46 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.4 spell_power points (3.10 DPS) | yes | Tree Bark Jacket (1486, -0.09 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.31 DPS) [dungeon]; Beguiler Robes (7728, -0.39 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.09 DPS) | yes | Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop]; Nightsky Wristbands (6407, -1.61 DPS) [world_drop]; Phoenix Bindings (210781, -2.29 DPS, sim-verified) [crafted] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | sim-verified (83.1 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Truefaith Gloves (7049, -0.23 DPS) [crafted]; Town Clerk's Mittens (270029, -0.89 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.0 spell_power points (2.78 DPS) | yes | Belt of Arugal (6392, -0.68 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.77 DPS) [crafted]; Crimson Silk Belt (7055, -0.84 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.78 DPS) | yes | Pristine Leggings (253987, -0.61 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.92 DPS) [crafted]; Abomination Skin Leggings (23173, -1.05 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (2.17 DPS) | yes | Acidic Walkers (9454, -0.39 DPS) [dungeon]; Nimbus Boots (6998, -0.78 DPS) [quest]; Spidersilk Boots (4320, -1.99 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.62 DPS) | yes | Minor Channeling Ring (1449, -0.31 DPS) [quest]; Electrocutioner Lagnut (9447, -0.93 DPS) [dungeon]; Snake Hoop (6750, -1.07 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.39 DPS) | yes | Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon]; Snake Hoop (6750, -0.84 DPS) [quest]; Minor Channeling Ring (1449, -2.15 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | 30.0 spell_power points (6.95 DPS) | yes | Glimmering Staff (249392, -6.09 DPS) [crafted]; Twisted Chanter's Staff (890, -6.17 DPS) [world_drop]; Staff of Soran'ruk (15109, -9.30 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 143.2 spell_power points (33.17 DPS) | yes | Scorching Wand (5213, -4.00 DPS) [world_drop]; Gravestone Scepter (7001, -4.17 DPS) [quest]; Greater Mystic Wand (217287, -4.77 DPS, sim-verified) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff; ranged: Starfaller

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053225103101321)

Set DPS (verified): 138.5. Weights run: 1.8s. Verify run: 2.2s. 334 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.621 ± 0.029, crit=0.119 ± 0.006 per rating point (14 rating = 1%, 1.670 per %), hit=0.294 ± 0.028 per rating point (10 rating = 1%, 2.942 per %), spell_haste=not significant (0.329 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.282 ± 0.001, fire_power=0.718 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.58 DPS) | yes | Augural Shroud (2620, -1.74 DPS, sim-verified) [world]; Living Cowl (5608, -2.12 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -2.13 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.0 spell_power points (3.18 DPS) | yes | Scorn's Icy Choker (23169, -0.94 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.02 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.02 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.1 spell_power points (4.00 DPS) | yes | Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest]; Green Silken Shoulders (7057, -0.85 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.42 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.56 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.60 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (6.83 DPS) | yes | Dreamweave Vest (10021, -0.57 DPS) [crafted]; Robe of Power (7054, -1.13 DPS) [crafted]; Elemental Raiment (9434, -1.26 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.3 spell_power points (2.48 DPS) | yes | Arcane Runed Bracers (4744, -0.09 DPS) [quest]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon]; Condor Bracers (15864, -0.62 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (5.44 DPS) | yes | Black Mageweave Gloves (10003, -1.46 DPS) [crafted]; Fiery Handwraps (254025, -1.62 DPS) [crafted]; Red Mageweave Gloves (10018, -2.20 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.5 spell_power points (4.38 DPS) | yes | Fiery Cord (254041, -0.39 DPS) [crafted]; Star Belt (4329, -0.93 DPS) [crafted]; Deathmage Sash (10771, -1.20 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.5 spell_power points (5.70 DPS) | yes | Crimson Silk Pantaloons (7062, -1.45 DPS, sim-verified) [crafted]; Flame Leggings (253991, -1.87 DPS) [crafted]; Abomination Skin Leggings (23173, -1.99 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.37 DPS) | yes | Gilded Slippers (254001, -3.36 DPS) [crafted]; Acidic Walkers (9454, -3.73 DPS) [dungeon]; Fiery Slippers (254005, -4.15 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.48 DPS) | yes | Ring of Forlorn Spirits (2043, -1.36 DPS) [quest]; Reedknot Ring (9622, -1.62 DPS) [quest]; Minor Channeling Ring (1449, -1.82 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.39 DPS) | yes | Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.53 DPS) [quest]; Minor Channeling Ring (1449, -0.73 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.33 DPS) [dungeon]; Searing Golden Blade (12260, -4.41 DPS) [crafted] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.72 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Noh'Orahil (15107, -0.22 DPS) [quest]; Orb of Lorica (11262, -0.47 DPS) [quest] |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (138.5 DPS) | yes | Twisted Nether Wand (249144, -0.08 DPS) [crafted]; Burning Sliver (5249, -1.39 DPS) [quest]; Jaina's Firestarter (13064, -1.53 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Umbral Wand

No-known-source sample (15 of 334, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 236.6. Weights run: 1.9s. Verify run: 2.3s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.202 ± 0.031, crit=0.168 ± 0.007 per rating point (14 rating = 1%, 2.346 per %), hit=0.447 ± 0.035 per rating point (10 rating = 1%, 4.472 per %), spell_haste=not significant (-0.308 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.280 ± 0.001, fire_power=0.720 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.71 DPS) | yes | Red Mageweave Headband (10033, -0.86 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.14 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.71 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.6 spell_power points (2.46 DPS) | yes | Mindburst Medallion (11196, -0.40 DPS) [quest]; Horizon Choker (13085, -1.65 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.83 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.6 spell_power points (4.75 DPS) | yes | Netherflame Shoulders (254053, -0.79 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.28 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.38 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 spell_power points (4.34 DPS) | yes | Runecloth Cloak (13860, -1.31 DPS) [crafted]; Cloak of Fire (14134, -1.47 DPS) [crafted]; Cindercloth Cloak (14044, -2.47 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 spell_power points (6.63 DPS) | yes | Acumen Robes (17775, -0.05 DPS) [quest]; Elemental Raiment (9434, -0.63 DPS) [world_drop]; Dreamweave Vest (10021, -0.97 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.8 spell_power points (3.08 DPS) | yes | Phoenix Bindings (210781, -0.40 DPS) [crafted]; Arcane Runed Bracers (4744, -0.51 DPS) [quest]; Spidertank Oilrag (9448, -0.51 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.37 DPS) | yes | Black Mageweave Gloves (10003, -1.09 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.14 DPS) [vendor]; Fiery Gloves (254099, -1.37 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.0 spell_power points (4.57 DPS) | yes | Fiery Waistcord (254085, -0.15 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.35 DPS) [rep]; Ghostweave Cord (254073, -0.58 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.43 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.64 DPS) [vendor]; Red Mageweave Pants (10009, -0.74 DPS) [crafted]; Cindercloth Leggings (12256, -1.52 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.85 DPS) | yes | Cindercloth Boots (10044, -2.54 DPS) [crafted]; Fiery Sandals (254111, -2.71 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.77 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.71 DPS) | yes | Philanthropist's Ring (281635, -0.57 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.43 DPS) [quest]; Reedknot Ring (9622, -1.71 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.43 DPS) | yes | Ring of Forlorn Spirits (2043, -1.14 DPS) [quest]; Philanthropist's Ring (281635, -1.35 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.43 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.71 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.12 DPS) [world]; Inventor's Focal Sword (17719, -0.69 DPS) [dungeon]; Hanzo Sword (8190, -27.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (236.6 DPS) | yes | Wand of Allistarj (13065, -3.01 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.39 DPS) [crafted]; Pyric Caduceus (11748, -9.87 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 469.6. Weights run: 1.9s. Verify run: 5.0s. 867 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.330 ± 0.036, crit=0.200 ± 0.008 per rating point (14 rating = 1%, 2.800 per %), hit=0.500 ± 0.039 per rating point (10 rating = 1%, 5.004 per %), spell_haste=not significant (0.991 ± 0.538), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.350 ± 0.001, fire_power=0.650 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | 30.9 spell_power points (13.16 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.50 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.65 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.36 DPS) | yes | Amulet of the Dawn (22657, -1.15 DPS) [quest]; Beads of Ogre Mojo (22149, -2.15 DPS) [quest]; Orb of the Darkmoon (19426, -2.21 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.7 spell_power points (14.79 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.76 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.54 DPS) [crafted]; Argent Shoulders (19059, -4.78 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.6 spell_power points (10.06 DPS) | yes | Hide of the Wild (18510, -2.03 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -3.26 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -4.10 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.0 spell_power points (20.85 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.85 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.39 DPS) [pvp]; Acumen Robes (17775, -10.09 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.6 spell_power points (10.49 DPS) | yes | Runecloth Cuffs (254123, -4.40 DPS) [crafted]; Deathmist Bracers (226907, -5.82 DPS) [quest]; Netherflame Cuffs (254065, -5.91 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.6 spell_power points (12.20 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -2.69 DPS) [pvp]; Inferno Gloves (18408, -4.60 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.3 spell_power points (15.86 DPS) | yes | Belt of the Archmage (18405, -6.27 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -6.80 DPS) [rep]; Highlander's Cloth Girdle (20047, -7.87 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.9 spell_power points (17.41 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -3.67 DPS) [pvp]; Ghostweave Pants (14144, -6.34 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.22 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.02 DPS) [pvp]; Omnicast Boots (11822, -1.97 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 36.0 spell_power points (15.31 DPS) | yes | Deep Rooted Ring (19109, -8.80 DPS) [world]; Blessed Band of Light (272407, -8.93 DPS) [vendor]; Lorekeeper's Ring (19522, -9.35 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 25.6 spell_power points (10.91 DPS) | yes | Deep Rooted Ring (19109, -4.04 DPS, sim-verified) [world]; Blessed Band of Light (272407, -4.52 DPS) [vendor]; Lorekeeper's Ring (19522, -4.95 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+12.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -0.85 DPS) [quest]; Weakness Analyzer (272438, -2.51 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.41 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.83 DPS) [world]; Teebu's Blazing Longsword (1728, -36.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (469.6 DPS) | yes | Wand of Biting Cold (19108, -0.39 DPS) [quest]; Stormrager (16997, -1.00 DPS) [quest]; Torch of Light (279246, -40.01 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 867, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 879.5. Weights run: 4.2s. Verify run: 4.9s. 867 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.060 ± 0.033, crit=0.256 ± 0.010 per rating point (14 rating = 1%, 3.588 per %), hit=0.738 ± 0.067 per rating point (10 rating = 1%, 7.376 per %), spell_haste=not significant (1.047 ± 0.857), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.279 ± 0.001, fire_power=0.721 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.10 DPS) | yes | Deathmist Mask (226909, +0.00 DPS, sim-verified) [quest]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.79 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.12 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.71 DPS) [quest]; Kezan's Taint (19604, -4.48 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.5 spell_power points (18.78 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -3.26 DPS) [pvp]; Argent Shoulders (19059, -5.21 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -6.03 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.9 spell_power points (14.22 DPS) | yes | Hide of the Wild (18510, -5.52 DPS) [crafted]; Spritecaster Cape (11623, -5.66 DPS) [dungeon]; Mageflame Cloak (13007, -9.04 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (27.75 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.81 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.13 DPS) [pvp]; Robe of the Magi (1716, -13.06 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.5 spell_power points (13.40 DPS) | yes | Runecloth Cuffs (254123, -6.49 DPS) [crafted]; Netherflame Cuffs (254065, -7.56 DPS) [crafted]; Phoenix Bindings (210781, -7.81 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (16.28 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.76 DPS) [crafted]; Earth Warder's Gloves (21318, -3.16 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.7 spell_power points (20.09 DPS) | yes | Belt of the Archmage (18405, -8.58 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -9.00 DPS) [rep]; Highlander's Cloth Girdle (20047, -9.39 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 39.8 spell_power points (23.72 DPS) | yes | Marshal's Dreadweave Leggings (231587, -0.98 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -6.56 DPS) [pvp]; Ghostweave Pants (14144, -8.21 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.31 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.19 DPS) [crafted]; Omnicast Boots (11822, -1.96 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.9 spell_power points (21.42 DPS) | yes | Blessed Band of Light (272407, -11.05 DPS) [vendor]; Lorekeeper's Ring (19522, -13.07 DPS) [rep]; Band of the Unicorn (7553, -13.66 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.7 spell_power points (13.51 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Lorekeeper's Ring (19522, -5.16 DPS) [rep]; Band of the Unicorn (7553, -5.76 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+19.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Abyss Shard (20534, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -1.19 DPS) [quest]; Weakness Analyzer (272438, -3.51 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.39 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.76 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -48.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (879.5 DPS) | yes | Wand of Biting Cold (19108, -1.20 DPS) [quest]; Stormrager (16997, -2.33 DPS) [quest]; Torch of Light (279246, -58.14 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Brilliant Wand

No-known-source sample (15 of 867, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 38.6. Weights run: 1.8s. Verify run: 1.6s. 137 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.062 ± 0.004, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.432 per %), hit=0.125 ± 0.005 per rating point (10 rating = 1%, 1.252 per %), spell_haste=0.602 ± 0.065, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.764 ± 0.002, fire_power=0.236 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.20 DPS) | yes | Shadow Goggles (4373, -1.13 DPS) [crafted]; Flame Circlet (253953, -2.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.6 spell_power points (1.11 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.06 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.31 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Battle Healer's Cloak (20427, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.32 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.3 spell_power points (1.06 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -1.35 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.60 DPS) | yes | Windsong Bangles (263336, -0.40 DPS) [quest]; Featherbead Bracers (15452, -0.54 DPS) [quest]; Owlbeard Bracers (16981, -0.59 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.39 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.56 DPS) [crafted]; Apothecary Gloves (10919, -0.60 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.2 spell_power points (0.85 DPS) | yes | Novice Ardent's Sash (253887, -0.41 DPS) [crafted]; Flame Sash (253929, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.90 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 spell_power points (1.89 DPS) | yes | Silk-threaded Trousers (1929, -0.61 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.2 spell_power points (1.44 DPS) | yes | Red Woolen Boots (4313, -0.62 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.81 DPS) [crafted]; Woolen Boots (2583, -1.02 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.00 DPS) | yes | Loop of Sacrifice (281673, -0.93 DPS) [quest]; Volcanic Rock Ring (12053, -0.96 DPS) [world_drop]; The 1 Ring (8350, -0.98 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.4 spell_power points (0.07 DPS) | yes | Loop of Sacrifice (281673, -0.01 DPS) [quest]; Volcanic Rock Ring (12053, -0.04 DPS) [world_drop]; The 1 Ring (8350, -0.06 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.6 spell_power points (0.12 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.05 DPS) [world]; Clear Crystal Rod (16894, -0.06 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 115.0 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.51 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.26 DPS) [quest]; Greater Magic Wand (11288, -5.01 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 137, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest

### Band 30 (troll, 00000000000000000-0000000000000000000-2053225101000000)

Set DPS (verified): 72.7. Weights run: 2.0s. Verify run: 1.8s. 243 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.338 ± 0.012, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.004 per %), hit=0.168 ± 0.011 per rating point (10 rating = 1%, 1.683 per %), spell_haste=not significant (0.095 ± 0.167), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.767 ± 0.002, fire_power=0.233 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.55 DPS) | yes | Silk Headband (7050, -0.46 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.70 DPS) [crafted]; Enchanter's Cowl (4322, -3.77 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.09 DPS) | yes | Darkspear Warding Pendant (272075, -1.59 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.78 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.78 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (2.79 DPS) | yes | Mantle of Woe (7750, -0.69 DPS) [quest]; Fairywing Mantle (9536, -0.70 DPS) [quest]; Death Speaker Mantle (6685, -0.71 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.16 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Battle Healer's Cloak (19529, -0.23 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.22 DPS) [dungeon]; Beguiler Robes (7728, -0.30 DPS) [dungeon]; Green Silk Armor (7065, -0.71 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.09 DPS) | yes | Tabitha's Cuffs (251486, -1.39 DPS) [quest]; Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop]; Phoenix Bindings (210781, -1.47 DPS, sim-verified) [crafted] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.7 spell_power points (1.78 DPS) | yes | Serpent Gloves (5970, -0.16 DPS) [dungeon]; Truefaith Gloves (7049, -0.39 DPS) [crafted]; Gnoll Casting Gloves (892, -0.39 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (2.78 DPS) | yes | Warsong Sash (16975, -0.24 DPS) [quest]; Belt of Arugal (6392, -0.46 DPS) [dungeon]; Invoker's Cord (215366, -0.77 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.78 DPS) | yes | Pristine Leggings (253987, -0.61 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.92 DPS) [crafted]; Abomination Skin Leggings (23173, -1.05 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (2.17 DPS) | yes | Acidic Walkers (9454, -0.39 DPS) [dungeon]; Fiery Slippers (254005, -0.92 DPS) [crafted]; Spidersilk Boots (4320, -2.09 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.62 DPS) | yes | Electrocutioner Lagnut (9447, -0.93 DPS) [dungeon]; Snake Hoop (6750, -1.07 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.15 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.39 DPS) | yes | Snake Hoop (6750, -0.84 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.92 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.41 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | 7.9 spell_power points (1.83 DPS) | yes | Twisted Chanter's Staff (890, -1.05 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.05 DPS) [quest]; Glimmering Staff (249392, -1.63 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Unstable Power Core (279847, -1.98 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.02 DPS) [crafted]; Scorching Wand (5213, -4.00 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Staff of Soran'ruk; ranged: Starfaller

No-known-source sample (15 of 243, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 00000000000000000-0000000000000000000-2053225103101321)

Set DPS (verified): 138.0. Weights run: 1.8s. Verify run: 2.0s. 313 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.621 ± 0.029, crit=0.119 ± 0.006 per rating point (14 rating = 1%, 1.670 per %), hit=0.294 ± 0.028 per rating point (10 rating = 1%, 2.942 per %), spell_haste=not significant (0.329 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.282 ± 0.001, fire_power=0.718 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.58 DPS) | yes | Living Cowl (5608, -2.12 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -2.13 DPS) [dungeon]; Augural Shroud (2620, -2.14 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.0 spell_power points (3.18 DPS) | yes | Scorn's Icy Choker (23169, -0.33 DPS) [dungeon]; Necklace of Calisea (1714, -2.02 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.02 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.1 spell_power points (4.00 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.42 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.56 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.60 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (6.83 DPS) | yes | Dreamweave Vest (10021, -0.57 DPS) [crafted]; Robe of Power (7054, -1.13 DPS) [crafted]; Zealot's Robe (17043, -1.20 DPS) [quest] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.3 spell_power points (2.48 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.10 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (5.44 DPS) | yes | Black Mageweave Gloves (10003, -1.46 DPS) [crafted]; Fiery Handwraps (254025, -1.62 DPS) [crafted]; Red Mageweave Gloves (10018, -1.93 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.5 spell_power points (4.38 DPS) | yes | Fiery Cord (254041, -0.39 DPS) [crafted]; Star Belt (4329, -0.93 DPS) [crafted]; Deathmage Sash (10771, -1.49 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.5 spell_power points (5.70 DPS) | yes | Flame Leggings (253991, -1.87 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.87 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.99 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.37 DPS) | yes | Gilded Slippers (254001, -3.36 DPS) [crafted]; Acidic Walkers (9454, -3.73 DPS) [dungeon]; Fiery Slippers (254005, -3.82 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.48 DPS) | yes | Reedknot Ring (9622, -1.62 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.89 DPS) [vendor]; Voodoo Band (1996, -2.33 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.39 DPS) | yes | Reedknot Ring (9622, -0.53 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.80 DPS) [vendor]; Voodoo Band (1996, -1.24 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.33 DPS) [dungeon]; Searing Golden Blade (12260, -4.41 DPS) [crafted] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.72 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Noh'Orahil (15107, -0.22 DPS) [quest]; Orb of Mystic Insight (249394, -0.87 DPS) [crafted] |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (138.0 DPS) | yes | Twisted Nether Wand (249144, -0.08 DPS) [crafted]; Dancing Flame (6806, -0.66 DPS) [quest]; Jaina's Firestarter (13064, -2.40 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Umbral Wand

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 234.9. Weights run: 1.9s. Verify run: 2.2s. 392 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.202 ± 0.031, crit=0.168 ± 0.007 per rating point (14 rating = 1%, 2.346 per %), hit=0.447 ± 0.035 per rating point (10 rating = 1%, 4.472 per %), spell_haste=not significant (-0.308 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.280 ± 0.001, fire_power=0.720 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.71 DPS) | yes | Dreamweave Circlet (10041, -1.14 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.71 DPS) [crafted]; Red Mageweave Headband (10033, -2.12 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.6 spell_power points (2.46 DPS) | yes | Mindburst Medallion (11196, -0.40 DPS) [quest]; Scorn's Icy Choker (23169, -1.49 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.65 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.6 spell_power points (4.75 DPS) | yes | Netherflame Shoulders (254053, -0.79 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.28 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.38 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 spell_power points (4.34 DPS) | yes | Cindercloth Cloak (14044, -1.21 DPS) [crafted]; Runecloth Cloak (13860, -1.31 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.48 DPS, sim-verified) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 spell_power points (6.63 DPS) | yes | Elemental Raiment (9434, -0.63 DPS) [world_drop]; Dreamweave Vest (10021, -0.97 DPS) [crafted]; Acumen Robes (17775, -1.02 DPS, sim-verified) [quest] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.8 spell_power points (3.08 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Phoenix Bindings (210781, -1.79 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.37 DPS) | yes | Fiery Gloves (254099, -0.89 DPS) [crafted]; Black Mageweave Gloves (10003, -1.09 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.14 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.0 spell_power points (4.57 DPS) | yes | Fiery Waistcord (254085, -0.15 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.35 DPS) [rep]; Ghostweave Cord (254073, -0.58 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.43 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.64 DPS) [vendor]; Red Mageweave Pants (10009, -0.74 DPS) [crafted]; Cindercloth Leggings (12256, -1.52 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.85 DPS) | yes | Cindercloth Boots (10044, -2.54 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.77 DPS) [vendor]; Fiery Sandals (254111, -3.25 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.71 DPS) | yes | Philanthropist's Ring (281635, -0.57 DPS) [quest]; Reedknot Ring (9622, -1.71 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.00 DPS) [vendor] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.43 DPS) | yes | Philanthropist's Ring (281635, -0.28 DPS) [quest]; Reedknot Ring (9622, -1.43 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.71 DPS) [vendor] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.71 DPS) [world_drop]; Rune of the Guard Captain (19120, -2.53 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.26 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.12 DPS) [world]; Inventor's Focal Sword (17719, -0.69 DPS) [dungeon]; Hanzo Sword (8190, -27.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (234.9 DPS) | yes | Wand of Allistarj (13065, -3.01 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.39 DPS) [crafted]; Pyric Caduceus (11748, -8.67 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 460.4. Weights run: 1.9s. Verify run: 4.8s. 857 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.330 ± 0.036, crit=0.200 ± 0.008 per rating point (14 rating = 1%, 2.800 per %), hit=0.500 ± 0.039 per rating point (10 rating = 1%, 5.004 per %), spell_haste=not significant (0.991 ± 0.538), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.350 ± 0.001, fire_power=0.650 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | 30.9 spell_power points (13.16 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.50 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.03 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.36 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.15 DPS) [quest]; Beads of Ogre Mojo (22149, -2.15 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.7 spell_power points (14.79 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.76 DPS) [pvp]; Argent Shoulders (19059, -4.15 DPS) [crafted]; Mantle of the Timbermaw (19050, -4.54 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.6 spell_power points (10.06 DPS) | yes | Hide of the Wild (18510, -2.70 DPS) [crafted]; Spritecaster Cape (11623, -3.26 DPS) [dungeon]; Deep Woodlands Cloak (19121, -3.69 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.0 spell_power points (20.85 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.85 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.39 DPS) [pvp]; Acumen Robes (17775, -8.61 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.6 spell_power points (10.49 DPS) | yes | Runecloth Cuffs (254123, -4.40 DPS) [crafted]; Deathmist Bracers (226907, -5.82 DPS) [quest]; Netherflame Cuffs (254065, -5.91 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.6 spell_power points (12.20 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.80 DPS) [crafted]; Blood Guard's Dreadweave Handwraps (227099, -2.69 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.3 spell_power points (15.86 DPS) | yes | Belt of the Archmage (18405, -3.22 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -6.80 DPS) [rep]; Defiler's Cloth Girdle (20163, -7.87 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.9 spell_power points (17.41 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -3.67 DPS) [pvp]; Outrider's Silk Leggings (22747, -5.16 DPS, sim-verified) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.22 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.02 DPS) [dungeon]; Blood Guard's Dreadweave Walkers (227098, -0.73 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 36.0 spell_power points (15.31 DPS) | yes | Deep Rooted Ring (19109, -8.80 DPS) [world]; Blessed Band of Light (272407, -8.93 DPS) [vendor]; Advisor's Ring (19518, -9.35 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 25.6 spell_power points (10.91 DPS) | yes | Blessed Band of Light (272407, -4.52 DPS) [vendor]; Deep Rooted Ring (19109, -4.64 DPS, sim-verified) [world]; Advisor's Ring (19518, -4.95 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -0.85 DPS) [quest]; Weakness Analyzer (272438, -2.45 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.41 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.83 DPS) [world]; Teebu's Blazing Longsword (1728, -30.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (460.4 DPS) | yes | Wand of Biting Cold (19108, -0.39 DPS) [quest]; Stormrager (16997, -1.00 DPS) [quest]; Torch of Light (279246, -37.27 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 857, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 885.0. Weights run: 4.2s. Verify run: 5.2s. 857 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.060 ± 0.033, crit=0.256 ± 0.010 per rating point (14 rating = 1%, 3.588 per %), hit=0.738 ± 0.067 per rating point (10 rating = 1%, 7.376 per %), spell_haste=not significant (1.047 ± 0.857), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.279 ± 0.001, fire_power=0.721 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.10 DPS) | yes | Deathmist Mask (226909, +0.00 DPS, sim-verified) [quest]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.79 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.12 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.71 DPS) [quest]; Kezan's Taint (19604, -4.48 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.5 spell_power points (18.78 DPS) | yes | Warlord's Dreadweave Mantle (231592, -3.26 DPS) [pvp]; Mantle of the Timbermaw (19050, -6.03 DPS) [crafted]; Argent Shoulders (19059, -7.58 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.9 spell_power points (14.22 DPS) | yes | Hide of the Wild (18510, -5.52 DPS) [crafted]; Spritecaster Cape (11623, -5.66 DPS) [dungeon]; Mageflame Cloak (13007, -9.41 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (27.75 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.81 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.13 DPS) [pvp]; Robe of the Magi (1716, -15.26 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.5 spell_power points (13.40 DPS) | yes | Runecloth Cuffs (254123, -6.49 DPS) [crafted]; Netherflame Cuffs (254065, -7.56 DPS) [crafted]; Phoenix Bindings (210781, -7.81 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (16.28 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -2.96 DPS, sim-verified) [crafted]; Earth Warder's Gloves (21318, -3.16 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.7 spell_power points (20.09 DPS) | yes | Belt of the Archmage (18405, -7.86 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -9.00 DPS) [rep]; Defiler's Cloth Girdle (20163, -9.39 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 39.8 spell_power points (23.72 DPS) | yes | General's Dreadweave Pants (231588, -0.98 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -6.56 DPS) [pvp]; Outrider's Silk Leggings (22747, -9.96 DPS, sim-verified) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.31 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.19 DPS) [crafted]; Omnicast Boots (11822, -1.96 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.9 spell_power points (21.42 DPS) | yes | Elemental Focus Band (20682, -7.90 DPS) [world]; Advisor's Ring (19518, -13.07 DPS) [rep]; Band of the Unicorn (7553, -13.66 DPS) [world_drop] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+8.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Ring (19518, -2.01 DPS) [rep]; Band of the Unicorn (7553, -2.61 DPS) [world_drop]; Elemental Focus Band (20682, -8.91 DPS, sim-verified) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+19.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Abyss Shard (20534, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -1.19 DPS) [quest]; Weakness Analyzer (272438, -3.48 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.39 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.76 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -50.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (+64.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Biting Cold (19108, -1.20 DPS) [quest]; Stormrager (16997, -2.33 DPS) [quest]; Torch of Light (279246, -64.63 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Brilliant Wand

No-known-source sample (15 of 857, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

