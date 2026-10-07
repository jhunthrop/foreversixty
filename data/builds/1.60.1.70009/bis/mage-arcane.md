# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 153002000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 28.2. Weights run: 0.8s. Verify run: 0.5s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.271 ± 0.006, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.079 per %), hit=0.269 ± 0.007 per rating point (10 rating = 1%, 2.693 per %), spell_haste=0.356 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.461 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -1.13 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (0.70 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.32 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.71 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.23 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.4 spell_power points (0.60 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -0.78 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.4 spell_power points (0.13 DPS) | yes | Windsong Bangles (263336, -0.03 DPS) [quest]; Repurposed Hair Band (281256, -0.08 DPS) [quest]; Bright Bracers (3647, -0.73 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.66 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.21 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.42 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.1 spell_power points (0.48 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.28 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.56 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.2 spell_power points (1.05 DPS) | yes | Filigreed Pristine Leggings (253937, -0.33 DPS) [crafted]; Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Rumpled Kilt (274741, -0.58 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (0.76 DPS) | yes | Red Woolen Boots (4313, -0.39 DPS) [crafted]; Pristine Boots (253889, -0.40 DPS) [crafted]; Feather Padded Treads (285345, -0.76 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (0.52 DPS) | yes | Sludge-Stained Band (286535, -0.24 DPS) [world]; Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.45 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.47 DPS) | yes | Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop]; Sludge-Stained Band (286535, -0.49 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.7 spell_power points (0.26 DPS) | yes | Lesser Staff of the Spire (1300, -0.10 DPS) [world]; Staff of Westfall (2042, -0.13 DPS) [quest]; Channeler's Staff (4437, -0.17 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 239.3 spell_power points (22.59 DPS) | yes | Skycaller (12984, -0.50 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.30 DPS) [dungeon]; Deepblaze (279896, -4.06 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 153005113100010000-00000000000000000-0000000000000000000)

Set DPS (verified): 98.0. Weights run: 1.0s. Verify run: 0.6s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.349 ± 0.013, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.897 per %), hit=0.252 ± 0.024 per rating point (10 rating = 1%, 2.519 per %), spell_haste=not significant (0.289 ± 0.188), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.915 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.91 DPS) | yes | Silk Headband (7050, -0.53 DPS) [crafted]; Embalmed Shroud (7691, -0.79 DPS) [dungeon]; Enchanter's Cowl (4322, -1.01 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.1 spell_power points (2.40 DPS) | yes | Crystal Starfire Medallion (5003, -2.04 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.04 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.59 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (3.21 DPS) | yes | Death Speaker Mantle (6685, -0.66 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.79 DPS) [quest]; Invoker's Mantle (215365, -0.90 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.32 DPS) | yes | Repairman's Cape (9605, -0.16 DPS) [quest]; Heavy Woolen Cloak (4311, -0.26 DPS) [crafted]; Prelacy Cape (7004, -0.26 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.5 spell_power points (3.58 DPS) | yes | Tree Bark Jacket (1486, -0.14 DPS) [dungeon]; Death Speaker Robes (6682, -0.71 DPS) [dungeon]; Pristine Gown (253961, -1.08 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.38 DPS) | yes | Nightsky Wristbands (6407, -1.83 DPS) [world_drop]; Stonecloth Bindings (14416, -1.92 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.71 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.8 spell_power points (2.07 DPS) | yes | Serpent Gloves (5970, -0.22 DPS) [dungeon]; Shilly Mitts (9609, -0.22 DPS) [quest]; Truefaith Gloves (7049, -0.47 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.0 spell_power points (3.19 DPS) | yes | Belt of Arugal (6392, -0.53 DPS) [dungeon]; Invoker's Cord (215366, -0.87 DPS) [crafted]; Crimson Silk Belt (7055, -0.95 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.17 DPS) | yes | Abomination Skin Leggings (23173, -0.67 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.68 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.03 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (2.50 DPS) | yes | Acidic Walkers (9454, -0.44 DPS) [dungeon]; Nimbus Boots (6998, -0.91 DPS) [quest]; Spidersilk Boots (4320, -2.61 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.85 DPS) | yes | Minor Channeling Ring (1449, -0.34 DPS) [quest]; Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon]; Sludge-Stained Band (286535, -1.06 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.59 DPS) | yes | Electrocutioner Lagnut (9447, -0.79 DPS) [dungeon]; Sludge-Stained Band (286535, -0.79 DPS) [world]; Minor Channeling Ring (1449, -2.04 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.38 DPS) | yes | Twisted Chanter's Staff (890, -1.46 DPS) [world_drop]; Channeler's Staff (4437, -1.64 DPS) [world]; Glimmering Staff (249392, -1.92 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.1 spell_power points (2.40 DPS) | yes | Dwarven Tome (279898, -0.70 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.35 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -1.35 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 128.6 spell_power points (34.00 DPS) | yes | Starfaller (13063, -0.64 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.68 DPS) [crafted]; Gravestone Scepter (7001, -5.00 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 153005113100011531-00000000000000000-0000000000000000000)

Set DPS (verified): 182.9. Weights run: 1.0s. Verify run: 0.6s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.583 ± 0.027, crit=0.174 ± 0.004 per rating point (14 rating = 1%, 2.442 per %), hit=0.399 ± 0.052 per rating point (10 rating = 1%, 3.993 per %), spell_haste=not significant (1.291 ± 0.388), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.918 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.38 DPS) | yes | Augural Shroud (2620, -1.27 DPS) [world]; Living Cowl (5608, -2.43 DPS) [world]; Enchanter's Cowl (4322, -2.78 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.5 spell_power points (3.19 DPS) | yes | Triune Amulet (7722, -1.95 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.95 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.03 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.43 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.53 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.2 spell_power points (4.33 DPS) | yes | Guardian Cloak (5965, -1.62 DPS) [crafted]; Icy Cloak (4327, -2.20 DPS) [crafted]; Long Silken Cloak (4326, -3.87 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (7.74 DPS) | yes | Dreamweave Vest (10021, -0.68 DPS) [crafted]; Elemental Raiment (9434, -1.37 DPS) [world_drop]; Robe of Power (7054, -1.37 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.73 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.61 DPS) [quest]; Windchaser Cuffs (14429, -1.14 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (6.17 DPS) | yes | Black Mageweave Gloves (10003, -1.62 DPS) [crafted]; Red Mageweave Gloves (10018, -1.90 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.51 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.3 spell_power points (4.96 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Star Belt (4329, -1.01 DPS) [crafted]; Gilded Cord (254037, -1.11 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.38 DPS) | yes | Abomination Skin Leggings (23173, -2.23 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.23 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.73 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.29 DPS) | yes | Gilded Slippers (254001, -1.83 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.35 DPS) [dungeon]; Spidersilk Boots (4320, -4.45 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.10 DPS) | yes | Ring of Forlorn Spirits (2043, -1.67 DPS) [quest]; Reedknot Ring (9622, -1.97 DPS) [quest]; Minor Channeling Ring (1449, -2.23 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.73 DPS) | yes | Reedknot Ring (9622, -0.61 DPS) [quest]; Minor Channeling Ring (1449, -0.86 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.24 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (182.9 DPS) | yes | Scorn's Focal Dagger (23168, -3.34 DPS) [dungeon]; Windweaver Staff (7757, -3.42 DPS) [dungeon]; Gut Ripper (2164, -10.30 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 133.2 spell_power points (40.44 DPS) | yes | Nether Force Wand (11263, -2.58 DPS) [quest]; Icefury Wand (7514, -2.72 DPS) [quest]; Ragefire Wand (7513, -2.77 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 153005113100011531-03202300000000000-0000000000000000000)

Set DPS (verified): 277.7. Weights run: 1.0s. Verify run: 0.7s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.714 ± 0.039, crit=0.276 ± 0.006 per rating point (14 rating = 1%, 3.864 per %), hit=0.678 ± 0.080 per rating point (10 rating = 1%, 6.780 per %), spell_haste=not significant (1.314 ± 0.598), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.920 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 33.3 spell_power points (10.32 DPS) | yes | Dreamweave Circlet (10041, -1.59 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.95 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.12 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.3 spell_power points (6.29 DPS) | yes | Mindburst Medallion (11196, -3.10 DPS) [quest]; Horizon Choker (13085, -3.19 DPS) [world_drop]; Scorn's Icy Choker (23169, -6.10 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.8 spell_power points (8.01 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.34 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.52 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.3 spell_power points (5.67 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.89 DPS) [dungeon]; Runecloth Cloak (13860, -1.11 DPS) [crafted]; Big Voodoo Cloak (8216, -2.13 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.3 spell_power points (10.32 DPS) | yes | Robe of the Magi (1716, -2.17 DPS) [world_drop]; Runecloth Tunic (13857, -2.61 DPS) [crafted]; Dreamweave Vest (10021, -2.74 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.0 spell_power points (3.72 DPS) | yes | Aristocratic Cuffs (12546, -0.40 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.93 DPS) [quest]; Bloodband Bracers (11469, -3.22 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.8 spell_power points (8.92 DPS) | yes | Dreamweave Gloves (10019, -2.45 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.90 DPS) [vendor]; Raider Handwraps (272098, -3.21 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.6 spell_power points (8.26 DPS) | yes | Dawnspire Cord (12466, -2.19 DPS) [dungeon]; Deathmage Sash (10771, -2.77 DPS) [dungeon]; Satyrmane Sash (17755, -5.91 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 30.1 spell_power points (9.34 DPS) | yes | Red Mageweave Pants (10009, -2.35 DPS) [crafted]; Wizardweave Leggings (14132, -3.45 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -5.47 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.44 DPS) | yes | Gilded Sandals (254107, -2.04 DPS) [crafted]; Black Mageweave Boots (10026, -2.48 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.00 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (4.43 DPS) | yes | Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.71 DPS) [rep]; Brainlash (6440, -1.11 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.0 spell_power points (4.34 DPS) | yes | Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.62 DPS) [rep]; Brainlash (6440, -1.02 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (277.7 DPS) | yes | Uther's Strength (11302, -2.16 DPS, sim-verified) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (277.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (277.7 DPS) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -8.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 169.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.27 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 484.6. Weights run: 1.1s. Verify run: 0.7s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.852 ± 0.052, crit=0.443 ± 0.010 per rating point (14 rating = 1%, 6.201 per %), hit=0.966 ± 0.105 per rating point (10 rating = 1%, 9.656 per %), spell_haste=3.490 ± 0.838, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.914 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 48.5 spell_power points (15.79 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.94 DPS) [pvp]; Crimson Felt Hat (18727, -3.80 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (484.6 DPS) | yes | Diana's Pearl Necklace (22403, -0.20 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.93 DPS) [quest]; Jewel of Kajaro (19601, -2.63 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.0 spell_power points (14.97 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.67 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.81 DPS) [crafted]; Darkspear Shoulderpads (272103, -9.19 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.5 spell_power points (10.57 DPS) | yes | Hide of the Wild (18510, -3.24 DPS) [crafted]; Spritecaster Cape (11623, -4.35 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.59 DPS, sim-verified) [world] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.4 spell_power points (18.36 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.89 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.80 DPS) [pvp]; Robe of Everlasting Night (18385, -5.36 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.8 spell_power points (9.38 DPS) | yes | Sublime Wristguards (18497, -2.70 DPS) [dungeon]; Runecloth Cuffs (254123, -3.03 DPS) [crafted]; Marshal's Silk Bracers (16438, -4.11 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 33.6 spell_power points (10.93 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.4 spell_power points (17.38 DPS) | yes | Magician's Cord (272393, -4.54 DPS) [vendor]; Ban'thok Sash (11662, -7.28 DPS) [dungeon]; Belt of the Archmage (18405, -8.24 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.9 spell_power points (17.22 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -3.46 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -3.66 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.6 spell_power points (11.27 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.98 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (484.6 DPS) | yes | Rune Band of Wizardry (22339, -1.86 DPS) [dungeon]; Maiden's Circle (13001, -2.41 DPS) [world_drop]; Naglering (11669, -12.81 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (484.6 DPS) | yes | Rune Band of Wizardry (22339, -0.43 DPS) [dungeon]; Maiden's Circle (13001, -0.98 DPS) [world_drop]; Naglering (11669, -11.90 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (484.6 DPS) | yes | Weakness Analyzer (272438, -2.28 DPS) [vendor]; Serenity Field (272439, -4.88 DPS) [vendor]; Blackhand's Breadth (13965, -5.40 DPS) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (484.6 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -3.04 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (484.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Crackling Staff (19102, -0.61 DPS) [rep]; Teebu's Blazing Longsword (1728, -28.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 241.9 spell_power points (78.71 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.28 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.39 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.26 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 698.9. Weights run: 1.2s. Verify run: 0.8s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.147 ± 0.007, crit=0.485 ± 0.010 per rating point (14 rating = 1%, 6.791 per %), hit=1.003 ± 0.031 per rating point (10 rating = 1%, 10.030 per %), spell_haste=4.887 ± 0.143, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.904 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.5 spell_power points (15.51 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.51 DPS) [pvp] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (698.9 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -0.88 DPS) [dungeon]; Jewel of Kajaro (19601, -5.72 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.0 spell_power points (17.73 DPS) | yes | Field Marshal's Silk Spaulders (231602, -4.33 DPS) [pvp]; Argent Shoulders (19059, -5.42 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.61 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.2 spell_power points (13.40 DPS) | yes | Amplifying Cloak (18350, -4.54 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.78 DPS, sim-verified) [world]; Hide of the Wild (18510, -5.78 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.6 spell_power points (23.92 DPS) | yes | Field Marshal's Silk Vestments (231603, -3.09 DPS) [pvp]; Robe of Everlasting Night (18385, -7.03 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Tunic (227108, -9.00 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.2 spell_power points (11.42 DPS) | yes | Sublime Wristguards (18497, -4.78 DPS) [dungeon]; Runecloth Cuffs (254123, -5.27 DPS) [crafted]; Arcane Runed Bracers (4744, -6.98 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (13.66 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.3 spell_power points (18.85 DPS) | yes | Ban'thok Sash (11662, -7.20 DPS) [dungeon]; Magician's Cord (272393, -7.75 DPS) [vendor]; Belt of the Archmage (18405, -8.14 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 47.1 spell_power points (23.18 DPS) | yes | Marshal's Silk Leggings (231605, -3.60 DPS) [pvp]; Skyshroud Leggings (13170, -5.85 DPS) [dungeon]; Sorcerer's Leggings (226933, -7.72 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.82 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.32 DPS) [quest]; Sorcerer's Sandals (226931, -0.32 DPS) [vendor] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (698.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.97 DPS) [quest]; Blessed Band of Light (272407, -2.96 DPS) [vendor]; Naglering (11669, -21.52 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (698.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.78 DPS) [quest]; Blessed Band of Light (272407, -1.77 DPS) [vendor]; Naglering (11669, -12.91 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (698.9 DPS) | yes | Weakness Analyzer (272438, -3.45 DPS) [vendor]; Serenity Field (272439, -7.39 DPS) [vendor]; Blackhand's Breadth (13965, -7.59 DPS) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (698.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -4.01 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (698.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -2.09 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -32.29 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 161.8 spell_power points (79.72 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.45 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.38 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.27 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 153002000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 26.5. Weights run: 0.8s. Verify run: 0.6s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.271 ± 0.006, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.079 per %), hit=0.269 ± 0.007 per rating point (10 rating = 1%, 2.693 per %), spell_haste=0.356 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.461 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -0.69 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (0.70 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.15 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.4 spell_power points (0.60 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -0.65 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.6 spell_power points (0.15 DPS) | yes | Owlbeard Bracers (16981, -0.01 DPS) [quest]; Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.66 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.21 DPS) [crafted]; Apothecary Gloves (10919, -0.28 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.1 spell_power points (0.48 DPS) | yes | Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Keller's Girdle (2911, -0.28 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.40 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (26.5 DPS) | yes | Silk-threaded Trousers (1929, -0.06 DPS) [dungeon]; Rumpled Kilt (274741, -0.25 DPS) [vendor]; Abomination Skin Leggings (23173, -0.43 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (0.76 DPS) | yes | Red Woolen Boots (4313, -0.39 DPS) [crafted]; Feather Padded Treads (285345, -0.39 DPS, sim-verified) [world]; Pristine Boots (253889, -0.40 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.47 DPS) | yes | Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Loop of Sacrifice (281673, -0.34 DPS) [quest]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Loop of Sacrifice (281673, -0.16 DPS) [quest]; Volcanic Rock Ring (12053, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.7 spell_power points (0.26 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.05 DPS) [world]; Lesser Staff of the Spire (1300, -0.10 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 239.3 spell_power points (22.59 DPS) | yes | Skycaller (12984, -0.33 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.30 DPS) [dungeon]; Sizzle Stick (8071, -4.47 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 153005113100010000-00000000000000000-0000000000000000000)

Set DPS (verified): 94.8. Weights run: 1.0s. Verify run: 0.6s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.349 ± 0.013, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.897 per %), hit=0.252 ± 0.024 per rating point (10 rating = 1%, 2.519 per %), spell_haste=not significant (0.289 ± 0.188), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.915 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.91 DPS) | yes | Silk Headband (7050, -0.53 DPS) [crafted]; Embalmed Shroud (7691, -0.79 DPS) [dungeon]; Enchanter's Cowl (4322, -1.34 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.1 spell_power points (2.40 DPS) | yes | Crystal Starfire Medallion (5003, -2.04 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.04 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.67 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (3.21 DPS) | yes | Death Speaker Mantle (6685, -0.72 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.79 DPS) [quest]; Invoker's Mantle (215365, -0.90 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.32 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.26 DPS) [crafted]; Battle Healer's Cloak (19529, -0.26 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.5 spell_power points (3.58 DPS) | yes | Tree Bark Jacket (1486, -0.14 DPS) [dungeon]; Death Speaker Robes (6682, -0.71 DPS) [dungeon]; Pristine Gown (253961, -1.08 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.38 DPS) | yes | Nightsky Wristbands (6407, -1.83 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.83 DPS) [quest]; Glowing Magical Bracelets (13106, -2.74 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.7 spell_power points (2.05 DPS) | yes | Serpent Gloves (5970, -0.20 DPS) [dungeon]; Truefaith Gloves (7049, -0.45 DPS) [crafted]; Gnoll Casting Gloves (892, -0.46 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (3.19 DPS) | yes | Warsong Sash (16975, -0.28 DPS) [quest]; Belt of Arugal (6392, -0.53 DPS) [dungeon]; Invoker's Cord (215366, -0.87 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.17 DPS) | yes | Pristine Leggings (253987, -0.68 DPS) [crafted]; Abomination Skin Leggings (23173, -0.87 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -1.03 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (2.50 DPS) | yes | Acidic Walkers (9454, -0.44 DPS) [dungeon]; Boots of the Enchanter (4325, -1.17 DPS) [crafted]; Spidersilk Boots (4320, -2.77 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.85 DPS) | yes | Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon]; Sludge-Stained Band (286535, -1.06 DPS) [world]; Black Widow Band (6199, -1.21 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.59 DPS) | yes | Electrocutioner Lagnut (9447, -0.79 DPS) [dungeon]; Black Widow Band (6199, -0.94 DPS) [world]; Sludge-Stained Band (286535, -2.96 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.38 DPS) | yes | Twisted Chanter's Staff (890, -1.46 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.46 DPS) [quest]; Glimmering Staff (249392, -2.20 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.1 spell_power points (2.40 DPS) | yes | Orb of Souls (249395, -1.35 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.51 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.78 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 128.6 spell_power points (34.00 DPS) | yes | Starfaller (13063, -0.88 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.68 DPS) [crafted]; Gravestone Scepter (7001, -5.00 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 153005113100011531-00000000000000000-0000000000000000000)

Set DPS (verified): 178.0. Weights run: 1.0s. Verify run: 0.6s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.583 ± 0.027, crit=0.174 ± 0.004 per rating point (14 rating = 1%, 2.442 per %), hit=0.399 ± 0.052 per rating point (10 rating = 1%, 3.993 per %), spell_haste=not significant (1.291 ± 0.388), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.918 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.38 DPS) | yes | Living Cowl (5608, -2.43 DPS) [world]; Augural Shroud (2620, -2.70 DPS, sim-verified) [world]; Enchanter's Cowl (4322, -2.78 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.5 spell_power points (3.19 DPS) | yes | Triune Amulet (7722, -1.95 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.95 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.32 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.43 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.53 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.2 spell_power points (4.33 DPS) | yes | Guardian Cloak (5965, -1.62 DPS) [crafted]; Icy Cloak (4327, -2.20 DPS) [crafted]; Long Silken Cloak (4326, -2.29 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (7.74 DPS) | yes | Dreamweave Vest (10021, -0.68 DPS) [crafted]; Elemental Raiment (9434, -1.37 DPS) [world_drop]; Robe of Power (7054, -1.37 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.73 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.61 DPS) [quest]; Radiant Silver Bracers (4545, -1.84 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (6.17 DPS) | yes | Black Mageweave Gloves (10003, -1.62 DPS) [crafted]; Red Mageweave Gloves (10018, -2.43 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.51 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.3 spell_power points (4.96 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Star Belt (4329, -1.01 DPS) [crafted]; Gilded Cord (254037, -1.11 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.38 DPS) | yes | Abomination Skin Leggings (23173, -2.23 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.51 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.73 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.29 DPS) | yes | Gilded Slippers (254001, -2.76 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.35 DPS) [dungeon]; Spidersilk Boots (4320, -4.45 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.10 DPS) | yes | Reedknot Ring (9622, -1.97 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.28 DPS) [vendor]; Black Widow Band (6199, -2.86 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.73 DPS) | yes | Reedknot Ring (9622, -0.61 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.91 DPS) [vendor]; Black Widow Band (6199, -1.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (178.0 DPS) | yes | Scorn's Focal Dagger (23168, -3.34 DPS) [dungeon]; Windweaver Staff (7757, -3.42 DPS) [dungeon]; Gut Ripper (2164, -10.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 133.2 spell_power points (40.44 DPS) | yes | Nether Force Wand (11263, -2.58 DPS) [quest]; Icefury Wand (7514, -2.72 DPS) [quest]; Ragefire Wand (7513, -2.77 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 153005113100011531-03202300000000000-0000000000000000000)

Set DPS (verified): 270.8. Weights run: 1.0s. Verify run: 0.7s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.714 ± 0.039, crit=0.276 ± 0.006 per rating point (14 rating = 1%, 3.864 per %), hit=0.678 ± 0.080 per rating point (10 rating = 1%, 6.780 per %), spell_haste=not significant (1.314 ± 0.598), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.920 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 33.3 spell_power points (10.32 DPS) | yes | Dreamweave Circlet (10041, -1.59 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.95 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.12 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.3 spell_power points (6.29 DPS) | yes | Mindburst Medallion (11196, -3.10 DPS) [quest]; Horizon Choker (13085, -3.19 DPS) [world_drop]; Scorn's Icy Choker (23169, -6.03 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.8 spell_power points (8.01 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.34 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.52 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.4 spell_power points (5.71 DPS) | yes | Spritecaster Cape (11623, -0.04 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.93 DPS) [dungeon]; Runecloth Cloak (13860, -1.15 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.3 spell_power points (10.32 DPS) | yes | Robe of the Magi (1716, -2.17 DPS) [world_drop]; Runecloth Tunic (13857, -2.61 DPS) [crafted]; Dreamweave Vest (10021, -2.74 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.0 spell_power points (3.72 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -3.42 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.8 spell_power points (8.92 DPS) | yes | Raider Handwraps (272098, -1.88 DPS) [vendor]; Dreamweave Gloves (10019, -2.45 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.90 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.6 spell_power points (8.26 DPS) | yes | Dawnspire Cord (12466, -2.19 DPS) [dungeon]; Deathmage Sash (10771, -2.77 DPS) [dungeon]; Satyrmane Sash (17755, -3.77 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 30.1 spell_power points (9.34 DPS) | yes | Red Mageweave Pants (10009, -2.35 DPS) [crafted]; Wizardweave Leggings (14132, -3.45 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -5.59 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.44 DPS) | yes | Gilded Sandals (254107, -2.04 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.34 DPS, sim-verified) [vendor]; Black Mageweave Boots (10026, -2.48 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (4.43 DPS) | yes | Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Advisor's Ring (19519, -0.71 DPS) [rep]; Brainlash (6440, -1.11 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.0 spell_power points (4.34 DPS) | yes | Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Advisor's Ring (19519, -0.62 DPS) [rep]; Brainlash (6440, -1.02 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (270.8 DPS) | yes | Uther's Strength (11302, -0.03 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (270.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (270.8 DPS) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -6.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 169.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.27 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 470.7. Weights run: 1.1s. Verify run: 0.7s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.852 ± 0.052, crit=0.443 ± 0.010 per rating point (14 rating = 1%, 6.201 per %), hit=0.966 ± 0.105 per rating point (10 rating = 1%, 9.656 per %), spell_haste=3.490 ± 0.838, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.914 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 48.5 spell_power points (15.79 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.94 DPS) [pvp]; Crimson Felt Hat (18727, -3.80 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (470.7 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -0.20 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.93 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.0 spell_power points (14.97 DPS) | yes | Warlord's Silk Amice (231594, -2.67 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.81 DPS) [crafted]; Darkspear Shoulderpads (272103, -9.30 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.5 spell_power points (10.57 DPS) | yes | Hide of the Wild (18510, -3.24 DPS) [crafted]; Crystalline Threaded Cape (20697, -3.40 DPS, sim-verified) [world]; Deep Woodlands Cloak (19121, -4.17 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.4 spell_power points (18.36 DPS) | yes | Warlord's Silk Raiment (231596, -0.89 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.80 DPS) [pvp]; Robe of Everlasting Night (18385, -5.16 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.8 spell_power points (9.38 DPS) | yes | Sublime Wristguards (18497, -2.70 DPS) [dungeon]; Runecloth Cuffs (254123, -3.03 DPS) [crafted]; General's Silk Cuffs (16538, -4.11 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 33.6 spell_power points (10.93 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.4 spell_power points (17.38 DPS) | yes | Magician's Cord (272393, -4.54 DPS) [vendor]; Belt of the Archmage (18405, -7.17 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -7.28 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.9 spell_power points (17.22 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -3.46 DPS) [quest]; Outrider's Silk Leggings (22747, -8.52 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.6 spell_power points (11.27 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.98 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (470.7 DPS) | yes | Rune Band of Wizardry (22339, -1.86 DPS) [dungeon]; Maiden's Circle (13001, -2.41 DPS) [world_drop]; Naglering (11669, -12.32 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (470.7 DPS) | yes | Rune Band of Wizardry (22339, -0.43 DPS) [dungeon]; Maiden's Circle (13001, -0.98 DPS) [world_drop]; Naglering (11669, -11.73 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (470.7 DPS) | yes | Weakness Analyzer (272438, -2.28 DPS) [vendor]; Serenity Field (272439, -4.88 DPS) [vendor]; Blackhand's Breadth (13965, -5.40 DPS) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (470.7 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -2.87 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (470.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -1.45 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -26.14 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 241.9 spell_power points (78.71 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.28 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.39 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.26 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (orc, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 683.4. Weights run: 1.2s. Verify run: 0.7s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.147 ± 0.007, crit=0.485 ± 0.010 per rating point (14 rating = 1%, 6.791 per %), hit=1.003 ± 0.031 per rating point (10 rating = 1%, 10.030 per %), spell_haste=4.887 ± 0.143, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.904 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.5 spell_power points (15.51 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.51 DPS) [pvp] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (683.4 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -0.88 DPS) [dungeon]; Jewel of Kajaro (19601, -5.42 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.0 spell_power points (17.73 DPS) | yes | Warlord's Silk Amice (231594, -4.33 DPS) [pvp]; Argent Shoulders (19059, -5.42 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.52 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.2 spell_power points (13.40 DPS) | yes | Crystalline Threaded Cape (20697, -4.31 DPS, sim-verified) [world]; Amplifying Cloak (18350, -4.54 DPS) [dungeon]; Hide of the Wild (18510, -5.78 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.6 spell_power points (23.92 DPS) | yes | Warlord's Silk Raiment (231596, -3.09 DPS) [pvp]; Robe of Everlasting Night (18385, -7.00 DPS, sim-verified) [dungeon]; Legionnaire's Silk Tunic (227106, -9.00 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.2 spell_power points (11.42 DPS) | yes | Sublime Wristguards (18497, -4.78 DPS) [dungeon]; Runecloth Cuffs (254123, -5.27 DPS) [crafted]; Spidertank Oilrag (9448, -6.98 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (13.66 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.3 spell_power points (18.85 DPS) | yes | Ban'thok Sash (11662, -7.20 DPS) [dungeon]; Belt of the Archmage (18405, -7.63 DPS, sim-verified) [crafted]; Magician's Cord (272393, -7.75 DPS) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 47.1 spell_power points (23.18 DPS) | yes | General's Silk Trousers (231595, -3.60 DPS) [pvp]; Skyshroud Leggings (13170, -5.85 DPS) [dungeon]; Sorcerer's Leggings (226933, -7.72 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.82 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.32 DPS) [quest]; Sorcerer's Sandals (226931, -0.32 DPS) [vendor] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (683.4 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.97 DPS) [quest]; Blessed Band of Light (272407, -2.96 DPS) [vendor]; Naglering (11669, -20.76 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (683.4 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.78 DPS) [quest]; Blessed Band of Light (272407, -1.77 DPS) [vendor]; Naglering (11669, -12.85 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (683.4 DPS) | yes | Weakness Analyzer (272438, -3.45 DPS) [vendor]; Serenity Field (272439, -7.39 DPS) [vendor]; Blackhand's Breadth (13965, -7.59 DPS) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (683.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -3.74 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (683.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -2.09 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -31.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 161.8 spell_power points (79.72 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.45 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.38 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.27 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

