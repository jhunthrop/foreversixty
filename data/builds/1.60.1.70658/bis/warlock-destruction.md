# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.7. Weights run: 2.0s. Verify run: 1.0s. 146 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.114 ± 0.003, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.389 per %), hit=0.116 ± 0.004 per rating point (10 rating = 1%, 1.156 per %), spell_haste=0.533 ± 0.050, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.002, fire_power=0.235 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.20 DPS) | yes | Shadow Goggles (4373, -1.08 DPS) [crafted]; Flame Circlet (253953, -2.39 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.0 spell_power points (1.20 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.19 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.41 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Caretaker's Cape (20428, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.24 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.33 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 spell_power points (1.11 DPS) | yes | Green Woolen Vest (2582, -0.31 DPS) [crafted]; Bloody Apron (6226, -0.31 DPS) [dungeon]; Gray Woolen Robe (2585, -1.39 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.09 DPS) [dungeon]; Bright Bracers (3647, -0.11 DPS) [world_drop]; Repurposed Hair Band (281256, -0.15 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.53 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.95 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 spell_power points (0.89 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Flame Sash (253929, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.96 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.9 spell_power points (1.98 DPS) | yes | Filigreed Pristine Leggings (253937, -0.65 DPS) [crafted]; Silk-threaded Trousers (1929, -0.66 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.98 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (1.49 DPS) | yes | Pristine Boots (253889, -0.82 DPS) [crafted]; Red Woolen Boots (4313, -0.83 DPS, sim-verified) [crafted]; Woolen Boots (2583, -1.04 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.2 spell_power points (1.04 DPS) | yes | Lavishly Jeweled Ring (1156, -0.91 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.98 DPS) [world_drop]; The 1 Ring (8350, -1.02 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.00 DPS) | yes | Volcanic Rock Ring (12053, -0.93 DPS) [world_drop]; The 1 Ring (8350, -0.98 DPS) [world]; Lavishly Jeweled Ring (1156, -1.28 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 1.1 spell_power points (0.23 DPS) | yes | Witching Stave (1484, -0.11 DPS) [dungeon]; Staff of Westfall (2042, -0.11 DPS) [quest]; Lesser Staff of the Spire (1300, -0.18 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.7 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.58 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.26 DPS) [quest]; Deepblaze (279896, -4.38 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 146, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 83.1. Weights run: 2.1s. Verify run: 1.1s. 264 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.328 ± 0.008, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.988 per %), hit=0.166 ± 0.007 per rating point (10 rating = 1%, 1.664 per %), spell_haste=not significant (0.095 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.764 ± 0.002, fire_power=0.236 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Silk Headband (7050, -0.46 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.69 DPS) [crafted]; Enchanter's Cowl (4322, -3.61 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.05 DPS) | yes | Darkspear Warding Pendant (272075, -1.67 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.75 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.81 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (2.73 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.69 DPS) [quest]; Invoker's Mantle (215365, -0.76 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.14 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Prelacy Cape (7004, -0.23 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.01 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.24 DPS) [dungeon]; Green Silk Armor (7065, -1.03 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.06 DPS) | yes | Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop]; Nightsky Wristbands (6407, -1.61 DPS) [world_drop]; Phoenix Bindings (210781, -1.64 DPS, sim-verified) [crafted] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.60 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.23 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.0 spell_power points (2.74 DPS) | yes | Belt of Arugal (6392, -0.46 DPS) [dungeon]; Invoker's Cord (215366, -0.76 DPS) [crafted]; Crimson Silk Belt (7055, -0.84 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.74 DPS) | yes | Abomination Skin Leggings (23173, -0.51 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.62 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.92 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.3 spell_power points (2.12 DPS) | yes | Acidic Walkers (9454, -0.38 DPS) [dungeon]; Glinteye Slippers (273024, -0.52 DPS) [dungeon]; Spidersilk Boots (4320, -1.44 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.6 spell_power points (1.74 DPS) | yes | Ring of Power Regulation (273030, -0.37 DPS) [dungeon]; Minor Channeling Ring (1449, -0.45 DPS) [quest]; Lorekeeper's Ring (19525, -1.42 DPS, sim-verified) [rep] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Minor Channeling Ring (1449, -0.08 DPS) [quest]; Lorekeeper's Ring (19525, -1.46 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.22 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.65 DPS) [quest]; Staff of Soran'ruk (15109, -6.39 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 145.2 spell_power points (33.16 DPS) | yes | Wand of Mana Concentration (273043, -1.79 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.02 DPS) [crafted]; Scorching Wand (5213, -3.99 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Sea Giant's Toe Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 134.1. Weights run: 1.7s. Verify run: 1.1s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.612 ± 0.019, crit=0.120 ± 0.004 per rating point (14 rating = 1%, 1.675 per %), hit=0.291 ± 0.018 per rating point (10 rating = 1%, 2.908 per %), spell_haste=not significant (0.522 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.283 ± 0.000, fire_power=0.717 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.44 DPS) | yes | Papal Fez (9431, -0.67 DPS) [dungeon]; Augural Shroud (2620, -1.01 DPS) [world]; Whitemane's Chapeau (7720, -2.04 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.9 spell_power points (3.08 DPS) | yes | Scorn's Icy Choker (23169, -0.74 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -1.01 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.97 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.0 spell_power points (3.87 DPS) | yes | Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest]; Green Silken Shoulders (7057, -0.65 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.35 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.53 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.60 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (6.65 DPS) | yes | Dreamweave Vest (10021, -0.56 DPS) [crafted]; Robe of Power (7054, -1.12 DPS) [crafted]; Elemental Raiment (9434, -1.21 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.3 spell_power points (2.41 DPS) | yes | Arcane Runed Bracers (4744, -0.08 DPS) [quest]; Spidertank Oilrag (9448, -0.08 DPS) [dungeon]; Condor Bracers (15864, -0.60 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (5.30 DPS) | yes | Black Mageweave Gloves (10003, -1.41 DPS) [crafted]; Fiery Handwraps (254025, -1.59 DPS) [crafted]; Red Mageweave Gloves (10018, -2.30 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.4 spell_power points (4.26 DPS) | yes | Fiery Cord (254041, -0.39 DPS) [crafted]; Star Belt (4329, -0.89 DPS) [crafted]; Deathmage Sash (10771, -1.16 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.3 spell_power points (5.53 DPS) | yes | Flame Leggings (253991, -1.82 DPS) [crafted]; Abomination Skin Leggings (23173, -1.93 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.98 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.22 DPS) | yes | Gilded Slippers (254001, -3.29 DPS) [crafted]; Acidic Walkers (9454, -3.65 DPS) [dungeon]; Fiery Slippers (254005, -3.73 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.38 DPS) | yes | Lorekeeper's Ring (19524, -1.05 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.31 DPS) [quest]; Reedknot Ring (9622, -1.57 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.1 spell_power points (2.35 DPS) | yes | Lorekeeper's Ring (19524, +0.00 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.53 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (134.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.32 DPS) [dungeon]; Spine of the Devourer (273040, -3.34 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.63 DPS) | yes | Orb of Noh'Orahil (15107, -0.24 DPS) [quest]; Orb of Lorica (11262, -0.49 DPS) [quest]; Orb of Mystic Insight (249394, -0.86 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 155.7 spell_power points (40.33 DPS) | yes | Earthen Rod (9381, -3.29 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.66 DPS) [dungeon]; Twisted Nether Wand (249144, -4.78 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 234.3. Weights run: 1.6s. Verify run: 1.4s. 429 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.223 ± 0.023, crit=0.172 ± 0.006 per rating point (14 rating = 1%, 2.409 per %), hit=0.448 ± 0.026 per rating point (10 rating = 1%, 4.484 per %), spell_haste=not significant (-0.047 ± 0.418), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.279 ± 0.000, fire_power=0.721 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.66 DPS) | yes | Red Mageweave Headband (10033, -1.03 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.70 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.8 spell_power points (2.49 DPS) | yes | Triune Amulet (7722, -0.22 DPS) [dungeon]; Mindburst Medallion (11196, -0.41 DPS) [quest]; Scorn's Icy Choker (23169, -1.93 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.0 spell_power points (4.83 DPS) | yes | Netherflame Shoulders (254053, -0.84 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.31 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.42 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.35 DPS) | yes | Runecloth Cloak (13860, -1.29 DPS) [crafted]; Cloak of Fire (14134, -1.49 DPS) [crafted]; Cindercloth Cloak (14044, -1.57 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.5 spell_power points (6.66 DPS) | yes | Robe of the Magi (1716, -0.04 DPS) [world_drop]; Elemental Raiment (9434, -0.70 DPS) [world_drop]; Dreamweave Vest (10021, -0.98 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.9 spell_power points (3.10 DPS) | yes | Arcane Runed Bracers (4744, -0.55 DPS) [quest]; Spidertank Oilrag (9448, -0.55 DPS) [dungeon]; Phoenix Bindings (210781, -1.70 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (5.36 DPS) | yes | Fiery Gloves (254099, -0.84 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.10 DPS) [vendor]; Black Mageweave Gloves (10003, -1.10 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.2 spell_power points (4.60 DPS) | yes | Fiery Waistcord (254085, -0.15 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.38 DPS) [rep]; Ghostweave Cord (254073, -0.63 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.39 DPS) | yes | Red Mageweave Pants (10009, -0.66 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -0.97 DPS, sim-verified) [vendor]; Cindercloth Leggings (12256, -1.50 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.81 DPS) | yes | Cindercloth Boots (10044, -2.51 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.70 DPS) [vendor]; Fiery Sandals (254111, -3.25 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.69 DPS) | yes | Philanthropist's Ring (281635, -0.53 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.42 DPS) [quest]; Snake Hoop (6750, -1.67 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.40 DPS) | yes | Philanthropist's Ring (281635, -0.89 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.13 DPS) [quest]; Snake Hoop (6750, -1.39 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.70 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.19 DPS) [world]; Inventor's Focal Sword (17719, -0.76 DPS) [dungeon]; Hanzo Sword (8190, -27.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (234.3 DPS) | yes | Wand of Allistarj (13065, -3.00 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.40 DPS) [crafted]; Pyric Caduceus (11748, -8.80 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 429, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 467.8. Weights run: 1.8s. Verify run: 2.6s. 882 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.324 ± 0.026, crit=0.202 ± 0.006 per rating point (14 rating = 1%, 2.835 per %), hit=0.527 ± 0.028 per rating point (10 rating = 1%, 5.266 per %), spell_haste=not significant (0.992 ± 0.383), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.344 ± 0.001, fire_power=0.656 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | 31.0 spell_power points (13.11 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.58 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.65 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.30 DPS) | yes | Amulet of the Dawn (22657, -1.18 DPS) [quest]; Orb of the Darkmoon (19426, -2.11 DPS, sim-verified) [quest]; Beads of Ogre Mojo (22149, -2.16 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.7 spell_power points (14.66 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.77 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.50 DPS) [crafted]; Argent Shoulders (19059, -4.75 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.9 spell_power points (10.08 DPS) | yes | Hide of the Wild (18510, -1.97 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -3.34 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -4.16 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.9 spell_power points (20.67 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.86 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.37 DPS) [pvp]; Acumen Robes (17775, -9.94 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.6 spell_power points (10.39 DPS) | yes | Runecloth Cuffs (254123, -4.37 DPS) [crafted]; Deathmist Bracers (226907, -5.79 DPS) [quest]; Netherflame Cuffs (254065, -5.83 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.6 spell_power points (12.09 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -2.60 DPS) [quest]; Inferno Gloves (18408, -4.41 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.4 spell_power points (15.80 DPS) | yes | Belt of the Archmage (18405, -6.19 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -6.82 DPS) [rep]; Highlander's Cloth Girdle (20047, -7.86 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.9 spell_power points (17.28 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -3.67 DPS) [pvp]; Deathmist Leggings (226910, -6.26 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.14 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.06 DPS) [pvp]; Omnicast Boots (11822, -1.89 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 36.2 spell_power points (15.29 DPS) | yes | Maiden's Circle (13001, -6.72 DPS) [world_drop]; Blessed Band of Light (272407, -8.84 DPS) [vendor]; Deep Rooted Ring (19109, -8.85 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 25.6 spell_power points (10.80 DPS) | yes | Maiden's Circle (13001, -2.62 DPS, sim-verified) [world_drop]; Blessed Band of Light (272407, -4.35 DPS) [vendor]; Deep Rooted Ring (19109, -4.36 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+12.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -0.85 DPS) [quest]; Weakness Analyzer (272438, -2.49 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.38 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.80 DPS) [world]; Teebu's Blazing Longsword (1728, -36.01 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (467.8 DPS) | yes | Wand of Biting Cold (19108, -0.34 DPS) [quest]; Stormrager (16997, -0.97 DPS) [quest]; Torch of Light (279246, -40.16 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 882, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 879.5. Weights run: 6.1s. Verify run: 2.7s. 882 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.056 ± 0.024, crit=0.272 ± 0.008 per rating point (14 rating = 1%, 3.810 per %), hit=0.716 ± 0.048 per rating point (10 rating = 1%, 7.164 per %), spell_haste=not significant (1.208 ± 0.600), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.278 ± 0.001, fire_power=0.722 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.08 DPS) | yes | Deathmist Mask (226909, +0.00 DPS, sim-verified) [quest]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.71 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.11 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.74 DPS) [quest]; Kezan's Taint (19604, -4.50 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.6 spell_power points (18.85 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -3.40 DPS) [pvp]; Argent Shoulders (19059, -5.21 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -6.02 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.6 spell_power points (14.06 DPS) | yes | Hide of the Wild (18510, -5.39 DPS) [crafted]; Spritecaster Cape (11623, -5.53 DPS) [dungeon]; Mageflame Cloak (13007, -9.04 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (27.70 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.84 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.15 DPS) [pvp]; Robe of the Magi (1716, -13.06 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.4 spell_power points (13.37 DPS) | yes | Runecloth Cuffs (254123, -6.49 DPS) [crafted]; Netherflame Cuffs (254065, -7.55 DPS) [crafted]; Phoenix Bindings (210781, -7.78 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (16.25 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.76 DPS) [crafted]; Earth Warder's Gloves (21318, -3.14 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.4 spell_power points (19.89 DPS) | yes | Belt of the Archmage (18405, -8.58 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -8.84 DPS) [rep]; Highlander's Cloth Girdle (20047, -9.08 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.2 spell_power points (23.93 DPS) | yes | Marshal's Dreadweave Leggings (231587, -1.26 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -6.82 DPS) [pvp]; Ghostweave Pants (14144, -8.44 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.30 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.19 DPS) [crafted]; Omnicast Boots (11822, -1.99 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.7 spell_power points (21.25 DPS) | yes | Maiden's Circle (13001, -10.29 DPS) [world_drop]; Blessed Band of Light (272407, -11.02 DPS) [vendor]; Lorekeeper's Ring (19522, -12.91 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.6 spell_power points (13.47 DPS) | yes | Maiden's Circle (13001, -2.52 DPS) [world_drop]; Blessed Band of Light (272407, -3.24 DPS) [vendor]; Lorekeeper's Ring (19522, -5.13 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+19.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Abyss Shard (20534, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -1.19 DPS) [quest]; Weakness Analyzer (272438, -3.51 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.50 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.79 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -48.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (879.5 DPS) | yes | Wand of Biting Cold (19108, -1.17 DPS) [quest]; Stormrager (16997, -2.32 DPS) [quest]; Torch of Light (279246, -58.14 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Brilliant Wand

No-known-source sample (15 of 882, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 38.6. Weights run: 2.0s. Verify run: 1.0s. 135 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.114 ± 0.003, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.389 per %), hit=0.116 ± 0.004 per rating point (10 rating = 1%, 1.156 per %), spell_haste=0.533 ± 0.050, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.002, fire_power=0.235 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.20 DPS) | yes | Shadow Goggles (4373, -1.08 DPS) [crafted]; Flame Circlet (253953, -2.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.0 spell_power points (1.20 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.11 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.41 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Battle Healer's Cloak (20427, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.32 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.33 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 spell_power points (1.11 DPS) | yes | Green Woolen Vest (2582, -0.31 DPS) [crafted]; Bloody Apron (6226, -0.31 DPS) [dungeon]; Gray Woolen Robe (2585, -1.35 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.60 DPS) | yes | Windsong Bangles (263336, -0.40 DPS) [quest]; Featherbead Bracers (15452, -0.49 DPS) [quest]; Owlbeard Bracers (16981, -0.59 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.53 DPS) [crafted]; Apothecary Gloves (10919, -0.60 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 spell_power points (0.89 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Flame Sash (253929, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.90 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.9 spell_power points (1.98 DPS) | yes | Silk-threaded Trousers (1929, -0.61 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.65 DPS) [crafted]; Rumpled Kilt (274741, -0.98 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (1.49 DPS) | yes | Red Woolen Boots (4313, -0.62 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.82 DPS) [crafted]; Woolen Boots (2583, -1.04 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.00 DPS) | yes | Loop of Sacrifice (281673, -0.88 DPS) [quest]; Volcanic Rock Ring (12053, -0.93 DPS) [world_drop]; The 1 Ring (8350, -0.98 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.7 spell_power points (0.14 DPS) | yes | Loop of Sacrifice (281673, -0.02 DPS) [quest]; Volcanic Rock Ring (12053, -0.07 DPS) [world_drop]; The 1 Ring (8350, -0.11 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 1.1 spell_power points (0.23 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.09 DPS) [world]; Clear Crystal Rod (16894, -0.11 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.7 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.51 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.26 DPS) [quest]; Greater Magic Wand (11288, -5.01 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 135, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest

### Band 30 (troll, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 81.2. Weights run: 2.1s. Verify run: 1.1s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.328 ± 0.008, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.988 per %), hit=0.166 ± 0.007 per rating point (10 rating = 1%, 1.664 per %), spell_haste=not significant (0.095 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.764 ± 0.002, fire_power=0.236 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Silk Headband (7050, -0.46 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.69 DPS) [crafted]; Enchanter's Cowl (4322, -3.64 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.05 DPS) | yes | Darkspear Warding Pendant (272075, -1.67 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.75 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.78 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (2.73 DPS) | yes | Death Speaker Mantle (6685, -0.51 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.69 DPS) [quest]; Invoker's Mantle (215365, -0.76 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.14 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Battle Healer's Cloak (19529, -0.23 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.3 spell_power points (3.03 DPS) | yes | Tree Bark Jacket (1486, -0.06 DPS) [dungeon]; Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.30 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.06 DPS) | yes | Tabitha's Cuffs (251486, -1.37 DPS) [quest]; Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop]; Phoenix Bindings (210781, -2.05 DPS, sim-verified) [crafted] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.6 spell_power points (1.74 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.37 DPS) [world]; Truefaith Gloves (7049, -0.38 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (2.74 DPS) | yes | Warsong Sash (16975, -0.22 DPS) [quest]; Belt of Arugal (6392, -0.46 DPS) [dungeon]; Invoker's Cord (215366, -0.76 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.74 DPS) | yes | Pristine Leggings (253987, -0.62 DPS) [crafted]; Abomination Skin Leggings (23173, -0.72 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.92 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.3 spell_power points (2.12 DPS) | yes | Acidic Walkers (9454, -0.38 DPS) [dungeon]; Glinteye Slippers (273024, -0.52 DPS) [dungeon]; Spidersilk Boots (4320, -1.61 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.6 spell_power points (1.74 DPS) | yes | Ring of Power Regulation (273030, -0.37 DPS) [dungeon]; Advisor's Ring (20426, -0.60 DPS) [rep] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (81.2 DPS) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Advisor's Ring (20426, -0.23 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.22 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -6.39 DPS) [quest]; Glimmering Staff (249392, -7.40 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 145.2 spell_power points (33.16 DPS) | yes | Wand of Mana Concentration (273043, -2.45 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.02 DPS) [crafted]; Unstable Power Core (279847, -3.45 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Sea Giant's Toe Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 133.6. Weights run: 1.7s. Verify run: 1.1s. 321 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.612 ± 0.019, crit=0.120 ± 0.004 per rating point (14 rating = 1%, 1.675 per %), hit=0.291 ± 0.018 per rating point (10 rating = 1%, 2.908 per %), spell_haste=not significant (0.522 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.283 ± 0.000, fire_power=0.717 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.44 DPS) | yes | Augural Shroud (2620, -1.01 DPS) [world]; Papal Fez (9431, -1.05 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -2.04 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.9 spell_power points (3.08 DPS) | yes | Triune Amulet (7722, -1.01 DPS) [dungeon]; Scorn's Icy Choker (23169, -1.41 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.97 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.0 spell_power points (3.87 DPS) | yes | Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest]; Green Silken Shoulders (7057, -1.32 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.35 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.53 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.60 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (6.65 DPS) | yes | Dreamweave Vest (10021, -0.95 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.12 DPS) [crafted]; Elemental Raiment (9434, -1.21 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.3 spell_power points (2.41 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.11 DPS) [quest]; Spidertank Oilrag (9448, -0.78 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (5.30 DPS) | yes | Black Mageweave Gloves (10003, -1.41 DPS) [crafted]; Fiery Handwraps (254025, -1.59 DPS) [crafted]; Red Mageweave Gloves (10018, -2.81 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.4 spell_power points (4.26 DPS) | yes | Fiery Cord (254041, -0.39 DPS) [crafted]; Star Belt (4329, -0.89 DPS) [crafted]; Deathmage Sash (10771, -1.95 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.3 spell_power points (5.53 DPS) | yes | Flame Leggings (253991, -1.82 DPS) [crafted]; Abomination Skin Leggings (23173, -1.93 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.05 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.22 DPS) | yes | Gilded Slippers (254001, -3.29 DPS) [crafted]; Acidic Walkers (9454, -3.65 DPS) [dungeon]; Fiery Slippers (254005, -4.52 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.38 DPS) | yes | Advisor's Ring (19520, -1.05 DPS) [rep]; Reedknot Ring (9622, -1.57 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.1 spell_power points (2.35 DPS) | yes | Advisor's Ring (19520, -0.02 DPS) [rep]; Reedknot Ring (9622, -0.53 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (133.6 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.32 DPS) [dungeon]; Spine of the Devourer (273040, -3.34 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.63 DPS) | yes | Orb of Mystic Insight (249394, -0.86 DPS) [crafted]; Rod of Molten Fire (2565, -1.21 DPS) [world_drop]; Orb of Noh'Orahil (15107, -1.56 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 155.7 spell_power points (40.33 DPS) | yes | Earthen Rod (9381, -4.36 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.66 DPS) [dungeon]; Twisted Nether Wand (249144, -4.78 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 321, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 234.8. Weights run: 1.6s. Verify run: 1.3s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.223 ± 0.023, crit=0.172 ± 0.006 per rating point (14 rating = 1%, 2.409 per %), hit=0.448 ± 0.026 per rating point (10 rating = 1%, 4.484 per %), spell_haste=not significant (-0.047 ± 0.418), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.279 ± 0.000, fire_power=0.721 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.66 DPS) | yes | Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Red Mageweave Headband (10033, -1.17 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.70 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.8 spell_power points (2.49 DPS) | yes | Triune Amulet (7722, -0.22 DPS) [dungeon]; Mindburst Medallion (11196, -0.41 DPS) [quest]; Scorn's Icy Choker (23169, -1.40 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.0 spell_power points (4.83 DPS) | yes | Netherflame Shoulders (254053, -0.91 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.31 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.42 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.35 DPS) | yes | Deep Woodlands Cloak (19121, -1.11 DPS, sim-verified) [quest]; Cindercloth Cloak (14044, -1.19 DPS) [crafted]; Runecloth Cloak (13860, -1.29 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.5 spell_power points (6.66 DPS) | yes | Robe of the Magi (1716, -0.04 DPS) [world_drop]; Elemental Raiment (9434, -0.70 DPS) [world_drop]; Dreamweave Vest (10021, -0.98 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.9 spell_power points (3.10 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Phoenix Bindings (210781, -0.96 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (5.36 DPS) | yes | Fiery Gloves (254099, -1.02 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.10 DPS) [vendor]; Black Mageweave Gloves (10003, -1.10 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.2 spell_power points (4.60 DPS) | yes | Fiery Waistcord (254085, -0.15 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.38 DPS) [rep]; Ghostweave Cord (254073, -0.63 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.39 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.54 DPS) [vendor]; Red Mageweave Pants (10009, -0.66 DPS) [crafted]; Cindercloth Leggings (12256, -1.50 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.81 DPS) | yes | Cindercloth Boots (10044, -2.51 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.70 DPS) [vendor]; Fiery Sandals (254111, -3.30 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.69 DPS) | yes | Philanthropist's Ring (281635, -0.53 DPS) [quest]; Snake Hoop (6750, -1.67 DPS) [quest]; Reedknot Ring (9622, -1.70 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.40 DPS) | yes | Philanthropist's Ring (281635, -0.25 DPS) [quest]; Snake Hoop (6750, -1.39 DPS) [quest]; Reedknot Ring (9622, -1.42 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.70 DPS) [world_drop]; Rune of the Guard Captain (19120, -2.51 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.25 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.19 DPS) [world]; Inventor's Focal Sword (17719, -0.76 DPS) [dungeon]; Hanzo Sword (8190, -26.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (234.8 DPS) | yes | Wand of Allistarj (13065, -3.00 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.40 DPS) [crafted]; Pyric Caduceus (11748, -10.14 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 458.7. Weights run: 1.8s. Verify run: 2.5s. 871 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.324 ± 0.026, crit=0.202 ± 0.006 per rating point (14 rating = 1%, 2.835 per %), hit=0.527 ± 0.028 per rating point (10 rating = 1%, 5.266 per %), spell_haste=not significant (0.992 ± 0.383), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.344 ± 0.001, fire_power=0.656 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | 31.0 spell_power points (13.11 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.58 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.05 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.30 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.18 DPS) [quest]; Beads of Ogre Mojo (22149, -2.16 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.7 spell_power points (14.66 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.77 DPS) [pvp]; Argent Shoulders (19059, -4.10 DPS) [crafted]; Mantle of the Timbermaw (19050, -4.50 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.9 spell_power points (10.08 DPS) | yes | Hide of the Wild (18510, -2.80 DPS) [crafted]; Spritecaster Cape (11623, -3.34 DPS) [dungeon]; Deep Woodlands Cloak (19121, -3.78 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.9 spell_power points (20.67 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.86 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.37 DPS) [pvp]; Acumen Robes (17775, -8.65 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.6 spell_power points (10.39 DPS) | yes | Runecloth Cuffs (254123, -4.37 DPS) [crafted]; Deathmist Bracers (226907, -5.79 DPS) [quest]; Netherflame Cuffs (254065, -5.83 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.6 spell_power points (12.09 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.72 DPS) [crafted]; Deathmist Wraps (226911, -2.60 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.4 spell_power points (15.80 DPS) | yes | Belt of the Archmage (18405, -3.22 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -6.82 DPS) [rep]; Defiler's Cloth Girdle (20163, -7.86 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.9 spell_power points (17.28 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -3.67 DPS) [pvp]; Outrider's Silk Leggings (22747, -5.22 DPS, sim-verified) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.14 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.05 DPS) [dungeon]; Blood Guard's Dreadweave Walkers (227098, -0.76 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 36.2 spell_power points (15.29 DPS) | yes | Maiden's Circle (13001, -6.72 DPS) [world_drop]; Blessed Band of Light (272407, -8.84 DPS) [vendor]; Deep Rooted Ring (19109, -8.85 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 25.6 spell_power points (10.80 DPS) | yes | Maiden's Circle (13001, -2.58 DPS, sim-verified) [world_drop]; Blessed Band of Light (272407, -4.35 DPS) [vendor]; Deep Rooted Ring (19109, -4.36 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -0.85 DPS) [quest]; Weakness Analyzer (272438, -2.43 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.38 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.80 DPS) [world]; Teebu's Blazing Longsword (1728, -30.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (458.7 DPS) | yes | Wand of Biting Cold (19108, -0.34 DPS) [quest]; Stormrager (16997, -0.97 DPS) [quest]; Torch of Light (279246, -37.37 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 871, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 881.0. Weights run: 6.1s. Verify run: 2.9s. 871 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.056 ± 0.024, crit=0.272 ± 0.008 per rating point (14 rating = 1%, 3.810 per %), hit=0.716 ± 0.048 per rating point (10 rating = 1%, 7.164 per %), spell_haste=not significant (1.208 ± 0.600), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.278 ± 0.001, fire_power=0.722 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.08 DPS) | yes | Deathmist Mask (226909, +0.00 DPS, sim-verified) [quest]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.71 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.11 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.74 DPS) [quest]; Kezan's Taint (19604, -4.50 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.6 spell_power points (18.85 DPS) | yes | Warlord's Dreadweave Mantle (231592, -3.40 DPS) [pvp]; Mantle of the Timbermaw (19050, -6.02 DPS) [crafted]; Argent Shoulders (19059, -7.58 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.6 spell_power points (14.06 DPS) | yes | Hide of the Wild (18510, -5.39 DPS) [crafted]; Spritecaster Cape (11623, -5.53 DPS) [dungeon]; Mageflame Cloak (13007, -9.41 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (27.70 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.84 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.15 DPS) [pvp]; Robe of the Magi (1716, -15.26 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.4 spell_power points (13.37 DPS) | yes | Runecloth Cuffs (254123, -6.49 DPS) [crafted]; Netherflame Cuffs (254065, -7.55 DPS) [crafted]; Phoenix Bindings (210781, -7.78 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (16.25 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -2.96 DPS, sim-verified) [crafted]; Earth Warder's Gloves (21318, -3.14 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.4 spell_power points (19.89 DPS) | yes | Belt of the Archmage (18405, -7.86 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -8.84 DPS) [rep]; Defiler's Cloth Girdle (20163, -9.08 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.2 spell_power points (23.93 DPS) | yes | General's Dreadweave Pants (231588, -1.26 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -6.82 DPS) [pvp]; Outrider's Silk Leggings (22747, -9.96 DPS, sim-verified) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.30 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.19 DPS) [crafted]; Omnicast Boots (11822, -1.99 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.7 spell_power points (21.25 DPS) | yes | Maiden's Circle (13001, -10.29 DPS) [world_drop]; Blessed Band of Light (272407, -11.02 DPS) [vendor]; Advisor's Ring (19518, -12.91 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.6 spell_power points (13.47 DPS) | yes | Blessed Band of Light (272407, -3.24 DPS) [vendor]; Maiden's Circle (13001, -3.58 DPS, sim-verified) [world_drop]; Advisor's Ring (19518, -5.13 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+19.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Abyss Shard (20534, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -1.19 DPS) [quest]; Weakness Analyzer (272438, -3.48 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.50 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.79 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -50.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (881.0 DPS) | yes | Wand of Biting Cold (19108, -1.17 DPS) [quest]; Stormrager (16997, -2.32 DPS) [quest]; Torch of Light (279246, -64.63 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Brilliant Wand

No-known-source sample (15 of 871, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

