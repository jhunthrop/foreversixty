# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 36.1. Weights run: 1.4s. Verify run: 1.3s. 129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.065, intellect=0.351 ± 0.004, crit=0.008 ± 0.001 per rating point (14 rating = 1%, 0.117 per %), hit=0.046 ± 0.001 per rating point (10 rating = 1%, 0.458 per %), spell_haste=not significant (-0.307 ± 0.080), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.065, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.28 DPS) | yes | Shadow Goggles (4373, -2.86 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.2 spell_power points (1.74 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.09 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.89 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.85 DPS) | yes | Feyscale Cloak (6632, -0.21 DPS) [dungeon]; Caretaker's Cape (20428, -0.21 DPS) [rep]; Pearl-clasped Cloak (5542, -0.67 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (1.44 DPS) | yes | Green Woolen Robe (6243, -0.58 DPS) [crafted]; Bloody Apron (6226, -0.59 DPS) [dungeon]; Gray Woolen Robe (2585, -1.47 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.1 spell_power points (0.45 DPS) | yes | Mindthrust Bracers (1974, -0.08 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.15 DPS) [world_drop]; Windsong Bangles (263336, -0.24 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.50 DPS) | yes | Gnoll Casting Gloves (892, -0.28 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.92 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (1.15 DPS) | yes | Novice Ardent's Sash (253887, -0.50 DPS) [crafted]; Keller's Girdle (2911, -0.55 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.10 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (36.1 DPS) | yes | Silk-threaded Trousers (1929, -0.24 DPS) [dungeon]; Rumpled Kilt (274741, -0.66 DPS) [vendor]; Abomination Skin Leggings (23173, -0.69 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (1.80 DPS) | yes | Feather Padded Treads (285345, -0.55 DPS, sim-verified) [world]; Pristine Boots (253889, -0.93 DPS) [crafted]; Red Woolen Boots (4313, -0.94 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (1.22 DPS) | yes | Sludge-Stained Band (286535, -0.58 DPS) [world]; Lavishly Jeweled Ring (1156, -0.77 DPS) [dungeon]; Loop of Sacrifice (281673, -0.84 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.07 DPS) | yes | Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; Loop of Sacrifice (281673, -0.69 DPS) [quest]; Sludge-Stained Band (286535, -0.77 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 3.5 spell_power points (0.75 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.30 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 107.4 spell_power points (22.95 DPS) | yes | Skycaller (12984, -1.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.66 DPS) [dungeon]; Sizzle Stick (8071, -4.23 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 129, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 55.8. Weights run: 1.3s. Verify run: 1.3s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.110, intellect=0.339 ± 0.007, crit=0.016 ± 0.001 per rating point (14 rating = 1%, 0.230 per %), hit=0.093 ± 0.001 per rating point (10 rating = 1%, 0.933 per %), spell_haste=not significant (0.034 ± 0.128), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.110, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.07 DPS) | yes | Silk Headband (7050, -0.38 DPS) [crafted]; Enchanter's Cowl (4322, -0.55 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.57 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Darkspear Warding Pendant (272075, -1.21 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.45 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.45 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (2.27 DPS) | yes | Death Speaker Mantle (6685, -0.40 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.57 DPS) [quest]; Invoker's Mantle (215365, -0.63 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.94 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Prelacy Cape (7004, -0.19 DPS) [quest]; Caretaker's Cape (19533, -0.19 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.4 spell_power points (2.53 DPS) | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.50 DPS) [dungeon]; Pristine Gown (253961, -0.76 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Nightsky Wristbands (6407, -1.31 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.31 DPS) [quest]; Glowing Magical Bracelets (13106, -1.73 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.7 spell_power points (1.46 DPS) | yes | Shilly Mitts (9609, +0.00 DPS, sim-verified) [quest]; Serpent Gloves (5970, -0.14 DPS) [dungeon]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.0 spell_power points (2.27 DPS) | yes | Belt of Arugal (6392, -0.20 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.63 DPS) [crafted]; Crimson Silk Belt (7055, -0.69 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.26 DPS) | yes | Abomination Skin Leggings (23173, -0.04 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.49 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.75 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (1.77 DPS) | yes | Acidic Walkers (9454, -0.31 DPS) [dungeon]; Nimbus Boots (6998, -0.64 DPS) [quest]; Spidersilk Boots (4320, -1.47 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.32 DPS) | yes | Minor Channeling Ring (1449, -0.25 DPS) [quest]; Lorekeeper's Ring (20431, -0.38 DPS) [rep]; Electrocutioner Lagnut (9447, -0.75 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.13 DPS) | yes | Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon]; Sludge-Stained Band (286535, -0.57 DPS) [world]; Minor Channeling Ring (1449, -1.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Twisted Chanter's Staff (890, -1.06 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.06 DPS) [quest]; Glimmering Staff (249392, -1.22 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.0 spell_power points (1.70 DPS) | yes | Dwarven Tome (279898, -0.71 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.95 DPS) [quest]; Orb of Souls (249395, -0.95 DPS) [crafted] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 179.2 spell_power points (33.78 DPS) | yes | Starfaller (13063, -0.13 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.83 DPS) [crafted]; Greater Mystic Wand (11290, -4.78 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 104.0. Weights run: 1.2s. Verify run: 1.1s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.209, intellect=0.485 ± 0.016, crit=0.043 ± 0.003 per rating point (14 rating = 1%, 0.608 per %), hit=0.210 ± 0.003 per rating point (10 rating = 1%, 2.097 per %), spell_haste=not significant (0.168 ± 0.262), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.209, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.54 DPS) | yes | Living Cowl (5608, -0.97 DPS) [world]; Holy Shroud (2721, -1.21 DPS) [world_drop]; Augural Shroud (2620, -2.27 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.9 spell_power points (1.20 DPS) | yes | Necklace of Calisea (1714, -0.79 DPS) [world_drop]; Triune Amulet (7722, -0.79 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.41 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.4 spell_power points (1.62 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.01 DPS) [dungeon]; Berylline Pads (4197, -0.18 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (104.0 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.17 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.99 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.9 spell_power points (3.02 DPS) | yes | Dreamweave Vest (10021, -0.17 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.47 DPS) [world_drop]; Robe of Power (7054, -0.62 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.09 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.24 DPS) [quest]; Windchaser Cuffs (14429, -0.56 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.9 spell_power points (2.42 DPS) | yes | Black Mageweave Gloves (10003, -0.60 DPS) [crafted]; Gilded Handwraps (254021, -1.03 DPS) [crafted]; Red Mageweave Gloves (10018, -1.07 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.9 spell_power points (1.93 DPS) | yes | Star Belt (4329, -0.36 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.42 DPS) [rep]; Deathmage Sash (10771, -1.43 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.8 spell_power points (2.40 DPS) | yes | Abomination Skin Leggings (23173, -0.84 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.95 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.50 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.91 DPS) | yes | Spidersilk Boots (4320, -1.82 DPS) [crafted]; Acidic Walkers (9454, -1.83 DPS) [dungeon]; Gilded Slippers (254001, -2.23 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (1.56 DPS) | yes | Ring of Forlorn Spirits (2043, -0.60 DPS) [quest]; Reedknot Ring (9622, -0.72 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.09 DPS) | yes | Reedknot Ring (9622, -0.24 DPS) [quest]; Lorekeeper's Ring (19525, -0.24 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.30 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Scorn's Focal Dagger (23168, -1.33 DPS) [dungeon]; Staff of Dar'Orahil (15106, -1.52 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 328.0 spell_power points (39.73 DPS) | yes | Umbral Wand (5216, -0.33 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.14 DPS) [dungeon]; Twisted Nether Wand (249144, -5.01 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 141.9. Weights run: 1.2s. Verify run: 1.0s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.208, intellect=0.255 ± 0.013, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.535 per %), hit=0.149 ± 0.002 per rating point (10 rating = 1%, 1.493 per %), spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -0.97 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.96 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.28 DPS) [quest]; Horizon Choker (13085, -1.39 DPS) [world_drop]; Arcane Crystal Pendant (20037, -2.00 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.6 spell_power points (4.93 DPS) | yes | Black Mageweave Shoulders (10027, -1.48 DPS) [crafted]; Bloodmage Mantle (7684, -1.76 DPS) [dungeon]; Kentic Amice (11624, -3.66 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (4.35 DPS) | yes | Runecloth Cloak (13860, -1.26 DPS) [crafted]; Nightfall Drape (12465, -1.83 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -3.93 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.1 spell_power points (6.75 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.87 DPS) [world_drop]; Dreamweave Vest (10021, -1.07 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, -0.06 DPS) [crafted]; Spidertank Oilrag (9448, -0.08 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -0.48 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (5.33 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -0.60 DPS, sim-verified) [vendor]; Black Mageweave Gloves (10003, -1.13 DPS) [crafted]; Runecloth Gloves (13863, -1.32 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (4.64 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.45 DPS) [dungeon]; Ghostweave Cord (254073, -0.71 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.54 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -0.95 DPS) [vendor]; Spellshock Leggings (9484, -5.04 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -0.36 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.14 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.42 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.41 DPS) [quest]; Cyclopean Band (11824, -0.62 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.36 DPS) | yes | Philanthropist's Ring (281635, -0.20 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.34 DPS) [dungeon]; Lorekeeper's Ring (19524, -0.84 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, -1.64 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Moonshadow Stave (22458, -0.55 DPS) [quest]; Arbiter's Blade (11784, -3.01 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Woestave (20082, -0.08 DPS) [quest]; Pyric Caduceus (11748, -2.30 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.98 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 318.0. Weights run: 1.3s. Verify run: 1.1s. 982 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.101, intellect=0.163 ± 0.008, crit=0.014 ± 0.001 per rating point (14 rating = 1%, 0.198 per %), hit=0.052 ± 0.001 per rating point (10 rating = 1%, 0.518 per %), spell_haste=0.740 ± 0.119, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.986 ± 0.101, fire_power=0.014 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 65.3 spell_power points (72.76 DPS) | yes | Crimson Felt Hat (18727, -25.00 DPS, sim-verified) [dungeon]; Field Marshal's Coronal (17578, -32.79 DPS) [vendor]; Field Marshal's Coronal (231584, -32.79 DPS) [pvp] |
| neck | Orb of the Darkmoon (19426) (or Chains of the Lich (23125)) | 1200 Tickets - Orb of the Darkmoon [quest] | 22.0 spell_power points (24.50 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Amulet of the Dawn (22657, -5.44 DPS) [quest]; Arcane Crystal Pendant (20037, -5.59 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 42.8 spell_power points (47.69 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -6.68 DPS, sim-verified) [vendor]; Field Marshal's Dreadweave Shoulders (17580, -16.77 DPS) [vendor]; Field Marshal's Dreadweave Shoulders (231583, -16.77 DPS) [pvp] |
| back | Crystalline Threaded Cape (20697) | World drop [world_drop] | 20.7 spell_power points (22.99 DPS) | yes | Amplifying Cloak (18350, +0.00 DPS, sim-verified) [dungeon]; Arcanoweave Cloak (272411, -3.15 DPS) [vendor]; Hide of the Wild (18510, -5.59 DPS) [crafted] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 68.2 spell_power points (75.89 DPS) | yes | Robe of the Void (14153, -7.64 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (17581, -35.92 DPS) [vendor]; Field Marshal's Dreadweave Robe (231582, -35.92 DPS) [pvp] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 28.6 spell_power points (31.88 DPS) | yes | Dryad's Wrist Bindings (19595, -3.19 DPS, sim-verified) [rep]; Dryad's Wrist Bindings (19596, -8.52 DPS) [rep]; Dryad's Wrist Bindings (19597, -12.98 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 39.6 spell_power points (44.13 DPS) | yes | Sandworm Skin Gloves (20716, -5.56 DPS, sim-verified) [quest]; Marshal's Dreadweave Gloves (17584, -9.64 DPS) [vendor]; Marshal's Dreadweave Gloves (231586, -9.64 DPS) [pvp] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 43.0 spell_power points (47.91 DPS) | yes | Knowledge of the Timbermaw (228190, -6.13 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -22.52 DPS) [crafted]; Felheart Belt (16806, -22.92 DPS) [world_drop] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 58.7 spell_power points (65.33 DPS) | yes | Flarecore Leggings (19165, -9.04 DPS, sim-verified) [crafted]; Marshal's Dreadweave Leggings (17579, -20.70 DPS) [vendor]; Marshal's Dreadweave Leggings (231587, -20.70 DPS) [pvp] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 38.8 spell_power points (43.19 DPS) | yes | Earthen Silk Slippers (254013, -6.14 DPS, sim-verified) [crafted]; Marshal's Dreadweave Boots (17583, -11.89 DPS) [vendor]; Marshal's Dreadweave Boots (231585, -11.89 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | sim-verified (318.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (234437, -1.29 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234032, -1.65 DPS) [vendor]; Naglering (11669, -9.61 DPS, sim-verified) [dungeon] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | sim-verified (318.0 DPS) | yes | Elemental Focus Band (20682, -1.57 DPS) [world]; Maiden's Circle (13001, -6.75 DPS) [world_drop]; Naglering (11669, -9.02 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (318.0 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (318.0 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, -2.69 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (318.0 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Electrified Dagger (19100, +0.00 DPS, sim-verified) [rep]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 74.9 spell_power points (83.44 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -6.34 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -7.84 DPS) [dungeon]; Brilliant Wand (249385, -12.57 DPS) [crafted] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Crystalline Threaded Cape; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ritssyn's Ring of Chaos; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 982, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 34.9. Weights run: 1.4s. Verify run: 1.2s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.065, intellect=0.351 ± 0.004, crit=0.008 ± 0.001 per rating point (14 rating = 1%, 0.117 per %), hit=0.046 ± 0.001 per rating point (10 rating = 1%, 0.458 per %), spell_haste=not significant (-0.307 ± 0.080), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.065, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.28 DPS) | yes | Shadow Goggles (4373, -2.72 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.2 spell_power points (1.74 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.02 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.89 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.85 DPS) | yes | Feyscale Cloak (6632, -0.21 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.21 DPS) [rep]; Pearl-clasped Cloak (5542, -0.44 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (1.44 DPS) | yes | Green Woolen Robe (6243, -0.58 DPS) [crafted]; Bloody Apron (6226, -0.59 DPS) [dungeon]; Gray Woolen Robe (2585, -1.37 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.1 spell_power points (0.45 DPS) | yes | Featherbead Bracers (15452, -0.07 DPS) [quest]; Mindthrust Bracers (1974, -0.08 DPS, sim-verified) [dungeon]; Owlbeard Bracers (16981, -0.09 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.50 DPS) | yes | Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.42 DPS) [crafted]; Apothecary Gloves (10919, -0.64 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (1.15 DPS) | yes | Novice Ardent's Sash (253887, -0.50 DPS) [crafted]; Keller's Girdle (2911, -0.55 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.92 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (34.9 DPS) | yes | Silk-threaded Trousers (1929, -0.24 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.66 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.66 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (1.80 DPS) | yes | Feather Padded Treads (285345, -0.67 DPS, sim-verified) [world]; Pristine Boots (253889, -0.93 DPS) [crafted]; Red Woolen Boots (4313, -0.94 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.07 DPS) | yes | Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; Loop of Sacrifice (281673, -0.69 DPS) [quest]; Volcanic Rock Ring (12053, -0.84 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.64 DPS) | yes | Loop of Sacrifice (281673, -0.27 DPS) [quest]; Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.60 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 3.5 spell_power points (0.75 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.30 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 107.4 spell_power points (22.95 DPS) | yes | Skycaller (12984, -1.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.66 DPS) [dungeon]; Sizzle Stick (8071, -4.23 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 54.2. Weights run: 1.3s. Verify run: 1.2s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.110, intellect=0.339 ± 0.007, crit=0.016 ± 0.001 per rating point (14 rating = 1%, 0.230 per %), hit=0.093 ± 0.001 per rating point (10 rating = 1%, 0.933 per %), spell_haste=not significant (0.034 ± 0.128), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.110, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.07 DPS) | yes | Silk Headband (7050, -0.38 DPS) [crafted]; Embalmed Shroud (7691, -0.57 DPS) [dungeon]; Enchanter's Cowl (4322, -0.72 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Crystal Starfire Medallion (5003, -1.45 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.45 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.49 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (2.27 DPS) | yes | Fairywing Mantle (9536, -0.57 DPS) [quest]; Invoker's Mantle (215365, -0.63 DPS) [crafted]; Death Speaker Mantle (6685, -0.68 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.94 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Battle Healer's Cloak (19529, -0.19 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.4 spell_power points (2.53 DPS) | yes | Tree Bark Jacket (1486, -0.09 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.50 DPS) [dungeon]; Pristine Gown (253961, -0.76 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Nightsky Wristbands (6407, -1.31 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.31 DPS) [quest]; Glowing Magical Bracelets (13106, -1.52 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.7 spell_power points (1.45 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.32 DPS) [crafted]; Gnoll Casting Gloves (892, -0.32 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (2.27 DPS) | yes | Warsong Sash (16975, -0.27 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.38 DPS) [dungeon]; Invoker's Cord (215366, -0.63 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.26 DPS) | yes | Abomination Skin Leggings (23173, -0.17 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.49 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.75 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (1.77 DPS) | yes | Acidic Walkers (9454, -0.31 DPS) [dungeon]; Boots of the Enchanter (4325, -0.82 DPS) [crafted]; Spidersilk Boots (4320, -1.79 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.32 DPS) | yes | Advisor's Ring (20426, -0.38 DPS) [rep]; Electrocutioner Lagnut (9447, -0.75 DPS) [dungeon]; Sludge-Stained Band (286535, -0.75 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.13 DPS) | yes | Sludge-Stained Band (286535, -0.57 DPS) [world]; Snake Hoop (6750, -0.68 DPS) [quest]; Electrocutioner Lagnut (9447, -2.09 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Twisted Chanter's Staff (890, -1.06 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.06 DPS) [quest]; Glimmering Staff (249392, -1.25 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.0 spell_power points (1.70 DPS) | yes | Dwarven Tome (279898, -0.87 DPS, sim-verified) [quest]; Orb of Souls (249395, -0.95 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -0.95 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 179.2 spell_power points (33.78 DPS) | yes | Starfaller (13063, -0.31 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.83 DPS) [crafted]; Greater Mystic Wand (11290, -4.78 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 102.1. Weights run: 1.2s. Verify run: 1.1s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.209, intellect=0.485 ± 0.016, crit=0.043 ± 0.003 per rating point (14 rating = 1%, 0.608 per %), hit=0.210 ± 0.003 per rating point (10 rating = 1%, 2.097 per %), spell_haste=not significant (0.168 ± 0.262), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.209, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.54 DPS) | yes | Living Cowl (5608, -0.97 DPS) [world]; Holy Shroud (2721, -1.21 DPS) [world_drop]; Augural Shroud (2620, -2.04 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.9 spell_power points (1.20 DPS) | yes | Necklace of Calisea (1714, -0.79 DPS) [world_drop]; Triune Amulet (7722, -0.79 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.12 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.4 spell_power points (1.62 DPS) | yes | Inquisitor's Shawl (19507, -0.01 DPS) [dungeon]; Green Silken Shoulders (7057, -0.15 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.18 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (102.1 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.17 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.68 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.9 spell_power points (3.02 DPS) | yes | Elemental Raiment (9434, -0.47 DPS) [world_drop]; Robe of Power (7054, -0.62 DPS) [crafted]; Dreamweave Vest (10021, -0.91 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.09 DPS) | yes | Condor Bracers (15864, -0.24 DPS) [quest]; Windchaser Cuffs (14429, -0.56 DPS) [world_drop]; Radiant Silver Bracers (4545, -0.86 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.9 spell_power points (2.42 DPS) | yes | Black Mageweave Gloves (10003, -0.60 DPS) [crafted]; Gilded Handwraps (254021, -1.03 DPS) [crafted]; Red Mageweave Gloves (10018, -1.72 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.9 spell_power points (1.93 DPS) | yes | Star Belt (4329, -0.36 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.42 DPS) [rep]; Deathmage Sash (10771, -1.11 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.8 spell_power points (2.40 DPS) | yes | Abomination Skin Leggings (23173, -0.84 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.95 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.09 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.91 DPS) | yes | Spidersilk Boots (4320, -1.82 DPS) [crafted]; Acidic Walkers (9454, -1.83 DPS) [dungeon]; Gilded Slippers (254001, -1.95 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (1.56 DPS) | yes | Reedknot Ring (9622, -0.72 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor]; Voodoo Band (1996, -1.15 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.09 DPS) | yes | Advisor's Ring (19521, -0.24 DPS) [rep]; Reedknot Ring (9622, -0.29 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.36 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Scorn's Focal Dagger (23168, -1.33 DPS) [dungeon]; Staff of Dar'Orahil (15106, -1.52 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 328.0 spell_power points (39.73 DPS) | yes | Umbral Wand (5216, -0.22 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.14 DPS) [dungeon]; Twisted Nether Wand (249144, -5.01 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 139.9. Weights run: 1.2s. Verify run: 1.1s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.208, intellect=0.255 ± 0.013, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.535 per %), hit=0.149 ± 0.002 per rating point (10 rating = 1%, 1.493 per %), spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -0.97 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.78 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.28 DPS) [quest]; Horizon Choker (13085, -1.39 DPS) [world_drop]; Arcane Crystal Pendant (20037, -1.50 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.6 spell_power points (4.93 DPS) | yes | Black Mageweave Shoulders (10027, -1.48 DPS) [crafted]; Bloodmage Mantle (7684, -1.76 DPS) [dungeon]; Kentic Amice (11624, -4.65 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (4.35 DPS) | yes | Deep Woodlands Cloak (19121, -0.61 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.19 DPS) [dungeon]; Runecloth Cloak (13860, -1.26 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.1 spell_power points (6.75 DPS) | yes | Robe of the Magi (1716, -0.66 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.87 DPS) [world_drop]; Dreamweave Vest (10021, -1.07 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -0.48 DPS) [quest]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (5.33 DPS) | yes | Black Mageweave Gloves (10003, -1.13 DPS) [crafted]; Runecloth Gloves (13863, -1.32 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.46 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (4.64 DPS) | yes | Ban'thok Sash (11662, -0.45 DPS) [dungeon]; Ghostweave Cord (254073, -0.71 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.25 DPS, sim-verified) [rep] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.54 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -0.95 DPS) [vendor]; Spellshock Leggings (9484, -4.56 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -0.60 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.14 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.42 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.41 DPS) [quest]; Cyclopean Band (11824, -0.62 DPS) [dungeon]; Runed Ring (862, -1.68 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.36 DPS) | yes | Cyclopean Band (11824, -0.34 DPS) [dungeon]; Philanthropist's Ring (281635, -0.68 DPS, sim-verified) [quest]; Advisor's Ring (19520, -0.84 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.99 DPS) [crafted]; Rune of the Guard Captain (19120, -3.07 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -1.30 DPS, sim-verified) [crafted]; Rune of the Guard Captain (19120, -1.39 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Moonshadow Stave (22458, -0.55 DPS) [quest]; Arbiter's Blade (11784, -3.01 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Woestave (20082, -0.08 DPS) [quest]; Pyric Caduceus (11748, -2.19 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.98 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 313.1. Weights run: 1.3s. Verify run: 1.1s. 976 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.101, intellect=0.163 ± 0.008, crit=0.014 ± 0.001 per rating point (14 rating = 1%, 0.198 per %), hit=0.052 ± 0.001 per rating point (10 rating = 1%, 0.518 per %), spell_haste=0.740 ± 0.119, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.986 ± 0.101, fire_power=0.014 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 65.3 spell_power points (72.76 DPS) | yes | Crimson Felt Hat (18727, -24.16 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Hood (17591, -32.79 DPS) [vendor]; Warlord's Dreadweave Hood (231590, -32.79 DPS) [pvp] |
| neck | Orb of the Darkmoon (19426) (or Chains of the Lich (23125)) | 1200 Tickets - Orb of the Darkmoon [quest] | 22.0 spell_power points (24.50 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Amulet of the Dawn (22657, -5.44 DPS) [quest]; Arcane Crystal Pendant (20037, -5.59 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 42.8 spell_power points (47.69 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -5.98 DPS, sim-verified) [vendor]; Warlord's Dreadweave Mantle (17590, -16.77 DPS) [vendor]; Warlord's Dreadweave Mantle (231592, -16.77 DPS) [pvp] |
| back | Crystalline Threaded Cape (20697) | World drop [world_drop] | 20.7 spell_power points (22.99 DPS) | yes | Amplifying Cloak (18350, +0.00 DPS, sim-verified) [dungeon]; Arcanoweave Cloak (272411, -3.15 DPS) [vendor]; Hide of the Wild (18510, -5.59 DPS) [crafted] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 68.2 spell_power points (75.89 DPS) | yes | Robe of the Void (14153, -7.92 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (17592, -35.92 DPS) [vendor]; Warlord's Dreadweave Robe (231591, -35.92 DPS) [pvp] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 28.6 spell_power points (31.88 DPS) | yes | Dryad's Wrist Bindings (19595, -3.52 DPS, sim-verified) [rep]; Dryad's Wrist Bindings (19596, -8.52 DPS) [rep]; Dryad's Wrist Bindings (19597, -12.98 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 39.6 spell_power points (44.13 DPS) | yes | Sandworm Skin Gloves (20716, -5.09 DPS, sim-verified) [quest]; General's Dreadweave Gloves (17588, -9.64 DPS) [vendor]; General's Dreadweave Gloves (231589, -9.64 DPS) [pvp] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 43.0 spell_power points (47.91 DPS) | yes | Knowledge of the Timbermaw (228190, -6.14 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -22.52 DPS) [crafted]; Felheart Belt (16806, -22.92 DPS) [world_drop] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 58.7 spell_power points (65.33 DPS) | yes | Flarecore Leggings (19165, -10.52 DPS, sim-verified) [crafted]; General's Dreadweave Pants (17593, -20.70 DPS) [vendor]; General's Dreadweave Pants (231588, -20.70 DPS) [pvp] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 38.8 spell_power points (43.19 DPS) | yes | Earthen Silk Slippers (254013, -4.77 DPS, sim-verified) [crafted]; General's Dreadweave Boots (17586, -11.89 DPS) [vendor]; General's Dreadweave Boots (231593, -11.89 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | sim-verified (313.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (234437, -1.29 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234032, -1.65 DPS) [vendor]; Naglering (11669, -8.44 DPS, sim-verified) [dungeon] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | sim-verified (313.1 DPS) | yes | Elemental Focus Band (20682, -1.57 DPS) [world]; Maiden's Circle (13001, -6.75 DPS) [world_drop]; Naglering (11669, -8.23 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (313.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (313.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, -1.86 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (313.1 DPS) | yes | Glacial Blade (19099, +0.00 DPS, sim-verified) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -4.22 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 74.9 spell_power points (83.44 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -5.16 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -7.84 DPS) [dungeon]; Brilliant Wand (249385, -12.57 DPS) [crafted] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Crystalline Threaded Cape; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ritssyn's Ring of Chaos; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 976, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

