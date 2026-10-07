# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 2.4s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.101, intellect=0.342 ± 0.017, crit=0.050 ± 0.002 per rating point (14 rating = 1%, 0.695 per %), hit=0.189 ± 0.001 per rating point (10 rating = 1%, 1.895 per %), spell_haste=not significant (0.091 ± 0.109), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.101, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.69 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.79 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.39 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.37 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.7 spell_power points (0.66 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.27 DPS) [crafted]; Gray Woolen Robe (2585, -1.09 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.7 spell_power points (0.17 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.07 DPS) [quest]; Repurposed Hair Band (281256, -0.10 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.42 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.53 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (36.0 DPS) | yes | Silk-threaded Trousers (1929, -0.10 DPS) [dungeon]; Rumpled Kilt (274741, -0.30 DPS) [vendor]; Abomination Skin Leggings (23173, -0.75 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.43 DPS) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Feather Padded Treads (285345, -0.50 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.56 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.46 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.49 DPS) | yes | Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Sludge-Stained Band (286535, -0.39 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.4 spell_power points (0.34 DPS) | yes | Channeler's Staff (4437, -0.07 DPS) [world]; Lesser Staff of the Spire (1300, -0.13 DPS) [world]; Staff of Westfall (2042, -0.17 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 230.1 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.29 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 56.8. Weights run: 2.4s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.072, intellect=0.051 ± 0.005, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.423 per %), hit=0.116 ± 0.001 per rating point (10 rating = 1%, 1.160 per %), spell_haste=0.976 ± 0.088, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.072, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.12 DPS) | yes | Silk Headband (7050, -0.49 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.58 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.58 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 spell_power points (1.41 DPS) | yes | Crystal Starfire Medallion (5003, -1.37 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.37 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.62 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (1.83 DPS) | yes | Moonlit Amice (11884, -0.48 DPS) [quest]; Death Speaker Mantle (6685, -0.56 DPS) [dungeon]; Invoker's Mantle (215365, -0.83 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.97 DPS) | yes | Prelacy Cape (7004, -0.19 DPS) [quest]; Caretaker's Cape (19533, -0.19 DPS) [rep]; Heavy Woolen Cloak (4311, -0.24 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.51 DPS) | yes | Green Silk Armor (7065, -0.64 DPS) [crafted]; Robes of Arcana (5770, -0.97 DPS) [crafted]; Death Speaker Robes (6682, -1.05 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.74 DPS) | yes | Glowing Magical Bracelets (13106, -1.66 DPS) [world_drop]; Nightsky Wristbands (6407, -1.68 DPS) [world_drop]; Windsong Bangles (263336, -1.96 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.35 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world]; Truefaith Gloves (7049, -0.36 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.2 spell_power points (2.15 DPS) | yes | Belt of Arugal (6392, -0.50 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.61 DPS) [dungeon]; Invoker's Cord (215366, -0.75 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.32 DPS) | yes | Abomination Skin Leggings (23173, -0.28 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.90 DPS) [crafted]; Silk-threaded Trousers (1929, -0.97 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.4 spell_power points (1.42 DPS) | yes | Nimbus Boots (6998, -0.26 DPS) [quest]; Acidic Walkers (9454, -0.38 DPS) [dungeon]; Spidersilk Boots (4320, -1.74 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.35 DPS) | yes | Minor Channeling Ring (1449, -0.37 DPS) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Sludge-Stained Band (286535, -0.77 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.16 DPS) | yes | Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon]; Sludge-Stained Band (286535, -0.58 DPS) [world]; Minor Channeling Ring (1449, -1.33 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.74 DPS) | yes | Glimmering Staff (249392, -1.48 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.64 DPS) [world_drop]; Channeler's Staff (4437, -1.66 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.3 spell_power points (1.41 DPS) | yes | Eye of Paleth (2943, -0.64 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.64 DPS) [world]; Dwarven Tome (279898, -0.85 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 174.9 spell_power points (33.79 DPS) | yes | Starfaller (13063, -0.35 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.82 DPS) [crafted]; Gravestone Scepter (7001, -4.79 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 124.6. Weights run: 1.8s. Verify run: 1.0s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.110, intellect=0.330 ± 0.008, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.479 per %), hit=0.143 ± 0.001 per rating point (10 rating = 1%, 1.432 per %), spell_haste=not significant (0.161 ± 0.129), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.110, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.76 DPS) | yes | Augural Shroud (2620, -2.14 DPS, sim-verified) [world]; Living Cowl (5608, -2.19 DPS) [world]; Holy Shroud (2721, -2.74 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.46 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.22 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.83 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.83 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (3.28 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Inquisitor's Shawl (19507, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.0 spell_power points (3.28 DPS) | yes | Guardian Cloak (5965, -1.18 DPS) [crafted]; Icy Cloak (4327, -1.36 DPS) [crafted]; Long Silken Cloak (4326, -1.46 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (6.58 DPS) | yes | Elemental Raiment (9434, -0.82 DPS) [world_drop]; Dreamweave Vest (10021, -0.83 DPS) [crafted]; Robe of Power (7054, -1.65 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.47 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.55 DPS) [quest]; Earthen Silk Cuffs (254019, -1.37 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (5.30 DPS) | yes | Black Mageweave Gloves (10003, -0.78 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.38 DPS) [crafted]; Gilded Handwraps (254021, -2.47 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.3 spell_power points (4.20 DPS) | yes | Star Belt (4329, -0.64 DPS) [crafted]; Deathmage Sash (10771, -0.93 DPS) [dungeon]; Gilded Cord (254037, -1.28 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.0 spell_power points (4.92 DPS) | yes | Gaze Dreamer Pants (6903, -1.63 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.72 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.73 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.58 DPS) | yes | Gilded Slippers (254001, -2.55 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.30 DPS) [crafted]; Acidic Walkers (9454, -4.49 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (3.29 DPS) | yes | Ring of Forlorn Spirits (2043, -1.09 DPS) [quest]; Reedknot Ring (9622, -1.37 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.64 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.47 DPS) | yes | Reedknot Ring (9622, -0.55 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.69 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.82 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (124.6 DPS) | yes | Scorn's Focal Dagger (23168, -3.02 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.10 DPS) [quest]; Gut Ripper (2164, -6.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 145.5 spell_power points (39.92 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.28 DPS) [crafted]; Earthen Rod (9381, -4.33 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 168.8. Weights run: 1.7s. Verify run: 1.2s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.171, intellect=0.233 ± 0.015, crit=0.062 ± 0.003 per rating point (14 rating = 1%, 0.872 per %), hit=0.242 ± 0.002 per rating point (10 rating = 1%, 2.425 per %), spell_haste=not significant (0.917 ± 0.261), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.927 ± 0.171, fire_power=0.071 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (5.75 DPS) | yes | Dreamweave Circlet (10041, -0.78 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.28 DPS) [crafted]; Red Mageweave Headband (10033, -1.47 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 spell_power points (1.79 DPS) | yes | Mindburst Medallion (11196, -0.21 DPS) [quest]; Horizon Choker (13085, -1.09 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.29 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.2 spell_power points (3.66 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.09 DPS) [crafted]; Bloodmage Mantle (7684, -1.30 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.4 spell_power points (3.28 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.92 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.97 DPS) [crafted]; Nightfall Drape (12465, -1.36 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.7 spell_power points (5.04 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.57 DPS) [world_drop]; Dreamweave Vest (10021, -0.76 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.92 DPS) | yes | Nethergeld Cuffs (254061, -0.08 DPS) [crafted]; Bloodband Bracers (11469, -0.40 DPS) [quest]; Spidertank Oilrag (9448, -0.60 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (4.03 DPS) | yes | Black Mageweave Gloves (10003, -0.84 DPS) [crafted]; Runecloth Gloves (13863, -1.03 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.75 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.3 spell_power points (3.48 DPS) | yes | Highlander's Cloth Girdle (20098, -0.30 DPS) [rep]; Ban'thok Sash (11662, -0.32 DPS) [dungeon]; Ghostweave Cord (254073, -0.50 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.3 spell_power points (5.40 DPS) | yes | Red Mageweave Pants (10009, -1.82 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.06 DPS) [vendor]; Wizardweave Leggings (14132, -3.33 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.11 DPS) | yes | Gilded Sandals (254107, -1.69 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -2.42 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.44 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.77 DPS) | yes | Philanthropist's Ring (281635, -0.34 DPS) [quest]; Cyclopean Band (11824, -0.50 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.06 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (2.56 DPS) | yes | Cyclopean Band (11824, -0.29 DPS) [dungeon]; Philanthropist's Ring (281635, -0.81 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -0.85 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.79 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.31 DPS) [dungeon]; Scorn's Focal Dagger (23168, -2.34 DPS) [dungeon]; Blade of Eternal Darkness (17780, -5.65 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (168.8 DPS) | yes | Pyric Caduceus (11748, -2.01 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.64 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.61 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 300.4. Weights run: 1.8s. Verify run: 1.2s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.568, intellect=0.665 ± 0.046, crit=0.234 ± 0.009 per rating point (14 rating = 1%, 3.272 per %), hit=0.775 ± 0.007 per rating point (10 rating = 1%, 7.752 per %), spell_haste=-9.800 ± 0.903, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.868 ± 0.568), fire_power=0.133 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -5.46 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) | Stratholme: Balzaphon [dungeon] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.12 DPS) [quest]; Amulet of the Dawn (22657, -3.51 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 40.3 spell_power points (4.74 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.46 DPS) [pvp]; Burial Shawl (18681, -1.13 DPS) [dungeon]; Darkspear Shoulderpads (272103, -1.21 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.1 spell_power points (3.42 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Hide of the Wild (18510, -0.99 DPS) [crafted]; Amplifying Cloak (18350, -1.30 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 52.0 spell_power points (6.12 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -0.47 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -1.61 DPS) [pvp]; Robe of Everlasting Night (18385, -2.83 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 27.3 spell_power points (3.22 DPS) | yes | Sublime Wristguards (18497, -1.02 DPS) [dungeon]; Runecloth Cuffs (254123, -1.14 DPS) [crafted]; Deathmist Bracers (226907, -1.45 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 30.3 spell_power points (3.57 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -0.11 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 47.4 spell_power points (5.58 DPS) | yes | Belt of the Archmage (18405, -1.92 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -2.68 DPS) [rep]; Highlander's Cloth Girdle (20047, -3.08 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.2 spell_power points (5.32 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.69 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -1.01 DPS) [pvp] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 28.6 spell_power points (3.37 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.08 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.78 DPS) [quest]; Maiden's Circle (13001, -0.78 DPS) [world_drop]; Naglering (11669, -5.23 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.35 DPS) [quest]; Maiden's Circle (13001, -0.35 DPS) [world_drop]; Naglering (11669, -6.34 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -4.83 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -0.82 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -1.65 DPS, sim-verified) [quest]; Serenity Field (272439, -1.77 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.19 DPS) [world]; Teebu's Blazing Longsword (1728, -6.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.91 DPS) [dungeon]; Wand of Biting Cold (19108, -1.15 DPS) [quest]; Torch of Light (279246, -3.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Crackling Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 34.9. Weights run: 2.4s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.101, intellect=0.342 ± 0.017, crit=0.050 ± 0.002 per rating point (14 rating = 1%, 0.695 per %), hit=0.189 ± 0.001 per rating point (10 rating = 1%, 1.895 per %), spell_haste=not significant (0.091 ± 0.109), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.101, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.72 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.79 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.39 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.43 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.7 spell_power points (0.66 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.27 DPS) [crafted]; Gray Woolen Robe (2585, -1.36 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest]; Owlbeard Bracers (16981, -0.04 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Pristine Gloves (253913, -0.19 DPS) [crafted]; Gnoll Casting Gloves (892, -0.24 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.29 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.53 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (34.9 DPS) | yes | Silk-threaded Trousers (1929, -0.10 DPS) [dungeon]; Rumpled Kilt (274741, -0.30 DPS) [vendor]; Abomination Skin Leggings (23173, -0.67 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.43 DPS) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Feather Padded Treads (285345, -0.67 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.49 DPS) | yes | Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Loop of Sacrifice (281673, -0.32 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.4 spell_power points (0.34 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.07 DPS) [world]; Lesser Staff of the Spire (1300, -0.13 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 230.1 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Sizzle Stick (8071, -4.46 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 54.9. Weights run: 2.4s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.072, intellect=0.051 ± 0.005, crit=0.030 ± 0.001 per rating point (14 rating = 1%, 0.423 per %), hit=0.116 ± 0.001 per rating point (10 rating = 1%, 1.160 per %), spell_haste=0.976 ± 0.088, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.072, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.12 DPS) | yes | Silk Headband (7050, -0.39 DPS) [crafted]; Embalmed Shroud (7691, -0.58 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.58 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 spell_power points (1.41 DPS) | yes | Crystal Starfire Medallion (5003, -1.37 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.37 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.40 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (1.83 DPS) | yes | Invoker's Mantle (215365, -0.43 DPS) [crafted]; Death Speaker Mantle (6685, -0.56 DPS) [dungeon]; Chestnut Mantle (17695, -0.77 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.97 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Battle Healer's Cloak (19529, -0.19 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.51 DPS) | yes | Green Silk Armor (7065, -0.64 DPS) [crafted]; High Robe of the Adjudicator (3461, -0.95 DPS) [quest]; Robes of Arcana (5770, -0.97 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.74 DPS) | yes | Owlbeard Bracers (16981, -1.47 DPS, sim-verified) [quest]; Windsong Bangles (263336, -1.55 DPS) [quest]; Glowing Magical Bracelets (13106, -1.66 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.35 DPS) | yes | Jutebraid Gloves (10654, -0.14 DPS) [quest]; Gnoll Casting Gloves (892, -0.19 DPS) [world]; Truefaith Gloves (7049, -0.36 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.2 spell_power points (2.15 DPS) | yes | Warsong Sash (16975, -0.03 DPS) [quest]; Belt of Arugal (6392, -0.39 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.61 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.32 DPS) | yes | Abomination Skin Leggings (23173, -0.50 DPS) [dungeon]; Pristine Leggings (253987, -0.90 DPS) [crafted]; Silk-threaded Trousers (1929, -0.97 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.4 spell_power points (1.42 DPS) | yes | Acidic Walkers (9454, -0.38 DPS) [dungeon]; Boots of the Enchanter (4325, -0.46 DPS) [crafted]; Spidersilk Boots (4320, -1.44 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.35 DPS) | yes | Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Sludge-Stained Band (286535, -0.77 DPS) [world]; Sacred Band (6669, -0.97 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.16 DPS) | yes | Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon]; Sacred Band (6669, -0.77 DPS) [quest]; Sludge-Stained Band (286535, -2.01 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.74 DPS) | yes | Glimmering Staff (249392, -1.02 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.64 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.64 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.3 spell_power points (1.41 DPS) | yes | Orb of Souls (249395, -0.64 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -0.90 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -0.99 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 174.9 spell_power points (33.79 DPS) | yes | Starfaller (13063, -0.24 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.82 DPS) [crafted]; Gravestone Scepter (7001, -4.79 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 123.1. Weights run: 1.8s. Verify run: 1.0s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.110, intellect=0.330 ± 0.008, crit=0.034 ± 0.001 per rating point (14 rating = 1%, 0.479 per %), hit=0.143 ± 0.001 per rating point (10 rating = 1%, 1.432 per %), spell_haste=not significant (0.161 ± 0.129), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.110, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.76 DPS) | yes | Living Cowl (5608, -2.19 DPS) [world]; Augural Shroud (2620, -2.41 DPS, sim-verified) [world]; Holy Shroud (2721, -2.74 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.46 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.54 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.83 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.83 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.0 spell_power points (3.28 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Inquisitor's Shawl (19507, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.46 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.0 spell_power points (3.28 DPS) | yes | Guardian Cloak (5965, -1.18 DPS) [crafted]; Icy Cloak (4327, -1.36 DPS) [crafted]; Long Silken Cloak (4326, -2.45 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (6.58 DPS) | yes | Dreamweave Vest (10021, -0.83 DPS) [crafted]; Elemental Raiment (9434, -1.07 DPS, sim-verified) [world_drop]; Robe of Power (7054, -1.65 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.47 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.65 DPS) [quest]; Condor Bracers (15864, -0.93 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (5.30 DPS) | yes | Black Mageweave Gloves (10003, -1.25 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.38 DPS) [crafted]; Gilded Handwraps (254021, -2.47 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.3 spell_power points (4.20 DPS) | yes | Star Belt (4329, -0.90 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -0.93 DPS) [dungeon]; Warsong Sash (16975, -1.18 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.0 spell_power points (4.92 DPS) | yes | Crimson Silk Pantaloons (7062, -0.99 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.63 DPS) [dungeon]; Abomination Skin Leggings (23173, -1.73 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.58 DPS) | yes | Gilded Slippers (254001, -2.44 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.30 DPS) [crafted]; Acidic Walkers (9454, -4.49 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (3.29 DPS) | yes | Reedknot Ring (9622, -1.37 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.64 DPS) [vendor]; Sludge-Stained Band (286535, -2.46 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.47 DPS) | yes | Sea Giant's Toe Ring (274746, -0.82 DPS) [vendor]; Reedknot Ring (9622, -0.93 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -1.65 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (123.1 DPS) | yes | Scorn's Focal Dagger (23168, -3.02 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.10 DPS) [quest]; Gut Ripper (2164, -6.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 145.5 spell_power points (39.92 DPS) | yes | Umbral Wand (5216, -0.50 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -4.28 DPS) [crafted]; Earthen Rod (9381, -4.33 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 167.7. Weights run: 1.7s. Verify run: 1.2s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.171, intellect=0.233 ± 0.015, crit=0.062 ± 0.003 per rating point (14 rating = 1%, 0.872 per %), hit=0.242 ± 0.002 per rating point (10 rating = 1%, 2.425 per %), spell_haste=not significant (0.917 ± 0.261), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.927 ± 0.171, fire_power=0.071 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (5.75 DPS) | yes | Dreamweave Circlet (10041, -0.78 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.28 DPS) [crafted]; Red Mageweave Headband (10033, -1.60 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 spell_power points (1.79 DPS) | yes | Mindburst Medallion (11196, -0.21 DPS) [quest]; Horizon Choker (13085, -1.09 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.29 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.2 spell_power points (3.66 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.09 DPS) [crafted]; Bloodmage Mantle (7684, -1.30 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.4 spell_power points (3.28 DPS) | yes | Deep Woodlands Cloak (19121, -0.28 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.92 DPS) [dungeon]; Runecloth Cloak (13860, -0.97 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.7 spell_power points (5.04 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.57 DPS) [world_drop]; Dreamweave Vest (10021, -0.76 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -1.81 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (4.03 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -0.61 DPS, sim-verified) [vendor]; Black Mageweave Gloves (10003, -0.84 DPS) [crafted]; Runecloth Gloves (13863, -1.03 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.3 spell_power points (3.48 DPS) | yes | Defiler's Cloth Girdle (20166, -0.30 DPS) [rep]; Ban'thok Sash (11662, -0.32 DPS) [dungeon]; Ghostweave Cord (254073, -0.50 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.3 spell_power points (5.40 DPS) | yes | Red Mageweave Pants (10009, -1.82 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.06 DPS) [vendor]; Wizardweave Leggings (14132, -3.32 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.11 DPS) | yes | Gilded Sandals (254107, -0.69 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -2.42 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.44 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.77 DPS) | yes | Philanthropist's Ring (281635, -0.34 DPS) [quest]; Cyclopean Band (11824, -0.50 DPS) [dungeon]; Runed Ring (862, -1.28 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (2.56 DPS) | yes | Philanthropist's Ring (281635, -0.13 DPS) [quest]; Cyclopean Band (11824, -0.29 DPS) [dungeon]; Runed Ring (862, -1.06 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.61 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -2.19 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, -0.10 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.31 DPS) [dungeon]; Scorn's Focal Dagger (23168, -2.34 DPS) [dungeon]; Blade of Eternal Darkness (17780, -4.61 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.64 DPS) [world_drop]; Pyric Caduceus (11748, -2.76 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.61 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 301.9. Weights run: 1.8s. Verify run: 1.2s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.568, intellect=0.665 ± 0.046, crit=0.234 ± 0.009 per rating point (14 rating = 1%, 3.272 per %), hit=0.775 ± 0.007 per rating point (10 rating = 1%, 7.752 per %), spell_haste=-9.800 ± 0.903, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.868 ± 0.568), fire_power=0.133 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | sim-verified (301.9 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -3.68 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.6 spell_power points (2.78 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Orb of the Darkmoon (19426, -0.19 DPS) [quest]; Beads of Ogre Mojo (22149, -0.31 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 40.3 spell_power points (4.74 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.46 DPS) [pvp]; Darkspear Shoulderpads (272103, -1.21 DPS) [vendor]; Burial Shawl (18681, -1.42 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.1 spell_power points (3.42 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Hide of the Wild (18510, -0.99 DPS) [crafted]; Amplifying Cloak (18350, -1.30 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 52.0 spell_power points (6.12 DPS) | yes | Warlord's Dreadweave Robe (231591, -0.47 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -1.61 DPS) [pvp]; Robe of Everlasting Night (18385, -3.66 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.3 spell_power points (3.22 DPS) | yes | Sublime Wristguards (18497, -1.02 DPS) [dungeon]; Runecloth Cuffs (254123, -1.14 DPS) [crafted]; Deathmist Bracers (226907, -1.45 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 30.3 spell_power points (3.57 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Deathmist Wraps (226911, -0.11 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 47.4 spell_power points (5.58 DPS) | yes | Frostwolf Cloth Belt (19090, -2.68 DPS) [rep]; Defiler's Cloth Girdle (20163, -3.08 DPS) [rep]; Belt of the Archmage (18405, -3.60 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.2 spell_power points (5.32 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.69 DPS) [dungeon]; Outrider's Silk Leggings (22747, -4.06 DPS, sim-verified) [rep] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 28.6 spell_power points (3.37 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Walkers (227098, -0.24 DPS) [pvp]; Omnicast Boots (11822, -2.48 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.78 DPS) [quest]; Maiden's Circle (13001, -0.78 DPS) [world_drop]; Naglering (11669, -7.92 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.35 DPS) [quest]; Maiden's Circle (13001, -0.35 DPS) [world_drop]; Naglering (11669, -7.58 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+14.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -0.82 DPS) [vendor]; Serenity Field (272439, -1.77 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.19 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.30 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -14.09 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 657.8 spell_power points (77.47 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.32 DPS) [dungeon]; Bonecreeper Stylus (13938, -13.23 DPS) [dungeon]; Wand of Biting Cold (19108, -13.47 DPS) [quest] |

**New at 60:** head: Crimson Felt Hat; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

