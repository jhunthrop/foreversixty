# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 43.1. Weights run: 2.3s. Verify run: 1.4s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.061, intellect=-0.053 ± 0.004, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.379 per %), hit=0.099 ± 0.001 per rating point (10 rating = 1%, 0.987 per %), spell_haste=not significant (0.088 ± 0.060), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.796 ± 0.061

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.18 DPS) | yes | Red Winter Hat (21524, -3.06 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.98 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.20 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.79 DPS) | yes | Feyscale Cloak (6632, -0.20 DPS) [dungeon]; Caretaker's Cape (20428, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.29 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.98 DPS) | yes | Green Woolen Vest (2582, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.20 DPS) [dungeon]; Gray Woolen Robe (2585, -1.39 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Ivycloth Bracelets (9793, -0.28 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.38 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.59 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.98 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (43.1 DPS) | yes | Novice Ardent's Sash (253887, -0.39 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.87 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.77 DPS) | yes | Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted]; Silk-threaded Trousers (1929, -0.77 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.79 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.38 DPS) | yes | Red Woolen Boots (4313, -0.59 DPS) [crafted]; Feather Padded Treads (285345, -0.69 DPS, sim-verified) [world]; Pristine Boots (253889, -0.79 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.98 DPS) | yes | Sludge-Stained Band (286535, -0.39 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.98 DPS) | yes | Sludge-Stained Band (286535, -0.58 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.28 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.98 DPS) | yes | Bouquet of Red Roses (22206, -1.26 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 116.5 spell_power points (22.90 DPS) | yes | Skycaller (12984, -2.03 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.61 DPS) [dungeon]; Sizzle Stick (8071, -4.27 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 71.4. Weights run: 2.3s. Verify run: 1.3s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.085, intellect=0.101 ± 0.005, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.419 per %), hit=0.109 ± 0.001 per rating point (10 rating = 1%, 1.089 per %), spell_haste=0.930 ± 0.082, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.807 ± 0.085

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.38 DPS) | yes | Embalmed Shroud (7691, -0.65 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.65 DPS) [crafted]; Silk Headband (7050, -0.69 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (1.64 DPS) | yes | Crystal Starfire Medallion (5003, -1.56 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.06 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (2.14 DPS) | yes | Invoker's Mantle (215365, -0.43 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.60 DPS) [dungeon]; Moonlit Amice (11884, -0.63 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.08 DPS) | yes | Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest]; Caretaker's Cape (19533, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.81 DPS) | yes | Death Speaker Robes (6682, -1.06 DPS) [dungeon]; Robes of Arcana (5770, -1.08 DPS) [crafted]; Green Silk Armor (7065, -1.16 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.94 DPS) | yes | Glowing Magical Bracelets (13106, -1.77 DPS) [world_drop]; Nightsky Wristbands (6407, -1.81 DPS) [world_drop]; Windsong Bangles (263336, -2.32 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.51 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.37 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.3 spell_power points (2.44 DPS) | yes | Belt of Arugal (6392, -0.67 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.71 DPS) [dungeon]; Invoker's Cord (215366, -0.82 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.59 DPS) | yes | Abomination Skin Leggings (23173, -0.77 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.93 DPS) [crafted]; Silk-threaded Trousers (1929, -1.08 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.7 spell_power points (1.67 DPS) | yes | Nimbus Boots (6998, -0.37 DPS) [quest]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.89 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.51 DPS) | yes | Minor Channeling Ring (1449, -0.39 DPS) [quest]; Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Sludge-Stained Band (286535, -0.86 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.30 DPS) | yes | Electrocutioner Lagnut (9447, -0.65 DPS) [dungeon]; Sludge-Stained Band (286535, -0.65 DPS) [world]; Minor Channeling Ring (1449, -1.85 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.94 DPS) | yes | Twisted Chanter's Staff (890, -1.73 DPS) [world_drop]; Channeler's Staff (4437, -1.77 DPS) [world]; Glimmering Staff (249392, -2.07 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.6 spell_power points (1.64 DPS) | yes | Dwarven Tome (279898, -0.46 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.78 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.78 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 156.7 spell_power points (33.86 DPS) | yes | Starfaller (13063, -1.07 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.78 DPS) [crafted]; Gravestone Scepter (7001, -4.86 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 129.0. Weights run: 1.9s. Verify run: 1.1s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.111, intellect=0.181 ± 0.007, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.491 per %), hit=0.134 ± 0.001 per rating point (10 rating = 1%, 1.342 per %), spell_haste=not significant (0.359 ± 0.114), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.830 ± 0.111

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.15 DPS) | yes | Augural Shroud (2620, -2.01 DPS) [world]; Living Cowl (5608, -2.41 DPS, sim-verified) [world]; Holy Shroud (2721, -2.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.1 spell_power points (1.98 DPS) | yes | Triune Amulet (7722, -1.67 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.67 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.17 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.6 spell_power points (2.61 DPS) | yes | Inquisitor's Shawl (19507, -0.31 DPS) [dungeon]; Green Silken Shoulders (7057, -0.39 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 10.6 spell_power points (2.61 DPS) | yes | Long Silken Cloak (4326, -0.91 DPS) [crafted]; Guardian Cloak (5965, -0.91 DPS) [crafted]; Icy Cloak (4327, -1.85 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (5.66 DPS) | yes | Elemental Raiment (9434, -0.42 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.85 DPS) [crafted]; Robe of Power (7054, -1.70 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.21 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.49 DPS) [quest]; Earthen Silk Cuffs (254019, -1.23 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (4.59 DPS) | yes | Black Mageweave Gloves (10003, -1.17 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.45 DPS) [crafted]; Gilded Handwraps (254021, -2.32 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.7 spell_power points (3.61 DPS) | yes | Star Belt (4329, -0.54 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.23 DPS) [dungeon]; Belt of Arugal (6392, -1.27 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.2 spell_power points (3.97 DPS) | yes | Gaze Dreamer Pants (6903, -1.02 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.40 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.43 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.89 DPS) | yes | Gilded Slippers (254001, -3.55 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.99 DPS) [crafted]; Acidic Walkers (9454, -4.31 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.1 spell_power points (2.72 DPS) | yes | Ring of Forlorn Spirits (2043, -0.76 DPS) [quest]; Reedknot Ring (9622, -1.00 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.25 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.21 DPS) | yes | Ring of Forlorn Spirits (2043, -0.25 DPS) [quest]; Reedknot Ring (9622, -0.49 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.74 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (129.0 DPS) | yes | Scorn's Focal Dagger (23168, -2.70 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.09 DPS) [quest]; Gut Ripper (2164, -7.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 161.6 spell_power points (39.65 DPS) | yes | Umbral Wand (5216, -3.98 DPS) [dungeon]; Earthen Rod (9381, -4.06 DPS) [dungeon]; Twisted Nether Wand (249144, -4.18 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 225.3. Weights run: 2.1s. Verify run: 1.5s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.138, intellect=0.159 ± 0.012, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.426 per %), hit=0.200 ± 0.002 per rating point (10 rating = 1%, 2.004 per %), spell_haste=not significant (-0.577 ± 0.169), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.899 ± 0.138

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.90 DPS) | yes | Red Mageweave Headband (10033, -1.59 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.98 DPS) [crafted]; Dreamweave Circlet (10041, -2.48 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (2.62 DPS) | yes | Mindburst Medallion (11196, -0.33 DPS) [quest]; Horizon Choker (13085, -1.89 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.10 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.1 spell_power points (5.29 DPS) | yes | Black Mageweave Shoulders (10027, -1.53 DPS) [crafted]; Bloodmage Mantle (7684, -1.86 DPS) [dungeon]; Rotgrip Mantle (17732, -1.96 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (4.93 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.11 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.54 DPS) [crafted]; Nightfall Drape (12465, -1.96 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (7.56 DPS) | yes | Elemental Raiment (9434, -0.64 DPS) [world_drop]; Acumen Robes (17775, -0.84 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -1.16 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.97 DPS) | yes | Nethergeld Cuffs (254061, -0.29 DPS) [crafted]; Spidertank Oilrag (9448, -0.61 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.66 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (6.14 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.39 DPS) [vendor]; Runecloth Gloves (13863, -1.72 DPS) [crafted]; Black Mageweave Gloves (10003, -1.90 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (5.14 DPS) | yes | Ghostweave Cord (254073, -0.52 DPS) [crafted]; Ban'thok Sash (11662, -0.54 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.03 DPS, sim-verified) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 24.6 spell_power points (8.10 DPS) | yes | Red Mageweave Pants (10009, -2.86 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.38 DPS) [vendor]; Wizardweave Leggings (14132, -5.03 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.91 DPS) | yes | Gilded Sandals (254107, -1.47 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.92 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.14 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.28 DPS) | yes | Philanthropist's Ring (281635, -0.67 DPS) [quest]; Cyclopean Band (11824, -0.95 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.65 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.95 DPS) | yes | Cyclopean Band (11824, -0.62 DPS) [dungeon]; Philanthropist's Ring (281635, -0.71 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.32 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.98 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.58 DPS, sim-verified) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.62 DPS) [dungeon]; Arbiter's Blade (11784, -3.69 DPS) [dungeon]; Blade of Eternal Darkness (17780, -8.28 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (225.3 DPS) | yes | Wand of Allistarj (13065, -3.23 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.26 DPS) [crafted]; Pyric Caduceus (11748, -6.75 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 399.4. Weights run: 2.2s. Verify run: 1.5s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.161, intellect=0.119 ± 0.015, crit=0.076 ± 0.002 per rating point (14 rating = 1%, 1.063 per %), hit=0.322 ± 0.003 per rating point (10 rating = 1%, 3.218 per %), spell_haste=not significant (-0.510 ± 0.220), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.892 ± 0.161

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.0 spell_power points (11.09 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -2.42 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.95 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.88 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.95 DPS) [quest]; Kezan's Taint (19604, -2.52 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.8 spell_power points (10.69 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.01 DPS) [pvp]; Argent Shoulders (19059, -1.54 DPS, sim-verified) [crafted]; Burial Shawl (18681, -2.85 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.5 spell_power points (7.33 DPS) | yes | Amplifying Cloak (18350, -0.89 DPS) [dungeon]; Arcanoweave Cloak (272411, -1.35 DPS, sim-verified) [vendor]; Hide of the Wild (18510, -1.89 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (16.86 DPS) | yes | Robe of Everlasting Night (18385, -3.49 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -4.38 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.05 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.0 spell_power points (8.22 DPS) | yes | Sublime Wristguards (18497, -3.50 DPS) [dungeon]; Runecloth Cuffs (254123, -3.85 DPS) [crafted]; Arcane Runed Bracers (4744, -5.00 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (9.88 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.00 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.8 spell_power points (11.05 DPS) | yes | Belt of the Archmage (18405, -3.69 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -4.17 DPS) [rep]; Oddly Magical Belt (18475, -5.31 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -1.68 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.60 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.92 DPS) [dungeon]; Venomspew Footpads (275606, -0.96 DPS, sim-verified) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.60 DPS) [quest]; Maiden's Circle (13001, -1.60 DPS) [world_drop]; Naglering (11669, -7.52 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.07 DPS) [quest]; Maiden's Circle (13001, -1.07 DPS) [world_drop]; Naglering (11669, -6.82 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Ascendance (22678, -1.26 DPS, sim-verified) [quest]; Royal Seal of Eldre'Thalas (18467, -2.15 DPS) [quest]; Weakness Analyzer (272438, -2.51 DPS) [vendor] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS, sim-verified) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.27 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -16.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+23.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.05 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.29 DPS) [world]; Torch of Light (279246, -23.14 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 40.7. Weights run: 2.3s. Verify run: 1.4s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.061, intellect=-0.053 ± 0.004, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.379 per %), hit=0.099 ± 0.001 per rating point (10 rating = 1%, 0.987 per %), spell_haste=not significant (0.088 ± 0.060), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.796 ± 0.061

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.18 DPS) | yes | Red Winter Hat (21524, -2.89 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.98 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.20 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.79 DPS) | yes | Feyscale Cloak (6632, -0.20 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.31 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.98 DPS) | yes | Green Woolen Vest (2582, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.20 DPS) [dungeon]; Gray Woolen Robe (2585, -1.21 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) (or Owlbeard Bracers (16981)) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.38 DPS) | yes | Gnoll Casting Gloves (892, -0.31 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.59 DPS) [quest]; Pristine Gloves (253913, -0.59 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (40.7 DPS) | yes | Novice Ardent's Sash (253887, -0.39 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.09 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.77 DPS) | yes | Silk-threaded Trousers (1929, -0.55 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted]; Rumpled Kilt (274741, -0.79 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.38 DPS) | yes | Feather Padded Treads (285345, -0.22 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.59 DPS) [crafted]; Pristine Boots (253889, -0.79 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.98 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.59 DPS) | yes | Ring of the Shadow (1462, -0.74 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 8 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 116.5 spell_power points (22.90 DPS) | yes | Skycaller (12984, -1.79 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.61 DPS) [dungeon]; Sizzle Stick (8071, -4.27 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 70.5. Weights run: 2.3s. Verify run: 1.3s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.085, intellect=0.101 ± 0.005, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.419 per %), hit=0.109 ± 0.001 per rating point (10 rating = 1%, 1.089 per %), spell_haste=0.930 ± 0.082, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.807 ± 0.085

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.38 DPS) | yes | Silk Headband (7050, -0.63 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.65 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.65 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (1.64 DPS) | yes | Crystal Starfire Medallion (5003, -1.56 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.12 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (2.14 DPS) | yes | Invoker's Mantle (215365, -0.52 DPS) [crafted]; Death Speaker Mantle (6685, -0.60 DPS) [dungeon]; Chestnut Mantle (17695, -0.96 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.08 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Battle Healer's Cloak (19529, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.81 DPS) | yes | High Robe of the Adjudicator (3461, -1.04 DPS) [quest]; Death Speaker Robes (6682, -1.06 DPS) [dungeon]; Green Silk Armor (7065, -1.13 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.94 DPS) | yes | Windsong Bangles (263336, -1.73 DPS) [quest]; Glowing Magical Bracelets (13106, -1.77 DPS) [world_drop]; Owlbeard Bracers (16981, -2.34 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.51 DPS) | yes | Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.37 DPS) [crafted]; Jutebraid Gloves (10654, -0.52 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.3 spell_power points (2.44 DPS) | yes | Warsong Sash (16975, -0.07 DPS) [quest]; Belt of Arugal (6392, -0.43 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.71 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.59 DPS) | yes | Abomination Skin Leggings (23173, -0.81 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.93 DPS) [crafted]; Silk-threaded Trousers (1929, -1.08 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.7 spell_power points (1.67 DPS) | yes | Acidic Walkers (9454, -0.41 DPS) [dungeon]; Boots of the Enchanter (4325, -0.59 DPS) [crafted]; Spidersilk Boots (4320, -1.92 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.51 DPS) | yes | Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Sludge-Stained Band (286535, -0.86 DPS) [world]; Sacred Band (6669, -1.08 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.30 DPS) | yes | Electrocutioner Lagnut (9447, -0.65 DPS) [dungeon]; Sacred Band (6669, -0.86 DPS) [quest]; Sludge-Stained Band (286535, -2.57 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.94 DPS) | yes | Twisted Chanter's Staff (890, -1.73 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.73 DPS) [quest]; Glimmering Staff (249392, -1.86 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.6 spell_power points (1.64 DPS) | yes | Orb of Souls (249395, -0.78 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.12 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.51 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 156.7 spell_power points (33.86 DPS) | yes | Starfaller (13063, -0.91 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.78 DPS) [crafted]; Gravestone Scepter (7001, -4.86 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 127.5. Weights run: 1.9s. Verify run: 1.1s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.111, intellect=0.181 ± 0.007, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.491 per %), hit=0.134 ± 0.001 per rating point (10 rating = 1%, 1.342 per %), spell_haste=not significant (0.359 ± 0.114), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.830 ± 0.111

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.15 DPS) | yes | Augural Shroud (2620, -2.01 DPS) [world]; Living Cowl (5608, -2.22 DPS, sim-verified) [world]; Holy Shroud (2721, -2.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.1 spell_power points (1.98 DPS) | yes | Triune Amulet (7722, -1.67 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.67 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.95 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.6 spell_power points (2.61 DPS) | yes | Inquisitor's Shawl (19507, -0.31 DPS) [dungeon]; Green Silken Shoulders (7057, -0.43 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 10.6 spell_power points (2.61 DPS) | yes | Long Silken Cloak (4326, -0.91 DPS) [crafted]; Guardian Cloak (5965, -0.91 DPS) [crafted]; Icy Cloak (4327, -1.57 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (5.66 DPS) | yes | Elemental Raiment (9434, -0.63 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.85 DPS) [crafted]; Robe of Power (7054, -1.70 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.21 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.65 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.87 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (4.59 DPS) | yes | Black Mageweave Gloves (10003, -1.07 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.45 DPS) [crafted]; Gilded Handwraps (254021, -2.32 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.7 spell_power points (3.61 DPS) | yes | Star Belt (4329, -0.42 DPS) [crafted]; Warsong Sash (16975, -0.91 DPS) [quest]; Deathmage Sash (10771, -1.23 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.2 spell_power points (3.97 DPS) | yes | Gaze Dreamer Pants (6903, -1.37 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.40 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.43 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.89 DPS) | yes | Gilded Slippers (254001, -2.98 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.99 DPS) [crafted]; Acidic Walkers (9454, -4.31 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.1 spell_power points (2.72 DPS) | yes | Reedknot Ring (9622, -1.00 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.25 DPS) [vendor]; Sludge-Stained Band (286535, -1.98 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.21 DPS) | yes | Reedknot Ring (9622, -0.65 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.74 DPS) [vendor]; Sludge-Stained Band (286535, -1.47 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (127.5 DPS) | yes | Scorn's Focal Dagger (23168, -2.70 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.09 DPS) [quest]; Gut Ripper (2164, -8.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 161.6 spell_power points (39.65 DPS) | yes | Umbral Wand (5216, -3.98 DPS) [dungeon]; Earthen Rod (9381, -4.06 DPS) [dungeon]; Twisted Nether Wand (249144, -4.18 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 224.6. Weights run: 2.1s. Verify run: 1.5s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.138, intellect=0.159 ± 0.012, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.426 per %), hit=0.200 ± 0.002 per rating point (10 rating = 1%, 2.004 per %), spell_haste=not significant (-0.577 ± 0.169), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.899 ± 0.138

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.90 DPS) | yes | Red Mageweave Headband (10033, -1.59 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.98 DPS) [crafted]; Dreamweave Circlet (10041, -2.90 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (2.62 DPS) | yes | Mindburst Medallion (11196, -0.68 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.89 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.10 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.1 spell_power points (5.29 DPS) | yes | Black Mageweave Shoulders (10027, -1.53 DPS) [crafted]; Bloodmage Mantle (7684, -1.86 DPS) [dungeon]; Rotgrip Mantle (17732, -2.33 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (4.93 DPS) | yes | Deep Woodlands Cloak (19121, -0.89 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.49 DPS) [dungeon]; Runecloth Cloak (13860, -1.54 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (7.56 DPS) | yes | Elemental Raiment (9434, -0.64 DPS) [world_drop]; Acumen Robes (17775, -0.90 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -1.16 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.97 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (6.14 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.39 DPS) [vendor]; Black Mageweave Gloves (10003, -1.46 DPS, sim-verified) [crafted]; Runecloth Gloves (13863, -1.72 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (5.14 DPS) | yes | Defiler's Cloth Girdle (20166, -0.31 DPS) [rep]; Ghostweave Cord (254073, -0.52 DPS) [crafted]; Ban'thok Sash (11662, -0.54 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 24.6 spell_power points (8.10 DPS) | yes | Red Mageweave Pants (10009, -2.86 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.38 DPS) [vendor]; Wizardweave Leggings (14132, -5.56 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.91 DPS) | yes | Gilded Sandals (254107, -1.11 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.92 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.14 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.28 DPS) | yes | Philanthropist's Ring (281635, -0.67 DPS) [quest]; Cyclopean Band (11824, -0.95 DPS) [dungeon]; Runed Ring (862, -1.98 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.95 DPS) | yes | Cyclopean Band (11824, -0.62 DPS) [dungeon]; Philanthropist's Ring (281635, -1.09 DPS, sim-verified) [quest]; Runed Ring (862, -1.65 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.98 DPS) [world_drop]; Rune of the Guard Captain (19120, -3.49 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.13 DPS) [quest]; Uther's Strength (11302, -0.50 DPS, sim-verified) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.62 DPS) [dungeon]; Arbiter's Blade (11784, -3.69 DPS) [dungeon]; Blade of Eternal Darkness (17780, -8.79 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (224.6 DPS) | yes | Wand of Allistarj (13065, -3.23 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.26 DPS) [crafted]; Pyric Caduceus (11748, -8.86 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 395.1. Weights run: 2.2s. Verify run: 1.5s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.161, intellect=0.119 ± 0.015, crit=0.076 ± 0.002 per rating point (14 rating = 1%, 1.063 per %), hit=0.322 ± 0.003 per rating point (10 rating = 1%, 3.218 per %), spell_haste=not significant (-0.510 ± 0.220), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.892 ± 0.161

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.0 spell_power points (11.09 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -2.42 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.68 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.88 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.95 DPS) [quest]; Kezan's Taint (19604, -2.52 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.8 spell_power points (10.69 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.01 DPS) [pvp]; Argent Shoulders (19059, -1.74 DPS) [crafted]; Burial Shawl (18681, -2.85 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.5 spell_power points (7.33 DPS) | yes | Amplifying Cloak (18350, -0.89 DPS) [dungeon]; Arcanoweave Cloak (272411, -1.25 DPS, sim-verified) [vendor]; Hide of the Wild (18510, -1.89 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (16.86 DPS) | yes | Robe of Everlasting Night (18385, -3.75 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -4.38 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.05 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.0 spell_power points (8.22 DPS) | yes | Sublime Wristguards (18497, -3.50 DPS) [dungeon]; Runecloth Cuffs (254123, -3.85 DPS) [crafted]; Spidertank Oilrag (9448, -5.00 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (9.88 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.00 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.8 spell_power points (11.05 DPS) | yes | Belt of the Archmage (18405, -2.68 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -4.17 DPS) [rep]; Oddly Magical Belt (18475, -5.31 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.68 DPS) [rep]; Sentinel's Silk Leggings (237815, -3.89 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.60 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.72 DPS) [crafted]; Omnicast Boots (11822, -0.92 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.60 DPS) [quest]; Maiden's Circle (13001, -1.60 DPS) [world_drop]; Naglering (11669, -6.83 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.07 DPS) [quest]; Maiden's Circle (13001, -1.07 DPS) [world_drop]; Naglering (11669, -7.09 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.15 DPS) [quest]; Weakness Analyzer (272438, -2.51 DPS) [vendor]; Draconic Infused Emblem (22268, -4.48 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -1.02 DPS, sim-verified) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.27 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -16.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+23.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.05 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.29 DPS) [world]; Torch of Light (279246, -23.25 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

