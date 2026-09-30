# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.1. Weights run: 1.0s. Verify run: 0.8s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.071, intellect=0.125 ± 0.010, crit=0.021 ± 0.001 per rating point (14 rating = 1%, 0.289 per %), hit=0.077 ± 0.001 per rating point (10 rating = 1%, 0.774 per %), spell_haste=0.435 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.071

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.43 DPS) | yes | Shadow Goggles (4373, -2.59 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.1 spell_power points (1.45 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.50 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.50 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.95 DPS) | yes | Feyscale Cloak (6632, -0.24 DPS) [dungeon]; Caretaker's Cape (20428, -0.24 DPS) [rep]; Black Whelp Cloak (7283, -0.27 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 spell_power points (1.34 DPS) | yes | Green Woolen Vest (2582, -0.39 DPS) [crafted]; Bloody Apron (6226, -0.39 DPS) [dungeon]; Gray Woolen Robe (2585, -1.22 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.24 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Repurposed Hair Band (281256, -0.18 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.66 DPS) | yes | Gnoll Casting Gloves (892, -0.31 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.62 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.13 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 spell_power points (1.07 DPS) | yes | Novice Ardent's Sash (253887, -0.50 DPS) [crafted]; Keller's Girdle (2911, -0.83 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.07 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.0 spell_power points (2.38 DPS) | yes | Silk-threaded Trousers (1929, -0.68 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.77 DPS) [crafted]; Rumpled Kilt (274741, -1.19 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (1.78 DPS) | yes | Feather Padded Treads (285345, -0.68 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.83 DPS) [crafted]; Pristine Boots (253889, -0.98 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.2 spell_power points (1.25 DPS) | yes | Sludge-Stained Band (286535, -0.53 DPS) [world]; Lavishly Jeweled Ring (1156, -1.07 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.16 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.19 DPS) | yes | Sludge-Stained Band (286535, -0.56 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -1.01 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.10 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 1.2 spell_power points (0.30 DPS) | yes | Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Staff of Westfall (2042, -0.15 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 96.9 spell_power points (23.02 DPS) | yes | Skycaller (12984, -1.32 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.73 DPS) [dungeon]; Sizzle Stick (8071, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 59.9. Weights run: 0.9s. Verify run: 0.8s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.156, intellect=0.346 ± 0.011, crit=0.034 ± 0.002 per rating point (14 rating = 1%, 0.478 per %), hit=0.134 ± 0.001 per rating point (10 rating = 1%, 1.336 per %), spell_haste=not significant (0.301 ± 0.194), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.156

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.85 DPS) | yes | Silk Headband (7050, -0.34 DPS) [crafted]; Embalmed Shroud (7691, -0.51 DPS) [dungeon]; Enchanter's Cowl (4322, -0.84 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.1 spell_power points (1.53 DPS) | yes | Crystal Starfire Medallion (5003, -1.30 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.30 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.15 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (2.04 DPS) | yes | Fairywing Mantle (9536, -0.51 DPS) [quest]; Invoker's Mantle (215365, -0.57 DPS) [crafted]; Death Speaker Mantle (6685, -1.07 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.84 DPS) | yes | Heavy Woolen Cloak (4311, -0.17 DPS) [crafted]; Prelacy Cape (7004, -0.17 DPS) [quest]; Repairman's Cape (9605, -0.78 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.5 spell_power points (2.27 DPS) | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.45 DPS) [dungeon]; Pristine Gown (253961, -0.69 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.52 DPS) | yes | Nightsky Wristbands (6407, -1.17 DPS) [world_drop]; Stonecloth Bindings (14416, -1.22 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.37 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.8 spell_power points (1.31 DPS) | yes | Serpent Gloves (5970, -0.14 DPS) [dungeon]; Truefaith Gloves (7049, -0.30 DPS) [crafted]; Shilly Mitts (9609, -0.33 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.0 spell_power points (2.03 DPS) | yes | Invoker's Cord (215366, -0.56 DPS) [crafted]; Crimson Silk Belt (7055, -0.61 DPS) [crafted]; Belt of Arugal (6392, -0.86 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.02 DPS) | yes | Pristine Leggings (253987, -0.43 DPS) [crafted]; Abomination Skin Leggings (23173, -0.47 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (1.59 DPS) | yes | Acidic Walkers (9454, -0.28 DPS) [dungeon]; Nimbus Boots (6998, -0.58 DPS) [quest]; Spidersilk Boots (4320, -2.21 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.18 DPS) | yes | Minor Channeling Ring (1449, -0.22 DPS) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Sludge-Stained Band (286535, -0.67 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.01 DPS) | yes | Electrocutioner Lagnut (9447, -0.51 DPS) [dungeon]; Sludge-Stained Band (286535, -0.51 DPS) [world]; Minor Channeling Ring (1449, -1.67 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.52 DPS) | yes | Twisted Chanter's Staff (890, -0.93 DPS) [world_drop]; Channeler's Staff (4437, -1.05 DPS) [world]; Glimmering Staff (249392, -2.23 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.1 spell_power points (1.53 DPS) | yes | Eye of Paleth (2943, -0.85 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.85 DPS) [world]; Dwarven Tome (279898, -1.37 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 200.2 spell_power points (33.72 DPS) | yes | Starfaller (13063, -0.61 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.87 DPS) [crafted]; Gravestone Scepter (7001, -4.72 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 118.0. Weights run: 0.8s. Verify run: 0.7s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.127, intellect=0.281 ± 0.009, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.387 per %), hit=0.103 ± 0.001 per rating point (10 rating = 1%, 1.034 per %), spell_haste=1.008 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.127

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.88 DPS) | yes | Living Cowl (5608, -2.24 DPS) [world]; Augural Shroud (2620, -2.37 DPS, sim-verified) [world]; Holy Shroud (2721, -2.80 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (2.43 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.59 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.88 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.88 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.5 spell_power points (3.23 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.25 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 11.5 spell_power points (3.23 DPS) | yes | Guardian Cloak (5965, -1.15 DPS) [crafted]; Icy Cloak (4327, -1.27 DPS) [crafted]; Long Silken Cloak (4326, -1.65 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.7 spell_power points (6.63 DPS) | yes | Elemental Raiment (9434, -0.33 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.88 DPS) [crafted]; Robe of Power (7054, -1.77 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest]; Earthen Silk Cuffs (254019, -1.40 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (5.36 DPS) | yes | Black Mageweave Gloves (10003, -1.03 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.49 DPS) [crafted]; Gilded Handwraps (254021, -2.56 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.1 spell_power points (4.24 DPS) | yes | Star Belt (4329, -0.23 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.10 DPS) [dungeon]; Gilded Cord (254037, -1.37 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.4 spell_power points (4.86 DPS) | yes | Gaze Dreamer Pants (6903, -0.47 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.60 DPS) [crafted]; Abomination Skin Leggings (23173, -1.71 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.72 DPS) | yes | Gilded Slippers (254001, -2.06 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.45 DPS) [crafted]; Acidic Walkers (9454, -4.69 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (3.27 DPS) | yes | Ring of Forlorn Spirits (2043, -1.03 DPS) [quest]; Reedknot Ring (9622, -1.31 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.59 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.52 DPS) | yes | Ring of Forlorn Spirits (2043, -0.23 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.56 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Scorn's Focal Dagger (23168, -3.08 DPS) [dungeon]; Windweaver Staff (7757, -4.42 DPS) [dungeon]; Gut Ripper (2164, -7.30 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (118.0 DPS) | yes | Umbral Wand (5216, -0.01 DPS) [dungeon]; Earthen Rod (9381, -0.09 DPS) [dungeon]; Jaina's Firestarter (13064, -1.84 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 205.0. Weights run: 1.0s. Verify run: 0.9s. 415 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.227, intellect=0.211 ± 0.015, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.457 per %), hit=0.203 ± 0.003 per rating point (10 rating = 1%, 2.027 per %), spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -2.16 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.3 spell_power points (2.32 DPS) | yes | Mindburst Medallion (11196, -0.38 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.49 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.73 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -1.36 DPS) [crafted]; Bloodmage Mantle (7684, -1.64 DPS) [dungeon]; Rotgrip Mantle (17732, -2.44 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.28 DPS) | yes | Mantle of Lady Falther'ess (23178, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.28 DPS) [crafted]; Nightfall Drape (12465, -1.76 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (6.52 DPS) | yes | Acumen Robes (17775, +0.00 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.64 DPS) [world_drop]; Dreamweave Vest (10021, -0.94 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.15 DPS) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.28 DPS) | yes | Black Mageweave Gloves (10003, -0.25 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.11 DPS) [vendor]; Runecloth Gloves (13863, -1.39 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (4.52 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Ghostweave Cord (254073, -0.59 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.1 spell_power points (7.04 DPS) | yes | Red Mageweave Pants (10009, -2.40 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.84 DPS) [vendor]; Wizardweave Leggings (14132, -4.03 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.23 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.39 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.49 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.36 DPS) | yes | Philanthropist's Ring (281635, +0.00 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.43 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.12 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.85 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -1.09 DPS, sim-verified) [crafted] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Arbiter's Blade (11784, -3.07 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.08 DPS) [dungeon]; Blade of Eternal Darkness (17780, -7.98 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted]; Pyric Caduceus (11748, -5.08 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Noxious Shooter

No-known-source sample (15 of 415, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 424.9. Weights run: 0.9s. Verify run: 0.9s. 1022 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.180, intellect=0.192 ± 0.014, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), hit=0.172 ± 0.002 per rating point (10 rating = 1%, 1.719 per %), spell_haste=-7.240 ± 0.381, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.180

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 68.5 spell_power points (37.77 DPS) | yes | Field Marshal's Coronal (17578, -16.19 DPS, sim-verified) [vendor]; Crimson Felt Hat (18727, -20.38 DPS) [dungeon] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (12.13 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS, sim-verified) [quest]; Amulet of the Dawn (22657, -2.48 DPS) [quest]; Kezan's Taint (19604, -3.56 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 44.8 spell_power points (24.69 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -7.42 DPS, sim-verified) [vendor]; Field Marshal's Dreadweave Shoulders (231583, -9.11 DPS) [pvp] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.8 spell_power points (11.45 DPS) | yes | Amplifying Cloak (18350, -1.53 DPS) [dungeon]; Hide of the Wild (18510, -2.67 DPS) [crafted]; Arcanoweave Cloak (272411, -3.43 DPS, sim-verified) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 69.4 spell_power points (38.27 DPS) | yes | Robe of the Void (14153, -11.30 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -18.09 DPS) [pvp] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 30.2 spell_power points (16.66 DPS) | yes | Dryad's Wrist Bindings (19596, -5.00 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 41.2 spell_power points (22.72 DPS) | yes | Marshal's Dreadweave Gloves (231586, -5.55 DPS) [pvp]; Sandworm Skin Gloves (20716, -7.31 DPS) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 45.4 spell_power points (25.01 DPS) | yes | Knowledge of the Timbermaw (228190, -8.29 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -11.97 DPS) [crafted]; Heretic Waistguard (240151, -13.24 DPS) [vendor] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 61.1 spell_power points (33.71 DPS) | yes | Marshal's Dreadweave Leggings (17579, -10.89 DPS, sim-verified) [vendor]; Skyshroud Leggings (13170, -14.12 DPS) [dungeon] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 40.4 spell_power points (22.27 DPS) | yes | Marshal's Dreadweave Boots (231585, -6.57 DPS) [pvp]; Earthen Silk Slippers (254013, -9.05 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Songstone of Ironforge (12543, -6.83 DPS) [quest]; Maiden's Circle (13001, -6.83 DPS) [world_drop]; Naglering (11669, -12.25 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 0.0 spell_power points (0.00 DPS) | yes | Songstone of Ironforge (12543, -2.63 DPS) [quest]; Maiden's Circle (13001, -2.63 DPS) [world_drop]; Naglering (11669, -8.95 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+16.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Weakness Analyzer (272438, -3.86 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -4.31 DPS, sim-verified) [quest]; Serenity Field (272439, -8.27 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | 0.0 spell_power points (0.00 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -2.00 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -17.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (424.9 DPS) | yes | Bonecreeper Stylus (13938, -0.80 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.87 DPS) [world]; Torch of Light (279246, -5.53 DPS, sim-verified) [crafted] |

**New at 60:** head: Heretic Cowl; neck: Chains of the Lich; shoulder: Heretic Shoulderpads; back: Crystalline Threaded Cape; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1022, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.5. Weights run: 1.0s. Verify run: 0.8s. 142 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.071, intellect=0.125 ± 0.010, crit=0.021 ± 0.001 per rating point (14 rating = 1%, 0.289 per %), hit=0.077 ± 0.001 per rating point (10 rating = 1%, 0.774 per %), spell_haste=0.435 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.071

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.43 DPS) | yes | Shadow Goggles (4373, -2.17 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.1 spell_power points (1.45 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.19 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.50 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.95 DPS) | yes | Black Whelp Cloak (7283, -0.15 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.24 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.24 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 spell_power points (1.34 DPS) | yes | Green Woolen Vest (2582, -0.39 DPS) [crafted]; Bloody Apron (6226, -0.39 DPS) [dungeon]; Gray Woolen Robe (2585, -0.88 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.2 spell_power points (0.30 DPS) | yes | Windsong Bangles (263336, -0.10 DPS, sim-verified) [quest]; Tabitha's Cuffs (251486, -0.12 DPS) [quest]; Featherbead Bracers (15452, -0.15 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.66 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.62 DPS) [crafted]; Apothecary Gloves (10919, -0.71 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 spell_power points (1.07 DPS) | yes | Novice Ardent's Sash (253887, -0.50 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.83 DPS) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.0 spell_power points (2.38 DPS) | yes | Filigreed Pristine Leggings (253937, -0.77 DPS) [crafted]; Silk-threaded Trousers (1929, -0.87 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -1.19 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (1.78 DPS) | yes | Feather Padded Treads (285345, -0.48 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.83 DPS) [crafted]; Pristine Boots (253889, -0.98 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.19 DPS) | yes | Lavishly Jeweled Ring (1156, -1.01 DPS) [dungeon]; Loop of Sacrifice (281673, -1.04 DPS) [quest]; Volcanic Rock Ring (12053, -1.10 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.71 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.56 DPS) [quest]; Volcanic Rock Ring (12053, -0.62 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 1.2 spell_power points (0.30 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.12 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 96.9 spell_power points (23.02 DPS) | yes | Skycaller (12984, -1.45 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.73 DPS) [dungeon]; Sizzle Stick (8071, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 142, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 58.1. Weights run: 0.9s. Verify run: 0.8s. 237 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.156, intellect=0.346 ± 0.011, crit=0.034 ± 0.002 per rating point (14 rating = 1%, 0.478 per %), hit=0.134 ± 0.001 per rating point (10 rating = 1%, 1.336 per %), spell_haste=not significant (0.301 ± 0.194), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.156

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.85 DPS) | yes | Silk Headband (7050, -0.34 DPS) [crafted]; Embalmed Shroud (7691, -0.51 DPS) [dungeon]; Enchanter's Cowl (4322, -0.97 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.1 spell_power points (1.53 DPS) | yes | Crystal Starfire Medallion (5003, -1.30 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.30 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.42 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (2.04 DPS) | yes | Fairywing Mantle (9536, -0.51 DPS) [quest]; Death Speaker Mantle (6685, -0.54 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.57 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.84 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.17 DPS) [crafted]; Battle Healer's Cloak (19529, -0.17 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.5 spell_power points (2.27 DPS) | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.45 DPS) [dungeon]; Pristine Gown (253961, -0.69 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.52 DPS) | yes | Nightsky Wristbands (6407, -1.17 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.17 DPS) [quest]; Glowing Magical Bracelets (13106, -1.42 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.7 spell_power points (1.30 DPS) | yes | Serpent Gloves (5970, -0.18 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.28 DPS) [crafted]; Gnoll Casting Gloves (892, -0.29 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (2.03 DPS) | yes | Belt of Arugal (6392, -0.34 DPS) [dungeon]; Warsong Sash (16975, -0.48 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.56 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.02 DPS) | yes | Abomination Skin Leggings (23173, -0.37 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.43 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (1.59 DPS) | yes | Acidic Walkers (9454, -0.28 DPS) [dungeon]; Boots of the Enchanter (4325, -0.74 DPS) [crafted]; Spidersilk Boots (4320, -1.57 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.18 DPS) | yes | Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Sludge-Stained Band (286535, -0.67 DPS) [world]; Black Widow Band (6199, -0.77 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.01 DPS) | yes | Electrocutioner Lagnut (9447, -0.51 DPS) [dungeon]; Black Widow Band (6199, -0.60 DPS) [world]; Sludge-Stained Band (286535, -1.81 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.52 DPS) | yes | Twisted Chanter's Staff (890, -0.93 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.93 DPS) [quest]; Glimmering Staff (249392, -1.19 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.1 spell_power points (1.53 DPS) | yes | Alliance Outrunner Healing Rod (285348, -0.59 DPS, sim-verified) [world]; Orb of Souls (249395, -0.85 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.96 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 200.2 spell_power points (33.72 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.87 DPS) [crafted]; Gravestone Scepter (7001, -4.72 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 237, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 116.7. Weights run: 0.8s. Verify run: 0.8s. 319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.127, intellect=0.281 ± 0.009, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.387 per %), hit=0.103 ± 0.001 per rating point (10 rating = 1%, 1.034 per %), spell_haste=1.008 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.127

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.88 DPS) | yes | Living Cowl (5608, -2.24 DPS) [world]; Augural Shroud (2620, -2.47 DPS, sim-verified) [world]; Holy Shroud (2721, -2.80 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (2.43 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.86 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.88 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.88 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.5 spell_power points (3.23 DPS) | yes | Inquisitor's Shawl (19507, -0.25 DPS) [dungeon]; Green Silken Shoulders (7057, -0.45 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 11.5 spell_power points (3.23 DPS) | yes | Guardian Cloak (5965, -1.15 DPS) [crafted]; Icy Cloak (4327, -1.27 DPS) [crafted]; Long Silken Cloak (4326, -2.09 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.7 spell_power points (6.63 DPS) | yes | Elemental Raiment (9434, -0.14 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.88 DPS) [crafted]; Robe of Power (7054, -1.77 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Condor Bracers (15864, -0.37 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.77 DPS) [quest]; Earthen Silk Cuffs (254019, -1.40 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (5.36 DPS) | yes | Black Mageweave Gloves (10003, -0.77 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.49 DPS) [crafted]; Gilded Handwraps (254021, -2.56 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.1 spell_power points (4.24 DPS) | yes | Star Belt (4329, +0.00 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.10 DPS) [dungeon]; Warsong Sash (16975, -1.15 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.4 spell_power points (4.86 DPS) | yes | Gaze Dreamer Pants (6903, -0.60 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.60 DPS) [crafted]; Abomination Skin Leggings (23173, -1.71 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.72 DPS) | yes | Gilded Slippers (254001, -2.78 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.45 DPS) [crafted]; Acidic Walkers (9454, -4.69 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (3.27 DPS) | yes | Reedknot Ring (9622, -1.31 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.59 DPS) [vendor]; Sludge-Stained Band (286535, -2.43 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.52 DPS) | yes | Reedknot Ring (9622, -0.37 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor]; Sludge-Stained Band (286535, -1.68 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Scorn's Focal Dagger (23168, -3.08 DPS) [dungeon]; Windweaver Staff (7757, -4.42 DPS) [dungeon]; Gut Ripper (2164, -7.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (116.7 DPS) | yes | Umbral Wand (5216, -0.01 DPS) [dungeon]; Earthen Rod (9381, -0.09 DPS) [dungeon]; Jaina's Firestarter (13064, -1.87 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 200.7. Weights run: 1.0s. Verify run: 0.9s. 406 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.227, intellect=0.211 ± 0.015, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.457 per %), hit=0.203 ± 0.003 per rating point (10 rating = 1%, 2.027 per %), spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.84 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.3 spell_power points (2.32 DPS) | yes | Mindburst Medallion (11196, -0.50 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.49 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.73 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.8 spell_power points (4.71 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.37 DPS) [crafted]; Bloodmage Mantle (7684, -1.65 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.28 DPS) | yes | Deep Woodlands Cloak (19121, -0.35 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.22 DPS) [dungeon]; Runecloth Cloak (13860, -1.28 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (6.52 DPS) | yes | Elemental Raiment (9434, -0.64 DPS) [world_drop]; Acumen Robes (17775, -0.92 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -0.94 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest]; Bloodband Bracers (11469, -0.59 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.28 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.11 DPS) [vendor]; Runecloth Gloves (13863, -1.39 DPS) [crafted]; Black Mageweave Gloves (10003, -1.51 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (4.52 DPS) | yes | Defiler's Cloth Girdle (20166, -0.29 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Ghostweave Cord (254073, -0.59 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.1 spell_power points (7.04 DPS) | yes | Red Mageweave Pants (10009, -2.40 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.84 DPS) [vendor]; Wizardweave Leggings (14132, -4.47 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -0.70 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.23 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.39 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.49 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon]; Runed Ring (862, -1.68 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.36 DPS) | yes | Philanthropist's Ring (281635, -0.32 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.43 DPS) [dungeon]; Runed Ring (862, -1.40 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.85 DPS) [crafted]; Rune of the Guard Captain (19120, -2.97 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -1.07 DPS, sim-verified) [crafted]; Rune of the Guard Captain (19120, -1.28 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Arbiter's Blade (11784, -3.07 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.08 DPS) [dungeon]; Blade of Eternal Darkness (17780, -6.89 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (200.7 DPS) | yes | Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted]; Pyric Caduceus (11748, -5.26 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Noxious Shooter

No-known-source sample (15 of 406, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 415.5. Weights run: 0.9s. Verify run: 0.9s. 1013 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.180, intellect=0.192 ± 0.014, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), hit=0.172 ± 0.002 per rating point (10 rating = 1%, 1.719 per %), spell_haste=-7.240 ± 0.381, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.180

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 68.5 spell_power points (37.77 DPS) | yes | Warlord's Dreadweave Hood (17591, -16.87 DPS, sim-verified) [vendor]; Crimson Felt Hat (18727, -20.38 DPS) [dungeon] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (12.13 DPS) | yes | Orb of the Darkmoon (19426, -0.15 DPS, sim-verified) [quest]; Amulet of the Dawn (22657, -2.48 DPS) [quest]; Kezan's Taint (19604, -3.56 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 44.8 spell_power points (24.69 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -7.54 DPS, sim-verified) [vendor]; Warlord's Dreadweave Mantle (231592, -9.11 DPS) [pvp] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.8 spell_power points (11.45 DPS) | yes | Amplifying Cloak (18350, -1.53 DPS) [dungeon]; Hide of the Wild (18510, -2.67 DPS) [crafted]; Arcanoweave Cloak (272411, -3.08 DPS, sim-verified) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 69.4 spell_power points (38.27 DPS) | yes | Robe of the Void (14153, -13.27 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -18.09 DPS) [pvp] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 30.2 spell_power points (16.66 DPS) | yes | Dryad's Wrist Bindings (19595, -4.58 DPS, sim-verified) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 41.2 spell_power points (22.72 DPS) | yes | General's Dreadweave Gloves (17588, -5.15 DPS, sim-verified) [vendor]; Sandworm Skin Gloves (20716, -7.31 DPS) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 45.4 spell_power points (25.01 DPS) | yes | Knowledge of the Timbermaw (228190, -7.47 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -11.97 DPS) [crafted]; Heretic Waistguard (240151, -13.24 DPS) [vendor] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 61.1 spell_power points (33.71 DPS) | yes | General's Dreadweave Pants (17593, -10.10 DPS, sim-verified) [vendor]; Skyshroud Leggings (13170, -14.12 DPS) [dungeon] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 40.4 spell_power points (22.27 DPS) | yes | General's Dreadweave Boots (17586, -5.98 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -9.05 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | sim-verified (415.5 DPS) | yes | Eye of Orgrimmar (12545, -6.83 DPS) [quest]; Maiden's Circle (13001, -6.83 DPS) [world_drop]; Naglering (11669, -14.86 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (415.5 DPS) | yes | Eye of Orgrimmar (12545, -2.63 DPS) [quest]; Maiden's Circle (13001, -2.63 DPS) [world_drop]; Naglering (11669, -12.99 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (415.5 DPS) | yes | Royal Seal of Eldre'Thalas (18467, -3.31 DPS) [quest]; Weakness Analyzer (272438, -3.86 DPS) [vendor]; Serenity Field (272439, -8.27 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (415.5 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -0.96 DPS, sim-verified) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (415.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -2.00 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -20.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 145.3 spell_power points (80.07 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.13 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -10.95 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.03 DPS) [world] |

**New at 60:** head: Heretic Cowl; neck: Chains of the Lich; shoulder: Heretic Shoulderpads; back: Crystalline Threaded Cape; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1013, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

