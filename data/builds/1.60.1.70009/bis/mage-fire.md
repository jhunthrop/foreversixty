# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 28.2. Weights run: 0.7s. Verify run: 0.6s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.304 ± 0.013, crit=0.129 ± 0.004 per rating point (14 rating = 1%, 1.809 per %), hit=0.428 ± 0.013 per rating point (10 rating = 1%, 4.280 per %), spell_haste=-1.028 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.40 DPS) | yes | Shadow Goggles (4373, -0.62 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.7 spell_power points (0.52 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.10 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.27 DPS) | yes | Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Caretaker's Cape (20428, -0.07 DPS) [rep]; Black Whelp Cloak (7283, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.44 DPS) | yes | Green Woolen Vest (2582, -0.17 DPS) [crafted]; Bloody Apron (6226, -0.17 DPS) [dungeon]; Gray Woolen Robe (2585, -0.54 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.5 spell_power points (0.10 DPS) | yes | Windsong Bangles (263336, -0.03 DPS) [quest]; Repurposed Hair Band (281256, -0.06 DPS) [quest]; Bright Bracers (3647, -0.56 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.47 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Heavy Woolen Gloves (4310, -0.29 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.35 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Keller's Girdle (2911, -0.19 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.38 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (28.2 DPS) | yes | Silk-threaded Trousers (1929, -0.06 DPS) [dungeon]; Rumpled Kilt (274741, -0.19 DPS) [vendor]; Abomination Skin Leggings (23173, -0.47 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.55 DPS) | yes | Red Woolen Boots (4313, -0.28 DPS) [crafted]; Pristine Boots (253889, -0.29 DPS) [crafted]; Feather Padded Treads (285345, -0.51 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.38 DPS) | yes | Sludge-Stained Band (286535, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.33 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Sludge-Stained Band (286535, -0.24 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.0 spell_power points (0.20 DPS) | yes | Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.08 DPS) [world]; Staff of Westfall (2042, -0.10 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 336.3 spell_power points (22.51 DPS) | yes | Skycaller (12984, -1.04 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.22 DPS) [dungeon]; Deepblaze (279896, -3.98 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 46.5. Weights run: 0.8s. Verify run: 0.7s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.661 ± 0.013, crit=0.166 ± 0.007 per rating point (14 rating = 1%, 2.327 per %), hit=0.349 ± 0.018 per rating point (10 rating = 1%, 3.494 per %), spell_haste=0.935 ± 0.221, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.6 spell_power points (1.18 DPS) | yes | Holy Shroud (2721, -0.15 DPS) [world_drop]; Silk Headband (7050, -0.34 DPS) [crafted]; Embalmed Shroud (7691, -0.43 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 spell_power points (1.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.78 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.78 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.93 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.9 spell_power points (1.40 DPS) | yes | Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.6 spell_power points (0.53 DPS) | yes | Cloak of Rot (4462, -0.03 DPS) [world]; Darkspear Raider's Cloak (272078, -0.03 DPS) [vendor]; Hillman's Cloak (3719, -0.06 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.6 spell_power points (1.65 DPS) | yes | Death Speaker Robes (6682, -0.31 DPS) [dungeon]; Tree Bark Jacket (1486, -0.43 DPS) [dungeon]; Pristine Gown (253961, -0.56 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Nightsky Wristbands (6407, -0.47 DPS) [world_drop]; Stonecloth Bindings (14416, -0.53 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.85 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.3 spell_power points (1.06 DPS) | yes | Serpent Gloves (5970, -0.40 DPS) [dungeon]; Shilly Mitts (9609, -0.40 DPS) [quest]; Truefaith Gloves (7049, -0.40 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.0 spell_power points (1.22 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Crimson Silk Belt (7055, -0.22 DPS) [crafted]; Invoker's Cord (215366, -0.25 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.3 spell_power points (1.34 DPS) | yes | Gaze Dreamer Pants (6903, -0.21 DPS) [dungeon]; Pristine Leggings (253987, -0.25 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.41 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.6 spell_power points (1.09 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Nimbus Boots (6998, -0.53 DPS) [quest]; Acidic Walkers (9454, -1.10 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.22 DPS) [world]; Snake Hoop (6750, -0.22 DPS) [quest]; Minor Channeling Ring (1449, -1.24 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (46.5 DPS) | yes | Black Widow Band (6199, -0.13 DPS) [world]; Snake Hoop (6750, -0.13 DPS) [quest]; Minor Channeling Ring (1449, -0.88 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Twisted Chanter's Staff (890, -0.22 DPS) [world_drop]; Channeler's Staff (4437, -0.35 DPS) [world]; Glimmering Staff (249392, -0.37 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.0 spell_power points (1.03 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.59 DPS) [vendor]; Eye of Paleth (2943, -0.65 DPS) [quest]; Dwarven Tome (279898, -0.78 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 356.6 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.38 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-23552100130103041-0000000000000000000)

Set DPS (verified): 82.7. Weights run: 0.8s. Verify run: 0.7s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.774 ± 0.026, crit=0.219 ± 0.015 per rating point (14 rating = 1%, 3.059 per %), hit=0.376 ± 0.034 per rating point (10 rating = 1%, 3.765 per %), spell_haste=2.538 ± 0.457, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.59 DPS) | yes | Augural Shroud (2620, -0.28 DPS) [world]; Corpseshroud (10574, -0.78 DPS) [dungeon]; Enchanter's Cowl (4322, -0.90 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.6 spell_power points (1.44 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.48 DPS) [quest]; Triune Amulet (7722, -0.77 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.77 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.1 spell_power points (2.11 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.0 spell_power points (1.97 DPS) | yes | Guardian Cloak (5965, -0.75 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.92 DPS) [vendor]; Long Silken Cloak (4326, -1.21 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.6 spell_power points (3.29 DPS) | yes | Dreamweave Vest (10021, -0.21 DPS) [crafted]; Robe of Power (7054, -0.41 DPS) [crafted]; Elemental Raiment (9434, -0.70 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (82.7 DPS) | yes | Condor Bracers (15864, -0.25 DPS) [quest]; Windchaser Cuffs (14429, -0.25 DPS) [world_drop]; Arcane Runed Bracers (4744, -0.98 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.60 DPS) | yes | Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Red Mageweave Gloves (10018, -0.81 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -0.95 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.6 spell_power points (2.30 DPS) | yes | Highlander's Cloth Girdle (20098, -0.19 DPS) [rep]; Gilded Cord (254037, -0.55 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.3 spell_power points (2.87 DPS) | yes | Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.25 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.96 DPS) | yes | Gilded Slippers (254001, -1.43 DPS) [crafted]; Acidic Walkers (9454, -1.58 DPS) [dungeon]; Spidersilk Boots (4320, -1.72 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.81 DPS) | yes | Ring of Forlorn Spirits (2043, -0.82 DPS) [quest]; Reedknot Ring (9622, -0.94 DPS) [quest]; Minor Channeling Ring (1449, -1.00 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.11 DPS) | yes | Reedknot Ring (9622, -0.25 DPS) [quest]; Minor Channeling Ring (1449, -0.30 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.80 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -1.04 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.36 DPS) [dungeon]; Gut Ripper (2164, -4.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 323.6 spell_power points (39.95 DPS) | yes | Nether Force Wand (11263, -0.82 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.48 DPS) [quest]; Ragefire Wand (7513, -2.53 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 145.5. Weights run: 0.8s. Verify run: 0.8s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.939 ± 0.038, crit=0.343 ± 0.024 per rating point (14 rating = 1%, 4.796 per %), hit=0.622 ± 0.048 per rating point (10 rating = 1%, 6.225 per %), spell_haste=-3.557 ± 0.738, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.8 spell_power points (5.51 DPS) | yes | Dreamweave Circlet (10041, -1.08 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.13 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.57 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.6 spell_power points (3.15 DPS) | yes | Scorn's Icy Choker (23169, -1.31 DPS) [dungeon]; Mindburst Medallion (11196, -1.46 DPS) [quest]; Horizon Choker (13085, -4.02 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.9 spell_power points (4.36 DPS) | yes | Kentic Amice (11624, -0.54 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.26 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.29 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.6 spell_power points (2.86 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.32 DPS) [dungeon]; Runecloth Cloak (13860, -0.46 DPS) [crafted]; Big Voodoo Cloak (8216, -0.90 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.8 spell_power points (5.51 DPS) | yes | Robe of the Magi (1716, -1.48 DPS) [world_drop]; Runecloth Tunic (13857, -1.52 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.55 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (145.5 DPS) | yes | Bloodband Bracers (11469, -0.02 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.34 DPS) [quest]; Aristocratic Cuffs (12546, -3.09 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.4 spell_power points (4.57 DPS) | yes | Raider Handwraps (272098, -0.54 DPS) [vendor]; Dreamweave Gloves (10019, -1.40 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.45 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 28.6 spell_power points (4.16 DPS) | yes | Satyrmane Sash (17755, -0.75 DPS) [dungeon]; Deathmage Sash (10771, -1.09 DPS) [dungeon]; Dawnspire Cord (12466, -2.33 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.4 spell_power points (4.72 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.78 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.92 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.50 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.19 DPS) [vendor]; Gilded Sandals (254107, -0.66 DPS) [crafted]; Southsea Mojo Boots (20641, -0.83 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.6 spell_power points (2.28 DPS) | yes | Brainlash (6440, -0.23 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.53 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.6 spell_power points (2.27 DPS) | yes | Brainlash (6440, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.52 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -0.82 DPS) [quest]; Spellforce Rod (1664, -1.05 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 360.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 419.9. Weights run: 0.9s. Verify run: 2.2s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.923 ± 0.054, crit=0.719 ± 0.043 per rating point (14 rating = 1%, 10.065 per %), hit=1.212 ± 0.083 per rating point (10 rating = 1%, 12.117 per %), spell_haste=-10.458 ± 1.052, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 54.1 spell_power points (14.58 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.74 DPS) [pvp]; Magister's Crown (16686, -11.94 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (419.9 DPS) | yes | Amulet of the Dawn (22657, -0.41 DPS) [quest]; Beads of Ogre Mojo (22149, -1.19 DPS) [quest]; Jewel of Kajaro (19601, -11.64 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 50.9 spell_power points (13.71 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.25 DPS) [pvp]; Darkspear Shoulderpads (272103, -3.77 DPS) [vendor]; Mantle of the Timbermaw (19050, -8.71 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.5 spell_power points (9.56 DPS) | yes | Hide of the Wild (18510, -3.31 DPS) [crafted]; Shroud of Arcane Mastery (22330, -3.56 DPS) [dungeon]; Crystalline Threaded Cape (20697, -13.11 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 61.1 spell_power points (16.47 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.64 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.88 DPS) [pvp]; Sorcerer's Robes (226932, -14.06 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.4 spell_power points (7.92 DPS) | yes | Sublime Wristguards (18497, -2.20 DPS) [dungeon]; Runecloth Cuffs (254123, -2.47 DPS) [crafted]; Marshal's Silk Bracers (16438, -3.19 DPS) [pvp] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 38.4 spell_power points (10.33 DPS) | yes | Marshal's Silk Gloves (16440, -0.08 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -0.08 DPS) [vendor]; Sorcerer's Gloves (22066, -9.70 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 57.4 spell_power points (15.47 DPS) | yes | Magician's Cord (272393, -4.38 DPS) [vendor]; Ban'thok Sash (11662, -6.24 DPS) [dungeon]; Belt of the Archmage (18405, -13.85 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 61.4 spell_power points (16.53 DPS) | yes | Marshal's Silk Leggings (231605, -0.76 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.94 DPS) [pvp]; Sorcerer's Leggings (226933, -15.43 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 35.8 spell_power points (9.63 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.81 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (419.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.26 DPS) [quest]; Channeler's Ring (272406, -1.92 DPS) [vendor]; Naglering (11669, -21.46 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (419.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.18 DPS) [quest]; Channeler's Ring (272406, -0.83 DPS) [vendor]; Naglering (11669, -21.14 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (419.9 DPS) | yes | Weakness Analyzer (272438, -1.89 DPS) [vendor]; Blackhand's Breadth (13965, -2.39 DPS) [quest]; Eye of the Beast (13968, -2.39 DPS) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (419.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (419.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -0.82 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -33.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 290.9 spell_power points (78.38 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.56 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.79 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.46 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 205015100000000000-23252100130133051-0050000000000000000)

Set DPS (verified): 792.9. Weights run: 1.0s. Verify run: 2.0s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.409 ± 0.029, crit=0.870 ± 0.043 per rating point (14 rating = 1%, 12.183 per %), hit=1.128 ± 0.065 per rating point (10 rating = 1%, 11.277 per %), spell_haste=not significant (-0.196 ± 0.515), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 43.4 spell_power points (22.35 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.47 DPS) [pvp]; Crimson Felt Hat (18727, -7.09 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (792.9 DPS) | yes | Orb of the Darkmoon (19426, -0.80 DPS) [quest]; Chains of the Lich (23125, -0.80 DPS) [dungeon]; Jewel of Kajaro (19601, -12.91 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.3 spell_power points (23.34 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, -7.02 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, -7.31 DPS) [pvp]; Mantle of the Timbermaw (19050, -7.96 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.5 spell_power points (15.73 DPS) | yes | Hide of the Wild (18510, -6.42 DPS) [crafted]; Amplifying Cloak (18350, -6.46 DPS) [dungeon]; Crystalline Threaded Cape (20697, -6.98 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 57.1 spell_power points (29.40 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.55 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -8.73 DPS) [pvp]; Robe of Everlasting Night (18385, -17.98 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.3 spell_power points (13.02 DPS) | yes | Sublime Wristguards (18497, -4.73 DPS) [dungeon]; Runecloth Cuffs (254123, -5.25 DPS) [crafted]; Arena Wristguards (18709, -6.74 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.5 spell_power points (19.29 DPS) | yes | Marshal's Silk Gloves (16440, -2.86 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.86 DPS) [vendor]; Sandworm Skin Gloves (20716, -18.83 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 45.3 spell_power points (23.32 DPS) | yes | Magician's Cord (272393, -8.48 DPS) [vendor]; Highlander's Cloth Girdle (20047, -8.57 DPS) [rep]; Belt of the Archmage (18405, -10.66 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 60.5 spell_power points (31.14 DPS) | yes | Marshal's Silk Leggings (231605, -5.20 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -10.47 DPS) [pvp]; Skyshroud Leggings (13170, -21.36 DPS, sim-verified) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 27.5 spell_power points (14.18 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -1.26 DPS) [pvp] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (792.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.76 DPS) [quest]; Don Julio's Band (19325, -1.97 DPS) [rep]; Naglering (11669, -26.99 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (792.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.36 DPS) [quest]; Don Julio's Band (19325, -1.56 DPS) [rep]; Naglering (11669, -19.23 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (792.9 DPS) | yes | Eye of the Beast (13968, -2.39 DPS) [quest]; Weakness Analyzer (272438, -3.61 DPS) [vendor]; Serenity Field (272439, -7.73 DPS) [vendor] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (792.9 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Weakness Analyzer (272438, -1.22 DPS) [vendor]; Serenity Field (272439, -5.34 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (792.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -0.05 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -42.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.0 spell_power points (79.85 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.33 DPS) [dungeon]; Bonecreeper Stylus (13938, -10.71 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.05 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Blackhand's Breadth; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 26.5. Weights run: 0.7s. Verify run: 0.6s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.304 ± 0.013, crit=0.129 ± 0.004 per rating point (14 rating = 1%, 1.809 per %), hit=0.428 ± 0.013 per rating point (10 rating = 1%, 4.280 per %), spell_haste=-1.028 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.40 DPS) | yes | Shadow Goggles (4373, -0.55 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.7 spell_power points (0.52 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.33 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.27 DPS) | yes | Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.07 DPS) [rep]; Black Whelp Cloak (7283, -0.22 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.44 DPS) | yes | Green Woolen Vest (2582, -0.17 DPS) [crafted]; Bloody Apron (6226, -0.17 DPS) [dungeon]; Gray Woolen Robe (2585, -0.65 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.8 spell_power points (0.12 DPS) | yes | Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Owlbeard Bracers (16981, -0.25 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.47 DPS) | yes | Gnoll Casting Gloves (892, -0.07 DPS) [world]; Pristine Gloves (253913, -0.14 DPS) [crafted]; Apothecary Gloves (10919, -0.20 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.35 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Keller's Girdle (2911, -0.19 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.34 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (26.5 DPS) | yes | Silk-threaded Trousers (1929, -0.06 DPS) [dungeon]; Rumpled Kilt (274741, -0.19 DPS) [vendor]; Abomination Skin Leggings (23173, -0.29 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.55 DPS) | yes | Red Woolen Boots (4313, -0.28 DPS) [crafted]; Pristine Boots (253889, -0.29 DPS) [crafted]; Feather Padded Treads (285345, -0.51 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.33 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Loop of Sacrifice (281673, -0.23 DPS) [quest]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.20 DPS) | yes | Lavishly Jeweled Ring (1156, -0.08 DPS) [dungeon]; Loop of Sacrifice (281673, -0.10 DPS) [quest]; Volcanic Rock Ring (12053, -0.14 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.0 spell_power points (0.20 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.08 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 336.3 spell_power points (22.51 DPS) | yes | Skycaller (12984, -0.73 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.22 DPS) [dungeon]; Sizzle Stick (8071, -4.53 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 43.6. Weights run: 0.8s. Verify run: 0.7s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.661 ± 0.013, crit=0.166 ± 0.007 per rating point (14 rating = 1%, 2.327 per %), hit=0.349 ± 0.018 per rating point (10 rating = 1%, 3.494 per %), spell_haste=0.935 ± 0.221, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.6 spell_power points (1.18 DPS) | yes | Holy Shroud (2721, -0.15 DPS) [world_drop]; Silk Headband (7050, -0.34 DPS) [crafted]; Embalmed Shroud (7691, -0.43 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 spell_power points (1.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.78 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.78 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.98 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.9 spell_power points (1.40 DPS) | yes | Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.3 spell_power points (0.50 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Hillman's Cloak (3719, -0.03 DPS) [crafted]; Windsong Drape (15468, -0.03 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.6 spell_power points (1.65 DPS) | yes | Death Speaker Robes (6682, -0.40 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.43 DPS) [dungeon]; Pristine Gown (253961, -0.56 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Nightsky Wristbands (6407, -0.47 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.47 DPS) [quest]; Glowing Magical Bracelets (13106, -0.61 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.3 spell_power points (0.87 DPS) | yes | Serpent Gloves (5970, -0.22 DPS) [dungeon]; Truefaith Gloves (7049, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.31 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.0 spell_power points (1.22 DPS) | yes | Warsong Sash (16975, -0.19 DPS) [quest]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Crimson Silk Belt (7055, -0.22 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.3 spell_power points (1.34 DPS) | yes | Gaze Dreamer Pants (6903, -0.21 DPS) [dungeon]; Pristine Leggings (253987, -0.25 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.41 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.6 spell_power points (1.09 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Boots of the Enchanter (4325, -0.62 DPS) [crafted]; Acidic Walkers (9454, -0.74 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.22 DPS) [world]; Snake Hoop (6750, -0.22 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.56 DPS) | yes | Snake Hoop (6750, -0.13 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Black Widow Band (6199, -0.48 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Glimmering Staff (249392, -0.16 DPS) [crafted]; Twisted Chanter's Staff (890, -0.22 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.22 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.0 spell_power points (1.03 DPS) | yes | Witch's Finger (16887, -0.60 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.65 DPS) [world]; Tome of the Darkspear Prophecy (272090, -0.76 DPS, sim-verified) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 356.6 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.38 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552100130103041-0000000000000000000)

Set DPS (verified): 79.3. Weights run: 0.8s. Verify run: 0.6s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.774 ± 0.026, crit=0.219 ± 0.015 per rating point (14 rating = 1%, 3.059 per %), hit=0.376 ± 0.034 per rating point (10 rating = 1%, 3.765 per %), spell_haste=2.538 ± 0.457, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.59 DPS) | yes | Augural Shroud (2620, -0.28 DPS) [world]; Corpseshroud (10574, -0.78 DPS) [dungeon]; Enchanter's Cowl (4322, -0.90 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.6 spell_power points (1.44 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.48 DPS) [quest]; Triune Amulet (7722, -0.77 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.77 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.1 spell_power points (2.11 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.0 spell_power points (1.97 DPS) | yes | Guardian Cloak (5965, -0.75 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.92 DPS) [vendor]; Long Silken Cloak (4326, -1.43 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.6 spell_power points (3.29 DPS) | yes | Dreamweave Vest (10021, -0.21 DPS) [crafted]; Robe of Power (7054, -0.41 DPS) [crafted]; Elemental Raiment (9434, -0.70 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.2 spell_power points (1.26 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.39 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.1 spell_power points (2.60 DPS) | yes | Red Mageweave Gloves (10018, -0.75 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Gilded Handwraps (254021, -0.95 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.6 spell_power points (2.30 DPS) | yes | Defiler's Cloth Girdle (20166, -0.19 DPS) [rep]; Gilded Cord (254037, -0.55 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.3 spell_power points (2.87 DPS) | yes | Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.25 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.96 DPS) | yes | Gilded Slippers (254001, -1.43 DPS) [crafted]; Acidic Walkers (9454, -1.58 DPS) [dungeon]; Spidersilk Boots (4320, -1.72 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.81 DPS) | yes | Reedknot Ring (9622, -0.94 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.07 DPS) [vendor]; Black Widow Band (6199, -1.14 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.11 DPS) | yes | Reedknot Ring (9622, -0.25 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.37 DPS) [vendor]; Black Widow Band (6199, -0.44 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (79.3 DPS) | yes | Windweaver Staff (7757, -1.04 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.36 DPS) [dungeon]; Gut Ripper (2164, -3.99 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 323.6 spell_power points (39.95 DPS) | yes | Nether Force Wand (11263, -0.88 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.48 DPS) [quest]; Ragefire Wand (7513, -2.53 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 139.0. Weights run: 0.8s. Verify run: 0.8s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.939 ± 0.038, crit=0.343 ± 0.024 per rating point (14 rating = 1%, 4.796 per %), hit=0.622 ± 0.048 per rating point (10 rating = 1%, 6.225 per %), spell_haste=-3.557 ± 0.738, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.8 spell_power points (5.51 DPS) | yes | Dreamweave Circlet (10041, -1.08 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.13 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.57 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.6 spell_power points (3.15 DPS) | yes | Scorn's Icy Choker (23169, -1.31 DPS) [dungeon]; Mindburst Medallion (11196, -1.46 DPS) [quest]; Horizon Choker (13085, -3.01 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.9 spell_power points (4.36 DPS) | yes | Kentic Amice (11624, -0.54 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.26 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.29 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.5 spell_power points (2.98 DPS) | yes | Spritecaster Cape (11623, -0.12 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.44 DPS) [dungeon]; Runecloth Cloak (13860, -0.57 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.8 spell_power points (5.51 DPS) | yes | Robe of the Magi (1716, -1.48 DPS) [world_drop]; Runecloth Tunic (13857, -1.52 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.55 DPS) [vendor] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 14.1 spell_power points (2.05 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.07 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.4 spell_power points (4.57 DPS) | yes | Raider Handwraps (272098, -0.54 DPS) [vendor]; Dreamweave Gloves (10019, -1.40 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.45 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 28.6 spell_power points (4.16 DPS) | yes | Satyrmane Sash (17755, -0.75 DPS) [dungeon]; Deathmage Sash (10771, -1.09 DPS) [dungeon]; Dawnspire Cord (12466, -3.20 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.4 spell_power points (4.72 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.63 DPS) [vendor]; Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.78 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.50 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.19 DPS) [vendor]; Gilded Sandals (254107, -0.66 DPS) [crafted]; Southsea Mojo Boots (20641, -0.83 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.6 spell_power points (2.28 DPS) | yes | Brainlash (6440, -0.23 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Advisor's Ring (19519, -0.53 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.6 spell_power points (2.27 DPS) | yes | Brainlash (6440, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Advisor's Ring (19519, -0.52 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (139.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (139.0 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 360.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 411.0. Weights run: 0.9s. Verify run: 2.7s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.923 ± 0.054, crit=0.719 ± 0.043 per rating point (14 rating = 1%, 10.065 per %), hit=1.212 ± 0.083 per rating point (10 rating = 1%, 12.117 per %), spell_haste=-10.458 ± 1.052, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 54.1 spell_power points (14.58 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.74 DPS) [pvp]; Magister's Crown (16686, -23.53 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (411.0 DPS) | yes | Amulet of the Dawn (22657, -0.41 DPS) [quest]; Beads of Ogre Mojo (22149, -1.19 DPS) [quest]; Jewel of Kajaro (19601, -13.19 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 50.9 spell_power points (13.71 DPS) | yes | Warlord's Silk Amice (231594, -3.25 DPS) [pvp]; Darkspear Shoulderpads (272103, -3.77 DPS) [vendor]; Mantle of the Timbermaw (19050, -7.33 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.5 spell_power points (9.56 DPS) | yes | Hide of the Wild (18510, -3.31 DPS) [crafted]; Shroud of Arcane Mastery (22330, -3.56 DPS) [dungeon]; Crystalline Threaded Cape (20697, -11.09 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 61.1 spell_power points (16.47 DPS) | yes | Warlord's Silk Raiment (231596, -0.64 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.88 DPS) [pvp]; Sorcerer's Robes (226932, -21.27 DPS, sim-verified) [quest] |
| wrist | Sorcerer's Bindings (226929) | An Earnest Proposition [quest] | sim-verified (411.0 DPS) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (411.0 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Gloves of Spell Mastery (14146, -14.26 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 57.4 spell_power points (15.47 DPS) | yes | Magician's Cord (272393, -4.38 DPS) [vendor]; Ban'thok Sash (11662, -6.24 DPS) [dungeon]; Belt of the Archmage (18405, -14.15 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 61.4 spell_power points (16.53 DPS) | yes | General's Silk Trousers (231595, -0.76 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -3.94 DPS) [pvp]; Sorcerer's Leggings (226933, -14.40 DPS, sim-verified) [quest] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (411.0 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.81 DPS) [dungeon]; Sorcerer's Boots (22064, -14.49 DPS, sim-verified) [quest] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (411.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.26 DPS) [quest]; Channeler's Ring (272406, -1.92 DPS) [vendor]; Naglering (11669, -17.10 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (411.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.18 DPS) [quest]; Channeler's Ring (272406, -0.83 DPS) [vendor]; Naglering (11669, -21.19 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (411.0 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -1.65 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (411.0 DPS) | yes | Weakness Analyzer (272438, -1.89 DPS) [vendor]; Eye of the Beast (13968, -2.39 DPS) [quest]; Burst of Knowledge (11832, -12.35 DPS, sim-verified) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (411.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -0.82 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -36.88 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 290.9 spell_power points (78.38 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.56 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.79 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.46 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Sorcerer's Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Blackhand's Breadth; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (orc, 205015100000000000-23252100130133051-0050000000000000000)

Set DPS (verified): 774.3. Weights run: 1.0s. Verify run: 2.1s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.409 ± 0.029, crit=0.870 ± 0.043 per rating point (14 rating = 1%, 12.183 per %), hit=1.128 ± 0.065 per rating point (10 rating = 1%, 11.277 per %), spell_haste=not significant (-0.196 ± 0.515), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 43.4 spell_power points (22.35 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.47 DPS) [pvp]; Crimson Felt Hat (18727, -5.22 DPS) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (774.3 DPS) | yes | Orb of the Darkmoon (19426, -0.80 DPS) [quest]; Chains of the Lich (23125, -0.80 DPS) [dungeon]; Jewel of Kajaro (19601, -12.45 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.3 spell_power points (23.34 DPS) | yes | Mantle of the Timbermaw (19050, -6.70 DPS, sim-verified) [crafted]; Champion's Silk Mantle (227104, -7.02 DPS) [pvp]; Warlord's Silk Amice (231594, -7.31 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.5 spell_power points (15.73 DPS) | yes | Crystalline Threaded Cape (20697, -4.59 DPS) [world]; Hide of the Wild (18510, -6.42 DPS) [crafted]; Amplifying Cloak (18350, -6.46 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 57.1 spell_power points (29.40 DPS) | yes | Warlord's Silk Raiment (231596, -2.55 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -8.73 DPS) [pvp]; Robe of Everlasting Night (18385, -14.10 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.3 spell_power points (13.02 DPS) | yes | Sublime Wristguards (18497, -4.73 DPS) [dungeon]; Runecloth Cuffs (254123, -5.25 DPS) [crafted]; Arena Wristguards (18709, -6.74 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.5 spell_power points (19.29 DPS) | yes | General's Silk Handguards (16540, -2.86 DPS) [vendor]; General's Silk Gauntlets (231599, -2.86 DPS) [vendor]; Sandworm Skin Gloves (20716, -10.90 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 45.3 spell_power points (23.32 DPS) | yes | Belt of the Archmage (18405, -3.37 DPS) [crafted]; Magician's Cord (272393, -8.48 DPS) [vendor]; Defiler's Cloth Girdle (20163, -8.57 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 60.5 spell_power points (31.14 DPS) | yes | General's Silk Trousers (231595, -5.20 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -10.47 DPS) [pvp]; Skyshroud Leggings (13170, -13.83 DPS, sim-verified) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 27.5 spell_power points (14.18 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.26 DPS) [pvp] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (774.3 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.76 DPS) [quest]; Don Julio's Band (19325, -1.97 DPS) [rep]; Naglering (11669, -24.79 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (774.3 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.36 DPS) [quest]; Don Julio's Band (19325, -1.56 DPS) [rep]; Naglering (11669, -16.82 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (774.3 DPS) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Eye of the Beast (13968, -2.39 DPS) [quest]; Weakness Analyzer (272438, -3.61 DPS) [vendor] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (774.3 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -1.22 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (774.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -0.05 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -40.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.0 spell_power points (79.85 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.33 DPS) [dungeon]; Bonecreeper Stylus (13938, -10.71 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.05 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Blackhand's Breadth; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

