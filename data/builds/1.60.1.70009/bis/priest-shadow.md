# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 22.3. Weights run: 0.9s. Verify run: 1.0s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-2.593 ± 0.062, crit=0.080 ± 0.004 per rating point (14 rating = 1%, 1.117 per %), hit=0.183 ± 0.004 per rating point (10 rating = 1%, 1.832 per %), spell_haste=not significant (-0.514 ± 0.244), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Resilient Cloth Headband (211500, -1.15 DPS, sim-verified) [vendor] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.44 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.09 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.35 DPS) | yes | Black Whelp Cloak (7283, -0.07 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Caretaker's Cape (20428, -0.09 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.44 DPS) | yes | Green Woolen Vest (2582, -0.09 DPS) [crafted]; Bloody Apron (6226, -0.09 DPS) [dungeon]; Gray Woolen Robe (2585, -0.50 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.09 DPS) | yes | Ivycloth Bracelets (9793, -0.10 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.62 DPS) | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.27 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.44 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (22.3 DPS) | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.33 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (0.80 DPS) | yes | Silk-threaded Trousers (1929, +0.00 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Rumpled Kilt (274741, -0.35 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (0.62 DPS) | yes | Feather Padded Treads (285345, -0.07 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.27 DPS) [crafted]; Pristine Boots (253889, -0.35 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.44 DPS) | yes | Sludge-Stained Band (286535, -0.18 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.44 DPS) | yes | Sludge-Stained Band (286535, -0.09 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 101 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -0.61 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.44 DPS) | yes | Bouquet of Red Roses (22206, -0.46 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 255.5 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.30 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 51.6. Weights run: 0.7s. Verify run: 0.7s. 243 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.609 ± 0.021, crit=0.082 ± 0.005 per rating point (14 rating = 1%, 1.149 per %), hit=0.335 ± 0.004 per rating point (10 rating = 1%, 3.350 per %), spell_haste=not significant (0.626 ± 0.345), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.1 spell_power points (1.03 DPS) | yes | Silk Headband (7050, -0.26 DPS) [crafted]; Holy Shroud (2721, -0.34 DPS, sim-verified) [world_drop]; Embalmed Shroud (7691, -0.35 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.7 spell_power points (0.91 DPS) | yes | Crystal Starfire Medallion (5003, -0.70 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.70 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.11 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.5 spell_power points (1.24 DPS) | yes | Fairywing Mantle (9536, -0.26 DPS) [quest]; Magician's Mantle (12998, -0.34 DPS) [world_drop]; Death Speaker Mantle (6685, -0.65 DPS, sim-verified) [dungeon] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.4 spell_power points (0.46 DPS) | yes | Cloak of Rot (4462, -0.05 DPS) [world]; Darkspear Raider's Cloak (272078, -0.05 DPS) [vendor]; Hillman's Cloak (3719, -0.77 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.9 spell_power points (1.44 DPS) | yes | Tree Bark Jacket (1486, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.48 DPS) [crafted]; Death Speaker Robes (6682, -0.74 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.77 DPS) | yes | Nightsky Wristbands (6407, -0.46 DPS) [world_drop]; Stonecloth Bindings (14416, -0.51 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.14 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.7 spell_power points (0.91 DPS) | yes | Serpent Gloves (5970, -0.32 DPS) [dungeon]; Truefaith Gloves (7049, -0.33 DPS) [crafted]; Shilly Mitts (9609, -0.58 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.8 spell_power points (1.09 DPS) | yes | Belt of Arugal (6392, -0.11 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.22 DPS) [crafted]; Invoker's Cord (215366, -0.24 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.9 spell_power points (1.18 DPS) | yes | Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.44 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.3 spell_power points (0.96 DPS) | yes | Spidersilk Boots (4320, -0.16 DPS) [crafted]; Nimbus Boots (6998, -0.45 DPS) [quest]; Acidic Walkers (9454, -1.30 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.60 DPS) | yes | Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor]; Black Widow Band (6199, -0.23 DPS) [world]; Snake Hoop (6750, -0.23 DPS) [quest] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.53 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Black Widow Band (6199, -0.17 DPS) [world]; Snake Hoop (6750, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.77 DPS) | yes | Twisted Chanter's Staff (890, -0.25 DPS) [world_drop]; Channeler's Staff (4437, -0.35 DPS) [world]; Glimmering Staff (249392, -0.70 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.7 spell_power points (0.91 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.53 DPS) [vendor]; Eye of Paleth (2943, -0.57 DPS) [quest]; Dwarven Tome (279898, -1.17 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 392.4 spell_power points (33.47 DPS) | yes | Starfaller (13063, -0.68 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.04 DPS) [crafted]; Gravestone Scepter (7001, -4.47 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 243, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 82.9. Weights run: 0.8s. Verify run: 0.7s. 327 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.984 ± 0.031, crit=0.098 ± 0.007 per rating point (14 rating = 1%, 1.372 per %), hit=0.456 ± 0.008 per rating point (10 rating = 1%, 4.562 per %), spell_haste=not significant (-1.521 ± 0.482), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (1.94 DPS) | yes | Augural Shroud (2620, -0.15 DPS, sim-verified) [world]; Corpseshroud (10574, -0.21 DPS) [dungeon]; Thinking Cap (2624, -0.39 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.9 spell_power points (1.19 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.43 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.56 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.56 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.8 spell_power points (1.83 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 17.9 spell_power points (1.65 DPS) | yes | Guardian Cloak (5965, -0.64 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.65 DPS) [vendor]; Long Silken Cloak (4326, -0.69 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.9 spell_power points (2.58 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.19 DPS) [crafted]; Crimson Silk Vest (7058, -0.56 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (0.83 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.01 DPS) [world_drop]; Mistscape Bracers (4045, -0.10 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.9 spell_power points (2.03 DPS) | yes | Red Mageweave Gloves (10018, +0.00 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.57 DPS) [crafted]; Black Mageweave Gloves (10003, -0.64 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.8 spell_power points (2.01 DPS) | yes | Gilded Cord (254037, -0.54 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.72 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.8 spell_power points (2.38 DPS) | yes | Crimson Silk Pantaloons (7062, -0.69 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.83 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.01 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.22 DPS) | yes | Gilded Slippers (254001, -0.44 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.03 DPS) [dungeon]; Spidersilk Boots (4320, -1.21 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.9 spell_power points (1.47 DPS) | yes | Ring of Forlorn Spirits (2043, -0.73 DPS) [quest]; Reedknot Ring (9622, -0.82 DPS) [quest]; Minor Channeling Ring (1449, -0.83 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.83 DPS) | yes | Reedknot Ring (9622, -0.18 DPS) [quest]; Minor Channeling Ring (1449, -0.19 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.82 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (82.9 DPS) | yes | Windweaver Staff (7757, -0.48 DPS) [dungeon]; Staff of Jordan (873, -0.85 DPS) [world_drop]; Gut Ripper (2164, -3.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 432.1 spell_power points (39.93 DPS) | yes | Umbral Wand (5216, -0.57 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.34 DPS) [dungeon]; Twisted Nether Wand (249144, -5.37 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 327, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 136.3. Weights run: 0.7s. Verify run: 0.8s. 422 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.099 ± 0.047, crit=0.115 ± 0.009 per rating point (14 rating = 1%, 1.607 per %), hit=0.638 ± 0.012 per rating point (10 rating = 1%, 6.384 per %), spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.0 spell_power points (4.27 DPS) | yes | Chief Architect's Monocle (11839, -1.18 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Hat (220889, -1.27 DPS) [vendor]; Dreamweave Circlet (10041, -1.56 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 15.4 spell_power points (1.60 DPS) | yes | Mindburst Medallion (11196, -0.29 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.46 DPS) [quest]; Scorn's Icy Choker (23169, -0.82 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.8 spell_power points (3.42 DPS) | yes | Kentic Amice (11624, -0.27 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -0.97 DPS) [crafted]; Inquisitor's Shawl (19507, -1.20 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.6 spell_power points (2.15 DPS) | yes | Mantle of Lady Falther'ess (23178, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.29 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.54 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.0 spell_power points (4.27 DPS) | yes | Runecloth Tunic (13857, -1.24 DPS) [crafted]; Robe of the Magi (1716, -1.29 DPS) [world_drop]; Runecloth Robe (13858, -1.89 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 16.5 spell_power points (1.72 DPS) | yes | Nethergeld Cuffs (254061, -0.19 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.34 DPS) [quest]; Bloodband Bracers (11469, -1.10 DPS, sim-verified) [quest] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 33.3 spell_power points (3.47 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.08 DPS) [vendor]; Dreamweave Gloves (10019, -1.14 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.9 spell_power points (2.80 DPS) | yes | Ban'thok Sash (11662, -0.22 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon]; Satyrmane Sash (17755, -0.88 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.0 spell_power points (3.54 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.75 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.17 DPS) [dungeon]; Red Mageweave Pants (10009, -1.87 DPS, sim-verified) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 24.3 spell_power points (2.53 DPS) | yes | Gilded Sandals (254107, -0.35 DPS) [crafted]; Southsea Mojo Boots (20641, -0.44 DPS) [quest]; Earthen Silk Slippers (254013, -0.83 DPS, sim-verified) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.7 spell_power points (1.74 DPS) | yes | Brainlash (6440, -0.02 DPS) [dungeon]; Mindseye Circle (10634, -0.37 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.6 spell_power points (1.73 DPS) | yes | Brainlash (6440, +0.00 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -0.35 DPS) [dungeon]; Band of the Unicorn (7553, -0.37 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (136.3 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (136.3 DPS) | yes | Blessed Prayer Beads (19990, -0.91 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (136.3 DPS) | yes | Spellshifter Rod (9527, -0.69 DPS) [quest]; Radiant Staff (249453, -1.15 DPS) [crafted]; Blade of Eternal Darkness (17780, -4.73 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 spell_power points (52.50 DPS) | yes | Woestave (20082, +0.00 DPS, sim-verified) [quest]; Noxious Shooter (17745, -1.98 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Virtuous Hands; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 255.1. Weights run: 0.8s. Verify run: 0.8s. 1096 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.190 ± 0.057, crit=0.224 ± 0.013 per rating point (14 rating = 1%, 3.142 per %), hit=1.121 ± 0.017 per rating point (10 rating = 1%, 11.211 per %), spell_haste=not significant (0.982 ± 1.111), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Field Marshal's Satin Crown (231616) | Captain Dirgehammer [vendor] | 66.3 spell_power points (7.43 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Virtuous Cowl (226957, -1.76 DPS) [vendor]; Field Marshal's Satin Hood (231622, -2.41 DPS, sim-verified) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.5 spell_power points (3.41 DPS) | yes | Archlight Talisman (15856, -0.74 DPS) [quest]; Beads of Ogre Mojo (22149, -0.87 DPS, sim-verified) [quest]; Lady Maye's Pendant (14558, -0.88 DPS) [world_drop] |
| shoulder | Field Marshal's Satin Mantle (17604) (or Field Marshal's Satin Epaulets (231621)) | Captain Dirgehammer [vendor] | 48.8 spell_power points (5.47 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -0.09 DPS) [vendor]; Virtuous Epaulets (226955, -0.48 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.7 spell_power points (4.11 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -1.34 DPS) [world]; Spritecaster Cape (11623, -1.75 DPS) [dungeon] |
| chest | Field Marshal's Satin Robe (231618) | Captain Dirgehammer [vendor] | 66.3 spell_power points (7.43 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Virtuous Gown (226960, -2.12 DPS) [vendor]; Field Marshal's Satin Tunic (231624, -2.41 DPS, sim-verified) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.5 spell_power points (3.53 DPS) | yes | Sublime Wristguards (18497, -0.85 DPS) [dungeon]; Runecloth Cuffs (254123, -0.97 DPS) [crafted]; Marshal's Satin Bracers (17606, -1.00 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 40.1 spell_power points (4.49 DPS) | yes | Marshal's Satin Gloves (17608, -0.19 DPS) [vendor]; Virtuous Hands (226958, -0.63 DPS) [vendor]; Marshal's Satin Grips (231617, -1.59 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 62.4 spell_power points (6.99 DPS) | yes | Magister's Belt (16685, -3.41 DPS) [dungeon]; Dustfeather Sash (12589, -3.58 DPS) [dungeon]; Belt of the Archmage (18405, -3.99 DPS, sim-verified) [crafted] |
| legs | Marshal's Satin Leggings (231619) | Captain Dirgehammer [vendor] | 57.0 spell_power points (6.38 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -0.76 DPS) [vendor] |
| feet | Marshal's Satin Treads (231620) | Captain Dirgehammer [vendor] | 44.4 spell_power points (4.98 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Virtuous Slippers (226959, +0.00 DPS, sim-verified) [vendor]; Dragonrider Boots (18102, -0.83 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (255.1 DPS) | yes | Channeler's Ring (272406, -2.47 DPS) [vendor]; Maiden's Circle (13001, -2.64 DPS) [world_drop]; Naglering (11669, -9.17 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (255.1 DPS) | yes | Channeler's Ring (272406, -0.81 DPS) [vendor]; Maiden's Circle (13001, -0.98 DPS) [world_drop]; Naglering (11669, -5.87 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (255.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -0.42 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (255.1 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -1.18 DPS, sim-verified) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (255.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.12 DPS) [world]; Hand of Edward the Odd (2243, -6.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 691.3 spell_power points (77.43 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -3.82 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -13.04 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.34 DPS) [world] |

**New at 60:** head: Field Marshal's Satin Crown; neck: Amulet of the Dawn; shoulder: Field Marshal's Satin Mantle; back: Arcanoweave Cloak; chest: Field Marshal's Satin Robe; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Marshal's Satin Leggings; feet: Marshal's Satin Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1096, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 27.4. Weights run: 0.9s. Verify run: 1.0s. 142 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-2.593 ± 0.062, crit=0.080 ± 0.004 per rating point (14 rating = 1%, 1.117 per %), hit=0.183 ± 0.004 per rating point (10 rating = 1%, 1.832 per %), spell_haste=not significant (-0.514 ± 0.244), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Resilient Cloth Headband (211500, -1.13 DPS, sim-verified) [vendor] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (0.44 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.09 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.56 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.35 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.14 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.44 DPS) | yes | Green Woolen Vest (2582, -0.09 DPS) [crafted]; Bloody Apron (6226, -0.09 DPS) [dungeon]; Gray Woolen Robe (2585, -0.48 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.09 DPS) | yes | Owlbeard Bracers (16981, -0.58 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.62 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.27 DPS) [quest]; Pristine Gloves (253913, -0.27 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (27.4 DPS) | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.51 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (0.80 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Rumpled Kilt (274741, -0.35 DPS) [vendor]; Silk-threaded Trousers (1929, -0.42 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (0.62 DPS) | yes | Red Woolen Boots (4313, -0.27 DPS) [crafted]; Pristine Boots (253889, -0.35 DPS) [crafted]; Feather Padded Treads (285345, -0.70 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.44 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.27 DPS) | yes | Ring of the Shadow (1462, -0.26 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Clear Crystal Rod (16894), and 111 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -0.50 DPS, sim-verified) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), Pulsating Hydra Heart (5183), and 7 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, -0.48 DPS, sim-verified) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 255.5 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.77 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 142, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 209613 Insignia of the Alliance

### Band 30 (undead, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 45.8. Weights run: 0.7s. Verify run: 0.7s. 235 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.609 ± 0.021, crit=0.082 ± 0.005 per rating point (14 rating = 1%, 1.149 per %), hit=0.335 ± 0.004 per rating point (10 rating = 1%, 3.350 per %), spell_haste=not significant (0.626 ± 0.345), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.1 spell_power points (1.03 DPS) | yes | Holy Shroud (2721, -0.05 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.26 DPS) [crafted]; Embalmed Shroud (7691, -0.35 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.7 spell_power points (0.91 DPS) | yes | Crystal Starfire Medallion (5003, -0.70 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.70 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.74 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.5 spell_power points (1.24 DPS) | yes | Death Speaker Mantle (6685, -0.08 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.26 DPS) [quest]; Magician's Mantle (12998, -0.34 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.43 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Cloak of Rot (4462, -0.01 DPS) [world]; Darkspear Raider's Cloak (272078, -0.01 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.9 spell_power points (1.44 DPS) | yes | Death Speaker Robes (6682, -0.07 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.48 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.77 DPS) | yes | Nightsky Wristbands (6407, -0.46 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.46 DPS) [quest]; Glowing Magical Bracelets (13106, -0.56 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.0 spell_power points (0.77 DPS) | yes | Serpent Gloves (5970, -0.16 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.19 DPS) [crafted]; Gnoll Casting Gloves (892, -0.26 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.8 spell_power points (1.09 DPS) | yes | Warsong Sash (16975, +0.00 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.17 DPS) [dungeon]; Crimson Silk Belt (7055, -0.22 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.9 spell_power points (1.18 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.3 spell_power points (0.96 DPS) | yes | Spidersilk Boots (4320, -0.16 DPS) [crafted]; Boots of the Enchanter (4325, -0.53 DPS) [crafted]; Acidic Walkers (9454, -0.77 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.60 DPS) | yes | Black Widow Band (6199, -0.23 DPS) [world]; Snake Hoop (6750, -0.23 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.51 DPS) | yes | Snake Hoop (6750, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Black Widow Band (6199, -0.37 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.77 DPS) | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.25 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.25 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.7 spell_power points (0.91 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.36 DPS, sim-verified) [vendor]; Witch's Finger (16887, -0.55 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.57 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 392.4 spell_power points (33.47 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.04 DPS) [crafted]; Gravestone Scepter (7001, -4.47 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 235, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 74.4. Weights run: 0.8s. Verify run: 0.7s. 319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.984 ± 0.031, crit=0.098 ± 0.007 per rating point (14 rating = 1%, 1.372 per %), hit=0.456 ± 0.008 per rating point (10 rating = 1%, 4.562 per %), spell_haste=not significant (-1.521 ± 0.482), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (74.4 DPS) | yes | Corpseshroud (10574, -0.20 DPS) [dungeon]; Thinking Cap (2624, -0.38 DPS) [world]; Spellpower Goggles Xtreme (10502, -0.82 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.9 spell_power points (1.19 DPS) | yes | Prodigious Shadowshard Pendant (17773, +0.00 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.56 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.56 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.8 spell_power points (1.83 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 17.9 spell_power points (1.65 DPS) | yes | Guardian Cloak (5965, -0.64 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.65 DPS) [vendor]; Long Silken Cloak (4326, -1.12 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.9 spell_power points (2.58 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.19 DPS) [crafted]; Crimson Silk Vest (7058, -0.56 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 11.9 spell_power points (1.10 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.28 DPS) [world_drop]; Mistscape Bracers (4045, -0.37 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.9 spell_power points (2.03 DPS) | yes | Red Mageweave Gloves (10018, -0.17 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.57 DPS) [crafted]; Black Mageweave Gloves (10003, -0.64 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.8 spell_power points (2.01 DPS) | yes | Defiler's Cloth Girdle (20166, -0.37 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.8 spell_power points (2.38 DPS) | yes | Crimson Silk Pantaloons (7062, +0.00 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.83 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.01 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.22 DPS) | yes | Gilded Slippers (254001, -0.17 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.03 DPS) [dungeon]; Spidersilk Boots (4320, -1.21 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.9 spell_power points (1.47 DPS) | yes | Reedknot Ring (9622, -0.82 DPS) [quest]; Voodoo Band (1996, -0.83 DPS) [world]; Black Widow Band (6199, -0.83 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.83 DPS) | yes | Voodoo Band (1996, -0.20 DPS) [world]; Black Widow Band (6199, -0.20 DPS) [world]; Reedknot Ring (9622, -0.74 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -0.48 DPS) [dungeon]; Staff of Jordan (873, -0.85 DPS) [world_drop]; Gut Ripper (2164, -2.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 432.1 spell_power points (39.93 DPS) | yes | Umbral Wand (5216, -1.01 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.34 DPS) [dungeon]; Twisted Nether Wand (249144, -5.37 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 123.1. Weights run: 0.7s. Verify run: 0.8s. 414 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.099 ± 0.047, crit=0.115 ± 0.009 per rating point (14 rating = 1%, 1.607 per %), hit=0.638 ± 0.012 per rating point (10 rating = 1%, 6.384 per %), spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.0 spell_power points (4.27 DPS) | yes | Dreamweave Circlet (10041, -0.70 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.18 DPS) [dungeon]; Blood Guard's Dreadweave Hat (220907, -1.27 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 15.4 spell_power points (1.60 DPS) | yes | Mindburst Medallion (11196, -0.29 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.46 DPS) [quest]; Scorn's Icy Choker (23169, -0.79 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.8 spell_power points (3.42 DPS) | yes | Kentic Amice (11624, -0.14 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -0.97 DPS) [crafted]; Inquisitor's Shawl (19507, -1.20 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.9 spell_power points (2.28 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.31 DPS) [dungeon]; Runecloth Cloak (13860, -0.43 DPS) [crafted]; Spritecaster Cape (11623, -0.51 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.0 spell_power points (4.27 DPS) | yes | Runecloth Tunic (13857, -1.24 DPS) [crafted]; Robe of the Magi (1716, -1.29 DPS) [world_drop]; Runecloth Robe (13858, -2.08 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 16.5 spell_power points (1.72 DPS) | yes | Nethergeld Cuffs (254061, -0.19 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.34 DPS) [quest]; Bloodband Bracers (11469, -0.57 DPS, sim-verified) [quest] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 33.3 spell_power points (3.47 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.08 DPS) [vendor]; Dreamweave Gloves (10019, -1.14 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.9 spell_power points (2.80 DPS) | yes | Ban'thok Sash (11662, -0.22 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon]; Satyrmane Sash (17755, -0.65 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.0 spell_power points (3.54 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.75 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.17 DPS) [dungeon]; Red Mageweave Pants (10009, -2.10 DPS, sim-verified) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 24.3 spell_power points (2.53 DPS) | yes | Earthen Silk Slippers (254013, -0.34 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -0.35 DPS) [crafted]; Southsea Mojo Boots (20641, -0.44 DPS) [quest] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.7 spell_power points (1.74 DPS) | yes | Brainlash (6440, -0.02 DPS) [dungeon]; Mindseye Circle (10634, -0.37 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.6 spell_power points (1.73 DPS) | yes | Brainlash (6440, +0.00 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -0.35 DPS) [dungeon]; Band of the Unicorn (7553, -0.37 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (123.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (123.1 DPS) | yes | Uther's Strength (11302, -0.19 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (123.1 DPS) | yes | Spellshifter Rod (9527, -0.69 DPS) [quest]; Radiant Staff (249453, -1.15 DPS) [crafted]; Blade of Eternal Darkness (17780, -5.98 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 spell_power points (52.50 DPS) | yes | Woestave (20082, +0.00 DPS, sim-verified) [quest]; Noxious Shooter (17745, -1.98 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Virtuous Hands; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 241.2. Weights run: 0.8s. Verify run: 0.8s. 1087 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.190 ± 0.057, crit=0.224 ± 0.013 per rating point (14 rating = 1%, 3.142 per %), hit=1.121 ± 0.017 per rating point (10 rating = 1%, 11.211 per %), spell_haste=not significant (0.982 ± 1.111), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warlord's Satin Crown (231615) | Lady Palanseer [vendor] | 66.3 spell_power points (7.43 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Virtuous Cowl (226957, -1.76 DPS) [vendor]; Warlord's Satin Hood (231635, -1.84 DPS, sim-verified) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.5 spell_power points (3.41 DPS) | yes | Beads of Ogre Mojo (22149, -0.16 DPS, sim-verified) [quest]; Archlight Talisman (15856, -0.74 DPS) [quest]; Lady Maye's Pendant (14558, -0.88 DPS) [world_drop] |
| shoulder | Warlord's Satin Mantle (17622) (or Warlord's Satin Epaulets (231611)) | Lady Palanseer [vendor] | 48.8 spell_power points (5.47 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -0.09 DPS) [vendor]; Virtuous Epaulets (226955, -0.48 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.7 spell_power points (4.11 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -1.34 DPS) [world]; Deep Woodlands Cloak (19121, -1.57 DPS) [quest] |
| chest | Warlord's Satin Robes (17624) | Lady Palanseer [vendor] | 66.3 spell_power points (7.43 DPS) | yes | Warlord's Satin Tunic (231632, -1.84 DPS, sim-verified) [vendor]; Virtuous Gown (226960, -2.12 DPS) [vendor]; Magister's Robes (16688, -2.35 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.5 spell_power points (3.53 DPS) | yes | Sublime Wristguards (18497, -0.85 DPS) [dungeon]; Runecloth Cuffs (254123, -0.97 DPS) [crafted]; General's Satin Bracers (17619, -1.00 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 40.1 spell_power points (4.49 DPS) | yes | General's Satin Gloves (17620, -0.19 DPS) [vendor]; Virtuous Hands (226958, -0.63 DPS) [vendor]; General's Satin Grips (231613, -2.14 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 62.4 spell_power points (6.99 DPS) | yes | Magister's Belt (16685, -3.41 DPS) [dungeon]; Dustfeather Sash (12589, -3.58 DPS) [dungeon]; Belt of the Archmage (18405, -4.41 DPS, sim-verified) [crafted] |
| legs | General's Satin Leggings (17625) | Lady Palanseer [vendor] | 57.0 spell_power points (6.38 DPS) | yes | Outrider's Silk Leggings (22747, -0.72 DPS, sim-verified) [rep]; Sentinel's Silk Leggings (237815, -0.76 DPS) [vendor]; General's Satin Legguards (231634, -1.34 DPS) [vendor] |
| feet | General's Satin Treads (231610) | Lady Palanseer [vendor] | 44.4 spell_power points (4.98 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; Virtuous Slippers (226959, +0.00 DPS, sim-verified) [vendor]; Dragonrider Boots (18102, -0.83 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (241.2 DPS) | yes | Channeler's Ring (272406, -2.47 DPS) [vendor]; Maiden's Circle (13001, -2.64 DPS) [world_drop]; Naglering (11669, -8.87 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (241.2 DPS) | yes | Channeler's Ring (272406, -0.81 DPS) [vendor]; Maiden's Circle (13001, -0.98 DPS) [world_drop]; Naglering (11669, -5.74 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (241.2 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (241.2 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -3.43 DPS, sim-verified) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (241.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -0.11 DPS) [dungeon]; Hand of Edward the Odd (2243, -4.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 691.3 spell_power points (77.43 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -3.85 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -13.04 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.34 DPS) [world] |

**New at 60:** head: Warlord's Satin Crown; neck: Amulet of the Dawn; shoulder: Warlord's Satin Mantle; back: Arcanoweave Cloak; chest: Warlord's Satin Robes; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: General's Satin Leggings; feet: General's Satin Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1087, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

