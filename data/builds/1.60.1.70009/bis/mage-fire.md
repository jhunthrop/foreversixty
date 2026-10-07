# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 27.8. Weights run: 1.1s. Verify run: 0.7s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.305 ± 0.013, crit=0.124 ± 0.004 per rating point (14 rating = 1%, 1.740 per %), hit=0.380 ± 0.002 per rating point (10 rating = 1%, 3.804 per %), spell_haste=-1.086 ± 0.157, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.40 DPS) | yes | Shadow Goggles (4373, -0.70 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.7 spell_power points (0.51 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.10 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Caretaker's Cape (20428, -0.07 DPS) [rep]; Black Whelp Cloak (7283, -0.23 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.43 DPS) | yes | Green Woolen Vest (2582, -0.17 DPS) [crafted]; Bloody Apron (6226, -0.17 DPS) [dungeon]; Gray Woolen Robe (2585, -0.57 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.5 spell_power points (0.10 DPS) | yes | Windsong Bangles (263336, -0.03 DPS) [quest]; Repurposed Hair Band (281256, -0.06 DPS) [quest]; Bright Bracers (3647, -0.56 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.46 DPS) | yes | Gnoll Casting Gloves (892, -0.07 DPS) [world]; Pristine Gloves (253913, -0.14 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.29 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.35 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Keller's Girdle (2911, -0.18 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.40 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (27.8 DPS) | yes | Silk-threaded Trousers (1929, -0.05 DPS) [dungeon]; Rumpled Kilt (274741, -0.19 DPS) [vendor]; Abomination Skin Leggings (23173, -0.40 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.54 DPS) | yes | Red Woolen Boots (4313, -0.28 DPS) [crafted]; Pristine Boots (253889, -0.29 DPS) [crafted]; Feather Padded Treads (285345, -0.47 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.37 DPS) | yes | Sludge-Stained Band (286535, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.33 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop]; Sludge-Stained Band (286535, -0.32 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.0 spell_power points (0.20 DPS) | yes | Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.08 DPS) [world]; Staff of Westfall (2042, -0.10 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 339.9 spell_power points (22.51 DPS) | yes | Skycaller (12984, -1.06 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.22 DPS) [dungeon]; Deepblaze (279896, -3.98 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 45.1. Weights run: 1.1s. Verify run: 0.8s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.623 ± 0.019, crit=0.175 ± 0.009 per rating point (14 rating = 1%, 2.446 per %), hit=0.340 ± 0.003 per rating point (10 rating = 1%, 3.400 per %), spell_haste=not significant (-0.012 ± 0.290), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.2 spell_power points (1.05 DPS) | yes | Holy Shroud (2721, -0.11 DPS) [world_drop]; Silk Headband (7050, -0.28 DPS) [crafted]; Embalmed Shroud (7691, -0.36 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.7 spell_power points (0.92 DPS) | yes | Crystal Starfire Medallion (5003, -0.71 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.94 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.6 spell_power points (1.25 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.26 DPS) [quest]; Magician's Mantle (12998, -0.34 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.5 spell_power points (0.47 DPS) | yes | Cloak of Rot (4462, -0.04 DPS) [world]; Darkspear Raider's Cloak (272078, -0.04 DPS) [vendor]; Hillman's Cloak (3719, -0.47 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.1 spell_power points (1.46 DPS) | yes | Tree Bark Jacket (1486, -0.35 DPS) [dungeon]; Pristine Gown (253961, -0.49 DPS) [crafted]; Death Speaker Robes (6682, -0.50 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.77 DPS) | yes | Nightsky Wristbands (6407, -0.45 DPS) [world_drop]; Stonecloth Bindings (14416, -0.50 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.08 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.9 spell_power points (0.93 DPS) | yes | Serpent Gloves (5970, -0.33 DPS) [dungeon]; Truefaith Gloves (7049, -0.34 DPS) [crafted]; Shilly Mitts (9609, -0.37 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.9 spell_power points (1.10 DPS) | yes | Belt of Arugal (6392, -0.17 DPS) [dungeon]; Crimson Silk Belt (7055, -0.21 DPS) [crafted]; Invoker's Cord (215366, -0.24 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.0 spell_power points (1.20 DPS) | yes | Gaze Dreamer Pants (6903, -0.17 DPS) [dungeon]; Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.4 spell_power points (0.97 DPS) | yes | Spidersilk Boots (4320, -0.16 DPS) [crafted]; Nimbus Boots (6998, -0.46 DPS) [quest]; Acidic Walkers (9454, -0.64 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.60 DPS) | yes | Black Widow Band (6199, -0.23 DPS) [world]; Snake Hoop (6750, -0.23 DPS) [quest]; Minor Channeling Ring (1449, -1.35 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (45.1 DPS) | yes | Black Widow Band (6199, -0.14 DPS) [world]; Snake Hoop (6750, -0.14 DPS) [quest]; Minor Channeling Ring (1449, -0.55 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.77 DPS) | yes | Twisted Chanter's Staff (890, -0.24 DPS) [world_drop]; Channeler's Staff (4437, -0.34 DPS) [world]; Glimmering Staff (249392, -0.42 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.7 spell_power points (0.92 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.53 DPS) [vendor]; Eye of Paleth (2943, -0.58 DPS) [quest]; Dwarven Tome (279898, -0.73 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 390.8 spell_power points (33.47 DPS) | yes | Starfaller (13063, -0.39 DPS) [world_drop]; Greater Mystic Wand (217287, -4.04 DPS) [crafted]; Gravestone Scepter (7001, -4.47 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-23552100130103050-0000000000000000000)

Set DPS (verified): 76.2. Weights run: 1.2s. Verify run: 0.8s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.788 ± 0.034, crit=0.259 ± 0.016 per rating point (14 rating = 1%, 3.632 per %), hit=0.437 ± 0.004 per rating point (10 rating = 1%, 4.369 per %), spell_haste=not significant (-0.273 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.21 DPS) | yes | Augural Shroud (2620, -0.22 DPS) [world]; Corpseshroud (10574, -0.63 DPS) [dungeon]; Enchanter's Cowl (4322, -0.75 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.7 spell_power points (1.23 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.40 DPS) [quest]; Triune Amulet (7722, -0.65 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.65 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.2 spell_power points (1.81 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.25 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.1 spell_power points (1.69 DPS) | yes | Guardian Cloak (5965, -0.65 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.78 DPS) [vendor]; Long Silken Cloak (4326, -1.68 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.7 spell_power points (2.81 DPS) | yes | Dreamweave Vest (10021, -0.17 DPS) [crafted]; Robe of Power (7054, -0.34 DPS) [crafted]; Elemental Raiment (9434, -0.60 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (76.2 DPS) | yes | Windchaser Cuffs (14429, -0.20 DPS) [world_drop]; Condor Bracers (15864, -0.21 DPS) [quest]; Arcane Runed Bracers (4744, -0.94 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (2.23 DPS) | yes | Black Mageweave Gloves (10003, -0.65 DPS) [crafted]; Gilded Handwraps (254021, -0.80 DPS) [crafted]; Red Mageweave Gloves (10018, -0.93 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.8 spell_power points (1.98 DPS) | yes | Highlander's Cloth Girdle (20098, -0.18 DPS) [rep]; Gilded Cord (254037, -0.48 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (2.47 DPS) | yes | Crimson Silk Pantaloons (7062, -0.72 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.86 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.53 DPS) | yes | Gilded Slippers (254001, -1.21 DPS) [crafted]; Acidic Walkers (9454, -1.34 DPS) [dungeon]; Spidersilk Boots (4320, -1.46 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (1.55 DPS) | yes | Ring of Forlorn Spirits (2043, -0.71 DPS) [quest]; Reedknot Ring (9622, -0.81 DPS) [quest]; Minor Channeling Ring (1449, -0.86 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.95 DPS) | yes | Reedknot Ring (9622, -0.21 DPS) [quest]; Minor Channeling Ring (1449, -0.25 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.19 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -0.86 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.16 DPS) [dungeon]; Gut Ripper (2164, -3.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 379.0 spell_power points (39.88 DPS) | yes | Nether Force Wand (11263, -1.84 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.44 DPS) [quest]; Ragefire Wand (7513, -2.49 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 132.1. Weights run: 1.1s. Verify run: 0.9s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.921 ± 0.043, crit=0.282 ± 0.021 per rating point (14 rating = 1%, 3.953 per %), hit=0.645 ± 0.008 per rating point (10 rating = 1%, 6.448 per %), spell_haste=not significant (-0.065 ± 0.852), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.4 spell_power points (4.92 DPS) | yes | Dreamweave Circlet (10041, -0.95 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.11 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.37 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (2.83 DPS) | yes | Scorn's Icy Choker (23169, -1.18 DPS) [dungeon]; Mindburst Medallion (11196, -1.31 DPS) [quest]; Horizon Choker (13085, -2.73 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.6 spell_power points (3.89 DPS) | yes | Kentic Amice (11624, -0.47 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.23 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 spell_power points (2.57 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon]; Runecloth Cloak (13860, -0.41 DPS) [crafted]; Big Voodoo Cloak (8216, -0.82 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.4 spell_power points (4.92 DPS) | yes | Runecloth Tunic (13857, -1.35 DPS) [crafted]; Runecloth Robe (13858, -1.41 DPS) [crafted]; Robe of the Magi (1716, -1.60 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.8 spell_power points (1.82 DPS) | yes | Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Bloodband Bracers (11469, -0.07 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.36 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.3 spell_power points (4.12 DPS) | yes | Raider Handwraps (272098, -0.54 DPS) [vendor]; Dreamweave Gloves (10019, -1.27 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.32 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.5 spell_power points (3.09 DPS) | yes | Satyrmane Sash (17755, -0.04 DPS) [dungeon]; Ban'thok Sash (11662, -0.09 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.2 spell_power points (4.23 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.68 DPS) [vendor]; Red Mageweave Pants (10009, -0.94 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.61 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (132.1 DPS) | yes | Gilded Sandals (254107, -0.45 DPS) [crafted]; Southsea Mojo Boots (20641, -0.61 DPS) [quest]; Earthen Silk Slippers (254013, -1.49 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.04 DPS) | yes | Brainlash (6440, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.33 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.46 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.21 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.45 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 399.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 324.7. Weights run: 1.3s. Verify run: 0.9s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.752 ± 0.060, crit=0.554 ± 0.039 per rating point (14 rating = 1%, 7.755 per %), hit=0.885 ± 0.011 per rating point (10 rating = 1%, 8.847 per %), spell_haste=not significant (2.312 ± 1.138), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.6 spell_power points (9.95 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.10 DPS) [pvp]; Crimson Felt Hat (18727, -2.41 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (324.7 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.58 DPS) [quest]; Chains of the Lich (23125, -0.58 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.0 spell_power points (9.63 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.04 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.87 DPS) [vendor]; Mantle of the Timbermaw (19050, -6.27 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.9 spell_power points (6.46 DPS) | yes | Crystalline Threaded Cape (20697, -1.64 DPS) [world]; Hide of the Wild (18510, -1.95 DPS) [crafted]; Spritecaster Cape (11623, -2.58 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.8 spell_power points (11.88 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.68 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.19 DPS) [pvp]; Robe of Everlasting Night (18385, -4.18 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.0 spell_power points (5.86 DPS) | yes | Sublime Wristguards (18497, -1.78 DPS) [dungeon]; Runecloth Cuffs (254123, -1.99 DPS) [crafted]; Sorcerer's Bindings (226929, -2.82 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 32.0 spell_power points (6.70 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.4 spell_power points (10.54 DPS) | yes | Magician's Cord (272393, -2.79 DPS) [vendor]; Highlander's Cloth Girdle (20047, -5.05 DPS) [rep]; Belt of the Archmage (18405, -6.71 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.0 spell_power points (11.51 DPS) | yes | Marshal's Silk Leggings (231605, -0.47 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -2.82 DPS) [pvp]; Skyshroud Leggings (13170, -3.14 DPS) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 33.0 spell_power points (6.91 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.63 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (324.7 DPS) | yes | Songstone of Ironforge (12543, -1.47 DPS) [quest]; Maiden's Circle (13001, -1.47 DPS) [world_drop]; Naglering (11669, -13.71 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (324.7 DPS) | yes | Songstone of Ironforge (12543, -0.63 DPS) [quest]; Maiden's Circle (13001, -0.63 DPS) [world_drop]; Naglering (11669, -14.65 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (324.7 DPS) | yes | Weakness Analyzer (272438, -1.46 DPS) [vendor]; Serenity Field (272439, -3.14 DPS) [vendor]; Burst of Knowledge (11832, -3.56 DPS) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (324.7 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Blackhand's Breadth (13965, -6.09 DPS, sim-verified) [quest] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (324.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.63 DPS) [world]; Teebu's Blazing Longsword (1728, -16.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 372.9 spell_power points (78.02 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.86 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.45 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.22 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 26.2. Weights run: 1.1s. Verify run: 0.8s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.305 ± 0.013, crit=0.124 ± 0.004 per rating point (14 rating = 1%, 1.740 per %), hit=0.380 ± 0.002 per rating point (10 rating = 1%, 3.804 per %), spell_haste=-1.086 ± 0.157, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.40 DPS) | yes | Shadow Goggles (4373, -0.56 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.7 spell_power points (0.51 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.33 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.26 DPS) | yes | Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.07 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.43 DPS) | yes | Green Woolen Vest (2582, -0.17 DPS) [crafted]; Bloody Apron (6226, -0.17 DPS) [dungeon]; Gray Woolen Robe (2585, -0.65 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.8 spell_power points (0.12 DPS) | yes | Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Owlbeard Bracers (16981, -0.28 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.46 DPS) | yes | Gnoll Casting Gloves (892, -0.07 DPS) [world]; Pristine Gloves (253913, -0.14 DPS) [crafted]; Apothecary Gloves (10919, -0.20 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.35 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Keller's Girdle (2911, -0.18 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.40 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (26.2 DPS) | yes | Silk-threaded Trousers (1929, -0.05 DPS) [dungeon]; Rumpled Kilt (274741, -0.19 DPS) [vendor]; Abomination Skin Leggings (23173, -0.26 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (0.54 DPS) | yes | Red Woolen Boots (4313, -0.28 DPS) [crafted]; Pristine Boots (253889, -0.29 DPS) [crafted]; Feather Padded Treads (285345, -0.54 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.33 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Loop of Sacrifice (281673, -0.23 DPS) [quest]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.20 DPS) | yes | Lavishly Jeweled Ring (1156, -0.08 DPS) [dungeon]; Loop of Sacrifice (281673, -0.10 DPS) [quest]; Volcanic Rock Ring (12053, -0.14 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.0 spell_power points (0.20 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.08 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 339.9 spell_power points (22.51 DPS) | yes | Skycaller (12984, -0.73 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.22 DPS) [dungeon]; Sizzle Stick (8071, -4.53 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 000000000000000000-23552100120000000-0000000000000000000)

Set DPS (verified): 42.4. Weights run: 1.1s. Verify run: 0.8s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.623 ± 0.019, crit=0.175 ± 0.009 per rating point (14 rating = 1%, 2.446 per %), hit=0.340 ± 0.003 per rating point (10 rating = 1%, 3.400 per %), spell_haste=not significant (-0.012 ± 0.290), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.2 spell_power points (1.05 DPS) | yes | Holy Shroud (2721, -0.11 DPS) [world_drop]; Silk Headband (7050, -0.28 DPS) [crafted]; Embalmed Shroud (7691, -0.36 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.7 spell_power points (0.92 DPS) | yes | Crystal Starfire Medallion (5003, -0.71 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.99 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.6 spell_power points (1.25 DPS) | yes | Death Speaker Mantle (6685, -0.15 DPS) [dungeon]; Fairywing Mantle (9536, -0.26 DPS) [quest]; Magician's Mantle (12998, -0.34 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.43 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Cloak of Rot (4462, -0.00 DPS) [world]; Darkspear Raider's Cloak (272078, -0.00 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.1 spell_power points (1.46 DPS) | yes | Tree Bark Jacket (1486, -0.35 DPS) [dungeon]; Pristine Gown (253961, -0.49 DPS) [crafted]; Death Speaker Robes (6682, -0.50 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.77 DPS) | yes | Nightsky Wristbands (6407, -0.45 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.45 DPS) [quest]; Glowing Magical Bracelets (13106, -0.69 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.1 spell_power points (0.78 DPS) | yes | Truefaith Gloves (7049, -0.19 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS) [world]; Serpent Gloves (5970, -0.71 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.9 spell_power points (1.10 DPS) | yes | Belt of Arugal (6392, -0.17 DPS) [dungeon]; Crimson Silk Belt (7055, -0.21 DPS) [crafted]; Warsong Sash (16975, -0.45 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.0 spell_power points (1.20 DPS) | yes | Gaze Dreamer Pants (6903, -0.17 DPS) [dungeon]; Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.4 spell_power points (0.97 DPS) | yes | Spidersilk Boots (4320, -0.16 DPS) [crafted]; Boots of the Enchanter (4325, -0.54 DPS) [crafted]; Acidic Walkers (9454, -0.93 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.60 DPS) | yes | Black Widow Band (6199, -0.23 DPS) [world]; Snake Hoop (6750, -0.23 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.51 DPS) | yes | Snake Hoop (6750, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Black Widow Band (6199, -0.37 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (0.77 DPS) | yes | Twisted Chanter's Staff (890, -0.24 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.24 DPS) [quest]; Glimmering Staff (249392, -0.33 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.7 spell_power points (0.92 DPS) | yes | Witch's Finger (16887, -0.55 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.58 DPS) [world]; Tome of the Darkspear Prophecy (272090, -0.83 DPS, sim-verified) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 390.8 spell_power points (33.47 DPS) | yes | Starfaller (13063, -0.39 DPS) [world_drop]; Greater Mystic Wand (217287, -4.04 DPS) [crafted]; Gravestone Scepter (7001, -4.47 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552100130103050-0000000000000000000)

Set DPS (verified): 72.2. Weights run: 1.2s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.788 ± 0.034, crit=0.259 ± 0.016 per rating point (14 rating = 1%, 3.632 per %), hit=0.437 ± 0.004 per rating point (10 rating = 1%, 4.369 per %), spell_haste=not significant (-0.273 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.21 DPS) | yes | Augural Shroud (2620, -0.22 DPS) [world]; Corpseshroud (10574, -0.63 DPS) [dungeon]; Enchanter's Cowl (4322, -0.75 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.7 spell_power points (1.23 DPS) | yes | Triune Amulet (7722, -0.65 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.65 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -0.76 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.2 spell_power points (1.81 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.25 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.1 spell_power points (1.69 DPS) | yes | Guardian Cloak (5965, -0.65 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.78 DPS) [vendor]; Long Silken Cloak (4326, -1.02 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.7 spell_power points (2.81 DPS) | yes | Dreamweave Vest (10021, -0.17 DPS) [crafted]; Robe of Power (7054, -0.34 DPS) [crafted]; Elemental Raiment (9434, -0.60 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.3 spell_power points (1.08 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Windchaser Cuffs (14429, -0.34 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (2.23 DPS) | yes | Black Mageweave Gloves (10003, -0.65 DPS) [crafted]; Red Mageweave Gloves (10018, -0.75 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -0.80 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.8 spell_power points (1.98 DPS) | yes | Gilded Cord (254037, -0.48 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.57 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.5 spell_power points (2.47 DPS) | yes | Abomination Skin Leggings (23173, -0.86 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.93 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.53 DPS) | yes | Gilded Slippers (254001, -1.21 DPS) [crafted]; Acidic Walkers (9454, -1.34 DPS) [dungeon]; Spidersilk Boots (4320, -1.46 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (1.55 DPS) | yes | Reedknot Ring (9622, -0.81 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.92 DPS) [vendor]; Black Widow Band (6199, -0.97 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.95 DPS) | yes | Reedknot Ring (9622, -0.21 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.32 DPS) [vendor]; Black Widow Band (6199, -0.37 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (72.2 DPS) | yes | Windweaver Staff (7757, -0.86 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.16 DPS) [dungeon]; Gut Ripper (2164, -3.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 379.0 spell_power points (39.88 DPS) | yes | Nether Force Wand (11263, -1.47 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.44 DPS) [quest]; Ragefire Wand (7513, -2.49 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 205011000000000000-23552100130103051-0000000000000000000)

Set DPS (verified): 126.7. Weights run: 1.1s. Verify run: 0.9s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.921 ± 0.043, crit=0.282 ± 0.021 per rating point (14 rating = 1%, 3.953 per %), hit=0.645 ± 0.008 per rating point (10 rating = 1%, 6.448 per %), spell_haste=not significant (-0.065 ± 0.852), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.4 spell_power points (4.92 DPS) | yes | Dreamweave Circlet (10041, -0.95 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.11 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.37 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (2.83 DPS) | yes | Scorn's Icy Choker (23169, -1.18 DPS) [dungeon]; Mindburst Medallion (11196, -1.31 DPS) [quest]; Horizon Choker (13085, -3.11 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.6 spell_power points (3.89 DPS) | yes | Kentic Amice (11624, -0.47 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.23 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.3 spell_power points (2.67 DPS) | yes | Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.39 DPS) [dungeon]; Runecloth Cloak (13860, -0.52 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.4 spell_power points (4.92 DPS) | yes | Robe of the Magi (1716, -1.30 DPS) [world_drop]; Runecloth Tunic (13857, -1.35 DPS) [crafted]; Runecloth Robe (13858, -1.41 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.8 spell_power points (1.82 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.05 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 31.3 spell_power points (4.12 DPS) | yes | Dreamweave Gloves (10019, -1.27 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.32 DPS) [vendor]; Raider Handwraps (272098, -1.44 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.5 spell_power points (3.09 DPS) | yes | Satyrmane Sash (17755, -0.04 DPS) [dungeon]; Ban'thok Sash (11662, -0.09 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.2 spell_power points (4.23 DPS) | yes | Red Mageweave Pants (10009, -0.94 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.61 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.85 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.17 DPS) [vendor]; Gilded Sandals (254107, -0.62 DPS) [crafted]; Southsea Mojo Boots (20641, -0.77 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.04 DPS) | yes | Brainlash (6440, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.33 DPS) [world_drop]; Advisor's Ring (19519, -0.46 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.03 DPS) | yes | Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Advisor's Ring (19519, -0.45 DPS) [rep]; Brainlash (6440, -1.84 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (126.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (126.7 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -0.73 DPS) [quest]; Spellforce Rod (1664, -0.88 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 399.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 205015100000000000-23552100130103051-0050000000000000000)

Set DPS (verified): 312.3. Weights run: 1.3s. Verify run: 0.9s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.752 ± 0.060, crit=0.554 ± 0.039 per rating point (14 rating = 1%, 7.755 per %), hit=0.885 ± 0.011 per rating point (10 rating = 1%, 8.847 per %), spell_haste=not significant (2.312 ± 1.138), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.6 spell_power points (9.95 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.10 DPS) [pvp]; Crimson Felt Hat (18727, -2.41 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (312.3 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.58 DPS) [quest]; Chains of the Lich (23125, -0.58 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.0 spell_power points (9.63 DPS) | yes | Warlord's Silk Amice (231594, -2.04 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.41 DPS) [crafted]; Darkspear Shoulderpads (272103, -2.87 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.9 spell_power points (6.46 DPS) | yes | Crystalline Threaded Cape (20697, -1.64 DPS) [world]; Hide of the Wild (18510, -1.95 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.53 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.8 spell_power points (11.88 DPS) | yes | Warlord's Silk Raiment (231596, -0.68 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.19 DPS) [pvp]; Robe of Everlasting Night (18385, -6.19 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.0 spell_power points (5.86 DPS) | yes | Sublime Wristguards (18497, -1.78 DPS) [dungeon]; Runecloth Cuffs (254123, -1.99 DPS) [crafted]; Sorcerer's Bindings (226929, -2.82 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 32.0 spell_power points (6.70 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.4 spell_power points (10.54 DPS) | yes | Belt of the Archmage (18405, -2.22 DPS) [crafted]; Magician's Cord (272393, -2.79 DPS) [vendor]; Defiler's Cloth Girdle (20163, -5.05 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.0 spell_power points (11.51 DPS) | yes | General's Silk Trousers (231595, -0.47 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.67 DPS) [rep]; Legionnaire's Silk Legguards (227107, -2.82 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 33.0 spell_power points (6.91 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.63 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (312.3 DPS) | yes | Eye of Orgrimmar (12545, -1.47 DPS) [quest]; Maiden's Circle (13001, -1.47 DPS) [world_drop]; Naglering (11669, -11.39 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (312.3 DPS) | yes | Eye of Orgrimmar (12545, -0.63 DPS) [quest]; Maiden's Circle (13001, -0.63 DPS) [world_drop]; Naglering (11669, -11.90 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (312.3 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Second Wind (11819, -5.59 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (312.3 DPS) | yes | Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, -1.46 DPS) [vendor]; Serenity Field (272439, -3.14 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (312.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.49 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 372.9 spell_power points (78.02 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.86 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.45 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.22 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

