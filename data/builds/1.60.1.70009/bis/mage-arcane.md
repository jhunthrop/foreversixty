# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 153002000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 28.5. Weights run: 0.7s. Verify run: 0.6s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.271 ± 0.006, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.072 per %), hit=0.269 ± 0.007 per rating point (10 rating = 1%, 2.690 per %), spell_haste=0.344 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.461 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -1.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (0.71 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.33 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.72 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Caretaker's Cape (20428, -0.10 DPS) [rep]; Black Whelp Cloak (7283, -0.24 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.4 spell_power points (0.61 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -0.78 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.4 spell_power points (0.13 DPS) | yes | Windsong Bangles (263336, -0.03 DPS) [quest]; Repurposed Hair Band (281256, -0.08 DPS) [quest]; Bright Bracers (3647, -0.76 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.67 DPS) | yes | Gnoll Casting Gloves (892, -0.18 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.21 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.42 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.1 spell_power points (0.48 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.28 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.55 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.2 spell_power points (1.06 DPS) | yes | Filigreed Pristine Leggings (253937, -0.34 DPS) [crafted]; Silk-threaded Trousers (1929, -0.40 DPS) [dungeon]; Rumpled Kilt (274741, -0.59 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (0.77 DPS) | yes | Red Woolen Boots (4313, -0.39 DPS) [crafted]; Pristine Boots (253889, -0.41 DPS) [crafted]; Feather Padded Treads (285345, -0.76 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (0.53 DPS) | yes | Sludge-Stained Band (286535, -0.24 DPS) [world]; Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.45 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.48 DPS) | yes | Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop]; Sludge-Stained Band (286535, -0.47 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.7 spell_power points (0.26 DPS) | yes | Lesser Staff of the Spire (1300, -0.10 DPS) [world]; Staff of Westfall (2042, -0.13 DPS) [quest]; Channeler's Staff (4437, -0.17 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 237.0 spell_power points (22.60 DPS) | yes | Skycaller (12984, -0.51 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 153005113100010000-00000000000000000-0000000000000000000)

Set DPS (verified): 98.9. Weights run: 0.9s. Verify run: 0.7s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.348 ± 0.013, crit=0.062 ± 0.002 per rating point (14 rating = 1%, 0.873 per %), hit=0.252 ± 0.024 per rating point (10 rating = 1%, 2.517 per %), spell_haste=not significant (0.279 ± 0.187), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.915 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.94 DPS) | yes | Silk Headband (7050, -0.53 DPS) [crafted]; Embalmed Shroud (7691, -0.80 DPS) [dungeon]; Enchanter's Cowl (4322, -1.04 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.1 spell_power points (2.43 DPS) | yes | Crystal Starfire Medallion (5003, -2.06 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.06 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.61 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (3.24 DPS) | yes | Death Speaker Mantle (6685, -0.66 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.80 DPS) [quest]; Invoker's Mantle (215365, -0.91 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.34 DPS) | yes | Repairman's Cape (9605, -0.16 DPS) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Prelacy Cape (7004, -0.27 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.5 spell_power points (3.61 DPS) | yes | Tree Bark Jacket (1486, -0.14 DPS) [dungeon]; Death Speaker Robes (6682, -0.72 DPS) [dungeon]; Pristine Gown (253961, -1.09 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.40 DPS) | yes | Nightsky Wristbands (6407, -1.85 DPS) [world_drop]; Stonecloth Bindings (14416, -1.94 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.75 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.8 spell_power points (2.09 DPS) | yes | Serpent Gloves (5970, -0.22 DPS) [dungeon]; Shilly Mitts (9609, -0.22 DPS) [quest]; Truefaith Gloves (7049, -0.48 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.0 spell_power points (3.22 DPS) | yes | Belt of Arugal (6392, -0.53 DPS) [dungeon]; Invoker's Cord (215366, -0.88 DPS) [crafted]; Crimson Silk Belt (7055, -0.96 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.21 DPS) | yes | Pristine Leggings (253987, -0.68 DPS) [crafted]; Abomination Skin Leggings (23173, -0.69 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -1.04 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (2.52 DPS) | yes | Acidic Walkers (9454, -0.44 DPS) [dungeon]; Nimbus Boots (6998, -0.92 DPS) [quest]; Spidersilk Boots (4320, -2.65 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.87 DPS) | yes | Minor Channeling Ring (1449, -0.35 DPS) [quest]; Electrocutioner Lagnut (9447, -1.07 DPS) [dungeon]; Sludge-Stained Band (286535, -1.07 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.60 DPS) | yes | Electrocutioner Lagnut (9447, -0.80 DPS) [dungeon]; Sludge-Stained Band (286535, -0.80 DPS) [world]; Minor Channeling Ring (1449, -2.05 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.40 DPS) | yes | Twisted Chanter's Staff (890, -1.47 DPS) [world_drop]; Channeler's Staff (4437, -1.66 DPS) [world]; Glimmering Staff (249392, -1.92 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.1 spell_power points (2.43 DPS) | yes | Dwarven Tome (279898, -0.71 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.36 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -1.36 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 127.3 spell_power points (34.01 DPS) | yes | Starfaller (13063, -0.78 DPS) [world_drop]; Greater Mystic Wand (217287, -3.68 DPS) [crafted]; Gravestone Scepter (7001, -5.01 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 153005113100011531-00000000000000000-0000000000000000000)

Set DPS (verified): 186.0. Weights run: 0.9s. Verify run: 0.7s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.585 ± 0.027, crit=0.172 ± 0.004 per rating point (14 rating = 1%, 2.410 per %), hit=0.398 ± 0.052 per rating point (10 rating = 1%, 3.978 per %), spell_haste=not significant (1.325 ± 0.390), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.918 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.49 DPS) | yes | Augural Shroud (2620, -1.69 DPS, sim-verified) [world]; Living Cowl (5608, -2.47 DPS) [world]; Enchanter's Cowl (4322, -2.83 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.5 spell_power points (3.25 DPS) | yes | Triune Amulet (7722, -1.98 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.98 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.13 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.51 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.54 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.3 spell_power points (4.41 DPS) | yes | Guardian Cloak (5965, -1.65 DPS) [crafted]; Icy Cloak (4327, -2.24 DPS) [crafted]; Long Silken Cloak (4326, -3.95 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (7.88 DPS) | yes | Dreamweave Vest (10021, -0.69 DPS) [crafted]; Robe of Power (7054, -1.39 DPS) [crafted]; Elemental Raiment (9434, -1.39 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.78 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.62 DPS) [quest]; Windchaser Cuffs (14429, -1.15 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (6.28 DPS) | yes | Black Mageweave Gloves (10003, -1.65 DPS) [crafted]; Red Mageweave Gloves (10018, -2.02 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.55 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.3 spell_power points (5.05 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Star Belt (4329, -1.03 DPS) [crafted]; Gilded Cord (254037, -1.13 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.49 DPS) | yes | Abomination Skin Leggings (23173, -2.27 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.28 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.79 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.41 DPS) | yes | Gilded Slippers (254001, -1.92 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.42 DPS) [dungeon]; Spidersilk Boots (4320, -4.53 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.17 DPS) | yes | Ring of Forlorn Spirits (2043, -1.70 DPS) [quest]; Reedknot Ring (9622, -2.01 DPS) [quest]; Minor Channeling Ring (1449, -2.27 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.78 DPS) | yes | Reedknot Ring (9622, -0.62 DPS) [quest]; Minor Channeling Ring (1449, -0.87 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.27 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (186.0 DPS) | yes | Scorn's Focal Dagger (23168, -3.40 DPS) [dungeon]; Windweaver Staff (7757, -3.47 DPS) [dungeon]; Gut Ripper (2164, -10.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 131.0 spell_power points (40.46 DPS) | yes | Nether Force Wand (11263, -2.59 DPS) [quest]; Icefury Wand (7514, -2.73 DPS) [quest]; Ragefire Wand (7513, -2.78 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 153005113100011531-03202300000000000-0000000000000000000)

Set DPS (verified): 281.9. Weights run: 0.9s. Verify run: 1.0s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.713 ± 0.039, crit=0.264 ± 0.006 per rating point (14 rating = 1%, 3.692 per %), hit=0.676 ± 0.080 per rating point (10 rating = 1%, 6.757 per %), spell_haste=not significant (1.333 ± 0.597), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.921 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 33.3 spell_power points (10.49 DPS) | yes | Dreamweave Circlet (10041, -1.62 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.98 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.21 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.3 spell_power points (6.39 DPS) | yes | Mindburst Medallion (11196, -3.15 DPS) [quest]; Horizon Choker (13085, -3.25 DPS) [world_drop]; Scorn's Icy Choker (23169, -6.19 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.8 spell_power points (8.15 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.44 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.57 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.3 spell_power points (5.76 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.90 DPS) [dungeon]; Runecloth Cloak (13860, -1.13 DPS) [crafted]; Big Voodoo Cloak (8216, -2.16 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.3 spell_power points (10.49 DPS) | yes | Robe of the Magi (1716, -2.20 DPS) [world_drop]; Runecloth Tunic (13857, -2.65 DPS) [crafted]; Dreamweave Vest (10021, -2.79 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.0 spell_power points (3.78 DPS) | yes | Aristocratic Cuffs (12546, -0.41 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.94 DPS) [quest]; Bloodband Bracers (11469, -3.29 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.7 spell_power points (9.06 DPS) | yes | Dreamweave Gloves (10019, -2.49 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.94 DPS) [vendor]; Raider Handwraps (272098, -3.28 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.6 spell_power points (8.39 DPS) | yes | Dawnspire Cord (12466, -2.22 DPS) [dungeon]; Deathmage Sash (10771, -2.81 DPS) [dungeon]; Satyrmane Sash (17755, -5.95 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 30.1 spell_power points (9.50 DPS) | yes | Red Mageweave Pants (10009, -2.39 DPS) [crafted]; Wizardweave Leggings (14132, -3.51 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -5.43 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.57 DPS) | yes | Gilded Sandals (254107, -2.07 DPS) [crafted]; Black Mageweave Boots (10026, -2.52 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.77 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (4.50 DPS) | yes | Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.72 DPS) [rep]; Brainlash (6440, -1.13 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.0 spell_power points (4.41 DPS) | yes | Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.63 DPS) [rep]; Brainlash (6440, -1.04 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (281.9 DPS) | yes | Uther's Strength (11302, -2.13 DPS, sim-verified) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (281.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (281.9 DPS) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -8.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 166.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.23 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 492.4. Weights run: 1.0s. Verify run: 2.2s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.844 ± 0.052, crit=0.412 ± 0.009 per rating point (14 rating = 1%, 5.769 per %), hit=0.957 ± 0.105 per rating point (10 rating = 1%, 9.569 per %), spell_haste=3.513 ± 0.839, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.914 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.9 spell_power points (15.83 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.95 DPS) [pvp]; Crimson Felt Hat (18727, -3.68 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (492.4 DPS) | yes | Diana's Pearl Necklace (22403, -0.22 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.94 DPS) [quest]; Jewel of Kajaro (19601, -2.65 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.4 spell_power points (15.03 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.57 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.87 DPS) [crafted]; Darkspear Shoulderpads (272103, -9.29 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.3 spell_power points (10.69 DPS) | yes | Hide of the Wild (18510, -3.27 DPS) [crafted]; Crystalline Threaded Cape (20697, -4.28 DPS, sim-verified) [world]; Spritecaster Cape (11623, -4.39 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.9 spell_power points (18.49 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.92 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.89 DPS) [pvp]; Robe of Everlasting Night (18385, -5.44 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.8 spell_power points (9.51 DPS) | yes | Sublime Wristguards (18497, -2.75 DPS) [dungeon]; Runecloth Cuffs (254123, -3.08 DPS) [crafted]; Marshal's Silk Bracers (16438, -4.21 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 33.4 spell_power points (11.04 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.1 spell_power points (17.58 DPS) | yes | Magician's Cord (272393, -4.59 DPS) [vendor]; Ban'thok Sash (11662, -7.37 DPS) [dungeon]; Belt of the Archmage (18405, -8.45 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.0 spell_power points (17.19 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -3.28 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -3.59 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.5 spell_power points (11.41 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.99 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (492.4 DPS) | yes | Rune Band of Wizardry (22339, -1.89 DPS) [dungeon]; Maiden's Circle (13001, -2.44 DPS) [world_drop]; Naglering (11669, -13.03 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (492.4 DPS) | yes | Rune Band of Wizardry (22339, -0.44 DPS) [dungeon]; Maiden's Circle (13001, -0.99 DPS) [world_drop]; Naglering (11669, -12.14 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (492.4 DPS) | yes | Weakness Analyzer (272438, -2.32 DPS) [vendor]; Serenity Field (272439, -4.96 DPS) [vendor]; Burst of Knowledge (11832, -5.62 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (492.4 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -3.11 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (492.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Crackling Staff (19102, -0.63 DPS) [rep]; Teebu's Blazing Longsword (1728, -28.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 238.0 spell_power points (78.74 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.26 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.36 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.25 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 153005113100011531-03000000000000000-0545000300000000000)

Set DPS (verified): 732.3. Weights run: 1.2s. Verify run: 2.1s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.142 ± 0.005, crit=0.470 ± 0.010 per rating point (14 rating = 1%, 6.578 per %), hit=0.993 ± 0.033 per rating point (10 rating = 1%, 9.926 per %), spell_haste=4.842 ± 0.141, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.892 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.1 spell_power points (15.98 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.51 DPS) [pvp]; Sorcerer's Crown (226935, -4.09 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (732.3 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -0.99 DPS) [dungeon]; Jewel of Kajaro (19601, -5.85 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.7 spell_power points (18.33 DPS) | yes | Field Marshal's Silk Spaulders (231602, -4.40 DPS) [pvp]; Argent Shoulders (19059, -5.50 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.90 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.1 spell_power points (13.89 DPS) | yes | Amplifying Cloak (18350, -4.65 DPS) [dungeon]; Crystalline Threaded Cape (20697, -5.35 DPS, sim-verified) [world]; Hide of the Wild (18510, -5.98 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.3 spell_power points (24.78 DPS) | yes | Field Marshal's Silk Vestments (231603, -3.23 DPS) [pvp]; Robe of Everlasting Night (18385, -7.30 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Tunic (227108, -9.39 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.1 spell_power points (11.88 DPS) | yes | Sublime Wristguards (18497, -4.99 DPS) [dungeon]; Runecloth Cuffs (254123, -5.50 DPS) [crafted]; Arcane Runed Bracers (4744, -7.26 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (14.22 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.1 spell_power points (19.53 DPS) | yes | Ban'thok Sash (11662, -7.48 DPS) [dungeon]; Magician's Cord (272393, -8.03 DPS) [vendor]; Belt of the Archmage (18405, -8.81 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.6 spell_power points (23.91 DPS) | yes | Skyshroud Leggings (13170, -3.54 DPS, sim-verified) [dungeon]; Marshal's Silk Leggings (231605, -3.67 DPS) [pvp]; Sorcerer's Leggings (226933, -7.89 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.32 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.37 DPS) [quest]; Sorcerer's Sandals (226931, -0.37 DPS) [vendor] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (732.3 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -2.02 DPS) [quest]; Blessed Band of Light (272407, -3.08 DPS) [vendor]; Naglering (11669, -22.61 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (732.3 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.81 DPS) [quest]; Blessed Band of Light (272407, -1.87 DPS) [vendor]; Naglering (11669, -13.41 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (732.3 DPS) | yes | Weakness Analyzer (272438, -3.59 DPS) [vendor]; Serenity Field (272439, -7.70 DPS) [vendor]; Blackhand's Breadth (13965, -8.13 DPS) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (732.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -4.16 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (732.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -2.23 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -33.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.5 spell_power points (79.84 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.34 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.27 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.29 DPS) [world] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 153002000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 26.7. Weights run: 0.7s. Verify run: 0.6s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.271 ± 0.006, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.072 per %), hit=0.269 ± 0.007 per rating point (10 rating = 1%, 2.690 per %), spell_haste=0.344 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.461 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -0.69 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (0.71 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.16 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.33 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.10 DPS) [rep]; Black Whelp Cloak (7283, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.4 spell_power points (0.61 DPS) | yes | Green Woolen Vest (2582, -0.22 DPS) [crafted]; Bloody Apron (6226, -0.22 DPS) [dungeon]; Gray Woolen Robe (2585, -0.67 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.6 spell_power points (0.16 DPS) | yes | Owlbeard Bracers (16981, -0.01 DPS) [quest]; Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.67 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.21 DPS) [crafted]; Apothecary Gloves (10919, -0.29 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.1 spell_power points (0.48 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.28 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.41 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (26.7 DPS) | yes | Silk-threaded Trousers (1929, -0.06 DPS) [dungeon]; Rumpled Kilt (274741, -0.25 DPS) [vendor]; Abomination Skin Leggings (23173, -0.43 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (0.77 DPS) | yes | Red Woolen Boots (4313, -0.39 DPS) [crafted]; Feather Padded Treads (285345, -0.40 DPS, sim-verified) [world]; Pristine Boots (253889, -0.41 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.48 DPS) | yes | Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Loop of Sacrifice (281673, -0.35 DPS) [quest]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.29 DPS) | yes | Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Loop of Sacrifice (281673, -0.16 DPS) [quest]; Volcanic Rock Ring (12053, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.7 spell_power points (0.26 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.05 DPS) [world]; Lesser Staff of the Spire (1300, -0.10 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 237.0 spell_power points (22.60 DPS) | yes | Skycaller (12984, -0.34 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Sizzle Stick (8071, -4.47 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 153005113100010000-00000000000000000-0000000000000000000)

Set DPS (verified): 95.5. Weights run: 0.9s. Verify run: 0.7s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.348 ± 0.013, crit=0.062 ± 0.002 per rating point (14 rating = 1%, 0.873 per %), hit=0.252 ± 0.024 per rating point (10 rating = 1%, 2.517 per %), spell_haste=not significant (0.279 ± 0.187), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.915 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.94 DPS) | yes | Silk Headband (7050, -0.53 DPS) [crafted]; Embalmed Shroud (7691, -0.80 DPS) [dungeon]; Enchanter's Cowl (4322, -1.34 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.1 spell_power points (2.43 DPS) | yes | Crystal Starfire Medallion (5003, -2.06 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.06 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.67 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (3.24 DPS) | yes | Death Speaker Mantle (6685, -0.72 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.80 DPS) [quest]; Invoker's Mantle (215365, -0.91 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.34 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Battle Healer's Cloak (19529, -0.27 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.5 spell_power points (3.61 DPS) | yes | Tree Bark Jacket (1486, -0.14 DPS) [dungeon]; Death Speaker Robes (6682, -0.72 DPS) [dungeon]; Pristine Gown (253961, -1.09 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.40 DPS) | yes | Nightsky Wristbands (6407, -1.85 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.85 DPS) [quest]; Glowing Magical Bracelets (13106, -2.75 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.7 spell_power points (2.07 DPS) | yes | Serpent Gloves (5970, -0.20 DPS) [dungeon]; Truefaith Gloves (7049, -0.45 DPS) [crafted]; Gnoll Casting Gloves (892, -0.46 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (3.22 DPS) | yes | Warsong Sash (16975, -0.28 DPS) [quest]; Belt of Arugal (6392, -0.53 DPS) [dungeon]; Invoker's Cord (215366, -0.88 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.21 DPS) | yes | Pristine Leggings (253987, -0.68 DPS) [crafted]; Abomination Skin Leggings (23173, -0.87 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -1.04 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.4 spell_power points (2.52 DPS) | yes | Acidic Walkers (9454, -0.44 DPS) [dungeon]; Boots of the Enchanter (4325, -1.18 DPS) [crafted]; Spidersilk Boots (4320, -2.78 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.87 DPS) | yes | Electrocutioner Lagnut (9447, -1.07 DPS) [dungeon]; Sludge-Stained Band (286535, -1.07 DPS) [world]; Black Widow Band (6199, -1.22 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.60 DPS) | yes | Electrocutioner Lagnut (9447, -0.80 DPS) [dungeon]; Black Widow Band (6199, -0.95 DPS) [world]; Sludge-Stained Band (286535, -2.98 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.40 DPS) | yes | Twisted Chanter's Staff (890, -1.47 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.47 DPS) [quest]; Glimmering Staff (249392, -2.19 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.1 spell_power points (2.43 DPS) | yes | Orb of Souls (249395, -1.36 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.52 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.77 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 127.3 spell_power points (34.01 DPS) | yes | Starfaller (13063, -0.89 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.68 DPS) [crafted]; Gravestone Scepter (7001, -5.01 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 153005113100011531-00000000000000000-0000000000000000000)

Set DPS (verified): 180.3. Weights run: 0.9s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.585 ± 0.027, crit=0.172 ± 0.004 per rating point (14 rating = 1%, 2.410 per %), hit=0.398 ± 0.052 per rating point (10 rating = 1%, 3.978 per %), spell_haste=not significant (1.325 ± 0.390), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.918 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.49 DPS) | yes | Living Cowl (5608, -2.47 DPS) [world]; Augural Shroud (2620, -2.71 DPS, sim-verified) [world]; Enchanter's Cowl (4322, -2.83 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.5 spell_power points (3.25 DPS) | yes | Triune Amulet (7722, -1.98 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.98 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.36 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.51 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.54 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.3 spell_power points (4.41 DPS) | yes | Guardian Cloak (5965, -1.65 DPS) [crafted]; Long Silken Cloak (4326, -2.20 DPS, sim-verified) [crafted]; Icy Cloak (4327, -2.24 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (7.88 DPS) | yes | Dreamweave Vest (10021, -0.69 DPS) [crafted]; Robe of Power (7054, -1.39 DPS) [crafted]; Elemental Raiment (9434, -1.39 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.78 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.62 DPS) [quest]; Radiant Silver Bracers (4545, -1.87 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (6.28 DPS) | yes | Black Mageweave Gloves (10003, -1.65 DPS) [crafted]; Red Mageweave Gloves (10018, -2.45 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.55 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.3 spell_power points (5.05 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Star Belt (4329, -1.03 DPS) [crafted]; Gilded Cord (254037, -1.13 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.49 DPS) | yes | Abomination Skin Leggings (23173, -2.27 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.55 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.79 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.41 DPS) | yes | Gilded Slippers (254001, -2.74 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.42 DPS) [dungeon]; Spidersilk Boots (4320, -4.53 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.17 DPS) | yes | Reedknot Ring (9622, -2.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.32 DPS) [vendor]; Black Widow Band (6199, -2.91 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.78 DPS) | yes | Reedknot Ring (9622, -0.62 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.93 DPS) [vendor]; Black Widow Band (6199, -1.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (180.3 DPS) | yes | Scorn's Focal Dagger (23168, -3.40 DPS) [dungeon]; Windweaver Staff (7757, -3.47 DPS) [dungeon]; Gut Ripper (2164, -10.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 131.0 spell_power points (40.46 DPS) | yes | Nether Force Wand (11263, -2.59 DPS) [quest]; Icefury Wand (7514, -2.73 DPS) [quest]; Ragefire Wand (7513, -2.78 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 153005113100011531-03202300000000000-0000000000000000000)

Set DPS (verified): 274.4. Weights run: 0.9s. Verify run: 0.9s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.713 ± 0.039, crit=0.264 ± 0.006 per rating point (14 rating = 1%, 3.692 per %), hit=0.676 ± 0.080 per rating point (10 rating = 1%, 6.757 per %), spell_haste=not significant (1.333 ± 0.597), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.921 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 33.3 spell_power points (10.49 DPS) | yes | Dreamweave Circlet (10041, -1.62 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.98 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.21 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.3 spell_power points (6.39 DPS) | yes | Mindburst Medallion (11196, -3.15 DPS) [quest]; Horizon Choker (13085, -3.25 DPS) [world_drop]; Scorn's Icy Choker (23169, -6.11 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.8 spell_power points (8.15 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.44 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.57 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.4 spell_power points (5.81 DPS) | yes | Spritecaster Cape (11623, -0.04 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.95 DPS) [dungeon]; Runecloth Cloak (13860, -1.17 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.3 spell_power points (10.49 DPS) | yes | Robe of the Magi (1716, -2.20 DPS) [world_drop]; Runecloth Tunic (13857, -2.65 DPS) [crafted]; Dreamweave Vest (10021, -2.79 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.0 spell_power points (3.78 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -3.44 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.7 spell_power points (9.06 DPS) | yes | Raider Handwraps (272098, -1.91 DPS) [vendor]; Dreamweave Gloves (10019, -2.49 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.94 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.6 spell_power points (8.39 DPS) | yes | Dawnspire Cord (12466, -2.22 DPS) [dungeon]; Deathmage Sash (10771, -2.81 DPS) [dungeon]; Satyrmane Sash (17755, -4.09 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 30.1 spell_power points (9.50 DPS) | yes | Red Mageweave Pants (10009, -2.39 DPS) [crafted]; Wizardweave Leggings (14132, -3.51 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -5.57 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.57 DPS) | yes | Gilded Sandals (254107, -2.07 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.32 DPS, sim-verified) [vendor]; Black Mageweave Boots (10026, -2.52 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (4.50 DPS) | yes | Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Advisor's Ring (19519, -0.72 DPS) [rep]; Brainlash (6440, -1.13 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.0 spell_power points (4.41 DPS) | yes | Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Advisor's Ring (19519, -0.63 DPS) [rep]; Brainlash (6440, -1.04 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (274.4 DPS) | yes | Uther's Strength (11302, -0.03 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (274.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (274.4 DPS) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -6.99 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 166.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.23 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 477.0. Weights run: 1.0s. Verify run: 2.1s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.844 ± 0.052, crit=0.412 ± 0.009 per rating point (14 rating = 1%, 5.769 per %), hit=0.957 ± 0.105 per rating point (10 rating = 1%, 9.569 per %), spell_haste=3.513 ± 0.839, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.914 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.9 spell_power points (15.83 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.95 DPS) [pvp]; Crimson Felt Hat (18727, -3.68 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (477.0 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -0.22 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.94 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.4 spell_power points (15.03 DPS) | yes | Warlord's Silk Amice (231594, -2.57 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.87 DPS) [crafted]; Darkspear Shoulderpads (272103, -9.51 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.3 spell_power points (10.69 DPS) | yes | Crystalline Threaded Cape (20697, -3.00 DPS, sim-verified) [world]; Hide of the Wild (18510, -3.27 DPS) [crafted]; Deep Woodlands Cloak (19121, -4.21 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.9 spell_power points (18.49 DPS) | yes | Warlord's Silk Raiment (231596, -0.92 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.89 DPS) [pvp]; Robe of Everlasting Night (18385, -5.26 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.8 spell_power points (9.51 DPS) | yes | Sublime Wristguards (18497, -2.75 DPS) [dungeon]; Runecloth Cuffs (254123, -3.08 DPS) [crafted]; General's Silk Cuffs (16538, -4.21 DPS) [pvp] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 33.4 spell_power points (11.04 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.1 spell_power points (17.58 DPS) | yes | Magician's Cord (272393, -4.59 DPS) [vendor]; Belt of the Archmage (18405, -7.27 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -7.37 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.0 spell_power points (17.19 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -3.28 DPS) [quest]; Outrider's Silk Leggings (22747, -8.69 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.5 spell_power points (11.41 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.99 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (477.0 DPS) | yes | Rune Band of Wizardry (22339, -1.89 DPS) [dungeon]; Maiden's Circle (13001, -2.44 DPS) [world_drop]; Naglering (11669, -12.60 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (477.0 DPS) | yes | Rune Band of Wizardry (22339, -0.44 DPS) [dungeon]; Maiden's Circle (13001, -0.99 DPS) [world_drop]; Naglering (11669, -12.03 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (477.0 DPS) | yes | Weakness Analyzer (272438, -2.32 DPS) [vendor]; Serenity Field (272439, -4.96 DPS) [vendor]; Burst of Knowledge (11832, -5.62 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (477.0 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -2.92 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (477.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -1.59 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -26.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 238.0 spell_power points (78.74 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.26 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.36 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.25 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60, raid preset (orc, 153005113100011531-03000000000000000-0545000300000000000)

Set DPS (verified): 714.6. Weights run: 1.2s. Verify run: 2.2s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.142 ± 0.005, crit=0.470 ± 0.010 per rating point (14 rating = 1%, 6.578 per %), hit=0.993 ± 0.033 per rating point (10 rating = 1%, 9.926 per %), spell_haste=4.842 ± 0.141, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.892 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.1 spell_power points (15.98 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.51 DPS) [pvp]; Sorcerer's Crown (226935, -3.64 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (714.6 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -0.99 DPS) [dungeon]; Jewel of Kajaro (19601, -5.61 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.7 spell_power points (18.33 DPS) | yes | Warlord's Silk Amice (231594, -4.40 DPS) [pvp]; Argent Shoulders (19059, -5.50 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.76 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.1 spell_power points (13.89 DPS) | yes | Amplifying Cloak (18350, -4.65 DPS) [dungeon]; Crystalline Threaded Cape (20697, -5.04 DPS, sim-verified) [world]; Hide of the Wild (18510, -5.98 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.3 spell_power points (24.78 DPS) | yes | Warlord's Silk Raiment (231596, -3.23 DPS) [pvp]; Robe of Everlasting Night (18385, -7.39 DPS, sim-verified) [dungeon]; Legionnaire's Silk Tunic (227106, -9.39 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.1 spell_power points (11.88 DPS) | yes | Sublime Wristguards (18497, -4.99 DPS) [dungeon]; Runecloth Cuffs (254123, -5.50 DPS) [crafted]; Spidertank Oilrag (9448, -7.26 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (14.22 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.1 spell_power points (19.53 DPS) | yes | Ban'thok Sash (11662, -7.48 DPS) [dungeon]; Magician's Cord (272393, -8.03 DPS) [vendor]; Belt of the Archmage (18405, -8.34 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.6 spell_power points (23.91 DPS) | yes | General's Silk Trousers (231595, -3.67 DPS) [pvp]; Skyshroud Leggings (13170, -3.97 DPS, sim-verified) [dungeon]; Sorcerer's Leggings (226933, -7.89 DPS) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.32 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.37 DPS) [quest]; Sorcerer's Sandals (226931, -0.37 DPS) [vendor] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (714.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -2.02 DPS) [quest]; Blessed Band of Light (272407, -3.08 DPS) [vendor]; Naglering (11669, -21.88 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (714.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.81 DPS) [quest]; Blessed Band of Light (272407, -1.87 DPS) [vendor]; Naglering (11669, -13.22 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (714.6 DPS) | yes | Weakness Analyzer (272438, -3.59 DPS) [vendor]; Serenity Field (272439, -7.70 DPS) [vendor]; Blackhand's Breadth (13965, -8.13 DPS) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (714.6 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -3.78 DPS, sim-verified) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (714.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Balzaphon (23124, -2.23 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -33.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 155.5 spell_power points (79.84 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -10.34 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.27 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.29 DPS) [world] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Lord Valthalak's Staff of Command; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

