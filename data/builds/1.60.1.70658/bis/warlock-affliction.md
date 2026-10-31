# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 42.1. Weights run: 2.1s. Verify run: 1.3s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.019 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.426 per %), hit=0.102 ± 0.003 per rating point (10 rating = 1%, 1.019 per %), spell_haste=not significant (-0.104 ± 0.048), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.746 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.14 DPS) | yes | Red Winter Hat (21524, -2.95 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.19 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.76 DPS) | yes | Pearl-clasped Cloak (5542, -0.38 DPS) [crafted]; Black Whelp Cloak (7283, -0.42 DPS, sim-verified) [crafted]; Reinforced Linen Cape (2580, -0.57 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -1.40 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.19 DPS) | yes | Ivycloth Bracelets (9793, -0.25 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.33 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.57 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.95 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.02 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.19 DPS) [crafted]; Heavy Woolen Pants (4316, -0.57 DPS) [crafted]; Silk-threaded Trousers (1929, -1.42 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.33 DPS) | yes | Red Woolen Boots (4313, -0.70 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.76 DPS) [crafted]; Soft-soled Linen Boots (4312, -0.95 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.95 DPS) | yes | Ring of the Shadow (1462, -1.22 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.57 DPS) | yes | Ring of the Shadow (1462, -0.83 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 89 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.10 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.95 DPS) | yes | Bouquet of Red Roses (22206, -1.21 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 120.8 spell_power points (22.88 DPS) | yes | Skycaller (12984, -2.31 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.28 DPS) [quest]; Deepblaze (279896, -4.35 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 75.1. Weights run: 2.1s. Verify run: 1.1s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.153 ± 0.005, crit=0.036 ± 0.001 per rating point (14 rating = 1%, 0.501 per %), hit=0.133 ± 0.005 per rating point (10 rating = 1%, 1.329 per %), spell_haste=not significant (0.197 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.778 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.46 DPS) | yes | Filigreed Pristine Circlet (253975, -0.67 DPS) [crafted]; Enchanter's Cowl (4322, -0.78 DPS) [crafted]; Silk Headband (7050, -3.27 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.9 spell_power points (0.21 DPS) | yes | Crystal Starfire Medallion (5003, -0.07 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.07 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (2.32 DPS) | yes | Death Speaker Mantle (6685, -0.60 DPS) [dungeon]; Invoker's Mantle (215365, -0.67 DPS, sim-verified) [crafted]; Fairywing Mantle (9536, -0.67 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.12 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.91 DPS) | yes | Beguiler Robes (7728, -0.48 DPS) [dungeon]; Green Silk Armor (7065, -0.60 DPS, sim-verified) [crafted]; Mechbuilder's Overalls (9508, -0.71 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.01 DPS) | yes | Windsong Bangles (263336, -1.79 DPS) [quest]; Nightsky Wristbands (6407, -1.81 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.46 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.57 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.5 spell_power points (2.12 DPS) | yes | Invoker's Cord (215366, -0.38 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.44 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.53 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.69 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.12 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.14 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (1.81 DPS) | yes | Glinteye Slippers (273024, -0.24 DPS) [dungeon]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.71 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (1.51 DPS) | yes | Minor Channeling Ring (1449, -0.33 DPS) [quest]; Electrocutioner Lagnut (9447, -0.84 DPS) [dungeon]; Runemender's Seal (273031, -0.84 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.34 DPS) | yes | Minor Channeling Ring (1449, -0.41 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Runemender's Seal (273031, -0.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.06 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.58 DPS) [quest]; Glimmering Staff (249392, -7.68 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 147.5 spell_power points (33.00 DPS) | yes | Wand of Mana Concentration (273043, -2.32 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -2.88 DPS) [crafted]; Gravestone Scepter (7001, -4.00 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 166.3. Weights run: 2.1s. Verify run: 1.5s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.708 ± 0.015, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.371 per %), hit=0.192 ± 0.010 per rating point (10 rating = 1%, 1.922 per %), spell_haste=0.542 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.11 DPS) | yes | Augural Shroud (2620, -0.99 DPS) [world]; Papal Fez (9431, -2.09 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -2.34 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (4.29 DPS) | yes | Triune Amulet (7722, -1.10 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.61 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.85 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (5.48 DPS) | yes | Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.72 DPS) [quest]; Green Silken Shoulders (7057, -1.62 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (3.23 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.86 DPS) [crafted]; Blackforge Cape (6424, -1.31 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (8.88 DPS) | yes | Dreamweave Vest (10021, -0.81 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.27 DPS) [crafted]; Elemental Raiment (9434, -1.78 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.05 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.68 DPS) [quest]; Windchaser Cuffs (14429, -0.89 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.05 DPS) | yes | Black Mageweave Gloves (10003, -1.97 DPS) [crafted]; Gilded Handwraps (254021, -2.67 DPS) [crafted]; Red Mageweave Gloves (10018, -2.70 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (166.3 DPS) | yes | Star Belt (4329, -0.22 DPS) [crafted]; Belt of Arugal (6392, -0.86 DPS) [dungeon]; Deathmage Sash (10771, -2.74 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (7.61 DPS) | yes | Crimson Silk Pantaloons (7062, -2.01 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -3.33 DPS) [dungeon]; Gaze Dreamer Pants (6903, -3.55 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.12 DPS) | yes | Gilded Slippers (254001, -3.43 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.51 DPS) [dungeon]; Spidersilk Boots (4320, -4.79 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.58 DPS) | yes | Ring of Forlorn Spirits (2043, -1.87 DPS) [quest]; Reedknot Ring (9622, -2.21 DPS) [quest]; Minor Channeling Ring (1449, -2.41 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (3.23 DPS) | yes | Ring of Forlorn Spirits (2043, -0.52 DPS) [quest]; Reedknot Ring (9622, -0.86 DPS) [quest]; Minor Channeling Ring (1449, -1.06 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.48 DPS) [dungeon]; Spine of the Devourer (273040, -4.62 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.74 DPS) | yes | Orb of Mystic Insight (249394, -0.93 DPS) [crafted]; Orb of Lorica (11262, -0.96 DPS, sim-verified) [quest]; Beacon of Hope (9393, -2.03 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 120.6 spell_power points (40.82 DPS) | yes | Earthen Rod (9381, -1.75 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.79 DPS) [crafted]; Umbral Wand (5216, -5.15 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 237.0. Weights run: 5.6s. Verify run: 1.7s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.043 ± 0.008, crit=0.033 ± 0.001 per rating point (14 rating = 1%, 0.466 per %), hit=0.302 ± 0.013 per rating point (10 rating = 1%, 3.023 per %), spell_haste=not significant (-0.082 ± 0.105), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.93 DPS) | yes | Dreamweave Circlet (10041, -1.43 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.99 DPS) [crafted]; Red Mageweave Headband (10033, -2.36 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (2.65 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.58 DPS) [quest]; Horizon Choker (13085, -2.45 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.8 spell_power points (4.56 DPS) | yes | Black Mageweave Shoulders (10027, -0.71 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -1.45 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.47 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (4.72 DPS) | yes | Nightfall Drape (12465, -1.74 DPS) [dungeon]; Runecloth Cloak (13860, -1.81 DPS, sim-verified) [crafted]; Icy Cloak (4327, -2.40 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (7.36 DPS) | yes | Elemental Raiment (9434, -0.66 DPS, sim-verified) [world_drop]; Acumen Robes (17775, -0.79 DPS) [quest]; Dreamweave Vest (10021, -1.28 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.98 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.56 DPS) [crafted]; Condor Bracers (15864, -0.66 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (6.01 DPS) | yes | Black Mageweave Gloves (10003, -1.58 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.58 DPS) [vendor]; Brightcloth Gloves (14101, -1.71 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.4 spell_power points (4.77 DPS) | yes | Ghostweave Cord (254073, +0.00 DPS, sim-verified) [crafted]; Star Belt (4329, -0.47 DPS) [crafted]; Gilded Waistcord (254081, -1.34 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.29 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.99 DPS) [vendor]; Red Mageweave Pants (10009, -2.02 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -2.22 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.94 DPS) | yes | Gilded Sandals (254107, -4.17 DPS) [crafted]; Black Mageweave Boots (10026, -4.20 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.56 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.30 DPS) | yes | Ring of Forlorn Spirits (2043, -1.65 DPS) [quest]; Reedknot Ring (9622, -1.99 DPS) [quest]; Snake Hoop (6750, -2.24 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.2 spell_power points (3.38 DPS) | yes | Ring of Forlorn Spirits (2043, -0.73 DPS) [quest]; Reedknot Ring (9622, -1.06 DPS) [quest]; Snake Hoop (6750, -1.32 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -4.23 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.49 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.66 DPS) [dungeon]; Hanzo Sword (8190, -29.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (237.0 DPS) | yes | Wand of Allistarj (13065, -3.23 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.26 DPS) [crafted]; Pyric Caduceus (11748, -9.18 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Triune Amulet; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 389.9. Weights run: 6.2s. Verify run: 2.7s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.076 ± 0.014, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.856 per %), hit=0.337 ± 0.023 per rating point (10 rating = 1%, 3.375 per %), spell_haste=not significant (0.888 ± 0.258), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.903 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.31 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.72 DPS) [pvp]; Dreamweave Circlet (10041, -2.39 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.8 spell_power points (5.82 DPS) | yes | Triune Amulet (7722, -2.05 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -2.35 DPS) [dungeon]; Mindburst Medallion (11196, -2.87 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (11.40 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.77 DPS, sim-verified) [crafted]; Ghostweave Mantle (254127, -4.56 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.5 spell_power points (6.59 DPS) | yes | Nightfall Drape (12465, -2.49 DPS) [dungeon]; Runecloth Cloak (13860, -2.67 DPS, sim-verified) [crafted]; Icy Cloak (4327, -3.40 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (21.29 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.87 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.20 DPS) [pvp]; Robe of the Magi (1716, -11.17 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.8 spell_power points (5.36 DPS) | yes | Arcane Runed Bracers (4744, -1.26 DPS) [quest]; Spidertank Oilrag (9448, -1.26 DPS) [dungeon]; Earthenweave Cuffs (254131, -1.71 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (10.03 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.32 DPS) [pvp]; Dreamweave Gloves (10019, -1.69 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.8 spell_power points (6.73 DPS) | yes | Ghostweave Cord (254073, -0.35 DPS) [crafted]; Star Belt (4329, -0.80 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.23 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.86 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -2.71 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.95 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.91 DPS) [crafted]; Omnicast Boots (11822, -1.41 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.5 spell_power points (8.45 DPS) | yes | Philanthropist's Ring (281635, -3.72 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.80 DPS) [quest]; Reedknot Ring (9622, -5.26 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.93 DPS) | yes | Philanthropist's Ring (281635, -1.20 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.28 DPS) [quest]; Reedknot Ring (9622, -2.74 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -5.25 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.50 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -37.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (389.9 DPS) | yes | Stormrager (16997, -1.06 DPS) [quest]; Greater Eternal Wand (249237, -6.49 DPS) [crafted]; Torch of Light (279246, -27.00 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Brilliant Wand

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25520300130201351-0005200000000000000-0550001000000000)

Set DPS (verified): 737.8. Weights run: 7.2s. Verify run: 2.9s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.078 ± 0.023, crit=0.192 ± 0.008 per rating point (14 rating = 1%, 2.690 per %), hit=0.632 ± 0.047 per rating point (10 rating = 1%, 6.324 per %), spell_haste=not significant (1.688 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.899 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.34 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.15 DPS) [pvp]; Dreamweave Circlet (10041, -3.16 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.8 spell_power points (7.74 DPS) | yes | Triune Amulet (7722, -2.89 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -3.12 DPS) [dungeon]; Mindburst Medallion (11196, -3.82 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (15.13 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -5.62 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.5 spell_power points (8.76 DPS) | yes | Runecloth Cloak (13860, -2.93 DPS) [crafted]; Nightfall Drape (12465, -3.31 DPS) [dungeon]; Icy Cloak (4327, -4.52 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (28.26 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.76 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.19 DPS) [pvp]; Robe of the Magi (1716, -12.67 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.8 spell_power points (7.13 DPS) | yes | Arcane Runed Bracers (4744, -1.68 DPS) [quest]; Spidertank Oilrag (9448, -1.68 DPS) [dungeon]; Earthenweave Cuffs (254131, -2.29 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.31 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.42 DPS) [pvp]; Dreamweave Gloves (10019, -2.23 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.8 spell_power points (8.95 DPS) | yes | Ghostweave Cord (254073, -0.47 DPS) [crafted]; Star Belt (4329, -1.08 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.63 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (15.73 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.24 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.52 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.21 DPS) [crafted]; Omnicast Boots (11822, -1.85 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.5 spell_power points (11.22 DPS) | yes | Philanthropist's Ring (281635, -4.94 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.38 DPS) [quest]; Reedknot Ring (9622, -6.99 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.87 DPS) | yes | Philanthropist's Ring (281635, -1.58 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.03 DPS) [quest]; Reedknot Ring (9622, -3.63 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.64 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -45.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (737.8 DPS) | yes | Stormrager (16997, -2.43 DPS) [quest]; Greater Eternal Wand (249237, -6.58 DPS) [crafted]; Torch of Light (279246, -47.03 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; ranged: Brilliant Wand

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 40.0. Weights run: 2.1s. Verify run: 1.3s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.019 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.426 per %), hit=0.102 ± 0.003 per rating point (10 rating = 1%, 1.019 per %), spell_haste=not significant (-0.104 ± 0.048), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.746 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.14 DPS) | yes | Red Winter Hat (21524, -2.84 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (0.95 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.19 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.26 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.76 DPS) | yes | Mystic Shawl (3449, -0.38 DPS) [quest]; Pearl-clasped Cloak (5542, -0.38 DPS) [crafted]; Black Whelp Cloak (7283, -0.46 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -1.32 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.57 DPS) | yes | Owlbeard Bracers (16981, -0.38 DPS) [quest]; Windsong Bangles (263336, -0.47 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.33 DPS) | yes | Gnoll Casting Gloves (892, -0.30 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.57 DPS) [quest]; Pristine Gloves (253913, -0.57 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.24 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.19 DPS) [crafted]; Embroidered Leggings (270011, -0.38 DPS) [quest]; Silk-threaded Trousers (1929, -1.93 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.33 DPS) | yes | Pristine Boots (253889, -0.76 DPS) [crafted]; Soft-soled Linen Boots (4312, -0.95 DPS) [crafted]; Red Woolen Boots (4313, -1.00 DPS, sim-verified) [crafted] |
| finger1 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.57 DPS) | yes | - |
| finger2 | - | - |  |  |  |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Evocator's Blade (2567, -0.57 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 120.8 spell_power points (22.88 DPS) | yes | Skycaller (12984, -2.01 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.28 DPS) [quest]; Greater Magic Wand (11288, -5.00 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 74.7. Weights run: 2.1s. Verify run: 1.2s. 283 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.153 ± 0.005, crit=0.036 ± 0.001 per rating point (14 rating = 1%, 0.501 per %), hit=0.133 ± 0.005 per rating point (10 rating = 1%, 1.329 per %), spell_haste=not significant (0.197 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.778 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.46 DPS) | yes | Filigreed Pristine Circlet (253975, -0.67 DPS) [crafted]; Enchanter's Cowl (4322, -0.78 DPS) [crafted]; Silk Headband (7050, -3.54 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.9 spell_power points (0.21 DPS) | yes | Crystal Starfire Medallion (5003, -0.07 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.07 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (2.32 DPS) | yes | Invoker's Mantle (215365, -0.58 DPS) [crafted]; Death Speaker Mantle (6685, -0.60 DPS) [dungeon]; Chestnut Mantle (17695, -1.16 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.12 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Black Whelp Cloak (7283, -0.45 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.91 DPS) | yes | Beguiler Robes (7728, -0.48 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.71 DPS) [dungeon]; Green Silk Armor (7065, -1.16 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.01 DPS) | yes | Tabitha's Cuffs (251486, -1.50 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -1.72 DPS) [quest]; Glowing Magical Bracelets (13106, -1.74 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.57 DPS) | yes | Jutebraid Gloves (10654, -0.05 DPS) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.46 DPS) | yes | Belt of Arugal (6392, -0.30 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.67 DPS) [dungeon]; Invoker's Cord (215366, -0.72 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.69 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.12 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.14 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (1.81 DPS) | yes | Glinteye Slippers (273024, -0.24 DPS) [dungeon]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.94 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (1.51 DPS) | yes | Solid Gold Band (279845, -0.62 DPS) [quest]; Electrocutioner Lagnut (9447, -0.84 DPS) [dungeon]; Runemender's Seal (273031, -0.84 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.34 DPS) | yes | Solid Gold Band (279845, -0.45 DPS) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Runemender's Seal (273031, -0.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.06 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -7.68 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -7.71 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 147.5 spell_power points (33.00 DPS) | yes | Greater Mystic Wand (217287, -2.88 DPS) [crafted]; Wand of Mana Concentration (273043, -3.20 DPS, sim-verified) [world]; Unstable Power Core (279847, -3.29 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 165.6. Weights run: 2.1s. Verify run: 1.5s. 381 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.708 ± 0.015, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.371 per %), hit=0.192 ± 0.010 per rating point (10 rating = 1%, 1.922 per %), spell_haste=0.542 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.11 DPS) | yes | Augural Shroud (2620, -0.99 DPS) [world]; Papal Fez (9431, -1.77 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -2.34 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (4.29 DPS) | yes | Triune Amulet (7722, -0.61 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.61 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.85 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (5.48 DPS) | yes | Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.72 DPS) [quest]; Green Silken Shoulders (7057, -1.22 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (3.23 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.86 DPS) [crafted]; Blackforge Cape (6424, -1.31 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (8.88 DPS) | yes | Dreamweave Vest (10021, -1.25 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.27 DPS) [crafted]; Elemental Raiment (9434, -1.78 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Condor Bracers (15864, -0.68 DPS) [quest]; Windchaser Cuffs (14429, -0.89 DPS) [world_drop]; Radiant Silver Bracers (4545, -1.76 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.05 DPS) | yes | Black Mageweave Gloves (10003, -1.97 DPS) [crafted]; Gilded Handwraps (254021, -2.67 DPS) [crafted]; Red Mageweave Gloves (10018, -2.84 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.22 DPS) [crafted]; Belt of Arugal (6392, -0.86 DPS) [dungeon]; Deathmage Sash (10771, -3.33 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (7.61 DPS) | yes | Crimson Silk Pantaloons (7062, -2.29 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -3.33 DPS) [dungeon]; Gaze Dreamer Pants (6903, -3.55 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.12 DPS) | yes | Gilded Slippers (254001, -4.14 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.51 DPS) [dungeon]; Spidersilk Boots (4320, -4.79 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.58 DPS) | yes | Reedknot Ring (9622, -2.21 DPS) [quest]; Ring of Power Regulation (273030, -2.55 DPS) [dungeon]; Voodoo Band (1996, -2.90 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (3.23 DPS) | yes | Reedknot Ring (9622, -0.86 DPS) [quest]; Ring of Power Regulation (273030, -1.20 DPS) [dungeon]; Voodoo Band (1996, -1.55 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.48 DPS) [dungeon]; Spine of the Devourer (273040, -4.62 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.74 DPS) | yes | Omega Orb (7749, -1.69 DPS) [quest]; Beacon of Hope (9393, -2.03 DPS) [dungeon]; Orb of Mystic Insight (249394, -2.64 DPS, sim-verified) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 120.6 spell_power points (40.82 DPS) | yes | Earthen Rod (9381, -1.47 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.79 DPS) [crafted]; Umbral Wand (5216, -5.15 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 381, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 235.4. Weights run: 5.6s. Verify run: 1.6s. 492 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.043 ± 0.008, crit=0.033 ± 0.001 per rating point (14 rating = 1%, 0.466 per %), hit=0.302 ± 0.013 per rating point (10 rating = 1%, 3.023 per %), spell_haste=not significant (-0.082 ± 0.105), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.93 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.99 DPS) [crafted]; Dreamweave Circlet (10041, -2.11 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.36 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (2.65 DPS) | yes | Glowing Eye of Mordresh (10769, -0.22 DPS) [dungeon]; Mindburst Medallion (11196, -0.58 DPS) [quest]; Horizon Choker (13085, -2.45 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.8 spell_power points (4.56 DPS) | yes | Bloodmage Mantle (7684, -1.45 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.47 DPS) [crafted]; Black Mageweave Shoulders (10027, -2.91 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (4.72 DPS) | yes | Deep Woodlands Cloak (19121, -1.35 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.63 DPS) [crafted]; Nightfall Drape (12465, -1.74 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (7.36 DPS) | yes | Acumen Robes (17775, -0.79 DPS) [quest]; Dreamweave Vest (10021, -1.28 DPS) [crafted]; Elemental Raiment (9434, -2.36 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.98 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.66 DPS) [quest]; Bloodband Bracers (11469, -1.20 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (6.01 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.58 DPS) [vendor]; Brightcloth Gloves (14101, -1.71 DPS) [crafted]; Black Mageweave Gloves (10003, -2.99 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.4 spell_power points (4.77 DPS) | yes | Star Belt (4329, -0.47 DPS) [crafted]; Ghostweave Cord (254073, -1.02 DPS, sim-verified) [crafted]; Warsong Sash (16975, -1.13 DPS) [quest] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.29 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.99 DPS) [vendor]; Senior Designer's Pantaloons (11841, -2.22 DPS) [dungeon]; Red Mageweave Pants (10009, -2.28 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.94 DPS) | yes | Gilded Sandals (254107, -4.17 DPS) [crafted]; Black Mageweave Boots (10026, -4.20 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -5.29 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.30 DPS) | yes | Reedknot Ring (9622, -1.99 DPS) [quest]; Snake Hoop (6750, -2.24 DPS) [quest]; Ring of Power Regulation (273030, -2.32 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.2 spell_power points (3.38 DPS) | yes | Snake Hoop (6750, -1.32 DPS) [quest]; Ring of Power Regulation (273030, -1.39 DPS) [dungeon]; Reedknot Ring (9622, -2.77 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.20 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -3.58 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.49 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.66 DPS) [dungeon]; Hanzo Sword (8190, -30.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (235.4 DPS) | yes | Wand of Allistarj (13065, -3.23 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.26 DPS) [crafted]; Pyric Caduceus (11748, -8.32 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Triune Amulet; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 492, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 381.5. Weights run: 6.2s. Verify run: 2.9s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.076 ± 0.014, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.856 per %), hit=0.337 ± 0.023 per rating point (10 rating = 1%, 3.375 per %), spell_haste=not significant (0.888 ± 0.258), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.903 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.31 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.72 DPS) [pvp]; Dreamweave Circlet (10041, -2.39 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.8 spell_power points (5.82 DPS) | yes | Glowing Eye of Mordresh (10769, -2.35 DPS) [dungeon]; Mindburst Medallion (11196, -2.87 DPS) [quest]; Triune Amulet (7722, -4.44 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (11.40 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.81 DPS) [crafted]; Ghostweave Mantle (254127, -4.56 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.5 spell_power points (6.59 DPS) | yes | Deep Woodlands Cloak (19121, -1.61 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.21 DPS) [crafted]; Nightfall Drape (12465, -2.49 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (21.29 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.87 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.20 DPS) [pvp]; Robe of the Magi (1716, -11.47 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.8 spell_power points (5.36 DPS) | yes | Earthenweave Cuffs (254131, -1.71 DPS) [crafted]; Nethergeld Cuffs (254061, -1.93 DPS) [crafted]; Spidertank Oilrag (9448, -6.15 DPS, sim-verified) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (10.03 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.32 DPS) [pvp]; Dreamweave Gloves (10019, -3.47 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.8 spell_power points (6.73 DPS) | yes | Star Belt (4329, -0.80 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.23 DPS) [crafted]; Ghostweave Cord (254073, -2.97 DPS, sim-verified) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.86 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -5.89 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.95 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.41 DPS) [dungeon]; Venomspew Footpads (275606, -1.79 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.5 spell_power points (8.45 DPS) | yes | Philanthropist's Ring (281635, -3.72 DPS) [quest]; Reedknot Ring (9622, -5.26 DPS) [quest]; Snake Hoop (6750, -5.54 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.93 DPS) | yes | Philanthropist's Ring (281635, -1.69 DPS, sim-verified) [quest]; Reedknot Ring (9622, -2.74 DPS) [quest]; Snake Hoop (6750, -3.02 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.50 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -36.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (381.5 DPS) | yes | Stormrager (16997, -1.06 DPS) [quest]; Greater Eternal Wand (249237, -6.49 DPS) [crafted]; Torch of Light (279246, -26.56 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Brilliant Wand

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25520300130201351-0005200000000000000-0550001000000000)

Set DPS (verified): 737.3. Weights run: 7.2s. Verify run: 2.9s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.078 ± 0.023, crit=0.192 ± 0.008 per rating point (14 rating = 1%, 2.690 per %), hit=0.632 ± 0.047 per rating point (10 rating = 1%, 6.324 per %), spell_haste=not significant (1.688 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.899 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.34 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.15 DPS) [pvp]; Dreamweave Circlet (10041, -6.58 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.8 spell_power points (7.74 DPS) | yes | Glowing Eye of Mordresh (10769, -3.12 DPS) [dungeon]; Mindburst Medallion (11196, -3.82 DPS) [quest]; Triune Amulet (7722, -7.34 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (15.13 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -5.62 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.72 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.5 spell_power points (8.76 DPS) | yes | Deep Woodlands Cloak (19121, -2.67 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.93 DPS) [crafted]; Nightfall Drape (12465, -3.31 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (28.26 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.76 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.19 DPS) [pvp]; Robe of the Magi (1716, -16.10 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.8 spell_power points (7.13 DPS) | yes | Earthenweave Cuffs (254131, -2.29 DPS) [crafted]; Nethergeld Cuffs (254061, -2.56 DPS) [crafted]; Spidertank Oilrag (9448, -5.48 DPS, sim-verified) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.31 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.42 DPS) [pvp]; Dreamweave Gloves (10019, -4.10 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.8 spell_power points (8.95 DPS) | yes | Star Belt (4329, -1.08 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.63 DPS) [crafted]; Ghostweave Cord (254073, -5.98 DPS, sim-verified) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (15.73 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -7.69 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.52 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.21 DPS) [crafted]; Omnicast Boots (11822, -1.85 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.5 spell_power points (11.22 DPS) | yes | Philanthropist's Ring (281635, -4.94 DPS) [quest]; Reedknot Ring (9622, -6.99 DPS) [quest]; Snake Hoop (6750, -7.36 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.87 DPS) | yes | Philanthropist's Ring (281635, -3.22 DPS, sim-verified) [quest]; Reedknot Ring (9622, -3.63 DPS) [quest]; Snake Hoop (6750, -4.00 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -3.41 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.64 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -48.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (737.3 DPS) | yes | Stormrager (16997, -2.43 DPS) [quest]; Greater Eternal Wand (249237, -6.58 DPS) [crafted]; Torch of Light (279246, -44.15 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Brilliant Wand

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

