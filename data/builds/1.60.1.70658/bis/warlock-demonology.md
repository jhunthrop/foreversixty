# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 31.4. Weights run: 1.8s. Verify run: 0.9s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.081 ± 0.002, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.362 per %), hit=0.064 ± 0.002 per rating point (10 rating = 1%, 0.644 per %), spell_haste=0.171 ± 0.012, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.251 ± 0.000, fire_power=0.749 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Circlet (253953, -0.81 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.16 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.7 spell_power points (1.19 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.22 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.36 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.83 DPS) | yes | Black Whelp Cloak (7283, -0.30 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS) [crafted]; Reinforced Linen Cape (2580, -0.62 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Gray Woolen Robe (2585, -0.21 DPS) [crafted]; Green Woolen Vest (2582, -0.29 DPS) [crafted]; Filigreed Flame Gown (253905, -0.56 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.21 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Bright Bracers (3647, -0.14 DPS) [world_drop]; Repurposed Hair Band (281256, -0.17 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.45 DPS) | yes | Phoenix Gloves (4331, -0.13 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.21 DPS) [world]; Flame Gloves (253917, -0.31 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.02 DPS) [crafted]; Novice Ardent's Sash (253887, -0.43 DPS) [crafted]; Flame Sash (253929, -0.52 DPS, sim-verified) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 8.7 spell_power points (1.81 DPS) | yes | Silk-threaded Trousers (1929, -0.36 DPS) [dungeon]; Phoenix Pants (4317, -0.40 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.3 spell_power points (1.52 DPS) | yes | Flame Boots (253893, -0.55 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.69 DPS) [crafted]; Pristine Boots (253889, -0.85 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.2 spell_power points (1.07 DPS) | yes | Volcanic Rock Ring (12053, -1.02 DPS) [world_drop]; The 1 Ring (8350, -1.05 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.5 spell_power points (0.10 DPS) | yes | The 1 Ring (8350, -0.08 DPS) [world]; Volcanic Rock Ring (12053, -0.26 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.8 spell_power points (0.17 DPS) | yes | Lesser Staff of the Spire (1300, -0.07 DPS) [world]; Witching Stave (1484, -0.08 DPS) [dungeon]; Staff of Westfall (2042, -0.08 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 110.7 spell_power points (22.93 DPS) | yes | Skycaller (12984, -0.67 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.25 DPS) [quest]; Deepblaze (279896, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 63.7. Weights run: 1.8s. Verify run: 0.9s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.089 ± 0.003, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.582 per %), hit=0.102 ± 0.004 per rating point (10 rating = 1%, 1.016 per %), spell_haste=0.321 ± 0.019, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.244 ± 0.000, fire_power=0.756 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.46 DPS) | yes | Silk Headband (7050, -0.45 DPS) [crafted]; Flame Circlet (253953, -0.60 DPS) [crafted]; Filigreed Flame Circlet (253979, -3.11 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.5 spell_power points (0.12 DPS) | yes | Crystal Starfire Medallion (5003, -0.04 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.04 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (2.19 DPS) | yes | Invoker's Mantle (215365, -0.56 DPS, sim-verified) [crafted]; Moonlit Amice (11884, -0.63 DPS) [quest]; Death Speaker Mantle (6685, -0.63 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.12 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.90 DPS) | yes | Flame Gown (253965, -0.33 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.63 DPS) [crafted]; Beguiler Robes (7728, -0.65 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.8 spell_power points (2.19 DPS) | yes | Spidertank Oilrag (9448, -0.19 DPS) [dungeon]; Windsong Bangles (263336, -1.97 DPS) [quest]; Glowing Magical Bracelets (13106, -2.04 DPS) [world_drop] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.56 DPS) | yes | Phoenix Gloves (4331, -0.04 DPS) [crafted]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Serpent Gloves (5970, -0.46 DPS, sim-verified) [dungeon] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.3 spell_power points (2.07 DPS) | yes | Invoker's Cord (215366, -0.41 DPS) [crafted]; Crimson Silk Belt (7055, -0.59 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.71 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.68 DPS) | yes | Flame Leggings (253991, -0.18 DPS) [crafted]; Smoldering Pants (3073, -0.48 DPS) [world]; Filigreed Flame Leggings (253941, -0.70 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (63.7 DPS) | yes | Spidersilk Boots (4320, -0.06 DPS) [crafted]; Glinteye Slippers (273024, -0.14 DPS) [dungeon]; Fiery Slippers (254005, -0.86 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.4 spell_power points (1.44 DPS) | yes | Minor Channeling Ring (1449, -0.28 DPS) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Runemender's Seal (273031, -0.77 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.34 DPS) | yes | Minor Channeling Ring (1449, -0.36 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Runemender's Seal (273031, -0.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.04 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -2.29 DPS) [quest]; Hardened Root Staff (1317, -3.57 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 147.6 spell_power points (32.94 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.82 DPS) [crafted]; Scorching Wand (5213, -3.08 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Phoenix Bindings; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 114.6. Weights run: 1.6s. Verify run: 1.0s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.090 ± 0.004, crit=0.064 ± 0.001 per rating point (14 rating = 1%, 0.898 per %), hit=0.161 ± 0.006 per rating point (10 rating = 1%, 1.613 per %), spell_haste=0.600 ± 0.044, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.254 ± 0.000, fire_power=0.746 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.07 DPS) | yes | Living Cowl (5608, -2.00 DPS, sim-verified) [world]; Augural Shroud (2620, -2.19 DPS) [world]; Holy Shroud (2721, -2.41 DPS) [world_drop] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (1.93 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.78 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.80 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (2.37 DPS) | yes | Green Silken Shoulders (7057, -0.20 DPS) [crafted]; Inquisitor's Shawl (19507, -0.40 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (1.69 DPS) | yes | Guardian Cloak (5965, -0.13 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.48 DPS) [quest]; Long Silken Cloak (4326, -0.52 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 spell_power points (5.44 DPS) | yes | Cindercloth Robe (10042, -0.58 DPS) [crafted]; Dreamweave Vest (10021, -0.90 DPS) [crafted]; Elemental Raiment (9434, -0.94 DPS, sim-verified) [world_drop] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.7 spell_power points (2.34 DPS) | yes | Arcane Runed Bracers (4744, -0.17 DPS) [quest]; Spidertank Oilrag (9448, -0.17 DPS) [dungeon]; Condor Bracers (15864, -0.65 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (4.43 DPS) | yes | Black Mageweave Gloves (10003, -1.37 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.56 DPS) [crafted]; Fiery Handwraps (254025, -1.76 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.14 DPS) | yes | Fiery Cord (254041, -0.50 DPS, sim-verified) [crafted]; Scorching Sash (4117, -0.80 DPS) [quest]; Belt of Arugal (6392, -0.90 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 spell_power points (3.64 DPS) | yes | Gaze Dreamer Pants (6903, -0.59 DPS, sim-verified) [dungeon]; Flame Leggings (253991, -0.97 DPS) [crafted]; Smoldering Pants (3073, -1.30 DPS) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.79 DPS) | yes | Fiery Slippers (254005, -3.48 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.95 DPS) [crafted]; Spidersilk Boots (4320, -4.01 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 spell_power points (2.52 DPS) | yes | Reedknot Ring (9622, -0.83 DPS) [quest]; Snake Hoop (6750, -0.97 DPS) [quest]; Ring of Power Regulation (273030, -1.07 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (1.93 DPS) | yes | Reedknot Ring (9622, -0.24 DPS) [quest]; Snake Hoop (6750, -0.37 DPS) [quest]; Ring of Power Regulation (273030, -0.48 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (114.6 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.04 DPS) [dungeon]; Spine of the Devourer (273040, -2.10 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.38 DPS) | yes | Orb of Noh'Orahil (15107, -0.65 DPS, sim-verified) [quest]; Rod of Molten Fire (2565, -1.04 DPS) [world_drop]; Beacon of Hope (9393, -1.45 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 163.8 spell_power points (39.51 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -3.84 DPS) [dungeon]; Twisted Nether Wand (249144, -4.06 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 193.5. Weights run: 4.0s. Verify run: 1.1s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.026 ± 0.013, crit=0.116 ± 0.004 per rating point (14 rating = 1%, 1.626 per %), hit=0.359 ± 0.018 per rating point (10 rating = 1%, 3.590 per %), spell_haste=-2.032 ± 0.320, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.256 ± 0.000, fire_power=0.744 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.86 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.52 DPS) [crafted]; Red Mageweave Headband (10033, -1.90 DPS) [crafted]; Dreamweave Circlet (10041, -2.39 DPS, sim-verified) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (2.03 DPS) | yes | Glowing Eye of Mordresh (10769, -0.20 DPS) [dungeon]; Mindburst Medallion (11196, -0.47 DPS) [quest]; Horizon Choker (13085, -1.94 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.5 spell_power points (3.42 DPS) | yes | Black Mageweave Shoulders (10027, -0.82 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -0.91 DPS) [vendor]; Netherflame Shoulders (254053, -1.13 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.2 spell_power points (3.59 DPS) | yes | Cindercloth Cloak (14044, -1.09 DPS) [crafted]; Runecloth Cloak (13860, -1.26 DPS) [crafted]; Cloak of Fire (14134, -3.02 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.2 spell_power points (5.63 DPS) | yes | Cindercloth Robe (10042, -0.53 DPS) [crafted]; Acumen Robes (17775, -0.67 DPS) [quest]; Elemental Raiment (9434, -2.41 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 9.8 spell_power points (2.50 DPS) | yes | Arcane Runed Bracers (4744, -0.21 DPS) [quest]; Spidertank Oilrag (9448, -0.21 DPS) [dungeon]; Phoenix Bindings (210781, -1.96 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.1 spell_power points (4.60 DPS) | yes | Fiery Gloves (254099, -0.94 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.24 DPS) [vendor]; Black Mageweave Gloves (10003, -1.87 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 14.4 spell_power points (3.65 DPS) | yes | Satyrmane Sash (17755, -0.03 DPS) [dungeon]; Ghostweave Cord (254073, -0.09 DPS) [crafted]; Star Belt (4329, -0.34 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (4.82 DPS) | yes | Cindercloth Leggings (12256, -1.24 DPS) [vendor]; Knight's Dreadweave Leggings (220888, -1.29 DPS) [vendor]; Red Mageweave Pants (10009, -1.98 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.09 DPS) | yes | Fiery Sandals (254111, -2.26 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.09 DPS) [vendor]; Cindercloth Boots (10044, -3.52 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.30 DPS) | yes | Ring of Forlorn Spirits (2043, -1.27 DPS) [quest]; Reedknot Ring (9622, -1.52 DPS) [quest]; Snake Hoop (6750, -1.75 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.1 spell_power points (2.57 DPS) | yes | Reedknot Ring (9622, -0.79 DPS) [quest]; Snake Hoop (6750, -1.02 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.47 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (193.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (193.5 DPS) | yes | Mark of the Chosen (17774, -3.21 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (193.5 DPS) | yes | Kindling Stave (11750, -0.43 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.51 DPS) [dungeon]; Hanzo Sword (8190, -23.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 206.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.23 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.72 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Fiery Waistcord; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 356.6. Weights run: 1.6s. Verify run: 3.0s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.105 ± 0.024, crit=0.147 ± 0.005 per rating point (14 rating = 1%, 2.051 per %), hit=0.481 ± 0.026 per rating point (10 rating = 1%, 4.810 per %), spell_haste=1.508 ± 0.304, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.335 ± 0.001, fire_power=0.665 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.18 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.78 DPS) [pvp]; Dreamweave Circlet (10041, -2.29 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.0 spell_power points (4.92 DPS) | yes | Glowing Eye of Mordresh (10769, -1.96 DPS) [dungeon]; Mindburst Medallion (11196, -2.42 DPS) [quest]; Triune Amulet (7722, -2.70 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.43 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.62 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.64 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.52 DPS) | yes | Mageflame Cloak (13007, -0.25 DPS) [world_drop]; Runecloth Cloak (13860, -1.81 DPS) [crafted]; Cindercloth Cloak (14044, -1.94 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.9 spell_power points (17.70 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.69 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.48 DPS) [pvp]; Robe of the Magi (1716, -9.12 DPS, sim-verified) [world_drop] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (356.6 DPS) | yes | Arcane Runed Bracers (4744, +0.00 DPS) [quest]; Netherflame Cuffs (254065, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -6.95 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (356.6 DPS) | yes | Knight-Lieutenant's Dreadweave Handwraps (227100, +0.00 DPS) [pvp]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -4.19 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (356.6 DPS) | yes | Ghostweave Cord (254073, +0.00 DPS) [crafted]; Fiery Waistcord (254085, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -6.73 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.81 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -3.30 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.05 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.75 DPS) [crafted]; Omnicast Boots (11822, -1.03 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.7 spell_power points (7.06 DPS) | yes | Philanthropist's Ring (281635, -3.10 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.05 DPS) [quest]; Reedknot Ring (9622, -4.43 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.90 DPS) | yes | Ring of Forlorn Spirits (2043, -1.89 DPS) [quest]; Philanthropist's Ring (281635, -1.95 DPS, sim-verified) [quest]; Reedknot Ring (9622, -2.26 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (356.6 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (356.6 DPS) | yes | Mark of the Chosen (17774, -4.72 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (356.6 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.28 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -32.88 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 209.5 spell_power points (79.02 DPS) | yes | Brilliant Wand (249385, -2.07 DPS, sim-verified) [crafted]; Stormrager (16997, -16.14 DPS) [quest]; Torch of Austen (13004, -22.23 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 05400000000000000-2035113111120001351-0050051000000000)

Set DPS (verified): 682.0. Weights run: 5.3s. Verify run: 2.1s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.113 ± 0.021, crit=0.231 ± 0.007 per rating point (14 rating = 1%, 3.230 per %), hit=0.702 ± 0.039 per rating point (10 rating = 1%, 7.016 per %), spell_haste=not significant (1.396 ± 0.474), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.266 ± 0.001, fire_power=0.734 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.94 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.38 DPS) [pvp]; Dreamweave Circlet (10041, -2.52 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.1 spell_power points (6.78 DPS) | yes | Triune Amulet (7722, -2.65 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.70 DPS) [dungeon]; Mindburst Medallion (11196, -3.33 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (12.91 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -4.29 DPS) [pvp] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 15.4 spell_power points (7.96 DPS) | yes | Spritecaster Cape (11623, -0.38 DPS) [dungeon]; Cindercloth Cloak (14044, -2.57 DPS) [crafted]; Cloak of Fire (14134, -2.65 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (24.27 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.35 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.20 DPS) [pvp]; Robe of the Magi (1716, -13.14 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.1 spell_power points (6.26 DPS) | yes | Netherflame Cuffs (254065, -0.93 DPS) [crafted]; Phoenix Bindings (210781, -1.33 DPS) [crafted]; Arcane Runed Bracers (4744, -1.61 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (11.36 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.28 DPS) [pvp]; Dreamweave Gloves (10019, -1.83 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.1 spell_power points (7.81 DPS) | yes | Fiery Waistcord (254085, -0.09 DPS) [crafted]; Ghostweave Cord (254073, -0.58 DPS) [crafted]; Star Belt (4329, -1.10 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (13.42 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -2.87 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.39 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.03 DPS) [crafted]; Omnicast Boots (11822, -1.37 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.8 spell_power points (9.70 DPS) | yes | Philanthropist's Ring (281635, -4.25 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.57 DPS) [quest]; Reedknot Ring (9622, -6.09 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.71 DPS) | yes | Philanthropist's Ring (281635, -1.26 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.58 DPS) [quest]; Reedknot Ring (9622, -3.10 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (682.0 DPS) | yes | Uther's Strength (11302, -0.16 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (682.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (682.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.33 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -42.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 154.7 spell_power points (79.86 DPS) | yes | Brilliant Wand (249385, -15.22 DPS) [crafted]; Stormrager (16997, -16.88 DPS) [quest]; Greater Eternal Wand (249237, -21.88 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; back: Mageflame Cloak; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; ranged: Torch of Light

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 30.3. Weights run: 1.8s. Verify run: 1.0s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.081 ± 0.002, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.362 per %), hit=0.064 ± 0.002 per rating point (10 rating = 1%, 0.644 per %), spell_haste=0.171 ± 0.012, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.251 ± 0.000, fire_power=0.749 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Circlet (253953, -0.80 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.16 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.7 spell_power points (1.19 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.36 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.83 DPS) | yes | Black Whelp Cloak (7283, -0.36 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS) [crafted]; Mystic Shawl (3449, -0.41 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Gray Woolen Robe (2585, -0.21 DPS) [crafted]; Green Woolen Vest (2582, -0.29 DPS) [crafted]; Filigreed Flame Gown (253905, -0.55 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.62 DPS) | yes | Owlbeard Bracers (16981, -0.39 DPS, sim-verified) [quest]; Windsong Bangles (263336, -0.41 DPS) [quest]; Featherbead Bracers (15452, -0.54 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.45 DPS) | yes | Phoenix Gloves (4331, -0.12 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.21 DPS) [world]; Flame Gloves (253917, -0.31 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.02 DPS) [crafted]; Novice Ardent's Sash (253887, -0.43 DPS) [crafted]; Flame Sash (253929, -0.52 DPS, sim-verified) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 8.7 spell_power points (1.81 DPS) | yes | Phoenix Pants (4317, -0.35 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.36 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.3 spell_power points (1.52 DPS) | yes | Flame Boots (253893, -0.54 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.69 DPS) [crafted]; Pristine Boots (253889, -0.85 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.5 spell_power points (0.10 DPS) | yes | Volcanic Rock Ring (12053, -0.05 DPS) [world_drop]; The 1 Ring (8350, -0.08 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 0.4 spell_power points (0.08 DPS) | yes | The 1 Ring (8350, -0.07 DPS) [world]; Volcanic Rock Ring (12053, -0.21 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.8 spell_power points (0.17 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.07 DPS) [world]; Clear Crystal Rod (16894, -0.08 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 110.7 spell_power points (22.93 DPS) | yes | Skycaller (12984, -0.66 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.25 DPS) [quest]; Greater Magic Wand (11288, -5.02 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 62.3. Weights run: 1.8s. Verify run: 0.9s. 283 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.089 ± 0.003, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.582 per %), hit=0.102 ± 0.004 per rating point (10 rating = 1%, 1.016 per %), spell_haste=0.321 ± 0.019, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.244 ± 0.000, fire_power=0.756 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.46 DPS) | yes | Silk Headband (7050, -0.45 DPS) [crafted]; Flame Circlet (253953, -0.60 DPS) [crafted]; Filigreed Flame Circlet (253979, -2.82 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.5 spell_power points (0.12 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.04 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (2.19 DPS) | yes | Chestnut Mantle (17695, -0.30 DPS, sim-verified) [quest]; Invoker's Mantle (215365, -0.53 DPS) [crafted]; Death Speaker Mantle (6685, -0.63 DPS) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.12 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Black Whelp Cloak (7283, -0.45 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.90 DPS) | yes | Flame Gown (253965, -0.57 DPS) [crafted]; Green Silk Armor (7065, -0.63 DPS) [crafted]; Beguiler Robes (7728, -0.65 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.8 spell_power points (2.19 DPS) | yes | Spidertank Oilrag (9448, -0.19 DPS) [dungeon]; Tabitha's Cuffs (251486, -1.53 DPS) [quest]; Owlbeard Bracers (16981, -1.93 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.56 DPS) | yes | Phoenix Gloves (4331, -0.04 DPS) [crafted]; Jutebraid Gloves (10654, -0.12 DPS) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.46 DPS) | yes | Belt of Arugal (6392, -0.39 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.67 DPS) [dungeon]; Invoker's Cord (215366, -0.79 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.68 DPS) | yes | Flame Leggings (253991, +0.00 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.48 DPS) [world]; Filigreed Flame Leggings (253941, -0.70 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (62.3 DPS) | yes | Spidersilk Boots (4320, -0.06 DPS) [crafted]; Glinteye Slippers (273024, -0.14 DPS) [dungeon]; Fiery Slippers (254005, -0.85 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.4 spell_power points (1.44 DPS) | yes | Solid Gold Band (279845, -0.55 DPS) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Runemender's Seal (273031, -0.77 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.34 DPS) | yes | Solid Gold Band (279845, -0.45 DPS) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Runemender's Seal (273031, -0.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (8.04 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -2.29 DPS) [quest]; Glimmering Staff (249392, -7.82 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 147.6 spell_power points (32.94 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -2.82 DPS) [crafted]; Scorching Wand (5213, -3.08 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Phoenix Bindings; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 112.7. Weights run: 1.6s. Verify run: 1.0s. 381 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.090 ± 0.004, crit=0.064 ± 0.001 per rating point (14 rating = 1%, 0.898 per %), hit=0.161 ± 0.006 per rating point (10 rating = 1%, 1.613 per %), spell_haste=0.600 ± 0.044, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.254 ± 0.000, fire_power=0.746 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.07 DPS) | yes | Living Cowl (5608, -1.96 DPS, sim-verified) [world]; Augural Shroud (2620, -2.19 DPS) [world]; Holy Shroud (2721, -2.41 DPS) [world_drop] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (1.93 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.78 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.80 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (2.37 DPS) | yes | Green Silken Shoulders (7057, -0.20 DPS) [crafted]; Inquisitor's Shawl (19507, -0.40 DPS) [dungeon]; Chestnut Mantle (17695, -0.44 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (1.69 DPS) | yes | Long Silken Cloak (4326, -0.13 DPS) [crafted]; Guardian Cloak (5965, -0.13 DPS) [crafted]; Hillman's Cloak (3719, -0.48 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 spell_power points (5.44 DPS) | yes | Cindercloth Robe (10042, -0.58 DPS) [crafted]; Elemental Raiment (9434, -0.65 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.90 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.7 spell_power points (2.34 DPS) | yes | Spidertank Oilrag (9448, -0.17 DPS) [dungeon]; Condor Bracers (15864, -0.65 DPS) [quest]; Radiant Silver Bracers (4545, -1.20 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (4.43 DPS) | yes | Black Mageweave Gloves (10003, -1.13 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.56 DPS) [crafted]; Fiery Handwraps (254025, -1.76 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.14 DPS) | yes | Fiery Cord (254041, -0.44 DPS) [crafted]; Warsong Sash (16975, -0.48 DPS) [quest]; Scorching Sash (4117, -0.80 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 spell_power points (3.64 DPS) | yes | Gaze Dreamer Pants (6903, -0.51 DPS, sim-verified) [dungeon]; Flame Leggings (253991, -0.97 DPS) [crafted]; Smoldering Pants (3073, -1.30 DPS) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.79 DPS) | yes | Fiery Slippers (254005, -3.57 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.95 DPS) [crafted]; Spidersilk Boots (4320, -4.01 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 spell_power points (2.52 DPS) | yes | Snake Hoop (6750, -0.97 DPS) [quest]; Ring of Power Regulation (273030, -1.07 DPS) [dungeon]; Solid Gold Band (279845, -1.56 DPS) [quest] |
| finger2 | Reedknot Ring (9622) | Jarl Needs a Blade [quest] | 7.0 spell_power points (1.69 DPS) | yes | Snake Hoop (6750, -0.13 DPS) [quest]; Ring of Power Regulation (273030, -0.24 DPS) [dungeon]; Solid Gold Band (279845, -0.72 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (112.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.04 DPS) [dungeon]; Spine of the Devourer (273040, -2.10 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.38 DPS) | yes | Orb of Noh'Orahil (15107, -0.49 DPS, sim-verified) [quest]; Rod of Molten Fire (2565, -1.04 DPS) [world_drop]; Omega Orb (7749, -1.21 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 163.8 spell_power points (39.51 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -3.84 DPS) [dungeon]; Twisted Nether Wand (249144, -4.06 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Reedknot Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 381, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 190.4. Weights run: 4.0s. Verify run: 1.1s. 492 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.026 ± 0.013, crit=0.116 ± 0.004 per rating point (14 rating = 1%, 1.626 per %), hit=0.359 ± 0.018 per rating point (10 rating = 1%, 3.590 per %), spell_haste=-2.032 ± 0.320, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.256 ± 0.000, fire_power=0.744 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.86 DPS) | yes | Dreamweave Circlet (10041, -1.46 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.52 DPS) [crafted]; Red Mageweave Headband (10033, -1.90 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (2.03 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.47 DPS) [quest]; Horizon Choker (13085, -1.94 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.5 spell_power points (3.42 DPS) | yes | Black Mageweave Shoulders (10027, -0.82 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -0.91 DPS) [vendor]; Netherflame Shoulders (254053, -1.06 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.2 spell_power points (3.59 DPS) | yes | Deep Woodlands Cloak (19121, -0.49 DPS) [quest]; Cloak of Fire (14134, -0.95 DPS) [crafted]; Cindercloth Cloak (14044, -1.09 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.2 spell_power points (5.63 DPS) | yes | Cindercloth Robe (10042, -0.53 DPS) [crafted]; Acumen Robes (17775, -0.67 DPS) [quest]; Elemental Raiment (9434, -1.14 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 9.8 spell_power points (2.50 DPS) | yes | Spidertank Oilrag (9448, -0.21 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.68 DPS) [crafted]; Phoenix Bindings (210781, -1.33 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.1 spell_power points (4.60 DPS) | yes | Black Mageweave Gloves (10003, -0.79 DPS) [crafted]; Fiery Gloves (254099, -0.94 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.24 DPS) [vendor] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 14.4 spell_power points (3.65 DPS) | yes | Satyrmane Sash (17755, -0.03 DPS) [dungeon]; Ghostweave Cord (254073, -0.09 DPS) [crafted]; Star Belt (4329, -0.34 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (4.82 DPS) | yes | Red Mageweave Pants (10009, -1.19 DPS) [crafted]; Cindercloth Leggings (12256, -1.24 DPS) [vendor]; Stone Guard's Dreadweave Leggings (220906, -1.29 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.09 DPS) | yes | Cindercloth Boots (10044, -1.93 DPS, sim-verified) [crafted]; Fiery Sandals (254111, -2.26 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.09 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.30 DPS) | yes | Reedknot Ring (9622, -1.52 DPS) [quest]; Snake Hoop (6750, -1.75 DPS) [quest]; Ring of Power Regulation (273030, -1.78 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.1 spell_power points (2.57 DPS) | yes | Snake Hoop (6750, -1.02 DPS) [quest]; Ring of Power Regulation (273030, -1.05 DPS) [dungeon]; Reedknot Ring (9622, -1.29 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (190.4 DPS) | yes | Rune of the Guard Captain (19120, -0.18 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (190.4 DPS) | yes | Rune of the Guard Captain (19120, -1.98 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (190.4 DPS) | yes | Kindling Stave (11750, -0.43 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.51 DPS) [dungeon]; Hanzo Sword (8190, -21.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 206.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.72 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Fiery Waistcord; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 492, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 349.8. Weights run: 1.6s. Verify run: 1.9s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.105 ± 0.024, crit=0.147 ± 0.005 per rating point (14 rating = 1%, 2.051 per %), hit=0.481 ± 0.026 per rating point (10 rating = 1%, 4.810 per %), spell_haste=1.508 ± 0.304, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.335 ± 0.001, fire_power=0.665 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.18 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.78 DPS) [pvp]; Dreamweave Circlet (10041, -3.78 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.0 spell_power points (4.92 DPS) | yes | Triune Amulet (7722, -1.94 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -1.96 DPS) [dungeon]; Mindburst Medallion (11196, -2.42 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.43 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -3.62 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.50 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.52 DPS) | yes | Mageflame Cloak (13007, -0.25 DPS) [world_drop]; Deep Woodlands Cloak (19121, -0.64 DPS) [quest]; Runecloth Cloak (13860, -1.81 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.9 spell_power points (17.70 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.69 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.48 DPS) [pvp]; Robe of the Magi (1716, -10.19 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.0 spell_power points (4.54 DPS) | yes | Spidertank Oilrag (9448, -1.15 DPS) [dungeon]; Phoenix Bindings (210781, -1.28 DPS) [crafted]; Netherflame Cuffs (254065, -1.75 DPS, sim-verified) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.30 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.22 DPS) [pvp]; Dreamweave Gloves (10019, -3.48 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.0 spell_power points (5.67 DPS) | yes | Ghostweave Cord (254073, -0.40 DPS) [crafted]; Fiery Waistcord (254085, -0.55 DPS) [crafted]; Star Belt (4329, -0.77 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.81 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.30 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.05 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.03 DPS) [dungeon]; Venomspew Footpads (275606, -1.60 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.7 spell_power points (7.06 DPS) | yes | Philanthropist's Ring (281635, -3.10 DPS) [quest]; Reedknot Ring (9622, -4.43 DPS) [quest]; Snake Hoop (6750, -4.60 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.90 DPS) | yes | Philanthropist's Ring (281635, -1.85 DPS, sim-verified) [quest]; Reedknot Ring (9622, -2.26 DPS) [quest]; Snake Hoop (6750, -2.44 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (349.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (349.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (349.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.28 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -33.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 209.5 spell_power points (79.02 DPS) | yes | Brilliant Wand (249385, -2.44 DPS, sim-verified) [crafted]; Stormrager (16997, -16.14 DPS) [quest]; Torch of Austen (13004, -22.23 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 05400000000000000-2035113111120001351-0050051000000000)

Set DPS (verified): 680.5. Weights run: 5.3s. Verify run: 2.1s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.113 ± 0.021, crit=0.231 ± 0.007 per rating point (14 rating = 1%, 3.230 per %), hit=0.702 ± 0.039 per rating point (10 rating = 1%, 7.016 per %), spell_haste=not significant (1.396 ± 0.474), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.266 ± 0.001, fire_power=0.734 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.94 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.38 DPS) [pvp]; Dreamweave Circlet (10041, -2.86 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.1 spell_power points (6.78 DPS) | yes | Triune Amulet (7722, -2.65 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.70 DPS) [dungeon]; Mindburst Medallion (11196, -3.33 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (12.91 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.71 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -4.29 DPS) [pvp] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 15.4 spell_power points (7.96 DPS) | yes | Spritecaster Cape (11623, -0.38 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.24 DPS) [quest]; Cindercloth Cloak (14044, -2.57 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (24.27 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.35 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.20 DPS) [pvp]; Robe of the Magi (1716, -11.33 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.1 spell_power points (6.26 DPS) | yes | Netherflame Cuffs (254065, -0.93 DPS) [crafted]; Phoenix Bindings (210781, -1.33 DPS) [crafted]; Spidertank Oilrag (9448, -1.61 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (11.36 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.28 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.1 spell_power points (7.81 DPS) | yes | Fiery Waistcord (254085, +0.00 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -0.58 DPS) [crafted]; Star Belt (4329, -1.10 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (13.42 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -2.80 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.39 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.03 DPS) [crafted]; Omnicast Boots (11822, -1.37 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.8 spell_power points (9.70 DPS) | yes | Philanthropist's Ring (281635, -4.25 DPS) [quest]; Reedknot Ring (9622, -6.09 DPS) [quest]; Snake Hoop (6750, -6.31 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.71 DPS) | yes | Philanthropist's Ring (281635, -1.26 DPS) [quest]; Reedknot Ring (9622, -3.10 DPS) [quest]; Snake Hoop (6750, -3.32 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (680.5 DPS) | yes | Uther's Strength (11302, -0.16 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (680.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (680.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.33 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -42.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 154.7 spell_power points (79.86 DPS) | yes | Brilliant Wand (249385, -15.22 DPS) [crafted]; Stormrager (16997, -16.88 DPS) [quest]; Greater Eternal Wand (249237, -21.88 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; back: Mageflame Cloak; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

