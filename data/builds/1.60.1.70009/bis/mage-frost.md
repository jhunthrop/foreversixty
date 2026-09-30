# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.5. Weights run: 0.6s. Verify run: 0.6s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.375 ± 0.009, crit=0.074 ± 0.003 per rating point (14 rating = 1%, 1.030 per %), hit=0.198 ± 0.002 per rating point (10 rating = 1%, 1.980 per %), spell_haste=0.719 ± 0.179, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.77 DPS) | yes | Shadow Goggles (4373, -1.42 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (1.07 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.56 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.51 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Caretaker's Cape (20428, -0.13 DPS) [rep]; Pearl-clasped Cloak (5542, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.88 DPS) | yes | Green Woolen Robe (6243, -0.35 DPS) [crafted]; Bloody Apron (6226, -0.37 DPS) [dungeon]; Gray Woolen Robe (2585, -1.02 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.9 spell_power points (0.24 DPS) | yes | Windsong Bangles (263336, -0.11 DPS) [quest]; Repurposed Hair Band (281256, -0.14 DPS) [quest]; Bright Bracers (3647, -0.59 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.90 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.54 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.70 DPS) | yes | Novice Ardent's Sash (253887, -0.30 DPS) [crafted]; Keller's Girdle (2911, -0.32 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.60 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.0 spell_power points (1.54 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.64 DPS) [dungeon]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (1.09 DPS) | yes | Pristine Boots (253889, -0.56 DPS) [crafted]; Red Woolen Boots (4313, -0.58 DPS) [crafted]; Feather Padded Treads (285345, -0.74 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.74 DPS) | yes | Sludge-Stained Band (286535, -0.35 DPS) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.59 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop]; Sludge-Stained Band (286535, -0.83 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.8 spell_power points (0.48 DPS) | yes | Lesser Staff of the Spire (1300, -0.19 DPS) [world]; Staff of Westfall (2042, -0.24 DPS) [quest]; Channeler's Staff (4437, -0.35 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 177.1 spell_power points (22.69 DPS) | yes | Skycaller (12984, -1.07 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.40 DPS) [dungeon]; Deepblaze (279896, -4.16 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 62.0. Weights run: 0.6s. Verify run: 0.5s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.471 ± 0.016, crit=0.138 ± 0.008 per rating point (14 rating = 1%, 1.937 per %), hit=0.230 ± 0.003 per rating point (10 rating = 1%, 2.304 per %), spell_haste=not significant (0.678 ± 0.285), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.65 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.47 DPS) | yes | Crystal Starfire Medallion (5003, -1.19 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.19 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.60 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (1.99 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.58 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Prelacy Cape (7004, -0.15 DPS) [quest]; Caretaker's Cape (19533, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.1 spell_power points (2.27 DPS) | yes | Death Speaker Robes (6682, -0.44 DPS) [dungeon]; Pristine Gown (253961, -0.72 DPS) [crafted]; Tree Bark Jacket (1486, -0.96 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Nightsky Wristbands (6407, -0.93 DPS) [world_drop]; Stonecloth Bindings (14416, -1.00 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.29 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 9.2 spell_power points (1.38 DPS) | yes | Serpent Gloves (5970, -0.33 DPS) [dungeon]; Truefaith Gloves (7049, -0.42 DPS) [crafted]; Shilly Mitts (9609, -0.82 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.4 spell_power points (1.86 DPS) | yes | Belt of Arugal (6392, -0.18 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.46 DPS) [crafted]; Crimson Silk Belt (7055, -0.47 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.8 spell_power points (1.92 DPS) | yes | Gaze Dreamer Pants (6903, -0.18 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.37 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.55 DPS) | yes | Acidic Walkers (9454, -0.23 DPS) [dungeon]; Nimbus Boots (6998, -0.64 DPS) [quest]; Spidersilk Boots (4320, -1.81 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.05 DPS) | yes | Minor Channeling Ring (1449, -0.16 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Snake Hoop (6750, -0.56 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.90 DPS) | yes | Black Widow Band (6199, -0.41 DPS) [world]; Snake Hoop (6750, -0.41 DPS) [quest]; Minor Channeling Ring (1449, -1.13 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Glimmering Staff (249392, -0.38 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.64 DPS) [world_drop]; Channeler's Staff (4437, -0.78 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.8 spell_power points (1.47 DPS) | yes | Eye of Paleth (2943, -0.87 DPS) [quest]; Orb of Souls (249395, -0.87 DPS) [crafted]; Dwarven Tome (279898, -0.95 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 224.3 spell_power points (33.66 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 98.3. Weights run: 0.6s. Verify run: 0.5s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.717 ± 0.034, crit=0.206 ± 0.012 per rating point (14 rating = 1%, 2.880 per %), hit=0.372 ± 0.005 per rating point (10 rating = 1%, 3.716 per %), spell_haste=not significant (1.576 ± 0.546), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.09 DPS) | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Corpseshroud (10574, -1.08 DPS) [dungeon]; Enchanter's Cowl (4322, -1.15 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.3 spell_power points (1.66 DPS) | yes | Necklace of Calisea (1714, -0.92 DPS) [dungeon]; Triune Amulet (7722, -0.92 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.15 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.3 spell_power points (2.40 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.5 spell_power points (2.27 DPS) | yes | Guardian Cloak (5965, -0.86 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.11 DPS) [vendor]; Long Silken Cloak (4326, -1.20 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.3 spell_power points (3.87 DPS) | yes | Dreamweave Vest (10021, -0.27 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.54 DPS) [crafted]; Elemental Raiment (9434, -0.78 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.32 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest]; Windchaser Cuffs (14429, -0.37 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.9 spell_power points (3.07 DPS) | yes | Black Mageweave Gloves (10003, -0.86 DPS) [crafted]; Gilded Handwraps (254021, -1.15 DPS) [crafted]; Red Mageweave Gloves (10018, -1.30 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.8 spell_power points (2.61 DPS) | yes | Highlander's Cloth Girdle (20098, -0.22 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.59 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.6 spell_power points (3.33 DPS) | yes | Crimson Silk Pantaloons (7062, -1.08 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.16 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.45 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.53 DPS) | yes | Gilded Slippers (254001, -1.62 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.95 DPS) [dungeon]; Spidersilk Boots (4320, -2.08 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (2.10 DPS) | yes | Ring of Forlorn Spirits (2043, -0.93 DPS) [quest]; Reedknot Ring (9622, -1.07 DPS) [quest]; Minor Channeling Ring (1449, -1.16 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.32 DPS) | yes | Reedknot Ring (9622, -0.29 DPS) [quest]; Lorekeeper's Ring (19525, -0.29 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.46 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (98.3 DPS) | yes | Windweaver Staff (7757, -1.36 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Gut Ripper (2164, -4.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 272.0 spell_power points (40.01 DPS) | yes | Nether Force Wand (11263, -0.91 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.51 DPS) [quest]; Ragefire Wand (7513, -2.56 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 134.9. Weights run: 0.6s. Verify run: 0.5s. 422 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.911 ± 0.050, crit=0.368 ± 0.021 per rating point (14 rating = 1%, 5.153 per %), hit=0.574 ± 0.009 per rating point (10 rating = 1%, 5.744 per %), spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.2 spell_power points (4.88 DPS) | yes | Dreamweave Circlet (10041, -0.14 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -0.94 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.34 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (2.82 DPS) | yes | Scorn's Icy Choker (23169, -1.18 DPS) [dungeon]; Mindburst Medallion (11196, -1.31 DPS) [quest]; Horizon Choker (13085, -2.98 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.4 spell_power points (3.86 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.06 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 spell_power points (2.55 DPS) | yes | Runecloth Cloak (13860, -0.42 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -0.68 DPS, sim-verified) [dungeon]; Big Voodoo Cloak (8216, -0.82 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.2 spell_power points (4.88 DPS) | yes | Robe of the Magi (1716, -1.09 DPS, sim-verified) [world_drop]; Knight's Dreadweave Vest (220886, -1.32 DPS) [vendor]; Runecloth Tunic (13857, -1.34 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.7 spell_power points (1.79 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -0.06 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.36 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.5 spell_power points (4.00 DPS) | yes | Raider Handwraps (272098, -0.58 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.16 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.22 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.3 spell_power points (3.06 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.09 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.1 spell_power points (4.21 DPS) | yes | Red Mageweave Pants (10009, -0.94 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.61 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.70 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | Gilded Sandals (254107, -0.63 DPS) [crafted]; Southsea Mojo Boots (20641, -0.78 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -0.85 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.45 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.02 DPS) | yes | Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.44 DPS) [rep]; Brainlash (6440, -1.69 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (134.9 DPS) | yes | Blade of Eternal Darkness (17780, -0.64 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.72 DPS) [quest]; Spellforce Rod (1664, -0.84 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 259.9. Weights run: 0.6s. Verify run: 0.6s. 1030 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.019 ± 0.072, crit=0.597 ± 0.035 per rating point (14 rating = 1%, 8.359 per %), hit=1.081 ± 0.014 per rating point (10 rating = 1%, 10.815 per %), spell_haste=not significant (4.959 ± 1.418), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 87.7 spell_power points (10.76 DPS) | yes | Fireleaf Hood (240048, -1.56 DPS, sim-verified) [vendor]; Field Marshal's Coronet (231604, -3.56 DPS) [pvp] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (259.9 DPS) | yes | Beads of Ogre Mojo (22149, -0.37 DPS) [quest]; Pebble of Kajaro (19600, -0.74 DPS) [quest]; Jewel of Kajaro (19601, -1.54 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 64.6 spell_power points (7.93 DPS) | yes | Fireleaf Mantle (240046, -1.37 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -1.72 DPS) [vendor]; Field Marshal's Silk Spaulders (16444, -2.99 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.0 spell_power points (4.29 DPS) | yes | Crystalline Threaded Cape (20697, -1.34 DPS) [world]; Spritecaster Cape (11623, -1.82 DPS) [dungeon]; Hide of the Wild (18510, -2.21 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 99.1 spell_power points (12.15 DPS) | yes | Fireleaf Garb (240051, -1.98 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -4.72 DPS) [crafted]; Field Marshal's Silk Vestments (16443, -4.95 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 59.7 spell_power points (7.32 DPS) | yes | Fireleaf Wristwraps (240044, -2.38 DPS, sim-verified) [vendor]; Dryad's Wrist Bindings (19595, -3.62 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 73.0 spell_power points (8.95 DPS) | yes | Fireleaf Mitts (240049, -1.42 DPS, sim-verified) [vendor]; Marshal's Silk Gloves (16440, -4.14 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -4.14 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 70.1 spell_power points (8.59 DPS) | yes | Knowledge of the Timbermaw (228190, -1.45 DPS) [vendor]; Fireleaf Waistguard (240045, -2.76 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -3.12 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 87.7 spell_power points (10.76 DPS) | yes | Fireleaf Pants (240047, -1.70 DPS, sim-verified) [vendor]; Sentinel's Silk Leggings (237815, -3.53 DPS) [vendor]; Marshal's Silk Leggings (231605, -3.55 DPS) [pvp] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 65.6 spell_power points (8.05 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (231606, -2.40 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (259.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.37 DPS) [vendor]; Naglering (11669, -12.54 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (259.9 DPS) | yes | Songstone of Ironforge (12543, -0.99 DPS) [quest]; Maiden's Circle (13001, -0.99 DPS) [world_drop]; Naglering (11669, -8.11 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (259.9 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Second Wind (11819, -0.05 DPS, sim-verified) [dungeon] |
| trinket2 | - | - |  |  |  |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (259.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.90 DPS) [world]; Teebu's Blazing Longsword (1728, -13.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 631.8 spell_power points (77.50 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -4.58 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -13.02 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.43 DPS) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1030, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 33.6. Weights run: 0.6s. Verify run: 0.5s. 141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.375 ± 0.009, crit=0.074 ± 0.003 per rating point (14 rating = 1%, 1.030 per %), hit=0.198 ± 0.002 per rating point (10 rating = 1%, 1.980 per %), spell_haste=0.719 ± 0.179, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.77 DPS) | yes | Shadow Goggles (4373, -1.21 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (1.07 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.50 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.56 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.51 DPS) | yes | Pearl-clasped Cloak (5542, -0.06 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.13 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.88 DPS) | yes | Green Woolen Robe (6243, -0.35 DPS) [crafted]; Bloody Apron (6226, -0.37 DPS) [dungeon]; Gray Woolen Robe (2585, -1.22 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.3 spell_power points (0.29 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.05 DPS) [quest]; Owlbeard Bracers (16981, -0.06 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.90 DPS) | yes | Gnoll Casting Gloves (892, -0.15 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Apothecary Gloves (10919, -0.38 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.70 DPS) | yes | Novice Ardent's Sash (253887, -0.30 DPS) [crafted]; Keller's Girdle (2911, -0.32 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.84 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.0 spell_power points (1.54 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.64 DPS) [dungeon]; Rumpled Kilt (274741, -0.90 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (1.09 DPS) | yes | Pristine Boots (253889, -0.56 DPS) [crafted]; Red Woolen Boots (4313, -0.58 DPS) [crafted]; Feather Padded Treads (285345, -0.90 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Loop of Sacrifice (281673, -0.40 DPS) [quest]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.38 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 3.8 spell_power points (0.48 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.10 DPS) [world]; Lesser Staff of the Spire (1300, -0.19 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 177.1 spell_power points (22.69 DPS) | yes | Skycaller (12984, -1.44 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.40 DPS) [dungeon]; Sizzle Stick (8071, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 56.1. Weights run: 0.6s. Verify run: 0.6s. 236 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.471 ± 0.016, crit=0.138 ± 0.008 per rating point (14 rating = 1%, 1.937 per %), hit=0.230 ± 0.003 per rating point (10 rating = 1%, 2.304 per %), spell_haste=not significant (0.678 ± 0.285), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.65 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.47 DPS) | yes | Crystal Starfire Medallion (5003, -1.19 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.19 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.49 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (1.99 DPS) | yes | Death Speaker Mantle (6685, -0.24 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.58 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.1 spell_power points (2.27 DPS) | yes | Death Speaker Robes (6682, -0.44 DPS) [dungeon]; Pristine Gown (253961, -0.72 DPS) [crafted]; Tree Bark Jacket (1486, -1.17 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Nightsky Wristbands (6407, -0.93 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.93 DPS) [quest]; Glowing Magical Bracelets (13106, -1.24 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.4 spell_power points (1.25 DPS) | yes | Truefaith Gloves (7049, -0.29 DPS) [crafted]; Gnoll Casting Gloves (892, -0.35 DPS) [world]; Serpent Gloves (5970, -0.39 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.4 spell_power points (1.86 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Invoker's Cord (215366, -0.46 DPS) [crafted]; Warsong Sash (16975, -0.57 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.8 spell_power points (1.92 DPS) | yes | Gaze Dreamer Pants (6903, -0.30 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.37 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.3 spell_power points (1.55 DPS) | yes | Acidic Walkers (9454, -0.23 DPS) [dungeon]; Boots of the Enchanter (4325, -0.79 DPS) [crafted]; Spidersilk Boots (4320, -1.78 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.05 DPS) | yes | Advisor's Ring (20426, -0.30 DPS) [rep]; Black Widow Band (6199, -0.56 DPS) [world]; Snake Hoop (6750, -0.56 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.90 DPS) | yes | Black Widow Band (6199, -0.41 DPS) [world]; Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Snake Hoop (6750, -1.04 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Glimmering Staff (249392, -0.60 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.64 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.64 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.8 spell_power points (1.47 DPS) | yes | Orb of Souls (249395, -0.25 DPS, sim-verified) [crafted]; Alliance Outrunner Healing Rod (285348, -0.87 DPS) [world]; Tome of the Darkspear Prophecy (272090, -0.89 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 224.3 spell_power points (33.66 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 236, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 89.5. Weights run: 0.6s. Verify run: 0.5s. 318 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.717 ± 0.034, crit=0.206 ± 0.012 per rating point (14 rating = 1%, 2.880 per %), hit=0.372 ± 0.005 per rating point (10 rating = 1%, 3.716 per %), spell_haste=not significant (1.576 ± 0.546), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.09 DPS) | yes | Augural Shroud (2620, -0.49 DPS, sim-verified) [world]; Corpseshroud (10574, -1.08 DPS) [dungeon]; Enchanter's Cowl (4322, -1.15 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.3 spell_power points (1.66 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.50 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -0.92 DPS) [dungeon]; Triune Amulet (7722, -0.92 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.3 spell_power points (2.40 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.5 spell_power points (2.27 DPS) | yes | Guardian Cloak (5965, -0.86 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.11 DPS) [vendor]; Long Silken Cloak (4326, -1.32 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.3 spell_power points (3.87 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.54 DPS) [crafted]; Elemental Raiment (9434, -0.78 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.7 spell_power points (1.43 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.40 DPS) [quest]; Windchaser Cuffs (14429, -0.48 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.9 spell_power points (3.07 DPS) | yes | Black Mageweave Gloves (10003, -0.86 DPS) [crafted]; Red Mageweave Gloves (10018, -1.07 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.15 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.8 spell_power points (2.61 DPS) | yes | Gilded Cord (254037, -0.59 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.68 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.6 spell_power points (3.33 DPS) | yes | Crimson Silk Pantaloons (7062, -0.76 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.16 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.45 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.53 DPS) | yes | Gilded Slippers (254001, -1.42 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.95 DPS) [dungeon]; Spidersilk Boots (4320, -2.08 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (2.10 DPS) | yes | Reedknot Ring (9622, -1.07 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.22 DPS) [vendor]; Voodoo Band (1996, -1.37 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.32 DPS) | yes | Advisor's Ring (19521, -0.29 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.44 DPS) [vendor]; Reedknot Ring (9622, -1.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (89.5 DPS) | yes | Windweaver Staff (7757, -1.36 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Gut Ripper (2164, -4.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 272.0 spell_power points (40.01 DPS) | yes | Nether Force Wand (11263, -2.36 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.51 DPS) [quest]; Ragefire Wand (7513, -2.56 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 122.7. Weights run: 0.6s. Verify run: 0.6s. 412 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.911 ± 0.050, crit=0.368 ± 0.021 per rating point (14 rating = 1%, 5.153 per %), hit=0.574 ± 0.009 per rating point (10 rating = 1%, 5.744 per %), spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.2 spell_power points (4.88 DPS) | yes | Dreamweave Circlet (10041, +0.00 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Hat (220907, -0.94 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.34 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (2.82 DPS) | yes | Horizon Choker (13085, -0.60 DPS, sim-verified) [world_drop]; Scorn's Icy Choker (23169, -1.18 DPS) [dungeon]; Mindburst Medallion (11196, -1.31 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Dreadweave Mantle (220905, -0.59 DPS) [vendor]; Red Mageweave Shoulders (10029, -0.68 DPS) [crafted]; Rotgrip Mantle (17732, -1.25 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.2 spell_power points (2.65 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.39 DPS) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.2 spell_power points (4.88 DPS) | yes | Robe of the Magi (1716, -0.68 DPS, sim-verified) [world_drop]; Stone Guard's Dreadweave Vest (220904, -1.32 DPS) [vendor]; Runecloth Tunic (13857, -1.34 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodband Bracers (11469, -0.02 DPS) [quest]; Radiant Silver Bracers (4545, -0.27 DPS) [quest]; Aristocratic Cuffs (12546, -2.66 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.5 spell_power points (4.00 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.16 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.22 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.3 spell_power points (3.06 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.09 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.1 spell_power points (4.21 DPS) | yes | Red Mageweave Pants (10009, -0.94 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -1.35 DPS, sim-verified) [vendor]; Crimson Silk Pantaloons (7062, -1.61 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -0.63 DPS) [crafted]; Southsea Mojo Boots (20641, -0.78 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Advisor's Ring (19519, -0.45 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.02 DPS) | yes | Brainlash (6440, +0.00 DPS, sim-verified) [dungeon]; Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Advisor's Ring (19519, -0.44 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.72 DPS) [quest]; Spellforce Rod (1664, -0.84 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Kentic Amice; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 412, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 242.3. Weights run: 0.6s. Verify run: 0.6s. 1021 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.019 ± 0.072, crit=0.597 ± 0.035 per rating point (14 rating = 1%, 8.359 per %), hit=1.081 ± 0.014 per rating point (10 rating = 1%, 10.815 per %), spell_haste=not significant (4.959 ± 1.418), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 87.7 spell_power points (10.76 DPS) | yes | Fireleaf Hood (240048, -1.69 DPS, sim-verified) [vendor]; Warlord's Silk Cowl (231601, -3.56 DPS) [pvp] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (242.3 DPS) | yes | Jewel of Kajaro (19601, -0.09 DPS, sim-verified) [quest]; Beads of Ogre Mojo (22149, -0.37 DPS) [quest]; Pebble of Kajaro (19600, -0.74 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 64.6 spell_power points (7.93 DPS) | yes | Fireleaf Mantle (240046, -1.51 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -1.72 DPS) [vendor]; Warlord's Silk Amice (16536, -2.99 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.0 spell_power points (4.29 DPS) | yes | Hide of the Wild (18510, -0.41 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -1.34 DPS) [world]; Deep Woodlands Cloak (19121, -1.69 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 99.1 spell_power points (12.15 DPS) | yes | Fireleaf Garb (240051, -0.86 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -4.72 DPS) [crafted]; Warlord's Silk Raiment (16535, -4.95 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 59.7 spell_power points (7.32 DPS) | yes | Fireleaf Wristwraps (240044, -2.79 DPS, sim-verified) [vendor]; Dryad's Wrist Bindings (19595, -3.62 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 73.0 spell_power points (8.95 DPS) | yes | Fireleaf Mitts (240049, -0.32 DPS, sim-verified) [vendor]; General's Silk Handguards (16540, -4.14 DPS) [vendor]; General's Silk Gauntlets (231599, -4.14 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 70.1 spell_power points (8.59 DPS) | yes | Knowledge of the Timbermaw (228190, -1.45 DPS) [vendor]; Fireleaf Waistguard (240045, -1.63 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -3.12 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 87.7 spell_power points (10.76 DPS) | yes | Fireleaf Pants (240047, -1.69 DPS, sim-verified) [vendor]; Sentinel's Silk Leggings (237815, -3.53 DPS) [vendor]; General's Silk Trousers (16534, -3.55 DPS) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 65.6 spell_power points (8.05 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (231597, -2.40 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (242.3 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.37 DPS) [vendor]; Naglering (11669, -12.48 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (242.3 DPS) | yes | Eye of Orgrimmar (12545, -0.99 DPS) [quest]; Maiden's Circle (13001, -0.99 DPS) [world_drop]; Naglering (11669, -7.77 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (242.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -12.81 DPS, sim-verified) [quest] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (242.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Talisman of Ascendance (22678, -6.18 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (242.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.23 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -12.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 631.8 spell_power points (77.50 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -4.19 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -13.02 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.43 DPS) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1021, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

