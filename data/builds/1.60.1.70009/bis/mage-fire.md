# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 28.4. Weights run: 0.8s. Verify run: 0.5s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.404 ± 0.011, crit=0.118 ± 0.004 per rating point (14 rating = 1%, 1.646 per %), hit=0.358 ± 0.002 per rating point (10 rating = 1%, 3.575 per %), spell_haste=not significant (0.485 ± 0.160), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.43 DPS) | yes | Shadow Goggles (4373, -0.79 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 spell_power points (0.62 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.33 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.44 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.29 DPS) | yes | Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Black Whelp Cloak (7283, -0.07 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.18 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 spell_power points (0.50 DPS) | yes | Green Woolen Robe (6243, -0.20 DPS) [crafted]; Green Woolen Vest (2582, -0.22 DPS) [crafted]; Gray Woolen Robe (2585, -1.00 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 2.0 spell_power points (0.14 DPS) | yes | Windsong Bangles (263336, -0.07 DPS) [quest]; Repurposed Hair Band (281256, -0.09 DPS) [quest]; Bright Bracers (3647, -0.77 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.50 DPS) | yes | Gnoll Casting Gloves (892, -0.07 DPS) [world]; Pristine Gloves (253913, -0.13 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.30 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 spell_power points (0.40 DPS) | yes | Keller's Girdle (2911, -0.17 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.17 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.86 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 spell_power points (0.87 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Silk-threaded Trousers (1929, -0.37 DPS) [dungeon]; Rumpled Kilt (274741, -0.52 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 spell_power points (0.61 DPS) | yes | Pristine Boots (253889, -0.31 DPS) [crafted]; Red Woolen Boots (4313, -0.33 DPS) [crafted]; Feather Padded Treads (285345, -0.67 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.41 DPS) | yes | Sludge-Stained Band (286535, -0.20 DPS) [world]; Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.36 DPS) | yes | Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop]; Sludge-Stained Band (286535, -0.67 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 4.0 spell_power points (0.29 DPS) | yes | Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Staff of Westfall (2042, -0.14 DPS) [quest]; Channeler's Staff (4437, -0.18 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 316.1 spell_power points (22.52 DPS) | yes | Skycaller (12984, -0.89 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.23 DPS) [dungeon]; Deepblaze (279896, -3.99 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 46.2. Weights run: 0.9s. Verify run: 0.6s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.560 ± 0.021, crit=0.173 ± 0.009 per rating point (14 rating = 1%, 2.420 per %), hit=0.334 ± 0.003 per rating point (10 rating = 1%, 3.340 per %), spell_haste=not significant (0.131 ± 0.295), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 11.6 spell_power points (1.04 DPS) | yes | Holy Shroud (2721, -0.05 DPS) [world_drop]; Silk Headband (7050, -0.23 DPS) [crafted]; Embalmed Shroud (7691, -0.32 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (0.93 DPS) | yes | Crystal Starfire Medallion (5003, -0.73 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.73 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.99 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.26 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.27 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.2 spell_power points (0.47 DPS) | yes | Cloak of Rot (4462, -0.07 DPS) [world]; Darkspear Raider's Cloak (272078, -0.07 DPS) [vendor]; Hillman's Cloak (3719, -0.65 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (1.46 DPS) | yes | Tree Bark Jacket (1486, -0.29 DPS) [dungeon]; Pristine Gown (253961, -0.48 DPS) [crafted]; Death Speaker Robes (6682, -0.60 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Nightsky Wristbands (6407, -0.51 DPS) [world_drop]; Stonecloth Bindings (14416, -0.56 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.99 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.2 spell_power points (0.91 DPS) | yes | Serpent Gloves (5970, -0.28 DPS) [dungeon]; Truefaith Gloves (7049, -0.31 DPS) [crafted]; Shilly Mitts (9609, -0.51 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (1.14 DPS) | yes | Belt of Arugal (6392, -0.18 DPS) [dungeon]; Crimson Silk Belt (7055, -0.25 DPS) [crafted]; Invoker's Cord (215366, -0.26 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.5 spell_power points (1.21 DPS) | yes | Gaze Dreamer Pants (6903, -0.13 DPS) [dungeon]; Pristine Leggings (253987, -0.23 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.37 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (0.98 DPS) | yes | Spidersilk Boots (4320, -0.15 DPS) [crafted]; Nimbus Boots (6998, -0.44 DPS) [quest]; Acidic Walkers (9454, -1.17 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.63 DPS) | yes | Black Widow Band (6199, -0.28 DPS) [world]; Snake Hoop (6750, -0.28 DPS) [quest]; Minor Channeling Ring (1449, -1.58 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (46.2 DPS) | yes | Black Widow Band (6199, -0.19 DPS) [world]; Snake Hoop (6750, -0.19 DPS) [quest]; Minor Channeling Ring (1449, -0.50 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Glimmering Staff (249392, -0.26 DPS) [crafted]; Twisted Chanter's Staff (890, -0.31 DPS) [world_drop]; Channeler's Staff (4437, -0.41 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.4 spell_power points (0.93 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.55 DPS) [vendor]; Eye of Paleth (2943, -0.57 DPS) [quest]; Dwarven Tome (279898, -0.60 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 372.3 spell_power points (33.48 DPS) | yes | Starfaller (13063, -0.42 DPS) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-23552100130103050-0000000000000000000)

Set DPS (verified): 78.4. Weights run: 0.9s. Verify run: 0.6s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.733 ± 0.032, crit=0.235 ± 0.016 per rating point (14 rating = 1%, 3.294 per %), hit=0.439 ± 0.004 per rating point (10 rating = 1%, 4.394 per %), spell_haste=not significant (-0.427 ± 0.532), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.29 DPS) | yes | Augural Shroud (2620, -0.29 DPS) [world]; Corpseshroud (10574, -0.77 DPS) [dungeon]; Enchanter's Cowl (4322, -0.84 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 spell_power points (1.24 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.44 DPS) [quest]; Triune Amulet (7722, -0.68 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.68 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.5 spell_power points (1.80 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.24 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.6 spell_power points (1.70 DPS) | yes | Guardian Cloak (5965, -0.65 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.82 DPS) [vendor]; Long Silken Cloak (4326, -1.62 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.4 spell_power points (2.88 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.39 DPS) [crafted]; Elemental Raiment (9434, -0.59 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (0.98 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.22 DPS) [quest]; Windchaser Cuffs (14429, -0.26 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.9 spell_power points (2.28 DPS) | yes | Red Mageweave Gloves (10018, -0.28 DPS) [crafted]; Black Mageweave Gloves (10003, -0.65 DPS) [crafted]; Gilded Handwraps (254021, -0.85 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.0 spell_power points (1.96 DPS) | yes | Gilded Cord (254037, -0.45 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.52 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.8 spell_power points (2.48 DPS) | yes | Crimson Silk Pantaloons (7062, -0.57 DPS) [crafted]; Abomination Skin Leggings (23173, -0.86 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.08 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.61 DPS) | yes | Gilded Slippers (254001, -0.79 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.43 DPS) [dungeon]; Spidersilk Boots (4320, -1.53 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.57 DPS) | yes | Ring of Forlorn Spirits (2043, -0.70 DPS) [quest]; Reedknot Ring (9622, -0.81 DPS) [quest]; Minor Channeling Ring (1449, -0.86 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.98 DPS) | yes | Reedknot Ring (9622, -0.22 DPS) [quest]; Minor Channeling Ring (1449, -0.28 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.46 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (78.4 DPS) | yes | Windweaver Staff (7757, -0.98 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.20 DPS) [dungeon]; Gut Ripper (2164, -3.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 365.8 spell_power points (39.86 DPS) | yes | Nether Force Wand (11263, -1.58 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.43 DPS) [quest]; Ragefire Wand (7513, -2.48 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 142.6. Weights run: 0.9s. Verify run: 0.7s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.821 ± 0.039, crit=0.297 ± 0.025 per rating point (14 rating = 1%, 4.164 per %), hit=0.605 ± 0.007 per rating point (10 rating = 1%, 6.046 per %), spell_haste=not significant (-2.454 ± 0.786), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.4 spell_power points (5.02 DPS) | yes | Dreamweave Circlet (10041, -0.88 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.05 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.19 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.9 spell_power points (2.97 DPS) | yes | Horizon Choker (13085, -1.34 DPS) [world_drop]; Mindburst Medallion (11196, -1.42 DPS) [quest]; Scorn's Icy Choker (23169, -2.90 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.8 spell_power points (3.94 DPS) | yes | Kentic Amice (11624, -0.44 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.17 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.20 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.9 spell_power points (2.68 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.36 DPS) [dungeon]; Runecloth Cloak (13860, -0.48 DPS) [crafted]; Big Voodoo Cloak (8216, -0.93 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.4 spell_power points (5.02 DPS) | yes | Runecloth Tunic (13857, -1.33 DPS) [crafted]; Dreamweave Vest (10021, -1.42 DPS) [crafted]; Robe of the Magi (1716, -2.79 DPS, sim-verified) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.7 spell_power points (1.81 DPS) | yes | Aristocratic Cuffs (12546, -0.06 DPS) [dungeon]; Shizzle's Nozzle Wiper (11917, -0.41 DPS) [quest]; Bloodband Bracers (11469, -2.13 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 29.5 spell_power points (4.19 DPS) | yes | Dreamweave Gloves (10019, -1.17 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.30 DPS) [vendor]; Raider Handwraps (272098, -1.79 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (142.6 DPS) | yes | Dawnspire Cord (12466, -0.01 DPS) [dungeon]; Deathmage Sash (10771, -0.33 DPS) [dungeon]; Satyrmane Sash (17755, -2.64 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.2 spell_power points (4.42 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Wizardweave Leggings (14132, -1.73 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.20 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.40 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.36 DPS) [vendor]; Gilded Sandals (254107, -0.80 DPS) [crafted]; Southsea Mojo Boots (20641, -0.99 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.9 spell_power points (2.11 DPS) | yes | Band of the Unicorn (7553, -0.27 DPS) [world_drop]; Brainlash (6440, -0.37 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.41 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.7 spell_power points (2.09 DPS) | yes | Brainlash (6440, -0.34 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.39 DPS) [rep]; Band of the Unicorn (7553, -1.69 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, -0.54 DPS) [world]; Spellshifter Rod (9527, -0.70 DPS) [quest]; Blade of Eternal Darkness (17780, -2.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 370.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.79 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 331.2. Weights run: 1.0s. Verify run: 0.7s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.785 ± 0.057, crit=0.506 ± 0.037 per rating point (14 rating = 1%, 7.087 per %), hit=0.851 ± 0.011 per rating point (10 rating = 1%, 8.506 per %), spell_haste=not significant (1.528 ± 1.117), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.7 spell_power points (10.36 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.19 DPS) [pvp]; Crimson Felt Hat (18727, -2.48 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (331.2 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.60 DPS) [quest]; Chains of the Lich (23125, -0.70 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.9 spell_power points (9.96 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.97 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.75 DPS) [vendor]; Mantle of the Timbermaw (19050, -9.87 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.8 spell_power points (6.69 DPS) | yes | Hide of the Wild (18510, -1.94 DPS) [crafted]; Spritecaster Cape (11623, -2.62 DPS) [dungeon]; Crystalline Threaded Cape (20697, -5.91 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.5 spell_power points (12.27 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.67 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.27 DPS) [pvp]; Robe of Everlasting Night (18385, -4.19 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.3 spell_power points (6.14 DPS) | yes | Sublime Wristguards (18497, -1.83 DPS) [dungeon]; Runecloth Cuffs (254123, -2.05 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.90 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 31.5 spell_power points (6.84 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.8 spell_power points (11.03 DPS) | yes | Magician's Cord (272393, -2.81 DPS) [vendor]; Stormpike Cloth Girdle (19094, -5.41 DPS) [rep]; Belt of the Archmage (18405, -8.77 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.0 spell_power points (11.73 DPS) | yes | Marshal's Silk Leggings (231605, -0.27 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -2.73 DPS) [pvp]; Skyshroud Leggings (13170, -2.99 DPS) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 33.6 spell_power points (7.29 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.65 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (331.2 DPS) | yes | Songstone of Ironforge (12543, -1.55 DPS) [quest]; Maiden's Circle (13001, -1.55 DPS) [world_drop]; Naglering (11669, -16.97 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (331.2 DPS) | yes | Songstone of Ironforge (12543, -0.65 DPS) [quest]; Maiden's Circle (13001, -0.65 DPS) [world_drop]; Naglering (11669, -16.29 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (331.2 DPS) | yes | Weakness Analyzer (272438, -1.52 DPS) [vendor]; Serenity Field (272439, -3.26 DPS) [vendor]; Burst of Knowledge (11832, -3.69 DPS) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (331.2 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (331.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.77 DPS) [world]; Teebu's Blazing Longsword (1728, -17.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 359.4 spell_power points (78.06 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.82 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.36 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.11 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 25.8. Weights run: 0.8s. Verify run: 0.5s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.404 ± 0.011, crit=0.118 ± 0.004 per rating point (14 rating = 1%, 1.646 per %), hit=0.358 ± 0.002 per rating point (10 rating = 1%, 3.575 per %), spell_haste=not significant (0.485 ± 0.160), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.43 DPS) | yes | Shadow Goggles (4373, -0.54 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 spell_power points (0.62 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.33 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.34 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.29 DPS) | yes | Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Black Whelp Cloak (7283, -0.07 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.27 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 spell_power points (0.50 DPS) | yes | Green Woolen Robe (6243, -0.20 DPS) [crafted]; Green Woolen Vest (2582, -0.22 DPS) [crafted]; Gray Woolen Robe (2585, -0.64 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 spell_power points (0.17 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest]; Owlbeard Bracers (16981, -0.04 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.50 DPS) | yes | Gnoll Casting Gloves (892, -0.07 DPS) [world]; Pristine Gloves (253913, -0.13 DPS) [crafted]; Apothecary Gloves (10919, -0.21 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 spell_power points (0.40 DPS) | yes | Keller's Girdle (2911, -0.17 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.17 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.39 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 spell_power points (0.87 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.37 DPS) [dungeon]; Rumpled Kilt (274741, -0.52 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 spell_power points (0.61 DPS) | yes | Pristine Boots (253889, -0.31 DPS) [crafted]; Red Woolen Boots (4313, -0.33 DPS) [crafted]; Feather Padded Treads (285345, -0.55 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.36 DPS) | yes | Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Loop of Sacrifice (281673, -0.21 DPS) [quest]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.21 DPS) | yes | Lavishly Jeweled Ring (1156, -0.04 DPS) [dungeon]; Loop of Sacrifice (281673, -0.07 DPS) [quest]; Volcanic Rock Ring (12053, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 4.0 spell_power points (0.29 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.12 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 316.1 spell_power points (22.52 DPS) | yes | Skycaller (12984, -0.72 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.23 DPS) [dungeon]; Sizzle Stick (8071, -4.52 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 42.2. Weights run: 0.9s. Verify run: 0.6s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.560 ± 0.021, crit=0.173 ± 0.009 per rating point (14 rating = 1%, 2.420 per %), hit=0.334 ± 0.003 per rating point (10 rating = 1%, 3.340 per %), spell_haste=not significant (0.131 ± 0.295), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 11.6 spell_power points (1.04 DPS) | yes | Holy Shroud (2721, -0.05 DPS) [world_drop]; Silk Headband (7050, -0.23 DPS) [crafted]; Embalmed Shroud (7691, -0.32 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (0.93 DPS) | yes | Crystal Starfire Medallion (5003, -0.73 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.73 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.97 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.26 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.27 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.45 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Cloak of Rot (4462, -0.05 DPS) [world]; Darkspear Raider's Cloak (272078, -0.05 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (1.46 DPS) | yes | Tree Bark Jacket (1486, -0.29 DPS) [dungeon]; Pristine Gown (253961, -0.48 DPS) [crafted]; Death Speaker Robes (6682, -0.50 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Nightsky Wristbands (6407, -0.51 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.51 DPS) [quest]; Glowing Magical Bracelets (13106, -0.67 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.8 spell_power points (0.79 DPS) | yes | Truefaith Gloves (7049, -0.19 DPS) [crafted]; Gnoll Casting Gloves (892, -0.25 DPS) [world]; Serpent Gloves (5970, -0.71 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.7 spell_power points (1.14 DPS) | yes | Belt of Arugal (6392, -0.18 DPS) [dungeon]; Crimson Silk Belt (7055, -0.25 DPS) [crafted]; Warsong Sash (16975, -0.45 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.5 spell_power points (1.21 DPS) | yes | Gaze Dreamer Pants (6903, -0.13 DPS) [dungeon]; Pristine Leggings (253987, -0.23 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.37 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (0.98 DPS) | yes | Spidersilk Boots (4320, -0.15 DPS) [crafted]; Boots of the Enchanter (4325, -0.53 DPS) [crafted]; Acidic Walkers (9454, -0.92 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.63 DPS) | yes | Black Widow Band (6199, -0.28 DPS) [world]; Snake Hoop (6750, -0.28 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.54 DPS) | yes | Snake Hoop (6750, -0.19 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Black Widow Band (6199, -0.35 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Twisted Chanter's Staff (890, -0.31 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.31 DPS) [quest]; Glimmering Staff (249392, -0.32 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.4 spell_power points (0.93 DPS) | yes | Orb of Souls (249395, -0.57 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -0.57 DPS) [world]; Tome of the Darkspear Prophecy (272090, -0.82 DPS, sim-verified) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 372.3 spell_power points (33.48 DPS) | yes | Starfaller (13063, -0.42 DPS) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552100130103050-0000000000000000000)

Set DPS (verified): 71.6. Weights run: 0.9s. Verify run: 0.6s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.733 ± 0.032, crit=0.235 ± 0.016 per rating point (14 rating = 1%, 3.294 per %), hit=0.439 ± 0.004 per rating point (10 rating = 1%, 4.394 per %), spell_haste=not significant (-0.427 ± 0.532), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.29 DPS) | yes | Augural Shroud (2620, -0.29 DPS) [world]; Corpseshroud (10574, -0.77 DPS) [dungeon]; Enchanter's Cowl (4322, -0.84 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 spell_power points (1.24 DPS) | yes | Triune Amulet (7722, -0.68 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.68 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -0.73 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.5 spell_power points (1.80 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.24 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.6 spell_power points (1.70 DPS) | yes | Guardian Cloak (5965, -0.65 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.82 DPS) [vendor]; Long Silken Cloak (4326, -0.97 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.4 spell_power points (2.88 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.39 DPS) [crafted]; Elemental Raiment (9434, -0.59 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.9 spell_power points (1.08 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon]; Condor Bracers (15864, -0.31 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.9 spell_power points (2.28 DPS) | yes | Black Mageweave Gloves (10003, -0.65 DPS) [crafted]; Red Mageweave Gloves (10018, -0.71 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -0.85 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.0 spell_power points (1.96 DPS) | yes | Gilded Cord (254037, -0.45 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.52 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.8 spell_power points (2.48 DPS) | yes | Abomination Skin Leggings (23173, -0.86 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.91 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.08 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.61 DPS) | yes | Gilded Slippers (254001, -1.29 DPS) [crafted]; Acidic Walkers (9454, -1.43 DPS) [dungeon]; Spidersilk Boots (4320, -1.53 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.57 DPS) | yes | Reedknot Ring (9622, -0.81 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.92 DPS) [vendor]; Black Widow Band (6199, -1.01 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.98 DPS) | yes | Reedknot Ring (9622, -0.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor]; Black Widow Band (6199, -0.42 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (71.6 DPS) | yes | Windweaver Staff (7757, -0.98 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.20 DPS) [dungeon]; Gut Ripper (2164, -3.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 365.8 spell_power points (39.86 DPS) | yes | Nether Force Wand (11263, -1.44 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.43 DPS) [quest]; Ragefire Wand (7513, -2.48 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 129.0. Weights run: 0.9s. Verify run: 0.7s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.821 ± 0.039, crit=0.297 ± 0.025 per rating point (14 rating = 1%, 4.164 per %), hit=0.605 ± 0.007 per rating point (10 rating = 1%, 6.046 per %), spell_haste=not significant (-2.454 ± 0.786), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.4 spell_power points (5.02 DPS) | yes | Dreamweave Circlet (10041, -0.88 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.05 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.19 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.9 spell_power points (2.97 DPS) | yes | Horizon Choker (13085, -1.34 DPS) [world_drop]; Mindburst Medallion (11196, -1.42 DPS) [quest]; Scorn's Icy Choker (23169, -2.59 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.8 spell_power points (3.94 DPS) | yes | Kentic Amice (11624, -0.44 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.17 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.20 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.4 spell_power points (2.75 DPS) | yes | Spritecaster Cape (11623, -0.07 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.43 DPS) [dungeon]; Runecloth Cloak (13860, -0.54 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.4 spell_power points (5.02 DPS) | yes | Robe of the Magi (1716, -1.20 DPS) [world_drop]; Runecloth Tunic (13857, -1.33 DPS) [crafted]; Dreamweave Vest (10021, -1.42 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.7 spell_power points (1.81 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -1.75 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 29.5 spell_power points (4.19 DPS) | yes | Raider Handwraps (272098, -0.63 DPS) [vendor]; Dreamweave Gloves (10019, -1.17 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.30 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (129.0 DPS) | yes | Dawnspire Cord (12466, -0.01 DPS) [dungeon]; Deathmage Sash (10771, -0.33 DPS) [dungeon]; Satyrmane Sash (17755, -1.95 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.2 spell_power points (4.42 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Wizardweave Leggings (14132, -1.73 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.59 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.40 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.36 DPS) [vendor]; Gilded Sandals (254107, -0.80 DPS) [crafted]; Southsea Mojo Boots (20641, -0.99 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.9 spell_power points (2.11 DPS) | yes | Band of the Unicorn (7553, -0.27 DPS) [world_drop]; Brainlash (6440, -0.37 DPS) [dungeon]; Advisor's Ring (19519, -0.41 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.7 spell_power points (2.09 DPS) | yes | Band of the Unicorn (7553, -0.25 DPS) [world_drop]; Brainlash (6440, -0.34 DPS) [dungeon]; Advisor's Ring (19519, -0.39 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellforce Rod (1664, -0.54 DPS) [world]; Spellshifter Rod (9527, -0.70 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 370.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.79 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 315.8. Weights run: 1.0s. Verify run: 0.7s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.785 ± 0.057, crit=0.506 ± 0.037 per rating point (14 rating = 1%, 7.087 per %), hit=0.851 ± 0.011 per rating point (10 rating = 1%, 8.506 per %), spell_haste=not significant (1.528 ± 1.117), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.7 spell_power points (10.36 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.19 DPS) [pvp]; Crimson Felt Hat (18727, -5.31 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (315.8 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.60 DPS) [quest]; Chains of the Lich (23125, -0.70 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.9 spell_power points (9.96 DPS) | yes | Warlord's Silk Amice (231594, -1.97 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.75 DPS) [vendor]; Mantle of the Timbermaw (19050, -7.00 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.8 spell_power points (6.69 DPS) | yes | Hide of the Wild (18510, -1.94 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.55 DPS) [quest]; Crystalline Threaded Cape (20697, -6.28 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.5 spell_power points (12.27 DPS) | yes | Warlord's Silk Raiment (231596, -0.67 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.27 DPS) [pvp]; Robe of Everlasting Night (18385, -6.88 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.3 spell_power points (6.14 DPS) | yes | Sublime Wristguards (18497, -1.83 DPS) [dungeon]; Runecloth Cuffs (254123, -2.05 DPS) [crafted]; General's Silk Cuffs (16538, -2.90 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 31.5 spell_power points (6.84 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.8 spell_power points (11.03 DPS) | yes | Magician's Cord (272393, -2.81 DPS) [vendor]; Frostwolf Cloth Belt (19090, -5.41 DPS) [rep]; Belt of the Archmage (18405, -8.47 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.0 spell_power points (11.73 DPS) | yes | General's Silk Trousers (231595, -0.27 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -2.73 DPS) [pvp]; Outrider's Silk Leggings (22747, -5.70 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 33.6 spell_power points (7.29 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.65 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (315.8 DPS) | yes | Eye of Orgrimmar (12545, -1.55 DPS) [quest]; Maiden's Circle (13001, -1.55 DPS) [world_drop]; Naglering (11669, -15.43 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (315.8 DPS) | yes | Eye of Orgrimmar (12545, -0.65 DPS) [quest]; Maiden's Circle (13001, -0.65 DPS) [world_drop]; Naglering (11669, -14.95 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (315.8 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.94 DPS) [crafted] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (315.8 DPS) | yes | Weakness Analyzer (272438, -1.52 DPS) [vendor]; Serenity Field (272439, -3.26 DPS) [vendor]; Blackhand's Breadth (13965, -5.61 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (315.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.51 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -20.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 359.4 spell_power points (78.06 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.82 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.36 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.11 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

