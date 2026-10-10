# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 31.7. Weights run: 0.7s. Verify run: 0.6s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.315 ± 0.007, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.078 per %), hit=0.256 ± 0.007 per rating point (10 rating = 1%, 2.558 per %), spell_haste=0.429 ± 0.105, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.63 DPS) | yes | Shadow Goggles (4373, -1.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.8 spell_power points (0.82 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.71 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.42 DPS) | yes | Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Caretaker's Cape (20428, -0.11 DPS) [rep]; Black Whelp Cloak (7283, -0.52 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.6 spell_power points (0.69 DPS) | yes | Green Woolen Vest (2582, -0.27 DPS) [crafted]; Bloody Apron (6226, -0.27 DPS) [dungeon]; Gray Woolen Robe (2585, -0.98 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.6 spell_power points (0.17 DPS) | yes | Windsong Bangles (263336, -0.06 DPS) [quest]; Repurposed Hair Band (281256, -0.10 DPS) [quest]; Bright Bracers (3647, -0.55 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.74 DPS) | yes | Pristine Gloves (253913, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Heavy Woolen Gloves (4310, -0.46 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.3 spell_power points (0.55 DPS) | yes | Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Keller's Girdle (2911, -0.29 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.76 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.5 spell_power points (1.21 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.48 DPS) [dungeon]; Rumpled Kilt (274741, -0.69 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.3 spell_power points (0.87 DPS) | yes | Red Woolen Boots (4313, -0.45 DPS) [crafted]; Pristine Boots (253889, -0.45 DPS) [crafted]; Feather Padded Treads (285345, -0.57 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.59 DPS) | yes | Sludge-Stained Band (286535, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.49 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.53 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop]; Sludge-Stained Band (286535, -0.74 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.1 spell_power points (0.33 DPS) | yes | Lesser Staff of the Spire (1300, -0.13 DPS) [world]; Staff of Westfall (2042, -0.17 DPS) [quest]; Channeler's Staff (4437, -0.50 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 215.1 spell_power points (22.63 DPS) | yes | Skycaller (12984, -1.01 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.34 DPS) [dungeon]; Deepblaze (279896, -4.10 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-00000000000000000-2535101301000000000)

Set DPS (verified): 49.1. Weights run: 0.9s. Verify run: 0.7s. 266 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.154 ± 0.004, crit=0.099 ± 0.003 per rating point (14 rating = 1%, 1.390 per %), hit=0.248 ± 0.006 per rating point (10 rating = 1%, 2.483 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.27 DPS) | yes | Embalmed Shroud (7691, -0.35 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.35 DPS) [crafted]; Silk Headband (7050, -1.78 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (0.91 DPS) | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.84 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.13 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (1.19 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Invoker's Mantle (215365, -0.60 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (0.69 DPS) | yes | Hillman's Cloak (3719, -0.12 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Prelacy Cape (7004, -0.23 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.09 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.13 DPS) [dungeon]; Tree Bark Jacket (1486, -0.96 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Glowing Magical Bracelets (13106, -0.56 DPS, sim-verified) [world_drop]; Windsong Bangles (263336, -0.92 DPS) [quest]; Nightsky Wristbands (6407, -0.93 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (0.81 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.12 DPS) [world]; Town Clerk's Mittens (270029, -0.15 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (1.32 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.40 DPS) [dungeon]; Invoker's Cord (215366, -0.42 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.38 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.45 DPS) [crafted]; Silk-threaded Trousers (1929, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (0.93 DPS) | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Nimbus Boots (6998, -0.24 DPS) [quest]; Spidersilk Boots (4320, -1.25 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.81 DPS) | yes | Minor Channeling Ring (1449, -0.19 DPS) [quest]; Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Sludge-Stained Band (286535, -0.35 DPS) [world]; Minor Channeling Ring (1449, -0.38 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Hardened Root Staff (1317, -1.49 DPS, sim-verified) [quest]; Scorn's Focal Dagger (23168, -1.70 DPS) [dungeon]; Glimmering Staff (249392, -2.54 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.8 spell_power points (33.56 DPS) | yes | Starfaller (13063, -0.62 DPS) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 266, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 93.1. Weights run: 0.8s. Verify run: 0.8s. 346 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.662 ± 0.027, crit=0.167 ± 0.007 per rating point (14 rating = 1%, 2.332 per %), hit=0.421 ± 0.029 per rating point (10 rating = 1%, 4.205 per %), spell_haste=not significant (1.417 ± 0.447), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.01 DPS) | yes | Augural Shroud (2620, -0.48 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -1.10 DPS) [dungeon]; Living Cowl (5608, -1.15 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.3 spell_power points (1.76 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.81 DPS) [quest]; Scorn's Icy Choker (23169, -0.97 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.10 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 spell_power points (2.24 DPS) | yes | Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest]; Green Silken Shoulders (7057, -0.91 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.0 spell_power points (2.14 DPS) | yes | Guardian Cloak (5965, -0.81 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.10 DPS) [vendor]; Long Silken Cloak (4326, -2.03 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 spell_power points (3.72 DPS) | yes | Dreamweave Vest (10021, -0.29 DPS) [crafted]; Robe of Power (7054, -0.58 DPS) [crafted]; Elemental Raiment (9434, -0.71 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (93.1 DPS) | yes | Condor Bracers (15864, -0.29 DPS) [quest]; Windchaser Cuffs (14429, -0.44 DPS) [world_drop]; Arcane Runed Bracers (4744, -1.04 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (2.96 DPS) | yes | Red Mageweave Gloves (10018, -0.43 DPS) [crafted]; Black Mageweave Gloves (10003, -0.81 DPS) [crafted]; Gilded Handwraps (254021, -1.15 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.9 spell_power points (2.43 DPS) | yes | Highlander's Cloth Girdle (20098, -0.04 DPS) [rep]; Gilded Cord (254037, -0.52 DPS) [crafted]; Star Belt (4329, -0.56 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.9 spell_power points (3.15 DPS) | yes | Crimson Silk Pantaloons (7062, -1.09 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.10 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.44 DPS) | yes | Gilded Slippers (254001, -1.77 DPS) [crafted]; Acidic Walkers (9454, -1.96 DPS) [dungeon]; Spidersilk Boots (4320, -2.06 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (1.91 DPS) | yes | Ring of Forlorn Spirits (2043, -0.76 DPS) [quest]; Reedknot Ring (9622, -0.90 DPS) [quest]; Minor Channeling Ring (1449, -1.00 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.29 DPS) | yes | Reedknot Ring (9622, -0.29 DPS) [quest]; Minor Channeling Ring (1449, -0.38 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.61 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Hardened Root Staff (1317, -2.77 DPS) [quest] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.0 spell_power points (2.15 DPS) | yes | Orb of the Forgotten Seer (7685, -0.14 DPS) [dungeon]; Thrash's Trash (276204, -0.14 DPS) [vendor]; Orb of Lorica (11262, -0.34 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 278.7 spell_power points (39.95 DPS) | yes | Nether Force Wand (11263, -2.33 DPS) [quest]; Icefury Wand (7514, -2.47 DPS) [quest]; Ragefire Wand (7513, -2.52 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 346, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 203005000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 236.0. Weights run: 0.8s. Verify run: 1.0s. 443 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.986 ± 0.043, crit=0.270 ± 0.011 per rating point (14 rating = 1%, 3.775 per %), hit=0.928 ± 0.057 per rating point (10 rating = 1%, 9.279 per %), spell_haste=7.675 ± 0.785, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 38.7 spell_power points (7.97 DPS) | yes | Dreamweave Circlet (10041, -1.62 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.88 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.41 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcane Crystal Pendant (20037, +0.00 DPS) [quest]; Horizon Choker (13085, -0.22 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.41 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 30.7 spell_power points (6.33 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.84 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.08 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.9 spell_power points (4.10 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.42 DPS) [dungeon]; Runecloth Cloak (13860, -0.62 DPS) [crafted]; Big Voodoo Cloak (8216, -1.24 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 38.7 spell_power points (7.97 DPS) | yes | Robe of the Magi (1716, -2.22 DPS) [world_drop]; Runecloth Tunic (13857, -2.24 DPS) [crafted]; Runecloth Robe (13858, -2.26 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (236.0 DPS) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Bloodband Bracers (11469, -0.01 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.43 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 35.1 spell_power points (7.22 DPS) | yes | Dreamweave Gloves (10019, -2.71 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.72 DPS) [vendor]; Raider Handwraps (272098, -3.82 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 32.1 spell_power points (6.61 DPS) | yes | Satyrmane Sash (17755, -1.70 DPS) [dungeon]; Deathmage Sash (10771, -2.13 DPS) [dungeon]; Dawnspire Cord (12466, -5.02 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.9 spell_power points (6.77 DPS) | yes | Red Mageweave Pants (10009, -1.45 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.48 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -4.54 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 26.2 spell_power points (5.39 DPS) | yes | Earthen Silk Slippers (254013, -0.44 DPS) [crafted]; Gilded Sandals (254107, -1.29 DPS) [crafted]; Southsea Mojo Boots (20641, -1.50 DPS) [quest] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.9 spell_power points (3.27 DPS) | yes | Brainlash (6440, -0.23 DPS) [dungeon]; Band of the Unicorn (7553, -0.60 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.80 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.9 spell_power points (3.07 DPS) | yes | Brainlash (6440, -0.03 DPS) [dungeon]; Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.60 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -4.00 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.00 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.64 DPS) [dungeon]; Hanzo Sword (8190, -29.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 255.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.47 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.10 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 443, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 203005000000000000-11302300000000000-2535101301000300250)

Set DPS (verified): 356.8. Weights run: 0.8s. Verify run: 2.9s. 1073 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.814 ± 0.056, crit=0.379 ± 0.014 per rating point (14 rating = 1%, 5.307 per %), hit=1.281 ± 0.072 per rating point (10 rating = 1%, 12.805 per %), spell_haste=7.471 ± 0.981, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 46.6 spell_power points (10.02 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.22 DPS) [pvp]; Crimson Felt Hat (18727, -6.59 DPS, sim-verified) [dungeon] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (356.8 DPS) | yes | Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Amulet of the Dawn (22657, +0.00 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.5 spell_power points (9.56 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.57 DPS) [pvp]; Burial Shawl (18681, -2.47 DPS) [dungeon]; Darkspear Shoulderpads (272103, -4.25 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.3 spell_power points (7.59 DPS) | yes | Crystalline Threaded Cape (20697, -2.59 DPS) [world]; Hide of the Wild (18510, -2.83 DPS) [crafted]; Shroud of Arcane Mastery (22330, -2.91 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.1 spell_power points (11.83 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.63 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.21 DPS) [pvp]; Robe of Everlasting Night (18385, -3.76 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.5 spell_power points (6.12 DPS) | yes | Sublime Wristguards (18497, -1.80 DPS) [dungeon]; Runecloth Cuffs (254123, -2.01 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.80 DPS) [pvp] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, -5.41 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 55.7 spell_power points (11.97 DPS) | yes | Belt of the Archmage (18405, -3.73 DPS) [crafted]; Ban'thok Sash (11662, -4.71 DPS) [dungeon]; Magician's Cord (272393, -4.87 DPS, sim-verified) [vendor] |
| legs | Sorcerer's Leggings (226933) | Anthion's Parting Words [quest] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -0.29 DPS) [rep] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.64 DPS) [dungeon]; Sorcerer's Boots (22064, -5.35 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -4.15 DPS) [dungeon]; Channeler's Ring (272406, -5.14 DPS) [vendor]; Naglering (11669, -18.03 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.46 DPS) [dungeon]; Channeler's Ring (272406, -1.45 DPS) [vendor]; Naglering (11669, -11.02 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+15.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.63 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -32.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 363.4 spell_power points (78.05 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.84 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.36 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.07 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Jewel of Kajaro; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sorcerer's Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Burst of Knowledge; ranged: Torch of Light

No-known-source sample (15 of 1073, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 203005000000000000-13102000000000000-2555100321000300250)

Set DPS (verified): 611.3. Weights run: 0.9s. Verify run: 2.3s. 1073 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.261 ± 0.019, crit=0.410 ± 0.013 per rating point (14 rating = 1%, 5.742 per %), hit=1.120 ± 0.047 per rating point (10 rating = 1%, 11.199 per %), spell_haste=7.612 ± 0.330, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 33.3 spell_power points (14.22 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.78 DPS) [pvp] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, -0.12 DPS) [quest]; Chains of the Lich (23125, -0.12 DPS) [dungeon]; Jewel of Kajaro (19601, -9.01 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.7 spell_power points (15.67 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.31 DPS) [pvp]; Argent Shoulders (19059, -4.98 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.08 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.3 spell_power points (12.52 DPS) | yes | Crystalline Threaded Cape (20697, -3.52 DPS) [world]; Amplifying Cloak (18350, -4.82 DPS) [dungeon]; Hide of the Wild (18510, -5.42 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.9 spell_power points (20.89 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.43 DPS) [pvp]; Robe of Everlasting Night (18385, -5.68 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Tunic (227108, -7.56 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.1 spell_power points (10.30 DPS) | yes | Sublime Wristguards (18497, -4.05 DPS) [dungeon]; Runecloth Cuffs (254123, -4.48 DPS) [crafted]; Sorcerer's Bindings (226929, -6.19 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.3 spell_power points (12.10 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 41.9 spell_power points (17.93 DPS) | yes | Belt of the Archmage (18405, -6.49 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.78 DPS) [dungeon]; Magician's Cord (272393, -7.13 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.1 spell_power points (19.70 DPS) | yes | Marshal's Silk Leggings (231605, -2.19 DPS) [pvp]; Skyshroud Leggings (13170, -4.28 DPS) [dungeon]; Sorcerer's Leggings (226933, -5.01 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 25.2 spell_power points (10.76 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.50 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Elemental Focus Band (20682, -7.13 DPS) [world]; Blessed Band of Light (272407, -8.70 DPS) [vendor]; Naglering (11669, -22.20 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Elemental Focus Band (20682, -1.00 DPS) [world]; Blessed Band of Light (272407, -2.56 DPS) [vendor]; Naglering (11669, -17.90 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Briarwood Reed (12930, -20.95 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (611.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -0.85 DPS) [dungeon] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.48 DPS) [world]; Teebu's Blazing Longsword (1728, -39.66 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 185.6 spell_power points (79.32 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.77 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.55 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.85 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Rune Band of Wizardry; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1073, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 29.1. Weights run: 0.7s. Verify run: 0.6s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.315 ± 0.007, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.078 per %), hit=0.256 ± 0.007 per rating point (10 rating = 1%, 2.558 per %), spell_haste=0.429 ± 0.105, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.63 DPS) | yes | Shadow Goggles (4373, -1.00 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.8 spell_power points (0.82 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.31 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.42 DPS) | yes | Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.11 DPS) [rep]; Black Whelp Cloak (7283, -0.26 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.6 spell_power points (0.69 DPS) | yes | Green Woolen Vest (2582, -0.27 DPS) [crafted]; Bloody Apron (6226, -0.27 DPS) [dungeon]; Gray Woolen Robe (2585, -0.65 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.32 DPS) | yes | Owlbeard Bracers (16981, -0.14 DPS) [quest]; Mindthrust Bracers (1974, -0.15 DPS) [dungeon]; Featherbead Bracers (15452, -0.15 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.74 DPS) | yes | Gnoll Casting Gloves (892, -0.21 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.22 DPS) [crafted]; Apothecary Gloves (10919, -0.32 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.3 spell_power points (0.55 DPS) | yes | Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Keller's Girdle (2911, -0.29 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.42 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.5 spell_power points (1.21 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.48 DPS) [dungeon]; Rumpled Kilt (274741, -0.69 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.3 spell_power points (0.87 DPS) | yes | Red Woolen Boots (4313, -0.45 DPS) [crafted]; Pristine Boots (253889, -0.45 DPS) [crafted]; Feather Padded Treads (285345, -0.61 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.53 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Loop of Sacrifice (281673, -0.36 DPS) [quest]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.32 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.15 DPS) [quest]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.1 spell_power points (0.33 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.07 DPS) [world]; Lesser Staff of the Spire (1300, -0.13 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 215.1 spell_power points (22.63 DPS) | yes | Skycaller (12984, -1.03 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.34 DPS) [dungeon]; Sizzle Stick (8071, -4.45 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 000000000000000000-00000000000000000-2535101301000000000)

Set DPS (verified): 47.2. Weights run: 0.9s. Verify run: 0.7s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.154 ± 0.004, crit=0.099 ± 0.003 per rating point (14 rating = 1%, 1.390 per %), hit=0.248 ± 0.006 per rating point (10 rating = 1%, 2.483 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.27 DPS) | yes | Embalmed Shroud (7691, -0.35 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.35 DPS) [crafted]; Silk Headband (7050, -1.83 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (0.91 DPS) | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.84 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.01 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (1.19 DPS) | yes | Invoker's Mantle (215365, -0.30 DPS) [crafted]; Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Chestnut Mantle (17695, -0.97 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.58 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Battle Healer's Cloak (19529, -0.12 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.09 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.13 DPS) [dungeon]; Tree Bark Jacket (1486, -1.04 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Owlbeard Bracers (16981, -0.88 DPS) [quest]; Glowing Magical Bracelets (13106, -0.89 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.26 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.81 DPS) | yes | Jutebraid Gloves (10654, -0.03 DPS) [quest]; Gnoll Casting Gloves (892, -0.12 DPS) [world]; Truefaith Gloves (7049, -0.18 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (1.32 DPS) | yes | Warsong Sash (16975, -0.05 DPS) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.40 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.38 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.45 DPS) [crafted]; Silk-threaded Trousers (1929, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (0.93 DPS) | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Boots of the Enchanter (4325, -0.35 DPS) [crafted]; Spidersilk Boots (4320, -1.27 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.81 DPS) | yes | Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world]; Sacred Band (6669, -0.58 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Sacred Band (6669, -0.46 DPS) [quest]; Sludge-Stained Band (286535, -1.04 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 23.8 spell_power points (2.73 DPS) | yes | Glimmering Staff (249392, -2.54 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -2.56 DPS) [quest]; Scorn's Focal Dagger (23168, -5.20 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Starfaller (13063, -0.62 DPS) [world_drop]; Unstable Power Core (279847, -1.02 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.98 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 88.0. Weights run: 0.8s. Verify run: 0.8s. 323 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.662 ± 0.027, crit=0.167 ± 0.007 per rating point (14 rating = 1%, 2.332 per %), hit=0.421 ± 0.029 per rating point (10 rating = 1%, 4.205 per %), spell_haste=not significant (1.417 ± 0.447), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.01 DPS) | yes | Augural Shroud (2620, -0.48 DPS) [world]; Electromagnetic Gigaflux Reactivator (9492, -1.10 DPS) [dungeon]; Living Cowl (5608, -1.15 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.3 spell_power points (1.76 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.81 DPS) [quest]; Darkspear Warding Pendant (272074, -1.10 DPS) [vendor]; Scorn's Icy Choker (23169, -1.36 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 spell_power points (2.24 DPS) | yes | Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest]; Green Silken Shoulders (7057, -1.30 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.0 spell_power points (2.14 DPS) | yes | Guardian Cloak (5965, -0.81 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.10 DPS) [vendor]; Long Silken Cloak (4326, -2.01 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 spell_power points (3.72 DPS) | yes | Dreamweave Vest (10021, -0.29 DPS) [crafted]; Robe of Power (7054, -0.58 DPS) [crafted]; Zealot's Robe (17043, -0.62 DPS) [quest] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.3 spell_power points (1.33 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.04 DPS) [dungeon]; Condor Bracers (15864, -0.33 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (2.96 DPS) | yes | Red Mageweave Gloves (10018, -0.43 DPS) [crafted]; Black Mageweave Gloves (10003, -0.81 DPS) [crafted]; Gilded Handwraps (254021, -1.15 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.9 spell_power points (2.43 DPS) | yes | Defiler's Cloth Girdle (20166, -0.04 DPS) [rep]; Gilded Cord (254037, -0.52 DPS) [crafted]; Star Belt (4329, -0.56 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.9 spell_power points (3.15 DPS) | yes | Abomination Skin Leggings (23173, -1.10 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.36 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.44 DPS) | yes | Gilded Slippers (254001, -0.83 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.96 DPS) [dungeon]; Spidersilk Boots (4320, -2.06 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (1.91 DPS) | yes | Reedknot Ring (9622, -0.90 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.05 DPS) [vendor]; Black Widow Band (6199, -1.24 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.29 DPS) | yes | Sea Giant's Toe Ring (274746, -0.43 DPS) [vendor]; Black Widow Band (6199, -0.63 DPS) [world]; Reedknot Ring (9622, -1.54 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (88.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Wind Spirit Staff (6689, -3.29 DPS) [dungeon] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.0 spell_power points (2.15 DPS) | yes | Orb of the Forgotten Seer (7685, -0.14 DPS) [dungeon]; Thrash's Trash (276204, -0.14 DPS) [vendor]; Orb of Mystic Insight (249394, -0.58 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 278.7 spell_power points (39.95 DPS) | yes | Nether Force Wand (11263, -2.17 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.47 DPS) [quest]; Ragefire Wand (7513, -2.52 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 203005000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 225.7. Weights run: 0.8s. Verify run: 1.0s. 416 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.986 ± 0.043, crit=0.270 ± 0.011 per rating point (14 rating = 1%, 3.775 per %), hit=0.928 ± 0.057 per rating point (10 rating = 1%, 9.279 per %), spell_haste=7.675 ± 0.785, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 38.7 spell_power points (7.97 DPS) | yes | Dreamweave Circlet (10041, -1.62 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.88 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.41 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.9 spell_power points (4.51 DPS) | yes | Horizon Choker (13085, -1.67 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.85 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -3.53 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 30.7 spell_power points (6.33 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.84 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.08 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.9 spell_power points (4.30 DPS) | yes | Spritecaster Cape (11623, -0.20 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.62 DPS) [dungeon]; Runecloth Cloak (13860, -0.82 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 38.7 spell_power points (7.97 DPS) | yes | Runecloth Tunic (13857, -2.24 DPS) [crafted]; Runecloth Robe (13858, -2.26 DPS) [crafted]; Robe of the Magi (1716, -4.52 DPS, sim-verified) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 35.1 spell_power points (7.22 DPS) | yes | Raider Handwraps (272098, -1.32 DPS) [vendor]; Dreamweave Gloves (10019, -2.71 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.72 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 32.1 spell_power points (6.61 DPS) | yes | Satyrmane Sash (17755, -1.70 DPS) [dungeon]; Deathmage Sash (10771, -2.13 DPS) [dungeon]; Dawnspire Cord (12466, -5.39 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | sim-verified (225.7 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Red Mageweave Pants (10009, -0.37 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.40 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.2 spell_power points (5.39 DPS) | yes | Earthen Silk Slippers (254013, -0.44 DPS) [crafted]; Gilded Sandals (254107, -1.29 DPS) [crafted]; Southsea Mojo Boots (20641, -1.50 DPS) [quest] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.9 spell_power points (3.27 DPS) | yes | Brainlash (6440, -0.23 DPS) [dungeon]; Band of the Unicorn (7553, -0.60 DPS) [world_drop]; Advisor's Ring (19519, -0.80 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.9 spell_power points (3.07 DPS) | yes | Brainlash (6440, -0.03 DPS) [dungeon]; Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Advisor's Ring (19519, -0.60 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.38 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Mark of the Chosen (17774, -3.85 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.00 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.64 DPS) [dungeon]; Hanzo Sword (8190, -30.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 255.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.47 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.10 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 203005000000000000-11302300000000000-2535101301000300250)

Set DPS (verified): 347.7. Weights run: 0.8s. Verify run: 2.9s. 1061 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.814 ± 0.056, crit=0.379 ± 0.014 per rating point (14 rating = 1%, 5.307 per %), hit=1.281 ± 0.072 per rating point (10 rating = 1%, 12.805 per %), spell_haste=7.471 ± 0.981, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 46.6 spell_power points (10.02 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.22 DPS) [pvp]; Crimson Felt Hat (18727, -5.77 DPS, sim-verified) [dungeon] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (347.7 DPS) | yes | Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Amulet of the Dawn (22657, +0.00 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.5 spell_power points (9.56 DPS) | yes | Warlord's Silk Amice (231594, -1.57 DPS) [pvp]; Burial Shawl (18681, -2.47 DPS) [dungeon]; Darkspear Shoulderpads (272103, -6.76 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.3 spell_power points (7.59 DPS) | yes | Hide of the Wild (18510, -2.83 DPS) [crafted]; Shroud of Arcane Mastery (22330, -2.91 DPS) [dungeon]; Crystalline Threaded Cape (20697, -6.32 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.1 spell_power points (11.83 DPS) | yes | Warlord's Silk Raiment (231596, -0.63 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.21 DPS) [pvp]; Robe of Everlasting Night (18385, -3.76 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.5 spell_power points (6.12 DPS) | yes | Sublime Wristguards (18497, -1.80 DPS) [dungeon]; Runecloth Cuffs (254123, -2.01 DPS) [crafted]; General's Silk Cuffs (16538, -2.80 DPS) [pvp] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 55.7 spell_power points (11.97 DPS) | yes | Belt of the Archmage (18405, -3.73 DPS) [crafted]; Ban'thok Sash (11662, -4.71 DPS) [dungeon]; Magician's Cord (272393, -7.78 DPS, sim-verified) [vendor] |
| legs | Sorcerer's Leggings (226933) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.29 DPS) [rep]; Sentinel's Silk Leggings (237815, -6.34 DPS, sim-verified) [vendor] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.64 DPS) [dungeon]; Sorcerer's Boots (22064, -5.36 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -4.15 DPS) [dungeon]; Channeler's Ring (272406, -5.14 DPS) [vendor]; Naglering (11669, -19.88 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.46 DPS) [dungeon]; Channeler's Ring (272406, -1.45 DPS) [vendor]; Naglering (11669, -9.11 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+15.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.50 DPS) [vendor]; Serenity Field (272439, -3.22 DPS) [vendor]; Burst of Knowledge (11832, -3.65 DPS) [dungeon] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.63 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -31.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 363.4 spell_power points (78.05 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.84 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.36 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.07 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Jewel of Kajaro; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sorcerer's Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; ranged: Torch of Light

No-known-source sample (15 of 1061, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (troll, 203005000000000000-13102000000000000-2555100321000300250)

Set DPS (verified): 604.8. Weights run: 0.9s. Verify run: 2.3s. 1061 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.261 ± 0.019, crit=0.410 ± 0.013 per rating point (14 rating = 1%, 5.742 per %), hit=1.120 ± 0.047 per rating point (10 rating = 1%, 11.199 per %), spell_haste=7.612 ± 0.330, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 33.3 spell_power points (14.22 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.78 DPS) [pvp] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, -0.12 DPS) [quest]; Chains of the Lich (23125, -0.12 DPS) [dungeon]; Jewel of Kajaro (19601, -7.27 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.7 spell_power points (15.67 DPS) | yes | Warlord's Silk Amice (231594, -3.31 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.81 DPS, sim-verified) [crafted]; Argent Shoulders (19059, -4.98 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.3 spell_power points (12.52 DPS) | yes | Crystalline Threaded Cape (20697, -3.52 DPS) [world]; Amplifying Cloak (18350, -4.82 DPS) [dungeon]; Hide of the Wild (18510, -5.42 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.9 spell_power points (20.89 DPS) | yes | Warlord's Silk Raiment (231596, -2.43 DPS) [pvp]; Robe of Everlasting Night (18385, -5.37 DPS, sim-verified) [dungeon]; Legionnaire's Silk Tunic (227106, -7.56 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.1 spell_power points (10.30 DPS) | yes | Sublime Wristguards (18497, -4.05 DPS) [dungeon]; Runecloth Cuffs (254123, -4.48 DPS) [crafted]; Sorcerer's Bindings (226929, -6.19 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.3 spell_power points (12.10 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 41.9 spell_power points (17.93 DPS) | yes | Belt of the Archmage (18405, -5.23 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.78 DPS) [dungeon]; Magician's Cord (272393, -7.13 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.1 spell_power points (19.70 DPS) | yes | General's Silk Trousers (231595, -2.19 DPS) [pvp]; Skyshroud Leggings (13170, -4.28 DPS) [dungeon]; Sorcerer's Leggings (226933, -5.01 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 25.2 spell_power points (10.76 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.50 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.13 DPS) [dungeon]; Blessed Band of Light (272407, -8.70 DPS) [vendor]; Naglering (11669, -20.73 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -1.57 DPS) [vendor]; Maiden's Circle (13001, -2.16 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Briarwood Reed (12930, -20.64 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (604.8 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -0.85 DPS) [dungeon] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.48 DPS) [world]; Teebu's Blazing Longsword (1728, -40.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 185.6 spell_power points (79.32 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.77 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.55 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.85 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1061, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

