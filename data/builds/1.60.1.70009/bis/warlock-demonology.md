# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 2.3s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.068, intellect=-0.018 ± 0.004, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.438 per %), hit=0.129 ± 0.001 per rating point (10 rating = 1%, 1.285 per %), spell_haste=not significant (0.131 ± 0.070), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.699 ± 0.068, fire_power=0.302 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.90 DPS) | yes | Red Winter Hat (21524, -2.62 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.75 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.15 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.60 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Caretaker's Cape (20428, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.22 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Green Woolen Vest (2582, -0.15 DPS) [crafted]; Bloody Apron (6226, -0.15 DPS) [dungeon]; Gray Woolen Robe (2585, -1.04 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Ivycloth Bracelets (9793, -0.22 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.05 DPS) | yes | Gnoll Casting Gloves (892, -0.26 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.45 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.75 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (36.0 DPS) | yes | Novice Ardent's Sash (253887, -0.30 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.95 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted]; Rumpled Kilt (274741, -0.60 DPS) [vendor]; Silk-threaded Trousers (1929, -0.65 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.05 DPS) | yes | Feather Padded Treads (285345, -0.35 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.45 DPS) [crafted]; Pristine Boots (253889, -0.60 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.75 DPS) | yes | Sludge-Stained Band (286535, -0.30 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.75 DPS) | yes | Sludge-Stained Band (286535, -0.40 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -0.94 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.75 DPS) | yes | Bouquet of Red Roses (22206, -1.04 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 151.6 spell_power points (22.76 DPS) | yes | Skycaller (12984, -1.53 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.47 DPS) [dungeon]; Deepblaze (279896, -4.23 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 58.0. Weights run: 2.4s. Verify run: 1.1s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.084, intellect=0.101 ± 0.006, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.429 per %), hit=0.128 ± 0.001 per rating point (10 rating = 1%, 1.279 per %), spell_haste=0.617 ± 0.090, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.775 ± 0.084, fire_power=0.228 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.17 DPS) | yes | Embalmed Shroud (7691, -0.59 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.59 DPS) [crafted]; Silk Headband (7050, -0.62 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (1.50 DPS) | yes | Crystal Starfire Medallion (5003, -1.42 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.42 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.62 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (1.96 DPS) | yes | Death Speaker Mantle (6685, -0.55 DPS) [dungeon]; Moonlit Amice (11884, -0.58 DPS) [quest]; Invoker's Mantle (215365, -0.59 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.99 DPS) | yes | Heavy Woolen Cloak (4311, -0.20 DPS) [crafted]; Prelacy Cape (7004, -0.20 DPS) [quest]; Caretaker's Cape (19533, -0.20 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.57 DPS) | yes | Green Silk Armor (7065, -0.64 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.97 DPS) [dungeon]; Robes of Arcana (5770, -0.99 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.78 DPS) | yes | Glowing Magical Bracelets (13106, -1.62 DPS) [world_drop]; Nightsky Wristbands (6407, -1.66 DPS) [world_drop]; Windsong Bangles (263336, -2.07 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.38 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.20 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.3 spell_power points (2.23 DPS) | yes | Belt of Arugal (6392, -0.60 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.65 DPS) [dungeon]; Invoker's Cord (215366, -0.75 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.37 DPS) | yes | Abomination Skin Leggings (23173, -0.53 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Silk-threaded Trousers (1929, -0.99 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.7 spell_power points (1.52 DPS) | yes | Nimbus Boots (6998, -0.34 DPS) [quest]; Acidic Walkers (9454, -0.38 DPS) [dungeon]; Spidersilk Boots (4320, -1.60 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.38 DPS) | yes | Minor Channeling Ring (1449, -0.36 DPS) [quest]; Electrocutioner Lagnut (9447, -0.79 DPS) [dungeon]; Sludge-Stained Band (286535, -0.79 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.19 DPS) | yes | Electrocutioner Lagnut (9447, -0.59 DPS) [dungeon]; Sludge-Stained Band (286535, -0.59 DPS) [world]; Minor Channeling Ring (1449, -1.60 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.78 DPS) | yes | Glimmering Staff (249392, -1.42 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.58 DPS) [world_drop]; Channeler's Staff (4437, -1.62 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.6 spell_power points (1.50 DPS) | yes | Dwarven Tome (279898, -0.58 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.71 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.71 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 171.1 spell_power points (33.80 DPS) | yes | Starfaller (13063, -0.62 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.81 DPS) [crafted]; Gravestone Scepter (7001, -4.80 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 129.1. Weights run: 1.8s. Verify run: 1.0s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.059, intellect=0.206 ± 0.009, crit=0.039 ± 0.001 per rating point (14 rating = 1%, 0.539 per %), hit=0.156 ± 0.001 per rating point (10 rating = 1%, 1.557 per %), spell_haste=0.932 ± 0.071, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.806 ± 0.059, fire_power=0.194 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.87 DPS) | yes | Living Cowl (5608, -2.24 DPS) [world]; Augural Shroud (2620, -2.24 DPS, sim-verified) [world]; Holy Shroud (2721, -2.79 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.2 spell_power points (2.30 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.75 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.90 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.90 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.8 spell_power points (3.03 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Inquisitor's Shawl (19507, -0.33 DPS) [dungeon]; Berylline Pads (4197, -0.50 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 10.8 spell_power points (3.03 DPS) | yes | Guardian Cloak (5965, -1.07 DPS) [crafted]; Icy Cloak (4327, -1.08 DPS) [crafted]; Long Silken Cloak (4326, -1.28 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 spell_power points (6.49 DPS) | yes | Elemental Raiment (9434, -0.62 DPS) [world_drop]; Dreamweave Vest (10021, -0.95 DPS) [crafted]; Robe of Power (7054, -1.89 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest]; Earthen Silk Cuffs (254019, -1.40 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.26 DPS) | yes | Black Mageweave Gloves (10003, -0.65 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.61 DPS) [crafted]; Gilded Handwraps (254021, -2.62 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.8 spell_power points (4.14 DPS) | yes | Star Belt (4329, -0.51 DPS) [crafted]; Deathmage Sash (10771, -1.32 DPS) [dungeon]; Gilded Cord (254037, -1.45 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.5 spell_power points (4.60 DPS) | yes | Gaze Dreamer Pants (6903, -0.74 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.62 DPS) [crafted]; Abomination Skin Leggings (23173, -1.63 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.71 DPS) | yes | Gilded Slippers (254001, -2.63 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.52 DPS) [crafted]; Acidic Walkers (9454, -4.85 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.2 spell_power points (3.14 DPS) | yes | Ring of Forlorn Spirits (2043, -0.90 DPS) [quest]; Reedknot Ring (9622, -1.18 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.46 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.52 DPS) | yes | Ring of Forlorn Spirits (2043, -0.28 DPS) [quest]; Reedknot Ring (9622, -0.56 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.07 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.52 DPS) [quest]; Gut Ripper (2164, -6.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (129.1 DPS) | yes | Umbral Wand (5216, -0.01 DPS) [dungeon]; Earthen Rod (9381, -0.09 DPS) [dungeon]; Jaina's Firestarter (13064, -2.42 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 174.2. Weights run: 1.7s. Verify run: 1.1s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.073, intellect=0.323 ± 0.014, crit=0.059 ± 0.002 per rating point (14 rating = 1%, 0.826 per %), hit=0.227 ± 0.002 per rating point (10 rating = 1%, 2.274 per %), spell_haste=not significant (-0.255 ± 0.140), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.802 ± 0.073, fire_power=0.199 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.96 DPS) | yes | Dreamweave Circlet (10041, -0.71 DPS) [crafted]; Red Mageweave Headband (10033, -1.40 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.55 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (2.30 DPS) | yes | Mindburst Medallion (11196, -0.26 DPS) [quest]; Horizon Choker (13085, -1.14 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.47 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 18.8 spell_power points (4.85 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.52 DPS) [crafted]; Bloodmage Mantle (7684, -1.78 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (4.11 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.04 DPS) [dungeon]; Runecloth Cloak (13860, -1.12 DPS) [crafted]; Nightfall Drape (12465, -1.79 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.5 spell_power points (6.56 DPS) | yes | Robe of the Magi (1716, -0.39 DPS) [world_drop]; Elemental Raiment (9434, -1.15 DPS) [world_drop]; Dreamweave Vest (10021, -1.17 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 9.3 spell_power points (2.39 DPS) | yes | Spidertank Oilrag (9448, -0.07 DPS) [dungeon]; Bloodband Bracers (11469, -0.35 DPS) [quest]; Arcane Runed Bracers (4744, -2.10 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (4.97 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -0.98 DPS, sim-verified) [vendor]; Black Mageweave Gloves (10003, -1.11 DPS) [crafted]; Runecloth Gloves (13863, -1.13 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Cloth Girdle (20098, -0.13 DPS) [rep]; Ghostweave Cord (254073, -0.46 DPS) [crafted]; Satyrmane Sash (17755, -1.95 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.2 spell_power points (6.76 DPS) | yes | Red Mageweave Pants (10009, -2.15 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.46 DPS) [vendor]; Wizardweave Leggings (14132, -3.29 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.19 DPS) | yes | Gilded Sandals (254107, -2.60 DPS) [crafted]; Black Mageweave Boots (10026, -2.77 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.79 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.35 DPS) | yes | Philanthropist's Ring (281635, -0.27 DPS) [quest]; Cyclopean Band (11824, -0.45 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.29 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.09 DPS) | yes | Philanthropist's Ring (281635, -0.02 DPS) [quest]; Cyclopean Band (11824, -0.19 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.03 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.24 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.68 DPS) [dungeon]; Glowing Brightwood Staff (812, -2.74 DPS) [world_drop]; Blade of Eternal Darkness (17780, -4.63 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Pyric Caduceus (11748, -2.24 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.87 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.48 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 360.8. Weights run: 1.9s. Verify run: 1.2s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.196, intellect=0.123 ± 0.019, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.925 per %), hit=0.258 ± 0.003 per rating point (10 rating = 1%, 2.584 per %), spell_haste=not significant (0.242 ± 0.308), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.883 ± 0.196, fire_power=0.117 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.0 spell_power points (14.64 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -3.24 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -3.62 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.40 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.55 DPS) [quest]; Kezan's Taint (19604, -3.32 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.8 spell_power points (14.07 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -1.27 DPS) [pvp]; Argent Shoulders (19059, -2.25 DPS) [crafted]; Burial Shawl (18681, -3.69 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.5 spell_power points (9.68 DPS) | yes | Arcanoweave Cloak (272411, -0.44 DPS) [vendor]; Amplifying Cloak (18350, -1.18 DPS) [dungeon]; Hide of the Wild (18510, -2.49 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (22.26 DPS) | yes | Robe of Everlasting Night (18385, -2.99 DPS, sim-verified) [dungeon]; Field Marshal's Dreadweave Robe (231582, -5.75 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.29 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.0 spell_power points (10.86 DPS) | yes | Sublime Wristguards (18497, -4.61 DPS) [dungeon]; Runecloth Cuffs (254123, -5.08 DPS) [crafted]; Arcane Runed Bracers (4744, -6.61 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (13.05 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.65 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.3 spell_power points (14.31 DPS) | yes | Belt of the Archmage (18405, -1.81 DPS, sim-verified) [crafted]; Stormpike Cloth Girdle (19094, -5.23 DPS) [rep]; Oddly Magical Belt (18475, -6.75 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -2.20 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.34 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.95 DPS) [crafted]; Omnicast Boots (11822, -1.19 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -2.12 DPS) [quest]; Maiden's Circle (13001, -2.12 DPS) [world_drop]; Naglering (11669, -6.18 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.42 DPS) [quest]; Maiden's Circle (13001, -1.42 DPS) [world_drop]; Naglering (11669, -5.73 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+11.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.84 DPS) [quest]; Weakness Analyzer (272438, -3.31 DPS) [vendor]; Serenity Field (272439, -7.09 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.71 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -14.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+9.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.99 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.83 DPS) [world]; Torch of Light (279246, -9.54 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 33.8. Weights run: 2.3s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.068, intellect=-0.018 ± 0.004, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.438 per %), hit=0.129 ± 0.001 per rating point (10 rating = 1%, 1.285 per %), spell_haste=not significant (0.131 ± 0.070), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.699 ± 0.068, fire_power=0.302 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.90 DPS) | yes | Red Winter Hat (21524, -2.80 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.75 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.15 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.60 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.21 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Green Woolen Vest (2582, -0.15 DPS) [crafted]; Bloody Apron (6226, -0.15 DPS) [dungeon]; Gray Woolen Robe (2585, -1.27 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) (or Owlbeard Bracers (16981)) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.05 DPS) | yes | Gnoll Casting Gloves (892, -0.23 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.45 DPS) [quest]; Pristine Gloves (253913, -0.45 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (33.8 DPS) | yes | Novice Ardent's Sash (253887, -0.30 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.76 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted]; Rumpled Kilt (274741, -0.60 DPS) [vendor]; Silk-threaded Trousers (1929, -0.68 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.05 DPS) | yes | Red Woolen Boots (4313, -0.45 DPS) [crafted]; Feather Padded Treads (285345, -0.52 DPS, sim-verified) [world]; Pristine Boots (253889, -0.60 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.75 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.45 DPS) | yes | Ring of the Shadow (1462, -0.68 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 8 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 151.6 spell_power points (22.76 DPS) | yes | Skycaller (12984, -1.79 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.47 DPS) [dungeon]; Sizzle Stick (8071, -4.36 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 57.2. Weights run: 2.4s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.084, intellect=0.101 ± 0.006, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.429 per %), hit=0.128 ± 0.001 per rating point (10 rating = 1%, 1.279 per %), spell_haste=0.617 ± 0.090, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.775 ± 0.084, fire_power=0.228 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.17 DPS) | yes | Silk Headband (7050, -0.47 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.59 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.59 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (1.50 DPS) | yes | Crystal Starfire Medallion (5003, -1.42 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.42 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.62 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (1.96 DPS) | yes | Invoker's Mantle (215365, -0.48 DPS) [crafted]; Chestnut Mantle (17695, -0.49 DPS, sim-verified) [quest]; Death Speaker Mantle (6685, -0.55 DPS) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.99 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.20 DPS) [crafted]; Battle Healer's Cloak (19529, -0.20 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.57 DPS) | yes | Green Silk Armor (7065, -0.44 DPS, sim-verified) [crafted]; High Robe of the Adjudicator (3461, -0.95 DPS) [quest]; Death Speaker Robes (6682, -0.97 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.78 DPS) | yes | Windsong Bangles (263336, -1.58 DPS) [quest]; Glowing Magical Bracelets (13106, -1.62 DPS) [world_drop]; Owlbeard Bracers (16981, -1.68 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.38 DPS) | yes | Jutebraid Gloves (10654, -0.10 DPS) [quest]; Gnoll Casting Gloves (892, -0.20 DPS) [world]; Truefaith Gloves (7049, -0.34 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.3 spell_power points (2.23 DPS) | yes | Warsong Sash (16975, -0.06 DPS) [quest]; Belt of Arugal (6392, -0.40 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.65 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.37 DPS) | yes | Abomination Skin Leggings (23173, -0.55 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Silk-threaded Trousers (1929, -0.99 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.7 spell_power points (1.52 DPS) | yes | Acidic Walkers (9454, -0.38 DPS) [dungeon]; Boots of the Enchanter (4325, -0.54 DPS) [crafted]; Spidersilk Boots (4320, -1.51 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.38 DPS) | yes | Electrocutioner Lagnut (9447, -0.79 DPS) [dungeon]; Sludge-Stained Band (286535, -0.79 DPS) [world]; Sacred Band (6669, -0.99 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.19 DPS) | yes | Electrocutioner Lagnut (9447, -0.59 DPS) [dungeon]; Sacred Band (6669, -0.79 DPS) [quest]; Sludge-Stained Band (286535, -2.14 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.78 DPS) | yes | Glimmering Staff (249392, -1.40 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.58 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.58 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.6 spell_power points (1.50 DPS) | yes | Orb of Souls (249395, -0.71 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -0.85 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -1.03 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 171.1 spell_power points (33.80 DPS) | yes | Starfaller (13063, -0.65 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.81 DPS) [crafted]; Gravestone Scepter (7001, -4.80 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 128.3. Weights run: 1.8s. Verify run: 1.0s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.059, intellect=0.206 ± 0.009, crit=0.039 ± 0.001 per rating point (14 rating = 1%, 0.539 per %), hit=0.156 ± 0.001 per rating point (10 rating = 1%, 1.557 per %), spell_haste=0.932 ± 0.071, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.806 ± 0.059, fire_power=0.194 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.87 DPS) | yes | Augural Shroud (2620, -2.06 DPS, sim-verified) [world]; Living Cowl (5608, -2.24 DPS) [world]; Holy Shroud (2721, -2.79 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.2 spell_power points (2.30 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.84 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.90 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.90 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.8 spell_power points (3.03 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Inquisitor's Shawl (19507, -0.33 DPS) [dungeon]; Berylline Pads (4197, -0.50 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 10.8 spell_power points (3.03 DPS) | yes | Guardian Cloak (5965, -1.07 DPS) [crafted]; Icy Cloak (4327, -1.08 DPS) [crafted]; Long Silken Cloak (4326, -1.32 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 spell_power points (6.49 DPS) | yes | Elemental Raiment (9434, -0.62 DPS) [world_drop]; Dreamweave Vest (10021, -0.95 DPS) [crafted]; Robe of Power (7054, -1.89 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.56 DPS) [quest]; Radiant Silver Bracers (4545, -0.94 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.26 DPS) | yes | Black Mageweave Gloves (10003, -0.69 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.61 DPS) [crafted]; Gilded Handwraps (254021, -2.62 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.8 spell_power points (4.14 DPS) | yes | Star Belt (4329, -0.51 DPS) [crafted]; Warsong Sash (16975, -1.07 DPS) [quest]; Deathmage Sash (10771, -1.32 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.5 spell_power points (4.60 DPS) | yes | Gaze Dreamer Pants (6903, -0.77 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.62 DPS) [crafted]; Abomination Skin Leggings (23173, -1.63 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.71 DPS) | yes | Gilded Slippers (254001, -2.70 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.52 DPS) [crafted]; Acidic Walkers (9454, -4.85 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.2 spell_power points (3.14 DPS) | yes | Reedknot Ring (9622, -1.18 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.46 DPS) [vendor]; Sludge-Stained Band (286535, -2.30 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.52 DPS) | yes | Reedknot Ring (9622, -0.56 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor]; Sludge-Stained Band (286535, -1.68 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.07 DPS) [dungeon]; Staff of Dar'Orahil (15106, -4.52 DPS) [quest]; Gut Ripper (2164, -6.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (128.3 DPS) | yes | Umbral Wand (5216, -0.01 DPS) [dungeon]; Earthen Rod (9381, -0.09 DPS) [dungeon]; Jaina's Firestarter (13064, -2.13 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 176.9. Weights run: 1.7s. Verify run: 1.1s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.073, intellect=0.323 ± 0.014, crit=0.059 ± 0.002 per rating point (14 rating = 1%, 0.826 per %), hit=0.227 ± 0.002 per rating point (10 rating = 1%, 2.274 per %), spell_haste=not significant (-0.255 ± 0.140), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.802 ± 0.073, fire_power=0.199 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (6.96 DPS) | yes | Dreamweave Circlet (10041, -0.71 DPS) [crafted]; Red Mageweave Headband (10033, -1.45 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.55 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (2.30 DPS) | yes | Mindburst Medallion (11196, -0.26 DPS) [quest]; Horizon Choker (13085, -1.14 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -1.47 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Shoulders (10027, -1.36 DPS) [crafted]; Bloodmage Mantle (7684, -1.62 DPS) [dungeon]; Rotgrip Mantle (17732, -2.18 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (4.11 DPS) | yes | Deep Woodlands Cloak (19121, -0.71 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.04 DPS) [dungeon]; Runecloth Cloak (13860, -1.12 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 25.5 spell_power points (6.56 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -1.15 DPS) [world_drop]; Dreamweave Vest (10021, -1.17 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 9.3 spell_power points (2.39 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.57 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.3 spell_power points (4.97 DPS) | yes | Black Mageweave Gloves (10003, -1.11 DPS) [crafted]; Runecloth Gloves (13863, -1.13 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.63 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Cloth Girdle (20166, -0.13 DPS) [rep]; Ghostweave Cord (254073, -0.46 DPS) [crafted]; Satyrmane Sash (17755, -2.21 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.2 spell_power points (6.76 DPS) | yes | Red Mageweave Pants (10009, -2.15 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.46 DPS) [vendor]; Wizardweave Leggings (14132, -2.89 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.19 DPS) | yes | Gilded Sandals (254107, -2.60 DPS) [crafted]; Black Mageweave Boots (10026, -2.77 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.79 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.35 DPS) | yes | Philanthropist's Ring (281635, -0.27 DPS) [quest]; Cyclopean Band (11824, -0.45 DPS) [dungeon]; Runed Ring (862, -1.55 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.09 DPS) | yes | Philanthropist's Ring (281635, -0.02 DPS) [quest]; Cyclopean Band (11824, -0.19 DPS) [dungeon]; Runed Ring (862, -1.29 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.94 DPS, sim-verified) [crafted]; Uther's Strength (11302, -1.55 DPS) [world_drop]; Rune of the Guard Captain (19120, -2.68 DPS) [quest] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.19 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, -2.57 DPS) [crafted]; Rune of the Guard Captain (19120, -2.68 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arbiter's Blade (11784, -2.68 DPS) [dungeon]; Glowing Brightwood Staff (812, -2.74 DPS) [world_drop]; Blade of Eternal Darkness (17780, -4.81 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.87 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.48 DPS) [crafted]; Pyric Caduceus (11748, -4.19 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 353.5. Weights run: 1.9s. Verify run: 1.2s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.196, intellect=0.123 ± 0.019, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.925 per %), hit=0.258 ± 0.003 per rating point (10 rating = 1%, 2.584 per %), spell_haste=not significant (0.242 ± 0.308), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.883 ± 0.196, fire_power=0.117 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 31.0 spell_power points (14.64 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -3.24 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -5.76 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.40 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.55 DPS) [quest]; Kezan's Taint (19604, -3.32 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.8 spell_power points (14.07 DPS) | yes | Warlord's Dreadweave Mantle (231592, -1.27 DPS) [pvp]; Argent Shoulders (19059, -2.38 DPS, sim-verified) [crafted]; Burial Shawl (18681, -3.69 DPS) [dungeon] |
| back | Crystalline Threaded Cape (20697) | Chillwind Ravager [world] | 20.5 spell_power points (9.68 DPS) | yes | Arcanoweave Cloak (272411, -0.44 DPS) [vendor]; Amplifying Cloak (18350, -1.18 DPS) [dungeon]; Hide of the Wild (18510, -2.49 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.1 spell_power points (22.26 DPS) | yes | Robe of Everlasting Night (18385, -3.40 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Robe (231591, -5.75 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.29 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.0 spell_power points (10.86 DPS) | yes | Sublime Wristguards (18497, -4.61 DPS) [dungeon]; Runecloth Cuffs (254123, -5.08 DPS) [crafted]; Spidertank Oilrag (9448, -6.61 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.6 spell_power points (13.05 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.65 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 30.3 spell_power points (14.31 DPS) | yes | Belt of the Archmage (18405, -3.95 DPS, sim-verified) [crafted]; Frostwolf Cloth Belt (19090, -5.23 DPS) [rep]; Oddly Magical Belt (18475, -6.75 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 35.1 spell_power points (16.58 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.24 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.34 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.95 DPS) [crafted]; Omnicast Boots (11822, -1.19 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.12 DPS) [quest]; Maiden's Circle (13001, -2.12 DPS) [world_drop]; Naglering (11669, -9.58 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.42 DPS) [quest]; Maiden's Circle (13001, -1.42 DPS) [world_drop]; Naglering (11669, -9.22 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.90 DPS, sim-verified) [quest]; Weakness Analyzer (272438, -3.31 DPS) [vendor]; Serenity Field (272439, -7.09 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.71 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -17.30 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (353.5 DPS) | yes | Bonecreeper Stylus (13938, -0.99 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.83 DPS) [world]; Torch of Light (279246, -8.85 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Crystalline Threaded Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

