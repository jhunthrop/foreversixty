# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 28.3. Weights run: 1.0s. Verify run: 0.7s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.229 ± 0.006, crit=0.084 ± 0.002 per rating point (14 rating = 1%, 1.182 per %), hit=0.254 ± 0.002 per rating point (10 rating = 1%, 2.540 per %), spell_haste=0.494 ± 0.080, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.432 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Shadow Goggles (4373, -1.09 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.1 spell_power points (0.64 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.28 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.45 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.20 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.1 spell_power points (0.56 DPS) | yes | Green Woolen Vest (2582, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.20 DPS) [dungeon]; Gray Woolen Robe (2585, -0.73 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.1 spell_power points (0.10 DPS) | yes | Bright Bracers (3647, -0.02 DPS) [world_drop]; Repurposed Hair Band (281256, -0.06 DPS) [quest]; Windsong Bangles (263336, -0.69 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.21 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.41 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.45 DPS) | yes | Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.28 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.50 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.8 spell_power points (0.99 DPS) | yes | Filigreed Pristine Leggings (253937, -0.32 DPS) [crafted]; Silk-threaded Trousers (1929, -0.35 DPS) [dungeon]; Rumpled Kilt (274741, -0.53 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.72 DPS) | yes | Red Woolen Boots (4313, -0.36 DPS) [crafted]; Pristine Boots (253889, -0.39 DPS) [crafted]; Feather Padded Treads (285345, -0.65 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (0.50 DPS) | yes | Sludge-Stained Band (286535, -0.22 DPS) [world]; Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.46 DPS) | yes | Sludge-Stained Band (286535, -0.33 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.3 spell_power points (0.21 DPS) | yes | Lesser Staff of the Spire (1300, -0.08 DPS) [world]; Staff of Westfall (2042, -0.10 DPS) [quest]; Channeler's Staff (4437, -0.15 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 247.7 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.47 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 47.6. Weights run: 1.1s. Verify run: 0.7s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.318 ± 0.010, crit=0.084 ± 0.002 per rating point (14 rating = 1%, 1.181 per %), hit=0.244 ± 0.002 per rating point (10 rating = 1%, 2.438 per %), spell_haste=not significant (0.180 ± 0.123), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.645 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.30 DPS) | yes | Enchanter's Cowl (4322, -0.22 DPS) [crafted]; Silk Headband (7050, -0.24 DPS) [crafted]; Embalmed Shroud (7691, -0.36 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (1.05 DPS) | yes | Darkspear Warding Pendant (272075, -0.77 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.90 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.90 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.9 spell_power points (1.40 DPS) | yes | Death Speaker Mantle (6685, -0.29 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.36 DPS) [quest]; Invoker's Mantle (215365, -0.39 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.59 DPS) | yes | Repairman's Cape (9605, -0.09 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Prelacy Cape (7004, -0.12 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.1 spell_power points (1.55 DPS) | yes | Death Speaker Robes (6682, -0.31 DPS) [dungeon]; Tree Bark Jacket (1486, -0.36 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.46 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.07 DPS) | yes | Nightsky Wristbands (6407, -0.84 DPS) [world_drop]; Stonecloth Bindings (14416, -0.88 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.00 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.5 spell_power points (0.89 DPS) | yes | Serpent Gloves (5970, -0.06 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.18 DPS) [world]; Shilly Mitts (9609, -0.27 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.0 spell_power points (1.41 DPS) | yes | Belt of Arugal (6392, -0.24 DPS) [dungeon]; Invoker's Cord (215366, -0.40 DPS) [crafted]; Crimson Silk Belt (7055, -0.44 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.42 DPS) | yes | Abomination Skin Leggings (23173, -0.05 DPS) [dungeon]; Pristine Leggings (253987, -0.33 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.2 spell_power points (1.09 DPS) | yes | Acidic Walkers (9454, -0.20 DPS) [dungeon]; Nimbus Boots (6998, -0.38 DPS) [quest]; Spidersilk Boots (4320, -1.45 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.83 DPS) | yes | Minor Channeling Ring (1449, -0.16 DPS) [quest]; Electrocutioner Lagnut (9447, -0.47 DPS) [dungeon]; Sludge-Stained Band (286535, -0.47 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.71 DPS) | yes | Electrocutioner Lagnut (9447, -0.36 DPS) [dungeon]; Sludge-Stained Band (286535, -0.36 DPS) [world]; Minor Channeling Ring (1449, -0.88 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.07 DPS) | yes | Glimmering Staff (249392, -0.66 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.69 DPS) [world_drop]; Channeler's Staff (4437, -0.76 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.9 spell_power points (1.05 DPS) | yes | Dwarven Tome (279898, -0.42 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.58 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.58 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 283.6 spell_power points (33.57 DPS) | yes | Starfaller (13063, -0.32 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.97 DPS) [crafted]; Gravestone Scepter (7001, -4.57 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 253225113100011400-00000000000000000-0000000000000000000)

Set DPS (verified): 166.2. Weights run: 1.3s. Verify run: 0.8s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.555 ± 0.025, crit=0.169 ± 0.004 per rating point (14 rating = 1%, 2.368 per %), hit=0.303 ± 0.005 per rating point (10 rating = 1%, 3.025 per %), spell_haste=not significant (0.952 ± 0.375), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.916 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.78 DPS) | yes | Augural Shroud (2620, -1.22 DPS) [world]; Living Cowl (5608, -2.20 DPS) [world]; Enchanter's Cowl (4322, -2.60 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (2.84 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.31 DPS) [quest]; Triune Amulet (7722, -1.77 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.77 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.2 spell_power points (3.91 DPS) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.0 spell_power points (3.85 DPS) | yes | Guardian Cloak (5965, -1.44 DPS) [crafted]; Long Silken Cloak (4326, -1.85 DPS, sim-verified) [crafted]; Icy Cloak (4327, -1.93 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.3 spell_power points (6.97 DPS) | yes | Dreamweave Vest (10021, -0.64 DPS) [crafted]; Elemental Raiment (9434, -1.19 DPS) [world_drop]; Robe of Power (7054, -1.28 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.48 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.55 DPS) [quest]; Windchaser Cuffs (14429, -1.10 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.2 spell_power points (5.56 DPS) | yes | Black Mageweave Gloves (10003, -1.44 DPS) [crafted]; Red Mageweave Gloves (10018, -1.78 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.29 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.2 spell_power points (4.46 DPS) | yes | Deathmage Sash (10771, -0.25 DPS) [dungeon]; Star Belt (4329, -0.89 DPS) [crafted]; Gilded Cord (254037, -1.04 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.7 spell_power points (5.69 DPS) | yes | Crimson Silk Pantaloons (7062, -1.50 DPS) [crafted]; Abomination Skin Leggings (23173, -1.99 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.60 DPS) | yes | Gilded Slippers (254001, -2.23 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.00 DPS) [dungeon]; Spidersilk Boots (4320, -4.07 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (3.67 DPS) | yes | Ring of Forlorn Spirits (2043, -1.47 DPS) [quest]; Reedknot Ring (9622, -1.74 DPS) [quest]; Minor Channeling Ring (1449, -1.99 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.48 DPS) | yes | Ring of Forlorn Spirits (2043, -0.28 DPS) [quest]; Reedknot Ring (9622, -0.55 DPS) [quest]; Minor Channeling Ring (1449, -0.79 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (166.2 DPS) | yes | Scorn's Focal Dagger (23168, -3.03 DPS) [dungeon]; Windweaver Staff (7757, -3.21 DPS) [dungeon]; Gut Ripper (2164, -9.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 146.5 spell_power points (40.30 DPS) | yes | Nether Force Wand (11263, -2.51 DPS) [quest]; Icefury Wand (7514, -2.65 DPS) [quest]; Ragefire Wand (7513, -2.70 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 253225113100011531-03200000000000000-0000000000000000000)

Set DPS (verified): 277.7. Weights run: 1.3s. Verify run: 1.0s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.701 ± 0.039, crit=0.277 ± 0.006 per rating point (14 rating = 1%, 3.874 per %), hit=0.464 ± 0.008 per rating point (10 rating = 1%, 4.640 per %), spell_haste=not significant (1.338 ± 0.588), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.920 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 33.0 spell_power points (10.24 DPS) | yes | Dreamweave Circlet (10041, -1.55 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.87 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.09 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.2 spell_power points (6.26 DPS) | yes | Mindburst Medallion (11196, -3.10 DPS) [quest]; Horizon Choker (13085, -3.22 DPS) [world_drop]; Scorn's Icy Choker (23169, -5.99 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.6 spell_power points (7.94 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.31 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.51 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.2 spell_power points (5.64 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.90 DPS) [dungeon]; Runecloth Cloak (13860, -1.12 DPS) [crafted]; Big Voodoo Cloak (8216, -2.14 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.0 spell_power points (10.24 DPS) | yes | Robe of the Magi (1716, -2.11 DPS) [world_drop]; Runecloth Tunic (13857, -2.58 DPS) [crafted]; Dreamweave Vest (10021, -2.70 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.9 spell_power points (3.69 DPS) | yes | Aristocratic Cuffs (12546, -0.43 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.90 DPS) [quest]; Bloodband Bracers (11469, -3.50 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 26.5 spell_power points (8.20 DPS) | yes | Raider Handwraps (272098, -1.25 DPS) [vendor]; Dreamweave Gloves (10019, -1.75 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.22 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (277.7 DPS) | yes | Dawnspire Cord (12466, -0.26 DPS) [dungeon]; Deathmage Sash (10771, -0.82 DPS) [dungeon]; Satyrmane Sash (17755, -5.90 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 30.0 spell_power points (9.30 DPS) | yes | Red Mageweave Pants (10009, -2.36 DPS) [crafted]; Wizardweave Leggings (14132, -3.41 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -5.49 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.44 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.56 DPS) [vendor]; Gilded Sandals (254107, -2.07 DPS) [crafted]; Black Mageweave Boots (10026, -2.51 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (4.40 DPS) | yes | Band of the Unicorn (7553, -0.37 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.68 DPS) [rep]; Brainlash (6440, -1.14 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.9 spell_power points (4.31 DPS) | yes | Band of the Unicorn (7553, -0.28 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.59 DPS) [rep]; Brainlash (6440, -1.05 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -7.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 169.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.27 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 253225113100011531-03202300000000000-0050000000000000000)

Set DPS (verified): 460.5. Weights run: 1.4s. Verify run: 1.0s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.815 ± 0.048, crit=0.443 ± 0.010 per rating point (14 rating = 1%, 6.203 per %), hit=0.718 ± 0.010 per rating point (10 rating = 1%, 7.181 per %), spell_haste=3.805 ± 0.776, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.914 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.6 spell_power points (15.39 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.84 DPS) [pvp]; Crimson Felt Hat (18727, -3.58 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (460.5 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.91 DPS) [quest]; Chains of the Lich (23125, -1.16 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.4 spell_power points (14.69 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.65 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.76 DPS) [crafted]; Darkspear Shoulderpads (272103, -8.49 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.7 spell_power points (9.61 DPS) | yes | Crystalline Threaded Cape (20697, -2.08 DPS) [world]; Hide of the Wild (18510, -2.44 DPS) [crafted]; Spritecaster Cape (11623, -3.50 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.0 spell_power points (18.10 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.95 DPS) [pvp]; Robe of Everlasting Night (18385, -4.11 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Tunic (227108, -4.83 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.5 spell_power points (9.22 DPS) | yes | Sublime Wristguards (18497, -2.71 DPS) [dungeon]; Runecloth Cuffs (254123, -3.03 DPS) [crafted]; Marshal's Silk Bracers (16438, -4.22 DPS) [pvp] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 31.1 spell_power points (10.05 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.1 spell_power points (16.21 DPS) | yes | Belt of the Archmage (18405, -3.51 DPS) [crafted]; Magician's Cord (272393, -3.74 DPS) [vendor]; Stormpike Cloth Girdle (19094, -7.75 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.6 spell_power points (17.00 DPS) | yes | Marshal's Silk Leggings (231605, -0.02 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.72 DPS) [pvp]; Skyshroud Leggings (13170, -3.89 DPS) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.0 spell_power points (11.01 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.97 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (460.5 DPS) | yes | Songstone of Ironforge (12543, -2.35 DPS) [quest]; Maiden's Circle (13001, -2.35 DPS) [world_drop]; Naglering (11669, -11.45 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (460.5 DPS) | yes | Songstone of Ironforge (12543, -0.97 DPS) [quest]; Maiden's Circle (13001, -0.97 DPS) [world_drop]; Naglering (11669, -10.89 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (460.5 DPS) | yes | Weakness Analyzer (272438, -2.26 DPS) [vendor]; Serenity Field (272439, -4.85 DPS) [vendor]; Burst of Knowledge (11832, -5.50 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (460.5 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -2.80 DPS, sim-verified) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (460.5 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.31 DPS) [world]; Teebu's Blazing Longsword (1728, -14.53 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 243.3 spell_power points (78.70 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.29 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.46 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.38 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 26.4. Weights run: 1.0s. Verify run: 0.7s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.229 ± 0.006, crit=0.084 ± 0.002 per rating point (14 rating = 1%, 1.182 per %), hit=0.254 ± 0.002 per rating point (10 rating = 1%, 2.540 per %), spell_haste=0.494 ± 0.080, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.432 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Shadow Goggles (4373, -0.80 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.1 spell_power points (0.64 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.28 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Black Whelp Cloak (7283, -0.34 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.1 spell_power points (0.56 DPS) | yes | Green Woolen Vest (2582, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.20 DPS) [dungeon]; Gray Woolen Robe (2585, -0.78 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.5 spell_power points (0.13 DPS) | yes | Tabitha's Cuffs (251486, -0.01 DPS) [quest]; Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.21 DPS) [crafted]; Apothecary Gloves (10919, -0.27 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.45 DPS) | yes | Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.28 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.43 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.8 spell_power points (0.99 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.35 DPS) [dungeon]; Rumpled Kilt (274741, -0.53 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.72 DPS) | yes | Red Woolen Boots (4313, -0.36 DPS) [crafted]; Pristine Boots (253889, -0.39 DPS) [crafted]; Feather Padded Treads (285345, -0.43 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Loop of Sacrifice (281673, -0.35 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.27 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon]; Loop of Sacrifice (281673, -0.17 DPS) [quest]; Volcanic Rock Ring (12053, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.3 spell_power points (0.21 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.08 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 247.7 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.45 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 44.7. Weights run: 1.1s. Verify run: 0.8s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.318 ± 0.010, crit=0.084 ± 0.002 per rating point (14 rating = 1%, 1.181 per %), hit=0.244 ± 0.002 per rating point (10 rating = 1%, 2.438 per %), spell_haste=not significant (0.180 ± 0.123), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.645 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.30 DPS) | yes | Enchanter's Cowl (4322, -0.22 DPS) [crafted]; Silk Headband (7050, -0.24 DPS) [crafted]; Embalmed Shroud (7691, -0.36 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (1.05 DPS) | yes | Crystal Starfire Medallion (5003, -0.90 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.90 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.18 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.9 spell_power points (1.40 DPS) | yes | Fairywing Mantle (9536, -0.36 DPS) [quest]; Invoker's Mantle (215365, -0.39 DPS) [crafted]; Death Speaker Mantle (6685, -0.60 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.59 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Battle Healer's Cloak (19529, -0.12 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.1 spell_power points (1.55 DPS) | yes | Tree Bark Jacket (1486, -0.30 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.31 DPS) [dungeon]; Pristine Gown (253961, -0.46 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.07 DPS) | yes | Nightsky Wristbands (6407, -0.84 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.84 DPS) [quest]; Glowing Magical Bracelets (13106, -0.92 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.6 spell_power points (0.90 DPS) | yes | Gnoll Casting Gloves (892, -0.19 DPS) [world]; Truefaith Gloves (7049, -0.19 DPS) [crafted]; Serpent Gloves (5970, -0.32 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.0 spell_power points (1.41 DPS) | yes | Belt of Arugal (6392, -0.24 DPS) [dungeon]; Invoker's Cord (215366, -0.40 DPS) [crafted]; Warsong Sash (16975, -0.46 DPS, sim-verified) [quest] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.42 DPS) | yes | Abomination Skin Leggings (23173, -0.05 DPS) [dungeon]; Pristine Leggings (253987, -0.33 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.2 spell_power points (1.09 DPS) | yes | Acidic Walkers (9454, -0.20 DPS) [dungeon]; Boots of the Enchanter (4325, -0.50 DPS) [crafted]; Spidersilk Boots (4320, -1.38 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.83 DPS) | yes | Electrocutioner Lagnut (9447, -0.47 DPS) [dungeon]; Sludge-Stained Band (286535, -0.47 DPS) [world]; Black Widow Band (6199, -0.57 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.71 DPS) | yes | Electrocutioner Lagnut (9447, -0.36 DPS) [dungeon]; Black Widow Band (6199, -0.45 DPS) [world]; Sludge-Stained Band (286535, -1.16 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.07 DPS) | yes | Twisted Chanter's Staff (890, -0.69 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.69 DPS) [quest]; Glimmering Staff (249392, -0.81 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.9 spell_power points (1.05 DPS) | yes | Orb of Souls (249395, -0.58 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -0.65 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -0.67 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 283.6 spell_power points (33.57 DPS) | yes | Starfaller (13063, -0.35 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.97 DPS) [crafted]; Gravestone Scepter (7001, -4.57 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 253225113100011400-00000000000000000-0000000000000000000)

Set DPS (verified): 163.0. Weights run: 1.3s. Verify run: 0.8s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.555 ± 0.025, crit=0.169 ± 0.004 per rating point (14 rating = 1%, 2.368 per %), hit=0.303 ± 0.005 per rating point (10 rating = 1%, 3.025 per %), spell_haste=not significant (0.952 ± 0.375), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.916 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.78 DPS) | yes | Living Cowl (5608, -2.20 DPS) [world]; Enchanter's Cowl (4322, -2.60 DPS) [crafted]; Augural Shroud (2620, -3.15 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (2.84 DPS) | yes | Triune Amulet (7722, -1.77 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.77 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.65 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.2 spell_power points (3.91 DPS) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 14.0 spell_power points (3.85 DPS) | yes | Guardian Cloak (5965, -1.44 DPS) [crafted]; Icy Cloak (4327, -1.93 DPS) [crafted]; Long Silken Cloak (4326, -3.68 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.3 spell_power points (6.97 DPS) | yes | Dreamweave Vest (10021, -0.64 DPS) [crafted]; Elemental Raiment (9434, -1.19 DPS) [world_drop]; Robe of Power (7054, -1.28 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.48 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.55 DPS) [quest]; Radiant Silver Bracers (4545, -1.59 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.2 spell_power points (5.56 DPS) | yes | Black Mageweave Gloves (10003, -1.44 DPS) [crafted]; Gilded Handwraps (254021, -2.29 DPS) [crafted]; Red Mageweave Gloves (10018, -2.29 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.2 spell_power points (4.46 DPS) | yes | Star Belt (4329, -0.89 DPS) [crafted]; Gilded Cord (254037, -1.04 DPS) [crafted]; Deathmage Sash (10771, -1.72 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.7 spell_power points (5.69 DPS) | yes | Abomination Skin Leggings (23173, -1.99 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.14 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.60 DPS) | yes | Gilded Slippers (254001, -3.76 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.00 DPS) [dungeon]; Spidersilk Boots (4320, -4.07 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (3.67 DPS) | yes | Reedknot Ring (9622, -1.74 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.02 DPS) [vendor]; Black Widow Band (6199, -2.60 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.48 DPS) | yes | Sea Giant's Toe Ring (274746, -0.83 DPS) [vendor]; Black Widow Band (6199, -1.41 DPS) [world]; Reedknot Ring (9622, -1.84 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (163.0 DPS) | yes | Scorn's Focal Dagger (23168, -3.03 DPS) [dungeon]; Windweaver Staff (7757, -3.21 DPS) [dungeon]; Gut Ripper (2164, -9.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 146.5 spell_power points (40.30 DPS) | yes | Nether Force Wand (11263, -2.51 DPS) [quest]; Icefury Wand (7514, -2.65 DPS) [quest]; Ragefire Wand (7513, -2.70 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253225113100011531-03200000000000000-0000000000000000000)

Set DPS (verified): 270.9. Weights run: 1.3s. Verify run: 1.0s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.701 ± 0.039, crit=0.277 ± 0.006 per rating point (14 rating = 1%, 3.874 per %), hit=0.464 ± 0.008 per rating point (10 rating = 1%, 4.640 per %), spell_haste=not significant (1.338 ± 0.588), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.920 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 33.0 spell_power points (10.24 DPS) | yes | Dreamweave Circlet (10041, -1.55 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.87 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.09 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 20.2 spell_power points (6.26 DPS) | yes | Mindburst Medallion (11196, -3.10 DPS) [quest]; Horizon Choker (13085, -3.22 DPS) [world_drop]; Scorn's Icy Choker (23169, -5.97 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.6 spell_power points (7.94 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.31 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.51 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.3 spell_power points (5.68 DPS) | yes | Spritecaster Cape (11623, -0.03 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.93 DPS) [dungeon]; Runecloth Cloak (13860, -1.15 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.0 spell_power points (10.24 DPS) | yes | Robe of the Magi (1716, -2.11 DPS) [world_drop]; Runecloth Tunic (13857, -2.58 DPS) [crafted]; Dreamweave Vest (10021, -2.70 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.9 spell_power points (3.69 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -3.71 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 26.5 spell_power points (8.20 DPS) | yes | Dreamweave Gloves (10019, -1.75 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.22 DPS) [vendor]; Raider Handwraps (272098, -2.79 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (270.9 DPS) | yes | Dawnspire Cord (12466, -0.26 DPS) [dungeon]; Deathmage Sash (10771, -0.82 DPS) [dungeon]; Satyrmane Sash (17755, -3.77 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 30.0 spell_power points (9.30 DPS) | yes | Red Mageweave Pants (10009, -2.36 DPS) [crafted]; Wizardweave Leggings (14132, -3.41 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -6.05 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.44 DPS) | yes | Gilded Sandals (254107, -2.07 DPS) [crafted]; Black Mageweave Boots (10026, -2.51 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.21 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (4.40 DPS) | yes | Band of the Unicorn (7553, -0.37 DPS) [world_drop]; Advisor's Ring (19519, -0.68 DPS) [rep]; Brainlash (6440, -1.14 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.9 spell_power points (4.31 DPS) | yes | Band of the Unicorn (7553, -0.28 DPS) [world_drop]; Advisor's Ring (19519, -0.59 DPS) [rep]; Brainlash (6440, -1.05 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, +0.00 DPS) [world]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -7.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 169.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.27 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 253225113100011531-03202300000000000-0050000000000000000)

Set DPS (verified): 456.7. Weights run: 1.4s. Verify run: 0.9s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.815 ± 0.048, crit=0.443 ± 0.010 per rating point (14 rating = 1%, 6.203 per %), hit=0.718 ± 0.010 per rating point (10 rating = 1%, 7.181 per %), spell_haste=3.805 ± 0.776, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.914 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.6 spell_power points (15.39 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.84 DPS) [pvp]; Crimson Felt Hat (18727, -3.58 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (456.7 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.91 DPS) [quest]; Chains of the Lich (23125, -1.16 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.4 spell_power points (14.69 DPS) | yes | Warlord's Silk Amice (231594, -2.65 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.76 DPS) [crafted]; Darkspear Shoulderpads (272103, -7.74 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.7 spell_power points (9.61 DPS) | yes | Crystalline Threaded Cape (20697, -2.08 DPS) [world]; Hide of the Wild (18510, -2.44 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.35 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.0 spell_power points (18.10 DPS) | yes | Warlord's Silk Raiment (231596, -0.95 DPS) [pvp]; Robe of Everlasting Night (18385, -4.79 DPS, sim-verified) [dungeon]; Legionnaire's Silk Tunic (227106, -4.83 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.5 spell_power points (9.22 DPS) | yes | Sublime Wristguards (18497, -2.71 DPS) [dungeon]; Runecloth Cuffs (254123, -3.03 DPS) [crafted]; General's Silk Cuffs (16538, -4.22 DPS) [pvp] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 31.1 spell_power points (10.05 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.1 spell_power points (16.21 DPS) | yes | Belt of the Archmage (18405, -3.51 DPS) [crafted]; Magician's Cord (272393, -3.74 DPS) [vendor]; Frostwolf Cloth Belt (19090, -7.75 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.6 spell_power points (17.00 DPS) | yes | General's Silk Trousers (231595, -0.02 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -3.72 DPS) [pvp]; Outrider's Silk Leggings (22747, -7.02 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.0 spell_power points (11.01 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.97 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (456.7 DPS) | yes | Eye of Orgrimmar (12545, -2.35 DPS) [quest]; Maiden's Circle (13001, -2.35 DPS) [world_drop]; Naglering (11669, -11.40 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (456.7 DPS) | yes | Eye of Orgrimmar (12545, -0.97 DPS) [quest]; Maiden's Circle (13001, -0.97 DPS) [world_drop]; Naglering (11669, -11.24 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (456.7 DPS) | yes | Weakness Analyzer (272438, -2.26 DPS) [vendor]; Serenity Field (272439, -4.85 DPS) [vendor]; Burst of Knowledge (11832, -5.50 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (456.7 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (456.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.80 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -20.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 243.3 spell_power points (78.70 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.29 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.46 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.38 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

