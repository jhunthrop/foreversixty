# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 41.9. Weights run: 1.4s. Verify run: 1.0s. 151 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.104 ± 0.003, crit=0.033 ± 0.001 per rating point (14 rating = 1%, 0.457 per %), hit=0.111 ± 0.004 per rating point (10 rating = 1%, 1.113 per %), spell_haste=not significant (-0.197 ± 0.052), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.746 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.14 DPS) | yes | Red Winter Hat (21524, -2.97 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.19 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.76 DPS) | yes | Feyscale Cloak (6632, -0.19 DPS) [dungeon]; Caretaker's Cape (20428, -0.19 DPS) [rep]; Black Whelp Cloak (7283, -0.39 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -1.44 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.19 DPS) | yes | Ivycloth Bracelets (9793, -0.25 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.33 DPS) | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.57 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.95 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (41.9 DPS) | yes | Novice Ardent's Sash (253887, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.87 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Filigreed Pristine Leggings (253937, -0.57 DPS) [crafted]; Silk-threaded Trousers (1929, -0.66 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.76 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.33 DPS) | yes | Feather Padded Treads (285345, -0.57 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.57 DPS) [crafted]; Pristine Boots (253889, -0.76 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.95 DPS) | yes | Sludge-Stained Band (286535, -0.38 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.95 DPS) | yes | Sludge-Stained Band (286535, -0.57 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.36 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.95 DPS) | yes | Bouquet of Red Roses (22206, -1.30 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 120.8 spell_power points (22.88 DPS) | yes | Skycaller (12984, -2.40 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.59 DPS) [dungeon]; Sizzle Stick (8071, -4.28 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 151, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 25552200000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 82.4. Weights run: 1.5s. Verify run: 0.9s. 266 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.144 ± 0.005, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.497 per %), hit=0.122 ± 0.005 per rating point (10 rating = 1%, 1.220 per %), spell_haste=not significant (0.035 ± 0.075), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.775 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Embalmed Shroud (7691, -0.68 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.68 DPS) [crafted]; Silk Headband (7050, -3.53 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (1.79 DPS) | yes | Crystal Starfire Medallion (5003, -1.66 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.66 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.83 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (2.35 DPS) | yes | Invoker's Mantle (215365, -0.59 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.62 DPS) [dungeon]; Fairywing Mantle (9536, -0.68 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (1.37 DPS) | yes | Hillman's Cloak (3719, -0.26 DPS, sim-verified) [crafted]; Heavy Woolen Cloak (4311, -0.46 DPS) [crafted]; Prelacy Cape (7004, -0.46 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.96 DPS) | yes | Beguiler Robes (7728, -0.65 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.74 DPS) [dungeon]; Green Silk Armor (7065, -0.77 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.05 DPS) | yes | Windsong Bangles (263336, -1.82 DPS) [quest]; Nightsky Wristbands (6407, -1.85 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.21 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.60 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Town Clerk's Mittens (270029, -0.32 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.4 spell_power points (2.61 DPS) | yes | Belt of Arugal (6392, -0.28 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.78 DPS) [dungeon]; Invoker's Cord (215366, -0.85 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.73 DPS) | yes | Pristine Leggings (253987, -0.91 DPS) [crafted]; Abomination Skin Leggings (23173, -0.91 DPS, sim-verified) [dungeon]; Silk-threaded Trousers (1929, -1.14 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (1.83 DPS) | yes | Acidic Walkers (9454, -0.42 DPS) [dungeon]; Nimbus Boots (6998, -0.46 DPS) [quest]; Spidersilk Boots (4320, -1.93 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.60 DPS) | yes | Minor Channeling Ring (1449, -0.39 DPS) [quest]; Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Sludge-Stained Band (286535, -0.91 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.37 DPS) | yes | Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon]; Sludge-Stained Band (286535, -0.68 DPS) [world]; Minor Channeling Ring (1449, -2.01 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (82.4 DPS) | yes | Hardened Root Staff (1317, -2.12 DPS, sim-verified) [quest]; Scorn's Focal Dagger (23168, -3.35 DPS) [dungeon]; Glimmering Staff (249392, -5.04 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 148.7 spell_power points (33.89 DPS) | yes | Starfaller (13063, -1.42 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.75 DPS) [crafted]; Gravestone Scepter (7001, -4.89 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 266, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 25552300120201201-0000000000000000000-0000000000000000)

Set DPS (verified): 173.2. Weights run: 1.5s. Verify run: 1.3s. 346 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.702 ± 0.015, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.364 per %), hit=0.190 ± 0.010 per rating point (10 rating = 1%, 1.899 per %), spell_haste=0.541 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.874 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.19 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -2.53 DPS) [dungeon]; Corpseshroud (10574, -2.62 DPS) [dungeon]; Augural Shroud (2620, -3.42 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.6 spell_power points (4.32 DPS) | yes | Scorn's Icy Choker (23169, -1.64 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.92 DPS) [quest]; Darkspear Warding Pendant (272074, -2.64 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 spell_power points (5.52 DPS) | yes | Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.72 DPS) [quest]; Green Silken Shoulders (7057, -1.52 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (173.2 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Mantle of Lady Falther'ess (23178, +0.00 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -0.61 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (8.97 DPS) | yes | Dreamweave Vest (10021, -0.84 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.30 DPS) [crafted]; Elemental Raiment (9434, -1.78 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.08 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.68 DPS) [quest]; Windchaser Cuffs (14429, -0.92 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.12 DPS) | yes | Black Mageweave Gloves (10003, -1.99 DPS) [crafted]; Red Mageweave Gloves (10018, -2.55 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.70 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -1.09 DPS) [crafted]; Star Belt (4329, -1.30 DPS) [crafted]; Deathmage Sash (10771, -1.82 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 spell_power points (7.68 DPS) | yes | Crimson Silk Pantaloons (7062, -2.17 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.67 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.36 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.22 DPS) | yes | Gilded Slippers (254001, -3.75 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.58 DPS) [dungeon]; Spidersilk Boots (4320, -4.86 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.62 DPS) | yes | Ring of Forlorn Spirits (2043, -1.89 DPS) [quest]; Reedknot Ring (9622, -2.23 DPS) [quest]; Minor Channeling Ring (1449, -2.43 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (3.08 DPS) | yes | Ring of Forlorn Spirits (2043, -0.64 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.68 DPS) [quest]; Minor Channeling Ring (1449, -0.89 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.48 DPS) [dungeon]; Hardened Root Staff (1317, -6.71 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.79 DPS) | yes | Thrash's Trash (276204, +0.00 DPS, sim-verified) [vendor]; Orb of Lorica (11262, -0.34 DPS) [quest]; Orb of Mystic Insight (249394, -0.95 DPS) [crafted] |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Umbral Wand (5216, -0.38 DPS) [dungeon]; Earthen Rod (9381, -0.46 DPS) [dungeon]; Jaina's Firestarter (13064, -3.70 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Twisted Nether Wand

No-known-source sample (15 of 346, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25552300120201351-2002000000000000000-0000000000000000)

Set DPS (verified): 252.3. Weights run: 3.6s. Verify run: 1.4s. 436 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.027 ± 0.008, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.441 per %), hit=0.299 ± 0.013 per rating point (10 rating = 1%, 2.988 per %), spell_haste=not significant (-0.119 ± 0.107), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.878 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.33 DPS) | yes | Dreamweave Circlet (10041, -0.63 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.07 DPS) [crafted]; Red Mageweave Headband (10033, -2.58 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.2 spell_power points (2.49 DPS) | yes | Mindburst Medallion (11196, -0.36 DPS) [quest]; Scorn's Icy Choker (23169, -1.73 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.36 DPS) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.4 spell_power points (4.96 DPS) | yes | Black Mageweave Shoulders (10027, -1.42 DPS) [crafted]; Bloodmage Mantle (7684, -1.77 DPS) [dungeon]; Rotgrip Mantle (17732, -2.96 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.2 spell_power points (4.89 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.08 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.71 DPS) [crafted]; Nightfall Drape (12465, -1.78 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.2 spell_power points (7.66 DPS) | yes | Elemental Raiment (9434, -0.40 DPS) [world_drop]; Acumen Robes (17775, -0.91 DPS) [quest]; Dreamweave Vest (10021, -1.35 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.11 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.63 DPS) [crafted]; Condor Bracers (15864, -0.69 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.1 spell_power points (6.26 DPS) | yes | Black Mageweave Gloves (10003, -1.49 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.68 DPS) [vendor]; Brightcloth Gloves (14101, -1.77 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | sim-verified (252.3 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.06 DPS) [rep]; Ghostweave Cord (254073, -0.09 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 23.3 spell_power points (8.04 DPS) | yes | Red Mageweave Pants (10009, -3.09 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.63 DPS) [vendor]; Wizardweave Leggings (14132, -4.29 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.29 DPS) | yes | Gilded Sandals (254107, -1.00 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.41 DPS) [vendor]; Black Mageweave Boots (10026, -4.43 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.49 DPS) | yes | Philanthropist's Ring (281635, -0.99 DPS) [quest]; Cyclopean Band (11824, -1.32 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.73 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.15 DPS) | yes | Philanthropist's Ring (281635, -0.74 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.97 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.38 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.07 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.58 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.69 DPS) [dungeon]; Blade of Eternal Darkness (17780, -28.90 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+10.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.21 DPS) [crafted]; Wand of Allistarj (13065, -3.31 DPS) [world_drop]; Pyric Caduceus (11748, -10.05 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 468.8. Weights run: 4.0s. Verify run: 3.1s. 1064 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.079 ± 0.017, crit=0.069 ± 0.002 per rating point (14 rating = 1%, 0.961 per %), hit=0.392 ± 0.024 per rating point (10 rating = 1%, 3.920 per %), spell_haste=not significant (0.746 ± 0.257), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.898 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.6 spell_power points (13.30 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -2.96 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -6.98 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.55 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.60 DPS) [quest]; Kezan's Taint (19604, -3.20 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.1 spell_power points (12.65 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.22 DPS) [pvp]; Burial Shawl (18681, -3.42 DPS) [dungeon]; Argent Shoulders (19059, -3.75 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.5 spell_power points (8.92 DPS) | yes | Crystalline Threaded Cape (20697, -0.10 DPS) [world]; Amplifying Cloak (18350, -1.11 DPS) [dungeon]; Hide of the Wild (18510, -2.50 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (20.28 DPS) | yes | Robe of Everlasting Night (18385, -4.88 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -5.57 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -8.74 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.6 spell_power points (9.83 DPS) | yes | Sublime Wristguards (18497, -4.27 DPS) [dungeon]; Runecloth Cuffs (254123, -4.71 DPS) [crafted]; Arcane Runed Bracers (4744, -5.92 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (11.90 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.34 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.6 spell_power points (13.31 DPS) | yes | Stormpike Cloth Girdle (19094, -5.15 DPS) [rep]; Ban'thok Sash (11662, -6.02 DPS) [dungeon]; Belt of the Archmage (18405, -6.67 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 34.7 spell_power points (15.07 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -2.47 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.42 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.87 DPS) [crafted]; Omnicast Boots (11822, -1.33 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -5.52 DPS) [dungeon]; Maiden's Circle (13001, -6.11 DPS) [world_drop]; Naglering (11669, -18.69 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.28 DPS) [dungeon]; Maiden's Circle (13001, -1.87 DPS) [world_drop]; Naglering (11669, -12.10 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+18.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.04 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.75 DPS, sim-verified) [quest]; Serenity Field (272439, -6.51 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.46 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -39.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (468.8 DPS) | yes | Bonecreeper Stylus (13938, -1.08 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.82 DPS) [world]; Torch of Light (279246, -17.20 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1064, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 25550300100201351-0005200000000000000-0550001000000000)

Set DPS (verified): 872.3. Weights run: 4.5s. Verify run: 3.1s. 1064 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.074 ± 0.025, crit=0.169 ± 0.007 per rating point (14 rating = 1%, 2.364 per %), hit=0.652 ± 0.047 per rating point (10 rating = 1%, 6.523 per %), spell_haste=not significant (1.374 ± 0.626), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.892 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.6 spell_power points (18.18 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -2.56 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -5.18 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.08 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -3.50 DPS) [dungeon]; Amulet of the Dawn (22657, -3.59 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.5 spell_power points (18.11 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.51 DPS) [pvp]; Argent Shoulders (19059, -3.25 DPS) [crafted]; Burial Shawl (18681, -5.52 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.1 spell_power points (13.74 DPS) | yes | Crystalline Threaded Cape (20697, -1.67 DPS) [world]; Amplifying Cloak (18350, -3.04 DPS) [dungeon]; Hide of the Wild (18510, -4.98 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (27.73 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.66 DPS) [pvp]; Robe of Everlasting Night (18385, -11.12 DPS) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -12.00 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.6 spell_power points (13.42 DPS) | yes | Sublime Wristguards (18497, -5.86 DPS) [dungeon]; Runecloth Cuffs (254123, -6.45 DPS) [crafted]; Arcane Runed Bracers (4744, -8.08 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (16.27 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.19 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.1 spell_power points (19.70 DPS) | yes | Belt of the Archmage (18405, -4.86 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -8.21 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -8.56 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 37.5 spell_power points (22.27 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -5.06 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.26 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.19 DPS) [crafted]; Omnicast Boots (11822, -1.85 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -7.53 DPS) [dungeon]; Maiden's Circle (13001, -9.91 DPS) [world_drop]; Naglering (11669, -20.33 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.17 DPS) [dungeon]; Maiden's Circle (13001, -2.55 DPS) [world_drop]; Naglering (11669, -10.88 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+23.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -3.57 DPS) [quest]; Weakness Analyzer (272438, -4.16 DPS) [vendor]; Serenity Field (272439, -8.91 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.66 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -48.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (872.3 DPS) | yes | Bonecreeper Stylus (13938, -1.05 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.69 DPS) [world]; Torch of Light (279246, -48.88 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1064, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 40.3. Weights run: 1.4s. Verify run: 1.0s. 140 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.104 ± 0.003, crit=0.033 ± 0.001 per rating point (14 rating = 1%, 0.457 per %), hit=0.111 ± 0.004 per rating point (10 rating = 1%, 1.113 per %), spell_haste=not significant (-0.197 ± 0.052), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.746 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.14 DPS) | yes | Red Winter Hat (21524, -2.94 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.19 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.76 DPS) | yes | Feyscale Cloak (6632, -0.19 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.19 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -1.35 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.57 DPS) | yes | Owlbeard Bracers (16981, -0.38 DPS) [quest]; Windsong Bangles (263336, -0.49 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.33 DPS) | yes | Gnoll Casting Gloves (892, -0.31 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.57 DPS) [quest]; Pristine Gloves (253913, -0.57 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (40.3 DPS) | yes | Novice Ardent's Sash (253887, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Filigreed Pristine Leggings (253937, -0.57 DPS) [crafted]; Rumpled Kilt (274741, -0.76 DPS) [vendor]; Silk-threaded Trousers (1929, -0.81 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.33 DPS) | yes | Feather Padded Treads (285345, -0.54 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.57 DPS) [crafted]; Pristine Boots (253889, -0.76 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.95 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.57 DPS) | yes | Ring of the Shadow (1462, -0.78 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 9 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 120.8 spell_power points (22.88 DPS) | yes | Skycaller (12984, -2.18 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.59 DPS) [dungeon]; Sizzle Stick (8071, -4.28 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 25552200000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 81.1. Weights run: 1.5s. Verify run: 1.0s. 249 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.144 ± 0.005, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.497 per %), hit=0.122 ± 0.005 per rating point (10 rating = 1%, 1.220 per %), spell_haste=not significant (0.035 ± 0.075), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.775 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Embalmed Shroud (7691, -0.68 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.68 DPS) [crafted]; Silk Headband (7050, -3.68 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (1.79 DPS) | yes | Crystal Starfire Medallion (5003, -1.66 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.66 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.77 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (2.35 DPS) | yes | Invoker's Mantle (215365, -0.59 DPS) [crafted]; Death Speaker Mantle (6685, -0.62 DPS) [dungeon]; Chestnut Mantle (17695, -2.11 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.14 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Battle Healer's Cloak (19529, -0.23 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.96 DPS) | yes | Beguiler Robes (7728, -0.65 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.74 DPS) [dungeon]; Green Silk Armor (7065, -1.04 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.05 DPS) | yes | Owlbeard Bracers (16981, -1.76 DPS) [quest]; Glowing Magical Bracelets (13106, -1.79 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.98 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.60 DPS) | yes | Gnoll Casting Gloves (892, -0.23 DPS) [world]; Jutebraid Gloves (10654, -0.35 DPS, sim-verified) [quest]; Truefaith Gloves (7049, -0.36 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.4 spell_power points (2.61 DPS) | yes | Warsong Sash (16975, -0.31 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.46 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.78 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.73 DPS) | yes | Abomination Skin Leggings (23173, -0.78 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.91 DPS) [crafted]; Silk-threaded Trousers (1929, -1.14 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (1.83 DPS) | yes | Acidic Walkers (9454, -0.42 DPS) [dungeon]; Boots of the Enchanter (4325, -0.69 DPS) [crafted]; Spidersilk Boots (4320, -2.20 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.60 DPS) | yes | Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Sludge-Stained Band (286535, -0.91 DPS) [world]; Sacred Band (6669, -1.14 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.37 DPS) | yes | Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon]; Sacred Band (6669, -0.91 DPS) [quest]; Sludge-Stained Band (286535, -2.16 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 23.7 spell_power points (5.41 DPS) | yes | Glimmering Staff (249392, -5.04 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -5.08 DPS) [quest]; Scorn's Focal Dagger (23168, -10.01 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (81.1 DPS) | yes | Starfaller (13063, -0.90 DPS) [world_drop]; Unstable Power Core (279847, -3.64 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.75 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552300120201201-0000000000000000000-0000000000000000)

Set DPS (verified): 173.9. Weights run: 1.5s. Verify run: 1.1s. 324 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.702 ± 0.015, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.364 per %), hit=0.190 ± 0.010 per rating point (10 rating = 1%, 1.899 per %), spell_haste=0.541 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.874 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.19 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -2.53 DPS) [dungeon]; Corpseshroud (10574, -2.62 DPS) [dungeon]; Augural Shroud (2620, -3.03 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.6 spell_power points (4.32 DPS) | yes | Scorn's Icy Choker (23169, -1.27 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.92 DPS) [quest]; Darkspear Warding Pendant (272074, -2.64 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 spell_power points (5.52 DPS) | yes | Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.72 DPS) [quest]; Green Silken Shoulders (7057, -1.15 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.3 spell_power points (5.24 DPS) | yes | Long Silken Cloak (4326, -1.76 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -1.99 DPS) [crafted]; Darkspear Raider's Cloak (272077, -2.60 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (8.97 DPS) | yes | Dreamweave Vest (10021, -0.54 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.30 DPS) [crafted]; Zealot's Robe (17043, -1.43 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.68 DPS) [quest]; Radiant Silver Bracers (4545, -2.26 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.12 DPS) | yes | Black Mageweave Gloves (10003, -1.99 DPS) [crafted]; Gilded Handwraps (254021, -2.70 DPS) [crafted]; Red Mageweave Gloves (10018, -2.84 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -1.09 DPS) [crafted]; Star Belt (4329, -1.30 DPS) [crafted]; Deathmage Sash (10771, -2.50 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 spell_power points (7.68 DPS) | yes | Crimson Silk Pantaloons (7062, -1.84 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.67 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.36 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.22 DPS) | yes | Gilded Slippers (254001, -3.76 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.58 DPS) [dungeon]; Spidersilk Boots (4320, -4.86 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.62 DPS) | yes | Reedknot Ring (9622, -2.23 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.57 DPS) [vendor]; Black Widow Band (6199, -2.94 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.08 DPS) | yes | Reedknot Ring (9622, -0.49 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor]; Black Widow Band (6199, -1.40 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.48 DPS) [dungeon]; Wind Spirit Staff (6689, -7.91 DPS) [dungeon] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.79 DPS) | yes | Thrash's Trash (276204, +0.00 DPS, sim-verified) [vendor]; Orb of Mystic Insight (249394, -0.95 DPS) [crafted]; Omega Orb (7749, -1.71 DPS) [quest] |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Dancing Flame (6806, -0.28 DPS) [quest]; Umbral Wand (5216, -0.38 DPS) [dungeon]; Jaina's Firestarter (13064, -4.90 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Twisted Nether Wand

No-known-source sample (15 of 324, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552300120201351-2002000000000000000-0000000000000000)

Set DPS (verified): 251.1. Weights run: 3.6s. Verify run: 1.4s. 410 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.027 ± 0.008, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.441 per %), hit=0.299 ± 0.013 per rating point (10 rating = 1%, 2.988 per %), spell_haste=not significant (-0.119 ± 0.107), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.878 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.33 DPS) | yes | Dreamweave Circlet (10041, -0.98 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.07 DPS) [crafted]; Red Mageweave Headband (10033, -2.58 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.2 spell_power points (2.49 DPS) | yes | Mindburst Medallion (11196, -0.36 DPS) [quest]; Scorn's Icy Choker (23169, -1.97 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.36 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (251.1 DPS) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Black Mageweave Shoulders (10027, -1.12 DPS) [crafted]; Bloodmage Mantle (7684, -1.47 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.2 spell_power points (4.89 DPS) | yes | Deep Woodlands Cloak (19121, -1.43 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.70 DPS) [dungeon]; Runecloth Cloak (13860, -1.71 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.2 spell_power points (7.66 DPS) | yes | Acumen Robes (17775, -0.91 DPS) [quest]; Elemental Raiment (9434, -1.04 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.35 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.11 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.1 spell_power points (6.26 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.68 DPS) [vendor]; Brightcloth Gloves (14101, -1.77 DPS) [crafted]; Black Mageweave Gloves (10003, -2.18 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.3 spell_power points (5.28 DPS) | yes | Defiler's Cloth Girdle (20166, -0.41 DPS) [rep]; Ghostweave Cord (254073, -0.44 DPS) [crafted]; Satyrmane Sash (17755, -3.12 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 23.3 spell_power points (8.04 DPS) | yes | Red Mageweave Pants (10009, -3.09 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.63 DPS) [vendor]; Wizardweave Leggings (14132, -4.95 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.29 DPS) | yes | Gilded Sandals (254107, -1.33 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.41 DPS) [vendor]; Black Mageweave Boots (10026, -4.43 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.49 DPS) | yes | Philanthropist's Ring (281635, -0.99 DPS) [quest]; Cyclopean Band (11824, -1.32 DPS) [dungeon]; Runed Ring (862, -2.07 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.15 DPS) | yes | Philanthropist's Ring (281635, -0.75 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.97 DPS) [dungeon]; Runed Ring (862, -1.73 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.07 DPS) [world_drop]; Rune of the Guard Captain (19120, -3.42 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.21 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.58 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.69 DPS) [dungeon]; Blade of Eternal Darkness (17780, -29.74 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+9.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.21 DPS) [crafted]; Wand of Allistarj (13065, -3.31 DPS) [world_drop]; Pyric Caduceus (11748, -9.67 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 462.5. Weights run: 4.0s. Verify run: 2.9s. 1052 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.079 ± 0.017, crit=0.069 ± 0.002 per rating point (14 rating = 1%, 0.961 per %), hit=0.392 ± 0.024 per rating point (10 rating = 1%, 3.920 per %), spell_haste=not significant (0.746 ± 0.257), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.898 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.6 spell_power points (13.30 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -2.96 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -8.06 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.55 DPS) | yes | Orb of the Darkmoon (19426, -1.97 DPS, sim-verified) [quest]; Amulet of the Dawn (22657, -2.60 DPS) [quest]; Kezan's Taint (19604, -3.20 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.1 spell_power points (12.65 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.22 DPS) [pvp]; Burial Shawl (18681, -3.42 DPS) [dungeon]; Argent Shoulders (19059, -3.58 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.5 spell_power points (8.92 DPS) | yes | Crystalline Threaded Cape (20697, -0.10 DPS) [world]; Amplifying Cloak (18350, -1.11 DPS) [dungeon]; Hide of the Wild (18510, -2.50 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (20.28 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.57 DPS) [pvp]; Robe of Everlasting Night (18385, -6.16 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -8.74 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.6 spell_power points (9.83 DPS) | yes | Sublime Wristguards (18497, -4.27 DPS) [dungeon]; Runecloth Cuffs (254123, -4.71 DPS) [crafted]; Spidertank Oilrag (9448, -5.92 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (11.90 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -0.40 DPS) [dungeon]; Earth Warder's Gloves (21318, -2.34 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.6 spell_power points (13.31 DPS) | yes | Frostwolf Cloth Belt (19090, -5.15 DPS) [rep]; Ban'thok Sash (11662, -6.02 DPS) [dungeon]; Belt of the Archmage (18405, -8.05 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 34.7 spell_power points (15.07 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.03 DPS) [dungeon]; Outrider's Silk Leggings (22747, -2.26 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.42 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.87 DPS) [crafted]; Omnicast Boots (11822, -1.33 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -5.52 DPS) [dungeon]; Maiden's Circle (13001, -6.11 DPS) [world_drop]; Naglering (11669, -19.63 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.28 DPS) [dungeon]; Maiden's Circle (13001, -1.87 DPS) [world_drop]; Naglering (11669, -12.81 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+17.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.61 DPS) [quest]; Weakness Analyzer (272438, -3.04 DPS) [vendor]; Serenity Field (272439, -6.51 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.46 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -39.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (462.5 DPS) | yes | Bonecreeper Stylus (13938, -1.08 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.82 DPS) [world]; Torch of Light (279246, -17.03 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1052, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 25550300100201351-0005200000000000000-0550001000000000)

Set DPS (verified): 864.7. Weights run: 4.5s. Verify run: 3.2s. 1052 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.074 ± 0.025, crit=0.169 ± 0.007 per rating point (14 rating = 1%, 2.364 per %), hit=0.652 ± 0.047 per rating point (10 rating = 1%, 6.523 per %), spell_haste=not significant (1.374 ± 0.626), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.892 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.6 spell_power points (18.18 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -2.56 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -11.19 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.08 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -3.50 DPS) [dungeon]; Amulet of the Dawn (22657, -3.59 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.5 spell_power points (18.11 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.51 DPS) [pvp]; Argent Shoulders (19059, -3.25 DPS) [crafted]; Burial Shawl (18681, -5.52 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.1 spell_power points (13.74 DPS) | yes | Amplifying Cloak (18350, -3.04 DPS) [dungeon]; Crystalline Threaded Cape (20697, -3.57 DPS, sim-verified) [world]; Hide of the Wild (18510, -4.98 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (27.73 DPS) | yes | Robe of Everlasting Night (18385, -4.67 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -7.66 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.00 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.6 spell_power points (13.42 DPS) | yes | Sublime Wristguards (18497, -5.86 DPS) [dungeon]; Runecloth Cuffs (254123, -6.45 DPS) [crafted]; Spidertank Oilrag (9448, -8.08 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (16.27 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.19 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.1 spell_power points (19.70 DPS) | yes | Ban'thok Sash (11662, -8.21 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -8.56 DPS) [rep]; Belt of the Archmage (18405, -9.91 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 37.5 spell_power points (22.27 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.71 DPS) [dungeon]; Outrider's Silk Leggings (22747, -4.79 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.26 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.19 DPS) [crafted]; Omnicast Boots (11822, -1.85 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -7.53 DPS) [dungeon]; Maiden's Circle (13001, -9.91 DPS) [world_drop]; Naglering (11669, -30.00 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.17 DPS) [dungeon]; Maiden's Circle (13001, -2.55 DPS) [world_drop]; Naglering (11669, -16.22 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -3.57 DPS) [quest]; Weakness Analyzer (272438, -4.16 DPS) [vendor]; Draconic Infused Emblem (22268, -23.21 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (864.7 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.66 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -52.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+50.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.05 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.69 DPS) [world]; Torch of Light (279246, -50.75 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1052, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

