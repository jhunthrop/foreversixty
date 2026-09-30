# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 33.0. Weights run: 0.6s. Verify run: 0.6s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.381 ± 0.010, crit=0.113 ± 0.004 per rating point (14 rating = 1%, 1.576 per %), hit=0.325 ± 0.003 per rating point (10 rating = 1%, 3.246 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.48 DPS) | yes | Shadow Goggles (4373, -0.71 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (0.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.36 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.32 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.56 DPS) | yes | Green Woolen Robe (6243, -0.22 DPS) [crafted]; Green Woolen Vest (2582, -0.23 DPS) [crafted]; Gray Woolen Robe (2585, -0.48 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.9 spell_power points (0.15 DPS) | yes | Windsong Bangles (263336, -0.07 DPS) [quest]; Repurposed Hair Band (281256, -0.09 DPS) [quest]; Bright Bracers (3647, -0.44 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.57 DPS) | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.34 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.45 DPS) | yes | Novice Ardent's Sash (253887, -0.19 DPS) [crafted]; Keller's Girdle (2911, -0.20 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.29 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (33.0 DPS) | yes | Silk-threaded Trousers (1929, -0.10 DPS) [dungeon]; Rumpled Kilt (274741, -0.27 DPS) [vendor]; Abomination Skin Leggings (23173, -0.34 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.69 DPS) | yes | Pristine Boots (253889, -0.35 DPS) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Feather Padded Treads (285345, -0.54 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.47 DPS) | yes | Sludge-Stained Band (286535, -0.22 DPS) [world]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.40 DPS) | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; Sludge-Stained Band (286535, -0.37 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.8 spell_power points (0.31 DPS) | yes | Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Channeler's Staff (4437, -0.14 DPS, sim-verified) [world]; Staff of Westfall (2042, -0.15 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 279.1 spell_power points (22.55 DPS) | yes | Skycaller (12984, -0.45 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.26 DPS) [dungeon]; Deepblaze (279896, -4.02 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 56.3. Weights run: 0.7s. Verify run: 0.6s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.444 ± 0.014, crit=0.110 ± 0.004 per rating point (14 rating = 1%, 1.540 per %), hit=0.309 ± 0.004 per rating point (10 rating = 1%, 3.087 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.26 DPS) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Enchanter's Cowl (4322, -0.25 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.34 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.7 spell_power points (1.11 DPS) | yes | Darkspear Warding Pendant (272075, -0.83 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.91 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.91 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (1.49 DPS) | yes | Death Speaker Mantle (6685, -0.13 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.57 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Prelacy Cape (7004, -0.11 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.8 spell_power points (1.70 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted]; Tree Bark Jacket (1486, -0.68 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Nightsky Wristbands (6407, -0.73 DPS) [world_drop]; Stonecloth Bindings (14416, -0.78 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.08 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.9 spell_power points (1.02 DPS) | yes | Shilly Mitts (9609, +0.00 DPS, sim-verified) [quest]; Serpent Gloves (5970, -0.22 DPS) [dungeon]; Truefaith Gloves (7049, -0.29 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.3 spell_power points (1.42 DPS) | yes | Belt of Arugal (6392, +0.00 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.36 DPS) [crafted]; Crimson Silk Belt (7055, -0.37 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | sim-verified (56.3 DPS) | yes | Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.38 DPS) [crafted]; Abomination Skin Leggings (23173, -0.58 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.1 spell_power points (1.16 DPS) | yes | Acidic Walkers (9454, -0.18 DPS) [dungeon]; Nimbus Boots (6998, -0.47 DPS) [quest]; Spidersilk Boots (4320, -1.14 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.80 DPS) | yes | Minor Channeling Ring (1449, -0.13 DPS) [quest]; Black Widow Band (6199, -0.45 DPS) [world]; Snake Hoop (6750, -0.45 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Black Widow Band (6199, -0.33 DPS) [world]; Snake Hoop (6750, -0.33 DPS) [quest]; Minor Channeling Ring (1449, -0.68 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Twisted Chanter's Staff (890, -0.52 DPS) [world_drop]; Channeler's Staff (4437, -0.63 DPS) [world]; Glimmering Staff (249392, -0.86 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.7 spell_power points (1.11 DPS) | yes | Dwarven Tome (279898, -0.37 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.65 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.65 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.9 spell_power points (33.55 DPS) | yes | Starfaller (13063, -0.36 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 188.9. Weights run: 0.7s. Verify run: 0.6s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.094 ± 0.003, crit=0.226 ± 0.006 per rating point (14 rating = 1%, 3.159 per %), hit=0.308 ± 0.004 per rating point (10 rating = 1%, 3.082 per %), spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.20 DPS) | yes | Living Cowl (5608, -2.77 DPS, sim-verified) [world]; Augural Shroud (2620, -3.11 DPS) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (2.59 DPS) | yes | Prodigious Shadowshard Pendant (17773, -2.18 DPS, sim-verified) [quest]; Triune Amulet (7722, -2.37 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.37 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (3.37 DPS) | yes | Green Silken Shoulders (7057, -0.23 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Berylline Pads (4197, -0.65 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 9.8 spell_power points (3.37 DPS) | yes | Long Silken Cloak (4326, -1.16 DPS) [crafted]; Guardian Cloak (5965, -1.16 DPS) [crafted]; Icy Cloak (4327, -2.20 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.6 spell_power points (7.73 DPS) | yes | Elemental Raiment (9434, -0.64 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.27 DPS) [crafted]; Robe of Power (7054, -2.55 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.08 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.69 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (6.30 DPS) | yes | Black Mageweave Gloves (10003, -1.25 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.4 spell_power points (4.93 DPS) | yes | Star Belt (4329, -0.55 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -1.75 DPS) [dungeon]; Gilded Cord (254037, -1.93 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 spell_power points (5.19 DPS) | yes | Gaze Dreamer Pants (6903, -1.34 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.02 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.23 DPS) | yes | Gilded Slippers (254001, -3.23 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Wingborne Boots (15104, -6.17 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.6 spell_power points (3.62 DPS) | yes | Ring of Forlorn Spirits (2043, -0.88 DPS) [quest]; Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (3.08 DPS) | yes | Ring of Forlorn Spirits (2043, -0.35 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.69 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (188.9 DPS) | yes | Scorn's Focal Dagger (23168, -3.77 DPS) [dungeon]; Illusionary Rod (7713, -5.58 DPS) [dungeon]; Gut Ripper (2164, -9.24 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 115.4 spell_power points (39.57 DPS) | yes | Nether Force Wand (11263, -0.15 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.29 DPS) [quest]; Ragefire Wand (7513, -2.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 277.2. Weights run: 0.8s. Verify run: 0.7s. 422 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.005, crit=0.336 ± 0.010 per rating point (14 rating = 1%, 4.698 per %), hit=0.469 ± 0.006 per rating point (10 rating = 1%, 4.690 per %), spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.46 DPS) | yes | Dreamweave Circlet (10041, -1.47 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.05 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.10 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 16.6 spell_power points (5.83 DPS) | yes | Mindburst Medallion (11196, -3.50 DPS) [quest]; Scorn's Icy Choker (23169, -5.10 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -5.30 DPS) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.4 spell_power points (5.40 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -0.61 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.55 DPS) [crafted]; Rotgrip Mantle (17732, -1.67 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.13 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.50 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.68 DPS) [crafted]; Nightfall Drape (12465, -1.98 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.6 spell_power points (7.93 DPS) | yes | Acumen Robes (17775, -0.13 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.58 DPS) [world_drop]; Dreamweave Vest (10021, -1.29 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.15 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.44 DPS) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (6.46 DPS) | yes | Sorcerer's Gauntlets (226930, +0.00 DPS, sim-verified) [vendor]; Black Mageweave Gloves (10003, -1.20 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.56 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.1 spell_power points (5.28 DPS) | yes | Highlander's Cloth Girdle (20097, -0.29 DPS) [rep]; Ghostweave Cord (254073, -0.38 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 24.1 spell_power points (8.44 DPS) | yes | Knight's Dreadweave Leggings (220888, -2.13 DPS) [vendor]; Red Mageweave Pants (10009, -3.08 DPS) [crafted]; Wizardweave Leggings (14132, -4.71 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.41 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -2.71 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -4.21 DPS) [crafted]; Black Mageweave Boots (10026, -4.29 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.55 DPS) | yes | Philanthropist's Ring (281635, -0.82 DPS) [quest]; Cyclopean Band (11824, -1.14 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.75 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.20 DPS) | yes | Philanthropist's Ring (281635, -0.30 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.79 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (277.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | - |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (277.2 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -3.85 DPS) [dungeon]; Arbiter's Blade (11784, -4.02 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; ranged: Pyric Caduceus

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 551.4. Weights run: 0.8s. Verify run: 0.7s. 1030 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.125 ± 0.006, crit=0.497 ± 0.014 per rating point (14 rating = 1%, 6.965 per %), hit=0.704 ± 0.009 per rating point (10 rating = 1%, 7.036 per %), spell_haste=4.546 ± 0.156, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 69.3 spell_power points (24.46 DPS) | yes | Fireleaf Hood (240048, -3.88 DPS, sim-verified) [vendor]; Field Marshal's Coronet (16441, -9.61 DPS) [vendor] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (551.4 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Arcane Crystal Pendant (20037, -1.85 DPS) [quest]; Jewel of Kajaro (19601, -4.12 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 49.8 spell_power points (17.58 DPS) | yes | Fireleaf Mantle (240046, -3.48 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -4.94 DPS) [vendor]; Field Marshal's Silk Spaulders (16444, -8.10 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.0 spell_power points (8.48 DPS) | yes | Amplifying Cloak (18350, -2.13 DPS) [dungeon]; Crystalline Threaded Cape (20697, -2.32 DPS, sim-verified) [world]; Hide of the Wild (18510, -3.10 DPS) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 79.3 spell_power points (27.97 DPS) | yes | Fireleaf Garb (240051, -3.77 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -10.88 DPS) [crafted]; Field Marshal's Silk Vestments (16443, -13.13 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 44.0 spell_power points (15.51 DPS) | yes | Fireleaf Wristwraps (240044, -3.14 DPS, sim-verified) [vendor]; Dryad's Wrist Bindings (19595, -7.39 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 57.7 spell_power points (20.35 DPS) | yes | Fireleaf Mitts (240049, -2.59 DPS, sim-verified) [vendor]; Marshal's Silk Gloves (16440, -10.29 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -10.29 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 54.7 spell_power points (19.28 DPS) | yes | Fireleaf Waistguard (240045, -2.27 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -7.01 DPS) [vendor]; Belt of the Archmage (18405, -9.07 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 69.3 spell_power points (24.46 DPS) | yes | Fireleaf Pants (240047, -4.22 DPS, sim-verified) [vendor]; Sentinel's Silk Leggings (237815, -7.82 DPS) [vendor]; Marshal's Silk Leggings (231605, -10.54 DPS) [pvp] |
| feet | Fireleaf Sandals (240058) | Leonid Barthalomew the Revered [vendor] | 51.7 spell_power points (18.24 DPS) | yes | Fireleaf Boots (240050, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (231606, -7.74 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (551.4 DPS) | yes | Songstone of Ironforge (12543, -6.10 DPS) [quest]; Maiden's Circle (13001, -6.10 DPS) [world_drop]; Naglering (11669, -17.88 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (551.4 DPS) | yes | Songstone of Ironforge (12543, -1.59 DPS) [quest]; Maiden's Circle (13001, -1.59 DPS) [world_drop]; Naglering (11669, -9.72 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (551.4 DPS) | yes | Weakness Analyzer (272438, -2.47 DPS) [vendor]; Serenity Field (272439, -5.29 DPS) [vendor]; Burst of Knowledge (11832, -6.00 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (551.4 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -3.24 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (551.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.12 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 223.6 spell_power points (78.88 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.19 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.39 DPS) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Chains of the Lich; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1030, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 29.4. Weights run: 0.6s. Verify run: 0.6s. 141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.381 ± 0.010, crit=0.113 ± 0.004 per rating point (14 rating = 1%, 1.576 per %), hit=0.325 ± 0.003 per rating point (10 rating = 1%, 3.246 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.48 DPS) | yes | Shadow Goggles (4373, -0.45 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (0.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.20 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.36 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.32 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.56 DPS) | yes | Green Woolen Robe (6243, -0.22 DPS) [crafted]; Green Woolen Vest (2582, -0.23 DPS) [crafted]; Gray Woolen Robe (2585, -0.47 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.3 spell_power points (0.18 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS, sim-verified) [quest]; Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Owlbeard Bracers (16981, -0.04 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.57 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Apothecary Gloves (10919, -0.24 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.45 DPS) | yes | Novice Arcanist's Sash (253885, -0.12 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.19 DPS) [crafted]; Keller's Girdle (2911, -0.20 DPS) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.0 spell_power points (0.97 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Rumpled Kilt (274741, -0.57 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.69 DPS) | yes | Pristine Boots (253889, -0.35 DPS) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Feather Padded Treads (285345, -0.40 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.40 DPS) | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Loop of Sacrifice (281673, -0.25 DPS) [quest]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.24 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.09 DPS) [quest]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.8 spell_power points (0.31 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.12 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 279.1 spell_power points (22.55 DPS) | yes | Skycaller (12984, -0.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.26 DPS) [dungeon]; Sizzle Stick (8071, -4.50 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 49.4. Weights run: 0.7s. Verify run: 0.6s. 236 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.444 ± 0.014, crit=0.110 ± 0.004 per rating point (14 rating = 1%, 1.540 per %), hit=0.309 ± 0.004 per rating point (10 rating = 1%, 3.087 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.26 DPS) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Enchanter's Cowl (4322, -0.28 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.34 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.7 spell_power points (1.11 DPS) | yes | Darkspear Warding Pendant (272075, -0.72 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.91 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.91 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (1.49 DPS) | yes | Death Speaker Mantle (6685, -0.25 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.57 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.8 spell_power points (1.70 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Tree Bark Jacket (1486, -0.44 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Nightsky Wristbands (6407, -0.73 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.73 DPS) [quest]; Glowing Magical Bracelets (13106, -0.92 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.2 spell_power points (0.94 DPS) | yes | Serpent Gloves (5970, -0.09 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.26 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.3 spell_power points (1.42 DPS) | yes | Warsong Sash (16975, -0.16 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.36 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.5 spell_power points (1.44 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.1 spell_power points (1.16 DPS) | yes | Acidic Walkers (9454, -0.18 DPS) [dungeon]; Boots of the Enchanter (4325, -0.59 DPS) [crafted]; Spidersilk Boots (4320, -1.08 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.80 DPS) | yes | Black Widow Band (6199, -0.45 DPS) [world]; Snake Hoop (6750, -0.45 DPS) [quest]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Snake Hoop (6750, -0.33 DPS) [quest]; Sludge-Stained Band (286535, -0.34 DPS) [world]; Black Widow Band (6199, -0.92 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Twisted Chanter's Staff (890, -0.52 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.52 DPS) [quest]; Glimmering Staff (249392, -0.57 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.7 spell_power points (1.11 DPS) | yes | Orb of Souls (249395, -0.65 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.68 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -0.71 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.9 spell_power points (33.55 DPS) | yes | Starfaller (13063, -0.34 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 236, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 181.7. Weights run: 0.7s. Verify run: 0.6s. 318 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.094 ± 0.003, crit=0.226 ± 0.006 per rating point (14 rating = 1%, 3.159 per %), hit=0.308 ± 0.004 per rating point (10 rating = 1%, 3.082 per %), spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.20 DPS) | yes | Living Cowl (5608, -2.66 DPS, sim-verified) [world]; Augural Shroud (2620, -3.11 DPS) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (2.59 DPS) | yes | Prodigious Shadowshard Pendant (17773, -2.13 DPS, sim-verified) [quest]; Triune Amulet (7722, -2.37 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.37 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (3.37 DPS) | yes | Green Silken Shoulders (7057, -0.24 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Chestnut Mantle (17695, -0.63 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 9.8 spell_power points (3.37 DPS) | yes | Long Silken Cloak (4326, -1.16 DPS) [crafted]; Guardian Cloak (5965, -1.16 DPS) [crafted]; Icy Cloak (4327, -2.15 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.6 spell_power points (7.73 DPS) | yes | Elemental Raiment (9434, -0.69 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.27 DPS) [crafted]; Robe of Power (7054, -2.55 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.08 DPS) | yes | Condor Bracers (15864, -0.67 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.46 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (6.30 DPS) | yes | Black Mageweave Gloves (10003, -1.22 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.4 spell_power points (4.93 DPS) | yes | Star Belt (4329, -0.55 DPS, sim-verified) [crafted]; Warsong Sash (16975, -1.16 DPS) [quest]; Belt of Arugal (6392, -1.75 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 spell_power points (5.19 DPS) | yes | Gaze Dreamer Pants (6903, -1.29 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.02 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.23 DPS) | yes | Gilded Slippers (254001, -3.10 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Acidic Walkers (9454, -6.26 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.6 spell_power points (3.62 DPS) | yes | Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor]; Sludge-Stained Band (286535, -2.59 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.08 DPS) | yes | Reedknot Ring (9622, -0.67 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor]; Sludge-Stained Band (286535, -2.06 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (181.7 DPS) | yes | Scorn's Focal Dagger (23168, -3.77 DPS) [dungeon]; Illusionary Rod (7713, -5.58 DPS) [dungeon]; Gut Ripper (2164, -8.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 115.4 spell_power points (39.57 DPS) | yes | Nether Force Wand (11263, -0.17 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.29 DPS) [quest]; Ragefire Wand (7513, -2.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 266.0. Weights run: 0.8s. Verify run: 0.7s. 412 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.005, crit=0.336 ± 0.010 per rating point (14 rating = 1%, 4.698 per %), hit=0.469 ± 0.006 per rating point (10 rating = 1%, 4.690 per %), spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.46 DPS) | yes | Dreamweave Circlet (10041, -1.48 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.05 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.10 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 16.6 spell_power points (5.83 DPS) | yes | Mindburst Medallion (11196, -3.50 DPS) [quest]; Scorn's Icy Choker (23169, -4.93 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -5.30 DPS) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.4 spell_power points (5.40 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -0.61 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.55 DPS) [crafted]; Rotgrip Mantle (17732, -1.67 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.13 DPS) | yes | Deep Woodlands Cloak (19121, -0.49 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.64 DPS) [dungeon]; Runecloth Cloak (13860, -1.68 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.6 spell_power points (7.93 DPS) | yes | Acumen Robes (17775, -0.27 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.58 DPS) [world_drop]; Dreamweave Vest (10021, -1.29 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.15 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest]; Bloodband Bracers (11469, -1.06 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (6.46 DPS) | yes | Sorcerer's Gauntlets (226930, +0.00 DPS, sim-verified) [vendor]; Black Mageweave Gloves (10003, -1.20 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.56 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.1 spell_power points (5.28 DPS) | yes | Defiler's Cloth Girdle (20165, -0.29 DPS) [rep]; Ghostweave Cord (254073, -0.38 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 24.1 spell_power points (8.44 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -2.13 DPS) [vendor]; Red Mageweave Pants (10009, -3.08 DPS) [crafted]; Wizardweave Leggings (14132, -4.56 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.41 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.32 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -4.21 DPS) [crafted]; Black Mageweave Boots (10026, -4.29 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.55 DPS) | yes | Philanthropist's Ring (281635, -0.82 DPS) [quest]; Cyclopean Band (11824, -1.14 DPS) [dungeon]; Runed Ring (862, -2.10 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.20 DPS) | yes | Philanthropist's Ring (281635, -0.36 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.79 DPS) [dungeon]; Runed Ring (862, -1.75 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (266.0 DPS) | yes | Uther's Strength (11302, -0.09 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (266.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (266.0 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -3.85 DPS) [dungeon]; Arbiter's Blade (11784, -4.02 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; ranged: Pyric Caduceus

No-known-source sample (15 of 412, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 529.3. Weights run: 0.8s. Verify run: 0.8s. 1021 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.125 ± 0.006, crit=0.497 ± 0.014 per rating point (14 rating = 1%, 6.965 per %), hit=0.704 ± 0.009 per rating point (10 rating = 1%, 7.036 per %), spell_haste=4.546 ± 0.156, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 69.3 spell_power points (24.46 DPS) | yes | Fireleaf Hood (240048, -3.76 DPS, sim-verified) [vendor]; Warlord's Silk Cowl (16533, -9.61 DPS) [vendor] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (529.3 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Arcane Crystal Pendant (20037, -1.85 DPS) [quest]; Jewel of Kajaro (19601, -3.91 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 49.8 spell_power points (17.58 DPS) | yes | Fireleaf Mantle (240046, -3.38 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -4.94 DPS) [vendor]; Warlord's Silk Amice (16536, -8.10 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.0 spell_power points (8.48 DPS) | yes | Crystalline Threaded Cape (20697, -1.98 DPS, sim-verified) [world]; Amplifying Cloak (18350, -2.13 DPS) [dungeon]; Hide of the Wild (18510, -3.10 DPS) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 79.3 spell_power points (27.97 DPS) | yes | Fireleaf Garb (240051, -3.66 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -10.88 DPS) [crafted]; Warlord's Silk Raiment (16535, -13.13 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 44.0 spell_power points (15.51 DPS) | yes | Fireleaf Wristwraps (240044, -3.03 DPS, sim-verified) [vendor]; Dryad's Wrist Bindings (19595, -7.39 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 57.7 spell_power points (20.35 DPS) | yes | Fireleaf Mitts (240049, -2.52 DPS, sim-verified) [vendor]; General's Silk Handguards (16540, -10.29 DPS) [vendor]; General's Silk Gauntlets (231599, -10.29 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 54.7 spell_power points (19.28 DPS) | yes | Fireleaf Waistguard (240045, -2.03 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -7.01 DPS) [vendor]; Belt of the Archmage (18405, -9.07 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 69.3 spell_power points (24.46 DPS) | yes | Fireleaf Pants (240047, -4.11 DPS, sim-verified) [vendor]; Sentinel's Silk Leggings (237815, -7.82 DPS) [vendor]; General's Silk Trousers (16534, -10.54 DPS) [vendor] |
| feet | Fireleaf Sandals (240058) | Leonid Barthalomew the Revered [vendor] | 51.7 spell_power points (18.24 DPS) | yes | Fireleaf Boots (240050, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (231597, -7.74 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (529.3 DPS) | yes | Eye of Orgrimmar (12545, -6.10 DPS) [quest]; Maiden's Circle (13001, -6.10 DPS) [world_drop]; Naglering (11669, -17.04 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (529.3 DPS) | yes | Eye of Orgrimmar (12545, -1.59 DPS) [quest]; Maiden's Circle (13001, -1.59 DPS) [world_drop]; Naglering (11669, -9.34 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (529.3 DPS) | yes | Weakness Analyzer (272438, -2.47 DPS) [vendor]; Serenity Field (272439, -5.29 DPS) [vendor]; Burst of Knowledge (11832, -6.00 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (529.3 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -2.43 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (529.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.12 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 223.6 spell_power points (78.88 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.19 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.39 DPS) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Chains of the Lich; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1021, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

