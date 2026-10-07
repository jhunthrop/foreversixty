# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.6. Weights run: 2.1s. Verify run: 1.4s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.072, intellect=-0.054 ± 0.005, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.480 per %), hit=0.133 ± 0.005 per rating point (10 rating = 1%, 1.326 per %), spell_haste=not significant (-0.047 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.733 ± 0.072

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.93 DPS) | yes | Red Winter Hat (21524, -3.06 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.15 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.62 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Caretaker's Cape (20428, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.77 DPS) | yes | Green Woolen Vest (2582, -0.15 DPS) [crafted]; Bloody Apron (6226, -0.15 DPS) [dungeon]; Gray Woolen Robe (2585, -1.32 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Ivycloth Bracelets (9793, -0.26 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.08 DPS) | yes | Gnoll Casting Gloves (892, -0.30 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.46 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.77 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (38.6 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.94 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted]; Silk-threaded Trousers (1929, -0.57 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.62 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.08 DPS) | yes | Red Woolen Boots (4313, -0.46 DPS) [crafted]; Feather Padded Treads (285345, -0.48 DPS, sim-verified) [world]; Pristine Boots (253889, -0.62 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.77 DPS) | yes | Sludge-Stained Band (286535, -0.31 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.77 DPS) | yes | Sludge-Stained Band (286535, -0.45 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.28 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.77 DPS) | yes | Bouquet of Red Roses (22206, -1.22 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 147.2 spell_power points (22.77 DPS) | yes | Skycaller (12984, -1.49 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.48 DPS) [dungeon]; Deepblaze (279896, -4.24 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 25552200000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 63.4. Weights run: 2.1s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.059, intellect=0.127 ± 0.004, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.367 per %), hit=0.104 ± 0.004 per rating point (10 rating = 1%, 1.039 per %), spell_haste=not significant (-0.140 ± 0.063), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.826 ± 0.059

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.78 DPS) | yes | Silk Headband (7050, -0.70 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.76 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.76 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.8 spell_power points (1.96 DPS) | yes | Crystal Starfire Medallion (5003, -1.83 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.83 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.20 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.1 spell_power points (2.57 DPS) | yes | Invoker's Mantle (215365, -0.64 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.69 DPS) [dungeon]; Fairywing Mantle (9536, -0.76 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.26 DPS) | yes | Prelacy Cape (7004, -0.25 DPS) [quest]; Caretaker's Cape (19533, -0.25 DPS) [rep]; Heavy Woolen Cloak (4311, -0.37 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.29 DPS) | yes | Green Silk Armor (7065, -1.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.16 DPS) [dungeon]; Robes of Arcana (5770, -1.26 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.28 DPS) | yes | Windsong Bangles (263336, -2.02 DPS) [quest]; Nightsky Wristbands (6407, -2.08 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.66 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.77 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.25 DPS) [world]; Town Clerk's Mittens (270029, -0.40 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.4 spell_power points (2.88 DPS) | yes | Belt of Arugal (6392, -0.68 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.86 DPS) [dungeon]; Invoker's Cord (215366, -0.95 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.03 DPS) | yes | Abomination Skin Leggings (23173, -0.84 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.04 DPS) [crafted]; Silk-threaded Trousers (1929, -1.26 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.9 spell_power points (2.00 DPS) | yes | Acidic Walkers (9454, -0.47 DPS) [dungeon]; Nimbus Boots (6998, -0.48 DPS) [quest]; Spidersilk Boots (4320, -1.89 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.77 DPS) | yes | Minor Channeling Ring (1449, -0.44 DPS) [quest]; Electrocutioner Lagnut (9447, -1.01 DPS) [dungeon]; Sludge-Stained Band (286535, -1.01 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.52 DPS) | yes | Electrocutioner Lagnut (9447, -0.76 DPS) [dungeon]; Sludge-Stained Band (286535, -0.76 DPS) [world]; Minor Channeling Ring (1449, -2.11 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.28 DPS) | yes | Twisted Chanter's Staff (890, -1.95 DPS) [world_drop]; Glimmering Staff (249392, -2.00 DPS, sim-verified) [crafted]; Channeler's Staff (4437, -2.02 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.8 spell_power points (1.96 DPS) | yes | Dwarven Tome (279898, -0.56 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.95 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.95 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 134.3 spell_power points (33.97 DPS) | yes | Starfaller (13063, -0.96 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.70 DPS) [crafted]; Gravestone Scepter (7001, -4.97 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 25552300120201300-0000000000000000000-0000000000000000)

Set DPS (verified): 110.1. Weights run: 1.7s. Verify run: 1.0s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.118, intellect=0.167 ± 0.008, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.514 per %), hit=0.157 ± 0.007 per rating point (10 rating = 1%, 1.567 per %), spell_haste=-1.356 ± 0.121, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.813 ± 0.118

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.96 DPS) | yes | Augural Shroud (2620, -1.97 DPS) [world]; Living Cowl (5608, -2.25 DPS, sim-verified) [world]; Holy Shroud (2721, -2.36 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (1.89 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.44 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.61 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.61 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (2.48 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Inquisitor's Shawl (19507, -0.31 DPS) [dungeon]; Berylline Pads (4197, -0.43 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 10.5 spell_power points (2.48 DPS) | yes | Long Silken Cloak (4326, -0.87 DPS) [crafted]; Guardian Cloak (5965, -0.87 DPS) [crafted]; Icy Cloak (4327, -1.27 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (5.44 DPS) | yes | Elemental Raiment (9434, -0.39 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.83 DPS) [crafted]; Robe of Power (7054, -1.65 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.13 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.47 DPS) [quest]; Earthen Silk Cuffs (254019, -1.18 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (4.41 DPS) | yes | Black Mageweave Gloves (10003, -0.55 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.42 DPS) [crafted]; Gilded Handwraps (254021, -2.24 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.7 spell_power points (3.47 DPS) | yes | Star Belt (4329, -0.39 DPS) [crafted]; Deathmage Sash (10771, -1.22 DPS) [dungeon]; Belt of Arugal (6392, -1.22 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.0 spell_power points (3.78 DPS) | yes | Gaze Dreamer Pants (6903, -0.90 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.34 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.38 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.67 DPS) | yes | Gilded Slippers (254001, -2.36 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.86 DPS) [crafted]; Acidic Walkers (9454, -4.17 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (2.60 DPS) | yes | Ring of Forlorn Spirits (2043, -0.71 DPS) [quest]; Reedknot Ring (9622, -0.95 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.18 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.13 DPS) | yes | Ring of Forlorn Spirits (2043, -0.41 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.47 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.71 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (110.1 DPS) | yes | Scorn's Focal Dagger (23168, -2.60 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.92 DPS) [quest]; Gut Ripper (2164, -7.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 167.7 spell_power points (39.62 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.03 DPS) [dungeon]; Twisted Nether Wand (249144, -4.20 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25552300120201351-2002000000000000000-0000000000000000)

Set DPS (verified): 232.9. Weights run: 2.0s. Verify run: 1.3s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.147, intellect=0.265 ± 0.012, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.411 per %), hit=0.260 ± 0.014 per rating point (10 rating = 1%, 2.602 per %), spell_haste=not significant (0.131 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.895 ± 0.147

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.53 DPS) | yes | Dreamweave Circlet (10041, -1.18 DPS) [crafted]; Red Mageweave Headband (10033, -1.74 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.12 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.6 spell_power points (3.03 DPS) | yes | Mindburst Medallion (11196, -0.35 DPS) [quest]; Horizon Choker (13085, -1.72 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.10 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -1.79 DPS) [crafted]; Bloodmage Mantle (7684, -2.14 DPS) [dungeon]; Rotgrip Mantle (17732, -2.30 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (5.50 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.61 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.58 DPS) [crafted]; Nightfall Drape (12465, -2.33 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.3 spell_power points (8.58 DPS) | yes | Robe of the Magi (1716, -0.25 DPS) [world_drop]; Elemental Raiment (9434, -1.17 DPS) [world_drop]; Dreamweave Vest (10021, -1.38 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.18 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Bloodband Bracers (11469, -0.57 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (6.73 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.39 DPS, sim-verified) [vendor]; Black Mageweave Gloves (10003, -1.43 DPS) [crafted]; Runecloth Gloves (13863, -1.65 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.5 spell_power points (6.18 DPS) | yes | Highlander's Cloth Girdle (20098, -0.87 DPS) [rep]; Ghostweave Cord (254073, -1.24 DPS) [crafted]; Satyrmane Sash (17755, -1.80 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.7 spell_power points (9.05 DPS) | yes | Red Mageweave Pants (10009, -2.99 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.55 DPS) [vendor]; Wizardweave Leggings (14132, -3.79 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.47 DPS) | yes | Gilded Sandals (254107, -3.75 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.89 DPS) [vendor]; Black Mageweave Boots (10026, -3.93 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.59 DPS) | yes | Philanthropist's Ring (281635, -0.50 DPS) [quest]; Cyclopean Band (11824, -0.76 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.76 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.24 DPS) | yes | Philanthropist's Ring (281635, -0.14 DPS) [quest]; Cyclopean Band (11824, -0.40 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.41 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.90 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -3.77 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.88 DPS) [dungeon]; Blade of Eternal Darkness (17780, -7.67 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.19 DPS) [crafted]; Wand of Allistarj (13065, -3.34 DPS) [world_drop]; Pyric Caduceus (11748, -9.21 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 481.4. Weights run: 5.7s. Verify run: 1.4s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.322, intellect=-0.072 ± 0.019, crit=0.081 ± 0.002 per rating point (14 rating = 1%, 1.134 per %), hit=0.505 ± 0.033 per rating point (10 rating = 1%, 5.052 per %), spell_haste=1.362 ± 0.276, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.883 ± 0.322)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.0 spell_power points (10.33 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -2.39 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -5.25 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.58 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.41 DPS) [quest]; Diana's Pearl Necklace (22403, -2.74 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.1 spell_power points (9.69 DPS) | yes | Argent Shoulders (19059, -1.08 DPS) [crafted]; Field Marshal's Dreadweave Shoulders (231583, -1.08 DPS) [pvp]; Burial Shawl (18681, -2.80 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 21.1 spell_power points (7.25 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -1.05 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -2.43 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.0 spell_power points (15.84 DPS) | yes | Robe of Everlasting Night (18385, -2.93 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -4.82 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -7.23 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.0 spell_power points (7.58 DPS) | yes | Sublime Wristguards (18497, -3.44 DPS) [dungeon]; Runecloth Cuffs (254123, -3.79 DPS) [crafted]; Arcane Runed Bracers (4744, -4.48 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.38 DPS) [quest]; Sandworm Skin Gloves (20716, -5.40 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.1 spell_power points (10.35 DPS) | yes | Belt of the Archmage (18405, -3.93 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -4.15 DPS) [rep]; Ban'thok Sash (11662, -4.48 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -2.07 DPS) [pvp]; Sentinel's Silk Leggings (237815, -7.00 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.27 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.69 DPS) [crafted]; Omnicast Boots (11822, -1.38 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.34 DPS) [quest]; Maiden's Circle (13001, -1.38 DPS) [world_drop]; Naglering (11669, -7.20 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.02 DPS) [quest]; Maiden's Circle (13001, -1.05 DPS) [world_drop]; Naglering (11669, -14.64 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+20.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.07 DPS) [quest]; Weakness Analyzer (272438, -2.41 DPS) [vendor]; Serenity Field (272439, -5.17 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -1.93 DPS, sim-verified) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.36 DPS) [world]; Teebu's Blazing Longsword (1728, -22.48 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+23.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.22 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.59 DPS) [world]; Torch of Light (279246, -23.03 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 785.8. Weights run: 6.1s. Verify run: 1.4s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.121, intellect=0.064 ± 0.018, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.901 per %), hit=0.450 ± 0.032 per rating point (10 rating = 1%, 4.499 per %), spell_haste=2.453 ± 0.340, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.924 ± 0.121

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.5 spell_power points (22.14 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -4.69 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -9.17 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (15.96 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -4.47 DPS) [quest]; Kezan's Taint (19604, -5.43 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.9 spell_power points (20.95 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.01 DPS) [pvp]; Burial Shawl (18681, -5.69 DPS) [dungeon]; Argent Shoulders (19059, -6.00 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 21.0 spell_power points (15.25 DPS) | yes | Crystalline Threaded Cape (20697, -0.55 DPS) [world]; Amplifying Cloak (18350, -2.19 DPS) [dungeon]; Hide of the Wild (18510, -4.62 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.6 spell_power points (33.80 DPS) | yes | Robe of Everlasting Night (18385, -4.90 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -9.46 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -14.72 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.5 spell_power points (16.34 DPS) | yes | Sublime Wristguards (18497, -7.16 DPS) [dungeon]; Runecloth Cuffs (254123, -7.89 DPS) [crafted]; Arcane Runed Bracers (4744, -9.81 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (19.83 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.86 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.9 spell_power points (22.44 DPS) | yes | Belt of the Archmage (18405, -7.33 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -8.91 DPS) [rep]; Ban'thok Sash (11662, -9.95 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 34.5 spell_power points (25.05 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -3.84 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (17.42 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.45 DPS) [crafted]; Omnicast Boots (11822, -2.34 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.60 DPS) [dungeon]; Maiden's Circle (13001, -3.09 DPS) [world_drop]; Naglering (11669, -17.32 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.69 DPS) [dungeon]; Maiden's Circle (13001, -2.18 DPS) [world_drop]; Naglering (11669, -15.23 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+22.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -3.65 DPS, sim-verified) [quest]; Weakness Analyzer (272438, -5.08 DPS) [vendor]; Serenity Field (272439, -10.89 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.41 DPS) [world]; Teebu's Blazing Longsword (1728, -31.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (785.8 DPS) | yes | Bonecreeper Stylus (13938, -1.03 DPS) [dungeon]; Brilliant Wand (249385, -5.39 DPS) [crafted]; Torch of Light (279246, -23.99 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.3. Weights run: 2.1s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.072, intellect=-0.054 ± 0.005, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.480 per %), hit=0.133 ± 0.005 per rating point (10 rating = 1%, 1.326 per %), spell_haste=not significant (-0.047 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.733 ± 0.072

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.93 DPS) | yes | Red Winter Hat (21524, -2.75 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.15 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.62 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.77 DPS) | yes | Green Woolen Vest (2582, -0.15 DPS) [crafted]; Bloody Apron (6226, -0.15 DPS) [dungeon]; Gray Woolen Robe (2585, -1.24 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) (or Owlbeard Bracers (16981)) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.08 DPS) | yes | Gnoll Casting Gloves (892, -0.29 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.46 DPS) [quest]; Pristine Gloves (253913, -0.46 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (36.3 DPS) | yes | Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Filigreed Pristine Leggings (253937, -0.46 DPS) [crafted]; Silk-threaded Trousers (1929, -0.60 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.62 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.08 DPS) | yes | Red Woolen Boots (4313, -0.46 DPS) [crafted]; Feather Padded Treads (285345, -0.55 DPS, sim-verified) [world]; Pristine Boots (253889, -0.62 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.77 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.46 DPS) | yes | Ring of the Shadow (1462, -0.65 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 8 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 147.2 spell_power points (22.77 DPS) | yes | Skycaller (12984, -1.62 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.48 DPS) [dungeon]; Sizzle Stick (8071, -4.35 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 25552200000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 62.5. Weights run: 2.1s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.059, intellect=0.127 ± 0.004, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.367 per %), hit=0.104 ± 0.004 per rating point (10 rating = 1%, 1.039 per %), spell_haste=not significant (-0.140 ± 0.063), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.826 ± 0.059

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.78 DPS) | yes | Silk Headband (7050, -0.42 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.76 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.76 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.8 spell_power points (1.96 DPS) | yes | Darkspear Warding Pendant (272075, -1.70 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.83 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.83 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.1 spell_power points (2.57 DPS) | yes | Invoker's Mantle (215365, -0.63 DPS) [crafted]; Death Speaker Mantle (6685, -0.69 DPS) [dungeon]; Chestnut Mantle (17695, -1.87 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.26 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.25 DPS) [crafted]; Battle Healer's Cloak (19529, -0.25 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.29 DPS) | yes | Green Silk Armor (7065, -0.91 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.16 DPS) [dungeon]; High Robe of the Adjudicator (3461, -1.20 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.28 DPS) | yes | Owlbeard Bracers (16981, -2.01 DPS, sim-verified) [quest]; Glowing Magical Bracelets (13106, -2.02 DPS) [world_drop]; Windsong Bangles (263336, -2.02 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.77 DPS) | yes | Jutebraid Gloves (10654, -0.09 DPS) [quest]; Gnoll Casting Gloves (892, -0.25 DPS) [world]; Truefaith Gloves (7049, -0.41 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.4 spell_power points (2.88 DPS) | yes | Warsong Sash (16975, -0.10 DPS) [quest]; Belt of Arugal (6392, -0.51 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.86 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.03 DPS) | yes | Abomination Skin Leggings (23173, -0.60 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.04 DPS) [crafted]; Silk-threaded Trousers (1929, -1.26 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.9 spell_power points (2.00 DPS) | yes | Acidic Walkers (9454, -0.47 DPS) [dungeon]; Boots of the Enchanter (4325, -0.73 DPS) [crafted]; Spidersilk Boots (4320, -1.76 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.77 DPS) | yes | Electrocutioner Lagnut (9447, -1.01 DPS) [dungeon]; Sludge-Stained Band (286535, -1.01 DPS) [world]; Sacred Band (6669, -1.26 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.52 DPS) | yes | Electrocutioner Lagnut (9447, -0.76 DPS) [dungeon]; Sacred Band (6669, -1.01 DPS) [quest]; Sludge-Stained Band (286535, -2.19 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.28 DPS) | yes | Glimmering Staff (249392, -1.53 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.95 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.95 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.8 spell_power points (1.96 DPS) | yes | Orb of Souls (249395, -0.95 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.33 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.43 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 134.3 spell_power points (33.97 DPS) | yes | Starfaller (13063, -0.60 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.70 DPS) [crafted]; Gravestone Scepter (7001, -4.97 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552300120201300-0000000000000000000-0000000000000000)

Set DPS (verified): 109.4. Weights run: 1.7s. Verify run: 1.0s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.118, intellect=0.167 ± 0.008, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.514 per %), hit=0.157 ± 0.007 per rating point (10 rating = 1%, 1.567 per %), spell_haste=-1.356 ± 0.121, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.813 ± 0.118

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.96 DPS) | yes | Augural Shroud (2620, -1.97 DPS) [world]; Living Cowl (5608, -2.04 DPS, sim-verified) [world]; Holy Shroud (2721, -2.36 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (1.89 DPS) | yes | Triune Amulet (7722, -1.61 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.61 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.29 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (2.48 DPS) | yes | Inquisitor's Shawl (19507, -0.31 DPS) [dungeon]; Berylline Pads (4197, -0.43 DPS) [quest]; Green Silken Shoulders (7057, -0.51 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 10.5 spell_power points (2.48 DPS) | yes | Long Silken Cloak (4326, -0.87 DPS) [crafted]; Guardian Cloak (5965, -0.87 DPS) [crafted]; Icy Cloak (4327, -1.35 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (5.44 DPS) | yes | Elemental Raiment (9434, -0.47 DPS) [world_drop]; Dreamweave Vest (10021, -0.83 DPS) [crafted]; Robe of Power (7054, -1.65 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.13 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.76 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.87 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (4.41 DPS) | yes | Black Mageweave Gloves (10003, -0.78 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.42 DPS) [crafted]; Gilded Handwraps (254021, -2.24 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.7 spell_power points (3.47 DPS) | yes | Star Belt (4329, -0.40 DPS, sim-verified) [crafted]; Warsong Sash (16975, -0.87 DPS) [quest]; Deathmage Sash (10771, -1.22 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.0 spell_power points (3.78 DPS) | yes | Gaze Dreamer Pants (6903, -0.36 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.34 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.38 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.67 DPS) | yes | Gilded Slippers (254001, -2.87 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.86 DPS) [crafted]; Acidic Walkers (9454, -4.17 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (2.60 DPS) | yes | Reedknot Ring (9622, -0.95 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.18 DPS) [vendor]; Sludge-Stained Band (286535, -1.89 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.13 DPS) | yes | Sea Giant's Toe Ring (274746, -0.71 DPS) [vendor]; Reedknot Ring (9622, -0.76 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -1.42 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (109.4 DPS) | yes | Scorn's Focal Dagger (23168, -2.60 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.92 DPS) [quest]; Gut Ripper (2164, -7.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 167.7 spell_power points (39.62 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.03 DPS) [dungeon]; Twisted Nether Wand (249144, -4.20 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552300120201351-2002000000000000000-0000000000000000)

Set DPS (verified): 231.8. Weights run: 2.0s. Verify run: 1.3s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.147, intellect=0.265 ± 0.012, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.411 per %), hit=0.260 ± 0.014 per rating point (10 rating = 1%, 2.602 per %), spell_haste=not significant (0.131 ± 0.154), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.895 ± 0.147

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.53 DPS) | yes | Dreamweave Circlet (10041, -1.18 DPS) [crafted]; Red Mageweave Headband (10033, -1.48 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.12 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.6 spell_power points (3.03 DPS) | yes | Mindburst Medallion (11196, -0.55 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.72 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.10 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -1.79 DPS) [crafted]; Bloodmage Mantle (7684, -2.14 DPS) [dungeon]; Rotgrip Mantle (17732, -2.72 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.6 spell_power points (5.50 DPS) | yes | Deep Woodlands Cloak (19121, -0.50 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.48 DPS) [dungeon]; Runecloth Cloak (13860, -1.58 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.3 spell_power points (8.58 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.17 DPS) [world_drop]; Dreamweave Vest (10021, -1.38 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.18 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (6.73 DPS) | yes | Black Mageweave Gloves (10003, -1.43 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.46 DPS, sim-verified) [vendor]; Runecloth Gloves (13863, -1.65 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.5 spell_power points (6.18 DPS) | yes | Defiler's Cloth Girdle (20166, -0.87 DPS) [rep]; Ghostweave Cord (254073, -1.24 DPS) [crafted]; Satyrmane Sash (17755, -1.65 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.7 spell_power points (9.05 DPS) | yes | Red Mageweave Pants (10009, -2.99 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.55 DPS) [vendor]; Wizardweave Leggings (14132, -4.25 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.47 DPS) | yes | Gilded Sandals (254107, -0.70 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.89 DPS) [vendor]; Black Mageweave Boots (10026, -3.93 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.59 DPS) | yes | Philanthropist's Ring (281635, -0.50 DPS) [quest]; Cyclopean Band (11824, -0.76 DPS) [dungeon]; Runed Ring (862, -2.12 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.24 DPS) | yes | Cyclopean Band (11824, -0.40 DPS) [dungeon]; Philanthropist's Ring (281635, -0.86 DPS, sim-verified) [quest]; Runed Ring (862, -1.76 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.12 DPS) [world_drop]; Rune of the Guard Captain (19120, -3.59 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.18 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -3.77 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.88 DPS) [dungeon]; Blade of Eternal Darkness (17780, -7.74 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Eternal Wand (249232, -3.19 DPS) [crafted]; Wand of Allistarj (13065, -3.34 DPS) [world_drop]; Pyric Caduceus (11748, -8.75 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 475.2. Weights run: 5.7s. Verify run: 1.4s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.322, intellect=-0.072 ± 0.019, crit=0.081 ± 0.002 per rating point (14 rating = 1%, 1.134 per %), hit=0.505 ± 0.033 per rating point (10 rating = 1%, 5.052 per %), spell_haste=1.362 ± 0.276, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.883 ± 0.322)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.0 spell_power points (10.33 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -2.39 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -3.56 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (7.58 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.41 DPS) [quest]; Diana's Pearl Necklace (22403, -2.74 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.1 spell_power points (9.69 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.08 DPS) [pvp]; Argent Shoulders (19059, -1.97 DPS, sim-verified) [crafted]; Burial Shawl (18681, -2.80 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 21.1 spell_power points (7.25 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -1.05 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -2.43 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.0 spell_power points (15.84 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.82 DPS) [pvp]; Robe of Everlasting Night (18385, -5.31 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -7.23 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.0 spell_power points (7.58 DPS) | yes | Sublime Wristguards (18497, -3.44 DPS) [dungeon]; Runecloth Cuffs (254123, -3.79 DPS) [crafted]; Spidertank Oilrag (9448, -4.48 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.0 spell_power points (9.30 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.72 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.1 spell_power points (10.35 DPS) | yes | Frostwolf Cloth Belt (19090, -4.15 DPS) [rep]; Ban'thok Sash (11662, -4.48 DPS) [dungeon]; Belt of the Archmage (18405, -5.43 DPS, sim-verified) [crafted] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -2.07 DPS) [pvp]; Sentinel's Silk Leggings (237815, -5.46 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.27 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.69 DPS) [crafted]; Omnicast Boots (11822, -1.38 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.34 DPS) [quest]; Maiden's Circle (13001, -1.38 DPS) [world_drop]; Naglering (11669, -11.90 DPS, sim-verified) [dungeon] |
| finger2 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.02 DPS) [quest]; Maiden's Circle (13001, -1.05 DPS) [world_drop]; Naglering (11669, -14.49 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+21.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.41 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.00 DPS, sim-verified) [quest]; Serenity Field (272439, -5.17 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.36 DPS) [world]; Teebu's Blazing Longsword (1728, -24.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+22.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.22 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.59 DPS) [world]; Torch of Light (279246, -22.40 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Rune Band of Wizardry; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 777.8. Weights run: 6.1s. Verify run: 1.5s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.121, intellect=0.064 ± 0.018, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.901 per %), hit=0.450 ± 0.032 per rating point (10 rating = 1%, 4.499 per %), spell_haste=2.453 ± 0.340, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.924 ± 0.121

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.5 spell_power points (22.14 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -4.69 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -9.85 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (15.96 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -4.47 DPS) [quest]; Kezan's Taint (19604, -5.43 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.9 spell_power points (20.95 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.01 DPS) [pvp]; Argent Shoulders (19059, -3.85 DPS, sim-verified) [crafted]; Burial Shawl (18681, -5.69 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 21.0 spell_power points (15.25 DPS) | yes | Amplifying Cloak (18350, -2.19 DPS) [dungeon]; Crystalline Threaded Cape (20697, -2.45 DPS, sim-verified) [world]; Hide of the Wild (18510, -4.62 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.6 spell_power points (33.80 DPS) | yes | Robe of Everlasting Night (18385, -6.26 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -9.46 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -14.72 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.5 spell_power points (16.34 DPS) | yes | Sublime Wristguards (18497, -7.16 DPS) [dungeon]; Runecloth Cuffs (254123, -7.89 DPS) [crafted]; Spidertank Oilrag (9448, -9.81 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (19.83 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.86 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.9 spell_power points (22.44 DPS) | yes | Frostwolf Cloth Belt (19090, -8.91 DPS) [rep]; Belt of the Archmage (18405, -9.76 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -9.95 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 34.5 spell_power points (25.05 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -3.84 DPS) [rep]; Sentinel's Silk Leggings (237815, -5.59 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (17.42 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.45 DPS) [crafted]; Omnicast Boots (11822, -2.34 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.60 DPS) [dungeon]; Maiden's Circle (13001, -3.09 DPS) [world_drop]; Naglering (11669, -15.13 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.69 DPS) [dungeon]; Maiden's Circle (13001, -2.18 DPS) [world_drop]; Naglering (11669, -15.15 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+22.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -4.35 DPS) [quest]; Weakness Analyzer (272438, -5.08 DPS) [vendor]; Serenity Field (272439, -10.89 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.41 DPS) [world]; Teebu's Blazing Longsword (1728, -34.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (777.8 DPS) | yes | Bonecreeper Stylus (13938, -1.03 DPS) [dungeon]; Brilliant Wand (249385, -5.39 DPS) [crafted]; Torch of Light (279246, -24.14 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

