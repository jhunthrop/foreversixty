# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.7. Weights run: 1.3s. Verify run: 1.3s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.491, intellect=0.360 ± 0.067, crit=0.135 ± 0.006 per rating point (14 rating = 1%, 1.886 per %), hit=0.466 ± 0.005 per rating point (10 rating = 1%, 4.660 per %), spell_haste=4.500 ± 0.499, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.491)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.30 DPS) | yes | Shadow Goggles (4373, -2.81 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.2 spell_power points (0.41 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.14 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.20 DPS) | yes | Feyscale Cloak (6632, -0.05 DPS) [dungeon]; Caretaker's Cape (20428, -0.05 DPS) [rep]; Pearl-clasped Cloak (5542, -0.33 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.34 DPS) | yes | Green Woolen Robe (6243, -0.14 DPS) [crafted]; Bloody Apron (6226, -0.14 DPS) [dungeon]; Gray Woolen Robe (2585, -1.25 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.8 spell_power points (0.09 DPS) | yes | Windsong Bangles (263336, -0.04 DPS) [quest]; Repurposed Hair Band (281256, -0.05 DPS) [quest]; Bright Bracers (3647, -0.12 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.35 DPS) | yes | Pristine Gloves (253913, -0.10 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.21 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS, sim-verified) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.27 DPS) | yes | Novice Ardent's Sash (253887, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.13 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.93 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (38.7 DPS) | yes | Silk-threaded Trousers (1929, -0.06 DPS) [dungeon]; Rumpled Kilt (274741, -0.16 DPS) [vendor]; Abomination Skin Leggings (23173, -0.69 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.42 DPS) | yes | Pristine Boots (253889, -0.22 DPS) [crafted]; Red Woolen Boots (4313, -0.22 DPS) [crafted]; Feather Padded Treads (285345, -0.58 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.28 DPS) | yes | Sludge-Stained Band (286535, -0.14 DPS) [world]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.25 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Sludge-Stained Band (286535, -0.58 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.6 spell_power points (0.18 DPS) | yes | Lesser Staff of the Spire (1300, -0.07 DPS) [world_drop]; Staff of Westfall (2042, -0.09 DPS) [quest]; Channeler's Staff (4437, -0.10 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 452.0 spell_power points (22.46 DPS) | yes | Skycaller (12984, -1.47 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.17 DPS) [dungeon]; Deepblaze (279896, -3.93 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 59.4. Weights run: 1.3s. Verify run: 1.2s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.160, intellect=0.439 ± 0.013, crit=0.041 ± 0.002 per rating point (14 rating = 1%, 0.567 per %), hit=0.158 ± 0.002 per rating point (10 rating = 1%, 1.578 per %), spell_haste=not significant (0.155 ± 0.172), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.160

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.86 DPS) | yes | Silk Headband (7050, -0.34 DPS) [crafted]; Embalmed Shroud (7691, -0.51 DPS) [dungeon]; Enchanter's Cowl (4322, -0.86 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (1.63 DPS) | yes | Crystal Starfire Medallion (5003, -1.33 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.33 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.85 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (2.19 DPS) | yes | Fairywing Mantle (9536, -0.51 DPS) [quest]; Death Speaker Mantle (6685, -0.62 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.64 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.85 DPS) | yes | Repairman's Cape (9605, -0.10 DPS, sim-verified) [quest]; Prelacy Cape (7004, -0.17 DPS) [quest]; Caretaker's Cape (19533, -0.17 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.7 spell_power points (2.49 DPS) | yes | Tree Bark Jacket (1486, -0.43 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.49 DPS) [dungeon]; Pristine Gown (253961, -0.78 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.52 DPS) | yes | Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Stonecloth Bindings (14416, -1.15 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.16 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.8 spell_power points (1.49 DPS) | yes | Shilly Mitts (9609, -0.17 DPS, sim-verified) [quest]; Serpent Gloves (5970, -0.31 DPS) [dungeon]; Truefaith Gloves (7049, -0.43 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.3 spell_power points (2.09 DPS) | yes | Invoker's Cord (215366, -0.53 DPS) [crafted]; Crimson Silk Belt (7055, -0.55 DPS) [crafted]; Belt of Arugal (6392, -0.87 DPS, sim-verified) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.5 spell_power points (2.12 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.41 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.1 spell_power points (1.71 DPS) | yes | Acidic Walkers (9454, -0.26 DPS) [dungeon]; Nimbus Boots (6998, -0.69 DPS) [quest]; Spidersilk Boots (4320, -2.41 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.19 DPS) | yes | Minor Channeling Ring (1449, -0.19 DPS) [quest]; Lorekeeper's Ring (20431, -0.34 DPS) [rep]; Snake Hoop (6750, -0.66 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.02 DPS) | yes | Black Widow Band (6199, -0.50 DPS) [world]; Snake Hoop (6750, -0.50 DPS) [quest]; Minor Channeling Ring (1449, -1.63 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.52 DPS) | yes | Twisted Chanter's Staff (890, -0.78 DPS) [world_drop]; Channeler's Staff (4437, -0.93 DPS) [world]; Glimmering Staff (249392, -1.58 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.6 spell_power points (1.63 DPS) | yes | Eye of Paleth (2943, -0.95 DPS) [quest]; Orb of Souls (249395, -0.95 DPS) [crafted]; Dwarven Tome (279898, -1.10 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 199.2 spell_power points (33.72 DPS) | yes | Starfaller (13063, -0.62 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.87 DPS) [crafted]; Gravestone Scepter (7001, -4.72 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 115.0. Weights run: 1.1s. Verify run: 1.0s. 329 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.422, intellect=0.796 ± 0.026, crit=0.086 ± 0.005 per rating point (14 rating = 1%, 1.207 per %), hit=0.359 ± 0.004 per rating point (10 rating = 1%, 3.589 per %), spell_haste=not significant (1.639 ± 0.466), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.422)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (1.88 DPS) | yes | Corpseshroud (10574, -0.52 DPS) [dungeon]; Enchanter's Cowl (4322, -0.63 DPS) [crafted]; Augural Shroud (2620, -2.89 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.8 spell_power points (1.05 DPS) | yes | Necklace of Calisea (1714, -0.55 DPS) [world_drop]; Triune Amulet (7722, -0.55 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.21 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.4 spell_power points (1.55 DPS) | yes | Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Green Silken Shoulders (7057, -0.19 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.21 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.2 spell_power points (1.45 DPS) | yes | Guardian Cloak (5965, -0.55 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.66 DPS) [vendor]; Long Silken Cloak (4326, -2.19 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.8 spell_power points (2.39 DPS) | yes | Robe of Power (7054, -0.29 DPS) [crafted]; Elemental Raiment (9434, -0.52 DPS) [world_drop]; Dreamweave Vest (10021, -1.13 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (0.80 DPS) | yes | Windchaser Cuffs (14429, -0.16 DPS) [world_drop]; Condor Bracers (15864, -0.18 DPS) [quest]; Spidertank Oilrag (9448, -0.25 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (1.89 DPS) | yes | Black Mageweave Gloves (10003, -0.55 DPS) [crafted]; Gilded Handwraps (254021, -0.68 DPS) [crafted]; Red Mageweave Gloves (10018, -2.18 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.9 spell_power points (1.69 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.41 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.50 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.6 spell_power points (2.11 DPS) | yes | Abomination Skin Leggings (23173, -0.73 DPS) [dungeon]; Stoneweaver Leggings (9407, -0.91 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.15 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.15 DPS) | yes | Acidic Walkers (9454, -1.13 DPS) [dungeon]; Spidersilk Boots (4320, -1.24 DPS) [crafted]; Gilded Slippers (254001, -3.35 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (1.32 DPS) | yes | Ring of Forlorn Spirits (2043, -0.61 DPS) [quest]; Reedknot Ring (9622, -0.70 DPS) [quest]; Minor Channeling Ring (1449, -0.73 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.80 DPS) | yes | Reedknot Ring (9622, -0.18 DPS) [quest]; Lorekeeper's Ring (19525, -0.18 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.77 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | sim-verified (115.0 DPS) | yes | Staff of Dar'Orahil (15106, -0.68 DPS) [quest]; Windweaver Staff (7757, -0.72 DPS) [dungeon]; Gut Ripper (2164, -7.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 445.2 spell_power points (39.81 DPS) | yes | Umbral Wand (5216, -0.19 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.22 DPS) [dungeon]; Twisted Nether Wand (249144, -5.27 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 329, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 208.0. Weights run: 1.3s. Verify run: 1.3s. 419 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.227, intellect=0.211 ± 0.015, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.457 per %), hit=0.203 ± 0.003 per rating point (10 rating = 1%, 2.027 per %), spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.84 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 17.3 spell_power points (4.84 DPS) | yes | Mindburst Medallion (11196, -2.80 DPS) [quest]; Horizon Choker (13085, -4.01 DPS) [world_drop]; Scorn's Icy Choker (23169, -5.00 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.8 spell_power points (4.71 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.37 DPS) [crafted]; Bloodmage Mantle (7684, -1.65 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.28 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.48 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.28 DPS) [crafted]; Nightfall Drape (12465, -1.76 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (6.52 DPS) | yes | Elemental Raiment (9434, -0.64 DPS) [world_drop]; Dreamweave Vest (10021, -0.94 DPS) [crafted]; Acumen Robes (17775, -1.17 DPS, sim-verified) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.15 DPS) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.28 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.11 DPS) [vendor]; Runecloth Gloves (13863, -1.39 DPS) [crafted]; Black Mageweave Gloves (10003, -1.57 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (4.52 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Ghostweave Cord (254073, -0.59 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.1 spell_power points (7.04 DPS) | yes | Red Mageweave Pants (10009, -2.40 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.84 DPS) [vendor]; Wizardweave Leggings (14132, -3.76 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -0.14 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.23 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.39 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.49 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.36 DPS) | yes | Cyclopean Band (11824, -0.43 DPS) [dungeon]; Philanthropist's Ring (281635, -0.64 DPS, sim-verified) [quest]; Lorekeeper's Ring (19524, -0.84 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -3.39 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -1.01 DPS, sim-verified) [crafted] |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Moonshadow Stave (22458, -0.55 DPS) [quest]; Arbiter's Blade (11784, -3.07 DPS) [dungeon]; Blade of Eternal Darkness (17780, -7.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (208.0 DPS) | yes | Woestave (20082, -0.08 DPS) [quest]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Pyric Caduceus (11748, -5.00 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Noxious Shooter

No-known-source sample (15 of 419, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 425.0. Weights run: 1.3s. Verify run: 1.3s. 1007 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.180, intellect=0.192 ± 0.014, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), hit=0.172 ± 0.002 per rating point (10 rating = 1%, 1.719 per %), spell_haste=-7.240 ± 0.381, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.180

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 68.5 spell_power points (37.77 DPS) | yes | Crimson Felt Hat (18727, -15.57 DPS, sim-verified) [dungeon]; Field Marshal's Coronal (231584, -17.59 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -22.88 DPS) [crafted] |
| neck | Orb of the Darkmoon (19426) (or Chains of the Lich (23125)) | 1200 Tickets - Orb of the Darkmoon [quest] | 22.0 spell_power points (12.13 DPS) | yes | Chains of the Lich (23125, -0.08 DPS, sim-verified) [dungeon]; Amulet of the Dawn (22657, -2.48 DPS) [quest]; Arcane Crystal Pendant (20037, -2.67 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 44.8 spell_power points (24.69 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -7.57 DPS, sim-verified) [vendor]; Field Marshal's Dreadweave Shoulders (231583, -9.11 DPS) [vendor]; Argent Shoulders (19059, -10.91 DPS) [crafted] |
| back | Crystalline Threaded Cape (20697) | World drop [world_drop] | 20.8 spell_power points (11.45 DPS) | yes | Amplifying Cloak (18350, -1.53 DPS) [dungeon]; Hide of the Wild (18510, -2.67 DPS) [crafted]; Arcanoweave Cloak (272411, -3.53 DPS, sim-verified) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 69.4 spell_power points (38.27 DPS) | yes | Robe of the Void (14153, -11.29 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -18.09 DPS) [vendor]; Robe of Everlasting Night (18385, -22.01 DPS) [dungeon] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 30.2 spell_power points (16.66 DPS) | yes | Dryad's Wrist Bindings (19596, -5.00 DPS) [rep]; Dryad's Wrist Bindings (19595, -5.68 DPS, sim-verified) [rep]; Dryad's Wrist Bindings (19597, -7.20 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 41.2 spell_power points (22.72 DPS) | yes | Marshal's Dreadweave Gloves (231586, -5.55 DPS) [vendor]; Hands of Power (13253, -7.75 DPS) [dungeon]; Sandworm Skin Gloves (20716, -7.88 DPS, sim-verified) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 45.4 spell_power points (25.01 DPS) | yes | Knowledge of the Timbermaw (228190, -8.41 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -11.97 DPS) [crafted]; Heretic Waistguard (240151, -13.24 DPS) [vendor] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 61.1 spell_power points (33.71 DPS) | yes | Skyshroud Leggings (13170, -9.68 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, -11.30 DPS) [vendor]; Sentinel's Silk Leggings (237815, -14.37 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 40.4 spell_power points (22.27 DPS) | yes | Marshal's Dreadweave Boots (231585, -6.57 DPS) [vendor]; Earthen Silk Slippers (254013, -6.80 DPS, sim-verified) [crafted]; Omnicast Boots (11822, -9.98 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.15 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234437, -0.66 DPS) [vendor]; Naglering (11669, -12.24 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 0.0 spell_power points (0.00 DPS) | yes | Songstone of Ironforge (12543, -2.63 DPS) [quest]; Maiden's Circle (13001, -2.63 DPS) [world_drop]; Naglering (11669, -9.12 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+16.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Weakness Analyzer (272438, -3.86 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -4.29 DPS, sim-verified) [quest]; Serenity Field (272439, -8.27 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | 0.0 spell_power points (0.00 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Lord Valthalak's Staff of Command (22335, -2.00 DPS) [dungeon]; Electrified Dagger (19100, -17.73 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (425.0 DPS) | yes | Bonecreeper Stylus (13938, -0.80 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.87 DPS) [world]; Torch of Light (279246, -5.54 DPS, sim-verified) [crafted] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Crystalline Threaded Cape; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1007, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 37.0. Weights run: 1.3s. Verify run: 1.2s. 142 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.491, intellect=0.360 ± 0.067, crit=0.135 ± 0.006 per rating point (14 rating = 1%, 1.886 per %), hit=0.466 ± 0.005 per rating point (10 rating = 1%, 4.660 per %), spell_haste=4.500 ± 0.499, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.491)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.30 DPS) | yes | Shadow Goggles (4373, -2.32 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.2 spell_power points (0.41 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.32 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.20 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.05 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.05 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.34 DPS) | yes | Green Woolen Robe (6243, -0.14 DPS) [crafted]; Bloody Apron (6226, -0.14 DPS) [dungeon]; Gray Woolen Robe (2585, -0.94 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.2 spell_power points (0.11 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Owlbeard Bracers (16981, -0.02 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.35 DPS) | yes | Gnoll Casting Gloves (892, -0.07 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.10 DPS) [crafted]; Apothecary Gloves (10919, -0.15 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.27 DPS) | yes | Novice Ardent's Sash (253887, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.13 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.69 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (37.0 DPS) | yes | Silk-threaded Trousers (1929, -0.06 DPS) [dungeon]; Rumpled Kilt (274741, -0.16 DPS) [vendor]; Abomination Skin Leggings (23173, -0.51 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.42 DPS) | yes | Pristine Boots (253889, -0.22 DPS) [crafted]; Red Woolen Boots (4313, -0.22 DPS) [crafted]; Feather Padded Treads (285345, -0.71 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.25 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Loop of Sacrifice (281673, -0.16 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.15 DPS) | yes | Loop of Sacrifice (281673, -0.06 DPS) [quest]; Volcanic Rock Ring (12053, -0.10 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.12 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 3.6 spell_power points (0.18 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.07 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 452.0 spell_power points (22.46 DPS) | yes | Skycaller (12984, -1.42 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.17 DPS) [dungeon]; Sizzle Stick (8071, -4.56 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 142, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 57.8. Weights run: 1.3s. Verify run: 1.2s. 237 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.160, intellect=0.439 ± 0.013, crit=0.041 ± 0.002 per rating point (14 rating = 1%, 0.567 per %), hit=0.158 ± 0.002 per rating point (10 rating = 1%, 1.578 per %), spell_haste=not significant (0.155 ± 0.172), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.160

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.86 DPS) | yes | Silk Headband (7050, -0.34 DPS) [crafted]; Embalmed Shroud (7691, -0.51 DPS) [dungeon]; Enchanter's Cowl (4322, -0.69 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (1.63 DPS) | yes | Crystal Starfire Medallion (5003, -1.33 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.33 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.68 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (2.19 DPS) | yes | Fairywing Mantle (9536, -0.51 DPS) [quest]; Death Speaker Mantle (6685, -0.53 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.64 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.85 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.17 DPS) [crafted]; Battle Healer's Cloak (19529, -0.17 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.7 spell_power points (2.49 DPS) | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.49 DPS) [dungeon]; Pristine Gown (253961, -0.78 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.52 DPS) | yes | Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.08 DPS) [quest]; Glowing Magical Bracelets (13106, -1.29 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.2 spell_power points (1.39 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.32 DPS) [crafted]; Gnoll Casting Gloves (892, -0.37 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.3 spell_power points (2.09 DPS) | yes | Warsong Sash (16975, +0.00 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.34 DPS) [dungeon]; Invoker's Cord (215366, -0.53 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.5 spell_power points (2.12 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.41 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.1 spell_power points (1.71 DPS) | yes | Acidic Walkers (9454, -0.26 DPS) [dungeon]; Boots of the Enchanter (4325, -0.86 DPS) [crafted]; Spidersilk Boots (4320, -1.19 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.19 DPS) | yes | Advisor's Ring (20426, -0.34 DPS) [rep]; Black Widow Band (6199, -0.66 DPS) [world]; Snake Hoop (6750, -0.66 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.02 DPS) | yes | Black Widow Band (6199, -0.50 DPS) [world]; Electrocutioner Lagnut (9447, -0.51 DPS) [dungeon]; Snake Hoop (6750, -2.20 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.52 DPS) | yes | Twisted Chanter's Staff (890, -0.78 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.78 DPS) [quest]; Glimmering Staff (249392, -0.84 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.6 spell_power points (1.63 DPS) | yes | Orb of Souls (249395, +0.00 DPS, sim-verified) [crafted]; Alliance Outrunner Healing Rod (285348, -0.95 DPS) [world]; Tome of the Darkspear Prophecy (272090, -1.00 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 199.2 spell_power points (33.72 DPS) | yes | Starfaller (13063, -0.48 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.87 DPS) [crafted]; Gravestone Scepter (7001, -4.72 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 237, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 112.0. Weights run: 1.1s. Verify run: 1.0s. 320 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.422, intellect=0.796 ± 0.026, crit=0.086 ± 0.005 per rating point (14 rating = 1%, 1.207 per %), hit=0.359 ± 0.004 per rating point (10 rating = 1%, 3.589 per %), spell_haste=not significant (1.639 ± 0.466), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.422)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (1.88 DPS) | yes | Corpseshroud (10574, -0.52 DPS) [dungeon]; Enchanter's Cowl (4322, -0.63 DPS) [crafted]; Augural Shroud (2620, -2.32 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.8 spell_power points (1.05 DPS) | yes | Necklace of Calisea (1714, -0.55 DPS) [world_drop]; Triune Amulet (7722, -0.55 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.45 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.4 spell_power points (1.55 DPS) | yes | Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.21 DPS) [quest]; Green Silken Shoulders (7057, -0.47 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.2 spell_power points (1.45 DPS) | yes | Guardian Cloak (5965, -0.55 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.66 DPS) [vendor]; Long Silken Cloak (4326, -2.21 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.8 spell_power points (2.39 DPS) | yes | Robe of Power (7054, -0.29 DPS) [crafted]; Elemental Raiment (9434, -0.52 DPS) [world_drop]; Dreamweave Vest (10021, -1.27 DPS, sim-verified) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.4 spell_power points (0.93 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.29 DPS) [world_drop]; Condor Bracers (15864, -0.30 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (1.89 DPS) | yes | Black Mageweave Gloves (10003, -0.55 DPS) [crafted]; Gilded Handwraps (254021, -0.68 DPS) [crafted]; Red Mageweave Gloves (10018, -1.94 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.9 spell_power points (1.69 DPS) | yes | Defiler's Cloth Girdle (20166, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.41 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.50 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.6 spell_power points (2.11 DPS) | yes | Abomination Skin Leggings (23173, -0.73 DPS) [dungeon]; Stoneweaver Leggings (9407, -0.91 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.55 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.15 DPS) | yes | Acidic Walkers (9454, -1.13 DPS) [dungeon]; Spidersilk Boots (4320, -1.24 DPS) [crafted]; Gilded Slippers (254001, -2.83 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (1.32 DPS) | yes | Reedknot Ring (9622, -0.70 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.78 DPS) [vendor]; Voodoo Band (1996, -0.82 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.80 DPS) | yes | Advisor's Ring (19521, -0.18 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.27 DPS) [vendor]; Reedknot Ring (9622, -1.03 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | sim-verified (112.0 DPS) | yes | Staff of Dar'Orahil (15106, -0.68 DPS) [quest]; Windweaver Staff (7757, -0.72 DPS) [dungeon]; Gut Ripper (2164, -7.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 445.2 spell_power points (39.81 DPS) | yes | Umbral Wand (5216, -0.23 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.22 DPS) [dungeon]; Twisted Nether Wand (249144, -5.27 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 320, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 205.0. Weights run: 1.3s. Verify run: 1.3s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.227, intellect=0.211 ± 0.015, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.457 per %), hit=0.203 ± 0.003 per rating point (10 rating = 1%, 2.027 per %), spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.89 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 17.3 spell_power points (4.84 DPS) | yes | Mindburst Medallion (11196, -2.80 DPS) [quest]; Horizon Choker (13085, -4.01 DPS) [world_drop]; Scorn's Icy Choker (23169, -4.72 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.8 spell_power points (4.71 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.37 DPS) [crafted]; Bloodmage Mantle (7684, -1.65 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.28 DPS) | yes | Deep Woodlands Cloak (19121, -0.40 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.22 DPS) [dungeon]; Runecloth Cloak (13860, -1.28 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (6.52 DPS) | yes | Elemental Raiment (9434, -0.64 DPS) [world_drop]; Dreamweave Vest (10021, -0.94 DPS) [crafted]; Acumen Robes (17775, -1.01 DPS, sim-verified) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest]; Bloodband Bracers (11469, -0.59 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.28 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.11 DPS) [vendor]; Black Mageweave Gloves (10003, -1.15 DPS, sim-verified) [crafted]; Runecloth Gloves (13863, -1.39 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (4.52 DPS) | yes | Defiler's Cloth Girdle (20166, -0.28 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Ghostweave Cord (254073, -0.59 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.1 spell_power points (7.04 DPS) | yes | Red Mageweave Pants (10009, -2.40 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.84 DPS) [vendor]; Wizardweave Leggings (14132, -4.47 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -1.41 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.23 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.39 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.49 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon]; Runed Ring (862, -1.68 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.36 DPS) | yes | Philanthropist's Ring (281635, -0.42 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.43 DPS) [dungeon]; Advisor's Ring (19520, -0.84 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.85 DPS) [crafted]; Rune of the Guard Captain (19120, -2.97 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.81 DPS, sim-verified) [crafted]; Rune of the Guard Captain (19120, -1.28 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Moonshadow Stave (22458, -0.55 DPS) [quest]; Arbiter's Blade (11784, -3.07 DPS) [dungeon]; Blade of Eternal Darkness (17780, -6.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (205.0 DPS) | yes | Woestave (20082, -0.08 DPS) [quest]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Pyric Caduceus (11748, -4.84 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Spidertank Oilrag; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Noxious Shooter

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 415.3. Weights run: 1.3s. Verify run: 1.2s. 998 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.180, intellect=0.192 ± 0.014, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), hit=0.172 ± 0.002 per rating point (10 rating = 1%, 1.719 per %), spell_haste=-7.240 ± 0.381, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.180

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 68.5 spell_power points (37.77 DPS) | yes | Crimson Felt Hat (18727, -14.89 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Hood (231590, -17.59 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -22.88 DPS) [crafted] |
| neck | Orb of the Darkmoon (19426) (or Chains of the Lich (23125)) | 1200 Tickets - Orb of the Darkmoon [quest] | 22.0 spell_power points (12.13 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Amulet of the Dawn (22657, -2.48 DPS) [quest]; Arcane Crystal Pendant (20037, -2.67 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 44.8 spell_power points (24.69 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -7.94 DPS, sim-verified) [vendor]; Warlord's Dreadweave Mantle (231592, -9.11 DPS) [vendor]; Argent Shoulders (19059, -10.91 DPS) [crafted] |
| back | Crystalline Threaded Cape (20697) | World drop [world_drop] | 20.8 spell_power points (11.45 DPS) | yes | Amplifying Cloak (18350, -1.53 DPS) [dungeon]; Hide of the Wild (18510, -2.67 DPS) [crafted]; Arcanoweave Cloak (272411, -3.14 DPS, sim-verified) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 69.4 spell_power points (38.27 DPS) | yes | Robe of the Void (14153, -13.04 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -18.09 DPS) [vendor]; Robe of Everlasting Night (18385, -22.01 DPS) [dungeon] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 30.2 spell_power points (16.66 DPS) | yes | Dryad's Wrist Bindings (19595, -4.88 DPS, sim-verified) [rep]; Dryad's Wrist Bindings (19596, -5.00 DPS) [rep]; Dryad's Wrist Bindings (19597, -7.20 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 41.2 spell_power points (22.72 DPS) | yes | General's Dreadweave Gloves (231589, -5.55 DPS) [vendor]; Sandworm Skin Gloves (20716, -7.49 DPS, sim-verified) [quest]; Hands of Power (13253, -7.75 DPS) [dungeon] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 45.4 spell_power points (25.01 DPS) | yes | Knowledge of the Timbermaw (228190, -7.56 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -11.97 DPS) [crafted]; Heretic Waistguard (240151, -13.24 DPS) [vendor] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 61.1 spell_power points (33.71 DPS) | yes | Skyshroud Leggings (13170, -10.32 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, -11.30 DPS) [vendor]; Sentinel's Silk Leggings (237815, -14.37 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 40.4 spell_power points (22.27 DPS) | yes | General's Dreadweave Boots (231593, -6.57 DPS) [vendor]; Omnicast Boots (11822, -9.98 DPS) [dungeon]; Earthen Silk Slippers (254013, -10.74 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | sim-verified (415.3 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.15 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234437, -0.66 DPS) [vendor]; Naglering (11669, -14.75 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (415.3 DPS) | yes | Eye of Orgrimmar (12545, -2.63 DPS) [quest]; Maiden's Circle (13001, -2.63 DPS) [world_drop]; Naglering (11669, -12.80 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (415.3 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (415.3 DPS) | yes | Royal Seal of Eldre'Thalas (18467, -3.31 DPS) [quest]; Weakness Analyzer (272438, -3.86 DPS) [vendor]; Draconic Infused Emblem (22268, -4.97 DPS, sim-verified) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (415.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Lord Valthalak's Staff of Command (22335, -2.00 DPS) [dungeon]; Glacial Blade (19099, -20.33 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 145.3 spell_power points (80.07 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -10.95 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.03 DPS) [world] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Crystalline Threaded Cape; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 998, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

