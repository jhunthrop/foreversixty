# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 43.9. Weights run: 2.6s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.055, intellect=0.067 ± 0.004, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.385 per %), hit=0.104 ± 0.001 per rating point (10 rating = 1%, 1.037 per %), spell_haste=0.247 ± 0.053, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.787 ± 0.055, fire_power=0.212 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.30 DPS) | yes | Shadow Goggles (4373, -3.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.6 spell_power points (1.22 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.29 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.35 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.87 DPS) | yes | Feyscale Cloak (6632, -0.22 DPS) [dungeon]; Caretaker's Cape (20428, -0.22 DPS) [rep]; Black Whelp Cloak (7283, -0.31 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.3 spell_power points (1.16 DPS) | yes | Green Woolen Vest (2582, -0.29 DPS) [crafted]; Bloody Apron (6226, -0.29 DPS) [dungeon]; Gray Woolen Robe (2585, -1.47 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.22 DPS) | yes | Bright Bracers (3647, -0.16 DPS) [world_drop]; Repurposed Hair Band (281256, -0.19 DPS) [quest]; Mindthrust Bracers (1974, -0.29 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.52 DPS) | yes | Gnoll Casting Gloves (892, -0.36 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.61 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.06 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.3 spell_power points (0.93 DPS) | yes | Novice Ardent's Sash (253887, -0.45 DPS) [crafted]; Keller's Girdle (2911, -0.81 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.94 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 spell_power points (2.07 DPS) | yes | Filigreed Pristine Leggings (253937, -0.68 DPS) [crafted]; Silk-threaded Trousers (1929, -0.83 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.99 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.3 spell_power points (1.58 DPS) | yes | Red Woolen Boots (4313, -0.71 DPS) [crafted]; Pristine Boots (253889, -0.88 DPS) [crafted]; Feather Padded Treads (285345, -0.99 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.1 spell_power points (1.12 DPS) | yes | Sludge-Stained Band (286535, -0.46 DPS) [world]; Lavishly Jeweled Ring (1156, -1.03 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.07 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.09 DPS) | yes | Sludge-Stained Band (286535, -0.65 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -1.00 DPS) [dungeon]; Volcanic Rock Ring (12053, -1.04 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.7 spell_power points (0.15 DPS) | yes | Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.06 DPS) [world]; Staff of Westfall (2042, -0.07 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 105.7 spell_power points (22.96 DPS) | yes | Skycaller (12984, -2.49 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.67 DPS) [dungeon]; Sizzle Stick (8071, -4.23 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 73.4. Weights run: 2.6s. Verify run: 1.3s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.127, intellect=0.254 ± 0.008, crit=0.068 ± 0.003 per rating point (14 rating = 1%, 0.956 per %), hit=0.163 ± 0.001 per rating point (10 rating = 1%, 1.626 per %), spell_haste=not significant (-0.047 ± 0.125), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.730 ± 0.127, fire_power=0.272 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.99 DPS) | yes | Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Embalmed Shroud (7691, -0.54 DPS) [dungeon]; Silk Headband (7050, -0.77 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (1.55 DPS) | yes | Crystal Starfire Medallion (5003, -1.36 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.36 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.05 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.3 spell_power points (2.05 DPS) | yes | Fairywing Mantle (9536, -0.54 DPS) [quest]; Invoker's Mantle (215365, -0.55 DPS) [crafted]; Death Speaker Mantle (6685, -0.84 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.91 DPS) | yes | Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Prelacy Cape (7004, -0.18 DPS) [quest]; Repairman's Cape (9605, -0.49 DPS, sim-verified) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.36 DPS) | yes | Death Speaker Robes (6682, -0.58 DPS) [dungeon]; Pristine Gown (253961, -0.77 DPS) [crafted]; Green Silk Armor (7065, -0.83 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.63 DPS) | yes | Nightsky Wristbands (6407, -1.36 DPS) [world_drop]; Stonecloth Bindings (14416, -1.40 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.69 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.27 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.04 DPS) [quest]; Gnoll Casting Gloves (892, -0.18 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.8 spell_power points (2.13 DPS) | yes | Invoker's Cord (215366, -0.63 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.68 DPS) [dungeon]; Belt of Arugal (6392, -0.73 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.18 DPS) | yes | Abomination Skin Leggings (23173, -0.47 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.58 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.81 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.8 spell_power points (1.59 DPS) | yes | Acidic Walkers (9454, -0.32 DPS) [dungeon]; Nimbus Boots (6998, -0.50 DPS) [quest]; Spidersilk Boots (4320, -1.81 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.27 DPS) | yes | Minor Channeling Ring (1449, -0.27 DPS) [quest]; Electrocutioner Lagnut (9447, -0.73 DPS) [dungeon]; Sludge-Stained Band (286535, -0.73 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.09 DPS) | yes | Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Sludge-Stained Band (286535, -0.54 DPS) [world]; Minor Channeling Ring (1449, -2.08 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.63 DPS) | yes | Twisted Chanter's Staff (890, -1.17 DPS) [world_drop]; Channeler's Staff (4437, -1.26 DPS) [world]; Glimmering Staff (249392, -1.97 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.5 spell_power points (1.55 DPS) | yes | Dwarven Tome (279898, -0.75 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.82 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.82 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 186.2 spell_power points (33.75 DPS) | yes | Starfaller (13063, -1.07 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.85 DPS) [crafted]; Gravestone Scepter (7001, -4.75 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 146.8. Weights run: 2.2s. Verify run: 1.2s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.210, intellect=0.323 ± 0.013, crit=0.073 ± 0.003 per rating point (14 rating = 1%, 1.022 per %), hit=0.210 ± 0.002 per rating point (10 rating = 1%, 2.100 per %), spell_haste=not significant (0.296 ± 0.181), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.557 ± 0.210), fire_power=0.444 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.62 DPS) | yes | Living Cowl (5608, -2.52 DPS) [world]; Augural Shroud (2620, -2.87 DPS, sim-verified) [world]; Holy Shroud (2721, -3.15 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (2.82 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.89 DPS, sim-verified) [quest]; Triune Amulet (7722, -2.10 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.10 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.9 spell_power points (3.75 DPS) | yes | Green Silken Shoulders (7057, -0.11 DPS) [crafted]; Inquisitor's Shawl (19507, -0.22 DPS) [dungeon]; Berylline Pads (4197, -0.53 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 11.9 spell_power points (3.75 DPS) | yes | Guardian Cloak (5965, -1.35 DPS) [crafted]; Icy Cloak (4327, -1.55 DPS) [crafted]; Long Silken Cloak (4326, -1.76 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.9 spell_power points (7.55 DPS) | yes | Elemental Raiment (9434, -0.52 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.96 DPS) [crafted]; Robe of Power (7054, -1.91 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.84 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.63 DPS) [quest]; Earthen Silk Cuffs (254019, -1.58 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (6.08 DPS) | yes | Black Mageweave Gloves (10003, -0.82 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.59 DPS) [crafted]; Gilded Handwraps (254021, -2.85 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.3 spell_power points (4.82 DPS) | yes | Star Belt (4329, -0.72 DPS) [crafted]; Deathmage Sash (10771, -1.09 DPS) [dungeon]; Gilded Cord (254037, -1.48 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.9 spell_power points (5.64 DPS) | yes | Gaze Dreamer Pants (6903, -1.85 DPS) [dungeon]; Abomination Skin Leggings (23173, -1.98 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.00 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.57 DPS) | yes | Gilded Slippers (254001, -2.80 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.95 DPS) [crafted]; Acidic Walkers (9454, -5.17 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.9 spell_power points (3.76 DPS) | yes | Ring of Forlorn Spirits (2043, -1.24 DPS) [quest]; Reedknot Ring (9622, -1.56 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.87 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.84 DPS) | yes | Ring of Forlorn Spirits (2043, -0.32 DPS) [quest]; Reedknot Ring (9622, -0.63 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.95 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.47 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.52 DPS) [quest]; Gut Ripper (2164, -7.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (146.8 DPS) | yes | Umbral Wand (5216, -0.22 DPS) [dungeon]; Earthen Rod (9381, -0.30 DPS) [dungeon]; Jaina's Firestarter (13064, -2.76 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 202.8. Weights run: 2.1s. Verify run: 1.4s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.205, intellect=0.315 ± 0.014, crit=0.076 ± 0.003 per rating point (14 rating = 1%, 1.060 per %), hit=0.198 ± 0.002 per rating point (10 rating = 1%, 1.979 per %), spell_haste=not significant (0.529 ± 0.198), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.711 ± 0.205), fire_power=0.290 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.89 DPS) | yes | Dreamweave Circlet (10041, -1.15 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.42 DPS) [crafted]; Red Mageweave Headband (10033, -2.72 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (3.59 DPS) | yes | Mindburst Medallion (11196, -0.40 DPS) [quest]; Horizon Choker (13085, -1.81 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.32 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.7 spell_power points (7.53 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -2.35 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.73 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (6.41 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.27 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.76 DPS) [crafted]; Nightfall Drape (12465, -2.78 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.3 spell_power points (10.20 DPS) | yes | Robe of the Magi (1716, -0.98 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.73 DPS) [world_drop]; Dreamweave Vest (10021, -1.80 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 9.2 spell_power points (3.71 DPS) | yes | Spidertank Oilrag (9448, -0.08 DPS) [dungeon]; Bloodband Bracers (11469, -0.55 DPS) [quest]; Arcane Runed Bracers (4744, -2.85 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (7.77 DPS) | yes | Black Mageweave Gloves (10003, -1.72 DPS) [crafted]; Runecloth Gloves (13863, -1.79 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.01 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 17.1 spell_power points (6.92 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.76 DPS) [rep]; Ghostweave Cord (254073, -1.27 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.1 spell_power points (10.55 DPS) | yes | Red Mageweave Pants (10009, -3.38 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.76 DPS) [vendor]; Wizardweave Leggings (14132, -4.60 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.68 DPS) | yes | Gilded Sandals (254107, -2.27 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -4.36 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.51 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.24 DPS) | yes | Philanthropist's Ring (281635, -0.45 DPS) [quest]; Cyclopean Band (11824, -0.72 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -2.02 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.84 DPS) | yes | Cyclopean Band (11824, -0.32 DPS) [dungeon]; Philanthropist's Ring (281635, -1.13 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.61 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.42 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -4.21 DPS) [dungeon]; Glowing Brightwood Staff (812, -4.39 DPS) [world_drop]; Blade of Eternal Darkness (17780, -5.85 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (202.8 DPS) | yes | Lesser Eternal Wand (249232, -3.04 DPS) [crafted]; Wand of Allistarj (13065, -3.60 DPS) [world_drop]; Pyric Caduceus (11748, -6.34 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 377.9. Weights run: 2.3s. Verify run: 1.4s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.250, intellect=0.338 ± 0.017, crit=0.134 ± 0.004 per rating point (14 rating = 1%, 1.874 per %), hit=0.322 ± 0.002 per rating point (10 rating = 1%, 3.222 per %), spell_haste=-2.608 ± 0.302, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.675 ± 0.250), fire_power=0.324 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.7 spell_power points (16.43 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.88 DPS) [pvp]; Deathmist Mask (226909, -4.59 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (11.05 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.31 DPS) [quest]; Beads of Ogre Mojo (22149, -2.48 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.0 spell_power points (17.06 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.61 DPS) [pvp]; Burial Shawl (18681, -4.29 DPS) [dungeon]; Argent Shoulders (19059, -4.50 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 21.9 spell_power points (11.02 DPS) | yes | Crystalline Threaded Cape (20697, -0.29 DPS) [world]; Amplifying Cloak (18350, -1.97 DPS) [dungeon]; Hide of the Wild (18510, -2.28 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.0 spell_power points (24.64 DPS) | yes | Robe of Everlasting Night (18385, -2.07 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -4.48 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -8.68 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.7 spell_power points (12.41 DPS) | yes | Sublime Wristguards (18497, -4.68 DPS) [dungeon]; Runecloth Cuffs (254123, -5.19 DPS) [crafted]; Deathmist Bracers (226907, -6.86 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.7 spell_power points (14.42 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -3.18 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 35.7 spell_power points (17.92 DPS) | yes | Belt of the Archmage (18405, -3.77 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -7.18 DPS) [rep]; Highlander's Cloth Girdle (20047, -8.92 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 39.1 spell_power points (19.66 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.22 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -3.38 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.1 spell_power points (12.09 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.03 DPS) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -2.69 DPS) [quest]; Maiden's Circle (13001, -2.69 DPS) [world_drop]; Naglering (11669, -9.02 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.51 DPS) [quest]; Maiden's Circle (13001, -1.51 DPS) [world_drop]; Naglering (11669, -8.69 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Ascendance (22678, -1.87 DPS, sim-verified) [quest]; Weakness Analyzer (272438, -3.52 DPS) [vendor]; Serenity Field (272439, -7.54 DPS) [vendor] |
| trinket2 | Royal Seal of Eldre'Thalas (18467) | Harnessing Shadows [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Ascendance (22678, +0.00 DPS, sim-verified) [quest]; Weakness Analyzer (272438, -0.50 DPS) [vendor]; Serenity Field (272439, -4.52 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.68 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -16.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (377.9 DPS) | yes | Bonecreeper Stylus (13938, -0.54 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.00 DPS) [world]; Torch of Light (279246, -22.51 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Royal Seal of Eldre'Thalas; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 42.5. Weights run: 2.6s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.055, intellect=0.067 ± 0.004, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.385 per %), hit=0.104 ± 0.001 per rating point (10 rating = 1%, 1.037 per %), spell_haste=0.247 ± 0.053, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.787 ± 0.055, fire_power=0.212 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.30 DPS) | yes | Shadow Goggles (4373, -3.26 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.6 spell_power points (1.22 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.30 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.35 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.87 DPS) | yes | Feyscale Cloak (6632, -0.22 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.22 DPS) [rep]; Black Whelp Cloak (7283, -0.36 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.3 spell_power points (1.16 DPS) | yes | Green Woolen Vest (2582, -0.29 DPS) [crafted]; Bloody Apron (6226, -0.29 DPS) [dungeon]; Gray Woolen Robe (2585, -1.65 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.1 spell_power points (0.25 DPS) | yes | Tabitha's Cuffs (251486, -0.16 DPS) [quest]; Featherbead Bracers (15452, -0.17 DPS) [quest]; Windsong Bangles (263336, -0.39 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.52 DPS) | yes | Gnoll Casting Gloves (892, -0.40 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.61 DPS) [crafted]; Apothecary Gloves (10919, -0.65 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.3 spell_power points (0.93 DPS) | yes | Novice Ardent's Sash (253887, -0.45 DPS) [crafted]; Keller's Girdle (2911, -0.81 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.14 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 spell_power points (2.07 DPS) | yes | Silk-threaded Trousers (1929, -0.65 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.68 DPS) [crafted]; Rumpled Kilt (274741, -0.99 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.3 spell_power points (1.58 DPS) | yes | Red Woolen Boots (4313, -0.71 DPS) [crafted]; Feather Padded Treads (285345, -0.84 DPS, sim-verified) [world]; Pristine Boots (253889, -0.88 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.09 DPS) | yes | Lavishly Jeweled Ring (1156, -1.00 DPS) [dungeon]; Loop of Sacrifice (281673, -1.01 DPS) [quest]; Volcanic Rock Ring (12053, -1.04 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.65 DPS) | yes | Loop of Sacrifice (281673, -0.58 DPS) [quest]; Volcanic Rock Ring (12053, -0.61 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.73 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 0.7 spell_power points (0.15 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.06 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 105.7 spell_power points (22.96 DPS) | yes | Firebelcher (5243, -2.67 DPS) [dungeon]; Skycaller (12984, -2.68 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.23 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 72.0. Weights run: 2.6s. Verify run: 1.3s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.127, intellect=0.254 ± 0.008, crit=0.068 ± 0.003 per rating point (14 rating = 1%, 0.956 per %), hit=0.163 ± 0.001 per rating point (10 rating = 1%, 1.626 per %), spell_haste=not significant (-0.047 ± 0.125), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.730 ± 0.127, fire_power=0.272 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.99 DPS) | yes | Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Embalmed Shroud (7691, -0.54 DPS) [dungeon]; Silk Headband (7050, -0.78 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (1.55 DPS) | yes | Crystal Starfire Medallion (5003, -1.36 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.36 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.98 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.3 spell_power points (2.05 DPS) | yes | Fairywing Mantle (9536, -0.54 DPS) [quest]; Invoker's Mantle (215365, -0.55 DPS) [crafted]; Death Speaker Mantle (6685, -0.77 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.91 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Battle Healer's Cloak (19529, -0.18 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.36 DPS) | yes | Death Speaker Robes (6682, -0.58 DPS) [dungeon]; Pristine Gown (253961, -0.77 DPS) [crafted]; Green Silk Armor (7065, -0.96 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.63 DPS) | yes | Nightsky Wristbands (6407, -1.36 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.36 DPS) [quest]; Glowing Magical Bracelets (13106, -2.80 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.3 spell_power points (1.32 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Truefaith Gloves (7049, -0.27 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.8 spell_power points (2.13 DPS) | yes | Belt of Arugal (6392, -0.36 DPS) [dungeon]; Warsong Sash (16975, -0.38 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.63 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.18 DPS) | yes | Pristine Leggings (253987, -0.58 DPS) [crafted]; Abomination Skin Leggings (23173, -0.60 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.81 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.8 spell_power points (1.59 DPS) | yes | Acidic Walkers (9454, -0.32 DPS) [dungeon]; Boots of the Enchanter (4325, -0.69 DPS) [crafted]; Spidersilk Boots (4320, -1.65 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.27 DPS) | yes | Electrocutioner Lagnut (9447, -0.73 DPS) [dungeon]; Sludge-Stained Band (286535, -0.73 DPS) [world]; Sacred Band (6669, -0.91 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.09 DPS) | yes | Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Sacred Band (6669, -0.73 DPS) [quest]; Sludge-Stained Band (286535, -2.60 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.63 DPS) | yes | Twisted Chanter's Staff (890, -1.17 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.17 DPS) [quest]; Glimmering Staff (249392, -1.99 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.5 spell_power points (1.55 DPS) | yes | Orb of Souls (249395, -0.82 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.00 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.04 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 186.2 spell_power points (33.75 DPS) | yes | Starfaller (13063, -0.92 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.85 DPS) [crafted]; Gravestone Scepter (7001, -4.75 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 144.5. Weights run: 2.2s. Verify run: 1.2s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.210, intellect=0.323 ± 0.013, crit=0.073 ± 0.003 per rating point (14 rating = 1%, 1.022 per %), hit=0.210 ± 0.002 per rating point (10 rating = 1%, 2.100 per %), spell_haste=not significant (0.296 ± 0.181), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.557 ± 0.210), fire_power=0.444 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.62 DPS) | yes | Augural Shroud (2620, -2.30 DPS, sim-verified) [world]; Living Cowl (5608, -2.52 DPS) [world]; Holy Shroud (2721, -3.15 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (2.82 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.94 DPS, sim-verified) [quest]; Triune Amulet (7722, -2.10 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.10 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.9 spell_power points (3.75 DPS) | yes | Green Silken Shoulders (7057, -0.11 DPS) [crafted]; Inquisitor's Shawl (19507, -0.22 DPS) [dungeon]; Berylline Pads (4197, -0.53 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 11.9 spell_power points (3.75 DPS) | yes | Guardian Cloak (5965, -1.35 DPS) [crafted]; Icy Cloak (4327, -1.55 DPS) [crafted]; Long Silken Cloak (4326, -1.67 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.9 spell_power points (7.55 DPS) | yes | Elemental Raiment (9434, -0.93 DPS) [world_drop]; Dreamweave Vest (10021, -0.96 DPS) [crafted]; Robe of Power (7054, -1.91 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.84 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.71 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.76 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (6.08 DPS) | yes | Black Mageweave Gloves (10003, -0.65 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.59 DPS) [crafted]; Gilded Handwraps (254021, -2.85 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.3 spell_power points (4.82 DPS) | yes | Star Belt (4329, -0.65 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.09 DPS) [dungeon]; Warsong Sash (16975, -1.35 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.9 spell_power points (5.64 DPS) | yes | Gaze Dreamer Pants (6903, -1.85 DPS) [dungeon]; Abomination Skin Leggings (23173, -1.98 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.05 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.57 DPS) | yes | Gilded Slippers (254001, -2.41 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.95 DPS) [crafted]; Acidic Walkers (9454, -5.17 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.9 spell_power points (3.76 DPS) | yes | Reedknot Ring (9622, -1.56 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.87 DPS) [vendor]; Sludge-Stained Band (286535, -2.82 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.84 DPS) | yes | Reedknot Ring (9622, -0.71 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.95 DPS) [vendor]; Sludge-Stained Band (286535, -1.89 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.47 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.52 DPS) [quest]; Gut Ripper (2164, -7.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (144.5 DPS) | yes | Umbral Wand (5216, -0.22 DPS) [dungeon]; Earthen Rod (9381, -0.30 DPS) [dungeon]; Jaina's Firestarter (13064, -2.84 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 202.4. Weights run: 2.1s. Verify run: 1.4s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.205, intellect=0.315 ± 0.014, crit=0.076 ± 0.003 per rating point (14 rating = 1%, 1.060 per %), hit=0.198 ± 0.002 per rating point (10 rating = 1%, 1.979 per %), spell_haste=not significant (0.529 ± 0.198), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.711 ± 0.205), fire_power=0.290 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.89 DPS) | yes | Dreamweave Circlet (10041, -1.15 DPS) [crafted]; Red Mageweave Headband (10033, -2.27 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.42 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (3.59 DPS) | yes | Mindburst Medallion (11196, -0.40 DPS) [quest]; Horizon Choker (13085, -1.81 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.32 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.7 spell_power points (7.53 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -2.35 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.73 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (6.41 DPS) | yes | Deep Woodlands Cloak (19121, -0.43 DPS) [quest]; Mantle of Lady Falther'ess (23178, -1.64 DPS) [dungeon]; Runecloth Cloak (13860, -1.76 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.3 spell_power points (10.20 DPS) | yes | Robe of the Magi (1716, -0.57 DPS) [world_drop]; Elemental Raiment (9434, -1.73 DPS) [world_drop]; Dreamweave Vest (10021, -1.80 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 9.2 spell_power points (3.71 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -1.79 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (7.77 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.38 DPS) [vendor]; Black Mageweave Gloves (10003, -1.72 DPS) [crafted]; Runecloth Gloves (13863, -1.79 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Cloth Girdle (20166, -0.16 DPS) [rep]; Ghostweave Cord (254073, -0.67 DPS) [crafted]; Satyrmane Sash (17755, -2.14 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.1 spell_power points (10.55 DPS) | yes | Red Mageweave Pants (10009, -3.38 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.76 DPS) [vendor]; Wizardweave Leggings (14132, -4.35 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.68 DPS) | yes | Gilded Sandals (254107, -4.10 DPS) [crafted]; Black Mageweave Boots (10026, -4.36 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.51 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.24 DPS) | yes | Philanthropist's Ring (281635, -0.45 DPS) [quest]; Cyclopean Band (11824, -0.72 DPS) [dungeon]; Runed Ring (862, -2.42 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.84 DPS) | yes | Philanthropist's Ring (281635, -0.04 DPS) [quest]; Cyclopean Band (11824, -0.32 DPS) [dungeon]; Runed Ring (862, -2.02 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.42 DPS) [world_drop]; Rune of the Guard Captain (19120, -4.28 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.16 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -4.21 DPS) [dungeon]; Glowing Brightwood Staff (812, -4.39 DPS) [world_drop]; Blade of Eternal Darkness (17780, -5.44 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.04 DPS) [crafted]; Wand of Allistarj (13065, -3.60 DPS) [world_drop]; Pyric Caduceus (11748, -7.35 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 374.5. Weights run: 2.3s. Verify run: 1.4s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.250, intellect=0.338 ± 0.017, crit=0.134 ± 0.004 per rating point (14 rating = 1%, 1.874 per %), hit=0.322 ± 0.002 per rating point (10 rating = 1%, 3.222 per %), spell_haste=-2.608 ± 0.302, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.675 ± 0.250), fire_power=0.324 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.7 spell_power points (16.43 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.88 DPS) [pvp]; Deathmist Mask (226909, -4.18 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (11.05 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.31 DPS) [quest]; Beads of Ogre Mojo (22149, -2.48 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.0 spell_power points (17.06 DPS) | yes | Burial Shawl (18681, -1.33 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Mantle (231592, -1.61 DPS) [pvp]; Argent Shoulders (19059, -4.50 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 21.9 spell_power points (11.02 DPS) | yes | Crystalline Threaded Cape (20697, -0.29 DPS) [world]; Amplifying Cloak (18350, -1.97 DPS) [dungeon]; Hide of the Wild (18510, -2.28 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.0 spell_power points (24.64 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.48 DPS) [pvp]; Robe of Everlasting Night (18385, -4.76 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -8.68 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.7 spell_power points (12.41 DPS) | yes | Sublime Wristguards (18497, -4.68 DPS) [dungeon]; Runecloth Cuffs (254123, -5.19 DPS) [crafted]; Deathmist Bracers (226907, -6.86 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.7 spell_power points (14.42 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -3.18 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 35.7 spell_power points (17.92 DPS) | yes | Belt of the Archmage (18405, -3.99 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -7.18 DPS) [rep]; Defiler's Cloth Girdle (20163, -8.92 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 39.1 spell_power points (19.66 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.22 DPS) [dungeon]; Outrider's Silk Leggings (22747, -2.36 DPS) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.1 spell_power points (12.09 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.03 DPS) [crafted]; Dragonrider Boots (18102, -0.32 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.69 DPS) [quest]; Maiden's Circle (13001, -2.69 DPS) [world_drop]; Naglering (11669, -9.20 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.51 DPS) [quest]; Maiden's Circle (13001, -1.51 DPS) [world_drop]; Naglering (11669, -8.03 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -3.01 DPS) [quest]; Weakness Analyzer (272438, -3.52 DPS) [vendor]; Talisman of Ascendance (22678, -5.93 DPS, sim-verified) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.68 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -16.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (374.5 DPS) | yes | Bonecreeper Stylus (13938, -0.54 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.00 DPS) [world]; Torch of Light (279246, -22.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

