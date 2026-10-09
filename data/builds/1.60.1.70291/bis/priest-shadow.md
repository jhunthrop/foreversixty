# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 29.3. Weights run: 1.0s. Verify run: 0.8s. 151 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.306 ± 0.008, crit=0.110 ± 0.003 per rating point (14 rating = 1%, 1.534 per %), hit=0.411 ± 0.012 per rating point (10 rating = 1%, 4.110 per %), spell_haste=not significant (-0.512 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Shadow Goggles (4373, -1.16 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.8 spell_power points (0.71 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.14 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.34 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.37 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.10 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.60 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -0.52 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.5 spell_power points (0.14 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.05 DPS) [quest]; Repurposed Hair Band (281256, -0.08 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.15 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.48 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.25 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.37 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.5 spell_power points (1.05 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.75 DPS) | yes | Feather Padded Treads (285345, -0.22 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.39 DPS) [crafted]; Pristine Boots (253889, -0.39 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.51 DPS) | yes | Sludge-Stained Band (286535, -0.24 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.46 DPS) | yes | Sludge-Stained Band (286535, -0.20 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.1 spell_power points (0.28 DPS) | yes | Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.11 DPS) [world]; Staff of Westfall (2042, -0.14 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 246.3 spell_power points (22.59 DPS) | yes | Skycaller (12984, -1.03 DPS) [world_drop]; Firebelcher (5243, -2.30 DPS) [dungeon]; Deepblaze (279896, -4.06 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 151, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 63.4. Weights run: 1.0s. Verify run: 0.8s. 262 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.465 ± 0.009, crit=0.044 ± 0.002 per rating point (14 rating = 1%, 0.622 per %), hit=0.231 ± 0.009 per rating point (10 rating = 1%, 2.309 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.24 DPS) | yes | Silk Headband (7050, -0.22 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Enchanter's Cowl (4322, -1.08 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.10 DPS) | yes | Crystal Starfire Medallion (5003, -0.89 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.89 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.48 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (1.48 DPS) | yes | Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (0.67 DPS) | yes | Hillman's Cloak (3719, -0.11 DPS) [crafted]; Repairman's Cape (9605, -0.13 DPS) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.0 spell_power points (1.69 DPS) | yes | Tree Bark Jacket (1486, -0.23 DPS) [dungeon]; Beguiler Robes (7728, -0.26 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.57 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.01 DPS) | yes | Nightsky Wristbands (6407, -0.70 DPS) [world_drop]; Stonecloth Bindings (14416, -0.75 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.94 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 9.1 spell_power points (1.02 DPS) | yes | Serpent Gloves (5970, -0.24 DPS) [dungeon]; Truefaith Gloves (7049, -0.31 DPS) [crafted]; Shilly Mitts (9609, -0.95 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.4 spell_power points (1.39 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Invoker's Cord (215366, -0.35 DPS) [crafted]; Crimson Silk Belt (7055, -0.35 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.7 spell_power points (1.43 DPS) | yes | Pristine Leggings (253987, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.44 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.50 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.15 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Nimbus Boots (6998, -0.48 DPS) [quest]; Spidersilk Boots (4320, -2.02 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.79 DPS) | yes | Minor Channeling Ring (1449, -0.12 DPS) [quest]; Black Widow Band (6199, -0.42 DPS) [world]; Snake Hoop (6750, -0.42 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.67 DPS) | yes | Black Widow Band (6199, -0.31 DPS) [world]; Snake Hoop (6750, -0.31 DPS) [quest]; Minor Channeling Ring (1449, -0.77 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (63.4 DPS) | yes | Royal Diplomatic Scepter (9457, -0.04 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.84 DPS) [dungeon]; Hardened Root Staff (1317, -2.01 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 298.3 spell_power points (33.55 DPS) | yes | Starfaller (13063, -0.48 DPS) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 98.7. Weights run: 1.0s. Verify run: 0.9s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.832 ± 0.014, crit=0.040 ± 0.002 per rating point (14 rating = 1%, 0.560 per %), hit=0.413 ± 0.018 per rating point (10 rating = 1%, 4.130 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.86 DPS) | yes | Augural Shroud (2620, -0.23 DPS) [world]; Corpseshroud (10574, -0.71 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.86 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.7 spell_power points (1.86 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.73 DPS) [quest]; Darkspear Warding Pendant (272074, -1.06 DPS) [vendor]; Scorn's Icy Choker (23169, -1.08 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.8 spell_power points (2.42 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.34 DPS) [quest]; Green Silken Shoulders (7057, -1.01 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.5 spell_power points (2.24 DPS) | yes | Guardian Cloak (5965, -0.86 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.00 DPS) [vendor]; Long Silken Cloak (4326, -2.00 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.0 spell_power points (3.67 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.41 DPS) [crafted]; Elemental Raiment (9434, -0.81 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.22 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.21 DPS) [world_drop]; Condor Bracers (15864, -0.27 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (2.90 DPS) | yes | Red Mageweave Gloves (10018, -0.63 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.86 DPS) [crafted]; Stormcloth Gloves (10011, -1.00 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.5 spell_power points (2.65 DPS) | yes | Gilded Cord (254037, -0.66 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.67 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.0 spell_power points (3.26 DPS) | yes | Crimson Silk Pantaloons (7062, -0.90 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.40 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.26 DPS) | yes | Gilded Slippers (254001, -1.02 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.68 DPS) [dungeon]; Spidersilk Boots (4320, -1.86 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (1.93 DPS) | yes | Ring of Forlorn Spirits (2043, -0.84 DPS) [quest]; Reedknot Ring (9622, -0.97 DPS) [quest]; Minor Channeling Ring (1449, -1.02 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.22 DPS) | yes | Reedknot Ring (9622, -0.27 DPS) [quest]; Minor Channeling Ring (1449, -0.32 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.16 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (98.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.23 DPS) [dungeon]; Hand of Righteousness (7721, -2.67 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 14.3 spell_power points (1.95 DPS) | yes | Thrash's Trash (276204, -0.04 DPS) [vendor]; Orb of Mystic Insight (249394, -0.32 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.11 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 294.6 spell_power points (40.06 DPS) | yes | Umbral Wand (5216, -1.04 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.47 DPS) [dungeon]; Twisted Nether Wand (249144, -5.24 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 159.8. Weights run: 1.0s. Verify run: 1.4s. 441 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.484 ± 0.024, crit=0.052 ± 0.002 per rating point (14 rating = 1%, 0.727 per %), hit=0.631 ± 0.034 per rating point (10 rating = 1%, 6.305 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.7 spell_power points (6.17 DPS) | yes | Soulcatcher Halo (10630, -1.47 DPS) [dungeon]; Dreamweave Circlet (10041, -1.63 DPS) [crafted]; Chief Architect's Monocle (11839, -1.73 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.8 spell_power points (2.63 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Scorn's Icy Choker (23169, -0.62 DPS) [dungeon]; Mindburst Medallion (11196, -0.74 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.7 spell_power points (5.03 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.32 DPS) [crafted]; Inquisitor's Shawl (19507, -1.70 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.9 spell_power points (2.90 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.07 DPS) [dungeon]; Runecloth Cloak (13860, -0.26 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.27 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.7 spell_power points (6.17 DPS) | yes | Robes of Insight (940, -1.18 DPS, sim-verified) [world_drop]; Runecloth Robe (13858, -1.58 DPS) [crafted]; Runecloth Tunic (13857, -1.95 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (159.8 DPS) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Forgotten Wraps (9433, -0.07 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.07 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.6 spell_power points (5.03 DPS) | yes | Virtuous Hands (226958, -0.17 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.69 DPS) [vendor]; Red Mageweave Gloves (10018, -1.75 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 34.6 spell_power points (4.39 DPS) | yes | Dawnspire Cord (12466, -0.05 DPS) [dungeon]; Deathmage Sash (10771, -0.68 DPS) [dungeon]; Satyrmane Sash (17755, -0.73 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.8 spell_power points (4.80 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.93 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.03 DPS) [dungeon]; Red Mageweave Pants (10009, -2.84 DPS, sim-verified) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 27.7 spell_power points (3.51 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Southsea Mojo Boots (20641, -0.42 DPS) [quest]; Earthen Silk Slippers (254013, -0.46 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 22.3 spell_power points (2.82 DPS) | yes | Mindseye Circle (10634, -0.56 DPS) [dungeon]; Philanthropist's Ring (281635, -0.61 DPS) [quest]; Woodseed Hoop (17768, -1.13 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.4 spell_power points (2.46 DPS) | yes | Philanthropist's Ring (281635, -0.25 DPS) [quest]; Woodseed Hoop (17768, -0.76 DPS) [quest]; Mindseye Circle (10634, -1.03 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, -2.51 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.31 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.45 DPS) [dungeon]; Barman Shanker (12791, -21.42 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 414.2 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 441, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 342.1. Weights run: 1.1s. Verify run: 3.6s. 1121 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.132 ± 0.028, crit=0.111 ± 0.004 per rating point (14 rating = 1%, 1.551 per %), hit=1.031 ± 0.040 per rating point (10 rating = 1%, 10.314 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 47.9 spell_power points (7.45 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Magister's Crown (16686, -9.72 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.7 spell_power points (4.62 DPS) | yes | Diana's Pearl Necklace (22403, -0.21 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.49 DPS) [quest]; Archlight Talisman (15856, -0.99 DPS) [quest] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | sim-verified (342.1 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.05 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -5.95 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.4 spell_power points (5.50 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -1.68 DPS) [world]; Shroud of Arcane Mastery (22330, -1.96 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 46.4 spell_power points (7.21 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -9.50 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.1 spell_power points (4.83 DPS) | yes | Sublime Wristguards (18497, -1.20 DPS) [dungeon]; Runecloth Cuffs (254123, -1.36 DPS) [crafted]; Marshal's Satin Bracers (17606, -1.48 DPS) [pvp] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | sim-verified (342.1 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Raider Handwraps (272097, -5.67 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 60.2 spell_power points (9.36 DPS) | yes | Ban'thok Sash (11662, -3.95 DPS) [dungeon]; Clutch of Andros (13956, -4.41 DPS) [dungeon]; Belt of the Archmage (18405, -6.15 DPS, sim-verified) [crafted] |
| legs | Virtuous Leggings (226956) | Mokvar [vendor] | sim-verified (342.1 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (22752, -4.76 DPS, sim-verified) [rep] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 40.1 spell_power points (6.23 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -8.55 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (342.1 DPS) | yes | Rune Band of Wizardry (22339, -3.45 DPS) [dungeon]; Channeler's Ring (272406, -3.47 DPS) [vendor]; Naglering (11669, -13.88 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (342.1 DPS) | yes | Rune Band of Wizardry (22339, -1.26 DPS) [dungeon]; Channeler's Ring (272406, -1.29 DPS) [vendor]; Naglering (11669, -10.27 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (342.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (342.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (342.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -1.00 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 500.0 spell_power points (77.69 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.13 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.65 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.00 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Virtuous Hands; waist: Knowledge of the Timbermaw; legs: Virtuous Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Second Wind; trinket2: Burst of Knowledge; ranged: Torch of Light

No-known-source sample (15 of 1121, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 325001031300000000-00000000000000000-523120501201300251)

Set DPS (verified): 608.8. Weights run: 1.2s. Verify run: 3.0s. 1121 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.594 ± 0.014, crit=0.039 ± 0.003 per rating point (14 rating = 1%, 0.547 per %), hit=0.716 ± 0.034 per rating point (10 rating = 1%, 7.161 per %), spell_haste=-3.603 ± 0.264, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 37.2 spell_power points (14.03 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 22.7 spell_power points (8.56 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Orb of the Darkmoon (19426, -0.27 DPS) [quest]; Diana's Pearl Necklace (22403, -0.68 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.5 spell_power points (13.73 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Virtuous Epaulets (226955, -1.19 DPS) [vendor]; Burial Shawl (18681, -2.62 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.9 spell_power points (10.52 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Hide of the Wild (18510, -3.00 DPS) [crafted]; Amplifying Cloak (18350, -3.73 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 36.7 spell_power points (13.82 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Robe of Everlasting Night (18385, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.7 spell_power points (10.08 DPS) | yes | Sublime Wristguards (18497, -3.32 DPS) [dungeon]; Runecloth Cuffs (254123, -3.70 DPS) [crafted]; Virtuous Wraps (226953, -4.60 DPS) [vendor] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 29.6 spell_power points (11.14 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Virtuous Hands (226958, -5.97 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 45.2 spell_power points (17.04 DPS) | yes | Belt of the Archmage (18405, -4.74 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -7.36 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -8.02 DPS) [rep] |
| legs | Sentinel's Silk Leggings (22752) | Silverwing Sentinels [rep] | 39.3 spell_power points (14.80 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Skyshroud Leggings (13170, -0.20 DPS) [dungeon] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 31.5 spell_power points (11.87 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Dragonrider Boots (18102, +0.00 DPS, sim-verified) [dungeon]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.53 DPS) [dungeon]; Maiden's Circle (13001, -6.91 DPS) [world_drop]; Naglering (11669, -14.08 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -2.02 DPS) [dungeon]; Maiden's Circle (13001, -2.40 DPS) [world_drop]; Naglering (11669, -8.71 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+19.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.64 DPS) [vendor]; Serenity Field (272439, -4.20 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -6.41 DPS) [dungeon] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.20 DPS) [dungeon]; Hand of Edward the Odd (2243, -28.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (608.8 DPS) | yes | Bonecreeper Stylus (13938, -0.33 DPS) [dungeon]; Sparkling Crystal Wand (20672, -1.77 DPS) [world]; Torch of Light (279246, -8.81 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1121, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 34.6. Weights run: 1.0s. Verify run: 0.8s. 141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.306 ± 0.008, crit=0.110 ± 0.003 per rating point (14 rating = 1%, 1.534 per %), hit=0.411 ± 0.012 per rating point (10 rating = 1%, 4.110 per %), spell_haste=not significant (-0.512 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Shadow Goggles (4373, -1.32 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.8 spell_power points (0.71 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.34 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.43 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.37 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.61 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.60 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -0.91 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.28 DPS) | yes | Mindthrust Bracers (1974, -0.13 DPS) [dungeon]; Featherbead Bracers (15452, -0.13 DPS) [quest]; Owlbeard Bracers (16981, -0.43 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Apothecary Gloves (10919, -0.28 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.48 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.25 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.61 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.5 spell_power points (1.05 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.75 DPS) | yes | Red Woolen Boots (4313, -0.39 DPS) [crafted]; Pristine Boots (253889, -0.39 DPS) [crafted]; Feather Padded Treads (285345, -0.61 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Loop of Sacrifice (281673, -0.32 DPS) [quest]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.1 spell_power points (0.28 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.11 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 246.3 spell_power points (22.59 DPS) | yes | Skycaller (12984, -2.25 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.30 DPS) [dungeon]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 61.8. Weights run: 1.0s. Verify run: 0.8s. 249 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.465 ± 0.009, crit=0.044 ± 0.002 per rating point (14 rating = 1%, 0.622 per %), hit=0.231 ± 0.009 per rating point (10 rating = 1%, 2.309 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.24 DPS) | yes | Silk Headband (7050, -0.22 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Enchanter's Cowl (4322, -0.90 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.10 DPS) | yes | Darkspear Warding Pendant (272075, -0.85 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.89 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.89 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (1.48 DPS) | yes | Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Mantle of Woe (7750, -0.29 DPS) [quest]; Fairywing Mantle (9536, -0.34 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.56 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.0 spell_power points (1.69 DPS) | yes | Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon]; Tree Bark Jacket (1486, -0.23 DPS) [dungeon]; Beguiler Robes (7728, -0.26 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.01 DPS) | yes | Glowing Magical Bracelets (13106, -0.59 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.67 DPS) [quest]; Nightsky Wristbands (6407, -0.70 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.3 spell_power points (0.94 DPS) | yes | Truefaith Gloves (7049, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.26 DPS) [world]; Serpent Gloves (5970, -0.34 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.4 spell_power points (1.39 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Warsong Sash (16975, -0.24 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.35 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.7 spell_power points (1.43 DPS) | yes | Pristine Leggings (253987, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.44 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.75 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.15 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Boots of the Enchanter (4325, -0.59 DPS) [crafted]; Spidersilk Boots (4320, -1.17 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.79 DPS) | yes | Black Widow Band (6199, -0.42 DPS) [world]; Snake Hoop (6750, -0.42 DPS) [quest]; Sludge-Stained Band (286535, -0.45 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.67 DPS) | yes | Snake Hoop (6750, -0.31 DPS) [quest]; Sludge-Stained Band (286535, -0.34 DPS) [world]; Black Widow Band (6199, -0.74 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 25.3 spell_power points (2.85 DPS) | yes | Royal Diplomatic Scepter (9457, -0.24 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -1.84 DPS) [dungeon]; Glimmering Staff (249392, -2.27 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (61.8 DPS) | yes | Starfaller (13063, -0.48 DPS) [world_drop]; Unstable Power Core (279847, -2.77 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.99 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 92.7. Weights run: 1.0s. Verify run: 1.0s. 326 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.832 ± 0.014, crit=0.040 ± 0.002 per rating point (14 rating = 1%, 0.560 per %), hit=0.413 ± 0.018 per rating point (10 rating = 1%, 4.130 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.86 DPS) | yes | Augural Shroud (2620, -0.23 DPS) [world]; Corpseshroud (10574, -0.71 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.86 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.7 spell_power points (1.86 DPS) | yes | Scorn's Icy Choker (23169, -0.57 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.73 DPS) [quest]; Darkspear Warding Pendant (272074, -1.06 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.8 spell_power points (2.42 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.34 DPS) [quest]; Green Silken Shoulders (7057, -0.63 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.5 spell_power points (2.24 DPS) | yes | Guardian Cloak (5965, -0.86 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.00 DPS) [vendor]; Long Silken Cloak (4326, -1.89 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.0 spell_power points (3.67 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.41 DPS) [crafted]; Zealot's Robe (17043, -0.50 DPS) [quest] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.7 spell_power points (1.45 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Windchaser Cuffs (14429, -0.43 DPS) [world_drop]; Spidertank Oilrag (9448, -0.60 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (2.90 DPS) | yes | Red Mageweave Gloves (10018, -0.27 DPS) [crafted]; Black Mageweave Gloves (10003, -0.86 DPS) [crafted]; Stormcloth Gloves (10011, -1.00 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.5 spell_power points (2.65 DPS) | yes | Gilded Cord (254037, -0.66 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.81 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.0 spell_power points (3.26 DPS) | yes | Crimson Silk Pantaloons (7062, -0.87 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.40 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.26 DPS) | yes | Gilded Slippers (254001, -1.52 DPS) [crafted]; Acidic Walkers (9454, -1.68 DPS) [dungeon]; Spidersilk Boots (4320, -1.86 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (1.93 DPS) | yes | Reedknot Ring (9622, -0.97 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.11 DPS) [vendor]; Black Widow Band (6199, -1.13 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.22 DPS) | yes | Sea Giant's Toe Ring (274746, -0.41 DPS) [vendor]; Black Widow Band (6199, -0.43 DPS) [world]; Reedknot Ring (9622, -1.13 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (92.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.23 DPS) [dungeon]; Skullbreaker (17039, -0.39 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.90 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Mystic Insight (249394, -0.27 DPS) [crafted]; Prophetic Cane (6803, -0.55 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 294.6 spell_power points (40.06 DPS) | yes | Umbral Wand (5216, -1.37 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.47 DPS) [dungeon]; Twisted Nether Wand (249144, -5.24 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 149.3. Weights run: 1.0s. Verify run: 1.3s. 421 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.484 ± 0.024, crit=0.052 ± 0.002 per rating point (14 rating = 1%, 0.727 per %), hit=0.631 ± 0.034 per rating point (10 rating = 1%, 6.305 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.7 spell_power points (6.17 DPS) | yes | Soulcatcher Halo (10630, -1.47 DPS) [dungeon]; Dreamweave Circlet (10041, -1.63 DPS) [crafted]; Chief Architect's Monocle (11839, -1.89 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.8 spell_power points (2.63 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Scorn's Icy Choker (23169, -0.62 DPS) [dungeon]; Mindburst Medallion (11196, -0.74 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.7 spell_power points (5.03 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.32 DPS) [crafted]; Inquisitor's Shawl (19507, -1.70 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 25.4 spell_power points (3.21 DPS) | yes | Spritecaster Cape (11623, -0.31 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.38 DPS) [dungeon]; Runecloth Cloak (13860, -0.57 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.7 spell_power points (6.17 DPS) | yes | Robes of Insight (940, -1.38 DPS, sim-verified) [world_drop]; Runecloth Robe (13858, -1.58 DPS) [crafted]; Runecloth Tunic (13857, -1.95 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.6 spell_power points (5.03 DPS) | yes | Virtuous Hands (226958, -0.17 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.69 DPS) [vendor]; Red Mageweave Gloves (10018, -1.75 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 34.6 spell_power points (4.39 DPS) | yes | Deathmage Sash (10771, -0.68 DPS) [dungeon]; Satyrmane Sash (17755, -0.73 DPS) [dungeon]; Dawnspire Cord (12466, -1.02 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.8 spell_power points (4.80 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.93 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.03 DPS) [dungeon]; Red Mageweave Pants (10009, -2.39 DPS, sim-verified) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 27.7 spell_power points (3.51 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Southsea Mojo Boots (20641, -0.42 DPS) [quest]; Earthen Silk Slippers (254013, -0.46 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (149.3 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, +0.00 DPS) [dungeon]; Woodseed Hoop (17768, -0.52 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.4 spell_power points (2.46 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Woodseed Hoop (17768, -0.76 DPS) [quest]; Mindseye Circle (10634, -1.26 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.31 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.45 DPS) [dungeon]; Blade of Eternal Darkness (17780, -17.44 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 414.2 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 342.6. Weights run: 1.1s. Verify run: 3.6s. 1112 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.132 ± 0.028, crit=0.111 ± 0.004 per rating point (14 rating = 1%, 1.551 per %), hit=1.031 ± 0.040 per rating point (10 rating = 1%, 10.314 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 47.9 spell_power points (7.45 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -9.90 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.7 spell_power points (4.62 DPS) | yes | Diana's Pearl Necklace (22403, -0.21 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.49 DPS) [quest]; Archlight Talisman (15856, -0.99 DPS) [quest] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | sim-verified (342.6 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.05 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -6.05 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.4 spell_power points (5.50 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -1.68 DPS) [world]; Shroud of Arcane Mastery (22330, -1.96 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 46.4 spell_power points (7.21 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -9.86 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.1 spell_power points (4.83 DPS) | yes | Sublime Wristguards (18497, -1.20 DPS) [dungeon]; Runecloth Cuffs (254123, -1.36 DPS) [crafted]; General's Satin Bracers (17619, -1.48 DPS) [pvp] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | sim-verified (342.6 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Raider Handwraps (272097, -5.79 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 60.2 spell_power points (9.36 DPS) | yes | Ban'thok Sash (11662, -3.95 DPS) [dungeon]; Clutch of Andros (13956, -4.41 DPS) [dungeon]; Belt of the Archmage (18405, -5.78 DPS, sim-verified) [crafted] |
| legs | Virtuous Leggings (226956) | Mokvar [vendor] | sim-verified (342.6 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (237815, +0.00 DPS) [vendor]; Outrider's Silk Leggings (22747, -4.77 DPS, sim-verified) [rep] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 40.1 spell_power points (6.23 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -8.49 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (342.6 DPS) | yes | Rune Band of Wizardry (22339, -3.45 DPS) [dungeon]; Channeler's Ring (272406, -3.47 DPS) [vendor]; Naglering (11669, -13.25 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (342.6 DPS) | yes | Rune Band of Wizardry (22339, -1.26 DPS) [dungeon]; Channeler's Ring (272406, -1.29 DPS) [vendor]; Naglering (11669, -9.68 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (342.6 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (342.6 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -2.93 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (342.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -1.00 DPS) [dungeon]; Hand of Edward the Odd (2243, -21.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 500.0 spell_power points (77.69 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.13 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.65 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.00 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Virtuous Hands; waist: Knowledge of the Timbermaw; legs: Virtuous Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Second Wind; ranged: Torch of Light

No-known-source sample (15 of 1112, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 325001031300000000-00000000000000000-523120501201300251)

Set DPS (verified): 622.0. Weights run: 1.2s. Verify run: 2.9s. 1112 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.594 ± 0.014, crit=0.039 ± 0.003 per rating point (14 rating = 1%, 0.547 per %), hit=0.716 ± 0.034 per rating point (10 rating = 1%, 7.161 per %), spell_haste=-3.603 ± 0.264, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 37.2 spell_power points (14.03 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 22.7 spell_power points (8.56 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Orb of the Darkmoon (19426, -0.27 DPS) [quest]; Diana's Pearl Necklace (22403, -0.68 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.5 spell_power points (13.73 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Virtuous Epaulets (226955, -1.19 DPS) [vendor]; Burial Shawl (18681, -2.62 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.9 spell_power points (10.52 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Hide of the Wild (18510, -3.00 DPS) [crafted]; Amplifying Cloak (18350, -3.73 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 36.7 spell_power points (13.82 DPS) | yes | Robe of Everlasting Night (18385, +0.00 DPS, sim-verified) [dungeon]; Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.7 spell_power points (10.08 DPS) | yes | Sublime Wristguards (18497, -3.32 DPS) [dungeon]; Runecloth Cuffs (254123, -3.70 DPS) [crafted]; Virtuous Wraps (226953, -4.60 DPS) [vendor] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 29.6 spell_power points (11.14 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Virtuous Hands (226958, -5.59 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 45.2 spell_power points (17.04 DPS) | yes | Belt of the Archmage (18405, -4.52 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -7.36 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -8.02 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 39.3 spell_power points (14.80 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (237815, +0.00 DPS, sim-verified) [vendor]; Skyshroud Leggings (13170, -0.20 DPS) [dungeon] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 31.5 spell_power points (11.87 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; Dragonrider Boots (18102, +0.00 DPS, sim-verified) [dungeon]; General's Satin Treads (231610, +0.00 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.53 DPS) [dungeon]; Maiden's Circle (13001, -6.91 DPS) [world_drop]; Naglering (11669, -14.09 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -2.02 DPS) [dungeon]; Maiden's Circle (13001, -2.40 DPS) [world_drop]; Naglering (11669, -8.69 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+19.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.64 DPS) [vendor]; Serenity Field (272439, -4.19 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -6.41 DPS) [dungeon] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.20 DPS) [dungeon]; The Lobotomizer (19324, -28.69 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (622.0 DPS) | yes | Bonecreeper Stylus (13938, -0.33 DPS) [dungeon]; Sparkling Crystal Wand (20672, -1.77 DPS) [world]; Torch of Light (279246, -11.33 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Outrider's Silk Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1112, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

