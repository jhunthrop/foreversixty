# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.2. Weights run: 2.6s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.067, intellect=0.082 ± 0.004, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.490 per %), hit=0.148 ± 0.005 per rating point (10 rating = 1%, 1.484 per %), spell_haste=not significant (0.199 ± 0.068), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.688 ± 0.067, fire_power=0.313 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.92 DPS) | yes | Shadow Goggles (4373, -3.21 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.7 spell_power points (0.88 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.27 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.39 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.61 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Caretaker's Cape (20428, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.4 spell_power points (0.83 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -1.47 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Bright Bracers (3647, -0.10 DPS) [world_drop]; Repurposed Hair Band (281256, -0.13 DPS) [quest]; Mindthrust Bracers (1974, -0.32 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.07 DPS) | yes | Gnoll Casting Gloves (892, -0.28 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.74 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.3 spell_power points (0.66 DPS) | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Keller's Girdle (2911, -0.56 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.7 spell_power points (1.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.49 DPS) [crafted]; Rumpled Kilt (274741, -0.71 DPS) [vendor]; Silk-threaded Trousers (1929, -0.81 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.3 spell_power points (1.12 DPS) | yes | Red Woolen Boots (4313, -0.51 DPS) [crafted]; Pristine Boots (253889, -0.63 DPS) [crafted]; Feather Padded Treads (285345, -0.93 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.2 spell_power points (0.79 DPS) | yes | Sludge-Stained Band (286535, -0.33 DPS) [world]; Lavishly Jeweled Ring (1156, -0.72 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.75 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.77 DPS) | yes | Sludge-Stained Band (286535, -0.59 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.69 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.8 spell_power points (0.13 DPS) | yes | Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.05 DPS) [world]; Staff of Westfall (2042, -0.06 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 148.5 spell_power points (22.77 DPS) | yes | Skycaller (12984, -1.74 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.48 DPS) [dungeon]; Deepblaze (279896, -4.24 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053225101000000)

Set DPS (verified): 71.6. Weights run: 2.7s. Verify run: 1.3s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.093, intellect=0.172 ± 0.006, crit=0.051 ± 0.002 per rating point (14 rating = 1%, 0.719 per %), hit=0.124 ± 0.006 per rating point (10 rating = 1%, 1.241 per %), spell_haste=not significant (-0.346 ± 0.090), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.803 ± 0.093, fire_power=0.195 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.93 DPS) | yes | Silk Headband (7050, -0.69 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.80 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.80 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (2.14 DPS) | yes | Crystal Starfire Medallion (5003, -1.95 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.95 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.22 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (2.81 DPS) | yes | Death Speaker Mantle (6685, -0.68 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.72 DPS) [crafted]; Fairywing Mantle (9536, -0.80 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.33 DPS) | yes | Prelacy Cape (7004, -0.27 DPS) [quest]; Caretaker's Cape (19533, -0.27 DPS) [rep]; Heavy Woolen Cloak (4311, -0.41 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.46 DPS) | yes | Green Silk Armor (7065, -0.87 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.09 DPS) [dungeon]; Pristine Gown (253961, -1.28 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.40 DPS) | yes | Nightsky Wristbands (6407, -2.12 DPS) [world_drop]; Windsong Bangles (263336, -2.13 DPS) [quest]; Glowing Magical Bracelets (13106, -2.42 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.86 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.27 DPS) [world]; Town Clerk's Mittens (270029, -0.30 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (3.07 DPS) | yes | Belt of Arugal (6392, -0.71 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.94 DPS) [dungeon]; Invoker's Cord (215366, -0.97 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.19 DPS) | yes | Abomination Skin Leggings (23173, -0.78 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.01 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.32 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.2 spell_power points (2.18 DPS) | yes | Acidic Walkers (9454, -0.49 DPS) [dungeon]; Nimbus Boots (6998, -0.59 DPS) [quest]; Spidersilk Boots (4320, -2.10 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.86 DPS) | yes | Minor Channeling Ring (1449, -0.44 DPS) [quest]; Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon]; Sludge-Stained Band (286535, -1.06 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.60 DPS) | yes | Electrocutioner Lagnut (9447, -0.80 DPS) [dungeon]; Sludge-Stained Band (286535, -0.80 DPS) [world]; Minor Channeling Ring (1449, -1.89 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.40 DPS) | yes | Twisted Chanter's Staff (890, -1.94 DPS) [world_drop]; Glimmering Staff (249392, -1.94 DPS, sim-verified) [crafted]; Channeler's Staff (4437, -2.03 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (2.14 DPS) | yes | Dwarven Tome (279898, -1.03 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.07 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -1.07 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 127.8 spell_power points (34.01 DPS) | yes | Starfaller (13063, -0.97 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.68 DPS) [crafted]; Gravestone Scepter (7001, -5.01 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053225103101330)

Set DPS (verified): 140.1. Weights run: 2.2s. Verify run: 1.2s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.269, intellect=0.494 ± 0.018, crit=0.097 ± 0.003 per rating point (14 rating = 1%, 1.353 per %), hit=0.293 ± 0.019 per rating point (10 rating = 1%, 2.931 per %), spell_haste=not significant (0.728 ± 0.240), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.660 ± 0.269), fire_power=0.341 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.58 DPS) | yes | Augural Shroud (2620, -1.97 DPS, sim-verified) [world]; Living Cowl (5608, -2.13 DPS) [world]; Holy Shroud (2721, -2.66 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.0 spell_power points (2.65 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.48 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.73 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.73 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.4 spell_power points (3.57 DPS) | yes | Green Silken Shoulders (7057, -0.00 DPS) [crafted]; Inquisitor's Shawl (19507, -0.01 DPS) [dungeon]; Berylline Pads (4197, -0.40 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.4 spell_power points (3.57 DPS) | yes | Guardian Cloak (5965, -1.32 DPS) [crafted]; Long Silken Cloak (4326, -1.43 DPS, sim-verified) [crafted]; Icy Cloak (4327, -1.71 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.0 spell_power points (6.63 DPS) | yes | Dreamweave Vest (10021, -0.94 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.05 DPS) [world_drop]; Robe of Power (7054, -1.34 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.39 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.53 DPS) [quest]; Windchaser Cuffs (14429, -1.21 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.0 spell_power points (5.31 DPS) | yes | Red Mageweave Gloves (10018, -1.28 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.32 DPS) [crafted]; Gilded Handwraps (254021, -2.26 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.0 spell_power points (4.24 DPS) | yes | Star Belt (4329, -0.79 DPS) [crafted]; Deathmage Sash (10771, -0.96 DPS, sim-verified) [dungeon]; Gilded Cord (254037, -1.07 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.9 spell_power points (5.30 DPS) | yes | Crimson Silk Pantaloons (7062, -1.67 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.85 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.11 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.38 DPS) | yes | Gilded Slippers (254001, -2.32 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.99 DPS) [crafted]; Acidic Walkers (9454, -4.00 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (3.44 DPS) | yes | Ring of Forlorn Spirits (2043, -1.32 DPS) [quest]; Reedknot Ring (9622, -1.58 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.85 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.39 DPS) | yes | Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.53 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.80 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -2.92 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.09 DPS) [quest]; Gut Ripper (2164, -7.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-verified (140.1 DPS) | yes | Twisted Nether Wand (249144, -0.08 DPS) [crafted]; Earthen Rod (9381, -0.08 DPS) [dungeon]; Jaina's Firestarter (13064, -1.92 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Umbral Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 197.9. Weights run: 8.7s. Verify run: 1.4s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.297, intellect=1.091 ± 0.020, crit=0.210 ± 0.004 per rating point (14 rating = 1%, 2.945 per %), hit=0.598 ± 0.020 per rating point (10 rating = 1%, 5.982 per %), spell_haste=2.334 ± 0.335, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.097 ± 0.297), fire_power=0.904 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.8 spell_power points (7.36 DPS) | yes | Dreamweave Circlet (10041, +0.00 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.94 DPS) [vendor]; Chief Architect's Monocle (11839, -2.05 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.18 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.47 DPS) [quest]; Horizon Choker (13085, -2.23 DPS, sim-verified) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Shoulders (10029, -0.87 DPS) [crafted]; Inquisitor's Shawl (19507, -1.26 DPS) [dungeon]; Rotgrip Mantle (17732, -1.95 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.5 spell_power points (3.70 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.31 DPS) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.95 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.8 spell_power points (7.36 DPS) | yes | Runecloth Tunic (13857, -2.13 DPS) [crafted]; Robe of the Magi (1716, -2.21 DPS) [world_drop]; Runecloth Robe (13858, -2.31 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 16.4 spell_power points (2.95 DPS) | yes | Bloodband Bracers (11469, +0.00 DPS, sim-verified) [quest]; Nethergeld Cuffs (254061, -0.31 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.59 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 31.0 spell_power points (5.59 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.47 DPS) [vendor]; Dreamweave Gloves (10019, -1.56 DPS) [crafted]; Red Mageweave Gloves (10018, -1.64 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 30.0 spell_power points (5.40 DPS) | yes | Satyrmane Sash (17755, -0.91 DPS) [dungeon]; Deathmage Sash (10771, -1.19 DPS) [dungeon]; Dawnspire Cord (12466, -4.37 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 33.9 spell_power points (6.11 DPS) | yes | Red Mageweave Pants (10009, -1.23 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.03 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -4.36 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.33 DPS) | yes | Gilded Sandals (254107, -0.57 DPS) [crafted]; Southsea Mojo Boots (20641, -0.72 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -3.46 DPS, sim-verified) [vendor] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.6 spell_power points (3.00 DPS) | yes | Brainlash (6440, -0.05 DPS) [dungeon]; Mindseye Circle (10634, -0.64 DPS) [dungeon]; Band of the Unicorn (7553, -0.66 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.5 spell_power points (2.98 DPS) | yes | Mindseye Circle (10634, -0.62 DPS) [dungeon]; Band of the Unicorn (7553, -0.64 DPS) [world_drop]; Brainlash (6440, -2.88 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.44 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Soul Harvester (20536, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -1.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.48 DPS) [world_drop]; Flaming Incinerator (9483, -3.68 DPS) [dungeon]; Pyric Caduceus (11748, -5.46 DPS, sim-verified) [dungeon] |

**New at 50:** head: Red Mageweave Headband; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Abyss Shard; main_hand: Blade of Eternal Darkness; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 425.6. Weights run: 2.3s. Verify run: 1.4s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.374, intellect=0.433 ± 0.025, crit=0.144 ± 0.004 per rating point (14 rating = 1%, 2.013 per %), hit=0.453 ± 0.026 per rating point (10 rating = 1%, 4.530 per %), spell_haste=-1.810 ± 0.408, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.602 ± 0.374), fire_power=0.397 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.5 spell_power points (16.74 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.33 DPS) [pvp]; Deathmist Mask (226909, -7.28 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (11.00 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.69 DPS) [quest]; Beads of Ogre Mojo (22149, -1.90 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.5 spell_power points (17.76 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.57 DPS) [pvp]; Burial Shawl (18681, -2.16 DPS, sim-verified) [dungeon]; Argent Shoulders (19059, -5.25 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.0 spell_power points (12.00 DPS) | yes | Crystalline Threaded Cape (20697, -1.58 DPS, sim-verified) [world]; Hide of the Wild (18510, -2.83 DPS) [crafted]; Amplifying Cloak (18350, -3.00 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.9 spell_power points (24.95 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.75 DPS) [pvp]; Robe of Everlasting Night (18385, -5.60 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -8.12 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.5 spell_power points (12.73 DPS) | yes | Sublime Wristguards (18497, -4.57 DPS) [dungeon]; Runecloth Cuffs (254123, -5.07 DPS) [crafted]; Deathmist Bracers (226907, -6.64 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.2 spell_power points (14.59 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -3.00 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.1 spell_power points (19.53 DPS) | yes | Belt of the Archmage (18405, -7.55 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -8.36 DPS) [rep]; Ban'thok Sash (11662, -8.88 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.4 spell_power points (20.18 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.45 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -3.36 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.2 spell_power points (12.60 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.13 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -2.87 DPS) [quest]; Maiden's Circle (13001, -2.87 DPS) [world_drop]; Naglering (11669, -12.75 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.50 DPS) [quest]; Maiden's Circle (13001, -1.50 DPS) [world_drop]; Naglering (11669, -10.87 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+19.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -3.00 DPS) [quest]; Weakness Analyzer (272438, -3.50 DPS) [vendor]; Serenity Field (272439, -7.50 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.48 DPS) [world]; Teebu's Blazing Longsword (1728, -24.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (425.6 DPS) | yes | Bonecreeper Stylus (13938, -0.35 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.57 DPS) [world]; Torch of Light (279246, -19.17 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 802.6. Weights run: 7.1s. Verify run: 1.5s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.184, intellect=0.084 ± 0.023, crit=0.197 ± 0.006 per rating point (14 rating = 1%, 2.752 per %), hit=0.580 ± 0.037 per rating point (10 rating = 1%, 5.797 per %), spell_haste=not significant (0.037 ± 0.500), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.501 ± 0.184), fire_power=0.499 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.7 spell_power points (21.52 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -3.40 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -12.80 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (15.43 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -4.14 DPS) [quest]; Diana's Pearl Necklace (22403, -4.58 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.0 spell_power points (21.76 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -3.22 DPS) [pvp]; Burial Shawl (18681, -6.78 DPS) [dungeon]; Argent Shoulders (19059, -9.58 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.5 spell_power points (15.77 DPS) | yes | Amplifying Cloak (18350, -3.14 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.23 DPS, sim-verified) [world]; Hide of the Wild (18510, -5.35 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.8 spell_power points (32.80 DPS) | yes | Robe of Everlasting Night (18385, -6.45 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -8.93 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -14.08 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.7 spell_power points (15.91 DPS) | yes | Sublime Wristguards (18497, -6.90 DPS) [dungeon]; Runecloth Cuffs (254123, -7.60 DPS) [crafted]; Arcane Runed Bracers (4744, -9.59 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (19.24 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.80 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.7 spell_power points (22.91 DPS) | yes | Stormpike Cloth Girdle (19094, -9.69 DPS) [rep]; Ban'thok Sash (11662, -9.77 DPS) [dungeon]; Belt of the Archmage (18405, -10.46 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 38.3 spell_power points (26.90 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -2.58 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -6.49 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (16.84 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.40 DPS) [crafted]; Omnicast Boots (11822, -2.10 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.94 DPS) [quest]; Maiden's Circle (13001, -3.04 DPS) [world_drop]; Naglering (11669, -16.89 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.14 DPS) [quest]; Maiden's Circle (13001, -2.25 DPS) [world_drop]; Naglering (11669, -20.93 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+25.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -4.21 DPS) [quest]; Weakness Analyzer (272438, -4.91 DPS) [vendor]; Serenity Field (272439, -10.52 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -6.04 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.32 DPS) [world]; Teebu's Blazing Longsword (1728, -33.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (802.6 DPS) | yes | Bonecreeper Stylus (13938, -0.98 DPS) [dungeon]; Sparkling Crystal Wand (20672, -5.20 DPS) [world]; Torch of Light (279246, -31.14 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 38.2. Weights run: 2.6s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.067, intellect=0.082 ± 0.004, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.490 per %), hit=0.148 ± 0.005 per rating point (10 rating = 1%, 1.484 per %), spell_haste=not significant (0.199 ± 0.068), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.688 ± 0.067, fire_power=0.313 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.92 DPS) | yes | Shadow Goggles (4373, -3.17 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.7 spell_power points (0.88 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.27 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.30 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.61 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.21 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.4 spell_power points (0.83 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -1.24 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.2 spell_power points (0.18 DPS) | yes | Windsong Bangles (263336, -0.03 DPS) [quest]; Tabitha's Cuffs (251486, -0.10 DPS) [quest]; Featherbead Bracers (15452, -0.12 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.07 DPS) | yes | Gnoll Casting Gloves (892, -0.26 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Apothecary Gloves (10919, -0.46 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.3 spell_power points (0.66 DPS) | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Keller's Girdle (2911, -0.56 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.90 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.7 spell_power points (1.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.49 DPS) [crafted]; Silk-threaded Trousers (1929, -0.71 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.71 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.3 spell_power points (1.12 DPS) | yes | Red Woolen Boots (4313, -0.51 DPS) [crafted]; Feather Padded Treads (285345, -0.62 DPS, sim-verified) [world]; Pristine Boots (253889, -0.63 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.77 DPS) | yes | Lavishly Jeweled Ring (1156, -0.69 DPS) [dungeon]; Loop of Sacrifice (281673, -0.70 DPS) [quest]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.46 DPS) | yes | Loop of Sacrifice (281673, -0.40 DPS) [quest]; Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.77 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.8 spell_power points (0.13 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.05 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 148.5 spell_power points (22.77 DPS) | yes | Skycaller (12984, -1.73 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.48 DPS) [dungeon]; Sizzle Stick (8071, -4.35 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-0000000000000000000-2053225101000000)

Set DPS (verified): 70.9. Weights run: 2.7s. Verify run: 1.3s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.093, intellect=0.172 ± 0.006, crit=0.051 ± 0.002 per rating point (14 rating = 1%, 0.719 per %), hit=0.124 ± 0.006 per rating point (10 rating = 1%, 1.241 per %), spell_haste=not significant (-0.346 ± 0.090), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.803 ± 0.093, fire_power=0.195 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.93 DPS) | yes | Embalmed Shroud (7691, -0.80 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.80 DPS) [crafted]; Silk Headband (7050, -0.91 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (2.14 DPS) | yes | Crystal Starfire Medallion (5003, -1.95 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.95 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.19 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (2.81 DPS) | yes | Death Speaker Mantle (6685, -0.71 DPS) [dungeon]; Invoker's Mantle (215365, -0.72 DPS) [crafted]; Chestnut Mantle (17695, -1.26 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.33 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Battle Healer's Cloak (19529, -0.27 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.46 DPS) | yes | Green Silk Armor (7065, -0.84 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.09 DPS) [dungeon]; High Robe of the Adjudicator (3461, -1.24 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.40 DPS) | yes | Owlbeard Bracers (16981, -2.04 DPS) [quest]; Nightsky Wristbands (6407, -2.12 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.50 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.86 DPS) | yes | Gnoll Casting Gloves (892, -0.27 DPS) [world]; Truefaith Gloves (7049, -0.40 DPS) [crafted]; Jutebraid Gloves (10654, -0.52 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (3.07 DPS) | yes | Warsong Sash (16975, -0.36 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.53 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.94 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.19 DPS) | yes | Abomination Skin Leggings (23173, -0.61 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.01 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.32 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.2 spell_power points (2.18 DPS) | yes | Acidic Walkers (9454, -0.49 DPS) [dungeon]; Boots of the Enchanter (4325, -0.85 DPS) [crafted]; Spidersilk Boots (4320, -2.34 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.86 DPS) | yes | Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon]; Sludge-Stained Band (286535, -1.06 DPS) [world]; Sacred Band (6669, -1.33 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.60 DPS) | yes | Electrocutioner Lagnut (9447, -0.80 DPS) [dungeon]; Sacred Band (6669, -1.06 DPS) [quest]; Sludge-Stained Band (286535, -2.94 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.40 DPS) | yes | Glimmering Staff (249392, -1.79 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.94 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.94 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (2.14 DPS) | yes | Orb of Souls (249395, -1.07 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.42 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.53 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 127.8 spell_power points (34.01 DPS) | yes | Starfaller (13063, -0.83 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.68 DPS) [crafted]; Gravestone Scepter (7001, -5.01 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2053225103101330)

Set DPS (verified): 137.6. Weights run: 2.2s. Verify run: 1.2s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.269, intellect=0.494 ± 0.018, crit=0.097 ± 0.003 per rating point (14 rating = 1%, 1.353 per %), hit=0.293 ± 0.019 per rating point (10 rating = 1%, 2.931 per %), spell_haste=not significant (0.728 ± 0.240), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.660 ± 0.269), fire_power=0.341 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.58 DPS) | yes | Living Cowl (5608, -2.13 DPS) [world]; Holy Shroud (2721, -2.66 DPS) [world_drop]; Augural Shroud (2620, -2.76 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.0 spell_power points (2.65 DPS) | yes | Triune Amulet (7722, -1.73 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.73 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.09 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.4 spell_power points (3.57 DPS) | yes | Green Silken Shoulders (7057, -0.00 DPS) [crafted]; Inquisitor's Shawl (19507, -0.01 DPS) [dungeon]; Berylline Pads (4197, -0.40 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.4 spell_power points (3.57 DPS) | yes | Guardian Cloak (5965, -1.32 DPS) [crafted]; Icy Cloak (4327, -1.71 DPS) [crafted]; Long Silken Cloak (4326, -1.92 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.0 spell_power points (6.63 DPS) | yes | Dreamweave Vest (10021, -0.74 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.05 DPS) [world_drop]; Robe of Power (7054, -1.34 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.39 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.53 DPS) [quest]; Radiant Silver Bracers (4545, -1.25 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.0 spell_power points (5.31 DPS) | yes | Black Mageweave Gloves (10003, -1.32 DPS) [crafted]; Red Mageweave Gloves (10018, -1.72 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.26 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.0 spell_power points (4.24 DPS) | yes | Star Belt (4329, -0.79 DPS) [crafted]; Gilded Cord (254037, -1.07 DPS) [crafted]; Deathmage Sash (10771, -1.30 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.9 spell_power points (5.30 DPS) | yes | Crimson Silk Pantaloons (7062, -1.83 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.85 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.11 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.38 DPS) | yes | Gilded Slippers (254001, -2.65 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.99 DPS) [crafted]; Acidic Walkers (9454, -4.00 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (3.44 DPS) | yes | Reedknot Ring (9622, -1.58 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.85 DPS) [vendor]; Black Widow Band (6199, -2.53 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.39 DPS) | yes | Sea Giant's Toe Ring (274746, -0.80 DPS) [vendor]; Reedknot Ring (9622, -0.96 DPS, sim-verified) [quest]; Black Widow Band (6199, -1.47 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (137.6 DPS) | yes | Scorn's Focal Dagger (23168, -2.92 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.09 DPS) [quest]; Gut Ripper (2164, -8.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 151.2 spell_power points (40.17 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.57 DPS) [crafted]; Earthen Rod (9381, -4.58 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25000000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 195.8. Weights run: 8.7s. Verify run: 1.4s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.297, intellect=1.091 ± 0.020, crit=0.210 ± 0.004 per rating point (14 rating = 1%, 2.945 per %), hit=0.598 ± 0.020 per rating point (10 rating = 1%, 5.982 per %), spell_haste=2.334 ± 0.335, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.097 ± 0.297), fire_power=0.904 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.8 spell_power points (7.36 DPS) | yes | Dreamweave Circlet (10041, -1.61 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.94 DPS) [vendor]; Chief Architect's Monocle (11839, -2.05 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.18 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.47 DPS) [quest]; Horizon Choker (13085, -2.54 DPS, sim-verified) [world_drop] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Shoulders (10029, -0.87 DPS) [crafted]; Inquisitor's Shawl (19507, -1.26 DPS) [dungeon]; Rotgrip Mantle (17732, -2.29 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.8 spell_power points (3.93 DPS) | yes | Spritecaster Cape (11623, -0.23 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.54 DPS) [dungeon]; Runecloth Cloak (13860, -0.74 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.8 spell_power points (7.36 DPS) | yes | Runecloth Tunic (13857, -2.13 DPS) [crafted]; Robe of the Magi (1716, -2.21 DPS) [world_drop]; Runecloth Robe (13858, -2.67 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 16.4 spell_power points (2.95 DPS) | yes | Bloodband Bracers (11469, +0.00 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 31.0 spell_power points (5.59 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -0.89 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.56 DPS) [crafted]; Red Mageweave Gloves (10018, -1.64 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 30.0 spell_power points (5.40 DPS) | yes | Satyrmane Sash (17755, -0.91 DPS) [dungeon]; Deathmage Sash (10771, -1.19 DPS) [dungeon]; Dawnspire Cord (12466, -4.54 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 33.9 spell_power points (6.11 DPS) | yes | Red Mageweave Pants (10009, -1.23 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.03 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -5.38 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.33 DPS) | yes | Gilded Sandals (254107, -0.57 DPS) [crafted]; Southsea Mojo Boots (20641, -0.72 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -2.91 DPS, sim-verified) [vendor] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.6 spell_power points (3.00 DPS) | yes | Brainlash (6440, -0.05 DPS) [dungeon]; Mindseye Circle (10634, -0.64 DPS) [dungeon]; Band of the Unicorn (7553, -0.66 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.5 spell_power points (2.98 DPS) | yes | Mindseye Circle (10634, -0.62 DPS) [dungeon]; Band of the Unicorn (7553, -0.64 DPS) [world_drop]; Brainlash (6440, -3.13 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.22 DPS) [quest] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.41 DPS) [quest]; Uther's Strength (11302, -2.06 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Soul Harvester (20536, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -3.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.48 DPS) [world_drop]; Flaming Incinerator (9483, -3.68 DPS) [dungeon]; Pyric Caduceus (11748, -4.68 DPS, sim-verified) [dungeon] |

**New at 50:** head: Red Mageweave Headband; shoulder: Kentic Amice; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Abyss Shard; main_hand: Blade of Eternal Darkness; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532000000000000-0000000000000000000-2053225103101351)

Set DPS (verified): 422.8. Weights run: 2.3s. Verify run: 1.4s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.374, intellect=0.433 ± 0.025, crit=0.144 ± 0.004 per rating point (14 rating = 1%, 2.013 per %), hit=0.453 ± 0.026 per rating point (10 rating = 1%, 4.530 per %), spell_haste=-1.810 ± 0.408, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.602 ± 0.374), fire_power=0.397 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.5 spell_power points (16.74 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.33 DPS) [pvp]; Deathmist Mask (226909, -6.05 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (11.00 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.69 DPS) [quest]; Beads of Ogre Mojo (22149, -1.90 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.5 spell_power points (17.76 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.57 DPS) [pvp]; Burial Shawl (18681, -4.29 DPS) [dungeon]; Argent Shoulders (19059, -5.25 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.0 spell_power points (12.00 DPS) | yes | Crystalline Threaded Cape (20697, -1.13 DPS) [world]; Hide of the Wild (18510, -2.83 DPS) [crafted]; Amplifying Cloak (18350, -3.00 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.9 spell_power points (24.95 DPS) | yes | Robe of Everlasting Night (18385, -3.33 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -3.75 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -8.12 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.5 spell_power points (12.73 DPS) | yes | Sublime Wristguards (18497, -4.57 DPS) [dungeon]; Runecloth Cuffs (254123, -5.07 DPS) [crafted]; Deathmist Bracers (226907, -6.64 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.2 spell_power points (14.59 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -3.00 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.1 spell_power points (19.53 DPS) | yes | Belt of the Archmage (18405, -4.27 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -8.36 DPS) [rep]; Ban'thok Sash (11662, -8.88 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.4 spell_power points (20.18 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.07 DPS) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.2 spell_power points (12.60 DPS) | yes | Dragonrider Boots (18102, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.60 DPS) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.87 DPS) [quest]; Maiden's Circle (13001, -2.87 DPS) [world_drop]; Naglering (11669, -10.20 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.50 DPS) [quest]; Maiden's Circle (13001, -1.50 DPS) [world_drop]; Naglering (11669, -9.48 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+16.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -3.00 DPS) [quest]; Weakness Analyzer (272438, -3.50 DPS) [vendor]; Serenity Field (272439, -7.50 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.48 DPS) [world]; Teebu's Blazing Longsword (1728, -22.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (422.8 DPS) | yes | Bonecreeper Stylus (13938, -0.35 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.57 DPS) [world]; Torch of Light (279246, -23.72 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 792.5. Weights run: 7.1s. Verify run: 1.6s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.184, intellect=0.084 ± 0.023, crit=0.197 ± 0.006 per rating point (14 rating = 1%, 2.752 per %), hit=0.580 ± 0.037 per rating point (10 rating = 1%, 5.797 per %), spell_haste=not significant (0.037 ± 0.500), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.501 ± 0.184), fire_power=0.499 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.7 spell_power points (21.52 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -3.40 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -4.00 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (15.43 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -4.14 DPS) [quest]; Diana's Pearl Necklace (22403, -4.58 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.0 spell_power points (21.76 DPS) | yes | Warlord's Dreadweave Mantle (231592, -3.22 DPS) [pvp]; Argent Shoulders (19059, -4.22 DPS) [crafted]; Burial Shawl (18681, -6.78 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.5 spell_power points (15.77 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -3.14 DPS) [dungeon]; Hide of the Wild (18510, -5.35 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.8 spell_power points (32.80 DPS) | yes | Warlord's Dreadweave Robe (231591, -8.93 DPS) [pvp]; Robe of Everlasting Night (18385, -13.09 DPS) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -14.08 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.7 spell_power points (15.91 DPS) | yes | Sublime Wristguards (18497, -6.90 DPS) [dungeon]; Runecloth Cuffs (254123, -7.60 DPS) [crafted]; Spidertank Oilrag (9448, -9.59 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (19.24 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -0.64 DPS) [dungeon]; Earth Warder's Gloves (21318, -3.80 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.7 spell_power points (22.91 DPS) | yes | Belt of the Archmage (18405, -3.80 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -9.69 DPS) [rep]; Ban'thok Sash (11662, -9.77 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 38.3 spell_power points (26.90 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -2.58 DPS) [dungeon]; Outrider's Silk Leggings (22747, -6.13 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (16.84 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.40 DPS) [crafted]; Omnicast Boots (11822, -2.10 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.94 DPS) [quest]; Maiden's Circle (13001, -3.04 DPS) [world_drop]; Naglering (11669, -9.53 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.14 DPS) [quest]; Maiden's Circle (13001, -2.25 DPS) [world_drop]; Naglering (11669, -17.82 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -4.21 DPS) [quest]; Weakness Analyzer (272438, -4.91 DPS) [vendor]; Burst of Knowledge (11832, -10.15 DPS, sim-verified) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.32 DPS) [world]; Teebu's Blazing Longsword (1728, -28.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (792.5 DPS) | yes | Bonecreeper Stylus (13938, -0.98 DPS) [dungeon]; Sparkling Crystal Wand (20672, -5.20 DPS) [world]; Torch of Light (279246, -35.21 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

