# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 46.1. Weights run: 2.1s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.051, intellect=-0.045 ± 0.003, crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.321 per %), hit=0.083 ± 0.001 per rating point (10 rating = 1%, 0.834 per %), spell_haste=not significant (0.075 ± 0.051), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.828 ± 0.051

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.40 DPS) | yes | Red Winter Hat (21524, -3.73 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.16 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.28 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.93 DPS) | yes | Feyscale Cloak (6632, -0.23 DPS) [dungeon]; Caretaker's Cape (20428, -0.23 DPS) [rep]; Black Whelp Cloak (7283, -0.38 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.16 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -1.67 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.23 DPS) | yes | Ivycloth Bracelets (9793, -0.37 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.63 DPS) | yes | Gnoll Casting Gloves (892, -0.40 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.70 DPS) [crafted]; Heavy Woolen Gloves (4310, -1.16 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (46.1 DPS) | yes | Novice Ardent's Sash (253887, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.12 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (2.09 DPS) | yes | Silk-threaded Trousers (1929, -0.59 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Rumpled Kilt (274741, -0.93 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.63 DPS) | yes | Red Woolen Boots (4313, -0.70 DPS) [crafted]; Feather Padded Treads (285345, -0.77 DPS, sim-verified) [world]; Pristine Boots (253889, -0.93 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (1.16 DPS) | yes | Sludge-Stained Band (286535, -0.47 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.16 DPS) | yes | Sludge-Stained Band (286535, -0.64 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.36 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (1.16 DPS) | yes | Bouquet of Red Roses (22206, -1.51 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 98.9 spell_power points (23.01 DPS) | yes | Skycaller (12984, -2.11 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.72 DPS) [dungeon]; Sizzle Stick (8071, -4.19 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 75.8. Weights run: 2.1s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.077, intellect=0.096 ± 0.005, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.371 per %), hit=0.097 ± 0.001 per rating point (10 rating = 1%, 0.969 per %), spell_haste=0.847 ± 0.074, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.831 ± 0.077

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.71 DPS) | yes | Silk Headband (7050, -0.69 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.74 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.74 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (1.87 DPS) | yes | Crystal Starfire Medallion (5003, -1.77 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.77 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.99 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (2.43 DPS) | yes | Invoker's Mantle (215365, -0.42 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.69 DPS) [dungeon]; Moonlit Amice (11884, -0.71 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.23 DPS) | yes | Prelacy Cape (7004, -0.25 DPS) [quest]; Caretaker's Cape (19533, -0.25 DPS) [rep]; Heavy Woolen Cloak (4311, -0.28 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.21 DPS) | yes | Green Silk Armor (7065, -0.64 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.22 DPS) [dungeon]; Robes of Arcana (5770, -1.23 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.22 DPS) | yes | Glowing Magical Bracelets (13106, -2.03 DPS) [world_drop]; Nightsky Wristbands (6407, -2.08 DPS) [world_drop]; Windsong Bangles (263336, -2.43 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.73 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.25 DPS) [world]; Truefaith Gloves (7049, -0.42 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.3 spell_power points (2.78 DPS) | yes | Belt of Arugal (6392, -0.66 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.81 DPS) [dungeon]; Invoker's Cord (215366, -0.94 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.96 DPS) | yes | Abomination Skin Leggings (23173, -0.60 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.07 DPS) [crafted]; Silk-threaded Trousers (1929, -1.23 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.7 spell_power points (1.89 DPS) | yes | Nimbus Boots (6998, -0.41 DPS) [quest]; Acidic Walkers (9454, -0.47 DPS) [dungeon]; Spidersilk Boots (4320, -1.85 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.73 DPS) | yes | Minor Channeling Ring (1449, -0.45 DPS) [quest]; Electrocutioner Lagnut (9447, -0.99 DPS) [dungeon]; Sludge-Stained Band (286535, -0.99 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.48 DPS) | yes | Electrocutioner Lagnut (9447, -0.74 DPS) [dungeon]; Sludge-Stained Band (286535, -0.74 DPS) [world]; Minor Channeling Ring (1449, -2.28 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.22 DPS) | yes | Twisted Chanter's Staff (890, -1.98 DPS) [world_drop]; Glimmering Staff (249392, -1.99 DPS, sim-verified) [crafted]; Channeler's Staff (4437, -2.03 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.6 spell_power points (1.87 DPS) | yes | Dwarven Tome (279898, -0.58 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.88 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.88 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 137.7 spell_power points (33.95 DPS) | yes | Starfaller (13063, -0.95 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.72 DPS) [crafted]; Gravestone Scepter (7001, -4.95 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 138.7. Weights run: 1.7s. Verify run: 1.1s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.102, intellect=0.176 ± 0.007, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.443 per %), hit=0.121 ± 0.001 per rating point (10 rating = 1%, 1.212 per %), spell_haste=not significant (0.326 ± 0.104), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.850 ± 0.102

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.86 DPS) | yes | Augural Shroud (2620, -2.30 DPS) [world]; Holy Shroud (2721, -2.79 DPS) [world_drop]; Living Cowl (5608, -3.25 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.1 spell_power points (2.25 DPS) | yes | Triune Amulet (7722, -1.90 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.90 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.33 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.6 spell_power points (2.95 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Inquisitor's Shawl (19507, -0.36 DPS) [dungeon]; Berylline Pads (4197, -0.51 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 10.6 spell_power points (2.95 DPS) | yes | Long Silken Cloak (4326, -1.03 DPS) [crafted]; Guardian Cloak (5965, -1.03 DPS) [crafted]; Icy Cloak (4327, -2.59 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (6.43 DPS) | yes | Elemental Raiment (9434, -0.89 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.97 DPS) [crafted]; Robe of Power (7054, -1.94 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.51 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest]; Earthen Silk Cuffs (254019, -1.39 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (5.22 DPS) | yes | Red Mageweave Gloves (10018, -1.66 DPS) [crafted]; Black Mageweave Gloves (10003, -2.01 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.64 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.7 spell_power points (4.10 DPS) | yes | Star Belt (4329, -0.99 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.41 DPS) [dungeon]; Belt of Arugal (6392, -1.44 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.1 spell_power points (4.49 DPS) | yes | Abomination Skin Leggings (23173, -1.59 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.62 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.62 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.69 DPS) | yes | Gilded Slippers (254001, -3.88 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.55 DPS) [crafted]; Acidic Walkers (9454, -4.91 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.1 spell_power points (3.08 DPS) | yes | Ring of Forlorn Spirits (2043, -0.85 DPS) [quest]; Reedknot Ring (9622, -1.13 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.41 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.51 DPS) | yes | Reedknot Ring (9622, -0.56 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor]; Ring of Forlorn Spirits (2043, -0.99 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.07 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.70 DPS) [quest]; Gut Ripper (2164, -9.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (138.7 DPS) | yes | Umbral Wand (5216, -0.00 DPS) [dungeon]; Earthen Rod (9381, -0.08 DPS) [dungeon]; Jaina's Firestarter (13064, -1.89 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 236.4. Weights run: 2.0s. Verify run: 1.4s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.126, intellect=0.148 ± 0.011, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.391 per %), hit=0.184 ± 0.002 per rating point (10 rating = 1%, 1.840 per %), spell_haste=not significant (-0.500 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.909 ± 0.126

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.81 DPS) | yes | Red Mageweave Headband (10033, -1.83 DPS) [crafted]; Dreamweave Circlet (10041, -2.16 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.18 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (2.87 DPS) | yes | Mindburst Medallion (11196, -0.53 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.11 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.33 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (5.79 DPS) | yes | Black Mageweave Shoulders (10027, -1.67 DPS) [crafted]; Bloodmage Mantle (7684, -2.03 DPS) [dungeon]; Rotgrip Mantle (17732, -2.53 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.9 spell_power points (5.41 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.57 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.71 DPS) [crafted]; Nightfall Drape (12465, -2.14 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (8.32 DPS) | yes | Elemental Raiment (9434, -0.69 DPS) [world_drop]; Acumen Robes (17775, -0.76 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -1.29 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.27 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.35 DPS) [crafted]; Condor Bracers (15864, -0.73 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (6.76 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.55 DPS) [vendor]; Runecloth Gloves (13863, -1.91 DPS) [crafted]; Black Mageweave Gloves (10003, -2.04 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.5 spell_power points (5.63 DPS) | yes | Ghostweave Cord (254073, -0.54 DPS) [crafted]; Ban'thok Sash (11662, -0.61 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.85 DPS, sim-verified) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 24.5 spell_power points (8.90 DPS) | yes | Red Mageweave Pants (10009, -3.16 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.75 DPS) [vendor]; Wizardweave Leggings (14132, -4.87 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.72 DPS) | yes | Gilded Sandals (254107, -1.23 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -4.35 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.66 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.72 DPS) | yes | Philanthropist's Ring (281635, -0.77 DPS) [quest]; Cyclopean Band (11824, -1.08 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.82 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.36 DPS) | yes | Cyclopean Band (11824, -0.71 DPS) [dungeon]; Philanthropist's Ring (281635, -1.08 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.45 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.11 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -4.00 DPS) [dungeon]; Arbiter's Blade (11784, -4.09 DPS) [dungeon]; Blade of Eternal Darkness (17780, -9.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (236.4 DPS) | yes | Lesser Eternal Wand (249232, -3.16 DPS) [crafted]; Wand of Allistarj (13065, -3.40 DPS) [world_drop]; Pyric Caduceus (11748, -7.73 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 421.6. Weights run: 2.1s. Verify run: 1.4s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.146, intellect=0.108 ± 0.013, crit=0.069 ± 0.002 per rating point (14 rating = 1%, 0.967 per %), hit=0.293 ± 0.002 per rating point (10 rating = 1%, 2.930 per %), spell_haste=not significant (-0.467 ± 0.201), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.902 ± 0.146

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.9 spell_power points (12.15 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -2.74 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.56 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.66 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.20 DPS) [quest]; Kezan's Taint (19604, -2.81 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.6 spell_power points (11.65 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.08 DPS) [pvp]; Argent Shoulders (19059, -2.93 DPS, sim-verified) [crafted]; Burial Shawl (18681, -3.09 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.4 spell_power points (8.04 DPS) | yes | Amplifying Cloak (18350, -0.96 DPS) [dungeon]; Arcanoweave Cloak (272411, -1.25 DPS, sim-verified) [vendor]; Hide of the Wild (18510, -2.11 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (18.49 DPS) | yes | Robe of Everlasting Night (18385, -3.90 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -4.87 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.80 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.9 spell_power points (9.00 DPS) | yes | Sublime Wristguards (18497, -3.85 DPS) [dungeon]; Runecloth Cuffs (254123, -4.25 DPS) [crafted]; Arcane Runed Bracers (4744, -5.46 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.83 DPS) [quest]; Sandworm Skin Gloves (20716, -4.44 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.3 spell_power points (11.93 DPS) | yes | Belt of the Archmage (18405, -3.37 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -4.42 DPS) [rep]; Oddly Magical Belt (18475, -5.63 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -1.89 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.45 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.06 DPS) [dungeon]; Venomspew Footpads (275606, -1.08 DPS, sim-verified) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.75 DPS) [quest]; Maiden's Circle (13001, -1.75 DPS) [world_drop]; Naglering (11669, -8.96 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.18 DPS) [quest]; Maiden's Circle (13001, -1.18 DPS) [world_drop]; Naglering (11669, -8.19 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.76 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.01 DPS, sim-verified) [quest]; Serenity Field (272439, -5.90 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.42 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+21.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.05 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.50 DPS) [world]; Torch of Light (279246, -21.19 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 43.2. Weights run: 2.1s. Verify run: 1.3s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.051, intellect=-0.045 ± 0.003, crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.321 per %), hit=0.083 ± 0.001 per rating point (10 rating = 1%, 0.834 per %), spell_haste=not significant (0.075 ± 0.051), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.828 ± 0.051

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.40 DPS) | yes | Red Winter Hat (21524, -3.40 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.16 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.23 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.93 DPS) | yes | Feyscale Cloak (6632, -0.23 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.23 DPS) [rep]; Black Whelp Cloak (7283, -0.26 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (1.16 DPS) | yes | Green Woolen Vest (2582, -0.23 DPS) [crafted]; Bloody Apron (6226, -0.23 DPS) [dungeon]; Gray Woolen Robe (2585, -1.55 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) (or Owlbeard Bracers (16981)) | Breaking the Breaker [quest] | 1.0 spell_power points (0.23 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.63 DPS) | yes | Gnoll Casting Gloves (892, -0.28 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.70 DPS) [quest]; Pristine Gloves (253913, -0.70 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (43.2 DPS) | yes | Novice Ardent's Sash (253887, -0.47 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.27 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (2.09 DPS) | yes | Silk-threaded Trousers (1929, -0.32 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.70 DPS) [crafted]; Rumpled Kilt (274741, -0.93 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.63 DPS) | yes | Feather Padded Treads (285345, -0.28 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.70 DPS) [crafted]; Pristine Boots (253889, -0.93 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.16 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.70 DPS) | yes | Ring of the Shadow (1462, -0.85 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 8 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 98.9 spell_power points (23.01 DPS) | yes | Skycaller (12984, -1.70 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.72 DPS) [dungeon]; Sizzle Stick (8071, -4.19 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 74.8. Weights run: 2.1s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.077, intellect=0.096 ± 0.005, crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.371 per %), hit=0.097 ± 0.001 per rating point (10 rating = 1%, 0.969 per %), spell_haste=0.847 ± 0.074, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.831 ± 0.077

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.71 DPS) | yes | Silk Headband (7050, -0.68 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.74 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.74 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (1.87 DPS) | yes | Crystal Starfire Medallion (5003, -1.77 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.77 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.32 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (2.43 DPS) | yes | Invoker's Mantle (215365, -0.59 DPS) [crafted]; Death Speaker Mantle (6685, -0.69 DPS) [dungeon]; Chestnut Mantle (17695, -0.99 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.23 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.25 DPS) [crafted]; Battle Healer's Cloak (19529, -0.25 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.21 DPS) | yes | Green Silk Armor (7065, -0.85 DPS, sim-verified) [crafted]; High Robe of the Adjudicator (3461, -1.19 DPS) [quest]; Death Speaker Robes (6682, -1.22 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.22 DPS) | yes | Windsong Bangles (263336, -1.97 DPS) [quest]; Glowing Magical Bracelets (13106, -2.03 DPS) [world_drop]; Owlbeard Bracers (16981, -2.57 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.73 DPS) | yes | Jutebraid Gloves (10654, -0.13 DPS) [quest]; Gnoll Casting Gloves (892, -0.25 DPS) [world]; Truefaith Gloves (7049, -0.42 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.3 spell_power points (2.78 DPS) | yes | Warsong Sash (16975, -0.07 DPS) [quest]; Belt of Arugal (6392, -0.49 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.81 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.96 DPS) | yes | Abomination Skin Leggings (23173, -0.49 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.07 DPS) [crafted]; Silk-threaded Trousers (1929, -1.23 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.7 spell_power points (1.89 DPS) | yes | Acidic Walkers (9454, -0.47 DPS) [dungeon]; Boots of the Enchanter (4325, -0.66 DPS) [crafted]; Spidersilk Boots (4320, -2.02 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.73 DPS) | yes | Electrocutioner Lagnut (9447, -0.99 DPS) [dungeon]; Sludge-Stained Band (286535, -0.99 DPS) [world]; Sacred Band (6669, -1.23 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.48 DPS) | yes | Electrocutioner Lagnut (9447, -0.74 DPS) [dungeon]; Sacred Band (6669, -0.99 DPS) [quest]; Sludge-Stained Band (286535, -2.87 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.22 DPS) | yes | Twisted Chanter's Staff (890, -1.98 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.98 DPS) [quest]; Glimmering Staff (249392, -2.14 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.6 spell_power points (1.87 DPS) | yes | Orb of Souls (249395, -0.88 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -1.25 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -1.28 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 137.7 spell_power points (33.95 DPS) | yes | Starfaller (13063, -0.99 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.72 DPS) [crafted]; Gravestone Scepter (7001, -4.95 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 136.6. Weights run: 1.7s. Verify run: 1.1s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.102, intellect=0.176 ± 0.007, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.443 per %), hit=0.121 ± 0.001 per rating point (10 rating = 1%, 1.212 per %), spell_haste=not significant (0.326 ± 0.104), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.850 ± 0.102

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.86 DPS) | yes | Augural Shroud (2620, -2.30 DPS) [world]; Living Cowl (5608, -2.52 DPS, sim-verified) [world]; Holy Shroud (2721, -2.79 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.1 spell_power points (2.25 DPS) | yes | Triune Amulet (7722, -1.90 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.90 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.00 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.6 spell_power points (2.95 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Inquisitor's Shawl (19507, -0.36 DPS) [dungeon]; Berylline Pads (4197, -0.51 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 10.6 spell_power points (2.95 DPS) | yes | Long Silken Cloak (4326, -1.03 DPS) [crafted]; Guardian Cloak (5965, -1.03 DPS) [crafted]; Icy Cloak (4327, -1.49 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (6.43 DPS) | yes | Elemental Raiment (9434, -0.57 DPS) [world_drop]; Dreamweave Vest (10021, -0.97 DPS) [crafted]; Robe of Power (7054, -1.94 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.51 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.61 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.00 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (5.22 DPS) | yes | Black Mageweave Gloves (10003, -0.86 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.66 DPS) [crafted]; Gilded Handwraps (254021, -2.64 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.7 spell_power points (4.10 DPS) | yes | Star Belt (4329, -0.48 DPS) [crafted]; Warsong Sash (16975, -1.03 DPS) [quest]; Deathmage Sash (10771, -1.41 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.1 spell_power points (4.49 DPS) | yes | Gaze Dreamer Pants (6903, -0.39 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.59 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.62 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.69 DPS) | yes | Gilded Slippers (254001, -3.09 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.55 DPS) [crafted]; Acidic Walkers (9454, -4.91 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.1 spell_power points (3.08 DPS) | yes | Reedknot Ring (9622, -1.13 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.41 DPS) [vendor]; Sludge-Stained Band (286535, -2.25 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.51 DPS) | yes | Reedknot Ring (9622, -0.61 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor]; Sludge-Stained Band (286535, -1.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.07 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.70 DPS) [quest]; Gut Ripper (2164, -8.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (136.6 DPS) | yes | Umbral Wand (5216, -0.00 DPS) [dungeon]; Earthen Rod (9381, -0.08 DPS) [dungeon]; Jaina's Firestarter (13064, -2.18 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 233.9. Weights run: 2.0s. Verify run: 1.4s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.126, intellect=0.148 ± 0.011, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.391 per %), hit=0.184 ± 0.002 per rating point (10 rating = 1%, 1.840 per %), spell_haste=not significant (-0.500 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.909 ± 0.126

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.81 DPS) | yes | Red Mageweave Headband (10033, -1.83 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.18 DPS) [crafted]; Dreamweave Circlet (10041, -3.35 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (2.87 DPS) | yes | Mindburst Medallion (11196, -0.36 DPS) [quest]; Horizon Choker (13085, -2.11 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.33 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (5.79 DPS) | yes | Black Mageweave Shoulders (10027, -1.67 DPS) [crafted]; Bloodmage Mantle (7684, -2.03 DPS) [dungeon]; Rotgrip Mantle (17732, -2.16 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.9 spell_power points (5.41 DPS) | yes | Deep Woodlands Cloak (19121, -0.68 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.66 DPS) [dungeon]; Runecloth Cloak (13860, -1.71 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (8.32 DPS) | yes | Elemental Raiment (9434, -0.69 DPS) [world_drop]; Dreamweave Vest (10021, -1.29 DPS) [crafted]; Acumen Robes (17775, -1.63 DPS, sim-verified) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.27 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (6.76 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.55 DPS) [vendor]; Runecloth Gloves (13863, -1.91 DPS) [crafted]; Black Mageweave Gloves (10003, -1.93 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.5 spell_power points (5.63 DPS) | yes | Ghostweave Cord (254073, -0.54 DPS) [crafted]; Ban'thok Sash (11662, -0.61 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.10 DPS, sim-verified) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 24.5 spell_power points (8.90 DPS) | yes | Red Mageweave Pants (10009, -3.16 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.75 DPS) [vendor]; Wizardweave Leggings (14132, -5.49 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.72 DPS) | yes | Gilded Sandals (254107, -1.24 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -4.35 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.66 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.72 DPS) | yes | Philanthropist's Ring (281635, -0.77 DPS) [quest]; Cyclopean Band (11824, -1.08 DPS) [dungeon]; Runed Ring (862, -2.18 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.36 DPS) | yes | Cyclopean Band (11824, -0.71 DPS) [dungeon]; Philanthropist's Ring (281635, -0.74 DPS, sim-verified) [quest]; Runed Ring (862, -1.82 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.18 DPS) [world_drop]; Rune of the Guard Captain (19120, -3.89 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.39 DPS, sim-verified) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -4.00 DPS) [dungeon]; Arbiter's Blade (11784, -4.09 DPS) [dungeon]; Blade of Eternal Darkness (17780, -9.30 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (233.9 DPS) | yes | Lesser Eternal Wand (249232, -3.16 DPS) [crafted]; Wand of Allistarj (13065, -3.40 DPS) [world_drop]; Pyric Caduceus (11748, -8.29 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 412.9. Weights run: 2.1s. Verify run: 1.4s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.146, intellect=0.108 ± 0.013, crit=0.069 ± 0.002 per rating point (14 rating = 1%, 0.967 per %), hit=0.293 ± 0.002 per rating point (10 rating = 1%, 2.930 per %), spell_haste=not significant (-0.467 ± 0.201), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.902 ± 0.146

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.9 spell_power points (12.15 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -2.74 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.07 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.66 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.20 DPS) [quest]; Kezan's Taint (19604, -2.81 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.6 spell_power points (11.65 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.08 DPS) [pvp]; Argent Shoulders (19059, -1.81 DPS) [crafted]; Burial Shawl (18681, -3.09 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.4 spell_power points (8.04 DPS) | yes | Arcanoweave Cloak (272411, -0.25 DPS) [vendor]; Amplifying Cloak (18350, -0.96 DPS) [dungeon]; Hide of the Wild (18510, -2.11 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.0 spell_power points (18.49 DPS) | yes | Robe of Everlasting Night (18385, -4.55 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -4.87 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -7.80 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.9 spell_power points (9.00 DPS) | yes | Sublime Wristguards (18497, -3.85 DPS) [dungeon]; Runecloth Cuffs (254123, -4.25 DPS) [crafted]; Spidertank Oilrag (9448, -5.46 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.5 spell_power points (10.84 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.18 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.3 spell_power points (11.93 DPS) | yes | Frostwolf Cloth Belt (19090, -4.42 DPS) [rep]; Belt of the Archmage (18405, -4.52 DPS, sim-verified) [crafted]; Oddly Magical Belt (18475, -5.63 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.89 DPS) [rep]; Sentinel's Silk Leggings (237815, -4.40 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.45 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.06 DPS) [dungeon]; Venomspew Footpads (275606, -1.38 DPS, sim-verified) [crafted] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.75 DPS) [quest]; Maiden's Circle (13001, -1.75 DPS) [world_drop]; Naglering (11669, -8.89 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.18 DPS) [quest]; Maiden's Circle (13001, -1.18 DPS) [world_drop]; Naglering (11669, -8.10 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -1.61 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.76 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.35 DPS, sim-verified) [quest]; Serenity Field (272439, -5.90 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.42 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -17.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+22.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.05 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.50 DPS) [world]; Torch of Light (279246, -22.43 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

