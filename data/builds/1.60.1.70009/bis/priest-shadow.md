# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 29.1. Weights run: 0.9s. Verify run: 0.7s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.233 ± 0.006, crit=0.105 ± 0.003 per rating point (14 rating = 1%, 1.467 per %), hit=0.326 ± 0.002 per rating point (10 rating = 1%, 3.261 per %), spell_haste=not significant (-0.280 ± 0.131), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -1.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.1 spell_power points (0.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.11 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Caretaker's Cape (20428, -0.10 DPS) [rep]; Black Whelp Cloak (7283, -0.10 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.2 spell_power points (0.59 DPS) | yes | Green Woolen Vest (2582, -0.21 DPS) [crafted]; Bloody Apron (6226, -0.21 DPS) [dungeon]; Gray Woolen Robe (2585, -0.52 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.2 spell_power points (0.11 DPS) | yes | Windsong Bangles (263336, -0.02 DPS) [quest]; Bright Bracers (3647, -0.02 DPS) [world_drop]; Repurposed Hair Band (281256, -0.07 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.67 DPS) | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.22 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.43 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.47 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.29 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.37 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.9 spell_power points (1.04 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.37 DPS) [dungeon]; Rumpled Kilt (274741, -0.56 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.76 DPS) | yes | Feather Padded Treads (285345, -0.22 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.40 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (0.52 DPS) | yes | Sludge-Stained Band (286535, -0.23 DPS) [world]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.45 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.48 DPS) | yes | Sludge-Stained Band (286535, -0.20 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.3 spell_power points (0.22 DPS) | yes | Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world]; Staff of Westfall (2042, -0.11 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 237.1 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.04 DPS) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-543110401200000000)

Set DPS (verified): 55.4. Weights run: 0.9s. Verify run: 0.7s. 244 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.401 ± 0.009, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.587 per %), hit=0.191 ± 0.002 per rating point (10 rating = 1%, 1.908 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.30 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.24 DPS) [crafted]; Embalmed Shroud (7691, -0.35 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (1.11 DPS) | yes | Crystal Starfire Medallion (5003, -0.92 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.92 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.00 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (1.49 DPS) | yes | Death Speaker Mantle (6685, -0.25 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.59 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Prelacy Cape (7004, -0.12 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.2 spell_power points (1.68 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.52 DPS) [crafted]; Tree Bark Jacket (1486, -1.22 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.06 DPS) | yes | Nightsky Wristbands (6407, -0.78 DPS) [world_drop]; Stonecloth Bindings (14416, -0.83 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.30 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.4 spell_power points (0.99 DPS) | yes | Serpent Gloves (5970, -0.17 DPS) [dungeon]; Truefaith Gloves (7049, -0.26 DPS) [crafted]; Shilly Mitts (9609, -0.45 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.2 spell_power points (1.44 DPS) | yes | Belt of Arugal (6392, -0.25 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.38 DPS) [crafted]; Crimson Silk Belt (7055, -0.40 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 spell_power points (1.44 DPS) | yes | Gaze Dreamer Pants (6903, -0.25 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.8 spell_power points (1.16 DPS) | yes | Acidic Walkers (9454, -0.19 DPS) [dungeon]; Nimbus Boots (6998, -0.45 DPS) [quest]; Spidersilk Boots (4320, -1.64 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.83 DPS) | yes | Minor Channeling Ring (1449, -0.14 DPS) [quest]; Electrocutioner Lagnut (9447, -0.47 DPS) [dungeon]; Sludge-Stained Band (286535, -0.47 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.71 DPS) | yes | Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Sludge-Stained Band (286535, -0.35 DPS) [world]; Minor Channeling Ring (1449, -0.88 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.06 DPS) | yes | Twisted Chanter's Staff (890, -0.59 DPS) [world_drop]; Glimmering Staff (249392, -0.64 DPS, sim-verified) [crafted]; Channeler's Staff (4437, -0.68 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.4 spell_power points (1.11 DPS) | yes | Eye of Paleth (2943, -0.64 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.64 DPS) [world]; Dwarven Tome (279898, -0.80 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 284.1 spell_power points (33.56 DPS) | yes | Starfaller (13063, -0.51 DPS) [world_drop]; Greater Mystic Wand (217287, -3.97 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-543110401201300240)

Set DPS (verified): 92.5. Weights run: 1.0s. Verify run: 0.7s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.711 ± 0.015, crit=0.040 ± 0.002 per rating point (14 rating = 1%, 0.554 per %), hit=0.289 ± 0.003 per rating point (10 rating = 1%, 2.889 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.79 DPS) | yes | Augural Shroud (2620, -0.38 DPS) [world]; Corpseshroud (10574, -1.00 DPS) [dungeon]; Enchanter's Cowl (4322, -1.05 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.3 spell_power points (1.50 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.50 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.84 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.84 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (2.16 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.4 spell_power points (2.05 DPS) | yes | Guardian Cloak (5965, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.01 DPS) [vendor]; Long Silken Cloak (4326, -1.77 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.3 spell_power points (3.50 DPS) | yes | Dreamweave Vest (10021, -0.25 DPS) [crafted]; Robe of Power (7054, -0.50 DPS) [crafted]; Elemental Raiment (9434, -0.70 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (92.5 DPS) | yes | Condor Bracers (15864, -0.27 DPS) [quest]; Windchaser Cuffs (14429, -0.35 DPS) [world_drop]; Arcane Runed Bracers (4744, -1.22 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (2.77 DPS) | yes | Red Mageweave Gloves (10018, -0.36 DPS) [crafted]; Black Mageweave Gloves (10003, -0.78 DPS) [crafted]; Gilded Handwraps (254021, -1.05 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.7 spell_power points (2.35 DPS) | yes | Gilded Cord (254037, -0.53 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.53 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (3.00 DPS) | yes | Crimson Silk Pantaloons (7062, -0.44 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.04 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.31 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.19 DPS) | yes | Gilded Slippers (254001, -0.56 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.77 DPS) [dungeon]; Spidersilk Boots (4320, -1.88 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.90 DPS) | yes | Ring of Forlorn Spirits (2043, -0.83 DPS) [quest]; Reedknot Ring (9622, -0.97 DPS) [quest]; Minor Channeling Ring (1449, -1.04 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.20 DPS) | yes | Reedknot Ring (9622, -0.27 DPS) [quest]; Minor Channeling Ring (1449, -0.34 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.08 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -1.24 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.46 DPS) [dungeon]; Gut Ripper (2164, -4.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 300.2 spell_power points (39.95 DPS) | yes | Umbral Wand (5216, -1.28 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.36 DPS) [dungeon]; Twisted Nether Wand (249144, -5.15 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 521000000000000000-00000000000000000-543110401201300251)

Set DPS (verified): 145.2. Weights run: 1.0s. Verify run: 0.8s. 423 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.020 ± 0.019, crit=0.047 ± 0.002 per rating point (14 rating = 1%, 0.661 per %), hit=0.408 ± 0.006 per rating point (10 rating = 1%, 4.077 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 39.4 spell_power points (5.20 DPS) | yes | Dreamweave Circlet (10041, -1.08 DPS) [crafted]; Chief Architect's Monocle (11839, -1.57 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -1.64 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.3 spell_power points (1.88 DPS) | yes | Scorn's Icy Choker (23169, -0.15 DPS) [dungeon]; Mindburst Medallion (11196, -0.28 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.54 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (145.2 DPS) | yes | Red Mageweave Shoulders (10029, -0.65 DPS) [crafted]; Inquisitor's Shawl (19507, -0.92 DPS) [dungeon]; Rotgrip Mantle (17732, -1.59 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.1 spell_power points (2.65 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.26 DPS) [dungeon]; Runecloth Cloak (13860, -0.39 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.77 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 39.4 spell_power points (5.20 DPS) | yes | Runecloth Tunic (13857, -1.48 DPS) [crafted]; Robe of the Magi (1716, -1.49 DPS) [world_drop]; Runecloth Robe (13858, -1.60 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 15.3 spell_power points (2.02 DPS) | yes | Bloodband Bracers (11469, -0.15 DPS) [quest]; Nethergeld Cuffs (254061, -0.15 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.40 DPS) [quest] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 32.3 spell_power points (4.26 DPS) | yes | Raider Handwraps (272098, -0.37 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.33 DPS) [vendor]; Dreamweave Gloves (10019, -1.34 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 25.4 spell_power points (3.35 DPS) | yes | Satyrmane Sash (17755, -0.16 DPS) [dungeon]; Ban'thok Sash (11662, -0.23 DPS) [dungeon]; Deathmage Sash (10771, -0.41 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 33.2 spell_power points (4.38 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.10 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.56 DPS) [dungeon]; Red Mageweave Pants (10009, -3.17 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.17 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.36 DPS) [vendor]; Gilded Sandals (254107, -0.50 DPS) [crafted]; Southsea Mojo Boots (20641, -0.63 DPS) [quest] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.1 spell_power points (2.13 DPS) | yes | Brainlash (6440, -0.11 DPS) [dungeon]; Band of the Unicorn (7553, -0.41 DPS) [world_drop]; Mindseye Circle (10634, -0.51 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (2.13 DPS) | yes | Brainlash (6440, -0.11 DPS) [dungeon]; Band of the Unicorn (7553, -0.41 DPS) [world_drop]; Mindseye Circle (10634, -0.51 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, -1.21 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.81 DPS) [quest]; Spellforce Rod (1664, -1.26 DPS) [world]; Barman Shanker (12791, -6.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 397.8 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Virtuous Hands; waist: Dawnspire Cord; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524111001300000000-00000000000000000-543110401201300251)

Set DPS (verified): 291.4. Weights run: 1.1s. Verify run: 0.8s. 1129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.995 ± 0.029, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.599 per %), hit=0.748 ± 0.009 per rating point (10 rating = 1%, 7.478 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 45.5 spell_power points (6.87 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Magister's Crown (16686, -4.52 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 27.9 spell_power points (4.22 DPS) | yes | Orb of the Darkmoon (19426, -0.90 DPS) [quest]; Chains of the Lich (23125, -0.90 DPS) [dungeon]; Beads of Ogre Mojo (22149, -1.93 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 43.5 spell_power points (6.57 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.70 DPS) [vendor]; Virtuous Epaulets (226955, -0.77 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.4 spell_power points (4.75 DPS) | yes | Hide of the Wild (18510, -1.13 DPS) [crafted]; Spritecaster Cape (11623, -1.73 DPS) [dungeon]; Crystalline Threaded Cape (20697, -3.30 DPS, sim-verified) [world] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 43.9 spell_power points (6.63 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Robe of Everlasting Night (18385, -2.89 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 30.0 spell_power points (4.52 DPS) | yes | Sublime Wristguards (18497, -1.21 DPS) [dungeon]; Runecloth Cuffs (254123, -1.36 DPS) [crafted]; Virtuous Wraps (226953, -1.66 DPS) [vendor] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 34.9 spell_power points (5.27 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Hands of Power (13253, -1.94 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 54.4 spell_power points (8.21 DPS) | yes | Stormpike Cloth Girdle (19094, -3.99 DPS) [rep]; Magister's Belt (16685, -4.00 DPS) [dungeon]; Belt of the Archmage (18405, -6.98 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (22752) | Silverwing Sentinels [rep] | 46.9 spell_power points (7.08 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Skyshroud Leggings (13170, -0.75 DPS) [dungeon] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 37.9 spell_power points (5.73 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -3.37 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (291.4 DPS) | yes | Songstone of Ironforge (12543, -1.21 DPS) [quest]; Maiden's Circle (13001, -1.21 DPS) [world_drop]; Naglering (11669, -9.88 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (291.4 DPS) | yes | Songstone of Ironforge (12543, -0.45 DPS) [quest]; Maiden's Circle (13001, -0.45 DPS) [world_drop]; Naglering (11669, -11.10 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (291.4 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (291.4 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -1.50 DPS, sim-verified) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (291.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.05 DPS) [world]; Hand of Edward the Odd (2243, -9.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 514.4 spell_power points (77.67 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -3.52 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.77 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.23 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Second Wind; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 30.3. Weights run: 0.9s. Verify run: 0.8s. 140 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.233 ± 0.006, crit=0.105 ± 0.003 per rating point (14 rating = 1%, 1.467 per %), hit=0.326 ± 0.002 per rating point (10 rating = 1%, 3.261 per %), spell_haste=not significant (-0.280 ± 0.131), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -0.72 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.1 spell_power points (0.68 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.29 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.31 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.10 DPS) [rep]; Black Whelp Cloak (7283, -0.14 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.2 spell_power points (0.59 DPS) | yes | Green Woolen Vest (2582, -0.21 DPS) [crafted]; Bloody Apron (6226, -0.21 DPS) [dungeon]; Gray Woolen Robe (2585, -0.53 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | sim-verified (30.3 DPS) | yes | Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Owlbeard Bracers (16981, -0.38 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.67 DPS) | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.22 DPS) [crafted]; Apothecary Gloves (10919, -0.29 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.47 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.29 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.37 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.9 spell_power points (1.04 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.37 DPS) [dungeon]; Rumpled Kilt (274741, -0.56 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.76 DPS) | yes | Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.40 DPS) [crafted]; Feather Padded Treads (285345, -0.62 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.48 DPS) | yes | Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Loop of Sacrifice (281673, -0.37 DPS) [quest]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.29 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.18 DPS) [quest]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.3 spell_power points (0.22 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 237.1 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.42 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Sizzle Stick (8071, -4.47 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 000000000000000000-00000000000000000-543110401200000000)

Set DPS (verified): 49.8. Weights run: 0.9s. Verify run: 0.7s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.401 ± 0.009, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.587 per %), hit=0.191 ± 0.002 per rating point (10 rating = 1%, 1.908 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.30 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.24 DPS) [crafted]; Embalmed Shroud (7691, -0.35 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (1.11 DPS) | yes | Crystal Starfire Medallion (5003, -0.92 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.92 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.09 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (1.49 DPS) | yes | Death Speaker Mantle (6685, -0.25 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.59 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Battle Healer's Cloak (19529, -0.12 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.2 spell_power points (1.68 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS) [dungeon]; Pristine Gown (253961, -0.52 DPS) [crafted]; Tree Bark Jacket (1486, -1.02 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.06 DPS) | yes | Glowing Magical Bracelets (13106, -0.77 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.78 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.78 DPS) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.0 spell_power points (0.95 DPS) | yes | Truefaith Gloves (7049, -0.21 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS) [world]; Serpent Gloves (5970, -0.35 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.2 spell_power points (1.44 DPS) | yes | Belt of Arugal (6392, -0.24 DPS) [dungeon]; Warsong Sash (16975, -0.35 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.38 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 spell_power points (1.44 DPS) | yes | Gaze Dreamer Pants (6903, -0.02 DPS) [dungeon]; Pristine Leggings (253987, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.8 spell_power points (1.16 DPS) | yes | Acidic Walkers (9454, -0.19 DPS) [dungeon]; Boots of the Enchanter (4325, -0.57 DPS) [crafted]; Spidersilk Boots (4320, -1.25 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.83 DPS) | yes | Electrocutioner Lagnut (9447, -0.47 DPS) [dungeon]; Sludge-Stained Band (286535, -0.47 DPS) [world]; Black Widow Band (6199, -0.50 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.71 DPS) | yes | Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Black Widow Band (6199, -0.38 DPS) [world]; Sludge-Stained Band (286535, -1.03 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.06 DPS) | yes | Glimmering Staff (249392, -0.39 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.59 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.59 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.4 spell_power points (1.11 DPS) | yes | Orb of Souls (249395, -0.64 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.69 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.18 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 284.1 spell_power points (33.56 DPS) | yes | Starfaller (13063, -0.51 DPS) [world_drop]; Greater Mystic Wand (217287, -3.97 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-543110401201300240)

Set DPS (verified): 81.7. Weights run: 1.0s. Verify run: 0.7s. 311 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.711 ± 0.015, crit=0.040 ± 0.002 per rating point (14 rating = 1%, 0.554 per %), hit=0.289 ± 0.003 per rating point (10 rating = 1%, 2.889 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.79 DPS) | yes | Augural Shroud (2620, -0.38 DPS) [world]; Corpseshroud (10574, -1.00 DPS) [dungeon]; Enchanter's Cowl (4322, -1.05 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.3 spell_power points (1.50 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.55 DPS) [quest]; Triune Amulet (7722, -0.84 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.84 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (2.16 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.4 spell_power points (2.05 DPS) | yes | Guardian Cloak (5965, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.01 DPS) [vendor]; Long Silken Cloak (4326, -1.50 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.3 spell_power points (3.50 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.50 DPS) [crafted]; Elemental Raiment (9434, -0.70 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.7 spell_power points (1.29 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.36 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (2.77 DPS) | yes | Red Mageweave Gloves (10018, -0.49 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.78 DPS) [crafted]; Gilded Handwraps (254021, -1.05 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.7 spell_power points (2.35 DPS) | yes | Defiler's Cloth Girdle (20166, -0.47 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.53 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (3.00 DPS) | yes | Crimson Silk Pantaloons (7062, -0.49 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.04 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.31 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.19 DPS) | yes | Gilded Slippers (254001, -0.62 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.77 DPS) [dungeon]; Spidersilk Boots (4320, -1.88 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.90 DPS) | yes | Reedknot Ring (9622, -0.97 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.10 DPS) [vendor]; Black Widow Band (6199, -1.24 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.20 DPS) | yes | Sea Giant's Toe Ring (274746, -0.40 DPS) [vendor]; Black Widow Band (6199, -0.53 DPS) [world]; Reedknot Ring (9622, -1.20 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (81.7 DPS) | yes | Windweaver Staff (7757, -1.24 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.46 DPS) [dungeon]; Gut Ripper (2164, -3.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 300.2 spell_power points (39.95 DPS) | yes | Umbral Wand (5216, -0.92 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.36 DPS) [dungeon]; Twisted Nether Wand (249144, -5.15 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 521000000000000000-00000000000000000-543110401201300251)

Set DPS (verified): 126.9. Weights run: 1.0s. Verify run: 0.8s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.020 ± 0.019, crit=0.047 ± 0.002 per rating point (14 rating = 1%, 0.661 per %), hit=0.408 ± 0.006 per rating point (10 rating = 1%, 4.077 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 39.4 spell_power points (5.20 DPS) | yes | Dreamweave Circlet (10041, -0.60 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.57 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -1.64 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.3 spell_power points (1.88 DPS) | yes | Scorn's Icy Choker (23169, -0.15 DPS) [dungeon]; Mindburst Medallion (11196, -0.28 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.54 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 31.4 spell_power points (4.14 DPS) | yes | Kentic Amice (11624, -0.54 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.20 DPS) [crafted]; Inquisitor's Shawl (19507, -1.46 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.2 spell_power points (2.79 DPS) | yes | Spritecaster Cape (11623, -0.14 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.40 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 39.4 spell_power points (5.20 DPS) | yes | Runecloth Tunic (13857, -1.48 DPS) [crafted]; Robe of the Magi (1716, -1.49 DPS) [world_drop]; Runecloth Robe (13858, -1.67 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 15.3 spell_power points (2.02 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.15 DPS) [quest] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 32.3 spell_power points (4.26 DPS) | yes | Raider Handwraps (272098, -0.37 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.33 DPS) [vendor]; Dreamweave Gloves (10019, -1.34 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 25.4 spell_power points (3.35 DPS) | yes | Satyrmane Sash (17755, -0.16 DPS) [dungeon]; Ban'thok Sash (11662, -0.23 DPS) [dungeon]; Deathmage Sash (10771, -0.41 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 33.2 spell_power points (4.38 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.10 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.56 DPS) [dungeon]; Red Mageweave Pants (10009, -2.48 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.17 DPS) | yes | Gilded Sandals (254107, -0.50 DPS) [crafted]; Southsea Mojo Boots (20641, -0.63 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -0.64 DPS, sim-verified) [vendor] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.1 spell_power points (2.13 DPS) | yes | Brainlash (6440, -0.11 DPS) [dungeon]; Band of the Unicorn (7553, -0.41 DPS) [world_drop]; Mindseye Circle (10634, -0.51 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (2.13 DPS) | yes | Brainlash (6440, -0.11 DPS) [dungeon]; Band of the Unicorn (7553, -0.41 DPS) [world_drop]; Mindseye Circle (10634, -0.51 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (126.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (126.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (126.9 DPS) | yes | Spellshifter Rod (9527, -0.81 DPS) [quest]; Spellforce Rod (1664, -1.26 DPS) [world]; Barman Shanker (12791, -5.83 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 397.8 spell_power points (52.50 DPS) | yes | Woestave (20082, -1.18 DPS) [quest]; Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Virtuous Hands; waist: Dawnspire Cord; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524111001300000000-00000000000000000-543110401201300251)

Set DPS (verified): 274.1. Weights run: 1.1s. Verify run: 0.8s. 1120 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.995 ± 0.029, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.599 per %), hit=0.748 ± 0.009 per rating point (10 rating = 1%, 7.478 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 45.5 spell_power points (6.87 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -1.75 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 27.9 spell_power points (4.22 DPS) | yes | Beads of Ogre Mojo (22149, -0.45 DPS) [quest]; Orb of the Darkmoon (19426, -0.90 DPS) [quest]; Chains of the Lich (23125, -0.90 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 43.5 spell_power points (6.57 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.70 DPS) [vendor]; Virtuous Epaulets (226955, -0.77 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.4 spell_power points (4.75 DPS) | yes | Crystalline Threaded Cape (20697, -1.13 DPS) [world]; Hide of the Wild (18510, -1.13 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.58 DPS) [quest] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 43.9 spell_power points (6.63 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robe of Everlasting Night (18385, -0.60 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 30.0 spell_power points (4.52 DPS) | yes | Sublime Wristguards (18497, -1.21 DPS) [dungeon]; Runecloth Cuffs (254123, -1.36 DPS) [crafted]; Virtuous Wraps (226953, -1.66 DPS) [vendor] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 34.9 spell_power points (5.27 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Hands of Power (13253, -0.44 DPS) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 54.4 spell_power points (8.21 DPS) | yes | Frostwolf Cloth Belt (19090, -3.99 DPS) [rep]; Magister's Belt (16685, -4.00 DPS) [dungeon]; Belt of the Archmage (18405, -4.17 DPS, sim-verified) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 46.9 spell_power points (7.08 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.75 DPS) [dungeon]; Sentinel's Silk Leggings (237815, -3.79 DPS, sim-verified) [vendor] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 37.9 spell_power points (5.73 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -0.60 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (274.1 DPS) | yes | Eye of Orgrimmar (12545, -1.21 DPS) [quest]; Maiden's Circle (13001, -1.21 DPS) [world_drop]; Naglering (11669, -7.69 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (274.1 DPS) | yes | Eye of Orgrimmar (12545, -0.45 DPS) [quest]; Maiden's Circle (13001, -0.45 DPS) [world_drop]; Naglering (11669, -7.11 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (274.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (274.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -2.05 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (274.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.34 DPS) [dungeon]; Hand of Edward the Odd (2243, -6.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 514.4 spell_power points (77.67 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.16 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.77 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.23 DPS) [world] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Outrider's Silk Leggings; feet: Virtuous Slippers; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

