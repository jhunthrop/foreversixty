# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.5. Weights run: 0.7s. Verify run: 0.5s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.379 ± 0.007, crit=0.075 ± 0.002 per rating point (14 rating = 1%, 1.054 per %), hit=0.196 ± 0.001 per rating point (10 rating = 1%, 1.956 per %), spell_haste=0.781 ± 0.127, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.77 DPS) | yes | Shadow Goggles (4373, -1.42 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (1.08 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.57 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.51 DPS) | yes | Pearl-clasped Cloak (5542, -0.11 DPS) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.89 DPS) | yes | Green Woolen Robe (6243, -0.35 DPS) [crafted]; Green Woolen Vest (2582, -0.37 DPS) [crafted]; Gray Woolen Robe (2585, -1.02 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.9 spell_power points (0.24 DPS) | yes | Windsong Bangles (263336, -0.12 DPS) [quest]; Repurposed Hair Band (281256, -0.15 DPS) [quest]; Bright Bracers (3647, -0.59 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.90 DPS) | yes | Gnoll Casting Gloves (892, -0.13 DPS) [world]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.54 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.71 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.32 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.60 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.0 spell_power points (1.54 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.65 DPS) [dungeon]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (1.09 DPS) | yes | Pristine Boots (253889, -0.56 DPS) [crafted]; Red Woolen Boots (4313, -0.58 DPS) [crafted]; Feather Padded Treads (285345, -0.74 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.74 DPS) | yes | Sludge-Stained Band (286535, -0.35 DPS) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.59 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop]; Sludge-Stained Band (286535, -0.83 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.8 spell_power points (0.49 DPS) | yes | Lesser Staff of the Spire (1300, -0.19 DPS) [world]; Staff of Westfall (2042, -0.24 DPS) [quest]; Channeler's Staff (4437, -0.35 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 176.8 spell_power points (22.70 DPS) | yes | Skycaller (12984, -1.07 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.41 DPS) [dungeon]; Deepblaze (279896, -4.17 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 62.0. Weights run: 0.7s. Verify run: 0.5s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.475 ± 0.012, crit=0.130 ± 0.005 per rating point (14 rating = 1%, 1.820 per %), hit=0.226 ± 0.002 per rating point (10 rating = 1%, 2.261 per %), spell_haste=not significant (0.631 ± 0.199), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.65 DPS) | yes | Enchanter's Cowl (4322, -0.04 DPS) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.9 spell_power points (1.48 DPS) | yes | Crystal Starfire Medallion (5003, -1.19 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.19 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.60 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.3 spell_power points (1.99 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.59 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Prelacy Cape (7004, -0.15 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.2 spell_power points (2.28 DPS) | yes | Death Speaker Robes (6682, -0.44 DPS) [dungeon]; Pristine Gown (253961, -0.73 DPS) [crafted]; Tree Bark Jacket (1486, -0.96 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Nightsky Wristbands (6407, -0.92 DPS) [world_drop]; Stonecloth Bindings (14416, -0.99 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.29 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 9.2 spell_power points (1.39 DPS) | yes | Serpent Gloves (5970, -0.33 DPS) [dungeon]; Truefaith Gloves (7049, -0.42 DPS) [crafted]; Shilly Mitts (9609, -0.82 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.4 spell_power points (1.87 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Invoker's Cord (215366, -0.46 DPS) [crafted]; Crimson Silk Belt (7055, -0.47 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.8 spell_power points (1.92 DPS) | yes | Gaze Dreamer Pants (6903, -0.12 DPS) [dungeon]; Pristine Leggings (253987, -0.37 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.55 DPS) | yes | Acidic Walkers (9454, -0.23 DPS) [dungeon]; Nimbus Boots (6998, -0.65 DPS) [quest]; Spidersilk Boots (4320, -1.81 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.05 DPS) | yes | Minor Channeling Ring (1449, -0.16 DPS) [quest]; Black Widow Band (6199, -0.55 DPS) [world]; Snake Hoop (6750, -0.55 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.90 DPS) | yes | Black Widow Band (6199, -0.40 DPS) [world]; Snake Hoop (6750, -0.40 DPS) [quest]; Minor Channeling Ring (1449, -1.13 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Glimmering Staff (249392, -0.57 DPS) [crafted]; Twisted Chanter's Staff (890, -0.64 DPS) [world_drop]; Channeler's Staff (4437, -0.78 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.9 spell_power points (1.48 DPS) | yes | Eye of Paleth (2943, -0.88 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.88 DPS) [world]; Dwarven Tome (279898, -0.95 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 224.2 spell_power points (33.66 DPS) | yes | Starfaller (13063, -0.51 DPS) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 100000000000000000-00000000000000000-2535111300000301050)

Set DPS (verified): 99.5. Weights run: 0.7s. Verify run: 0.5s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.722 ± 0.024, crit=0.208 ± 0.009 per rating point (14 rating = 1%, 2.914 per %), hit=0.374 ± 0.004 per rating point (10 rating = 1%, 3.744 per %), spell_haste=1.758 ± 0.380, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.10 DPS) | yes | Augural Shroud (2620, -0.41 DPS) [world]; Corpseshroud (10574, -1.08 DPS) [dungeon]; Enchanter's Cowl (4322, -1.15 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.3 spell_power points (1.67 DPS) | yes | Triune Amulet (7722, -0.93 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.93 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.16 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.4 spell_power points (2.42 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.5 spell_power points (2.28 DPS) | yes | Guardian Cloak (5965, -0.87 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.11 DPS) [vendor]; Long Silken Cloak (4326, -1.20 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.3 spell_power points (3.88 DPS) | yes | Dreamweave Vest (10021, -0.27 DPS) [crafted]; Robe of Power (7054, -0.54 DPS) [crafted]; Elemental Raiment (9434, -0.79 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.33 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest]; Windchaser Cuffs (14429, -0.37 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.9 spell_power points (3.08 DPS) | yes | Black Mageweave Gloves (10003, -0.87 DPS) [crafted]; Gilded Handwraps (254021, -1.16 DPS) [crafted]; Red Mageweave Gloves (10018, -1.30 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.8 spell_power points (2.63 DPS) | yes | Highlander's Cloth Girdle (20098, -0.14 DPS) [rep]; Gilded Cord (254037, -0.60 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.7 spell_power points (3.34 DPS) | yes | Crimson Silk Pantaloons (7062, -1.08 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.16 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.46 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.54 DPS) | yes | Gilded Slippers (254001, -1.63 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.95 DPS) [dungeon]; Spidersilk Boots (4320, -2.08 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (2.11 DPS) | yes | Ring of Forlorn Spirits (2043, -0.93 DPS) [quest]; Reedknot Ring (9622, -1.08 DPS) [quest]; Minor Channeling Ring (1449, -1.16 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.33 DPS) | yes | Ring of Forlorn Spirits (2043, -0.15 DPS) [quest]; Reedknot Ring (9622, -0.29 DPS) [quest]; Minor Channeling Ring (1449, -0.38 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (99.5 DPS) | yes | Windweaver Staff (7757, -1.35 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Gut Ripper (2164, -4.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 271.4 spell_power points (40.02 DPS) | yes | Nether Force Wand (11263, -0.97 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.51 DPS) [quest]; Ragefire Wand (7513, -2.56 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 203005000100000000-00000000000000000-2535111300000301050)

Set DPS (verified): 201.9. Weights run: 0.7s. Verify run: 0.6s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.858 ± 0.037, crit=0.315 ± 0.013 per rating point (14 rating = 1%, 4.406 per %), hit=0.641 ± 0.008 per rating point (10 rating = 1%, 6.410 per %), spell_haste=12.188 ± 0.716, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.2 spell_power points (7.23 DPS) | yes | Dreamweave Circlet (10041, -1.32 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.49 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.83 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.1 spell_power points (4.23 DPS) | yes | Horizon Choker (13085, -1.83 DPS) [world_drop]; Mindburst Medallion (11196, -2.00 DPS) [quest]; Scorn's Icy Choker (23169, -4.25 DPS, sim-verified) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -1.00 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.06 DPS) [crafted]; Rotgrip Mantle (17732, -2.85 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.1 spell_power points (3.83 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.49 DPS) [dungeon]; Runecloth Cloak (13860, -0.66 DPS) [crafted]; Big Voodoo Cloak (8216, -1.29 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.2 spell_power points (7.23 DPS) | yes | Robe of the Magi (1716, -1.80 DPS) [world_drop]; Runecloth Tunic (13857, -1.94 DPS) [crafted]; Knight's Dreadweave Vest (220886, -2.06 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 13.0 spell_power points (2.60 DPS) | yes | Aristocratic Cuffs (12546, -0.03 DPS) [dungeon]; Bloodband Bracers (11469, -0.06 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.54 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.4 spell_power points (6.08 DPS) | yes | Raider Handwraps (272098, -0.91 DPS) [vendor]; Dreamweave Gloves (10019, -1.80 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.94 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.6 spell_power points (4.52 DPS) | yes | Dawnspire Cord (12466, -0.06 DPS) [dungeon]; Ban'thok Sash (11662, -0.10 DPS) [dungeon]; Deathmage Sash (10771, -0.54 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.6 spell_power points (6.32 DPS) | yes | Red Mageweave Pants (10009, -1.46 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.49 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.96 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.80 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.37 DPS) [vendor]; Gilded Sandals (254107, -1.06 DPS) [crafted]; Southsea Mojo Boots (20641, -1.31 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (3.03 DPS) | yes | Band of the Unicorn (7553, -0.43 DPS) [world_drop]; Brainlash (6440, -0.46 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.63 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.0 spell_power points (3.00 DPS) | yes | Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Brainlash (6440, -0.43 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.60 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellforce Rod (1664, -0.97 DPS) [world]; Spellshifter Rod (9527, -1.03 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.58 DPS) [world_drop]; Pyric Caduceus (11748, -3.21 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.65 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Noxious Shooter

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 203005000100000000-11302300000000000-2535111300000301050)

Set DPS (verified): 325.4. Weights run: 0.8s. Verify run: 0.6s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.993 ± 0.051, crit=0.496 ± 0.019 per rating point (14 rating = 1%, 6.948 per %), hit=1.032 ± 0.011 per rating point (10 rating = 1%, 10.317 per %), spell_haste=16.128 ± 1.020, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 52.8 spell_power points (9.92 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.31 DPS) [pvp]; Magister's Crown (16686, -2.25 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (325.4 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.56 DPS) [quest]; Chains of the Lich (23125, -1.11 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.8 spell_power points (9.18 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.68 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.25 DPS) [crafted]; Darkspear Shoulderpads (272103, -4.99 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 34.3 spell_power points (6.44 DPS) | yes | Crystalline Threaded Cape (20697, -1.93 DPS) [world]; Hide of the Wild (18510, -1.94 DPS) [crafted]; Spritecaster Cape (11623, -2.69 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.9 spell_power points (11.06 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.38 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -2.64 DPS) [pvp]; Sorcerer's Robes (226932, -5.92 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.9 spell_power points (5.63 DPS) | yes | Sublime Wristguards (18497, -1.51 DPS) [dungeon]; Runecloth Cuffs (254123, -1.69 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.08 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 36.2 spell_power points (6.81 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 57.2 spell_power points (10.74 DPS) | yes | Magician's Cord (272393, -2.69 DPS) [vendor]; Stormpike Cloth Girdle (19094, -5.49 DPS) [rep]; Belt of the Archmage (18405, -5.67 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.8 spell_power points (10.49 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -2.00 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -2.07 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 36.9 spell_power points (6.93 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.56 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (325.4 DPS) | yes | Songstone of Ironforge (12543, -1.50 DPS) [quest]; Maiden's Circle (13001, -1.50 DPS) [world_drop]; Naglering (11669, -9.55 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (325.4 DPS) | yes | Songstone of Ironforge (12543, -0.56 DPS) [quest]; Maiden's Circle (13001, -0.56 DPS) [world_drop]; Naglering (11669, -11.09 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (325.4 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Second Wind (11819, -6.57 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (325.4 DPS) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -1.32 DPS) [vendor]; Serenity Field (272439, -2.82 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (325.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.29 DPS) [world]; Teebu's Blazing Longsword (1728, -13.81 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 414.5 spell_power points (77.89 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.97 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.44 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.94 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 33.8. Weights run: 0.7s. Verify run: 0.6s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.379 ± 0.007, crit=0.075 ± 0.002 per rating point (14 rating = 1%, 1.054 per %), hit=0.196 ± 0.001 per rating point (10 rating = 1%, 1.956 per %), spell_haste=0.781 ± 0.127, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.77 DPS) | yes | Shadow Goggles (4373, -0.89 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (1.08 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.27 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.57 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.51 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.89 DPS) | yes | Green Woolen Robe (6243, -0.35 DPS) [crafted]; Green Woolen Vest (2582, -0.37 DPS) [crafted]; Gray Woolen Robe (2585, -0.77 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.3 spell_power points (0.29 DPS) | yes | Mindthrust Bracers (1974, -0.05 DPS) [dungeon]; Featherbead Bracers (15452, -0.05 DPS) [quest]; Owlbeard Bracers (16981, -0.07 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.90 DPS) | yes | Gnoll Casting Gloves (892, -0.13 DPS) [world]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Apothecary Gloves (10919, -0.39 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.71 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.32 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.51 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.0 spell_power points (1.54 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.65 DPS) [dungeon]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (1.09 DPS) | yes | Pristine Boots (253889, -0.56 DPS) [crafted]; Red Woolen Boots (4313, -0.58 DPS) [crafted]; Feather Padded Treads (285345, -0.72 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.64 DPS) | yes | Loop of Sacrifice (281673, -0.40 DPS) [quest]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop]; Sludge-Stained Band (286535, -0.95 DPS, sim-verified) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (33.8 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS) [quest]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop]; Sludge-Stained Band (286535, -0.46 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.8 spell_power points (0.49 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.10 DPS) [world]; Lesser Staff of the Spire (1300, -0.19 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 176.8 spell_power points (22.70 DPS) | yes | Skycaller (12984, -1.13 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.41 DPS) [dungeon]; Sizzle Stick (8071, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 56.1. Weights run: 0.7s. Verify run: 0.6s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.475 ± 0.012, crit=0.130 ± 0.005 per rating point (14 rating = 1%, 1.820 per %), hit=0.226 ± 0.002 per rating point (10 rating = 1%, 2.261 per %), spell_haste=not significant (0.631 ± 0.199), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.65 DPS) | yes | Enchanter's Cowl (4322, -0.04 DPS) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.9 spell_power points (1.48 DPS) | yes | Crystal Starfire Medallion (5003, -1.19 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.19 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.49 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.3 spell_power points (1.99 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.59 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.2 spell_power points (2.28 DPS) | yes | Death Speaker Robes (6682, -0.44 DPS) [dungeon]; Pristine Gown (253961, -0.73 DPS) [crafted]; Tree Bark Jacket (1486, -1.17 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Nightsky Wristbands (6407, -0.92 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.92 DPS) [quest]; Glowing Magical Bracelets (13106, -1.24 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.4 spell_power points (1.26 DPS) | yes | Truefaith Gloves (7049, -0.29 DPS) [crafted]; Gnoll Casting Gloves (892, -0.36 DPS) [world]; Serpent Gloves (5970, -0.39 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.4 spell_power points (1.87 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Invoker's Cord (215366, -0.46 DPS) [crafted]; Warsong Sash (16975, -0.57 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.8 spell_power points (1.92 DPS) | yes | Gaze Dreamer Pants (6903, -0.12 DPS) [dungeon]; Pristine Leggings (253987, -0.37 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.55 DPS) | yes | Acidic Walkers (9454, -0.23 DPS) [dungeon]; Boots of the Enchanter (4325, -0.80 DPS) [crafted]; Spidersilk Boots (4320, -1.78 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.05 DPS) | yes | Black Widow Band (6199, -0.55 DPS) [world]; Snake Hoop (6750, -0.55 DPS) [quest]; Sludge-Stained Band (286535, -0.60 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.90 DPS) | yes | Snake Hoop (6750, -0.40 DPS) [quest]; Sludge-Stained Band (286535, -0.45 DPS) [world]; Black Widow Band (6199, -1.25 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Glimmering Staff (249392, -0.60 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.64 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.64 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.9 spell_power points (1.48 DPS) | yes | Orb of Souls (249395, -0.88 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.89 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.21 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 224.2 spell_power points (33.66 DPS) | yes | Starfaller (13063, -0.51 DPS) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 100000000000000000-00000000000000000-2535111300000301050)

Set DPS (verified): 90.9. Weights run: 0.7s. Verify run: 0.5s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.722 ± 0.024, crit=0.208 ± 0.009 per rating point (14 rating = 1%, 2.914 per %), hit=0.374 ± 0.004 per rating point (10 rating = 1%, 3.744 per %), spell_haste=1.758 ± 0.380, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.10 DPS) | yes | Augural Shroud (2620, -0.41 DPS) [world]; Corpseshroud (10574, -1.08 DPS) [dungeon]; Enchanter's Cowl (4322, -1.15 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.3 spell_power points (1.67 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.61 DPS) [quest]; Triune Amulet (7722, -0.93 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.93 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.4 spell_power points (2.42 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.5 spell_power points (2.28 DPS) | yes | Guardian Cloak (5965, -0.87 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.11 DPS) [vendor]; Long Silken Cloak (4326, -1.31 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.3 spell_power points (3.88 DPS) | yes | Dreamweave Vest (10021, -0.27 DPS) [crafted]; Robe of Power (7054, -0.54 DPS) [crafted]; Elemental Raiment (9434, -0.79 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.8 spell_power points (1.44 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.11 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.9 spell_power points (3.08 DPS) | yes | Black Mageweave Gloves (10003, -0.87 DPS) [crafted]; Red Mageweave Gloves (10018, -1.07 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.16 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.8 spell_power points (2.63 DPS) | yes | Gilded Cord (254037, -0.60 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.69 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.7 spell_power points (3.34 DPS) | yes | Crimson Silk Pantaloons (7062, -0.77 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.16 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.46 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.54 DPS) | yes | Gilded Slippers (254001, -1.42 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.95 DPS) [dungeon]; Spidersilk Boots (4320, -2.08 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (2.11 DPS) | yes | Reedknot Ring (9622, -1.08 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.23 DPS) [vendor]; Black Widow Band (6199, -1.37 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.33 DPS) | yes | Sea Giant's Toe Ring (274746, -0.44 DPS) [vendor]; Black Widow Band (6199, -0.58 DPS) [world]; Reedknot Ring (9622, -1.22 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (90.9 DPS) | yes | Windweaver Staff (7757, -1.35 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Gut Ripper (2164, -4.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 271.4 spell_power points (40.02 DPS) | yes | Nether Force Wand (11263, -2.39 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.51 DPS) [quest]; Ragefire Wand (7513, -2.56 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 203005000100000000-00000000000000000-2535111300000301050)

Set DPS (verified): 189.2. Weights run: 0.7s. Verify run: 0.6s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.858 ± 0.037, crit=0.315 ± 0.013 per rating point (14 rating = 1%, 4.406 per %), hit=0.641 ± 0.008 per rating point (10 rating = 1%, 6.410 per %), spell_haste=12.188 ± 0.716, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.2 spell_power points (7.23 DPS) | yes | Dreamweave Circlet (10041, -1.32 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.49 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.83 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.1 spell_power points (4.23 DPS) | yes | Horizon Choker (13085, -1.83 DPS) [world_drop]; Mindburst Medallion (11196, -2.00 DPS) [quest]; Scorn's Icy Choker (23169, -3.97 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.4 spell_power points (5.69 DPS) | yes | Kentic Amice (11624, -0.66 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.66 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.71 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.7 spell_power points (3.94 DPS) | yes | Spritecaster Cape (11623, -0.11 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.60 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.2 spell_power points (7.23 DPS) | yes | Robe of the Magi (1716, -1.80 DPS) [world_drop]; Runecloth Tunic (13857, -1.94 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -2.06 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 13.0 spell_power points (2.60 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Aristocratic Cuffs (12546, -0.03 DPS) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.4 spell_power points (6.08 DPS) | yes | Raider Handwraps (272098, -0.91 DPS) [vendor]; Dreamweave Gloves (10019, -1.80 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.94 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.6 spell_power points (4.52 DPS) | yes | Dawnspire Cord (12466, -0.06 DPS) [dungeon]; Ban'thok Sash (11662, -0.10 DPS) [dungeon]; Deathmage Sash (10771, -0.54 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.6 spell_power points (6.32 DPS) | yes | Red Mageweave Pants (10009, -1.46 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.49 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.29 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.80 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.37 DPS) [vendor]; Gilded Sandals (254107, -1.06 DPS) [crafted]; Southsea Mojo Boots (20641, -1.31 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (3.03 DPS) | yes | Band of the Unicorn (7553, -0.43 DPS) [world_drop]; Brainlash (6440, -0.46 DPS) [dungeon]; Advisor's Ring (19519, -0.63 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.0 spell_power points (3.00 DPS) | yes | Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Brainlash (6440, -0.43 DPS) [dungeon]; Advisor's Ring (19519, -0.60 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (189.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (189.2 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellforce Rod (1664, -0.97 DPS) [world]; Spellshifter Rod (9527, -1.03 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 262.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.50 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.15 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 203005000100000000-11302300000000000-2535111300000301050)

Set DPS (verified): 305.6. Weights run: 0.8s. Verify run: 0.6s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.993 ± 0.051, crit=0.496 ± 0.019 per rating point (14 rating = 1%, 6.948 per %), hit=1.032 ± 0.011 per rating point (10 rating = 1%, 10.317 per %), spell_haste=16.128 ± 1.020, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 52.8 spell_power points (9.92 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.31 DPS) [pvp]; Magister's Crown (16686, -2.25 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (305.6 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.56 DPS) [quest]; Chains of the Lich (23125, -1.11 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.8 spell_power points (9.18 DPS) | yes | Warlord's Silk Amice (231594, -1.68 DPS) [pvp]; Darkspear Shoulderpads (272103, -1.89 DPS) [vendor]; Mantle of the Timbermaw (19050, -2.25 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 34.3 spell_power points (6.44 DPS) | yes | Crystalline Threaded Cape (20697, -1.93 DPS) [world]; Hide of the Wild (18510, -1.94 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.50 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.9 spell_power points (11.06 DPS) | yes | Warlord's Silk Raiment (231596, -0.38 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.64 DPS) [pvp]; Sorcerer's Robes (226932, -3.39 DPS) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.9 spell_power points (5.63 DPS) | yes | Sublime Wristguards (18497, -1.51 DPS) [dungeon]; Runecloth Cuffs (254123, -1.69 DPS) [crafted]; General's Silk Cuffs (16538, -2.08 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 36.2 spell_power points (6.81 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 57.2 spell_power points (10.74 DPS) | yes | Belt of the Archmage (18405, -2.69 DPS) [crafted]; Magician's Cord (272393, -2.69 DPS) [vendor]; Frostwolf Cloth Belt (19090, -5.49 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.8 spell_power points (10.49 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.68 DPS) [rep]; Sorcerer's Leggings (226933, -2.00 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 36.9 spell_power points (6.93 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.56 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (305.6 DPS) | yes | Eye of Orgrimmar (12545, -1.50 DPS) [quest]; Maiden's Circle (13001, -1.50 DPS) [world_drop]; Naglering (11669, -8.12 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (305.6 DPS) | yes | Eye of Orgrimmar (12545, -0.56 DPS) [quest]; Maiden's Circle (13001, -0.56 DPS) [world_drop]; Naglering (11669, -10.41 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (305.6 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (305.6 DPS) | yes | Weakness Analyzer (272438, -1.32 DPS) [vendor]; Serenity Field (272439, -2.82 DPS) [vendor]; Burst of Knowledge (11832, -3.19 DPS) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (305.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.37 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -13.99 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 414.5 spell_power points (77.89 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.97 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.44 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.94 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

