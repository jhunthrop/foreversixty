# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 33.9. Weights run: 1.5s. Verify run: 0.9s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.112 ± 0.004, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.430 per %), hit=0.081 ± 0.004 per rating point (10 rating = 1%, 0.806 per %), spell_haste=not significant (-0.123 ± 0.063), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.753 ± 0.002, fire_power=0.247 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.11 DPS) | yes | Flame Circlet (253953, -2.01 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.92 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.18 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.74 DPS) | yes | Pearl-clasped Cloak (5542, -0.37 DPS) [crafted]; Black Whelp Cloak (7283, -0.41 DPS, sim-verified) [crafted]; Reinforced Linen Cape (2580, -0.55 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.92 DPS) | yes | Green Woolen Vest (2582, -0.18 DPS) [crafted]; Bloody Apron (6226, -0.18 DPS) [dungeon]; Gray Woolen Robe (2585, -1.38 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.18 DPS) | yes | Ivycloth Bracelets (9793, -0.21 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.29 DPS) | yes | Gnoll Casting Gloves (892, -0.31 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.55 DPS) [crafted]; Phoenix Gloves (4331, -0.88 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.37 DPS) [crafted]; Flame Sash (253929, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.76 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.18 DPS) [crafted]; Heavy Woolen Pants (4316, -0.55 DPS) [crafted]; Silk-threaded Trousers (1929, -1.41 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.29 DPS) | yes | Red Woolen Boots (4313, -0.40 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.74 DPS) [crafted]; Soft-soled Linen Boots (4312, -0.92 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.92 DPS) | yes | Ring of the Shadow (1462, -1.00 DPS, sim-verified) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.55 DPS) | yes | Ring of the Shadow (1462, -0.89 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 89 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -0.95 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.92 DPS) | yes | Antipodean Rod (2879, -0.85 DPS, sim-verified) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 124.0 spell_power points (22.86 DPS) | yes | Skycaller (12984, -0.63 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.29 DPS) [quest]; Deepblaze (279896, -4.33 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 64.0. Weights run: 1.5s. Verify run: 0.8s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.062 ± 0.005, crit=0.046 ± 0.001 per rating point (14 rating = 1%, 0.639 per %), hit=0.121 ± 0.005 per rating point (10 rating = 1%, 1.208 per %), spell_haste=0.309 ± 0.064, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.002, fire_power=0.235 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.31 DPS) | yes | Filigreed Pristine Circlet (253975, -0.63 DPS) [crafted]; Enchanter's Cowl (4322, -0.92 DPS) [crafted]; Silk Headband (7050, -3.40 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.4 spell_power points (0.08 DPS) | yes | Crystal Starfire Medallion (5003, -0.03 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.03 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.6 spell_power points (2.01 DPS) | yes | Invoker's Mantle (215365, -0.45 DPS, sim-verified) [crafted]; Moonlit Amice (11884, -0.54 DPS) [quest]; Death Speaker Mantle (6685, -0.60 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.05 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted]; Prelacy Cape (7004, -0.21 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.73 DPS) | yes | Beguiler Robes (7728, -0.68 DPS) [dungeon]; Green Silk Armor (7065, -0.75 DPS, sim-verified) [crafted]; Mechbuilder's Overalls (9508, -0.89 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.89 DPS) | yes | Windsong Bangles (263336, -1.68 DPS) [quest]; Phoenix Bindings (210781, -1.73 DPS, sim-verified) [crafted]; Glowing Magical Bracelets (13106, -1.79 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.47 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.21 DPS) [world]; Truefaith Gloves (7049, -0.38 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.2 spell_power points (1.93 DPS) | yes | Invoker's Cord (215366, -0.39 DPS) [crafted]; Crimson Silk Belt (7055, -0.58 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.91 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.52 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.05 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -1.18 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.4 spell_power points (1.56 DPS) | yes | Glinteye Slippers (273024, -0.09 DPS) [dungeon]; Nimbus Boots (6998, -0.30 DPS) [quest]; Spidersilk Boots (4320, -2.07 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.3 spell_power points (1.33 DPS) | yes | Minor Channeling Ring (1449, -0.25 DPS) [quest]; Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon]; Runemender's Seal (273031, -0.70 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.26 DPS) | yes | Minor Channeling Ring (1449, -0.18 DPS) [quest]; Electrocutioner Lagnut (9447, -0.63 DPS) [dungeon]; Runemender's Seal (273031, -0.63 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (7.56 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.36 DPS) [quest]; Staff of Soran'ruk (15109, -5.88 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 156.7 spell_power points (32.91 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.86 DPS) [crafted]; Scorching Wand (5213, -3.77 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 120.2. Weights run: 1.4s. Verify run: 1.0s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.236 ± 0.007, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.924 per %), hit=0.179 ± 0.008 per rating point (10 rating = 1%, 1.788 per %), spell_haste=0.890 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.796 ± 0.002, fire_power=0.204 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.22 DPS) | yes | Augural Shroud (2620, -1.57 DPS, sim-verified) [world]; Living Cowl (5608, -1.99 DPS) [world]; Papal Fez (9431, -2.23 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.9 spell_power points (2.21 DPS) | yes | Triune Amulet (7722, -0.22 DPS) [dungeon]; Necklace of Calisea (1714, -1.80 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.86 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 spell_power points (2.76 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.44 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 spell_power points (1.78 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.04 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.54 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.4 spell_power points (5.82 DPS) | yes | Elemental Raiment (9434, -0.60 DPS) [world_drop]; Dreamweave Vest (10021, -0.82 DPS) [crafted]; Robe of Power (7054, -1.64 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.24 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.50 DPS) [quest]; Earthen Silk Cuffs (254019, -1.24 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (4.71 DPS) | yes | Black Mageweave Gloves (10003, -0.91 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.39 DPS) [crafted]; Gilded Handwraps (254021, -2.31 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.23 DPS) | yes | Gilded Cord (254037, -0.77 DPS) [crafted]; Belt of Arugal (6392, -0.82 DPS) [dungeon]; Deathmage Sash (10771, -0.95 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.8 spell_power points (4.18 DPS) | yes | Gaze Dreamer Pants (6903, -1.20 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.43 DPS) [crafted]; Stoneweaver Leggings (9407, -1.97 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.96 DPS) | yes | Gilded Slippers (254001, -2.75 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.99 DPS) [crafted]; Glinteye Slippers (273024, -4.22 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.2 spell_power points (2.78 DPS) | yes | Snake Hoop (6750, -0.99 DPS) [quest]; Reedknot Ring (9622, -1.04 DPS) [quest]; Ring of Power Regulation (273030, -1.29 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (1.99 DPS) | yes | Snake Hoop (6750, -0.20 DPS) [quest]; Reedknot Ring (9622, -0.25 DPS) [quest]; Ring of Power Regulation (273030, -0.50 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.12 DPS) [dungeon]; Spine of the Devourer (273040, -2.46 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.48 DPS) | yes | Orb of Lorica (11262, -1.40 DPS) [quest]; Beacon of Hope (9393, -1.49 DPS) [dungeon]; Orb of Mystic Insight (249394, -1.63 DPS, sim-verified) [crafted] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | sim-verified (120.2 DPS) | yes | Umbral Wand (5216, -0.67 DPS) [dungeon]; Twisted Nether Wand (249144, -0.84 DPS) [crafted]; Jaina's Firestarter (13064, -1.48 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Earthen Rod

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 179.1. Weights run: 1.4s. Verify run: 1.0s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.233 ± 0.019, crit=0.115 ± 0.004 per rating point (14 rating = 1%, 1.605 per %), hit=0.354 ± 0.019 per rating point (10 rating = 1%, 3.540 per %), spell_haste=-1.487 ± 0.300, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.509 ± 0.001, fire_power=0.491 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.54 DPS) | yes | Dreamweave Circlet (10041, -0.89 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.45 DPS) [crafted]; Red Mageweave Headband (10033, -1.91 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.9 spell_power points (2.15 DPS) | yes | Mindburst Medallion (11196, -0.35 DPS) [quest]; Triune Amulet (7722, -0.95 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.36 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.2 spell_power points (4.16 DPS) | yes | Black Mageweave Shoulders (10027, -1.23 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.33 DPS) [vendor]; Bloodmage Mantle (7684, -1.48 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.4 spell_power points (3.73 DPS) | yes | Nightfall Drape (12465, -1.55 DPS) [dungeon]; Cindercloth Cloak (14044, -1.73 DPS) [crafted]; Runecloth Cloak (13860, -1.86 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.7 spell_power points (5.73 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.64 DPS) [world_drop]; Dreamweave Vest (10021, -0.86 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.18 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.09 DPS) [crafted]; Netherflame Cuffs (254065, -0.24 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (4.58 DPS) | yes | Black Mageweave Gloves (10003, -0.95 DPS) [crafted]; Runecloth Gloves (13863, -1.17 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.23 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.3 spell_power points (3.95 DPS) | yes | Ghostweave Cord (254073, -0.56 DPS) [crafted]; Star Belt (4329, -0.81 DPS) [crafted]; Gilded Waistcord (254081, -1.02 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (4.60 DPS) | yes | Red Mageweave Pants (10009, -0.53 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -0.63 DPS) [vendor]; Senior Designer's Pantaloons (11841, -1.30 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.81 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -2.63 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -2.64 DPS) [crafted]; Black Mageweave Boots (10026, -2.75 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.15 DPS) | yes | Ring of Forlorn Spirits (2043, -1.21 DPS) [quest]; Snake Hoop (6750, -1.41 DPS) [quest]; Reedknot Ring (9622, -1.45 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.2 spell_power points (2.70 DPS) | yes | Ring of Forlorn Spirits (2043, -0.77 DPS) [quest]; Snake Hoop (6750, -0.97 DPS) [quest]; Reedknot Ring (9622, -1.01 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (179.1 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (179.1 DPS) | yes | Mark of the Chosen (17774, -2.99 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (179.1 DPS) | yes | Spire of Hakkar (10844, -0.19 DPS) [world]; Inventor's Focal Sword (17719, -0.68 DPS) [dungeon]; Hanzo Sword (8190, -20.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 216.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.29 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.81 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 353.5. Weights run: 1.4s. Verify run: 1.9s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.273 ± 0.029, crit=0.145 ± 0.005 per rating point (14 rating = 1%, 2.025 per %), hit=0.506 ± 0.030 per rating point (10 rating = 1%, 5.057 per %), spell_haste=1.814 ± 0.394, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.627 ± 0.001, fire_power=0.373 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.20 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -2.00 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.7 spell_power points (5.57 DPS) | yes | Glowing Eye of Mordresh (10769, -2.10 DPS) [dungeon]; Triune Amulet (7722, -2.54 DPS) [dungeon]; Mindburst Medallion (11196, -2.68 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.44 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.11 DPS, sim-verified) [crafted]; Rotgrip Mantle (17732, -2.67 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (5.91 DPS) | yes | Runecloth Cloak (13860, -2.25 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.51 DPS) [dungeon]; Mageflame Cloak (13007, -2.95 DPS) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.5 spell_power points (18.31 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.74 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -6.80 DPS) [pvp]; Acumen Robes (17775, -13.17 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.7 spell_power points (5.19 DPS) | yes | Arcane Runed Bracers (4744, -1.79 DPS) [quest]; Spidertank Oilrag (9448, -1.79 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.82 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.31 DPS) | yes | Knight-Lieutenant's Dreadweave Handwraps (227100, +0.00 DPS) [pvp]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Dreamweave Gloves (10019, -2.01 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.7 spell_power points (6.32 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -1.03 DPS) [crafted]; Star Belt (4329, -1.41 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.82 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -2.49 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.07 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.34 DPS) [pvp]; Omnicast Boots (11822, -2.01 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.9 spell_power points (7.52 DPS) | yes | Philanthropist's Ring (281635, -3.23 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.50 DPS) [quest]; Snake Hoop (6750, -4.74 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.91 DPS) | yes | Philanthropist's Ring (281635, -0.62 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.89 DPS) [quest]; Snake Hoop (6750, -2.13 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (353.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (353.5 DPS) | yes | Mark of the Chosen (17774, -4.69 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (353.5 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.48 DPS) [world]; Teebu's Blazing Longsword (1728, -32.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 209.2 spell_power points (79.03 DPS) | yes | Brilliant Wand (249385, -15.32 DPS) [crafted]; Stormrager (16997, -15.82 DPS) [quest]; Torch of Austen (13004, -22.24 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 05000000000000000-2035113011220001351-0550005000000000)

Set DPS (verified): 741.5. Weights run: 1.7s. Verify run: 2.6s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.266 ± 0.043, crit=0.374 ± 0.013 per rating point (14 rating = 1%, 5.242 per %), hit=0.846 ± 0.055 per rating point (10 rating = 1%, 8.458 per %), spell_haste=3.791 ± 0.798, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.629 ± 0.002, fire_power=0.371 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (14.88 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -4.73 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.7 spell_power points (8.08 DPS) | yes | Triune Amulet (7722, -3.67 DPS) [dungeon]; Mindburst Medallion (11196, -3.89 DPS) [quest]; Glowing Eye of Mordresh (10769, -4.35 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.7 spell_power points (14.16 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.38 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.75 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (8.59 DPS) | yes | Runecloth Cloak (13860, -2.46 DPS) [crafted]; Nightfall Drape (12465, -3.63 DPS) [dungeon]; Mageflame Cloak (13007, -4.30 DPS) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.4 spell_power points (26.66 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.52 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.96 DPS) [pvp]; Acumen Robes (17775, -16.55 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.7 spell_power points (7.53 DPS) | yes | Arcane Runed Bracers (4744, -2.57 DPS) [quest]; Spidertank Oilrag (9448, -2.57 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.64 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.1 spell_power points (12.20 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.04 DPS) [pvp]; Earth Warder's Gloves (21318, -4.81 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.7 spell_power points (9.18 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -1.46 DPS) [crafted]; Star Belt (4329, -2.02 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (14.33 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -4.21 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.22 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.45 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.56 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.9 spell_power points (10.94 DPS) | yes | Philanthropist's Ring (281635, -4.70 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.54 DPS) [quest]; Snake Hoop (6750, -6.90 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.16 DPS) | yes | Philanthropist's Ring (281635, -0.92 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.75 DPS) [quest]; Snake Hoop (6750, -3.12 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (741.5 DPS) | yes | Uther's Strength (11302, -0.89 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (741.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (741.5 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.66 DPS) [world]; Teebu's Blazing Longsword (1728, -50.14 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 145.3 spell_power points (80.07 DPS) | yes | Brilliant Wand (249385, -14.50 DPS) [crafted]; Stormrager (16997, -16.64 DPS) [quest]; Greater Eternal Wand (249237, -21.78 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; ranged: Torch of Light

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 31.4. Weights run: 1.5s. Verify run: 0.9s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.112 ± 0.004, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.430 per %), hit=0.081 ± 0.004 per rating point (10 rating = 1%, 0.806 per %), spell_haste=not significant (-0.123 ± 0.063), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.753 ± 0.002, fire_power=0.247 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.11 DPS) | yes | Flame Circlet (253953, -1.91 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (0.92 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.18 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.76 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.74 DPS) | yes | Mystic Shawl (3449, -0.37 DPS) [quest]; Pearl-clasped Cloak (5542, -0.37 DPS) [crafted]; Black Whelp Cloak (7283, -0.62 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.92 DPS) | yes | Green Woolen Vest (2582, -0.18 DPS) [crafted]; Bloody Apron (6226, -0.18 DPS) [dungeon]; Gray Woolen Robe (2585, -1.34 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.55 DPS) | yes | Owlbeard Bracers (16981, -0.37 DPS) [quest]; Windsong Bangles (263336, -0.40 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.29 DPS) | yes | Gnoll Casting Gloves (892, -0.29 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.55 DPS) [quest]; Pristine Gloves (253913, -0.55 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Ardent's Sash (253887, -0.37 DPS) [crafted]; Flame Sash (253929, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.67 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Colorful Kilt (10048, -0.18 DPS) [crafted]; Embroidered Leggings (270011, -0.37 DPS) [quest]; Silk-threaded Trousers (1929, -1.61 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.29 DPS) | yes | Pristine Boots (253889, -0.74 DPS) [crafted]; Red Woolen Boots (4313, -0.85 DPS, sim-verified) [crafted]; Soft-soled Linen Boots (4312, -0.92 DPS) [crafted] |
| finger1 | Sludge-Stained Band (286535) | Sludge Beast (rare) [world] | 3.0 spell_power points (0.55 DPS) | yes | - |
| finger2 | - | - |  |  |  |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 98 more) | Blackfathom Deeps [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 1.5 spell_power points (0.27 DPS) | yes | Bouquet of Red Roses (22206, -0.27 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 124.0 spell_power points (22.86 DPS) | yes | Skycaller (12984, -0.60 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.29 DPS) [quest]; Greater Magic Wand (11288, -4.99 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 63.4. Weights run: 1.5s. Verify run: 0.9s. 283 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.062 ± 0.005, crit=0.046 ± 0.001 per rating point (14 rating = 1%, 0.639 per %), hit=0.121 ± 0.005 per rating point (10 rating = 1%, 1.208 per %), spell_haste=0.309 ± 0.064, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.002, fire_power=0.235 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.31 DPS) | yes | Filigreed Pristine Circlet (253975, -0.63 DPS) [crafted]; Enchanter's Cowl (4322, -0.92 DPS) [crafted]; Silk Headband (7050, -3.38 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.4 spell_power points (0.08 DPS) | yes | Crystal Starfire Medallion (5003, -0.03 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.03 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.6 spell_power points (2.01 DPS) | yes | Chestnut Mantle (17695, -0.40 DPS, sim-verified) [quest]; Invoker's Mantle (215365, -0.47 DPS) [crafted]; Death Speaker Mantle (6685, -0.60 DPS) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.05 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted]; Black Whelp Cloak (7283, -0.42 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.73 DPS) | yes | Green Silk Armor (7065, -0.65 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.68 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.89 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.89 DPS) | yes | Tabitha's Cuffs (251486, -1.26 DPS) [quest]; Owlbeard Bracers (16981, -1.65 DPS) [quest]; Phoenix Bindings (210781, -1.70 DPS, sim-verified) [crafted] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.47 DPS) | yes | Jutebraid Gloves (10654, -0.14 DPS) [quest]; Gnoll Casting Gloves (892, -0.21 DPS) [world]; Truefaith Gloves (7049, -0.38 DPS) [crafted] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.31 DPS) | yes | Belt of Arugal (6392, -0.32 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.63 DPS) [dungeon]; Invoker's Cord (215366, -0.78 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (63.4 DPS) | yes | Silk-threaded Trousers (1929, -0.09 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.22 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.66 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.4 spell_power points (1.56 DPS) | yes | Glinteye Slippers (273024, -0.09 DPS) [dungeon]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -2.08 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.3 spell_power points (1.33 DPS) | yes | Solid Gold Band (279845, -0.49 DPS) [quest]; Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon]; Runemender's Seal (273031, -0.70 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.26 DPS) | yes | Solid Gold Band (279845, -0.42 DPS) [quest]; Electrocutioner Lagnut (9447, -0.63 DPS) [dungeon]; Runemender's Seal (273031, -0.63 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (7.56 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -5.88 DPS) [quest]; Glimmering Staff (249392, -7.42 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 156.7 spell_power points (32.91 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.86 DPS) [crafted]; Unstable Power Core (279847, -3.20 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 119.2. Weights run: 1.4s. Verify run: 1.0s. 381 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.236 ± 0.007, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.924 per %), hit=0.179 ± 0.008 per rating point (10 rating = 1%, 1.788 per %), spell_haste=0.890 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.796 ± 0.002, fire_power=0.204 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.22 DPS) | yes | Augural Shroud (2620, -1.97 DPS, sim-verified) [world]; Living Cowl (5608, -1.99 DPS) [world]; Papal Fez (9431, -2.23 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.9 spell_power points (2.21 DPS) | yes | Triune Amulet (7722, -0.79 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.80 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.86 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 spell_power points (2.76 DPS) | yes | Green Silken Shoulders (7057, -0.13 DPS) [crafted]; Inquisitor's Shawl (19507, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.44 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 spell_power points (1.78 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.04 DPS) [crafted]; Hillman's Cloak (3719, -0.54 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.4 spell_power points (5.82 DPS) | yes | Elemental Raiment (9434, -0.60 DPS) [world_drop]; Dreamweave Vest (10021, -0.82 DPS) [crafted]; Robe of Power (7054, -1.64 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.24 DPS) | yes | Condor Bracers (15864, -0.50 DPS) [quest]; Radiant Silver Bracers (4545, -0.77 DPS) [quest]; Earthen Silk Cuffs (254019, -1.24 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (4.71 DPS) | yes | Black Mageweave Gloves (10003, -0.63 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.39 DPS) [crafted]; Gilded Handwraps (254021, -2.31 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.23 DPS) | yes | Deathmage Sash (10771, -0.61 DPS) [dungeon]; Warsong Sash (16975, -0.74 DPS, sim-verified) [quest]; Gilded Cord (254037, -0.77 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.8 spell_power points (4.18 DPS) | yes | Gaze Dreamer Pants (6903, -1.20 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.43 DPS) [crafted]; Stoneweaver Leggings (9407, -1.97 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.96 DPS) | yes | Gilded Slippers (254001, -2.52 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.99 DPS) [crafted]; Glinteye Slippers (273024, -4.22 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.2 spell_power points (2.78 DPS) | yes | Reedknot Ring (9622, -1.04 DPS) [quest]; Ring of Power Regulation (273030, -1.29 DPS) [dungeon]; Solid Gold Band (279845, -1.78 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.2 spell_power points (1.78 DPS) | yes | Reedknot Ring (9622, -0.04 DPS) [quest]; Ring of Power Regulation (273030, -0.29 DPS) [dungeon]; Solid Gold Band (279845, -0.79 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.12 DPS) [dungeon]; Spine of the Devourer (273040, -2.46 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.48 DPS) | yes | Omega Orb (7749, -0.64 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -1.39 DPS) [crafted]; Beacon of Hope (9393, -1.49 DPS) [dungeon] |
| ranged | Earthen Rod (9381) | Uldaman [dungeon] | sim-verified (119.2 DPS) | yes | Umbral Wand (5216, -0.67 DPS) [dungeon]; Twisted Nether Wand (249144, -0.84 DPS) [crafted]; Jaina's Firestarter (13064, -1.47 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Earthen Rod

No-known-source sample (15 of 381, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 177.9. Weights run: 1.4s. Verify run: 1.2s. 492 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.233 ± 0.019, crit=0.115 ± 0.004 per rating point (14 rating = 1%, 1.605 per %), hit=0.354 ± 0.019 per rating point (10 rating = 1%, 3.540 per %), spell_haste=-1.487 ± 0.300, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.509 ± 0.001, fire_power=0.491 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.54 DPS) | yes | Dreamweave Circlet (10041, -0.89 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.45 DPS) [crafted]; Red Mageweave Headband (10033, -1.75 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.9 spell_power points (2.15 DPS) | yes | Triune Amulet (7722, -0.21 DPS) [dungeon]; Mindburst Medallion (11196, -0.35 DPS) [quest]; Horizon Choker (13085, -1.36 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.2 spell_power points (4.16 DPS) | yes | Black Mageweave Shoulders (10027, -1.18 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.33 DPS) [vendor]; Bloodmage Mantle (7684, -1.48 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.4 spell_power points (3.73 DPS) | yes | Deep Woodlands Cloak (19121, -0.72 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.10 DPS) [crafted]; Nightfall Drape (12465, -1.55 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.7 spell_power points (5.73 DPS) | yes | Robe of the Magi (1716, -0.06 DPS) [world_drop]; Elemental Raiment (9434, -0.64 DPS) [world_drop]; Dreamweave Vest (10021, -0.86 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.18 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Netherflame Cuffs (254065, -0.24 DPS) [crafted]; Bloodband Bracers (11469, -0.46 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (4.58 DPS) | yes | Black Mageweave Gloves (10003, -0.95 DPS) [crafted]; Runecloth Gloves (13863, -1.17 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.50 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.3 spell_power points (3.95 DPS) | yes | Ghostweave Cord (254073, -0.56 DPS) [crafted]; Star Belt (4329, -0.81 DPS) [crafted]; Gilded Waistcord (254081, -1.02 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (4.60 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.63 DPS) [vendor]; Red Mageweave Pants (10009, -1.07 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -1.30 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.81 DPS) | yes | Gilded Sandals (254107, -2.64 DPS) [crafted]; Black Mageweave Boots (10026, -2.75 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.95 DPS, sim-verified) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.15 DPS) | yes | Snake Hoop (6750, -1.41 DPS) [quest]; Reedknot Ring (9622, -1.45 DPS) [quest]; Ring of Power Regulation (273030, -1.69 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.2 spell_power points (2.70 DPS) | yes | Snake Hoop (6750, -0.99 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.01 DPS) [quest]; Ring of Power Regulation (273030, -1.25 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (177.9 DPS) | yes | Rune of the Guard Captain (19120, -0.17 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (177.9 DPS) | yes | Rune of the Guard Captain (19120, -2.40 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (177.9 DPS) | yes | Spire of Hakkar (10844, -0.19 DPS) [world]; Inventor's Focal Sword (17719, -0.68 DPS) [dungeon]; Hanzo Sword (8190, -19.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 216.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.29 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.81 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 492, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 355.1. Weights run: 1.4s. Verify run: 1.9s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.273 ± 0.029, crit=0.145 ± 0.005 per rating point (14 rating = 1%, 2.025 per %), hit=0.506 ± 0.030 per rating point (10 rating = 1%, 5.057 per %), spell_haste=1.814 ± 0.394, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.627 ± 0.001, fire_power=0.373 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.20 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -0.96 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.7 spell_power points (5.57 DPS) | yes | Glowing Eye of Mordresh (10769, -2.10 DPS) [dungeon]; Triune Amulet (7722, -2.54 DPS) [dungeon]; Mindburst Medallion (11196, -2.68 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.44 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.80 DPS, sim-verified) [crafted]; Rotgrip Mantle (17732, -2.67 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (5.91 DPS) | yes | Deep Woodlands Cloak (19121, -0.45 DPS) [quest]; Runecloth Cloak (13860, -1.68 DPS) [crafted]; Nightfall Drape (12465, -2.51 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.5 spell_power points (18.31 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.74 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -6.80 DPS) [pvp]; Acumen Robes (17775, -9.67 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.7 spell_power points (5.19 DPS) | yes | Spidertank Oilrag (9448, -1.79 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.82 DPS) [crafted]; Dreadmist Bracers (16703, -2.16 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.31 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Blood Guard's Dreadweave Handwraps (227099, +0.00 DPS) [pvp]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (355.1 DPS) | yes | Ghostweave Cord (254073, -0.31 DPS) [crafted]; Star Belt (4329, -0.69 DPS) [crafted]; Satyrmane Sash (17755, -4.64 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.82 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.76 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.07 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.76 DPS) [crafted]; Omnicast Boots (11822, -2.00 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.9 spell_power points (7.52 DPS) | yes | Philanthropist's Ring (281635, -3.23 DPS) [quest]; Snake Hoop (6750, -4.74 DPS) [quest]; Reedknot Ring (9622, -4.88 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.91 DPS) | yes | Snake Hoop (6750, -2.13 DPS) [quest]; Reedknot Ring (9622, -2.27 DPS) [quest]; Philanthropist's Ring (281635, -2.27 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.38 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -2.24 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.48 DPS) [world]; Teebu's Blazing Longsword (1728, -32.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 209.2 spell_power points (79.03 DPS) | yes | Brilliant Wand (249385, -15.32 DPS) [crafted]; Stormrager (16997, -15.82 DPS) [quest]; Torch of Austen (13004, -22.24 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Wisdom of the Timbermaw; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 05000000000000000-2035113011220001351-0550005000000000)

Set DPS (verified): 740.4. Weights run: 1.7s. Verify run: 2.4s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.266 ± 0.043, crit=0.374 ± 0.013 per rating point (14 rating = 1%, 5.242 per %), hit=0.846 ± 0.055 per rating point (10 rating = 1%, 8.458 per %), spell_haste=3.791 ± 0.798, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.629 ± 0.002, fire_power=0.371 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (14.88 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -1.48 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.7 spell_power points (8.08 DPS) | yes | Glowing Eye of Mordresh (10769, -3.05 DPS) [dungeon]; Triune Amulet (7722, -3.67 DPS) [dungeon]; Mindburst Medallion (11196, -3.89 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.7 spell_power points (14.16 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.38 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -2.75 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (8.59 DPS) | yes | Deep Woodlands Cloak (19121, -0.66 DPS) [quest]; Runecloth Cloak (13860, -2.46 DPS) [crafted]; Nightfall Drape (12465, -3.63 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.4 spell_power points (26.66 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.52 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.96 DPS) [pvp]; Acumen Robes (17775, -11.11 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.7 spell_power points (7.53 DPS) | yes | Spidertank Oilrag (9448, -2.57 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.64 DPS) [crafted]; Earthenweave Cuffs (254131, -3.12 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.1 spell_power points (12.20 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.04 DPS) [pvp]; Earth Warder's Gloves (21318, -0.08 DPS) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.7 spell_power points (9.18 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -1.46 DPS) [crafted]; Star Belt (4329, -2.02 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (14.33 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS) [vendor]; Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.22 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.45 DPS) [dungeon]; Venomspew Footpads (275606, -1.10 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.9 spell_power points (10.94 DPS) | yes | Philanthropist's Ring (281635, -4.70 DPS) [quest]; Snake Hoop (6750, -6.90 DPS) [quest]; Reedknot Ring (9622, -7.09 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.16 DPS) | yes | Philanthropist's Ring (281635, -0.92 DPS) [quest]; Snake Hoop (6750, -3.12 DPS) [quest]; Reedknot Ring (9622, -3.31 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (740.4 DPS) | yes | Uther's Strength (11302, -0.89 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (740.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (740.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.66 DPS) [world]; Teebu's Blazing Longsword (1728, -49.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 145.3 spell_power points (80.07 DPS) | yes | Brilliant Wand (249385, -14.50 DPS) [crafted]; Stormrager (16997, -16.64 DPS) [quest]; Greater Eternal Wand (249237, -21.78 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

