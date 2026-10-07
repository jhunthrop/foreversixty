# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 39.6. Weights run: 2.4s. Verify run: 1.1s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.062, intellect=0.089 ± 0.004, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.489 per %), hit=0.129 ± 0.001 per rating point (10 rating = 1%, 1.288 per %), spell_haste=0.358 ± 0.066, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.701 ± 0.062, fire_power=0.298 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.97 DPS) | yes | Shadow Goggles (4373, -3.22 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.8 spell_power points (0.94 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.07 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.65 DPS) | yes | Feyscale Cloak (6632, -0.16 DPS) [dungeon]; Caretaker's Cape (20428, -0.16 DPS) [rep]; Black Whelp Cloak (7283, -0.37 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.4 spell_power points (0.88 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -1.50 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.16 DPS) | yes | Mindthrust Bracers (1974, -0.09 DPS) [dungeon]; Bright Bracers (3647, -0.10 DPS) [world_drop]; Repurposed Hair Band (281256, -0.13 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.13 DPS) | yes | Gnoll Casting Gloves (892, -0.41 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.44 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.78 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.4 spell_power points (0.71 DPS) | yes | Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Keller's Girdle (2911, -0.59 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.97 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.7 spell_power points (1.57 DPS) | yes | Silk-threaded Trousers (1929, -0.45 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.51 DPS) [crafted]; Rumpled Kilt (274741, -0.76 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.4 spell_power points (1.19 DPS) | yes | Feather Padded Treads (285345, -0.43 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.54 DPS) [crafted]; Pristine Boots (253889, -0.66 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.2 spell_power points (0.84 DPS) | yes | Sludge-Stained Band (286535, -0.35 DPS) [world]; Lavishly Jeweled Ring (1156, -0.75 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.80 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.81 DPS) | yes | Sludge-Stained Band (286535, -0.60 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.72 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.77 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.9 spell_power points (0.14 DPS) | yes | Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.06 DPS) [world]; Staff of Westfall (2042, -0.07 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 140.7 spell_power points (22.80 DPS) | yes | Skycaller (12984, -1.87 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.51 DPS) [dungeon]; Deepblaze (279896, -4.27 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 67.7. Weights run: 2.4s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.108, intellect=0.295 ± 0.007, crit=0.068 ± 0.003 per rating point (14 rating = 1%, 0.950 per %), hit=0.144 ± 0.001 per rating point (10 rating = 1%, 1.444 per %), spell_haste=not significant (0.464 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.738 ± 0.108, fire_power=0.262 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.12 DPS) | yes | Enchanter's Cowl (4322, -0.39 DPS) [crafted]; Silk Headband (7050, -0.40 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.58 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 spell_power points (1.69 DPS) | yes | Crystal Starfire Medallion (5003, -1.46 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.46 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.30 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (2.25 DPS) | yes | Fairywing Mantle (9536, -0.58 DPS) [quest]; Invoker's Mantle (215365, -0.61 DPS) [crafted]; Death Speaker Mantle (6685, -0.67 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.96 DPS) | yes | Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Prelacy Cape (7004, -0.19 DPS) [quest]; Repairman's Cape (9605, -0.73 DPS, sim-verified) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.50 DPS) | yes | Death Speaker Robes (6682, -0.53 DPS) [dungeon]; Pristine Gown (253961, -0.76 DPS) [crafted]; Green Silk Armor (7065, -0.95 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.73 DPS) | yes | Nightsky Wristbands (6407, -1.39 DPS) [world_drop]; Stonecloth Bindings (14416, -1.45 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.86 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.2 spell_power points (1.40 DPS) | yes | Shilly Mitts (9609, +0.00 DPS, sim-verified) [quest]; Serpent Gloves (5970, -0.05 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.24 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 spell_power points (2.29 DPS) | yes | Belt of Arugal (6392, -0.33 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.66 DPS) [crafted]; Crimson Silk Belt (7055, -0.74 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.31 DPS) | yes | Pristine Leggings (253987, -0.57 DPS) [crafted]; Abomination Skin Leggings (23173, -0.72 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.81 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.1 spell_power points (1.75 DPS) | yes | Acidic Walkers (9454, -0.33 DPS) [dungeon]; Nimbus Boots (6998, -0.59 DPS) [quest]; Spidersilk Boots (4320, -2.23 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.35 DPS) | yes | Minor Channeling Ring (1449, -0.27 DPS) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Sludge-Stained Band (286535, -0.77 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.16 DPS) | yes | Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon]; Sludge-Stained Band (286535, -0.58 DPS) [world]; Minor Channeling Ring (1449, -2.17 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.73 DPS) | yes | Twisted Chanter's Staff (890, -1.16 DPS) [world_drop]; Channeler's Staff (4437, -1.28 DPS) [world]; Glimmering Staff (249392, -1.91 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.69 DPS) | yes | Dwarven Tome (279898, -0.84 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.92 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.92 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 175.4 spell_power points (33.79 DPS) | yes | Starfaller (13063, -1.02 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.82 DPS) [crafted]; Gravestone Scepter (7001, -4.79 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 137.5. Weights run: 2.1s. Verify run: 1.1s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.220, intellect=0.290 ± 0.013, crit=0.086 ± 0.003 per rating point (14 rating = 1%, 1.201 per %), hit=0.229 ± 0.002 per rating point (10 rating = 1%, 2.293 per %), spell_haste=0.941 ± 0.204, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.436 ± 0.220), fire_power=0.565 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.80 DPS) | yes | Augural Shroud (2620, -2.06 DPS, sim-verified) [world]; Living Cowl (5608, -2.21 DPS) [world]; Holy Shroud (2721, -2.76 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (2.41 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.46 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.85 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.85 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.6 spell_power points (3.21 DPS) | yes | Green Silken Shoulders (7057, -0.12 DPS) [crafted]; Inquisitor's Shawl (19507, -0.23 DPS) [dungeon]; Berylline Pads (4197, -0.47 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 11.6 spell_power points (3.21 DPS) | yes | Guardian Cloak (5965, -1.15 DPS) [crafted]; Icy Cloak (4327, -1.27 DPS) [crafted]; Long Silken Cloak (4326, -1.89 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.7 spell_power points (6.56 DPS) | yes | Elemental Raiment (9434, -0.60 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.86 DPS) [crafted]; Robe of Power (7054, -1.73 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.49 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.55 DPS) [quest]; Earthen Silk Cuffs (254019, -1.38 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (5.29 DPS) | yes | Black Mageweave Gloves (10003, -0.76 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.45 DPS) [crafted]; Gilded Handwraps (254021, -2.52 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.2 spell_power points (4.19 DPS) | yes | Star Belt (4329, -0.60 DPS) [crafted]; Deathmage Sash (10771, -1.05 DPS) [dungeon]; Gilded Cord (254037, -1.34 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.5 spell_power points (4.83 DPS) | yes | Gaze Dreamer Pants (6903, -1.51 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.58 DPS) [crafted]; Abomination Skin Leggings (23173, -1.70 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.63 DPS) | yes | Gilded Slippers (254001, -2.36 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.38 DPS) [crafted]; Acidic Walkers (9454, -4.61 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (3.24 DPS) | yes | Ring of Forlorn Spirits (2043, -1.03 DPS) [quest]; Reedknot Ring (9622, -1.31 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.59 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.49 DPS) | yes | Ring of Forlorn Spirits (2043, -0.28 DPS) [quest]; Reedknot Ring (9622, -0.55 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.83 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.04 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.01 DPS) [quest]; Gut Ripper (2164, -7.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (137.5 DPS) | yes | Twisted Nether Wand (249144, -0.01 DPS) [crafted]; Earthen Rod (9381, -0.08 DPS) [dungeon]; Jaina's Firestarter (13064, -1.69 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Umbral Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 191.2. Weights run: 2.0s. Verify run: 1.3s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.263, intellect=0.460 ± 0.017, crit=0.100 ± 0.004 per rating point (14 rating = 1%, 1.403 per %), hit=0.242 ± 0.002 per rating point (10 rating = 1%, 2.425 per %), spell_haste=1.195 ± 0.278, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.589 ± 0.263), fire_power=0.407 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.47 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.01 DPS) [crafted]; Red Mageweave Headband (10033, -2.56 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (3.27 DPS) | yes | Mindburst Medallion (11196, -0.34 DPS) [quest]; Horizon Choker (13085, -1.11 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.73 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 21.3 spell_power points (7.13 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -2.39 DPS) [crafted]; Red Mageweave Shoulders (10029, -2.47 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (5.62 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.73 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.37 DPS) [crafted]; Big Voodoo Cloak (8216, -2.55 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.2 spell_power points (9.45 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -2.03 DPS) [crafted]; Runecloth Tunic (13857, -2.06 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 10.2 spell_power points (3.43 DPS) | yes | Arcane Runed Bracers (4744, -0.41 DPS) [quest]; Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Bloodband Bracers (11469, -2.73 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (6.65 DPS) | yes | Raider Handwraps (272098, -0.91 DPS) [vendor]; Runecloth Gloves (13863, -1.24 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.63 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Cloth Girdle (20098, -0.49 DPS) [rep]; Dawnspire Cord (12466, -0.86 DPS) [dungeon]; Satyrmane Sash (17755, -2.40 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.6 spell_power points (9.25 DPS) | yes | Wizardweave Leggings (14132, -2.88 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.91 DPS) [vendor]; Red Mageweave Pants (10009, -4.76 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.05 DPS) | yes | Gilded Sandals (254107, -2.97 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.16 DPS) [vendor]; Black Mageweave Boots (10026, -3.28 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.36 DPS) | yes | Cyclopean Band (11824, -0.26 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.34 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (4.28 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19523, -0.25 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.37 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.23 DPS) [world_drop]; Spellshifter Rod (9527, -3.16 DPS) [quest]; Blade of Eternal Darkness (17780, -5.54 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.24 DPS) [crafted]; Wand of Allistarj (13065, -3.26 DPS) [world_drop]; Pyric Caduceus (11748, -6.83 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 392.6. Weights run: 9.2s. Verify run: 1.3s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± 1.380), intellect=not significant (-4.202 ± 0.108), crit=not significant (-1.319 ± 0.020) per rating point (14 rating = 1%, -18.463 per %), hit=not significant (-3.578 ± 0.016) per rating point (10 rating = 1%, -35.778 per %), spell_haste=not significant (7.767 ± 1.813), spell_penetration=not significant (-0.000 ± 0.000), shadow_power=not significant (4.588 ± 1.380), fire_power=not significant (-3.597 ± 0.004)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.7 spell_power points | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.00 DPS) [pvp]; Deathmist Mask (226909, -5.01 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.34 DPS) [quest]; Beads of Ogre Mojo (22149, -1.17 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.3 spell_power points | yes | Field Marshal's Dreadweave Shoulders (231583, -0.83 DPS) [pvp]; Burial Shawl (18681, -2.66 DPS) [dungeon]; Argent Shoulders (19059, -3.45 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.1 spell_power points | yes | Crystalline Threaded Cape (20697, -0.09 DPS) [world]; Hide of the Wild (18510, -1.17 DPS) [crafted]; Amplifying Cloak (18350, -1.38 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.1 spell_power points | yes | Field Marshal's Dreadweave Robe (231582, -2.38 DPS) [pvp]; Robe of Everlasting Night (18385, -4.45 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -5.34 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.7 spell_power points | yes | Sublime Wristguards (18497, -3.04 DPS) [dungeon]; Runecloth Cuffs (254123, -3.38 DPS) [crafted]; Deathmist Bracers (226907, -4.41 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.3 spell_power points | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -2.17 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.5 spell_power points | yes | Stormpike Cloth Girdle (19094, -5.01 DPS) [rep]; Belt of the Archmage (18405, -6.09 DPS, sim-verified) [crafted]; Satyrmane Sash (17755, -6.35 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 39.4 spell_power points | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -1.82 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.5 spell_power points | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.05 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.96 DPS) [quest]; Maiden's Circle (13001, -1.96 DPS) [world_drop]; Naglering (11669, -9.38 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.01 DPS) [quest]; Maiden's Circle (13001, -1.01 DPS) [world_drop]; Naglering (11669, -9.37 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+13.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.01 DPS) [quest]; Weakness Analyzer (272438, -2.35 DPS) [vendor]; Talisman of Ascendance (22678, -4.26 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Crackling Staff (19102, -0.55 DPS) [rep]; Teebu's Blazing Longsword (1728, -18.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (392.6 DPS) | yes | Bonecreeper Stylus (13938, -0.60 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.14 DPS) [world]; Torch of Light (279246, -18.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 38.4. Weights run: 2.4s. Verify run: 1.1s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.062, intellect=0.089 ± 0.004, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.489 per %), hit=0.129 ± 0.001 per rating point (10 rating = 1%, 1.288 per %), spell_haste=0.358 ± 0.066, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.701 ± 0.062, fire_power=0.298 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.97 DPS) | yes | Shadow Goggles (4373, -2.89 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.8 spell_power points (0.94 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.07 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.65 DPS) | yes | Feyscale Cloak (6632, -0.16 DPS) [dungeon]; Black Whelp Cloak (7283, -0.16 DPS) [crafted]; Battle Healer's Cloak (20427, -0.16 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.4 spell_power points (0.88 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -1.44 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.2 spell_power points (0.19 DPS) | yes | Windsong Bangles (263336, -0.03 DPS) [quest]; Tabitha's Cuffs (251486, -0.10 DPS) [quest]; Featherbead Bracers (15452, -0.12 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.13 DPS) | yes | Gnoll Casting Gloves (892, -0.18 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.44 DPS) [crafted]; Apothecary Gloves (10919, -0.49 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.4 spell_power points (0.71 DPS) | yes | Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Keller's Girdle (2911, -0.59 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.78 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.7 spell_power points (1.57 DPS) | yes | Silk-threaded Trousers (1929, -0.47 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.51 DPS) [crafted]; Rumpled Kilt (274741, -0.76 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.4 spell_power points (1.19 DPS) | yes | Red Woolen Boots (4313, -0.54 DPS) [crafted]; Feather Padded Treads (285345, -0.62 DPS, sim-verified) [world]; Pristine Boots (253889, -0.66 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.81 DPS) | yes | Lavishly Jeweled Ring (1156, -0.72 DPS) [dungeon]; Loop of Sacrifice (281673, -0.74 DPS) [quest]; Volcanic Rock Ring (12053, -0.77 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.49 DPS) | yes | Loop of Sacrifice (281673, -0.41 DPS) [quest]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.54 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.9 spell_power points (0.14 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.06 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 140.7 spell_power points (22.80 DPS) | yes | Skycaller (12984, -1.69 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.51 DPS) [dungeon]; Sizzle Stick (8071, -4.34 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 66.6. Weights run: 2.4s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.108, intellect=0.295 ± 0.007, crit=0.068 ± 0.003 per rating point (14 rating = 1%, 0.950 per %), hit=0.144 ± 0.001 per rating point (10 rating = 1%, 1.444 per %), spell_haste=not significant (0.464 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.738 ± 0.108, fire_power=0.262 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.12 DPS) | yes | Silk Headband (7050, -0.30 DPS, sim-verified) [crafted]; Enchanter's Cowl (4322, -0.39 DPS) [crafted]; Embalmed Shroud (7691, -0.58 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 spell_power points (1.69 DPS) | yes | Crystal Starfire Medallion (5003, -1.46 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.46 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.91 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (2.25 DPS) | yes | Fairywing Mantle (9536, -0.58 DPS) [quest]; Invoker's Mantle (215365, -0.61 DPS) [crafted]; Death Speaker Mantle (6685, -0.87 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.96 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Battle Healer's Cloak (19529, -0.19 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.50 DPS) | yes | Green Silk Armor (7065, -0.39 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.53 DPS) [dungeon]; Pristine Gown (253961, -0.76 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.73 DPS) | yes | Nightsky Wristbands (6407, -1.39 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.39 DPS) [quest]; Glowing Magical Bracelets (13106, -2.43 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.5 spell_power points (1.44 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.28 DPS) [world]; Truefaith Gloves (7049, -0.31 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 spell_power points (2.29 DPS) | yes | Warsong Sash (16975, -0.17 DPS) [quest]; Belt of Arugal (6392, -0.39 DPS) [dungeon]; Invoker's Cord (215366, -0.66 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.31 DPS) | yes | Abomination Skin Leggings (23173, -0.54 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.57 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.81 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.1 spell_power points (1.75 DPS) | yes | Acidic Walkers (9454, -0.33 DPS) [dungeon]; Boots of the Enchanter (4325, -0.78 DPS) [crafted]; Spidersilk Boots (4320, -2.02 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.35 DPS) | yes | Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Sludge-Stained Band (286535, -0.77 DPS) [world]; Black Widow Band (6199, -0.95 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.16 DPS) | yes | Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon]; Black Widow Band (6199, -0.76 DPS) [world]; Sludge-Stained Band (286535, -2.47 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.73 DPS) | yes | Twisted Chanter's Staff (890, -1.16 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.16 DPS) [quest]; Glimmering Staff (249392, -1.79 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.69 DPS) | yes | Orb of Souls (249395, -0.92 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.08 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.09 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 175.4 spell_power points (33.79 DPS) | yes | Starfaller (13063, -0.74 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.82 DPS) [crafted]; Gravestone Scepter (7001, -4.79 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 136.3. Weights run: 2.1s. Verify run: 1.2s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.220, intellect=0.290 ± 0.013, crit=0.086 ± 0.003 per rating point (14 rating = 1%, 1.201 per %), hit=0.229 ± 0.002 per rating point (10 rating = 1%, 2.293 per %), spell_haste=0.941 ± 0.204, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.436 ± 0.220), fire_power=0.565 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.80 DPS) | yes | Living Cowl (5608, -2.21 DPS) [world]; Holy Shroud (2721, -2.76 DPS) [world_drop]; Augural Shroud (2620, -2.78 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (2.41 DPS) | yes | Triune Amulet (7722, -1.85 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.85 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.37 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.6 spell_power points (3.21 DPS) | yes | Green Silken Shoulders (7057, -0.12 DPS) [crafted]; Inquisitor's Shawl (19507, -0.23 DPS) [dungeon]; Berylline Pads (4197, -0.47 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 11.6 spell_power points (3.21 DPS) | yes | Guardian Cloak (5965, -1.15 DPS) [crafted]; Icy Cloak (4327, -1.27 DPS) [crafted]; Long Silken Cloak (4326, -2.48 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.7 spell_power points (6.56 DPS) | yes | Elemental Raiment (9434, -0.86 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.86 DPS) [crafted]; Robe of Power (7054, -1.73 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.49 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.72 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.74 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (5.29 DPS) | yes | Black Mageweave Gloves (10003, -1.44 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.45 DPS) [crafted]; Gilded Handwraps (254021, -2.52 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.2 spell_power points (4.19 DPS) | yes | Star Belt (4329, -0.93 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.05 DPS) [dungeon]; Warsong Sash (16975, -1.15 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.5 spell_power points (4.83 DPS) | yes | Crimson Silk Pantaloons (7062, -1.58 DPS) [crafted]; Abomination Skin Leggings (23173, -1.70 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.92 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.63 DPS) | yes | Gilded Slippers (254001, -3.04 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.38 DPS) [crafted]; Acidic Walkers (9454, -4.61 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (3.24 DPS) | yes | Reedknot Ring (9622, -1.31 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.59 DPS) [vendor]; Sludge-Stained Band (286535, -2.41 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.49 DPS) | yes | Reedknot Ring (9622, -0.72 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.83 DPS) [vendor]; Sludge-Stained Band (286535, -1.66 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.04 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.01 DPS) [quest]; Gut Ripper (2164, -8.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (136.3 DPS) | yes | Twisted Nether Wand (249144, -0.01 DPS) [crafted]; Earthen Rod (9381, -0.08 DPS) [dungeon]; Jaina's Firestarter (13064, -1.47 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Umbral Wand

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 189.2. Weights run: 2.0s. Verify run: 1.3s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.263, intellect=0.460 ± 0.017, crit=0.100 ± 0.004 per rating point (14 rating = 1%, 1.403 per %), hit=0.242 ± 0.002 per rating point (10 rating = 1%, 2.425 per %), spell_haste=1.195 ± 0.278, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.589 ± 0.263), fire_power=0.407 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 28.2 spell_power points (9.45 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, +0.00 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -0.87 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.41 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (3.27 DPS) | yes | Mindburst Medallion (11196, -0.34 DPS) [quest]; Horizon Choker (13085, -1.11 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.73 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rotgrip Mantle (17732, -1.84 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.96 DPS) [crafted]; Red Mageweave Shoulders (10029, -2.04 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (5.62 DPS) | yes | Deep Woodlands Cloak (19121, -0.72 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.21 DPS) [dungeon]; Runecloth Cloak (13860, -1.37 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.2 spell_power points (9.45 DPS) | yes | Robe of the Magi (1716, -1.15 DPS) [world_drop]; Dreamweave Vest (10021, -2.03 DPS) [crafted]; Runecloth Tunic (13857, -2.06 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 10.2 spell_power points (3.43 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -2.05 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (6.65 DPS) | yes | Raider Handwraps (272098, -0.91 DPS) [vendor]; Runecloth Gloves (13863, -1.24 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.58 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Cloth Girdle (20166, -0.49 DPS) [rep]; Dawnspire Cord (12466, -0.86 DPS) [dungeon]; Satyrmane Sash (17755, -2.88 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.6 spell_power points (9.25 DPS) | yes | Wizardweave Leggings (14132, -2.88 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.91 DPS) [vendor]; Red Mageweave Pants (10009, -4.46 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.05 DPS) | yes | Gilded Sandals (254107, -2.97 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.16 DPS) [vendor]; Black Mageweave Boots (10026, -3.28 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.36 DPS) | yes | Cyclopean Band (11824, -0.26 DPS) [dungeon]; Advisor's Ring (19519, -0.34 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (4.28 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Advisor's Ring (19519, -0.25 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.69 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -3.45 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.16 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.23 DPS) [world_drop]; Spellshifter Rod (9527, -3.16 DPS) [quest]; Blade of Eternal Darkness (17780, -3.93 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.24 DPS) [crafted]; Wand of Allistarj (13065, -3.26 DPS) [world_drop]; Pyric Caduceus (11748, -7.28 DPS, sim-verified) [dungeon] |

**New at 50:** head: Red Mageweave Headband; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 388.1. Weights run: 9.2s. Verify run: 1.3s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± 1.380), intellect=not significant (-4.202 ± 0.108), crit=not significant (-1.319 ± 0.020) per rating point (14 rating = 1%, -18.463 per %), hit=not significant (-3.578 ± 0.016) per rating point (10 rating = 1%, -35.778 per %), spell_haste=not significant (7.767 ± 1.813), spell_penetration=not significant (-0.000 ± 0.000), shadow_power=not significant (4.588 ± 1.380), fire_power=not significant (-3.597 ± 0.004)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.7 spell_power points | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.00 DPS) [pvp]; Deathmist Mask (226909, -2.70 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.34 DPS) [quest]; Beads of Ogre Mojo (22149, -1.17 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.3 spell_power points | yes | Warlord's Dreadweave Mantle (231592, -0.83 DPS) [pvp]; Burial Shawl (18681, -1.37 DPS, sim-verified) [dungeon]; Argent Shoulders (19059, -3.45 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.1 spell_power points | yes | Crystalline Threaded Cape (20697, -0.09 DPS) [world]; Hide of the Wild (18510, -1.17 DPS) [crafted]; Amplifying Cloak (18350, -1.38 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.1 spell_power points | yes | Warlord's Dreadweave Robe (231591, -2.38 DPS) [pvp]; Robe of Everlasting Night (18385, -4.34 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -5.34 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.7 spell_power points | yes | Sublime Wristguards (18497, -3.04 DPS) [dungeon]; Runecloth Cuffs (254123, -3.38 DPS) [crafted]; Deathmist Bracers (226907, -4.41 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.3 spell_power points | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -2.17 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.5 spell_power points | yes | Frostwolf Cloth Belt (19090, -5.01 DPS) [rep]; Belt of the Archmage (18405, -6.10 DPS, sim-verified) [crafted]; Satyrmane Sash (17755, -6.35 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 39.4 spell_power points | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.58 DPS) [dungeon]; Outrider's Silk Leggings (22747, -0.89 DPS) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.5 spell_power points | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.05 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.51 DPS) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.96 DPS) [quest]; Maiden's Circle (13001, -1.96 DPS) [world_drop]; Naglering (11669, -9.10 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.01 DPS) [quest]; Maiden's Circle (13001, -1.01 DPS) [world_drop]; Naglering (11669, -8.58 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.01 DPS) [quest]; Weakness Analyzer (272438, -2.35 DPS) [vendor]; Talisman of Ascendance (22678, -6.60 DPS, sim-verified) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.31 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.11 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (388.1 DPS) | yes | Bonecreeper Stylus (13938, -0.60 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.14 DPS) [world]; Torch of Light (279246, -20.72 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

