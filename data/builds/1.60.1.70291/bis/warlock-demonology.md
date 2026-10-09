# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 34.5. Weights run: 1.5s. Verify run: 1.0s. 151 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.019 ± 0.002, crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.324 per %), hit=0.092 ± 0.003 per rating point (10 rating = 1%, 0.922 per %), spell_haste=1.980 ± 0.050, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.246 ± 0.000, fire_power=0.754 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Circlet (253953, -0.77 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.18 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.2 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Feyscale Cloak (6632, -0.20 DPS) [dungeon]; Caretaker's Cape (20428, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.24 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gray Woolen Robe (2585, -0.20 DPS) [crafted]; Green Woolen Vest (2582, -0.22 DPS) [crafted]; Filigreed Flame Gown (253905, -0.52 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.18 DPS) [dungeon]; Bright Bracers (3647, -0.18 DPS) [world_drop]; Repurposed Hair Band (281256, -0.19 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Phoenix Gloves (4331, -0.04 DPS) [crafted]; Gnoll Casting Gloves (892, -0.20 DPS) [world]; Flame Gloves (253917, -0.33 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.00 DPS) [crafted]; Novice Ardent's Sash (253887, -0.40 DPS) [crafted]; Flame Sash (253929, -0.49 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.2 spell_power points (1.83 DPS) | yes | Filigreed Flame Leggings (253941, -0.15 DPS) [crafted]; Phoenix Pants (4317, -0.31 DPS) [crafted]; Silk-threaded Trousers (1929, -0.43 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.1 spell_power points (1.41 DPS) | yes | Feather Padded Treads (285345, -0.41 DPS, sim-verified) [world]; Flame Boots (253893, -0.50 DPS) [crafted]; Red Woolen Boots (4313, -0.61 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.00 DPS) | yes | Sludge-Stained Band (286535, -0.41 DPS) [world]; Lavishly Jeweled Ring (1156, -0.98 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.99 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.00 DPS) | yes | Sludge-Stained Band (286535, -0.39 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.97 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.99 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.2 spell_power points (0.04 DPS) | yes | Channeler's Staff (4437, -0.01 DPS) [world]; Lesser Staff of the Spire (1300, -0.02 DPS) [world]; Staff of Westfall (2042, -0.02 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.9 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.40 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.62 DPS) [dungeon]; Sizzle Stick (8071, -4.26 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 151, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 70.3. Weights run: 1.5s. Verify run: 1.0s. 266 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.044 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.423 per %), hit=0.125 ± 0.004 per rating point (10 rating = 1%, 1.247 per %), spell_haste=-1.894 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.239 ± 0.000, fire_power=0.761 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.42 DPS) | yes | Silk Headband (7050, -0.44 DPS) [crafted]; Flame Circlet (253953, -0.58 DPS) [crafted]; Filigreed Flame Circlet (253979, -2.87 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 spell_power points (1.60 DPS) | yes | Crystal Starfire Medallion (5003, -1.56 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.71 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.4 spell_power points (2.07 DPS) | yes | Moonlit Amice (11884, -0.53 DPS) [quest]; Death Speaker Mantle (6685, -0.64 DPS) [dungeon]; Invoker's Mantle (215365, -0.79 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | 6.0 spell_power points (1.32 DPS) | yes | Hillman's Cloak (3719, -0.22 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.44 DPS) [crafted]; Prelacy Cape (7004, -0.44 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.86 DPS) | yes | Flame Gown (253965, -0.38 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.75 DPS) [crafted]; Beguiler Robes (7728, -0.80 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.9 spell_power points (2.18 DPS) | yes | Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Windsong Bangles (263336, -1.96 DPS) [quest]; Glowing Magical Bracelets (13106, -2.10 DPS) [world_drop] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.54 DPS) | yes | Phoenix Gloves (4331, -0.03 DPS) [crafted]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Serpent Gloves (5970, -0.29 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.1 spell_power points (2.45 DPS) | yes | Belt of Arugal (6392, -0.44 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.69 DPS) [dungeon]; Invoker's Cord (215366, -0.86 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.64 DPS) | yes | Flame Leggings (253991, -0.23 DPS) [crafted]; Smoldering Pants (3073, -0.46 DPS) [world]; Abomination Skin Leggings (23173, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (70.3 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Nimbus Boots (6998, -0.29 DPS) [quest]; Fiery Slippers (254005, -0.84 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.54 DPS) | yes | Minor Channeling Ring (1449, -0.42 DPS) [quest]; Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon]; Sludge-Stained Band (286535, -0.88 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.32 DPS) | yes | Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon]; Sludge-Stained Band (286535, -0.66 DPS) [world]; Minor Channeling Ring (1449, -1.56 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | 30.0 spell_power points (6.60 DPS) | yes | Wind Spirit Staff (6689, -1.49 DPS) [dungeon]; Staff of Soran'ruk (15109, -3.86 DPS, sim-verified) [quest]; Scorn's Focal Dagger (23168, -4.62 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 153.9 spell_power points (33.87 DPS) | yes | Starfaller (13063, -1.04 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.77 DPS) [crafted]; Scorching Wand (5213, -4.02 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Phoenix Bindings; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff; ranged: Necrotic Wand

No-known-source sample (15 of 266, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 114.8. Weights run: 1.3s. Verify run: 1.0s. 346 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.060 ± 0.006, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.591 per %), hit=0.183 ± 0.006 per rating point (10 rating = 1%, 1.825 per %), spell_haste=-6.799 ± 0.224, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.247 ± 0.000, fire_power=0.753 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.95 DPS) | yes | Living Cowl (5608, -1.90 DPS, sim-verified) [world]; Augural Shroud (2620, -2.22 DPS) [world]; Holy Shroud (2721, -2.36 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.5 spell_power points (1.76 DPS) | yes | Scorn's Icy Choker (23169, -0.98 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.62 DPS) [quest]; Darkspear Warding Pendant (272074, -1.67 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.25 DPS) | yes | Green Silken Shoulders (7057, -0.21 DPS) [crafted]; Inquisitor's Shawl (19507, -0.42 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (114.8 DPS) | yes | Mantle of Lady Falther'ess (23178, +0.00 DPS) [dungeon]; Long Silken Cloak (4326, -0.17 DPS) [crafted]; Guardian Cloak (5965, -0.17 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.4 spell_power points (5.27 DPS) | yes | Elemental Raiment (9434, -0.41 DPS, sim-verified) [world_drop]; Cindercloth Robe (10042, -0.48 DPS) [crafted]; Dreamweave Vest (10021, -0.90 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.8 spell_power points (2.31 DPS) | yes | Arcane Runed Bracers (4744, -0.18 DPS) [quest]; Spidertank Oilrag (9448, -0.18 DPS) [dungeon]; Condor Bracers (15864, -0.66 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.30 DPS) | yes | Black Mageweave Gloves (10003, -1.09 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.57 DPS) [crafted]; Fiery Handwraps (254025, -1.72 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.2 spell_power points (3.36 DPS) | yes | Star Belt (4329, -0.29 DPS) [crafted]; Fiery Cord (254041, -0.76 DPS) [crafted]; Scorching Sash (4117, -1.05 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.7 spell_power points (3.47 DPS) | yes | Gaze Dreamer Pants (6903, -0.64 DPS) [dungeon]; Flame Leggings (253991, -0.89 DPS) [crafted]; Smoldering Pants (3073, -1.16 DPS) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.66 DPS) | yes | Fiery Slippers (254005, -2.98 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.91 DPS) [crafted]; Spidersilk Boots (4320, -3.95 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.43 DPS) | yes | Ring of Forlorn Spirits (2043, -0.54 DPS) [quest]; Reedknot Ring (9622, -0.78 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.01 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.12 DPS) | yes | Ring of Forlorn Spirits (2043, -0.24 DPS) [quest]; Reedknot Ring (9622, -0.47 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.71 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.03 DPS) [dungeon]; Staff of Noh'Orahil (15105, -3.37 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.30 DPS) | yes | Thrash's Trash (276204, +0.00 DPS, sim-verified) [vendor]; Orb of Noh'Orahil (15107, -0.75 DPS) [quest]; Rod of Molten Fire (2565, -0.99 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 167.3 spell_power points (39.46 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -3.87 DPS) [dungeon]; Twisted Nether Wand (249144, -4.05 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 346, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 194.1. Weights run: 3.4s. Verify run: 1.4s. 436 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.052 ± 0.017, crit=0.091 ± 0.004 per rating point (14 rating = 1%, 1.272 per %), hit=0.454 ± 0.022 per rating point (10 rating = 1%, 4.543 per %), spell_haste=-11.708 ± 0.522, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.256 ± 0.000, fire_power=0.744 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (5.96 DPS) | yes | Dreamweave Circlet (10041, -1.21 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.32 DPS) [crafted]; Red Mageweave Headband (10033, -1.53 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.4 spell_power points (1.64 DPS) | yes | Mindburst Medallion (11196, -0.24 DPS) [quest]; Horizon Choker (13085, -1.48 DPS) [world_drop]; Scorn's Icy Choker (23169, -3.25 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (194.1 DPS) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherflame Shoulders (254053, -0.19 DPS) [crafted]; Black Mageweave Shoulders (10027, -0.77 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (3.16 DPS) | yes | Cindercloth Cloak (14044, -0.93 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.07 DPS) [dungeon]; Cloak of Fire (14134, -5.35 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (4.93 DPS) | yes | Cindercloth Robe (10042, -0.49 DPS) [crafted]; Acumen Robes (17775, -0.50 DPS) [quest]; Elemental Raiment (9434, -4.77 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.0 spell_power points (2.22 DPS) | yes | Arcane Runed Bracers (4744, -0.23 DPS) [quest]; Spidertank Oilrag (9448, -0.23 DPS) [dungeon]; Phoenix Bindings (210781, -2.38 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.02 DPS) | yes | Fiery Gloves (254099, -0.78 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.05 DPS) [vendor]; Black Mageweave Gloves (10003, -5.60 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.1 spell_power points (3.78 DPS) | yes | Satyrmane Sash (17755, -0.57 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.64 DPS) [rep]; Fiery Waistcord (254085, -4.57 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 23.5 spell_power points (5.19 DPS) | yes | Red Mageweave Pants (10009, -1.96 DPS) [crafted]; Cindercloth Leggings (12256, -2.07 DPS) [vendor]; Wizardweave Leggings (14132, -4.51 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.30 DPS) | yes | Fiery Sandals (254111, -1.91 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.43 DPS) [vendor]; Cindercloth Boots (10044, -5.64 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.87 DPS) | yes | Philanthropist's Ring (281635, -0.60 DPS) [quest]; Cyclopean Band (11824, -0.80 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.10 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (2.65 DPS) | yes | Philanthropist's Ring (281635, -0.38 DPS) [quest]; Cyclopean Band (11824, -0.58 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -0.88 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.32 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.30 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.44 DPS) [dungeon]; Blade of Eternal Darkness (17780, -17.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.68 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.59 DPS) [crafted]; Pyric Caduceus (11748, -5.39 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532000000000000-2332113101220001350-0040000000000000)

Set DPS (verified): 410.8. Weights run: 1.4s. Verify run: 2.8s. 1064 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.114 ± 0.027, crit=0.118 ± 0.004 per rating point (14 rating = 1%, 1.652 per %), hit=0.603 ± 0.030 per rating point (10 rating = 1%, 6.027 per %), spell_haste=-10.450 ± 0.587, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.341 ± 0.001, fire_power=0.659 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.9 spell_power points (10.71 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -1.44 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -4.75 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.62 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.91 DPS) [quest]; Diana's Pearl Necklace (22403, -2.10 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.4 spell_power points (10.51 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.19 DPS) [pvp]; Argent Shoulders (19059, -1.86 DPS) [crafted]; Burial Shawl (18681, -2.96 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.9 spell_power points (7.94 DPS) | yes | Crystalline Threaded Cape (20697, -0.86 DPS) [world]; Amplifying Cloak (18350, -1.71 DPS) [dungeon]; Hide of the Wild (18510, -2.70 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (16.29 DPS) | yes | Robe of Everlasting Night (18385, -3.20 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -4.26 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -6.84 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.9 spell_power points (7.94 DPS) | yes | Sublime Wristguards (18497, -3.38 DPS) [dungeon]; Runecloth Cuffs (254123, -3.73 DPS) [crafted]; Netherflame Cuffs (254065, -4.69 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (9.55 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.66 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.5 spell_power points (11.61 DPS) | yes | Ban'thok Sash (11662, -4.94 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -4.99 DPS) [rep]; Belt of the Archmage (18405, -7.47 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 36.4 spell_power points (12.62 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -2.41 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.31 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.69 DPS) [crafted]; Omnicast Boots (11822, -0.91 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -4.51 DPS) [dungeon]; Maiden's Circle (13001, -5.63 DPS) [world_drop]; Naglering (11669, -15.57 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.42 DPS) [dungeon]; Maiden's Circle (13001, -1.54 DPS) [world_drop]; Naglering (11669, -6.90 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+18.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.42 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.88 DPS, sim-verified) [quest]; Serenity Field (272439, -5.20 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.22 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -29.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (410.8 DPS) | yes | Bonecreeper Stylus (13938, -1.06 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.24 DPS) [world]; Torch of Light (279246, -11.16 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1064, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 05500000000000000-2035113111120001350-0050051000000000)

Set DPS (verified): 764.6. Weights run: 4.5s. Verify run: 2.8s. 1064 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.077 ± 0.021, crit=0.179 ± 0.006 per rating point (14 rating = 1%, 2.511 per %), hit=0.724 ± 0.042 per rating point (10 rating = 1%, 7.236 per %), spell_haste=-21.265 ± 0.868, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.265 ± 0.001, fire_power=0.735 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -0.04 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.92 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.55 DPS) [dungeon]; Amulet of the Dawn (22657, -2.98 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.7 spell_power points (15.22 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.16 DPS) [pvp]; Argent Shoulders (19059, -4.07 DPS, sim-verified) [crafted]; Burial Shawl (18681, -4.68 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.9 spell_power points (11.84 DPS) | yes | Crystalline Threaded Cape (20697, -1.76 DPS) [world]; Amplifying Cloak (18350, -2.90 DPS) [dungeon]; Mageflame Cloak (13007, -4.18 DPS) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (23.17 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.37 DPS) [pvp]; Robe of Everlasting Night (18385, -9.27 DPS) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -10.00 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.6 spell_power points (11.22 DPS) | yes | Sublime Wristguards (18497, -4.89 DPS) [dungeon]; Runecloth Cuffs (254123, -5.38 DPS) [crafted]; Netherflame Cuffs (254065, -6.21 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (13.59 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.21 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.9 spell_power points (16.84 DPS) | yes | Belt of the Archmage (18405, -6.39 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.87 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -7.52 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 37.8 spell_power points (18.75 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -4.36 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.91 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.99 DPS) [crafted]; Omnicast Boots (11822, -1.52 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.30 DPS) [dungeon]; Maiden's Circle (13001, -8.63 DPS) [world_drop]; Naglering (11669, -21.62 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (764.6 DPS) | yes | Rune Band of Wizardry (22339, +0.00 DPS) [dungeon]; Songstone of Ironforge (12543, -2.14 DPS) [quest]; Maiden's Circle (13001, -2.14 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+23.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.47 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -4.18 DPS, sim-verified) [quest]; Serenity Field (272439, -7.44 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.53 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -44.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 160.7 spell_power points (79.74 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -11.50 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.58 DPS) [world] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; ranged: Torch of Light

No-known-source sample (15 of 1064, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 34.1. Weights run: 1.5s. Verify run: 0.9s. 140 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.019 ± 0.002, crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.324 per %), hit=0.092 ± 0.003 per rating point (10 rating = 1%, 0.922 per %), spell_haste=1.980 ± 0.050, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.246 ± 0.000, fire_power=0.754 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Flame Circlet (253953, -0.75 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.18 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.2 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.02 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.80 DPS) | yes | Feyscale Cloak (6632, -0.20 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.20 DPS) [rep]; Black Whelp Cloak (7283, -0.26 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gray Woolen Robe (2585, -0.20 DPS) [crafted]; Green Woolen Vest (2582, -0.22 DPS) [crafted]; Filigreed Flame Gown (253905, -0.51 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.60 DPS) | yes | Windsong Bangles (263336, -0.40 DPS) [quest]; Owlbeard Bracers (16981, -0.47 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.58 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.40 DPS) | yes | Phoenix Gloves (4331, -0.04 DPS) [crafted]; Gnoll Casting Gloves (892, -0.20 DPS) [world]; Flame Gloves (253917, -0.33 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.00 DPS) [crafted]; Novice Ardent's Sash (253887, -0.40 DPS) [crafted]; Flame Sash (253929, -0.48 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.2 spell_power points (1.83 DPS) | yes | Filigreed Flame Leggings (253941, -0.15 DPS) [crafted]; Phoenix Pants (4317, -0.31 DPS) [crafted]; Silk-threaded Trousers (1929, -0.43 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.1 spell_power points (1.41 DPS) | yes | Flame Boots (253893, -0.50 DPS) [crafted]; Feather Padded Treads (285345, -0.53 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.61 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.00 DPS) | yes | Lavishly Jeweled Ring (1156, -0.97 DPS) [dungeon]; Loop of Sacrifice (281673, -0.98 DPS) [quest]; Volcanic Rock Ring (12053, -0.99 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.60 DPS) | yes | Loop of Sacrifice (281673, -0.58 DPS) [quest]; Volcanic Rock Ring (12053, -0.59 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.64 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.2 spell_power points (0.04 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.01 DPS) [world]; Lesser Staff of the Spire (1300, -0.02 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 114.9 spell_power points (22.91 DPS) | yes | Skycaller (12984, -1.48 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.62 DPS) [dungeon]; Sizzle Stick (8071, -4.26 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 70.4. Weights run: 1.5s. Verify run: 1.0s. 249 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.044 ± 0.003, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.423 per %), hit=0.125 ± 0.004 per rating point (10 rating = 1%, 1.247 per %), spell_haste=-1.894 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.239 ± 0.000, fire_power=0.761 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.42 DPS) | yes | Silk Headband (7050, -0.44 DPS) [crafted]; Flame Circlet (253953, -0.58 DPS) [crafted]; Filigreed Flame Circlet (253979, -2.83 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 spell_power points (1.60 DPS) | yes | Darkspear Warding Pendant (272075, -1.46 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.56 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.4 spell_power points (2.07 DPS) | yes | Invoker's Mantle (215365, -0.48 DPS) [crafted]; Death Speaker Mantle (6685, -0.64 DPS) [dungeon]; Chestnut Mantle (17695, -0.74 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.10 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Battle Healer's Cloak (19529, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.86 DPS) | yes | Flame Gown (253965, -0.60 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.75 DPS) [crafted]; Beguiler Robes (7728, -0.80 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.9 spell_power points (2.18 DPS) | yes | Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Tabitha's Cuffs (251486, -1.52 DPS) [quest]; Owlbeard Bracers (16981, -1.94 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.54 DPS) | yes | Phoenix Gloves (4331, -0.03 DPS) [crafted]; Jutebraid Gloves (10654, -0.17 DPS) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.1 spell_power points (2.45 DPS) | yes | Warsong Sash (16975, -0.31 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.44 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.69 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.64 DPS) | yes | Flame Leggings (253991, -0.38 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.46 DPS) [world]; Abomination Skin Leggings (23173, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Acidic Walkers (9454, -0.43 DPS) [dungeon]; Fiery Slippers (254005, -0.81 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.54 DPS) | yes | Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon]; Sludge-Stained Band (286535, -0.88 DPS) [world]; Sacred Band (6669, -1.10 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.32 DPS) | yes | Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon]; Sacred Band (6669, -0.88 DPS) [quest]; Sludge-Stained Band (286535, -1.84 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Focal Dagger (23168, -3.13 DPS) [dungeon]; Staff of Soran'ruk (15109, -4.97 DPS, sim-verified) [quest]; Glimmering Staff (249392, -5.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Starfaller (13063, -0.97 DPS) [world_drop]; Unstable Power Core (279847, -2.46 DPS, sim-verified) [quest]; Greater Mystic Wand (217287, -3.77 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Phoenix Bindings; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff; ranged: Necrotic Wand

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 114.3. Weights run: 1.3s. Verify run: 1.0s. 324 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.060 ± 0.006, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.591 per %), hit=0.183 ± 0.006 per rating point (10 rating = 1%, 1.825 per %), spell_haste=-6.799 ± 0.224, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.247 ± 0.000, fire_power=0.753 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.95 DPS) | yes | Living Cowl (5608, -1.86 DPS, sim-verified) [world]; Augural Shroud (2620, -2.22 DPS) [world]; Holy Shroud (2721, -2.36 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.5 spell_power points (1.76 DPS) | yes | Scorn's Icy Choker (23169, -0.83 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.62 DPS) [quest]; Darkspear Warding Pendant (272074, -1.67 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.25 DPS) | yes | Green Silken Shoulders (7057, -0.21 DPS) [crafted]; Chestnut Mantle (17695, -0.36 DPS) [quest]; Inquisitor's Shawl (19507, -0.42 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 9.5 spell_power points (2.25 DPS) | yes | Long Silken Cloak (4326, -0.76 DPS) [crafted]; Guardian Cloak (5965, -0.76 DPS) [crafted]; Icy Cloak (4327, -1.18 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.4 spell_power points (5.27 DPS) | yes | Elemental Raiment (9434, -0.33 DPS, sim-verified) [world_drop]; Cindercloth Robe (10042, -0.48 DPS) [crafted]; Dreamweave Vest (10021, -0.90 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 9.8 spell_power points (2.31 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.18 DPS) [dungeon]; Condor Bracers (15864, -0.66 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.30 DPS) | yes | Black Mageweave Gloves (10003, -0.81 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.57 DPS) [crafted]; Fiery Handwraps (254025, -1.72 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.2 spell_power points (3.36 DPS) | yes | Star Belt (4329, -0.29 DPS) [crafted]; Fiery Cord (254041, -0.76 DPS) [crafted]; Warsong Sash (16975, -0.76 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.7 spell_power points (3.47 DPS) | yes | Gaze Dreamer Pants (6903, -0.87 DPS, sim-verified) [dungeon]; Flame Leggings (253991, -0.89 DPS) [crafted]; Smoldering Pants (3073, -1.16 DPS) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.66 DPS) | yes | Fiery Slippers (254005, -3.18 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -3.91 DPS) [crafted]; Spidersilk Boots (4320, -3.95 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.3 spell_power points (2.43 DPS) | yes | Reedknot Ring (9622, -0.78 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.01 DPS) [vendor]; Sludge-Stained Band (286535, -1.72 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.12 DPS) | yes | Reedknot Ring (9622, -0.38 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.71 DPS) [vendor]; Sludge-Stained Band (286535, -1.42 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (114.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.03 DPS) [dungeon]; Staff of Noh'Orahil (15105, -3.37 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.30 DPS) | yes | Thrash's Trash (276204, +0.00 DPS, sim-verified) [vendor]; Orb of Noh'Orahil (15107, -0.75 DPS) [quest]; Rod of Molten Fire (2565, -0.99 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 167.3 spell_power points (39.46 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -3.87 DPS) [dungeon]; Twisted Nether Wand (249144, -4.05 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 324, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 190.5. Weights run: 3.4s. Verify run: 1.3s. 410 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.052 ± 0.017, crit=0.091 ± 0.004 per rating point (14 rating = 1%, 1.272 per %), hit=0.454 ± 0.022 per rating point (10 rating = 1%, 4.543 per %), spell_haste=-11.708 ± 0.522, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.256 ± 0.000, fire_power=0.744 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (5.96 DPS) | yes | Dreamweave Circlet (10041, -1.21 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.32 DPS) [crafted]; Red Mageweave Headband (10033, -1.53 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 7.4 spell_power points (1.64 DPS) | yes | Mindburst Medallion (11196, -0.24 DPS) [quest]; Scorn's Icy Choker (23169, -1.31 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.48 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherflame Shoulders (254053, -0.19 DPS) [crafted]; Black Mageweave Shoulders (10027, -0.77 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 spell_power points (3.16 DPS) | yes | Deep Woodlands Cloak (19121, -0.41 DPS) [quest]; Cloak of Fire (14134, -0.86 DPS) [crafted]; Cindercloth Cloak (14044, -0.93 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 spell_power points (4.93 DPS) | yes | Elemental Raiment (9434, -0.29 DPS) [world_drop]; Cindercloth Robe (10042, -0.49 DPS) [crafted]; Acumen Robes (17775, -0.50 DPS) [quest] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.0 spell_power points (2.22 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Phoenix Bindings (210781, -0.08 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 spell_power points (4.02 DPS) | yes | Fiery Gloves (254099, -0.78 DPS) [crafted]; Black Mageweave Gloves (10003, -0.87 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.05 DPS) [vendor] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | sim-verified (190.5 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, -0.02 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.09 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 23.5 spell_power points (5.19 DPS) | yes | Red Mageweave Pants (10009, -1.96 DPS) [crafted]; Cindercloth Leggings (12256, -2.07 DPS) [vendor]; Wizardweave Leggings (14132, -2.39 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.30 DPS) | yes | Fiery Sandals (254111, -1.91 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.43 DPS) [vendor]; Cindercloth Boots (10044, -2.88 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.87 DPS) | yes | Philanthropist's Ring (281635, -0.60 DPS) [quest]; Cyclopean Band (11824, -0.80 DPS) [dungeon]; Runed Ring (862, -1.32 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (2.65 DPS) | yes | Philanthropist's Ring (281635, -0.38 DPS) [quest]; Cyclopean Band (11824, -0.58 DPS) [dungeon]; Runed Ring (862, -1.10 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.32 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.95 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.20 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.30 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.44 DPS) [dungeon]; Blade of Eternal Darkness (17780, -16.33 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.68 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.59 DPS) [crafted]; Pyric Caduceus (11748, -7.65 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Fiery Waistcord; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Spire of Hakkar; ranged: Noxious Shooter

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532000000000000-2332113101220001350-0040000000000000)

Set DPS (verified): 404.1. Weights run: 1.4s. Verify run: 2.7s. 1052 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.114 ± 0.027, crit=0.118 ± 0.004 per rating point (14 rating = 1%, 1.652 per %), hit=0.603 ± 0.030 per rating point (10 rating = 1%, 6.027 per %), spell_haste=-10.450 ± 0.587, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.341 ± 0.001, fire_power=0.659 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.9 spell_power points (10.71 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -1.44 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -4.28 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.62 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.91 DPS) [quest]; Diana's Pearl Necklace (22403, -2.10 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.4 spell_power points (10.51 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.19 DPS) [pvp]; Argent Shoulders (19059, -1.86 DPS) [crafted]; Burial Shawl (18681, -2.96 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.9 spell_power points (7.94 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -1.71 DPS) [dungeon]; Hide of the Wild (18510, -2.70 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (16.29 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.26 DPS) [pvp]; Robe of Everlasting Night (18385, -6.42 DPS) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -6.84 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.9 spell_power points (7.94 DPS) | yes | Sublime Wristguards (18497, -3.38 DPS) [dungeon]; Runecloth Cuffs (254123, -3.73 DPS) [crafted]; Netherflame Cuffs (254065, -4.69 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (9.55 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.66 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.5 spell_power points (11.61 DPS) | yes | Belt of the Archmage (18405, -2.71 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -4.94 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -4.99 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 36.4 spell_power points (12.62 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.17 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.31 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.69 DPS) [crafted]; Omnicast Boots (11822, -0.91 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -4.51 DPS) [dungeon]; Maiden's Circle (13001, -5.63 DPS) [world_drop]; Naglering (11669, -13.84 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.42 DPS) [dungeon]; Maiden's Circle (13001, -1.54 DPS) [world_drop]; Naglering (11669, -8.15 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+17.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.42 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.83 DPS, sim-verified) [quest]; Serenity Field (272439, -5.20 DPS) [vendor] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.22 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -29.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (404.1 DPS) | yes | Bonecreeper Stylus (13938, -1.06 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.24 DPS) [world]; Torch of Light (279246, -11.34 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1052, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 05500000000000000-2035113111120001350-0050051000000000)

Set DPS (verified): 765.4. Weights run: 4.5s. Verify run: 3.0s. 1052 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.077 ± 0.021, crit=0.179 ± 0.006 per rating point (14 rating = 1%, 2.511 per %), hit=0.724 ± 0.042 per rating point (10 rating = 1%, 7.236 per %), spell_haste=-21.265 ± 0.868, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.265 ± 0.001, fire_power=0.735 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -0.04 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.92 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.55 DPS) [dungeon]; Amulet of the Dawn (22657, -2.98 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 30.7 spell_power points (15.22 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.16 DPS) [pvp]; Argent Shoulders (19059, -2.81 DPS) [crafted]; Burial Shawl (18681, -4.68 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.9 spell_power points (11.84 DPS) | yes | Crystalline Threaded Cape (20697, -1.76 DPS) [world]; Amplifying Cloak (18350, -2.90 DPS) [dungeon]; Mageflame Cloak (13007, -4.18 DPS) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.7 spell_power points (23.17 DPS) | yes | Robe of Everlasting Night (18385, -4.81 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -6.37 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.00 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.6 spell_power points (11.22 DPS) | yes | Sublime Wristguards (18497, -4.89 DPS) [dungeon]; Runecloth Cuffs (254123, -5.38 DPS) [crafted]; Netherflame Cuffs (254065, -6.21 DPS) [crafted] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.4 spell_power points (13.59 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Inferno Gloves (18408, -1.21 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 33.9 spell_power points (16.84 DPS) | yes | Belt of the Archmage (18405, -6.37 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.87 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -7.52 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 37.8 spell_power points (18.75 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -4.13 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.91 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.99 DPS) [crafted]; Omnicast Boots (11822, -1.52 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.30 DPS) [dungeon]; Maiden's Circle (13001, -8.63 DPS) [world_drop]; Naglering (11669, -20.70 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (765.4 DPS) | yes | Rune Band of Wizardry (22339, +0.00 DPS) [dungeon]; Eye of Orgrimmar (12545, -2.14 DPS) [quest]; Maiden's Circle (13001, -2.14 DPS) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+21.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.98 DPS) [quest]; Weakness Analyzer (272438, -3.47 DPS) [vendor]; Burst of Knowledge (11832, -6.49 DPS, sim-verified) [dungeon] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.53 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -43.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+10.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.07 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.15 DPS) [world]; Torch of Light (279246, -10.12 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1052, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

