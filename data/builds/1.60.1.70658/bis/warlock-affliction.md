# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 41.9. Weights run: 2.2s. Verify run: 1.3s. 146 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.019 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.426 per %), hit=0.102 ± 0.003 per rating point (10 rating = 1%, 1.019 per %), spell_haste=not significant (-0.104 ± 0.048), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.746 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.14 DPS) | yes | Red Winter Hat (21524, -2.97 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.19 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.76 DPS) | yes | Caretaker's Cape (20428, -0.19 DPS) [rep]; Pearl-clasped Cloak (5542, -0.38 DPS) [crafted]; Black Whelp Cloak (7283, -0.39 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -1.44 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.19 DPS) | yes | Ivycloth Bracelets (9793, -0.25 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.33 DPS) | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.57 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.95 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (41.9 DPS) | yes | Novice Ardent's Sash (253887, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.87 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Filigreed Pristine Leggings (253937, -0.57 DPS) [crafted]; Silk-threaded Trousers (1929, -0.66 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.76 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.33 DPS) | yes | Pristine Boots (253889, -0.76 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS, sim-verified) [crafted]; Soft-soled Linen Boots (4312, -0.95 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.95 DPS) | yes | - |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.95 DPS) | yes | Ring of the Shadow (1462, -1.31 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.36 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.95 DPS) | yes | Bouquet of Red Roses (22206, -1.30 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 120.8 spell_power points (22.88 DPS) | yes | Skycaller (12984, -2.40 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.28 DPS) [quest]; Deepblaze (279896, -4.35 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 146, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde

### Band 30 (gnome, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 79.1. Weights run: 2.1s. Verify run: 1.2s. 264 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.153 ± 0.005, crit=0.036 ± 0.001 per rating point (14 rating = 1%, 0.501 per %), hit=0.133 ± 0.005 per rating point (10 rating = 1%, 1.329 per %), spell_haste=not significant (0.197 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.778 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.46 DPS) | yes | Filigreed Pristine Circlet (253975, -0.67 DPS) [crafted]; Enchanter's Cowl (4322, -0.78 DPS) [crafted]; Silk Headband (7050, -3.61 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (1.77 DPS) | yes | Darkspear Warding Pendant (272075, -1.60 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.63 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.79 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (2.32 DPS) | yes | Death Speaker Mantle (6685, -0.60 DPS) [dungeon]; Invoker's Mantle (215365, -0.61 DPS, sim-verified) [crafted]; Fairywing Mantle (9536, -0.67 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.12 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.91 DPS) | yes | Beguiler Robes (7728, -0.48 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.71 DPS) [dungeon]; Green Silk Armor (7065, -0.82 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.01 DPS) | yes | Windsong Bangles (263336, -1.79 DPS) [quest]; Nightsky Wristbands (6407, -1.81 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.55 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.57 DPS) | yes | Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted]; Serpent Gloves (5970, -0.39 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (2.56 DPS) | yes | Belt of Arugal (6392, -0.51 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.77 DPS) [dungeon]; Invoker's Cord (215366, -0.83 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.69 DPS) | yes | Pristine Leggings (253987, -0.88 DPS) [crafted]; Silk-threaded Trousers (1929, -1.12 DPS) [dungeon]; Abomination Skin Leggings (23173, -1.34 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (1.81 DPS) | yes | Glinteye Slippers (273024, -0.24 DPS) [dungeon]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.82 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.57 DPS) | yes | Ring of Power Regulation (273030, -0.22 DPS) [dungeon]; Minor Channeling Ring (1449, -0.38 DPS) [quest]; Snake Hoop (6750, -2.02 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (79.1 DPS) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Minor Channeling Ring (1449, -0.16 DPS) [quest]; Snake Hoop (6750, -1.38 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.06 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.58 DPS) [quest]; Glimmering Staff (249392, -7.68 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 147.5 spell_power points (33.00 DPS) | yes | Greater Mystic Wand (217287, -2.88 DPS) [crafted]; Wand of Mana Concentration (273043, -2.99 DPS, sim-verified) [world]; Gravestone Scepter (7001, -4.00 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (gnome, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 165.7. Weights run: 2.1s. Verify run: 1.4s. 343 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.708 ± 0.015, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.371 per %), hit=0.192 ± 0.010 per rating point (10 rating = 1%, 1.922 per %), spell_haste=0.542 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.11 DPS) | yes | Augural Shroud (2620, -0.99 DPS) [world]; Papal Fez (9431, -2.09 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -2.34 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (4.29 DPS) | yes | Triune Amulet (7722, -1.58 DPS) [dungeon]; Scorn's Icy Choker (23169, -1.74 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -2.61 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (5.48 DPS) | yes | Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.72 DPS) [quest]; Green Silken Shoulders (7057, -1.62 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (3.23 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.59 DPS) [vendor]; Icy Cloak (4327, -0.86 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (8.88 DPS) | yes | Dreamweave Vest (10021, -0.81 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.27 DPS) [crafted]; Elemental Raiment (9434, -1.78 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.05 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.68 DPS) [quest]; Windchaser Cuffs (14429, -0.89 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.05 DPS) | yes | Black Mageweave Gloves (10003, -1.97 DPS) [crafted]; Gilded Handwraps (254021, -2.67 DPS) [crafted]; Red Mageweave Gloves (10018, -2.70 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (165.7 DPS) | yes | Gilded Cord (254037, -1.07 DPS) [crafted]; Star Belt (4329, -1.30 DPS) [crafted]; Deathmage Sash (10771, -2.10 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (7.61 DPS) | yes | Crimson Silk Pantaloons (7062, -2.01 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.65 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.33 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.12 DPS) | yes | Gilded Slippers (254001, -3.43 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.51 DPS) [dungeon]; Spidersilk Boots (4320, -4.79 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.58 DPS) | yes | Lorekeeper's Ring (19524, -1.54 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.87 DPS) [quest]; Reedknot Ring (9622, -2.21 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (3.23 DPS) | yes | Lorekeeper's Ring (19524, +0.00 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.52 DPS) [quest]; Reedknot Ring (9622, -0.86 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.48 DPS) [dungeon]; Spine of the Devourer (273040, -4.62 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.74 DPS) | yes | Orb of Mystic Insight (249394, -0.93 DPS) [crafted]; Orb of Lorica (11262, -0.96 DPS, sim-verified) [quest]; Beacon of Hope (9393, -2.03 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 120.6 spell_power points (40.82 DPS) | yes | Earthen Rod (9381, -1.75 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.79 DPS) [crafted]; Umbral Wand (5216, -5.15 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 237.4. Weights run: 5.5s. Verify run: 1.7s. 429 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.043 ± 0.008, crit=0.033 ± 0.001 per rating point (14 rating = 1%, 0.466 per %), hit=0.302 ± 0.013 per rating point (10 rating = 1%, 3.023 per %), spell_haste=not significant (-0.082 ± 0.105), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.93 DPS) | yes | Dreamweave Circlet (10041, -1.63 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.99 DPS) [crafted]; Red Mageweave Headband (10033, -2.36 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (2.65 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Scorn's Icy Choker (23169, -0.25 DPS) [dungeon]; Mindburst Medallion (11196, -0.58 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.8 spell_power points (4.56 DPS) | yes | Black Mageweave Shoulders (10027, -0.68 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -1.45 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.47 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (4.72 DPS) | yes | Runecloth Cloak (13860, -1.71 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.74 DPS) [dungeon]; Icy Cloak (4327, -2.40 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (7.36 DPS) | yes | Elemental Raiment (9434, -0.42 DPS) [world_drop]; Acumen Robes (17775, -0.79 DPS) [quest]; Dreamweave Vest (10021, -1.28 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.98 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.56 DPS) [crafted]; Condor Bracers (15864, -0.66 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (6.01 DPS) | yes | Black Mageweave Gloves (10003, -1.04 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.58 DPS) [vendor]; Brightcloth Gloves (14101, -1.71 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.4 spell_power points (4.77 DPS) | yes | Highlander's Cloth Girdle (20098, -0.09 DPS) [rep]; Ghostweave Cord (254073, -0.14 DPS) [crafted]; Star Belt (4329, -0.47 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.29 DPS) | yes | Red Mageweave Pants (10009, -1.43 DPS, sim-verified) [crafted]; Knight's Dreadweave Leggings (220888, -1.99 DPS) [vendor]; Senior Designer's Pantaloons (11841, -2.22 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.94 DPS) | yes | Gilded Sandals (254107, -4.17 DPS) [crafted]; Black Mageweave Boots (10026, -4.20 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.57 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.30 DPS) | yes | Philanthropist's Ring (281635, -0.92 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.65 DPS) [quest]; Reedknot Ring (9622, -1.99 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.97 DPS) | yes | Philanthropist's Ring (281635, -0.68 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.32 DPS) [quest]; Reedknot Ring (9622, -1.65 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.99 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.49 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.66 DPS) [dungeon]; Hanzo Sword (8190, -29.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (237.4 DPS) | yes | Wand of Allistarj (13065, -3.23 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.26 DPS) [crafted]; Pyric Caduceus (11748, -8.81 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Triune Amulet; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 429, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 430.4. Weights run: 6.1s. Verify run: 3.1s. 882 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.076 ± 0.014, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.856 per %), hit=0.337 ± 0.023 per rating point (10 rating = 1%, 3.375 per %), spell_haste=not significant (0.888 ± 0.258), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.903 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.31 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.72 DPS) [pvp]; Deathmist Mask (226909, -1.74 DPS) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.03 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.74 DPS) [quest]; Kezan's Taint (19604, -3.37 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.0 spell_power points (13.22 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.23 DPS) [pvp]; Argent Shoulders (19059, -3.94 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -4.63 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.0 spell_power points (9.11 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -2.52 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -2.73 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (21.29 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.87 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.20 DPS) [pvp]; Robe of the Magi (1716, -11.27 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.6 spell_power points (10.31 DPS) | yes | Runecloth Cuffs (254123, -4.95 DPS) [crafted]; Arcane Runed Bracers (4744, -6.21 DPS) [quest]; Spidertank Oilrag (9448, -6.21 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (12.49 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -2.77 DPS) [pvp]; Earth Warder's Gloves (21318, -4.38 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.0 spell_power points (13.70 DPS) | yes | Stormpike Cloth Girdle (19094, -5.15 DPS) [rep]; Belt of the Archmage (18405, -6.67 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -6.72 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 34.5 spell_power points (15.72 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -2.50 DPS) [pvp]; Ghostweave Pants (14144, -3.86 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.95 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.91 DPS) [crafted]; Omnicast Boots (11822, -1.41 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 32.1 spell_power points (14.62 DPS) | yes | Maiden's Circle (13001, -6.17 DPS) [world_drop]; Lorekeeper's Ring (19522, -8.24 DPS) [rep]; Blessed Band of Light (272407, -8.52 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.8 spell_power points (10.41 DPS) | yes | Maiden's Circle (13001, -2.43 DPS, sim-verified) [world_drop]; Lorekeeper's Ring (19522, -4.03 DPS) [rep]; Blessed Band of Light (272407, -4.31 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+17.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -0.91 DPS) [quest]; Weakness Analyzer (272438, -2.70 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.65 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.50 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -38.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | sim-verified (430.4 DPS) | yes | Brilliant Wand (249385, -0.07 DPS) [crafted]; Stormrager (16997, -1.14 DPS) [quest]; Torch of Light (279246, -5.06 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Wand of Biting Cold

No-known-source sample (15 of 882, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 25520300130201351-0005200000000000000-0550001000000000)

Set DPS (verified): 843.2. Weights run: 7.2s. Verify run: 3.4s. 882 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.078 ± 0.023, crit=0.192 ± 0.008 per rating point (14 rating = 1%, 2.690 per %), hit=0.632 ± 0.047 per rating point (10 rating = 1%, 6.324 per %), spell_haste=not significant (1.688 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.899 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.34 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -0.48 DPS) [quest]; Lieutenant Commander's Dreadweave Cowl (227093, -1.15 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.31 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.62 DPS) [quest]; Kezan's Taint (19604, -4.46 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.9 spell_power points (18.68 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.74 DPS) [pvp]; Mantle of the Timbermaw (19050, -6.15 DPS) [crafted]; Argent Shoulders (19059, -6.76 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.0 spell_power points (13.89 DPS) | yes | Hide of the Wild (18510, -4.94 DPS) [crafted]; Spritecaster Cape (11623, -5.13 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -5.42 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (28.26 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.76 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.19 DPS) [pvp]; Robe of the Magi (1716, -15.12 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.6 spell_power points (13.69 DPS) | yes | Runecloth Cuffs (254123, -6.56 DPS) [crafted]; Arcane Runed Bracers (4744, -8.25 DPS) [quest]; Spidertank Oilrag (9448, -8.25 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (16.58 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -3.68 DPS) [pvp]; Earth Warder's Gloves (21318, -5.81 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.0 spell_power points (20.00 DPS) | yes | Stormpike Cloth Girdle (19094, -8.63 DPS) [rep]; Highlander's Cloth Girdle (20047, -9.61 DPS) [rep]; Belt of the Archmage (18405, -9.86 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 38.2 spell_power points (23.09 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -5.53 DPS) [pvp]; Ghostweave Pants (14144, -7.36 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.52 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.21 DPS) [crafted]; Omnicast Boots (11822, -1.85 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.0 spell_power points (21.20 DPS) | yes | Maiden's Circle (13001, -9.97 DPS) [world_drop]; Blessed Band of Light (272407, -11.32 DPS) [vendor]; Lorekeeper's Ring (19522, -12.73 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.9 spell_power points (13.83 DPS) | yes | Maiden's Circle (13001, -2.61 DPS) [world_drop]; Blessed Band of Light (272407, -3.96 DPS) [vendor]; Lorekeeper's Ring (19522, -5.36 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+21.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -1.21 DPS) [quest]; Weakness Analyzer (272438, -3.55 DPS, sim-verified) [vendor]; Uther's Strength (11302, -4.84 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.64 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -52.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (843.2 DPS) | yes | Wand of Biting Cold (19108, -1.36 DPS) [quest]; Stormrager (16997, -2.43 DPS) [quest]; Torch of Light (279246, -51.01 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 882, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 40.3. Weights run: 2.2s. Verify run: 1.3s. 135 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.019 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.426 per %), hit=0.102 ± 0.003 per rating point (10 rating = 1%, 1.019 per %), spell_haste=not significant (-0.104 ± 0.048), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.746 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.14 DPS) | yes | Red Winter Hat (21524, -2.94 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.19 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.76 DPS) | yes | Battle Healer's Cloak (20427, -0.19 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted]; Mystic Shawl (3449, -0.38 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.95 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -1.35 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.57 DPS) | yes | Owlbeard Bracers (16981, -0.38 DPS) [quest]; Windsong Bangles (263336, -0.49 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.33 DPS) | yes | Gnoll Casting Gloves (892, -0.31 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.57 DPS) [quest]; Pristine Gloves (253913, -0.57 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (40.3 DPS) | yes | Novice Ardent's Sash (253887, -0.38 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Filigreed Pristine Leggings (253937, -0.57 DPS) [crafted]; Rumpled Kilt (274741, -0.76 DPS) [vendor]; Silk-threaded Trousers (1929, -0.81 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.33 DPS) | yes | Pristine Boots (253889, -0.76 DPS) [crafted]; Red Woolen Boots (4313, -0.93 DPS, sim-verified) [crafted]; Soft-soled Linen Boots (4312, -0.95 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.95 DPS) | yes | Ring of the Shadow (1462, -1.22 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.57 DPS) | yes | Ring of the Shadow (1462, -0.78 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 9 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 120.8 spell_power points (22.88 DPS) | yes | Skycaller (12984, -2.18 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.28 DPS) [quest]; Greater Magic Wand (11288, -5.00 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 135, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest

### Band 30 (troll, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 78.2. Weights run: 2.1s. Verify run: 1.3s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.153 ± 0.005, crit=0.036 ± 0.001 per rating point (14 rating = 1%, 0.501 per %), hit=0.133 ± 0.005 per rating point (10 rating = 1%, 1.329 per %), spell_haste=not significant (0.197 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.778 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.46 DPS) | yes | Filigreed Pristine Circlet (253975, -0.67 DPS) [crafted]; Enchanter's Cowl (4322, -0.78 DPS) [crafted]; Silk Headband (7050, -3.40 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (1.77 DPS) | yes | Darkspear Warding Pendant (272075, -1.60 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.63 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.76 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (2.32 DPS) | yes | Invoker's Mantle (215365, -0.58 DPS) [crafted]; Death Speaker Mantle (6685, -0.60 DPS) [dungeon]; Chestnut Mantle (17695, -0.85 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.12 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Battle Healer's Cloak (19529, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.91 DPS) | yes | Beguiler Robes (7728, -0.48 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.71 DPS) [dungeon]; Green Silk Armor (7065, -0.73 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.01 DPS) | yes | Tabitha's Cuffs (251486, -1.65 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -1.72 DPS) [quest]; Glowing Magical Bracelets (13106, -1.74 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.57 DPS) | yes | Gnoll Casting Gloves (892, -0.22 DPS) [world]; Jutebraid Gloves (10654, -0.27 DPS, sim-verified) [quest]; Truefaith Gloves (7049, -0.34 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (2.56 DPS) | yes | Warsong Sash (16975, -0.10 DPS) [quest]; Belt of Arugal (6392, -0.45 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.77 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.69 DPS) | yes | Pristine Leggings (253987, -0.88 DPS) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS, sim-verified) [dungeon]; Silk-threaded Trousers (1929, -1.12 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 spell_power points (1.81 DPS) | yes | Glinteye Slippers (273024, -0.24 DPS) [dungeon]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -2.02 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.57 DPS) | yes | Ring of Power Regulation (273030, -0.22 DPS) [dungeon]; Solid Gold Band (279845, -0.67 DPS) [quest]; Snake Hoop (6750, -1.81 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (78.2 DPS) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Solid Gold Band (279845, -0.45 DPS) [quest]; Snake Hoop (6750, -1.42 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.06 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -7.68 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -7.71 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 147.5 spell_power points (33.00 DPS) | yes | Greater Mystic Wand (217287, -2.88 DPS) [crafted]; Unstable Power Core (279847, -3.29 DPS) [quest]; Wand of Mana Concentration (273043, -3.30 DPS, sim-verified) [world] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 164.2. Weights run: 2.1s. Verify run: 1.4s. 321 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.708 ± 0.015, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.371 per %), hit=0.192 ± 0.010 per rating point (10 rating = 1%, 1.922 per %), spell_haste=0.542 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.11 DPS) | yes | Augural Shroud (2620, -0.99 DPS) [world]; Papal Fez (9431, -1.77 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -2.34 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.7 spell_power points (4.29 DPS) | yes | Scorn's Icy Choker (23169, -1.33 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -1.58 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.61 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.2 spell_power points (5.48 DPS) | yes | Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Berylline Pads (4197, -0.72 DPS) [quest]; Green Silken Shoulders (7057, -1.22 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (3.23 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.59 DPS) [vendor]; Icy Cloak (4327, -0.86 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 spell_power points (8.88 DPS) | yes | Dreamweave Vest (10021, -1.25 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.27 DPS) [crafted]; Elemental Raiment (9434, -1.78 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.68 DPS) [quest]; Radiant Silver Bracers (4545, -1.76 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.05 DPS) | yes | Black Mageweave Gloves (10003, -1.97 DPS) [crafted]; Gilded Handwraps (254021, -2.67 DPS) [crafted]; Red Mageweave Gloves (10018, -2.84 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -1.07 DPS) [crafted]; Star Belt (4329, -1.30 DPS) [crafted]; Deathmage Sash (10771, -2.21 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.5 spell_power points (7.61 DPS) | yes | Crimson Silk Pantaloons (7062, -2.29 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.65 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.33 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.12 DPS) | yes | Gilded Slippers (254001, -4.14 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.51 DPS) [dungeon]; Spidersilk Boots (4320, -4.79 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.58 DPS) | yes | Advisor's Ring (19520, -1.54 DPS) [rep]; Reedknot Ring (9622, -2.21 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (3.23 DPS) | yes | Advisor's Ring (19520, +0.00 DPS, sim-verified) [rep]; Reedknot Ring (9622, -0.86 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.48 DPS) [dungeon]; Spine of the Devourer (273040, -4.62 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.74 DPS) | yes | Omega Orb (7749, -1.69 DPS) [quest]; Beacon of Hope (9393, -2.03 DPS) [dungeon]; Orb of Mystic Insight (249394, -2.64 DPS, sim-verified) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 120.6 spell_power points (40.82 DPS) | yes | Earthen Rod (9381, -1.47 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.79 DPS) [crafted]; Umbral Wand (5216, -5.15 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 321, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 236.4. Weights run: 5.5s. Verify run: 1.6s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.043 ± 0.008, crit=0.033 ± 0.001 per rating point (14 rating = 1%, 0.466 per %), hit=0.302 ± 0.013 per rating point (10 rating = 1%, 3.023 per %), spell_haste=not significant (-0.082 ± 0.105), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.93 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.99 DPS) [crafted]; Dreamweave Circlet (10041, -2.07 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.36 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (2.65 DPS) | yes | Glowing Eye of Mordresh (10769, -0.22 DPS) [dungeon]; Scorn's Icy Choker (23169, -0.25 DPS) [dungeon]; Mindburst Medallion (11196, -0.58 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.8 spell_power points (4.56 DPS) | yes | Black Mageweave Shoulders (10027, -1.39 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -1.45 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.47 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (4.72 DPS) | yes | Deep Woodlands Cloak (19121, -0.63 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.63 DPS) [crafted]; Nightfall Drape (12465, -1.74 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (7.36 DPS) | yes | Acumen Robes (17775, -0.79 DPS) [quest]; Dreamweave Vest (10021, -1.28 DPS) [crafted]; Elemental Raiment (9434, -1.39 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.98 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (6.01 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.58 DPS) [vendor]; Brightcloth Gloves (14101, -1.71 DPS) [crafted]; Black Mageweave Gloves (10003, -2.41 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.4 spell_power points (4.77 DPS) | yes | Ghostweave Cord (254073, -0.14 DPS) [crafted]; Star Belt (4329, -0.47 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.20 DPS, sim-verified) [rep] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (6.29 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.99 DPS) [vendor]; Red Mageweave Pants (10009, -2.13 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -2.22 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.94 DPS) | yes | Gilded Sandals (254107, -4.17 DPS) [crafted]; Black Mageweave Boots (10026, -4.20 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.54 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.30 DPS) | yes | Philanthropist's Ring (281635, -0.92 DPS) [quest]; Reedknot Ring (9622, -1.99 DPS) [quest]; Snake Hoop (6750, -2.24 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.97 DPS) | yes | Philanthropist's Ring (281635, -0.59 DPS) [quest]; Reedknot Ring (9622, -1.65 DPS) [quest]; Snake Hoop (6750, -1.91 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.99 DPS) [world_drop]; Rune of the Guard Captain (19120, -3.27 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.20 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.49 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.66 DPS) [dungeon]; Hanzo Sword (8190, -30.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (236.4 DPS) | yes | Wand of Allistarj (13065, -3.23 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.26 DPS) [crafted]; Pyric Caduceus (11748, -8.85 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Triune Amulet; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 424.6. Weights run: 6.1s. Verify run: 3.1s. 871 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.076 ± 0.014, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.856 per %), hit=0.337 ± 0.023 per rating point (10 rating = 1%, 3.375 per %), spell_haste=not significant (0.888 ± 0.258), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.903 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.31 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.72 DPS) [pvp]; Deathmist Mask (226909, -1.74 DPS) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.03 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.74 DPS) [quest]; Kezan's Taint (19604, -3.37 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.0 spell_power points (13.22 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.23 DPS) [pvp]; Argent Shoulders (19059, -3.49 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -4.63 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.0 spell_power points (9.11 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -2.52 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -2.73 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (21.29 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.87 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.20 DPS) [pvp]; Robe of the Magi (1716, -10.93 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.6 spell_power points (10.31 DPS) | yes | Runecloth Cuffs (254123, -4.95 DPS) [crafted]; Spidertank Oilrag (9448, -6.21 DPS) [dungeon]; Earthenweave Cuffs (254131, -6.66 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (12.49 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -2.77 DPS) [pvp]; Earth Warder's Gloves (21318, -2.99 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.0 spell_power points (13.70 DPS) | yes | Belt of the Archmage (18405, -3.06 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -5.15 DPS) [rep]; Defiler's Cloth Girdle (20163, -6.72 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 34.5 spell_power points (15.72 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -2.50 DPS) [pvp]; Outrider's Silk Leggings (22747, -3.73 DPS, sim-verified) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.95 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.91 DPS) [crafted]; Omnicast Boots (11822, -1.41 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 32.1 spell_power points (14.62 DPS) | yes | Maiden's Circle (13001, -6.17 DPS) [world_drop]; Advisor's Ring (19518, -8.24 DPS) [rep]; Blessed Band of Light (272407, -8.52 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.8 spell_power points (10.41 DPS) | yes | Maiden's Circle (13001, -1.48 DPS, sim-verified) [world_drop]; Advisor's Ring (19518, -4.03 DPS) [rep]; Blessed Band of Light (272407, -4.31 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+16.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -0.91 DPS) [quest]; Weakness Analyzer (272438, -2.65 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.65 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.50 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -37.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | sim-verified (424.6 DPS) | yes | Brilliant Wand (249385, -0.07 DPS) [crafted]; Stormrager (16997, -1.14 DPS) [quest]; Torch of Light (279246, -5.57 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Wand of Biting Cold

No-known-source sample (15 of 871, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25520300130201351-0005200000000000000-0550001000000000)

Set DPS (verified): 833.2. Weights run: 7.2s. Verify run: 3.3s. 871 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.078 ± 0.023, crit=0.192 ± 0.008 per rating point (14 rating = 1%, 2.690 per %), hit=0.632 ± 0.047 per rating point (10 rating = 1%, 6.324 per %), spell_haste=not significant (1.688 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.899 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.34 DPS) | yes | Deathmist Mask (226909, +0.00 DPS, sim-verified) [quest]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.15 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (13.31 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.62 DPS) [quest]; Kezan's Taint (19604, -4.46 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.9 spell_power points (18.68 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.74 DPS) [pvp]; Argent Shoulders (19059, -3.55 DPS) [crafted]; Mantle of the Timbermaw (19050, -6.15 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.0 spell_power points (13.89 DPS) | yes | Hide of the Wild (18510, -4.94 DPS) [crafted]; Spritecaster Cape (11623, -5.13 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -5.42 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (28.26 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.76 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.19 DPS) [pvp]; Robe of the Magi (1716, -13.74 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.6 spell_power points (13.69 DPS) | yes | Runecloth Cuffs (254123, -6.56 DPS) [crafted]; Spidertank Oilrag (9448, -8.25 DPS) [dungeon]; Earthenweave Cuffs (254131, -8.85 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (16.58 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.26 DPS) [quest]; Blood Guard's Dreadweave Handwraps (227099, -3.68 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.0 spell_power points (20.00 DPS) | yes | Belt of the Archmage (18405, -3.74 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -8.63 DPS) [rep]; Defiler's Cloth Girdle (20163, -9.61 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 38.2 spell_power points (23.09 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -3.53 DPS, sim-verified) [rep]; Legionnaire's Dreadweave Legguards (227097, -5.53 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (14.52 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.21 DPS) [crafted]; Omnicast Boots (11822, -1.85 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.0 spell_power points (21.20 DPS) | yes | Maiden's Circle (13001, -9.97 DPS) [world_drop]; Blessed Band of Light (272407, -11.32 DPS) [vendor]; Advisor's Ring (19518, -12.73 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.9 spell_power points (13.83 DPS) | yes | Maiden's Circle (13001, -2.61 DPS) [world_drop]; Blessed Band of Light (272407, -3.96 DPS) [vendor]; Advisor's Ring (19518, -5.36 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+20.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -1.21 DPS) [quest]; Weakness Analyzer (272438, -3.47 DPS, sim-verified) [vendor]; Uther's Strength (11302, -4.84 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.64 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -48.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (833.2 DPS) | yes | Wand of Biting Cold (19108, -1.36 DPS) [quest]; Stormrager (16997, -2.43 DPS) [quest]; Torch of Light (279246, -54.45 DPS, sim-verified) [crafted] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 871, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

