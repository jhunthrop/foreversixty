# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.8. Weights run: 1.6s. Verify run: 1.0s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.202 ± 0.003, crit=0.033 ± 0.001 per rating point (14 rating = 1%, 0.466 per %), hit=0.088 ± 0.003 per rating point (10 rating = 1%, 0.877 per %), spell_haste=not significant (-0.095 ± 0.052), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.756 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.22 DPS) | yes | Red Winter Hat (21524, -2.94 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.01 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.20 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.81 DPS) | yes | Black Whelp Cloak (7283, -0.39 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.41 DPS) [crafted]; Reinforced Linen Cape (2580, -0.61 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.01 DPS) | yes | Green Woolen Vest (2582, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.20 DPS) [dungeon]; Gray Woolen Robe (2585, -1.38 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Ivycloth Bracelets (9793, -0.25 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.42 DPS) | yes | Gnoll Casting Gloves (892, -0.37 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.61 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.01 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.41 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.04 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.20 DPS) [crafted]; Heavy Woolen Pants (4316, -0.61 DPS) [crafted]; Silk-threaded Trousers (1929, -1.58 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.42 DPS) | yes | Red Woolen Boots (4313, -0.43 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.81 DPS) [crafted]; Soft-soled Linen Boots (4312, -1.01 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.01 DPS) | yes | Ring of the Shadow (1462, -1.13 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.61 DPS) | yes | Ring of the Shadow (1462, -0.81 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 89 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -0.86 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (1.01 DPS) | yes | Bouquet of Red Roses (22206, -1.27 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 113.1 spell_power points (22.92 DPS) | yes | Skycaller (12984, -0.74 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.25 DPS) [quest]; Deepblaze (279896, -4.39 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 70.8. Weights run: 4.1s. Verify run: 0.9s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.009 ± 0.003, crit=0.050 ± 0.001 per rating point (14 rating = 1%, 0.703 per %), hit=0.105 ± 0.004 per rating point (10 rating = 1%, 1.051 per %), spell_haste=not significant (0.198 ± 0.063), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.777 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Filigreed Pristine Circlet (253975, -0.68 DPS) [crafted]; Enchanter's Cowl (4322, -1.12 DPS) [crafted]; Silk Headband (7050, -3.81 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.1 spell_power points (0.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.00 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.1 spell_power points (2.07 DPS) | yes | Invoker's Mantle (215365, -0.40 DPS, sim-verified) [crafted]; Moonlit Amice (11884, -0.47 DPS) [quest]; Death Speaker Mantle (6685, -0.68 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.14 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Prelacy Cape (7004, -0.23 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.96 DPS) | yes | Green Silk Armor (7065, -0.30 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.89 DPS) [dungeon]; Mechbuilder's Overalls (9508, -1.11 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.05 DPS) | yes | Glowing Magical Bracelets (13106, -2.03 DPS) [world_drop]; Nightsky Wristbands (6407, -2.04 DPS) [world_drop]; Windsong Bangles (263336, -2.30 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.59 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.45 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.0 spell_power points (2.06 DPS) | yes | Invoker's Cord (215366, -0.45 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.48 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.68 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.73 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.14 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.35 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.1 spell_power points (1.61 DPS) | yes | Glinteye Slippers (273024, -0.01 DPS) [dungeon]; Nimbus Boots (6998, -0.24 DPS) [quest]; Spidersilk Boots (4320, -2.01 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.0 spell_power points (1.38 DPS) | yes | Minor Channeling Ring (1449, -0.23 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.69 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.37 DPS) | yes | Minor Channeling Ring (1449, -0.38 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon]; Runemender's Seal (273031, -0.68 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.20 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.65 DPS) [quest]; Glimmering Staff (249392, -8.18 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 144.3 spell_power points (32.87 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.73 DPS) [crafted]; Gravestone Scepter (7001, -3.87 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 166.7. Weights run: 1.5s. Verify run: 1.1s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.742 ± 0.014, crit=0.121 ± 0.002 per rating point (14 rating = 1%, 1.700 per %), hit=0.161 ± 0.009 per rating point (10 rating = 1%, 1.610 per %), spell_haste=-1.881 ± 0.155, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.35 DPS) | yes | Augural Shroud (2620, -0.90 DPS) [world]; Papal Fez (9431, -1.94 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -2.30 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.9 spell_power points (4.53 DPS) | yes | Triune Amulet (7722, -1.22 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.71 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.97 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.6 spell_power points (5.83 DPS) | yes | Bloodmage Mantle (7684, -0.34 DPS) [dungeon]; Berylline Pads (4197, -0.78 DPS) [quest]; Green Silken Shoulders (7057, -1.64 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.7 spell_power points (3.40 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.95 DPS) [crafted]; Blackforge Cape (6424, -1.32 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.5 spell_power points (9.26 DPS) | yes | Dreamweave Vest (10021, -0.72 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.24 DPS) [crafted]; Elemental Raiment (9434, -1.91 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.15 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.70 DPS) [quest]; Windchaser Cuffs (14429, -0.81 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.0 spell_power points (7.34 DPS) | yes | Black Mageweave Gloves (10003, -2.09 DPS) [crafted]; Red Mageweave Gloves (10018, -2.68 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.72 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (166.7 DPS) | yes | Star Belt (4329, -0.33 DPS) [crafted]; Belt of Arugal (6392, -0.95 DPS) [dungeon]; Deathmage Sash (10771, -3.32 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.9 spell_power points (8.02 DPS) | yes | Crimson Silk Pantaloons (7062, -2.68 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -3.49 DPS) [dungeon]; Pristine Leggings (253987, -3.75 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.40 DPS) | yes | Gilded Slippers (254001, -4.06 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.57 DPS) [dungeon]; Spidersilk Boots (4320, -4.91 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (4.80 DPS) | yes | Ring of Forlorn Spirits (2043, -2.00 DPS) [quest]; Reedknot Ring (9622, -2.35 DPS) [quest]; Minor Channeling Ring (1449, -2.53 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (3.40 DPS) | yes | Ring of Forlorn Spirits (2043, +0.00 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.95 DPS) [quest]; Minor Channeling Ring (1449, -1.13 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.52 DPS) [dungeon]; Spine of the Devourer (273040, -4.88 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.90 DPS) | yes | Orb of Mystic Insight (249394, -0.89 DPS) [crafted]; Orb of Lorica (11262, -0.95 DPS, sim-verified) [quest]; Beacon of Hope (9393, -2.10 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 116.9 spell_power points (40.94 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.84 DPS) [crafted]; Umbral Wand (5216, -5.27 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 239.9. Weights run: 1.5s. Verify run: 1.2s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.157 ± 0.006, crit=0.140 ± 0.002 per rating point (14 rating = 1%, 1.957 per %), hit=0.186 ± 0.008 per rating point (10 rating = 1%, 1.863 per %), spell_haste=-1.861 ± 0.110, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.885 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.36 DPS) | yes | Dreamweave Circlet (10041, -1.70 DPS) [crafted]; Red Mageweave Headband (10033, -1.87 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.30 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.3 spell_power points (3.17 DPS) | yes | Mindburst Medallion (11196, -0.50 DPS) [quest]; Triune Amulet (7722, -1.07 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.33 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 15.8 spell_power points (6.07 DPS) | yes | Black Mageweave Shoulders (10027, -1.44 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.71 DPS) [vendor]; Bloodmage Mantle (7684, -2.08 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.9 spell_power points (5.73 DPS) | yes | Nightfall Drape (12465, -2.28 DPS) [dungeon]; Icy Cloak (4327, -3.05 DPS) [crafted]; Runecloth Cloak (13860, -3.10 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (8.80 DPS) | yes | Elemental Raiment (9434, -0.74 DPS) [world_drop]; Dreamweave Vest (10021, -1.35 DPS) [crafted]; Acumen Robes (17775, -1.47 DPS, sim-verified) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.45 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.35 DPS) [crafted]; Condor Bracers (15864, -0.77 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (7.15 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.62 DPS) [vendor]; Runecloth Gloves (13863, -2.00 DPS) [crafted]; Black Mageweave Gloves (10003, -2.24 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (5.97 DPS) | yes | Ghostweave Cord (254073, +0.00 DPS, sim-verified) [crafted]; Star Belt (4329, -0.98 DPS) [crafted]; Gilded Waistcord (254081, -1.59 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (7.29 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.21 DPS) [vendor]; Red Mageweave Pants (10009, -1.54 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -2.27 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.21 DPS) | yes | Gilded Sandals (254107, -1.67 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -4.57 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.88 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.99 DPS) | yes | Ring of Forlorn Spirits (2043, -1.92 DPS) [quest]; Reedknot Ring (9622, -2.30 DPS) [quest]; Snake Hoop (6750, -2.39 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (4.14 DPS) | yes | Ring of Forlorn Spirits (2043, -1.06 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.45 DPS) [quest]; Snake Hoop (6750, -1.53 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (239.9 DPS) | yes | Frozen Heart of the Mountain (249469, -1.92 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (239.9 DPS) | yes | Frozen Heart of the Mountain (249469, -1.66 DPS) [crafted]; Mark of the Chosen (17774, -4.80 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (239.9 DPS) | yes | Kindling Stave (11750, -0.05 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.77 DPS) [dungeon]; Hanzo Sword (8190, -33.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 136.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.68 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 374.6. Weights run: 1.5s. Verify run: 3.1s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.110 ± 0.024, crit=0.182 ± 0.007 per rating point (14 rating = 1%, 2.553 per %), hit=0.280 ± 0.027 per rating point (10 rating = 1%, 2.798 per %), spell_haste=not significant (0.562 ± 0.343), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.901 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.18 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.67 DPS) [pvp]; Dreamweave Circlet (10041, -2.21 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.1 spell_power points (5.91 DPS) | yes | Glowing Eye of Mordresh (10769, -2.35 DPS) [dungeon]; Mindburst Medallion (11196, -2.90 DPS) [quest]; Triune Amulet (7722, -3.42 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (11.28 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.21 DPS, sim-verified) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -4.07 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.7 spell_power points (6.61 DPS) | yes | Runecloth Cloak (13860, -2.01 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.55 DPS) [dungeon]; Icy Cloak (4327, -3.45 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (21.20 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.57 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -8.93 DPS) [pvp]; Robe of the Magi (1716, -11.89 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (374.6 DPS) | yes | Arcane Runed Bracers (4744, +0.00 DPS) [quest]; Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Runecloth Cuffs (254123, -7.98 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (374.6 DPS) | yes | Knight-Lieutenant's Dreadweave Handwraps (227100, +0.00 DPS) [pvp]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -7.31 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (374.6 DPS) | yes | Star Belt (4329, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -7.73 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.73 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -6.12 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.83 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.90 DPS) [crafted]; Omnicast Boots (11822, -1.21 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.8 spell_power points (8.47 DPS) | yes | Philanthropist's Ring (281635, -3.71 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.86 DPS) [quest]; Reedknot Ring (9622, -5.31 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.86 DPS) | yes | Philanthropist's Ring (281635, -1.11 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.26 DPS) [quest]; Reedknot Ring (9622, -2.71 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (374.6 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (374.6 DPS) | yes | Mark of the Chosen (17774, -5.93 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (374.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.31 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -43.53 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 176.1 spell_power points (79.47 DPS) | yes | Brilliant Wand (249385, -3.33 DPS, sim-verified) [crafted]; Stormrager (16997, -16.53 DPS) [quest]; Greater Eternal Wand (249237, -22.08 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25520300130201351-0005200000000000000-0550001000000000)

Set DPS (verified): 743.2. Weights run: 1.9s. Verify run: 2.2s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.164 ± 0.035, crit=0.357 ± 0.011 per rating point (14 rating = 1%, 5.003 per %), hit=0.508 ± 0.045 per rating point (10 rating = 1%, 5.083 per %), spell_haste=not significant (0.668 ± 0.610), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.896 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.88 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -3.99 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.6 spell_power points (8.53 DPS) | yes | Glowing Eye of Mordresh (10769, -3.33 DPS) [dungeon]; Triune Amulet (7722, -3.53 DPS) [dungeon]; Mindburst Medallion (11196, -4.16 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (15.63 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.54 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.67 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (9.37 DPS) | yes | Nightfall Drape (12465, -3.74 DPS) [dungeon]; Icy Cloak (4327, -4.99 DPS) [crafted]; Runecloth Cloak (13860, -5.04 DPS, sim-verified) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.5 spell_power points (29.69 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.22 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.00 DPS) [pvp]; Robe of the Magi (1716, -18.92 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.6 spell_power points (7.91 DPS) | yes | Arcane Runed Bracers (4744, -2.28 DPS) [quest]; Spidertank Oilrag (9448, -2.28 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.81 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.76 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.21 DPS) [pvp]; Gloves of Spell Mastery (14146, -0.85 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (9.78 DPS) | yes | Ghostweave Cord (254073, -1.03 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.47 DPS) [crafted]; Star Belt (4329, -1.65 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (16.26 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -9.99 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.01 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.27 DPS) [dungeon]; Venomspew Footpads (275606, -3.41 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.1 spell_power points (11.97 DPS) | yes | Philanthropist's Ring (281635, -5.21 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.97 DPS) [quest]; Reedknot Ring (9622, -7.60 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (8.13 DPS) | yes | Ring of Forlorn Spirits (2043, -3.13 DPS) [quest]; Reedknot Ring (9622, -3.75 DPS) [quest]; Philanthropist's Ring (281635, -5.33 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (743.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (743.2 DPS) | yes | Mark of the Chosen (17774, -7.84 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (743.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.02 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -56.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 128.7 spell_power points (80.51 DPS) | yes | Brilliant Wand (249385, -4.62 DPS, sim-verified) [crafted]; Stormrager (16997, -17.31 DPS) [quest]; Greater Eternal Wand (249237, -21.55 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.4. Weights run: 1.6s. Verify run: 1.0s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.202 ± 0.003, crit=0.033 ± 0.001 per rating point (14 rating = 1%, 0.466 per %), hit=0.088 ± 0.003 per rating point (10 rating = 1%, 0.877 per %), spell_haste=not significant (-0.095 ± 0.052), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.756 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.22 DPS) | yes | Red Winter Hat (21524, -2.83 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.01 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.20 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.72 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.81 DPS) | yes | Black Whelp Cloak (7283, -0.38 DPS, sim-verified) [crafted]; Mystic Shawl (3449, -0.41 DPS) [quest]; Pearl-clasped Cloak (5542, -0.41 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.01 DPS) | yes | Green Woolen Vest (2582, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.20 DPS) [dungeon]; Gray Woolen Robe (2585, -1.37 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.61 DPS) | yes | Owlbeard Bracers (16981, -0.41 DPS) [quest]; Windsong Bangles (263336, -0.47 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.42 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.61 DPS) [quest]; Pristine Gloves (253913, -0.61 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.41 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.19 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.20 DPS) [crafted]; Embroidered Leggings (270011, -0.41 DPS) [quest]; Silk-threaded Trousers (1929, -1.84 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.42 DPS) | yes | Pristine Boots (253889, -0.81 DPS) [crafted]; Red Woolen Boots (4313, -0.83 DPS, sim-verified) [crafted]; Soft-soled Linen Boots (4312, -1.01 DPS) [crafted] |
| finger1 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.61 DPS) | yes | - |
| finger2 | - | - |  |  |  |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Evocator's Blade (2567, -0.49 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 113.1 spell_power points (22.92 DPS) | yes | Skycaller (12984, -0.81 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.25 DPS) [quest]; Greater Magic Wand (11288, -5.01 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 69.7. Weights run: 4.1s. Verify run: 0.9s. 283 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.009 ± 0.003, crit=0.050 ± 0.001 per rating point (14 rating = 1%, 0.703 per %), hit=0.105 ± 0.004 per rating point (10 rating = 1%, 1.051 per %), spell_haste=not significant (0.198 ± 0.063), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.777 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.51 DPS) | yes | Filigreed Pristine Circlet (253975, -0.68 DPS) [crafted]; Enchanter's Cowl (4322, -1.12 DPS) [crafted]; Silk Headband (7050, -3.71 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.1 spell_power points (0.01 DPS) | yes | Crystal Starfire Medallion (5003, -0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.00 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.1 spell_power points (2.07 DPS) | yes | Invoker's Mantle (215365, -0.46 DPS) [crafted]; Death Speaker Mantle (6685, -0.68 DPS) [dungeon]; Chestnut Mantle (17695, -1.21 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.14 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.23 DPS) [crafted]; Black Whelp Cloak (7283, -0.46 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.96 DPS) | yes | Beguiler Robes (7728, -0.89 DPS) [dungeon]; Green Silk Armor (7065, -1.02 DPS, sim-verified) [crafted]; Mechbuilder's Overalls (9508, -1.11 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.05 DPS) | yes | Tabitha's Cuffs (251486, -1.75 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -1.82 DPS) [quest]; Windsong Bangles (263336, -1.82 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.59 DPS) | yes | Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.45 DPS) [crafted]; Jutebraid Gloves (10654, -0.46 DPS, sim-verified) [quest] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.51 DPS) | yes | Belt of Arugal (6392, -0.26 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.68 DPS) [dungeon]; Invoker's Cord (215366, -0.90 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.73 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.14 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.35 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.1 spell_power points (1.61 DPS) | yes | Glinteye Slippers (273024, -0.01 DPS) [dungeon]; Acidic Walkers (9454, -0.45 DPS) [dungeon]; Spidersilk Boots (4320, -2.23 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.0 spell_power points (1.38 DPS) | yes | Solid Gold Band (279845, -0.47 DPS) [quest]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Runemender's Seal (273031, -0.69 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.37 DPS) | yes | Solid Gold Band (279845, -0.46 DPS) [quest]; Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon]; Runemender's Seal (273031, -0.68 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.20 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -8.18 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -8.18 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 144.3 spell_power points (32.87 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.73 DPS) [crafted]; Unstable Power Core (279847, -3.16 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 165.4. Weights run: 1.5s. Verify run: 1.1s. 381 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.742 ± 0.014, crit=0.121 ± 0.002 per rating point (14 rating = 1%, 1.700 per %), hit=0.161 ± 0.009 per rating point (10 rating = 1%, 1.610 per %), spell_haste=-1.881 ± 0.155, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.877 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.35 DPS) | yes | Augural Shroud (2620, -0.90 DPS) [world]; Papal Fez (9431, -1.87 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -2.30 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.9 spell_power points (4.53 DPS) | yes | Triune Amulet (7722, -1.29 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.71 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.97 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.6 spell_power points (5.83 DPS) | yes | Bloodmage Mantle (7684, -0.34 DPS) [dungeon]; Berylline Pads (4197, -0.78 DPS) [quest]; Green Silken Shoulders (7057, -2.27 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.7 spell_power points (3.40 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.95 DPS) [crafted]; Blackforge Cape (6424, -1.32 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.5 spell_power points (9.26 DPS) | yes | Dreamweave Vest (10021, -0.99 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.24 DPS) [crafted]; Elemental Raiment (9434, -1.91 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Condor Bracers (15864, -0.70 DPS) [quest]; Windchaser Cuffs (14429, -0.81 DPS) [world_drop]; Radiant Silver Bracers (4545, -1.87 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.0 spell_power points (7.34 DPS) | yes | Black Mageweave Gloves (10003, -2.09 DPS) [crafted]; Gilded Handwraps (254021, -2.72 DPS) [crafted]; Red Mageweave Gloves (10018, -3.26 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.33 DPS) [crafted]; Belt of Arugal (6392, -0.95 DPS) [dungeon]; Deathmage Sash (10771, -2.92 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.9 spell_power points (8.02 DPS) | yes | Crimson Silk Pantaloons (7062, -2.53 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -3.49 DPS) [dungeon]; Pristine Leggings (253987, -3.75 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.40 DPS) | yes | Gilded Slippers (254001, -3.83 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.57 DPS) [dungeon]; Spidersilk Boots (4320, -4.91 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (4.80 DPS) | yes | Reedknot Ring (9622, -2.35 DPS) [quest]; Ring of Power Regulation (273030, -2.70 DPS) [dungeon]; Voodoo Band (1996, -2.98 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (3.40 DPS) | yes | Reedknot Ring (9622, -0.95 DPS) [quest]; Ring of Power Regulation (273030, -1.30 DPS) [dungeon]; Voodoo Band (1996, -1.58 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.52 DPS) [dungeon]; Spine of the Devourer (273040, -4.88 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.90 DPS) | yes | Omega Orb (7749, -1.75 DPS) [quest]; Beacon of Hope (9393, -2.10 DPS) [dungeon]; Orb of Mystic Insight (249394, -2.50 DPS, sim-verified) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 116.9 spell_power points (40.94 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.84 DPS) [crafted]; Umbral Wand (5216, -5.27 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 381, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 235.9. Weights run: 1.5s. Verify run: 1.2s. 492 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.157 ± 0.006, crit=0.140 ± 0.002 per rating point (14 rating = 1%, 1.957 per %), hit=0.186 ± 0.008 per rating point (10 rating = 1%, 1.863 per %), spell_haste=-1.861 ± 0.110, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.885 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.36 DPS) | yes | Dreamweave Circlet (10041, -1.37 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -1.87 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.30 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.3 spell_power points (3.17 DPS) | yes | Mindburst Medallion (11196, -0.50 DPS) [quest]; Triune Amulet (7722, -1.50 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.33 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 15.8 spell_power points (6.07 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.71 DPS) [vendor]; Bloodmage Mantle (7684, -2.08 DPS) [dungeon]; Black Mageweave Shoulders (10027, -2.13 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.9 spell_power points (5.73 DPS) | yes | Deep Woodlands Cloak (19121, -1.18 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.80 DPS) [crafted]; Nightfall Drape (12465, -2.28 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (8.80 DPS) | yes | Acumen Robes (17775, -0.31 DPS) [quest]; Elemental Raiment (9434, -0.74 DPS) [world_drop]; Dreamweave Vest (10021, -1.35 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.45 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.77 DPS) [quest]; Bloodband Bracers (11469, -0.99 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (7.15 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.62 DPS) [vendor]; Runecloth Gloves (13863, -2.00 DPS) [crafted]; Black Mageweave Gloves (10003, -2.34 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (5.97 DPS) | yes | Ghostweave Cord (254073, +0.00 DPS, sim-verified) [crafted]; Star Belt (4329, -0.98 DPS) [crafted]; Gilded Waistcord (254081, -1.59 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (7.29 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.21 DPS) [vendor]; Red Mageweave Pants (10009, -1.40 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -2.27 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.21 DPS) | yes | Gilded Sandals (254107, -1.73 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -4.57 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.88 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.99 DPS) | yes | Reedknot Ring (9622, -2.30 DPS) [quest]; Snake Hoop (6750, -2.39 DPS) [quest]; Ring of Power Regulation (273030, -2.69 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (4.14 DPS) | yes | Snake Hoop (6750, -1.53 DPS) [quest]; Reedknot Ring (9622, -1.77 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -1.84 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (235.9 DPS) | yes | Rune of the Guard Captain (19120, -1.80 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (235.9 DPS) | yes | Rune of the Guard Captain (19120, -1.61 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (235.9 DPS) | yes | Kindling Stave (11750, -0.05 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.77 DPS) [dungeon]; Hanzo Sword (8190, -34.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 136.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.68 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 492, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 364.9. Weights run: 1.5s. Verify run: 3.2s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.110 ± 0.024, crit=0.182 ± 0.007 per rating point (14 rating = 1%, 2.553 per %), hit=0.280 ± 0.027 per rating point (10 rating = 1%, 2.798 per %), spell_haste=not significant (0.562 ± 0.343), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.901 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.18 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.67 DPS) [pvp]; Dreamweave Circlet (10041, -2.02 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.1 spell_power points (5.91 DPS) | yes | Glowing Eye of Mordresh (10769, -2.35 DPS) [dungeon]; Mindburst Medallion (11196, -2.90 DPS) [quest]; Triune Amulet (7722, -3.47 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (11.28 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.70 DPS, sim-verified) [crafted]; Champion's Dreadweave Spaulders (227092, -4.07 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.7 spell_power points (6.61 DPS) | yes | Runecloth Cloak (13860, -2.16 DPS) [crafted]; Nightfall Drape (12465, -2.55 DPS) [dungeon]; Deep Woodlands Cloak (19121, -3.21 DPS, sim-verified) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (21.20 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.57 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -8.93 DPS) [pvp]; Robe of the Magi (1716, -10.31 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (364.9 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -8.89 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (364.9 DPS) | yes | Blood Guard's Dreadweave Handwraps (227099, +0.00 DPS) [pvp]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -5.55 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (364.9 DPS) | yes | Star Belt (4329, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -8.32 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.73 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -6.85 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.83 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.21 DPS) [dungeon]; Venomspew Footpads (275606, -1.55 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.8 spell_power points (8.47 DPS) | yes | Philanthropist's Ring (281635, -3.71 DPS) [quest]; Reedknot Ring (9622, -5.31 DPS) [quest]; Snake Hoop (6750, -5.51 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.86 DPS) | yes | Philanthropist's Ring (281635, -1.11 DPS) [quest]; Reedknot Ring (9622, -2.71 DPS) [quest]; Snake Hoop (6750, -2.91 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (364.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (364.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (364.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.31 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -42.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 176.1 spell_power points (79.47 DPS) | yes | Brilliant Wand (249385, -1.53 DPS, sim-verified) [crafted]; Stormrager (16997, -16.53 DPS) [quest]; Greater Eternal Wand (249237, -22.08 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25520300130201351-0005200000000000000-0550001000000000)

Set DPS (verified): 725.0. Weights run: 1.9s. Verify run: 2.5s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.164 ± 0.035, crit=0.357 ± 0.011 per rating point (14 rating = 1%, 5.003 per %), hit=0.508 ± 0.045 per rating point (10 rating = 1%, 5.083 per %), spell_haste=not significant (0.668 ± 0.610), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.896 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (16.88 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.73 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.6 spell_power points (8.53 DPS) | yes | Glowing Eye of Mordresh (10769, -3.33 DPS) [dungeon]; Triune Amulet (7722, -3.53 DPS) [dungeon]; Mindburst Medallion (11196, -4.16 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (15.63 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.54 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -3.67 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (9.37 DPS) | yes | Deep Woodlands Cloak (19121, -0.94 DPS) [quest]; Runecloth Cloak (13860, -2.92 DPS) [crafted]; Nightfall Drape (12465, -3.74 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.5 spell_power points (29.69 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.22 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.00 DPS) [pvp]; Robe of the Magi (1716, -15.44 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.6 spell_power points (7.91 DPS) | yes | Spidertank Oilrag (9448, -2.28 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.81 DPS) [crafted]; Earthenweave Cuffs (254131, -2.90 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (13.76 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS, sim-verified) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.21 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (9.78 DPS) | yes | Ghostweave Cord (254073, +0.00 DPS, sim-verified) [crafted]; Wisdom of the Timbermaw (19047, -1.47 DPS) [crafted]; Star Belt (4329, -1.65 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (16.26 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -3.00 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.01 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.25 DPS) [crafted]; Omnicast Boots (11822, -1.27 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.1 spell_power points (11.97 DPS) | yes | Philanthropist's Ring (281635, -5.21 DPS) [quest]; Reedknot Ring (9622, -7.60 DPS) [quest]; Snake Hoop (6750, -7.71 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (8.13 DPS) | yes | Philanthropist's Ring (281635, -1.36 DPS) [quest]; Reedknot Ring (9622, -3.75 DPS) [quest]; Snake Hoop (6750, -3.86 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (725.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (725.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (725.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.02 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -53.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 128.7 spell_power points (80.51 DPS) | yes | Brilliant Wand (249385, -14.59 DPS) [crafted]; Stormrager (16997, -17.31 DPS) [quest]; Greater Eternal Wand (249237, -21.55 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

