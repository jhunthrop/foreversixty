# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 31.3. Weights run: 1.1s. Verify run: 0.8s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.258 ± 0.009, crit=0.109 ± 0.004 per rating point (14 rating = 1%, 1.522 per %), hit=0.306 ± 0.002 per rating point (10 rating = 1%, 3.059 per %), spell_haste=not significant (-0.490 ± 0.165), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.52 DPS) | yes | Shadow Goggles (4373, -0.90 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.3 spell_power points (0.63 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.34 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.35 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.30 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.3 spell_power points (0.54 DPS) | yes | Green Woolen Vest (2582, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.20 DPS) [dungeon]; Gray Woolen Robe (2585, -0.85 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.3 spell_power points (0.11 DPS) | yes | Windsong Bangles (263336, -0.02 DPS) [quest]; Repurposed Hair Band (281256, -0.07 DPS) [quest]; Bright Bracers (3647, -0.41 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.60 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.39 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.0 spell_power points (0.43 DPS) | yes | Novice Ardent's Sash (253887, -0.19 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.59 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (31.3 DPS) | yes | Silk-threaded Trousers (1929, -0.05 DPS) [dungeon]; Rumpled Kilt (274741, -0.22 DPS) [vendor]; Abomination Skin Leggings (23173, -0.41 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.0 spell_power points (0.69 DPS) | yes | Feather Padded Treads (285345, -0.31 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.35 DPS) [crafted]; Pristine Boots (253889, -0.37 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (0.48 DPS) | yes | Sludge-Stained Band (286535, -0.22 DPS) [world]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.43 DPS) | yes | Sludge-Stained Band (286535, -0.24 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.6 spell_power points (0.22 DPS) | yes | Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world]; Staff of Westfall (2042, -0.11 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 261.4 spell_power points (22.57 DPS) | yes | Skycaller (12984, -0.82 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.28 DPS) [dungeon]; Deepblaze (279896, -4.04 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 48.7. Weights run: 1.2s. Verify run: 0.9s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.603 ± 0.020, crit=0.170 ± 0.009 per rating point (14 rating = 1%, 2.376 per %), hit=0.332 ± 0.003 per rating point (10 rating = 1%, 3.322 per %), spell_haste=not significant (-0.234 ± 0.295), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.0 spell_power points (1.15 DPS) | yes | Holy Shroud (2721, -0.10 DPS) [world_drop]; Silk Headband (7050, -0.29 DPS) [crafted]; Embalmed Shroud (7691, -0.39 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.6 spell_power points (1.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.78 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.78 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.79 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (1.38 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.4 spell_power points (0.52 DPS) | yes | Hillman's Cloak (3719, -0.04 DPS) [crafted]; Cloak of Rot (4462, -0.06 DPS) [world]; Darkspear Raider's Cloak (272078, -0.06 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.8 spell_power points (1.61 DPS) | yes | Tree Bark Jacket (1486, -0.37 DPS) [dungeon]; Death Speaker Robes (6682, -0.39 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Nightsky Wristbands (6407, -0.51 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.52 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -0.57 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.6 spell_power points (1.02 DPS) | yes | Serpent Gloves (5970, -0.35 DPS) [dungeon]; Shilly Mitts (9609, -0.35 DPS) [quest]; Truefaith Gloves (7049, -0.37 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.8 spell_power points (1.22 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Crimson Silk Belt (7055, -0.25 DPS) [crafted]; Invoker's Cord (215366, -0.27 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.8 spell_power points (1.32 DPS) | yes | Gaze Dreamer Pants (6903, -0.17 DPS) [dungeon]; Pristine Leggings (253987, -0.25 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.2 spell_power points (1.07 DPS) | yes | Spidersilk Boots (4320, -0.17 DPS) [crafted]; Nimbus Boots (6998, -0.50 DPS) [quest]; Acidic Walkers (9454, -0.95 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.67 DPS) | yes | Black Widow Band (6199, -0.27 DPS) [world]; Snake Hoop (6750, -0.27 DPS) [quest]; Minor Channeling Ring (1449, -1.30 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (48.7 DPS) | yes | Black Widow Band (6199, -0.17 DPS) [world]; Snake Hoop (6750, -0.17 DPS) [quest]; Minor Channeling Ring (1449, -0.76 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Glimmering Staff (249392, -0.23 DPS) [crafted]; Twisted Chanter's Staff (890, -0.28 DPS) [world_drop]; Channeler's Staff (4437, -0.40 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.6 spell_power points (1.01 DPS) | yes | Dwarven Tome (279898, -0.46 DPS, sim-verified) [quest]; Tome of the Darkspear Prophecy (272090, -0.59 DPS) [vendor]; Eye of Paleth (2943, -0.63 DPS) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 350.7 spell_power points (33.50 DPS) | yes | Starfaller (13063, -0.41 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-23552100130103050-0000000000000000000)

Set DPS (verified): 82.2. Weights run: 1.2s. Verify run: 0.8s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.605 ± 0.032, crit=0.273 ± 0.016 per rating point (14 rating = 1%, 3.817 per %), hit=0.429 ± 0.004 per rating point (10 rating = 1%, 4.288 per %), spell_haste=not significant (0.673 ± 0.512), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.43 DPS) | yes | Augural Shroud (2620, -0.46 DPS) [world]; Living Cowl (5608, -0.93 DPS) [world]; Enchanter's Cowl (4322, -1.04 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.6 spell_power points (1.23 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.53 DPS) [quest]; Triune Amulet (7722, -0.74 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.74 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.9 spell_power points (1.72 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Berylline Pads (4197, -0.21 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.4 spell_power points (1.67 DPS) | yes | Guardian Cloak (5965, -0.63 DPS) [crafted]; Icy Cloak (4327, -0.86 DPS) [crafted]; Long Silken Cloak (4326, -1.96 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.6 spell_power points (2.97 DPS) | yes | Dreamweave Vest (10021, -0.25 DPS) [crafted]; Robe of Power (7054, -0.51 DPS) [crafted]; Elemental Raiment (9434, -0.54 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.04 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.23 DPS) [quest]; Windchaser Cuffs (14429, -0.41 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (2.37 DPS) | yes | Red Mageweave Gloves (10018, -0.39 DPS) [crafted]; Black Mageweave Gloves (10003, -0.63 DPS) [crafted]; Gilded Handwraps (254021, -0.95 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (82.2 DPS) | yes | Star Belt (4329, -0.36 DPS) [crafted]; Gilded Cord (254037, -0.37 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.39 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.3 spell_power points (2.46 DPS) | yes | Abomination Skin Leggings (23173, -0.86 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.93 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.78 DPS) | yes | Gilded Slippers (254001, -1.48 DPS) [crafted]; Acidic Walkers (9454, -1.64 DPS) [dungeon]; Spidersilk Boots (4320, -1.69 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.6 spell_power points (1.58 DPS) | yes | Ring of Forlorn Spirits (2043, -0.65 DPS) [quest]; Reedknot Ring (9622, -0.77 DPS) [quest]; Minor Channeling Ring (1449, -0.86 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.04 DPS) | yes | Reedknot Ring (9622, -0.23 DPS) [quest]; Minor Channeling Ring (1449, -0.32 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.84 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -1.27 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.27 DPS) [dungeon]; Gut Ripper (2164, -3.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 343.4 spell_power points (39.80 DPS) | yes | Nether Force Wand (11263, -1.52 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.40 DPS) [quest]; Ragefire Wand (7513, -2.45 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 141.9. Weights run: 1.2s. Verify run: 0.9s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.915 ± 0.039, crit=0.263 ± 0.021 per rating point (14 rating = 1%, 3.681 per %), hit=0.589 ± 0.007 per rating point (10 rating = 1%, 5.887 per %), spell_haste=not significant (0.096 ± 0.803), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.3 spell_power points (5.58 DPS) | yes | Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.29 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.54 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (3.22 DPS) | yes | Scorn's Icy Choker (23169, -1.35 DPS) [dungeon]; Mindburst Medallion (11196, -1.50 DPS) [quest]; Horizon Choker (13085, -3.65 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.5 spell_power points (4.41 DPS) | yes | Kentic Amice (11624, -0.53 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.31 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.43 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 spell_power points (2.92 DPS) | yes | Runecloth Cloak (13860, -0.47 DPS) [crafted]; Big Voodoo Cloak (8216, -0.94 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.85 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.3 spell_power points (5.58 DPS) | yes | Robe of the Magi (1716, -1.47 DPS) [world_drop]; Runecloth Tunic (13857, -1.53 DPS) [crafted]; Runecloth Robe (13858, -1.61 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.7 spell_power points (2.05 DPS) | yes | Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Bloodband Bracers (11469, -0.07 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.41 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.7 spell_power points (4.59 DPS) | yes | Raider Handwraps (272098, -0.53 DPS) [vendor]; Dreamweave Gloves (10019, -1.35 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.42 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.4 spell_power points (3.50 DPS) | yes | Satyrmane Sash (17755, -0.04 DPS) [dungeon]; Ban'thok Sash (11662, -0.11 DPS) [dungeon]; Deathmage Sash (10771, -0.40 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.2 spell_power points (4.81 DPS) | yes | Red Mageweave Pants (10009, -1.07 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.83 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.43 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.28 DPS) [vendor]; Gilded Sandals (254107, -0.71 DPS) [crafted]; Southsea Mojo Boots (20641, -0.89 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.32 DPS) | yes | Brainlash (6440, -0.26 DPS) [dungeon]; Band of the Unicorn (7553, -0.37 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.52 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.30 DPS) | yes | Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.51 DPS) [rep]; Brainlash (6440, -2.29 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (141.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (141.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (141.9 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 350.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.75 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 370.1. Weights run: 1.3s. Verify run: 0.9s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.836 ± 0.062, crit=0.543 ± 0.037 per rating point (14 rating = 1%, 7.601 per %), hit=0.905 ± 0.013 per rating point (10 rating = 1%, 9.047 per %), spell_haste=not significant (-4.127 ± 1.161), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 49.5 spell_power points (11.83 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.40 DPS) [pvp]; Crimson Felt Hat (18727, -3.06 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (370.1 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.68 DPS) [quest]; Chains of the Lich (23125, -0.92 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 47.1 spell_power points (11.26 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.29 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.79 DPS) [crafted]; Darkspear Shoulderpads (272103, -3.00 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.7 spell_power points (7.58 DPS) | yes | Crystalline Threaded Cape (20697, -2.00 DPS) [world]; Hide of the Wild (18510, -2.24 DPS) [crafted]; Spritecaster Cape (11623, -3.04 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 57.6 spell_power points (13.77 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.67 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.54 DPS) [pvp]; Robe of Everlasting Night (18385, -4.72 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.7 spell_power points (6.85 DPS) | yes | Sublime Wristguards (18497, -1.99 DPS) [dungeon]; Runecloth Cuffs (254123, -2.23 DPS) [crafted]; Marshal's Silk Bracers (16438, -3.06 DPS) [pvp] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 32.8 spell_power points (7.82 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.4 spell_power points (12.53 DPS) | yes | Magician's Cord (272393, -3.19 DPS) [vendor]; Highlander's Cloth Girdle (20047, -6.17 DPS) [rep]; Belt of the Archmage (18405, -9.42 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.6 spell_power points (13.27 DPS) | yes | Marshal's Silk Leggings (231605, -0.30 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.05 DPS) [pvp]; Sorcerer's Leggings (226933, -3.38 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.4 spell_power points (8.21 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.72 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (370.1 DPS) | yes | Songstone of Ironforge (12543, -1.75 DPS) [quest]; Maiden's Circle (13001, -1.75 DPS) [world_drop]; Naglering (11669, -13.38 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (370.1 DPS) | yes | Songstone of Ironforge (12543, -0.72 DPS) [quest]; Maiden's Circle (13001, -0.72 DPS) [world_drop]; Naglering (11669, -13.53 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (370.1 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.92 DPS) [crafted] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (370.1 DPS) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -1.67 DPS) [vendor]; Serenity Field (272439, -3.58 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (370.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.05 DPS) [world]; Teebu's Blazing Longsword (1728, -18.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 327.3 spell_power points (78.19 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.72 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.14 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.87 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 28.3. Weights run: 1.1s. Verify run: 0.8s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.258 ± 0.009, crit=0.109 ± 0.004 per rating point (14 rating = 1%, 1.522 per %), hit=0.306 ± 0.002 per rating point (10 rating = 1%, 3.059 per %), spell_haste=not significant (-0.490 ± 0.165), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.52 DPS) | yes | Shadow Goggles (4373, -0.49 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.3 spell_power points (0.63 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.31 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.35 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.3 spell_power points (0.54 DPS) | yes | Green Woolen Vest (2582, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.20 DPS) [dungeon]; Gray Woolen Robe (2585, -0.65 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.5 spell_power points (0.13 DPS) | yes | Owlbeard Bracers (16981, -0.00 DPS) [quest]; Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.60 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Apothecary Gloves (10919, -0.26 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.0 spell_power points (0.43 DPS) | yes | Novice Ardent's Sash (253887, -0.19 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.42 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (28.3 DPS) | yes | Silk-threaded Trousers (1929, -0.05 DPS) [dungeon]; Rumpled Kilt (274741, -0.22 DPS) [vendor]; Abomination Skin Leggings (23173, -0.43 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.0 spell_power points (0.69 DPS) | yes | Feather Padded Treads (285345, -0.15 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.35 DPS) [crafted]; Pristine Boots (253889, -0.37 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.43 DPS) | yes | Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon]; Loop of Sacrifice (281673, -0.32 DPS) [quest]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.26 DPS) | yes | Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Loop of Sacrifice (281673, -0.15 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.6 spell_power points (0.22 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 261.4 spell_power points (22.57 DPS) | yes | Skycaller (12984, -0.69 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.28 DPS) [dungeon]; Sizzle Stick (8071, -4.49 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 44.7. Weights run: 1.2s. Verify run: 0.8s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.603 ± 0.020, crit=0.170 ± 0.009 per rating point (14 rating = 1%, 2.376 per %), hit=0.332 ± 0.003 per rating point (10 rating = 1%, 3.322 per %), spell_haste=not significant (-0.234 ± 0.295), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.0 spell_power points (1.15 DPS) | yes | Silk Headband (7050, -0.29 DPS) [crafted]; Embalmed Shroud (7691, -0.39 DPS) [dungeon]; Holy Shroud (2721, -0.52 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.6 spell_power points (1.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.78 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.78 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.04 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (1.38 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.48 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Cloak of Rot (4462, -0.02 DPS) [world]; Darkspear Raider's Cloak (272078, -0.02 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.8 spell_power points (1.61 DPS) | yes | Tree Bark Jacket (1486, -0.37 DPS) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted]; Death Speaker Robes (6682, -0.71 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Nightsky Wristbands (6407, -0.51 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.51 DPS) [quest]; Glowing Magical Bracelets (13106, -0.69 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.0 spell_power points (0.86 DPS) | yes | Truefaith Gloves (7049, -0.21 DPS) [crafted]; Gnoll Casting Gloves (892, -0.29 DPS) [world]; Serpent Gloves (5970, -0.45 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.8 spell_power points (1.22 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Crimson Silk Belt (7055, -0.25 DPS) [crafted]; Warsong Sash (16975, -0.35 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.8 spell_power points (1.32 DPS) | yes | Gaze Dreamer Pants (6903, -0.17 DPS) [dungeon]; Pristine Leggings (253987, -0.25 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.2 spell_power points (1.07 DPS) | yes | Spidersilk Boots (4320, -0.17 DPS) [crafted]; Boots of the Enchanter (4325, -0.59 DPS) [crafted]; Acidic Walkers (9454, -1.00 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.67 DPS) | yes | Black Widow Band (6199, -0.27 DPS) [world]; Snake Hoop (6750, -0.27 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.57 DPS) | yes | Snake Hoop (6750, -0.17 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Black Widow Band (6199, -0.65 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Glimmering Staff (249392, -0.23 DPS) [crafted]; Twisted Chanter's Staff (890, -0.28 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.28 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.6 spell_power points (1.01 DPS) | yes | Witch's Finger (16887, -0.61 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.63 DPS) [world]; Tome of the Darkspear Prophecy (272090, -1.05 DPS, sim-verified) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 350.7 spell_power points (33.50 DPS) | yes | Starfaller (13063, -0.41 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552100130103050-0000000000000000000)

Set DPS (verified): 76.0. Weights run: 1.2s. Verify run: 0.8s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.605 ± 0.032, crit=0.273 ± 0.016 per rating point (14 rating = 1%, 3.817 per %), hit=0.429 ± 0.004 per rating point (10 rating = 1%, 4.288 per %), spell_haste=not significant (0.673 ± 0.512), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.43 DPS) | yes | Augural Shroud (2620, -0.46 DPS) [world]; Living Cowl (5608, -0.93 DPS) [world]; Enchanter's Cowl (4322, -1.04 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.6 spell_power points (1.23 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.53 DPS) [quest]; Triune Amulet (7722, -0.74 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.74 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.9 spell_power points (1.72 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Berylline Pads (4197, -0.21 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.4 spell_power points (1.67 DPS) | yes | Guardian Cloak (5965, -0.63 DPS) [crafted]; Icy Cloak (4327, -0.86 DPS) [crafted]; Long Silken Cloak (4326, -1.62 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.25 DPS) [crafted]; Elemental Raiment (9434, -0.28 DPS) [world_drop]; Robe of the Magi (1716, -0.77 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.02 DPS) [quest]; Condor Bracers (15864, -0.23 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (2.37 DPS) | yes | Red Mageweave Gloves (10018, -0.39 DPS) [crafted]; Black Mageweave Gloves (10003, -0.63 DPS) [crafted]; Gilded Handwraps (254021, -0.95 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.36 DPS) [crafted]; Gilded Cord (254037, -0.37 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.89 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.3 spell_power points (2.46 DPS) | yes | Crimson Silk Pantaloons (7062, -0.63 DPS) [crafted]; Abomination Skin Leggings (23173, -0.86 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.78 DPS) | yes | Gilded Slippers (254001, -0.75 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.64 DPS) [dungeon]; Spidersilk Boots (4320, -1.69 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.6 spell_power points (1.58 DPS) | yes | Reedknot Ring (9622, -0.77 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.88 DPS) [vendor]; Black Widow Band (6199, -1.09 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.04 DPS) | yes | Reedknot Ring (9622, -0.23 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.35 DPS) [vendor]; Black Widow Band (6199, -0.55 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -1.27 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.27 DPS) [dungeon]; Gut Ripper (2164, -3.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 343.4 spell_power points (39.80 DPS) | yes | Nether Force Wand (11263, -0.84 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.40 DPS) [quest]; Ragefire Wand (7513, -2.45 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 131.2. Weights run: 1.2s. Verify run: 0.9s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.915 ± 0.039, crit=0.263 ± 0.021 per rating point (14 rating = 1%, 3.681 per %), hit=0.589 ± 0.007 per rating point (10 rating = 1%, 5.887 per %), spell_haste=not significant (0.096 ± 0.803), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.3 spell_power points (5.58 DPS) | yes | Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.29 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.54 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (3.22 DPS) | yes | Scorn's Icy Choker (23169, -1.35 DPS) [dungeon]; Mindburst Medallion (11196, -1.50 DPS) [quest]; Horizon Choker (13085, -3.66 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.5 spell_power points (4.41 DPS) | yes | Kentic Amice (11624, -0.53 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.31 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.43 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.2 spell_power points (3.03 DPS) | yes | Spritecaster Cape (11623, -0.11 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.45 DPS) [dungeon]; Runecloth Cloak (13860, -0.59 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.3 spell_power points (5.58 DPS) | yes | Runecloth Tunic (13857, -1.53 DPS) [crafted]; Runecloth Robe (13858, -1.61 DPS) [crafted]; Robe of the Magi (1716, -2.35 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.7 spell_power points (2.05 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.05 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.7 spell_power points (4.59 DPS) | yes | Raider Handwraps (272098, -0.53 DPS) [vendor]; Dreamweave Gloves (10019, -1.35 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.42 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.4 spell_power points (3.50 DPS) | yes | Satyrmane Sash (17755, -0.04 DPS) [dungeon]; Ban'thok Sash (11662, -0.11 DPS) [dungeon]; Deathmage Sash (10771, -0.40 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.2 spell_power points (4.81 DPS) | yes | Red Mageweave Pants (10009, -1.07 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -1.77 DPS, sim-verified) [vendor]; Crimson Silk Pantaloons (7062, -1.83 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.28 DPS) [vendor]; Gilded Sandals (254107, -0.71 DPS) [crafted]; Southsea Mojo Boots (20641, -0.89 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.32 DPS) | yes | Brainlash (6440, -0.26 DPS) [dungeon]; Band of the Unicorn (7553, -0.37 DPS) [world_drop]; Advisor's Ring (19519, -0.52 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.30 DPS) | yes | Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Advisor's Ring (19519, -0.51 DPS) [rep]; Brainlash (6440, -2.04 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (131.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.18 DPS) [quest] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (131.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (131.2 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -0.82 DPS) [quest]; Spellforce Rod (1664, -0.98 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 350.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.75 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 349.7. Weights run: 1.3s. Verify run: 1.0s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.836 ± 0.062, crit=0.543 ± 0.037 per rating point (14 rating = 1%, 7.601 per %), hit=0.905 ± 0.013 per rating point (10 rating = 1%, 9.047 per %), spell_haste=not significant (-4.127 ± 1.161), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 49.5 spell_power points (11.83 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.40 DPS) [pvp]; Crimson Felt Hat (18727, -3.06 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (349.7 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.68 DPS) [quest]; Chains of the Lich (23125, -0.92 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 47.1 spell_power points (11.26 DPS) | yes | Warlord's Silk Amice (231594, -2.29 DPS) [pvp]; Darkspear Shoulderpads (272103, -3.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -7.50 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.7 spell_power points (7.58 DPS) | yes | Crystalline Threaded Cape (20697, -2.00 DPS) [world]; Hide of the Wild (18510, -2.24 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.92 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 57.6 spell_power points (13.77 DPS) | yes | Warlord's Silk Raiment (231596, -0.67 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.54 DPS) [pvp]; Robe of Everlasting Night (18385, -4.72 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.7 spell_power points (6.85 DPS) | yes | Sublime Wristguards (18497, -1.99 DPS) [dungeon]; Runecloth Cuffs (254123, -2.23 DPS) [crafted]; General's Silk Cuffs (16538, -3.06 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 32.8 spell_power points (7.82 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.4 spell_power points (12.53 DPS) | yes | Magician's Cord (272393, -3.19 DPS) [vendor]; Defiler's Cloth Girdle (20163, -6.17 DPS) [rep]; Belt of the Archmage (18405, -7.46 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.6 spell_power points (13.27 DPS) | yes | General's Silk Trousers (231595, -0.30 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.79 DPS) [rep]; Legionnaire's Silk Legguards (227107, -3.05 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.4 spell_power points (8.21 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.72 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (349.7 DPS) | yes | Eye of Orgrimmar (12545, -1.75 DPS) [quest]; Maiden's Circle (13001, -1.75 DPS) [world_drop]; Naglering (11669, -11.19 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (349.7 DPS) | yes | Eye of Orgrimmar (12545, -0.72 DPS) [quest]; Maiden's Circle (13001, -0.72 DPS) [world_drop]; Naglering (11669, -9.87 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (349.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (349.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (349.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.54 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -19.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 327.3 spell_power points (78.19 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.72 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.14 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.87 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

