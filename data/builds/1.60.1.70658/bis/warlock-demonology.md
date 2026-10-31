# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 33.6. Weights run: 1.6s. Verify run: 0.8s. 155 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.028 ± 0.002, crit=0.022 ± 0.001 per rating point (14 rating = 1%, 0.308 per %), hit=0.089 ± 0.003 per rating point (10 rating = 1%, 0.887 per %), spell_haste=0.927 ± 0.034, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.246 ± 0.000, fire_power=0.754 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Circlet (253953, -0.77 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.17 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.3 spell_power points (1.05 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.19 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Pearl-clasped Cloak (5542, -0.38 DPS) [crafted]; Black Whelp Cloak (7283, -0.43 DPS, sim-verified) [crafted]; Reinforced Linen Cape (2580, -0.60 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gray Woolen Robe (2585, -0.20 DPS) [crafted]; Green Woolen Vest (2582, -0.23 DPS) [crafted]; Filigreed Flame Gown (253905, -0.52 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.18 DPS) [world_drop]; Repurposed Hair Band (281256, -0.19 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Phoenix Gloves (4331, -0.04 DPS) [crafted]; Gnoll Casting Gloves (892, -0.20 DPS) [world]; Flame Gloves (253917, -0.33 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.01 DPS) [crafted]; Novice Ardent's Sash (253887, -0.40 DPS) [crafted]; Flame Sash (253929, -0.50 DPS, sim-verified) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 8.5 spell_power points (1.69 DPS) | yes | Phoenix Pants (4317, -0.28 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.29 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.1 spell_power points (1.42 DPS) | yes | Flame Boots (253893, -0.53 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.62 DPS) [crafted]; Pristine Boots (253889, -0.80 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.1 spell_power points (1.01 DPS) | yes | Volcanic Rock Ring (12053, -0.99 DPS) [world_drop]; The 1 Ring (8350, -1.00 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.2 spell_power points (0.03 DPS) | yes | The 1 Ring (8350, -0.03 DPS) [world]; Volcanic Rock Ring (12053, -0.22 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.3 spell_power points (0.06 DPS) | yes | Witching Stave (1484, -0.03 DPS) [dungeon]; Staff of Westfall (2042, -0.03 DPS) [quest]; Lesser Staff of the Spire (1300, -0.23 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.9 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.44 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.26 DPS) [quest]; Deepblaze (279896, -4.38 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 155, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 66.8. Weights run: 1.5s. Verify run: 0.9s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.044 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.423 per %), hit=0.125 ± 0.004 per rating point (10 rating = 1%, 1.247 per %), spell_haste=-1.894 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.239 ± 0.000, fire_power=0.761 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.42 DPS) | yes | Silk Headband (7050, -0.44 DPS) [crafted]; Flame Circlet (253953, -0.58 DPS) [crafted]; Filigreed Flame Circlet (253979, -2.58 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.3 spell_power points (0.06 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.02 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.4 spell_power points (2.07 DPS) | yes | Invoker's Mantle (215365, -0.29 DPS, sim-verified) [crafted]; Moonlit Amice (11884, -0.53 DPS) [quest]; Death Speaker Mantle (6685, -0.64 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.10 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.86 DPS) | yes | Flame Gown (253965, -0.62 DPS) [crafted]; Green Silk Armor (7065, -0.75 DPS) [crafted]; Beguiler Robes (7728, -0.76 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.9 spell_power points (2.18 DPS) | yes | Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Windsong Bangles (263336, -1.96 DPS) [quest]; Glowing Magical Bracelets (13106, -2.10 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.54 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Phoenix Gloves (4331, -0.03 DPS) [crafted]; Gnoll Casting Gloves (892, -0.22 DPS) [world] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.1 spell_power points (2.01 DPS) | yes | Ghamoo-ra's Bind (6908, -0.25 DPS) [dungeon]; Invoker's Cord (215366, -0.42 DPS) [crafted]; Crimson Silk Belt (7055, -0.62 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.64 DPS) | yes | Flame Leggings (253991, -0.23 DPS) [crafted]; Smoldering Pants (3073, -0.46 DPS) [world]; Filigreed Flame Leggings (253941, -0.74 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (66.8 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Glinteye Slippers (273024, -0.07 DPS) [dungeon]; Fiery Slippers (254005, -0.85 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.2 spell_power points (1.37 DPS) | yes | Minor Channeling Ring (1449, -0.25 DPS) [quest]; Electrocutioner Lagnut (9447, -0.71 DPS) [dungeon]; Runemender's Seal (273031, -0.71 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.32 DPS) | yes | Minor Channeling Ring (1449, -0.33 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon]; Runemender's Seal (273031, -0.66 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (7.92 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -2.23 DPS) [quest]; Hardened Root Staff (1317, -3.52 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 149.5 spell_power points (32.90 DPS) | yes | Wand of Mana Concentration (273043, -1.46 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -2.80 DPS) [crafted]; Scorching Wand (5213, -3.04 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Phoenix Bindings; hands: Shilly Mitts; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 116.3. Weights run: 1.3s. Verify run: 0.9s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.060 ± 0.006, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.582 per %), hit=0.187 ± 0.006 per rating point (10 rating = 1%, 1.868 per %), spell_haste=-6.930 ± 0.226, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.247 ± 0.000, fire_power=0.753 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.08 DPS) | yes | Living Cowl (5608, -1.95 DPS, sim-verified) [world]; Augural Shroud (2620, -2.27 DPS) [world]; Holy Shroud (2721, -2.42 DPS) [world_drop] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (1.94 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.83 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.85 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.31 DPS) | yes | Inquisitor's Shawl (19507, -0.43 DPS) [dungeon]; Berylline Pads (4197, -0.47 DPS) [quest]; Green Silken Shoulders (7057, -0.60 DPS, sim-verified) [crafted] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (1.69 DPS) | yes | Guardian Cloak (5965, -0.17 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.48 DPS) [quest]; Long Silken Cloak (4326, -0.81 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.4 spell_power points (5.41 DPS) | yes | Elemental Raiment (9434, -0.47 DPS, sim-verified) [world_drop]; Cindercloth Robe (10042, -0.49 DPS) [crafted]; Dreamweave Vest (10021, -0.92 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.8 spell_power points (2.37 DPS) | yes | Arcane Runed Bracers (4744, -0.19 DPS) [quest]; Spidertank Oilrag (9448, -0.19 DPS) [dungeon]; Condor Bracers (15864, -0.67 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.41 DPS) | yes | Black Mageweave Gloves (10003, -1.28 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.61 DPS) [crafted]; Fiery Handwraps (254025, -1.76 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.15 DPS) | yes | Fiery Cord (254041, -0.48 DPS) [crafted]; Scorching Sash (4117, -0.78 DPS) [quest]; Belt of Arugal (6392, -0.92 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.7 spell_power points (3.56 DPS) | yes | Gaze Dreamer Pants (6903, -0.73 DPS, sim-verified) [dungeon]; Flame Leggings (253991, -0.91 DPS) [crafted]; Smoldering Pants (3073, -1.20 DPS) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.81 DPS) | yes | Fiery Slippers (254005, -3.48 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -4.01 DPS) [crafted]; Spidersilk Boots (4320, -4.05 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.49 DPS) | yes | Reedknot Ring (9622, -0.80 DPS) [quest]; Snake Hoop (6750, -0.97 DPS) [quest]; Ring of Power Regulation (273030, -1.04 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (1.94 DPS) | yes | Reedknot Ring (9622, -0.24 DPS) [quest]; Snake Hoop (6750, -0.41 DPS) [quest]; Ring of Power Regulation (273030, -0.48 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (116.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.03 DPS) [dungeon]; Spine of the Devourer (273040, -2.05 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.39 DPS) | yes | Orb of Noh'Orahil (15107, -0.67 DPS, sim-verified) [quest]; Rod of Molten Fire (2565, -1.02 DPS) [world_drop]; Beacon of Hope (9393, -1.45 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 163.1 spell_power points (39.47 DPS) | yes | Earthen Rod (9381, -3.15 DPS) [dungeon]; Umbral Wand (5216, -3.80 DPS) [dungeon]; Twisted Nether Wand (249144, -4.02 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 193.7. Weights run: 3.5s. Verify run: 1.1s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.057 ± 0.018, crit=0.091 ± 0.004 per rating point (14 rating = 1%, 1.278 per %), hit=0.502 ± 0.024 per rating point (10 rating = 1%, 5.018 per %), spell_haste=-10.063 ± 0.531, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.258 ± 0.000, fire_power=0.742 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.06 DPS) | yes | Dreamweave Circlet (10041, -1.22 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.35 DPS) [crafted]; Red Mageweave Headband (10033, -1.54 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (1.79 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.37 DPS) [quest]; Horizon Choker (13085, -1.61 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 14.0 spell_power points (3.15 DPS) | yes | Netherflame Shoulders (254053, -0.22 DPS) [crafted]; Black Mageweave Shoulders (10027, -0.79 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -0.95 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (3.22 DPS) | yes | Cindercloth Cloak (14044, -0.95 DPS) [crafted]; Runecloth Cloak (13860, -1.10 DPS) [crafted]; Cloak of Fire (14134, -1.88 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (5.01 DPS) | yes | Acumen Robes (17775, -0.49 DPS) [quest]; Cindercloth Robe (10042, -0.52 DPS) [crafted]; Elemental Raiment (9434, -1.75 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.0 spell_power points (2.25 DPS) | yes | Arcane Runed Bracers (4744, -0.23 DPS) [quest]; Spidertank Oilrag (9448, -0.23 DPS) [dungeon]; Phoenix Bindings (210781, -0.74 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.09 DPS) | yes | Fiery Gloves (254099, -0.80 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.06 DPS) [vendor]; Black Mageweave Gloves (10003, -2.67 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 14.6 spell_power points (3.28 DPS) | yes | Satyrmane Sash (17755, -0.01 DPS) [dungeon]; Ghostweave Cord (254073, -0.14 DPS) [crafted]; Star Belt (4329, -0.36 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (4.26 DPS) | yes | Cindercloth Leggings (12256, -1.10 DPS) [vendor]; Knight's Dreadweave Leggings (220888, -1.13 DPS) [vendor]; Red Mageweave Pants (10009, -1.66 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.38 DPS) | yes | Fiery Sandals (254111, -1.94 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.35 DPS) [vendor]; Cindercloth Boots (10044, -2.62 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.92 DPS) | yes | Ring of Forlorn Spirits (2043, -1.12 DPS) [quest]; Reedknot Ring (9622, -1.35 DPS) [quest]; Snake Hoop (6750, -1.51 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.31 DPS) | yes | Reedknot Ring (9622, -0.74 DPS) [quest]; Snake Hoop (6750, -0.90 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.76 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.92 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.29 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.45 DPS) [dungeon]; Hanzo Sword (8190, -21.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (193.7 DPS) | yes | Wand of Allistarj (13065, -2.70 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.58 DPS) [crafted]; Pyric Caduceus (11748, -7.64 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Fiery Waistcord; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 362.0. Weights run: 1.5s. Verify run: 1.8s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.097 ± 0.022, crit=0.106 ± 0.003 per rating point (14 rating = 1%, 1.489 per %), hit=0.478 ± 0.024 per rating point (10 rating = 1%, 4.784 per %), spell_haste=-3.787 ± 0.447, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.322 ± 0.001, fire_power=0.678 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.15 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.04 DPS) [pvp]; Dreamweave Circlet (10041, -1.89 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.0 spell_power points (4.88 DPS) | yes | Triune Amulet (7722, -1.87 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.95 DPS) [dungeon]; Mindburst Medallion (11196, -2.40 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.40 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.97 DPS) [crafted]; Ghostweave Mantle (254127, -3.76 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.48 DPS) | yes | Mageflame Cloak (13007, +0.00 DPS, sim-verified) [world_drop]; Runecloth Cloak (13860, -1.81 DPS) [crafted]; Cindercloth Cloak (14044, -1.88 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.9 spell_power points (17.62 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.71 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.49 DPS) [pvp]; Robe of the Magi (1716, -9.03 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.0 spell_power points (4.50 DPS) | yes | Netherflame Cuffs (254065, -0.93 DPS) [crafted]; Arcane Runed Bracers (4744, -1.12 DPS) [quest]; Spidertank Oilrag (9448, -1.12 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.27 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.23 DPS) [pvp]; Dreamweave Gloves (10019, -1.36 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.0 spell_power points (5.63 DPS) | yes | Ghostweave Cord (254073, -0.37 DPS) [crafted]; Fiery Waistcord (254085, -0.46 DPS) [crafted]; Star Belt (4329, -0.74 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.77 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -2.63 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.02 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.75 DPS) [crafted]; Omnicast Boots (11822, -1.06 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.7 spell_power points (7.02 DPS) | yes | Philanthropist's Ring (281635, -3.08 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.02 DPS) [quest]; Reedknot Ring (9622, -4.39 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.89 DPS) | yes | Philanthropist's Ring (281635, -0.94 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.88 DPS) [quest]; Reedknot Ring (9622, -2.26 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -4.41 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.31 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -28.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (362.0 DPS) | yes | Stormrager (16997, -0.35 DPS) [quest]; Torch of Austen (13004, -6.43 DPS) [world_drop]; Torch of Light (279246, -10.31 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Brilliant Wand

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 05500000000000000-2035113111120001350-0050051000000000)

Set DPS (verified): 693.0. Weights run: 4.6s. Verify run: 1.9s. 1126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.051 ± 0.020, crit=0.168 ± 0.005 per rating point (14 rating = 1%, 2.352 per %), hit=0.669 ± 0.040 per rating point (10 rating = 1%, 6.686 per %), spell_haste=-10.878 ± 0.714, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.262 ± 0.001, fire_power=0.738 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.83 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.40 DPS) [pvp]; Dreamweave Circlet (10041, -2.81 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.5 spell_power points (6.41 DPS) | yes | Triune Amulet (7722, -2.31 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -2.61 DPS) [dungeon]; Mindburst Medallion (11196, -3.18 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (12.81 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.55 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -5.11 DPS) [pvp] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 15.5 spell_power points (7.94 DPS) | yes | Spritecaster Cape (11623, -0.62 DPS) [dungeon]; Cloak of Fire (14134, -2.65 DPS) [crafted]; Cindercloth Cloak (14044, -2.82 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (23.80 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.78 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.47 DPS) [pvp]; Robe of the Magi (1716, -13.39 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.5 spell_power points (5.90 DPS) | yes | Netherflame Cuffs (254065, -0.80 DPS) [crafted]; Phoenix Bindings (210781, -0.98 DPS) [crafted]; Arcane Runed Bracers (4744, -1.29 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (11.27 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.41 DPS) [pvp]; Dreamweave Gloves (10019, -1.94 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.5 spell_power points (7.43 DPS) | yes | Fiery Waistcord (254085, -0.01 DPS) [crafted]; Ghostweave Cord (254073, -0.26 DPS) [crafted]; Star Belt (4329, -0.77 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (13.32 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -3.17 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.29 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.02 DPS) [crafted]; Omnicast Boots (11822, -1.74 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.4 spell_power points (9.40 DPS) | yes | Philanthropist's Ring (281635, -4.15 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.30 DPS) [quest]; Reedknot Ring (9622, -5.82 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.66 DPS) | yes | Philanthropist's Ring (281635, -1.41 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.56 DPS) [quest]; Reedknot Ring (9622, -3.07 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.01 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.71 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -41.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (693.0 DPS) | yes | Stormrager (16997, -1.55 DPS) [quest]; Greater Eternal Wand (249237, -6.43 DPS) [crafted]; Torch of Light (279246, -8.34 DPS, sim-verified) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; back: Mageflame Cloak; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; ranged: Brilliant Wand

No-known-source sample (15 of 1126, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 32.5. Weights run: 1.6s. Verify run: 0.9s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.028 ± 0.002, crit=0.022 ± 0.001 per rating point (14 rating = 1%, 0.308 per %), hit=0.089 ± 0.003 per rating point (10 rating = 1%, 0.887 per %), spell_haste=0.927 ± 0.034, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.246 ± 0.000, fire_power=0.754 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Circlet (253953, -0.76 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.17 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.3 spell_power points (1.05 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.03 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Black Whelp Cloak (7283, -0.32 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.38 DPS) [crafted]; Mystic Shawl (3449, -0.40 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gray Woolen Robe (2585, -0.20 DPS) [crafted]; Green Woolen Vest (2582, -0.23 DPS) [crafted]; Filigreed Flame Gown (253905, -0.51 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.60 DPS) | yes | Windsong Bangles (263336, -0.40 DPS) [quest]; Owlbeard Bracers (16981, -0.43 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.57 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Phoenix Gloves (4331, -0.04 DPS) [crafted]; Gnoll Casting Gloves (892, -0.20 DPS) [world]; Flame Gloves (253917, -0.33 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.01 DPS) [crafted]; Novice Ardent's Sash (253887, -0.40 DPS) [crafted]; Flame Sash (253929, -0.49 DPS, sim-verified) [crafted] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 8.5 spell_power points (1.69 DPS) | yes | Silk-threaded Trousers (1929, -0.29 DPS) [dungeon]; Phoenix Pants (4317, -0.34 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.1 spell_power points (1.42 DPS) | yes | Flame Boots (253893, -0.57 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.62 DPS) [crafted]; Pristine Boots (253889, -0.80 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.2 spell_power points (0.03 DPS) | yes | Volcanic Rock Ring (12053, -0.02 DPS) [world_drop]; The 1 Ring (8350, -0.03 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 0.1 spell_power points (0.03 DPS) | yes | Volcanic Rock Ring (12053, -0.01 DPS) [world_drop]; The 1 Ring (8350, -0.02 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.3 spell_power points (0.06 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.02 DPS) [world]; Clear Crystal Rod (16894, -0.03 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.9 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.54 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.26 DPS) [quest]; Greater Magic Wand (11288, -5.01 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 65.9. Weights run: 1.5s. Verify run: 0.9s. 283 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.044 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.423 per %), hit=0.125 ± 0.004 per rating point (10 rating = 1%, 1.247 per %), spell_haste=-1.894 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.239 ± 0.000, fire_power=0.761 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.42 DPS) | yes | Silk Headband (7050, -0.44 DPS) [crafted]; Flame Circlet (253953, -0.58 DPS) [crafted]; Filigreed Flame Circlet (253979, -2.48 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.3 spell_power points (0.06 DPS) | yes | Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.02 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.4 spell_power points (2.07 DPS) | yes | Chestnut Mantle (17695, -0.37 DPS, sim-verified) [quest]; Invoker's Mantle (215365, -0.48 DPS) [crafted]; Death Speaker Mantle (6685, -0.64 DPS) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.10 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Black Whelp Cloak (7283, -0.44 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (2.86 DPS) | yes | Flame Gown (253965, -0.26 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.75 DPS) [crafted]; Beguiler Robes (7728, -0.76 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.9 spell_power points (2.18 DPS) | yes | Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Tabitha's Cuffs (251486, -1.52 DPS) [quest]; Owlbeard Bracers (16981, -1.94 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.54 DPS) | yes | Phoenix Gloves (4331, -0.03 DPS) [crafted]; Jutebraid Gloves (10654, -0.17 DPS) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.42 DPS) | yes | Belt of Arugal (6392, -0.41 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.66 DPS) [dungeon]; Invoker's Cord (215366, -0.83 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.64 DPS) | yes | Flame Leggings (253991, -0.23 DPS) [crafted]; Smoldering Pants (3073, -0.46 DPS) [world]; Filigreed Flame Leggings (253941, -0.74 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (65.9 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Glinteye Slippers (273024, -0.07 DPS) [dungeon]; Fiery Slippers (254005, -0.83 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.2 spell_power points (1.37 DPS) | yes | Solid Gold Band (279845, -0.49 DPS) [quest]; Electrocutioner Lagnut (9447, -0.71 DPS) [dungeon]; Runemender's Seal (273031, -0.71 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.32 DPS) | yes | Solid Gold Band (279845, +0.00 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon]; Runemender's Seal (273031, -0.66 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (7.92 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Staff of Soran'ruk (15109, -2.23 DPS) [quest]; Glimmering Staff (249392, -7.82 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 149.5 spell_power points (32.90 DPS) | yes | Wand of Mana Concentration (273043, -1.71 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -2.80 DPS) [crafted]; Scorching Wand (5213, -3.04 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Phoenix Bindings; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 114.9. Weights run: 1.3s. Verify run: 0.8s. 381 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.060 ± 0.006, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.582 per %), hit=0.187 ± 0.006 per rating point (10 rating = 1%, 1.868 per %), spell_haste=-6.930 ± 0.226, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.247 ± 0.000, fire_power=0.753 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.08 DPS) | yes | Living Cowl (5608, -1.91 DPS, sim-verified) [world]; Augural Shroud (2620, -2.27 DPS) [world]; Holy Shroud (2721, -2.42 DPS) [world_drop] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.0 spell_power points (1.94 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.83 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.85 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.31 DPS) | yes | Chestnut Mantle (17695, -0.37 DPS) [quest]; Green Silken Shoulders (7057, -0.39 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.43 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (1.69 DPS) | yes | Guardian Cloak (5965, -0.17 DPS) [crafted]; Hillman's Cloak (3719, -0.48 DPS) [crafted]; Long Silken Cloak (4326, -0.72 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.4 spell_power points (5.41 DPS) | yes | Cindercloth Robe (10042, -0.49 DPS) [crafted]; Elemental Raiment (9434, -0.67 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.92 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.8 spell_power points (2.37 DPS) | yes | Spidertank Oilrag (9448, -0.19 DPS) [dungeon]; Condor Bracers (15864, -0.67 DPS) [quest]; Radiant Silver Bracers (4545, -1.28 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.41 DPS) | yes | Black Mageweave Gloves (10003, -1.36 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.61 DPS) [crafted]; Fiery Handwraps (254025, -1.76 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.15 DPS) | yes | Fiery Cord (254041, -0.44 DPS, sim-verified) [crafted]; Warsong Sash (16975, -0.48 DPS) [quest]; Scorching Sash (4117, -0.78 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.7 spell_power points (3.56 DPS) | yes | Gaze Dreamer Pants (6903, -0.87 DPS, sim-verified) [dungeon]; Flame Leggings (253991, -0.91 DPS) [crafted]; Smoldering Pants (3073, -1.20 DPS) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.81 DPS) | yes | Fiery Slippers (254005, -3.54 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -4.01 DPS) [crafted]; Spidersilk Boots (4320, -4.05 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.49 DPS) | yes | Snake Hoop (6750, -0.97 DPS) [quest]; Ring of Power Regulation (273030, -1.04 DPS) [dungeon]; Solid Gold Band (279845, -1.52 DPS) [quest] |
| finger2 | Reedknot Ring (9622) | Jarl Needs a Blade [quest] | 7.0 spell_power points (1.69 DPS) | yes | Ring of Power Regulation (273030, -0.24 DPS) [dungeon]; Snake Hoop (6750, -0.57 DPS, sim-verified) [quest]; Solid Gold Band (279845, -0.73 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (114.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.03 DPS) [dungeon]; Spine of the Devourer (273040, -2.05 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.39 DPS) | yes | Rod of Molten Fire (2565, -1.02 DPS) [world_drop]; Orb of Noh'Orahil (15107, -1.15 DPS, sim-verified) [quest]; Omega Orb (7749, -1.21 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 163.1 spell_power points (39.47 DPS) | yes | Earthen Rod (9381, -0.50 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -3.80 DPS) [dungeon]; Twisted Nether Wand (249144, -4.02 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Reedknot Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 381, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 190.1. Weights run: 3.5s. Verify run: 1.1s. 492 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.057 ± 0.018, crit=0.091 ± 0.004 per rating point (14 rating = 1%, 1.278 per %), hit=0.502 ± 0.024 per rating point (10 rating = 1%, 5.018 per %), spell_haste=-10.063 ± 0.531, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.258 ± 0.000, fire_power=0.742 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Spellpower Goggles Xtreme (10502, -0.13 DPS) [crafted]; Red Mageweave Headband (10033, -0.32 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.33 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.25 DPS) [quest]; Horizon Choker (13085, -1.49 DPS) [world_drop]; Triune Amulet (7722, -4.69 DPS, sim-verified) [dungeon] |
| shoulder | Netherflame Shoulders (254053) | Tailoring [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -0.57 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -0.74 DPS) [vendor]; Rotgrip Mantle (17732, -1.88 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (3.22 DPS) | yes | Deep Woodlands Cloak (19121, -0.41 DPS) [quest]; Cloak of Fire (14134, -0.89 DPS) [crafted]; Cindercloth Cloak (14044, -0.95 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (5.01 DPS) | yes | Elemental Raiment (9434, +0.00 DPS, sim-verified) [world_drop]; Acumen Robes (17775, -0.49 DPS) [quest]; Cindercloth Robe (10042, -0.52 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Spidertank Oilrag (9448, -0.14 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.50 DPS) [crafted]; Netherflame Cuffs (254065, -2.02 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.09 DPS) | yes | Black Mageweave Gloves (10003, +0.00 DPS, sim-verified) [crafted]; Fiery Gloves (254099, -0.80 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.06 DPS) [vendor] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 14.6 spell_power points (3.28 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ghostweave Cord (254073, -0.14 DPS) [crafted]; Star Belt (4329, -0.36 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Cindercloth Leggings (12256, -0.13 DPS) [vendor]; Stone Guard's Dreadweave Leggings (220906, -0.16 DPS) [vendor]; Wizardweave Leggings (14132, -2.71 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.38 DPS) | yes | Cindercloth Boots (10044, -1.89 DPS) [crafted]; Fiery Sandals (254111, -1.94 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.35 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.92 DPS) | yes | Reedknot Ring (9622, -1.35 DPS) [quest]; Snake Hoop (6750, -1.51 DPS) [quest]; Ring of Power Regulation (273030, -1.57 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.31 DPS) | yes | Reedknot Ring (9622, +0.00 DPS, sim-verified) [quest]; Snake Hoop (6750, -0.90 DPS) [quest]; Ring of Power Regulation (273030, -0.96 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.23 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.70 DPS, sim-verified) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.29 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.45 DPS) [dungeon]; Hanzo Sword (8190, -17.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+10.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.70 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.58 DPS) [crafted]; Pyric Caduceus (11748, -10.05 DPS, sim-verified) [dungeon] |

**New at 50:** head: Dreamweave Circlet; neck: Glowing Eye of Mordresh; shoulder: Netherflame Shoulders; back: Spritecaster Cape; waist: Fiery Waistcord; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 492, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 355.2. Weights run: 1.5s. Verify run: 1.8s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.097 ± 0.022, crit=0.106 ± 0.003 per rating point (14 rating = 1%, 1.489 per %), hit=0.478 ± 0.024 per rating point (10 rating = 1%, 4.784 per %), spell_haste=-3.787 ± 0.447, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.322 ± 0.001, fire_power=0.678 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.15 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.04 DPS) [pvp]; Dreamweave Circlet (10041, -1.68 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.0 spell_power points (4.88 DPS) | yes | Glowing Eye of Mordresh (10769, -1.95 DPS) [dungeon]; Mindburst Medallion (11196, -2.40 DPS) [quest]; Triune Amulet (7722, -2.63 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.40 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.71 DPS, sim-verified) [crafted]; Ghostweave Mantle (254127, -3.76 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.48 DPS) | yes | Deep Woodlands Cloak (19121, -0.64 DPS) [quest]; Mageflame Cloak (13007, -1.44 DPS, sim-verified) [world_drop]; Runecloth Cloak (13860, -1.81 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.9 spell_power points (17.62 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.71 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.49 DPS) [pvp]; Robe of the Magi (1716, -8.40 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.0 spell_power points (4.50 DPS) | yes | Netherflame Cuffs (254065, -0.93 DPS) [crafted]; Spidertank Oilrag (9448, -1.12 DPS) [dungeon]; Phoenix Bindings (210781, -1.19 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.27 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.23 DPS) [pvp] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.0 spell_power points (5.63 DPS) | yes | Ghostweave Cord (254073, -0.37 DPS) [crafted]; Fiery Waistcord (254085, -0.46 DPS) [crafted]; Star Belt (4329, -0.74 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.77 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Brightcloth Pants (14104, -4.11 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.02 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.75 DPS) [crafted]; Omnicast Boots (11822, -1.06 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.7 spell_power points (7.02 DPS) | yes | Philanthropist's Ring (281635, -3.08 DPS) [quest]; Reedknot Ring (9622, -4.39 DPS) [quest]; Snake Hoop (6750, -4.58 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.89 DPS) | yes | Philanthropist's Ring (281635, -0.94 DPS) [quest]; Reedknot Ring (9622, -2.26 DPS) [quest]; Snake Hoop (6750, -2.45 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.31 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -31.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (355.2 DPS) | yes | Stormrager (16997, -0.35 DPS) [quest]; Torch of Austen (13004, -6.43 DPS) [world_drop]; Torch of Light (279246, -8.24 DPS, sim-verified) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Brilliant Wand

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 05500000000000000-2035113111120001350-0050051000000000)

Set DPS (verified): 688.1. Weights run: 4.6s. Verify run: 2.0s. 1118 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.051 ± 0.020, crit=0.168 ± 0.005 per rating point (14 rating = 1%, 2.352 per %), hit=0.669 ± 0.040 per rating point (10 rating = 1%, 6.686 per %), spell_haste=-10.878 ± 0.714, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.262 ± 0.001, fire_power=0.738 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.83 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.40 DPS) [pvp]; Dreamweave Circlet (10041, -2.72 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.5 spell_power points (6.41 DPS) | yes | Triune Amulet (7722, -2.35 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -2.61 DPS) [dungeon]; Mindburst Medallion (11196, -3.18 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (12.81 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.67 DPS, sim-verified) [crafted]; Champion's Dreadweave Spaulders (227092, -5.11 DPS) [pvp] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 15.5 spell_power points (7.94 DPS) | yes | Spritecaster Cape (11623, -0.62 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.56 DPS) [quest]; Cloak of Fire (14134, -2.65 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (23.80 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.78 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.47 DPS) [pvp]; Robe of the Magi (1716, -10.55 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.5 spell_power points (5.90 DPS) | yes | Netherflame Cuffs (254065, -0.80 DPS) [crafted]; Phoenix Bindings (210781, -0.98 DPS) [crafted]; Spidertank Oilrag (9448, -1.29 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (11.27 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.41 DPS) [pvp]; Dreamweave Gloves (10019, -2.70 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.5 spell_power points (7.43 DPS) | yes | Fiery Waistcord (254085, -0.01 DPS) [crafted]; Ghostweave Cord (254073, -0.26 DPS) [crafted]; Star Belt (4329, -0.77 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (13.32 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -4.61 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.29 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.02 DPS) [crafted]; Omnicast Boots (11822, -1.74 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.4 spell_power points (9.40 DPS) | yes | Philanthropist's Ring (281635, -4.15 DPS) [quest]; Reedknot Ring (9622, -5.82 DPS) [quest]; Snake Hoop (6750, -6.20 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.66 DPS) | yes | Philanthropist's Ring (281635, -1.41 DPS) [quest]; Reedknot Ring (9622, -3.07 DPS) [quest]; Snake Hoop (6750, -3.46 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.01 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.71 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -40.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (688.1 DPS) | yes | Stormrager (16997, -1.55 DPS) [quest]; Greater Eternal Wand (249237, -6.43 DPS) [crafted]; Torch of Light (279246, -7.89 DPS, sim-verified) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Mageflame Cloak; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; ranged: Brilliant Wand

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

