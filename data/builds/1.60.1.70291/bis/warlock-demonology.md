# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 34.5. Weights run: 1.7s. Verify run: 1.8s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.022 ± 0.003, crit=0.024 ± 0.001 per rating point (14 rating = 1%, 0.338 per %), hit=0.094 ± 0.004 per rating point (10 rating = 1%, 0.943 per %), spell_haste=1.932 ± 0.068, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.246 ± 0.000, fire_power=0.754 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Circlet (253953, -0.77 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.17 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.2 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.24 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Feyscale Cloak (6632, -0.20 DPS) [dungeon]; Caretaker's Cape (20428, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.24 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gray Woolen Robe (2585, -0.20 DPS) [crafted]; Green Woolen Vest (2582, -0.22 DPS) [crafted]; Filigreed Flame Gown (253905, -0.52 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.18 DPS) [dungeon]; Bright Bracers (3647, -0.18 DPS) [world_drop]; Repurposed Hair Band (281256, -0.19 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.39 DPS) | yes | Phoenix Gloves (4331, -0.04 DPS) [crafted]; Gnoll Casting Gloves (892, -0.20 DPS) [world]; Flame Gloves (253917, -0.33 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.00 DPS) [crafted]; Novice Ardent's Sash (253887, -0.40 DPS) [crafted]; Flame Sash (253929, -0.49 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.2 spell_power points (1.83 DPS) | yes | Filigreed Flame Leggings (253941, -0.15 DPS) [crafted]; Phoenix Pants (4317, -0.31 DPS) [crafted]; Silk-threaded Trousers (1929, -0.43 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.1 spell_power points (1.41 DPS) | yes | Feather Padded Treads (285345, -0.41 DPS, sim-verified) [world]; Flame Boots (253893, -0.50 DPS) [crafted]; Red Woolen Boots (4313, -0.61 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.00 DPS) | yes | Sludge-Stained Band (286535, -0.41 DPS) [world]; Lavishly Jeweled Ring (1156, -0.98 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.99 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.00 DPS) | yes | Sludge-Stained Band (286535, -0.39 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.97 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.98 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.2 spell_power points (0.04 DPS) | yes | Channeler's Staff (4437, -0.01 DPS) [world]; Lesser Staff of the Spire (1300, -0.02 DPS) [world]; Staff of Westfall (2042, -0.02 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 115.1 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.40 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.62 DPS) [dungeon]; Sizzle Stick (8071, -4.26 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 70.3. Weights run: 1.7s. Verify run: 1.8s. 260 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.044 ± 0.004, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.431 per %), hit=0.124 ± 0.006 per rating point (10 rating = 1%, 1.243 per %), spell_haste=-1.939 ± 0.121, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.239 ± 0.000, fire_power=0.761 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.41 DPS) | yes | Silk Headband (7050, -0.44 DPS) [crafted]; Flame Circlet (253953, -0.58 DPS) [crafted]; Filigreed Flame Circlet (253979, -2.87 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 spell_power points (1.59 DPS) | yes | Crystal Starfire Medallion (5003, -1.55 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.55 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.71 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.4 spell_power points (2.06 DPS) | yes | Moonlit Amice (11884, -0.53 DPS) [quest]; Death Speaker Mantle (6685, -0.64 DPS) [dungeon]; Invoker's Mantle (215365, -0.79 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (1.32 DPS) | yes | Hillman's Cloak (3719, -0.22 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.44 DPS) [crafted]; Prelacy Cape (7004, -0.44 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.85 DPS) | yes | Flame Gown (253965, -0.38 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.75 DPS) [crafted]; Beguiler Robes (7728, -0.80 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.9 spell_power points (2.17 DPS) | yes | Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Windsong Bangles (263336, -1.95 DPS) [quest]; Glowing Magical Bracelets (13106, -2.09 DPS) [world_drop] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.54 DPS) | yes | Phoenix Gloves (4331, -0.03 DPS) [crafted]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Serpent Gloves (5970, -0.29 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.1 spell_power points (2.44 DPS) | yes | Belt of Arugal (6392, -0.44 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.69 DPS) [dungeon]; Invoker's Cord (215366, -0.86 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.63 DPS) | yes | Flame Leggings (253991, -0.23 DPS) [crafted]; Smoldering Pants (3073, -0.46 DPS) [world]; Abomination Skin Leggings (23173, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (70.3 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Nimbus Boots (6998, -0.29 DPS) [quest]; Fiery Slippers (254005, -0.84 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.54 DPS) | yes | Minor Channeling Ring (1449, -0.42 DPS) [quest]; Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon]; Sludge-Stained Band (286535, -0.88 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.32 DPS) | yes | Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon]; Sludge-Stained Band (286535, -0.66 DPS) [world]; Minor Channeling Ring (1449, -1.56 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | 30.0 spell_power points (6.58 DPS) | yes | Wind Spirit Staff (6689, -1.49 DPS) [dungeon]; Staff of Soran'ruk (15109, -3.86 DPS, sim-verified) [quest]; Glimmering Staff (249392, -6.47 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 154.4 spell_power points (33.87 DPS) | yes | Starfaller (13063, -1.04 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.77 DPS) [crafted]; Scorching Wand (5213, -4.02 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Phoenix Bindings; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff; ranged: Necrotic Wand

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 114.8. Weights run: 1.6s. Verify run: 1.7s. 334 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.057 ± 0.008, crit=0.044 ± 0.002 per rating point (14 rating = 1%, 0.616 per %), hit=0.190 ± 0.009 per rating point (10 rating = 1%, 1.903 per %), spell_haste=-6.909 ± 0.316, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.248 ± 0.000, fire_power=0.752 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.95 DPS) | yes | Living Cowl (5608, -1.91 DPS, sim-verified) [world]; Augural Shroud (2620, -2.22 DPS) [world]; Holy Shroud (2721, -2.36 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.5 spell_power points (1.76 DPS) | yes | Scorn's Icy Choker (23169, -1.23 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.66 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.66 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.24 DPS) | yes | Green Silken Shoulders (7057, -0.21 DPS) [crafted]; Inquisitor's Shawl (19507, -0.42 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (1.65 DPS) | yes | Guardian Cloak (5965, -0.17 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.24 DPS) [quest]; Long Silken Cloak (4326, -0.54 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (5.27 DPS) | yes | Elemental Raiment (9434, -0.45 DPS, sim-verified) [world_drop]; Cindercloth Robe (10042, -0.48 DPS) [crafted]; Dreamweave Vest (10021, -0.90 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.8 spell_power points (2.30 DPS) | yes | Arcane Runed Bracers (4744, -0.18 DPS) [quest]; Spidertank Oilrag (9448, -0.18 DPS) [dungeon]; Condor Bracers (15864, -0.65 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.30 DPS) | yes | Black Mageweave Gloves (10003, -1.19 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.57 DPS) [crafted]; Fiery Handwraps (254025, -1.72 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.2 spell_power points (3.35 DPS) | yes | Star Belt (4329, -0.74 DPS, sim-verified) [crafted]; Fiery Cord (254041, -0.76 DPS) [crafted]; Scorching Sash (4117, -1.05 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.7 spell_power points (3.46 DPS) | yes | Flame Leggings (253991, -0.89 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.02 DPS, sim-verified) [dungeon]; Smoldering Pants (3073, -1.16 DPS) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.66 DPS) | yes | Fiery Slippers (254005, -3.54 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.91 DPS) [crafted]; Spidersilk Boots (4320, -3.95 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.42 DPS) | yes | Ring of Forlorn Spirits (2043, -0.54 DPS) [quest]; Reedknot Ring (9622, -0.77 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.01 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.12 DPS) | yes | Reedknot Ring (9622, -0.47 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.56 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.71 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (114.8 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.03 DPS) [dungeon]; Staff of Noh'Orahil (15105, -3.37 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.30 DPS) | yes | Thrash's Trash (276204, +0.00 DPS, sim-verified) [vendor]; Orb of Noh'Orahil (15107, -0.75 DPS) [quest]; Rod of Molten Fire (2565, -1.00 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 167.4 spell_power points (39.46 DPS) | yes | Umbral Wand (5216, -3.79 DPS) [dungeon]; Twisted Nether Wand (249144, -4.05 DPS) [crafted]; Burning Sliver (5249, -5.36 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 334, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 189.1. Weights run: 3.5s. Verify run: 2.1s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.049 ± 0.024, crit=0.091 ± 0.006 per rating point (14 rating = 1%, 1.276 per %), hit=0.492 ± 0.032 per rating point (10 rating = 1%, 4.918 per %), spell_haste=-11.670 ± 0.748, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.256 ± 0.000, fire_power=0.744 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (5.95 DPS) | yes | Dreamweave Circlet (10041, +0.00 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.32 DPS) [crafted]; Red Mageweave Headband (10033, -1.55 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.4 spell_power points (1.63 DPS) | yes | Mindburst Medallion (11196, -0.24 DPS) [quest]; Scorn's Icy Choker (23169, -1.16 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.48 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.9 spell_power points (3.06 DPS) | yes | Netherflame Shoulders (254053, +0.00 DPS, sim-verified) [crafted]; Black Mageweave Shoulders (10027, -0.76 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -0.92 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (3.15 DPS) | yes | Cindercloth Cloak (14044, -0.93 DPS) [crafted]; Runecloth Cloak (13860, -1.08 DPS) [crafted]; Cloak of Fire (14134, -1.17 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (4.91 DPS) | yes | Elemental Raiment (9434, -0.28 DPS) [world_drop]; Cindercloth Robe (10042, -0.49 DPS) [crafted]; Acumen Robes (17775, -0.51 DPS) [quest] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.0 spell_power points (2.21 DPS) | yes | Arcane Runed Bracers (4744, -0.22 DPS) [quest]; Spidertank Oilrag (9448, -0.22 DPS) [dungeon]; Phoenix Bindings (210781, -0.81 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.01 DPS) | yes | Black Mageweave Gloves (10003, -0.70 DPS) [crafted]; Fiery Gloves (254099, -0.79 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.05 DPS) [vendor] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 14.6 spell_power points (3.21 DPS) | yes | Highlander's Cloth Girdle (20098, -0.08 DPS) [rep]; Ghostweave Cord (254073, -0.12 DPS) [crafted]; Satyrmane Sash (17755, -1.71 DPS, sim-verified) [dungeon] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (4.19 DPS) | yes | Red Mageweave Pants (10009, -0.97 DPS) [crafted]; Cindercloth Leggings (12256, -1.07 DPS) [vendor]; Knight's Dreadweave Leggings (220888, -1.13 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.29 DPS) | yes | Cindercloth Boots (10044, -1.83 DPS, sim-verified) [crafted]; Fiery Sandals (254111, -1.92 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.35 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.87 DPS) | yes | Philanthropist's Ring (281635, -0.61 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.10 DPS) [quest]; Reedknot Ring (9622, -1.32 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (2.65 DPS) | yes | Ring of Forlorn Spirits (2043, -0.88 DPS) [quest]; Reedknot Ring (9622, -1.10 DPS) [quest]; Philanthropist's Ring (281635, -2.84 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.32 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.31 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.44 DPS) [dungeon]; Hanzo Sword (8190, -19.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (189.1 DPS) | yes | Wand of Allistarj (13065, -2.68 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.59 DPS) [crafted]; Pyric Caduceus (11748, -7.53 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Fiery Waistcord; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (gnome, 25532000000000000-2332113101220001350-0040000000000000)

Set DPS (verified): 396.3. Weights run: 3.5s. Verify run: 4.3s. 867 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.099 ± 0.025, crit=0.103 ± 0.005 per rating point (14 rating = 1%, 1.440 per %), hit=0.564 ± 0.041 per rating point (10 rating = 1%, 5.636 per %), spell_haste=-7.699 ± 0.738, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.345 ± 0.001, fire_power=0.655 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.64 DPS) | yes | Deathmist Mask (226909, +0.00 DPS, sim-verified) [quest]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.99 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.85 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.04 DPS) [quest]; Kezan's Taint (19604, -2.57 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.9 spell_power points (10.68 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.16 DPS) [pvp]; Argent Shoulders (19059, -1.76 DPS) [crafted]; Mantle of the Timbermaw (19050, -3.64 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.4 spell_power points (8.00 DPS) | yes | Hide of the Wild (18510, -2.65 DPS) [crafted]; Spritecaster Cape (11623, -2.80 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -3.01 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.9 spell_power points (16.73 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.47 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.11 DPS) [pvp]; Robe of the Magi (1716, -8.72 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.8 spell_power points (8.13 DPS) | yes | Runecloth Cuffs (254123, -3.85 DPS) [crafted]; Netherflame Cuffs (254065, -4.85 DPS) [crafted]; Arcane Runed Bracers (4744, -4.92 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.5 spell_power points (9.81 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.96 DPS) [quest]; Inferno Gloves (18408, -4.44 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.8 spell_power points (11.71 DPS) | yes | Belt of the Archmage (18405, -4.90 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -4.93 DPS) [rep]; Highlander's Cloth Girdle (20047, -5.99 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 35.9 spell_power points (12.80 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -2.35 DPS) [pvp]; Ghostweave Pants (14144, -3.52 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.56 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.71 DPS) [crafted]; Omnicast Boots (11822, -1.00 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 34.5 spell_power points (12.32 DPS) | yes | Elemental Focus Band (20682, -4.08 DPS) [world]; Lorekeeper's Ring (19522, -7.32 DPS) [rep]; Band of the Unicorn (7553, -7.68 DPS) [world_drop] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (396.3 DPS) | yes | Lorekeeper's Ring (19522, -0.58 DPS) [rep]; Band of the Unicorn (7553, -0.94 DPS) [world_drop]; Elemental Focus Band (20682, -5.25 DPS, sim-verified) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+18.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -0.71 DPS) [quest]; Weakness Analyzer (272438, -2.16 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.85 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.29 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -29.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 221.1 spell_power points (78.90 DPS) | yes | Wand of Biting Cold (19108, -4.99 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.86 DPS) [crafted]; Stormrager (16997, -16.03 DPS) [quest] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 867, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60, raid preset (gnome, 05500000000000000-2035113111120001350-0050051000000000)

Set DPS (verified): 754.4. Weights run: 3.6s. Verify run: 4.3s. 867 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.051 ± 0.028, crit=0.168 ± 0.007 per rating point (14 rating = 1%, 2.355 per %), hit=0.667 ± 0.056 per rating point (10 rating = 1%, 6.669 per %), spell_haste=-11.062 ± 0.977, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.261 ± 0.001, fire_power=0.739 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.42 DPS) | yes | Deathmist Mask (226909, +0.00 DPS, sim-verified) [quest]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.35 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.93 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -3.15 DPS) [quest]; Kezan's Taint (19604, -3.77 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.1 spell_power points (14.97 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.11 DPS) [pvp]; Mantle of the Timbermaw (19050, -5.02 DPS) [crafted]; Argent Shoulders (19059, -5.31 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.1 spell_power points (11.47 DPS) | yes | Hide of the Wild (18510, -4.26 DPS) [crafted]; Spritecaster Cape (11623, -4.36 DPS) [dungeon]; Mageflame Cloak (13007, -8.04 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (23.09 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.57 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.16 DPS) [pvp]; Robe of the Magi (1716, -15.26 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.4 spell_power points (11.14 DPS) | yes | Runecloth Cuffs (254123, -5.42 DPS) [crafted]; Netherflame Cuffs (254065, -6.19 DPS) [crafted]; Phoenix Bindings (210781, -6.37 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (13.55 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.20 DPS) [crafted]; Earth Warder's Gloves (21318, -2.61 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.8 spell_power points (16.30 DPS) | yes | Belt of the Archmage (18405, -6.30 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -7.10 DPS) [rep]; Highlander's Cloth Girdle (20047, -8.02 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 37.2 spell_power points (18.50 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -4.25 DPS) [pvp]; Ghostweave Pants (14144, -5.58 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.93 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.99 DPS) [crafted]; Omnicast Boots (11822, -1.68 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.1 spell_power points (17.46 DPS) | yes | Blessed Band of Light (272407, -9.18 DPS) [vendor]; Lorekeeper's Ring (19522, -10.50 DPS) [rep]; Band of the Unicorn (7553, -11.00 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.6 spell_power points (11.21 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Lorekeeper's Ring (19522, -4.26 DPS) [rep]; Band of the Unicorn (7553, -4.75 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (754.4 DPS) | yes | Abyss Shard (20534, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (754.4 DPS) | yes | Abyss Shard (20534, -0.99 DPS) [quest]; Weakness Analyzer (272438, -3.10 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -3.97 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (754.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.69 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -45.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 160.5 spell_power points (79.74 DPS) | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Wand of Biting Cold (19108, -15.74 DPS) [quest]; Stormrager (16997, -16.92 DPS) [quest] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 867, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 34.1. Weights run: 1.7s. Verify run: 1.7s. 137 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.022 ± 0.003, crit=0.024 ± 0.001 per rating point (14 rating = 1%, 0.338 per %), hit=0.094 ± 0.004 per rating point (10 rating = 1%, 0.943 per %), spell_haste=1.932 ± 0.068, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.246 ± 0.000, fire_power=0.754 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Circlet (253953, -0.75 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.17 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.2 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.02 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.24 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Feyscale Cloak (6632, -0.20 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.26 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gray Woolen Robe (2585, -0.20 DPS) [crafted]; Green Woolen Vest (2582, -0.22 DPS) [crafted]; Filigreed Flame Gown (253905, -0.51 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.60 DPS) | yes | Windsong Bangles (263336, -0.40 DPS) [quest]; Owlbeard Bracers (16981, -0.47 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.58 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.39 DPS) | yes | Phoenix Gloves (4331, -0.04 DPS) [crafted]; Gnoll Casting Gloves (892, -0.20 DPS) [world]; Flame Gloves (253917, -0.33 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.00 DPS) [crafted]; Novice Ardent's Sash (253887, -0.40 DPS) [crafted]; Flame Sash (253929, -0.48 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.2 spell_power points (1.83 DPS) | yes | Filigreed Flame Leggings (253941, -0.15 DPS) [crafted]; Phoenix Pants (4317, -0.31 DPS) [crafted]; Silk-threaded Trousers (1929, -0.43 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.1 spell_power points (1.41 DPS) | yes | Flame Boots (253893, -0.50 DPS) [crafted]; Feather Padded Treads (285345, -0.53 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.61 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.97 DPS) [dungeon]; Loop of Sacrifice (281673, -0.97 DPS) [quest]; Volcanic Rock Ring (12053, -0.98 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.60 DPS) | yes | Loop of Sacrifice (281673, -0.58 DPS) [quest]; Volcanic Rock Ring (12053, -0.58 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.64 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.2 spell_power points (0.04 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.01 DPS) [world]; Lesser Staff of the Spire (1300, -0.02 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 115.1 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.48 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.62 DPS) [dungeon]; Sizzle Stick (8071, -4.26 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 137, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 70.4. Weights run: 1.7s. Verify run: 1.7s. 243 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.044 ± 0.004, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.431 per %), hit=0.124 ± 0.006 per rating point (10 rating = 1%, 1.243 per %), spell_haste=-1.939 ± 0.121, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.239 ± 0.000, fire_power=0.761 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.41 DPS) | yes | Silk Headband (7050, -0.44 DPS) [crafted]; Flame Circlet (253953, -0.58 DPS) [crafted]; Filigreed Flame Circlet (253979, -2.83 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 spell_power points (1.59 DPS) | yes | Darkspear Warding Pendant (272075, -1.46 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.55 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.55 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.4 spell_power points (2.06 DPS) | yes | Invoker's Mantle (215365, -0.48 DPS) [crafted]; Death Speaker Mantle (6685, -0.64 DPS) [dungeon]; Chestnut Mantle (17695, -0.74 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.10 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Battle Healer's Cloak (19529, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.85 DPS) | yes | Flame Gown (253965, -0.60 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.75 DPS) [crafted]; Beguiler Robes (7728, -0.80 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.9 spell_power points (2.17 DPS) | yes | Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Tabitha's Cuffs (251486, -1.51 DPS) [quest]; Owlbeard Bracers (16981, -1.93 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.54 DPS) | yes | Phoenix Gloves (4331, -0.03 DPS) [crafted]; Jutebraid Gloves (10654, -0.17 DPS) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.1 spell_power points (2.44 DPS) | yes | Warsong Sash (16975, -0.31 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.44 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.69 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.63 DPS) | yes | Flame Leggings (253991, -0.38 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.46 DPS) [world]; Abomination Skin Leggings (23173, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Acidic Walkers (9454, -0.43 DPS) [dungeon]; Fiery Slippers (254005, -0.81 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.54 DPS) | yes | Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon]; Sludge-Stained Band (286535, -0.88 DPS) [world]; Sacred Band (6669, -1.10 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.32 DPS) | yes | Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon]; Sacred Band (6669, -0.88 DPS) [quest]; Sludge-Stained Band (286535, -1.84 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of Soran'ruk (15109, -4.97 DPS, sim-verified) [quest]; Glimmering Staff (249392, -4.99 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -5.00 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Starfaller (13063, -0.97 DPS) [world_drop]; Unstable Power Core (279847, -2.46 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.77 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Phoenix Bindings; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 243, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 113.1. Weights run: 1.6s. Verify run: 1.6s. 313 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.057 ± 0.008, crit=0.044 ± 0.002 per rating point (14 rating = 1%, 0.616 per %), hit=0.190 ± 0.009 per rating point (10 rating = 1%, 1.903 per %), spell_haste=-6.909 ± 0.316, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.248 ± 0.000, fire_power=0.752 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.95 DPS) | yes | Living Cowl (5608, -1.86 DPS, sim-verified) [world]; Augural Shroud (2620, -2.22 DPS) [world]; Holy Shroud (2721, -2.36 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.5 spell_power points (1.76 DPS) | yes | Scorn's Icy Choker (23169, -1.15 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.66 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.66 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.24 DPS) | yes | Green Silken Shoulders (7057, -0.21 DPS) [crafted]; Chestnut Mantle (17695, -0.36 DPS) [quest]; Inquisitor's Shawl (19507, -0.42 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (1.65 DPS) | yes | Long Silken Cloak (4326, -0.17 DPS) [crafted]; Guardian Cloak (5965, -0.17 DPS) [crafted]; Battle Healer's Cloak (19528, -0.24 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (5.27 DPS) | yes | Elemental Raiment (9434, -0.38 DPS, sim-verified) [world_drop]; Cindercloth Robe (10042, -0.48 DPS) [crafted]; Dreamweave Vest (10021, -0.90 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.8 spell_power points (2.30 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.18 DPS) [dungeon]; Condor Bracers (15864, -0.65 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.30 DPS) | yes | Black Mageweave Gloves (10003, -1.13 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.57 DPS) [crafted]; Fiery Handwraps (254025, -1.72 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.2 spell_power points (3.35 DPS) | yes | Star Belt (4329, -0.40 DPS, sim-verified) [crafted]; Warsong Sash (16975, -0.76 DPS) [quest]; Fiery Cord (254041, -0.76 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.7 spell_power points (3.46 DPS) | yes | Gaze Dreamer Pants (6903, -0.57 DPS, sim-verified) [dungeon]; Flame Leggings (253991, -0.89 DPS) [crafted]; Smoldering Pants (3073, -1.16 DPS) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.66 DPS) | yes | Fiery Slippers (254005, -3.22 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.91 DPS) [crafted]; Spidersilk Boots (4320, -3.95 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.42 DPS) | yes | Reedknot Ring (9622, -0.77 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.01 DPS) [vendor]; Sludge-Stained Band (286535, -1.72 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.12 DPS) | yes | Reedknot Ring (9622, -0.65 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.71 DPS) [vendor]; Sludge-Stained Band (286535, -1.41 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (113.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.03 DPS) [dungeon]; Staff of Noh'Orahil (15105, -3.37 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.30 DPS) | yes | Thrash's Trash (276204, +0.00 DPS, sim-verified) [vendor]; Orb of Noh'Orahil (15107, -0.75 DPS) [quest]; Rod of Molten Fire (2565, -1.00 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 167.4 spell_power points (39.46 DPS) | yes | Umbral Wand (5216, -3.79 DPS) [dungeon]; Twisted Nether Wand (249144, -4.05 DPS) [crafted]; Unstable Power Core (279847, -5.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 186.9. Weights run: 3.5s. Verify run: 2.0s. 392 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.049 ± 0.024, crit=0.091 ± 0.006 per rating point (14 rating = 1%, 1.276 per %), hit=0.492 ± 0.032 per rating point (10 rating = 1%, 4.918 per %), spell_haste=-11.670 ± 0.748, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.256 ± 0.000, fire_power=0.744 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (5.95 DPS) | yes | Dreamweave Circlet (10041, -0.89 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.32 DPS) [crafted]; Red Mageweave Headband (10033, -1.55 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.4 spell_power points (1.63 DPS) | yes | Mindburst Medallion (11196, -0.24 DPS) [quest]; Scorn's Icy Choker (23169, -0.86 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.48 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 13.9 spell_power points (3.06 DPS) | yes | Netherflame Shoulders (254053, -0.19 DPS) [crafted]; Black Mageweave Shoulders (10027, -0.76 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -0.92 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (3.15 DPS) | yes | Cloak of Fire (14134, -0.86 DPS) [crafted]; Cindercloth Cloak (14044, -0.93 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.95 DPS, sim-verified) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (4.91 DPS) | yes | Elemental Raiment (9434, -0.28 DPS) [world_drop]; Cindercloth Robe (10042, -0.49 DPS) [crafted]; Acumen Robes (17775, -0.51 DPS) [quest] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.0 spell_power points (2.21 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Phoenix Bindings (210781, -0.07 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.01 DPS) | yes | Fiery Gloves (254099, -0.79 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.05 DPS) [vendor]; Black Mageweave Gloves (10003, -1.65 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 14.6 spell_power points (3.21 DPS) | yes | Satyrmane Sash (17755, -0.02 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.08 DPS) [rep]; Ghostweave Cord (254073, -0.12 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (4.19 DPS) | yes | Cindercloth Leggings (12256, -1.07 DPS) [vendor]; Stone Guard's Dreadweave Leggings (220906, -1.13 DPS) [vendor]; Red Mageweave Pants (10009, -1.90 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.29 DPS) | yes | Cindercloth Boots (10044, -1.75 DPS, sim-verified) [crafted]; Fiery Sandals (254111, -1.92 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.35 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.87 DPS) | yes | Philanthropist's Ring (281635, -0.61 DPS) [quest]; Reedknot Ring (9622, -1.32 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.54 DPS) [vendor] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (2.65 DPS) | yes | Philanthropist's Ring (281635, -0.89 DPS, sim-verified) [quest]; Reedknot Ring (9622, -1.10 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.32 DPS) [vendor] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.32 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.89 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.22 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.31 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.44 DPS) [dungeon]; Hanzo Sword (8190, -19.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (186.9 DPS) | yes | Wand of Allistarj (13065, -2.68 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.59 DPS) [crafted]; Pyric Caduceus (11748, -6.39 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Fiery Waistcord; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25532000000000000-2332113101220001350-0040000000000000)

Set DPS (verified): 384.7. Weights run: 3.5s. Verify run: 4.4s. 857 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.099 ± 0.025, crit=0.103 ± 0.005 per rating point (14 rating = 1%, 1.440 per %), hit=0.564 ± 0.041 per rating point (10 rating = 1%, 5.636 per %), spell_haste=-7.699 ± 0.738, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.345 ± 0.001, fire_power=0.655 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.64 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -0.35 DPS) [quest]; Champion's Dreadweave Cowl (227090, -0.99 DPS) [pvp] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.85 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.04 DPS) [quest]; Kezan's Taint (19604, -2.57 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.9 spell_power points (10.68 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.16 DPS) [pvp]; Argent Shoulders (19059, -3.17 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -3.64 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.4 spell_power points (8.00 DPS) | yes | Hide of the Wild (18510, -2.65 DPS) [crafted]; Spritecaster Cape (11623, -2.80 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -3.01 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.9 spell_power points (16.73 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.47 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.11 DPS) [pvp]; Robe of the Magi (1716, -8.67 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.8 spell_power points (8.13 DPS) | yes | Runecloth Cuffs (254123, -3.85 DPS) [crafted]; Netherflame Cuffs (254065, -4.85 DPS) [crafted]; Spidertank Oilrag (9448, -4.92 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.5 spell_power points (9.81 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.96 DPS, sim-verified) [crafted]; Earth Warder's Gloves (21318, -1.96 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.8 spell_power points (11.71 DPS) | yes | Frostwolf Cloth Belt (19090, -4.93 DPS) [rep]; Belt of the Archmage (18405, -5.47 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -5.99 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 35.9 spell_power points (12.80 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -2.35 DPS) [pvp]; Outrider's Silk Leggings (22747, -3.58 DPS, sim-verified) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.56 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.71 DPS) [crafted]; Omnicast Boots (11822, -1.00 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 34.5 spell_power points (12.32 DPS) | yes | Blessed Band of Light (272407, -6.74 DPS) [vendor]; Advisor's Ring (19518, -7.32 DPS) [rep]; Band of the Unicorn (7553, -7.68 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 23.1 spell_power points (8.24 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Advisor's Ring (19518, -3.24 DPS) [rep]; Band of the Unicorn (7553, -3.60 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (384.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Abyss Shard (20534, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (384.7 DPS) | yes | Abyss Shard (20534, -0.71 DPS) [quest]; Weakness Analyzer (272438, -2.12 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.85 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (384.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.29 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -30.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 221.1 spell_power points (78.90 DPS) | yes | Wand of Biting Cold (19108, -3.03 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.86 DPS) [crafted]; Stormrager (16997, -16.03 DPS) [quest] |

**New at 60:** neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 857, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 05500000000000000-2035113111120001350-0050051000000000)

Set DPS (verified): 761.2. Weights run: 3.6s. Verify run: 4.2s. 857 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.051 ± 0.028, crit=0.168 ± 0.007 per rating point (14 rating = 1%, 2.355 per %), hit=0.667 ± 0.056 per rating point (10 rating = 1%, 6.669 per %), spell_haste=-11.062 ± 0.977, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.261 ± 0.001, fire_power=0.739 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | sim-verified (+9.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.81 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -9.02 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.93 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS, sim-verified) [quest]; Amulet of the Dawn (22657, -3.15 DPS) [quest]; Kezan's Taint (19604, -3.77 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.1 spell_power points (14.97 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.11 DPS) [pvp]; Argent Shoulders (19059, -2.55 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.02 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.1 spell_power points (11.47 DPS) | yes | Mageflame Cloak (13007, -3.04 DPS, sim-verified) [world_drop]; Hide of the Wild (18510, -4.26 DPS) [crafted]; Spritecaster Cape (11623, -4.36 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (23.09 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.57 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.16 DPS) [pvp]; Robe of the Magi (1716, -11.12 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.4 spell_power points (11.14 DPS) | yes | Runecloth Cuffs (254123, -5.42 DPS) [crafted]; Netherflame Cuffs (254065, -6.19 DPS) [crafted]; Phoenix Bindings (210781, -6.37 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (13.55 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.20 DPS) [crafted]; Earth Warder's Gloves (21318, -2.61 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.8 spell_power points (16.30 DPS) | yes | Belt of the Archmage (18405, -4.78 DPS) [crafted]; Frostwolf Cloth Belt (19090, -7.10 DPS) [rep]; Defiler's Cloth Girdle (20163, -8.02 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 37.2 spell_power points (18.50 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -4.10 DPS) [rep]; Legionnaire's Dreadweave Legguards (227097, -4.25 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.93 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.99 DPS) [crafted]; Omnicast Boots (11822, -1.68 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 35.1 spell_power points (17.46 DPS) | yes | Blessed Band of Light (272407, -9.18 DPS) [vendor]; Advisor's Ring (19518, -10.50 DPS) [rep]; Band of the Unicorn (7553, -11.00 DPS) [world_drop] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.6 spell_power points (11.21 DPS) | yes | Blessed Band of Light (272407, +0.00 DPS, sim-verified) [vendor]; Advisor's Ring (19518, -4.26 DPS) [rep]; Band of the Unicorn (7553, -4.75 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+21.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Abyss Shard (20534, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Abyss Shard (20534, -0.99 DPS) [quest]; Weakness Analyzer (272438, -3.04 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -3.97 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.69 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -41.09 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (+15.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Biting Cold (19108, -0.23 DPS) [quest]; Stormrager (16997, -1.41 DPS) [quest]; Torch of Light (279246, -15.09 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 857, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

