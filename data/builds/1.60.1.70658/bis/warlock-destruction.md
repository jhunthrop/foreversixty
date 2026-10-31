# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 39.2. Weights run: 2.4s. Verify run: 1.2s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.114 ± 0.003, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.389 per %), hit=0.116 ± 0.004 per rating point (10 rating = 1%, 1.156 per %), spell_haste=0.533 ± 0.050, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.002, fire_power=0.235 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.20 DPS) | yes | Shadow Goggles (4373, -1.08 DPS) [crafted]; Flame Circlet (253953, -2.39 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.0 spell_power points (1.20 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.41 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Pearl-clasped Cloak (5542, -0.33 DPS) [crafted]; Black Whelp Cloak (7283, -0.41 DPS, sim-verified) [crafted]; Reinforced Linen Cape (2580, -0.60 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 spell_power points (1.11 DPS) | yes | Green Woolen Vest (2582, -0.31 DPS) [crafted]; Bloody Apron (6226, -0.31 DPS) [dungeon]; Gray Woolen Robe (2585, -1.37 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Bright Bracers (3647, -0.11 DPS) [world_drop]; Repurposed Hair Band (281256, -0.15 DPS) [quest]; Mindthrust Bracers (1974, -0.37 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Gnoll Casting Gloves (892, -0.31 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.53 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.95 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 spell_power points (0.89 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Flame Sash (253929, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.00 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (39.2 DPS) | yes | Colorful Kilt (10048, -0.34 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.68 DPS) [crafted]; Silk-threaded Trousers (1929, -1.15 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (1.49 DPS) | yes | Pristine Boots (253889, -0.82 DPS) [crafted]; Red Woolen Boots (4313, -0.86 DPS, sim-verified) [crafted]; Woolen Boots (2583, -1.04 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.2 spell_power points (1.04 DPS) | yes | Volcanic Rock Ring (12053, -0.98 DPS) [world_drop]; The 1 Ring (8350, -1.02 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.7 spell_power points (0.14 DPS) | yes | The 1 Ring (8350, -0.11 DPS) [world]; Volcanic Rock Ring (12053, -0.33 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 1.1 spell_power points (0.23 DPS) | yes | Witching Stave (1484, -0.11 DPS) [dungeon]; Staff of Westfall (2042, -0.11 DPS) [quest]; Lesser Staff of the Spire (1300, -0.46 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.7 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.50 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.26 DPS) [quest]; Deepblaze (279896, -4.38 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 80.4. Weights run: 2.4s. Verify run: 1.3s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.328 ± 0.008, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.988 per %), hit=0.166 ± 0.007 per rating point (10 rating = 1%, 1.664 per %), spell_haste=not significant (0.095 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.764 ± 0.002, fire_power=0.236 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Silk Headband (7050, -0.46 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.69 DPS) [crafted]; Enchanter's Cowl (4322, -3.73 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.0 spell_power points (0.45 DPS) | yes | Crystal Starfire Medallion (5003, -0.15 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.15 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (2.73 DPS) | yes | Death Speaker Mantle (6685, -0.37 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.69 DPS) [quest]; Invoker's Mantle (215365, -0.76 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.14 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Prelacy Cape (7004, -0.23 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.01 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.24 DPS) [dungeon]; Green Silk Armor (7065, -1.03 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.06 DPS) | yes | Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop]; Nightsky Wristbands (6407, -1.61 DPS) [world_drop]; Phoenix Bindings (210781, -1.66 DPS, sim-verified) [crafted] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.60 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.23 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 10.0 spell_power points (2.28 DPS) | yes | Invoker's Cord (215366, -0.31 DPS) [crafted]; Crimson Silk Belt (7055, -0.39 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.45 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.30 DPS) [crafted]; Silk-threaded Trousers (1929, -0.52 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.13 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.3 spell_power points (2.12 DPS) | yes | Acidic Walkers (9454, -0.38 DPS) [dungeon]; Glinteye Slippers (273024, -0.52 DPS) [dungeon]; Spidersilk Boots (4320, -1.61 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.6 spell_power points (1.74 DPS) | yes | Minor Channeling Ring (1449, -0.45 DPS) [quest]; Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon]; Runemender's Seal (273031, -1.06 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.37 DPS) | yes | Minor Channeling Ring (1449, -0.08 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.69 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.22 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.65 DPS) [quest]; Staff of Soran'ruk (15109, -6.39 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 145.2 spell_power points (33.16 DPS) | yes | Wand of Mana Concentration (273043, -1.75 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.02 DPS) [crafted]; Scorching Wand (5213, -3.99 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 132.9. Weights run: 2.2s. Verify run: 1.4s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.612 ± 0.019, crit=0.120 ± 0.004 per rating point (14 rating = 1%, 1.675 per %), hit=0.291 ± 0.018 per rating point (10 rating = 1%, 2.908 per %), spell_haste=not significant (0.522 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.283 ± 0.000, fire_power=0.717 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.44 DPS) | yes | Augural Shroud (2620, -1.01 DPS) [world]; Papal Fez (9431, -1.15 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -2.04 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.9 spell_power points (3.08 DPS) | yes | Triune Amulet (7722, -0.69 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.97 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.13 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.0 spell_power points (3.87 DPS) | yes | Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest]; Green Silken Shoulders (7057, -1.33 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.35 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.53 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.05 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (6.65 DPS) | yes | Dreamweave Vest (10021, -0.56 DPS) [crafted]; Robe of Power (7054, -1.12 DPS) [crafted]; Elemental Raiment (9434, -1.21 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.3 spell_power points (2.41 DPS) | yes | Arcane Runed Bracers (4744, -0.08 DPS) [quest]; Spidertank Oilrag (9448, -0.08 DPS) [dungeon]; Condor Bracers (15864, -0.60 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (5.30 DPS) | yes | Black Mageweave Gloves (10003, -1.41 DPS) [crafted]; Fiery Handwraps (254025, -1.59 DPS) [crafted]; Red Mageweave Gloves (10018, -2.44 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.2 spell_power points (4.19 DPS) | yes | Fiery Cord (254041, +0.00 DPS, sim-verified) [crafted]; Star Belt (4329, -0.82 DPS) [crafted]; Gilded Cord (254037, -0.85 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.3 spell_power points (5.53 DPS) | yes | Crimson Silk Pantaloons (7062, -1.13 DPS, sim-verified) [crafted]; Flame Leggings (253991, -1.82 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.42 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.22 DPS) | yes | Gilded Slippers (254001, -3.29 DPS) [crafted]; Fiery Slippers (254005, -3.48 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.65 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.38 DPS) | yes | Ring of Forlorn Spirits (2043, -1.31 DPS) [quest]; Reedknot Ring (9622, -1.57 DPS) [quest]; Minor Channeling Ring (1449, -1.77 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.1 spell_power points (2.35 DPS) | yes | Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.53 DPS) [quest]; Minor Channeling Ring (1449, -0.73 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (132.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.32 DPS) [dungeon]; Spine of the Devourer (273040, -3.34 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.63 DPS) | yes | Orb of Noh'Orahil (15107, -0.24 DPS) [quest]; Orb of Lorica (11262, -0.49 DPS) [quest]; Orb of Mystic Insight (249394, -0.86 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 155.7 spell_power points (40.33 DPS) | yes | Earthen Rod (9381, -2.89 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.66 DPS) [dungeon]; Twisted Nether Wand (249144, -4.78 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 235.7. Weights run: 2.1s. Verify run: 1.6s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.223 ± 0.023, crit=0.172 ± 0.006 per rating point (14 rating = 1%, 2.409 per %), hit=0.448 ± 0.026 per rating point (10 rating = 1%, 4.484 per %), spell_haste=not significant (-0.047 ± 0.418), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.279 ± 0.000, fire_power=0.721 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.66 DPS) | yes | Red Mageweave Headband (10033, -1.00 DPS) [crafted]; Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.70 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.8 spell_power points (2.49 DPS) | yes | Mindburst Medallion (11196, -0.41 DPS) [quest]; Triune Amulet (7722, -1.01 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.61 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.0 spell_power points (4.83 DPS) | yes | Netherflame Shoulders (254053, -0.84 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.31 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.42 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.35 DPS) | yes | Runecloth Cloak (13860, -1.29 DPS) [crafted]; Cloak of Fire (14134, -1.49 DPS) [crafted]; Cindercloth Cloak (14044, -1.61 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.5 spell_power points (6.66 DPS) | yes | Elemental Raiment (9434, -0.70 DPS) [world_drop]; Robe of the Magi (1716, -0.95 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.98 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.9 spell_power points (3.10 DPS) | yes | Arcane Runed Bracers (4744, -0.55 DPS) [quest]; Spidertank Oilrag (9448, -0.55 DPS) [dungeon]; Phoenix Bindings (210781, -1.22 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (5.36 DPS) | yes | Fiery Gloves (254099, -0.84 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.10 DPS) [vendor]; Black Mageweave Gloves (10003, -1.10 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.2 spell_power points (4.60 DPS) | yes | Fiery Waistcord (254085, -0.15 DPS) [crafted]; Ghostweave Cord (254073, -0.63 DPS) [crafted]; Star Belt (4329, -0.92 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.39 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.54 DPS) [vendor]; Red Mageweave Pants (10009, -0.66 DPS) [crafted]; Cindercloth Leggings (12256, -1.50 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.81 DPS) | yes | Cindercloth Boots (10044, -2.51 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.70 DPS) [vendor]; Fiery Sandals (254111, -3.34 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.69 DPS) | yes | Ring of Forlorn Spirits (2043, -1.42 DPS) [quest]; Snake Hoop (6750, -1.67 DPS) [quest]; Reedknot Ring (9622, -1.70 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.1 spell_power points (3.15 DPS) | yes | Ring of Forlorn Spirits (2043, -1.13 DPS, sim-verified) [quest]; Snake Hoop (6750, -1.13 DPS) [quest]; Reedknot Ring (9622, -1.17 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -3.88 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.19 DPS) [world]; Inventor's Focal Sword (17719, -0.76 DPS) [dungeon]; Hanzo Sword (8190, -27.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (235.7 DPS) | yes | Wand of Allistarj (13065, -3.00 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.40 DPS) [crafted]; Pyric Caduceus (11748, -11.10 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 401.5. Weights run: 2.3s. Verify run: 4.2s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.324 ± 0.026, crit=0.202 ± 0.006 per rating point (14 rating = 1%, 2.835 per %), hit=0.527 ± 0.028 per rating point (10 rating = 1%, 5.266 per %), spell_haste=not significant (0.992 ± 0.383), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.344 ± 0.001, fire_power=0.656 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (401.5 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -6.13 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.2 spell_power points (6.44 DPS) | yes | Glowing Eye of Mordresh (10769, -2.39 DPS) [dungeon]; Triune Amulet (7722, -3.06 DPS) [dungeon]; Mindburst Medallion (11196, -3.08 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.56 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.40 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.52 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (6.74 DPS) | yes | Mageflame Cloak (13007, -1.67 DPS, sim-verified) [world_drop]; Runecloth Cloak (13860, -1.84 DPS) [crafted]; Cindercloth Cloak (14044, -2.04 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.9 spell_power points (20.67 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.86 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.37 DPS) [pvp]; Acumen Robes (17775, -11.79 DPS, sim-verified) [quest] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (401.5 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Netherflame Cuffs (254065, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -7.03 DPS, sim-verified) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (9.30 DPS) | yes | Knight-Lieutenant's Dreadweave Handwraps (227100, +0.00 DPS) [pvp]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Dreamweave Gloves (10019, -1.14 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (401.5 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Fiery Waistcord (254085, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -6.75 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (10.99 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -4.06 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.14 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.05 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.06 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.3 spell_power points (8.56 DPS) | yes | Philanthropist's Ring (281635, -3.65 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.18 DPS) [quest]; Snake Hoop (6750, -5.34 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.49 DPS) | yes | Philanthropist's Ring (281635, -0.58 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.11 DPS) [quest]; Snake Hoop (6750, -2.27 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (401.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (401.5 DPS) | yes | Mark of the Chosen (17774, -5.51 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (401.5 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.80 DPS) [world]; Teebu's Blazing Longsword (1728, -37.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (401.5 DPS) | yes | Stormrager (16997, -0.97 DPS) [quest]; Greater Eternal Wand (249237, -7.21 DPS) [crafted]; Torch of Light (279246, -33.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Earth Warder's Gloves; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Brilliant Wand

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 775.3. Weights run: 7.8s. Verify run: 3.1s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.056 ± 0.024, crit=0.272 ± 0.008 per rating point (14 rating = 1%, 3.810 per %), hit=0.716 ± 0.048 per rating point (10 rating = 1%, 7.164 per %), spell_haste=not significant (1.208 ± 0.600), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.278 ± 0.001, fire_power=0.722 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.08 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.71 DPS) [pvp]; Dreamweave Circlet (10041, -3.24 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.6 spell_power points (7.48 DPS) | yes | Glowing Eye of Mordresh (10769, -3.04 DPS) [dungeon]; Mindburst Medallion (11196, -3.71 DPS) [quest]; Triune Amulet (7722, -5.91 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (14.89 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.16 DPS, sim-verified) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -5.04 DPS) [pvp] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 15.2 spell_power points (9.03 DPS) | yes | Spritecaster Cape (11623, -0.49 DPS) [dungeon]; Cloak of Fire (14134, -3.01 DPS) [crafted]; Cindercloth Cloak (14044, -3.18 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (27.70 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.84 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.15 DPS) [pvp]; Robe of the Magi (1716, -14.03 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.6 spell_power points (6.88 DPS) | yes | Netherflame Cuffs (254065, -1.06 DPS) [crafted]; Phoenix Bindings (210781, -1.29 DPS) [crafted]; Arcane Runed Bracers (4744, -1.52 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.11 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.46 DPS) [pvp]; Dreamweave Gloves (10019, -2.25 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.6 spell_power points (8.67 DPS) | yes | Fiery Waistcord (254085, -0.20 DPS) [crafted]; Ghostweave Cord (254073, -0.33 DPS) [crafted]; Star Belt (4329, -0.93 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (15.49 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -6.05 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.30 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.19 DPS) [crafted]; Omnicast Boots (11822, -1.99 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.4 spell_power points (10.95 DPS) | yes | Philanthropist's Ring (281635, -4.83 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.19 DPS) [quest]; Reedknot Ring (9622, -6.78 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.74 DPS) | yes | Ring of Forlorn Spirits (2043, -2.98 DPS) [quest]; Philanthropist's Ring (281635, -3.13 DPS, sim-verified) [quest]; Reedknot Ring (9622, -3.57 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.27 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.79 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -48.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (775.3 DPS) | yes | Stormrager (16997, -2.32 DPS) [quest]; Greater Eternal Wand (249237, -6.48 DPS) [crafted]; Torch of Light (279246, -53.33 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; back: Mageflame Cloak; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; main_hand: Spire of Hakkar; ranged: Brilliant Wand

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 38.2. Weights run: 2.4s. Verify run: 1.2s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.114 ± 0.003, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.389 per %), hit=0.116 ± 0.004 per rating point (10 rating = 1%, 1.156 per %), spell_haste=0.533 ± 0.050, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.002, fire_power=0.235 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.20 DPS) | yes | Shadow Goggles (4373, -1.08 DPS) [crafted]; Flame Circlet (253953, -2.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.0 spell_power points (1.20 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.11 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.41 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Pearl-clasped Cloak (5542, -0.33 DPS) [crafted]; Mystic Shawl (3449, -0.40 DPS) [quest]; Black Whelp Cloak (7283, -0.44 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 spell_power points (1.11 DPS) | yes | Green Woolen Vest (2582, -0.31 DPS) [crafted]; Bloody Apron (6226, -0.31 DPS) [dungeon]; Gray Woolen Robe (2585, -1.51 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.60 DPS) | yes | Windsong Bangles (263336, -0.40 DPS) [quest]; Featherbead Bracers (15452, -0.49 DPS) [quest]; Owlbeard Bracers (16981, -0.65 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.53 DPS) [crafted]; Apothecary Gloves (10919, -0.60 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 spell_power points (0.89 DPS) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Flame Sash (253929, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.98 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (38.2 DPS) | yes | Colorful Kilt (10048, -0.34 DPS) [crafted]; Embroidered Leggings (270011, -0.54 DPS) [quest]; Silk-threaded Trousers (1929, -1.14 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (1.49 DPS) | yes | Pristine Boots (253889, -0.82 DPS) [crafted]; Red Woolen Boots (4313, -0.92 DPS, sim-verified) [crafted]; Woolen Boots (2583, -1.04 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.7 spell_power points (0.14 DPS) | yes | Volcanic Rock Ring (12053, -0.07 DPS) [world_drop]; The 1 Ring (8350, -0.11 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 0.6 spell_power points (0.11 DPS) | yes | The 1 Ring (8350, -0.09 DPS) [world]; Volcanic Rock Ring (12053, -0.19 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 1.1 spell_power points (0.23 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.09 DPS) [world]; Clear Crystal Rod (16894, -0.11 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.7 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.72 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.26 DPS) [quest]; Greater Magic Wand (11288, -5.01 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 78.9. Weights run: 2.4s. Verify run: 1.3s. 283 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.328 ± 0.008, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.988 per %), hit=0.166 ± 0.007 per rating point (10 rating = 1%, 1.664 per %), spell_haste=not significant (0.095 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.764 ± 0.002, fire_power=0.236 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Silk Headband (7050, -0.46 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.69 DPS) [crafted]; Enchanter's Cowl (4322, -3.64 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.0 spell_power points (0.45 DPS) | yes | Crystal Starfire Medallion (5003, -0.15 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.15 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (2.73 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.69 DPS) [quest]; Invoker's Mantle (215365, -0.76 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.14 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Black Whelp Cloak (7283, -0.46 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.01 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.24 DPS) [dungeon]; Green Silk Armor (7065, -0.81 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.06 DPS) | yes | Tabitha's Cuffs (251486, -1.37 DPS) [quest]; Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop]; Phoenix Bindings (210781, -1.55 DPS, sim-verified) [crafted] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.6 spell_power points (1.74 DPS) | yes | Serpent Gloves (5970, -0.15 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.37 DPS) [world]; Truefaith Gloves (7049, -0.38 DPS) [crafted] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.51 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.54 DPS) [crafted]; Crimson Silk Belt (7055, -0.62 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.30 DPS) [crafted]; Silk-threaded Trousers (1929, -0.52 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.03 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.3 spell_power points (2.12 DPS) | yes | Acidic Walkers (9454, -0.38 DPS) [dungeon]; Glinteye Slippers (273024, -0.52 DPS) [dungeon]; Spidersilk Boots (4320, -1.63 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.6 spell_power points (1.74 DPS) | yes | Solid Gold Band (279845, -0.83 DPS) [quest]; Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon]; Runemender's Seal (273031, -1.06 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.37 DPS) | yes | Solid Gold Band (279845, -0.46 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.69 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.22 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -6.39 DPS) [quest]; Glimmering Staff (249392, -7.40 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 145.2 spell_power points (33.16 DPS) | yes | Wand of Mana Concentration (273043, -2.52 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.02 DPS) [crafted]; Unstable Power Core (279847, -3.45 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Warsong Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 133.1. Weights run: 2.2s. Verify run: 1.4s. 381 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.612 ± 0.019, crit=0.120 ± 0.004 per rating point (14 rating = 1%, 1.675 per %), hit=0.291 ± 0.018 per rating point (10 rating = 1%, 2.908 per %), spell_haste=not significant (0.522 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.283 ± 0.000, fire_power=0.717 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.44 DPS) | yes | Papal Fez (9431, -0.67 DPS) [dungeon]; Augural Shroud (2620, -1.01 DPS) [world]; Whitemane's Chapeau (7720, -2.04 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.9 spell_power points (3.08 DPS) | yes | Triune Amulet (7722, -1.01 DPS) [dungeon]; Necklace of Calisea (1714, -1.97 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.13 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.0 spell_power points (3.87 DPS) | yes | Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest]; Green Silken Shoulders (7057, -0.71 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.35 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.53 DPS) [crafted]; Hillman's Cloak (3719, -1.05 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (6.65 DPS) | yes | Dreamweave Vest (10021, -0.56 DPS) [crafted]; Robe of Power (7054, -1.12 DPS) [crafted]; Elemental Raiment (9434, -1.21 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.3 spell_power points (2.41 DPS) | yes | Spidertank Oilrag (9448, -0.08 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.11 DPS) [quest]; Condor Bracers (15864, -0.60 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.4 spell_power points (5.30 DPS) | yes | Black Mageweave Gloves (10003, -1.41 DPS) [crafted]; Fiery Handwraps (254025, -1.59 DPS) [crafted]; Red Mageweave Gloves (10018, -1.74 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | sim-verified (133.1 DPS) | yes | Star Belt (4329, -0.50 DPS) [crafted]; Gilded Cord (254037, -0.53 DPS) [crafted]; Deathmage Sash (10771, -1.46 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.3 spell_power points (5.53 DPS) | yes | Crimson Silk Pantaloons (7062, -1.35 DPS, sim-verified) [crafted]; Flame Leggings (253991, -1.82 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.42 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.22 DPS) | yes | Fiery Slippers (254005, -2.95 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.29 DPS) [crafted]; Acidic Walkers (9454, -3.65 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.38 DPS) | yes | Reedknot Ring (9622, -1.57 DPS) [quest]; Ring of Power Regulation (273030, -1.83 DPS) [dungeon]; Voodoo Band (1996, -2.27 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.1 spell_power points (2.35 DPS) | yes | Reedknot Ring (9622, -0.53 DPS) [quest]; Ring of Power Regulation (273030, -0.79 DPS) [dungeon]; Voodoo Band (1996, -1.24 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.32 DPS) [dungeon]; Spine of the Devourer (273040, -3.34 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.63 DPS) | yes | Orb of Noh'Orahil (15107, -0.83 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -0.86 DPS) [crafted]; Rod of Molten Fire (2565, -1.21 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 155.7 spell_power points (40.33 DPS) | yes | Earthen Rod (9381, -3.17 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.66 DPS) [dungeon]; Twisted Nether Wand (249144, -4.78 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 381, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 233.6. Weights run: 2.1s. Verify run: 1.6s. 492 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.223 ± 0.023, crit=0.172 ± 0.006 per rating point (14 rating = 1%, 2.409 per %), hit=0.448 ± 0.026 per rating point (10 rating = 1%, 4.484 per %), spell_haste=not significant (-0.047 ± 0.418), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.279 ± 0.000, fire_power=0.721 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.66 DPS) | yes | Red Mageweave Headband (10033, -0.85 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.70 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.8 spell_power points (2.49 DPS) | yes | Mindburst Medallion (11196, -0.41 DPS) [quest]; Triune Amulet (7722, -1.42 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.61 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.0 spell_power points (4.83 DPS) | yes | Netherflame Shoulders (254053, -0.84 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.31 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.42 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.35 DPS) | yes | Cindercloth Cloak (14044, -1.19 DPS) [crafted]; Runecloth Cloak (13860, -1.29 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.80 DPS, sim-verified) [quest] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.5 spell_power points (6.66 DPS) | yes | Robe of the Magi (1716, -0.04 DPS) [world_drop]; Elemental Raiment (9434, -0.70 DPS) [world_drop]; Dreamweave Vest (10021, -0.98 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.9 spell_power points (3.10 DPS) | yes | Spidertank Oilrag (9448, -0.55 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.67 DPS) [crafted]; Phoenix Bindings (210781, -1.04 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (5.36 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.10 DPS) [vendor]; Black Mageweave Gloves (10003, -1.10 DPS) [crafted]; Fiery Gloves (254099, -1.27 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.2 spell_power points (4.60 DPS) | yes | Fiery Waistcord (254085, -0.15 DPS) [crafted]; Ghostweave Cord (254073, -0.63 DPS) [crafted]; Star Belt (4329, -0.92 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.39 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.54 DPS) [vendor]; Red Mageweave Pants (10009, -0.66 DPS) [crafted]; Cindercloth Leggings (12256, -1.50 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.81 DPS) | yes | Cindercloth Boots (10044, -2.51 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.70 DPS) [vendor]; Fiery Sandals (254111, -3.26 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.69 DPS) | yes | Snake Hoop (6750, -1.67 DPS) [quest]; Reedknot Ring (9622, -1.70 DPS) [quest]; Ring of Power Regulation (273030, -1.99 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.1 spell_power points (3.15 DPS) | yes | Reedknot Ring (9622, -1.17 DPS) [quest]; Snake Hoop (6750, -1.28 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -1.45 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.25 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -2.63 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.19 DPS) [world]; Inventor's Focal Sword (17719, -0.76 DPS) [dungeon]; Hanzo Sword (8190, -26.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (233.6 DPS) | yes | Wand of Allistarj (13065, -3.00 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.40 DPS) [crafted]; Pyric Caduceus (11748, -9.62 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 492, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 398.3. Weights run: 2.3s. Verify run: 2.7s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.324 ± 0.026, crit=0.202 ± 0.006 per rating point (14 rating = 1%, 2.835 per %), hit=0.527 ± 0.028 per rating point (10 rating = 1%, 5.266 per %), spell_haste=not significant (0.992 ± 0.383), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.344 ± 0.001, fire_power=0.656 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.41 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -2.34 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.2 spell_power points (6.44 DPS) | yes | Glowing Eye of Mordresh (10769, -2.39 DPS) [dungeon]; Triune Amulet (7722, -3.06 DPS) [dungeon]; Mindburst Medallion (11196, -3.08 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.56 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.81 DPS, sim-verified) [crafted]; Champion's Dreadweave Spaulders (227092, -2.52 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (6.74 DPS) | yes | Mageflame Cloak (13007, -0.92 DPS) [world_drop]; Runecloth Cloak (13860, -1.84 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.86 DPS, sim-verified) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.9 spell_power points (20.67 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.86 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.37 DPS) [pvp]; Acumen Robes (17775, -10.71 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 14.2 spell_power points (6.02 DPS) | yes | Netherflame Cuffs (254065, -1.74 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -2.10 DPS) [crafted]; Spidertank Oilrag (9448, -2.21 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (9.30 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Blood Guard's Dreadweave Handwraps (227099, +0.00 DPS) [pvp]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Fiery Waistcord (254085, -0.07 DPS) [crafted]; Ghostweave Cord (254073, -0.65 DPS) [crafted]; Satyrmane Sash (17755, -3.64 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (10.99 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -6.01 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.14 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.05 DPS) [dungeon]; Blood Guard's Dreadweave Walkers (227098, -0.76 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.3 spell_power points (8.56 DPS) | yes | Philanthropist's Ring (281635, -3.65 DPS) [quest]; Snake Hoop (6750, -5.34 DPS) [quest]; Reedknot Ring (9622, -5.61 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.49 DPS) | yes | Philanthropist's Ring (281635, -2.05 DPS, sim-verified) [quest]; Snake Hoop (6750, -2.27 DPS) [quest]; Reedknot Ring (9622, -2.54 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.80 DPS) [world]; Teebu's Blazing Longsword (1728, -32.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (+31.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormrager (16997, -0.97 DPS) [quest]; Greater Eternal Wand (249237, -7.21 DPS) [crafted]; Torch of Light (279246, -31.32 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Wisdom of the Timbermaw; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Brilliant Wand

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 774.2. Weights run: 7.8s. Verify run: 3.1s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.056 ± 0.024, crit=0.272 ± 0.008 per rating point (14 rating = 1%, 3.810 per %), hit=0.716 ± 0.048 per rating point (10 rating = 1%, 7.164 per %), spell_haste=not significant (1.208 ± 0.600), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.278 ± 0.001, fire_power=0.722 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.08 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.71 DPS) [pvp]; Dreamweave Circlet (10041, -3.24 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.6 spell_power points (7.48 DPS) | yes | Glowing Eye of Mordresh (10769, -3.04 DPS) [dungeon]; Mindburst Medallion (11196, -3.71 DPS) [quest]; Triune Amulet (7722, -5.09 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (14.89 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.07 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -5.04 DPS) [pvp] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 15.2 spell_power points (9.03 DPS) | yes | Spritecaster Cape (11623, -0.49 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.58 DPS) [quest]; Cloak of Fire (14134, -3.01 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (27.70 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.84 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.15 DPS) [pvp]; Robe of the Magi (1716, -14.78 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.6 spell_power points (6.88 DPS) | yes | Netherflame Cuffs (254065, -1.06 DPS) [crafted]; Phoenix Bindings (210781, -1.29 DPS) [crafted]; Spidertank Oilrag (9448, -1.52 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.11 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.46 DPS) [pvp]; Dreamweave Gloves (10019, -2.68 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.6 spell_power points (8.67 DPS) | yes | Fiery Waistcord (254085, -0.20 DPS) [crafted]; Ghostweave Cord (254073, -0.33 DPS) [crafted]; Star Belt (4329, -0.93 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (15.49 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -4.60 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.30 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.19 DPS) [crafted]; Omnicast Boots (11822, -1.99 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.4 spell_power points (10.95 DPS) | yes | Philanthropist's Ring (281635, -4.83 DPS) [quest]; Reedknot Ring (9622, -6.78 DPS) [quest]; Snake Hoop (6750, -7.21 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.74 DPS) | yes | Philanthropist's Ring (281635, -1.62 DPS) [quest]; Reedknot Ring (9622, -3.57 DPS) [quest]; Snake Hoop (6750, -4.00 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.27 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.79 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -47.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (774.2 DPS) | yes | Stormrager (16997, -2.32 DPS) [quest]; Greater Eternal Wand (249237, -6.48 DPS) [crafted]; Torch of Light (279246, -54.25 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; back: Mageflame Cloak; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; main_hand: Spire of Hakkar; ranged: Brilliant Wand

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

