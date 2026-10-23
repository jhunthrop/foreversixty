# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 1531001000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 40.4. Weights run: 1.0s. Verify run: 0.7s. 145 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.289 ± 0.009, crit=0.056 ± 0.002 per rating point (14 rating = 1%, 0.783 per %), hit=0.193 ± 0.009 per rating point (10 rating = 1%, 1.932 per %), spell_haste=0.467 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.265 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.85 DPS) | yes | Shadow Goggles (4373, -1.63 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.6 spell_power points (1.08 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.60 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.57 DPS) | yes | Caretaker's Cape (20428, -0.14 DPS) [rep]; Pearl-clasped Cloak (5542, -0.16 DPS) [crafted]; Black Whelp Cloak (7283, -0.51 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.4 spell_power points (0.91 DPS) | yes | Green Woolen Vest (2582, -0.35 DPS) [crafted]; Bloody Apron (6226, -0.35 DPS) [dungeon]; Gray Woolen Robe (2585, -1.31 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.4 spell_power points (0.20 DPS) | yes | Windsong Bangles (263336, -0.06 DPS) [quest]; Repurposed Hair Band (281256, -0.12 DPS) [quest]; Bright Bracers (3647, -0.83 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.99 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.30 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.63 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.73 DPS) | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Keller's Girdle (2911, -0.40 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.02 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.3 spell_power points (1.60 DPS) | yes | Filigreed Pristine Leggings (253937, -0.51 DPS) [crafted]; Silk-threaded Trousers (1929, -0.61 DPS) [dungeon]; Rumpled Kilt (274741, -0.89 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (1.15 DPS) | yes | Pristine Boots (253889, -0.61 DPS) [crafted]; Red Woolen Boots (4313, -0.62 DPS, sim-verified) [crafted]; Woolen Boots (2583, -0.79 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.79 DPS) | yes | Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.67 DPS) [world_drop]; The 1 Ring (8350, -0.75 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.71 DPS) | yes | Lavishly Jeweled Ring (1156, -0.54 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.58 DPS) [world_drop]; The 1 Ring (8350, -0.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.9 spell_power points (0.41 DPS) | yes | Witching Stave (1484, -0.20 DPS) [dungeon]; Staff of Westfall (2042, -0.20 DPS) [quest]; Lesser Staff of the Spire (1300, -0.57 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 160.7 spell_power points (22.73 DPS) | yes | Skycaller (12984, -0.64 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.20 DPS) [quest]; Sizzle Stick (8071, -4.38 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 145, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 1531005113100000000-00000000000000000-000000000000000000)

Set DPS (verified): 97.0. Weights run: 1.2s. Verify run: 0.9s. 264 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.573 ± 0.021, crit=0.069 ± 0.002 per rating point (14 rating = 1%, 0.964 per %), hit=0.352 ± 0.038 per rating point (10 rating = 1%, 3.525 per %), spell_haste=-6.291 ± 0.521, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.575 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (97.0 DPS) | yes | Silk Headband (7050, -0.40 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.60 DPS) [crafted]; Enchanter's Cowl (4322, -3.07 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (2.10 DPS) | yes | Darkspear Warding Pendant (272075, -1.52 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.64 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.97 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.2 spell_power points (2.84 DPS) | yes | Death Speaker Mantle (6685, -0.37 DPS) [dungeon]; Fairywing Mantle (9536, -0.60 DPS) [quest]; Magician's Mantle (12998, -0.80 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.00 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272078, -0.08 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.20 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.4 spell_power points (3.30 DPS) | yes | Beguiler Robes (7728, -0.12 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.32 DPS) [dungeon]; Death Speaker Robes (6682, -0.63 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.81 DPS) | yes | Nightsky Wristbands (6407, -1.12 DPS) [world_drop]; Stonecloth Bindings (14416, -1.23 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.74 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.0 spell_power points (1.41 DPS) | yes | Serpent Gloves (5970, -0.00 DPS) [dungeon]; Shilly Mitts (9609, -0.00 DPS) [quest]; Truefaith Gloves (7049, -0.06 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (2.55 DPS) | yes | Belt of Arugal (6392, -0.40 DPS) [dungeon]; Crimson Silk Belt (7055, -0.54 DPS) [crafted]; Invoker's Cord (215366, -0.57 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.6 spell_power points (2.73 DPS) | yes | Gaze Dreamer Pants (6903, -0.32 DPS) [dungeon]; Pristine Leggings (253987, -0.52 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.83 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (2.21 DPS) | yes | Spidersilk Boots (4320, -0.35 DPS) [crafted]; Glinteye Slippers (273024, -0.81 DPS) [dungeon]; Acidic Walkers (9454, -2.13 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.9 spell_power points (1.78 DPS) | yes | Minor Channeling Ring (1449, -0.55 DPS) [quest]; Ring of Power Regulation (273030, -0.58 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.58 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.41 DPS) | yes | Minor Channeling Ring (1449, -0.17 DPS) [quest]; Ring of Power Regulation (273030, -0.20 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.20 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (7.23 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.21 DPS) [quest]; Glimmering Staff (249392, -5.96 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 166.0 spell_power points (33.32 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.32 DPS) [crafted]; Gravestone Scepter (7001, -4.32 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 1531005113100001531-00000000000000000-000000000000000000)

Set DPS (verified): 153.7. Weights run: 1.3s. Verify run: 0.9s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.870 ± 0.051, crit=0.168 ± 0.005 per rating point (14 rating = 1%, 2.352 per %), hit=0.541 ± 0.079 per rating point (10 rating = 1%, 5.411 per %), spell_haste=-13.241 ± 1.144, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.599 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 22.8 spell_power points (4.88 DPS) | yes | Augural Shroud (2620, -0.66 DPS) [world]; Whitemane's Chapeau (7720, -1.52 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.67 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.0 spell_power points (2.99 DPS) | yes | Triune Amulet (7722, -1.28 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.68 DPS) [vendor]; Scorn's Icy Choker (23169, -2.35 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.3 spell_power points (3.92 DPS) | yes | Bloodmage Mantle (7684, -0.32 DPS) [dungeon]; Berylline Pads (4197, -0.56 DPS) [quest]; Green Silken Shoulders (7057, -2.24 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.4 spell_power points (2.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.17 DPS) [vendor]; Icy Cloak (4327, -0.72 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.2 spell_power points (5.83 DPS) | yes | Dreamweave Vest (10021, -0.30 DPS) [crafted]; Robe of Power (7054, -0.59 DPS) [crafted]; Elemental Raiment (9434, -1.33 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.93 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Windchaser Cuffs (14429, -0.25 DPS) [world_drop]; Condor Bracers (15864, -0.43 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.5 spell_power points (4.60 DPS) | yes | Red Mageweave Gloves (10018, -0.38 DPS) [crafted]; Black Mageweave Gloves (10003, -1.39 DPS) [crafted]; Stormcloth Gloves (10011, -1.51 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.1 spell_power points (4.29 DPS) | yes | Highlander's Cloth Girdle (20098, -0.55 DPS) [rep]; Gilded Cord (254037, -1.09 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.4 spell_power points (5.23 DPS) | yes | Crimson Silk Pantaloons (7062, -1.10 DPS) [crafted]; Abomination Skin Leggings (23173, -1.82 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.24 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.14 DPS) | yes | Gilded Slippers (254001, -2.33 DPS) [crafted]; Acidic Walkers (9454, -2.58 DPS) [dungeon]; Spidersilk Boots (4320, -2.89 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (3.07 DPS) | yes | Lorekeeper's Ring (19524, -1.15 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.36 DPS) [quest]; Reedknot Ring (9622, -1.57 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.4 spell_power points (2.22 DPS) | yes | Lorekeeper's Ring (19524, -0.29 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.50 DPS) [quest]; Reedknot Ring (9622, -0.72 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (153.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.37 DPS) [dungeon]; Spine of the Devourer (273040, -3.20 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.6 spell_power points (3.34 DPS) | yes | Orb of Lorica (11262, -0.19 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.34 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.73 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 189.2 spell_power points (40.50 DPS) | yes | Nether Force Wand (11263, -2.61 DPS) [quest]; Icefury Wand (7514, -2.75 DPS) [quest]; Ragefire Wand (7513, -2.80 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 1531005113100001531-03202300000000000-000000000000000000)

Set DPS (verified): 253.3. Weights run: 1.2s. Verify run: 1.2s. 436 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.114 ± 0.075, crit=0.257 ± 0.007 per rating point (14 rating = 1%, 3.593 per %), hit=0.763 ± 0.121 per rating point (10 rating = 1%, 7.627 per %), spell_haste=-22.057 ± 1.737, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.613 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.3 spell_power points (8.48 DPS) | yes | Dreamweave Circlet (10041, -1.88 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.12 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.93 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.9 spell_power points (3.27 DPS) | yes | Horizon Choker (13085, -0.07 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.46 DPS) [dungeon]; Mindburst Medallion (11196, -0.66 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.0 spell_power points (6.79 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -2.35 DPS) [vendor]; Inquisitor's Shawl (19507, -2.38 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.65 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.7 spell_power points (4.25 DPS) | yes | Runecloth Cloak (13860, -0.57 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.05 DPS) [vendor]; Big Voodoo Cloak (8216, -1.16 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.3 spell_power points (8.48 DPS) | yes | Runecloth Tunic (13857, -2.47 DPS) [crafted]; Robe of the Magi (1716, -2.59 DPS) [world_drop]; Runecloth Robe (13858, -4.32 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 15.0 spell_power points (3.09 DPS) | yes | Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Forgotten Wraps (9433, -0.34 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.34 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 35.2 spell_power points (7.24 DPS) | yes | Raider Handwraps (272098, -0.76 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.51 DPS) [vendor]; Dreamweave Gloves (10019, -2.62 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 27.2 spell_power points (5.58 DPS) | yes | Deathmage Sash (10771, -0.71 DPS) [dungeon]; Runecloth Belt (13856, -1.19 DPS) [crafted]; Satyrmane Sash (17755, -3.37 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 29.0 spell_power points (5.95 DPS) | yes | Red Mageweave Pants (10009, -0.33 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.21 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.33 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 25.6 spell_power points (5.27 DPS) | yes | Earthen Silk Slippers (254013, -0.34 DPS) [crafted]; Gilded Sandals (254107, -0.95 DPS) [crafted]; Southsea Mojo Boots (20641, -1.11 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.6 spell_power points (3.20 DPS) | yes | Lorekeeper's Ring (19523, -0.73 DPS) [rep]; Snake Hoop (6750, -0.82 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.67 DPS) | yes | Lorekeeper's Ring (19523, -0.21 DPS) [rep]; Snake Hoop (6750, -0.29 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (253.3 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (253.3 DPS) | yes | Mark of the Chosen (17774, -4.09 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (253.3 DPS) | yes | Spellshifter Rod (9527, -0.55 DPS) [quest]; Kindling Stave (11750, -1.84 DPS) [dungeon]; Hanzo Sword (8190, -43.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 255.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.47 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.11 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 1531005113100001531-03202300000000000-053200000000000000)

Set DPS (verified): 453.6. Weights run: 1.4s. Verify run: 2.5s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.069 ± 0.082, crit=0.409 ± 0.011 per rating point (14 rating = 1%, 5.722 per %), hit=1.170 ± 0.154 per rating point (10 rating = 1%, 11.700 per %), spell_haste=-28.978 ± 2.102, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.571 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 53.5 spell_power points (12.94 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.81 DPS) [pvp]; Magister's Crown (16686, -7.09 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (453.6 DPS) | yes | Beads of Ogre Mojo (22149, -0.74 DPS) [quest]; Pebble of Kajaro (19600, -1.45 DPS) [quest]; Jewel of Kajaro (19601, -1.45 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.8 spell_power points (11.80 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.87 DPS) [pvp]; Darkspear Shoulderpads (272103, -1.91 DPS) [vendor]; Mantle of the Timbermaw (19050, -2.94 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.3 spell_power points (8.78 DPS) | yes | Hide of the Wild (18510, -2.80 DPS) [crafted]; Spritecaster Cape (11623, -3.83 DPS) [dungeon]; Runecloth Cloak (13860, -4.53 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.6 spell_power points (14.17 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.40 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.31 DPS) [pvp]; Sorcerer's Robes (226932, -6.74 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 30.6 spell_power points (7.40 DPS) | yes | Runecloth Cuffs (254123, -2.14 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.48 DPS) [pvp]; Magiskull Cuffs (13107, -3.00 DPS) [world_drop] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 38.7 spell_power points (9.36 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 60.2 spell_power points (14.58 DPS) | yes | Belt of the Archmage (18405, -4.21 DPS) [crafted]; Magister's Belt (16685, -7.45 DPS) [dungeon]; Magician's Cord (272393, -9.86 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.1 spell_power points (13.10 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -1.53 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -2.24 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 38.1 spell_power points (9.22 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.28 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 49.3 spell_power points (11.94 DPS) | yes | Channeler's Ring (272406, -5.48 DPS) [vendor]; Maiden's Circle (13001, -5.77 DPS) [world_drop]; Deep Rooted Ring (19109, -6.45 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 33.8 spell_power points (8.17 DPS) | yes | Channeler's Ring (272406, -1.72 DPS) [vendor]; Maiden's Circle (13001, -2.00 DPS) [world_drop]; Deep Rooted Ring (19109, -2.68 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (453.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (453.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.84 DPS) [crafted]; Uther's Strength (11302, -1.94 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (453.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.58 DPS) [quest]; Teebu's Blazing Longsword (1728, -51.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 323.1 spell_power points (78.21 DPS) | yes | Wand of Biting Cold (19108, -5.14 DPS, sim-verified) [quest]; Stormrager (16997, -14.23 DPS) [quest]; Brilliant Wand (249385, -14.64 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 1531005113100101531-03000000000000000-053500030000000000)

Set DPS (verified): 678.8. Weights run: 1.5s. Verify run: 2.5s. 891 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.430 ± 0.013, crit=0.467 ± 0.011 per rating point (14 rating = 1%, 6.536 per %), hit=0.997 ± 0.073 per rating point (10 rating = 1%, 9.965 per %), spell_haste=-12.652 ± 0.936, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.511 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 38.3 spell_power points (17.21 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.35 DPS) [pvp]; Red Mageweave Headband (10033, -7.33 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Jewel of Kajaro (19601, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.64 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 40.0 spell_power points (17.97 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.84 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.50 DPS, sim-verified) [crafted]; Lieutenant Commander's Silk Mantle (227102, -6.17 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.4 spell_power points (13.22 DPS) | yes | Hide of the Wild (18510, -4.99 DPS) [crafted]; Spritecaster Cape (11623, -5.77 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -6.92 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 51.7 spell_power points (23.24 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.18 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -7.58 DPS) [pvp]; Acumen Robes (17775, -13.94 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.4 spell_power points (11.44 DPS) | yes | Runecloth Cuffs (254123, -4.56 DPS) [crafted]; Sorcerer's Bindings (226929, -6.36 DPS) [quest]; Bracers of Hope (22667, -6.61 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.1 spell_power points (13.10 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS, sim-verified) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 44.4 spell_power points (19.97 DPS) | yes | Magician's Cord (272393, -6.79 DPS) [vendor]; Belt of the Archmage (18405, -7.06 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -9.58 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.4 spell_power points (22.19 DPS) | yes | Marshal's Silk Leggings (231605, -1.91 DPS) [pvp]; Sorcerer's Leggings (226933, -6.08 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -6.53 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 27.9 spell_power points (12.53 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -1.16 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 41.8 spell_power points (18.81 DPS) | yes | Maiden's Circle (13001, -9.36 DPS) [world_drop]; Blessed Band of Light (272407, -9.83 DPS) [vendor]; Don Julio's Band (19325, -11.39 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 26.7 spell_power points (12.02 DPS) | yes | Maiden's Circle (13001, -2.57 DPS) [world_drop]; Blessed Band of Light (272407, -3.04 DPS) [vendor]; Don Julio's Band (19325, -4.60 DPS) [rep] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (678.8 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.26 DPS) [crafted]; Darkmoon Card: Blue Dragon (19288, -8.52 DPS, sim-verified) [quest] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.57 DPS) [quest]; Teebu's Blazing Longsword (1728, -43.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 176.7 spell_power points (79.46 DPS) | yes | Brilliant Wand (249385, -14.48 DPS) [crafted]; Wand of Biting Cold (19108, -15.46 DPS) [quest]; Stormrager (16997, -15.80 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Serenity Field; trinket2: Talisman of Ascendance; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (orc, 1531001000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 37.9. Weights run: 1.0s. Verify run: 0.8s. 134 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.289 ± 0.009, crit=0.056 ± 0.002 per rating point (14 rating = 1%, 0.783 per %), hit=0.193 ± 0.009 per rating point (10 rating = 1%, 1.932 per %), spell_haste=0.467 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.265 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.85 DPS) | yes | Shadow Goggles (4373, -1.30 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.6 spell_power points (1.08 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.49 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.57 DPS) | yes | Black Whelp Cloak (7283, -0.14 DPS) [crafted]; Battle Healer's Cloak (20427, -0.14 DPS) [rep]; Pearl-clasped Cloak (5542, -0.16 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.4 spell_power points (0.91 DPS) | yes | Green Woolen Vest (2582, -0.35 DPS) [crafted]; Bloody Apron (6226, -0.35 DPS) [dungeon]; Gray Woolen Robe (2585, -1.16 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.42 DPS) | yes | Owlbeard Bracers (16981, -0.20 DPS) [quest]; Mindthrust Bracers (1974, -0.22 DPS) [dungeon]; Featherbead Bracers (15452, -0.22 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.99 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.30 DPS) [crafted]; Apothecary Gloves (10919, -0.42 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.73 DPS) | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Keller's Girdle (2911, -0.40 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.87 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (37.9 DPS) | yes | Silk-threaded Trousers (1929, -0.10 DPS) [dungeon]; Rumpled Kilt (274741, -0.39 DPS) [vendor]; Abomination Skin Leggings (23173, -0.87 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (1.15 DPS) | yes | Pristine Boots (253889, -0.61 DPS) [crafted]; Woolen Boots (2583, -0.79 DPS) [crafted]; Red Woolen Boots (4313, -0.82 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.71 DPS) | yes | Loop of Sacrifice (281673, -0.50 DPS) [quest]; Volcanic Rock Ring (12053, -0.58 DPS) [world_drop]; The 1 Ring (8350, -0.67 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.7 spell_power points (0.25 DPS) | yes | Loop of Sacrifice (281673, -0.04 DPS) [quest]; Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; The 1 Ring (8350, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.9 spell_power points (0.41 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.16 DPS) [world]; Clear Crystal Rod (16894, -0.20 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 160.7 spell_power points (22.73 DPS) | yes | Skycaller (12984, -0.54 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.38 DPS) [quest]; Greater Magic Wand (11288, -4.95 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 134, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (orc, 1531005113100000000-00000000000000000-000000000000000000)

Set DPS (verified): 94.5. Weights run: 1.2s. Verify run: 0.8s. 245 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.573 ± 0.021, crit=0.069 ± 0.002 per rating point (14 rating = 1%, 0.964 per %), hit=0.352 ± 0.038 per rating point (10 rating = 1%, 3.525 per %), spell_haste=-6.291 ± 0.521, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.575 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (94.5 DPS) | yes | Silk Headband (7050, -0.40 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.60 DPS) [crafted]; Enchanter's Cowl (4322, -2.09 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (2.10 DPS) | yes | Darkspear Warding Pendant (272075, -1.52 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.64 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.93 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.2 spell_power points (2.84 DPS) | yes | Death Speaker Mantle (6685, -0.37 DPS) [dungeon]; Fairywing Mantle (9536, -0.60 DPS) [quest]; Magician's Mantle (12998, -0.80 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.00 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.08 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.20 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.4 spell_power points (3.30 DPS) | yes | Beguiler Robes (7728, -0.12 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.32 DPS) [dungeon]; Death Speaker Robes (6682, -0.63 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.81 DPS) | yes | Nightsky Wristbands (6407, -1.12 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.20 DPS) [quest]; Glowing Magical Bracelets (13106, -2.28 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.9 spell_power points (1.78 DPS) | yes | Serpent Gloves (5970, -0.37 DPS) [dungeon]; Truefaith Gloves (7049, -0.43 DPS) [crafted]; Gnoll Casting Gloves (892, -0.58 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.7 spell_power points (2.55 DPS) | yes | Warsong Sash (16975, -0.35 DPS) [quest]; Belt of Arugal (6392, -0.40 DPS) [dungeon]; Crimson Silk Belt (7055, -0.54 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.6 spell_power points (2.73 DPS) | yes | Gaze Dreamer Pants (6903, -0.32 DPS) [dungeon]; Pristine Leggings (253987, -0.52 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.83 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (2.21 DPS) | yes | Spidersilk Boots (4320, -0.35 DPS) [crafted]; Glinteye Slippers (273024, -0.81 DPS) [dungeon]; Acidic Walkers (9454, -2.46 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.9 spell_power points (1.78 DPS) | yes | Ring of Power Regulation (273030, -0.58 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.58 DPS) [vendor]; Solid Gold Band (279845, -0.98 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.41 DPS) | yes | Ring of Power Regulation (273030, -0.20 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.20 DPS) [vendor]; Solid Gold Band (279845, -0.60 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (7.23 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -5.96 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -6.08 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 166.0 spell_power points (33.32 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.32 DPS) [crafted]; Unstable Power Core (279847, -3.61 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 245, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (orc, 1531005113100001531-00000000000000000-000000000000000000)

Set DPS (verified): 148.0. Weights run: 1.3s. Verify run: 1.0s. 320 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.870 ± 0.051, crit=0.168 ± 0.005 per rating point (14 rating = 1%, 2.352 per %), hit=0.541 ± 0.079 per rating point (10 rating = 1%, 5.411 per %), spell_haste=-13.241 ± 1.144, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.599 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 22.8 spell_power points (4.88 DPS) | yes | Augural Shroud (2620, -0.66 DPS) [world]; Whitemane's Chapeau (7720, -1.52 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -3.02 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.0 spell_power points (2.99 DPS) | yes | Triune Amulet (7722, -1.28 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.68 DPS) [vendor]; Scorn's Icy Choker (23169, -2.17 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.3 spell_power points (3.92 DPS) | yes | Bloodmage Mantle (7684, -0.32 DPS) [dungeon]; Berylline Pads (4197, -0.56 DPS) [quest]; Green Silken Shoulders (7057, -2.06 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.4 spell_power points (2.22 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.17 DPS) [vendor]; Icy Cloak (4327, -0.72 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.2 spell_power points (5.83 DPS) | yes | Dreamweave Vest (10021, -0.30 DPS) [crafted]; Robe of Power (7054, -0.59 DPS) [crafted]; Elemental Raiment (9434, -1.33 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 11.0 spell_power points (2.35 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.42 DPS) [dungeon]; Windchaser Cuffs (14429, -0.67 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.5 spell_power points (4.60 DPS) | yes | Red Mageweave Gloves (10018, -0.38 DPS) [crafted]; Black Mageweave Gloves (10003, -1.39 DPS) [crafted]; Stormcloth Gloves (10011, -1.51 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.1 spell_power points (4.29 DPS) | yes | Defiler's Cloth Girdle (20166, -0.55 DPS) [rep]; Gilded Cord (254037, -1.09 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.4 spell_power points (5.23 DPS) | yes | Crimson Silk Pantaloons (7062, -1.10 DPS) [crafted]; Abomination Skin Leggings (23173, -1.82 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.24 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.14 DPS) | yes | Gilded Slippers (254001, -2.33 DPS) [crafted]; Acidic Walkers (9454, -2.58 DPS) [dungeon]; Spidersilk Boots (4320, -2.89 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (3.07 DPS) | yes | Advisor's Ring (19520, -1.15 DPS) [rep]; Reedknot Ring (9622, -1.57 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.4 spell_power points (2.22 DPS) | yes | Advisor's Ring (19520, -0.29 DPS) [rep]; Reedknot Ring (9622, -0.72 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (148.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.37 DPS) [dungeon]; Spine of the Devourer (273040, -3.20 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.6 spell_power points (3.34 DPS) | yes | Orb of the Forgotten Seer (7685, -0.34 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.73 DPS) [crafted]; Omega Orb (7749, -1.41 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 189.2 spell_power points (40.50 DPS) | yes | Nether Force Wand (11263, -2.61 DPS) [quest]; Icefury Wand (7514, -2.75 DPS) [quest]; Ragefire Wand (7513, -2.80 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 320, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 1531005113100001531-03202300000000000-000000000000000000)

Set DPS (verified): 244.6. Weights run: 1.2s. Verify run: 1.2s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.114 ± 0.075, crit=0.257 ± 0.007 per rating point (14 rating = 1%, 3.593 per %), hit=0.763 ± 0.121 per rating point (10 rating = 1%, 7.627 per %), spell_haste=-22.057 ± 1.737, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.613 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.3 spell_power points (8.48 DPS) | yes | Dreamweave Circlet (10041, -1.88 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.12 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.93 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.9 spell_power points (3.27 DPS) | yes | Horizon Choker (13085, -0.07 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.46 DPS) [dungeon]; Mindburst Medallion (11196, -0.66 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.0 spell_power points (6.79 DPS) | yes | Red Mageweave Shoulders (10029, -1.92 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.35 DPS) [vendor]; Inquisitor's Shawl (19507, -2.38 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.0 spell_power points (4.53 DPS) | yes | Spritecaster Cape (11623, -0.28 DPS) [dungeon]; Runecloth Cloak (13860, -0.85 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.32 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.3 spell_power points (8.48 DPS) | yes | Runecloth Robe (13858, -2.33 DPS) [crafted]; Runecloth Tunic (13857, -2.47 DPS) [crafted]; Robe of the Magi (1716, -2.59 DPS) [world_drop] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 15.0 spell_power points (3.09 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.05 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 35.2 spell_power points (7.24 DPS) | yes | Raider Handwraps (272098, -0.76 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.51 DPS) [vendor]; Dreamweave Gloves (10019, -2.62 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 27.2 spell_power points (5.58 DPS) | yes | Satyrmane Sash (17755, -0.42 DPS) [dungeon]; Deathmage Sash (10771, -0.71 DPS) [dungeon]; Runecloth Belt (13856, -1.19 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 29.0 spell_power points (5.95 DPS) | yes | Red Mageweave Pants (10009, -0.33 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.21 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.33 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (244.6 DPS) | yes | Gilded Sandals (254107, -0.61 DPS) [crafted]; Southsea Mojo Boots (20641, -0.77 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -3.72 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.6 spell_power points (3.20 DPS) | yes | Advisor's Ring (19519, -0.73 DPS) [rep]; Snake Hoop (6750, -0.82 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.67 DPS) | yes | Advisor's Ring (19519, -0.21 DPS) [rep]; Snake Hoop (6750, -0.29 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.31 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.14 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.55 DPS) [quest]; Kindling Stave (11750, -1.84 DPS) [dungeon]; Hanzo Sword (8190, -40.66 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 255.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.47 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.11 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 1531005113100001531-03202300000000000-053200000000000000)

Set DPS (verified): 439.0. Weights run: 1.4s. Verify run: 2.5s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.069 ± 0.082, crit=0.409 ± 0.011 per rating point (14 rating = 1%, 5.722 per %), hit=1.170 ± 0.154 per rating point (10 rating = 1%, 11.700 per %), spell_haste=-28.978 ± 2.102, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.571 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 53.5 spell_power points (12.94 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.81 DPS) [pvp]; Magister's Crown (16686, -5.51 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (439.0 DPS) | yes | Beads of Ogre Mojo (22149, -0.74 DPS) [quest]; Pebble of Kajaro (19600, -1.45 DPS) [quest]; Jewel of Kajaro (19601, -1.45 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 48.8 spell_power points (11.80 DPS) | yes | Warlord's Silk Amice (231594, -1.87 DPS) [pvp]; Darkspear Shoulderpads (272103, -1.91 DPS) [vendor]; Mantle of the Timbermaw (19050, -2.94 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.3 spell_power points (8.78 DPS) | yes | Hide of the Wild (18510, -2.80 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.54 DPS) [quest]; Spritecaster Cape (11623, -3.83 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.6 spell_power points (14.17 DPS) | yes | Warlord's Silk Raiment (231596, -0.40 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.31 DPS) [pvp]; Sorcerer's Robes (226932, -7.81 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 30.6 spell_power points (7.40 DPS) | yes | Runecloth Cuffs (254123, -2.14 DPS) [crafted]; General's Silk Cuffs (16538, -2.48 DPS) [pvp]; Magiskull Cuffs (13107, -3.00 DPS) [world_drop] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 38.7 spell_power points (9.36 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 60.2 spell_power points (14.58 DPS) | yes | Belt of the Archmage (18405, -4.21 DPS) [crafted]; Magister's Belt (16685, -7.45 DPS) [dungeon]; Magician's Cord (272393, -9.46 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.1 spell_power points (13.10 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.41 DPS) [rep]; Sorcerer's Leggings (226933, -1.53 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 38.1 spell_power points (9.22 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.28 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 49.3 spell_power points (11.94 DPS) | yes | Channeler's Ring (272406, -5.48 DPS) [vendor]; Maiden's Circle (13001, -5.77 DPS) [world_drop]; Deep Rooted Ring (19109, -6.45 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 33.8 spell_power points (8.17 DPS) | yes | Channeler's Ring (272406, -1.72 DPS) [vendor]; Maiden's Circle (13001, -2.00 DPS) [world_drop]; Deep Rooted Ring (19109, -2.68 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (439.0 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (439.0 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.84 DPS) [crafted]; Rune of the Guard Captain (19120, -1.41 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (439.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.58 DPS) [quest]; Teebu's Blazing Longsword (1728, -54.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 323.1 spell_power points (78.21 DPS) | yes | Wand of Biting Cold (19108, -6.93 DPS, sim-verified) [quest]; Stormrager (16997, -14.23 DPS) [quest]; Brilliant Wand (249385, -14.64 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 1531005113100101531-03000000000000000-053500030000000000)

Set DPS (verified): 670.6. Weights run: 1.5s. Verify run: 2.4s. 880 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.430 ± 0.013, crit=0.467 ± 0.011 per rating point (14 rating = 1%, 6.536 per %), hit=0.997 ± 0.073 per rating point (10 rating = 1%, 9.965 per %), spell_haste=-12.652 ± 0.936, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.511 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 38.3 spell_power points (17.21 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.35 DPS) [pvp]; Red Mageweave Headband (10033, -7.75 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (670.6 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Jewel of Kajaro (19601, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.64 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 40.0 spell_power points (17.97 DPS) | yes | Warlord's Silk Amice (231594, -3.84 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.53 DPS, sim-verified) [crafted]; Champion's Silk Mantle (227104, -6.17 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.4 spell_power points (13.22 DPS) | yes | Hide of the Wild (18510, -4.99 DPS) [crafted]; Spritecaster Cape (11623, -5.77 DPS) [dungeon]; Deep Woodlands Cloak (19121, -6.09 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 51.7 spell_power points (23.24 DPS) | yes | Warlord's Silk Raiment (231596, -2.18 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -7.58 DPS) [pvp]; Acumen Robes (17775, -14.20 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.4 spell_power points (11.44 DPS) | yes | Runecloth Cuffs (254123, -4.56 DPS) [crafted]; Sorcerer's Bindings (226929, -6.36 DPS) [quest]; Bracers of Hope (22667, -6.61 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 29.1 spell_power points (13.10 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 44.4 spell_power points (19.97 DPS) | yes | Magician's Cord (272393, -6.79 DPS) [vendor]; Belt of the Archmage (18405, -8.65 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -9.58 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.4 spell_power points (22.19 DPS) | yes | General's Silk Trousers (231595, -1.91 DPS) [pvp]; Sorcerer's Leggings (226933, -6.08 DPS) [quest]; Outrider's Silk Leggings (22747, -10.64 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 27.9 spell_power points (12.53 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -1.16 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 41.8 spell_power points (18.81 DPS) | yes | Maiden's Circle (13001, -9.36 DPS) [world_drop]; Blessed Band of Light (272407, -9.83 DPS) [vendor]; Don Julio's Band (19325, -11.39 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 26.7 spell_power points (12.02 DPS) | yes | Maiden's Circle (13001, -2.57 DPS) [world_drop]; Blessed Band of Light (272407, -3.04 DPS) [vendor]; Don Julio's Band (19325, -4.60 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (670.6 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (670.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.26 DPS) [crafted]; Rune of the Guard Captain (19120, -3.16 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (670.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.57 DPS) [quest]; Teebu's Blazing Longsword (1728, -43.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 176.7 spell_power points (79.46 DPS) | yes | Brilliant Wand (249385, -14.48 DPS) [crafted]; Wand of Biting Cold (19108, -15.46 DPS) [quest]; Stormrager (16997, -15.80 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 880, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

