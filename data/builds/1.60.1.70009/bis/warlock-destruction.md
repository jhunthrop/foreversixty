# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 39.5. Weights run: 2.4s. Verify run: 1.1s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.064, intellect=0.108 ± 0.004, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.485 per %), hit=0.128 ± 0.001 per rating point (10 rating = 1%, 1.277 per %), spell_haste=0.465 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.707 ± 0.064, fire_power=0.291 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.98 DPS) | yes | Shadow Goggles (4373, -3.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.0 spell_power points (0.97 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.32 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.39 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.65 DPS) | yes | Feyscale Cloak (6632, -0.16 DPS) [dungeon]; Caretaker's Cape (20428, -0.16 DPS) [rep]; Black Whelp Cloak (7283, -0.20 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.5 spell_power points (0.90 DPS) | yes | Green Woolen Vest (2582, -0.25 DPS) [crafted]; Bloody Apron (6226, -0.25 DPS) [dungeon]; Gray Woolen Robe (2585, -1.52 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.16 DPS) | yes | Bright Bracers (3647, -0.09 DPS) [world_drop]; Repurposed Hair Band (281256, -0.13 DPS) [quest]; Mindthrust Bracers (1974, -0.35 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.14 DPS) | yes | Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.44 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.78 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.4 spell_power points (0.72 DPS) | yes | Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Keller's Girdle (2911, -0.58 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.11 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.9 spell_power points (1.61 DPS) | yes | Filigreed Pristine Leggings (253937, -0.52 DPS) [crafted]; Rumpled Kilt (274741, -0.79 DPS) [vendor]; Silk-threaded Trousers (1929, -0.80 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.4 spell_power points (1.21 DPS) | yes | Red Woolen Boots (4313, -0.56 DPS) [crafted]; Pristine Boots (253889, -0.67 DPS) [crafted]; Feather Padded Treads (285345, -0.72 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.2 spell_power points (0.85 DPS) | yes | Sludge-Stained Band (286535, -0.36 DPS) [world]; Lavishly Jeweled Ring (1156, -0.75 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.80 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.82 DPS) | yes | Lavishly Jeweled Ring (1156, -0.71 DPS) [dungeon]; Sludge-Stained Band (286535, -0.73 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.76 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 1.1 spell_power points (0.18 DPS) | yes | Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.07 DPS) [world]; Staff of Westfall (2042, -0.09 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 139.7 spell_power points (22.80 DPS) | yes | Skycaller (12984, -1.79 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.51 DPS) [dungeon]; Deepblaze (279896, -4.27 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 67.2. Weights run: 2.4s. Verify run: 1.3s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.191, intellect=0.468 ± 0.014, crit=0.102 ± 0.004 per rating point (14 rating = 1%, 1.429 per %), hit=0.230 ± 0.002 per rating point (10 rating = 1%, 2.299 per %), spell_haste=not significant (0.475 ± 0.198), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.567 ± 0.191), fire_power=0.430 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.27 DPS) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Embalmed Shroud (7691, -0.35 DPS) [dungeon]; Enchanter's Cowl (4322, -1.19 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.13 DPS) | yes | Crystal Starfire Medallion (5003, -0.91 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.91 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.24 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (1.52 DPS) | yes | Fairywing Mantle (9536, -0.35 DPS) [quest]; Invoker's Mantle (215365, -0.45 DPS) [crafted]; Death Speaker Mantle (6685, -0.97 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.58 DPS) | yes | Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Prelacy Cape (7004, -0.12 DPS) [quest]; Repairman's Cape (9605, -0.75 DPS, sim-verified) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.10 DPS) [dungeon]; Pristine Gown (253961, -0.31 DPS) [crafted]; Green Silk Armor (7065, -1.10 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Nightsky Wristbands (6407, -0.71 DPS) [world_drop]; Stonecloth Bindings (14416, -0.77 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.75 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Truefaith Gloves (7049, -0.07 DPS) [crafted]; Town Clerk's Mittens (270029, -0.72 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.4 spell_power points (1.43 DPS) | yes | Invoker's Cord (215366, -0.35 DPS) [crafted]; Crimson Silk Belt (7055, -0.36 DPS) [crafted]; Belt of Arugal (6392, -0.67 DPS, sim-verified) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.7 spell_power points (1.47 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.18 DPS) | yes | Acidic Walkers (9454, -0.18 DPS) [dungeon]; Nimbus Boots (6998, -0.49 DPS) [quest]; Spidersilk Boots (4320, -2.12 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.81 DPS) | yes | Minor Channeling Ring (1449, -0.12 DPS) [quest]; Black Widow Band (6199, -0.43 DPS) [world]; Snake Hoop (6750, -0.43 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Black Widow Band (6199, -0.31 DPS) [world]; Snake Hoop (6750, -0.31 DPS) [quest]; Minor Channeling Ring (1449, -2.11 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Twisted Chanter's Staff (890, -0.50 DPS) [world_drop]; Channeler's Staff (4437, -0.61 DPS) [world]; Glimmering Staff (249392, -1.82 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.8 spell_power points (1.13 DPS) | yes | Eye of Paleth (2943, -0.67 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.67 DPS) [world]; Dwarven Tome (279898, -0.79 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.1 spell_power points (33.56 DPS) | yes | Starfaller (13063, -0.94 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 136.7. Weights run: 2.1s. Verify run: 1.1s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.392, intellect=0.637 ± 0.023, crit=0.144 ± 0.005 per rating point (14 rating = 1%, 2.009 per %), hit=0.353 ± 0.003 per rating point (10 rating = 1%, 3.529 per %), spell_haste=not significant (1.212 ± 0.340), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.049 ± 0.392), fire_power=0.952 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.42 DPS) | yes | Living Cowl (5608, -1.30 DPS) [world]; Enchanter's Cowl (4322, -1.40 DPS) [crafted]; Augural Shroud (2620, -2.16 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.76 DPS) | yes | Triune Amulet (7722, -1.04 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.04 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.74 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.3 spell_power points (2.49 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.31 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.7 spell_power points (2.40 DPS) | yes | Guardian Cloak (5965, -0.90 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.26 DPS) [vendor]; Long Silken Cloak (4326, -2.30 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.8 spell_power points (4.20 DPS) | yes | Robe of Power (7054, -0.68 DPS) [crafted]; Elemental Raiment (9434, -0.79 DPS) [world_drop]; Dreamweave Vest (10021, -1.11 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.46 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.33 DPS) [quest]; Windchaser Cuffs (14429, -0.53 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (3.34 DPS) | yes | Black Mageweave Gloves (10003, -0.90 DPS) [crafted]; Red Mageweave Gloves (10018, -1.03 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.32 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -0.56 DPS) [crafted]; Star Belt (4329, -0.58 DPS) [crafted]; Deathmage Sash (10771, -2.24 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.6 spell_power points (3.52 DPS) | yes | Abomination Skin Leggings (23173, -1.23 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.26 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.55 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.90 DPS) | yes | Gilded Slippers (254001, -2.04 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.26 DPS) [dungeon]; Spidersilk Boots (4320, -2.35 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (2.25 DPS) | yes | Ring of Forlorn Spirits (2043, -0.95 DPS) [quest]; Reedknot Ring (9622, -1.11 DPS) [quest]; Minor Channeling Ring (1449, -1.23 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.46 DPS) | yes | Ring of Forlorn Spirits (2043, -0.16 DPS) [quest]; Reedknot Ring (9622, -0.33 DPS) [quest]; Minor Channeling Ring (1449, -0.44 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Dar'Orahil (15106, -1.54 DPS) [quest]; Windweaver Staff (7757, -1.70 DPS) [dungeon]; Gut Ripper (2164, -8.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Earthen Rod (9381, -0.08 DPS) [dungeon]; Twisted Nether Wand (249144, -0.69 DPS) [crafted]; Jaina's Firestarter (13064, -1.45 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Umbral Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 192.4. Weights run: 2.0s. Verify run: 1.3s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.197, intellect=0.442 ± 0.014, crit=0.083 ± 0.003 per rating point (14 rating = 1%, 1.165 per %), hit=0.193 ± 0.002 per rating point (10 rating = 1%, 1.927 per %), spell_haste=not significant (-0.249 ± 0.229), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.672 ± 0.197), fire_power=0.327 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 27.8 spell_power points (11.50 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, +0.00 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.00 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.82 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.7 spell_power points (3.99 DPS) | yes | Mindburst Medallion (11196, -0.41 DPS) [quest]; Horizon Choker (13085, -1.43 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.16 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Rotgrip Mantle (17732, -2.08 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -2.38 DPS) [crafted]; Red Mageweave Shoulders (10029, -2.53 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.7 spell_power points (6.88 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.52 DPS) [dungeon]; Runecloth Cloak (13860, -1.70 DPS) [crafted]; Nightfall Drape (12465, -3.16 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 27.8 spell_power points (11.50 DPS) | yes | Robe of the Magi (1716, -1.32 DPS) [world_drop]; Dreamweave Vest (10021, -2.42 DPS) [crafted]; Runecloth Tunic (13857, -2.47 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 10.1 spell_power points (4.17 DPS) | yes | Spidertank Oilrag (9448, -0.45 DPS) [dungeon]; Bloodband Bracers (11469, -0.46 DPS) [quest]; Arcane Runed Bracers (4744, -1.29 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (8.16 DPS) | yes | Raider Handwraps (272098, -1.26 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.44 DPS, sim-verified) [vendor]; Runecloth Gloves (13863, -1.57 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Cloth Girdle (20098, -0.53 DPS) [rep]; Dawnspire Cord (12466, -1.10 DPS) [dungeon]; Satyrmane Sash (17755, -3.19 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.4 spell_power points (11.32 DPS) | yes | Wizardweave Leggings (14132, -3.48 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.70 DPS) [vendor]; Red Mageweave Pants (10009, -5.10 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.91 DPS) | yes | Gilded Sandals (254107, -1.05 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -4.09 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.17 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.37 DPS) | yes | Cyclopean Band (11824, -0.37 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.41 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (5.23 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19523, -0.27 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.29 DPS, sim-verified) [crafted]; Uther's Strength (11302, -2.48 DPS) [world_drop] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.83 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, -4.24 DPS) [crafted] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.97 DPS) [world_drop]; Arbiter's Blade (11784, -4.04 DPS) [dungeon]; Blade of Eternal Darkness (17780, -4.83 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+6.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.01 DPS) [crafted]; Wand of Allistarj (13065, -3.64 DPS) [world_drop]; Pyric Caduceus (11748, -6.26 DPS, sim-verified) [dungeon] |

**New at 50:** head: Red Mageweave Headband; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Abyss Shard; trinket2: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 392.5. Weights run: 2.2s. Verify run: 1.3s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.342, intellect=0.402 ± 0.023, crit=0.143 ± 0.004 per rating point (14 rating = 1%, 2.006 per %), hit=0.371 ± 0.003 per rating point (10 rating = 1%, 3.715 per %), spell_haste=not significant (-1.237 ± 0.374), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.625 ± 0.342), fire_power=0.375 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.2 spell_power points (15.36 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.37 DPS) [pvp]; Deathmist Mask (226909, -5.31 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.18 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.82 DPS) [quest]; Beads of Ogre Mojo (22149, -1.93 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.0 spell_power points (16.21 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.48 DPS) [pvp]; Burial Shawl (18681, -1.76 DPS, sim-verified) [dungeon]; Argent Shoulders (19059, -4.64 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.9 spell_power points (10.61 DPS) | yes | Crystalline Threaded Cape (20697, -0.61 DPS) [world]; Hide of the Wild (18510, -2.27 DPS) [crafted]; Amplifying Cloak (18350, -2.28 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.6 spell_power points (22.95 DPS) | yes | Robe of Everlasting Night (18385, -3.31 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -3.68 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.67 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.2 spell_power points (11.66 DPS) | yes | Sublime Wristguards (18497, -4.25 DPS) [dungeon]; Runecloth Cuffs (254123, -4.72 DPS) [crafted]; Deathmist Bracers (226907, -6.19 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.0 spell_power points (13.42 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -2.96 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.6 spell_power points (17.38 DPS) | yes | Belt of the Archmage (18405, -4.80 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -7.19 DPS) [rep]; Highlander's Cloth Girdle (20047, -8.86 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.0 spell_power points (18.52 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -3.15 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.8 spell_power points (11.48 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.18 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -2.59 DPS) [quest]; Maiden's Circle (13001, -2.59 DPS) [world_drop]; Naglering (11669, -9.93 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.39 DPS) [quest]; Maiden's Circle (13001, -1.39 DPS) [world_drop]; Naglering (11669, -9.49 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+16.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.78 DPS) [quest]; Weakness Analyzer (272438, -3.24 DPS) [vendor]; Serenity Field (272439, -6.94 DPS) [vendor] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -1.78 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Crackling Staff (19102, -1.18 DPS) [rep]; Teebu's Blazing Longsword (1728, -19.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (392.5 DPS) | yes | Bonecreeper Stylus (13938, -0.48 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.62 DPS) [world]; Torch of Light (279246, -19.56 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 38.6. Weights run: 2.4s. Verify run: 1.1s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.064, intellect=0.108 ± 0.004, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.485 per %), hit=0.128 ± 0.001 per rating point (10 rating = 1%, 1.277 per %), spell_haste=0.465 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.707 ± 0.064, fire_power=0.291 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.98 DPS) | yes | Shadow Goggles (4373, -3.25 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.0 spell_power points (0.97 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.32 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.65 DPS) | yes | Feyscale Cloak (6632, -0.16 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.16 DPS) [rep]; Black Whelp Cloak (7283, -0.34 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.5 spell_power points (0.90 DPS) | yes | Green Woolen Vest (2582, -0.25 DPS) [crafted]; Bloody Apron (6226, -0.25 DPS) [dungeon]; Gray Woolen Robe (2585, -1.51 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.2 spell_power points (0.20 DPS) | yes | Tabitha's Cuffs (251486, -0.09 DPS) [quest]; Featherbead Bracers (15452, -0.11 DPS) [quest]; Windsong Bangles (263336, -0.19 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.14 DPS) | yes | Gnoll Casting Gloves (892, -0.40 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.44 DPS) [crafted]; Apothecary Gloves (10919, -0.49 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.4 spell_power points (0.72 DPS) | yes | Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Keller's Girdle (2911, -0.58 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.17 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.9 spell_power points (1.61 DPS) | yes | Filigreed Pristine Leggings (253937, -0.52 DPS) [crafted]; Silk-threaded Trousers (1929, -0.75 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.79 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.4 spell_power points (1.21 DPS) | yes | Red Woolen Boots (4313, -0.56 DPS) [crafted]; Pristine Boots (253889, -0.67 DPS) [crafted]; Feather Padded Treads (285345, -0.68 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.82 DPS) | yes | Lavishly Jeweled Ring (1156, -0.71 DPS) [dungeon]; Loop of Sacrifice (281673, -0.73 DPS) [quest]; Volcanic Rock Ring (12053, -0.76 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.49 DPS) | yes | Loop of Sacrifice (281673, -0.40 DPS) [quest]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.76 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 1.1 spell_power points (0.18 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.07 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 139.7 spell_power points (22.80 DPS) | yes | Skycaller (12984, -1.80 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.51 DPS) [dungeon]; Sizzle Stick (8071, -4.33 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 66.4. Weights run: 2.4s. Verify run: 1.3s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.191, intellect=0.468 ± 0.014, crit=0.102 ± 0.004 per rating point (14 rating = 1%, 1.429 per %), hit=0.230 ± 0.002 per rating point (10 rating = 1%, 2.299 per %), spell_haste=not significant (0.475 ± 0.198), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.567 ± 0.191), fire_power=0.430 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.27 DPS) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Embalmed Shroud (7691, -0.35 DPS) [dungeon]; Enchanter's Cowl (4322, -0.46 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.13 DPS) | yes | Crystal Starfire Medallion (5003, -0.91 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.91 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.61 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (1.52 DPS) | yes | Fairywing Mantle (9536, -0.35 DPS) [quest]; Death Speaker Mantle (6685, -0.38 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.45 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.58 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Battle Healer's Cloak (19529, -0.12 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.10 DPS) [dungeon]; Pristine Gown (253961, -0.31 DPS) [crafted]; Green Silk Armor (7065, -1.14 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Nightsky Wristbands (6407, -0.71 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.71 DPS) [quest]; Glowing Magical Bracelets (13106, -1.98 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.3 spell_power points (0.96 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.4 spell_power points (1.43 DPS) | yes | Warsong Sash (16975, +0.00 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.35 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.20 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.37 DPS) [crafted]; Abomination Skin Leggings (23173, -1.12 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.18 DPS) | yes | Acidic Walkers (9454, -0.18 DPS) [dungeon]; Boots of the Enchanter (4325, -0.61 DPS) [crafted]; Spidersilk Boots (4320, -1.51 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.81 DPS) | yes | Black Widow Band (6199, -0.43 DPS) [world]; Snake Hoop (6750, -0.43 DPS) [quest]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Snake Hoop (6750, -0.31 DPS) [quest]; Sludge-Stained Band (286535, -0.35 DPS) [world]; Black Widow Band (6199, -2.80 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Twisted Chanter's Staff (890, -0.50 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.50 DPS) [quest]; Glimmering Staff (249392, -1.38 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.8 spell_power points (1.13 DPS) | yes | Alliance Outrunner Healing Rod (285348, -0.53 DPS, sim-verified) [world]; Orb of Souls (249395, -0.67 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.68 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.1 spell_power points (33.56 DPS) | yes | Starfaller (13063, -0.36 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 131.8. Weights run: 2.1s. Verify run: 1.1s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.392, intellect=0.637 ± 0.023, crit=0.144 ± 0.005 per rating point (14 rating = 1%, 2.009 per %), hit=0.353 ± 0.003 per rating point (10 rating = 1%, 3.529 per %), spell_haste=not significant (1.212 ± 0.340), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.049 ± 0.392), fire_power=0.952 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.42 DPS) | yes | Living Cowl (5608, -1.30 DPS) [world]; Enchanter's Cowl (4322, -1.40 DPS) [crafted]; Augural Shroud (2620, -2.73 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.76 DPS) | yes | Triune Amulet (7722, -1.04 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.04 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.29 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.3 spell_power points (2.49 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.31 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.7 spell_power points (2.40 DPS) | yes | Guardian Cloak (5965, -0.90 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.26 DPS) [vendor]; Long Silken Cloak (4326, -2.41 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.8 spell_power points (4.20 DPS) | yes | Robe of Power (7054, -0.68 DPS) [crafted]; Elemental Raiment (9434, -0.79 DPS) [world_drop]; Dreamweave Vest (10021, -1.30 DPS, sim-verified) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.1 spell_power points (1.48 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.34 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (3.34 DPS) | yes | Black Mageweave Gloves (10003, -0.90 DPS) [crafted]; Gilded Handwraps (254021, -1.32 DPS) [crafted]; Red Mageweave Gloves (10018, -2.11 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.6 spell_power points (2.69 DPS) | yes | Defiler's Cloth Girdle (20166, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.56 DPS) [crafted]; Star Belt (4329, -0.58 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.6 spell_power points (3.52 DPS) | yes | Abomination Skin Leggings (23173, -1.23 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.55 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.40 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.90 DPS) | yes | Acidic Walkers (9454, -2.26 DPS) [dungeon]; Spidersilk Boots (4320, -2.35 DPS) [crafted]; Gilded Slippers (254001, -3.07 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (2.25 DPS) | yes | Reedknot Ring (9622, -1.11 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.27 DPS) [vendor]; Black Widow Band (6199, -1.52 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.46 DPS) | yes | Sea Giant's Toe Ring (274746, -0.49 DPS) [vendor]; Black Widow Band (6199, -0.74 DPS) [world]; Reedknot Ring (9622, -1.33 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (131.8 DPS) | yes | Staff of Dar'Orahil (15106, -1.54 DPS) [quest]; Windweaver Staff (7757, -1.70 DPS) [dungeon]; Gut Ripper (2164, -8.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 245.9 spell_power points (40.00 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.41 DPS) [dungeon]; Twisted Nether Wand (249144, -5.03 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 190.4. Weights run: 2.0s. Verify run: 1.3s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.197, intellect=0.442 ± 0.014, crit=0.083 ± 0.003 per rating point (14 rating = 1%, 1.165 per %), hit=0.193 ± 0.002 per rating point (10 rating = 1%, 1.927 per %), spell_haste=not significant (-0.249 ± 0.229), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.672 ± 0.197), fire_power=0.327 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.65 DPS) [crafted]; Red Mageweave Headband (10033, -2.28 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.7 spell_power points (3.99 DPS) | yes | Mindburst Medallion (11196, -0.41 DPS) [quest]; Horizon Choker (13085, -1.43 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.16 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Rotgrip Mantle (17732, -1.99 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -2.38 DPS) [crafted]; Red Mageweave Shoulders (10029, -2.53 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.7 spell_power points (6.88 DPS) | yes | Deep Woodlands Cloak (19121, -0.28 DPS) [quest]; Mantle of Lady Falther'ess (23178, -1.52 DPS) [dungeon]; Runecloth Cloak (13860, -1.70 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 27.8 spell_power points (11.50 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -2.42 DPS) [crafted]; Runecloth Tunic (13857, -2.47 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 10.1 spell_power points (4.17 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.82 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (8.16 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.20 DPS, sim-verified) [vendor]; Raider Handwraps (272098, -1.26 DPS) [vendor]; Runecloth Gloves (13863, -1.57 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Cloth Girdle (20166, -0.53 DPS) [rep]; Dawnspire Cord (12466, -1.10 DPS) [dungeon]; Satyrmane Sash (17755, -3.03 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.4 spell_power points (11.32 DPS) | yes | Wizardweave Leggings (14132, -3.48 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.70 DPS) [vendor]; Red Mageweave Pants (10009, -4.30 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.91 DPS) | yes | Gilded Sandals (254107, -3.73 DPS) [crafted]; Black Mageweave Boots (10026, -4.09 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.17 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.37 DPS) | yes | Cyclopean Band (11824, -0.37 DPS) [dungeon]; Advisor's Ring (19519, -0.41 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (5.23 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Advisor's Ring (19519, -0.27 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.27 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -4.40 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.16 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.97 DPS) [world_drop]; Blade of Eternal Darkness (17780, -3.62 DPS, sim-verified) [dungeon]; Arbiter's Blade (11784, -4.04 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.01 DPS) [crafted]; Wand of Allistarj (13065, -3.64 DPS) [world_drop]; Pyric Caduceus (11748, -7.69 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 386.8. Weights run: 2.2s. Verify run: 1.3s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.342, intellect=0.402 ± 0.023, crit=0.143 ± 0.004 per rating point (14 rating = 1%, 2.006 per %), hit=0.371 ± 0.003 per rating point (10 rating = 1%, 3.715 per %), spell_haste=not significant (-1.237 ± 0.374), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.625 ± 0.342), fire_power=0.375 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.2 spell_power points (15.36 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.37 DPS) [pvp]; Deathmist Mask (226909, -5.64 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.18 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.82 DPS) [quest]; Beads of Ogre Mojo (22149, -1.93 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.0 spell_power points (16.21 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.48 DPS) [pvp]; Burial Shawl (18681, -2.04 DPS, sim-verified) [dungeon]; Argent Shoulders (19059, -4.64 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.9 spell_power points (10.61 DPS) | yes | Crystalline Threaded Cape (20697, -0.61 DPS) [world]; Hide of the Wild (18510, -2.27 DPS) [crafted]; Amplifying Cloak (18350, -2.28 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.6 spell_power points (22.95 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.68 DPS) [pvp]; Robe of Everlasting Night (18385, -3.97 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -7.67 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.2 spell_power points (11.66 DPS) | yes | Sublime Wristguards (18497, -4.25 DPS) [dungeon]; Runecloth Cuffs (254123, -4.72 DPS) [crafted]; Deathmist Bracers (226907, -6.19 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.0 spell_power points (13.42 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -2.96 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.6 spell_power points (17.38 DPS) | yes | Belt of the Archmage (18405, -5.89 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -7.19 DPS) [rep]; Defiler's Cloth Girdle (20163, -8.86 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.0 spell_power points (18.52 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.30 DPS) [dungeon]; Outrider's Silk Leggings (22747, -2.03 DPS) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.8 spell_power points (11.48 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.18 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.38 DPS) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.59 DPS) [quest]; Maiden's Circle (13001, -2.59 DPS) [world_drop]; Naglering (11669, -10.74 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.39 DPS) [quest]; Maiden's Circle (13001, -1.39 DPS) [world_drop]; Naglering (11669, -10.35 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.78 DPS) [quest]; Weakness Analyzer (272438, -3.24 DPS) [vendor]; Talisman of Ascendance (22678, -8.16 DPS, sim-verified) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.49 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (386.8 DPS) | yes | Bonecreeper Stylus (13938, -0.48 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.62 DPS) [world]; Torch of Light (279246, -18.19 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

