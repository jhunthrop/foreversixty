# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.7. Weights run: 2.1s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.086, intellect=-0.074 ± 0.005, crit=0.039 ± 0.001 per rating point (14 rating = 1%, 0.551 per %), hit=0.138 ± 0.001 per rating point (10 rating = 1%, 1.385 per %), spell_haste=not significant (-0.010 ± 0.086), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.685 ± 0.086

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.80 DPS) | yes | Red Winter Hat (21524, -3.12 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.67 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.13 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.54 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Caretaker's Cape (20428, -0.13 DPS) [rep]; Black Whelp Cloak (7283, -0.38 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.67 DPS) | yes | Green Woolen Vest (2582, -0.13 DPS) [crafted]; Bloody Apron (6226, -0.13 DPS) [dungeon]; Gray Woolen Robe (2585, -1.37 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.13 DPS) | yes | Ivycloth Bracelets (9793, -0.38 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.94 DPS) | yes | Pristine Gloves (253913, -0.40 DPS) [crafted]; Gnoll Casting Gloves (892, -0.43 DPS, sim-verified) [world]; Heavy Woolen Gloves (4310, -0.67 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (38.7 DPS) | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.82 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.20 DPS) | yes | Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Rumpled Kilt (274741, -0.54 DPS) [vendor]; Silk-threaded Trousers (1929, -0.57 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (0.94 DPS) | yes | Red Woolen Boots (4313, -0.40 DPS) [crafted]; Pristine Boots (253889, -0.54 DPS) [crafted]; Feather Padded Treads (285345, -0.58 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.67 DPS) | yes | Sludge-Stained Band (286535, -0.27 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.67 DPS) | yes | Sludge-Stained Band (286535, -0.47 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.21 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.67 DPS) | yes | Bouquet of Red Roses (22206, -1.27 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 169.8 spell_power points (22.71 DPS) | yes | Skycaller (12984, -1.67 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 63.5. Weights run: 2.1s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.057, intellect=0.105 ± 0.004, crit=0.025 ± 0.001 per rating point (14 rating = 1%, 0.348 per %), hit=0.086 ± 0.001 per rating point (10 rating = 1%, 0.859 per %), spell_haste=not significant (0.205 ± 0.060), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.839 ± 0.057

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.99 DPS) | yes | Silk Headband (7050, -0.59 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.82 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.82 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (2.08 DPS) | yes | Darkspear Warding Pendant (272075, -1.96 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.96 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.96 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (2.70 DPS) | yes | Invoker's Mantle (215365, -0.58 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.76 DPS) [dungeon]; Moonlit Amice (11884, -0.80 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.36 DPS) | yes | Heavy Woolen Cloak (4311, -0.24 DPS, sim-verified) [crafted]; Prelacy Cape (7004, -0.27 DPS) [quest]; Caretaker's Cape (19533, -0.27 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.54 DPS) | yes | Green Silk Armor (7065, -1.04 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.32 DPS) [dungeon]; Robes of Arcana (5770, -1.36 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.45 DPS) | yes | Glowing Magical Bracelets (13106, -2.22 DPS) [world_drop]; Nightsky Wristbands (6407, -2.28 DPS) [world_drop]; Windsong Bangles (263336, -2.38 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.90 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.27 DPS) [world]; Truefaith Gloves (7049, -0.46 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.3 spell_power points (3.08 DPS) | yes | Belt of Arugal (6392, -0.54 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.90 DPS) [dungeon]; Invoker's Cord (215366, -1.03 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.26 DPS) | yes | Abomination Skin Leggings (23173, -0.84 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.16 DPS) [crafted]; Silk-threaded Trousers (1929, -1.36 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.7 spell_power points (2.10 DPS) | yes | Nimbus Boots (6998, -0.47 DPS) [quest]; Acidic Walkers (9454, -0.52 DPS) [dungeon]; Spidersilk Boots (4320, -1.72 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.90 DPS) | yes | Minor Channeling Ring (1449, -0.49 DPS) [quest]; Electrocutioner Lagnut (9447, -1.09 DPS) [dungeon]; Sludge-Stained Band (286535, -1.09 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.63 DPS) | yes | Electrocutioner Lagnut (9447, -0.82 DPS) [dungeon]; Sludge-Stained Band (286535, -0.82 DPS) [world]; Minor Channeling Ring (1449, -1.87 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.45 DPS) | yes | Glimmering Staff (249392, -1.86 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -2.16 DPS) [world_drop]; Channeler's Staff (4437, -2.22 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.6 spell_power points (2.08 DPS) | yes | Dwarven Tome (279898, -0.59 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.99 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.99 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 125.1 spell_power points (34.03 DPS) | yes | Starfaller (13063, -0.86 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.67 DPS) [crafted]; Gravestone Scepter (7001, -5.03 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 118.2. Weights run: 1.7s. Verify run: 1.0s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.131, intellect=0.378 ± 0.009, crit=0.043 ± 0.001 per rating point (14 rating = 1%, 0.598 per %), hit=0.149 ± 0.001 per rating point (10 rating = 1%, 1.488 per %), spell_haste=-0.819 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.800 ± 0.131

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.49 DPS) | yes | Living Cowl (5608, -1.71 DPS) [world]; Holy Shroud (2721, -2.14 DPS) [world_drop]; Augural Shroud (2620, -2.59 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 spell_power points (1.98 DPS) | yes | Triune Amulet (7722, -1.42 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.42 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.17 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 spell_power points (2.65 DPS) | yes | Inquisitor's Shawl (19507, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.35 DPS) [quest]; Green Silken Shoulders (7057, -0.49 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.4 spell_power points (2.65 DPS) | yes | Guardian Cloak (5965, -0.97 DPS) [crafted]; Icy Cloak (4327, -1.16 DPS) [crafted]; Long Silken Cloak (4326, -1.84 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.3 spell_power points (5.19 DPS) | yes | Elemental Raiment (9434, -0.70 DPS) [world_drop]; Robe of Power (7054, -1.23 DPS) [crafted]; Dreamweave Vest (10021, -1.24 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.93 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.43 DPS) [quest]; Earthen Silk Cuffs (254019, -1.07 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.5 spell_power points (4.17 DPS) | yes | Black Mageweave Gloves (10003, -0.71 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.01 DPS) [crafted]; Gilded Handwraps (254021, -1.90 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.5 spell_power points (3.32 DPS) | yes | Star Belt (4329, -0.54 DPS) [crafted]; Deathmage Sash (10771, -0.61 DPS) [dungeon]; Gilded Cord (254037, -0.96 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.5 spell_power points (3.96 DPS) | yes | Abomination Skin Leggings (23173, -1.39 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.40 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.10 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.13 DPS) | yes | Gilded Slippers (254001, -2.91 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.31 DPS) [crafted]; Acidic Walkers (9454, -3.42 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.3 spell_power points (2.62 DPS) | yes | Ring of Forlorn Spirits (2043, -0.91 DPS) [quest]; Reedknot Ring (9622, -1.13 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.34 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.93 DPS) | yes | Ring of Forlorn Spirits (2043, -0.21 DPS) [quest]; Reedknot Ring (9622, -0.43 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.64 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (118.2 DPS) | yes | Scorn's Focal Dagger (23168, -2.35 DPS) [dungeon]; Windweaver Staff (7757, -3.07 DPS) [dungeon]; Gut Ripper (2164, -8.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 186.4 spell_power points (39.87 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.28 DPS) [dungeon]; Twisted Nether Wand (249144, -4.58 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 205.8. Weights run: 2.0s. Verify run: 1.4s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.316, intellect=0.534 ± 0.026, crit=0.074 ± 0.003 per rating point (14 rating = 1%, 1.041 per %), hit=0.474 ± 0.005 per rating point (10 rating = 1%, 4.738 per %), spell_haste=-1.557 ± 0.318, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.728 ± 0.316)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.09 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -0.78 DPS) [vendor]; Red Mageweave Headband (10033, -2.06 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 spell_power points (1.44 DPS) | yes | Mindburst Medallion (11196, -0.14 DPS) [quest]; Horizon Choker (13085, -0.39 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.69 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 22.6 spell_power points (3.19 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.07 DPS) [crafted]; Black Mageweave Shoulders (10027, -1.10 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.2 spell_power points (2.43 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.55 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.55 DPS) [crafted]; Big Voodoo Cloak (8216, -1.04 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 29.7 spell_power points (4.19 DPS) | yes | Robe of the Magi (1716, -0.63 DPS) [world_drop]; Runecloth Tunic (13857, -0.96 DPS) [crafted]; Dreamweave Vest (10021, -0.97 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 10.7 spell_power points (1.51 DPS) | yes | Arcane Runed Bracers (4744, -0.24 DPS) [quest]; Spidertank Oilrag (9448, -0.24 DPS) [dungeon]; Bloodband Bracers (11469, -2.76 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 spell_power points (2.84 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -0.33 DPS) [vendor]; Runecloth Gloves (13863, -0.47 DPS) [crafted]; Raider Handwraps (272098, -1.37 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 19.3 spell_power points (2.73 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -0.45 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.45 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 28.3 spell_power points (4.00 DPS) | yes | Knight's Dreadweave Leggings (220888, -1.25 DPS) [vendor]; Wizardweave Leggings (14132, -1.32 DPS) [crafted]; Red Mageweave Pants (10009, -5.72 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.39 DPS) | yes | Gilded Sandals (254107, -1.16 DPS) [crafted]; Black Mageweave Boots (10026, -1.31 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.76 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (1.86 DPS) | yes | Cyclopean Band (11824, -0.07 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.17 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.83 DPS) | yes | Lorekeeper's Ring (19523, -0.14 DPS) [rep]; Cyclopean Band (11824, -0.75 DPS, sim-verified) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.02 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -0.64 DPS) [world_drop]; Soul Harvester (20536, -0.95 DPS) [quest]; Blade of Eternal Darkness (17780, -7.26 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.29 DPS) [world_drop]; Flaming Incinerator (9483, -3.49 DPS) [dungeon]; Pyric Caduceus (11748, -6.96 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Spellshock Leggings; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 420.8. Weights run: 5.8s. Verify run: 1.4s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.213, intellect=0.061 ± 0.014, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.857 per %), hit=0.281 ± 0.003 per rating point (10 rating = 1%, 2.808 per %), spell_haste=not significant (0.385 ± 0.290), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.911 ± 0.213

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.5 spell_power points (13.64 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -3.37 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -6.55 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.85 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.78 DPS) [quest]; Kezan's Taint (19604, -3.36 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.8 spell_power points (12.88 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.22 DPS) [pvp]; Argent Shoulders (19059, -2.16 DPS, sim-verified) [crafted]; Burial Shawl (18681, -3.49 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.2 spell_power points (9.06 DPS) | yes | Arcanoweave Cloak (272411, -0.42 DPS) [vendor]; Amplifying Cloak (18350, -1.00 DPS) [dungeon]; Hide of the Wild (18510, -2.52 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (20.83 DPS) | yes | Robe of Everlasting Night (18385, -2.37 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -5.86 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.10 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.5 spell_power points (10.06 DPS) | yes | Sublime Wristguards (18497, -4.42 DPS) [dungeon]; Runecloth Cuffs (254123, -4.87 DPS) [crafted]; Arcane Runed Bracers (4744, -6.04 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.95 DPS) [quest]; Sandworm Skin Gloves (20716, -4.12 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 29.2 spell_power points (13.04 DPS) | yes | Stormpike Cloth Girdle (19094, -4.72 DPS) [rep]; Belt of the Archmage (18405, -5.31 DPS, sim-verified) [crafted]; Oddly Magical Belt (18475, -5.88 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 34.5 spell_power points (15.43 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -2.38 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.74 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.90 DPS) [crafted]; Omnicast Boots (11822, -1.46 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.90 DPS) [quest]; Maiden's Circle (13001, -1.90 DPS) [world_drop]; Naglering (11669, -10.64 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.34 DPS) [quest]; Maiden's Circle (13001, -1.34 DPS) [world_drop]; Naglering (11669, -9.84 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.99 DPS, sim-verified) [quest]; Weakness Analyzer (272438, -3.13 DPS) [vendor]; Serenity Field (272439, -6.71 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.64 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -20.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+18.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.11 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.96 DPS) [world]; Torch of Light (279246, -18.43 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.4. Weights run: 2.1s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.086, intellect=-0.074 ± 0.005, crit=0.039 ± 0.001 per rating point (14 rating = 1%, 0.551 per %), hit=0.138 ± 0.001 per rating point (10 rating = 1%, 1.385 per %), spell_haste=not significant (-0.010 ± 0.086), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.685 ± 0.086

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.80 DPS) | yes | Red Winter Hat (21524, -2.75 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.67 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.13 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.54 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.13 DPS) [rep]; Black Whelp Cloak (7283, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.67 DPS) | yes | Green Woolen Vest (2582, -0.13 DPS) [crafted]; Bloody Apron (6226, -0.13 DPS) [dungeon]; Gray Woolen Robe (2585, -1.22 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) (or Owlbeard Bracers (16981)) | Breaking the Breaker [quest] | 1.0 spell_power points (0.13 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.94 DPS) | yes | Gnoll Casting Gloves (892, -0.35 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.40 DPS) [quest]; Pristine Gloves (253913, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (36.4 DPS) | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.01 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.20 DPS) | yes | Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Silk-threaded Trousers (1929, -0.53 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.54 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (0.94 DPS) | yes | Red Woolen Boots (4313, -0.40 DPS) [crafted]; Feather Padded Treads (285345, -0.42 DPS, sim-verified) [world]; Pristine Boots (253889, -0.54 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.67 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.40 DPS) | yes | Ring of the Shadow (1462, -0.72 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 8 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 169.8 spell_power points (22.71 DPS) | yes | Skycaller (12984, -1.67 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Sizzle Stick (8071, -4.39 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 62.5. Weights run: 2.1s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.057, intellect=0.105 ± 0.004, crit=0.025 ± 0.001 per rating point (14 rating = 1%, 0.348 per %), hit=0.086 ± 0.001 per rating point (10 rating = 1%, 0.859 per %), spell_haste=not significant (0.205 ± 0.060), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.839 ± 0.057

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.99 DPS) | yes | Silk Headband (7050, -0.63 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.82 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.82 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (2.08 DPS) | yes | Darkspear Warding Pendant (272075, -1.88 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.96 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.96 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (2.70 DPS) | yes | Invoker's Mantle (215365, -0.66 DPS) [crafted]; Death Speaker Mantle (6685, -0.76 DPS) [dungeon]; Chestnut Mantle (17695, -1.15 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.36 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.27 DPS) [crafted]; Battle Healer's Cloak (19529, -0.27 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.54 DPS) | yes | Green Silk Armor (7065, -0.72 DPS, sim-verified) [crafted]; High Robe of the Adjudicator (3461, -1.30 DPS) [quest]; Death Speaker Robes (6682, -1.32 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.45 DPS) | yes | Windsong Bangles (263336, -2.18 DPS) [quest]; Glowing Magical Bracelets (13106, -2.22 DPS) [world_drop]; Owlbeard Bracers (16981, -2.30 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.90 DPS) | yes | Jutebraid Gloves (10654, -0.13 DPS) [quest]; Gnoll Casting Gloves (892, -0.27 DPS) [world]; Truefaith Gloves (7049, -0.46 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.3 spell_power points (3.08 DPS) | yes | Warsong Sash (16975, -0.33 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.54 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.90 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.26 DPS) | yes | Abomination Skin Leggings (23173, -0.62 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.16 DPS) [crafted]; Silk-threaded Trousers (1929, -1.36 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.7 spell_power points (2.10 DPS) | yes | Acidic Walkers (9454, -0.52 DPS) [dungeon]; Boots of the Enchanter (4325, -0.74 DPS) [crafted]; Spidersilk Boots (4320, -2.09 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.90 DPS) | yes | Electrocutioner Lagnut (9447, -1.09 DPS) [dungeon]; Sludge-Stained Band (286535, -1.09 DPS) [world]; Sacred Band (6669, -1.36 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.63 DPS) | yes | Electrocutioner Lagnut (9447, -0.82 DPS) [dungeon]; Sacred Band (6669, -1.09 DPS) [quest]; Sludge-Stained Band (286535, -2.35 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.45 DPS) | yes | Glimmering Staff (249392, -1.79 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -2.16 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -2.16 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.6 spell_power points (2.08 DPS) | yes | Orb of Souls (249395, -0.99 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.42 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.48 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 125.1 spell_power points (34.03 DPS) | yes | Starfaller (13063, -0.88 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.67 DPS) [crafted]; Gravestone Scepter (7001, -5.03 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 117.2. Weights run: 1.7s. Verify run: 1.0s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.131, intellect=0.378 ± 0.009, crit=0.043 ± 0.001 per rating point (14 rating = 1%, 0.598 per %), hit=0.149 ± 0.001 per rating point (10 rating = 1%, 1.488 per %), spell_haste=-0.819 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.800 ± 0.131

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.49 DPS) | yes | Living Cowl (5608, -1.71 DPS) [world]; Holy Shroud (2721, -2.14 DPS) [world_drop]; Augural Shroud (2620, -3.00 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 spell_power points (1.98 DPS) | yes | Triune Amulet (7722, -1.42 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.42 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.65 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 spell_power points (2.65 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS) [crafted]; Inquisitor's Shawl (19507, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.35 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.4 spell_power points (2.65 DPS) | yes | Guardian Cloak (5965, -0.97 DPS) [crafted]; Icy Cloak (4327, -1.16 DPS) [crafted]; Long Silken Cloak (4326, -2.34 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.3 spell_power points (5.19 DPS) | yes | Elemental Raiment (9434, -0.70 DPS) [world_drop]; Dreamweave Vest (10021, -0.98 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.23 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.93 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.43 DPS) [quest]; Radiant Silver Bracers (4545, -2.04 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.5 spell_power points (4.17 DPS) | yes | Red Mageweave Gloves (10018, -1.01 DPS) [crafted]; Black Mageweave Gloves (10003, -1.15 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.90 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.5 spell_power points (3.32 DPS) | yes | Deathmage Sash (10771, -0.61 DPS) [dungeon]; Star Belt (4329, -0.68 DPS, sim-verified) [crafted]; Gilded Cord (254037, -0.96 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.5 spell_power points (3.96 DPS) | yes | Abomination Skin Leggings (23173, -1.39 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.40 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.10 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.13 DPS) | yes | Spidersilk Boots (4320, -3.31 DPS) [crafted]; Acidic Walkers (9454, -3.42 DPS) [dungeon]; Gilded Slippers (254001, -3.52 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.3 spell_power points (2.62 DPS) | yes | Reedknot Ring (9622, -1.13 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.34 DPS) [vendor]; Sludge-Stained Band (286535, -1.98 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.93 DPS) | yes | Sea Giant's Toe Ring (274746, -0.64 DPS) [vendor]; Reedknot Ring (9622, -0.66 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -1.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (117.2 DPS) | yes | Scorn's Focal Dagger (23168, -2.35 DPS) [dungeon]; Windweaver Staff (7757, -3.07 DPS) [dungeon]; Gut Ripper (2164, -8.14 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 186.4 spell_power points (39.87 DPS) | yes | Umbral Wand (5216, -4.20 DPS) [dungeon]; Earthen Rod (9381, -4.28 DPS) [dungeon]; Twisted Nether Wand (249144, -4.58 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 208.0. Weights run: 2.0s. Verify run: 1.4s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.316, intellect=0.534 ± 0.026, crit=0.074 ± 0.003 per rating point (14 rating = 1%, 1.041 per %), hit=0.474 ± 0.005 per rating point (10 rating = 1%, 4.738 per %), spell_haste=-1.557 ± 0.318, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.728 ± 0.316)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.09 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -0.78 DPS) [vendor]; Red Mageweave Headband (10033, -2.66 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 spell_power points (1.44 DPS) | yes | Mindburst Medallion (11196, -0.14 DPS) [quest]; Horizon Choker (13085, -0.39 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.69 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Shoulders (10029, -0.84 DPS) [crafted]; Black Mageweave Shoulders (10027, -0.87 DPS) [crafted]; Rotgrip Mantle (17732, -2.55 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.2 spell_power points (2.43 DPS) | yes | Deep Woodlands Cloak (19121, -0.06 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.48 DPS) [dungeon]; Runecloth Cloak (13860, -0.55 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 29.7 spell_power points (4.19 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Runecloth Tunic (13857, -0.96 DPS) [crafted]; Dreamweave Vest (10021, -0.97 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 10.7 spell_power points (1.51 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -2.66 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 spell_power points (2.84 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -0.33 DPS) [vendor]; Runecloth Gloves (13863, -0.47 DPS) [crafted]; Raider Handwraps (272098, -0.81 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Dawnspire Cord (12466, -0.31 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.31 DPS) [rep]; Satyrmane Sash (17755, -2.36 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 28.3 spell_power points (4.00 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.25 DPS) [vendor]; Wizardweave Leggings (14132, -1.32 DPS) [crafted]; Red Mageweave Pants (10009, -5.04 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.39 DPS) | yes | Gilded Sandals (254107, -1.16 DPS) [crafted]; Black Mageweave Boots (10026, -1.31 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.39 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (1.86 DPS) | yes | Cyclopean Band (11824, -0.07 DPS) [dungeon]; Advisor's Ring (19519, -0.17 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.83 DPS) | yes | Cyclopean Band (11824, -0.04 DPS) [dungeon]; Advisor's Ring (19519, -0.14 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.04 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -1.23 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.13 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -0.64 DPS) [world_drop]; Soul Harvester (20536, -0.95 DPS) [quest]; Blade of Eternal Darkness (17780, -6.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.29 DPS) [world_drop]; Flaming Incinerator (9483, -3.49 DPS) [dungeon]; Pyric Caduceus (11748, -6.96 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 410.4. Weights run: 5.8s. Verify run: 1.4s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.213, intellect=0.061 ± 0.014, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.857 per %), hit=0.281 ± 0.003 per rating point (10 rating = 1%, 2.808 per %), spell_haste=not significant (0.385 ± 0.290), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.911 ± 0.213

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.5 spell_power points (13.64 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -3.37 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -4.69 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (9.85 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.78 DPS) [quest]; Kezan's Taint (19604, -3.36 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.8 spell_power points (12.88 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.22 DPS) [pvp]; Argent Shoulders (19059, -1.69 DPS) [crafted]; Burial Shawl (18681, -3.49 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.2 spell_power points (9.06 DPS) | yes | Amplifying Cloak (18350, -1.00 DPS) [dungeon]; Hide of the Wild (18510, -2.52 DPS) [crafted]; Arcanoweave Cloak (272411, -2.83 DPS, sim-verified) [vendor] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.5 spell_power points (20.83 DPS) | yes | Robe of Everlasting Night (18385, -4.01 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -5.86 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.10 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.5 spell_power points (10.06 DPS) | yes | Sublime Wristguards (18497, -4.42 DPS) [dungeon]; Runecloth Cuffs (254123, -4.87 DPS) [crafted]; Spidertank Oilrag (9448, -6.04 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.3 spell_power points (12.22 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.37 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 29.2 spell_power points (13.04 DPS) | yes | Belt of the Archmage (18405, -4.64 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -4.72 DPS) [rep]; Oddly Magical Belt (18475, -5.88 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 34.5 spell_power points (15.43 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.38 DPS) [rep]; Sentinel's Silk Leggings (237815, -4.61 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.74 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.90 DPS) [crafted]; Omnicast Boots (11822, -1.46 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.90 DPS) [quest]; Maiden's Circle (13001, -1.90 DPS) [world_drop]; Naglering (11669, -8.27 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.34 DPS) [quest]; Maiden's Circle (13001, -1.34 DPS) [world_drop]; Naglering (11669, -8.87 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -1.68 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.13 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.83 DPS, sim-verified) [quest]; Serenity Field (272439, -6.71 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.64 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -18.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (410.4 DPS) | yes | Bonecreeper Stylus (13938, -1.11 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.96 DPS) [world]; Torch of Light (279246, -17.78 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

