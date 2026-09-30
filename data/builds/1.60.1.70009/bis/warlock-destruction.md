# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 43.9. Weights run: 1.5s. Verify run: 1.2s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.102, intellect=0.267 ± 0.020, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.606 per %), hit=0.152 ± 0.001 per rating point (10 rating = 1%, 1.516 per %), spell_haste=not significant (0.109 ± 0.098), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.774 ± 0.102, fire_power=0.228 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.13 DPS) | yes | Shadow Goggles (4373, -3.28 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (1.39 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.64 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Feyscale Cloak (6632, +0.00 DPS) [dungeon]; Black Whelp Cloak (7283, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.59 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.3 spell_power points (1.19 DPS) | yes | Green Woolen Vest (2582, -0.44 DPS) [crafted]; Bloody Apron (6226, -0.44 DPS) [dungeon]; Gray Woolen Robe (2585, -1.63 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.3 spell_power points (0.25 DPS) | yes | Windsong Bangles (263336, -0.06 DPS) [quest]; Repurposed Hair Band (281256, -0.15 DPS) [quest]; Bright Bracers (3647, -0.22 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.32 DPS) | yes | Gnoll Casting Gloves (892, -0.27 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.41 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.84 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.1 spell_power points (0.95 DPS) | yes | Novice Ardent's Sash (253887, -0.43 DPS) [crafted]; Keller's Girdle (2911, -0.55 DPS) [world_drop]; Novice Arcanist's Sash (253885, -1.04 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk-threaded Trousers (1929, -0.11 DPS) [dungeon]; Rumpled Kilt (274741, -0.49 DPS) [vendor]; Abomination Skin Leggings (23173, -0.76 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (1.52 DPS) | yes | Feather Padded Treads (285345, -0.71 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.77 DPS) [crafted]; Pristine Boots (253889, -0.80 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (1.04 DPS) | yes | Sludge-Stained Band (286535, -0.48 DPS) [world]; Lavishly Jeweled Ring (1156, -0.74 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.89 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.94 DPS) | yes | Lavishly Jeweled Ring (1156, -0.64 DPS) [dungeon]; Sludge-Stained Band (286535, -0.69 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.79 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.7 spell_power points (0.50 DPS) | yes | Channeler's Staff (4437, -0.18 DPS, sim-verified) [world]; Lesser Staff of the Spire (1300, -0.20 DPS) [world_drop]; Staff of Westfall (2042, -0.25 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 121.5 spell_power points (22.87 DPS) | yes | Skycaller (12984, -2.39 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.58 DPS) [dungeon]; Sizzle Stick (8071, -4.28 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 66.9. Weights run: 1.5s. Verify run: 1.3s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.103, intellect=0.190 ± 0.009, crit=0.065 ± 0.003 per rating point (14 rating = 1%, 0.912 per %), hit=0.154 ± 0.002 per rating point (10 rating = 1%, 1.542 per %), spell_haste=0.727 ± 0.096, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.799 ± 0.103, fire_power=0.202 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.28 DPS) | yes | Silk Headband (7050, -0.57 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.62 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.62 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.1 spell_power points (1.69 DPS) | yes | Crystal Starfire Medallion (5003, -1.53 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.53 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.83 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (2.22 DPS) | yes | Invoker's Mantle (215365, -0.57 DPS) [crafted]; Fairywing Mantle (9536, -0.62 DPS) [quest]; Death Speaker Mantle (6685, -0.79 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-verified (66.9 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Prelacy Cape (7004, +0.00 DPS) [quest]; Hillman's Cloak (3719, -0.68 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.70 DPS) | yes | Green Silk Armor (7065, -0.18 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.81 DPS) [dungeon]; Pristine Gown (253961, -0.97 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.87 DPS) | yes | Nightsky Wristbands (6407, -1.63 DPS) [world_drop]; Windsong Bangles (263336, -1.66 DPS) [quest]; Glowing Magical Bracelets (13106, -1.77 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.45 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Town Clerk's Mittens (270029, -0.19 DPS) [quest]; Gnoll Casting Gloves (892, -0.21 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.6 spell_power points (2.40 DPS) | yes | Belt of Arugal (6392, -0.53 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.74 DPS) [dungeon]; Invoker's Cord (215366, -0.75 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.49 DPS) | yes | Abomination Skin Leggings (23173, -0.06 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.76 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.01 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.3 spell_power points (1.73 DPS) | yes | Acidic Walkers (9454, -0.38 DPS) [dungeon]; Nimbus Boots (6998, -0.48 DPS) [quest]; Spidersilk Boots (4320, -1.95 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.45 DPS) | yes | Minor Channeling Ring (1449, -0.34 DPS) [quest]; Lorekeeper's Ring (20431, -0.42 DPS) [rep]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.25 DPS) | yes | Electrocutioner Lagnut (9447, -0.62 DPS) [dungeon]; Sludge-Stained Band (286535, -0.62 DPS) [world]; Minor Channeling Ring (1449, -1.65 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.87 DPS) | yes | Glimmering Staff (249392, -1.15 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.47 DPS) [world_drop]; Channeler's Staff (4437, -1.55 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.1 spell_power points (1.69 DPS) | yes | Eye of Paleth (2943, -0.86 DPS) [quest]; Orb of Souls (249395, -0.86 DPS) [crafted]; Dwarven Tome (279898, -1.21 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 163.0 spell_power points (33.83 DPS) | yes | Starfaller (13063, -0.83 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.79 DPS) [crafted]; Gravestone Scepter (7001, -4.83 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 107.5. Weights run: 1.4s. Verify run: 1.1s. 329 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.661, intellect=1.306 ± 0.042, crit=0.233 ± 0.011 per rating point (14 rating = 1%, 3.260 per %), hit=0.588 ± 0.006 per rating point (10 rating = 1%, 5.884 per %), spell_haste=not significant (1.377 ± 0.598), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.268 ± 0.661), fire_power=1.253 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Miner's Hat of the Deep (9429, -0.17 DPS) [dungeon]; Papal Fez (9431, -0.17 DPS) [dungeon]; Corpseshroud (10574, -2.41 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.8 spell_power points (1.34 DPS) | yes | Necklace of Calisea (1714, -0.51 DPS) [world_drop]; Triune Amulet (7722, -0.51 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.44 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.0 spell_power points (2.16 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.29 DPS) [dungeon]; Death Speaker Mantle (6685, -0.33 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 20.8 spell_power points (1.87 DPS) | yes | Long Silken Cloak (4326, -0.74 DPS) [crafted]; Guardian Cloak (5965, -0.74 DPS) [crafted]; Darkspear Raider's Cloak (272077, -2.97 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.8 spell_power points (2.69 DPS) | yes | Robe of Power (7054, -0.01 DPS) [crafted]; Green Silk Armor (7065, -0.35 DPS) [crafted]; Dreamweave Vest (10021, -0.95 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 11.8 spell_power points (1.06 DPS) | yes | Mistscape Bracers (4045, -0.11 DPS, sim-verified) [world_drop]; Aurora Bracers (4043, -0.12 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.12 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Gloves (10011, -0.32 DPS) [crafted]; Town Clerk's Mittens (270029, -0.44 DPS) [quest]; Red Mageweave Gloves (10018, -1.35 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -0.07 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.32 DPS) [quest]; Deathmage Sash (10771, -1.12 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.7 spell_power points (2.68 DPS) | yes | Abomination Skin Leggings (23173, -0.92 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.10 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.38 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.16 DPS) | yes | Acidic Walkers (9454, -0.77 DPS) [dungeon]; Spidersilk Boots (4320, -1.06 DPS) [crafted]; Gilded Slippers (254001, -1.99 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.8 spell_power points (1.61 DPS) | yes | Ogremind Ring (1993, -0.78 DPS) [world_drop]; Mindbender Loop (5009, -0.78 DPS) [world_drop]; Snake Hoop (6750, -0.78 DPS) [quest] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009), Snake Hoop (6750), Black Widow Band (6199)) | World drop [world_drop] | 9.1 spell_power points (0.82 DPS) | yes | Ogremind Ring (1993, +0.00 DPS, sim-verified) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Snake Hoop (6750, +0.00 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 0.0 spell_power points (0.00 DPS) | yes | Spellforce Rod (1664, -0.02 DPS) [world_drop]; Windweaver Staff (7757, -0.06 DPS) [dungeon]; Gut Ripper (2164, -1.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 444.4 spell_power points (40.09 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.50 DPS) [dungeon]; Twisted Nether Wand (249144, -5.55 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Voodoo Band; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 329, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 153.0. Weights run: 1.4s. Verify run: 1.2s. 419 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.605, intellect=1.132 ± 0.038, crit=0.187 ± 0.010 per rating point (14 rating = 1%, 2.617 per %), hit=0.418 ± 0.005 per rating point (10 rating = 1%, 4.179 per %), spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.6 spell_power points (6.45 DPS) | yes | Dreamweave Circlet (10041, -0.14 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.72 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Hat (220889, -1.77 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 22.8 spell_power points (3.53 DPS) | yes | Scorn's Icy Choker (23169, -1.39 DPS) [dungeon]; Mindburst Medallion (11196, -1.55 DPS) [quest]; Horizon Choker (13085, -5.54 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.4 spell_power points (5.17 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.46 DPS) [crafted]; Inquisitor's Shawl (19507, -1.81 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.8 spell_power points (3.22 DPS) | yes | Runecloth Cloak (13860, -0.42 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -0.62 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272076, -0.77 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.6 spell_power points (6.45 DPS) | yes | Runecloth Tunic (13857, -1.89 DPS) [crafted]; Robe of the Magi (1716, -1.99 DPS) [world_drop]; Runecloth Robe (13858, -2.80 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Bloodband Bracers (11469, +0.00 DPS, sim-verified) [quest]; Nethergeld Cuffs (254061, -0.32 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.53 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 31.9 spell_power points (4.94 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -0.87 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.45 DPS) [crafted]; Red Mageweave Gloves (10018, -1.48 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 27.5 spell_power points (4.26 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.41 DPS) [dungeon]; Deathmage Sash (10771, -0.55 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.3 spell_power points (5.32 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.70 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -4.01 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.72 DPS) | yes | Gilded Sandals (254107, -0.44 DPS) [crafted]; Southsea Mojo Boots (20641, -0.55 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -2.70 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Philanthropist's Ring (281635, -0.03 DPS) [quest]; Mindseye Circle (10634, -0.53 DPS) [dungeon]; Band of the Unicorn (7553, -0.62 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.9 spell_power points (2.62 DPS) | yes | Mindseye Circle (10634, -0.52 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Philanthropist's Ring (281635, -1.10 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (153.0 DPS) | yes | Frozen Heart of the Mountain (249469, -1.28 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (153.0 DPS) | yes | Frozen Heart of the Mountain (249469, -0.19 DPS, sim-verified) [crafted] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (153.0 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Soul Harvester (20536, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -2.30 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 338.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.73 DPS) [dungeon]; Woestave (20082, -2.30 DPS, sim-verified) [quest]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; ranged: Pyric Caduceus

No-known-source sample (15 of 419, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 328.6. Weights run: 1.4s. Verify run: 1.3s. 1007 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± -5.700), intellect=not significant (-11.635 ± -0.334), crit=not significant (-2.647 ± -0.119) per rating point (14 rating = 1%, -37.059 per %), hit=not significant (-5.135 ± -0.053) per rating point (10 rating = 1%, -51.347 per %), spell_haste=not significant (-13.428 ± -4.422), spell_penetration=not significant (-0.000 ± -0.000), shadow_power=not significant (8.066 ± -5.701), fire_power=not significant (-7.192 ± -0.026)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | sim-verified (+26.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Sun Shades (208424, -26.93 DPS, sim-verified) [vendor] |
| neck | - | - |  |  |  |
| shoulder | - | - |  |  |  |
| back | - | - |  |  |  |
| chest | - | - |  |  |  |
| wrist | - | - |  |  |  |
| hands | - | - |  |  |  |
| waist | - | - |  |  |  |
| legs | - | - |  |  |  |
| feet | - | - |  |  |  |
| finger1 | Signet Ring of the Bronze Dragonflight (21200) | The Protector of Kalimdor [quest] | 0.0 spell_power points | yes | Naglering (11669, +0.00 DPS, sim-verified) [dungeon] |
| finger2 | Cauterizing Band (19140) | World drop [world_drop] | 0.0 spell_power points | yes | Naglering (11669, -10.33 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+10.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | 0.0 spell_power points | yes | Burst of Knowledge (11832, -3.30 DPS, sim-verified) [dungeon] |
| main_hand | Runesword of the Red (21521) | Treasure of the Timeless One [quest] | 0.0 spell_power points | yes | Teebu's Blazing Longsword (1728, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | sim-verified (+8.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Tome of the Ice Lord (19310, -8.37 DPS, sim-verified) [rep] |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Torch of Light (279246, -4.92 DPS, sim-verified) [crafted] |

**New at 60:** head: Heretic Cowl; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cauterizing Band; trinket1: Talisman of Ascendance; trinket2: Draconic Infused Emblem; main_hand: Runesword of the Red; off_hand: Lei of the Lifegiver; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1007, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 42.3. Weights run: 1.5s. Verify run: 1.2s. 142 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.102, intellect=0.267 ± 0.020, crit=0.043 ± 0.002 per rating point (14 rating = 1%, 0.606 per %), hit=0.152 ± 0.001 per rating point (10 rating = 1%, 1.516 per %), spell_haste=not significant (0.109 ± 0.098), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.774 ± 0.102, fire_power=0.228 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (1.13 DPS) | yes | Shadow Goggles (4373, -3.12 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (1.39 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.12 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.64 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Feyscale Cloak (6632, +0.00 DPS) [dungeon]; Black Whelp Cloak (7283, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.73 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.3 spell_power points (1.19 DPS) | yes | Green Woolen Vest (2582, -0.44 DPS) [crafted]; Bloody Apron (6226, -0.44 DPS) [dungeon]; Gray Woolen Robe (2585, -1.35 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 1.6 spell_power points (0.30 DPS) | yes | Owlbeard Bracers (16981, +0.00 DPS, sim-verified) [quest]; Mindthrust Bracers (1974, -0.05 DPS) [dungeon]; Featherbead Bracers (15452, -0.05 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.32 DPS) | yes | Pristine Gloves (253913, -0.41 DPS) [crafted]; Gnoll Casting Gloves (892, -0.41 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.56 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.1 spell_power points (0.95 DPS) | yes | Novice Ardent's Sash (253887, -0.43 DPS) [crafted]; Keller's Girdle (2911, -0.55 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.87 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk-threaded Trousers (1929, -0.11 DPS) [dungeon]; Rumpled Kilt (274741, -0.49 DPS) [vendor]; Abomination Skin Leggings (23173, -0.83 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (1.52 DPS) | yes | Feather Padded Treads (285345, -0.38 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.77 DPS) [crafted]; Pristine Boots (253889, -0.80 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.94 DPS) | yes | Lavishly Jeweled Ring (1156, -0.64 DPS) [dungeon]; Loop of Sacrifice (281673, -0.69 DPS) [quest]; Volcanic Rock Ring (12053, -0.79 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.56 DPS) | yes | Loop of Sacrifice (281673, -0.31 DPS) [quest]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.69 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 2.7 spell_power points (0.50 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.10 DPS) [world]; Lesser Staff of the Spire (1300, -0.20 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 121.5 spell_power points (22.87 DPS) | yes | Skycaller (12984, -2.29 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.58 DPS) [dungeon]; Sizzle Stick (8071, -4.28 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 142, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 65.0. Weights run: 1.5s. Verify run: 1.2s. 237 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.103, intellect=0.190 ± 0.009, crit=0.065 ± 0.003 per rating point (14 rating = 1%, 0.912 per %), hit=0.154 ± 0.002 per rating point (10 rating = 1%, 1.542 per %), spell_haste=0.727 ± 0.096, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.799 ± 0.103, fire_power=0.202 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.28 DPS) | yes | Embalmed Shroud (7691, -0.62 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.62 DPS) [crafted]; Silk Headband (7050, -1.26 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.1 spell_power points (1.69 DPS) | yes | Crystal Starfire Medallion (5003, -1.53 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.53 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.31 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (2.22 DPS) | yes | Chestnut Mantle (17695, -0.56 DPS) [quest]; Invoker's Mantle (215365, -0.57 DPS) [crafted]; Death Speaker Mantle (6685, -1.04 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.04 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted]; Battle Healer's Cloak (19529, -0.21 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (2.70 DPS) | yes | Green Silk Armor (7065, -0.45 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.81 DPS) [dungeon]; High Robe of the Adjudicator (3461, -0.96 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.87 DPS) | yes | Owlbeard Bracers (16981, -1.58 DPS) [quest]; Nightsky Wristbands (6407, -1.63 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.18 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.45 DPS) | yes | Gnoll Casting Gloves (892, -0.21 DPS) [world]; Truefaith Gloves (7049, -0.30 DPS) [crafted]; Jutebraid Gloves (10654, -0.34 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.6 spell_power points (2.40 DPS) | yes | Belt of Arugal (6392, -0.42 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.74 DPS) [dungeon]; Warsong Sash (16975, -1.13 DPS, sim-verified) [quest] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (2.49 DPS) | yes | Abomination Skin Leggings (23173, -0.60 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.76 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.01 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.3 spell_power points (1.73 DPS) | yes | Acidic Walkers (9454, -0.38 DPS) [dungeon]; Boots of the Enchanter (4325, -0.69 DPS) [crafted]; Spidersilk Boots (4320, -2.50 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.45 DPS) | yes | Advisor's Ring (20426, -0.42 DPS) [rep]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Sludge-Stained Band (286535, -0.83 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.25 DPS) | yes | Sludge-Stained Band (286535, -0.62 DPS) [world]; Sacred Band (6669, -0.83 DPS) [quest]; Electrocutioner Lagnut (9447, -3.00 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.87 DPS) | yes | Twisted Chanter's Staff (890, -1.47 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.47 DPS) [quest]; Glimmering Staff (249392, -1.72 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.1 spell_power points (1.69 DPS) | yes | Alliance Outrunner Healing Rod (285348, -0.86 DPS) [world]; Orb of Souls (249395, -0.88 DPS, sim-verified) [crafted]; Tome of the Darkspear Prophecy (272090, -1.12 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 163.0 spell_power points (33.83 DPS) | yes | Starfaller (13063, -0.83 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.79 DPS) [crafted]; Gravestone Scepter (7001, -4.83 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 237, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 107.4. Weights run: 1.4s. Verify run: 1.1s. 320 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.661, intellect=1.306 ± 0.042, crit=0.233 ± 0.011 per rating point (14 rating = 1%, 3.260 per %), hit=0.588 ± 0.006 per rating point (10 rating = 1%, 5.884 per %), spell_haste=not significant (1.377 ± 0.598), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.268 ± 0.661), fire_power=1.253 ± 0.006

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Miner's Hat of the Deep (9429, -0.17 DPS) [dungeon]; Papal Fez (9431, -0.17 DPS) [dungeon]; Corpseshroud (10574, -2.25 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.8 spell_power points (1.34 DPS) | yes | Necklace of Calisea (1714, -0.51 DPS) [world_drop]; Triune Amulet (7722, -0.51 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.39 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.0 spell_power points (2.16 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.29 DPS) [dungeon]; Death Speaker Mantle (6685, -0.33 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 20.8 spell_power points (1.87 DPS) | yes | Long Silken Cloak (4326, -0.74 DPS) [crafted]; Guardian Cloak (5965, -0.74 DPS) [crafted]; Darkspear Raider's Cloak (272077, -2.88 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.8 spell_power points (2.69 DPS) | yes | Robe of Power (7054, -0.01 DPS) [crafted]; Green Silk Armor (7065, -0.35 DPS) [crafted]; Dreamweave Vest (10021, -0.76 DPS, sim-verified) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 14.4 spell_power points (1.30 DPS) | yes | Mistscape Bracers (4045, -0.36 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.36 DPS) [quest]; Windchaser Cuffs (14429, -1.08 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Gloves (10011, -0.32 DPS) [crafted]; Gilded Handwraps (254021, -0.55 DPS) [crafted]; Red Mageweave Gloves (10018, -1.68 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Cord (254037, -0.07 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.32 DPS) [quest]; Deathmage Sash (10771, -1.33 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.7 spell_power points (2.68 DPS) | yes | Abomination Skin Leggings (23173, -0.92 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.10 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.63 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.16 DPS) | yes | Acidic Walkers (9454, -0.77 DPS) [dungeon]; Spidersilk Boots (4320, -1.06 DPS) [crafted]; Gilded Slippers (254001, -2.44 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.8 spell_power points (1.61 DPS) | yes | Ogremind Ring (1993, -0.78 DPS) [world_drop]; Mindbender Loop (5009, -0.78 DPS) [world_drop]; Snake Hoop (6750, -0.78 DPS) [quest] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009), Snake Hoop (6750), Black Widow Band (6199)) | World drop [world_drop] | 9.1 spell_power points (0.82 DPS) | yes | Ogremind Ring (1993, +0.00 DPS, sim-verified) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Snake Hoop (6750, +0.00 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 0.0 spell_power points (0.00 DPS) | yes | Spellforce Rod (1664, -0.02 DPS) [world_drop]; Windweaver Staff (7757, -0.06 DPS) [dungeon]; Gut Ripper (2164, -1.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 444.4 spell_power points (40.09 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.50 DPS) [dungeon]; Twisted Nether Wand (249144, -5.55 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Voodoo Band; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 320, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 147.3. Weights run: 1.4s. Verify run: 1.2s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.605, intellect=1.132 ± 0.038, crit=0.187 ± 0.010 per rating point (14 rating = 1%, 2.617 per %), hit=0.418 ± 0.005 per rating point (10 rating = 1%, 4.179 per %), spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.6 spell_power points (6.45 DPS) | yes | Dreamweave Circlet (10041, -0.24 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.72 DPS) [dungeon]; Blood Guard's Dreadweave Hat (220907, -1.77 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 22.8 spell_power points (3.53 DPS) | yes | Scorn's Icy Choker (23169, -1.39 DPS) [dungeon]; Mindburst Medallion (11196, -1.55 DPS) [quest]; Horizon Choker (13085, -5.01 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.4 spell_power points (5.17 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.46 DPS) [crafted]; Inquisitor's Shawl (19507, -1.81 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.2 spell_power points (3.44 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.46 DPS) [dungeon]; Runecloth Cloak (13860, -0.64 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.6 spell_power points (6.45 DPS) | yes | Runecloth Tunic (13857, -1.89 DPS) [crafted]; Robe of the Magi (1716, -1.99 DPS) [world_drop]; Runecloth Robe (13858, -2.53 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Bloodband Bracers (11469, +0.00 DPS, sim-verified) [quest]; Nethergeld Cuffs (254061, -0.32 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.53 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 31.9 spell_power points (4.94 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.12 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.45 DPS) [crafted]; Red Mageweave Gloves (10018, -1.48 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 27.5 spell_power points (4.26 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.41 DPS) [dungeon]; Deathmage Sash (10771, -0.55 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.3 spell_power points (5.32 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.70 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -4.65 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.72 DPS) | yes | Gilded Sandals (254107, -0.44 DPS) [crafted]; Southsea Mojo Boots (20641, -0.55 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -2.45 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Philanthropist's Ring (281635, -0.03 DPS) [quest]; Mindseye Circle (10634, -0.53 DPS) [dungeon]; Band of the Unicorn (7553, -0.62 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.9 spell_power points (2.62 DPS) | yes | Mindseye Circle (10634, -0.52 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Philanthropist's Ring (281635, -1.07 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (147.3 DPS) | yes | Frozen Heart of the Mountain (249469, -1.28 DPS) [crafted]; Rune of the Guard Captain (19120, -1.41 DPS) [quest]; Uther's Strength (11302, -1.76 DPS, sim-verified) [world_drop] |
| trinket2 | Fire Ruby (20036) | Destroy Morphaz [quest] | sim-verified (147.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (147.3 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -1.05 DPS) [quest]; Soul Harvester (20536, -1.63 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 338.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.73 DPS) [dungeon]; Woestave (20082, -1.93 DPS, sim-verified) [quest]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Abyss Shard; trinket2: Fire Ruby; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 327.4. Weights run: 1.4s. Verify run: 1.3s. 998 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± -5.700), intellect=not significant (-11.635 ± -0.334), crit=not significant (-2.647 ± -0.119) per rating point (14 rating = 1%, -37.059 per %), hit=not significant (-5.135 ± -0.053) per rating point (10 rating = 1%, -51.347 per %), spell_haste=not significant (-13.428 ± -4.422), spell_penetration=not significant (-0.000 ± -0.000), shadow_power=not significant (8.066 ± -5.701), fire_power=not significant (-7.192 ± -0.026)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | sim-verified (+27.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Sun Shades (208424, -27.84 DPS, sim-verified) [vendor] |
| neck | - | - |  |  |  |
| shoulder | - | - |  |  |  |
| back | - | - |  |  |  |
| chest | - | - |  |  |  |
| wrist | - | - |  |  |  |
| hands | - | - |  |  |  |
| waist | - | - |  |  |  |
| legs | - | - |  |  |  |
| feet | - | - |  |  |  |
| finger1 | Signet Ring of the Bronze Dragonflight (21200) | The Protector of Kalimdor [quest] | 0.0 spell_power points | yes | Naglering (11669, +0.00 DPS, sim-verified) [dungeon] |
| finger2 | Cauterizing Band (19140) | World drop [world_drop] | 0.0 spell_power points | yes | Naglering (11669, -10.84 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+9.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points | yes | Burst of Knowledge (11832, -4.01 DPS, sim-verified) [dungeon] |
| main_hand | Runesword of the Red (21521) | Treasure of the Timeless One [quest] | 0.0 spell_power points | yes | Teebu's Blazing Longsword (1728, +0.00 DPS, sim-verified) [world_drop] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | sim-verified (+7.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Tome of the Ice Lord (19310, -7.36 DPS, sim-verified) [rep] |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+7.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Torch of Light (279246, -7.39 DPS, sim-verified) [crafted] |

**New at 60:** head: Heretic Cowl; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cauterizing Band; trinket1: Draconic Infused Emblem; trinket2: Talisman of Ascendance; main_hand: Runesword of the Red; off_hand: Lei of the Lifegiver; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 998, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

