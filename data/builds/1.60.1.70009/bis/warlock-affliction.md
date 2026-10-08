# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 40.2. Weights run: 2.0s. Verify run: 1.2s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.065, intellect=-0.045 ± 0.004, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.524 per %), hit=0.138 ± 0.004 per rating point (10 rating = 1%, 1.383 per %), spell_haste=not significant (0.047 ± 0.068), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.721 ± 0.065

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.90 DPS) | yes | Red Winter Hat (21524, -3.20 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (0.75 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.15 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.60 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Caretaker's Cape (20428, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Green Woolen Vest (2582, -0.15 DPS) [crafted]; Bloody Apron (6226, -0.15 DPS) [dungeon]; Gray Woolen Robe (2585, -1.40 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Ivycloth Bracelets (9793, -0.25 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.05 DPS) | yes | Gnoll Casting Gloves (892, -0.30 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.45 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.75 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (40.2 DPS) | yes | Novice Ardent's Sash (253887, -0.30 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.04 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted]; Silk-threaded Trousers (1929, -0.52 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.60 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.05 DPS) | yes | Red Woolen Boots (4313, -0.45 DPS) [crafted]; Pristine Boots (253889, -0.60 DPS) [crafted]; Feather Padded Treads (285345, -0.65 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 spell_power points (0.75 DPS) | yes | Sludge-Stained Band (286535, -0.30 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.75 DPS) | yes | Sludge-Stained Band (286535, -0.62 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Staff of Westfall (2042), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), Lesser Staff of the Spire (1300), and 88 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, -1.20 DPS, sim-verified) [world_drop] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 spell_power points (0.75 DPS) | yes | Bouquet of Red Roses (22206, -1.33 DPS, sim-verified) [dungeon] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 151.6 spell_power points (22.76 DPS) | yes | Skycaller (12984, -1.71 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.47 DPS) [dungeon]; Deepblaze (279896, -4.23 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 25552200000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 65.7. Weights run: 2.1s. Verify run: 1.2s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.047, intellect=0.118 ± 0.003, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.360 per %), hit=0.091 ± 0.004 per rating point (10 rating = 1%, 0.915 per %), spell_haste=not significant (-0.054 ± 0.051), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.845 ± 0.047

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (3.15 DPS) | yes | Silk Headband (7050, -0.64 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.86 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.86 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.7 spell_power points (2.21 DPS) | yes | Crystal Starfire Medallion (5003, -2.07 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.07 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.07 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.1 spell_power points (2.88 DPS) | yes | Invoker's Mantle (215365, -0.73 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.79 DPS) [dungeon]; Fairywing Mantle (9536, -0.86 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.43 DPS) | yes | Prelacy Cape (7004, -0.29 DPS) [quest]; Caretaker's Cape (19533, -0.29 DPS) [rep]; Heavy Woolen Cloak (4311, -0.34 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.72 DPS) | yes | Green Silk Armor (7065, -1.25 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.35 DPS) [dungeon]; Robes of Arcana (5770, -1.43 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.58 DPS) | yes | Glowing Magical Bracelets (13106, -2.31 DPS) [world_drop]; Nightsky Wristbands (6407, -2.37 DPS) [world_drop]; Windsong Bangles (263336, -2.68 DPS, sim-verified) [quest] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (2.00 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.29 DPS) [world]; Truefaith Gloves (7049, -0.47 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.4 spell_power points (3.25 DPS) | yes | Belt of Arugal (6392, -0.63 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.96 DPS) [dungeon]; Invoker's Cord (215366, -1.08 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.44 DPS) | yes | Abomination Skin Leggings (23173, -1.05 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.19 DPS) [crafted]; Silk-threaded Trousers (1929, -1.43 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.8 spell_power points (2.24 DPS) | yes | Nimbus Boots (6998, -0.52 DPS) [quest]; Acidic Walkers (9454, -0.54 DPS) [dungeon]; Spidersilk Boots (4320, -1.94 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.00 DPS) | yes | Minor Channeling Ring (1449, -0.50 DPS) [quest]; Electrocutioner Lagnut (9447, -1.15 DPS) [dungeon]; Sludge-Stained Band (286535, -1.15 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.72 DPS) | yes | Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Sludge-Stained Band (286535, -0.86 DPS) [world]; Minor Channeling Ring (1449, -1.89 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.58 DPS) | yes | Glimmering Staff (249392, -1.98 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -2.24 DPS) [world_drop]; Channeler's Staff (4437, -2.31 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.7 spell_power points (2.21 DPS) | yes | Dwarven Tome (279898, -0.76 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.06 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -1.06 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 119.0 spell_power points (34.07 DPS) | yes | Starfaller (13063, -0.91 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.64 DPS) [crafted]; Gravestone Scepter (7001, -5.07 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 25552300120201201-0000000000000000000-0000000000000000)

Set DPS (verified): 175.8. Weights run: 2.0s. Verify run: 1.2s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.219, intellect=0.529 ± 0.012, crit=0.025 ± 0.001 per rating point (14 rating = 1%, 0.348 per %), hit=0.147 ± 0.011 per rating point (10 rating = 1%, 1.468 per %), spell_haste=not significant (-0.042 ± 0.162), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.885 ± 0.219

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.76 DPS) | yes | Living Cowl (5608, -2.58 DPS) [world]; Enchanter's Cowl (4322, -3.12 DPS) [crafted]; Augural Shroud (2620, -3.67 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 spell_power points (3.28 DPS) | yes | Triune Amulet (7722, -2.08 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.08 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.87 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.9 spell_power points (4.47 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Bloodmage Mantle (7684, -0.04 DPS) [dungeon]; Berylline Pads (4197, -0.51 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.8 spell_power points (4.43 DPS) | yes | Guardian Cloak (5965, -1.65 DPS) [crafted]; Icy Cloak (4327, -2.18 DPS) [crafted]; Long Silken Cloak (4326, -2.90 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.2 spell_power points (8.11 DPS) | yes | Elemental Raiment (9434, -1.34 DPS) [world_drop]; Robe of Power (7054, -1.55 DPS) [crafted]; Dreamweave Vest (10021, -1.82 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.90 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.64 DPS) [quest]; Windchaser Cuffs (14429, -1.36 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 spell_power points (6.48 DPS) | yes | Black Mageweave Gloves (10003, -1.65 DPS) [crafted]; Gilded Handwraps (254021, -2.71 DPS) [crafted]; Red Mageweave Gloves (10018, -2.94 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.1 spell_power points (5.19 DPS) | yes | Star Belt (4329, -1.00 DPS) [crafted]; Gilded Cord (254037, -1.25 DPS) [crafted]; Deathmage Sash (10771, -2.22 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.4 spell_power points (6.55 DPS) | yes | Abomination Skin Leggings (23173, -2.29 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.36 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.69 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.73 DPS) | yes | Gilded Slippers (254001, -3.88 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.75 DPS) [dungeon]; Spidersilk Boots (4320, -4.79 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (4.24 DPS) | yes | Ring of Forlorn Spirits (2043, -1.67 DPS) [quest]; Reedknot Ring (9622, -1.99 DPS) [quest]; Minor Channeling Ring (1449, -2.29 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.90 DPS) | yes | Ring of Forlorn Spirits (2043, -0.52 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.64 DPS) [quest]; Minor Channeling Ring (1449, -0.95 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.54 DPS) [dungeon]; Windweaver Staff (7757, -3.88 DPS) [dungeon]; Gut Ripper (2164, -10.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (175.8 DPS) | yes | Umbral Wand (5216, -0.26 DPS) [dungeon]; Earthen Rod (9381, -0.34 DPS) [dungeon]; Jaina's Firestarter (13064, -4.39 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25552300120201351-2002000000000000000-0000000000000000)

Set DPS (verified): 234.3. Weights run: 1.9s. Verify run: 1.3s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.139, intellect=0.155 ± 0.011, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.386 per %), hit=0.231 ± 0.011 per rating point (10 rating = 1%, 2.309 per %), spell_haste=not significant (-0.239 ± 0.114), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.908 ± 0.139

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.99 DPS) | yes | Dreamweave Circlet (10041, -1.17 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -1.99 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.44 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (3.23 DPS) | yes | Mindburst Medallion (11196, -0.41 DPS) [quest]; Horizon Choker (13085, -2.34 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.60 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.0 spell_power points (6.52 DPS) | yes | Rotgrip Mantle (17732, -1.84 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.88 DPS) [crafted]; Bloodmage Mantle (7684, -2.29 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.9 spell_power points (6.08 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.85 DPS) [dungeon]; Runecloth Cloak (13860, -1.91 DPS) [crafted]; Nightfall Drape (12465, -2.41 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (9.33 DPS) | yes | Acumen Robes (17775, -0.34 DPS) [quest]; Elemental Raiment (9434, -0.79 DPS) [world_drop]; Dreamweave Vest (10021, -1.44 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.66 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.37 DPS) [crafted]; Condor Bracers (15864, -0.81 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (7.58 DPS) | yes | Black Mageweave Gloves (10003, -0.78 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.72 DPS) [vendor]; Runecloth Gloves (13863, -2.13 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.0 spell_power points (6.52 DPS) | yes | Highlander's Cloth Girdle (20098, -0.57 DPS) [rep]; Ghostweave Cord (254073, -0.82 DPS) [crafted]; Satyrmane Sash (17755, -1.76 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 24.6 spell_power points (9.99 DPS) | yes | Red Mageweave Pants (10009, -3.54 DPS) [crafted]; Wizardweave Leggings (14132, -3.81 DPS, sim-verified) [crafted]; Knight's Dreadweave Leggings (220888, -4.19 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.77 DPS) | yes | Gilded Sandals (254107, -4.72 DPS) [crafted]; Black Mageweave Boots (10026, -4.85 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -5.00 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.29 DPS) | yes | Philanthropist's Ring (281635, -0.84 DPS) [quest]; Cyclopean Band (11824, -1.19 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -2.03 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.88 DPS) | yes | Philanthropist's Ring (281635, -0.43 DPS) [quest]; Cyclopean Band (11824, -0.78 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.63 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.62 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -4.48 DPS) [dungeon]; Arbiter's Blade (11784, -4.57 DPS) [dungeon]; Blade of Eternal Darkness (17780, -8.14 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (234.3 DPS) | yes | Lesser Eternal Wand (249232, -3.03 DPS) [crafted]; Wand of Allistarj (13065, -3.61 DPS) [world_drop]; Pyric Caduceus (11748, -9.32 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 483.8. Weights run: 5.5s. Verify run: 1.3s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.448, intellect=0.035 ± 0.028, crit=0.114 ± 0.003 per rating point (14 rating = 1%, 1.596 per %), hit=0.644 ± 0.043 per rating point (10 rating = 1%, 6.437 per %), spell_haste=not significant (1.360 ± 0.388), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.849 ± 0.448)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.3 spell_power points (8.13 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Deathmist Mask (226909, -1.34 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -7.04 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (5.91 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -1.69 DPS) [dungeon]; Amulet of the Dawn (22657, -1.76 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.1 spell_power points (7.81 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.95 DPS) [pvp]; Argent Shoulders (19059, -2.14 DPS, sim-verified) [crafted]; Burial Shawl (18681, -2.30 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.7 spell_power points (6.10 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -1.27 DPS) [dungeon]; Hide of the Wild (18510, -2.25 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.3 spell_power points (12.43 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.62 DPS) [pvp]; Robe of Everlasting Night (18385, -5.17 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -5.53 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.3 spell_power points (5.98 DPS) | yes | Sublime Wristguards (18497, -2.67 DPS) [dungeon]; Runecloth Cuffs (254123, -2.93 DPS) [crafted]; Arcane Runed Bracers (4744, -3.56 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.13 DPS) [quest]; Sandworm Skin Gloves (20716, -4.94 DPS, sim-verified) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.2 spell_power points (8.64 DPS) | yes | Ban'thok Sash (11662, -3.59 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -3.72 DPS) [rep]; Belt of the Archmage (18405, -6.21 DPS, sim-verified) [crafted] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Sentinel's Silk Leggings (22752, -1.51 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.44 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.54 DPS) [crafted]; Omnicast Boots (11822, -0.96 DPS) [dungeon] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.32 DPS) [quest]; Maiden's Circle (13001, -1.13 DPS) [world_drop]; Naglering (11669, -15.59 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.31 DPS) [quest]; Maiden's Circle (13001, -1.11 DPS) [world_drop]; Naglering (11669, -11.52 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -1.61 DPS) [quest]; Weakness Analyzer (272438, -1.88 DPS) [vendor]; Draconic Infused Emblem (22268, -7.37 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -1.70 DPS, sim-verified) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.66 DPS) [world]; Teebu's Blazing Longsword (1728, -24.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+22.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.18 DPS) [dungeon]; Wand of Biting Cold (19108, -2.80 DPS) [quest]; Torch of Light (279246, -22.21 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60, raid preset (gnome, 25550300100201351-0005200000000000000-0550001000000000)

Set DPS (verified): 903.9. Weights run: 6.4s. Verify run: 1.5s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.311, intellect=0.102 ± 0.039, crit=0.300 ± 0.013 per rating point (14 rating = 1%, 4.197 per %), hit=0.760 ± 0.075 per rating point (10 rating = 1%, 7.601 per %), spell_haste=not significant (-0.108 ± 0.966), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.871 ± 0.311)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.8 spell_power points (14.19 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.74 DPS) [pvp]; Deathmist Mask (226909, -7.22 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.13 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.11 DPS) [dungeon]; Amulet of the Dawn (22657, -2.62 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 32.7 spell_power points (15.07 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.76 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.70 DPS) [crafted]; Argent Shoulders (19059, -7.12 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.4 spell_power points (11.25 DPS) | yes | Crystalline Threaded Cape (20697, -1.85 DPS) [world]; Amplifying Cloak (18350, -2.95 DPS) [dungeon]; Hide of the Wild (18510, -4.33 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.9 spell_power points (21.61 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.75 DPS) [pvp]; Robe of Everlasting Night (18385, -8.57 DPS) [dungeon]; Knight-Captain's Dreadweave Tunic (227096, -9.16 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.8 spell_power points (10.51 DPS) | yes | Sublime Wristguards (18497, -4.51 DPS) [dungeon]; Runecloth Cuffs (254123, -4.97 DPS) [crafted]; Arcane Runed Bracers (4744, -6.36 DPS) [quest] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.5 spell_power points (12.67 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.54 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 34.8 spell_power points (16.05 DPS) | yes | Belt of the Archmage (18405, -5.75 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -6.50 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -7.29 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 41.4 spell_power points (19.08 DPS) | yes | Skyshroud Leggings (13170, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Leggings (231587, -1.14 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -5.57 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.06 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.92 DPS) [crafted]; Omnicast Boots (11822, -1.28 DPS) [dungeon] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.87 DPS) [quest]; Maiden's Circle (13001, -2.25 DPS) [world_drop]; Naglering (11669, -20.60 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.65 DPS) [quest]; Maiden's Circle (13001, -2.03 DPS) [world_drop]; Naglering (11669, -17.26 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (+25.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18467, -2.76 DPS) [quest]; Weakness Analyzer (272438, -3.22 DPS) [vendor]; Serenity Field (272439, -6.91 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.71 DPS) [world]; Teebu's Blazing Longsword (1728, -33.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (903.9 DPS) | yes | Bonecreeper Stylus (13938, -1.03 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.86 DPS) [world]; Torch of Light (279246, -52.35 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.0. Weights run: 2.0s. Verify run: 1.2s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.065, intellect=-0.045 ± 0.004, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.524 per %), hit=0.138 ± 0.004 per rating point (10 rating = 1%, 1.383 per %), spell_haste=not significant (0.047 ± 0.068), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.721 ± 0.065

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.90 DPS) | yes | Red Winter Hat (21524, -2.93 DPS, sim-verified) [dungeon] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (0.75 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.15 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.29 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.60 DPS) | yes | Feyscale Cloak (6632, -0.15 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.15 DPS) [rep]; Black Whelp Cloak (7283, -0.27 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 spell_power points (0.75 DPS) | yes | Green Woolen Vest (2582, -0.15 DPS) [crafted]; Bloody Apron (6226, -0.15 DPS) [dungeon]; Gray Woolen Robe (2585, -1.35 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) (or Owlbeard Bracers (16981)) | Breaking the Breaker [quest] | 1.0 spell_power points (0.15 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.05 DPS) | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.45 DPS) [quest]; Pristine Gloves (253913, -0.45 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (38.0 DPS) | yes | Novice Ardent's Sash (253887, -0.30 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.14 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted]; Rumpled Kilt (274741, -0.60 DPS) [vendor]; Silk-threaded Trousers (1929, -0.67 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.05 DPS) | yes | Red Woolen Boots (4313, -0.45 DPS) [crafted]; Pristine Boots (253889, -0.60 DPS) [crafted]; Feather Padded Treads (285345, -0.63 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.75 DPS) | yes | - |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.45 DPS) | yes | Ring of the Shadow (1462, -0.66 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) (or Twisted Chanter's Staff (890), Gnarled Necromancer's Staff (251534), Quilboar Toothpick (276888), Channeler's Staff (4437), Witching Stave (1484), and 96 more) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) (or Seer's Fine Stein (7608), Tork Wrench (11855), Spellbinder Orb (15926), Ancestral Orb (15944), Mystic's Sphere (15946), and 8 more) | Winklespark [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Seer's Fine Stein (7608, +0.00 DPS) [world_drop] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 151.6 spell_power points (22.76 DPS) | yes | Skycaller (12984, -1.82 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.47 DPS) [dungeon]; Sizzle Stick (8071, -4.36 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Evocator's Blade; off_hand: Defective Samophlange; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 25552200000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 64.6. Weights run: 2.1s. Verify run: 1.2s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.047, intellect=0.118 ± 0.003, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.360 per %), hit=0.091 ± 0.004 per rating point (10 rating = 1%, 0.915 per %), spell_haste=not significant (-0.054 ± 0.051), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.845 ± 0.047

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (3.15 DPS) | yes | Silk Headband (7050, -0.70 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.86 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.86 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.7 spell_power points (2.21 DPS) | yes | Darkspear Warding Pendant (272075, -1.95 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.07 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.07 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.1 spell_power points (2.88 DPS) | yes | Invoker's Mantle (215365, -0.71 DPS) [crafted]; Death Speaker Mantle (6685, -0.79 DPS) [dungeon]; Chestnut Mantle (17695, -2.05 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.43 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.29 DPS) [crafted]; Battle Healer's Cloak (19529, -0.29 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.72 DPS) | yes | Green Silk Armor (7065, -0.67 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.35 DPS) [dungeon]; High Robe of the Adjudicator (3461, -1.36 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.58 DPS) | yes | Owlbeard Bracers (16981, -2.22 DPS, sim-verified) [quest]; Windsong Bangles (263336, -2.29 DPS) [quest]; Glowing Magical Bracelets (13106, -2.31 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.00 DPS) | yes | Jutebraid Gloves (10654, -0.12 DPS) [quest]; Gnoll Casting Gloves (892, -0.29 DPS) [world]; Truefaith Gloves (7049, -0.47 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.4 spell_power points (3.25 DPS) | yes | Warsong Sash (16975, -0.10 DPS) [quest]; Belt of Arugal (6392, -0.57 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.96 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.44 DPS) | yes | Abomination Skin Leggings (23173, -0.49 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.19 DPS) [crafted]; Silk-threaded Trousers (1929, -1.43 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.8 spell_power points (2.24 DPS) | yes | Acidic Walkers (9454, -0.54 DPS) [dungeon]; Boots of the Enchanter (4325, -0.81 DPS) [crafted]; Spidersilk Boots (4320, -1.89 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.00 DPS) | yes | Electrocutioner Lagnut (9447, -1.15 DPS) [dungeon]; Sludge-Stained Band (286535, -1.15 DPS) [world]; Sacred Band (6669, -1.43 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.72 DPS) | yes | Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Sacred Band (6669, -1.15 DPS) [quest]; Sludge-Stained Band (286535, -2.41 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.58 DPS) | yes | Glimmering Staff (249392, -1.76 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -2.24 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -2.24 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.7 spell_power points (2.21 DPS) | yes | Orb of Souls (249395, -1.06 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -1.49 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -1.50 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 119.0 spell_power points (34.07 DPS) | yes | Starfaller (13063, -0.67 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.64 DPS) [crafted]; Gravestone Scepter (7001, -5.07 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552300120201201-0000000000000000000-0000000000000000)

Set DPS (verified): 174.8. Weights run: 2.0s. Verify run: 1.2s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.219, intellect=0.529 ± 0.012, crit=0.025 ± 0.001 per rating point (14 rating = 1%, 0.348 per %), hit=0.147 ± 0.011 per rating point (10 rating = 1%, 1.468 per %), spell_haste=not significant (-0.042 ± 0.162), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.885 ± 0.219

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.76 DPS) | yes | Living Cowl (5608, -2.58 DPS) [world]; Enchanter's Cowl (4322, -3.12 DPS) [crafted]; Augural Shroud (2620, -3.64 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 spell_power points (3.28 DPS) | yes | Triune Amulet (7722, -2.08 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.08 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.68 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.9 spell_power points (4.47 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Bloodmage Mantle (7684, -0.04 DPS) [dungeon]; Berylline Pads (4197, -0.51 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.8 spell_power points (4.43 DPS) | yes | Guardian Cloak (5965, -1.65 DPS) [crafted]; Icy Cloak (4327, -2.18 DPS) [crafted]; Long Silken Cloak (4326, -2.86 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.2 spell_power points (8.11 DPS) | yes | Elemental Raiment (9434, -1.34 DPS) [world_drop]; Robe of Power (7054, -1.55 DPS) [crafted]; Dreamweave Vest (10021, -1.80 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.90 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.64 DPS) [quest]; Radiant Silver Bracers (4545, -1.93 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 spell_power points (6.48 DPS) | yes | Black Mageweave Gloves (10003, -1.65 DPS) [crafted]; Red Mageweave Gloves (10018, -2.56 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.71 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.1 spell_power points (5.19 DPS) | yes | Star Belt (4329, -1.00 DPS) [crafted]; Gilded Cord (254037, -1.25 DPS) [crafted]; Deathmage Sash (10771, -2.05 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.4 spell_power points (6.55 DPS) | yes | Abomination Skin Leggings (23173, -2.29 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.66 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.69 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.73 DPS) | yes | Gilded Slippers (254001, -4.16 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.75 DPS) [dungeon]; Spidersilk Boots (4320, -4.79 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 spell_power points (4.24 DPS) | yes | Reedknot Ring (9622, -1.99 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.31 DPS) [vendor]; Black Widow Band (6199, -3.05 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.90 DPS) | yes | Sea Giant's Toe Ring (274746, -0.97 DPS) [vendor]; Reedknot Ring (9622, -1.51 DPS, sim-verified) [quest]; Black Widow Band (6199, -1.70 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -3.54 DPS) [dungeon]; Windweaver Staff (7757, -3.88 DPS) [dungeon]; Gut Ripper (2164, -10.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (174.8 DPS) | yes | Umbral Wand (5216, -0.26 DPS) [dungeon]; Earthen Rod (9381, -0.34 DPS) [dungeon]; Jaina's Firestarter (13064, -5.01 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552300120201351-2002000000000000000-0000000000000000)

Set DPS (verified): 233.6. Weights run: 1.9s. Verify run: 1.3s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.139, intellect=0.155 ± 0.011, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.386 per %), hit=0.231 ± 0.011 per rating point (10 rating = 1%, 2.309 per %), spell_haste=not significant (-0.239 ± 0.114), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.908 ± 0.139

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.99 DPS) | yes | Red Mageweave Headband (10033, -1.99 DPS) [crafted]; Dreamweave Circlet (10041, -2.11 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -2.44 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 spell_power points (3.23 DPS) | yes | Mindburst Medallion (11196, -0.41 DPS) [quest]; Horizon Choker (13085, -2.34 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -2.60 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.0 spell_power points (6.52 DPS) | yes | Rotgrip Mantle (17732, -1.78 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.88 DPS) [crafted]; Bloodmage Mantle (7684, -2.29 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.9 spell_power points (6.08 DPS) | yes | Deep Woodlands Cloak (19121, -0.61 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.85 DPS) [dungeon]; Runecloth Cloak (13860, -1.91 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (9.33 DPS) | yes | Elemental Raiment (9434, -0.79 DPS) [world_drop]; Acumen Robes (17775, -0.98 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -1.44 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.66 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (7.58 DPS) | yes | Black Mageweave Gloves (10003, -0.90 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.72 DPS) [vendor]; Runecloth Gloves (13863, -2.13 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.0 spell_power points (6.52 DPS) | yes | Defiler's Cloth Girdle (20166, -0.57 DPS) [rep]; Ghostweave Cord (254073, -0.82 DPS) [crafted]; Satyrmane Sash (17755, -2.12 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 24.6 spell_power points (9.99 DPS) | yes | Red Mageweave Pants (10009, -3.54 DPS) [crafted]; Wizardweave Leggings (14132, -4.05 DPS, sim-verified) [crafted]; Stone Guard's Dreadweave Leggings (220906, -4.19 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.77 DPS) | yes | Gilded Sandals (254107, -0.85 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -4.85 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -5.00 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.29 DPS) | yes | Philanthropist's Ring (281635, -0.84 DPS) [quest]; Cyclopean Band (11824, -1.19 DPS) [dungeon]; Runed Ring (862, -2.44 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.88 DPS) | yes | Philanthropist's Ring (281635, -0.43 DPS) [quest]; Cyclopean Band (11824, -0.78 DPS) [dungeon]; Runed Ring (862, -2.03 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.44 DPS) [world_drop]; Rune of the Guard Captain (19120, -4.23 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.19 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -4.48 DPS) [dungeon]; Arbiter's Blade (11784, -4.57 DPS) [dungeon]; Blade of Eternal Darkness (17780, -8.94 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (233.6 DPS) | yes | Lesser Eternal Wand (249232, -3.03 DPS) [crafted]; Wand of Allistarj (13065, -3.61 DPS) [world_drop]; Pyric Caduceus (11748, -9.14 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Kentic Amice; back: Spritecaster Cape; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552300120201351-2005220000000000000-0030000000000000)

Set DPS (verified): 474.0. Weights run: 5.5s. Verify run: 1.4s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.448, intellect=0.035 ± 0.028, crit=0.114 ± 0.003 per rating point (14 rating = 1%, 1.596 per %), hit=0.644 ± 0.043 per rating point (10 rating = 1%, 6.437 per %), spell_haste=not significant (1.360 ± 0.388), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.849 ± 0.448)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.3 spell_power points (8.13 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Deathmist Mask (226909, -1.34 DPS) [quest]; Spellpower Goggles Xtreme Plus (15999, -7.34 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (5.91 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -1.69 DPS) [dungeon]; Amulet of the Dawn (22657, -1.76 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 29.1 spell_power points (7.81 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.95 DPS) [pvp]; Burial Shawl (18681, -2.30 DPS) [dungeon]; Argent Shoulders (19059, -5.11 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 22.7 spell_power points (6.10 DPS) | yes | Crystalline Threaded Cape (20697, +0.00 DPS, sim-verified) [world]; Amplifying Cloak (18350, -1.27 DPS) [dungeon]; Hide of the Wild (18510, -2.25 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.3 spell_power points (12.43 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.62 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -5.53 DPS) [pvp]; Robe of Everlasting Night (18385, -6.44 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.3 spell_power points (5.98 DPS) | yes | Sublime Wristguards (18497, -2.67 DPS) [dungeon]; Runecloth Cuffs (254123, -2.93 DPS) [crafted]; Spidertank Oilrag (9448, -3.56 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.2 spell_power points (7.29 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -1.39 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 32.2 spell_power points (8.64 DPS) | yes | Ban'thok Sash (11662, -3.59 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -3.72 DPS) [rep]; Belt of the Archmage (18405, -7.69 DPS, sim-verified) [crafted] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.51 DPS) [rep]; Sentinel's Silk Leggings (237815, -5.36 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.44 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.54 DPS) [crafted]; Omnicast Boots (11822, -0.96 DPS) [dungeon] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.32 DPS) [quest]; Maiden's Circle (13001, -1.13 DPS) [world_drop]; Naglering (11669, -15.76 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.31 DPS) [quest]; Maiden's Circle (13001, -1.11 DPS) [world_drop]; Naglering (11669, -14.39 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.88 DPS) [vendor]; Serenity Field (272439, -4.03 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -4.55 DPS, sim-verified) [quest] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.66 DPS) [world]; Teebu's Blazing Longsword (1728, -27.38 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+24.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -1.18 DPS) [dungeon]; Wand of Biting Cold (19108, -2.80 DPS) [quest]; Torch of Light (279246, -24.81 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Skyshroud Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Draconic Infused Emblem; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (troll, 25550300100201351-0005200000000000000-0550001000000000)

Set DPS (verified): 901.8. Weights run: 6.4s. Verify run: 1.5s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.311, intellect=0.102 ± 0.039, crit=0.300 ± 0.013 per rating point (14 rating = 1%, 4.197 per %), hit=0.760 ± 0.075 per rating point (10 rating = 1%, 7.601 per %), spell_haste=not significant (-0.108 ± 0.966), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.871 ± 0.311)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 30.8 spell_power points (14.19 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.74 DPS) [pvp]; Deathmist Mask (226909, -6.10 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.13 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, -2.11 DPS) [dungeon]; Amulet of the Dawn (22657, -2.62 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 32.7 spell_power points (15.07 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.76 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.70 DPS) [crafted]; Argent Shoulders (19059, -7.51 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.4 spell_power points (11.25 DPS) | yes | Amplifying Cloak (18350, -2.95 DPS) [dungeon]; Crystalline Threaded Cape (20697, -3.52 DPS, sim-verified) [world]; Hide of the Wild (18510, -4.33 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.9 spell_power points (21.61 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.75 DPS) [pvp]; Robe of Everlasting Night (18385, -7.76 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Tunic (227094, -9.16 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.8 spell_power points (10.51 DPS) | yes | Sublime Wristguards (18497, -4.51 DPS) [dungeon]; Runecloth Cuffs (254123, -4.97 DPS) [crafted]; Spidertank Oilrag (9448, -6.36 DPS) [dungeon] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 27.5 spell_power points (12.67 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.54 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 34.8 spell_power points (16.05 DPS) | yes | Ban'thok Sash (11662, -6.50 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -7.29 DPS) [rep]; Belt of the Archmage (18405, -8.89 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 41.4 spell_power points (19.08 DPS) | yes | General's Dreadweave Pants (231588, -1.14 DPS) [pvp]; Skyshroud Leggings (13170, -3.04 DPS) [dungeon]; Outrider's Silk Leggings (22747, -5.29 DPS) [rep] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.06 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.92 DPS) [crafted]; Omnicast Boots (11822, -1.28 DPS) [dungeon] |
| finger1 | Rune Band of Wizardry (22339) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.87 DPS) [quest]; Maiden's Circle (13001, -2.25 DPS) [world_drop]; Naglering (11669, -24.75 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.65 DPS) [quest]; Maiden's Circle (13001, -2.03 DPS) [world_drop]; Naglering (11669, -17.89 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -2.76 DPS) [quest]; Weakness Analyzer (272438, -3.22 DPS) [vendor]; Draconic Infused Emblem (22268, -9.23 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.71 DPS) [world]; Teebu's Blazing Longsword (1728, -36.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (901.8 DPS) | yes | Bonecreeper Stylus (13938, -1.03 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.86 DPS) [world]; Torch of Light (279246, -58.16 DPS, sim-verified) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; finger1: Rune Band of Wizardry; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

