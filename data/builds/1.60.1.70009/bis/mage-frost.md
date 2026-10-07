# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 32.1. Weights run: 0.9s. Verify run: 0.7s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.196 ± 0.005, crit=0.062 ± 0.002 per rating point (14 rating = 1%, 0.866 per %), hit=0.165 ± 0.001 per rating point (10 rating = 1%, 1.651 per %), spell_haste=0.706 ± 0.082, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.77 DPS) | yes | Shadow Goggles (4373, -1.56 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.8 spell_power points (0.87 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.35 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.45 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.51 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Caretaker's Cape (20428, -0.13 DPS) [rep]; Black Whelp Cloak (7283, -0.36 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.0 spell_power points (0.77 DPS) | yes | Green Woolen Vest (2582, -0.25 DPS) [crafted]; Bloody Apron (6226, -0.25 DPS) [dungeon]; Gray Woolen Robe (2585, -1.01 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (32.1 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Repurposed Hair Band (281256, -0.08 DPS) [quest]; Windsong Bangles (263336, -0.51 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.90 DPS) | yes | Gnoll Casting Gloves (892, -0.13 DPS) [world]; Pristine Gloves (253913, -0.31 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.59 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.8 spell_power points (0.61 DPS) | yes | Novice Ardent's Sash (253887, -0.28 DPS) [crafted]; Keller's Girdle (2911, -0.41 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.60 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.6 spell_power points (1.36 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.46 DPS) [dungeon]; Rumpled Kilt (274741, -0.71 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.8 spell_power points (1.00 DPS) | yes | Red Woolen Boots (4313, -0.49 DPS) [crafted]; Feather Padded Treads (285345, -0.52 DPS, sim-verified) [world]; Pristine Boots (253889, -0.54 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.4 spell_power points (0.69 DPS) | yes | Sludge-Stained Band (286535, -0.31 DPS) [world]; Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.62 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; Sludge-Stained Band (286535, -0.64 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.0 spell_power points (0.25 DPS) | yes | Lesser Staff of the Spire (1300, -0.10 DPS) [world]; Staff of Westfall (2042, -0.13 DPS) [quest]; Channeler's Staff (4437, -0.42 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 176.8 spell_power points (22.70 DPS) | yes | Skycaller (12984, -1.15 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.41 DPS) [dungeon]; Deepblaze (279896, -4.17 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 53.6. Weights run: 0.9s. Verify run: 0.7s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.297 ± 0.008, crit=0.101 ± 0.004 per rating point (14 rating = 1%, 1.411 per %), hit=0.183 ± 0.002 per rating point (10 rating = 1%, 1.826 per %), spell_haste=0.601 ± 0.136, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.65 DPS) | yes | Enchanter's Cowl (4322, -0.31 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon]; Silk Headband (7050, -0.63 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 spell_power points (1.32 DPS) | yes | Crystal Starfire Medallion (5003, -1.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.14 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.44 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (1.75 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS) [dungeon]; Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.48 DPS) [crafted] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.03 DPS) [crafted]; Prelacy Cape (7004, -0.03 DPS) [quest]; Hillman's Cloak (3719, -0.54 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.39 DPS) [dungeon]; Pristine Gown (253961, -0.57 DPS) [crafted]; Tree Bark Jacket (1486, -0.90 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Glowing Magical Bracelets (13106, -0.94 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Stonecloth Bindings (14416, -1.13 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.3 spell_power points (1.09 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world]; Shilly Mitts (9609, -0.39 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 spell_power points (1.78 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Invoker's Cord (215366, -0.51 DPS) [crafted]; Crimson Silk Belt (7055, -0.57 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.80 DPS) | yes | Abomination Skin Leggings (23173, -0.09 DPS) [dungeon]; Pristine Leggings (253987, -0.44 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.63 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.1 spell_power points (1.36 DPS) | yes | Acidic Walkers (9454, -0.26 DPS) [dungeon]; Nimbus Boots (6998, -0.46 DPS) [quest]; Spidersilk Boots (4320, -1.62 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.05 DPS) | yes | Minor Channeling Ring (1449, -0.21 DPS) [quest]; Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon]; Sludge-Stained Band (286535, -0.60 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.90 DPS) | yes | Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Sludge-Stained Band (286535, -0.45 DPS) [world]; Minor Channeling Ring (1449, -0.94 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Glimmering Staff (249392, -0.64 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.91 DPS) [world_drop]; Channeler's Staff (4437, -1.00 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.32 DPS) | yes | Eye of Paleth (2943, -0.72 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.72 DPS) [world]; Dwarven Tome (279898, -1.03 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 224.2 spell_power points (33.66 DPS) | yes | Starfaller (13063, -0.62 DPS) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 100000000000000000-00000000000000000-2535111300000301050)

Set DPS (verified): 86.9. Weights run: 1.0s. Verify run: 0.6s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.516 ± 0.018, crit=0.166 ± 0.007 per rating point (14 rating = 1%, 2.328 per %), hit=0.308 ± 0.003 per rating point (10 rating = 1%, 3.080 per %), spell_haste=1.428 ± 0.283, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.10 DPS) | yes | Augural Shroud (2620, -0.93 DPS, sim-verified) [world]; Living Cowl (5608, -1.18 DPS) [world]; Enchanter's Cowl (4322, -1.45 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.1 spell_power points (1.49 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.73 DPS) [quest]; Triune Amulet (7722, -0.96 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.7 spell_power points (2.02 DPS) | yes | Green Silken Shoulders (7057, -0.00 DPS) [crafted]; Bloodmage Mantle (7684, -0.01 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.6 spell_power points (2.01 DPS) | yes | Guardian Cloak (5965, -0.75 DPS) [crafted]; Icy Cloak (4327, -0.98 DPS) [crafted]; Long Silken Cloak (4326, -1.11 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.1 spell_power points (3.70 DPS) | yes | Dreamweave Vest (10021, -0.36 DPS) [crafted]; Elemental Raiment (9434, -0.60 DPS) [world_drop]; Robe of Power (7054, -0.72 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.33 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest]; Windchaser Cuffs (14429, -0.64 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 spell_power points (2.96 DPS) | yes | Red Mageweave Gloves (10018, -0.58 DPS) [crafted]; Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Gilded Handwraps (254021, -1.25 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.1 spell_power points (2.37 DPS) | yes | Deathmage Sash (10771, -0.19 DPS) [dungeon]; Star Belt (4329, -0.45 DPS) [crafted]; Gilded Cord (254037, -0.58 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.2 spell_power points (2.98 DPS) | yes | Crimson Silk Pantaloons (7062, -1.02 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.04 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.21 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.54 DPS) | yes | Gilded Slippers (254001, -0.96 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.19 DPS) [dungeon]; Spidersilk Boots (4320, -2.20 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (1.93 DPS) | yes | Ring of Forlorn Spirits (2043, -0.75 DPS) [quest]; Reedknot Ring (9622, -0.90 DPS) [quest]; Minor Channeling Ring (1449, -1.04 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.33 DPS) | yes | Reedknot Ring (9622, -0.29 DPS) [quest]; Minor Channeling Ring (1449, -0.44 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (86.9 DPS) | yes | Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Windweaver Staff (7757, -1.81 DPS) [dungeon]; Gut Ripper (2164, -4.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 270.1 spell_power points (39.84 DPS) | yes | Nether Force Wand (11263, -1.92 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.42 DPS) [quest]; Ragefire Wand (7513, -2.47 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 203005000100000000-00000000000000000-2535111300000301050)

Set DPS (verified): 186.0. Weights run: 0.9s. Verify run: 0.8s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.692 ± 0.030, crit=0.267 ± 0.011 per rating point (14 rating = 1%, 3.743 per %), hit=0.543 ± 0.007 per rating point (10 rating = 1%, 5.434 per %), spell_haste=9.959 ± 0.579, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 32.8 spell_power points (6.57 DPS) | yes | Dreamweave Circlet (10041, -0.98 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.17 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.36 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.2 spell_power points (4.03 DPS) | yes | Mindburst Medallion (11196, -2.00 DPS) [quest]; Horizon Choker (13085, -2.09 DPS) [world_drop]; Scorn's Icy Choker (23169, -4.25 DPS, sim-verified) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -1.01 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.12 DPS) [crafted]; Rotgrip Mantle (17732, -2.69 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.2 spell_power points (3.63 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.58 DPS) [dungeon]; Runecloth Cloak (13860, -0.72 DPS) [crafted]; Big Voodoo Cloak (8216, -1.38 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 32.8 spell_power points (6.57 DPS) | yes | Robe of the Magi (1716, -1.34 DPS) [world_drop]; Runecloth Tunic (13857, -1.65 DPS) [crafted]; Dreamweave Vest (10021, -1.72 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.8 spell_power points (2.37 DPS) | yes | Bloodband Bracers (11469, -0.12 DPS) [quest]; Aristocratic Cuffs (12546, -0.29 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.57 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 27.1 spell_power points (5.43 DPS) | yes | Raider Handwraps (272098, -0.98 DPS) [vendor]; Dreamweave Gloves (10019, -1.27 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.58 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dawnspire Cord (12466, -0.20 DPS) [dungeon]; Deathmage Sash (10771, -0.55 DPS) [dungeon]; Satyrmane Sash (17755, -3.84 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.9 spell_power points (5.98 DPS) | yes | Red Mageweave Pants (10009, -1.52 DPS) [crafted]; Wizardweave Leggings (14132, -2.18 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.27 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.80 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.87 DPS) [vendor]; Gilded Sandals (254107, -1.35 DPS) [crafted]; Black Mageweave Boots (10026, -1.63 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (2.83 DPS) | yes | Band of the Unicorn (7553, -0.23 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.43 DPS) [rep]; Brainlash (6440, -0.75 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.8 spell_power points (2.77 DPS) | yes | Band of the Unicorn (7553, -0.17 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.37 DPS) [rep]; Brainlash (6440, -0.69 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellforce Rod (1664, -0.02 DPS) [world]; Spellshifter Rod (9527, -0.83 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.58 DPS) [world_drop]; Pyric Caduceus (11748, -2.95 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.65 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Noxious Shooter

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 203005000100000000-11302300000000000-2535111300000301050)

Set DPS (verified): 304.3. Weights run: 1.0s. Verify run: 0.7s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.830 ± 0.044, crit=0.450 ± 0.017 per rating point (14 rating = 1%, 6.299 per %), hit=0.933 ± 0.010 per rating point (10 rating = 1%, 9.330 per %), spell_haste=13.730 ± 0.867, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 48.0 spell_power points (9.03 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.09 DPS) [pvp]; Crimson Felt Hat (18727, -2.14 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (304.3 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.53 DPS) [quest]; Chains of the Lich (23125, -0.71 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.7 spell_power points (8.60 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.56 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.19 DPS) [crafted]; Darkspear Shoulderpads (272103, -6.19 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.0 spell_power points (6.01 DPS) | yes | Hide of the Wild (18510, -1.82 DPS) [crafted]; Spritecaster Cape (11623, -2.44 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.95 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.3 spell_power points (10.57 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.54 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -2.79 DPS) [pvp]; Robe of Everlasting Night (18385, -4.24 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.6 spell_power points (5.38 DPS) | yes | Sublime Wristguards (18497, -1.57 DPS) [dungeon]; Runecloth Cuffs (254123, -1.76 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.42 DPS) [pvp] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 32.9 spell_power points (6.19 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.6 spell_power points (9.88 DPS) | yes | Magician's Cord (272393, -2.57 DPS) [vendor]; Stormpike Cloth Girdle (19094, -4.94 DPS) [rep]; Belt of the Archmage (18405, -6.76 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.9 spell_power points (9.94 DPS) | yes | Marshal's Silk Leggings (231605, -0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -2.12 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -2.16 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.3 spell_power points (6.44 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.56 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (304.3 DPS) | yes | Songstone of Ironforge (12543, -1.38 DPS) [quest]; Maiden's Circle (13001, -1.38 DPS) [world_drop]; Naglering (11669, -10.82 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (304.3 DPS) | yes | Songstone of Ironforge (12543, -0.56 DPS) [quest]; Maiden's Circle (13001, -0.56 DPS) [world_drop]; Naglering (11669, -11.37 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (304.3 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (304.3 DPS) | yes | Weakness Analyzer (272438, -1.32 DPS) [vendor]; Serenity Field (272439, -2.82 DPS) [vendor]; Second Wind (11819, -6.42 DPS, sim-verified) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (304.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.80 DPS) [world]; Teebu's Blazing Longsword (1728, -15.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 414.5 spell_power points (77.89 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.97 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.57 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.21 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 28.5. Weights run: 0.9s. Verify run: 0.7s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.196 ± 0.005, crit=0.062 ± 0.002 per rating point (14 rating = 1%, 0.866 per %), hit=0.165 ± 0.001 per rating point (10 rating = 1%, 1.651 per %), spell_haste=0.706 ± 0.082, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.77 DPS) | yes | Shadow Goggles (4373, -1.07 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.8 spell_power points (0.87 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.34 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.35 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.51 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Battle Healer's Cloak (20427, -0.13 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.0 spell_power points (0.77 DPS) | yes | Green Woolen Vest (2582, -0.25 DPS) [crafted]; Bloody Apron (6226, -0.25 DPS) [dungeon]; Gray Woolen Robe (2585, -0.59 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.4 spell_power points (0.18 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Windsong Bangles (263336, -0.05 DPS) [quest]; Featherbead Bracers (15452, -0.05 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.90 DPS) | yes | Gnoll Casting Gloves (892, -0.13 DPS) [world]; Pristine Gloves (253913, -0.31 DPS) [crafted]; Apothecary Gloves (10919, -0.39 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.8 spell_power points (0.61 DPS) | yes | Novice Ardent's Sash (253887, -0.28 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.38 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.41 DPS) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.6 spell_power points (1.36 DPS) | yes | Filigreed Pristine Leggings (253937, -0.44 DPS) [crafted]; Silk-threaded Trousers (1929, -0.46 DPS) [dungeon]; Rumpled Kilt (274741, -0.71 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.8 spell_power points (1.00 DPS) | yes | Red Woolen Boots (4313, -0.49 DPS) [crafted]; Pristine Boots (253889, -0.54 DPS) [crafted]; Feather Padded Treads (285345, -0.64 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon]; Loop of Sacrifice (281673, -0.52 DPS) [quest]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.39 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.26 DPS) [quest]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.0 spell_power points (0.25 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.05 DPS) [world]; Lesser Staff of the Spire (1300, -0.10 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 176.8 spell_power points (22.70 DPS) | yes | Skycaller (12984, -1.20 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.41 DPS) [dungeon]; Sizzle Stick (8071, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 48.4. Weights run: 0.9s. Verify run: 0.7s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.297 ± 0.008, crit=0.101 ± 0.004 per rating point (14 rating = 1%, 1.411 per %), hit=0.183 ± 0.002 per rating point (10 rating = 1%, 1.826 per %), spell_haste=0.601 ± 0.136, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.65 DPS) | yes | Enchanter's Cowl (4322, -0.31 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon]; Silk Headband (7050, -0.60 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 spell_power points (1.32 DPS) | yes | Crystal Starfire Medallion (5003, -1.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.14 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.16 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (1.75 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.48 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (48.4 DPS) | yes | Death Speaker Robes (6682, -0.39 DPS) [dungeon]; Pristine Gown (253961, -0.57 DPS) [crafted]; Tree Bark Jacket (1486, -0.79 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.08 DPS) [quest]; Glowing Magical Bracelets (13106, -1.16 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.5 spell_power points (1.12 DPS) | yes | Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.59 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 spell_power points (1.78 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Warsong Sash (16975, -0.49 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.51 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.80 DPS) | yes | Abomination Skin Leggings (23173, -0.09 DPS) [dungeon]; Pristine Leggings (253987, -0.44 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.63 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.1 spell_power points (1.36 DPS) | yes | Acidic Walkers (9454, -0.26 DPS) [dungeon]; Boots of the Enchanter (4325, -0.61 DPS) [crafted]; Spidersilk Boots (4320, -1.60 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.05 DPS) | yes | Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon]; Sludge-Stained Band (286535, -0.60 DPS) [world]; Black Widow Band (6199, -0.74 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.90 DPS) | yes | Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Black Widow Band (6199, -0.59 DPS) [world]; Sludge-Stained Band (286535, -1.56 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Glimmering Staff (249392, -0.49 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.91 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.91 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.8 spell_power points (1.32 DPS) | yes | Orb of Souls (249395, -0.72 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.84 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.20 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 224.2 spell_power points (33.66 DPS) | yes | Starfaller (13063, -0.62 DPS) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 100000000000000000-00000000000000000-2535111300000301050)

Set DPS (verified): 79.9. Weights run: 1.0s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.516 ± 0.018, crit=0.166 ± 0.007 per rating point (14 rating = 1%, 2.328 per %), hit=0.308 ± 0.003 per rating point (10 rating = 1%, 3.080 per %), spell_haste=1.428 ± 0.283, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.10 DPS) | yes | Augural Shroud (2620, -0.71 DPS) [world]; Living Cowl (5608, -1.18 DPS) [world]; Enchanter's Cowl (4322, -1.45 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.1 spell_power points (1.49 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.73 DPS) [quest]; Triune Amulet (7722, -0.96 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.7 spell_power points (2.02 DPS) | yes | Green Silken Shoulders (7057, -0.00 DPS) [crafted]; Bloodmage Mantle (7684, -0.01 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.6 spell_power points (2.01 DPS) | yes | Guardian Cloak (5965, -0.75 DPS) [crafted]; Icy Cloak (4327, -0.98 DPS) [crafted]; Long Silken Cloak (4326, -1.41 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.1 spell_power points (3.70 DPS) | yes | Dreamweave Vest (10021, -0.36 DPS) [crafted]; Elemental Raiment (9434, -0.60 DPS) [world_drop]; Robe of Power (7054, -0.72 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.33 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.13 DPS) [quest]; Condor Bracers (15864, -0.29 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 spell_power points (2.96 DPS) | yes | Red Mageweave Gloves (10018, -0.58 DPS) [crafted]; Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Gilded Handwraps (254021, -1.25 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (79.9 DPS) | yes | Star Belt (4329, -0.26 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.32 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.2 spell_power points (2.98 DPS) | yes | Crimson Silk Pantaloons (7062, -0.88 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.04 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.21 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.54 DPS) | yes | Gilded Slippers (254001, -0.60 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.19 DPS) [dungeon]; Spidersilk Boots (4320, -2.20 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (1.93 DPS) | yes | Reedknot Ring (9622, -0.90 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.05 DPS) [vendor]; Black Widow Band (6199, -1.40 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.33 DPS) | yes | Sea Giant's Toe Ring (274746, -0.44 DPS) [vendor]; Reedknot Ring (9622, -0.69 DPS, sim-verified) [quest]; Black Widow Band (6199, -0.79 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Windweaver Staff (7757, -1.81 DPS) [dungeon]; Gut Ripper (2164, -4.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 270.1 spell_power points (39.84 DPS) | yes | Nether Force Wand (11263, -1.06 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.42 DPS) [quest]; Ragefire Wand (7513, -2.47 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 203005000100000000-00000000000000000-2535111300000301050)

Set DPS (verified): 172.1. Weights run: 0.9s. Verify run: 0.8s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.692 ± 0.030, crit=0.267 ± 0.011 per rating point (14 rating = 1%, 3.743 per %), hit=0.543 ± 0.007 per rating point (10 rating = 1%, 5.434 per %), spell_haste=9.959 ± 0.579, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 32.8 spell_power points (6.57 DPS) | yes | Dreamweave Circlet (10041, -0.98 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.17 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.36 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.2 spell_power points (4.03 DPS) | yes | Mindburst Medallion (11196, -2.00 DPS) [quest]; Horizon Choker (13085, -2.09 DPS) [world_drop]; Scorn's Icy Choker (23169, -3.97 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.5 spell_power points (5.09 DPS) | yes | Kentic Amice (11624, -0.49 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.50 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.62 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.2 spell_power points (3.65 DPS) | yes | Spritecaster Cape (11623, -0.02 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.60 DPS) [dungeon]; Runecloth Cloak (13860, -0.74 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 32.8 spell_power points (6.57 DPS) | yes | Robe of the Magi (1716, -1.34 DPS) [world_drop]; Runecloth Tunic (13857, -1.65 DPS) [crafted]; Dreamweave Vest (10021, -1.72 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.8 spell_power points (2.37 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.12 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 27.1 spell_power points (5.43 DPS) | yes | Raider Handwraps (272098, -0.98 DPS) [vendor]; Dreamweave Gloves (10019, -1.27 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.58 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (172.1 DPS) | yes | Dawnspire Cord (12466, -0.20 DPS) [dungeon]; Deathmage Sash (10771, -0.55 DPS) [dungeon]; Satyrmane Sash (17755, -3.31 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.9 spell_power points (5.98 DPS) | yes | Red Mageweave Pants (10009, -1.52 DPS) [crafted]; Wizardweave Leggings (14132, -2.18 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.51 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.80 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.87 DPS) [vendor]; Gilded Sandals (254107, -1.35 DPS) [crafted]; Black Mageweave Boots (10026, -1.63 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (2.83 DPS) | yes | Band of the Unicorn (7553, -0.23 DPS) [world_drop]; Advisor's Ring (19519, -0.43 DPS) [rep]; Brainlash (6440, -0.75 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.8 spell_power points (2.77 DPS) | yes | Band of the Unicorn (7553, -0.17 DPS) [world_drop]; Advisor's Ring (19519, -0.37 DPS) [rep]; Brainlash (6440, -0.69 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellforce Rod (1664, -0.02 DPS) [world]; Spellshifter Rod (9527, -0.83 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 262.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.50 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.15 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 203005000100000000-11302300000000000-2535111300000301050)

Set DPS (verified): 285.5. Weights run: 1.0s. Verify run: 0.7s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.830 ± 0.044, crit=0.450 ± 0.017 per rating point (14 rating = 1%, 6.299 per %), hit=0.933 ± 0.010 per rating point (10 rating = 1%, 9.330 per %), spell_haste=13.730 ± 0.867, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 48.0 spell_power points (9.03 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.09 DPS) [pvp]; Crimson Felt Hat (18727, -2.14 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (285.5 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.53 DPS) [quest]; Chains of the Lich (23125, -0.71 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.7 spell_power points (8.60 DPS) | yes | Warlord's Silk Amice (231594, -1.56 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.13 DPS) [vendor]; Mantle of the Timbermaw (19050, -2.19 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.0 spell_power points (6.01 DPS) | yes | Crystalline Threaded Cape (20697, -1.63 DPS) [world]; Hide of the Wild (18510, -1.82 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.35 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.3 spell_power points (10.57 DPS) | yes | Warlord's Silk Raiment (231596, -0.54 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.79 DPS) [pvp]; Robe of Everlasting Night (18385, -3.47 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.6 spell_power points (5.38 DPS) | yes | Sublime Wristguards (18497, -1.57 DPS) [dungeon]; Runecloth Cuffs (254123, -1.76 DPS) [crafted]; General's Silk Cuffs (16538, -2.42 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 32.9 spell_power points (6.19 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.6 spell_power points (9.88 DPS) | yes | Belt of the Archmage (18405, -2.44 DPS) [crafted]; Magician's Cord (272393, -2.57 DPS) [vendor]; Frostwolf Cloth Belt (19090, -4.94 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.9 spell_power points (9.94 DPS) | yes | General's Silk Trousers (231595, -0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.72 DPS) [rep]; Sorcerer's Leggings (226933, -2.12 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.3 spell_power points (6.44 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.56 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (285.5 DPS) | yes | Eye of Orgrimmar (12545, -1.38 DPS) [quest]; Maiden's Circle (13001, -1.38 DPS) [world_drop]; Naglering (11669, -7.96 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (285.5 DPS) | yes | Eye of Orgrimmar (12545, -0.56 DPS) [quest]; Maiden's Circle (13001, -0.56 DPS) [world_drop]; Naglering (11669, -10.05 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (285.5 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (285.5 DPS) | yes | Weakness Analyzer (272438, -1.32 DPS) [vendor]; Serenity Field (272439, -2.82 DPS) [vendor]; Burst of Knowledge (11832, -3.19 DPS) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (285.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.42 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -13.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 414.5 spell_power points (77.89 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.97 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.57 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.21 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

