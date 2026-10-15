# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 153002000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 29.9. Weights run: 0.8s. Verify run: 0.6s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.353 ± 0.008, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.079 per %), hit=0.275 ± 0.009 per rating point (10 rating = 1%, 2.754 per %), spell_haste=0.464 ± 0.089, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.465 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -1.09 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.2 spell_power points (0.77 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.39 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.76 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Pearl-clasped Cloak (5542, -0.38 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.64 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.80 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.8 spell_power points (0.17 DPS) | yes | Windsong Bangles (263336, -0.07 DPS) [quest]; Repurposed Hair Band (281256, -0.10 DPS) [quest]; Bright Bracers (3647, -0.80 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.66 DPS) | yes | Gnoll Casting Gloves (892, -0.18 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.18 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.51 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.24 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.57 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.8 spell_power points (1.11 DPS) | yes | Filigreed Pristine Leggings (253937, -0.35 DPS) [crafted]; Silk-threaded Trousers (1929, -0.45 DPS) [dungeon]; Rumpled Kilt (274741, -0.64 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.79 DPS) | yes | Pristine Boots (253889, -0.24 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.42 DPS) [crafted]; Woolen Boots (2583, -0.54 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.54 DPS) | yes | Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.44 DPS) [world_drop]; The 1 Ring (8350, -0.50 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.47 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.44 DPS) [world]; Lavishly Jeweled Ring (1156, -0.68 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.5 spell_power points (0.33 DPS) | yes | Witching Stave (1484, -0.17 DPS) [dungeon]; Staff of Westfall (2042, -0.17 DPS) [quest]; Lesser Staff of the Spire (1300, -0.53 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 239.6 spell_power points (22.59 DPS) | yes | Skycaller (12984, -0.49 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.06 DPS) [quest]; Sizzle Stick (8071, -4.47 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 153005113100010000-00000000000000000-0000000000000000000)

Set DPS (verified): 118.4. Weights run: 1.0s. Verify run: 0.6s. 261 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.407 ± 0.015, crit=0.067 ± 0.002 per rating point (14 rating = 1%, 0.936 per %), hit=0.266 ± 0.027 per rating point (10 rating = 1%, 2.663 per %), spell_haste=not significant (0.544 ± 0.229), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.900 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.94 DPS) | yes | Silk Headband (7050, -0.53 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.80 DPS) [crafted]; Enchanter's Cowl (4322, -4.73 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (2.52 DPS) | yes | Darkspear Warding Pendant (272075, -1.98 DPS) [vendor]; Crystal Starfire Medallion (5003, -2.09 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.42 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.7 spell_power points (3.38 DPS) | yes | Death Speaker Mantle (6685, -0.58 DPS) [dungeon]; Fairywing Mantle (9536, -0.80 DPS) [quest]; Invoker's Mantle (215365, -0.97 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.34 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Prelacy Cape (7004, -0.27 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.3 spell_power points (3.82 DPS) | yes | Beguiler Robes (7728, -0.11 DPS) [dungeon]; Tree Bark Jacket (1486, -0.35 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.38 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.41 DPS) | yes | Nightsky Wristbands (6407, -1.75 DPS) [world_drop]; Stonecloth Bindings (14416, -1.86 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.68 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.87 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Truefaith Gloves (7049, -0.21 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.2 spell_power points (3.27 DPS) | yes | Belt of Arugal (6392, -0.53 DPS) [dungeon]; Invoker's Cord (215366, -0.85 DPS) [crafted]; Crimson Silk Belt (7055, -0.90 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.57 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.95 DPS) [crafted]; Abomination Skin Leggings (23173, -1.18 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.9 spell_power points (2.63 DPS) | yes | Acidic Walkers (9454, -0.43 DPS) [dungeon]; Glinteye Slippers (273024, -0.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.73 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 spell_power points (2.15 DPS) | yes | Ring of Power Regulation (273030, -0.54 DPS) [dungeon]; Minor Channeling Ring (1449, -0.59 DPS) [quest]; Lorekeeper's Ring (19525, -2.24 DPS, sim-verified) [rep] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Minor Channeling Ring (1449, -0.05 DPS) [quest]; Lorekeeper's Ring (19525, -1.48 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (9.62 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -4.28 DPS) [quest]; Glimmering Staff (249392, -8.42 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 124.6 spell_power points (33.30 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.96 DPS) [crafted]; Gravestone Scepter (7001, -4.30 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Sea Giant's Toe Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 261, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 153005113100011531-00000000000000000-0000000000000000000)

Set DPS (verified): 181.9. Weights run: 1.0s. Verify run: 0.7s. 339 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.585 ± 0.027, crit=0.172 ± 0.004 per rating point (14 rating = 1%, 2.410 per %), hit=0.398 ± 0.052 per rating point (10 rating = 1%, 3.978 per %), spell_haste=not significant (1.325 ± 0.390), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.918 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.49 DPS) | yes | Papal Fez (9431, -0.94 DPS) [dungeon]; Augural Shroud (2620, -1.28 DPS) [world]; Living Cowl (5608, -2.47 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (3.61 DPS) | yes | Triune Amulet (7722, -1.14 DPS) [dungeon]; Scorn's Icy Choker (23169, -1.93 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -2.34 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.51 DPS) | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.54 DPS) [quest]; Green Silken Shoulders (7057, -1.80 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.9 spell_power points (2.76 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.59 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.77 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (7.88 DPS) | yes | Dreamweave Vest (10021, -0.69 DPS) [crafted]; Robe of Power (7054, -1.39 DPS) [crafted]; Elemental Raiment (9434, -1.39 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.78 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.62 DPS) [quest]; Windchaser Cuffs (14429, -1.15 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (6.28 DPS) | yes | Black Mageweave Gloves (10003, -1.65 DPS) [crafted]; Red Mageweave Gloves (10018, -2.04 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.55 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.3 spell_power points (5.05 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Star Belt (4329, -1.03 DPS) [crafted]; Gilded Cord (254037, -1.13 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.49 DPS) | yes | Abomination Skin Leggings (23173, -2.27 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.28 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.79 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.41 DPS) | yes | Gilded Slippers (254001, -1.97 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.42 DPS) [dungeon]; Spidersilk Boots (4320, -4.53 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (3.99 DPS) | yes | Snake Hoop (6750, -1.24 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.52 DPS) [quest]; Reedknot Ring (9622, -1.83 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.78 DPS) | yes | Snake Hoop (6750, -0.02 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.31 DPS) [quest]; Reedknot Ring (9622, -0.62 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (181.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.36 DPS) [dungeon]; Spine of the Devourer (273040, -3.92 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (4.56 DPS) | yes | Orb of the Forgotten Seer (7685, -0.23 DPS) [dungeon]; Thrash's Trash (276204, -0.23 DPS) [vendor]; Orb of Lorica (11262, -0.90 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 131.0 spell_power points (40.46 DPS) | yes | Nether Force Wand (11263, -2.59 DPS) [quest]; Icefury Wand (7514, -2.73 DPS) [quest]; Ragefire Wand (7513, -2.78 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 339, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 153005113100011531-03202300000000000-0000000000000000000)

Set DPS (verified): 294.0. Weights run: 1.0s. Verify run: 1.1s. 432 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.713 ± 0.039, crit=0.264 ± 0.006 per rating point (14 rating = 1%, 3.692 per %), hit=0.676 ± 0.080 per rating point (10 rating = 1%, 6.757 per %), spell_haste=not significant (1.333 ± 0.597), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.921 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 33.3 spell_power points (10.49 DPS) | yes | Dreamweave Circlet (10041, -1.62 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.98 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.21 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (4.01 DPS) | yes | Scorn's Icy Choker (23169, -0.45 DPS) [dungeon]; Mindburst Medallion (11196, -0.77 DPS) [quest]; Horizon Choker (13085, -0.86 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.8 spell_power points (8.15 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -2.44 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.57 DPS) [crafted]; Black Mageweave Shoulders (10027, -2.97 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.3 spell_power points (5.76 DPS) | yes | Runecloth Cloak (13860, -2.08 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -2.16 DPS) [crafted]; Darkspear Raider's Cloak (272076, -2.61 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.3 spell_power points (10.49 DPS) | yes | Robe of the Magi (1716, -2.20 DPS) [world_drop]; Runecloth Tunic (13857, -2.65 DPS) [crafted]; Dreamweave Vest (10021, -2.79 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.0 spell_power points (3.78 DPS) | yes | Arcane Runed Bracers (4744, -0.94 DPS) [quest]; Spidertank Oilrag (9448, -0.94 DPS) [dungeon]; Bloodband Bracers (11469, -3.32 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.7 spell_power points (9.06 DPS) | yes | Raider Handwraps (272098, -1.91 DPS) [vendor]; Dreamweave Gloves (10019, -2.49 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.94 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 21.1 spell_power points (6.66 DPS) | yes | Dawnspire Cord (12466, -0.50 DPS) [dungeon]; Deathmage Sash (10771, -1.08 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.35 DPS) [rep] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 24.3 spell_power points (7.65 DPS) | yes | Red Mageweave Pants (10009, -0.53 DPS) [crafted]; Wizardweave Leggings (14132, -1.66 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.20 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.57 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.89 DPS) [vendor]; Gilded Sandals (254107, -2.07 DPS) [crafted]; Black Mageweave Boots (10026, -2.52 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.6 spell_power points (4.28 DPS) | yes | Lorekeeper's Ring (19523, -0.49 DPS) [rep]; Snake Hoop (6750, -1.26 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.10 DPS) | yes | Lorekeeper's Ring (19523, -0.32 DPS) [rep]; Snake Hoop (6750, -1.08 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (294.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (294.0 DPS) | yes | Mark of the Chosen (17774, -2.21 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (294.0 DPS) | yes | Kindling Stave (11750, -0.58 DPS) [dungeon]; Spire of Hakkar (10844, -2.65 DPS) [world]; Hanzo Sword (8190, -33.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 166.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.23 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Knight's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 432, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 498.6. Weights run: 1.1s. Verify run: 1.9s. 886 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.820 ± 0.046, crit=0.411 ± 0.009 per rating point (14 rating = 1%, 5.750 per %), hit=1.066 ± 0.109 per rating point (10 rating = 1%, 10.661 per %), spell_haste=not significant (2.031 ± 0.783), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.913 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.2 spell_power points (15.84 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.92 DPS) [pvp]; Magister's Crown (16686, -7.91 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (498.6 DPS) | yes | Beads of Ogre Mojo (22149, -0.95 DPS) [quest]; Chains of the Lich (23125, -1.23 DPS) [dungeon]; Jewel of Kajaro (19601, -2.57 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.0 spell_power points (15.10 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.60 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.90 DPS) [crafted]; Darkspear Shoulderpads (272103, -9.40 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.2 spell_power points (11.14 DPS) | yes | Hide of the Wild (18510, -3.70 DPS) [crafted]; Spritecaster Cape (11623, -4.79 DPS) [dungeon]; Runecloth Cloak (13860, -5.92 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.6 spell_power points (18.64 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.97 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -5.00 DPS) [vendor]; Sorcerer's Robes (226932, -13.85 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.6 spell_power points (9.58 DPS) | yes | Runecloth Cuffs (254123, -3.14 DPS) [crafted]; Marshal's Silk Bracers (16438, -4.35 DPS) [pvp]; Sorcerer's Bindings (226929, -4.48 DPS) [quest] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 34.1 spell_power points (11.45 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.7 spell_power points (18.00 DPS) | yes | Magician's Cord (272393, -5.04 DPS) [vendor]; Belt of the Archmage (18405, -7.15 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -9.22 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.7 spell_power points (17.33 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -2.99 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -3.69 DPS) [vendor] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.1 spell_power points (11.44 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.43 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 46.0 spell_power points (15.44 DPS) | yes | Maiden's Circle (13001, -7.48 DPS) [world_drop]; Channeler's Ring (272406, -8.01 DPS) [vendor]; Blessed Band of Light (272407, -8.51 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 31.0 spell_power points (10.40 DPS) | yes | Maiden's Circle (13001, -2.44 DPS) [world_drop]; Channeler's Ring (272406, -2.98 DPS) [vendor]; Blessed Band of Light (272407, -3.47 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (498.6 DPS) | yes | Frozen Heart of the Mountain (249469, -1.48 DPS) [crafted]; Uther's Strength (11302, -2.68 DPS) [world_drop]; Weakness Analyzer (272438, -2.83 DPS, sim-verified) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (498.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (498.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -1.01 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -37.44 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 234.9 spell_power points (78.77 DPS) | yes | Brilliant Wand (249385, -14.25 DPS) [crafted]; Stormrager (16997, -14.71 DPS) [quest]; Wand of Biting Cold (19108, -14.77 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Serenity Field; trinket2: Talisman of Ascendance; ranged: Torch of Light

No-known-source sample (15 of 886, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 153005113100011531-03000000000000000-0545000300000000000)

Set DPS (verified): 737.7. Weights run: 1.3s. Verify run: 2.1s. 886 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.138 ± 0.005, crit=0.465 ± 0.010 per rating point (14 rating = 1%, 6.515 per %), hit=0.958 ± 0.033 per rating point (10 rating = 1%, 9.576 per %), spell_haste=4.730 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.892 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.0 spell_power points (15.80 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.49 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.03 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.66 DPS) [quest]; Jewel of Kajaro (19601, -5.41 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.6 spell_power points (18.16 DPS) | yes | Field Marshal's Silk Spaulders (231602, -4.35 DPS) [pvp]; Argent Shoulders (19059, -5.40 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.84 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.7 spell_power points (13.62 DPS) | yes | Hide of the Wild (18510, -5.77 DPS) [crafted]; Spritecaster Cape (11623, -6.05 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -6.47 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.2 spell_power points (24.58 DPS) | yes | Field Marshal's Silk Vestments (231603, -3.22 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.34 DPS) [vendor]; Robe of the Magi (1716, -15.89 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.1 spell_power points (11.79 DPS) | yes | Runecloth Cuffs (254123, -5.47 DPS) [crafted]; Arcane Runed Bracers (4744, -7.20 DPS) [quest]; Spidertank Oilrag (9448, -7.20 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (14.13 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.6 spell_power points (19.19 DPS) | yes | Belt of the Archmage (18405, -4.54 DPS) [crafted]; Magician's Cord (272393, -7.81 DPS) [vendor]; Highlander's Cloth Girdle (20047, -8.30 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.4 spell_power points (23.68 DPS) | yes | Marshal's Silk Leggings (231605, -3.64 DPS) [pvp]; Sorcerer's Leggings (226933, -6.66 DPS, sim-verified) [quest]; Knight-Captain's Silk Legguards (227109, -8.45 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.25 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.40 DPS) [quest]; Sorcerer's Sandals (226931, -0.40 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 38.8 spell_power points (19.81 DPS) | yes | Elemental Focus Band (20682, -7.81 DPS) [world]; Maiden's Circle (13001, -10.13 DPS) [world_drop]; Don Julio's Band (19325, -11.60 DPS) [rep] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (737.7 DPS) | yes | Maiden's Circle (13001, -0.31 DPS) [world_drop]; Don Julio's Band (19325, -1.78 DPS) [rep]; Elemental Focus Band (20682, -8.34 DPS, sim-verified) [world] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.75 DPS) [crafted]; Weakness Analyzer (272438, -3.34 DPS, sim-verified) [vendor]; Uther's Strength (11302, -4.08 DPS) [world_drop] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.18 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -45.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 156.4 spell_power points (79.82 DPS) | yes | Brilliant Wand (249385, -15.16 DPS) [crafted]; Wand of Biting Cold (19108, -15.82 DPS) [quest]; Stormrager (16997, -16.78 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket1: Serenity Field; trinket2: Talisman of Ascendance; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 886, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (orc, 153002000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 27.9. Weights run: 0.8s. Verify run: 0.6s. 136 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.353 ± 0.008, crit=0.077 ± 0.002 per rating point (14 rating = 1%, 1.079 per %), hit=0.275 ± 0.009 per rating point (10 rating = 1%, 2.754 per %), spell_haste=0.464 ± 0.089, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.465 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.57 DPS) | yes | Shadow Goggles (4373, -0.65 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Roadwatcher's Confidence (281265, -0.56 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.2 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.39 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.38 DPS) | yes | Pearl-clasped Cloak (5542, -0.09 DPS) [crafted]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Battle Healer's Cloak (20427, -0.09 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.64 DPS) | yes | Gray Woolen Robe (2585, -0.09 DPS) [crafted]; Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.26 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.28 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS, sim-verified) [quest]; Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Owlbeard Bracers (16981, -0.12 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.66 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.18 DPS) [crafted]; Apothecary Gloves (10919, -0.28 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.51 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.24 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.34 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk-threaded Trousers (1929, -0.11 DPS) [dungeon]; Rumpled Kilt (274741, -0.29 DPS) [vendor]; Abomination Skin Leggings (23173, -0.72 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.79 DPS) | yes | Pristine Boots (253889, -0.41 DPS) [crafted]; Red Woolen Boots (4313, -0.42 DPS) [crafted]; Woolen Boots (2583, -0.54 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.47 DPS) | yes | Loop of Sacrifice (281673, -0.31 DPS) [quest]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.44 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.1 spell_power points (0.20 DPS) | yes | Loop of Sacrifice (281673, -0.03 DPS) [quest]; Volcanic Rock Ring (12053, -0.10 DPS) [world_drop]; The 1 Ring (8350, -0.17 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.5 spell_power points (0.33 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.13 DPS) [world]; Clear Crystal Rod (16894, -0.17 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 239.6 spell_power points (22.59 DPS) | yes | Skycaller (12984, -0.23 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.47 DPS) [quest]; Greater Magic Wand (11288, -4.90 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 136, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (orc, 153005113100010000-00000000000000000-0000000000000000000)

Set DPS (verified): 112.8. Weights run: 1.0s. Verify run: 0.7s. 242 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.407 ± 0.015, crit=0.067 ± 0.002 per rating point (14 rating = 1%, 0.936 per %), hit=0.266 ± 0.027 per rating point (10 rating = 1%, 2.663 per %), spell_haste=not significant (0.544 ± 0.229), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.900 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.94 DPS) | yes | Silk Headband (7050, -0.53 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.80 DPS) [crafted]; Enchanter's Cowl (4322, -4.29 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (2.52 DPS) | yes | Darkspear Warding Pendant (272075, -1.98 DPS) [vendor]; Crystal Starfire Medallion (5003, -2.09 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.34 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.7 spell_power points (3.38 DPS) | yes | Death Speaker Mantle (6685, -0.58 DPS) [dungeon]; Fairywing Mantle (9536, -0.80 DPS) [quest]; Invoker's Mantle (215365, -0.97 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.34 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Battle Healer's Cloak (19529, -0.27 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.3 spell_power points (3.82 DPS) | yes | Beguiler Robes (7728, -0.11 DPS) [dungeon]; Tree Bark Jacket (1486, -0.35 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.38 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.41 DPS) | yes | Tabitha's Cuffs (251486, -1.60 DPS) [quest]; Nightsky Wristbands (6407, -1.75 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.28 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.0 spell_power points (2.15 DPS) | yes | Serpent Gloves (5970, -0.28 DPS) [dungeon]; Truefaith Gloves (7049, -0.48 DPS) [crafted]; Gnoll Casting Gloves (892, -0.54 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.2 spell_power points (3.27 DPS) | yes | Warsong Sash (16975, -0.33 DPS) [quest]; Belt of Arugal (6392, -0.53 DPS) [dungeon]; Invoker's Cord (215366, -0.85 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.3 spell_power points (3.28 DPS) | yes | Gaze Dreamer Pants (6903, -0.07 DPS) [dungeon]; Pristine Leggings (253987, -0.64 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.02 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.9 spell_power points (2.63 DPS) | yes | Acidic Walkers (9454, -0.43 DPS) [dungeon]; Glinteye Slippers (273024, -0.76 DPS) [dungeon]; Spidersilk Boots (4320, -2.45 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 spell_power points (2.15 DPS) | yes | Ring of Power Regulation (273030, -0.54 DPS) [dungeon]; Advisor's Ring (20426, -0.81 DPS) [rep] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (112.8 DPS) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Advisor's Ring (20426, -0.27 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (9.62 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -8.42 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -8.53 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 124.6 spell_power points (33.30 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.96 DPS) [crafted]; Unstable Power Core (279847, -3.59 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Sea Giant's Toe Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (orc, 153005113100011531-00000000000000000-0000000000000000000)

Set DPS (verified): 176.3. Weights run: 1.0s. Verify run: 0.7s. 316 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.585 ± 0.027, crit=0.172 ± 0.004 per rating point (14 rating = 1%, 2.410 per %), hit=0.398 ± 0.052 per rating point (10 rating = 1%, 3.978 per %), spell_haste=not significant (1.325 ± 0.390), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.918 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.49 DPS) | yes | Papal Fez (9431, -0.94 DPS) [dungeon]; Augural Shroud (2620, -1.28 DPS) [world]; Living Cowl (5608, -2.47 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (3.61 DPS) | yes | Triune Amulet (7722, -1.14 DPS) [dungeon]; Scorn's Icy Choker (23169, -1.64 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -2.34 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.6 spell_power points (4.51 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.54 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.9 spell_power points (2.76 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.59 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.77 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.5 spell_power points (7.88 DPS) | yes | Dreamweave Vest (10021, -0.69 DPS) [crafted]; Robe of Power (7054, -1.39 DPS) [crafted]; Elemental Raiment (9434, -1.39 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.78 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.62 DPS) [quest]; Radiant Silver Bracers (4545, -1.88 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (6.28 DPS) | yes | Black Mageweave Gloves (10003, -1.65 DPS) [crafted]; Red Mageweave Gloves (10018, -2.46 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.55 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.3 spell_power points (5.05 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Star Belt (4329, -1.03 DPS) [crafted]; Gilded Cord (254037, -1.13 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.0 spell_power points (6.49 DPS) | yes | Abomination Skin Leggings (23173, -2.27 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.54 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.79 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.41 DPS) | yes | Gilded Slippers (254001, -2.77 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.42 DPS) [dungeon]; Spidersilk Boots (4320, -4.53 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (3.99 DPS) | yes | Snake Hoop (6750, -1.24 DPS) [quest]; Reedknot Ring (9622, -1.83 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.14 DPS) [vendor] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.78 DPS) | yes | Snake Hoop (6750, -0.02 DPS) [quest]; Reedknot Ring (9622, -0.62 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.93 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (176.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.36 DPS) [dungeon]; Spine of the Devourer (273040, -3.92 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (4.56 DPS) | yes | Orb of the Forgotten Seer (7685, -0.23 DPS) [dungeon]; Thrash's Trash (276204, -0.23 DPS) [vendor]; Orb of Mystic Insight (249394, -1.31 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 131.0 spell_power points (40.46 DPS) | yes | Nether Force Wand (11263, -2.59 DPS) [quest]; Icefury Wand (7514, -2.73 DPS) [quest]; Ragefire Wand (7513, -2.78 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 316, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 153005113100011531-03202300000000000-0000000000000000000)

Set DPS (verified): 284.8. Weights run: 1.0s. Verify run: 0.9s. 405 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.713 ± 0.039, crit=0.264 ± 0.006 per rating point (14 rating = 1%, 3.692 per %), hit=0.676 ± 0.080 per rating point (10 rating = 1%, 6.757 per %), spell_haste=not significant (1.333 ± 0.597), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.921 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 33.3 spell_power points (10.49 DPS) | yes | Dreamweave Circlet (10041, -1.62 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.98 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.21 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (4.01 DPS) | yes | Scorn's Icy Choker (23169, -0.45 DPS) [dungeon]; Mindburst Medallion (11196, -0.77 DPS) [quest]; Horizon Choker (13085, -0.86 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.8 spell_power points (8.15 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -2.44 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.69 DPS, sim-verified) [crafted]; Black Mageweave Shoulders (10027, -2.97 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.4 spell_power points (5.81 DPS) | yes | Spritecaster Cape (11623, -0.04 DPS) [dungeon]; Runecloth Cloak (13860, -1.17 DPS) [crafted]; Big Voodoo Cloak (8216, -2.21 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.3 spell_power points (10.49 DPS) | yes | Robe of the Magi (1716, -2.20 DPS) [world_drop]; Runecloth Tunic (13857, -2.65 DPS) [crafted]; Dreamweave Vest (10021, -2.79 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.0 spell_power points (3.78 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -3.34 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 28.7 spell_power points (9.06 DPS) | yes | Raider Handwraps (272098, -1.91 DPS) [vendor]; Dreamweave Gloves (10019, -2.49 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.94 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 21.1 spell_power points (6.66 DPS) | yes | Dawnspire Cord (12466, -0.50 DPS) [dungeon]; Deathmage Sash (10771, -1.08 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.35 DPS) [rep] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 24.3 spell_power points (7.65 DPS) | yes | Red Mageweave Pants (10009, -0.53 DPS) [crafted]; Wizardweave Leggings (14132, -1.66 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.20 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.57 DPS) | yes | Gilded Sandals (254107, -2.07 DPS) [crafted]; Black Mageweave Boots (10026, -2.52 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.93 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.6 spell_power points (4.28 DPS) | yes | Advisor's Ring (19519, -0.49 DPS) [rep]; Snake Hoop (6750, -1.26 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.10 DPS) | yes | Advisor's Ring (19519, -0.32 DPS) [rep]; Snake Hoop (6750, -1.08 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (284.8 DPS) | yes | Rune of the Guard Captain (19120, -0.43 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (284.8 DPS) | yes | Rune of the Guard Captain (19120, -2.68 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (284.8 DPS) | yes | Kindling Stave (11750, -0.58 DPS) [dungeon]; Spire of Hakkar (10844, -2.65 DPS) [world]; Hanzo Sword (8190, -33.65 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 166.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.23 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Stone Guard's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 405, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 153005113100011531-03202300000000000-0550000000000000000)

Set DPS (verified): 482.9. Weights run: 1.1s. Verify run: 1.8s. 875 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.820 ± 0.046, crit=0.411 ± 0.009 per rating point (14 rating = 1%, 5.750 per %), hit=1.066 ± 0.109 per rating point (10 rating = 1%, 10.661 per %), spell_haste=not significant (2.031 ± 0.783), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.913 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.2 spell_power points (15.84 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.92 DPS) [pvp]; Magister's Crown (16686, -7.29 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (482.9 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.95 DPS) [quest]; Chains of the Lich (23125, -1.23 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.0 spell_power points (15.10 DPS) | yes | Warlord's Silk Amice (231594, -2.60 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.90 DPS) [crafted]; Darkspear Shoulderpads (272103, -8.26 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.2 spell_power points (11.14 DPS) | yes | Hide of the Wild (18510, -3.70 DPS) [crafted]; Deep Woodlands Cloak (19121, -4.64 DPS) [quest]; Spritecaster Cape (11623, -4.79 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.6 spell_power points (18.64 DPS) | yes | Warlord's Silk Raiment (231596, -0.97 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -5.00 DPS) [pvp]; Sorcerer's Robes (226932, -12.81 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.6 spell_power points (9.58 DPS) | yes | Runecloth Cuffs (254123, -3.14 DPS) [crafted]; General's Silk Cuffs (16538, -4.35 DPS) [pvp]; Sorcerer's Bindings (226929, -4.48 DPS) [quest] |
| hands | Sorcerer's Gloves (22066) (or Sorcerer's Gauntlets (226930)) | Just Compensation [quest] | 34.1 spell_power points (11.45 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gauntlets (226930, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 53.7 spell_power points (18.00 DPS) | yes | Magician's Cord (272393, -5.04 DPS) [vendor]; Belt of the Archmage (18405, -6.08 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -9.22 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.7 spell_power points (17.33 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -2.99 DPS) [quest]; Outrider's Silk Leggings (22747, -7.51 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 34.1 spell_power points (11.44 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.43 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 46.0 spell_power points (15.44 DPS) | yes | Maiden's Circle (13001, -7.48 DPS) [world_drop]; Channeler's Ring (272406, -8.01 DPS) [vendor]; Blessed Band of Light (272407, -8.51 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 31.0 spell_power points (10.40 DPS) | yes | Maiden's Circle (13001, -2.44 DPS) [world_drop]; Channeler's Ring (272406, -2.98 DPS) [vendor]; Blessed Band of Light (272407, -3.47 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (482.9 DPS) | yes | Frozen Heart of the Mountain (249469, -1.48 DPS) [crafted]; Rune of the Guard Captain (19120, -2.19 DPS) [quest]; Weakness Analyzer (272438, -2.77 DPS, sim-verified) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (482.9 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (482.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -1.01 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -36.81 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 234.9 spell_power points (78.77 DPS) | yes | Brilliant Wand (249385, -14.25 DPS) [crafted]; Stormrager (16997, -14.71 DPS) [quest]; Wand of Biting Cold (19108, -14.77 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sorcerer's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Serenity Field; trinket2: Talisman of Ascendance; ranged: Torch of Light

No-known-source sample (15 of 875, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 153005113100011531-03000000000000000-0545000300000000000)

Set DPS (verified): 721.0. Weights run: 1.3s. Verify run: 1.9s. 875 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.138 ± 0.005, crit=0.465 ± 0.010 per rating point (14 rating = 1%, 6.515 per %), hit=0.958 ± 0.033 per rating point (10 rating = 1%, 9.576 per %), spell_haste=4.730 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.892 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.0 spell_power points (15.80 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.49 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.12 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.66 DPS) [quest]; Jewel of Kajaro (19601, -5.68 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.6 spell_power points (18.16 DPS) | yes | Warlord's Silk Amice (231594, -4.35 DPS) [pvp]; Argent Shoulders (19059, -5.40 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.60 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.7 spell_power points (13.62 DPS) | yes | Hide of the Wild (18510, -5.77 DPS) [crafted]; Spritecaster Cape (11623, -6.05 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -6.47 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.2 spell_power points (24.58 DPS) | yes | Warlord's Silk Raiment (231596, -3.22 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -9.34 DPS) [pvp]; Robe of the Magi (1716, -15.69 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.1 spell_power points (11.79 DPS) | yes | Runecloth Cuffs (254123, -5.47 DPS) [crafted]; Spidertank Oilrag (9448, -7.20 DPS) [dungeon]; Sorcerer's Bindings (226929, -7.51 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.7 spell_power points (14.13 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 37.6 spell_power points (19.19 DPS) | yes | Belt of the Archmage (18405, -5.50 DPS, sim-verified) [crafted]; Magician's Cord (272393, -7.81 DPS) [vendor]; Defiler's Cloth Girdle (20163, -8.30 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.4 spell_power points (23.68 DPS) | yes | General's Silk Trousers (231595, -3.64 DPS) [pvp]; Sorcerer's Leggings (226933, -6.92 DPS, sim-verified) [quest]; Outrider's Silk Leggings (22747, -8.06 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.25 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Sorcerer's Boots (22064, -0.40 DPS) [quest]; Sorcerer's Sandals (226931, -0.40 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 38.8 spell_power points (19.81 DPS) | yes | Elemental Focus Band (20682, -7.81 DPS) [world]; Maiden's Circle (13001, -10.13 DPS) [world_drop]; Don Julio's Band (19325, -11.60 DPS) [rep] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (721.0 DPS) | yes | Maiden's Circle (13001, -0.31 DPS) [world_drop]; Don Julio's Band (19325, -1.78 DPS) [rep]; Elemental Focus Band (20682, -8.29 DPS, sim-verified) [world] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.75 DPS) [crafted]; Weakness Analyzer (272438, -3.29 DPS, sim-verified) [vendor]; Rune of the Guard Captain (19120, -3.72 DPS) [quest] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Darkmoon Card: Blue Dragon (19288, -5.10 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.18 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -45.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 156.4 spell_power points (79.82 DPS) | yes | Brilliant Wand (249385, -15.16 DPS) [crafted]; Wand of Biting Cold (19108, -15.82 DPS) [quest]; Stormrager (16997, -16.78 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket1: Serenity Field; trinket2: Talisman of Ascendance; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 875, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

