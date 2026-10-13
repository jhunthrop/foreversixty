# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 41.9. Weights run: 1.4s. Verify run: 1.5s. 192 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.101 ± 0.005, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.476 per %), hit=0.118 ± 0.006 per rating point (10 rating = 1%, 1.176 per %), spell_haste=not significant (-0.216 ± 0.075), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.747 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.14 DPS) | yes | Red Winter Hat (21524, -2.97 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.19 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.76 DPS) | yes | Caretaker's Cape (20428, -0.19 DPS) [rep]; Pearl-clasped Cloak (5542, -0.38 DPS) [crafted]; Black Whelp Cloak (7283, -0.39 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -1.44 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.19 DPS) | yes | Ivycloth Bracelets (9793, -0.25 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.33 DPS) | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.57 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.95 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (41.9 DPS) | yes | Novice Ardent's Sash (253887, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.87 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.71 DPS) | yes | Filigreed Pristine Leggings (253937, -0.57 DPS) [crafted]; Silk-threaded Trousers (1929, -0.66 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.76 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.33 DPS) | yes | Pristine Boots (253889, -0.76 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS, sim-verified) [crafted]; Soft-soled Linen Boots (4312, -0.95 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.95 DPS) | yes | - |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.95 DPS) | yes | Ring of the Shadow (1462, -1.31 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Witching Stave (1484), Lesser Staff of the Spire (1300), Riverpaw Mystic Staff (1391), and 75 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.36 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.95 DPS) | yes | Bouquet of Red Roses (22206, -1.30 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 120.2 spell_power points (22.88 DPS) | yes | Skycaller (12984, -2.40 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.28 DPS) [quest]; Deepblaze (279896, -4.35 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1917 Jeweled Dagger; 2235 Brackclaw; 2241 Desperado Cape; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash; 3563 Seafarer's Pantaloons; 3738 Brewing Rod; 3902 Staff of Nobles

### Band 30 (gnome, 25552200000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 78.9. Weights run: 1.4s. Verify run: 1.3s. 319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.146 ± 0.007, crit=0.036 ± 0.002 per rating point (14 rating = 1%, 0.507 per %), hit=0.131 ± 0.007 per rating point (10 rating = 1%, 1.310 per %), spell_haste=not significant (-0.013 ± 0.106), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.775 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Filigreed Pristine Circlet (253975, -0.68 DPS) [crafted]; Enchanter's Cowl (4322, -0.81 DPS) [crafted]; Silk Headband (7050, -3.58 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (1.80 DPS) | yes | Crystal Starfire Medallion (5003, -1.66 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.66 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.86 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (2.35 DPS) | yes | Invoker's Mantle (215365, -0.47 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.62 DPS) [dungeon]; Fairywing Mantle (9536, -0.68 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (1.37 DPS) | yes | Hillman's Cloak (3719, -0.26 DPS, sim-verified) [crafted]; Heavy Woolen Cloak (4311, -0.46 DPS) [crafted]; Prelacy Cape (7004, -0.46 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.97 DPS) | yes | Green Silk Armor (7065, -0.44 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.65 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.74 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.05 DPS) | yes | Windsong Bangles (263336, -1.83 DPS) [quest]; Nightsky Wristbands (6407, -1.85 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.08 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.60 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Town Clerk's Mittens (270029, -0.32 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.4 spell_power points (2.61 DPS) | yes | Belt of Arugal (6392, -0.46 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.78 DPS) [dungeon]; Invoker's Cord (215366, -0.85 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.74 DPS) | yes | Abomination Skin Leggings (23173, -0.66 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.91 DPS) [crafted]; Silk-threaded Trousers (1929, -1.14 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (1.83 DPS) | yes | Acidic Walkers (9454, -0.42 DPS) [dungeon]; Nimbus Boots (6998, -0.46 DPS) [quest]; Spidersilk Boots (4320, -1.80 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.60 DPS) | yes | Minor Channeling Ring (1449, -0.39 DPS) [quest]; Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Snake Hoop (6750, -1.36 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.37 DPS) | yes | Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon]; Snake Hoop (6750, -1.14 DPS) [quest]; Minor Channeling Ring (1449, -1.52 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | 30.0 spell_power points (6.85 DPS) | yes | Twisted Chanter's Staff (890, -6.51 DPS) [world_drop]; Lesser Staff of the Spire (1300, -6.65 DPS) [world]; Glimmering Staff (249392, -10.38 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 144.5 spell_power points (32.99 DPS) | yes | Greater Mystic Wand (217287, -3.55 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -3.99 DPS) [quest]; Scorching Wand (5213, -4.14 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff; ranged: Starfaller

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash; 3563 Seafarer's Pantaloons

### Band 40 (gnome, 25552300120201201-0000000000000000000-0000000000000000)

Set DPS (verified): 173.2. Weights run: 1.5s. Verify run: 1.6s. 402 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.709 ± 0.020, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.369 per %), hit=0.198 ± 0.015 per rating point (10 rating = 1%, 1.977 per %), spell_haste=not significant (0.495 ± 0.182), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.875 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.19 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -2.51 DPS) [dungeon]; Enchanter's Cowl (4322, -2.71 DPS) [crafted]; Augural Shroud (2620, -3.31 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (4.34 DPS) | yes | Scorn's Icy Choker (23169, -1.57 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.64 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.64 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (5.55 DPS) | yes | Bloodmage Mantle (7684, -0.29 DPS) [dungeon]; Berylline Pads (4197, -0.73 DPS) [quest]; Green Silken Shoulders (7057, -1.45 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (3.27 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.60 DPS) [vendor]; Icy Cloak (4327, -0.87 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.3 spell_power points (8.99 DPS) | yes | Dreamweave Vest (10021, -0.66 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.28 DPS) [crafted]; Elemental Raiment (9434, -1.80 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.08 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.69 DPS) [quest]; Windchaser Cuffs (14429, -0.90 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.14 DPS) | yes | Black Mageweave Gloves (10003, -2.00 DPS) [crafted]; Red Mageweave Gloves (10018, -2.01 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.70 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -1.08 DPS) [crafted]; Star Belt (4329, -1.31 DPS) [crafted]; Deathmage Sash (10771, -1.99 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (7.71 DPS) | yes | Crimson Silk Pantaloons (7062, -1.98 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.68 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.37 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.22 DPS) | yes | Gilded Slippers (254001, -3.17 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.57 DPS) [dungeon]; Spidersilk Boots (4320, -4.85 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.64 DPS) | yes | Ring of Forlorn Spirits (2043, -1.90 DPS) [quest]; Reedknot Ring (9622, -2.24 DPS) [quest]; Minor Channeling Ring (1449, -2.44 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (3.08 DPS) | yes | Ring of Forlorn Spirits (2043, -0.49 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.69 DPS) [quest]; Minor Channeling Ring (1449, -0.88 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.49 DPS) [dungeon]; Hardened Root Staff (1317, -6.74 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.80 DPS) | yes | Thrash's Trash (276204, +0.00 DPS, sim-verified) [vendor]; Orb of Lorica (11262, -0.31 DPS) [quest]; Orb of Mystic Insight (249394, -0.94 DPS) [crafted] |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Umbral Wand (5216, -0.39 DPS) [dungeon]; Burning Sliver (5249, -1.31 DPS) [quest]; Jaina's Firestarter (13064, -4.25 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Twisted Nether Wand

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash

### Band 50 (gnome, 25552300120201351-2002000000000000000-0000000000000000)

Set DPS (verified): 243.6. Weights run: 2.9s. Verify run: 1.9s. 488 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.029 ± 0.011, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.468 per %), hit=0.307 ± 0.018 per rating point (10 rating = 1%, 3.066 per %), spell_haste=not significant (-0.106 ± 0.146), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.876 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.14 DPS) | yes | Spellpower Goggles Xtreme (10502, -2.03 DPS) [crafted]; Dreamweave Circlet (10041, -2.04 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.51 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.2 spell_power points (2.45 DPS) | yes | Mindburst Medallion (11196, -0.36 DPS) [quest]; Scorn's Icy Choker (23169, -1.17 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.31 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.5 spell_power points (4.58 DPS) | yes | Black Mageweave Shoulders (10027, -0.85 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -1.44 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.45 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.2 spell_power points (4.80 DPS) | yes | Runecloth Cloak (13860, -1.61 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.75 DPS) [dungeon]; Icy Cloak (4327, -2.43 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.2 spell_power points (7.51 DPS) | yes | Elemental Raiment (9434, -0.40 DPS) [world_drop]; Acumen Robes (17775, -0.88 DPS) [quest]; Dreamweave Vest (10021, -1.32 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.05 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.61 DPS) [crafted]; Condor Bracers (15864, -0.68 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.1 spell_power points (6.13 DPS) | yes | Black Mageweave Gloves (10003, -0.60 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.64 DPS) [vendor]; Brightcloth Gloves (14101, -1.73 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.3 spell_power points (4.84 DPS) | yes | Highlander's Cloth Girdle (20098, -0.06 DPS) [rep]; Ghostweave Cord (254073, -0.10 DPS) [crafted]; Star Belt (4329, -0.44 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.43 DPS) | yes | Knight's Dreadweave Leggings (220888, -2.09 DPS) [vendor]; Senior Designer's Pantaloons (11841, -2.30 DPS) [dungeon]; Red Mageweave Pants (10009, -2.81 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.13 DPS) | yes | Gilded Sandals (254107, -4.31 DPS) [crafted]; Black Mageweave Boots (10026, -4.33 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.37 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.40 DPS) | yes | Philanthropist's Ring (281635, -0.97 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.69 DPS) [quest]; Reedknot Ring (9622, -2.03 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.06 DPS) | yes | Philanthropist's Ring (281635, -1.14 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.35 DPS) [quest]; Reedknot Ring (9622, -1.69 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -3.13 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.97 DPS, sim-verified) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.56 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.68 DPS) [dungeon]; Hanzo Sword (8190, -29.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (243.6 DPS) | yes | Lesser Eternal Wand (249232, -3.23 DPS) [crafted]; Wand of Allistarj (13065, -3.27 DPS) [world_drop]; Pyric Caduceus (11748, -9.70 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 488, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash

### Band 60 (gnome, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 431.3. Weights run: 3.0s. Verify run: 4.1s. 938 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.006 ± 0.022, crit=0.062 ± 0.003 per rating point (14 rating = 1%, 0.868 per %), hit=0.340 ± 0.032 per rating point (10 rating = 1%, 3.400 per %), spell_haste=not significant (0.218 ± 0.368), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.902 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.24 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -2.28 DPS) [pvp]; Deathmist Mask (226909, -2.47 DPS) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.97 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.14 DPS) [quest]; Kezan's Taint (19604, -3.60 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.0 spell_power points (12.67 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.29 DPS) [pvp]; Argent Shoulders (19059, -3.43 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -4.54 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 19.4 spell_power points (8.82 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -2.45 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -2.47 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.1 spell_power points (20.88 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.30 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.49 DPS) [pvp]; Robe of the Magi (1716, -11.71 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.0 spell_power points (9.99 DPS) | yes | Runecloth Cuffs (254123, -4.98 DPS) [crafted]; Arcane Runed Bracers (4744, -5.92 DPS) [quest]; Spidertank Oilrag (9448, -5.92 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.0 spell_power points (12.25 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -2.72 DPS) [pvp]; Earth Warder's Gloves (21318, -4.08 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 28.5 spell_power points (12.94 DPS) | yes | Stormpike Cloth Girdle (19094, -4.75 DPS) [rep]; Highlander's Cloth Girdle (20047, -6.18 DPS) [rep]; Belt of the Archmage (18405, -7.87 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 33.8 spell_power points (15.32 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -2.59 DPS) [pvp]; Ghostweave Pants (14144, -3.53 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.88 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.91 DPS) [crafted]; Omnicast Boots (11822, -1.78 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 31.5 spell_power points (14.26 DPS) | yes | Lorekeeper's Ring (19522, -7.91 DPS) [rep]; Blessed Band of Light (272407, -8.18 DPS) [vendor]; Band of the Unicorn (7553, -8.37 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.1 spell_power points (10.00 DPS) | yes | Blessed Band of Light (272407, -3.93 DPS) [vendor]; Band of the Unicorn (7553, -4.11 DPS) [world_drop]; Lorekeeper's Ring (19522, -6.47 DPS, sim-verified) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (431.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (431.3 DPS) | yes | Abyss Shard (20534, -0.91 DPS) [quest]; Weakness Analyzer (272438, -2.74 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.63 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (431.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.87 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -39.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 175.3 spell_power points (79.48 DPS) | yes | Wand of Biting Cold (19108, -6.69 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.80 DPS) [crafted]; Stormrager (16997, -16.78 DPS) [quest] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 938, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash

### Band 60, raid preset (gnome, 25550300100201351-0005200000000000000-0550001000000000)

Set DPS (verified): 850.3. Weights run: 3.2s. Verify run: 4.4s. 938 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.075 ± 0.034, crit=0.198 ± 0.012 per rating point (14 rating = 1%, 2.778 per %), hit=0.675 ± 0.067 per rating point (10 rating = 1%, 6.747 per %), spell_haste=not significant (1.858 ± 0.794), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.897 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.57 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -0.28 DPS) [quest]; Lieutenant Commander's Dreadweave Cowl (227093, -1.15 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.50 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.70 DPS) [quest]; Kezan's Taint (19604, -4.54 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.9 spell_power points (18.96 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.84 DPS) [pvp]; Mantle of the Timbermaw (19050, -6.23 DPS) [crafted]; Argent Shoulders (19059, -6.71 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.3 spell_power points (14.33 DPS) | yes | Hide of the Wild (18510, -5.28 DPS) [crafted]; Spritecaster Cape (11623, -5.46 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -5.74 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (28.64 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.91 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.38 DPS) [pvp]; Robe of the Magi (1716, -15.32 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.6 spell_power points (13.87 DPS) | yes | Runecloth Cuffs (254123, -6.66 DPS) [crafted]; Arcane Runed Bracers (4744, -8.34 DPS) [quest]; Spidertank Oilrag (9448, -8.34 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (16.80 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -3.73 DPS) [pvp]; Earth Warder's Gloves (21318, -5.87 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.4 spell_power points (20.49 DPS) | yes | Stormpike Cloth Girdle (19094, -8.99 DPS) [rep]; Belt of the Archmage (18405, -9.84 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -9.92 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 38.3 spell_power points (23.51 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -5.73 DPS) [pvp]; Ghostweave Pants (14144, -7.55 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.73 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.23 DPS) [crafted]; Omnicast Boots (11822, -1.91 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.4 spell_power points (21.74 DPS) | yes | Blessed Band of Light (272407, -11.46 DPS) [vendor]; Lorekeeper's Ring (19522, -13.15 DPS) [rep]; Band of the Unicorn (7553, -13.76 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.8 spell_power points (14.01 DPS) | yes | Blessed Band of Light (272407, -3.73 DPS) [vendor]; Lorekeeper's Ring (19522, -5.41 DPS) [rep]; Band of the Unicorn (7553, -6.03 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+21.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Abyss Shard (20534, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -1.23 DPS) [quest]; Weakness Analyzer (272438, -3.60 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.87 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.68 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -52.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (850.3 DPS) | yes | Wand of Biting Cold (19108, -1.42 DPS) [quest]; Stormrager (16997, -2.51 DPS) [quest]; Torch of Light (279246, -50.15 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 938, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 39.6. Weights run: 1.4s. Verify run: 1.5s. 181 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.101 ± 0.005, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.476 per %), hit=0.118 ± 0.006 per rating point (10 rating = 1%, 1.176 per %), spell_haste=not significant (-0.216 ± 0.075), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.747 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.14 DPS) | yes | Red Winter Hat (21524, -2.93 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.19 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.76 DPS) | yes | Battle Healer's Cloak (20427, -0.19 DPS) [rep]; Black Whelp Cloak (7283, -0.24 DPS, sim-verified) [crafted]; Mystic Shawl (3449, -0.38 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -1.25 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.57 DPS) | yes | Owlbeard Bracers (16981, -0.38 DPS) [quest]; Windsong Bangles (263336, -0.49 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.33 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.57 DPS) [quest]; Pristine Gloves (253913, -0.57 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (39.6 DPS) | yes | Novice Ardent's Sash (253887, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.99 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.71 DPS) | yes | Filigreed Pristine Leggings (253937, -0.57 DPS) [crafted]; Silk-threaded Trousers (1929, -0.67 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.76 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.33 DPS) | yes | Red Woolen Boots (4313, -0.73 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.76 DPS) [crafted]; Soft-soled Linen Boots (4312, -0.95 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.95 DPS) | yes | - |
| finger2 | - | - |  |  |  |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Witching Stave (1484), Clear Crystal Rod (16894), and 83 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 7 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 120.2 spell_power points (22.88 DPS) | yes | Skycaller (12984, -2.35 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.28 DPS) [quest]; Greater Magic Wand (11288, -5.00 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 181, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1917 Jeweled Dagger; 2235 Brackclaw; 2241 Desperado Cape; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash; 3563 Seafarer's Pantaloons; 3738 Brewing Rod

### Band 30 (troll, 25552200000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 67.4. Weights run: 1.4s. Verify run: 1.4s. 302 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.146 ± 0.007, crit=0.036 ± 0.002 per rating point (14 rating = 1%, 0.507 per %), hit=0.131 ± 0.007 per rating point (10 rating = 1%, 1.310 per %), spell_haste=not significant (-0.013 ± 0.106), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.775 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Filigreed Pristine Circlet (253975, -0.68 DPS) [crafted]; Enchanter's Cowl (4322, -0.81 DPS) [crafted]; Silk Headband (7050, -3.44 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (1.80 DPS) | yes | Crystal Starfire Medallion (5003, -1.66 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.66 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.82 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (2.35 DPS) | yes | Invoker's Mantle (215365, -0.59 DPS) [crafted]; Death Speaker Mantle (6685, -0.62 DPS) [dungeon]; Chestnut Mantle (17695, -1.06 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.14 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Battle Healer's Cloak (19529, -0.23 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.97 DPS) | yes | Beguiler Robes (7728, -0.65 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.74 DPS) [dungeon]; Green Silk Armor (7065, -0.78 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.05 DPS) | yes | Tabitha's Cuffs (251486, -1.72 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -1.76 DPS) [quest]; Glowing Magical Bracelets (13106, -1.79 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.60 DPS) | yes | Jutebraid Gloves (10654, -0.06 DPS) [quest]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.36 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.4 spell_power points (2.61 DPS) | yes | Warsong Sash (16975, -0.10 DPS) [quest]; Belt of Arugal (6392, -0.46 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.78 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.74 DPS) | yes | Abomination Skin Leggings (23173, -0.75 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.91 DPS) [crafted]; Silk-threaded Trousers (1929, -1.14 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (1.83 DPS) | yes | Acidic Walkers (9454, -0.42 DPS) [dungeon]; Boots of the Enchanter (4325, -0.69 DPS) [crafted]; Spidersilk Boots (4320, -1.71 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.60 DPS) | yes | Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Sacred Band (6669, -1.14 DPS) [quest]; Snake Hoop (6750, -1.36 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.37 DPS) | yes | Sacred Band (6669, -0.91 DPS) [quest]; Snake Hoop (6750, -1.14 DPS) [quest]; Electrocutioner Lagnut (9447, -1.53 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 1.6 spell_power points (0.37 DPS) | yes | Twisted Chanter's Staff (890, -0.03 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.03 DPS) [quest]; Polished Walking Staff (16889, -0.13 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-verified (67.4 DPS) | yes | Unstable Power Core (279847, -2.07 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -2.85 DPS) [crafted]; Gravestone Scepter (7001, -3.99 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Glimmering Staff; ranged: Starfaller

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff; 3229 Tarantula Silk Sash

### Band 40 (troll, 25552300120201201-0000000000000000000-0000000000000000)

Set DPS (verified): 171.1. Weights run: 1.5s. Verify run: 1.6s. 381 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.709 ± 0.020, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.369 per %), hit=0.198 ± 0.015 per rating point (10 rating = 1%, 1.977 per %), spell_haste=not significant (0.495 ± 0.182), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.875 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.19 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -2.51 DPS) [dungeon]; Enchanter's Cowl (4322, -2.71 DPS) [crafted]; Augural Shroud (2620, -3.69 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (4.34 DPS) | yes | Scorn's Icy Choker (23169, -1.71 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.64 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.64 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (5.55 DPS) | yes | Bloodmage Mantle (7684, -0.29 DPS) [dungeon]; Berylline Pads (4197, -0.73 DPS) [quest]; Green Silken Shoulders (7057, -1.60 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (3.27 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.60 DPS) [vendor]; Icy Cloak (4327, -0.87 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.3 spell_power points (8.99 DPS) | yes | Dreamweave Vest (10021, -1.22 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.28 DPS) [crafted]; Zealot's Robe (17043, -1.43 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.69 DPS) [quest]; Radiant Silver Bracers (4545, -1.84 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.14 DPS) | yes | Black Mageweave Gloves (10003, -2.00 DPS) [crafted]; Gilded Handwraps (254021, -2.70 DPS) [crafted]; Red Mageweave Gloves (10018, -2.85 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -1.08 DPS) [crafted]; Star Belt (4329, -1.31 DPS) [crafted]; Deathmage Sash (10771, -2.65 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (7.71 DPS) | yes | Crimson Silk Pantaloons (7062, -2.11 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.68 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.37 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.22 DPS) | yes | Gilded Slippers (254001, -3.88 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.57 DPS) [dungeon]; Spidersilk Boots (4320, -4.85 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.64 DPS) | yes | Reedknot Ring (9622, -2.24 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.58 DPS) [vendor]; Voodoo Band (1996, -2.94 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.08 DPS) | yes | Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor]; Reedknot Ring (9622, -1.05 DPS, sim-verified) [quest]; Voodoo Band (1996, -1.38 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.49 DPS) [dungeon]; Windweaver Staff (7757, -13.38 DPS) [dungeon] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.80 DPS) | yes | Thrash's Trash (276204, +0.00 DPS, sim-verified) [vendor]; Orb of Mystic Insight (249394, -0.94 DPS) [crafted]; Omega Orb (7749, -1.71 DPS) [quest] |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dancing Flame (6806, -0.27 DPS) [quest]; Umbral Wand (5216, -0.39 DPS) [dungeon]; Jaina's Firestarter (13064, -4.34 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Twisted Nether Wand

No-known-source sample (15 of 381, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff

### Band 50 (troll, 25552300120201351-2002000000000000000-0000000000000000)

Set DPS (verified): 242.0. Weights run: 2.9s. Verify run: 1.9s. 463 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.029 ± 0.011, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.468 per %), hit=0.307 ± 0.018 per rating point (10 rating = 1%, 3.066 per %), spell_haste=not significant (-0.106 ± 0.146), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.876 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.14 DPS) | yes | Spellpower Goggles Xtreme (10502, -2.03 DPS) [crafted]; Dreamweave Circlet (10041, -2.11 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.51 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.2 spell_power points (2.45 DPS) | yes | Mindburst Medallion (11196, -0.36 DPS) [quest]; Scorn's Icy Choker (23169, -1.32 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.31 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.5 spell_power points (4.58 DPS) | yes | Bloodmage Mantle (7684, -1.44 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.45 DPS) [crafted]; Black Mageweave Shoulders (10027, -2.10 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.2 spell_power points (4.80 DPS) | yes | Deep Woodlands Cloak (19121, -0.71 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.67 DPS) [crafted]; Nightfall Drape (12465, -1.75 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.2 spell_power points (7.51 DPS) | yes | Acumen Robes (17775, -0.88 DPS) [quest]; Dreamweave Vest (10021, -1.32 DPS) [crafted]; Elemental Raiment (9434, -1.61 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.05 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.1 spell_power points (6.13 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.64 DPS) [vendor]; Brightcloth Gloves (14101, -1.73 DPS) [crafted]; Black Mageweave Gloves (10003, -2.83 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.3 spell_power points (4.84 DPS) | yes | Ghostweave Cord (254073, -0.10 DPS) [crafted]; Star Belt (4329, -0.44 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.77 DPS, sim-verified) [rep] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.43 DPS) | yes | Red Mageweave Pants (10009, -1.71 DPS, sim-verified) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.09 DPS) [vendor]; Senior Designer's Pantaloons (11841, -2.30 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.13 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -4.08 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -4.31 DPS) [crafted]; Black Mageweave Boots (10026, -4.33 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.40 DPS) | yes | Philanthropist's Ring (281635, -0.97 DPS) [quest]; Reedknot Ring (9622, -2.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.37 DPS) [vendor] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.06 DPS) | yes | Philanthropist's Ring (281635, -0.90 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.69 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.03 DPS) [vendor] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -3.13 DPS) [crafted]; Rune of the Guard Captain (19120, -3.34 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.30 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.77 DPS, sim-verified) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.56 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.68 DPS) [dungeon]; Hanzo Sword (8190, -29.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (242.0 DPS) | yes | Lesser Eternal Wand (249232, -3.23 DPS) [crafted]; Wand of Allistarj (13065, -3.27 DPS) [world_drop]; Pyric Caduceus (11748, -9.64 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 463, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff

### Band 60 (troll, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 424.3. Weights run: 3.0s. Verify run: 4.0s. 928 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.006 ± 0.022, crit=0.062 ± 0.003 per rating point (14 rating = 1%, 0.868 per %), hit=0.340 ± 0.032 per rating point (10 rating = 1%, 3.400 per %), spell_haste=not significant (0.218 ± 0.368), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.902 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.24 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -2.28 DPS) [pvp]; Deathmist Mask (226909, -2.47 DPS) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.97 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.14 DPS) [quest]; Kezan's Taint (19604, -3.60 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.0 spell_power points (12.67 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.29 DPS) [pvp]; Argent Shoulders (19059, -1.80 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -4.54 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 19.4 spell_power points (8.82 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -2.45 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -2.47 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.1 spell_power points (20.88 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.30 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.49 DPS) [pvp]; Robe of the Magi (1716, -10.87 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.0 spell_power points (9.99 DPS) | yes | Runecloth Cuffs (254123, -4.98 DPS) [crafted]; Spidertank Oilrag (9448, -5.92 DPS) [dungeon]; Earthenweave Cuffs (254131, -6.37 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.0 spell_power points (12.25 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.01 DPS, sim-verified) [quest]; Blood Guard's Dreadweave Handwraps (227099, -2.72 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 28.5 spell_power points (12.94 DPS) | yes | Frostwolf Cloth Belt (19090, -4.75 DPS) [rep]; Belt of the Archmage (18405, -5.03 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -6.18 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 33.8 spell_power points (15.32 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.66 DPS, sim-verified) [rep]; Legionnaire's Dreadweave Legguards (227097, -2.59 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.88 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.91 DPS) [crafted]; Omnicast Boots (11822, -1.78 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 31.5 spell_power points (14.26 DPS) | yes | Advisor's Ring (19518, -7.91 DPS) [rep]; Blessed Band of Light (272407, -8.18 DPS) [vendor]; Band of the Unicorn (7553, -8.37 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.1 spell_power points (10.00 DPS) | yes | Advisor's Ring (19518, -3.68 DPS, sim-verified) [rep]; Blessed Band of Light (272407, -3.93 DPS) [vendor]; Band of the Unicorn (7553, -4.11 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (424.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (424.3 DPS) | yes | Abyss Shard (20534, -0.91 DPS) [quest]; Weakness Analyzer (272438, -2.69 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.63 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (424.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.87 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -37.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 175.3 spell_power points (79.48 DPS) | yes | Wand of Biting Cold (19108, -6.59 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.80 DPS) [crafted]; Stormrager (16997, -16.78 DPS) [quest] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 928, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff

### Band 60, raid preset (troll, 25550300100201351-0005200000000000000-0550001000000000)

Set DPS (verified): 840.2. Weights run: 3.2s. Verify run: 4.1s. 928 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.075 ± 0.034, crit=0.198 ± 0.012 per rating point (14 rating = 1%, 2.778 per %), hit=0.675 ± 0.067 per rating point (10 rating = 1%, 6.747 per %), spell_haste=not significant (1.858 ± 0.794), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.897 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.57 DPS) | yes | Deathmist Mask (226909, +0.00 DPS, sim-verified) [quest]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.15 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.50 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.70 DPS) [quest]; Kezan's Taint (19604, -4.54 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.9 spell_power points (18.96 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.84 DPS) [pvp]; Argent Shoulders (19059, -3.62 DPS) [crafted]; Mantle of the Timbermaw (19050, -6.23 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.3 spell_power points (14.33 DPS) | yes | Hide of the Wild (18510, -5.28 DPS) [crafted]; Spritecaster Cape (11623, -5.46 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -5.74 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (28.64 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.91 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.38 DPS) [pvp]; Robe of the Magi (1716, -13.93 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.6 spell_power points (13.87 DPS) | yes | Runecloth Cuffs (254123, -6.66 DPS) [crafted]; Spidertank Oilrag (9448, -8.34 DPS) [dungeon]; Earthenweave Cuffs (254131, -8.96 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (16.80 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.30 DPS) [quest]; Blood Guard's Dreadweave Handwraps (227099, -3.73 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.4 spell_power points (20.49 DPS) | yes | Belt of the Archmage (18405, -3.60 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -8.99 DPS) [rep]; Defiler's Cloth Girdle (20163, -9.92 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 38.3 spell_power points (23.51 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -3.36 DPS, sim-verified) [rep]; Legionnaire's Dreadweave Legguards (227097, -5.73 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.73 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.23 DPS) [crafted]; Omnicast Boots (11822, -1.91 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.4 spell_power points (21.74 DPS) | yes | Blessed Band of Light (272407, -11.46 DPS) [vendor]; Advisor's Ring (19518, -13.15 DPS) [rep]; Band of the Unicorn (7553, -13.76 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.8 spell_power points (14.01 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Advisor's Ring (19518, -5.41 DPS) [rep]; Band of the Unicorn (7553, -6.03 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+21.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Abyss Shard (20534, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -1.23 DPS) [quest]; Weakness Analyzer (272438, -3.52 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -4.87 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.68 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -48.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (840.2 DPS) | yes | Wand of Biting Cold (19108, -1.42 DPS) [quest]; Stormrager (16997, -2.51 DPS) [quest]; Torch of Light (279246, -53.78 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 928, see the JSON for more): 1189 Overseer's Ring; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1917 Jeweled Dagger; 2035 Sword of the Night Sky; 2235 Brackclaw; 2241 Desperado Cape; 2622 Nimar's Tribal Headdress; 2944 Cursed Eye of Paleth; 3018 Hide of Lupos; 3019 Noble's Robe; 3224 Silver-lined Bracers; 3225 Bloodstained Knife; 3227 Nightbane Staff

