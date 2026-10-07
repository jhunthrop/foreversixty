# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 29.0. Weights run: 0.9s. Verify run: 0.7s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.306 ± 0.008, crit=0.108 ± 0.003 per rating point (14 rating = 1%, 1.517 per %), hit=0.343 ± 0.002 per rating point (10 rating = 1%, 3.430 per %), spell_haste=not significant (-0.520 ± 0.135), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Shadow Goggles (4373, -1.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.7 spell_power points (0.71 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.14 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.34 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.37 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.10 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.60 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -0.52 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.5 spell_power points (0.14 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.05 DPS) [quest]; Repurposed Hair Band (281256, -0.08 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.48 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.25 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.37 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.4 spell_power points (1.04 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.75 DPS) | yes | Feather Padded Treads (285345, -0.22 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.39 DPS) [crafted]; Pristine Boots (253889, -0.39 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.51 DPS) | yes | Sludge-Stained Band (286535, -0.24 DPS) [world]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.46 DPS) | yes | Sludge-Stained Band (286535, -0.20 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.1 spell_power points (0.28 DPS) | yes | Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.11 DPS) [world]; Staff of Westfall (2042, -0.14 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 247.4 spell_power points (22.58 DPS) | yes | Skycaller (12984, -1.02 DPS) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-543110401200000000)

Set DPS (verified): 53.6. Weights run: 0.9s. Verify run: 0.7s. 244 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.481 ± 0.008, crit=0.045 ± 0.002 per rating point (14 rating = 1%, 0.632 per %), hit=0.196 ± 0.002 per rating point (10 rating = 1%, 1.961 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.21 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.22 DPS) [crafted]; Embalmed Shroud (7691, -0.33 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.9 spell_power points (1.09 DPS) | yes | Crystal Starfire Medallion (5003, -0.88 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.88 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.00 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.3 spell_power points (1.47 DPS) | yes | Death Speaker Mantle (6685, -0.22 DPS) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.55 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Prelacy Cape (7004, -0.11 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.3 spell_power points (1.68 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted]; Tree Bark Jacket (1486, -0.86 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.99 DPS) | yes | Nightsky Wristbands (6407, -0.67 DPS) [world_drop]; Stonecloth Bindings (14416, -0.73 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.01 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 9.3 spell_power points (1.03 DPS) | yes | Serpent Gloves (5970, -0.25 DPS) [dungeon]; Truefaith Gloves (7049, -0.31 DPS) [crafted]; Shilly Mitts (9609, -0.54 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.4 spell_power points (1.37 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Invoker's Cord (215366, -0.34 DPS) [crafted]; Crimson Silk Belt (7055, -0.34 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.9 spell_power points (1.42 DPS) | yes | Gaze Dreamer Pants (6903, -0.09 DPS) [dungeon]; Pristine Leggings (253987, -0.27 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.4 spell_power points (1.14 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Nimbus Boots (6998, -0.48 DPS) [quest]; Spidersilk Boots (4320, -1.34 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.77 DPS) | yes | Minor Channeling Ring (1449, -0.11 DPS) [quest]; Black Widow Band (6199, -0.40 DPS) [world]; Snake Hoop (6750, -0.40 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.29 DPS) [world]; Snake Hoop (6750, -0.29 DPS) [quest]; Minor Channeling Ring (1449, -0.74 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.99 DPS) | yes | Glimmering Staff (249392, -0.44 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.46 DPS) [world_drop]; Channeler's Staff (4437, -0.57 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.9 spell_power points (1.09 DPS) | yes | Eye of Paleth (2943, -0.65 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.65 DPS) [world]; Dwarven Tome (279898, -0.85 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 304.0 spell_power points (33.54 DPS) | yes | Starfaller (13063, -0.47 DPS) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.54 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-543110401201300240)

Set DPS (verified): 88.3. Weights run: 1.0s. Verify run: 0.7s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.793 ± 0.012, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.592 per %), hit=0.276 ± 0.003 per rating point (10 rating = 1%, 2.760 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.59 DPS) | yes | Augural Shroud (2620, -0.26 DPS) [world]; Corpseshroud (10574, -0.73 DPS) [dungeon]; Enchanter's Cowl (4322, -0.87 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.8 spell_power points (1.45 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.74 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.77 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.77 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.3 spell_power points (2.14 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.1 spell_power points (1.99 DPS) | yes | Guardian Cloak (5965, -0.76 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.92 DPS) [vendor]; Long Silken Cloak (4326, -1.77 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.8 spell_power points (3.30 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.40 DPS) [crafted]; Elemental Raiment (9434, -0.71 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.11 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.23 DPS) [world_drop]; Condor Bracers (15864, -0.25 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (2.61 DPS) | yes | Red Mageweave Gloves (10018, -0.55 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.76 DPS) [crafted]; Gilded Handwraps (254021, -0.94 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.9 spell_power points (2.33 DPS) | yes | Gilded Cord (254037, -0.56 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.68 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (2.90 DPS) | yes | Crimson Silk Pantaloons (7062, -0.52 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.01 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.26 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.96 DPS) | yes | Gilded Slippers (254001, -0.96 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.56 DPS) [dungeon]; Spidersilk Boots (4320, -1.71 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (1.82 DPS) | yes | Ring of Forlorn Spirits (2043, -0.83 DPS) [quest]; Reedknot Ring (9622, -0.96 DPS) [quest]; Minor Channeling Ring (1449, -1.01 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.11 DPS) | yes | Reedknot Ring (9622, -0.25 DPS) [quest]; Minor Channeling Ring (1449, -0.30 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.18 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (88.3 DPS) | yes | Windweaver Staff (7757, -1.00 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.36 DPS) [dungeon]; Gut Ripper (2164, -4.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 323.7 spell_power points (39.97 DPS) | yes | Umbral Wand (5216, -1.25 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.38 DPS) [dungeon]; Twisted Nether Wand (249144, -5.23 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 521000000000000000-00000000000000000-543110401201300251)

Set DPS (verified): 138.6. Weights run: 0.9s. Verify run: 0.8s. 423 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.262 ± 0.021, crit=0.054 ± 0.002 per rating point (14 rating = 1%, 0.761 per %), hit=0.440 ± 0.006 per rating point (10 rating = 1%, 4.402 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 44.2 spell_power points (5.38 DPS) | yes | Dreamweave Circlet (10041, -1.29 DPS) [crafted]; Soulcatcher Halo (10630, -1.54 DPS) [dungeon]; Chief Architect's Monocle (11839, -2.61 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.7 spell_power points (2.15 DPS) | yes | Scorn's Icy Choker (23169, -0.38 DPS) [dungeon]; Mindburst Medallion (11196, -0.50 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.61 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 35.7 spell_power points (4.34 DPS) | yes | Kentic Amice (11624, -0.65 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.19 DPS) [crafted]; Inquisitor's Shawl (19507, -1.50 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.6 spell_power points (2.62 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.15 DPS) [dungeon]; Runecloth Cloak (13860, -0.30 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.47 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 44.2 spell_power points (5.38 DPS) | yes | Robes of Insight (940, -1.54 DPS) [world_drop]; Runecloth Tunic (13857, -1.62 DPS) [crafted]; Runecloth Robe (13858, -2.47 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 18.9 spell_power points (2.30 DPS) | yes | Bloodband Bracers (11469, -0.31 DPS) [quest]; Nethergeld Cuffs (254061, -0.38 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.46 DPS) [quest] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 35.4 spell_power points (4.30 DPS) | yes | Raider Handwraps (272098, -0.08 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.34 DPS) [vendor]; Red Mageweave Gloves (10018, -1.43 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 30.0 spell_power points (3.64 DPS) | yes | Satyrmane Sash (17755, -0.41 DPS) [dungeon]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Deathmage Sash (10771, -0.49 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 35.6 spell_power points (4.33 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.94 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.20 DPS) [dungeon]; Red Mageweave Pants (10009, -3.03 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.92 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.03 DPS) [vendor]; Gilded Sandals (254107, -0.20 DPS) [crafted]; Southsea Mojo Boots (20641, -0.26 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 18.9 spell_power points (2.30 DPS) | yes | Philanthropist's Ring (281635, -0.17 DPS) [quest]; Mindseye Circle (10634, -0.46 DPS) [dungeon]; Band of the Unicorn (7553, -0.72 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.8 spell_power points (2.17 DPS) | yes | Mindseye Circle (10634, -0.33 DPS) [dungeon]; Band of the Unicorn (7553, -0.59 DPS) [world_drop]; Philanthropist's Ring (281635, -0.81 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (138.6 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (138.6 DPS) | yes | Blessed Prayer Beads (19990, -1.19 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (138.6 DPS) | yes | Spellshifter Rod (9527, -0.92 DPS) [quest]; Radiant Staff (249453, -1.53 DPS) [crafted]; Barman Shanker (12791, -6.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 431.8 spell_power points (52.50 DPS) | yes | Woestave (20082, -0.71 DPS, sim-verified) [quest]; Noxious Shooter (17745, -1.89 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Virtuous Hands; waist: Dawnspire Cord; legs: Spellshock Leggings; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524111001300000000-00000000000000000-543110401201300251)

Set DPS (verified): 280.7. Weights run: 1.1s. Verify run: 0.8s. 1129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.199 ± 0.031, crit=0.128 ± 0.004 per rating point (14 rating = 1%, 1.796 per %), hit=0.809 ± 0.009 per rating point (10 rating = 1%, 8.090 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 49.4 spell_power points (6.89 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Magister's Crown (16686, -2.86 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.6 spell_power points (4.27 DPS) | yes | Beads of Ogre Mojo (22149, -0.45 DPS) [quest]; Archlight Talisman (15856, -0.92 DPS) [quest]; Lady Maye's Pendant (14558, -1.09 DPS) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.8 spell_power points (6.53 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.34 DPS) [vendor]; Virtuous Epaulets (226955, -0.74 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.7 spell_power points (4.70 DPS) | yes | Hide of the Wild (18510, -1.07 DPS) [crafted]; Crystalline Threaded Cape (20697, -1.24 DPS) [world]; Spritecaster Cape (11623, -1.74 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 47.6 spell_power points (6.64 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -2.51 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.6 spell_power points (4.41 DPS) | yes | Sublime Wristguards (18497, -1.06 DPS) [dungeon]; Runecloth Cuffs (254123, -1.20 DPS) [crafted]; Marshal's Satin Bracers (17606, -1.23 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 40.4 spell_power points (5.64 DPS) | yes | Marshal's Satin Gloves (17608, -0.25 DPS) [vendor]; Marshal's Satin Grips (231617, -0.25 DPS) [vendor]; Virtuous Hands (226958, -0.81 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 59.5 spell_power points (8.30 DPS) | yes | Magister's Belt (16685, -3.81 DPS) [dungeon]; Dustfeather Sash (12589, -4.03 DPS) [dungeon]; Belt of the Archmage (18405, -5.98 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (22752) | Silverwing Sentinels [rep] | 50.8 spell_power points (7.09 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Marshal's Satin Legguards (231626, -0.78 DPS) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 41.2 spell_power points (5.75 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -2.55 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (280.7 DPS) | yes | Songstone of Ironforge (12543, -1.23 DPS) [quest]; Maiden's Circle (13001, -1.23 DPS) [world_drop]; Naglering (11669, -8.11 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (280.7 DPS) | yes | Songstone of Ironforge (12543, -0.42 DPS) [quest]; Maiden's Circle (13001, -0.42 DPS) [world_drop]; Naglering (11669, -8.92 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (280.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (280.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -1.74 DPS, sim-verified) [quest] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (280.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.42 DPS) [world]; Hand of Edward the Odd (2243, -7.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 555.8 spell_power points (77.60 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -2.34 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.76 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.06 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Second Wind; trinket2: Burst of Knowledge; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 34.5. Weights run: 0.9s. Verify run: 0.7s. 140 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.306 ± 0.008, crit=0.108 ± 0.003 per rating point (14 rating = 1%, 1.517 per %), hit=0.343 ± 0.002 per rating point (10 rating = 1%, 3.430 per %), spell_haste=not significant (-0.520 ± 0.135), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Shadow Goggles (4373, -1.34 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.7 spell_power points (0.71 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.34 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.50 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.37 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.42 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.60 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -0.95 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.8 spell_power points (0.17 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest]; Owlbeard Bracers (16981, -0.55 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Apothecary Gloves (10919, -0.27 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.48 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.25 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.59 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.4 spell_power points (1.04 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.75 DPS) | yes | Red Woolen Boots (4313, -0.39 DPS) [crafted]; Pristine Boots (253889, -0.39 DPS) [crafted]; Feather Padded Treads (285345, -0.71 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Loop of Sacrifice (281673, -0.32 DPS) [quest]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.27 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.24 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.1 spell_power points (0.28 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.11 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 247.4 spell_power points (22.58 DPS) | yes | Firebelcher (5243, -2.29 DPS) [dungeon]; Skycaller (12984, -2.44 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 000000000000000000-00000000000000000-543110401200000000)

Set DPS (verified): 52.8. Weights run: 0.9s. Verify run: 0.7s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.481 ± 0.008, crit=0.045 ± 0.002 per rating point (14 rating = 1%, 0.632 per %), hit=0.196 ± 0.002 per rating point (10 rating = 1%, 1.961 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.21 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.22 DPS) [crafted]; Embalmed Shroud (7691, -0.33 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.9 spell_power points (1.09 DPS) | yes | Darkspear Warding Pendant (272075, -0.77 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.88 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.88 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.3 spell_power points (1.47 DPS) | yes | Death Speaker Mantle (6685, -0.22 DPS) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.55 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.3 spell_power points (1.68 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted]; Tree Bark Jacket (1486, -0.69 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.99 DPS) | yes | Glowing Magical Bracelets (13106, -0.51 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.67 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.67 DPS) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.4 spell_power points (0.93 DPS) | yes | Serpent Gloves (5970, -0.22 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.4 spell_power points (1.37 DPS) | yes | Warsong Sash (16975, -0.16 DPS) [quest]; Belt of Arugal (6392, -0.22 DPS) [dungeon]; Invoker's Cord (215366, -0.34 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.9 spell_power points (1.42 DPS) | yes | Gaze Dreamer Pants (6903, -0.09 DPS) [dungeon]; Pristine Leggings (253987, -0.27 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.4 spell_power points (1.14 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Boots of the Enchanter (4325, -0.59 DPS) [crafted]; Spidersilk Boots (4320, -1.04 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.77 DPS) | yes | Black Widow Band (6199, -0.40 DPS) [world]; Snake Hoop (6750, -0.40 DPS) [quest]; Sludge-Stained Band (286535, -0.44 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.66 DPS) | yes | Snake Hoop (6750, -0.29 DPS) [quest]; Sludge-Stained Band (286535, -0.33 DPS) [world]; Black Widow Band (6199, -0.58 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.99 DPS) | yes | Glimmering Staff (249392, -0.41 DPS) [crafted]; Twisted Chanter's Staff (890, -0.46 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.46 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.9 spell_power points (1.09 DPS) | yes | Orb of Souls (249395, -0.65 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.66 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -0.96 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 304.0 spell_power points (33.54 DPS) | yes | Starfaller (13063, -0.47 DPS) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.54 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-543110401201300240)

Set DPS (verified): 84.6. Weights run: 1.0s. Verify run: 0.7s. 311 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.793 ± 0.012, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.592 per %), hit=0.276 ± 0.003 per rating point (10 rating = 1%, 2.760 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.59 DPS) | yes | Augural Shroud (2620, -0.26 DPS) [world]; Corpseshroud (10574, -0.73 DPS) [dungeon]; Enchanter's Cowl (4322, -0.87 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.8 spell_power points (1.45 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.44 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.77 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.77 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.3 spell_power points (2.14 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.1 spell_power points (1.99 DPS) | yes | Guardian Cloak (5965, -0.76 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.92 DPS) [vendor]; Long Silken Cloak (4326, -1.41 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.8 spell_power points (3.30 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.40 DPS) [crafted]; Elemental Raiment (9434, -0.71 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.3 spell_power points (1.28 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Windchaser Cuffs (14429, -0.40 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (2.61 DPS) | yes | Red Mageweave Gloves (10018, -0.57 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.76 DPS) [crafted]; Gilded Handwraps (254021, -0.94 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.9 spell_power points (2.33 DPS) | yes | Defiler's Cloth Girdle (20166, -0.36 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.56 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (2.90 DPS) | yes | Crimson Silk Pantaloons (7062, -0.64 DPS) [crafted]; Abomination Skin Leggings (23173, -1.01 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.26 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.96 DPS) | yes | Gilded Slippers (254001, -0.45 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.56 DPS) [dungeon]; Spidersilk Boots (4320, -1.71 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (1.82 DPS) | yes | Reedknot Ring (9622, -0.96 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.08 DPS) [vendor]; Black Widow Band (6199, -1.14 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.11 DPS) | yes | Sea Giant's Toe Ring (274746, -0.37 DPS) [vendor]; Black Widow Band (6199, -0.43 DPS) [world]; Reedknot Ring (9622, -1.29 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (84.6 DPS) | yes | Windweaver Staff (7757, -1.00 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.36 DPS) [dungeon]; Gut Ripper (2164, -3.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 323.7 spell_power points (39.97 DPS) | yes | Umbral Wand (5216, -1.10 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.38 DPS) [dungeon]; Twisted Nether Wand (249144, -5.23 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 521000000000000000-00000000000000000-543110401201300251)

Set DPS (verified): 131.3. Weights run: 0.9s. Verify run: 0.8s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.262 ± 0.021, crit=0.054 ± 0.002 per rating point (14 rating = 1%, 0.761 per %), hit=0.440 ± 0.006 per rating point (10 rating = 1%, 4.402 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 44.2 spell_power points (5.38 DPS) | yes | Dreamweave Circlet (10041, -1.29 DPS) [crafted]; Soulcatcher Halo (10630, -1.54 DPS) [dungeon]; Chief Architect's Monocle (11839, -1.74 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.7 spell_power points (2.15 DPS) | yes | Scorn's Icy Choker (23169, -0.38 DPS) [dungeon]; Mindburst Medallion (11196, -0.50 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.61 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 35.7 spell_power points (4.34 DPS) | yes | Kentic Amice (11624, -0.65 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.19 DPS) [crafted]; Inquisitor's Shawl (19507, -1.50 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 23.4 spell_power points (2.84 DPS) | yes | Spritecaster Cape (11623, -0.22 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.36 DPS) [dungeon]; Runecloth Cloak (13860, -0.52 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 44.2 spell_power points (5.38 DPS) | yes | Robes of Insight (940, -1.54 DPS) [world_drop]; Runecloth Tunic (13857, -1.62 DPS) [crafted]; Runecloth Robe (13858, -1.69 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 18.9 spell_power points (2.30 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.60 DPS, sim-verified) [quest] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 35.4 spell_power points (4.30 DPS) | yes | Raider Handwraps (272098, -0.08 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.34 DPS) [vendor]; Red Mageweave Gloves (10018, -1.43 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 30.0 spell_power points (3.64 DPS) | yes | Satyrmane Sash (17755, -0.41 DPS) [dungeon]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Deathmage Sash (10771, -0.49 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 35.6 spell_power points (4.33 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.94 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.20 DPS) [dungeon]; Red Mageweave Pants (10009, -2.85 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.92 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.03 DPS) [vendor]; Gilded Sandals (254107, -0.20 DPS) [crafted]; Southsea Mojo Boots (20641, -0.26 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 18.9 spell_power points (2.30 DPS) | yes | Philanthropist's Ring (281635, -0.17 DPS) [quest]; Mindseye Circle (10634, -0.46 DPS) [dungeon]; Band of the Unicorn (7553, -0.72 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.8 spell_power points (2.17 DPS) | yes | Mindseye Circle (10634, -0.33 DPS) [dungeon]; Band of the Unicorn (7553, -0.59 DPS) [world_drop]; Philanthropist's Ring (281635, -1.04 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (131.3 DPS) | yes | Uther's Strength (11302, -0.87 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (131.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (131.3 DPS) | yes | Spellshifter Rod (9527, -0.92 DPS) [quest]; Radiant Staff (249453, -1.53 DPS) [crafted]; Barman Shanker (12791, -5.36 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 431.8 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.89 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Virtuous Hands; waist: Dawnspire Cord; legs: Spellshock Leggings; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524111001300000000-00000000000000000-543110401201300251)

Set DPS (verified): 281.1. Weights run: 1.1s. Verify run: 0.8s. 1120 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.199 ± 0.031, crit=0.128 ± 0.004 per rating point (14 rating = 1%, 1.796 per %), hit=0.809 ± 0.009 per rating point (10 rating = 1%, 8.090 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 49.4 spell_power points (6.89 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -2.03 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.6 spell_power points (4.27 DPS) | yes | Beads of Ogre Mojo (22149, -0.45 DPS) [quest]; Archlight Talisman (15856, -0.92 DPS) [quest]; Lady Maye's Pendant (14558, -1.09 DPS) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.8 spell_power points (6.53 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.34 DPS) [vendor]; Virtuous Epaulets (226955, -0.74 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.7 spell_power points (4.70 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -1.24 DPS) [world]; Deep Woodlands Cloak (19121, -1.52 DPS) [quest] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 47.6 spell_power points (6.64 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -1.63 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.6 spell_power points (4.41 DPS) | yes | Sublime Wristguards (18497, -1.06 DPS) [dungeon]; Runecloth Cuffs (254123, -1.20 DPS) [crafted]; General's Satin Bracers (17619, -1.23 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 40.4 spell_power points (5.64 DPS) | yes | General's Satin Gloves (17620, -0.25 DPS) [vendor]; General's Satin Grips (231613, -0.25 DPS) [vendor]; Virtuous Hands (226958, -0.81 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 59.5 spell_power points (8.30 DPS) | yes | Magister's Belt (16685, -3.81 DPS) [dungeon]; Dustfeather Sash (12589, -4.03 DPS) [dungeon]; Belt of the Archmage (18405, -4.18 DPS, sim-verified) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 50.8 spell_power points (7.09 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, -0.78 DPS) [vendor]; Sentinel's Silk Leggings (237815, -3.43 DPS, sim-verified) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 41.2 spell_power points (5.75 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -0.56 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (281.1 DPS) | yes | Eye of Orgrimmar (12545, -1.23 DPS) [quest]; Maiden's Circle (13001, -1.23 DPS) [world_drop]; Naglering (11669, -7.91 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (281.1 DPS) | yes | Eye of Orgrimmar (12545, -0.42 DPS) [quest]; Maiden's Circle (13001, -0.42 DPS) [world_drop]; Naglering (11669, -7.38 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (281.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (281.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -2.16 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (281.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -0.11 DPS) [dungeon]; Hand of Edward the Odd (2243, -6.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 555.8 spell_power points (77.60 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.21 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.76 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.06 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Outrider's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

