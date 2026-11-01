# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 49.6. Weights run: 1.8s. Verify run: 1.2s. 161 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.044 ± 0.004, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.389 per %), hit=0.081 ± 0.003 per rating point (10 rating = 1%, 0.808 per %), spell_haste=0.258 ± 0.048, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.735 ± 0.002, fire_power=0.265 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.39 DPS) | yes | Flame Circlet (253953, -2.34 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.16 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.37 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.46 DPS) [crafted]; Black Whelp Cloak (7283, -0.50 DPS, sim-verified) [crafted]; Reinforced Linen Cape (2580, -0.69 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.16 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -1.88 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.23 DPS) | yes | Ivycloth Bracelets (9793, -0.27 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.62 DPS) | yes | Gnoll Casting Gloves (892, -0.38 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.69 DPS) [crafted]; Phoenix Gloves (4331, -1.07 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.46 DPS) [crafted]; Flame Sash (253929, -0.50 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.05 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.23 DPS) [crafted]; Heavy Woolen Pants (4316, -0.69 DPS) [crafted]; Silk-threaded Trousers (1929, -1.45 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.62 DPS) | yes | Red Woolen Boots (4313, -0.66 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.93 DPS) [crafted]; Soft-soled Linen Boots (4312, -1.16 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.16 DPS) | yes | Ring of the Shadow (1462, -1.59 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.69 DPS) | yes | Ring of the Shadow (1462, -1.26 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 89 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.15 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (1.16 DPS) | yes | Antipodean Rod (2879, -1.23 DPS, sim-verified) [world_drop] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 105.4 spell_power points (24.39 DPS) | yes | Cookie's Stirring Rod (5198, -0.53 DPS, sim-verified) [dungeon]; Sable Wand (7607, -1.48 DPS) [quest]; Skycaller (12984, -2.83 DPS) [world_drop] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 161, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 92.1. Weights run: 1.8s. Verify run: 1.0s. 298 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.162 ± 0.006, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.121 per %), hit=0.128 ± 0.006 per rating point (10 rating = 1%, 1.278 per %), spell_haste=not significant (-0.035 ± 0.091), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.734 ± 0.002, fire_power=0.266 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.81 DPS) | yes | Filigreed Pristine Circlet (253975, -0.77 DPS) [crafted]; Enchanter's Cowl (4322, -0.87 DPS) [crafted]; Silk Headband (7050, -4.16 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 1.0 spell_power points (0.25 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (2.67 DPS) | yes | Death Speaker Mantle (6685, -0.68 DPS) [dungeon]; Fairywing Mantle (9536, -0.77 DPS) [quest]; Invoker's Mantle (215365, -0.78 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.28 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.26 DPS) [crafted]; Prelacy Cape (7004, -0.26 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.33 DPS) | yes | Beguiler Robes (7728, -0.53 DPS) [dungeon]; Green Silk Armor (7065, -0.59 DPS, sim-verified) [crafted]; Mechbuilder's Overalls (9508, -0.78 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.30 DPS) | yes | Phoenix Bindings (210781, -1.78 DPS, sim-verified) [crafted]; Glowing Magical Bracelets (13106, -1.97 DPS) [world_drop]; Windsong Bangles (263336, -2.05 DPS) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.79 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.26 DPS) [world]; Truefaith Gloves (7049, -0.39 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.5 spell_power points (2.43 DPS) | yes | Invoker's Cord (215366, -0.43 DPS) [crafted]; Crimson Silk Belt (7055, -0.60 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.88 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.07 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.28 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.29 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (2.08 DPS) | yes | Glinteye Slippers (273024, -0.29 DPS) [dungeon]; Acidic Walkers (9454, -0.47 DPS) [dungeon]; Spidersilk Boots (4320, -2.50 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (1.74 DPS) | yes | Minor Channeling Ring (1449, -0.38 DPS) [quest]; Electrocutioner Lagnut (9447, -0.97 DPS) [dungeon]; Runemender's Seal (273031, -0.97 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.53 DPS) | yes | Minor Channeling Ring (1449, -0.36 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Runemender's Seal (273031, -0.77 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (9.21 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -4.09 DPS) [quest]; Staff of Soran'ruk (15109, -6.89 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 129.1 spell_power points (33.02 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Gyromatic Icemaker (9489, -2.67 DPS) [dungeon]; Greater Mystic Wand (217287, -2.74 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 129.3. Weights run: 1.9s. Verify run: 1.2s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.568 ± 0.016, crit=0.117 ± 0.004 per rating point (14 rating = 1%, 1.633 per %), hit=0.234 ± 0.015 per rating point (10 rating = 1%, 2.341 per %), spell_haste=not significant (0.135 ± 0.190), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.315 ± 0.001, fire_power=0.685 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.70 DPS) | yes | Papal Fez (9431, -0.66 DPS, sim-verified) [dungeon]; Augural Shroud (2620, -1.17 DPS) [world]; Living Cowl (5608, -2.17 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.5 spell_power points (3.13 DPS) | yes | Triune Amulet (7722, -0.74 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.05 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.21 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.4 spell_power points (3.90 DPS) | yes | Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest]; Green Silken Shoulders (7057, -1.70 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.8 spell_power points (2.40 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.50 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.04 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.4 spell_power points (6.90 DPS) | yes | Dreamweave Vest (10021, -0.62 DPS) [crafted]; Elemental Raiment (9434, -1.20 DPS) [world_drop]; Robe of Power (7054, -1.25 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.44 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Phoenix Bindings (210781, -0.02 DPS) [crafted]; Condor Bracers (15864, -0.54 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (5.50 DPS) | yes | Black Mageweave Gloves (10003, -1.43 DPS) [crafted]; Fiery Handwraps (254025, -1.82 DPS) [crafted]; Red Mageweave Gloves (10018, -2.43 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.5 spell_power points (4.21 DPS) | yes | Fiery Cord (254041, +0.00 DPS, sim-verified) [crafted]; Star Belt (4329, -0.68 DPS) [crafted]; Gilded Cord (254037, -0.81 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.8 spell_power points (5.65 DPS) | yes | Flame Leggings (253991, -1.97 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.07 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.39 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.51 DPS) | yes | Gilded Slippers (254001, -3.54 DPS) [crafted]; Acidic Walkers (9454, -3.92 DPS) [dungeon]; Fiery Slippers (254005, -4.55 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (3.48 DPS) | yes | Ring of Forlorn Spirits (2043, -1.31 DPS) [quest]; Reedknot Ring (9622, -1.58 DPS) [quest]; Minor Channeling Ring (1449, -1.82 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (2.40 DPS) | yes | Ring of Forlorn Spirits (2043, -0.23 DPS) [quest]; Reedknot Ring (9622, -0.50 DPS) [quest]; Minor Channeling Ring (1449, -0.73 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (129.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.31 DPS) [dungeon]; Spine of the Devourer (273040, -3.40 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.80 DPS) | yes | Orb of Lorica (11262, -0.63 DPS) [quest]; Orb of Noh'Orahil (15107, -0.72 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -0.98 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 148.5 spell_power points (40.30 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.63 DPS) [dungeon]; Twisted Nether Wand (249144, -4.67 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 235.5. Weights run: 1.9s. Verify run: 1.4s. 507 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.300 ± 0.016, crit=0.172 ± 0.005 per rating point (14 rating = 1%, 2.413 per %), hit=0.352 ± 0.018 per rating point (10 rating = 1%, 3.522 per %), spell_haste=-1.676 ± 0.233, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.267 ± 0.001, fire_power=0.733 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.63 DPS) | yes | Red Mageweave Headband (10033, -0.93 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.14 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.4 spell_power points (3.35 DPS) | yes | Mindburst Medallion (11196, -0.57 DPS) [quest]; Triune Amulet (7722, -0.91 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.85 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.4 spell_power points (6.56 DPS) | yes | Netherflame Shoulders (254053, -1.26 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.89 DPS) [vendor]; Black Mageweave Shoulders (10027, -2.03 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.8 spell_power points (5.63 DPS) | yes | Runecloth Cloak (13860, -1.57 DPS) [crafted]; Cloak of Fire (14134, -1.97 DPS) [crafted]; Cindercloth Cloak (14044, -2.11 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.0 spell_power points (8.91 DPS) | yes | Robe of the Magi (1716, -0.43 DPS) [world_drop]; Elemental Raiment (9434, -1.43 DPS) [world_drop]; Dreamweave Vest (10021, -1.53 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 11.6 spell_power points (4.15 DPS) | yes | Nethergeld Cuffs (254061, -0.90 DPS) [crafted]; Arcane Runed Bracers (4744, -0.94 DPS) [quest]; Phoenix Bindings (210781, -1.25 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (6.84 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.25 DPS) [vendor]; Fiery Gloves (254099, -1.32 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.50 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.0 spell_power points (6.06 DPS) | yes | Fiery Waistcord (254085, -0.13 DPS) [crafted]; Ghostweave Cord (254073, -1.07 DPS) [crafted]; Star Belt (4329, -1.43 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.77 DPS) | yes | Red Mageweave Pants (10009, -0.50 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -0.95 DPS, sim-verified) [vendor]; Senior Designer's Pantaloons (11841, -1.75 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.56 DPS) | yes | Cindercloth Boots (10044, -3.07 DPS) [crafted]; Fiery Sandals (254111, -3.17 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.49 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.63 DPS) | yes | Ring of Forlorn Spirits (2043, -1.78 DPS) [quest]; Snake Hoop (6750, -1.96 DPS) [quest]; Reedknot Ring (9622, -2.14 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 spell_power points (4.10 DPS) | yes | Snake Hoop (6750, -1.43 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.47 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.60 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (235.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (235.5 DPS) | yes | Mark of the Chosen (17774, -4.60 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (235.5 DPS) | yes | Spire of Hakkar (10844, -0.57 DPS) [world]; Spellshifter Rod (9527, -0.96 DPS) [quest]; Hanzo Sword (8190, -32.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 147.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.90 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 373.0. Weights run: 1.9s. Verify run: 2.5s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.214 ± 0.020, crit=0.208 ± 0.006 per rating point (14 rating = 1%, 2.910 per %), hit=0.520 ± 0.025 per rating point (10 rating = 1%, 5.205 per %), spell_haste=not significant (-0.531 ± 0.335), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.375 ± 0.001, fire_power=0.625 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.53 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -2.91 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.1 spell_power points (6.04 DPS) | yes | Glowing Eye of Mordresh (10769, -2.32 DPS) [dungeon]; Triune Amulet (7722, -2.62 DPS) [dungeon]; Mindburst Medallion (11196, -2.93 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.68 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.99 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.12 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (6.53 DPS) | yes | Mageflame Cloak (13007, -0.92 DPS) [world_drop]; Runecloth Cloak (13860, -1.95 DPS) [crafted]; Cindercloth Cloak (14044, -2.33 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.9 spell_power points (20.47 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.61 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.97 DPS) [pvp]; Robe of the Magi (1716, -10.54 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.1 spell_power points (5.61 DPS) | yes | Netherflame Cuffs (254065, -1.50 DPS) [crafted]; Arcane Runed Bracers (4744, -1.77 DPS) [quest]; Spidertank Oilrag (9448, -1.77 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (9.40 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.06 DPS) [pvp]; Dreamweave Gloves (10019, -1.34 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (6.89 DPS) | yes | Ghostweave Cord (254073, -0.91 DPS) [crafted]; Wisdom of the Timbermaw (19047, -0.92 DPS) [crafted]; Fiery Waistcord (254085, -1.00 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.11 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.92 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.25 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.61 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.77 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.5 spell_power points (8.33 DPS) | yes | Philanthropist's Ring (281635, -3.60 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.91 DPS) [quest]; Snake Hoop (6750, -5.31 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.55 DPS) | yes | Philanthropist's Ring (281635, -0.82 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.14 DPS) [quest]; Snake Hoop (6750, -2.53 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (373.0 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (373.0 DPS) | yes | Mark of the Chosen (17774, -5.36 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (373.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.24 DPS) [world]; Teebu's Blazing Longsword (1728, -36.88 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 185.7 spell_power points (79.33 DPS) | yes | Brilliant Wand (249385, -15.26 DPS) [crafted]; Stormrager (16997, -16.18 DPS) [quest]; Greater Eternal Wand (249237, -22.15 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25500000020000000-0000000000000000000-2053345103101351)

Set DPS (verified): 700.9. Weights run: 2.1s. Verify run: 2.7s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.167 ± 0.027, crit=0.325 ± 0.009 per rating point (14 rating = 1%, 4.551 per %), hit=0.804 ± 0.038 per rating point (10 rating = 1%, 8.043 per %), spell_haste=not significant (0.203 ± 0.490), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.290 ± 0.001, fire_power=0.710 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (17.03 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.73 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (8.62 DPS) | yes | Glowing Eye of Mordresh (10769, -3.36 DPS) [dungeon]; Triune Amulet (7722, -3.57 DPS) [dungeon]; Mindburst Medallion (11196, -4.20 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (15.76 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.81 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.96 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (9.46 DPS) | yes | Cindercloth Cloak (14044, -2.80 DPS) [crafted]; Mageflame Cloak (13007, -2.92 DPS, sim-verified) [world_drop]; Runecloth Cloak (13860, -2.94 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.5 spell_power points (29.95 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.25 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.09 DPS) [pvp]; Robe of the Magi (1716, -19.54 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (7.99 DPS) | yes | Netherflame Cuffs (254065, -1.43 DPS) [crafted]; Phoenix Bindings (210781, -2.16 DPS) [crafted]; Arcane Runed Bracers (4744, -2.31 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.87 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.21 DPS) [pvp]; Gloves of Spell Mastery (14146, -1.41 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.7 spell_power points (9.88 DPS) | yes | Fiery Waistcord (254085, -0.42 DPS) [crafted]; Ghostweave Cord (254073, -1.05 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.47 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (16.39 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -8.79 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.13 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.26 DPS) [crafted]; Omnicast Boots (11822, -1.26 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (12.09 DPS) | yes | Philanthropist's Ring (281635, -5.25 DPS) [quest]; Ring of Forlorn Spirits (2043, -7.04 DPS) [quest]; Reedknot Ring (9622, -7.67 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (8.20 DPS) | yes | Philanthropist's Ring (281635, -1.37 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.15 DPS) [quest]; Reedknot Ring (9622, -3.78 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (700.9 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (700.9 DPS) | yes | Mark of the Chosen (17774, -7.92 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (700.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.00 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -56.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 127.7 spell_power points (80.55 DPS) | yes | Brilliant Wand (249385, -14.55 DPS) [crafted]; Stormrager (16997, -17.33 DPS) [quest]; Greater Eternal Wand (249237, -21.54 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 46.9. Weights run: 1.8s. Verify run: 1.2s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.044 ± 0.004, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.389 per %), hit=0.081 ± 0.003 per rating point (10 rating = 1%, 0.808 per %), spell_haste=0.258 ± 0.048, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.735 ± 0.002, fire_power=0.265 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.39 DPS) | yes | Flame Circlet (253953, -2.18 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.16 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.93 DPS) | yes | Mystic Shawl (3449, -0.46 DPS) [quest]; Pearl-clasped Cloak (5542, -0.46 DPS) [crafted]; Black Whelp Cloak (7283, -0.70 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.16 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -1.70 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.69 DPS) | yes | Owlbeard Bracers (16981, -0.46 DPS) [quest]; Windsong Bangles (263336, -0.50 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.62 DPS) | yes | Gnoll Casting Gloves (892, -0.36 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.69 DPS) [quest]; Pristine Gloves (253913, -0.69 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.46 DPS) [crafted]; Flame Sash (253929, -0.50 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.17 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.23 DPS) [crafted]; Embroidered Leggings (270011, -0.46 DPS) [quest]; Silk-threaded Trousers (1929, -2.01 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.62 DPS) | yes | Pristine Boots (253889, -0.93 DPS) [crafted]; Red Woolen Boots (4313, -1.15 DPS, sim-verified) [crafted]; Soft-soled Linen Boots (4312, -1.16 DPS) [crafted] |
| finger1 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.69 DPS) | yes | - |
| finger2 | - | - |  |  |  |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 98 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 1.6 spell_power points (0.37 DPS) | yes | Bouquet of Red Roses (22206, -0.45 DPS, sim-verified) [dungeon] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 105.4 spell_power points (24.39 DPS) | yes | Cookie's Stirring Rod (5198, -0.58 DPS, sim-verified) [dungeon]; Skycaller (12984, -2.83 DPS) [world_drop]; Greater Magic Wand (11288, -6.43 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Sizzle Stick

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 90.5. Weights run: 1.8s. Verify run: 1.0s. 289 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.162 ± 0.006, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.121 per %), hit=0.128 ± 0.006 per rating point (10 rating = 1%, 1.278 per %), spell_haste=not significant (-0.035 ± 0.091), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.734 ± 0.002, fire_power=0.266 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.81 DPS) | yes | Filigreed Pristine Circlet (253975, -0.77 DPS) [crafted]; Enchanter's Cowl (4322, -0.87 DPS) [crafted]; Silk Headband (7050, -3.58 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 1.0 spell_power points (0.25 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (2.67 DPS) | yes | Chestnut Mantle (17695, -0.55 DPS, sim-verified) [quest]; Invoker's Mantle (215365, -0.68 DPS) [crafted]; Death Speaker Mantle (6685, -0.68 DPS) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.28 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.26 DPS) [crafted]; Black Whelp Cloak (7283, -0.51 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.33 DPS) | yes | Green Silk Armor (7065, -0.38 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.53 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.78 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.30 DPS) | yes | Tabitha's Cuffs (251486, -1.53 DPS) [quest]; Phoenix Bindings (210781, -1.61 DPS, sim-verified) [crafted]; Owlbeard Bracers (16981, -1.96 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.79 DPS) | yes | Jutebraid Gloves (10654, -0.05 DPS) [quest]; Gnoll Casting Gloves (892, -0.26 DPS) [world]; Truefaith Gloves (7049, -0.39 DPS) [crafted] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.81 DPS) | yes | Belt of Arugal (6392, +0.00 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.77 DPS) [dungeon]; Invoker's Cord (215366, -0.82 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.07 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.28 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.29 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (2.08 DPS) | yes | Glinteye Slippers (273024, -0.29 DPS) [dungeon]; Acidic Walkers (9454, -0.47 DPS) [dungeon]; Spidersilk Boots (4320, -2.47 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (1.74 DPS) | yes | Solid Gold Band (279845, -0.72 DPS) [quest]; Electrocutioner Lagnut (9447, -0.97 DPS) [dungeon]; Runemender's Seal (273031, -0.97 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.53 DPS) | yes | Solid Gold Band (279845, -0.51 DPS) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Runemender's Seal (273031, -0.77 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (9.21 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -6.89 DPS) [quest]; Glimmering Staff (249392, -8.75 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 129.1 spell_power points (33.02 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -2.67 DPS) [dungeon]; Greater Mystic Wand (217287, -2.74 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 289, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 129.2. Weights run: 1.9s. Verify run: 1.3s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.568 ± 0.016, crit=0.117 ± 0.004 per rating point (14 rating = 1%, 1.633 per %), hit=0.234 ± 0.015 per rating point (10 rating = 1%, 2.341 per %), spell_haste=not significant (0.135 ± 0.190), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.315 ± 0.001, fire_power=0.685 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.70 DPS) | yes | Papal Fez (9431, -1.08 DPS, sim-verified) [dungeon]; Augural Shroud (2620, -1.17 DPS) [world]; Living Cowl (5608, -2.17 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.5 spell_power points (3.13 DPS) | yes | Triune Amulet (7722, -0.88 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.05 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.21 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 14.4 spell_power points (3.90 DPS) | yes | Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest]; Green Silken Shoulders (7057, -1.74 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.8 spell_power points (2.40 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.50 DPS) [crafted]; Hillman's Cloak (3719, -1.04 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.4 spell_power points (6.90 DPS) | yes | Dreamweave Vest (10021, -0.68 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.20 DPS) [world_drop]; Robe of Power (7054, -1.25 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.44 DPS) | yes | Phoenix Bindings (210781, -0.02 DPS) [crafted]; Radiant Silver Bracers (4545, -0.12 DPS) [quest]; Condor Bracers (15864, -0.54 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (5.50 DPS) | yes | Black Mageweave Gloves (10003, -1.43 DPS) [crafted]; Fiery Handwraps (254025, -1.82 DPS) [crafted]; Red Mageweave Gloves (10018, -2.38 DPS, sim-verified) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | sim-verified (129.2 DPS) | yes | Star Belt (4329, -0.31 DPS) [crafted]; Gilded Cord (254037, -0.43 DPS) [crafted]; Deathmage Sash (10771, -1.28 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.8 spell_power points (5.65 DPS) | yes | Crimson Silk Pantaloons (7062, -1.89 DPS, sim-verified) [crafted]; Flame Leggings (253991, -1.97 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.39 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.51 DPS) | yes | Gilded Slippers (254001, -3.54 DPS) [crafted]; Fiery Slippers (254005, -3.82 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.92 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (3.48 DPS) | yes | Reedknot Ring (9622, -1.58 DPS) [quest]; Ring of Power Regulation (273030, -1.86 DPS) [dungeon]; Solid Gold Band (279845, -2.40 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (2.40 DPS) | yes | Reedknot Ring (9622, -0.50 DPS) [quest]; Ring of Power Regulation (273030, -0.77 DPS) [dungeon]; Solid Gold Band (279845, -1.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.31 DPS) [dungeon]; Spine of the Devourer (273040, -3.40 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.80 DPS) | yes | Orb of Noh'Orahil (15107, -0.83 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -0.98 DPS) [crafted]; Omega Orb (7749, -1.36 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 148.5 spell_power points (40.30 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.63 DPS) [dungeon]; Twisted Nether Wand (249144, -4.67 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 235.4. Weights run: 1.9s. Verify run: 1.4s. 498 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.300 ± 0.016, crit=0.172 ± 0.005 per rating point (14 rating = 1%, 2.413 per %), hit=0.352 ± 0.018 per rating point (10 rating = 1%, 3.522 per %), spell_haste=-1.676 ± 0.233, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.267 ± 0.001, fire_power=0.733 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.63 DPS) | yes | Red Mageweave Headband (10033, +0.00 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.14 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.4 spell_power points (3.35 DPS) | yes | Triune Amulet (7722, -0.50 DPS) [dungeon]; Mindburst Medallion (11196, -0.57 DPS) [quest]; Horizon Choker (13085, -1.85 DPS) [world_drop] |
| shoulder | Netherflame Shoulders (254053) | Tailoring [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Dreadweave Mantle (220905, -0.62 DPS) [vendor]; Black Mageweave Shoulders (10027, -0.77 DPS) [crafted]; Rotgrip Mantle (17732, -2.34 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.8 spell_power points (5.63 DPS) | yes | Deep Woodlands Cloak (19121, -0.39 DPS) [quest]; Cindercloth Cloak (14044, -1.38 DPS) [crafted]; Runecloth Cloak (13860, -1.57 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.0 spell_power points (8.91 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.43 DPS) [world_drop]; Dreamweave Vest (10021, -1.53 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 11.6 spell_power points (4.15 DPS) | yes | Phoenix Bindings (210781, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.90 DPS) [crafted]; Spidertank Oilrag (9448, -0.94 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (6.84 DPS) | yes | Fiery Gloves (254099, -0.81 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.25 DPS) [vendor]; Black Mageweave Gloves (10003, -1.50 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.0 spell_power points (6.06 DPS) | yes | Fiery Waistcord (254085, +0.00 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -1.07 DPS) [crafted]; Star Belt (4329, -1.43 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.77 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.35 DPS) [vendor]; Red Mageweave Pants (10009, -0.50 DPS) [crafted]; Senior Designer's Pantaloons (11841, -1.75 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.56 DPS) | yes | Fiery Sandals (254111, -2.44 DPS, sim-verified) [crafted]; Cindercloth Boots (10044, -3.07 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.49 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.63 DPS) | yes | Snake Hoop (6750, -1.96 DPS) [quest]; Reedknot Ring (9622, -2.14 DPS) [quest]; Ring of Power Regulation (273030, -2.50 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 spell_power points (4.10 DPS) | yes | Snake Hoop (6750, -1.51 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.60 DPS) [quest]; Ring of Power Regulation (273030, -1.96 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.25 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -2.67 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.57 DPS) [world]; Spellshifter Rod (9527, -0.96 DPS) [quest]; Hanzo Sword (8190, -29.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Pyric Caduceus (11748, -2.62 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.18 DPS) [crafted]; Wand of Allistarj (13065, -3.36 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Netherflame Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Noxious Shooter

No-known-source sample (15 of 498, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 371.4. Weights run: 1.9s. Verify run: 3.7s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.214 ± 0.020, crit=0.208 ± 0.006 per rating point (14 rating = 1%, 2.910 per %), hit=0.520 ± 0.025 per rating point (10 rating = 1%, 5.205 per %), spell_haste=not significant (-0.531 ± 0.335), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.375 ± 0.001, fire_power=0.625 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.53 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -3.25 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.1 spell_power points (6.04 DPS) | yes | Glowing Eye of Mordresh (10769, -1.43 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -2.62 DPS) [dungeon]; Mindburst Medallion (11196, -2.93 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.68 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.66 DPS, sim-verified) [crafted]; Champion's Dreadweave Spaulders (227092, -3.12 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (6.53 DPS) | yes | Deep Woodlands Cloak (19121, -0.58 DPS) [quest]; Mageflame Cloak (13007, -0.92 DPS) [world_drop]; Runecloth Cloak (13860, -1.95 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.9 spell_power points (20.47 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.61 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.97 DPS) [pvp]; Robe of the Magi (1716, -9.96 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (371.4 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Netherflame Cuffs (254065, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -9.03 DPS, sim-verified) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (9.40 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.06 DPS) [pvp]; Dreamweave Gloves (10019, -3.28 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (371.4 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -9.25 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (371.4 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -6.87 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.25 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.85 DPS) [crafted]; Omnicast Boots (11822, -3.09 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.5 spell_power points (8.33 DPS) | yes | Philanthropist's Ring (281635, -3.60 DPS) [quest]; Snake Hoop (6750, -5.31 DPS) [quest]; Reedknot Ring (9622, -5.34 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.55 DPS) | yes | Philanthropist's Ring (281635, -2.50 DPS, sim-verified) [quest]; Snake Hoop (6750, -2.53 DPS) [quest]; Reedknot Ring (9622, -2.56 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (371.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (371.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (371.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.24 DPS) [world]; Teebu's Blazing Longsword (1728, -39.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 185.7 spell_power points (79.33 DPS) | yes | Brilliant Wand (249385, -1.46 DPS, sim-verified) [crafted]; Stormrager (16997, -16.18 DPS) [quest]; Greater Eternal Wand (249237, -22.15 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Earth Warder's Gloves; waist: Devout Belt; legs: Devout Skirt; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25500000020000000-0000000000000000000-2053345103101351)

Set DPS (verified): 687.9. Weights run: 2.1s. Verify run: 2.7s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.167 ± 0.027, crit=0.325 ± 0.009 per rating point (14 rating = 1%, 4.551 per %), hit=0.804 ± 0.038 per rating point (10 rating = 1%, 8.043 per %), spell_haste=not significant (0.203 ± 0.490), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.290 ± 0.001, fire_power=0.710 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (17.03 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.73 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (8.62 DPS) | yes | Glowing Eye of Mordresh (10769, -3.36 DPS) [dungeon]; Triune Amulet (7722, -3.57 DPS) [dungeon]; Mindburst Medallion (11196, -4.20 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (15.76 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS, sim-verified) [crafted]; Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -3.96 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (9.46 DPS) | yes | Mageflame Cloak (13007, -0.05 DPS) [world_drop]; Deep Woodlands Cloak (19121, -0.95 DPS) [quest]; Cindercloth Cloak (14044, -2.80 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.5 spell_power points (29.95 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.25 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.09 DPS) [pvp]; Robe of the Magi (1716, -17.41 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (7.99 DPS) | yes | Netherflame Cuffs (254065, -1.43 DPS) [crafted]; Phoenix Bindings (210781, -2.16 DPS) [crafted]; Spidertank Oilrag (9448, -2.31 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.87 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.21 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.7 spell_power points (9.88 DPS) | yes | Fiery Waistcord (254085, -0.42 DPS) [crafted]; Ghostweave Cord (254073, -1.05 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.47 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (16.39 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -4.13 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.13 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.26 DPS) [crafted]; Omnicast Boots (11822, -1.26 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (12.09 DPS) | yes | Philanthropist's Ring (281635, -5.25 DPS) [quest]; Reedknot Ring (9622, -7.67 DPS) [quest]; Snake Hoop (6750, -7.78 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (8.20 DPS) | yes | Philanthropist's Ring (281635, -1.37 DPS) [quest]; Reedknot Ring (9622, -3.78 DPS) [quest]; Snake Hoop (6750, -3.89 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (687.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (687.9 DPS) | yes | Uther's Strength (11302, -0.78 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (687.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.00 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -55.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 127.7 spell_power points (80.55 DPS) | yes | Brilliant Wand (249385, -14.55 DPS) [crafted]; Stormrager (16997, -17.33 DPS) [quest]; Greater Eternal Wand (249237, -21.54 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

