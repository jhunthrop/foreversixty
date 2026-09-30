# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.4. Weights run: 1.0s. Verify run: 0.9s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.222, intellect=0.726 ± 0.012, crit=0.021 ± 0.002 per rating point (14 rating = 1%, 0.290 per %), hit=0.089 ± 0.002 per rating point (10 rating = 1%, 0.887 per %), spell_haste=-1.704 ± 0.181, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.74 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 spell_power points (1.13 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.74 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.41 DPS) | yes | Heavy Woolen Cloak (4311, +0.33 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.12 DPS) [dungeon]; Caretaker's Cape (20428, -0.12 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.6 spell_power points (0.85 DPS) | yes | Green Woolen Robe (6243, -0.34 DPS) [crafted]; Mystic's Wrap (14369, -0.35 DPS) [world_drop]; Gray Woolen Robe (2585, -1.25 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 3.6 spell_power points (0.36 DPS) | yes | Bright Bracers (3647, -0.19 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -0.21 DPS) [world_drop]; Repurposed Hair Band (281256, -0.21 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Pristine Gloves (253913, +0.16 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Tomb Robber's Gloves (280096, -0.26 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.9 spell_power points (0.68 DPS) | yes | Keller's Girdle (2911, -0.11 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (38.4 DPS) | yes | Silk-threaded Trousers (1929, -0.33 DPS) [dungeon]; Darkweave Breeches (12987, -0.52 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.74 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (0.97 DPS) | yes | Pristine Boots (253889, -0.40 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.48 DPS) [world]; Red Woolen Boots (4313, -0.58 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (0.63 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Sludge-Stained Band (286535, -0.34 DPS) [world]; Volcanic Rock Ring (12053, -0.42 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.49 DPS) | yes | Sludge-Stained Band (286535, -0.20 DPS) [world]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.11 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 7.3 spell_power points (0.71 DPS) | yes | Channeler's Staff (4437, -0.11 DPS, sim-verified) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world_drop]; Staff of Westfall (2042, -0.36 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 229.9 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.65 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.1. Weights run: 1.0s. Verify run: 0.9s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.064, intellect=0.165 ± 0.004, crit=0.009 ± 0.001 per rating point (14 rating = 1%, 0.129 per %), hit=0.037 ± 0.001 per rating point (10 rating = 1%, 0.366 per %), spell_haste=0.807 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.064

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (4.20 DPS) | yes | Silk Headband (7050, -0.68 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -1.15 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -1.15 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (3.05 DPS) | yes | Darkspear Warding Pendant (272075, -2.24 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.80 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.80 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (4.00 DPS) | yes | Death Speaker Mantle (6685, -1.02 DPS) [dungeon]; Fairywing Mantle (9536, -1.15 DPS) [quest]; Invoker's Mantle (215365, -1.40 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.91 DPS) | yes | Caretaker's Cape (19533, +0.50 DPS, sim-verified) [rep]; Heavy Woolen Cloak (4311, -0.38 DPS) [crafted]; Prelacy Cape (7004, -0.38 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (4.96 DPS) | yes | Green Silk Armor (7065, -0.55 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.60 DPS) [dungeon]; Pristine Gown (253961, -1.85 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Glowing Magical Bracelets (13106, -2.21 DPS, sim-verified) [world_drop]; Windsong Bangles (263336, -3.05 DPS) [quest]; Nightsky Wristbands (6407, -3.06 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (2.67 DPS) | yes | Serpent Gloves (5970, +0.09 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world]; Town Clerk's Mittens (270029, -0.45 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (4.39 DPS) | yes | Belt of Arugal (6392, -0.63 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -1.33 DPS) [dungeon]; Invoker's Cord (215366, -1.40 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (4.58 DPS) | yes | Abomination Skin Leggings (23173, -0.89 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.47 DPS) [crafted]; Silk-threaded Trousers (1929, -1.91 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.2 spell_power points (3.11 DPS) | yes | Acidic Walkers (9454, -0.70 DPS) [dungeon]; Nimbus Boots (6998, -0.82 DPS) [quest]; Spidersilk Boots (4320, -2.36 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.67 DPS) | yes | Minor Channeling Ring (1449, -0.64 DPS) [quest]; Lorekeeper's Ring (20431, -0.76 DPS) [rep]; Electrocutioner Lagnut (9447, -1.53 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (2.29 DPS) | yes | Electrocutioner Lagnut (9447, -1.15 DPS) [dungeon]; Sludge-Stained Band (286535, -1.15 DPS) [world]; Minor Channeling Ring (1449, -1.42 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Glimmering Staff (249392, -2.08 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -2.81 DPS) [world_drop]; Channeler's Staff (4437, -2.93 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (3.05 DPS) | yes | Eye of Paleth (2943, -1.52 DPS) [quest]; Orb of Souls (249395, -1.52 DPS) [crafted]; Dwarven Tome (279898, -1.59 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 90.0 spell_power points (34.36 DPS) | yes | Starfaller (13063, -1.04 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.45 DPS) [crafted]; Thunderwood (13062, -4.95 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 107.7. Weights run: 0.9s. Verify run: 0.8s. 329 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.122, intellect=0.301 ± 0.010, crit=0.021 ± 0.001 per rating point (14 rating = 1%, 0.287 per %), hit=0.087 ± 0.001 per rating point (10 rating = 1%, 0.869 per %), spell_haste=not significant (-0.202 ± 0.176), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.122

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.90 DPS) | yes | Living Cowl (5608, -2.25 DPS) [world]; Augural Shroud (2620, -2.63 DPS, sim-verified) [world]; Holy Shroud (2721, -2.81 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 spell_power points (2.47 DPS) | yes | Necklace of Calisea (1714, -1.88 DPS) [world_drop]; Triune Amulet (7722, -1.88 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.35 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (3.29 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.22 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (107.7 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.14 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.80 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (6.69 DPS) | yes | Elemental Raiment (9434, -0.57 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.87 DPS) [crafted]; Robe of Power (7054, -1.74 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.53 DPS) | yes | Spidertank Oilrag (9448, -0.47 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest]; Earthen Silk Cuffs (254019, -1.41 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (5.40 DPS) | yes | Black Mageweave Gloves (10003, -1.30 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.46 DPS) [crafted]; Gilded Handwraps (254021, -2.56 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.2 spell_power points (4.27 DPS) | yes | Highlander's Cloth Girdle (20099, -0.93 DPS) [rep]; Star Belt (4329, -0.94 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.04 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.6 spell_power points (4.95 DPS) | yes | Crimson Silk Pantaloons (7062, -1.60 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.72 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.74 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.74 DPS) | yes | Gilded Slippers (254001, -2.89 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.44 DPS) [crafted]; Acidic Walkers (9454, -4.66 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.8 spell_power points (3.32 DPS) | yes | Ring of Forlorn Spirits (2043, -1.07 DPS) [quest]; Reedknot Ring (9622, -1.35 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.63 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.53 DPS) | yes | Ring of Forlorn Spirits (2043, -0.48 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.56 DPS) [quest]; Lorekeeper's Ring (19525, -0.56 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Scorn's Focal Dagger (23168, -3.09 DPS) [dungeon]; Windweaver Staff (7757, -4.35 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 141.9 spell_power points (39.89 DPS) | yes | Twisted Nether Wand (249144, +0.98 DPS, sim-verified) [crafted]; Umbral Wand (5216, -4.22 DPS) [world_drop]; Earthen Rod (9381, -4.30 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 329, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 190.2. Weights run: 0.9s. Verify run: 1.0s. 419 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.227, intellect=0.211 ± 0.015, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.457 per %), hit=0.203 ± 0.003 per rating point (10 rating = 1%, 2.027 per %), spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Red Mageweave Headband (10033, -1.42 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.28 DPS) [quest]; Horizon Choker (13085, -1.49 DPS) [world_drop]; Arcane Crystal Pendant (20037, -3.13 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.8 spell_power points (4.71 DPS) | yes | Black Mageweave Shoulders (10027, -1.37 DPS) [crafted]; Bloodmage Mantle (7684, -1.65 DPS) [dungeon]; Kentic Amice (11624, -4.44 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.28 DPS) | yes | Runecloth Cloak (13860, -1.28 DPS) [crafted]; Nightfall Drape (12465, -1.76 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -4.71 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (6.52 DPS) | yes | Acumen Robes (17775, -0.27 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.64 DPS) [world_drop]; Dreamweave Vest (10021, -0.94 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, -0.15 DPS) [crafted]; Spidertank Oilrag (9448, -0.29 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.28 DPS) | yes | Black Mageweave Gloves (10003, -1.03 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.11 DPS) [vendor]; Runecloth Gloves (13863, -1.39 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (4.52 DPS) | yes | Highlander's Cloth Girdle (20098, +0.29 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Ghostweave Cord (254073, -0.59 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.69 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -1.13 DPS) [vendor]; Spellshock Leggings (9484, -6.72 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, +0.20 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.23 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.39 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.49 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (3.36 DPS) | yes | Cyclopean Band (11824, -0.43 DPS) [dungeon]; Lorekeeper's Ring (19524, -0.84 DPS) [rep]; Philanthropist's Ring (281635, -1.31 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.85 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.46 DPS, sim-verified) [crafted] |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Moonshadow Stave (22458, -0.55 DPS) [quest]; Arbiter's Blade (11784, -3.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Woestave (20082, -0.08 DPS) [quest]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Pyric Caduceus (11748, -4.68 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Noxious Shooter

No-known-source sample (15 of 419, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 396.0. Weights run: 0.9s. Verify run: 1.0s. 1012 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.250, intellect=0.107 ± 0.021, crit=0.030 ± 0.002 per rating point (14 rating = 1%, 0.421 per %), hit=0.201 ± 0.003 per rating point (10 rating = 1%, 2.006 per %), spell_haste=-3.778 ± 0.375, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.250

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 67.8 spell_power points (23.92 DPS) | yes | Field Marshal's Coronal (17578, -11.73 DPS) [vendor]; Field Marshal's Coronal (231584, -11.73 DPS) [pvp]; Crimson Felt Hat (18727, -28.83 DPS, sim-verified) [dungeon] |
| neck | Orb of the Darkmoon (19426) (or Chains of the Lich (23125)) | 1200 Tickets - Orb of the Darkmoon [quest] | 22.0 spell_power points (7.76 DPS) | yes | Chains of the Lich (23125, +0.00 DPS, sim-verified) [dungeon]; Arcane Crystal Pendant (20037, -1.89 DPS) [quest]; Amulet of the Dawn (22657, -1.98 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 43.8 spell_power points (15.46 DPS) | yes | Field Marshal's Dreadweave Shoulders (17580, -6.00 DPS) [vendor]; Field Marshal's Dreadweave Shoulders (231583, -6.00 DPS) [pvp]; Rugged Mantle of the Timbermaw (227808, -6.05 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (396.0 DPS) | yes | Amplifying Cloak (18350, -0.30 DPS) [dungeon]; Hide of the Wild (18510, -1.34 DPS) [crafted]; Crystalline Threaded Cape (20697, -9.06 DPS, sim-verified) [world_drop] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 67.7 spell_power points (23.86 DPS) | yes | Robe of the Void (14153, -9.27 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (17581, -11.67 DPS) [vendor]; Field Marshal's Dreadweave Robe (231582, -11.67 DPS) [pvp] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 29.4 spell_power points (10.37 DPS) | yes | Dryad's Wrist Bindings (19596, -3.09 DPS) [rep]; Dryad's Wrist Bindings (19595, -3.99 DPS, sim-verified) [rep]; Dryad's Wrist Bindings (19597, -4.50 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 40.4 spell_power points (14.25 DPS) | yes | Marshal's Dreadweave Gloves (17584, -3.44 DPS) [vendor]; Marshal's Dreadweave Gloves (231586, -3.44 DPS) [pvp]; Sandworm Skin Gloves (20716, -5.87 DPS, sim-verified) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 44.2 spell_power points (15.60 DPS) | yes | Knowledge of the Timbermaw (228190, -6.15 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -7.80 DPS) [crafted]; Felheart Belt (16806, -7.98 DPS) [world_drop] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 59.7 spell_power points (21.04 DPS) | yes | Marshal's Dreadweave Leggings (17579, -7.28 DPS) [vendor]; Marshal's Dreadweave Leggings (231587, -7.28 DPS) [pvp]; Flarecore Leggings (19165, -8.97 DPS, sim-verified) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 39.5 spell_power points (13.93 DPS) | yes | Marshal's Dreadweave Boots (17583, -4.27 DPS) [vendor]; Marshal's Dreadweave Boots (231585, -4.27 DPS) [pvp]; Earthen Silk Slippers (254013, -6.90 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234436, -0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234437, -0.39 DPS) [vendor]; Naglering (11669, -13.29 DPS, sim-verified) [dungeon] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Elemental Focus Band (20682, -0.79 DPS) [world]; Maiden's Circle (13001, -2.35 DPS) [world_drop]; Naglering (11669, -10.67 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, -2.55 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | 0.0 spell_power points (0.00 DPS) | yes | Electrified Dagger (19100, +0.47 DPS, sim-verified) [rep]; Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 223.6 spell_power points (78.88 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -5.76 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.22 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.44 DPS) [world] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ritssyn's Ring of Chaos; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1012, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.7. Weights run: 1.0s. Verify run: 0.9s. 142 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.222, intellect=0.726 ± 0.012, crit=0.021 ± 0.002 per rating point (14 rating = 1%, 0.290 per %), hit=0.089 ± 0.002 per rating point (10 rating = 1%, 0.887 per %), spell_haste=-1.704 ± 0.181, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.71 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 spell_power points (1.13 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.31 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.74 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.41 DPS) | yes | Heavy Woolen Cloak (4311, +0.24 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.12 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.12 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.6 spell_power points (0.85 DPS) | yes | Green Woolen Robe (6243, -0.34 DPS) [crafted]; Mystic's Wrap (14369, -0.35 DPS) [world_drop]; Gray Woolen Robe (2585, -1.40 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 spell_power points (0.43 DPS) | yes | Mindthrust Bracers (1974, -0.07 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Bright Bracers (3647, -0.14 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Blight Gloves (279877, -0.19 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.9 spell_power points (0.68 DPS) | yes | Keller's Girdle (2911, -0.11 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.10 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (36.7 DPS) | yes | Silk-threaded Trousers (1929, -0.33 DPS) [dungeon]; Darkweave Breeches (12987, -0.52 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.61 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (0.97 DPS) | yes | Pristine Boots (253889, -0.41 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.48 DPS) [world]; Red Woolen Boots (4313, -0.58 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.49 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Sludge-Stained Band (286535, -0.20 DPS) [world]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.4 spell_power points (0.43 DPS) | yes | Sludge-Stained Band (286535, -0.13 DPS) [world]; Loop of Sacrifice (281673, -0.14 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.3 spell_power points (0.71 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.14 DPS) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 229.9 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.88 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Sizzle Stick (8071, -4.46 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 142, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 58.6. Weights run: 1.0s. Verify run: 0.9s. 237 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.064, intellect=0.165 ± 0.004, crit=0.009 ± 0.001 per rating point (14 rating = 1%, 0.129 per %), hit=0.037 ± 0.001 per rating point (10 rating = 1%, 0.366 per %), spell_haste=0.807 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.064

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (4.20 DPS) | yes | Silk Headband (7050, -0.76 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -1.15 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -1.15 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (3.05 DPS) | yes | Darkspear Warding Pendant (272075, -1.68 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.80 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.80 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (4.00 DPS) | yes | Invoker's Mantle (215365, -1.02 DPS) [crafted]; Death Speaker Mantle (6685, -1.02 DPS) [dungeon]; Chestnut Mantle (17695, -1.38 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.91 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.38 DPS) [crafted]; Battle Healer's Cloak (19529, -0.38 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (4.96 DPS) | yes | Green Silk Armor (7065, -0.69 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.60 DPS) [dungeon]; High Robe of the Adjudicator (3461, -1.78 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Owlbeard Bracers (16981, -1.65 DPS, sim-verified) [quest]; Glowing Magical Bracelets (13106, -2.93 DPS) [world_drop]; Windsong Bangles (263336, -3.05 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.67 DPS) | yes | Jutebraid Gloves (10654, -0.06 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.38 DPS) [world]; Truefaith Gloves (7049, -0.57 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (4.39 DPS) | yes | Warsong Sash (16975, +0.00 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.76 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -1.33 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (4.58 DPS) | yes | Abomination Skin Leggings (23173, +0.10 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.47 DPS) [crafted]; Silk-threaded Trousers (1929, -1.91 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.2 spell_power points (3.11 DPS) | yes | Acidic Walkers (9454, -0.70 DPS) [dungeon]; Boots of the Enchanter (4325, -1.20 DPS) [crafted]; Spidersilk Boots (4320, -1.95 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.67 DPS) | yes | Advisor's Ring (20426, -0.76 DPS) [rep]; Electrocutioner Lagnut (9447, -1.53 DPS) [dungeon]; Sludge-Stained Band (286535, -1.53 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (2.29 DPS) | yes | Sludge-Stained Band (286535, -1.15 DPS) [world]; Sacred Band (6669, -1.53 DPS) [quest]; Electrocutioner Lagnut (9447, -1.93 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Glimmering Staff (249392, -1.62 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -2.81 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -2.81 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (3.05 DPS) | yes | Orb of Souls (249395, -0.61 DPS, sim-verified) [crafted]; Alliance Outrunner Healing Rod (285348, -1.52 DPS) [world]; Tome of the Darkspear Prophecy (272090, -2.04 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 90.0 spell_power points (34.36 DPS) | yes | Starfaller (13063, -1.06 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.45 DPS) [crafted]; Thunderwood (13062, -4.95 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 237, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 107.4. Weights run: 0.9s. Verify run: 0.8s. 320 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.122, intellect=0.301 ± 0.010, crit=0.021 ± 0.001 per rating point (14 rating = 1%, 0.287 per %), hit=0.087 ± 0.001 per rating point (10 rating = 1%, 0.869 per %), spell_haste=not significant (-0.202 ± 0.176), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.122

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.90 DPS) | yes | Augural Shroud (2620, -2.21 DPS, sim-verified) [world]; Living Cowl (5608, -2.25 DPS) [world]; Holy Shroud (2721, -2.81 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 spell_power points (2.47 DPS) | yes | Necklace of Calisea (1714, -1.88 DPS) [world_drop]; Triune Amulet (7722, -1.88 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.14 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (3.29 DPS) | yes | Inquisitor's Shawl (19507, -0.22 DPS) [dungeon]; Green Silken Shoulders (7057, -0.26 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.14 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.80 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (6.69 DPS) | yes | Elemental Raiment (9434, -0.29 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.87 DPS) [crafted]; Robe of Power (7054, -1.74 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.53 DPS) | yes | Condor Bracers (15864, -0.64 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.73 DPS) [quest]; Earthen Silk Cuffs (254019, -1.41 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (5.40 DPS) | yes | Black Mageweave Gloves (10003, -0.96 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.46 DPS) [crafted]; Gilded Handwraps (254021, -2.56 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.2 spell_power points (4.27 DPS) | yes | Star Belt (4329, -0.46 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.93 DPS) [rep]; Deathmage Sash (10771, -1.04 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.6 spell_power points (4.95 DPS) | yes | Gaze Dreamer Pants (6903, -0.69 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.60 DPS) [crafted]; Abomination Skin Leggings (23173, -1.74 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.74 DPS) | yes | Gilded Slippers (254001, -2.25 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.44 DPS) [crafted]; Acidic Walkers (9454, -4.66 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.8 spell_power points (3.32 DPS) | yes | Reedknot Ring (9622, -1.35 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.63 DPS) [vendor]; Electrocutioner Lagnut (9447, -2.47 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.53 DPS) | yes | Advisor's Ring (19521, -0.56 DPS) [rep]; Reedknot Ring (9622, -0.64 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Scorn's Focal Dagger (23168, -3.09 DPS) [dungeon]; Windweaver Staff (7757, -4.35 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Umbral Wand (5216, -0.02 DPS) [world_drop]; Earthen Rod (9381, -0.10 DPS) [dungeon]; Jaina's Firestarter (13064, -1.56 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 320, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 187.2. Weights run: 0.9s. Verify run: 0.9s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.227, intellect=0.211 ± 0.015, crit=0.033 ± 0.002 per rating point (14 rating = 1%, 0.457 per %), hit=0.203 ± 0.003 per rating point (10 rating = 1%, 2.027 per %), spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.68 DPS) [crafted]; Red Mageweave Headband (10033, -1.76 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.28 DPS) [quest]; Horizon Choker (13085, -1.49 DPS) [world_drop]; Arcane Crystal Pendant (20037, -2.74 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.8 spell_power points (4.71 DPS) | yes | Black Mageweave Shoulders (10027, -1.37 DPS) [crafted]; Bloodmage Mantle (7684, -1.65 DPS) [dungeon]; Kentic Amice (11624, -5.40 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.28 DPS) | yes | Deep Woodlands Cloak (19121, -0.45 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.22 DPS) [dungeon]; Runecloth Cloak (13860, -1.28 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (6.52 DPS) | yes | Acumen Robes (17775, -0.62 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.64 DPS) [world_drop]; Dreamweave Vest (10021, -0.94 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, +1.57 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest]; Bloodband Bracers (11469, -0.59 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.28 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.11 DPS) [vendor]; Black Mageweave Gloves (10003, -1.22 DPS, sim-verified) [crafted]; Runecloth Gloves (13863, -1.39 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (4.52 DPS) | yes | Defiler's Cloth Girdle (20166, +0.12 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -0.44 DPS) [dungeon]; Ghostweave Cord (254073, -0.59 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | sim-verified (+7.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.69 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -1.13 DPS) [vendor]; Spellshock Leggings (9484, -7.11 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.73 DPS) | yes | Gilded Sandals (254107, -0.64 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -3.23 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.39 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.49 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon]; Runed Ring (862, -1.68 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (3.36 DPS) | yes | Cyclopean Band (11824, -0.43 DPS) [dungeon]; Advisor's Ring (19520, -0.84 DPS) [rep]; Philanthropist's Ring (281635, -1.03 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.85 DPS) [crafted]; Rune of the Guard Captain (19120, -2.97 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.97 DPS, sim-verified) [crafted]; Rune of the Guard Captain (19120, -1.28 DPS) [quest] |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Moonshadow Stave (22458, -0.55 DPS) [quest]; Arbiter's Blade (11784, -3.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Woestave (20082, -0.08 DPS) [quest]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Pyric Caduceus (11748, -4.84 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; ranged: Noxious Shooter

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 388.9. Weights run: 0.9s. Verify run: 1.0s. 1003 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.250, intellect=0.107 ± 0.021, crit=0.030 ± 0.002 per rating point (14 rating = 1%, 0.421 per %), hit=0.201 ± 0.003 per rating point (10 rating = 1%, 2.006 per %), spell_haste=-3.778 ± 0.375, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.250

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 67.8 spell_power points (23.92 DPS) | yes | Warlord's Dreadweave Hood (17591, -11.73 DPS) [vendor]; Warlord's Dreadweave Hood (231590, -11.73 DPS) [pvp]; Crimson Felt Hat (18727, -27.67 DPS, sim-verified) [dungeon] |
| neck | Orb of the Darkmoon (19426) (or Chains of the Lich (23125)) | 1200 Tickets - Orb of the Darkmoon [quest] | 22.0 spell_power points (7.76 DPS) | yes | Chains of the Lich (23125, -0.20 DPS, sim-verified) [dungeon]; Arcane Crystal Pendant (20037, -1.89 DPS) [quest]; Amulet of the Dawn (22657, -1.98 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 43.8 spell_power points (15.46 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -5.22 DPS, sim-verified) [vendor]; Warlord's Dreadweave Mantle (17590, -6.00 DPS) [vendor]; Warlord's Dreadweave Mantle (231592, -6.00 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (388.9 DPS) | yes | Amplifying Cloak (18350, -0.30 DPS) [dungeon]; Hide of the Wild (18510, -1.34 DPS) [crafted]; Crystalline Threaded Cape (20697, -7.47 DPS, sim-verified) [world_drop] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 67.7 spell_power points (23.86 DPS) | yes | Robe of the Void (14153, -7.10 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (17592, -11.67 DPS) [vendor]; Warlord's Dreadweave Robe (231591, -11.67 DPS) [pvp] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 29.4 spell_power points (10.37 DPS) | yes | Dryad's Wrist Bindings (19595, -1.74 DPS, sim-verified) [rep]; Dryad's Wrist Bindings (19596, -3.09 DPS) [rep]; Dryad's Wrist Bindings (19597, -4.50 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 40.4 spell_power points (14.25 DPS) | yes | General's Dreadweave Gloves (17588, -3.44 DPS) [vendor]; General's Dreadweave Gloves (231589, -3.44 DPS) [pvp]; Sandworm Skin Gloves (20716, -3.80 DPS, sim-verified) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 44.2 spell_power points (15.60 DPS) | yes | Knowledge of the Timbermaw (228190, -7.14 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -7.80 DPS) [crafted]; Felheart Belt (16806, -7.98 DPS) [world_drop] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 59.7 spell_power points (21.04 DPS) | yes | Flarecore Leggings (19165, -6.44 DPS, sim-verified) [crafted]; General's Dreadweave Pants (17593, -7.28 DPS) [vendor]; General's Dreadweave Pants (231588, -7.28 DPS) [pvp] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 39.5 spell_power points (13.93 DPS) | yes | General's Dreadweave Boots (17586, -4.27 DPS) [vendor]; General's Dreadweave Boots (231593, -4.27 DPS) [pvp]; Earthen Silk Slippers (254013, -5.47 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234436, -0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234437, -0.39 DPS) [vendor]; Naglering (11669, -12.05 DPS, sim-verified) [dungeon] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Elemental Focus Band (20682, -0.79 DPS) [world]; Maiden's Circle (13001, -2.35 DPS) [world_drop]; Naglering (11669, -8.44 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+15.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, -0.87 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | 0.0 spell_power points (0.00 DPS) | yes | Glacial Blade (19099, +0.26 DPS, sim-verified) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -1.30 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 223.6 spell_power points (78.88 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -5.18 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.22 DPS) [dungeon]; Sparkling Crystal Wand (20672, -14.44 DPS) [world] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ritssyn's Ring of Chaos; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1003, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

