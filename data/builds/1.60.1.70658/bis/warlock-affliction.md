# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 45.4. Weights run: 2.2s. Verify run: 1.2s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.078 ± 0.003, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.428 per %), hit=0.066 ± 0.003 per rating point (10 rating = 1%, 0.660 per %), spell_haste=not significant (0.012 ± 0.037), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.807 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.53 DPS) | yes | Red Winter Hat (21524, -3.46 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.28 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (1.02 DPS) | yes | Black Whelp Cloak (7283, -0.33 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.51 DPS) [crafted]; Reinforced Linen Cape (2580, -0.77 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.28 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -1.52 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.26 DPS) | yes | Ivycloth Bracelets (9793, -0.29 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.79 DPS) | yes | Gnoll Casting Gloves (892, -0.43 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.77 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.28 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.51 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.12 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.26 DPS) [crafted]; Heavy Woolen Pants (4316, -0.77 DPS) [crafted]; Silk-threaded Trousers (1929, -1.53 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.79 DPS) | yes | Red Woolen Boots (4313, -0.81 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.02 DPS) [crafted]; Soft-soled Linen Boots (4312, -1.28 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.28 DPS) | yes | Ring of the Shadow (1462, -1.41 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.77 DPS) | yes | Ring of the Shadow (1462, -0.85 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 89 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.30 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (1.28 DPS) | yes | Bouquet of Red Roses (22206, -1.37 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 90.4 spell_power points (23.08 DPS) | yes | Skycaller (12984, -0.87 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.15 DPS) [quest]; Deepblaze (279896, -4.55 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 82.8. Weights run: 2.2s. Verify run: 1.1s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.018 ± 0.003, crit=0.053 ± 0.001 per rating point (14 rating = 1%, 0.746 per %), hit=0.098 ± 0.004 per rating point (10 rating = 1%, 0.982 per %), spell_haste=-0.425 ± 0.062, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.818 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.96 DPS) | yes | Filigreed Pristine Circlet (253975, -0.81 DPS) [crafted]; Enchanter's Cowl (4322, -1.30 DPS) [crafted]; Silk Headband (7050, -4.16 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.1 spell_power points (0.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.01 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.2 spell_power points (2.46 DPS) | yes | Moonlit Amice (11884, -0.58 DPS) [quest]; Invoker's Mantle (215365, -0.78 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.80 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.34 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Prelacy Cape (7004, -0.27 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.50 DPS) | yes | Green Silk Armor (7065, -0.90 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.02 DPS) [dungeon]; Mechbuilder's Overalls (9508, -1.29 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.42 DPS) | yes | Glowing Magical Bracelets (13106, -2.38 DPS) [world_drop]; Nightsky Wristbands (6407, -2.39 DPS) [world_drop]; Windsong Bangles (263336, -2.60 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.88 DPS) | yes | Gnoll Casting Gloves (892, -0.27 DPS) [world]; Serpent Gloves (5970, -0.40 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.52 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.1 spell_power points (2.44 DPS) | yes | Invoker's Cord (215366, -0.53 DPS) [crafted]; Crimson Silk Belt (7055, -0.79 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.81 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.23 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.34 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.59 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.1 spell_power points (1.92 DPS) | yes | Glinteye Slippers (273024, -0.03 DPS) [dungeon]; Nimbus Boots (6998, -0.30 DPS) [quest]; Spidersilk Boots (4320, -2.34 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.1 spell_power points (1.64 DPS) | yes | Minor Channeling Ring (1449, -0.28 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Runemender's Seal (273031, -0.83 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.61 DPS) | yes | Minor Channeling Ring (1449, -0.38 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.81 DPS) [dungeon]; Runemender's Seal (273031, -0.81 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (9.68 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -4.30 DPS) [quest]; Glimmering Staff (249392, -9.63 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 122.2 spell_power points (32.88 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.53 DPS) [crafted]; Gravestone Scepter (7001, -3.88 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 200.2. Weights run: 2.0s. Verify run: 1.5s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.651 ± 0.013, crit=0.134 ± 0.002 per rating point (14 rating = 1%, 1.879 per %), hit=0.173 ± 0.011 per rating point (10 rating = 1%, 1.733 per %), spell_haste=-0.973 ± 0.157, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.908 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.81 DPS) | yes | Papal Fez (9431, -1.11 DPS, sim-verified) [dungeon]; Augural Shroud (2620, -1.46 DPS) [world]; Whitemane's Chapeau (7720, -3.14 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.2 spell_power points (5.12 DPS) | yes | Triune Amulet (7722, -1.53 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -3.21 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -3.48 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.5 spell_power points (6.48 DPS) | yes | Bloodmage Mantle (7684, -0.25 DPS) [dungeon]; Berylline Pads (4197, -0.82 DPS) [quest]; Green Silken Shoulders (7057, -2.24 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 spell_power points (3.88 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.94 DPS) [crafted]; Blackforge Cape (6424, -1.70 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.9 spell_power points (10.86 DPS) | yes | Dreamweave Vest (10021, -0.72 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.72 DPS) [crafted]; Elemental Raiment (9434, -2.06 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.77 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.84 DPS) [quest]; Windchaser Cuffs (14429, -1.32 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (8.64 DPS) | yes | Black Mageweave Gloves (10003, -2.35 DPS) [crafted]; Red Mageweave Gloves (10018, -3.11 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -3.37 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.09 DPS) [crafted]; Belt of Arugal (6392, -0.94 DPS) [dungeon]; Deathmage Sash (10771, -3.60 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.8 spell_power points (9.14 DPS) | yes | Crimson Silk Pantaloons (7062, -2.66 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -4.03 DPS) [dungeon]; Gaze Dreamer Pants (6903, -4.11 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.06 DPS) | yes | Gilded Slippers (254001, -4.00 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -5.78 DPS) [dungeon]; Spidersilk Boots (4320, -6.04 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (5.56 DPS) | yes | Ring of Forlorn Spirits (2043, -2.20 DPS) [quest]; Reedknot Ring (9622, -2.62 DPS) [quest]; Minor Channeling Ring (1449, -2.92 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.3 spell_power points (3.88 DPS) | yes | Ring of Forlorn Spirits (2043, -0.53 DPS) [quest]; Reedknot Ring (9622, -0.94 DPS) [quest]; Minor Channeling Ring (1449, -1.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.55 DPS) [dungeon]; Spine of the Devourer (273040, -5.54 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (5.87 DPS) | yes | Orb of Lorica (11262, -1.27 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -1.30 DPS) [crafted]; Beacon of Hope (9393, -2.52 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Twisted Nether Wand (249144, -0.33 DPS) [crafted]; Umbral Wand (5216, -1.18 DPS) [dungeon]; Jaina's Firestarter (13064, -2.24 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Earthen Rod

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 288.2. Weights run: 2.1s. Verify run: 1.6s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.166 ± 0.008, crit=0.163 ± 0.002 per rating point (14 rating = 1%, 2.283 per %), hit=0.190 ± 0.010 per rating point (10 rating = 1%, 1.903 per %), spell_haste=-0.726 ± 0.120, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.916 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.34 DPS) | yes | Dreamweave Circlet (10041, -0.69 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.14 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.74 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.3 spell_power points (3.81 DPS) | yes | Mindburst Medallion (11196, -0.61 DPS) [quest]; Triune Amulet (7722, -2.13 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.75 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.0 spell_power points (7.31 DPS) | yes | Black Mageweave Shoulders (10027, -1.56 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.92 DPS) [vendor]; Bloodmage Mantle (7684, -2.51 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (6.85 DPS) | yes | Nightfall Drape (12465, -2.74 DPS) [dungeon]; Runecloth Cloak (13860, -2.93 DPS, sim-verified) [crafted]; Icy Cloak (4327, -3.65 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (10.51 DPS) | yes | Acumen Robes (17775, -0.31 DPS) [quest]; Elemental Raiment (9434, -0.91 DPS) [world_drop]; Dreamweave Vest (10021, -1.60 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (4.11 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.38 DPS) [crafted]; Condor Bracers (15864, -0.91 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (8.53 DPS) | yes | Black Mageweave Gloves (10003, -1.78 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.91 DPS) [vendor]; Runecloth Gloves (13863, -2.36 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.7 spell_power points (7.16 DPS) | yes | Ghostweave Cord (254073, -0.76 DPS) [crafted]; Star Belt (4329, -1.21 DPS) [crafted]; Gilded Waistcord (254081, -1.90 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (8.69 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.25 DPS) [vendor]; Red Mageweave Pants (10009, -1.38 DPS) [crafted]; Senior Designer's Pantaloons (11841, -2.67 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.97 DPS) | yes | Gilded Sandals (254107, -1.41 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -5.41 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -5.76 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.94 DPS) | yes | Ring of Forlorn Spirits (2043, -2.29 DPS) [quest]; Reedknot Ring (9622, -2.74 DPS) [quest]; Snake Hoop (6750, -2.82 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (4.95 DPS) | yes | Ring of Forlorn Spirits (2043, -1.29 DPS) [quest]; Reedknot Ring (9622, -1.75 DPS) [quest]; Snake Hoop (6750, -1.83 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (288.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (288.2 DPS) | yes | Mark of the Chosen (17774, -2.93 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (288.2 DPS) | yes | Kindling Stave (11750, -0.01 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.91 DPS) [dungeon]; Hanzo Sword (8190, -40.30 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 114.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.09 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 388.7. Weights run: 2.1s. Verify run: 2.5s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.147 ± 0.009, crit=0.195 ± 0.003 per rating point (14 rating = 1%, 2.730 per %), hit=0.243 ± 0.011 per rating point (10 rating = 1%, 2.426 per %), spell_haste=not significant (0.112 ± 0.117), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.912 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.76 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.30 DPS) [pvp]; Dreamweave Circlet (10041, -1.97 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.5 spell_power points (6.37 DPS) | yes | Glowing Eye of Mordresh (10769, -2.50 DPS) [dungeon]; Triune Amulet (7722, -2.58 DPS) [dungeon]; Mindburst Medallion (11196, -3.11 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (11.82 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.58 DPS, sim-verified) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.95 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.9 spell_power points (7.04 DPS) | yes | Runecloth Cloak (13860, -2.18 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.78 DPS) [dungeon]; Icy Cloak (4327, -3.73 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.3 spell_power points (22.37 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.58 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.16 DPS) [pvp]; Robe of the Magi (1716, -11.53 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.5 spell_power points (5.89 DPS) | yes | Spidertank Oilrag (9448, -1.64 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.10 DPS) [crafted]; Arcane Runed Bracers (4744, -2.73 DPS, sim-verified) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (10.40 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.20 DPS) [pvp]; Dreamweave Gloves (10019, -1.33 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.5 spell_power points (7.31 DPS) | yes | Ghostweave Cord (254073, -0.69 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.14 DPS) [crafted]; Star Belt (4329, -1.17 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (12.29 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.62 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.35 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.95 DPS) [crafted]; Omnicast Boots (11822, -1.06 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.0 spell_power points (9.00 DPS) | yes | Philanthropist's Ring (281635, -3.92 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.21 DPS) [quest]; Reedknot Ring (9622, -5.69 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.15 DPS) | yes | Philanthropist's Ring (281635, -1.81 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -2.36 DPS) [quest]; Reedknot Ring (9622, -2.84 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (388.7 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (388.7 DPS) | yes | Mark of the Chosen (17774, -6.05 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (388.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.11 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -43.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 168.4 spell_power points (79.60 DPS) | yes | Brilliant Wand (249385, -2.23 DPS, sim-verified) [crafted]; Stormrager (16997, -16.56 DPS) [quest]; Greater Eternal Wand (249237, -22.01 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25520300130201351-0005200000000000000-0550001000000000)

Set DPS (verified): 767.6. Weights run: 2.6s. Verify run: 2.8s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.160 ± 0.016, crit=0.362 ± 0.007 per rating point (14 rating = 1%, 5.068 per %), hit=0.465 ± 0.025 per rating point (10 rating = 1%, 4.654 per %), spell_haste=not significant (0.871 ± 0.304), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.906 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (17.47 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.84 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.6 spell_power points (8.80 DPS) | yes | Glowing Eye of Mordresh (10769, -3.44 DPS) [dungeon]; Triune Amulet (7722, -3.63 DPS) [dungeon]; Mindburst Medallion (11196, -4.30 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (16.18 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.55 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.78 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (9.68 DPS) | yes | Nightfall Drape (12465, -3.86 DPS) [dungeon]; Runecloth Cloak (13860, -4.03 DPS, sim-verified) [crafted]; Icy Cloak (4327, -5.15 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.4 spell_power points (30.70 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -7.50 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.45 DPS) [pvp]; Robe of the Magi (1716, -15.10 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.6 spell_power points (8.16 DPS) | yes | Spidertank Oilrag (9448, -2.33 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.90 DPS) [crafted]; Arcane Runed Bracers (4744, -4.84 DPS, sim-verified) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (14.24 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS, sim-verified) [crafted]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.23 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (10.10 DPS) | yes | Ghostweave Cord (254073, -1.04 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.53 DPS) [crafted]; Star Belt (4329, -1.69 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (16.82 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -7.51 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.53 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.29 DPS) [crafted]; Omnicast Boots (11822, -1.34 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.1 spell_power points (12.37 DPS) | yes | Philanthropist's Ring (281635, -5.38 DPS) [quest]; Ring of Forlorn Spirits (2043, -7.20 DPS) [quest]; Reedknot Ring (9622, -7.84 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (8.41 DPS) | yes | Ring of Forlorn Spirits (2043, -3.24 DPS) [quest]; Philanthropist's Ring (281635, -3.66 DPS, sim-verified) [quest]; Reedknot Ring (9622, -3.88 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (767.6 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (767.6 DPS) | yes | Mark of the Chosen (17774, -8.08 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (767.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.05 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -56.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 124.6 spell_power points (80.64 DPS) | yes | Brilliant Wand (249385, -14.51 DPS) [crafted]; Stormrager (16997, -17.43 DPS) [quest]; Greater Eternal Wand (249237, -21.49 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 42.3. Weights run: 2.2s. Verify run: 1.3s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.078 ± 0.003, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.428 per %), hit=0.066 ± 0.003 per rating point (10 rating = 1%, 0.660 per %), spell_haste=not significant (0.012 ± 0.037), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.807 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.53 DPS) | yes | Red Winter Hat (21524, -3.36 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.28 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.61 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (1.02 DPS) | yes | Black Whelp Cloak (7283, -0.34 DPS, sim-verified) [crafted]; Mystic Shawl (3449, -0.51 DPS) [quest]; Pearl-clasped Cloak (5542, -0.51 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.28 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -1.86 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.77 DPS) | yes | Owlbeard Bracers (16981, -0.51 DPS) [quest]; Windsong Bangles (263336, -0.56 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.79 DPS) | yes | Gnoll Casting Gloves (892, -0.40 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.77 DPS) [quest]; Pristine Gloves (253913, -0.77 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.51 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.12 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.26 DPS) [crafted]; Embroidered Leggings (270011, -0.51 DPS) [quest]; Silk-threaded Trousers (1929, -1.94 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.79 DPS) | yes | Red Woolen Boots (4313, -0.92 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.02 DPS) [crafted]; Soft-soled Linen Boots (4312, -1.28 DPS) [crafted] |
| finger1 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.77 DPS) | yes | - |
| finger2 | - | - |  |  |  |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 98 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Seer's Fine Stein (7608) (or Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), Antipodean Rod (2879), Pulsating Hydra Heart (5183), and 7 more) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Tork Wrench (11855, -0.17 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 90.4 spell_power points (23.08 DPS) | yes | Skycaller (12984, -0.81 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.15 DPS) [quest]; Greater Magic Wand (11288, -5.07 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Seer's Fine Stein; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 82.0. Weights run: 2.2s. Verify run: 1.2s. 283 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.018 ± 0.003, crit=0.053 ± 0.001 per rating point (14 rating = 1%, 0.746 per %), hit=0.098 ± 0.004 per rating point (10 rating = 1%, 0.982 per %), spell_haste=-0.425 ± 0.062, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.818 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.96 DPS) | yes | Filigreed Pristine Circlet (253975, -0.81 DPS) [crafted]; Enchanter's Cowl (4322, -1.30 DPS) [crafted]; Silk Headband (7050, -3.90 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.1 spell_power points (0.03 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.01 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.2 spell_power points (2.46 DPS) | yes | Invoker's Mantle (215365, -0.56 DPS) [crafted]; Death Speaker Mantle (6685, -0.80 DPS) [dungeon]; Chestnut Mantle (17695, -1.14 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.34 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Black Whelp Cloak (7283, -0.54 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.50 DPS) | yes | Green Silk Armor (7065, -0.62 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.02 DPS) [dungeon]; Mechbuilder's Overalls (9508, -1.29 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.42 DPS) | yes | Tabitha's Cuffs (251486, -1.78 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -2.14 DPS) [quest]; Windsong Bangles (263336, -2.15 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.88 DPS) | yes | Gnoll Casting Gloves (892, -0.27 DPS) [world]; Jutebraid Gloves (10654, -0.27 DPS, sim-verified) [quest]; Truefaith Gloves (7049, -0.52 DPS) [crafted] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.96 DPS) | yes | Belt of Arugal (6392, -0.52 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.81 DPS) [dungeon]; Invoker's Cord (215366, -1.05 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (82.0 DPS) | yes | Silk-threaded Trousers (1929, -0.03 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.02 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.1 spell_power points (1.92 DPS) | yes | Glinteye Slippers (273024, -0.03 DPS) [dungeon]; Acidic Walkers (9454, -0.53 DPS) [dungeon]; Spidersilk Boots (4320, -2.23 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.1 spell_power points (1.64 DPS) | yes | Solid Gold Band (279845, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Runemender's Seal (273031, -0.83 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.61 DPS) | yes | Solid Gold Band (279845, -0.54 DPS) [quest]; Electrocutioner Lagnut (9447, -0.81 DPS) [dungeon]; Runemender's Seal (273031, -0.81 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (9.68 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -9.63 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -9.64 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 122.2 spell_power points (32.88 DPS) | yes | Wand of Mana Concentration (273043, +0.00 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -2.53 DPS) [crafted]; Unstable Power Core (279847, -3.17 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 196.9. Weights run: 2.0s. Verify run: 1.5s. 381 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.651 ± 0.013, crit=0.134 ± 0.002 per rating point (14 rating = 1%, 1.879 per %), hit=0.173 ± 0.011 per rating point (10 rating = 1%, 1.733 per %), spell_haste=-0.973 ± 0.157, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.908 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.81 DPS) | yes | Augural Shroud (2620, -1.46 DPS) [world]; Papal Fez (9431, -2.05 DPS, sim-verified) [dungeon]; Whitemane's Chapeau (7720, -3.14 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.2 spell_power points (5.12 DPS) | yes | Triune Amulet (7722, -1.83 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -3.21 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -3.48 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.5 spell_power points (6.48 DPS) | yes | Bloodmage Mantle (7684, -0.25 DPS) [dungeon]; Berylline Pads (4197, -0.82 DPS) [quest]; Green Silken Shoulders (7057, -2.24 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 spell_power points (3.88 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.94 DPS) [crafted]; Blackforge Cape (6424, -1.70 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.9 spell_power points (10.86 DPS) | yes | Dreamweave Vest (10021, -0.86 DPS) [crafted]; Robe of Power (7054, -1.72 DPS) [crafted]; Elemental Raiment (9434, -2.06 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.2 spell_power points (3.86 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.92 DPS) [quest]; Windchaser Cuffs (14429, -1.40 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (8.64 DPS) | yes | Black Mageweave Gloves (10003, -2.35 DPS) [crafted]; Red Mageweave Gloves (10018, -3.00 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -3.37 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.09 DPS) [crafted]; Warsong Sash (16975, -0.92 DPS) [quest]; Deathmage Sash (10771, -3.53 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.8 spell_power points (9.14 DPS) | yes | Crimson Silk Pantaloons (7062, -2.29 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -4.03 DPS) [dungeon]; Gaze Dreamer Pants (6903, -4.11 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.06 DPS) | yes | Gilded Slippers (254001, -4.61 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -5.78 DPS) [dungeon]; Spidersilk Boots (4320, -6.04 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (5.56 DPS) | yes | Reedknot Ring (9622, -2.62 DPS) [quest]; Ring of Power Regulation (273030, -3.04 DPS) [dungeon]; Voodoo Band (1996, -3.65 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.3 spell_power points (3.88 DPS) | yes | Reedknot Ring (9622, -0.94 DPS) [quest]; Ring of Power Regulation (273030, -1.36 DPS) [dungeon]; Voodoo Band (1996, -1.97 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.55 DPS) [dungeon]; Spine of the Devourer (273040, -5.54 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (5.87 DPS) | yes | Omega Orb (7749, -2.10 DPS) [quest]; Orb of Mystic Insight (249394, -2.26 DPS, sim-verified) [crafted]; Beacon of Hope (9393, -2.52 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Twisted Nether Wand (249144, -0.33 DPS) [crafted]; Umbral Wand (5216, -1.18 DPS) [dungeon]; Jaina's Firestarter (13064, -2.03 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Earthen Rod

No-known-source sample (15 of 381, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 290.9. Weights run: 2.1s. Verify run: 1.6s. 492 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.166 ± 0.008, crit=0.163 ± 0.002 per rating point (14 rating = 1%, 2.283 per %), hit=0.190 ± 0.010 per rating point (10 rating = 1%, 1.903 per %), spell_haste=-0.726 ± 0.120, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.916 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.34 DPS) | yes | Dreamweave Circlet (10041, -1.99 DPS) [crafted]; Red Mageweave Headband (10033, -2.14 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.74 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.3 spell_power points (3.81 DPS) | yes | Mindburst Medallion (11196, -0.61 DPS) [quest]; Triune Amulet (7722, -1.88 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.75 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.0 spell_power points (7.31 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.92 DPS) [vendor]; Black Mageweave Shoulders (10027, -2.10 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -2.51 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (6.85 DPS) | yes | Deep Woodlands Cloak (19121, -0.69 DPS) [quest]; Runecloth Cloak (13860, -2.13 DPS) [crafted]; Nightfall Drape (12465, -2.74 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (10.51 DPS) | yes | Acumen Robes (17775, -0.31 DPS) [quest]; Elemental Raiment (9434, -0.91 DPS) [world_drop]; Dreamweave Vest (10021, -1.60 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (290.9 DPS) | yes | Condor Bracers (15864, -0.53 DPS) [quest]; Bloodband Bracers (11469, -0.76 DPS) [quest]; Spidertank Oilrag (9448, -3.35 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (8.53 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.91 DPS) [vendor]; Runecloth Gloves (13863, -2.36 DPS) [crafted]; Black Mageweave Gloves (10003, -2.69 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.7 spell_power points (7.16 DPS) | yes | Ghostweave Cord (254073, -0.76 DPS) [crafted]; Star Belt (4329, -1.21 DPS) [crafted]; Gilded Waistcord (254081, -1.90 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (8.69 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.25 DPS) [vendor]; Red Mageweave Pants (10009, -1.38 DPS) [crafted]; Senior Designer's Pantaloons (11841, -2.67 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.97 DPS) | yes | Gilded Sandals (254107, -1.00 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -5.41 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -5.76 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.94 DPS) | yes | Reedknot Ring (9622, -2.74 DPS) [quest]; Snake Hoop (6750, -2.82 DPS) [quest]; Ring of Power Regulation (273030, -3.20 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (4.95 DPS) | yes | Snake Hoop (6750, -1.83 DPS) [quest]; Ring of Power Regulation (273030, -2.21 DPS) [dungeon]; Reedknot Ring (9622, -2.45 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -2.13 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.27 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.96 DPS) [crafted]; Rune of the Guard Captain (19120, -4.20 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.01 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.91 DPS) [dungeon]; Hanzo Sword (8190, -40.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 114.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.09 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 492, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 387.6. Weights run: 2.1s. Verify run: 4.1s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.147 ± 0.009, crit=0.195 ± 0.003 per rating point (14 rating = 1%, 2.730 per %), hit=0.243 ± 0.011 per rating point (10 rating = 1%, 2.426 per %), spell_haste=not significant (0.112 ± 0.117), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.912 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.76 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.30 DPS) [pvp]; Dreamweave Circlet (10041, -1.32 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.5 spell_power points (6.37 DPS) | yes | Glowing Eye of Mordresh (10769, -2.50 DPS) [dungeon]; Triune Amulet (7722, -2.58 DPS) [dungeon]; Mindburst Medallion (11196, -3.11 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (11.82 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.59 DPS, sim-verified) [crafted]; Champion's Dreadweave Spaulders (227092, -3.95 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.9 spell_power points (7.04 DPS) | yes | Deep Woodlands Cloak (19121, -1.35 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.23 DPS) [crafted]; Nightfall Drape (12465, -2.78 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.3 spell_power points (22.37 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.58 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.16 DPS) [pvp]; Robe of the Magi (1716, -11.98 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (387.6 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -8.59 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (387.6 DPS) | yes | Blood Guard's Dreadweave Handwraps (227099, +0.00 DPS) [pvp]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -7.04 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (387.6 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -9.29 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (12.29 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -5.25 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.35 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.95 DPS) [crafted]; Omnicast Boots (11822, -1.06 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.0 spell_power points (9.00 DPS) | yes | Philanthropist's Ring (281635, -3.92 DPS) [quest]; Reedknot Ring (9622, -5.69 DPS) [quest]; Snake Hoop (6750, -5.81 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.15 DPS) | yes | Philanthropist's Ring (281635, -1.51 DPS, sim-verified) [quest]; Reedknot Ring (9622, -2.84 DPS) [quest]; Snake Hoop (6750, -2.96 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (387.6 DPS) | yes | Rune of the Guard Captain (19120, -0.23 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (387.6 DPS) | yes | Rune of the Guard Captain (19120, -4.82 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (387.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.11 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -42.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 168.4 spell_power points (79.60 DPS) | yes | Brilliant Wand (249385, -3.08 DPS, sim-verified) [crafted]; Stormrager (16997, -16.56 DPS) [quest]; Greater Eternal Wand (249237, -22.01 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25520300130201351-0005200000000000000-0550001000000000)

Set DPS (verified): 756.2. Weights run: 2.6s. Verify run: 2.8s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.160 ± 0.016, crit=0.362 ± 0.007 per rating point (14 rating = 1%, 5.068 per %), hit=0.465 ± 0.025 per rating point (10 rating = 1%, 4.654 per %), spell_haste=not significant (0.871 ± 0.304), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.906 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (17.47 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.84 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.6 spell_power points (8.80 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.63 DPS) [dungeon]; Mindburst Medallion (11196, -4.30 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (16.18 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.55 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -3.78 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (9.68 DPS) | yes | Deep Woodlands Cloak (19121, -0.98 DPS) [quest]; Runecloth Cloak (13860, -3.03 DPS) [crafted]; Nightfall Drape (12465, -3.86 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.4 spell_power points (30.70 DPS) | yes | Warlord's Dreadweave Robe (231591, -7.50 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.45 DPS) [pvp]; Robe of the Magi (1716, -14.01 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.6 spell_power points (8.16 DPS) | yes | Spidertank Oilrag (9448, -2.37 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -2.90 DPS) [crafted]; Earthenweave Cuffs (254131, -2.98 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (14.24 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS, sim-verified) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.23 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (10.10 DPS) | yes | Ghostweave Cord (254073, -1.04 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.53 DPS) [crafted]; Star Belt (4329, -1.69 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (16.82 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -5.74 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.53 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.29 DPS) [crafted]; Omnicast Boots (11822, -1.34 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.1 spell_power points (12.37 DPS) | yes | Philanthropist's Ring (281635, -5.38 DPS) [quest]; Reedknot Ring (9622, -7.84 DPS) [quest]; Snake Hoop (6750, -7.97 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (8.41 DPS) | yes | Philanthropist's Ring (281635, -1.42 DPS) [quest]; Reedknot Ring (9622, -3.88 DPS) [quest]; Snake Hoop (6750, -4.01 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (756.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (756.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (756.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.05 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -56.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 124.6 spell_power points (80.64 DPS) | yes | Brilliant Wand (249385, -14.51 DPS) [crafted]; Stormrager (16997, -17.43 DPS) [quest]; Greater Eternal Wand (249237, -21.49 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

