# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 1531001000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 34.3. Weights run: 1.1s. Verify run: 0.7s. 154 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.235 ± 0.007, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.854 per %), hit=0.215 ± 0.007 per rating point (10 rating = 1%, 2.150 per %), spell_haste=0.593 ± 0.102, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.135 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.72 DPS) | yes | Shadow Goggles (4373, -1.68 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.1 spell_power points (0.86 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.37 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.48 DPS) | yes | Pearl-clasped Cloak (5542, -0.16 DPS) [crafted]; Reinforced Linen Cape (2580, -0.36 DPS) [crafted]; Black Whelp Cloak (7283, -0.38 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.2 spell_power points (0.74 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -1.08 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.2 spell_power points (0.14 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Repurposed Hair Band (281256, -0.08 DPS) [quest]; Windsong Bangles (263336, -0.97 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.84 DPS) | yes | Gnoll Casting Gloves (892, -0.12 DPS) [world]; Pristine Gloves (253913, -0.28 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.54 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.59 DPS) | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Keller's Girdle (2911, -0.37 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 7.4 spell_power points (0.89 DPS) | yes | Colorful Kilt (10048, -0.29 DPS) [crafted]; Heavy Woolen Pants (4316, -0.53 DPS) [crafted]; Silk-threaded Trousers (1929, -1.62 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.96 DPS) | yes | Pristine Boots (253889, -0.51 DPS) [crafted]; Woolen Boots (2583, -0.66 DPS) [crafted]; Red Woolen Boots (4313, -1.10 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (0.66 DPS) | yes | Volcanic Rock Ring (12053, -0.57 DPS) [world_drop]; The 1 Ring (8350, -0.63 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.4 spell_power points (0.17 DPS) | yes | The 1 Ring (8350, -0.14 DPS) [world]; Volcanic Rock Ring (12053, -0.36 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.3 spell_power points (0.28 DPS) | yes | Witching Stave (1484, -0.14 DPS) [dungeon]; Staff of Westfall (2042, -0.14 DPS) [quest]; Lesser Staff of the Spire (1300, -0.62 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 188.4 spell_power points (22.67 DPS) | yes | Skycaller (12984, -0.90 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.14 DPS) [quest]; Sizzle Stick (8071, -4.42 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 154, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 1531005113100000000-00000000000000000-000000000000000000)

Set DPS (verified): 74.6. Weights run: 1.2s. Verify run: 0.8s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.598 ± 0.022, crit=0.087 ± 0.003 per rating point (14 rating = 1%, 1.212 per %), hit=0.343 ± 0.026 per rating point (10 rating = 1%, 3.432 per %), spell_haste=1.512 ± 0.375, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.517 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.27 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.40 DPS) [crafted]; Enchanter's Cowl (4322, -2.04 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.6 spell_power points (0.48 DPS) | yes | Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (1.94 DPS) | yes | Death Speaker Mantle (6685, -0.24 DPS) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Magician's Mantle (12998, -0.54 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.67 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Prelacy Cape (7004, -0.13 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.8 spell_power points (2.26 DPS) | yes | Beguiler Robes (7728, -0.08 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.22 DPS) [dungeon]; Death Speaker Robes (6682, -0.43 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.21 DPS) | yes | Nightsky Wristbands (6407, -0.73 DPS) [world_drop]; Stonecloth Bindings (14416, -0.81 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.62 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.2 spell_power points (0.97 DPS) | yes | Serpent Gloves (5970, -0.02 DPS) [dungeon]; Shilly Mitts (9609, -0.02 DPS) [quest]; Truefaith Gloves (7049, -0.05 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 10.8 spell_power points (1.46 DPS) | yes | Crimson Silk Belt (7055, -0.08 DPS) [crafted]; Invoker's Cord (215366, -0.11 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.38 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.22 DPS) [crafted]; Silk-threaded Trousers (1929, -0.56 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.28 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.2 spell_power points (1.51 DPS) | yes | Spidersilk Boots (4320, -0.24 DPS) [crafted]; Glinteye Slippers (273024, -0.56 DPS) [dungeon]; Acidic Walkers (9454, -1.75 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.21 DPS) | yes | Ring of Power Regulation (273030, -0.40 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.73 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.81 DPS) [dungeon] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.84 DPS) | yes | Ring of Power Regulation (273030, -0.03 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.43 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.85 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -2.16 DPS) [quest]; Glimmering Staff (249392, -3.97 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 246.1 spell_power points (33.18 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.51 DPS) [crafted]; Gravestone Scepter (7001, -4.18 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Minor Channeling Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 1531005113100001531-00000000000000000-000000000000000000)

Set DPS (verified): 111.0. Weights run: 1.3s. Verify run: 1.0s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.705 ± 0.044, crit=0.161 ± 0.005 per rating point (14 rating = 1%, 2.253 per %), hit=0.462 ± 0.044 per rating point (10 rating = 1%, 4.618 per %), spell_haste=not significant (2.785 ± 0.713), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.561 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (111.0 DPS) | yes | Augural Shroud (2620, -0.29 DPS) [world]; Whitemane's Chapeau (7720, -0.89 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -3.04 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.6 spell_power points (1.89 DPS) | yes | Necklace of Calisea (1714, -1.15 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.26 DPS) [world]; Triune Amulet (7722, -1.84 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (2.42 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (1.42 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.38 DPS) [crafted]; Blackforge Cape (6424, -0.58 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (3.92 DPS) | yes | Dreamweave Vest (10021, -0.28 DPS) [crafted]; Robe of Power (7054, -0.56 DPS) [crafted]; Elemental Raiment (9434, -0.78 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.35 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.30 DPS) [quest]; Windchaser Cuffs (14429, -0.40 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (3.11 DPS) | yes | Red Mageweave Gloves (10018, -0.41 DPS) [crafted]; Black Mageweave Gloves (10003, -0.87 DPS) [crafted]; Gilded Handwraps (254021, -1.18 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.6 spell_power points (2.63 DPS) | yes | Gilded Cord (254037, -0.59 DPS) [crafted]; Star Belt (4329, -0.68 DPS) [crafted]; Belt of Arugal (6392, -0.97 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (3.36 DPS) | yes | Stoneweaver Leggings (9407, -1.47 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.56 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.67 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Gilded Slippers (254001, -1.80 DPS) [crafted]; Acidic Walkers (9454, -2.00 DPS) [dungeon]; Spidersilk Boots (4320, -2.12 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (2.02 DPS) | yes | Ring of Forlorn Spirits (2043, -0.83 DPS) [quest]; Reedknot Ring (9622, -0.98 DPS) [quest]; Minor Channeling Ring (1449, -1.06 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (1.42 DPS) | yes | Ring of Forlorn Spirits (2043, -0.23 DPS) [quest]; Reedknot Ring (9622, -0.38 DPS) [quest]; Minor Channeling Ring (1449, -0.47 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.21 DPS) [dungeon]; Spine of the Devourer (273040, -2.04 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.1 spell_power points (2.26 DPS) | yes | Orb of the Forgotten Seer (7685, -0.17 DPS) [dungeon]; Orb of Lorica (11262, -0.31 DPS) [quest]; Orb of Mystic Insight (249394, -0.58 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 267.7 spell_power points (40.01 DPS) | yes | Nether Force Wand (11263, -2.37 DPS) [quest]; Icefury Wand (7514, -2.51 DPS) [quest]; Ragefire Wand (7513, -2.56 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 1531005113100001531-03202300000000000-000000000000000000)

Set DPS (verified): 185.5. Weights run: 1.2s. Verify run: 1.2s. 506 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.827 ± 0.071, crit=0.263 ± 0.009 per rating point (14 rating = 1%, 3.680 per %), hit=0.726 ± 0.078 per rating point (10 rating = 1%, 7.263 per %), spell_haste=5.491 ± 1.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.584 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.5 spell_power points (5.16 DPS) | yes | Dreamweave Circlet (10041, -0.91 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.15 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.24 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 spell_power points (1.98 DPS) | yes | Horizon Choker (13085, -0.30 DPS) [world_drop]; Mindburst Medallion (11196, -0.39 DPS) [quest]; Triune Amulet (7722, -0.82 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.9 spell_power points (4.05 DPS) | yes | Red Mageweave Shoulders (10029, -1.23 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.27 DPS) [vendor]; Inquisitor's Shawl (19507, -1.47 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.0 spell_power points (2.75 DPS) | yes | Big Voodoo Cloak (8216, -0.95 DPS) [crafted]; Long Silken Cloak (4326, -1.28 DPS) [crafted]; Runecloth Cloak (13860, -3.03 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.5 spell_power points (5.16 DPS) | yes | Robe of the Magi (1716, -1.25 DPS) [world_drop]; Runecloth Tunic (13857, -1.37 DPS) [crafted]; Dreamweave Vest (10021, -1.47 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.8 spell_power points (1.86 DPS) | yes | Bloodband Bracers (11469, -0.05 DPS) [quest]; Forgotten Wraps (9433, -0.42 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.42 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (3.09 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -0.13 DPS) [vendor]; Runecloth Gloves (13863, -0.27 DPS) [crafted]; Red Mageweave Gloves (10018, -0.30 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 22.3 spell_power points (3.23 DPS) | yes | Dawnspire Cord (12466, -0.08 DPS) [dungeon]; Deathmage Sash (10771, -0.42 DPS) [dungeon]; Runecloth Belt (13856, -0.63 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 25.6 spell_power points (3.72 DPS) | yes | Red Mageweave Pants (10009, -0.24 DPS) [crafted]; Wizardweave Leggings (14132, -0.96 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.00 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.49 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.19 DPS) [vendor]; Gilded Sandals (254107, -0.81 DPS) [crafted]; Southsea Mojo Boots (20641, -1.00 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.05 DPS) | yes | Snake Hoop (6750, -0.58 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.89 DPS) [quest]; Reedknot Ring (9622, -1.04 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.89 DPS) | yes | Snake Hoop (6750, -0.42 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.73 DPS) [quest]; Reedknot Ring (9622, -0.87 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (185.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (185.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (185.5 DPS) | yes | Spellshifter Rod (9527, -0.14 DPS) [quest]; Kindling Stave (11750, -0.59 DPS) [dungeon]; Hanzo Sword (8190, -32.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 361.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.59 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Knight's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 506, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 1531005113100001531-03202300000000000-053200000000000000)

Set DPS (verified): 305.7. Weights run: 1.3s. Verify run: 3.0s. 1125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.007 ± 0.083, crit=0.430 ± 0.013 per rating point (14 rating = 1%, 6.027 per %), hit=1.172 ± 0.106 per rating point (10 rating = 1%, 11.718 per %), spell_haste=not significant (1.799 ± 1.258), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.529 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 41.2 spell_power points (7.09 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -21.38 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 22.1 spell_power points (3.80 DPS) | yes | Glowing Eye of Mordresh (10769, -1.21 DPS) [dungeon]; Horizon Choker (13085, -1.37 DPS) [world_drop]; Opaline Medallion (16623, -1.55 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (305.7 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -16.88 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.0 spell_power points (3.45 DPS) | yes | Runecloth Cloak (13860, -0.51 DPS) [crafted]; Big Voodoo Cloak (8216, -1.03 DPS) [crafted]; Imperial Red Cloak (8248, -1.54 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.1 spell_power points (10.00 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.34 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -2.40 DPS) [pvp]; Magister's Robes (16688, -3.08 DPS) [dungeon] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (305.7 DPS) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; Marshal's Silk Bracers (16438, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, -16.24 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.1 spell_power points (5.36 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -5.57 DPS, sim-verified) [vendor] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 28.1 spell_power points (4.84 DPS) | yes | Wisdom of the Timbermaw (19047, -0.52 DPS) [crafted]; Devout Belt (16696, -0.69 DPS) [dungeon]; Dawnspire Cord (12466, -19.12 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 32.6 spell_power points (5.62 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -23.24 DPS, sim-verified) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 32.1 spell_power points (5.52 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.17 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.57 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.0 spell_power points (4.31 DPS) | yes | Band of the Unicorn (7553, -2.07 DPS) [world_drop]; Glowing Crystal Ring (18402, -2.23 DPS) [quest]; Snake Hoop (6750, -2.41 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.0 spell_power points (2.59 DPS) | yes | Band of the Unicorn (7553, -0.35 DPS) [world_drop]; Glowing Crystal Ring (18402, -0.51 DPS) [quest]; Snake Hoop (6750, -0.69 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (305.7 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (305.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (305.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.35 DPS) [quest]; Teebu's Blazing Longsword (1728, -43.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 452.0 spell_power points (77.79 DPS) | yes | Stormrager (16997, -5.30 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.45 DPS) [crafted]; Torch of Austen (13004, -21.00 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Magister's Mantle; chest: Robe of the Archmage; wrist: Magister's Bindings; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; ranged: Torch of Light

No-known-source sample (15 of 1125, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 0532005113100101531-03000000000000000-053500030000000000)

Set DPS (verified): 623.7. Weights run: 1.6s. Verify run: 2.1s. 1125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.906 ± 0.048, crit=0.540 ± 0.012 per rating point (14 rating = 1%, 7.558 per %), hit=1.023 ± 0.096 per rating point (10 rating = 1%, 10.234 per %), spell_haste=4.214 ± 0.745, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.599 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 38.2 spell_power points (16.08 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -0.45 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 21.1 spell_power points (8.87 DPS) | yes | Glowing Eye of Mordresh (10769, -2.87 DPS) [dungeon]; Horizon Choker (13085, -3.53 DPS) [world_drop]; Opaline Medallion (16623, -3.91 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.3 spell_power points (15.31 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.61 DPS) [pvp]; Rotgrip Mantle (17732, -5.37 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.4 spell_power points (8.19 DPS) | yes | Runecloth Cloak (13860, -1.34 DPS) [crafted]; Big Voodoo Cloak (8216, -2.65 DPS) [crafted]; Long Silken Cloak (4326, -3.75 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.4 spell_power points (24.61 DPS) | yes | Field Marshal's Silk Vestments (231603, -1.04 DPS) [pvp]; Magister's Robes (16688, -5.63 DPS, sim-verified) [dungeon]; Knight-Captain's Silk Tunic (227108, -6.10 DPS) [pvp] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 20.1 spell_power points (8.45 DPS) | yes | Marshal's Silk Bracers (16438, -1.20 DPS) [pvp]; Dreadmist Bracers (16703, -2.15 DPS) [dungeon]; Magiskull Cuffs (13107, -3.60 DPS, sim-verified) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 33.2 spell_power points (13.98 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -5.81 DPS, sim-verified) [quest] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 26.0 spell_power points (10.96 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.18 DPS) [dungeon]; Satyrmane Sash (17755, -1.25 DPS) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (623.7 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -9.92 DPS, sim-verified) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 30.9 spell_power points (13.01 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.34 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -1.89 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.3 spell_power points (10.25 DPS) | yes | Band of the Unicorn (7553, -4.78 DPS) [world_drop]; Glowing Crystal Ring (18402, -5.67 DPS) [quest]; Snake Hoop (6750, -5.82 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (6.12 DPS) | yes | Band of the Unicorn (7553, -0.64 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.54 DPS) [quest]; Snake Hoop (6750, -1.68 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.35 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.61 DPS) [quest]; Teebu's Blazing Longsword (1728, -58.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 188.2 spell_power points (79.29 DPS) | yes | Brilliant Wand (249385, -13.24 DPS) [crafted]; Stormrager (16997, -14.69 DPS) [quest]; Greater Eternal Wand (249237, -22.17 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Mark of the Chosen; ranged: Torch of Light

No-known-source sample (15 of 1125, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 1531001000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 31.5. Weights run: 1.1s. Verify run: 0.7s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.235 ± 0.007, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.854 per %), hit=0.215 ± 0.007 per rating point (10 rating = 1%, 2.150 per %), spell_haste=0.593 ± 0.102, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.135 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.72 DPS) | yes | Shadow Goggles (4373, -1.38 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.1 spell_power points (0.86 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.37 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.43 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.48 DPS) | yes | Black Whelp Cloak (7283, -0.12 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.16 DPS) [crafted]; Mystic Shawl (3449, -0.24 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.2 spell_power points (0.74 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -0.89 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.36 DPS) | yes | Owlbeard Bracers (16981, -0.18 DPS) [quest]; Mindthrust Bracers (1974, -0.22 DPS) [dungeon]; Featherbead Bracers (15452, -0.22 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.84 DPS) | yes | Gnoll Casting Gloves (892, -0.12 DPS) [world]; Pristine Gloves (253913, -0.28 DPS) [crafted]; Apothecary Gloves (10919, -0.36 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.59 DPS) | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Keller's Girdle (2911, -0.37 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.53 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 7.4 spell_power points (0.89 DPS) | yes | Colorful Kilt (10048, -0.29 DPS) [crafted]; Embroidered Leggings (270011, -0.41 DPS) [quest]; Silk-threaded Trousers (1929, -1.39 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.96 DPS) | yes | Red Woolen Boots (4313, -0.50 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.51 DPS) [crafted]; Woolen Boots (2583, -0.66 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.4 spell_power points (0.17 DPS) | yes | Volcanic Rock Ring (12053, -0.08 DPS) [world_drop]; The 1 Ring (8350, -0.14 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 1.2 spell_power points (0.14 DPS) | yes | Volcanic Rock Ring (12053, -0.06 DPS) [world_drop]; The 1 Ring (8350, -0.11 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 2.3 spell_power points (0.28 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.11 DPS) [world]; Clear Crystal Rod (16894, -0.14 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 188.4 spell_power points (22.67 DPS) | yes | Skycaller (12984, -0.35 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.42 DPS) [quest]; Greater Magic Wand (11288, -4.93 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 1531005113100000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.0. Weights run: 1.2s. Verify run: 0.8s. 282 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.598 ± 0.022, crit=0.087 ± 0.003 per rating point (14 rating = 1%, 1.212 per %), hit=0.343 ± 0.026 per rating point (10 rating = 1%, 3.432 per %), spell_haste=1.512 ± 0.375, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.517 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.27 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.40 DPS) [crafted]; Enchanter's Cowl (4322, -2.49 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.6 spell_power points (0.48 DPS) | yes | Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.4 spell_power points (1.94 DPS) | yes | Death Speaker Mantle (6685, -0.24 DPS) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Magician's Mantle (12998, -0.54 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.16 DPS) [crafted]; Hillman's Cloak (3719, -1.29 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.8 spell_power points (2.26 DPS) | yes | Beguiler Robes (7728, -0.08 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.22 DPS) [dungeon]; Death Speaker Robes (6682, -0.43 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.21 DPS) | yes | Glowing Magical Bracelets (13106, -0.57 DPS) [world_drop]; Nightsky Wristbands (6407, -0.73 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.81 DPS) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.0 spell_power points (1.21 DPS) | yes | Serpent Gloves (5970, -0.27 DPS) [dungeon]; Truefaith Gloves (7049, -0.30 DPS) [crafted]; Gnoll Casting Gloves (892, -0.40 DPS) [world] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Crimson Silk Belt (7055, -0.08 DPS) [crafted]; Invoker's Cord (215366, -0.11 DPS) [crafted]; Warsong Sash (16975, -1.82 DPS, sim-verified) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.22 DPS) [crafted]; Silk-threaded Trousers (1929, -0.56 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.11 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.2 spell_power points (1.51 DPS) | yes | Acidic Walkers (9454, -0.19 DPS) [dungeon]; Spidersilk Boots (4320, -0.24 DPS) [crafted]; Glinteye Slippers (273024, -0.56 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.0 spell_power points (1.21 DPS) | yes | Solid Gold Band (279845, -0.67 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.73 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.81 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.81 DPS) | yes | Solid Gold Band (279845, -0.27 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.40 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.85 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -3.97 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -4.05 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 246.1 spell_power points (33.18 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Unstable Power Core (279847, -3.47 DPS) [quest]; Greater Mystic Wand (217287, -3.51 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 282, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (orc, 1531005113100001531-00000000000000000-000000000000000000)

Set DPS (verified): 105.5. Weights run: 1.3s. Verify run: 1.0s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.705 ± 0.044, crit=0.161 ± 0.005 per rating point (14 rating = 1%, 2.253 per %), hit=0.462 ± 0.044 per rating point (10 rating = 1%, 4.618 per %), spell_haste=not significant (2.785 ± 0.713), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.561 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (105.5 DPS) | yes | Augural Shroud (2620, -0.29 DPS) [world]; Whitemane's Chapeau (7720, -0.89 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.77 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.6 spell_power points (1.89 DPS) | yes | Necklace of Calisea (1714, -1.15 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.26 DPS) [world]; Triune Amulet (7722, -1.97 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (2.42 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.32 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (1.42 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.38 DPS) [crafted]; Blackforge Cape (6424, -0.58 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (3.92 DPS) | yes | Dreamweave Vest (10021, -0.28 DPS) [crafted]; Robe of Power (7054, -0.56 DPS) [crafted]; Elemental Raiment (9434, -0.78 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.6 spell_power points (1.44 DPS) | yes | Spidertank Oilrag (9448, -0.10 DPS) [dungeon]; Condor Bracers (15864, -0.39 DPS) [quest]; Windchaser Cuffs (14429, -0.49 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (3.11 DPS) | yes | Black Mageweave Gloves (10003, -0.87 DPS) [crafted]; Gilded Handwraps (254021, -1.18 DPS) [crafted]; Red Mageweave Gloves (10018, -1.68 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.6 spell_power points (2.63 DPS) | yes | Gilded Cord (254037, -0.59 DPS) [crafted]; Star Belt (4329, -0.68 DPS) [crafted]; Belt of Arugal (6392, -0.97 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (3.36 DPS) | yes | Stoneweaver Leggings (9407, -1.47 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.56 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.06 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Gilded Slippers (254001, -1.80 DPS) [crafted]; Acidic Walkers (9454, -2.00 DPS) [dungeon]; Spidersilk Boots (4320, -2.12 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (2.02 DPS) | yes | Reedknot Ring (9622, -0.98 DPS) [quest]; Ring of Power Regulation (273030, -1.13 DPS) [dungeon]; Voodoo Band (1996, -1.28 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (1.42 DPS) | yes | Reedknot Ring (9622, -0.38 DPS) [quest]; Ring of Power Regulation (273030, -0.53 DPS) [dungeon]; Voodoo Band (1996, -0.69 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.21 DPS) [dungeon]; Spine of the Devourer (273040, -2.04 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.1 spell_power points (2.26 DPS) | yes | Orb of the Forgotten Seer (7685, -0.17 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.58 DPS) [crafted]; Omega Orb (7749, -0.91 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 267.7 spell_power points (40.01 DPS) | yes | Nether Force Wand (11263, -2.37 DPS) [quest]; Icefury Wand (7514, -2.51 DPS) [quest]; Ragefire Wand (7513, -2.56 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 1531005113100001531-03202300000000000-000000000000000000)

Set DPS (verified): 179.3. Weights run: 1.2s. Verify run: 1.2s. 496 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=0.827 ± 0.071, crit=0.263 ± 0.009 per rating point (14 rating = 1%, 3.680 per %), hit=0.726 ± 0.078 per rating point (10 rating = 1%, 7.263 per %), spell_haste=5.491 ± 1.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.584 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 35.5 spell_power points (5.16 DPS) | yes | Blood Guard's Dreadweave Hat (220907, -1.15 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.24 DPS) [crafted]; Dreamweave Circlet (10041, -3.00 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 spell_power points (1.98 DPS) | yes | Horizon Choker (13085, -0.30 DPS) [world_drop]; Mindburst Medallion (11196, -0.39 DPS) [quest]; Triune Amulet (7722, -0.82 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 27.9 spell_power points (4.05 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.27 DPS) [vendor]; Inquisitor's Shawl (19507, -1.47 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.48 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.4 spell_power points (2.82 DPS) | yes | Spritecaster Cape (11623, -0.07 DPS) [dungeon]; Runecloth Cloak (13860, -0.56 DPS) [crafted]; Big Voodoo Cloak (8216, -1.02 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 35.5 spell_power points (5.16 DPS) | yes | Robe of the Magi (1716, -1.25 DPS) [world_drop]; Runecloth Tunic (13857, -1.37 DPS) [crafted]; Dreamweave Vest (10021, -1.47 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 12.8 spell_power points (1.86 DPS) | yes | Bloodband Bracers (11469, -0.05 DPS) [quest]; Radiant Silver Bracers (4545, -0.32 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.42 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (3.09 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -0.13 DPS) [vendor]; Runecloth Gloves (13863, -0.27 DPS) [crafted]; Red Mageweave Gloves (10018, -0.30 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (179.3 DPS) | yes | Deathmage Sash (10771, -0.34 DPS) [dungeon]; Runecloth Belt (13856, -0.55 DPS) [crafted]; Satyrmane Sash (17755, -3.12 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 25.6 spell_power points (3.72 DPS) | yes | Red Mageweave Pants (10009, -0.24 DPS) [crafted]; Wizardweave Leggings (14132, -0.96 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.00 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.49 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.19 DPS) [vendor]; Gilded Sandals (254107, -0.81 DPS) [crafted]; Southsea Mojo Boots (20641, -1.00 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.05 DPS) | yes | Snake Hoop (6750, -0.58 DPS) [quest]; Reedknot Ring (9622, -1.04 DPS) [quest]; Ring of Power Regulation (273030, -1.18 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.89 DPS) | yes | Snake Hoop (6750, -0.42 DPS) [quest]; Reedknot Ring (9622, -0.87 DPS) [quest]; Ring of Power Regulation (273030, -1.02 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.08 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.14 DPS) [quest]; Kindling Stave (11750, -0.59 DPS) [dungeon]; Hanzo Sword (8190, -30.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 361.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.59 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 496, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 1531005113100001531-03202300000000000-053200000000000000)

Set DPS (verified): 292.7. Weights run: 1.3s. Verify run: 3.1s. 1117 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.006, intellect=1.007 ± 0.083, crit=0.430 ± 0.013 per rating point (14 rating = 1%, 6.027 per %), hit=1.172 ± 0.106 per rating point (10 rating = 1%, 11.718 per %), spell_haste=not significant (1.799 ± 1.258), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.529 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (292.7 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Magister's Crown (16686, -14.41 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 22.1 spell_power points (3.80 DPS) | yes | Glowing Eye of Mordresh (10769, -1.21 DPS) [dungeon]; Horizon Choker (13085, -1.37 DPS) [world_drop]; Mindburst Medallion (11196, -1.73 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (292.7 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Magister's Mantle (16689, -14.84 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.1 spell_power points (3.62 DPS) | yes | Spritecaster Cape (11623, -0.18 DPS) [dungeon]; Runecloth Cloak (13860, -0.69 DPS) [crafted]; Big Voodoo Cloak (8216, -1.20 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.1 spell_power points (10.00 DPS) | yes | Warlord's Silk Raiment (231596, -0.34 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.40 DPS) [pvp]; Magister's Robes (16688, -3.08 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (292.7 DPS) | yes | General's Silk Cuffs (16538, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Magister's Bindings (16683, -13.80 DPS, sim-verified) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.1 spell_power points (5.36 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (292.7 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Magister's Belt (16685, -15.11 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (292.7 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -14.35 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 32.1 spell_power points (5.52 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.17 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -4.74 DPS, sim-verified) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.0 spell_power points (4.31 DPS) | yes | Band of the Unicorn (7553, -2.07 DPS) [world_drop]; Glowing Crystal Ring (18402, -2.23 DPS) [quest]; Snake Hoop (6750, -2.41 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.0 spell_power points (2.59 DPS) | yes | Band of the Unicorn (7553, -0.35 DPS) [world_drop]; Glowing Crystal Ring (18402, -0.51 DPS) [quest]; Snake Hoop (6750, -0.69 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (292.7 DPS) | yes | Uther's Strength (11302, -0.38 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (292.7 DPS) | yes | Uther's Strength (11302, -0.78 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (292.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.35 DPS) [quest]; Teebu's Blazing Longsword (1728, -39.99 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 452.0 spell_power points (77.79 DPS) | yes | Stormrager (16997, -5.28 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.45 DPS) [crafted]; Torch of Austen (13004, -21.00 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 0532005113100101531-03000000000000000-053500030000000000)

Set DPS (verified): 605.1. Weights run: 1.6s. Verify run: 2.4s. 1117 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.906 ± 0.048, crit=0.540 ± 0.012 per rating point (14 rating = 1%, 7.558 per %), hit=1.023 ± 0.096 per rating point (10 rating = 1%, 10.234 per %), spell_haste=4.214 ± 0.745, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.599 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 38.2 spell_power points (16.08 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -0.45 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 21.1 spell_power points (8.87 DPS) | yes | Glowing Eye of Mordresh (10769, -2.87 DPS) [dungeon]; Horizon Choker (13085, -3.53 DPS) [world_drop]; Mindburst Medallion (11196, -4.05 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.3 spell_power points (15.31 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.61 DPS) [pvp]; Rotgrip Mantle (17732, -4.92 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.2 spell_power points (8.49 DPS) | yes | Spritecaster Cape (11623, -0.30 DPS) [dungeon]; Runecloth Cloak (13860, -1.65 DPS) [crafted]; Big Voodoo Cloak (8216, -2.95 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 58.4 spell_power points (24.61 DPS) | yes | Warlord's Silk Raiment (231596, -1.04 DPS) [pvp]; Magister's Robes (16688, -4.83 DPS, sim-verified) [dungeon]; Legionnaire's Silk Tunic (227106, -6.10 DPS) [pvp] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 20.1 spell_power points (8.45 DPS) | yes | General's Silk Cuffs (16538, -1.20 DPS) [pvp]; Dreadmist Bracers (16703, -2.15 DPS) [dungeon]; Magiskull Cuffs (13107, -3.42 DPS, sim-verified) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 33.2 spell_power points (13.98 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -5.72 DPS, sim-verified) [quest] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 26.0 spell_power points (10.96 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.18 DPS) [dungeon]; Satyrmane Sash (17755, -1.25 DPS) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 30.4 spell_power points (12.82 DPS) | yes | Magister's Leggings (16687, +0.00 DPS, sim-verified) [dungeon]; Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 30.9 spell_power points (13.01 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.34 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -4.50 DPS, sim-verified) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.3 spell_power points (10.25 DPS) | yes | Band of the Unicorn (7553, -4.78 DPS) [world_drop]; Glowing Crystal Ring (18402, -5.67 DPS) [quest]; Snake Hoop (6750, -5.82 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (6.12 DPS) | yes | Band of the Unicorn (7553, -0.64 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.54 DPS) [quest]; Snake Hoop (6750, -1.68 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (605.1 DPS) | yes | Uther's Strength (11302, -1.35 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (605.1 DPS) | yes | Uther's Strength (11302, -0.49 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (605.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.61 DPS) [quest]; Teebu's Blazing Longsword (1728, -58.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 188.2 spell_power points (79.29 DPS) | yes | Brilliant Wand (249385, -13.24 DPS) [crafted]; Stormrager (16997, -14.69 DPS) [quest]; Greater Eternal Wand (249237, -22.17 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; ranged: Torch of Light

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

