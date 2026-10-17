# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 0000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 31.7. Weights run: 0.6s. Verify run: 0.4s. 145 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.365 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.008 per %), hit=0.238 ± 0.007 per rating point (10 rating = 1%, 2.381 per %), spell_haste=not significant (0.287 ± 0.111), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.63 DPS) | yes | Shadow Goggles (4373, -1.14 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.3 spell_power points (0.87 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.45 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.71 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.42 DPS) | yes | Pearl-clasped Cloak (5542, -0.10 DPS) [crafted]; Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Caretaker's Cape (20428, -0.11 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.72 DPS) | yes | Green Woolen Robe (6243, -0.29 DPS) [crafted]; Green Woolen Vest (2582, -0.30 DPS) [crafted]; Gray Woolen Robe (2585, -0.98 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.8 spell_power points (0.19 DPS) | yes | Windsong Bangles (263336, -0.09 DPS) [quest]; Repurposed Hair Band (281256, -0.12 DPS) [quest]; Bright Bracers (3647, -0.55 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.74 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Heavy Woolen Gloves (4310, -0.45 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.57 DPS) | yes | Novice Ardent's Sash (253887, -0.25 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.76 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.9 spell_power points (1.25 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.52 DPS) [dungeon]; Rumpled Kilt (274741, -0.73 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.89 DPS) | yes | Pristine Boots (253889, -0.37 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.47 DPS) [crafted]; Woolen Boots (2583, -0.60 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.60 DPS) | yes | Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.49 DPS) [world_drop]; The 1 Ring (8350, -0.56 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.53 DPS) | yes | Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.6 spell_power points (0.38 DPS) | yes | Witching Stave (1484, -0.19 DPS) [dungeon]; Staff of Westfall (2042, -0.19 DPS) [quest]; Lesser Staff of the Spire (1300, -0.28 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 215.1 spell_power points (22.63 DPS) | yes | Skycaller (12984, -1.01 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.10 DPS) [quest]; Sizzle Stick (8071, -4.45 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 145, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 0000000000000000000-00000000000000000-2535101301000000000)

Set DPS (verified): 48.1. Weights run: 0.8s. Verify run: 0.5s. 259 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.154 ± 0.004, crit=0.099 ± 0.003 per rating point (14 rating = 1%, 1.390 per %), hit=0.248 ± 0.006 per rating point (10 rating = 1%, 2.483 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.27 DPS) | yes | Filigreed Pristine Circlet (253975, -0.35 DPS) [crafted]; Enchanter's Cowl (4322, -0.40 DPS) [crafted]; Silk Headband (7050, -1.93 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (0.91 DPS) | yes | Darkspear Warding Pendant (272075, -0.82 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.97 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (1.19 DPS) | yes | Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Invoker's Mantle (215365, -0.83 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.58 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Prelacy Cape (7004, -0.12 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.02 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.13 DPS) [dungeon]; Tree Bark Jacket (1486, -0.89 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Windsong Bangles (263336, -0.92 DPS) [quest]; Nightsky Wristbands (6407, -0.93 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.95 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (0.81 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.12 DPS) [world]; Truefaith Gloves (7049, -0.18 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (1.32 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.40 DPS) [dungeon]; Invoker's Cord (215366, -0.42 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.25 DPS) [crafted]; Silk-threaded Trousers (1929, -0.37 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.51 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (0.93 DPS) | yes | Glinteye Slippers (273024, -0.12 DPS) [dungeon]; Acidic Walkers (9454, -0.21 DPS) [dungeon]; Spidersilk Boots (4320, -1.36 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.81 DPS) | yes | Ring of Power Regulation (273030, -0.12 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.12 DPS) [vendor]; Minor Channeling Ring (1449, -0.19 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (0.78 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Ring of Power Regulation (273030, -0.09 DPS) [dungeon]; Minor Channeling Ring (1449, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.14 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.84 DPS) [quest]; Glimmering Staff (249392, -3.95 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 286.3 spell_power points (32.93 DPS) | yes | Wand of Mana Concentration (273043, -0.86 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.36 DPS) [crafted]; Gravestone Scepter (7001, -3.93 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 259, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 0000000000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 84.8. Weights run: 0.8s. Verify run: 0.5s. 338 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.600 ± 0.023, crit=0.174 ± 0.007 per rating point (14 rating = 1%, 2.437 per %), hit=0.399 ± 0.024 per rating point (10 rating = 1%, 3.987 per %), spell_haste=1.674 ± 0.374, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.16 DPS) [world]; Living Cowl (5608, -0.71 DPS) [world]; Spellpower Goggles Xtreme (10502, -2.29 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (1.61 DPS) | yes | Triune Amulet (7722, -0.52 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.03 DPS) [vendor]; Scorn's Icy Choker (23169, -1.12 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (2.02 DPS) | yes | Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Berylline Pads (4197, -0.25 DPS) [quest]; Green Silken Shoulders (7057, -1.07 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.0 spell_power points (1.23 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.27 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.33 DPS) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.30 DPS) [crafted]; Elemental Raiment (9434, -0.33 DPS) [world_drop]; Robe of the Magi (1716, -0.87 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Condor Bracers (15864, -0.27 DPS) [quest]; Windchaser Cuffs (14429, -0.49 DPS) [world_drop]; Arcane Runed Bracers (4744, -1.27 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (2.78 DPS) | yes | Red Mageweave Gloves (10018, -0.46 DPS) [crafted]; Black Mageweave Gloves (10003, -0.74 DPS) [crafted]; Gilded Handwraps (254021, -1.12 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.41 DPS) [crafted]; Gilded Cord (254037, -0.44 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.89 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (2.89 DPS) | yes | Crimson Silk Pantaloons (7062, -0.74 DPS) [crafted]; Abomination Skin Leggings (23173, -1.01 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.25 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.27 DPS) | yes | Gilded Slippers (254001, -1.74 DPS) [crafted]; Acidic Walkers (9454, -1.93 DPS) [dungeon]; Spidersilk Boots (4320, -1.99 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (1.77 DPS) | yes | Lorekeeper's Ring (19524, -0.54 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.68 DPS) [quest]; Reedknot Ring (9622, -0.82 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.23 DPS) | yes | Lorekeeper's Ring (19524, -0.00 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.14 DPS) [quest]; Reedknot Ring (9622, -0.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.16 DPS) [dungeon]; Spine of the Devourer (273040, -1.74 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (2.01 DPS) | yes | Orb of the Forgotten Seer (7685, -0.11 DPS) [dungeon]; Orb of Lorica (11262, -0.38 DPS) [quest]; Orb of Mystic Insight (249394, -0.57 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 292.9 spell_power points (39.87 DPS) | yes | Nether Force Wand (11263, -2.30 DPS) [quest]; Icefury Wand (7514, -2.44 DPS) [quest]; Ragefire Wand (7513, -2.49 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 338, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 2030005000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 203.5. Weights run: 0.7s. Verify run: 0.6s. 431 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.888 ± 0.038, crit=0.273 ± 0.011 per rating point (14 rating = 1%, 3.822 per %), hit=0.842 ± 0.050 per rating point (10 rating = 1%, 8.423 per %), spell_haste=7.404 ± 0.691, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.8 spell_power points (7.10 DPS) | yes | Dreamweave Circlet (10041, -1.33 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.60 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.89 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.1 spell_power points (2.72 DPS) | yes | Horizon Choker (13085, -0.32 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.34 DPS) [dungeon]; Mindburst Medallion (11196, -0.54 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.0 spell_power points (5.60 DPS) | yes | Red Mageweave Shoulders (10029, -1.67 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.77 DPS) [vendor]; Inquisitor's Shawl (19507, -2.02 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.3 spell_power points (3.73 DPS) | yes | Runecloth Cloak (13860, -0.62 DPS) [crafted]; Big Voodoo Cloak (8216, -1.22 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.33 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.8 spell_power points (7.10 DPS) | yes | Robe of the Magi (1716, -1.82 DPS) [world_drop]; Runecloth Tunic (13857, -1.93 DPS) [crafted]; Runecloth Robe (13858, -2.06 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 13.2 spell_power points (2.55 DPS) | yes | Bloodband Bracers (11469, -0.04 DPS) [quest]; Forgotten Wraps (9433, -0.49 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.49 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 32.9 spell_power points (6.34 DPS) | yes | Raider Handwraps (272098, -1.22 DPS) [vendor]; Dreamweave Gloves (10019, -2.18 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.29 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.9 spell_power points (4.42 DPS) | yes | Dawnspire Cord (12466, -0.00 DPS) [dungeon]; Deathmage Sash (10771, -0.49 DPS) [dungeon]; Runecloth Belt (13856, -0.82 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 26.5 spell_power points (5.11 DPS) | yes | Red Mageweave Pants (10009, -0.35 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.34 DPS) [crafted]; Wizardweave Leggings (14132, -1.44 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 24.4 spell_power points (4.71 DPS) | yes | Earthen Silk Slippers (254013, -0.08 DPS) [crafted]; Gilded Sandals (254107, -1.05 DPS) [crafted]; Southsea Mojo Boots (20641, -1.28 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (2.79 DPS) | yes | Snake Hoop (6750, -0.77 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.24 DPS) [quest]; Band of the Unicorn (7553, -3.21 DPS, sim-verified) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-verified (203.5 DPS) | yes | Snake Hoop (6750, -0.30 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.77 DPS) [quest]; Band of the Unicorn (7553, -2.94 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -3.56 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.73 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.23 DPS) [dungeon]; Hanzo Sword (8190, -26.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 272.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.53 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.21 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 431, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 2030005000000000000-11302300000000000-2535101301000300250)

Set DPS (verified): 320.8. Weights run: 0.8s. Verify run: 1.8s. 886 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.653 ± 0.045, crit=0.382 ± 0.013 per rating point (14 rating = 1%, 5.344 per %), hit=1.191 ± 0.066 per rating point (10 rating = 1%, 11.905 per %), spell_haste=7.468 ± 0.785, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.7 spell_power points (8.72 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.93 DPS) [pvp]; Red Mageweave Headband (10033, -7.68 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (320.8 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.30 DPS) [quest]; Chains of the Lich (23125, -0.30 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 42.1 spell_power points (8.61 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.50 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.31 DPS) [crafted]; Darkspear Shoulderpads (272103, -2.56 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.1 spell_power points (6.77 DPS) | yes | Hide of the Wild (18510, -2.57 DPS) [crafted]; Spritecaster Cape (11623, -3.11 DPS) [dungeon]; Runecloth Cloak (13860, -3.86 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 53.2 spell_power points (10.87 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.76 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.22 DPS) [pvp]; Sorcerer's Robes (226932, -7.17 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 27.2 spell_power points (5.56 DPS) | yes | Runecloth Cuffs (254123, -1.98 DPS) [crafted]; Sorcerer's Bindings (226929, -2.80 DPS) [quest]; Bracers of Hope (22667, -2.87 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (320.8 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, -4.46 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 51.3 spell_power points (10.48 DPS) | yes | Magician's Cord (272393, -3.39 DPS) [vendor]; Stormpike Cloth Girdle (19094, -5.46 DPS) [rep]; Belt of the Archmage (18405, -6.37 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.2 spell_power points (10.06 DPS) | yes | Marshal's Silk Leggings (231605, -0.17 DPS) [pvp]; Sorcerer's Leggings (226933, -1.61 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -2.41 DPS) [pvp] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (320.8 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.74 DPS) [dungeon]; Sorcerer's Boots (22064, -5.89 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.8 spell_power points (9.36 DPS) | yes | Maiden's Circle (13001, -4.74 DPS) [world_drop]; Blessed Band of Light (272407, -4.88 DPS) [vendor]; Channeler's Ring (272406, -5.06 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 29.2 spell_power points (5.96 DPS) | yes | Maiden's Circle (13001, -1.35 DPS) [world_drop]; Blessed Band of Light (272407, -1.49 DPS) [vendor]; Channeler's Ring (272406, -1.66 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (320.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (320.8 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.67 DPS) [crafted]; Uther's Strength (11302, -1.63 DPS) [world_drop] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (320.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.24 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -31.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 381.6 spell_power points (77.99 DPS) | yes | Wand of Biting Cold (19108, -6.52 DPS, sim-verified) [quest]; Stormrager (16997, -14.63 DPS) [quest]; Brilliant Wand (249385, -15.63 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 886, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 2030005000000000000-13102000000000000-2555100321000300250)

Set DPS (verified): 547.0. Weights run: 0.8s. Verify run: 1.2s. 886 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.191 ± 0.016, crit=0.420 ± 0.013 per rating point (14 rating = 1%, 5.884 per %), hit=0.966 ± 0.041 per rating point (10 rating = 1%, 9.661 per %), spell_haste=7.070 ± 0.277, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.7 spell_power points (12.51 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -0.53 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.34 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (547.0 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.79 DPS) [quest]; Jewel of Kajaro (19601, -5.04 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.7 spell_power points (14.13 DPS) | yes | Field Marshal's Silk Spaulders (231602, -3.12 DPS) [pvp]; Argent Shoulders (19059, -4.25 DPS) [crafted]; Mantle of the Timbermaw (19050, -4.61 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.2 spell_power points (10.74 DPS) | yes | Hide of the Wild (18510, -4.46 DPS) [crafted]; Spritecaster Cape (11623, -4.76 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -5.21 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.2 spell_power points (19.04 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.39 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -7.13 DPS) [pvp]; Robe of the Magi (1716, -12.06 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.5 spell_power points (9.30 DPS) | yes | Runecloth Cuffs (254123, -4.20 DPS) [crafted]; Arcane Runed Bracers (4744, -5.74 DPS) [quest]; Spidertank Oilrag (9448, -5.74 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.0 spell_power points (11.05 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.9 spell_power points (15.36 DPS) | yes | Belt of the Archmage (18405, -4.50 DPS, sim-verified) [crafted]; Magician's Cord (272393, -6.04 DPS) [vendor]; Highlander's Cloth Girdle (20047, -7.05 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.7 spell_power points (18.05 DPS) | yes | Marshal's Silk Leggings (231605, -2.36 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -6.15 DPS) [pvp]; Sorcerer's Leggings (226933, -7.72 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 24.1 spell_power points (9.51 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.02 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 39.4 spell_power points (15.56 DPS) | yes | Blessed Band of Light (272407, -7.79 DPS) [vendor]; Maiden's Circle (13001, -7.92 DPS) [world_drop]; Don Julio's Band (19325, -9.42 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 24.1 spell_power points (9.52 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Maiden's Circle (13001, -1.88 DPS) [world_drop]; Don Julio's Band (19325, -3.38 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (547.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (547.0 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.10 DPS) [crafted]; Uther's Strength (11302, -3.16 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (547.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.11 DPS) [world]; Teebu's Blazing Longsword (1728, -36.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 200.2 spell_power points (79.13 DPS) | yes | Wand of Biting Cold (19108, -5.17 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.47 DPS) [crafted]; Stormrager (16997, -16.06 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 886, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 0000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 29.4. Weights run: 0.6s. Verify run: 0.4s. 134 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.365 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.008 per %), hit=0.238 ± 0.007 per rating point (10 rating = 1%, 2.381 per %), spell_haste=not significant (0.287 ± 0.111), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.63 DPS) | yes | Shadow Goggles (4373, -1.11 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.3 spell_power points (0.87 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.45 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.60 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.42 DPS) | yes | Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Battle Healer's Cloak (20427, -0.11 DPS) [rep]; Pearl-clasped Cloak (5542, -0.28 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.8 spell_power points (0.72 DPS) | yes | Green Woolen Robe (6243, -0.29 DPS) [crafted]; Green Woolen Vest (2582, -0.30 DPS) [crafted]; Gray Woolen Robe (2585, -0.91 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.32 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Featherbead Bracers (15452, -0.12 DPS) [quest]; Owlbeard Bracers (16981, -0.13 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.74 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.21 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.32 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.57 DPS) | yes | Novice Ardent's Sash (253887, -0.25 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 11.9 spell_power points (1.25 DPS) | yes | Filigreed Pristine Leggings (253937, -0.39 DPS) [crafted]; Silk-threaded Trousers (1929, -0.52 DPS) [dungeon]; Rumpled Kilt (274741, -0.73 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.89 DPS) | yes | Pristine Boots (253889, -0.36 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.47 DPS) [crafted]; Woolen Boots (2583, -0.60 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.53 DPS) | yes | Loop of Sacrifice (281673, -0.33 DPS) [quest]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; The 1 Ring (8350, -0.49 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.2 spell_power points (0.23 DPS) | yes | Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; The 1 Ring (8350, -0.19 DPS) [world]; Loop of Sacrifice (281673, -0.25 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.6 spell_power points (0.38 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.15 DPS) [world]; Clear Crystal Rod (16894, -0.19 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 215.1 spell_power points (22.63 DPS) | yes | Skycaller (12984, -1.16 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.45 DPS) [quest]; Greater Magic Wand (11288, -4.92 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 134, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (troll, 0000000000000000000-00000000000000000-2535101301000000000)

Set DPS (verified): 46.0. Weights run: 0.8s. Verify run: 0.5s. 240 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.154 ± 0.004, crit=0.099 ± 0.003 per rating point (14 rating = 1%, 1.390 per %), hit=0.248 ± 0.006 per rating point (10 rating = 1%, 2.483 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.27 DPS) | yes | Filigreed Pristine Circlet (253975, -0.35 DPS) [crafted]; Enchanter's Cowl (4322, -0.40 DPS) [crafted]; Silk Headband (7050, -1.76 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (0.91 DPS) | yes | Darkspear Warding Pendant (272075, -0.82 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.91 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (1.19 DPS) | yes | Invoker's Mantle (215365, -0.30 DPS) [crafted]; Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Chestnut Mantle (17695, -1.11 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.58 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Battle Healer's Cloak (19529, -0.12 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (46.0 DPS) | yes | Beguiler Robes (7728, -0.02 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.13 DPS) [dungeon]; Tree Bark Jacket (1486, -0.92 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Owlbeard Bracers (16981, -0.88 DPS) [quest]; Glowing Magical Bracelets (13106, -0.89 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.17 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.81 DPS) | yes | Jutebraid Gloves (10654, -0.03 DPS) [quest]; Gnoll Casting Gloves (892, -0.12 DPS) [world]; Truefaith Gloves (7049, -0.18 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (1.32 DPS) | yes | Warsong Sash (16975, -0.05 DPS) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.40 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.38 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.45 DPS) [crafted]; Silk-threaded Trousers (1929, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (0.93 DPS) | yes | Glinteye Slippers (273024, -0.12 DPS) [dungeon]; Acidic Walkers (9454, -0.21 DPS) [dungeon]; Spidersilk Boots (4320, -1.14 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.81 DPS) | yes | Ring of Power Regulation (273030, -0.12 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.12 DPS) [vendor]; Solid Gold Band (279845, -0.35 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (0.78 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Ring of Power Regulation (273030, -0.09 DPS) [dungeon]; Solid Gold Band (279845, -0.32 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.14 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -3.95 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -3.96 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 286.3 spell_power points (32.93 DPS) | yes | Wand of Mana Concentration (273043, -1.06 DPS, sim-verified) [world]; Unstable Power Core (279847, -3.22 DPS) [quest]; Greater Mystic Wand (217287, -3.36 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 240, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 20162 Defiler's Cloth Boots

### Band 40 (troll, 0000000000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 80.1. Weights run: 0.8s. Verify run: 0.5s. 315 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.600 ± 0.023, crit=0.174 ± 0.007 per rating point (14 rating = 1%, 2.437 per %), hit=0.399 ± 0.024 per rating point (10 rating = 1%, 3.987 per %), spell_haste=1.674 ± 0.374, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (80.1 DPS) | yes | Augural Shroud (2620, -0.16 DPS) [world]; Living Cowl (5608, -0.71 DPS) [world]; Spellpower Goggles Xtreme (10502, -1.00 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.8 spell_power points (1.61 DPS) | yes | Triune Amulet (7722, -0.52 DPS) [dungeon]; Scorn's Icy Choker (23169, -0.99 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.03 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.8 spell_power points (2.02 DPS) | yes | Bloodmage Mantle (7684, -0.05 DPS) [dungeon]; Berylline Pads (4197, -0.25 DPS) [quest]; Green Silken Shoulders (7057, -0.94 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.0 spell_power points (1.23 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.27 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.33 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.6 spell_power points (3.49 DPS) | yes | Dreamweave Vest (10021, -0.30 DPS) [crafted]; Robe of Power (7054, -0.60 DPS) [crafted]; Elemental Raiment (9434, -0.63 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.23 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.27 DPS) [quest]; Radiant Silver Bracers (4545, -0.78 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (2.78 DPS) | yes | Black Mageweave Gloves (10003, -0.74 DPS) [crafted]; Gilded Handwraps (254021, -1.12 DPS) [crafted]; Red Mageweave Gloves (10018, -1.20 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.4 spell_power points (2.23 DPS) | yes | Deathmage Sash (10771, -0.05 DPS) [dungeon]; Star Belt (4329, -0.46 DPS) [crafted]; Gilded Cord (254037, -0.49 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.2 spell_power points (2.89 DPS) | yes | Abomination Skin Leggings (23173, -1.01 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.05 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.25 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.27 DPS) | yes | Gilded Slippers (254001, -0.68 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.93 DPS) [dungeon]; Spidersilk Boots (4320, -1.99 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.0 spell_power points (1.77 DPS) | yes | Advisor's Ring (19520, -0.54 DPS) [rep]; Reedknot Ring (9622, -0.82 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.23 DPS) | yes | Advisor's Ring (19520, -0.00 DPS) [rep]; Reedknot Ring (9622, -0.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.16 DPS) [dungeon]; Spine of the Devourer (273040, -1.74 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 14.8 spell_power points (2.01 DPS) | yes | Orb of the Forgotten Seer (7685, -0.11 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.57 DPS) [crafted]; Omega Orb (7749, -0.79 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 292.9 spell_power points (39.87 DPS) | yes | Nether Force Wand (11263, -1.81 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.44 DPS) [quest]; Ragefire Wand (7513, -2.49 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (troll, 2030005000000000000-00000000000000000-2535101301000300250)

Set DPS (verified): 193.3. Weights run: 0.7s. Verify run: 0.6s. 404 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.888 ± 0.038, crit=0.273 ± 0.011 per rating point (14 rating = 1%, 3.822 per %), hit=0.842 ± 0.050 per rating point (10 rating = 1%, 8.423 per %), spell_haste=7.404 ± 0.691, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.8 spell_power points (7.10 DPS) | yes | Dreamweave Circlet (10041, -1.33 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.60 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.89 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.1 spell_power points (2.72 DPS) | yes | Horizon Choker (13085, -0.32 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.34 DPS) [dungeon]; Mindburst Medallion (11196, -0.54 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.0 spell_power points (5.60 DPS) | yes | Red Mageweave Shoulders (10029, -1.67 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.77 DPS) [vendor]; Inquisitor's Shawl (19507, -2.02 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.0 spell_power points (3.86 DPS) | yes | Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Runecloth Cloak (13860, -0.75 DPS) [crafted]; Big Voodoo Cloak (8216, -1.35 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.8 spell_power points (7.10 DPS) | yes | Robe of the Magi (1716, -1.82 DPS) [world_drop]; Runecloth Tunic (13857, -1.93 DPS) [crafted]; Runecloth Robe (13858, -2.06 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 13.2 spell_power points (2.55 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.04 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 32.9 spell_power points (6.34 DPS) | yes | Raider Handwraps (272098, -1.22 DPS) [vendor]; Dreamweave Gloves (10019, -2.18 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.29 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.9 spell_power points (4.42 DPS) | yes | Dawnspire Cord (12466, -0.00 DPS) [dungeon]; Deathmage Sash (10771, -0.49 DPS) [dungeon]; Runecloth Belt (13856, -0.82 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 26.5 spell_power points (5.11 DPS) | yes | Red Mageweave Pants (10009, -0.35 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.34 DPS) [crafted]; Wizardweave Leggings (14132, -1.44 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 24.4 spell_power points (4.71 DPS) | yes | Earthen Silk Slippers (254013, -0.08 DPS) [crafted]; Gilded Sandals (254107, -1.05 DPS) [crafted]; Southsea Mojo Boots (20641, -1.28 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (2.79 DPS) | yes | Advisor's Ring (19519, -0.47 DPS) [rep]; Snake Hoop (6750, -0.77 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.51 DPS) | yes | Advisor's Ring (19519, -0.19 DPS) [rep]; Snake Hoop (6750, -0.49 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (193.3 DPS) | yes | Uther's Strength (11302, -0.31 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (193.3 DPS) | yes | Kindling Stave (11750, -0.73 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.23 DPS) [dungeon]; Hanzo Sword (8190, -25.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 272.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.53 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.21 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Satyrmane Sash; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 404, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (troll, 2030005000000000000-11302300000000000-2535101301000300250)

Set DPS (verified): 309.4. Weights run: 0.8s. Verify run: 1.8s. 875 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.653 ± 0.045, crit=0.382 ± 0.013 per rating point (14 rating = 1%, 5.344 per %), hit=1.191 ± 0.066 per rating point (10 rating = 1%, 11.905 per %), spell_haste=7.468 ± 0.785, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 42.7 spell_power points (8.72 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.93 DPS) [pvp]; Red Mageweave Headband (10033, -10.24 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (309.4 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Orb of the Darkmoon (19426, -0.30 DPS) [quest]; Chains of the Lich (23125, -0.30 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 42.1 spell_power points (8.61 DPS) | yes | Warlord's Silk Amice (231594, -1.50 DPS) [pvp]; Darkspear Shoulderpads (272103, -2.56 DPS) [vendor]; Mantle of the Timbermaw (19050, -6.10 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.1 spell_power points (6.77 DPS) | yes | Hide of the Wild (18510, -2.57 DPS) [crafted]; Spritecaster Cape (11623, -3.11 DPS) [dungeon]; Deep Woodlands Cloak (19121, -3.12 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 53.2 spell_power points (10.87 DPS) | yes | Warlord's Silk Raiment (231596, -0.76 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.22 DPS) [pvp]; Sorcerer's Robes (226932, -6.25 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.2 spell_power points (5.56 DPS) | yes | Runecloth Cuffs (254123, -1.98 DPS) [crafted]; Sorcerer's Bindings (226929, -2.80 DPS) [quest]; Bracers of Hope (22667, -2.87 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | sim-verified (309.4 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, -5.85 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 51.3 spell_power points (10.48 DPS) | yes | Magician's Cord (272393, -3.39 DPS) [vendor]; Belt of the Archmage (18405, -5.29 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -5.46 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 49.2 spell_power points (10.06 DPS) | yes | General's Silk Trousers (231595, -0.17 DPS) [pvp]; Sorcerer's Leggings (226933, -1.61 DPS) [quest]; Outrider's Silk Leggings (22747, -1.80 DPS) [rep] |
| feet | Sorcerer's Sandals (226931) | Mokvar [vendor] | sim-verified (309.4 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.74 DPS) [dungeon]; Sorcerer's Boots (22064, -6.35 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.8 spell_power points (9.36 DPS) | yes | Maiden's Circle (13001, -4.74 DPS) [world_drop]; Blessed Band of Light (272407, -4.88 DPS) [vendor]; Channeler's Ring (272406, -5.06 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 29.2 spell_power points (5.96 DPS) | yes | Maiden's Circle (13001, -1.35 DPS) [world_drop]; Blessed Band of Light (272407, -1.49 DPS) [vendor]; Channeler's Ring (272406, -1.66 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (309.4 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (309.4 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.67 DPS) [crafted]; Rune of the Guard Captain (19120, -1.16 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (309.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.24 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -30.54 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 381.6 spell_power points (77.99 DPS) | yes | Wand of Biting Cold (19108, -13.99 DPS) [quest]; Stormrager (16997, -14.63 DPS) [quest]; Brilliant Wand (249385, -15.63 DPS) [crafted] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 875, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (troll, 2030005000000000000-13102000000000000-2555100321000300250)

Set DPS (verified): 547.6. Weights run: 0.8s. Verify run: 1.2s. 875 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.191 ± 0.016, crit=0.420 ± 0.013 per rating point (14 rating = 1%, 5.884 per %), hit=0.966 ± 0.041 per rating point (10 rating = 1%, 9.661 per %), spell_haste=7.070 ± 0.277, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 31.7 spell_power points (12.51 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -0.53 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -4.43 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (547.6 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.79 DPS) [quest]; Jewel of Kajaro (19601, -4.64 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 35.7 spell_power points (14.13 DPS) | yes | Warlord's Silk Amice (231594, -3.12 DPS) [pvp]; Argent Shoulders (19059, -4.25 DPS) [crafted]; Mantle of the Timbermaw (19050, -4.44 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 27.2 spell_power points (10.74 DPS) | yes | Hide of the Wild (18510, -4.46 DPS) [crafted]; Spritecaster Cape (11623, -4.76 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -5.21 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 48.2 spell_power points (19.04 DPS) | yes | Warlord's Silk Raiment (231596, -2.39 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -7.13 DPS) [pvp]; Robe of the Magi (1716, -11.98 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.5 spell_power points (9.30 DPS) | yes | Runecloth Cuffs (254123, -4.20 DPS) [crafted]; Spidertank Oilrag (9448, -5.74 DPS) [dungeon]; Sorcerer's Bindings (226929, -5.78 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.0 spell_power points (11.05 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.9 spell_power points (15.36 DPS) | yes | Belt of the Archmage (18405, -5.25 DPS, sim-verified) [crafted]; Magician's Cord (272393, -6.04 DPS) [vendor]; Defiler's Cloth Girdle (20163, -7.05 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.7 spell_power points (18.05 DPS) | yes | General's Silk Trousers (231595, -2.36 DPS) [pvp]; Outrider's Silk Leggings (22747, -5.55 DPS) [rep]; Sorcerer's Leggings (226933, -7.55 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 24.1 spell_power points (9.51 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.02 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 39.4 spell_power points (15.56 DPS) | yes | Blessed Band of Light (272407, -7.79 DPS) [vendor]; Maiden's Circle (13001, -7.92 DPS) [world_drop]; Don Julio's Band (19325, -9.42 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 24.1 spell_power points (9.52 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Maiden's Circle (13001, -1.88 DPS) [world_drop]; Don Julio's Band (19325, -3.38 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (547.6 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (547.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.10 DPS) [crafted]; Rune of the Guard Captain (19120, -2.86 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (547.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.11 DPS) [world]; Teebu's Blazing Longsword (1728, -36.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 200.2 spell_power points (79.13 DPS) | yes | Wand of Biting Cold (19108, -5.13 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.47 DPS) [crafted]; Stormrager (16997, -16.06 DPS) [quest] |

**New at 60:** head: Sorcerer's Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 875, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

