# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.6. Weights run: 2.1s. Verify run: 1.3s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.072, intellect=-0.054 ± 0.005, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.480 per %), hit=0.119 ± 0.001 per rating point (10 rating = 1%, 1.194 per %), spell_haste=not significant (-0.047 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.733 ± 0.072

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.93 DPS) | yes | Red Winter Hat (21524, -3.06 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.15 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.62 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Caretaker's Cape (20428, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.77 DPS) | yes | Green Woolen Vest (2582, -0.15 DPS) [crafted]; Bloody Apron (6226, -0.15 DPS) [dungeon]; Gray Woolen Robe (2585, -1.32 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Ivycloth Bracelets (9793, -0.26 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.08 DPS) | yes | Gnoll Casting Gloves (892, -0.30 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.46 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.77 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (38.6 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.94 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted]; Silk-threaded Trousers (1929, -0.57 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.62 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.08 DPS) | yes | Red Woolen Boots (4313, -0.46 DPS) [crafted]; Feather Padded Treads (285345, -0.48 DPS, sim-verified) [world]; Pristine Boots (253889, -0.62 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.77 DPS) | yes | Sludge-Stained Band (286535, -0.31 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.77 DPS) | yes | Sludge-Stained Band (286535, -0.45 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.28 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.77 DPS) | yes | Bouquet of Red Roses (22206, -1.22 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 147.2 spell_power points (22.77 DPS) | yes | Skycaller (12984, -1.49 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.48 DPS) [dungeon]; Deepblaze (279896, -4.24 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 63.3. Weights run: 2.1s. Verify run: 1.1s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.070, intellect=0.145 ± 0.005, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.443 per %), hit=0.112 ± 0.001 per rating point (10 rating = 1%, 1.123 per %), spell_haste=not significant (0.127 ± 0.079), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.793 ± 0.070

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.34 DPS) | yes | Silk Headband (7050, -0.56 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.64 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.64 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (1.68 DPS) | yes | Crystal Starfire Medallion (5003, -1.55 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.55 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.09 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (2.20 DPS) | yes | Invoker's Mantle (215365, -0.48 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.58 DPS) [dungeon]; Fairywing Mantle (9536, -0.64 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.07 DPS) | yes | Prelacy Cape (7004, -0.21 DPS) [quest]; Caretaker's Cape (19533, -0.21 DPS) [rep]; Heavy Woolen Cloak (4311, -0.29 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.77 DPS) | yes | Death Speaker Robes (6682, -0.94 DPS) [dungeon]; Green Silk Armor (7065, -1.01 DPS, sim-verified) [crafted]; Pristine Gown (253961, -1.06 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.92 DPS) | yes | Windsong Bangles (263336, -1.71 DPS) [quest]; Nightsky Wristbands (6407, -1.73 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.50 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.49 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.21 DPS) [world]; Town Clerk's Mittens (270029, -0.30 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.4 spell_power points (2.44 DPS) | yes | Belt of Arugal (6392, -0.54 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.73 DPS) [dungeon]; Invoker's Cord (215366, -0.79 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.56 DPS) | yes | Pristine Leggings (253987, -0.85 DPS) [crafted]; Abomination Skin Leggings (23173, -0.87 DPS, sim-verified) [dungeon]; Silk-threaded Trousers (1929, -1.07 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (1.71 DPS) | yes | Acidic Walkers (9454, -0.40 DPS) [dungeon]; Nimbus Boots (6998, -0.43 DPS) [quest]; Spidersilk Boots (4320, -1.74 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.49 DPS) | yes | Minor Channeling Ring (1449, -0.36 DPS) [quest]; Electrocutioner Lagnut (9447, -0.85 DPS) [dungeon]; Sludge-Stained Band (286535, -0.85 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.28 DPS) | yes | Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon]; Sludge-Stained Band (286535, -0.64 DPS) [world]; Minor Channeling Ring (1449, -2.01 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.92 DPS) | yes | Twisted Chanter's Staff (890, -1.61 DPS) [world_drop]; Channeler's Staff (4437, -1.67 DPS) [world]; Glimmering Staff (249392, -1.93 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.9 spell_power points (1.68 DPS) | yes | Dwarven Tome (279898, -0.52 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.82 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.82 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 158.8 spell_power points (33.85 DPS) | yes | Starfaller (13063, -0.80 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.78 DPS) [crafted]; Gravestone Scepter (7001, -4.85 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 117.9. Weights run: 1.7s. Verify run: 1.0s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.138, intellect=0.361 ± 0.009, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.590 per %), hit=0.150 ± 0.001 per rating point (10 rating = 1%, 1.499 per %), spell_haste=not significant (0.167 ± 0.130), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.806 ± 0.138

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.65 DPS) | yes | Living Cowl (5608, -1.77 DPS) [world]; Holy Shroud (2721, -2.21 DPS) [world_drop]; Augural Shroud (2620, -2.84 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.2 spell_power points (2.03 DPS) | yes | Triune Amulet (7722, -1.47 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.47 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.45 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.2 spell_power points (2.71 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Inquisitor's Shawl (19507, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.36 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.2 spell_power points (2.71 DPS) | yes | Guardian Cloak (5965, -0.98 DPS) [crafted]; Icy Cloak (4327, -1.16 DPS) [crafted]; Long Silken Cloak (4326, -2.13 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.2 spell_power points (5.35 DPS) | yes | Elemental Raiment (9434, -0.70 DPS) [world_drop]; Robe of Power (7054, -1.29 DPS) [crafted]; Dreamweave Vest (10021, -1.29 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.99 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.44 DPS) [quest]; Earthen Silk Cuffs (254019, -1.11 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.4 spell_power points (4.30 DPS) | yes | Black Mageweave Gloves (10003, -1.06 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.07 DPS) [crafted]; Gilded Handwraps (254021, -1.97 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.4 spell_power points (3.42 DPS) | yes | Star Belt (4329, -0.51 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -0.67 DPS) [dungeon]; Gilded Cord (254037, -1.01 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.3 spell_power points (4.06 DPS) | yes | Gaze Dreamer Pants (6903, -1.40 DPS) [dungeon]; Abomination Skin Leggings (23173, -1.43 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.75 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.31 DPS) | yes | Gilded Slippers (254001, -3.21 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.44 DPS) [crafted]; Acidic Walkers (9454, -3.57 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.2 spell_power points (2.69 DPS) | yes | Ring of Forlorn Spirits (2043, -0.92 DPS) [quest]; Reedknot Ring (9622, -1.14 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.36 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.99 DPS) | yes | Reedknot Ring (9622, -0.44 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.52 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.66 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (117.9 DPS) | yes | Scorn's Focal Dagger (23168, -2.43 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.22 DPS) [quest]; Gut Ripper (2164, -8.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 180.1 spell_power points (39.86 DPS) | yes | Umbral Wand (5216, -4.19 DPS) [dungeon]; Earthen Rod (9381, -4.27 DPS) [dungeon]; Twisted Nether Wand (249144, -4.53 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 207.2. Weights run: 1.9s. Verify run: 1.3s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.132, intellect=0.253 ± 0.011, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.441 per %), hit=0.194 ± 0.002 per rating point (10 rating = 1%, 1.944 per %), spell_haste=not significant (-0.086 ± 0.139), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.889 ± 0.132

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.01 DPS) | yes | Dreamweave Circlet (10041, -1.16 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.00 DPS) [crafted]; Red Mageweave Headband (10033, -2.35 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (2.84 DPS) | yes | Mindburst Medallion (11196, -0.33 DPS) [quest]; Horizon Choker (13085, -1.66 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.00 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.5 spell_power points (5.86 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.76 DPS) [crafted]; Bloodmage Mantle (7684, -2.09 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (5.18 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.67 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.50 DPS) [crafted]; Nightfall Drape (12465, -2.18 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.1 spell_power points (8.03 DPS) | yes | Robe of the Magi (1716, -0.18 DPS) [world_drop]; Elemental Raiment (9434, -1.02 DPS) [world_drop]; Dreamweave Vest (10021, -1.26 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.00 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.08 DPS) [crafted]; Bloodband Bracers (11469, -0.58 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (6.35 DPS) | yes | Black Mageweave Gloves (10003, -1.34 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.35 DPS, sim-verified) [vendor]; Runecloth Gloves (13863, -1.58 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (5.52 DPS) | yes | Highlander's Cloth Girdle (20098, -0.51 DPS) [rep]; Ban'thok Sash (11662, -0.52 DPS) [dungeon]; Ghostweave Cord (254073, -0.84 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (8.52 DPS) | yes | Red Mageweave Pants (10009, -2.84 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.36 DPS) [vendor]; Wizardweave Leggings (14132, -4.89 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.01 DPS) | yes | Gilded Sandals (254107, -3.58 DPS) [crafted]; Black Mageweave Boots (10026, -3.75 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.93 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.34 DPS) | yes | Philanthropist's Ring (281635, -0.50 DPS) [quest]; Cyclopean Band (11824, -0.75 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.67 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.01 DPS) | yes | Cyclopean Band (11824, -0.41 DPS) [dungeon]; Philanthropist's Ring (281635, -0.68 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.34 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.13 DPS, sim-verified) [crafted] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.40 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, -3.42 DPS) [crafted] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -3.58 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.67 DPS) [dungeon]; Blade of Eternal Darkness (17780, -7.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (207.2 DPS) | yes | Lesser Eternal Wand (249232, -3.25 DPS) [crafted]; Wand of Allistarj (13065, -3.25 DPS) [world_drop]; Pyric Caduceus (11748, -7.74 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 419.8. Weights run: 5.6s. Verify run: 1.4s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.266, intellect=0.050 ± 0.022, crit=0.073 ± 0.002 per rating point (14 rating = 1%, 1.015 per %), hit=0.338 ± 0.004 per rating point (10 rating = 1%, 3.377 per %), spell_haste=not significant (-0.457 ± 0.357), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.894 ± 0.266)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.4 spell_power points (11.50 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -2.83 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -6.59 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.33 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.40 DPS) [quest]; Kezan's Taint (19604, -2.88 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.8 spell_power points (10.89 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.10 DPS) [pvp]; Argent Shoulders (19059, -1.56 DPS, sim-verified) [crafted]; Burial Shawl (18681, -3.01 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.2 spell_power points (7.64 DPS) | yes | Arcanoweave Cloak (272411, -0.16 DPS) [vendor]; Amplifying Cloak (18350, -0.83 DPS) [dungeon]; Hide of the Wild (18510, -2.16 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.4 spell_power points (17.58 DPS) | yes | Robe of Everlasting Night (18385, -2.07 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -5.02 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.74 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.4 spell_power points (8.48 DPS) | yes | Sublime Wristguards (18497, -3.75 DPS) [dungeon]; Runecloth Cuffs (254123, -4.13 DPS) [crafted]; Arcane Runed Bracers (4744, -5.07 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.63 DPS) [quest]; Sandworm Skin Gloves (20716, -4.31 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 29.5 spell_power points (11.15 DPS) | yes | Stormpike Cloth Girdle (19094, -4.15 DPS) [rep]; Belt of the Archmage (18405, -4.49 DPS, sim-verified) [crafted]; Oddly Magical Belt (18475, -5.10 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -2.06 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.08 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.76 DPS) [crafted]; Omnicast Boots (11822, -1.29 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.59 DPS) [quest]; Maiden's Circle (13001, -1.59 DPS) [world_drop]; Naglering (11669, -8.27 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.14 DPS) [quest]; Maiden's Circle (13001, -1.14 DPS) [world_drop]; Naglering (11669, -8.92 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.65 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.71 DPS, sim-verified) [quest]; Serenity Field (272439, -5.68 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.37 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+20.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.14 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.62 DPS) [world]; Torch of Light (279246, -20.66 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.3. Weights run: 2.1s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.072, intellect=-0.054 ± 0.005, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.480 per %), hit=0.119 ± 0.001 per rating point (10 rating = 1%, 1.194 per %), spell_haste=not significant (-0.047 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.733 ± 0.072

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.93 DPS) | yes | Red Winter Hat (21524, -2.75 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.15 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.62 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.77 DPS) | yes | Green Woolen Vest (2582, -0.15 DPS) [crafted]; Bloody Apron (6226, -0.15 DPS) [dungeon]; Gray Woolen Robe (2585, -1.24 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) (or Owlbeard Bracers (16981)) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.08 DPS) | yes | Gnoll Casting Gloves (892, -0.29 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.46 DPS) [quest]; Pristine Gloves (253913, -0.46 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (36.3 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted]; Silk-threaded Trousers (1929, -0.60 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.62 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.08 DPS) | yes | Red Woolen Boots (4313, -0.46 DPS) [crafted]; Feather Padded Treads (285345, -0.55 DPS, sim-verified) [world]; Pristine Boots (253889, -0.62 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.77 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.46 DPS) | yes | Ring of the Shadow (1462, -0.65 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 8 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 147.2 spell_power points (22.77 DPS) | yes | Skycaller (12984, -1.62 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.48 DPS) [dungeon]; Sizzle Stick (8071, -4.35 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 62.6. Weights run: 2.1s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.070, intellect=0.145 ± 0.005, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.443 per %), hit=0.112 ± 0.001 per rating point (10 rating = 1%, 1.123 per %), spell_haste=not significant (0.127 ± 0.079), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.793 ± 0.070

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.34 DPS) | yes | Silk Headband (7050, -0.51 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.64 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.64 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (1.68 DPS) | yes | Crystal Starfire Medallion (5003, -1.55 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.55 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.76 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (2.20 DPS) | yes | Invoker's Mantle (215365, -0.55 DPS) [crafted]; Death Speaker Mantle (6685, -0.58 DPS) [dungeon]; Chestnut Mantle (17695, -1.99 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.07 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted]; Battle Healer's Cloak (19529, -0.21 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.77 DPS) | yes | Death Speaker Robes (6682, -0.94 DPS) [dungeon]; Green Silk Armor (7065, -0.95 DPS, sim-verified) [crafted]; High Robe of the Adjudicator (3461, -1.00 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.92 DPS) | yes | Glowing Magical Bracelets (13106, -1.67 DPS) [world_drop]; Windsong Bangles (263336, -1.71 DPS) [quest]; Owlbeard Bracers (16981, -2.09 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.49 DPS) | yes | Gnoll Casting Gloves (892, -0.21 DPS) [world]; Jutebraid Gloves (10654, -0.28 DPS, sim-verified) [quest]; Truefaith Gloves (7049, -0.33 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.4 spell_power points (2.44 DPS) | yes | Warsong Sash (16975, -0.09 DPS) [quest]; Belt of Arugal (6392, -0.43 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.73 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.56 DPS) | yes | Abomination Skin Leggings (23173, -0.62 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Silk-threaded Trousers (1929, -1.07 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (1.71 DPS) | yes | Acidic Walkers (9454, -0.40 DPS) [dungeon]; Boots of the Enchanter (4325, -0.64 DPS) [crafted]; Spidersilk Boots (4320, -1.68 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.49 DPS) | yes | Electrocutioner Lagnut (9447, -0.85 DPS) [dungeon]; Sludge-Stained Band (286535, -0.85 DPS) [world]; Sacred Band (6669, -1.07 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.28 DPS) | yes | Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon]; Sacred Band (6669, -0.85 DPS) [quest]; Sludge-Stained Band (286535, -2.19 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.92 DPS) | yes | Glimmering Staff (249392, -1.45 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.61 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.61 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.9 spell_power points (1.68 DPS) | yes | Orb of Souls (249395, -0.82 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.13 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.70 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 158.8 spell_power points (33.85 DPS) | yes | Starfaller (13063, -0.62 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.78 DPS) [crafted]; Gravestone Scepter (7001, -4.85 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 116.7. Weights run: 1.7s. Verify run: 1.0s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.138, intellect=0.361 ± 0.009, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.590 per %), hit=0.150 ± 0.001 per rating point (10 rating = 1%, 1.499 per %), spell_haste=not significant (0.167 ± 0.130), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.806 ± 0.138

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.65 DPS) | yes | Living Cowl (5608, -1.77 DPS) [world]; Holy Shroud (2721, -2.21 DPS) [world_drop]; Augural Shroud (2620, -2.46 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.2 spell_power points (2.03 DPS) | yes | Triune Amulet (7722, -1.47 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.47 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.23 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.2 spell_power points (2.71 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Inquisitor's Shawl (19507, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.36 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.2 spell_power points (2.71 DPS) | yes | Guardian Cloak (5965, -0.98 DPS) [crafted]; Icy Cloak (4327, -1.16 DPS) [crafted]; Long Silken Cloak (4326, -1.40 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.2 spell_power points (5.35 DPS) | yes | Elemental Raiment (9434, -0.70 DPS) [world_drop]; Dreamweave Vest (10021, -1.25 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.29 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.99 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.44 DPS) [quest]; Radiant Silver Bracers (4545, -0.47 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.4 spell_power points (4.30 DPS) | yes | Black Mageweave Gloves (10003, -0.98 DPS) [crafted]; Red Mageweave Gloves (10018, -1.07 DPS) [crafted]; Gilded Handwraps (254021, -1.97 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.4 spell_power points (3.42 DPS) | yes | Star Belt (4329, -0.54 DPS) [crafted]; Deathmage Sash (10771, -0.67 DPS) [dungeon]; Warsong Sash (16975, -0.98 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.3 spell_power points (4.06 DPS) | yes | Gaze Dreamer Pants (6903, -1.40 DPS) [dungeon]; Abomination Skin Leggings (23173, -1.43 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.94 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.31 DPS) | yes | Gilded Slippers (254001, -2.99 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.44 DPS) [crafted]; Acidic Walkers (9454, -3.57 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.2 spell_power points (2.69 DPS) | yes | Reedknot Ring (9622, -1.14 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.36 DPS) [vendor]; Sludge-Stained Band (286535, -2.03 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.99 DPS) | yes | Reedknot Ring (9622, -0.44 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.66 DPS) [vendor]; Sludge-Stained Band (286535, -1.33 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (116.7 DPS) | yes | Scorn's Focal Dagger (23168, -2.43 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.22 DPS) [quest]; Gut Ripper (2164, -7.53 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 180.1 spell_power points (39.86 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.27 DPS) [dungeon]; Twisted Nether Wand (249144, -4.53 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 207.0. Weights run: 1.9s. Verify run: 1.3s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.132, intellect=0.253 ± 0.011, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.441 per %), hit=0.194 ± 0.002 per rating point (10 rating = 1%, 1.944 per %), spell_haste=not significant (-0.086 ± 0.139), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.889 ± 0.132

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.01 DPS) | yes | Red Mageweave Headband (10033, -1.15 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.16 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.00 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (2.84 DPS) | yes | Mindburst Medallion (11196, -0.33 DPS) [quest]; Horizon Choker (13085, -1.66 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.00 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -1.67 DPS) [crafted]; Bloodmage Mantle (7684, -2.01 DPS) [dungeon]; Rotgrip Mantle (17732, -2.24 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (5.18 DPS) | yes | Deep Woodlands Cloak (19121, -0.81 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.42 DPS) [dungeon]; Runecloth Cloak (13860, -1.50 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.1 spell_power points (8.03 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.02 DPS) [world_drop]; Dreamweave Vest (10021, -1.26 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -1.98 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (6.35 DPS) | yes | Black Mageweave Gloves (10003, -1.34 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.34 DPS, sim-verified) [vendor]; Runecloth Gloves (13863, -1.58 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (5.52 DPS) | yes | Defiler's Cloth Girdle (20166, -0.51 DPS) [rep]; Ban'thok Sash (11662, -0.52 DPS) [dungeon]; Ghostweave Cord (254073, -0.84 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (8.52 DPS) | yes | Red Mageweave Pants (10009, -2.84 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.36 DPS) [vendor]; Wizardweave Leggings (14132, -4.44 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.01 DPS) | yes | Gilded Sandals (254107, -3.58 DPS) [crafted]; Black Mageweave Boots (10026, -3.75 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.93 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.34 DPS) | yes | Philanthropist's Ring (281635, -0.50 DPS) [quest]; Cyclopean Band (11824, -0.75 DPS) [dungeon]; Runed Ring (862, -2.00 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.01 DPS) | yes | Philanthropist's Ring (281635, -0.16 DPS) [quest]; Cyclopean Band (11824, -0.41 DPS) [dungeon]; Runed Ring (862, -1.67 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.96 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -3.55 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.13 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -3.58 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.67 DPS) [dungeon]; Blade of Eternal Darkness (17780, -7.52 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.25 DPS) [crafted]; Wand of Allistarj (13065, -3.25 DPS) [world_drop]; Pyric Caduceus (11748, -7.55 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 407.7. Weights run: 5.6s. Verify run: 1.4s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.266, intellect=0.050 ± 0.022, crit=0.073 ± 0.002 per rating point (14 rating = 1%, 1.015 per %), hit=0.338 ± 0.004 per rating point (10 rating = 1%, 3.377 per %), spell_haste=not significant (-0.457 ± 0.357), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.894 ± 0.266)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.4 spell_power points (11.50 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -2.83 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -6.03 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.33 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.40 DPS) [quest]; Kezan's Taint (19604, -2.88 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.8 spell_power points (10.89 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.10 DPS) [pvp]; Argent Shoulders (19059, -1.42 DPS) [crafted]; Burial Shawl (18681, -3.01 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.2 spell_power points (7.64 DPS) | yes | Amplifying Cloak (18350, -0.83 DPS) [dungeon]; Hide of the Wild (18510, -2.16 DPS) [crafted]; Arcanoweave Cloak (272411, -2.28 DPS, sim-verified) [vendor] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.4 spell_power points (17.58 DPS) | yes | Robe of Everlasting Night (18385, -4.24 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -5.02 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.74 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.4 spell_power points (8.48 DPS) | yes | Sublime Wristguards (18497, -3.75 DPS) [dungeon]; Runecloth Cuffs (254123, -4.13 DPS) [crafted]; Spidertank Oilrag (9448, -5.07 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.2 spell_power points (10.31 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.99 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 29.5 spell_power points (11.15 DPS) | yes | Frostwolf Cloth Belt (19090, -4.15 DPS) [rep]; Belt of the Archmage (18405, -4.72 DPS, sim-verified) [crafted]; Oddly Magical Belt (18475, -5.10 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 34.5 spell_power points (13.07 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.11 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.08 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.76 DPS) [crafted]; Omnicast Boots (11822, -1.29 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.59 DPS) [quest]; Maiden's Circle (13001, -1.59 DPS) [world_drop]; Naglering (11669, -10.30 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.14 DPS) [quest]; Maiden's Circle (13001, -1.14 DPS) [world_drop]; Naglering (11669, -9.72 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+17.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.27 DPS) [quest]; Weakness Analyzer (272438, -2.65 DPS) [vendor]; Serenity Field (272439, -3.83 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.37 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -19.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (407.7 DPS) | yes | Bonecreeper Stylus (13938, -1.14 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.62 DPS) [world]; Torch of Light (279246, -19.30 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

