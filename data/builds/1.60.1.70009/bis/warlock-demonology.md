# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 2.4s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.074, intellect=-0.024 ± 0.005, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.481 per %), hit=0.140 ± 0.001 per rating point (10 rating = 1%, 1.397 per %), spell_haste=not significant (0.213 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.672 ± 0.074, fire_power=0.328 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.83 DPS) | yes | Red Winter Hat (21524, -2.62 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.69 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.14 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.55 DPS) | yes | Feyscale Cloak (6632, -0.14 DPS) [dungeon]; Caretaker's Cape (20428, -0.14 DPS) [rep]; Black Whelp Cloak (7283, -0.22 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.69 DPS) | yes | Green Woolen Vest (2582, -0.14 DPS) [crafted]; Bloody Apron (6226, -0.14 DPS) [dungeon]; Gray Woolen Robe (2585, -1.04 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.14 DPS) | yes | Ivycloth Bracelets (9793, -0.22 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.97 DPS) | yes | Gnoll Casting Gloves (892, -0.26 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.41 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.69 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (36.0 DPS) | yes | Novice Ardent's Sash (253887, -0.28 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.95 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.24 DPS) | yes | Filigreed Pristine Leggings (253937, -0.41 DPS) [crafted]; Rumpled Kilt (274741, -0.55 DPS) [vendor]; Silk-threaded Trousers (1929, -0.65 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (0.97 DPS) | yes | Feather Padded Treads (285345, -0.35 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.41 DPS) [crafted]; Pristine Boots (253889, -0.55 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.69 DPS) | yes | Sludge-Stained Band (286535, -0.28 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.69 DPS) | yes | Sludge-Stained Band (286535, -0.40 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -0.94 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.69 DPS) | yes | Bouquet of Red Roses (22206, -1.04 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 164.4 spell_power points (22.72 DPS) | yes | Skycaller (12984, -1.53 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.43 DPS) [dungeon]; Deepblaze (279896, -4.19 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 59.7. Weights run: 2.4s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.080, intellect=0.114 ± 0.006, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.406 per %), hit=0.121 ± 0.001 per rating point (10 rating = 1%, 1.206 per %), spell_haste=0.535 ± 0.088, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.793 ± 0.080, fire_power=0.210 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.37 DPS) | yes | Silk Headband (7050, -0.64 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.65 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.65 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.7 spell_power points (1.66 DPS) | yes | Crystal Starfire Medallion (5003, -1.56 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.69 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.0 spell_power points (2.16 DPS) | yes | Death Speaker Mantle (6685, -0.60 DPS) [dungeon]; Invoker's Mantle (215365, -0.62 DPS, sim-verified) [crafted]; Fairywing Mantle (9536, -0.65 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.08 DPS) | yes | Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest]; Caretaker's Cape (19533, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.80 DPS) | yes | Green Silk Armor (7065, -0.67 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.02 DPS) [dungeon]; Robes of Arcana (5770, -1.08 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.94 DPS) | yes | Glowing Magical Bracelets (13106, -1.74 DPS) [world_drop]; Nightsky Wristbands (6407, -1.79 DPS) [world_drop]; Windsong Bangles (263336, -2.17 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.51 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.36 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.3 spell_power points (2.44 DPS) | yes | Belt of Arugal (6392, -0.62 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.72 DPS) [dungeon]; Invoker's Cord (215366, -0.81 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.58 DPS) | yes | Abomination Skin Leggings (23173, -0.56 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.90 DPS) [crafted]; Silk-threaded Trousers (1929, -1.08 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.8 spell_power points (1.68 DPS) | yes | Nimbus Boots (6998, -0.39 DPS) [quest]; Acidic Walkers (9454, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.66 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.51 DPS) | yes | Minor Channeling Ring (1449, -0.38 DPS) [quest]; Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Sludge-Stained Band (286535, -0.86 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.29 DPS) | yes | Electrocutioner Lagnut (9447, -0.65 DPS) [dungeon]; Sludge-Stained Band (286535, -0.65 DPS) [world]; Minor Channeling Ring (1449, -1.66 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.94 DPS) | yes | Glimmering Staff (249392, -1.48 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.69 DPS) [world_drop]; Channeler's Staff (4437, -1.74 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.7 spell_power points (1.66 DPS) | yes | Dwarven Tome (279898, -0.60 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.79 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.79 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 157.2 spell_power points (33.86 DPS) | yes | Starfaller (13063, -0.65 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.78 DPS) [crafted]; Gravestone Scepter (7001, -4.86 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 130.8. Weights run: 1.9s. Verify run: 1.0s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.084, intellect=0.460 ± 0.016, crit=0.052 ± 0.002 per rating point (14 rating = 1%, 0.733 per %), hit=0.210 ± 0.002 per rating point (10 rating = 1%, 2.101 per %), spell_haste=1.806 ± 0.106, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.748 ± 0.084, fire_power=0.251 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.47 DPS) | yes | Living Cowl (5608, -1.70 DPS) [world]; Holy Shroud (2721, -2.13 DPS) [world_drop]; Augural Shroud (2620, -2.62 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (2.08 DPS) | yes | Triune Amulet (7722, -1.39 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.39 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.83 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.1 spell_power points (2.80 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Inquisitor's Shawl (19507, -0.03 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.1 spell_power points (2.80 DPS) | yes | Guardian Cloak (5965, -1.03 DPS) [crafted]; Icy Cloak (4327, -1.31 DPS) [crafted]; Long Silken Cloak (4326, -1.66 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.8 spell_power points (5.27 DPS) | yes | Elemental Raiment (9434, -0.80 DPS) [world_drop]; Dreamweave Vest (10021, -1.04 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.11 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.91 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.43 DPS) [quest]; Windchaser Cuffs (14429, -1.03 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (4.22 DPS) | yes | Black Mageweave Gloves (10003, -1.03 DPS) [crafted]; Red Mageweave Gloves (10018, -1.55 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.83 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.8 spell_power points (3.37 DPS) | yes | Star Belt (4329, -0.60 DPS) [crafted]; Gilded Cord (254037, -0.89 DPS) [crafted]; Deathmage Sash (10771, -1.84 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.5 spell_power points (4.15 DPS) | yes | Abomination Skin Leggings (23173, -1.46 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.60 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.66 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.11 DPS) | yes | Gilded Slippers (254001, -2.26 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.23 DPS) [crafted]; Acidic Walkers (9454, -3.26 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (2.71 DPS) | yes | Ring of Forlorn Spirits (2043, -1.01 DPS) [quest]; Reedknot Ring (9622, -1.23 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.44 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.91 DPS) | yes | Ring of Forlorn Spirits (2043, -0.21 DPS) [quest]; Reedknot Ring (9622, -0.43 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.64 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (130.8 DPS) | yes | Scorn's Focal Dagger (23168, -2.34 DPS) [dungeon]; Staff of Dar'Orahil (15106, -2.73 DPS) [quest]; Gut Ripper (2164, -7.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 187.9 spell_power points (39.97 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.38 DPS) [dungeon]; Twisted Nether Wand (249144, -4.69 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 179.7. Weights run: 1.8s. Verify run: 1.2s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.064, intellect=0.250 ± 0.012, crit=0.059 ± 0.002 per rating point (14 rating = 1%, 0.820 per %), hit=0.220 ± 0.002 per rating point (10 rating = 1%, 2.203 per %), spell_haste=not significant (-0.062 ± 0.127), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.813 ± 0.064, fire_power=0.187 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.49 DPS) | yes | Dreamweave Circlet (10041, -0.97 DPS) [crafted]; Red Mageweave Headband (10033, -1.31 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.67 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (2.36 DPS) | yes | Mindburst Medallion (11196, -0.28 DPS) [quest]; Horizon Choker (13085, -1.39 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.66 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -1.39 DPS) [crafted]; Bloodmage Mantle (7684, -1.67 DPS) [dungeon]; Rotgrip Mantle (17732, -2.02 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (4.30 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.18 DPS) [dungeon]; Runecloth Cloak (13860, -1.25 DPS) [crafted]; Nightfall Drape (12465, -1.80 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.0 spell_power points (6.66 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.83 DPS) [world_drop]; Dreamweave Vest (10021, -1.04 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.50 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.07 DPS) [crafted]; Bloodband Bracers (11469, -0.49 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (5.27 DPS) | yes | Black Mageweave Gloves (10003, -1.11 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.32 DPS, sim-verified) [vendor]; Runecloth Gloves (13863, -1.32 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (4.58 DPS) | yes | Highlander's Cloth Girdle (20098, -0.42 DPS) [rep]; Ban'thok Sash (11662, -0.42 DPS) [dungeon]; Ghostweave Cord (254073, -0.69 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (7.08 DPS) | yes | Red Mageweave Pants (10009, -2.36 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.69 DPS) [vendor]; Wizardweave Leggings (14132, -3.84 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.66 DPS) | yes | Gilded Sandals (254107, -2.98 DPS) [crafted]; Black Mageweave Boots (10026, -3.12 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.20 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.61 DPS) | yes | Philanthropist's Ring (281635, -0.42 DPS) [quest]; Cyclopean Band (11824, -0.62 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.39 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.33 DPS) | yes | Cyclopean Band (11824, -0.35 DPS) [dungeon]; Philanthropist's Ring (281635, -0.75 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.11 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.44 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.98 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.05 DPS) [dungeon]; Blade of Eternal Darkness (17780, -4.47 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.97 DPS) [world_drop]; Pyric Caduceus (11748, -3.22 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.42 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532000000000000-2332113101220001350-0040000000000000)

Set DPS (verified): 390.9. Weights run: 8.3s. Verify run: 1.2s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± 0.341), intellect=not significant (-0.809 ± 0.025), crit=not significant (-0.183 ± 0.003) per rating point (14 rating = 1%, -2.558 per %), hit=not significant (-0.758 ± 0.004) per rating point (10 rating = 1%, -7.581 per %), spell_haste=not significant (-1.404 ± 0.442), spell_penetration=not significant (-0.000 ± 0.000), shadow_power=not significant (1.292 ± 0.341), fire_power=not significant (-0.292 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.0 spell_power points | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.58 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -6.05 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.04 DPS) [quest]; Beads of Ogre Mojo (22149, -1.66 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.6 spell_power points | yes | Field Marshal's Dreadweave Shoulders (231583, -0.64 DPS) [pvp]; Burial Shawl (18681, -2.10 DPS) [dungeon]; Argent Shoulders (19059, -3.09 DPS, sim-verified) [crafted] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 21.0 spell_power points | yes | Amplifying Cloak (18350, -0.83 DPS) [dungeon]; Hide of the Wild (18510, -1.25 DPS) [crafted]; Arcanoweave Cloak (272411, -2.26 DPS, sim-verified) [vendor] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.3 spell_power points | yes | Field Marshal's Dreadweave Robe (231582, -2.84 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -5.06 DPS) [pvp]; Robe of Everlasting Night (18385, -5.49 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.0 spell_power points | yes | Sublime Wristguards (18497, -2.64 DPS) [dungeon]; Runecloth Cuffs (254123, -2.91 DPS) [crafted]; Deathmist Bracers (226907, -3.88 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.3 spell_power points | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -1.73 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.7 spell_power points | yes | Stormpike Cloth Girdle (19094, -3.39 DPS) [rep]; Satyrmane Sash (17755, -4.50 DPS) [dungeon]; Belt of the Archmage (18405, -5.41 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 36.1 spell_power points | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -1.36 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.35 DPS) [pvp]; Omnicast Boots (11822, -2.22 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.39 DPS) [quest]; Maiden's Circle (13001, -1.39 DPS) [world_drop]; Naglering (11669, -9.46 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.83 DPS) [quest]; Maiden's Circle (13001, -0.83 DPS) [world_drop]; Naglering (11669, -9.33 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.94 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.18 DPS, sim-verified) [quest]; Serenity Field (272439, -4.16 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.98 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (390.9 DPS) | yes | Bonecreeper Stylus (13938, -0.94 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.56 DPS) [world]; Torch of Light (279246, -9.59 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 33.8. Weights run: 2.4s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.074, intellect=-0.024 ± 0.005, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.481 per %), hit=0.140 ± 0.001 per rating point (10 rating = 1%, 1.397 per %), spell_haste=not significant (0.213 ± 0.077), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.672 ± 0.074, fire_power=0.328 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.83 DPS) | yes | Red Winter Hat (21524, -2.80 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.69 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.14 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.55 DPS) | yes | Feyscale Cloak (6632, -0.14 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.14 DPS) [rep]; Black Whelp Cloak (7283, -0.21 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.69 DPS) | yes | Green Woolen Vest (2582, -0.14 DPS) [crafted]; Bloody Apron (6226, -0.14 DPS) [dungeon]; Gray Woolen Robe (2585, -1.27 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) (or Owlbeard Bracers (16981)) | Breaking the Breaker [quest] | 1.0 spell_power points (0.14 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.97 DPS) | yes | Gnoll Casting Gloves (892, -0.23 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.41 DPS) [quest]; Pristine Gloves (253913, -0.41 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (33.8 DPS) | yes | Novice Ardent's Sash (253887, -0.28 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.76 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.24 DPS) | yes | Filigreed Pristine Leggings (253937, -0.41 DPS) [crafted]; Rumpled Kilt (274741, -0.55 DPS) [vendor]; Silk-threaded Trousers (1929, -0.68 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (0.97 DPS) | yes | Red Woolen Boots (4313, -0.41 DPS) [crafted]; Feather Padded Treads (285345, -0.52 DPS, sim-verified) [world]; Pristine Boots (253889, -0.55 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.69 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.41 DPS) | yes | Ring of the Shadow (1462, -0.68 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 8 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 164.4 spell_power points (22.72 DPS) | yes | Skycaller (12984, -1.79 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.43 DPS) [dungeon]; Sizzle Stick (8071, -4.38 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 58.8. Weights run: 2.4s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.080, intellect=0.114 ± 0.006, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.406 per %), hit=0.121 ± 0.001 per rating point (10 rating = 1%, 1.206 per %), spell_haste=0.535 ± 0.088, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.793 ± 0.080, fire_power=0.210 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.37 DPS) | yes | Silk Headband (7050, -0.48 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.65 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.65 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.7 spell_power points (1.66 DPS) | yes | Crystal Starfire Medallion (5003, -1.56 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.69 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.0 spell_power points (2.16 DPS) | yes | Chestnut Mantle (17695, -0.52 DPS, sim-verified) [quest]; Invoker's Mantle (215365, -0.53 DPS) [crafted]; Death Speaker Mantle (6685, -0.60 DPS) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.08 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Battle Healer's Cloak (19529, -0.22 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.80 DPS) | yes | Green Silk Armor (7065, -0.44 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.02 DPS) [dungeon]; High Robe of the Adjudicator (3461, -1.03 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.94 DPS) | yes | Windsong Bangles (263336, -1.72 DPS) [quest]; Glowing Magical Bracelets (13106, -1.74 DPS) [world_drop]; Owlbeard Bracers (16981, -1.75 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.51 DPS) | yes | Jutebraid Gloves (10654, -0.09 DPS) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Truefaith Gloves (7049, -0.36 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.3 spell_power points (2.44 DPS) | yes | Warsong Sash (16975, -0.07 DPS) [quest]; Belt of Arugal (6392, -0.43 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.72 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.58 DPS) | yes | Abomination Skin Leggings (23173, -0.56 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.90 DPS) [crafted]; Silk-threaded Trousers (1929, -1.08 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.8 spell_power points (1.68 DPS) | yes | Acidic Walkers (9454, -0.41 DPS) [dungeon]; Boots of the Enchanter (4325, -0.60 DPS) [crafted]; Spidersilk Boots (4320, -1.57 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.51 DPS) | yes | Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Sludge-Stained Band (286535, -0.86 DPS) [world]; Sacred Band (6669, -1.08 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.29 DPS) | yes | Electrocutioner Lagnut (9447, -0.65 DPS) [dungeon]; Sacred Band (6669, -0.86 DPS) [quest]; Sludge-Stained Band (286535, -2.22 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.94 DPS) | yes | Glimmering Staff (249392, -1.45 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.69 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.69 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.7 spell_power points (1.66 DPS) | yes | Orb of Souls (249395, -0.79 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -0.89 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -1.13 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 157.2 spell_power points (33.86 DPS) | yes | Starfaller (13063, -0.67 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.78 DPS) [crafted]; Gravestone Scepter (7001, -4.86 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 130.4. Weights run: 1.9s. Verify run: 1.0s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.084, intellect=0.460 ± 0.016, crit=0.052 ± 0.002 per rating point (14 rating = 1%, 0.733 per %), hit=0.210 ± 0.002 per rating point (10 rating = 1%, 2.101 per %), spell_haste=1.806 ± 0.106, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.748 ± 0.084, fire_power=0.251 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.47 DPS) | yes | Living Cowl (5608, -1.70 DPS) [world]; Holy Shroud (2721, -2.13 DPS) [world_drop]; Augural Shroud (2620, -2.67 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (2.08 DPS) | yes | Triune Amulet (7722, -1.39 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.39 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.36 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.1 spell_power points (2.80 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Inquisitor's Shawl (19507, -0.03 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.1 spell_power points (2.80 DPS) | yes | Guardian Cloak (5965, -1.03 DPS) [crafted]; Icy Cloak (4327, -1.31 DPS) [crafted]; Long Silken Cloak (4326, -2.14 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.8 spell_power points (5.27 DPS) | yes | Elemental Raiment (9434, -0.80 DPS) [world_drop]; Robe of Power (7054, -1.11 DPS) [crafted]; Dreamweave Vest (10021, -1.18 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.91 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.43 DPS) [quest]; Radiant Silver Bracers (4545, -1.65 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (4.22 DPS) | yes | Black Mageweave Gloves (10003, -1.03 DPS) [crafted]; Gilded Handwraps (254021, -1.83 DPS) [crafted]; Red Mageweave Gloves (10018, -1.90 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.8 spell_power points (3.37 DPS) | yes | Star Belt (4329, -0.60 DPS) [crafted]; Gilded Cord (254037, -0.89 DPS) [crafted]; Deathmage Sash (10771, -1.52 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.5 spell_power points (4.15 DPS) | yes | Abomination Skin Leggings (23173, -1.46 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.60 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.78 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.11 DPS) | yes | Gilded Slippers (254001, -2.95 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.23 DPS) [crafted]; Acidic Walkers (9454, -3.26 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (2.71 DPS) | yes | Reedknot Ring (9622, -1.23 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.44 DPS) [vendor]; Black Widow Band (6199, -2.03 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.91 DPS) | yes | Sea Giant's Toe Ring (274746, -0.64 DPS) [vendor]; Reedknot Ring (9622, -1.03 DPS, sim-verified) [quest]; Black Widow Band (6199, -1.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (130.4 DPS) | yes | Scorn's Focal Dagger (23168, -2.34 DPS) [dungeon]; Staff of Dar'Orahil (15106, -2.73 DPS) [quest]; Gut Ripper (2164, -7.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 187.9 spell_power points (39.97 DPS) | yes | Umbral Wand (5216, -4.30 DPS) [dungeon]; Earthen Rod (9381, -4.38 DPS) [dungeon]; Twisted Nether Wand (249144, -4.69 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 177.2. Weights run: 1.8s. Verify run: 1.2s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.064, intellect=0.250 ± 0.012, crit=0.059 ± 0.002 per rating point (14 rating = 1%, 0.820 per %), hit=0.220 ± 0.002 per rating point (10 rating = 1%, 2.203 per %), spell_haste=not significant (-0.062 ± 0.127), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.813 ± 0.064, fire_power=0.187 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.49 DPS) | yes | Dreamweave Circlet (10041, -0.97 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.67 DPS) [crafted]; Red Mageweave Headband (10033, -2.04 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 spell_power points (2.36 DPS) | yes | Mindburst Medallion (11196, -0.28 DPS) [quest]; Horizon Choker (13085, -1.39 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.66 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.5 spell_power points (4.86 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.46 DPS) [crafted]; Bloodmage Mantle (7684, -1.73 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (4.30 DPS) | yes | Deep Woodlands Cloak (19121, -0.35 DPS) [quest]; Mantle of Lady Falther'ess (23178, -1.18 DPS) [dungeon]; Runecloth Cloak (13860, -1.25 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.0 spell_power points (6.66 DPS) | yes | Robe of the Magi (1716, -0.14 DPS) [world_drop]; Elemental Raiment (9434, -0.83 DPS) [world_drop]; Dreamweave Vest (10021, -1.04 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (5.27 DPS) | yes | Black Mageweave Gloves (10003, -1.11 DPS) [crafted]; Runecloth Gloves (13863, -1.32 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.41 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (4.58 DPS) | yes | Ban'thok Sash (11662, -0.42 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.68 DPS, sim-verified) [rep]; Ghostweave Cord (254073, -0.69 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (7.08 DPS) | yes | Red Mageweave Pants (10009, -2.36 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.69 DPS) [vendor]; Wizardweave Leggings (14132, -4.26 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.66 DPS) | yes | Gilded Sandals (254107, -0.66 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.12 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.20 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.61 DPS) | yes | Philanthropist's Ring (281635, -0.42 DPS) [quest]; Cyclopean Band (11824, -0.62 DPS) [dungeon]; Runed Ring (862, -1.67 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.33 DPS) | yes | Philanthropist's Ring (281635, -0.14 DPS) [quest]; Cyclopean Band (11824, -0.35 DPS) [dungeon]; Runed Ring (862, -1.39 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.72 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -2.90 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.12 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.98 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.05 DPS) [dungeon]; Blade of Eternal Darkness (17780, -4.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (177.2 DPS) | yes | Wand of Allistarj (13065, -2.97 DPS) [world_drop]; Pyric Caduceus (11748, -3.18 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.42 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532000000000000-2332113101220001350-0040000000000000)

Set DPS (verified): 390.1. Weights run: 8.3s. Verify run: 1.2s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± 0.341), intellect=not significant (-0.809 ± 0.025), crit=not significant (-0.183 ± 0.003) per rating point (14 rating = 1%, -2.558 per %), hit=not significant (-0.758 ± 0.004) per rating point (10 rating = 1%, -7.581 per %), spell_haste=not significant (-1.404 ± 0.442), spell_penetration=not significant (-0.000 ± 0.000), shadow_power=not significant (1.292 ± 0.341), fire_power=not significant (-0.292 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.0 spell_power points | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.58 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.02 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.04 DPS) [quest]; Beads of Ogre Mojo (22149, -1.66 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 31.6 spell_power points | yes | Warlord's Dreadweave Mantle (231592, -0.64 DPS) [pvp]; Burial Shawl (18681, -2.10 DPS) [dungeon]; Argent Shoulders (19059, -3.20 DPS, sim-verified) [crafted] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 21.0 spell_power points | yes | Arcanoweave Cloak (272411, -0.22 DPS) [vendor]; Amplifying Cloak (18350, -0.83 DPS) [dungeon]; Hide of the Wild (18510, -1.25 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.3 spell_power points | yes | Robe of Everlasting Night (18385, -2.31 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -2.84 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -5.06 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.0 spell_power points | yes | Sublime Wristguards (18497, -2.64 DPS) [dungeon]; Runecloth Cuffs (254123, -2.91 DPS) [crafted]; Deathmist Bracers (226907, -3.88 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 28.3 spell_power points | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -1.73 DPS) [pvp] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.7 spell_power points | yes | Frostwolf Cloth Belt (19090, -3.39 DPS) [rep]; Satyrmane Sash (17755, -4.50 DPS) [dungeon]; Belt of the Archmage (18405, -4.78 DPS, sim-verified) [crafted] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.90 DPS) [rep]; Sentinel's Silk Leggings (237815, -3.95 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.28 DPS) [dungeon]; Dragonrider Boots (18102, -0.55 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.39 DPS) [quest]; Maiden's Circle (13001, -1.39 DPS) [world_drop]; Naglering (11669, -9.27 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.83 DPS) [quest]; Maiden's Circle (13001, -0.83 DPS) [world_drop]; Naglering (11669, -8.18 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+14.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.94 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.45 DPS, sim-verified) [quest]; Serenity Field (272439, -4.16 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.98 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+11.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.94 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.56 DPS) [world]; Torch of Light (279246, -11.73 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

