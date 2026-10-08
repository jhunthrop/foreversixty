# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 29.2. Weights run: 1.0s. Verify run: 0.7s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.216 ± 0.006, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.995 per %), hit=0.223 ± 0.006 per rating point (10 rating = 1%, 2.235 per %), spell_haste=0.592 ± 0.082, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.62 DPS) | yes | Shadow Goggles (4373, -1.20 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.9 spell_power points (0.72 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.30 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.61 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.41 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Caretaker's Cape (20428, -0.10 DPS) [rep]; Black Whelp Cloak (7283, -0.46 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.1 spell_power points (0.63 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -0.93 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.1 spell_power points (0.11 DPS) | yes | Bright Bracers (3647, -0.02 DPS) [world_drop]; Repurposed Hair Band (281256, -0.07 DPS) [quest]; Windsong Bangles (263336, -0.75 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.73 DPS) | yes | Gnoll Casting Gloves (892, -0.10 DPS) [world]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.47 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.50 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.33 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.72 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.7 spell_power points (1.11 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.81 DPS) | yes | Red Woolen Boots (4313, -0.40 DPS) [crafted]; Pristine Boots (253889, -0.44 DPS) [crafted]; Feather Padded Treads (285345, -0.53 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.4 spell_power points (0.56 DPS) | yes | Sludge-Stained Band (286535, -0.25 DPS) [world]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.52 DPS) | yes | Lavishly Jeweled Ring (1156, -0.38 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.45 DPS) [world_drop]; Sludge-Stained Band (286535, -0.67 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.2 spell_power points (0.22 DPS) | yes | Lesser Staff of the Spire (1300, -0.09 DPS) [world]; Staff of Westfall (2042, -0.11 DPS) [quest]; Channeler's Staff (4437, -0.43 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 218.3 spell_power points (22.62 DPS) | yes | Skycaller (12984, -0.97 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.33 DPS) [dungeon]; Deepblaze (279896, -4.09 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-00000000000000000-2535101301000000000)

Set DPS (verified): 40.6. Weights run: 1.1s. Verify run: 0.8s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.137 ± 0.005, crit=0.098 ± 0.003 per rating point (14 rating = 1%, 1.373 per %), hit=0.240 ± 0.006 per rating point (10 rating = 1%, 2.402 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.20 DPS) | yes | Silk Headband (7050, -0.23 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.33 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.33 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.8 spell_power points (0.85 DPS) | yes | Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.89 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.2 spell_power points (1.11 DPS) | yes | Death Speaker Mantle (6685, -0.30 DPS) [dungeon]; Fairywing Mantle (9536, -0.33 DPS) [quest]; Invoker's Mantle (215365, -0.34 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.54 DPS) | yes | Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Prelacy Cape (7004, -0.11 DPS) [quest]; Caretaker's Cape (19533, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (40.6 DPS) | yes | Death Speaker Robes (6682, -0.25 DPS) [dungeon]; Robes of Arcana (5770, -0.30 DPS) [crafted]; Tree Bark Jacket (1486, -0.77 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Glowing Magical Bracelets (13106, -0.58 DPS, sim-verified) [world_drop]; Windsong Bangles (263336, -0.87 DPS) [quest]; Nightsky Wristbands (6407, -0.89 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (0.76 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.11 DPS) [world]; Town Clerk's Mittens (270029, -0.16 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.4 spell_power points (1.24 DPS) | yes | Belt of Arugal (6392, -0.19 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.37 DPS) [dungeon]; Invoker's Cord (215366, -0.40 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.30 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.44 DPS) [crafted]; Silk-threaded Trousers (1929, -0.54 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (0.86 DPS) | yes | Acidic Walkers (9454, -0.20 DPS) [dungeon]; Nimbus Boots (6998, -0.21 DPS) [quest]; Spidersilk Boots (4320, -0.81 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.76 DPS) | yes | Minor Channeling Ring (1449, -0.19 DPS) [quest]; Electrocutioner Lagnut (9447, -0.43 DPS) [dungeon]; Sludge-Stained Band (286535, -0.43 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.65 DPS) | yes | Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Sludge-Stained Band (286535, -0.33 DPS) [world]; Minor Channeling Ring (1449, -0.67 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Glimmering Staff (249392, -0.30 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.83 DPS) [world_drop]; Channeler's Staff (4437, -0.86 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.8 spell_power points (0.85 DPS) | yes | Eye of Paleth (2943, -0.42 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.42 DPS) [world]; Dwarven Tome (279898, -0.72 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 308.7 spell_power points (33.54 DPS) | yes | Starfaller (13063, -0.62 DPS) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.54 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 87.7. Weights run: 1.0s. Verify run: 0.7s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.638 ± 0.024, crit=0.179 ± 0.008 per rating point (14 rating = 1%, 2.503 per %), hit=0.406 ± 0.025 per rating point (10 rating = 1%, 4.064 per %), spell_haste=not significant (0.742 ± 0.384), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.93 DPS) | yes | Augural Shroud (2620, -0.51 DPS) [world]; Living Cowl (5608, -1.12 DPS) [world]; Enchanter's Cowl (4322, -1.20 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.51 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.62 DPS) [quest]; Triune Amulet (7722, -0.89 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.89 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.3 spell_power points (2.14 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.7 spell_power points (2.06 DPS) | yes | Guardian Cloak (5965, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.08 DPS) [vendor]; Long Silken Cloak (4326, -1.55 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.8 spell_power points (3.61 DPS) | yes | Dreamweave Vest (10021, -0.29 DPS) [crafted]; Robe of Power (7054, -0.58 DPS) [crafted]; Elemental Raiment (9434, -0.67 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.26 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.28 DPS) [quest]; Windchaser Cuffs (14429, -0.46 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (2.87 DPS) | yes | Black Mageweave Gloves (10003, -0.78 DPS) [crafted]; Red Mageweave Gloves (10018, -0.97 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.13 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.6 spell_power points (2.31 DPS) | yes | Highlander's Cloth Girdle (20098, -0.00 DPS) [rep]; Gilded Cord (254037, -0.48 DPS) [crafted]; Star Belt (4329, -0.50 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.7 spell_power points (3.02 DPS) | yes | Crimson Silk Pantaloons (7062, -0.95 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.05 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.33 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.35 DPS) | yes | Gilded Slippers (254001, -1.16 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.94 DPS) [dungeon]; Spidersilk Boots (4320, -2.02 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (1.93 DPS) | yes | Ring of Forlorn Spirits (2043, -0.81 DPS) [quest]; Reedknot Ring (9622, -0.95 DPS) [quest]; Minor Channeling Ring (1449, -1.05 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.26 DPS) | yes | Reedknot Ring (9622, -0.28 DPS) [quest]; Minor Channeling Ring (1449, -0.38 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.29 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (87.7 DPS) | yes | Windweaver Staff (7757, -1.46 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.54 DPS) [dungeon]; Gut Ripper (2164, -4.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 285.8 spell_power points (39.91 DPS) | yes | Nether Force Wand (11263, -1.36 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.46 DPS) [quest]; Ragefire Wand (7513, -2.51 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 203005000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 196.7. Weights run: 0.9s. Verify run: 0.7s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.924 ± 0.039, crit=0.279 ± 0.011 per rating point (14 rating = 1%, 3.904 per %), hit=0.857 ± 0.051 per rating point (10 rating = 1%, 8.572 per %), spell_haste=5.800 ± 0.683, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.5 spell_power points (7.50 DPS) | yes | Dreamweave Circlet (10041, -1.45 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.70 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.10 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (4.31 DPS) | yes | Scorn's Icy Choker (23169, -1.80 DPS) [dungeon]; Mindburst Medallion (11196, -2.00 DPS) [quest]; Horizon Choker (13085, -5.23 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.6 spell_power points (5.93 DPS) | yes | Kentic Amice (11624, -0.72 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.75 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.88 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 spell_power points (3.91 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.45 DPS) [dungeon]; Runecloth Cloak (13860, -0.63 DPS) [crafted]; Big Voodoo Cloak (8216, -1.25 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.5 spell_power points (7.50 DPS) | yes | Robe of the Magi (1716, -1.99 DPS) [world_drop]; Runecloth Tunic (13857, -2.06 DPS) [crafted]; Runecloth Robe (13858, -2.15 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.9 spell_power points (2.77 DPS) | yes | Nethergeld Cuffs (254061, -0.08 DPS) [crafted]; Bloodband Bracers (11469, -0.11 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.55 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 33.5 spell_power points (6.70 DPS) | yes | Dreamweave Gloves (10019, -2.36 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.44 DPS) [vendor]; Raider Handwraps (272098, -4.04 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 30.7 spell_power points (6.15 DPS) | yes | Satyrmane Sash (17755, -1.50 DPS) [dungeon]; Deathmage Sash (10771, -1.97 DPS) [dungeon]; Dawnspire Cord (12466, -6.08 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.2 spell_power points (6.45 DPS) | yes | Red Mageweave Pants (10009, -1.43 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.45 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.87 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 24.9 spell_power points (4.98 DPS) | yes | Earthen Silk Slippers (254013, -0.18 DPS) [crafted]; Gilded Sandals (254107, -1.11 DPS) [crafted]; Southsea Mojo Boots (20641, -1.34 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (3.11 DPS) | yes | Brainlash (6440, -0.34 DPS) [dungeon]; Band of the Unicorn (7553, -0.51 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.71 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.5 spell_power points (3.09 DPS) | yes | Brainlash (6440, -0.32 DPS) [dungeon]; Band of the Unicorn (7553, -0.49 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.69 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (196.7 DPS) | yes | Uther's Strength (11302, -4.07 DPS, sim-verified) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (196.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (196.7 DPS) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -3.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 262.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.50 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.15 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 203005000000000000-11302300000000000-2535101301000300250)

Set DPS (verified): 323.6. Weights run: 1.0s. Verify run: 0.7s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.747 ± 0.048, crit=0.365 ± 0.013 per rating point (14 rating = 1%, 5.108 per %), hit=1.194 ± 0.063 per rating point (10 rating = 1%, 11.937 per %), spell_haste=7.179 ± 0.845, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 44.8 spell_power points (9.48 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.11 DPS) [pvp]; Crimson Felt Hat (18727, -1.86 DPS) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (323.6 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.47 DPS) [quest]; Chains of the Lich (23125, -1.04 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 43.3 spell_power points (9.16 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.50 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.36 DPS) [vendor]; Burial Shawl (18681, -2.40 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.9 spell_power points (7.17 DPS) | yes | Crystalline Threaded Cape (20697, -2.31 DPS) [world]; Hide of the Wild (18510, -2.63 DPS) [crafted]; Shroud of Arcane Mastery (22330, -2.91 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 54.1 spell_power points (11.44 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.69 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.23 DPS) [pvp]; Robe of Everlasting Night (18385, -3.67 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.0 spell_power points (5.92 DPS) | yes | Sublime Wristguards (18497, -1.80 DPS) [dungeon]; Runecloth Cuffs (254123, -2.01 DPS) [crafted]; Sorcerer's Bindings (226929, -2.86 DPS) [quest] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 34.4 spell_power points (7.28 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.4 spell_power points (11.29 DPS) | yes | Magician's Cord (272393, -3.48 DPS) [vendor]; Ban'thok Sash (11662, -4.49 DPS) [dungeon]; Belt of the Archmage (18405, -6.65 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.7 spell_power points (10.51 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -1.44 DPS) [quest]; Skyshroud Leggings (13170, -2.05 DPS) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 33.0 spell_power points (6.97 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.63 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (323.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.84 DPS) [quest]; Maiden's Circle (13001, -1.48 DPS) [world_drop]; Naglering (11669, -11.29 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (323.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.36 DPS) [quest]; Maiden's Circle (13001, -1.00 DPS) [world_drop]; Naglering (11669, -14.09 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (323.6 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (323.6 DPS) | yes | Weakness Analyzer (272438, -1.48 DPS) [vendor]; Serenity Field (272439, -3.17 DPS) [vendor]; Burst of Knowledge (11832, -3.60 DPS) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (323.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Crackling Staff (19102, -1.21 DPS) [rep]; Teebu's Blazing Longsword (1728, -24.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 368.8 spell_power points (78.03 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.85 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.44 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.22 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 203005000000000000-13102000000000000-2555100321000300250)

Set DPS (verified): 560.1. Weights run: 1.0s. Verify run: 0.8s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.272 ± 0.019, crit=0.430 ± 0.013 per rating point (14 rating = 1%, 6.019 per %), hit=1.055 ± 0.043 per rating point (10 rating = 1%, 10.553 per %), spell_haste=7.238 ± 0.333, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 33.8 spell_power points (13.96 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.79 DPS) [pvp] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (560.1 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -0.11 DPS) [dungeon]; Jewel of Kajaro (19601, -4.60 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.1 spell_power points (15.32 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.31 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.68 DPS, sim-verified) [crafted]; Argent Shoulders (19059, -4.99 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.7 spell_power points (11.86 DPS) | yes | Crystalline Threaded Cape (20697, -3.16 DPS) [world]; Amplifying Cloak (18350, -4.43 DPS) [dungeon]; Hide of the Wild (18510, -4.96 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 49.3 spell_power points (20.35 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.33 DPS) [pvp]; Robe of Everlasting Night (18385, -5.56 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Tunic (227108, -7.29 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.2 spell_power points (9.98 DPS) | yes | Sublime Wristguards (18497, -3.91 DPS) [dungeon]; Runecloth Cuffs (254123, -4.32 DPS) [crafted]; Sorcerer's Bindings (226929, -5.97 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.4 spell_power points (11.71 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 41.5 spell_power points (17.15 DPS) | yes | Belt of the Archmage (18405, -5.46 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.60 DPS) [dungeon]; Magician's Cord (272393, -6.61 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.8 spell_power points (19.31 DPS) | yes | Marshal's Silk Leggings (231605, -2.19 DPS) [pvp]; Skyshroud Leggings (13170, -4.37 DPS) [dungeon]; Sorcerer's Leggings (226933, -5.31 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 25.3 spell_power points (10.47 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.56 DPS) [crafted] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (560.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.51 DPS) [quest]; Blessed Band of Light (272407, -2.48 DPS) [vendor]; Naglering (11669, -16.23 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (560.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.86 DPS) [quest]; Blessed Band of Light (272407, -1.83 DPS) [vendor]; Naglering (11669, -10.99 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (560.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (560.1 DPS) | yes | Weakness Analyzer (272438, -2.89 DPS) [vendor]; Serenity Field (272439, -4.96 DPS, sim-verified) [vendor]; Blackhand's Breadth (13965, -7.00 DPS) [quest] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (560.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -2.29 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -25.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 191.9 spell_power points (79.24 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.85 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.62 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.83 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 27.0. Weights run: 1.0s. Verify run: 0.7s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.216 ± 0.006, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.995 per %), hit=0.223 ± 0.006 per rating point (10 rating = 1%, 2.235 per %), spell_haste=0.592 ± 0.082, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.62 DPS) | yes | Shadow Goggles (4373, -1.22 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.9 spell_power points (0.72 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.30 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.47 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.41 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.10 DPS) [rep]; Black Whelp Cloak (7283, -0.29 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.1 spell_power points (0.63 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -0.90 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.4 spell_power points (0.15 DPS) | yes | Tabitha's Cuffs (251486, -0.01 DPS) [quest]; Mindthrust Bracers (1974, -0.04 DPS) [dungeon]; Featherbead Bracers (15452, -0.04 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.73 DPS) | yes | Gnoll Casting Gloves (892, -0.10 DPS) [world]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Apothecary Gloves (10919, -0.31 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.50 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.33 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.77 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.7 spell_power points (1.11 DPS) | yes | Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted]; Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.81 DPS) | yes | Red Woolen Boots (4313, -0.40 DPS) [crafted]; Pristine Boots (253889, -0.44 DPS) [crafted]; Feather Padded Treads (285345, -0.79 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.52 DPS) | yes | Lavishly Jeweled Ring (1156, -0.38 DPS) [dungeon]; Loop of Sacrifice (281673, -0.41 DPS) [quest]; Volcanic Rock Ring (12053, -0.45 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.31 DPS) | yes | Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Loop of Sacrifice (281673, -0.20 DPS) [quest]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.2 spell_power points (0.22 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.09 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 218.3 spell_power points (22.62 DPS) | yes | Skycaller (12984, -1.02 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.33 DPS) [dungeon]; Sizzle Stick (8071, -4.45 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 000000000000000000-00000000000000000-2535101301000000000)

Set DPS (verified): 39.1. Weights run: 1.1s. Verify run: 0.8s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.137 ± 0.005, crit=0.098 ± 0.003 per rating point (14 rating = 1%, 1.373 per %), hit=0.240 ± 0.006 per rating point (10 rating = 1%, 2.402 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.20 DPS) | yes | Embalmed Shroud (7691, -0.33 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.33 DPS) [crafted]; Silk Headband (7050, -0.37 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.8 spell_power points (0.85 DPS) | yes | Crystal Starfire Medallion (5003, -0.79 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.00 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.2 spell_power points (1.11 DPS) | yes | Invoker's Mantle (215365, -0.28 DPS) [crafted]; Death Speaker Mantle (6685, -0.30 DPS) [dungeon]; Chestnut Mantle (17695, -0.85 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.54 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (39.1 DPS) | yes | Death Speaker Robes (6682, -0.25 DPS) [dungeon]; High Robe of the Adjudicator (3461, -0.27 DPS) [quest]; Tree Bark Jacket (1486, -0.39 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Glowing Magical Bracelets (13106, -0.86 DPS) [world_drop]; Windsong Bangles (263336, -0.87 DPS) [quest]; Owlbeard Bracers (16981, -1.45 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.76 DPS) | yes | Jutebraid Gloves (10654, -0.03 DPS) [quest]; Gnoll Casting Gloves (892, -0.11 DPS) [world]; Truefaith Gloves (7049, -0.17 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.4 spell_power points (1.24 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.37 DPS) [dungeon]; Warsong Sash (16975, -0.45 DPS, sim-verified) [quest] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.30 DPS) | yes | Abomination Skin Leggings (23173, -0.21 DPS) [dungeon]; Pristine Leggings (253987, -0.44 DPS) [crafted]; Silk-threaded Trousers (1929, -0.54 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.0 spell_power points (0.86 DPS) | yes | Acidic Walkers (9454, -0.20 DPS) [dungeon]; Boots of the Enchanter (4325, -0.32 DPS) [crafted]; Spidersilk Boots (4320, -1.22 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.76 DPS) | yes | Electrocutioner Lagnut (9447, -0.43 DPS) [dungeon]; Sludge-Stained Band (286535, -0.43 DPS) [world]; Sacred Band (6669, -0.54 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.65 DPS) | yes | Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Sacred Band (6669, -0.43 DPS) [quest]; Sludge-Stained Band (286535, -1.05 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.98 DPS) | yes | Glimmering Staff (249392, -0.72 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.83 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.83 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.8 spell_power points (0.85 DPS) | yes | Orb of Souls (249395, -0.42 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.57 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -0.95 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 308.7 spell_power points (33.54 DPS) | yes | Starfaller (13063, -0.24 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.99 DPS) [crafted]; Gravestone Scepter (7001, -4.54 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 83.7. Weights run: 1.0s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.638 ± 0.024, crit=0.179 ± 0.008 per rating point (14 rating = 1%, 2.503 per %), hit=0.406 ± 0.025 per rating point (10 rating = 1%, 4.064 per %), spell_haste=not significant (0.742 ± 0.384), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.93 DPS) | yes | Augural Shroud (2620, -0.51 DPS) [world]; Living Cowl (5608, -1.12 DPS) [world]; Enchanter's Cowl (4322, -1.20 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.51 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.62 DPS) [quest]; Triune Amulet (7722, -0.89 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.89 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.3 spell_power points (2.14 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Bloodmage Mantle (7684, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.7 spell_power points (2.06 DPS) | yes | Guardian Cloak (5965, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.08 DPS) [vendor]; Long Silken Cloak (4326, -1.67 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.8 spell_power points (3.61 DPS) | yes | Dreamweave Vest (10021, -0.29 DPS) [crafted]; Robe of Power (7054, -0.58 DPS) [crafted]; Elemental Raiment (9434, -0.67 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.1 spell_power points (1.27 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.01 DPS) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (2.87 DPS) | yes | Black Mageweave Gloves (10003, -0.78 DPS) [crafted]; Gilded Handwraps (254021, -1.13 DPS) [crafted]; Red Mageweave Gloves (10018, -1.40 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.6 spell_power points (2.31 DPS) | yes | Defiler's Cloth Girdle (20166, -0.00 DPS) [rep]; Gilded Cord (254037, -0.48 DPS) [crafted]; Star Belt (4329, -0.50 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.7 spell_power points (3.02 DPS) | yes | Crimson Silk Pantaloons (7062, -0.85 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.05 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.33 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.35 DPS) | yes | Gilded Slippers (254001, -1.15 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.94 DPS) [dungeon]; Spidersilk Boots (4320, -2.02 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (1.93 DPS) | yes | Reedknot Ring (9622, -0.95 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.09 DPS) [vendor]; Black Widow Band (6199, -1.31 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.26 DPS) | yes | Sea Giant's Toe Ring (274746, -0.42 DPS) [vendor]; Black Widow Band (6199, -0.63 DPS) [world]; Reedknot Ring (9622, -1.44 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (83.7 DPS) | yes | Windweaver Staff (7757, -1.46 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.54 DPS) [dungeon]; Gut Ripper (2164, -4.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 285.8 spell_power points (39.91 DPS) | yes | Nether Force Wand (11263, -1.36 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.46 DPS) [quest]; Ragefire Wand (7513, -2.51 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 203005000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 190.7. Weights run: 0.9s. Verify run: 0.7s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.924 ± 0.039, crit=0.279 ± 0.011 per rating point (14 rating = 1%, 3.904 per %), hit=0.857 ± 0.051 per rating point (10 rating = 1%, 8.572 per %), spell_haste=5.800 ± 0.683, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.5 spell_power points (7.50 DPS) | yes | Dreamweave Circlet (10041, -1.45 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.70 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.10 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (4.31 DPS) | yes | Scorn's Icy Choker (23169, -1.80 DPS) [dungeon]; Mindburst Medallion (11196, -2.00 DPS) [quest]; Horizon Choker (13085, -4.97 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.6 spell_power points (5.93 DPS) | yes | Kentic Amice (11624, -0.72 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.75 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.88 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.3 spell_power points (4.06 DPS) | yes | Spritecaster Cape (11623, -0.15 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.60 DPS) [dungeon]; Runecloth Cloak (13860, -0.78 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.5 spell_power points (7.50 DPS) | yes | Robe of the Magi (1716, -1.99 DPS) [world_drop]; Runecloth Tunic (13857, -2.06 DPS) [crafted]; Runecloth Robe (13858, -2.15 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.9 spell_power points (2.77 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.08 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 33.5 spell_power points (6.70 DPS) | yes | Raider Handwraps (272098, -1.24 DPS) [vendor]; Dreamweave Gloves (10019, -2.36 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.44 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 30.7 spell_power points (6.15 DPS) | yes | Dawnspire Cord (12466, -1.44 DPS) [dungeon]; Satyrmane Sash (17755, -1.50 DPS) [dungeon]; Deathmage Sash (10771, -1.97 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.2 spell_power points (6.45 DPS) | yes | Red Mageweave Pants (10009, -1.43 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.45 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.37 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 24.9 spell_power points (4.98 DPS) | yes | Earthen Silk Slippers (254013, -0.18 DPS) [crafted]; Gilded Sandals (254107, -1.11 DPS) [crafted]; Southsea Mojo Boots (20641, -1.34 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (3.11 DPS) | yes | Brainlash (6440, -0.34 DPS) [dungeon]; Band of the Unicorn (7553, -0.51 DPS) [world_drop]; Advisor's Ring (19519, -0.71 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.5 spell_power points (3.09 DPS) | yes | Brainlash (6440, -0.32 DPS) [dungeon]; Band of the Unicorn (7553, -0.49 DPS) [world_drop]; Advisor's Ring (19519, -0.69 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (190.7 DPS) | yes | Uther's Strength (11302, -0.34 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (190.7 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 262.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.50 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.15 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 203005000000000000-11302300000000000-2535101301000300250)

Set DPS (verified): 311.0. Weights run: 1.0s. Verify run: 0.8s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.747 ± 0.048, crit=0.365 ± 0.013 per rating point (14 rating = 1%, 5.108 per %), hit=1.194 ± 0.063 per rating point (10 rating = 1%, 11.937 per %), spell_haste=7.179 ± 0.845, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 44.8 spell_power points (9.48 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.11 DPS) [pvp]; Crimson Felt Hat (18727, -5.11 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (311.0 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.47 DPS) [quest]; Chains of the Lich (23125, -1.04 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 43.3 spell_power points (9.16 DPS) | yes | Warlord's Silk Amice (231594, -1.50 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.36 DPS) [vendor]; Burial Shawl (18681, -2.40 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.9 spell_power points (7.17 DPS) | yes | Crystalline Threaded Cape (20697, -2.31 DPS) [world]; Hide of the Wild (18510, -2.63 DPS) [crafted]; Shroud of Arcane Mastery (22330, -2.91 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 54.1 spell_power points (11.44 DPS) | yes | Warlord's Silk Raiment (231596, -0.69 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.23 DPS) [pvp]; Robe of Everlasting Night (18385, -3.67 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.0 spell_power points (5.92 DPS) | yes | Sublime Wristguards (18497, -1.80 DPS) [dungeon]; Runecloth Cuffs (254123, -2.01 DPS) [crafted]; Sorcerer's Bindings (226929, -2.86 DPS) [quest] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 34.4 spell_power points (7.28 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.4 spell_power points (11.29 DPS) | yes | Magician's Cord (272393, -3.48 DPS) [vendor]; Ban'thok Sash (11662, -4.49 DPS) [dungeon]; Belt of the Archmage (18405, -5.23 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.7 spell_power points (10.51 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -1.44 DPS) [quest]; Outrider's Silk Leggings (22747, -1.58 DPS) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 33.0 spell_power points (6.97 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.63 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (311.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.84 DPS) [quest]; Maiden's Circle (13001, -1.48 DPS) [world_drop]; Naglering (11669, -13.51 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (311.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.36 DPS) [quest]; Maiden's Circle (13001, -1.00 DPS) [world_drop]; Naglering (11669, -10.73 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (311.0 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (311.0 DPS) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -1.48 DPS) [vendor]; Serenity Field (272439, -3.17 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (311.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -1.66 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -22.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 368.8 spell_power points (78.03 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.85 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.44 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.22 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Diana's Pearl Necklace; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (troll, 203005000000000000-13102000000000000-2555100321000300250)

Set DPS (verified): 561.1. Weights run: 1.0s. Verify run: 0.8s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.272 ± 0.019, crit=0.430 ± 0.013 per rating point (14 rating = 1%, 6.019 per %), hit=1.055 ± 0.043 per rating point (10 rating = 1%, 10.553 per %), spell_haste=7.238 ± 0.333, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 33.8 spell_power points (13.96 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Crimson Felt Hat (18727, -0.68 DPS) [dungeon]; Champion's Silk Cowl (227105, -0.79 DPS) [pvp] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (561.1 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -0.11 DPS) [dungeon]; Jewel of Kajaro (19601, -4.55 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.1 spell_power points (15.32 DPS) | yes | Warlord's Silk Amice (231594, -3.31 DPS) [pvp]; Argent Shoulders (19059, -4.99 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.49 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.7 spell_power points (11.86 DPS) | yes | Crystalline Threaded Cape (20697, -3.16 DPS) [world]; Amplifying Cloak (18350, -4.43 DPS) [dungeon]; Hide of the Wild (18510, -4.96 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 49.3 spell_power points (20.35 DPS) | yes | Warlord's Silk Raiment (231596, -2.33 DPS) [pvp]; Robe of Everlasting Night (18385, -5.70 DPS, sim-verified) [dungeon]; Legionnaire's Silk Tunic (227106, -7.29 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.2 spell_power points (9.98 DPS) | yes | Sublime Wristguards (18497, -3.91 DPS) [dungeon]; Runecloth Cuffs (254123, -4.32 DPS) [crafted]; Sorcerer's Bindings (226929, -5.97 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.4 spell_power points (11.71 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 41.5 spell_power points (17.15 DPS) | yes | Belt of the Archmage (18405, -5.99 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.60 DPS) [dungeon]; Magician's Cord (272393, -6.61 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.8 spell_power points (19.31 DPS) | yes | General's Silk Trousers (231595, -2.19 DPS) [pvp]; Skyshroud Leggings (13170, -2.93 DPS, sim-verified) [dungeon]; Sorcerer's Leggings (226933, -5.31 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 25.3 spell_power points (10.47 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.56 DPS) [crafted] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (561.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.51 DPS) [quest]; Blessed Band of Light (272407, -2.48 DPS) [vendor]; Naglering (11669, -15.67 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (561.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.86 DPS) [quest]; Blessed Band of Light (272407, -1.83 DPS) [vendor]; Naglering (11669, -11.79 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (561.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (561.1 DPS) | yes | Weakness Analyzer (272438, -2.89 DPS) [vendor]; Serenity Field (272439, -4.96 DPS, sim-verified) [vendor]; Blackhand's Breadth (13965, -7.00 DPS) [quest] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (561.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -2.29 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -25.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 191.9 spell_power points (79.24 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.85 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.62 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.83 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

