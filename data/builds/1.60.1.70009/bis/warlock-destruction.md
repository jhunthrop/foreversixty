# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 43.3. Weights run: 2.5s. Verify run: 1.3s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.087, intellect=0.559 ± 0.015, crit=0.051 ± 0.002 per rating point (14 rating = 1%, 0.720 per %), hit=0.161 ± 0.001 per rating point (10 rating = 1%, 1.608 per %), spell_haste=0.772 ± 0.077, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.669 ± 0.087, fire_power=0.336 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -3.28 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.0 spell_power points (1.33 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.80 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.50 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.8 spell_power points (1.03 DPS) | yes | Green Woolen Robe (6243, -0.41 DPS) [crafted]; Green Woolen Vest (2582, -0.50 DPS) [crafted]; Gray Woolen Robe (2585, -1.63 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 2.8 spell_power points (0.37 DPS) | yes | Bright Bracers (3647, -0.22 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -0.22 DPS) [world_drop]; Repurposed Hair Band (281256, -0.22 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.93 DPS) | yes | Pristine Gloves (253913, -0.18 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.48 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.2 spell_power points (0.83 DPS) | yes | Keller's Girdle (2911, -0.23 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.04 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (43.3 DPS) | yes | Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Rumpled Kilt (274741, -0.58 DPS) [vendor]; Abomination Skin Leggings (23173, -0.76 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.2 spell_power points (1.22 DPS) | yes | Pristine Boots (253889, -0.60 DPS) [crafted]; Red Woolen Boots (4313, -0.69 DPS) [crafted]; Feather Padded Treads (285345, -0.71 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.1 spell_power points (0.81 DPS) | yes | Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon]; Sludge-Stained Band (286535, -0.41 DPS) [world]; Volcanic Rock Ring (12053, -0.59 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.66 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.29 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 5.6 spell_power points (0.74 DPS) | yes | Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.30 DPS) [world]; Staff of Westfall (2042, -0.37 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 171.5 spell_power points (22.71 DPS) | yes | Skycaller (12984, -2.39 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 66.2. Weights run: 2.5s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.083, intellect=0.236 ± 0.005, crit=0.052 ± 0.002 per rating point (14 rating = 1%, 0.730 per %), hit=0.102 ± 0.001 per rating point (10 rating = 1%, 1.022 per %), spell_haste=0.650 ± 0.080, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.833 ± 0.083, fire_power=0.167 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.80 DPS) | yes | Silk Headband (7050, -0.57 DPS, sim-verified) [crafted]; Enchanter's Cowl (4322, -0.67 DPS) [crafted]; Embalmed Shroud (7691, -0.76 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 spell_power points (2.14 DPS) | yes | Darkspear Warding Pendant (272075, -1.83 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.90 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.90 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 spell_power points (2.83 DPS) | yes | Invoker's Mantle (215365, -0.75 DPS) [crafted]; Fairywing Mantle (9536, -0.76 DPS) [quest]; Death Speaker Mantle (6685, -0.79 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.27 DPS) | yes | Prelacy Cape (7004, -0.25 DPS) [quest]; Caretaker's Cape (19533, -0.25 DPS) [rep]; Heavy Woolen Cloak (4311, -0.65 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.31 DPS) | yes | Green Silk Armor (7065, -0.24 DPS) [crafted]; Death Speaker Robes (6682, -0.86 DPS) [dungeon]; Pristine Gown (253961, -1.11 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.29 DPS) | yes | Glowing Magical Bracelets (13106, -1.77 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.93 DPS) [world_drop]; Stonecloth Bindings (14416, -1.99 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.78 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Town Clerk's Mittens (270029, -0.10 DPS) [quest]; Gnoll Casting Gloves (892, -0.25 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.7 spell_power points (2.98 DPS) | yes | Belt of Arugal (6392, -0.53 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.90 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.94 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.05 DPS) | yes | Abomination Skin Leggings (23173, -0.28 DPS) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.17 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.7 spell_power points (2.20 DPS) | yes | Acidic Walkers (9454, -0.45 DPS) [dungeon]; Nimbus Boots (6998, -0.67 DPS) [quest]; Spidersilk Boots (4320, -1.95 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.78 DPS) | yes | Minor Channeling Ring (1449, -0.39 DPS) [quest]; Electrocutioner Lagnut (9447, -1.02 DPS) [dungeon]; Sludge-Stained Band (286535, -1.02 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.53 DPS) | yes | Electrocutioner Lagnut (9447, -0.76 DPS) [dungeon]; Sludge-Stained Band (286535, -0.76 DPS) [world]; Minor Channeling Ring (1449, -1.65 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.29 DPS) | yes | Glimmering Staff (249392, -1.15 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.69 DPS) [world_drop]; Channeler's Staff (4437, -1.81 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.4 spell_power points (2.14 DPS) | yes | Eye of Paleth (2943, -1.12 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -1.12 DPS) [world]; Dwarven Tome (279898, -1.21 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 133.6 spell_power points (33.97 DPS) | yes | Starfaller (13063, -0.83 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.70 DPS) [crafted]; Gravestone Scepter (7001, -4.97 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 117.3. Weights run: 2.1s. Verify run: 1.1s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.190, intellect=0.439 ± 0.012, crit=0.084 ± 0.003 per rating point (14 rating = 1%, 1.177 per %), hit=0.188 ± 0.002 per rating point (10 rating = 1%, 1.876 per %), spell_haste=0.495 ± 0.123, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.449 ± 0.190), fire_power=0.551 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.38 DPS) | yes | Augural Shroud (2620, -1.23 DPS, sim-verified) [world]; Living Cowl (5608, -2.05 DPS) [world]; Holy Shroud (2721, -2.56 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (2.47 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.68 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.68 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.68 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (3.32 DPS) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Inquisitor's Shawl (19507, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.40 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.0 spell_power points (3.32 DPS) | yes | Long Silken Cloak (4326, -0.52 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -1.22 DPS) [crafted]; Icy Cloak (4327, -1.53 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.6 spell_power points (6.32 DPS) | yes | Dreamweave Vest (10021, -0.69 DPS) [crafted]; Elemental Raiment (9434, -0.93 DPS) [world_drop]; Robe of Power (7054, -1.38 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.31 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.51 DPS) [quest]; Earthen Silk Cuffs (254019, -1.28 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (5.07 DPS) | yes | Black Mageweave Gloves (10003, -1.22 DPS) [crafted]; Red Mageweave Gloves (10018, -1.28 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.23 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.8 spell_power points (4.04 DPS) | yes | Deathmage Sash (10771, -0.60 DPS, sim-verified) [dungeon]; Star Belt (4329, -0.71 DPS) [crafted]; Gilded Cord (254037, -1.09 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.3 spell_power points (4.94 DPS) | yes | Crimson Silk Pantaloons (7062, -1.04 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.73 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.86 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.15 DPS) | yes | Gilded Slippers (254001, -1.67 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.91 DPS) [crafted]; Acidic Walkers (9454, -3.97 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.6 spell_power points (3.24 DPS) | yes | Ring of Forlorn Spirits (2043, -1.19 DPS) [quest]; Reedknot Ring (9622, -1.45 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.70 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.31 DPS) | yes | Ring of Forlorn Spirits (2043, -0.26 DPS) [quest]; Reedknot Ring (9622, -0.51 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.77 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (117.3 DPS) | yes | Scorn's Focal Dagger (23168, -2.82 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.41 DPS) [quest]; Gut Ripper (2164, -5.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 156.2 spell_power points (40.06 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.47 DPS) [dungeon]; Twisted Nether Wand (249144, -4.52 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 149.4. Weights run: 2.0s. Verify run: 1.3s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.735, intellect=1.960 ± 0.046, crit=0.327 ± 0.012 per rating point (14 rating = 1%, 4.579 per %), hit=0.721 ± 0.006 per rating point (10 rating = 1%, 7.206 per %), spell_haste=not significant (-0.632 ± 0.764), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.480 ± 0.735), fire_power=1.502 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 58.2 spell_power points (5.16 DPS) | yes | Soulcatcher Halo (10630, -0.82 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Hat (220889, -1.43 DPS) [vendor]; Chief Architect's Monocle (11839, -5.43 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 27.4 spell_power points (2.43 DPS) | yes | Gemshard Heart (17707, -0.69 DPS) [dungeon]; Scorn's Icy Choker (23169, -0.77 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.82 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 48.3 spell_power points (4.28 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.05 DPS) [crafted]; Inquisitor's Shawl (19507, -1.40 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -0.08 DPS) [dungeon]; Runecloth Cloak (13860, -0.17 DPS) [crafted]; Darkspear Raider's Cloak (272076, -2.72 DPS, sim-verified) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 58.2 spell_power points (5.16 DPS) | yes | Runecloth Robe (13858, -1.23 DPS) [crafted]; Hibernal Robe (8113, -1.68 DPS) [world_drop]; Robes of Insight (940, -5.36 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 29.4 spell_power points (2.61 DPS) | yes | Forgotten Wraps (9433, -0.52 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.57 DPS, sim-verified) [quest]; Bloodband Bracers (11469, -0.60 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 50.1 spell_power points (4.44 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -0.96 DPS, sim-verified) [vendor]; Red Mageweave Gloves (10018, -1.73 DPS) [crafted]; Runecloth Gloves (13863, -1.82 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 43.2 spell_power points (3.83 DPS) | yes | Deathmage Sash (10771, -0.61 DPS) [dungeon]; Ban'thok Sash (11662, -0.79 DPS) [dungeon]; Satyrmane Sash (17755, -0.85 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 42.6 spell_power points (3.78 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.38 DPS) [dungeon]; Red Mageweave Pants (10009, -0.45 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -4.52 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 32.8 spell_power points (2.91 DPS) | yes | Gilded Sandals (254107, -0.37 DPS) [crafted]; Coldstone Slippers (18697, -0.48 DPS) [dungeon]; Southsea Mojo Boots (20641, -1.57 DPS, sim-verified) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.4 spell_power points (2.61 DPS) | yes | Philanthropist's Ring (281635, -0.68 DPS) [quest]; Mindseye Circle (10634, -0.84 DPS, sim-verified) [dungeon]; Woodseed Hoop (17768, -1.04 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Philanthropist's Ring (281635, -0.09 DPS) [quest]; Woodseed Hoop (17768, -0.45 DPS) [quest]; Mindseye Circle (10634, -2.68 DPS, sim-verified) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.13 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.04 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Soul Harvester (20536, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -3.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.02 DPS) [world_drop]; Flaming Incinerator (9483, -3.22 DPS) [dungeon]; Pyric Caduceus (11748, -4.07 DPS, sim-verified) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 278.6. Weights run: 8.2s. Verify run: 1.2s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± 0.378), intellect=not significant (-2.313 ± 0.023), crit=not significant (-0.457 ± 0.008) per rating point (14 rating = 1%, -6.402 per %), hit=not significant (-0.990 ± 0.004) per rating point (10 rating = 1%, -9.899 per %), spell_haste=not significant (-2.393 ± 0.292), spell_penetration=not significant (-0.000 ± 0.000), shadow_power=not significant (2.392 ± 0.378), fire_power=not significant (-1.419 ± 0.002)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | 72.2 spell_power points | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; The Postmaster's Band (13390, -0.82 DPS) [dungeon]; Magister's Crown (16686, -2.69 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 40.5 spell_power points | yes | Beads of Ogre Mojo (22149, -0.35 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -0.46 DPS) [world_drop]; Lady Maye's Pendant (14558, -2.11 DPS, sim-verified) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.24 DPS) [pvp]; Burial Shawl (18681, -0.85 DPS) [dungeon]; Darkspear Shoulderpads (272103, -6.09 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 38.9 spell_power points | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Royal Tribunal Cloak (13376, -0.67 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -0.67 DPS) [vendor] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | sim-verified (+8.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Captain's Dreadweave Tunic (227096, +0.00 DPS) [pvp]; Field Marshal's Dreadweave Robe (231582, +0.00 DPS) [pvp]; Magister's Robes (16688, -8.38 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 37.7 spell_power points | yes | Marshal's Dreadweave Cuffs (17582, -0.04 DPS) [pvp]; Sublime Wristguards (18497, -0.54 DPS) [dungeon]; Magiskull Cuffs (13107, -5.28 DPS, sim-verified) [world_drop] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 60.9 spell_power points | yes | Deathmist Wraps (226911, -1.35 DPS) [quest]; Mooncloth Gloves (18409, -1.64 DPS) [crafted]; Marshal's Dreadweave Gloves (231586, -1.70 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 75.3 spell_power points | yes | Belt of the Archmage (18405, -1.15 DPS, sim-verified) [crafted]; Magister's Belt (16685, -2.41 DPS) [dungeon]; Marshal's Dreadweave Sash (17585, -2.68 DPS) [pvp] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Deathmist Leggings (226910, -0.29 DPS) [quest]; Knight-Captain's Dreadweave Leggings (17567, -0.49 DPS) [pvp] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 49.4 spell_power points | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.09 DPS) [pvp]; Omnicast Boots (11822, -0.52 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -0.79 DPS) [vendor]; Seal of Rivendare (13345, -0.91 DPS) [dungeon]; Naglering (11669, -6.61 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -0.01 DPS) [vendor]; Seal of Rivendare (13345, -0.12 DPS) [dungeon]; Naglering (11669, -5.60 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+10.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -0.53 DPS) [quest]; Weakness Analyzer (272438, -0.62 DPS) [vendor]; Serenity Field (272439, -1.33 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -1.28 DPS, sim-verified) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -0.98 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -7.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+8.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Oblivion's Touch (18761, -0.41 DPS) [dungeon]; Sparkling Crystal Wand (20672, -0.49 DPS) [world]; Torch of Light (279246, -8.51 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 41.8. Weights run: 2.5s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.087, intellect=0.559 ± 0.015, crit=0.051 ± 0.002 per rating point (14 rating = 1%, 0.720 per %), hit=0.161 ± 0.001 per rating point (10 rating = 1%, 1.608 per %), spell_haste=0.772 ± 0.077, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.669 ± 0.087, fire_power=0.336 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -3.00 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.0 spell_power points (1.33 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.37 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.80 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.44 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.8 spell_power points (1.03 DPS) | yes | Green Woolen Robe (6243, -0.41 DPS) [crafted]; Green Woolen Vest (2582, -0.50 DPS) [crafted]; Gray Woolen Robe (2585, -1.44 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.4 spell_power points (0.44 DPS) | yes | Mindthrust Bracers (1974, -0.07 DPS) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Bright Bracers (3647, -0.15 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.93 DPS) | yes | Pristine Gloves (253913, -0.18 DPS) [crafted]; Gnoll Casting Gloves (892, -0.26 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.40 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.2 spell_power points (0.83 DPS) | yes | Keller's Girdle (2911, -0.23 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.94 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Rumpled Kilt (274741, -0.58 DPS) [vendor]; Abomination Skin Leggings (23173, -0.82 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.2 spell_power points (1.22 DPS) | yes | Pristine Boots (253889, -0.60 DPS) [crafted]; Red Woolen Boots (4313, -0.69 DPS) [crafted]; Feather Padded Treads (285345, -0.72 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.29 DPS) [quest]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.44 DPS, sim-verified) [dungeon] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Loop of Sacrifice (281673, -0.03 DPS) [quest]; Volcanic Rock Ring (12053, -0.18 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.69 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 5.6 spell_power points (0.74 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.30 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 171.5 spell_power points (22.71 DPS) | yes | Skycaller (12984, -2.23 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Sizzle Stick (8071, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 64.7. Weights run: 2.5s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.083, intellect=0.236 ± 0.005, crit=0.052 ± 0.002 per rating point (14 rating = 1%, 0.730 per %), hit=0.102 ± 0.001 per rating point (10 rating = 1%, 1.022 per %), spell_haste=0.650 ± 0.080, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.833 ± 0.083, fire_power=0.167 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.80 DPS) | yes | Silk Headband (7050, -0.62 DPS, sim-verified) [crafted]; Enchanter's Cowl (4322, -0.67 DPS) [crafted]; Embalmed Shroud (7691, -0.76 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 spell_power points (2.14 DPS) | yes | Crystal Starfire Medallion (5003, -1.90 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.90 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.90 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 spell_power points (2.83 DPS) | yes | Death Speaker Mantle (6685, -0.60 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.75 DPS) [crafted]; Fairywing Mantle (9536, -0.76 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.27 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.25 DPS) [crafted]; Battle Healer's Cloak (19529, -0.25 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.31 DPS) | yes | Green Silk Armor (7065, -0.24 DPS) [crafted]; Death Speaker Robes (6682, -0.86 DPS) [dungeon]; Pristine Gown (253961, -1.11 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.29 DPS) | yes | Glowing Magical Bracelets (13106, -1.76 DPS, sim-verified) [world_drop]; Owlbeard Bracers (16981, -1.91 DPS) [quest]; Nightsky Wristbands (6407, -1.93 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.2 spell_power points (1.83 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.30 DPS) [world]; Truefaith Gloves (7049, -0.37 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.7 spell_power points (2.98 DPS) | yes | Warsong Sash (16975, -0.18 DPS) [quest]; Belt of Arugal (6392, -0.51 DPS) [dungeon]; Invoker's Cord (215366, -0.90 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.05 DPS) | yes | Abomination Skin Leggings (23173, -0.28 DPS) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.17 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.7 spell_power points (2.20 DPS) | yes | Acidic Walkers (9454, -0.45 DPS) [dungeon]; Boots of the Enchanter (4325, -0.93 DPS) [crafted]; Spidersilk Boots (4320, -1.85 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.78 DPS) | yes | Electrocutioner Lagnut (9447, -1.02 DPS) [dungeon]; Sludge-Stained Band (286535, -1.02 DPS) [world]; Sacred Band (6669, -1.27 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.53 DPS) | yes | Electrocutioner Lagnut (9447, -0.76 DPS) [dungeon]; Sacred Band (6669, -1.02 DPS) [quest]; Sludge-Stained Band (286535, -2.34 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.29 DPS) | yes | Twisted Chanter's Staff (890, -1.69 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.69 DPS) [quest]; Glimmering Staff (249392, -1.93 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.4 spell_power points (2.14 DPS) | yes | Orb of Souls (249395, -1.12 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -1.27 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -1.39 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 133.6 spell_power points (33.97 DPS) | yes | Starfaller (13063, -0.58 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.70 DPS) [crafted]; Gravestone Scepter (7001, -4.97 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 115.9. Weights run: 2.1s. Verify run: 1.1s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.190, intellect=0.439 ± 0.012, crit=0.084 ± 0.003 per rating point (14 rating = 1%, 1.177 per %), hit=0.188 ± 0.002 per rating point (10 rating = 1%, 1.876 per %), spell_haste=0.495 ± 0.123, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.449 ± 0.190), fire_power=0.551 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.38 DPS) | yes | Augural Shroud (2620, -1.58 DPS, sim-verified) [world]; Living Cowl (5608, -2.05 DPS) [world]; Holy Shroud (2721, -2.56 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (2.47 DPS) | yes | Triune Amulet (7722, -1.68 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.68 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.95 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 spell_power points (3.32 DPS) | yes | Green Silken Shoulders (7057, -0.03 DPS) [crafted]; Inquisitor's Shawl (19507, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.40 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.0 spell_power points (3.32 DPS) | yes | Guardian Cloak (5965, -1.22 DPS) [crafted]; Icy Cloak (4327, -1.53 DPS) [crafted]; Long Silken Cloak (4326, -1.60 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.6 spell_power points (6.32 DPS) | yes | Dreamweave Vest (10021, -0.73 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.93 DPS) [world_drop]; Robe of Power (7054, -1.38 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.31 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.51 DPS) [quest]; Radiant Silver Bracers (4545, -1.22 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (5.07 DPS) | yes | Black Mageweave Gloves (10003, -1.22 DPS) [crafted]; Red Mageweave Gloves (10018, -1.28 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.23 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.8 spell_power points (4.04 DPS) | yes | Star Belt (4329, -0.71 DPS) [crafted]; Gilded Cord (254037, -1.09 DPS) [crafted]; Deathmage Sash (10771, -1.52 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.3 spell_power points (4.94 DPS) | yes | Crimson Silk Pantaloons (7062, -1.64 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.73 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.86 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.15 DPS) | yes | Gilded Slippers (254001, -1.77 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.91 DPS) [crafted]; Acidic Walkers (9454, -3.97 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.6 spell_power points (3.24 DPS) | yes | Reedknot Ring (9622, -1.45 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.70 DPS) [vendor]; Black Widow Band (6199, -2.45 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.31 DPS) | yes | Sea Giant's Toe Ring (274746, -0.77 DPS) [vendor]; Reedknot Ring (9622, -1.39 DPS, sim-verified) [quest]; Black Widow Band (6199, -1.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (115.9 DPS) | yes | Scorn's Focal Dagger (23168, -2.82 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.41 DPS) [quest]; Gut Ripper (2164, -6.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 156.2 spell_power points (40.06 DPS) | yes | Umbral Wand (5216, -0.61 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.47 DPS) [dungeon]; Twisted Nether Wand (249144, -4.52 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 146.4. Weights run: 2.0s. Verify run: 1.3s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.735, intellect=1.960 ± 0.046, crit=0.327 ± 0.012 per rating point (14 rating = 1%, 4.579 per %), hit=0.721 ± 0.006 per rating point (10 rating = 1%, 7.206 per %), spell_haste=not significant (-0.632 ± 0.764), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.480 ± 0.735), fire_power=1.502 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 58.2 spell_power points (5.16 DPS) | yes | Soulcatcher Halo (10630, -0.82 DPS) [dungeon]; Blood Guard's Dreadweave Hat (220907, -1.43 DPS) [vendor]; Chief Architect's Monocle (11839, -4.22 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 27.4 spell_power points (2.43 DPS) | yes | Gemshard Heart (17707, -0.69 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.69 DPS) [quest]; Scorn's Icy Choker (23169, -0.77 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 48.3 spell_power points (4.28 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.05 DPS) [crafted]; Inquisitor's Shawl (19507, -1.40 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 29.6 spell_power points (2.63 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.27 DPS) [dungeon]; Spritecaster Cape (11623, -0.34 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -2.46 DPS, sim-verified) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 58.2 spell_power points (5.16 DPS) | yes | Runecloth Robe (13858, -1.23 DPS) [crafted]; Hibernal Robe (8113, -1.68 DPS) [world_drop]; Robes of Insight (940, -4.21 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 29.4 spell_power points (2.61 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Shizzle's Nozzle Wiper (11917, -0.52 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 50.1 spell_power points (4.44 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.02 DPS, sim-verified) [vendor]; Red Mageweave Gloves (10018, -1.73 DPS) [crafted]; Runecloth Gloves (13863, -1.82 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 43.2 spell_power points (3.83 DPS) | yes | Deathmage Sash (10771, -0.61 DPS) [dungeon]; Ban'thok Sash (11662, -0.79 DPS) [dungeon]; Satyrmane Sash (17755, -0.85 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 42.6 spell_power points (3.78 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.38 DPS) [dungeon]; Red Mageweave Pants (10009, -0.45 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.49 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 32.8 spell_power points (2.91 DPS) | yes | Gilded Sandals (254107, -0.37 DPS) [crafted]; Coldstone Slippers (18697, -0.48 DPS) [dungeon]; Southsea Mojo Boots (20641, -1.15 DPS, sim-verified) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.4 spell_power points (2.61 DPS) | yes | Philanthropist's Ring (281635, -0.68 DPS) [quest]; Mindseye Circle (10634, -1.00 DPS, sim-verified) [dungeon]; Woodseed Hoop (17768, -1.04 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Philanthropist's Ring (281635, -0.09 DPS) [quest]; Woodseed Hoop (17768, -0.45 DPS) [quest]; Mindseye Circle (10634, -2.65 DPS, sim-verified) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.62 DPS) [quest]; Uther's Strength (11302, -1.28 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.04 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.13 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, +0.00 DPS) [quest]; Hanzo Sword (8190, -4.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.02 DPS) [world_drop]; Flaming Incinerator (9483, -3.22 DPS) [dungeon]; Pyric Caduceus (11748, -3.96 DPS, sim-verified) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 271.7. Weights run: 8.2s. Verify run: 1.2s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± 0.378), intellect=not significant (-2.313 ± 0.023), crit=not significant (-0.457 ± 0.008) per rating point (14 rating = 1%, -6.402 per %), hit=not significant (-0.990 ± 0.004) per rating point (10 rating = 1%, -9.899 per %), spell_haste=not significant (-2.393 ± 0.292), spell_penetration=not significant (-0.000 ± 0.000), shadow_power=not significant (2.392 ± 0.378), fire_power=not significant (-1.419 ± 0.002)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | 72.2 spell_power points | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; The Postmaster's Band (13390, -0.82 DPS) [dungeon]; Magister's Crown (16686, -2.06 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 40.5 spell_power points | yes | Beads of Ogre Mojo (22149, -0.35 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -0.46 DPS) [world_drop]; Lady Maye's Pendant (14558, -2.08 DPS, sim-verified) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dreadweave Mantle (231592, -0.24 DPS) [pvp]; Burial Shawl (18681, -0.85 DPS) [dungeon]; Darkspear Shoulderpads (272103, -6.14 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 38.9 spell_power points | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Royal Tribunal Cloak (13376, -0.67 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -0.67 DPS) [vendor] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Dreadweave Tunic (227094, +0.00 DPS) [pvp]; Warlord's Dreadweave Robe (231591, +0.00 DPS) [pvp]; Magister's Robes (16688, -8.66 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 37.7 spell_power points | yes | General's Dreadweave Bracers (17587, -0.04 DPS) [pvp]; Sublime Wristguards (18497, -0.54 DPS) [dungeon]; Magiskull Cuffs (13107, -3.03 DPS, sim-verified) [world_drop] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 60.9 spell_power points | yes | Deathmist Wraps (226911, -1.35 DPS) [quest]; Mooncloth Gloves (18409, -1.64 DPS) [crafted]; General's Dreadweave Gloves (231589, -1.70 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 75.3 spell_power points | yes | Belt of the Archmage (18405, -1.46 DPS, sim-verified) [crafted]; Magister's Belt (16685, -2.41 DPS) [dungeon]; General's Dreadweave Belt (17589, -2.68 DPS) [pvp] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Deathmist Leggings (226910, -0.29 DPS) [quest]; Outrider's Silk Leggings (22747, -4.43 DPS, sim-verified) [rep] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 49.4 spell_power points | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.52 DPS) [dungeon]; Blood Guard's Dreadweave Walkers (227098, -0.52 DPS) [pvp] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -0.79 DPS) [vendor]; Seal of Rivendare (13345, -0.91 DPS) [dungeon]; Naglering (11669, -4.21 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -0.01 DPS) [vendor]; Seal of Rivendare (13345, -0.12 DPS) [dungeon]; Naglering (11669, -4.65 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, -0.53 DPS) [quest]; Weakness Analyzer (272438, -0.62 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Trindlehaven Staff (13161) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Argent Crusader (13249, -0.37 DPS) [quest]; Teebu's Blazing Longsword (1728, -5.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Oblivion's Touch (18761, -0.41 DPS) [dungeon]; Sparkling Crystal Wand (20672, -0.49 DPS) [world]; Torch of Light (279246, -8.66 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Trindlehaven Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

