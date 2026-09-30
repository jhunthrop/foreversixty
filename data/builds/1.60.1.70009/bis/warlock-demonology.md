# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 1.4s. Verify run: 1.2s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.554, intellect=0.515 ± 0.091, crit=0.168 ± 0.008 per rating point (14 rating = 1%, 2.347 per %), hit=0.671 ± 0.007 per rating point (10 rating = 1%, 6.711 per %), spell_haste=not significant (2.349 ± 0.675), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.554), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.22 DPS) | yes | Shadow Goggles (4373, -2.69 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.6 spell_power points (0.36 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.15 DPS) | yes | Feyscale Cloak (6632, -0.04 DPS) [dungeon]; Caretaker's Cape (20428, -0.04 DPS) [rep]; Pearl-clasped Cloak (5542, -0.37 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.6 spell_power points (0.28 DPS) | yes | Green Woolen Robe (6243, -0.11 DPS) [crafted]; Bloody Apron (6226, -0.13 DPS) [dungeon]; Gray Woolen Robe (2585, -1.09 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 2.6 spell_power points (0.09 DPS) | yes | Bright Bracers (3647, +0.00 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -0.06 DPS) [world_drop]; Repurposed Hair Band (281256, -0.06 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.26 DPS) | yes | Pristine Gloves (253913, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.14 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.1 spell_power points (0.22 DPS) | yes | Keller's Girdle (2911, -0.07 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.09 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (36.0 DPS) | yes | Silk-threaded Trousers (1929, -0.08 DPS) [dungeon]; Rumpled Kilt (274741, -0.15 DPS) [vendor]; Abomination Skin Leggings (23173, -0.75 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.1 spell_power points (0.33 DPS) | yes | Pristine Boots (253889, -0.17 DPS) [crafted]; Red Woolen Boots (4313, -0.19 DPS) [crafted]; Feather Padded Treads (285345, -0.50 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.0 spell_power points (0.22 DPS) | yes | Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sludge-Stained Band (286535, -0.11 DPS) [world]; Volcanic Rock Ring (12053, -0.17 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.18 DPS) | yes | Sludge-Stained Band (286535, -0.07 DPS) [world]; Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.09 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 5.2 spell_power points (0.19 DPS) | yes | Channeler's Staff (4437, +0.00 DPS, sim-verified) [world]; Lesser Staff of the Spire (1300, -0.08 DPS) [world]; Staff of Westfall (2042, -0.09 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 607.9 spell_power points (22.42 DPS) | yes | Skycaller (12984, -1.29 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.13 DPS) [dungeon]; Deepblaze (279896, -3.89 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 55.7. Weights run: 1.4s. Verify run: 1.2s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.094, intellect=0.086 ± 0.009, crit=0.034 ± 0.002 per rating point (14 rating = 1%, 0.477 per %), hit=0.143 ± 0.001 per rating point (10 rating = 1%, 1.431 per %), spell_haste=0.573 ± 0.112, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.094, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.13 DPS) | yes | Silk Headband (7050, -0.54 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.58 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.58 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 spell_power points (1.46 DPS) | yes | Crystal Starfire Medallion (5003, -1.39 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.39 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.51 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (1.89 DPS) | yes | Moonlit Amice (11884, -0.54 DPS) [quest]; Death Speaker Mantle (6685, -0.55 DPS) [dungeon]; Invoker's Mantle (215365, -0.57 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.97 DPS) | yes | Caretaker's Cape (19533, +0.00 DPS, sim-verified) [rep]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Prelacy Cape (7004, -0.19 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.52 DPS) | yes | Green Silk Armor (7065, +0.00 DPS, sim-verified) [crafted]; Robes of Arcana (5770, -0.97 DPS) [crafted]; Death Speaker Robes (6682, -0.98 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.74 DPS) | yes | Glowing Magical Bracelets (13106, -1.61 DPS) [world_drop]; Nightsky Wristbands (6407, -1.64 DPS) [world_drop]; Windsong Bangles (263336, -1.88 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.36 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.19 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.3 spell_power points (2.18 DPS) | yes | Belt of Arugal (6392, -0.54 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.63 DPS) [dungeon]; Invoker's Cord (215366, -0.74 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.33 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Silk-threaded Trousers (1929, -0.97 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.6 spell_power points (1.47 DPS) | yes | Nimbus Boots (6998, -0.31 DPS) [quest]; Acidic Walkers (9454, -0.37 DPS) [dungeon]; Spidersilk Boots (4320, -1.58 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.36 DPS) | yes | Minor Channeling Ring (1449, -0.35 DPS) [quest]; Lorekeeper's Ring (20431, -0.39 DPS) [rep]; Electrocutioner Lagnut (9447, -0.78 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.16 DPS) | yes | Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon]; Sludge-Stained Band (286535, -0.58 DPS) [world]; Minor Channeling Ring (1449, -1.58 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.74 DPS) | yes | Glimmering Staff (249392, -1.42 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.58 DPS) [world_drop]; Channeler's Staff (4437, -1.61 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.5 spell_power points (1.46 DPS) | yes | Eye of Paleth (2943, -0.68 DPS) [quest]; Orb of Souls (249395, -0.68 DPS) [crafted]; Dwarven Tome (279898, -1.00 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 174.4 spell_power points (33.79 DPS) | yes | Starfaller (13063, -0.17 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.82 DPS) [crafted]; Gravestone Scepter (7001, -4.79 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 112.1. Weights run: 1.3s. Verify run: 1.0s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.149, intellect=0.339 ± 0.009, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.456 per %), hit=0.151 ± 0.002 per rating point (10 rating = 1%, 1.508 per %), spell_haste=not significant (0.160 ± 0.159), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.149, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.82 DPS) | yes | Living Cowl (5608, -1.84 DPS) [world]; Augural Shroud (2620, -2.09 DPS, sim-verified) [world]; Holy Shroud (2721, -2.30 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.07 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.08 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.53 DPS) [dungeon]; Triune Amulet (7722, -1.53 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (2.77 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.38 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.1 spell_power points (2.77 DPS) | yes | Guardian Cloak (5965, -1.00 DPS) [crafted]; Icy Cloak (4327, -1.16 DPS) [crafted]; Long Silken Cloak (4326, -1.53 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (5.52 DPS) | yes | Dreamweave Vest (10021, -0.43 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.70 DPS) [world_drop]; Robe of Power (7054, -1.37 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.07 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.46 DPS) [quest]; Earthen Silk Cuffs (254019, -1.15 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.4 spell_power points (4.44 DPS) | yes | Black Mageweave Gloves (10003, -1.13 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.14 DPS) [crafted]; Gilded Handwraps (254021, -2.06 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.4 spell_power points (3.53 DPS) | yes | Star Belt (4329, -0.24 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -0.75 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.77 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.1 spell_power points (4.15 DPS) | yes | Gaze Dreamer Pants (6903, -1.39 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.44 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.46 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.51 DPS) | yes | Gilded Slippers (254001, -2.39 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.59 DPS) [crafted]; Acidic Walkers (9454, -3.74 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (2.76 DPS) | yes | Ring of Forlorn Spirits (2043, -0.93 DPS) [quest]; Reedknot Ring (9622, -1.16 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.39 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.07 DPS) | yes | Reedknot Ring (9622, -0.46 DPS) [quest]; Lorekeeper's Ring (19525, -0.46 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.59 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (112.1 DPS) | yes | Scorn's Focal Dagger (23168, -2.53 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.39 DPS) [quest]; Gut Ripper (2164, -6.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 173.6 spell_power points (39.85 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.26 DPS) [dungeon]; Twisted Nether Wand (249144, -4.47 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 156.5. Weights run: 1.3s. Verify run: 1.2s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.208, intellect=0.255 ± 0.013, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.535 per %), hit=0.149 ± 0.002 per rating point (10 rating = 1%, 1.493 per %), spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -0.97 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -2.27 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 17.5 spell_power points (4.91 DPS) | yes | Mindburst Medallion (11196, -2.80 DPS) [quest]; Horizon Choker (13085, -3.91 DPS) [world_drop]; Scorn's Icy Choker (23169, -4.13 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.6 spell_power points (4.93 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.48 DPS) [crafted]; Bloodmage Mantle (7684, -1.76 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (4.35 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.65 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.26 DPS) [crafted]; Nightfall Drape (12465, -1.83 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.1 spell_power points (6.75 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.87 DPS) [world_drop]; Dreamweave Vest (10021, -1.07 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, -0.06 DPS) [crafted]; Bloodband Bracers (11469, -0.48 DPS) [quest]; Spidertank Oilrag (9448, -0.95 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (5.33 DPS) | yes | Black Mageweave Gloves (10003, -1.13 DPS) [crafted]; Runecloth Gloves (13863, -1.32 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.10 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (4.64 DPS) | yes | Highlander's Cloth Girdle (20098, -0.16 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.45 DPS) [dungeon]; Ghostweave Cord (254073, -0.71 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (7.16 DPS) | yes | Red Mageweave Pants (10009, -2.38 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.79 DPS) [vendor]; Wizardweave Leggings (14132, -3.57 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -1.79 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.14 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.42 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.41 DPS) [quest]; Cyclopean Band (11824, -0.62 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.36 DPS) | yes | Philanthropist's Ring (281635, -0.24 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.34 DPS) [dungeon]; Lorekeeper's Ring (19524, -0.84 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, -1.52 DPS, sim-verified) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | 0.0 spell_power points (0.00 DPS) | yes | Moonshadow Stave (22458, -0.55 DPS) [quest]; Arbiter's Blade (11784, -3.01 DPS) [dungeon]; Blade of Eternal Darkness (17780, -6.17 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (156.5 DPS) | yes | Woestave (20082, -0.08 DPS) [quest]; Pyric Caduceus (11748, -2.02 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.98 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 342.4. Weights run: 1.3s. Verify run: 1.1s. 1027 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.972, intellect=not significant (0.230 ± 0.095), crit=0.233 ± 0.014 per rating point (14 rating = 1%, 3.263 per %), hit=0.719 ± 0.010 per rating point (10 rating = 1%, 7.192 per %), spell_haste=not significant (-0.152 ± 1.139), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.849 ± 0.972), fire_power=0.152 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 82.6 spell_power points (8.44 DPS) | yes | Field Marshal's Coronal (231584, -4.61 DPS) [pvp]; Crimson Felt Hat (18727, -5.19 DPS) [dungeon] |
| neck | Orb of the Darkmoon (19426) (or Chains of the Lich (23125)) | 1200 Tickets - Orb of the Darkmoon [quest] | 22.0 spell_power points (2.25 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Amulet of the Dawn (22657, -0.41 DPS) [quest]; Arcane Crystal Pendant (20037, -0.47 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 53.4 spell_power points (5.46 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.51 DPS) [pvp]; Rugged Mantle of the Timbermaw (227808, -6.61 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.0 spell_power points (2.56 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -0.72 DPS) [dungeon]; Hide of the Wild (18510, -0.89 DPS) [crafted] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 75.4 spell_power points (7.71 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.87 DPS) [pvp]; Robe of the Void (14153, -8.25 DPS, sim-verified) [crafted] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 36.2 spell_power points (3.70 DPS) | yes | Dryad's Wrist Bindings (19596, -1.51 DPS) [rep]; Heretic Wristguards (240152, -1.53 DPS) [vendor] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 47.2 spell_power points (4.82 DPS) | yes | Marshal's Dreadweave Gloves (231586, -1.61 DPS) [pvp]; Sandworm Skin Gloves (20716, -1.94 DPS) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 56.7 spell_power points (5.79 DPS) | yes | Heretic Waistguard (240151, -2.45 DPS) [vendor]; Belt of the Archmage (18405, -3.04 DPS) [crafted]; Knowledge of the Timbermaw (228190, -4.90 DPS, sim-verified) [vendor] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 72.6 spell_power points (7.42 DPS) | yes | Marshal's Dreadweave Leggings (231587, -3.19 DPS) [pvp]; Sentinel's Silk Leggings (237815, -3.25 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 46.4 spell_power points (4.74 DPS) | yes | Marshal's Dreadweave Boots (231585, -1.78 DPS) [pvp]; Earthen Silk Slippers (254013, -2.29 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (342.4 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.23 DPS) [vendor]; Naglering (11669, -11.16 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (342.4 DPS) | yes | Songstone of Ironforge (12543, -0.50 DPS) [quest]; Maiden's Circle (13001, -0.50 DPS) [world_drop]; Naglering (11669, -7.78 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (342.4 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (342.4 DPS) | yes | Weakness Analyzer (272438, -0.72 DPS) [vendor]; Serenity Field (272439, -1.53 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.46 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (342.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.31 DPS) [dungeon]; Electrified Dagger (19100, -14.00 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 757.4 spell_power points (77.37 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS, sim-verified) [dungeon]; Wand of Biting Cold (19108, -13.37 DPS) [quest]; Bonecreeper Stylus (13938, -13.53 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1027, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 34.9. Weights run: 1.4s. Verify run: 1.2s. 142 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.554, intellect=0.515 ± 0.091, crit=0.168 ± 0.008 per rating point (14 rating = 1%, 2.347 per %), hit=0.671 ± 0.007 per rating point (10 rating = 1%, 6.711 per %), spell_haste=not significant (2.349 ± 0.675), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.554), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.22 DPS) | yes | Shadow Goggles (4373, -2.80 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.6 spell_power points (0.36 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.21 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.15 DPS) | yes | Feyscale Cloak (6632, -0.04 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.04 DPS) [rep]; Pearl-clasped Cloak (5542, -0.51 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.6 spell_power points (0.28 DPS) | yes | Green Woolen Robe (6243, -0.11 DPS) [crafted]; Bloody Apron (6226, -0.13 DPS) [dungeon]; Gray Woolen Robe (2585, -1.53 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.1 spell_power points (0.11 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Bright Bracers (3647, -0.04 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.26 DPS) | yes | Pristine Gloves (253913, -0.05 DPS) [crafted]; Apothecary Gloves (10919, -0.11 DPS) [quest]; Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.1 spell_power points (0.22 DPS) | yes | Keller's Girdle (2911, -0.07 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.09 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.88 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk-threaded Trousers (1929, -0.08 DPS) [dungeon]; Rumpled Kilt (274741, -0.15 DPS) [vendor]; Abomination Skin Leggings (23173, -0.53 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.1 spell_power points (0.33 DPS) | yes | Pristine Boots (253889, -0.17 DPS) [crafted]; Red Woolen Boots (4313, -0.19 DPS) [crafted]; Feather Padded Treads (285345, -0.53 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.18 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.34 DPS, sim-verified) [dungeon] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Loop of Sacrifice (281673, -0.02 DPS) [quest]; Volcanic Rock Ring (12053, -0.05 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.60 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 5.2 spell_power points (0.19 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.04 DPS) [world]; Lesser Staff of the Spire (1300, -0.08 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 607.9 spell_power points (22.42 DPS) | yes | Skycaller (12984, -1.69 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.13 DPS) [dungeon]; Sizzle Stick (8071, -4.59 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 142, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 54.1. Weights run: 1.4s. Verify run: 1.1s. 237 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.094, intellect=0.086 ± 0.009, crit=0.034 ± 0.002 per rating point (14 rating = 1%, 0.477 per %), hit=0.143 ± 0.001 per rating point (10 rating = 1%, 1.431 per %), spell_haste=0.573 ± 0.112, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.094, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.13 DPS) | yes | Silk Headband (7050, -0.42 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.58 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.58 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 spell_power points (1.46 DPS) | yes | Darkspear Warding Pendant (272075, -1.38 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.39 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.39 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (1.89 DPS) | yes | Invoker's Mantle (215365, -0.45 DPS) [crafted]; Death Speaker Mantle (6685, -0.55 DPS) [dungeon]; Chestnut Mantle (17695, -1.28 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.97 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Battle Healer's Cloak (19529, -0.19 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.52 DPS) | yes | Green Silk Armor (7065, +0.00 DPS, sim-verified) [crafted]; High Robe of the Adjudicator (3461, -0.94 DPS) [quest]; Robes of Arcana (5770, -0.97 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.74 DPS) | yes | Owlbeard Bracers (16981, -1.48 DPS, sim-verified) [quest]; Windsong Bangles (263336, -1.55 DPS) [quest]; Glowing Magical Bracelets (13106, -1.61 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.36 DPS) | yes | Jutebraid Gloves (10654, +0.00 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.19 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.3 spell_power points (2.18 DPS) | yes | Warsong Sash (16975, -0.28 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.39 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.63 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.33 DPS) | yes | Abomination Skin Leggings (23173, -0.08 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Silk-threaded Trousers (1929, -0.97 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.6 spell_power points (1.47 DPS) | yes | Acidic Walkers (9454, -0.37 DPS) [dungeon]; Boots of the Enchanter (4325, -0.50 DPS) [crafted]; Spidersilk Boots (4320, -1.75 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.36 DPS) | yes | Advisor's Ring (20426, -0.39 DPS) [rep]; Electrocutioner Lagnut (9447, -0.78 DPS) [dungeon]; Sludge-Stained Band (286535, -0.78 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.16 DPS) | yes | Sludge-Stained Band (286535, -0.58 DPS) [world]; Sacred Band (6669, -0.78 DPS) [quest]; Electrocutioner Lagnut (9447, -1.63 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.74 DPS) | yes | Glimmering Staff (249392, -1.25 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.58 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.58 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.5 spell_power points (1.46 DPS) | yes | Orb of Souls (249395, -0.41 DPS, sim-verified) [crafted]; Alliance Outrunner Healing Rod (285348, -0.68 DPS) [world]; Tome of the Darkspear Prophecy (272090, -1.00 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 174.4 spell_power points (33.79 DPS) | yes | Starfaller (13063, -0.59 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.82 DPS) [crafted]; Gravestone Scepter (7001, -4.79 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 237, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 110.9. Weights run: 1.3s. Verify run: 0.9s. 319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.149, intellect=0.339 ± 0.009, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.456 per %), hit=0.151 ± 0.002 per rating point (10 rating = 1%, 1.508 per %), spell_haste=not significant (0.160 ± 0.159), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.149, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.82 DPS) | yes | Living Cowl (5608, -1.84 DPS) [world]; Holy Shroud (2721, -2.30 DPS) [world_drop]; Augural Shroud (2620, -2.82 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.07 DPS) | yes | Necklace of Calisea (1714, -1.53 DPS) [dungeon]; Triune Amulet (7722, -1.53 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.11 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (2.77 DPS) | yes | Inquisitor's Shawl (19507, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.38 DPS) [quest]; Green Silken Shoulders (7057, -0.78 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.1 spell_power points (2.77 DPS) | yes | Guardian Cloak (5965, -1.00 DPS) [crafted]; Icy Cloak (4327, -1.16 DPS) [crafted]; Long Silken Cloak (4326, -2.44 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (5.52 DPS) | yes | Elemental Raiment (9434, -0.70 DPS) [world_drop]; Robe of Power (7054, -1.37 DPS) [crafted]; Dreamweave Vest (10021, -1.58 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.07 DPS) | yes | Radiant Silver Bracers (4545, -0.52 DPS) [quest]; Condor Bracers (15864, -1.06 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.15 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.4 spell_power points (4.44 DPS) | yes | Red Mageweave Gloves (10018, -1.14 DPS) [crafted]; Black Mageweave Gloves (10003, -1.50 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.06 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.4 spell_power points (3.53 DPS) | yes | Deathmage Sash (10771, -0.75 DPS) [dungeon]; Defiler's Cloth Girdle (20164, -0.77 DPS) [rep]; Star Belt (4329, -1.01 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.1 spell_power points (4.15 DPS) | yes | Gaze Dreamer Pants (6903, -1.39 DPS) [dungeon]; Abomination Skin Leggings (23173, -1.46 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.83 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.51 DPS) | yes | Gilded Slippers (254001, -2.95 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.59 DPS) [crafted]; Acidic Walkers (9454, -3.74 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (2.76 DPS) | yes | Reedknot Ring (9622, -1.16 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.39 DPS) [vendor]; Electrocutioner Lagnut (9447, -2.07 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.07 DPS) | yes | Advisor's Ring (19521, -0.46 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.69 DPS) [vendor]; Reedknot Ring (9622, -1.06 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (110.9 DPS) | yes | Scorn's Focal Dagger (23168, -2.53 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.39 DPS) [quest]; Gut Ripper (2164, -7.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 173.6 spell_power points (39.85 DPS) | yes | Umbral Wand (5216, -0.52 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.26 DPS) [dungeon]; Twisted Nether Wand (249144, -4.47 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 152.6. Weights run: 1.3s. Verify run: 1.1s. 408 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.208, intellect=0.255 ± 0.013, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.535 per %), hit=0.149 ± 0.002 per rating point (10 rating = 1%, 1.493 per %), spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -0.97 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.91 DPS, sim-verified) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 17.5 spell_power points (4.91 DPS) | yes | Mindburst Medallion (11196, -2.80 DPS) [quest]; Horizon Choker (13085, -3.91 DPS) [world_drop]; Scorn's Icy Choker (23169, -4.51 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.6 spell_power points (4.93 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.48 DPS) [crafted]; Bloodmage Mantle (7684, -1.76 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (4.35 DPS) | yes | Deep Woodlands Cloak (19121, -0.56 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.19 DPS) [dungeon]; Runecloth Cloak (13860, -1.26 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 24.1 spell_power points (6.75 DPS) | yes | Robe of the Magi (1716, -0.46 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.87 DPS) [world_drop]; Dreamweave Vest (10021, -1.07 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -0.48 DPS) [quest]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (5.33 DPS) | yes | Black Mageweave Gloves (10003, -1.13 DPS) [crafted]; Runecloth Gloves (13863, -1.32 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.59 DPS, sim-verified) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.5 spell_power points (4.64 DPS) | yes | Ban'thok Sash (11662, -0.45 DPS) [dungeon]; Ghostweave Cord (254073, -0.71 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.15 DPS, sim-verified) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 25.5 spell_power points (7.16 DPS) | yes | Red Mageweave Pants (10009, -2.38 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.79 DPS) [vendor]; Wizardweave Leggings (14132, -4.41 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -0.77 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.14 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.42 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.41 DPS) [quest]; Cyclopean Band (11824, -0.62 DPS) [dungeon]; Runed Ring (862, -1.68 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.36 DPS) | yes | Cyclopean Band (11824, -0.34 DPS) [dungeon]; Advisor's Ring (19520, -0.84 DPS) [rep]; Philanthropist's Ring (281635, -1.05 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (152.6 DPS) | yes | Frozen Heart of the Mountain (249469, -2.99 DPS) [crafted]; Rune of the Guard Captain (19120, -3.07 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (152.6 DPS) | yes | Fire Ruby (20036, -1.15 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -1.30 DPS) [crafted]; Rune of the Guard Captain (19120, -1.39 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (152.6 DPS) | yes | Moonshadow Stave (22458, -0.55 DPS) [quest]; Arbiter's Blade (11784, -3.01 DPS) [dungeon]; Blade of Eternal Darkness (17780, -5.24 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 187.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Woestave (20082, -1.18 DPS) [quest]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; waist: Satyrmane Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Pyric Caduceus

No-known-source sample (15 of 408, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 340.0. Weights run: 1.3s. Verify run: 1.1s. 1018 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.972, intellect=not significant (0.230 ± 0.095), crit=0.233 ± 0.014 per rating point (14 rating = 1%, 3.263 per %), hit=0.719 ± 0.010 per rating point (10 rating = 1%, 7.192 per %), spell_haste=not significant (-0.152 ± 1.139), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.849 ± 0.972), fire_power=0.152 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 82.6 spell_power points (8.44 DPS) | yes | Warlord's Dreadweave Hood (231590, -4.61 DPS) [pvp]; Crimson Felt Hat (18727, -5.19 DPS) [dungeon] |
| neck | Orb of the Darkmoon (19426) (or Chains of the Lich (23125)) | 1200 Tickets - Orb of the Darkmoon [quest] | 22.0 spell_power points (2.25 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Amulet of the Dawn (22657, -0.41 DPS) [quest]; Arcane Crystal Pendant (20037, -0.47 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 53.4 spell_power points (5.46 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.51 DPS) [pvp]; Rugged Mantle of the Timbermaw (227808, -8.54 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.0 spell_power points (2.56 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -0.72 DPS) [dungeon]; Hide of the Wild (18510, -0.89 DPS) [crafted] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 75.4 spell_power points (7.71 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.87 DPS) [pvp]; Robe of the Void (14153, -8.47 DPS, sim-verified) [crafted] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 36.2 spell_power points (3.70 DPS) | yes | Dryad's Wrist Bindings (19596, -1.51 DPS) [rep]; Heretic Wristguards (240152, -1.53 DPS) [vendor] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 47.2 spell_power points (4.82 DPS) | yes | General's Dreadweave Gloves (231589, -1.61 DPS) [pvp]; Sandworm Skin Gloves (20716, -1.94 DPS) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 56.7 spell_power points (5.79 DPS) | yes | Heretic Waistguard (240151, -2.45 DPS) [vendor]; Belt of the Archmage (18405, -3.04 DPS) [crafted]; Knowledge of the Timbermaw (228190, -6.07 DPS, sim-verified) [vendor] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 72.6 spell_power points (7.42 DPS) | yes | General's Dreadweave Pants (231588, -3.19 DPS) [pvp]; Sentinel's Silk Leggings (237815, -3.25 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 46.4 spell_power points (4.74 DPS) | yes | General's Dreadweave Boots (231593, -1.78 DPS) [pvp]; Earthen Silk Slippers (254013, -2.29 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (340.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.23 DPS) [vendor]; Naglering (11669, -11.46 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (340.0 DPS) | yes | Eye of Orgrimmar (12545, -0.50 DPS) [quest]; Maiden's Circle (13001, -0.50 DPS) [world_drop]; Naglering (11669, -7.76 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (340.0 DPS) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -8.25 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (340.0 DPS) | yes | Royal Seal of Eldre'Thalas (18467, -0.61 DPS) [quest]; Weakness Analyzer (272438, -0.72 DPS) [vendor]; Serenity Field (272439, -3.65 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (340.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.31 DPS) [dungeon]; Glacial Blade (19099, -17.48 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 757.4 spell_power points (77.37 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.09 DPS, sim-verified) [dungeon]; Wand of Biting Cold (19108, -13.37 DPS) [quest]; Bonecreeper Stylus (13938, -13.53 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1018, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

