# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 29.2. Weights run: 1.0s. Verify run: 0.7s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.216 ± 0.006, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.995 per %), hit=0.199 ± 0.001 per rating point (10 rating = 1%, 1.991 per %), spell_haste=0.592 ± 0.082, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

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

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 49.5. Weights run: 0.9s. Verify run: 0.7s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.359 ± 0.009, crit=0.110 ± 0.004 per rating point (14 rating = 1%, 1.543 per %), hit=0.199 ± 0.002 per rating point (10 rating = 1%, 1.989 per %), spell_haste=not significant (0.532 ± 0.158), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.44 DPS) | yes | Enchanter's Cowl (4322, -0.18 DPS) [crafted]; Silk Headband (7050, -0.26 DPS) [crafted]; Embalmed Shroud (7691, -0.39 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.2 spell_power points (1.20 DPS) | yes | Crystal Starfire Medallion (5003, -1.01 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.01 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.27 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.2 spell_power points (1.60 DPS) | yes | Death Speaker Mantle (6685, -0.30 DPS) [dungeon]; Fairywing Mantle (9536, -0.39 DPS) [quest]; Invoker's Mantle (215365, -0.45 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.65 DPS) | yes | Repairman's Cape (9605, -0.07 DPS) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Prelacy Cape (7004, -0.13 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.7 spell_power points (1.79 DPS) | yes | Death Speaker Robes (6682, -0.36 DPS) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted]; Tree Bark Jacket (1486, -0.95 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.18 DPS) | yes | Nightsky Wristbands (6407, -0.89 DPS) [world_drop]; Stonecloth Bindings (14416, -0.94 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.32 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.0 spell_power points (1.04 DPS) | yes | Serpent Gloves (5970, -0.12 DPS) [dungeon]; Truefaith Gloves (7049, -0.25 DPS) [crafted]; Shilly Mitts (9609, -0.49 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.1 spell_power points (1.58 DPS) | yes | Belt of Arugal (6392, -0.26 DPS) [dungeon]; Invoker's Cord (215366, -0.43 DPS) [crafted]; Crimson Silk Belt (7055, -0.47 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.57 DPS) | yes | Abomination Skin Leggings (23173, -0.02 DPS) [dungeon]; Pristine Leggings (253987, -0.32 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.50 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.5 spell_power points (1.24 DPS) | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Nimbus Boots (6998, -0.46 DPS) [quest]; Spidersilk Boots (4320, -1.64 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.92 DPS) | yes | Minor Channeling Ring (1449, -0.17 DPS) [quest]; Electrocutioner Lagnut (9447, -0.52 DPS) [dungeon]; Sludge-Stained Band (286535, -0.52 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.78 DPS) | yes | Electrocutioner Lagnut (9447, -0.39 DPS) [dungeon]; Sludge-Stained Band (286535, -0.39 DPS) [world]; Minor Channeling Ring (1449, -0.91 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.18 DPS) | yes | Glimmering Staff (249392, -0.57 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.71 DPS) [world_drop]; Channeler's Staff (4437, -0.80 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.2 spell_power points (1.20 DPS) | yes | Eye of Paleth (2943, -0.67 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.67 DPS) [world]; Dwarven Tome (279898, -1.03 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 257.0 spell_power points (33.60 DPS) | yes | Starfaller (13063, -0.55 DPS) [world_drop]; Greater Mystic Wand (217287, -3.95 DPS) [crafted]; Gravestone Scepter (7001, -4.60 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 100000000000000000-00000000000000000-2535111300000301050)

Set DPS (verified): 81.1. Weights run: 1.0s. Verify run: 0.6s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.569 ± 0.019, crit=0.180 ± 0.008 per rating point (14 rating = 1%, 2.526 per %), hit=0.331 ± 0.003 per rating point (10 rating = 1%, 3.306 per %), spell_haste=not significant (1.187 ± 0.305), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.75 DPS) | yes | Augural Shroud (2620, -0.56 DPS) [world]; Living Cowl (5608, -1.05 DPS) [world]; Enchanter's Cowl (4322, -1.22 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (1.36 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.62 DPS) [quest]; Triune Amulet (7722, -0.84 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.84 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.4 spell_power points (1.88 DPS) | yes | Bloodmage Mantle (7684, -0.04 DPS) [dungeon]; Berylline Pads (4197, -0.22 DPS) [quest]; Green Silken Shoulders (7057, -0.95 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.1 spell_power points (1.85 DPS) | yes | Guardian Cloak (5965, -0.69 DPS) [crafted]; Icy Cloak (4327, -0.93 DPS) [crafted]; Long Silken Cloak (4326, -1.65 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.4 spell_power points (3.32 DPS) | yes | Dreamweave Vest (10021, -0.30 DPS) [crafted]; Elemental Raiment (9434, -0.58 DPS) [world_drop]; Robe of Power (7054, -0.60 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.18 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.26 DPS) [quest]; Windchaser Cuffs (14429, -0.51 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (2.65 DPS) | yes | Red Mageweave Gloves (10018, -0.65 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.69 DPS) [crafted]; Gilded Handwraps (254021, -1.08 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.3 spell_power points (2.13 DPS) | yes | Deathmage Sash (10771, -0.10 DPS) [dungeon]; Star Belt (4329, -0.43 DPS) [crafted]; Gilded Cord (254037, -0.49 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.8 spell_power points (2.73 DPS) | yes | Crimson Silk Pantaloons (7062, -0.71 DPS) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.16 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.14 DPS) | yes | Gilded Slippers (254001, -1.03 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.89 DPS) [dungeon]; Spidersilk Boots (4320, -1.93 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (1.76 DPS) | yes | Ring of Forlorn Spirits (2043, -0.71 DPS) [quest]; Reedknot Ring (9622, -0.84 DPS) [quest]; Minor Channeling Ring (1449, -0.95 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.18 DPS) | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Minor Channeling Ring (1449, -0.37 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.44 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (81.1 DPS) | yes | Scorn's Focal Dagger (23168, -1.44 DPS) [dungeon]; Windweaver Staff (7757, -1.50 DPS) [dungeon]; Gut Ripper (2164, -4.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 304.4 spell_power points (39.83 DPS) | yes | Nether Force Wand (11263, -1.36 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.41 DPS) [quest]; Ragefire Wand (7513, -2.46 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 203005000100000000-00000000000000000-2535111300000301050)

Set DPS (verified): 172.7. Weights run: 0.9s. Verify run: 0.7s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.835 ± 0.034, crit=0.269 ± 0.011 per rating point (14 rating = 1%, 3.769 per %), hit=0.569 ± 0.007 per rating point (10 rating = 1%, 5.688 per %), spell_haste=11.366 ± 0.706, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.7 spell_power points (6.51 DPS) | yes | Dreamweave Circlet (10041, -1.16 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.44 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.59 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.0 spell_power points (3.83 DPS) | yes | Horizon Choker (13085, -1.70 DPS) [world_drop]; Mindburst Medallion (11196, -1.82 DPS) [quest]; Scorn's Icy Choker (23169, -4.08 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.0 spell_power points (5.11 DPS) | yes | Kentic Amice (11624, -0.58 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.55 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.59 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.0 spell_power points (3.46 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.45 DPS) [dungeon]; Runecloth Cloak (13860, -0.61 DPS) [crafted]; Big Voodoo Cloak (8216, -1.18 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.7 spell_power points (6.51 DPS) | yes | Robe of the Magi (1716, -1.58 DPS) [world_drop]; Runecloth Tunic (13857, -1.73 DPS) [crafted]; Dreamweave Vest (10021, -1.86 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.8 spell_power points (2.34 DPS) | yes | Bloodband Bracers (11469, -0.06 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.51 DPS) [quest]; Aristocratic Cuffs (12546, -2.69 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 29.4 spell_power points (5.36 DPS) | yes | Raider Handwraps (272098, -0.73 DPS) [vendor]; Dreamweave Gloves (10019, -1.47 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.62 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.4 spell_power points (4.07 DPS) | yes | Dawnspire Cord (12466, -0.09 DPS) [dungeon]; Ban'thok Sash (11662, -0.11 DPS) [dungeon]; Deathmage Sash (10771, -0.51 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.4 spell_power points (5.71 DPS) | yes | Red Mageweave Pants (10009, -1.34 DPS) [crafted]; Wizardweave Leggings (14132, -2.25 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.46 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.37 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.51 DPS) [vendor]; Gilded Sandals (254107, -1.00 DPS) [crafted]; Southsea Mojo Boots (20641, -1.24 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.0 spell_power points (2.74 DPS) | yes | Band of the Unicorn (7553, -0.37 DPS) [world_drop]; Brainlash (6440, -0.45 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.55 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.8 spell_power points (2.71 DPS) | yes | Band of the Unicorn (7553, -0.34 DPS) [world_drop]; Brainlash (6440, -0.42 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.52 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (172.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (172.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (172.7 DPS) | yes | Spellforce Rod (1664, -0.77 DPS) [world]; Spellshifter Rod (9527, -0.91 DPS) [quest]; Blade of Eternal Darkness (17780, -3.15 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 288.1 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.59 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 203005000100000000-11302300000000000-2535111300000301050)

Set DPS (verified): 284.2. Weights run: 1.0s. Verify run: 0.8s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.661 ± 0.034, crit=0.385 ± 0.014 per rating point (14 rating = 1%, 5.384 per %), hit=0.791 ± 0.008 per rating point (10 rating = 1%, 7.912 per %), spell_haste=11.080 ± 0.614, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.9 spell_power points (8.26 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.89 DPS) [pvp]; Crimson Felt Hat (18727, -1.47 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.31 DPS) [quest]; Chains of the Lich (23125, -0.31 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 42.3 spell_power points (8.14 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.42 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.18 DPS) [crafted]; Burial Shawl (18681, -2.26 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.2 spell_power points (5.62 DPS) | yes | Crystalline Threaded Cape (20697, -1.26 DPS) [world]; Hide of the Wild (18510, -1.65 DPS) [crafted]; Amplifying Cloak (18350, -2.15 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 53.3 spell_power points (10.26 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.71 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.02 DPS) [pvp]; Robe of Everlasting Night (18385, -3.41 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 27.3 spell_power points (5.25 DPS) | yes | Sublime Wristguards (18497, -1.67 DPS) [dungeon]; Runecloth Cuffs (254123, -1.86 DPS) [crafted]; Sorcerer's Bindings (226929, -2.63 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (284.2 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Sandworm Skin Gloves (20716, -3.80 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 47.5 spell_power points (9.13 DPS) | yes | Magician's Cord (272393, -2.42 DPS) [vendor]; Stormpike Cloth Girdle (19094, -4.40 DPS) [rep]; Belt of the Archmage (18405, -4.98 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.4 spell_power points (9.50 DPS) | yes | Marshal's Silk Leggings (231605, -0.15 DPS) [pvp]; Skyshroud Leggings (13170, -1.94 DPS) [dungeon]; Knight-Captain's Silk Legguards (227109, -2.26 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 31.6 spell_power points (6.07 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.58 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.28 DPS) [quest]; Maiden's Circle (13001, -1.28 DPS) [world_drop]; Naglering (11669, -11.02 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.58 DPS) [quest]; Maiden's Circle (13001, -0.58 DPS) [world_drop]; Naglering (11669, -10.67 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+14.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.35 DPS) [vendor]; Serenity Field (272439, -2.89 DPS) [vendor]; Burst of Knowledge (11832, -3.27 DPS) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.30 DPS) [world]; Teebu's Blazing Longsword (1728, -14.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 405.0 spell_power points (77.91 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.95 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.66 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.48 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 27.0. Weights run: 1.0s. Verify run: 0.7s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.216 ± 0.006, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.995 per %), hit=0.199 ± 0.001 per rating point (10 rating = 1%, 1.991 per %), spell_haste=0.592 ± 0.082, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

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

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 47.0. Weights run: 0.9s. Verify run: 0.7s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.359 ± 0.009, crit=0.110 ± 0.004 per rating point (14 rating = 1%, 1.543 per %), hit=0.199 ± 0.002 per rating point (10 rating = 1%, 1.989 per %), spell_haste=not significant (0.532 ± 0.158), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.44 DPS) | yes | Enchanter's Cowl (4322, -0.18 DPS) [crafted]; Silk Headband (7050, -0.26 DPS) [crafted]; Embalmed Shroud (7691, -0.39 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.2 spell_power points (1.20 DPS) | yes | Crystal Starfire Medallion (5003, -1.01 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.01 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.14 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.2 spell_power points (1.60 DPS) | yes | Fairywing Mantle (9536, -0.39 DPS) [quest]; Invoker's Mantle (215365, -0.45 DPS) [crafted]; Death Speaker Mantle (6685, -0.50 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.65 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Battle Healer's Cloak (19529, -0.13 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.7 spell_power points (1.79 DPS) | yes | Death Speaker Robes (6682, -0.36 DPS) [dungeon]; Pristine Gown (253961, -0.54 DPS) [crafted]; Tree Bark Jacket (1486, -0.84 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.18 DPS) | yes | Nightsky Wristbands (6407, -0.89 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.89 DPS) [quest]; Glowing Magical Bracelets (13106, -1.13 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.8 spell_power points (1.02 DPS) | yes | Truefaith Gloves (7049, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Serpent Gloves (5970, -0.42 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.1 spell_power points (1.58 DPS) | yes | Belt of Arugal (6392, -0.26 DPS) [dungeon]; Warsong Sash (16975, -0.41 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.43 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.57 DPS) | yes | Abomination Skin Leggings (23173, -0.02 DPS) [dungeon]; Pristine Leggings (253987, -0.32 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.50 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.5 spell_power points (1.24 DPS) | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Boots of the Enchanter (4325, -0.59 DPS) [crafted]; Spidersilk Boots (4320, -1.75 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.92 DPS) | yes | Electrocutioner Lagnut (9447, -0.52 DPS) [dungeon]; Sludge-Stained Band (286535, -0.52 DPS) [world]; Black Widow Band (6199, -0.59 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.78 DPS) | yes | Electrocutioner Lagnut (9447, -0.39 DPS) [dungeon]; Black Widow Band (6199, -0.46 DPS) [world]; Sludge-Stained Band (286535, -1.37 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.18 DPS) | yes | Glimmering Staff (249392, -0.52 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.71 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.71 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.2 spell_power points (1.20 DPS) | yes | Orb of Souls (249395, -0.67 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.75 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.15 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 257.0 spell_power points (33.60 DPS) | yes | Starfaller (13063, -0.55 DPS) [world_drop]; Greater Mystic Wand (217287, -3.95 DPS) [crafted]; Gravestone Scepter (7001, -4.60 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 100000000000000000-00000000000000000-2535111300000301050)

Set DPS (verified): 77.8. Weights run: 1.0s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.569 ± 0.019, crit=0.180 ± 0.008 per rating point (14 rating = 1%, 2.526 per %), hit=0.331 ± 0.003 per rating point (10 rating = 1%, 3.306 per %), spell_haste=not significant (1.187 ± 0.305), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.75 DPS) | yes | Augural Shroud (2620, -0.56 DPS) [world]; Living Cowl (5608, -1.05 DPS) [world]; Enchanter's Cowl (4322, -1.22 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (1.36 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.76 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.84 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.84 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.4 spell_power points (1.88 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Bloodmage Mantle (7684, -0.04 DPS) [dungeon]; Berylline Pads (4197, -0.22 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.1 spell_power points (1.85 DPS) | yes | Guardian Cloak (5965, -0.69 DPS) [crafted]; Icy Cloak (4327, -0.93 DPS) [crafted]; Long Silken Cloak (4326, -1.77 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.4 spell_power points (3.32 DPS) | yes | Dreamweave Vest (10021, -0.30 DPS) [crafted]; Elemental Raiment (9434, -0.58 DPS) [world_drop]; Robe of Power (7054, -0.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.18 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.06 DPS) [quest]; Condor Bracers (15864, -0.26 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (2.65 DPS) | yes | Black Mageweave Gloves (10003, -0.69 DPS) [crafted]; Red Mageweave Gloves (10018, -0.89 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.08 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.3 spell_power points (2.13 DPS) | yes | Deathmage Sash (10771, -0.10 DPS) [dungeon]; Star Belt (4329, -0.43 DPS) [crafted]; Gilded Cord (254037, -0.49 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.8 spell_power points (2.73 DPS) | yes | Crimson Silk Pantaloons (7062, -0.69 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.16 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.14 DPS) | yes | Gilded Slippers (254001, -1.07 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.89 DPS) [dungeon]; Spidersilk Boots (4320, -1.93 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (1.76 DPS) | yes | Reedknot Ring (9622, -0.84 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.97 DPS) [vendor]; Black Widow Band (6199, -1.23 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.18 DPS) | yes | Sea Giant's Toe Ring (274746, -0.39 DPS) [vendor]; Black Widow Band (6199, -0.66 DPS) [world]; Reedknot Ring (9622, -1.10 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (77.8 DPS) | yes | Scorn's Focal Dagger (23168, -1.44 DPS) [dungeon]; Windweaver Staff (7757, -1.50 DPS) [dungeon]; Gut Ripper (2164, -4.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 304.4 spell_power points (39.83 DPS) | yes | Nether Force Wand (11263, -1.63 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.41 DPS) [quest]; Ragefire Wand (7513, -2.46 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 203005000100000000-00000000000000000-2535111300000301050)

Set DPS (verified): 169.0. Weights run: 0.9s. Verify run: 0.7s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.835 ± 0.034, crit=0.269 ± 0.011 per rating point (14 rating = 1%, 3.769 per %), hit=0.569 ± 0.007 per rating point (10 rating = 1%, 5.688 per %), spell_haste=11.366 ± 0.706, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.7 spell_power points (6.51 DPS) | yes | Dreamweave Circlet (10041, -1.16 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.44 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.59 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.0 spell_power points (3.83 DPS) | yes | Horizon Choker (13085, -1.70 DPS) [world_drop]; Mindburst Medallion (11196, -1.82 DPS) [quest]; Scorn's Icy Choker (23169, -3.98 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.0 spell_power points (5.11 DPS) | yes | Kentic Amice (11624, -0.58 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.55 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.59 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.5 spell_power points (3.56 DPS) | yes | Spritecaster Cape (11623, -0.09 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.55 DPS) [dungeon]; Runecloth Cloak (13860, -0.70 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.7 spell_power points (6.51 DPS) | yes | Runecloth Tunic (13857, -1.73 DPS) [crafted]; Dreamweave Vest (10021, -1.86 DPS) [crafted]; Robe of the Magi (1716, -2.49 DPS, sim-verified) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.8 spell_power points (2.34 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Aristocratic Cuffs (12546, -3.68 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 29.4 spell_power points (5.36 DPS) | yes | Raider Handwraps (272098, -0.73 DPS) [vendor]; Dreamweave Gloves (10019, -1.47 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.62 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.4 spell_power points (4.07 DPS) | yes | Dawnspire Cord (12466, -0.09 DPS) [dungeon]; Ban'thok Sash (11662, -0.11 DPS) [dungeon]; Deathmage Sash (10771, -0.51 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.4 spell_power points (5.71 DPS) | yes | Red Mageweave Pants (10009, -1.34 DPS) [crafted]; Wizardweave Leggings (14132, -2.25 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -4.44 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.37 DPS) | yes | Gilded Sandals (254107, -1.00 DPS) [crafted]; Southsea Mojo Boots (20641, -1.24 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -2.81 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.0 spell_power points (2.74 DPS) | yes | Band of the Unicorn (7553, -0.37 DPS) [world_drop]; Brainlash (6440, -0.45 DPS) [dungeon]; Advisor's Ring (19519, -0.55 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.8 spell_power points (2.71 DPS) | yes | Brainlash (6440, -0.42 DPS) [dungeon]; Advisor's Ring (19519, -0.52 DPS) [rep]; Band of the Unicorn (7553, -2.94 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (169.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (169.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (169.0 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellforce Rod (1664, -0.77 DPS) [world]; Spellshifter Rod (9527, -0.91 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 288.1 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.59 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 203005000100000000-11302300000000000-2535111300000301050)

Set DPS (verified): 271.6. Weights run: 1.0s. Verify run: 0.7s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.661 ± 0.034, crit=0.385 ± 0.014 per rating point (14 rating = 1%, 5.384 per %), hit=0.791 ± 0.008 per rating point (10 rating = 1%, 7.912 per %), spell_haste=11.080 ± 0.614, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.9 spell_power points (8.26 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.89 DPS) [pvp]; Crimson Felt Hat (18727, -1.47 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (271.6 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.31 DPS) [quest]; Chains of the Lich (23125, -0.31 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 42.3 spell_power points (8.14 DPS) | yes | Warlord's Silk Amice (231594, -1.42 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.18 DPS) [crafted]; Burial Shawl (18681, -2.26 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.2 spell_power points (5.62 DPS) | yes | Crystalline Threaded Cape (20697, -1.26 DPS) [world]; Hide of the Wild (18510, -1.65 DPS) [crafted]; Amplifying Cloak (18350, -2.15 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 53.3 spell_power points (10.26 DPS) | yes | Warlord's Silk Raiment (231596, -0.71 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.02 DPS) [pvp]; Robe of Everlasting Night (18385, -3.41 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.3 spell_power points (5.25 DPS) | yes | Sublime Wristguards (18497, -1.67 DPS) [dungeon]; Runecloth Cuffs (254123, -1.86 DPS) [crafted]; Sorcerer's Bindings (226929, -2.63 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 30.3 spell_power points (5.83 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 47.5 spell_power points (9.13 DPS) | yes | Magician's Cord (272393, -2.42 DPS) [vendor]; Frostwolf Cloth Belt (19090, -4.40 DPS) [rep]; Belt of the Archmage (18405, -4.80 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.4 spell_power points (9.50 DPS) | yes | General's Silk Trousers (231595, -0.15 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.70 DPS) [rep]; Skyshroud Leggings (13170, -1.94 DPS) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 31.6 spell_power points (6.07 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.58 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (271.6 DPS) | yes | Eye of Orgrimmar (12545, -1.28 DPS) [quest]; Maiden's Circle (13001, -1.28 DPS) [world_drop]; Naglering (11669, -8.04 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (271.6 DPS) | yes | Eye of Orgrimmar (12545, -0.58 DPS) [quest]; Maiden's Circle (13001, -0.58 DPS) [world_drop]; Naglering (11669, -6.62 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (271.6 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (271.6 DPS) | yes | Weakness Analyzer (272438, -1.35 DPS) [vendor]; Serenity Field (272439, -2.89 DPS) [vendor]; Burst of Knowledge (11832, -3.27 DPS) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (271.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.49 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -14.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 405.0 spell_power points (77.91 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.95 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.66 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.48 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

