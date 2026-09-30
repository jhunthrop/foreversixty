# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 1.1s. Verify run: 0.9s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.142, intellect=0.365 ± 0.026, crit=0.050 ± 0.002 per rating point (14 rating = 1%, 0.693 per %), hit=0.203 ± 0.002 per rating point (10 rating = 1%, 2.033 per %), spell_haste=not significant (0.216 ± 0.166), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.142, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.56 DPS) | yes | Shadow Goggles (4373, -2.69 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.3 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.37 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.37 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.63 DPS) | yes | Green Woolen Robe (6243, -0.25 DPS) [crafted]; Green Woolen Vest (2582, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -1.09 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.8 spell_power points (0.17 DPS) | yes | Bright Bracers (3647, +0.00 DPS, sim-verified) [world_drop]; Windsong Bangles (263336, -0.08 DPS) [quest]; Repurposed Hair Band (281256, -0.10 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.65 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.18 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.51 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.24 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (36.0 DPS) | yes | Silk-threaded Trousers (1929, -0.11 DPS) [dungeon]; Rumpled Kilt (274741, -0.30 DPS) [vendor]; Abomination Skin Leggings (23173, -0.75 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.78 DPS) | yes | Pristine Boots (253889, -0.40 DPS) [crafted]; Red Woolen Boots (4313, -0.41 DPS) [crafted]; Feather Padded Treads (285345, -0.50 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.53 DPS) | yes | Sludge-Stained Band (286535, -0.25 DPS) [world]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Sludge-Stained Band (286535, -0.39 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.6 spell_power points (0.34 DPS) | yes | Channeler's Staff (4437, +0.00 DPS, sim-verified) [world]; Lesser Staff of the Spire (1300, -0.14 DPS) [world]; Staff of Westfall (2042, -0.17 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 243.7 spell_power points (22.59 DPS) | yes | Skycaller (12984, -1.29 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.30 DPS) [dungeon]; Deepblaze (279896, -4.06 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 55.7. Weights run: 1.2s. Verify run: 0.9s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.132, intellect=0.164 ± 0.010, crit=0.032 ± 0.002 per rating point (14 rating = 1%, 0.454 per %), hit=0.147 ± 0.002 per rating point (10 rating = 1%, 1.471 per %), spell_haste=1.022 ± 0.148, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.132, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.73 DPS) | yes | Embalmed Shroud (7691, -0.47 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.47 DPS) [crafted]; Silk Headband (7050, -0.54 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (1.25 DPS) | yes | Crystal Starfire Medallion (5003, -1.15 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.15 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.51 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (1.65 DPS) | yes | Death Speaker Mantle (6685, -0.42 DPS) [dungeon]; Fairywing Mantle (9536, -0.47 DPS) [quest]; Invoker's Mantle (215365, -0.57 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.79 DPS) | yes | Prelacy Cape (7004, -0.16 DPS) [quest]; Caretaker's Cape (19533, -0.16 DPS) [rep]; Heavy Woolen Cloak (4311, -0.30 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.04 DPS) | yes | Green Silk Armor (7065, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.66 DPS) [dungeon]; Pristine Gown (253961, -0.76 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.41 DPS) | yes | Windsong Bangles (263336, -1.26 DPS) [quest]; Nightsky Wristbands (6407, -1.26 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.51 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.10 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.16 DPS) [world]; Town Clerk's Mittens (270029, -0.19 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (1.81 DPS) | yes | Belt of Arugal (6392, -0.54 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.55 DPS) [dungeon]; Invoker's Cord (215366, -0.58 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.89 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.61 DPS) [crafted]; Silk-threaded Trousers (1929, -0.79 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (1.28 DPS) | yes | Acidic Walkers (9454, -0.29 DPS) [dungeon]; Nimbus Boots (6998, -0.34 DPS) [quest]; Spidersilk Boots (4320, -1.58 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.10 DPS) | yes | Minor Channeling Ring (1449, -0.26 DPS) [quest]; Electrocutioner Lagnut (9447, -0.63 DPS) [dungeon]; Sludge-Stained Band (286535, -0.63 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.94 DPS) | yes | Electrocutioner Lagnut (9447, -0.47 DPS) [dungeon]; Sludge-Stained Band (286535, -0.47 DPS) [world]; Minor Channeling Ring (1449, -1.58 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.41 DPS) | yes | Twisted Chanter's Staff (890, -1.16 DPS) [world_drop]; Channeler's Staff (4437, -1.21 DPS) [world]; Glimmering Staff (249392, -1.42 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (1.25 DPS) | yes | Eye of Paleth (2943, -0.63 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.63 DPS) [world]; Dwarven Tome (279898, -1.00 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 214.4 spell_power points (33.68 DPS) | yes | Starfaller (13063, -0.17 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.90 DPS) [crafted]; Gravestone Scepter (7001, -4.68 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 111.8. Weights run: 1.0s. Verify run: 0.8s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.308, intellect=0.697 ± 0.020, crit=0.060 ± 0.003 per rating point (14 rating = 1%, 0.836 per %), hit=0.289 ± 0.003 per rating point (10 rating = 1%, 2.892 per %), spell_haste=1.901 ± 0.341, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.308), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.29 DPS) | yes | Corpseshroud (10574, -0.84 DPS) [dungeon]; Living Cowl (5608, -0.87 DPS) [world]; Augural Shroud (2620, -2.22 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 spell_power points (1.22 DPS) | yes | Triune Amulet (7722, -0.69 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.69 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.66 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 spell_power points (1.75 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.3 spell_power points (1.66 DPS) | yes | Guardian Cloak (5965, -0.63 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.83 DPS) [vendor]; Long Silken Cloak (4326, -1.95 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (2.85 DPS) | yes | Robe of Power (7054, -0.42 DPS) [crafted]; Elemental Raiment (9434, -0.56 DPS) [world_drop]; Dreamweave Vest (10021, -0.97 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (0.98 DPS) | yes | Spidertank Oilrag (9448, -0.14 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.22 DPS) [quest]; Windchaser Cuffs (14429, -0.30 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (2.26 DPS) | yes | Black Mageweave Gloves (10003, -0.63 DPS) [crafted]; Gilded Handwraps (254021, -0.86 DPS) [crafted]; Red Mageweave Gloves (10018, -1.47 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (111.8 DPS) | yes | Gilded Cord (254037, -0.35 DPS) [crafted]; Star Belt (4329, -0.41 DPS) [crafted]; Deathmage Sash (10771, -1.40 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 spell_power points (2.44 DPS) | yes | Abomination Skin Leggings (23173, -0.85 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.07 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.37 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.61 DPS) | yes | Acidic Walkers (9454, -1.46 DPS) [dungeon]; Spidersilk Boots (4320, -1.55 DPS) [crafted]; Gilded Slippers (254001, -2.25 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (1.54 DPS) | yes | Ring of Forlorn Spirits (2043, -0.67 DPS) [quest]; Reedknot Ring (9622, -0.78 DPS) [quest]; Minor Channeling Ring (1449, -0.85 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.98 DPS) | yes | Ring of Forlorn Spirits (2043, -0.06 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.22 DPS) [quest]; Minor Channeling Ring (1449, -0.28 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Staff of Dar'Orahil (15106, -1.03 DPS) [quest]; Windweaver Staff (7757, -1.04 DPS) [dungeon]; Gut Ripper (2164, -6.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 365.8 spell_power points (39.84 DPS) | yes | Umbral Wand (5216, -0.15 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.25 DPS) [dungeon]; Twisted Nether Wand (249144, -5.18 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 154.4. Weights run: 1.0s. Verify run: 1.0s. 415 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.208, intellect=0.255 ± 0.013, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.535 per %), hit=0.149 ± 0.002 per rating point (10 rating = 1%, 1.493 per %), spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -0.97 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.85 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (2.39 DPS) | yes | Mindburst Medallion (11196, -0.25 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.39 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.68 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -1.41 DPS) [crafted]; Bloodmage Mantle (7684, -1.69 DPS) [dungeon]; Rotgrip Mantle (17732, -1.87 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (4.35 DPS) | yes | Mantle of Lady Falther'ess (23178, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.26 DPS) [crafted]; Nightfall Drape (12465, -1.83 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.1 spell_power points (6.75 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.87 DPS) [world_drop]; Dreamweave Vest (10021, -1.07 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.06 DPS) [crafted]; Bloodband Bracers (11469, -0.48 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (5.33 DPS) | yes | Black Mageweave Gloves (10003, -1.13 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.16 DPS, sim-verified) [vendor]; Runecloth Gloves (13863, -1.32 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (4.64 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.45 DPS) [dungeon]; Ghostweave Cord (254073, -0.71 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (7.16 DPS) | yes | Red Mageweave Pants (10009, -2.38 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.79 DPS) [vendor]; Wizardweave Leggings (14132, -2.87 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -0.92 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.14 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.42 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.41 DPS) [quest]; Cyclopean Band (11824, -0.62 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.36 DPS) | yes | Cyclopean Band (11824, -0.34 DPS) [dungeon]; Philanthropist's Ring (281635, -0.54 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.12 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.99 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.12 DPS, sim-verified) [crafted] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Arbiter's Blade (11784, -3.01 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.08 DPS) [dungeon]; Blade of Eternal Darkness (17780, -5.63 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Pyric Caduceus (11748, -2.23 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Noxious Shooter

No-known-source sample (15 of 415, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 342.4. Weights run: 1.1s. Verify run: 0.9s. 1022 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.972, intellect=not significant (0.230 ± 0.095), crit=0.233 ± 0.014 per rating point (14 rating = 1%, 3.263 per %), hit=0.719 ± 0.010 per rating point (10 rating = 1%, 7.192 per %), spell_haste=not significant (-0.152 ± 1.139), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.849 ± 0.972), fire_power=0.152 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 82.6 spell_power points (8.44 DPS) | yes | Field Marshal's Coronal (231584, -4.61 DPS) [pvp]; Crimson Felt Hat (18727, -5.19 DPS) [dungeon] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (2.25 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS, sim-verified) [quest]; Amulet of the Dawn (22657, -0.41 DPS) [quest]; Kezan's Taint (19604, -0.63 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 53.4 spell_power points (5.46 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.51 DPS) [pvp]; Rugged Mantle of the Timbermaw (227808, -6.72 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.0 spell_power points (2.56 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -0.72 DPS) [dungeon]; Hide of the Wild (18510, -0.89 DPS) [crafted] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 75.4 spell_power points (7.71 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.87 DPS) [pvp]; Robe of the Void (14153, -8.31 DPS, sim-verified) [crafted] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 36.2 spell_power points (3.70 DPS) | yes | Dryad's Wrist Bindings (19596, -1.51 DPS) [rep]; Heretic Wristguards (240152, -1.53 DPS) [vendor] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 47.2 spell_power points (4.82 DPS) | yes | Marshal's Dreadweave Gloves (231586, -1.61 DPS) [pvp]; Sandworm Skin Gloves (20716, -1.94 DPS) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 56.7 spell_power points (5.79 DPS) | yes | Heretic Waistguard (240151, -2.45 DPS) [vendor]; Belt of the Archmage (18405, -3.04 DPS) [crafted]; Knowledge of the Timbermaw (228190, -4.91 DPS, sim-verified) [vendor] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 72.6 spell_power points (7.42 DPS) | yes | Marshal's Dreadweave Leggings (231587, -3.19 DPS) [pvp]; Sentinel's Silk Leggings (237815, -3.25 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 46.4 spell_power points (4.74 DPS) | yes | Marshal's Dreadweave Boots (231585, -1.78 DPS) [pvp]; Earthen Silk Slippers (254013, -2.29 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (342.4 DPS) | yes | Songstone of Ironforge (12543, -1.80 DPS) [quest]; Maiden's Circle (13001, -1.80 DPS) [world_drop]; Naglering (11669, -11.22 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (342.4 DPS) | yes | Songstone of Ironforge (12543, -0.50 DPS) [quest]; Maiden's Circle (13001, -0.50 DPS) [world_drop]; Naglering (11669, -7.77 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (342.4 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (342.4 DPS) | yes | Weakness Analyzer (272438, -0.72 DPS) [vendor]; Serenity Field (272439, -1.53 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.55 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (342.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.31 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -14.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 757.4 spell_power points (77.37 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS, sim-verified) [dungeon]; Wand of Biting Cold (19108, -13.37 DPS) [quest]; Bonecreeper Stylus (13938, -13.53 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Chains of the Lich; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1022, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 34.9. Weights run: 1.1s. Verify run: 0.9s. 142 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.142, intellect=0.365 ± 0.026, crit=0.050 ± 0.002 per rating point (14 rating = 1%, 0.693 per %), hit=0.203 ± 0.002 per rating point (10 rating = 1%, 2.033 per %), spell_haste=not significant (0.216 ± 0.166), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.142, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.56 DPS) | yes | Shadow Goggles (4373, -2.72 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.3 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.37 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.43 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.63 DPS) | yes | Green Woolen Robe (6243, -0.25 DPS) [crafted]; Green Woolen Vest (2582, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -1.36 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.2 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Owlbeard Bracers (16981, -0.04 DPS) [quest]; Featherbead Bracers (15452, -0.06 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.65 DPS) | yes | Pristine Gloves (253913, -0.18 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.28 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.51 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.24 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (34.9 DPS) | yes | Silk-threaded Trousers (1929, -0.11 DPS) [dungeon]; Rumpled Kilt (274741, -0.30 DPS) [vendor]; Abomination Skin Leggings (23173, -0.67 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.78 DPS) | yes | Pristine Boots (253889, -0.40 DPS) [crafted]; Red Woolen Boots (4313, -0.41 DPS) [crafted]; Feather Padded Treads (285345, -0.67 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Loop of Sacrifice (281673, -0.29 DPS) [quest]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Volcanic Rock Ring (12053, -0.18 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.6 spell_power points (0.34 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.07 DPS) [world]; Lesser Staff of the Spire (1300, -0.14 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 243.7 spell_power points (22.59 DPS) | yes | Skycaller (12984, -1.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.30 DPS) [dungeon]; Sizzle Stick (8071, -4.47 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 142, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 54.1. Weights run: 1.2s. Verify run: 0.9s. 237 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.132, intellect=0.164 ± 0.010, crit=0.032 ± 0.002 per rating point (14 rating = 1%, 0.454 per %), hit=0.147 ± 0.002 per rating point (10 rating = 1%, 1.471 per %), spell_haste=1.022 ± 0.148, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.132, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.73 DPS) | yes | Silk Headband (7050, -0.42 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.47 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.47 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (1.25 DPS) | yes | Crystal Starfire Medallion (5003, -1.15 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.15 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.38 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (1.65 DPS) | yes | Invoker's Mantle (215365, -0.42 DPS) [crafted]; Death Speaker Mantle (6685, -0.42 DPS) [dungeon]; Chestnut Mantle (17695, -1.28 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.79 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.16 DPS) [crafted]; Battle Healer's Cloak (19529, -0.16 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.04 DPS) | yes | Green Silk Armor (7065, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.66 DPS) [dungeon]; High Robe of the Adjudicator (3461, -0.73 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.41 DPS) | yes | Glowing Magical Bracelets (13106, -1.21 DPS) [world_drop]; Windsong Bangles (263336, -1.26 DPS) [quest]; Owlbeard Bracers (16981, -1.48 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.10 DPS) | yes | Jutebraid Gloves (10654, +0.00 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.16 DPS) [world]; Truefaith Gloves (7049, -0.24 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (1.81 DPS) | yes | Warsong Sash (16975, -0.28 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.31 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.55 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.89 DPS) | yes | Abomination Skin Leggings (23173, -0.08 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.61 DPS) [crafted]; Silk-threaded Trousers (1929, -0.79 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (1.28 DPS) | yes | Acidic Walkers (9454, -0.29 DPS) [dungeon]; Boots of the Enchanter (4325, -0.49 DPS) [crafted]; Spidersilk Boots (4320, -1.75 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.10 DPS) | yes | Electrocutioner Lagnut (9447, -0.63 DPS) [dungeon]; Sludge-Stained Band (286535, -0.63 DPS) [world]; Sacred Band (6669, -0.79 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.94 DPS) | yes | Electrocutioner Lagnut (9447, -0.47 DPS) [dungeon]; Sacred Band (6669, -0.63 DPS) [quest]; Sludge-Stained Band (286535, -1.65 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.41 DPS) | yes | Twisted Chanter's Staff (890, -1.16 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.16 DPS) [quest]; Glimmering Staff (249392, -1.25 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (1.25 DPS) | yes | Orb of Souls (249395, -0.63 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.84 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.68 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 214.4 spell_power points (33.68 DPS) | yes | Starfaller (13063, -0.59 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.90 DPS) [crafted]; Gravestone Scepter (7001, -4.68 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 237, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 108.8. Weights run: 1.0s. Verify run: 0.8s. 319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.308, intellect=0.697 ± 0.020, crit=0.060 ± 0.003 per rating point (14 rating = 1%, 0.836 per %), hit=0.289 ± 0.003 per rating point (10 rating = 1%, 2.892 per %), spell_haste=1.901 ± 0.341, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.308), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.29 DPS) | yes | Corpseshroud (10574, -0.84 DPS) [dungeon]; Living Cowl (5608, -0.87 DPS) [world]; Augural Shroud (2620, -1.49 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 spell_power points (1.22 DPS) | yes | Triune Amulet (7722, -0.69 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.69 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.76 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 spell_power points (1.75 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.3 spell_power points (1.66 DPS) | yes | Guardian Cloak (5965, -0.63 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.83 DPS) [vendor]; Long Silken Cloak (4326, -1.60 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (2.85 DPS) | yes | Robe of Power (7054, -0.42 DPS) [crafted]; Elemental Raiment (9434, -0.56 DPS) [world_drop]; Dreamweave Vest (10021, -0.93 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (108.8 DPS) | yes | Condor Bracers (15864, -0.22 DPS) [quest]; Windchaser Cuffs (14429, -0.30 DPS) [world_drop]; Radiant Silver Bracers (4545, -1.26 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (2.26 DPS) | yes | Black Mageweave Gloves (10003, -0.63 DPS) [crafted]; Gilded Handwraps (254021, -0.86 DPS) [crafted]; Red Mageweave Gloves (10018, -1.40 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 spell_power points (1.90 DPS) | yes | Defiler's Cloth Girdle (20166, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.42 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 spell_power points (2.44 DPS) | yes | Abomination Skin Leggings (23173, -0.85 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.07 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.62 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.61 DPS) | yes | Acidic Walkers (9454, -1.46 DPS) [dungeon]; Spidersilk Boots (4320, -1.55 DPS) [crafted]; Gilded Slippers (254001, -2.00 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (1.54 DPS) | yes | Reedknot Ring (9622, -0.78 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.89 DPS) [vendor]; Black Widow Band (6199, -1.01 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.98 DPS) | yes | Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor]; Black Widow Band (6199, -0.45 DPS) [world]; Reedknot Ring (9622, -0.55 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Staff of Dar'Orahil (15106, -1.03 DPS) [quest]; Windweaver Staff (7757, -1.04 DPS) [dungeon]; Gut Ripper (2164, -6.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 365.8 spell_power points (39.84 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.25 DPS) [dungeon]; Twisted Nether Wand (249144, -5.18 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 149.7. Weights run: 1.0s. Verify run: 0.9s. 406 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.208, intellect=0.255 ± 0.013, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.535 per %), hit=0.149 ± 0.002 per rating point (10 rating = 1%, 1.493 per %), spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -0.97 DPS) [crafted]; Red Mageweave Headband (10033, -1.42 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (2.39 DPS) | yes | Mindburst Medallion (11196, -0.30 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.39 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.68 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (149.7 DPS) | yes | Black Mageweave Shoulders (10027, -1.41 DPS) [crafted]; Rotgrip Mantle (17732, -1.64 DPS, sim-verified) [dungeon]; Bloodmage Mantle (7684, -1.69 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (4.35 DPS) | yes | Deep Woodlands Cloak (19121, -0.45 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.19 DPS) [dungeon]; Runecloth Cloak (13860, -1.26 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.1 spell_power points (6.75 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.87 DPS) [world_drop]; Dreamweave Vest (10021, -1.07 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -0.48 DPS) [quest]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (5.33 DPS) | yes | Black Mageweave Gloves (10003, -1.13 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.29 DPS, sim-verified) [vendor]; Runecloth Gloves (13863, -1.32 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (4.64 DPS) | yes | Ban'thok Sash (11662, -0.45 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.58 DPS, sim-verified) [rep]; Ghostweave Cord (254073, -0.71 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (7.16 DPS) | yes | Red Mageweave Pants (10009, -2.38 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.79 DPS) [vendor]; Wizardweave Leggings (14132, -3.10 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -0.46 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.14 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.42 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.41 DPS) [quest]; Cyclopean Band (11824, -0.62 DPS) [dungeon]; Runed Ring (862, -1.68 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.36 DPS) | yes | Cyclopean Band (11824, -0.34 DPS) [dungeon]; Philanthropist's Ring (281635, -0.59 DPS, sim-verified) [quest]; Runed Ring (862, -1.40 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.99 DPS) [crafted]; Rune of the Guard Captain (19120, -3.07 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.69 DPS, sim-verified) [crafted]; Rune of the Guard Captain (19120, -1.39 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Arbiter's Blade (11784, -3.01 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.08 DPS) [dungeon]; Blade of Eternal Darkness (17780, -4.77 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 187.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.51 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Pyric Caduceus

No-known-source sample (15 of 406, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 340.0. Weights run: 1.1s. Verify run: 0.9s. 1013 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.972, intellect=not significant (0.230 ± 0.095), crit=0.233 ± 0.014 per rating point (14 rating = 1%, 3.263 per %), hit=0.719 ± 0.010 per rating point (10 rating = 1%, 7.192 per %), spell_haste=not significant (-0.152 ± 1.139), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.849 ± 0.972), fire_power=0.152 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 82.6 spell_power points (8.44 DPS) | yes | Warlord's Dreadweave Hood (231590, -4.61 DPS) [pvp]; Crimson Felt Hat (18727, -5.19 DPS) [dungeon] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (2.25 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS, sim-verified) [quest]; Amulet of the Dawn (22657, -0.41 DPS) [quest]; Kezan's Taint (19604, -0.63 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 53.4 spell_power points (5.46 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.51 DPS) [pvp]; Rugged Mantle of the Timbermaw (227808, -8.48 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.0 spell_power points (2.56 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -0.72 DPS) [dungeon]; Hide of the Wild (18510, -0.89 DPS) [crafted] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 75.4 spell_power points (7.71 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.87 DPS) [pvp]; Robe of the Void (14153, -8.47 DPS, sim-verified) [crafted] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 36.2 spell_power points (3.70 DPS) | yes | Dryad's Wrist Bindings (19596, -1.51 DPS) [rep]; Heretic Wristguards (240152, -1.53 DPS) [vendor] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 47.2 spell_power points (4.82 DPS) | yes | General's Dreadweave Gloves (231589, -1.61 DPS) [pvp]; Sandworm Skin Gloves (20716, -1.94 DPS) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 56.7 spell_power points (5.79 DPS) | yes | Heretic Waistguard (240151, -2.45 DPS) [vendor]; Belt of the Archmage (18405, -3.04 DPS) [crafted]; Knowledge of the Timbermaw (228190, -6.03 DPS, sim-verified) [vendor] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 72.6 spell_power points (7.42 DPS) | yes | General's Dreadweave Pants (231588, -3.19 DPS) [pvp]; Sentinel's Silk Leggings (237815, -3.25 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 46.4 spell_power points (4.74 DPS) | yes | General's Dreadweave Boots (231593, -1.78 DPS) [pvp]; Earthen Silk Slippers (254013, -2.29 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (340.0 DPS) | yes | Eye of Orgrimmar (12545, -1.80 DPS) [quest]; Maiden's Circle (13001, -1.80 DPS) [world_drop]; Naglering (11669, -11.47 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (340.0 DPS) | yes | Eye of Orgrimmar (12545, -0.50 DPS) [quest]; Maiden's Circle (13001, -0.50 DPS) [world_drop]; Naglering (11669, -7.76 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (340.0 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -8.26 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (340.0 DPS) | yes | Royal Seal of Eldre'Thalas (18467, -0.61 DPS) [quest]; Weakness Analyzer (272438, -0.72 DPS) [vendor]; Serenity Field (272439, -3.66 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (340.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.31 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -17.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 757.4 spell_power points (77.37 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.06 DPS, sim-verified) [dungeon]; Wand of Biting Cold (19108, -13.37 DPS) [quest]; Bonecreeper Stylus (13938, -13.53 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Chains of the Lich; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1013, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

