# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 1531001000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 38.7. Weights run: 1.1s. Verify run: 0.7s. 154 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.294 ± 0.008, crit=0.058 ± 0.002 per rating point (14 rating = 1%, 0.807 per %), hit=0.201 ± 0.007 per rating point (10 rating = 1%, 2.015 per %), spell_haste=0.568 ± 0.100, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.236 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.81 DPS) | yes | Shadow Goggles (4373, -1.84 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.6 spell_power points (1.04 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.49 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.62 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.54 DPS) | yes | Black Whelp Cloak (7283, -0.14 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.15 DPS) [crafted]; Sanguine Cape (14376, -0.38 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.88 DPS) | yes | Green Woolen Vest (2582, -0.33 DPS) [crafted]; Bloody Apron (6226, -0.33 DPS) [dungeon]; Gray Woolen Robe (2585, -1.24 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.5 spell_power points (0.20 DPS) | yes | Windsong Bangles (263336, -0.06 DPS) [quest]; Repurposed Hair Band (281256, -0.12 DPS) [quest]; Bright Bracers (3647, -1.02 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.95 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.29 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.60 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.70 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.38 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.01 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 7.8 spell_power points (1.05 DPS) | yes | Colorful Kilt (10048, -0.37 DPS) [crafted]; Heavy Woolen Pants (4316, -0.65 DPS) [crafted]; Silk-threaded Trousers (1929, -2.01 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (1.11 DPS) | yes | Pristine Boots (253889, -0.58 DPS) [crafted]; Woolen Boots (2583, -0.76 DPS) [crafted]; Red Woolen Boots (4313, -0.96 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.76 DPS) | yes | Volcanic Rock Ring (12053, -0.64 DPS) [world_drop]; The 1 Ring (8350, -0.72 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.8 spell_power points (0.24 DPS) | yes | The 1 Ring (8350, -0.20 DPS) [world]; Volcanic Rock Ring (12053, -0.46 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.9 spell_power points (0.40 DPS) | yes | Witching Stave (1484, -0.20 DPS) [dungeon]; Staff of Westfall (2042, -0.20 DPS) [quest]; Lesser Staff of the Spire (1300, -0.71 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 167.6 spell_power points (22.72 DPS) | yes | Skycaller (12984, -0.96 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.19 DPS) [quest]; Sizzle Stick (8071, -4.39 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 154, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 1531005113100000000-00000000000000000-000000000000000000)

Set DPS (verified): 91.5. Weights run: 1.2s. Verify run: 0.8s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.600 ± 0.019, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.986 per %), hit=0.270 ± 0.028 per rating point (10 rating = 1%, 2.701 per %), spell_haste=-7.141 ± 0.517, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.541 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.38 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.57 DPS) [crafted]; Enchanter's Cowl (4322, -2.41 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.6 spell_power points (0.69 DPS) | yes | Crystal Starfire Medallion (5003, -0.23 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (2.75 DPS) | yes | Fairywing Mantle (9536, -0.57 DPS) [quest]; Magician's Mantle (12998, -0.76 DPS) [world_drop]; Death Speaker Mantle (6685, -1.21 DPS, sim-verified) [dungeon] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.95 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Prelacy Cape (7004, -0.19 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.8 spell_power points (3.21 DPS) | yes | Beguiler Robes (7728, -0.11 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.31 DPS) [dungeon]; Death Speaker Robes (6682, -0.61 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.72 DPS) | yes | Nightsky Wristbands (6407, -1.03 DPS) [world_drop]; Stonecloth Bindings (14416, -1.15 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.03 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.2 spell_power points (1.37 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Shilly Mitts (9609, -0.04 DPS) [quest]; Truefaith Gloves (7049, -0.08 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 10.8 spell_power points (2.06 DPS) | yes | Invoker's Cord (215366, -0.15 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.53 DPS) [dungeon]; Crimson Silk Belt (7055, -1.09 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.31 DPS) [crafted]; Silk-threaded Trousers (1929, -0.80 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.84 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.2 spell_power points (2.14 DPS) | yes | Spidersilk Boots (4320, -0.34 DPS) [crafted]; Glinteye Slippers (273024, -0.80 DPS) [dungeon]; Acidic Walkers (9454, -2.67 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.72 DPS) | yes | Ring of Power Regulation (273030, -0.57 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -1.03 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.14 DPS) [dungeon] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (1.18 DPS) | yes | Ring of Power Regulation (273030, -0.04 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (6.87 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.05 DPS) [quest]; Glimmering Staff (249392, -5.61 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 174.6 spell_power points (33.32 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.36 DPS) [crafted]; Gravestone Scepter (7001, -4.32 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Minor Channeling Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 1531005113100001531-00000000000000000-000000000000000000)

Set DPS (verified): 141.0. Weights run: 1.3s. Verify run: 0.9s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.817 ± 0.043, crit=0.160 ± 0.005 per rating point (14 rating = 1%, 2.238 per %), hit=0.553 ± 0.057 per rating point (10 rating = 1%, 5.534 per %), spell_haste=-14.786 ± 1.089, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.570 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 21.9 spell_power points (4.46 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.18 DPS) [crafted]; Augural Shroud (2620, -0.55 DPS) [world]; Whitemane's Chapeau (7720, -1.37 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.5 spell_power points (2.76 DPS) | yes | Triune Amulet (7722, -1.13 DPS) [dungeon]; Necklace of Calisea (1714, -1.59 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.76 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.6 spell_power points (3.59 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.50 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.1 spell_power points (2.05 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.63 DPS) [crafted]; Blackforge Cape (6424, -0.72 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.9 spell_power points (5.48 DPS) | yes | Dreamweave Vest (10021, -0.32 DPS) [crafted]; Robe of Power (7054, -0.63 DPS) [crafted]; Elemental Raiment (9434, -1.20 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.83 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Windchaser Cuffs (14429, -0.34 DPS) [world_drop]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (4.33 DPS) | yes | Red Mageweave Gloves (10018, -0.43 DPS) [crafted]; Black Mageweave Gloves (10003, -1.28 DPS) [crafted]; Stormcloth Gloves (10011, -1.52 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.3 spell_power points (3.92 DPS) | yes | Gilded Cord (254037, -0.96 DPS) [crafted]; Star Belt (4329, -1.27 DPS) [crafted]; Crimson Silk Belt (7055, -1.53 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.8 spell_power points (4.85 DPS) | yes | Crimson Silk Pantaloons (7062, -1.06 DPS) [crafted]; Stoneweaver Leggings (9407, -2.09 DPS) [dungeon]; Pristine Leggings (253987, -2.26 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Gilded Slippers (254001, -2.04 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.54 DPS) [dungeon]; Spidersilk Boots (4320, -2.80 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.87 DPS) | yes | Ring of Forlorn Spirits (2043, -1.24 DPS) [quest]; Reedknot Ring (9622, -1.44 DPS) [quest]; Minor Channeling Ring (1449, -1.52 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.1 spell_power points (2.05 DPS) | yes | Ring of Forlorn Spirits (2043, -0.42 DPS) [quest]; Reedknot Ring (9622, -0.63 DPS) [quest]; Minor Channeling Ring (1449, -0.70 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (141.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.33 DPS) [dungeon]; Spine of the Devourer (273040, -2.96 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.5 spell_power points (3.15 DPS) | yes | Orb of Lorica (11262, -0.26 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.30 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.72 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 198.3 spell_power points (40.38 DPS) | yes | Nether Force Wand (11263, -2.55 DPS) [quest]; Icefury Wand (7514, -2.69 DPS) [quest]; Ragefire Wand (7513, -2.74 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 1531005113100001531-03202300000000000-000000000000000000)

Set DPS (verified): 232.4. Weights run: 1.3s. Verify run: 1.1s. 506 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.060 ± 0.065, crit=0.239 ± 0.007 per rating point (14 rating = 1%, 3.346 per %), hit=0.702 ± 0.090 per rating point (10 rating = 1%, 7.021 per %), spell_haste=-21.271 ± 1.665, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.585 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.2 spell_power points (7.90 DPS) | yes | Knight-Lieutenant's Dreadweave Hat (220889, -1.99 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.59 DPS) [crafted]; Dreamweave Circlet (10041, -2.80 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.5 spell_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.13 DPS) [world_drop]; Mindburst Medallion (11196, -0.61 DPS) [quest]; Triune Amulet (7722, -1.47 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.1 spell_power points (6.31 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -2.20 DPS) [vendor]; Inquisitor's Shawl (19507, -2.22 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.53 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.4 spell_power points (4.00 DPS) | yes | Runecloth Cloak (13860, -0.57 DPS) [crafted]; Big Voodoo Cloak (8216, -1.14 DPS) [crafted]; Imperial Red Cloak (8248, -1.71 DPS) [world_drop] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.2 spell_power points (7.90 DPS) | yes | Runecloth Tunic (13857, -2.27 DPS) [crafted]; Robe of the Magi (1716, -2.33 DPS) [world_drop]; Runecloth Robe (13858, -4.15 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.5 spell_power points (2.86 DPS) | yes | Nethergeld Cuffs (254061, -0.02 DPS) [crafted]; Forgotten Wraps (9433, -0.36 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.36 DPS) [quest] |
| hands | Sergeant Major's Dreadweave Gloves (220890) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 22.5 spell_power points (4.43 DPS) | yes | Dreamweave Gloves (10019, -0.06 DPS) [crafted]; Red Mageweave Gloves (10018, -0.18 DPS) [crafted]; Runecloth Gloves (13863, -0.20 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.1 spell_power points (5.14 DPS) | yes | Deathmage Sash (10771, -0.64 DPS) [dungeon]; Runecloth Belt (13856, -1.07 DPS) [crafted]; Satyrmane Sash (17755, -2.92 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 28.1 spell_power points (5.52 DPS) | yes | Red Mageweave Pants (10009, -0.26 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.18 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.24 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 24.6 spell_power points (4.83 DPS) | yes | Earthen Silk Slippers (254013, -0.11 DPS) [crafted]; Gilded Sandals (254107, -0.79 DPS) [crafted]; Southsea Mojo Boots (20641, -0.96 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.3 spell_power points (3.01 DPS) | yes | Snake Hoop (6750, -0.79 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.44 DPS) [quest]; Voodoo Band (1996, -1.55 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.56 DPS) | yes | Snake Hoop (6750, -0.33 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.98 DPS) [quest]; Voodoo Band (1996, -1.10 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (232.4 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (232.4 DPS) | yes | Mark of the Chosen (17774, -3.75 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (232.4 DPS) | yes | Spellshifter Rod (9527, -0.46 DPS) [quest]; Kindling Stave (11750, -1.58 DPS) [dungeon]; Hanzo Sword (8190, -39.44 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 267.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.52 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.18 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sergeant Major's Dreadweave Gloves; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 506, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 1531005113100001531-03202300000000000-053200000000000000)

Set DPS (verified): 367.8. Weights run: 1.4s. Verify run: 2.9s. 1125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.093 ± 0.070, crit=0.399 ± 0.011 per rating point (14 rating = 1%, 5.591 per %), hit=0.932 ± 0.107 per rating point (10 rating = 1%, 9.322 per %), spell_haste=-30.393 ± 1.876, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.527 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 43.8 spell_power points (10.36 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -9.65 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 22.9 spell_power points (5.43 DPS) | yes | Glowing Eye of Mordresh (10769, -1.70 DPS) [dungeon]; Horizon Choker (13085, -1.81 DPS) [world_drop]; Opaline Medallion (16623, -2.06 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (367.8 DPS) | yes | Elder Wizard's Mantle (13013, +0.00 DPS) [world_drop]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.96 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.6 spell_power points (4.87 DPS) | yes | Runecloth Cloak (13860, -0.67 DPS) [crafted]; Big Voodoo Cloak (8216, -1.35 DPS) [crafted]; Imperial Red Cloak (8248, -2.02 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.7 spell_power points (13.89 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.36 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.20 DPS) [pvp]; Magister's Robes (16688, -6.49 DPS, sim-verified) [dungeon] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (367.8 DPS) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; Marshal's Silk Bracers (16438, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, -5.45 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.1 spell_power points (7.36 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -5.39 DPS, sim-verified) [vendor] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 30.0 spell_power points (7.09 DPS) | yes | Wisdom of the Timbermaw (19047, -0.86 DPS) [crafted]; Devout Belt (16696, -0.97 DPS) [dungeon]; Dawnspire Cord (12466, -8.62 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 31.8 spell_power points (7.52 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -12.95 DPS, sim-verified) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 33.1 spell_power points (7.84 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.28 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.7 spell_power points (6.07 DPS) | yes | Glowing Crystal Ring (18402, -2.97 DPS) [quest]; Band of the Unicorn (7553, -2.99 DPS) [world_drop]; Snake Hoop (6750, -3.36 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (3.66 DPS) | yes | Glowing Crystal Ring (18402, -0.56 DPS) [quest]; Band of the Unicorn (7553, -0.58 DPS) [world_drop]; Snake Hoop (6750, -0.95 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (367.8 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (367.8 DPS) | yes | Mark of the Chosen (17774, -4.57 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (367.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.61 DPS) [quest]; Teebu's Blazing Longsword (1728, -47.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 330.3 spell_power points (78.18 DPS) | yes | Stormrager (16997, -4.23 DPS, sim-verified) [quest]; Brilliant Wand (249385, -14.66 DPS) [crafted]; Torch of Austen (13004, -21.39 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Magister's Mantle; chest: Robe of the Archmage; wrist: Magister's Bindings; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; ranged: Torch of Light

No-known-source sample (15 of 1125, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 1531005113100101531-03000000000000000-053500030000000000)

Set DPS (verified): 576.7. Weights run: 1.5s. Verify run: 2.2s. 1125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.453 ± 0.013, crit=0.465 ± 0.012 per rating point (14 rating = 1%, 6.512 per %), hit=0.996 ± 0.038 per rating point (10 rating = 1%, 9.959 per %), spell_haste=-13.693 ± 0.827, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.475 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 28.1 spell_power points (12.61 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -0.48 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.5 spell_power points (7.43 DPS) | yes | Glowing Eye of Mordresh (10769, -2.65 DPS) [dungeon]; Mindburst Medallion (11196, -3.51 DPS) [quest]; Triune Amulet (7722, -3.83 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.4 spell_power points (13.21 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.31 DPS) [pvp]; Argent Shoulders (19059, -3.35 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.7 spell_power points (7.51 DPS) | yes | Runecloth Cloak (13860, -1.84 DPS) [crafted]; Big Voodoo Cloak (8216, -3.43 DPS) [crafted]; Nightfall Drape (12465, -3.47 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 51.9 spell_power points (23.35 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.13 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -7.52 DPS) [pvp]; Acumen Robes (17775, -12.91 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 15.5 spell_power points (6.98 DPS) | yes | Nethergeld Cuffs (254061, -2.41 DPS) [crafted]; Dreadmist Bracers (16703, -2.49 DPS) [dungeon]; Bloodband Bracers (11469, -2.90 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 26.6 spell_power points (11.93 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -5.17 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 18.5 spell_power points (8.33 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Magister's Belt (16685, -0.91 DPS) [dungeon]; Dreadmist Belt (16702, -1.43 DPS) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.68 DPS) | yes | Knight's Dreadweave Leggings (220888, +0.00 DPS) [vendor]; Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.4 spell_power points (11.43 DPS) | yes | Knight-Lieutenant's Silk Walkers (227112, +0.00 DPS) [pvp]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.65 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.2 spell_power points (9.51 DPS) | yes | Philanthropist's Ring (281635, -4.00 DPS) [quest]; Snake Hoop (6750, -5.80 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.92 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.84 DPS) | yes | Philanthropist's Ring (281635, -0.33 DPS) [quest]; Snake Hoop (6750, -2.13 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.25 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (576.7 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (576.7 DPS) | yes | Mark of the Chosen (17774, -5.79 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (576.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.46 DPS) [quest]; Teebu's Blazing Longsword (1728, -41.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 176.8 spell_power points (79.46 DPS) | yes | Brilliant Wand (249385, -14.41 DPS) [crafted]; Stormrager (16997, -15.75 DPS) [quest]; Greater Eternal Wand (249237, -22.08 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Omnicast Boots; finger1: Maiden's Circle; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1125, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 1531001000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.3. Weights run: 1.1s. Verify run: 0.7s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.294 ± 0.008, crit=0.058 ± 0.002 per rating point (14 rating = 1%, 0.807 per %), hit=0.201 ± 0.007 per rating point (10 rating = 1%, 2.015 per %), spell_haste=0.568 ± 0.100, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.236 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.81 DPS) | yes | Shadow Goggles (4373, -1.70 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.6 spell_power points (1.04 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.20 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.49 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.54 DPS) | yes | Black Whelp Cloak (7283, -0.14 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.15 DPS) [crafted]; Mystic Shawl (3449, -0.27 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.5 spell_power points (0.88 DPS) | yes | Green Woolen Vest (2582, -0.33 DPS) [crafted]; Bloody Apron (6226, -0.33 DPS) [dungeon]; Gray Woolen Robe (2585, -1.01 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.41 DPS) | yes | Owlbeard Bracers (16981, -0.19 DPS) [quest]; Mindthrust Bracers (1974, -0.21 DPS) [dungeon]; Featherbead Bracers (15452, -0.21 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.95 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.29 DPS) [crafted]; Apothecary Gloves (10919, -0.41 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.2 spell_power points (0.70 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.38 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.50 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 7.8 spell_power points (1.05 DPS) | yes | Colorful Kilt (10048, -0.37 DPS) [crafted]; Embroidered Leggings (270011, -0.51 DPS) [quest]; Silk-threaded Trousers (1929, -1.49 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.2 spell_power points (1.11 DPS) | yes | Pristine Boots (253889, -0.58 DPS) [crafted]; Red Woolen Boots (4313, -0.61 DPS, sim-verified) [crafted]; Woolen Boots (2583, -0.76 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.8 spell_power points (0.24 DPS) | yes | Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; The 1 Ring (8350, -0.20 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 1.5 spell_power points (0.20 DPS) | yes | Volcanic Rock Ring (12053, -0.08 DPS) [world_drop]; The 1 Ring (8350, -0.16 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.9 spell_power points (0.40 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.16 DPS) [world]; Clear Crystal Rod (16894, -0.20 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 167.6 spell_power points (22.72 DPS) | yes | Skycaller (12984, -1.16 DPS) [world_drop]; Sizzle Stick (8071, -4.39 DPS) [quest]; Greater Magic Wand (11288, -4.95 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 1531005113100000000-00000000000000000-000000000000000000)

Set DPS (verified): 88.4. Weights run: 1.2s. Verify run: 0.8s. 282 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.600 ± 0.019, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.986 per %), hit=0.270 ± 0.028 per rating point (10 rating = 1%, 2.701 per %), spell_haste=-7.141 ± 0.517, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.541 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.38 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.57 DPS) [crafted]; Enchanter's Cowl (4322, -3.01 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.6 spell_power points (0.69 DPS) | yes | Crystal Starfire Medallion (5003, -0.23 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (2.75 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.57 DPS) [quest]; Magician's Mantle (12998, -0.76 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.23 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.8 spell_power points (3.21 DPS) | yes | Beguiler Robes (7728, -0.11 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.31 DPS) [dungeon]; Death Speaker Robes (6682, -0.61 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.72 DPS) | yes | Nightsky Wristbands (6407, -1.03 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.15 DPS) [quest]; Glowing Magical Bracelets (13106, -1.43 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.0 spell_power points (1.72 DPS) | yes | Serpent Gloves (5970, -0.38 DPS) [dungeon]; Truefaith Gloves (7049, -0.42 DPS) [crafted]; Gnoll Casting Gloves (892, -0.57 DPS) [world] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.10 DPS) | yes | Belt of Arugal (6392, -0.04 DPS) [dungeon]; Crimson Silk Belt (7055, -0.15 DPS) [crafted]; Invoker's Cord (215366, -0.19 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.31 DPS) [crafted]; Silk-threaded Trousers (1929, -0.80 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.56 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.2 spell_power points (2.14 DPS) | yes | Spidersilk Boots (4320, -0.34 DPS) [crafted]; Glinteye Slippers (273024, -0.80 DPS) [dungeon]; Acidic Walkers (9454, -1.67 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.72 DPS) | yes | Solid Gold Band (279845, -0.95 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.03 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.14 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.15 DPS) | yes | Solid Gold Band (279845, -0.38 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (6.87 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -5.61 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -5.73 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 174.6 spell_power points (33.32 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.36 DPS) [crafted]; Unstable Power Core (279847, -3.61 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Warsong Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 282, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (orc, 1531005113100001531-00000000000000000-000000000000000000)

Set DPS (verified): 137.6. Weights run: 1.3s. Verify run: 1.0s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.817 ± 0.043, crit=0.160 ± 0.005 per rating point (14 rating = 1%, 2.238 per %), hit=0.553 ± 0.057 per rating point (10 rating = 1%, 5.534 per %), spell_haste=-14.786 ± 1.089, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.570 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 21.9 spell_power points (4.46 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.18 DPS) [crafted]; Augural Shroud (2620, -0.55 DPS) [world]; Whitemane's Chapeau (7720, -1.37 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.5 spell_power points (2.76 DPS) | yes | Necklace of Calisea (1714, -1.59 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.76 DPS) [world]; Triune Amulet (7722, -2.06 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.6 spell_power points (3.59 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.50 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.1 spell_power points (2.05 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.63 DPS) [crafted]; Blackforge Cape (6424, -0.72 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.9 spell_power points (5.48 DPS) | yes | Dreamweave Vest (10021, -0.32 DPS) [crafted]; Robe of Power (7054, -0.63 DPS) [crafted]; Elemental Raiment (9434, -1.20 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (137.6 DPS) | yes | Windchaser Cuffs (14429, -0.34 DPS) [world_drop]; Condor Bracers (15864, -0.41 DPS) [quest]; Radiant Silver Bracers (4545, -2.22 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (4.33 DPS) | yes | Red Mageweave Gloves (10018, -0.43 DPS) [crafted]; Black Mageweave Gloves (10003, -1.28 DPS) [crafted]; Stormcloth Gloves (10011, -1.52 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.3 spell_power points (3.92 DPS) | yes | Gilded Cord (254037, -0.96 DPS) [crafted]; Star Belt (4329, -1.27 DPS) [crafted]; Crimson Silk Belt (7055, -1.53 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.8 spell_power points (4.85 DPS) | yes | Crimson Silk Pantaloons (7062, -1.06 DPS) [crafted]; Stoneweaver Leggings (9407, -2.09 DPS) [dungeon]; Pristine Leggings (253987, -2.26 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Gilded Slippers (254001, -2.30 DPS) [crafted]; Acidic Walkers (9454, -2.54 DPS) [dungeon]; Spidersilk Boots (4320, -2.80 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.87 DPS) | yes | Reedknot Ring (9622, -1.44 DPS) [quest]; Ring of Power Regulation (273030, -1.65 DPS) [dungeon]; Voodoo Band (1996, -1.70 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.1 spell_power points (2.05 DPS) | yes | Reedknot Ring (9622, -0.63 DPS) [quest]; Ring of Power Regulation (273030, -0.83 DPS) [dungeon]; Voodoo Band (1996, -0.89 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.33 DPS) [dungeon]; Spine of the Devourer (273040, -2.96 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.5 spell_power points (3.15 DPS) | yes | Orb of the Forgotten Seer (7685, -0.30 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.72 DPS) [crafted]; Omega Orb (7749, -1.31 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 198.3 spell_power points (40.38 DPS) | yes | Nether Force Wand (11263, -2.55 DPS) [quest]; Icefury Wand (7514, -2.69 DPS) [quest]; Ragefire Wand (7513, -2.74 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 1531005113100001531-03202300000000000-000000000000000000)

Set DPS (verified): 220.9. Weights run: 1.3s. Verify run: 1.1s. 496 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.060 ± 0.065, crit=0.239 ± 0.007 per rating point (14 rating = 1%, 3.346 per %), hit=0.702 ± 0.090 per rating point (10 rating = 1%, 7.021 per %), spell_haste=-21.271 ± 1.665, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.585 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.2 spell_power points (7.90 DPS) | yes | Dreamweave Circlet (10041, -1.69 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.99 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.59 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.5 spell_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.13 DPS) [world_drop]; Mindburst Medallion (11196, -0.61 DPS) [quest]; Triune Amulet (7722, -1.47 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.1 spell_power points (6.31 DPS) | yes | Red Mageweave Shoulders (10029, -1.80 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.20 DPS) [vendor]; Inquisitor's Shawl (19507, -2.22 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.5 spell_power points (4.23 DPS) | yes | Spritecaster Cape (11623, -0.23 DPS) [dungeon]; Runecloth Cloak (13860, -0.80 DPS) [crafted]; Big Voodoo Cloak (8216, -1.38 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.2 spell_power points (7.90 DPS) | yes | Runecloth Tunic (13857, -2.27 DPS) [crafted]; Robe of the Magi (1716, -2.33 DPS) [world_drop]; Runecloth Robe (13858, -2.88 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.5 spell_power points (2.86 DPS) | yes | Nethergeld Cuffs (254061, -0.02 DPS) [crafted]; Forgotten Wraps (9433, -0.36 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.36 DPS) [quest] |
| hands | First Sergeant's Dreadweave Gloves (220908) | PvP rank 9 · First Sergeant · Horde [vendor] | 22.5 spell_power points (4.43 DPS) | yes | Dreamweave Gloves (10019, -0.06 DPS) [crafted]; Red Mageweave Gloves (10018, -0.18 DPS) [crafted]; Runecloth Gloves (13863, -0.20 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.1 spell_power points (5.14 DPS) | yes | Satyrmane Sash (17755, -0.30 DPS) [dungeon]; Deathmage Sash (10771, -0.64 DPS) [dungeon]; Runecloth Belt (13856, -1.07 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 28.1 spell_power points (5.52 DPS) | yes | Red Mageweave Pants (10009, -0.26 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.18 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.24 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 24.6 spell_power points (4.83 DPS) | yes | Earthen Silk Slippers (254013, -0.11 DPS) [crafted]; Gilded Sandals (254107, -0.79 DPS) [crafted]; Southsea Mojo Boots (20641, -0.96 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.3 spell_power points (3.01 DPS) | yes | Snake Hoop (6750, -0.79 DPS) [quest]; Ogremind Ring (1993, -1.55 DPS) [world_drop]; Voodoo Band (1996, -1.55 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.56 DPS) | yes | Snake Hoop (6750, -0.33 DPS) [quest]; Ogremind Ring (1993, -1.10 DPS) [world_drop]; Voodoo Band (1996, -1.10 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (220.9 DPS) | yes | Uther's Strength (11302, -0.06 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (220.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (220.9 DPS) | yes | Spellshifter Rod (9527, -0.46 DPS) [quest]; Kindling Stave (11750, -1.58 DPS) [dungeon]; Hanzo Sword (8190, -38.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 267.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.52 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.18 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: First Sergeant's Dreadweave Gloves; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 496, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 1531005113100001531-03202300000000000-053200000000000000)

Set DPS (verified): 355.3. Weights run: 1.4s. Verify run: 3.0s. 1117 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.093 ± 0.070, crit=0.399 ± 0.011 per rating point (14 rating = 1%, 5.591 per %), hit=0.932 ± 0.107 per rating point (10 rating = 1%, 9.322 per %), spell_haste=-30.393 ± 1.876, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.527 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 43.8 spell_power points (10.36 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -11.74 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 22.9 spell_power points (5.43 DPS) | yes | Glowing Eye of Mordresh (10769, -1.70 DPS) [dungeon]; Horizon Choker (13085, -1.81 DPS) [world_drop]; Mindburst Medallion (11196, -2.45 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (355.3 DPS) | yes | Elder Wizard's Mantle (13013, +0.00 DPS) [world_drop]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -7.86 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.8 spell_power points (5.17 DPS) | yes | Spritecaster Cape (11623, -0.30 DPS) [dungeon]; Runecloth Cloak (13860, -0.97 DPS) [crafted]; Big Voodoo Cloak (8216, -1.66 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.7 spell_power points (13.89 DPS) | yes | Warlord's Silk Raiment (231596, -0.36 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.20 DPS) [pvp]; Magister's Robes (16688, -5.80 DPS, sim-verified) [dungeon] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (355.3 DPS) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; General's Silk Cuffs (16538, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, -8.89 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.1 spell_power points (7.36 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 30.0 spell_power points (7.09 DPS) | yes | Wisdom of the Timbermaw (19047, -0.86 DPS) [crafted]; Devout Belt (16696, -0.97 DPS) [dungeon]; Dawnspire Cord (12466, -10.83 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 31.8 spell_power points (7.52 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Stone Guard's Dreadweave Leggings (220906, -15.91 DPS, sim-verified) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 33.1 spell_power points (7.84 DPS) | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.28 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.7 spell_power points (6.07 DPS) | yes | Glowing Crystal Ring (18402, -2.97 DPS) [quest]; Band of the Unicorn (7553, -2.99 DPS) [world_drop]; Snake Hoop (6750, -3.36 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (3.66 DPS) | yes | Glowing Crystal Ring (18402, -0.56 DPS) [quest]; Band of the Unicorn (7553, -0.58 DPS) [world_drop]; Snake Hoop (6750, -0.95 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (355.3 DPS) | yes | Uther's Strength (11302, -0.57 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (355.3 DPS) | yes | Uther's Strength (11302, -0.12 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (355.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.61 DPS) [quest]; Teebu's Blazing Longsword (1728, -44.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 330.3 spell_power points (78.18 DPS) | yes | Stormrager (16997, -14.20 DPS) [quest]; Brilliant Wand (249385, -14.66 DPS) [crafted]; Torch of Austen (13004, -21.39 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Magister's Mantle; chest: Robe of the Archmage; wrist: Magister's Bindings; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; ranged: Torch of Light

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 1531005113100101531-03000000000000000-053500030000000000)

Set DPS (verified): 565.9. Weights run: 1.5s. Verify run: 2.2s. 1117 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.453 ± 0.013, crit=0.465 ± 0.012 per rating point (14 rating = 1%, 6.512 per %), hit=0.996 ± 0.038 per rating point (10 rating = 1%, 9.959 per %), spell_haste=-13.693 ± 0.827, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.475 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 28.1 spell_power points (12.61 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -0.48 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.5 spell_power points (7.43 DPS) | yes | Glowing Eye of Mordresh (10769, -2.65 DPS) [dungeon]; Mindburst Medallion (11196, -3.51 DPS) [quest]; Triune Amulet (7722, -3.83 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.4 spell_power points (13.21 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.31 DPS) [pvp]; Argent Shoulders (19059, -1.98 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.7 spell_power points (7.51 DPS) | yes | Deep Woodlands Cloak (19121, -0.29 DPS) [quest]; Runecloth Cloak (13860, -1.84 DPS) [crafted]; Big Voodoo Cloak (8216, -3.43 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 51.9 spell_power points (23.35 DPS) | yes | Warlord's Silk Raiment (231596, -2.13 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -7.52 DPS) [pvp]; Acumen Robes (17775, -12.53 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 15.5 spell_power points (6.98 DPS) | yes | Nethergeld Cuffs (254061, -2.41 DPS) [crafted]; Dreadmist Bracers (16703, -2.49 DPS) [dungeon]; Bloodband Bracers (11469, -2.90 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 26.6 spell_power points (11.93 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -4.76 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 18.5 spell_power points (8.33 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Magister's Belt (16685, -0.91 DPS) [dungeon]; Dreadmist Belt (16702, -1.43 DPS) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.68 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS) [vendor]; Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.4 spell_power points (11.43 DPS) | yes | Blood Guard's Silk Walkers (227110, +0.00 DPS) [pvp]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.65 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.2 spell_power points (9.51 DPS) | yes | Philanthropist's Ring (281635, -4.00 DPS) [quest]; Snake Hoop (6750, -5.80 DPS) [quest]; Reedknot Ring (9622, -6.37 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.84 DPS) | yes | Philanthropist's Ring (281635, -0.33 DPS) [quest]; Snake Hoop (6750, -2.13 DPS) [quest]; Reedknot Ring (9622, -2.70 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (565.9 DPS) | yes | Rune of the Guard Captain (19120, -0.90 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (565.9 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Mark of the Chosen (17774, -5.73 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (565.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.46 DPS) [quest]; Teebu's Blazing Longsword (1728, -40.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 176.8 spell_power points (79.46 DPS) | yes | Brilliant Wand (249385, -14.41 DPS) [crafted]; Stormrager (16997, -15.75 DPS) [quest]; Greater Eternal Wand (249237, -22.08 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Spritecaster Cape; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Omnicast Boots; finger1: Maiden's Circle; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

