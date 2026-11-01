# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 45.4. Weights run: 2.2s. Verify run: 1.3s. 161 eligible items had no known source.

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

No-known-source sample (15 of 161, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 82.8. Weights run: 2.2s. Verify run: 1.2s. 298 eligible items had no known source.

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

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 128.3. Weights run: 1.8s. Verify run: 1.3s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.188 ± 0.006, crit=0.090 ± 0.002 per rating point (14 rating = 1%, 1.257 per %), hit=0.129 ± 0.006 per rating point (10 rating = 1%, 1.293 per %), spell_haste=0.387 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.842 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.06 DPS) | yes | Augural Shroud (2620, -2.35 DPS) [world]; Living Cowl (5608, -2.49 DPS, sim-verified) [world]; Papal Fez (9431, -2.83 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.5 spell_power points (2.45 DPS) | yes | Triune Amulet (7722, -1.37 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.07 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.13 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (3.09 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Inquisitor's Shawl (19507, -0.36 DPS) [dungeon]; Berylline Pads (4197, -0.52 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.02 DPS) | yes | Guardian Cloak (5965, -0.02 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.58 DPS) [quest]; Long Silken Cloak (4326, -0.61 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (6.68 DPS) | yes | Dreamweave Vest (10021, -0.99 DPS) [crafted]; Elemental Raiment (9434, -1.10 DPS, sim-verified) [world_drop]; Robe of Power (7054, -1.98 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.60 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.58 DPS) [quest]; Earthen Silk Cuffs (254019, -1.44 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.41 DPS) | yes | Black Mageweave Gloves (10003, -1.48 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.70 DPS) [crafted]; Gilded Handwraps (254021, -2.72 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.75 DPS) | yes | Belt of Arugal (6392, -0.99 DPS) [dungeon]; Gilded Cord (254037, -1.01 DPS) [crafted]; Deathmage Sash (10771, -1.51 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.3 spell_power points (4.69 DPS) | yes | Gaze Dreamer Pants (6903, -1.13 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.68 DPS) [crafted]; Stoneweaver Leggings (9407, -2.24 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.93 DPS) | yes | Gilded Slippers (254001, -3.83 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.69 DPS) [crafted]; Glinteye Slippers (273024, -4.91 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.9 spell_power points (3.16 DPS) | yes | Reedknot Ring (9622, -1.14 DPS) [quest]; Snake Hoop (6750, -1.15 DPS) [quest]; Ring of Power Regulation (273030, -1.43 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (2.31 DPS) | yes | Reedknot Ring (9622, -0.29 DPS) [quest]; Snake Hoop (6750, -0.31 DPS) [quest]; Ring of Power Regulation (273030, -0.58 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (128.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.11 DPS) [dungeon]; Spine of the Devourer (273040, -2.74 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.04 DPS) | yes | Beacon of Hope (9393, -1.73 DPS) [dungeon]; Orb of Lorica (11262, -1.77 DPS) [quest]; Orb of Mystic Insight (249394, -2.21 DPS, sim-verified) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 137.5 spell_power points (39.70 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -3.97 DPS) [crafted]; Umbral Wand (5216, -4.03 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 207.2. Weights run: 1.9s. Verify run: 1.5s. 507 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.295 ± 0.008, crit=0.137 ± 0.002 per rating point (14 rating = 1%, 1.915 per %), hit=0.179 ± 0.009 per rating point (10 rating = 1%, 1.791 per %), spell_haste=not significant (0.225 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.856 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.38 DPS) | yes | Dreamweave Circlet (10041, -0.95 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.86 DPS) [crafted]; Red Mageweave Headband (10033, -1.86 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.4 spell_power points (2.90 DPS) | yes | Mindburst Medallion (11196, -0.49 DPS) [quest]; Triune Amulet (7722, -1.01 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.62 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.3 spell_power points (5.68 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -1.78 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.87 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -2.07 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.8 spell_power points (4.89 DPS) | yes | Runecloth Cloak (13860, -1.81 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.10 DPS) [dungeon]; Big Voodoo Cloak (8216, -2.52 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.9 spell_power points (7.73 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.21 DPS) [world_drop]; Dreamweave Vest (10021, -1.32 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 9.1 spell_power points (2.81 DPS) | yes | Spidertank Oilrag (9448, -0.02 DPS) [dungeon]; Bloodband Bracers (11469, -0.44 DPS) [quest]; Arcane Runed Bracers (4744, -2.88 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (5.95 DPS) | yes | Black Mageweave Gloves (10003, -1.30 DPS) [crafted]; Runecloth Gloves (13863, -1.40 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.66 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.0 spell_power points (5.26 DPS) | yes | Ghostweave Cord (254073, -0.92 DPS) [crafted]; Star Belt (4329, -1.23 DPS) [crafted]; Gilded Waistcord (254081, -1.33 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.89 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.48 DPS) [vendor]; Red Mageweave Pants (10009, -1.17 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -1.53 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.45 DPS) | yes | Gilded Sandals (254107, -1.06 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.39 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.58 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.03 DPS) | yes | Ring of Forlorn Spirits (2043, -1.55 DPS) [quest]; Snake Hoop (6750, -1.71 DPS) [quest]; Reedknot Ring (9622, -1.86 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 spell_power points (3.56 DPS) | yes | Ring of Forlorn Spirits (2043, -0.57 DPS, sim-verified) [quest]; Snake Hoop (6750, -1.24 DPS) [quest]; Reedknot Ring (9622, -1.39 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (207.2 DPS) | yes | Frozen Heart of the Mountain (249469, -1.96 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (207.2 DPS) | yes | Frozen Heart of the Mountain (249469, -1.36 DPS) [crafted]; Mark of the Chosen (17774, -4.01 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (207.2 DPS) | yes | Spire of Hakkar (10844, -0.48 DPS) [world]; Spellshifter Rod (9527, -0.85 DPS) [quest]; Hanzo Sword (8190, -27.48 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 169.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.27 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 307.0. Weights run: 2.0s. Verify run: 3.9s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.130 ± 0.010, crit=0.192 ± 0.003 per rating point (14 rating = 1%, 2.684 per %), hit=0.296 ± 0.013 per rating point (10 rating = 1%, 2.965 per %), spell_haste=not significant (0.038 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.856 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.77 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.32 DPS) [pvp]; Dreamweave Circlet (10041, -1.71 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.3 spell_power points (4.32 DPS) | yes | Glowing Eye of Mordresh (10769, -1.21 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -1.72 DPS) [dungeon]; Mindburst Medallion (11196, -2.12 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (8.12 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.69 DPS, sim-verified) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.80 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.8 spell_power points (4.80 DPS) | yes | Nightfall Drape (12465, -1.88 DPS) [dungeon]; Runecloth Cloak (13860, -2.52 DPS, sim-verified) [crafted]; Icy Cloak (4327, -2.53 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.2 spell_power points (15.32 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.92 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -6.36 DPS) [pvp]; Robe of the Magi (1716, -9.65 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (307.0 DPS) | yes | Arcane Runed Bracers (4744, +0.00 DPS) [quest]; Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Runecloth Cuffs (254123, -6.54 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (307.0 DPS) | yes | Knight-Lieutenant's Dreadweave Handwraps (227100, +0.00 DPS) [pvp]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -4.95 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (307.0 DPS) | yes | Star Belt (4329, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -5.80 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (8.45 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.35 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.80 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.65 DPS) [crafted]; Omnicast Boots (11822, -0.79 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.9 spell_power points (6.14 DPS) | yes | Philanthropist's Ring (281635, -2.68 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.54 DPS) [quest]; Reedknot Ring (9622, -3.87 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.22 DPS) | yes | Ring of Forlorn Spirits (2043, -1.62 DPS) [quest]; Philanthropist's Ring (281635, -1.71 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.95 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (307.0 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (307.0 DPS) | yes | Mark of the Chosen (17774, -4.42 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (307.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.14 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -32.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 242.3 spell_power points (78.71 DPS) | yes | Stormrager (16997, -5.24 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.91 DPS) [crafted]; Torch of Austen (13004, -21.92 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25520100135001050-0005200000000000000-0550004000000000)

Set DPS (verified): 649.0. Weights run: 2.3s. Verify run: 3.0s. 1132 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.167 ± 0.020, crit=0.432 ± 0.010 per rating point (14 rating = 1%, 6.045 per %), hit=0.518 ± 0.032 per rating point (10 rating = 1%, 5.178 per %), spell_haste=not significant (0.457 ± 0.414), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.860 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.64 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.96 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (6.91 DPS) | yes | Glowing Eye of Mordresh (10769, -2.65 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -2.86 DPS) [dungeon]; Mindburst Medallion (11196, -3.37 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.2 spell_power points (12.74 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.11 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.53 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (7.58 DPS) | yes | Nightfall Drape (12465, -3.03 DPS) [dungeon]; Runecloth Cloak (13860, -3.26 DPS, sim-verified) [crafted]; Icy Cloak (4327, -4.04 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.5 spell_power points (24.00 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.81 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.68 DPS) [pvp]; Robe of the Magi (1716, -12.99 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (6.40 DPS) | yes | Arcane Runed Bracers (4744, -1.85 DPS) [quest]; Spidertank Oilrag (9448, -1.85 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.27 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.8 spell_power points (11.50 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.55 DPS) [pvp]; Earth Warder's Gloves (21318, -3.91 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.7 spell_power points (7.92 DPS) | yes | Ghostweave Cord (254073, -0.84 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.18 DPS) [crafted]; Star Belt (4329, -1.35 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (13.13 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -3.01 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.12 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.01 DPS) [dungeon]; Venomspew Footpads (275606, -1.01 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (9.68 DPS) | yes | Philanthropist's Ring (281635, -4.21 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.64 DPS) [quest]; Reedknot Ring (9622, -6.15 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.57 DPS) | yes | Philanthropist's Ring (281635, -1.09 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.53 DPS) [quest]; Reedknot Ring (9622, -3.03 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (649.0 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (649.0 DPS) | yes | Mark of the Chosen (17774, -6.52 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (649.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.00 DPS) [world]; Teebu's Blazing Longsword (1728, -42.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 157.9 spell_power points (79.79 DPS) | yes | Brilliant Wand (249385, -15.07 DPS) [crafted]; Stormrager (16997, -16.68 DPS) [quest]; Greater Eternal Wand (249237, -21.91 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; ranged: Torch of Light

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 42.3. Weights run: 2.2s. Verify run: 1.3s. 155 eligible items had no known source.

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

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 25522300110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 82.0. Weights run: 2.2s. Verify run: 1.2s. 289 eligible items had no known source.

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

No-known-source sample (15 of 289, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 127.5. Weights run: 1.8s. Verify run: 1.4s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.188 ± 0.006, crit=0.090 ± 0.002 per rating point (14 rating = 1%, 1.257 per %), hit=0.129 ± 0.006 per rating point (10 rating = 1%, 1.293 per %), spell_haste=0.387 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.842 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.06 DPS) | yes | Augural Shroud (2620, -2.35 DPS) [world]; Living Cowl (5608, -2.45 DPS, sim-verified) [world]; Papal Fez (9431, -2.83 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.5 spell_power points (2.45 DPS) | yes | Triune Amulet (7722, -0.74 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.07 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.13 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (3.09 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Inquisitor's Shawl (19507, -0.36 DPS) [dungeon]; Berylline Pads (4197, -0.52 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.02 DPS) | yes | Long Silken Cloak (4326, -0.02 DPS) [crafted]; Guardian Cloak (5965, -0.02 DPS) [crafted]; Hillman's Cloak (3719, -0.58 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (6.68 DPS) | yes | Elemental Raiment (9434, -0.40 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.99 DPS) [crafted]; Robe of Power (7054, -1.98 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.60 DPS) | yes | Condor Bracers (15864, -0.51 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.01 DPS) [quest]; Earthen Silk Cuffs (254019, -1.44 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.41 DPS) | yes | Black Mageweave Gloves (10003, -1.09 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.70 DPS) [crafted]; Gilded Handwraps (254021, -2.72 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.75 DPS) | yes | Warsong Sash (16975, -0.65 DPS, sim-verified) [quest]; Deathmage Sash (10771, -0.92 DPS) [dungeon]; Belt of Arugal (6392, -0.99 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.3 spell_power points (4.69 DPS) | yes | Gaze Dreamer Pants (6903, -1.23 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.68 DPS) [crafted]; Stoneweaver Leggings (9407, -2.24 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.93 DPS) | yes | Gilded Slippers (254001, -2.87 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.69 DPS) [crafted]; Glinteye Slippers (273024, -4.91 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.9 spell_power points (3.16 DPS) | yes | Snake Hoop (6750, -1.15 DPS) [quest]; Ring of Power Regulation (273030, -1.43 DPS) [dungeon]; Solid Gold Band (279845, -2.00 DPS) [quest] |
| finger2 | Reedknot Ring (9622) | Jarl Needs a Blade [quest] | 7.0 spell_power points (2.02 DPS) | yes | Snake Hoop (6750, -0.02 DPS) [quest]; Ring of Power Regulation (273030, -0.29 DPS) [dungeon]; Solid Gold Band (279845, -0.87 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.11 DPS) [dungeon]; Spine of the Devourer (273040, -2.74 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (4.04 DPS) | yes | Omega Orb (7749, -1.14 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -1.70 DPS) [crafted]; Beacon of Hope (9393, -1.73 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | sim-verified (127.5 DPS) | yes | Twisted Nether Wand (249144, -0.72 DPS) [crafted]; Umbral Wand (5216, -0.79 DPS) [dungeon]; Jaina's Firestarter (13064, -1.37 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Reedknot Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Earthen Rod

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 202.5. Weights run: 1.9s. Verify run: 1.5s. 498 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.295 ± 0.008, crit=0.137 ± 0.002 per rating point (14 rating = 1%, 1.915 per %), hit=0.179 ± 0.009 per rating point (10 rating = 1%, 1.791 per %), spell_haste=not significant (0.225 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.856 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.38 DPS) | yes | Dreamweave Circlet (10041, -0.95 DPS) [crafted]; Red Mageweave Headband (10033, -1.42 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.86 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.4 spell_power points (2.90 DPS) | yes | Mindburst Medallion (11196, -0.49 DPS) [quest]; Triune Amulet (7722, -1.33 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.62 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.3 spell_power points (5.68 DPS) | yes | Black Mageweave Shoulders (10027, -1.74 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.78 DPS) [vendor]; Bloodmage Mantle (7684, -2.07 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.8 spell_power points (4.89 DPS) | yes | Deep Woodlands Cloak (19121, -1.01 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.37 DPS) [crafted]; Nightfall Drape (12465, -2.10 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.9 spell_power points (7.73 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.21 DPS) [world_drop]; Dreamweave Vest (10021, -1.32 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 9.1 spell_power points (2.81 DPS) | yes | Bloodband Bracers (11469, -0.44 DPS) [quest]; Condor Bracers (15864, -0.64 DPS) [quest]; Spidertank Oilrag (9448, -2.56 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (5.95 DPS) | yes | Black Mageweave Gloves (10003, -1.30 DPS) [crafted]; Runecloth Gloves (13863, -1.40 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.70 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.0 spell_power points (5.26 DPS) | yes | Ghostweave Cord (254073, -0.92 DPS) [crafted]; Star Belt (4329, -1.23 DPS) [crafted]; Gilded Waistcord (254081, -1.33 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (5.89 DPS) | yes | Red Mageweave Pants (10009, -0.45 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -0.48 DPS) [vendor]; Senior Designer's Pantaloons (11841, -1.53 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.45 DPS) | yes | Gilded Sandals (254107, -0.85 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.39 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.58 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.03 DPS) | yes | Snake Hoop (6750, -1.71 DPS) [quest]; Reedknot Ring (9622, -1.86 DPS) [quest]; Ring of Power Regulation (273030, -2.17 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 spell_power points (3.56 DPS) | yes | Snake Hoop (6750, -1.32 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.39 DPS) [quest]; Ring of Power Regulation (273030, -1.70 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (202.5 DPS) | yes | Rune of the Guard Captain (19120, -1.47 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (202.5 DPS) | yes | Rune of the Guard Captain (19120, -1.10 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (202.5 DPS) | yes | Spire of Hakkar (10844, -0.48 DPS) [world]; Spellshifter Rod (9527, -0.85 DPS) [quest]; Hanzo Sword (8190, -26.44 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 169.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.27 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 498, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 302.1. Weights run: 2.0s. Verify run: 4.1s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.130 ± 0.010, crit=0.192 ± 0.003 per rating point (14 rating = 1%, 2.684 per %), hit=0.296 ± 0.013 per rating point (10 rating = 1%, 2.965 per %), spell_haste=not significant (0.038 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.856 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreadmist Mask (16698) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -9.45 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.3 spell_power points (4.32 DPS) | yes | Glowing Eye of Mordresh (10769, -1.71 DPS) [dungeon]; Triune Amulet (7722, -1.72 DPS) [dungeon]; Mindburst Medallion (11196, -2.12 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (8.12 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.18 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -2.80 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.8 spell_power points (4.80 DPS) | yes | Deep Woodlands Cloak (19121, -1.08 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.54 DPS) [crafted]; Nightfall Drape (12465, -1.88 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.2 spell_power points (15.32 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.92 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -6.36 DPS) [pvp]; Robe of the Magi (1716, -8.74 DPS, sim-verified) [world_drop] |
| wrist | Dreadmist Bracers (16703) | Blackrock Spire: Summoned Blackhand Dreadweaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -4.48 DPS, sim-verified) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (7.15 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.16 DPS) [pvp] |
| waist | Dreadmist Belt (16702) | Scholomance: Scholomance Necromancer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star Belt (4329, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -5.94 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (8.45 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.39 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.80 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.65 DPS) [crafted]; Omnicast Boots (11822, -0.79 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.9 spell_power points (6.14 DPS) | yes | Philanthropist's Ring (281635, -2.68 DPS) [quest]; Reedknot Ring (9622, -3.87 DPS) [quest]; Snake Hoop (6750, -3.98 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.22 DPS) | yes | Philanthropist's Ring (281635, -0.76 DPS) [quest]; Reedknot Ring (9622, -1.95 DPS) [quest]; Snake Hoop (6750, -2.06 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (-1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.19 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (302.1 DPS) | yes | Rune of the Guard Captain (19120, -3.02 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.14 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -30.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 242.3 spell_power points (78.71 DPS) | yes | Stormrager (16997, -4.97 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.91 DPS) [crafted]; Torch of Austen (13004, -21.92 DPS) [world_drop] |

**New at 60:** head: Dreadmist Mask; neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Dreadmist Bracers; hands: Earth Warder's Gloves; waist: Dreadmist Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 25520100135001050-0005200000000000000-0550004000000000)

Set DPS (verified): 646.8. Weights run: 2.3s. Verify run: 4.5s. 1124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.167 ± 0.020, crit=0.432 ± 0.010 per rating point (14 rating = 1%, 6.045 per %), hit=0.518 ± 0.032 per rating point (10 rating = 1%, 5.178 per %), spell_haste=not significant (0.457 ± 0.414), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.860 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.64 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -3.04 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (6.91 DPS) | yes | Glowing Eye of Mordresh (10769, -2.69 DPS) [dungeon]; Triune Amulet (7722, -2.86 DPS) [dungeon]; Mindburst Medallion (11196, -3.37 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.2 spell_power points (12.74 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -2.53 DPS) [pvp]; Argent Shoulders (19059, -3.64 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (7.58 DPS) | yes | Runecloth Cloak (13860, -2.36 DPS) [crafted]; Nightfall Drape (12465, -3.03 DPS) [dungeon]; Deep Woodlands Cloak (19121, -3.34 DPS, sim-verified) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.5 spell_power points (24.00 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.81 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.68 DPS) [pvp]; Robe of the Magi (1716, -16.17 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (646.8 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -11.91 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.8 spell_power points (11.50 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.55 DPS) [pvp]; Earth Warder's Gloves (21318, -8.13 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (646.8 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -11.63 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (646.8 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -9.53 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.12 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.01 DPS) [dungeon]; Venomspew Footpads (275606, -1.01 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (9.68 DPS) | yes | Philanthropist's Ring (281635, -4.21 DPS) [quest]; Reedknot Ring (9622, -6.15 DPS) [quest]; Snake Hoop (6750, -6.23 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.57 DPS) | yes | Philanthropist's Ring (281635, -2.79 DPS, sim-verified) [quest]; Reedknot Ring (9622, -3.03 DPS) [quest]; Snake Hoop (6750, -3.11 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (646.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (646.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (646.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.00 DPS) [world]; Teebu's Blazing Longsword (1728, -46.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 157.9 spell_power points (79.79 DPS) | yes | Brilliant Wand (249385, -5.68 DPS, sim-verified) [crafted]; Stormrager (16997, -16.68 DPS) [quest]; Greater Eternal Wand (249237, -21.91 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Void; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Devout Skirt; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

