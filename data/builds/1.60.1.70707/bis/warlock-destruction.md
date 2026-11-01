# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 38.3. Weights run: 2.3s. Verify run: 1.2s. 161 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.046 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.425 per %), hit=0.081 ± 0.004 per rating point (10 rating = 1%, 0.814 per %), spell_haste=0.277 ± 0.041, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.762 ± 0.002, fire_power=0.238 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.25 DPS) | yes | Flame Circlet (253953, -2.31 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.04 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.83 DPS) | yes | Black Whelp Cloak (7283, -0.38 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.42 DPS) [crafted]; Reinforced Linen Cape (2580, -0.62 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.04 DPS) | yes | Green Woolen Vest (2582, -0.21 DPS) [crafted]; Bloody Apron (6226, -0.21 DPS) [dungeon]; Gray Woolen Robe (2585, -1.47 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.21 DPS) | yes | Ivycloth Bracelets (9793, -0.24 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.46 DPS) | yes | Gnoll Casting Gloves (892, -0.34 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.62 DPS) [crafted]; Phoenix Gloves (4331, -1.01 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Flame Sash (253929, -0.49 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.79 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.21 DPS) [crafted]; Heavy Woolen Pants (4316, -0.62 DPS) [crafted]; Silk-threaded Trousers (1929, -1.29 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.46 DPS) | yes | Red Woolen Boots (4313, -0.68 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.83 DPS) [crafted]; Soft-soled Linen Boots (4312, -1.04 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.04 DPS) | yes | Ring of the Shadow (1462, -1.37 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.62 DPS) | yes | Ring of the Shadow (1462, -0.91 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 89 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.10 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (1.04 DPS) | yes | Antipodean Rod (2879, -1.05 DPS, sim-verified) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 110.2 spell_power points (22.93 DPS) | yes | Skycaller (12984, -0.73 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.24 DPS) [quest]; Deepblaze (279896, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 161, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 74.8. Weights run: 2.3s. Verify run: 1.1s. 298 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.178 ± 0.005, crit=0.087 ± 0.002 per rating point (14 rating = 1%, 1.214 per %), hit=0.136 ± 0.006 per rating point (10 rating = 1%, 1.360 per %), spell_haste=not significant (-0.019 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.757 ± 0.002, fire_power=0.243 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.57 DPS) | yes | Filigreed Pristine Circlet (253975, -0.70 DPS) [crafted]; Enchanter's Cowl (4322, -0.75 DPS) [crafted]; Silk Headband (7050, -3.67 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 1.1 spell_power points (0.25 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.6 spell_power points (2.47 DPS) | yes | Invoker's Mantle (215365, -0.63 DPS) [crafted]; Death Speaker Mantle (6685, -0.68 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.70 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.17 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Prelacy Cape (7004, -0.23 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.03 DPS) | yes | Green Silk Armor (7065, -0.39 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.44 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.67 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.10 DPS) | yes | Glowing Magical Bracelets (13106, -1.77 DPS) [world_drop]; Phoenix Bindings (210781, -1.80 DPS, sim-verified) [crafted]; Nightsky Wristbands (6407, -1.85 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.63 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.5 spell_power points (2.22 DPS) | yes | Invoker's Cord (215366, -0.38 DPS) [crafted]; Crimson Silk Belt (7055, -0.53 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.86 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.80 DPS) | yes | Pristine Leggings (253987, -0.88 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.15 DPS) [crafted]; Silk-threaded Trousers (1929, -1.17 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.2 spell_power points (1.92 DPS) | yes | Glinteye Slippers (273024, -0.29 DPS) [dungeon]; Acidic Walkers (9454, -0.43 DPS) [dungeon]; Spidersilk Boots (4320, -2.23 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.9 spell_power points (1.61 DPS) | yes | Minor Channeling Ring (1449, -0.36 DPS) [quest]; Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Runemender's Seal (273031, -0.91 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.40 DPS) | yes | Minor Channeling Ring (1449, -0.15 DPS) [quest]; Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon]; Runemender's Seal (273031, -0.70 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.40 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.73 DPS) [quest]; Staff of Soran'ruk (15109, -6.47 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 141.5 spell_power points (33.03 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.86 DPS) [crafted]; Scorching Wand (5213, -3.84 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 127.4. Weights run: 1.9s. Verify run: 1.2s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.621 ± 0.017, crit=0.128 ± 0.004 per rating point (14 rating = 1%, 1.786 per %), hit=0.247 ± 0.015 per rating point (10 rating = 1%, 2.467 per %), spell_haste=not significant (0.172 ± 0.200), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.297 ± 0.000, fire_power=0.703 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.21 DPS) | yes | Papal Fez (9431, -0.61 DPS) [dungeon]; Augural Shroud (2620, -0.94 DPS) [world]; Whitemane's Chapeau (7720, -1.93 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.0 spell_power points (2.97 DPS) | yes | Triune Amulet (7722, -0.66 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.89 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.05 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.1 spell_power points (3.74 DPS) | yes | Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest]; Green Silken Shoulders (7057, -1.61 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.26 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.52 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.02 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (6.39 DPS) | yes | Dreamweave Vest (10021, -0.53 DPS) [crafted]; Robe of Power (7054, -1.06 DPS) [crafted]; Elemental Raiment (9434, -1.17 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.1 spell_power points (2.27 DPS) | yes | Arcane Runed Bracers (4744, -0.03 DPS) [quest]; Spidertank Oilrag (9448, -0.03 DPS) [dungeon]; Condor Bracers (15864, -0.53 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (5.08 DPS) | yes | Black Mageweave Gloves (10003, -1.36 DPS) [crafted]; Fiery Handwraps (254025, -1.56 DPS) [crafted]; Red Mageweave Gloves (10018, -2.30 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 16.3 spell_power points (4.05 DPS) | yes | Fiery Cord (254041, +0.00 DPS, sim-verified) [crafted]; Star Belt (4329, -0.82 DPS) [crafted]; Gilded Cord (254037, -0.83 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.4 spell_power points (5.32 DPS) | yes | Flame Leggings (253991, -1.80 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.92 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.35 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.96 DPS) | yes | Gilded Slippers (254001, -3.14 DPS) [crafted]; Acidic Walkers (9454, -3.48 DPS) [dungeon]; Fiery Slippers (254005, -4.32 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.25 DPS) | yes | Ring of Forlorn Spirits (2043, -1.27 DPS) [quest]; Reedknot Ring (9622, -1.52 DPS) [quest]; Minor Channeling Ring (1449, -1.70 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.1 spell_power points (2.26 DPS) | yes | Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.52 DPS) [quest]; Minor Channeling Ring (1449, -0.71 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (127.4 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.31 DPS) [dungeon]; Spine of the Devourer (273040, -3.22 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.48 DPS) | yes | Orb of Lorica (11262, -0.44 DPS) [quest]; Orb of Noh'Orahil (15107, -0.73 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -0.81 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 162.4 spell_power points (40.30 DPS) | yes | Earthen Rod (9381, -3.97 DPS) [dungeon]; Umbral Wand (5216, -4.63 DPS) [dungeon]; Twisted Nether Wand (249144, -4.82 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 229.5. Weights run: 2.0s. Verify run: 1.3s. 507 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.312 ± 0.017, crit=0.185 ± 0.005 per rating point (14 rating = 1%, 2.594 per %), hit=0.300 ± 0.018 per rating point (10 rating = 1%, 3.000 per %), spell_haste=-1.864 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.250 ± 0.000, fire_power=0.750 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.98 DPS) | yes | Red Mageweave Headband (10033, -0.59 DPS) [crafted]; Dreamweave Circlet (10041, -0.96 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.00 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.5 spell_power points (3.16 DPS) | yes | Triune Amulet (7722, -0.50 DPS) [dungeon]; Mindburst Medallion (11196, -0.54 DPS) [quest]; Horizon Choker (13085, -1.71 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.6 spell_power points (6.19 DPS) | yes | Netherflame Shoulders (254053, -1.12 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.73 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.93 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (5.28 DPS) | yes | Runecloth Cloak (13860, -1.46 DPS) [crafted]; Cloak of Fire (14134, -1.79 DPS) [crafted]; Cindercloth Cloak (14044, -1.87 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.2 spell_power points (8.39 DPS) | yes | Robe of the Magi (1716, -0.45 DPS) [world_drop]; Elemental Raiment (9434, -1.41 DPS) [world_drop]; Dreamweave Vest (10021, -1.47 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 11.9 spell_power points (3.97 DPS) | yes | Phoenix Bindings (210781, -0.73 DPS) [crafted]; Nethergeld Cuffs (254061, -0.92 DPS) [crafted]; Arcane Runed Bracers (4744, -0.98 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (6.40 DPS) | yes | Fiery Gloves (254099, -0.62 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.14 DPS) [vendor]; Black Mageweave Gloves (10003, -1.41 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.1 spell_power points (5.69 DPS) | yes | Fiery Waistcord (254085, -0.02 DPS) [crafted]; Ghostweave Cord (254073, -1.04 DPS) [crafted]; Star Belt (4329, -1.37 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.32 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.22 DPS) [vendor]; Red Mageweave Pants (10009, -0.42 DPS) [crafted]; Cindercloth Leggings (12256, -1.58 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.98 DPS) | yes | Fiery Sandals (254111, -1.99 DPS, sim-verified) [crafted]; Cindercloth Boots (10044, -2.74 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.39 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.32 DPS) | yes | Ring of Forlorn Spirits (2043, -1.66 DPS) [quest]; Snake Hoop (6750, -1.81 DPS) [quest]; Reedknot Ring (9622, -2.00 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.6 spell_power points (3.84 DPS) | yes | Ring of Forlorn Spirits (2043, -1.14 DPS, sim-verified) [quest]; Snake Hoop (6750, -1.33 DPS) [quest]; Reedknot Ring (9622, -1.52 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (229.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (229.5 DPS) | yes | Mark of the Chosen (17774, -2.21 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (229.5 DPS) | yes | Spire of Hakkar (10844, -0.58 DPS) [world]; Spellshifter Rod (9527, -0.85 DPS) [quest]; Hanzo Sword (8190, -29.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 157.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.09 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 369.7. Weights run: 2.0s. Verify run: 2.3s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.242 ± 0.021, crit=0.222 ± 0.006 per rating point (14 rating = 1%, 3.105 per %), hit=0.418 ± 0.026 per rating point (10 rating = 1%, 4.179 per %), spell_haste=not significant (-0.747 ± 0.353), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.368 ± 0.001, fire_power=0.632 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.89 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -3.51 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.4 spell_power points (5.82 DPS) | yes | Glowing Eye of Mordresh (10769, -2.21 DPS) [dungeon]; Triune Amulet (7722, -2.59 DPS) [dungeon]; Mindburst Medallion (11196, -2.81 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.08 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.71 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.72 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (6.23 DPS) | yes | Runecloth Cloak (13860, -1.82 DPS) [crafted]; Cindercloth Cloak (14044, -2.14 DPS) [crafted]; Mageflame Cloak (13007, -2.24 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (19.43 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.18 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.40 DPS) [pvp]; Acumen Robes (17775, -10.97 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.4 spell_power points (5.41 DPS) | yes | Arcane Runed Bracers (4744, -1.78 DPS) [quest]; Spidertank Oilrag (9448, -1.78 DPS) [dungeon]; Netherflame Cuffs (254065, -1.85 DPS, sim-verified) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.87 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.01 DPS) [pvp]; Dreamweave Gloves (10019, -1.22 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (369.7 DPS) | yes | Fiery Waistcord (254085, -0.08 DPS) [crafted]; Ghostweave Cord (254073, -0.16 DPS) [crafted]; Satyrmane Sash (17755, -3.83 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (10.49 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -4.59 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.68 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.55 DPS) [pvp]; Omnicast Boots (11822, -1.57 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (7.94 DPS) | yes | Philanthropist's Ring (281635, -3.42 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.72 DPS) [quest]; Snake Hoop (6750, -5.04 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.24 DPS) | yes | Philanthropist's Ring (281635, -0.72 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.02 DPS) [quest]; Snake Hoop (6750, -2.34 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -5.08 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.36 DPS) [world]; Teebu's Blazing Longsword (1728, -35.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 196.3 spell_power points (79.18 DPS) | yes | Brilliant Wand (249385, -15.29 DPS) [crafted]; Stormrager (16997, -16.00 DPS) [quest]; Greater Eternal Wand (249237, -22.22 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Wisdom of the Timbermaw; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 725.7. Weights run: 2.5s. Verify run: 3.0s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.180 ± 0.026, crit=0.318 ± 0.008 per rating point (14 rating = 1%, 4.446 per %), hit=0.692 ± 0.036 per rating point (10 rating = 1%, 6.918 per %), spell_haste=not significant (0.287 ± 0.495), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.308 ± 0.001, fire_power=0.692 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (15.59 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.43 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.8 spell_power points (7.97 DPS) | yes | Glowing Eye of Mordresh (10769, -3.10 DPS) [dungeon]; Triune Amulet (7722, -3.35 DPS) [dungeon]; Mindburst Medallion (11196, -3.88 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (14.44 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.59 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.1 spell_power points (8.71 DPS) | yes | Mageflame Cloak (13007, -0.31 DPS) [world_drop]; Cindercloth Cloak (14044, -2.68 DPS) [crafted]; Runecloth Cloak (13860, -2.68 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (27.50 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.53 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.98 DPS) [pvp]; Robe of the Magi (1716, -13.76 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.8 spell_power points (7.39 DPS) | yes | Netherflame Cuffs (254065, -1.47 DPS) [crafted]; Phoenix Bindings (210781, -2.19 DPS) [crafted]; Arcane Runed Bracers (4744, -2.19 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (12.70 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.16 DPS) [pvp]; Gloves of Spell Mastery (14146, -1.33 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.8 spell_power points (9.12 DPS) | yes | Fiery Waistcord (254085, -0.59 DPS) [crafted]; Ghostweave Cord (254073, -1.04 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.32 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (15.01 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -8.20 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.86 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.06 DPS) [dungeon]; Venomspew Footpads (275606, -1.15 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.3 spell_power points (11.12 DPS) | yes | Philanthropist's Ring (281635, -4.83 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.50 DPS) [quest]; Reedknot Ring (9622, -7.08 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.51 DPS) | yes | Philanthropist's Ring (281635, -1.21 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.89 DPS) [quest]; Reedknot Ring (9622, -3.46 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (725.7 DPS) | yes | Uther's Strength (11302, -0.13 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (725.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (725.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.09 DPS) [world]; Teebu's Blazing Longsword (1728, -50.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 138.9 spell_power points (80.22 DPS) | yes | Brilliant Wand (249385, -3.55 DPS, sim-verified) [crafted]; Stormrager (16997, -17.02 DPS) [quest]; Greater Eternal Wand (249237, -21.70 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 35.7. Weights run: 2.3s. Verify run: 1.3s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.046 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.425 per %), hit=0.081 ± 0.004 per rating point (10 rating = 1%, 0.814 per %), spell_haste=0.277 ± 0.041, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.762 ± 0.002, fire_power=0.238 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.25 DPS) | yes | Flame Circlet (253953, -2.16 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.04 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.86 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.83 DPS) | yes | Black Whelp Cloak (7283, -0.39 DPS, sim-verified) [crafted]; Mystic Shawl (3449, -0.42 DPS) [quest]; Pearl-clasped Cloak (5542, -0.42 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.04 DPS) | yes | Green Woolen Vest (2582, -0.21 DPS) [crafted]; Bloody Apron (6226, -0.21 DPS) [dungeon]; Gray Woolen Robe (2585, -1.44 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.62 DPS) | yes | Owlbeard Bracers (16981, -0.42 DPS) [quest]; Windsong Bangles (263336, -0.45 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.46 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.62 DPS) [quest]; Pristine Gloves (253913, -0.62 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Flame Sash (253929, -0.49 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.27 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.21 DPS) [crafted]; Embroidered Leggings (270011, -0.42 DPS) [quest]; Silk-threaded Trousers (1929, -2.18 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.46 DPS) | yes | Pristine Boots (253889, -0.83 DPS) [crafted]; Red Woolen Boots (4313, -1.03 DPS, sim-verified) [crafted]; Soft-soled Linen Boots (4312, -1.04 DPS) [crafted] |
| finger1 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.62 DPS) | yes | - |
| finger2 | - | - |  |  |  |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Evocator's Blade (2567, -0.63 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 110.2 spell_power points (22.93 DPS) | yes | Skycaller (12984, -0.73 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.24 DPS) [quest]; Greater Magic Wand (11288, -5.02 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 73.2. Weights run: 2.3s. Verify run: 1.1s. 289 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.178 ± 0.005, crit=0.087 ± 0.002 per rating point (14 rating = 1%, 1.214 per %), hit=0.136 ± 0.006 per rating point (10 rating = 1%, 1.360 per %), spell_haste=not significant (-0.019 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.757 ± 0.002, fire_power=0.243 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.57 DPS) | yes | Filigreed Pristine Circlet (253975, -0.70 DPS) [crafted]; Enchanter's Cowl (4322, -0.75 DPS) [crafted]; Silk Headband (7050, -4.13 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 1.1 spell_power points (0.25 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.6 spell_power points (2.47 DPS) | yes | Death Speaker Mantle (6685, -0.62 DPS) [dungeon]; Invoker's Mantle (215365, -0.63 DPS) [crafted]; Chestnut Mantle (17695, -1.24 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.17 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Black Whelp Cloak (7283, -0.47 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.03 DPS) | yes | Beguiler Robes (7728, -0.44 DPS) [dungeon]; Green Silk Armor (7065, -0.50 DPS, sim-verified) [crafted]; Mechbuilder's Overalls (9508, -0.67 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.10 DPS) | yes | Tabitha's Cuffs (251486, -1.40 DPS) [quest]; Glowing Magical Bracelets (13106, -1.77 DPS) [world_drop]; Phoenix Bindings (210781, -2.13 DPS, sim-verified) [crafted] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.63 DPS) | yes | Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted]; Jutebraid Gloves (10654, -0.49 DPS, sim-verified) [quest] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.57 DPS) | yes | Belt of Arugal (6392, -0.34 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.70 DPS) [dungeon]; Invoker's Cord (215366, -0.73 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.80 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -1.15 DPS) [crafted]; Silk-threaded Trousers (1929, -1.17 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.2 spell_power points (1.92 DPS) | yes | Glinteye Slippers (273024, -0.29 DPS) [dungeon]; Acidic Walkers (9454, -0.43 DPS) [dungeon]; Spidersilk Boots (4320, -2.61 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.9 spell_power points (1.61 DPS) | yes | Solid Gold Band (279845, -0.67 DPS) [quest]; Electrocutioner Lagnut (9447, -0.91 DPS) [dungeon]; Runemender's Seal (273031, -0.91 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.40 DPS) | yes | Solid Gold Band (279845, -0.47 DPS) [quest]; Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon]; Runemender's Seal (273031, -0.70 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.40 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -6.47 DPS) [quest]; Glimmering Staff (249392, -7.94 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 141.5 spell_power points (33.03 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.86 DPS) [crafted]; Unstable Power Core (279847, -3.32 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 289, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 127.0. Weights run: 1.9s. Verify run: 1.2s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.621 ± 0.017, crit=0.128 ± 0.004 per rating point (14 rating = 1%, 1.786 per %), hit=0.247 ± 0.015 per rating point (10 rating = 1%, 2.467 per %), spell_haste=not significant (0.172 ± 0.200), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.297 ± 0.000, fire_power=0.703 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.21 DPS) | yes | Papal Fez (9431, -0.61 DPS) [dungeon]; Augural Shroud (2620, -0.94 DPS) [world]; Whitemane's Chapeau (7720, -1.93 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.0 spell_power points (2.97 DPS) | yes | Triune Amulet (7722, -0.98 DPS) [dungeon]; Necklace of Calisea (1714, -1.89 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.05 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.1 spell_power points (3.74 DPS) | yes | Bloodmage Mantle (7684, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest]; Green Silken Shoulders (7057, -1.07 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.26 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.52 DPS) [crafted]; Hillman's Cloak (3719, -1.02 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (6.39 DPS) | yes | Dreamweave Vest (10021, -0.53 DPS) [crafted]; Robe of Power (7054, -1.06 DPS) [crafted]; Elemental Raiment (9434, -1.17 DPS) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.1 spell_power points (2.27 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -0.04 DPS) [quest]; Condor Bracers (15864, -0.53 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (5.08 DPS) | yes | Black Mageweave Gloves (10003, -1.36 DPS) [crafted]; Fiery Handwraps (254025, -1.56 DPS) [crafted]; Red Mageweave Gloves (10018, -1.85 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | sim-verified (127.0 DPS) | yes | Star Belt (4329, -0.45 DPS) [crafted]; Gilded Cord (254037, -0.46 DPS) [crafted]; Deathmage Sash (10771, -1.52 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.4 spell_power points (5.32 DPS) | yes | Crimson Silk Pantaloons (7062, -1.75 DPS, sim-verified) [crafted]; Flame Leggings (253991, -1.80 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.35 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.96 DPS) | yes | Gilded Slippers (254001, -3.14 DPS) [crafted]; Acidic Walkers (9454, -3.48 DPS) [dungeon]; Fiery Slippers (254005, -3.68 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.25 DPS) | yes | Reedknot Ring (9622, -1.52 DPS) [quest]; Ring of Power Regulation (273030, -1.76 DPS) [dungeon]; Voodoo Band (1996, -2.17 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.1 spell_power points (2.26 DPS) | yes | Reedknot Ring (9622, -0.52 DPS) [quest]; Ring of Power Regulation (273030, -0.77 DPS) [dungeon]; Voodoo Band (1996, -1.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.31 DPS) [dungeon]; Spine of the Devourer (273040, -3.22 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.48 DPS) | yes | Orb of Noh'Orahil (15107, -0.61 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -0.81 DPS) [crafted]; Rod of Molten Fire (2565, -1.21 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 162.4 spell_power points (40.30 DPS) | yes | Earthen Rod (9381, -3.97 DPS) [dungeon]; Umbral Wand (5216, -4.63 DPS) [dungeon]; Twisted Nether Wand (249144, -4.82 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Phoenix Bindings; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 228.9. Weights run: 2.0s. Verify run: 1.3s. 498 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.312 ± 0.017, crit=0.185 ± 0.005 per rating point (14 rating = 1%, 2.594 per %), hit=0.300 ± 0.018 per rating point (10 rating = 1%, 3.000 per %), spell_haste=-1.864 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.250 ± 0.000, fire_power=0.750 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.98 DPS) | yes | Red Mageweave Headband (10033, -0.59 DPS) [crafted]; Dreamweave Circlet (10041, -0.96 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.00 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.5 spell_power points (3.16 DPS) | yes | Triune Amulet (7722, -0.50 DPS) [dungeon]; Mindburst Medallion (11196, -0.54 DPS) [quest]; Horizon Choker (13085, -1.71 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.6 spell_power points (6.19 DPS) | yes | Netherflame Shoulders (254053, +0.00 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.73 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.93 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (5.28 DPS) | yes | Deep Woodlands Cloak (19121, -1.00 DPS, sim-verified) [quest]; Cindercloth Cloak (14044, -1.21 DPS) [crafted]; Runecloth Cloak (13860, -1.46 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.2 spell_power points (8.39 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.41 DPS) [world_drop]; Dreamweave Vest (10021, -1.47 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 11.9 spell_power points (3.97 DPS) | yes | Phoenix Bindings (210781, -0.73 DPS) [crafted]; Nethergeld Cuffs (254061, -0.92 DPS) [crafted]; Spidertank Oilrag (9448, -0.98 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (6.40 DPS) | yes | Fiery Gloves (254099, -0.62 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.14 DPS) [vendor]; Black Mageweave Gloves (10003, -1.41 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.1 spell_power points (5.69 DPS) | yes | Fiery Waistcord (254085, -0.02 DPS) [crafted]; Ghostweave Cord (254073, -1.04 DPS) [crafted]; Star Belt (4329, -1.37 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.32 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.22 DPS) [vendor]; Red Mageweave Pants (10009, -0.42 DPS) [crafted]; Cindercloth Leggings (12256, -1.58 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.98 DPS) | yes | Cindercloth Boots (10044, -2.74 DPS) [crafted]; Fiery Sandals (254111, -2.78 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.39 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.32 DPS) | yes | Snake Hoop (6750, -1.81 DPS) [quest]; Reedknot Ring (9622, -2.00 DPS) [quest]; Ring of Power Regulation (273030, -2.33 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.6 spell_power points (3.84 DPS) | yes | Snake Hoop (6750, -1.40 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.52 DPS) [quest]; Ring of Power Regulation (273030, -1.85 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (228.9 DPS) | yes | Rune of the Guard Captain (19120, -1.30 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.06 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (228.9 DPS) | yes | Frozen Heart of the Mountain (249469, -1.10 DPS) [crafted]; Rune of the Guard Captain (19120, -2.73 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (228.9 DPS) | yes | Spire of Hakkar (10844, -0.58 DPS) [world]; Spellshifter Rod (9527, -0.85 DPS) [quest]; Hanzo Sword (8190, -27.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 157.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.09 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 498, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 365.6. Weights run: 2.0s. Verify run: 2.3s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.242 ± 0.021, crit=0.222 ± 0.006 per rating point (14 rating = 1%, 3.105 per %), hit=0.418 ± 0.026 per rating point (10 rating = 1%, 4.179 per %), spell_haste=not significant (-0.747 ± 0.353), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.368 ± 0.001, fire_power=0.632 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.89 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -2.51 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.4 spell_power points (5.82 DPS) | yes | Glowing Eye of Mordresh (10769, -2.21 DPS) [dungeon]; Triune Amulet (7722, -2.59 DPS) [dungeon]; Mindburst Medallion (11196, -2.81 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.08 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.49 DPS, sim-verified) [crafted]; Champion's Dreadweave Spaulders (227092, -2.72 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (6.23 DPS) | yes | Deep Woodlands Cloak (19121, -0.51 DPS) [quest]; Mageflame Cloak (13007, -0.88 DPS) [world_drop]; Runecloth Cloak (13860, -1.82 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (19.43 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.18 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.40 DPS) [pvp]; Acumen Robes (17775, -11.06 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.4 spell_power points (5.41 DPS) | yes | Netherflame Cuffs (254065, -1.41 DPS) [crafted]; Spidertank Oilrag (9448, -1.78 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.91 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.87 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.01 DPS) [pvp]; Dreamweave Gloves (10019, -2.21 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (365.6 DPS) | yes | Fiery Waistcord (254085, -0.08 DPS) [crafted]; Ghostweave Cord (254073, -0.16 DPS) [crafted]; Satyrmane Sash (17755, -4.24 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (10.49 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -6.36 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.68 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.44 DPS) [dungeon]; Venomspew Footpads (275606, -0.81 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (7.94 DPS) | yes | Philanthropist's Ring (281635, -3.42 DPS) [quest]; Snake Hoop (6750, -5.04 DPS) [quest]; Reedknot Ring (9622, -5.12 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.24 DPS) | yes | Philanthropist's Ring (281635, -0.72 DPS) [quest]; Snake Hoop (6750, -2.34 DPS) [quest]; Reedknot Ring (9622, -2.42 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -1.24 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.34 DPS) [quest]; Mark of the Chosen (17774, -2.03 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.36 DPS) [world]; Teebu's Blazing Longsword (1728, -35.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 196.3 spell_power points (79.18 DPS) | yes | Brilliant Wand (249385, -15.29 DPS) [crafted]; Stormrager (16997, -16.00 DPS) [quest]; Greater Eternal Wand (249237, -22.22 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Wisdom of the Timbermaw; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25500000000000000-0005000000000000000-2053045103101351)

Set DPS (verified): 722.5. Weights run: 2.5s. Verify run: 3.2s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.180 ± 0.026, crit=0.318 ± 0.008 per rating point (14 rating = 1%, 4.446 per %), hit=0.692 ± 0.036 per rating point (10 rating = 1%, 6.918 per %), spell_haste=not significant (0.287 ± 0.495), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.308 ± 0.001, fire_power=0.692 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (15.59 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.43 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.8 spell_power points (7.97 DPS) | yes | Glowing Eye of Mordresh (10769, -3.10 DPS) [dungeon]; Triune Amulet (7722, -3.35 DPS) [dungeon]; Mindburst Medallion (11196, -3.88 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (14.44 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.70 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -3.59 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.1 spell_power points (8.71 DPS) | yes | Mageflame Cloak (13007, -0.31 DPS) [world_drop]; Deep Woodlands Cloak (19121, -0.84 DPS) [quest]; Cindercloth Cloak (14044, -2.68 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (27.50 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.53 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.98 DPS) [pvp]; Robe of the Magi (1716, -14.11 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.8 spell_power points (7.39 DPS) | yes | Netherflame Cuffs (254065, -1.47 DPS) [crafted]; Phoenix Bindings (210781, -2.19 DPS) [crafted]; Spidertank Oilrag (9448, -2.19 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (12.70 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.16 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.8 spell_power points (9.12 DPS) | yes | Fiery Waistcord (254085, -0.59 DPS) [crafted]; Ghostweave Cord (254073, -1.04 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.32 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (15.01 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -9.71 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.86 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.15 DPS) [crafted]; Omnicast Boots (11822, -4.29 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.3 spell_power points (11.12 DPS) | yes | Philanthropist's Ring (281635, -4.83 DPS) [quest]; Reedknot Ring (9622, -7.08 DPS) [quest]; Snake Hoop (6750, -7.14 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.51 DPS) | yes | Reedknot Ring (9622, -3.46 DPS) [quest]; Snake Hoop (6750, -3.52 DPS) [quest]; Philanthropist's Ring (281635, -3.71 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (722.5 DPS) | yes | Uther's Strength (11302, -0.13 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (722.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (722.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.09 DPS) [world]; Teebu's Blazing Longsword (1728, -51.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 138.9 spell_power points (80.22 DPS) | yes | Brilliant Wand (249385, -5.78 DPS, sim-verified) [crafted]; Stormrager (16997, -17.02 DPS) [quest]; Greater Eternal Wand (249237, -21.70 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

