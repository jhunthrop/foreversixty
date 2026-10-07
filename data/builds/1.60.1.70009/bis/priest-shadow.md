# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 30.6. Weights run: 0.7s. Verify run: 0.5s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.290 ± 0.007, crit=0.113 ± 0.003 per rating point (14 rating = 1%, 1.589 per %), hit=0.355 ± 0.003 per rating point (10 rating = 1%, 3.550 per %), spell_haste=not significant (-0.497 ± 0.144), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.54 DPS) | yes | Shadow Goggles (4373, -1.07 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.6 spell_power points (0.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.13 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.09 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.4 spell_power points (0.58 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -0.49 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.4 spell_power points (0.13 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.04 DPS) [quest]; Repurposed Hair Band (281256, -0.08 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.63 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.46 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.35 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.3 spell_power points (1.02 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Rumpled Kilt (274741, -0.57 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.73 DPS) | yes | Feather Padded Treads (285345, -0.21 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.39 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.50 DPS) | yes | Sludge-Stained Band (286535, -0.23 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.42 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.45 DPS) | yes | Sludge-Stained Band (286535, -0.18 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.9 spell_power points (0.26 DPS) | yes | Channeler's Staff (4437, -0.05 DPS) [world]; Lesser Staff of the Spire (1300, -0.10 DPS) [world]; Staff of Westfall (2042, -0.13 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 250.8 spell_power points (22.58 DPS) | yes | Skycaller (12984, -1.02 DPS) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-543110401200000000)

Set DPS (verified): 62.6. Weights run: 0.7s. Verify run: 0.6s. 244 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.651 ± 0.010, crit=0.046 ± 0.002 per rating point (14 rating = 1%, 0.640 per %), hit=0.240 ± 0.003 per rating point (10 rating = 1%, 2.400 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.5 spell_power points (1.36 DPS) | yes | Holy Shroud (2721, -0.16 DPS) [world_drop]; Silk Headband (7050, -0.38 DPS) [crafted]; Embalmed Shroud (7691, -0.49 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.9 spell_power points (1.18 DPS) | yes | Crystal Starfire Medallion (5003, -0.90 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.90 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.06 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.9 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Magician's Mantle (12998, -0.43 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.6 spell_power points (0.61 DPS) | yes | Cloak of Rot (4462, +0.00 DPS, sim-verified) [world]; Darkspear Raider's Cloak (272078, -0.04 DPS) [vendor]; Hillman's Cloak (3719, -0.07 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.5 spell_power points (1.89 DPS) | yes | Death Speaker Robes (6682, -0.36 DPS) [dungeon]; Tree Bark Jacket (1486, -0.48 DPS) [dungeon]; Pristine Gown (253961, -0.64 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Glowing Magical Bracelets (13106, -0.55 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.55 DPS) [world_drop]; Stonecloth Bindings (14416, -0.62 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.2 spell_power points (1.21 DPS) | yes | Shilly Mitts (9609, -0.39 DPS, sim-verified) [quest]; Serpent Gloves (5970, -0.45 DPS) [dungeon]; Truefaith Gloves (7049, -0.46 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.0 spell_power points (1.40 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Crimson Silk Belt (7055, -0.26 DPS) [crafted]; Invoker's Cord (215366, -0.29 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.54 DPS) | yes | Pristine Leggings (253987, -0.29 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.38 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.47 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.6 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.21 DPS) [crafted]; Nimbus Boots (6998, -0.60 DPS) [quest]; Acidic Walkers (9454, -1.16 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.76 DPS) | yes | Black Widow Band (6199, -0.26 DPS) [world]; Snake Hoop (6750, -0.26 DPS) [quest]; Minor Channeling Ring (1449, -1.35 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (62.6 DPS) | yes | Black Widow Band (6199, -0.16 DPS) [world]; Snake Hoop (6750, -0.16 DPS) [quest]; Minor Channeling Ring (1449, -1.19 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Glimmering Staff (249392, -0.20 DPS) [crafted]; Twisted Chanter's Staff (890, -0.27 DPS) [world_drop]; Channeler's Staff (4437, -0.41 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.9 spell_power points (1.18 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.68 DPS) [vendor]; Eye of Paleth (2943, -0.75 DPS) [quest]; Dwarven Tome (279898, -0.85 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 309.5 spell_power points (33.54 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.54 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-543110401201300240)

Set DPS (verified): 101.3. Weights run: 0.8s. Verify run: 0.5s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.059 ± 0.019, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.601 per %), hit=0.381 ± 0.005 per rating point (10 rating = 1%, 3.812 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 21.6 spell_power points (2.64 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.07 DPS) [crafted]; Corpseshroud (10574, -0.18 DPS) [dungeon]; Thinking Cap (2624, -0.44 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.4 spell_power points (1.64 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest]; Triune Amulet (7722, -0.73 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.73 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.8 spell_power points (2.54 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Bloodmage Mantle (7684, -0.27 DPS) [dungeon]; Death Speaker Mantle (6685, -0.38 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 18.5 spell_power points (2.27 DPS) | yes | Darkspear Raider's Cloak (272077, -0.78 DPS, sim-verified) [vendor]; Long Silken Cloak (4326, -0.89 DPS) [crafted]; Guardian Cloak (5965, -0.89 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.4 spell_power points (3.47 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.20 DPS) [crafted]; Green Silk Armor (7065, -0.68 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 9.5 spell_power points (1.17 DPS) | yes | Arcane Runed Bracers (4744, -0.06 DPS) [quest]; Spidertank Oilrag (9448, -0.06 DPS) [dungeon]; Mistscape Bracers (4045, -0.13 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.2 spell_power points (2.72 DPS) | yes | Red Mageweave Gloves (10018, -0.08 DPS) [crafted]; Stormcloth Gloves (10011, -0.68 DPS) [crafted]; Town Clerk's Mittens (270029, -0.81 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.9 spell_power points (2.80 DPS) | yes | Highlander's Cloth Girdle (20098, -0.64 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.79 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.7 spell_power points (3.27 DPS) | yes | Crimson Silk Pantaloons (7062, -0.61 DPS) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.94 DPS) | yes | Gilded Slippers (254001, -1.17 DPS) [crafted]; Acidic Walkers (9454, -1.29 DPS) [dungeon]; Spidersilk Boots (4320, -1.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.4 spell_power points (2.00 DPS) | yes | Ring of Forlorn Spirits (2043, -1.02 DPS) [quest]; Voodoo Band (1996, -1.10 DPS) [world]; Black Widow Band (6199, -1.10 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.10 DPS) | yes | Voodoo Band (1996, -0.19 DPS) [world]; Black Widow Band (6199, -0.19 DPS) [world]; Ring of Forlorn Spirits (2043, -1.06 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (101.3 DPS) | yes | Windweaver Staff (7757, -0.50 DPS) [dungeon]; Staff of Jordan (873, -1.02 DPS) [world_drop]; Gut Ripper (2164, -4.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 327.8 spell_power points (40.16 DPS) | yes | Umbral Wand (5216, -0.83 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.57 DPS) [dungeon]; Twisted Nether Wand (249144, -5.42 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 521000000000000000-00000000000000000-543110401201300251)

Set DPS (verified): 158.2. Weights run: 0.8s. Verify run: 0.6s. 423 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.431 ± 0.029, crit=0.053 ± 0.002 per rating point (14 rating = 1%, 0.743 per %), hit=0.522 ± 0.007 per rating point (10 rating = 1%, 5.221 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 47.6 spell_power points (5.74 DPS) | yes | Chief Architect's Monocle (11839, -1.23 DPS, sim-verified) [dungeon]; Soulcatcher Halo (10630, -1.43 DPS) [dungeon]; Dreamweave Circlet (10041, -1.48 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.0 spell_power points (2.41 DPS) | yes | Scorn's Icy Choker (23169, -0.54 DPS) [dungeon]; Mindburst Medallion (11196, -0.66 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.69 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 38.8 spell_power points (4.67 DPS) | yes | Kentic Amice (11624, -0.74 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.24 DPS) [crafted]; Inquisitor's Shawl (19507, -1.58 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.6 spell_power points (2.72 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.08 DPS) [dungeon]; Runecloth Cloak (13860, -0.26 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.31 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 47.6 spell_power points (5.74 DPS) | yes | Runecloth Robe (13858, -1.48 DPS) [crafted]; Runecloth Tunic (13857, -1.79 DPS) [crafted]; Robes of Insight (940, -2.40 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 21.5 spell_power points (2.59 DPS) | yes | Bloodband Bracers (11469, -0.43 DPS) [quest]; Forgotten Wraps (9433, -0.52 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.52 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 38.5 spell_power points (4.64 DPS) | yes | Virtuous Hands (226958, -0.11 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.52 DPS) [vendor]; Red Mageweave Gloves (10018, -1.59 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 33.2 spell_power points (4.00 DPS) | yes | Satyrmane Sash (17755, -0.59 DPS) [dungeon]; Ban'thok Sash (11662, -0.59 DPS) [dungeon]; Deathmage Sash (10771, -1.07 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.3 spell_power points (4.49 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.89 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.03 DPS) [dungeon]; Red Mageweave Pants (10009, -2.91 DPS, sim-verified) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 26.1 spell_power points (3.14 DPS) | yes | Earthen Silk Slippers (254013, -0.25 DPS) [crafted]; Gilded Sandals (254107, -0.27 DPS) [crafted]; Southsea Mojo Boots (20641, -0.28 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 21.5 spell_power points (2.59 DPS) | yes | Philanthropist's Ring (281635, -0.35 DPS) [quest]; Mindseye Circle (10634, -0.52 DPS) [dungeon]; Band of the Unicorn (7553, -1.02 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.0 spell_power points (2.29 DPS) | yes | Mindseye Circle (10634, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.72 DPS) [world_drop]; Philanthropist's Ring (281635, -1.22 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (158.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (158.2 DPS) | yes | Blessed Prayer Beads (19990, -1.20 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (158.2 DPS) | yes | Spellshifter Rod (9527, -1.03 DPS) [quest]; Radiant Staff (249453, -1.72 DPS) [crafted]; Barman Shanker (12791, -6.68 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 435.9 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.90 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524111001300000000-00000000000000000-543110401201300251)

Set DPS (verified): 307.8. Weights run: 0.8s. Verify run: 0.6s. 1129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.320 ± 0.037, crit=0.127 ± 0.005 per rating point (14 rating = 1%, 1.772 per %), hit=0.908 ± 0.011 per rating point (10 rating = 1%, 9.081 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 51.5 spell_power points (7.11 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Magister's Crown (16686, -4.37 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 32.2 spell_power points (4.44 DPS) | yes | Archlight Talisman (15856, -0.96 DPS) [quest]; Lady Maye's Pendant (14558, -0.98 DPS) [world_drop]; Beads of Ogre Mojo (22149, -2.00 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.6 spell_power points (6.70 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.13 DPS) [vendor]; Virtuous Epaulets (226955, -0.64 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.6 spell_power points (4.92 DPS) | yes | Hide of the Wild (18510, -1.16 DPS) [crafted]; Crystalline Threaded Cape (20697, -1.43 DPS) [world]; Spritecaster Cape (11623, -1.89 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 49.8 spell_power points (6.87 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -4.20 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 32.6 spell_power points (4.49 DPS) | yes | Sublime Wristguards (18497, -1.02 DPS) [dungeon]; Marshal's Satin Bracers (17606, -1.03 DPS) [pvp]; Runecloth Cuffs (254123, -1.15 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 43.6 spell_power points (6.02 DPS) | yes | Marshal's Satin Gloves (17608, -0.48 DPS) [vendor]; Marshal's Satin Grips (231617, -0.48 DPS) [vendor]; Virtuous Hands (226958, -2.42 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 63.1 spell_power points (8.71 DPS) | yes | Magister's Belt (16685, -3.92 DPS) [dungeon]; Dustfeather Sash (12589, -4.19 DPS) [dungeon]; Belt of the Archmage (18405, -8.43 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (22752) | Silverwing Sentinels [rep] | 53.1 spell_power points (7.32 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Marshal's Satin Legguards (231626, -0.74 DPS) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 43.1 spell_power points (5.95 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -3.54 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (307.8 DPS) | yes | Channeler's Ring (272406, -1.24 DPS) [vendor]; Maiden's Circle (13001, -1.28 DPS) [world_drop]; Naglering (11669, -10.33 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (307.8 DPS) | yes | Channeler's Ring (272406, -0.37 DPS) [vendor]; Maiden's Circle (13001, -0.41 DPS) [world_drop]; Naglering (11669, -11.94 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (307.8 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -6.67 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (307.8 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (307.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -1.52 DPS) [dungeon]; Hand of Edward the Odd (2243, -9.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 562.3 spell_power points (77.59 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -4.76 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.71 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.93 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Second Wind; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 31.8. Weights run: 0.7s. Verify run: 0.6s. 140 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.290 ± 0.007, crit=0.113 ± 0.003 per rating point (14 rating = 1%, 1.589 per %), hit=0.355 ± 0.003 per rating point (10 rating = 1%, 3.550 per %), spell_haste=not significant (-0.497 ± 0.144), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.54 DPS) | yes | Shadow Goggles (4373, -0.83 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.6 spell_power points (0.68 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.32 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.46 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.40 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.4 spell_power points (0.58 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -0.78 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.7 spell_power points (0.16 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest]; Owlbeard Bracers (16981, -0.41 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.63 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Apothecary Gloves (10919, -0.27 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.46 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.51 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.3 spell_power points (1.02 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Rumpled Kilt (274741, -0.57 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.73 DPS) | yes | Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.39 DPS) [crafted]; Feather Padded Treads (285345, -0.69 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.45 DPS) | yes | Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Loop of Sacrifice (281673, -0.32 DPS) [quest]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.27 DPS) | yes | Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.9 spell_power points (0.26 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.05 DPS) [world]; Lesser Staff of the Spire (1300, -0.10 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 250.8 spell_power points (22.58 DPS) | yes | Skycaller (12984, -1.68 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 000000000000000000-00000000000000000-543110401200000000)

Set DPS (verified): 55.5. Weights run: 0.7s. Verify run: 0.6s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.651 ± 0.010, crit=0.046 ± 0.002 per rating point (14 rating = 1%, 0.640 per %), hit=0.240 ± 0.003 per rating point (10 rating = 1%, 2.400 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.5 spell_power points (1.36 DPS) | yes | Holy Shroud (2721, -0.37 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.38 DPS) [crafted]; Embalmed Shroud (7691, -0.49 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.9 spell_power points (1.18 DPS) | yes | Darkspear Warding Pendant (272075, -0.87 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.90 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.90 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.9 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Magician's Mantle (12998, -0.43 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.2 spell_power points (0.56 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.02 DPS) [crafted]; Windsong Drape (15468, -0.02 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.5 spell_power points (1.89 DPS) | yes | Tree Bark Jacket (1486, -0.48 DPS) [dungeon]; Death Speaker Robes (6682, -0.52 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.64 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Nightsky Wristbands (6407, -0.55 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.55 DPS) [quest]; Glowing Magical Bracelets (13106, -0.82 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.3 spell_power points (1.00 DPS) | yes | Truefaith Gloves (7049, -0.25 DPS) [crafted]; Gnoll Casting Gloves (892, -0.35 DPS) [world]; Serpent Gloves (5970, -0.51 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.0 spell_power points (1.40 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Crimson Silk Belt (7055, -0.26 DPS) [crafted]; Warsong Sash (16975, -0.34 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.54 DPS) | yes | Gaze Dreamer Pants (6903, -0.29 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.29 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.47 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.6 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.21 DPS) [crafted]; Boots of the Enchanter (4325, -0.71 DPS) [crafted]; Acidic Walkers (9454, -0.79 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.76 DPS) | yes | Black Widow Band (6199, -0.26 DPS) [world]; Snake Hoop (6750, -0.26 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.65 DPS) | yes | Snake Hoop (6750, -0.16 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Black Widow Band (6199, -0.56 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Glimmering Staff (249392, -0.20 DPS) [crafted]; Twisted Chanter's Staff (890, -0.27 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.27 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.9 spell_power points (1.18 DPS) | yes | Witch's Finger (16887, -0.69 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.75 DPS) [world]; Tome of the Darkspear Prophecy (272090, -0.90 DPS, sim-verified) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 309.5 spell_power points (33.54 DPS) | yes | Starfaller (13063, -0.39 DPS) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.54 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-543110401201300240)

Set DPS (verified): 90.1. Weights run: 0.8s. Verify run: 0.5s. 311 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.059 ± 0.019, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.601 per %), hit=0.381 ± 0.005 per rating point (10 rating = 1%, 3.812 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 21.6 spell_power points (2.64 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.07 DPS) [crafted]; Corpseshroud (10574, -0.18 DPS) [dungeon]; Thinking Cap (2624, -0.44 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.4 spell_power points (1.64 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest]; Triune Amulet (7722, -0.73 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.73 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.8 spell_power points (2.54 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Bloodmage Mantle (7684, -0.27 DPS) [dungeon]; Death Speaker Mantle (6685, -0.38 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 18.5 spell_power points (2.27 DPS) | yes | Long Silken Cloak (4326, -0.89 DPS) [crafted]; Guardian Cloak (5965, -0.89 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.13 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.4 spell_power points (3.47 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.20 DPS) [crafted]; Green Silk Armor (7065, -0.68 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.5 spell_power points (1.53 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Windchaser Cuffs (14429, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.43 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.2 spell_power points (2.72 DPS) | yes | Red Mageweave Gloves (10018, -0.08 DPS) [crafted]; Stormcloth Gloves (10011, -0.68 DPS) [crafted]; Gilded Handwraps (254021, -0.84 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.9 spell_power points (2.80 DPS) | yes | Gilded Cord (254037, -0.79 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.07 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.7 spell_power points (3.27 DPS) | yes | Crimson Silk Pantaloons (7062, -0.61 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.94 DPS) | yes | Gilded Slippers (254001, -1.17 DPS) [crafted]; Acidic Walkers (9454, -1.29 DPS) [dungeon]; Spidersilk Boots (4320, -1.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.4 spell_power points (2.00 DPS) | yes | Ogremind Ring (1993, -1.10 DPS) [world_drop]; Voodoo Band (1996, -1.10 DPS) [world]; Black Widow Band (6199, -1.10 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.10 DPS) | yes | Ogremind Ring (1993, -0.19 DPS) [world_drop]; Voodoo Band (1996, -0.19 DPS) [world]; Black Widow Band (6199, -0.64 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (90.1 DPS) | yes | Windweaver Staff (7757, -0.50 DPS) [dungeon]; Staff of Jordan (873, -1.02 DPS) [world_drop]; Gut Ripper (2164, -3.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 327.8 spell_power points (40.16 DPS) | yes | Umbral Wand (5216, -1.20 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.57 DPS) [dungeon]; Twisted Nether Wand (249144, -5.42 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 521000000000000000-00000000000000000-543110401201300251)

Set DPS (verified): 140.5. Weights run: 0.8s. Verify run: 0.6s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.431 ± 0.029, crit=0.053 ± 0.002 per rating point (14 rating = 1%, 0.743 per %), hit=0.522 ± 0.007 per rating point (10 rating = 1%, 5.221 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 47.6 spell_power points (5.74 DPS) | yes | Soulcatcher Halo (10630, -1.43 DPS) [dungeon]; Dreamweave Circlet (10041, -1.48 DPS) [crafted]; Chief Architect's Monocle (11839, -6.78 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.0 spell_power points (2.41 DPS) | yes | Mindburst Medallion (11196, -0.66 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.69 DPS) [quest]; Scorn's Icy Choker (23169, -1.87 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 38.8 spell_power points (4.67 DPS) | yes | Kentic Amice (11624, -0.74 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.24 DPS) [crafted]; Inquisitor's Shawl (19507, -1.58 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 24.9 spell_power points (3.00 DPS) | yes | Spritecaster Cape (11623, -0.28 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.36 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 47.6 spell_power points (5.74 DPS) | yes | Runecloth Robe (13858, -1.48 DPS) [crafted]; Runecloth Tunic (13857, -1.79 DPS) [crafted]; Robes of Insight (940, -6.48 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 21.5 spell_power points (2.59 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -1.71 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 38.5 spell_power points (4.64 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.52 DPS) [vendor]; Virtuous Hands (226958, -1.52 DPS, sim-verified) [vendor]; Red Mageweave Gloves (10018, -1.59 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 33.2 spell_power points (4.00 DPS) | yes | Satyrmane Sash (17755, -0.59 DPS) [dungeon]; Ban'thok Sash (11662, -0.59 DPS) [dungeon]; Deathmage Sash (10771, -0.98 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.3 spell_power points (4.49 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.89 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.03 DPS) [dungeon]; Red Mageweave Pants (10009, -2.94 DPS, sim-verified) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.1 spell_power points (3.14 DPS) | yes | Gilded Sandals (254107, -0.27 DPS) [crafted]; Southsea Mojo Boots (20641, -0.28 DPS) [quest]; Earthen Silk Slippers (254013, -1.34 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 21.5 spell_power points (2.59 DPS) | yes | Philanthropist's Ring (281635, -0.35 DPS) [quest]; Mindseye Circle (10634, -0.52 DPS) [dungeon]; Band of the Unicorn (7553, -1.02 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.0 spell_power points (2.29 DPS) | yes | Mindseye Circle (10634, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.72 DPS) [world_drop]; Philanthropist's Ring (281635, -1.18 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (140.5 DPS) | yes | Uther's Strength (11302, -1.41 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (140.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (140.5 DPS) | yes | Spellshifter Rod (9527, -1.03 DPS) [quest]; Radiant Staff (249453, -1.72 DPS) [crafted]; Blade of Eternal Darkness (17780, -7.03 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 435.9 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.90 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524111001300000000-00000000000000000-543110401201300251)

Set DPS (verified): 288.9. Weights run: 0.8s. Verify run: 0.6s. 1120 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.320 ± 0.037, crit=0.127 ± 0.005 per rating point (14 rating = 1%, 1.772 per %), hit=0.908 ± 0.011 per rating point (10 rating = 1%, 9.081 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 51.5 spell_power points (7.11 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -2.27 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 32.2 spell_power points (4.44 DPS) | yes | Beads of Ogre Mojo (22149, -0.46 DPS) [quest]; Archlight Talisman (15856, -0.96 DPS) [quest]; Lady Maye's Pendant (14558, -0.98 DPS) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.6 spell_power points (6.70 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.13 DPS) [vendor]; Virtuous Epaulets (226955, -0.64 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.6 spell_power points (4.92 DPS) | yes | Hide of the Wild (18510, -1.16 DPS) [crafted]; Crystalline Threaded Cape (20697, -1.43 DPS) [world]; Deep Woodlands Cloak (19121, -1.62 DPS) [quest] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 49.8 spell_power points (6.87 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -2.03 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 32.6 spell_power points (4.49 DPS) | yes | Sublime Wristguards (18497, -1.02 DPS) [dungeon]; General's Satin Bracers (17619, -1.03 DPS) [pvp]; Runecloth Cuffs (254123, -1.15 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 43.6 spell_power points (6.02 DPS) | yes | General's Satin Gloves (17620, -0.48 DPS) [vendor]; General's Satin Grips (231613, -0.48 DPS) [vendor]; Virtuous Hands (226958, -1.03 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 63.1 spell_power points (8.71 DPS) | yes | Magister's Belt (16685, -3.92 DPS) [dungeon]; Dustfeather Sash (12589, -4.19 DPS) [dungeon]; Belt of the Archmage (18405, -5.20 DPS, sim-verified) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 53.1 spell_power points (7.32 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, -0.74 DPS) [vendor]; Sentinel's Silk Leggings (237815, -3.75 DPS, sim-verified) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 43.1 spell_power points (5.95 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -0.55 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (288.9 DPS) | yes | Channeler's Ring (272406, -1.24 DPS) [vendor]; Maiden's Circle (13001, -1.28 DPS) [world_drop]; Naglering (11669, -7.75 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (288.9 DPS) | yes | Channeler's Ring (272406, -0.37 DPS) [vendor]; Maiden's Circle (13001, -0.41 DPS) [world_drop]; Naglering (11669, -7.88 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (288.9 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (288.9 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -4.14 DPS, sim-verified) [quest] |
| main_hand | Trindlehaven Staff (13161) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-verified (288.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.15 DPS) [world]; Hand of Edward the Odd (2243, -4.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 562.3 spell_power points (77.59 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.22 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.71 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.93 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Outrider's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Trindlehaven Staff; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

