# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 28.6. Weights run: 1.0s. Verify run: 0.8s. 151 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.339 ± 0.008, crit=0.112 ± 0.003 per rating point (14 rating = 1%, 1.571 per %), hit=0.430 ± 0.012 per rating point (10 rating = 1%, 4.303 per %), spell_haste=not significant (-0.406 ± 0.124), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Shadow Goggles (4373, -1.12 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.0 spell_power points (0.71 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.15 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.36 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.35 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.18 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.7 spell_power points (0.59 DPS) | yes | Green Woolen Robe (6243, -0.24 DPS) [crafted]; Green Woolen Vest (2582, -0.24 DPS) [crafted]; Gray Woolen Robe (2585, -0.51 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.7 spell_power points (0.15 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.06 DPS) [quest]; Repurposed Hair Band (281256, -0.09 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.62 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.38 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.47 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.23 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.36 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.7 spell_power points (1.03 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.74 DPS) | yes | Feather Padded Treads (285345, -0.22 DPS, sim-verified) [world]; Pristine Boots (253889, -0.38 DPS) [crafted]; Red Woolen Boots (4313, -0.38 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.50 DPS) | yes | Sludge-Stained Band (286535, -0.24 DPS) [world]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.44 DPS) | yes | Sludge-Stained Band (286535, -0.19 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.4 spell_power points (0.30 DPS) | yes | Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Staff of Westfall (2042, -0.15 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 256.4 spell_power points (22.57 DPS) | yes | Skycaller (12984, -1.01 DPS) [world_drop]; Firebelcher (5243, -2.28 DPS) [dungeon]; Deepblaze (279896, -4.04 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 151, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 61.1. Weights run: 1.0s. Verify run: 0.8s. 262 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.441 ± 0.008, crit=0.049 ± 0.002 per rating point (14 rating = 1%, 0.683 per %), hit=0.234 ± 0.008 per rating point (10 rating = 1%, 2.343 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.18 DPS) | yes | Silk Headband (7050, -0.21 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.32 DPS) [crafted]; Enchanter's Cowl (4322, -1.02 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (1.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.27 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (1.39 DPS) | yes | Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Fairywing Mantle (9536, -0.32 DPS) [quest]; Invoker's Mantle (215365, -0.40 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (0.64 DPS) | yes | Hillman's Cloak (3719, -0.11 DPS) [crafted]; Repairman's Cape (9605, -0.13 DPS) [quest]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.7 spell_power points (1.58 DPS) | yes | Tree Bark Jacket (1486, -0.19 DPS) [dungeon]; Beguiler Robes (7728, -0.24 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.40 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.97 DPS) | yes | Nightsky Wristbands (6407, -0.68 DPS) [world_drop]; Stonecloth Bindings (14416, -0.73 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.01 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.9 spell_power points (0.95 DPS) | yes | Serpent Gloves (5970, -0.20 DPS) [dungeon]; Truefaith Gloves (7049, -0.27 DPS) [crafted]; Shilly Mitts (9609, -0.62 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.3 spell_power points (1.32 DPS) | yes | Belt of Arugal (6392, -0.21 DPS) [dungeon]; Invoker's Cord (215366, -0.33 DPS) [crafted]; Crimson Silk Belt (7055, -0.35 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.5 spell_power points (1.34 DPS) | yes | Pristine Leggings (253987, -0.26 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.29 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.1 spell_power points (1.08 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Nimbus Boots (6998, -0.44 DPS) [quest]; Spidersilk Boots (4320, -1.83 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.75 DPS) | yes | Minor Channeling Ring (1449, -0.12 DPS) [quest]; Black Widow Band (6199, -0.42 DPS) [world]; Snake Hoop (6750, -0.42 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.64 DPS) | yes | Black Widow Band (6199, -0.31 DPS) [world]; Snake Hoop (6750, -0.31 DPS) [quest]; Minor Channeling Ring (1449, -0.82 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (61.1 DPS) | yes | Royal Diplomatic Scepter (9457, -0.02 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.74 DPS) [dungeon]; Hardened Root Staff (1317, -1.95 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 312.5 spell_power points (33.53 DPS) | yes | Starfaller (13063, -0.28 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.00 DPS) [crafted]; Gravestone Scepter (7001, -4.53 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 130.1. Weights run: 1.0s. Verify run: 0.9s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.036 ± 0.010, crit=0.022 ± 0.001 per rating point (14 rating = 1%, 0.301 per %), hit=0.295 ± 0.013 per rating point (10 rating = 1%, 2.945 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 21.4 spell_power points (3.95 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.07 DPS) [crafted]; Corpseshroud (10574, -0.31 DPS) [dungeon]; Thinking Cap (2624, -0.69 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.3 spell_power points (2.83 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.91 DPS) [quest]; Darkspear Warding Pendant (272074, -1.49 DPS) [vendor]; Scorn's Icy Choker (23169, -1.68 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.5 spell_power points (3.79 DPS) | yes | Bloodmage Mantle (7684, -0.40 DPS) [dungeon]; Death Speaker Mantle (6685, -0.57 DPS) [dungeon]; Green Silken Shoulders (7057, -1.59 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 18.3 spell_power points (3.39 DPS) | yes | Long Silken Cloak (4326, -1.32 DPS) [crafted]; Guardian Cloak (5965, -1.32 DPS) [crafted]; Darkspear Raider's Cloak (272077, -2.82 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.2 spell_power points (5.22 DPS) | yes | Dreamweave Vest (10021, -0.17 DPS) [crafted]; Robe of Power (7054, -0.33 DPS) [crafted]; Green Silk Armor (7065, -1.06 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 9.3 spell_power points (1.73 DPS) | yes | Arcane Runed Bracers (4744, -0.06 DPS) [quest]; Spidertank Oilrag (9448, -0.06 DPS) [dungeon]; Mistscape Bracers (4045, -0.19 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.1 spell_power points (4.10 DPS) | yes | Red Mageweave Gloves (10018, -0.72 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -1.06 DPS) [crafted]; Town Clerk's Mittens (270029, -1.25 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.5 spell_power points (4.17 DPS) | yes | Highlander's Cloth Girdle (20098, -1.03 DPS, sim-verified) [rep]; Gilded Cord (254037, -1.16 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.4 spell_power points (4.89 DPS) | yes | Crimson Silk Pantaloons (7062, -1.26 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.69 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.06 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.44 DPS) | yes | Gilded Slippers (254001, -1.49 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.98 DPS) [dungeon]; Spidersilk Boots (4320, -2.38 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (2.81 DPS) | yes | Ring of Forlorn Spirits (2043, -1.33 DPS) [quest]; Voodoo Band (1996, -1.47 DPS) [world]; Black Widow Band (6199, -1.47 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.67 DPS) | yes | Voodoo Band (1996, -0.32 DPS) [world]; Black Widow Band (6199, -0.32 DPS) [world]; Ring of Forlorn Spirits (2043, -1.93 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (130.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.38 DPS) [dungeon]; Hand of Righteousness (7721, -3.94 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 16.4 spell_power points (3.03 DPS) | yes | Thrash's Trash (276204, -0.44 DPS) [vendor]; Orb of Mystic Insight (249394, -0.58 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.63 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 219.0 spell_power points (40.53 DPS) | yes | Umbral Wand (5216, -0.71 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.94 DPS) [dungeon]; Twisted Nether Wand (249144, -5.42 DPS) [crafted] |

**New at 40:** head: Augural Shroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 213.4. Weights run: 1.0s. Verify run: 1.4s. 441 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.473 ± 0.019, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.393 per %), hit=0.497 ± 0.026 per rating point (10 rating = 1%, 4.974 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.5 spell_power points (8.58 DPS) | yes | Soulcatcher Halo (10630, -2.06 DPS) [dungeon]; Dreamweave Circlet (10041, -2.25 DPS) [crafted]; Chief Architect's Monocle (11839, -3.77 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.6 spell_power points (3.65 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Scorn's Icy Choker (23169, -0.85 DPS) [dungeon]; Mindburst Medallion (11196, -1.02 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.5 spell_power points (6.99 DPS) | yes | Kentic Amice (11624, -1.13 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.84 DPS) [crafted]; Inquisitor's Shawl (19507, -2.37 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.8 spell_power points (4.04 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.36 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.39 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.5 spell_power points (8.58 DPS) | yes | Runecloth Robe (13858, -2.20 DPS) [crafted]; Runecloth Tunic (13857, -2.70 DPS) [crafted]; Robes of Insight (940, -4.06 DPS, sim-verified) [world_drop] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Forgotten Wraps (9433, -0.10 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.10 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.4 spell_power points (6.97 DPS) | yes | Virtuous Hands (226958, -1.04 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.33 DPS) [vendor]; Red Mageweave Gloves (10018, -2.42 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 34.0 spell_power points (6.01 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS, sim-verified) [dungeon]; Deathmage Sash (10771, -0.87 DPS) [dungeon]; Satyrmane Sash (17755, -0.93 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.7 spell_power points (6.68 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.36 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.45 DPS) [dungeon]; Red Mageweave Pants (10009, -4.20 DPS, sim-verified) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Southsea Mojo Boots (20641, -0.01 DPS) [quest]; Earthen Silk Slippers (254013, -0.05 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.69 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (213.4 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, +0.00 DPS) [dungeon]; Woodseed Hoop (17768, -0.73 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.3 spell_power points (3.42 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Woodseed Hoop (17768, -1.07 DPS) [quest]; Mindseye Circle (10634, -2.22 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, -3.49 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.81 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -3.39 DPS) [dungeon]; Blade of Eternal Darkness (17780, -25.10 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 296.7 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.62 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Gilded Sandals; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 441, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 400.6. Weights run: 1.1s. Verify run: 2.8s. 1121 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.509 ± 0.018, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.458 per %), hit=0.627 ± 0.029 per rating point (10 rating = 1%, 6.273 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Magister's Crown (16686, -5.08 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.6 spell_power points (9.25 DPS) | yes | Beads of Ogre Mojo (22149, -0.94 DPS) [quest]; Lady Maye's Pendant (14558, -1.59 DPS) [world_drop]; Diana's Pearl Necklace (22403, -1.94 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, -1.05 DPS) [pvp]; Darkspear Shoulderpads (272103, -5.08 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 34.3 spell_power points (9.18 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -2.22 DPS) [world]; Royal Tribunal Cloak (13376, -2.73 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -4.57 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 34.1 spell_power points (9.10 DPS) | yes | Marshal's Satin Bracers (17606, -1.44 DPS) [pvp]; Sublime Wristguards (18497, -1.87 DPS) [dungeon]; Runecloth Cuffs (254123, -2.13 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 48.7 spell_power points (13.02 DPS) | yes | Marshal's Satin Gloves (17608, -1.64 DPS) [vendor]; Marshal's Satin Grips (231617, -1.64 DPS) [vendor]; Virtuous Hands (226958, -2.71 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 64.5 spell_power points (17.23 DPS) | yes | Belt of the Archmage (18405, -4.32 DPS, sim-verified) [crafted]; Magister's Belt (16685, -6.89 DPS) [dungeon]; Dustfeather Sash (12589, -7.56 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (22752) | Silverwing Sentinels [rep] | 56.7 spell_power points (15.14 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Virtuous Leggings (226956, -0.98 DPS, sim-verified) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 46.1 spell_power points (12.33 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -5.15 DPS) [quest]; Maiden's Circle (13001, -5.15 DPS) [world_drop]; Naglering (11669, -12.02 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -2.68 DPS) [quest]; Maiden's Circle (13001, -2.68 DPS) [world_drop]; Naglering (11669, -7.22 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Briarwood Reed (12930, -15.41 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (400.6 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -0.53 DPS) [dungeon] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -2.83 DPS) [dungeon]; Hand of Edward the Odd (2243, -23.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 293.3 spell_power points (78.36 DPS) | yes | Sparkling Crystal Wand (20672, +0.00 DPS, sim-verified) [world]; Bonecreeper Stylus (13938, -11.18 DPS) [dungeon]; Oblivion's Touch (18761, -11.43 DPS) [dungeon] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1121, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 325001031300000000-00000000000000000-523120501201300251)

Set DPS (verified): 573.6. Weights run: 1.3s. Verify run: 3.3s. 1121 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.015 ± 0.002, crit=0.067 ± 0.005 per rating point (14 rating = 1%, 0.944 per %), hit=0.496 ± 0.035 per rating point (10 rating = 1%, 4.963 per %), spell_haste=-11.032 ± 0.408, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.08 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.50 DPS) [quest]; Diana's Pearl Necklace (22403, -2.91 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.2 spell_power points (10.35 DPS) | yes | Field Marshal's Satin Epaulets (231621, -1.05 DPS) [vendor]; Argent Shoulders (19059, -1.13 DPS, sim-verified) [crafted]; Virtuous Epaulets (226955, -1.48 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 21.1 spell_power points (7.75 DPS) | yes | Crystalline Threaded Cape (20697, -0.38 DPS) [world]; Amplifying Cloak (18350, -1.13 DPS) [dungeon]; Hide of the Wild (18510, -2.55 DPS) [crafted] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Robe of Everlasting Night (18385, +0.00 DPS) [dungeon]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.1 spell_power points (8.13 DPS) | yes | Sublime Wristguards (18497, -3.66 DPS) [dungeon]; Runecloth Cuffs (254123, -4.03 DPS) [crafted]; Arcane Runed Bracers (4744, -4.82 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.3 spell_power points (11.13 DPS) | yes | Stormpike Cloth Girdle (19094, -4.46 DPS) [rep]; Ban'thok Sash (11662, -4.84 DPS) [dungeon]; Belt of the Archmage (18405, -5.04 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; Marshal's Satin Pants (17603, -0.63 DPS) [vendor]; Marshal's Satin Leggings (231619, -0.63 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.82 DPS) | yes | Marshal's Satin Sandals (17607, -0.27 DPS) [vendor]; Marshal's Satin Treads (231620, -0.27 DPS) [vendor]; Virtuous Slippers (226959, -0.64 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -4.46 DPS) [dungeon]; Maiden's Circle (13001, -5.51 DPS) [world_drop]; Naglering (11669, -13.79 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.44 DPS) [dungeon]; Maiden's Circle (13001, -1.49 DPS) [world_drop]; Naglering (11669, -8.20 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Briarwood Reed (12930, -17.49 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (573.6 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -0.73 DPS) [dungeon] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.67 DPS) [dungeon]; The Lobotomizer (19324, -30.51 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+9.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.20 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.67 DPS) [world]; Torch of Light (279246, -9.39 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Earthen Silk Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1121, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 33.5. Weights run: 1.0s. Verify run: 0.8s. 141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.339 ± 0.008, crit=0.112 ± 0.003 per rating point (14 rating = 1%, 1.571 per %), hit=0.430 ± 0.012 per rating point (10 rating = 1%, 4.303 per %), spell_haste=not significant (-0.406 ± 0.124), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Shadow Goggles (4373, -1.19 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.0 spell_power points (0.71 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.36 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.37 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.35 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.7 spell_power points (0.59 DPS) | yes | Green Woolen Robe (6243, -0.24 DPS) [crafted]; Green Woolen Vest (2582, -0.24 DPS) [crafted]; Gray Woolen Robe (2585, -0.89 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.26 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Featherbead Bracers (15452, -0.12 DPS) [quest]; Owlbeard Bracers (16981, -0.12 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.62 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Apothecary Gloves (10919, -0.26 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.47 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.23 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.62 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.7 spell_power points (1.03 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.74 DPS) | yes | Pristine Boots (253889, -0.38 DPS) [crafted]; Red Woolen Boots (4313, -0.38 DPS) [crafted]; Feather Padded Treads (285345, -0.52 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.44 DPS) | yes | Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Loop of Sacrifice (281673, -0.29 DPS) [quest]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.26 DPS) | yes | Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Loop of Sacrifice (281673, -0.12 DPS) [quest]; Volcanic Rock Ring (12053, -0.17 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.4 spell_power points (0.30 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.06 DPS) [world]; Lesser Staff of the Spire (1300, -0.12 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 256.4 spell_power points (22.57 DPS) | yes | Skycaller (12984, -2.11 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.28 DPS) [dungeon]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 59.5. Weights run: 1.0s. Verify run: 0.8s. 249 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.441 ± 0.008, crit=0.049 ± 0.002 per rating point (14 rating = 1%, 0.683 per %), hit=0.234 ± 0.008 per rating point (10 rating = 1%, 2.343 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.18 DPS) | yes | Silk Headband (7050, -0.21 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.32 DPS) [crafted]; Enchanter's Cowl (4322, -0.69 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (1.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.95 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (1.39 DPS) | yes | Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Mantle of Woe (7750, -0.29 DPS) [quest]; Fairywing Mantle (9536, -0.32 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.54 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.7 spell_power points (1.58 DPS) | yes | Mechbuilder's Overalls (9508, -0.15 DPS) [dungeon]; Tree Bark Jacket (1486, -0.19 DPS) [dungeon]; Beguiler Robes (7728, -0.24 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.97 DPS) | yes | Tabitha's Cuffs (251486, -0.64 DPS) [quest]; Glowing Magical Bracelets (13106, -0.65 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.68 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.2 spell_power points (0.88 DPS) | yes | Truefaith Gloves (7049, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS) [world]; Serpent Gloves (5970, -0.46 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.3 spell_power points (1.32 DPS) | yes | Belt of Arugal (6392, -0.21 DPS) [dungeon]; Warsong Sash (16975, -0.26 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.33 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.5 spell_power points (1.34 DPS) | yes | Pristine Leggings (253987, -0.26 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.86 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.1 spell_power points (1.08 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Boots of the Enchanter (4325, -0.55 DPS) [crafted]; Spidersilk Boots (4320, -1.17 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.75 DPS) | yes | Black Widow Band (6199, -0.42 DPS) [world]; Snake Hoop (6750, -0.42 DPS) [quest]; Sludge-Stained Band (286535, -0.43 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.64 DPS) | yes | Snake Hoop (6750, -0.31 DPS) [quest]; Sludge-Stained Band (286535, -0.32 DPS) [world]; Black Widow Band (6199, -0.66 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 25.2 spell_power points (2.70 DPS) | yes | Royal Diplomatic Scepter (9457, -0.46 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -1.74 DPS) [dungeon]; Glimmering Staff (249392, -2.18 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (59.5 DPS) | yes | Starfaller (13063, -0.48 DPS) [world_drop]; Unstable Power Core (279847, -2.18 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -4.00 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 121.1. Weights run: 1.0s. Verify run: 1.0s. 326 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.036 ± 0.010, crit=0.022 ± 0.001 per rating point (14 rating = 1%, 0.301 per %), hit=0.295 ± 0.013 per rating point (10 rating = 1%, 2.945 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 21.4 spell_power points (3.95 DPS) | yes | Corpseshroud (10574, -0.31 DPS) [dungeon]; Thinking Cap (2624, -0.69 DPS) [world]; Spellpower Goggles Xtreme (10502, -1.36 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.3 spell_power points (2.83 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.91 DPS) [quest]; Darkspear Warding Pendant (272074, -1.49 DPS) [vendor]; Scorn's Icy Choker (23169, -2.04 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.5 spell_power points (3.79 DPS) | yes | Bloodmage Mantle (7684, -0.40 DPS) [dungeon]; Mantle of Woe (7750, -0.56 DPS) [quest]; Green Silken Shoulders (7057, -2.07 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 18.3 spell_power points (3.39 DPS) | yes | Long Silken Cloak (4326, -1.32 DPS) [crafted]; Guardian Cloak (5965, -1.32 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.86 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.2 spell_power points (5.22 DPS) | yes | Dreamweave Vest (10021, -0.17 DPS) [crafted]; Robe of Power (7054, -0.33 DPS) [crafted]; Zealot's Robe (17043, -0.53 DPS) [quest] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.3 spell_power points (2.27 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.61 DPS) [dungeon]; Windchaser Cuffs (14429, -0.96 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.1 spell_power points (4.10 DPS) | yes | Red Mageweave Gloves (10018, -0.15 DPS) [crafted]; Stormcloth Gloves (10011, -1.06 DPS) [crafted]; Gilded Handwraps (254021, -1.28 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.5 spell_power points (4.17 DPS) | yes | Gilded Cord (254037, -1.16 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.56 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.4 spell_power points (4.89 DPS) | yes | Crimson Silk Pantaloons (7062, -1.46 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.69 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.06 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.44 DPS) | yes | Gilded Slippers (254001, -1.02 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.98 DPS) [dungeon]; Spidersilk Boots (4320, -2.38 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (2.81 DPS) | yes | Ogremind Ring (1993, -1.47 DPS) [world_drop]; Voodoo Band (1996, -1.47 DPS) [world]; Black Widow Band (6199, -1.47 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.67 DPS) | yes | Ogremind Ring (1993, -0.32 DPS) [world_drop]; Voodoo Band (1996, -0.32 DPS) [world]; Black Widow Band (6199, -2.34 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (121.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.38 DPS) [dungeon]; Skullbreaker (17039, -0.56 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.59 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Mystic Insight (249394, -0.15 DPS) [crafted]; Prophetic Cane (6803, -0.29 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 219.0 spell_power points (40.53 DPS) | yes | Umbral Wand (5216, -4.86 DPS) [dungeon]; Earthen Rod (9381, -4.94 DPS) [dungeon]; Twisted Nether Wand (249144, -5.42 DPS) [crafted] |

**New at 40:** head: Augural Shroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 197.8. Weights run: 1.0s. Verify run: 1.3s. 421 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.473 ± 0.019, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.393 per %), hit=0.497 ± 0.026 per rating point (10 rating = 1%, 4.974 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.5 spell_power points (8.58 DPS) | yes | Soulcatcher Halo (10630, -2.06 DPS) [dungeon]; Dreamweave Circlet (10041, -2.25 DPS) [crafted]; Chief Architect's Monocle (11839, -2.34 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.6 spell_power points (3.65 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Scorn's Icy Choker (23169, -0.85 DPS) [dungeon]; Mindburst Medallion (11196, -1.02 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.5 spell_power points (6.99 DPS) | yes | Kentic Amice (11624, -1.13 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.84 DPS) [crafted]; Inquisitor's Shawl (19507, -2.37 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 25.3 spell_power points (4.47 DPS) | yes | Spritecaster Cape (11623, -0.43 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.53 DPS) [dungeon]; Runecloth Cloak (13860, -0.79 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.5 spell_power points (8.58 DPS) | yes | Runecloth Robe (13858, -2.20 DPS) [crafted]; Runecloth Tunic (13857, -2.70 DPS) [crafted]; Robes of Insight (940, -2.93 DPS, sim-verified) [world_drop] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.4 spell_power points (6.97 DPS) | yes | Virtuous Hands (226958, -0.22 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.33 DPS) [vendor]; Red Mageweave Gloves (10018, -2.42 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 34.0 spell_power points (6.01 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS, sim-verified) [dungeon]; Deathmage Sash (10771, -0.87 DPS) [dungeon]; Satyrmane Sash (17755, -0.93 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.7 spell_power points (6.68 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.36 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.45 DPS) [dungeon]; Red Mageweave Pants (10009, -3.81 DPS, sim-verified) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.2 spell_power points (4.64 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Southsea Mojo Boots (20641, -0.36 DPS) [quest]; Earthen Silk Slippers (254013, -0.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (197.8 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, +0.00 DPS) [dungeon]; Woodseed Hoop (17768, -0.73 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.3 spell_power points (3.42 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Woodseed Hoop (17768, -1.07 DPS) [quest]; Mindseye Circle (10634, -1.49 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.81 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -3.39 DPS) [dungeon]; Barman Shanker (12791, -25.00 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 296.7 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.62 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 411.0. Weights run: 1.1s. Verify run: 2.7s. 1112 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.509 ± 0.018, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.458 per %), hit=0.627 ± 0.029 per rating point (10 rating = 1%, 6.273 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -5.25 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.6 spell_power points (9.25 DPS) | yes | Beads of Ogre Mojo (22149, -0.79 DPS, sim-verified) [quest]; Lady Maye's Pendant (14558, -1.59 DPS) [world_drop]; Diana's Pearl Necklace (22403, -1.94 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, -1.05 DPS) [pvp]; Darkspear Shoulderpads (272103, -4.72 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 34.3 spell_power points (9.18 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -2.22 DPS) [world]; Deep Woodlands Cloak (19121, -2.34 DPS) [quest] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -4.70 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 34.1 spell_power points (9.10 DPS) | yes | General's Satin Bracers (17619, -1.44 DPS) [pvp]; Sublime Wristguards (18497, -1.87 DPS) [dungeon]; Runecloth Cuffs (254123, -2.13 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 48.7 spell_power points (13.02 DPS) | yes | General's Satin Gloves (17620, -1.64 DPS) [vendor]; General's Satin Grips (231613, -1.64 DPS) [vendor]; Virtuous Hands (226958, -2.71 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 64.5 spell_power points (17.23 DPS) | yes | Belt of the Archmage (18405, -4.10 DPS, sim-verified) [crafted]; Magister's Belt (16685, -6.89 DPS) [dungeon]; Dustfeather Sash (12589, -7.56 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 56.7 spell_power points (15.14 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Virtuous Leggings (226956, -1.21 DPS, sim-verified) [vendor]; General's Satin Legguards (231634, -1.33 DPS) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 46.1 spell_power points (12.33 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; Dragonrider Boots (18102, +0.00 DPS) [dungeon]; General's Satin Treads (231610, +0.00 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -5.15 DPS) [quest]; Maiden's Circle (13001, -5.15 DPS) [world_drop]; Naglering (11669, -12.04 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.68 DPS) [quest]; Maiden's Circle (13001, -2.68 DPS) [world_drop]; Naglering (11669, -7.29 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Briarwood Reed (12930, -15.45 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (411.0 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -0.53 DPS) [dungeon] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -2.83 DPS) [dungeon]; Hand of Edward the Odd (2243, -23.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 293.3 spell_power points (78.36 DPS) | yes | Sparkling Crystal Wand (20672, +0.00 DPS, sim-verified) [world]; Bonecreeper Stylus (13938, -11.18 DPS) [dungeon]; Oblivion's Touch (18761, -11.43 DPS) [dungeon] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Outrider's Silk Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1112, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 325001031300000000-00000000000000000-523120501201300251)

Set DPS (verified): 587.7. Weights run: 1.3s. Verify run: 3.2s. 1112 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.015 ± 0.002, crit=0.067 ± 0.005 per rating point (14 rating = 1%, 0.944 per %), hit=0.496 ± 0.035 per rating point (10 rating = 1%, 4.963 per %), spell_haste=-11.032 ± 0.408, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.08 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.50 DPS) [quest]; Diana's Pearl Necklace (22403, -2.91 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.2 spell_power points (10.35 DPS) | yes | Warlord's Satin Epaulets (231611, -1.05 DPS) [vendor]; Virtuous Epaulets (226955, -1.48 DPS) [vendor]; Argent Shoulders (19059, -1.56 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 21.1 spell_power points (7.75 DPS) | yes | Crystalline Threaded Cape (20697, -0.38 DPS) [world]; Amplifying Cloak (18350, -1.13 DPS) [dungeon]; Hide of the Wild (18510, -2.55 DPS) [crafted] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of Everlasting Night (18385, +0.00 DPS) [dungeon]; Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Truefaith Vestments (14154, -0.47 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.1 spell_power points (8.13 DPS) | yes | Sublime Wristguards (18497, -3.66 DPS) [dungeon]; Runecloth Cuffs (254123, -4.03 DPS) [crafted]; Spidertank Oilrag (9448, -4.82 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.3 spell_power points (11.13 DPS) | yes | Frostwolf Cloth Belt (19090, -4.46 DPS) [rep]; Ban'thok Sash (11662, -4.84 DPS) [dungeon]; Belt of the Archmage (18405, -5.73 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Satin Leggings (231614, -0.63 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.11 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.82 DPS) | yes | General's Satin Boots (17618, -0.27 DPS) [vendor]; General's Satin Treads (231610, -0.27 DPS) [vendor]; Virtuous Slippers (226959, -0.64 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -4.46 DPS) [dungeon]; Maiden's Circle (13001, -5.51 DPS) [world_drop]; Naglering (11669, -14.77 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.44 DPS) [dungeon]; Maiden's Circle (13001, -1.49 DPS) [world_drop]; Naglering (11669, -8.06 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Briarwood Reed (12930, -17.28 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (587.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Burst of Knowledge (11832, -0.73 DPS) [dungeon] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.67 DPS) [dungeon]; The Lobotomizer (19324, -30.64 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+13.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.20 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.67 DPS) [world]; Torch of Light (279246, -13.40 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Earthen Silk Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1112, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

