# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 32.6. Weights run: 0.7s. Verify run: 0.6s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.263 ± 0.009, crit=0.098 ± 0.002 per rating point (14 rating = 1%, 1.368 per %), hit=0.291 ± 0.002 per rating point (10 rating = 1%, 2.913 per %), spell_haste=not significant (0.543 ± 0.142), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.536 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.54 DPS) | yes | Shadow Goggles (4373, -0.77 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (0.66 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.30 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.42 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Caretaker's Cape (20428, -0.09 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.3 spell_power points (0.57 DPS) | yes | Green Woolen Vest (2582, -0.21 DPS) [crafted]; Bloody Apron (6226, -0.21 DPS) [dungeon]; Gray Woolen Robe (2585, -0.63 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.3 spell_power points (0.12 DPS) | yes | Windsong Bangles (263336, -0.03 DPS) [quest]; Repurposed Hair Band (281256, -0.07 DPS) [quest]; Bright Bracers (3647, -0.44 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.63 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.20 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.1 spell_power points (0.46 DPS) | yes | Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.28 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (32.6 DPS) | yes | Silk-threaded Trousers (1929, -0.05 DPS) [dungeon]; Rumpled Kilt (274741, -0.23 DPS) [vendor]; Abomination Skin Leggings (23173, -0.33 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (0.73 DPS) | yes | Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.38 DPS) [crafted]; Feather Padded Treads (285345, -0.60 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (0.50 DPS) | yes | Sludge-Stained Band (286535, -0.23 DPS) [world]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.45 DPS) | yes | Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Sludge-Stained Band (286535, -0.69 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.6 spell_power points (0.24 DPS) | yes | Lesser Staff of the Spire (1300, -0.09 DPS) [world]; Staff of Westfall (2042, -0.12 DPS) [quest]; Channeler's Staff (4437, -0.39 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 250.6 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.52 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 55.2. Weights run: 0.8s. Verify run: 0.6s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.448 ± 0.015, crit=0.101 ± 0.003 per rating point (14 rating = 1%, 1.411 per %), hit=0.292 ± 0.003 per rating point (10 rating = 1%, 2.924 per %), spell_haste=not significant (0.334 ± 0.194), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.653 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.35 DPS) | yes | Silk Headband (7050, -0.24 DPS) [crafted]; Embalmed Shroud (7691, -0.37 DPS) [dungeon]; Enchanter's Cowl (4322, -0.65 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.7 spell_power points (1.18 DPS) | yes | Crystal Starfire Medallion (5003, -0.97 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.97 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.25 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (1.59 DPS) | yes | Death Speaker Mantle (6685, -0.26 DPS) [dungeon]; Fairywing Mantle (9536, -0.37 DPS) [quest]; Invoker's Mantle (215365, -0.46 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.61 DPS) | yes | Repairman's Cape (9605, -0.03 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Prelacy Cape (7004, -0.12 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.8 spell_power points (1.81 DPS) | yes | Death Speaker Robes (6682, -0.35 DPS) [dungeon]; Tree Bark Jacket (1486, -0.52 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.57 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.10 DPS) | yes | Nightsky Wristbands (6407, -0.77 DPS) [world_drop]; Stonecloth Bindings (14416, -0.83 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.44 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.9 spell_power points (1.09 DPS) | yes | Serpent Gloves (5970, -0.24 DPS) [dungeon]; Shilly Mitts (9609, -0.24 DPS) [quest]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.3 spell_power points (1.51 DPS) | yes | Belt of Arugal (6392, -0.24 DPS) [dungeon]; Invoker's Cord (215366, -0.38 DPS) [crafted]; Crimson Silk Belt (7055, -0.39 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.6 spell_power points (1.54 DPS) | yes | Gaze Dreamer Pants (6903, -0.07 DPS) [dungeon]; Pristine Leggings (253987, -0.30 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.1 spell_power points (1.24 DPS) | yes | Acidic Walkers (9454, -0.19 DPS) [dungeon]; Nimbus Boots (6998, -0.51 DPS) [quest]; Spidersilk Boots (4320, -1.58 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.86 DPS) | yes | Minor Channeling Ring (1449, -0.14 DPS) [quest]; Black Widow Band (6199, -0.47 DPS) [world]; Snake Hoop (6750, -0.47 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.73 DPS) | yes | Black Widow Band (6199, -0.35 DPS) [world]; Snake Hoop (6750, -0.35 DPS) [quest]; Minor Channeling Ring (1449, -0.86 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.10 DPS) | yes | Twisted Chanter's Staff (890, -0.55 DPS) [world_drop]; Channeler's Staff (4437, -0.66 DPS) [world]; Glimmering Staff (249392, -0.93 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.7 spell_power points (1.18 DPS) | yes | Eye of Paleth (2943, -0.70 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.70 DPS) [world]; Dwarven Tome (279898, -0.87 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 274.6 spell_power points (33.58 DPS) | yes | Starfaller (13063, -0.69 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.97 DPS) [crafted]; Gravestone Scepter (7001, -4.58 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 253225113100011400-00000000000000000-0000000000000000000)

Set DPS (verified): 177.3. Weights run: 1.0s. Verify run: 0.6s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.677 ± 0.030, crit=0.207 ± 0.005 per rating point (14 rating = 1%, 2.897 per %), hit=0.379 ± 0.006 per rating point (10 rating = 1%, 3.785 per %), spell_haste=not significant (1.420 ± 0.468), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.900 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.63 DPS) | yes | Living Cowl (5608, -2.14 DPS) [world]; Corpseshroud (10574, -2.18 DPS) [dungeon]; Augural Shroud (2620, -2.56 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.1 spell_power points (2.96 DPS) | yes | Triune Amulet (7722, -1.69 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.69 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.22 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.8 spell_power points (4.23 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.54 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.1 spell_power points (4.04 DPS) | yes | Guardian Cloak (5965, -1.53 DPS) [crafted]; Darkspear Raider's Cloak (272077, -2.05 DPS) [vendor]; Long Silken Cloak (4326, -2.51 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.1 spell_power points (6.98 DPS) | yes | Dreamweave Vest (10021, -0.53 DPS) [crafted]; Robe of Power (7054, -1.05 DPS) [crafted]; Elemental Raiment (9434, -1.36 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.41 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.54 DPS) [quest]; Windchaser Cuffs (14429, -0.78 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 spell_power points (5.55 DPS) | yes | Black Mageweave Gloves (10003, -1.53 DPS) [crafted]; Red Mageweave Gloves (10018, -2.00 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.13 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.2 spell_power points (4.60 DPS) | yes | Highlander's Cloth Girdle (20098, -0.12 DPS) [rep]; Gilded Cord (254037, -1.00 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.1 spell_power points (5.93 DPS) | yes | Crimson Silk Pantaloons (7062, -1.80 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.07 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.60 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.43 DPS) | yes | Gilded Slippers (254001, -2.56 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.64 DPS) [dungeon]; Spidersilk Boots (4320, -3.83 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (3.77 DPS) | yes | Ring of Forlorn Spirits (2043, -1.62 DPS) [quest]; Reedknot Ring (9622, -1.89 DPS) [quest]; Minor Channeling Ring (1449, -2.07 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.41 DPS) | yes | Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.54 DPS) [quest]; Minor Channeling Ring (1449, -0.71 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (177.3 DPS) | yes | Windweaver Staff (7757, -2.64 DPS) [dungeon]; Scorn's Focal Dagger (23168, -2.95 DPS) [dungeon]; Gut Ripper (2164, -8.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 151.0 spell_power points (40.47 DPS) | yes | Nether Force Wand (11263, -2.59 DPS) [quest]; Icefury Wand (7514, -2.73 DPS) [quest]; Ragefire Wand (7513, -2.78 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 253225113100011531-03200000000000000-0000000000000000000)

Set DPS (verified): 292.8. Weights run: 1.0s. Verify run: 0.8s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.804 ± 0.043, crit=0.333 ± 0.008 per rating point (14 rating = 1%, 4.669 per %), hit=0.559 ± 0.009 per rating point (10 rating = 1%, 5.593 per %), spell_haste=not significant (1.367 ± 0.663), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.908 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.1 spell_power points (10.36 DPS) | yes | Dreamweave Circlet (10041, -1.78 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.39 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.8 spell_power points (6.15 DPS) | yes | Horizon Choker (13085, -2.82 DPS) [world_drop]; Mindburst Medallion (11196, -2.95 DPS) [quest]; Scorn's Icy Choker (23169, -5.54 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.5 spell_power points (8.11 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.23 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.48 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.8 spell_power points (5.56 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.76 DPS) [dungeon]; Runecloth Cloak (13860, -1.00 DPS) [crafted]; Big Voodoo Cloak (8216, -1.94 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.1 spell_power points (10.36 DPS) | yes | Robe of the Magi (1716, -2.44 DPS) [world_drop]; Runecloth Tunic (13857, -2.73 DPS) [crafted]; Knight's Dreadweave Vest (220886, -2.82 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.6 spell_power points (3.73 DPS) | yes | Aristocratic Cuffs (12546, -0.17 DPS) [dungeon]; Shizzle's Nozzle Wiper (11917, -0.88 DPS) [quest]; Bloodband Bracers (11469, -2.92 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.9 spell_power points (8.52 DPS) | yes | Dreamweave Gloves (10019, -2.25 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.54 DPS) [vendor]; Raider Handwraps (272098, -3.25 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (292.8 DPS) | yes | Dawnspire Cord (12466, -0.04 DPS) [dungeon]; Deathmage Sash (10771, -0.69 DPS) [dungeon]; Satyrmane Sash (17755, -3.43 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.0 spell_power points (9.16 DPS) | yes | Red Mageweave Pants (10009, -2.18 DPS) [crafted]; Wizardweave Leggings (14132, -3.55 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -4.48 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.09 DPS) | yes | Gilded Sandals (254107, -1.70 DPS) [crafted]; Southsea Mojo Boots (20641, -2.11 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -2.83 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (4.38 DPS) | yes | Band of the Unicorn (7553, -0.54 DPS) [world_drop]; Brainlash (6440, -0.82 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.83 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.6 spell_power points (4.32 DPS) | yes | Band of the Unicorn (7553, -0.48 DPS) [world_drop]; Brainlash (6440, -0.76 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.78 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -9.48 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 177.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.39 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 253225113100011531-03202300000000000-0050000000000000000)

Set DPS (verified): 470.2. Weights run: 1.1s. Verify run: 0.7s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.960 ± 0.057, crit=0.520 ± 0.011 per rating point (14 rating = 1%, 7.279 per %), hit=0.846 ± 0.012 per rating point (10 rating = 1%, 8.457 per %), spell_haste=not significant (2.338 ± 0.894), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.904 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 52.3 spell_power points (15.57 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -2.00 DPS) [pvp]; Magister's Crown (16686, -6.37 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (470.2 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.88 DPS) [quest]; Chains of the Lich (23125, -1.63 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.7 spell_power points (14.50 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.76 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.55 DPS) [crafted]; Darkspear Shoulderpads (272103, -8.36 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.1 spell_power points (9.57 DPS) | yes | Hide of the Wild (18510, -2.54 DPS) [crafted]; Spritecaster Cape (11623, -3.69 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.07 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.8 spell_power points (17.52 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.66 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.23 DPS) [pvp]; Sorcerer's Robes (226932, -11.21 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.7 spell_power points (8.84 DPS) | yes | Sublime Wristguards (18497, -2.41 DPS) [dungeon]; Runecloth Cuffs (254123, -2.70 DPS) [crafted]; Marshal's Silk Bracers (16438, -3.41 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 33.9 spell_power points (10.10 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 54.6 spell_power points (16.26 DPS) | yes | Magician's Cord (272393, -3.73 DPS) [vendor]; Belt of the Archmage (18405, -6.66 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -8.04 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 56.2 spell_power points (16.73 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.44 DPS) [pvp]; Sorcerer's Leggings (226933, -3.97 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 36.4 spell_power points (10.83 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.89 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (470.2 DPS) | yes | Songstone of Ironforge (12543, -2.34 DPS) [quest]; Maiden's Circle (13001, -2.34 DPS) [world_drop]; Naglering (11669, -9.77 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (470.2 DPS) | yes | Songstone of Ironforge (12543, -0.89 DPS) [quest]; Maiden's Circle (13001, -0.89 DPS) [world_drop]; Naglering (11669, -9.70 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (470.2 DPS) | yes | Weakness Analyzer (272438, -2.09 DPS) [vendor]; Serenity Field (272439, -4.47 DPS) [vendor]; Burst of Knowledge (11832, -5.06 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (470.2 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -2.80 DPS, sim-verified) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (470.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.89 DPS) [world]; Teebu's Blazing Longsword (1728, -12.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 263.7 spell_power points (78.55 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.42 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.50 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.15 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 29.5. Weights run: 0.7s. Verify run: 0.6s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.263 ± 0.009, crit=0.098 ± 0.002 per rating point (14 rating = 1%, 1.368 per %), hit=0.291 ± 0.002 per rating point (10 rating = 1%, 2.913 per %), spell_haste=not significant (0.543 ± 0.142), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.536 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.54 DPS) | yes | Shadow Goggles (4373, -0.72 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (0.66 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.12 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.30 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.26 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.3 spell_power points (0.57 DPS) | yes | Green Woolen Vest (2582, -0.21 DPS) [crafted]; Bloody Apron (6226, -0.21 DPS) [dungeon]; Gray Woolen Robe (2585, -0.28 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.6 spell_power points (0.14 DPS) | yes | Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Owlbeard Bracers (16981, -0.17 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.63 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.20 DPS) [crafted]; Apothecary Gloves (10919, -0.27 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.1 spell_power points (0.46 DPS) | yes | Novice Arcanist's Sash (253885, -0.02 DPS) [crafted]; Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (29.5 DPS) | yes | Silk-threaded Trousers (1929, -0.05 DPS) [dungeon]; Rumpled Kilt (274741, -0.23 DPS) [vendor]; Abomination Skin Leggings (23173, -0.30 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (0.73 DPS) | yes | Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.38 DPS) [crafted]; Feather Padded Treads (285345, -0.46 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.45 DPS) | yes | Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; Loop of Sacrifice (281673, -0.33 DPS) [quest]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.27 DPS) | yes | Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Loop of Sacrifice (281673, -0.15 DPS) [quest]; Volcanic Rock Ring (12053, -0.20 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.6 spell_power points (0.24 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.05 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 250.6 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.27 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 49.4. Weights run: 0.8s. Verify run: 0.6s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.448 ± 0.015, crit=0.101 ± 0.003 per rating point (14 rating = 1%, 1.411 per %), hit=0.292 ± 0.003 per rating point (10 rating = 1%, 2.924 per %), spell_haste=not significant (0.334 ± 0.194), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.653 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.35 DPS) | yes | Enchanter's Cowl (4322, -0.06 DPS) [crafted]; Silk Headband (7050, -0.24 DPS) [crafted]; Embalmed Shroud (7691, -0.37 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.7 spell_power points (1.18 DPS) | yes | Darkspear Warding Pendant (272075, -0.52 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.97 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.97 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (1.59 DPS) | yes | Death Speaker Mantle (6685, -0.26 DPS) [dungeon]; Fairywing Mantle (9536, -0.37 DPS) [quest]; Invoker's Mantle (215365, -0.46 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.61 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Battle Healer's Cloak (19529, -0.12 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.8 spell_power points (1.81 DPS) | yes | Tree Bark Jacket (1486, -0.22 DPS) [dungeon]; Death Speaker Robes (6682, -0.35 DPS) [dungeon]; Pristine Gown (253961, -0.57 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.10 DPS) | yes | Nightsky Wristbands (6407, -0.77 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.77 DPS) [quest]; Glowing Magical Bracelets (13106, -1.39 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.2 spell_power points (1.01 DPS) | yes | Truefaith Gloves (7049, -0.23 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS) [world]; Serpent Gloves (5970, -0.43 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.3 spell_power points (1.51 DPS) | yes | Belt of Arugal (6392, -0.24 DPS) [dungeon]; Invoker's Cord (215366, -0.38 DPS) [crafted]; Warsong Sash (16975, -0.49 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.6 spell_power points (1.54 DPS) | yes | Gaze Dreamer Pants (6903, -0.07 DPS) [dungeon]; Pristine Leggings (253987, -0.30 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.1 spell_power points (1.24 DPS) | yes | Acidic Walkers (9454, -0.19 DPS) [dungeon]; Boots of the Enchanter (4325, -0.63 DPS) [crafted]; Spidersilk Boots (4320, -1.33 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.86 DPS) | yes | Black Widow Band (6199, -0.47 DPS) [world]; Snake Hoop (6750, -0.47 DPS) [quest]; Sludge-Stained Band (286535, -0.49 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.73 DPS) | yes | Snake Hoop (6750, -0.35 DPS) [quest]; Sludge-Stained Band (286535, -0.37 DPS) [world]; Black Widow Band (6199, -1.22 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.10 DPS) | yes | Twisted Chanter's Staff (890, -0.55 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.55 DPS) [quest]; Glimmering Staff (249392, -0.85 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.7 spell_power points (1.18 DPS) | yes | Orb of Souls (249395, -0.70 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.72 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -0.99 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 274.6 spell_power points (33.58 DPS) | yes | Starfaller (13063, -0.69 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.97 DPS) [crafted]; Gravestone Scepter (7001, -4.58 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 253225113100011400-00000000000000000-0000000000000000000)

Set DPS (verified): 171.1. Weights run: 1.0s. Verify run: 0.6s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.677 ± 0.030, crit=0.207 ± 0.005 per rating point (14 rating = 1%, 2.897 per %), hit=0.379 ± 0.006 per rating point (10 rating = 1%, 3.785 per %), spell_haste=not significant (1.420 ± 0.468), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.900 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.63 DPS) | yes | Living Cowl (5608, -2.14 DPS) [world]; Corpseshroud (10574, -2.18 DPS) [dungeon]; Augural Shroud (2620, -2.71 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.1 spell_power points (2.96 DPS) | yes | Triune Amulet (7722, -1.69 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.69 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.24 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.8 spell_power points (4.23 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.54 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.1 spell_power points (4.04 DPS) | yes | Guardian Cloak (5965, -1.53 DPS) [crafted]; Darkspear Raider's Cloak (272077, -2.05 DPS) [vendor]; Long Silken Cloak (4326, -2.18 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.1 spell_power points (6.98 DPS) | yes | Dreamweave Vest (10021, -0.53 DPS) [crafted]; Robe of Power (7054, -1.05 DPS) [crafted]; Elemental Raiment (9434, -1.36 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.4 spell_power points (2.52 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.11 DPS) [dungeon]; Condor Bracers (15864, -0.65 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 spell_power points (5.55 DPS) | yes | Black Mageweave Gloves (10003, -1.53 DPS) [crafted]; Gilded Handwraps (254021, -2.13 DPS) [crafted]; Red Mageweave Gloves (10018, -2.34 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.2 spell_power points (4.60 DPS) | yes | Defiler's Cloth Girdle (20166, -0.12 DPS) [rep]; Gilded Cord (254037, -1.00 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.1 spell_power points (5.93 DPS) | yes | Crimson Silk Pantaloons (7062, -1.73 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.07 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.60 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.43 DPS) | yes | Gilded Slippers (254001, -3.10 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.64 DPS) [dungeon]; Spidersilk Boots (4320, -3.83 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (3.77 DPS) | yes | Reedknot Ring (9622, -1.89 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.16 DPS) [vendor]; Black Widow Band (6199, -2.50 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.41 DPS) | yes | Reedknot Ring (9622, -0.54 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.80 DPS) [vendor]; Black Widow Band (6199, -1.14 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (171.1 DPS) | yes | Windweaver Staff (7757, -2.64 DPS) [dungeon]; Scorn's Focal Dagger (23168, -2.95 DPS) [dungeon]; Gut Ripper (2164, -8.53 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 151.0 spell_power points (40.47 DPS) | yes | Nether Force Wand (11263, -2.59 DPS) [quest]; Icefury Wand (7514, -2.73 DPS) [quest]; Ragefire Wand (7513, -2.78 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253225113100011531-03200000000000000-0000000000000000000)

Set DPS (verified): 285.8. Weights run: 1.0s. Verify run: 0.8s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.804 ± 0.043, crit=0.333 ± 0.008 per rating point (14 rating = 1%, 4.669 per %), hit=0.559 ± 0.009 per rating point (10 rating = 1%, 5.593 per %), spell_haste=not significant (1.367 ± 0.663), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.908 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.1 spell_power points (10.36 DPS) | yes | Dreamweave Circlet (10041, -1.78 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.39 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.8 spell_power points (6.15 DPS) | yes | Horizon Choker (13085, -2.82 DPS) [world_drop]; Mindburst Medallion (11196, -2.95 DPS) [quest]; Scorn's Icy Choker (23169, -5.38 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.5 spell_power points (8.11 DPS) | yes | Kentic Amice (11624, -0.89 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.23 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.48 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.2 spell_power points (5.68 DPS) | yes | Spritecaster Cape (11623, -0.12 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.89 DPS) [dungeon]; Runecloth Cloak (13860, -1.12 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.1 spell_power points (10.36 DPS) | yes | Robe of the Magi (1716, -2.44 DPS) [world_drop]; Runecloth Tunic (13857, -2.73 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -2.82 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.6 spell_power points (3.73 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -3.18 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.9 spell_power points (8.52 DPS) | yes | Raider Handwraps (272098, -1.23 DPS) [vendor]; Dreamweave Gloves (10019, -2.25 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.54 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (285.8 DPS) | yes | Dawnspire Cord (12466, -0.04 DPS) [dungeon]; Deathmage Sash (10771, -0.69 DPS) [dungeon]; Satyrmane Sash (17755, -6.04 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.0 spell_power points (9.16 DPS) | yes | Red Mageweave Pants (10009, -2.18 DPS) [crafted]; Wizardweave Leggings (14132, -3.55 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -4.62 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.09 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.94 DPS) [vendor]; Gilded Sandals (254107, -1.70 DPS) [crafted]; Southsea Mojo Boots (20641, -2.11 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (4.38 DPS) | yes | Band of the Unicorn (7553, -0.54 DPS) [world_drop]; Brainlash (6440, -0.82 DPS) [dungeon]; Advisor's Ring (19519, -0.83 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.6 spell_power points (4.32 DPS) | yes | Band of the Unicorn (7553, -0.48 DPS) [world_drop]; Brainlash (6440, -0.76 DPS) [dungeon]; Advisor's Ring (19519, -0.78 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -7.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 177.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.39 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 253225113100011531-03202300000000000-0050000000000000000)

Set DPS (verified): 459.3. Weights run: 1.1s. Verify run: 0.7s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.960 ± 0.057, crit=0.520 ± 0.011 per rating point (14 rating = 1%, 7.279 per %), hit=0.846 ± 0.012 per rating point (10 rating = 1%, 8.457 per %), spell_haste=not significant (2.338 ± 0.894), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.904 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 52.3 spell_power points (15.57 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -2.00 DPS) [pvp]; Magister's Crown (16686, -6.83 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (459.3 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.88 DPS) [quest]; Chains of the Lich (23125, -1.63 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.7 spell_power points (14.50 DPS) | yes | Warlord's Silk Amice (231594, -2.76 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.55 DPS) [crafted]; Darkspear Shoulderpads (272103, -6.89 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.1 spell_power points (9.57 DPS) | yes | Crystalline Threaded Cape (20697, -2.47 DPS) [world]; Hide of the Wild (18510, -2.54 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.42 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.8 spell_power points (17.52 DPS) | yes | Warlord's Silk Raiment (231596, -0.66 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.23 DPS) [pvp]; Sorcerer's Robes (226932, -10.48 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.7 spell_power points (8.84 DPS) | yes | Sublime Wristguards (18497, -2.41 DPS) [dungeon]; Runecloth Cuffs (254123, -2.70 DPS) [crafted]; General's Silk Cuffs (16538, -3.41 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 33.9 spell_power points (10.10 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 54.6 spell_power points (16.26 DPS) | yes | Magician's Cord (272393, -3.73 DPS) [vendor]; Belt of the Archmage (18405, -5.56 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -8.04 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 56.2 spell_power points (16.73 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -3.44 DPS) [pvp]; Outrider's Silk Leggings (22747, -6.75 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 36.4 spell_power points (10.83 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.89 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (459.3 DPS) | yes | Eye of Orgrimmar (12545, -2.34 DPS) [quest]; Maiden's Circle (13001, -2.34 DPS) [world_drop]; Naglering (11669, -9.27 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (459.3 DPS) | yes | Eye of Orgrimmar (12545, -0.89 DPS) [quest]; Maiden's Circle (13001, -0.89 DPS) [world_drop]; Naglering (11669, -9.06 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (459.3 DPS) | yes | Weakness Analyzer (272438, -2.09 DPS) [vendor]; Serenity Field (272439, -4.47 DPS) [vendor]; Burst of Knowledge (11832, -5.06 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (459.3 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (459.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.65 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 263.7 spell_power points (78.55 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.42 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.50 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.15 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

