# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 31.3. Weights run: 1.2s. Verify run: 0.9s. 149 eligible items had no known source.

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

### Band 30 (gnome, 000000000000000000-23552100030000000-0000000000000000000)

Set DPS (verified): 50.7. Weights run: 1.3s. Verify run: 0.9s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.746 ± 0.022, crit=0.173 ± 0.009 per rating point (14 rating = 1%, 2.425 per %), hit=0.348 ± 0.003 per rating point (10 rating = 1%, 3.479 per %), spell_haste=not significant (0.636 ± 0.392), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.5 spell_power points (1.28 DPS) | yes | Holy Shroud (2721, -0.23 DPS) [world_drop]; Silk Headband (7050, -0.42 DPS) [crafted]; Nightsky Cowl (4039, -0.43 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.5 spell_power points (1.09 DPS) | yes | Darkspear Warding Pendant (272075, -0.76 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.81 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.81 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.7 spell_power points (1.49 DPS) | yes | Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 6.0 spell_power points (0.57 DPS) | yes | Cloak of Rot (4462, -0.00 DPS) [world]; Darkspear Raider's Cloak (272078, -0.00 DPS) [vendor]; Hillman's Cloak (3719, -0.09 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.7 spell_power points (1.78 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Tree Bark Jacket (1486, -0.54 DPS) [dungeon]; Pristine Gown (253961, -0.62 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Nightsky Wristbands (6407, -0.43 DPS) [world_drop]; Stonecloth Bindings (14416, -0.50 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.98 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 12.2 spell_power points (1.16 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.49 DPS) [dungeon]; Shilly Mitts (9609, -0.49 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.2 spell_power points (1.26 DPS) | yes | Belt of Arugal (6392, +0.00 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.19 DPS) [crafted]; Invoker's Cord (215366, -0.24 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gaze Dreamer Pants (6903, -0.02 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.17 DPS) [crafted]; Abomination Skin Leggings (23173, -1.14 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.2 spell_power points (1.16 DPS) | yes | Spidersilk Boots (4320, -0.21 DPS) [crafted]; Nimbus Boots (6998, -0.59 DPS) [quest]; Acidic Walkers (9454, -0.78 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.17 DPS) [world]; Snake Hoop (6750, -0.17 DPS) [quest]; Minor Channeling Ring (1449, -1.02 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.07 DPS) [world]; Snake Hoop (6750, -0.07 DPS) [quest]; Minor Channeling Ring (1449, -0.75 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Twisted Chanter's Staff (890, -0.15 DPS) [world_drop]; Channeler's Staff (4437, -0.29 DPS) [world]; Glimmering Staff (249392, -0.51 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.5 spell_power points (1.09 DPS) | yes | Dwarven Tome (279898, -0.62 DPS) [quest]; Tome of the Darkspear Prophecy (272090, -0.62 DPS) [vendor]; Satyr's Rod (15962, -0.66 DPS) [world_drop] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 352.7 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.35 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-23552100030023050-0000000000000000000)

Set DPS (verified): 83.1. Weights run: 1.3s. Verify run: 0.8s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.660 ± 0.033, crit=0.243 ± 0.020 per rating point (14 rating = 1%, 3.404 per %), hit=0.425 ± 0.005 per rating point (10 rating = 1%, 4.247 per %), spell_haste=not significant (0.426 ± 0.567), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.46 DPS) | yes | Augural Shroud (2620, -0.40 DPS) [world]; Living Cowl (5608, -0.94 DPS) [world]; Enchanter's Cowl (4322, -0.99 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 spell_power points (1.29 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.51 DPS) [quest]; Triune Amulet (7722, -0.74 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.74 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 spell_power points (1.83 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.9 spell_power points (1.75 DPS) | yes | Guardian Cloak (5965, -0.66 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.90 DPS) [vendor]; Long Silken Cloak (4326, -1.77 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 spell_power points (3.05 DPS) | yes | Dreamweave Vest (10021, -0.24 DPS) [crafted]; Robe of Power (7054, -0.47 DPS) [crafted]; Elemental Raiment (9434, -0.58 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (83.1 DPS) | yes | Condor Bracers (15864, -0.23 DPS) [quest]; Windchaser Cuffs (14429, -0.36 DPS) [world_drop]; Arcane Runed Bracers (4744, -0.90 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (2.42 DPS) | yes | Red Mageweave Gloves (10018, -0.36 DPS) [crafted]; Black Mageweave Gloves (10003, -0.66 DPS) [crafted]; Gilded Handwraps (254021, -0.94 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.9 spell_power points (1.98 DPS) | yes | Highlander's Cloth Girdle (20098, -0.03 DPS) [rep]; Gilded Cord (254037, -0.42 DPS) [crafted]; Star Belt (4329, -0.46 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.9 spell_power points (2.57 DPS) | yes | Crimson Silk Pantaloons (7062, -0.63 DPS) [crafted]; Abomination Skin Leggings (23173, -0.90 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.13 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.82 DPS) | yes | Gilded Slippers (254001, -0.80 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.61 DPS) [dungeon]; Spidersilk Boots (4320, -1.68 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.64 DPS) | yes | Ring of Forlorn Spirits (2043, -0.70 DPS) [quest]; Reedknot Ring (9622, -0.82 DPS) [quest]; Minor Channeling Ring (1449, -0.90 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.06 DPS) | yes | Ring of Forlorn Spirits (2043, -0.12 DPS) [quest]; Reedknot Ring (9622, -0.23 DPS) [quest]; Minor Channeling Ring (1449, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -1.18 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.29 DPS) [dungeon]; Gut Ripper (2164, -3.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 339.6 spell_power points (39.84 DPS) | yes | Nether Force Wand (11263, -1.61 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.42 DPS) [quest]; Ragefire Wand (7513, -2.47 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 205011000000000000-23552100030023051-0000000000000000000)

Set DPS (verified): 145.8. Weights run: 1.2s. Verify run: 1.0s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.623 ± 0.041, crit=0.251 ± 0.023 per rating point (14 rating = 1%, 3.514 per %), hit=0.593 ± 0.007 per rating point (10 rating = 1%, 5.934 per %), spell_haste=5.274 ± 0.809, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 31.5 spell_power points (4.60 DPS) | yes | Dreamweave Circlet (10041, -0.62 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.65 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -0.94 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.7 spell_power points (2.88 DPS) | yes | Mindburst Medallion (11196, -1.46 DPS) [quest]; Horizon Choker (13085, -1.61 DPS) [world_drop]; Scorn's Icy Choker (23169, -2.78 DPS, sim-verified) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -0.73 DPS) [vendor]; Red Mageweave Shoulders (10029, -0.84 DPS) [crafted]; Rotgrip Mantle (17732, -2.05 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.7 spell_power points (2.59 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.46 DPS) [dungeon]; Runecloth Cloak (13860, -0.55 DPS) [crafted]; Big Voodoo Cloak (8216, -1.04 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 31.5 spell_power points (4.60 DPS) | yes | Robe of the Magi (1716, -0.84 DPS) [world_drop]; Runecloth Tunic (13857, -1.11 DPS) [crafted]; Dreamweave Vest (10021, -1.15 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.4 spell_power points (1.66 DPS) | yes | Aristocratic Cuffs (12546, -0.30 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.34 DPS) [quest]; Bloodband Bracers (11469, -2.63 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 26.6 spell_power points (3.89 DPS) | yes | Raider Handwraps (272098, -0.87 DPS) [vendor]; Dreamweave Gloves (10019, -0.90 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.18 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Dawnspire Cord (12466, -0.24 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.43 DPS) [rep]; Satyrmane Sash (17755, -2.02 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.2 spell_power points (4.27 DPS) | yes | Red Mageweave Pants (10009, -1.13 DPS) [crafted]; Wizardweave Leggings (14132, -1.49 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.20 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.51 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.65 DPS) [vendor]; Gilded Sandals (254107, -1.08 DPS) [crafted]; Black Mageweave Boots (10026, -1.26 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (2.01 DPS) | yes | Band of the Unicorn (7553, -0.11 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.25 DPS) [rep]; Brainlash (6440, -0.64 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.4 spell_power points (1.95 DPS) | yes | Band of the Unicorn (7553, -0.05 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.20 DPS) [rep]; Brainlash (6440, -0.59 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -0.28 DPS) [world_drop]; Spellshifter Rod (9527, -0.83 DPS) [quest]; Blade of Eternal Darkness (17780, -2.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 359.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 205015100000000000-23552100030023051-0050000000000000000)

Set DPS (verified): 352.7. Weights run: 1.4s. Verify run: 0.9s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.854 ± 0.058, crit=0.448 ± 0.032 per rating point (14 rating = 1%, 6.277 per %), hit=0.887 ± 0.012 per rating point (10 rating = 1%, 8.867 per %), spell_haste=not significant (2.472 ± 1.154), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 48.6 spell_power points (11.40 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.40 DPS) [pvp]; Crimson Felt Hat (18727, -2.77 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (352.7 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.67 DPS) [quest]; Chains of the Lich (23125, -0.96 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.1 spell_power points (10.81 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.94 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.75 DPS) [crafted]; Darkspear Shoulderpads (272103, -5.69 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.7 spell_power points (7.43 DPS) | yes | Crystalline Threaded Cape (20697, -1.94 DPS) [world]; Hide of the Wild (18510, -2.15 DPS) [crafted]; Spritecaster Cape (11623, -2.95 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.5 spell_power points (13.26 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.64 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.45 DPS) [pvp]; Robe of Everlasting Night (18385, -4.32 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.8 spell_power points (6.76 DPS) | yes | Sublime Wristguards (18497, -1.94 DPS) [dungeon]; Runecloth Cuffs (254123, -2.18 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.96 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 32.8 spell_power points (7.70 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.7 spell_power points (12.35 DPS) | yes | Magician's Cord (272393, -3.09 DPS) [vendor]; Stormpike Cloth Girdle (19094, -6.12 DPS) [rep]; Belt of the Archmage (18405, -10.64 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 53.1 spell_power points (12.45 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -2.65 DPS) [pvp]; Sorcerer's Leggings (226933, -2.71 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.7 spell_power points (8.13 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.70 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (352.7 DPS) | yes | Songstone of Ironforge (12543, -1.74 DPS) [quest]; Maiden's Circle (13001, -1.74 DPS) [world_drop]; Naglering (11669, -8.14 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (352.7 DPS) | yes | Songstone of Ironforge (12543, -0.70 DPS) [quest]; Maiden's Circle (13001, -0.70 DPS) [world_drop]; Naglering (11669, -7.95 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (352.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (352.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -7.91 DPS, sim-verified) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (352.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.09 DPS) [world]; Teebu's Blazing Longsword (1728, -17.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 333.3 spell_power points (78.17 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.74 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.16 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.86 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 28.3. Weights run: 1.2s. Verify run: 0.9s. 138 eligible items had no known source.

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

### Band 30 (orc, 000000000000000000-23552100030000000-0000000000000000000)

Set DPS (verified): 46.2. Weights run: 1.3s. Verify run: 1.0s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.746 ± 0.022, crit=0.173 ± 0.009 per rating point (14 rating = 1%, 2.425 per %), hit=0.348 ± 0.003 per rating point (10 rating = 1%, 3.479 per %), spell_haste=not significant (0.636 ± 0.392), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.5 spell_power points (1.28 DPS) | yes | Holy Shroud (2721, -0.23 DPS) [world_drop]; Silk Headband (7050, -0.42 DPS) [crafted]; Nightsky Cowl (4039, -0.43 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.5 spell_power points (1.09 DPS) | yes | Crystal Starfire Medallion (5003, -0.81 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.81 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.93 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.7 spell_power points (1.49 DPS) | yes | Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.0 spell_power points (0.57 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Hillman's Cloak (3719, -0.09 DPS) [crafted]; Windsong Drape (15468, -0.09 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.7 spell_power points (1.78 DPS) | yes | Tree Bark Jacket (1486, -0.54 DPS) [dungeon]; Pristine Gown (253961, -0.62 DPS) [crafted]; Death Speaker Robes (6682, -0.63 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Nightsky Wristbands (6407, -0.43 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.43 DPS) [quest]; Glowing Magical Bracelets (13106, -0.51 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.7 spell_power points (0.92 DPS) | yes | Truefaith Gloves (7049, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Pristine Gloves (253913, -0.33 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.2 spell_power points (1.26 DPS) | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Crimson Silk Belt (7055, -0.19 DPS) [crafted]; Warsong Sash (16975, -0.21 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (46.2 DPS) | yes | Gaze Dreamer Pants (6903, -0.02 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.17 DPS) [crafted]; Abomination Skin Leggings (23173, -0.75 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.2 spell_power points (1.16 DPS) | yes | Spidersilk Boots (4320, -0.21 DPS) [crafted]; Pristine Boots (253889, -0.66 DPS) [crafted]; Acidic Walkers (9454, -0.88 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.17 DPS) [world]; Snake Hoop (6750, -0.17 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.57 DPS) | yes | Snake Hoop (6750, -0.07 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Black Widow Band (6199, -0.45 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Glimmering Staff (249392, -0.07 DPS) [crafted]; Twisted Chanter's Staff (890, -0.15 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.15 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.5 spell_power points (1.09 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.62 DPS) [vendor]; Satyr's Rod (15962, -0.66 DPS) [world_drop]; Witch's Finger (16887, -0.69 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 352.7 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.35 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552100030023050-0000000000000000000)

Set DPS (verified): 75.5. Weights run: 1.3s. Verify run: 0.8s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.660 ± 0.033, crit=0.243 ± 0.020 per rating point (14 rating = 1%, 3.404 per %), hit=0.425 ± 0.005 per rating point (10 rating = 1%, 4.247 per %), spell_haste=not significant (0.426 ± 0.567), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.46 DPS) | yes | Augural Shroud (2620, -0.40 DPS) [world]; Living Cowl (5608, -0.94 DPS) [world]; Enchanter's Cowl (4322, -0.99 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 spell_power points (1.29 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.51 DPS) [quest]; Triune Amulet (7722, -0.74 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.74 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 spell_power points (1.83 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.9 spell_power points (1.75 DPS) | yes | Guardian Cloak (5965, -0.66 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.90 DPS) [vendor]; Long Silken Cloak (4326, -1.45 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 spell_power points (3.05 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.47 DPS) [crafted]; Elemental Raiment (9434, -0.58 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.3 spell_power points (1.09 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.03 DPS) [dungeon]; Condor Bracers (15864, -0.27 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (2.42 DPS) | yes | Black Mageweave Gloves (10003, -0.66 DPS) [crafted]; Red Mageweave Gloves (10018, -0.74 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -0.94 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.9 spell_power points (1.98 DPS) | yes | Defiler's Cloth Girdle (20166, -0.03 DPS) [rep]; Gilded Cord (254037, -0.42 DPS) [crafted]; Star Belt (4329, -0.46 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.9 spell_power points (2.57 DPS) | yes | Abomination Skin Leggings (23173, -0.90 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.90 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.13 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.82 DPS) | yes | Gilded Slippers (254001, -1.28 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.61 DPS) [dungeon]; Spidersilk Boots (4320, -1.68 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.64 DPS) | yes | Reedknot Ring (9622, -0.82 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.93 DPS) [vendor]; Black Widow Band (6199, -1.10 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.06 DPS) | yes | Sea Giant's Toe Ring (274746, -0.35 DPS) [vendor]; Black Widow Band (6199, -0.51 DPS) [world]; Reedknot Ring (9622, -1.61 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (75.5 DPS) | yes | Windweaver Staff (7757, -1.18 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.29 DPS) [dungeon]; Gut Ripper (2164, -3.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 339.6 spell_power points (39.84 DPS) | yes | Nether Force Wand (11263, -1.02 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.42 DPS) [quest]; Ragefire Wand (7513, -2.47 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 205011000000000000-23552100030023051-0000000000000000000)

Set DPS (verified): 132.1. Weights run: 1.2s. Verify run: 1.0s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.623 ± 0.041, crit=0.251 ± 0.023 per rating point (14 rating = 1%, 3.514 per %), hit=0.593 ± 0.007 per rating point (10 rating = 1%, 5.934 per %), spell_haste=5.274 ± 0.809, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 31.5 spell_power points (4.60 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.65 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -0.94 DPS) [vendor]; Dreamweave Circlet (10041, -1.71 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.7 spell_power points (2.88 DPS) | yes | Mindburst Medallion (11196, -1.46 DPS) [quest]; Horizon Choker (13085, -1.61 DPS) [world_drop]; Scorn's Icy Choker (23169, -2.58 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 24.2 spell_power points (3.54 DPS) | yes | Kentic Amice (11624, -0.31 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.04 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.7 spell_power points (2.59 DPS) | yes | Deep Woodlands Cloak (19121, -0.02 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.46 DPS) [dungeon]; Runecloth Cloak (13860, -0.55 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 31.5 spell_power points (4.60 DPS) | yes | Runecloth Tunic (13857, -1.11 DPS) [crafted]; Dreamweave Vest (10021, -1.15 DPS) [crafted]; Robe of the Magi (1716, -2.43 DPS, sim-verified) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.4 spell_power points (1.66 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.11 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 26.6 spell_power points (3.89 DPS) | yes | Raider Handwraps (272098, -0.87 DPS) [vendor]; Dreamweave Gloves (10019, -0.90 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.18 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (132.1 DPS) | yes | Dawnspire Cord (12466, -0.24 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.43 DPS) [rep]; Satyrmane Sash (17755, -2.26 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.2 spell_power points (4.27 DPS) | yes | Red Mageweave Pants (10009, -1.13 DPS) [crafted]; Wizardweave Leggings (14132, -1.49 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.91 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.51 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.65 DPS) [vendor]; Gilded Sandals (254107, -1.08 DPS) [crafted]; Black Mageweave Boots (10026, -1.26 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (2.01 DPS) | yes | Band of the Unicorn (7553, -0.11 DPS) [world_drop]; Advisor's Ring (19519, -0.25 DPS) [rep]; Brainlash (6440, -0.64 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.4 spell_power points (1.95 DPS) | yes | Advisor's Ring (19519, -0.20 DPS) [rep]; Brainlash (6440, -0.59 DPS) [dungeon]; Band of the Unicorn (7553, -2.61 DPS, sim-verified) [world_drop] |
| trinket1 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 359.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Mark of the Chosen; trinket2: Smoking Heart of the Mountain; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 205015100000000000-23552100030023051-0050000000000000000)

Set DPS (verified): 338.3. Weights run: 1.4s. Verify run: 1.0s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.854 ± 0.058, crit=0.448 ± 0.032 per rating point (14 rating = 1%, 6.277 per %), hit=0.887 ± 0.012 per rating point (10 rating = 1%, 8.867 per %), spell_haste=not significant (2.472 ± 1.154), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 48.6 spell_power points (11.40 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.40 DPS) [pvp]; Crimson Felt Hat (18727, -6.55 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (338.3 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.67 DPS) [quest]; Chains of the Lich (23125, -0.96 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.1 spell_power points (10.81 DPS) | yes | Warlord's Silk Amice (231594, -1.94 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.75 DPS) [crafted]; Darkspear Shoulderpads (272103, -10.99 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.7 spell_power points (7.43 DPS) | yes | Hide of the Wild (18510, -2.15 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.82 DPS) [quest]; Crystalline Threaded Cape (20697, -8.49 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.5 spell_power points (13.26 DPS) | yes | Warlord's Silk Raiment (231596, -0.64 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.45 DPS) [pvp]; Robe of Everlasting Night (18385, -7.42 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.8 spell_power points (6.76 DPS) | yes | Sublime Wristguards (18497, -1.94 DPS) [dungeon]; Runecloth Cuffs (254123, -2.18 DPS) [crafted]; General's Silk Cuffs (16538, -2.96 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 32.8 spell_power points (7.70 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.7 spell_power points (12.35 DPS) | yes | Magician's Cord (272393, -3.09 DPS) [vendor]; Frostwolf Cloth Belt (19090, -6.12 DPS) [rep]; Belt of the Archmage (18405, -7.38 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 53.1 spell_power points (12.45 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -2.65 DPS) [pvp]; Outrider's Silk Leggings (22747, -7.07 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.7 spell_power points (8.13 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.70 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (338.3 DPS) | yes | Eye of Orgrimmar (12545, -1.74 DPS) [quest]; Maiden's Circle (13001, -1.74 DPS) [world_drop]; Naglering (11669, -11.82 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (338.3 DPS) | yes | Eye of Orgrimmar (12545, -0.70 DPS) [quest]; Maiden's Circle (13001, -0.70 DPS) [world_drop]; Naglering (11669, -12.51 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (338.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Smolderweb's Eye (13213, -12.85 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (338.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Smolderweb's Eye (13213, -11.99 DPS, sim-verified) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (338.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.53 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -20.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 333.3 spell_power points (78.17 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.74 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.16 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.86 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

