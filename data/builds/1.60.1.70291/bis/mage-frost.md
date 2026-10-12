# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 31.7. Weights run: 1.2s. Verify run: 1.3s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.304 ± 0.009, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.014 per %), hit=0.262 ± 0.011 per rating point (10 rating = 1%, 2.616 per %), spell_haste=not significant (0.583 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.63 DPS) | yes | Shadow Goggles (4373, -1.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.7 spell_power points (0.81 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.39 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.71 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.42 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Caretaker's Cape (20428, -0.10 DPS) [rep]; Black Whelp Cloak (7283, -0.52 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.68 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -0.98 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.5 spell_power points (0.16 DPS) | yes | Windsong Bangles (263336, -0.05 DPS) [quest]; Repurposed Hair Band (281256, -0.10 DPS) [quest]; Bright Bracers (3647, -0.55 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.73 DPS) | yes | Pristine Gloves (253913, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Heavy Woolen Gloves (4310, -0.46 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.55 DPS) | yes | Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Keller's Girdle (2911, -0.29 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.76 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.4 spell_power points (1.20 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.46 DPS) [dungeon]; Rumpled Kilt (274741, -0.67 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.86 DPS) | yes | Red Woolen Boots (4313, -0.44 DPS) [crafted]; Pristine Boots (253889, -0.45 DPS) [crafted]; Feather Padded Treads (285345, -0.57 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.59 DPS) | yes | Sludge-Stained Band (286535, -0.27 DPS) [world]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.49 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.52 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop]; Sludge-Stained Band (286535, -0.74 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.0 spell_power points (0.32 DPS) | yes | Lesser Staff of the Spire (1300, -0.13 DPS) [world]; Staff of Westfall (2042, -0.16 DPS) [quest]; Channeler's Staff (4437, -0.50 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 215.7 spell_power points (22.62 DPS) | yes | Skycaller (12984, -1.01 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.33 DPS) [dungeon]; Deepblaze (279896, -4.09 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 000000000000000000-00000000000000000-2535101301000000000)

Set DPS (verified): 49.1. Weights run: 1.2s. Verify run: 1.4s. 260 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.148 ± 0.006, crit=0.097 ± 0.004 per rating point (14 rating = 1%, 1.355 per %), hit=0.253 ± 0.009 per rating point (10 rating = 1%, 2.533 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.26 DPS) | yes | Embalmed Shroud (7691, -0.34 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Silk Headband (7050, -1.78 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (0.91 DPS) | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.84 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.13 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (1.19 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Invoker's Mantle (215365, -0.60 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (0.69 DPS) | yes | Hillman's Cloak (3719, -0.11 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Prelacy Cape (7004, -0.23 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.08 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.13 DPS) [dungeon]; Tree Bark Jacket (1486, -0.96 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Glowing Magical Bracelets (13106, -0.56 DPS, sim-verified) [world_drop]; Windsong Bangles (263336, -0.92 DPS) [quest]; Nightsky Wristbands (6407, -0.93 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (0.80 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.11 DPS) [world]; Town Clerk's Mittens (270029, -0.16 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.4 spell_power points (1.31 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.40 DPS) [dungeon]; Invoker's Cord (215366, -0.43 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.38 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.46 DPS) [crafted]; Silk-threaded Trousers (1929, -0.57 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (0.92 DPS) | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Nimbus Boots (6998, -0.23 DPS) [quest]; Spidersilk Boots (4320, -1.25 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.80 DPS) | yes | Minor Channeling Ring (1449, -0.20 DPS) [quest]; Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Sludge-Stained Band (286535, -0.34 DPS) [world]; Minor Channeling Ring (1449, -0.38 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Hardened Root Staff (1317, -1.49 DPS, sim-verified) [quest]; Glimmering Staff (249392, -2.54 DPS) [crafted]; Twisted Chanter's Staff (890, -2.56 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 292.3 spell_power points (33.55 DPS) | yes | Starfaller (13063, -0.63 DPS) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 000000000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 91.3. Weights run: 1.2s. Verify run: 1.6s. 334 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.682 ± 0.040, crit=0.168 ± 0.011 per rating point (14 rating = 1%, 2.349 per %), hit=0.420 ± 0.043 per rating point (10 rating = 1%, 4.196 per %), spell_haste=not significant (1.845 ± 0.644), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.00 DPS) | yes | Augural Shroud (2620, -0.45 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -1.08 DPS) [dungeon]; Living Cowl (5608, -1.14 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.5 spell_power points (1.78 DPS) | yes | Scorn's Icy Choker (23169, -0.94 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.10 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.10 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.9 spell_power points (2.26 DPS) | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -0.88 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.4 spell_power points (1.34 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.27 DPS) [vendor]; Icy Cloak (4327, -0.34 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.1 spell_power points (3.72 DPS) | yes | Dreamweave Vest (10021, -0.28 DPS) [crafted]; Robe of Power (7054, -0.56 DPS) [crafted]; Elemental Raiment (9434, -0.73 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (91.3 DPS) | yes | Condor Bracers (15864, -0.29 DPS) [quest]; Windchaser Cuffs (14429, -0.41 DPS) [world_drop]; Arcane Runed Bracers (4744, -1.20 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 spell_power points (2.96 DPS) | yes | Black Mageweave Gloves (10003, -0.82 DPS) [crafted]; Red Mageweave Gloves (10018, -0.85 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.14 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.2 spell_power points (2.46 DPS) | yes | Gilded Cord (254037, -0.54 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.60 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (3.17 DPS) | yes | Crimson Silk Pantaloons (7062, -0.76 DPS) [crafted]; Abomination Skin Leggings (23173, -1.10 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.39 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.43 DPS) | yes | Gilded Slippers (254001, -1.75 DPS) [crafted]; Acidic Walkers (9454, -1.93 DPS) [dungeon]; Spidersilk Boots (4320, -2.04 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (1.91 DPS) | yes | Ring of Forlorn Spirits (2043, -0.77 DPS) [quest]; Reedknot Ring (9622, -0.91 DPS) [quest]; Minor Channeling Ring (1449, -1.01 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.28 DPS) | yes | Ring of Forlorn Spirits (2043, -0.14 DPS) [quest]; Reedknot Ring (9622, -0.29 DPS) [quest]; Minor Channeling Ring (1449, -0.38 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Hardened Root Staff (1317, -2.78 DPS) [quest] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.0 spell_power points (2.15 DPS) | yes | Orb of the Forgotten Seer (7685, -0.15 DPS) [dungeon]; Thrash's Trash (276204, -0.15 DPS) [vendor]; Orb of Lorica (11262, -0.32 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 279.9 spell_power points (39.96 DPS) | yes | Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest]; Ragefire Wand (7513, -2.53 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 334, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 203005000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 223.8. Weights run: 1.2s. Verify run: 1.9s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.011, intellect=0.969 ± 0.061, crit=0.285 ± 0.015 per rating point (14 rating = 1%, 3.985 per %), hit=0.943 ± 0.083 per rating point (10 rating = 1%, 9.428 per %), spell_haste=8.086 ± 1.154, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 38.4 spell_power points (7.88 DPS) | yes | Dreamweave Circlet (10041, -1.58 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.80 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.34 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.8 spell_power points (3.03 DPS) | yes | Horizon Choker (13085, -0.24 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.40 DPS) [dungeon]; Mindburst Medallion (11196, -0.60 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 30.4 spell_power points (6.25 DPS) | yes | Red Mageweave Shoulders (10029, -1.83 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.00 DPS) [vendor]; Inquisitor's Shawl (19507, -2.23 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.8 spell_power points (4.07 DPS) | yes | Runecloth Cloak (13860, -0.63 DPS) [crafted]; Big Voodoo Cloak (8216, -1.25 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.28 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 38.4 spell_power points (7.88 DPS) | yes | Robe of the Magi (1716, -2.17 DPS) [world_drop]; Runecloth Tunic (13857, -2.20 DPS) [crafted]; Runecloth Robe (13858, -2.24 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 13.8 spell_power points (2.83 DPS) | yes | Bloodband Bracers (11469, -0.01 DPS) [quest]; Forgotten Wraps (9433, -0.44 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.44 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 35.0 spell_power points (7.19 DPS) | yes | Dreamweave Gloves (10019, -2.69 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.73 DPS) [vendor]; Raider Handwraps (272098, -3.80 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 24.4 spell_power points (5.01 DPS) | yes | Satyrmane Sash (17755, -0.15 DPS) [dungeon]; Deathmage Sash (10771, -0.59 DPS) [dungeon]; Runecloth Belt (13856, -0.98 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 27.6 spell_power points (5.67 DPS) | yes | Red Mageweave Pants (10009, -0.41 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.44 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.47 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 26.1 spell_power points (5.37 DPS) | yes | Earthen Silk Slippers (254013, -0.44 DPS) [crafted]; Gilded Sandals (254107, -1.32 DPS) [crafted]; Southsea Mojo Boots (20641, -1.54 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (3.05 DPS) | yes | Lorekeeper's Ring (19523, -0.58 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.41 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.67 DPS) | yes | Lorekeeper's Ring (19523, -0.21 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.03 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (223.8 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (223.8 DPS) | yes | Mark of the Chosen (17774, -3.89 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (223.8 DPS) | yes | Kindling Stave (11750, -0.96 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.59 DPS) [dungeon]; Hanzo Sword (8190, -28.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 255.6 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.47 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.11 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 203005000000000000-11302300000000000-2535101301000300250)

Set DPS (verified): 356.5. Weights run: 1.2s. Verify run: 5.8s. 876 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.010, intellect=0.737 ± 0.077, crit=0.377 ± 0.020 per rating point (14 rating = 1%, 5.279 per %), hit=1.314 ± 0.108 per rating point (10 rating = 1%, 13.137 per %), spell_haste=8.458 ± 1.366, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.010

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 44.7 spell_power points (9.77 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.13 DPS) [pvp]; Red Mageweave Headband (10033, -8.99 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (356.5 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.56 DPS) [quest]; Chains of the Lich (23125, -0.56 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 43.3 spell_power points (9.47 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.59 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.51 DPS) [crafted]; Darkspear Shoulderpads (272103, -7.45 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.0 spell_power points (7.66 DPS) | yes | Hide of the Wild (18510, -2.99 DPS) [crafted]; Spritecaster Cape (11623, -3.63 DPS) [dungeon]; Runecloth Cloak (13860, -4.40 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 54.1 spell_power points (11.83 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.72 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.35 DPS) [pvp]; Sorcerer's Robes (226932, -9.09 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 27.9 spell_power points (6.10 DPS) | yes | Runecloth Cuffs (254123, -2.08 DPS) [crafted]; Sorcerer's Bindings (226929, -2.96 DPS) [quest]; Bracers of Hope (22667, -3.01 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (356.5 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, -6.01 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 54.4 spell_power points (11.88 DPS) | yes | Magician's Cord (272393, -3.86 DPS) [vendor]; Stormpike Cloth Girdle (19094, -6.33 DPS) [rep]; Belt of the Archmage (18405, -7.86 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.9 spell_power points (10.91 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -1.31 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -2.43 DPS) [pvp] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (356.5 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.86 DPS) [dungeon]; Sorcerer's Boots (22064, -7.64 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 47.8 spell_power points (10.44 DPS) | yes | Channeler's Ring (272406, -5.31 DPS) [vendor]; Blessed Band of Light (272407, -5.38 DPS) [vendor]; Deep Rooted Ring (19109, -6.21 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 30.1 spell_power points (6.58 DPS) | yes | Blessed Band of Light (272407, -1.52 DPS) [vendor]; Deep Rooted Ring (19109, -2.35 DPS) [world]; Channeler's Ring (272406, -5.00 DPS, sim-verified) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (356.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (356.5 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.48 DPS) [crafted]; Uther's Strength (11302, -1.75 DPS) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (356.5 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.46 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -34.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 357.2 spell_power points (78.07 DPS) | yes | Wand of Biting Cold (19108, -7.16 DPS, sim-verified) [quest]; Stormrager (16997, -14.58 DPS) [quest]; Brilliant Wand (249385, -15.40 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 876, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 203005000000000000-13102000000000000-2555100321000300250)

Set DPS (verified): 595.7. Weights run: 1.3s. Verify run: 4.2s. 876 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.210 ± 0.025, crit=0.399 ± 0.018 per rating point (14 rating = 1%, 5.583 per %), hit=1.060 ± 0.065 per rating point (10 rating = 1%, 10.598 per %), spell_haste=6.961 ± 0.427, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.8 spell_power points (13.58 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.63 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.40 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (595.7 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.82 DPS) [quest]; Jewel of Kajaro (19601, -5.22 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.7 spell_power points (15.25 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.24 DPS) [pvp]; Argent Shoulders (19059, -4.58 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.03 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.3 spell_power points (12.07 DPS) | yes | Hide of the Wild (18510, -5.20 DPS) [crafted]; Spritecaster Cape (11623, -5.55 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -6.09 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.1 spell_power points (20.52 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.54 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -7.66 DPS) [pvp]; Robe of the Magi (1716, -12.52 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.7 spell_power points (10.10 DPS) | yes | Runecloth Cuffs (254123, -4.51 DPS) [crafted]; Sorcerer's Bindings (226929, -6.22 DPS) [quest]; Arcane Runed Bracers (4744, -6.26 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.1 spell_power points (11.97 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 40.2 spell_power points (17.16 DPS) | yes | Belt of the Archmage (18405, -5.74 DPS, sim-verified) [crafted]; Magician's Cord (272393, -6.90 DPS) [vendor]; Highlander's Cloth Girdle (20047, -8.27 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.3 spell_power points (19.31 DPS) | yes | Marshal's Silk Leggings (231605, -2.34 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -6.45 DPS) [pvp]; Sorcerer's Leggings (226933, -6.84 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 24.4 spell_power points (10.39 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.15 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 40.5 spell_power points (17.28 DPS) | yes | Blessed Band of Light (272407, -8.49 DPS) [vendor]; Don Julio's Band (19325, -10.37 DPS) [rep]; Deep Rooted Ring (19109, -11.26 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 24.3 spell_power points (10.37 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Don Julio's Band (19325, -3.47 DPS) [rep]; Deep Rooted Ring (19109, -4.36 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (595.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -7.10 DPS, sim-verified) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (595.7 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Uther's Strength (11302, -1.51 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (595.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.22 DPS) [world]; Teebu's Blazing Longsword (1728, -39.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 185.9 spell_power points (79.32 DPS) | yes | Brilliant Wand (249385, -15.27 DPS) [crafted]; Wand of Biting Cold (19108, -15.32 DPS) [quest]; Stormrager (16997, -16.18 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 876, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 29.1. Weights run: 1.2s. Verify run: 1.4s. 136 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.304 ± 0.009, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.014 per %), hit=0.262 ± 0.011 per rating point (10 rating = 1%, 2.616 per %), spell_haste=not significant (0.583 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.63 DPS) | yes | Shadow Goggles (4373, -1.00 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.7 spell_power points (0.81 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.31 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.39 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.42 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.10 DPS) [rep]; Black Whelp Cloak (7283, -0.26 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.68 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -0.65 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.31 DPS) | yes | Owlbeard Bracers (16981, -0.15 DPS) [quest]; Mindthrust Bracers (1974, -0.16 DPS) [dungeon]; Featherbead Bracers (15452, -0.16 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.73 DPS) | yes | Gnoll Casting Gloves (892, -0.21 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.22 DPS) [crafted]; Apothecary Gloves (10919, -0.31 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.55 DPS) | yes | Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Keller's Girdle (2911, -0.29 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.42 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.4 spell_power points (1.20 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.46 DPS) [dungeon]; Rumpled Kilt (274741, -0.67 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.86 DPS) | yes | Red Woolen Boots (4313, -0.44 DPS) [crafted]; Pristine Boots (253889, -0.45 DPS) [crafted]; Feather Padded Treads (285345, -0.61 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.52 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Loop of Sacrifice (281673, -0.37 DPS) [quest]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.31 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.16 DPS) [quest]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.0 spell_power points (0.32 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.13 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 215.7 spell_power points (22.62 DPS) | yes | Skycaller (12984, -1.03 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.33 DPS) [dungeon]; Sizzle Stick (8071, -4.45 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 136, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (troll, 000000000000000000-00000000000000000-2535101301000000000)

Set DPS (verified): 47.2. Weights run: 1.2s. Verify run: 1.4s. 242 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.148 ± 0.006, crit=0.097 ± 0.004 per rating point (14 rating = 1%, 1.355 per %), hit=0.253 ± 0.009 per rating point (10 rating = 1%, 2.533 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.26 DPS) | yes | Embalmed Shroud (7691, -0.34 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Silk Headband (7050, -1.83 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (0.91 DPS) | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.84 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.01 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (1.19 DPS) | yes | Invoker's Mantle (215365, -0.30 DPS) [crafted]; Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Chestnut Mantle (17695, -0.97 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.57 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.08 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.13 DPS) [dungeon]; Tree Bark Jacket (1486, -1.04 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Owlbeard Bracers (16981, -0.88 DPS) [quest]; Glowing Magical Bracelets (13106, -0.90 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.26 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.80 DPS) | yes | Jutebraid Gloves (10654, -0.03 DPS) [quest]; Gnoll Casting Gloves (892, -0.11 DPS) [world]; Truefaith Gloves (7049, -0.18 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.4 spell_power points (1.31 DPS) | yes | Warsong Sash (16975, -0.05 DPS) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.40 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.38 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.46 DPS) [crafted]; Silk-threaded Trousers (1929, -0.57 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (0.92 DPS) | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Boots of the Enchanter (4325, -0.35 DPS) [crafted]; Spidersilk Boots (4320, -1.27 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.80 DPS) | yes | Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world]; Sacred Band (6669, -0.57 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Sacred Band (6669, -0.46 DPS) [quest]; Sludge-Stained Band (286535, -1.04 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 23.7 spell_power points (2.73 DPS) | yes | Twisted Chanter's Staff (890, -2.56 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -2.56 DPS) [quest]; Glimmering Staff (249392, -5.66 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Starfaller (13063, -0.63 DPS) [world_drop]; Unstable Power Core (279847, -1.02 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.98 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (troll, 000000000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 86.9. Weights run: 1.2s. Verify run: 1.6s. 312 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.682 ± 0.040, crit=0.168 ± 0.011 per rating point (14 rating = 1%, 2.349 per %), hit=0.420 ± 0.043 per rating point (10 rating = 1%, 4.196 per %), spell_haste=not significant (1.845 ± 0.644), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.00 DPS) | yes | Augural Shroud (2620, -0.45 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -1.08 DPS) [dungeon]; Living Cowl (5608, -1.14 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.5 spell_power points (1.78 DPS) | yes | Scorn's Icy Choker (23169, -0.19 DPS) [dungeon]; Necklace of Calisea (1714, -1.10 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.10 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.9 spell_power points (2.26 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.4 spell_power points (1.34 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.27 DPS) [vendor]; Icy Cloak (4327, -0.34 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.1 spell_power points (3.72 DPS) | yes | Dreamweave Vest (10021, -0.28 DPS) [crafted]; Robe of Power (7054, -0.56 DPS) [crafted]; Zealot's Robe (17043, -0.61 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (86.9 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.29 DPS) [quest]; Radiant Silver Bracers (4545, -0.87 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 spell_power points (2.96 DPS) | yes | Red Mageweave Gloves (10018, -0.42 DPS) [crafted]; Black Mageweave Gloves (10003, -0.82 DPS) [crafted]; Gilded Handwraps (254021, -1.14 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.2 spell_power points (2.46 DPS) | yes | Gilded Cord (254037, -0.54 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.60 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (3.17 DPS) | yes | Crimson Silk Pantaloons (7062, -0.76 DPS) [crafted]; Abomination Skin Leggings (23173, -1.10 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.39 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.43 DPS) | yes | Gilded Slippers (254001, -0.73 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.93 DPS) [dungeon]; Spidersilk Boots (4320, -2.04 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (1.91 DPS) | yes | Reedknot Ring (9622, -0.91 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.06 DPS) [vendor]; Black Widow Band (6199, -1.23 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.28 DPS) | yes | Sea Giant's Toe Ring (274746, -0.43 DPS) [vendor]; Black Widow Band (6199, -0.60 DPS) [world]; Reedknot Ring (9622, -0.79 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Wind Spirit Staff (6689, -3.29 DPS) [dungeon] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.0 spell_power points (2.15 DPS) | yes | Orb of the Forgotten Seer (7685, -0.15 DPS) [dungeon]; Thrash's Trash (276204, -0.15 DPS) [vendor]; Orb of Mystic Insight (249394, -0.56 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 279.9 spell_power points (39.96 DPS) | yes | Nether Force Wand (11263, -1.58 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.48 DPS) [quest]; Ragefire Wand (7513, -2.53 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 312, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (troll, 203005000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 216.3. Weights run: 1.2s. Verify run: 1.8s. 398 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.011, intellect=0.969 ± 0.061, crit=0.285 ± 0.015 per rating point (14 rating = 1%, 3.985 per %), hit=0.943 ± 0.083 per rating point (10 rating = 1%, 9.428 per %), spell_haste=8.086 ± 1.154, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 38.4 spell_power points (7.88 DPS) | yes | Dreamweave Circlet (10041, -1.58 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.80 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.34 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.8 spell_power points (3.03 DPS) | yes | Horizon Choker (13085, -0.24 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.40 DPS) [dungeon]; Mindburst Medallion (11196, -0.60 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 30.4 spell_power points (6.25 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -2.00 DPS) [vendor]; Inquisitor's Shawl (19507, -2.23 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.94 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.7 spell_power points (4.26 DPS) | yes | Spritecaster Cape (11623, -0.19 DPS) [dungeon]; Runecloth Cloak (13860, -0.82 DPS) [crafted]; Big Voodoo Cloak (8216, -1.44 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 38.4 spell_power points (7.88 DPS) | yes | Robe of the Magi (1716, -2.17 DPS) [world_drop]; Runecloth Tunic (13857, -2.20 DPS) [crafted]; Runecloth Robe (13858, -2.24 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 13.8 spell_power points (2.83 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.01 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 35.0 spell_power points (7.19 DPS) | yes | Dreamweave Gloves (10019, -2.69 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.73 DPS) [vendor]; Raider Handwraps (272098, -3.96 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 24.4 spell_power points (5.01 DPS) | yes | Satyrmane Sash (17755, -0.15 DPS) [dungeon]; Deathmage Sash (10771, -0.59 DPS) [dungeon]; Runecloth Belt (13856, -0.98 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 27.6 spell_power points (5.67 DPS) | yes | Red Mageweave Pants (10009, -0.41 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.44 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.47 DPS) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.1 spell_power points (5.37 DPS) | yes | Earthen Silk Slippers (254013, -0.44 DPS) [crafted]; Gilded Sandals (254107, -1.32 DPS) [crafted]; Southsea Mojo Boots (20641, -1.54 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (3.05 DPS) | yes | Advisor's Ring (19519, -0.58 DPS) [rep]; Reedknot Ring (9622, -1.61 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.67 DPS) | yes | Advisor's Ring (19519, -0.21 DPS) [rep]; Reedknot Ring (9622, -1.23 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (216.3 DPS) | yes | Uther's Strength (11302, -0.51 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (216.3 DPS) | yes | Kindling Stave (11750, -0.96 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.59 DPS) [dungeon]; Hanzo Sword (8190, -31.09 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 255.6 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.47 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.11 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 398, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (troll, 203005000000000000-11302300000000000-2535101301000300250)

Set DPS (verified): 343.3. Weights run: 1.2s. Verify run: 5.7s. 866 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.010, intellect=0.737 ± 0.077, crit=0.377 ± 0.020 per rating point (14 rating = 1%, 5.279 per %), hit=1.314 ± 0.108 per rating point (10 rating = 1%, 13.137 per %), spell_haste=8.458 ± 1.366, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.010

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 44.7 spell_power points (9.77 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.13 DPS) [pvp]; Red Mageweave Headband (10033, -12.08 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (343.3 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.56 DPS) [quest]; Chains of the Lich (23125, -0.56 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 43.3 spell_power points (9.47 DPS) | yes | Warlord's Silk Amice (231594, -1.59 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.50 DPS) [vendor]; Mantle of the Timbermaw (19050, -2.51 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.0 spell_power points (7.66 DPS) | yes | Hide of the Wild (18510, -2.99 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.58 DPS) [quest]; Spritecaster Cape (11623, -3.63 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 54.1 spell_power points (11.83 DPS) | yes | Warlord's Silk Raiment (231596, -0.72 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.35 DPS) [pvp]; Sorcerer's Robes (226932, -8.04 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.9 spell_power points (6.10 DPS) | yes | Runecloth Cuffs (254123, -2.08 DPS) [crafted]; Sorcerer's Bindings (226929, -2.96 DPS) [quest]; Bracers of Hope (22667, -3.01 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (343.3 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, -6.77 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 54.4 spell_power points (11.88 DPS) | yes | Magician's Cord (272393, -3.86 DPS) [vendor]; Belt of the Archmage (18405, -6.33 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -6.33 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.9 spell_power points (10.91 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -1.31 DPS) [quest]; Outrider's Silk Leggings (22747, -1.73 DPS) [rep] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (343.3 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.86 DPS) [dungeon]; Sorcerer's Boots (22064, -7.34 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 47.8 spell_power points (10.44 DPS) | yes | Channeler's Ring (272406, -5.31 DPS) [vendor]; Blessed Band of Light (272407, -5.38 DPS) [vendor]; Deep Rooted Ring (19109, -6.21 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 30.1 spell_power points (6.58 DPS) | yes | Channeler's Ring (272406, -1.45 DPS) [vendor]; Blessed Band of Light (272407, -1.52 DPS) [vendor]; Deep Rooted Ring (19109, -2.35 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (343.3 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (343.3 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.48 DPS) [crafted]; Rune of the Guard Captain (19120, -1.05 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (343.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.46 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -33.54 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 357.2 spell_power points (78.07 DPS) | yes | Wand of Biting Cold (19108, -14.07 DPS) [quest]; Stormrager (16997, -14.58 DPS) [quest]; Brilliant Wand (249385, -15.40 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 866, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (troll, 203005000000000000-13102000000000000-2555100321000300250)

Set DPS (verified): 602.3. Weights run: 1.3s. Verify run: 4.1s. 866 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.210 ± 0.025, crit=0.399 ± 0.018 per rating point (14 rating = 1%, 5.583 per %), hit=1.060 ± 0.065 per rating point (10 rating = 1%, 10.598 per %), spell_haste=6.961 ± 0.427, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.8 spell_power points (13.58 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.63 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -4.10 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (602.3 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.82 DPS) [quest]; Jewel of Kajaro (19601, -4.56 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.7 spell_power points (15.25 DPS) | yes | Warlord's Silk Amice (231594, -3.24 DPS) [pvp]; Argent Shoulders (19059, -4.58 DPS) [crafted]; Mantle of the Timbermaw (19050, -4.72 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.3 spell_power points (12.07 DPS) | yes | Hide of the Wild (18510, -5.20 DPS) [crafted]; Spritecaster Cape (11623, -5.55 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -6.09 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.1 spell_power points (20.52 DPS) | yes | Warlord's Silk Raiment (231596, -2.54 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -7.66 DPS) [pvp]; Robe of the Magi (1716, -12.08 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.7 spell_power points (10.10 DPS) | yes | Runecloth Cuffs (254123, -4.51 DPS) [crafted]; Sorcerer's Bindings (226929, -6.22 DPS) [quest]; Spidertank Oilrag (9448, -6.26 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.1 spell_power points (11.97 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 40.2 spell_power points (17.16 DPS) | yes | Belt of the Archmage (18405, -4.98 DPS, sim-verified) [crafted]; Magician's Cord (272393, -6.90 DPS) [vendor]; Defiler's Cloth Girdle (20163, -8.27 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.3 spell_power points (19.31 DPS) | yes | General's Silk Trousers (231595, -2.34 DPS) [pvp]; Outrider's Silk Leggings (22747, -5.66 DPS) [rep]; Sorcerer's Leggings (226933, -7.33 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 24.4 spell_power points (10.39 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.15 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 40.5 spell_power points (17.28 DPS) | yes | Blessed Band of Light (272407, -8.49 DPS) [vendor]; Don Julio's Band (19325, -10.37 DPS) [rep]; Deep Rooted Ring (19109, -11.26 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 24.3 spell_power points (10.37 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Don Julio's Band (19325, -3.47 DPS) [rep]; Deep Rooted Ring (19109, -4.36 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (602.3 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (602.3 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.90 DPS) [crafted]; Rune of the Guard Captain (19120, -2.81 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (602.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.22 DPS) [world]; Teebu's Blazing Longsword (1728, -38.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 185.9 spell_power points (79.32 DPS) | yes | Brilliant Wand (249385, -15.27 DPS) [crafted]; Wand of Biting Cold (19108, -15.32 DPS) [quest]; Stormrager (16997, -16.18 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 866, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

