# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.1. Weights run: 2.0s. Verify run: 1.1s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.050, intellect=0.126 ± 0.007, crit=0.022 ± 0.001 per rating point (14 rating = 1%, 0.312 per %), hit=0.077 ± 0.001 per rating point (10 rating = 1%, 0.768 per %), spell_haste=0.412 ± 0.047, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.050

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.43 DPS) | yes | Shadow Goggles (4373, -2.59 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.1 spell_power points (1.46 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.50 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.95 DPS) | yes | Feyscale Cloak (6632, -0.24 DPS) [dungeon]; Caretaker's Cape (20428, -0.24 DPS) [rep]; Black Whelp Cloak (7283, -0.27 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 spell_power points (1.34 DPS) | yes | Green Woolen Vest (2582, -0.39 DPS) [crafted]; Bloody Apron (6226, -0.39 DPS) [dungeon]; Gray Woolen Robe (2585, -1.22 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.24 DPS) | yes | Mindthrust Bracers (1974, -0.09 DPS) [dungeon]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Repurposed Hair Band (281256, -0.18 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.66 DPS) | yes | Gnoll Casting Gloves (892, -0.31 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.62 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.13 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 spell_power points (1.07 DPS) | yes | Novice Ardent's Sash (253887, -0.51 DPS) [crafted]; Keller's Girdle (2911, -0.83 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.07 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.0 spell_power points (2.38 DPS) | yes | Silk-threaded Trousers (1929, -0.68 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.77 DPS) [crafted]; Rumpled Kilt (274741, -1.19 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (1.78 DPS) | yes | Feather Padded Treads (285345, -0.68 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.83 DPS) [crafted]; Pristine Boots (253889, -0.98 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.3 spell_power points (1.25 DPS) | yes | Sludge-Stained Band (286535, -0.54 DPS) [world]; Lavishly Jeweled Ring (1156, -1.07 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.16 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.19 DPS) | yes | Sludge-Stained Band (286535, -0.56 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -1.01 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.10 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 1.3 spell_power points (0.30 DPS) | yes | Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Staff of Westfall (2042, -0.15 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 96.9 spell_power points (23.02 DPS) | yes | Skycaller (12984, -1.32 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.73 DPS) [dungeon]; Sizzle Stick (8071, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.6. Weights run: 2.0s. Verify run: 1.1s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.126, intellect=0.389 ± 0.009, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.582 per %), hit=0.148 ± 0.001 per rating point (10 rating = 1%, 1.477 per %), spell_haste=not significant (0.415 ± 0.159), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.126

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.65 DPS) | yes | Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon]; Enchanter's Cowl (4322, -0.88 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 spell_power points (1.40 DPS) | yes | Crystal Starfire Medallion (5003, -1.17 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.17 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.88 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (1.88 DPS) | yes | Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.53 DPS) [crafted]; Death Speaker Mantle (6685, -0.64 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Repairman's Cape (9605, -0.07 DPS) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Prelacy Cape (7004, -0.15 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.1 spell_power points (2.11 DPS) | yes | Death Speaker Robes (6682, -0.42 DPS) [dungeon]; Tree Bark Jacket (1486, -0.42 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.65 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Nightsky Wristbands (6407, -1.00 DPS) [world_drop]; Stonecloth Bindings (14416, -1.06 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.20 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.3 spell_power points (1.24 DPS) | yes | Serpent Gloves (5970, -0.19 DPS) [dungeon]; Shilly Mitts (9609, -0.19 DPS) [quest]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.2 spell_power points (1.83 DPS) | yes | Invoker's Cord (215366, -0.48 DPS) [crafted]; Crimson Silk Belt (7055, -0.52 DPS) [crafted]; Belt of Arugal (6392, -0.88 DPS, sim-verified) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.1 spell_power points (1.82 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.36 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.57 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (1.46 DPS) | yes | Acidic Walkers (9454, -0.24 DPS) [dungeon]; Nimbus Boots (6998, -0.56 DPS) [quest]; Spidersilk Boots (4320, -2.45 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.05 DPS) | yes | Minor Channeling Ring (1449, -0.18 DPS) [quest]; Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon]; Sludge-Stained Band (286535, -0.60 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.90 DPS) | yes | Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Sludge-Stained Band (286535, -0.45 DPS) [world]; Minor Channeling Ring (1449, -1.66 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Twisted Chanter's Staff (890, -0.77 DPS) [world_drop]; Channeler's Staff (4437, -0.88 DPS) [world]; Glimmering Staff (249392, -1.61 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.3 spell_power points (1.40 DPS) | yes | Eye of Paleth (2943, -0.80 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.80 DPS) [world]; Dwarven Tome (279898, -1.11 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 224.4 spell_power points (33.66 DPS) | yes | Starfaller (13063, -0.63 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 119.9. Weights run: 1.7s. Verify run: 1.1s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.090, intellect=0.270 ± 0.006, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.392 per %), hit=0.099 ± 0.001 per rating point (10 rating = 1%, 0.992 per %), spell_haste=0.913 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.090

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.09 DPS) | yes | Living Cowl (5608, -2.32 DPS) [world]; Augural Shroud (2620, -2.43 DPS, sim-verified) [world]; Holy Shroud (2721, -2.90 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.6 spell_power points (2.50 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.62 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.95 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.95 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.4 spell_power points (3.31 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Inquisitor's Shawl (19507, -0.27 DPS) [dungeon]; Berylline Pads (4197, -0.50 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 11.4 spell_power points (3.31 DPS) | yes | Guardian Cloak (5965, -1.18 DPS) [crafted]; Icy Cloak (4327, -1.28 DPS) [crafted]; Long Silken Cloak (4326, -1.68 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.6 spell_power points (6.84 DPS) | yes | Elemental Raiment (9434, -0.76 DPS) [world_drop]; Dreamweave Vest (10021, -0.92 DPS) [crafted]; Robe of Power (7054, -1.85 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.61 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.58 DPS) [quest]; Earthen Silk Cuffs (254019, -1.45 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (5.53 DPS) | yes | Black Mageweave Gloves (10003, -1.05 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.56 DPS) [crafted]; Gilded Handwraps (254021, -2.66 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.1 spell_power points (4.37 DPS) | yes | Star Belt (4329, -0.60 DPS) [crafted]; Deathmage Sash (10771, -1.17 DPS) [dungeon]; Gilded Cord (254037, -1.43 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.2 spell_power points (4.99 DPS) | yes | Gaze Dreamer Pants (6903, -0.48 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.66 DPS) [crafted]; Abomination Skin Leggings (23173, -1.76 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.96 DPS) | yes | Gilded Slippers (254001, -2.11 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.61 DPS) [crafted]; Acidic Walkers (9454, -4.88 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.6 spell_power points (3.37 DPS) | yes | Ring of Forlorn Spirits (2043, -1.05 DPS) [quest]; Reedknot Ring (9622, -1.34 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.63 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.61 DPS) | yes | Ring of Forlorn Spirits (2043, -0.29 DPS) [quest]; Reedknot Ring (9622, -0.58 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.87 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.19 DPS) [dungeon]; Windweaver Staff (7757, -4.62 DPS) [dungeon]; Gut Ripper (2164, -7.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (119.9 DPS) | yes | Umbral Wand (5216, -0.07 DPS) [dungeon]; Earthen Rod (9381, -0.15 DPS) [dungeon]; Jaina's Firestarter (13064, -1.88 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 211.8. Weights run: 1.9s. Verify run: 1.3s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.212, intellect=0.251 ± 0.014, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.603 per %), hit=0.257 ± 0.003 per rating point (10 rating = 1%, 2.569 per %), spell_haste=not significant (-0.121 ± 0.270), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.212

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.02 DPS) | yes | Dreamweave Circlet (10041, -0.78 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.34 DPS) [crafted]; Red Mageweave Headband (10033, -2.97 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (1.90 DPS) | yes | Mindburst Medallion (11196, -0.22 DPS) [quest]; Horizon Choker (13085, -1.11 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.34 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.5 spell_power points (3.90 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.17 DPS) [crafted]; Bloodmage Mantle (7684, -1.39 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (3.46 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.55 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.00 DPS) [crafted]; Nightfall Drape (12465, -1.45 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.0 spell_power points (5.35 DPS) | yes | Robe of the Magi (1716, -0.11 DPS) [world_drop]; Elemental Raiment (9434, -0.67 DPS) [world_drop]; Dreamweave Vest (10021, -0.84 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.01 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Bloodband Bracers (11469, -0.39 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (4.24 DPS) | yes | Black Mageweave Gloves (10003, -0.89 DPS) [crafted]; Runecloth Gloves (13863, -1.06 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.05 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Cloth Girdle (20098, -0.00 DPS) [rep]; Ghostweave Cord (254073, -0.23 DPS) [crafted]; Satyrmane Sash (17755, -2.25 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (5.69 DPS) | yes | Red Mageweave Pants (10009, -1.89 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.21 DPS) [vendor]; Wizardweave Leggings (14132, -5.13 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.35 DPS) | yes | Gilded Sandals (254107, -1.12 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.49 DPS) [vendor]; Black Mageweave Boots (10026, -2.51 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.90 DPS) | yes | Philanthropist's Ring (281635, -0.33 DPS) [quest]; Cyclopean Band (11824, -0.50 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.11 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (2.68 DPS) | yes | Cyclopean Band (11824, -0.28 DPS) [dungeon]; Philanthropist's Ring (281635, -0.77 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -0.89 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.68 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.40 DPS) [dungeon]; Scorn's Focal Dagger (23168, -2.45 DPS) [dungeon]; Blade of Eternal Darkness (17780, -8.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.69 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.58 DPS) [crafted]; Pyric Caduceus (11748, -5.38 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 360.9. Weights run: 1.9s. Verify run: 1.2s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.146, intellect=0.227 ± 0.012, crit=0.046 ± 0.002 per rating point (14 rating = 1%, 0.647 per %), hit=0.198 ± 0.002 per rating point (10 rating = 1%, 1.981 per %), spell_haste=-8.680 ± 0.316, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.146

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.8 spell_power points (15.26 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -2.92 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -6.88 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.55 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.94 DPS) [quest]; Kezan's Taint (19604, -2.97 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.1 spell_power points (14.90 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.05 DPS) [pvp]; Argent Shoulders (19059, -3.52 DPS, sim-verified) [crafted]; Burial Shawl (18681, -3.56 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.9 spell_power points (10.03 DPS) | yes | Arcanoweave Cloak (272411, -1.18 DPS, sim-verified) [vendor]; Amplifying Cloak (18350, -1.40 DPS) [dungeon]; Hide of the Wild (18510, -2.22 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.0 spell_power points (23.04 DPS) | yes | Robe of Everlasting Night (18385, -3.79 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -5.08 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -8.87 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.8 spell_power points (11.42 DPS) | yes | Sublime Wristguards (18497, -4.58 DPS) [dungeon]; Runecloth Cuffs (254123, -5.06 DPS) [crafted]; Deathmist Bracers (226907, -6.76 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.1 spell_power points (13.50 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.94 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.0 spell_power points (15.34 DPS) | yes | Belt of the Archmage (18405, -3.95 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -5.62 DPS) [rep]; Satyrmane Sash (17755, -7.53 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 35.8 spell_power points (17.18 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -1.68 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.51 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.76 DPS) [pvp]; Omnicast Boots (11822, -2.13 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (360.9 DPS) | yes | Songstone of Ironforge (12543, -2.35 DPS) [quest]; Maiden's Circle (13001, -2.35 DPS) [world_drop]; Naglering (11669, -10.73 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (360.9 DPS) | yes | Songstone of Ironforge (12543, -1.44 DPS) [quest]; Maiden's Circle (13001, -1.44 DPS) [world_drop]; Naglering (11669, -10.05 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (360.9 DPS) | yes | Royal Seal of Eldre'Thalas (18467, -2.88 DPS) [quest]; Weakness Analyzer (272438, -3.36 DPS) [vendor]; Serenity Field (272439, -7.19 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (360.9 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.30 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (360.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.71 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -19.24 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 166.0 spell_power points (79.64 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -11.30 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.93 DPS) [world] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.5. Weights run: 2.0s. Verify run: 1.1s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.050, intellect=0.126 ± 0.007, crit=0.022 ± 0.001 per rating point (14 rating = 1%, 0.312 per %), hit=0.077 ± 0.001 per rating point (10 rating = 1%, 0.768 per %), spell_haste=0.412 ± 0.047, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.050

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.43 DPS) | yes | Shadow Goggles (4373, -2.17 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.1 spell_power points (1.46 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.19 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.95 DPS) | yes | Black Whelp Cloak (7283, -0.15 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.24 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.24 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 spell_power points (1.34 DPS) | yes | Green Woolen Vest (2582, -0.39 DPS) [crafted]; Bloody Apron (6226, -0.39 DPS) [dungeon]; Gray Woolen Robe (2585, -0.88 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.3 spell_power points (0.30 DPS) | yes | Windsong Bangles (263336, -0.06 DPS) [quest]; Tabitha's Cuffs (251486, -0.12 DPS) [quest]; Featherbead Bracers (15452, -0.15 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.66 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.62 DPS) [crafted]; Apothecary Gloves (10919, -0.71 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 spell_power points (1.07 DPS) | yes | Novice Ardent's Sash (253887, -0.51 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.83 DPS) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.0 spell_power points (2.38 DPS) | yes | Filigreed Pristine Leggings (253937, -0.77 DPS) [crafted]; Silk-threaded Trousers (1929, -0.87 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -1.19 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (1.78 DPS) | yes | Feather Padded Treads (285345, -0.48 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.83 DPS) [crafted]; Pristine Boots (253889, -0.98 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.19 DPS) | yes | Lavishly Jeweled Ring (1156, -1.01 DPS) [dungeon]; Loop of Sacrifice (281673, -1.04 DPS) [quest]; Volcanic Rock Ring (12053, -1.10 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.71 DPS) | yes | Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Loop of Sacrifice (281673, -0.56 DPS) [quest]; Volcanic Rock Ring (12053, -0.62 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 1.3 spell_power points (0.30 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.12 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 96.9 spell_power points (23.02 DPS) | yes | Skycaller (12984, -1.45 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.73 DPS) [dungeon]; Sizzle Stick (8071, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 58.9. Weights run: 2.0s. Verify run: 1.1s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.126, intellect=0.389 ± 0.009, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.582 per %), hit=0.148 ± 0.001 per rating point (10 rating = 1%, 1.477 per %), spell_haste=not significant (0.415 ± 0.159), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.126

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.65 DPS) | yes | Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon]; Enchanter's Cowl (4322, -0.71 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 spell_power points (1.40 DPS) | yes | Crystal Starfire Medallion (5003, -1.17 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.17 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.71 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (1.88 DPS) | yes | Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.53 DPS) [crafted]; Death Speaker Mantle (6685, -0.54 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.1 spell_power points (2.11 DPS) | yes | Tree Bark Jacket (1486, -0.16 DPS) [dungeon]; Death Speaker Robes (6682, -0.42 DPS) [dungeon]; Pristine Gown (253961, -0.65 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Nightsky Wristbands (6407, -1.00 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.00 DPS) [quest]; Glowing Magical Bracelets (13106, -1.32 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.9 spell_power points (1.19 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.27 DPS) [crafted]; Gnoll Casting Gloves (892, -0.29 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.2 spell_power points (1.83 DPS) | yes | Warsong Sash (16975, +0.00 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.30 DPS) [dungeon]; Invoker's Cord (215366, -0.48 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.1 spell_power points (1.82 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.36 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.57 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (1.46 DPS) | yes | Acidic Walkers (9454, -0.24 DPS) [dungeon]; Boots of the Enchanter (4325, -0.71 DPS) [crafted]; Spidersilk Boots (4320, -1.22 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.05 DPS) | yes | Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon]; Sludge-Stained Band (286535, -0.60 DPS) [world]; Black Widow Band (6199, -0.64 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.90 DPS) | yes | Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Black Widow Band (6199, -0.49 DPS) [world]; Sludge-Stained Band (286535, -2.00 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Twisted Chanter's Staff (890, -0.77 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.77 DPS) [quest]; Glimmering Staff (249392, -0.86 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.3 spell_power points (1.40 DPS) | yes | Orb of Souls (249395, -0.80 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.87 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -0.87 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 224.4 spell_power points (33.66 DPS) | yes | Starfaller (13063, -0.49 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 118.6. Weights run: 1.7s. Verify run: 1.0s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.090, intellect=0.270 ± 0.006, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.392 per %), hit=0.099 ± 0.001 per rating point (10 rating = 1%, 0.992 per %), spell_haste=0.913 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.090

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.09 DPS) | yes | Living Cowl (5608, -2.32 DPS) [world]; Augural Shroud (2620, -2.51 DPS, sim-verified) [world]; Holy Shroud (2721, -2.90 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.6 spell_power points (2.50 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.89 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.95 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.95 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.4 spell_power points (3.31 DPS) | yes | Inquisitor's Shawl (19507, -0.27 DPS) [dungeon]; Green Silken Shoulders (7057, -0.46 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.50 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 11.4 spell_power points (3.31 DPS) | yes | Guardian Cloak (5965, -1.18 DPS) [crafted]; Icy Cloak (4327, -1.28 DPS) [crafted]; Long Silken Cloak (4326, -2.12 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.6 spell_power points (6.84 DPS) | yes | Elemental Raiment (9434, -0.76 DPS) [world_drop]; Dreamweave Vest (10021, -0.92 DPS) [crafted]; Robe of Power (7054, -1.85 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.61 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.58 DPS) [quest]; Radiant Silver Bracers (4545, -0.82 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (5.53 DPS) | yes | Black Mageweave Gloves (10003, -0.78 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.56 DPS) [crafted]; Gilded Handwraps (254021, -2.66 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.1 spell_power points (4.37 DPS) | yes | Star Belt (4329, -0.60 DPS) [crafted]; Deathmage Sash (10771, -1.17 DPS) [dungeon]; Warsong Sash (16975, -1.18 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.2 spell_power points (4.99 DPS) | yes | Gaze Dreamer Pants (6903, -0.61 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.66 DPS) [crafted]; Abomination Skin Leggings (23173, -1.76 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.96 DPS) | yes | Gilded Slippers (254001, -2.83 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.61 DPS) [crafted]; Acidic Walkers (9454, -4.88 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.6 spell_power points (3.37 DPS) | yes | Reedknot Ring (9622, -1.34 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.63 DPS) [vendor]; Sludge-Stained Band (286535, -2.50 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.61 DPS) | yes | Reedknot Ring (9622, -0.58 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.87 DPS) [vendor]; Sludge-Stained Band (286535, -1.74 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.19 DPS) [dungeon]; Windweaver Staff (7757, -4.62 DPS) [dungeon]; Gut Ripper (2164, -7.54 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (118.6 DPS) | yes | Umbral Wand (5216, -0.07 DPS) [dungeon]; Earthen Rod (9381, -0.15 DPS) [dungeon]; Jaina's Firestarter (13064, -1.91 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 209.6. Weights run: 1.9s. Verify run: 1.3s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.212, intellect=0.251 ± 0.014, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.603 per %), hit=0.257 ± 0.003 per rating point (10 rating = 1%, 2.569 per %), spell_haste=not significant (-0.121 ± 0.270), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.212

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.02 DPS) | yes | Red Mageweave Headband (10033, -0.62 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -0.78 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.34 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (1.90 DPS) | yes | Mindburst Medallion (11196, -0.22 DPS) [quest]; Horizon Choker (13085, -1.11 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.34 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -1.12 DPS) [crafted]; Bloodmage Mantle (7684, -1.34 DPS) [dungeon]; Rotgrip Mantle (17732, -2.43 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (3.46 DPS) | yes | Deep Woodlands Cloak (19121, -0.28 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.95 DPS) [dungeon]; Runecloth Cloak (13860, -1.00 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.0 spell_power points (5.35 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.67 DPS) [world_drop]; Dreamweave Vest (10021, -0.84 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -2.46 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (4.24 DPS) | yes | Black Mageweave Gloves (10003, -0.89 DPS) [crafted]; Runecloth Gloves (13863, -1.06 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.45 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (3.68 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -0.34 DPS) [rep]; Ghostweave Cord (254073, -0.56 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (5.69 DPS) | yes | Red Mageweave Pants (10009, -1.89 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.21 DPS) [vendor]; Wizardweave Leggings (14132, -4.13 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.35 DPS) | yes | Gilded Sandals (254107, -2.40 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.49 DPS) [vendor]; Black Mageweave Boots (10026, -2.51 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.90 DPS) | yes | Philanthropist's Ring (281635, -0.33 DPS) [quest]; Cyclopean Band (11824, -0.50 DPS) [dungeon]; Runed Ring (862, -1.34 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (2.68 DPS) | yes | Philanthropist's Ring (281635, -0.11 DPS) [quest]; Cyclopean Band (11824, -0.28 DPS) [dungeon]; Runed Ring (862, -1.11 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.34 DPS) [world_drop]; Rune of the Guard Captain (19120, -2.27 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.31 DPS, sim-verified) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.40 DPS) [dungeon]; Scorn's Focal Dagger (23168, -2.45 DPS) [dungeon]; Blade of Eternal Darkness (17780, -6.93 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.69 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.58 DPS) [crafted]; Pyric Caduceus (11748, -5.28 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 359.7. Weights run: 1.9s. Verify run: 1.3s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.146, intellect=0.227 ± 0.012, crit=0.046 ± 0.002 per rating point (14 rating = 1%, 0.647 per %), hit=0.198 ± 0.002 per rating point (10 rating = 1%, 1.981 per %), spell_haste=-8.680 ± 0.316, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.146

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.8 spell_power points (15.26 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -2.92 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.67 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.55 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.94 DPS) [quest]; Kezan's Taint (19604, -2.97 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.1 spell_power points (14.90 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.05 DPS) [pvp]; Burial Shawl (18681, -3.56 DPS) [dungeon]; Argent Shoulders (19059, -3.79 DPS, sim-verified) [crafted] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.9 spell_power points (10.03 DPS) | yes | Amplifying Cloak (18350, -1.40 DPS) [dungeon]; Arcanoweave Cloak (272411, -1.52 DPS, sim-verified) [vendor]; Hide of the Wild (18510, -2.22 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.0 spell_power points (23.04 DPS) | yes | Robe of Everlasting Night (18385, -3.74 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -5.08 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -8.87 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.8 spell_power points (11.42 DPS) | yes | Sublime Wristguards (18497, -4.58 DPS) [dungeon]; Runecloth Cuffs (254123, -5.06 DPS) [crafted]; Deathmist Bracers (226907, -6.76 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (359.7 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.57 DPS) [quest]; Sandworm Skin Gloves (20716, -3.94 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.0 spell_power points (15.34 DPS) | yes | Belt of the Archmage (18405, -3.38 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -5.62 DPS) [rep]; Satyrmane Sash (17755, -7.53 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 35.8 spell_power points (17.18 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.68 DPS) [rep]; Sentinel's Silk Leggings (237815, -2.88 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.51 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.96 DPS) [crafted]; Omnicast Boots (11822, -1.25 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.35 DPS) [quest]; Maiden's Circle (13001, -2.35 DPS) [world_drop]; Naglering (11669, -8.75 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.44 DPS) [quest]; Maiden's Circle (13001, -1.44 DPS) [world_drop]; Naglering (11669, -8.45 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+14.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.53 DPS, sim-verified) [quest]; Weakness Analyzer (272438, -3.36 DPS) [vendor]; Serenity Field (272439, -7.19 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.71 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -16.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 166.0 spell_power points (79.64 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.51 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.30 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.93 DPS) [world] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

