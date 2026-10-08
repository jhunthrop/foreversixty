# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 34.4. Weights run: 1.7s. Verify run: 0.9s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.013 ± 0.002, crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.283 per %), hit=0.087 ± 0.003 per rating point (10 rating = 1%, 0.866 per %), spell_haste=1.678 ± 0.043, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.250 ± 0.000, fire_power=0.750 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.24 DPS) | yes | Shadow Goggles (4373, -2.46 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.1 spell_power points (1.06 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.01 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.83 DPS) | yes | Feyscale Cloak (6632, -0.21 DPS) [dungeon]; Caretaker's Cape (20428, -0.21 DPS) [rep]; Black Whelp Cloak (7283, -0.37 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.1 spell_power points (1.05 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -1.24 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.21 DPS) | yes | Mindthrust Bracers (1974, -0.19 DPS) [dungeon]; Bright Bracers (3647, -0.20 DPS) [world_drop]; Repurposed Hair Band (281256, -0.20 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.45 DPS) | yes | Gnoll Casting Gloves (892, -0.25 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.61 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.03 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.1 spell_power points (0.84 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.80 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.82 DPS) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.1 spell_power points (1.89 DPS) | yes | Silk-threaded Trousers (1929, -0.56 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.63 DPS) [crafted]; Rumpled Kilt (274741, -0.85 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.1 spell_power points (1.46 DPS) | yes | Feather Padded Treads (285345, -0.50 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.63 DPS) [crafted]; Pristine Boots (253889, -0.83 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.04 DPS) | yes | Sludge-Stained Band (286535, -0.42 DPS) [world]; Lavishly Jeweled Ring (1156, -1.03 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.03 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.04 DPS) | yes | Sludge-Stained Band (286535, -0.45 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -1.02 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.03 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.1 spell_power points (0.03 DPS) | yes | Channeler's Staff (4437, -0.01 DPS) [world]; Lesser Staff of the Spire (1300, -0.01 DPS) [world]; Staff of Westfall (2042, -0.01 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 110.7 spell_power points (22.93 DPS) | yes | Skycaller (12984, -1.57 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.64 DPS) [dungeon]; Sizzle Stick (8071, -4.25 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 58.7. Weights run: 1.8s. Verify run: 1.0s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.034 ± 0.003, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.375 per %), hit=0.116 ± 0.004 per rating point (10 rating = 1%, 1.165 per %), spell_haste=-1.395 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.240 ± 0.000, fire_power=0.760 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.45 DPS) | yes | Silk Headband (7050, -0.59 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.67 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.67 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.2 spell_power points (1.61 DPS) | yes | Crystal Starfire Medallion (5003, -1.58 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.58 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.63 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.3 spell_power points (2.07 DPS) | yes | Moonlit Amice (11884, -0.51 DPS) [quest]; Invoker's Mantle (215365, -0.51 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.65 DPS) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.11 DPS) | yes | Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest]; Caretaker's Cape (19533, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.90 DPS) | yes | Green Silk Armor (7065, -0.65 DPS, sim-verified) [crafted]; Robes of Arcana (5770, -1.11 DPS) [crafted]; Death Speaker Robes (6682, -1.26 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.01 DPS) | yes | Windsong Bangles (263336, -1.88 DPS, sim-verified) [quest]; Glowing Magical Bracelets (13106, -1.95 DPS) [world_drop]; Nightsky Wristbands (6407, -1.96 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.56 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.42 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.1 spell_power points (2.48 DPS) | yes | Belt of Arugal (6392, -0.30 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.69 DPS) [dungeon]; Invoker's Cord (215366, -0.88 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.68 DPS) | yes | Abomination Skin Leggings (23173, -0.74 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.06 DPS) [crafted]; Silk-threaded Trousers (1929, -1.11 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.2 spell_power points (1.61 DPS) | yes | Nimbus Boots (6998, -0.28 DPS) [quest]; Acidic Walkers (9454, -0.44 DPS) [dungeon]; Spidersilk Boots (4320, -1.63 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.56 DPS) | yes | Minor Channeling Ring (1449, -0.43 DPS) [quest]; Electrocutioner Lagnut (9447, -0.89 DPS) [dungeon]; Sludge-Stained Band (286535, -0.89 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.34 DPS) | yes | Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Sludge-Stained Band (286535, -0.67 DPS) [world]; Minor Channeling Ring (1449, -1.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.01 DPS) | yes | Glimmering Staff (249392, -1.66 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.93 DPS) [world_drop]; Channeler's Staff (4437, -1.95 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.2 spell_power points (1.61 DPS) | yes | Dwarven Tome (279898, -0.54 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.71 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.71 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 151.9 spell_power points (33.88 DPS) | yes | Starfaller (13063, -0.85 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.76 DPS) [crafted]; Gravestone Scepter (7001, -4.88 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 126.5. Weights run: 1.4s. Verify run: 0.9s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.053 ± 0.006, crit=0.043 ± 0.001 per rating point (14 rating = 1%, 0.604 per %), hit=0.186 ± 0.006 per rating point (10 rating = 1%, 1.861 per %), spell_haste=-6.904 ± 0.227, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.232 ± 0.000, fire_power=0.768 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.31 DPS) | yes | Living Cowl (5608, -2.04 DPS, sim-verified) [world]; Augural Shroud (2620, -2.40 DPS) [world]; Holy Shroud (2721, -2.53 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 spell_power points (1.85 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.43 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.76 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.76 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.40 DPS) | yes | Green Silken Shoulders (7057, -0.23 DPS) [crafted]; Inquisitor's Shawl (19507, -0.45 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 9.5 spell_power points (2.40 DPS) | yes | Long Silken Cloak (4326, -0.81 DPS) [crafted]; Guardian Cloak (5965, -0.81 DPS) [crafted]; Icy Cloak (4327, -1.10 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (5.65 DPS) | yes | Elemental Raiment (9434, -0.51 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.97 DPS) [crafted]; Robe of Power (7054, -1.94 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.28 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.51 DPS) [quest]; Earthen Silk Cuffs (254019, -1.27 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.61 DPS) | yes | Black Mageweave Gloves (10003, -1.28 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.69 DPS) [crafted]; Gilded Handwraps (254021, -2.49 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.2 spell_power points (3.60 DPS) | yes | Star Belt (4329, -0.31 DPS) [crafted]; Belt of Arugal (6392, -1.28 DPS) [dungeon]; Gilded Cord (254037, -1.46 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.6 spell_power points (3.71 DPS) | yes | Gaze Dreamer Pants (6903, -0.51 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.32 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.50 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.07 DPS) | yes | Gilded Slippers (254001, -2.37 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.25 DPS) [crafted]; Wingborne Boots (15104, -4.56 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.61 DPS) | yes | Ring of Forlorn Spirits (2043, -0.59 DPS) [quest]; Reedknot Ring (9622, -0.84 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.09 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.28 DPS) | yes | Ring of Forlorn Spirits (2043, -0.37 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.51 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.76 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (126.5 DPS) | yes | Scorn's Focal Dagger (23168, -2.78 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.44 DPS) [quest]; Gut Ripper (2164, -6.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 155.9 spell_power points (39.46 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -3.87 DPS) [dungeon]; Twisted Nether Wand (249144, -3.94 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 192.5. Weights run: 3.8s. Verify run: 1.1s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.029 ± 0.018, crit=0.098 ± 0.004 per rating point (14 rating = 1%, 1.373 per %), hit=0.445 ± 0.023 per rating point (10 rating = 1%, 4.451 per %), spell_haste=-12.811 ± 0.550, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.226 ± 0.000, fire_power=0.774 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.26 DPS) | yes | Dreamweave Circlet (10041, +0.00 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.39 DPS) [crafted]; Red Mageweave Headband (10033, -1.72 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.2 spell_power points (1.66 DPS) | yes | Mindburst Medallion (11196, -0.23 DPS) [quest]; Horizon Choker (13085, -1.57 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.60 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.4 spell_power points (3.33 DPS) | yes | Black Mageweave Shoulders (10027, -0.95 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.10 DPS) [vendor]; Rotgrip Mantle (17732, -2.80 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.2 spell_power points (3.29 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.98 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.15 DPS) [crafted]; Nightfall Drape (12465, -1.20 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.2 spell_power points (5.14 DPS) | yes | Elemental Raiment (9434, -0.27 DPS) [world_drop]; Acumen Robes (17775, -0.60 DPS) [quest]; Dreamweave Vest (10021, -0.91 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.09 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.42 DPS) [crafted]; Condor Bracers (15864, -0.46 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.1 spell_power points (4.20 DPS) | yes | Black Mageweave Gloves (10003, -0.72 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.13 DPS) [vendor]; Brightcloth Gloves (14101, -1.19 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.8 spell_power points (3.89 DPS) | yes | Highlander's Cloth Girdle (20098, -0.62 DPS) [rep]; Ghostweave Cord (254073, -0.64 DPS) [crafted]; Satyrmane Sash (17755, -2.67 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 23.3 spell_power points (5.40 DPS) | yes | Red Mageweave Pants (10009, -2.07 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.22 DPS) [vendor]; Wizardweave Leggings (14132, -4.24 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.57 DPS) | yes | Gilded Sandals (254107, -2.95 DPS) [crafted]; Black Mageweave Boots (10026, -2.97 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.33 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.01 DPS) | yes | Philanthropist's Ring (281635, -0.65 DPS) [quest]; Cyclopean Band (11824, -0.88 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.16 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (2.78 DPS) | yes | Philanthropist's Ring (281635, -0.42 DPS) [quest]; Cyclopean Band (11824, -0.65 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -0.93 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.50 DPS, sim-verified) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -2.55 DPS) [dungeon]; Arbiter's Blade (11784, -2.75 DPS) [dungeon]; Blade of Eternal Darkness (17780, -3.36 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (192.5 DPS) | yes | Wand of Allistarj (13065, -2.74 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.55 DPS) [crafted]; Pyric Caduceus (11748, -7.24 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532000000000000-2332113101220001350-0040000000000000)

Set DPS (verified): 418.2. Weights run: 1.5s. Verify run: 2.6s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.115 ± 0.028, crit=0.127 ± 0.004 per rating point (14 rating = 1%, 1.784 per %), hit=0.603 ± 0.031 per rating point (10 rating = 1%, 6.033 per %), spell_haste=-12.127 ± 0.610, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.309 ± 0.001, fire_power=0.691 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.9 spell_power points (11.22 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -1.50 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -5.45 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.98 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.20 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.5 spell_power points (11.07 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.29 DPS) [pvp]; Argent Shoulders (19059, -1.67 DPS, sim-verified) [crafted]; Burial Shawl (18681, -3.15 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.9 spell_power points (8.33 DPS) | yes | Crystalline Threaded Cape (20697, -0.90 DPS) [world]; Amplifying Cloak (18350, -1.80 DPS) [dungeon]; Hide of the Wild (18510, -2.83 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (17.07 DPS) | yes | Robe of Everlasting Night (18385, -4.08 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -4.46 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.16 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.9 spell_power points (8.32 DPS) | yes | Sublime Wristguards (18497, -3.55 DPS) [dungeon]; Runecloth Cuffs (254123, -3.91 DPS) [crafted]; Arcane Runed Bracers (4744, -5.05 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (10.01 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.02 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.6 spell_power points (12.18 DPS) | yes | Ban'thok Sash (11662, -5.17 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -5.23 DPS) [rep]; Belt of the Archmage (18405, -6.26 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 36.7 spell_power points (13.32 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.65 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -2.62 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.71 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.73 DPS) [crafted]; Omnicast Boots (11822, -0.95 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.53 DPS) [quest]; Maiden's Circle (13001, -1.62 DPS) [world_drop]; Naglering (11669, -9.16 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.08 DPS) [quest]; Maiden's Circle (13001, -1.17 DPS) [world_drop]; Naglering (11669, -13.14 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+20.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.54 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.06 DPS, sim-verified) [quest]; Serenity Field (272439, -5.44 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.78 DPS) [world]; Teebu's Blazing Longsword (1728, -20.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (418.2 DPS) | yes | Bonecreeper Stylus (13938, -1.05 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.32 DPS) [world]; Torch of Light (279246, -8.33 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 05500000000000000-2035113111120001350-0050051000000000)

Set DPS (verified): 837.8. Weights run: 5.2s. Verify run: 2.8s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.087 ± 0.023, crit=0.298 ± 0.009 per rating point (14 rating = 1%, 4.175 per %), hit=0.738 ± 0.047 per rating point (10 rating = 1%, 7.383 per %), spell_haste=-24.073 ± 0.955, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.226 ± 0.001, fire_power=0.774 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.7 spell_power points (16.92 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -1.78 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -2.04 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (12.12 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.71 DPS) [dungeon]; Amulet of the Dawn (22657, -3.24 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 32.5 spell_power points (17.90 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -3.31 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.61 DPS) [crafted]; Argent Shoulders (19059, -9.89 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.1 spell_power points (13.27 DPS) | yes | Amplifying Cloak (18350, -3.35 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.68 DPS, sim-verified) [world]; Hide of the Wild (18510, -5.08 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.8 spell_power points (25.78 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.00 DPS) [pvp]; Robe of Everlasting Night (18385, -7.41 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -11.05 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.7 spell_power points (12.51 DPS) | yes | Sublime Wristguards (18497, -5.42 DPS) [dungeon]; Runecloth Cuffs (254123, -5.97 DPS) [crafted]; Arcane Runed Bracers (4744, -7.55 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (15.12 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.99 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 34.3 spell_power points (18.90 DPS) | yes | Ban'thok Sash (11662, -7.69 DPS) [dungeon]; Belt of the Archmage (18405, -8.22 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -8.50 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 41.2 spell_power points (22.72 DPS) | yes | Marshal's Dreadweave Leggings (231587, -1.42 DPS) [pvp]; Skyshroud Leggings (13170, -3.60 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -6.66 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.23 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.10 DPS) [crafted]; Omnicast Boots (11822, -1.63 DPS) [dungeon] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.98 DPS) [quest]; Maiden's Circle (13001, -2.63 DPS) [world_drop]; Naglering (11669, -20.98 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.74 DPS) [quest]; Maiden's Circle (13001, -2.40 DPS) [world_drop]; Naglering (11669, -15.81 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+25.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.86 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -4.71 DPS, sim-verified) [quest]; Serenity Field (272439, -8.27 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.91 DPS) [world]; Teebu's Blazing Longsword (1728, -32.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (837.8 DPS) | yes | Bonecreeper Stylus (13938, -1.03 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.40 DPS) [world]; Torch of Light (279246, -9.84 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 33.5. Weights run: 1.7s. Verify run: 1.0s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.013 ± 0.002, crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.283 per %), hit=0.087 ± 0.003 per rating point (10 rating = 1%, 0.866 per %), spell_haste=1.678 ± 0.043, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.250 ± 0.000, fire_power=0.750 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.24 DPS) | yes | Shadow Goggles (4373, -2.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.1 spell_power points (1.06 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.13 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.83 DPS) | yes | Feyscale Cloak (6632, -0.21 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.21 DPS) [rep]; Black Whelp Cloak (7283, -0.26 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.1 spell_power points (1.05 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -1.15 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.0 spell_power points (0.21 DPS) | yes | Windsong Bangles (263336, -0.01 DPS) [quest]; Tabitha's Cuffs (251486, -0.20 DPS) [quest]; Featherbead Bracers (15452, -0.20 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.45 DPS) | yes | Gnoll Casting Gloves (892, -0.25 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.61 DPS) [crafted]; Apothecary Gloves (10919, -0.62 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.1 spell_power points (0.84 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.67 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.82 DPS) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.1 spell_power points (1.89 DPS) | yes | Silk-threaded Trousers (1929, -0.48 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.63 DPS) [crafted]; Rumpled Kilt (274741, -0.85 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.1 spell_power points (1.46 DPS) | yes | Feather Padded Treads (285345, -0.46 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.63 DPS) [crafted]; Pristine Boots (253889, -0.83 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.04 DPS) | yes | Lavishly Jeweled Ring (1156, -1.02 DPS) [dungeon]; Loop of Sacrifice (281673, -1.02 DPS) [quest]; Volcanic Rock Ring (12053, -1.03 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.62 DPS) | yes | Lavishly Jeweled Ring (1156, -0.45 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.61 DPS) [quest]; Volcanic Rock Ring (12053, -0.61 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.1 spell_power points (0.03 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.01 DPS) [world]; Lesser Staff of the Spire (1300, -0.01 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 110.7 spell_power points (22.93 DPS) | yes | Skycaller (12984, -1.41 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.64 DPS) [dungeon]; Sizzle Stick (8071, -4.25 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 57.8. Weights run: 1.8s. Verify run: 1.0s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.034 ± 0.003, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.375 per %), hit=0.116 ± 0.004 per rating point (10 rating = 1%, 1.165 per %), spell_haste=-1.395 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.240 ± 0.000, fire_power=0.760 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.45 DPS) | yes | Embalmed Shroud (7691, -0.67 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.67 DPS) [crafted]; Silk Headband (7050, -0.75 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.2 spell_power points (1.61 DPS) | yes | Crystal Starfire Medallion (5003, -1.58 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.58 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.74 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.3 spell_power points (2.07 DPS) | yes | Chestnut Mantle (17695, -0.41 DPS, sim-verified) [quest]; Invoker's Mantle (215365, -0.48 DPS) [crafted]; Death Speaker Mantle (6685, -0.65 DPS) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.11 DPS) | yes | Windsong Drape (15468, -0.20 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Battle Healer's Cloak (19529, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.90 DPS) | yes | Green Silk Armor (7065, -0.67 DPS, sim-verified) [crafted]; High Robe of the Adjudicator (3461, -1.10 DPS) [quest]; Robes of Arcana (5770, -1.11 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.01 DPS) | yes | Windsong Bangles (263336, -1.78 DPS) [quest]; Owlbeard Bracers (16981, -1.87 DPS, sim-verified) [quest]; Glowing Magical Bracelets (13106, -1.95 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.56 DPS) | yes | Jutebraid Gloves (10654, -0.20 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.42 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.1 spell_power points (2.48 DPS) | yes | Warsong Sash (16975, -0.02 DPS) [quest]; Belt of Arugal (6392, -0.45 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.69 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.68 DPS) | yes | Abomination Skin Leggings (23173, -0.52 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.06 DPS) [crafted]; Silk-threaded Trousers (1929, -1.11 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.2 spell_power points (1.61 DPS) | yes | Acidic Walkers (9454, -0.44 DPS) [dungeon]; Boots of the Enchanter (4325, -0.50 DPS) [crafted]; Spidersilk Boots (4320, -1.37 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.56 DPS) | yes | Electrocutioner Lagnut (9447, -0.89 DPS) [dungeon]; Sludge-Stained Band (286535, -0.89 DPS) [world]; Sacred Band (6669, -1.11 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.34 DPS) | yes | Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Sacred Band (6669, -0.89 DPS) [quest]; Sludge-Stained Band (286535, -2.17 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.01 DPS) | yes | Glimmering Staff (249392, -1.52 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.93 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.93 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.2 spell_power points (1.61 DPS) | yes | Orb of Souls (249395, -0.71 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -0.78 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -1.13 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 151.9 spell_power points (33.88 DPS) | yes | Starfaller (13063, -0.84 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.76 DPS) [crafted]; Gravestone Scepter (7001, -4.88 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 125.3. Weights run: 1.4s. Verify run: 0.9s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.053 ± 0.006, crit=0.043 ± 0.001 per rating point (14 rating = 1%, 0.604 per %), hit=0.186 ± 0.006 per rating point (10 rating = 1%, 1.861 per %), spell_haste=-6.904 ± 0.227, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.232 ± 0.000, fire_power=0.768 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.31 DPS) | yes | Living Cowl (5608, -2.00 DPS, sim-verified) [world]; Augural Shroud (2620, -2.40 DPS) [world]; Holy Shroud (2721, -2.53 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 spell_power points (1.85 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.53 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.76 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.76 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.40 DPS) | yes | Green Silken Shoulders (7057, -0.23 DPS) [crafted]; Chestnut Mantle (17695, -0.37 DPS) [quest]; Inquisitor's Shawl (19507, -0.45 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 9.5 spell_power points (2.40 DPS) | yes | Long Silken Cloak (4326, -0.81 DPS) [crafted]; Guardian Cloak (5965, -0.81 DPS) [crafted]; Icy Cloak (4327, -1.42 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (5.65 DPS) | yes | Elemental Raiment (9434, -0.46 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.97 DPS) [crafted]; Robe of Power (7054, -1.94 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.28 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.66 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.16 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.61 DPS) | yes | Black Mageweave Gloves (10003, -1.22 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.69 DPS) [crafted]; Gilded Handwraps (254021, -2.49 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.2 spell_power points (3.60 DPS) | yes | Star Belt (4329, -0.31 DPS) [crafted]; Warsong Sash (16975, -0.81 DPS) [quest]; Belt of Arugal (6392, -1.28 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.6 spell_power points (3.71 DPS) | yes | Gaze Dreamer Pants (6903, -0.88 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.32 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.50 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.07 DPS) | yes | Gilded Slippers (254001, -2.68 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.25 DPS) [crafted]; Acidic Walkers (9454, -4.70 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.61 DPS) | yes | Reedknot Ring (9622, -0.84 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.09 DPS) [vendor]; Sludge-Stained Band (286535, -1.85 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.28 DPS) | yes | Reedknot Ring (9622, -0.66 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.76 DPS) [vendor]; Sludge-Stained Band (286535, -1.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (125.3 DPS) | yes | Scorn's Focal Dagger (23168, -2.78 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.44 DPS) [quest]; Gut Ripper (2164, -6.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 155.9 spell_power points (39.46 DPS) | yes | Umbral Wand (5216, -3.79 DPS) [dungeon]; Earthen Rod (9381, -3.87 DPS) [dungeon]; Twisted Nether Wand (249144, -3.94 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 192.0. Weights run: 3.8s. Verify run: 1.1s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.029 ± 0.018, crit=0.098 ± 0.004 per rating point (14 rating = 1%, 1.373 per %), hit=0.445 ± 0.023 per rating point (10 rating = 1%, 4.451 per %), spell_haste=-12.811 ± 0.550, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.226 ± 0.000, fire_power=0.774 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.26 DPS) | yes | Dreamweave Circlet (10041, -1.32 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.39 DPS) [crafted]; Red Mageweave Headband (10033, -1.72 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.2 spell_power points (1.66 DPS) | yes | Mindburst Medallion (11196, -0.23 DPS) [quest]; Horizon Choker (13085, -1.57 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.60 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.4 spell_power points (3.33 DPS) | yes | Rotgrip Mantle (17732, -0.20 DPS) [dungeon]; Black Mageweave Shoulders (10027, -0.95 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.10 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.2 spell_power points (3.29 DPS) | yes | Deep Woodlands Cloak (19121, +0.00 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.14 DPS) [dungeon]; Runecloth Cloak (13860, -1.15 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.2 spell_power points (5.14 DPS) | yes | Acumen Robes (17775, -0.60 DPS) [quest]; Dreamweave Vest (10021, -0.91 DPS) [crafted]; Elemental Raiment (9434, -2.58 DPS, sim-verified) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -2.45 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.1 spell_power points (4.20 DPS) | yes | Black Mageweave Gloves (10003, -0.72 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.13 DPS) [vendor]; Brightcloth Gloves (14101, -1.19 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.8 spell_power points (3.89 DPS) | yes | Defiler's Cloth Girdle (20166, -0.62 DPS) [rep]; Ghostweave Cord (254073, -0.64 DPS) [crafted]; Satyrmane Sash (17755, -1.69 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 23.3 spell_power points (5.40 DPS) | yes | Red Mageweave Pants (10009, -2.07 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.22 DPS) [vendor]; Wizardweave Leggings (14132, -4.45 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.57 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.39 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -2.95 DPS) [crafted]; Black Mageweave Boots (10026, -2.97 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.01 DPS) | yes | Philanthropist's Ring (281635, -0.65 DPS) [quest]; Cyclopean Band (11824, -0.88 DPS) [dungeon]; Runed Ring (862, -1.39 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (2.78 DPS) | yes | Philanthropist's Ring (281635, -0.42 DPS) [quest]; Cyclopean Band (11824, -0.65 DPS) [dungeon]; Runed Ring (862, -1.16 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.48 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -2.06 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.21 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, -2.43 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -2.55 DPS) [dungeon]; Arbiter's Blade (11784, -2.75 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.74 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.55 DPS) [crafted]; Pyric Caduceus (11748, -7.60 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532000000000000-2332113101220001350-0040000000000000)

Set DPS (verified): 413.0. Weights run: 1.5s. Verify run: 2.5s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.115 ± 0.028, crit=0.127 ± 0.004 per rating point (14 rating = 1%, 1.784 per %), hit=0.603 ± 0.031 per rating point (10 rating = 1%, 6.033 per %), spell_haste=-12.127 ± 0.610, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.309 ± 0.001, fire_power=0.691 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.9 spell_power points (11.22 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -1.50 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -2.66 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.98 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.20 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.5 spell_power points (11.07 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.29 DPS) [pvp]; Argent Shoulders (19059, -2.00 DPS) [crafted]; Burial Shawl (18681, -3.15 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.9 spell_power points (8.33 DPS) | yes | Crystalline Threaded Cape (20697, -0.90 DPS) [world]; Amplifying Cloak (18350, -1.80 DPS) [dungeon]; Hide of the Wild (18510, -2.83 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (17.07 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.46 DPS) [pvp]; Robe of Everlasting Night (18385, -6.73 DPS) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -7.16 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.9 spell_power points (8.32 DPS) | yes | Sublime Wristguards (18497, -3.55 DPS) [dungeon]; Runecloth Cuffs (254123, -3.91 DPS) [crafted]; Spidertank Oilrag (9448, -5.05 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (10.01 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.02 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.6 spell_power points (12.18 DPS) | yes | Belt of the Archmage (18405, -3.96 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -5.17 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -5.23 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 36.7 spell_power points (13.32 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.37 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.71 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.73 DPS) [crafted]; Omnicast Boots (11822, -0.95 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.53 DPS) [quest]; Maiden's Circle (13001, -1.62 DPS) [world_drop]; Naglering (11669, -5.75 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.08 DPS) [quest]; Maiden's Circle (13001, -1.17 DPS) [world_drop]; Naglering (11669, -10.77 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+17.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.54 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.99 DPS, sim-verified) [quest]; Serenity Field (272439, -5.44 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.78 DPS) [world]; Teebu's Blazing Longsword (1728, -17.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (413.0 DPS) | yes | Bonecreeper Stylus (13938, -1.05 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.32 DPS) [world]; Torch of Light (279246, -13.07 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 05500000000000000-2035113111120001350-0050051000000000)

Set DPS (verified): 828.4. Weights run: 5.2s. Verify run: 2.8s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.087 ± 0.023, crit=0.298 ± 0.009 per rating point (14 rating = 1%, 4.175 per %), hit=0.738 ± 0.047 per rating point (10 rating = 1%, 7.383 per %), spell_haste=-24.073 ± 0.955, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.226 ± 0.001, fire_power=0.774 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.7 spell_power points (16.92 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -2.04 DPS) [crafted]; Deathmist Mask (226909, -3.59 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (12.12 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.71 DPS) [dungeon]; Amulet of the Dawn (22657, -3.24 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 32.5 spell_power points (17.90 DPS) | yes | Warlord's Dreadweave Mantle (231592, -3.31 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.61 DPS) [crafted]; Argent Shoulders (19059, -6.00 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.1 spell_power points (13.27 DPS) | yes | Crystalline Threaded Cape (20697, -2.06 DPS) [world]; Amplifying Cloak (18350, -3.35 DPS) [dungeon]; Hide of the Wild (18510, -5.08 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.8 spell_power points (25.78 DPS) | yes | Robe of Everlasting Night (18385, -5.12 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -7.00 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -11.05 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.7 spell_power points (12.51 DPS) | yes | Sublime Wristguards (18497, -5.42 DPS) [dungeon]; Runecloth Cuffs (254123, -5.97 DPS) [crafted]; Spidertank Oilrag (9448, -7.55 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (15.12 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.99 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 34.3 spell_power points (18.90 DPS) | yes | Belt of the Archmage (18405, -7.52 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -7.69 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -8.50 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 41.2 spell_power points (22.72 DPS) | yes | General's Dreadweave Pants (231588, -1.42 DPS) [pvp]; Skyshroud Leggings (13170, -3.60 DPS) [dungeon]; Outrider's Silk Leggings (22747, -6.38 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.23 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.10 DPS) [crafted]; Omnicast Boots (11822, -1.63 DPS) [dungeon] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.98 DPS) [quest]; Maiden's Circle (13001, -2.63 DPS) [world_drop]; Naglering (11669, -19.31 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.74 DPS) [quest]; Maiden's Circle (13001, -2.40 DPS) [world_drop]; Naglering (11669, -14.87 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+24.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.86 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -4.65 DPS, sim-verified) [quest]; Serenity Field (272439, -8.27 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.91 DPS) [world]; Teebu's Blazing Longsword (1728, -31.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (828.4 DPS) | yes | Bonecreeper Stylus (13938, -1.03 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.40 DPS) [world]; Torch of Light (279246, -9.58 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

