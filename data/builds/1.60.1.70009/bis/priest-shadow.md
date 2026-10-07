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

### Band 30 (gnome, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 55.2. Weights run: 0.9s. Verify run: 0.7s. 244 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.468 ± 0.009, crit=0.045 ± 0.002 per rating point (14 rating = 1%, 0.624 per %), hit=0.192 ± 0.002 per rating point (10 rating = 1%, 1.923 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.19 DPS) [crafted]; Embalmed Shroud (7691, -0.30 DPS) [dungeon]; Holy Shroud (2721, -0.56 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.10 DPS) | yes | Crystal Starfire Medallion (5003, -0.89 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.89 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.05 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (1.49 DPS) | yes | Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Invoker's Mantle (215365, -0.44 DPS) [crafted] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.10 DPS) [crafted]; Prelacy Cape (7004, -0.10 DPS) [quest]; Hillman's Cloak (3719, -0.58 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.1 spell_power points (1.70 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted]; Tree Bark Jacket (1486, -1.36 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.01 DPS) | yes | Nightsky Wristbands (6407, -0.70 DPS) [world_drop]; Stonecloth Bindings (14416, -0.75 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.95 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 9.2 spell_power points (1.03 DPS) | yes | Serpent Gloves (5970, -0.24 DPS) [dungeon]; Truefaith Gloves (7049, -0.31 DPS) [crafted]; Shilly Mitts (9609, -0.94 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.4 spell_power points (1.39 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Invoker's Cord (215366, -0.34 DPS) [crafted]; Crimson Silk Belt (7055, -0.35 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.7 spell_power points (1.43 DPS) | yes | Pristine Leggings (253987, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.44 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.65 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.16 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Nimbus Boots (6998, -0.48 DPS) [quest]; Spidersilk Boots (4320, -1.43 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.79 DPS) | yes | Minor Channeling Ring (1449, -0.12 DPS) [quest]; Black Widow Band (6199, -0.42 DPS) [world]; Snake Hoop (6750, -0.42 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.67 DPS) | yes | Black Widow Band (6199, -0.31 DPS) [world]; Snake Hoop (6750, -0.31 DPS) [quest]; Minor Channeling Ring (1449, -0.67 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.01 DPS) | yes | Glimmering Staff (249392, -0.34 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.49 DPS) [world_drop]; Channeler's Staff (4437, -0.59 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.8 spell_power points (1.10 DPS) | yes | Eye of Paleth (2943, -0.65 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.65 DPS) [world]; Dwarven Tome (279898, -1.05 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 298.3 spell_power points (33.55 DPS) | yes | Starfaller (13063, -0.48 DPS) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 97.9. Weights run: 1.0s. Verify run: 0.7s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.835 ± 0.014, crit=0.041 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), hit=0.289 ± 0.004 per rating point (10 rating = 1%, 2.886 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.86 DPS) | yes | Augural Shroud (2620, -0.22 DPS) [world]; Corpseshroud (10574, -0.70 DPS) [dungeon]; Enchanter's Cowl (4322, -0.90 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.0 spell_power points (1.63 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.72 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.84 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.84 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.9 spell_power points (2.43 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.34 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.5 spell_power points (2.25 DPS) | yes | Guardian Cloak (5965, -0.86 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.00 DPS) [vendor]; Long Silken Cloak (4326, -1.56 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.0 spell_power points (3.67 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.41 DPS) [crafted]; Elemental Raiment (9434, -0.82 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.22 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.20 DPS) [world_drop]; Condor Bracers (15864, -0.27 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (2.90 DPS) | yes | Red Mageweave Gloves (10018, -0.76 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.86 DPS) [crafted]; Stormcloth Gloves (10011, -1.00 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.5 spell_power points (2.66 DPS) | yes | Gilded Cord (254037, -0.66 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.80 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.0 spell_power points (3.27 DPS) | yes | Crimson Silk Pantaloons (7062, -0.98 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.41 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.26 DPS) | yes | Gilded Slippers (254001, -1.00 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.67 DPS) [dungeon]; Spidersilk Boots (4320, -1.86 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.0 spell_power points (2.04 DPS) | yes | Ring of Forlorn Spirits (2043, -0.95 DPS) [quest]; Reedknot Ring (9622, -1.09 DPS) [quest]; Minor Channeling Ring (1449, -1.13 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.22 DPS) | yes | Reedknot Ring (9622, -0.27 DPS) [quest]; Minor Channeling Ring (1449, -0.32 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.14 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (97.9 DPS) | yes | Windweaver Staff (7757, -1.02 DPS) [dungeon]; Staff of Jordan (873, -1.47 DPS) [world_drop]; Gut Ripper (2164, -4.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 294.6 spell_power points (40.06 DPS) | yes | Umbral Wand (5216, -1.06 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.47 DPS) [dungeon]; Twisted Nether Wand (249144, -5.25 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 145.5. Weights run: 0.9s. Verify run: 0.8s. 423 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.484 ± 0.024, crit=0.052 ± 0.002 per rating point (14 rating = 1%, 0.724 per %), hit=0.436 ± 0.006 per rating point (10 rating = 1%, 4.357 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.7 spell_power points (6.17 DPS) | yes | Soulcatcher Halo (10630, -1.47 DPS) [dungeon]; Dreamweave Circlet (10041, -1.63 DPS) [crafted]; Chief Architect's Monocle (11839, -3.10 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.8 spell_power points (2.63 DPS) | yes | Scorn's Icy Choker (23169, -0.62 DPS) [dungeon]; Mindburst Medallion (11196, -0.74 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.75 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.7 spell_power points (5.03 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.33 DPS) [crafted]; Inquisitor's Shawl (19507, -1.70 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.9 spell_power points (2.90 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.07 DPS) [dungeon]; Runecloth Cloak (13860, -0.26 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.27 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.7 spell_power points (6.17 DPS) | yes | Runecloth Robe (13858, -1.58 DPS) [crafted]; Runecloth Tunic (13857, -1.95 DPS) [crafted]; Robes of Insight (940, -2.71 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 22.3 spell_power points (2.82 DPS) | yes | Bloodband Bracers (11469, -0.50 DPS) [quest]; Forgotten Wraps (9433, -0.56 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.56 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.7 spell_power points (5.03 DPS) | yes | Virtuous Hands (226958, -0.17 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.69 DPS) [vendor]; Red Mageweave Gloves (10018, -1.75 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 34.2 spell_power points (4.34 DPS) | yes | Satyrmane Sash (17755, -0.68 DPS) [dungeon]; Ban'thok Sash (11662, -0.69 DPS) [dungeon]; Deathmage Sash (10771, -0.84 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.8 spell_power points (4.80 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.93 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.03 DPS) [dungeon]; Red Mageweave Pants (10009, -3.18 DPS, sim-verified) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (145.5 DPS) | yes | Southsea Mojo Boots (20641, -0.00 DPS) [quest]; Earthen Silk Slippers (254013, -0.05 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -1.61 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 22.3 spell_power points (2.82 DPS) | yes | Philanthropist's Ring (281635, -0.43 DPS) [quest]; Mindseye Circle (10634, -0.56 DPS) [dungeon]; Woodseed Hoop (17768, -1.13 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.4 spell_power points (2.46 DPS) | yes | Mindseye Circle (10634, -0.20 DPS) [dungeon]; Woodseed Hoop (17768, -0.76 DPS) [quest]; Philanthropist's Ring (281635, -0.88 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, -1.29 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -1.13 DPS) [quest]; Radiant Staff (249453, -1.88 DPS) [crafted]; Barman Shanker (12791, -6.49 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 414.2 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 299.4. Weights run: 1.1s. Verify run: 0.8s. 1129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.174 ± 0.031, crit=0.116 ± 0.004 per rating point (14 rating = 1%, 1.620 per %), hit=0.753 ± 0.009 per rating point (10 rating = 1%, 7.527 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 48.8 spell_power points (7.34 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Magister's Crown (16686, -3.79 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.3 spell_power points (4.56 DPS) | yes | Beads of Ogre Mojo (22149, -0.48 DPS) [quest]; Archlight Talisman (15856, -0.98 DPS) [quest]; Lady Maye's Pendant (14558, -1.20 DPS) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.2 spell_power points (6.96 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.38 DPS) [vendor]; Virtuous Epaulets (226955, -0.85 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.9 spell_power points (4.96 DPS) | yes | Hide of the Wild (18510, -1.08 DPS) [crafted]; Crystalline Threaded Cape (20697, -1.24 DPS) [world]; Spritecaster Cape (11623, -1.79 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 47.1 spell_power points (7.10 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -3.69 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.4 spell_power points (4.73 DPS) | yes | Sublime Wristguards (18497, -1.15 DPS) [dungeon]; Runecloth Cuffs (254123, -1.30 DPS) [crafted]; Marshal's Satin Bracers (17606, -1.37 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 39.7 spell_power points (5.98 DPS) | yes | Marshal's Satin Gloves (17608, -0.22 DPS) [vendor]; Marshal's Satin Grips (231617, -0.22 DPS) [vendor]; Virtuous Hands (226958, -0.82 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.4 spell_power points (8.79 DPS) | yes | Magister's Belt (16685, -4.02 DPS) [dungeon]; Dustfeather Sash (12589, -4.25 DPS) [dungeon]; Belt of the Archmage (18405, -5.15 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (22752) | Silverwing Sentinels [rep] | 50.3 spell_power points (7.58 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Marshal's Satin Legguards (231626, -0.85 DPS) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 40.8 spell_power points (6.14 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -1.86 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (299.4 DPS) | yes | Songstone of Ironforge (12543, -1.31 DPS) [quest]; Maiden's Circle (13001, -1.31 DPS) [world_drop]; Naglering (11669, -9.42 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (299.4 DPS) | yes | Songstone of Ironforge (12543, -0.45 DPS) [quest]; Maiden's Circle (13001, -0.45 DPS) [world_drop]; Naglering (11669, -9.97 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (299.4 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -4.15 DPS, sim-verified) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (299.4 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (299.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.47 DPS) [world]; Hand of Edward the Odd (2243, -7.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 515.8 spell_power points (77.66 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -1.91 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.67 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.99 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Second Wind; trinket2: Burst of Knowledge; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 574.5. Weights run: 1.2s. Verify run: 0.8s. 1129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.737 ± 0.018, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.610 per %), hit=0.520 ± 0.009 per rating point (10 rating = 1%, 5.204 per %), spell_haste=-2.160 ± 0.208, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 39.9 spell_power points (14.13 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 24.6 spell_power points (8.71 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Orb of the Darkmoon (19426, -0.91 DPS) [quest]; Beads of Ogre Mojo (22149, -0.97 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 38.7 spell_power points (13.70 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Virtuous Epaulets (226955, -1.99 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.40 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.1 spell_power points (9.60 DPS) | yes | Crystalline Threaded Cape (20697, -1.47 DPS) [world]; Hide of the Wild (18510, -2.03 DPS) [crafted]; Spritecaster Cape (11623, -3.07 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 39.3 spell_power points (13.91 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Robe of Everlasting Night (18385, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 27.9 spell_power points (9.88 DPS) | yes | Sublime Wristguards (18497, -3.02 DPS) [dungeon]; Runecloth Cuffs (254123, -3.37 DPS) [crafted]; Virtuous Wraps (226953, -4.18 DPS) [vendor] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 30.4 spell_power points (10.78 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Virtuous Hands (226958, -5.85 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 46.4 spell_power points (16.44 DPS) | yes | Belt of the Archmage (18405, -5.02 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -7.46 DPS) [rep]; Magister's Belt (16685, -8.48 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (22752) | Silverwing Sentinels [rep] | 42.0 spell_power points (14.88 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Skyshroud Leggings (13170, -0.75 DPS) [dungeon] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 33.8 spell_power points (11.97 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -1.42 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -2.46 DPS) [quest]; Maiden's Circle (13001, -2.46 DPS) [world_drop]; Naglering (11669, -8.54 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.06 DPS) [quest]; Maiden's Circle (13001, -1.06 DPS) [world_drop]; Naglering (11669, -8.33 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+18.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.48 DPS) [vendor]; Serenity Field (272439, -4.12 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -6.02 DPS) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.99 DPS) [world]; Hand of Edward the Odd (2243, -9.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (574.5 DPS) | yes | Bonecreeper Stylus (13938, -0.18 DPS) [dungeon]; Sparkling Crystal Wand (20672, -1.30 DPS) [world]; Torch of Light (279246, -9.09 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Crackling Staff; ranged: Ritssyn's Wand of Bad Mojo

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

### Band 30 (undead, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 53.6. Weights run: 0.9s. Verify run: 0.7s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.468 ± 0.009, crit=0.045 ± 0.002 per rating point (14 rating = 1%, 0.624 per %), hit=0.192 ± 0.002 per rating point (10 rating = 1%, 1.923 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.24 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.22 DPS) [crafted]; Embalmed Shroud (7691, -0.34 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.10 DPS) | yes | Darkspear Warding Pendant (272075, -0.85 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.89 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.89 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (1.49 DPS) | yes | Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Invoker's Mantle (215365, -0.44 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.56 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.1 spell_power points (1.70 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted]; Tree Bark Jacket (1486, -1.14 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.01 DPS) | yes | Glowing Magical Bracelets (13106, -0.45 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.70 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.70 DPS) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.3 spell_power points (0.94 DPS) | yes | Truefaith Gloves (7049, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.26 DPS) [world]; Serpent Gloves (5970, -0.55 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.4 spell_power points (1.39 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Warsong Sash (16975, -0.26 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.34 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.7 spell_power points (1.43 DPS) | yes | Gaze Dreamer Pants (6903, -0.08 DPS) [dungeon]; Pristine Leggings (253987, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.16 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Boots of the Enchanter (4325, -0.59 DPS) [crafted]; Spidersilk Boots (4320, -1.09 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.79 DPS) | yes | Black Widow Band (6199, -0.42 DPS) [world]; Snake Hoop (6750, -0.42 DPS) [quest]; Sludge-Stained Band (286535, -0.45 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.67 DPS) | yes | Snake Hoop (6750, -0.31 DPS) [quest]; Sludge-Stained Band (286535, -0.34 DPS) [world]; Black Widow Band (6199, -0.62 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.01 DPS) | yes | Glimmering Staff (249392, -0.43 DPS) [crafted]; Twisted Chanter's Staff (890, -0.49 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.49 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.8 spell_power points (1.10 DPS) | yes | Orb of Souls (249395, -0.65 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.67 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.20 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 298.3 spell_power points (33.55 DPS) | yes | Starfaller (13063, -0.48 DPS) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 93.3. Weights run: 1.0s. Verify run: 0.7s. 311 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.835 ± 0.014, crit=0.041 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), hit=0.289 ± 0.004 per rating point (10 rating = 1%, 2.886 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.86 DPS) | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Corpseshroud (10574, -0.70 DPS) [dungeon]; Enchanter's Cowl (4322, -0.90 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.0 spell_power points (1.63 DPS) | yes | Triune Amulet (7722, -0.84 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.84 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -0.90 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.9 spell_power points (2.43 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.34 DPS) [quest]; Green Silken Shoulders (7057, -0.90 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.5 spell_power points (2.25 DPS) | yes | Guardian Cloak (5965, -0.86 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.00 DPS) [vendor]; Long Silken Cloak (4326, -2.52 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.0 spell_power points (3.67 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.41 DPS) [crafted]; Elemental Raiment (9434, -0.82 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.7 spell_power points (1.45 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.23 DPS) [dungeon]; Windchaser Cuffs (14429, -0.43 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (2.90 DPS) | yes | Red Mageweave Gloves (10018, -0.27 DPS) [crafted]; Black Mageweave Gloves (10003, -0.86 DPS) [crafted]; Stormcloth Gloves (10011, -1.00 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.5 spell_power points (2.66 DPS) | yes | Gilded Cord (254037, -0.66 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.82 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.0 spell_power points (3.27 DPS) | yes | Crimson Silk Pantaloons (7062, -0.99 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.41 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.26 DPS) | yes | Gilded Slippers (254001, -1.52 DPS) [crafted]; Acidic Walkers (9454, -1.67 DPS) [dungeon]; Spidersilk Boots (4320, -1.86 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.0 spell_power points (2.04 DPS) | yes | Reedknot Ring (9622, -1.09 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.23 DPS) [vendor]; Black Widow Band (6199, -1.25 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.22 DPS) | yes | Sea Giant's Toe Ring (274746, -0.41 DPS) [vendor]; Black Widow Band (6199, -0.43 DPS) [world]; Reedknot Ring (9622, -1.69 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (93.3 DPS) | yes | Windweaver Staff (7757, -1.02 DPS) [dungeon]; Staff of Jordan (873, -1.47 DPS) [world_drop]; Gut Ripper (2164, -4.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 294.6 spell_power points (40.06 DPS) | yes | Umbral Wand (5216, -1.44 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.47 DPS) [dungeon]; Twisted Nether Wand (249144, -5.25 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 137.5. Weights run: 0.9s. Verify run: 0.8s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.484 ± 0.024, crit=0.052 ± 0.002 per rating point (14 rating = 1%, 0.724 per %), hit=0.436 ± 0.006 per rating point (10 rating = 1%, 4.357 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.7 spell_power points (6.17 DPS) | yes | Soulcatcher Halo (10630, -1.47 DPS) [dungeon]; Dreamweave Circlet (10041, -1.63 DPS) [crafted]; Chief Architect's Monocle (11839, -4.71 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.8 spell_power points (2.63 DPS) | yes | Scorn's Icy Choker (23169, -0.62 DPS) [dungeon]; Mindburst Medallion (11196, -0.74 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.75 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.7 spell_power points (5.03 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.33 DPS) [crafted]; Inquisitor's Shawl (19507, -1.70 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 25.4 spell_power points (3.21 DPS) | yes | Spritecaster Cape (11623, -0.31 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.38 DPS) [dungeon]; Runecloth Cloak (13860, -0.57 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.7 spell_power points (6.17 DPS) | yes | Runecloth Robe (13858, -1.58 DPS) [crafted]; Runecloth Tunic (13857, -1.95 DPS) [crafted]; Robes of Insight (940, -4.46 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 22.3 spell_power points (2.82 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.50 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.7 spell_power points (5.03 DPS) | yes | Virtuous Hands (226958, -0.17 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.69 DPS) [vendor]; Red Mageweave Gloves (10018, -1.75 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 34.2 spell_power points (4.34 DPS) | yes | Satyrmane Sash (17755, -0.68 DPS) [dungeon]; Ban'thok Sash (11662, -0.69 DPS) [dungeon]; Deathmage Sash (10771, -0.69 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.8 spell_power points (4.80 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.93 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.03 DPS) [dungeon]; Red Mageweave Pants (10009, -3.05 DPS, sim-verified) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (137.5 DPS) | yes | Southsea Mojo Boots (20641, -0.00 DPS) [quest]; Earthen Silk Slippers (254013, -0.05 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -1.48 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 22.3 spell_power points (2.82 DPS) | yes | Philanthropist's Ring (281635, -0.43 DPS) [quest]; Mindseye Circle (10634, -0.56 DPS) [dungeon]; Woodseed Hoop (17768, -1.13 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.4 spell_power points (2.46 DPS) | yes | Mindseye Circle (10634, -0.20 DPS) [dungeon]; Woodseed Hoop (17768, -0.76 DPS) [quest]; Philanthropist's Ring (281635, -0.90 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.77 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -1.13 DPS) [quest]; Radiant Staff (249453, -1.88 DPS) [crafted]; Blade of Eternal Darkness (17780, -5.18 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 414.2 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 300.1. Weights run: 1.1s. Verify run: 0.7s. 1120 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.174 ± 0.031, crit=0.116 ± 0.004 per rating point (14 rating = 1%, 1.620 per %), hit=0.753 ± 0.009 per rating point (10 rating = 1%, 7.527 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 48.8 spell_power points (7.34 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -4.07 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.3 spell_power points (4.56 DPS) | yes | Archlight Talisman (15856, -0.98 DPS) [quest]; Lady Maye's Pendant (14558, -1.20 DPS) [world_drop]; Beads of Ogre Mojo (22149, -1.43 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.2 spell_power points (6.96 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.38 DPS) [vendor]; Virtuous Epaulets (226955, -0.85 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.9 spell_power points (4.96 DPS) | yes | Hide of the Wild (18510, -1.08 DPS) [crafted]; Crystalline Threaded Cape (20697, -1.24 DPS) [world]; Deep Woodlands Cloak (19121, -1.56 DPS) [quest] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 47.1 spell_power points (7.10 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -3.84 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.4 spell_power points (4.73 DPS) | yes | Sublime Wristguards (18497, -1.15 DPS) [dungeon]; Runecloth Cuffs (254123, -1.30 DPS) [crafted]; General's Satin Bracers (17619, -1.37 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 39.7 spell_power points (5.98 DPS) | yes | General's Satin Gloves (17620, -0.22 DPS) [vendor]; General's Satin Grips (231613, -0.22 DPS) [vendor]; Virtuous Hands (226958, -2.15 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.4 spell_power points (8.79 DPS) | yes | Magister's Belt (16685, -4.02 DPS) [dungeon]; Dustfeather Sash (12589, -4.25 DPS) [dungeon]; Belt of the Archmage (18405, -6.82 DPS, sim-verified) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 50.3 spell_power points (7.58 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, -0.85 DPS) [vendor]; Sentinel's Silk Leggings (237815, -5.23 DPS, sim-verified) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 40.8 spell_power points (6.14 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -3.43 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (300.1 DPS) | yes | Eye of Orgrimmar (12545, -1.31 DPS) [quest]; Maiden's Circle (13001, -1.31 DPS) [world_drop]; Naglering (11669, -10.43 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (300.1 DPS) | yes | Eye of Orgrimmar (12545, -0.45 DPS) [quest]; Maiden's Circle (13001, -0.45 DPS) [world_drop]; Naglering (11669, -10.91 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (300.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (300.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -3.74 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (300.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -0.18 DPS) [dungeon]; Hand of Edward the Odd (2243, -7.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 515.8 spell_power points (77.66 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -1.39 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.67 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.99 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Outrider's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Second Wind; trinket2: Burst of Knowledge; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 591.9. Weights run: 1.2s. Verify run: 0.9s. 1120 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.737 ± 0.018, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.610 per %), hit=0.520 ± 0.009 per rating point (10 rating = 1%, 5.204 per %), spell_haste=-2.160 ± 0.208, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 39.9 spell_power points (14.13 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 24.6 spell_power points (8.71 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Orb of the Darkmoon (19426, -0.91 DPS) [quest]; Beads of Ogre Mojo (22149, -0.97 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 38.7 spell_power points (13.70 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Virtuous Epaulets (226955, -1.99 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.40 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.1 spell_power points (9.60 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Hide of the Wild (18510, -2.03 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.00 DPS) [quest] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 39.3 spell_power points (13.91 DPS) | yes | Robe of Everlasting Night (18385, +0.00 DPS, sim-verified) [dungeon]; Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.9 spell_power points (9.88 DPS) | yes | Sublime Wristguards (18497, -3.02 DPS) [dungeon]; Runecloth Cuffs (254123, -3.37 DPS) [crafted]; Virtuous Wraps (226953, -4.18 DPS) [vendor] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 30.4 spell_power points (10.78 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Virtuous Hands (226958, -5.45 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 46.4 spell_power points (16.44 DPS) | yes | Belt of the Archmage (18405, -4.25 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -7.46 DPS) [rep]; Magister's Belt (16685, -8.48 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 42.0 spell_power points (14.88 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (237815, +0.00 DPS, sim-verified) [vendor]; Skyshroud Leggings (13170, -0.75 DPS) [dungeon] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 33.8 spell_power points (11.97 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -1.42 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.46 DPS) [quest]; Maiden's Circle (13001, -2.46 DPS) [world_drop]; Naglering (11669, -8.97 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.06 DPS) [quest]; Maiden's Circle (13001, -1.06 DPS) [world_drop]; Naglering (11669, -8.22 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+18.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.48 DPS) [vendor]; Serenity Field (272439, -4.10 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -6.02 DPS) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.97 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (591.9 DPS) | yes | Bonecreeper Stylus (13938, -0.18 DPS) [dungeon]; Sparkling Crystal Wand (20672, -1.30 DPS) [world]; Torch of Light (279246, -8.92 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Outrider's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

